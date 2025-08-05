package router

import (
	"github.com/gin-gonic/gin"

	"github.com/fadhilkholaf/go-gorm/internal/handler"
	"github.com/fadhilkholaf/go-gorm/internal/middleware"
	"github.com/fadhilkholaf/go-gorm/internal/model"
)

func userRoute(r *gin.Engine, h *handler.Handler) {
	u := r.Group("/user")

	authorizedRoute := u.Group("/")
	authorizedRoute.Use(middleware.Auth([]model.Role{model.AdminRole, model.UserRole}))
	authorizedRoute.GET("/", h.FindUser)
	authorizedRoute.GET("/:id", h.FirstUser)

	adminRoute := u.Group("/")
	adminRoute.Use(middleware.Auth([]model.Role{model.AdminRole}))
	adminRoute.POST("/", h.CreateUser)
	adminRoute.PUT("/:id", h.UpdateUser)
	adminRoute.DELETE("/:id", h.DeleteUser)
}
