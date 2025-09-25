# AfroChat iOS App

A native iOS application built with Swift 5.9+ and SwiftUI, providing a seamless Slack-like communication experience on iOS devices.

## 🏗️ Architecture Overview

### Technology Stack
- **Swift 5.9+** with **SwiftUI** for modern UI development
- **Combine** for reactive programming and data flow
- **URLSession** for networking and API communication
- **Core Data** for local data persistence
- **UserNotifications** for push notifications
- **Keychain Services** for secure token storage
- **WebSocket** for real-time communication

### Project Structure
```
ios/
├── AfroChat/
│   ├── Models/
│   │   ├── User.swift
│   │   ├── Workspace.swift
│   │   ├── Channel.swift
│   │   ├── Message.swift
│   │   ├── File.swift
│   │   ├── Reaction.swift
│   │   └── APIResponse.swift
│   ├── ViewModels/
│   │   ├── AuthViewModel.swift
│   │   ├── WorkspaceViewModel.swift
│   │   ├── ChannelListViewModel.swift
│   │   ├── ChatViewModel.swift
│   │   ├── ProfileViewModel.swift
│   │   └── SettingsViewModel.swift
│   ├── Views/
│   │   ├── Authentication/
│   │   │   ├── LoginView.swift
│   │   │   ├── RegisterView.swift
│   │   │   ├── ForgotPasswordView.swift
│   │   │   └── AuthLoadingView.swift
│   │   ├── Workspace/
│   │   │   ├── WorkspaceView.swift
│   │   │   ├── ChannelListView.swift
│   │   │   ├── ChannelItemView.swift
│   │   │   ├── UserListView.swift
│   │   │   └── WorkspaceHeaderView.swift
│   │   ├── Chat/
│   │   │   ├── ChatView.swift
│   │   │   ├── MessageListView.swift
│   │   │   ├── MessageItemView.swift
│   │   │   ├── MessageInputView.swift
│   │   │   ├── TypingIndicatorView.swift
│   │   │   ├── MessageReactionView.swift
│   │   │   └── ThreadView.swift
│   │   ├── Profile/
│   │   │   ├── ProfileView.swift
│   │   │   ├── EditProfileView.swift
│   │   │   └── SettingsView.swift
│   │   └── Common/
│   │       ├── LoadingView.swift
│   │       ├── ErrorView.swift
│   │       ├── EmptyStateView.swift
│   │       └── CustomNavigationView.swift
│   ├── Services/
│   │   ├── APIService.swift
│   │   ├── WebSocketService.swift
│   │   ├── AuthService.swift
│   │   ├── ChatService.swift
│   │   ├── WorkspaceService.swift
│   │   ├── FileService.swift
│   │   ├── NotificationService.swift
│   │   └── CoreDataService.swift
│   ├── Utils/
│   │   ├── NetworkManager.swift
│   │   ├── KeychainManager.swift
│   │   ├── DateFormatter+Extensions.swift
│   │   ├── Color+Extensions.swift
│   │   ├── View+Extensions.swift
│   │   └── Constants.swift
│   ├── Resources/
│   │   ├── Assets.xcassets
│   │   ├── Info.plist
│   │   ├── Localizable.strings
│   │   └── AfroChat.xcdatamodeld
│   ├── App.swift
│   └── ContentView.swift
├── AfroChat.xcodeproj
├── Package.swift
└── README.md
```

## 🚀 Getting Started

### Prerequisites
- Xcode 15.0+
- iOS 16.0+ deployment target
- Backend API running on localhost:8080
- CocoaPods or Swift Package Manager

### Installation

1. **Clone and setup**
```bash
cd frontend/ios
open AfroChat.xcodeproj
```

2. **Install Dependencies**
```bash
# Using CocoaPods
pod install

# Using Swift Package Manager (built into Xcode)
# Dependencies are managed in Xcode
```

3. **Configure Environment**
```swift
// Utils/Constants.swift
struct Constants {
    static let apiBaseURL = "http://localhost:8080/api/v1"
    static let websocketURL = "ws://localhost:8080/api/v1/ws"
    static let appName = "AfroChat"
    static let appVersion = "1.0.0"
}
```

