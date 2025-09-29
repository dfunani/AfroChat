# AfroChat Android App

A native Android application built with Kotlin and Jetpack Compose, providing a modern Slack-like communication experience on Android devices.

## 🏗️ Architecture Overview

### Technology Stack
- **Kotlin** with **Jetpack Compose** for modern UI development
- **Flow** for reactive programming and data streams
- **Retrofit** for networking and API communication
- **Room** for local database persistence
- **Hilt** for dependency injection
- **Firebase Cloud Messaging** for push notifications
- **WebSocket** for real-time communication
- **Material Design 3** for modern UI components

### Project Structure
```
android/
├── app/
│   ├── src/main/java/com/afrochat/
│   │   ├── data/
│   │   │   ├── models/
│   │   │   │   ├── User.kt
│   │   │   │   ├── Workspace.kt
│   │   │   │   ├── Channel.kt
│   │   │   │   ├── Message.kt
│   │   │   │   ├── File.kt
│   │   │   │   └── Reaction.kt
│   │   │   ├── repository/
│   │   │   │   ├── AuthRepository.kt
│   │   │   │   ├── ChatRepository.kt
│   │   │   │   ├── WorkspaceRepository.kt
│   │   │   │   └── FileRepository.kt
│   │   │   ├── network/
│   │   │   │   ├── ApiService.kt
│   │   │   │   ├── WebSocketService.kt
│   │   │   │   ├── ApiClient.kt
│   │   │   │   └── NetworkModule.kt
│   │   │   └── local/
│   │   │       ├── database/
│   │   │       │   ├── AfroChatDatabase.kt
│   │   │       │   ├── entities/
│   │   │       │   │   ├── UserEntity.kt
│   │   │       │   │   ├── MessageEntity.kt
│   │   │       │   │   └── ChannelEntity.kt
│   │   │       │   └── dao/
│   │   │       │       ├── UserDao.kt
│   │   │       │       ├── MessageDao.kt
│   │   │       │       └── ChannelDao.kt
│   │   │       └── preferences/
│   │   │           └── AuthPreferences.kt
│   │   ├── domain/
│   │   │   ├── usecases/
│   │   │   │   ├── LoginUseCase.kt
│   │   │   │   ├── SendMessageUseCase.kt
│   │   │   │   ├── GetMessagesUseCase.kt
│   │   │   │   └── UploadFileUseCase.kt
│   │   │   └── repository/
│   │   │       ├── AuthRepository.kt
│   │   │       ├── ChatRepository.kt
│   │   │       └── WorkspaceRepository.kt
│   │   ├── presentation/
│   │   │   ├── viewmodels/
│   │   │   │   ├── AuthViewModel.kt
│   │   │   │   ├── ChatViewModel.kt
│   │   │   │   ├── WorkspaceViewModel.kt
│   │   │   │   └── ProfileViewModel.kt
│   │   │   ├── screens/
│   │   │   │   ├── auth/
│   │   │   │   │   ├── LoginScreen.kt
│   │   │   │   │   ├── RegisterScreen.kt
│   │   │   │   │   └── ForgotPasswordScreen.kt
│   │   │   │   ├── workspace/
│   │   │   │   │   ├── WorkspaceScreen.kt
│   │   │   │   │   ├── ChannelListScreen.kt
│   │   │   │   │   └── UserListScreen.kt
│   │   │   │   ├── chat/
│   │   │   │   │   ├── ChatScreen.kt
│   │   │   │   │   ├── MessageListScreen.kt
│   │   │   │   │   └── ThreadScreen.kt
│   │   │   │   ├── video/
│   │   │   │   │   ├── VideoCallScreen.kt
│   │   │   │   │   ├── ParticipantGridScreen.kt
│   │   │   │   │   ├── VideoCallControls.kt
│   │   │   │   │   └── ScreenShareScreen.kt
│   │   │   │   └── profile/
│   │   │   │       ├── ProfileScreen.kt
│   │   │   │       └── SettingsScreen.kt
│   │   │   └── components/
│   │   │       ├── MessageItem.kt
│   │   │       ├── ChannelItem.kt
│   │   │       ├── UserItem.kt
│   │   │       ├── MessageInput.kt
│   │   │       └── TypingIndicator.kt
│   │   ├── di/
│   │   │   ├── AppModule.kt
│   │   │   ├── NetworkModule.kt
│   │   │   ├── DatabaseModule.kt
│   │   │   └── ViewModelModule.kt
│   │   ├── utils/
│   │   │   ├── Constants.kt
│   │   │   ├── Extensions.kt
│   │   │   ├── Formatters.kt
│   │   │   └── Validators.kt
│   │   └── AfroChatApplication.kt
│   ├── src/main/res/
│   │   ├── values/
│   │   │   ├── strings.xml
│   │   │   ├── colors.xml
│   │   │   ├── themes.xml
│   │   │   └── dimens.xml
│   │   ├── layout/
│   │   └── drawable/
│   ├── src/main/AndroidManifest.xml
│   └── build.gradle
├── build.gradle
├── settings.gradle
├── gradle.properties
└── README.md
```

## 🚀 Getting Started

### Prerequisites
- Android Studio Hedgehog (2023.1.1) or later
- Android SDK 34
- Kotlin 1.9.0+
- Backend API running on localhost:8080

### Installation

1. **Clone and setup**
```bash
cd frontend/android
# Open in Android Studio
```

2. **Configure Environment**
```kotlin
// app/src/main/java/com/afrochat/utils/Constants.kt
object Constants {
    const val API_BASE_URL = "http://10.0.2.2:8080/api/v1" // Use 10.0.2.2 for Android emulator
    const val WEBSOCKET_URL = "ws://10.0.2.2:8080/api/v1/ws"
    const val APP_NAME = "AfroChat"
    const val APP_VERSION = "1.0.0"
}
```

3. **Build and Run**
```bash
# Build debug APK
./gradlew assembleDebug

# Run on device/emulator
./gradlew installDebug

# Run tests
./gradlew test
```

## 🔧 Configuration

