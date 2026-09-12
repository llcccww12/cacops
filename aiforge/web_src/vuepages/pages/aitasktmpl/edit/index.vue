<template>
  <div v-if="emptyPage" style="padding-top:50px">
    <NotFound></NotFound>
  </div>
  <div v-else class="bg">
    <div class="ui container main-content">
      <div class="left">
        <div class="title-c">
          <div class="title">{{ title }}</div>
        </div>
        <div class="form-body" v-loading="dataPreparing">
          <!-- <div class="main-title">{{ $t('taskTmplObj.tmplInfo') }}：</div> -->
          <div class="main-title">{{ $t('cloudbrainObj.basicInfo') }}：</div>
          <TmplName ref="nameRef" v-model="state.name"></TmplName>
          <div class="form-row">
            <div class="left-area">
              <div class="title">{{ $t('taskTmplObj.tmplAccessRight') }}</div>
              <div class="content" style="display:flex;align-items:center;">
                <el-radio-group v-model="state.isPrivate">
                  <el-radio :label="false">{{ $t('modelManage.modelAccessPublic') }}</el-radio>
                  <el-radio :label="true">{{ $t('modelManage.modelAccessPrivate') }}</el-radio>
                </el-radio-group>
              </div>
            </div>
            <div class="right-area"></div>
          </div>
          <div class="form-row">
            <div class="left-area">
              <div class="title">{{ $t('taskTmplObj.tmplTags') }}</div>
              <div class="content">
                <Tags v-model="state.tags"></Tags>
              </div>
            </div>
            <div class="right-area"></div>
          </div>
          <div class="form-row">
            <div class="left-area">
              <div class="title">{{ $t('taskTmplObj.tmplDescr') }}</div>
              <div class="content">
                <el-input class="field-input task-descr" type="textarea" v-model="state.descr" :rows="2"
                  :placeholder="$t('cloudbrainObj.taskDescrPlaceholder')" :maxlength="255"></el-input>
              </div>
            </div>
            <div class="right-area"></div>
          </div>
          <TmplTaskType v-model="state.taskType" @change="changeTaskType"></TmplTaskType>
          <TmplTaskCluster v-show="false" :configs="taskTypeObj" @change="changeClusterAndComputeResouce">
          </TmplTaskCluster>
          <TmplTaskComputeResource :configs="taskTypeObj" @change="changeClusterAndComputeResouce">
          </TmplTaskComputeResource>
          <div class="required-hide">
            <NetworkType class="required-hide" ref="networkTypeRef" v-if="formCfg.networkType"
              v-model="state.networkType">
            </NetworkType>
            <Visualization ref="visualizationRef" v-if="formCfg.visualization" v-model="state.visualizeRequired">
            </Visualization>
            <SpecSelect class="required-hide" ref="specRef" v-if="formCfg.spec" v-model="state.spec"
              :required="formCfg.spec.required" :configs="specConfigs" :workServerNum="state.workServerNum"
              :networkType="state.networkType" :visualize="state.visualizeRequired">
            </SpecSelect>
            <div class="main-title">{{ $t('cloudbrainObj.paramsSetting') }}：</div>
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
            <RepoSelect ref="repoRef" v-if="formCfg.repo" v-model="repoList" :repoSize="repoSize"
              :required="formCfg.repo.required" :multiple="formCfg.repo.multiple">
            </RepoSelect>
            <BranchName ref="branchNameRef" v-if="formCfg.branchName" v-model="state.branchName"
              :required="formCfg.branchName.required" :branches="branchList">
            </BranchName>
            <BootFile ref="bootFileRef" v-if="formCfg.bootFile" v-model="state.bootFile"
              :required="formCfg.bootFile.required" :sampleUrl="formCfg.bootFile.sampleUrl"></BootFile>
            <RunParameters ref="runParametersRef" v-if="formCfg.runParameters" v-model="state.runParameters"
              :required="formCfg.runParameters.required"></RunParameters>
          </div>
          <div class="form-row">
            <div class="left-area">
              <div class="title"></div>
              <div class="content">
                <el-button type="primary" size="default" class="submit-btn" @click="submit">{{
                  $t('confirm')
                }}</el-button>
                <el-button class="cancel-btn" size="default" @click="cancel">{{ $t('cancel') }}</el-button>
              </div>
            </div>
            <div class="right-area"></div>
          </div>
        </div>
      </div>
    </div>
    <LoadingMask :loading="maskLoading" :content="maskLoadingContent"></LoadingMask>
  </div>
