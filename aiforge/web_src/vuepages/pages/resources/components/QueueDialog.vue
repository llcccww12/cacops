<template>
  <div class="base-dlg">
    <BaseDialog :visible.sync="dialogShow" :width="`750px`"
      :title="type === 'add' ? $t('resourcesManagement.addResQueue') : $t('resourcesManagement.editResQueue')"
      @open="open" @opened="opened" @close="close" @closed="closed">
      <div class="dlg-content">
        <div class="form">
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.resQueueCode') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.QueueCode" placeholder="" :disabled="type === 'edit'" maxlength="255">
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title">
              <span>{{ $t('resourcesManagement.resQueueName') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.QueueName" placeholder="" maxlength="255">
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title">
              <span>{{ $t('resourcesManagement.resQueueType') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.QueueType" clearable>
                <el-option v-for="item in queueTypeList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.whichCluster') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.Cluster" :disabled="type === 'edit'" @change="changeCluster">
                <el-option v-for="item in clusterList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.aiCenter') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.AiCenterCode" :disabled="type === 'edit'">
                <el-option v-for="item in computingCenterList[clusterValue]" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.computeResource') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.ComputeResource" :disabled="type === 'edit'">
                <el-option v-for="item in computingTypeList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.accCardType') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.AccCardType" :disabled="type === 'edit'">
                <el-option v-for="item in cardTypeList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.cardsTotalNum') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.CardsTotalNum" type="number" placeholder=""></el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('cloudbrainObj.networkType') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.HasInternet">
                <el-option v-for="item in networkTypeList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.resQueueIsAvailable') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.IsAvailable">
                <el-option v-for="item in isAvailableList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row" style="margin-top: 10px">
            <div class="title"><span>{{ $t('resourcesManagement.remark') }}</span></div>
            <div class="content" style="width: 400px">
              <el-input type="textarea" :autosize="{ minRows: 3, maxRows: 4 }" maxlength="2000"
                :placeholder="$t('resourcesManagement.pleaseEnterRemark')" v-model="dataInfo.Remark">
              </el-input>
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
import { addResQueue, updateResQueue } from '~/apis/modules/resources';
import { CLUSTERS, AI_CENTER, COMPUTER_RESOURCES, ACC_CARD_TYPE, NETWORK_TYPE_VALUE } from '~/const';

export default {
  name: "QueueDialog",
  props: {
    visible: { type: Boolean, default: false },
    title: { type: String, default: '' },
    type: { type: String, defalut: 'add' },
    data: { type: Object, default: () => ({}) },
  },
  components: {
    BaseDialog
  },
  data() {
    return {
      dialogShow: false,
      queueTypeList: [{ k: 'public', v: 'public' }, { k: 'exclusive', v: 'exclusive' }],
      clusterList: [CLUSTERS[0], CLUSTERS[2]],
      clusterValue:'OpenI',
      computingCenterList: {
        OpenI: [AI_CENTER[0], AI_CENTER[1], AI_CENTER[2]],
        IFLYTEKTraining: [AI_CENTER[6]]
      },
      computingTypeList: [...COMPUTER_RESOURCES],
      cardTypeList: [...ACC_CARD_TYPE],
      networkTypeList: [...NETWORK_TYPE_VALUE],
      isAvailableList: [{ k: 1, v: this.$t('resourcesManagement.notAvailable') }, { k: 2, v: this.$t('resourcesManagement.available') }],
      dataInfo: {Cluster: 'OpenI'},
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
        ID: '',
        QueueCode: '',
        QueueName: '',
        QueueType: '',
        Cluster: '',
        AiCenterCode: '',
        ComputeResource: '',
        AccCardType: '',
        CardsTotalNum: '',
        HasInternet: '',
        IsAvailable: '',
        Remark: '',
      }
    },
    changeCluster(val) {
      this.clusterValue = val
      this.dataInfo.AiCenterCode = ''
    },
    open() {
      this.resetDataInfo();
      if (this.type === 'add') {
        //
      } else if (this.type === 'edit') {
        this.dataInfo = Object.assign(this.dataInfo, { ...this.data });
        this.dataInfo.IsAvailable = this.data.IsAvailable ? 2 : 1;
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
      if (!this.dataInfo.QueueCode || !this.dataInfo.Cluster || !this.dataInfo.AiCenterCode || !this.dataInfo.ComputeResource
        || !this.dataInfo.AccCardType || !this.dataInfo.CardsTotalNum || this.dataInfo.HasInternet === '' || this.dataInfo.IsAvailable === '') {
        this.$message({
          type: 'info',
          message: this.$t('pleaseCompleteTheInformationFirst'),
        });
        return;
      }
      if (parseInt(this.dataInfo.CardsTotalNum) != Number(this.dataInfo.CardsTotalNum)) {
        this.$message({
          type: 'info',
          message: this.$t('pleaseEnterPositiveIntegerCardsTotalNum')
        });
        return;
      }
      const setApi = this.type === 'add' ? addResQueue : updateResQueue;
      setApi({ ...this.dataInfo, CardsTotalNum: Number(this.dataInfo.CardsTotalNum) }).then(res => {
        res = res.data;
        if (res.Code === 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully')
          });
          this.$emit("confirm");
        } else {
          this.$message({
            type: 'error',
            message: this.$t('submittedFailed')
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
