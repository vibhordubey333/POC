package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vibhordubey333/poc/go-redis-postgres/pkg/products"
)

//go:generate openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout go-server.key -out go-server.crt
func main() {
	fmt.Println("Starting Server:")
	http.HandleFunc("/products", httpProductsHandler)
	serverObject := &http.Server{
		Addr:    ":8888",
		Handler: nil,
	}

	//TODO: Use HTTP server
	errorResponse := serverObject.ListenAndServe()
	if errorResponse != nil {
		fmt.Println("Error in ListenAndServe", errorResponse)
	}

}

func httpHelloWorld(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Server is up and running.")
}

func httpProductsHandler(w http.ResponseWriter, req *http.Request) {
	response, errorResponse := products.GetProducts()
	if errorResponse != nil {
		fmt.Fprintf(w, errorResponse.Error()+"\r\n")
	} else {
		encoding := json.NewEncoder(w)
		encoding.SetIndent("", " ")
		if errorResponse := encoding.Encode(response); errorResponse != nil {
			fmt.Println(errorResponse.Error())
		}
	}
}
