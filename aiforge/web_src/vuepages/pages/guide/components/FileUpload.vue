<template>
    <div class="ui form">
        <div class="row-c">
            <div class="row">
                <div class="r-title"><label class="required">{{ $t('modelManage.uploadPath') }}</label></div>
                <div class="r-content">
                    <el-input size="medium" class="input-disabled" :value="uploadDir || '/'" readonly>
                    </el-input>
                </div>
            </div>
           
            <div class="row">
                <div class="r-title"></div>
                <div class="r-content">
                    <el-radio-group class="upload-type-sel" v-model="uploadType" :disabled="uploading"
                    @change="changeUploadType">
                    <el-radio :label="'file'">{{ $t('modelManage.uploadFile') }}</el-radio>
                    <el-radio :label="'folder'">{{ $t('modelManage.uploadFolder') }}</el-radio>
                    </el-radio-group>
                </div>
            </div>
            <div class="row" v-if="storageShow">
              <div class="r-title"></div>
              <div class="r-content" style="display:flex;line-height: 28px;flex-wrap: wrap;">
                <div class="storage-t">
                  {{ $t('storage.remain_storage') }}：<span class="storage-v">{{formattedRemaining[0]}}</span> {{formattedRemaining[1]}}
                  ({{ $t('storage.selected_file_size') }}：
                </div>
                <div class="storage-limit-wrap">
                  <span class="select-v">{{formattedTotal[0]}}</span>{{formattedTotal[1]}}
                  <div class="limit-tip-wrap" v-if="isStorageExceeded">
                    <div class="limit-tip"><i class="ri-information-line"></i>{{ $t('storage.exceedStorage') }}</div>
                    <a v-if="!showOwenerTips" href="/storages">{{ $t('storage.capacity_details') }}</a>
                    <span v-else>{{ $t('storage.owenerTips',{ownerName:ownerName}) }}</span>
                  </div>
                  )
                </div>
              </div>
            </div>
            <div class="row" style="align-items:flex-start;">
                <div class="r-title">
                    <label class="required" v-if="uploadType == 'file'">{{ $t('modelManage.fileUpload') }}</label>
                    <label class="required" v-if="uploadType == 'folder'">{{ $t('modelManage.folderUpload') }}</label>
                </div>
                <div class="r-content">
                    <div style="position:relative" v-show="uploadType == 'file'">
                    <form class="dropzone" ref="dropzoneRef">
                        <div class="dropzon-err-tips ui red message" v-show="showUploadErr"
                        style="display:none;margin:2.5rem" v-html="uploadErrTxt">
                        </div>
                    </form>
                    <div class="tips" v-html="getDefaultErrTxt()"></div>
                    <div class="not-allowed-placeholder" v-show="uploading"></div>
                    </div>
                    <FolderUploadSelect ref="folderUploadSelect" v-if="uploadType == 'folder'"
                    :maxFilesSize="maxModelFilesSize" :uploading="uploading" @folderSelectChange="folderSelectChange">
                    </FolderUploadSelect>
                </div>
            </div>
            <div class="row" style="margin-top:10px">
                <div class="r-title"><label></label></div>
                <div class="r-content">
                    <!-- <el-button type="info" @click="back" :disabled="uploading">{{ $t('modelManage.back') }}</el-button> -->
                    <el-button size="medium" round class="green" @click="submit" :disabled="uploading || btnFlag || isStorageExceeded">{{ $t('modelManage.upload')}}</el-button>
                    <el-button size="medium" round type="info" @click="cancel" :disabled="uploading">{{ $t('modelManage.cancel') }}</el-button>
                </div>
            </div>
            <div class="row" style="align-items:flex-start;">
                <div class="r-title"><label>{{ $t('modelManage.uploadStatus') }}：</label></div>
                <div class="r-content">
                    <div v-for="(item, index) in uploadFiles" :key="item.upload.uuid" class="datast-upload-progress">
                        <span class="dataset-name nowrap" :title="item.fullname">{{ item.fullname }}</span>
                        <div class="dataset-progress">
                            <el-progress :text-inside="true" :stroke-width="14" :percentage="uploadStatusList[index].progress">
                            </el-progress>
                        </div>
                        <div class="dataset-status nowrap">
                            <div class="status-flex">
                            <i v-if="uploadStatusList[index].infoCode === 1 || uploadStatusList[index].infoCode === 2"
                                class="ri-close-circle-line failed"></i>
                            <i v-if="uploadStatusList[index].infoCode === 0" class="ri-checkbox-circle-line success"></i>
                            <span>{{ uploadStatusList[index].status }}</span>
                            <el-tooltip v-if="uploadStatusList[index].infoCode === 1" class="item" effect="dark"
                                placement="top">
                                <div slot="content"> {{ uploadStatusList[index].failedInfo }} </div>
                                <i style="font-size: 16px; margin-left: 0.5rem; cursor: pointer" class="ri-question-fill"></i>
                            </el-tooltip>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>

