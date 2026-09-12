<template>
  <div class="base-dlg">
    <BaseDialog :visible.sync="dialogShow" :width="`750px`" :title="$t('thirdPartyApp.editOauth2AppAuthInfo')"
      @open="open" @opened="opened" @close="close" @closed="closed">
      <div class="dlg-content">
        <div class="form">
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('thirdPartyApp.clientID') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.ClientID" placeholder="" disabled>
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title">
              <span>{{ $t('thirdPartyApp.appName') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.Name" placeholder="" disabled>
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title">
              <span>{{ $t('thirdPartyApp.appCreator') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.UserName" placeholder="" disabled>
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('thirdPartyApp.authStatus') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.status" @change="changeStatus">
                <el-option v-for="item in statusList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row" v-if="dataInfo.status == 2">
            <div class="auths-c">
              <div class="title required">{{ $t('thirdPartyApp.authenticatedCanGetInfoTips') }} ：</div>
              <el-checkbox v-model="authsInfo.checkedAll" :indeterminate="authsInfo.isIndeterminate"
                @change="changeCheckAll">{{ $t('userRole.roleSelectedTips') }}</el-checkbox>
              <el-checkbox-group v-model="dataInfo.auths" @change="changeAuths">
                <el-checkbox v-for="auth in authsList" :label="auth.k" :key="auth.k" :disabled="auth.disabled">{{ auth.v
                  }}</el-checkbox>
              </el-checkbox-group>
            </div>
          </div>
          <div class="form-row" style="margin-top: 20px">
            <div class="content" style="justify-content: center;">
              <el-button type="primary" class="btn confirm-btn" @click="confirm">{{ $t('confirm') }}</el-button>
              <el-button class="btn" @click="cancel">{{ $t('cancel') }}</el-button>
            </div>
          </div>
        </div>
      </div>
    </BaseDialog>
  </div>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';
import { i18n } from '~/langs';
import { setEditApplicaiton } from '~/apis/modules/thirdpartyapp';

const StatusList = [{ k: 2, v: i18n.t('thirdPartyApp.setAuthenticated') }, { k: 3, v: i18n.t('thirdPartyApp.cancelAuthentication') }]

export default {
  name: "EditDialog",
  props: {
    visible: { type: Boolean, default: false },
    title: { type: String, default: '' },
    type: { type: String, defalut: '' },
    data: { type: Object, default: () => ({}) },
  },
  components: {
    BaseDialog
  },
  data() {
    return {
      dialogShow: false,
      statusList: [],
      authsList: [
        { k: 'user.base', v: this.$t('thirdPartyApp.basicInfoAll'), disabled: true },
        { k: 'user.email', v: this.$t('thirdPartyApp.email') },
        { k: 'user.phone', v: this.$t('thirdPartyApp.phoneNumber') },
      ],
      authsInfo: {
        checkedAll: false,
        isIndeterminate: false,
      },
      dataInfo: {},
    };
  },
  watch: {
    visible: function (val) {
      this.dialogShow = val;
    },
  },
  methods: {
    resetDataInfo() {
      this.dataInfo = {
        status: '',
        auths: ['user.base'],
      }
    },
    changeStatus(val) {
      if (val == 3) {
        this.changeCheckAll(false);
      } else {
        this.changeCheckAll(true);
      }
    },
    changeAuths(val) {
      this.refreshCheckBox();
    },
    refreshCheckBox() {
      this.authsInfo.isIndeterminate = this.dataInfo.auths.length && this.dataInfo.auths.length != this.authsList.length;
      this.authsInfo.checkedAll = this.dataInfo.auths.length == this.authsList.length;
    },
    changeCheckAll(val) {
      if (val) {
        this.dataInfo.auths = this.authsList.map(item => item.k);
      } else {
        this.dataInfo.auths = this.authsList.filter(item => item.disabled).map(item => item.k);
      }
      this.refreshCheckBox();
    },
    open() {
      this.resetDataInfo();
      this.dataInfo = Object.assign(this.dataInfo, { ...this.data });
      this.dataInfo.auths = [...this.data.AllowedScopes];
      if (this.type == 'verify' || this.type == 'verify_again') {
        this.statusList = [StatusList[0]];
        this.dataInfo.status = StatusList[0].k;
        this.changeCheckAll(true);
      }
      if (this.type == 'edit') {
        this.statusList = [...StatusList];
        this.dataInfo.status = this.data.VerifyFlag;
      }
      this.refreshCheckBox();
      this.$emit("open");
    },
    opened() {
      this.$emit("opened");
    },
    close() {
      this.$emit("close");
    },
    closed() {
      this.$emit("closed");
      this.$emit("update:visible", false);
    },
    async confirm() {
      const subData = {
        ID: this.dataInfo.ID,
        VerifyFlag: this.dataInfo.status,
        ScopeList: this.dataInfo.auths.join(',')
      }
      if (subData.VerifyFlag == 3) {
        let canSubFlag = false;
        await this.$confirm(this.$t('thirdPartyApp.cancelVerifyTips'), this.$t('tips'), {
          confirmButtonText: this.$t('confirm'),
          cancelButtonText: this.$t('cancel'),
          type: 'warning'
        }).then(() => { canSubFlag = true }).catch(err => { });
        if (!canSubFlag) return;
      }
      // console.log('subData', subData)
      // return;
      setEditApplicaiton(subData).then(res => {
        res = res.data;
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully')
          });
          this.$emit("confirm");
        } else {
          this.$message({
            type: 'error',
            message: res.msg || this.$t('submittedFailed')
          });
        }
      }).catch(err => {
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('submittedFailed')
        });
      })
    },
    cancel() {
      this.dialogShow = false;
      this.$emit("update:visible", false);
    }
  },
  mounted() {
    this.resetDataInfo();
  },
};
</script>

<style scoped lang="less">
.dlg-content {
  margin: 20px 40px 25px 0;
  display: flex;
  justify-content: center;

  .form {
    width: 600px;

    .form-row {
      display: flex;
      min-height: 42px;
      margin-bottom: 4px;

      .title {
        width: 140px;
        display: flex;
        justify-content: flex-start;
        align-items: center;
        margin-right: 20px;
        color: rgb(136, 136, 136);
        font-size: 14px;

        &.required {
          span {
            position: relative;
          }

          span::after {
            position: absolute;
            right: -10px;
            top: -2px;
            vertical-align: top;
            content: '*';
            color: #db2828;
          }
        }
      }

      .content {
        width: 0;
        flex: 1;
        display: flex;
        align-items: center;

        /deep/ .el-select {
          width: 100%;
        }
      }

      .auths-c {
        width: 100%;
        border: 1px solid #DCDFE6;
        border-radius: 4px;
        padding: 12px;
        position: relative;
        margin-top: 20px;

        .title {
          position: absolute;
          top: -10px;
          left: 7px;
          padding: 0 5px;
          width: auto;
          background: white;
        }

        .el-checkbox {
          margin-top: 10px;
          margin-left: 10px;
        }

        .el-checkbox-group {
          display: flex;
          flex-direction: column;
          margin-left: 10px;
          padding-left: 10px;
        }
      }
    }
  }

  .btn {
    color: rgb(2, 0, 4);
    background-color: rgb(194, 199, 204);
    border-color: rgb(194, 199, 204);

    &.confirm-btn {
      color: #fff;
      background-color: rgb(56, 158, 13);
      border-color: rgb(56, 158, 13);
    }
  }
}
</style>
