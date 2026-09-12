<template>
  <div>
    <a class="operate-btn" @click.stop.prevent="dlgShow = true" :class="disabled ? 'disabled' : ''">{{
      $t('cloudbrainObj.saveNewModel') }}</a>
    <BaseDialog class="export-model-dlg base-dlg" :visible.sync="dlgShow" :title="$t('cloudbrainObj.saveNewModel')"
      width="950px" :modal="true" :modalAppendToBody="true" :appendToBody="true" :close-on-click-modal="false"
      :show-close="true" :lockScroll="true" :destroy-on-close="false" @open="open" @closed="closed">
      <div class="dlg-content">
        <div class="row-c" :class="loading ? 'disabled' : ''">
          <div class="row">
            <div class="r-title"><label class="required">{{ $t('modelManage.trainTask') }}</label></div>
            <div class="r-content" style="display:flex">
              <el-input class="input-disabled" size="medium" v-model="state.taskName" readonly></el-input>
            </div>
          </div>
          <div class="row" :class="nameErr ? 'error' : ''">
            <div class="r-title"><label class="required">{{ $t('modelSquare.model_name') }}</label></div>
            <div class="r-content">
              <el-input size="medium" v-model="state.name" @keydown.native="handleKeyDown" @blur="checkName"
                :maxlength="100"></el-input>
            </div>
          </div>
          <div class="row" style="margin: -6px 0 0 -190px;">
            <div class="r-title"></div>
            <div class="r-content">
              <span style="font-size: 12px;color: #888;line-height: 1;margin-top: 0.5em;display: inline-block;">
                {{ $t('datasetObj.dataset_name_tooltips') }}
              </span>
            </div>
          </div>
          <div class="row" :class="aliasErr ? 'error' : ''">
            <div class="r-title"><label>{{ $t('modelSquare.model_zh_name') }}</label></div>
            <div class="r-content">
              <el-input size="medium" :maxLength="100" v-model="state.alias" @keydown.native="handleKeyDown"
                @blur="checkAlias"></el-input>
            </div>
          </div>
          <div class="row" style="margin: -6px 0 0 -190px;">
            <div class="r-title"></div>
            <div class="r-content">
              <span style="font-size: 12px;color: #888;line-height: 1;margin-top: 0.5em;display: inline-block;">
                {{ $t('datasetObj.dataset_name_cn_tooltips') }}
              </span>
            </div>
          </div>
          <div class="row">
            <div class="r-title"><label class="required">{{ $t('datasetObj.dataset_owner') }}</label></div>
            <div class="r-content">
              <el-select v-model="state.owner_id" :placeholder="$t('datasetObj.select_category')" style="width:312px;">
                <el-option v-for="item in owners" :key="item.ID" :value="item.ID" :label="item.Name"></el-option>
              </el-select>
            </div>
          </div>
          <div class="row">
            <div class="r-title"><label class="required">{{ $t('modelManage.modelEngine') }}</label></div>
            <div class="r-content">
              <el-select style="width:312px;" size="medium" v-model="state.engine" placeholder="">
                <el-option v-for="item in engineList" :key="item.k" :label="item.v" :value="item.k">
                </el-option>
              </el-select>
            </div>
          </div>

          <div class="row">
            <div class="r-title"><label>{{ $t('modelManage.license') }}</label></div>
            <div class="r-content">
              <el-select style="width:100%;" size="medium" v-model="state.license" class="license-sel"
                :placeholder="$t('modelManage.selectLicense')">
                <template slot="prefix">
                  <i v-if="state.license" class="el-select__caret el-input__icon el-icon-close"
                    @click.stop.prevent="handleClearLicenseClick"></i>
                </template>
                <el-option v-for="item in licenseList" :key="item.id" :label="item.name" :value="item.id">
                </el-option>
              </el-select>
            </div>
          </div>
          <div class="row">
            <div class="r-title"><label>{{ $t('modelManage.modelLabel') }}</label></div>
            <div class="r-content">
              <el-input size="medium" :maxLength="255" v-model="state.label"
                :placeholder="$t('modelManage.modelLabelInputTips')" @input="labelInput"></el-input>
            </div>
          </div>
          <div class="row">
            <div class="r-title"><label>{{ $t('modelManage.modelAccess') }}</label></div>
            <div class="r-content">
              <el-radio v-model="state.isPrivate" label="0">{{ $t('modelManage.modelAccessPublic') }}</el-radio>
              <el-radio v-model="state.isPrivate" label="1">{{ $t('modelManage.modelAccessPrivate') }}</el-radio>
            </div>
          </div>
          <div class="row" style="margin: 0;margin-left: -190px;">
            <div class="r-title"></div>
            <div class="r-content" style="display:flex;line-height: 28px;flex-wrap: wrap;">
              <div class="storage-t">
                {{ $t('storage.remain_storage') }}：<span class="storage-v">{{ formattedRemaining[0] }}</span>
                {{ formattedRemaining[1] }}
                ({{ $t('storage.selected_file_size') }}：
              </div>
              <div class="storage-limit-wrap">
                <span class="select-v">{{ formattedTotal[0] }}</span>{{ formattedTotal[1] }}
                <div class="limit-tip-wrap" v-if="isStorageExceeded">
                  <div class="limit-tip"><i class="ri-information-line"></i>{{ $t('storage.exceedStorage') }}</div>
                  <span>{{ $t('storage.owenerTips', { ownerName: orgName }) }}</span>
                </div>
                )
              </div>
            </div>
          </div>
          <div class="row" :class="modelFileErr ? 'error' : ''">
            <div class="r-title"><label class="required">{{ $t('modelManage.modelFiles') }}</label></div>
            <div class="r-content">
              <el-popover placement="bottom" width="508" trigger="click">
                <div class="treeContainer">
                  <el-tree :data="treeData" show-checkbox default-expand-all node-key="id" ref="fileTreeRef"
                    :props="defaultProps" @check="onFileCheckChange">
                    <span slot-scope="{ data }" class="slot-wrap" style="display: flex;flex:1;">
                      <i class="icon" :class="data.isDir ? 'folder' : 'file'" width="16" height="16"
                        aria-hidden="true"></i>
                      <span>{{ data.label }}</span>
                      <span v-if="!data.isDir"
                        style="margin-left:auto">{{ formatBytes(data.Size)[0] }}{{ formatBytes(data.Size)[1] }}</span>
                    </span>
                  </el-tree>
                </div>
                <div class="add-param-btn" slot="reference">
                  <a href="javascript:;">
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
                  <div class="file-status"
                    v-if="item.statusCode == -1 || item.statusCode == -2 || item.statusCode == -3">
                    <i class="icon ri-close-circle-line" style="color:red"></i>
                    <span>{{ $t('cloudbrainObj.exportDataset.export_failed') }}</span>
                    <el-tooltip placement="top" effect="dark" v-if="item.statusCode == -2">
                      <i class="question circle icon"></i>
                      <div slot="content">
                        <div>{{ $t('cloudbrainObj.exportDataset.export_has_same_file1') }}</div>
                      </div>
                    </el-tooltip>
                    <el-tooltip placement="top" effect="dark" v-if="item.statusCode == -3">
                      <i class="question circle icon"></i>
                      <div slot="content">
                        <div>{{ $t('cloudbrainObj.exportDataset.export_exceed_storage') }}</div>
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
              <el-button size="medium" class="green" @click="submit" :disabled="isStorageExceeded">{{
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
import { MODEL_ENGINES } from '~/const';
import { getListValueWithKey } from '~/utils';
import { getAiTaskOutputResultAll, setAiTaskExportModel, getAiTaskExportModelProgress } from '~/apis/modules/cloudbrain';
import { getModelLicenseList } from '~/apis/modules/modelmanage';
import { getStorageSummary } from "~/apis/modules/storage";
import { getOrgStorageSummary } from '~/apis/modules/organization';
import { getAvailableUsers } from '~/apis/modules/dataset';
const UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB'];

const MAX_LABEL_COUNT = 5;

export default {
  name: 'ExportModel',
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
        taskName: '',
        name: '',
        alias: '',
        owner_id: '',
        engine: 0,
        filesStr: '',
        label: '',
        license: '',
        isPrivate: '0',
      },
      licenseList: [],
      nameErr: false,
      aliasErr: false,
      modelFileErr: false,
      engineList: MODEL_ENGINES,
      treeData: [],
      selectedData: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      remaining_storage: 0,
      totalSize: 0,
      isStorageExceeded: false,
      owners: [],
      progressId: '',
      progressTimer: null,
      orgName: ''
    };
  },
  watch: {
    'state.owner_id': {
      handler(newVal, oldVal) {
        // 在这里处理owner_id变化后的逻辑
        if (newVal !== oldVal) {
          console.log('owner_id从', oldVal, '变更为', newVal, this.owners)
          // 调用相关处理函数

          const findItem = this.owners.filter((item) => {
            return item.ID === newVal
          })
          console.log("findItem", findItem)
          if (findItem[0].IsOrganization) {
            this.orgName = findItem[0].Name
          } else {
            this.orgName = ''
          }
          this.getStorageSummary()
          // this.handleOwnerIdChange(newVal, oldVal)
        }
      },
      deep: false      // 由于是基本类型监听，不需要深度监听
    }
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
    handleKeyDown(e) {
      if (e.key === ' ' || e.keyCode === 32) {
        e.preventDefault(); // 阻止空格输入
      }
    },
    checkName() {
      console.log("checkName", /^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,98}[a-zA-Z0-9]$/.test(this.state.name))
      if (/^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,98}[a-zA-Z0-9]$/.test(this.state.name)) {
        this.nameErr = false;
      } else {
        this.nameErr = true;
      }
      return !this.nameErr;
    },
    checkAlias() {
      console.log("checkAlias", /^[\u4e00-\u9fa5a-zA-Z0-9-_.]{0,100}$/.test(this.state.alias))
      if (/^[\u4e00-\u9fa5a-zA-Z0-9-_.]{0,100}$/.test(this.state.alias)) {
        this.aliasErr = false;
      } else {
        this.aliasErr = true;
      }
      return !this.aliasErr;
    },
    checkModelFile() {
      this.modelFileErr = !this.state.filesStr;
      return !this.modelFileErr;
    },
    labelInput() {
      const hasEndSpace = this.state.label[this.state.label.length - 1] == ' ';
      const list = this.state.label.trim().split(' ').filter(label => label != '');
      this.state.label = list.slice(0, MAX_LABEL_COUNT).join(' ') + (hasEndSpace && list.length < MAX_LABEL_COUNT ? ' ' : '');
    },
    handleClearLicenseClick() {
      this.state.license = '';
    },
    startGetProgressTimer() {
      this.progressTimer && clearInterval(this.progressTimer);
      this.progressTimer = setInterval(() => {
        this.getProgress();
      }, 5 * 1000);
    },
    getProgress(isFirst) {
      getAiTaskExportModelProgress({
        id: this.progressId, // ${this.state.tab}_
      }).then(res => {
        console.log(res);
        res = res.data;
        if (res.code == 0) {
          const result = res.data
          if (isFirst) {
            console.log("xxxxxxxxxx")
            if (result && Object.keys(result).length > 0) {
              if (result['##type##'] !== undefined) {
                this.state.tab = Number(result['##type##']);
              }
              const files = Object.keys(result).filter((item) => item !== '##type##');
              files.forEach(item => {
                const statusCode = result[item]
                this.selectedData.push({
                  label: item,
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
              console.log("files.length", files.length)
              console.log("selectedData", this.selectedData.length)
              for (let i = 0, iLen = files.length; i < iLen; i++) {
                const item = files[i];
                for (let j = 0, jLen = this.selectedData.length; j < jLen; j++) {
                  const selectedFile = this.selectedData[j];
                  if (selectedFile.id == item) {
                    const statusCode = result[item];
                    if (statusCode == -1 || statusCode == -2 || statusCode == -3 || statusCode == 100) {
                      endStatusCount++;
                    }
                    selectedFile.statusCode = statusCode;
                    break;
                  }
                }
              }
              if (endStatusCount == this.selectedData.length) {
                this.loading = false;
                this.progressTimer && clearInterval(this.progressTimer);
              }
            }
          }
        }
      }).catch(err => {
        this.loading = false
        this.progressTimer && clearInterval(this.progressTimer);
        console.log(err);
      });
    },
    submit() {
      console.log("sbmit", this.checkName(), this.checkAlias())
      this.state.name = this.state.name.trim();
      if (!this.checkName()) {
        this.$message({
          type: 'error',
          message: this.$t('modelManage.pleaseInputModelName'),
        });
        return;
      }
      if (!this.checkAlias()) {
        return;
      }
      if (!this.checkModelFile()) {
        this.$message({
          type: 'error',
          message: this.$t('modelObj.model_export_placeholder'),
        });
        return;
      }
      const subData = {
        name: this.state.name,
        alias: this.state.alias,
        owner_id: this.state.owner_id,
        aimodel_type: 0,
        is_private: this.state.isPrivate == 1 ? true : false,
        engine: this.state.engine,
        label: this.state.label.split(/\s+/).join(' ').trim(),
        license: this.state.license,
        task_id: this.data.id,
        file_list: this.state.filesStr,
      }
      this.loading = true
      this.selectedData.forEach(item => item.statusCode = 0);
      setAiTaskExportModel({ id: this.data.id }, subData).then(res => {
        res = res.data;
        console.log("res", res)
        if (res.code == 0) {
          this.progressId = res.data.process_id;
          this.startGetProgressTimer();
        } else {
          this.selectedData.forEach(item => item.statusCode = -99);
          this.loading = false
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      }).catch(err => {
        this.loading = false
        this.selectedData.forEach(item => item.statusCode = -99);
        console.log(err);
      });
    },
    cancel() {
      this.dlgShow = false;
    },
    isMindSporeEngine(obj) {
      if (obj.engine_name != null && obj.engine_name.toLowerCase().startsWith("mindspore")) {
        return true;
      }
      if (obj.engine_id == 122 || obj.engine_id == 35 || obj.engine_id == -1 || obj.engine_id == 37) {
        return true;
      }
      return false;
    },
    onFileCheckChange(data) {
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
      this.checkModelFile();
    },
    removeFile(data) {
      console.log("data", data)
      this.$refs.fileTreeRef.setChecked(data.id, false, false);
      const index = this.selectedData.findIndex((item) => {
        return item.id === data.id;
      });
      this.selectedData.splice(index, 1);
      this.totalSize = this.selectedData.reduce((pre, cur) => {
        return cur.isDir ? pre : pre + cur.Size;
      }, 0);
      this.state.filesStr = this.selectedData.reduce((pre, cur) => {
        return cur.IsDir ? pre : (pre ? pre + ',' : pre) + cur.FileName;
      }, '');
      this.isStorageExceeded = this.totalSize > this.remaining_storage
    },
    open() {
      this.state.filesStr = '';
      this.state.taskName = this.data.display_job_name;
      this.state.name = this.data.display_job_name + '_model_' + Math.random().toString(36).substr(2, 4);
      if (this.isMindSporeEngine(this.data)) {
        this.state.engine = 2;
      } else {
        if (this.data.engine_id == 121 || this.data.engine_id == 38) {
          this.state.engine = 1;
        } else {
          this.state.engine = 0;
        }
      }
      this.selectedData = []
      this.state.label = '';
      this.state.description = '';
      this.loading = true;
      getAiTaskOutputResultAll({
        id: this.data.id,
      }).then(res => {
        res = res.data;
        this.loading = false;
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
        }
      }).catch(err => {
        console.log(err);
        this.loading = false;
      });
      getModelLicenseList().then(res => {
        res = res.data;
        try {
          const license = JSON.parse(res) || [];
          this.licenseList = license;
          if (license.length) {
            this.state.license = license[0].id;
          }
        } catch (err) {
          console.log(err);
        }
      }).catch(err => {
        console.log(err);
      });

    },
    closed() {

      this.progressTimer && clearInterval(this.progressTimer);
    },
    async getStorageSummary() {
      const queryParams = !this.orgName
        ? getStorageSummary({})
        : getOrgStorageSummary(this.orgName, {});

      const response = await queryParams
      let res = !this.orgName ? response.data : response.data.data
      console.log("res", res)
      this.remaining_storage = res.remaining_storage || 0
      this.totalSize = this.selectedData.reduce((pre, cur) => {
        return cur.isDir ? pre : pre + cur.Size;
      }, 0);
      this.isStorageExceeded = this.totalSize > this.remaining_storage
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
    async getUsersList() {
      try {
        const response = await getAvailableUsers({}, 'aimodel')
        const res = response.data
        this.loading = false
        if (res.code === 0) {
          this.owners = res.data.users || []
          if (this.owners.length === 0) {
            return
          }
          this.state.owner_id = this.owners[0].ID
        } else {
          this.$message.error(response.data.msg);
        }
      } catch (error) {
        this.loading = false
        this.$message.error(error);
      }
    },
  },
  beforeMount() {
    this.getUsersList()
  },
  mounted() {
    // this.getStorageSummary()
  }
};
</script>

<style scoped lang="less">
.dlg-content {
  padding: 30px 0;

  .row-c {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    margin: 0 auto;
    // width: 80%;

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
        width: 218px;
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
      }
    }
  }
}

.license-sel {
  /deep/.el-input--prefix .el-input__inner {
    padding-left: 15px;
  }

  /deep/.el-input__prefix {
    position: absolute;
    right: 0;

    .el-icon-close {
      position: absolute;
      right: 24px;
      color: rgba(0, 0, 0, .87);
      font-weight: bold;
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

/deep/ .el-radio.is-checked {
  .el-radio__inner {
    // border-color: #409EFF;
    // background: #409EFF;
    border-color: rgb(16, 16, 16);
    background: rgb(16, 16, 16);
  }

  .el-radio__label {
    // color: #409EFF;
    color: rgb(16, 16, 16);
  }
}

.treeContainer {
  max-height: 600px;
  overflow: auto;
  min-height: 232px;

}
</style>
