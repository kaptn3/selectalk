<template>
  <div class="achieves">
    <div class="achieves__points">
      <h2>У вас <b>11 400 баллов</b></h2>
      <button class="achieves__btn">
        В процессе
      </button>
      <button class="achieves__btn">
        Накоплено
      </button>
    </div>
    <div class="achieves__list">
      <card
        v-for="achieve in achieves"
        :key="achieve.title"
        :title="achieve.title"
        :deadline="achieve.deadline"
        :img="achieve.img"
        :points="achieve.points"
        :pick-up="achieve.pickUp"
        class="achieves__item"
      />
    </div>
  </div>
</template>

<script>
  import Card from '../components/achieves/Card';

  export default {
    layout: 'lists',
    components: { Card },
    data() {
      return {
        achieves: [
          {
            title: 'День опозданий',
            img: '/img/achieve.png',
            points: 600,
            deadline: 'До 12 окт 2019 г.'
          },
          {
            title: 'Китайский язык',
            img: '/img/6.png',
            points: 11000,
            deadline: 'До 10 окт 2019 г.',
            pickUp: true
          }
        ]
      }
    }
  };
</script>

<style scoped>
  .achieves {
    margin-top: 40px;
  }

  .achieves__points {
    display: flex;
    align-items: center;
  }

  .achieves__btn {
    font-size: 13px;
    font-weight: 500;
    padding: 10px;
    border-radius: 8px;
    border: 1px solid transparent;
  }

  .achieves__btn:first-of-type {
    background-color: var(--red);
    color: #fff;
    margin-right: 20px;
    margin-left: auto;
  }

  .achieves__btn:nth-of-type(2) {
    border-color:rgba(53, 53, 71, 0.2);
    color:rgba(53, 53, 71, 0.2);
  }

  .achieves__list {
    display: flex;
    flex-wrap: wrap;
    margin-top: 40px;
  }

  .achieves__item {
    width: 31%;
    margin-right: 2%;
  }

  h2 {
    font-size: 24px;
    font-weight: 400;
  }
</style>
