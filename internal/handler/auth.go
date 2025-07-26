package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/fadhilkholaf/go-gorm/internal/config"
	"github.com/fadhilkholaf/go-gorm/internal/model"
)

func (h *Handler) Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	var user model.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding request!",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	p, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error hashing password!",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	r := h.db.WithContext(ctx).Create(&model.User{Username: user.Username, Password: string(p)})

	var pgErr *pgconn.PgError

	if r.Error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error request timeout!",
				"data":    nil,
				"error":   r.Error.Error(),
			})
			return
		}

		isErr := errors.As(r.Error, &pgErr)

		if isErr {
			switch pgErr.Code {
			case "23505":
				log.Println(pgErr.Error())
			default:
				log.Println("err skipped")
			}
			return
		}

		if r.Error == gorm.ErrDuplicatedKey {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error username has been taken!",
				"data":    nil,
				"error":   r.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unexpected error registering user!",
			"data":    nil,
			"error":   r.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register success.",
		"data":    nil,
		"error":   nil,
	})
}

func (h *Handler) LogIn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	var body model.User
	var user model.User

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding request!",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	r := h.db.WithContext(ctx).First(&user, "username = ?", body.Username)

	if r.Error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"message": "Error request timeout!",
				"data":    nil,
				"error":   r.Error.Error(),
			})
			return
		}

		if r.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Incorrect username or password!",
				"data":    nil,
				"error":   r.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unexpected error finding user!",
			"data":    nil,
			"error":   r.Error.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Incorrect username or password!",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &config.JwtClaims{
		Role: user.Role,
		RegisteredClaims: &jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(user.ID), 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	s, err := t.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error signin token!",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	c.SetCookie("token", s, 28800, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Log in success.",
		"data": gin.H{
			"token": s,
		},
		"error": nil,
	})
}

func (*Handler) LogOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Log out success.",
		"data":    nil,
		"error":   nil,
	})
}
