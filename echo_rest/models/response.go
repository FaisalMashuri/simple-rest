package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseLogin struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
