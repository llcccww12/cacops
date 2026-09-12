<template>
  <div class="model-select">
    <div class="title" v-if="showTitle"><span :class="required ? 'required' : ''">{{ selfTitle }}</span></div>
    <div class="content">
      <input type="hidden" name="model_id" :value="form_model_id" required />
      <input type="hidden" name="model_name" :value="form_model_name" required />
      <input type="hidden" name="model_version" :value="form_model_version" required />
      <input type="hidden" name="pre_train_model_url" :value="form_pre_train_model_url" required />
      <input type="hidden" name="ckpt_name" :value="form_ckpt_name" required />
      <div v-if="selectList.length" class="model-list-c" :class="errStatus ? 'error' : ''">
        <div class="model-item" v-for="(item, index) in selectList" :key="item.id"
          :title="item._modelName + '/' + item.name">
          {{ item._modelName + '/' + item.name }}{{ selectList.length > 1 ? ';' : '' }}
        </div>
      </div>
      <input v-if="selectList.length == 0" type="text" class="disabled" style="width:100%"
        :required="required ? 'required' : undefined" :placeholder="$t('modelObj.model_select_placeholder')" />
    </div>
    <div class="btn-select" @click="dlgShow = true">
      <i class="el-icon-plus"></i>
      <span>{{ $t('modelObj.model_select') }}</span>
    </div>
    <!-- <div style="display:flex;align-items:center;justify-content:center;margin-left:6px;">
      <el-tooltip placement="top" effect="light">
        <i class="question circle icon link" style="margin-top:-7px"></i>
        <div slot="content">
          <div style="width:200px;text-align:center;">{{ $t('modelObj.model_suport_file_tips') }}</div>
        </div>
      </el-tooltip>
    </div> -->
    <el-dialog class="model-dlg" :visible.sync="dlgShow" :title="$t('modelObj.model_select')" width="1000px" :modal="true"
      :close-on-click-modal="false" :show-close="true" :destroy-on-close="false" :before-close="beforeClose" @open="open"
      @closed="closed">
      <div class="dlg-content">
        <div class="left-area" v-loading="dlgLoading">
          <div class="model-tabs-c">
            <el-tabs class="model-tabs" v-model="dlgActiveName" @tab-click="dlgTabClick">
              <el-tab-pane :label="$t('modelObj.model_current_repo')" name="first"></el-tab-pane>
              <el-tab-pane :label="$t('modelObj.model_my')" name="second"></el-tab-pane>
              <el-tab-pane :label="$t('modelObj.model_public')" name="third"></el-tab-pane>
              <el-tab-pane :label="$t('modelObj.model_collected')" name="fourth"></el-tab-pane>
            </el-tabs>
            <el-input size="small" class="search-inp" :placeholder="$t('modelObj.model_search_placeholder')"
              v-model="dlgSearchValue" @keydown.enter.stop.native.prevent="inputSearch">
              <div slot="suffix" class="search-inp-icon" @click="inputSearch">
                <i class="el-icon-search"></i>
              </div>
            </el-input>
          </div>
          <el-tree :data="dlgModelTreeData" ref="dlgTreeRef" highlight-current show-checkbox node-key="id"
            :default-expanded-keys="dlgInitTreeNode" :props="dlgTreeProps" :index="10" accordion
            @check="onTreeCheckChange">
            <span slot-scope="{ node, data }" class="slot-wrap">
              <span v-if="data.parent" class="custom-tree-node">
                <el-tooltip v-if="data.description" placement="top-start">
                  <div slot="content" class="multiple-wrap"> {{ data.description }}</div>
                  <span class="model-title model-nowrap">
                    <div class="model_flex">
                      <span style="flex: inherit" class="model-nowrap">{{ node.label }}<span
                          v-if="(data.version != '0.0.1' || data.new != 1)" class="model-version">
                          {{ data.version }} </span>
                      </span>
                      <img v-if="data.recommend == 1" style="margin-left: 0.4rem" src="/img/jian.svg" />
                    </div>
                  </span>
                </el-tooltip>
                <span v-else class="model-title model-nowrap">
                  <div class="model_flex">
                    <span style="flex: inherit" class="model-nowrap">{{ node.label }}<span
                        v-if="(data.version != '0.0.1' || data.new != 1)" class="model-version">
                        {{ data.version }} </span>
                    </span>
                    <img v-if="data.recommend == 1" style="margin-left: 0.4rem" src="/img/jian.svg" />
                  </div>
                </span>
                <span class="model-repolink model-nowrap" @click.stop="return false;">
                  <a :href=" '/' + data.repoOwnerName + '/' + data.repoName + '/modelmanage/model_readme_tmpl?name=' + data.name "
                    target="_blank">
                    {{ data.repoOwnerName }}/{{ data.repoDisplayName }}
                  </a>
                </span>
              </span>
              <span v-else style="display: flex">
                <span class="model-nowrap" :title=" node.label ">
                  {{ node.label }}
                </span>
              </span>
            </span>
          </el-tree>
          <div class="pagination-c">
            <el-pagination background @current-change=" dlgPageChange " :current-page=" dlgPage "
              :page-size=" dlgPageSize " layout="total, prev, pager, next" :total=" dlgTotal ">
            </el-pagination>
          </div>
        </div>
        <div class="right-area">
          <div class="right-title"><span>{{ $t('modelObj.model_selected') }}</span></div>
          <div class="right-selected-list">
            <el-checkbox-group v-model=" dlgSelectedModel ">
              <el-checkbox v-for="(   item, index   ) in    dlgSelectedModelList   " :key=" item.id " :label=" item.id "
                :true-label=" item.id " :title=" item.name " @change=" (checked) => dlgChangeSelect(checked, item) ">
                {{ item.name }}
              </el-checkbox>
            </el-checkbox-group>
          </div>
          <div class="right-btn-c">
            <el-button type="primary" size="small" @click=" confirm ">{{ $t('modelObj.model_ok') }}</el-button>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>

