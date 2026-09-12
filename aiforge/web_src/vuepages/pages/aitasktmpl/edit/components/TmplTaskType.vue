<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center">
        <span class="required">{{ $t('resourcesManagement.jobType') }}</span>
      </div>
      <div class="content">
        <el-select class="field-input" v-model="task" @change="changeType">
          <el-option v-for="(item, index) in tasks" :key="index" :value="item.k" :label="item.v"></el-option>
        </el-select>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>
import { TmplTaskTypes } from '~/pages/aitasktmpl/tools';

export default {
  name: 'TmplTaskType',
  props: {
    value: { type: String, required: true },
  },
  data() {
    return {
      tasks: [...TmplTaskTypes],
      task: 'DEBUG',
    };
  },
  watch: {
    value: {
      immediate: true,
      handler(newVal) {
        newVal = newVal === undefined ? '' : newVal;
        this.task = newVal.toString();
        if (this.task) {
          this.emitChange();
        }
      }
    }
  },
  methods: {
    changeType(type) {
      if (this.disabled) return;
      this.task = type;
      this.emitChange();
    },
    emitChange() {
      const findTask = this.tasks.filter(item => item.k == this.task)[0];
      this.$emit('input', this.task, findTask);
      this.$emit('change', this.task, findTask);
    }
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';
</style>
