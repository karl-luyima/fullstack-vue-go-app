<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { connectWebSocket } from '@/services/websocket.service'
import Navbar from '@/components/Navbar.vue'
import StatCard from '@/components/StatCard.vue'
import SignupsChart from '@/components/SignupsChart.vue'

const auth = useAuthStore()

const notifications = ref<string[]>([])
let socket: WebSocket | null = null

onMounted(() => {
  socket = connectWebSocket((data) => {
    if (data.event === 'new_signup') {
      notifications.value.unshift(`${data.name} just signed up`)
    }
  })
})

onUnmounted(() => {
  socket?.close()
})
</script>

<template>
  <div class="min-h-screen bg-slate-50">
    <Navbar />

    <main class="mx-auto max-w-4xl px-6 py-8">
      <h1 class="text-2xl font-semibold text-slate-900">Welcome back, {{ auth.user?.name }}</h1>
      <p class="mt-1 text-sm text-slate-500">Here's what's happening with your account.</p>

      <div class="mt-6 grid grid-cols-3 gap-3">
        <StatCard label="Name" :value="auth.user?.name ?? ''" />
        <StatCard label="Email" :value="auth.user?.email ?? ''" />
        <StatCard label="Status" value="Active" value-class="text-green-600" />
      </div>

      <div v-if="notifications.length > 0" class="mt-6 rounded-xl border border-slate-200 bg-white p-5">
        <p class="mb-3 text-sm font-medium text-slate-500">Live activity</p>
        <ul class="space-y-2">
          <li
            v-for="(note, i) in notifications"
            :key="i"
            class="flex items-center gap-2 rounded-lg bg-blue-50 px-3 py-2 text-sm text-blue-700"
          >
            {{ note }}
          </li>
        </ul>
      </div>

      <div class="mt-6">
        <SignupsChart />
      </div>
    </main>
  </div>
</template>