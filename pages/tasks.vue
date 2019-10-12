<template>
  <div>
    <header class="tasks__header">
      <span class="tasks__index">Приоритет</span>
      <span class="tasks__title">Тема задания</span>
      <span class="tasks__leader">Руководитель</span>
      <span class="tasks__deadline">Дедлайн</span>
    </header>
    <task-item
      v-for="task in tasks"
      :key="task.index"
      :index="task.index"
      :title="task.title"
      :leader="task.leader"
      :deadline="task.deadline"
      :skills="task.skills"
      :points="task.points"
      class="tasks__item"
    />
  </div>
</template>

<script>
  import TaskItem from '../components/tasks/TaskItem';

  export default {
    layout: 'lists',
    components: {
      TaskItem
    },
    data() {
      return {
        tasks: [
          {
            index: 1,
            title: 'Оформить ТЗ',
            leader: 'Никита',
            deadline: '13.10.2019',
            skills: [
              '/icons/skills/1.svg',
              '/icons/skills/2.svg'
            ],
            points: 200
          },
          {
            index: 2,
            title: 'Оформить ТЗ',
            leader: 'Никита',
            deadline: '13.10.2019',
            skills: [
              '/icons/skills/1.svg',
              '/icons/skills/2.svg'
            ],
            points: 200
          },
          {
            index: 3,
            title: 'Оформить ТЗ',
            leader: 'Никита',
            deadline: '13.10.2019',
            skills: [
              '/icons/skills/1.svg',
              '/icons/skills/2.svg'
            ],
            points: 200
          }
        ]
      }
    }
  };
</script>

<style scoped>
  .tasks__header {
    display: flex;
    align-items: center;
    padding: 1rem 0;
  }

  .tasks__item {
    margin-bottom: 1rem;
  }

  span {
    color: #B7B6B6;
    font-size: 11px;
    display: block;
  }

  .tasks__index {
    width: 15%;
  }

  .tasks__title {
    width: 39%;
  }

  .tasks__leader {
    width: 15%;
  }

  .tasks__deadline {
    width: 15%;
  }
</style>
