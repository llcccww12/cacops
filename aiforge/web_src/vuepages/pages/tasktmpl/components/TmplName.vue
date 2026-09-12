<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center">
        <span :class="required ? 'required' : ''">{{ $t('taskTmplObj.tmplName') }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <el-input class="field-input" v-model="currentValue" @input="handleInput" @change="handleInputChange"
          :maxlength="20" :placeholder="$t('taskTmplObj.tmplNamePlacehoulder')" :autofocus="autofocus"></el-input>
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
      if (!this.required && this.currentValue == '') {
        this.errStatus = false;
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
