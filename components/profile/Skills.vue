<template>
  <section>
    <h2>
      Навыки
      <button-plus class="button-plus"/>
    </h2>
    <div class="skills">
      <skills-item
        v-for="skill in skills"
        :key="skill.name"
        :percent="skill.percent"
        :icon="skill.icon"
        :name="skill.name"
        class="skills__item"
      />
    </div>
  </section>
</template>

<script>
  import ButtonPlus from '../ButtonPlus';
  import SkillsItem from './SkillsItem';

  export default {
    name: 'Skills',
    components: {
      SkillsItem,
      ButtonPlus
    },
    data() {
      return {
        skills: [
          {
            icon: '/icons/skills/1.svg',
            name: 'Аккуратность',
            percent: 22
          },
          {
            icon: '/icons/skills/2.svg',
            name: 'Ответственность',
            percent: 56
          },
          {
            icon: '/icons/skills/3.svg',
            name: 'Общительность',
            percent: 100
          },
          {
            icon: '/icons/skills/4.svg',
            name: 'Пунктуальность',
            percent: 31
          },
          {
            icon: '/icons/skills/5.svg',
            name: 'Наблюдательность',
            percent: 76
          }
        ]
      }
    }
  };
</script>

<style scoped>
  section {
    padding: 50px 102px;
  }

  h2 {
    font-weight: 400;
    margin-bottom: 35px;
  }

  .skills {
    display: flex;
    flex-wrap: wrap;
  }

  .skills__item {
    width: 33%;
    margin-bottom: 50px;
  }

  .button-plus {
    margin-left: 50px;
  }
</style>
