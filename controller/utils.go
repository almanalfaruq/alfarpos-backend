package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/response"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
)

func renderJSONError(w http.ResponseWriter, status int, err error, message string) {
	golog.Error(err)
	renderJSON(w, status, err.Error(), message)
}

func renderJSONSuccess(w http.ResponseWriter, status int, data interface{}, message string) {
	renderJSON(w, status, data, message)
}

func renderJSON(w http.ResponseWriter, status int, data interface{}, message string) {
	responseMapper := parseToResponseMapper(status, data, message)
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func checkUser(user model.User, roles ...model.Role) bool {
	for _, role := range roles {
		if user.RoleID == role {
			return true
		}
	}
	return false
}

func parseJwtToUser(authHeader string, secretKey string) (model.User, error) {
	if !strings.Contains(authHeader, "Bearer") {
		return model.User{}, errors.New("Authorization header is empty")
	}
	authToken := strings.Replace(authHeader, "Bearer ", "", -1)
	token, err := jwt.ParseWithClaims(authToken, &response.TokenResponse{}, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return model.User{}, err
	}
	claims, ok := token.Claims.(*response.TokenResponse)
	if !ok || !token.Valid {
		return model.User{}, errors.New("Token cannot be processed")
	}
	return claims.User, nil
}

func parseToResponseMapper(code int, data interface{}, message string) response.ResponseMapper {
	return response.ResponseMapper{
		Code:    code,
		Data:    data,
		Message: message,
	}
}
