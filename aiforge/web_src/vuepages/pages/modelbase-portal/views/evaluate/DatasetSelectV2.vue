<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title" v-if="showTitle">
        <span :class="required ? 'required' : ''"> {{ $t('cloudbrainObj.dataset') }}</span>
      </div>
      <div class="content">
        <div class="model-list-c" :class="errStatus ? 'error' : ''">
          <div class="model-item" v-for="(item) in selectList" :key="item.name" :title="item.name">
            <span>{{ item.name }}</span>
            <svg @click="deleteItem(item)" xmlns="http://www.w3.org/2000/svg" class="fill-close" viewBox="0 0 32 32" width="14" height="14" fill="rgb(184, 188, 197);"><defs></defs><g><path d="M16 30.667c-8.1 0-14.667-6.566-14.667-14.667s6.566-14.667 14.667-14.667 14.667 6.566 14.667 14.667-6.566 14.667-14.667 14.667zM21.492 20.238l-4.267-4.267 4.267-4.267-1.254-1.254-4.267 4.267-4.267-4.267-1.254 1.254 4.267 4.267-4.267 4.267 1.254 1.254 4.267-4.267 4.267 4.267 1.254-1.254z"></path></g></svg>
          </div>
          <div v-if="selectList.length == 0" class="model-item-placeholder">
            {{ $t('datasetObj.eval_select_placeholder') }}
          </div>
        </div>
      </div>
    </div>
    <div class="right-area">
      <div class="btn-select" @click="dlgShow = true">
        <i class="el-icon-plus"></i>
        <span>{{ $t('datasetObj.dataset_select') }}</span>
      </div>
    </div>
    <el-dialog class="model-dlg" :visible.sync="dlgShow" :title="$t('datasetObj.dataset_select')" width="1100px"
      :modal="true" :close-on-click-modal="false" :show-close="true" :destroy-on-close="false"
      :before-close="beforeClose" @open="open" @closed="closed">
      <div class="dlg-content">
        <div class="dlg-left-area">
          <div class="model-tabs-c">
            <el-tabs class="model-tabs" v-model="dlgActiveName" @tab-click="dlgTabClick">
              <el-tab-pane v-for="item in tabList" :label="$t('evalCaterary.'+ item)" :name="item" :key="item"></el-tab-pane>
            </el-tabs>
          </div>
          <el-tree :data="dlgModelTreeData" ref="dlgTreeRef" highlight-current show-checkbox node-key="name"
            :props="dlgTreeProps" @check="onTreeCheckChange">
            <span slot-scope="{ node, data }" class="slot-wrap">
              <span style="display: flex;justify-content: space-between;align-items: center;width: 100%;">
                <div class="nowrap" style="color: rgba(16, 16, 16, 0.8);display: flex;align-items: center;width: 100%;">
                  <span class="nowrap" style="max-width: 93%">
                    <span class="model-nowrap" :title="data.name">
                      {{ data.name }}
                    </span>
                  </span>
                </div>
                <a :href="`/datasets/detail/${data.url}`" target="_blank">
                  <svg width="16" height="16" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M28 6H42V20" stroke="#10101080" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M42 29.4737V39C42 40.6569 40.6569 42 39 42H9C7.34315 42 6 40.6569 6 39V9C6 7.34315 7.34315 6 9 6L18 6" stroke="#10101080" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M25.7998 22.1999L41.0998 6.8999" stroke="#10101080" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
                </a>
              </span>
            </span>
          </el-tree>
          <div class="pagination-c">
            <el-pagination background @current-change="handlePageChange" :current-page="dlgPage" :page-size="dlgPageSize"
              layout="total, prev, pager, next" :total="dlgTotal">
            </el-pagination>
          </div>
        </div>
        <div class="dlg-right-area">
          <div class="right-title"><span>{{ $t('datasetObj.dataset_selected') }}</span>
          </div>
          <div class="right-selected-list">
            <el-checkbox-group v-model="dlgSelectedModel">
              <el-checkbox v-for="(item) in dlgSelectedModelList " :key="item.name" :label="item.name" :true-label="item.name"
                :title="item.name" @change="(checked) => dlgChangeSelect(checked, item)">
                <span class="model-nowrap">{{ item.name }}</span>
              </el-checkbox>
            </el-checkbox-group>
          </div>
          <div class="right-btn-c">
            <el-button type="default" size="small" @click="clearSelect">{{ $t('clear') }}</el-button>
            <el-button type="primary" size="small" @click="confirm">{{ $t('datasetObj.dataset_ok') }}</el-button>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getDatasets } from "~/apis/modules/dataset";
