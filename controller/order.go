package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/almanalfaruq/alfarpos-backend/util/logger"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
	"github.com/gorilla/mux"
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
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	orders, err := c.order.GetAllOrder()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, orders, "Success getting all orders")
}

func (c *OrderController) GetOrderUsingFilterHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	logger.Log.Infof("%s - Order: GetOrderUsingFilterHandler (/orders/filter)", r.Method)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID < 1 {
		err := errors.New("User not found")
		response.RenderJSONError(w, http.StatusForbidden, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	var param orderentity.GetOrderUsingFilterParam
	err = json.Unmarshal(body, &param)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if param.Sort != "" && (param.Sort != orderentity.SortAsc && param.Sort != orderentity.SortDesc) {
		err = errors.New("Unsupported sort param")
		response.RenderJSONError(w, http.StatusBadRequest, err)
		return
	}

	orders, err := c.order.GetOrderUsingFilter(param)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, orders, "Success getting orders with filter")
}

// NewOrder godoc
// @Summary New order
// @Description New order is used for creating a new order by the order details per product.
// @Description Order only need to specify the primitive data like the product_id, user_id, etc. without having to specify its object.
// @Tags order
// @Accept json
// @Produce json
// @Param body body order.Order true "Order with order detail"
// @Success 200 {object} response.ResponseMapper{data=order.Order} "Return order data"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /orders [post]
func (c *OrderController) NewOrderHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	logger.Log.Infof("%s - Order: NewOrderHandler (/orders)", r.Method)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID == 0 {
		err := fmt.Errorf("User must logged in!")
		response.RenderJSONError(w, http.StatusForbidden, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	var data orderentity.Order
	err = json.Unmarshal(body, &data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	order, err := c.order.NewOrder(data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if order.Status == orderentity.StatusPending {
		response.RenderJSONSuccess(w, http.StatusOK, order, "Success hold an order")
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, order, "Success creating a new order")
}

func (c *OrderController) GetOrderByIDHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	logger.Log.Infof("%s - Order: GetOrderByIDHandler (/orders/%d)", r.Method, id)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		err := errors.New("User must be Admin or Manager")
		response.RenderJSONError(w, http.StatusForbidden, err)
		return
	}

	orders, err := c.order.GetOneOrder(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, orders, fmt.Sprintf("Success getting order by id: %d", id))
}

func (c *OrderController) UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	logger.Log.Infof("%s - Order: UpdateStatusHandler (/orders/status)", r.Method)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID < 1 {
		err := fmt.Errorf("User must logged in!")
		response.RenderJSONError(w, http.StatusForbidden, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	var data orderentity.Order
	err = json.Unmarshal(body, &data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if data.ID < 1 {
		response.RenderJSONError(w, http.StatusInternalServerError, errors.New("OrderID < 1"))
		return
	}

	order, err := c.order.UpdateOrderStatus(data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, order, "Success updating the order status")
}
