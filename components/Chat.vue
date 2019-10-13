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
    border-radius: 30px;
    transition: .2s width, .2s background-position-x;
  }

  button::after {
    content: 'Появились вопросы?';
    width: 0;
    overflow: hidden;
    color: #fff;
    display: inline-block;
    margin-right: 10px;
    transition: .2s width;
    white-space: nowrap;
  }

  button:hover,
  button:focus {
    width: 240px;
    background-position-x: 20px;
    text-align: right;
  }

  button:hover::after,
  button:focus::after {
    width: 160px;
  }

  div {
    position: fixed;
    bottom: 42px;
    right: 42px;
  }

  img {
    max-width: initial;
    position: absolute;
    right: -16px;
    bottom: -22px;
    cursor: pointer;
  }
</style>
