package types

type RPCRequest struct {
	ID     string            `json:"id"`
	Method string            `json:"method"`
	Params map[string]string `json:"params,omitempty"`
}

type RPCResponse struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data string `json:"data,omitempty"`
}