4. **Build and Run**
```bash
# Build for simulator
xcodebuild -scheme AfroChat -destination 'platform=iOS Simulator,name=iPhone 15'

# Run in Xcode
# Press Cmd+R to build and run
```

## 🔧 Configuration

### Info.plist Configuration
```xml
<!-- Info.plist -->
<key>CFBundleDisplayName</key>
<string>AfroChat</string>
<key>CFBundleIdentifier</key>
<string>com.afrochat.app</string>
<key>CFBundleVersion</key>
<string>1.0.0</string>
<key>CFBundleShortVersionString</key>
<string>1.0.0</string>

<!-- Network Security -->
<key>NSAppTransportSecurity</key>
<dict>
    <key>NSAllowsArbitraryLoads</key>
    <true/>
    <key>NSExceptionDomains</key>
    <dict>
        <key>localhost</key>
        <dict>
            <key>NSExceptionAllowsInsecureHTTPLoads</key>
            <true/>
        </dict>
    </dict>
</dict>

<!-- Background Modes -->
<key>UIBackgroundModes</key>
<array>
    <string>background-processing</string>
    <string>background-fetch</string>
</array>

<!-- Push Notifications -->
<key>UIBackgroundModes</key>
<array>
    <string>remote-notification</string>
</array>
```

### App Configuration
```swift
// App.swift
import SwiftUI

@main
struct AfroChatApp: App {
    let persistenceController = PersistenceController.shared
    
    var body: some Scene {
        WindowGroup {
            ContentView()
                .environment(\.managedObjectContext, persistenceController.container.viewContext)
                .onAppear {
                    configureAppearance()
                }
        }
    }
    
    private func configureAppearance() {
        // Configure navigation bar appearance
        let appearance = UINavigationBarAppearance()
        appearance.configureWithOpaqueBackground()
        appearance.backgroundColor = UIColor.systemBackground
        UINavigationBar.appearance().standardAppearance = appearance
        UINavigationBar.appearance().scrollEdgeAppearance = appearance
    }
}
```

## 📱 Data Models

### Core Models
```swift
// Models/User.swift
import Foundation

struct User: Codable, Identifiable, Hashable {
    let id: UUID
    let email: String
    let username: String
    let displayName: String
    let avatarURL: String?
    let status: UserStatus
    let createdAt: Date
    let updatedAt: Date
    
    enum CodingKeys: String, CodingKey {
        case id, email, username, displayName
        case avatarURL = "avatar_url"
        case status, createdAt, updatedAt
    }
}

enum UserStatus: String, Codable, CaseIterable {
    case online = "online"
    case offline = "offline"
    case away = "away"
    case busy = "busy"
    
    var color: Color {
        switch self {
        case .online: return .green
        case .offline: return .gray
        case .away: return .yellow
        case .busy: return .red
        }
    }
}
```

### Message Model
```swift
// Models/Message.swift
import Foundation

struct Message: Codable, Identifiable, Hashable {
    let id: UUID
    let content: String
    let channelID: UUID
    let userID: UUID
    let user: User
    let threadID: UUID?
    let editedAt: Date?
    let reactions: [Reaction]
    let createdAt: Date
    
    enum CodingKeys: String, CodingKey {
        case id, content
        case channelID = "channel_id"
        case userID = "user_id"
        case user, threadID, editedAt, reactions, createdAt
    }
}

struct Reaction: Codable, Identifiable, Hashable {
    let id = UUID()
    let emoji: String
    let count: Int
    let users: [User]
}
```

## 🔄 State Management

