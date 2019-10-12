<template>
  <div class="task">
    <div class="task__inner">
      <span class="task__index">
        {{ index }}
      </span>
      <span class="task__title">
        {{ title }}
      </span>
      <span class="task__leader">
        {{ leader }}
      </span>
      <span class="task__deadline">
        {{ deadline }}
      </span>
      <div class="task__action">
        <button class="task__cancel"/>
      </div>
      <div class="task__action">
        <button class="task__check"/>
      </div>
    </div>
    <div class="task__awards">
      <div class="task__skills">
        + к навыкам
        <img
          v-for="icon in skills"
          :key="icon"
          :src="icon"
          class="task__skill-icon"
        >
      </div>
      <div class="task__points">
        {{ points }} баллов
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'TaskItem',
    props: {
      index: {
        type: Number,
        required: true
      },
      title: {
        type: String,
        required: true
      },
      leader: {
        type: String,
        required: true
      },
      deadline: {
        type: String,
        required: true
      },
      skills: {
        type: Array,
        required: true
      },
      points: {
        type: Number,
        required: true
      }
    }
  };
</script>

<style scoped>
  .task {
    background-color: #fff;
    height: 116px;
    overflow: hidden;
    cursor: pointer;
    transition: .2s height;
    border-left: 6px solid transparent;
  }

  .task__inner {
    height: 116px;
    display: flex;
    align-items: center;
    padding: 20px;
    font-size: 18px;
  }

  .task__index {
    width: 15%;
  }

  .task__title {
    width: 39%;
  }

  .task__leader {
    width: 15%;
  }

  .task__deadline {
    color: var(--red);
    width: 15%;
    font-weight: 500;
  }

  .task__action {
    width: 8%;
    text-align: center;
    opacity: 0;
    transition: .2s opacity;
  }

  .task__awards {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    border-top: 1px solid rgba(53, 53, 71, .1);
    height: 40px;
  }

  .task__skills {
    opacity: .5;
    font-size: 10px;
    margin-right: 60px;
  }

  .task__skill-icon {
    width: 24px;
  }

  .task__points {
    color: var(--red);
    padding: 10px;
    width: 140px;
    text-align: center;
    border: 1px solid var(--red);
    font-weight: 500;
    font-size: 13px;
  }

  .task:hover,
  .task:focus {
    height: 157px;
    border-left-color: var(--red);
  }

  .task:hover .task__action,
  .task:focus .task__action {
    opacity: 1;
  }

  span {
    opacity: .5;
  }

  button {
    border-radius: 100%;
    width: 39px;
    height: 39px;
  }

  .task__cancel {
    background: url(/icons/cancel.png) no-repeat center;
    border: 1px solid #CDCDCD;
  }

  .task__check {
    background: var(--red) url(/icons/check.png) no-repeat center;
  }
</style>
