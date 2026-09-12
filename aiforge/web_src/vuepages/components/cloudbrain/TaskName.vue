<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title">
        <span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.taskName') }}</span>
      </div>
      <div class="content" :class="errStatus ? 'error' : ''">
        <el-input class="field-input" v-model="currentValue" @input="handleInput" @change="handleInputChange"
          :maxlength="type == 'supercompute' ? 26 : 36" :placeholder="$t('cloudbrainObj.taskName')"
          :autofocus="autofocus"></el-input>
        <div v-if="type === 'supercompute'" class="tips">{{ $t('cloudbrainObj.taskNameTips1') }}</div>
        <div v-else class="tips">{{ $t('cloudbrainObj.taskNameTips') }}</div>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>
import dayjs from 'dayjs';

export default {
  name: 'TaskName',
  props: {
    value: { type: String, required: true },
    userName: { type: String, required: true },
    autofocus: { type: Boolean, default: true, },
    generate: { type: Boolean, default: false, },
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
      const reg = /^[a-z0-9][a-z0-9\-_]{1,34}[a-z0-9\-]$/;
      const reg1 = /^[a-z][a-z0-9\-]{4,25}$/;
      if (this.type == 'supercompute') {
        this.errStatus = !reg1.test(this.currentValue);
      } else {
        this.errStatus = !reg.test(this.currentValue);
      }
      if (!this.required && this.currentValue == '') {
        this.errStatus = false;
      }
      return !this.errStatus;
    },
    generateName() {
      let str = this.userName.toLocaleLowerCase();
      const reg1 = /[^a-z0-9_\-]+/g;
      const reg2 = /^[_\-]+/g;
      const reg3 = /[_]+$/g;
      str = str.replace(reg1, '').replace(reg2, '').replace(reg3, '');
      str = str.slice(0, 5);
      const now = Date.now();
      return str + dayjs(now).format('YYYYMMDDHH') + (now / 1000).toFixed(0).slice(-5);
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
    if (!this.value && this.generate) {
      this.currentValue = this.generateName();
      this.$emit('input', this.currentValue);
    }
  }
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';
</style>
