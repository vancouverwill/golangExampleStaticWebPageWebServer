package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func Handler(response http.ResponseWriter, request *http.Request) {
	log.Print("index.html handler")
	response.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(response, fmt.Sprintf("index.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
}

func AboutUsHandler(response http.ResponseWriter, request *http.Request) {
	log.Print("aboutus.html handler")
	response.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("aboutus.html")
	if err != nil {
		http.Error(response, fmt.Sprintf("aboutus.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
}

func SimpleServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	fmt.Fprint(response, "<h1>android</h1>Hello! what is your name")
	fmt.Fprint(response, "what is your name!<br/>")
	fmt.Fprint(response, "send me a message!")
}

func main() {
	port := 4000

	var err string
	portstring := strconv.Itoa(port)

	mux := http.NewServeMux()
	mux.Handle("/aboutus", http.HandlerFunc(AboutUsHandler))
	mux.Handle("/contact", http.HandlerFunc(SimpleServeHTTP))
	mux.Handle("/", http.HandlerFunc(Handler))

	log.Print("Listening on port " + portstring + " ... ")
	errs := http.ListenAndServe(":"+portstring, mux)
	if errs != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
