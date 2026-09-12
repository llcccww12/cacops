<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center">
        <span class="required">{{ $t('cloudbrainObj.cluster') }}</span>
      </div>
      <div class="content">
        <el-select class="field-input" :value="configs.cluster" @change="change">
          <el-option v-for="item in configs.clusters" :key="item.key" :value="item.key" :label="item.name"></el-option>
        </el-select>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: 'TmplTaskCluster',
  props: {
    configs: { type: Object, default: () => ({}) },
  },
  data() {
    return {};
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
