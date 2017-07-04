package response

import (
	"encoding/json"
	"net/http"
)

type sessionAttr map[string]string

type AlexaResponse struct {
	Version           string      `json:"version"`
	SessionAttributes sessionAttr `json:"sessionAttributes"`
	Response          Response    `json:"response"`
}

type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	Card             Card         `json:"card,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession"`
	Reprompt         struct {
		OutputSpeech OutputSpeech `json:"outputSpeech"`
	} `json:"reprompt"`
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
	session := make(map[string]string)

	return &AlexaResponse{
		Version:           "1.0",
		Response:          Response{OutputSpeech: outputSpeech, ShouldEndSession: false},
		SessionAttributes: session,
	}
}

func (a *AlexaResponse) SessionAttr(key, value string) *AlexaResponse {
	a.SessionAttributes[key] = value

	return a
}

func (a *AlexaResponse) SimpleCard(title, content string) *AlexaResponse {
	a.Response.Card = Card{
		Type:    "Simple",
		Title:   title,
		Content: content,
	}

	return a
}

func (a *AlexaResponse) Respond(w http.ResponseWriter, status int, closeSession bool) {
	a.Response.ShouldEndSession = closeSession
	resp, _ := json.Marshal(a)

	w.WriteHeader(status)
	w.Write(resp)
}

func (a *AlexaResponse) RepromptText(speech string) *AlexaResponse {
	outputSpeech := OutputSpeech{
		Type: "PlainText",
		Text: speech,
		SSML: "",
	}

	a.Response.Reprompt.OutputSpeech = outputSpeech
	return a
}
