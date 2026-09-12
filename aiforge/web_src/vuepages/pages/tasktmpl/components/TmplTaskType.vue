<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center">
        <span class="required">{{ $t('resourcesManagement.jobType') }}</span>
      </div>
      <div class="content">
        <el-select class="field-input" v-model="task" @change="changeType">
          <el-option v-for="(item, index) in tasks" :key="index" :value="item.k" :label="item.name"></el-option>
        </el-select>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>
import { i18n } from '~/langs';

export default {
  name: 'TmplTaskType',
  props: {
    value: { type: String, required: true },
  },
  data() {
    return {
      tasks: [{
        k: 'DEBUG',
        name: i18n.t('cloudbrainObj.tabTitDebug'),
        defaultSelect: {
          cluster: 'C2Net',
          computerResouce: 'NPU',
        },
      }, {
        k: 'TRAIN',
        name: i18n.t('cloudbrainObj.tabTitTrain'),
        defaultSelect: {
          cluster: 'C2Net',
          computerResouce: 'NPU',
        },
      }, {
        k: 'INFERENCE',
        name: i18n.t('cloudbrainObj.tabTitInference'),
        defaultSelect: {
          cluster: 'OpenI',
          computerResouce: 'NPU',
        },
      }, {
        k: 'ONLINEINFERENCE',
        name: i18n.t('cloudbrainObj.tabTitOnlineInference'),
        defaultSelect: {
          cluster: 'C2Net',
          computerResouce: 'GPU',
        },
      }, {
        k: 'GENERAL',
        name: i18n.t('cloudbrainObj.tabTitGeneral'),
        defaultSelect: {
          cluster: 'C2Net',
          computerResouce: 'GPU',
        },
      }],
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
