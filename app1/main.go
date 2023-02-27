package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "app1 -> hello\n")
}

func otherHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Trying to hit app2 from app1")
	addr := fmt.Sprint("http://app2.demo.svc:8002/home")

	resp, err := http.Get(addr)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintln(w, string(body))
}

func main() {
	log.Println("==== app1 ====")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/other", otherHandler)
	http.ListenAndServe(":8001", nil)
}
