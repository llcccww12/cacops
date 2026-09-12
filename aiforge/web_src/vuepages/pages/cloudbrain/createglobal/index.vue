<template>
  <div class="ai-task-create-global-c">
    <div class="create-task-warp">
      <div class="ui container" v-if="errorMsgBoxShow">
        <div class="err-msg-box">
          <p v-html="errorMsg"></p>
        </div>
      </div>
      <div class="area-l-title">
        <a href="/cloudbrains" class="link label">{{ $t('notebook.sameTaskTips6') }}</a>
        <div class="separator"> / </div>
        <div class="label">{{ $t('cloudbrainObj.createNewAitask') }}</div>
      </div>
      <div class="create-t-box">
        <div class="form-container">
          <div class="form-content" :class="showFormRight ? '' : 'hidden-right'">
            <div class="form-left">
              <div class="form-body">
                <div class="form-row" v-if="this.isUseTmpl">
                  <div class="left-area">
                    <div class="title align-items-center">{{ $t('taskTmplObj.useTaskTmpl') }}</div>
                    <div class="content" style="display:flex;align-items:center;">
                      <span style="color:rgb(0, 122, 255)">{{ useTmplData.Name }}</span>
                    </div>
                  </div>
                </div>
                <TaskTypeSelect v-model="taskType" :limitCount='limitCount' @change="changeTaskType"></TaskTypeSelect>
                <div class="form-body-content" style="display:flex" v-if="taskType === 'HPC'">
                  <div class="main-title">{{ $t('cloudbrainObj.selectApp') }}：</div>
                  <div class="hpc-app-list">
                    <el-tooltip class="item-tool-tips" effect="light" placement="bottom">
                      <div class="app-item">
                        <img src="/img/supercompute/mmlspark.png" alt="">
                      </div>
                      <div slot="content" class="item-tool-tip-content" v-html="$t('superComputeObj.mmlSparkDescr')">
                      </div>
                    </el-tooltip>

                  </div>
                </div>
                <div class="form-body-content">
                  <div class="main-title">{{ $t('cloudbrainObj.basicInfo') }}：</div>
                  <FormTopV2 ref="formTopRef" :taskTypeObj="taskTypeObj" :queueNum="queueNum"
                    @change="changeClusterAndComputeResouce">
                  </FormTopV2>
                  <NetworkType ref="networkTypeRef" v-if="formCfg.networkType" v-model="state.networkType">
                  </NetworkType>
                  <Visualization ref="visualizationRef" v-if="formCfg.visualization" v-model="state.visualizeRequired">
                  </Visualization>
                  <SpecSelect ref="specRef" v-if="formCfg.spec" v-model="state.spec" :required="formCfg.spec.required"
                    :configs="specConfigs" :workServerNum="state.workServerNum" :networkType="state.networkType"
                    :visualize="state.visualizeRequired">
                  </SpecSelect>
                  <WorkServerNum ref="workServerNumRef" v-if="formCfg.workServerNum && workServerNumList.length > 1"
                    v-model="state.workServerNum" :required="formCfg.workServerNum.required" :data="workServerNumList">
                  </WorkServerNum>
                  <TaskName ref="taskNameRef" v-if="formCfg.taskName" v-model="state.taskName" autofocus
                    :required="formCfg.taskName.required" :userName="repoOwnerName">
                  </TaskName>
                  <TaskDescr ref="taskDescrRef" v-if="formCfg.taskDescr" v-model="state.taskDescr"
                    :required="formCfg.taskDescr.required"></TaskDescr>
                </div>
                <div class="line"></div>
                <div class="form-body-content">
                  <div class="main-title params-setting">{{ $t('cloudbrainObj.paramsSetting') }}：</div>
                  <ImageSelectV1 ref="imagev1Ref" v-if="formCfg.imagev1" v-model="state.image" :configs="taskTypeObj"
                    :spec="state.spec" :required="formCfg.imagev1.required"
                    :type="formCfg.imagev1.type != undefined ? formCfg.imagev1.type : 0" :useId="formCfg.imagev1.useId"
                    :networkType="state.networkType" :visualize="state.visualizeRequired" @changeImage="changeImage">
                  </ImageSelectV1>
                  <DatasetSelect ref="datasetRef" v-if="formCfg.dataset" v-model="state.dataset"
                    :required="formCfg.dataset.required" :multiple="formCfg.model.multiple" :maxCount="datasetCount"
                    :useExceedSize="formCfg.dataset.useExceedSize || false" :exceedSize="datasetSize">
                  </DatasetSelect>
                  <ModelSelect ref="modelRef" v-if="formCfg.model" v-model="state.model" :maxCount="modelCount"
                    :required="formCfg.model.required" :multiple="formCfg.model.multiple" :exceedSize="modelSize"
                    :useExceedSize="formCfg.model.useExceedSize">
                  </ModelSelect>
                  <RepoSelect ref="repoRef" v-if="formCfg.repo" v-model="state.repoList" :repoSize="repoSize"
                    :required="formCfg.repo.required" :multiple="formCfg.repo.multiple">
                  </RepoSelect>
                  <BranchName v-loading="branchLoading" ref="branchNameRef" v-if="formCfg.branchName"
                    v-model="state.branchName" :required="formCfg.branchName.required" :branches="branchList">
                  </BranchName>
                  <BootFile ref="bootFileRef" v-if="formCfg.bootFile" v-model="state.bootFile"
                    :required="formCfg.bootFile.required" :sampleUrl="formCfg.bootFile.sampleUrl"></BootFile>
                  <RunParameters ref="runParametersRef" v-if="formCfg.runParameters" v-model="state.runParameters"
                    :required="formCfg.runParameters.required">
                  </RunParameters>
                  <RunTimeLimit ref="runTimeLimitRef" v-if="formCfg.runTimeLimit && subscriberUser"
                    v-model="state.time_limit" :required="formCfg.runTimeLimit.required"></RunTimeLimit>
                  <SelfSshAddress ref="selfSshAddressRef" v-if="formCfg.selfSshAddress && subscriberUser"
                    v-model="state.endpoint_port" :required="formCfg.selfSshAddress.required"> </SelfSshAddress>
                  <ForbidPenetrationTipCheck v-model="agreeForbidPenetration"></ForbidPenetrationTipCheck>
                  <div class="form-row">
                    <div class="left-area">
                      <div class="title"></div>
                      <div class="content">
                        <el-button type="primary"
                          :disabled="!agreeForbidPenetration || maskLoading || alreadyMsgBoxShow || !specConfigs.specs[state.networkType].length"
                          size="default" class="submit-btn" @click="submit">
                          {{ $t('cloudbrainObj.createTask') }}
                        </el-button>
                        <el-button class="cancel-btn" size="default" @click="cancel">{{ $t('cancel') }}</el-button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="form-right">
              <div class="form-right-content" v-show="showFormRight">
                <SDKCode ref="sdkCodeRef" :data="state" :pageConfigs="taskTypeObj" :formConfigs="formCfg"
                  :specConfigs="specConfigs"></SDKCode>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <LoadingMask :loading="maskLoading" :content="maskLoadingContent"></LoadingMask>
    <DialogTips :visible="visible" @closeDialog="visible = false"></DialogTips>
    <!-- 新建任务限制提示弹框 -->
    <TaskLimitDialog :visible="notebookMsgBoxShow" :title="limitDialogTitle" :desc="limitDialogDesc"
      :taskList="notebookTaskList" @close="closeNotebookDialog" @stopSuccess="onStopSuccess"
      @deleteSuccess="onDeleteSuccess" @taskStatusUpdate="onTaskStatusUpdate" @allTasksTerminal="onAllTasksTerminal">
    </TaskLimitDialog>
  </div>
