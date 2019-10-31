package response

import (
	"../../model"
	"github.com/dgrijalva/jwt-go"
)

type TokenResponse struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}
