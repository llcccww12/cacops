<template>
  <div class="demand-c">
    <div class="bg only-mobile-hidden"><img src="/img/cp-resource-bg.jpg" alt="" /></div>
    <div class="ui container">
      <div class="top-head only-mobile-hidden">
        <div class="title">{{ $t('computingPowerObj.computeResource') }}</div>
        <div class="descr" style="color:transparent;user-select:none;">{{ $t('computingPowerObj.computeResourceDescr')
          }}
        </div>
      </div>
      <div class="main-content">
        <div class="top-area">
          <div class="tab-c">
            <div class="tab-item" :class="activeTab == 0 ? 'active' : ''" @click="changeTab(0)">{{
              $t('computingPowerObj.freeComputingPower') }}</div>
            <div class="tab-item" :class="activeTab == 1 ? 'active' : ''" @click="changeTab(1)">
              {{ $t('computingPowerObj.paidComputingPower') }}
              <el-tooltip effect="dark" placement="top">
                <div slot="content">{{ $t('computingPowerObj.paidComputingPowerTip') }} </div>
                <svg xmlns="http://www.w3.org/2000/svg" style="margin-left:5px;" fill="rgb(255, 98, 0)"
                  viewBox="0 0 32 32" width="16" height="16">
                  <defs></defs>
                  <g>
                    <path
                      d="M16 29.333c-7.364 0-13.333-5.969-13.333-13.333s5.969-13.333 13.333-13.333 13.333 5.969 13.333 13.333-5.969 13.333-13.333 13.333zM16 26.667c5.891 0 10.667-4.776 10.667-10.667s-4.776-10.667-10.667-10.667v0c-5.891 0-10.667 4.776-10.667 10.667s4.776 10.667 10.667 10.667v0zM14.667 20h2.667v2.667h-2.667v-2.667zM17.333 17.807v0.86h-2.667v-2c0-0.736 0.597-1.333 1.333-1.333v0c1.105-0 2-0.895 2-2s-0.895-2-2-2c-0.966 0-1.772 0.685-1.959 1.595l-0.002 0.013-2.616-0.524c0.442-2.155 2.323-3.752 4.577-3.752 2.578 0 4.668 2.090 4.668 4.668 0 2.103-1.39 3.881-3.302 4.465l-0.033 0.009z">
                    </path>
                  </g>
                </svg>
              </el-tooltip>
            </div>
          </div>
          <div class="operate-c only-mobile-hidden">
            <JumpButton :title="$t('computingPowerObj.createNewComputingTask')" @click="goAiTask" />
          </div>
        </div>
        <div class="middle-area">
          <div class="left-area" v-loading="loading">
            <div class="tab-content" v-show="activeTab == 0">
              <Resources :active="activeTab == 0">
              </Resources>
            </div>
            <div class="tab-content" v-show="activeTab == 1">
              <ExtraResources :active="activeTab == 1">
              </ExtraResources>
            </div>
          </div>
        </div>
        <div>
          <Partner />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Resources from '../components/Resources.vue';
import ExtraResources from '../components/ExtraResources.vue';
import JumpButton from '../components/JumpButton.vue';
import Partner from '../components/Partner.vue';

