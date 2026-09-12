<template>
  <div class="item-container">
    <div class="content">
      <div class="files-info" v-loading="loading">
        <div class="top" v-if="resultStatus == 0">
          <div>
            <div class="title files-path-c">
              <div class="file-path" v-for="(item, index) in filePath" :key="index">
                <span v-if="index == filePath.length - 1" class="path-name">{{ item.label }}</span>
                <a v-if="index != filePath.length - 1" class="path-name canback" @click="goBackDir(item)">{{
                  item.label
                  }}</a>
                <span style="color:rgba(0,0,0,.4);" class="divider"> / </span>
              </div>
            </div>
          </div>
          <div class="right-btn-c">
            <a class="op-btn" v-if="filesList.length" :href="downloadAllUrl"
              :class="!canDownload ? 'disabled-download' : ''">
              <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 32 32" width="20" height="20"
                fill="none">
                <defs></defs>
                <g>
                  <path
                    d="M24 8h-4v2h2.8l4 8h-8.8v4h-4v-4h-8.8l4-8h2.8v-2h-4l-6 12v10h28v-10l-6-12zM28 28h-24v-8h8v4h8v-4h8v8z">
                  </path>
                  <path d="M11.2 12.8l4.8 4.6 4.8-4.6-1.6-1.6-2.2 2.4v-9.6h-2v9.6l-2.2-2.4z"></path>
                </g>
              </svg>
              <span>{{ $t('cloudbrainObj.allResultDownload') }}</span>
            </a>
            <a class="op-btn" href="javascript:;" @click="refresh">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="20" height="20">
                <defs></defs>
                <g>
                  <path fill="#202565"
                    d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z">
                  </path>
                </g>
              </svg>
              <span>{{ $t('cloudbrainObj.refresh') }}</span>
            </a>
          </div>
        </div>
        <div class="table-container" v-if="resultStatus == 0">
          <el-table ref="tableRef" :data="filesList" row-key="sn" height="100%" tyle="min-width:100%">
            <el-table-column column-key="FileName" prop="FileName" sortable
              :sort-method="(a, b) => a.FileName.toLocaleLowerCase().localeCompare(b.FileName.toLocaleLowerCase())"
              :label="$t('modelManage.fileName')" align="left" header-align="left" min-width="160">
              <template slot-scope="scope">
                <div class="tbl-file-name">
                  <a v-if="scope.row.IsDir" @click="goNextDir(scope.row)" href="javascript:;">
                    <div class="fitted" :title="scope.row.FileName">
                      <i class="icon folder" width="16" height="16" aria-hidden="true"></i>
                      <span>{{ scope.row.FileName }}</span>
                    </div>
                  </a>
                  <a v-else :class="!canDownload ? 'disabled-download' : ''" :href="scope.row.downloadUrl">
                    <div class="fitted" :title="scope.row.FileName">
                      <i class="icon file" width="16" height="16" aria-hidden="true"></i>
                      <span>{{ scope.row.FileName }}</span>
                    </div>
                  </a>
                </div>
              </template>
            </el-table-column>
            <el-table-column column-key="SizeShow" prop="SizeShow" sortable :sort-method="(a, b) => a.Size - b.Size"
              :label="$t('modelManage.fileSize')" align="left" header-align="left" width="200">
            </el-table-column>
            <el-table-column column-key="ModTime" prop="ModTime" sortable
              :sort-method="(a, b) => a.ModTimeNum - b.ModTimeNum" :label="$t('modelManage.updateTime')" align="center"
              header-align="center" width="200">
            </el-table-column>
          </el-table>
          <div class="max-count-tips"><i class="el-icon-warning-outline"></i>{{
            this.$t('cloudbrainObj.downloadDisplayMaxCountTips', { count: 100, size: size }) }}</div>
        </div>
        <div v-if="resultStatus != 0">
          <div class="status-tips" v-if="isTaskTerminal === false">
            <div>
              <i class="ri-time-line"></i>
              <span>{{ this.$t('cloudbrainObj.task_not_finished') }}</span>
            </div>
          </div>
          <div class="status-tips" v-else>
            <div v-if="resultStatus == 1">
              <svg xmlns="http://www.w3.org/2000/svg" style="margin-right:5px" viewBox="0 0 24 24" width="14"
                height="14" class="rotating" fill="#101010">
                <path
                  d="M6 4H4V2H20V4H18V6C18 7.61543 17.1838 8.91468 16.1561 9.97667C15.4532 10.703 14.598 11.372 13.7309 12C14.598 12.628 15.4532 13.297 16.1561 14.0233C17.1838 15.0853 18 16.3846 18 18V20H20V22H4V20H6V18C6 16.3846 6.81616 15.0853 7.8439 14.0233C8.54682 13.297 9.40202 12.628 10.2691 12C9.40202 11.372 8.54682 10.703 7.8439 9.97667C6.81616 8.91468 6 7.61543 6 6V4ZM8 4V6C8 6.88457 8.43384 7.71032 9.2811 8.58583C10.008 9.33699 10.9548 10.0398 12 10.7781C13.0452 10.0398 13.992 9.33699 14.7189 8.58583C15.5662 7.71032 16 6.88457 16 6V4H8ZM12 13.2219C10.9548 13.9602 10.008 14.663 9.2811 15.4142C8.43384 16.2897 8 17.1154 8 18V20H16V18C16 17.1154 15.5662 16.2897 14.7189 15.4142C13.992 14.663 13.0452 13.9602 12 13.2219Z">
                </path>
              </svg>
              <span>{{ this.$t('cloudbrainObj.file_sync_ing') }}</span>
            </div>
            <div v-if="resultStatus == 2">
              <i class="ri-alert-line"></i>
              <span>{{ this.$t('cloudbrainObj.file_sync_fail') }}</span>
              <a v-if="canReschedule" class="retry" href="javascript:void(0)" @click="retry">{{
                this.$t('cloudbrainObj.retrieve_results') }}</a>
            </div>
            <div v-if="resultStatus == 3">
              <i class="ri-time-line"></i>
              <span>{{ this.$t('cloudbrainObj.file_sync_wait') }}</span>
            </div>
            <div v-if="resultStatus == 4">
              <i class="ri-emotion-unhappy-line"></i>
              <span>{{ this.$t('cloudbrainObj.no_file_to_download') }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {
  getAiTaskOutputResult, setAiTaskOutputReschedule,
  getDownLoadAiTaskResultFileUrl, getDownLoadAiTaskResultFileAllUrl
} from '~/apis/modules/cloudbrain';
import { transFileSize } from '~/utils';

export default {
  name: 'ResultDownload',
  props: {
    configs: { type: Object, default: () => { return {} } },
    data: { type: Object, default: () => { return {} } },
  },
  data() {
    return {
      filesList: [],
      filePath: [],
      loading: false,
      isTaskTerminal: null,
      canDownload: false,
      canReschedule: false,
      resultStatus: 0,
      resultPath: '',
      downloadAllUrl: '',
      size: 0
    };
  },
  methods: {
    getDirFiles(dir) {
      dir = dir.length ? dir.slice(1) : '';
      const task = this.data.task;
      this.loading = true;
      getAiTaskOutputResult({
        repoOwnerName: task.repoOwnerName,
        repoName: task.repoName,
        id: task.id,
        parent_dir: dir,
      }).then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 0 && res.data && res.data.output) {
          const data = res.data.output;
          this.isTaskTerminal = data.is_task_terminal;
          this.canDownload = data.can_download;
          this.canReschedule = data.can_reschedule;
          this.resultStatus = data.status;
          this.resultPath = data.path;
          this.size = data.output_size_limit
          this.downloadAllUrl = getDownLoadAiTaskResultFileAllUrl({
            repoOwnerName: task.repoOwnerName,
            repoName: task.repoName,
            id: task.id,
          });
          if (this.resultStatus == 0) { // 成功 0
            const list = data.file_list || [];
            list.forEach(item => {
              item.SizeShow = item.IsDir ? '' : transFileSize(item.Size);
              item.ModTimeNum = new Date(item.ModTime).getTime();
              item.downloadUrl = getDownLoadAiTaskResultFileUrl({
                repoOwnerName: task.repoOwnerName,
                repoName: task.repoName,
                id: task.id,
                parent_dir: dir,
                file_name: item.FileName,
              });
            });
            list.sort((a, b) => a.FileName.localeCompare(b.FileName));
            list.sort((a, b) => b.ModTimeNum - a.ModTimeNum);
            this.filesList = list;
            this.$refs['tableRef']?.clearSort();
          }
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    goNextDir(item) {
      this.filePath.push({
        label: item.FileName,
        path: item.FileName
      });
      const dir = this.filePath.map((item) => item.path).join('/');
      this.getDirFiles(dir);
    },
    goBackDir(item) {
      const index = this.filePath.findIndex(pth => item === pth);
      this.filePath = this.filePath.slice(0, index + 1);
      const dir = this.filePath.map((item) => item.path).join('/');
      this.getDirFiles(dir);
    },
    retry() {
      const task = this.data.task;
      const experienceParams = {}
      setAiTaskOutputReschedule({
        id: task.id
      }).then(res => {
        res = res.data;
        if (res.code == 0) {
          this.refresh();
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      }).catch(err => {
        console.log(err);
      })
    },
    refresh() {
      // console.log(this.configs);
      // console.log(this.data);
      const task = this.data.task;
      this.filesList = [];
      const version = task.current_version_name;
      if (version) {
        this.filePath = [{ label: version, path: `` }];
        this.getDirFiles(``);
      } else {
        this.filePath = [{ label: 'result', path: '' }];
        this.getDirFiles('');
      }
    }
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
.item-container {
  height: 100%;
}

.content {
  padding: 20px 30px 0px 30px;
  height: 100%;

  .files-info {
    height: 100%;
    display: flex;
    flex-direction: column;

    .top {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 10px;
      flex-wrap: wrap;
    }

    .files-path-c {
      height: 20px;

      .file-path {
        margin-right: 6px;
        float: left;
        font-size: 14px;
        font-weight: 550;
        color: #101010;

        .path-name {
          &.canback {
            color: #4183c4;
          }
        }
      }
    }

    .right-btn-c {
      display: flex;
      align-items: center;
      justify-content: flex-end;

      .op-btn {
        display: flex;
        align-items: center;
        background-color: rgba(16, 16, 16, 0.1);
        color: rgba(32, 37, 101, 0.9);
        height: 30px;
        padding: 0 10px;

        &:first-child {
          margin-right: 1px;
          border-radius: 6px 0px 0px 6px;
        }

        &:last-child {
          border-radius: 0px 6px 6px 0px;
        }

        .fill:not([stroke]) {
          fill: rgba(32, 37, 101, 0.9);
        }

        svg {
          margin-right: 4px;
        }
      }
    }

    .table-container {
      flex: 1;
      height: 0;
      display: flex;
      flex-direction: column;

      /deep/ .el-table__header {
        th {
          background: rgb(245, 245, 246);
          color: rgb(16, 16, 16);
          font-weight: 400;
          font-size: 14px;
        }
      }

      /deep/ .el-table__body {
        td {
          color: rgb(16, 16, 16);
          font-weight: 400;
          font-size: 14px;
        }
      }

      .tbl-file-name {
        height: 32px;
        display: flex;
        align-items: center;
        overflow: hidden;
        font-size: 16px;
        font-weight: 500;
        position: relative;

        a {
          max-width: 100%;

          .fitted {
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
            max-width: 100%;
          }
        }
      }
    }

    .max-count-tips {
      display: flex;
      align-items: center;
      justify-content: flex-end;;
      color: #f2711c;
      font-size: 14px;
      margin-top: 10px;
      margin-bottom: 10px;

      i {
        margin-right: 5px;
      }
    }
  }
}

@media screen and (max-width: 767px) {
  .content {
    padding: 20px 10px 0px 10px;
  }
}

.status-tips {
  height: 200px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 14px;
  color: rgb(16, 16, 16);

  div {
    display: flex;
    justify-content: center;
    align-items: center;

    i {
      margin-right: 5px;
    }
  }

  .retry {
    text-decoration: underline;
    margin-left: 0.5rem
  }
}

.disabled-download {
  cursor: default;
  pointer-events: none;
  color: rgba(0, 0, 0, .6) !important;
  opacity: .45 !important;
}

@keyframes rotation {
  from {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(360deg);
  }
}

.rotating {
  animation: rotation 4s linear infinite;
}

/deep/ .el-table .el-table__body-wrapper::-webkit-scrollbar {
  -webkit-appearance: none;
  width: 7px;
  height: 7px
}

/deep/ .el-table .el-table__body-wrapper::-webkit-scrollbar-track {
  background: transparent;
  border-radius: 7px !important;
}

/deep/ .el-table .el-table__body-wrapper::-webkit-scrollbar-thumb {
  background: rgb(210, 210, 216) !important;
  border-radius: 7px !important;
}

/deep/ .el-table .el-table__body-wrapper::-webkit-scrollbar-thumb:hover {
  background: rgba(210, 210, 216, 0.8) !important;
  border-radius: 7px !important;
}

/deep/ .el-table .el-table__body-wrapper::-webkit-scrollbar-thumb:active {
  background: rgba(210, 210, 216, 0.8) !important;
  border-radius: 7px !important;
}
</style>
