<script setup lang="ts">
import AveragePart from "@/components/AveragePart.vue";
import RealtimePart from "@/components/RealtimePart.vue";
</script>

<template>
  <div class="container">
    <Suspense>
      <AveragePart
          airport="NTE"
      />
    </Suspense>

    <Suspense>
      <RealtimePart
          airport="NTE"
      />
    </Suspense>
  </div>
</template>