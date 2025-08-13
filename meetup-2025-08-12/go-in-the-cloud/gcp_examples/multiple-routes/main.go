package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type echoRequest struct {
	Message string `json:"message"`
}

type echoResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		service := getenv("SERVICE_NAME", "multiple-routes")
		host := getenv("HOSTNAME", "unknown")
		fmt.Fprintf(w, "service=%s host=%s time=%s\n", service, host, time.Now().Format(time.RFC3339))
	})

	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "use POST", http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()

		var in echoRequest
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "bad json: "+err.Error(), http.StatusBadRequest)
			return
		}

		resp := echoResponse{
			Message:   in.Message,
			Timestamp: time.Now(),
			Service:   getenv("SERVICE_NAME", "multiple-routes"),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&resp)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Try /healthz, /info, or POST /echo")
	})

	port := getenv("PORT", "8080") // Cloud Run erwartet 8080
	log.Printf("listening on :%s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
