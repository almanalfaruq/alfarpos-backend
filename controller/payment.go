package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/response"
	"github.com/almanalfaruq/alfarpos-backend/service"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
)

type PaymentController struct {
	service.IPaymentService
	util.Config
}

func (controller *PaymentController) GetPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var payments []model.Payment
	var err error
	var responseMapper response.ResponseMapper
	query := r.URL.Query().Get("query")

	if query == "" {
		golog.Info("GET - Payment: GetAllPaymentHandler (/payments)")
		payments, err = controller.GetAllPayment()
		if err != nil {
			golog.Error(err)
			responseMapper = response.ResponseMapper{
				Code:    http.StatusInternalServerError,
				Data:    err.Error(),
				Message: "Cannot get all payments",
			}
			w.WriteHeader(http.StatusNotFound)
			err = json.NewEncoder(w).Encode(responseMapper)
			if err != nil {
				golog.Error("Cannot encode json")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	} else {
		golog.Infof("GET - Product: GetPaymentsByNameHandler (/payments?query=%s)", query)
		payments, err = controller.GetPaymentsByName(query)
		if err != nil {
			golog.Error(err)
			responseMapper = response.ResponseMapper{
				Code:    http.StatusNotFound,
				Data:    err.Error(),
				Message: "Cannot get payments by name",
			}
			w.WriteHeader(http.StatusNotFound)
			err = json.NewEncoder(w).Encode(responseMapper)
			if err != nil {
				golog.Error("Cannot encode json")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	responseMapper = response.ResponseMapper{
		Code:    http.StatusOK,
		Data:    payments,
		Message: "Success getting payments",
	}
	err = json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *PaymentController) GetPaymentByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var responseMapper response.ResponseMapper
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("GET - Product: GetPaymentByIdHandler (/payments/id/%v)", id)
	payment, err := controller.GetOnePayment(int(id))
	if err != nil {
		golog.Error(err)
		responseMapper = response.ResponseMapper{
			Code:    http.StatusNotFound,
			Data:    err.Error(),
			Message: fmt.Sprintf("Cannot find payment with id: %v", id),
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
		Data:    payment,
		Message: fmt.Sprintf("Success getting payment with id: %v", id),
	}
	err = json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *PaymentController) NewPaymentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - Payment: NewPaymentHandler (/payments)")

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

	payment, err := controller.NewPayment(string(body))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(payment)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *PaymentController) UpdatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("PUT - Payment: UpdatePaymentHandler (/payments/%v)", id)

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

	payment, err := controller.UpdatePayment(string(body))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(payment)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *PaymentController) DeletePaymentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("DELETE - Payment: DeletePaymentHandler (/payments/%v)", id)

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

	payment, err := controller.DeletePayment(int(id))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(payment)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
