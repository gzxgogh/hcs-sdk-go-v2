package model

type ItemNotFound struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type NeutronError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
