<template>
  <div class="task-name">
    <div class="title">
      <span :class="required ? 'required' : ''">{{ title }}</span>
    </div>
    <div class="content" :class="errStatus ? 'error' : ''">
      <el-input class="name-input" v-model="taskName" @input="check" maxlength="36" placeholder="任务名称"></el-input>
      <div class="tips">只能以小写字母或数字开头且只包含小写字母、数字、_和-，不能以_结尾，最长36个字符。</div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs';

export default {
  name: "TaskName",
  props: {
    title: { type: String, default: "任务名称" },
    required: { type: Boolean, default: true },
    userName: { type: String, default: "" },
    oriData: { type: String, default: "" },
  },
  data() {
    return {
      taskName: '',
      errStatus: false,
    };
  },
  watch: {
    oriData: function (val) {
      this.taskName = val;
    },
  },
  methods: {
    check() {
      const reg = /^[a-z0-9][a-z0-9\-_]{1,34}[a-z0-9\-]$/;
      this.errStatus = !reg.test(this.taskName);
      return !this.errStatus;
    },
    generateName() {
      let str = this.userName.toLocaleLowerCase();
      const reg1 = /[^a-z0-9_\-]+/g;
      const reg2 = /^[_\-]+/g;
      const reg3 = /[_]+$/g;
      str = str.replace(reg1, '').replace(reg2, '').replace(reg3, '');
      str = str.slice(0, 5);
      const now = Date.now();
      return str + dayjs(now).format('YYYYMMDDHH') + (now / 1000).toFixed(0).slice(-5);
    },
    getName() {
      return this.taskName;
    }
  },
  mounted() {
    if (!this.oriData) {
      this.taskName = this.generateName();
    }
  },
};
</script>

<style scoped lang="less">
.task-name {
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

    .name-input {
      width: 100%;
    }

    .tips {
      font-size: 12px;
      color: rgba(136, 136, 136, 1);
      margin-top: 10px;
    }

    &.error {
      .name-input {
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
</style>
