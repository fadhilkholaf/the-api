package model

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type Model struct {
	Id        uint64         `gorm:"column:id;primaryKey;autoIncrement;not null;uniqueIndex;" json:"id"`
	CreatedAt time.Time      `gorm:"column:createdAt;not null;" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updatedAt;" json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt;index;" json:"deletedAt,omitempty"`
}

type User struct {
	Model
	Username string `gorm:"column:username;not null;uniqueIndex;" json:"username"`
	Password string `gorm:"column:password;not null;" json:"password"`
	Role     Role   `gorm:"column:role;not null;" json:"role"`
	Posts    []Post `gorm:"foreignKey:authorId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"posts,omitempty"`
}

type Post struct {
	Model
	Title       string `gorm:"column:title;not null;" json:"title"`
	Description string `gorm:"column:description;not null;" json:"description"`
	AuthorId    uint64 `gorm:"column:authorId;not null;index;" json:"authorId"`
	Author      User   `gorm:"foreignKey:authorId;references:id;" json:"author,omitempty"`
}
