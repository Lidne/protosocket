package main

import (
	"main/internal/application/handlers"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logger.SetLevel(logrus.DebugLevel)

	wsHanler := handlers.NewWebSocketHandler(logger)
	http.HandleFunc("/ws", wsHanler.HandleStream)
	logger.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		logger.Fatal("ListenAndServe: ", err)
	}
}