### Build Configuration
```kotlin
// app/build.gradle
android {
    compileSdk 34
    
    defaultConfig {
        applicationId "com.afrochat.app"
        minSdk 24
        targetSdk 34
        versionCode 1
        versionName "1.0.0"
        
        testInstrumentationRunner "androidx.test.runner.AndroidJUnitRunner"
    }
    
    buildTypes {
        debug {
            buildConfigField "String", "API_BASE_URL", "\"http://10.0.2.2:8080/api/v1\""
            buildConfigField "String", "WEBSOCKET_URL", "\"ws://10.0.2.2:8080/api/v1/ws\""
        }
        release {
            buildConfigField "String", "API_BASE_URL", "\"https://api.afrochat.com/v1\""
            buildConfigField "String", "WEBSOCKET_URL", "\"wss://api.afrochat.com/v1/ws\""
            minifyEnabled true
            proguardFiles getDefaultProguardFile('proguard-android-optimize.txt'), 'proguard-rules.pro'
        }
    }
    
    compileOptions {
        sourceCompatibility JavaVersion.VERSION_1_8
        targetCompatibility JavaVersion.VERSION_1_8
    }
    
    kotlinOptions {
        jvmTarget = '1.8'
    }
    
    buildFeatures {
        compose true
        buildConfig true
    }
    
    composeOptions {
        kotlinCompilerExtensionVersion '1.5.4'
    }
}
```

### Dependencies
```kotlin
// app/build.gradle
dependencies {
    // Core Android
    implementation 'androidx.core:core-ktx:1.12.0'
    implementation 'androidx.lifecycle:lifecycle-runtime-ktx:2.7.0'
    implementation 'androidx.activity:activity-compose:1.8.2'
    
    // Compose
    implementation platform('androidx.compose:compose-bom:2023.10.01')
    implementation 'androidx.compose.ui:ui'
    implementation 'androidx.compose.ui:ui-graphics'
    implementation 'androidx.compose.ui:ui-tooling-preview'
    implementation 'androidx.compose.material3:material3'
    implementation 'androidx.compose.material:material-icons-extended'
    
    // Navigation
    implementation 'androidx.navigation:navigation-compose:2.7.5'
    
    // ViewModel
    implementation 'androidx.lifecycle:lifecycle-viewmodel-compose:2.7.0'
    implementation 'androidx.lifecycle:lifecycle-runtime-compose:2.7.0'
    
    // Networking
    implementation 'com.squareup.retrofit2:retrofit:2.9.0'
    implementation 'com.squareup.retrofit2:converter-gson:2.9.0'
    implementation 'com.squareup.okhttp3:okhttp:4.12.0'
    implementation 'com.squareup.okhttp3:logging-interceptor:4.12.0'
    
    // WebSocket
    implementation 'com.squareup.okhttp3:okhttp:4.12.0'
    
    // Dependency Injection
    implementation 'com.google.dagger:hilt-android:2.48'
    implementation 'androidx.hilt:hilt-navigation-compose:1.1.0'
    kapt 'com.google.dagger:hilt-compiler:2.48'
    
    // Database
    implementation 'androidx.room:room-runtime:2.6.1'
    implementation 'androidx.room:room-ktx:2.6.1'
    kapt 'androidx.room:room-compiler:2.6.1'
    
    // Preferences
    implementation 'androidx.datastore:datastore-preferences:1.0.0'
    
    // Image Loading
    implementation 'io.coil-kt:coil-compose:2.5.0'
    
    // Firebase
    implementation platform('com.google.firebase:firebase-bom:32.7.0')
    implementation 'com.google.firebase:firebase-messaging-ktx'
    implementation 'com.google.firebase:firebase-analytics-ktx'
    
    // WebRTC
    implementation 'org.webrtc:google-webrtc:1.0.32006'
    implementation 'io.getstream:stream-webrtc-android:1.0.0'
    
    // Testing
    testImplementation 'junit:junit:4.13.2'
    testImplementation 'org.mockito:mockito-core:5.7.0'
    testImplementation 'org.jetbrains.kotlinx:kotlinx-coroutines-test:1.7.3'
    androidTestImplementation 'androidx.test.ext:junit:1.1.5'
    androidTestImplementation 'androidx.test.espresso:espresso-core:3.5.1'
    androidTestImplementation platform('androidx.compose:compose-bom:2023.10.01')
    androidTestImplementation 'androidx.compose.ui:ui-test-junit4'
    debugImplementation 'androidx.compose.ui:ui-tooling'
    debugImplementation 'androidx.compose.ui:ui-test-manifest'
}
```

## 📱 Data Models

### Core Models
```kotlin
// data/models/User.kt
import com.google.gson.annotations.SerializedName

data class User(
    @SerializedName("id")
    val id: String,
    
    @SerializedName("email")
    val email: String,
    
    @SerializedName("username")
    val username: String,
    
    @SerializedName("display_name")
    val displayName: String,
    
    @SerializedName("avatar_url")
    val avatarUrl: String?,
    
    @SerializedName("status")
    val status: UserStatus,
    
    @SerializedName("created_at")
    val createdAt: String,
    
    @SerializedName("updated_at")
    val updatedAt: String
)

enum class UserStatus(val value: String) {
    ONLINE("online"),
    OFFLINE("offline"),
    AWAY("away"),
    BUSY("busy")
}
```

### Message Model
```kotlin
// data/models/Message.kt
import com.google.gson.annotations.SerializedName

data class Message(
    @SerializedName("id")
    val id: String,
    
    @SerializedName("content")
    val content: String,
    
    @SerializedName("channel_id")
    val channelId: String,
    
    @SerializedName("user_id")
    val userId: String,
    
    @SerializedName("user")
    val user: User,
    
    @SerializedName("thread_id")
    val threadId: String?,
    
    @SerializedName("edited_at")
    val editedAt: String?,
    
    @SerializedName("reactions")
    val reactions: List<Reaction>,
    
    @SerializedName("created_at")
    val createdAt: String
)

data class Reaction(
    val emoji: String,
    val count: Int,
    val users: List<User>
)

data class SendMessageRequest(
    val content: String,
    val channelId: String,
    val threadId: String? = null
)
```

