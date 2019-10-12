<template>
  <section class="top">
    <div class="top__wrapper">
      <profile-avatar size="104px"/>
      <div class="top__text">
        <span class="top__name">Тирекс Селектел</span>
        <span class="top__points">11400 баллов</span>
        <span class="top__location">Санкт-Петербург</span>
        <div class="top__company">
          <div>
            <span class="top__title">Должность</span>
            <span>Маскот</span>
          </div>
          <div>
            <span class="top__title">Отдел</span>
            <span>Реклама</span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
  import ProfileAvatar from '../ProfileAvatar';

  export default {
    name: 'Top',
    components: { ProfileAvatar }
  };
</script>

<style scoped>
  .top {
    height: 337px;
    background: var(--red) url(/img/bg.jpg) 36% 60% no-repeat;
    margin-top: -72px;
    display: flex;
    align-items: flex-end;
    padding: 12px 102px;
    color: var(--white);
  }

  .top__wrapper {
    display: flex;
    align-items: flex-start;
  }

  .top__text {
    margin-left: 40px;
  }

  .top__name {
    color: var(--black);
    font-weight: 500;
    font-size: 1rem;
  }

  .top__points::before,
  .top__location::before {
    content: '';
    display: inline-block;
    width: 11px;
    height: 15px;
    vertical-align: middle;
    margin-right: 8px;
  }

  .top__points::before {
    background: url(/icons/achieve.png) center no-repeat;
  }

  .top__location::before {
    background: url(/icons/map.png) center no-repeat;
  }

  span {
    font-size: 14px;
    display: block;
  }

  .top__company {
    margin-top: 50px;
    display: flex;
  }

  .top__company div:first-child {
    padding-right: 50px;
    border-right: 1px solid rgba(216, 216, 216, .7);
    margin-right: 50px;
  }

  .top__title {
    font-size: 10px;
    color: #D8D8D8;
  }
</style>
