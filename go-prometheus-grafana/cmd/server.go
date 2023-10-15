package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	deviceList = []Device{}
)

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

func init() {
	deviceList = []Device{
		{
			ID:       0,
			Mac:      "34-34-3422-2243-43",
			Firmware: "1.0",
		},
		{
			ID:       1,
			Mac:      "34-34-3422-2243-44",
			Firmware: "2.0",
		},
	}
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/devices", GetDevices)
	httpErrorResponse := http.ListenAndServe(":8081", nil)
	if httpErrorResponse != nil {
		log.Fatalln("Error while starting server:", httpErrorResponse)
	}
}

func GetDevices(w http.ResponseWriter, r *http.Request) {
	//Converting go struct to json string
	deviceJson, jsonErrorResponse := json.Marshal(deviceList)
	if jsonErrorResponse != nil {
		http.Error(w, jsonErrorResponse.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(deviceJson)
}
