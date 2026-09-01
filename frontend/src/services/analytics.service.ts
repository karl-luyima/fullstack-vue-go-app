import api from './api'
import type { SignupsByDayPoint } from '@/types'

export const analyticsService = {
  async signupsByDay(days: number = 30) {
    const { data } = await api.get<SignupsByDayPoint[]>('/analytics/signups', {
      params: { days },
    })
    return data
  },
}