<template>
  <div>
    <div class="bg-container">
      <div class="bg">
        <div class="bg-01"></div>
        <div class="bg-02"></div>
        <div class="bg-03"></div>
        <div class="bg-04"></div>
      </div>
      <div class="title-c">
        <div class="title-main">科技创新2030</div>
        <div class="title-second">新一代人工智能重大项目成果展示</div>
      </div>
      <div class="menu">
        <TopMenu :menu="menu"></TopMenu>
      </div>
    </div>
  </div>
</template>

<script>
import TopMenu from './TopMenu.vue';

export default {
  name: "TopHeader",
  props: {
    menu: { type: String, default: '' },
  },
  components: {
    TopMenu
  },
  data() {
    return {
      list: [],
    };
  },
  methods: {},
  mounted() { },
};
</script>

<style scoped lang="less">
.bg-container {
  height: 200px;
  width: 100%;
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.9059999999999998%2C%20-0.8390000000000001%2C%200.016184413580246918%2C%200.9059999999999998%2C%20-0.183%2C%200.683)%22%3E%3Cstop%20stop-color%3D%22%233bb6fe%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%232250cf%22%20stop-opacity%3D%221%22%20offset%3D%220.61%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%237c23b4%22%20stop-opacity%3D%221%22%20offset%3D%220.99%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  overflow: hidden;
  position: relative;

  .bg {
    position: absolute;
    width: 100%;
    height: 100%;
    left: 67%;
    top: 23%;

    .bg-01 {
      position: absolute;
      top: -296.85px;
      left: -100.21px;
      width: 11.52px;
      height: 430.36px;
      transform: rotate(46deg);
      border-radius: 45px;
      display: flex;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(2.6570000000000005%2C%20-1.366%2C%201906.3802884596842%2C%202.6570000000000005%2C%20-0.714%2C%201.313)%22%3E%3Cstop%20stop-color%3D%22%2315cf16%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ed6bc9%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }

    .bg-02 {
      position: absolute;
      top: -130.68px;
      left: -213.01px;
      width: 3.64px;
      height: 430.36px;
      transform: rotate(46deg);
      border-radius: 45px;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(2.6570000000000005%2C%20-1.366%2C%2019094.65120710059%2C%202.6570000000000005%2C%20-0.714%2C%201.313)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ed6bc9%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }

    .bg-03 {
      position: absolute;
      top: -178.07px;
      left: -50.045px;
      width: 111.15px;
      height: 306.28px;
      transform: rotate(46deg);
      border-radius: 63px;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(2.6570000000000005%2C%20-1.366%2C%2010.372156049382713%2C%202.6570000000000005%2C%20-0.714%2C%201.313)%22%3E%3Cstop%20stop-color%3D%22%23f8da77%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ed6bc9%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }

    .bg-04 {
      position: absolute;
      top: -218.72px;
      left: 239.81px;
      width: 20.38px;
      height: 306.28px;
      transform: rotate(46deg);
      border-radius: 45px;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(2.6570000000000005%2C%20-1.366%2C%20308.5173674049777%2C%202.6570000000000005%2C%20-0.714%2C%201.313)%22%3E%3Cstop%20stop-color%3D%22%2377f8f8%22%20stop-opacity%3D%220.6%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ed6bc9%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }
  }

  .title-c {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 1;

    .title-main {
      text-align: center;
      color: rgba(255, 230, 21, 1);
      font-weight: 700;
      text-shadow: rgb(34 7 94) 0px 2px 6px;
      line-height: 39px;
      font-size: 28px;
      margin: 26px 0 4px;
    }

    .title-second {
      text-align: center;
      font-weight: 700;
      text-shadow: rgb(34 7 94) 0px 2px 6px;
      color: rgb(255, 255, 255);
      font-size: 36px;
      font-weight: 500;
      line-height: 50px;
    }
  }

  .menu {
    position: absolute;
    bottom: 0;
    height: 40px;
    width: 100%;
    z-index: 2;
    padding: 0 140px;
  }
}
</style>
