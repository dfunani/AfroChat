# AfroChat Backend

A high-performance, real-time communication backend built with Go, designed to power the AfroChat Slack clone across web, iOS, and Android platforms.

## 🏗️ Architecture Overview

### Core Technologies
- **Go 1.21+** with **Gin** framework for HTTP API
- **PostgreSQL** for primary data storage
- **Redis** for caching, sessions, and pub/sub
- **WebSocket** for real-time communication
- **Docker** for containerization

### Microservices Structure
```
backend/
├── cmd/                    # Application entry points
│   ├── api/               # Main API server
│   ├── websocket/         # WebSocket service
│   ├── worker/            # Background workers
│   └── migration/         # Database migrations
├── internal/              # Private application code
│   ├── auth/              # Authentication service
│   ├── chat/              # Chat service
│   ├── workspace/         # Workspace service
│   ├── notification/      # Notification service
│   ├── file/              # File service
│   └── search/            # Search service
├── pkg/                   # Public library code
│   ├── database/          # Database connections
│   ├── redis/             # Redis client
│   ├── websocket/         # WebSocket handler
│   └── middleware/        # Middleware functions
├── api/                   # API definitions
│   ├── v1/                # API version 1
│   └── websocket/         # WebSocket endpoints
├── migrations/            # Database migrations
├── configs/               # Configuration files
├── scripts/               # Build and deployment scripts
└── tests/                 # Test files
```

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- PostgreSQL 14+
- Redis 6+
- Docker & Docker Compose
- Make (for build automation)

### Installation

1. **Clone and setup**
```bash
git clone <repository-url>
cd AfroChat/backend
go mod download
```

2. **Environment Configuration**
```bash
cp .env.example .env
# Edit .env with your configuration
```

3. **Database Setup**
```bash
# Start PostgreSQL and Redis with Docker
docker-compose up -d postgres redis

# Run migrations
make migrate-up

# Seed initial data (optional)
make seed
```

4. **Start Development Server**
```bash
# Start API server
make dev

# Start WebSocket server (separate terminal)
make dev-websocket

# Start background workers (separate terminal)
make dev-worker
```

## 🔧 Go Commands Reference

### **Backend Initialization Commands**

#### **1. Initialize Go Module**
```bash
cd backend
go mod init github.com/dfunani/AfroChat/backend
```

#### **2. Install Dependencies**
```bash
go mod tidy
```

#### **3. Create Project Structure**
```bash
mkdir -p cmd/{api,websocket,worker,migration}
mkdir -p internal/{auth,chat,workspace,notification,file,search}
mkdir -p pkg/{database,redis,websocket,middleware}
mkdir -p api/{v1,websocket}
mkdir -p migrations
mkdir -p configs
mkdir -p scripts
mkdir -p tests
```

#### **4. Install Required Go Packages**
```bash
go get github.com/gin-gonic/gin
go get github.com/golang-jwt/jwt/v5
go get github.com/lib/pq
go get github.com/redis/go-redis/v9
go get github.com/gorilla/websocket
go get github.com/google/uuid
go get github.com/golang-migrate/migrate/v4
go get github.com/golang-migrate/migrate/v4/database/postgres
go get github.com/golang-migrate/migrate/v4/source/file
go get github.com/joho/godotenv
go get github.com/sirupsen/logrus
```

#### **5. Set Up Environment**
```bash
cp .env.example .env
# Edit .env with your configuration
```

#### **6. Run Database Migrations**
```bash
# Start PostgreSQL and Redis with Docker
docker-compose up -d postgres redis

# Run migrations
make migrate-up
```

#### **7. Start Development Servers**
```bash
# Start API server
make dev

# Start WebSocket server (separate terminal)
make dev-websocket

# Start background workers (separate terminal)
make dev-worker
```

### **Alternative Manual Commands**

If you don't have the Makefile set up yet, you can run these directly:

#### **Start API Server**
```bash
go run cmd/api/main.go
```

#### **Start WebSocket Server**
```bash
go run cmd/websocket/main.go
```

