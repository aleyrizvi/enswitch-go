package enswitch

type Response struct {
	Responses []struct {
		Key     string `json:"key"`
		Code    uint32 `json:"code,string"`
		Message string `json:"message"`
	} `json:"responses"`
}
