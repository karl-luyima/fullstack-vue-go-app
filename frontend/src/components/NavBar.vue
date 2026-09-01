<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const auth = useAuthStore()

function initials(name: string | undefined) {
  if (!name) return '?'
  return name
    .split(' ')
    .map((p) => p[0])
    .join('')
    .slice(0, 2)
    .toUpperCase()
}

function handleLogout() {
  auth.logout()
  router.push('/login')
}
</script>

<template>
  <header class="border-b border-slate-200 bg-white">
    <div class="mx-auto flex max-w-4xl items-center justify-between px-6 py-4">
      <div class="flex items-center gap-2">
        <div class="flex h-7 w-7 items-center justify-center rounded-lg bg-blue-600 text-sm font-medium text-white">
          A
        </div>
        <span class="font-medium text-slate-900">Acme app</span>
      </div>

      <div class="flex items-center gap-4">
        <div
          class="flex h-8 w-8 items-center justify-center rounded-full bg-blue-50 text-xs font-medium text-blue-700"
        >
          {{ initials(auth.user?.name) }}
        </div>
        <button class="text-sm text-slate-500 hover:text-slate-900" @click="handleLogout">Log out</button>
      </div>
    </div>
  </header>
</template>