package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"vibhordubey333/go-prometheus-grafana/prometheus/metrics"
)

var (
	deviceList = make([]Device, 0)
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

/*
- In the main() function, create a non-global registry without any pre-registered Collectors.
- Then create metrics using the NewMetrics function.
- Now we can use the devices property of the metrics struct and set it to the current number of connected devices. For that, we simply set it to the number of items in the devices slice.
- Let's also create a custom prometheus handler with the newly created register.
- We also need to update the /metrics handler to promHandler.
*/
func main() {
	registryObject := prometheus.NewRegistry()

	newMetricsObject := metrics.NewMetrics(registryObject)
	fmt.Println(len(deviceList))

	newMetricsObject.Devices.Set(float64(3))

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
