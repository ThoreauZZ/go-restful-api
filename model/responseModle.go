package model

type ResponseModel struct {
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}
