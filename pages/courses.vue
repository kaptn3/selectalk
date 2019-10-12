<template>
  <div class="courses">
    <card
      v-for="course in courses"
      :key="course.title"
      :percent="60"
      :img="course.img"
      :title="course.title"
      :type="course.type"
      class="courses__item"
    />
  </div>
</template>

<script>
  import Card from '../components/courses/Card';

  export default {
    layout: 'lists',
    components: { Card },
    data() {
      return {
        courses: [
          {
            img: '/img/1.png',
            title: 'Сетевая безопасность',
            type: 'Упражнение'
          },
          {
            img: '/img/2.png',
            title: 'Размещение серверов',
            type: 'Статья'
          },
          {
            img: '/img/3.png',
            title: 'Общение с поставщиками',
            type: 'Тест'
          },
          {
            img: '/img/4.png',
            title: 'Безопасность в Data центрах',
            type: 'Курс'
          },
          {
            img: '/img/5.png',
            title: 'Сотрудники Data центров',
            type: 'Видео'
          }
        ]
      };
    }
  };
</script>

<style scoped>
  .courses {
    display: flex;
    flex-wrap: wrap;
    align-items: flex-start;
    margin-top: 50px;
  }

  .courses__item {
    width: 31%;
    margin-right: 3%;
    margin-bottom: 40px;
  }

  .courses__item:nth-child(3n) {
    margin-right: 0;
  }
</style>
