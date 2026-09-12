<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title"><span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.resourceSpec') }}</span></div>
      <div class="content" :class="errStatus ? 'error' : ''" v-loading="loading">
        <div class="spec-list-c" v-if="!list.length">
          <div class="spec-item-placeholder" style="color: red;">
            {{ $t('specObj.no_use_resource') }}
          </div>
        </div>
        <div v-else class="spec-info">
          <el-select class="spec-sel field-input" :class="configs.showPoint ? 'spec-show-point' : ''" v-model="spec"
            :placeholder="$t('cloudbrainObj.specPlaceholder')" @change="changeSpec">
            <div slot="prefix" class="spec-sel-icon spec-op-icon">
              <div :class="selIconType + '_icon _icon'">{{ selIconType[0] }}</div>
            </div>
            <el-option v-for="(item) in list" :key="item.id" :label="item.specStr" :value="item.id">
              <span style="float: left" class="spec-op-icon">
                <div :class="item.type + '_icon _icon'">{{ item.type[0] }}</div>
              </span>
              <span class="spec-op-spec" style="float: left">{{ item.specStr }}</span>
              <span class="spec-op-point" style="float: right;" v-if="configs.showPoint">{{ item.pointStr }}</span>
            </el-option>
            <div slot="prefix" v-if="configs.showPoint" class="spec-sel-point"> {{ selPointStr }} </div>
          </el-select>
          <div class="self-point-info" v-if="configs.showPoint">
            <span>{{ $t('cloudbrainObj.balanceOfPoints') }}：<span style="color:red"> {{ configs.blance.toFixed(2) }}</span>
              {{ $t('cloudbrainObj.points') }}<span v-if="showUseTime">{{ $t('cloudbrainObj.canUseTime') }} <span
                  style="color:red">{{ canUseTime
                  }}</span> {{ $t('cloudbrainObj.hours') }}</span></span>
            <span>
              <i class="el-icon-question"></i>
              <a target="_blank" href="/reward/point/rule">{{ $t('cloudbrainObj.PointGainDescr') }}</a>
            </span>
          </div>
          <div class="dynamic-tips" v-if="tips" v-html="tips"></div>
        </div>
      </div>
    </div>
    <div class="right-area resource-descr-c">
      <div class="resource-descr" v-if="!configs.hideHelpLink">
        <el-tooltip placement="top" effect="light">
          <i class="question circle icon link" style="margin-top:-7px"></i>
          <div slot="content">
            <div style="width:200px;text-align:center;">{{ $t('specObj.resSelectTips') }}</div>
          </div>
        </el-tooltip>
        <a target="_blank" href="https://openi.pcl.ac.cn/docs/index.html#/quickstart/resources">{{
          $t('cloudbrainObj.specDescr') }}</a>
      </div>
    </div>
  </div>
</template>
<script>
import { renderSpecObject } from '~/utils';

