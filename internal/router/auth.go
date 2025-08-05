package router

import (
	"github.com/gin-gonic/gin"

	"github.com/fadhilkholaf/go-gorm/internal/handler"
	"github.com/fadhilkholaf/go-gorm/internal/middleware"
	"github.com/fadhilkholaf/go-gorm/internal/model"
)

func authRoute(r *gin.Engine, h *handler.Handler) {
	a := r.Group("/auth")

	authenticated := a.Group("/")
	authenticated.Use(middleware.Auth([]model.Role{model.AdminRole, model.UserRole}))
	authenticated.POST("/logout", h.LogOut)

	a.POST("/register", h.Register)
	a.POST("/login", h.LogIn)
}