import NotFound from '~/components/NotFound.vue';
import 'dropzone/dist/dropzone.css';
import Dropzone from 'dropzone';
import SparkMD5 from "spark-md5";
import { getChunks, getNewMultipart, getMultipartUrl, setCompleteMultipart } from '~/apis/modules/dataset';
import {  transFileSize } from '~/utils';
import { getStorageSummary } from "~/apis/modules/storage";
import FolderUploadSelect from '~/pages/modelmanage/components/FolderUploadSelect.vue';
Dropzone.autoDiscover = false;

const uploadChunkSize = 1024 * 1024 * 64;
const md5ChunkSize = 1024 * 1024 * 64;
const maxFileCount = 100 ;
const maxModelFilesSize = window.MAX_MODEL_SIZE || 536870912; // 200 GB
const UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB'];

export default {
    props: {
        ownerName: { type: String, default: '' },
        modelId: { type: String, default: '' },
        uploadDir: { type: String, default: '' },
        subjectType: { type: String, default: '1' },
    },
    data() {
      return {
        emptyPage: false,
        modelData: {},
        uploadType: 'file', // file|folder
        dropzoneHandler: null,
        type: '1', // 1-修改，其它-新增
        state: {
          type: '',
          id: '',
          name: '',
          version: '',
          engine: '',
          label: '',
          description: '',
          size: 0,
        },
        originData: null,
        showUploadErr: false,
        uploadErrTxt: '',
  
        uploadFiles: [],
        uploadLength: 0,
        uploadSuccessLength: 0,
        uploadStatusList: [],
        maxModelFilesSize: maxModelFilesSize,
        uploading: false,
        btnFlag: false,
        
        remaining_storage: 0,
        totalSize: 0,
        isStorageExceeded: false,
        storageShow: false,
        showOwenerTips: false,    
        canShowError: false
      };
    },
    components: {  NotFound,FolderUploadSelect },
    computed: {
        // 格式化的剩余空间显示（自动单位转换）
        formattedRemaining() {
            return this.formatBytes(Math.max(this.remaining_storage,0));
        },
        formattedTotal() {
            return this.formatBytes(this.totalSize);
        },
        formatUploadir(){
            if(this.uploadDir){
                return this.uploadDir.slice(1) + '/';
            }else{
                return '';
            }
        }
    },
    methods: {
        initDropZone() {
            let previewTemplate = `
                <div class="dz-preview dz-file-preview"> 
                <div class="dz-image"> 
                    <img data-dz-thumbnail /> 
                </div> 
                <div class="dz-details"> 
                    <div class="dz-size"><span data-dz-size></span></div> 
                    <div class="dz-filename"><span data-dz-name></span></div> 
                </div> 

                <div style="opacity:0" class="dz-progress"><span class="dz-upload" data-dz-uploadprogress></span></div> 
                <div class="dz-error-message" style="line-height: 1.5;"><span data-dz-errormessage></span></div> 
                <div class="dz-success-mark"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="54" height="54"><path fill="none" d="M0 0h24v24H0z"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10-4.477 10-10 10zm0-2a8 8 0 1 0 0-16 8 8 0 0 0 0 16zm-.997-4L6.76 11.757l1.414-1.414 2.829 2.829 5.656-5.657 1.415 1.414L11.003 16z" fill="rgba(47,204,113,1)"/></svg></div> 
                <div class="dz-error-mark"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="54" height="54"><path fill="none" d="M0 0h24v24H0z"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10-4.477 10-10 10zm0-2a8 8 0 1 0 0-16 8 8 0 0 0 0 16zm0-9.414l2.828-2.829 1.415 1.415L13.414 12l2.829 2.828-1.415 1.415L12 13.414l-2.828 2.829-1.415-1.415L10.586 12 7.757 9.172l1.415-1.415L12 10.586z" fill="rgba(231,76,60,1)"/></svg></div> 
                </div> `;
            let vm = this
            this.dropzoneHandler = new Dropzone(this.$refs.dropzoneRef, {
                url: '/',
                maxFiles: maxFileCount,
                parallelUploads: 20,
                uploadMultiple: true,
                filesizeBase: 1024,
                maxFilesize: this.maxModelFilesSize/(1024*1024),
                timeout: 0,
                accept: function(file, done) {
                    if (file.name.length > 128) {
                        vm.btnFlag = true;
                        vm.showUploadErrInfo(true, vm.getDefaultErrTxt());
                        done(vm.$t('modelManage.modelFileNameTips'))
                    }
                    else {
                        done()
                    }
                },
                addRemoveLinks: true,
                autoProcessQueue: false,
                dictDefaultMessage: this.$t('modelManage.modelFileUploadDefaultTips'),
                dictFileTooBig: this.$t('modelManage.fileIstoBig'),
                dictRemoveFile: this.$t('modelManage.removeFile'),
                dictMaxFilesExceeded: this.getDefaultErrTxt(),
                previewTemplate: previewTemplate,
            });
            this.dropzoneHandler.on("addedfile", file => {
                file.fullname = file.name;
                this.checkFiles(file);
                setTimeout(() => this.calculateTotalSize([]), 0);
            });
            this.dropzoneHandler.on("removedfile", file => {
                if (this.dropzoneHandler.getRejectedFiles().length === 0) {
                    this.showUploadErrInfo(false);
                    this.btnFlag = false
                }
                this.checkFiles();
                setTimeout(() => this.calculateTotalSize([]), 0);
            });
        },
        getDefaultErrTxt() {
            return this.$t('modelManage.modelFileUploadErrTips', { maxCount: maxFileCount, size: transFileSize(maxModelFilesSize) ,url:'https://openi.pcl.ac.cn/docs/index.html#/dataset/sdk'});
        },
        showUploadErrInfo(state, info) {
            if (state) {
            this.uploadErrTxt = info;
            this.showUploadErr = true;
            } else {
            this.uploadErrTxt = '';
            this.showUploadErr = false;
            }
        },
        changeUploadType() {
            this.dropzoneHandler.removeAllFiles();
            this.resetFileStatus();
        },
        checkFiles(file, countStay) {
            const fileList = this.dropzoneHandler.getAcceptedFiles();
            fileList.forEach(item => item.fullname = item.name);
            const filesCount = fileList.length + (file && !countStay ? 1 : 0);
            if (file && filesCount > maxFileCount) {
                this.dropzoneHandler.removeFile(file);
                this.showUploadErrInfo(true, this.getDefaultErrTxt());
                return false;
            }
            if (file && file.size/(1024*1024) > this.dropzoneHandler.options.maxFilesize) {
                this.btnFlag = true;
                this.showUploadErrInfo(true, this.getDefaultErrTxt());
                return false;
            }
            return true;
        },
        resetFileStatus() {
            this.uploadFiles = [];
            this.uploadLength = 0;
            this.uploadSuccessLength = 0;
            this.uploadStatusList = [];
            this.canShowError = false
        },
        updateFileStatus(file, status, progress, infoCode, failedInfo = "") {
            this.uploadStatusList.forEach((item, index) => {
            if (item.uploadUuid === file.upload.uuid) {
                this.uploadStatusList[index].status = status;
                this.uploadStatusList[index].progress = progress;
                this.uploadStatusList[index].infoCode = infoCode;
                this.uploadStatusList[index].failedInfo = failedInfo;
            }
            });
        },
        calcFileMd5(file) {
            const blobSlice = File.prototype.slice || File.prototype.mozSlice || File.prototype.webkitSlice;
            const chunkSize = md5ChunkSize;
            const chunks = Math.ceil(file.size / chunkSize);
            const spark = new SparkMD5.ArrayBuffer();
            const fileReader = new FileReader();
            file.totalChunkCounts = chunks;
            if (file.size == 0) {
                file.totalChunkCounts = 1;
            }
            let currentChunk = 0;
            this.uploadStatusList.push({
                uploadUuid: file.upload.uuid,
                name: file.fullname,
                status: this.$t('modelManage.calcFileMd5'),
                progress: 0,
                infoCode: 3,
            });
            return new Promise((resolve, reject) => {
                fileReader.onload = function (e) {
                    spark.append(e.target.result);
                        currentChunk++;
                    if (currentChunk < chunks) {
                        loadNext();
                    } else {
                        const md5 = spark.end();
                        spark.destroy();
                        file.uniqueIdentifier = md5;
                        resolve(md5);
                    }
                };
                fileReader.onerror = function (e) {
                    console.warn(file.fullname + ': calcFileMd5 went wrong.');
                    reject(e);
                };
        
                function loadNext() {
                    const start = currentChunk * chunkSize;
                    const end = ((start + chunkSize) >= file.size) ? file.size : start + chunkSize;
                    fileReader.readAsArrayBuffer(blobSlice.call(file, start, end));
                }
                loadNext();
            });
        },
        getChunksInfo(file) {
            return getChunks({
                md5: file.uniqueIdentifier,
                file_name: `${this.formatUploadir}${file.fullname}`,
                subject_id: this.modelId,
                subject_type: this.subjectType,
            }).then(resonpse => {
                console.log('getChunksInfo', resonpse);
                const res = resonpse.data;
                if(res.code === 0){
                    console.log("xxxxxxxxxxxxxx",res)
                    file.uuid = res.data.uuid;
                    file.uploaded = res.data.uploaded;
                    file.chunks = res.data.chunks;
                    file.realName = res.data.fileName;
                    return file;
                }else{
                    throw new Error(res.msg); // 抛出一个错误
                }
                
            }).catch(err => {
                console.info('getChunksInfo', err);
                return err;
            });
        },
        newUpload(file) {
            return getNewMultipart({
                md5: file.uniqueIdentifier,
                file_name: `${this.formatUploadir}${file.fullname}`,
                subject_id: this.modelId,
                subject_type: this.subjectType,
                file_type: file.type,
                size: file.size,
                total_chunk_counts: file.totalChunkCounts,
            }).then(resonpse => {
                const res = resonpse.data;
                if(res.code === 0){
                    file.uuid = res.data.uuid;
                    if ( file.uuid ) {
                        file.chunks = [];
                        this.breakpointUpload(file);
                    } else {
                        this.uploadError(file, info);
                        this.updateFileStatus(file, this.$t('modelManage.uploadFailed'), 0, 2);
                    }
                    return file;
                } else {
                    if (this.canShowError) {
                        this.canShowError = false
                        throw new Error(res.msg); // 抛出一个错误
                    }
                    
                }
            }).catch(err => {
                console.log('getNewMultipart', err);
                this.$message({
                    type: 'error',
                    message: err.message,
                });
                this.uploading = false
                this.uploadError(file, this.$t('modelManage.uploadFailed'));
                this.updateFileStatus(file, this.$t('modelManage.uploadFailed'), 0, 2);
                return err;
            });
        },
        breakpointUpload(file) {
            const blobSlice = File.prototype.slice || File.prototype.mozSlice || File.prototype.webkitSlice;
            const fileReader = new FileReader();
            const time = new Date().getTime();
            const chunkSize = uploadChunkSize;
            const chunks = Math.ceil(file.size / chunkSize);
            const _this = this;
            let currentChunk = 0;
            const successChunks = file.chunks;
            // const successParts = file.chunks.split(",");
            // for (let i = 0; i < successParts.length; i++) {
            //     successChunks[i] = successParts[i].split("-")[0];
            // }
            const urls = [];
            const etags = [];
            const checkSuccessChunks = () => {
                const index = successChunks.indexOf((currentChunk + 1).toString());
                if (index == -1) {
                    return false;
                }
                return true;
            }
    
            const getUploadChunkUrl = async (currentChunk, partSize) => {
                try {
                    const resonpse = await getMultipartUrl({
                        uuid: file.uuid,
                        size: partSize,
                        chunk_number: currentChunk + 1,
                    });
                    const res = resonpse.data;
                    if(res.code  == 0){
                        urls[currentChunk] = res.data.url;
                    }else{
                        throw new Error(res.message);
                    }
                    console.log("getUploadChunkUrl",urls,etags)
                } catch (error) {
                    this.$message.error(error);
                }
            };
    
            const uploadMinioNewMethod = async (url, e) => {
                const xhr = new XMLHttpRequest();
                xhr.open("PUT", url, false);
                xhr.send(e.target.result);
                // if (_this.state.type == 0) {
                //     xhr.setRequestHeader("Content-Type", "text/plain");
                    
                //     const etagValue = xhr.getResponseHeader("etag");
                //     etags[currentChunk] = etagValue;
                // } else if (_this.state.type == 1) {
                //     xhr.setRequestHeader("Content-Type", "");
                //     xhr.send(e.target.result);
                //     const etagValue = xhr.getResponseHeader("ETag");
                //     etags[currentChunk] = etagValue;
                // }
            }
    
            const uploadChunk = async (e) => {
                try {
                    if (!checkSuccessChunks()) {
                        const start = currentChunk * chunkSize;
                        const partSize = start + chunkSize >= file.size ? file.size - start : chunkSize;
                        // 获取分片上传url
                        await getUploadChunkUrl(currentChunk, partSize);
                    if (urls[currentChunk] != '') {
                        console.log("ssssssssssssssssssss")
                        await uploadMinioNewMethod(urls[currentChunk], e);
                        if (etags[currentChunk] != '') {

                        } else {
                            console.log("上传到minio uploadChunk etags[currentChunk] == ''"); // TODO
                        }
                    } else {
                        console.log("uploadChunk urls[currentChunk] != ''"); // TODO
                    }
                    }
                } catch (error) {
                    console.log(error);
                    this.$message({
                        type: 'error',
                        message: err,
                    });
                }
            }
    
            const completeUpload = async () => {
                return await setCompleteMultipart({
                    uuid: file.uuid,
                });
            }
    
            function loadNext() {
                const start = currentChunk * chunkSize;
                const end = start + chunkSize >= file.size ? file.size : start + chunkSize;
                fileReader.readAsArrayBuffer(blobSlice.call(file, start, end));
            }
    
            fileReader.onload = async (e) => {
            try {
                await uploadChunk(e);
                fileReader.abort();
                currentChunk++;
                if (currentChunk < chunks) {
                    console.log(`第${currentChunk}个分片上传完成, 开始第${currentChunk + 1}/${chunks}个分片上传`);
                    this.updateFileStatus(file, this.$t('modelManage.uploading'), Number(((currentChunk / chunks) * 100).toFixed(2)), 3);
                    loadNext();
                } else {
                    try {
                        const response = await completeUpload();
                        console.log("xxxxxxxxxxa",response)
                        console.log(`文件上传完成：${file.fullname} \n分片：${chunks} 大小:${file.size} 用时：${(new Date().getTime() - time) / 1000} s`);
                        this.uploadLength++;
                        this.uploadSuccessLength++;
                        this.updateFileStatus(file, this.$t('modelManage.uploadSuccess'), 100, 0);
                        this.uploadSuccess(file);
                    } catch (err) {
                        const info = this.$t('modelManage.uploadFailed');
                        console.log(info, file)
                        this.uploadLength++;
                        this.uploadError(file, info);
                        this.updateFileStatus(file, info, Number(((currentChunk / chunks) * 100).toFixed(2)) - 1, 2);
                        this.$message({
                            type: 'error',
                            message: err,
                        });
                    }
                }
            } catch (err) {
                console.log(err);
                const info = this.$t('modelManage.uploadFailed');
                console.log(info, file)
                this.uploadLength++;
                this.uploadError(file, info);
                this.updateFileStatus(file, info, Number(((currentChunk / chunks) * 100).toFixed(2)) - 1, 2);
                this.$message({
                    type: 'error',
                    message: err,
                });
            }
            };
            console.log("上传分片...");
            loadNext();
        },
        uploadError(file, info) {
            file.status = 'error';
            file.errTips = info;
            if (file.previewTemplate) {
                file.previewTemplate.querySelector('.dz-success-mark').style.opacity = 0;
                file.previewTemplate.querySelector('.dz-error-mark').style.opacity = 1;
                file.previewTemplate.querySelector('.dz-error-message span').innerHTML = info;
                file.previewTemplate.querySelector(".dz-error-message").style.display = 'block';
                file.previewTemplate.querySelector(".dz-details").onmouseover = () => {
                    file.previewTemplate.querySelector('.dz-error-message').style.opacity = 1;
                };
                file.previewTemplate.querySelector(".dz-details").onmouseout = () => {
                    file.previewTemplate.querySelector('.dz-error-message').style.opacity = 0;
                };
            }
            this.uploadFinishCheck(file);
        },
        uploadSuccess(file) {
            file.status = 'success';
            file.errTips = '';
            if (file.previewTemplate) {
                file.previewTemplate.querySelector('.dz-error-mark').style.opacity = 0;
                file.previewTemplate.querySelector('.dz-error-message span').innerHTML = '';
                file.previewTemplate.querySelector('.dz-success-mark').style.opacity = 1;
                file.previewTemplate.querySelector(".dz-error-message").style.display = 'none';
                file.previewTemplate.querySelector(".dz-details").onmouseover = null;
                file.previewTemplate.querySelector(".dz-details").onmouseout = null;
            }
            this.uploadFinishCheck(file);
        },
        uploadFinishCheck(file) {
            console.log('uploadFinishCheck', file, this.uploadLength, '/', this.uploadFiles.length);
            if (this.uploadLength === this.uploadFiles.length && this.uploadFiles.length != 0) {
                console.log('All file has finish, success ' + this.uploadSuccessLength);
                this.uploading = false;
                if (this.uploadSuccessLength == this.uploadLength) {
                    this.$emit('uploadFinish')
                    // window.setTimeout(() => {
                    //     location.href = `/explore/datasets/${this.modelId}?tab=files`
                    // }, 1000);
                } else {
                    if (this.uploadSuccessLength > 0) {
                        
                    }
                }
            }
        },
        submit() {
            let fileList = [];
            if (this.uploadType == 'file') {
                fileList = this.dropzoneHandler.getAcceptedFiles();
                if (!fileList.length) return;
                for (let i = 0, iLen = fileList.length; i < iLen; i++) {
                    if (!this.checkFiles(fileList[i], true)) return;
                }
            } else if (this.uploadType == 'folder') {
                fileList = this.$refs['folderUploadSelect'].getAcceptedFiles();
                if (!fileList.length) return;
                if (!this.$refs['folderUploadSelect'].checkFiles()) return;
            }
            this.resetFileStatus();
            this.canShowError = true
            this.uploadFiles = fileList;
            this.uploading = true;
            const fileNameList = [];
            for (let i = 0, iLen = fileList.length; i < iLen; i++) {
                const file = fileList[i];
                // file.modelUuid = this.state.id;
                // file.modelName = this.state.name;
                this.calcFileMd5(file).then(res => { // 计算MD5
                    if (fileNameList.indexOf(file.fullname) > -1) {
                        var info = `${this.$t('modelManage.fileExistInTheModel')} ${file.fullname}`;
                        console.log("info=" + info);
                        this.uploadLength++;
                        this.uploadError(file, this.$t('modelManage.uploadFailed'));
                        this.updateFileStatus(file, this.$t('modelManage.uploadFailed'), 0, 1, info);
                    } else {
                        fileNameList.push(file.fullname);
                        this.getChunksInfo(file).then(res => { // 获取Chunk信息
                            console.log(file)
                            if (file.uuid == '') { // 未上传过
                                console.log(`file.uuid == ''`,file.uuid == '');
                                this.newUpload(file);
                            } else if (file.uploaded == '1') { // 已上传成功 
                                // if (file.attachID == '0') { // 删除数据集记录，未删除文件
                                //     // await addAttachment(file);
                                //     console.log(`file.attachID == '0'`);
                                // }
                                this.uploadLength++;
                                // 同一模型上传同一个文件
                                if (file.subjectId) {
                                    const info = `${this.$t('modelManage.fileExistInTheModel')} ${file.realName}`;
                                    this.uploadError(file, info);
                                    this.updateFileStatus(file, this.$t('modelManage.uploadFailed'), 0, 1, info);
                                } else { // 秒传
                                    this.uploadSuccessLength++;
                                    this.updateFileStatus(file, this.$t('modelManage.uploadSuccess'), 100, 0);
                                    this.uploadSuccess(file);
                                }
                                console.log(file.fullname, '文件处理完成');
                            } else { // 断点续传
                                this.breakpointUpload(file);
                            }
                        }).catch(err => {
                            console.log("xxxxxxxxxxxxxssssss")
                            console.info('getChunksInfo', err);
                            this.uploadLength++;
                            this.uploadError(file, this.$t('modelManage.uploadFailed'));
                            this.updateFileStatus(file, this.$t('modelManage.uploadFailed'), 0, 2);
                            this.$message({
                                type: 'error',
                                message: err,
                            });
                        });
                    }
                }).catch(err => {
                    console.info('calcFileMd5', err);
                    this.uploadLength++;
                    this.uploadError(file, this.$t('modelManage.uploadFailed'));
                    this.updateFileStatus(file, this.$t('modelManage.uploadFailed'), 0, 2);
                });
            }
        },
        cancel() {
            this.$emit('cancel')
        },
        async getStorageSummary(){
            const res = await getStorageSummary({
                subject_id: this.modelId,
                subject_type: this.subjectType
            })
            this.remaining_storage = res.data.remaining_storage
            this.storageShow = res.data.storage_limit !=  -1
        },
        // 智能单位格式化
        formatBytes(bytes) {
            let unitIndex = 0;
            let value = bytes;
            
            while (value >= 1024 && unitIndex < UNITS.length - 1) {
                value /= 1024;
                unitIndex++;
            }
            return [this.toPrecision(value, 2),UNITS[unitIndex]]
        },
        // 精确小数处理
        toPrecision(value, decimals = 2) {
            if (value < 0.001 && value > 0) return '<0.001';
            return Number(value.toFixed(decimals)).toString();
        },
        calculateTotalSize(files) {
            if(!this.storageShow) return
            if(files.length===0){
                files = this.dropzoneHandler ? this.dropzoneHandler.getQueuedFiles() : []
            }
            this.totalSize = files.reduce((sum, file) => sum + file.size, 0)
            this.isStorageExceeded = this.totalSize > this.remaining_storage
        },
        folderSelectChange(fileList){
            this.calculateTotalSize(fileList)
        }
    },
    beforeMount() { },
    mounted() {
      this.initDropZone();
      this.getStorageSummary()
      const loginName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext')
      console.log('loginName', loginName,this.ownerName)
      if( this.ownerName != loginName ){
        this.showOwenerTips = true
      }
    },
    beforeDestroy() {
    },
  };
