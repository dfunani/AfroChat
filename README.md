# AfroChat - Slack Clone

A fully-featured real-time communication platform built with modern technologies, offering web, iOS, and Android applications.

## 🚀 Project Overview

AfroChat is a comprehensive Slack clone that provides real-time messaging, file sharing, workspace management, and cross-platform synchronization. The platform is designed for scalability, performance, and type safety.

## 🏗️ System Architecture

### Multi-Platform Frontend Strategy

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Web Frontend  │    │   iOS App       │    │  Android App    │
│   (React + TS)  │    │   (Swift + UIK) │    │  (Kotlin + Jet) │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                    ┌─────────────────┐
                    │   API Gateway   │
                    │   (Go + Gin)    │
                    └─────────────────┘
                                 │
                    ┌─────────────────┐
                    │   Core Services │
                    │   (Go + Gin)    │
                    └─────────────────┘
                                 │
         ┌───────────────────────┼───────────────────────┐
         │                       │                       │
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   PostgreSQL    │    │     Redis        │    │   File Storage  │
│   (Primary DB)   │    │   (Cache/PubSub) │    │   (S3/MinIO)    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🛠️ Tech Stack

### Frontend Technologies

#### Web Application
- **React 18** with TypeScript
- **Vite** (build tool)
- **Tailwind CSS** (styling)
- **Zustand** (state management)
- **Socket.io-client** (real-time communication)
- **React Query** (server state management)

#### iOS Application
- **Swift** with **SwiftUI**
- **Combine** (reactive programming)
- **URLSession** (networking)
- **Core Data** (local storage)
- **UserNotifications** (push notifications)

#### Android Application
- **Kotlin** with **Jetpack Compose**
- **Flow** (reactive programming)
- **Retrofit** (networking)
- **Room** (local database)
- **Firebase Cloud Messaging** (push notifications)

### Backend Technologies

#### Core Backend
- **Go** with **Gin** framework
- **WebSocket** for real-time messaging
- **PostgreSQL** (primary database)
- **Redis** (caching, sessions, pub/sub)
- **Docker** (containerization)

#### Additional Services
- **Elasticsearch** (search functionality)
- **AWS S3** or **MinIO** (file storage)
- **SendGrid** or **AWS SES** (email notifications)
- **Nginx** (reverse proxy)

## 📱 Mobile Architecture

### iOS (Swift + SwiftUI)
```
├── Models/
│   ├── User.swift
│   ├── Channel.swift
│   ├── Message.swift
│   └── Workspace.swift
├── ViewModels/
│   ├── AuthViewModel.swift
│   ├── ChatViewModel.swift
│   └── WorkspaceViewModel.swift
├── Views/
│   ├── LoginView.swift
│   ├── ChannelListView.swift
│   ├── ChatView.swift
│   └── ProfileView.swift
├── Services/
│   ├── APIService.swift
│   ├── WebSocketService.swift
│   └── NotificationService.swift
└── Utils/
    ├── NetworkManager.swift
    └── KeychainManager.swift
```

### Android (Kotlin + Jetpack Compose)
```
├── data/
│   ├── models/
│   │   ├── User.kt
│   │   ├── Channel.kt
│   │   └── Message.kt
│   ├── repository/
│   │   ├── AuthRepository.kt
│   │   └── ChatRepository.kt
│   └── network/
│       ├── ApiService.kt
│       └── WebSocketService.kt
├── domain/
│   ├── usecases/
│   └── repository/
├── presentation/
│   ├── viewmodels/
│   ├── screens/
│   └── components/
└── di/
    └── AppModule.kt
```

## 🗄️ Database Schema

### Core Tables

```sql
-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    display_name VARCHAR(100),
    avatar_url TEXT,
    status VARCHAR(20) DEFAULT 'offline',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Workspaces table
CREATE TABLE workspaces (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    slug VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    owner_id UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Channels table
CREATE TABLE channels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    workspace_id UUID REFERENCES workspaces(id),
    is_private BOOLEAN DEFAULT FALSE,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Messages table
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL,
    channel_id UUID REFERENCES channels(id),
    user_id UUID REFERENCES users(id),
    thread_id UUID REFERENCES messages(id),
    edited_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Workspace members
CREATE TABLE workspace_members (
    workspace_id UUID REFERENCES workspaces(id),
    user_id UUID REFERENCES users(id),
    role VARCHAR(20) DEFAULT 'member',
    joined_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (workspace_id, user_id)
);

-- Channel members
CREATE TABLE channel_members (
    channel_id UUID REFERENCES channels(id),
    user_id UUID REFERENCES users(id),
    joined_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (channel_id, user_id)
);
```

## 🌐 API Design

### REST API Endpoints

```go
// Authentication
POST   /api/v1/auth/register
POST   /api/v1/auth/login
POST   /api/v1/auth/refresh
POST   /api/v1/auth/logout

// Users
GET    /api/v1/users/profile
PUT    /api/v1/users/profile
GET    /api/v1/users/search

// Workspaces
GET    /api/v1/workspaces
POST   /api/v1/workspaces
GET    /api/v1/workspaces/{id}
PUT    /api/v1/workspaces/{id}

// Channels
GET    /api/v1/workspaces/{id}/channels
POST   /api/v1/workspaces/{id}/channels
GET    /api/v1/channels/{id}
PUT    /api/v1/channels/{id}

// Messages
GET    /api/v1/channels/{id}/messages
POST   /api/v1/channels/{id}/messages
PUT    /api/v1/messages/{id}
DELETE /api/v1/messages/{id}

// Files
POST   /api/v1/files/upload
GET    /api/v1/files/{id}
DELETE /api/v1/files/{id}

// Video Calls
POST   /api/v1/calls/start
POST   /api/v1/calls/{id}/join
POST   /api/v1/calls/{id}/leave
POST   /api/v1/calls/{id}/end
GET    /api/v1/calls/{id}
GET    /api/v1/calls/{id}/participants
```

