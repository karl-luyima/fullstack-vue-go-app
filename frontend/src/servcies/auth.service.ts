import api from './api'
import type { AuthResponse, LoginPayload, RegisterPayload, User } from '@/types'

export const authService = {
  async login(payload: LoginPayload) {
    const { data } = await api.post<AuthResponse>('/auth/login', payload)
    return data
  },

  async register(payload: RegisterPayload) {
    const { data } = await api.post<AuthResponse>('/auth/register', payload)
    return data
  },

  async me() {
    const { data } = await api.get<User>('/users/me')
    return data
  },
}