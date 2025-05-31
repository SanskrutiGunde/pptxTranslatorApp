<!-- systemPatterns.md -->

# System Patterns: Audit Service

## 1. Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                    HTTP Requests                             │
│                  (GET /sessions/{id}/history)                │
└───────────────────────┬─────────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────────┐
│                    Gin Router                                │
│  ┌─────────────────────────────────────────────────────┐   │
│  │                Middleware Stack                      │   │
│  │  • Request ID Generator                              │   │
│  │  • Zap Logger (structured)                          │   │
│  │  • Auth Middleware (JWT/Share Token)                │   │
│  │  • Error Handler                                    │   │
│  └─────────────────────────────────────────────────────┘   │
└───────────────────────┬─────────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────────┐
│                    Handlers Layer                            │
│              (AuditHandler.GetHistory)                       │
└───────────────────────┬─────────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────────┐
│                    Service Layer                             │
│              (AuditService.GetAuditLogs)                     │
│  • Business logic                                            │
│  • Permission validation                                     │
│  • Response formatting                                       │
└───────────────────────┬─────────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────────┐
│                  Repository Layer                            │
│           (AuditRepository.FindBySessionID)                  │
│  • Supabase REST API calls                                   │
│  • HTTP connection pooling                                   │
│  • Response parsing                                          │
└───────────────────────┬─────────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────────┐
│                 External Services                            │
│  ┌─────────────────┐  ┌──────────────────┐                 │
│  │  Token Cache    │  │  Supabase REST   │                 │
│  │  (In-Memory)    │  │  API             │                 │
│  └─────────────────┘  └──────────────────┘                 │
└─────────────────────────────────────────────────────────────┘
```

## 2. Design Patterns

### 2.1 Domain-Driven Design (DDD)
```go
// Clear separation of concerns
internal/
├── domain/      // Business entities & rules
├── handlers/    // HTTP layer
├── service/     // Business logic
└── repository/  // Data access
```

### 2.2 Dependency Injection
```go
// Constructor injection for testability
type AuditHandler struct {
    service Service
    logger  *zap.Logger
}

func NewAuditHandler(service Service, logger *zap.Logger) *AuditHandler {
    return &AuditHandler{
        service: service,
        logger:  logger,
    }
}
```

### 2.3 Interface Segregation
```go
// Small, focused interfaces
type AuditService interface {
    GetAuditLogs(ctx context.Context, sessionID string, limit, offset int) (*AuditResponse, error)
}

type AuditRepository interface {
    FindBySessionID(ctx context.Context, sessionID string, limit, offset int) ([]AuditEntry, int, error)
}
```

### 2.4 Repository Pattern
- Abstracts data access behind interfaces
- Enables easy mocking for tests
- Centralizes Supabase REST API logic

### 2.5 Middleware Chain Pattern
```go
router.Use(
    middleware.RequestID(),
    middleware.Logger(logger),
    middleware.ErrorHandler(),
)

protected.Use(middleware.Auth(tokenValidator))
```

## 3. Authentication Flow

```
┌──────────┐     ┌──────────────┐     ┌────────────┐     ┌──────────┐
│  Client  │────▶│   Auth MW    │────▶│Token Cache │────▶│ Validate │
└──────────┘     └──────────────┘     └────────────┘     └──────────┘
                         │                    │ Miss              │
                         │                    └──────────────────▶│
                         │                                        ▼
                         │                              ┌──────────────┐
                         │                              │  Supabase    │
                         │                              │  Validation  │
                         │                              └──────────────┘
                         ▼
                 ┌──────────────┐
                 │   Handler    │
                 └──────────────┘
```

## 4. Caching Strategy

### Token Cache Design
```go
type TokenCache struct {
    cache *cache.Cache  // go-cache with TTL
}

// Cache JWT tokens for 5 minutes
// Cache share tokens for 1 minute
// Reduce auth overhead by 90%+
```

### Cache Key Patterns
- JWT: `jwt:{token_hash}`
- Share: `share:{token}:{sessionID}`

## 5. Error Handling Patterns

### Structured Errors
```go
type APIError struct {
    Code    string `json:"error"`
    Message string `json:"message"`
    Status  int    `json:"-"`
}

// Consistent error responses
var (
    ErrUnauthorized = &APIError{
        Code:    "unauthorized",
        Message: "Invalid or missing authentication",
        Status:  401,
    }
    ErrForbidden = &APIError{
        Code:    "forbidden", 
        Message: "Access denied to this resource",
        Status:  403,
    }
    ErrNotFound = &APIError{
        Code:    "not_found",
        Message: "Session not found",
        Status:  404,
    }
)
```

## 6. Logging Patterns

### Structured Logging with Context
```go
logger.Info("audit logs retrieved",
    zap.String("request_id", requestID),
    zap.String("session_id", sessionID),
    zap.Int("count", len(entries)),
    zap.Duration("duration", time.Since(start)),
)
```

### Request Tracing
- Generate UUID for each request
- Pass through all layers via context
- Include in all log entries

## 7. Configuration Management

### Environment-Based Config
```go
type Config struct {
    Port              string
    SupabaseURL       string
    SupabaseAnonKey   string
    SupabaseJWTSecret string
    LogLevel          string
    
    // HTTP Client settings
    HTTPTimeout           time.Duration
    HTTPMaxIdleConns      int
    HTTPMaxConnsPerHost   int
}
```

### Viper Integration
- Load from environment variables
- Support for config files
- Default values for development

## 8. Testing Patterns

### Table-Driven Tests
```go
func TestGetAuditLogs(t *testing.T) {
    tests := []struct {
        name      string
        sessionID string
        mockData  []AuditEntry
        wantErr   error
    }{
        // Test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### Mock Interfaces
```go
type MockRepository struct {
    mock.Mock
}

func (m *MockRepository) FindBySessionID(ctx context.Context, sessionID string, limit, offset int) ([]AuditEntry, int, error) {
    args := m.Called(ctx, sessionID, limit, offset)
    return args.Get(0).([]AuditEntry), args.Int(1), args.Error(2)
}
```

## 9. Performance Patterns

### Connection Pooling
```go
httpClient := &http.Client{
    Timeout: config.HTTPTimeout,
    Transport: &http.Transport{
        MaxIdleConns:        config.HTTPMaxIdleConns,
        MaxIdleConnsPerHost: config.HTTPMaxConnsPerHost,
        IdleConnTimeout:     90 * time.Second,
    },
}
```

### Context Deadlines
```go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
```

--- 