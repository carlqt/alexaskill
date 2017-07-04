// A typical incoming alexa request looks like this.
// {
//   "session": {
//     "sessionId": "SessionId.ec3aa84e-c489-4f0a-ba50-8b67d9673452",
//     "application": {
//       "applicationId": "amzn1.ask.skill.3aebac54-38a0-4dd3-9f17-4942972e4136"
//     },
//     "attributes": {},
//     "user": {
//       "userId": "amzn1.ask.account.AFWSUCAS2MHVONEYZ6YDMODQ4FDXFJTHOATXZZ2XODC7L5SECMNIXBDUIP3WKGDPCWE26LXCPOB6EII3H46CSPKJGLOIB3CLZTTNNJHFGFQEG73YBO65UPFNLXVKJXZ3RTT4EOMSU2IJWOQ6JDQYNNYYWAZX6E46ERLJKCTGGZQEFBU6XM5ZBS7YL2G2PEEVI6VOJ4P5WAYZJGY"
//     },
//     "new": false
//   },
//   "request": {
//     "type": "IntentRequest",
//     "requestId": "EdwRequestId.c30686c6-3e00-4da6-b49b-e1a49e2ade27",
//     "locale": "en-US",
//     "timestamp": "2017-07-02T02:10:24Z",
//     "intent": {
//       "name": "AMAZON.CancelIntent",
//       "slots": {}
//     }
//   },
//   "version": "1.0"
// }

package alexaskill

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type sessionAttr map[string]string

// AlexaRequest is the struct of the incoming alexa request
type AlexaRequest struct {
	Request struct {
		Type      string `json:"type"`
		RequestID string `json:"requestId"`
		Intent    Intent `json:"intent"`
	} `json:"request"`

	Session Session `json:"session"`
}

// Session object of alexa request
type Session struct {
	New         bool
	SessionID   string `json:"sessionId"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`

	Attributes sessionAttr `json:"attributes"`
}

// Intent object of alexa request
type Intent struct {
	Name  string `json:"name"`
	Slots Slots  `json:"slots"`
}

type Slots struct {
	Answer struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"answer"`
}

// AlexaNewRequest is a constructor that reads the request.Body from Amazon
func AlexaNewRequest(body io.ReadCloser) (*AlexaRequest, error) {
	rBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	alexa := new(AlexaRequest)
	err = json.Unmarshal(rBody, alexa)
	if err != nil {
		return nil, err
	}

	return alexa, nil
}

//Type is a function shortcut to get AlexaRequest.Request.Type
func (a *AlexaRequest) Type() string {
	return a.Request.Type
}

//IntentName is a function shortcut to get AlexaRequest.Request.Intent.Name
func (a *AlexaRequest) IntentName() string {
	return a.Request.Intent.Name
}

//AppID is a function shortcut to get AlexaRequest.Request.Intent.Name
func (a *AlexaRequest) AppID() string {
	return a.Session.Application.ApplicationID
}

func (a *AlexaRequest) GetUserAnswer() string {
	return a.Request.Intent.Slots.Answer.Value
}

func (a *AlexaRequest) GetSessionAttr(key string) string {
	return a.Session.Attributes[key]
}

func (s *Session) Get(key string) string {
	return s.Attributes[key]
}
