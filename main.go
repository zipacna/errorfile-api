package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func errorFilePresent() bool {
	_, err := os.Stat("errorfile")
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Content-Type", "text/plain")
	h.Set("Content-Length", "0")

	var code int

	if errorFilePresent() {
		code = http.StatusTeapot
	} else {
		code = http.StatusOK
	}

	w.WriteHeader(code)
}

func main() {
	port := flag.Uint("p", 9876, "port number")
	flag.Parse()

	var srv *server
	fmt.Printf("Listening on port %v\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), srv))
}
