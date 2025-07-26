package router

import (
	"github.com/gin-gonic/gin"

	"github.com/fadhilkholaf/go-gorm/internal/handler"
	"github.com/fadhilkholaf/go-gorm/internal/middleware"
)

func authRoute(r *gin.Engine, h *handler.Handler) {
	a := r.Group("/auth")

	a.POST("/register", h.Register)
	a.POST("/login", h.LogIn)
	a.POST("/logout", middleware.Auth(&middleware.IsAdmin), h.LogOut)
}
