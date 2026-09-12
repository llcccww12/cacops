<template>
  <div class="main-layout">
    <div id="nav-left" class="left-menu">
      <Menu />
    </div>
    <div class="right-content">
      <div class="bg-1"></div>
      <div class="bg-2"></div>
      <div class="bg-3"></div>
      <div class="bg-4"></div>
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Menu from '../components/Menu.vue';

export default {
  name: 'Layout',
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
    interpolate-size: allow-keywords;
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

    .bg-1 {
      position: absolute;
      width: 100%;
      height: 100%;
      top: 0;
      z-index: -2;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3CradialGradient%20id%3D%221%22%20cx%3D%220%22%20cy%3D%220%22%20r%3D%221%22%20gradientTransform%3D%22matrix(0.42999999999999994%2C%200.01200000000000001%2C%20-0.009048275494276803%2C%200.6127499999999999%2C%200.154%2C%200.447)%22%3E%3Cstop%20stop-color%3D%22%23defaff%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23f6feff%22%20stop-opacity%3D%220.3%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FradialGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }

    .bg-2 {
      position: absolute;
      width: 100%;
      height: 100%;
      top: 0;
      z-index: -2;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3CradialGradient%20id%3D%221%22%20cx%3D%220%22%20cy%3D%220%22%20r%3D%221%22%20gradientTransform%3D%22matrix(0.533%2C%200.009000000000000001%2C%20-0.004843208514568158%2C%200.542061%2C%20-0.002%2C%200.003)%22%3E%3Cstop%20stop-color%3D%22%23e8e4ff%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23eae8f9%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23eae8f9%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FradialGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }

    .bg-3 {
      position: absolute;
      right: 0;
      top: 0;
      z-index: -2;
      width: 1200px;
      height: 227px;
      background: url("/img/model/modelbase-bg.png");
    }

    .bg-4 {
      position: absolute;
      width: 100%;
      height: 100%;
      top: 0;
      z-index: -2;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3CradialGradient%20id%3D%221%22%20cx%3D%220%22%20cy%3D%220%22%20r%3D%221%22%20gradientTransform%3D%22matrix(-0.003000000000000119%2C%200.8010000000000002%2C%20-0.650838234521332%2C%20-0.004050000000000161%2C%201.014%2C%20-0.025)%22%3E%3Cstop%20stop-color%3D%22%237bb2ff%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23bdabff%22%20stop-opacity%3D%220.41%22%20offset%3D%220.35%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%2392ffaf%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FradialGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }
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
