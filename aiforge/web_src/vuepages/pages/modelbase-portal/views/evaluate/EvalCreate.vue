<template>
  <div class="content-box">
    <Header></Header>
    <div class="main-body">
      <div class="main-content">
        <div class="err-msg-box-c" v-if="errorMsgBoxShow">
          <div class="err-msg-box">
            <p>{{ errorMsg }}</p>
          </div>
        </div>
        <div v-if="!alreadyMsgBoxShow" class="create-tips">
          <i class="ri-information-line"></i>
          <span style="margin-left:4px;">{{ $t('cloudbrainObj.runningLimit', {
            count: limitCount, taskType:
              $t('modelSquare.modelEvaluateTask')
          }) }}</span>
        </div>
        <TaskLimitDialog :visible="limitDialogVisible" :title="limitDialogTitle" :desc="limitDialogDesc"
          :taskList="limitTaskList" @close="limitDialogVisible = false" @stopSuccess="onLimitStopSuccess"
          @deleteSuccess="onLimitDeleteSuccess" @taskStatusUpdate="onLimitTaskStatusUpdate"
          @allTasksTerminal="onLimitAllTasksTerminal">
        </TaskLimitDialog>

        <div class="area" v-loading="globalLoading">
          <div class="area-content">
            <div class="form-row">
              <div class="left-area">
                <div class="title">

                </div>
                <div class="content">
                  <el-radio-group v-model="modelRadio" size="medium" @change="changeModelType">
                    <el-radio label="base" border>{{ $t('modelFinetune.selectBaseModel') }}</el-radio>
                    <el-radio label="finetune" border>{{ $t('modelFinetune.selectFtModel') }}</el-radio>
                  </el-radio-group>
                </div>
              </div>
              <div class="right-area"></div>
            </div>
            <div class="form-row" v-if="modelRadio === 'base'">
              <div class="left-area">
                <div class="title">
                  <span class="required">{{ $t('repos.model') }}</span>
                </div>
                <div class="content">
                  <el-select v-model="modelName" class="field-input" @change="changeModel">
                    <el-option v-for="item in modelList" :value="item.name" :key="item.name"
                      :label="item.name"></el-option>
                  </el-select>
                </div>
              </div>
              <div class="right-area"></div>
            </div>
            <div class="form-row" v-else>
              <div class="left-area">
                <div class="title">
                  <span class="required">{{ $t('modelFinetune.sftFinetuneName') }}</span>
                </div>
                <div class="content">
                  <el-select v-model="ftModelName" class="field-input" @change="changeFtModel">
                    <el-option v-for="item in ftModelList" :value="item.display_job_name" :key="item.id"
                      :label="item.display_job_name"></el-option>
                  </el-select>
                </div>
              </div>
              <div class="right-area"></div>
            </div>
            <FormTopV2 ref="formTopRef" :resourceObj="pageCfg" :queueNum="queueNum" @change="changeComputeResouce">
            </FormTopV2>
            <SpecSelect ref="specRef" v-model="state.spec" :required="true" :configs="specConfigs" networkType="all"
              :loading="loading"></SpecSelect>
            <DatasetSelectV2 ref="datasetRef" v-if="Object.keys(datasetList).length > 0" v-model="datasets"
              :required="true" :datasetList="datasetList"></DatasetSelectV2>

            <div class="form-row">
              <div class="left-area">
                <div class="title">
                  <span class="required">{{ $t('modelFinetune.evalDatasetLimit') }}</span>
                </div>
                <div class="content">
                  <el-select v-model="limitNum">
                    <el-option v-for="item in limitNumList" :value="item.v" :key="item.name"
                      :label="item.k"></el-option>
                  </el-select>
                </div>
              </div>
              <div class="right-area"></div>
            </div>
            <TaskName ref="taskNameRef" v-model="state.taskName" :required="true" :userName="loginName"></TaskName>
            <ForbidPenetrationTipCheck v-model="agreeForbidPenetration"></ForbidPenetrationTipCheck>
            <div class="form-row">
              <div class="left-area">
                <div class="title"></div>
                <div class="content">
                  <el-button
                    :disabled="maskLoading || !agreeForbidPenetration || !specConfigs.specs['all'] || alreadyMsgBoxShow"
                    size="default" class="submit-btn" @click="submit">{{ $t('cloudbrainObj.createTask')
                    }}</el-button>
                  <el-button class="cancel-btn" size="default" @click="cancel">{{ $t('cancel') }}</el-button>
                </div>
              </div>
            </div>
            <div class="form-row" style="display:flex;justify-content:center;margin-top:45px;margin-bottom:-10px;">
              <AcknowledgementsTips></AcknowledgementsTips>
            </div>
          </div>
          <LoadingMask :loading="maskLoading" :content="maskLoadingContent"></LoadingMask>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from '../../components/Header.vue';
