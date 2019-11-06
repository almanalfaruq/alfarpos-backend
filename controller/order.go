package controller

import (
	"encoding/json"
	"net/http"

	"../model"
	"../model/response"
	"../service"
	"../util"
	"github.com/kataras/golog"
)

type OrderController struct {
	service.IOrderService
	util.Config
}

func (controller *OrderController) GetAllOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var orders []model.Order
	var err error
	var responseMapper response.ResponseMapper

	authHeader := r.Header.Get("Authorization")
	user, err := ParseJwtToUser(authHeader, controller.SecretKey)
	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusBadRequest,
			Data:    err.Error(),
			Message: "Cannot parse token",
		}
		err := json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}
	if user.RoleID == model.Cashier {
		golog.Error("User must be Admin or Manager")
		w.WriteHeader(http.StatusForbidden)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusForbidden,
			Data:    "User must be Admin or Manager",
			Message: "User must be Admin or Manager",
		}
		err := json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	orders, err = controller.GetAllOrder()
	if err != nil {
		golog.Error(err)
		responseMapper = response.ResponseMapper{
			Code:    http.StatusInternalServerError,
			Data:    err.Error(),
			Message: "Cannot get all categories",
		}
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error("Cannot encode json")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	responseMapper = response.ResponseMapper{
		Code:    http.StatusOK,
		Data:    orders,
		Message: "Success getting all orders",
	}
	err = json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
