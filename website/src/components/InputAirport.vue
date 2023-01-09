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
        <option v-for="airport in airports">{{ airport }}</option>
      </select>
      <div class="arrow"/>
    </div>
  </div>
</template>

<script>
export default {
  name: "InputAirport",
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
  data() {
    return {
      airports: []
    }
  },
  mounted() {
    this.loadOptions();
  },
  methods: {
    async loadOptions() {
      this.airports = await fetch(`http://localhost:8080/airports`).then(api => api.json())

      this.$emit("update:modelValue", this.airports[0]);
    },
    updateInput(event) {
      this.$emit("update:modelValue", event.target.value);
    }
  }
};
</script>

<style lang="scss" scoped>
.label {
  display: block;
  color: #A39CAD;
}

.select-wrapper {
  position: relative;
  width: min-content;

  .select {
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