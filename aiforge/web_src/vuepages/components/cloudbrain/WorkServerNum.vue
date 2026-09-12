<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center" v-if="showTitle"><span class="required">{{
        $t('cloudbrainObj.computeNodeCount') }}</span></div>
      <div class="content">
        <el-select class="work-server-num-sel field-input" v-model="workServerNum" @change="handleChange">
          <el-option v-for="item in data" :key="item" :value="item" :label="item"></el-option>
        </el-select>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: "WorkServerNum",
  props: {
    value: { type: Number, required: true, },
    required: { type: Boolean, default: true },
    data: { type: Array, required: true, },
    title: { type: String, default: '' },
    showTitle: { type: Boolean, default: true },
  },
  data() {
    return {
      workServerNum: 1,
      errStatus: false,
    };
  },
  watch: {
    value: {
      immediate: true,
      handler(newVal) {
        newVal = newVal === undefined ? 1 : newVal;
        this.workServerNum = Number(newVal)
      }
    }
  },
  methods: {
    handleChange(item) {
      this.$emit('input', this.workServerNum);
      this.$emit('change', this.workServerNum);
    },
    check() {
      return true;
    },
  },
  beforeMount() { },
  mounted() { },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.form-row {
  .left-area {
    .content {
      .work-server-num-sel {
        width: 90px;
      }
    }
  }
}
</style>
