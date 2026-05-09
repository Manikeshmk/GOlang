#!/bin/bash

# Health check script for monitoring API

API_URL=${1:-"http://localhost:8080"}
MAX_RETRIES=${2:-5}
RETRY_DELAY=${3:-5}

echo "Checking API health at $API_URL..."

for i in $(seq 1 $MAX_RETRIES); do
    echo "Attempt $i/$MAX_RETRIES..."
    
    RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "$API_URL/health")
    
    if [ "$RESPONSE" = "200" ]; then
        echo "✅ API is healthy (HTTP $RESPONSE)"
        exit 0
    else
        echo "❌ API returned HTTP $RESPONSE"
        
        if [ $i -lt $MAX_RETRIES ]; then
            echo "Retrying in ${RETRY_DELAY}s..."
            sleep $RETRY_DELAY
        fi
    fi
done

echo "❌ API health check failed after $MAX_RETRIES attempts"
exit 1