</template>

<script>

import TaskTypeSelect from '~/components/cloudbrain/TaskTypeSelect.vue';
import FormTopV2 from '~/components/cloudbrain/FormTopV2.vue';
import TaskName from '~/components/cloudbrain/TaskName.vue';
import TaskDescr from '~/components/cloudbrain/TaskDescr.vue';
import RepoSelect from '~/components/cloudbrain/RepoSelectV2.vue';
import BranchName from '~/components/cloudbrain/BranchName.vue';
import ModelSelect from '~/components/cloudbrain/ModelSelectV2.vue';
import ImageSelectV1 from '~/components/cloudbrain/ImageSelectV1.vue';
import BootFile from '~/components/cloudbrain/BootFile.vue';
import DatasetSelect from '~/components/cloudbrain/DatasetSelectV2.vue';
import RunParameters from '~/components/cloudbrain/RunParameters.vue';
import NetworkType from '~/components/cloudbrain/NetworkType.vue';
import SpecSelect from '~/components/cloudbrain/SpecSelect.vue';
import Visualization from '~/components/cloudbrain/Visualization.vue';
import WorkServerNum from '~/components/cloudbrain/WorkServerNum.vue';
import RunTimeLimit from '~/components/cloudbrain/RunTimeLimit.vue';
import SelfSshAddress from '~/components/cloudbrain/SelfSshAddress.vue';
import SDKCode from '~/components/cloudbrain/SDKCode.vue';
import DialogTips from '~/components/cloudbrain/DialogTips.vue';
import LoadingMask from '~/components/cloudbrain/LoadingMask.vue';
import TaskLimitDialog from '~/components/cloudbrain/TaskLimitDialog.vue';
import ForbidPenetrationTipCheck from '~/components/cloudbrain/ForbidPenetrationTipCheck.vue';

import { JOB_TYPE, ACC_CARD_TYPE } from '~/const';
import { configCreateManager } from '../configs';
import { getUrlSearchParams, initClipboard, toBoolean, renderSpecObject, getListValueWithKey } from '~/utils';
import { getAiTaskPrepareInfo, createAiTask, getAiTask, getMyAiTasks } from '~/apis/modules/cloudbrain';
import { getAiTaskTmpl } from '~/apis/modules/aitasktmpl';
import { getRepoBranch } from "~/apis/modules/repos";
import { getStaticFile, getPromoteData, getMarkdownHtml } from '~/apis/modules/common';
import { getCheckoldDataset } from "~/apis/modules/dataset";
import SparkMD5 from "spark-md5";
import hljs from 'highlight.js';
import { lang, i18n } from '~/langs';
import dayjs from 'dayjs';

