<template>
  <div class="dropzone-wrapper dataset-files">
    <div class="data-storage-limit">
      <div class="storage-t">
        {{ i18n.remain_storage }}：<span class="storage-v">{{formattedRemaining[0]}}</span> {{formattedRemaining[1]}}
        ({{ i18n.selected_file_size }}：
      </div>
      <div class="storage-limit-wrap">
        <span class="select-v">{{formattedTotal[0]}}</span>{{formattedTotal[1]}}
        <div class="limit-tip-wrap" v-if="isStorageExceeded">
          <div class="limit-tip"><i class="ri-information-line"></i>{{ i18n.exceedStorage }}</div>
          <a href="/storages">{{ i18n.capacity_details }}</a>
        </div>
        )
      </div>
    </div>
    <div id="dataset" class="dropzone">
      <div
        class="maxfilesize ui red message"
        style="display: none; margin: 1rem"
        v-html="errorTips"
      ></div>
    </div>
    <el-button
      style="background-color: #21ba45; margin-top: 2rem"
      type="success"
      :disabled="btnFlag || isStorageExceeded"
      @click="startUpload"
      >{{ upload }}</el-button
    >
    <el-button type="info" @click="cancelDataset">{{ cancel }}</el-button>
    <div style="margin-top: 2rem; position: relative">
      <label
        class="el-form-item__label"
        style="width: 140px; position: absolute; left: -140px"
        >{{ upload_status }}:</label
      >
      <div v-for="(item, index) in uploadFiles" class="datast-upload-progress">
        <span class="dataset-name nowrap" :title="item.name">{{
          item.name
        }}</span>
        <div class="dataset-progress">
          <el-progress
            :text-inside="true"
            :stroke-width="14"
            :percentage="uploadProgressList[index].progress"
          ></el-progress>
        </div>
        <div class="dataset-status nowrap">
          <div class="status-flex">
            <i
              v-if="
                uploadProgressList[index].infoCode === 1 ||
                uploadProgressList[index].infoCode === 2
              "
              class="ri-close-circle-line failed"
            ></i>
            <i
              v-if="uploadProgressList[index].infoCode === 0"
              class="ri-checkbox-circle-line success"
            >
            </i>
            <span>{{ uploadProgressList[index].status }}</span>
            <el-tooltip
              v-if="uploadProgressList[index].infoCode === 1"
              class="item"
              effect="dark"
              placement="top"
            >
              <div slot="content">
                {{ uploadProgressList[index].failedInfo }}
              </div>
              <i
                style="font-size: 16px; margin-left: 0.5rem; cursor: pointer"
                class="ri-question-fill"
              ></i>
            </el-tooltip>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SparkMD5 from "spark-md5";
import axios from "axios";
import qs from "qs";
import createDropzone from "../features/dropzone.js";

const { _AppSubUrl, _StaticUrlPrefix, csrf } = window.config;
const chunkSize = 1024 * 1024 * 64;
const md5ChunkSize = 1024 * 1024 * 1;
const UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB'];


