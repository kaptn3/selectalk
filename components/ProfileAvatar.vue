<template>
  <div
    :style="{ 'width': size, 'height': size }"
  >
    <img :src="avatar" alt="">
  </div>
</template>

<script>
  export default {
    name: 'ProfileAvatar',
    props: {
      size: {
        type: String,
        default: '46px'
      }
    },
    data() {
      return {
        avatar: '/img/photo.jpg'
      }
    }
  };
</script>

<style scoped>
  div {
    border-radius: 100%;
    flex-shrink: 0;
    overflow: hidden;
    margin-right: 10px;
  }

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
</style>