### ViewModels with Combine
```swift
// ViewModels/ChatViewModel.swift
import Foundation
import Combine

@MainActor
class ChatViewModel: ObservableObject {
    @Published var messages: [Message] = []
    @Published var isLoading = false
    @Published var errorMessage: String?
    @Published var typingUsers: [User] = []
    @Published var isTyping = false
    
    private let chatService: ChatService
    private let websocketService: WebSocketService
    private var cancellables = Set<AnyCancellable>()
    private var currentChannelID: UUID?
    
    init(chatService: ChatService = ChatService(), 
         websocketService: WebSocketService = WebSocketService()) {
        self.chatService = chatService
        self.websocketService = websocketService
        
        setupWebSocketObservers()
    }
    
    func loadMessages(for channelID: UUID) {
        currentChannelID = channelID
        isLoading = true
        
        chatService.getMessages(channelID: channelID)
            .receive(on: DispatchQueue.main)
            .sink(
                receiveCompletion: { [weak self] completion in
                    self?.isLoading = false
                    if case .failure(let error) = completion {
                        self?.errorMessage = error.localizedDescription
                    }
                },
                receiveValue: { [weak self] messages in
                    self?.messages = messages
                }
            )
            .store(in: &cancellables)
    }
    
    func sendMessage(content: String) {
        guard let channelID = currentChannelID else { return }
        
        let message = SendMessageRequest(
            content: content,
            channelID: channelID
        )
        
        chatService.sendMessage(message)
            .receive(on: DispatchQueue.main)
            .sink(
                receiveCompletion: { [weak self] completion in
                    if case .failure(let error) = completion {
                        self?.errorMessage = error.localizedDescription
                    }
                },
                receiveValue: { [weak self] newMessage in
                    self?.messages.append(newMessage)
                }
            )
            .store(in: &cancellables)
    }
    
    func startTyping() {
        guard let channelID = currentChannelID else { return }
        websocketService.sendTyping(channelID: channelID)
    }
    
    func stopTyping() {
        guard let channelID = currentChannelID else { return }
        websocketService.stopTyping(channelID: channelID)
    }
    
    private func setupWebSocketObservers() {
        websocketService.messagePublisher
            .receive(on: DispatchQueue.main)
            .sink { [weak self] message in
                self?.messages.append(message)
            }
            .store(in: &cancellables)
        
        websocketService.typingUsersPublisher
            .receive(on: DispatchQueue.main)
            .sink { [weak self] users in
                self?.typingUsers = users
            }
            .store(in: &cancellables)
    }
}
```

## 🌐 Networking

### API Service
```swift
// Services/APIService.swift
import Foundation
import Combine

class APIService {
    private let session: URLSession
    private let baseURL: URL
    
    init(baseURL: URL = URL(string: Constants.apiBaseURL)!) {
        self.baseURL = baseURL
        self.session = URLSession.shared
    }
    
    func request<T: Codable>(
        endpoint: String,
        method: HTTPMethod = .GET,
        body: Data? = nil,
        responseType: T.Type
    ) -> AnyPublisher<T, Error> {
        guard let url = URL(string: endpoint, relativeTo: baseURL) else {
            return Fail(error: APIError.invalidURL)
                .eraseToAnyPublisher()
        }
        
        var request = URLRequest(url: url)
        request.httpMethod = method.rawValue
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        
        if let token = KeychainManager.shared.getAccessToken() {
            request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        }
        
        if let body = body {
            request.httpBody = body
        }
        
        return session.dataTaskPublisher(for: request)
            .map(\.data)
            .decode(type: T.self, decoder: JSONDecoder())
            .eraseToAnyPublisher()
    }
}

enum HTTPMethod: String {
    case GET = "GET"
    case POST = "POST"
    case PUT = "PUT"
    case DELETE = "DELETE"
}

enum APIError: Error, LocalizedError {
    case invalidURL
    case noData
    case decodingError
    case networkError(Error)
    
    var errorDescription: String? {
        switch self {
        case .invalidURL:
            return "Invalid URL"
        case .noData:
            return "No data received"
        case .decodingError:
            return "Failed to decode response"
        case .networkError(let error):
            return error.localizedDescription
        }
    }
}
```

