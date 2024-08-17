package types





type AuthResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Token string      `json:"token,omitempty"`
	Error string      `json:"error,omitempty"`
	StatusCode int `json:"status_code,omitempty"`
}