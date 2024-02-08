package alphasms

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Exporter struct {
	client *Client

	balance prometheus.Gauge
	error   prometheus.Gauge
	date    prometheus.Gauge
}

func NewExporter(client *Client) *Exporter {
	return &Exporter{
		client: client,
		balance: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "alphasms",
			Name:      "balance_amount",
			Help:      "The current balance amount.",
		}),
		error: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "alphasms",
			Name:      "error",
			Help:      "The current error code.",
		}),
		date: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "alphasms",
			Name:      "date",
			Help:      "The last successful call date.",
		}),
	}
}

// Descriptors
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.balance.Describe(ch)
	e.error.Describe(ch)
	e.date.Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	resp, err := e.client.GetUserBalance()
	if err != nil {
		// Handle error situation
	}

	balance, err := strconv.ParseFloat(resp.Data.Balance, 64)
	if err != nil {
		// Handle error situation
	}

	e.balance.Set(balance)
	e.error.Set(float64(resp.Error))

	// Convert date string to unix timestamp
	t, err := time.Parse(time.RFC3339, resp.Data.Validity)
	if err != nil {
		// Handle error situation
	}

	e.date.Set(float64(t.Unix()))

	e.balance.Collect(ch)
	e.error.Collect(ch)
	e.date.Collect(ch)
}
