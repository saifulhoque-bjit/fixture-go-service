package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	store, err := OpenStore(":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte("ok")) })
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	log.Printf("listening on %s", addr)
	log.Fatal(server.ListenAndServe())
}
