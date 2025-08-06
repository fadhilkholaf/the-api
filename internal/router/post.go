package router

import (
	"github.com/fadhilkholaf/go-gorm/internal/handler"
	"github.com/fadhilkholaf/go-gorm/internal/middleware"
	"github.com/fadhilkholaf/go-gorm/internal/model"
	"github.com/gin-gonic/gin"
)

func postRoute(r *gin.Engine, h *handler.Handler) {
	p := r.Group("/post")

	adminRoute := p.Group("/")
	adminRoute.Use(middleware.Auth([]model.Role{model.AdminRole}))
	adminRoute.POST("/", h.CreatePost)
}