### WebSocket Service
```swift
// Services/WebSocketService.swift
import Foundation
import Combine

class WebSocketService: NSObject, ObservableObject {
    @Published var isConnected = false
    @Published var connectionError: String?
    
    private var webSocketTask: URLSessionWebSocketTask?
    private var urlSession: URLSession?
    private let messageSubject = PassthroughSubject<Message, Never>()
    private let typingUsersSubject = PassthroughSubject<[User], Never>()
    
    var messagePublisher: AnyPublisher<Message, Never> {
        messageSubject.eraseToAnyPublisher()
    }
    
    var typingUsersPublisher: AnyPublisher<[User], Never> {
        typingUsersSubject.eraseToAnyPublisher()
    }
    
    func connect() {
        guard let url = URL(string: Constants.websocketURL),
              let token = KeychainManager.shared.getAccessToken() else {
            connectionError = "Invalid WebSocket URL or missing token"
            return
        }
        
        var components = URLComponents(url: url, resolvingAgainstBaseURL: false)
        components?.queryItems = [URLQueryItem(name: "token", value: token)]
        
        guard let finalURL = components?.url else {
            connectionError = "Failed to construct WebSocket URL"
            return
        }
        
        urlSession = URLSession(configuration: .default, delegate: self, delegateQueue: nil)
        webSocketTask = urlSession?.webSocketTask(with: finalURL)
        webSocketTask?.resume()
        
        receiveMessage()
    }
    
    func disconnect() {
        webSocketTask?.cancel(with: .goingAway, reason: nil)
        webSocketTask = nil
        urlSession = nil
        isConnected = false
    }
    
    func joinChannel(channelID: UUID) {
        let message = WebSocketMessage(type: "join_channel", data: ["channel_id": channelID.uuidString])
        sendMessage(message)
    }
    
    func leaveChannel(channelID: UUID) {
        let message = WebSocketMessage(type: "leave_channel", data: ["channel_id": channelID.uuidString])
        sendMessage(message)
    }
    
    func sendMessage(_ message: SendMessageRequest) {
        let wsMessage = WebSocketMessage(type: "send_message", data: message)
        sendMessage(wsMessage)
    }
    
    func sendTyping(channelID: UUID) {
        let message = WebSocketMessage(type: "typing", data: ["channel_id": channelID.uuidString])
        sendMessage(message)
    }
    
    func stopTyping(channelID: UUID) {
        let message = WebSocketMessage(type: "stop_typing", data: ["channel_id": channelID.uuidString])
        sendMessage(message)
    }
    
    private func receiveMessage() {
        webSocketTask?.receive { [weak self] result in
            switch result {
            case .success(let message):
                self?.handleMessage(message)
                self?.receiveMessage()
            case .failure(let error):
                DispatchQueue.main.async {
                    self?.connectionError = error.localizedDescription
                    self?.isConnected = false
                }
            }
        }
    }
    
    private func handleMessage(_ message: URLSessionWebSocketTask.Message) {
        switch message {
        case .string(let text):
            if let data = text.data(using: .utf8),
               let wsMessage = try? JSONDecoder().decode(WebSocketMessage.self, from: data) {
                handleWebSocketMessage(wsMessage)
            }
        case .data(let data):
            if let wsMessage = try? JSONDecoder().decode(WebSocketMessage.self, from: data) {
                handleWebSocketMessage(wsMessage)
            }
        @unknown default:
            break
        }
    }
    
    private func handleWebSocketMessage(_ message: WebSocketMessage) {
        switch message.type {
        case "message_sent":
            if let messageData = message.data as? [String: Any],
               let messageJSON = try? JSONSerialization.data(withJSONObject: messageData),
               let newMessage = try? JSONDecoder().decode(Message.self, from: messageJSON) {
                messageSubject.send(newMessage)
            }
        case "user_typing":
            // Handle typing indicator
            break
        default:
            break
        }
    }
    
    private func sendMessage(_ message: WebSocketMessage) {
        guard let data = try? JSONEncoder().encode(message) else { return }
        webSocketTask?.send(.data(data)) { error in
            if let error = error {
                print("WebSocket send error: \(error)")
            }
        }
    }
}

extension WebSocketService: URLSessionWebSocketDelegate {
    func urlSession(_ session: URLSession, webSocketTask: URLSessionWebSocketTask, didOpenWithProtocol protocol: String?) {
        DispatchQueue.main.async {
            self.isConnected = true
            self.connectionError = nil
        }
    }
    
    func urlSession(_ session: URLSession, webSocketTask: URLSessionWebSocketTask, didCloseWith closeCode: URLSessionWebSocketTask.CloseCode, reason: Data?) {
        DispatchQueue.main.async {
            self.isConnected = false
        }
    }
}
```

