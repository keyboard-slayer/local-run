package types

type PingSchema struct {
	Response string `json:"response"`
}

func NewPing() PingSchema {
	return PingSchema{Response: "ping"}
}
