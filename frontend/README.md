# AfroChat Frontend

Multi-platform frontend implementation for AfroChat, supporting Web (React), iOS (Swift), and Android (Kotlin) applications with shared API integration.

## 🏗️ Frontend Architecture Overview

### Platform-Specific Implementations
```
frontend/
├── web/                    # React + TypeScript web application
│   ├── src/
│   │   ├── components/     # Reusable UI components
│   │   ├── pages/         # Page components
│   │   ├── hooks/         # Custom React hooks
│   │   ├── services/      # API and WebSocket services
│   │   ├── store/         # State management (Zustand)
│   │   ├── utils/         # Utility functions
│   │   └── types/         # TypeScript type definitions
│   ├── public/            # Static assets
│   └── package.json       # Dependencies and scripts
├── ios/                   # Swift + SwiftUI iOS application
│   ├── AfroChat/
│   │   ├── Models/        # Data models
│   │   ├── ViewModels/    # MVVM view models
│   │   ├── Views/         # SwiftUI views
│   │   ├── Services/      # API and WebSocket services
│   │   └── Utils/         # Utility classes
│   └── AfroChat.xcodeproj # Xcode project
└── android/               # Kotlin + Jetpack Compose Android app
    ├── app/
    │   ├── src/main/java/com/afrochat/
    │   │   ├── data/      # Data layer (models, repository)
    │   │   ├── domain/    # Domain layer (use cases)
    │   │   ├── presentation/ # Presentation layer (UI)
    │   │   └── di/        # Dependency injection
    │   └── build.gradle   # Android build configuration
    └── build.gradle       # Project build configuration
```

## 🌐 Web Application (React + TypeScript)

### Technology Stack
- **React 18** with TypeScript
- **Vite** for fast development and building
- **Tailwind CSS** for styling
- **Zustand** for state management
- **React Query** for server state
- **Socket.io-client** for real-time communication
- **React Router** for navigation

### Project Structure
```
web/
├── src/
│   ├── components/
│   │   ├── ui/            # Base UI components
│   │   ├── chat/          # Chat-specific components
│   │   ├── workspace/    # Workspace components
│   │   └── common/        # Shared components
│   ├── pages/
│   │   ├── Login.tsx      # Authentication pages
│   │   ├── Workspace.tsx  # Workspace management
│   │   ├── Channel.tsx    # Channel view
│   │   └── Profile.tsx    # User profile
│   ├── hooks/
│   │   ├── useAuth.ts     # Authentication hook
│   │   ├── useWebSocket.ts # WebSocket hook
│   │   └── useChat.ts     # Chat functionality
│   ├── services/
│   │   ├── api.ts         # REST API client
│   │   ├── websocket.ts   # WebSocket service
│   │   └── auth.ts        # Authentication service
│   ├── store/
│   │   ├── authStore.ts   # Authentication state
│   │   ├── chatStore.ts   # Chat state
│   │   └── workspaceStore.ts # Workspace state
│   └── types/
│       ├── api.ts         # API type definitions
│       ├── user.ts        # User types
│       └── message.ts     # Message types
```

### Key Features Implementation

#### Authentication Flow
```typescript
// hooks/useAuth.ts
export const useAuth = () => {
  const { user, setUser, login, logout } = useAuthStore();
  
  const handleLogin = async (email: string, password: string) => {
    try {
      const response = await authService.login({ email, password });
      setUser(response.user);
      localStorage.setItem('accessToken', response.accessToken);
      localStorage.setItem('refreshToken', response.refreshToken);
    } catch (error) {
      throw new Error('Login failed');
    }
  };
  
  return { user, login: handleLogin, logout };
};
```

#### Real-time Messaging
```typescript
// hooks/useWebSocket.ts
export const useWebSocket = (channelId: string) => {
  const [socket, setSocket] = useState<Socket | null>(null);
  const [messages, setMessages] = useState<Message[]>([]);
  
  useEffect(() => {
    const newSocket = io(API_BASE_URL, {
      auth: {
        token: localStorage.getItem('accessToken')
      }
    });
    
    newSocket.emit('join_channel', channelId);
    
    newSocket.on('message_sent', (message: Message) => {
      setMessages(prev => [...prev, message]);
    });
    
    setSocket(newSocket);
    
    return () => newSocket.close();
  }, [channelId]);
  
  const sendMessage = (content: string) => {
    if (socket) {
      socket.emit('send_message', { channelId, content });
    }
  };
  
  return { messages, sendMessage };
};
```

