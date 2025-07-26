package router

import (
	"github.com/gin-gonic/gin"

	"github.com/fadhilkholaf/go-gorm/internal/handler"
)

func userRoute(r *gin.Engine, h *handler.Handler) {
	u := r.Group("/user")

	u.POST("/", h.CreateUser)
	u.GET("/", h.FindUser)
	u.GET("/:id", h.FirstUser)
	u.PUT("/:id", h.UpdateUser)
	u.DELETE("/:id", h.DeleteUser)
}
