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
        <div class="create-tips" v-if="!alreadyMsgBoxShow">
          <i class="ri-information-line"></i>
          <span style="margin-left:4px;">{{ $t('cloudbrainObj.runningLimit', {
            count: limitCount, taskType:
              $t('modelSquare.sftFinetuneTask')
          }) }}</span>
        </div>
        <TaskLimitDialog :visible="limitDialogVisible" :title="limitDialogTitle" :desc="limitDialogDesc"
          :taskList="limitTaskList" @close="limitDialogVisible = false" @stopSuccess="onLimitStopSuccess"
          @deleteSuccess="onLimitDeleteSuccess" @taskStatusUpdate="onLimitTaskStatusUpdate"
          @allTasksTerminal="onLimitAllTasksTerminal">
        </TaskLimitDialog>
        <div class="area">
          <div class="area-content">
            <div class="form-row">
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
            <ModelBaseDatasetSelect ref="modelBaseDatasetSelectRef" :tabindex="datasetTab" :oriData="datasetList"
              :datasetCount="datasetCount">
            </ModelBaseDatasetSelect>
            <IflySftParams v-if="modelName === 'Spark13B_0206'" ref="sftParamsRef"></IflySftParams>
            <SftParams v-else ref="sftParamsRef" :dtype="dtype"></SftParams>
            <FormTopV2 ref="formTopRef" :resourceObj="pageCfg" :queueNum="queueNum" @change="changeComputeResouce">
            </FormTopV2>
            <SpecSelect ref="specRef" v-model="state.spec" :required="true" :configs="specConfigs" networkType="all"
              :loading="loading"></SpecSelect>
            <TaskName ref="taskNameRef" v-model="state.taskName" :required="true" :userName="loginName"
              :autofocus="false"></TaskName>
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
        </div>

      </div>
    </div>
    <LoadingMask :loading="maskLoading" :content="maskLoadingContent"></LoadingMask>
  </div>
</template>

<script>
import Header from '../../components/Header.vue';
import ModelBaseDatasetSelect from './components/ModelBaseDatasetSelect.vue';
import SftParams from './SftParams.vue';
import IflySftParams from './IflySftParams.vue';
import FormTopV2 from '~/components/cloudbrain/FormTopV2.vue';
import TaskName from '~/components/cloudbrain/TaskName.vue';
import SpecSelect from '~/components/cloudbrain/SpecSelect.vue';
import LoadingMask from '~/components/cloudbrain/LoadingMask.vue';
import AcknowledgementsTips from '~/components/AcknowledgementsTips.vue';
import TaskLimitDialog from '~/components/cloudbrain/TaskLimitDialog.vue';
import ForbidPenetrationTipCheck from '~/components/cloudbrain/ForbidPenetrationTipCheck.vue';
import { getAiTaskPrepareInfo, createAiTask, getMyAiTasks } from '~/apis/modules/cloudbrain';
import { getPromoteData } from '~/apis/modules/common';
import { CLUSTERS, JOB_TYPE, ACC_CARD_TYPE } from '~/const';
import { COMPUTER_RESOURCES_TITLE } from '~/pages/cloudbrain/configs';
import { getListValueWithKey } from '~/utils';


