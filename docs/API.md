# API Documentation

## Authentication

All protected endpoints require a Bearer token in the Authorization header:

```
Authorization: Bearer <jwt_token>
```

### Register User

**POST** `/auth/register`

Register a new user account.

**Request Body:**

```json
{
  "email": "user@example.com",
  "name": "John Doe",
  "password": "securepassword123"
}
```

**Response (201 Created):**

```json
{
  "id": "uuid",
  "email": "user@example.com",
  "name": "John Doe"
}
```

### Login

**POST** `/auth/login`

Authenticate and receive JWT token.

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

**Response (200 OK):**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

## Meetings

### Create Meeting

**POST** `/meetings`

Start a new meeting session.

**Headers:**

```
Authorization: Bearer <token>
```

**Request Body:**

```json
{
  "title": "Q2 Planning",
  "description": "Strategic planning for Q2"
}
```

**Response (201 Created):**

```json
{
  "id": "meeting_uuid",
  "title": "Q2 Planning",
  "status": "active",
  "startTime": "2026-05-09T10:00:00Z",
  "participantCount": 0,
  "duration": 0
}
```

### List Meetings

**GET** `/meetings`

Retrieve all meetings for authenticated user.

**Headers:**

```
Authorization: Bearer <token>
```

**Response (200 OK):**

```json
[
  {
    "id": "meeting_uuid",
    "title": "Q2 Planning",
    "status": "active",
    "startTime": "2026-05-09T10:00:00Z",
    "endTime": null,
    "participantCount": 5,
    "duration": 3600
  }
]
```

### Get Meeting

**GET** `/meetings/{id}`

Retrieve details of a specific meeting.

**Headers:**

```
Authorization: Bearer <token>
```

**Response (200 OK):**

```json
{
  "id": "meeting_uuid",
  "title": "Q2 Planning",
  "description": "Strategic planning for Q2",
  "status": "completed",
  "startTime": "2026-05-09T10:00:00Z",
  "endTime": "2026-05-09T11:00:00Z",
  "participantCount": 5,
  "duration": 3600
}
```

### End Meeting

**POST** `/meetings/{id}/end`

Mark a meeting as completed.

**Headers:**

```
Authorization: Bearer <token>
```

**Response (200 OK):**

```json
{
  "message": "meeting ended"
}
```

## Tasks

### Get Meeting Tasks

**GET** `/meetings/{meetingId}/tasks`

Retrieve all extracted tasks from a meeting.

**Headers:**

```
Authorization: Bearer <token>
```

**Response (200 OK):**

```json
[
  {
    "id": "task_uuid",
    "meetingId": "meeting_uuid",
    "title": "Deploy staging build",
    "status": "pending",
    "priority": "high",
    "owner": "John Doe",
    "deadline": "2026-05-15T00:00:00Z"
  }
]
```

## Error Responses

### 400 Bad Request

```json
{
  "error": "invalid request"
}
```

### 401 Unauthorized

```json
{
  "error": "invalid token"
}
```

### 404 Not Found

```json
{
  "error": "resource not found"
}
```

### 500 Internal Server Error

```json
{
  "error": "internal server error"
}
```

## Rate Limiting

Rate limits are applied per user:

- Authentication endpoints: 10 requests per minute
- Regular endpoints: 100 requests per minute

## Pagination

List endpoints support pagination:

- `?limit=50` - Number of results (default: 50, max: 100)
- `?offset=0` - Number of results to skip (default: 0)

## WebSocket Connections

Live updates via WebSocket:

```
ws://localhost:8080/ws/meetings/{meetingId}
```

**Messages:**

```json
{
  "type": "transcript",
  "data": {
    "speaker": "John Doe",
    "text": "Let's discuss the roadmap",
    "timestamp": "2026-05-09T10:05:00Z"
  }
}
```

---

For more examples and integration guides, see the [main README](../README.md).
