<template>
  <div :class="isAdminPage ? 'admin-page' : 'custom-page'">
    <div v-if="!isAdminPage" class="cb-task-header">
      <div class="header-wrap">
        <div>
          <div class="title">{{ $t('notebook.sameTaskTips6') }}</div>
          <div class="tips" v-if="!isAdminPage" v-html="$t('cloudbrainObj.maxTaskTips')"></div>
        </div>
      </div>
    </div>
    <div v-if="isAdminPage" class="ui attached segment">
      <div class="ui form ignore-dirty" style="display: flex;justify-content: space-between;">
        <div class="ui action input">
          <input :value="conds.q" @input="onInput" @keyup.enter="search"
            :placeholder="$t('cloudbrainObj.searchTaskName')" />
          <button class="ui blue button" @click="search">{{ $t('repos.search') }}</button>
        </div>
        <div class="ui buttons">
          <button v-for="(item, index) in btnGroup" class="ui button" :key="item.time"
            :class="currentTime === item.time ? 'blue' : 'basic'" :disabled="item.disabled"
            @click="timeChange(item.time)">
            {{ item.value }}
          </button>
        </div>
      </div>
    </div>
    <div class="content">
      <div class="list-head">
        <div class="filters-c">
          <!-- <el-select v-model="conds.cluster" @change="changeConds">
            <el-option v-for="(item) in clusterList" :key="item.k" :value="item.k" :label="item.v"></el-option>
          </el-select> -->
          <el-select v-model="conds.aicenter" @change="changeConds" filterable>
            <el-option v-for="(item) in aicenterList" :key="item.k" :value="item.k" :label="item.v"></el-option>
          </el-select>
          <el-select v-model="conds.taskType" @change="changeConds" filterable>
            <el-option v-for="(item) in taskTypeList" :key="item.k" :value="item.k" :label="item.v"></el-option>
          </el-select>
          <el-select v-model="conds.computeResource" @change="changeConds" filterable>
            <el-option v-for="(item) in computeResourceList" :key="item.k" :value="item.k" :label="item.v"></el-option>
          </el-select>
          <el-select v-model="conds.status" @change="changeConds" filterable>
            <el-option v-for="(item) in statusList" :key="item.k" :value="item.k" :label="item.v"></el-option>
          </el-select>
        </div>
        <div class="right">
          <a v-if="isAdminPage" class="ui compact blue basic icon button"
            style="box-shadow: none !important; padding: 0.8em;"
            :href="`/api/v1/admin/ai_task/download/list?type=${currentTime}&_csrf=${csrf}`"><i
              class="ri-download-line middle aligned icon"></i>{{ $t('cloudbrainObj.downloadReport') }}</a>
          <template v-else>
            <div class="ui small icon input" style="height: 36px;">
              <input :value="conds.q" @input="onInput" @keyup.enter="search"
                :placeholder="$t('cloudbrainObj.searchTaskName')">
              <i class="search icon" style="cursor: pointer;pointer-events: auto" @click="search"></i>
            </div>
            <el-button class="op-btn" type="primary" @click="goCreate">
              <div class="btn-content">
                <i class="ri-add-box-line"></i>
                <span>{{ $t('cloudbrainObj.createNewAitask') }}</span>
              </div>
            </el-button>
          </template>

        </div>
      </div>
      <div class="list-body table-container" v-loading="loading">
        <el-table :data="tableData" class="top-align-table" v-loading="loading" style="min-width:100%"
          :cell-style="{ verticalAlign: 'top' }" @selection-change="handleSelectionChange" stripe>
          <el-table-column v-if="!isMiniScreen" type="selection" width="45" :selectable="isRowSelectable"
            fixed></el-table-column>
          <el-table-column :label="$t('cloudbrainObj.taskName')" align="left" header-align="center" width="240"
            :fixed="!isMiniScreen">
            <template slot-scope="scope">
              <div v-if="scope.row.editing">
                <el-input v-model="scope.row.editName" class="edit-input" size="small" ref="editInput"
                  @keyup.enter.native="saveEdit(scope.row)" @keyup.esc.native="cancelEdit(scope.row)">
                </el-input>
                <div class="edit-buttons">
                  <el-button type="primary" size="mini" @click="saveEdit(scope.row)">保存</el-button>
                  <el-button size="mini" @click="cancelEdit(scope.row)">取消</el-button>
                </div>
              </div>
              <div v-else>
                <a class="dispaly-job-name" :href="scope.row.task.jobLink"
                  @contextmenu.prevent="handleTaskNameContextMenu(scope.row, $event)">
                  {{ scope.row.task.display_job_name }}
                </a>
              </div>

            </template>
          </el-table-column>
          <el-table-column :label="$t('cloudbrainObj.cloudbrainTaskType')" align="center" header-align="center"
            width="120">
            <template slot-scope="scope">
              <div class="cb-job" :class="scope.row.task.job_type">
                <span>{{ scope.row.task.jobTypeShow }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="Status" :label="$t('status')" align="left" header-align="center" width="180">
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

          <el-table-column :label="$t('cloudbrainObj.ComputingResourceInfo')" align="center" header-align="center"
            width="235">
            <template slot-scope="scope">
              <!-- <div class="task-resource-wrap">
                <span class="label nowrap" :title="$t('cloudbrainObj.cluster')">{{$t('cloudbrainObj.cluster')}}:</span>
                <span class="value nowrap">{{scope.row.task.clusterName}}</span>
              </div>-->
              <div class="task-resource-wrap">
                <span class="label nowrap" :title="$t('resourcesManagement.aiCenter')">{{
                  $t('resourcesManagement.aiCenter') }}:</span>
                <span class="value nowrap">{{ scope.row.task.ai_center }}</span>
              </div>
              <div class="task-resource-wrap">
                <span class="label nowrap" :title="$t('cloudbrainObj.computeResource')">{{
                  $t('cloudbrainObj.computeResource') }}:</span>
                <span class="value nowrap">{{ scope.row.task.computeSourceShow }}</span>
              </div>
              <div class="task-resource-wrap">
                <span class="label nowrap" :title="$t('resourcesManagement.accCardType')">{{
                  $t('resourcesManagement.accCardType') }}:</span>
                <span class="value nowrap">{{ scope.row.task.accCardTypeShow }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="created_unix" :label="$t('cloudbrainObj.createTime')" align="center"
            header-align="center" width="140">
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
          <el-table-column v-if="isAdminPage" :label="$t('modelManage.creator')" align="left" width="80"
            header-align="center">
            <template slot-scope="scope">
              <div class="creator-wrap">
                <a :href="'/' + scope.row.creator.name" :title="scope.row.creator.full_name">
                  <img :src="scope.row.creator.rel_avatar_link">
                </a>
              </div>
            </template>
          </el-table-column>
          <el-table-column :label="$t('cloudbrainObj.repo')" align="center" header-align="center" width="250">
            <template slot-scope="scope">
              <a v-if="scope.row.task.repo_id !== 0 && scope.row.task.repoName" class="row-repo"
                :href="`/${scope.row.task.repoOwnerName}/${scope.row.task.repoName}`">
                {{ scope.row.task.repoOwnerName }} / <span style="font-weight: bold;">{{ scope.row.task.repoAlias
                  }}</span>
              </a>
              <span style="color:#888888" v-else-if="scope.row.task.repo_id !== 0 && !scope.row.task.repoName">--
                ({{ $t('repositoryWasDel') }})</span>
              <span v-else>--</span>
            </template>
          </el-table-column>
          <el-table-column v-if="isAdminPage" :label="$t('cloudbrainObj.cloudbrainTaskName')" align="center"
            header-align="center" width="200">
            <template slot-scope="scope">
              <span class="ui poping up clipboard" :id="`clipboard-${scope.row.task.jobNameShow}`"
                data-position="top center" data-variation="inverted tiny" :data-success="$t('copySuccess')"
                :data-content="$t('copy')" :data-original="$t('copy')"
                :data-clipboard-text="scope.row.task.jobNameShow">
                <span :title="scope.row.task.jobNameShow" style="cursor:pointer;">
                  {{ scope.row.task.jobNameShow }}
                </span>
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="operation" :label="$t('operation')" align="left" :min-width="isMiniScreen ? 170 : 315"
            header-align="center" fixed="right">
            <template slot-scope="scope">
              <div class="op-wrap">
                <OperationBtnGroup :usekey="scope.row.task.id" :key="scope.row.task.id">
                  <a href="javascript:;"
                    v-if="scope.row.task.isSubscriber && scope.row.task.operations.indexOf('consoleHome') >= 0"
                    @click="opDebug(scope.row)"
                    :class="scope.row.can_modify && scope.row.task.canDebug ? '' : 'disabled'">{{
                      $t('consoleHome')
                    }}</a>
                  <a href="javascript:;"
                    v-if="scope.row.task.isAdmin && scope.row.task.operations.indexOf('consoleHome1') >= 0"
                    @click="opDebug(scope.row)"
                    :class="scope.row.can_modify && scope.row.task.canDebug ? '' : 'disabled'">{{
                      $t('consoleHome')
                    }}</a>
                  <a href="javascript:;"
                    v-if="scope.row.task.operations.indexOf('onlineInfer') >= 0 && scope.row.task.canDebug"
                    @click="opVisual(scope.row)"
                    :class="scope.row.can_modify && scope.row.task.canDebug ? '' : 'disabled'">{{
                      $t('onlineinfer')
                    }}</a>
                  <a href="javascript:;" v-if="scope.row.task.operations.indexOf('onlinexperience') >= 0"
                    @click="opOnline(scope.row)"
                    :class="scope.row.can_experience && scope.row.task.canDebug ? '' : 'disabled'">{{
                      $t('modelManage.onlineInference')
                    }}</a>
                  <a href="javascript:;" v-if="scope.row.task.operations.indexOf('onlineLoraTrain') >= 0"
                    :class="scope.row.can_experience && scope.row.task.canDebug ? '' : 'disabled'"
                    @click="opLoraTrain(scope.row)">{{ $t('modelManage.onlineLoraTrain') }}</a>
                  <a href="javascript:;" v-if="scope.row.task.operations.indexOf('onlineComfyUI') >= 0"
                    :class="scope.row.can_experience && scope.row.task.canDebug ? '' : 'disabled'"
                    @click="opComfyUI(scope.row)">{{ $t('modelManage.onlineWorkflow') }}</a>
                  <a href="javascript:;"
                    v-if="scope.row.task.operations.indexOf('debug') >= 0 && scope.row.task.canDebug"
                    @click="opDebug(scope.row)"
                    :class="scope.row.can_modify && scope.row.task.canDebug ? '' : 'disabled'">{{
                      $t('cloudbrainObj.debug')
                    }}</a>
                  <a href="javascript:;"
                    v-if="scope.row.task.operations.indexOf('start') >= 0 && scope.row.can_modify && scope.row.task.canDebug"
                    @click="opDebug(scope.row)"
                    :class="scope.row.can_modify && scope.row.task.canDebug ? '' : 'disabled'">{{
                      $t('cloudbrainObj.startUse')
                    }}</a>
                  <a href="javascript:;"
                    v-if="scope.row.task.operations.indexOf('redebug') >= 0 && !scope.row.task.canDebug && !scope.row.task.is_file_notebook"
                    @click="opReDebug(scope.row)"
                    :class="scope.row.can_modify && scope.row.task.canReDebug ? '' : 'disabled'">{{
                      $t('cloudbrainObj.reDebug') }}</a>
                  <a href="javascript:;"
                    v-if="scope.row.task.operations.indexOf('reinfer') >= 0 && !scope.row.task.canDebug && !scope.row.task.is_file_notebook"
                    @click="opReDebug(scope.row)"
                    :class="scope.row.can_modify && scope.row.task.canReDebug ? '' : 'disabled'">{{
                      $t('cloudbrainObj.reInfer') }}</a>
                  <a href="javascript:;"
                    v-if="scope.row.task.operations.indexOf('deploy') >= 0 && scope.row.task.can_fintune_experience"
                    @click="opDeploy(scope.row)">{{
                      $t('cloudbrainObj.deploymentExperience') }}</a>
                  <a href="javascript:;" v-if="scope.row.task.operations.indexOf('stop') >= 0"
                    @click="opStop(scope.row)"
                    :class="scope.row.can_delete && scope.row.task.canStop ? '' : 'disabled'">{{
                      $t('cloudbrainObj.stop')
                    }}</a>
                  <a href="javascript:;" v-if="scope.row.task.operations.indexOf('modify') >= 0 && scope.row.task.can_modify
                    && !scope.row.task.is_fine_tune_task && !scope.row.task.is_file_notebook"
                    @click="opModify(scope.row)"
                    :class="scope.row.can_modify && scope.row.task.canModify ? '' : 'disabled'">{{
                      $t('cloudbrainObj.modify')
                    }}</a>
                  <!--  comfyui 专属  -->
                  <a href="javascript:;" v-if="scope.row.task.operations.indexOf('modifyComfyUI') >= 0"
                    @click="opModify(scope.row)" :class="scope.row.can_experience ? '' : 'disabled'">{{
                      $t('cloudbrainObj.modify')
                    }}</a>
                  <template v-if="['ComfyuiExperience'].includes(scope.row.task.job_type)">
                    <a target="_blank" v-if="scope.row.can_experience && scope.row.task.canSaveImage"
                      :href="scope.row.task.saveImageUrl">
                      {{ $t('cloudbrainObj.commitImage') }}
                    </a>
                  </template>
                  <!-- -->
                  <template v-else>
                    <a target="_blank" v-if="scope.row.task.can_modify && scope.row.task.canSaveImage"
                      :href="scope.row.task.saveImageUrl">
                      {{ $t('cloudbrainObj.commitImage') }}
                    </a>
                  </template>
                  <a href="javascript:;" v-if="scope.row.task.operations.indexOf('delete') >= 0"
                    @click="opDelete(scope.row)"
                    :class="scope.row.can_delete && scope.row.task.canDelete ? '' : 'disabled'">{{
                      $t('cloudbrainObj.delete') }}</a>
                  <a href="javascript:;" v-if="scope.row.can_create_template && scope.row.task.canSaveTmpl"
                    @click="opSaveTmpl(scope.row)">{{ $t('cloudbrainObj.saveTaskTmpl') }}</a>
                </OperationBtnGroup>
              </div>
            </template>
          </el-table-column>
        </el-table>
        <!-- 右键菜单 -->
      </div>
      <div class="list-foot" v-if="tableData.length">
        <el-button class="left-btn" plain icon="el-icon-delete" @click="batchDelete">{{ $t('cloudbrainObj.batchDelete')
          }}</el-button>
        <el-pagination class="center-btn" ref="paginationRef" background :small="isMiniScreen"
          @current-change="currentChange" @size-change="sizeChange" :current-page.sync="page" :page-sizes="pageSizes"
          :page-size.sync="pageSize" layout="total, sizes, prev, pager, next, jumper" :total="total">
        </el-pagination>
      </div>
    </div>
    <LoadingMask :loading="maskLoading" :content="maskLoadingContent"></LoadingMask>
    <el-dialog :visible.sync="dialogVisible" width="25%" top="35vh" :show-close="false" :close-on-click-modal="false"
      :close-on-press-escape="false" custom-class="del-container">
      <p v-for="(delItem, index) in delFaildList" :key="index">{{ delItem.task.display_job_name }} {{
        $t('cloudbrainObj.deleteFailed') }}；</p>
      <span slot="footer">
        <el-button @click="dialogVisible = false" type="info"
          style="color: #020004;background-color: #c2c7cc;border-color: #c2c7cc;">{{
            $t('cloudbrainObj.dialogTips.tips8')
          }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import LoadingMask from '~/components/cloudbrain/LoadingMask.vue';
import OperationBtnGroup from '~/components/cloudbrain/OperationBtnGroup.vue';
import { timeSinceUnix, getListValueWithKey, initClipboard } from '~/utils';
import { CloudBrainTools } from '../tools';
import { CLUSTERS, COMPUTER_RESOURCES, JOB_TYPE, ACC_CARD_TYPE, BenchmarkTypeList } from '~/const';
import {
  getMyAiTasks, getAdminAiTasks, getAiTaskDebugUrl, getAiTaskRestart, stopAiTask, deleteAiTask,
  getAiTask, getTmplEditAdress, getAiEndpointUrl, deleteMulAiTask
} from '~/apis/modules/cloudbrain';
import { getAiCenterList, commonFormPost, getComfyuiUrl } from '~/apis/modules/common';

import relativeTime from 'dayjs/plugin/relativeTime';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import 'dayjs/locale/zh-cn';
import 'dayjs/locale/en';
import dayjs from 'dayjs';
import { lang } from '~/langs';

dayjs.locale(lang == 'zh-CN' ? 'zh-cn' : 'en');
dayjs.extend(relativeTime);
dayjs.extend(localizedFormat);

const cloudBrainTools = new CloudBrainTools();

const OperationsMap = {
  DEBUG: ['debug', 'redebug', 'stop', 'modify', 'delete'],
  TRAIN: ['stop', 'modify', 'delete'],
  INFERENCE: ['stop', 'modify', 'delete'],

  BENCHMARK: ['stop', 'delete'],
  SIM2BRAIN_SNN: ['stop', 'delete'],
  SNN4ECOSET: ['stop', 'delete'],
  SNN4IMAGENET: ['stop', 'delete'],
  BRAINSCORE: ['stop', 'delete'],
  MODELSAFETY: ['stop', 'delete'],

  ONLINEINFERENCE: ['consoleHome', 'onlineInfer', 'reinfer', 'stop', 'modify', 'delete'],
  HPC: ['start', 'stop', 'modify', 'delete'],
  GENERAL: ['debug', 'stop', 'modify', 'delete'],
  MODELEXPERIENCE: ['onlinexperience', 'stop', 'delete'],
  FINETUNE: ['deploy', 'stop', 'delete'],
  SDFINETUNE: ['consoleHome1', 'onlineLoraTrain', 'stop', 'delete'],
  ComfyuiExperience: ['consoleHome1', 'onlineComfyUI', 'stop', 'modifyComfyUI', 'delete'],
  EVAL: ['stop', 'delete'],
};

const UseJobStatusList = ['STARTING', 'RUNNING', 'RESTARTING', 'START_FAILED', 'STOPPING', 'STOPPED',
  'WAITING', 'COMPLETED', 'SUCCEEDED', 'FAILED', 'KILLED', 'CREATED_FAILED'
];

const btnGroup = [
  { disabled: false, time: 'today', value: '今天' },
  { disabled: false, time: 'yesterday', value: '昨日' },
  { disabled: false, time: 'last_7day', value: '近7天' },
  { disabled: false, time: 'current_month', value: '本月' },
  { disabled: false, time: 'last_month', value: '上月' },
  { disabled: false, time: 'last_30day', value: '近30天' },
  { disabled: false, time: 'current_year', value: '今年' },
  { disabled: false, time: 'last_year', value: '去年' },
  { disabled: false, time: 'all', value: '所有' },
]
const jobLinkGenerators = {
  MODELEXPERIENCE: (id) => `/modelbase/experience/detail/${id}`,
  FINETUNE: (id) => `/modelbase/nlp/sft/detail/${id}`,
  SDFINETUNE: (id) => `/modelbase/cv/sft/detail/${id}`,
  ComfyuiExperience: (id) => `/modelbase/cv/comfyui/detail/${id}`,
  EVAL: (id) => `/modelbase/eval/evaluate/detail/${id}`,
  DEFAULT: (id) => `/cloudbrains/detail/${id}`
};

const { csrf } = window.config;
export default {
  data() {
    return {
      isAdminPage: true,
      conds: {
        q: '',
        cluster: '',
        aicenter: '',
        taskType: '',
        computeResource: '',
        status: '',
      },
      clusterList: [{ k: '', v: this.$t('resourcesManagement.allCluster') }, ...CLUSTERS],
      aicenterList: [{ k: '', v: this.$t('resourcesManagement.allAiCenter') }],
      taskTypeList: [{ k: '', v: this.$t('resourcesManagement.allJobType') }],
      computeResourceList: [{ k: '', v: this.$t('resourcesManagement.allComputeResource') }, ...COMPUTER_RESOURCES],
      statusList: [{ k: '', v: this.$t('resourcesManagement.allJobStatus') }, ...[...UseJobStatusList, 'other']
        .map(item => ({ k: item, v: item.toLocaleUpperCase() }))],
      tableData: [],
      page: 1,
      pageSize: 10,
      total: 0,
      pageSizes: [10, 20, 30, 50],
      loading: false,
      maskLoading: false,
      maskLoadingContent: '',
      clipboardHandler: null,
      isMiniScreen: false,
      multipleSelection: [],
      dialogVisible: false,
      delFaildList: [],
      currentTime: 'last_30day',
      btnGroup: btnGroup,
      csrf: csrf,
      searchValue: '',
    };
  },
  components: { LoadingMask, OperationBtnGroup },
  watch: {
    'searchValue': function (val) {
      if (val === '') {
        setTimeout(() => {
          this.search();
        }, 100);
      }
    }
  },
  methods: {
    onInput(event) {
      this.searchValue = event.target.value;
    },
    timeChange(time) {
      if (this.currentTime === time) return
      this.currentTime = time
      this.page = 1
      this.getTableData()
    },
    goCreate() {
      location.href = '/cloudbrains/create'
    },
    getTableData() {
      this.loading = true;
      const params = {
        job_type: this.conds.taskType,
        job_status: this.conds.status,
        ai_center: this.conds.aicenter,
        cluster: this.conds.cluster == 'OpenI' ? 'resource_cluster_openi' : this.conds.cluster == 'C2Net' ? 'resource_cluster_c2net' : this.conds.cluster,
        compute_source: this.conds.computeResource,
        q: this.conds.q,
        page: this.page,
        pageSize: this.pageSize,
        exclude_status: this.conds.status == 'other' ? UseJobStatusList.join(',') : undefined,
        type: this.currentTime
      };
      console.log("params", params)
      const getApi = this.isAdminPage ? getAdminAiTasks : getMyAiTasks;
      getApi(params).then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 0) {
          const data = res.data;
          this.canCreateTask = data.can_create_task;
          let isSubscriber = data.is_subscriber;
          let isAdmin = data.is_admin;
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
            task.repoOwnerName = item.owner_name;
            task.repoName = item.repo_name;
            task.repoAlias = item?.repo_alias || item.repo_name;
            const generateJobLink = jobLinkGenerators[task.job_type] || jobLinkGenerators.DEFAULT;
            task.jobLink = generateJobLink(task.id);
            task.createdFromNow = this.calcFromNow(task.created_unix);
            task.jobNameShow = task.job_name;
            task.jobTypeShow = BenchmarkTypeList.indexOf(task.job_type) >= 0 ? this.$t('benchmarkTask') : getListValueWithKey(JOB_TYPE, task.job_type);
            task.operations = OperationsMap[task.job_type] || [];
            task.isSubscriber = isSubscriber && task.end_point
            task.isAdmin = isAdmin

            cloudBrainTools.checkOperation(task);
          });
          if (data.tasks && data.tasks.length > 0) {
            this.tableData = data.tasks.map(row => {
              return {
                ...row,
                editing: false,
                editName: '',
                originalName: '',
              }
            })
          } else {
            this.tableData = []
          }
          cloudBrainTools.initRefreshData(this.tableData);
          console.log(this.tableData)

        }
        this.$nextTick(() => {
          this.clipboardHandler && this.clipboardHandler.destroy();
          this.clipboardHandler = initClipboard();
        });
      }).catch(err => {
        this.loading = false;
        this.$message.error(err)
      });
    },
    search() {
      this.conds.q = this.searchValue;
      this.page = 1;
      this.getTableData();
    },
    handleSelectionChange(taskList) {
      this.multipleSelection = []
      taskList.forEach((item) => {
        this.multipleSelection.push(item.task.id)
      })
    },
    isRowSelectable(row, rowIndex) {
      return row.can_delete && row.task.canDelete
    },
    changeConds() {
      this.page = 1;
      this.getTableData();
    },
    currentChange(page) {
      this.page = page;
      this.getTableData();
    },
    sizeChange(pageSize) {
      this.page = 1;
      this.pageSize = pageSize;
      this.getTableData();
    },
    calcFromNow(unix) {
      return timeSinceUnix(unix, Date.now() / 1000);
    },
    dateFormat(unix) {
      return dayjs(unix * 1000).format('YYYY-MM-DD HH:mm:ss');
    },
    calcScreenInfo() {
      this.isMiniScreen = document.documentElement.clientWidth <= 800;
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
    // operations
    opDebug(row) {
      if (this.operating) return;
      this.operating = true;
      getAiTaskDebugUrl({
        repoOwnerName: row.task.repoOwnerName,
        repoName: row.task.repoName,
        id: row.task.id,
      }).then(res => {
        this.operating = false;
        res = res.data;
        if (res.code == 0) {
          if (!!(res.data?.ptoken)) {
            this.scowFunc(res.data.url, res.data.ptoken)
            return
          }
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
        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      });
    },
    opVisual(row) {
      if (this.operating) return;
      this.operating = true;
      getAiEndpointUrl({
        repoOwnerName: row.task.repoOwnerName,
        repoName: row.task.repoName,
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
        repoOwnerName: row.task.repoOwnerName,
        repoName: row.task.repoName,
        id: row.task.id,
      }).then(res => {
        this.operating = false;
        res = res.data;
        if (res.code == 0) {
          this.page = 1;
          this.getTableData();
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
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
    opDeploy(row) {
      const item = row.task
      let url = `/modelbase/experience/create?id=${item.id}&job_name=${item.display_job_name}&compute_resource=${item.compute_source}&modelName=${item.app_name}`;
      window.location.href = url;
    },
    opStop(row) {
      if (this.operating) return;
      this.operating = true;
      if (BenchmarkTypeList.indexOf(row.task.job_type) >= 0) {
        let url = '';
        if (row.task.job_type == 'MODELSAFETY') {
          url = `/${row.task.repoOwnerName}/${row.task.repoName}/modelsafety/${row.task.id}/stop`;
        } else if (row.task.job_type == 'BENCHMARK') {
          url = `/${row.task.repoOwnerName}/${row.task.repoName}/cloudbrain/benchmark/${row.task.id}/stop`;
        } else {
          url = `/${row.task.repoOwnerName}/${row.task.repoName}/cloudbrain/${row.task.id}/stop`;
        }
        url = url + (this.isAdminPage ? '?isadminpage=true' : '');
        const formData = { id: row.task.id };
        commonFormPost(url, formData).then(res => {
          this.operating = false;
          row.task.createdFromNow = this.calcFromNow(row.task.created_unix);
          cloudBrainTools.checkOperation(row.task);
          this.getTableData();
        }).catch(err => {
          this.operating = false;
          this.$message({
            type: 'error',
            message: this.$t('operationFailed'),
          });
        });
      } else {
        const experienceParams = {}
        if (row.task.job_type === "MODELEXPERIENCE") {
          experienceParams.model_experience = true
        }
        stopAiTask({
          repoOwnerName: row.task.repoOwnerName,
          repoName: row.task.repoName,
          id: row.task.id,
        }, experienceParams).then(res => {
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
      }
    },
    opModify(row) {
      if (this.operating) return;
      const task = row.task;
      let url = `/cloudbrains/create`;
      if (task.job_type === 'ComfyuiExperience') {
        url = `/modelbase/cv/comfyui/create`;
      }
      window.location.href = url + (url.indexOf('?') >= 0 ? '&' : '?') + `modify=true&id=${task.id}&backurl=${encodeURIComponent(window.location.href)}`;
    },
    opDelete(row) {
      if (this.operating) return;
      this.$confirm(this.$t('cloudbrainObj.deleteConfirmTips'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        this.operating = true;
        this.maskLoading = true;
        this.maskLoadingContent = this.$t('cloudbrainObj.deletingTips');
        if (BenchmarkTypeList.indexOf(row.task.job_type) >= 0
          || (row.task.job_type == 'INFERENCE' && row.task.cluster == 'OpenI'
            && (row.task.compute_source == 'GPU' || row.task.compute_source == 'CPU/GPU'))
        ) {
          let url = '';
          if (row.task.job_type == 'MODELSAFETY') {
            url = `/${row.task.repoOwnerName}/${row.task.repoName}/modelsafety/${row.task.id}/del`;
          } else if (row.task.job_type == 'BENCHMARK') {
            url = `/${row.task.repoOwnerName}/${row.task.repoName}/cloudbrain/benchmark/${row.task.id}/del`;
          } else {
            url = `/${row.task.repoOwnerName}/${row.task.repoName}/cloudbrain/${row.task.id}/del`;
          }
          url = url + (this.isAdminPage ? '?isadminpage=true' : '');
          const formData = { id: row.task.id };
          commonFormPost(url, formData).then(res => {
            this.operating = false;
            this.maskLoading = false;
            if (this.total % this.pageSize == 1 && this.page > 1) {
              this.page -= 1;
            }
            this.getTableData();
          }).catch(err => {
            this.operating = false;
            this.maskLoading = false;
            this.$message({
              type: 'error',
              message: this.$t('operationFailed'),
            });
          });
        } else {
          deleteAiTask({
            repoOwnerName: row.task.repoOwnerName,
            repoName: row.task.repoName,
            id: row.task.id,
          }).then(res => {
            this.operating = false;
            this.maskLoading = false;
            if (this.total % this.pageSize == 1 && this.page > 1) {
              this.page -= 1;
            }
            this.getTableData();
          }).catch(err => {
            this.maskLoading = false;
            this.operating = false;
            this.$message({
              type: 'error',
              message: this.$t('operationFailed'),
            });
          });
        }
      }).catch(() => {
        this.$message({
          type: 'info',
          message: this.$t('cancelOperate'),
        });
      });
    },
    batchDelete() {
      if (this.operating) return;
      if (this.multipleSelection.length === 0) return
      this.$confirm(this.$t('cloudbrainObj.deleteBatchConfirmTips'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        this.operating = true;
        this.maskLoading = true;
        this.maskLoadingContent = this.$t('cloudbrainObj.deletingTips');
        deleteMulAiTask({ id: this.multipleSelection }).then(res => {
          if (res.data.code === 99) {
            this.delFaildList = this.tableData.filter((item) => {
              return res.data.data.includes(item.task.id)
            })

            this.dialogVisible = true
          }

          this.operating = false;
          this.maskLoading = false;
          let deLength = this.multipleSelection.length - this.delFaildList.length
          // if (this.page === Math.ceil((this.total / this.pageSize)) && deLength === this.total % this.pageSize) {
          //   this.page -= 1
          // }
          if (this.shouldDecrementPage(deLength)) {
            this.page -= 1
          }
          this.getTableData();
        }).catch(err => {
          this.maskLoading = false;
          this.operating = false;
          this.$message({
            type: 'error',
            message: this.$t('operationFailed'),
          });
        });
      }).catch(() => {
        this.$message({
          type: 'info',
          message: this.$t('cancelOperate'),
        });
      });
    },
    shouldDecrementPage(deletedCount) {
      const newTotal = this.total - deletedCount;
      const newTotalPages = Math.ceil(newTotal / this.pageSize);
      return this.page > newTotalPages && this.page > 1;
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
        if (!res.data.url) {
          this.$message.warning(this.$t('modelSquare.ComfyUiWarn'))
        } else {
          if (!window.open(res.data.url)) {
            window.location.href = res.data.url;
          }
        }
      })
    },
    opSaveTmpl(row) {
      if (this.operating) return;
      window.location.href = `/ai_task_tmpl/create?task=${row.task.id}`;
    }
  },
  beforeMount() {
    this.isAdminPage = window.location.href.indexOf('admin/cloudbrains') >= 0;
    this.calcScreenInfo();
    getAiCenterList().then(res => {
      res = res.data;
      if (res.Code == 0) {
        const data = res.Data || [];
        data.forEach(item => {
          this.aicenterList.push({
            k: item.AiCenterCode,
            v: item.AiCenterName,
          });
        });
      }
    }).catch(err => {
      this.$message.error(err)
    });
    this.getTableData();
    if (this.isAdminPage) {
      this.taskTypeList = [
        { k: '', v: this.$t('resourcesManagement.allJobType') },
        ...JOB_TYPE
      ]
    } else {
      this.taskTypeList = [
        { k: '', v: this.$t('resourcesManagement.allJobType') },
        ...JOB_TYPE.filter(item => ['BENCHMARK'].indexOf(item.k) < 0)
      ]
    }
  },
  mounted() {
    window.addEventListener('resize', this.calcScreenInfo);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.calcScreenInfo);
  },
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';

.cb-task-header {
  margin: 30px 30px 0 40px;

  .header-wrap {
    display: flex;
    justify-content: space-between;
    align-items: flex-end;

    .title {
      color: rgb(16, 16, 16);
      font-size: 18px;
      font-weight: 700;
      padding: 6px 0;
    }

    .tips {
      margin-top: 8px;
      color: #888888;
      font-size: 14px;

      /deep/ span {
        color: #ff6200;
      }
    }

    .adv-space {
      min-width: 446px;
      height: 54px;
      background-image: url('/img/svg-icon/cb-adv-right.svg');
      display: flex;
      align-items: center;

      img {
        margin-bottom: 12px;
        margin-left: -14px;
      }

      .adv-text {
        line-height: 40px;
        color: rgba(255, 255, 255, 1);
        font-size: 18px;
        text-align: left;
        font-weight: 700;
        margin-left: -12px;
      }

      .adv-btn {
        background-color: rgba(255, 255, 255, 1);
        color: rgba(255, 48, 67, 1);
        border-radius: 5px;
        font-size: 14px;
        font-weight: 700;
        height: 26px;
        line-height: 26px;
        padding: 0 10px;
        margin-left: 20px;
      }
    }
  }
}

.content {
  margin: 16px 30px 0 40px;

  .list-head {
    margin-top: 12px;
    margin-bottom: 30px;
    display: flex;
    justify-content: space-between;

    .filters-c {
      .el-select {
        margin-right: 12px;
        height: 36px;
        width: 200px;

        /deep/.el-input--small {
          height: 36px;
        }

        /deep/.el-input__inner {
          height: 36px;

          &:visited {
            border-color: #85b7d9;
          }

          &:focus {
            border-color: #85b7d9;
          }

          &:active {
            border-color: #85b7d9;
          }
        }
      }
    }

    .right {
      display: flex;

      .op-btn {
        margin-left: 14px;
        display: flex;
        align-items: center;
        height: 36px;
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
    }
  }
}

.table-container {
  /deep/ .el-table__body {
    td {
      height: 64px;
    }
  }

  /deep/ .el-table__header {
    th {
      border: 1px solid rgb(208, 208, 208);
      border-right: 0;

      &:last-child {
        border-right: 1px solid rgb(208, 208, 208);
      }
    }
  }

  .edit-buttons {
    display: flex;
    margin-top: 6px;
    justify-content: flex-end;
  }

  .context-menu {
    position: fixed;
    z-index: 3000;
    border-radius: 5px;
    background-color: rgba(255, 255, 255, 1);
    box-shadow: 0px 2px 12px 0px rgba(0, 0, 0, 0.1);
    border: 1px solid rgba(235, 235, 235, 1);
    padding: 10px 8px;
    color: rgba(16, 16, 16, 1);
    font-size: 12px;

    .context-menu-item {
      height: 24px;
      display: flex;
      align-items: center;
      padding: 0 6px;

      i {
        color: rgba(11, 63, 247, 0.88);
        margin-right: 4px;
        font-size: 16px;
      }

      &:hover {
        background-color: rgba(240, 240, 240, 1);
        ;
      }
    }
  }

  .dispaly-job-name {
    font-size: 14px;
    font-weight: bold;
    color: rgb(0, 92, 255);
  }

  .status-wrap {
    display: flex;
    align-items: center;
    margin-left: 20px;

    i {
      margin-right: 4px;
      flex-shrink: 0;

      &:last-child {
        margin-left: 4px;
      }
    }
  }

  .task-resource-wrap {
    display: flex;

    .label {
      color: rgba(16, 16, 16, 0.6);
      width: 80px;
      text-align: right;
    }

    .value {
      margin-left: 8px;
      color: rgba(16, 16, 16, 1);
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

  .row-repo {
    font-size: 14px;
    color: rgb(0, 92, 255);
  }

  .op-wrap {
    display: flex;
    align-items: center;
    justify-content: center;

    a {
      margin: 0 10px;
      font-size: 14px;
      color: #005cff;

      &.disabled {
        color: rgba(0, 0, 0, .87);
      }
    }

    .el-dropdown {
      margin: 0 10px;
      color: rgba(0, 0, 0, .87);
      cursor: pointer;
    }
  }
}

.list-foot {
  display: flex;
  justify-content: space-between;
  /* 初始设置为两端对齐 */
  align-items: center;
  /* 垂直方向居中 */
  margin-top: 20px;
  margin-bottom: 20px;

  .left-btn {
    // margin-right: auto;
    border-color: #ff4d4f;
    color: #ff4d4f;
  }

  .center-btn {
    margin-right: auto;
    margin-left: auto;
  }
}

/deep/ .del-container {
  min-height: 240px;
  background-image: url('/img/model/del_failed.png');
  background-repeat: no-repeat;
  /* 背景图像不重复 */
  background-position: right bottom;
  display: flex;
  flex-direction: column;

  .el-dialog__body {
    flex: 1;
  }
}

/deep/ .el-dropdown-menu__item {
  a {
    color: #1678c2;
  }

  &.is-disabled {
    color: rgba(0, 0, 0, .87);
    pointer-events: none;
    opacity: 0.45 !important;

    a {
      color: rgba(0, 0, 0, .87);
      pointer-events: none;
      opacity: 0.45 !important;
    }
  }
}

.admin-page {
  border: 1px solid #d4d4d5;
  margin-top: 0px;
  padding-top: 0;
  box-shadow: 0 1px 2px 0 rgba(34, 36, 38, 0.15);

  .content {
    margin: 10px 10px;

    .list-head {
      margin-top: 16px;
      margin-bottom: 2px;

      .right {
        margin-bottom: 8px;
        min-width: 120px;
      }
    }
  }
}

.custom-page {}

@media only screen and (max-width: 1200px) {
  .adv-space {
    display: none !important;
  }
}

@media only screen and (max-width: 800px) {
  .custom-page {
    .content {
      margin: 10px 10px;

      .list-head {
        margin-bottom: 14px !important;

        .filters-c {
          display: none;
        }
      }

      .table-container {
        .op-wrap {
          a {
            margin: 0 4px;

            &.disabled {
              // display: none;
            }
          }
        }
      }

      .list-foot {
        .left-btn {
          display: none;
        }

        .el-pagination {

          /deep/ .el-pagination__total,
          /deep/ .el-pagination__sizes,
          /deep/ .el-pagination__jump {
            display: none;
          }
        }
      }
    }
  }
}
</style>
<style>
@media only screen and (max-width: 800px) {
  .el-message-box__wrapper .el-message-box {
    width: calc(100% - 20px);
  }

  .el-popup-parent--hidden {
    padding-right: 0px !important;
  }

  .explore.users {
    margin-right: 0px !important;
  }
}
</style>