#### **Start Background Workers**
```bash
go run cmd/worker/main.go
```

#### **Run Tests**
```bash
go test ./...
```

#### **Build for Production**
```bash
go build -o bin/api cmd/api/main.go
go build -o bin/websocket cmd/websocket/main.go
go build -o bin/worker cmd/worker/main.go
```

### **Prerequisites Setup**

Before running the Go commands, make sure you have:

1. **Go 1.21+** installed
2. **PostgreSQL 14+** running
3. **Redis 6+** running
4. **Docker & Docker Compose** (for easy database setup)

## 🔧 Configuration

### Environment Variables
```bash
# Database
DATABASE_URL=postgres://user:password@localhost:5432/afrochat?sslmode=disable
REDIS_URL=redis://localhost:6379

# Server
PORT=8080
WEBSOCKET_PORT=8081
WORKER_PORT=8082

# JWT
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRE_HOURS=24
REFRESH_TOKEN_EXPIRE_DAYS=7

# File Storage
FILE_STORAGE_TYPE=local  # or s3
FILE_UPLOAD_PATH=./uploads
AWS_S3_BUCKET=afrochat-files
AWS_S3_REGION=us-east-1

# Email
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password

# Search
ELASTICSEARCH_URL=http://localhost:9200
```

## 🗄️ Database Schema

### Core Tables Implementation

```go
// User model
type User struct {
    ID          uuid.UUID `json:"id" db:"id"`
    Email       string    `json:"email" db:"email"`
    Username    string    `json:"username" db:"username"`
    DisplayName string    `json:"display_name" db:"display_name"`
    AvatarURL   *string   `json:"avatar_url" db:"avatar_url"`
    Status      string    `json:"status" db:"status"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Workspace model
type Workspace struct {
    ID          uuid.UUID `json:"id" db:"id"`
    Name        string    `json:"name" db:"name"`
    Slug        string    `json:"slug" db:"slug"`
    Description *string   `json:"description" db:"description"`
    OwnerID     uuid.UUID `json:"owner_id" db:"owner_id"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Channel model
type Channel struct {
    ID          uuid.UUID `json:"id" db:"id"`
    Name        string    `json:"name" db:"name"`
    Description *string   `json:"description" db:"description"`
    WorkspaceID uuid.UUID `json:"workspace_id" db:"workspace_id"`
    IsPrivate   bool      `json:"is_private" db:"is_private"`
    CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Message model
type Message struct {
    ID        uuid.UUID  `json:"id" db:"id"`
    Content   string    `json:"content" db:"content"`
    ChannelID uuid.UUID `json:"channel_id" db:"channel_id"`
    UserID    uuid.UUID `json:"user_id" db:"user_id"`
    ThreadID  *uuid.UUID `json:"thread_id" db:"thread_id"`
    EditedAt  *time.Time `json:"edited_at" db:"edited_at"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}
```

### Database Migrations
```bash
# Create new migration
make migrate-create name=add_user_preferences

# Run migrations
make migrate-up

# Rollback migration
make migrate-down

# Check migration status
make migrate-status
```

## 🌐 API Implementation

### HTTP Server Setup
```go
// main.go
func main() {
    // Load configuration
    cfg := config.Load()
    
    // Initialize database
    db := database.New(cfg.DatabaseURL)
    
    // Initialize Redis
    redis := redis.New(cfg.RedisURL)
    
    // Initialize services
    authService := auth.NewService(db, redis)
    chatService := chat.NewService(db, redis)
    
    // Setup routes
    router := setupRoutes(authService, chatService)
    
    // Start server
    log.Fatal(router.Run(":" + cfg.Port))
}
```

### Authentication Middleware
```go
func AuthMiddleware(authService *auth.Service) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(401, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }
        
        user, err := authService.ValidateToken(token)
        if err != nil {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        c.Set("user", user)
        c.Next()
    }
}
```

### WebSocket Implementation
```go
// websocket/handler.go
type Hub struct {
    clients    map[*Client]bool
    register   chan *Client
    unregister chan *Client
    broadcast  chan []byte
    rooms      map[string]map[*Client]bool
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
            h.joinRoom(client)
            
        case client := <-h.unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
                h.leaveRoom(client)
            }
            
        case message := <-h.broadcast:
            h.broadcastToRoom(message)
        }
    }
}
```

## 🔐 Security Implementation

### JWT Token Management
```go
type JWTService struct {
    secretKey []byte
    issuer    string
}

