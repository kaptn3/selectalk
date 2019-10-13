<template>
  <div class="rest">
    <div
      v-for="item in rating"
      :key="item.place"
      class="rest__item"
    >
      <span class="rest__place">{{ item.place }} место</span>
      <div class="rest__img">
        <img :src="item.img" alt="">
      </div>
      <span class="rest__name">{{ item.name }}</span>
      <div class="rest__points">
        {{ item.points }}
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Rest',
    data() {
      return {
        rating: [
          {
            place: 4,
            img: '/img/photo-4.png',
            name: 'Кирилл',
            points: '20400'
          },
          {
            place: 4,
            img: '/img/photo-5.png',
            name: 'Владимир',
            points: '20400'
          },
          {
            place: 4,
            img: '/img/photo.jpg',
            name: 'Тирекс',
            points: '20400'
          },
          {
            place: 4,
            img: '/img/photo-7.png',
            name: 'Николай',
            points: '20400'
          }
        ]
      }
    }
  }
</script>

<style scoped>
  .rest {
    padding: 20px 102px;
  }

  .rest__item {
    background-color: #fff;
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 20px;
    height: 98px;
  }

  .rest__img {
    width: 65px;
    height: 65px;
    margin-right: 22px;
  }

  .rest__name {
    font-size: 24px;
    width: 200px;
  }

  .rest__points {
    border: 1px solid var(--red);
    font-size: 14px;
    color: var(--red);
    padding: 9px;
    width: 141px;
    text-align: center;
  }

  .rest__place {
    font-weight: 500;
    font-size: 18px;
    margin-right: 50px;
  }
</style>