## 🔐 Security

### Keychain Manager
```swift
// Utils/KeychainManager.swift
import Foundation
import Security

class KeychainManager {
    static let shared = KeychainManager()
    
    private init() {}
    
    func store(token: String, for key: KeychainKey) {
        let data = token.data(using: .utf8)!
        
        let query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrAccount as String: key.rawValue,
            kSecValueData as String: data
        ]
        
        SecItemDelete(query as CFDictionary)
        SecItemAdd(query as CFDictionary, nil)
    }
    
    func getToken(for key: KeychainKey) -> String? {
        let query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrAccount as String: key.rawValue,
            kSecReturnData as String: true,
            kSecMatchLimit as String: kSecMatchLimitOne
        ]
        
        var dataTypeRef: AnyObject?
        let status = SecItemCopyMatching(query as CFDictionary, &dataTypeRef)
        
        if status == noErr, let data = dataTypeRef as? Data {
            return String(data: data, encoding: .utf8)
        }
        
        return nil
    }
    
    func deleteToken(for key: KeychainKey) {
        let query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrAccount as String: key.rawValue
        ]
        
        SecItemDelete(query as CFDictionary)
    }
    
    func deleteAll() {
        let query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword
        ]
        
        SecItemDelete(query as CFDictionary)
    }
}

enum KeychainKey: String {
    case accessToken = "access_token"
    case refreshToken = "refresh_token"
}
```

## 🔔 Push Notifications

### Notification Service
```swift
// Services/NotificationService.swift
import Foundation
import UserNotifications
import UIKit

class NotificationService: NSObject, ObservableObject {
    static let shared = NotificationService()
    
    @Published var authorizationStatus: UNAuthorizationStatus = .notDetermined
    
    override init() {
        super.init()
        UNUserNotificationCenter.current().delegate = self
    }
    
    func requestPermission() async -> Bool {
        do {
            let granted = try await UNUserNotificationCenter.current().requestAuthorization(
                options: [.alert, .badge, .sound]
            )
            
            await MainActor.run {
                self.authorizationStatus = granted ? .authorized : .denied
            }
            
            return granted
        } catch {
            print("Notification permission error: \(error)")
            return false
        }
    }
    
    func scheduleNotification(
        title: String,
        body: String,
        userInfo: [String: Any] = [:]
    ) {
        let content = UNMutableNotificationContent()
        content.title = title
        content.body = body
        content.sound = .default
        content.userInfo = userInfo
        
        let request = UNNotificationRequest(
            identifier: UUID().uuidString,
            content: content,
            trigger: nil
        )
        
        UNUserNotificationCenter.current().add(request)
    }
    
    func handleRemoteNotification(_ userInfo: [AnyHashable: Any]) {
        // Handle push notification data
        if let messageData = userInfo["message"] as? [String: Any] {
            // Process message notification
        }
    }
}

extension NotificationService: UNUserNotificationCenterDelegate {
    func userNotificationCenter(
        _ center: UNUserNotificationCenter,
        willPresent notification: UNNotification,
        withCompletionHandler completionHandler: @escaping (UNNotificationPresentationOptions) -> Void
    ) {
        // Show notification even when app is in foreground
        completionHandler([.banner, .sound, .badge])
    }
    
    func userNotificationCenter(
        _ center: UNUserNotificationCenter,
        didReceive response: UNNotificationResponse,
        withCompletionHandler completionHandler: @escaping () -> Void
    ) {
        // Handle notification tap
        let userInfo = response.notification.request.content.userInfo
        
        if let channelID = userInfo["channel_id"] as? String {
            // Navigate to specific channel
            NotificationCenter.default.post(
                name: .navigateToChannel,
                object: nil,
                userInfo: ["channel_id": channelID]
            )
        }
        
        completionHandler()
    }
}

extension Notification.Name {
    static let navigateToChannel = Notification.Name("navigateToChannel")
}
```

