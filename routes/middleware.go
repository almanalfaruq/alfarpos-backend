package routes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
)

type AuthConfig struct {
	SecretKey string `yaml:"secret_key"`
}

type AuthMiddleware struct {
	secretKey string
}

func New(cfg AuthConfig) *AuthMiddleware {
	return &AuthMiddleware{
		secretKey: cfg.SecretKey,
	}
}

func (m *AuthMiddleware) CheckJWTToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		authHeaderSplit := strings.Split(authHeader, " ")
		if len(authHeaderSplit) != 2 {
			m.renderJSONError(w, http.StatusUnauthorized, errors.New("Missing header Authorization"), "Missing header Authorization")
			return
		}

		if strings.ToLower(authHeaderSplit[0]) != "bearer" {
			m.renderJSONError(w, http.StatusUnauthorized, errors.New("Missing header Authorization"), "Missing header Authorization")
			return
		}

		token, err := jwt.ParseWithClaims(authHeaderSplit[1], &response.TokenResponse{}, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return []byte(m.secretKey), nil
		})
		if err != nil {
			m.renderJSONError(w, http.StatusUnauthorized, err, err.Error())
		}
		claims, ok := token.Claims.(*response.TokenResponse)
		if !ok || !token.Valid {
			m.renderJSONError(w, http.StatusUnauthorized, errors.New("Token is not valid"), "Token is not valid")
			return
		}

		ctxUser := context.WithValue(r.Context(), "user", claims.User)

		r = r.WithContext(ctxUser)

		next(w, r)
	}
}

func (m *AuthMiddleware) renderJSONError(w http.ResponseWriter, status int, err error, message string) {
	golog.Error(err)
	m.renderJSON(w, status, err.Error(), message)
}

func (m *AuthMiddleware) renderJSON(w http.ResponseWriter, status int, data interface{}, message string) {
	responseMapper := response.ResponseMapper{}
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
