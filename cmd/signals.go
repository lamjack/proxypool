package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// HandleSignals listens for OS signals and calls the provided stopFunction when a SIGINT or SIGTERM signal is received.
func HandleSignals(stopFunction func()) {
	var callback sync.Once

	// On ^C or SIGTERM, gracefully stop the sniffer
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigc
		log.Println("service", "Received sigterm/sigint, stopping")
		callback.Do(stopFunction)
	}()
}
