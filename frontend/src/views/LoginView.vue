<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const auth = useAuthStore()

const form = reactive({ email: '', password: '' })

async function handleSubmit() {
  try {
    await auth.login(form)
    router.push('/dashboard')
  } catch {
    // auth.error already holds the message.
  }
}
</script>

<template>
  <div class="flex min-h-screen items-center justify-center bg-slate-50 px-4">
    <div class="w-full max-w-sm">
      <div class="mb-6 flex items-center justify-center gap-2">
        <div class="flex h-8 w-8 items-center justify-center rounded-lg bg-blue-600 text-sm font-medium text-white">
          A
        </div>
        <span class="text-lg font-medium text-slate-900">Acme app</span>
      </div>

      <div class="rounded-xl border border-slate-200 bg-white p-8 shadow-sm">
        <h1 class="mb-1 text-xl font-semibold text-slate-900">Welcome back</h1>
        <p class="mb-6 text-sm text-slate-500">Log in to your account to continue.</p>

        <p v-if="auth.error" class="mb-4 rounded-lg bg-red-50 px-3 py-2 text-sm text-red-700">
          {{ auth.error }}
        </p>

        <form class="space-y-4" @submit.prevent="handleSubmit">
          <div>
            <label class="mb-1.5 block text-sm font-medium text-slate-700">Email</label>
            <input
              v-model="form.email"
              type="email"
              required
              placeholder="you@example.com"
              class="w-full rounded-lg border border-slate-300 px-3 py-2.5 text-sm shadow-sm placeholder:text-slate-400 focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
            />
          </div>

          <div>
            <label class="mb-1.5 block text-sm font-medium text-slate-700">Password</label>
            <input
              v-model="form.password"
              type="password"
              required
              placeholder="••••••••"
              class="w-full rounded-lg border border-slate-300 px-3 py-2.5 text-sm shadow-sm placeholder:text-slate-400 focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
            />
          </div>

          <button
            type="submit"
            :disabled="auth.isLoading"
            class="w-full rounded-lg bg-blue-600 px-4 py-2.5 text-sm font-medium text-white shadow-sm transition-colors hover:bg-blue-700 disabled:cursor-not-allowed disabled:opacity-60"
          >
            {{ auth.isLoading ? 'Logging in...' : 'Log in' }}
          </button>
        </form>
      </div>

      <p class="mt-6 text-center text-sm text-slate-500">
        No account?
        <router-link to="/register" class="font-medium text-blue-600 hover:text-blue-700">Register</router-link>
      </p>
    </div>
  </div>
</template>