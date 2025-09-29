# AfroChat Web Frontend

A modern, real-time web application built with React 18, TypeScript, and Vite, providing a Slack-like communication experience in the browser.

## 🏗️ Architecture Overview

### Technology Stack
- **React 18** with TypeScript for type safety
- **Vite** for fast development and optimized builds
- **Tailwind CSS** for utility-first styling
- **Zustand** for lightweight state management
- **React Query** for server state and caching
- **Socket.io-client** for real-time communication
- **React Router** for client-side routing
- **React Hook Form** for form handling
- **Zod** for runtime validation

### Project Structure
```
web/
├── public/
│   ├── favicon.ico
│   ├── manifest.json
│   └── sw.js                    # Service worker
├── src/
│   ├── components/
│   │   ├── ui/                  # Base UI components
│   │   │   ├── Button.tsx
│   │   │   ├── Input.tsx
│   │   │   ├── Modal.tsx
│   │   │   ├── Avatar.tsx
│   │   │   ├── Badge.tsx
│   │   │   ├── Dropdown.tsx
│   │   │   └── index.ts
│   │   ├── chat/                # Chat-specific components
│   │   │   ├── MessageList.tsx
│   │   │   ├── MessageItem.tsx
│   │   │   ├── MessageInput.tsx
│   │   │   ├── TypingIndicator.tsx
│   │   │   ├── MessageReactions.tsx
│   │   │   └── ThreadView.tsx
│   │   ├── video/               # Video calling components
│   │   │   ├── VideoCallModal.tsx
│   │   │   ├── VideoCallControls.tsx
│   │   │   ├── ParticipantGrid.tsx
│   │   │   ├── ParticipantVideo.tsx
│   │   │   ├── ScreenShare.tsx
│   │   │   └── CallSettings.tsx
│   │   ├── workspace/           # Workspace components
│   │   │   ├── WorkspaceSidebar.tsx
│   │   │   ├── ChannelList.tsx
│   │   │   ├── ChannelItem.tsx
│   │   │   ├── UserList.tsx
│   │   │   ├── UserItem.tsx
│   │   │   └── WorkspaceHeader.tsx
│   │   ├── auth/                # Authentication components
│   │   │   ├── LoginForm.tsx
│   │   │   ├── RegisterForm.tsx
│   │   │   ├── ForgotPasswordForm.tsx
│   │   │   └── AuthLayout.tsx
│   │   └── common/              # Shared components
│   │       ├── Layout.tsx
│   │       ├── Header.tsx
│   │       ├── Sidebar.tsx
│   │       ├── LoadingSpinner.tsx
│   │       ├── ErrorBoundary.tsx
│   │       └── Toast.tsx
│   ├── pages/
│   │   ├── auth/
│   │   │   ├── LoginPage.tsx
│   │   │   ├── RegisterPage.tsx
│   │   │   └── ForgotPasswordPage.tsx
│   │   ├── workspace/
│   │   │   ├── WorkspacePage.tsx
│   │   │   ├── ChannelPage.tsx
│   │   │   ├── DirectMessagePage.tsx
│   │   │   └── SettingsPage.tsx
│   │   ├── profile/
│   │   │   ├── ProfilePage.tsx
│   │   │   └── EditProfilePage.tsx
│   │   └── error/
│   │       ├── NotFoundPage.tsx
│   │       └── ErrorPage.tsx
│   ├── hooks/
│   │   ├── useAuth.ts
│   │   ├── useWebSocket.ts
│   │   ├── useChat.ts
│   │   ├── useWorkspace.ts
│   │   ├── useFileUpload.ts
│   │   ├── useNotifications.ts
│   │   ├── useLocalStorage.ts
│   │   ├── useDebounce.ts
│   │   ├── useVideoCall.ts
│   │   ├── useWebRTC.ts
│   │   └── useScreenShare.ts
│   ├── services/
│   │   ├── api.ts
│   │   ├── websocket.ts
│   │   ├── auth.ts
│   │   ├── chat.ts
│   │   ├── workspace.ts
│   │   ├── file.ts
│   │   ├── notification.ts
│   │   ├── videoCall.ts
│   │   └── webrtc.ts
│   ├── store/
│   │   ├── authStore.ts
│   │   ├── chatStore.ts
│   │   ├── workspaceStore.ts
│   │   ├── uiStore.ts
│   │   └── index.ts
│   ├── types/
│   │   ├── api.ts
│   │   ├── user.ts
│   │   ├── message.ts
│   │   ├── channel.ts
│   │   ├── workspace.ts
│   │   ├── websocket.ts
│   │   └── index.ts
│   ├── utils/
│   │   ├── constants.ts
│   │   ├── helpers.ts
│   │   ├── formatters.ts
│   │   ├── validators.ts
│   │   ├── storage.ts
│   │   └── api.ts
│   ├── styles/
│   │   ├── globals.css
│   │   ├── components.css
│   │   └── utilities.css
│   ├── App.tsx
│   ├── main.tsx
│   └── vite-env.d.ts
├── package.json
├── vite.config.ts
├── tailwind.config.js
├── tsconfig.json
├── .env.example
└── README.md
```

