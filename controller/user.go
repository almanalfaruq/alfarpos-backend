package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/kataras/golog"
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
// @Param json body model.User true "These field must be present: username, password, fullname, address, phone, and role_id (1 = Admin; 2 = Manager; 3 = Cashier)"
// @Success 200 {object} response.ResponseMapper{data=model.User} "Return the new registered user"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /users/register [post]
func (c *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - User: RegisterHandler (/users/register)")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, "Cannot read request body")
		return
	}

	user, err := c.user.NewUser(string(body))
	if err != nil {
		renderJSONError(w, http.StatusUnprocessableEntity, err, "Cannot create user")
		return
	}

	renderJSONSuccess(w, http.StatusCreated, user, "User created")
}

// LoginUser godoc
// @Summary Login user
// @Description Endpoint for getting the token for the logged in user
// @Tags user
// @Produce json
// @Param json body model.User true "These field must be present: username, password"
// @Success 200 {object} response.ResponseMapper{data=string} "Return a jwt token to be used for other requests"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /users/login [post]
func (c *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - User: LoginHandler (/users/login)")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, "Cannot read request body")
		return
	}

	token, err := c.user.LoginUser(string(body))
	if err != nil {
		renderJSONError(w, http.StatusUnauthorized, err, "Cannot login user")
		return
	}

	renderJSONSuccess(w, http.StatusOK, token, "User logged in")
}