## 🔄 State Management

### ViewModels with StateFlow
```kotlin
// presentation/viewmodels/ChatViewModel.kt
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class ChatViewModel @Inject constructor(
    private val chatRepository: ChatRepository,
    private val websocketService: WebSocketService
) : ViewModel() {
    
    private val _uiState = MutableStateFlow(ChatUiState())
    val uiState: StateFlow<ChatUiState> = _uiState.asStateFlow()
    
    private val _messages = MutableStateFlow<List<Message>>(emptyList())
    val messages: StateFlow<List<Message>> = _messages.asStateFlow()
    
    private val _typingUsers = MutableStateFlow<List<User>>(emptyList())
    val typingUsers: StateFlow<List<User>> = _typingUsers.asStateFlow()
    
    fun loadMessages(channelId: String) {
        viewModelScope.launch {
            _uiState.value = _uiState.value.copy(isLoading = true)
            
            chatRepository.getMessages(channelId)
                .catch { error ->
                    _uiState.value = _uiState.value.copy(
                        isLoading = false,
                        error = error.message
                    )
                }
                .collect { messages ->
                    _messages.value = messages
                    _uiState.value = _uiState.value.copy(isLoading = false)
                }
        }
    }
    
    fun sendMessage(content: String, channelId: String) {
        viewModelScope.launch {
            try {
                chatRepository.sendMessage(SendMessageRequest(content, channelId))
            } catch (e: Exception) {
                _uiState.value = _uiState.value.copy(error = e.message)
            }
        }
    }
    
    fun startTyping(channelId: String) {
        websocketService.sendTyping(channelId)
    }
    
    fun stopTyping(channelId: String) {
        websocketService.stopTyping(channelId)
    }
    
    fun clearError() {
        _uiState.value = _uiState.value.copy(error = null)
    }
}

data class ChatUiState(
    val isLoading: Boolean = false,
    val error: String? = null
)
```

## 🌐 Networking

### Retrofit API Service
```kotlin
// data/network/ApiService.kt
import retrofit2.Response
import retrofit2.http.*

interface ApiService {
    @POST("auth/login")
    suspend fun login(@Body request: LoginRequest): Response<AuthResponse>
    
    @POST("auth/register")
    suspend fun register(@Body request: RegisterRequest): Response<AuthResponse>
    
    @POST("auth/refresh")
    suspend fun refresh(@Body request: RefreshRequest): Response<AuthResponse>
    
    @GET("users/profile")
    suspend fun getProfile(): Response<User>
    
    @PUT("users/profile")
    suspend fun updateProfile(@Body user: UpdateUserRequest): Response<User>
    
    @GET("workspaces")
    suspend fun getWorkspaces(): Response<List<Workspace>>
    
    @POST("workspaces")
    suspend fun createWorkspace(@Body request: CreateWorkspaceRequest): Response<Workspace>
    
    @GET("workspaces/{workspaceId}/channels")
    suspend fun getChannels(@Path("workspaceId") workspaceId: String): Response<List<Channel>>
    
    @POST("workspaces/{workspaceId}/channels")
    suspend fun createChannel(
        @Path("workspaceId") workspaceId: String,
        @Body request: CreateChannelRequest
    ): Response<Channel>
    
    @GET("channels/{channelId}/messages")
    suspend fun getMessages(
        @Path("channelId") channelId: String,
        @Query("limit") limit: Int = 50,
        @Query("before") before: String? = null,
        @Query("after") after: String? = null
    ): Response<MessageResponse>
    
    @POST("channels/{channelId}/messages")
    suspend fun sendMessage(
        @Path("channelId") channelId: String,
        @Body request: SendMessageRequest
    ): Response<Message>
    
    @POST("files/upload")
    suspend fun uploadFile(@Body file: MultipartBody.Part): Response<File>
}

data class MessageResponse(
    val messages: List<Message>,
    val hasMore: Boolean,
    val nextCursor: String?
)
```

### WebSocket Service
```kotlin
// data/network/WebSocketService.kt
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.*
import okhttp3.*
import okio.ByteString
import java.util.concurrent.TimeUnit
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class WebSocketService @Inject constructor(
    private val okHttpClient: OkHttpClient
) {
    private var webSocket: WebSocket? = null
    private val _isConnected = MutableStateFlow(false)
    val isConnected: StateFlow<Boolean> = _isConnected.asStateFlow()
    
    private val _messageFlow = MutableSharedFlow<Message>()
    val messageFlow: SharedFlow<Message> = _messageFlow.asSharedFlow()
    
    private val _typingUsersFlow = MutableSharedFlow<List<User>>()
    val typingUsersFlow: SharedFlow<List<User>> = _typingUsersFlow.asSharedFlow()
    
    fun connect(token: String) {
        val request = Request.Builder()
            .url("${Constants.WEBSOCKET_URL}?token=$token")
            .build()
        
        webSocket = okHttpClient.newWebSocket(request, object : WebSocketListener() {
            override fun onOpen(webSocket: WebSocket, response: Response) {
                _isConnected.value = true
            }
            
            override fun onMessage(webSocket: WebSocket, text: String) {
                handleMessage(text)
            }
            
            override fun onMessage(webSocket: WebSocket, bytes: ByteString) {
                handleMessage(bytes.utf8())
            }
            
            override fun onClosing(webSocket: WebSocket, code: Int, reason: String) {
                _isConnected.value = false
            }
            
            override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
                _isConnected.value = false
            }
        })
    }
    
    fun disconnect() {
        webSocket?.close(1000, "Normal closure")
        webSocket = null
        _isConnected.value = false
    }
    
    fun joinChannel(channelId: String) {
        sendMessage(WebSocketMessage("join_channel", mapOf("channel_id" to channelId)))
    }
    
    fun leaveChannel(channelId: String) {
        sendMessage(WebSocketMessage("leave_channel", mapOf("channel_id" to channelId)))
    }
    
    fun sendMessage(request: SendMessageRequest) {
        sendMessage(WebSocketMessage("send_message", request))
    }
    
    fun sendTyping(channelId: String) {
        sendMessage(WebSocketMessage("typing", mapOf("channel_id" to channelId)))
    }
    
    fun stopTyping(channelId: String) {
        sendMessage(WebSocketMessage("stop_typing", mapOf("channel_id" to channelId)))
    }
    
    private fun sendMessage(message: WebSocketMessage) {
        val json = gson.toJson(message)
        webSocket?.send(json)
    }
    
    private fun handleMessage(text: String) {
        try {
            val message = gson.fromJson(text, WebSocketMessage::class.java)
            when (message.type) {
                "message_sent" -> {
                    val msg = gson.fromJson(gson.toJson(message.data), Message::class.java)
                    _messageFlow.tryEmit(msg)
                }
                "user_typing" -> {
                    // Handle typing indicator
                }
            }
        } catch (e: Exception) {
            // Handle parsing error
        }
    }
}

data class WebSocketMessage(
    val type: String,
    val data: Any
)
```

