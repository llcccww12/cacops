<template>
  <div class="steps-warp">
    <div class="step-item" :class="index + 1 == step ? 'active' : ''" v-for="(item, index) in steps" :key="index">
      <div class="step-link-line" v-if="index != 0"></div>
      <div class="step-num">{{ index + 1 }}</div>
      <div class="step-text">{{ item.name }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Steps',
  props: {
    step: { type: Number, default: 1 },
    steps: { type: Array, required: true, },
  },
  data() {
    return {};
  },
  components: {},
  methods: {},
  mounted() { },
};
</script>

<style scoped lang="less">
.steps-warp {
  display: flex;
  align-items: center;
  margin-bottom: 28px;

  .step-item {
    display: flex;
    align-items: center;

    .step-num {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 46px;
      height: 46px;
      border: 1px solid rgb(230, 231, 239);
      background: rgb(249, 249, 249);
      color: rgba(136, 136, 136, 1);
      font-size: 18px;
      font-weight: 700;
      border-radius: 50%;
      box-sizing: border-box;
    }

    .step-text {
      margin-left: 8px;
      color: rgba(64, 64, 64, 1);
      font-size: 14px;
    }

    &.active {
      .step-num {
        border-color: rgb(50, 145, 248);
        color: rgba(0, 102, 255, 1);
      }

      .step-text {
        color: rgba(50, 145, 248, 1);
      }
    }
  }

  .step-link-line {
    margin: 0 10px;
    border-top: 1px solid rgba(230, 231, 239);
    height: 1px;
    width: 140px;
  }
}
</style>
