package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

var port = "8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d, err := httputil.DumpRequest(r, true)
		if err != nil {
			msg := fmt.Sprintf("couldn't dump request: %s", err)
			log.Printf(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		lg := fmt.Sprintf("%q", d)
		lg = strings.TrimPrefix(lg, "\"")
		lg = strings.TrimSuffix(lg, "\"")
		fmt.Println(lg)

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		if _, err := fmt.Fprintf(w, ""); err != nil {
			msg := fmt.Sprintf("couldn't write response: %s", err)
			log.Printf(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
	})

	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	addr := ":" + port

	log.Printf("http-dump is listening at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
