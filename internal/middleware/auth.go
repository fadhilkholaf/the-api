package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var IsAdmin = "ADMIN"

func Auth(role *string) gin.HandlerFunc {
	log.Println("Auth middleware triggered.")

	return func(c *gin.Context) {
		if role == nil {
			c.Next()
			return
		}

		s, err := c.Cookie("token")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Error token not found!",
				"data":    nil,
				"error":   err.Error(),
			})
			return
		}

		log.Println(s)

		t, err := jwt.ParseWithClaims(s, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			log.Fatalf("Error parsing token: %s", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		log.Println(t)

		if *role != "ADMIN" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
