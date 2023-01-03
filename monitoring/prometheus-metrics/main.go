package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	myCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "adamik_mycounter",
		Help: "sth",
	})

	myGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "adamik_mygauge",
		Help: "sth",
	})

	mySummary = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "adamik_summaru",
		Help: "sth",
	})

	myHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "adamik_histogram",
		Help:    "sth",
		Buckets: []float64{10, 50, 100, 200, 500, 1000},
	})
)

func recordMetrics() {
	go func() {
		for {
			myValue := rand.Int63n(1000)

			myCounter.Inc()
			myGauge.Set(float64(myValue))
			mySummary.Observe(float64(myValue))
			myHistogram.Observe(float64(myValue))

			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9876", nil)
}
