#!/bin/bash

# Deploy to Kubernetes

NAMESPACE=${1:-default}
IMAGE_TAG=${2:-latest}

echo "Deploying to Kubernetes namespace: $NAMESPACE"

# Create namespace if doesn't exist
kubectl create namespace $NAMESPACE 2>/dev/null || true

echo "Creating secrets..."
kubectl create secret generic summarizer-secrets \
  --from-literal=db_password=postgres \
  --from-literal=jwt_secret=your-secret-key \
  -n $NAMESPACE \
  2>/dev/null || kubectl set env secret/summarizer-secrets db_password=postgres jwt_secret=your-secret-key -n $NAMESPACE

echo "Applying infrastructure..."
kubectl apply -f deployments/kubernetes/infrastructure.yml -n $NAMESPACE

echo "Applying API deployment..."
kubectl apply -f deployments/kubernetes/api-deployment.yml -n $NAMESPACE

echo "Applying monitoring..."
kubectl apply -f deployments/kubernetes/monitoring.yml -n $NAMESPACE

echo "✅ Deployment completed"
echo ""
echo "Check status with:"
echo "  kubectl get pods -n $NAMESPACE"
echo "  kubectl get services -n $NAMESPACE"
