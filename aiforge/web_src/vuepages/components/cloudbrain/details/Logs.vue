<template>
  <div class="item-container">
    <div class="log-container">
      <div class="op-tool-c">
        <div>
          <el-select v-if="configs.multiNodes && multiNodesData.length > 1" v-model="common.nodeSel" @change="changeNode">
            <el-option v-for="(item, index) in multiNodesData" :key="item.id" :value="index"
              :label="`${$t('cloudbrainObj.computeNode')} ${index + 1}`"></el-option>
          </el-select>
        </div>
        <div class="op-btn-c">
          <a class="op-btn" :href="common.downloadUrl" :class="downloading || !canLogDownload ? 'disabled' : ''">
            <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 32 32" width="20" height="20" fill="none"><defs></defs><g><path d="M24 8h-4v2h2.8l4 8h-8.8v4h-4v-4h-8.8l4-8h2.8v-2h-4l-6 12v10h28v-10l-6-12zM28 28h-24v-8h8v4h8v-4h8v8z"></path><path d="M11.2 12.8l4.8 4.6 4.8-4.6-1.6-1.6-2.2 2.4v-9.6h-2v9.6l-2.2-2.4z"></path></g></svg>
            <span>{{ $t('modelManage.download') }}</span>
          </a>
          <a class="op-btn" href="javascript:;" @click="dialogShow = true">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" width="20" height="20" fill="none"><defs></defs><g>
              <rect id="全屏-imiijmsard5yz2g-V3pOBaFoDeaT9o" width="20.000000" height="20.000000" x="0.000000" y="0.000000"></rect>
              <g id="组合 7-imiijmsard5yz2g-V3pOBaFoDeaT9o">
                <path id="合并-imiijmsard5yz2g-V3pOBaFoDeaT9o" d="M11.9333 7.21802L12.7819 8.06655L16.7998 4.04861L16.7998 8.00002L17.9998 8.00002L17.9998 2.84861L17.9999 2.84853L17.9998 2.84845L17.9998 2.00002L17.1514 2.00002L17.1514 2L17.1514 2.00002L11.9998 2.00002L11.9998 3.20002L15.9514 3.20002L11.9333 7.21802Z" fill="rgba(32,37,101,0.9)" fill-rule="evenodd"></path>
                <path id="合并-imiijmsard5yz2g-V3pOBaFoDeaT9o" d="M0 5.21802L0.848528 6.06655L4.86647 2.04861L4.86647 6.00002L6.06647 6.00002L6.06647 0.84861L6.06655 0.848528L6.06647 0.848446L6.06647 1.74046e-05L5.21804 1.74046e-05L5.21802 0L5.218 1.74046e-05L0.0664654 1.74046e-05L0.0664654 1.20002L4.018 1.20002L0 5.21802Z" fill="rgba(32,37,101,0.9)" fill-rule="evenodd" transform="matrix(-1,0,0,-1,8.06665,18)"></path>
                <path id="合并-imiijmsard5yz2g-V3pOBaFoDeaT9o" d="M6 1.2L6 0L0 0L0 1.2L4.8 1.2L4.8 6L6 6L6 1.2Z" fill="rgba(32,37,101,0.9)" fill-rule="evenodd" transform="matrix(0,1,-1,0,18,12)"></path>
                <path id="合并-imiijmsard5yz2g-V3pOBaFoDeaT9o" d="M6 1.2L6 0L0 0L0 1.2L4.8 1.2L4.8 6L6 6L6 1.2Z" fill="rgba(32,37,101,0.9)" fill-rule="evenodd" transform="matrix(-0,-1,1,0,2,8)"></path>
              </g>
            </g></svg>
            <span>{{ $t('cloudbrainObj.viewFullScreen') }}</span>
          </a>
          <a class="op-btn" href="javascript:;" @click="refresh">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="20" height="20"><defs></defs><g><path fill="#202565" d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"></path></g></svg>
            <span>{{ $t('cloudbrainObj.refresh') }}</span>
          </a>
        </div>
      </div>
      <div class="log-content-c" v-loading="common.loading">
        <div class="log-content" ref="commonLogContentRef" @scroll="scrollHandler" v-html="common.content">
        </div>
        <div class="right-btn-c">
          <div class="icon-wrap" style="margin-top:8px;" @click="goTop" :title="$t('cloudbrainObj.scrollToTop')">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="16" height="16" fill="none"><defs></defs><g><path d="M24.0083 14.1005V42" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path d="M12 26L24 14L36 26" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path d="M12 6H36" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path></g></svg>
          </div>
          <div class="icon-wrap" style="margin-bottom:8px;" @click="goBottom" :title="$t('cloudbrainObj.scrollToBottom')">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="16" height="16"><path d="M24.0083 33.8995V6" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M36 22L24 34L12 22" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M36 42H12" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
          </div>
        </div>
      </div>
    </div>
    <el-dialog class="log-fullscreen-dlg" :visible.sync="dialogShow" :fullscreen="true" :modal="true"
      :modal-append-to-body="true" :lock-scroll="true" @open="open" @opened="opened" @closed="closed">
      <div class="log-fullscreen-container">
        <div class="log-fullscreen-header-c">
          <div class="title">{{ $t('cloudbrainObj.logFile') }}</div>
          <div class="op-tool-c">
            <div style="margin-right:20px">
              <el-select v-if="configs.multiNodes && multiNodesData.length > 1" v-model="fullscreen.nodeSel"
                @change="changeNode">
                <el-option v-for="(item, index) in multiNodesData" :key="item.id" :value="index"
                  :label="`${$t('cloudbrainObj.computeNode')} ${index + 1}`"></el-option>
              </el-select>
            </div>
            <div class="op-btn-c">
              <a class="op-btn" :href="fullscreen.downloadUrl" :class="downloading || !canLogDownload ? 'disabled' : ''">
                <i class="ri-download-cloud-2-line"></i>
                <span>{{ $t('modelManage.download') }}</span>
              </a>
              <a class="op-btn" href="javascript:;" @click="dialogShow = false">
                <i class="ri-fullscreen-exit-fill"></i>
                <span>{{ $t('cloudbrainObj.exitFullScreen') }}</span>
              </a>
              <a class="op-btn" href="javascript:;" @click="refresh">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="20" height="20"><defs></defs><g><path fill="#202565" d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"></path></g></svg>
                <span>{{ $t('cloudbrainObj.refresh') }}</span>
              </a>
            </div>
          </div>
        </div>
        <div class="log-fullscreen-content-c" v-loading="fullscreen.loading">
          <div class="log-content" ref="fullscreenLogContentRef" @scroll="scrollHandler" v-html="fullscreen.content">
          </div>
          <div class="right-btn-c">
            <div class="icon-wrap" style="margin-top:8px;" @click="goTop" :title="$t('cloudbrainObj.scrollToTop')">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="16" height="16" fill="none"><defs></defs><g><path d="M24.0083 14.1005V42" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path d="M12 26L24 14L36 26" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path><path d="M12 6H36" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"></path></g></svg>
            </div>
            <div class="icon-wrap" style="margin-bottom:8px;" @click="goBottom" :title="$t('cloudbrainObj.scrollToBottom')">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="16" height="16"><path d="M24.0083 33.8995V6" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M36 22L24 34L12 22" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M36 42H12" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
            </div>
            <!--<i class="icon-to-top" @click="goTop" :title="$t('cloudbrainObj.scrollToTop')"></i>
            <i class="icon-to-bottom" @click="goBottom" :title="$t('cloudbrainObj.scrollToBottom')"></i>-->
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getAiTaskNodeInfo, getAiTaskLogs, getAiTaskLogsDownloadUrl } from '~/apis/modules/cloudbrain';

