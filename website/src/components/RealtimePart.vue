<template>
  <div class="realtime-part">
    <div class="realtime-card">
      <h2 class="realtime-card-title">Pressure <span class="realtime-card-time">(realtime)</span></h2>

      <RealtimeChart
          class="chart"
          sensor="pressure"
          :min="980"
          :max="1030"
          :airport="airport"
          color="#8455E9"
      />
    </div>

    <div class="realtime-card">
      <h2 class="realtime-card-title">Temperature <span class="realtime-card-time">(realtime)</span></h2>

      <RealtimeChart
          class="chart"
          sensor="temperature"
          :min="0"
          :max="40"
          :airport="airport"
          color="#5ADD77"
      />
    </div>

    <div class="realtime-card">
      <h2 class="realtime-card-title">Wind <span class="realtime-card-time">(realtime)</span></h2>

      <RealtimeChart
          class="chart"
          sensor="wind"
          :min="0"
          :max="80"
          :airport="airport"
          color="#E03BB2"
      />
    </div>
  </div>
</template>

<script>
import RealtimeChart from "@/components/RealtimeChart.vue";

export default {
  name: "RealtimePart",
  components: {RealtimeChart},
  props: {
    airport: { type: String }
  }
}
</script>

<style lang="scss" scoped>
.realtime-part {
  margin-top: 20px;
  gap: 20px;
  display: flex;
  flex-flow: column;
  align-content: center;
  justify-content: center;

  .realtime-card {
    position: relative;

    width: 100%;
    background-color: #1C1A1F;
    border-radius: 15px;
    padding: 10px;

    .chart {
      height: 100%;
    }

    .realtime-card-title {
      color: #FFFFFF;
      font-size: 13px;

      .realtime-card-time {
        color: #716B7A;
      }
    }
  }
}
</style>