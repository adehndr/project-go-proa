package web

type WebResponseRequest struct{
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    []TaskResponse `json:"data,omitempty"`
}