# Planning Service

A task planning and distribution service written in Go that helps allocate tasks among developers based on their capacity and task complexity.

## Features

- Task fetching from multiple providers
- SQLite database for task persistence
- Automated task distribution based on developer capacity
- RESTful API endpoints
- Docker support

## Prerequisites

- Go 1.23.6 or higher
- Fiber web framework
- SQLite
- Docker (optional)

## Getting Started

### Local Development

1. Clone the repository:
```bash
git clone <repository-url>
cd planning
```

2. Install dependencies:
```bash
make download-deps
```

3. Run tests:
```bash
make test
```

4. Generate mocks:
```bash
make mocks
```

5. Run the application:
```bash
make run
```

The service will be available at `http://localhost:7878`

## API Endpoints

### UI

The web interface is served at the root URL.


### Distribute Tasks

Distributes tasks among developers based on their capacity and task difficulty.

```
POST /api/tasks/schedule


Request Body:
{
    "developers": [
        {
            "name": "DEV1",
            "capacity": 1,
            "weeklyHours": 45,
            "currentHours": 45
        }
    ]
}

Response:
{
    "schedule": {
        "DEV1": [
            {
                "name": "Task 1",
                "duration": 60,
                "difficulty": 3
            }
        ]
    },
    "totalWeeks": 1
}
```

## Project Structure

```
.
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── clients/          # External service clients
│   ├── models/           # Data models
│   ├── repository/       # Data access layer
│   ├── resource/         # Request/Response structures
│   └── services/         # Business logic
├── └── handler/          # handlers for API endpoints
├── └── router/           # routes for API endpoints  
├── pkg/
│   └── sqlite/          # SQLite database wrapper
├── web/                 # Web interface assets
│   └── static/
│   |    ├── js/         # JavaScript files
│   └── templates/        # HTML templatess
├── Dockerfile           # Docker configuration
├── Makefile            # Build and development commands
└── go.mod              # Go module definition
```

## Web Interface

The web interface provides a user-friendly way to:
- Add and manage developers
- Set developer capacities and working hours
- Distribute tasks automatically
- View task assignments and schedules

The interface is served from the `web/` directory and uses:
- Bootstrap for styling
- jQuery for DOM manipulation and AJAX calls
- Custom JavaScript for developer management and task distribution

## Testing

Run the test suite:

```bash
make test
```

This will generate a coverage report at `coverage.html`
