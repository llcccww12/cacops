<template>
  <div class="jp-btn" :class="[size]" @click="handleClick">
    <div class="bg-1">
      <div class="bg-2">
        <el-tooltip v-if="tooltip" effect="dark" placement="top">
          <div slot="content">{{ tooltip }}</div>
          <div class="content">
            <div class="icon-c">
              <svg xmlns="http://www.w3.org/2000/svg" fill="rgb(255, 255, 255)" viewBox="0 0 512 512" width="16"
                height="16">
                <defs></defs>
                <g>
                  <path
                    d="M478.21 334.093L336 256l142.21-78.093c11.795-6.477 15.961-21.384 9.232-33.037l-19.48-33.741c-6.728-11.653-21.72-15.499-33.227-8.523L296 186.718l3.475-162.204C299.763 11.061 288.937 0 275.48 0h-38.96c-13.456 0-24.283 11.061-23.994 24.514L216 186.718 77.265 102.607c-11.506-6.976-26.499-3.13-33.227 8.523l-19.48 33.741c-6.728 11.653-2.562 26.56 9.233 33.037L176 256 33.79 334.093c-11.795 6.477-15.961 21.384-9.232 33.037l19.48 33.741c6.728 11.653 21.721 15.499 33.227 8.523L216 325.282l-3.475 162.204C212.237 500.939 223.064 512 236.52 512h38.961c13.456 0 24.283-11.061 23.995-24.514L296 325.282l138.735 84.111c11.506 6.976 26.499 3.13 33.227-8.523l19.48-33.741c6.728-11.653 2.563-26.559-9.232-33.036z">
                  </path>
                </g>
              </svg>
            </div>
            <div class="title">{{ title }}</div>
          </div>
        </el-tooltip>
        <div v-else class="content">
          <div class="icon-c">
            <svg xmlns="http://www.w3.org/2000/svg" fill="rgb(255, 255, 255)" viewBox="0 0 512 512" width="16"
              height="16">
              <defs></defs>
              <g>
                <path
                  d="M478.21 334.093L336 256l142.21-78.093c11.795-6.477 15.961-21.384 9.232-33.037l-19.48-33.741c-6.728-11.653-21.72-15.499-33.227-8.523L296 186.718l3.475-162.204C299.763 11.061 288.937 0 275.48 0h-38.96c-13.456 0-24.283 11.061-23.994 24.514L216 186.718 77.265 102.607c-11.506-6.976-26.499-3.13-33.227 8.523l-19.48 33.741c-6.728 11.653-2.562 26.56 9.233 33.037L176 256 33.79 334.093c-11.795 6.477-15.961 21.384-9.232 33.037l19.48 33.741c6.728 11.653 21.721 15.499 33.227 8.523L216 325.282l-3.475 162.204C212.237 500.939 223.064 512 236.52 512h38.961c13.456 0 24.283-11.061 23.995-24.514L296 325.282l138.735 84.111c11.506 6.976 26.499 3.13 33.227-8.523l19.48-33.741c6.728-11.653 2.563-26.559-9.232-33.036z">
                </path>
              </g>
            </svg>
          </div>
          <div class="title">{{ title }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "JumpButton",
  props: {
    size: { type: String, default: 'big' }, // big|small
    title: { type: String, default: '' },
    tooltip: { type: String, default: '' },
  },
  data() {
    return {};
  },
  methods: {
    handleClick() {
      this.$emit('click');
    },
  },
  mounted() { },
};
</script>
<style scoped lang="less">
.jp-btn {
  position: relative;
  height: 40px;
  border-radius: 4px;
  cursor: pointer;

  .bg-1 {
    height: 100%;
    border-radius: 4px;
    background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.596%2C%200.9139999999999999%2C%20-0.07461224489795917%2C%200.596%2C%20-0.016%2C%200.047)%22%3E%3Cstop%20stop-color%3D%22%2350a0f7%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%230b3ff7%22%20stop-opacity%3D%220.88%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");

    .bg-2 {
      height: 100%;
      border-radius: 4px;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3CradialGradient%20id%3D%221%22%20cx%3D%220%22%20cy%3D%220%22%20r%3D%221%22%20gradientTransform%3D%22matrix(-1.295%2C%200.9540000000000001%2C%20-0.16112865306122448%2C%20-2.6793549999999997%2C%200.998%2C%200.061)%22%3E%3Cstop%20stop-color%3D%22%23ffb798%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23650bf7%22%20stop-opacity%3D%220%22%20offset%3D%220.82%22%3E%3C%2Fstop%3E%3C%2FradialGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    }
  }

  .content {
    height: 100%;
    top: 0px;
    left: 0px;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0 10px;
    z-index: 1;

    .icon-c {
      position: relative;
      border: 0.8px solid rgb(255, 255, 255);
      border-radius: 100%;
      height: 22px;
      width: 22px;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 6px;

      svg {
        width: calc(100% - 4px);
        height: calc(100% - 4px);
      }
    }

    .title {
      color: rgb(255, 255, 255);
      font-size: 14px;
    }
  }

  &.small {
    height: 30px;

    .content {
      padding: 0 8px;

      .icon-c {
        height: 20px;
        width: 20px;
      }
    }
  }
}
</style>