## 💾 Local Storage

### Core Data Setup
```swift
// Services/CoreDataService.swift
import CoreData
import Foundation

class PersistenceController {
    static let shared = PersistenceController()
    
    let container: NSPersistentContainer
    
    init(inMemory: Bool = false) {
        container = NSPersistentContainer(name: "AfroChat")
        
        if inMemory {
            container.persistentStoreDescriptions.first?.url = URL(fileURLWithPath: "/dev/null")
        }
        
        container.loadPersistentStores { _, error in
            if let error = error as NSError? {
                fatalError("Core Data error: \(error), \(error.userInfo)")
            }
        }
        
        container.viewContext.automaticallyMergesChangesFromParent = true
    }
}

// Core Data Models
extension Message {
    @nonobjc public class func fetchRequest() -> NSFetchRequest<Message> {
        return NSFetchRequest<Message>(entityName: "Message")
    }
    
    @NSManaged public var id: UUID
    @NSManaged public var content: String
    @NSManaged public var channelID: UUID
    @NSManaged public var userID: UUID
    @NSManaged public var createdAt: Date
    @NSManaged public var editedAt: Date?
    @NSManaged public var threadID: UUID?
}
```

## 🧪 Testing

### Unit Tests
```swift
// Tests/AfroChatTests/ChatViewModelTests.swift
import XCTest
import Combine
@testable import AfroChat

class ChatViewModelTests: XCTestCase {
    var viewModel: ChatViewModel!
    var mockChatService: MockChatService!
    var mockWebSocketService: MockWebSocketService!
    
    override func setUp() {
        super.setUp()
        mockChatService = MockChatService()
        mockWebSocketService = MockWebSocketService()
        viewModel = ChatViewModel(
            chatService: mockChatService,
            websocketService: mockWebSocketService
        )
    }
    
    func testLoadMessages() async {
        // Given
        let channelID = UUID()
        let expectedMessages = [Message.mockMessage()]
        mockChatService.messages = expectedMessages
        
        // When
        await viewModel.loadMessages(for: channelID)
        
        // Then
        XCTAssertEqual(viewModel.messages.count, 1)
        XCTAssertFalse(viewModel.isLoading)
    }
    
    func testSendMessage() async {
        // Given
        let channelID = UUID()
        let content = "Test message"
        
        // When
        await viewModel.loadMessages(for: channelID)
        await viewModel.sendMessage(content: content)
        
        // Then
        XCTAssertTrue(mockChatService.sendMessageCalled)
        XCTAssertEqual(mockChatService.lastMessageContent, content)
    }
}

class MockChatService: ChatService {
    var messages: [Message] = []
    var sendMessageCalled = false
    var lastMessageContent: String?
    
    func getMessages(channelID: UUID) -> AnyPublisher<[Message], Error> {
        return Just(messages)
            .setFailureType(to: Error.self)
            .eraseToAnyPublisher()
    }
    
    func sendMessage(_ request: SendMessageRequest) -> AnyPublisher<Message, Error> {
        sendMessageCalled = true
        lastMessageContent = request.content
        
        return Just(Message.mockMessage())
            .setFailureType(to: Error.self)
            .eraseToAnyPublisher()
    }
}
```

### UI Tests
```swift
// Tests/AfroChatUITests/ChatUITests.swift
import XCTest

class ChatUITests: XCTestCase {
    var app: XCUIApplication!
    
    override func setUp() {
        super.setUp()
        app = XCUIApplication()
        app.launch()
    }
    
    func testSendMessage() {
        // Login
        let emailTextField = app.textFields["Email"]
        emailTextField.tap()
        emailTextField.typeText("test@example.com")
        
        let passwordTextField = app.secureTextFields["Password"]
        passwordTextField.tap()
        passwordTextField.typeText("password")
        
        app.buttons["Login"].tap()
        
        // Navigate to channel
        app.buttons["General"].tap()
        
        // Send message
        let messageInput = app.textFields["Message input"]
        messageInput.tap()
        messageInput.typeText("Hello everyone!")
        
        app.buttons["Send"].tap()
        
        // Verify message appears
        XCTAssertTrue(app.staticTexts["Hello everyone!"].exists)
    }
}
```

