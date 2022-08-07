package web

type TaskCreateResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message,omitempty"`
	Data    TaskUpdateRequest `json:"data,omitempty"`
}
