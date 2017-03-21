package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var response string
var filename string
var listen string

func homeHandler(w http.ResponseWriter, r *http.Request) {
	errorHandler(w, r, http.StatusNotFound)
}

func initBlackhole() {
	filename = os.Getenv("FILENAME")
	if len(filename) == 0 {
		filename = "/var/www/index.html"
	}
	listen = os.Getenv("LISTEN")
	if len(listen) == 0 {
		listen = ":8080"
	}
	output, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	response = string(output)
}

func main() {
	initBlackhole()
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(listen, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(http.StatusUnavailableForLegalReasons)
	fmt.Fprint(w, response)
}