#### State Management
```typescript
// store/chatStore.ts
interface ChatState {
  channels: Channel[];
  currentChannel: Channel | null;
  messages: Message[];
  typingUsers: string[];
}

export const useChatStore = create<ChatState>((set, get) => ({
  channels: [],
  currentChannel: null,
  messages: [],
  typingUsers: [],
  
  setCurrentChannel: (channel: Channel) => 
    set({ currentChannel: channel }),
    
  addMessage: (message: Message) =>
    set(state => ({ 
      messages: [...state.messages, message] 
    })),
    
  setTypingUsers: (users: string[]) =>
    set({ typingUsers: users }),
}));
```

## 📱 iOS Application (Swift + SwiftUI)

### Technology Stack
- **Swift 5.9+** with **SwiftUI**
- **Combine** for reactive programming
- **URLSession** for networking
- **Core Data** for local storage
- **UserNotifications** for push notifications

### Project Structure
```
ios/AfroChat/
├── Models/
│   ├── User.swift
│   ├── Channel.swift
│   ├── Message.swift
│   └── Workspace.swift
├── ViewModels/
│   ├── AuthViewModel.swift
│   ├── ChatViewModel.swift
│   ├── WorkspaceViewModel.swift
│   └── ChannelListViewModel.swift
├── Views/
│   ├── Authentication/
│   │   ├── LoginView.swift
│   │   └── RegisterView.swift
│   ├── Workspace/
│   │   ├── WorkspaceView.swift
│   │   └── ChannelListView.swift
│   ├── Chat/
│   │   ├── ChatView.swift
│   │   ├── MessageView.swift
│   │   └── MessageInputView.swift
│   └── Profile/
│       └── ProfileView.swift
├── Services/
│   ├── APIService.swift
│   ├── WebSocketService.swift
│   ├── AuthService.swift
│   └── NotificationService.swift
└── Utils/
    ├── NetworkManager.swift
    ├── KeychainManager.swift
    └── Extensions.swift
```

### Key Features Implementation

#### Authentication Service
```swift
// Services/AuthService.swift
class AuthService: ObservableObject {
    @Published var isAuthenticated = false
    @Published var currentUser: User?
    
    private let apiService = APIService()
    private let keychain = KeychainManager()
    
    func login(email: String, password: String) async throws {
        let request = LoginRequest(email: email, password: password)
        let response: AuthResponse = try await apiService.post("/auth/login", body: request)
        
        await MainActor.run {
            self.currentUser = response.user
            self.isAuthenticated = true
        }
        
        keychain.store(token: response.accessToken, for: .accessToken)
        keychain.store(token: response.refreshToken, for: .refreshToken)
    }
    
    func logout() {
        currentUser = nil
        isAuthenticated = false
        keychain.deleteAll()
    }
}
```

#### WebSocket Service
```swift
// Services/WebSocketService.swift
class WebSocketService: ObservableObject {
    @Published var isConnected = false
    @Published var messages: [Message] = []
    
    private var webSocket: URLSessionWebSocketTask?
    private var urlSession: URLSession?
    
    func connect() {
        guard let url = URL(string: "ws://localhost:8080/api/v1/ws?token=\(getAccessToken())") else { return }
        
        urlSession = URLSession(configuration: .default)
        webSocket = urlSession?.webSocketTask(with: url)
        webSocket?.resume()
        
        receiveMessage()
    }
    
    private func receiveMessage() {
        webSocket?.receive { [weak self] result in
            switch result {
            case .success(let message):
                switch message {
                case .string(let text):
                    self?.handleMessage(text)
                case .data(let data):
                    self?.handleData(data)
                @unknown default:
                    break
                }
                self?.receiveMessage()
            case .failure(let error):
                print("WebSocket error: \(error)")
            }
        }
    }
    
    func sendMessage(_ message: Message) {
        let messageData = try? JSONEncoder().encode(message)
        webSocket?.send(.data(messageData ?? Data())) { error in
            if let error = error {
                print("Send error: \(error)")
            }
        }
    }
}
```

#### Chat View Model
```swift
// ViewModels/ChatViewModel.swift
class ChatViewModel: ObservableObject {
    @Published var messages: [Message] = []
    @Published var isTyping = false
    @Published var typingUsers: [User] = []
    
    private let webSocketService = WebSocketService()
    private let apiService = APIService()
    
    func loadMessages(channelId: String) async {
        do {
            let messages: [Message] = try await apiService.get("/channels/\(channelId)/messages")
            await MainActor.run {
                self.messages = messages
            }
        } catch {
            print("Failed to load messages: \(error)")
        }
    }
    
    func sendMessage(content: String, channelId: String) {
        let message = Message(
            id: UUID(),
            content: content,
            channelId: UUID(uuidString: channelId)!,
            userId: AuthService.shared.currentUser?.id ?? UUID(),
            createdAt: Date()
        )
        
        webSocketService.sendMessage(message)
    }
}
```

