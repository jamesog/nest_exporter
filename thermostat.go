package main

import (
	"github.com/jamesog/nest_exporter/starling"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	thermostatLabels = []string{"id", "name", "where"}

	thermostatCurrentTemperature = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemThermostat, "current_temperature_celsius"),
		"Current temperature in celsius",
		thermostatLabels,
		nil,
	)
	thermostatTargetTemperature = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemThermostat, "target_temperature_celsius"),
		"Target temperature in celsius",
		thermostatLabels,
		nil,
	)
	thermostatEcoMode = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemThermostat, "eco_mode"),
		"Eco mode enabled",
		thermostatLabels,
		nil,
	)
	thermostatHumidityPct = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemThermostat, "humidity_percent"),
		"Relative humidity at the thermostat",
		thermostatLabels,
		nil,
	)

	// Metrics not directly exposed by Starling, but we compute them.
	thermostatIsEnabled = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemThermostat, "is_enabled"),
		"HVAC mode is not set to off",
		thermostatLabels,
		nil,
	)
	thermostatIsHeating = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemThermostat, "is_heating"),
		"HVAC mode is set to heating and heating is on",
		thermostatLabels,
		nil,
	)
	thermostatIsCooling = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystemThermostat, "is_cooling"),
		"HVAC mode is set to cooling and cooling is on",
		thermostatLabels,
		nil,
	)
)

func thermostatMetrics(thermostat starling.ThermostatProperties, ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		thermostatCurrentTemperature,
		prometheus.GaugeValue,
		thermostat.CurrentTemperature,
		thermostat.ID, thermostat.Name, thermostat.Where,
	)
	ch <- prometheus.MustNewConstMetric(
		thermostatTargetTemperature,
		prometheus.GaugeValue,
		thermostat.TargetTemperature,
		thermostat.ID, thermostat.Name, thermostat.Where,
	)
	ch <- prometheus.MustNewConstMetric(
		thermostatEcoMode,
		prometheus.GaugeValue,
		boolToFloat64(thermostat.EcoMode),
		thermostat.ID, thermostat.Name, thermostat.Where,
	)
	ch <- prometheus.MustNewConstMetric(
		thermostatHumidityPct,
		prometheus.GaugeValue,
		thermostat.HumidityPercent,
		thermostat.ID, thermostat.Name, thermostat.Where,
	)

	// Computed metrics
	ch <- prometheus.MustNewConstMetric(
		thermostatIsEnabled,
		prometheus.GaugeValue,
		boolToFloat64(thermostat.HVACMode != "off"),
		thermostat.ID, thermostat.Name, thermostat.Where,
	)
	ch <- prometheus.MustNewConstMetric(
		thermostatIsHeating,
		prometheus.GaugeValue,
		boolToFloat64(thermostat.HVACMode == "heat" && thermostat.HVACState == "heating"),
		thermostat.ID, thermostat.Name, thermostat.Where,
	)
	ch <- prometheus.MustNewConstMetric(
		thermostatIsCooling,
		prometheus.GaugeValue,
		boolToFloat64(thermostat.HVACMode == "cool" && thermostat.HVACState == "cooling"),
		thermostat.ID, thermostat.Name, thermostat.Where,
	)
}
