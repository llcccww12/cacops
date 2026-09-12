<template>
  <div>
    <TopHeader>
      <div class="title">
        <div class="title-1">鹏城·脑海(原鹏城·盘古)大模型</div>
        <div class="title-2">以中文为核心的鹏城·脑海大语言模型，更强的任务理解与处理能力。</div>
        <div class="title-3">
          <a target="_blank" href="https://openi.pcl.ac.cn/PCL-Platform.Intelligence/pcl_pangu">项目主页</a>
        </div>
      </div>
    </TopHeader>
    <div class="type-list-container">
      <div class="ui container">
        <div class="content-tips">
          <span>* 社区提供了三个示例场景，用户可以使用提供的示例数据集新建微调任务，也可以上传自己的数据集，定制符合自己的应用场景的鹏城·脑海大模型。</span>
          <a target="_blank"
            href="https://openi.pcl.ac.cn/PCL-Platform.Intelligence/pcl_pangu/src/branch/master/docs/README_DATASET.md">如何构造数据集</a>
        </div>
        <div class="type-list">
          <div class="item" :class="item.class ? item.class : ''" v-for="(item, index) in list" :key="index">
            <div class="item-title">{{ item.name }}</div>
            <div class="item-descr" :title="item.descr">{{ item.descr }}</div>
            <div class="item-op">
              <a class="item-op-btn" :href="item.href">
                <span>新建任务</span>
                <i class="el-icon-right"></i>
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div>
      <div class="ui container">
        <div class="table-container">
          <el-table ref="tableRef" border :data="tableData" style="width:100%;" v-loading="loading" stripe>
            <el-table-column label="任务名称" align="center" header-align="center">
              <template slot-scope="scope">
                <a :href="`/${userName}/${repoName}/modelarts/train-job/${scope.row.id}`" target="_blank"> {{
                  scope.row.display_job_name }}</a>
              </template>
            </el-table-column>
            <el-table-column label="任务场景" prop="jobCategoryStr" align="center" header-align="center"></el-table-column>
            <el-table-column label="训练状态" prop="status" align="center" header-align="center">
              <template slot-scope="scope">
                <div class="job-status-c">
                  <i class="job-status-icon" :class="scope.row.status"></i>
                  <span>{{ scope.row.status }}</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column label="部署状态" align="center" header-align="center">
              <template slot-scope="scope">
                <div class="job-status-c">
                  <i class="job-status-icon"
                    :class="[scope.row.deploy_status, deployStateMap[scope.row.deploy_status]]"></i>
                  <span>{{ scope.row.deploy_status }}</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column label="创建时间" prop="createdTimeStr" align="center" header-align="center"></el-table-column>
            <el-table-column label="操作" align="center" header-align="center" width="220">
              <template slot-scope="scope">
                <div class="op-btn-c">
                  <a class="btn" href="javascript:;" v-if="scope.row.canInference"
                    @click="startInference(scope.row)">体验</a>
                  <a class="btn" href="javascript:;" v-if="scope.row.canReDeploy"
                    @click="updateDeploy(scope.row, 'running')">启动</a>
                  <a class="btn" href="javascript:;" v-if="scope.row.canStopTrain" @click="stopTrain(scope.row)">停止</a>
                  <a class="btn" href="javascript:;" v-if="scope.row.canStopDeploy"
                    @click="updateDeploy(scope.row, 'stopped')">停止</a>
                  <a class="btn" href="javascript:;" v-if="scope.row.canDeploy" @click="deploy(scope.row)">立即部署</a>
                  <a class="btn" href="javascript:;" v-if="scope.row.canDelete" @click="deleteTask(scope.row)">删除</a>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <div class="table-tips">* 最多创建 5 条任务，超过 30 天的任务将不能部署或者再次启动。</div>
      </div>
    </div>
    <LoadingMask :loading="maskLoading" :tips="maskLoadingTips"></LoadingMask>
  </div>
</template>

<script>
import dayjs from 'dayjs';
import TopHeader from '../components/TopHeader.vue';
import LoadingMask from '../components/cloudbrain/LoadingMask.vue';
import { getFinetuneList, getJobStatus, stopTrainJob, deleteTrainJob, setFinetuneService, getFinetuneServiceStatus, updateFinetuneService } from '~/apis/modules/modelbase';
import { getListValueWithKey } from '~/utils';