</template>

<script>
import NotFound from '~/components/NotFound.vue';
import Tags from '../components/Tags.vue'

import TmplName from './components/TmplName.vue';
import TmplTaskType from './components/TmplTaskType.vue';
import TmplTaskCluster from './components/TmplTaskCluster.vue';
import TmplTaskComputeResource from './components/TmplTaskComputeResource.vue';
import NetworkType from '~/components/cloudbrain/NetworkType.vue';
import Visualization from '~/components/cloudbrain/Visualization.vue';
import TmplDescr from './components/TmplDescr.vue';
import BranchName from '~/components/cloudbrain/BranchName.vue';
import ModelSelect from '~/components/cloudbrain/ModelSelectV2.vue';
import RepoSelect from '~/components/cloudbrain/RepoSelectV2.vue';
import ImageSelectV1 from '~/components/cloudbrain/ImageSelectV1.vue';
import ImageSelectV2 from '~/components/cloudbrain/ImageSelectV2.vue';
import BootFile from '~/components/cloudbrain/BootFile.vue';
import DatasetSelect from '~/components/cloudbrain/DatasetSelectV2.vue';
import RunParameters from '~/components/cloudbrain/RunParameters.vue';
import SpecSelect from '~/components/cloudbrain/SpecSelect.vue';
import LoadingMask from '~/components/cloudbrain/LoadingMask.vue';
import { configCreateManager } from '~/pages/cloudbrain/configs';
import { getAiTaskPrepareInfo, getAiTask } from '~/apis/modules/cloudbrain';
import { getRepoBranch } from "~/apis/modules/repos";
import { createAiTaskTmpl, getAiTaskTmpl, editAiTaskTmpl } from '~/apis/modules/aitasktmpl';
import { TaskTmplTools } from '../tools';
import { uuidv4 } from '~/utils';


const taskTmplTools = new TaskTmplTools();

const pageType = window.location.href.includes('/edit/') ? 'edit' : 'create';

