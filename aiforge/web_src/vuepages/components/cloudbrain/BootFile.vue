<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center" v-if="showTitle">
        <span :class="required ? 'required' : ''">{{ $t('modelManage.bootFile') }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <el-input class="field-input" v-model="currentValue" :placeholder="$t('bootPlaceholder')" @input="handleInput"
          @change="handleInputChange"
          @keyup.native="$event.target.value = $event.target.value.replace(/^\s+|\s+$/gm, '')"></el-input>
      </div>
    </div>
    <div class="right-area">
      <div class="btn-select">
        <el-tooltip placement="top" effect="light">
          <i class="question circle icon link" style="margin-top:-7px"></i>
          <div slot="content">
            <div style="width:200px;text-align:center;">{{ $t('modelObj.boot_file_helper') }}</div>
          </div>
        </el-tooltip>
        <a style="color: rgb(0, 102, 255);" :href="sampleUrl" target="_blank">{{ $t('modelManage.viewSamples') }}</a>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "BootFile",
  props: {
    value: { type: String, required: true },
    required: { type: Boolean, default: true },
    showTitle: { type: Boolean, default: true },
    sampleUrl: { type: String, default: '' },
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
      } else {
        const reg = /.+\.py$/;
        this.errStatus = !reg.test(this.currentValue.replace(/^\s+|\s+$/gm, ''));
      }
      return !this.errStatus;
    },
    handleInput(value) {
      this.currentValue = value.replace(/^\s+|\s+$/gm, '');
      this.$emit('input', value.replace(/^\s+|\s+$/gm, ''));
    },
    handleInputChange(value) {
      this.$emit('change', value.replace(/^\s+|\s+$/gm, ''));
      this.check();
    },
  },
  beforeMount() { },
  mounted() { },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';
</style>
