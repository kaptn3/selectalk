<template>
  <div>
    <top/>
    <div class="container">
      SelecTalk
    </div>
  </div>
</template>

<script>
  import Top from '../components/profile/Top';

  export default {
    components: {
      Top
    }
  };
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
