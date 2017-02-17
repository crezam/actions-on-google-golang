package main

import (
	"encoding/json"
	"fmt"
	"github.com/crezam/actions-on-google-golang/model"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", HandleWebhook)
	http.ListenAndServe("localhost:8000", nil)
}

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var webhookRequest model.ApiAiRequest
	err := decoder.Decode(&webhookRequest)
	if err != nil {
		panic(err)
	}
	w.Write([]byte("OK"))
}
