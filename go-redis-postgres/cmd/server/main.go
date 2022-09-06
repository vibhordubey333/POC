package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//go:generate openssl genrsa -out server.key 2048
//go:generate openssl ecparam -genkey -name secp384r1 -out server.key
//go:generate openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
func main() {
	http.HandleFunc("/products", httpHandler)
	http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
}

func httpHandler(w http.ResponseWriter, req *http.Request) {
	response, errorResponse := getProducts()
	if errorResponse != nil {
		fmt.Fprintf(w, errorResponse.Error()+"\r\n")
	} else {
		encoding := json.NewEncoder(w)
		encoding.SetIndent("", " ")
		if errorResponse := encoding.Encode(response); errResponse != nil {
			fmt.Println(errorResponse.Error())
		}
	}
}
