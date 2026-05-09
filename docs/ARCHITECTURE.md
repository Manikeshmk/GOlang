# Architecture Guide

## System Architecture

The Silent Meeting Summarizer is built with a clean, layered architecture following industry best practices.

### Layers

```
┌─────────────────────────────────────────┐
│          Presentation Layer             │
│    (REST API, WebSocket, HTTP)          │
├─────────────────────────────────────────┤
│          Application Layer              │
│    (Handlers, Middleware, Routes)       │
├─────────────────────────────────────────┤
│          Business Logic Layer           │
│    (Services, AI, Analytics)            │
├─────────────────────────────────────────┤
│          Data Access Layer              │
│    (Repositories, Queries)              │
├─────────────────────────────────────────┤
│          Infrastructure Layer           │
│    (Database, Cache, Messaging)         │
└─────────────────────────────────────────┘
```

### Key Components

#### 1. API Handler Layer (`internal/handler/`)

- REST endpoint implementations
- Request validation
- Response formatting
- Error handling

#### 2. Service Layer (`internal/service/`)

- Business logic
- Orchestration
- Transaction management
- Domain logic

#### 3. AI Services (`internal/ai/`)

- Transcription service (speech-to-text)
- Summarization engine
- Task extraction
- Conflict detection
- Sentiment analysis
- Decision confidence scoring

#### 4. Repository Layer (`internal/repository/`)

- Data access abstraction
- Query builders
- Database operations
- Connection pooling

#### 5. Domain Models (`internal/domain/`)

- Core entities
- Value objects
- Business rules

### Data Flow

```
User Request
    ↓
[Middleware] (Auth, Logging, CORS)
    ↓
[Handler] (Parse request)
    ↓
[Service] (Business logic)
    ↓
[Repository] (Data access)
    ↓
[Database/Cache]
    ↓
[Response]
```

## Concurrency Model

### Worker Pool Pattern

The system uses worker pools for parallel processing:

```golang
workers := make(chan Task, buffSize)
for i := 0; i < numWorkers; i++ {
    go worker(workers)
}

for task := range taskChan {
    workers <- task
}
```

### Channels & Goroutines

- Audio ingestion: Multiple reader goroutines
- Processing: Worker pool with backpressure
- Broadcasting: Fan-out pattern for events

### Context Propagation

All operations use `context.Context` for:

- Cancellation
- Timeout
- Deadline tracking
- Value passing

## Database Schema

### Core Entities

```sql
users
├── id (UUID)
├── email (UNIQUE)
├── name
├── password (hashed)
├── role
└── timestamps

meetings
├── id (UUID)
├── user_id (FK)
├── title
├── status
├── start_time
├── end_time
└── timestamps

speakers
├── id (UUID)
├── meeting_id (FK)
├── name
├── email
├── embedding (vector)
└── stats

transcripts
├── id (UUID)
├── meeting_id (FK)
├── speaker_id (FK)
├── text
├── sentiment
├── timestamps
└── confidence

tasks
├── id (UUID)
├── meeting_id (FK)
├── title
├── owner
├── status
├── priority
└── deadline
```

### Indexes

- `meetings(user_id, status)` - Quick user lookups
- `transcripts(meeting_id, start_time)` - Timeline queries
- `speakers(meeting_id)` - Speaker lookups
- `tasks(meeting_id, status)` - Task filtering

## Caching Strategy

### Redis Cache Layers

```
├── User Sessions (TTL: 24h)
├── Meeting Data (TTL: 1h)
├── Transcript Cache (TTL: 30m)
├── Summarization Cache (TTL: 2h)
└── Feature Toggles (TTL: 5m)
```

### Cache Invalidation

```
Update Meeting → Invalidate:
├── meeting:{id}
├── transcripts:{id}
├── summary:{id}
└── tasks:{id}
```

## Message Queue Architecture

### NATS for Internal Events

```
Events Published:
├── TranscriptReceived
├── SummaryGenerated
├── TaskExtracted
├── ConflictDetected
└── MeetingEnded

Subscribers:
├── Analytics Service
├── Notification Service
├── Search Indexer
└── Analytics Pipeline
```

### Kafka for High-Volume Streaming

```
Topics:
├── audio-stream (partitioned by meeting)
├── transcripts (partitioned by meeting)
├── analysis-results
└── metrics
```

## Observability Stack

### Structured Logging

```golang
logger.Info("meeting started",
    "meeting_id", meetingID,
    "participants", count,
    "duration", timeElapsed,
)
```

### Metrics Collection

```
Prometheus Metrics:
├── http_requests_total (counter)
├── http_request_duration (histogram)
├── db_query_duration (histogram)
├── active_meetings (gauge)
└── transcripts_processed (counter)
```

### Distributed Tracing

```
Jaeger Spans:
├── api_request
│   ├── auth_check
│   ├── database_query
│   ├── ai_processing
│   └── cache_operation
```

## Security Architecture

### Authentication Flow

```
User Input
    ↓
[JWT Verification]
    ↓
[User ID Extraction]
    ↓
[RBAC Check]
    ↓
[Request Proceed]
```

### Authorization Levels

```
Anonymous
├── /auth/register
├── /auth/login
└── /health

User
├── /meetings (own only)
├── /profile
└── /settings

Admin
├── /users
├── /analytics
└── /system
```

## Scaling Considerations

### Horizontal Scaling

```
Load Balancer
├── API Instance 1
├── API Instance 2
├── API Instance 3
└── API Instance N
    ↓
Shared Resources:
├── PostgreSQL Replicas
├── Redis Cluster
├── NATS Cluster
└── Kafka Brokers
```

### Database Scaling

- Read replicas for analytics queries
- Partitioning by meeting_id for large tables
- Connection pooling (25-50 connections)
- Query optimization with indexes

### Cache Warming

- Pre-load user preferences
- Cache popular meetings
- Warm up trending topics
- Regular TTL rotation

## Deployment Architecture

### Development Environment

```
Docker Compose
├── API (localhost:8080)
├── PostgreSQL (localhost:5432)
├── Redis (localhost:6379)
├── Ollama (localhost:11434)
└── Prometheus (localhost:9090)
```

### Production Environment

```
Kubernetes Cluster
├── API Deployment (replicas: 3)
├── StatefulSet PostgreSQL
├── Redis Deployment
├── ConfigMaps & Secrets
└── Services & Ingress
```

## Error Handling

### Error Propagation

```golang
if err != nil {
    logger.Error("operation failed", err,
        "context", context,
    )
    return wrappedErr
}
```

### Graceful Degradation

- Primary provider fails → Use fallback
- Database unavailable → Use cache
- Cache unavailable → Proceed without cache
- All services down → Health check fails

## Performance Targets

- Response time: <200ms (p99)
- Database query time: <50ms (p99)
- Throughput: 1000+ RPS
- Concurrency: 10,000+ connections
- Memory per instance: <512MB
- CPU per instance: <500m

---

For deployment details, see [DEPLOYMENT.md](DEPLOYMENT.md)
