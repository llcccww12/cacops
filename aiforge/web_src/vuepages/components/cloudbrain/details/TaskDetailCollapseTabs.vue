<template>
  <div class="task-detail-container">
    <div class="body-content">
      <div class="content" v-for="(item, index) in mainData">
        <div class="operation-area">
          <!-- task  -->
          <div class="operate-item" v-if="item.task.isSubscriber && operationList.indexOf('consoleHome') >= 0">
            <div class="operate-btn-c" :class="item.can_modify && item.task.canDebug ? '' : 'btn-disabled'">
              <Icons type="1"></Icons>
              <div class="operate-btn" @click="opDebug(item)">{{ $t('cloudbrainObj.consoleHome') }}</div>
            </div>
          </div>
          <div class="operate-item" v-if="operationList.indexOf('onlineInfer') >= 0 && item.task.canDebug">
            <div class="operate-btn-c" :class="item.can_modify && item.task.canDebug ? '' : 'btn-disabled'">
              <Icons type="1"></Icons>
              <div class="operate-btn" @click="opVisual(item)">{{ $t('onlineinfer') }}</div>
            </div>
          </div>
          <div class="operate-item" v-if="operationList.indexOf('onlinexperience') >= 0">
            <div class="operate-btn-c" :class="item.can_experience && item.task.canDebug ? '' : 'btn-disabled'">
              <Icons type="1"></Icons>
              <div class="operate-btn" @click="opOnline(item)">{{ $t('modelManage.onlineInference') }}</div>
            </div>
          </div>
          <div class="operate-item" v-if="operationList.indexOf('onlineLoraTrain') >= 0">
            <div class="operate-btn-c" :class="item.can_experience && item.task.canDebug ? '' : 'btn-disabled'">
              <Icons type="1"></Icons>
              <div class="operate-btn" @click="opLoraTrain(item)">{{ $t('modelManage.onlineLoraTrain') }}</div>
            </div>
          </div>
          <div class="operate-item" v-if="operationList.indexOf('onlineWorkflow') >= 0">
            <div class="operate-btn-c" :class="item.can_experience && item.task.canDebug ? '' : 'btn-disabled'">
              <Icons type="1"></Icons>
              <div class="operate-btn" @click="opComfyUI(item)">{{ $t('modelManage.onlineWorkflow') }}</div>
            </div>
          </div>
          <div class="operate-item" v-if="operationList.indexOf('debug') >= 0 && item.task.canDebug">
            <div class="operate-btn-c" :class="item.can_modify && item.task.canDebug ? '' : 'btn-disabled'">
              <Icons type="1"></Icons>
              <div class="operate-btn" @click="opDebug(item)">{{ $t('cloudbrainObj.startDebug') }}</div>
            </div>
          </div>
          <div class="operate-item"
            v-if="operationList.indexOf('redebug') >= 0 && !item.task.canDebug && !item.task.is_file_notebook">
            <div class="operate-btn-c" :class="item.can_modify && item.task.canReDebug ? '' : 'btn-disabled'">
              <Icons type="1"></Icons>
              <div class="operate-btn" @click="opReDebug(item)">{{ $t('cloudbrainObj.reDebug') }}</div>
            </div>
          </div>
          <div class="operate-item"
            v-if="operationList.indexOf('reinfer') >= 0 && !item.task.canDebug && !item.task.is_file_notebook">
            <div class="operate-btn-c" :class="item.can_modify && item.task.canReDebug ? '' : 'btn-disabled'">
              <Icons type="1"></Icons>
              <div class="operate-btn" @click="opReDebug(item)">{{ $t('cloudbrainObj.reInfer') }}</div>
            </div>
          </div>
          <div class="operate-item" v-if="operationList.indexOf('tensorBoard') >= 0 && item.task.visualize_required">
            <div class="operate-btn-c" :class="item.can_modify && item.task.canVisualize ? '' : 'btn-disabled'">
              <Icons type="2"></Icons>
              <div class="operate-btn" @click="opTensorBoard(item)">{{ $t('cloudbrainObj.tensorBoardVisualization')
                }}</div>
            </div>
          </div>
          <div class="operate-item" v-if="item.can_modify && item.task.canAim">
            <div class="operate-btn-c">
              <Icons type="1"></Icons>
              <div class="operate-btn" @click="opAim(item)">{{ $t('cloudbrainObj.aimVisualization') }}</div>
            </div>
          </div>
          <div class="operate-item" v-if="operationList.indexOf('stop') >= 0">
            <div class="operate-btn-c" :class="item.can_delete && item.task.canStop ? '' : 'btn-disabled'">
              <Icons type="1"></Icons>
              <div class="operate-btn" @click="opStop(item)">{{ $t('cloudbrainObj.stopTask') }}</div>
            </div>
          </div>
          <!--  comfyui 专属  -->
          <template v-if="['ComfyuiExperience'].includes(item.task.job_type)">
            <div class="operate-item" v-if="item.can_experience && item.task.canSaveImage">
              <a class="operate-btn-c" target="_blank" v-if="item.can_experience && item.task.canSaveImage"
                :href="item.task.saveImageUrl">
                <Icons type="1"></Icons>
                <div class="operate-btn"> {{ $t('cloudbrainObj.commitImage') }}</div>
              </a>
            </div>
          </template>
          <template v-else>
            <div class="operate-item" v-if="item.can_modify && item.task.canSaveImage">
              <a class="operate-btn-c" target="_blank" v-if="item.can_modify && item.task.canSaveImage"
                :href="item.task.saveImageUrl">
                <Icons type="1"></Icons>
                <div class="operate-btn"> {{ $t('cloudbrainObj.commitImage') }}</div>
              </a>
            </div>
          </template>

          <!-- others  -->


          <div class="operate-item" v-if="item.can_create_template && item.task.canSaveTmpl">
            <div class="operate-btn-c" :class="item.can_create_template && item.task.canSaveTmpl ? '' : 'btn-disabled'">
              <Icons type="3"></Icons>
              <div class="operate-btn" @click="saveTmpl(item)">{{ $t('cloudbrainObj.saveTaskTmpl') }}</div>
            </div>
          </div>
          <div class="operate-item export-output" v-if="operationList.indexOf('saveModel') >= 0">
            <div class="operate-btn-c" :class="!item.task.canExportOutput ? 'btn-disabled' : ''">
              <Icons type="4"></Icons>
              <ExportModel :data="item.task" :configs="pageCfg" :disabled="!item.task.canExportOutput">
              </ExportModel>
            </div>
          </div>
          <div class="operate-item export-output" v-if="operationList.indexOf('exportDataset') >= 0">
            <div class="operate-btn-c" :class="!item.task.canExportOutput ? 'btn-disabled' : ''">
              <Icons type="5"></Icons>
              <ExportDataset :data="item.task" :configs="pageCfg" :disabled="!item.task.canExportOutput">
              </ExportDataset>
            </div>
          </div>
          <div class="operate-item" v-if="operationList.indexOf('deployModel') >= 0">
            <div class="operate-btn-c" :class="!item.can_fintune_experience ? 'btn-disabled' : ''">
              <Icons type="5"></Icons>
              <div class="operate-btn" @click="deployModel(item)">{{ $t('cloudbrainObj.deploymentExperience') }}</div>
            </div>
          </div>
          <div class="mobile-export-output"
            v-if="operationList.indexOf('exportDataset') >= 0 || operationList.indexOf('saveModel') >= 0">
            {{ $t('cloudbrainObj.exportOutputTis') }}
          </div>
        </div>
        <el-tabs v-model="item.activeName" @tab-click="tabChange(item)">
          <el-tab-pane v-if="tabNameList.indexOf('configInfo') >= 0" :label="$t('cloudbrainObj.configurationInfo')"
            :name="`configInfo-` + item.task.id">
            <ConfigInfo :ref="`configInfo-` + item.task.id + '-Ref'" :data="item" :configs="tabConfigs['configInfo']">
            </ConfigInfo>
          </el-tab-pane>
          <el-tab-pane v-if="tabNameList.indexOf('operationProfile') >= 0" :label="$t('cloudbrainObj.taskRuntimeInfo')"
            :name="`operationProfile-` + item.task.id">
            <OperationProfile :ref="`operationProfile-` + item.task.id + '-Ref'" :data="item"
              :configs="tabConfigs['operationProfile']"></OperationProfile>
          </el-tab-pane>
          <el-tab-pane v-if="tabNameList.indexOf('logs') >= 0" :label="$t('cloudbrainObj.log')"
            :name="`logs-` + item.task.id">
            <Logs :ref="`logs-` + item.task.id + '-Ref'" :data="item" :configs="tabConfigs['logs']"></Logs>
          </el-tab-pane>
          <el-tab-pane v-if="tabNameList.indexOf('resourceUseage') >= 0" :label="$t('cloudbrainObj.resourceOccupancy')"
            :name="'resourceUseage-' + item.task.id">
            <ResourceUseage :ref="`resourceUseage-` + item.task.id + '-Ref'" :data="item"
              :configs="tabConfigs['resourceUseage']"></ResourceUseage>
          </el-tab-pane>
          <el-tab-pane v-if="tabNameList.indexOf('resultDownload') >= 0" :label="$t('cloudbrainObj.modelDownload')"
            :name="`resultDownload-` + item.task.id">
            <ResultDownload :ref="`resultDownload-` + item.task.id + '-Ref'" :data="item"
              :configs="tabConfigs['resultDownload']"></ResultDownload>
          </el-tab-pane>
          <el-tab-pane v-if="tabNameList.indexOf('loss') >= 0" :label="$t('cloudbrainObj.lossPlot')"
            :name="`loss-` + item.task.id">
            <Loss :ref="`loss-` + item.task.id + '-Ref'" :data="item"></Loss>
          </el-tab-pane>
          <el-tab-pane v-if="tabNameList.indexOf('evalOverview') >= 0" :label="$t('cloudbrainObj.evalOverview')"
            :name="`evalOverview-` + item.task.id">
            <EvalOverview :ref="`evalOverview-` + item.task.id + '-Ref'" :data="item"></EvalOverview>
          </el-tab-pane>
          <el-tab-pane v-if="tabNameList.indexOf('evalDetail') >= 0" :label="$t('cloudbrainObj.evalDetail')"
            :name="`evalDetail-` + item.task.id">
            <EvalDetail :ref="`evalDetail-` + item.task.id + '-Ref'" :data="item"></EvalDetail>
          </el-tab-pane>
        </el-tabs>
      </div>
      <el-dialog class="task-already-dlg" :visible.sync="taskAlreadyDialogShow" :lock-scroll="false"
        :show-close="false">
        <div class="err-msg-box-already">
          <div class="msg-content">
            <i class="ri-information-line"></i>
            <div class="msg-content-tip">
              <div class="line-1" v-html="taskAlreadyDialogShowMsg"></div>
              <div class="line-2" v-html="$t('cloudbrainObj.sameTaskTips2')"></div>
            </div>
          </div>
        </div>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import NotFound from '~/components/NotFound.vue';
