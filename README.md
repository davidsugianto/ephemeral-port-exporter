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
- Reads from:
  - `/proc/sys/net/ipv4/ip_local_port_range`
  - `/proc/net/tcp`, `/proc/net/tcp6`

---

## 📊 Exported Metrics

| Metric Name                              | Type   | Description                                      |
|------------------------------------------|--------|--------------------------------------------------|
| `ephemeral_ports_used`                   | Gauge  | Number of ephemeral ports currently in use       |
| `ephemeral_ports_total`                  | Gauge  | Total ephemeral ports available on the system    |
| `ephemeral_ports_usage_ratio`            | Gauge  | Ratio of used/total ephemeral ports              |
| `ephemeral_port_exporter_scrape_success` | Gauge  | 1 if scrape was successful, 0 otherwise          |

## Example Output

```
# HELP ephemeral_ports_used Number of ephemeral ports in use
# TYPE ephemeral_ports_used gauge
ephemeral_ports_used 43
```

## 🚀 Getting Started

### Requirements
- Linux system (uses /proc)
- Go 1.18 or later
- Prometheus for scraping

### 🔧 Build & Run

```
git clone https://github.com/yourname/ephemeral-port-exporter.git
cd ephemeral-port-exporter

go build -o exporter ./cmd/exporter
./exporter
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
│   └── system/          # Linux port range parsing
└── go.mod
```

## 📄 License
MIT License. See LICENSE for details.
