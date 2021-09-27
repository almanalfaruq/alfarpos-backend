package response

import (
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/golang-jwt/jwt/v4"
)

type TokenData struct {
	User userentity.User `json:"user"`
	jwt.StandardClaims
}