const TrainJobFinalState = [
  'STOPPED', 'CREATE_FAILED', 'UNAVAILABLE', 'DELETED', 'RESIZE_FAILED', 'SUCCEEDED', 'IMAGE_FAILED', 'SUBMIT_FAILED', 'DELETE_FAILED',
  'KILLED', 'COMPLETED', 'FAILED', 'CANCELED', 'LOST', 'START_FAILED', 'SUBMIT_MODEL_FAILED', 'DEPLOY_SERVICE_FAILED', 'CHECK_FAILED',
];

const DeployStateMap = {
  'BUILDING': 'WAITING',
  'WAITING': 'WAITING',
  'DEPLOYING': 'RUNNING',
  'SUCCEEDED': 'SUCCEEDED',
  'STOP': 'STOPPED',
  'FAIL': 'FAILED',
}

export default {
  data() {
    return {
      deployStateMap: DeployStateMap,
      userName: '',
      repoName: 'openi-notebook',
      list: [{
        _name: '文本分类',
        name: '文本分类',
        descr: '根据文本分类选项对文本内容进行自动分类，每段文本内容均可属于某一种选项类别。',
        key: 1,
        href: '/extension/modelbase/pangufinetune/create?type=1',
      }, {
        _name: '中英翻译',
        name: '中英翻译',
        descr: '指定翻译的目标语言，将源语言翻译成目标语言。支持中译英、英译中。',
        key: 2,
        href: '/extension/modelbase/pangufinetune/create?type=2',
      }, {
        _name: '情感分类',
        name: '情感分类',
        descr: '输入一段文本，对其情感进行分类，判断其属于积极的或者消极的。',
        key: 3,
        href: '/extension/modelbase/pangufinetune/create?type=3',
      }, {
        _name: '自定义',
        name: '自定义更多场景',
        descr: '支持用户自定义数据模板，微调以适用于更多业务场景，例如摘要生成、对对联、阅读理解等。',
        key: 0,
        class: 'item-green',
        href: '/extension/modelbase/pangufinetune/create?type=0',
      }],
      loading: false,
      tableData: [],
      maxJobNum: '-',

      refreshTimer: null,
      isOperating: false,

      maskLoading: false,
      maskLoadingTips: '',
    };
  },
  components: { TopHeader, LoadingMask },
  methods: {
    getTableData() {
      if (!this.userName) return;
      this.loading = true;
      this.refreshTimer && clearInterval(this.refreshTimer);
      getFinetuneList({ userName: this.userName }).then(res => {
        res = res.data;
        this.loading = false;
        this.maxJobNum = res.maxJobNum || this.maxJobNum;
        const data = res.fineTuneJobs;
        this.tableData = data.map((item) => {
          const taskJob = {
            ...item,
            jobCategoryStr: getListValueWithKey(this.list, item.job_category, 'key', '_name'),
            createdTimeStr: dayjs(item.created_unix * 1000).format('YYYY-MM-DD HH:mm:ss'),
          }
          this.checkOperation(taskJob);
          return taskJob;
        });
        this.startRefreshTaskTimer();
      }).catch(err => {
        console.log(err);
        this.loading = false;
      });
    },
    checkOperation(taskJob) {
      taskJob.canStopTrain = ['RUNNING', 'WAITING'].includes(taskJob.status);
      taskJob.canDelete = ['STOPPED', 'FAILED', 'START_FAILED', 'KILLED', 'COMPLETED', 'SUCCEEDED', 'CREATE_FAILED'].includes(taskJob.status)
        && ['', 'STOP', 'FAILED'].includes(taskJob.deploy_status);
      taskJob.canDeploy = !taskJob.cleared && ['SUCCEEDED', 'COMPLETED'].includes(taskJob.status) && !['BUILDING', 'WAITING', 'DEPLOYING', 'SUCCEEDED', 'STOP', 'FAILED'].includes(taskJob.deploy_status);
      taskJob.canReDeploy = !taskJob.cleared && ['SUCCEEDED', 'COMPLETED'].includes(taskJob.status) && ['STOP'].includes(taskJob.deploy_status);
      taskJob.canStopDeploy = ['SUCCEEDED', 'COMPLETED'].includes(taskJob.status) && ['DEPLOYING', 'SUCCEEDED'].includes(taskJob.deploy_status);
      taskJob.canInference = ['SUCCEEDED', 'COMPLETED'].includes(taskJob.status) && ['SUCCEEDED'].includes(taskJob.deploy_status);
    },
    startRefreshTaskTimer() {
      this.refreshTimer = setInterval(() => {
        for (let i = 0, iLen = this.tableData.length; i < iLen; i++) {
          const taskJob = this.tableData[i];
          this.refreshTaskStatus(taskJob);
        }
      }, 6 * 1000);
    },
    refreshTaskStatus(taskJob, notCheck) {
      if (notCheck || !TrainJobFinalState.includes(taskJob.status)) {
        getJobStatus({
          userName: this.userName,
          jobId: taskJob.job_id,
        }).then(res => {
          const data = res.data || {};
          taskJob.status = data.JobStatus;
          taskJob.start_time = data.StartTime;
          taskJob.ai_center = data.AiCenter;
          taskJob.job_duration = data.JobDuration;
          this.checkOperation(taskJob);
        }).catch(err => {
          console.log(err);
        });
      }
      if (notCheck || ['BUILDING', 'WAITING', 'DEPLOYING', 'SUCCEEDED'].includes(taskJob.deploy_status)) {
        getFinetuneServiceStatus({
          userName: this.userName,
          jobId: taskJob.job_id,
        }).then(res => {
          const data = res.data;
          taskJob.deploy_status = data.fineTuneDeployStatus;
          this.checkOperation(taskJob);
        }).catch(err => {
          console.log(err);
        });
      }
    },
    stopTrain(taskJob) {
      if (this.isOperating) return;
      this.isOperating = true;
      this.maskLoadingTips = '正在停止，请稍后';
      this.maskLoading = true;
      stopTrainJob({
        userName: this.userName,
        jobId: taskJob.job_id,
      }).then(res => {
        this.isOperating = false;
        this.maskLoading = false;
        this.refreshTaskStatus(taskJob, true);
        setTimeout(() => {
          this.checkOperation(taskJob);
        }, 200);
      }).catch(err => {
        this.isOperating = false;
        this.maskLoading = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: '操作失败',
        });
      })
    },
    deleteTask(taskJob) {
      if (this.isOperating) return;
      this.$confirm(`您确认删除该任务么？此任务一旦删除不可恢复。`, '提示', {
        confirmButtonText: '确定操作',
        cancelButtonText: '取消操作',
        type: 'warning',
      }).then(async () => {
        this.isOperating = true;
        this.maskLoadingTips = '任务删除中，请稍后';
        this.maskLoading = true;
        deleteTrainJob({
          userName: this.userName,
          jobId: taskJob.job_id,
          id: taskJob.id,
        }).then(res => {
          this.isOperating = false;
          this.maskLoading = false;
          this.getTableData();
        }).catch(err => {
          this.isOperating = false;
          this.maskLoading = false;
          console.log(err);
          this.$message({
            type: 'error',
            message: '操作失败',
          });
        })
      }).catch(() => { });
    },
    deploy(taskJob) {
      if (this.isOperating) return;
      this.isOperating = true;
      this.maskLoadingTips = '正在部署，请稍后';
      this.maskLoading = true;
      setFinetuneService({
        userName: this.userName,
        jobId: taskJob.job_id,
      }).then(res => {
        this.isOperating = false;
        this.maskLoading = false;
        res = res.data;
        if (res.code != 0) {
          this.$message({
            type: 'error',
            message: res.message,
          });
        }
        this.refreshTaskStatus(taskJob, true);
      }).catch(err => {
        console.log(err);
        this.isOperating = false;
        this.maskLoading = false;
        this.$message({
          type: 'error',
          message: '操作失败',
        });
      })
    },
    updateDeploy(taskJob, type) { // type=stop,running
      if (this.isOperating) return;
      this.isOperating = true;
      this.maskLoadingTips = '正在操作，请稍后';
      this.maskLoading = true;
      updateFinetuneService({
        userName: this.userName,
        jobId: taskJob.job_id,
        status: type,
      }).then(res => {
        this.isOperating = false;
        this.maskLoading = false;
        res = res.data;
        if (res.code != 0) {
          this.$message({
            type: 'error',
            message: res.message,
          });
        }
        this.refreshTaskStatus(taskJob, true);
      }).catch(err => {
        console.log(err);
        this.isOperating = false;
        this.maskLoading = false;
        this.$message({
          type: 'error',
          message: '操作失败',
        });
      });
    },
    startInference(taskJob) {
      if (this.isOperating) return;
      window.open(`/extension/modelbase/pangufinetune/inference?jobid=${taskJob.job_id}&type=${taskJob.type}&jobcategory=${taskJob.job_category}`);
    },
  },
  beforeMount() {
    const metaEl = document.querySelectorAll('meta[name="_uid"]');
    if (metaEl.length) {
      const uid = metaEl[0].getAttribute('content');
      const uname = metaEl[0].getAttribute('content-ext');
      this.userName = uname;
      this.getTableData();
    }
  },
  mounted() { },
  beforeDestroy() {
    this.refreshTimer && clearInterval(this.refreshTimer);
  },
};
</script>

