package controller

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/almanalfaruq/alfarpos-backend/model"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/almanalfaruq/alfarpos-backend/util/logger"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
	"github.com/gorilla/mux"
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

// GetPayments godoc
// @Summary Get all payments or search by name
// @Description Get all payments or search by name
// @Tags payment
// @Produce json
// @Param query query string false "Payment name query"
// @Success 200 {object} response.ResponseStruct{data=[]model.Payment} "Return array of payments"
// @Failure 404 {object} response.ResponseStruct{data=string} "Not Found"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Router /payments [get]
func (c *PaymentController) GetPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var payments []model.Payment
	var err error
	query := r.URL.Query().Get("query")

	if query == "" {
		logger.Log.Info("GET - Payment: GetAllPaymentHandler (/payments)")
		payments, err = c.payment.GetAllPayment()
		if err != nil {
			logger.Log.Debug(err)
			response.RenderJSONError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		logger.Log.Infof("GET - Product: GetPaymentsByNameHandler (/payments?query=%s)", query)
		payments, err = c.payment.GetPaymentsByName(query)
		if err != nil {
			logger.Log.Debug(err)
			response.RenderJSONError(w, http.StatusNotFound, err)
			return
		}
	}

	response.RenderJSONSuccess(w, http.StatusOK, payments, "Success getting payments")
}

// GetPaymentByID godoc
// @Summary Get payment by ID
// @Description Get payment by ID
// @Tags payment
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} response.ResponseStruct{data=model.Payment} "Return payment"
// @Failure 404 {object} response.ResponseStruct{data=string} "Not Found"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Router /payments/id/{id} [get]
func (c *PaymentController) GetPaymentByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	logger.Log.Infof("GET - Product: GetPaymentByIdHandler (/payments/id/%d)", id)
	payment, err := c.payment.GetOnePayment(id)
	if err != nil {
		logger.Log.Debug(err)
		response.RenderJSONError(w, http.StatusNotFound, err)
		return
	}

	message := fmt.Sprintf("Success getting payment with id: %v", id)
	response.RenderJSONSuccess(w, http.StatusOK, payment, message)
}

// NewPayment godoc
// @Summary Create a new payment
// @Description Create a new payment
// @Tags payment
// @Accept json
// @Produce json
// @Param payment body model.Payment true "Payment data"
// @Success 201 {object} response.ResponseStruct{data=model.Payment} "Payment created"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /payments [post]
func (c *PaymentController) NewPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	logger.Log.Info("POST - Payment: NewPaymentHandler (/payments)")

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, errors.New(message))
		return
	}

	body, err := io.ReadAll(r.Body)
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

// UpdatePayment godoc
// @Summary Update a payment
// @Description Update a payment
// @Tags payment
// @Accept json
// @Produce json
// @Param id path int true "Payment ID"
// @Param payment body model.Payment true "Payment data"
// @Success 200 {object} response.ResponseStruct{data=model.Payment} "Payment updated"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /payments/{id} [put]
func (c *PaymentController) UpdatePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	logger.Log.Infof("PUT - Payment: UpdatePaymentHandler (/payments/%v)", id)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, errors.New(message))
		return
	}

	body, err := io.ReadAll(r.Body)
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

// DeletePayment godoc
// @Summary Delete a payment
// @Description Delete a payment by ID
// @Tags payment
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} response.ResponseStruct{data=model.Payment} "Payment deleted"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /payments/{id} [delete]
func (c *PaymentController) DeletePaymentHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	logger.Log.Infof("DELETE - Payment: DeletePaymentHandler (/payments/%d)", id)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, errors.New(message))
		return
	}

	payment, err := c.payment.DeletePayment(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, payment, "Payment deleted")
}
