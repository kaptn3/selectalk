<template>
  <div class="card">
    <div class="card__img">
      <img :src="img">
    </div>
    <div class="card__text">
      <h3 class="card__title">
        {{ title }}
      </h3>
      <span class="card__type">
        {{ type }}
      </span>
    </div>
    <div class="card__progress">
      {{ percent }}%
      <div
        :style="{ 'transform': `rotate(${360 * percent / 100}deg)` }"
        class="card__circle"
      />
      <div class="card__circle"/>
    </div>
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
      type: {
        type: String,
        required: true
      },
      percent: {
        type: Number,
        required: true
      }
    }
  };
</script>

<style scoped>
  .card {
    background-color: var(--white);
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    font-size: 14px;
    transition: .2s box-shadow;
    cursor: pointer;
  }

  .card:hover {
    box-shadow: 0 4px 8px rgba(0, 0, 0, .15);
  }

  .card__img {
    width: 100%;
    height: 170px;
  }

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .card__progress {
    font-size: 11px;
    width: 28px;
    height: 28px;
    border: 1px solid var(--red);
    border-radius: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
    margin: 20px 20px 20px 0;
  }

  .card__circle {
    width: 26px;
    height: 26px;
    border: 1px solid var(--red);
    border-radius: 100%;
    flex-shrink: 0;
    position: absolute;
    clip: rect(0, 13px, 26px, 0);
    left: 0;
    top: 0;
    transform: rotate(180deg);
  }

  .card__text {
    padding: 20px 0 20px 20px;
  }

  span {
    opacity: .5;
  }
</style>