## 💾 Local Storage

### Room Database
```kotlin
// data/local/database/AfroChatDatabase.kt
import androidx.room.Database
import androidx.room.Room
import androidx.room.RoomDatabase
import android.content.Context

@Database(
    entities = [UserEntity::class, MessageEntity::class, ChannelEntity::class],
    version = 1,
    exportSchema = false
)
abstract class AfroChatDatabase : RoomDatabase() {
    abstract fun userDao(): UserDao
    abstract fun messageDao(): MessageDao
    abstract fun channelDao(): ChannelDao
    
    companion object {
        @Volatile
        private var INSTANCE: AfroChatDatabase? = null
        
        fun getDatabase(context: Context): AfroChatDatabase {
            return INSTANCE ?: synchronized(this) {
                val instance = Room.databaseBuilder(
                    context.applicationContext,
                    AfroChatDatabase::class.java,
                    "afrochat_database"
                ).build()
                INSTANCE = instance
                instance
            }
        }
    }
}
```

### Entity Models
```kotlin
// data/local/database/entities/MessageEntity.kt
import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "messages")
data class MessageEntity(
    @PrimaryKey
    val id: String,
    val content: String,
    val channelId: String,
    val userId: String,
    val threadId: String?,
    val editedAt: String?,
    val createdAt: String
)

// data/local/database/dao/MessageDao.kt
import androidx.room.*
import kotlinx.coroutines.flow.Flow

@Dao
interface MessageDao {
    @Query("SELECT * FROM messages WHERE channelId = :channelId ORDER BY createdAt DESC")
    fun getMessagesByChannel(channelId: String): Flow<List<MessageEntity>>
    
    @Query("SELECT * FROM messages WHERE channelId = :channelId ORDER BY createdAt DESC LIMIT :limit")
    suspend fun getMessagesByChannelLimit(channelId: String, limit: Int): List<MessageEntity>
    
    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertMessage(message: MessageEntity)
    
    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertMessages(messages: List<MessageEntity>)
    
    @Update
    suspend fun updateMessage(message: MessageEntity)
    
    @Delete
    suspend fun deleteMessage(message: MessageEntity)
}
```

## 📹 Video Calling Implementation

