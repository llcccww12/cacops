<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title">
        <span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.runParameter') }}</span>
      </div>
      <div class="content">
        <div class="param-list field-input">
          <div class="param-item" v-for="(item) in list" :key="item.id">
            <div class="param-k-v">
              <el-input class="param-k" :class="item.labelError ? 'error' : ''"
                :placeholder="$t('cloudbrainObj.parameterName')" @input="handleInput(item)"
                v-model="item.label"></el-input>
              <el-input class="param-v" :class="item.valueError ? 'error' : ''"
                :placeholder="$t('cloudbrainObj.parameterValue')" @input="handleInput(item)"
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
            <span>{{ $t('cloudbrainObj.addRunParameter') }}</span>
          </a>
        </div>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>
export default {
  name: "RunParameters",
  props: {
    value: { type: Array, required: true, },
    required: { type: Boolean, default: false },
  },
  data() {
    return {
      list: [],
    };
  },
  watch: {
    value: {
      immediate: true,
      deep: true,
      handler(newVal) {
        newVal = newVal || [];
        this.list = newVal.map((item) => {
          return {
            id: Math.random(),
            label: item.label,
            value: item.value,
            labelError: false,
            valueError: false,
            ...item,
          }
        })
      }
    }
  },
  methods: {
    addParameter() {
      this.list.push({
        id: Math.random(),
        label: '',
        value: '',
        labelError: false,
        valueError: false,
      });
      this.$emit('input', this.list);
      this.$emit('change', this.list);
    },
    removeParameter(item) {
      const index = this.list.findIndex((_item) => _item.id == item.id);
      this.list.splice(index, 1);
      this.$emit('input', this.list);
      this.$emit('change', this.list);
    },
    handleInput(item) {
      if (item.label.trim() != '') {
        item.labelError = false;
      }
      if (item.value.trim() != '') {
        item.valueError = false;
      }
      this.$emit('input', this.list);
      this.$emit('change', this.list);
    },
    check() {
      let hasError = false;
      for (let i = 0, iLen = this.list.length; i < iLen; i++) {
        const item = this.list[i];
        item.label = item.label.trim();
        item.value = item.value.trim();
        if (item.label == '') {
          hasError = true;
          item.labelError = true;
        } else {
          item.labelError = false;
        }
        if (item.value == '') {
          hasError = true;
          item.valueError = true;
        } else {
          item.valueError = false;
        }
      }
      this.$emit('input', this.list);
      this.$emit('change', this.list);
      return !hasError;
    }
  },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.form-row {

  .content {
    flex: 1;

    .add-param-btn {
      margin-left: 2px;
      margin-top: 8px;

      a {
        color: rgb(0, 102, 255);
      }
    }

    .param-list {

      .param-item {
        display: flex;
        align-items: center;
        margin-bottom: 12px;
        position: relative;

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
          position: absolute;
          right: -28px;
          width: 20px;

          i {
            cursor: pointer;
          }
        }
      }
    }
  }
}
</style>
