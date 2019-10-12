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
  </div>
</template>

<script>
  import Badge from './Badge';
  import ProfileAvatar from './ProfileAvatar';

  export default {
    name: 'ProfileNav',
    components: {
      Badge,
      ProfileAvatar
    }
  };
</script>

<style scoped>
  .profile-nav__top {
    display: flex;
    padding: 1.4rem;
  }
</style>
