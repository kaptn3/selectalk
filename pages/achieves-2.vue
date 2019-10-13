<template>
  <div class="achieves">
    <div class="achieves__points">
      <h2>У вас <b>11 400 баллов</b></h2>
      <nuxt-link
        to="achieves-2"
        class="achieves__btn"
      >
        В процессе
      </nuxt-link>
      <nuxt-link
        to="achieves"
        class="achieves__btn"
      >
        Накоплено
      </nuxt-link>
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
  import axios from 'axios';
  import Card from '../components/achieves/Card';

  export default {
    layout: 'lists',
    components: { Card },
    data() {
      return {
        img: [
          '/img/11.png',
          '/img/10.png',
          '/img/9.png',
          '/img/8.png',
          '/img/7.png'
        ],
        achieves: []
      }
    },
    mounted() {
      axios.get(`${process.env.api}achievements`)
        .then((res) => {
          for (let i = 0; i < res.data.length; i++) {
            const data = res.data[i];

            this.achieves.push({
              img: this.img[i],
              title: data.Name,
              points: data.Costs,
              deadline: `Осталось ${data.Costs - 11400} баллов`
            })
          }
        });
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

  .achieves__btn {
    border-color:rgba(53, 53, 71, 0.2);
    color:rgba(53, 53, 71, 0.2);
  }

  .nuxt-link-active {
    background-color: var(--red);
    color: #fff;
    margin-right: 20px;
    margin-left: auto;
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
