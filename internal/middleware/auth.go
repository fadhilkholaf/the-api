package middleware

import (
	"net/http"
	"os"
	"slices"

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
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Error parsing token!",
				"data":    nil,
				"error":   err.Error(),
			})
			return
		}

		claims, ok := t.Claims.(*config.JwtClaims)

		if !ok {
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
