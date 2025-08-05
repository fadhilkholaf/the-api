package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fadhilkholaf/go-gorm/internal/handler"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := handler.NewHandler(db)

	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "The Go api.",
			"data":    nil,
			"error":   nil,
		})
	})

	userRoute(r, h)
	authRoute(r, h)

	return r
}
