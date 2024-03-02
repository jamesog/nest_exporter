package main

import (
	"github.com/jamesog/nest_exporter/starling"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	protectLabels = []string{"id", "name", "where"}

	protectCODetected = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemProtect, "carbon_monoxide_detected"),
		"Carbon Monoxide detected",
		protectLabels,
		nil,
	)
	protectManualTestActive = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemProtect, "manual_test_active"),
		"Manual Test is active",
		protectLabels,
		nil,
	)
	protectOccupancyDetected = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemProtect, "occupancy_detected"),
		"Occupancy detected",
		protectLabels,
		nil,
	)
	protectSmokeDetected = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemProtect, "smoke_detected"),
		"Smoke detected",
		protectLabels,
		nil,
	)

	// Metrics not directly exposed by Starling, but we compute them.
	protectBatteryLow = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemProtect, "battery_low"),
		"Battery is low",
		protectLabels,
		nil,
	)
)

func protectMetrics(protect starling.ProtectProperties, ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		protectCODetected,
		prometheus.GaugeValue,
		boolToFloat64(protect.CODetected),
		protect.ID, protect.Name, protect.Where,
	)
	ch <- prometheus.MustNewConstMetric(
		protectManualTestActive,
		prometheus.GaugeValue,
		boolToFloat64(protect.ManualTestActive),
		protect.ID, protect.Name, protect.Where,
	)
	ch <- prometheus.MustNewConstMetric(
		protectOccupancyDetected,
		prometheus.GaugeValue,
		boolToFloat64(protect.OccupancyDetected),
		protect.ID, protect.Name, protect.Where,
	)
	ch <- prometheus.MustNewConstMetric(
		protectSmokeDetected,
		prometheus.GaugeValue,
		boolToFloat64(protect.SmokeDetected),
		protect.ID, protect.Name, protect.Where,
	)

	// Computed metrics
	ch <- prometheus.MustNewConstMetric(
		protectBatteryLow,
		prometheus.GaugeValue,
		boolToFloat64(protect.batteryStatus == "low"),
		protect.ID, protect.Name, protect.Where,
	)
}
