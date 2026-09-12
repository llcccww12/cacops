<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title align-items-center">
        <span class="required">{{ $t('cloudbrainObj.computeResource') }}</span>
      </div>
      <div class="content">
        <el-select class="field-input" :value="configs.computerResouce" @change="change">
          <el-option v-for="item in configs.computerResouces" :key="item.key" :value="item.key"
            :label="item.label"></el-option>
        </el-select>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: 'TmplTaskComputerResource',
  props: {
    configs: { type: Object, default: () => ({}) },
  },
  data() {
    return {};
  },
  methods: {
    change(computerResouce) {
      this.$emit('change', {
        computerResouce: computerResouce,
      })
    },
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';
</style>
