<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { Bar } from 'vue-chartjs'
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
} from 'chart.js'
import { analyticsService } from '@/services/analytics.service'
import type { SignupsByDayPoint } from '@/types'

// chart.js requires you to explicitly register the pieces you're using —
// this keeps the library's bundle size small by not including chart
// types/features you never use.
ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)

const points = ref<SignupsByDayPoint[]>([])
const isLoading = ref(true)

onMounted(async () => {
  points.value = await analyticsService.signupsByDay(30)
  isLoading.value = false
})

// vue-chartjs expects data in this specific shape: an array of labels
// (the x-axis) and one or more "datasets" (the actual bars/lines).
const chartData = computed(() => ({
  labels: points.value.map((p) => p.date),
  datasets: [
    {
      label: 'New signups',
      backgroundColor: '#2563eb',
      data: points.value.map((p) => p.count),
    },
  ],
}))

const chartOptions = {
  responsive: true,
  plugins: { legend: { display: false } },
  scales: { y: { beginAtZero: true, ticks: { stepSize: 1 } } },
}
</script>

<template>
  <div class="rounded-xl border border-slate-200 bg-white p-5">
    <p class="mb-3 text-sm font-medium text-slate-500">Signups, last 30 days</p>

    <p v-if="isLoading" class="text-sm text-slate-500">Loading...</p>
    <p v-else-if="points.length === 0" class="text-sm text-slate-500">No signups yet.</p>
    <div v-else class="relative h-48">
      <Bar :data="chartData" :options="chartOptions" />
    </div>
  </div>
</template>