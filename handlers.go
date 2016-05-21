package main

import (
	"net/http"
	"os"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("hub.verify_token") == os.Getenv("VERIFY_TOKEN") {
		w.WriteHeader(200)
		w.Write([]byte(r.URL.Query().Get("hub.challenge")))
	} else {
		w.WriteHeader(500)
		w.Write([]byte("Error, wrong validation token"))
	}
}