export default {
  name: 'FinetuningSFTCreate',
  data() {
    return {
      modelList: [],
      modelName: '',
      datasetTab: 0,
      loginName: '',
      state: {
        spec: '',
        taskName: '',
        modelBaseDatasetSelect: ''
      },
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
      loading: false,
      specConfigs: {
        specs: {},
        blance: 0,
        showPoint: false,
      },
      queueNum: 1,
      repoName: '',
      repoOwnerName: '',
      maskLoading: false,
      maskLoadingContent: '',
      errorMsgBoxShow: false,
      errorMsg: '',
      alreadyMsgBoxShow: false,
      limitDialogVisible: false,
      limitTaskList: [],
      DepModelInfo: {},
      images: {},
      bootfiles: {},
      tmpBootfiles: {},
      npuImageInfo: [],
      datasetList: [],
      tmpDatasetList: [],
      notStopTaskCount: 0,
      limitCount: 0,
      dtype: '',
      datasetCount: 0,
      agreeForbidPenetration: true,
    };
  },
  components: { Header, ModelBaseDatasetSelect, SpecSelect, FormTopV2, TaskName, SftParams, LoadingMask, AcknowledgementsTips, IflySftParams, ForbidPenetrationTipCheck, TaskLimitDialog },
  computed: {
    limitDialogTitle() {
      return this.$t('cloudbrainObj.runningLimit', { count: this.limitCount, taskType: this.$t('modelSquare.sftFinetuneTask') });
    },
    limitDialogDesc() {
      return this.$t('cloudbrainObj.sameTaskTips1', { count: this.notStopTaskCount });
    },
  },
  methods: {
    columnStyle({ row, column, rowIndex, columnIndex }) {
      if (columnIndex == 0) {
        return "color:#101010";
      }
    },
    changeModel(_, noPrepare) {
      this.DepModelInfo = this.modelList.filter(item => item.name == this.modelName)[0]
      if ('dataset' in this.DepModelInfo) {
        this.datasetList = this.DepModelInfo.dataset
      } else {
        this.datasetList = this.tmpDatasetList
      }
      if ('bootfile' in this.DepModelInfo) {
        this.bootfiles = this.DepModelInfo.bootfile
      } else {
        this.bootfiles = this.tmpBootfiles
      }
      const compute_sources = Object.keys(this.DepModelInfo?.compute_sources || {})
      const useComputeResources = COMPUTER_RESOURCES_TITLE
        .filter(item => compute_sources.indexOf(item.k) >= 0)
        .map(item => ({ ...item, key: item.k, label: item.v }));
      this.pageCfg.computerResouces = useComputeResources;
      this.pageCfg.computerResouce = useComputeResources[0]?.k || '';
      this.pageCfg.cluster = this.modelName === 'Spark13B_0206' ? 'IFLYTEKTraining' : 'C2Net'
      if (noPrepare) return;
      this.prepare();
    },

    submit() {
      let canSubmit = true;
      for (let key in this.state) {
        if (this.$refs[key + 'Ref']) {
          if (!this.$refs[key + 'Ref'].check()) {
            canSubmit = false
          }
        }
      }
      if (!canSubmit) return;
      const subObj = {
        repoOwnerName: this.repoOwnerName,
        repoName: this.repoName,
        job_type: 'FINETUNE',
        cluster: this.pageCfg.cluster,
        compute_source: this.pageCfg.computerResouce,
      };
      subObj['display_job_name'] = this.state.taskName;
      subObj['description'] = '';
      subObj['branch_name'] = 'master'
      subObj['pretrain_model_id_str'] = this.DepModelInfo.id //'4239503b-4170-4ad1-9e3e-1b952b1669f2'
      subObj['app_name'] = this.DepModelInfo.name
      let datasets = this.$refs.modelBaseDatasetSelectRef.getPlatformDataset();
      subObj['dataset_uuid_str'] = datasets.map(item => item.id).join(';');
      if (this.pageCfg.computerResouce !== 'NPU') {
        subObj['image_url'] = this.DepModelInfo.compute_sources[this.pageCfg.computerResouce].image_url;
      } else {
        if (this.npuImageInfo.length > 0) {
          subObj['image_id'] = this.npuImageInfo[0].image_id;
          subObj['image_name'] = this.npuImageInfo[0].image_name;
        }
      }
      subObj['boot_file'] = this.bootfiles[this.pageCfg.computerResouce]

      let parameter = this.$refs.sftParamsRef.getParams();
      Object.keys(this.DepModelInfo.paramaters).forEach((key) => {
        if (key !== 'compute_type') {
          let matchKeyIndex = parameter.findIndex((item) => {
            return item.label === key
          })
          if (matchKeyIndex !== -1) {
            parameter[matchKeyIndex].value = this.DepModelInfo.paramaters[key][this.pageCfg.computerResouce]
          } else {
            parameter.push({ label: key, value: this.DepModelInfo.paramaters[key][this.pageCfg.computerResouce] })
          }
        }
      })
      parameter.push({ label: 'template', value: this.DepModelInfo.template })
      let params = { parameter: parameter };
      subObj['params'] = JSON.stringify(params);
      subObj['has_internet'] = 2;
      subObj['spec_id'] = Number(this.state.spec);
      subObj['work_server_number'] = 1;
      this.maskLoadingContent = this.$t('cloudbrainObj.taskPrepareTips');
      this.maskLoading = true;
      this.errorMsg = '';
      this.errorMsgBoxShow = false;
      createAiTask(subObj).then(res => {
        const data = res.data;
        if (data.code == 0) {
          this.$router.replace('/nlp/sft');
        } else {
          this.maskLoading = false;
          this.errorMsg = data.msg;
          if (data.code == 2023) {
            this.errorMsg = data.msg + this.$t('cloudbrainObj.sameTaskTips2');
          }
          this.errorMsgBoxShow = true;
          document.querySelector('.main-body').scrollTo({ top: 0, behavior: 'smooth' });
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
      this.$router.replace('/nlp/sft');
    },
    prepare() {
      getAiTaskPrepareInfo({
        jobType: 'FINETUNE',
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
          if (data.config && data.config.dataset_max_num) {
            this.datasetCount = data.config.dataset_max_num;
            console.log("this.datasetCount,this.datasetCount", this.datasetCount)
          }
          if (this.pageCfg.computerResouce === 'NPU') {
            if (data.images !== null) {
              this.npuImageInfo = data.images.filter((item) => {
                return item.image_name === this.DepModelInfo.compute_sources["NPU"].image_url
              })
            }

          }
          this.dtype = this.DepModelInfo.paramaters.compute_type[this.pageCfg.computerResouce]
        }
      }).catch(err => {
        this.loading = false;
        this.$message.error(err?.response?.data?.message || err);
      });
    },
    changeComputeResouce(data) {
      if (this.pageCfg.computerResouce != data.computerResouce) {
        this.pageCfg.computerResouce = data.computerResouce;
        this.prepare();
      }
    },
    fetchLimitTaskList() {
      const params = {
        job_type: 'FINETUNE',
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
            const evalText = this.$t('modelSquare.sftFinetune');
            task.jobTypeShow = [evalText.slice(0, 3), evalText.slice(3)]
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
  },
  mounted() {
    let modelName = this.$route.query.model
    if (modelName) {
      modelName = decodeURIComponent(modelName)
    }
    getPromoteData('model/modelfinetune.json').then(res => {
      // this.loading = false;
      try {
        const data = JSON.parse(res.data);
        this.modelList = data.llm.model
        // this.modelList = data;

        if (modelName) {
          this.DepModelInfo = this.modelList.filter(item => item.name == modelName)[0]
          if (this.DepModelInfo.name) {
            this.modelName = modelName;
          }
        } else {
          if (this.modelList.length) {
            this.modelName = this.modelList[0].name;
            this.DepModelInfo = this.modelList[0]
          }
        }
        this.repoName = data.llm.repo_name
        this.repoOwnerName = data.llm.repo_owner_name
        this.bootfiles = this.DepModelInfo.bootfile ? this.DepModelInfo.bootfile : data.llm.bootfile
        this.tmpBootfiles = data.llm.bootfile
        this.datasetList = this.DepModelInfo.dataset ? this.DepModelInfo.dataset : data.llm.dataset
        this.tmpDatasetList = data.llm.dataset
        console.log(this.datasetList)
        console.log(this.tmpDatasetList)
        this.changeModel('', true);
        this.prepare();
      } catch (err) {
        this.loading = false;
        this.$message.error(err);
      }
    })
  },
  beforeMount() {
    const isLogin = !!document.querySelector('meta[name="_uid"]');
    if (isLogin) {
      this.loginName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext')
    }
  },
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';
@import '../../components/createcommon.less';
</style>
