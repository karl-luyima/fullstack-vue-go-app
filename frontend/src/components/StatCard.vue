<script setup lang="ts">
defineProps<{
  label: string
  value: string | number
  valueClass?: string
}>()
</script>

<template>
  <div class="rounded-lg bg-slate-50 p-4">
    <p class="text-sm text-slate-500">{{ label }}</p>
    <p class="mt-1 text-2xl font-medium" :class="valueClass ?? 'text-slate-900'">{{ value }}</p>
  </div>
</template>