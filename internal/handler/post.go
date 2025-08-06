package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/fadhilkholaf/go-gorm/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	var post model.Post

	err := c.ShouldBindJSON(&post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding request!",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	h.db.WithContext(ctx).Create(&post)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created.",
		"data":    "data",
		"error":   nil,
	})
}
