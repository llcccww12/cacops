<template>
     <div class="dlg-content" v-loading="loading">
        <div class="row-c">
        <div class="row">
            <div class="r-title"><label class="required">{{ $t('modelManage.trainTask') }}</label></div>
            <div class="r-content" style="display:flex">
            <el-select v-model="currentTrainJob" value-key="k" placeholder="" @change="changeTrain" v-loading="loadingJob">
                <el-option v-for="item in trainJobList ":key="item.k" :label="item.v" :value="item">
                </el-option>
            </el-select>
            </div>
        </div>
        <div class="row" :class="nameErr ? 'error' : ''">
            <div class="r-title"><label class="required">{{ $t('modelManage.modelName') }}</label></div>
            <div class="r-content">
            <el-input :maxLength="127" v-model="state.name" @blur="checkName"
                :placeholder="$t('modelManage.pleaseInputModelName')">
            </el-input>
            </div>
        </div>
        <div class="row">
            <div class="r-title"><label class="required">{{ $t('modelManage.modelVersion') }}</label></div>
            <div class="r-content">
            <el-input class="input-disabled" v-model="state.version" readonly>
            </el-input>
            </div>
        </div>
        <div class="row">
            <div class="r-title"><label class="required">{{ $t('modelManage.modelEngine') }}</label></div>
            <div class="r-content">
            <el-select v-model="state.engine" placeholder="">
                <el-option v-for="item in engineList " :key="item.k" :label="item.v" :value="item.k">
                </el-option>
            </el-select>
            </div>
        </div>
        <div class="row">
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
                <a href="/storages">{{ $t('storage.capacity_details') }}</a>
              </div>
              )
            </div>
          </div>
        </div>
        <div class="row" :class="modelFileErr ? 'error' : ''">
            <div class="r-title"><label class="required">{{ $t('modelManage.modelFiles') }}</label></div>
            <div class="r-content">
              <el-popover placement="bottom-start" :width="treeWidth" trigger="click">
                  <div class="treeContainer">
                    <el-tree :data="treeData" show-checkbox default-expand-all node-key="id" ref="fileTreeRef"
                        :props="defaultProps" @check="onFileCheckChange">
                        <span slot-scope="{ data }" class="slot-wrap" style="display: flex;flex:1;">
                        <i class="icon" :class="data.isDir ? 'folder' : 'file'" width="16" height="16"
                            aria-hidden="true"></i>
                        <span>{{ data.label }}</span>
                        <span v-if="!data.isDir" style="margin-left:auto">{{formatBytes(data.Size)[0]}}{{formatBytes(data.Size)[1]}}</span>
                        </span>
                    </el-tree>
                  </div>
                  <el-input slot="reference" id="test" :placeholder="$t('modelObj.model_select_placeholder')"
                  v-model="state.filesStr" readonly v-loading="loadingFile">
                  </el-input>
              </el-popover>
            </div>
        </div>
        <div class="row">
            <div class="r-title"><label>{{ $t('modelManage.license') }}</label></div>
            <div class="r-content">
            <el-select v-model="state.license" class="license-sel"
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
            <el-input :maxLength="255" v-model="state.label"
                :placeholder="$t('modelManage.modelLabelInputTips')" @input="labelInput"></el-input>
            </div>
        </div>
        <div class="row" v-if="repoIsPrivate == false">
            <div class="r-title"><label>{{ $t('modelManage.modelAccess') }}</label></div>
            <div class="r-content">
            <el-radio v-model="state.isPrivate" label="0">{{ $t('modelManage.modelAccessPublic') }}</el-radio>
            <el-radio v-model="state.isPrivate" label="1">{{ $t('modelManage.modelAccessPrivate') }}</el-radio>
            </div>
        </div>
        <div class="row" style="align-items:flex-start;">
            <div class="r-title"><label>{{ $t('modelManage.modelBriefIntro') }}</label></div>
            <div class="r-content">
            <el-input type="textarea" :maxLength="255" v-model="state.description" :rows="3"
                :placeholder="$t('modelManage.modelDescrInputTips')">
            </el-input>
            </div>
        </div>
        <div class="row" style="margin-top:20px">
            <div class="r-title"><label></label></div>
            <div class="r-content btn-c">
            <el-button type="info" @click="back">{{ $t('modelManage.back') }}</el-button>
            <el-button class="green" @click="submit" :disabled="isStorageExceeded">{{ $t('modelManage.createModel') }}</el-button>
            <el-button type="info" @click="cancel">{{ $t('modelManage.cancel') }}</el-button>
            </div>
        </div>
        </div>
    </div>
</template>
  
