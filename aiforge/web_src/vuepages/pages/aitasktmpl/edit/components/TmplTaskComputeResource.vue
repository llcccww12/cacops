<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title">
        <span class="required">{{ $t('cloudbrainObj.computeResource') }}</span>
      </div>
      <div class="content">
        <div class="item" :class="computerResouce == item.key ? 'focus' : ''" @click="handlerClick(item)"
          v-for="item in computerResouces" :key="item.key" :value="item.key"> {{ item.label }}</div>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>
import { configCreateManager, COMPUTER_RESOURCES_TITLE } from '~/pages/cloudbrain/configs';
import { CLUSTERS } from '~/const';
import { getListValueWithKey } from '~/utils';

export default {
  name: 'TmplTaskComputerResource',
  props: {
    configs: { type: Object, default: () => ({}) },
  },
  data() {
    return {
      clusters: [],
      cluster: '',
      computerResouces: [],
      computerResouce: '',
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
          const _computerResouces = configCreateManager.getComputerResources(newVal.taskType)
          this.computerResouces = _computerResouces.map(item => {
            return {
              key: item,
              label: getListValueWithKey(COMPUTER_RESOURCES_TITLE, item),
            }
          });
          this.computerResouce = newVal.computerResouce
        }
      }
    }
  },
  methods: {
    handlerClick(item) {
      if (item.key == this.computerResouce) return;
      this.change(item.key);
    },
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

.form-row {
  margin-bottom: 20px !important;

  .content {
    display: flex;
    align-items: center;
    flex-wrap: wrap;

    .item {
      height: 40px;
      display: flex;
      align-items: center;
      padding: 0 12px;
      background: rgb(245, 245, 245);
      border-right: 1px solid rgb(214, 215, 217);
      color: rgb(16, 16, 16);
      margin-bottom: 8px;
      cursor: pointer;

      &:first-child {
        border-top-left-radius: 5px;
        border-bottom-left-radius: 5px;
      }

      &:last-child {
        border-right: none;
        border-top-right-radius: 5px;
        border-bottom-right-radius: 5px;
      }

      &.focus {
        color: rgb(0, 102, 255);
        border: 1px solid rgb(50, 145, 248);
      }
    }
  }
}
</style>