## 🚀 Getting Started

### Prerequisites
- Node.js 18+ and npm
- Backend API running on localhost:8080
- Redis and PostgreSQL (for backend)

### Installation

1. **Clone and setup**
```bash
cd frontend/web
npm install
```

2. **Environment Configuration**
```bash
cp .env.example .env
# Edit .env with your configuration
```

3. **Start Development Server**
```bash
npm run dev
```

4. **Build for Production**
```bash
npm run build
npm run preview
```

## 🔧 Configuration

### Environment Variables
```bash
# API Configuration
VITE_API_BASE_URL=http://localhost:8080/api/v1
VITE_WS_URL=ws://localhost:8080/api/v1/ws

# App Configuration
VITE_APP_NAME=AfroChat
VITE_APP_VERSION=1.0.0

# Feature Flags
VITE_ENABLE_PWA=true
VITE_ENABLE_NOTIFICATIONS=true
VITE_ENABLE_OFFLINE_MODE=true
```

### Vite Configuration
```typescript
// vite.config.ts
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { resolve } from 'path'

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': resolve(__dirname, './src'),
      '@components': resolve(__dirname, './src/components'),
      '@pages': resolve(__dirname, './src/pages'),
      '@hooks': resolve(__dirname, './src/hooks'),
      '@services': resolve(__dirname, './src/services'),
      '@store': resolve(__dirname, './src/store'),
      '@types': resolve(__dirname, './src/types'),
      '@utils': resolve(__dirname, './src/utils'),
    },
  },
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
  build: {
    outDir: 'dist',
    sourcemap: true,
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ['react', 'react-dom'],
          router: ['react-router-dom'],
          state: ['zustand', '@tanstack/react-query'],
        },
      },
    },
  },
})
```

## 🎨 UI Components

### Base UI Components
```typescript
// components/ui/Button.tsx
import { ButtonHTMLAttributes, forwardRef } from 'react'
import { cva, type VariantProps } from 'class-variance-authority'
import { cn } from '@/utils/helpers'

const buttonVariants = cva(
  'inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:opacity-50 disabled:pointer-events-none ring-offset-background',
  {
    variants: {
      variant: {
        default: 'bg-primary text-primary-foreground hover:bg-primary/90',
        destructive: 'bg-destructive text-destructive-foreground hover:bg-destructive/90',
        outline: 'border border-input hover:bg-accent hover:text-accent-foreground',
        secondary: 'bg-secondary text-secondary-foreground hover:bg-secondary/80',
        ghost: 'hover:bg-accent hover:text-accent-foreground',
        link: 'underline-offset-4 hover:underline text-primary',
      },
      size: {
        default: 'h-10 py-2 px-4',
        sm: 'h-9 px-3 rounded-md',
        lg: 'h-11 px-8 rounded-md',
        icon: 'h-10 w-10',
      },
    },
    defaultVariants: {
      variant: 'default',
      size: 'default',
    },
  }
)

export interface ButtonProps
  extends ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {}

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant, size, ...props }, ref) => {
    return (
      <button
        className={cn(buttonVariants({ variant, size, className }))}
        ref={ref}
        {...props}
      />
    )
  }
)
Button.displayName = 'Button'

export { Button, buttonVariants }
```

