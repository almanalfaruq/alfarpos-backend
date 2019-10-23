package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kataras/golog"

	"../model/response"
	"../service"
)

type UserController struct {
	service.IUserService
}

func (controller *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("POST - User: RegisterHandler (/users/register)")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	user, err := controller.NewUser(string(body))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
}

func (controller *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("POST - User: LoginHandler (/users/login)")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := controller.LoginUser(string(body))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	loginResponse := response.LoginResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(loginResponse)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
}
