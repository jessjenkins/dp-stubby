package zebedee

import (
	_ "embed"
	"log/slog"
	"net/http"
)

func Stub() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/data/", Data)
	mux.HandleFunc("/data", Data)
	mux.HandleFunc("/", NotFound)

	//Page examples

	slog.Info("starting zebedee")
	http.ListenAndServe("localhost:8082", mux)
}

func Health(w http.ResponseWriter, req *http.Request) {
	slog.Info("zebedee health")
	w.WriteHeader(http.StatusOK)
}

func NotFound(w http.ResponseWriter, req *http.Request) {
	slog.Info("zebedee request not found", slog.String("uri", req.RequestURI))
	http.NotFound(w, req)
}

//

var zebedeeContent string = "/Users/jessjenkins/dev/ons/dp/zebedee-content/zebedee/master"

//go:embed content/dataset_example.json
var dsExample []byte

func Data(w http.ResponseWriter, req *http.Request) {
	uri := req.FormValue("uri")
	slog.Info("zebedee request data", slog.String("uri", uri))

	var data []byte
	switch uri {
	case "/peoplepopulationandcommunity/householdcharacteristics/homeinternetandsocialmediausage/datasets/internetusers":
		data = dsExample
	}

	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
