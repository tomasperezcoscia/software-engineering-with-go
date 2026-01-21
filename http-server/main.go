package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	const filepathRoot = "./http-server/"
	const port = "8080"

	apiCfg := apiConfig{fileserverHits: atomic.Int32{}}

	mux := http.NewServeMux()
	rootHandler := http.FileServer(http.Dir(filepathRoot))
	mux.Handle("/app/", http.StripPrefix("/app", apiCfg.middlewareMetricsInc(rootHandler)))
	mux.HandleFunc("GET /api/health", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)
	mux.HandleFunc("POST /api/validate_text", handlerTextValidator)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
