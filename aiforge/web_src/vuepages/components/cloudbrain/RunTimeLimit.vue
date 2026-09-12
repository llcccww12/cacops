<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title">
        <span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.taskIsAutomaticStop') }}</span>
      </div>
      <div class="content">
        <el-select class="spec-sel field-input" v-model="type" @change="changeType">
          <el-option :value="1" :label="$t('cloudbrainObj.automaticStop')" key="1"></el-option>
          <el-option :value="2" :label="$t('cloudbrainObj.manualStop')" key="2"></el-option>
        </el-select>
        <div v-if="type == 1" class="tips">{{ $t('cloudbrainObj.automaticStopTips') }}</div>
        <div v-if="type == 2" class="tips">{{ $t('cloudbrainObj.manualStopTips') }}</div>
        <el-radio-group class="time-group-sel" v-if="type == 1" v-model="limitTime" @change="changeTime">
          <el-radio :label="item" v-for="item in limitTimeList" :key="item">
            {{ $t('cloudbrainObj.numOfHours', { num: item }) }}
          </el-radio>
          <br>
          <el-radio class="custom" :label="-1">{{ $t('cloudbrainObj.customize') }}</el-radio>
          <div class="limit-time-inp-c content" :class="errStatus ? 'error' : ''" v-if="type == 1 && limitTime == -1">
            <div class="limit-time-inp-wrap">
              <el-input class="limit-time-inp field-input" v-model="limitTimeInputValue"
                @input="handleInput"></el-input>
              <div class="unit">
                {{ $t('cloudbrainObj.hours') }} ({{ $t('cloudbrainObj.customizeTimeLimitPlaceholder') }})
              </div>
            </div>
          </div>
        </el-radio-group>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>

export default {
  name: 'RunTimeLimit',
  props: {
    value: { type: String, required: true },
    required: { type: Boolean, default: true },
  },
  data() {
    return {
      type: 1,
      limitTime: 4,
      limitTimeList: [1, 2, 4, 6],
      limitTimeInputValue: '',
      currentValue: '',
      errStatus: false,
    };
  },
  watch: {
    value: {
      immediate: true,
      handler(newVal) {
        newVal = newVal === undefined ? '' : newVal;
        this.currentValue = newVal.toString();
        if (this.currentValue == '-1') {
          this.type = 2;
        } else if (this.limitTimeList.indexOf(Number(this.currentValue)) >= 0 && this.limitTime != -1) {
          this.type = 1;
          this.limitTime = Number(this.currentValue);
        } else {
          this.type = 1;
          this.limitTime = -1
          this.limitTimeInputValue = this.currentValue;
        }
      }
    }
  },
  methods: {
    check() {
      if (this.type == 2) {
        this.currentValue = '-1';
        this.errStatus = false;
      } else {
        if (this.limitTime == -1) {
          const value = parseInt(this.limitTimeInputValue);
          if (value >= 1 && value <= 24) {
            this.limitTimeInputValue = value;
            this.currentValue = this.limitTimeInputValue.toString();
            this.errStatus = false;
          } else {
            this.currentValue = '';
            this.limitTimeInputValue = '';
            this.errStatus = true;
          }
        } else {
          this.currentValue = this.limitTime.toString();
          this.errStatus = false;
        }
      }
      return !this.errStatus;
    },
    changeType() {
      this.check();
      this.errStatus = false;
      this.$emit('input', this.currentValue);
    },
    changeTime() {
      this.check();
      this.errStatus = false;
      this.$emit('input', this.currentValue);
    },
    handleInput() {
      const value = parseInt(this.limitTimeInputValue);
      if (value >= 1 && value <= 24) {
        this.limitTimeInputValue = value;
      } else {
        this.limitTimeInputValue = this.limitTimeInputValue.slice(0, -1);
      }
      this.check();
      this.$emit('input', this.currentValue);
    },
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.time-group-sel {
  margin-top: 12px;
  width: 100%;

  /deep/ .el-radio {
    margin-bottom: 16px;
  }

  .custom {
    margin-right: 10px;
  }

  .limit-time-inp-c {
    display: inline-block;
    position: relative;
    width: auto !important;

    .limit-time-inp-wrap {
      display: flex;
      align-items: center;

      /deep/.el-input__inner {
        height: 32px !important;
        line-height: 32px !important;
        font-size: 13px !important;
        text-align: center;
      }
    }

    .limit-time-inp {
      width: 60px !important;
    }

    .unit {
      margin-left: 5px;
      font-size: 14px;
      color: #606266;
    }
  }
}
</style>