### Chat Components
```typescript
// components/chat/MessageList.tsx
import { useEffect, useRef } from 'react'
import { Message } from '@/types/message'
import { MessageItem } from './MessageItem'
import { TypingIndicator } from './TypingIndicator'

interface MessageListProps {
  messages: Message[]
  typingUsers: string[]
  isLoading?: boolean
}

export const MessageList = ({ messages, typingUsers, isLoading }: MessageListProps) => {
  const messagesEndRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' })
  }, [messages])

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
      </div>
    )
  }

  return (
    <div className="flex-1 overflow-y-auto p-4 space-y-4">
      {messages.map((message) => (
        <MessageItem key={message.id} message={message} />
      ))}
      {typingUsers.length > 0 && (
        <TypingIndicator users={typingUsers} />
      )}
      <div ref={messagesEndRef} />
    </div>
  )
}
```

## 🔄 State Management

### Zustand Stores
```typescript
// store/authStore.ts
import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import { User } from '@/types/user'

interface AuthState {
  user: User | null
  isAuthenticated: boolean
  isLoading: boolean
  login: (user: User, tokens: { accessToken: string; refreshToken: string }) => void
  logout: () => void
  setLoading: (loading: boolean) => void
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      isAuthenticated: false,
      isLoading: false,
      login: (user, tokens) => {
        localStorage.setItem('accessToken', tokens.accessToken)
        localStorage.setItem('refreshToken', tokens.refreshToken)
        set({ user, isAuthenticated: true, isLoading: false })
      },
      logout: () => {
        localStorage.removeItem('accessToken')
        localStorage.removeItem('refreshToken')
        set({ user: null, isAuthenticated: false, isLoading: false })
      },
      setLoading: (loading) => set({ isLoading: loading }),
    }),
    {
      name: 'auth-storage',
      partialize: (state) => ({ user: state.user, isAuthenticated: state.isAuthenticated }),
    }
  )
)
```

### React Query Integration
```typescript
// hooks/useChat.ts
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { chatService } from '@/services/chat'
import { Message, SendMessageRequest } from '@/types/message'

export const useMessages = (channelId: string) => {
  return useQuery({
    queryKey: ['messages', channelId],
    queryFn: () => chatService.getMessages(channelId),
    enabled: !!channelId,
  })
}

export const useSendMessage = () => {
  const queryClient = useQueryClient()
  
  return useMutation({
    mutationFn: (data: SendMessageRequest) => chatService.sendMessage(data),
    onSuccess: (newMessage, variables) => {
      queryClient.setQueryData(['messages', variables.channelId], (old: Message[]) => [
        ...(old || []),
        newMessage,
      ])
    },
  })
}
```

## 🌐 Real-time Communication

