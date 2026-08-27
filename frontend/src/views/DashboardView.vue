<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import SignupsChart from '@/components/SignupsChart.vue'

const router = useRouter()
const auth = useAuthStore()

function handleLogout() {
  auth.logout()
  router.push('/login')
}
</script>

<template>
  <div class="min-h-screen bg-slate-100 p-8">
    <div class="mx-auto max-w-2xl space-y-6">
      <div class="rounded-xl bg-white p-8 shadow-sm">
        <div class="flex items-center justify-between">
          <h1 class="text-xl font-semibold text-slate-900">Dashboard</h1>
          <button
            class="rounded-lg border border-slate-300 px-3 py-1.5 text-sm text-slate-700 hover:bg-slate-50"
            @click="handleLogout"
          >
            Log out
          </button>
        </div>

        <div class="mt-6 space-y-2 text-sm text-slate-700">
          <p><span class="font-medium">Name:</span> {{ auth.user?.name }}</p>
          <p><span class="font-medium">Email:</span> {{ auth.user?.email }}</p>
        </div>
      </div>

      <SignupsChart />
    </div>
  </div>
</template>