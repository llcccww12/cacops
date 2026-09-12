<template>
  <div class="resource-specification">
    <div class="title"><span :class="required ? 'required' : ''">{{ title }}</span></div>
    <div class="content" :class="errStatus ? 'error' : ''">
      <div class="spec-info">
        <el-select class="spec-sel" :class="showPoint ? 'spec-show-point' : ''" v-model="spec" placeholder="请选择资源规格"
          @change="changeSpec">
          <div slot="prefix" class="spec-sel-icon spec-op-icon">
            <div :class="selIconType + '_icon'">{{ selIconType[0] }}</div>
          </div>
          <el-option v-for="(item, index) in list" :key="item.id" :label="item.specStr" :value="item.id">
            <span style="float: left" class="spec-op-icon">
              <div :class="item.type + '_icon'">{{ item.type[0] }}</div>
            </span>
            <span class="spec-op-spec" style="float: left">{{ item.specStr }}</span>
            <span class="spec-op-point" style="float: right;" v-if="showPoint">{{ item.pointStr }}</span>
          </el-option>
          <div slot="prefix" v-if="showPoint" class="spec-sel-point"> {{ selPointStr }} </div>
        </el-select>
        <div class="self-point-info" v-if="showPoint">
          <span>积分余额：<span style="color:red"> {{ blance }}</span> 积分<span v-if="showUseTime">，预计可用 <span
                style="color:red">{{ canUseTime }}</span> 小时</span></span>
          <span>
            <i class="el-icon-question"></i>
            <a target="_blank" href="/reward/point/rule">积分获取说明</a>
          </span>
        </div>
      </div>
      <div class="resource-descr-c">
        <div class="resource-descr">
          <i class="el-icon-question"></i>
          <a target="_blank" href="https://openi.pcl.ac.cn/docs/index.html#/quickstart/resources">资源说明</a>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { renderSpecObject } from '~/utils';

export default {
  name: "ResourceSpecification",
  props: {
    title: { type: String, default: "资源规格" },
    required: { type: Boolean, default: true },
    specData: { type: Array, default: () => [] },
    specOri: { type: String, default: '' },
    showPoint: { type: Boolean, default: false, },
    blance: { type: Number, default: 0 },
    workServerNum: { type: Number, default: 1 },
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
    specData: function (val) {
      this.renderSpec();
    },
    showPoint: function (val) {
      this.renderSpec();
    },
    blance: function (val) {
      this.renderSpec();
    },
    specOri: function (val) {
      this.spec = val;
      this.renderSpec();
    }
  },
  methods: {
    renderSpec(specs) {
      this.list = this.specData.map((item) => {
        return renderSpecObject(item, this.showPoint)
      });
      this.changeSpec(this.spec);
    },
    changeSpec(specID) {
      const seldSpecItem = this.list.filter(item => item.id == this.spec)[0];
      if (seldSpecItem) {
        this.selIconType = seldSpecItem.type;
        this.selPointStr = seldSpecItem.pointStr;
        this.errStatus = false;
      }
      this.refreshPointInfo();
    },
    refreshPointInfo() {
      if (!this.showPoint) return;
      const seldSpecItem = this.list.filter(item => item.id == this.spec)[0];
      if (seldSpecItem) {
        const unitPrice = seldSpecItem.unit_price;
        const blance = this.blance;
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
.resource-specification {
  display: flex;
  margin-bottom: 28px;

  .title {
    width: 200px;
    text-align: right;
    margin-right: 24px;
    color: #101010;
    font-size: 14px;
    display: flex;
    justify-content: flex-end;
    padding-top: 6px;

    .required {
      position: relative;

      &::after {
        position: absolute;
        content: "*";
        top: -3px;
        right: -10px;
        color: red;
      }
    }
  }

  .content {
    flex: 1;
    display: flex;

    .spec-info {
      flex: 1;

      .spec-sel {
        width: 100%;

        /deep/.el-input__prefix {
          width: 100%;
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

    .resource-descr-c {
      padding-top: 6px;

      .resource-descr {
        display: flex;
        margin-left: 10px;
        align-items: center;

        i {
          margin-right: 5px;
          cursor: pointer;
          color: rgba(0, 0, 0, 0.87);
        }
      }
    }

    &.error {
      .spec-sel {
        /deep/.el-input__inner {
          color: #9f3a38;
          background: #fff6f6;
          border-color: #e0b4b4;

          &:visited {
            border-color: #e0b4b4;
          }

          &:focus {
            border-color: #e0b4b4;
          }

          &:active {
            border-color: #e0b4b4;
          }
        }
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

  .NPU_icon {
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

// .spec-op-spec {}

.spec-op-point {
  color: rgba(136, 136, 136, 1);
  font-size: 14px
}
</style>
