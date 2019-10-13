<template>
  <div class="carousel">
    <nuxt-link
      to="profile"
      class="carousel__close"
    >
      <img src="/icons/close.png" alt="">
    </nuxt-link>
    <h1>Добро пожаловать!</h1>
    <div
      :style="{ 'transform': `translateX(${current * -969}px)` }"
      class="carousel__box"
    >
      <div
        v-for="(slide, index) in slides"
        :key="index"
        class="carousel__item"
      >
        <img :src="slide.img" alt="">
        <p>{{ slide.text }}</p>
      </div>
    </div>
    <div class="buttons">
      <nuxt-link
        to="profile"
        class="btn skip"
      >
        Пропустить
      </nuxt-link>
      <div>
        <button
          v-for="(slide, index) in slides"
          :key="index"
          :class="{ 'active': current === index }"
          class="button"
          @click="current = index"
        />
      </div>
      <button
        :disabled="current === slides.length - 1"
        class="btn next"
        @click="current++"
      >
        Далее
      </button>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Carousel',
    data() {
      return {
        current: 0,
        slides: [
          {
            text: 'Шаг 1. Выполняйте задания по приоритетам. Не бойтесь дедлайнов и повышайте навыки.',
            img: '/img/1.gif'
          },
          {
            text: 'Шаг 2. Проходите тесты, читайте статьи, делайте упражнения, смотрите обучающие видео и получайте баллы.',
            img: '/img/2.gif'
          },
          {
            text: 'Шаг 3. Копите баллы и получайте награды от Дня Опозданий до командировок за границу.',
            img: '/img/screen-3.png'
          },
          {
            text: 'Шаг 4. Отслеживайте прогресс подчиненных по заданиям, курсам и достижениям в виде графиков.',
            img: '/img/screen-4.png'
          },
          {
            text: 'Шаг 5. Остались вопросы? Пишите чат боту в любое время.',
            img: '/img/5.gif'
          }
        ]
      }
    }
  }
</script>

<style scoped>
  h1 {
    font-size: 24px;
  }

  .carousel {
    text-align: center;
    display: flex;
    justify-content: space-between;
    flex-direction: column;
    overflow: hidden;
    height: 100%;
  }

  .carousel__close {
    position: absolute;
    display: block;
    width: 46px;
    height: 46px;
    top: 30px;
    right: 30px;
  }

  .carousel__box {
    width: 500%;
    display: flex;
    transition: .4s transform;
  }

  .carousel__item {
    width: 969px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-direction: column;
  }

  .carousel__item:nth-child(4) img {
    width: 550px;
  }

  img {
    width: 700px;
  }

  .buttons {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .button {
    width: 20px;
    height: 20px;
    border: 1px solid rgba(53, 53, 71, .2);
    border-radius: 50%;
    margin: 0 11px;
  }

  p {
    color: #9B9B9B;
    font-size: 18px;
    margin: 15px;
  }

  .button.active {
    background-color: rgba(53, 53, 71, .2);
  }

  .btn {
    width: 224px;
    border-radius: 8px;
    font-size: 13px;
    font-weight: 500;
    padding: 12px;
    border: 1px solid transparent;
  }

  .skip {
    color: rgba(53, 53, 71, 0.3);
    border-color: rgba(53, 53, 71, 0.3);
  }

  .next {
    color: #fff;
    background-color: var(--red);
  }
</style>