export default {
  name: 'Logs',
  props: {
    configs: { type: Object, default: () => { return {} } },
    data: { type: Object, default: () => { return {} } },
  },
  data() {
    return {
      common: {
        nodeSel: 0,
        content: '',
        lines: 60,
        order: 'up',
        startLine: '',
        endLine: '',
        loading: false,
        scrollTop: 0,
        downloadUrl: '',
      },
      fullscreen: {
        nodeSel: 0,
        content: '',
        lines: 100,
        order: 'up',
        startLine: '',
        endLine: '',
        loading: false,
        scrollTop: 0,
        downloadUrl: '',
      },
      multiNodesData: [],
      dialogShow: false,
      canLogDownload: 0,
      downloading: false,
    };
  },
  methods: {
    escapeHTML(str) {
      return str.toString().replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;").replace(/'/g, "&apos;");
    },
    getLogs(callback) {
      const key = this.dialogShow ? 'fullscreen' : 'common';
      this[key].loading = true;
      const task = this.data.task;
      const params = {
        id: task.id,
        base_line: this[key].order == 'up' ? this[key].startLine : this[key].endLine,
        lines: this[key].lines,
        order: this[key].order,
        node_id: this.configs.multiNodes ? this.multiNodesData[this[key].nodeSel]?.id : undefined,
        log_file_name: this.configs.multiNodes ? this.multiNodesData[this[key].nodeSel]?.log_file_name : undefined,
      };
      getAiTaskLogs(params).then(res => {
        this[key].loading = false;
        res = res.data;
        if (res.code == 0) {
          res = res.data || {};
          if (this.canLogDownload === 0) {
            this.canLogDownload = res.can_log_download;
          }
          if (res.lines > 0) {
            if (this[key].order == 'up') {
              this[key].content = `<pre>${this.escapeHTML(res.content)}</pre>` + this[key].content;
              this[key].startLine = res.start_line;
              if (this[key].endLine) {
                this.stayScrollPos();
              }
              if (!this[key].endLine) {
                this[key].endLine = res.end_line;
              }
            } else if (this[key].order == 'down') {
              this[key].content += `<pre>${this.escapeHTML(res.content)}</pre>`;
              this[key].endLine = res.end_line;
              if (!this[key].startLine) {
                this[key].startLine = res.start_line;
              }
            }
          } else {
            if (this.configs.noScroll) {
              this[key].content = `<pre>${this.escapeHTML(res.content)}</pre>`;
            } else {
              let msg = '';
              if (!this[key].content) { // log content is empty
                msg = this[key].order == 'down' ? this.$t('cloudbrainObj.scrolledToTopTip')
                  : this.$t('cloudbrainObj.scrolledToBottomTip');
              } else {
                msg = this[key].order == 'up' ? this.$t('cloudbrainObj.scrolledToTopTip')
                  : this.$t('cloudbrainObj.scrolledToBottomTip');
              }
              this.$message({
                type: 'info',
                message: msg,
              });
            }
          }
          callback && callback();
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      }).catch(err => {
        this[key].loading = false;
        console.log(err);
      });
    },
    scrollHandler(evt) {
      if (this.configs.noScroll) return;
      const key = this.dialogShow ? 'fullscreen' : 'common';
      if (this[key].loading) return;
      const logEle = this.$refs[key + 'LogContentRef'];
      if (logEle) {
        const scrollHeight = logEle.scrollHeight;
        const clientHeight = logEle.clientHeight;
        const scrollTop = logEle.scrollTop;
        if (scrollTop != this[key].scrollTop) {
          if (scrollTop == 0) {
            this[key].order = 'up';
            this.getLogs();
          } else if (scrollTop + clientHeight >= scrollHeight - 1) {
            this[key].order = 'down';
            this.getLogs();
          }
        }
        this[key].scrollTop = scrollTop;
      }
    },
    goTop() {
      const key = this.dialogShow ? 'fullscreen' : 'common';
      this[key].content = '';
      this[key].startLine = '';
      this[key].endLine = '';
      this[key].order = 'down';
      this.getLogs(() => {
        this.stayScrollPos();
        this.$nextTick(() => {
          const logEle = this.$refs[key + 'LogContentRef'];
          if (logEle) {
            logEle.scrollTo({
              top: 0,
              behavior: 'smooth',
            });
          }
        });
      });
    },
    goBottom() {
      this.refresh();
    },
    stayScrollPos() {
      const key = this.dialogShow ? 'fullscreen' : 'common';
      const logEle = this.$refs[key + 'LogContentRef'];
      if (logEle) {
        const scrollHeight = logEle.scrollHeight;
        const scrollTop = logEle.scrollTop;
        this.$nextTick(() => {
          const scrollHeightNew = logEle.scrollHeight;
          const scrollTopNew = logEle.scrollTop;
          logEle.scrollTo({
            top: scrollTop + (scrollHeightNew - scrollHeight),
            behavior: 'instant',
          });
        });
      }
    },
    scrollBottomAnimation() {
      const key = this.dialogShow ? 'fullscreen' : 'common';
      const logEle = this.$refs[key + 'LogContentRef'];
      if (logEle) {
        const scrollHeight = logEle.scrollHeight;
        const clientHeight = logEle.clientHeight;
        const scrollTop = logEle.scrollTop;
        logEle.scrollTo({
          top: scrollHeight - clientHeight,
          behavior: 'smooth',
        });
      }
    },
    getFirstLogs() {
      this.getLogs(() => {
        this.$nextTick(() => {
          this.scrollBottomAnimation();
        });
      });
      const key = this.dialogShow ? 'fullscreen' : 'common';
      const task = this.data.task;
      
      this[key].downloadUrl = getAiTaskLogsDownloadUrl({
        id: task.id,
        node_id: this.configs.multiNodes ? this.multiNodesData[this[key].nodeSel]?.id : undefined,
        log_file_name: this.configs.multiNodes ? this.multiNodesData[this[key].nodeSel]?.log_file_name : undefined,
      });
      console.log("this[key].downloadUrl",this[key].downloadUrl)
    },
    refresh() {
      const key = this.dialogShow ? 'fullscreen' : 'common';
      this[key].nodeSel = 0;
      this[key].content = '';
      this[key].startLine = '';
      this[key].endLine = '';
      this[key].order = 'up';
      this[key].downloadUrl = '';
      if (this.configs.multiNodes) {
        const task = this.data.task;
        getAiTaskNodeInfo({
          repoOwnerName: task.repoOwnerName,
          repoName: task.repoName,
          id: task.id,
        }).then(res => {
          res = res.data;
          if (res && res.code == 0) {
            this.multiNodesData = res.data?.nodes || [];
          }
          this.getFirstLogs();
        }).catch(err => {
          console.log(err);
        });
      } else {
        this.getFirstLogs();
      }
    },
    changeNode() {
      const key = this.dialogShow ? 'fullscreen' : 'common';
      this[key].content = '';
      this[key].startLine = '';
      this[key].endLine = '';
      this[key].order = 'up';
      this[key].downloadUrl = '';
      this.getFirstLogs();
    },
    open() {
      const key = 'fullscreen';
      this[key].nodeSel = 0;
      this[key].content = '';
      this[key].startLine = '';
      this[key].endLine = '';
      this[key].order = 'up';
      this[key].downloadUrl = '';
    },
    opened() {
      this.refresh();
    },
    closed() {
      const key = 'fullscreen';
      this[key].nodeSel = 0;
      this[key].content = '';
      this[key].startLine = '';
      this[key].endLine = '';
      this[key].order = 'up';
      this[key].downloadUrl = '';
    },
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
.item-container{
  height: 100%;
  .log-container {
    margin-bottom: 20px;
    height: 100%;
    display: flex;
    flex-direction: column;
    .op-tool-c {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-top: 10px;
      .op-btn-c {
        display: flex;
        align-items: center;
        font-size: 14px;
        
        .op-btn {
          margin-right: 1px;
          display: flex;
          align-items: center;
          background-color: rgba(16,16,16,0.1);
          color: rgba(32,37,101,0.9);
          height: 30px;
          padding: 0 10px;
          &:first-child {
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
    }

    .log-content-c {
      display: flex;
      flex: 1;
      height: 0;
      .log-content {
        flex: 1;
        width: 0;
        overflow: auto;
        padding: 0 10px;
        /deep/ pre {
          margin: 0;
        }
      }

      .right-btn-c {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        margin-left: 10px;
        .icon-wrap{
          display: flex;
          align-items: center;
          justify-content: center;
          width: 30px;
          height: 30px;
          background-color: rgba(16,16,16,0.1);
          border-radius: 6px;
          cursor: pointer;
          svg:not([stroke]) {
            fill: rgba(32, 37, 101, 0.9);
          }
        }
      }
    }
  }
}


.log-fullscreen-dlg {
  /deep/ .el-dialog__header {
    display: none;
  }

  /deep/ .el-dialog__body {
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    padding: 25px 30px 30px 50px;

    .log-fullscreen-container {
      height: 100%;
      display: flex;
      flex-direction: column;

      .log-fullscreen-header-c {
        margin-bottom: 10px;
        display: flex;
        align-items: center;
        justify-content: space-between;

        .title {
          display: flex;
          align-items: center;
          font-size: 16px;
          font-weight: 600;
          color: rgba(0, 0, 0, .87);
        }

        .op-tool-c {
          display: flex;
          align-items: center;
          justify-content: space-between;
          margin-bottom: 10px;

          .op-btn-c {
            display: flex;
            align-items: center;
            font-size: 14px;
            justify-content: space-between;
            .op-btn {
              margin-right: 1px;
              display: flex;
              align-items: center;
              background-color: rgba(16,16,16,0.1);
              color: rgba(32,37,101,0.9);
              height: 30px;
              padding: 0 10px;
              &:first-child {
                border-radius: 6px 0px 0px 6px;
              }
              &:last-child {
                border-radius: 0px 6px 6px 0px;
              }
              svg {
                margin-right: 4px;
              }
            }
          }
        }
      }

      .log-fullscreen-content-c {
        flex: 1;
        display: flex;
        height: 0;

        .log-content {
          border: 1px solid rgba(0, 0, 0, .2);
          flex: 1;
          width: 0;
          overflow: auto;
          padding: 0 10px;
          color: #303133;
        }

        .right-btn-c {
          display: flex;
          flex-direction: column;
          justify-content: space-between;
          margin-left: 10px;
          .icon-wrap{
            display: flex;
            align-items: center;
            justify-content: center;
            width: 30px;
            height: 30px;
            background-color: rgba(16,16,16,0.1);
            border-radius: 6px;
            cursor: pointer;
            svg:not([stroke]) {
              fill: rgba(32, 37, 101, 0.9);
            }
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
</style>
