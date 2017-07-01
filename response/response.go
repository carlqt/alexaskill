package response

import (
	"encoding/json"
	"net/http"
)

type AlexaResponse struct {
	Version  string   `json:"version"`
	Response Response `json:"response"`
}

type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	Card             Card         `json:"card,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
	SSML string `json:"ssml"`
}

type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Text    string `json:"text"`
}

func AlexaText(speech string) *AlexaResponse {
	outputSpeech := OutputSpeech{
		Type: "PlainText",
		Text: speech,
		SSML: "",
	}

	return &AlexaResponse{
		Version:  "1.0",
		Response: Response{OutputSpeech: outputSpeech, ShouldEndSession: true},
	}
}

func (a *AlexaResponse) SimpleCard(title, content string) *AlexaResponse {
	a.Response.Card = Card{
		Type:    "Simple",
		Title:   title,
		Content: content,
	}

	return a
}

func (a *AlexaResponse) Respond(w http.ResponseWriter, status int) {
	resp, _ := json.Marshal(a)

	w.WriteHeader(status)
	w.Write(resp)
}
