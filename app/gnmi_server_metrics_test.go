package app

import (
	"testing"

	dto "github.com/prometheus/client_model/go"
)

func TestGnmiServerMetrics(t *testing.T) {
	var gnmiServerGrpcGetCacheHitTotalMetric *dto.Metric = new(dto.Metric)
	gnmiServerGrpcGetCacheHitTotal.Write(gnmiServerGrpcGetCacheHitTotalMetric)
	var initialValue float64 = 0.0
	if *gnmiServerGrpcGetCacheHitTotalMetric.GetCounter().Value != initialValue {
		t.Errorf("got %f, expected %f", *gnmiServerGrpcGetCacheHitTotalMetric.GetCounter().Value, initialValue)
	}

	incrementGnmiServerGrpcGetCacheHitTotalMetric()
	gnmiServerGrpcGetCacheHitTotal.Write(gnmiServerGrpcGetCacheHitTotalMetric)
	initialValue++
	if *gnmiServerGrpcGetCacheHitTotalMetric.GetCounter().Value != initialValue {
		t.Errorf("got %f, expected %f", *gnmiServerGrpcGetCacheHitTotalMetric.GetCounter().Value, initialValue)
	}
}
