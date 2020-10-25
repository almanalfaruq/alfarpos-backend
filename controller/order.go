package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
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

	user, ok := r.Context().Value(model.CTX_USER).(model.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(model.RoleManager, model.RoleAdmin); !ok {
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

// NewOrder godoc
// @Summary New order
// @Description New order is used for creating a new order by the order details per product.
// @Description Order only need to specify the primitive data like the product_id, user_id, etc. without having to specify its object.
// @Tags order
// @Accept json
// @Produce json
// @Param body body model.Order true "Order with order detail"
// @Success 200 {object} response.ResponseMapper{data=model.Order} "Return array of product"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /orders [post]
func (c *OrderController) NewOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user, ok := r.Context().Value(model.CTX_USER).(model.User)
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

	var data model.Order
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

	response.RenderJSONSuccess(w, http.StatusOK, order, "Success creating a new order")
}
