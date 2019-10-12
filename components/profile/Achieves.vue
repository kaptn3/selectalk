<template>
  <section>
    <h2>
      Достижения
      <button-plus class="button-plus"/>
    </h2>
    <div class="achieves">
      <achieves-item
        v-for="achieve in achieves"
        :key="achieve.name"
        :name="achieve.name"
        :img="achieve.img"
        :points="achieve.points"
        :deadline="achieve.deadline"
        class="achieves__item"
      />
    </div>
  </section>
</template>

<script>
  import ButtonPlus from '../ButtonPlus';
  import AchievesItem from './AchievesItem';

  export default {
    name: 'Skills',
    components: {
      AchievesItem,
      ButtonPlus
    },
    data() {
      return {
        achieves: [
          {
            name: 'День опозданий',
            img: '/img/achieve.png',
            points: 600,
            deadline: 'До 12 окт 2019 г.'
          }
        ]
      }
    }
  };
</script>

<style scoped>
  section {
    padding: 50px 102px;
  }

  h2 {
    margin-bottom: 35px;
  }

  .achieves {
    display: flex;
    flex-wrap: wrap;
  }

  .achieves__item {
    width: 33%;
  }
</style>
