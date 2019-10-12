<template>
  <div class="app">
    <sidebar class="app__sidebar"/>
    <app-header
      :is-profile="$route.name == 'profile'"
      class="app__header"
    />
    <main class="app__main">
      <nuxt/>
    </main>
    <chat/>
  </div>
</template>

<script>
  import Sidebar from '../components/Sidebar';
  import AppHeader from '../components/AppHeader';
  import Chat from '../components/Chat';

  export default {
    components: {
      Sidebar,
      AppHeader,
      Chat
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
    grid-template-rows: 72px 1fr;
    background-color: #E5E5E5;
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
