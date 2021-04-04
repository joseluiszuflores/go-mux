package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
)

const httpAddr = ":8080"

const (
	get = "GET"
	post = "POST"
	put = "PUT"
)

func main()  {
	fmt.Println("Server runnig in ", httpAddr)
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",helloHandlerGet )
	log.Fatal(http.ListenAndServe(httpAddr, mux))
}

func helloHandlerGet(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	if method == get {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello get data"))
		return
	}
	var str string
	if method == post {
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&str)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(str))
		return
	}
}