// const supportCheckPointFileExt = ["ckpt", "pb", "h5", "json", "pkl", "pth", "t7", "pdparams", "onnx", "pbtxt", "keras", "mlmodel", "cfg", "pt"];
export default {
  name: "ModelSelect",
  props: {
    title: { type: String, default: '' },
    showTitle: { type: Boolean, default: true },
  },
  data() {
    return {
      maxCount: 1,
      required: false,
      userName: location.pathname.split('/')[1],
      repoName: location.pathname.split('/')[2],
      type: '1',

      form_model_id: '',
      form_model_name: '',
      form_model_version: '',
      form_pre_train_model_url: '',
      form_ckpt_name: '',

      selectList: [],
      dlgShow: false,
      dlgLoading: false,
      dlgActiveName: 'first',
      dlgSearchValue: '',

      dlgTreeProps: {
        children: "children",
        label: "name",
      },
      dlgModelTreeData: [],
      dlgInitTreeNode: [],

      dlgPage: 1,
      dlgPageSize: 10,
      dlgTotal: 0,

      dlgSelectedModel: [],
      dlgSelectedModelList: [],

      errStatus: false,
    };
  },
  computed: {
    selfTitle() {
      return this.title || this.$t('modelObj.model_label');
    }
  },
  watch: {
    oriData: function (val) { },
  },
  methods: {
    $t(str, object) {
      const strList = str.split('.');
      let obj = window.i18n || {};
      let rStr = str;
      for (let i = 0, iLen = strList.length; i < iLen; i++) {
        if (typeof obj[strList[i]] == 'object') {
          obj = obj[strList[i]];
        }
        if (typeof obj[strList[i]] == 'string') {
          rStr = obj[strList[i]];
          break;
        }
      }
      if (object) {
        for (let key in object) {
          rStr = rStr.replace(new RegExp('\{\\s*' + key + '\\s*\}', 'g'), object[key]);
        }
      }
      return rStr;
    },
    open() {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.dlgSelectedModel = [];
      this.dlgSelectedModelList = [];
      for (let i = 0, iLen = this.selectList.length; i < iLen; i++) {
        const item = this.selectList[i];
        this.dlgSelectedModel.push(item.id);
        this.dlgSelectedModelList.push({
          ...item,
          id: item.id,
          name: item.name,
        });
      }
      this.searchModelData();
    },
    beforeClose(done) {
      done();
    },
    closed() { },
    dlgTabClick(tab, event) {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.searchModelData();
    },
    transformTreeData(data) {
      for (let i = 0, iLen = data.length; i < iLen; i++) {
        const dataI = data[i];
        const _children = dataI.modelFileList || [];
        const children = [];
        dataI.parent = true;
        dataI.disabled = true;
        _children.forEach(item => {
          item.ModTimeNum = new Date(item.ModTime).getTime();
        })
        _children.sort((a, b) => a.FileName.localeCompare(b.FileName));
        _children.sort((a, b) => b.ModTimeNum - a.ModTimeNum);
        for (let j = 0, jLen = _children.length; j < jLen; j++) {
          const file = _children[j];
          if (file.IsDir) continue;
          // const arr = file.FileName.split('.');
          // if (!supportCheckPointFileExt.includes(arr[arr.length - 1])) continue;
          file._modelID = dataI.id;
          file._modelName = dataI.name;
          file._modelVersion = dataI.version;
          file._preTrainModelUrl = dataI.path;
          file.name = file.FileName;
          file.id = `${file._modelID}|${file._modelName}|${file._modelVersion}|${file._preTrainModelUrl}|${file.name}`;
          children.push(file);
        }
        dataI.children = children;
      }
      return data;
    },
    inputSearch() {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.searchModelData();
    },
    searchModelData() {
      const tabName = this.dlgActiveName;
      const tabPrams = {
        'first': 4,
        'second': 2,
        'third': 1,
        'fourth': 3,
      }
      const params = {
        type: this.type,
        queryType: tabPrams[tabName],
        repoOwnerName: this.userName,
        repoName: this.repoName,
        q: this.dlgSearchValue.trim(),
        page: this.dlgPage,
        needModelFile: true,
        notNeedEmpty: true,
      };
      if (params.queryType == 2 || params.queryType == 4) {
        params.orderBy = 'created_unix';
      }
      this.dlgLoading = true;
      this.$axios.get(`/modelsquare/main_query_data`, {
        params: params,
      }).then(res => {
        this.dlgLoading = false;
        const data = res.data?.data || [];
        this.dlgModelTreeData = this.transformTreeData(data);
        this.dlgInitTreeNode = this.dlgModelTreeData[0]?.id
          ? [this.dlgModelTreeData[0].id]
          : [];
        this.dlgTotal = parseInt(res.data?.count || 0);
        const setCheckedKeysList = this.dlgModelTreeData.reduce((pre, cur) => {
          cur.children.forEach((item) => {
            if (this.dlgSelectedModel.includes(item.id)) {
              pre.push(item.id);
            }
          });
          return pre;
        }, []);
        this.$refs.dlgTreeRef.setCheckedKeys(setCheckedKeysList);
      }).catch(err => {
        this.dlgLoading = false;
        console.log(err);
      });
    },
    dlgChangeSelect(checked, data) {
      this.$refs.dlgTreeRef.setChecked(data.id, false, false);
      const index = this.dlgSelectedModelList.findIndex((item) => {
        return item.id === data.id;
      });
      this.dlgSelectedModelList.splice(index, 1);
      this.dlgSelectedModel = this.dlgSelectedModelList.map(item => item.id);
    },
    dlgPageChange(page) {
      this.dlgPage = page;
      this.searchModelData();
    },
    onTreeCheckChange(data) {
      if (
        this.dlgSelectedModelList.length === 0 ||
        this.dlgSelectedModelList.every((item) => item.id !== data.id)
      ) {
        if (this.maxCount == 1) {
          this.dlgSelectedModelList.forEach((item) => {
            this.$refs.dlgTreeRef.setChecked(item.id, false, false);
          });
          this.dlgSelectedModelList.splice(0, Infinity);
          this.dlgSelectedModelList.push(data);
        } else {
          if (
            this.dlgSelectedModelList.some((item) => {
              return item._modelID !== data._modelID;
            })
          ) {
            this.$refs.dlgTreeRef.setChecked(data.id, false, false);
            this.$message.warning(this.$t('modelObj.model_should_same_model'));
          } else if (this.dlgSelectedModelList.length === this.maxCount) {
            this.$refs.dlgTreeRef.setChecked(data.id, false, false);
            this.$message.warning(this.$t('modelObj.model_most', { msg: this.maxCount }));
          } else {
            this.dlgSelectedModelList.push(data);
          }
        }
      } else {
        const index = this.dlgSelectedModelList.findIndex((item) => {
          return item.id === data.id;
        });
        this.dlgSelectedModelList.splice(index, 1);
      }
      this.dlgSelectedModel = this.dlgSelectedModelList.map((item) => {
        return item.id;
      });
    },
    confirm() {
      const len = this.dlgSelectedModelList.length;
      this.selectList.splice(0, Infinity);
      if (len) {
        const ckptNames = [];
        for (let i = 0; i < len; i++) {
          const item = this.dlgSelectedModelList[i];
          this.form_model_id = item._modelID;
          this.form_model_name = item._modelName;
          this.form_model_version = item._modelVersion;
          this.form_pre_train_model_url = item._preTrainModelUrl;
          ckptNames.push(item.name);
          this.selectList.push({ ...item });
        }
        this.form_ckpt_name = ckptNames.join(';');
      } else {
        this.form_model_id = '';
        this.form_model_name = '';
        this.form_model_version = '';
        this.form_pre_train_model_url = '';
        this.form_ckpt_name = '';
      }
      if (this.selectList.length) {
        this.errStatus = false;
      }
      this.dlgShow = false;
    },
    check() {
      if (!this.selectList.length) {
        this.errStatus = true;
        return false;
      }
      this.errStatus = false;
      return true;
    },
  },
  beforeMount() { },
  mounted() {
    const dataEl = this.$el.previousElementSibling;
    if (dataEl) {
      const dataset = dataEl.dataset;
      const multiple = dataset.multiple;
      if (multiple) {
        this.maxCount = Infinity;
      }
      if (dataset.modelId) {
        this.form_model_id = dataset.modelId;
        this.form_model_name = dataset.modelName;
        this.form_model_version = dataset.modelVersion;
        this.form_pre_train_model_url = dataset.preTrainModelUrl;
        this.form_ckpt_name = dataset.ckptName;
        const ckptNameList = dataset.ckptName.split(';');
        for (let i = 0, iLen = ckptNameList.length; i < iLen; i++) {
          const ckptName = ckptNameList[i];
          const file = {
            _modelID: dataset.modelId,
            _modelName: dataset.modelName,
            _modelVersion: dataset.modelVersion,
            _preTrainModelUrl: dataset.preTrainModelUrl,
            name: ckptName,
          };
          file.id = `${file._modelID}|${file._modelName}|${file._modelVersion}|${file._preTrainModelUrl}|${file.name}`;
          this.selectList.push(file);
        }
      }
      if (dataset.required) {
        this.required = true;
      }
    }
  },
};
</script>

