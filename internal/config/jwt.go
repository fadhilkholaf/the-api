package config

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/fadhilkholaf/go-gorm/internal/model"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	Role model.Role `json:"role"`
}
