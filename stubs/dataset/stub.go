package dataset

import (
	_ "embed"
	"log/slog"
	"net/http"
)

func Stub() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/datasets/{id...}", Datasets)
	mux.HandleFunc("/", NotFound)

	slog.Info("starting datasets")
	http.ListenAndServe("localhost:22000", mux)
}

//go:embed health.json
var health []byte

func Health(w http.ResponseWriter, req *http.Request) {
	slog.Info("dataset health")
	w.WriteHeader(http.StatusOK)
	w.Write(health)
}

func NotFound(w http.ResponseWriter, req *http.Request) {
	slog.Info("dataset request not found", slog.String("uri", req.RequestURI))
	http.NotFound(w, req)
}

//go:embed valid.json
var valid []byte
var notFound []byte = []byte("dataset not found")

func Datasets(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	slog.Info("datasets request data", slog.String("id", id))

	if id == "wibble-wobble-dibble-dobble" {
		w.WriteHeader(http.StatusNotFound)
		w.Write(notFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(valid)
}
