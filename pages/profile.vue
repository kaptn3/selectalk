<template>
  <div>
    <top/>
    <skills/>
    <achieves/>
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
