package models

// ResponsePayload represents the response structure
type ResponsePayload struct {
	EventName string   `json:"event_name"`
	Message   string   `json:"message"`
	URLs      []string `json:"urls"`
	Status    string   `json:"status"`
	Username  string   `json:"username"`
}
