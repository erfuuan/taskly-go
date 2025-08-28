
# Taskly

Taskly is a lightweight task management application built with **Golang** and **PostgreSQL**.  
It provides a **CLI for task management** and a **REST API** for health checks.

---

## Features

- Add tasks via CLI with title and due date.
- Check tasks due in the next 30 minutes via CLI.
- REST API endpoint for health check.
- PostgreSQL database for task storage.
- Dockerized setup with Go multi-stage build.
- Concurrent REST API server and task runner.
- Configuration via `.env` file.

---

## Prerequisites

- Docker & Docker Compose
- Go 1.21+ (for local development, optional if using Docker)
- Git

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/your-username/taskly-go.git
cd taskly-go
````

2. Create a `.env` file in the root directory:

```env
DB_HOST=taskly-postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=taskly
SERVER_HOST=0.0.0.0
SERVER_PORT=3000
TASK_CHECK_INTERVAL=30s
```

---

## Run with Docker

```bash
docker-compose up --build -d
```

This will start:

* `taskly-postgres`: PostgreSQL database
* `taskly-app-go`: Go application with REST API and task runner

---

## CLI Usage

### Add a Task

```bash
docker exec taskly-app-go ./taskly-cli add --title="Do something" --duedate="2025-11-01T11:23:00"
```

### Run Tasks Manually

```bash
docker exec taskly-app-go ./taskly-cli run
```

### Output Example

```
⏰ Task 1: Do something
```

### CLI Help

```bash
docker exec taskly-app-go ./taskly-cli --help
```

---

# REST API

## Health Check Endpoint

```
curl -X GET http://localhost:3000/api/v1/health
```

### Response Example

```json
{
    "status": "ok",
    "timestamp": "2025-08-28T14:14:45.664Z"
}
```

---

## Notes

* The `tasks` table is automatically created on app start if it does not exist.
* REST API server and task runner run concurrently in the same container.
* Modify `.env` for custom ports or database credentials.