### WebRTC Service
```kotlin
// services/WebRTCService.kt
import org.webrtc.*
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class WebRTCService @Inject constructor() {
    private var peerConnectionFactory: PeerConnectionFactory? = null
    private val peerConnections = mutableMapOf<String, PeerConnection>()
    private var localVideoSource: VideoSource? = null
    private var localVideoTrack: VideoTrack? = null
    private var localAudioTrack: AudioTrack? = null
    
    fun initialize() {
        val initializationOptions = PeerConnectionFactory.InitializationOptions.builder()
            .setEnableInternalTracer(true)
            .setFieldTrials("")
            .createInitializationOptions()
        
        PeerConnectionFactory.initialize(initializationOptions)
        
        val options = PeerConnectionFactory.Options()
        peerConnectionFactory = PeerConnectionFactory.builder()
            .setOptions(options)
            .createPeerConnectionFactory()
    }
    
    fun createLocalVideoTrack(): VideoTrack? {
        val videoCapturer = createCameraCapturer()
        localVideoSource = peerConnectionFactory?.createVideoSource(false)
        localVideoTrack = peerConnectionFactory?.createVideoTrack("video_track", localVideoSource)
        
        videoCapturer?.initialize(
            SurfaceTextureHelper.create("CaptureThread", null),
            localVideoSource?.capturerObserver
        )
        videoCapturer?.startCapture(1280, 720, 30)
        
        return localVideoTrack
    }
    
    fun createLocalAudioTrack(): AudioTrack? {
        val audioSource = peerConnectionFactory?.createAudioSource(MediaConstraints())
        localAudioTrack = peerConnectionFactory?.createAudioTrack("audio_track", audioSource)
        return localAudioTrack
    }
    
    fun createPeerConnection(userId: String): PeerConnection? {
        val rtcConfig = PeerConnection.RTCConfiguration(
            listOf(
                PeerConnection.IceServer.builder("stun:stun.l.google.com:19302").createIceServer(),
                PeerConnection.IceServer.builder("stun:stun1.l.google.com:19302").createIceServer()
            )
        )
        
        val peerConnection = peerConnectionFactory?.createPeerConnection(rtcConfig, object : PeerConnection.Observer {
            override fun onSignalingChange(signalingState: PeerConnection.SignalingState?) {}
            override fun onIceConnectionChange(iceConnectionState: PeerConnection.IceConnectionState?) {}
            override fun onIceConnectionReceivingChange(receiving: Boolean) {}
            override fun onIceGatheringChange(iceGatheringState: PeerConnection.IceGatheringState?) {}
            override fun onIceCandidate(iceCandidate: IceCandidate?) {
                // Send ICE candidate via WebSocket
                sendICECandidate(userId, iceCandidate)
            }
            override fun onIceCandidatesRemoved(candidates: Array<out IceCandidate>?) {}
            override fun onAddStream(mediaStream: MediaStream?) {
                // Handle remote stream
                handleRemoteStream(userId, mediaStream)
            }
            override fun onRemoveStream(mediaStream: MediaStream?) {}
            override fun onDataChannel(dataChannel: DataChannel?) {}
            override fun onRenegotiationNeeded() {}
            override fun onAddTrack(rtpReceiver: RtpReceiver?, streams: Array<out MediaStream>?) {}
        })
        
        peerConnections[userId] = peerConnection
        return peerConnection
    }
    
    fun createOffer(userId: String, callback: (SessionDescription?) -> Unit) {
        val peerConnection = createPeerConnection(userId)
        val mediaConstraints = MediaConstraints()
        
        peerConnection?.createOffer(object : SdpObserver {
            override fun onCreateSuccess(sessionDescription: SessionDescription?) {
                peerConnection.setLocalDescription(object : SdpObserver {
                    override fun onCreateSuccess(p0: SessionDescription?) {}
                    override fun onSetSuccess() {
                        callback(sessionDescription)
                    }
                    override fun onCreateFailure(p0: String?) {}
                    override fun onSetFailure(p0: String?) {}
                }, sessionDescription)
            }
            override fun onSetSuccess() {}
            override fun onCreateFailure(p0: String?) {}
            override fun onSetFailure(p0: String?) {}
        }, mediaConstraints)
    }
    
    fun handleOffer(userId: String, offer: SessionDescription) {
        val peerConnection = createPeerConnection(userId)
        peerConnection?.setRemoteDescription(object : SdpObserver {
            override fun onCreateSuccess(p0: SessionDescription?) {}
            override fun onSetSuccess() {
                createAnswer(userId, peerConnection)
            }
            override fun onCreateFailure(p0: String?) {}
            override fun onSetFailure(p0: String?) {}
        }, offer)
    }
    
    private fun createAnswer(userId: String, peerConnection: PeerConnection) {
        val mediaConstraints = MediaConstraints()
        peerConnection.createAnswer(object : SdpObserver {
            override fun onCreateSuccess(sessionDescription: SessionDescription?) {
                peerConnection.setLocalDescription(object : SdpObserver {
                    override fun onCreateSuccess(p0: SessionDescription?) {}
                    override fun onSetSuccess() {
                        // Send answer via WebSocket
                        sendAnswer(userId, sessionDescription)
                    }
                    override fun onCreateFailure(p0: String?) {}
                    override fun onSetFailure(p0: String?) {}
                }, sessionDescription)
            }
            override fun onSetSuccess() {}
            override fun onCreateFailure(p0: String?) {}
            override fun onSetFailure(p0: String?) {}
        }, mediaConstraints)
    }
    
    private fun createCameraCapturer(): CameraVideoCapturer? {
        val cameraEnumerator = Camera2Enumerator()
        val deviceNames = cameraEnumerator.deviceNames
        
        for (deviceName in deviceNames) {
            if (cameraEnumerator.isFrontFacing(deviceName)) {
                return cameraEnumerator.createCapturer(deviceName, null)
            }
        }
        
        return null
    }
    
    private fun sendICECandidate(userId: String, candidate: IceCandidate?) {
        // Send via WebSocket
    }
    
    private fun sendAnswer(userId: String, answer: SessionDescription?) {
        // Send via WebSocket
    }
    
    private fun handleRemoteStream(userId: String, stream: MediaStream?) {
        // Handle remote video stream
    }
}
```

### Video Call ViewModel
```kotlin
// presentation/viewmodels/VideoCallViewModel.kt
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class VideoCallViewModel @Inject constructor(
    private val videoCallRepository: VideoCallRepository,
    private val webrtcService: WebRTCService
) : ViewModel() {
    
    private val _uiState = MutableStateFlow(VideoCallUiState())
    val uiState: StateFlow<VideoCallUiState> = _uiState.asStateFlow()
    
    private val _participants = MutableStateFlow<List<CallParticipant>>(emptyList())
    val participants: StateFlow<List<CallParticipant>> = _participants.asStateFlow()
    
    fun startCall(channelId: String, title: String) {
        viewModelScope.launch {
            _uiState.value = _uiState.value.copy(isLoading = true)
            
            try {
                val call = videoCallRepository.startCall(StartCallRequest(channelId, title))
                _uiState.value = _uiState.value.copy(
                    isLoading = false,
                    isInCall = true,
                    currentCall = call
                )
                
                // Initialize WebRTC
                webrtcService.initialize()
                webrtcService.createLocalVideoTrack()
                webrtcService.createLocalAudioTrack()
                
            } catch (e: Exception) {
                _uiState.value = _uiState.value.copy(
                    isLoading = false,
                    error = e.message
                )
            }
        }
    }
    
    fun joinCall(callId: String) {
        viewModelScope.launch {
            try {
                val call = videoCallRepository.joinCall(callId)
                _uiState.value = _uiState.value.copy(
                    isInCall = true,
                    currentCall = call
                )
                
                // Initialize WebRTC
                webrtcService.initialize()
                webrtcService.createLocalVideoTrack()
                webrtcService.createLocalAudioTrack()
                
            } catch (e: Exception) {
                _uiState.value = _uiState.value.copy(error = e.message)
            }
        }
    }
    
    fun leaveCall() {
        viewModelScope.launch {
            try {
                videoCallRepository.leaveCall()
                _uiState.value = _uiState.value.copy(
                    isInCall = false,
                    currentCall = null
                )
                _participants.value = emptyList()
            } catch (e: Exception) {
                _uiState.value = _uiState.value.copy(error = e.message)
            }
        }
    }
    
    fun toggleMute() {
        webrtcService.toggleMute()
        _uiState.value = _uiState.value.copy(isMuted = !_uiState.value.isMuted)
    }
    
    fun toggleVideo() {
        webrtcService.toggleVideo()
        _uiState.value = _uiState.value.copy(isVideoEnabled = !_uiState.value.isVideoEnabled)
    }
    
    fun toggleScreenShare() {
        viewModelScope.launch {
            try {
                if (_uiState.value.isScreenSharing) {
                    webrtcService.stopScreenShare()
                } else {
                    webrtcService.startScreenShare()
                }
                _uiState.value = _uiState.value.copy(
                    isScreenSharing = !_uiState.value.isScreenSharing
                )
            } catch (e: Exception) {
                _uiState.value = _uiState.value.copy(error = e.message)
            }
        }
    }
}

data class VideoCallUiState(
    val isLoading: Boolean = false,
    val isInCall: Boolean = false,
    val isMuted: Boolean = false,
    val isVideoEnabled: Boolean = true,
    val isScreenSharing: Boolean = false,
    val currentCall: VideoCall? = null,
    val error: String? = null
)
```

