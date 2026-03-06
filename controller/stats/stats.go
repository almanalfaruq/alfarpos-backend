package stats

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	statsentity "github.com/almanalfaruq/alfarpos-backend/model/stats"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util/logger"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
)

var _ = statsentity.ShopStats{}

type StatsController struct {
	statsService statsService
}

func New(statsService statsService) *StatsController {
	return &StatsController{
		statsService: statsService,
	}
}

// GetShopStats godoc
// @Summary Get shop stats by date
// @Description Get aggregated shop stats (sales, expenses, profit) for a specific date
// @Tags stats
// @Produce json
// @Param date query string true "Date in YYYY-MM-DD format"
// @Success 200 {object} response.ResponseStruct{data=statsentity.ShopStats} "Return shop stats"
// @Failure 400 {object} response.ResponseStruct{data=string} "Bad Request"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /stats [get]
func (c *StatsController) GetShopStats(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	logger.Log.Infof("%s - Stats: GetShopStats (/stats)", r.Method)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if !user.HasRole(userentity.RoleAdmin) {
		err := errors.New("User must be an admin!")
		response.RenderJSONError(w, http.StatusForbidden, err)
		return
	}

	urlQueries := r.URL.Query()
	date := urlQueries.Get("date")
	if date == "" {
		date = r.FormValue("date")
		if date == "" {
			err := errors.New("date cannot be empty")
			response.RenderJSONError(w, http.StatusBadRequest, err)
			return
		}
	}

	result, err := c.statsService.GetShopStats(r.Context(), date)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, result, fmt.Sprintf("Success getting shop stats for date: %s", date))
}

// GetShopStatsByDateRangeV2 godoc
// @Summary Get shop stats by date range (V2)
// @Description Get daily aggregated shop stats for a specific date range with V2 response structure
// @Tags stats
// @Produce json
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Success 200 {object} response.ResponseStruct{data=[]statsentity.ShopStatsV2} "Return daily stats V2"
// @Failure 400 {object} response.ResponseStruct{data=string} "Bad Request"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /v2/stats/sales [get]
func (c *StatsController) GetShopStatsByDateRangeV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		response.RenderJSONError(w, http.StatusInternalServerError, errors.New("Cannot parse user context"))
		return
	}

	if !user.HasRole(userentity.RoleAdmin) {
		response.RenderJSONError(w, http.StatusForbidden, errors.New("User must be an admin!"))
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" || endDate == "" {
		response.RenderJSONError(w, http.StatusBadRequest, errors.New("start_date and end_date are required"))
		return
	}

	result, err := c.statsService.GetShopStatsByDateRange(r.Context(), startDate, endDate)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	var resultV2 []statsentity.ShopStatsV2
	for _, s := range result {
		resultV2 = append(resultV2, s.ToV2())
	}

	response.RenderJSONSuccess(w, http.StatusOK, resultV2, "Success getting shop stats by date range V2")
}

// GetTrendingProductsV2 godoc
// @Summary Get trending products (V2)
// @Description Get list of trending products based on sales volume for a specific period with V2 response structure
// @Tags stats
// @Produce json
// @Param period query string false "Period (week, month, year). Default: month"
// @Success 200 {object} response.ResponseStruct{data=[]statsentity.ProductStatsV2} "Return trending products V2"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /v2/stats/products/trending [get]
func (c *StatsController) GetTrendingProductsV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		response.RenderJSONError(w, http.StatusInternalServerError, errors.New("Cannot parse user context"))
		return
	}

	if !user.HasRole(userentity.RoleAdmin) {
		response.RenderJSONError(w, http.StatusForbidden, errors.New("User must be an admin!"))
		return
	}

	period := r.URL.Query().Get("period")
	if period == "" {
		period = "month"
	}

	result, err := c.statsService.GetTrendingProducts(r.Context(), period)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	var resultV2 []statsentity.ProductStatsV2
	for _, p := range result {
		resultV2 = append(resultV2, p.ToV2())
	}

	response.RenderJSONSuccess(w, http.StatusOK, resultV2, "Success getting trending products V2")
}

// GetProductSalesV2 godoc
// @Summary Get product sales stats (V2)
// @Description Get sales stats for a specific product over a period with V2 response structure
// @Tags stats
// @Produce json
// @Param id path int true "Product ID"
// @Param period query string false "Period (week, month, year). Default: month"
// @Success 200 {object} response.ResponseStruct{data=statsentity.ProductStatsV2} "Return product sales stats V2"
// @Failure 400 {object} response.ResponseStruct{data=string} "Invalid ID"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /v2/stats/products/{id}/sales [get]
func (c *StatsController) GetProductSalesV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		response.RenderJSONError(w, http.StatusInternalServerError, errors.New("Cannot parse user context"))
		return
	}

	if !user.HasRole(userentity.RoleAdmin) {
		response.RenderJSONError(w, http.StatusForbidden, errors.New("User must be an admin!"))
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.RenderJSONError(w, http.StatusBadRequest, errors.New("Invalid product ID"))
		return
	}

	period := r.URL.Query().Get("period")
	if period == "" {
		period = "month"
	}

	result, err := c.statsService.GetProductSales(r.Context(), id, period)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, result.ToV2(), "Success getting product sales V2")
}
