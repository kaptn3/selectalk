<template>
  <form>
    <a-input
      class="add-task__input"
      label="Приоритет"
    />
    <a-input
      class="add-task__input"
      label="Дедлайн"
    />
    <a-input
      class="add-task__input"
      label="Выбрать руководителя"
    />
    <button>
      Добавить задание
    </button>
  </form>
</template>

<script>
  import AInput from './AInput';

  export default {
    name: 'AddTask',
    components: { AInput }
  };
</script>

<style scoped>
  form {
    background-color: #fff;
    padding: 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin: 2rem 0;
  }

  button {
    background-color: var(--red);
    color: var(--white);
    border-radius: 8px;
    margin-left: auto;
    padding: 10px;
  }

  button,
  .add-task__input {
    width: 20%;
  }

  .add-task__input {
    margin-right: 50px;
  }
</style>
