package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/jessjenkins/dp-stubby/stubs/apipoc"
	"github.com/jessjenkins/dp-stubby/stubs/dataset"
	"github.com/jessjenkins/dp-stubby/stubs/images"
	"github.com/jessjenkins/dp-stubby/stubs/zebedee"
)

func main() {
	slog.Info("starting")
	signals := make(chan os.Signal, 1)

	go zebedee.Stub()
	go images.Stub()
	go apipoc.Stub()
	go dataset.Stub()

	select {
	case sig := <-signals:
		log.Printf("os signal received:%v", sig)
	}
}
