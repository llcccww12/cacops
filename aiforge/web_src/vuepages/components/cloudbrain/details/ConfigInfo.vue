<template>
<div class="item-container">
  <div class="layout-wrapper">
    <!-- 左侧：三个字段组 -->
    <div class="left-column">
      <div v-if="getCurrentFields('basicFiled').length > 0" class="field-group">
        <div class="field-title">{{$t('cloudbrainObj.basicInfo')}}</div>
        <div class="field-divider"></div>
        <div class="field-grid">
            <div v-for="fieldKey in getCurrentFields('basicFiled')" :key="fieldKey" class="field-item">
                <label :title="renderTitle(fieldKey)">{{ renderTitle(fieldKey) }}:</label>
                <span class="nowrap" v-html="renderContent(fieldKey)"></span>
            </div>
        </div>
      </div>
    
      <div v-if="getCurrentFields('resourceFiled').length > 0" class="field-group">
        <div class="field-title">{{$t('cloudbrainObj.paramsSetting')}}</div>
        <div class="field-divider"></div>
        <div class="field-grid">
            <div v-for="fieldKey in getCurrentFields('resourceFiled')" :key="fieldKey" class="field-item">
                <label :title="renderTitle(fieldKey)">{{ renderTitle(fieldKey) }}:</label>
                <span class="nowrap" v-html="renderContent(fieldKey)"></span>
            </div>
        </div>
      </div>
    
      <div v-if="getCurrentFields('paramsFiled').length > 0" class="field-group">
        <div class="field-title">{{$t('cloudbrainObj.resourceSetting')}}</div>
        <div class="field-divider"></div>
        <div class="field-grid">
            <div v-for="fieldKey in getCurrentFields('paramsFiled')" :key="fieldKey" class="field-item">
                <label :title="renderTitle(fieldKey)">{{ renderTitle(fieldKey) }}:</label>
                <span class="nowrap" v-html="renderContent(fieldKey)"></span>
            </div>
        </div>
      </div>
    </div>
    
    <!-- 右侧：运行状态 -->
    <div class="right-column" >
      <div v-if="getCurrentFields('runstatusFiled').length > 0"  class="field-group">
        <div class="field-title">{{$t('cloudbrainObj.runningStatus')}}</div>
        <div class="field-divider"></div>
        <div class="field-grid">
          <div v-for="fieldKey in getCurrentFields('runstatusFiled')" :key="fieldKey" class="field-item">
            <label :title="renderTitle(fieldKey)">{{ renderTitle(fieldKey) }}:</label>
            <span class="nowrap" v-html="renderContent(fieldKey)"></span>
          </div>
        </div>
      </div>

      <div v-if="getCurrentFields('showSdkCode').length > 0" class="field-group item-block item-model-code">
        <div class="field-title">{{ $t('cloudbrainObj.sdkUseWay') }}</div>
        <div class="field-grid">
          <div class="code-content-c">
            <div class="code-content">
              <pre><code class="python hljs" v-html="renderHljs(data.task.sdk_code)"></code></pre>
            </div>
            <div class="copy-btn">
              <a href="javascript:;" class="ui poping up clipboard"
                :id="`clipboard-${sparkMD5Hash(data.task.sdk_code)}`" data-position="top center"
                data-variation="inverted tiny" :data-success="$t('copySuccess')" :data-content="$t('copy')"
                :data-original="$t('copy')" :data-clipboard-text="data.task.sdk_code">
                <i style="font-size:14px;" class="copy outline icon"></i>
              </a>
            </div>
          </div>
        </div>
      </div>
      <div v-if="getCurrentFields('generalTaskCodeTips').length > 0 && data.can_modify"
        class="field-group item-block item-model-code">
        <div class="field-title">{{ $t('cloudbrainObj.generalTaskSdkCodeTip0') }}</div>
        <div class="field-grid">
          <GeneralTaskCodeTips :isTaskDetail="true"></GeneralTaskCodeTips>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import { getPromoteData } from '~/apis/modules/common';
