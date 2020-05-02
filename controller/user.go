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
