package main

import (
	"github.com/jamesog/nest_exporter/starling"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"

	"net/http"
	"os"
	"time"

	flag "github.com/spf13/pflag"
)

const (
	namespace           = "nest"
	subsystemThermostat = "thermostat"
	subsystemThermostat = "protect"
)

var (
	up = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Status of the connection to the API",
		nil,
		nil,
	)
)

func boolToFloat64(b bool) float64 {
	if b {
		return 1.0
	}
	return 0
}

type Collector struct {
	starlingClient starling.Client
}

func NewCollector(client starling.Client) *Collector {
	return &Collector{client}
}

func (e Collector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(e, ch)
}

func (e Collector) Collect(ch chan<- prometheus.Metric) {
	status, err := e.starlingClient.Status()
	if err != nil {
		log.Err(err).Msg("error fetching status")
	}

	upValue := 0.0
	if status != nil && status.APIReady {
		upValue = 1.0
	}
	ch <- prometheus.MustNewConstMetric(
		up,
		prometheus.GaugeValue,
		upValue,
	)

	devices, err := e.starlingClient.Devices()
	if err != nil {
		log.Err(err).Msg("error getting devices")
		return
	}

	for _, device := range devices.Devices {
		switch device.Type {
		case "thermostat":
			t, err := e.starlingClient.ThermostatProperties(device.ID)
			if err != nil {
				continue
			}
			thermostatMetrics(*t, ch)
		case "protect":
			t, err := e.starlingClient.ProtectProperties(device.ID)
			if err != nil {
				continue
			}
			protectMetrics(*t, ch)
		default:
			log.Warn().Str("type", device.Type).Msg("unsupported device type")
		}
	}
	// TODO: Check status endpoint, return either metric or label for whether connectedToNest is true
	// false should note that it's a cached response
}

func requestLog(next http.Handler) http.Handler {
	h := hlog.NewHandler(log.Logger)
	accessHandler := hlog.AccessHandler(func(r *http.Request, status int, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("proto", r.Proto).
			Str("scheme", r.URL.Scheme).
			Str("host", r.Host).
			Str("uri", r.RequestURI).
			Str("user_agent", r.UserAgent()).
			Int("status_code", status).
			Int("response_size", size).
			Dur("duration", duration).
			Send()
	})
	return h(accessHandler(next))
}

func main() {
	starlingAPIKey := os.Getenv("STARLING_API_KEY")
	starlingAPIURLEnv := os.Getenv("STARLING_API_URL")

	starlingAPIFlag := flag.String("starling.api", "", "The base URL of the Starling API (overrides STARLING_API_URL)")
	logLevel := flag.String("log.level", "info", "The level of logging detail")
	listen := flag.String("listen", ":3081", "The address:port to listen on")
	flag.Parse()

	if level, err := zerolog.ParseLevel(*logLevel); err == nil {
		zerolog.SetGlobalLevel(level)
	}

	var starlingAPI string
	switch {
	case *starlingAPIFlag != "":
		starlingAPI = *starlingAPIFlag
	case starlingAPIURLEnv != "":
		starlingAPI = starlingAPIURLEnv
	default:
		log.Fatal().Msg("Starling API not set; must set either STARLING_API_URL or --starling.api")
	}

	if starlingAPIKey == "" {
		log.Fatal().Msg("Starling API key not set; must set STARLING_API_URL")
	}

	client := starling.NewClient(starlingAPI, starlingAPIKey)
	c := NewCollector(client)
	prometheus.MustRegister(collectors.NewBuildInfoCollector())
	prometheus.MustRegister(c)

	http.Handle("/metrics", requestLog(promhttp.Handler()))
	log.Fatal().Err(http.ListenAndServe(*listen, nil)).Send()
}
