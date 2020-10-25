package response

import (
	"encoding/json"
	"net/http"

	"github.com/kataras/golog"
)

type responseStruct struct {
	Code    int         `json:"code" example:"200"`
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"Success getting all products"`
}

func RenderJSONError(w http.ResponseWriter, status int, err error) {
	renderJSON(w, status, nil, err.Error())
}

func RenderJSONSuccess(w http.ResponseWriter, status int, data interface{}, message string) {
	renderJSON(w, status, data, message)
}

func renderJSON(w http.ResponseWriter, status int, data interface{}, message string) {
	responseMapper := responseStruct{
		Code:    status,
		Data:    data,
		Message: message,
	}
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
