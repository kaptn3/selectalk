<template>
  <div class="profile-nav">
    <div class="profile-nav__top">
      <div class="profile-nav__avatar">
        <img src="/photo.jpg" alt="">
      </div>
      <div>
        <span class="profile-nav__name">Тирекс Селектел</span>
        <badge name="Мастер"/>
      </div>
    </div>
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

  export default {
    name: 'ProfileNav',
    components: {
      Badge,
      SidebarMenuItem
    },
    data() {
      return {
        menu: [
          {
            name: 'Мои задания',
            link: '#',
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
    background-color: #3B3A44;
    display: flex;
    padding: 1.4rem;
  }

  .profile-nav__avatar {
    border-radius: 100%;
    width: 46px;
    height: 46px;
    flex-shrink: 0;
    overflow: hidden;
    margin-right: 10px;
  }

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  ul {
    list-style: none;
    padding: 0;
    margin-bottom: 2rem;
  }
</style>
