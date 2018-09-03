/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import "github.com/prometheus/client_golang/prometheus"

// Metrics holds the prometheus metrics used internally by the controller
type Metrics struct {
	QueueLength     *prometheus.GaugeVec
	ReconcileErrors *prometheus.CounterVec
	ReconcileTime   *prometheus.HistogramVec
}

// GetCollectors implements the metrics.Collector interface
func (m *Metrics) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		m.QueueLength,
		m.ReconcileErrors,
		m.ReconcileTime,
	}
}

// NewMetrics returns a new initialised Metrics
func NewMetrics() *Metrics {
	return &Metrics{
		QueueLength:     newQueueLengthMetric(),
		ReconcileErrors: newReconcileErrorsMetric(),
		ReconcileTime:   newReconcileTimeMetric(),
	}
}

func newQueueLengthMetric() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "controller_runtime_reconcile_queue_length",
		Help: "Length of reconcile queue per controller",
	}, []string{"controller"})
}

func newReconcileErrorsMetric() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "controller_runtime_reconcile_errors_total",
		Help: "Total number of reconcile errors per controller",
	}, []string{"controller"})
}

func newReconcileTimeMetric() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "controller_runtime_reconcile_time_second",
		Help: "Length of time per reconcile per controller",
	}, []string{"controller"})
}
