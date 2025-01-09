package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request /%s", r.URL.Path[1:])
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handler_eric(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request /%s", r.URL.Path[1:])
	fmt.Fprintf(w, "Hola mon %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/eric", handler_eric)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
