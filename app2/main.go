package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome to App2 Home"))
}

func main() {
	log.Println("==== app2 ====")
	http.HandleFunc("/home", homeHandler)
	http.ListenAndServe(":8002", nil)
}
