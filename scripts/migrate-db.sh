#!/bin/bash

# Database migration script

DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_NAME=${DB_NAME:-meeting_summarizer}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}

PSQL_CMD="psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME"

echo "Running database migrations..."

# Create tables
$PSQL_CMD <<EOF
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS meetings (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    status VARCHAR(50) DEFAULT 'active',
    duration INTEGER,
    participant_count INTEGER DEFAULT 0,
    user_id VARCHAR(36) NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_meetings_user_id ON meetings(user_id);
CREATE INDEX IF NOT EXISTS idx_meetings_status ON meetings(status);

CREATE TABLE IF NOT EXISTS speakers (
    id VARCHAR(36) PRIMARY KEY,
    meeting_id VARCHAR(36) NOT NULL REFERENCES meetings(id),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    embedding TEXT,
    speak_time INTEGER DEFAULT 0,
    turn_count INTEGER DEFAULT 0,
    interruptions INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_speakers_meeting_id ON speakers(meeting_id);

CREATE TABLE IF NOT EXISTS transcripts (
    id VARCHAR(36) PRIMARY KEY,
    meeting_id VARCHAR(36) NOT NULL REFERENCES meetings(id),
    speaker_id VARCHAR(36) REFERENCES speakers(id),
    speaker_name VARCHAR(255),
    text TEXT NOT NULL,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    confidence DECIMAL(4,2),
    language VARCHAR(10) DEFAULT 'en',
    sentiment VARCHAR(50),
    emotions TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_transcripts_meeting_id ON transcripts(meeting_id);
EOF

echo "✅ Database migrations completed successfully"
