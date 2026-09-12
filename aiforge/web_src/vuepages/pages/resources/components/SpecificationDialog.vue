<template>
  <div class="base-dlg">
    <BaseDialog :visible.sync="dialogShow" :width="`700px`"
      :title="type === 'add' ? $t('resourcesManagement.addResSpecificationAndPriceInfo') : $t('resourcesManagement.editResSpecificationAndPriceInfo')"
      @open="open" @opened="opened" @close="close" @closed="closed">
      <div class="dlg-content">
        <div class="form">
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.resQueue') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.QueueId" :disabled="type === 'edit'">
                <el-option v-for="item in this.queueList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row">
            <div class="title">
              <span>{{ $t('resourcesManagement.sourceSpecCode') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.SourceSpecId" :placeholder="$t('resourcesManagement.sourceSpecCodeTips')"
                maxlength="255" :disabled="type === 'edit'">
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.accCardsNum') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.AccCardsNum" type="number" placeholder="" :disabled="type === 'edit'">
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.cpuNum') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.CpuCores" type="number" placeholder="" :disabled="type === 'edit'"></el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.gpuMem') }}(GB)</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.GPUMemGiB" type="number" placeholder="" :disabled="type === 'edit'">
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.mem') }}(GB)</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.MemGiB" type="number" placeholder="" :disabled="type === 'edit'"></el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.shareMem') }}(GB)</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.ShareMemGiB" type="number" placeholder="" :disabled="type === 'edit'">
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.unitPrice') }}({{ $t('resourcesManagement.point_hr') }})</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.UnitPrice" type="number" @change="checkUnitPrice" placeholder=""></el-input>
            </div>
          </div>
          <div class="form-row" v-if="type != 'add'">
            <div class="title required">
              <span>{{ $t('resourcesManagement.reason') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.Comment" :placeholder="$t('enterReason')"
                maxlength="255">
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('status') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.Status" :disabled="type === 'edit'">
                <el-option v-for="item in this.statusList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row" style="margin-top: 20px">
            <div class="title"></div>
            <div class="content">
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
import { getResQueueCode, addResSpecification, updateResSpecification, getAiCenterList } from '~/apis/modules/resources';
import { SPECIFICATION_STATUS, CLUSTERS } from '~/const';
import { getListValueWithKey } from '~/utils';

export default {
  name: "SpecificationDialog",
  props: {
    visible: { type: Boolean, default: false },
    title: { type: String, default: '' },
    type: { type: String, defalut: 'add' },
    editOr: { type: Boolean, defalut: false },
    data: { type: Object, default: () => ({}) },
  },
  components: {
    BaseDialog
  },
  data() {
    return {
      dialogShow: false,
      dataInfo: {},
      queueList: [],
      statusList: [...SPECIFICATION_STATUS],
      clusterList: [...CLUSTERS],
      aiCenterList: [],
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
        QueueId: '',
        AccCardsNum: '',
        CpuCores: '',
        MemGiB: '',
        ShareMemGiB: '',
        GPUMemGiB: '',
        UnitPrice: '',
        Comment: '',
        Status: '1',
      }
    },
    getAiCenterList() {
      getAiCenterList().then(res => {
        res = res.data;
        if (res.Code === 0) {
          const list = res.Data;
          const data = list.map(item => {
            return {
              k: item.AiCenterCode,
              v: item.AiCenterName
            };
          });
          this.aiCenterList.splice(0, Infinity, ...data);
        }
      }).catch(err => {
        console.log(err);
      });
    },
    getQueueList() {
      getResQueueCode({ cluster: this.type === 'add' ? 'OpenI||IFLYTEKTraining' : undefined}).then(res => {
        res = res.data;
        if (res.Code === 0) {
          const data = res.Data;
          const list = [];
          for (let i = 0, iLen = data.length; i < iLen; i++) {
            const item = data[i];
            const queueName = item.QueueName ? `【${item.QueueName}】` : '';
            const queueType = item.QueueType ? `【${item.QueueType}】` : '';
            list.push({
              k: item.ID,
              v: `${item.QueueCode}${queueName}${queueType}(${getListValueWithKey(this.clusterList, item.Cluster)} - ${item.AiCenterName})`,
            });
          }
          this.queueList.splice(0, Infinity, ...list);
        }
      }).catch(err => {
        console.log(err);
      });
    },
    checkUnitPrice() {
      if (this.dataInfo.UnitPrice) {
        this.dataInfo.UnitPrice = Math.abs(Number(this.dataInfo.UnitPrice)).toFixed(2);
      } else {
        this.dataInfo.UnitPrice = '';
      }
    },
    open() {
      this.resetDataInfo();
      this.getQueueList();
      this.getAiCenterList();
      if (this.type === 'add') {
        //
      } else if (this.type === 'edit') {
        this.dataInfo = Object.assign(this.dataInfo, { ...this.data, Status: '2' });
      }
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
    confirm() {
      const isAdd = this.type === 'add';
      const commonCondition = this.dataInfo.AccCardsNum === '' 
        || this.dataInfo.CpuCores === '' 
        || this.dataInfo.MemGiB === '' 
        || this.dataInfo.ShareMemGiB === '' 
        || this.dataInfo.GPUMemGiB === ''
        || this.dataInfo.UnitPrice === '' 
        || !this.dataInfo.Status;

      const shouldShowMessage = isAdd 
        ? commonCondition 
        : (commonCondition || !this.dataInfo.Comment);

      if (shouldShowMessage) {
        this.$message({
          type: 'info',
          message: this.$t('pleaseCompleteTheInformationFirst')
        });
        return;
      }
      if (parseInt(this.dataInfo.AccCardsNum) != Number(this.dataInfo.AccCardsNum)) {
        this.$message({
          type: 'info',
          message: this.$t('pleaseEnterPositiveIntegerCardsTotalNum')
        });
        return;
      }
      const setApi = this.type === 'add' ? addResSpecification : updateResSpecification;
      const action = this.editOr ? 'edit' : this.type === 'edit' ? 'on-shelf' : undefined;
      setApi({
        ...this.dataInfo,
        comment: this.dataInfo.Comment,
        action: action,
        AccCardsNum: Number(this.dataInfo.AccCardsNum),
        CpuCores: Number(this.dataInfo.CpuCores),
        MemGiB: Number(this.dataInfo.MemGiB),
        ShareMemGiB: Number(this.dataInfo.ShareMemGiB),
        GPUMemGiB: Number(this.dataInfo.GPUMemGiB),
        UnitPrice: Number(this.dataInfo.UnitPrice),
        Status: Number(this.dataInfo.Status)
      }).then(res => {
        res = res.data;
        if (res.Code === 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully')
          });
          this.$emit("confirm");
        } else {
          if (action === 'on-shelf' && res.Code === 1001) {
            this.$message({
              type: 'info',
              message: this.$t('resourcesManagement.onShelfCode1001')
            });
          } else if (action === 'on-shelf' && res.Code === 1003) {
            this.$message({
              type: 'info',
              message: this.$t('resourcesManagement.onShelfCode1003')
            });
          } else {
            this.$message({
              type: 'error',
              message: this.$t('submittedFailed')
            });
          }
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
  margin: 20px 0 25px 0;
  display: flex;
  justify-content: center;

  .form {
    width: 600px;

    .form-row {
      display: flex;
      min-height: 42px;
      margin-bottom: 4px;

      .title {
        width: 160px;
        display: flex;
        justify-content: flex-end;
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
        width: 300px;
        display: flex;
        align-items: center;

        /deep/ .el-select {
          width: 100%;
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
