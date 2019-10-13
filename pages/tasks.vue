<template>
  <div class="tasks">
    <add-task/>
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
  import axios from 'axios';
  import AddTask from '../components/tasks/AddTask';
  import TaskItem from '../components/tasks/TaskItem';

  export default {
    layout: 'lists',
    components: {
      TaskItem,
      AddTask
    },
    data() {
      return {
        tasks: []
      }
    },
    mounted() {
      axios.get(`${process.env.api}tasks.json`)
        .then((res) => {
          for (let i = 0; i < res.data.length; i++) {
            const data = res.data[i];
            let date = new Date(data.ExpirationDate);
            date = `${date.getDate()}.${date.getUTCMonth()}.${date.getUTCFullYear()}`;

            axios.get(`${process.env.api}users/${data.ManagerID}.json`)
              .then((res2) => {
                this.tasks.push({
                  title: data.Name,
                  points: data.Score,
                  leader: res2.data.Name,
                  index: data.TaskPriorityID,
                  skills: [
                    '/icons/skills/1.svg',
                    '/icons/skills/2.svg'
                  ],
                  deadline: date
                })
              })
          }
        });
    }
  };
</script>

<style scoped>
  .tasks {
    overflow: hidden;
  }

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
