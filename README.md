![Repo visits](https://hits.sh/github.com/Manikeshmk/GOlang.svg?label=repo%20visits)
![GitHub stars](https://img.shields.io/github/stars/Manikeshmk/GOlang?style=logo&logo=github&label=⭐%20Stars) 
![GitHub forks](https://img.shields.io/github/forks/Manikeshmk/GOlang?style=social)

# Silent Meeting Summarizer

<p align="center">
  <img src="https://capsule-render.vercel.app/api?type=waving&height=220&color=0:00ADD8,45:111827,100:7C3AED&text=Silent%20Meeting%20Summarizer&fontColor=ffffff&fontSize=42&fontAlignY=38&desc=Local-first%20AI%20meeting%20intelligence%20for%20summaries,%20tasks,%20decisions,%20and%20conflicts&descAlignY=60&descSize=16" alt="Silent Meeting Summarizer banner" />
</p>

<p align="center">
  <a href="https://github.com/Manikeshmk/GOlang/actions"><img src="https://github.com/Manikeshmk/GOlang/workflows/CI%2FCD%20Pipeline/badge.svg" alt="CI/CD Pipeline" /></a>
  <a href="https://codecov.io/gh/Manikeshmk/GOlang"><img src="https://codecov.io/gh/Manikeshmk/GOlang/branch/main/graph/badge.svg" alt="Codecov" /></a>
  <a href="https://goreportcard.com/report/github.com/Manikeshmk/GOlang"><img src="https://goreportcard.com/badge/github.com/Manikeshmk/GOlang" alt="Go Report Card" /></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="MIT License" /></a>
  <img src="https://img.shields.io/badge/status-production--ready-16a34a?style=flat-square" alt="Production Ready" />
</p>

<p align="center">
  <img src="https://skillicons.dev/icons?i=go,docker,kubernetes,postgres,redis,kafka,prometheus,react,nextjs,ts,tailwind,githubactions" alt="Technology logos" />
</p>

<p align="center">
  <b>A production-grade AI meeting assistant that records audio, performs speaker diarization, generates summaries, extracts action items, detects conflict, and keeps processing local-first with optional cloud AI.</b>
</p>

---

## Table of Contents

- [Why It Exists](#why-it-exists)
- [Feature Highlights](#feature-highlights)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Docker Services](#docker-services)
- [Testing](#testing)
- [Monitoring](#monitoring)
- [Security](#security)
- [Deployment](#deployment)
- [Documentation](#documentation)
- [Contributing](#contributing)

## Why It Exists

Meetings create a lot of information, but most of it is hard to reuse: scattered decisions, unclear ownership, repeated topics, unresolved disagreements, and minutes that arrive too late. Silent Meeting Summarizer turns meeting audio into structured intelligence:

- clear summaries for different audiences
- speaker-aware transcripts
- decisions with confidence signals
- action items with owners and due dates
- conflict, confusion, sentiment, and repeated-topic analysis
- local-first AI pipelines with optional OpenAI/Ollama integrations

## Feature Highlights

<table>
  <tr>
    <td width="50%">
      <h3>Audio Intelligence</h3>
      <ul>
        <li>Real-time streamed audio ingestion</li>
        <li>Buffering and backpressure handling</li>
        <li>Speech-to-text with provider fallback</li>
        <li>Speaker diarization and participant tracking</li>
      </ul>
    </td>
    <td width="50%">
      <h3>Meeting Understanding</h3>
      <ul>
        <li>Concise, detailed, and executive summaries</li>
        <li>Task extraction with owners and deadlines</li>
        <li>Decision confidence analysis</li>
        <li>Topic clustering and wasted-time estimation</li>
      </ul>
    </td>
  </tr>
  <tr>
    <td width="50%">
      <h3>Risk Signals</h3>
      <ul>
        <li>Conflict detection and disagreement scoring</li>
        <li>Unresolved topic identification</li>
        <li>Confusion and clarification-request detection</li>
        <li>Sentiment tracking across the meeting</li>
      </ul>
    </td>
    <td width="50%">
      <h3>Production Foundation</h3>
      <ul>
        <li>JWT authentication and RBAC</li>
        <li>Prometheus metrics and Jaeger tracing</li>
        <li>Docker and Kubernetes manifests</li>
        <li>Structured logging and health checks</li>
      </ul>
    </td>
  </tr>
</table>

## Tech Stack

### Backend

<p>
  <img src="https://img.shields.io/badge/Go-1.23+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Gin-HTTP%20API-008ECF?style=for-the-badge&logo=gin&logoColor=white" alt="Gin" />
  <img src="https://img.shields.io/badge/gRPC-services-244C5A?style=for-the-badge&logo=grpc&logoColor=white" alt="gRPC" />
  <img src="https://img.shields.io/badge/WebSocket-realtime-010101?style=for-the-badge&logo=socketdotio&logoColor=white" alt="WebSocket" />
  <img src="https://img.shields.io/badge/Zap-logging-f97316?style=for-the-badge&logo=go&logoColor=white" alt="Zap logging" />
</p>

### Data and Messaging

<p>
  <img src="https://img.shields.io/badge/PostgreSQL-16+-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/Redis-7+-DC382D?style=for-the-badge&logo=redis&logoColor=white" alt="Redis" />
  <img src="https://img.shields.io/badge/NATS-messaging-27AAE1?style=for-the-badge&logo=natsdotio&logoColor=white" alt="NATS" />
  <img src="https://img.shields.io/badge/Apache%20Kafka-events-231F20?style=for-the-badge&logo=apachekafka&logoColor=white" alt="Apache Kafka" />
  <img src="https://img.shields.io/badge/ZooKeeper-coordination-6DB33F?style=for-the-badge&logo=apache&logoColor=white" alt="Apache ZooKeeper" />
</p>

### AI and Observability

<p>
  <img src="https://img.shields.io/badge/OpenAI-optional%20AI-412991?style=for-the-badge&logo=openai&logoColor=white" alt="OpenAI" />
  <img src="https://img.shields.io/badge/Ollama-local%20LLM-000000?style=for-the-badge&logo=ollama&logoColor=white" alt="Ollama" />
  <img src="https://img.shields.io/badge/Whisper-speech--to--text-111827?style=for-the-badge&logo=openai&logoColor=white" alt="Whisper" />
  <img src="https://img.shields.io/badge/Prometheus-metrics-E6522C?style=for-the-badge&logo=prometheus&logoColor=white" alt="Prometheus" />
  <img src="https://img.shields.io/badge/Jaeger-tracing-65A2C8?style=for-the-badge&logo=jaeger&logoColor=white" alt="Jaeger" />
</p>

### Frontend and Platform

<p>
  <img src="https://img.shields.io/badge/Next.js-14-000000?style=for-the-badge&logo=nextdotjs&logoColor=white" alt="Next.js" />
  <img src="https://img.shields.io/badge/React-18-61DAFB?style=for-the-badge&logo=react&logoColor=111827" alt="React" />
  <img src="https://img.shields.io/badge/TypeScript-5-3178C6?style=for-the-badge&logo=typescript&logoColor=white" alt="TypeScript" />
  <img src="https://img.shields.io/badge/Tailwind%20CSS-3-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white" alt="Tailwind CSS" />
  <img src="https://img.shields.io/badge/Docker-containers-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
  <img src="https://img.shields.io/badge/Kubernetes-orchestration-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white" alt="Kubernetes" />
  <img src="https://img.shields.io/badge/GitHub%20Actions-CI/CD-2088FF?style=for-the-badge&logo=githubactions&logoColor=white" alt="GitHub Actions" />
</p>

## Architecture

```mermaid
flowchart TB
    client["Web / Mobile / Integrations"]
    gateway["API Gateway / Ingress"]
    api["Go API<br/>Gin + REST + WebSocket"]
    worker["Background Workers"]
    ai["AI Services<br/>Whisper + Ollama + OpenAI"]
    repo["Repository Layer"]
    postgres[("PostgreSQL")]
    redis[("Redis Cache")]
    nats["NATS"]
    kafka["Kafka"]
    prom["Prometheus"]
    jaeger["Jaeger"]

    client --> gateway --> api
    api --> worker
    api --> ai
    api --> repo
    worker --> ai
    worker --> repo
    repo --> postgres
    api --> redis
    worker --> redis
    api --> nats
    worker --> kafka
    api -. metrics .-> prom
    worker -. metrics .-> prom
    api -. traces .-> jaeger
    worker -. traces .-> jaeger
```

## Quick Start

### Prerequisites

- Go 1.23+
- Docker and Docker Compose
- PostgreSQL 16+
- Redis 7+
- Node.js 18+ for the frontend

### Local Development

```bash
git clone https://github.com/Manikeshmk/GOlang.git
cd GOlang
cp .env.example .env
go mod download
make install-deps
make docker-up
make run
```

The API runs at:

```text
http://localhost:8080
```

### Docker Workflow

```bash
make docker-build
make docker-up
make docker-logs
make docker-down
```

## API Endpoints

### Authentication

| Method | Endpoint | Description |
| --- | --- | --- |
| `POST` | `/auth/register` | Register a new user |
| `POST` | `/auth/login` | Login and receive a JWT |

### Meetings

| Method | Endpoint | Description |
| --- | --- | --- |
| `POST` | `/meetings` | Create a meeting |
| `GET` | `/meetings` | List user meetings |
| `GET` | `/meetings/{id}` | Get meeting details |
| `POST` | `/meetings/{id}/end` | End a meeting |

### Analysis

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/meetings/{meetingId}/summary` | Get generated summary |
| `GET` | `/meetings/{meetingId}/tasks` | Get extracted action items |
| `POST` | `/meetings/{meetingId}/tasks` | Create a task manually |
| `GET` | `/meetings/{meetingId}/conflicts` | Get detected conflicts |
| `GET` | `/meetings/{meetingId}/decisions` | Get meeting decisions |

## Project Structure

```text
.
|-- cmd/
|   |-- api/                 # API server entry point
|   `-- worker/              # Background workers
|-- internal/
|   |-- ai/                  # AI/ML services
|   |-- audio/               # Audio processing
|   |-- config/              # Configuration
|   |-- domain/              # Domain models
|   |-- handler/             # HTTP handlers
|   |-- logger/              # Logging
|   |-- middleware/          # HTTP middleware
|   |-- metrics/             # Prometheus metrics
|   |-- repository/          # Data access
|   |-- service/             # Business logic
|   `-- streaming/           # WebSocket/gRPC streaming
|-- api/proto/               # Protocol Buffer definitions
|-- db/migrations/           # Database migrations
|-- deployments/
|   |-- docker/              # Docker files and compose
|   `-- kubernetes/          # Kubernetes manifests
|-- web/frontend/            # Next.js frontend
|-- tests/                   # Unit and integration tests
|-- scripts/                 # Helper scripts
`-- docs/                    # Project documentation
```

## Docker Services

| Service | Purpose | Default Port |
| --- | --- | --- |
| API | Go application server | `8080` |
| PostgreSQL | Primary database | `5432` |
| Redis | Cache and sessions | `6379` |
| NATS | Lightweight messaging | `4222` |
| Kafka | Event streaming | `9092` |
| ZooKeeper | Kafka coordination | `2181` |
| Ollama | Local LLM runtime | `11434` |
| Prometheus | Metrics collection | `9090` |
| Jaeger | Distributed tracing | `16686` |

## Database Schema

```sql
users              -- User accounts
meetings           -- Meeting sessions
transcripts        -- Speech-to-text output
speakers           -- Unique meeting participants
summaries          -- AI-generated summaries
tasks              -- Extracted action items
decisions          -- Decision records
conflicts          -- Detected disagreements
confusions         -- Confusion signals
repeated_topics    -- Topic clustering results
```

## Testing

```bash
make test
make test-unit
make test-integration
make coverage
go test -race ./...
go test -bench=. ./...
```

## Monitoring

| Tool | URL | Use |
| --- | --- | --- |
| Prometheus | `http://localhost:9090` | Metrics, latency, error rates, runtime stats |
| Jaeger | `http://localhost:16686` | Distributed traces and request flow |
| API Health | `http://localhost:8080/health` | Service readiness |

```bash
curl http://localhost:8080/health
```

## Security

- JWT token-based authentication
- Role-based access control
- Input validation and sanitization
- Encrypted database connections
- CORS protection
- Rate limiting middleware
- Non-root Docker containers
- Kubernetes security policies
- Environment-based secret management

```bash
cp .env.example .env

kubectl create secret generic summarizer-secrets \
  --from-literal=db_password=secure_password \
  --from-literal=jwt_secret=secure_secret
```

## Deployment

### Local Binary

```bash
make build
./bin/summarizer-api
```

### Docker

```bash
make docker-build
make docker-up
```

### Kubernetes

```bash
kubectl create namespace meeting-summarizer
kubectl apply -f deployments/kubernetes/
kubectl get pods -n meeting-summarizer
```

### Production Checklist

- [ ] Rotate `JWT_SECRET`
- [ ] Change database credentials
- [ ] Enable HTTPS/TLS
- [ ] Configure production domain name
- [ ] Set up automated backups
- [ ] Enable centralized log aggregation
- [ ] Configure alert thresholds
- [ ] Review RBAC policies
- [ ] Test disaster recovery

## Development Commands

| Command | Description |
| --- | --- |
| `make fmt` | Format Go code |
| `make lint` | Run linter |
| `make vet` | Run Go vet |
| `make security-check` | Run security scan |
| `make build` | Build API binary |
| `make dev` | Run development server with auto-reload |
| `make docs` | Generate API docs |
| `make docker-logs` | Tail API container logs |

## Environment Variables

```env
SERVER_PORT=8080
ENVIRONMENT=production

DB_HOST=localhost
DB_NAME=meeting_summarizer
DB_USER=postgres
DB_PASSWORD=secure_password

REDIS_HOST=localhost
JWT_SECRET=your-secret-key

WHISPER_MODEL=base
OLLAMA_URL=http://localhost:11434
OPENAI_API_KEY=sk-...

NATS_URL=nats://localhost:4222
KAFKA_URL=localhost:9092
```

See [.env.example](.env.example) for the complete configuration.

## Documentation

- [API Documentation](docs/API.md)
- [Architecture Guide](docs/ARCHITECTURE.md)
- [Deployment Guide](docs/DEPLOYMENT.md)
- [Development Guide](docs/DEVELOPMENT.md)
- [Contributing Guidelines](CONTRIBUTING.md)

## Troubleshooting

### Database Connection Failed

```bash
docker ps | grep postgres
```

Verify the database host, port, username, password, and database name in `.env`.

### Port Already in Use

```bash
lsof -i :8080
kill -9 <PID>
```

### Docker Issues

```bash
docker-compose -f deployments/docker/docker-compose.yml down -v
make docker-clean
make docker-build
```

## Contributing

Contributions are welcome. Please read [CONTRIBUTING.md](CONTRIBUTING.md), open an issue for larger changes, and keep pull requests focused.

## Acknowledgments

Built with the Go ecosystem, Gin, PostgreSQL, Redis, NATS, Kafka, Docker, Kubernetes, OpenAI Whisper, Ollama, Prometheus, Jaeger, Next.js, React, TypeScript, and Tailwind CSS.

## License

MIT License. See [LICENSE](LICENSE) for details.

---

<p align="center">
  <b>Latest Version:</b> 1.0.0 &nbsp;|&nbsp;
  <b>Last Updated:</b> 2026-05-09 &nbsp;|&nbsp;
  <b>Status:</b> Production Ready
</p>
