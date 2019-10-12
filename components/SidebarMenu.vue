<template>
  <nav>
    <ul>
      <li
        v-for="item in menu"
        :key="item.name"
      >
        <sidebar-menu-item
          :name="item.name"
          :link="item.link"
          :icon="item.icon"
        />
      </li>
    </ul>
  </nav>
</template>

<script>
  import SidebarMenuItem from './SidebarMenuItem';

  export default {
    name: 'SidebarMenu',
    components: {
      SidebarMenuItem
    },
    data() {
      return {
        menu: [
          {
            link: 'achieves',
            name: 'Достижения',
            icon: '/icons/menu/achievement.png'
          },
          {
            link: 'calendar',
            name: 'Календарь',
            icon: '/icons/menu/calendar.png'
          },
          {
            link: 'docs',
            name: 'Документы',
            icon: ''
          },
          {
            link: 'news',
            name: 'Новости',
            icon: ''
          }
        ]
      }
    }
  }
</script>

<style scoped>
  ul {
    list-style: none;
    padding: 0;
  }
</style>
