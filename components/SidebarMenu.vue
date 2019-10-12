<template>
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
</template>

<script>
  export default {
    name: 'SidebarMenu',
    data() {
      return {
        menu: [
          {
            link: '#',
            name: 'Достижения'
          },
          {
            link: '#',
            name: 'Календарь'
          },
          {
            link: '#',
            name: 'Документы'
          },
          {
            link: '#',
            name: 'Новости'
          }
        ]
      }
    }
  }
</script>

<style scoped>
  ul {
    list-style: none;
    padding: 1rem;
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