func (j *JWTService) GenerateToken(user *User) (*TokenPair, error) {
    // Generate access token
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "email":   user.Email,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
        "iat":     time.Now().Unix(),
        "iss":     j.issuer,
    })
    
    // Generate refresh token
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
        "iat":     time.Now().Unix(),
        "iss":     j.issuer,
    })
    
    return &TokenPair{
        AccessToken:  accessToken.SignedString(j.secretKey),
        RefreshToken: refreshToken.SignedString(j.secretKey),
    }, nil
}
```

### Rate Limiting
```go
func RateLimitMiddleware() gin.HandlerFunc {
    store := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
    return gin_rate_limit.RateLimiter(store, gin_rate_limit.WithLimit(100))
}
```

## 📁 File Upload Implementation

### File Service
```go
type FileService struct {
    storage Storage
    db      *sql.DB
}

func (f *FileService) UploadFile(file *multipart.FileHeader, channelID uuid.UUID, userID uuid.UUID) (*File, error) {
    // Generate unique filename
    filename := generateUniqueFilename(file.Filename)
    
    // Upload to storage
    url, err := f.storage.Upload(file, filename)
    if err != nil {
        return nil, err
    }
    
    // Save to database
    fileRecord := &File{
        ID:        uuid.New(),
        Filename:  filename,
        OriginalName: file.Filename,
        MimeType:  file.Header.Get("Content-Type"),
        Size:      file.Size,
        URL:       url,
        ChannelID: channelID,
        UploadedBy: userID,
        CreatedAt: time.Now(),
    }
    
    return f.saveFile(fileRecord)
}
```

## 🔍 Search Implementation

### Elasticsearch Integration
```go
type SearchService struct {
    client *elasticsearch.Client
    index  string
}

func (s *SearchService) IndexMessage(message *Message) error {
    _, err := s.client.Index(
        s.index,
        esutil.NewJSONReader(message),
        s.client.Index.WithDocumentID(message.ID.String()),
    )
    return err
}

func (s *SearchService) SearchMessages(query string, workspaceID uuid.UUID) ([]*Message, error) {
    searchRequest := map[string]interface{}{
        "query": map[string]interface{}{
            "bool": map[string]interface{}{
                "must": []map[string]interface{}{
                    {
                        "multi_match": map[string]interface{}{
                            "query":  query,
                            "fields": []string{"content"},
                        },
                    },
                    {
                        "term": map[string]interface{}{
                            "workspace_id": workspaceID.String(),
                        },
                    },
                },
            },
        },
    }
    
    // Execute search
    // ... implementation
}
```

## 📱 Real-time Features

### Message Broadcasting
```go
func (h *Hub) BroadcastMessage(message *Message) {
    // Get channel members
    members := h.getChannelMembers(message.ChannelID)
    
    // Create WebSocket message
    wsMessage := WebSocketMessage{
        Type: "message_sent",
        Data: message,
    }
    
    // Broadcast to all channel members
    for member := range members {
        select {
        case member.send <- wsMessage:
        default:
            close(member.send)
            delete(h.clients, member)
        }
    }
}
```

### Typing Indicators
```go
func (h *Hub) HandleTyping(client *Client, channelID uuid.UUID) {
    typingMessage := WebSocketMessage{
        Type: "user_typing",
        Data: map[string]interface{}{
            "user_id":    client.UserID,
            "channel_id": channelID,
            "timestamp":  time.Now(),
        },
    }
    
    h.broadcastToChannel(channelID, typingMessage)
}
```

## 🧪 Testing

### Unit Tests
```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific package tests
go test ./internal/auth/...
```

### Integration Tests
```bash
# Start test database
make test-db-up

