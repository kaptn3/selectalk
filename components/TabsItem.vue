<template>
  <nuxt-link :to="link">
    {{ name }}
  </nuxt-link>
</template>

<script>
  export default {
    name: 'TabsItem',
    props: {
      name: {
        type: String,
        required: true
      },
      link: {
        type: String,
        required: true
      }
    }
  };
</script>

<style scoped>
  a {
    display: block;
    border-bottom: 4px solid transparent;
    text-align: center;
    font-size: 14px;
    padding: .7rem;
  }

  .nuxt-link-active {
    border-bottom-color: var(--red);
  }
</style>
