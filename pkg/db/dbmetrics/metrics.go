package dbmetrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type IMetric interface {
	Complete(err error)
}

type metric struct {
	funcName     string
	metricPrefix string
	metricSuffix string
	start        time.Time

	dbLatency *prometheus.SummaryVec
	dbVolume  *prometheus.CounterVec
	dbErrors  *prometheus.CounterVec
}

func NewMetric(opts ...Option) IMetric {
	m := &metric{start: time.Now()}

	for _, o := range opts {
		o(m)
	}

	m.dbLatency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       m.getMetricName("db_latency"),
			Help:       "The latency quantiles for the given database request",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"func"},
	)

	m.dbVolume = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: m.getMetricName("db_volume"),
			Help: "Number of times a given database request was made",
		},
		[]string{"func"},
	)

	m.dbErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: m.getMetricName("db_errors"),
			Help: "Number of times a given database request failed",
		},
		[]string{"func"},
	)

	prometheus.MustRegister(m.dbVolume)
	prometheus.MustRegister(m.dbLatency)
	prometheus.MustRegister(m.dbErrors)

	if m.funcName != "" {
		m.dbVolume.With(prometheus.Labels{"func": m.funcName}).Inc()
	}

	return m
}

func (m *metric) Complete(err error) {
	if err != nil {
		m.dbErrors.With(prometheus.Labels{"func": m.funcName}).Inc()
	}
	m.dbLatency.WithLabelValues(m.funcName).Observe(float64(time.Since(m.start).Milliseconds()))
}

func (m *metric) getMetricName(v string) string {
	if m.metricPrefix != "" {
		v = m.metricPrefix + "_" + v
	}

	if m.metricSuffix != "" {
		v = v + "_" + m.metricSuffix
	}

	return v
}
