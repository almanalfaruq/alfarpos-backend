package controller

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/almanalfaruq/alfarpos-backend/model"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
)

type PaymentController struct {
	payment paymentServiceIface
	conf    util.Config
}

func NewPaymentController(conf util.Config, paymentService paymentServiceIface) *PaymentController {
	return &PaymentController{
		payment: paymentService,
		conf:    conf,
	}
}

func (c *PaymentController) GetPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var payments []model.Payment
	var err error
	query := r.URL.Query().Get("query")

	if query == "" {
		golog.Info("GET - Payment: GetAllPaymentHandler (/payments)")
		payments, err = c.payment.GetAllPayment()
		if err != nil {
			response.RenderJSONError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		golog.Infof("GET - Product: GetPaymentsByNameHandler (/payments?query=%s)", query)
		payments, err = c.payment.GetPaymentsByName(query)
		if err != nil {
			response.RenderJSONError(w, http.StatusNotFound, err)
			return
		}
	}

	response.RenderJSONSuccess(w, http.StatusOK, payments, "Success getting payments")
}

func (c *PaymentController) GetPaymentByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	golog.Infof("GET - Product: GetPaymentByIdHandler (/payments/id/%d)", id)
	payment, err := c.payment.GetOnePayment(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusNotFound, err)
		return
	}

	message := fmt.Sprintf("Success getting payment with id: %v", id)
	response.RenderJSONSuccess(w, http.StatusOK, payment, message)
}

func (c *PaymentController) NewPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - Payment: NewPaymentHandler (/payments)")

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, fmt.Errorf(message))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	payment, err := c.payment.NewPayment(string(body))
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusCreated, payment, "Payment created")
}

func (c *PaymentController) UpdatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("PUT - Payment: UpdatePaymentHandler (/payments/%v)", id)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, fmt.Errorf(message))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	payment, err := c.payment.UpdatePayment(string(body))
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, payment, "Payment updated")
}

func (c *PaymentController) DeletePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("DELETE - Payment: DeletePaymentHandler (/payments/%d)", id)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, fmt.Errorf(message))
		return
	}

	payment, err := c.payment.DeletePayment(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, payment, "Payment deleted")
}
