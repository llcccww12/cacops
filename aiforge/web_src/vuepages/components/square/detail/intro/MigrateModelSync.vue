<template>
  <div>
    <a class="operate-btn" @click.stop.prevent="dlgShow = true">
      <i class="ri-restart-line"></i>{{ $t('modelObj.sync_now') }}
    </a>
    <BaseDialog class="export-model-dlg base-dlg" :visible.sync="dlgShow"
      :title="`${data.name} ${$t('modelObj.model_sync')}`" width="950px" :modal="true" :modalAppendToBody="true"
      :appendToBody="true" :close-on-click-modal="false" :show-close="true" :lockScroll="true" :destroy-on-close="false"
      @open="open" @closed="closed">
      <div class="dlg-content" v-loading="loading || submitLoading">
        <div class="compare-files" v-if="!loading && hasUpdate">
          <div class="table-container">
            <el-table ref="tableRef" :data="filesList" :height="400" style="width:100%;height:100%;" stripe>
              <el-table-column prop="FileName" :label="$t('modelManage.fileName')" align="left" header-align="left">
                <template slot-scope="scope">
                  <div class="fitted file-name">
                    <i class="icon file alternate outline" width="16" height="16" aria-hidden="true"></i>
                    <span :title="scope.row.FileName" :class="scope.row.Status == 1 ? 'color-add' :
                      scope.row.Status == 2 ? 'color-del' :
                        scope.row.Status == 3 ? 'color-update' : ''">{{ scope.row.FileName }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="SizeShow" :label="$t('modelObj.current_size')" align="center" header-align="center"
                width="200">
              </el-table-column>
              <el-table-column prop="NewSizeShow" :label="$t('modelObj.new_size')" align="center" header-align="center"
                width="200">
                <template slot-scope="scope">
                  <span :class="scope.row.Status == 1 ? 'color-add' :
                    scope.row.Status == 2 ? 'color-del' :
                      scope.row.Status == 3 ? 'color-update' : ''">{{ scope.row.NewSizeShow }}</span>
                </template>
              </el-table-column>
              <el-table-column prop="Status" :label="$t('status')" align="center" header-align="center" width="200">
                <template slot-scope="scope">
                  <span v-if="scope.row.Status == 0"></span>
                  <span v-if="scope.row.Status == 1" class="color-add">{{ $t('modelObj.status_add') }}</span>
                  <span v-if="scope.row.Status == 2" class="color-del">{{ $t('modelObj.status_del') }}</span>
                  <span v-if="scope.row.Status == 3" class="color-update">{{ $t('modelObj.status_update') }}</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
          <div class="op-btns-c">
            <div>
              <div class="hf-token-c" v-if="canUpdate && gated">
                <span class="title-requred">{{ $t('modelManage.externalToken') }}</span>
                <el-input v-model="hf_token" :placeholder="$t('modelManage.modelLabelGatedTips')" />
              </div>
            </div>
            <div class="btns-c">
              <el-button type="primary" v-if="canUpdate" size="default" class="submit-btn" @click="update">{{
                $t('modelObj.start_sync')
                }}</el-button>
              <el-button class="cancel-btn" size="default" @click="cancel">{{ $t('cancel') }}</el-button>
            </div>
          </div>
        </div>
        <div class="no-update" v-if="!loading && !hasUpdate">
          <div class="content">
            <div class="icon-c"><i class="ri-information-line"></i></div>
            <div class="txt">{{ $t('modelObj.model_no_update_tips') }}</div>
          </div>
          <div class="op-btns-c">
            <div></div>
            <el-button class="cancel-btn" size="default" @click="cancel">{{ $t('cloudbrainObj.dialogTips.tips8')
              }}</el-button>
          </div>
        </div>
      </div>
    </BaseDialog>
  </div>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';
import { getModelMigrateUpdateInfo, setModelMigrateUpdate } from '~/apis/modules/modelmanage';
import { transFileSize } from '~/utils';