## 🤖 Android Application (Kotlin + Jetpack Compose)

### Technology Stack
- **Kotlin** with **Jetpack Compose**
- **Flow** for reactive programming
- **Retrofit** for networking
- **Room** for local database
- **Hilt** for dependency injection
- **Firebase Cloud Messaging** for push notifications

### Project Structure
```
android/app/src/main/java/com/afrochat/
├── data/
│   ├── models/
│   │   ├── User.kt
│   │   ├── Channel.kt
│   │   └── Message.kt
│   ├── repository/
│   │   ├── AuthRepository.kt
│   │   ├── ChatRepository.kt
│   │   └── WorkspaceRepository.kt
│   ├── network/
│   │   ├── ApiService.kt
│   │   ├── WebSocketService.kt
│   │   └── ApiClient.kt
│   └── local/
│       ├── database/
│       │   ├── AfroChatDatabase.kt
│       │   ├── entities/
│       │   └── dao/
│       └── preferences/
│           └── AuthPreferences.kt
├── domain/
│   ├── usecases/
│   │   ├── LoginUseCase.kt
│   │   ├── SendMessageUseCase.kt
│   │   └── GetMessagesUseCase.kt
│   └── repository/
│       ├── AuthRepository.kt
│       └── ChatRepository.kt
├── presentation/
│   ├── viewmodels/
│   │   ├── AuthViewModel.kt
│   │   ├── ChatViewModel.kt
│   │   └── WorkspaceViewModel.kt
│   ├── screens/
│   │   ├── LoginScreen.kt
│   │   ├── WorkspaceScreen.kt
│   │   ├── ChannelScreen.kt
│   │   └── ChatScreen.kt
│   └── components/
│       ├── MessageItem.kt
│       ├── ChannelItem.kt
│       └── UserItem.kt
└── di/
    └── AppModule.kt
```

### Key Features Implementation

#### Repository Pattern
```kotlin
// data/repository/ChatRepository.kt
@Singleton
class ChatRepository @Inject constructor(
    private val apiService: ApiService,
    private val webSocketService: WebSocketService,
    private val messageDao: MessageDao
) {
    fun getMessages(channelId: String): Flow<List<Message>> {
        return messageDao.getMessagesByChannel(channelId)
            .onStart {
                // Load from API and cache locally
                val messages = apiService.getMessages(channelId)
                messageDao.insertMessages(messages)
            }
    }
    
    suspend fun sendMessage(content: String, channelId: String) {
        val message = Message(
            id = UUID.randomUUID().toString(),
            content = content,
            channelId = channelId,
            userId = getCurrentUserId(),
            createdAt = System.currentTimeMillis()
        )
        
        webSocketService.sendMessage(message)
    }
}
```

#### ViewModel with State Management
```kotlin
// presentation/viewmodels/ChatViewModel.kt
@HiltViewModel
class ChatViewModel @Inject constructor(
    private val chatRepository: ChatRepository,
    private val authRepository: AuthRepository
) : ViewModel() {
    
    private val _uiState = MutableStateFlow(ChatUiState())
    val uiState: StateFlow<ChatUiState> = _uiState.asStateFlow()
    
    fun loadMessages(channelId: String) {
        viewModelScope.launch {
            chatRepository.getMessages(channelId)
                .catch { error ->
                    _uiState.value = _uiState.value.copy(
                        error = error.message
                    )
                }
                .collect { messages ->
                    _uiState.value = _uiState.value.copy(
                        messages = messages,
                        isLoading = false
                    )
                }
        }
    }
    
    fun sendMessage(content: String, channelId: String) {
        viewModelScope.launch {
            try {
                chatRepository.sendMessage(content, channelId)
            } catch (e: Exception) {
                _uiState.value = _uiState.value.copy(
                    error = e.message
                )
            }
        }
    }
}

data class ChatUiState(
    val messages: List<Message> = emptyList(),
    val isLoading: Boolean = true,
    val error: String? = null
)
```

#### Compose UI
```kotlin
// presentation/screens/ChatScreen.kt
@Composable
fun ChatScreen(
    channelId: String,
    viewModel: ChatViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsState()
    
    LaunchedEffect(channelId) {
        viewModel.loadMessages(channelId)
    }
    
    Column(
        modifier = Modifier.fillMaxSize()
    ) {
        // Messages list
        LazyColumn(
            modifier = Modifier.weight(1f),
            contentPadding = PaddingValues(16.dp)
        ) {
            items(uiState.messages) { message ->
                MessageItem(
                    message = message,
                    modifier = Modifier.padding(vertical = 4.dp)
                )
            }
        }
        
        // Message input
        MessageInput(
            onSendMessage = { content ->
                viewModel.sendMessage(content, channelId)
            }
        )
    }
}
```

