package app

import "github.com/prometheus/client_golang/prometheus"

var gnmiServerGrpcGetCacheHitTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "gnmic",
	Subsystem: "gnmi_server",
	Name:      "grpc_get_cache_hit_total",
	Help:      "gNMI Server : Total number of gRPC Get cache hit.",
})

func registerGnmiServerMetrics(reg *prometheus.Registry) error {
	if reg == nil {
		return nil
	}

	if err := reg.Register(gnmiServerGrpcGetCacheHitTotal); err != nil {
		return err
	}

	return nil
}

func incrementGnmiServerGrpcGetCacheHitTotalMetric() {
	gnmiServerGrpcGetCacheHitTotal.Inc()
}
