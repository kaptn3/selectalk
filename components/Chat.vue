<template>
  <div>
    <button @click="display = 'block'"/>
    <img
      :style="{ 'display': display }"
      src="/img/chat.svg"
      alt=""
      @click="display = 'none'"
    >
  </div>
</template>

<script>
  export default {
    name: 'Chat',
    data() {
      return {
        display: 'none'
      }
    }
  };
</script>

<style scoped>
  button {
    box-shadow: 0 4px 10px rgba(0, 0, 0, .5);
    width: 60px;
    height: 60px;
    background: var(--black) url(/icons/chat.png) no-repeat center;
    border-radius: 100%;
  }

  div {
    position: fixed;
    bottom: 42px;
    right: 42px;
  }

  img {
    max-width: initial;
    position: absolute;
    right: -8px;
    bottom: -8px;
    cursor: pointer;
  }
</style>
