<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const auth = useAuthStore()

const form = reactive({ name: '', email: '', password: '' })

async function handleSubmit() {
  try {
    await auth.register(form)
    router.push('/dashboard')
  } catch {
    // auth.error already holds the message.
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-slate-100">
    <div class="w-full max-w-sm rounded-xl bg-white p-8 shadow-sm">
      <h1 class="mb-6 text-xl font-semibold text-slate-900">Create an account</h1>

      <p v-if="auth.error" class="mb-4 rounded-lg bg-red-50 px-3 py-2 text-sm text-red-700">
        {{ auth.error }}
      </p>

      <form class="space-y-4" @submit.prevent="handleSubmit">
        <div>
          <label class="mb-1 block text-sm font-medium text-slate-700">Name</label>
          <input
            v-model="form.name"
            type="text"
            required
            class="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm focus:border-blue-500 focus:outline-none"
          />
        </div>

        <div>
          <label class="mb-1 block text-sm font-medium text-slate-700">Email</label>
          <input
            v-model="form.email"
            type="email"
            required
            class="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm focus:border-blue-500 focus:outline-none"
          />
        </div>

        <div>
          <label class="mb-1 block text-sm font-medium text-slate-700">Password</label>
          <input
            v-model="form.password"
            type="password"
            required
            minlength="8"
            class="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm focus:border-blue-500 focus:outline-none"
          />
        </div>

        <button
          type="submit"
          :disabled="auth.isLoading"
          class="w-full rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 disabled:opacity-60"
        >
          {{ auth.isLoading ? 'Creating account...' : 'Create account' }}
        </button>
      </form>

      <p class="mt-4 text-center text-sm text-slate-500">
        Already have an account?
        <router-link to="/login" class="font-medium text-blue-600">Log in</router-link>
      </p>
    </div>
  </div>
</template>