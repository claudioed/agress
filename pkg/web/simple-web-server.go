package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.Info("Starting go Server...")
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		log.Info(fmt.Printf("Receiving Request on %s",request.URL.Path))
		userAgent := request.Header.Get("User-Agent")
		name := request.URL.Query().Get("name")
		log.WithFields(log.Fields{"user": name , "from" : userAgent}).Info("Calling hello")
		writer.Write([]byte(fmt.Sprintf("Hello %s ", name)))
	})
	http.ListenAndServe(":8080", nil)
}
