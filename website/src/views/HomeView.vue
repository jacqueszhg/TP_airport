<script setup lang="ts">
import LineGraph from "@/components/LineGraph.vue";
</script>

<template>
  <main>
    <Suspense>
      <LineGraph sensor-type="temperature"/>
    </Suspense>
    <Suspense>
      <LineGraph sensor-type="wind"/>
    </Suspense>
    <Suspense>
      <LineGraph sensor-type="pressure"/>
    </Suspense>
  </main>
</template>