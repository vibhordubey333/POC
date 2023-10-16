package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"vibhordubey333/go-prometheus-grafana/prometheus/metrics"
)

var (
	deviceList = make([]Device, 0)
	version    string
)

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

func init() {
	version = "3.0.0"
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

	metricsObject := metrics.NewMetrics(registryObject)
	log.Println("Starting Server...")

	metricsObject.Devices.Set(float64(len(deviceList)))
	metricsObject.Info.With(prometheus.Labels{"version": version})

	/*
		    ---Output-of-metrics-api---
			vibhor@vibhor-virtualbox:~$ curl localhost:8081/metrics
			# HELP POC1_Info Details of environment
			# TYPE POC1_Info gauge
			POC1_Info{version="3.0.0"} 0
			# HELP POC1_device_connected_list List of connected devices
			# TYPE POC1_device_connected_list gauge
			POC1_device_connected_list 2

	*/
	/*
		Since we don't have 3rd party routers and frameworks, We will use multiple servers using goroutines.
			- Why two ?
				- Serve /metrics API
				- Serve /devices API

	*/
	deviceMuxObject := http.NewServeMux()
	deviceMuxObject.HandleFunc("/devices", GetDevices)

	prometheusMuxObject := http.NewServeMux()
	prometheusHandler := promhttp.HandlerFor(registryObject, promhttp.HandlerOpts{})
	prometheusMuxObject.Handle("/metrics", prometheusHandler)

	//Device
	go func() {
		/*
				vibhor@vibhor-virtualbox:~/code-repositories/POC/go-prometheus-grafana (Refactor/go-prometheus-grafana *)$ curl -XGET localhost:8081/devices
			[{"id":0,"mac":"34-34-3422-2243-43","firmware":"1.0"},{"id":1,"mac":"34-34-3422-2243-44","firmware":"2.0"}]
		*/
		log.Fatal(http.ListenAndServe(":8080", deviceMuxObject))
	}()
	//Metric
	go func() {
		log.Fatal(http.ListenAndServe(":8081", prometheusMuxObject))
		/*
			vibhor@vibhor-virtualbox:~/code-repositories/POC/go-prometheus-grafana (Refactor/go-prometheus-grafana *)$ curl localhost:8081/metrics
			# HELP POC1_Info Details of environment
			# TYPE POC1_Info gauge
			POC1_Info{version="3.0.0"} 0
			# HELP POC1_device_connected_list List of connected devices
			# TYPE POC1_device_connected_list gauge
			POC1_device_connected_list 2

		*/
	}()

	//To indefinitely run program
	select {}

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
