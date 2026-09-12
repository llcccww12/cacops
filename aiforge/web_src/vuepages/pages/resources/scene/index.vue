<template>
  <div>
    <div class="title"><span>{{ $t('resourcesManagement.resSceneManagement') }}</span></div>
    <div class="tools-bar">
      <div>
        <el-select class="select" size="medium" v-model="selTaskType" @change="selectChange">
          <el-option v-for="item in taskTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selSceneType" @change="selectChange">
          <el-option v-for="item in sceneTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selIsSpecExclusive" @change="selectChange">
          <el-option v-for="item in IsSpecExclusiveList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selCluster" @change="selectChange">
          <el-option v-for="item in clusterList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selAiCenter" @change="selectChange">
          <el-option v-for="item in aiCenterList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selQueue" @change="selectChange">
          <el-option v-for="item in queueList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selResource" @change="selectChange">
          <el-option v-for="item in resourceList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selCardType" @change="selectChange">
          <el-option v-for="item in cardTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selNetworkType" @change="selectChange">
          <el-option v-for="item in networkTypeList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="selIsEnableVisualization" @change="selectChange">
          <el-option v-for="item in IsEnableVisualizationList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
      </div>
      <div>
        <el-button type="primary" icon="el-icon-plus" size="medium" @click="showDialog('add')">
          {{ $t('resourcesManagement.addResSceneBtn') }}</el-button>
      </div>
    </div>
    <div class="table-container">
      <div style="min-height:600px;">
        <el-table border :data="tableData" v-tableSticky style="width:100%" v-loading="loading" stripe :span-method="tableSpanMethod">
          <el-table-column prop="_id_" label="ID" align="center" header-align="center" width="60"></el-table-column>
          <el-table-column prop="SceneName" :label="$t('resourcesManagement.resSceneName')" align="center"
            header-align="center"></el-table-column>
          <el-table-column prop="JobTypeStr" :label="$t('resourcesManagement.jobType')" align="center"
            header-align="center" width="120">
          </el-table-column>
          <el-table-column prop="SceneTypeStr" :label="$t('resourcesManagement.sceneType')" align="center"
            header-align="center" width="120">
            <template slot-scope="scope">
              <span :style="{ color: scope.row.SceneType == 'exclusive' ? 'red' : '' }">
                {{ scope.row.SceneTypeStr }}
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="IsSpecExclusiveStr" :label="$t('resourcesManagement.isExclusiveSpec')" align="center"
            header-align="center" width="120">
            <template slot-scope="scope">
              <span :style="{ color: scope.row.IsSpecExclusive == 'exclusive' ? 'red' : '' }">
                {{ scope.row.IsSpecExclusiveStr }}
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="ExclusiveOrg" :label="$t('resourcesManagement.exclusiveOrg')" align="center"
            header-align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.ExclusiveOrg ? scope.row.ExclusiveOrg : '--' }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="specStr" :label="$t('resourcesManagement.resourceSpecification')" align="left"
            header-align="center" min-width="380">
            <template slot-scope="scope">
              <span v-html="scope.row.specStr"></span>
            </template>
          </el-table-column>
          <el-table-column prop="AiCenterStr" :label="$t('resourcesManagement.aiCenter')" align="center"
            header-align="center" width="230">
            <template slot-scope="scope">
              <div v-if="!scope.row.queues.length">--</div>
              <div v-for="item in scope.row.queues " :key="item.key">
                <span>{{ item.AiCenterName }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="QueueStr" :label="$t('resourcesManagement.resQueue')" align="left"
            header-align="center" width="350">
            <template slot-scope="scope">
              <div v-if="!scope.row.queues.length">--</div>
              <div v-for=" item in scope.row.queues " :key="item.key">
                <span
                  v-html="item.QueueStr + item.priceStr + item.networkTypeStr + item.visualizationStr + item.statusStr"></span>
              </div>
            </template>
          </el-table-column>
          <el-table-column :label="$t('operation')" align="center" header-align="center" width="100" fixed="right">
            <template slot-scope="scope">
              <span class="op-btn" @click="showDialog('edit', scope.row)">{{ $t('edit') }}</span>
              <span class="op-btn" style="color:red" @click="deleteRow(scope.row)">{{ $t('delete') }}</span>
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
    <SceneDialog :visible.sync="sceneDialogShow" :type="sceneDialogType" :data="sceneDialogData"
      @confirm="sceneDialogConfirm"></SceneDialog>
  </div>
</template>

<script>
import SceneDialog from '../components/SceneDialog.vue';
import { getResQueueCode, getResSceneList, updateResScene, getAiCenterList } from '~/apis/modules/resources';
import { JOB_TYPE, CLUSTERS, ACC_CARD_TYPE, COMPUTER_RESOURCES, SPECIFICATION_STATUS, NETWORK_TYPE, NETWORK_TYPE_VALUE } from '~/const';
import { getListValueWithKey } from '~/utils';

export default {
  data() {
    return {
      selTaskType: '',
      taskTypeList: [{ k: '', v: this.$t('resourcesManagement.allJobType') }, ...JOB_TYPE],
      selIsSpecExclusive: '',
      IsSpecExclusiveList: [{ k: '', v: this.$t('resourcesManagement.allExclusiveAndCommonUseSpec') }, { k: 'exclusive', v: this.$t('resourcesManagement.exclusiveSpec') }, { k: 'public', v: this.$t('resourcesManagement.commonUseSpec') }],
      selSceneType: '',
      sceneTypeList: [{ k: '', v: this.$t('resourcesManagement.allSceneType') }, { k: 'public', v: this.$t('resourcesManagement.public') }, { k: 'exclusive', v: this.$t('resourcesManagement.exclusive') }],
      selQueue: '',
      queueList: [{ k: '', v: this.$t('resourcesManagement.allResQueue') }],
      selCluster: '',
      clusterList: [{ k: '', v: this.$t('resourcesManagement.allCluster') }, ...CLUSTERS],
      selAiCenter: '',
      aiCenterList: [{ k: '', v: this.$t('resourcesManagement.allAiCenter') }],
      accCardTypeList: [...ACC_CARD_TYPE],
      statusList: [{ k: '', v: this.$t('resourcesManagement.allStatus') }, ...SPECIFICATION_STATUS],
      selResource: '',
      resourceList: [{ k: '', v: this.$t('resourcesManagement.allComputeResource') }, ...COMPUTER_RESOURCES],
      selCardType: '',
      cardTypeList: [{ k: '', v: this.$t('resourcesManagement.allAccCardType') }, ...ACC_CARD_TYPE],
      selNetworkType: 0,
      networkTypeList: [{ k: 0, v: this.$t('resourcesManagement.allNetworkType') }, ...NETWORK_TYPE],
      selIsEnableVisualization: 0,
      IsEnableVisualizationList: [{ k: 0, v: this.$t('resourcesManagement.allIsEnableVisualization') }, { k: 1, v: this.$t('resourcesManagement.notEnable') }, { k: 2, v: this.$t('resourcesManagement.enable') }],
      loading: false,
      tableData: [],
      pageInfo: {
        curpage: 1,
        pageSize: 10,
        pageSizes: [10, 50, 100],
        total: 0,
      },
      sceneDialogShow: false,
      sceneDialogType: 'add',
      sceneDialogData: {},
    };
  },
  components: { SceneDialog },
  methods: {
    tableSpanMethod({ row, column, rowIndex, columnIndex }) {
      if ([0, 1, 2, 3, 4, 5, 9].indexOf(columnIndex) > -1) {
        if (row.rowspan) {
          return {
            rowspan: row.rowspan,
            colspan: 1,
          };
        } else {
          return {
            rowspan: 0,
            colspan: 0
          };
        }
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
          this.aiCenterList.splice(1, Infinity, ...data);
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
        jobType: this.selTaskType,
        sceneType: this.selSceneType,
        isSpecExclusive: this.selIsSpecExclusive,
        cluster: this.selCluster,
        queue: this.selQueue,
        center: this.selAiCenter,
        resource: this.selResource,
        cardType: this.selCardType,
        hasInternet: this.selNetworkType,
        enableVisualization: this.selIsEnableVisualization,
        page: this.pageInfo.curpage,
        pageSize: this.pageInfo.pageSize,
      };
      this.loading = true;
      getResSceneList(params).then(res => {
        this.loading = false;
        res = res.data;
        if (res.Code === 0) {
          const list = res.Data.List;
          const data = [];
          for (let i = 0, iLen = list.length; i < iLen; i++) {
            const item = list[i];
            const specs = item.Specs;
            const sourceSpecIdMap = {};
            for (let j = 0, jLen = specs.length; j < jLen; j++) {
              const spec = specs[j];
              const NGPU = `${spec.ComputeResource}:${spec.AccCardsNum + '*' + getListValueWithKey(this.accCardTypeList, spec.AccCardType)}`;
              const statusStr = spec.Status != '2' ? `<span style="color:rgb(245, 34, 45)">(${getListValueWithKey(this.statusList, spec.Status.toString())})</span>` : '';
              spec.specStr = `${NGPU}(${this.$t('resourcesManagement.gpuMem')}:${spec.GPUMemGiB}GB), CPU:${spec.CpuCores}, ${this.$t('resourcesManagement.mem')}:${spec.MemGiB}GB, ${this.$t('resourcesManagement.shareMem')}:${spec.ShareMemGiB}GB`;
              spec.JobTypeStr = getListValueWithKey(this.taskTypeList, item.JobType);
              spec.SceneTypeStr = getListValueWithKey(this.sceneTypeList.slice(1, Infinity), item.SceneType);
              spec.IsSpecExclusiveStr = item.SceneType == 'exclusive' ? '--' : getListValueWithKey(this.IsSpecExclusiveList.slice(1, Infinity), item.IsSpecExclusive);
              spec.statusStr = statusStr;
              spec.priceStr = `, ${this.$t('resourcesManagement.unitPrice')}:${spec.UnitPrice.toFixed(2)}${this.$t('resourcesManagement.point_hr')}`
              spec.networkTypeStr = `, ${this.$t('cloudbrainObj.networkType')}:${getListValueWithKey(NETWORK_TYPE_VALUE, spec.HasInternet)}`;
              spec.visualizationStr = `, ${this.$t('cloudbrainObj.visualization')}:${spec.EnableVisualization ? this.$t('resourcesManagement.enable') : this.$t('resourcesManagement.notEnable')}`;
              spec._id_ = item.ID;
              const sourceSpecId = spec.SourceSpecId;
              if (sourceSpecIdMap[sourceSpecId]) {
                sourceSpecIdMap[sourceSpecId].push({ ...spec });
              } else {
                sourceSpecIdMap[sourceSpecId] = [{ ...item, ...spec }];
              }
            }
            for (let key in sourceSpecIdMap) {
              const _specs = sourceSpecIdMap[key];
              if (key) {
                const aiCenters = [];
                const queues = [];
                for (let k = 0, kLen = _specs.length; k < kLen; k++) {
                  const _spec = _specs[k];
                  const queueName = _spec.QueueName ? `【${_spec.QueueName}】` : '';
                  const queueType = _spec.QueueType ? `【${_spec.QueueType}】` : '';
                  queues.push({
                    QueueId: _spec.QueueId,
                    QueueCode: _spec.QueueCode,
                    QueueName: _spec.QueueName,
                    QueueType: _spec.QueueType,
                    AiCenterCode: _spec.AiCenterCode,
                    AiCenterName: _spec.AiCenterName,
                    QueueStr: `${_spec.QueueCode}${queueName}${queueType}(${getListValueWithKey(this.clusterList, _spec.Cluster)} - ${_spec.AiCenterName})`,
                    statusStr: _spec.statusStr,
                    priceStr: _spec.priceStr,
                    networkTypeStr: _spec.networkTypeStr,
                    visualizationStr: _spec.visualizationStr,
                    key: Math.random(),
                  });
                }
                data.push({ ..._specs[0], aiCenters, queues })
              } else {
                for (let k = 0, kLen = _specs.length; k < kLen; k++) {
                  const _spec = _specs[k];
                  const queueName = _spec.QueueName ? `【${_spec.QueueName}】` : '';
                  const queueType = _spec.QueueType ? `【${_spec.QueueType}】` : '';
                  data.push({
                    ..._spec,
                    queues: [{
                      QueueId: _spec.QueueId,
                      QueueCode: _spec.QueueCode,
                      QueueName: _spec.QueueName,
                      QueueType: _spec.QueueType,
                      AiCenterCode: _spec.AiCenterCode,
                      AiCenterName: _spec.AiCenterName,
                      QueueStr: `${_spec.QueueCode}${queueName}${queueType}(${getListValueWithKey(this.clusterList, _spec.Cluster)} - ${_spec.AiCenterName})`,
                      statusStr: _spec.statusStr,
                      priceStr: _spec.priceStr,
                      networkTypeStr: _spec.networkTypeStr,
                      visualizationStr: _spec.visualizationStr,
                      key: Math.random(),
                    }]
                  })
                }
              }
            }
          }

          let rowspan = 1;
          let index = 0;
          for (let i = 0, iLen = data.length; i < iLen; i++) {
            if (i == 0) {
              data[i].rowspan = 1;
              continue;
            }
            if (data[i]._id_ == data[i - 1]._id_) {
              data[i].rowspan = 0;
              rowspan++;
              data[index].rowspan = rowspan;
            } else {
              index = i;
              rowspan = 1;
              data[index].rowspan = rowspan;
            }
          }
          this.tableData = data;
          this.pageInfo.total = res.Data.TotalSize;
        }
      }).catch(err => {
        console.log(err);
        this.loading = false;
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
    deleteRow(row) {
      this.$confirm(this.$t('resourcesManagement.resSceneDeleteConfirm'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm1'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        updateResScene({
          action: 'delete',
          ID: row._id_,
        }).then(res => {
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
          this.$message({
            type: 'error',
            message: this.$t('submittedFailed')
          });
        });
      }).catch(() => { });
    },
    showDialog(type, data) {
      this.sceneDialogType = type;
      this.sceneDialogData = data ? {
        ID: data._id_,
        SceneName: data.SceneName,
        JobType: data.JobType,
        SceneType: data.SceneType,
        IsSpecExclusive: data.IsSpecExclusive,
        ExclusiveOrg: data.ExclusiveOrg,
        Cluster: data.Cluster,
        ComputeResource: data.ComputeResource,
        SpecIds: data.Specs.map(item => item.ID),
      } : {};
      this.sceneDialogShow = true;
    },
    sceneDialogConfirm() {
      this.sceneDialogShow = false;
      this.getTableData();
    }
  },
  mounted() {
    this.getAiCenterList();
    this.getQueueList();
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
  align-items: center;
  justify-content: space-between;

  .select {
    margin-right: 10px;
    margin-bottom: 10px;

    /deep/ .el-input__inner {
      border-radius: 0;
    }
  }

  .el-button {
    margin-bottom: 10px;
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
