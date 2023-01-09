<template>
  <apexchart ref="chart" height="300" type="line" :options="chartOptions" :series="series"/>
</template>

<script>
export default {
  name: "RealtimeChart",
  props: {
    color: { type: String },
    airport: { type: String },
    sensor: { type: String },
    min: { type: Number },
    max: { type: Number },
    period: { type: Number },
  },
  data() {
    return {
      series: [],
      chartOptions: {
        colors: [this.color],
        chart: {
          id: 'realtime',
          height: 350,
          type: 'line',
          animations: {
            enabled: true,
            easing: 'linear',
            dynamicAnimation: {
              speed: 1000
            }
          },
          toolbar: {
            show: false
          },
          zoom: {
            enabled: false
          }
        },
        grid: {
          borderColor: 'rgb(59,54,82)',
          strokeDashArray: 7,
        },
        dataLabels: {
          enabled: false
        },
        stroke: {
          curve: 'smooth'
        },
        markers: {
          size: 0
        },
        xaxis: {
          type: 'datetime',
          range: 1000 * 60 * this.period,
          labels: {
            style: {
              colors: '#A39CAD'
            }
          },
          axisBorder: {
            show: false,
          },
          axisTicks: {
            show: false
          }
        },
        axisBorder: {
          show: false,
        },
        yaxis: {
          max: this.max,
          min: this.min,
          labels: {
            style: {
              colors: '#A39CAD'
            }
          }
        },
        legend: {
          show: false
        },
      }
    }
  },
  mounted() {
    this.updateChart();

    setInterval(this.updateChart, 10000)
  },
  methods: {
    async updateChart() {
      console.log(this.period)
      const from = new Date();
      const to = new Date();

      from.setMinutes(-this.period);

      const json = await (await fetch(`http://localhost:8080/airport/${this.airport}/measure?type=${this.sensor}&startDate=${from.toISOString()}&endDate=${to.toISOString()}`)).json()

      const data = [{ name: this.sensor, data: json.map(obj => { return { y: obj.value.toFixed(2), x: new Date(obj.date) } })}];

      console.log(this.series)

      this.$refs.chart.updateSeries(data);
    }
  }
}
</script>

<style scoped>

</style>