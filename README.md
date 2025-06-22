# 📘 Ephemeral Port Exporter

A Prometheus exporter written in Go that exposes metrics about ephemeral (dynamic) ports in use on Linux systems.

Built with:

- 🟦 [Go](https://golang.org)
- 🔀 [Chi router](https://github.com/go-chi/chi)
- 📈 [Prometheus Go client](https://github.com/prometheus/client_golang)

---

## 📦 Features

- Reports:
  - Number of ephemeral ports in use
  - Total available ephemeral port range
  - Ratio of usage
  - Scrape success status
- Reference from:
  - `/proc/sys/net/ipv4/ip_local_port_range`
  - `/proc/net/tcp`, `/proc/net/tcp6`
  - `ss`

---

## 📊 Exported Metrics

| Metric Name                              | Type   | Description                                      |
|------------------------------------------|--------|--------------------------------------------------|
| `ephemeral_ports_used`                   | Gauge  | Number of ephemeral ports currently in use       |
| `ephemeral_ports_total`                  | Gauge  | Total ephemeral ports                            |
| `ephemeral_ports_available`              | Gauge  | Total ephemeral ports available on the system    |
| `ephemeral_ports_usage_ratio`            | Gauge  | Ratio of used/total ephemeral ports              |
| `ephemeral_port_exporter_scrape_success` | Gauge  | 1 if scrape was successful, 0 otherwise          |

## Example Output

```
# HELP ephemeral_port_exporter_scrape_success Whether scraping ephemeral ports succeeded (1 = yes, 0 = no)
# TYPE ephemeral_port_exporter_scrape_success gauge
ephemeral_port_exporter_scrape_success 1
# HELP ephemeral_ports_available Total ephemeral ports available
# TYPE ephemeral_ports_available gauge
ephemeral_ports_available 28231
# HELP ephemeral_ports_total Total ephemeral ports
# TYPE ephemeral_ports_total gauge
ephemeral_ports_total 28232
# HELP ephemeral_ports_usage_ratio Ratio of ephemeral ports used
# TYPE ephemeral_ports_usage_ratio gauge
ephemeral_ports_usage_ratio 3.5420799093227545e-05
# HELP ephemeral_ports_used Number of ephemeral ports in use
# TYPE ephemeral_ports_used gauge
ephemeral_ports_used 1
```

## 🚀 Getting Started

### Requirements
- Linux system (uses /proc)
- Go 1.18 or later
- Prometheus for scraping

### 🔧 Build & Run

```
git clone https://github.com/davidsugianto/ephemeral-port-exporter.git
cd ephemeral-port-exporter

make deps
make run
```
Default port: :2112

### 🧪 Test

```
curl http://localhost:2112/health
# OK

curl http://localhost:2112/metrics
# Prometheus metrics output
```

### 📍 Prometheus Config Example

```
scrape_configs:
  - job_name: "ephemeral-port-exporter"
    static_configs:
      - targets: ["localhost:2112"]
```

## 🛠️ Project Structure

```
ephemeral-port-exporter/
├── cmd/                 # Entrypoint
├── internal/            # Logic & modules
│   ├── collector/       # Prometheus collector
│   ├── router/          # Chi router
│   └── system/          # Linux port parsing
└── go.mod
```

## 📄 License
MIT License. See LICENSE for details.
