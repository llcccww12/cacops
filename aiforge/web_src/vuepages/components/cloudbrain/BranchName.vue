<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title">
        <span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.codeBranch') }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <el-select class="field-input" v-model="currentValue" @change="handleChange">
          <el-option v-for="item in branches" :key="item" :value="item" :label="item"></el-option>
        </el-select>
        <div class="dynamic-tips" v-if="tips" v-html="tips"></div>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: 'BranchName',
  props: {
    value: { type: String, required: true },
    branches: { type: Array, required: true, },
    required: { type: Boolean, default: true },
    tips: { type: String, default: '' },
  },
  data() {
    return {
      currentValue: '',
      errStatus: false,
    };
  },
  watch: {
    value: {
      immediate: true,
      handler(newVal) {
        newVal = newVal === undefined ? '' : newVal;
        this.currentValue = newVal.toString();
      }
    }
  },
  methods: {
    check() {
      if (this.required && !this.currentValue) {
        this.errStatus = true;
      } else {
        this.errStatus = false;
      }
      return !this.errStatus;
    },
    handleChange(value) {
      this.currentValue = value;
      this.$emit('input', value);
      this.$emit('change', value);
    },
  },
  mounted() { },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';
</style>