import FormTopV2 from '~/components/cloudbrain/FormTopV2.vue';
import TaskName from '~/components/cloudbrain/TaskName.vue';
import SpecSelect from '~/components/cloudbrain/SpecSelect.vue';
import LoadingMask from '../../../modelbase/components/cloudbrain/LoadingMask.vue';
import AcknowledgementsTips from '~/components/AcknowledgementsTips.vue';
import TaskLimitDialog from '~/components/cloudbrain/TaskLimitDialog.vue';
import ForbidPenetrationTipCheck from '~/components/cloudbrain/ForbidPenetrationTipCheck.vue';
import DatasetSelectV2 from './DatasetSelectV2.vue';
import { getAiTaskPrepareInfo, createAiTask, getMyAiTasks } from '~/apis/modules/cloudbrain';
import { getPromoteData } from '~/apis/modules/common';
import { CLUSTERS, JOB_TYPE, ACC_CARD_TYPE } from '~/const';
import { COMPUTER_RESOURCES_TITLE } from '~/pages/cloudbrain/configs';
import { getUrlSearchParams, getListValueWithKey } from '~/utils';
export default {
  name: 'ExpCreate',
  data() {
    return {
      loginName: '',
      pageCfg: {
        configs: {
          hideCluster: true,
          hideTips2: true,
        },
        cluster: 'C2Net',
        clusters: CLUSTERS.filter(item => item.k == 'C2Net').map(item => ({ ...item, key: item.k, label: item.v })),
        computerResouce: '',
        computerResouces: [],
      },
      modelList: [],
      modelName: '',
      ftModelName: '',
      ftModelList: [],
      state: {
        spec: '',
        taskName: '',
        dataset: ''
      },
      specConfigs: {
        specs: {},
        blance: 0,
        showPoint: false,
      },
      datasets: [],
      datasetList: {},
      limitNum: '10',
      limitNumList: [{ k: this.$t('all'), v: "0" }, { k: "10", v: "10" }, { k: "20", v: "20" }, { k: "50", v: "50" }, { k: "100", v: "100" }],
      dataLoading: false,
      maskLoading: false,
      maskLoadingContent: '',
      errorMsgBoxShow: false,
      errorMsg: '',
      DepModelInfo: {},
      ftModelInfo: [],
      queueNum: 1,
      computeSource: '',
      loading: false,
      alreadyMsgBoxShow: false,
      notStopTaskCount: 0,
      limitCount: 0,
      limitDialogVisible: false,
      limitTaskList: [],
      modelRadio: 'base',
      urlParams: null,
      globalLoading: false,
      promoteFtList: [],
      restoreModelInfo: [],
      agreeForbidPenetration: true,
    };
  },
  components: { Header, LoadingMask, FormTopV2, SpecSelect, TaskName, AcknowledgementsTips, ForbidPenetrationTipCheck, DatasetSelectV2, TaskLimitDialog },
  computed: {
    limitDialogTitle() {
      return this.$t('cloudbrainObj.runningLimit', { count: this.limitCount, taskType: this.$t('modelSquare.modelEvaluateTask') });
    },
    limitDialogDesc() {
      return this.$t('cloudbrainObj.sameTaskTips1', { count: this.notStopTaskCount });
    },
  },
  methods: {
    changeModel(_, noPrepare) {
      console.log("changeModel", this.modelName)
      this.DepModelInfo = this.modelList.filter((item) => {
        return item.name === this.modelName;
      });
      console.log("DepModelInfo", this.DepModelInfo)
      const modelInfo = this.DepModelInfo[0];
      const compute_sources = Object.keys(modelInfo.configs?.compute_sources || {});
      const useComputeResources = COMPUTER_RESOURCES_TITLE
        .filter(item => compute_sources.indexOf(item.k) >= 0)
        .map(item => ({ ...item, key: item.k, label: item.v }));
      this.pageCfg.computerResouces = useComputeResources;
      this.pageCfg.computerResouce = useComputeResources[0]?.k || '';
      if (noPrepare) return;
      this.prepare();
    },
    changeFtModel(_, noPrepare) {
      this.ftModelInfo = this.ftModelList.filter((item) => {
        return item.display_job_name === this.ftModelName;
      });
      if (!this.ftModelInfo.length) return
      console.log('this.ftModelInfo', this.ftModelInfo)
      let modelName = this.ftModelInfo[0].app_name
      let computerResouce = this.ftModelInfo[0].compute_source
      if (modelName) {
        this.restoreModelInfo = this.promoteFtList.filter(item => item.name == modelName)[0].experience
      }

      const useComputeResources = COMPUTER_RESOURCES_TITLE
        .filter(item => computerResouce.indexOf(item.k) >= 0)
        .map(item => ({ ...item, key: item.k, label: item.v }));
      console.log('useComputeResources', useComputeResources)
      this.pageCfg.computerResouces = useComputeResources;
      this.pageCfg.computerResouce = useComputeResources[0]?.k || '';
      console.log("noPrepare", noPrepare)
      if (noPrepare) return;
      let { repo_owner_name, repo_name } = this.restoreModelInfo
      console.log("repo_owner_name", repo_owner_name)
      this.prepare(repo_owner_name, repo_name);

    },
    changeComputeResouce(data) {
      if (this.pageCfg.cluster != data.cluster || this.pageCfg.computerResouce != data.computerResouce) {
        this.pageCfg.cluster = data.cluster;
        this.pageCfg.computerResouce = data.computerResouce;
        this.prepare();
      }
    },
    fetchLimitTaskList() {
      const params = {
        job_type: 'EVAL',
        job_status: 'nostop',
        ai_center: '', cluster: '', compute_source: '', q: '',
        page: 1, pageSize: 5
      };
      getMyAiTasks(params).then(res => {
        if (res.data.code === 0) {
          this.limitTaskList = res.data.data.tasks || [];
          this.limitTaskList.forEach((item) => {
            const task = item.task;
            task.computeSourceShow = task.compute_source == 'GPU' ? 'CPU/GPU' : task.compute_source;
            task.accCardTypeShow = getListValueWithKey(ACC_CARD_TYPE, task.acc_card_type);
            task.jobTypeShow = this.$t('modelSquare.modelEvaluate');
          });
          this.limitDialogVisible = true;
        }
      });
    },
    onLimitStopSuccess() { this.prepare(); },
    onLimitDeleteSuccess() { this.fetchLimitTaskList(); this.prepare(); },
    onLimitTaskStatusUpdate(newTask) {
      const target = this.limitTaskList.find(t => t.task.id === newTask.id);
      if (target) target.task = newTask;
    },
    onLimitAllTasksTerminal() { this.prepare(); },
    prepare(repo_owner_name, repo_name) {
      this.loading = true;
      getAiTaskPrepareInfo({
        jobType: 'EVAL',
        clusterType: this.pageCfg.cluster,
        computeSource: this.pageCfg.computerResouce,
      }).then((res) => {
        res = res.data
        this.loading = false
        if (res.code === 0) {
          const data = res.data;
          this.queueNum = data.wait_count || 1;
          this.alreadyMsgBoxShow = !data.can_create_more;
          this.notStopTaskCount = data.not_stop_task_count || 0;
          this.limitCount = data.limit_count || 0;
          if (!data.can_create_more) {
            this.fetchLimitTaskList();
          } else {
            this.limitDialogVisible = false;
          }
          this.specConfigs.showPoint = data.pay_switch;
          this.specConfigs.blance = data.point_account ? data.point_account.balance : 0;
          this.specConfigs.specs = { 'all': data.specs.all };
          this.state.spec = this.specConfigs.specs['all'][0]?.id?.toString() || '';
          this.state.taskName = data.display_job_name;
        }
      }).catch(err => {
        this.loading = false;
        this.$message.error(err?.response?.data?.message || err);
      });
    },
    createBasetaskData() {
      const subObj = {
        repoOwnerName: this.DepModelInfo[0].configs.repo_owner_name,
        repoName: this.DepModelInfo[0].configs.repo_name,
        job_type: 'EVAL',
        cluster: this.pageCfg.cluster,
        compute_source: this.pageCfg.computerResouce,
      }

      const modelComputeSourceInfo = this.DepModelInfo[0].configs.compute_sources[this.pageCfg.computerResouce];
      subObj['app_name'] = this.DepModelInfo[0].name
      subObj['pretrain_model_id_str'] = this.DepModelInfo[0].configs.model_id;
      subObj['dataset_uuid_str'] = this.DepModelInfo[0].configs.dataset_uuid;

      let datasetSelectList = []
      if (this.datasets.length) {
        datasetSelectList = this.datasets.map(item => item.name)
      }

      if (this.pageCfg.computerResouce === 'NPU') {
        modelComputeSourceInfo.parameters.push({ label: 'datasets', value: datasetSelectList.join(' ') })
        modelComputeSourceInfo.parameters.push({ label: 'limit', value: this.limitNum })
        subObj['params'] = JSON.stringify({ parameter: modelComputeSourceInfo?.parameters });
      } else {
        subObj['boot_file'] = modelComputeSourceInfo?.boot_file;
        subObj['image_url'] = modelComputeSourceInfo?.image_url;
        let parameters = [{ label: 'datasets', value: datasetSelectList.join(' ') }, { label: 'limit', value: this.limitNum }]
        subObj['params'] = JSON.stringify({ parameter: parameters });
      }

      subObj['label_names'] = this.DepModelInfo[0].type;
      subObj['description'] = ''
      subObj['branch_name'] = 'master'
      subObj['display_job_name'] = this.state.taskName
      subObj['has_internet'] = this.specConfigs.specs.all[0].has_internet;
      subObj['spec_id'] = +this.state.spec;
      subObj['work_server_number'] = 1;
      return subObj
    },
    createDeploytaskData() {
      let { repo_owner_name, repo_name } = this.restoreModelInfo
      const subObj = {
        repoOwnerName: repo_owner_name,
        repoName: repo_name,
        job_type: 'EVAL',
        cluster: this.pageCfg.cluster,
        compute_source: this.pageCfg.computerResouce,
      }
      subObj['app_name'] = this.ftModelInfo[0].app_name + ` (${this.ftModelName})`
      subObj['pretrain_model_id_str'] = ''

      subObj['dataset_uuid_str'] = ''

      let datasetSelectList = []
      if (this.datasets.length) {
        datasetSelectList = this.datasets.map(item => item.name)
      }

      const modelComputeSourceInfo = this.restoreModelInfo.compute_sources[this.pageCfg.computerResouce];
      if (this.pageCfg.computerResouce === "GPU") {
        subObj['boot_file'] = modelComputeSourceInfo?.boot_file;
        subObj['image_url'] = modelComputeSourceInfo?.image_url;
      }
      const parameters = JSON.parse(modelComputeSourceInfo.parameters)
      parameters.push({ label: 'source_task_id', value: String(this.ftModelInfo[0].id) })
      parameters.push({ label: 'datasets', value: datasetSelectList.join(' ') })
      parameters.push({ label: 'limit', value: this.limitNum })
      subObj['label_names'] = 'chat'
      subObj['description'] = ''
      subObj['branch_name'] = 'master'
      subObj['display_job_name'] = this.state.taskName
      subObj['has_internet'] = this.specConfigs.specs.all[0].has_internet;
      subObj['spec_id'] = +this.state.spec;
      subObj['work_server_number'] = 1;

      subObj['source_cloudbrain_id'] = this.ftModelInfo[0].id;
      subObj['params'] = JSON.stringify({ parameter: parameters });

      return subObj
    },
    submit() {
      if (this.maskLoading) return;
      let canSubmit = true;
      for (let key in this.state) {
        if (this.$refs[key + 'Ref']) {
          if (!this.$refs[key + 'Ref'].check()) {
            canSubmit = false
          }
        }
      }
      if (!canSubmit) return;
      let subData = {}
      if (this.modelRadio == 'finetune') {
        subData = this.createDeploytaskData()
      } else {
        subData = this.createBasetaskData()
      }

      this.maskLoadingContent = this.$t('cloudbrainObj.taskPrepareTips');
      this.maskLoading = true;
      this.errorMsg = '';
      this.errorMsgBoxShow = false;
      this.postCreateAiTask(subData)
    },
    postCreateAiTask(data) {
      createAiTask(data).then(res => {
        const data = res.data;
        if (data.code == 0) {
          this.$router.replace('/eval/evaluate');
        } else {
          this.maskLoading = false;
          this.errorMsg = data.msg;
          this.errorMsgBoxShow = true;
          document.querySelector('.main-content').scrollTo({ top: 0, behavior: 'smooth' });
        }
      }).catch(err => {
        this.maskLoading = false;
        this.$message.error(err)
      });
    },
    cancel() {
      if (this.$route.query.backurl) {
        window.location.href = this.$route.query.backurl;
        return;
      }
      if (this.$route.query.backpath) {
        this.$router.replace(this.$route.query.backpath);
        return;
      }
      this.$router.replace('/eval/evaluate');
    },
    changeModelType(val) {
      if (val === 'finetune') {
        console.log('finetune', this.ftModelName)
        if (!this.ftModelName) {
          this.ftModelName = this.ftModelList[0].display_job_name
        }
        this.changeFtModel('')
      } else {
        this.changeModel('')
      }
    },
    async getFtTaskList() {
      try {
        const params = {
          job_type: 'FINETUNE',
          job_status: 'SUCCEEDED',//'SUCCEEDED',
          page: 1,
          pageSize: 1000,
        };

        const response = await getMyAiTasks(params)
        const res = response.data

        if (res.code == 0) {
          const data = res.data
          if (data.tasks.length > 0) {
            data.tasks.forEach(item => {
              if (item.can_fintune_experience) { //item.can_fintune_experience
                this.ftModelList.push(item.task)
              }
            });
            console.log("this.ftModelList", this.ftModelList)
          }
        } else {
          throw new Error(res.message || '请求失败')
        }
      } catch (error) {
        console.error('获取任务列表失败:', error)
        throw error
      }
    },
    async getPromoteExperience(modelName, computerResouce) {
      try {
        const res = await getPromoteData('model/modelexperiencenew.json')
        const data = JSON.parse(res.data);
        console.log("getPromoteExperience", data)
        console.log("getPromoteExperience modelName", modelName, computerResouce)
        this.modelList = data.filter(item => item.type == 'chat');
        if (modelName && data.filter(item => item.name == modelName).length) {
          this.modelName = modelName;
        } else {
          if (data.length) {
            this.modelName = data[0].name;
          }
        }
        this.changeModel('', true);
        if (computerResouce && this.pageCfg.computerResouces.filter(item => item.k == computerResouce).length) {
          this.pageCfg.computerResouce = computerResouce;
        }
        this.prepare();
      } catch (error) {
        this.$message.error(error)
        throw error
      }
    },
    async getPromoteFinetune() {
      try {
        const res = await getPromoteData('model/modelfinetune.json')
        const data = JSON.parse(res.data);
        this.promoteFtList = data.llm.model
        // this.modelList = data;
        console.log("promoteFtList", this.promoteFtList)
        this.changeFtModel('')
      } catch (error) {
        this.$message.error(error)
        throw error
      }
    },
    async restoreFtModelExperience(jobName) {
      this.ftModelName = jobName
      await this.initData()
      this.ftModelInfo = this.ftModelList.filter((item) => {
        return item.display_job_name === this.ftModelName;
      });
    },
    async initData(modelName, computerResouce) {
      this.globalLoading = true;
      await this.getFtTaskList()
      await this.getPromoteExperience(modelName, computerResouce)
      await this.getPromoteFinetune()
      this.globalLoading = false;
    },
    getEvalPromoteData(modelName, computerResouce) {
      this.dataLoading = true;
      getPromoteData('model/evaldataset.json').then(res => {
        this.dataLoading = false;
        try {
          console.log("saddddddddddd", res.data)
          const data = JSON.parse(res.data);

          console.log("sddddddddddddddddddd", data)
          this.datasetList = data;

        } catch (err) {
          this.dataLoading = false;
          this.$message.error(err);
        }
      })
    },
  },
  beforeMount() {
    const isLogin = !!document.querySelector('meta[name="_uid"]');
    if (isLogin) {
      this.loginName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext');
    }
    this.getEvalPromoteData()
  },
  async mounted() {
    this.urlParams = getUrlSearchParams();
    const computerResouce = this.urlParams.compute_resource;
    let modelName = this.urlParams.modelName
    if (modelName) {
      modelName = decodeURIComponent(modelName)
    }
    if (this.urlParams.job_name) {
      this.modelRadio = 'finetune'
      await this.restoreFtModelExperience(this.urlParams.job_name)
    } else {
      this.initData(modelName, computerResouce)
    }

  },
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';
@import '../../components/createcommon.less';

.option-content {
  height: 30px;
  display: flex;
  align-items: center;
  border-bottom: 1px solid rgb(232, 232, 232);

  .option-name {
    flex: 1;
    color: #101010;
    font-weight: bold;
    margin-left: 14px;
  }

  .option-label {
    display: flex;
    align-items: center;

    .label-text {
      color: #606266;
      font-weight: 400;
      margin-right: 10px;
    }

  }
}

.el-select {
  /deep/ .el-checkbox__label {
    display: none;
  }
}

.data-eval {
  /deep/ .el-select-dropdown {
    padding: 10px 0;
  }
}
</style>