## 🚀 Build and Deployment

### Build Configuration
```swift
// Build settings in Xcode
// - Deployment Target: iOS 16.0
// - Swift Language Version: Swift 5.9
// - Build Configuration: Debug/Release
// - Code Signing: Automatic
```

### App Store Configuration
```swift
// App Store Connect settings
// - App Name: AfroChat
// - Bundle ID: com.afrochat.app
// - Version: 1.0.0
// - Build: 1
// - Category: Social Networking
// - Age Rating: 4+
```

### CI/CD Pipeline
```yaml
# .github/workflows/ios.yml
name: iOS Build

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build:
    runs-on: macos-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Setup Xcode
      uses: maxim-lobanov/setup-xcode@v1
      with:
        xcode-version: '15.0'
    
    - name: Install dependencies
      run: pod install
    
    - name: Build
      run: xcodebuild -workspace AfroChat.xcworkspace -scheme AfroChat -destination 'platform=iOS Simulator,name=iPhone 15' build
    
    - name: Test
      run: xcodebuild -workspace AfroChat.xcworkspace -scheme AfroChat -destination 'platform=iOS Simulator,name=iPhone 15' test
```

## 📊 Performance Optimization

### Memory Management
```swift
// ViewModels should be @MainActor for UI updates
@MainActor
class ChatViewModel: ObservableObject {
    // Automatic memory management with Combine
    private var cancellables = Set<AnyCancellable>()
    
    deinit {
        cancellables.removeAll()
    }
}
```

### Image Loading
```swift
// Utils/AsyncImageLoader.swift
import SwiftUI

struct AsyncImageLoader: View {
    let url: String?
    let placeholder: Image
    
    var body: some View {
        AsyncImage(url: URL(string: url ?? "")) { image in
            image
                .resizable()
                .aspectRatio(contentMode: .fill)
        } placeholder: {
            placeholder
        }
        .frame(width: 40, height: 40)
        .clipShape(Circle())
    }
}
```

## 🛡️ Security Best Practices

### Certificate Pinning
```swift
// Utils/NetworkSecurity.swift
import Foundation

class NetworkSecurityManager: NSObject, URLSessionDelegate {
    func urlSession(
        _ session: URLSession,
        didReceive challenge: URLAuthenticationChallenge,
        completionHandler: @escaping (URLSession.AuthChallengeDisposition, URLCredential?) -> Void
    ) {
        // Implement certificate pinning
        guard let serverTrust = challenge.protectionSpace.serverTrust else {
            completionHandler(.cancelAuthenticationChallenge, nil)
            return
        }
        
        // Validate certificate
        let credential = URLCredential(trust: serverTrust)
        completionHandler(.useCredential, credential)
    }
}
```

### Input Validation
```swift
// Utils/Validators.swift
import Foundation

struct InputValidator {
    static func validateEmail(_ email: String) -> Bool {
        let emailRegex = "[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,64}"
        let emailPredicate = NSPredicate(format: "SELF MATCHES %@", emailRegex)
        return emailPredicate.evaluate(with: email)
    }
    
    static func validatePassword(_ password: String) -> Bool {
        return password.count >= 8
    }
    
    static func sanitizeInput(_ input: String) -> String {
        return input.trimmingCharacters(in: .whitespacesAndNewlines)
    }
}
```

## 📚 Development Tools

### SwiftLint Configuration
```yaml
# .swiftlint.yml
disabled_rules:
  - trailing_whitespace
  - line_length

opt_in_rules:
  - empty_count
  - force_unwrapping

included:
  - AfroChat

excluded:
  - Pods
  - Tests
```

### Xcode Schemes
```swift
// Debug scheme configuration
// - Build Configuration: Debug
// - Run: Debug
// - Test: Debug
// - Profile: Release
// - Analyze: Debug
// - Archive: Release
```

## 🤝 Contributing

1. Follow Swift style guidelines
2. Write unit tests for new features
3. Update documentation
4. Ensure all tests pass
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
