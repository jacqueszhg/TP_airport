<script setup lang="ts">
import AveragePart from "@/components/AveragePart.vue";
import RealtimePart from "@/components/RealtimePart.vue";
import HeaderPart from "@/components/HeaderPart.vue";
import InputAiport from "@/components/InputAirport.vue";
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

    <Suspense>
      <RealtimePart
          :key="airport"
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
      airport: ''
    }
  }
}
</script>