package middleware

import (
	"errors"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/fadhilkholaf/go-gorm/internal/config"
	"github.com/fadhilkholaf/go-gorm/internal/model"
)

func Auth(roles []model.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		s, err := c.Cookie("token")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Error token not found!",
				"data":    nil,
				"error":   err.Error(),
			})
			return
		}

		t, err := jwt.ParseWithClaims(s, &config.JwtClaims{}, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		}, jwt.WithLeeway(5*time.Second))

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Error token expired!",
					"data":    nil,
					"error":   err.Error(),
				})
				return
			}

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Error parsing token!",
				"data":    nil,
				"error":   err.Error(),
			})
			return
		}

		claims, ok := t.Claims.(*config.JwtClaims)

		if !ok || !t.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Error parsing claims!",
				"data":    nil,
				"error":   nil,
			})
			return
		}

		authorized := slices.Contains(roles, claims.Role)

		if !authorized {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Error unauthorized role!",
				"data":    nil,
				"error":   nil,
			})
			return
		}

		c.Next()
	}
}
