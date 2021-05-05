package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	profileentity "github.com/almanalfaruq/alfarpos-backend/model/profile"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
)

type ProfileController struct {
	profileService profileService
}

func NewProfile(profileService profileService) *ProfileController {
	return &ProfileController{
		profileService: profileService,
	}
}

func (c *ProfileController) GetProfileByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("%s - Profile: GetProfileByIDHandler (/profiles/%d)", r.Method, id)
	if err != nil {
		response.RenderJSONError(w, http.StatusBadRequest, err)
		return
	}

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

	profile, err := c.profileService.GetByID(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, profile, "Success getting profile by ID")
}

func (c *ProfileController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Infof("%s - Profile: UpdateHandler (/profiles)", r.Method)

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

	if user.HasRole(userentity.RoleAdmin, userentity.RoleManager) {
		err := fmt.Errorf("User must be Admin or Manager!")
		response.RenderJSONError(w, http.StatusForbidden, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	var data profileentity.Profile
	err = json.Unmarshal(body, &data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	profile, err := c.profileService.Update(data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, profile, "Success updating profile")
}