### Video Call Screen
```kotlin
// presentation/screens/video/VideoCallScreen.kt
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.grid.GridCells
import androidx.compose.foundation.lazy.grid.LazyVerticalGrid
import androidx.compose.foundation.lazy.grid.items
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.compose.collectAsStateWithLifecycle

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun VideoCallScreen(
    callId: String?,
    channelId: String,
    viewModel: VideoCallViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsStateWithLifecycle()
    val participants by viewModel.participants.collectAsStateWithLifecycle()
    
    LaunchedEffect(callId) {
        if (callId != null) {
            viewModel.joinCall(callId)
        }
    }
    
    Box(modifier = Modifier.fillMaxSize()) {
        // Participant grid
        LazyVerticalGrid(
            columns = GridCells.Fixed(2),
            modifier = Modifier.fillMaxSize(),
            contentPadding = PaddingValues(8.dp),
            verticalArrangement = Arrangement.spacedBy(8.dp),
            horizontalArrangement = Arrangement.spacedBy(8.dp)
        ) {
            items(participants) { participant ->
                ParticipantVideoItem(
                    participant = participant,
                    modifier = Modifier.aspectRatio(16f / 9f)
                )
            }
        }
        
        // Call controls
        VideoCallControls(
            isMuted = uiState.isMuted,
            isVideoEnabled = uiState.isVideoEnabled,
            isScreenSharing = uiState.isScreenSharing,
            onToggleMute = { viewModel.toggleMute() },
            onToggleVideo = { viewModel.toggleVideo() },
            onToggleScreenShare = { viewModel.toggleScreenShare() },
            onLeaveCall = { viewModel.leaveCall() },
            modifier = Modifier.align(Alignment.BottomCenter)
        )
    }
    
    if (uiState.isLoading) {
        Box(
            modifier = Modifier.fillMaxSize(),
            contentAlignment = Alignment.Center
        ) {
            CircularProgressIndicator()
        }
    }
    
    uiState.error?.let { error ->
        LaunchedEffect(error) {
            // Show error snackbar
        }
    }
}
```

## 🔔 Push Notifications

### Firebase Messaging
```kotlin
// services/NotificationService.kt
import com.google.firebase.messaging.FirebaseMessagingService
import com.google.firebase.messaging.RemoteMessage
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import javax.inject.Inject

@AndroidEntryPoint
class AfroChatMessagingService : FirebaseMessagingService() {
    
    @Inject
    lateinit var notificationManager: NotificationManager
    
    override fun onMessageReceived(remoteMessage: RemoteMessage) {
        super.onMessageReceived(remoteMessage)
        
        remoteMessage.data.let { data ->
            val title = data["title"] ?: "New Message"
            val body = data["body"] ?: "You have a new message"
            val channelId = data["channel_id"]
            val callId = data["call_id"]
            
            CoroutineScope(Dispatchers.Main).launch {
                if (callId != null) {
                    // Video call notification
                    notificationManager.showCallNotification(
                        title = title,
                        body = body,
                        callId = callId
                    )
                } else {
                    // Regular message notification
                    notificationManager.showNotification(
                        title = title,
                        body = body,
                        channelId = channelId
                    )
                }
            }
        }
    }
    
    override fun onNewToken(token: String) {
        super.onNewToken(token)
        // Send token to server
        CoroutineScope(Dispatchers.IO).launch {
            // Update token on server
        }
    }
}
```

### Notification Manager
```kotlin
// utils/NotificationManager.kt
import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.PendingIntent
import android.content.Context
import android.content.Intent
import android.os.Build
import androidx.core.app.NotificationCompat
import dagger.hilt.android.qualifiers.ApplicationContext
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class NotificationManager @Inject constructor(
    @ApplicationContext private val context: Context
) {
    private val notificationManager = context.getSystemService(Context.NOTIFICATION_SERVICE) as NotificationManager
    
    init {
        createNotificationChannels()
    }
    
    private fun createNotificationChannels() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            val messageChannel = NotificationChannel(
                CHANNEL_MESSAGES,
                "Messages",
                NotificationManager.IMPORTANCE_HIGH
            ).apply {
                description = "New message notifications"
                enableVibration(true)
                setShowBadge(true)
            }
            
            notificationManager.createNotificationChannel(messageChannel)
        }
    }
    
    suspend fun showNotification(
        title: String,
        body: String,
        channelId: String? = null
    ) {
        val intent = Intent(context, MainActivity::class.java).apply {
            flags = Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TASK
            putExtra("channel_id", channelId)
        }
        
        val pendingIntent = PendingIntent.getActivity(
            context,
            0,
            intent,
            PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE
        )
        
        val notification = NotificationCompat.Builder(context, CHANNEL_MESSAGES)
            .setSmallIcon(R.drawable.ic_notification)
            .setContentTitle(title)
            .setContentText(body)
            .setPriority(NotificationCompat.PRIORITY_HIGH)
            .setContentIntent(pendingIntent)
            .setAutoCancel(true)
            .build()
        
        notificationManager.notify(System.currentTimeMillis().toInt(), notification)
    }
    
    companion object {
        private const val CHANNEL_MESSAGES = "messages"
    }
}
```

## 🎨 UI Components

