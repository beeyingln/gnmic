package app

import "github.com/prometheus/client_golang/prometheus"

var gnmiIncomingGetRequestsCacheHitTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "enc_nwi",
	Subsystem: "gnmi_proxy",
	Name:      "incoming_getrequests_cache_hit_total",
	Help:      "ENC NwI gNMI Proxy: Total number of cache hits for incoming gNMI GetRequest.",
})

var gnmiOutgoingSetRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "enc_nwi",
	Subsystem: "gnmi_proxy",
	Name:      "outgoing_setrequests_total",
	Help:      "ENC NwI gNMI Proxy: Total number of outgoing gNMI SetRequests from proxy to device.",
})

var gnmiIncomingSetRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "enc_nwi",
	Subsystem: "gnmi_proxy",
	Name:      "incoming_setrequests_total",
	Help:      "ENC NwI gNMI Proxy: Total number of incoming gNMI SetRequests from controller to proxy",
})

var gnmiIncomingSetRequestsBatchingTransactionFailureImpactedTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "enc_nwi",
	Subsystem: "gnmi_proxy",
	Name:      "incoming_setrequests_batching_transaction_failure_impacted_total",
	Help:      "ENC NwI gNMI Proxy: Total number of incoming gNMI SetRequests impacted by a failing batch",
})

func registerGnmiServerMetrics(reg *prometheus.Registry) error {
	if reg == nil {
		return nil
	}

	if err := reg.Register(gnmiIncomingGetRequestsCacheHitTotal); err != nil {
		return err
	}
	if err := reg.Register(gnmiOutgoingSetRequestsTotal); err != nil {
		return err
	}
	if err := reg.Register(gnmiIncomingSetRequestsBatchingTransactionFailureImpactedTotal); err != nil {
		return err
	}
	if err := reg.Register(gnmiIncomingSetRequestsTotal); err != nil {
		return err
	}

	return nil
}

func incrementGnmiIncomingGetRequestsCacheHitTotalMetric() {
	gnmiIncomingGetRequestsCacheHitTotal.Inc()
}

func incrementGnmiOutgoingSetRequestsTotalMetric() {
	gnmiOutgoingSetRequestsTotal.Inc()
}

func incrementGnmiIncomingSetRequestsBatchingTransactionFailureImpactedTotalMetric(x float64) {
	gnmiIncomingSetRequestsBatchingTransactionFailureImpactedTotal.Add(x)
}

func incrementGnmiIncomingSetRequestsTotalMetric(x float64) {
	gnmiIncomingSetRequestsTotal.Add(x)
}