<style scoped lang="less">
.title {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  margin-top: -10px;

  .title-1 {
    font-weight: 400;
    font-size: 28px;
    color: rgb(16, 16, 16);
    height: 42px;
  }

  .title-2 {
    font-weight: 400;
    font-size: 14px;
    color: rgba(16, 16, 16, 1);
    margin: 12px 0 10px;
  }

  .title-3 {
    font-weight: 400;
    font-size: 14px;

    a {
      color: rgba(3, 102, 214, 1);
      text-decoration: underline;
    }
  }
}

.content-tips {
  font-size: 14px;
  color: rgb(136, 136, 136);
  padding-left: 10px;

  a {
    font-size: 14px;
    color: rgba(3, 102, 214, 1);
    text-decoration: underline;
  }
}

.type-list-container {
  padding: 60px 0 60px 0;
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(0.11899999999999993%2C%201.217%2C%20-0.0901857098765432%2C%200.11899999999999993%2C%200.269%2C%20-0.22)%22%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%220.47%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23e5e7eb%22%20stop-opacity%3D%220.3%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");

  .type-list {
    display: flex;
    justify-content: center;

    .item {
      width: 318px;
      height: 160px;
      border-color: rgba(157, 197, 226, 0.4);
      border-width: 1px;
      border-style: solid;
      box-shadow: rgba(157, 197, 226, 0.2) 0px 5px 10px 0px;
      color: rgb(16, 16, 16);
      border-radius: 6px;
      margin: 20px 10px;
      padding: 20px 15px;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.2531545429373838%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23eef2ff%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");

      .item-title {
        text-align: center;
        font-weight: 500;
        font-size: 16px;
        color: rgb(16, 16, 16);
      }

      .item-descr {
        font-weight: 300;
        font-size: 12px;
        color: rgb(136, 136, 136);
        margin: 10px 0px;
        height: 60px;
        overflow: hidden;
        text-overflow: ellipsis;
        word-break: break-all;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 3;
      }

      .item-op {
        display: flex;
        align-items: center;
        justify-content: center;

        .item-op-btn {
          font-size: 14px;
          color: rgb(50, 145, 248);
        }
      }

      &.item-green {
        background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.2531545429373838%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23cffff0%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
      }
    }
  }
}

.table-container {
  margin-top: 40px;

  /deep/ .el-table__header {
    th {
      background: rgb(249, 249, 249);
      font-size: 12px;
      color: rgb(136, 136, 136);
      font-weight: normal;
    }
  }

  /deep/ .el-table__body {
    td {
      font-size: 12px;
    }
  }

  /deep/ .el-radio__label {
    display: none;
  }

  .job-status-c {
    display: flex;
    align-items: center;
    padding-left: 10px;

    .job-status-icon {
      margin-right: 4px;
    }
  }

  .op-btn-c {
    display: flex;
    align-items: center;
    justify-content: center;

    .btn {
      margin: 0 4px;
    }
  }
}

.table-tips {
  margin-top: 10px;
  font-weight: 300;
  font-size: 14px;
  color: rgba(250, 140, 22, 1);
}
</style>
