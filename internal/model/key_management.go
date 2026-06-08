package model

type WrapKeyRequest struct {
	Key      string `json:"key"`
	Password string `json:"password"`
	AD       string `json:"ad"`
	IsHex    bool   `json:"isHex"`
}