### WebSocket Service
```typescript
// services/websocket.ts
import { io, Socket } from 'socket.io-client'
import { WebSocketMessage } from '@/types/websocket'

class WebSocketService {
  private socket: Socket | null = null
  private listeners: Map<string, Function[]> = new Map()

  connect(token: string) {
    this.socket = io(import.meta.env.VITE_WS_URL, {
      auth: { token },
      transports: ['websocket'],
    })

    this.socket.on('connect', () => {
      console.log('WebSocket connected')
    })

    this.socket.on('disconnect', () => {
      console.log('WebSocket disconnected')
    })

    this.socket.on('message_sent', (data) => {
      this.emit('message_sent', data)
    })

    this.socket.on('user_typing', (data) => {
      this.emit('user_typing', data)
    })

    this.socket.on('user_joined', (data) => {
      this.emit('user_joined', data)
    })
  }

  joinChannel(channelId: string) {
    this.socket?.emit('join_channel', channelId)
  }

  leaveChannel(channelId: string) {
    this.socket?.emit('leave_channel', channelId)
  }

  sendMessage(data: { channelId: string; content: string }) {
    this.socket?.emit('send_message', data)
  }

  sendTyping(channelId: string) {
    this.socket?.emit('typing', channelId)
  }

  on(event: string, callback: Function) {
    if (!this.listeners.has(event)) {
      this.listeners.set(event, [])
    }
    this.listeners.get(event)?.push(callback)
  }

  off(event: string, callback: Function) {
    const listeners = this.listeners.get(event)
    if (listeners) {
      const index = listeners.indexOf(callback)
      if (index > -1) {
        listeners.splice(index, 1)
      }
    }
  }

  private emit(event: string, data: any) {
    const listeners = this.listeners.get(event)
    if (listeners) {
      listeners.forEach(callback => callback(data))
    }
  }

  disconnect() {
    this.socket?.disconnect()
    this.socket = null
    this.listeners.clear()
  }
}

export const websocketService = new WebSocketService()
```

### WebSocket Hook
```typescript
// hooks/useWebSocket.ts
import { useEffect, useCallback } from 'react'
import { websocketService } from '@/services/websocket'
import { useAuthStore } from '@/store/authStore'

export const useWebSocket = () => {
  const { user, isAuthenticated } = useAuthStore()

  useEffect(() => {
    if (isAuthenticated && user) {
      const token = localStorage.getItem('accessToken')
      if (token) {
        websocketService.connect(token)
      }
    } else {
      websocketService.disconnect()
    }

    return () => {
      websocketService.disconnect()
    }
  }, [isAuthenticated, user])

  const joinChannel = useCallback((channelId: string) => {
    websocketService.joinChannel(channelId)
  }, [])

  const leaveChannel = useCallback((channelId: string) => {
    websocketService.leaveChannel(channelId)
  }, [])

  const sendMessage = useCallback((data: { channelId: string; content: string }) => {
    websocketService.sendMessage(data)
  }, [])

  const sendTyping = useCallback((channelId: string) => {
    websocketService.sendTyping(channelId)
  }, [])

  return {
    joinChannel,
    leaveChannel,
    sendMessage,
    sendTyping,
    on: websocketService.on.bind(websocketService),
    off: websocketService.off.bind(websocketService),
  }
}
```

## 📱 PWA Features

### Service Worker
```typescript
// public/sw.js
const CACHE_NAME = 'afrochat-v1'
const urlsToCache = [
  '/',
  '/static/js/bundle.js',
  '/static/css/main.css',
  '/manifest.json',
]

self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open(CACHE_NAME)
      .then((cache) => cache.addAll(urlsToCache))
  )
})

self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches.match(event.request)
      .then((response) => {
        if (response) {
          return response
        }
        return fetch(event.request)
      })
  )
})
```

### Manifest Configuration
```json
// public/manifest.json
{
  "name": "AfroChat",
  "short_name": "AfroChat",
  "description": "Real-time communication platform",
  "start_url": "/",
  "display": "standalone",
  "background_color": "#ffffff",
  "theme_color": "#000000",
  "icons": [
    {
      "src": "/icon-192x192.png",
      "sizes": "192x192",
      "type": "image/png"
    },
    {
      "src": "/icon-512x512.png",
      "sizes": "512x512",
      "type": "image/png"
    }
  ]
}
```

## 📹 Video Calling Implementation

