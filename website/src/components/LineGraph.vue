<script setup lang="ts">
import {Chart, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend} from "chart.js";
import {Line} from 'vue-chartjs'

const props = defineProps({
  sensorType: { type: String }
})

Chart.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
)

const airport = "NTE"
const startDate = new Date();
const endDate = new Date().toISOString()

startDate.setHours(new Date().getHours() - 3)

const fetchData = await fetch(`http://localhost:8080/airport/${airport}/measure?type=${props.sensorType}&startDate=${startDate.toISOString()}&endDate=${endDate}`)

const objs: {
  date: string,
  value: number
}[] = await fetchData.json() || []

const data = {
  labels: objs.map(obj => new Date(obj.date).toLocaleTimeString()),
  datasets: [{
    label: props.sensorType,
    borderColor: '#f87979',
    data: objs.map(obj => obj.value),
    pointStyle: false,
    cubicInterpolationMode: 'monotone',
    tension: 0.4,

  }]
}

const options = {
  responsive: true,
  maintainAspectRatio: false
}
</script>

<template>
  <Line class="chart" :data="data" :options="options"/>
</template>

<style scoped>
.chart {
  width: 100%;
  max-height: 200px;
}
</style>