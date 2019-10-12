<template>
  <div class="profile-nav">
    <nuxt-link
      class="profile-nav__top"
      to="/profile"
    >
      <profile-avatar/>
      <div>
        <span class="profile-nav__name">Тирекс Селектел</span>
        <badge name="Мастер"/>
      </div>
    </nuxt-link>
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
  </div>
</template>

<script>
  import Badge from './Badge';
  import SidebarMenuItem from './SidebarMenuItem';
  import ProfileAvatar from './ProfileAvatar';

  export default {
    name: 'ProfileNav',
    components: {
      Badge,
      SidebarMenuItem,
      ProfileAvatar
    },
    data() {
      return {
        menu: [
          {
            name: 'Мои задания',
            link: '/tasks',
            icon: '/icons/menu/task.png'
          },
          {
            name: 'Мои курсы',
            link: '#',
            icon: '/icons/menu/course.png'
          }
        ]
      };
    }
  };
</script>

<style scoped>
  .profile-nav__top {
    display: flex;
    padding: 1.4rem;
  }

  ul {
    list-style: none;
    padding: 0;
    margin-bottom: 2rem;
  }

  .nuxt-link-active {
    background-color: #3B3A44;
  }
</style>
