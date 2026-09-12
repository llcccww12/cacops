<template>
  <div class="base-dlg">
    <BaseDialog :visible.sync="dialogShow" width="50%"
      :title="dialogTitle"
      @open="open" @opened="opened" @close="close" @closed="closed">
      <div class="dlg-content">
        <div class="form">
            <div class="form-row">
                <label class="required" for="">{{$t('userRole.roleType')}}：</label>
                <div class="content">
                    <span style="color:#101010">{{$t('userRole.containerStorageCategory')}}</span>
                </div>
            </div>
            <div class="form-row">
                <label class="required" for="">{{$t('userRole.roleName')}}：</label>
                <div class="content">
                    <el-input v-model="params.name" :disabled="type==='view'" style="width:60%;"
                     maxlength="80" :placeholder="$t('resourcesManagement.roleNameTips')">
                    </el-input>
                </div>
            </div>
            <div class="form-row">
                <label class="required">{{$t('userRole.isItDefault')}}：</label>
                <div class="content">
                    <el-select v-model="params.isCommon" :disabled="type !== 'add'" style="width:20%;">
                        <el-option v-for="item in commonTypeList" :key="item.k" :label="item.v" :value="item.k" />
                    </el-select>
                </div>
            </div>
            <div class="form-row">
                <label for="">{{$t('userRole.roleDescription')}}：</label>
                <div class="content">
                    <el-input type="textarea" :rows="2" v-model="params.description" :disabled="type==='view'"
                    style="width:60%;" maxlength="800" :placeholder="$t('resourcesManagement.roleDescTips')">
                    </el-input>
                </div>
            </div>
            <div class="form-row">
                <label class="required" for="">{{$t('userRole.codeSizeLimit')}}：</label>
                <div class="content">
                    <el-input v-model="params.codeSize" @input="params.codeSize = parseInt($event) || 0" :disabled="type==='view'" style="width:80px"></el-input> GB
                </div>
            </div>
            <div class="form-row">
                <label class="required" for="">{{$t('userRole.outputSizeLimit')}}：</label>
                <div class="content">
                    <el-input v-model="params.outputSize" @input="params.outputSize = parseInt($event) || 0" :disabled="type==='view'" style="width:80px"></el-input> GB
                </div>
            </div>
            <div class="btn-wrap" v-if="type!=='view'">
                <el-button type="primary" :disabled="type==='view'" class="btn confirm-btn" @click="confirm">{{ $t('confirm') }}</el-button>
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
  name: "containerStorageRoleDialog",
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
        type: 3,
        isCommon: 1,
        description: '',
        codeSize: 0,
        outputSize: 0,
      },
      commonTypeList: [{ k: 0, v: this.$t('userRole.default') }, { k: 1, v: this.$t('userRole.notDefault') }],
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
        case 'add': return this.$t('userRole.newContainerStorageRole')
        case 'edit': return this.$t('userRole.editContainerStorageRole')
        case 'view': return this.$t('userRole.viewContainerStorageRole')
      }
    },
  },
  methods: {
    initParams() {
      if (this.type === 'add') {
        this.resetParams()
      } else {
        this.restoreParams()
      }
    },
    restoreParams() {
      this.params.name = this.data.Name
      this.params.description = this.data.Description
      this.params.isCommon = this.data.IsCommon
      this.params.codeSize = this.data.codeSize || 0
      this.params.outputSize = this.data.outputSize || 0
    },
    resetParams() {
      this.params = { name: "", type: 3, isCommon: 1, description: '', codeSize: 0, outputSize: 0 }
    },
    open() {
      this.initParams()
      this.$emit("open");
    },
    opened() { this.$emit("opened"); },
    close() { this.$emit("close"); },
    closed() {
      this.$emit("closed");
      this.$emit("update:visible", false);
    },
    submitRole(params) {
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
            this.submitRole(this.params)
          }).catch(() => {})
        } else {
          this.submitRole(this.params)
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
  mounted() {},
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
      label {
        &.required::before {
          content: "*";
          color: red;
          margin-right: 5px;
        }
        width: 170px;
        color: rgba(136,136,136,1);
        font-size: 14px;
        box-sizing: border-box;
        text-align: right;
      }
      .content {
        flex: 1;
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
/deep/ .el-input.is-disabled .el-input__inner {
  color: #101010
}
</style>
