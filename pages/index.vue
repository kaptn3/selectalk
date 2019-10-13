<template>
  <div class="index__modal">
    <carousel/>
  </div>
</template>

<script>
  import Carousel from '../components/Carousel';

  export default {
    components: { Carousel },
    layout: 'index'
  };
</script>

<style scoped>
  .index__modal {
    background-color: #fff;
    width: 100%;
    height: 730px;
    padding: 80px 102px;
    position: relative;
  }
</style>
