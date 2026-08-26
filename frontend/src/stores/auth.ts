import { defineStore } from 'pinia'
import { authService } from '@/services/auth.service'
import { tokenStorage } from '@/services/api'
import type { LoginPayload, RegisterPayload, User } from '@/types'

interface AuthState {
  user: User | null
  isLoading: boolean
  error: string | null
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    user: null,
    isLoading: false,
    error: null,
  }),

  getters: {
    isAuthenticated: (state) => !!state.user,
  },

  actions: {
    async login(payload: LoginPayload) {
      this.isLoading = true
      this.error = null
      try {
        const result = await authService.login(payload)
        tokenStorage.set(result.token)
        this.user = result.user
      } catch (err: any) {
        this.error = err?.response?.data?.error ?? 'Login failed. Please try again.'
        throw err
      } finally {
        this.isLoading = false
      }
    },

    async register(payload: RegisterPayload) {
      this.isLoading = true
      this.error = null
      try {
        const result = await authService.register(payload)
        tokenStorage.set(result.token)
        this.user = result.user
      } catch (err: any) {
        this.error = err?.response?.data?.error ?? 'Registration failed. Please try again.'
        throw err
      } finally {
        this.isLoading = false
      }
    },

    logout() {
      tokenStorage.clear()
      this.user = null
    },
  },
})