### Video Call Hook
```typescript
// hooks/useVideoCall.ts
import { useState, useEffect, useCallback } from 'react'
import { videoCallService } from '@/services/videoCall'
import { webrtcService } from '@/services/webrtc'

export const useVideoCall = (channelId: string) => {
  const [isInCall, setIsInCall] = useState(false)
  const [participants, setParticipants] = useState<Participant[]>([])
  const [isMuted, setIsMuted] = useState(false)
  const [isVideoEnabled, setIsVideoEnabled] = useState(true)
  const [isScreenSharing, setIsScreenSharing] = useState(false)

  const startCall = useCallback(async (title: string) => {
    try {
      const call = await videoCallService.startCall({
        channelId,
        title,
        maxParticipants: 10,
        isPrivate: false
      })
      
      setIsInCall(true)
      await webrtcService.initializeCall(call.id)
      
      return call
    } catch (error) {
      console.error('Failed to start call:', error)
      throw error
    }
  }, [channelId])

  const joinCall = useCallback(async (callId: string) => {
    try {
      await videoCallService.joinCall(callId)
      setIsInCall(true)
      await webrtcService.joinCall(callId)
    } catch (error) {
      console.error('Failed to join call:', error)
      throw error
    }
  }, [])

  const leaveCall = useCallback(async () => {
    try {
      await webrtcService.leaveCall()
      await videoCallService.leaveCall()
      setIsInCall(false)
      setParticipants([])
    } catch (error) {
      console.error('Failed to leave call:', error)
    }
  }, [])

  const toggleMute = useCallback(() => {
    webrtcService.toggleMute()
    setIsMuted(!isMuted)
  }, [isMuted])

  const toggleVideo = useCallback(() => {
    webrtcService.toggleVideo()
    setIsVideoEnabled(!isVideoEnabled)
  }, [isVideoEnabled])

  const toggleScreenShare = useCallback(async () => {
    if (isScreenSharing) {
      await webrtcService.stopScreenShare()
    } else {
      await webrtcService.startScreenShare()
    }
    setIsScreenSharing(!isScreenSharing)
  }, [isScreenSharing])

  return {
    isInCall,
    participants,
    isMuted,
    isVideoEnabled,
    isScreenSharing,
    startCall,
    joinCall,
    leaveCall,
    toggleMute,
    toggleVideo,
    toggleScreenShare
  }
}
```

