package alexaskill

type AlexaRequest struct {
	Request struct {
		Type      string `json:"type"`
		RequestID string `json:"requestId"`
		Intent    Intent `intent`
	} `json:"request"`
}

type Intent struct {
	Name string `json:"name"`
}