import { getAiTaskBrief } from '~/apis/modules/cloudbrain';
import GeneralTaskCodeTips from '../GeneralTaskCodeTips.vue'
import { renderSpecObject, initClipboard, escapeHTML } from '~/utils';
import { i18n } from '~/langs';
import { formatDate } from 'element-ui/lib/utils/date-util';
import SparkMD5 from "spark-md5";
import hljs from 'highlight.js';

export default {
  name: 'ConfigInfo',
  props: {
    configs: { type: Object, default: () => { return {} } },
    data: { type: Object, default: () => { return {} } },
    nosdkcode: { type: Boolean, default: false },
    isSubscriber: { type: Boolean, default: false }
  },
  components: { GeneralTaskCodeTips },
  data() {
    return {
      clipboardHandler: null,
      sourceFtList: [],
      flag: false,
    };
  },
  watch: {
    data: {
      immediate: true,
      deep: true,
      async handler(newVal) {
        if(this.data.task && this.data.task.source_id && !this.flag){
          const res = await getPromoteData('model/modelfinetune.json')
          const data = JSON.parse(res.data);
          let repoName = data.llm.repo_name
          let repoOwnerName = data.llm.repo_owner_name
          
          this.sourceFtList.push({
            name: this.data.task.app_name,
            url: `/modelbase/nlp/sft/detail/${repoOwnerName}/${repoName}/${this.data.task.source_id}`
          })
          console.log('this.sourceFtList',this.sourceFtList)
          this.flag = true

        }
        this.$nextTick(() => {
          this.clipboardHandler && this.clipboardHandler.destroy();
          this.clipboardHandler = initClipboard();
        });
      }
    }
  },
  computed: {
    // 计算字段分组配置
  },
  methods: {
    getCurrentFields(groupName) {
      return this.configs.fields[groupName] || [];
    },
    sparkMD5Hash(str = '') {
      return SparkMD5.hash(str) + Math.random().toString().replace('0.', '');
    },
    renderHljs(str = '') {
      return hljs.highlight('python', str).value;
    },
    renderTitle(key) {
      if (!this.data.is_subscriber && (key === 'endPort' || key === 'port')) {
        return
      }
      const task = this.data.task;
      let result = key;
      switch (key) {
        case 'taskName':
          result = i18n.t('cloudbrainObj.taskName');
          break;
        case 'appName':
          result = i18n.t('cloudbrainObj.appName');
          break;
        case 'spec':
          result = i18n.t('cloudbrainObj.resourceSpec');
          break;
        case 'computerRes':
          result = i18n.t('cloudbrainObj.computeResource');
          break;
        case 'status':
          result = i18n.t('status');
          break;
        case 'imagev1':
          result = i18n.t('cloudbrainObj.image');
          break;
        case 'imagev2':
          result = i18n.t('cloudbrainObj.image');
          break;
        case 'codeObsPath':
          result = i18n.t('cloudbrainObj.codeObsPath');
          break;
        case 'aiCenter':
          result = i18n.t('resourcesManagement.aiCenter');
          break;
        case 'hasInternet':
          result = i18n.t('cloudbrainObj.networkType');
          break;
        case 'creator':
          result = i18n.t('modelManage.creator');
          break;
        case 'repo':
          result = i18n.t('repos.repos');
          break;
        case 'branch':
          result = i18n.t('cloudbrainObj.codeBranch');
          break;
        case 'runVersion':
          result = i18n.t('cloudbrainObj.runVersion');
          break;
        case 'createTime':
          result = i18n.t('cloudbrainObj.createTime');
          break;
        case 'startTime':
          result = i18n.t('cloudbrainObj.startRunTime');
          break;
        case 'endTime':
          result = i18n.t('cloudbrainObj.endRunTime');
          break;
        case 'duration':
          result = i18n.t('cloudbrainObj.runDuration');
          break;
        case 'modelName':
          result = i18n.t('modelManage.modelName');
          break;
        case 'modelVersion':
          result = i18n.t('modelManage.modelVersion');
          break;
        case 'modelFiles':
          result = i18n.t('modelManage.modelFiles');
          break;
        case 'bootFile':
          result = i18n.t('modelManage.bootFile');
          break;
        case 'runParameters':
          result = i18n.t('modelManage.runParameters');
          break;
        case 'workServerNum':
          result = i18n.t('modelManage.workServerNumber');
          break;
        case 'descr':
          result = i18n.t('modelManage.descr');
          break;
        case 'codePath':
          result = i18n.t('cloudbrainObj.codePath');
          break;
        case 'datasetPath':
          result = i18n.t('cloudbrainObj.datasetPath');
          break;
        case 'modelPath':
          result = i18n.t('cloudbrainObj.modelPath');
          break;
        case 'outputPath':
          result = i18n.t('cloudbrainObj.outputPath');
          break;
        case 'bootFile':
          result = i18n.t('modelManage.bootFile');
          break;
        case 'failedReason':
          result = i18n.t('cloudbrainObj.failedReason');
          break;
        case 'timeLimit':
          result = i18n.t('cloudbrainObj.taskIsAutomaticStop');
          break;
        case 'endPort':
          result = task.end_point == '' ? '' : i18n.t('cloudbrainObj.customPath');
          break
        case 'port':
          result = task.port == 0 ? '' : i18n.t('cloudbrainObj.port');
          break
        case 'visualization':
          result = i18n.t('cloudbrainObj.visualization');
          break;
        case 'datasetList':
          result = i18n.t('dataset');
          break;
        case 'modelList':
          result = i18n.t('repos.model');
          break;
        case 'sourceFtName':
          result = i18n.t('cloudbrainObj.sourceFtName');
          break;   
        default:
          break;
      }
      return result;
    },
    renderContent(key) {
      if (!this.data.is_subscriber && (key === 'endPort' || key === 'port')) {
        return
      }
      const task = this.data.task;
      let result = '--';
      switch (key) {
        case 'taskName':
          if(task.display_job_name){
            result = task.display_job_name;
          }
          break;
        case 'appName':
          if(task.app_name){
            result = task.app_name;
          }
          break;
        case 'imagev1':
          if(task.image_name || task.image_url){
            result = `<span class="ui poping up clipboard"
            id="clipboard-${SparkMD5.hash(task.image_url || task.image_name)}"
            data-position="top center"
            data-variation="inverted tiny"
            data-success="${this.$t('copyedLink')}"
            data-content="${this.$t('copy')}" 
            data-original="${this.$t('copy')}"
            data-clipboard-text="${task.image_url || task.image_name}"><span title="${task.image_name || task.image_url}">${ task.image_name || task.image_url }</span></span>`;
          }
          break;
        case 'imagev2':
          if(task.image_name){
            result = `<span class="ui poping up clipboard"
            id="clipboard-${SparkMD5.hash(task.image_name)}"
            data-position="top center"
            data-variation="inverted tiny"
            data-success="${this.$t('copySuccess')}"
            data-content="${this.$t('copy')}" 
            data-original="${this.$t('copy')}"
            data-clipboard-text="${task.image_name}"><span title="${task.image_name}">${task.image_name}</span></span>`;
          }
          break;
        case 'codeObsPath':
          if(task.code_url){
            result = `<span class="ui poping up clipboard"
            id="clipboard-${SparkMD5.hash(task.code_url)}"
            data-position="top center"
            data-variation="inverted tiny"
            data-success="${this.$t('copyedLink')}"
            data-content="${this.$t('copy')}" 
            data-original="${this.$t('copy')}"
            data-clipboard-text="${task.code_url}"><span title="${task.code_url}">${task.code_url}</span></span>`;
          }
          break;
        case 'spec':
          if(task.spec){
            const specObj = renderSpecObject(task.spec, false);
            result = specObj.specStr;
          }
          break;
        case 'computerRes':
          if(task.computeSourceShow || task.compute_source){
            result = task.computeSourceShow || task.compute_source;
          }
          break;
        case 'status':
          if(task.status){
            result = `<span style="display:flex;align-items: center;"><i style="margin-right:4px;" class="${task.status}"></i><span>${task.status}</span>`;
            if(task.status === 'WAITING'){
              if(task.detailed_status === 'dataMigrating'){
                result += `<i style="margin-left:8px;" class="dataMigrating" title="${i18n.t('cloudbrainObj.migratingData')}"></i>`;
              }
              if(task.detailed_status === 'centerPending'){
                result += `<i style="margin-left:8px;" class="centerPending" title="${i18n.t('cloudbrainObj.centerPending')}"></i>`;
              }
              if(task.detailed_status === 'ImagePulling'){
                result += `<i style="margin-left:8px;" class="ImagePulling" title="${i18n.t('cloudbrainObj.ImagePulling')}"></i>`;
              }
            }
            result += `</span>`
          }
          break;
        case 'aiCenter':
          if(task.ai_center){
            result = task.ai_center
          }
          break;
        case 'hasInternet':
          if (task.has_internet == 1) {
            result = i18n.t('cloudbrainObj.noInternet');
          } else if (task.has_internet == 2) {
            result = i18n.t('cloudbrainObj.hasInternet');
          }
          break;
        case 'repo':
          if(task.repo_name && task.repo_owner_name){
            result = `<a style="color:#0066ff;" target="_blank" href="/${task.repo_owner_name}/${task.repo_name}">${task.repo_owner_name} / ${task.repo_alias}</a>`;
          }else if(task.repo_id){
            result = `<span style="color:#888888">-- （${i18n.t('repositoryWasDel')}）</span>`
          }
          break;
        case 'branch':
          if(task.branch_name){
            result = `${task.branch_name}` + (task.commit_id ? `<span class="commit-id">${task.commit_id.slice(0, 10)}</span>` : '');
          }
          
          break;
        case 'creator':
          if(task.creator_name){
            result = task.creator_name;
          }
          break;
        case 'runVersion':
          if(task.current_version_name){
            task.current_version_name;
          }
          break;
        case 'createTime':
          if(task.createTimeStr){
            result = task.createTimeStr;
          }
          break;
        case 'startTime':
          if(task.start_time){
            result = formatDate(new Date(task.start_time * 1000), 'yyyy-MM-dd HH:mm:ss');
          }
          break;
        case 'endTime':
          if(task.end_time){
            result = formatDate(new Date(task.end_time * 1000), 'yyyy-MM-dd HH:mm:ss');
          }
          break;
        case 'duration':
          if(task.formatted_duration){
            result = task.formatted_duration;
          }
          break;
        case 'modelName':
          if(task.pretrain_model_name){
            result = task.pretrain_model_name;
          }
          break;
        case 'modelVersion':
          if(task.pretrain_model_version){
            result = task.pretrain_model_version;
          }
          break;
        case 'modelFiles':
          if(task.pretrain_model_ckpt_name){
            result = `<span title="${task.pretrain_model_ckpt_name}">${task.pretrain_model_ckpt_name}</span>`
          }
          break;
        case 'bootFile':
          if(task.boot_file){
            result = task.boot_file;
          }
          break;
        case 'runParameters':
          const parameters = task.parameters || { parameter: [] };
          const paramList = parameters.parameter || [];
          const paramsStr = paramList.map(item => (`${item.label} = ${item.value}`)).join('; ');
          if(paramsStr){
            result = `<span class="ui poping up clipboard"
            id="clipboard-${SparkMD5.hash(paramsStr)}"
            data-position="top center"
            data-variation="inverted tiny"
            data-success="${this.$t('copySuccess')}"
            data-content="${this.$t('copy')}" 
            data-original="${this.$t('copy')}"
            data-clipboard-text="${paramsStr}"><span title="${paramsStr}">${paramsStr}</span></span>`;
          }
          break;
        case 'workServerNum':
          if(task.work_server_number){
            result = task.work_server_number;
          }
          break;
        case 'descr':
          if(task.description){
            result = `<span title="${escapeHTML(task.description)}">${escapeHTML(task.description)}</span>`;
          }
          break;
        case 'codePath':
          result = task.code_path;
          break;
        case 'datasetPath':
          result = task.dataset_path;
          break;
        case 'modelPath':
          result = task.pretrain_model_path;
          break;
        case 'outputPath':
          result = task.output_path;
          break;
        case 'bootFile':
          if(task.boot_file){
            result = task.boot_file;
          }
          break;
        case 'failedReason':
          if(task.failed_reason){
            result = task.failed_reason;
          }
          break;
        case 'timeLimit':
          if (task.time_limit == 0) {
            result = '';
          } else if (task.time_limit == -1) {
            result = i18n.t('cloudbrainObj.manualStop');
          } else {
            result = i18n.t('cloudbrainObj.automaticStop') + `(${i18n.t('cloudbrainObj.numOfHours', { num: task.time_limit })})`;
          }
          break;
        case 'endPort':
          if(task.end_point){
            result = task.end_point
          }
          break;
        case 'port':
          if(task.port){
            result = task.port
          }
          break;
        case 'visualization':
          result = !!task.visualize_required ? i18n.t('cloudbrainObj.tensorBoardVisualization') : i18n.t('none');
          break;
        case 'datasetList':
          if (task.dataset_list && task.dataset_list.length) {
            result = '<div class="ui list">'
            task.dataset_list.forEach((item) =>{
              if(item.is_delete){
                result += `<div style="line-height:20px" class="item nowrap" title=${item.dataset_alias}>${item.dataset_alias} (${i18n.t('cloudbrainObj.fileWasDeleted')})</div>`
              }else{
                result += `<a style="line-height:20px;color:#0066ff;" target="_blank" href="/datasets/detail/${item.owner_name}/${item.dataset_name}" class="item nowrap" title=${item.owner_name}/${item.dataset_alias}>${item.owner_name}/${item.dataset_alias}</a>`
              }
            })
            result += '</div>'
          }
          if(task.job_type === 'EVAL'){
            const datasetParam = task.parameters.parameter.find(item => item.label === 'datasets');
            let selectDatasetTemp = []
            if (datasetParam && datasetParam.value) {
              selectDatasetTemp = datasetParam.value.split(' ')
                .sort((a, b) => a.localeCompare(b))
                .map(item => ({ k: item, v: item }));
              if( selectDatasetTemp.length>0 )
              result = '<div class="ui list">'
              selectDatasetTemp.forEach(item => {
                result += `<a style="line-height:20px;color:#0066ff;" target="_blank" href="/datasets/detail/Open_Dataset/${item.k}" class="item nowrap" title=${item.k}>${item.k}</a>`
              })
            }
          }
          break;
        case 'modelList':
          if (task.pretrain_model_list && task.pretrain_model_list.length) {
            result = '<div class="ui list">'
            task.pretrain_model_list.forEach((item) =>{
              if(item.is_delete){
                result += `<div style="line-height:20px" class="item nowrap" title=${item.alias}>${item.alias} (${i18n.t('cloudbrainObj.fileWasDeleted')})</div>`
              }else{
                result += `<a style="line-height:20px;color:#0066ff;" target="_blank" href="/models/detail/${item.owner_name}/${item.name}" class="item nowrap" title=${item.owner_name}/${item.alias}>${item.owner_name}/${item.alias}</a>`
              }
            })
            result += '</div>'
          }
          break;
        case 'sourceFtName':
          if (task.source_id && task.app_name) {
            result = '<div class="ui list">'
            result += `<a style="line-height:20px;color:#0066ff;" target="_blank" href="/modelbase/nlp/sft/detail/${task.source_id}" class="item nowrap" title=${task.app_name}>${task.app_name}</a>`
            result += '</div>'
          }
          break;
        default:
          break;
      }
      return result;
    },
    refresh() { }
  },
  beforeMount() { },
  mounted() {
  }
};
</script>