export default {
  data() {
    return {
      isLogin: false,
      isOperator: false,
      loading: false,
      activeTab: 0,
    };
  },
  components: { Resources, ExtraResources, JumpButton, Partner },
  methods: {
    changeTab(tab) {
      if (tab == this.activeTab) return;
      this.activeTab = tab;
      this.getData();
    },
    getData() { },
    goAiTask() {
      window.location.href = '/cloudbrains/create';
    },
  },
  created() {
    this.getData();
  },
  mounted() {
    this.isLogin = !!document.querySelector('meta[name="_uid"]');
    this.isOperator = window.IS_OPERATOR;
    const searchParams = new URLSearchParams(window.location.search);
    const active = Number(searchParams.get('active') || 0);
    if (active) {
      this.$nextTick(() => {
        this.changeTab(active);
      })
    }
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.demand-c {
  position: relative;
}

.bg {
  position: absolute;
  top: -40px;
  left: 0;
  width: 100%;
  height: 150px;
  z-index: -1;
  overflow: hidden;

  img {
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -48%);
    width: 135%;
    height: 300%;
  }
}

.new-demand-tab-mobile-show {
  display: none !important;
}

.new-demand-content {
  padding-top: 22px;
}

@media only screen and (max-width: 767px) {
  .new-demand-tab-mobile-show {
    display: flex !important;
    color: #007AFF;
  }

  .main-content {
    margin-top: 0 !important;
  }

  .top-area {
    justify-content: center !important;
    margin: 0 -1rem;
    background: #F7F7F7 100%;
    border-bottom: 1px solid rgba(229, 229, 229, 1);
    padding-top: 12px;
  }

  .only-mobile-hidden {
    display: none !important;
  }

  .tab-item {
    border: none !important;
    padding: 0 12px !important;
  }

  /deep/ .el-pagination__total,
  /deep/ .el-pagination__sizes,
  /deep/ .el-pagination__jump {
    display: none !important;
  }

  .demand-item-left {
    flex: unset !important;
    width: 100% !important;
  }

  .demand-item-line {
    display: grid !important;
    align-content: center;
    justify-content: start;
    align-items: stretch;
    justify-items: stretch;
  }
}

.top-head {
  margin-top: 40px;

  .title {
    text-align: center;
    color: rgba(16, 16, 16, 1);
    font-size: 28px;
    margin-bottom: 20px;
  }

  .descr {
    color: rgba(136, 136, 136, 0.87);
    font-size: 14px;
    text-align: center;
  }
}

.main-content {
  margin-top: 75px;

  .top-area {
    display: flex;
    align-items: center;
    justify-content: space-between;

    .tab-c {
      display: flex;
      align-items: center;

      .tab-item {
        border: 1px solid rgb(229, 229, 229);
        height: 40px;
        display: flex;
        align-items: center;
        text-align: center;
        padding: 0 16px;
        cursor: pointer;

        &:first-child {
          border-top-left-radius: 4px;
          border-bottom-left-radius: 4px;
        }

        &:last-child {
          border-top-right-radius: 4px;
          border-bottom-right-radius: 4px;
        }

        &.active {
          border-color: rgb(1, 72, 255);
          color: rgb(0, 60, 213);
        }
      }
    }

    .operate-c {
      display: flex;
      align-items: center;
    }
  }

  .middle-area {
    display: flex;

    .left-area {
      flex: 1;
      width: 0;

      .tab-content {
        min-height: 390px;
      }

      .pagination-c {
        text-align: center;
        margin: 10px 0;
      }
    }

    .right-area {
      width: 460px;
      margin-left: 22px;
      padding-top: 22px;

      .form-container {
        top: 10px;
        position: sticky;
        border: 1px solid rgb(229, 231, 235);
        background: rgb(249, 250, 251);
        border-radius: 10px;
        color: rgb(16, 16, 16);
        padding: 20px;
        padding-right: 26px;
      }
    }
  }
}

.demand-item {
  padding: 22px 26px 16px 26px;
  border-radius: 4px;
  margin: 22px 0;
  border-color: rgba(157, 197, 226, 0.4);
  border-width: 1px;
  border-style: solid;
  box-shadow: rgba(157, 197, 226, 0.2) 0px 5px 10px 0px;
  color: rgb(16, 16, 16);
  border-radius: 15px;
  background: rgb(255, 255, 255);
  font-size: 14px;

  .demand-item-bottom {
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-top: 1px solid rgba(157, 197, 226, 0.2);
    padding-top: 12px;
    height: 30px;

    .demand-item-left {
      flex: 1;
      width: 0;
      color: rgba(16, 16, 16, 0.6);
    }

    .demand-item-right {
      width: 320px;
      text-align: right;
      display: flex;
      justify-content: flex-end;
      align-items: center;

      .edit-btn {
        cursor: pointer;
        margin-right: 10px;
        border: 1px solid rgb(50, 145, 248);
        color: rgb(50, 145, 248);
        border-radius: 4px;
        height: 28px;
        line-height: 26px;
        text-align: center;
        padding: 0 8px;
        background: rgba(50, 145, 248, 0.1);
      }

      .status {
        margin-left: 8px;

        i {
          margin-right: 3px;
        }

        &.accepted {
          color: rgb(39, 177, 72);
        }

        &.refuse {
          color: rgb(140, 162, 170);
        }

        &.pending {
          color: rgb(50, 145, 248);
        }
      }
    }
  }

  .demand-item-line {
    display: flex;
    margin-bottom: 16px;

    .demand-item-line-block {
      flex: 1;
      margin-right: 20px;

      span:first-child {
        color: rgb(136, 136, 136);
        margin-right: 2px;
      }
    }
  }
}

.no-data {
  display: flex;
  justify-content: center;
  padding: 0 0;
  width: 100%;

  .item-empty {
    height: 391px;
    width: 100%;
    overflow: hidden;
    padding: 15px;
    background: transparent;
    display: flex;
    flex-direction: column;
    justify-content: center;
    background-color: rgba(245, 245, 246, 0.5);

    .item-empty-icon {
      height: 80px;
      width: 100%;
      background: url(/img/empty-box.svg) center center no-repeat;
    }

    .item-empty-tips {
      text-align: center;
      margin-top: 20px;
      font-size: 18px;
      color: rgb(63, 63, 64);
    }
  }
}
</style>
