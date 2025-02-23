package types

type AuthRequest struct {
	Username string
	Password string
}

type AuthResponse struct {
	Status string
	Msg    string `json:"msg,omitempty"`
}
