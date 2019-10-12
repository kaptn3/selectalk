<template>
  <label>
    <span :class="{ 'label_top': onTop || value.length > 0 }">
      {{ label }}
    </span>
    <input
      v-model="value"
      type="text"
      @focus="onTop = true"
      @blur="onTop = false"
    >
  </label>
</template>

<script>
  export default {
    name: 'AInput',
    props: {
      label: {
        type: String,
        required: true
      }
    },
    data() {
      return {
        onTop: false,
        value: ''
      };
    }
  };
</script>

<style scoped>
  label {
    position: relative;
    font-size: 13px;
  }

  input {
    background-color: #fff;
    border: 1px solid rgba(22, 22, 31, .1);
    border-radius: 8px;
    outline: none;
    padding: 8px;
    width: 100%;
  }

  span {
    position: absolute;
    left: 8px;
    top: 6px;
    transition: .2s top;
  }

  .label_top {
    font-size: 10px;
    top: -7px;
    background-color: #fff;
  }
</style>
