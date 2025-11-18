package apipoc

import (
	_ "embed"
	"log/slog"
	"net/http"
)

func Stub() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ops/ping", Ping)
	mux.HandleFunc("/dataset/", Dataset)

	slog.Info("starting apipoc")
	http.ListenAndServe("localhost:3000", mux)
}

func Ping(w http.ResponseWriter, req *http.Request) {
	slog.Info("apipoc ping")
	w.WriteHeader(http.StatusOK)
}

//go:embed dataset.json
var dataset []byte

func Dataset(w http.ResponseWriter, req *http.Request) {
	slog.Info("apipoc dataset")
	w.WriteHeader(http.StatusOK)
	w.Write(dataset)
}