### WebRTC Service
```typescript
// services/webrtc.ts
class WebRTCService {
  private peerConnections: Map<string, RTCPeerConnection> = new Map()
  private localStream: MediaStream | null = null
  private screenStream: MediaStream | null = null

  async initializeCall(callId: string) {
    // Get user media
    this.localStream = await navigator.mediaDevices.getUserMedia({
      video: true,
      audio: true
    })

    // Setup WebSocket listeners for WebRTC signaling
    websocketService.on('webrtc_offer', this.handleOffer.bind(this))
    websocketService.on('webrtc_answer', this.handleAnswer.bind(this))
    websocketService.on('webrtc_ice_candidate', this.handleICECandidate.bind(this))
  }

  async createPeerConnection(userId: string): Promise<RTCPeerConnection> {
    const configuration = {
      iceServers: [
        { urls: 'stun:stun.l.google.com:19302' },
        { urls: 'stun:stun1.l.google.com:19302' }
      ]
    }

    const peerConnection = new RTCPeerConnection(configuration)
    
    // Add local stream
    if (this.localStream) {
      this.localStream.getTracks().forEach(track => {
        peerConnection.addTrack(track, this.localStream!)
      })
    }

    // Handle ICE candidates
    peerConnection.onicecandidate = (event) => {
      if (event.candidate) {
        this.sendICECandidate(userId, event.candidate)
      }
    }

    // Handle remote stream
    peerConnection.ontrack = (event) => {
      this.handleRemoteStream(userId, event.streams[0])
    }

    this.peerConnections.set(userId, peerConnection)
    return peerConnection
  }

  async createOffer(userId: string): Promise<RTCSessionDescriptionInit> {
    const peerConnection = await this.createPeerConnection(userId)
    const offer = await peerConnection.createOffer()
    await peerConnection.setLocalDescription(offer)
    
    this.sendOffer(userId, offer)
    return offer
  }

  async handleOffer(offer: WebRTCOffer) {
    const peerConnection = await this.createPeerConnection(offer.fromUserId)
    await peerConnection.setRemoteDescription(offer.offer)
    
    const answer = await peerConnection.createAnswer()
    await peerConnection.setLocalDescription(answer)
    
    this.sendAnswer(offer.fromUserId, answer)
  }

  async handleAnswer(answer: WebRTCAnswer) {
    const peerConnection = this.peerConnections.get(answer.fromUserId)
    if (peerConnection) {
      await peerConnection.setRemoteDescription(answer.answer)
    }
  }

  async handleICECandidate(candidate: ICECandidate) {
    const peerConnection = this.peerConnections.get(candidate.fromUserId)
    if (peerConnection) {
      await peerConnection.addIceCandidate(candidate.candidate)
    }
  }

  async startScreenShare() {
    try {
      this.screenStream = await navigator.mediaDevices.getDisplayMedia({
        video: true,
        audio: true
      })

      // Replace video track in all peer connections
      this.peerConnections.forEach(peerConnection => {
        const videoTrack = this.screenStream!.getVideoTracks()[0]
        const sender = peerConnection.getSenders().find(s => 
          s.track && s.track.kind === 'video'
        )
        if (sender) {
          sender.replaceTrack(videoTrack)
        }
      })

      return this.screenStream
    } catch (error) {
      console.error('Failed to start screen share:', error)
      throw error
    }
  }

  async stopScreenShare() {
    if (this.screenStream) {
      this.screenStream.getTracks().forEach(track => track.stop())
      this.screenStream = null
    }
  }

  toggleMute() {
    if (this.localStream) {
      const audioTrack = this.localStream.getAudioTracks()[0]
      if (audioTrack) {
        audioTrack.enabled = !audioTrack.enabled
      }
    }
  }

  toggleVideo() {
    if (this.localStream) {
      const videoTrack = this.localStream.getVideoTracks()[0]
      if (videoTrack) {
        videoTrack.enabled = !videoTrack.enabled
      }
    }
  }

  private sendOffer(userId: string, offer: RTCSessionDescriptionInit) {
    websocketService.send('webrtc_offer', {
      toUserId: userId,
      offer,
      timestamp: new Date().toISOString()
    })
  }

  private sendAnswer(userId: string, answer: RTCSessionDescriptionInit) {
    websocketService.send('webrtc_answer', {
      toUserId: userId,
      answer,
      timestamp: new Date().toISOString()
    })
  }

  private sendICECandidate(userId: string, candidate: RTCIceCandidate) {
    websocketService.send('webrtc_ice_candidate', {
      toUserId: userId,
      candidate,
      timestamp: new Date().toISOString()
    })
  }

  private handleRemoteStream(userId: string, stream: MediaStream) {
    // Emit event for UI to handle remote stream
    window.dispatchEvent(new CustomEvent('remoteStream', {
      detail: { userId, stream }
    }))
  }
}

export const webrtcService = new WebRTCService()
```

### Video Call Component
```typescript
// components/video/VideoCallModal.tsx
import { useState, useEffect, useRef } from 'react'
import { useVideoCall } from '@/hooks/useVideoCall'
import { ParticipantGrid } from './ParticipantGrid'
import { VideoCallControls } from './VideoCallControls'

interface VideoCallModalProps {
  channelId: string
  callId?: string
  onClose: () => void
}

export const VideoCallModal = ({ channelId, callId, onClose }: VideoCallModalProps) => {
  const {
    isInCall,
    participants,
    isMuted,
    isVideoEnabled,
    isScreenSharing,
    startCall,
    joinCall,
    leaveCall,
    toggleMute,
    toggleVideo,
    toggleScreenShare
  } = useVideoCall(channelId)

  const [callTitle, setCallTitle] = useState('')
  const [isStarting, setIsStarting] = useState(false)

  useEffect(() => {
    if (callId) {
      joinCall(callId)
    }
  }, [callId, joinCall])

  const handleStartCall = async () => {
    if (!callTitle.trim()) return
    
    setIsStarting(true)
    try {
      await startCall(callTitle)
    } catch (error) {
      console.error('Failed to start call:', error)
    } finally {
      setIsStarting(false)
    }
  }

  const handleLeaveCall = () => {
    leaveCall()
    onClose()
  }

  if (!isInCall && !callId) {
    return (
      <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div className="bg-white rounded-lg p-6 w-96">
          <h2 className="text-xl font-bold mb-4">Start Video Call</h2>
          <input
            type="text"
            placeholder="Call title"
            value={callTitle}
            onChange={(e) => setCallTitle(e.target.value)}
            className="w-full p-2 border rounded mb-4"
          />
          <div className="flex gap-2">
            <button
              onClick={handleStartCall}
              disabled={isStarting || !callTitle.trim()}
              className="flex-1 bg-blue-500 text-white py-2 px-4 rounded disabled:opacity-50"
            >
              {isStarting ? 'Starting...' : 'Start Call'}
            </button>
            <button
              onClick={onClose}
              className="flex-1 bg-gray-500 text-white py-2 px-4 rounded"
            >
              Cancel
            </button>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="fixed inset-0 bg-black flex flex-col z-50">
      <div className="flex-1 relative">
        <ParticipantGrid participants={participants} />
      </div>
      
      <VideoCallControls
        isMuted={isMuted}
        isVideoEnabled={isVideoEnabled}
        isScreenSharing={isScreenSharing}
        onToggleMute={toggleMute}
        onToggleVideo={toggleVideo}
        onToggleScreenShare={toggleScreenShare}
        onLeaveCall={handleLeaveCall}
      />
    </div>
  )
}
```

