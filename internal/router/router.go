package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fadhilkholaf/go-gorm/internal/handler"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := handler.NewHandler(db)

	userRoute(r, h)
	authRoute(r, h)

	return r
}
