package model

type ResponseFormmater struct {
	Error  *string     `json:"error"`
	Result interface{} `json:"result"`
}
