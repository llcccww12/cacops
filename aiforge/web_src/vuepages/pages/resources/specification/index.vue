<template>
  <div>
    <div class="title"><span>{{ $t('resourcesManagement.resSpecificationAndPriceManagement') }}</span></div>
    <div class="tools-bar">
      <div>
        <el-select class="select" size="medium" filterable v-model="selCenter" @change="selectChange">
          <el-option v-for="item in centerList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" filterable v-model="selQueue" @change="selectChange">
          <el-option v-for="item in queueList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selStatus" @change="selectChange">
          <el-option v-for="item in statusList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selAvailable" @change="selectChange">
          <el-option v-for="item in availableList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" filterable v-model="selResource" @change="selectChange">
          <el-option v-for="item in resourceList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" filterable v-model="selCardType" @change="selectChange">
          <el-option v-for="item in cardTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selCardsNum" @change="selectChange">
          <el-option v-for="item in cardsNumList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selNetworkType" @change="selectChange">
          <el-option v-for="item in networkTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selIsEnableVisualization" @change="selectChange">
          <el-option v-for="item in IsEnableVisualizationList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
      </div>
      <div class="tools-bar-btn-c">
        <el-button size="medium" icon="el-icon-refresh" @click="syncComputerNetwork" v-loading="syncLoading">
          {{ $t('resourcesManagement.syncAiNetwork') }}</el-button>
        <el-button type="primary" icon="el-icon-plus" size="medium" @click="showDialog('add')">
          {{ $t('resourcesManagement.addResSpecificationBtn') }}</el-button>
      </div>
    </div>
    <div class="table-container">
      <div style="min-height:600px;">
        <el-table border :data="tableData" v-tableSticky style="width:100%;" v-loading="loading" stripe>
          <el-table-column prop="ID" label="ID" align="center" header-align="center" width="60" fixed></el-table-column>
          <el-table-column prop="SpecStr" :label="$t('resourcesManagement.resourceSpecification')" align="left" fixed
            header-align="center" min-width="200">
          </el-table-column>
          <el-table-column prop="QueueInfo" :label="$t('resourcesManagement.resQueue')" align="left" fixed
            header-align="center" min-width="100">
          </el-table-column>
          <el-table-column prop="SourceSpecId" :label="$t('resourcesManagement.sourceSpecCode')" align="center"
            min-width="120" header-align="center">
          </el-table-column>
          <el-table-column prop="AccCardsNum" :label="$t('resourcesManagement.accCardsNum')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="CpuCores" :label="$t('resourcesManagement.cpuNum')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="GPUMemGiB" :label="`${$t('resourcesManagement.gpuMem')}(GB)`" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="MemGiB" :label="`${$t('resourcesManagement.mem')}(GB)`" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="ShareMemGiB" :label="`${$t('resourcesManagement.shareMem')}(GB)`" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="UpdatedTimeStr" :label="$t('resourcesManagement.lastUpdateTime')" align="center"
            min-width="140" header-align="center"></el-table-column>
          <el-table-column prop="UnitPrice"
            :label="`${$t('resourcesManagement.unitPrice')}(${$t('resourcesManagement.point_hr')})`" align="center"
            header-align="center">
            <template slot-scope="scope">
              <span style="font-weight:600;font-size:14px;">{{ scope.row.UnitPrice.toFixed(2) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="HasInternetStr" :label="$t('cloudbrainObj.networkType')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="EnableVisualizationStr" :label="$t('cloudbrainObj.visualization')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="IsAvailableStr" :label="$t('resourcesManagement.resourceSpecificationIsAvailable')"
            align="center" header-align="center" width="100">
            <template slot-scope="scope">
              <span :style="{ color: scope.row.IsAvailable ? 'rgb(82, 196, 26)' : 'rgb(245, 34, 45)' }">{{
                scope.row.IsAvailableStr
                }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="StatusStr" :label="$t('resourcesManagement.status')" align="center"
            header-align="center" width="100" fixed="right">
            <template slot-scope="scope">
              <span :style="{ color: scope.row.Status == '2' ? 'rgb(82, 196, 26)' : 'rgb(245, 34, 45)' }">{{
                scope.row.StatusStr
                }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('operation')" align="center" header-align="center" width="100" fixed="right">
            <template slot-scope="scope">
              <span v-if="scope.row.Status == '1' && !scope.row.UnitPrice">
                <span v-if="scope.row.IsAvailable" class="op-btn" @click="showDialog('edit', scope.row)">{{
                  $t('resourcesManagement.toSetPriceAndOnShelf')
                  }}</span>
                <span v-else class="op-btn" style="color: rgb(187, 187, 187); cursor: not-allowed;">{{
                  $t('resourcesManagement.toSetPriceAndOnShelf')
                  }}</span>
              </span>
              <span v-if="scope.row.Status == '2'">
                <span class="op-btn" @click="showDialog('edit', scope.row, true)">{{ $t('edit') }}</span>
                <span class="op-btn" @click="offShelfPrev(scope.row)">{{
                  $t('resourcesManagement.toOffShelf')
                  }}</span>
              </span>
              <span v-if="scope.row.Status == '3' || scope.row.Status == '1' && scope.row.UnitPrice">
                <span v-if="scope.row.IsAvailable" class="op-btn" @click="onShelf(scope.row)">{{
                  $t('resourcesManagement.toOnShelf')
                  }}</span>
                <span v-else class="op-btn" style="color: rgb(187, 187, 187); cursor: not-allowed;">{{
                  $t('resourcesManagement.toSetPriceAndOnShelf')
                  }}</span>
              </span>
            </template>
          </el-table-column>
          <template slot="empty">
            <span style="font-size: 12px">{{
              loading ? $t('loading') : $t('noData')
              }}</span>
          </template>
        </el-table>
      </div>
      <div class="__r_p_pagination">
        <div style="margin-top: 2rem">
          <div class="center">
            <el-pagination background @current-change="currentChange" @size-change="pageSizeChange"
              :current-page="pageInfo.curpage" :page-sizes="pageInfo.pageSizes" :page-size="pageInfo.pageSize"
              layout="total, sizes, prev, pager, next, jumper" :total="pageInfo.total">
            </el-pagination>
          </div>
        </div>
      </div>
    </div>
    <SpecificationDialog :visible.sync="specificationDialogShow" :type="specificationDialogType"
      :editOr="specificationDialogEditOr" :data="specificationDialogData" @confirm="specificationDialogConfirm">
    </SpecificationDialog>

    <BaseDialog :visible.sync="offShelfDialogShow" :width="`600px`" :title="$t('tips')">
      <div class="form">
        <div class="form-row" style="flex-direction:column;">
          <div class="content" style="margin:8px 0">{{ $t('resourcesManagement.offShelfDlgTip1') }}</div>
          <div class="content" style="margin:8px 0;font-weight: bold;">{{ offSelfDialogContent }}</div>
          <div class="content" style="margin:8px 0">{{ $t('resourcesManagement.offShelfDlgTip2') }}</div>
        </div>
        <div class="form-row" style="margin-top: 20px">
          <div class="content">
            <el-button type="primary" class="btn confirm-btn" @click="offShelf">{{ $t('confirm') }}</el-button>
            <el-button class="btn" @click="offShelfDialogShow = false">{{ $t('cancel') }}</el-button>
          </div>
        </div>
      </div>
    </BaseDialog>
  </div>
</template>

<script>
import SpecificationDialog from '../components/SpecificationDialog.vue';
import BaseDialog from '~/components/BaseDialog.vue';
import { getAiCenterCode, getResQueueCode, getResSpecificationList, updateResSpecification, syncResSpecification, getResSpecificationScenes } from '~/apis/modules/resources';
import { SPECIFICATION_STATUS, CLUSTERS, ACC_CARD_TYPE, COMPUTER_RESOURCES, NETWORK_TYPE, NETWORK_TYPE_VALUE } from '~/const';
import { getListValueWithKey } from '~/utils';
import { formatDate } from 'element-ui/lib/utils/date-util';

export default {
  data() {
    return {
      selCenter: '',
      centerList: [{ k: '', v: this.$t('resourcesManagement.allAiCenter') }],
      selQueue: '',
      queueList: [{ k: '', v: this.$t('resourcesManagement.allResQueue') }],
      selStatus: '',
      statusList: [{ k: '', v: this.$t('resourcesManagement.allStatus') }, ...SPECIFICATION_STATUS],
      selAvailable: '',
      availableList: [{ k: '', v: this.$t('resourcesManagement.resourceSpecificationIsAvailableAll') }, { k: '1', v: this.$t('resourcesManagement.available') }, { k: '2', v: this.$t('resourcesManagement.notAvailable') }],
      clusterList: [...CLUSTERS],
      accCardTypeList: [...ACC_CARD_TYPE],
      selResource: '',
      resourceList: [{ k: '', v: this.$t('resourcesManagement.allComputeResource') }, ...COMPUTER_RESOURCES],
      selCardType: '',
      cardTypeList: [{ k: '', v: this.$t('resourcesManagement.allAccCardType') }, ...ACC_CARD_TYPE],
      selCardsNum: -1,
      cardsNumList: [{ k: -1, v: this.$t('resourcesManagement.allCardsNum') }, ...[0, 1, 2, 4, 6, 8, 16, 32, 64].map(item => { return { k: item, v: `${this.$t('resourcesManagement.accCardsNum')}(${item})` } })],
      selNetworkType: 0,
      networkTypeList: [{ k: 0, v: this.$t('resourcesManagement.allNetworkType') }, ...NETWORK_TYPE],
      selIsEnableVisualization: 0,
      IsEnableVisualizationList: [{ k: 0, v: this.$t('resourcesManagement.allIsEnableVisualization') }, { k: 1, v: this.$t('resourcesManagement.notEnable') }, { k: 2, v: this.$t('resourcesManagement.enable') }],
      syncLoading: false,
      loading: false,
      tableData: [],
      pageInfo: {
        curpage: 1,
        pageSize: 10,
        pageSizes: [10, 50, 100],
        total: 0,
      },
      specificationDialogShow: false,
      specificationDialogType: 'add',
      specificationDialogEditOr: false,
      specificationDialogData: {},

      offShelfDialogShow: false,
      offShelfDialogData: {},
      offSelfDialogContent: '',
      comment: undefined
    };
  },
  components: { BaseDialog, SpecificationDialog },
  methods: {
    getCenterList() {
      getAiCenterCode().then(res => {
        res = res.data;
        if (res.Code === 0) {
          const data = res.Data;
          const list = [];
          for (let i = 0, iLen = data.length; i < iLen; i++) {
            const item = data[i];
            list.push({
              k: item.AiCenterCode,
              v: item.AiCenterName,
            });
          }
          this.centerList.push(...list);
        }
      }).catch(err => {
        console.log(err);
      });
    },
    getQueueList() {
      getResQueueCode().then(res => {
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
              v: `${item.QueueCode}${queueName}${queueType}(${getListValueWithKey(this.clusterList, item.Cluster)} - ${item.AiCenterName}) ${item.ComputeResource}(${item.AccCardType})`,
            });
          }
          this.queueList.push(...list);
        }
      }).catch(err => {
        console.log(err);
      });
    },
    getTableData() {
      const params = {
        center: this.selCenter,
        queue: this.selQueue,
        status: this.selStatus,
        available: this.selAvailable,
        resource: this.selResource,
        cardType: this.selCardType,
        cardsNum: this.selCardsNum,
        hasInternet: this.selNetworkType,
        enableVisualization: this.selIsEnableVisualization,
        page: this.pageInfo.curpage,
        pageSize: this.pageInfo.pageSize,
      };
      this.loading = true;
      getResSpecificationList(params).then(res => {
        this.loading = false;
        res = res.data;
        if (res.Code === 0) {
          const list = res.Data.List;
          const data = list.map((item) => {
            const Queue = item.Queue;
            const Spec = item.Spec;
            const NGPU = `${Queue.ComputeResource}:${Spec.AccCardsNum + '*' + getListValueWithKey(this.accCardTypeList, Queue.AccCardType)}`;
            const queueName = Queue.QueueName ? `【${Queue.QueueName}】` : '';
            const queueType = Queue.QueueType ? `【${Queue.QueueType}】` : '';
            return {
              ...Spec,
              SourceSpecId: Spec.SourceSpecId || '--',
              SpecStr: `${NGPU}(${this.$t('resourcesManagement.gpuMem')}:${Spec.GPUMemGiB}GB), CPU:${Spec.CpuCores}, ${this.$t('resourcesManagement.mem')}:${Spec.MemGiB}GB, ${this.$t('resourcesManagement.shareMem')}:${Spec.ShareMemGiB}GB`,
              QueueId: Queue.ID,
              QueueInfo: `${Queue.QueueCode}${queueName}${queueType}(${getListValueWithKey(this.clusterList, Queue.Cluster)} - ${Queue.AiCenterName})`,
              UpdatedTimeStr: formatDate(new Date(Spec.UpdatedTime * 1000), 'yyyy-MM-dd HH:mm:ss'),
              Status: Spec.Status.toString(),
              StatusStr: getListValueWithKey(this.statusList, Spec.Status.toString()),
              HasInternetStr: getListValueWithKey(NETWORK_TYPE_VALUE, Queue.HasInternet),
              EnableVisualizationStr: Queue.EnableVisualization ? this.$t('resourcesManagement.enable') : this.$t('resourcesManagement.notEnable'),
              IsAvailable: Spec.IsAvailable,
              IsAvailableStr: Spec.IsAvailable ? this.$t('resourcesManagement.available') : this.$t('resourcesManagement.notAvailable'),
            }
          });
          this.tableData = data;
          this.pageInfo.total = res.Data.TotalSize;
        }
      }).catch(err => {
        console.log(err);
        this.loading = false;
      });
    },
    syncComputerNetwork() {
      this.syncLoading = true;
      syncResSpecification().then(res => {
        this.syncLoading = false;
        res = res.data;
        if (res.Code === 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully')
          });
          this.getTableData();
        } else {
          this.$message({
            type: 'error',
            message: this.$t('submittedFailed')
          });
        }
      }).catch(err => {
        console.log(err);
        this.syncLoading = false;
        this.$message({
          type: 'error',
          message: this.$t('submittedFailed')
        });
      });
    },
    selectChange() {
      this.pageInfo.curpage = 1;
      this.getTableData();
    },
    currentChange(val) {
      this.pageInfo.curpage = val;
      this.getTableData();
    },
    pageSizeChange(val) {
      this.pageInfo.pageSize = val;
      this.getTableData();
    },
    showDialog(type, data, editOr) {
      this.specificationDialogType = type;
      this.specificationDialogEditOr = !!editOr;
      this.specificationDialogData = data ? { ...data } : {};
      this.specificationDialogShow = true;
    },
    specificationDialogConfirm() {
      this.specificationDialogShow = false;
      this.getTableData();
    },
    onShelf(data) {
      const type = 'on-shelf';
      this.$prompt(type === 'on-shelf' ? this.$t('resourcesManagement.onShelfConfirm') : this.$t('resourcesManagement.offShelfConfirm'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm1'),
        cancelButtonText: this.$t('cancel'),
        inputPlaceholder: this.$t('enterReason'),
        inputPattern: /.+/,
        inputValidator: (value) => {
          if (!value || value.trim() === '') {
            return this.$t('modelSquare.inputNotEmpty');
          }
          if (value.length > 255) {
            return this.$t('resourcesManagement.characterLengthPrompt');
          }
          return true;
        },
        inputErrorMessage: this.$t('enterReason'),
        type: 'warning',
        lockScroll: false,
        customClass: 'OnOff-shelf-prompt'
      }).then(({ value }) => {
        const comment = value.trim()
        updateResSpecification({
          ID: data.ID,
          action: type,
          comment
        }).then(res => {
          res = res.data;
          if (res.Code === 0) {
            this.$message({
              type: 'success',
              message: this.$t('submittedSuccessfully')
            });
            this.getTableData();
          } else {
            if (type === 'on-shelf' && res.Code === 1001) {
              this.$message({
                type: 'info',
                message: this.$t('resourcesManagement.onShelfCode1001')
              });
            } else if (type === 'on-shelf' && res.Code === 1003) {
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
        });
      }).catch(() => { });
    },
    offShelfPrev(data) {
      this.$prompt(this.$t('resourcesManagement.offShelfConfirm'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm1'),
        cancelButtonText: this.$t('cancel'),
        inputPlaceholder: this.$t('enterReason'),
        inputPattern: /.+/,
        inputValidator: (value) => {
          if (!value || value.trim() === '') {
            return this.$t('modelSquare.inputNotEmpty');
          }
          if (value.length > 255) {
            return this.$t('resourcesManagement.characterLengthPrompt');
          }
          return true;
        },
        inputErrorMessage: this.$t('enterReason'),
        lockScroll: false,
        type: 'warning',
        customClass: 'OnOff-shelf-prompt'
      }).then(({ value }) => {
        this.offShelfDialogData = data;
        getResSpecificationScenes({ ID: data.ID }).then(res => {
          res = res.data;
          this.comment = value.trim()
          if (res.Code === 0) {
            if (res.Data.RoleName.length) {
              this.offShelfDialogShow = true;
              this.offSelfDialogContent = res.Data.RoleName.join(', ');
            } else {
              // const comment = value.trim()
              this.offShelf();
            }
          } else {
            console.log(res);
          }
        }).catch(err => {
          console.log(err);
        });
      }).catch(() => { });
    },
    offShelf() {
      updateResSpecification({
        ID: this.offShelfDialogData.ID,
        action: 'off-shelf',
        comment: this.comment
      }).then(res => {
        res = res.data;
        if (res.Code === 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully')
          });
          this.offShelfDialogShow = false;
          this.getTableData();
        } else {
          this.$message({
            type: 'error',
            message: this.$t('submittedFailed')
          });
          this.comment = undefined
        }
      }).catch(err => {
        console.log(err);
        this.comment = undefined
        this.$message({
          type: 'error',
          message: this.$t('submittedFailed')
        });
      });
    },
  },
  mounted: function () {
    this.getQueueList();
    this.getCenterList();
    this.getTableData();
  },
  beforeDestroy: function () {
  },
};
</script>