export default {
  name: 'MigrateModelSync',
  props: {
    data: { type: Object, default: () => ({}) },
  },
  components: { BaseDialog, },
  data() {
    return {
      dlgShow: false,
      loading: true,
      submitLoading: false,
      hasUpdate: false,
      canUpdate: false,
      filesList: [],
      mainData: {},
      gated: false,
      hf_token: '',
    };
  },
  methods: {
    open() {
      this.loading = true;
      this.submitLoading = false;
      this.hf_token = '';
      this.gated = false;
      this.hasUpdate = true;
      this.canUpdate = false;
      this.filesList = [];
      getModelMigrateUpdateInfo({
        hf_repo_id: this.data.external_name,
      }).then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 200) {
          const data = res.data;
          this.mainData = data;
          this.hasUpdate = data.is_update;
          this.canUpdate = this.hasUpdate;
          this.gated = data.gated;
          this.filesList = (data.files || []).map(item => {
            return {
              ...item,
              FileName: item.filename,
              SizeShow: item.update_status == 1 ? '--' : transFileSize(item.size),
              NewSizeShow: item.update_status == 2 ? '--' : transFileSize(item.size_new),
              Status: item.update_status,
            }
          });
        } else {
          this.$message.error(res.msg || this.$t('operationFailed'));
        }
      }).catch(err => {
        this.loading = false;
        if (err.response.status == 401) {
          window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
        } else {
          this.$message.error(err.response.data.message || this.$t('operationFailed'));
        }
      });
    },
    closed() { },
    update() {
      if (this.submitLoading) return;
      if (this.canUpdate && this.gated && this.hf_token.trim() == '') {
        this.$message.info(this.$t('modelManage.modelLabelGatedTips'));
        return;
      }
      this.submitLoading = true;
      setModelMigrateUpdate({
        ...this.mainData,
        hf_token: this.hf_token.trim(),
      }).then(res => {
        const data = res.data;
        if (data.code == 1) {
          window.location.reload();
        } else {
          this.submitLoading = false;
          this.$message.error(data.msg || this.$t('operationFailed'));
        }
      }).catch(err => {
        this.submitLoading = false;
        this.$message.error(err.response.data.message || this.$t('operationFailed'));
      });
    },
    cancel() {
      this.dlgShow = false;
    },
  },
  beforeMount() { },
  mounted() { }
};
</script>

<style scoped lang="less">
.operate-btn {
  display: flex;
  align-items: center;

  i {
    margin-right: 2px;
  }
}

.dlg-content {
  padding: 10px;
  min-height: 300px;

  .table-container {
    min-height: 300px;
    max-height: 400px;

    /deep/ .el-table__header {
      th {
        background: rgb(245, 245, 246);
        color: rgb(16, 16, 16);
        font-weight: 400;
        font-size: 14px;
      }
    }

    /deep/ .el-table__body {
      td {
        color: rgb(16, 16, 16);
        font-weight: 400;
        font-size: 14px;
      }
    }
  }

  .no-update {
    .content {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      height: 300px;
      color: rgba(56, 158, 13, 1);
      margin-bottom: 30px;

      .icon-c {
        margin-bottom: 10px;

        i {
          font-size: 48px;
        }
      }

      .txt {
        font-size: 14px;
      }
    }
  }

  .op-btns-c {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-top: 20px;
    padding-right: 10px;
    padding-bottom: 10px;

    .hf-token-c {
      display: flex;
      align-items: center;

      .title-requred {
        position: relative;
        margin-right: 20px;

        &::after {
          position: absolute;
          margin: -0.2em 0 0 0.2em;
          content: '*';
          color: #db2828;
        }
      }

      .el-input {
        width: 340px;
      }
    }

    .submit-btn {
      background-color: #5bb973;
      color: #fff;
      border-color: #5bb973;

      &:hover {
        background-color: #16ab39;
      }

      &:active {
        background-color: #198f35;
      }

      &.is-disabled,
      &.is-disabled:active,
      &.is-disabled:focus,
      &.is-disabled:hover {
        background-color: #5bb973;
        opacity: .45;
      }

      a {
        color: #fff;
        font-weight: 700;
      }
    }

    .cancel-btn {
      background: #e0e1e2;
      color: rgba(0, 0, 0, .6);
      border-color: #e0e1e2;

      &:hover {
        background-color: #cacbcd;
      }

      &:active {
        background-color: #babbbc;
      }
    }
  }

  .color-update {
    color: rgb(255, 169, 64);
  }

  .color-add {
    color: rgb(56, 158, 13);
  }

  .color-del {
    color: rgb(245, 34, 45);
  }
}
</style>
