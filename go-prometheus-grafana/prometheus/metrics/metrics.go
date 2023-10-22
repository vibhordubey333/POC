package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

/*
Gauge represents a single numerical value when GaugeVec is a collection that bundles a set of Gauges with the same name but different labels.

For example, if you want to count all connected devices and you don't care about the different types, use a Gauge. On the other hand, if you have
different device types, such as routers, switches, and access points, and you want to count them separately, use GaugeVec with a type label. You'll see a
bunch of examples during this tutorial.
*/
type Metrics struct {
	Devices prometheus.Gauge
	Info    *prometheus.GaugeVec
}

/*
		We need to create devices metric using the NewGauge function.
		- A Namespace is just a metric prefix; usually, you use a single word that matches the name of your app. In my case, it's POC-1.
		- Then the metric Name. It's very important to follow the naming conventions provided by Prometheus. You can find it on the official website. Let's call it
	      device-connected-list.
		- You also need to include a metric description.
		- Then register it with the prometheus registry and return a pointer
*/
func NewMetrics(reg prometheus.Registerer) *Metrics {
	metricsObject := &Metrics{
		Devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace:   "POC1",
			Name:        "device_connected_list",
			Help:        "List of connected devices",
			ConstLabels: nil,
		}),
		Info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "POC1",
			Name:      "Info",
			Help:      "Details of environment",
		},

			[]string{"version"}),
	}
	reg.MustRegister(metricsObject.Devices, metricsObject.Info)
	return metricsObject
}

/*

   - In the main() function, create a non-global registry without any pre-registered Collectors.
   - Then create metrics using the NewMetrics function.
   - Now we can use the devices property of the metrics struct and set it to the current number of connected devices. For that, we simply set it to the number of items in the devices slice.
   - Let's also create a custom prometheus handler with the newly created register.
   - We also need to update the /metrics handler to promHandler.

*/
