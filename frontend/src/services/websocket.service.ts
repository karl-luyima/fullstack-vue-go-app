import { tokenStorage } from './api'

type MessageHandler = (data: any) => void

export function connectWebSocket(onMessage: MessageHandler): WebSocket {
  const token = tokenStorage.get()
  const socket = new WebSocket(`ws://localhost:8080/ws?token=${token}`)

  socket.onmessage = (event) => {
    const data = JSON.parse(event.data)
    onMessage(data)
  }

  socket.onerror = (event) => {
    console.error('WebSocket error:', event)
  }

  return socket
}