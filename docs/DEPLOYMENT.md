# Deployment Guide

## Local Development Deployment

### Prerequisites

```bash
Go 1.23+
Docker Desktop
PostgreSQL 16
Redis 7
```

### Quick Start

```bash
# Clone repository
git clone https://github.com/Manikeshmk/GOlang.git
cd GOlang

# Setup environment
cp .env.example .env

# Start services
docker-compose -f deployments/docker/docker-compose.yml up -d

# Run migrations and start API
make run
```

Visit `http://localhost:8080/health`

## Docker Deployment

### Build Image

```bash
docker build -f deployments/docker/Dockerfile -t silent-meeting-summarizer:latest .
```

### Run Container

```bash
docker run -p 8080:8080 \
  -e DB_HOST=postgres \
  -e DB_NAME=meeting_summarizer \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres \
  --network summarizer-network \
  silent-meeting-summarizer:latest
```

### Docker Compose

```bash
# Development environment
docker-compose -f deployments/docker/docker-compose.yml up

# Production environment
docker-compose -f deployments/docker/docker-compose.prod.yml up
```

## Kubernetes Deployment

### Prerequisites

```bash
kubectl 1.24+
Kubernetes 1.24+
4GB RAM minimum
```

### Installation Steps

```bash
# 1. Create namespace
kubectl create namespace meeting-summarizer

# 2. Create secrets
kubectl create secret generic summarizer-secrets \
  --from-literal=db_password=secure_password \
  --from-literal=jwt_secret=secure_secret \
  -n meeting-summarizer

# 3. Apply infrastructure
kubectl apply -f deployments/kubernetes/infrastructure.yml \
  -n meeting-summarizer

# 4. Apply API deployment
kubectl apply -f deployments/kubernetes/api-deployment.yml \
  -n meeting-summarizer

# 5. Apply monitoring
kubectl apply -f deployments/kubernetes/monitoring.yml \
  -n meeting-summarizer

# 6. Verify deployment
kubectl get pods -n meeting-summarizer
kubectl get services -n meeting-summarizer
```

### Access Services

```bash
# Port forward API
kubectl port-forward -n meeting-summarizer \
  svc/summarizer-api-service 8080:80

# Port forward Prometheus
kubectl port-forward -n meeting-summarizer \
  svc/prometheus 9090:9090

# Port forward PostgreSQL
kubectl port-forward -n meeting-summarizer \
  svc/postgres-service 5432:5432
```

### Monitoring Deployment

```bash
# Check pod status
kubectl describe pod <pod-name> -n meeting-summarizer

# View logs
kubectl logs <pod-name> -n meeting-summarizer
kubectl logs -f <pod-name> -n meeting-summarizer

# Check events
kubectl get events -n meeting-summarizer

# Resource usage
kubectl top nodes
kubectl top pods -n meeting-summarizer
```

### Scaling

```bash
# Manual scaling
kubectl scale deployment summarizer-api --replicas=5 \
  -n meeting-summarizer

# Auto-scaling
kubectl get hpa -n meeting-summarizer

# Check current replicas
kubectl get deployment summarizer-api -n meeting-summarizer
```

## AWS Deployment

### ECS Deployment

```bash
# Create ECR repository
aws ecr create-repository --repository-name silent-meeting-summarizer

# Push image
docker tag silent-meeting-summarizer:latest \
  123456789.dkr.ecr.us-east-1.amazonaws.com/silent-meeting-summarizer:latest

docker push 123456789.dkr.ecr.us-east-1.amazonaws.com/silent-meeting-summarizer:latest

# Deploy to ECS (use task definition)
aws ecs create-service --cluster my-cluster \
  --service-name summarizer-api \
  --task-definition summarizer-api \
  --desired-count 3
```

### RDS Database

```bash
# Create RDS instance
aws rds create-db-instance \
  --db-instance-identifier meeting-summarizer-db \
  --db-instance-class db.t3.micro \
  --engine postgres \
  --master-username postgres \
  --master-user-password <password>
```

## Azure Deployment

### App Service Deployment

