<template>
  <div class="forbid-penetration-check-container">
    <el-popover ref="penaltyPopover" placement="right" trigger="hover" popper-class="penalty-tip-popover">
      <div class="penalty-tip-content" v-html="$t('cloudbrainObj.penaltyTipContent')"></div>
      <el-checkbox slot="reference" v-model="currentValue" @change="handleChange">
        <span class="checkbox-label">{{ $t('cloudbrainObj.forbidPenetrationStatement') }}</span>
      </el-checkbox>
    </el-popover>
  </div>
</template>

<script>
export default {
  name: 'ForbidPenetrationTipCheck',
  props: {
    value: { type: Boolean, required: true },
    required: { type: Boolean, default: false },
  },
  data() {
    return {
      currentValue: false,
      errStatus: false,
    };
  },
  watch: {
    value: {
      immediate: true,
      handler(newVal) {
        this.currentValue = !!newVal;
      }
    }
  },
  methods: {
    check() {
      const isValid = this.currentValue || !this.required;
      this.errStatus = !isValid;
      return isValid;
    },
    handleChange(value) {
      this.currentValue = value;
      this.$emit('input', value);
      this.check();
    },
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
.forbid-penetration-check-container {
  display: flex;
  align-items: center;
  padding: 0;
  margin-left: 154px;
  margin-right: 24px;
  margin-bottom: 26px;

  .el-checkbox {
    margin-right: 0;

    /deep/ .el-checkbox__input {
      margin-top: -1px;
    }

    /deep/ .el-checkbox__label {
      padding-left: 10px;
      line-height: normal;
      white-space: break-spaces;
    }

    /deep/ .el-checkbox__input.is-checked+.el-checkbox__label {
      /deep/ .el-checkbox__inner {
        background-color: #f04848;
        border-color: #f04848;
      }
    }
  }

  .checkbox-label {
    color: #f04848;
    font-size: 14px;
    text-decoration: underline;
  }
}

@media (max-width: 768px) {
  .forbid-penetration-check-container {
    margin-left: 0;
  }
}
</style>

<style lang="less">
.penalty-tip-popover {
  padding: 12px 16px !important;
  background-color: #fff7f0 !important;
  border: 1px solid #f2711c !important;
  border-radius: 6px !important;
  color: #101010 !important;

  .popper__arrow {
    border-right-color: #f2711c !important;
    border-width: 8px;
    left: -8px !important;

    &::after {
      border-right-color: #f2711c !important;
    }
  }

  .penalty-tip-content {
    max-width: 320px;
    color: #101010;
    font-size: 13px;
    line-height: 1.6;
    word-break: normal;

    /deep/.highlight {
      color: #f04848;
      font-weight: 600;
    }
  }
}
</style>