</script>
  
<style scoped lang="less">
  ::v-deep .el-button--info{
    color: #020004;
    background-color: #c2c7cc;
    border-color: #c2c7cc;
  }
  .row-c {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    margin: 0 auto;

    .row {
    width: 100%;
    display: flex;
    align-items: center;
    margin: 8px 0;
    margin-left: -190px;

    .r-title {
        text-align: right;
        font-size: .92857143em;
        color: #606266;;
        width: 200px;
        margin-right: 28px;
        position: relative;

        .required {
        &::after {
            position: absolute;
            margin: -0.2em 0 0 0.2em;
            content: '*';
            color: #db2828;
        }
        }
    }

    .r-content {
        flex: 1;
        overflow: hidden;
        .cluster-type-btn {
            display: flex;
            align-items: center;
            justify-content: center;
            border: 1px solid #DCDFE6;
            height: 36px;
            padding: 10px;
            cursor: pointer;
            border-radius: 4px;

        .icon {
            margin-right: 5px;
        }

        &.focused {
            border-color: rgb(50, 145, 248);
            color: rgb(50, 145, 248);
            cursor: default;

            .icon {
            :not([stroke]) {
                fill: rgb(50, 145, 248);
            }
            }
        }
        }
        .storage-t{
          font-family: Arial;
          color: rgba(2, 0, 4, 0.5);
          .storage-v{
            color: rgba(16, 16, 16, 1);
            font-weight: 700;
          }
        }
        .storage-limit-wrap{
          color: #101010;
          font-family: Arial;
          display: flex;
          margin: 0 6px;
          .select-v{
            color: rgba(246, 106, 0, 1);
            font-weight: 700;
            margin-right: 4px;
          }
          .limit-tip-wrap{
            display: flex;
            .limit-tip{
              display: flex;
              align-items: center;
              height: 28px;
              border-radius: 4px;
              background-color: rgba(250,140,22,1);
              color: rgba(255,255,255,1);
              padding: 0 8px;
              margin: 0 6px;
              i{
                font-size: 14px;
                margin-right: 6px;
              }
            }
            a{
              color: #101010;
              text-decoration:underline;
            }
          }
        }
    }
    }
}
  .not-allowed-placeholder {
    position: absolute;
    width: 100%;
    height: 100%;
    z-index: 20;
    cursor: no-drop;
    top: 0;
  }
  
  .tips {
    margin-top: 10px;
    font-size: 13px;
    color: rgba(0, 0, 0, 0.6);
  }
  
  .dropzone {
    min-height: 186px !important;
  
    /deep/ .dz-default.dz-message {
      margin: 76px 0 !important;
    }
  
    /deep/ .dz-remove {
      margin-top: 10px !important;
    }
  }
  
  .input-disabled {
    /deep/ .el-input__inner {
      background-color: #f5f5f6 !important;
      color: #888888 !important;
    }
  }
  
  /deep/ .dz-progress {
    display: none;
  }
  
  .success {
    color: #21ba45;
    font-size: 16px;
    margin-right: 0.5rem;
  }
  
  .failed {
    color: red;
    font-size: 16px;
    margin-right: 0.5rem;
  }
  
  .datast-upload-progress {
    display: flex;
    align-items: center;
  }
  
  .datast-upload-progress .dataset-name {
    text-align: right;
    width: 240px;
    margin-right: 1rem;
  }
  
  .datast-upload-progress .dataset-progress {
    flex: 1;
  }
  
  .datast-upload-progress .dataset-status {
    width: 140px;
    margin-left: 1rem;
  }
  
  .datast-upload-progress .dataset-status .status-flex {
    display: flex;
    align-items: center;
  }
  
  /deep/ .el-progress-bar__inner {
    background-color: #21ba45;
  }
  
  .el-select-dropdown__item.selected {
    color: #85b7d9;
  }
  
  /deep/ .el-button.el-button--medium {
    background-color: #e0e1e2;
    color: rgba(0, 0, 0, .6);
    border-color: transparent;
    transition: opacity .1s ease, background-color .1s ease, color .1s ease, box-shadow .1s ease, background .1s ease, -webkit-box-shadow .1s ease;
    will-change: auto;
    -webkit-tap-highlight-color: transparent;
  
    &:hover {
      border-color: transparent;
      background-color: #cacbcd;
      color: rgba(0, 0, 0, .8);
    }
  
    &:focus {
      background-color: #cacbcd;
      color: rgba(0, 0, 0, .8);
      border-color: transparent;
    }
  
    &:active {
      background-color: #babbbc;
      color: rgba(0, 0, 0, .9);
      border-color: transparent;
    }
  
    &.green {
      background-color: #0066ff;
      color: #fff;
  
      &:hover {
        background-color: #0066ff;
        border-color: transparent;
      }
  
      &:focus {
        background-color: #0066ff;
        border-color: transparent;
      }
  
      &:active {
        background-color: #0066ff;
        border-color: transparent;
      }
    }
  }
  
  /deep/ .el-select {
    .is-focus {
      .el-input__inner {
        border-color: #85b7d9;
      }
    }
  }
  
  /deep/ .el-input__inner {
    &:focus {
      border-color: #85b7d9;
    }
  }
  
  /deep/ .el-textarea__inner {
    &:focus {
      border-color: #85b7d9;
    }
  }

@media only screen and (max-width: 768px) {
  .row-c{
    padding: 10px;
    .row {
      margin-left: 0 !important;
      .r-title{
         width: auto !important;
      }
    }
  }
  /deep/ .dz-image{
    width: 80px !important;
    height: 100px !important;
  }
  /deep/ .dropzone .dz-preview.dz-file-preview .dz-details {
    padding: 20px 6px 20px 6px !important;
  }
  .dataset-name{
    width: auto !important;
  }
  .dataset-status{
    width: 80px !important;
  }
}

</style>
  