# Run integration tests
make test-integration

# Cleanup test database
make test-db-down
```

### Load Testing
```bash
# Install k6
brew install k6

# Run load tests
k6 run tests/load/websocket-test.js
```

## 🚀 Deployment

### Docker Configuration
```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
CMD ["./main"]
```

### Docker Compose
```yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://user:pass@postgres:5432/afrochat
      - REDIS_URL=redis://redis:6379
    depends_on:
      - postgres
      - redis

  websocket:
    build: .
    command: ["./main", "websocket"]
    ports:
      - "8081:8081"
    environment:
      - DATABASE_URL=postgres://user:pass@postgres:5432/afrochat
      - REDIS_URL=redis://redis:6379
    depends_on:
      - postgres
      - redis

  postgres:
    image: postgres:14
    environment:
      POSTGRES_DB: afrochat
      POSTGRES_USER: user
      POSTGRES_PASSWORD: pass
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:6-alpine
    ports:
      - "6379:6379"

  elasticsearch:
    image: elasticsearch:7.17.0
    environment:
      - discovery.type=single-node
    ports:
      - "9200:9200"
```

### Production Deployment
```bash
# Build for production
make build-prod

# Deploy with Docker Compose
docker-compose -f docker-compose.prod.yml up -d

# Run migrations in production
make migrate-up-prod
```

## 📊 Monitoring & Logging

### Structured Logging
```go
import "github.com/sirupsen/logrus"

func setupLogger() *logrus.Logger {
    log := logrus.New()
    log.SetFormatter(&logrus.JSONFormatter{})
    log.SetLevel(logrus.InfoLevel)
    return log
}
```

### Health Checks
```go
func healthCheck(c *gin.Context) {
    // Check database connection
    if err := db.Ping(); err != nil {
        c.JSON(503, gin.H{"status": "unhealthy", "database": "down"})
        return
    }
    
    // Check Redis connection
    if err := redis.Ping(); err != nil {
        c.JSON(503, gin.H{"status": "unhealthy", "redis": "down"})
        return
    }
    
    c.JSON(200, gin.H{"status": "healthy"})
}
```

## 🔧 Development Tools

### Makefile Commands
```makefile
.PHONY: dev dev-websocket dev-worker test build migrate-up migrate-down

# Development
dev:
	go run cmd/api/main.go

dev-websocket:
	go run cmd/websocket/main.go

dev-worker:
	go run cmd/worker/main.go

# Testing
test:
	go test ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Database
migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

# Build
build:
	go build -o bin/api cmd/api/main.go
	go build -o bin/websocket cmd/websocket/main.go
	go build -o bin/worker cmd/worker/main.go

build-prod:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/api cmd/api/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/websocket cmd/websocket/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/worker cmd/worker/main.go
```

## 📈 Performance Optimization

### Database Optimization
- **Connection Pooling**: Configure proper connection pool settings
- **Query Optimization**: Use EXPLAIN ANALYZE for slow queries
- **Indexing**: Create appropriate indexes for search queries
- **Caching**: Use Redis for frequently accessed data

### WebSocket Optimization
- **Connection Limits**: Implement connection limits per user
- **Message Queuing**: Queue messages for offline users
- **Room Management**: Efficient room-based broadcasting
- **Compression**: Enable WebSocket compression

## 🛡️ Security Best Practices

- **Input Validation**: Validate all input data
- **SQL Injection Prevention**: Use parameterized queries
- **CORS Configuration**: Proper CORS settings
- **Rate Limiting**: Implement rate limiting
- **HTTPS**: Use HTTPS in production
- **Secrets Management**: Use environment variables for secrets

## 📚 API Documentation

The complete API specification is available in the root directory as `api-spec.yaml`. You can:

1. **View Interactive Docs**: Use Swagger UI to explore endpoints
2. **Generate Client SDKs**: Use OpenAPI Generator for client libraries
3. **API Testing**: Use the spec for automated testing

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
