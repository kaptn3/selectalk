<template>
  <div>
    <div class="card">
      <img :src="img">
      <div class="card__inner">
        <h2>{{ title }} ({{ points }} баллов) </h2>
        <span>{{ deadline }}</span>
      </div>
    </div>
    <button v-if="pickUp">
      Забрать
    </button>
  </div>
</template>

<script>
  export default {
    name: 'Card',
    props: {
      img: {
        type: String,
        required: true
      },
      title: {
        type: String,
        required: true
      },
      deadline: {
        type: String,
        required: true
      },
      points: {
        type: Number,
        required: true
      },
      pickUp: {
        type: Boolean,
        default: false
      }
    }
  };
</script>

<style scoped>
  .card {
    background-color: #fff;
    font-size: 14px;
    text-align: center;
    cursor: pointer;
  }

  .card:hover,
  .card:focus {
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.15);
  }

  .card__inner {
    padding: 20px;
  }

  h2 {
    font-size: 14px;
  }

  button {
    width: 100%;
    background-color: var(--red);
    margin-top: 10px;
    font-size: 13px;
    color: #fff;
    padding: 14px;
    box-shadow: 0 4px 4px rgba(0, 0, 0, .25);
  }
</style>
