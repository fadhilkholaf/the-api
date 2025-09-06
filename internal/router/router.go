package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fadhilkholaf/go-gorm/internal/handler"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.New()
	h := handler.NewHandler(db)

	r.GET("/status", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.Use(gin.Logger(), gin.Recovery())
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.Static("/public", "./public")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "The API.",
			"data":    nil,
			"error":   nil,
		})
	})

	userRoute(r, h)
	authRoute(r, h)
	postRoute(r, h)

	return r
}
