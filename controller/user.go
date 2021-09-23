package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/almanalfaruq/alfarpos-backend/util/logger"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
)

type UserController struct {
	user userServiceIface
}

func NewUserController(userService userServiceIface) *UserController {
	return &UserController{
		user: userService,
	}
}

// RegisterUser godoc
// @Summary Register new user
// @Description Register new user (TODO: API caller should be a Manager or an Admin)
// @Tags user
// @Produce json
// @Param json body user.User true "These field must be present: username, password, fullname, address, phone, and role_id (1 = Admin; 2 = Manager; 3 = Cashier)"
// @Success 200 {object} response.ResponseMapper{data=user.User} "Return the new registered user"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /users/register [post]
func (c *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	logger.Log.Info("POST - User: RegisterHandler (/users/register)")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Log.Debug(err)
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	user, err := c.user.NewUser(string(body))
	if err != nil {
		logger.Log.Debug(err)
		response.RenderJSONError(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusCreated, user, "User created")
}

// LoginUser godoc
// @Summary Login user
// @Description Endpoint for getting the token for the logged in user
// @Tags user
// @Produce json
// @Param json body user.User true "These field must be present: username, password"
// @Success 200 {object} response.ResponseMapper{data=string} "Return a jwt token to be used for other requests"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /users/login [post]
func (c *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	logger.Log.Info("POST - User: LoginHandler (/users/login)")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Log.Debug(err)
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	data, err := c.user.LoginUser(string(body))
	if err != nil {
		logger.Log.Debug(err)
		response.RenderJSONError(w, http.StatusUnauthorized, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, data, "User logged in")
}
