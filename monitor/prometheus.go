package monitor

import (
	"fmt"
	"js_statistics/config"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Start 开启监控服务
func Start() {
	cfg := config.GetConfig()
	port := fmt.Sprintf(":%d", cfg.Prometheus.Port)
	log.Printf("monitor service start, listening on port: %d", cfg.Prometheus.Port)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(port, nil))
}