## 🔔 Notifications

### Notification Service
```typescript
// services/notification.ts
class NotificationService {
  private permission: NotificationPermission = 'default'

  async requestPermission(): Promise<boolean> {
    if (!('Notification' in window)) {
      console.warn('This browser does not support notifications')
      return false
    }

    if (this.permission === 'granted') {
      return true
    }

    this.permission = await Notification.requestPermission()
    return this.permission === 'granted'
  }

  showNotification(title: string, options: NotificationOptions = {}) {
    if (this.permission !== 'granted') {
      return
    }

    const notification = new Notification(title, {
      icon: '/icon-192x192.png',
      badge: '/icon-192x192.png',
      ...options,
    })

    notification.onclick = () => {
      window.focus()
      notification.close()
    }

    return notification
  }

  showMessageNotification(message: { content: string; user: { displayName: string } }) {
    this.showNotification(
      `New message from ${message.user.displayName}`,
      {
        body: message.content,
        tag: 'message',
        requireInteraction: false,
      }
    )
  }

  showCallNotification(call: { title: string; host: { displayName: string } }) {
    this.showNotification(
      `Video call: ${call.title}`,
      {
        body: `Started by ${call.host.displayName}`,
        tag: 'video_call',
        requireInteraction: true,
        actions: [
          { action: 'join', title: 'Join Call' },
          { action: 'decline', title: 'Decline' }
        ]
      }
    )
  }
}

export const notificationService = new NotificationService()
```

## 🧪 Testing

### Unit Tests
```typescript
// __tests__/components/MessageItem.test.tsx
import { render, screen } from '@testing-library/react'
import { MessageItem } from '@/components/chat/MessageItem'
import { Message } from '@/types/message'

const mockMessage: Message = {
  id: '1',
  content: 'Hello world!',
  userId: 'user1',
  channelId: 'channel1',
  createdAt: new Date().toISOString(),
  user: {
    id: 'user1',
    username: 'johndoe',
    displayName: 'John Doe',
    avatarUrl: null,
  },
}

describe('MessageItem', () => {
  it('renders message content', () => {
    render(<MessageItem message={mockMessage} />)
    expect(screen.getByText('Hello world!')).toBeInTheDocument()
  })

  it('displays user information', () => {
    render(<MessageItem message={mockMessage} />)
    expect(screen.getByText('John Doe')).toBeInTheDocument()
  })
})
```

