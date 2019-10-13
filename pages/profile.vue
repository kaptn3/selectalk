<template>
  <div>
    <top/>
    <skills/>
    <achieves/>
    <img src="/img/wall.svg" alt="">
  </div>
</template>

<script>
  import Top from '../components/profile/Top';
  import Skills from '../components/profile/Skills';
  import Achieves from '../components/profile/Achieves';

  export default {
    components: {
      Top,
      Skills,
      Achieves
    }
  };
</script>

<style scoped>
  img {
    margin: 50px 0 50px 102px;
  }
</style>
