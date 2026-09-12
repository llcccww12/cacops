<template>
  <div>
    <div v-if="emptyPage" style="padding-top:50px">
      <NotFound></NotFound>
    </div>
    <div v-else>
        <div class="ui container">
          <div class="content" v-if="!showUploadPage">
            <div class="top-wrap">
              <!-- 文件导航路径 -->
              <div class="left-path-wrap">
                <div class="file-path nowrap" @copy="handleCopy" v-for="(item, index) in filePath" :key="index">
                  <span v-if="index == filePath.length - 1">{{ item }}</span>
                  <a v-if="index != filePath.length - 1" class="canback" @click="goBackDir(index)">{{item}}</a>
                  <span class="divider"> / </span>
                </div>
              </div>
              <!-- 顶部操作区域 -->
              <div class="right-btn-c">
                <a v-if="dataObj.can_download" @click="downloadAllModel()" class="download-btn">
                  {{ type==='dataset' ? $t('datasetObj.dataDownloadAll') : $t('modelManage.modelDownloadAll')}}
                </a>
                <el-button v-if="dataObj.can_edit_file" style="background-color: #0066ff;" type="primary" @click="goUploadPage">
                  {{type==='dataset' ?  $t('uploadDatasetFile') : $t('modelManage.uploadModelFiles')}}
                </el-button>
              </div>
            </div>
            <div class="table-container">
              <el-table ref="tableRef" 
                :data="filesList" 
                :header-cell-style="headerStyle" 
                style="width: 100%" 
                v-loading="loading" 
                @sort-change="sortChange"
                :cell-style="cellStyle"
                max-height="800px"
              >
                  <el-table-column column-key="FileName" prop="FileName" sortable="custom"
                    :label="$t('modelManage.fileName')" align="left" header-align="center" min-width="260">
                    <template slot-scope="scope">
                      <div class="tbl-file-name">
                        <div v-if="scope.row.IsDir" @click="goNextDir(scope.row)" class="fitted folder">
                          <i class="ri-folder-5-fill" style="color: #ffb02c;font-size:16px;"></i>
                          <span class="nowrap" :title="scope.row.FileName">{{ scope.row.FileName }}</span>
                        </div>
                        <div v-else class="fitted">
                          <i class="ri-file-text-fill" style="color: #8ca2aa;font-size:16px"></i>
                          <span class="nowrap" :title="scope.row.FileName">{{ scope.row.FileName }}</span>
                          <i class="ri-file-copy-2-line ui poping up clipboard-model-name" data-position="top center"
                            data-variation="inverted tiny" :data-success="$t('copySuccess')" :data-content="$t('copy')"
                            :data-original="$t('copy')" :data-clipboard-text="scope.row.FileName">
                          </i>
                        </div>
                      </div>
                    </template>
                  </el-table-column>
                  <el-table-column column-key="SizeShow" prop="Size" sortable="custom" :label="$t('modelManage.fileSize')" 
                    align="center" header-align="center" width="200">
                    <template slot-scope="scope">
                        {{ scope.row.IsDir ? '--' : formatFileSize(scope.row.Size) }}
                    </template>
                  </el-table-column>
                  <el-table-column column-key="ModTime" prop="ModTime" sortable="custom" :label="$t('modelManage.updateTime')" 
                    align="center" header-align="center" width="200">
                  </el-table-column>
                  <el-table-column column-key="operate" prop="operate" fixed="right" :label="$t('modelManage.operate')" align="center" header-align="center" width="200">
                    <template slot-scope="scope">
                      <span v-if="scope.row.IsSupportPreview && type=='dataset'" 
                        class="btn-del" 
                        @click="previewFile(scope.row)"
                      >
                        {{$t('modelManage.preview')}}
                      </span>
                      <span v-if="!scope.row.IsDir && dataObj.can_download" 
                        class="btn-del" 
                        @click="downLoadFile(scope.row)"
                      >
                        {{ $t('modelManage.download') }}
                      </span>
                      <span v-if="!scope.row.IsDir && dataObj.can_edit_file " 
                        class="btn-del" 
                        @click="deleteFile(scope.row)"
                      >
                        {{ $t('modelManage.delete') }}
                      </span>
                    </template>
                  </el-table-column>
                </el-table>
                <el-button v-if="has_next" style="margin-top:24px" @click="getDatasetFileList">{{$t('datasets.loadMore')}}</el-button>
            </div>
          </div>
          <CreateForm v-else :showSteps="false" :title="[type==='dataset'? $t('uploadDatasetFile') : $t('modelManage.uploadModelFiles')]" class="upload-wrap">
            <FileUpload :modelId="dataObj.id" :subjectType="type ==='dataset' ? '1' : '2'" :ownerName="dataObj.owner_name" @cancel="cancel" @uploadFinish="uploadFinish" :uploadDir="params.parent_dir"></FileUpload>
          </CreateForm>
        </div>
    </div>
    <BaseDialog :visible="visible" :title="title" @closed="closeDialog">
      <div class="body-c">
        <div class="not-support" v-if="showFlag">
          <svg width="168" height="168" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" style="enable-background:new 0 0 200 200" viewBox="0 0 200 200"><path d="M64.2 137.3H40s1.3-1.6 1.3-10.7c0-10.2-5.8-11.1-5.8-24.5 0-10.5 4-15.1 4-24.3 0-8-3.2-15.9-3.2-15.9h27.9s-4.3 30.6-4.3 39.7c0 9 4.3 35.7 4.3 35.7zm110.7 0-19.4.6s-4.2.4-4.2-8.7c0-10.2 5.6-13.7 5.6-27 0-10.5-5.9-16.3-5.9-25.5 0-8 2-14.7 2-14.7h21.8s4.4 30.6 4.4 39.7-4.3 35.6-4.3 35.6z" style="fill:#66c8ff"/><path d="M185.4 6.8H14.6c-6.8 0-12.3 5.5-12.3 12.3v117.5c0 6.8 5.5 12.3 12.3 12.3h9.3c.6 1.7 1.6 3.5 3.1 5 .4.4.8.7 1.4.8 15.5 4 19.3 9 20.3 11.6 1 2.9-.7 5.2-.8 5.4-1 1.3-.8 3.1.5 4.1 1.3 1 3.1.8 4.2-.5 1.4-1.8 3.4-6 1.8-10.8-2.2-6.5-10.1-11.6-23.5-15.2a5.03 5.03 0 0 1-1.4-3.3c0-1.2.7-2.4 2.1-3.4 0 0 .1 0 .1-.1.4-.3.8-.6 1.3-.8.1-.1.2-.1.4-.2.6-.3 1.2-.5 1.9-.7 1-.2 2.9-.6 5.9-.6 5.2 0 11.2 1 18 3 1.1.5 2.3 1 3.4 1.6.4.2.7.4 1 .6 1.4.8 2.9 1.7 4.3 2.6.1.1.2.1.4.2 22.6 15.9 20.1 24 20.1 24-.8 1.4-.2 3.2 1.2 4 .4.2.9.4 1.4.4 1.1 0 2.1-.6 2.6-1.6 1-1.8 4-10.6-14.7-26H131c-18.8 15.5-15.7 24.3-14.8 26.1.6 1 1.6 1.6 2.7 1.6.5 0 .9-.1 1.3-.4 1.4-.8 2-2.5 1.3-3.9 0-.1-2.7-8.1 20.1-24.1 1.6-1.1 3.2-2 4.8-3 .3-.2.6-.3.8-.5 1.2-.6 2.4-1.2 3.5-1.7 6.7-2 12.7-3 17.9-3 3 0 4.8.4 5.8.6.6.2 1.3.5 1.9.7.1 0 .2.1.3.1.5.2 1 .6 1.4.8 1.4 1 2.1 2.2 2.1 3.4 0 .9-.4 2-1.4 3.3-13.4 3.6-21.3 8.7-23.5 15.3-1.6 4.8.4 9.1 1.8 10.8a3 3 0 0 0 4.2.5 3 3 0 0 0 .5-4.2s-1.8-2.4-.8-5.3c.9-2.6 4.7-7.7 20.3-11.7.5-.1 1-.4 1.4-.8 1.6-1.6 2.6-3.4 3.1-5.1 6.7-.1 12.1-5.6 12.1-12.3V19.1a12.5 12.5 0 0 0-12.4-12.3zM33.7 63.7h26.4c1.1 3.9 3 13.9-1.4 26.4l-.4.9c-2.2 6-4.8 13.5 1.1 29 3.6 9.4 1.4 14.8-.1 17.1a46.37 46.37 0 0 0-11-2.9l-.8-.1c-1.5-.2-3-.3-4.4-.3h-1c-1.2 0-2.3.1-3.4.2-.3 0-.6 0-.8.1-.2 0-.5.1-.7.1 1-4.9 1.7-12.2-1.5-17-3.5-5.2-4.3-21.3-.1-33.4 1.9-5.3-.3-14.7-1.9-20.1zm116.4 56.2c6-15.5 3.3-23 1.1-29l-.3-.8c-4.4-12.6-2.5-22.5-1.4-26.4h26.4c-1.6 5.4-3.7 14.9-1.9 20.1 4.2 12.2 3.4 28.3 0 33.4-3.2 4.9-2.5 12.1-1.5 17-.2 0-.5-.1-.7-.1-.2 0-.4 0-.6-.1-1.2-.2-2.5-.2-3.8-.3h-.5c-4.9 0-10.5 1-16.4 3.3-1.8-2.2-4-7.6-.4-17.1zm41.7 16.8c0 3.5-2.8 6.3-6.2 6.4-.4-1.1-1-2.2-1.9-3.2l-.4-.4-.1-.1c-.4-.5-1-.9-1.4-1.3-.2-.1-.4-.2-.5-.4-.4-.3-.9-.6-1.4-.8-.2-.1-.4-.2-.6-.4-.1 0-.1 0-.2-.1-1.1-3.8-2.8-12-.2-15.8 4.9-7.4 5.3-25.5.7-38.7-1.2-3.6.7-12.3 2.5-18.2h1.2a3 3 0 0 0 3-3 3 3 0 0 0-3-3h-38.4a3 3 0 0 0-3 3c0 1.2.7 2.1 1.6 2.6-1.3 5.3-2.8 15.7 1.9 28.8l.3.9c1.9 5.4 4.1 11.4-1.1 24.8-4.2 10.9-2 18 .2 21.7-2 1.1-4 2.2-6 3.6H70.7c-.2-.1-.4-.2-.6-.4-1.6-1.1-3.2-2-4.8-3-.2-.1-.4-.2-.6-.4 2.2-3.8 4.4-10.9.2-21.7-5.2-13.4-3-19.4-1.1-24.8l.4-.9c4.6-13.1 3.2-23.5 1.9-28.8 1-.5 1.6-1.4 1.6-2.6a3 3 0 0 0-3-3H26.4a3 3 0 0 0-3 3 3 3 0 0 0 3 3h1.2c1.8 5.9 3.7 14.6 2.5 18.2-4.4 12.9-4.1 31.4.7 38.7 2.6 3.8.9 12-.2 15.8-.1 0-.2.1-.2.1-.2.1-.4.2-.6.4-.5.2-.8.5-1.3.8l-.6.4c-.5.4-1 .8-1.4 1.2l-.2.2-.4.4c-.4.4-.6.8-1 1.3-.4.6-.7 1.3-1 1.9h-9.3c-3.6 0-6.4-2.9-6.4-6.4V19.2c0-3.6 2.9-6.4 6.4-6.4h170.8c3.6 0 6.4 2.9 6.4 6.4v117.5zM93.7 49.9a3 3 0 0 0-3-3c-10.1 0-18.6-6.2-22.6-9.2-1-.7-1.7-1.2-2.2-1.5-1.4-.8-3.2-.4-4.1 1-.8 1.4-.4 3.2 1 4.1.4.2 1 .7 1.7 1.2 4.6 3.4 14.1 10.3 26.1 10.3 1.8 0 3.1-1.3 3.1-2.9zm42.5-13.7c-.5.3-1.2.8-2.2 1.5-4.1 3-12.6 9.2-22.6 9.2a3 3 0 0 0-3 3 3 3 0 0 0 3 3c12 0 21.6-7 26.1-10.3.7-.6 1.3-1 1.7-1.2 1.4-.8 1.9-2.6 1-4.1-.8-1.5-2.6-1.9-4-1.1zm-15 80.5-19.9-9c-.8-.4-1.8-.3-2.5 0l-17.6 9c-1.4.7-2 2.5-1.3 4 .7 1.4 2.5 2 4 1.3l16.4-8.3 18.6 8.4c.4.2.8.2 1.2.2 1.1 0 2.2-.6 2.7-1.8.5-1.4-.2-3.1-1.6-3.8z" style="fill:#78c6ff"/><path d="M185.4 6.8H14.6c-6.8 0-12.3 5.5-12.3 12.3v117.5c0 6.8 5.5 12.3 12.3 12.3h9.3c.6 1.7 1.6 3.5 3.1 5 .4.4.8.7 1.4.8 15.5 4 19.3 9 20.3 11.6 1 2.9-.7 5.2-.8 5.4-1 1.3-.8 3.1.5 4.1 1.3 1 3.1.8 4.2-.5 1.4-1.8 3.4-6 1.8-10.8-2.2-6.5-10.1-11.6-23.5-15.2a5.03 5.03 0 0 1-1.4-3.3c0-1.2.7-2.4 2.1-3.4 0 0 .1 0 .1-.1.4-.3.8-.6 1.3-.8.1-.1.2-.1.4-.2.6-.3 1.2-.5 1.9-.7 1-.2 2.9-.6 5.9-.6 5.2 0 11.2 1 18 3 1.1.5 2.3 1 3.4 1.6.4.2.7.4 1 .6 1.4.8 2.9 1.7 4.3 2.6.1.1.2.1.4.2 22.6 15.9 20.1 24 20.1 24-.8 1.4-.2 3.2 1.2 4 .4.2.9.4 1.4.4 1.1 0 2.1-.6 2.6-1.6 1-1.8 4-10.6-14.7-26H131c-18.8 15.5-15.7 24.3-14.8 26.1.6 1 1.6 1.6 2.7 1.6.5 0 .9-.1 1.3-.4 1.4-.8 2-2.5 1.3-3.9 0-.1-2.7-8.1 20.1-24.1 1.6-1.1 3.2-2 4.8-3 .3-.2.6-.3.8-.5 1.2-.6 2.4-1.2 3.5-1.7 6.7-2 12.7-3 17.9-3 3 0 4.8.4 5.8.6.6.2 1.3.5 1.9.7.1 0 .2.1.3.1.5.2 1 .6 1.4.8 1.4 1 2.1 2.2 2.1 3.4 0 .9-.4 2-1.4 3.3-13.4 3.6-21.3 8.7-23.5 15.3-1.6 4.8.4 9.1 1.8 10.8a3 3 0 0 0 4.2.5 3 3 0 0 0 .5-4.2s-1.8-2.4-.8-5.3c.9-2.6 4.7-7.7 20.3-11.7.5-.1 1-.4 1.4-.8 1.6-1.6 2.6-3.4 3.1-5.1 6.7-.1 12.1-5.6 12.1-12.3V19.1a12.5 12.5 0 0 0-12.4-12.3zM33.7 63.7h26.4c1.1 3.9 3 13.9-1.4 26.4l-.4.9c-2.2 6-4.8 13.5 1.1 29 3.6 9.4 1.4 14.8-.1 17.1a46.37 46.37 0 0 0-11-2.9l-.8-.1c-1.5-.2-3-.3-4.4-.3h-1c-1.2 0-2.3.1-3.4.2-.3 0-.6 0-.8.1-.2 0-.5.1-.7.1 1-4.9 1.7-12.2-1.5-17-3.5-5.2-4.3-21.3-.1-33.4 1.9-5.3-.3-14.7-1.9-20.1zm116.4 56.2c6-15.5 3.3-23 1.1-29l-.3-.8c-4.4-12.6-2.5-22.5-1.4-26.4h26.4c-1.6 5.4-3.7 14.9-1.9 20.1 4.2 12.2 3.4 28.3 0 33.4-3.2 4.9-2.5 12.1-1.5 17-.2 0-.5-.1-.7-.1-.2 0-.4 0-.6-.1-1.2-.2-2.5-.2-3.8-.3h-.5c-4.9 0-10.5 1-16.4 3.3-1.8-2.2-4-7.6-.4-17.1zm41.7 16.8c0 3.5-2.8 6.3-6.2 6.4-.4-1.1-1-2.2-1.9-3.2l-.4-.4-.1-.1c-.4-.5-1-.9-1.4-1.3-.2-.1-.4-.2-.5-.4-.4-.3-.9-.6-1.4-.8-.2-.1-.4-.2-.6-.4-.1 0-.1 0-.2-.1-1.1-3.8-2.8-12-.2-15.8 4.9-7.4 5.3-25.5.7-38.7-1.2-3.6.7-12.3 2.5-18.2h1.2a3 3 0 0 0 3-3 3 3 0 0 0-3-3h-38.4a3 3 0 0 0-3 3c0 1.2.7 2.1 1.6 2.6-1.3 5.3-2.8 15.7 1.9 28.8l.3.9c1.9 5.4 4.1 11.4-1.1 24.8-4.2 10.9-2 18 .2 21.7-2 1.1-4 2.2-6 3.6H70.7c-.2-.1-.4-.2-.6-.4-1.6-1.1-3.2-2-4.8-3-.2-.1-.4-.2-.6-.4 2.2-3.8 4.4-10.9.2-21.7-5.2-13.4-3-19.4-1.1-24.8l.4-.9c4.6-13.1 3.2-23.5 1.9-28.8 1-.5 1.6-1.4 1.6-2.6a3 3 0 0 0-3-3H26.4a3 3 0 0 0-3 3 3 3 0 0 0 3 3h1.2c1.8 5.9 3.7 14.6 2.5 18.2-4.4 12.9-4.1 31.4.7 38.7 2.6 3.8.9 12-.2 15.8-.1 0-.2.1-.2.1-.2.1-.4.2-.6.4-.5.2-.8.5-1.3.8l-.6.4c-.5.4-1 .8-1.4 1.2l-.2.2-.4.4c-.4.4-.6.8-1 1.3-.4.6-.7 1.3-1 1.9h-9.3c-3.6 0-6.4-2.9-6.4-6.4V19.2c0-3.6 2.9-6.4 6.4-6.4h170.8c3.6 0 6.4 2.9 6.4 6.4v117.5zM93.7 49.9a3 3 0 0 0-3-3c-10.1 0-18.6-6.2-22.6-9.2-1-.7-1.7-1.2-2.2-1.5-1.4-.8-3.2-.4-4.1 1-.8 1.4-.4 3.2 1 4.1.4.2 1 .7 1.7 1.2 4.6 3.4 14.1 10.3 26.1 10.3 1.8 0 3.1-1.3 3.1-2.9zm42.5-13.7c-.5.3-1.2.8-2.2 1.5-4.1 3-12.6 9.2-22.6 9.2a3 3 0 0 0-3 3 3 3 0 0 0 3 3c12 0 21.6-7 26.1-10.3.7-.6 1.3-1 1.7-1.2 1.4-.8 1.9-2.6 1-4.1-.8-1.5-2.6-1.9-4-1.1zm-15 80.5-19.9-9c-.8-.4-1.8-.3-2.5 0l-17.6 9c-1.4.7-2 2.5-1.3 4 .7 1.4 2.5 2 4 1.3l16.4-8.3 18.6 8.4c.4.2.8.2 1.2.2 1.1 0 2.2-.6 2.7-1.8.5-1.4-.2-3.1-1.6-3.8z" style="fill:#444"/></svg>
          <div>{{ showMessage }}</div>
        </div>
        <template v-else>
          <div class="img-c" v-if="fileImg">
            <img :src="fileImg" alt="">
          </div>
          
          <div class="file-content" v-if="textContent">
            <pre>{{textContent}}</pre>
          </div>
        </template>
      </div>
    </BaseDialog>
    <CommonSDK ref="childModelTips" :data="sdkData" :title="$t('modelObj.codeDownDlgTitle')" :closeText="$t('cloudbrainObj.dialogTips.tips8')"></CommonSDK>
  </div>
