package stats

import (
	"errors"
	"fmt"
	"net/http"

	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
	"github.com/kataras/golog"
)

type StatsController struct {
	statsService statsService
}

func New(statsService statsService) *StatsController {
	return &StatsController{
		statsService: statsService,
	}
}

func (c *StatsController) GetShopStats(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	golog.Infof("%s - Stats: GetShopStats (/stats)", r.Method)

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
