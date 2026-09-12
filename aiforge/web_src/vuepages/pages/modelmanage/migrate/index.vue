<template>
  <div>
    <div class="ui middle very relaxed page grid content-wrap" style="overflow-x: auto;">
      <div class="column">
        <div class="form-wrap">
          <div class="header">
            <span class="title">{{ $t('modelObj.migrate_external_model') }}</span>
          </div>
          <div class="content">
            <div class="row-c">
              <div class="row" :class="nameErr ? 'error' : ''">
                <div class="r-title flex-start"><label class="required">{{ $t('modelObj.external_model_name') }}</label>
                </div>
                <div class="r-content">
                  <el-input size="medium" :maxLength="255" v-model="state.externalName" @blur="checkExternalName"
                    :placeholder="$t('modelManage.pleaseInputModelName')">
                  </el-input>
                  <div class="tips">{{ $t('modelObj.migrate_external_model_name_tips') }}</div>
                </div>
              </div>
              <div class="row" v-if="state.externalName && !this.nameErr && state.name">
                <div class="r-title"><label>{{ $t('modelManage.modelName') }}</label></div>
                <div class="r-content r-wrap"><span>{{ state.name }}</span></div>
              </div>
              <div class="row" v-if="state.externalName && !this.nameErr && state.url">
                <div class="r-title"><label>{{ $t('modelObj.openi_model_repo_url') }}</label></div>
                <div class="r-content r-wrap"><span>{{ state.url }}</span></div>
              </div>
              <div class="row">
                <div class="r-title"><label class="required">{{ $t('modelManage.modelEngine') }}</label></div>
                <div class="r-content">
                  <el-select style="width:100%;" size="medium" v-model="state.engine" placeholder="">
                    <el-option v-for="item in engineList" :key="item.k" :label="item.v" :value="item.k">
                    </el-option>
                  </el-select>
                </div>
              </div>
              <div class="row" v-if="false">
                <div class="r-title"><label>{{ $t('modelManage.license') }}</label></div>
                <div class="r-content">
                  <el-select style="width:100%;" size="medium" v-model="state.license" class="license-sel"
                    :placeholder="$t('modelManage.selectLicense')">
                    <template slot="prefix">
                      <i v-if="state.license" class="el-select__caret el-input__icon el-icon-close"
                        @click.stop.prevent="handleClearLicenseClick"></i>
                    </template>
                    <el-option v-for="item in licenseList" :key="item.id" :label="item.name" :value="item.id">
                    </el-option>
                  </el-select>
                </div>
              </div>
              <div class="row">
                <div class="r-title"><label>{{ $t('modelManage.modelLabel') }}</label></div>
                <div class="r-content">
                  <el-input size="medium" :maxLength="255" v-model="state.label"
                     @input="labelInput"></el-input>
                </div>
              </div>
              <div class="row">
                <div class="r-title"><label>{{ $t('modelManage.externalToken') }}</label></div>
                <div class="r-content">
                  <el-input size="medium" :maxLength="255" v-model="state.hf_token"
                    :placeholder="$t('modelManage.modelLabelGatedTips')"></el-input>
                </div>
              </div>
              <div class="row">
                <div class="r-title"><label>{{ $t('modelManage.modelAccess') }}</label></div>
                <div class="r-content">
                  <el-radio v-model="state.isPrivate" label="0">{{
                    $t('modelManage.modelAccessPublic') }}<span>
                      ({{ $t('modelObj.migrate_model_should_be_public') }})</span>
                  </el-radio>
                </div>
              </div>
              
              <div class="row" style="margin-top:20px">
                <div class="r-title"><label></label></div>
                <div class="r-content btn-c">
                  <el-button size="medium" class="green" @click="submit">{{ $t('modelObj.migrate_model') }}</el-button>
                  <el-button size="medium" @click="cancel">{{ $t('modelManage.cancel') }}</el-button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <LoadingMask :loading="maskLoading" :content="maskLoadingContent"></LoadingMask>
  </div>
</template>

<script>
import { setModelMigrate, getModelLicenseList } from '~/apis/modules/modelmanage';
import { MODEL_ENGINES } from '~/const';
import LoadingMask from '~/components/cloudbrain/LoadingMask.vue';

