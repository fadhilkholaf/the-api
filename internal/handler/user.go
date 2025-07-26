package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/fadhilkholaf/go-gorm/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user model.User

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

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

	if r.Error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"message": "Error request timeout!",
				"data":    nil,
				"error":   r.Error.Error(),
			})
			return
		}

		if r.Error == gorm.ErrDuplicatedKey {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Error conflict.",
				"data":    nil,
				"error":   r.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating user!",
			"data":    nil,
			"error":   r.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created.",
		"data":    user,
		"error":   nil,
	})
}

func (h *Handler) FindUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	var users []model.User

	r := h.db.WithContext(ctx).Find(&users)

	if r.Error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"message": "Error request timeout!",
				"data":    nil,
				"error":   r.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating user!",
			"data":    nil,
			"error":   r.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users retrieved.",
		"data":    users,
		"error":   nil,
	})
}

func (h *Handler) FirstUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	var user model.User

	r := h.db.WithContext(ctx).First(&user, c.Param("id"))

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
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Error user not found!",
				"data":    nil,
				"error":   r.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error finding user!",
			"data":    nil,
			"error":   r.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User retrieved.",
		"data":    user,
		"error":   nil,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
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

	r := h.db.WithContext(ctx).Model(&user).Where("id = ?", c.Param("id")).Updates(&model.User{Username: user.Username, Password: string(p)})

	if r.Error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"message": "Error request timeout!",
				"data":    nil,
				"error":   r.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating user!",
			"data":    nil,
			"error":   r.Error.Error(),
		})
		return
	}

	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error user not found!",
			"data":    nil,
			"error":   "0 rows affected!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated.",
		"data":    user,
		"error":   nil,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	r := h.db.WithContext(ctx).Delete(&model.User{}, c.Param("id"))

	if r.Error != nil {
		if ctx.Err() == context.DeadlineExceeded {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"message": "Error request timeout!",
				"data":    nil,
				"error":   r.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting user!",
			"data":    nil,
			"error":   r.Error.Error(),
		})
		return
	}

	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error user not found!",
			"data":    nil,
			"error":   "0 rows affected!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted.",
		"data":    nil,
		"error":   nil,
	})
}