import LoadingMask from '~/components/cloudbrain/LoadingMask.vue';
import ConfigInfo from '~/components/cloudbrain/details/ConfigInfo.vue';
import OperationProfile from '~/components/cloudbrain/details/OperationProfile.vue';
import Logs from '~/components/cloudbrain/details/Logs.vue';
import Loss from '~/components/cloudbrain/details/Loss.vue';
import ResourceUseage from '~/components/cloudbrain/details/ResourceUseage.vue';
import ResultDownload from '~/components/cloudbrain/details/ResultDownload.vue';
import ExportModel from '~/components/cloudbrain/details/ExportModel.vue';
import ExportDataset from '~/components/cloudbrain/details/ExportDataset.vue';
import EvalOverview from '~/components/cloudbrain/details/EvalOverview.vue';
import EvalDetail from '~/components/cloudbrain/details/EvalDetail.vue';
import Icons from './Icons.vue';
import { CloudBrainTools } from '~/pages/cloudbrain/tools';
import { TaskTmplTools } from '~/pages/tasktmpl/tools';
import { CLUSTERS, BenchmarkTypeList, JOB_TYPE } from '~/const';
import { timeSinceUnix, getListValueWithKey } from '~/utils';
import { configDetailManager } from '~/pages/cloudbrain/configs';
import {
  getAiTask, getAiTaskDebugUrl, getAiTaskVisualizeUrl, getAiTaskRestart,
  stopAiTask, getTmplEditAdress, getAiEndpointUrl, getAimUrl
} from '~/apis/modules/cloudbrain';
import { commonFormPost, getComfyuiUrl } from '~/apis/modules/common';
import { formatDate } from 'element-ui/lib/utils/date-util';

