<template>
  <div>
    <img :src="img" alt="">
    <h3>{{ name }} ({{ points }} баллов)</h3>
    <span>{{ deadline }}</span>
  </div>
</template>

<script>
  export default {
    name: 'AchievesItem',
    props: {
      name: {
        type: String,
        required: true
      },
      points: {
        type: Number,
        required: true
      },
      deadline: {
        type: String,
        required: true
      },
      img: {
        type: String,
        required: true
      }
    }
  };
</script>

<style scoped>
  div {
    background: #fff;
    box-shadow: 0 4px 8px rgba(0, 0, 0, .15);
    text-align: center;
    font-size: 14px;
  }

  img {
    margin-bottom: 22px;
  }

  span {
    color: rgba(0, 0, 0, .5);
  }
</style>