<style scoped lang="less">
.model-select {
  display: flex;
  margin-bottom: 28px;

  .title {
    width: 130px;
    text-align: right;
    margin-right: 24px;
    color: #101010;
    font-size: 14px;
    align-items: center;
    display: flex;
    justify-content: flex-end;

    .required {
      position: relative;

      &::after {
        position: absolute;
        content: "*";
        top: -3px;
        right: -10px;
        color: red;
      }
    }
  }

  .content {
    display: inline-block;
    width: 48.5%;
    margin-right: 5px;

    .model-list-c {
      min-height: 36px;
      border-radius: 4px;
      border: 1px solid #DCDFE6;
      box-sizing: border-box;
      color: #606266;
      padding: 5px 15px;

      .model-item {
        line-height: 26px;
        font-size: 14px;
        color: rgba(0, 0, 0, 0.87);
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }

      .model-item-model {
        color: rgba(136, 136, 136, 1);
      }

      .model-item-file {
        padding-left: 5px;
      }

      .model-item-placeholder {
        height: 26px;
        line-height: 26px;
        color: rgba(0, 0, 0, 0.4);
        opacity: 0.45 !important;
        font-size: 14px;
      }

      &.error {
        color: #9f3a38;
        background: #fff6f6;
        border-color: #e0b4b4;
      }
    }
  }

  .btn-select {
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    color: rgb(3, 102, 214);

    i {
      margin-right: 4px;
    }
  }
}

