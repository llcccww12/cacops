<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title">
        <span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.taskDescr') }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <el-input class="field-input" type="textarea" v-model="currentValue" :rows="3"
          :placeholder="$t('cloudbrainObj.taskDescrPlaceholder')" :maxlength="255" @blur="handleBlur" @focus="handleFocus"
          @input="handleInput" @change="handleInputChange"></el-input>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: 'TaskDescr',
  props: {
    value: { type: String, required: true },
    required: { type: Boolean, default: false },
    userName: { type: String, default: '' },
    oriData: { type: String, default: '' },
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
      if (this.required && this.currentValue.trim() == '') {
        this.errStatus = true;
      } else {
        this.errStatus = false;
      }
      return !this.errStatus;
    },
    handleBlur(event) {
      this.$emit('blur', event);
    },
    handleFocus(event) {
      this.$emit('focus', event);
    },
    handleInput(value) {
      // this.currentValue = this.truncateByBytes(value, 255);
      this.currentValue = value;
      
      this.$emit('input', this.currentValue);
      this.check();
    },
    handleInputChange(value) {
      // this.currentValue = this.truncateByBytes(value, 255);
      this.$emit('change', value);
    },
    truncateByBytes(str, maxBytes) {
      const encoder = new TextEncoder();
      const decoder = new TextDecoder('utf-8', { fatal: false });
      
      const encoded = encoder.encode(str);
      console.log(encoded.length,maxBytes);
      if (encoded.length <= maxBytes) return str;
      
      const truncated = encoded.slice(0, maxBytes);
      console.log(decoder.decode(truncated));
      return decoder.decode(truncated).replace(/\uFFFD/g, '');
    }
  },
  beforeMount() { },
  mounted() { },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.form-row .left-area .content .field-input /deep/ .el-textarea__inner {
  line-height: 1.2857;
  padding: 0.78571429em 1em;
}
</style>
