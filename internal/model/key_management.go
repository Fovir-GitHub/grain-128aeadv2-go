package model

type WrapKeyRequest struct {
	Key      string `json:"key"`
	Password string `json:"password"`
	AD       string `json:"ad"`
	IsHex    bool   `json:"isHex"`
}

type UnwrapKeyRequest struct {
	Base64Content string `json:"b64"`
	Passphrase    string `json:"passphrase"`
	AD            string `json:"ad"`
	IsHex         bool   `json:"isHex"`
}

type WrapKeyResp struct {
	Key string `json:"key"`
}

type UnwrapKeyResp struct {
	Key string `json:"key"`
}
