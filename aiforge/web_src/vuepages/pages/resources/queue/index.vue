<template>
  <div>
    <div class="title"><span>{{ $t('resourcesManagement.resQueue') }}</span></div>
    <div class="tools-bar">
      <div class="left">
        <el-select class="select" size="medium" v-model="selCluster" @change="selectChange">
          <el-option v-for="item in clusterList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selQueueType" @change="selectChange">
          <el-option v-for="item in queueTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selQueueIsExclusiveType" @change="selectChange">
          <el-option v-for="item in queueIsExclusiveTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" filterable v-model="selComputingCenter" @change="selectChange">
          <el-option v-for="item in computingCenterList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" filterable v-model="selComputingType" @change="selectChange">
          <el-option v-for="item in computingTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" filterable v-model="selCardType" @change="selectChange">
          <el-option v-for="item in cardTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selNetworkType" @change="selectChange">
          <el-option v-for="item in networkTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selIsAvailable" @change="selectChange">
          <el-option v-for="item in isAvailableList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selIsEnableVisualization" @change="selectChange">
          <el-option v-for="item in IsEnableVisualizationList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
      </div>
      <div class="tools-bar-btn-c">
        <el-button size="medium" icon="el-icon-refresh" @click="syncComputerNetwork" v-loading="syncLoading">
          {{ $t('resourcesManagement.syncAiNetwork') }}</el-button>
        <el-button type="primary" icon="el-icon-plus" size="medium" @click="showDialog('add')">
          {{ $t('resourcesManagement.addResQueueBtn') }}</el-button>
      </div>
    </div>
    <div class="table-container">
      <div style="min-height:600px;">
        <el-table border :data="tableData" v-tableSticky style="width: 100%" v-loading="loading" stripe>
          <el-table-column prop="ID" label="ID" align="center" header-align="center" width="80"></el-table-column>
          <el-table-column prop="QueueCode" :label="$t('resourcesManagement.resQueueName')" align="center"
            header-align="center">
            <template slot-scope="scope">
              <span :title="scope.row.Cluster">{{ `${scope.row.QueueCode}${scope.row.QueueName ?
                '【' + scope.row.QueueName + '】' : ''}` }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="QueueType" :label="$t('resourcesManagement.resQueueType')" align="center"
            header-align="center" width="130">
          </el-table-column>
          <el-table-column prop="SceneTypeStr" :label="$t('resourcesManagement.sceneType')" align="center"
            header-align="center" width="120">
            <template slot-scope="scope">
              <span :style="{ color: scope.row.IsQueueExclusive ? 'red' : '' }">
                {{ scope.row.IsQueueExclusiveStr }}
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="ClusterName" :label="$t('resourcesManagement.whichCluster')" align="center"
            header-align="center">
            <template slot-scope="scope">
              <span :title="scope.row.Cluster">{{ scope.row.ClusterName }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="AiCenterCode" :label="$t('resourcesManagement.aiCenterID')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="AiCenterName" :label="$t('resourcesManagement.aiCenter')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="ComputeResourceName" :label="$t('resourcesManagement.computeResource')" align="center"
            header-align="center">
          </el-table-column>
          <el-table-column prop="AccCardTypeName" :label="$t('resourcesManagement.accCardType')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="CardsTotalNum" :label="$t('resourcesManagement.cardsTotalNum')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="HasInternetStr" :label="$t('cloudbrainObj.networkType')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="EnableVisualizationStr" :label="$t('cloudbrainObj.visualization')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="IsAvailableStr" :label="$t('resourcesManagement.resQueueIsAvailable')" align="center"
            header-align="center">
            <template slot-scope="scope">
              <span :style="{ color: scope.row.IsAvailable ? 'rgb(82, 196, 26)' : 'rgb(245, 34, 45)' }">{{
                scope.row.IsAvailableStr
                }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="UpdatedTimeStr" :label="$t('resourcesManagement.lastUpdateTime')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="Remark" :label="$t('resourcesManagement.remark')" align="left" header-align="center"
            min-width="160">
          </el-table-column>
          <el-table-column :label="$t('operation')" align="center" header-align="center" width="80" fixed="right">
            <template slot-scope="scope">
              <span v-if="scope.row.Cluster !== 'C2Net'" class="op-btn" @click="showDialog('edit', scope.row)">{{
                $t('edit')
                }}</span>
              <span v-else class="op-btn" style="color:rgb(187, 187, 187);cursor:not-allowed">{{
                $t('edit')
                }}</span>
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
    <QueueDialog :visible.sync="queueDialogShow" :type="queueDialogType" :data="queueDialogData"
      @confirm="queueDialogConfirm"></QueueDialog>
  </div>
</template>

<script>
import QueueDialog from '../components/QueueDialog.vue';
import { getAiCenterList, getResQueueList, syncResQueue } from '~/apis/modules/resources';
import { CLUSTERS, COMPUTER_RESOURCES, ACC_CARD_TYPE, NETWORK_TYPE, NETWORK_TYPE_VALUE } from '~/const';
import { getListValueWithKey } from '~/utils';
import { formatDate } from 'element-ui/lib/utils/date-util';

export default {
  data() {
    return {
      selCluster: '',
      clusterList: [{ k: '', v: this.$t('resourcesManagement.allCluster') }, ...CLUSTERS],
      selQueueType: '',
      queueTypeList: [{ k: '', v: this.$t('resourcesManagement.allResQueueType') }, { k: 'public', v: 'public' }, { k: 'exclusive', v: 'exclusive' }],
      selQueueIsExclusiveType: '',
      queueIsExclusiveTypeList: [{ k: '', v: this.$t('resourcesManagement.allSceneType') }, { k: 'public', v: this.$t('resourcesManagement.public') }, { k: 'exclusive', v: this.$t('resourcesManagement.exclusive') }],
      selComputingCenter: '',
      computingCenterList: [{ k: '', v: this.$t('resourcesManagement.allAiCenter') }],
      selComputingType: '',
      computingTypeList: [{ k: '', v: this.$t('resourcesManagement.allComputeResource') }, ...COMPUTER_RESOURCES],
      selCardType: '',
      cardTypeList: [{ k: '', v: this.$t('resourcesManagement.allAccCardType') }, ...ACC_CARD_TYPE],
      selNetworkType: 0,
      networkTypeList: [{ k: 0, v: this.$t('resourcesManagement.allNetworkType') }, ...NETWORK_TYPE],
      selIsAvailable: 0,
      isAvailableList: [{ k: 0, v: this.$t('resourcesManagement.resQueueIsAvailableAll') }, { k: 1, v: this.$t('resourcesManagement.notAvailable') }, { k: 2, v: this.$t('resourcesManagement.available') }],
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
      queueDialogShow: false,
      queueDialogType: 'add',
      queueDialogData: {},
    };
  },
  components: { QueueDialog },
  methods: {
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
          this.computingCenterList.splice(1, Infinity, ...data);
        }
      }).catch(err => {
        console.log(err);
      });
    },
    getTableData() {
      const params = {
        cluster: this.selCluster,
        queueType: this.selQueueType,
        isQueueExclusive: this.selQueueIsExclusiveType ? this.selQueueIsExclusiveType == 'exclusive' ? 2 : 1 : '',
        center: this.selComputingCenter,
        resource: this.selComputingType,
        card: this.selCardType,
        hasInternet: this.selNetworkType,
        isAvailable: this.selIsAvailable,
        enableVisualization: this.selIsEnableVisualization,
        page: this.pageInfo.curpage,
        pageSize: this.pageInfo.pageSize,
      };
      this.loading = true;
      getResQueueList(params).then(res => {
        this.loading = false;
        res = res.data;
        if (res.Code === 0) {
          const list = res.Data.List;
          const data = list.map((item) => {
            return {
              ...item,
              QueueCode: item.QueueCode || '--',
              IsQueueExclusiveStr: getListValueWithKey(this.queueIsExclusiveTypeList.slice(1, Infinity), item.IsQueueExclusive ? 'exclusive' : 'public'),
              ClusterName: getListValueWithKey(this.clusterList, item.Cluster),
              ComputeResourceName: getListValueWithKey(this.computingTypeList, item.ComputeResource),
              AccCardTypeName: getListValueWithKey(this.cardTypeList, item.AccCardType),
              HasInternetStr: getListValueWithKey(NETWORK_TYPE_VALUE, item.HasInternet),
              EnableVisualizationStr: item.EnableVisualization ? this.$t('resourcesManagement.enable') : this.$t('resourcesManagement.notEnable'),
              IsAvailableStr: item.IsAvailable ? this.$t('resourcesManagement.available') : this.$t('resourcesManagement.notAvailable'),
              UpdatedTimeStr: formatDate(new Date(item.UpdatedTime * 1000), 'yyyy-MM-dd HH:mm:ss'),
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
      syncResQueue().then(res => {
        this.syncLoading = false;
        res = res.data;
        if (res.Code === 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully')
          });
          this.getAiCenterList();
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
    showDialog(type, data) {
      this.queueDialogType = type;
      this.queueDialogData = data ? { ...data } : {};
      this.queueDialogShow = true;
    },
    queueDialogConfirm() {
      this.queueDialogShow = false;
      this.getAiCenterList();
      this.getTableData();
    }
  },
  mounted() {
    this.getAiCenterList();
    this.getTableData();
  },
  beforeDestroy() {
  },
};
</script>

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
  align-items: flex-start;
  ;
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
    margin: 0 5px;
  }
}

.center {
  display: flex;
  justify-content: center;
}
</style>
