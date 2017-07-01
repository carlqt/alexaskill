## Install
go get -u github.com/carlqt/alexaskill

## Middleware
The library comes with a middleware to handle amazons validation of Signature. Just wrap `AlexaValidation`  middleware to any of your handlers.

## Example
```go
package main

import (
	"log"
	"net/http"

	"github.com/carlqt/alexaskill"
	"github.com/carlqt/alexaskill/middleware"
	"github.com/carlqt/alexaskill/response"
	"github.com/gorilla/mux"
)

func main() {
	var appID := "amazon.app.id.kdjakjf"
	r := mux.NewRouter()
	alexaHandler := http.HandlerFunc(index)
	r.Handle("/", middleware.AlexaValidation(alexaHandler)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func index(w http.ResponseWriter, r *http.Request) {
	advice := "Carl says: Star this repository now"

	alexaReq, err := alexaskill.AlexaNewRequest(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if alexaReq.AppID() != appID {
		w.WriteHeader(404)
		return
	}

	if alexaReq.Type() == "IntentRequest" {
		switch alexaReq.IntentName() {
		case "AMAZON.CancelIntent":
			response.AlexaText("Cancelled").SimpleCard("Cancel", "cancel").Respond(w, 200)
		case "RandomAdvice":
			response.AlexaText(advice).SimpleCard("Advice", advice).Respond(w, 200)
		default:
			log.Fatal("unrecognized")
		}
	}
}


```