export default {
  data() {
    return {
      codeUsePromotePath: `tips/cloudbrain/sdkcode${lang == 'zh-CN' ? '' : '_en'}.md`,
      codeUseContent: '',
      taskType: '',
      taskTypeObj: {},

      formCfg: {},
      cancelUrl: '',
      repoOwnerName: '',
      repoName: '',

      state: {
        taskName: '',
        taskDescr: '',
        branchName: '',
        repoList: [],
        model: [],
        image: {
          image_url: '',
          image_id: '',
          image_name: '',
        },
        bootFile: '',
        dataset: [],
        runParameters: [],
        networkType: 'has_internet',
        visualizeRequired: false,
        spec: '',
        workServerNum: 1,
        time_limit: '4',
        endpoint_port: { endPoint: '', port: '' },
      },
      branchList: [],
      specConfigs: {
        specs: {
          'all': [],
          'no_internet': [],
          'has_internet': [],
        },
        blance: 0,
        showPoint: false,
      },
      workServerNumList: [1],
      queueNum: 1,
      agreeForbidPenetration: true,

      errorMsgBoxShow: false,
      errorMsg: '',
      alreadyMsgBoxShow: false,
      notStopTaskCount: 0,
      limitCount: 0,

      // 调试任务相关限制
      notebookCanCreateMore: true,
      notebookLimitCount: 0,
      notebookMsgBoxShow: false,
      notebookTaskList: [],
      maskLoading: false,
      maskLoadingContent: '',
      datasetSize: 0,
      datasetCount: 0,
      modelSize: 0,
      modelCount: 0,
      repoSize: 0,
      noSpecFlag: false,
      visible: false,
      showFormRight: true,
      subscriberUser: false,
      branchLoading: false,
      appName: '',
      modify: false,
      modifyTask: {},
      urlParams: {},
      firstModifyFlag: true,

      isUseTmpl: false,
      isFillTmplInfo: false,
      useTmplData: {},
      tmplTips: {
        spec: '',
        image: '',
        branch: '',
      },

      useSpecData: null,
    };
  },
  components: {
    TaskTypeSelect, RepoSelect, FormTopV2, TaskName, TaskDescr, BranchName, BootFile, ImageSelectV1,
    ModelSelect, DatasetSelect, RunParameters, NetworkType, Visualization, SpecSelect, WorkServerNum,
    RunTimeLimit, SelfSshAddress, SDKCode, LoadingMask, DialogTips, TaskLimitDialog, ForbidPenetrationTipCheck,
  },
  watch: {
    'state.repoList': {
      handler(newVal, oldVal) {
        // 只有当数组内容实际发生变化时才执行
        if (JSON.stringify(newVal) !== JSON.stringify(oldVal)) {
          if (newVal.length > 0) {
            let branch = newVal[0]?.branch ?? false;
            this.getRepoBranchList(newVal[0].owner_name, newVal[0].name, branch);
          } else {
            this.state.branchName = ''
            this.repoOwnerName = ''
            this.repoName = ''
            this.branchList = []
          }
        }
      },
      deep: true
    }
  },
  computed: {
    canCreateMore() {
      return this.notStopTaskCount < this.limitCount;
    },
    limitDialogTitle() {
      return this.$t('notebook.limitTitle', { type: this.getTaskTypeName(this.taskType) });
    },
    limitDialogDesc() {
      if (!this.canCreateMore && !this.notebookCanCreateMore) {
        return this.$t('notebook.limitReason', { count: this.notebookLimitCount, runCount: this.limitCount });
      } else if (!this.canCreateMore) {
        return this.$t('notebook.stopReason', { count: this.limitCount });
      } else {
        return this.$t('notebook.deleteReason', { count: this.notebookLimitCount, type: this.getTaskTypeName(this.taskType) });
      }
    },
  },
  methods: {
    dateFormat(unix) {
      return dayjs(unix * 1000).format('YYYY-MM-DD HH:mm:ss');
    },
    // 获取受限任务列表
    fetchTaskListForLimit() {
      const isDebugTask = this.taskType === 'DEBUG';
      // DEBUG任务且数量超限时显示所有状态，其他情况只显示运行/等待状态
      const jobStatus = (isDebugTask && !this.notebookCanCreateMore) ? '' : 'nostop';
      let params = {
        job_type: isDebugTask ? 'DEBUG' : this.taskType,
        job_status: jobStatus,
        ai_center: '',
        cluster: '',
        compute_source: '',
        q: '',
        page: 1,
        pageSize: 5
      };

      getMyAiTasks(params).then(res => {
        if (res.data.code === 0) {
          this.notebookTaskList = res.data.data.tasks || [];
          this.notebookTaskList.forEach((item) => {
            const task = item.task;
            task.computeSourceShow = task.compute_source == 'GPU' ? 'CPU/GPU' : task.compute_source;
            task.accCardTypeShow = getListValueWithKey(ACC_CARD_TYPE, task.acc_card_type);
            task.jobTypeShow = getListValueWithKey(JOB_TYPE, task.job_type);
          })
          this.notebookMsgBoxShow = true;
        }
      });
    },

    // 重新检查创建任务限制
    recheckCreationLimits() {
      if (!this.taskTypeObj.cluster || !this.taskTypeObj.computerResouce) return;
      getAiTaskPrepareInfo({
        jobType: this.taskType,
        clusterType: this.taskTypeObj.cluster,
        computeSource: this.taskTypeObj.computerResouce,
        appName: this.appName
      }).then(res => {
        res = res.data;
        if (res.code == 0) {
          const data = res.data;
          this.alreadyMsgBoxShow = !data.can_create_more || data.notebook_can_create_more === false;
          this.notebookCanCreateMore = data.notebook_can_create_more !== false;
          this.notebookLimitCount = data.notebook_limit_count || 0;
          this.notStopTaskCount = data.not_stop_task_count || 0;
          this.limitCount = data.limit_count || 0;
          // 检查是否还需要弹框
          if (data.can_create_more && data.notebook_can_create_more !== false) {
            this.notebookMsgBoxShow = false;
          }
        }
      });
    },

    // 关闭弹框
    closeNotebookDialog() {
      this.notebookMsgBoxShow = false;
    },

    // 获取任务类型的中文名称
    getTaskTypeName(type) {
      return getListValueWithKey(JOB_TYPE, type) || type;
    },

    // TaskLimitDialog 停止成功回调
    onStopSuccess(task) {
      const isDebugTask = this.taskType === 'DEBUG';
      if (isDebugTask) {
        this.fetchTaskListForLimit();
        this.recheckCreationLimits();
      } else {
        this.recheckCreationLimits();
        this.notebookMsgBoxShow = false;
      }
    },

    // TaskLimitDialog 删除成功回调
    onDeleteSuccess(task) {
      this.fetchTaskListForLimit();
      this.recheckCreationLimits();
    },

    // TaskLimitDialog 状态更新回调
    onTaskStatusUpdate(newTask) {
      const target = this.notebookTaskList.find(t => t.task.id === newTask.id);
      if (target) {
        target.task = newTask;
      }
    },

    // TaskLimitDialog 所有轮询任务都到达终态回调
    onAllTasksTerminal() {
      const isDebugTask = this.taskType === 'DEBUG';
      if (isDebugTask) {
        this.fetchTaskListForLimit();
        this.recheckCreationLimits();
      } else {
        this.recheckCreationLimits();
      }
    },

    // step 1
    sparkMD5Hash(str = '') {
      return SparkMD5.hash(str) + Math.random().toString().replace('0.', '');
    },
    hljsAndInsertCopyButton(htmlStr) {
      const html = document.createElement('div');
      html.innerHTML = htmlStr;
      const codeBlocks = html.querySelectorAll('.code-block')
      for (let i = 0, iLen = codeBlocks.length; i < iLen; i++) {
        const codeBlockI = codeBlocks[i];
        const txt = codeBlockI.textContent;
        const code = codeBlockI.querySelector('code').textContent;
        codeBlockI.querySelector('code').innerHTML = hljs.highlight('python', code).value;
        const copyBtn = document.createElement('div');
        copyBtn.classList = ['copy-btn'];
        copyBtn.innerHTML = `<a href="javascript:;" class="ui poping inline up clipboard" id="clipboard-${this.sparkMD5Hash(txt)}"
        data-position="top center" data-variation="inverted tiny" data-success="${this.$t('copySuccess')}"
        data-content="${this.$t('copy')}" data-original="${this.$t('copy')}" 
        data-clipboard-text=""><i style="font-size:14px;" class="copy outline icon"></i></a>`;
        copyBtn.querySelector('a').setAttribute('data-clipboard-text', txt);
        codeBlockI.appendChild(copyBtn);
      }
      return html.innerHTML;
    },
    getMarkdown(str) {
      getMarkdownHtml(str).then(res => {
        const html = res.data;
        this.codeUseContent = this.hljsAndInsertCopyButton(html);
        this.$nextTick(() => {
          initClipboard('.code-use-guide .clipboard');
        });
      }).catch(err => {
        console.log(err);
      });
    },
    getCodeUseGuide() {
      if (this.codeUseContent) return;
      getPromoteData(this.codeUsePromotePath).then(res => {
        this.loading = false;
        let contentStr = res.data;
        if (contentStr) {
          this.getMarkdown(contentStr);
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    async getRepoBranchList(repoOwnerName, repoName, branch) {
      try {
        this.branchLoading = true;
        this.repoOwnerName = repoOwnerName
        this.repoName = repoName
        const response = await getRepoBranch({ repoOwnerName, repoName });
        const res = response.data;
        if (res.code === 0) {
          this.branchList = res.data.branches
          if (branch) {
            this.state.branchName = branch
          } else {
            this.state.branchName = res.data.default_branch
          }

        }

      } catch (err) {
        console.log(err);
      } finally {
        this.branchLoading = false;
      }
    },
    async changeTaskType(type, typeObj) {
      if (typeObj) {
        if (!(this.modify && this.modifyTask.taskType === type)) {
          this.resetForm()
        }
        // 从仓库内新建
        if (this.urlParams.id && this.urlParams.onwer && this.urlParams.alias) {
          this.$nextTick(() => {
            this.state.repoList = [{
              id: this.urlParams.id,
              owner_name: this.urlParams.onwer,
              alias: this.urlParams.alias,
              name: this.urlParams.name,
            }]
          });
        }
        let taskObj = JSON.parse(JSON.stringify(typeObj));
        this.taskTypeObj = { taskType: type, ...taskObj }
        if (this.modify && this.modifyTask.taskType === type) {
          this.taskTypeObj.cluster = this.modifyTask.cluster
          this.taskTypeObj.computerResouce = this.modifyTask.computerResouce
        }
        if (this.isUseTmpl && this.firstModifyFlag) {
          this.taskTypeObj.cluster = this.useTmplData.Cluster
          this.taskTypeObj.computerResouce = this.useTmplData.ComputeSource
        }
        if (this.firstModifyFlag && this.useSpecData && configCreateManager.getResourceConfig(this.taskType, this.useSpecData.ComputeSource)) {
          this.taskTypeObj.computerResouce = this.useSpecData.ComputeSource
        } else {
          this.useSpecData = null
        }
        this.changeClusterAndComputeResouce({
          cluster: this.taskTypeObj.cluster,
          computerResouce: this.taskTypeObj.computerResouce,
        });
      }
    },
    // step 2
    resetForm() {
      const newState = {
        taskName: '',
        taskDescr: '',
        branchName: '',
        repoList: [],
        model: [],
        image: {
          image_url: '',
          image_id: '',
          image_name: '',
        },
        bootFile: '',
        dataset: [],
        runParameters: [],
        networkType: 'has_internet',
        visualizeRequired: false,
        spec: '',
        workServerNum: 1,
        time_limit: '4',
        endpoint_port: { endPoint: '', port: '' },
      };
      Object.keys(this.state).forEach(key => {
        if (key in newState) {
          this.state[key] = newState[key];
        }
      });
      this.branchList = [];
      this.specConfigs = {
        specs: {
          'all': [],
          'no_internet': [],
          'has_internet': [],
        },
        blance: 0,
        showPoint: false,
      };
      this.workServerNumList = [1];
      this.queueNum = 1;
      this.errorMsgBoxShow = false;
      this.errorMsg = '';
      this.alreadyMsgBoxShow = false;
      this.notStopTaskCount = 0;
      this.limitCount = 0;
      this.maskLoading = false;
      this.maskLoadingContent = '';
      this.datasetSize = 0;
      this.datasetCount = 0;
      this.modelSize = 0;
      this.modelCount = 0;
      this.repoSize = 0
      this.noSpecFlag = false;
      this.visible = false;
    },
    async changeClusterAndComputeResouce(obj) {
      console.log('changeClusterAndComputeResouce', obj)
      this.taskTypeObj = {
        ...this.taskTypeObj,
        ...obj
      };
      const configs = configCreateManager.getResourceConfig(this.taskType, obj.computerResouce)
      if (configs === null) {
        console.warn('get page configs error', this.taskType, obj.cluster, obj.computerResouce);
        return;
      }
      this.formCfg = configs.form
      this.appName = configs.appName ? configs.appName : '';
      // this.resetForm();
      this.maskLoadingContent = this.$t('cloudbrainObj.dataPreparing');
      this.maskLoading = true;
      if (this.modify && this.taskType === this.modifyTask.taskType && !this.firstModifyFlag) {
        this.state.image = {}
        this.state.spec = ''
        await this.getTaskInfo(this.urlParams.id)
      }

      getAiTaskPrepareInfo({
        jobType: this.taskType,
        clusterType: obj.cluster,
        computeSource: obj.computerResouce,
        appName: this.appName
      }).then(res => {
        this.firstModifyFlag = false
        res = res.data;
        this.maskLoadingContent = '';
        this.maskLoading = false;
        if (res.code == 0) {
          const data = res.data;
          this.subscriberUser = data.is_subscriber;
          if (!this.subscriberUser) {
            delete this.formCfg['runTimeLimit']
            delete this.formCfg['selfSshAddress']
          };
          if (this.branchList.length == 0) {
            this.branchList = data.branches || [];
            this.state.branchName = data.default_branch;
          }
          this.alreadyMsgBoxShow = !data.can_create_more || data.notebook_can_create_more === false;
          this.notStopTaskCount = data.not_stop_task_count || 0;
          this.limitCount = data.limit_count || 0;
          this.notebookCanCreateMore = data.notebook_can_create_more !== false;
          this.notebookLimitCount = data.notebook_limit_count || 0;

          // 根据限制类型决定显示哪个弹框
          if (!data.can_create_more || !data.notebook_can_create_more) {
            this.fetchTaskListForLimit();
          }
          this.specConfigs.showPoint = data.pay_switch;
          this.specConfigs.blance = data.point_account ? data.point_account.balance : 0;
          this.specConfigs.specs = data.specs || {
            'all': [],
            'no_internet': [],
            'has_internet': [],
          };
          this.state.networkType = this.formCfg.networkType ? 'no_internet' : 'all';
          this.state.visualizeRequired = false;
          if (this.state.networkType == 'no_internet'
            && !this.specConfigs.specs['no_internet'].length
            && this.specConfigs.specs['has_internet'].length) {
            this.state.networkType = 'has_internet';
          }
          if (this.modify && this.taskType === this.modifyTask.taskType && this.taskTypeObj.computerResouce === this.modifyTask.computerResouce) {
            this.state.networkType = this.modifyTask.networkType || 'all';
            this.state.visualizeRequired = this.modifyTask.visualizeRequired;
            this.$nextTick(() => { this.state.spec = ''; });
            const specs = this.specConfigs.specs[this.state.networkType] || [];

            if (specs?.length) {
              const targetSpec = specs.find(item => item.source_spec_id === this.modifyTask.specId);
              if (targetSpec) {
                this.$nextTick(() => { this.state.spec = targetSpec.id.toString(); });
                // this.state.spec = targetSpec.id.toString();
              }
            }
            if (this.modifyTask.imageId || this.modifyTask.imageUrl) {
              this.state.image.image_id = this.modifyTask.imageId;
              this.state.image.image_name = this.modifyTask.imageName;
              this.state.image.image_url = this.modifyTask.imageUrl;
            }
          } else {
            this.state.spec = this.specConfigs.specs[this.state.networkType][0] ? this.specConfigs.specs[this.state.networkType][0].id.toString() : '';
          }

          this.queueNum = data.wait_count || 1;
          this.repoSize = data.code_size_limit;
          this.state.taskName = data.display_job_name;
          if (data.config && data.config.dataset_max_num) {
            this.datasetSize = data.config.dataset_max_size;
            this.datasetCount = data.config.dataset_max_num;
          }
          if (data.config && data.config.model_max_num) {
            this.modelSize = data.config.model_max_size;
            this.modelCount = data.config.model_max_num;
          }
          this.workServerNumList = data.allowed_worker_num || [1];
          if (this.workServerNumList.length > 1) {
            this.state.workServerNum = this.workServerNumList[0];
          }
          const isCloseOnlineTips = localStorage.getItem("isCloseOnlineTips");
          if (this.taskType === "ONLINEINFERENCE" && !JSON.parse(isCloseOnlineTips)) {
            this.visible = true;
          }
          if (this.isUseTmpl && !this.isFillTmplInfo) {
            setTimeout(() => {
              this.fillTmplInfo();
            }, 10);
          }
          if (this.useSpecData) {
            setTimeout(() => {
              this.fillSpecInfo();
            }, 10);
          }
        } else { }
      }).catch(err => {
        this.maskLoadingContent = '';
        this.maskLoading = false;
        console.log(err);
      });
    },
    changeImage() {
      if (this.modify && this.taskType === this.modifyTask.taskType && this.taskTypeObj.computerResouce === this.modifyTask.computerResouce) return;
      this.state.image.image_url = '';
      this.state.image.image_id = '';
      this.state.image.image_name = '';
    },
    fillTmplInfo() {
      // console.log('fillTmplInfo', this.useTmplData);
      const tmplData = this.useTmplData;
      this.state.visualizeRequired = !!tmplData.VisualizeRequired;
      this.state.networkType = tmplData.HasInternet == 1 ? 'no_internet' : tmplData.HasInternet == 2 ? 'has_internet' : 'has_internet';
      this.$nextTick(async () => {
        const findSpec = this.specConfigs.specs[this.state.networkType].find(item => {
          return item.acc_cards_num == tmplData.AccCardsNum
            && item.acc_card_type == tmplData.AccCardType
            && item.cpu_cores == tmplData.CpuCores
            && item.mem_gi_b == tmplData.MemGiB
            && item.gpu_mem_gi_b == tmplData.GPUMemGiB
            && item.share_mem_gi_b == tmplData.ShareMemGiB
        });
        if (findSpec && findSpec.id.toString() != this.state.spec) {
          this.state.spec = findSpec.id.toString();
          this.tmplTips.spec = '';
        }
        if (!findSpec) {
          const specObj = renderSpecObject({
            id: '',
            compute_resource: tmplData.ComputeSource,
            acc_cards_num: tmplData.AccCardsNum,
            acc_card_type: tmplData.AccCardType,
            cpu_cores: tmplData.CpuCores,
            mem_gi_b: tmplData.MemGiB,
            gpu_mem_gi_b: tmplData.GPUMemGiB,
            share_mem_gi_b: tmplData.ShareMemGiB,
          }, false);
          this.state.spec = '';
          this.tmplTips.spec = tmplData.AccCardType ? this.$t('taskTmplObj.unAvailableTmplSpec', { msg: specObj.specStr }) : '';
        }
        if (tmplData.RepoOwnerName && tmplData.RepoName) {
          this.state.repoList = [{
            owner_name: tmplData.RepoOwnerName,
            alias: tmplData.RepoName,
            name: tmplData.RepoName,
            branch: tmplData.BranchName,
          }]
        }
        if (tmplData.BootFile) {
          this.state.bootFile = tmplData.BootFile;
        }
        if (tmplData.DatasetLists) {
          const datasetList = tmplData.DatasetLists;
          const dataset = [];
          for (let i = 0, iLen = datasetList.length; i < iLen; i++) {
            const _dataset = datasetList[i];
            if (!_dataset.IsDeleted) {
              dataset.push({
                id: _dataset.DatasetID,
                name: _dataset.DatasetName,
                owner_name: _dataset.OwnerName,
                alias: _dataset.DatasetAlias,
              });
            }
          }
          this.state.dataset = dataset;
        }
        if (tmplData.ModelLists) {
          const modelList = [];
          const models = tmplData.ModelLists;
          for (let i = 0, iLen = models.length; i < iLen; i++) {
            const _model = models[i];
            if (_model.IsDeleted) continue;
            modelList.push({
              id: _model.ModelID,
              name: _model.ModelName,
              owner_name: _model.OwnerName,
              alias: _model.ModelAlias,
            });
          }
          this.state.model = modelList;
        }
        if (tmplData.Parameters) {
          const runParameters = JSON.parse(tmplData.Parameters || '[]').map(item => {
            return {
              label: item.Label || item.label,
              value: item.Value || item.value,
            }
          });
          this.state.runParameters = runParameters;
        }
        this.$nextTick(() => {
          this.state.image.image_id = tmplData.ImageID;
          this.state.image.image_name = tmplData.ImageName;
          this.state.image.image_url = tmplData.ImageUrl;
        });
        this.isFillTmplInfo = true;
      });
    },
    fillSpecInfo() {
      if (!this.useSpecData) return;
      const specData = this.useSpecData;
      const findSpecs = this.specConfigs.specs['all'].filter(item => {
        return item.acc_cards_num == specData.AccCardsNum
          && item.acc_card_type == specData.AccCardType
          && item.cpu_cores == specData.CpuCores
          && item.mem_gi_b == specData.MemGiB
          && item.gpu_mem_gi_b == specData.GPUMemGiB
          && item.share_mem_gi_b == specData.ShareMemGiB
      });
      const hasInternetSpecs = findSpecs.filter(item => item.has_internet == 2);
      if (hasInternetSpecs.length || findSpecs.length) {
        const useSpec = hasInternetSpecs[0] || findSpecs[0];
        this.state.networkType = useSpec.has_internet == 1 ? 'no_internet' : useSpec.has_internet == 2 ? 'has_internet' : 'has_internet';
        this.$nextTick(async () => {
          if (useSpec && useSpec.id.toString() != this.state.spec) {
            this.state.spec = useSpec.id.toString();
          }
        });
      } else {
        this.state.spec = '';
        if (!this.specConfigs.specs['has_internet'].length && this.specConfigs.specs['no_internet'].length) {
          this.state.networkType = 'no_internet';
          this.$nextTick(async () => {
            this.state.spec = '';
          })
        }
      }
      this.useSpecData = null;
    },
    submit() {
      if (this.maskLoading) return;
      let canSubmit = true;
      for (let key in this.formCfg) {
        if (this.$refs[key + 'Ref']) {
          if (!this.$refs[key + 'Ref'].check()) {
            canSubmit = false;
          }
        }
      }
      if (!canSubmit) return;
      const subObj = {
        repoOwnerName: this.repoOwnerName || '-',
        repoName: this.repoName || '-',
        job_type: this.taskType,
        cluster: this.taskTypeObj.cluster,
        compute_source: this.taskTypeObj.computerResouce,
      };
      if (this.taskType === 'HPC') {
        subObj.app_name = this.appName
      }
      if (this.isUseTmpl) {
        subObj.template_id = this.useTmplData.ID;
      }
      for (let key in this.formCfg) {
        switch (key) {
          case 'taskName':
            subObj['display_job_name'] = this.state.taskName;
            break;
          case 'taskDescr':
            subObj['description'] = this.state.taskDescr;
            break;
          case 'branchName':
            subObj['branch_name'] = this.state.branchName;
            break;
          case 'model':
            const modelIDStr = this.state.model.map(item => item.id).join(';');
            subObj['pretrain_model_id_str'] = modelIDStr;
            break;
          case 'imagev1':
            if (this.formCfg.imagev1.useId) {
              subObj['image_id'] = this.state.image.image_id;
              subObj['image_name'] = this.state.image.image_name;
            } else {
              if (this.taskTypeObj.computerResouce == 'GPU' && this.taskType != 'GENERAL') {
                if (this.state.image.image_url) {
                  subObj['image_url'] = this.state.image.image_url;
                } else {
                  subObj['image_id'] = this.state.image.image_id;
                }
              } else {
                subObj['image_url'] = this.state.image.image_url;
                subObj['image_id'] = this.state.image.image_id;
              }

            }
            break;
          case 'imagev2':
            subObj['image_id'] = this.state.image.image_id;
            subObj['image_name'] = this.state.image.image_name;
            break;
          case 'bootFile':
            subObj['boot_file'] = this.state.bootFile;
            break;
          case 'dataset':
            subObj['dataset_uuid_str'] = this.state.dataset.map(item => item.id).join(';');
            break;
          case 'runParameters':
            const runParameters = this.state.runParameters;
            let params = { parameter: [] };
            if (runParameters.length) {
              params.parameter = runParameters.map(item => {
                return {
                  label: item.label,
                  value: item.value,
                }
              });
            }
            subObj['params'] = JSON.stringify(params);
            break;
          case 'networkType':
            let networkType = 0; // all
            if (this.state.networkType == 'no_internet') {
              networkType = 1;
            } else if (this.state.networkType == 'has_internet') {
              networkType = 2;
            }
            subObj['has_internet'] = networkType;
            break;
          case 'visualization':
            subObj['visualize_required'] = this.state.visualizeRequired;
            break;
          case 'spec':
            subObj['spec_id'] = Number(this.state.spec);
            break;
          case 'workServerNum':
            subObj['work_server_number'] = this.state.workServerNum;
            break;
          case 'runTimeLimit':
            subObj['time_limit'] = Number(this.state.time_limit);
            break;
          case 'selfSshAddress':
            subObj['endPoint'] = this.state.endpoint_port.endPoint
            subObj['port'] = Number(this.state.endpoint_port.port)
            break;
          default:
            break;
        }
      }
      // return;
      this.maskLoadingContent = this.$t('cloudbrainObj.taskPrepareTips');
      this.maskLoading = true;
      this.errorMsg = '';
      this.errorMsgBoxShow = false;
      createAiTask(subObj).then(res => {
        const data = res.data;
        if (data.code == 0) {
          this.goList();
        } else {
          this.maskLoading = false;
          this.errorMsg = data.msg;
          if (data.code == 2023) {
            this.errorMsg = data.msg + this.$t('cloudbrainObj.sameTaskTips2');
          }
          this.errorMsgBoxShow = true;
          this.scrollToTop();
        }
      }).catch(err => {
        this.maskLoading = false;
        console.log(err);
      });
    },
    scrollToTop() {
      document.querySelector('html').scrollTo({ top: 0, behavior: 'smooth' });
      document.querySelector('body').scrollTo({ top: 0, behavior: 'smooth' });
    },
    goList() {
      window.location.href = '/cloudbrains';
    },
    cancel() {
      if (this.cancelUrl) {
        window.location.href = this.cancelUrl;
      } else {
        const referrer = document.referrer || '';
        if (referrer.indexOf('cloudbrains/create') > -1) {
          window.location.href = '/dashboard';
        } else {
          window.history.back();
        }
      }
    },
    async getTaskInfo(taskId) {
      try {
        this.maskLoading = true;
        const response = await getAiTask({
          id: taskId
        });
        const res = response.data;
        if (res.code == 0) {
          const task = {};
          Object.assign(task, res.data.task);
          delete res.data.task;
          delete res.data.early_version_list;
          Object.assign(task, res.data);
          if (task.description) {
            this.state.taskDescr = task.description;
          }

          if (task.compute_source && task.job_type && task.cluster) {
            this.taskType = task.job_type;
            this.modifyTask.cluster = task.cluster;
            this.modifyTask.computerResouce = task.compute_source;
            this.modifyTask.taskType = task.job_type;
            this.modifyTask.specId = task?.spec?.source_spec_id;
            this.modifyTask.visualizeRequired = task?.visualize_required

            this.modifyTask.imageId = task?.image_id
            this.modifyTask.imageName = task?.image_name
            this.modifyTask.imageUrl = task?.image_url
          }

          if (task.has_internet == 0 || !!task.has_internet) {
            let networkType = 'has_internet'
            if (task.has_internet == 1) {
              networkType = 'no_internet';
            } else if (task.has_internet == 2) {
              networkType = 'has_internet'
            } else {
              networkType = 'all'
            }
            this.modifyTask.networkType = networkType;
          }
          if (task.boot_file) {
            this.state.bootFile = task.boot_file;
          }
          if (task.parameters) {
            const parameters = task.parameters.parameter || [];
            this.state.runParameters = [...parameters];
          }
          if (task.dataset_list) {
            const datasetList = task.dataset_list;
            const dataset = [];
            for (let i = 0, iLen = datasetList.length; i < iLen; i++) {
              const _dataset = datasetList[i];
              if (!_dataset.is_delete) {
                dataset.push({
                  id: _dataset.uuid,
                  name: _dataset.dataset_name,
                  owner_name: _dataset.owner_name,
                  alias: _dataset.dataset_alias,
                });
              }
            }
            const filterDataList = await this.clearOldDataset(dataset, taskId)
            this.state.dataset = filterDataList;
          }
          if (task.pretrain_model_list) {
            const modelList = [];
            const models = task.pretrain_model_list;
            for (let i = 0, iLen = models.length; i < iLen; i++) {
              const _model = models[i];
              if (_model.is_delete) continue;
              modelList.push({
                id: _model.id,
                name: _model.name,
                owner_name: _model.owner_name,
                alias: _model.alias,
              });
            }
            this.state.model = modelList;
          }
          if (task.repo_name && task.repo_owner_name) {
            this.state.repoList = [{
              id: task.repo_id,
              owner_name: task.repo_owner_name,
              alias: task.repo_alias,
              name: task.repo_name,
              branch: task.branch_name,
            }]
          }
          if (task.work_server_number) {
            if (this.workServerNumList.indexOf(task.work_server_number) >= 0) {
              this.state.workServerNum = task.work_server_number;
            }
          }
          if (task.time_limit != undefined) {
            this.state.time_limit = task.time_limit.toString();
          }
          if (task.end_point && task.port) {
            this.state.endpoint_port.endPoint = task.end_point;
            this.state.endpoint_port.port = task.port;
          }
        }
      } catch (err) {
        console.log(err);
        this.maskLoading = false;
        throw err; // 抛出错误以便 await 捕获
      }

    },
    async clearOldDataset(list, taskId) {
      try {
        // 使用 Promise.all 处理异步过滤
        let selectListCopy = [...list];
        let selectList
        const filteredItems = await Promise.all(
          selectListCopy.map(async (item) => {
            try {
              let response = await getCheckoldDataset({ uuid: item.id, source_id: taskId });
              let res = response.data;
              if (res.code === 0) {
                if (res.data.is_changed) {
                  if (!this.showMessageFlag) {
                    this.$message.error('历史数据集名称已被修改，请重新选择！');
                    this.showMessageFlag = true;
                  }
                  return null; // 返回 null 表示过滤掉
                }
                return item; // 返回 item 表示保留
              }
            } catch (error) {
              console.log(error);
              return null; // 出错时过滤掉
            }
            return null; // 默认过滤掉
          })
        );
        // 过滤掉 null 值
        selectList = filteredItems.filter(item => item !== null);
        return selectList;
      } catch (error) {
        console.log(error);
      } finally {

      }
    },
  },
  beforeMount() {

  },
  mounted() {
    this.urlParams = getUrlSearchParams();
    if (this.urlParams.backurl) {
      this.cancelUrl = this.urlParams.backurl;
    }

    //任务修改赋值
    this.modify = toBoolean(this.urlParams.modify);
    if (this.modify && this.urlParams.id) {
      this.getTaskInfo(this.urlParams.id);
    } else if (this.urlParams.tmpl) {
      getAiTaskTmpl({ id: this.urlParams.tmpl }).then(res => {
        res = res.data;
        if (res.code == 0) {
          this.useTmplData = res.data;
          this.isUseTmpl = true;
          this.taskType = this.useTmplData.JobType;
        } else {
          this.taskType = 'DEBUG';
        }
      }).catch(err => {
        console.log(err);
        this.taskType = 'DEBUG';
      });
    } else if (this.urlParams.spec) {
      this.taskType = 'DEBUG';
      try {
        this.useSpecData = JSON.parse(this.urlParams.spec);
      } catch (err) {
        console.log(err);
      }
    } else {
      this.taskType = 'DEBUG';
    }
    this.getCodeUseGuide();
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';

.ai-task-create-global-c {
  height: 100%;
  padding: 32px 40px 0 40px;
  background-color: rgba(249, 249, 249, 1);

  .create-task-warp {
    height: 100%;
    display: flex;
    flex-direction: column;

    .area-l-title {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      font-size: 18px;
      margin-bottom: 8px;

      .link {
        color: rgba(0, 102, 255, 1);
        font-weight: 400;
        cursor: pointer;
      }

      .separator {
        color: #6f76a5;
        font-family: SourceHanSansSC;
        margin-right: 10px;
      }

      .label {
        line-height: 39px;
        margin-right: 10px;
        font-family: SourceHanSansSC;
        font-weight: 500;
      }
    }

    .create-t-box {
      flex: 1;
      height: 0;

      .form-container {
        height: 100%;

        .form-content {
          display: flex;

          .form-left {
            flex: 1;
            width: 0;
            border-radius: 10px;
            background-color: rgba(255, 255, 255, 1);
            border: 1px solid rgba(157, 197, 226, 0.4);
            margin-bottom: 20px;
          }

          .form-right {
            width: 370px;
            position: relative;
            margin-left: 20px;

            .form-right-content {
              position: sticky;
              top: 10px;

              /deep/ code {
                white-space: pre-wrap;
                word-break: break-all;
              }
            }
          }
        }

        .form-body {
          max-width: 866px;
          margin: 36px auto;

          .form-body-content {
            .main-title {
              color: #101010;
              font-size: 16px;
              font-weight: 700;
              line-height: 1.28571429em;
              margin: 16px 0;
            }

            .hpc-app-list {
              border-radius: 10px;
              width: 80px;
              height: 80px;
              border: 2px solid rgb(50, 145, 248);
              padding: 2px;

              .app-item {
                border-radius: 8px;
                width: 100%;
                height: 100%;
                border: 2px solid rgba(255, 255, 255, 1);

                img {
                  height: 100%;
                  width: 100%;
                  object-fit: cover;
                }
              }

            }

          }

          .line {
            border-top: 1px solid rgba(34, 36, 38, .15);
            border-bottom: 1px solid rgba(255, 255, 255, .1);
            margin: 1rem 0 2rem;
          }
        }
      }
    }
  }
}


.code-use-guide.markdown {
  font-size: 14px;
  padding-left: 20px;
  padding-right: 20px;
  margin-top: 6px;
  padding: 2px;

  /deep/ .code-block {
    position: relative;

    .copy-btn {
      position: absolute;
      right: 12px;
      top: 12px;
    }
  }
}

.err-msg-box {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #fff6f6;
  color: #9f3a38;
  box-shadow: 0 0 0 1px #e0b4b4 inset, 0 0 0 0 transparent;
  margin: 1em 0;
  padding: 1em 1.5em;
  border-radius: 0.28571429rem;
  transition: opacity .1s ease, color .1s ease, background .1s ease, box-shadow .1s ease, -webkit-box-shadow .1s ease;

  >p {
    opacity: .85;
  }
}

.err-msg-box-already {
  margin: 1em 0;
  padding: 1em 1.5em;
  background-color: rgba(242, 113, 28, 0.05);
  border: 1px solid rgba(242, 113, 28, 1);
  border-radius: 5px;

  .msg-content {
    display: flex;
    align-items: center;

    i {
      font-size: 35px;
      color: rgba(242, 113, 28, 1);
    }

    .msg-content-tip {
      text-align: left;
      margin-left: 1rem;

      .line-1 {
        font-weight: 600;
        line-height: 2;

        span {
          color: rgba(242, 113, 28, 1);
        }
      }

      .line-2 {
        color: #939393
      }
    }
  }
}

.item-tool-tip-content {
  width: 360px;
  word-wrap: break-word;
  white-space: pre-wrap;
}

@media only screen and (max-width: 1500px) {
  .form-container {
    .form-right {
      display: none;
    }
  }
}

@media only screen and (max-width: 1200px) {
  .form-container .form-content:not(.hidden-right) .form-body {
    padding: 0 1em;
  }
}

@media screen and (max-width: 767px) {
  .ai-task-create-global-c {
    padding: 20px 0 0 0;

    .form-body {
      width: 100% !important;
      padding: 0 10px;
    }
  }

  .form-right {
    display: none;
  }
}
</style>