export default {
  data() {
    return {
      emptyPage: false,
      pageType: pageType,
      title: pageType == 'edit' ? this.$t('taskTmplObj.editTaskTmpl') : this.$t('taskTmplObj.createNewTaskTmpl'),
      userName: document.querySelector('meta[name="_uid"]').getAttribute('content-ext'),
      formCfg: {},
      taskTypeObj: {},

      state: {
        name: '',
        descr: '',
        isPrivate: false,
        tags: [],
        taskType: '',
        cluster: '',
        computeResource: '',
        networkType: 'has_internet',
        visualizeRequired: false,
        spec: '',
        image: {
          image_url: '',
          image_id: '',
          image_name: '',
        },
        model: [],
        dataset: [],
        repoOwnerName: '',
        repoName: '',
        branchName: '',
        bootFile: '',
        runParameters: [],
        appName: '',
      },
      repoList: [],
      branchList: [],
      imageList: [],
      specConfigs: {
        specs: {
          'all': [],
          'no_internet': [],
          'has_internet': [],
        },
        blance: 0,
        showPoint: false,
        hideHelpLink: true,
      },

      dataPreparing: false,
      datasetSize: 0,
      datasetCount: 0,
      modelSize: 0,
      modelCount: 0,
      repoSize: 0,

      maskLoadingContent: '',
      maskLoading: false,
      isInited: false,
    };
  },
  components: {
    NotFound, Tags, TmplName, TmplTaskType, TmplTaskCluster, TmplTaskComputeResource, NetworkType, Visualization, TmplDescr, BranchName,
    ModelSelect, RepoSelect, ImageSelectV1, ImageSelectV2, BootFile, DatasetSelect, RunParameters, SpecSelect, LoadingMask
  },
  watch: {
    state: {
      deep: true,
      handler(nVal, oVal) {
        const selectSpec = this.specConfigs.specs.all.find(item => item.id == nVal.spec);
        nVal.acc_cards_num = selectSpec?.acc_cards_num || 0;
        nVal.acc_card_type = selectSpec?.acc_card_type || '';
        nVal.cpu_cores = selectSpec?.cpu_cores || 0;
        nVal.mem_gi_b = selectSpec?.mem_gi_b || 0;
        nVal.gpu_mem_gi_b = selectSpec?.gpu_mem_gi_b || 0;
        nVal.share_mem_gi_b = selectSpec?.share_mem_gi_b || 0;
        // console.log('[TmplEdit]  state change', nVal);
        this.$emit('input', nVal);
        this.$emit('change', nVal);
      }
    },
    'repoList': {
      handler(newVal, oldVal) {
        if (JSON.stringify(newVal) !== JSON.stringify(oldVal)) {
          if (newVal.length > 0) {
            let branch = newVal[0]?.branch_name ?? this.state.branchName;
            this.getRepoBranchList(newVal[0].owner_name, newVal[0].name, branch);
          } else {
            this.state.branchName = ''
            this.state.repoOwnerName = ''
            this.state.repoName = ''
            this.branchList = []
          }
        }
      },
      deep: true
    }
  },
  methods: {
    resetForm() {
      this.state.networkType = 'has_internet';
      this.state.visualizeRequired = false;
      this.state.spec = '';
      this.state.image = {
        image_url: '',
        image_id: '',
        image_name: '',
      };
      this.imageList = [];
      this.state.bootFile = '';
      this.state.runParameters = [];
      // this.branchList = [];
      // this.repoList = [];
      // this.state.repoOwnerName = '';
      // this.state.repoName = '';
      // this.specConfigs = {
      //   specs: {
      //     'all': [],
      //     'no_internet': [],
      //     'has_internet': [],
      //   },
      //   blance: 0,
      //   showPoint: false,
      //   hideHelpLink: true,
      // };
    },
    async getRepoBranchList(repoOwnerName, repoName, branch) {
      try {
        this.branchLoading = true;
        this.state.repoOwnerName = repoOwnerName
        this.state.repoName = repoName
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
    changeTaskType(type, typeObj) {
      // console.log('[tmplEdit] changeTaskType', type, typeObj);
      if (typeObj) {
        this.taskTypeObj = { taskType: type, ...JSON.parse(JSON.stringify(typeObj)) }
        if (!this.isInited) {
          this.state.cluster = this.state.cluster || this.taskTypeObj.cluster;
          this.state.computeResource = this.state.computeResource || this.taskTypeObj.computerResouce;
        } else {
          this.state.cluster = this.taskTypeObj.cluster;
          this.state.computeResource = this.taskTypeObj.computerResouce;
        }
        this.changeClusterAndComputeResouce({
          cluster: this.state.cluster,
          computerResouce: this.state.computeResource,
        });
      }
    },
    changeClusterAndComputeResouce(obj) {
      this.taskTypeObj = {
        ...this.taskTypeObj,
        ...obj
      };
      this.state.cluster = obj.cluster || this.state.cluster;
      this.state.computeResource = obj.computerResouce || this.state.computeResource;
      const params = {
        taskType: this.state.taskType,
        cluster: this.state.cluster,
        computeResource: this.state.computeResource,
      };
      this.configTimer && clearTimeout(this.configTimer);
      this.configTimer = setTimeout(() => {
        // console.log('[tmplEdit] changeClusterAndComputeResouce');
        const configs = configCreateManager.getResourceConfig(params.taskType, params.computeResource);
        if (!configs) {
          console.error('[tmplEdit] get configs error, params=', params);
          return;
        }
        // console.log('[tmplEdit] configs=', configs);
        this.formCfg = configs.form || {};
        this.state.appName = configs.appName ? configs.appName : '';
        this.resetForm();
        this.getPrepareInfo();
      }, 20);
    },
    getPrepareInfo() {
      this.prepareTimer && clearTimeout(this.prepareTimer);
      this.prepareTimer = setTimeout(() => {
        this.maskLoadingContent = this.$t('cloudbrainObj.dataPreparing');
        this.maskLoading = true;
        getAiTaskPrepareInfo({
          jobType: this.taskTypeObj.taskType,
          clusterType: this.taskTypeObj.cluster,
          computeSource: this.taskTypeObj.computerResouce,
          appName: this.state.appName
        }).then(res => {
          res = res.data;
          this.maskLoadingContent = '';
          this.maskLoading = false;
          if (res.code == 0) {
            const data = res.data;
            this.subscriberUser = data.is_subscriber;
            if (!this.subscriberUser) delete this.formCfg['runTimeLimit'];
            this.imageList = data.images || [];
            this.specConfigs.specs = data.specs || {
              'all': [],
              'no_internet': [],
              'has_internet': [],
            };
            this.state.networkType = this.formCfg.networkType ? 'no_internet' : 'all';
            this.state.visualizeRequired = false;

            this.repoSize = data.code_size_limit;
            if (this.state.networkType == 'no_internet'
              && !this.specConfigs.specs['no_internet'].length
              && this.specConfigs.specs['has_internet'].length) {
              this.state.networkType = 'has_internet';
            }
            if (data.config && data.config.dataset_max_num) {
              this.datasetSize = data.config.dataset_max_size;
              this.datasetCount = data.config.dataset_max_num;
            }
            if (data.config && data.config.model_max_num) {
              this.modelSize = data.config.model_max_size;
              this.modelCount = data.config.model_max_num;
            }
            if (this.branchList.length == 0) {
              this.branchList = data.branches || [];
              this.state.branchName = data.default_branch || '';
            }
            if (!this.isInited && Object.keys(this.oriData).length) {
              this.fillingInfo();
            } else {
              this.state.spec = this.specConfigs.specs[this.state.networkType][0] ? this.specConfigs.specs[this.state.networkType][0].id.toString() : '';
              this.isInited = true;
            }
          } else { }
        }).catch(err => {
          this.maskLoadingContent = '';
          this.maskLoading = false;
          console.log(err);
        });
      }, 20);
    },
    changeImage() {
      this.state.image.image_url = '';
      this.state.image.image_id = '';
      this.state.image.image_name = '';
      if (this.imageChangeShouldRefreshInit) {
        this.fillTmplImage();
      }
    },
    changeImages(images) {
      this.imageList = images || [];
      const image = this.imageList[0];
      if (image) {
        this.state.image.image_id = image.image_id;
        this.state.image.image_name = image.image_name;
      } else {
        this.state.image.image_id = '';
        this.state.image.image_name = '';
      }
      if (this.imageChangeShouldRefreshInit) {
        this.fillTmplImage();
      }
    },
    getRandomTmplName() {
      return this.userName.slice(0, 5) + '_tmpl_' + uuidv4().slice(-5);
    },
    fillTmplImage() {
      const oriData = this.oriData;
      this.fillImageTimer && clearTimeout(this.fillImageTimer);
      this.fillImageTimer = setTimeout(() => {
        this.state.image.image_id = '';
        this.state.image.image_name = '';
        this.state.image.image_url = '';
        if (this.formCfg['imagev1'] && (oriData.image?.image_id || oriData.image?.image_url)) {
          this.state.image.image_url = oriData.image?.image_url;
          this.state.image.image_id = oriData.image?.image_id || '';
          this.state.image.image_name = oriData.image?.image_name || '';
        }
        if (this.formCfg['imagev2'] && oriData.image) {
          if (this.imageList.filter(item => item.image_id == oriData.image.image_id).length) {
            this.state.image.image_id = oriData.image.image_id;
            this.state.image.image_name = oriData.image.image_name;
            this.state.image.image_url = ''
          }
        }
      }, 80);
    },
    fillingInfo() {
      // console.log('fillingInfo', this.oriData);
      const oriData = this.oriData;
      this.state.networkType = this.oriData.networkType || 'has_internet';
      this.state.visualizeRequired = !!this.oriData.visualizeRequired;
      this.$nextTick(() => {
        const findSpec = this.specConfigs.specs[this.state.networkType].find(item => {
          return item.acc_cards_num == oriData.acc_cards_num
            && item.acc_card_type == oriData.acc_card_type
            && item.cpu_cores == oriData.cpu_cores
            && item.mem_gi_b == oriData.mem_gi_b
            && item.gpu_mem_gi_b == oriData.gpu_mem_gi_b
            && item.share_mem_gi_b == oriData.share_mem_gi_b
        });
        this.imageChangeShouldRefreshInit = true;
        if (findSpec && findSpec.id.toString() != this.state.spec) {
          this.state.spec = findSpec.id.toString();
        }
        if (oriData.repoOwnerName) {
          this.state.repoOwnerName = oriData.repoOwnerName;
        }
        if (oriData.repoName) {
          this.state.repoName = oriData.repoName;
        }
        if (oriData.repoOwnerName && oriData.repoName) {
          this.repoList = [{
            id: oriData.repoID,
            owner_name: oriData.repoOwnerName,
            alias: oriData.repoName,
            name: oriData.repoName,
            branch_name: oriData.branchName,
          }]
        }
        setTimeout(() => {
          if (this.formCfg['branchName'] && oriData.branchName) {
            if (this.branchList.indexOf(oriData.branchName) > -1) {
              this.state.branchName = oriData.branchName;
            }
          }
        }, 200);
        if (oriData.bootFile) {
          this.state.bootFile = oriData.bootFile;
        }
        if (oriData.bootFile) {
          this.state.bootFile = oriData.bootFile;
        }
        if (oriData.runParameters) {
          this.state.runParameters = oriData.runParameters;
        }
        if (oriData.dataset) {
          this.state.dataset = oriData.dataset;
        }
        if (oriData.model) {
          this.state.model = oriData.model;
        }
        this.fillTmplImage();
        setTimeout(() => {
          this.imageChangeShouldRefreshInit = false;
        }, 400);
        this.isInited = true;
      });
    },
    submit() {
      // console.log('submit', this.state);
      if (!this.$refs.nameRef.check() || !this.state.taskType || !this.state.cluster || !this.state.computeResource) {
        this.$message({
          type: 'info',
          message: this.$t('taskTmplObj.createTaskTmplErrTips'),
        });
        return;
      }
      this.state.descr = this.state.descr.trim();
      const tmplObj = taskTmplTools.transformDataFormToTmpl(this.state);
      const setApi = this.pageType == 'edit' ? editAiTaskTmpl : createAiTaskTmpl;
      if (this.pageType == 'edit') {
        tmplObj.id = this.oriData.id;
        tmplObj.OwnerId = this.oriData.ownerId;
      }
      // console.log('submit tmplObj', tmplObj);
      // return;
      this.maskLoadingContent = this.$t('taskTmplObj.taskTmplSaveTips');
      this.maskLoading = true;
      setApi(tmplObj).then(res => {
        res = res.data;
        if (res.code == 0) {
          window.location.href = `/ai_task_tmpl/detail/${tmplObj.id || res.data.id}`;
        } else {
          this.$message({
            type: 'error',
            message: res.msg || this.$t('taskTmplObj.taskTmplSaveFailedTips'),
          });
        }
      }).catch(err => {
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('taskTmplObj.taskTmplSaveFailedTips'),
        });
      }).finally(() => {
        this.maskLoading = false;
      })
    },
    cancel() {
      if (this.cancelUrl) {
        window.location.href = this.cancelUrl;
      } else {
        window.history.back();
      }
    },
  },
  beforeMount() {
    if (this.pageType == 'edit') {
      const pathname = window.location.pathname;
      const id = pathname.split('/')[3];
      if (!id) {
        this.emptyPage = true;
        return;
      }
      this.dataPreparing = true;
      getAiTaskTmpl({ id }).then(res => {
        res = res.data;
        if (res.code == 0) {
          const formData = taskTmplTools.transformDataTmplToForm(res.data);
          this.oriData = {
            ...formData,
          }
          // console.log('oriData', this.oriData);
          this.state = {
            ...this.state,
            ...formData
          };
        } else {
          this.emptyPage = true;
        }
      }).catch(err => {
        console.log(err);
        this.emptyPage = true;
      }).finally(() => {
        this.dataPreparing = false;
      });
    } else {
      const params = new URLSearchParams(window.location.href.split('?')[1]);
      const taskId = params.get('task');
      if (taskId && Number.isInteger(Number(taskId))) {
        getAiTask({ id: taskId }).then((res) => {
          res = res.data;
          if (res.code == 0) {
            const task = res.data.task;
            const formData = taskTmplTools.transformDataTaskToForm(task);
            this.oriData = {
              ...formData,
            }
            // console.log('oriData', this.oriData);
            this.state = {
              ...this.state,
              ...formData,
              name: this.getRandomTmplName(),
            };
          }
        }).catch(err => {
          console.log(err);
          this.oriData = {};
          this.state.taskType = 'DEBUG';
        });
      } else {
        this.oriData = {};
        this.state.taskType = 'DEBUG';
        this.state.name = this.getRandomTmplName();
      }
    }
    this.maskLoadingContent = this.$t('cloudbrainObj.dataPreparing');
    this.maskLoading = true;
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style>
.full.height {
  background-color: rgba(249, 249, 249, 1);
}
</style>
<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';

.bg {}

.main-content {
  display: flex;
  border-radius: 10px;
  background-color: rgba(255, 255, 255, 1);
  margin-top: 44px;
  box-shadow: 0px 2px 6px 0px rgba(178, 192, 255, 0.5);
  border: 1px solid rgba(224, 227, 234, 1);
  min-height: 300px;

  .left {
    flex: 1;
    width: 0;
    padding: 20px;
    padding-right: 220px;
    padding-top: 12px;

    .title-c {
      display: flex;
      width: 100%;
      margin-bottom: 24px;

      .title {
        height: 32px;
        line-height: 32px;
        display: inline-block;
        color: rgba(16, 16, 16, 0.8);
        font-size: 16px;
        border-bottom: 2px solid rgba(16, 16, 16, 0.8);
      }
    }

    .form-body {
      padding-left: 220px;

      /deep/ .form-row .left-area .title {
        width: 133px;
      }
    }

    .required-hide {
      /deep/ .title {
        .required {
          &::after {
            color: transparent;
          }
        }
      }
    }

    .task-descr {
      /deep/ .el-textarea__inner {
        line-height: 1.2857 !important;
        padding: 0.78571429em 1em;
        min-height: 76px !important;
      }
    }
  }

  .right {
    width: 375px;
    padding: 20px;
    background-color: rgba(247, 247, 247, 1);

    .main-title {
      padding-left: 0;
    }

    .row {
      display: flex;
      align-items: center;
      margin-bottom: 12px;

      .label {
        color: rgba(16, 16, 16, 0.8);
        margin-right: 10px;
      }
    }

    .multi-row {
      margin-bottom: 12px;

      .label {
        color: rgba(16, 16, 16, 0.8);
        margin-bottom: 6px;
      }

      .el-textarea {
        /deep/ textarea {
          border: 1px solid rgba(157, 197, 226, 0.4);
          border-radius: 5px;

          &:focus {
            border-color: #409EFF;
          }
        }
      }
    }
  }

  .main-title {
    font-size: 16px;
    color: rgb(16, 16, 16);
    padding-left: 1rem;
    font-weight: 700;
    line-height: 1.28571429em;
    margin: 16px 0;
    font-family: sans-serif;
  }
}

@media only screen and (max-width: 1600px) {
  .main-content {
    .left {
      padding-right: 100px;

      .form-body {
        padding-left: 100px;
      }
    }
  }
}

@media only screen and (max-width: 1200px) {
  .main-content {
    .left {
      padding-right: 30px;

      .form-body {
        padding-left: 30px;
      }
    }
  }
}

@media screen and (max-width: 767px) {
  .main-content {
    padding-right: 10px;

    .form-body {
      padding-left: 10px;
    }

    .main-title {
      padding-left: 0;
    }
  }
}
</style>
