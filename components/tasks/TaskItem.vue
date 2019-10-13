<template>
  <div
    class="task"
    :style="{ 'transform': `translateX(${left})`, 'display': display }"
  >
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
        <button
          class="task__check"
          @click="check"
        />
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
    },
    data() {
      return {
        left: 0,
        display: 'block'
      }
    },
    methods: {
      check() {
        this.left = '100%';
        setTimeout(() => {
          this.display = 'none';
        }, 300);
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
    border-left: 6px solid transparent;
    transition: .3s all;
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
    position: relative;
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

  .task:hover span,
  .task:focus span {
    opacity: 1;
  }

  .task:hover .task__action,
  .task:focus .task__action {
    opacity: 1;
  }

  span {
    opacity: .5;
  }

  button {
    border-radius: 20px;
    width: 40px;
    height: 40px;
    transition: .2s background-poisiton-x, .2s width;
    position: absolute;
    transform: translateY(-50%);
  }

  button:hover,
  button:focus {
    width: 132px;
    text-align: right;
    background-position-x: 10px;
  }

  button::after {
    white-space: nowrap;
    width: 0;
    display: inline-block;
    overflow: hidden;
    transition: .2s width;
    left: 10px;
  }

  button:hover::after,
  button:focus::after {
    width: 84px;
  }

  .task__cancel {
    background: #fff url(/icons/cancel.png) no-repeat center;
    border: 1px solid #CDCDCD;
    left: 0;
  }

  .task__check {
    background: var(--red) url(/icons/check.png) no-repeat center;
    right: 20px;
  }

  .task__cancel::after {
    content: 'Отменить';
    opacity: .3;
  }

  .task__check::after {
    content: 'Выполнить';
    color: #fff;
  }

  .task__cancel:hover,
  .task__cancel:focus {
    z-index: 9;
  }
</style>
