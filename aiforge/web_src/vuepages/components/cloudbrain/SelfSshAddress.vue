<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title">
        <span :class="required ? 'required' : ''">{{ $t('cloudbrainObj.customPath') }}</span>
      </div>
      <div class="content">
        <div style="display:flex;">
            <div style="flex:1;display: flex;">
              <el-input style="width:75px;" disabled value="https://"></el-input>
              <el-input v-model="endPoint" @input="handleInput" 
                  :class="labelError ? 'error' : ''" style="flex:1" maxlength="32">
              </el-input>
              <el-input style="width:215px;" disabled value=".user.mlunotebook.openi.org.cn"></el-input>
            </div>
            <div style="margin-left: 1rem;">
                {{ $t('cloudbrainObj.port') }}
                <el-input v-model="port" :class="valueError ? 'error' : ''" @input="handleInput" style="width: 100px;"></el-input>
            </div>
        </div>
        <div class="tips">{{ $t('cloudbrainObj.customPathTips') }}</div>
      </div>
    </div>
    <div class="right-area"></div>
  </div>
</template>

<script>
const regex = /^[a-z][a-z0-9]*$/
const regex1 = /^[8][0-8]{3}$/
export default {
  name: 'SelfSshAddress',
  props: {
    value: {
      type: Object,
      default: () => {
        return {
            endPoint:'',
            port:'',
        }
      },
    },
    required: { type: Boolean, default: true },
  },
  data() {
    return {
      errStatus: false,
      endPoint:'',
      port:'',
      labelError:false,
      valueError:false
    };
  },
  watch: {
    value: {
      immediate: true,
      deep: true,
      handler(newVal) {
        this.endPoint = newVal.endPoint
        this.port = newVal.port
      }
    }
  },
  methods: {
    check() {
        if (!this.endPoint.trim() && !this.port.toString().trim()) {
          this.$emit('input', { endPoint: this.endPoint.trim(), port: this.port.toString().trim() });
          return true
        } else if (!this.endPoint.trim() && this.port.toString().trim()) {
          this.labelError = true;
          if (Number(this.port.toString().trim()) > 8800 || Number(this.port.toString().trim()) < 8000) {
            this.valueError = true;
          }
          return false
        } else if (this.endPoint.trim() && !this.port.toString().trim()) {
          this.valueError = true;
          if (!regex.test(this.endPoint.trim())) {
            this.labelError = true;
          }
          return false
        } else {
          if (!regex.test(this.endPoint.trim())) {
            this.labelError = true;
            return false
          }
          if (Number(this.port.toString().trim()) > 8800 || Number(this.port.toString().trim()) < 8000 || !regex1.test(Number(this.port.toString().trim()))) {
            this.valueError = true;
            return false
          }
          
          this.$emit('input', { endPoint: this.endPoint.trim(), port: this.port.toString().trim() });
          return true
        }
        
    },
    handleInput(value) {
        if ( this.endPoint.trim()) {
            this.labelError = false;
        }
        if ( this.port?.toString().trim()) {
            this.valueError = false;
        }
    },
  },
  beforeMount() {
  }
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';
.error {
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
</style>
