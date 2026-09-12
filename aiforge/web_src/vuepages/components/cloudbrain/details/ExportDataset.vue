<template>
  <div>
    <a class="operate-btn" @click.stop.prevent="dlgShow = true" :class="disabled ? 'disabled' : ''">{{
      $t('cloudbrainObj.exportToDataset') }}</a>
    <BaseDialog class="export-model-dlg base-dlg" :visible.sync="dlgShow" :title="$t('cloudbrainObj.exportToDataset')"
      width="900px" :modal="true" :modalAppendToBody="true" :appendToBody="true" :close-on-click-modal="false"
      :show-close="true" :lockScroll="true" :destroy-on-close="false" @open="open" @closed="closed">
      <div class="dlg-content" style="min-width: 900px;">
        <div class="row-c" v-loading="loading">
          <div class="tips" v-html="$t('cloudbrainObj.exportDataset.exportDatasetTips')"></div>

          <div class="row" style="align-items:flex-start;">
            <div class="r-title" style="margin-top:5px"><label class="required">{{ $t('datasetObj.dataset_name1')
                }}</label>
            </div>
            <div class="r-content">
              <div class="select-dataset">
                <el-input size="medium" :maxLength="255" v-model="state.dataset_name" readonly
                  style="flex:1"></el-input>
                <el-popover placement="right" width="500" @show="searchModelData" popper-class="export_dataset"
                  trigger="click" v-model="popoverVisible">
                  <div class="dataset-wrap">
                    <div class="dataset-title">
                      <span>选择数据集</span>
                      <i class="el-icon-close" style="cursor:pointer;" @click="popoverVisible = false"></i>
                    </div>
                    <div class="datatset-main">
                      <div class="model-tabs-c">
                        <div class="model-tabs-w">
                          <el-tabs class="model-tabs" v-model="dlgActiveName" @tab-click="dlgTabClick">
                            <el-tab-pane :label="$t('cloudbrainObj.ihave')" name="first"></el-tab-pane>
                            <el-tab-pane :label="$t('cloudbrainObj.iCollaborate')" name="second"></el-tab-pane>
                          </el-tabs>
                          <el-input size="small" class="search-inp"
                            :placeholder="$t('datasetObj.dataset_search_placeholder')" v-model="dlgSearchValue"
                            @keydown.enter.stop.native.prevent="inputSearch">
                            <div slot="suffix" class="search-inp-icon" @click="inputSearch">
                              <i class="el-icon-search"></i>
                            </div>
                          </el-input>
                        </div>

                        <div class="datalist-wrap">
                          <el-radio-group class="data-wrap" v-model="dataset_id" @input="changeDataset">
                            <el-radio class="data-item" v-for="(item, index) in dlgModelTreeData" :key="item.id"
                              :label="item.id" :title="`${item.owner_name}/${item.alias}`">
                              <div class="data-info-c">
                                <div class="data-info nowrap">
                                  <span> {{ item.owner_name }} / </span>
                                  <span style="color: #101010;font-weight: 700;">{{ item.alias }}</span>
                                </div>
                                <a :href="`/datasets/detail/${item.owner_name}/${item.name}`" target="_blank">
                                  <svg width="16" height="16" viewBox="0 0 48 48" fill="none"
                                    xmlns="http://www.w3.org/2000/svg">
                                    <path d="M28 6H42V20" stroke="#10101080" stroke-width="4" stroke-linecap="round"
                                      stroke-linejoin="round" />
                                    <path
                                      d="M42 29.4737V39C42 40.6569 40.6569 42 39 42H9C7.34315 42 6 40.6569 6 39V9C6 7.34315 7.34315 6 9 6L18 6"
                                      stroke="#10101080" stroke-width="4" stroke-linecap="round"
                                      stroke-linejoin="round" />
                                    <path d="M25.7998 22.1999L41.0998 6.8999" stroke="#10101080" stroke-width="4"
                                      stroke-linecap="round" stroke-linejoin="round" />
                                  </svg>
                                </a>
                              </div>
                            </el-radio>
                          </el-radio-group>
                        </div>
                        <div class="pagination-c">
                          <el-pagination background @current-change="dlgPageChange" :current-page="dlgPage"
                            :page-size="dlgPageSize" :pager-count="5" layout="total, prev, pager, next"
                            :total="dlgTotal">
                          </el-pagination>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="btn-select" slot="reference">
                    <i class="el-icon-plus"></i>
                    <span>{{ $t('datasetObj.dataset_select') }}</span>
                  </div>
                </el-popover>
              </div>
            </div>
          </div>
          <div class="row" v-if="!!dataset_id" style="margin: 0;margin-left: -190px;">
            <div class="r-title"></div>
            <div class="r-content" style="display:flex;line-height: 28px;flex-wrap: wrap;">
              <span class="storage-t">
                {{ $t('storage.remain_storage') }}：<span class="storage-v">{{ formattedRemaining[0] }}
                </span>{{ formattedRemaining[1] }}
                ({{ $t('storage.selected_file_size') }}：
              </span>
              <span class="storage-limit-wrap" style="display: inline;">
                <span class="select-v">{{ formattedTotal[0] }}</span>{{ formattedTotal[1] }}
                <span class="limit-tip-wrap" style="display:inline" v-if="isStorageExceeded">
                  <span class="limit-tip" style="display: inline-flex;"><i class="ri-information-line"></i>{{
                    $t('storage.exceedStorage') }}</span>
                  <a v-if="!showOwenerTips" href="/storages">{{ $t('storage.capacity_details') }}</a>
                  <span v-else>{{ $t('storage.owenerTips', { ownerName: ownerName }) }}</span>
                </span>
                )
              </span>
            </div>
          </div>
          <div class="row">
            <div class="r-title"><label class="required">{{ $t('cloudbrainObj.exportDataset.please_select_output_file')
                }}</label></div>
            <div class="r-content" style="display:flex">
              <el-popover placement="bottom-start" width="735" trigger="click" :disabled="uploading">
                <div class="treeContainer">
                  <el-tree :data="treeData" show-checkbox default-expand-all node-key="id" ref="fileTreeRef"
                    :props="defaultProps" @check="onFileCheckChange">
                    <span slot-scope="{ data }" class="slot-wrap" style="display: flex;flex:1;">
                      <i class="icon" :class="data.IsDir ? 'folder' : 'file'" width="16" height="16"
                        aria-hidden="true"></i>
                      <span>{{ data.label }}</span>
                      <span v-if="!data.isDir" style="margin-left:auto">{{ formatBytes(data.Size)[0] }}{{
                        formatBytes(data.Size)[1] }}</span>
                    </span>
                  </el-tree>
                </div>
                <div class="add-param-btn" slot="reference">
                  <a href="javascript:;" :class="uploading ? 'disabled' : ''">
                    <i class="plus square outline icon"></i>
                    <span>{{ $t('cloudbrainObj.exportDataset.select_file') }}</span>
                  </a>
                </div>
              </el-popover>
            </div>
          </div>
          <div class="row" style="margin-top:-2px">
            <div class="r-title"><label></label></div>
            <div class="r-content">
              <div class="file-item-list">
                <div class="file-item" v-for="item in selectedData" :key="item.FileName">
                  <span class="file-name" :title="item.FileName">{{ item.FileName }}</span>
                  <i class="icon delete icon-delete" v-if="item.statusCode == -99" @click="removeFile(item)"></i>
                  <div class="file-status" v-if="item.statusCode == 0">
                    <i class="icon el-icon-loading" style="color:#21ba45;margin-top:0"></i>
                    <span>{{ $t('cloudbrainObj.exportDataset.exporting') }}</span>
                  </div>
                  <div class="file-status" v-if="item.statusCode == -1 || item.statusCode == -2">
                    <i class="icon ri-close-circle-line" style="color:red"></i>
                    <span>{{ $t('cloudbrainObj.exportDataset.export_failed') }}</span>
                    <el-tooltip placement="top" effect="dark" v-if="item.statusCode == -2">
                      <i class="question circle icon"></i>
                      <div slot="content">
                        <div>{{ $t('cloudbrainObj.exportDataset.export_has_same_file') }}</div>
                      </div>
                    </el-tooltip>
                  </div>
                  <div class="file-status" v-if="item.statusCode == 100">
                    <i class="icon ri-checkbox-circle-line" style="color:#21ba45"></i>
                    <span>{{ $t('cloudbrainObj.exportDataset.export_success') }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="row" style="margin-top:20px">
            <div class="r-title"><label></label></div>
            <div class="r-content btn-c">
              <el-button size="medium" class="green" @click="submit" :disabled="uploading || isStorageExceeded">{{
                $t('modelManage.confirm')
                }}</el-button>
              <el-button size="medium" @click="cancel">{{ $t('modelManage.cancel') }}</el-button>
            </div>
          </div>
        </div>
      </div>
    </BaseDialog>
  </div>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';
import { getRepoDatasetInfo, getAiTaskOutputResultAll, getAiTaskExportDatasetProgress, setAiTaskExportDataset } from '~/apis/modules/cloudbrain';
import { getStorageSummary } from "~/apis/modules/storage";
import { getDatasets } from "~/apis/modules/dataset";
const UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB'];

export default {
  name: 'ExportDataset',
  props: {
    disabled: { type: Boolean, default: true },
    configs: { type: Object, default: () => { return {} } },
    data: { type: Object, default: () => { return {} } },
  },
  components: { BaseDialog, },
  data() {
    return {
      dlgShow: false,
      loading: false,
      state: {
        dataset_name: '',
        tab: 1,
        filesStr: '',
      },
      dataset_id: '',
      progressId: '',
      firstOpen: false,
      treeData: [],
      selectedData: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      uploading: false,
      progressTimer: null,
      remaining_storage: 0,
      totalSize: 0,
      isStorageExceeded: false,
      dlgActiveName: 'first',
      dlgSearchValue: '',
      dlgLoading: false,
      dlgModelTreeData: [],
      dlgTotal: 0,
      dlgPage: 1,
      dlgPageSize: 6,
      popoverVisible: false,
      loginName: '',
      ownerName: '',
      showOwenerTips: false,
    };
  },
  computed: {
    // 格式化的剩余空间显示（自动单位转换）
    formattedRemaining() {
      return this.formatBytes(Math.max(this.remaining_storage, 0));
    },
    formattedTotal() {
      return this.formatBytes(this.totalSize);
    },
  },
  methods: {
    changeDataset(id) {
      this.dlgModelTreeData.forEach((item) => {
        if (item.id == id) {
          this.state.dataset_name = item.alias;
          console.log(item.alias)
          this.ownerName = item.owner_name;
          if (this.loginName != this.ownerName) {
            this.showOwenerTips = true
          }
        }
      })
      this.getStorageSummary(id)
      console.log(this.state.dataset_name)
    },
    dlgTabClick(tab, event) {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.searchModelData();
    },
    inputSearch() {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.searchModelData();
    },
    dlgPageChange(page) {
      this.dlgPage = page;
      this.searchModelData();
    },
    searchModelData() {
      const tabName = this.dlgActiveName;
      const tabPrams = {
        'first': '/owned',
        'second': '/collaborated',
      }
      const params = {
        url: tabPrams[tabName],
        q: this.dlgSearchValue.trim(),
        page: this.dlgPage,
        page_size: this.dlgPageSize,
      };
      this.dlgLoading = true;
      getDatasets(params).then(res => {
        console.log(res)
        this.dlgLoading = false;
        if (res.data.code === 0) {
          const data = res.data?.data?.datasets || [];
          this.dlgModelTreeData = data
          this.dlgTotal = res.data?.data?.total || 0;
        } else {
          this.$message.error(rs.data.msg || '获取数据集失败')
          console.log(err);
        }

      }).catch(err => {
        this.dlgLoading = false;
        console.log(err);
      });
    },
    onFileCheckChange() {
      const selectedData = this.$refs.fileTreeRef.getCheckedNodes();
      const fliterfile = selectedData.filter((item) => {
        return !item.isDir
      })
      this.state.filesStr = fliterfile.reduce((pre, cur) => {
        return cur.isDir ? pre : (pre ? pre + ',' : pre) + cur.FileName;
      }, '');
      this.selectedData = [...fliterfile.map(item => ({ ...item, statusCode: -99 }))];
      this.totalSize = fliterfile.reduce((pre, cur) => {
        return cur.isDir ? pre : pre + cur.Size;
      }, 0);
      this.isStorageExceeded = this.totalSize > this.remaining_storage
    },
    removeFile(data) {
      this.$refs.fileTreeRef.setChecked(data.id, false, false);
      const index = this.selectedData.findIndex((item) => {
        return item.id === data.id;
      });
      this.selectedData.splice(index, 1);
      this.state.filesStr = this.selectedData.reduce((pre, cur) => {
        return cur.IsDir ? pre : (pre ? pre + ',' : pre) + cur.FileName;
      }, '');
    },
    getFiles() {
      getAiTaskOutputResultAll({
        id: this.data.id,
      }).then(res => {
        res = res.data;
        if (res.code == 0) {
          const data = res.data?.output?.file_list || [];
          const nodeMap = {};
          for (let i = 0, iLen = data.length; i < iLen; i++) {
            let dataI = data[i];
            const path = dataI.FileName.split('/');
            let curNode = nodeMap;
            for (let j = 0, jLen = path.length; j < jLen; j++) {
              const cur = path[j];
              if (!curNode[cur]) {
                curNode[cur] = {};
              }
              if (j == jLen - 1) {
                dataI._isLeaf = true;
                curNode[cur] = dataI;
              }
              curNode = curNode[cur];
            }
          }
          const nodeData = [];
          const walkNode = (curNode, nodeList) => {
            if (curNode._isLeaf) return;
            for (let key in curNode) {
              const node = {
                label: key,
                isDir: !curNode[key]._isLeaf,
                children: [],
                id: curNode[key].FileName,
              };
              nodeList.push(node);
              if (curNode[key]._isLeaf) {
                delete node.children;
                Object.assign(node, curNode[key]);
              }
              walkNode(curNode[key], node.children);
            }
          };
          walkNode(nodeMap, nodeData);
          this.treeData = nodeData;
        } else {
          this.treeData = [];
        }
      }).catch(err => {
        console.log(err);
      });
    },
    startGetProgressTimer() {
      this.progressTimer && clearInterval(this.progressTimer);
      this.progressTimer = setInterval(() => {
        this.getProgress();
      }, 5 * 1000);
    },
    getProgress(isFirst) {
      getAiTaskExportDatasetProgress({
        id: this.progressId, // ${this.state.tab}_
      }).then(res => {
        console.log(res);
        res = res.data;
        if (res.code == 0) {
          const result = res.data
          if (isFirst) {
            if (result && Object.keys(result).length > 0) {
              if (result['##type##'] !== undefined) {
                this.state.tab = Number(result['##type##']);
              }
              const files = Object.keys(result).filter((item) => item !== '##type##');
              files.forEach(item => {
                const statusCode = result[item]
                this.selectedData.push({
                  label: item,
                  id: item,
                  FileName: item,
                  children: [],
                  statusCode: statusCode,
                });
                if (statusCode == 0) {
                  this.uploading = true;
                  this.startGetProgressTimer();
                }
              });
            }
          } else {
            if (result && Object.keys(result).length > 0) {
              const files = Object.keys(result).filter((item) => item !== '##type##');
              let endStatusCount = 0;
              for (let i = 0, iLen = files.length; i < iLen; i++) {
                const item = files[i];
                for (let j = 0, jLen = this.selectedData.length; j < jLen; j++) {
                  const selectedFile = this.selectedData[j];
                  if (selectedFile.id == item) {
                    const statusCode = result[item];
                    if (statusCode == -1 || statusCode == -2 || statusCode == 100) {
                      endStatusCount++;
                    }
                    selectedFile.statusCode = statusCode;
                    break;
                  }
                }
              }
              if (endStatusCount == this.selectedData.length) {
                this.uploading = false;
                this.progressTimer && clearInterval(this.progressTimer);
              }
            }
          }
        }
      }).catch(err => {
        this.progressTimer && clearInterval(this.progressTimer);
        console.log(err);
      });
    },
    submit() {
      if (!this.dataset_id) {
        this.$message({
          type: 'info',
          message: this.$t('cloudbrainObj.exportDataset.please_select_dataset'),
        });
        return;
      }
      if (!this.selectedData.length || !this.state.filesStr.length) {
        this.$message({
          type: 'info',
          message: this.$t('cloudbrainObj.exportDataset.please_select_file'),
        });
        return;
      }
      const subData = {
        task_id: this.data.id,
        file_list: this.state.filesStr,
      };
      this.uploading = true;
      this.selectedData.forEach(item => item.statusCode = 0);
      console.log("this.selectedData", this.selectedData)
      console.log("subData", subData)

      setAiTaskExportDataset({ dataset_id: this.dataset_id, id: this.data.id }, subData).then(res => {
        console.log("res", res)
        res = res.data;
        if (res.code == 0) {
          this.progressId = res.data.id;
          this.startGetProgressTimer();
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      }).catch(err => {
        this.uploading = false;
        console.log(err);
      });
    },
    cancel() {
      this.dlgShow = false;
    },
    open() {
      if (this.firstOpen) return
      this.uploading = false;
      this.getFiles()
      this.firstOpen = true;
    },
    closed() {
      this.progressTimer && clearInterval(this.progressTimer);
    },
    async getStorageSummary(id) {
      const res = await getStorageSummary({
        subject_id: id,
        subject_type: '1'
      })
      this.remaining_storage = res.data.remaining_storage
      this.onFileCheckChange()
    },
    // 智能单位格式化
    formatBytes(bytes) {
      let unitIndex = 0;
      let value = bytes;

      while (value >= 1024 && unitIndex < UNITS.length - 1) {
        value /= 1024;
        unitIndex++;
      }

      return [this.toPrecision(value, 2), UNITS[unitIndex]]
    },
    // 精确小数处理
    toPrecision(value, decimals = 2) {
      if (value < 0.001 && value > 0) return '<0.001';
      return Number(value.toFixed(decimals)).toString();
    },
  },
  beforeMount() { },
  mounted() {
    console.log('this.data', this.data)
    this.loginName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext')

  }
};
</script>

<style scoped lang="less">
.dlg-content {
  padding: 30px 0;
  min-height: 288px;

  .row-c {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    margin: 0 auto;
    width: 90%;

    .tips {
      color: #888;
      font-size: 12px;
      margin-top: -10px;
      margin-bottom: 10px;

      span {
        color: red;
      }
    }

    .row {
      width: 100%;
      display: flex;
      align-items: center;
      margin: 8px 0;
      margin-left: -190px;

      .r-title {
        text-align: right;
        font-size: .92857143em;
        font-weight: 700;
        color: rgba(0, 0, 0, .87);
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

      &.error {
        .r-title {
          color: #9f3a38;
        }

        .r-content {
          /deep/.el-input__inner {
            color: #9f3a38;
            background: #fff6f6;
            border-color: #e0b4b4;

            &::placeholder {
              color: #e0b4b4;
            }
          }
        }
      }

      .r-content {
        flex: 1;

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

        .select-dataset {
          display: flex;
          align-items: center;

          .btn-select {
            cursor: pointer;
            color: #0366d6;
            margin-left: 10px;
          }

        }

        .storage-t {
          font-family: Arial;
          color: rgba(2, 0, 4, 0.5);

          .storage-v {
            color: rgba(16, 16, 16, 1);
            font-weight: 700;
          }
        }

        .storage-limit-wrap {
          color: #101010;
          font-family: Arial;
          display: flex;
          margin: 0 6px;

          .select-v {
            color: rgba(246, 106, 0, 1);
            font-weight: 700;
            margin-right: 4px;
          }

          .limit-tip-wrap {
            display: flex;

            .limit-tip {
              display: flex;
              align-items: center;
              height: 28px;
              border-radius: 4px;
              background-color: rgba(250, 140, 22, 1);
              color: rgba(255, 255, 255, 1);
              padding: 0 8px;
              margin: 0 6px;

              i {
                font-size: 14px;
                margin-right: 6px;
              }
            }

            a {
              color: #101010;
              text-decoration: underline;
            }
          }
        }
      }
    }
  }
}

.input-disabled {
  /deep/ .el-input__inner {
    background-color: #f5f5f6 !important;
    color: #888888 !important;
  }
}

.el-select-dropdown__item.selected {
  color: rgba(0, 0, 0, .95);
}

.btn-c {
  /deep/ .el-button {
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
      background-color: #5bb973;
      color: #fff;

      &:hover {
        background-color: #16ab39;
        border-color: transparent;
      }

      &:focus {
        background-color: #0ea432;
        border-color: transparent;
      }

      &:active {
        background-color: #198f35;
        border-color: transparent;
      }
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

// /deep/ .el-radio.is-checked {
//   .el-radio__inner {
//     // border-color: #409EFF;
//     // background: #409EFF;
//     border-color: rgb(16, 16, 16);
//     background: rgb(16, 16, 16);
//   }

//   .el-radio__label {
//     // color: #409EFF;
//     color: rgb(16, 16, 16);
//   }
// }

.tab-c {
  display: flex;
  align-items: center;

  .tab {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 12px;
    color: rgba(0, 0, 0, .87);
    border: 1px solid rgba(34, 36, 38, .15);
    margin-left: -1px;
    height: 34px;
    padding: 0 12px;
    border-left: none;

    i {
      margin-top: -7px;
      font-size: 14px;
    }

    &.focus {
      color: #0087f5;
      border-color: #0087f5;
      border-left: 1px solid #0087f5;
    }

    &:first-child {
      border-top-left-radius: 0.28571429rem;
      border-bottom-left-radius: 0.28571429rem;
      border-left: 1px solid rgba(34, 36, 38, .15);

      &.focus {
        border-color: #0087f5;
      }
    }

    &:last-child {
      border-top-right-radius: 0.28571429rem;
      border-bottom-right-radius: 0.28571429rem;
    }

    &:hover:not(.focus) {
      background: rgba(0, 0, 0, .03);
    }
  }
}

.treeContainer {
  max-height: 600px;
  overflow: auto;
  padding-right: 16px;
}

.dataset-wrap {
  .dataset-title {
    display: flex;
    justify-content: space-between;
  }

  .datatset-main {
    .model-tabs-c {
      display: flex;
      flex-direction: column;

      .model-tabs-w {
        display: flex;
        align-items: center;

        .model-tabs {
          flex: 1;
          overflow: hidden;
          margin-right: 5px;
        }

        .search-inp {
          overflow: hidden;
          width: 200px;
          z-index: 5;
          position: relative;
          margin-top: -10px;

          .search-inp-icon {
            height: 100%;
            display: flex;
            justify-content: center;
            align-items: center;
            width: 22px;
            cursor: pointer;
          }
        }
      }

    }

    .datalist-wrap {
      .data-wrap {
        display: flex;
        flex-direction: column;

        .data-item {
          display: flex;
          align-items: center;
          height: 40px;
          border-bottom: 1px solid rgba(16, 16, 16, .1);

          .data-info-c {
            display: flex;
            justify-content: space-between;

            .data-info {
              max-width: 90%;
            }
          }

          &.el-radio {
            margin-right: 20px;
          }

          /deep/ .el-radio__label {
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
            flex: 1;
          }
        }
      }
    }

    .pagination-c {
      margin-top: 18px;
      text-align: center;
      overflow: hidden;

      .el-pagination {
        flex-wrap: wrap;
        /* 允许换行 */
        justify-content: center;
        max-width: 100%;
        /* 不超过父容器 */
      }
    }
  }
}

.file-item-list {
  max-height: 380px;
  overflow-y: auto;

  .file-item {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 4px 0;

    .file-name {
      flex: 1;
      width: 0;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .icon-delete {
      width: 20px;
      cursor: pointer;
    }

    .file-status {
      display: flex;
      font-size: 14px;
      align-items: center;

      i {
        margin-top: -7px;
        margin-left: 3px;
      }
    }
  }
}

.empty-dataset {
  .item-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 50px;

    .item-empty-icon {
      height: 100px;
      width: 100px;
      background: url(/img/empty-box.svg) center center no-repeat;
      margin: 2rem 0 1rem;
    }

    .item-empty-tips {
      text-align: center;
      margin-top: 16px;
      font-size: 16px;
      color: #3f3f40;
    }
  }
}
</style>
<style lang="less"></style>