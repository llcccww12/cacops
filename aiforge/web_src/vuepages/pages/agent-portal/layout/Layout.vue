<template>
  <div class="main-layout">
    <div id="nav-left" class="left-menu">
      <Menu />
    </div>
    <div class="right-content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Menu from '../components/Menu.vue';

export default {
  name: 'AgentPortalLayout',
  data() {
    return {};
  },
  components: { Menu },
  methods: {},
  beforeCreate() { },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.main-layout {
  display: flex;

  .left-menu {
    position: relative;
    interpolate-size: allow-keyword;
    transition: all .2s ease-in-out;
    width: 200px;

    &.collapsed {
      width: 0;
      overflow-x: hidden;
    }
  }

  .right-content {
    flex: 1;
    width: 0;
    height: calc(100vh - 62px);
    position: relative;
  }
}

@media only screen and (max-width: 768px) {
  .main-layout {
    .left-menu {
      width: auto;
    }
  }
}
</style>