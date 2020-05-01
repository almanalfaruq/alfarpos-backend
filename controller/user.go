package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kataras/golog"

	"github.com/almanalfaruq/alfarpos-backend/model/response"
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
		golog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusInternalServerError,
			Data:    err.Error(),
			Message: "Cannot read request body",
		}
		err = json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	user, err := c.user.NewUser(string(body))
	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusUnprocessableEntity,
			Data:    err.Error(),
			Message: "Cannot create user",
		}
		err = json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	responseMapper := response.ResponseMapper{
		Code:    http.StatusCreated,
		Data:    user,
		Message: "User created!",
	}
	err = json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
}

func (c *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - User: LoginHandler (/users/login)")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusInternalServerError,
			Data:    err.Error(),
			Message: "Cannot read request body",
		}
		err = json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
			return
		}
		return
	}

	token, err := c.user.LoginUser(string(body))
	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusUnauthorized)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusUnauthorized,
			Data:    err.Error(),
			Message: "Cannot login user",
		}
		err = json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	loginResponse := response.ResponseMapper{
		Code:    http.StatusOK,
		Data:    token,
		Message: "User logged in!",
	}

	err = json.NewEncoder(w).Encode(loginResponse)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
}