### Compose Screens
```kotlin
// presentation/screens/chat/ChatScreen.kt
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.compose.collectAsStateWithLifecycle

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ChatScreen(
    channelId: String,
    viewModel: ChatViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsStateWithLifecycle()
    val messages by viewModel.messages.collectAsStateWithLifecycle()
    val typingUsers by viewModel.typingUsers.collectAsStateWithLifecycle()
    
    LaunchedEffect(channelId) {
        viewModel.loadMessages(channelId)
    }
    
    Column(
        modifier = Modifier.fillMaxSize()
    ) {
        // Messages list
        LazyColumn(
            modifier = Modifier.weight(1f),
            contentPadding = PaddingValues(16.dp),
            verticalArrangement = Arrangement.spacedBy(8.dp)
        ) {
            items(messages) { message ->
                MessageItem(
                    message = message,
                    modifier = Modifier.fillMaxWidth()
                )
            }
            
            if (typingUsers.isNotEmpty()) {
                item {
                    TypingIndicator(users = typingUsers)
                }
            }
        }
        
        // Message input
        MessageInput(
            onSendMessage = { content ->
                viewModel.sendMessage(content, channelId)
            },
            onTyping = {
                viewModel.startTyping(channelId)
            },
            onStopTyping = {
                viewModel.stopTyping(channelId)
            }
        )
    }
    
    if (uiState.isLoading) {
        Box(
            modifier = Modifier.fillMaxSize(),
            contentAlignment = Alignment.Center
        ) {
            CircularProgressIndicator()
        }
    }
    
    uiState.error?.let { error ->
        LaunchedEffect(error) {
            // Show error snackbar
        }
    }
}
```

### Message Component
```kotlin
// presentation/components/MessageItem.kt
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage

@Composable
fun MessageItem(
    message: Message,
    modifier: Modifier = Modifier
) {
    Row(
        modifier = modifier.padding(vertical = 4.dp),
        verticalAlignment = Alignment.Top
    ) {
        // User avatar
        AsyncImage(
            model = message.user.avatarUrl,
            contentDescription = "User avatar",
            modifier = Modifier
                .size(40.dp)
                .clip(CircleShape),
            contentScale = ContentScale.Crop
        )
        
        Spacer(modifier = Modifier.width(12.dp))
        
        Column(
            modifier = Modifier.weight(1f)
        ) {
            // User name and timestamp
            Row(
                verticalAlignment = Alignment.CenterVertically
            ) {
                Text(
                    text = message.user.displayName,
                    style = MaterialTheme.typography.bodyMedium,
                    fontWeight = FontWeight.Bold
                )
                
                Spacer(modifier = Modifier.width(8.dp))
                
                Text(
                    text = formatTimestamp(message.createdAt),
                    style = MaterialTheme.typography.bodySmall,
                    color = MaterialTheme.colorScheme.onSurfaceVariant
                )
            }
            
            Spacer(modifier = Modifier.height(4.dp))
            
            // Message content
            Text(
                text = message.content,
                style = MaterialTheme.typography.bodyLarge
            )
            
            // Reactions
            if (message.reactions.isNotEmpty()) {
                Spacer(modifier = Modifier.height(8.dp))
                MessageReactions(reactions = message.reactions)
            }
        }
    }
}
```

## 🧪 Testing

### Unit Tests
```kotlin
// test/java/com/afrochat/ChatViewModelTest.kt
import kotlinx.coroutines.flow.flowOf
import kotlinx.coroutines.test.runTest
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import org.mockito.Mock
import org.mockito.junit.MockitoJUnitRunner
import org.mockito.kotlin.whenever
import kotlin.test.assertEquals
import kotlin.test.assertFalse
import kotlin.test.assertTrue

@RunWith(MockitoJUnitRunner::class)
class ChatViewModelTest {
    
    @Mock
    private lateinit var chatRepository: ChatRepository
    
    @Mock
    private lateinit var websocketService: WebSocketService
    
    private lateinit var viewModel: ChatViewModel
    
    @Before
    fun setup() {
        viewModel = ChatViewModel(chatRepository, websocketService)
    }
    
    @Test
    fun `loadMessages should update messages when successful`() = runTest {
        // Given
        val channelId = "channel1"
        val messages = listOf(
            Message(
                id = "1",
                content = "Hello",
                channelId = channelId,
                userId = "user1",
                user = User("user1", "test@example.com", "testuser", "Test User", null, UserStatus.ONLINE, "", ""),
                threadId = null,
                editedAt = null,
                reactions = emptyList(),
                createdAt = "2023-01-01T00:00:00Z"
            )
        )
        whenever(chatRepository.getMessages(channelId)).thenReturn(flowOf(messages))
        
        // When
        viewModel.loadMessages(channelId)
        
        // Then
        assertEquals(messages, viewModel.messages.value)
        assertFalse(viewModel.uiState.value.isLoading)
    }
    
    @Test
    fun `sendMessage should call repository`() = runTest {
        // Given
        val content = "Test message"
        val channelId = "channel1"
        
        // When
        viewModel.sendMessage(content, channelId)
        
        // Then
        // Verify repository was called
        // This would require a more sophisticated mock setup
    }
}
```

### UI Tests
```kotlin
// androidTest/java/com/afrochat/ChatScreenTest.kt
import androidx.compose.ui.test.*
import androidx.compose.ui.test.junit4.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

@RunWith(AndroidJUnit4::class)
class ChatScreenTest {
    
    @get:Rule
    val composeTestRule = createComposeRule()
    
    @Test
    fun chatScreen_displaysMessages() {
        composeTestRule.setContent {
            ChatScreen(channelId = "channel1")
        }
        
        // Verify message input is displayed
        composeTestRule.onNodeWithText("Type a message...").assertIsDisplayed()
        
        // Type and send message
        composeTestRule.onNodeWithText("Type a message...").performTextInput("Hello world!")
        composeTestRule.onNodeWithContentDescription("Send message").performClick()
        
        // Verify message appears
        composeTestRule.onNodeWithText("Hello world!").assertIsDisplayed()
    }
}
```

## 🚀 Build and Deployment

