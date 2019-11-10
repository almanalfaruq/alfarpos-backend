package controller

import (
	"errors"
	"fmt"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/response"
	jwt "github.com/dgrijalva/jwt-go"
)

func ParseJwtToUser(authHeader string, secretKey string) (model.User, error) {
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

func ParseToResponseMapper(code int, data interface{}, message string) response.ResponseMapper {
	return response.ResponseMapper{
		Code:    code,
		Data:    data,
		Message: message,
	}
}
