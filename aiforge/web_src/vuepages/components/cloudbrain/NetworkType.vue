<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center">
        <span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.networkType') }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <el-radio-group v-model="currentValue" @input="handleInput">
          <el-radio label="no_internet">{{ $t('cloudbrainObj.noInternet') }}</el-radio>
          <el-radio label="has_internet">{{ $t('cloudbrainObj.hasInternet') }}</el-radio>
        </el-radio-group>
        <el-tooltip class="tooltip" placement="top" effect="light">
          <i class="question circle icon link" style="margin-top:-7px"></i>
          <div slot="content">
            <div style="width:200px;text-align:center;" v-html="$t('cloudbrainObj.networkTypeDesc')"></div>
          </div>
        </el-tooltip>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: 'NetworkType',
  props: {
    value: { type: String, required: true },
    required: { type: Boolean, default: true },
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
      return !this.errStatus;
    },
    handleInput(value) {
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
