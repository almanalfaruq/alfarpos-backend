package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/response"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/kataras/golog"
)

type OrderController struct {
	order orderServiceIface
	conf  util.Config
}

func NewOrderController(conf util.Config, orderService orderServiceIface) *OrderController {
	return &OrderController{
		conf:  conf,
		order: orderService,
	}
}

func (c *OrderController) GetAllOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var orders []model.Order
	var err error
	var responseMapper response.ResponseMapper

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)
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

	orders, err = c.order.GetAllOrder()
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

func (c *OrderController) NewOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var order model.Order
	var err error
	var responseMapper response.ResponseMapper

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)
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
	if user.ID == 0 {
		golog.Error("User must logged in!")
		w.WriteHeader(http.StatusForbidden)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusForbidden,
			Data:    "User must logged in!",
			Message: "User must logged in!",
		}
		err := json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusInternalServerError,
			Data:    err.Error(),
			Message: "Cannot read request body",
		}
		err = json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	order, err = c.order.NewOrder(string(body))
	if err != nil {
		golog.Error(err)
		responseMapper = response.ResponseMapper{
			Code:    http.StatusInternalServerError,
			Data:    err.Error(),
			Message: "Cannot create a new order",
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
		Data:    order,
		Message: "Success creating a new order",
	}
	err = json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
