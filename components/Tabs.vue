<template>
  <nav>
    <ul>
      <li
        v-for="tab in tabs"
        :key="tab.name"
        class="tabs__item"
      >
        <tabs-item
          :link="tab.link"
          :name="tab.name"
        />
      </li>
    </ul>
  </nav>
</template>

<script>
  import TabsItem from './TabsItem';

  export default {
    name: 'Tabs',
    components: { TabsItem },
    data() {
      return {
        tabs: [
          {
            name: 'Мои задания',
            link: 'tasks'
          },
          {
            name: 'Мои курсы',
            link: 'courses'
          },
          {
            name: 'Достижения',
            link: 'achieves'
          },
          {
            name: 'Календарь',
            link: 'calendar'
          }
        ]
      }
    }
  };
</script>

<style scoped>
  ul {
    display: flex;
    list-style: none;
    padding: 0;
    border-bottom: 1px solid rgba(53, 53, 71, .1);
    padding: 0 102px;
  }

  .tabs__item {
    width: 25%;
    margin-bottom: -2px;
  }
</style>
