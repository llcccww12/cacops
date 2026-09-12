<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center">
        <span class="required">{{ $t('cloudbrainObj.cluster') }}</span>
      </div>
      <div class="content">
        <el-select class="field-input" :value="cluster" @change="change">
          <el-option v-for="item in clusters" :key="item.key" :value="item.key" :label="item.label"></el-option>
        </el-select>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>
import { configCreateManager } from '~/pages/cloudbrain/configs';
import { CLUSTERS } from '~/const';
import { getListValueWithKey } from '~/utils';

export default {
  name: 'TmplTaskCluster',
  props: {
    configs: { type: Object, default: () => ({}) },
  },
  data() {
    return {
      cluster: '',
      clusters: [],
    };
  },
  watch: {
    configs: {
      deep: true,
      handler(newVal) {
        if (JSON.stringify(newVal) !== '{}') {
          const _clusters = configCreateManager.getTaskTypeAllClusters(newVal.taskType)
          this.clusters = _clusters.map(item => {
            return {
              key: item,
              label: getListValueWithKey(CLUSTERS, item),
            }
          });
          this.cluster = newVal.cluster
        }
      }
    }
  },
  methods: {
    change(cluster) {
      this.$emit('change', {
        cluster: cluster,
      })
    },
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';
</style>