```bash
# Create resource group
az group create -n summarizer-rg -l eastus

# Create App Service Plan
az appservice plan create -n summarizer-plan \
  -g summarizer-rg --sku B1 --is-linux

# Deploy from Docker
az webapp create -g summarizer-rg \
  -p summarizer-plan -n summarizer-api \
  -i silent-meeting-summarizer:latest
```

### Cosmos DB

```bash
az cosmosdb create \
  -g summarizer-rg \
  -n summarizer-cosmos \
  --kind MongoDB
```

## GCP Deployment

### Cloud Run

```bash
# Build and push to Container Registry
gcloud builds submit --tag gcr.io/PROJECT_ID/silent-meeting-summarizer

# Deploy to Cloud Run
gcloud run deploy summarizer-api \
  --image gcr.io/PROJECT_ID/silent-meeting-summarizer \
  --platform managed \
  --region us-central1
```

### Cloud SQL

```bash
# Create Cloud SQL instance
gcloud sql instances create meeting-summarizer \
  --database-version POSTGRES_16 \
  --tier db-f1-micro
```

## Configuration Management

### Environment Variables

```env
# .env production file
SERVER_PORT=8080
ENVIRONMENT=production

DB_HOST=postgres.example.com
DB_USER=postgres
DB_PASSWORD=<strong_password>

REDIS_HOST=redis.example.com
REDIS_PASSWORD=<strong_password>

JWT_SECRET=<strong_secret>

OPENAI_API_KEY=sk-...
OLLAMA_URL=http://ollama:11434

NATS_URL=nats://nats:4222
KAFKA_URL=kafka:9092

JAEGER_URL=http://jaeger:14268/api/traces
```

### Kubernetes ConfigMap

```bash
kubectl create configmap summarizer-config \
  --from-literal=db_name=meeting_summarizer \
  --from-literal=environment=production \
  -n meeting-summarizer
```

## Database Migrations

### PostgreSQL Setup

```bash
# Connect to database
psql -h localhost -U postgres -d meeting_summarizer

# Run migrations (automatic on startup)
# Or manual:
\i db/migrations/001_initial.sql
```

### Backup & Recovery

```bash
# Backup database
pg_dump -h localhost -U postgres meeting_summarizer > backup.sql

# Restore database
psql -h localhost -U postgres meeting_summarizer < backup.sql

# Automated backups (AWS RDS)
aws rds create-db-snapshot \
  --db-instance-identifier meeting-summarizer-db \
  --db-snapshot-identifier backup-$(date +%Y%m%d)
```

## Health Checks & Monitoring

### Health Endpoints

```bash
# API health
curl http://localhost:8080/health

# Prometheus metrics
curl http://localhost:9090/api/v1/targets
```

### Alerts Setup

Configure Prometheus alerts:

```yaml
groups:
  - name: api
    rules:
      - alert: HighErrorRate
        expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.05
        for: 5m
```

## SSL/TLS Configuration

### Self-Signed Certificate

```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
```

### Let's Encrypt

```bash
# Using cert-manager in Kubernetes
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.0/cert-manager.yaml
```

## CDN and Caching

### CloudFlare Configuration

- Enable caching for static assets
- Cache API responses (if applicable)
- Enable compression
- Set security headers

### Edge Caching

- Cache static assets at CDN edge
- Short TTL for API responses
- Longer TTL for user assets

## Disaster Recovery

### Backup Strategy

- Daily database backups
- Backup replication to secondary region
- Point-in-time recovery capability
- Test restores monthly

### Recovery Procedures

```bash
# Restore from backup
pg_restore -d meeting_summarizer backup.dump

# Verify restoration
SELECT COUNT(*) FROM meetings;
```

### RTO/RPO Targets

- RTO (Recovery Time Objective): 1 hour
- RPO (Recovery Point Objective): 15 minutes

## Production Checklist

- [ ] Update all passwords and secrets
- [ ] Enable SSL/TLS
- [ ] Configure automated backups
- [ ] Set up monitoring and alerts
- [ ] Implement log aggregation
- [ ] Configure CDN
- [ ] Test disaster recovery
- [ ] Document runbooks
- [ ] Set up on-call rotation
- [ ] Review security policies
- [ ] Load test the system
- [ ] Plan capacity

---

For more information, see [README.md](../README.md)
