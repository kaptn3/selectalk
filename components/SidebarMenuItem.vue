<template>
  <nuxt-link :to="link">
    <div class="icon">
      <img :src="icon" alt="">
    </div>
    {{ name }}
  </nuxt-link>
</template>

<script>
  export default {
    name: 'SidebarMenuItem',
    props: {
      name: {
        type: String,
        required: true
      },
      link: {
        type: String,
        required: true
      },
      icon: {
        type: String,
        required: true
      }
    }
  };
</script>

<style scoped>
  a {
    opacity: .7;
    display: block;
    padding: .6rem 1.4rem;
  }

  a:hover,
  a:focus {
    opacity: 1;
    background-color: #3B3A44;
  }

  .icon {
    opacity: 0;
    width: 46px;
    height: 46px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    margin-right: 10px;
    vertical-align: middle;
  }

  a:hover .icon,
  a:focus .icon {
    opacity: 1;
  }
</style>
