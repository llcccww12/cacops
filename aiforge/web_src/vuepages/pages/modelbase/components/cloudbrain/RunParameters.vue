<template>
  <div class="run-parameters">
    <div class="title">
      <span :class="required ? 'required' : ''">{{ title }}</span>
    </div>
    <div class="content">
      <div class="param-list">
        <div class="param-item" v-for="(item, index) in list" :key="item.id">
          <div class="param-k-v">
            <el-input class="param-k" readonly :class="item.keyError ? 'error' : ''" placeholder="参数名" @input="inputChange(item)"
              v-model="item.key"></el-input>
            <el-input class="param-v" :class="item.valueError ? 'error' : ''" placeholder="参数值" @input="inputChange(item)"
              v-model="item.value"></el-input>
          </div>
          <div class="param-del-btn">
            <i class="trash icon" @click="removeParameter(item)"></i>
          </div>
        </div>
      </div>
      <div class="add-param-btn">
        <a href="javascript:;" @click="addParameter">
          <i class="plus square outline icon"></i>
          <span>新增运行参数</span>
        </a>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "RunParameters",
  props: {
    title: { type: String, default: "运行参数" },
    required: { type: Boolean, default: false },
    oriData: { type: Array, default: () => [] },
  },
  data() {
    return {
      list: [],
    };
  },
  watch: {
    oriData: {
      immediate: true,
      deep: true,
      handler(newVal) {
        this.list = newVal.map((item) => {
          return {
            id: Math.random(),
            key: item.k,
            value: item.v,
            keyError: false,
            valueError: false,
          }
        })
      }
    }
  },
  methods: {
    addParameter() {
      this.list.push({
        id: Math.random(),
        key: '',
        value: '',
        keyError: false,
        valueError: false,
      })
    },
    removeParameter(item) {
      const index = this.list.findIndex((_item) => _item.id == item.id);
      this.list.splice(index, 1);
    },
    inputChange(item) {
      if (item.key.trim() != '') {
        item.keyError = false;
      }
      if (item.value.trim() != '') {
        item.valueError = false;
      }
    },
    check() {
      let hasError = false;
      for (let i = 0, iLen = this.list.length; i < iLen; i++) {
        const item = this.list[i];
        if (item.key.trim() == '') {
          hasError = true;
          item.keyError = true;
        } else {
          item.keyError = false;
        }
        if (item.value.trim() == '') {
          hasError = true;
          item.valueError = true;
        } else {
          item.valueError = false;
        }
      }
      return !hasError;
    },
    getParameter() {
      return this.list.map(item => {
        return {
          label: item.key,
          value: item.value
        }
      })
    },
  },
};
</script>

<style scoped lang="less">
.run-parameters {
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

    .add-param-btn {
      margin-left: 2px;
      margin-top: 6px;
      display: none;

      a {
        color: rgba(3, 102, 214, 100);
      }
    }

    .param-list {

      .param-item {
        display: flex;
        align-items: center;
        margin-bottom: 12px;

        .param-k-v {
          flex: 1;
          display: flex;

          .param-k {
            margin-right: 10px;

            &.error {
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

          .param-v {
            margin-right: 10px;

            &.error {
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

        .param-del-btn {
          width: 20px;
          display: none;

          i {
            cursor: pointer;
          }
        }
      }
    }
  }
}
</style>
