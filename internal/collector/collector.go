package collector

import (
	"github.com/davidsugianto/ephemeral-port-exporter/internal/system"
	"github.com/prometheus/client_golang/prometheus"
)

type EphemeralPortCollector struct {
	totalPorts     *prometheus.Desc
	usedPorts      *prometheus.Desc
	availablePorts *prometheus.Desc
	usageRatio     *prometheus.Desc
	scrapeSuccess  *prometheus.Desc
}

func NewEphemeralPortCollector() *EphemeralPortCollector {
	return &EphemeralPortCollector{
		totalPorts:     prometheus.NewDesc("ephemeral_ports_total", "Total ephemeral ports", nil, nil),
		usedPorts:      prometheus.NewDesc("ephemeral_ports_used", "Number of ephemeral ports in use", nil, nil),
		availablePorts: prometheus.NewDesc("ephemeral_ports_available", "Total ephemeral ports available", nil, nil),
		usageRatio:     prometheus.NewDesc("ephemeral_ports_usage_ratio", "Ratio of ephemeral ports used", nil, nil),
		scrapeSuccess:  prometheus.NewDesc("ephemeral_port_exporter_scrape_success", "Whether scraping ephemeral ports succeeded (1 = yes, 0 = no)", nil, nil),
	}
}

func (c *EphemeralPortCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.totalPorts
	ch <- c.usedPorts
	ch <- c.availablePorts
	ch <- c.usageRatio
	ch <- c.scrapeSuccess
}

func (c *EphemeralPortCollector) Collect(ch chan<- prometheus.Metric) {
	start, end, err := system.GetEphemeralPortRange()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.scrapeSuccess, prometheus.GaugeValue, 0)
		return
	}

	total := end - start + 1
	used, err := system.CountUsedEphemeralPorts(start, end)
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.scrapeSuccess, prometheus.GaugeValue, 0)
		return
	}

	available := total - used
	usageRatio := float64(used) / float64(total)

	ch <- prometheus.MustNewConstMetric(c.totalPorts, prometheus.GaugeValue, float64(total))
	ch <- prometheus.MustNewConstMetric(c.usedPorts, prometheus.GaugeValue, float64(used))
	ch <- prometheus.MustNewConstMetric(c.availablePorts, prometheus.GaugeValue, float64(available))
	ch <- prometheus.MustNewConstMetric(c.usageRatio, prometheus.GaugeValue, usageRatio)
	ch <- prometheus.MustNewConstMetric(c.scrapeSuccess, prometheus.GaugeValue, 1)
}
