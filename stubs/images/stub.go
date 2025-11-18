package images

import (
	_ "embed"
	"log/slog"
	"net/http"
)

func Stub() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/images", Images)

	slog.Info("starting images")
	http.ListenAndServe("localhost:24700", mux)
}

//go:embed health.json
var health []byte

func Health(w http.ResponseWriter, req *http.Request) {
	slog.Info("image health")
	w.WriteHeader(http.StatusOK)
	w.Write(health)
}

//go:embed images.json
var images []byte

func Images(w http.ResponseWriter, req *http.Request) {
	slog.Info("images")
	w.WriteHeader(http.StatusOK)
	w.Write(images)
}
