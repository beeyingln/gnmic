package app

import "github.com/prometheus/client_golang/prometheus"

var gnmiServerGrpcGetCacheHitTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "gnmic",
	Subsystem: "gnmi_server",
	Name:      "grpc_get_cache_hit_total",
	Help:      "gNMI Server : Total number of gRPC Get cache hit.",
})

var gnmiServerGrpcDeviceSetReqTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "gnmic",
	Subsystem: "gnmi_server",
	Name:      "grpc_device_set_req_total",
	Help:      "gNMI Server : Total number of gRPC Set requests towards the actual device.",
})

var gnmiServerBatchingTransactionFailureImpactedTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "gnmic",
	Subsystem: "gnmi_server",
	Name:      "batching_transaction_failure_impacted_total",
	Help:      "gNMI Server : Total number of Set requests impacted by a failing batch",
})

func registerGnmiServerMetrics(reg *prometheus.Registry) error {
	if reg == nil {
		return nil
	}

	if err := reg.Register(gnmiServerGrpcGetCacheHitTotal); err != nil {
		return err
	}
	if err := reg.Register(gnmiServerGrpcDeviceSetReqTotal); err != nil {
		return err
	}
	if err := reg.Register(gnmiServerBatchingTransactionFailureImpactedTotal); err != nil {
		return err
	}

	return nil
}

func incrementGnmiServerGrpcGetCacheHitTotalMetric() {
	gnmiServerGrpcGetCacheHitTotal.Inc()
}

func incrementGnmiServerGrpcDeviceSetReqTotalMetric() {
	gnmiServerGrpcDeviceSetReqTotal.Inc()
}

func addGnmiServerBatchingTransactionFailureImpactedTotalMetric(x float64) {
	gnmiServerBatchingTransactionFailureImpactedTotal.Add(x)
}
