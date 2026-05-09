"# Silent Meeting Summarizer

[![CI/CD Pipeline](https://github.com/Manikeshmk/GOlang/workflows/CI%2FCD%20Pipeline/badge.svg)](https://github.com/Manikeshmk/GOlang/actions)
[![codecov](https://codecov.io/gh/Manikeshmk/GOlang/branch/main/graph/badge.svg)](https://codecov.io/gh/Manikeshmk/GOlang)
[![Go Report Card](https://goreportcard.com/badge/github.com/Manikeshmk/GOlang)](https://goreportcard.com/report/github.com/Manikeshmk/GOlang)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A production-grade AI-powered meeting assistant that records audio, performs speaker diarization, generates intelligent summaries, extracts action items, and detects conflicts—all with a focus on local-first processing and optional cloud AI integrations.

## 🎯 Features

### Core Capabilities

- **Real-Time Audio Processing**: Streamed audio ingestion with buffering and backpressure handling
- **Speech-to-Text**: Multi-provider support (Whisper, OpenAI) with automatic fallback
- **Speaker Diarization**: Unique speaker detection and tracking with embeddings
- **AI Summarization**: Multiple summary types (concise, detailed, executive)
- **Task Extraction**: Automatic identification of action items with owners and deadlines
- **Conflict Detection**: Disagreement intensity scoring and unresolved topic identification
- **Confusion Detection**: Recognition of hesitation patterns and clarification requests
- **Decision Analysis**: Confidence scoring for decision finalization
- **Sentiment Analysis**: Real-time emotion tracking throughout meetings
- **Topic Clustering**: Detection of repeated discussion topics and wasted time estimation

### Technical Features

- **Production-Ready**: Structured logging, metrics, tracing, and health checks
- **Scalable**: Horizontal scaling with Kubernetes, worker pools, pub/sub messaging
- **Secure**: JWT authentication, RBAC, input validation, encrypted secrets
- **Observable**: Prometheus metrics, Jaeger tracing, structured logging
- **Dockerized**: Complete Docker and Kubernetes deployment manifests
- **Tested**: Unit, integration, and load testing infrastructure
- **Documented**: API docs, deployment guides, architecture diagrams

## 🚀 Quick Start

### Prerequisites

- Go 1.23+
- Docker & Docker Compose (optional but recommended)
- PostgreSQL 16+
- Redis 7+
- Node.js 18+ (for frontend)

### Local Development

1. **Clone and Setup**

```bash
git clone https://github.com/Manikeshmk/GOlang.git
cd GOlang
cp .env.example .env
```

2. **Install Dependencies**

```bash
go mod download
make install-deps
```

3. **Start Services with Docker Compose**

```bash
make docker-up
```

4. **Run Migrations and Start API**

```bash
make run
```

The API will be available at `http://localhost:8080`

### Using Docker

```bash
# Build image
make docker-build

# Start all services
make docker-up

# View logs
make docker-logs

# Stop services
make docker-down
```

## 📁 Project Structure

```
.
├── cmd/
│   ├── api/                 # API server entry point
│   └── worker/              # Background worker processes
├── internal/
│   ├── ai/                  # AI/ML services
│   ├── audio/               # Audio processing
│   ├── config/              # Configuration management
│   ├── domain/              # Domain models
│   ├── handler/             # HTTP handlers
│   ├── logger/              # Logging setup
│   ├── middleware/          # HTTP middleware
│   ├── metrics/             # Prometheus metrics
│   ├── repository/          # Data access layer
│   ├── service/             # Business logic
│   └── streaming/           # WebSocket/gRPC streaming
├── pkg/
│   └── utils/               # Utility functions
├── api/
│   └── proto/               # Protocol Buffer definitions
├── db/
│   └── migrations/          # Database migrations
├── deployments/
│   ├── docker/              # Docker files and compose
│   └── kubernetes/          # K8s manifests
├── web/
│   └── frontend/            # Next.js frontend
├── tests/
│   ├── unit/                # Unit tests
│   └── integration/         # Integration tests
├── scripts/                 # Helper scripts
├── docs/                    # Documentation
└── Makefile                 # Make commands
```

## 🔧 API Endpoints

### Authentication

- `POST /auth/register` - Register new user
- `POST /auth/login` - Login and get JWT token

### Meetings

- `POST /meetings` - Create new meeting
- `GET /meetings` - List user's meetings
- `GET /meetings/{id}` - Get meeting details
- `POST /meetings/{id}/end` - End meeting

### Tasks

- `GET /meetings/{meetingId}/tasks` - Get extracted tasks
- `POST /meetings/{meetingId}/tasks` - Create task

### Analysis

- `GET /meetings/{meetingId}/summary` - Get meeting summary
- `GET /meetings/{meetingId}/conflicts` - Get detected conflicts
- `GET /meetings/{meetingId}/decisions` - Get decisions made

## 🗄️ Database Schema

```sql
-- Core tables
users              -- User accounts
meetings           -- Meeting sessions
transcripts        -- Speech-to-text results
speakers           -- Unique participants
summaries          -- AI-generated summaries
tasks              -- Extracted action items
decisions          -- Decision records
conflicts          -- Detected disagreements
confusions         -- Confusion signals
repeated_topics    -- Topic clustering results
```

## 🐳 Docker Compose Services

The `docker-compose.yml` includes:

- **API** - Go application server
- **PostgreSQL** - Primary database
- **Redis** - Caching and sessions
- **NATS** - Message broker
- **Kafka** - Event streaming
- **Zookeeper** - Kafka coordination
- **Ollama** - Local LLM support
- **Prometheus** - Metrics collection
- **Jaeger** - Distributed tracing

## ☸️ Kubernetes Deployment

Deploy to Kubernetes:

```bash
# Create namespace
kubectl create namespace meeting-summarizer

# Apply configs and secrets
kubectl apply -f deployments/kubernetes/

# Monitor deployment
kubectl get pods -n meeting-summarizer
kubectl logs -f pod/<pod-name> -n meeting-summarizer
```

## 🧪 Testing

```bash
# Run all tests
make test

# Run unit tests only
make test-unit

# Run integration tests
make test-integration

# Generate coverage report
make coverage

# Run with race detector
go test -race ./...

# Benchmark tests
go test -bench=. ./...
```

## 📊 Monitoring

### Prometheus

Access at `http://localhost:9090`

- CPU, memory, goroutine metrics
- Request latency and error rates
- Custom application metrics

### Jaeger

Access at `http://localhost:16686`

- Distributed tracing
- Request flow visualization
- Performance bottleneck identification

### Health Checks

```bash
curl http://localhost:8080/health
```

## 🔐 Security

### Implemented

- ✅ JWT token-based authentication
- ✅ Role-based access control (RBAC)
- ✅ Input validation and sanitization
- ✅ Encrypted database connections
- ✅ CORS protection
- ✅ Rate limiting middleware
- ✅ Non-root Docker containers
- ✅ K8s security policies

### Secrets Management

```bash
# Create .env with secrets (never commit)
cp .env.example .env

# Or use Kubernetes secrets
kubectl create secret generic summarizer-secrets \
  --from-literal=db_password=secure_password \
  --from-literal=jwt_secret=secure_secret
```

## 📈 Performance

### Optimization Features

- Connection pooling (25 max connections)
- Redis caching layer
- Worker pool pattern for concurrency
- Streaming audio processing
- Incremental summarization
- Query optimization with indexes

### Benchmarks

```bash
make test
# Detailed performance metrics in coverage report
```

## 🚀 Deployment

### Local Deployment

```bash
make build
./bin/summarizer-api
```

### Docker Deployment

```bash
make docker-build
make docker-up
```

### Kubernetes Deployment

```bash
kubectl apply -f deployments/kubernetes/api-deployment.yml
kubectl apply -f deployments/kubernetes/infrastructure.yml
kubectl apply -f deployments/kubernetes/monitoring.yml
```

### Production Checklist

- [ ] Update JWT_SECRET in .env
- [ ] Change database password
- [ ] Enable HTTPS/TLS
- [ ] Configure domain name
- [ ] Set up automated backups
- [ ] Enable log aggregation
- [ ] Configure alert thresholds
- [ ] Review RBAC policies
- [ ] Test disaster recovery

## 📚 Documentation

- [API Documentation](docs/API.md)
- [Architecture Guide](docs/ARCHITECTURE.md)
- [Deployment Guide](docs/DEPLOYMENT.md)
- [Development Guide](docs/DEVELOPMENT.md)
- [Contributing Guidelines](CONTRIBUTING.md)

## 🛠️ Development Commands

```bash
# Code formatting
make fmt

# Linting
make lint

# Security scan
make security-check

# Build binary
make build

# Run development server with auto-reload
make dev

# View Docker logs
make docker-logs

# Generate API docs
make docs
```

## 📝 Environment Variables

Key configuration variables:

```env
# Server
SERVER_PORT=8080
ENVIRONMENT=production

# Database
DB_HOST=localhost
DB_NAME=meeting_summarizer
DB_USER=postgres
DB_PASSWORD=secure_password

# Redis
REDIS_HOST=localhost

# JWT
JWT_SECRET=your-secret-key

# AI/ML
WHISPER_MODEL=base
OLLAMA_URL=http://localhost:11434
OPENAI_API_KEY=sk-...

# Infrastructure
NATS_URL=nats://localhost:4222
KAFKA_URL=localhost:9092
```

See `.env.example` for all variables.

## 🚨 Troubleshooting

### Database Connection Failed

```bash
# Check PostgreSQL is running
docker ps | grep postgres

# Verify credentials in .env
# Check connection string format
```

### Port Already in Use

```bash
# Find process on port 8080
lsof -i :8080

# Kill the process
kill -9 <PID>
```

### Docker Issues

```bash
# Clear Docker cache
docker-compose down -v

# Rebuild from scratch
make docker-clean
make docker-build
```

## 📊 Architecture

```
┌─────────────────────────────────────────┐
│         Client Applications              │
│  (Web, Mobile, Third-party integrations) │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────┴──────────────────────┐
│      API Gateway / Load Balancer         │
│    (Kubernetes Ingress / nginx)          │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────┴──────────────────────┐
│        REST API / WebSocket              │
│    (Gin + Gorilla WebSocket)             │
└──────────────────┬──────────────────────┘
                   │
       ┌───────────┼───────────┐
       │           │           │
┌──────┴────┐  ┌───┴──────┐  ┌┴─────────┐
│  Services │  │  Handler │  │Middleware│
│           │  │           │  │          │
└──────┬────┘  └───┬──────┘  └┬─────────┘
       │           │          │
       └───────────┼──────────┘
                   │
┌──────────────────┴──────────────────────┐
│    Data Access & AI Services            │
│  (Repository, AI, Streaming)            │
└──────────────────┬──────────────────────┘
                   │
       ┌───────────┼───────────────────┐
       │           │                   │
  ┌────┴─────┐ ┌──┴──────┐ ┌─────┴───┐
  │PostgreSQL│ │ Redis   │ │NATS/Msg │
  │Database  │ │ Cache   │ │ Broker  │
  └──────────┘ └─────────┘ └─────────┘
```

## 📄 License

MIT License - see LICENSE file for details

## 🤝 Contributing

Contributions welcome! See CONTRIBUTING.md for guidelines.

## 📧 Contact

For questions or issues, open a GitHub issue or contact the maintainers.

## 🙏 Acknowledgments

Built with:

- Golang ecosystem
- Gin web framework
- PostgreSQL
- Redis
- NATS messaging
- OpenAI Whisper
- Kubernetes

## 📚 Additional Resources

- [Golang Documentation](https://golang.org/doc/)
- [Gin Documentation](https://gin-gonic.com/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [Docker Documentation](https://docs.docker.com/)

---

**Latest Version**: 1.0.0  
**Last Updated**: 2026-05-09  
**Status**: Production Ready ✅"