.model-dlg {
  /deep/.el-dialog__body {
    padding-top: 0;
  }

  .dlg-content {
    display: flex;
    min-height: 300px;

    .left-area {
      flex: 1;
      margin-right: 15px;
      border-right: 1px solid rgb(245, 245, 246);
      padding-right: 15px;
      position: relative;
      overflow: hidden;

      .model-tabs-c {
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

      .pagination-c {
        margin-top: 25px;
        text-align: center;
      }
    }

    .right-area {
      width: 300px;
      padding-right: 30px;
      position: relative;


      .right-title {
        font-size: 14px;
        height: 40px;
        text-align: left;
        color: rgb(0, 102, 255);
        line-height: 40px;
      }

      .right-selected-list {
        margin: 14px 0;
      }

      .right-btn-c {
        text-align: right;
        position: absolute;
        bottom: 15px;
        width: 100%;
        padding-right: 30px;

        /deep/ .el-button {
          background: rgb(56, 158, 13);
          color: rgb(255, 255, 255);
          border: 1px solid rgb(56, 158, 13);
        }
      }
    }
  }
}

.el-tree {
  max-height: 400px;
  overflow-y: auto;
  overflow-x: hidden;
  position: relative;
  cursor: default;
  background: #fff;
  color: #606266;
  font-family: SourceHanSansSC-medium;
}

.custom-tree-node {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.custom-tree-node .model-title {
  font-size: 14px;
  color: #101010;
  font-weight: 600;
  flex: 1;
}

.custom-tree-node .model-repolink {
  flex: 1;
  text-align: right;
  font-size: 12px;
}

.el-tree /deep/ .el-tree-node__content {
  height: 40px;
  background-color: #f5f5f6;
}

.el-tree /deep/ .el-tree-node__children .el-tree-node__content {
  height: 30px;
  background-color: #fff;
  line-height: 20px;
  font-size: 12px;
}

/deep/ .el-checkbox-group .el-checkbox {
  max-width: 100%;
  min-width: 80%;
}

/deep/ .el-checkbox-group .el-checkbox .el-checkbox__label {
  max-width: 100%;
  overflow: hidden;
  vertical-align: middle;
  text-overflow: ellipsis;
}

.model-nowrap {
  overflow: hidden;
  text-overflow: ellipsis;
}

.slot-wrap {
  flex: 1;
  padding-right: 2rem;
  max-width: 93%;
}

.multiple-wrap {
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  display: -webkit-box;
  max-width: 400px;
  overflow: hidden;
}

.unzip-failed {
  margin-left: 1rem;
  color: red;
}

.zip-loading {
  margin-left: 1rem;
  color: #fcca00;
}

.model-search-vue {
  z-index: 9999;
  position: absolute;
  right: 31%;
  height: 30px;
  top: 6px;
}

.select-model-label {
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-right: 1rem;
  white-space: nowrap;
}

.model_flex {
  display: flex;
  align-items: center;
}

.model-version {
  margin-left: 4px;
  border-radius: 4px;
  color: rgba(16, 16, 16, 0.8);
  border-radius: 4px;
  font-size: 12px;
  background: rgba(220, 220, 220, 0.8);
  padding: 1px 3px;
}
</style>
