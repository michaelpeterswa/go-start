# go-start

A template repository for Go projects with built-in observability support.

## Features

- Structured JSON logging with configurable log levels via [slog](https://pkg.go.dev/log/slog)
- OpenTelemetry metrics and tracing via [ootel](https://alpineworks.io/ootel)
- Runtime and host metrics instrumentation
- Environment-based configuration via [env](https://github.com/caarlos0/env)
- Multi-stage Docker build with distroless final image
- Local development setup with Grafana LGTM stack

## Configuration

Configuration is managed through environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `LOG_LEVEL` | Logging level (debug, info, warn, error) | `error` |
| `METRICS_ENABLED` | Enable Prometheus metrics | `true` |
| `METRICS_PORT` | Port for metrics endpoint | `8081` |
| `LOCAL` | Use OTLP gRPC exporter instead of Prometheus | `false` |
| `TRACING_ENABLED` | Enable distributed tracing | `false` |
| `TRACING_SAMPLERATE` | Trace sampling rate | `0.01` |
| `TRACING_SERVICE` | Service name for traces | `katalog-agent` |
| `TRACING_VERSION` | Service version for traces | - |

## Getting Started

### Run Locally with Docker Compose

```bash
docker-compose up
```

This starts the application along with the Grafana LGTM (Loki, Grafana, Tempo, Mimir) stack for local observability:

- **Application**: Port 8081 (metrics)
- **Grafana UI**: http://localhost:3000
- **OTLP gRPC**: Port 4317
- **OTLP HTTP**: Port 4318

### Build and Run

```bash
go build -o go-start ./cmd/go-start
./go-start
```

## Project Structure

```
.
├── cmd/go-start/       # Application entrypoint
├── internal/
│   ├── config/         # Environment-based configuration
│   └── logging/        # Logging utilities
├── docker/
│   └── grafana/        # Grafana dashboard provisioning
├── Dockerfile          # Multi-stage build
└── docker-compose.yml  # Local development stack
```

## CI/CD

Pull requests are validated with:

- **commitlint**: Conventional commit message enforcement
- **golangci-lint**: Go linting
- **yamllint**: YAML linting
- **hadolint**: Dockerfile linting
- **go test**: Unit tests
