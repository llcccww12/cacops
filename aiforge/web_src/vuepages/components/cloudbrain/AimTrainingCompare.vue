<template>
  <div class="aim-visual-compare-btn-c">
    <el-button v-if="hasAimRight" class="op-btn" type="primary" @click="dlgShow = true">
      <div class="btn-content">
        <span>{{ $t('cloudbrainObj.aimTrainCompare') }}</span>
      </div>
    </el-button>
    <el-dialog class="aim-visual-compare-task-select-dlg" :visible.sync="dlgShow"
      :title="$t('cloudbrainObj.aimTrainCompareDlgTitle')" width="1000px" :modal="true" :close-on-click-modal="false"
      :show-close="true" :destroy-on-close="false" :before-close="beforeClose" @open="open" @closed="closed">
      <div class="dlg-content">
        <div class="list-body table-container">
          <el-table :data="tableData" style="min-width:100%" v-loading="loading" stripe :max-height="450"
            @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="55" align="center" header-align="center" fixed>
            </el-table-column>
            <el-table-column :label="$t('cloudbrainObj.taskName')" align="left" header-align="left" width="200">
              <template slot-scope="scope">
                <a class="dispaly-job-name" :href="scope.row.task.jobLink">
                  {{ scope.row.task.display_job_name }}
                </a>
              </template>
            </el-table-column>
            <el-table-column :label="$t('cloudbrainObj.cluster')" align="center" header-align="center" width="125">
              <template slot-scope="scope">
                <span>{{ scope.row.task.clusterName }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="Status" :label="$t('status')" align="left" header-align="center" width="145">
              <template slot-scope="scope">
                <div class="status-wrap">
                  <i :class="scope.row.task.status"></i>
                  <span>{{ scope.row.task.status }}</span>
                  <i v-if="scope.row.task.detailed_status === 'dataMigrating' && scope.row.task.status === 'WAITING'"
                    :class="scope.row.task.detailed_status" :title="$t('cloudbrainObj.migratingData')"></i>
                  <i v-if="scope.row.task.detailed_status === 'centerPending' && scope.row.task.status === 'WAITING'"
                    :class="scope.row.task.detailed_status" :title="$t('cloudbrainObj.centerPending')"></i>
                  <i v-if="scope.row.task.detailed_status === 'ImagePulling' && scope.row.task.status === 'WAITING'"
                    :class="scope.row.task.detailed_status" :title="$t('cloudbrainObj.imagePulling')"></i>
                </div>
              </template>
            </el-table-column>
            <el-table-column :label="$t('cloudbrainObj.cloudbrainTaskType')" align="center" header-align="center"
              width="125">
              <template slot-scope="scope">
                <span>{{ scope.row.task.jobTypeShow }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="created_unix" :label="$t('cloudbrainObj.createTime')" align="center"
              header-align="center" width="160">
              <template slot-scope="scope">
                <span>{{ dateFormat(scope.row.task.created_unix) }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="formatted_duration" :label="$t('cloudbrainObj.runDuration')" align="center"
              header-align="center" width="110">
              <template slot-scope="scope">
                <span>{{ scope.row.task.formatted_duration }}</span>
              </template>
            </el-table-column>
            <el-table-column :label="$t('cloudbrainObj.computeResource')" align="center" header-align="center"
              width="160">
              <template slot-scope="scope">
                <span>{{ scope.row.task.computeSourceShow }}</span>
              </template>
            </el-table-column>
            <el-table-column :label="$t('resourcesManagement.aiCenter')" align="center" header-align="center"
              width="200">
              <template slot-scope="scope">
                <span>{{ scope.row.task.ai_center }}</span>
              </template>
            </el-table-column>
            <el-table-column :label="$t('resourcesManagement.accCardType')" align="center" header-align="center"
              width="150">
              <template slot-scope="scope">
                <span>{{ scope.row.task.accCardTypeShow }}</span>
              </template>
            </el-table-column>
            <el-table-column prop="creator" :label="$t('modelManage.creator')" align="left" min-width="80"
              :fixed="isMiniScreen ? undefined : 'right'" header-align="center">
              <template slot-scope="scope">
                <div class="creator-wrap">
                  <a v-if="scope.row.creator.name" :href="'/' + scope.row.creator.name"
                    :title="scope.row.creator.full_name">
                    <img :src="scope.row.creator.rel_avatar_link">
                  </a>
                  <span v-else title="Ghost">
                    <img src="/user/avatar/Ghost/-1">
                  </span>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div class="list-foot">
          <el-button type="primary" @click="confirm">{{ $t('datasetObj.dataset_ok') }}</el-button>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { CloudBrainTools } from '~/pages/cloudbrain/tools';

import { CLUSTERS, JOB_TYPE, ACC_CARD_TYPE } from '~/const';
import { getAimRight, getAiTaskList, getAimUrl } from '~/apis/modules/cloudbrain';
import { timeSinceUnix, getListValueWithKey } from '~/utils';
import dayjs from 'dayjs';

const cloudBrainTools = new CloudBrainTools();

export default {
  name: "AimTrainingCompare",
  props: {
    repoOwnerName: {
      type: String,
    },
    repoName: {
      type: String,
    }
  },
  data() {
    return {
      hasAimRight: false,
      loading: false,
      dlgShow: false,
      tableData: [],
      selectedData: [],
      page: 1,
      pageSize: 500,
      total: 0,
      isMiniScreen: false,
      operating: false,
    };
  },
  methods: {
    open() {
      this.calcScreenInfo();
      this.tableData = [];
      this.selectedData = [];
      this.getTableData();
    },
    getTableData() {
      this.loading = true;
      getAiTaskList({
        repoOwnerName: this.repoOwnerName,
        repoName: this.repoName,
        job_type: 'TRAIN',
        exclude_job_types: 'HPC',
        page: this.page,
        page_size: this.pageSize,
        aim: true,
      }).then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 0) {
          const data = res.data;
          this.canCreateTask = data.can_create_task;
          let isSubscriber = data.is_subscriber;
          this.isRepoEmpty = data.is_repo_empty;
          this.total = data.total;
          data.tasks.forEach(item => {
            const obj = Object.assign({}, item);
            delete obj.task;
            const task = item.task;
            Object.assign(task, obj);
            task.computeSourceShow = task.compute_source == 'GPU' ? 'CPU/GPU' : task.compute_source;
            task.clusterName = getListValueWithKey(CLUSTERS, task.cluster);
            task.accCardTypeShow = getListValueWithKey(ACC_CARD_TYPE, task.acc_card_type);
            task.repoOwnerName = this.repoOwnerName;
            task.repoName = this.repoName;
            task.jobLink = '';
            task.createdFromNow = this.calcFromNow(task.created_unix);
            task.jobTypeShow = getListValueWithKey(JOB_TYPE, task.job_type);
            task.isSubscriber = isSubscriber && task.end_point;
            cloudBrainTools.checkOperation(task);
          });
          this.tableData = (data.tasks || []).filter(item => ["RUNNING", "STOPPED", "FAILED", "SUCCEEDED", "COMPLETED"].includes(item.task.status));
          cloudBrainTools.initRefreshData(this.tableData);
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    calcFromNow(unix) {
      return timeSinceUnix(unix, Date.now() / 1000);
    },
    dateFormat(unix) {
      return dayjs(unix * 1000).format('YYYY-MM-DD HH:mm:ss');
    },
    beforeClose(done) {
      cloudBrainTools.stop();
      done();
    },
    closed() { },
    handleSelectionChange(val) {
      this.selectedData = val;
    },
    confirm() {
      if (this.operating) return;
      if (this.selectedData.length <= 1 || this.selectedData.length > 10) {
        this.$message.info(this.$t('cloudbrainObj.aimTrainCompareSelectTips', { minCount: 2, maxCount: 10 }));
        return;
      }
      this.operating = true;
      getAimUrl({
        repoOwnerName: this.repoOwnerName,
        repoName: this.repoName,
        job_name: this.selectedData.map(item => item.task.display_job_name)
      }).then(res => {
        res = res.data;
        this.operating = false;
        if (res.code == 0) {
          if (res.data && res.data.url) {
            if (!window.open(res.data.url)) {
              window.location.href = res.data.url;
            }
          }
        } else {
          this.$message({
            type: 'error',
            message: res.msg || this.$t('operationFailed'),
          });
        }
      }).catch(err => {
        console.log(err);
        this.operating = false;
        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      })
    },
    calcScreenInfo() {
      this.isMiniScreen = document.documentElement.clientWidth <= 800;
    },
  },
  beforeMount() {
    getAimRight().then(res => {
      this.hasAimRight = !!res.data.aim_right;
    }).catch(err => console.log(err))
  },
  beforeDestroy() {
    cloudBrainTools.stop();
  },
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';

.aim-visual-compare-task-select-dlg {
  /deep/.el-dialog__body {
    padding-top: 5px;
    padding-bottom: 20px;
  }

  .op-btn {
    display: flex;
    align-items: center;
    height: 38px;
    font-size: 14px;
    background: rgba(22, 132, 252, 0.9);
    border-radius: 4px;

    &:active {
      background: rgb(22, 132, 252, 1);
    }


    &:focus,
    &:hover {
      background: rgba(22, 132, 252, 0.8);
    }

    .btn-content {
      display: flex;
      align-items: center;

      i {
        font-size: 14px;
        margin-right: 10px;
      }
    }
  }

  .dlg-content {
    .table-container {
      /deep/ .el-table__body {
        td {
          height: 64px;
        }
      }

      .dispaly-job-name {
        font-size: 14px;
        font-weight: bold;
      }

      .status-wrap {
        display: flex;
        align-items: center;

        i,
        span {
          margin-right: 4px;
        }
      }

      .creator-wrap {
        display: flex;
        align-items: center;
        justify-content: center;

        img {
          height: 28px;
          width: 28px;
          border-radius: 100%;
        }
      }
    }

    .list-foot {
      margin-top: 20px;
      display: flex;
      justify-content: flex-end;
      padding: 0 10px;
    }
  }
}
</style>
