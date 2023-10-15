package main

import (
	"bytes"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	URL = "http://localhost:8080"
)

// Need to register the counter so prometheus can collect this metric.
var userStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "API_Request_User_Count",
		Help: "User Count",
	},
	[]string{"user", "status"},
)

// Registering the counter so prometheus can collect this metric.
func init() {
	prometheus.MustRegister(userStatus)
}

type MyRequest struct {
	User string
}

func producer() {
	poolOfUsers := []string{"Alpha", "Mike", "Foxtrot"}
	for {
		postBodyRequest, _ := json.Marshal(MyRequest{
			User: poolOfUsers[rand.Intn(len(poolOfUsers))],
		})
		requestBody := bytes.NewBuffer(postBodyRequest)
		http.Post(URL, "application/json", requestBody)
		time.Sleep(time.Second * 3)
	}
}

func server(w http.ResponseWriter, r *http.Request) {
	var status, user string

	defer func() {
		userStatus.WithLabelValues(user, status).Inc()
	}()

	var requestObject MyRequest
	json.NewDecoder(r.Body).Decode(&requestObject)

	if rand.Float32() > 0.8 {
		status = "200"
	} else {
		status = "501"
	}
	user = requestObject.User
	log.Println(user, status)
	w.Write([]byte(status))
}

//Un-comment to run local server.
/*
func main() {
	go producer()

	http.Handle("/metrics", prometheusHttp.Handler())
	http.HandleFunc("/", server)

	err := http.ListenAndServe(":8080", nil)
	log.Println("Server started", err)
	if err != nil {
		log.Fatal("Listen & Serve: ", err)
	}

}
*/
//https://medium.com/@alcbotta/monitoring-you-golang-server-with-prometheus-and-grafana-97e64bb1d0e9
//https://antonputra.com/monitoring/monitor-golang-with-prometheus/#gauge
