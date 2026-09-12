<template>
  <div class="dataset-file-uploader">
    <div class="content">
      <input type="file" ref="fileInputRef" style="display:none;" accept=".zip" @change="fileChange">
      <div class="file-list-c" :class="errStatus ? 'error' : ''">
        <div class="file-item" v-for="(item, index) in selectList" :key="item.index">
          {{ item.name }};
        </div>
        <div v-if="selectList.length == 0" class="file-item-placeholder">{{ '请选择本地数据集文件' }}</div>
      </div>
      <div class="btn-select" @click="selectClick">
        <i class="el-icon-plus"></i>
        <span>选择本地数据集文件</span>
      </div>
    </div>
  </div>
</template>

<script>

import SparkMD5 from "spark-md5";
import { getChunks, getNewMultipart, getMultipartUrl, setCompleteMultipart } from '~/apis/modules/dataset';

const uploadChunkSize = 1024 * 1024 * 64;
const md5ChunkSize = 1024 * 1024 * 64;
const calcMd5ChunkSize = 1024 * 1024 * 1;
const maxFileSize = 20;

export default {
  name: "DatasetFileUploader",
  props: {
    type: { type: Number, default: 0 },
    datasetId: { type: String, default: '' },
  },
  data() {
    return {
      selectList: [],

      uploading: false,
      uploadFiles: [],
      uploadLength: 0,
      uploadSuccessLength: 0,
      uploadStatusList: [],

      resolveHandler: null,
      rejectHandler: null,

      errStatus: false,
    };
  },
  watch: {},
  methods: {
    fileChange(evt) {
      const files = evt.target.files || [];
      this.selectList = [...files];
      if (this.selectList.length) {
        this.errStatus = false;
      }
    },
    selectClick() {
      this.$refs.fileInputRef.value = '';
      this.$refs.fileInputRef.click();
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
        uploadUuid: file.__uuid,
        name: file.name,
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
          console.warn(file.name + ': calcFileMd5 went wrong.');
          reject(e);
        };

        function loadNext() {
          const start = currentChunk * chunkSize;
          const end = ((start + calcMd5ChunkSize) >= file.size) ? file.size : start + calcMd5ChunkSize;
          fileReader.readAsArrayBuffer(blobSlice.call(file, start, end));
        }
        loadNext();
      });
    },
    getChunksInfo(file) {
      return getChunks({
        md5: file.uniqueIdentifier,
        type: this.type,
        dataset_id: this.datasetId,
        file_name: file.name
      }).then(res => {
        const data = res.data;
        file.uploadID = data.uploadID;
        file.uuid = data.uuid;
        file.uploaded = data.uploaded;
        file.chunks = data.chunks;
        file.attachID = data.attachID;
        file._datasetID = data.datasetID;
        file._datasetName = data.datasetName;
        file._fileName = data.fileName;
        return file;
      }).catch(err => {
        console.info('getChunksInfo', err);
        return err;
      });
    },
    newUpload(file) {
      return getNewMultipart({
        totalChunkCounts: file.totalChunkCounts,
        md5: file.uniqueIdentifier,
        size: file.size,
        fileType: file.type,
        type: this.type,
        file_name: file.name,
        dataset_id: this.datasetId,
      }).then(res => {
        const data = res.data;
        file.uploadID = data.uploadID;
        file.uuid = data.uuid;
        if (file.uploadID && file.uuid) {
          file.chunks = '';
          this.breakpointUpload(file);
        } else {
          console.log('getNewMultipart Error', file);
          this.uploadError(file, info);
          this.updateFileStatus(file, this.$t('modelManage.uploadFailed'), 0, 2);
        }
        return file;
      }).catch(err => {
        console.log('getNewMultipart', err);
        this.uploadError(file, info);
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
      const successChunks = [];
      const successParts = file.chunks.split(",");
      for (let i = 0; i < successParts.length; i++) {
        successChunks[i] = successParts[i].split("-")[0];
      }
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
        const res = await getMultipartUrl({
          uuid: file.uuid,
          uploadID: file.uploadID,
          size: partSize,
          chunkNumber: currentChunk + 1,
          type: _this.type,
          file_name: file.name,
          dataset_id: this.datasetId,
        });
        urls[currentChunk] = res.data.url;
      };

      const uploadMinioNewMethod = async (url, e) => {
        const xhr = new XMLHttpRequest();
        xhr.open("PUT", url, false);
        if (_this.type == 0) {
          xhr.setRequestHeader("Content-Type", "text/plain");
          xhr.send(e.target.result);
          const etagValue = xhr.getResponseHeader("etag");
          etags[currentChunk] = etagValue;
        } else if (_this.type == 1) {
          xhr.setRequestHeader("Content-Type", "");
          xhr.send(e.target.result);
          const etagValue = xhr.getResponseHeader("ETag");
          etags[currentChunk] = etagValue;
        }
      }

      const uploadChunk = async (e) => {
        try {
          if (!checkSuccessChunks()) {
            const start = currentChunk * chunkSize;
            const partSize = start + chunkSize >= file.size ? file.size - start : chunkSize;
            // 获取分片上传url
            await getUploadChunkUrl(currentChunk, partSize);
            if (urls[currentChunk] != '') {
              await uploadMinioNewMethod(urls[currentChunk], e);
              if (etags[currentChunk] != '') {
              } else {
                console.log("上传到minio uploadChunk etags[currentChunk] == ''");
              }
            } else {
              console.log("uploadChunk urls[currentChunk] != ''");
            }
          }
        } catch (error) {
          console.log(error);
        }
      }

      const completeUpload = async () => {
        return await setCompleteMultipart({
          uuid: file.uuid,
          uploadID: file.uploadID,
          file_name: file.name,
          size: file.size,
          type: _this.type,
          dataset_id: _this.datasetId,
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
              await completeUpload();
              console.log(`文件上传完成：${file.name} \n分片：${chunks} 大小:${file.size} 用时：${(new Date().getTime() - time) / 1000} s`);
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
            }
          }
        } catch (err) {
          console.log(err);
          const info = this.$t('modelManage.uploadFailed');
          console.log(info, file)
          this.uploadLength++;
          this.uploadError(file, info);
          this.updateFileStatus(file, info, Number(((currentChunk / chunks) * 100).toFixed(2)) - 1, 2);
        }
      };
      console.log("上传分片...");
      loadNext();
    },
    resetFileStatus() {
      this.uploadFiles = [];
      this.uploadLength = 0;
      this.uploadSuccessLength = 0;
      this.uploadStatusList = [];
    },
    updateFileStatus(file, status, progress, infoCode, failedInfo = "") {
      this.uploadStatusList.forEach((item, index) => {
        if (item.uploadUuid === file.__uuid) {
          this.uploadStatusList[index].status = status;
          this.uploadStatusList[index].progress = progress;
          this.uploadStatusList[index].infoCode = infoCode;
          this.uploadStatusList[index].failedInfo = failedInfo;
        }
      });
    },
    uploadError(file, info) {
      this.uploadFinishCheck(file);
    },
    uploadSuccess(file) {
      this.uploadFinishCheck(file);
    },
    uploadFinishCheck(file) {
      console.log('uploadFinishCheck', file, this.uploadLength, '/', this.uploadFiles.length);
      if (this.uploadLength === this.uploadFiles.length && this.uploadFiles.length != 0) {
        console.log('All file has finish, success ' + this.uploadSuccessLength);
        this.uploading = false;
        this.resolveHandler && this.resolveHandler({
          status: '0',
          total: this.uploadFiles.length,
          success: this.uploadSuccessLength,
          msg: `处理完成，总数${this.uploadFiles.length},成功${this.uploadSuccessLength}`,
          files: this.uploadFiles,
          filesStatus: this.uploadStatusList,
        });
      }
    },
    checkFiles(file) {
      return true;
    },
    upload() {
      if (this.uploading) return;
      return new Promise((resolve, reject) => {
        this.resolveHandler = resolve;
        this.rejectHandler = reject;
        const fileList = this.selectList;
        if (!fileList.length) {
          this.$message({
            type: 'error',
            message: '请先选择文件',
          });
          resolve({
            status: '1',
            msg: '请选择文件',
          });
          return;
        };
        for (let i = 0, iLen = fileList.length; i < iLen; i++) {
          if (!this.checkFiles(fileList[i], true)) {
            resolve({
              status: '1',
              msg: '文件不合法'
            });
            return;
          }
        }
        this.resetFileStatus();
        this.uploadFiles = fileList;
        this.uploading = true;
        for (let i = 0, iLen = fileList.length; i < iLen; i++) {
          const file = fileList[i];
          file.dataset_id = this.datasetId;
          file.__uuid = Math.random();
          this.calcFileMd5(file).then(res => { // 计算MD5
            this.getChunksInfo(file).then(res => { // 获取Chunk信息
              if (file.uploadID == '' || file.uuid == '') { // 未上传过
                this.newUpload(file);
              } else if (file.uploaded == '1') { // 已上传成功 
                if (file.attachID == '0') { // 删除数据集记录，未删除文件
                  console.log(`file.attachID == '0'`);
                }
                this.uploadLength++;
                // 同一数据集上传同一个文件
                if (file._datasetID) {
                  const info = `${this.$t('modelManage.fileHasAlreadyInTheModel')} ${file._modelName}`;
                  this.uploadError(file, info);
                  this.updateFileStatus(file, this.$t('modelManage.uploadFailed'), 0, 1, info);
                } else { // 秒传
                  this.uploadSuccessLength++;
                  this.updateFileStatus(file, this.$t('modelManage.uploadSuccess'), 100, 0);
                  this.uploadSuccess(file);
                }
                console.log(file.name, '文件处理完成');
              } else { // 断点续传
                this.breakpointUpload(file);
              }
            }).catch(err => {
              console.info('getChunksInfo', err);
              this.uploadLength++;
              this.uploadError(file, this.$t('modelManage.uploadFailed'));
              this.updateFileStatus(file, this.$t('modelManage.uploadFailed'), 0, 2);
            });
          }).catch(err => {
            console.info('calcFileMd5', err);
            this.uploadLength++;
            this.uploadError(file, this.$t('modelManage.uploadFailed'));
            this.updateFileStatus(file, this.$t('modelManage.uploadFailed'), 0, 2);
          });
        }
      });

    },
    check() {
      const fileList = this.selectList;
      if (!fileList.length) {
        this.errStatus = true;
        return false;
      }
      for (let i = 0, iLen = fileList.length; i < iLen; i++) {
        if (!this.checkFiles(fileList[i])) {
          this.errStatus = true;
          return fasle;
        }
      }
      this.errStatus = false;
      return true;
    },
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
.dataset-file-uploader {
  display: flex;

  .content {
    flex: 1;
    display: flex;

    .file-list-c {
      margin-right: 5px;
      flex: 1;
      min-height: 32px;
      border-radius: 4px;
      border: 1px solid #DCDFE6;
      box-sizing: border-box;
      color: #606266;
      padding: 0 15px;

      .file-item {
        line-height: 32px;
        font-size: 13px;
      }

      .file-item-placeholder {
        line-height: 32px;
        color: rgba(0, 0, 0, 0.6);
        opacity: 0.45 !important;
        font-size: 13px;
      }

      &.error {
        color: #9f3a38;
        background: #fff6f6;
        border-color: #e0b4b4;
      }
    }

    .btn-select {
      display: flex;
      justify-content: center;
      align-items: center;
      cursor: pointer;
      color: rgb(3, 102, 214);

      i {
        margin-right: 2px;
      }
    }
  }
}
</style>
