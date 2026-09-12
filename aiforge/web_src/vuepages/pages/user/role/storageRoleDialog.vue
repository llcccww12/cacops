<template>
  <div class="base-dlg">
    <BaseDialog :visible.sync="dialogShow" width="50%" :title="dialogTitle" @open="open" @opened="opened" @close="close"
      @closed="closed">
      <div class="dlg-content">
        <div class="form">
          <div class="form-row">
            <label class="required" for="">{{ $t('userRole.roleType') }}：</label>
            <div class="content">
              <span style="color:#101010">{{ $t('userRole.storageCategory') }}</span>
            </div>
          </div>
          <div class="form-row">
            <label class="required" for="">{{ $t('userRole.roleName') }}：</label>
            <div class="content">
              <el-input v-model="params.name" :disabled="type === 'view'" style="width:60%;" maxlength="80"
                :placeholder="$t('resourcesManagement.roleNameTips')">
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <label class="required">{{ $t('userRole.isItDefault') }}：</label>
            <div class="content">
              <el-select v-model="params.isCommon" :disabled="type !== 'add'" style="width:20%;">
                <el-option v-for="item in commonTypeList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row">
            <label for="">{{ $t('userRole.roleDescription') }}：</label>
            <div class="content">
              <el-input type="textarea" :rows="2" v-model="params.description" :disabled="type === 'view'"
                style="width:60%;" maxlength="800" :placeholder="$t('resourcesManagement.roleDescTips')">
              </el-input>
            </div>
          </div>
          <div class="form-row border">
            <div class="title" style="position:absolute">请指定角色拥有的权限（可多选）</div>
            <div class="content box">
              <el-checkbox :disabled="type === 'view'" :indeterminate="isIndeterminate" v-model="checkAll"
                @change="handleCheckAllChange">{{ $t('userRole.roleSelectedTips') }}</el-checkbox>
              <el-checkbox-group :disabled="type === 'view'" v-model="checkedItems" @change="handleCheckedChange">
                <el-checkbox label="operNum">
                  数据集和模型存储配额一共&nbsp;&nbsp;<el-input v-model="params.operNum"
                    @input="params.operNum = parseInt($event) || 0"
                    :disabled="type === 'view' || !checkedItems.includes('operNum')" size="mini"
                    style="width:80px"></el-input>&nbsp;GiB
                </el-checkbox>
                <el-checkbox label="codeSize">
                  运行的计算任务容器内code目录存储配额&nbsp;&nbsp;<el-input v-model="params.codeSize"
                    @input="params.codeSize = parseInt($event) || 0"
                    :disabled="type === 'view' || !checkedItems.includes('codeSize')" size="mini"
                    style="width:80px"></el-input>&nbsp;GB
                </el-checkbox>
                <el-checkbox label="outputSize">
                  运行的计算任务容器内output目录存储配额&nbsp;&nbsp;<el-input v-model="params.outputSize"
                    @input="params.outputSize = parseInt($event) || 0"
                    :disabled="type === 'view' || !checkedItems.includes('outputSize')" size="mini"
                    style="width:80px"></el-input>&nbsp;GB
                </el-checkbox>
              </el-checkbox-group>
            </div>
          </div>
          <div class="btn-wrap" v-if="type !== 'view'">
            <el-button type="primary" :disabled="type === 'view'" class="btn confirm-btn" @click="confirm">{{
              $t('confirm') }}</el-button>
            <el-button class="btn" @click="cancel">{{ $t('cancel') }}</el-button>
          </div>
        </div>
      </div>
    </BaseDialog>
  </div>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';
