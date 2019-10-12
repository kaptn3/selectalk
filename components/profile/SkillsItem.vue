<template>
  <div class="skill">
    <div class="skill__icon">
      <img :src="icon" alt="">
    </div>
    <div class="skill__count">
      <h3>{{ name }}</h3>
      <span class="skill__progress">
        <span
          :style="{ 'width': `${percent}%` }"
          class="skill__percent"
        >
          {{ percent }} %
        </span>
      </span>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'SkillsItem',
    props: {
      icon: {
        type: String,
        required: true
      },
      name: {
        type: String,
        required: true
      },
      percent: {
        type: Number,
        default: 0
      }
    }
  };
</script>

<style scoped>
  .skill {
    display: flex;
    align-items: center;
  }

  .skill__icon {
    margin-right: 14px;
  }

  .skill__progress {
    display: block;
    position: relative;
    background-color: #fff;
    width: 162px;
    height: 19px;
    border-radius: 8px;
    border: 1px solid #EBEBEB;
  }

  .skill__percent {
    background-color: var(--black);
    display: block;
    position: absolute;
    top: -1px;
    left: 0;
    font-size: 10px;
    padding: 2px;
    text-align: center;
    color: var(--white);
    border-radius: 8px;
  }
</style>
