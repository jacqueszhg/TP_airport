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
const startDate = "2021-04-04T22%3A08%3A41Z"
const endDate = new Date().toISOString()

const fetchData = await fetch(`http://localhost:8080/airport/${airport}/measure?type=${props.sensorType}&startDate=${startDate}&endDate=${endDate}`)

const objs: {
  date: string,
  value: number
}[] = await fetchData.json()

const data = {
  labels: objs.map(obj => new Date(obj.date).toUTCString()),
  datasets: [{
    label: props.sensorType,
    backgroundColor: '#f87979',
    data: objs.map(obj => obj.value)
  }]
}

const options = {
  responsive: false,
  maintainAspectRatio: false
}
</script>

<template>
  <Line :data="data" :options="options"/>
</template>

<style scoped>

</style>