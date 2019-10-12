<template>
  <div class="app">
    <sidebar class="app__sidebar"/>
    <app-header class="app__header"/>
    <main class="app__main">
      <nuxt/>
    </main>
  </div>
</template>

<script>
  import Sidebar from '../components/Sidebar';
  import AppHeader from '../components/AppHeader';

  export default {
    components: {
      Sidebar,
      AppHeader
    }
  }
</script>

<style>
  .app {
    display: grid;
    grid-template-areas:
      'sidebar header'
      'sidebar main';
    grid-template-columns: 245px 1fr;
  }

  .app__sidebar {
    grid-area: sidebar;
  }

  .app__header {
    grid-area: header;
  }

  .app__main {
    grid-area: main;
  }
</style>
