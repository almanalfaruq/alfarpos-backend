package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
	"github.com/golang-jwt/jwt/v4"
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}

		authHeader := r.Header.Get("Authorization")

		authHeaderSplit := strings.Split(authHeader, " ")
		if len(authHeaderSplit) != 2 {
			response.RenderJSONError(w, http.StatusUnauthorized, errors.New("Missing header Authorization"))
			return
		}

		if strings.ToLower(authHeaderSplit[0]) != "bearer" {
			response.RenderJSONError(w, http.StatusUnauthorized, errors.New("Missing header Authorization"))
			return
		}

		token, err := jwt.ParseWithClaims(authHeaderSplit[1], &response.TokenData{}, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return []byte(m.secretKey), nil
		})
		if err != nil {
			response.RenderJSONError(w, http.StatusUnauthorized, err)
			return
		}
		claims, ok := token.Claims.(*response.TokenData)
		if !ok || !token.Valid {
			response.RenderJSONError(w, http.StatusUnauthorized, errors.New("Token is not valid"))
			return
		}

		ctxUser := context.WithValue(r.Context(), userentity.CTX_USER, claims.User)

		r = r.WithContext(ctxUser)

		next(w, r)
	}
}

func (m *AuthMiddleware) CheckCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}

		next(w, r)
	}
}