const MAX_LABEL_COUNT = 5;
const ModelOwner = window._HfModelOwner;

export default {
  data() {
    return {
      state: {
        aimodel_type: 2,
        externalName: '',
        name: '',
        url: '',
        version: '0.0.1',
        engine: 0,
        label: '',
        license: '',
        isPrivate: '0',
        hf_token: '',
      },
      licenseList: [],
      nameErr: false,
      engineList: MODEL_ENGINES,
      maskLoading: false,
      maskLoadingContent: '',
      oldUrl: '',
    };
  },
  components: { LoadingMask },
  methods: {
    checkExternalName() {
      const reg = /^[\.a-zA-Z0-9_-]+\/[\.a-zA-Z0-9_-]+$/;
      if (reg.test(this.state.externalName)) {
        this.state.name = this.state.externalName.split('/')[1];
        this.oldUrl = window.origin + '/' + ModelOwner + '/' + this.state.externalName.split('/')[0];
        this.state.url = `${window.origin}/models/detail/${ModelOwner}/${this.state.name}`
        this.nameErr = false;
      } else {
        this.state.name = '';
        this.state.url = '';
        this.nameErr = true;
      }
      return !this.nameErr;
    },
    labelInput() {
      const hasEndSpace = this.state.label[this.state.label.length - 1] == ' ';
      const list = this.state.label.trim().split(' ').filter(label => label != '');
      this.state.label = list.slice(0, MAX_LABEL_COUNT).join(' ') + (hasEndSpace && list.length < MAX_LABEL_COUNT ? ' ' : '');
    },
    handleClearLicenseClick() {
      this.state.license = '';
    },
    cancel() {
      window.history.back();
    },
    goMigrating(id) {
      window.location.href = `${window.origin}/models/detail/model_migrating?model_id=${id}`;
    },
    getDetailUrl() {
      return this.state.url
    },
    submit() {
      this.state.name = this.state.name.trim();
      if (!this.checkExternalName()) {
        this.$message({
          type: 'info',
          message: this.$t('modelObj.please_enter_right_external_model_name'),
        });
        return;
      }
      const subData = {
        ...this.state,
        hf_repo_id: this.state.externalName,
        engine: Number(this.state.engine),
        label: this.state.label.split(/\s+/).join(' ').trim(),
        isPrivate: false,
      };
      this.maskLoading = true;
      setModelMigrate(subData).then(res => {
        res = res.data;
        console.log(res)
        // console.log(c)
        if (res.code == '1') { // success
          console.log("xxxxxxxxxxxx")
          if(res.data && res.data.id){
            console.log("xxxxxxxxxxxx1")
            this.goMigrating(res.data.id);
          }
        } else if (res.code == '2') { // exists
          this.maskLoading = false;
          const url = this.getDetailUrl();
          this.$alert(this.$t('modelObj.create_model_migrate_exists', { url }), this.$t('tips'), {
            dangerouslyUseHTMLString: true,
          });
        } else if (res.code == '-1') { // error
          this.maskLoading = false;
          this.$message({
            type: 'error',
            message: res.msg,
          });
        } else {
          this.maskLoading = false;
          this.$message({
            type: 'error',
            message: this.$t('modelObj.migrate_external_model_failed'),
          });
        }
      }).catch(err => {
        this.maskLoading = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('modelObj.migrate_external_model_failed'),
        });
      });
    },
  },
  beforeMount() {
    this.maskLoadingContent = this.$t('modelObj.create_model_migrate_loading_content');
    // getModelLicenseList().then(res => {
    //   res = res.data;
    //   try {
    //     const license = JSON.parse(res) || [];
    //     this.licenseList = license;
    //   } catch (err) {
    //     console.log(err);
    //   }
    // }).catch(err => {
    //   console.log(err);
    // });
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.content-wrap {
  margin-top: 26px;

  .form-wrap {
    width: 800px;
    margin: auto;
  }

  .header {
    height: 46.7px;
    border-color: rgb(212, 212, 213);
    border-width: 1px;
    border-style: solid;
    border-radius: 5px 5px 0px 0px;
    font-size: 14px;
    background: rgb(240, 240, 240);
    display: flex;
    align-items: center;
    justify-content: center;

    .title {
      font-weight: 600;
      font-size: 18px;
      color: rgb(16, 16, 16);
    }
  }

  .content {
    margin-top: -1px;
    border-color: rgb(212, 212, 213);
    border-width: 1px;
    border-style: solid;
    padding: 30px 0;
    border-top: none;

    .row-c {
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
      margin: 0 auto;
      width: 80%;

      .row {
        width: 100%;
        display: flex;
        align-items: center;
        margin: 8px 0;
        margin-left: -190px;

        .r-title {
          text-align: right;
          font-size: .92857143em;
          font-weight: 700;
          color: rgba(0, 0, 0, .87);
          width: 200px;
          margin-right: 28px;
          position: relative;

          .required {
            &::after {
              position: absolute;
              margin: -0.2em 0 0 0.2em;
              content: '*';
              color: #db2828;
            }
          }

          &.flex-start {
            align-self: flex-start;
            margin-top: 10px;
          }
        }

        &.error {
          .r-title {
            color: #9f3a38;
          }

          .r-content {
            /deep/.el-input__inner {
              color: #9f3a38;
              background: #fff6f6;
              border-color: #e0b4b4;

              &::placeholder {
                color: #e0b4b4;
              }
            }
          }
        }

        .r-content {
          flex: 1;
          width: 0;

          .tips {
            color: rgb(153, 153, 153);
            font-size: 14px;
            margin-top: 6px;
            line-height: 20px;
          }

          .cluster-type-btn {
            display: flex;
            align-items: center;
            justify-content: center;
            border: 1px solid #DCDFE6;
            height: 36px;
            padding: 10px;
            cursor: pointer;
            border-radius: 4px;

            .icon {
              margin-right: 5px;
            }

            &.focused {
              border-color: rgb(50, 145, 248);
              color: rgb(50, 145, 248);
              cursor: default;

              .icon {
                :not([stroke]) {
                  fill: rgb(50, 145, 248);
                }
              }
            }
          }

          &.r-wrap {
            word-break: break-all;
          }
        }
      }
    }
  }
}

.license-sel {
  /deep/.el-input--prefix .el-input__inner {
    padding-left: 15px;
  }

  /deep/.el-input__prefix {
    position: absolute;
    right: 0;

    .el-icon-close {
      position: absolute;
      right: 24px;
      color: rgba(0, 0, 0, .87);
      font-weight: bold;
    }
  }
}

.input-disabled {
  /deep/ .el-input__inner {
    background-color: #f5f5f6 !important;
    color: #888888 !important;
  }
}

.el-select-dropdown__item.selected {
  color: rgba(0, 0, 0, .95);
}

.btn-c {
  /deep/ .el-button {
    background-color: #e0e1e2;
    color: rgba(0, 0, 0, .6);
    border-color: transparent;
    transition: opacity .1s ease, background-color .1s ease, color .1s ease, box-shadow .1s ease, background .1s ease, -webkit-box-shadow .1s ease;
    will-change: auto;
    -webkit-tap-highlight-color: transparent;

    &:hover {
      border-color: transparent;
      background-color: #cacbcd;
      color: rgba(0, 0, 0, .8);
    }

    &:focus {
      background-color: #cacbcd;
      color: rgba(0, 0, 0, .8);
      border-color: transparent;
    }

    &:active {
      background-color: #babbbc;
      color: rgba(0, 0, 0, .9);
      border-color: transparent;
    }

    &.green {
      background-color: #5bb973;
      color: #fff;

      &:hover {
        background-color: #16ab39;
        border-color: transparent;
      }

      &:focus {
        background-color: #0ea432;
        border-color: transparent;
      }

      &:active {
        background-color: #198f35;
        border-color: transparent;
      }
    }
  }
}

/deep/ .el-select {
  .is-focus {
    .el-input__inner {
      border-color: #85b7d9;
    }
  }
}

/deep/ .el-input__inner {
  &:focus {
    border-color: #85b7d9;
  }
}

/deep/ .el-textarea__inner {
  &:focus {
    border-color: #85b7d9;
  }
}

/deep/ .el-radio.is-checked {
  .el-radio__inner {
    border-color: rgb(16, 16, 16);
    background: rgb(16, 16, 16);
  }

  .el-radio__label {
    color: rgb(16, 16, 16);
  }
}
</style>
