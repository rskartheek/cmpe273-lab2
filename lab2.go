package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Request struct
type Request struct {
	Name string `json:"name"`
}

//Response struct
type Response struct {
	Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func postFunction(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var requestObject Request
	json.NewDecoder(req.Body).Decode(&requestObject)
	var respObject Response
	respObject.Greeting = "Hello " + requestObject.Name
	uj, _ := json.Marshal(respObject)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(201)
	fmt.Fprintf(rw, "%s", uj)
}
func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hello", postFunction)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
