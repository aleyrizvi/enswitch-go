package enswitch

type Response struct {
	Responses []struct {
		Key     string `json:"key"`
		Code    uint   `json:"code,string"`
		Message string `json:"message"`
	} `json:"responses"`
}
