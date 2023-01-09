<script setup lang="ts">
import AveragePart from "@/components/AveragePart.vue";
import RealtimePart from "@/components/RealtimePart.vue";
import HeaderPart from "@/components/HeaderPart.vue";
import InputAiport from "@/components/InputAirport.vue";
import InputLong from "@/components/InputLong.vue";
</script>

<template>
  <div class="container">
    <HeaderPart/>

    <InputAiport
        id="airport"
        label="Airport"
        v-model="airport"
    />

    <Suspense>
      <AveragePart
          :key="airport"
          :airport="airport"
      />
    </Suspense>

    <InputLong
        id="long"
        label="Period"
        v-model="long"
    />

    <Suspense>
      <RealtimePart
          :key="airport + long"
          :long="parseInt(long)"
          :airport="airport"
      />
    </Suspense>
  </div>
</template>

<script lang="ts">
export default {
  name: "HomeView",
  data() {
    return {
      airport: '',
      long: 10
    }
  }
}
</script>