## 🔄 Shared API Integration

### API Client Generation
All platforms use the same OpenAPI specification to generate type-safe clients:

```bash
# Generate TypeScript client for web
npx @openapitools/openapi-generator-cli generate \
  -i ../api-spec.yaml \
  -g typescript-fetch \
  -o ./src/generated

# Generate Swift client for iOS
npx @openapitools/openapi-generator-cli generate \
  -i ../api-spec.yaml \
  -g swift5 \
  -o ./AfroChat/Generated

# Generate Kotlin client for Android
npx @openapitools/openapi-generator-cli generate \
  -i ../api-spec.yaml \
  -g kotlin \
  -o ./app/src/main/java/com/afrochat/generated
```

### WebSocket Implementation
All platforms implement the same WebSocket protocol:

```typescript
// WebSocket message format (shared across all platforms)
interface WebSocketMessage {
  type: 'message_sent' | 'message_edited' | 'user_typing' | 'user_joined';
  data: any;
  timestamp: string;
}
```

## 🚀 Getting Started

### Web Application
```bash
cd frontend/web
npm install
npm run dev
```

### iOS Application
```bash
cd frontend/ios
open AfroChat.xcodeproj
# Build and run in Xcode
```

### Android Application
```bash
cd frontend/android
./gradlew build
# Open in Android Studio
```

## 🧪 Testing

### Web Testing
```bash
# Unit tests
npm run test

# E2E tests
npm run test:e2e

# Coverage
npm run test:coverage
```

### iOS Testing
```swift
// Unit tests
func testLogin() async {
    let authService = AuthService()
    try await authService.login(email: "test@example.com", password: "password")
    XCTAssertTrue(authService.isAuthenticated)
}
```

### Android Testing
```kotlin
// Unit tests
@Test
fun `login should set authenticated state`() = runTest {
    val authRepository = mockk<AuthRepository>()
    coEvery { authRepository.login(any(), any()) } returns User(...)
    
    val viewModel = AuthViewModel(authRepository)
    viewModel.login("test@example.com", "password")
    
    assertTrue(viewModel.isAuthenticated.value)
}
```

## 📱 Platform-Specific Features

### iOS Features
- **Push Notifications**: UserNotifications framework
- **Background App Refresh**: Keep connection alive
- **Keychain**: Secure token storage
- **Haptic Feedback**: Enhanced user experience

### Android Features
- **Firebase Cloud Messaging**: Push notifications
- **Background Services**: Maintain WebSocket connection
- **SharedPreferences**: Secure token storage
- **Material Design**: Native Android UI

### Web Features
- **Service Workers**: Offline support
- **Push Notifications**: Browser notifications
- **Local Storage**: Token persistence
- **PWA Support**: Installable web app

## 🔧 Development Tools

### Code Generation
- **OpenAPI Generator**: Generate API clients
- **TypeScript**: Type safety for web
- **Swift**: Type safety for iOS
- **Kotlin**: Type safety for Android

### State Management
- **Zustand**: Lightweight state management (Web)
- **Combine**: Reactive programming (iOS)
- **Flow**: Reactive streams (Android)

### Testing
- **Jest**: Web testing framework
- **XCTest**: iOS testing framework
- **JUnit**: Android testing framework

## 📊 Performance Optimization

### Web Optimization
- **Code Splitting**: Lazy load components
- **Bundle Analysis**: Optimize bundle size
- **Caching**: React Query caching
- **Virtual Scrolling**: Large message lists

### iOS Optimization
- **Lazy Loading**: Load messages on demand
- **Image Caching**: Efficient image handling
- **Memory Management**: Proper object lifecycle
- **Background Processing**: Efficient WebSocket handling

### Android Optimization
- **RecyclerView**: Efficient list rendering
- **Image Loading**: Glide/Coil for images
- **Memory Management**: Proper lifecycle handling
- **Background Processing**: WorkManager for tasks

## 🛡️ Security

### Token Management
- **Secure Storage**: Keychain (iOS), EncryptedSharedPreferences (Android)
- **Token Refresh**: Automatic token renewal
- **Logout Cleanup**: Clear all stored data

### Network Security
- **HTTPS Only**: Secure API communication
- **Certificate Pinning**: Prevent MITM attacks
- **Request Validation**: Validate all API responses

## 📚 Documentation

Each platform has detailed documentation:
- **Web**: Component documentation with Storybook
- **iOS**: Code documentation with Swift DocC
- **Android**: KDoc documentation

## 🤝 Contributing

1. Follow platform-specific coding standards
2. Write tests for new features
3. Update documentation
4. Ensure cross-platform consistency

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
