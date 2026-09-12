<template>
  <div>
    <div class="form-row" v-if="configs.showBenchmarkMode === true">
      <div class="left-area">
        <div class="title align-items-center"><span class="required">评测场景</span></div>
        <div class="content">
          <div class="list">
            <a class="item" :href="`/${repoOwnerName}/${repoName}/${item.url}`"
              :class="item.key == configs.bechmarkMode ? 'focus' : ''" v-for="item in benchmarkModeList"
              :key="item.key">
              <span>{{ item.label }}</span>
            </a>
          </div>
        </div>
      </div>
    </div>
    <div class="form-row form-row-cluster" v-if="configs.hideCluster !== true">
      <div class="left-area">
        <div class="title align-items-center"><span class="required">{{ $t('cloudbrainObj.cluster') }}</span></div>
        <div class="content">
          <div class="list">
            <a class="item" :href="`/${repoOwnerName}/${repoName}/${item.url}`"
              :class="item.key == configs.cluster ? 'focus' : ''" v-for="item in configs.clusters" :key="item.key">
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
            <a class="item" :href="`/${repoOwnerName}/${repoName}/${item.url}`"
              :class="item.key == configs.computerResouce ? 'focus' : ''" v-for="item in configs.computerResouces"
              :key="item.key">
              <i class="icon ri-archive-drawer-line"></i>
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
          <div class="tips tips-2" v-if="configs.hideTips2 !== true">
            <i class="ri-error-warning-line"></i>
            <span v-html="configs.tips2"></span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "FormTop",
  props: {
    repoOwnerName: { type: String, default: "" },
    repoName: { type: String, default: "" },
    queueNum: { type: Number, default: 1 },
    configs: { type: Object, default: () => ({}) },
  },
  data() {
    return {
      benchmarkModeList: [{
        key: 'alogrithm',
        label: '算法评测',
        url: 'cloudbrain/benchmark/create?benchmarkMode=alogrithm',
      }, {
        key: 'model',
        label: '模型评测',
        url: 'cloudbrain/benchmark/create?benchmarkMode=model',
      }, {
        key: 'modelsafety',
        label: '模型安全评测',
        url: 'modelsafety/create_gpu',
      }],
    };
  },
  watch: {},
  methods: {
    check() { return true; }
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
    font-size: 12px;
    color: rgba(0, 0, 0, .87);
    border: 1px solid rgba(34, 36, 38, .15);
    margin-left: -1px;
    height: 38px;
    padding: 0 12px;
    margin-bottom: 5px;

    i {
      margin-top: -7px;
      font-size: 14px;
    }

    &.focus {
      color: #0087f5;
      border-color: #0087f5;
      border-left: 1px solid #0087f5;
    }

    &:first-child {
      border-top-left-radius: 0.28571429rem;
      border-bottom-left-radius: 0.28571429rem;
      border-left: 1px solid rgba(34, 36, 38, .15);

      &.focus {
        border-color: #0087f5;
      }
    }

    &:last-child {
      border-top-right-radius: 0.28571429rem;
      border-bottom-right-radius: 0.28571429rem;
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