### E2E Tests
```typescript
// e2e/chat.spec.ts
import { test, expect } from '@playwright/test'

test('user can send and receive messages', async ({ page }) => {
  await page.goto('/login')
  await page.fill('[data-testid=email]', 'test@example.com')
  await page.fill('[data-testid=password]', 'password')
  await page.click('[data-testid=login-button]')

  await page.waitForURL('/workspace')
  await page.click('[data-testid=channel-general]')

  await page.fill('[data-testid=message-input]', 'Hello everyone!')
  await page.click('[data-testid=send-button]')

  await expect(page.locator('[data-testid=message-list]')).toContainText('Hello everyone!')
})
```

## 🚀 Build and Deployment

### Build Configuration
```json
// package.json
{
  "scripts": {
    "dev": "vite",
    "build": "tsc && vite build",
    "preview": "vite preview",
    "test": "vitest",
    "test:ui": "vitest --ui",
    "test:e2e": "playwright test",
    "lint": "eslint . --ext ts,tsx --report-unused-disable-directives --max-warnings 0",
    "lint:fix": "eslint . --ext ts,tsx --fix"
  }
}
```

### Docker Configuration
```dockerfile
# Dockerfile
FROM node:18-alpine AS builder

WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production

COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

## 📊 Performance Optimization

### Code Splitting
```typescript
// App.tsx
import { lazy, Suspense } from 'react'
import { Routes, Route } from 'react-router-dom'

const WorkspacePage = lazy(() => import('@/pages/workspace/WorkspacePage'))
const ChannelPage = lazy(() => import('@/pages/workspace/ChannelPage'))
const ProfilePage = lazy(() => import('@/pages/profile/ProfilePage'))

function App() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Routes>
        <Route path="/workspace" element={<WorkspacePage />} />
        <Route path="/channel/:id" element={<ChannelPage />} />
        <Route path="/profile" element={<ProfilePage />} />
      </Routes>
    </Suspense>
  )
}
```

### Virtual Scrolling
```typescript
// components/chat/VirtualMessageList.tsx
import { FixedSizeList as List } from 'react-window'

interface VirtualMessageListProps {
  messages: Message[]
  height: number
}

export const VirtualMessageList = ({ messages, height }: VirtualMessageListProps) => {
  const Row = ({ index, style }: { index: number; style: React.CSSProperties }) => (
    <div style={style}>
      <MessageItem message={messages[index]} />
    </div>
  )

  return (
    <List
      height={height}
      itemCount={messages.length}
      itemSize={80}
      width="100%"
    >
      {Row}
    </List>
  )
}
```

## 🛡️ Security

### Input Validation
```typescript
// utils/validators.ts
import { z } from 'zod'

export const loginSchema = z.object({
  email: z.string().email('Invalid email address'),
  password: z.string().min(8, 'Password must be at least 8 characters'),
})

export const messageSchema = z.object({
  content: z.string().min(1).max(4000, 'Message too long'),
  channelId: z.string().uuid('Invalid channel ID'),
})

export const validateInput = <T>(schema: z.ZodSchema<T>, data: unknown): T => {
  return schema.parse(data)
}
```

### XSS Protection
```typescript
// utils/helpers.ts
import DOMPurify from 'dompurify'

export const sanitizeHtml = (html: string): string => {
  return DOMPurify.sanitize(html, {
    ALLOWED_TAGS: ['b', 'i', 'em', 'strong', 'a'],
    ALLOWED_ATTR: ['href', 'target'],
  })
}
```

## 📚 Development Tools

### ESLint Configuration
```json
// .eslintrc.json
{
  "extends": [
    "eslint:recommended",
    "@typescript-eslint/recommended",
    "plugin:react/recommended",
    "plugin:react-hooks/recommended"
  ],
  "rules": {
    "react/react-in-jsx-scope": "off",
    "@typescript-eslint/no-unused-vars": "error",
    "@typescript-eslint/explicit-function-return-type": "warn"
  }
}
```

### Prettier Configuration
```json
// .prettierrc
{
  "semi": true,
  "trailingComma": "es5",
  "singleQuote": true,
  "printWidth": 80,
  "tabWidth": 2
}
```

## 🤝 Contributing

1. Follow the established code style
2. Write tests for new features
3. Update documentation
4. Ensure all tests pass
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
