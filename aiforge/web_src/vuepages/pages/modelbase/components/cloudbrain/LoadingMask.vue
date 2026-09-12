<template>
  <div v-if="loading" class="loading-mask">
    <div class="loading-wrap">
      <div class="loading-icon">
        <div class="rect1"></div>
        <div class="rect2"></div>
        <div class="rect3"></div>
        <div class="rect4"></div>
        <div class="rect5"></div>
      </div>
      <div class="loading-tips"> {{ tipsContent }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "LoadingMask",
  props: {
    loading: { type: Boolean, default: false },
    tips: { type: String, default: '' },
  },
  data() {
    return {};
  },
  computed: {
    tipsContent() {
      return this.tips || '加载中，请稍后'
    }
  },
};
</script>

<style scoped lang="less">
.loading-mask {
  position: fixed;
  top: 0px;
  left: 0px;
  right: 0px;
  bottom: 0px;
  filter: alpha(opacity=60);
  background-color: rgba(119, 119, 119, 0.8);
  z-index: 1000;
  color: #000000;

  .loading-wrap {
    width: 20%;
    height: 25%;
    border: 1px solid #bbb;
    background-color: #fff;
    margin: auto;
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    border-radius: 20px;

    .loading-icon {
      width: 50px;
      height: 40px;
      text-align: center;
      font-size: 10px;
      display: block;

      >div {
        background-color: green;
        height: 100%;
        width: 6px;
        display: inline-block;
        -webkit-animation: sk-stretchdelay 1.2s infinite ease-in-out;
        animation: sk-stretchdelay 1.2s infinite ease-in-out;
      }

      .rect2 {
        -webkit-animation-delay: -1.1s;
        animation-delay: -1.1s;
      }

      .rect3 {
        -webkit-animation-delay: -1s;
        animation-delay: -1s;
      }

      .rect4 {
        -webkit-animation-delay: -0.9s;
        animation-delay: -0.9s;
      }

      .rect5 {
        -webkit-animation-delay: -0.8s;
        animation-delay: -0.8s;
      }
    }

    .loading-tips {
      margin-top: 10%;
      font-size: 14px;
      color: #101010;
      font-family: SourceHanSansSC-regular;
    }
  }
}

@-webkit-keyframes sk-stretchdelay {

  0%,
  40%,
  100% {
    -webkit-transform: scaleY(0.4);
  }

  20% {
    -webkit-transform: scaleY(1);
  }
}

@keyframes sk-stretchdelay {

  0%,
  40%,
  100% {
    transform: scaleY(0.4);
    -webkit-transform: scaleY(0.4);
  }

  20% {
    transform: scaleY(1);
    -webkit-transform: scaleY(1);
  }
}
</style>
