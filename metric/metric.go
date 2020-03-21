package metric

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Monitor() {
	p := &http.ServeMux{}
	p.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8081", p))
}

func NewHistogram() prometheus.Histogram {
	h := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "shorturl",
		Name:      "histogram",
	})

	if err := prometheus.Register(h); err != nil {
		if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
			h = are.ExistingCollector.(prometheus.Histogram)
		} else {
			panic(err)
		}
	}

	return h
}

func NewCounter() prometheus.Counter {
	c := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "shorturl",
		Name:      "counter",
	})

	if err := prometheus.Register(c); err != nil {
		if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
			c = are.ExistingCollector.(prometheus.Counter)
		} else {
			panic(err)
		}
	}

	return c
}