export default {
  name: "ModelSelect",
  props: {
    datasetList: { type: Object, required: true },
    required: { type: Boolean, default: false },
    multiple: { type: Boolean, default: true },
    title: { type: String, default: '' },
    showTitle: { type: Boolean, default: true },
  },
  data() {
    return {

      selectList: [],
      dlgShow: false,

      tabList: [],
      dlgActiveName: '',

      dlgTreeProps: {
        children: "children",
        label: "name",
      },
      dlgModelTreeData: [],

      dlgPage: 1,
      dlgPageSize: 10,
      dlgTotal: 0,

      dlgSelectedModel: [],
      dlgSelectedModelList: [],

      errStatus: false,
      maxCount: 5
    };
  },
  watch: {
    
  },
  methods: {
    open() {
      console.log("open",this.datasetList)
      // this.dlgTotal = 0;
      // this.dlgPage = 1;
      // this.dlgSelectedModel = [];
      // this.dlgSelectedModelList = [];
      // for (let i = 0, iLen = this.selectList.length; i < iLen; i++) {
      //   const item = this.selectList[i];
      //   this.dlgSelectedModel.push(item.id);
      //   this.dlgSelectedModelList.push({
      //     ...item,
      //     id: item.id,
      //     name: item.name,
      //   });
      // }
    },
    beforeClose(done) {
      done();
    },
    closed() { 
    },
    dlgTabClick(tab, event){
      this.allDlgModelTreeData = this.datasetList[this.dlgActiveName]
      this.handlePageChange()
      this.dlgSelectedModel = this.dlgSelectedModelList.map(item => item.name);
      this.$refs.dlgTreeRef.setCheckedKeys(this.dlgSelectedModel);
    },
    dlgChangeSelect(checked, data) {
      const index = this.dlgSelectedModelList.findIndex((item) => {
        return item.name === data.name;
      });
      this.dlgSelectedModelList.splice(index, 1);
      this.dlgSelectedModel = this.dlgSelectedModelList.map(item => item.name);
      this.$refs.dlgTreeRef.setCheckedKeys(this.dlgSelectedModel);
    },
    onTreeCheckChange(data) {
      if (this.multiple) {
        if (this.dlgSelectedModel.indexOf(data.name) >= 0) {
          const index = this.dlgSelectedModelList.findIndex(item => item.name == data.name);
          this.dlgSelectedModelList.splice(index, 1);
        } else {
          if (this.dlgSelectedModel.length >= this.maxCount) {
            this.$message.warning(this.$t('datasetObj.dataset_most', { msg: this.maxCount }));
          }else{
            this.dlgSelectedModelList.push(data);
          }
          
        }
        this.dlgSelectedModel = this.dlgSelectedModelList.map(item => item.name);
      }
      this.$refs.dlgTreeRef.setCheckedKeys(this.dlgSelectedModel);
    },
    clearSelect() {
      this.dlgSelectedModelList = [];
      this.dlgSelectedModel = this.dlgSelectedModelList.map((item) => {
        return item.id;
      });
      this.$refs.dlgTreeRef.setCheckedKeys([], true);
    },
    confirm() {
      const len = this.dlgSelectedModelList.length;
      this.selectList.splice(0, Infinity);
      if (len) {
        for (let i = 0; i < len; i++) {
          const item = this.dlgSelectedModelList[i];
          this.selectList.push({ ...item });
        }
      }
      if (this.selectList.length) {
        this.errStatus = false;
      }
      console.log("confirm",this.selectList)
      this.dlgShow = false;
      this.$emit('input', this.selectList);
      this.$emit('change', this.selectList);
    },
    deleteItem(item){
      this.dlgChangeSelect(false,item)
      this.confirm()
    },
    check() {
      if (this.required && !this.selectList.length) {
        this.errStatus = true;
        return false;
      }
      this.errStatus = false;
      return true;
    },
    tranformDataset(data){
      this.tabList = Object.keys(data)
      if(this.tabList.length){
        this.dlgActiveName = this.tabList[0]
        this.allDlgModelTreeData = data[this.dlgActiveName]
        this.handlePageChange()
      }
    },
    handlePageChange(currentPage) {
      this.dlgPage = currentPage || 1;
      const startIndex = (this.dlgPage - 1) * this.dlgPageSize;
      const endIndex = this.dlgPage * this.dlgPageSize;
      // 从完整数据中截取当前页的数据
      this.dlgModelTreeData = this.allDlgModelTreeData.slice(startIndex, endIndex);
      // 更新总条数
      this.dlgTotal = this.allDlgModelTreeData.length;
    },
  },
  beforeMount() { },
  mounted() {
    this.tranformDataset(this.datasetList)
  },
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';

.form-row {
  .content {
    .model-list-c {
      min-height: 37.6px;
      border-radius: 4px;
      border: 1px solid #DCDFE6;
      box-sizing: border-box;
      color: #606266;
      padding: 4px 8px;
      background-color: rgb(255, 255, 255);
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      gap: 6px;
      .model-item {
        height: 26px;
        border-radius: 3px;
        background-color: rgba(238,240,243,1);
        color: rgba(16, 16, 16, 1);
        font-size: 12px;
        display: flex;
        align-items: center;
        padding: 0 6px;
        gap: 4px;

        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        .fill-close{
          cursor: pointer;
        }
        .fill-close:not([stroke]) {
            fill: rgb(184, 188, 197);
        }
      }

      .model-item-model {
        color: rgba(136, 136, 136, 1);
      }

      .model-item-file {
        padding-left: 5px;
      }

      .model-item-placeholder {
        height: 27px;
        line-height: 27px;
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

    .model-item-placeholder {
      line-height: 32px;
      color: rgba(0, 0, 0, 0.6);
      opacity: 0.45 !important;
      font-size: 13px;

    }
  }

  .right-area {
    display: flex;
    align-items: center;

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
}

.model-dlg {
  /deep/.el-dialog__body {
    padding-top: 0;
  }

  .dlg-content {
    display: flex;
    min-height: 300px;

    .dlg-left-area {
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

    .dlg-right-area {
      width: 300px;
      padding-right: 30px;
      position: relative;
      display: flex;
      flex-direction: column;


      .right-title {
        font-size: 14px;
        height: 40px;
        text-align: left;
        color: rgb(0, 102, 255);
        line-height: 40px;
      }

      .right-selected-list {
        margin: 14px 0;
        overflow-y: auto;
        flex: 1;
        height: 0;
        max-height: 390px;
      }

      .right-btn-c {
        width: 100%;
        display: flex;
        justify-content: space-between;
        height: 32px;
        margin-bottom: 10px;

        /deep/ .el-button--default {
          color: #409EFF;
          border: 1px solid #409EFF;
        }

        /deep/ .el-button--primary {
          background-color: rgba(25,144,255,1);
          color: rgb(255, 255, 255);
          border: 1px solid rgba(25,144,255,1);
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
  border-bottom: 1px solid rgba(16, 16, 16, 0.1);
}
.el-tree /deep/ .is-checked>.el-tree-node__content{
  background-color: #f7f7f7;
}
.el-tree /deep/ .is-checked .el-tree-node__content .model-nowrap{
  color: rgba(0, 102, 255, 1);
  font-weight: 700;
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
  max-width: 95%;
  overflow: hidden;
  vertical-align: middle;
  text-overflow: ellipsis;
  color: rgba(16, 16, 16, 1);
}

.model-nowrap {
  overflow: hidden;
  text-overflow: ellipsis;
  color: rgba(16, 16, 16, 1);
  font-weight: 700;
  margin-right: 10px;
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

.exceed-size-tips {
  margin-left: 1rem;
  font-size: 12px;
  color: red;
}
</style>
