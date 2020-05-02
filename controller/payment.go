package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/util"
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
			renderJSONError(w, http.StatusInternalServerError, err, "Cannot get all payments")
			return
		}
	} else {
		golog.Infof("GET - Product: GetPaymentsByNameHandler (/payments?query=%s)", query)
		payments, err = c.payment.GetPaymentsByName(query)
		if err != nil {
			renderJSONError(w, http.StatusNotFound, err, "Cannot get payments by name")
			return
		}
	}

	renderJSONSuccess(w, http.StatusOK, payments, "Success getting payments")
}

func (c *PaymentController) GetPaymentByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	golog.Infof("GET - Product: GetPaymentByIdHandler (/payments/id/%d)", id)
	payment, err := c.payment.GetOnePayment(id)
	if err != nil {
		message := fmt.Sprintf("Cannot find payment with id: %v", id)
		renderJSONError(w, http.StatusNotFound, err, message)
		return
	}

	message := fmt.Sprintf("Success getting payment with id: %v", id)
	renderJSONSuccess(w, http.StatusOK, payment, message)
}

func (c *PaymentController) NewPaymentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - Payment: NewPaymentHandler (/payments)")

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)

	if err != nil {
		renderJSONError(w, http.StatusBadRequest, err, "Cannot parse token")
		return
	}

	if ok := user.HasRole(model.Manager, model.Admin); !ok {
		message := "User must be Admin or Manager"
		renderJSONError(w, http.StatusForbidden, fmt.Errorf(message), message)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, "Cannot read request body")
		return
	}

	payment, err := c.payment.NewPayment(string(body))
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusCreated, payment, "Payment created")
}

func (c *PaymentController) UpdatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("PUT - Payment: UpdatePaymentHandler (/payments/%v)", id)

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)

	if err != nil {
		renderJSONError(w, http.StatusBadRequest, err, "Cannot parse token")
		return
	}

	if ok := user.HasRole(model.Manager, model.Admin); !ok {
		message := "User must be Admin or Manager"
		renderJSONError(w, http.StatusForbidden, fmt.Errorf(message), message)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, "Cannot read request body")
		return
	}

	payment, err := c.payment.UpdatePayment(string(body))
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusOK, payment, "Payment updated")
}

func (c *PaymentController) DeletePaymentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("DELETE - Payment: DeletePaymentHandler (/payments/%d)", id)

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)

	if err != nil {
		renderJSONError(w, http.StatusBadRequest, err, "Cannot parse token")
		return
	}

	if ok := user.HasRole(model.Manager, model.Admin); !ok {
		message := "User must be Admin or Manager"
		renderJSONError(w, http.StatusForbidden, fmt.Errorf(message), message)
		return
	}

	payment, err := c.payment.DeletePayment(id)
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusOK, payment, "Payment deleted")
}
