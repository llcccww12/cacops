<template>
  <div class="folder-upload-select-c">
    <div class="folder-upload-select">
      <div class="op-c">
        <el-button size="mini" icon="el-icon-circle-plus" :disabled="uploading" @click="addFolder">{{
          $t('modelManage.addUploadFolder') }}</el-button>
        <div class="op-right">
          <div class="summary">{{ $t('modelManage.fileCountAndSize', { count: fileList.length, size: allFilesSize }) }}
          </div>
          <el-button size="mini" type="text" style="color:red" :disabled="uploading" @click="remove()">{{
            $t('modelManage.clearAll') }}</el-button>
        </div>
      </div>
      <div class="err-tips-c">
        <div class="err-tips ui red message" v-show="errorState" v-html="errorInfo"></div>
      </div>
      <div class="file-list-c" @dragover="handleDragover" @drop="handleDrop">
        <div class="file-item" v-for="(file, index) in fileList" :key="index" @click.stop.prevent="">
          <div class="file-name" :style="{color: file.sizeStatus == 'error' ? 'red' : ''}">{{ file.fullname }}</div>
          <div class="file-size">{{ file.file_size }}</div>
          <div class="file-status">
            <el-tooltip v-if="file.status == 'error'" effect="dark" :content="file.errTips" placement="top">
              <i class="el-icon-warning-outline" style="color:red"></i>
            </el-tooltip>
            <i class="el-icon-circle-check" v-if="file.status == 'success'" style="color:#21ba45"></i>
          </div>
          <div class="file-op"><i class="el-icon-circle-close" :disabled="uploading"
              @click.stop.prevent="remove(file)"></i></div>
        </div>
        <div class="empty-place-holder" v-if="!fileList.length" @click="addFolder"
          v-html="$t('modelManage.folderUploadPlaceholder')">
        </div>
      </div>
      <input type="file" ref="filepicker" @change="folderSelectChange" webkitdirectory directory multiple
        style="display:none" />
    </div>
    <div class="tips"
      v-html="getDefaultErrTxt()"></div>
  </div>
</template>

<script>
import { transFileSize, uuidv4 } from '~/utils';
export default {
  name: "FolderUploadSelect",
  props: {
    uploading: { type: Boolean, default: false },
    maxFilesCount: { type: Number, default: 100 },
    maxFilesSize: { type: Number, default: 200 * 1024 * 1024 * 1024 },
  },
  data() {
    return {
      fileList: [],
      errorState: false,
      errorInfo: '',
      errorList: [],
    };
  },
  computed: {
    allFilesSize() {
      const allFilesSize = this.fileList.reduce((acc, item, index) => acc + item.size, 0);
      return transFileSize(allFilesSize);
    },
  },
  methods: {
    addFolder() {
      this.$refs['filepicker'].value = '';
      this.$refs['filepicker'].click();
    },
    folderSelectChange(evt) {
      const files = evt.target.files || [];
      for (let i = 0, iLen = files.length; i < iLen; i++) {
        const file = files[i];
        file._webkitRelativePath = file.webkitRelativePath;
      }
      this.handleFiles(files);
    },
    handleFiles(files) {
      for (let i = 0, iLen = files.length; i < iLen; i++) {
        const file = files[i];
        if (this.fileList.findIndex(item => item._webkitRelativePath == file._webkitRelativePath) >= 0) continue;
        file.fullname = file._webkitRelativePath ? file._webkitRelativePath : file.name;
        file.file_size = transFileSize(file.size);
        file.upload = { uuid: uuidv4() };
        this.fileList.push(file);
      }
      this.checkFiles();
      this.$emit('folderSelectChange', this.fileList);
    },
    handleDragover(evt) {
      evt.preventDefault();
    },
    handleDrop(evt) {
      evt.preventDefault();
      const filesList = [];
      const dataTransferItemList = evt.dataTransfer.items;
      for (const dataTransferItem of dataTransferItemList) {
        const fileEntry = dataTransferItem.webkitGetAsEntry();
        this.handleDropedFileEntry(fileEntry, filesList);
      }
    },
    handleDropedFileEntry(fileEntry) {
      const self = this;
      if (fileEntry.isFile) {
        fileEntry.file(function (file) {
          file._webkitRelativePath = fileEntry.fullPath.slice(1);
          self.handleFiles([file]);
          return;
        })
      } else {
        const dirReader = fileEntry.createReader()
        dirReader.readEntries(function (entries) {
          for (let i = 0; i < entries.length; i++) {
            self.handleDropedFileEntry(entries[i]);
          }
        })
      }
    },
    remove(file) {
      if (file) {
        const index = this.fileList.findIndex(item => item._webkitRelativePath === file._webkitRelativePath);
        this.fileList.splice(index, 1);
      } else {
        this.fileList.splice(0);
      }
      this.checkFiles();
      this.$emit('folderSelectChange', this.fileList);
    },
    showErrInfo(state, info) {
      this.errorState = state;
      this.errorInfo = info;
    },
    getDefaultErrTxt() {
      return this.$t('modelManage.modelFileUploadErrTips', { maxCount: this.maxFilesCount, size: transFileSize(this.maxFilesSize) ,url:'https://openi.pcl.ac.cn/docs/index.html#/model/sdk'});
    },
    checkFiles() {
      let flag = false
      this.fileList.forEach(item => {
        if (item.size > this.maxFilesSize || item.name.length > 128) {
          flag = true
          item.sizeStatus = 'error'
        }
      });
      if (this.fileList.length > this.maxFilesCount || flag) {
        this.showErrInfo(true, this.getDefaultErrTxt());
        return false;
      }
      this.showErrInfo(false, '');
      return true;
    },
    getAcceptedFiles() {
      return this.fileList;
    },
  },
  beforeMount() { },
};
</script>

<style scoped lang="less">
.folder-upload-select-c {
  position: relative;

  .folder-upload-select {
    border: 2px dashed #0087f5;
    border-radius: 4px;
    padding: 4px;

    .op-c {
      margin: 4px;
      display: flex;
      align-items: center;
      justify-content: space-between;
      flex-wrap: wrap;
      .op-right {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        .summary {
          margin-right: 12px;
          font-size: 12px;
          color: #606266;
        }
      }
    }

    .err-tips-c {
      margin: 8px 4px;
    }

    .file-list-c {
      min-height: 130px;
      max-height: 300px;
      overflow: auto;
      border: 1px solid burlywood;
      margin: 4px;

      .file-item {
        color: #606266;
        font-size: 14px;
        display: flex;
        align-items: center;
        margin: 0 4px;
        padding: 4px 0;
        border-bottom: 1px dashed burlywood;

        .file-name {
          flex: 1;
          width: 0;
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
          padding-left: 5px;
          padding-right: 10px;
        }

        .file-size {
          width: 100px;
          margin: 0 10px;
        }

        .file-status {
          width: 30px;
          display: flex;
          align-items: center;
          justify-content: center;
        }

        .file-op {
          width: 30px;
          display: flex;
          align-items: center;
          justify-content: center;

          i {
            cursor: pointer;
            color: #606266;
          }
        }
      }

      .empty-place-holder {
        color: rgba(0, 0, 0, .6);
        padding: 54px 0 54px 0;
        text-align: center;
        cursor: pointer;
      }
    }
  }

  .tips {
    margin-top: 10px;
    font-size: 13px;
    color: rgba(0, 0, 0, 0.6);
  }
}
</style>
