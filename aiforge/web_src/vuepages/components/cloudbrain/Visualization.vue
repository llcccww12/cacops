<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center">
        <span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.visualization') }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <el-checkbox v-model="currentValue" @change="handleChange">{{ $t('cloudbrainObj.viewVisualization')
          }}</el-checkbox>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: 'Visualization',
  props: {
    value: { type: Boolean, required: true },
    required: { type: Boolean, default: false },
  },
  data() {
    return {
      currentValue: false,
      errStatus: false,
    };
  },
  watch: {
    value: {
      immediate: true,
      handler(newVal) {
        this.currentValue = !!newVal;
      }
    }
  },
  methods: {
    check() {
      return !this.errStatus;
    },
    handleChange(value) {
      this.currentValue = value;
      this.$emit('input', value);
      this.check();
    },
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.content {
  display: flex;
  align-items: center;
}

.tooltip {
  margin-left: 25px;
}
</style>
