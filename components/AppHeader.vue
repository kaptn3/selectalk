<template>
  <header
    :class="{ 'profile__header': isProfile }"
    class="header"
  >
    <ul>
      <button/>
      <li
        v-for="item in menu"
        :key="item.name"
      >
        <nuxt-link :to="item.link">
          {{ item.name }}
        </nuxt-link>
      </li>
    </ul>
  </header>
</template>

<script>
  export default {
    name: 'AppHeader',
    props: {
      isProfile: {
        type: Boolean,
        default: false
      }
    },
    data() {
      return {
        menu: [
          {
            link: 'chat',
            name: 'Чат'
          },
          {
            link: 'rating',
            name: 'Рейтинг'
          },
          {
            link: '#',
            name: 'Войти'
          },
        ]
      }
    }
  };
</script>

<style scoped>
  ul {
    display: flex;
    justify-content: flex-end;
    list-style: none;
    padding: 0;
  }

  a {
    display: block;
    padding: 1.5rem;
  }

  button {
    width: 18px;
    height: 18px;
    padding: 0;
    background: url(/icons/search.png) no-repeat center;
    background-size: contain;
    margin: 1.5rem;
  }

  .header {
    background-color: #fff;
    box-shadow: 0 4px 8px rgba(0, 0, 0, .15);
  }

  .profile__header {
    background-color: transparent;
    box-shadow: none;
    z-index: 9;
  }
</style>
