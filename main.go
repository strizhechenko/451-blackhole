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
	w.WriteHeader(http.StatusUnavailableForLegalReasons)
	fmt.Fprint(w, response)
}

func init() {
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
	http.HandleFunc("/", homeHandler)
	fmt.Printf("Listen at %s...\r\n", listen)
	http.ListenAndServe(listen, nil)
}
