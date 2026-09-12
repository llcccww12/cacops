<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title">
        <span :class="required ? 'required' : ''">{{ $t('taskTmplObj.tmplName') }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <el-input class="field-input" v-model="currentValue" @input="handleInput" @change="handleInputChange"
          :maxlength="100" :placeholder="$t('taskTmplObj.tmplName')" :autofocus="autofocus"></el-input>
        <div class="tips">{{ $t('datasetObj.dataset_name_cn_tooltips') }}</div>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: 'TmplName',
  props: {
    value: { type: String, required: true },
    autofocus: { type: Boolean, default: true, },
    required: { type: Boolean, default: true },
    type: { type: String, default: '' },
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
      this.errStatus = false;
      this.currentValue = this.currentValue.trim();
      if (!this.currentValue || !/^[\u4e00-\u9fa5a-zA-Z0-9-_.]{0,100}$/.test(this.currentValue)) {
        this.errStatus = true;
      }
      return !this.errStatus;
    },
    handleInput(value) {
      this.currentValue = value;
      this.$emit('input', value);
      this.check();
    },
    handleInputChange(value) {
      this.$emit('change', value);
    },
  },
  beforeMount() {

  }
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';
</style>