export default {
  props: {
    uploadtype: {
      type: Number,
      required: true,
    },
    desc: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      dropzoneUploader: null,
      maxFiles: 10,
      acceptedFiles: "*/*",
      progress: 0,
      status: "",
      dropzoneParams: {},
      file_status_text: "",
      file: {},
      repoPath: "",
      btnFlag: false,
      cancel: "",
      upload: "",
      upload_status: "",
      uploadFiles: [],
      uploadFilesAddId: [],
      // allUploadFiles: [],
      uploadLength: 0,
      allUploadLength: 0,
      uploadProgressList: [],
      errorTips: '',
      failed:'',
      upload_success: '',

      remaining_storage: 0,
      totalSize: 0,
      isStorageExceeded: false
    };
  },
  computed: {
    // 格式化的剩余空间显示（自动单位转换）
    formattedRemaining() {
        return this.formatBytes(Math.max(this.remaining_storage,0));
    },
    formattedTotal() {
        return this.formatBytes(this.totalSize);
    },
  },
  async mounted() {
    this.dropzoneParams = $("div#minioUploader-params");
    
    this.file_status_text = this.dropzoneParams.data("file-status");
    this.status = this.dropzoneParams.data("file-init-status");
    this.repoPath = this.dropzoneParams.data("repopath");
    this.cancel = this.dropzoneParams.data("cancel");
    this.upload = this.dropzoneParams.data("upload");
    this.upload_status = this.dropzoneParams.data("upload-status");
    this.failed = this.dropzoneParams.data("failed");
    this.upload_success = this.dropzoneParams.data("upload-success");
    let maxModelFilesSize = this.dropzoneParams.data("max-size");

    let error_file_tips = this.dropzoneParams.data("error-tips").format(this.transFileSize(maxModelFilesSize))
    let error_space_tips = this.dropzoneParams.data("error-space-tips")
    
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
    const $dropzone = $("div#dataset");
    let vm = this
    const dropzoneUploader = await createDropzone($dropzone[0], {
      url: "/todouploader",
      maxFiles: this.maxFiles,
      maxFilesize: maxModelFilesSize/(1024*1024),
      filesizeBase: 1024,
      parallelUploads: this.maxFiles,
      timeout: 0,
      accept: function(file, done) {
        if (/\s/g.test(file.name)) {
          vm.errorTips = error_file_tips
          $(".maxfilesize.ui.red.message").css("display", "block");
          vm.btnFlag = true;
          done(error_space_tips) 
        } else {
          done()
        }
      },
      addRemoveLinks: true,
      // autoQueue: false,
      autoProcessQueue: false, //自动上传
      dictDefaultMessage: this.dropzoneParams.data("default-message"),
      dictInvalidFileType: this.dropzoneParams.data("invalid-input-type"),
      dictFileTooBig: this.dropzoneParams.data("file-too-big"),
      dictRemoveFile: this.dropzoneParams.data("remove-file"),
      previewTemplate: previewTemplate,
    });
    dropzoneUploader.on("addedfile", (file) => {
      if (file.size / (1024 * 1024) > dropzoneUploader.options.maxFilesize) {
        this.errorTips = error_file_tips
        $(".maxfilesize.ui.red.message").css("display", "block");
        this.btnFlag = true
      } else {
        this.file = file;
      }
      setTimeout(() => this.calculateTotalSize(), 0);
    });
    dropzoneUploader.on("removedfile", (file) => {
      if (this.dropzoneUploader.getRejectedFiles().length === 0) {
        this.btnFlag = false
        $(".maxfilesize.ui.red.message").css("display", "none");
      }
      setTimeout(() => this.calculateTotalSize(), 0);
    });
    dropzoneUploader.on("maxfilesexceeded", function (file) {
      dropzoneUploader.removeFile(file);
      this.errorTips = error_file_tips
      $(".maxfilesize.ui.red.message").css("display", "block");
    });

    this.dropzoneUploader = dropzoneUploader;

    this.getStorageSummary()
  },
  created() {
    if (document.documentElement.attributes["lang"].nodeValue == "en-US") {
      this.i18n = this.$locale.US;
    } else {
      this.i18n = this.$locale.CN;
    }
  },
  watch: {
    allUploadLength(len) {
      if (len === this.uploadFiles.length) {
        setTimeout(() => {
          this.dropzoneUploader.removeAllFiles(true);
          this.btnFlag = false;
          this.$emit("setcluster", this.btnFlag);
        }, 2000);
      }
    },
  },
  methods: {
    transFileSize(srcSize){
      if (null == srcSize || srcSize == '') {
        return '0 Bytes';
      }
      const unitArr = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PB', 'EB', 'ZB', 'YB'];
      srcSize = parseFloat(srcSize);
      const index = Math.floor(Math.log(srcSize) / Math.log(1024));
      let size
      if ((srcSize % Math.pow(1024, index)) !== 0) {
        size = (srcSize / Math.pow(1024, index)).toFixed(2);
      } else {
        size = (srcSize / Math.pow(1024, index))
      }
      return size + ' ' + unitArr[index];
    },
    startUpload() {
      this.uploadFiles = this.dropzoneUploader.getQueuedFiles();
      if (this.uploadFiles.length === 0) {
        return;
      }
      this.resetStatus();
      $(".dz-remove").remove();
      $(".maxfilesize.ui.red.message").css("display", "none");
      this.btnFlag = true;
      this.$emit("setcluster", this.btnFlag);
      this.uploadFiles.forEach((element) => {
        element.datasetId = document
          .getElementById("datasetId")
          .getAttribute("datasetId");
        this.computeMD5(element);
      });
    },
    cancelDataset() {
      location.href = this.repoPath;
      this.dropzoneUploader.removeAllFiles(true);
    },
    resetStatus() {
      this.uploadLength = 0;
      this.allUploadLength = 0;
      // this.allUploadFiles = [];
      this.uploadProgressList = [];
    },
    updateProgress(file, status, progress, infoCode, failedInfo = "") {
      this.uploadProgressList.forEach((item, index) => {
        if (item.name === file.name) {
          this.uploadProgressList[index].status = status;
          this.uploadProgressList[index].progress = progress;
          this.uploadProgressList[index].infoCode = infoCode;
          this.uploadProgressList[index].failedInfo = failedInfo;
        }
      });
    },
    uploadError(file, info) {
      file.previewTemplate.querySelector(".dz-error-mark").style.opacity = 1;
      file.previewTemplate.querySelector(".dz-progress").style.opacity = 0;
      file.previewTemplate.querySelector(".dz-error-message span").innerHTML =
        info;
      file.previewTemplate.querySelector(".dz-error-message").style.display =
        "block";
      file.previewTemplate.querySelector(".dz-details").onmouseover =
        function () {
          file.previewTemplate.querySelector(
            ".dz-error-message"
          ).style.opacity = 1;
        };
      file.previewTemplate.querySelector(".dz-details").onmouseout =
        function () {
          file.previewTemplate.querySelector(
            ".dz-error-message"
          ).style.opacity = 0;
        };
    },
    emitDropzoneSuccess(file) {
      file.status = "success";
      this.dropzoneUploader.emit("success", file);
      this.dropzoneUploader.emit("complete", file);
    },
    emitDropzoneFailed(file) {
      this.status = this.dropzoneParams.data("falied");
      file.status = "error";
      this.dropzoneUploader.emit("error", file);
      // this.dropzoneUploader.emit('complete', file);
    },
    finishUpload(file) {
      file.previewTemplate.querySelector(".dz-success-mark").style.opacity = 1;
      file.previewTemplate.querySelector(".dz-progress").style.opacity = 0;
      if (this.uploadLength === this.uploadFiles.length) {
        setTimeout(() => {
          window.location.href = this.repoPath;
        }, 1000);
      }
    },

    computeMD5(file) {
      const blobSlice =
          File.prototype.slice ||
          File.prototype.mozSlice ||
          File.prototype.webkitSlice,
        chunks = Math.ceil(file.size / chunkSize),
        spark = new SparkMD5.ArrayBuffer(),
        fileReader = new FileReader();
      let currentChunk = 0;
      const time = new Date().getTime();
      this.status = this.dropzoneParams.data("md5-computing");
      this.uploadProgressList.push({
        name: file.name,
        status: this.dropzoneParams.data("md5-computing"),
        progress: 0,
        infoCode: 3,
      });
      file.totalChunkCounts = chunks;
      if (file.size == 0) {
        file.totalChunkCounts = 1;
      }
      loadMd5Next();

      fileReader.onload = (e) => {
        fileLoaded.call(this, e);
      };
      fileReader.onerror = (err) => {
        console.warn("oops, something went wrong.", err);
        file.cancel();
      };

      function fileLoaded(e) {
        spark.append(e.target.result); // Append array buffer
        currentChunk++;
        if (currentChunk < chunks) {
          this.status = `${this.dropzoneParams.data("loading-file")} ${(
            (currentChunk / chunks) *
            100
          ).toFixed(2)}% (${currentChunk}/${chunks})`;
          this.updateProgress(
            file,
            this.dropzoneParams.data("md5-computing"),
            Number(((currentChunk / chunks) *100).toFixed(2)),
            3,
          );
          loadMd5Next();
          return;
        }

        const md5 = spark.end();
        console.log(
          `MD5计算完成：${file.name} \nMD5：${md5} \n分片：${chunks} 大小:${
            file.size
          } 用时：${(new Date().getTime() - time) / 1000} s`
        );
        spark.destroy(); // 释放缓存
        file.uniqueIdentifier = md5; // 将文件md5赋值给文件唯一标识
        file.cmd5 = false; // 取消计算md5状态
        this.updateProgress(
          file,
          this.dropzoneParams.data("md5-computing"),
          100,
          3,
        );
        this.computeMD5Success(file);
      }

      function loadNext() {
        const start = currentChunk * chunkSize;
        const end =
          start + chunkSize >= file.size ? file.size : start + chunkSize;
        fileReader.readAsArrayBuffer(blobSlice.call(file, start, end));
      }

      function loadMd5Next() {
        const start = currentChunk * chunkSize;
        const end =
          start + md5ChunkSize >= file.size ? file.size : start + md5ChunkSize;
        fileReader.readAsArrayBuffer(blobSlice.call(file, start, end));
      }
    },

    async computeMD5Success(md5edFile) {
      const file = await this.getSuccessChunks(md5edFile);
      try {
        if (file.uploadID == "" || file.uuid == "") {
          // 未上传过
          await this.newMultiUpload(file);
          if (file.uploadID != "" && file.uuid != "") {
            file.chunks = "";
            this.multipartUpload(file);
          } else {
            // 失败如何处理
            let info = this.failed;
            this.allUploadLength++;
            this.uploadError(file, info);
            this.updateProgress(file, this.failed, 0, 2);
            // this.allUploadFiles.push({
            //   name: file.name,
            //   status: 2,
            //   info: info,
            // });
            return;
          }
          return;
        }

        if (file.uploaded == "1") {
          // 已上传成功
          // 秒传
          if (file.attachID == "0") {
            // 删除数据集记录，未删除文件
            await addAttachment(file);
          }
          //不同数据集上传同一个文件
          if (file.datasetID != "") {
            if (file.datasetName != "" && file.realName != "") {
              let info = `${this.dropzoneParams.data("file-exists-in-dataset")} ${file.datasetName}`;
              this.uploadError(file, info);
              this.allUploadLength++;
              this.updateProgress(file, this.failed, 0, 1, info);
            }else{
              this.uploadLength++;
              this.allUploadLength++;
              this.updateProgress(file, this.upload_success, 100, 0);
              this.finishUpload(file);
              return;
            }
          }
          console.log("文件已上传完成");
          this.progress = 100;
          this.status = this.dropzoneParams.data("upload-complete");
         
        } else {
          // 断点续传
          this.multipartUpload(file);
        }
      } catch (error) {
        this.emitDropzoneFailed(file);
        console.log(error);
      }

      async function addAttachment(file) {
        return await axios.post(
          "/attachments/add",
          qs.stringify({
            uuid: file.uuid,
            file_name: file.name,
            size: file.size,
            dataset_id: file.datasetId,
            type: this.uploadtype,
            _csrf: csrf,
          })
        );
      }
    },

    async getSuccessChunks(file) {
      const params = {
        params: {
          md5: file.uniqueIdentifier,
          type: this.uploadtype,
          file_name: file.name,
          dataset_id:file.datasetId,
          size: file.size,
          _csrf: csrf,
        },
      };
      try {
        this.updateProgress(file, this.dropzoneParams.data("file-init-status"), 0, 3);
        const response = await axios.get("/attachments/get_chunks", params);
        file.uploadID = response.data.uploadID;
        file.uuid = response.data.uuid;
        file.uploaded = response.data.uploaded;
        file.chunks = response.data.chunks;
        file.attachID = response.data.attachID;
        file.datasetID = response.data.datasetID;
        file.datasetName = response.data.datasetName;
        file.realName = response.data.fileName;
        return file;
      } catch (error) {
        this.emitDropzoneFailed(file);
        console.log("getSuccessChunks catch: ", error);
        return null;
      }
    },

    async newMultiUpload(file) {
      const res = await axios.get("/attachments/new_multipart", {
        params: {
          totalChunkCounts: file.totalChunkCounts,
          md5: file.uniqueIdentifier,
          size: file.size,
          fileType: file.type,
          type: this.uploadtype,
          file_name: file.name,
          _csrf: csrf,
        },
      });
      file.uploadID = res.data.uploadID;
      file.uuid = res.data.uuid;
    },

    multipartUpload(file) {
      const blobSlice =
          File.prototype.slice ||
          File.prototype.mozSlice ||
          File.prototype.webkitSlice,
        chunks = Math.ceil(file.size / chunkSize),
        fileReader = new FileReader(),
        time = new Date().getTime();
      let currentChunk = 0;
      let _this = this;

     function loadNext() {
        const start = currentChunk * chunkSize;
        const end =
          start + chunkSize >= file.size ? file.size : start + chunkSize;
        fileReader.readAsArrayBuffer(blobSlice.call(file, start, end));
      }

      function checkSuccessChunks() {
        const index = successChunks.indexOf((currentChunk + 1).toString());
        if (index == -1) {
          return false;
        }
        return true;
      }

      async function getUploadChunkUrl(currentChunk, partSize) {
        const res = await axios.get("/attachments/get_multipart_url", {
          params: {
            uuid: file.uuid,
            uploadID: file.uploadID,
            size: partSize,
            chunkNumber: currentChunk + 1,
            type: _this.uploadtype,
            file_name: file.name,
            _csrf: csrf,
          },
        });
        urls[currentChunk] = res.data.url;
      }

      async function uploadMinio(url, e) {
        const res = await axios.put(url, e.target.result);
        delete e.target.result;
        etags[currentChunk] = res.headers.etag;
      }

      async function uploadMinioNewMethod(url, e) {
        var xhr = new XMLHttpRequest();
        xhr.open("PUT", url, false);
        if (_this.uploadtype === 0) {
          xhr.setRequestHeader("Content-Type", "text/plain");
          xhr.send(e.target.result);

          var etagValue = xhr.getResponseHeader("etag");
          etags[currentChunk] = etagValue;
        } else if (_this.uploadtype === 1) {
          xhr.setRequestHeader("Content-Type", "");
          xhr.send(e.target.result);
          var etagValue = xhr.getResponseHeader("ETag");
          etags[currentChunk] = etagValue;
        }
      }

      async function updateChunk(currentChunk) {
        await axios.post(
          "/attachments/update_chunk",
          qs.stringify({
            uuid: file.uuid,
            chunkNumber: currentChunk + 1,
            etag: etags[currentChunk],
            _csrf: csrf,
          })
        );
      }
      async function uploadChunk(e,_this) {
        try {
          if (!checkSuccessChunks()) {
            const start = currentChunk * chunkSize;
            const partSize =
              start + chunkSize >= file.size ? file.size - start : chunkSize;
            // 获取分片上传url
            await getUploadChunkUrl(currentChunk, partSize);
            if (urls[currentChunk] != "") {
              // 上传到minio
              //await uploadMinio(urls[currentChunk], e);
              await uploadMinioNewMethod(urls[currentChunk], e);
              if (etags[currentChunk] != "") {
                // 更新数据库：分片上传结果
                //await updateChunk(currentChunk);
                succesLength++
                _this.updateProgress(
                  file,
                  _this.dropzoneParams.data("uploading"),
                  Number(((succesLength / chunks) * 100).toFixed(2)),
                  3
                );
              } else {
                console.log(
                  "上传到minio uploadChunk etags[currentChunk] == ''"
                ); // TODO
              }
            } else {
              console.log("uploadChunk urls[currentChunk] != ''"); // TODO
            }
          }
        } catch (error) {
          // this.emitDropzoneFailed(file);
          //stopFlag = true
          errorChunkLength++
          console.log(error);
        }
      }

      async function completeUpload() {
        return await axios.post(
          "/attachments/complete_multipart",
          qs.stringify({
            uuid: file.uuid,
            uploadID: file.uploadID,
            file_name: file.name,
            size: file.size,
            dataset_id: file.datasetId,
            type: _this.uploadtype,
            _csrf: csrf,
            description: _this.desc,
          })
        );
      }
      async function checkCompleteUpload() {
        if(succesLength===chunks) {
          try {
            await completeUpload();
          } catch (err) {
            let info = _this.failed;
            _this.allUploadLength++;
            _this.uploadError(file, info);
            let completeProgress=Number(((succesLength / chunks) * 100).toFixed(2))
            if(completeProgress==100){
              completeProgress = completeProgress -1
            }
            _this.updateProgress(
              file,
              info,
              completeProgress,
              2
            );
            // this.allUploadFiles.push({
            //   name: file.name,
            //   status: 2,
            //   info: info,
            // });
            if (err) {
              return;
            }
          }

          console.log(
            `文件上传完成：${file.name} \n分片：${chunks} 大小:${
              file.size
            } 用时：${(new Date().getTime() - time) / 1000} s`
          );
          _this.uploadLength++;
          _this.allUploadLength++;
          // this.allUploadFiles.push({
          //   name: file.name,
          //   status: 0,
          //   info: "上传成功",
          // });

          _this.updateProgress(file, _this.upload_success, 100, 0);
          // this.progress = 100;
          // this.status = this.dropzoneParams.data("upload-complete");
          _this.finishUpload(file);
        }
      }
      const successChunks = [];
      let successParts = [];
      successParts = file.chunks.split(",");
      for (let i = 0; i < successParts.length; i++) {
        successChunks[i] = successParts[i].split("-")[0];
      }
      if(successChunks.length>1){
        this.updateProgress(
            file,
            this.dropzoneParams.data("uploading"),
            Number((((successChunks.length-1) / chunks) * 100).toFixed(2)),
            3
          );
      }
      let succesLength = successChunks.length-1
      let errorChunkLength = 0
      const urls = []; // TODO const ?
      const etags = [];
      
      console.log("上传分片...");
      this.status = this.dropzoneParams.data("uploading");
      
      for (currentChunk; currentChunk < chunks; currentChunk++) {
          if(!checkSuccessChunks()){
            loadNext();
            break
          }
      }
      checkCompleteUpload()
      fileReader.onload = async (e) => {
        try {
          await uploadChunk(e,_this);
        } catch (err) {
          console.log(err);
        }
        fileReader.abort();
        currentChunk++;
        if(succesLength<chunks && errorChunkLength>20){
          this.updateProgress(
            file,
            this.failed,
            Number(((succesLength / chunks) * 100).toFixed(2)),
            2
          );
          return
        }
        if (currentChunk < chunks) {
          console.log(
            `第${currentChunk}个分片上传完成, 开始第${
              currentChunk + 1
            }/${chunks}个分片上传`
          );
          for (currentChunk; currentChunk < chunks; currentChunk++) {
            if(!checkSuccessChunks()){
              loadNext();
              break
            }
          }
        }
        if(succesLength<chunks && currentChunk === chunks){
          this.updateProgress(
            file,
            this.failed,
            Number(((succesLength / chunks) * 100).toFixed(2)),
            2
          );
        }
        if(succesLength===chunks){
          checkCompleteUpload()
        }
      };
    },
    async getStorageSummary(){
      const params = {
        params: {
          _csrf: csrf,
        },
      };
      const res = await axios.get('/api/v1/storage/summary', params)
      this.remaining_storage = res.data.remaining_storage
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
    calculateTotalSize() {
      const files = this.dropzoneUploader ? this.dropzoneUploader.getQueuedFiles() : []
      this.totalSize = files.reduce((sum, file) => sum + file.size, 0)
      this.isStorageExceeded = this.totalSize > this.remaining_storage
    },
  },
};
</script>

<style scoped lang="less">
.dropzone-wrapper {
  margin: 0;
}
.ui .dropzone {
  border: 2px dashed #0087f5;
  box-shadow: none !important;
  padding: 0;
  min-height: 5rem;
  border-radius: 4px;
}
.dataset .dataset-files #dataset .dz-preview.dz-file-preview,
.dataset .dataset-files #dataset .dz-preview.dz-processing {
  display: flex;
  align-items: center;
}
.dataset-files .data-storage-limit{
  position: absolute;
  top: -40px;
  display: flex;
  flex-wrap: wrap;
  width: 100%;
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
      align-items: center;
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
.dataset .dataset-files #dataset .dz-preview {
  border-bottom: 1px solid #dadce0;
  min-height: 0;
}
.upload-info {
  margin-top: 1em;
  margin-bottom: 3em;
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
  width: 200px;
  margin-right: 1rem;
}
.datast-upload-progress .dataset-progress {
  flex: 1;
}
.datast-upload-progress .dataset-status {
  margin-left: 1rem;
  min-width: 102px;
}
.datast-upload-progress .dataset-status .status-flex {
  display: flex;
  align-items: center;
}
/deep/ .el-progress-bar__inner {
  background-color: #21ba45;
}
</style>
