<template>
  <div>
    <label v-if="label" :for="id" class="label text-white">
      {{ label }}
    </label>
    <div class="select-wrapper">
      <select
          :id="id"
          :value="modelValue"
          class="select"
          @input="updateInput"
      >
        <option value="1">1 minute</option>
        <option selected value="10">10 minutes</option>
        <option value="60">1 heure</option>
        <option value="1440">1 jour</option>
        <option value="10080">1 semaine</option>
      </select>
    </div>
  </div>
</template>

<script>
export default {
  name: "InputLong",
  props: {
    id: {
      type: String,
      default: "",
    },
    label: {
      type: String,
      default: "",
    },
    modelValue: {
      type: [String, Number],
      default: "",
    }
  },
  mounted() {
    this.$emit("update:modelValue", 10);
  },
  methods: {
    updateInput(event) {
      this.$emit("update:modelValue", event.target.value);
    }
  }
};
</script>

<style lang="scss" scoped>
.label {
  display: block;
  margin-top: 20px;
  color: #A39CAD;
}

.select-wrapper {
  position: relative;
  width: min-content;

  .select {
    border-radius: 10px;
    width: 200px;
    cursor: pointer;
    appearance: none;
    background-color: #1C1A1F;
    border: 0;
    padding: 10px 20px;
    color: #FFFFFF;
    font-size: 13px;
    font-weight: bold;
  }
}
</style>