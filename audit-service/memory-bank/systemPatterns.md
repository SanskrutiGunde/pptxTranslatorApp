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

### 8.1 Mock Generation Strategy
```yaml
# .mockery.yaml
with-expecter: true
dir: "mocks"
outpkg: "mocks"
mockname: "Mock{{.InterfaceName}}"
filename: "mock_{{.InterfaceName | snakecase}}.go"
interfaces:
  AuditService:
    config:
      dir: "internal/service/mocks"
  AuditRepository:
    config:
      dir: "internal/repository/mocks"
  TokenValidator:
    config:
      dir: "pkg/jwt/mocks"
```

### 8.2 Unit Test Patterns

#### Table-Driven Tests
```go
func TestAuditService_GetAuditLogs(t *testing.T) {
    tests := []struct {
        name         string
        sessionID    string
        limit        int
        offset       int
        mockSetup    func(*mocks.MockAuditRepository)
        expectedResp *domain.AuditResponse
        expectedErr  error
    }{
        {
            name:      "successful retrieval",
            sessionID: "valid-session-id",
            limit:     10,
            offset:    0,
            mockSetup: func(repo *mocks.MockAuditRepository) {
                repo.EXPECT().FindBySessionID(
                    mock.Anything, "valid-session-id", 10, 0,
                ).Return(mockEntries, 25, nil)
            },
            expectedResp: &domain.AuditResponse{
                TotalCount: 25,
                Items:      mockEntries,
            },
            expectedErr: nil,
        },
        // Additional test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation with proper setup/teardown
        })
    }
}
```

#### Mock Interface Usage
```go
type MockAuditRepository struct {
    mock.Mock
}

func (m *MockAuditRepository) FindBySessionID(
    ctx context.Context, 
    sessionID string, 
    limit, offset int,
) ([]domain.AuditEntry, int, error) {
    args := m.Called(ctx, sessionID, limit, offset)
    return args.Get(0).([]domain.AuditEntry), args.Int(1), args.Error(2)
}
```

### 8.3 HTTP Testing Patterns

#### Handler Testing with httptest
```go
func TestAuditHandler_GetHistory(t *testing.T) {
    gin.SetMode(gin.TestMode)
    
    tests := []struct {
        name           string
        sessionID      string
        queryParams    string
        mockSetup      func(*mocks.MockAuditService)
        expectedStatus int
        expectedBody   string
    }{
        // Test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup mock service
            mockService := mocks.NewMockAuditService(t)
            tt.mockSetup(mockService)
            
            // Create handler and router
            handler := handlers.NewAuditHandler(mockService, logger)
            router := gin.New()
            router.GET("/sessions/:sessionId/history", handler.GetHistory)
            
            // Create request and recorder
            req := httptest.NewRequest("GET", 
                fmt.Sprintf("/sessions/%s/history%s", tt.sessionID, tt.queryParams), 
                nil)
            w := httptest.NewRecorder()
            
            // Execute request
            router.ServeHTTP(w, req)
            
            // Assertions
            assert.Equal(t, tt.expectedStatus, w.Code)
            assert.JSONEq(t, tt.expectedBody, w.Body.String())
        })
    }
}
```

### 8.4 Integration Test Patterns

#### Supabase Integration Setup
```go
func setupTestSupabase(t *testing.T) *repository.SupabaseClient {
    config := &config.Config{
        SupabaseURL:           "http://localhost:54321",
        SupabaseServiceKey:    os.Getenv("TEST_SUPABASE_SERVICE_KEY"),
        HTTPTimeout:           30 * time.Second,
        HTTPMaxIdleConns:      10,
        HTTPMaxConnsPerHost:   2,
    }
    
    client, err := repository.NewSupabaseClient(config, logger)
    require.NoError(t, err)
    
    // Verify connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    err = client.HealthCheck(ctx)
    require.NoError(t, err, "Supabase connection failed")
    
    return client
}
```

#### Complete API Flow Testing
```go
func TestAuditAPI_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration tests in short mode")
    }
    
    // Setup test server with real dependencies
    supabaseClient := setupTestSupabase(t)
    tokenCache := cache.NewTokenCache(5*time.Minute, 1*time.Minute)
    jwtValidator := jwt.NewValidator(testJWTSecret)
    
    repo := repository.NewAuditRepository(supabaseClient, logger)
    service := service.NewAuditService(repo, logger)
    handler := handlers.NewAuditHandler(service, logger)
    
    router := setupRouter(handler, jwtValidator, tokenCache, logger)
    server := httptest.NewServer(router)
    defer server.Close()
    
    tests := []struct {
        name           string
        setupData      func() (sessionID string, token string)
        expectedStatus int
        validateResp   func(t *testing.T, body []byte)
    }{
        // Integration test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### 8.5 Test Utilities and Helpers

#### Test Fixtures
```go
// tests/helpers/fixtures.go
package helpers

func CreateTestAuditEntry(sessionID, userID string) domain.AuditEntry {
    return domain.AuditEntry{
        ID:        uuid.New(),
        SessionID: sessionID,
        UserID:    userID,
        Action:    domain.ActionEdit,
        Timestamp: time.Now(),
        Details:   json.RawMessage(`{"field": "content", "old": "old", "new": "new"}`),
    }
}

func CreateTestJWT(userID string, sessionID string, secret []byte) string {
    claims := jwt.MapClaims{
        "sub": userID,
        "exp": time.Now().Add(time.Hour).Unix(),
        "iat": time.Now().Unix(),
        "aud": "authenticated",
        "iss": "supabase",
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, _ := token.SignedString(secret)
    return tokenString
}
```

### 8.6 Coverage and Quality Patterns

#### Coverage Configuration
```makefile
# Makefile targets for testing
test:
	go test ./... -v

test-coverage:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

test-coverage-report:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -func=coverage.out
	@echo "Coverage threshold: 80%"

test-integration:
	go test ./tests/integration/... -v -tags=integration

test-all: test test-integration
```

#### Quality Gates
```go
// TestCoverageThreshold ensures minimum coverage
func TestCoverageThreshold(t *testing.T) {
    cmd := exec.Command("go", "test", "./...", "-coverprofile=coverage.out")
    err := cmd.Run()
    require.NoError(t, err)
    
    coverage := parseCoverageOutput("coverage.out")
    assert.GreaterOrEqual(t, coverage, 80.0, 
        "Test coverage below threshold: %.2f%% < 80%%", coverage)
}
```

--- 