<style lang="less">
.OnOff-shelf-prompt {
  .el-input.el-input--small {
    position: relative;

    &::before {
      content: "*";
      color: red;
      position: absolute;
      left: -10px;
      top: 50%;
      transform: translateY(-50%);
    }
  }
}
</style>

<style scoped lang="less">
.title {
  height: 30px;
  display: flex;
  align-items: center;
  margin-bottom: 5px;

  span {
    font-weight: 700;
    font-size: 16px;
    color: rgb(16, 16, 16);
  }
}

.tools-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;

  .select {
    margin-right: 10px;
    margin-bottom: 10px;

    /deep/ .el-input__inner {
      border-radius: 0;
    }
  }

  .tools-bar-btn-c {
    display: flex;

    .el-button {
      margin-bottom: 10px;
    }
  }
}

.table-container {
  margin-bottom: 16px;

  /deep/ .el-table__header {
    th {
      background: rgb(245, 245, 246);
      font-size: 12px;
      color: rgb(36, 36, 36);
    }
  }

  /deep/ .el-table__body {
    td {
      font-size: 12px;
    }
  }

  .op-btn {
    cursor: pointer;
    font-size: 12px;
    color: rgb(25, 103, 252);
    margin-right: 4px;
  }
}

.center {
  display: flex;
  justify-content: center;
}

.form {
  margin: 5px 0 5px 0;
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
}

.form-row {
  display: flex;
  min-height: 42px;
  margin-bottom: 4px;

  .content {
    width: 500px;
    display: flex;
    align-items: center;
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
</style>
