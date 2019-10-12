<template>
  <aside>
    <logo/>
    <profile-nav/>
    <sidebar-menu/>
  </aside>
</template>

<script>
  import Logo from './Logo';
  import ProfileNav from './ProfileNav';
  import SidebarMenu from './SidebarMenu';

  export default {
    name: 'Sidebar',
    components: {
      Logo,
      ProfileNav,
      SidebarMenu
    }
  };
</script>

<style scoped>
  aside {
    background-color: var(--black);
    color: var(--white);
  }
</style>
