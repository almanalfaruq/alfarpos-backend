package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	transactionentity "github.com/almanalfaruq/alfarpos-backend/model/transaction"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util/logger"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
)

type MoneyController struct {
	moneyUC moneyUsecase
}

func NewMoney(moneyUC moneyUsecase) *MoneyController {
	return &MoneyController{
		moneyUC: moneyUC,
	}
}

// NewMoney godoc
// @Summary New money
// @Description New money is used for creating a new money transaction (in/out)
// @Tags money
// @Accept json
// @Produce json
// @Param body body transactionentity.Money true "Money transaction"
// @Success 200 {object} response.ResponseMapper{data=transactionentity.Money} "Return money transaction data"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /money [post]
func (c *MoneyController) NewMoneyHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	logger.Log.Infof("%s - Money: NewMoneyHandler (/money)", r.Method)

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

	var data transactionentity.Money
	err = json.Unmarshal(body, &data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	money, err := c.moneyUC.InsertMoney(data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, money, "Success creating a new money transaction")
}

// GetMoneyTransactionWithFilter godoc
// @Summary Get money transaction with filter
// @Description Get money transaction with filter status and date
// @Tags money
// @Accept json
// @Produce json
// @Param body body transactionentity.GetMoneyWithFilterReq true "Money transaction"
// @Success 200 {object} response.ResponseMapper{data=[]transactionentity.Money} "Return monies transaction data"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /money/filters [post]
func (c *MoneyController) GetMoneyTransactionWithFilterHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	logger.Log.Infof("%s - Money: GetMoneyTransactionWithFilterHandler (/money/filters)", r.Method)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleAdmin); !ok {
		message := "User must be Admin"
		response.RenderJSONError(w, http.StatusForbidden, fmt.Errorf(message))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	var param transactionentity.GetMoneyWithFilterReq
	err = json.Unmarshal(body, &param)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	monies, err := c.moneyUC.GetMoneyWithFilter(param)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, monies, "Success getting monies with filter")
}
