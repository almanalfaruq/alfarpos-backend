package response

import (
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/dgrijalva/jwt-go"
)

type TokenData struct {
	User userentity.User `json:"user"`
	jwt.StandardClaims
}