<style scoped lang="less">
@media only screen and (max-width: 767px) {
  .item-container{
    padding: 20px 10px 10px 10px !important;
    .layout-wrapper{
      flex-direction: column;
      .left-column, .right-column{
        width: 100% !important;
      }
      
      .field-title{
        padding-left: 0 !important;
      }
      .field-grid{
        margin-left: 0 !important;
      }
    }
    
  }
  .item-block {
    min-width: 100% !important;
    padding-left: 0 !important;
  }

  .content {
    white-space: normal !important;
  }

  .title {
    text-align: right;
    width: 105px !important;
  }
}

.item-container {
  width: 100%;
  height: 100%;
  padding: 30px 50px 24px 20px;
  overflow: auto;
  .layout-wrapper {
    display: flex;
    gap: 40px; /* 左右栏间距 */
    align-items: flex-start;
    .left-column, .right-column {
      flex: 1 1 50%; /* 各占50% */
      min-width: 0; /* 重要：防止内容撑开 */
      display: flex;
      flex-direction: column;
      gap: 24px; /* 左侧三个字段组之间的间距 */
      .field-title{
        font-weight: bold;
        color: rgb(118, 162, 211);
        font-size: 16px;
        line-height: 22px;
        padding-left: 30px;
      }
      .field-divider{
        height: 1px;
        background-color: rgba(235,238,245,1);
        margin-left: 30px;
        margin-top: 10px;
      }
      .field-grid{
        margin-left: 40px;
        margin-top: 14px;
        .field-item{
          color: rgba(16,16,16,0.5);
          margin-bottom: 10px;
          display: flex;
          label{
            min-width: 130px;
            text-align: right;
          }
          span{
            margin-left: 10px;
            color: #101010
          }
        }
        .code-content-c {
          position: relative;

          .code-content {
            position: relative;
            border-radius: 6px;
            background-color: rgba(240,240,240,1);
            padding: 0 0.8em;
            overflow: auto;
          }

          .copy-btn {
            position: absolute;
            top: 0.5em;
            right: 0.8em;
          }
        }
      }
    }
  }   
  .item-block {
    display: flex;
    width: 50%;
    padding-bottom: 20px;
    padding-right: 20px;
    padding-left: 10px;
    min-width: 200px;

    .title {
      width: 118px;
      padding-right: 20px;
      color: #8a8e99;
      font-size: 12px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .content {
      flex: 1;
      width: 0;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-size: 12px;
      color: rgba(0, 0, 0, .87);

      &.wrap {
        white-space: normal;
        word-break: break-all;
      }

      
    }

    &.item-dataset,
    &.item-model-list,
    &.item-model-code {
      display: block;
      width: 100%;
      padding-right: 0;

      .table-container {
        /deep/ .el-table__header {
          th {
            background: rgb(245, 245, 246);
            color: #8a8e99;
            font-size: 12px;
            font-weight: 700;
          }
        }

        /deep/ .el-table__body {
          td {
            color: rgb(16, 16, 16);
            font-weight: 400;
            font-size: 12px;
          }
        }

        .download-link-c {
          word-wrap: break-word;
          word-break: break-all;
          overflow: hidden;
          text-overflow: ellipsis;
          display: -webkit-box;
          -webkit-box-orient: vertical;
          -webkit-line-clamp: 2;
          max-height: 50px;
        }
      }
    }

    &.item-failed-reason {
      width: 100%;

      .content {
        overflow: auto;
        text-overflow: clip;
        white-space: break-spaces;
      }
    }

    .code-c {
      margin-top: -2px;

      .code-title {
        color: #8a8e99;
        font-size: 12px;
        font-weight: 700;
        margin-bottom: 4px;
      }

      
    }
  }

  /deep/.clipboard {
    cursor: pointer;
  }
  /deep/.commit-id {
    margin-left: 8px;
    color: rgba(16,16,16,0.8);;
    background-color: #e8e8e8;
    font-weight: 700;
    border-radius: 6px;
    padding: 0.3em 0.5em;
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
</style>
