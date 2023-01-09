<template>
  <div class="average-part">
    <div class="average-card">
      <h2 class="average-card-title">Pressure <span class="average-card-time">(today)</span></h2>

      <AverageChart
          class="chart"
          :data="data.pressure.ratio"
          color="#8455E9"
      />

      <div class="average-card-value-container">
        <div class="average-card-value text-purple">
          {{ data.pressure.average.toFixed(2) }}
        </div>
        <div class="average-card-unit">
          {{ data.pressure.unit }}
        </div>
      </div>
    </div>
    <div class="average-card">
      <h2 class="average-card-title">Temperature <span class="average-card-time">(today)</span></h2>

      <AverageChart
          class="chart"
          :data="data.temp.ratio"
          color="#5ADD77"
      />

      <div class="average-card-value-container">
        <div class="average-card-value text-green">
          {{ data.temp.average.toFixed(2) }}
        </div>
        <div class="average-card-unit">
          {{ data.temp.unit }}
        </div>
      </div>
    </div>
    <div class="average-card">
      <h2 class="average-card-title">Wind <span class="average-card-time">(today)</span></h2>

      <AverageChart
          class="chart"
          :data="data.wind.ratio"
          color="#E03BB2"
      />

      <div class="average-card-value-container">
        <div class="average-card-value text-pink">
          {{ data.wind.average.toFixed(2) }}
        </div>
        <div class="average-card-unit">
          {{ data.wind.unit }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AverageChart from "@/components/AverageChart.vue";

export default {
  name: "AveragePart",
  components: {AverageChart},
  props: {
    airport: { type: String }
  },
  data() {
    return {
      data: {
        temp: {
          average: 0,
          ratio: 0,
          unit: ''
        },
        wind: {
          average: 0,
          ratio: 0,
          unit: ''
        },
        pressure: {
          average: 0,
          ratio: 0,
          unit: ''
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
      console.log("test")
      const json = await (await fetch(`http://localhost:8080/airport/${this.airport}/averages?date=${new Date().toISOString().split('T')[0]}`)).json()

      this.data.temp = json.find(el => el.sensortype === "temperature")
      this.data.wind = json.find(el => el.sensortype === "wind")
      this.data.pressure = json.find(el => el.sensortype === "pressure")

      this.data.temp.ratio = ((this.data.temp.average - 0) / (40 - 0)) * 100;
      this.data.wind.ratio = ((this.data.wind.average - 0) / (80 - 0)) * 100;
      this.data.pressure.ratio = ((this.data.pressure.average - 950) / (1050 - 950)) * 100;
    }
  }
}
</script>

<style lang="scss" scoped>
.average-part {
  gap: 20px;
  display: flex;
  align-content: center;
  justify-content: center;

  .average-card {
    position: relative;

    width: 100%;
    max-width: 350px;
    background-color: #1C1A1F;
    border-radius: 15px;
    padding: 10px;

    .average-card-value-container {
      width: calc(100% - 20px);
      bottom: 30px;
      position: absolute;
      display: flex;
      flex-flow: column;
      align-items: center;
      justify-content: center;

      .average-card-value {
        line-height: 1;
        font-size: 24px;
        font-weight: bold;
      }

      .average-card-unit {
        line-height: 1;
        color: #716B7A;
        font-size: 12px;
        font-weight: bold;
      }
    }

    .chart {
      height: 100%;
    }

    .average-card-title {
      color: #FFFFFF;
      font-size: 13px;

      .average-card-time {
        color: #716B7A;
      }
    }
  }
}
</style>