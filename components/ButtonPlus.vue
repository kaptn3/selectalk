<template>
  <button>
    <span>+</span>
  </button>
</template>

<script>
  export default {
    name: 'ButtonPlus'
  };
</script>

<style scoped>
  button {
    color: var(--white);
    background-color: var(--red);
    border-radius: 100%;
    padding: 0;
    height: 28px;
    width: 28px;
    font-size: 24px;
  }

  span {
    margin-bottom: 10px;
    display: block;
    line-height: 1;
  }
</style>