### Build Variants
```kotlin
// app/build.gradle
android {
    buildTypes {
        debug {
            applicationIdSuffix ".debug"
            versionNameSuffix "-debug"
            buildConfigField "boolean", "DEBUG_MODE", "true"
        }
        
        release {
            minifyEnabled true
            proguardFiles getDefaultProguardFile('proguard-android-optimize.txt'), 'proguard-rules.pro'
            buildConfigField "boolean", "DEBUG_MODE", "false"
        }
    }
    
    flavorDimensions "environment"
    productFlavors {
        dev {
            dimension "environment"
            applicationIdSuffix ".dev"
            versionNameSuffix "-dev"
            buildConfigField "String", "API_BASE_URL", "\"http://10.0.2.2:8080/api/v1\""
        }
        
        prod {
            dimension "environment"
            buildConfigField "String", "API_BASE_URL", "\"https://api.afrochat.com/v1\""
        }
    }
}
```

### CI/CD Pipeline
```yaml
# .github/workflows/android.yml
name: Android Build

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up JDK 17
      uses: actions/setup-java@v3
      with:
        java-version: '17'
        distribution: 'temurin'
    
    - name: Cache Gradle packages
      uses: actions/cache@v3
      with:
        path: |
          ~/.gradle/caches
          ~/.gradle/wrapper
        key: ${{ runner.os }}-gradle-${{ hashFiles('**/*.gradle*', '**/gradle-wrapper.properties') }}
        restore-keys: |
          ${{ runner.os }}-gradle-
    
    - name: Grant execute permission for gradlew
      run: chmod +x gradlew
    
    - name: Build with Gradle
      run: ./gradlew build
    
    - name: Run tests
      run: ./gradlew test
    
    - name: Upload APK
      uses: actions/upload-artifact@v3
      with:
        name: app-debug
        path: app/build/outputs/apk/debug/app-debug.apk
```

## 📊 Performance Optimization

### Image Loading
```kotlin
// utils/ImageLoader.kt
import coil.compose.AsyncImage
import coil.request.ImageRequest
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier

@Composable
fun UserAvatar(
    imageUrl: String?,
    modifier: Modifier = Modifier
) {
    AsyncImage(
        model = ImageRequest.Builder(LocalContext.current)
            .data(imageUrl)
            .crossfade(true)
            .build(),
        contentDescription = "User avatar",
        modifier = modifier,
        contentScale = ContentScale.Crop
    )
}
```

### Memory Management
```kotlin
// utils/MemoryManager.kt
import android.app.ActivityManager
import android.content.Context
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class MemoryManager @Inject constructor(
    private val context: Context
) {
    fun getAvailableMemory(): Long {
        val activityManager = context.getSystemService(Context.ACTIVITY_SERVICE) as ActivityManager
        val memoryInfo = ActivityManager.MemoryInfo()
        activityManager.getMemoryInfo(memoryInfo)
        return memoryInfo.availMem
    }
    
    fun isLowMemory(): Boolean {
        val activityManager = context.getSystemService(Context.ACTIVITY_SERVICE) as ActivityManager
        return activityManager.isLowRamDevice
    }
}
```

## 🛡️ Security

### Certificate Pinning
```kotlin
// network/CertificatePinner.kt
import okhttp3.CertificatePinner
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class CertificatePinner @Inject constructor() {
    fun create(): CertificatePinner {
        return CertificatePinner.Builder()
            .add("api.afrochat.com", "sha256/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=")
            .build()
    }
}
```

### Data Encryption
```kotlin
// utils/EncryptionManager.kt
import android.security.keystore.KeyGenParameterSpec
import android.security.keystore.KeyProperties
import java.security.KeyStore
import javax.crypto.Cipher
import javax.crypto.KeyGenerator
import javax.crypto.SecretKey
import javax.crypto.spec.GCMParameterSpec

class EncryptionManager {
    private val keyStore = KeyStore.getInstance("AndroidKeyStore")
    
    init {
        keyStore.load(null)
    }
    
    fun encrypt(data: String): String {
        val cipher = Cipher.getInstance("AES/GCM/NoPadding")
        val key = getOrCreateKey()
        cipher.init(Cipher.ENCRYPT_MODE, key)
        
        val iv = cipher.iv
        val encryptedBytes = cipher.doFinal(data.toByteArray())
        
        return Base64.encodeToString(encryptedBytes, Base64.DEFAULT)
    }
    
    private fun getOrCreateKey(): SecretKey {
        val keyGenerator = KeyGenerator.getInstance(KeyProperties.KEY_ALGORITHM_AES, "AndroidKeyStore")
        val keyGenParameterSpec = KeyGenParameterSpec.Builder(
            "AfroChatKey",
            KeyProperties.PURPOSE_ENCRYPT or KeyProperties.PURPOSE_DECRYPT
        )
            .setBlockModes(KeyProperties.BLOCK_MODE_GCM)
            .setEncryptionPaddings(KeyProperties.ENCRYPTION_PADDING_NONE)
            .build()
        
        keyGenerator.init(keyGenParameterSpec)
        return keyGenerator.generateKey()
    }
}
```

## 📚 Development Tools

### Ktlint Configuration
```kotlin
// build.gradle (project level)
plugins {
    id "org.jlleitschuh.gradle.ktlint" version "11.6.1"
}

ktlint {
    version = "0.50.0"
    debug = true
    verbose = true
    android = true
    outputToConsole = true
    outputColorName = "RED"
    ignoreFailures = false
    enableExperimentalRules = true
}
```

### Detekt Configuration
```yaml
# detekt.yml
config:
  validation: true
  warningsAsErrors: false
  checkExhaustiveness: true

complexity:
  ComplexMethod:
    threshold: 15
  LongParameterList:
    functionThreshold: 6
    constructorThreshold: 7
    ignoreDefaultParameters: false
    ignoreDataClasses: true
    ignoreAnnotatedParameter: []

style:
  ForbiddenComment:
    values: ['FIXME:', 'STOPSHIP:', 'TODO:']
  MaxLineLength:
    maxLineLength: 120
    ignoreBackTicks: true
```

## 🤝 Contributing

1. Follow Kotlin coding standards
2. Write unit tests for new features
3. Update documentation
4. Ensure all tests pass
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
