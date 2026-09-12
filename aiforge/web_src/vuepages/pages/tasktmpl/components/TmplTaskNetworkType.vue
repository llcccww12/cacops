<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center">
        <span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.networkType') }}</span>
      </div>
      <div class="content">
        <el-select class="field-input" v-model="currentValue" @change="handleChange">
          <el-option value="no_internet" :label="$t('cloudbrainObj.noInternet')"></el-option>
          <el-option value="has_internet" :label="$t('cloudbrainObj.hasInternet')"></el-option>
        </el-select>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: 'TmplTaskNetworkType',
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
    handleChange(value) {
      this.currentValue = value;
      this.$emit('input', value);
      this.$emit('change', value);
    },
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';

.content {
  display: flex;
  align-items: center;
}
</style>
