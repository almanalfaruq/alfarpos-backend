package response

type ResponseMapper struct {
	Code    int         `json:"code" example:"200"`
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"Success getting all products"`
}