const cloudBrainTools = new CloudBrainTools();
const taskTmplTools = new TaskTmplTools();

export default {
  name: 'TaskDetailCollapseTabs',
  props: {
    taskId: { type: String, default: '', required: true, },
    autoInit: { type: Boolean, default: true, }
  },
  data() {
    return {
      notFound: false,
      pageCfg: {},
      tabNameList: [],
      tabConfigs: {},
      operationList: [],
      state: {},
      collapseValue: '0',
      mainData: [],

      loading: false,
      errorMsgBoxShow: false,
      errorMsg: '',
      maskLoading: false,
      maskLoadingContent: '',

      operating: false,
      taskAlreadyDialogShow: false,
      taskAlreadyDialogShowMsg: '',
    };
  },
  components: { NotFound, LoadingMask, ConfigInfo, OperationProfile, Logs, ResourceUseage, ResultDownload, ExportModel, ExportDataset, Loss, EvalOverview, EvalDetail, Icons },
  methods: {
    tabChange(item) {
      console.log('tabChange', item);
      this.$nextTick(() => {
        this.refresh(item);
      });
    },
    refresh(item, useRefreshBtn) {
      const activeTab = item.activeName;
      const tabContent = this.$refs[activeTab + '-Ref'];
      console.log("activeTab.indexOf('configInfo-') > -1", activeTab.indexOf('configInfo-') > -1)
      if (activeTab.indexOf('configInfo-') > -1) {
        const task = item.task;
        const taskId = task.id;
        getAiTask({
          id: taskId,
        }).then(res => {
          res = res.data;
          if (res.code == 0) {
            Object.assign(task, res.data.task);
            delete res.data.task;
            delete res.data.early_version_list;
            Object.assign(item, res.data);
            Object.assign(task, res.data);
            task.createdFromNow = timeSinceUnix(task.created_unix, Date.now() / 1000);
            task.createTimeStr = formatDate(new Date(task.created_unix * 1000), 'yyyy-MM-dd HH:mm:ss');
            if (['FINETUNE', 'MODELEXPERIENCE', 'SDFINETUNE', 'ComfyuiExperience', 'EVAL'].includes(task.job_type)) {
              task.sdk_code = ''
            }
            cloudBrainTools.checkRunningLeftTime(task);
            cloudBrainTools.checkOperation(task);
          }
        }).catch(err => {
          console.log(err);
        });
      } else {
        tabContent && tabContent[0] && tabContent[0].refresh && tabContent[0].refresh(useRefreshBtn);
      }
    },
    calcFromNow(unix) {
      return timeSinceUnix(unix, Date.now() / 1000);
    },
    dateFormat(unix) {
      return dayjs(unix * 1000).format('YYYY-MM-DD HH:mm:ss');
    },
    // ops
    opVisual(row) {
      if (this.operating) return;
      this.operating = true;
      getAiEndpointUrl({
        id: row.task.id,
      }).then(res => {
        this.operating = false;
        res = res.data;
        if (res.code == 0) {
          if (res.data && res.data.url) {
            if (!window.open(res.data.url)) {
              window.location.href = res.data.url;
            }
          }
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      }).catch(err => {
        this.operating = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      });
    },
    opReDebug(row) {
      if (this.operating) return;
      this.operating = true;
      getAiTaskRestart({
        id: row.task.id,
      }).then(res => {
        this.operating = false;
        res = res.data;
        if (res.code == 0) {
          window.location.href = `/cloudbrains/detail/${res.data.id}`;
        } else if (res.code == 2004) {
          this.taskAlreadyDialogShow = true;
          this.taskAlreadyDialogShowMsg = res.msg;
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      }).catch(err => {
        console.log(err);
        this.operating = false;
        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      });
    },
    opStop(row) {
      if (this.operating) return;
      this.operating = true;
      stopAiTask({ id: row.task.id }).then(res => {
        this.operating = false;
        res = res.data;
        if (res.code == 0) {
          const data = res.data;
          Object.assign(row.task, data);
          row.task.createdFromNow = this.calcFromNow(row.task.created_unix);
          cloudBrainTools.checkOperation(row.task);
        } else {
          this.$message({
            type: 'error',
            message: res.msg
          });
        }
      }).catch(err => {
        this.operating = false;

        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      });
    },
    opOnline(row) {
      let url = ''
      let task = row.task
      if (task.label_name === "chat") {
        url = `/extension/modelexperience/chat?id=${btoa(task.id)}&modelName=${task.app_name}`
      }
      if (row.task.label_name === "sd") {
        url = `/extension/modelexperience/sd?id=${btoa(task.id)}&modelName=${task.app_name}`
      }
      if (row.task.label_name === "tts") {
        url = `/extension/modelexperience/tts?id=${btoa(task.id)}&modelName=${task.app_name}`
      }
      window.open(url, '_blank')
    },
    opLoraTrain(row) {
      const url = `/modelbase/cv/sft/lora?id=${btoa(row.task.id)}&model=${row.task.app_name}`
      window.open(url, '_blank')
    },
    opComfyUI(row) {
      getComfyuiUrl({ task_id: row.task.id }).then((res) => {
        window.open(res.data.url, '_blank')
      })
    },
    opAim(row) {
      if (this.operating) return;
      this.operating = true;
      getAimUrl({
        job_name: [row.task.display_job_name],
      }).then(res => {
        this.operating = false;
        res = res.data;
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
        this.operating = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      });
    },
    saveTmpl(row) {
      if (this.operating) return;
      window.location.href = `/ai_task_tmpl/create?task=${row.task.id}`;
    },
    scowFunc(url, ptoken) {
      const form = document.createElement('form');
      form.style.display = 'none';
      form.action = url;
      form.method = 'post';
      form.target = '_blank';
      const input = document.createElement('input');
      input.type = 'hidden';
      input.name = 'password';
      input.value = ptoken;
      form.appendChild(input);
      document.body.appendChild(form);
      form.submit();
      document.body.removeChild(form);
    },
    deployModel(row) {
      const item = row.task
      let url = `/modelbase/experience/create?id=${item.id}&job_name=${item.display_job_name}&compute_resource=${item.compute_source}&modelName=${item.app_name}`;
      window.location.href = url;
    },
    opDebug(row) {
      if (this.operating) return;
      this.operating = true;
      getAiTaskDebugUrl({
        id: row.task.id,
      }).then(res => {
        this.operating = false;
        res = res.data;
        if (res.code == 0) {
          if (res.data && res.data.url) {
            if (!!(res.data?.ptoken)) {
              this.scowFunc(res.data.url, res.data.ptoken)
              return
            }
            if (!window.open(res.data.url)) {
              window.location.href = res.data.url;
            }
          }
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      }).catch(err => {
        this.operating = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      });
    },
    opTensorBoard(row) {
      if (this.operating) return;
      this.operating = true;
      getAiTaskVisualizeUrl({
        id: row.task.id,
      }).then(res => {
        this.operating = false;
        res = res.data;
        if (res.code == 0) {
          if (res.data && res.data.url) {
            if (!window.open(res.data.url)) {
              window.location.href = res.data.url;
            }
          }
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      }).catch(err => {
        this.operating = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      });
    },
    emitUpdate() {
      console.log("emitUpdate")
      this.$nextTick(() => {
        this.$emit('update', {
          notFound: this.notFound,
          mainData: this.mainData,
          pageCfg: this.pageCfg,
        })
      });
    },
    init() {
      this.loading = true;
      getAiTask({
        id: this.taskId,
      }).then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 0 && res.data) {
          const earlyVersionList = (res.data.early_version_list || []).map(item => {
            return Object.assign({}, res.data, { task: item });
          });
          const extraData = Object.assign({}, res.data);
          delete extraData.early_version_list;
          delete extraData.task;
          const data = [res.data, ...earlyVersionList];
          let computeSource = '', taskType = '', cluster = '';
          data.forEach(item => {
            const task = item.task;
            Object.assign(task, extraData);
            task.computeSourceShow = task.compute_source == 'GPU' ? 'CPU/GPU' : task.compute_source;
            computeSource = task.compute_source;
            taskType = task.job_type;
            cluster = task.cluster;
            item.activeName = 'configInfo-' + task.id;
            task.clusterName = getListValueWithKey(CLUSTERS, task.cluster);
            task.jobTypeShow = BenchmarkTypeList.indexOf(taskType) >= 0 ? this.$t('benchmarkTask') : getListValueWithKey(JOB_TYPE, taskType);
            task.createdFromNow = timeSinceUnix(task.created_unix, Date.now() / 1000);
            task.createTimeStr = formatDate(new Date(task.created_unix * 1000), 'yyyy-MM-dd HH:mm:ss');
            cloudBrainTools.checkRunningLeftTime(task);
            cloudBrainTools.checkOperation(task);
          });
          console.log(taskType, cluster, computeSource)
          const configs = configDetailManager.getConfig(taskType, cluster, computeSource);
          this.pageCfg = configs;
          console.log(this.pageCfg);
          const tabs = configs.tabs || [];
          const tabsName = tabs.map(item => item.name);
          const tabConfigs = {};
          tabs.forEach(item => {
            tabConfigs[item.name] = item;
          });
          this.tabNameList = tabsName;
          this.tabConfigs = tabConfigs;
          this.operationList = configs.operations || [];
          for (let i = 0, iLen = data.length; i < iLen; i++) {
            const _task = data[i].task;
            if (['PREPARING', 'CONNECTING', 'CREATING', 'WAITING', 'INIT', 'STARTING'].includes(_task.status)
              && tabsName.indexOf('operationProfile') >= 0) {
              data[i].activeName = 'operationProfile-' + _task.id;
            }
            if (['RUNNING'].includes(_task.status) && tabsName.indexOf('logs') >= 0) {
              data[i].activeName = 'logs-' + _task.id;
            }
            if (data[i].activeName.indexOf('configInfo-') < 0) {
              setTimeout(() => {
                this.tabChange(data[i]);
              }, 80);
            }
          }
          this.mainData = data;
          cloudBrainTools.initRefreshData(this.mainData);
        } else {
          this.notFound = true;
        }
        this.emitUpdate();
      }).catch(err => {
        this.loading = false;
        this.notFound = true;
        console.log("xxxxxxxxxxx")
        console.log(err);
        this.emitUpdate();
      });
    }
  },
  beforeMount() { },
  mounted() {
    if (this.autoInit) {
      this.init();
    }
  },
};
</script>

<style scoped lang="less">
.task-detail-container {
  flex: 1;
  height: 0;

  .body-content {
    height: 100%;

    .content {
      height: 100%;
      position: relative;
      display: flex;
      flex-direction: column;

      .operation-area {
        display: flex;
        align-items: center;
        flex-wrap: wrap;

        .operate-item {
          margin: 0 16px 20px 0;
          min-width: 140px;

          .operate-btn-c {
            height: 40px;
            line-height: 20px;
            border-radius: 5px;
            background-color: rgba(50, 145, 248, 1);
            color: rgba(255, 255, 255, 1);
            font-size: 14px;
            text-align: center;
            box-shadow: 0px 2px 6px 0px rgba(50, 145, 248, 1);
            font-family: Roboto;
            border: 1px solid rgba(28, 126, 232, 1);
            display: flex;
            align-items: center;
            cursor: pointer;
            padding: 0 16px;

            /deep/.operate-btn,
            .operate-btn {
              margin-left: 6px;
              color: rgba(255, 255, 255, 1);
            }

            &:hover {
              opacity: 0.85;
            }

            &.btn-disabled {
              cursor: not-allowed;
              opacity: 1 !important;
              background-color: rgba(16, 16, 16, 0.1);
              color: rgba(16, 16, 16, 0.5);
              font-size: 14px;
              box-shadow: 0px 2px 6px 0px rgba(178, 178, 178, 1);
              border: 1px solid rgba(178, 178, 178, 1);
              cursor: not-allowed;
              pointer-events: none;

              &:hover {
                pacity: 1;
              }

              /deep/.operate-btn,
              .operate-btn {
                color: rgba(16, 16, 16, 0.5);
              }
            }
          }
        }

        .mobile-export-output {
          display: none;
        }
      }

      /deep/ .el-tabs__header {
        margin: 0;
      }

      .el-tab-pane {
        background-color: #fff;
        height: 100%;
      }

      .el-tabs {
        flex: 1;
        height: 0;
        display: flex;
        flex-direction: column;

        /deep/ .el-tabs__content {
          flex: 1;
        }
      }
    }

    .task-already-dlg {
      margin-top: 15vh;

      /deep/.el-dialog__header {
        display: none;
      }

      /deep/.el-dialog__body {
        padding: 0px 0px;
      }
    }

    .err-msg-box-already {
      padding: 1em 1.5em;
      background-color: rgba(242, 113, 28, 0.05);
      border: 2px solid rgba(242, 113, 28, 1);
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
  }
}

::-webkit-scrollbar {
  -webkit-appearance: none;
  width: 7px;
  height: 7px
}

::-webkit-scrollbar-track {
  background: transparent;
  border-radius: 7px !important;
}

::-webkit-scrollbar-thumb {
  background: rgb(210, 210, 216) !important;
  border-radius: 7px !important;
}

::-webkit-scrollbar-thumb:hover {
  background: rgba(210, 210, 216, 0.8) !important;
  border-radius: 7px !important;
}

::-webkit-scrollbar-thumb:active {
  background: rgba(210, 210, 216, 0.8) !important;
  border-radius: 7px !important;
}

@media only screen and (max-width: 768px) {
  .export-output {
    display: none;
  }

  .mobile-export-output {
    display: block !important;
    color: #888;
    margin-bottom: 12px;
  }
}
</style>
