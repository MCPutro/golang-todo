package web

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
