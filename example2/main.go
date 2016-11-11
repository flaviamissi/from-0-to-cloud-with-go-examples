package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/vrischmann/go-metrics-influxdb"
)

func main() {

	r := metrics.NewRegistry()
	metrics.RegisterDebugGCStats(r)
	metrics.RegisterRuntimeMemStats(r)

	go metrics.CaptureDebugGCStats(r, time.Second*5)
	go metrics.CaptureRuntimeMemStats(r, time.Second*5)

	go influxdb.InfluxDB(
		r,                                                        // metrics registry
		time.Second*5,                                            // interval
		fmt.Sprintf("http://%s:8086", os.Getenv("INFLUXDB_URL")), // the InfluxDB url
		"helloWorld", // your InfluxDB database
		"",           // your InfluxDB user
		"",           // your InfluxDB password
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		body := fmt.Sprintf(`<h1>Hello web ✖‿✖</h1>
	<h2>My host name is %s!</h2>`, hostname)
		fmt.Fprint(w, body)
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:4000", nil))
}
