package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TaskDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "task_duration_seconds",
			Help:    "Time spent processing tasks in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"task_type"},
	)

	TaskProcessedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "task_processed_total",
			Help: "Total number of processed tasks",
		},
		[]string{"task_type"},
	)

	TaskErrorCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "task_errors_total",
			Help: "Total number of task processing errors",
		},
		[]string{"task_type"},
	)
)
