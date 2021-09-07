package models

type GetResponse struct {
	Quote string `json:"quote"`
	Error string `json:"error"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
