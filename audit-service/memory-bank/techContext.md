<!-- techContext.md -->

# Technology Context: Audit Service

## 1. Technology Stack

### 1.1 Core Technologies
- **Language**: Go 1.21+
- **HTTP Framework**: Gin v1.9+
- **Logging**: Uber Zap v1.26+
- **Configuration**: Viper v1.17+
- **Testing**: Testify v1.8+
- **Mocking**: Mockery v2.36+

### 1.2 Key Dependencies
```go
// go.mod key dependencies
require (
    github.com/gin-gonic/gin v1.9.1
    github.com/golang-jwt/jwt/v5 v5.0.0
    github.com/patrickmn/go-cache v2.1.0+incompatible
    github.com/spf13/viper v1.17.0
    github.com/stretchr/testify v1.8.4
    github.com/swaggo/gin-swagger v1.6.0
    github.com/swaggo/swag v1.16.2
    go.uber.org/zap v1.26.0
)
```

### 1.3 Development Tools
- **Docker**: v24.0+
- **Docker Compose**: v2.20+
- **Make**: GNU Make 4.3+
- **golangci-lint**: v1.54+
- **swag**: CLI for OpenAPI generation

## 2. Development Setup

### 2.1 Prerequisites
```bash
# Install Go
brew install go  # macOS
# or download from https://golang.org/dl/

# Install development tools
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/vektra/mockery/v2@latest

# Docker (for local development)
# Install Docker Desktop from https://www.docker.com/products/docker-desktop/
```

### 2.2 Project Structure
```
audit-service/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go        # Configuration management
│   ├── domain/
│   │   ├── audit.go         # Domain models
│   │   └── errors.go        # Domain errors
│   ├── handlers/
│   │   └── audit_handler.go # HTTP handlers
│   ├── middleware/
│   │   ├── auth.go          # Authentication middleware
│   │   ├── logger.go        # Logging middleware
│   │   └── request_id.go    # Request ID middleware
│   ├── repository/
│   │   ├── audit_repository.go    # Data access
│   │   └── supabase_client.go     # Supabase REST client
│   └── service/
│       └── audit_service.go        # Business logic
├── pkg/
│   ├── cache/
│   │   └── token_cache.go          # Token caching
│   └── jwt/
│       └── validator.go            # JWT validation
├── api/
│   └── openapi.yaml               # OpenAPI specification
├── docs/                          # Generated Swagger docs
├── scripts/
│   └── generate_docs.sh          # Documentation generator
├── tests/
│   └── integration/              # Integration tests
├── .env.example                  # Environment template
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

### 2.3 Environment Configuration
```bash
# .env.example
PORT=4006
LOG_LEVEL=info

# Supabase Configuration
SUPABASE_URL=http://localhost:54321
SUPABASE_ANON_KEY=your-anon-key
SUPABASE_SERVICE_ROLE_KEY=your-service-role-key
SUPABASE_JWT_SECRET=your-jwt-secret

# HTTP Client Configuration
HTTP_TIMEOUT=30s
HTTP_MAX_IDLE_CONNS=100
HTTP_MAX_CONNS_PER_HOST=10

# Cache Configuration
CACHE_JWT_TTL=5m
CACHE_SHARE_TOKEN_TTL=1m
```

## 3. Build & Run

### 3.1 Local Development
```bash
# Clone and setup
cd audit-service
cp .env.example .env
# Edit .env with your values

# Install dependencies
go mod download

# Generate mocks
make generate-mocks

# Generate OpenAPI docs
make docs

# Run locally
make run

# Or with hot reload
air  # requires: go install github.com/cosmtrek/air@latest
```

### 3.2 Docker Development
```bash
# Build Docker image
make docker-build

# Run with docker-compose
docker-compose up

# Run tests in Docker
docker-compose run --rm audit-service make test
```

### 3.3 Makefile Commands
```makefile
# Common commands
make build          # Build binary
make run           # Run locally
make test          # Run unit tests
make test-coverage # Run tests with coverage
make lint          # Run linter
make docs          # Generate OpenAPI docs
make docker-build  # Build Docker image
make docker-run    # Run in Docker
make clean         # Clean build artifacts
```

## 4. Technical Constraints

### 4.1 Performance Requirements
- Response time: < 200ms (p95)
- Memory usage: < 100MB under normal load
- CPU usage: < 10% for 100 req/s
- Startup time: < 2 seconds

### 4.2 Security Constraints
- All endpoints require authentication
- JWT validation with RS256 algorithm
- Token expiry validation
- Rate limiting per IP/user

### 4.3 Operational Constraints
- Graceful shutdown handling
- Health check endpoint
- Structured JSON logging
- Request ID tracing

## 5. External Dependencies

### 5.1 Supabase Integration
- **REST API**: For database queries
- **Auth**: JWT token validation
- **Tables**: audit_logs, sessions, session_shares

### 5.2 Required Endpoints
```
# Supabase REST endpoints used
GET /rest/v1/audit_logs?session_id=eq.{id}&order=timestamp.desc
GET /rest/v1/sessions?id=eq.{id}
GET /rest/v1/session_shares?token=eq.{token}
```

## 6. Monitoring & Observability

### 6.1 Logging
- Structured JSON logs with Zap
- Log levels: debug, info, warn, error
- Request/response logging
- Performance metrics in logs

### 6.2 Health Checks
```go
// GET /health
{
    "status": "healthy",
    "version": "1.0.0",
    "uptime": "2h30m",
    "checks": {
        "supabase": "ok",
        "cache": "ok"
    }
}
```

### 6.3 Metrics (Future)
- Prometheus metrics endpoint
- Request rate, latency, errors
- Cache hit/miss ratios
- Supabase API latency

## 7. Testing Strategy

### 7.1 Unit Tests
- Minimum 80% code coverage
- Mock all external dependencies
- Table-driven test approach
- Parallel test execution

### 7.2 Integration Tests
- Test against local Supabase
- End-to-end API tests
- Authentication flow tests
- Error scenario coverage

### 7.3 Load Testing
```bash
# Using k6 (future)
k6 run --vus 100 --duration 30s load-test.js
```

--- 