<template>
  <apexchart type="radialBar" :options="chartOptions" :series="[data]"/>
</template>

<script>
export default {
  name: "AverageChart",
  props: {
    color: { type: String },
    data: { type: Number }
  },
  data() {
    return {
      chartOptions: {
        chart: {
          type: 'radialBar'
        },
        plotOptions: {
          radialBar: {
            startAngle: -90,
            endAngle: 90,
            track: {
              background: "#2A282E",
            },
            dataLabels: {
              name: {
                show: false
              },
              value: {
                show: false
              }
            }
          }
        },
        grid: {
          padding: {
            top: -10,
            left: 0,
            right: 0,
          }
        },
        fill: {
          colors: [this.color]
        },
      }
    }
  }
}
</script>

<style scoped>

</style>