</template>

<script>
import FileUpload from '~/pages/guide/components/FileUpload.vue';
import CommonSDK from '~/components/square/CommonSDK.vue';
import BaseDialog from '~/components/BaseDialog.vue';
import CreateForm from '~/pages/guide/components/CreateForm.vue';
import { getFileList, delSingleDataset, getFilePreview } from "~/apis/modules/dataset";
import {  transFileSize, initClipboard } from '~/utils';
import { lang } from '~/langs';
const maxModelFilesSize = window.MAX_MODEL_SIZE ||  1024*1024*1024*0.5;

export default {
  name: 'AiforgeFileList1',
  props: {
    dataObj: { type: Object, default: () => ({}) },
    type: { type: String, default: 'dataset' },
  },
  data() {
    return {
      emptyPage: false,
      filesList: [],
      loading: false,
      params: {
        parent_dir: '',
        page_size: 10,
        marker: ''
      },
      has_next: false,
      filePath: [this.dataObj.name],
      showUploadPage:false,

      visible: false,
      title: '',
      previewCache: new Map(), // 使用Map来缓存预览数据
      fileImg: '',
      textContent: '',
      supportImgReg: /(\.jpg|\.jpeg|\.png|\.gif|\.bmp|\.webp)$/i,
      supportTxtReg: /(\.txt|\.xml|\.html|\.json|\.py|\.sh|\.md|\.csv|\.log|\.js|\.css|\.ipynb)$/i,


      sdkData: {},
      codeUsePromotePath: `tips/modelDown/sdkcode${lang == 'zh-CN' ? '' : '_en'}.md`,
      showFlag: false,
      showMessage: ''
    };
  },
  components: { FileUpload, CreateForm, CommonSDK, BaseDialog },
  mounted() {
    if(this.type==='dataset'){
      this.params.dataset_id = this.dataObj.id;
    }else{
      this.params.aimodel_id = this.dataObj.id;
    }
    this.getDatasetFileList()
  },
  methods: {
    handleCopy(e) {
      e.preventDefault();
  
      // 构造要复制的文本
      const textToCopy = this.filePath.join('/');
      
      // 将文本放入剪贴板
      e.clipboardData.setData('text/plain', textToCopy);
    },
    cancel(){
      this.showUploadPage = false;
    },
    uploadFinish(){
      this.showUploadPage = false;
      this.getDatasetFileList();
      this.$emit('editSuccess')
    },
    sortChange(options) {
      let { prop, order } = options;
      console.log(prop)
      if (!order) {
        prop = 'FileName';
        order = 'ascending'
      }
      this.filesList.sort((a, b) => {
        if (a.IsDir != b.IsDir && order == 'ascending') {
          return a.IsDir ? -1 : 1;
        }
        if (a.IsDir != b.IsDir && order == 'descending') {
          return a.IsDir ? 1 : -1;
        }
        if (prop == 'FileName' && order == 'ascending') {
          return a.FileName.toLocaleLowerCase().localeCompare(b.FileName.toLocaleLowerCase());
        }
        if (prop == 'FileName' && order == 'descending') {
          return b.FileName.toLocaleLowerCase().localeCompare(a.FileName.toLocaleLowerCase());
        }
        if (prop == 'Size' && order == 'ascending') {
          return a.Size - b.Size;
        }
        if (prop == 'Size' && order == 'descending') {
          return b.Size - a.Size;
        }
        if (prop == 'ModTime' && order == 'ascending') {
          return a.ModTime.localeCompare(b.ModTime)
        }
        if (prop == 'ModTime' && order == 'descending') {
          return b.ModTime.localeCompare(a.ModTime)
        }
        return 0;
      });
    },
    goNextDir(row) {
      this.params.parent_dir = `${row.ParenDir}/${row.FileName}`;
      this.filePath.push(row.FileName);
      this.params.marker = ''; // 重置分页标记
      this.getDatasetFileList();
    },
    goBackDir(index) {
      if (index === this.filePath.length - 1) return;
      // 计算要返回的路径
      const pathParts = this.filePath.slice(0, index + 1);
      this.filePath = pathParts;
      // 如果是根目录
      if (index === 0) {
        this.params.parent_dir = '';
      } else {
        // 找到对应的目录对象
        const result = '/' + pathParts.slice(1).join('/');
        this.params.parent_dir = result;
        
      }
      
      this.params.marker = ''; // 重置分页标记
      this.getDatasetFileList();
    },
    async previewFile(file){
      this.visible = true
      this.title = file.FileName
      const cacheKey = `${file.ParenDir}/${file.FileName}`;
      console.log("cacheKey",cacheKey)
      console.log("previewCache",this.previewCache)
      // 检查缓存中是否已有数据
      if (this.previewCache.has(cacheKey)) {
        const cachedData = this.previewCache.get(cacheKey);
        if (cachedData.type === 'image') {
          this.fileImg = cachedData.content;
          this.textContent = '';
        } else {
          this.fileImg = '';
          this.textContent = cachedData.content;
        }
        return;
      }
      try {
        const params = {
          [this.type === 'dataset' ? 'dataset_id' : 'aimodel_id']: this.dataObj.id,
          parent_dir: file.ParenDir,
          file_name: file.FileName
        }
        console.log("this.supportImgReg.test(file.FileName)",this.supportImgReg.test(file.FileName))
        if (this.supportImgReg.test(file.FileName)){
          this.textContent = ''
          const queryParams = new URLSearchParams(params)
          this.fileImg = `/api/v1/dataset/preview?${queryParams.toString()}`
          // 缓存图片URL
          this.previewCache.set(cacheKey, {
            type: 'image',
            content: this.fileImg
          });
        }else if(this.supportTxtReg.test(file.FileName)){
          this.fileImg = ''
          const response = await getFilePreview(params,this.type);
          console.log(response)
          if(response?.data && response?.data.code === 4004 ){
            this.showMessage = response.data.msg
            this.showFlag = true
            this.textContent = ''
          }
          this.textContent = response.data
          // 缓存文本内容
          this.previewCache.set(cacheKey, {
            type: 'text',
            content: this.textContent
          });
        }
      } catch (error) {
        this.$message.error(error)
      }
    },
    closeDialog() {
      this.visible = false;
      this.showFlag = false
    },
    downloadAllModel(){
      if (!this.filesList.length) return;
      if(this.dataObj.size <= maxModelFilesSize ){
        let downloadElement = document.createElement('a')
        let href = `/api/v1/${this.type}/download?${this.type}_id=${this.dataObj.id}`
        downloadElement.href = href
        document.body.appendChild(downloadElement)
        downloadElement.click() //点击下载
        document.body.removeChild(downloadElement) //下载完成移除元素
      }else{
        if(this.type === 'dataset'){
          this.sdkData.dataset_id = this.dataObj.id
        }else{
          this.sdkData.aimodel_id = this.dataObj.id
        }
        this.sdkData.type = this.type
        this.sdkData.file_name = ''
        this.$nextTick(() => {
          this.$refs.childModelTips.handlerOpen('', maxModelFilesSize)
        })
      }
    },
    downLoadFile(file){
      if(file.Size <= maxModelFilesSize ){
        const params = `${this.type}_id=${this.dataObj.id}&parent_dir=${file.ParenDir}&file_name=${file.FileName}`
        let downloadElement = document.createElement('a')
        let href = `/api/v1/${this.type}/file?${params}`;
        downloadElement.href = href
        document.body.appendChild(downloadElement)
        downloadElement.click() //点击下载
        document.body.removeChild(downloadElement) //下载完成移除元素
      }else{
        if(this.type === 'dataset'){
          this.sdkData.dataset_id = this.dataObj.id
        }else{
          this.sdkData.aimodel_id = this.dataObj.id
        }
        this.sdkData.type = this.type
        this.sdkData.file_name = ''
        this.sdkData.file_name = file.FileName
        this.$nextTick(() => {
          this.$refs.childModelTips.handlerOpen('', maxModelFilesSize)
        })
      }
    },
    deleteFile(file){
      this.$confirm(this.$t('datasetObj.deleteDataFileConfirmTips',{name: file.FileName}), this.$t('tips'), {
        confirmButtonText: this.$t('confirm1'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        this.loading = true;
        const params = {
          [this.type === 'dataset' ? 'dataset_id' : 'aimodel_id']: this.dataObj.id,
          parent_dir: file.ParenDir,
          file_name: file.FileName,
        }
        delSingleDataset(params,this.type).then(response => {
          const res = response.data;
          if (res.code ===0) {
            this.params.marker = ''
            this.getDatasetFileList();
            this.$emit('editSuccess')
          } else {
            this.loading = false;
            this.$message({
              type: 'error',
              message: res.msg || this.$t('datasetObj.dataFileDeleteFailed'),
            });
          }
        }).catch(err => {
          this.loading = false;
          console.log(err);
          this.$message({
            type: 'error',
            message: err || this.$t('datasetObj.dataFileDeleteFailed'),
          });
        });
      }).catch(() => { });
    },
    goUploadPage(){
      this.showUploadPage = true;
    },
    async getDatasetFileList(){
        this.loading = true;
        console.log(this.params)
        try {
            const response = await getFileList(this.params,this.type)
            console.log(response)
            const res = response.data
            if(res.code === 0){
                if(this.params.marker){
                    console.log("marker",this.params.marker)
                    console.log(res.data.file_list)
                    
                    this.filesList = [...this.filesList, ...res.data.file_list];
                    console.log(this.filesList)
                }else{
                    this.filesList = res.data.file_list
                }
                this.has_next = res.data.has_next
                this.params.marker = res.data.marker
                this.$nextTick(() => {
                  initClipboard('.tbl-file-name .clipboard-model-name');
                });
            }else{
                this.$message.error(res.msg || '加载文件列表出错')
            }
        } catch (error) {
            console.error('加载文件列表出错:', error);
            this.$message.error(error);
        } finally {
            this.loading = false;
        }
    },
    formatFileSize(size){
        return transFileSize(size)
    },
    headerStyle() {
      return {
        backgroundColor: 'rgba(247,247,247,1)',
        color: 'rgba(16, 16, 16)',
        fontSize: '14px',
        fontWeight: '400'
      }
    },
    cellStyle() {
      return {
        fontSize: '14px',
        color: 'rgba(16, 16, 16)',
        fontWeight: '400'
      }
    }
  },
};
</script>

<style lang="less" scoped>
.content{
    margin-top: 28px;
    .top-wrap{
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 22px;
      flex-wrap: wrap;
      width: 100%;
      .left-path-wrap{
          display: flex;
          flex-wrap: wrap;
          width: 100%;
          .file-path{
              margin-right: 6px;
              font-size: 14px;
              font-weight: 550;
              color: #101010;
              .canback{
                  color: rgb(50, 145, 248);
              }
              .divider{
                  color:rgba(0,0,0,.4);
              }
              
          }
          
      }
      .right-btn-c{
          // min-width: 300px;
          // display: flex;
          // align-items: center;
          // justify-content: flex-end;
          margin-left: auto;
          .download-btn{
              // text-align: right;
            color: rgba(0, 102, 255, 1);
            margin-right: 16px;
          }
      }
    }
    .table-container{
      .tbl-file-name{
        .fitted{
          display: flex;
          align-items: center;
          span{
            margin: 0 6px;
          }
          .clipboard-model-name{
            color: #919191;
            display: none;
            cursor: pointer;
            
          }
          &:hover {
            .clipboard-model-name {
              display: inline-block;
            }
          }
        }
        .folder{
          cursor: pointer;
        }
      }
      .btn-del{
        color: rgb(0, 102, 255);
        cursor: pointer;
        margin: 0 5px;
      }
    }
}
.upload-wrap{
  margin-top: 32px;
}
.body-c{
  // max-height: 600px;
  // min-height: 400px;
  height: 500px;
  display: flex;
  flex-direction: column;
  overflow: hidden; /* 防止内容溢出 */
  .not-support{
    display: flex;
    align-items: center;
    flex-direction: column;
    justify-content: center;
    svg{
      margin-bottom: 12px;
    }
  }
  .img-c{
    position: absolute; /* 绝对定位，覆盖整个父容器 */
    top: 45px;
    left: 0;
    right: 0;
    bottom: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 20px;
    /* 不需要overflow: hidden; 因为object-fit:contain已经确保不溢出 */
    // flex: 1; /* 撑满剩余空间 */
    // display: flex;
    // justify-content: center;
    // align-items: center;
    // overflow: hidden;
    img{
      max-width: 100%;
      max-height: 100%;
      object-fit: contain; /* 保持比例 */
      width: auto;
      height: auto;
      display: block;
    }
  }
  .file-content{
    flex: 1;
    overflow: auto;
    height: 100%;
    width: 100%;
  }
}

/* 通用滚动条样式 */
::-webkit-scrollbar {
  width: 4px;
  height: 4px;
}
::-webkit-scrollbar-track {
  background: #f1f1f1;
}
::-webkit-scrollbar-thumb {
  background: #88888880;
  border-radius: 3px;
}
@media only screen and (max-width: 768px) {
  .upload-wrap{
    margin-left: 0 !important;
    margin-right: 0 !important;
  }
}
</style>