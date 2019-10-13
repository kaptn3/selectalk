<template>
  <div>
    <div class="top">
      <div class="top__inner">
        <h1><b>Топ 3</b> сотрудника</h1>
        <div class="top__item">
          <img src="/img/2-place.png" alt="">
          <span>Николай</span>
          <span class="top__points">26 400 баллов</span>
        </div>
        <div class="top__item">
          <img src="/img/1-place.png" alt="">
          <span>Даниил</span>
          <span class="top__points">28 000 баллов</span>
        </div>
        <div class="top__item">
          <img src="/img/3-place.png" alt="">
          <span>Максим</span>
          <span class="top__points">21 900 баллов</span>
        </div>
      </div>
    </div>
    <rest/>
  </div>
</template>

<script>
  import Rest from '../components/rating/Rest';

  export default {
    components: { Rest }
  };
</script>

<style scoped>
  .top {
    background: var(--red) url(/img/bg.jpg) 36% 60% no-repeat;
    background-size: cover;
    height: 540px;
    margin-top: -72px;
    display: flex;
    align-items: flex-end;
    padding: 40px 102px;
    text-align: center;
  }

  .top__inner {
    background-color: #fff;
    opacity: .8;
    height: 366px;
    padding: 40px;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: flex-end;
    flex-wrap: wrap;
  }

  h1 {
    font-size: 24px;
    font-weight: 400;
    width: 100%;
  }

  img {
    max-width: initial;
  }

  .top__item {
    width: 188px;
  }

  span {
    font-size: 18px;
    opacity: .6;
    display: block;
  }

  .top__points {
    color: var(--red);
    font-size: 14px;
  }

  .top__item:nth-child(3) span {
    opacity: 1;
  }
</style>
