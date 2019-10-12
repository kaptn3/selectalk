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
          <nuxt-link :to="item.link">
            {{ item.name }}
          </nuxt-link>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script>
  import Badge from './Badge';

  export default {
    name: 'ProfileNav',
    components: { Badge },
    data() {
      return {
        menu: [
          {
            name: 'Мои задания',
            link: '#'
          },
          {
            name: 'Мои курсы',
            link: '#'
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
    padding: 2rem 1rem;
    list-style: none;
  }

  li {
    margin-bottom: 3rem;
  }

  a {
    opacity: .7;
  }

  a:hover,
  a:focus {
    opacity: 1;
  }
</style>
