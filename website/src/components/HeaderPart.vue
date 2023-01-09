<template>
  <h1 class="welcome text-white">Welcome back, <span class="text-purple">User</span>!</h1>
</template>

<script>
export default {
  name: "HeaderPart"
}
</script>

<style scoped lang="scss">
.welcome {
  font-size: 24px;
  font-weight: bold;
  margin: 20px 0;

  span {
    font-weight: bold;
  }
}
</style>