<script>
  import { MODEL_ENGINES } from '~/const';
  import { getAiTaskOutputResultAll, setAiTaskResultToModelApi } from '~/apis/modules/cloudbrain';
  import { getModelLicenseList,getTrainJobList } from '~/apis/modules/modelmanage';
  import { getStorageSummary } from "~/apis/modules/storage";
  const UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB'];
  const MAX_LABEL_COUNT = 5;
  const {AppSubUrl} = window.config;
  export default {
    props: {
        repoOwnerName: { type: String, default: '' },
        repoName: { type: String, default: '' },
        isprivate: { type: Boolean, default: false },
        repoId:{type: Number, default: 0}
    },
    data() {
      return {
        loading: false,
        loadingJob: false,
        loadingFile:false,
        state: {
            jobId: '',
            versionName:'V0001',
            name: '',
            version: '0.0.1',
            engine: '0',
            filesStr: '',
            label: '',
            license: '',
            description: '',
            isPrivate: this.isprivate ? '1' : '0',
        },
        licenseList: [],
        nameErr: false,
        modelFileErr: false,

        engineList: MODEL_ENGINES,
        repoIsPrivate: this.isprivate,
        treeData: [],
        defaultProps: {
          children: 'children',
          label: 'label'
          },
        trainJobList:[],
        currentTrainJob: {},
        treeWidth: '',
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
    methods: {
        changeTrain(val) {
            this.state.jobId = val.jobId
            this.state.filesStr = ''
            this.getReusltList(val.k)
        },
        initTrainList() {
            this.loadingJob = true
            this.state.name = this.repoName + '_model_' + Math.random().toString(36).substr(2, 4);
            const params = {repo:`/${this.repoOwnerName}/${this.repoName}`,repoId:this.repoId}
            getTrainJobList(params).then((res) => {
                this.loadingJob = false
                if (res.data.length >0){
                    res.data.map((item) => {
                        this.trainJobList.push({k:item.ID,v:item.DisplayJobName,jobId:item.JobID})
                    })
                    this.state.jobId = this.trainJobList[0].jobId
                    this.currentTrainJob = this.trainJobList[0]
                    this.getReusltList(this.trainJobList[0].k)
                }
            }).catch((err) => {
                this.loadingJob = false
                this.$message({
                    type: 'error',
                    message: err,
                });
            })
        },
        getReusltList(id) {
            this.loadingFile = true
            getAiTaskOutputResultAll({
                repoOwnerName: this.repoOwnerName,
                repoName: this.repoName,
                id: id,
            }).then(res => {
                this.loadingFile = false
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
                this.loadingFile = false
                this.$message({
                    type: 'error',
                    message: err,
                });
            });
      },
        checkName() {
            this.nameErr = !this.state.name;
            return !this.nameErr;
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
        submit() {
            this.state.name = this.state.name.trim();
            if (!this.checkName()) {
                this.$message({
                    type: 'info',
                    message: this.$t('modelManage.pleaseInputModelName'),
                });
                return;
            }
            if (!this.checkModelFile()) {
            this.$message({
                type: 'info',
                message: this.$t('modelObj.model_select_placeholder'),
            });
            return;
            }
            this.loading = true
            const subData = {
                repoOwnerName: this.repoOwnerName,
                repoName: this.repoName,
                jobId:this.state.jobId,
                versionName:this.state.versionName,
                name: this.state.name,
                version: this.state.version,
                engine: this.state.engine,
                modelSelectedFile: this.state.filesStr,
                label: this.state.label.split(/\s+/).join(' ').trim(),
                license: this.state.license,
                isPrivate: (this.repoIsPrivate || this.state.isPrivate == 1) ? true : false,
                description: this.state.description,
            }
            setAiTaskResultToModelApi(subData).then(res => {
                this.loading = false
                res = res.data;
                if (res.code == 0) {
                    window.location.href = `/${this.repoOwnerName}/${this.repoName}/modelmanage/model_readme_tmpl?name=${encodeURIComponent(this.state.name)}`;
                } else {
                    this.$message({
                        type: 'error',
                        message: res.msg,
                    });
                }
            }).catch(err => {
                this.loading = false
                this.$message({
                    type: 'error',
                    message: err.response.data ?? err,
                });
            });
        },
        cancel() {
            window.location.href = `${AppSubUrl}/dashboard`
        },
        back() {
            this.$emit('backOne');
        },
        onFileCheckChange() {
            const selectedData = this.$refs.fileTreeRef.getCheckedNodes();
            this.state.filesStr = selectedData.reduce((pre, cur) => {
              return cur.isDir ? pre : (pre ? pre + ';' : pre) + cur.FileName;
            }, '');
            this.totalSize = selectedData.reduce((pre, cur) => {
              return cur.isDir ? pre : pre + cur.Size;
            }, 0);
            this.isStorageExceeded = this.totalSize > this.remaining_storage
            this.checkModelFile();
        },
        async getStorageSummary(){
          const res = await getStorageSummary({})
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
    },
    mounted() {
      // this.open()
        this.treeWidth = document.getElementsByClassName('r-content')[0].clientWidth
        this.initTrainList()
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
        this.getStorageSummary()
     }
  };
</script>
  
<style scoped lang="less">
  ::v-deep .el-button--info{
    color: #020004;
    background-color: #c2c7cc;
    border-color: #c2c7cc;
  }
  .dlg-content {
    margin: 3rem 2rem 0 6rem;
    .row-c {
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
      width: 100%;
  
      .row {
        width: 100%;
        display: flex;
        align-items: center;
        margin: 8px 0;
        
  
        .r-title {
          text-align: right;
          font-size: .92857143em;
          font-weight: 700;
          color: rgba(0, 0, 0, .87);
          width: 100px;
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
  
  .license-sel {
    width:100%;
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
    
  }
</style>
  