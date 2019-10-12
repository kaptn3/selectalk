<template>
  <span>
    {{ name }}
  </span>
</template>

<script>
  export default {
    name: 'Badge',
    props: {
      name: {
        type: String,
        required: true
      }
    }
  };
</script>

<style scoped>
  span {
    display: inline-block;
    color: var(--white);
    background-color: var(--red);
    border-radius: 8px;
    padding: 4px 1rem;
    font-size: .5rem;
    line-height: 1;
  }
</style>