export default {
  name: "SpecSelect",
  props: {
    value: { type: String, required: true, },
    configs: { type: Object, required: true, },
    required: { type: Boolean, default: true },
    networkType: { type: String, default: 'no_internet' },
    visualize: { type: Boolean, default: false },
    workServerNum: { type: Number, default: 1 },
    loading: { type: Boolean, default: false },
    tips: { type: String, default: '' },
  },
  data() {
    return {
      list: [],
      spec: '',
      selIconType: '',
      selPointStr: '',
      showUseTime: false,
      canUseTime: '',

      errStatus: false,
    };
  },
  watch: {
    value: {
      immediate: true,
      handler(newVal) {
        this.spec = newVal.toString();
        this.renderSpec();
      }
    },
    configs: { // {specs: [],  blance: 0, showPoint: true },
      immediate: true,
      deep: true,
      handler(newVal) {
        this.renderSpec();
      }
    },
    networkType: {
      immediate: true,
      handler() {
        const resetSpec = true;
        this.renderSpec(resetSpec);
      }
    },
    visualize: {
      immediate: true,
      handler() {
        const resetSpec = true;
        this.renderSpec(resetSpec);
      }
    },
    workServerNum: {
      immediate: true,
      handler() {
        this.renderSpec();
      }
    },
  },
  methods: {
    renderSpec(resetSpec) {
      const showPoint = this.configs.showPoint || false;
      let specs = this.configs.specs[this.networkType] || [];
      if (this.visualize) {
        specs = specs.filter(item => item.VisualizeCapableQueuesExist);
      }
      this.list = specs.map((item) => {
        return renderSpecObject(item, showPoint)
      });
      if (resetSpec) {
        this.spec = specs.length ? specs[0].id.toString() : '';
      }
      this.changeSpec(this.spec);
    },
    changeSpec() {
      const seldSpecItem = this.list.filter(item => item.id == this.spec)[0];
      if (seldSpecItem) {
        this.selIconType = seldSpecItem.type;
        this.selPointStr = seldSpecItem.pointStr;
        this.errStatus = false;
      } else {
        this.selIconType = '';
        this.selPointStr = '';
      }
      this.refreshPointInfo();
      this.$emit('input', this.spec);
      this.$emit('change', this.spec);
    },
    refreshPointInfo() {
      const showPoint = this.configs.showPoint || false;
      if (!showPoint) return;
      const seldSpecItem = this.list.filter(item => item.id == this.spec)[0];
      if (seldSpecItem) {
        const unitPrice = seldSpecItem.unit_price;
        const blance = this.configs.blance || 0;
        const workServerNum = this.workServerNum;
        if (unitPrice == 0) {
          this.showUseTime = false;
        } else {
          let canUseTime = Number(blance) / (Number(unitPrice) * Number(workServerNum));
          if (Number(blance) < Number(unitPrice) * Number(workServerNum)) {
            // 余额不足一个单位单价时可用时间提示为 0
            canUseTime = 0;
          }
          this.canUseTime = canUseTime.toFixed(2);
          this.showUseTime = true;
        }
      }
    },
    check() {
      const seldSpecItem = this.list.filter(item => item.id == this.spec)[0];
      if (!seldSpecItem) {
        this.errStatus = true;
        return false;
      }
      this.errStatus = false;
      return true;
    },
    getSpecData() {
      return this.list.filter(item => item.id == this.spec)[0];
    },
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.form-row {
  .content {
    .spec-list-c {
      min-height: 37.6px;
      border-radius: 4px;
      // border: 1px solid #DCDFE6;
      background: rgb(245, 245, 245);
      box-sizing: border-box;
      color: #606266;
      padding: 4px 15px;
      .spec-item-placeholder {
        height: 27px;
        line-height: 27px;
        color: rgba(0, 0, 0, 0.4);
        opacity: 0.45 !important;
        font-size: 14px;
      }
    }

    .spec-info {
      flex: 1;

      .spec-sel {
        width: 100%;

        /deep/.el-input__prefix {
          width: 100%;
          cursor: pointer;
        }

        /deep/.el-input__inner {
          padding-left: 36px;
        }

        .spec-sel-icon {
          position: absolute;
          left: 5px;
          height: 96%;
        }

        .spec-sel-point {
          position: absolute;
          right: 40px;
          display: flex;
          align-items: center;
          height: 99%;
        }

        &.spec-show-point {
          /deep/.el-input__inner {
            padding-right: 120px;
          }
        }
      }

      .self-point-info {
        padding: 0 10px;
        margin-top: 5px;
        display: flex;
        justify-content: space-between;
        font-size: 12px;

        i {
          margin-right: 4px;
          cursor: pointer;
          color: rgba(0, 0, 0, 0.87);
        }
      }
    }
  }

  .resource-descr-c {
    align-items: flex-start;
    margin-top: 8px;

    .resource-descr {
      display: flex;
      align-items: center;

      i {
        margin-right: 6px;
        cursor: pointer;
        color: rgba(0, 0, 0, 0.87);
      }
      a{
        color: rgb(0, 102, 255);
      }
    }
  }
}

.spec-op-icon {
  margin-right: 6px;
  display: flex;
  align-items: center;
  height: 100%;

  ._icon {
    color: white;
    width: 22px;
    height: 22px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 100%;
    background: rgba(136, 136, 136, 0.5);
  }

  .CPU_icon {
    color: white;
    width: 22px;
    height: 22px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 100%;
    background: rgb(54, 207, 201);
  }

  .GPU_icon {
    color: white;
    width: 22px;
    height: 22px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 100%;
    background: rgb(252, 202, 0);
  }

  .NPU_icon,
  .GCU_icon,
  .MLU_icon,
  .DCU_icon,
  .ILUVATAR-GPGPU_icon,
  .METAX-GPGPU_icon {
    color: white;
    width: 22px;
    height: 22px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 100%;
    background: rgb(123, 50, 178)
  }
}

.spec-op-point {
  color: rgba(136, 136, 136, 1);
  font-size: 14px
}
</style>