import { addAiforgeRole, updateAiforgeRole } from '~/apis/modules/resources';
export default {
  name: "opRoleDialog",
  props: {
    visible: { type: Boolean, default: false },
    type: { type: String, defalut: 'add' },
    data: { type: Object, default: () => ({}) },
  },
  components: {
    BaseDialog
  },
  data() {
    return {
      dialogShow: false,
      params: {
        name: "",
        type: 2,
        isCommon: 0,
        description: '',
        operNum: 0,
        codeSize: 0,
        outputSize: 0,
      },
      checkedItems: [],
      checkAll: false,
      isIndeterminate: false,
      commonTypeList: [{ k: 0, v: this.$t('userRole.default'), }, { k: 1, v: this.$t('userRole.notDefault') }],
    };
  },
  watch: {
    visible: function (val) {
      this.dialogShow = val;
    },
  },
  computed: {
    dialogTitle() {
      switch (this.type) {
        case 'add':
          return this.$t('userRole.newStorageRole')
        case 'edit':
          return this.$t('userRole.editStorageRole')
        case 'view':
          return this.$t('userRole.viewStorageRole')
      }
    },
  },
  methods: {
    initStorageParams() {
      if (this.type === 'add') {
        this.restParams()
      } else {
        this.restorePamras()
      }
    },
    restorePamras() {
      this.params.name = this.data.Name
      this.params.description = this.data.Description
      this.params.operNum = this.data.operNum[0] || 0
      this.params.isCommon = this.data.IsCommon
      this.params.codeSize = this.data.codeSize || 0
      this.params.outputSize = this.data.outputSize || 0
      const checked = []
      if (parseInt(this.data.operNum[0]) > 0) checked.push('operNum')
      if (parseInt(this.data.codeSize) > 0) checked.push('codeSize')
      if (parseInt(this.data.outputSize) > 0) checked.push('outputSize')
      this.checkedItems = checked
      this.handleCheckedChange(checked)
    },
    restParams() {
      this.params = { name: "", type: 2, isCommon: 0, description: '', operNum: 0, codeSize: 0, outputSize: 0, }
      this.checkedItems = []
      this.checkAll = false
      this.isIndeterminate = false
    },
    handleCheckedChange(val) {
      const total = 3
      this.checkAll = val.length === total
      this.isIndeterminate = val.length > 0 && val.length < total
    },
    handleCheckAllChange(val) {
      this.checkedItems = val ? ['operNum', 'codeSize', 'outputSize'] : []
      this.isIndeterminate = false
    },
    open() {
      this.initStorageParams()
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
    addStorageRole(params) {
      addAiforgeRole(params).then((res) => {
        if (res.data.code == 0) {
          this.$message.success(this.$t('submittedSuccessfully'))
          this.dialogShow = false
          this.$emit("update:visible", false);
          this.$emit("refresh");
        } else {
          this.$message.error(res.data.msg)
        }
      }).catch((err) => {
        this.dialogShow = false
        this.$emit("update:visible", false);
        this.$message.error(err)
      })
    },
    confirm() {
      if (this.params.name === '') {
        this.$message.error("请输入角色名称")
        return
      }
      if (!this.checkedItems.includes('operNum')) this.params.operNum = 0
      if (!this.checkedItems.includes('codeSize')) this.params.codeSize = 0
      if (!this.checkedItems.includes('outputSize')) this.params.outputSize = 0
      if (this.type === 'add') {
        if (!this.params.isCommon) {
          this.$confirm('当前创建角色应用于所有用户', "提示", {
            confirmButtonText: this.$t('confirm1'),
            cancelButtonText: this.$t('cancel'),
            type: 'warning',
            lockScroll: false,
            closeOnClickModal: false,
            showClose: false,
            closeOnPressEscape: false,
          }).then(() => {
            this.addStorageRole(this.params)
          }).catch(() => {

          })
        } else {
          this.addStorageRole(this.params)
        }

      } else {
        updateAiforgeRole({ id: this.data.ID, ...this.params }).then((res) => {
          if (res.data.code == 0) {
            this.$message.success(this.$t('submittedSuccessfully'))
            this.dialogShow = false
            this.$emit("update:visible", false);
            this.$emit("refresh");
          } else {
            this.$message.error(res.data.msg)
          }
        }).catch((err) => {
          this.dialogShow = false
          this.$emit("update:visible", false);
          this.$message.error(err)
        })
      }

    },
    cancel() {
      this.dialogShow = false;
      this.$emit("update:visible", false);
    }
  },
  mounted() {
  },
};
</script>

<style scoped lang="less">
.dlg-content {
  margin: 24px 0px;

  .form {
    .form-row {
      display: flex;
      margin-bottom: 20px;
      align-items: center;

      &.border {
        border: 1px solid #d4d4d5;
        border-radius: 5px;
        min-height: 100px;
        position: relative;
        margin: 32px 80px 32px 80px;

        .title {
          position: absolute;
          top: -10px;
          left: 12px;
          padding: 0 4px;
          background: #fff;
          color: #101010;
        }
      }

      label {
        &.required::before {
          content: "*";
          color: red;
          margin-right: 5px;
        }

        width: 170px;
        color: rgba(136, 136, 136, 1);
        font-size: 14px;
        box-sizing: border-box;
        text-align: right;
      }

      .content {
        flex: 1;

        &.box {
          padding: 20px;

          .el-checkbox {
            display: flex;
            margin-bottom: 12px;
            align-items: center;
          }

          .el-checkbox-group {
            margin-left: 40px;
            margin-top: 20px;
          }
        }

        &.error {

          /deep/.el-input__inner,
          /deep/.el-textarea__inner {
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

        .tips {
          font-size: 12px;
          color: rgba(136, 136, 136, 1);
          margin-top: 10px;
        }
      }
    }
  }

  .btn-wrap {
    margin-left: 170px;

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

}

/deep/ .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color: #409eff;
  border-color: #DCDFE6;
}

/deep/ .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner::after {
  border-color: #fff;
}

/deep/ .el-input.is-disabled .el-input__inner {
  color: #101010
}

/deep/ .el-checkbox__input.is-disabled+span.el-checkbox__label {
  color: #101010
}
</style>
