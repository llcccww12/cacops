<template>
  <div>
    <div class="bg-container">
      <div class="bg-00"></div>
      <div class="bg-01"></div>
      <div class="bg-02"></div>
      <div class="bg-03"></div>
      <div class="bg-04">
        <div class="bg-04-1"></div>
        <div class="bg-04-2"></div>
        <div class="bg-04-3"></div>
        <div class="bg-04-4"></div>
        <div class="bg-04-5"></div>
        <div class="bg-04-6"></div>
      </div>
      <div class="content ui container">
        <slot name="default"></slot>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "TopHeader",
  props: {
    // menu: { type: String, default: '' },
  },
  components: {

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
  overflow: hidden;
  position: relative;

  .bg-00 {
    position: absolute;
    right: 0;
    top: 0;
    height: 100%;
    width: 50%;
    background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.11899999999999993%2C%201.217%2C%20-0.0901857098765432%2C%200.11899999999999993%2C%200.269%2C%20-0.22)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.47%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23e5e7eb%22%20stop-opacity%3D%220.3%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  }

  .bg-01 {
    position: absolute;
    top: 0;
    left: 50%;
    z-index: 1;
    width: 2501px;
    height: 400px;
    border-radius: 100%;
    transform: translate(-50%, -50%);
    background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.06299999999999999%2C%201.5829999999999997%2C%20-0.06686206253204435%2C%200.06299999999999999%2C%200.339%2C%20-0.115)%22%3E%3Cstop%20stop-color%3D%22%23c4e4db%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%236dffd0%22%20stop-opacity%3D%221%22%20offset%3D%220.43%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%239eb3ff%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  }

  .bg-02 {
    position: absolute;
    top: -180px;
    right: 1%;
    width: 70%;
    height: 410px;
    border-radius: 100%;
    transform-origin: 583.5px 205px;
    z-index: 2;
    background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.002000000000000203%2C%200.9999999999999999%2C%20-0.12343149845545413%2C%20-0.002000000000000203%2C%200%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.2%22%20offset%3D%220.34%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220%22%20offset%3D%220.76%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  }

  .bg-03 {
    position: absolute;
    top: -100px;
    right: 10%;
    transform-origin: 501px 150px;
    width: 60%;
    height: 300px;
    border-radius: 100%;
    z-index: 3;
    background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3CradialGradient%20id%3D%221%22%20cx%3D%220%22%20cy%3D%220%22%20r%3D%221%22%20gradientTransform%3D%22matrix(-0.6689999999999999%2C%200.225%2C%20-0.01901959553946%2C%20-0.6308669999999998%2C%200.7%2C%200.342)%22%3E%3Cstop%20stop-color%3D%22%23fbffb1%22%20stop-opacity%3D%220.6%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23fbffb8%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FradialGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  }

  .bg-04 {
    position: absolute;
    z-index: 4;
    left: 10%;

    .bg-04-1 {
      position: absolute;
      top: -25px;
      left: 2px;
      z-index: 4;
      transform-origin: 71px 25px;
      width: 142px;
      height: 50px;
      border-radius: 100%;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.12398333663955563%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.3%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.1%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }

    .bg-04-2 {
      position: absolute;
      top: 40px;
      left: 23px;
      z-index: 4;
      transform-origin: 50px 17.5px;
      width: 100px;
      height: 35px;
      border-radius: 100%;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.12249999999999998%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.4%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }

    .bg-04-3 {
      position: absolute;
      top: 88px;
      left: 57px;
      z-index: 4;
      transform-origin: 38.5px 13.5px;
      width: 77px;
      height: 27px;
      border-radius: 100%;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.12295496711081128%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.45%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }

    .bg-04-4 {
      position: absolute;
      top: 122px;
      left: 105px;
      z-index: 4;
      transform-origin: 29px 10.5px;
      width: 58px;
      height: 21px;
      border-radius: 100%;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.13109393579072534%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.5%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }

    .bg-04-5 {
      position: absolute;
      top: 145px;
      left: 144px;
      z-index: 4;
      transform-origin: 22.5px 8.5px;
      width: 45px;
      height: 17px;
      border-radius: 100%;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.14271604938271604%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.55%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }

    .bg-04-6 {
      position: absolute;
      top: 162px;
      left: 176px;
      z-index: 4;
      transform-origin: 19.5px 7px;
      width: 39px;
      height: 14px;
      border-radius: 100%;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.12886259040105194%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.6%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }
  }

  .content {
    position: absolute;
    z-index: 5;
    height: 100%;
    left: 50%;
    transform: translateX(-50%);
  }
}
</style>
