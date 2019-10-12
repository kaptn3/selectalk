<template>
  <div>
    <top/>
    <skills/>
    <div class="container">
      SelecTalk
    </div>
  </div>
</template>

<script>
  import Top from '../components/profile/Top';
  import Skills from '../components/profile/Skills';

  export default {
    components: {
      Top,
      Skills
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
