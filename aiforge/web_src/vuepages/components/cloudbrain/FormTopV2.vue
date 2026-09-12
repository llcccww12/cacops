<template>
  <div>
    <div class="form-row form-row-cluster" v-if="configs.hideCluster !== true">
      <div class="left-area">
        <div class="title align-items-center"><span class="required">{{ $t('cloudbrainObj.cluster') }}</span></div>
        <div class="content">
          <div class="list">
            <a class="item" @click="changeCluster(item.key)" :class="item.key == cluster ? 'focus' : ''"
              v-for="item in clusters" :key="item.key">
              <i class="icon ri-global-line"></i>
              <span>{{ item.label }}</span>
            </a>
          </div>
        </div>
      </div>
    </div>
    <div class="form-row form-row-computer-resource" v-if="configs.hideComputerResource !== true">
      <div class="left-area">
        <div class="title"><span class="required">{{ $t('cloudbrainObj.computeResource') }}</span></div>
        <div class="content">
          <div class="list">
            <a class="item" @click="changeComputerResouce(item.key)"
              :class="item.key == computerResouce ? 'focus' : ''" v-for="item in computerResouces"
              :key="item.key">
              <!-- <i class="icon ri-archive-drawer-line"></i> -->
              <span>{{ item.label }}</span>
            </a>
          </div>
        </div>
      </div>
    </div>
    <div class="form-row tips-c">
      <div class="left-area">
        <div class="title"></div>
        <div class="content">
          <div class="tips tips-1">
            <span class="wait-count-c">
              <i class="ri-error-warning-line"></i>
              <span>
                {{ $t('cloudbrainObj.waitCountStart') }}
                <span>{{ queueNum }}</span>
                {{ $t('cloudbrainObj.waitCountEnd') }}
              </span>
            </span>
            <a class="mind-torch-helper" v-if="configs.showMindTorchHelper"
              href="https://openi.pcl.ac.cn/OpenI/mindtorch_tutorial" target="_blank">
              <i class="ri-arrow-right-line"></i>
              <span>{{ $t('cloudbrainObj.mindTorchHelper') }}</span>
            </a>
          </div>
          <!-- <div class="tips tips-2" v-if="configs.hideTips2 !== true">
            <i class="ri-error-warning-line"></i>
            <span v-html="configs.tips2"></span>
          </div> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { configCreateManager,COMPUTER_RESOURCES_TITLE } from '~/pages/cloudbrain/configs';
import { JOB_TYPE, CLUSTERS } from '~/const';
import { getListValueWithKey } from '~/utils';
export default {
  name: "FormTopV2",
  props: {
    queueNum: { type: Number, default: 1 },
    taskTypeObj: { type: Object, default: () => ({}) }, //静态资源-基础任务
    resourceObj: { type: Object, default: () => ({}) }, //动态资源配置项-大模型基地任务
  },
  data() {
    return {
      clusters: [],
      cluster: '',
      computerResouces: [],
      computerResouce: '',
      configs: {}
    };
  },
  watch: {
    taskTypeObj: {
      deep: true,
      handler(newVal) {
        if(JSON.stringify(newVal) !== '{}'){
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
          const _configParams = configCreateManager.getResourceConfig(newVal.taskType,newVal.computerResouce)
          this.configs = _configParams || {}
        }
      }
    },
    resourceObj:{
      deep: true,
      immediate: true,
      handler(newVal){
        if(JSON.stringify(newVal) !== '{}'){
          this.computerResouces = newVal?.computerResouces || []
          this.computerResouce = newVal.computerResouce || ''
          this.clusters = newVal?.clusters || []
          this.cluster = newVal?.cluster || ''
          this.configs = newVal?.configs || {}
        }
      }
    }
  },
  methods: {
    check() { return true; },
    changeCluster(cluster) {
      this.cluster = cluster
      return
      // this.$emit('change', {
      //   cluster: cluster,
      //   computerResouce: '',
      // })
    },
    changeComputerResouce(computerResouce) {
      this.computerResouce = computerResouce
      this.$emit('change', {
        cluster: this.cluster,
        computerResouce: this.computerResouce,
      })
    },
  },
  mounted() { 
    // console.log('taskType', this.taskType)
    // const clusters = configCreateManager.getTaskTypeAllClusters(this.taskType)
  },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.list {
  display: flex;
  align-items: center;
  flex-wrap: wrap;

  .item {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    border-right: 1px solid rgb(194, 199, 204);
    height: 40px;
    padding: 0 12px;
    margin-bottom: 5px;
    background-color: rgba(245,245,245,1);
    span{
      color: rgb(0, 102, 255);
    }
    i {
      margin-top: -7px;
      font-size: 14px;
    }
    &:first-child {
      border-top-left-radius: 0.28571429rem;
      border-bottom-left-radius: 0.28571429rem;
      // border-left: 1px solid rgba(34, 36, 38, .15);

      // &.focus {
      //   border-color: #0087f5;
      // }
    }

    &:last-child {
      border-top-right-radius: 0.28571429rem;
      border-bottom-right-radius: 0.28571429rem;
      border-right: 0;
    }
    &.focus {
      background-color: rgba(255,255,255,1);
      border: 1px solid rgb(50, 145, 248);
      border-radius: 5px;
      margin-left: -4px;
    }

    

    &:hover:not(.focus) {
      background: rgba(0, 0, 0, .03);
    }
  }
}

.tips-c {
  margin-top: -20px;

  .tips {
    display: flex;
    font-size: 12px;
    align-items: center;

    i {
      color: #f2711c;
      margin-right: 5px;
      font-size: 14px;
    }
  }

  .tips-1 {
    flex-wrap: wrap;

    span {
      color: #f2711c;
    }
  }

  .tips-2 {
    span {
      color: #888;
    }
  }

  .wait-count-c {
    display: flex;
    align-items: center;
    margin-right: 20px;
  }

  .mind-torch-helper {
    display: flex;
    align-items: center;

    i {
      color: rgba(0, 122, 255, 1);
      font-size: 15px;
      margin-right: 2px;
    }

    span {
      color: rgba(0, 122, 255, 1);
    }
  }
}
</style>
