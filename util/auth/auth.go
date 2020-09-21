package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/util"
	render "github.com/almanalfaruq/alfarpos-backend/util/response"
	"github.com/dgrijalva/jwt-go"
)

type AuthMw struct {
	secretKey string
}

func New(cfg util.Config) *AuthMw {
	return &AuthMw{
		secretKey: cfg.SecretKey,
	}
}

type TokenData struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

func (m *AuthMw) CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tokenData TokenData
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			err := errors.New("Missing authorization header")
			render.RenderJSONError(w, http.StatusUnauthorized, err)
			return
		}
		authToken := strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwt.ParseWithClaims(authToken, &tokenData, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return []byte(m.secretKey), nil
		})
		if err != nil {
			render.RenderJSONError(w, http.StatusUnauthorized, err)
			return
		}
		claims, ok := token.Claims.(*TokenData)
		if !ok || !token.Valid {
			err = errors.New("Failed parsing token data")
			render.RenderJSONError(w, http.StatusUnauthorized, err)
			return
		}

		if claims.User.ID == 0 {
			err = errors.New("User data is null")
			render.RenderJSONError(w, http.StatusUnauthorized, err)
			return
		}

		userCtx := context.WithValue(r.Context(), model.CTX_USER, claims.User)
		r = r.WithContext(userCtx)
		next(w, r)
	}
}