### WebSocket Events

```go
// Client to Server
type ClientMessage struct {
    Type string `json:"type"`
    Data interface{} `json:"data"`
}

// Server to Client
type ServerMessage struct {
    Type string `json:"type"`
    Data interface{} `json:"data"`
    Timestamp time.Time `json:"timestamp"`
}

// Event Types
const (
    EventMessageSent = "message_sent"
    EventMessageEdited = "message_edited"
    EventMessageDeleted = "message_deleted"
    EventUserTyping = "user_typing"
    EventUserJoined = "user_joined"
    EventUserLeft = "user_left"
    EventChannelCreated = "channel_created"
    EventChannelUpdated = "channel_updated"
    EventCallStarted = "call_started"
    EventCallEnded = "call_ended"
    EventUserJoinedCall = "user_joined_call"
    EventUserLeftCall = "user_left_call"
    EventWebRTCOffer = "webrtc_offer"
    EventWebRTCAnswer = "webrtc_answer"
    EventWebRTCIceCandidate = "webrtc_ice_candidate"
)
```

## 🔄 Real-time Synchronization

### Cross-Platform Sync Strategy
1. **WebSocket Connections**: All platforms connect to same WebSocket server
2. **Message Ordering**: Use timestamps and sequence numbers
3. **Offline Support**: Queue messages when offline, sync when reconnected
4. **Conflict Resolution**: Last-write-wins with user awareness

### Data Flow
```
User Action → Mobile App → API Gateway → Go Backend → Database
                ↓
WebSocket Broadcast → All Connected Clients (Web, iOS, Android)
```

## 🚀 Key Features

### Core Functionalities

#### Authentication & User Management
- **Backend**: JWT token management, OAuth integration, user profiles
- **Frontend**: Login/signup forms, profile editing, workspace invitations

#### Real-time Messaging
- **Backend**: WebSocket connections, message persistence, threading
- **Frontend**: Real-time message display, rich text formatting, reactions

#### Video Calling & Conferencing
- **Backend**: WebRTC signaling server, call management, participant tracking
- **Frontend**: HD video calls, screen sharing, call controls, participant management
- **Features**: Multi-participant calls, mute/unmute, video on/off, call recording

#### Workspace & Channel Management
- **Backend**: Workspace configuration, channel permissions, member management
- **Frontend**: Channel sidebar, creation modals, settings interface

#### File Sharing & Media
- **Backend**: File upload handling, storage management, thumbnails
- **Frontend**: Drag-and-drop upload, file previews, gallery view

#### Notifications
- **Backend**: Push notification service, email notifications, mention detection
- **Frontend**: Desktop notifications, in-app center, sound controls

#### Search & Discovery
- **Backend**: Full-text search indexing, result ranking
- **Frontend**: Global search bar, filters, advanced search

#### Mobile & Responsive
- **Backend**: Mobile-optimized APIs, push notifications
- **Frontend**: Responsive design, touch interactions, mobile navigation

## 🏗️ Backend Architecture

### Microservices Structure
```
├── cmd/
│   ├── api/                 # Main API server
│   ├── websocket/           # WebSocket service
│   ├── worker/              # Background workers
│   └── migration/           # Database migrations
├── internal/
│   ├── auth/                # Authentication service
│   ├── chat/                # Chat service
│   ├── workspace/           # Workspace service
│   ├── notification/        # Notification service
│   ├── file/                # File service
│   ├── search/              # Search service
│   └── video/               # Video calling service
├── pkg/
│   ├── database/            # Database connections
│   ├── redis/               # Redis client
│   ├── websocket/           # WebSocket handler
│   └── middleware/          # Middleware functions
└── api/
    ├── v1/                  # API version 1
    └── websocket/           # WebSocket endpoints
```

## 📱 Mobile Implementation Examples

### iOS WebSocket Manager
```swift
class WebSocketManager: ObservableObject {
    private var webSocket: URLSessionWebSocketTask?
    @Published var isConnected = false
    
    func connect() {
        // WebSocket connection logic
    }
    
    func sendMessage(_ message: Message) {
        // Send message via WebSocket
    }
}
```

### Android Repository Pattern
```kotlin
class ChatRepository @Inject constructor(
    private val apiService: ApiService,
    private val webSocketService: WebSocketService
) {
    suspend fun getMessages(channelId: String): Flow<List<Message>> {
        // Combine API and WebSocket data
    }
}
```

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL 14+
- Redis 6+
- Docker & Docker Compose

### Development Setup
1. Clone the repository
2. Set up environment variables
3. Run database migrations
4. Start the development servers
5. Access the web application at `http://localhost:3000`

## 📋 Development Roadmap

### Phase 1: Core Backend
- [ ] Authentication system
- [ ] Basic messaging
- [ ] WebSocket implementation
- [ ] Database schema

### Phase 2: Web Frontend
- [ ] React application setup
- [ ] Authentication UI
- [ ] Chat interface
- [ ] Real-time messaging

### Phase 3: Mobile Applications
- [ ] iOS app development
- [ ] Android app development
- [ ] Push notifications
- [ ] Offline support

### Phase 4: Advanced Features
- [ ] File sharing
- [ ] Search functionality
- [ ] Workspace management
- [ ] Integrations

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For support and questions, please open an issue in the repository or contact the development team.

---

**AfroChat** - Building the future of team communication 🚀
