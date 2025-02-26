package types

import "github.com/golang-jwt/jwt/v5"

type AuthRequest struct {
	Username string
	Password string
}

type AuthResponse struct {
	Status string
	Msg    string `json:"msg,omitempty"`
	Token  string `json:"token,omitempty"`
}

type AuthClaims struct {
	Id uint
	jwt.RegisteredClaims
}
