<template>
  <div class="form-row">
    <div class="left-area">
      <div class="title" v-if="showTitle">
        <span :class="required ? 'required' : ''"> {{ $t('cloudbrainObj.dataset') }}</span>
      </div>
      <div class="content">
        <div class="model-list-c" :class="errStatus ? 'error' : ''">
          <div class="model-item" v-for="(item) in selectList" :key="item.id" :title="item.alias">
            {{ item.alias }}{{ multiple ? ';' : '' }}
          </div>
          <div v-if="selectList.length == 0" class="model-item-placeholder">
            {{ $t('datasetObj.dataset_select_placeholder') }}
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
    <el-dialog class="model-dlg" :visible.sync="dlgShow" width="1100px"
      :modal="true" :close-on-click-modal="false" :show-close="true" :destroy-on-close="false"
      :before-close="beforeClose" @open="open" @closed="closed">
      <div class="dlg-content">
        <div class="dlg-left-area" v-loading="dlgLoading">
          <div class="header-c">
            <span class="header-text">{{$t('datasetObj.dataset_select')}}</span>
            <el-input size="small" class="search-inp" :placeholder="$t('datasetObj.dataset_search_placeholder')"
              v-model="dlgSearchValue" @keydown.enter.stop.native.prevent="inputSearch">
              <div slot="suffix" class="search-inp-icon" @click="inputSearch">
                <i class="el-icon-search"></i>
              </div>
            </el-input>
          </div>
          <div class="model-tabs-c">
            <el-tabs class="model-tabs" v-model="dlgActiveName" @tab-click="dlgTabClick">
              <el-tab-pane :label="$t('cloudbrainObj.all')" name="first"></el-tab-pane>
              <el-tab-pane :label="$t('cloudbrainObj.ihave')" name="second"></el-tab-pane>
              <el-tab-pane :label="$t('cloudbrainObj.iCollaborate')" name="third"></el-tab-pane>
              <el-tab-pane :label="$t('cloudbrainObj.iCollect')" name="fourth"></el-tab-pane>
            </el-tabs>
          </div>
          <el-tree :data="dlgModelTreeData" ref="dlgTreeRef" highlight-current show-checkbox node-key="id"
            :props="dlgTreeProps" @check="onTreeCheckChange">
            <span slot-scope="{ node, data }" class="slot-wrap">
              <span class="item-wrap">
                <div class="item-c nowrap">
                  <span class="item nowrap" style="max-width: 93%">
                    <Icons style="flex-shrink: 0;width:16px;height:16px;" type="dataset" :isPrivate="data.is_private"></Icons>
                    <span tyle="color: rgba(16, 16, 16, 0.8);"> {{ data.owner_name }} / </span>
                    <span class="model-nowrap" :title="node.label">
                      {{ node.label }}
                    </span>
                  </span>
                  <svg v-if="data.recommend" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20"><defs></defs><g><path fill="#FF6200" d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z"></path></g></svg>
                  <span class="exceed-size-tips" v-if="useExceedSize && exceedSize && data.size > exceedSize">
                      {{ $t('datasetObj.dataset_exceeds_failed') }}{{ exceedSize / (1024 * 1024 * 1024) }}G
                    </span>
                </div>
                <a :href="`/datasets/detail/${data.owner_name}/${data.name}`" target="_blank">
                  <svg width="16" height="16" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M28 6H42V20" stroke="#10101080" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M42 29.4737V39C42 40.6569 40.6569 42 39 42H9C7.34315 42 6 40.6569 6 39V9C6 7.34315 7.34315 6 9 6L18 6" stroke="#10101080" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M25.7998 22.1999L41.0998 6.8999" stroke="#10101080" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
                </a>
              </span>
            </span>
          </el-tree>
          <div class="pagination-c">
            <el-pagination background @current-change="dlgPageChange" :current-page="dlgPage" :page-size="dlgPageSize"
              layout="total, prev, pager, next" :total="dlgTotal">
            </el-pagination>
          </div>
        </div>
        <div class="dlg-right-area">
          <div class="right-title"><span>{{ $t('datasetObj.dataset_selected') }}</span>
          </div>
          <div class="right-selected-list">
            <el-checkbox-group v-model="dlgSelectedModel">
              <el-checkbox v-for="(item) in dlgSelectedModelList " :key="item.id" :label="item.id" :true-label="item.id"
                :title="item.alias" @change="(checked) => dlgChangeSelect(checked, item)">
                <span style="color: rgba(16, 16, 16, 0.8);">{{ item.owner_name }} / </span><span class="model-nowrap">{{ item.alias }}</span>
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
import { getUrlSearchParams } from '~/utils';
import Icons from '~/components/square/Icons.vue';
export default {
  name: "DatasetSelectV2",
  props: {
    value: { type: Array, required: true },
    required: { type: Boolean, default: false },
    multiple: { type: Boolean, default: false },
    title: { type: String, default: '' },
    showTitle: { type: Boolean, default: true },
    useExceedSize: { type: Boolean, default: false, },
    exceedSize: { type: Number, default: 0 },
    maxCount: { type: Number, default: 0 },
  },
  components: { Icons },
  data() {
    return {
      type: '1',
      selectList: [],
      dlgShow: false,
      dlgLoading: false,
      dlgActiveName: 'first',
      dlgSearchValue: '',

      dlgTreeProps: {
        children: "children",
        label: "alias",
      },
      dlgModelTreeData: [],

      dlgPage: 1,
      dlgPageSize: 10,
      dlgTotal: 0,

      dlgSelectedModel: [],
      dlgSelectedModelList: [],

      errStatus: false,
    };
  },
  watch: {
    value: {
      immediate: true,
      deep: true,
      handler(newVal) {
        newVal = newVal === undefined ? [] : newVal;
        this.selectList = [...newVal];
      }
    },
    dlgSearchValue: {
      handler(newVal) {
        if (newVal === '') {
          this.searchModelData();
        }
      }
    }
  },
  methods: {
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
    inputSearch() {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.searchModelData();
    },
    searchModelData() {
      const tabName = this.dlgActiveName;
      const tabPrams = {
        'first': '/accessible',
        'second': '/owned',
        'third': '/collaborated',
        'fourth': '/collected',
      }
      const params = {
        url: tabPrams[tabName],
        q: this.dlgSearchValue.trim(),
        page: this.dlgPage,
        page_size: this.dlgPageSize,
      };
      params.owner_type = tabName==='second' ?  'individual' : ''
      this.dlgLoading = true;
      getDatasets(params).then(res => {
        this.dlgLoading = false;
        if(res.data.code === 0){
            const data = res.data?.data?.datasets || [];
            this.dlgModelTreeData = data.map(item => { 
              item.disabled = false
              if(this.useExceedSize && item.size > this.exceedSize && this.exceedSize){
                item.disabled = true
              }
              return item
            })
            this.dlgTotal = res.data?.data?.total || 0;
            const setCheckedKeysList = this.dlgModelTreeData.reduce((pre, cur) => {
            if (this.dlgSelectedModel.includes(cur.id)) {
                pre.push(cur.id);
            }
            return pre;
            }, []);
            this.$refs.dlgTreeRef.setCheckedKeys(setCheckedKeysList);
        }else{
            this.$message.error(rs.data.msg || '获取数据集失败')
            console.log(err);
        }
        
      }).catch(err => {
        this.dlgLoading = false;
        console.log(err);
      });
    },
    dlgChangeSelect(checked, data) {
      const index = this.dlgSelectedModelList.findIndex((item) => {
        return item.id === data.id;
      });
      this.dlgSelectedModelList.splice(index, 1);
      this.dlgSelectedModel = this.dlgSelectedModelList.map(item => item.id);
      this.$refs.dlgTreeRef.setCheckedKeys(this.dlgSelectedModel);
    },
    dlgPageChange(page) {
      this.dlgPage = page;
      this.searchModelData();
    },
    onTreeCheckChange(data) {
        if (this.multiple) {
          if (this.dlgSelectedModel.indexOf(data.id) >= 0) {
            const index = this.dlgSelectedModelList.findIndex(item => item.id == data.id);
            this.dlgSelectedModelList.splice(index, 1);
          } else {
            if (this.dlgSelectedModel.length >= this.maxCount) {
              console.log('maxCount');
              this.$message.warning(this.$t('datasetObj.dataset_most', { msg: this.maxCount }));
            } else if (this.dlgSelectedModelList.find(item => item.name === data.name)) {
              this.$message.warning(this.$t('datasetObj.dataset_not_equal_file'));
            } else {
              const curSize = data.size;
              const selectedSize = this.dlgSelectedModelList.reduce((pre, _data) => {
                return pre + _data.size;
              }, 0);
              if (this.useExceedSize && this.exceedSize && (curSize + selectedSize) > this.exceedSize) {
                this.$message.warning(this.$t('datasetObj.dataset_exceeds_failed') + `${(this.exceedSize) / (1024 * 1024 * 1024)}G`);
              } else {
                this.dlgSelectedModelList.push(data);
              }
            }
          }
          this.dlgSelectedModel = this.dlgSelectedModelList.map(item => item.id);
        } else {
          if (this.dlgSelectedModel.indexOf(data.id) >= 0) {
            this.dlgSelectedModelList = [];
            this.dlgSelectedModel = [];
          } else {
            this.dlgSelectedModelList = [data];
            this.dlgSelectedModel = this.dlgSelectedModelList.map(item => item.id);
          }
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
      this.dlgShow = false;
      this.$emit('input', this.selectList);
      this.$emit('change', this.selectList);
    },
    check() {
      if (this.required && !this.selectList.length) {
        this.errStatus = true;
        return false;
      }
      this.errStatus = false;
      return true;
    },
  },
  beforeMount() { },
  mounted() {
    const urlParams = getUrlSearchParams();
  },
};
</script>

<style scoped lang="less">
@import 'cloudbrain.less';

.form-row {
  .content {
    .model-list-c {
      min-height: 37.6px;
      border-radius: 4px;
      background: rgb(245, 245, 245);
      box-sizing: border-box;
      color: rgba(0, 0, 0, 0.95);
      padding: 4px 15px;
      display: flex;
      flex-wrap: wrap;
      gap: 4px;
      .model-item {
        line-height: 26px;
        font-size: 14px;
        color: rgba(0, 0, 0, 0.87);
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

  // .right-area {
  //   display: flex;
  //   align-items: center;

  //   .btn-select {
  //     display: flex;
  //     justify-content: center;
  //     align-items: center;
  //     cursor: pointer;
  //     color: rgb(3, 102, 214);

  //     i {
  //       margin-right: 4px;
  //     }
  //   }
  // }
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
      .header-c{
        display: flex;
        align-items: center;
        height: 40px;
        .header-text{
          font-size: 18px;
          font-weight: 700;
          line-height: 20px;
          color: #101010;
        }
        .search-inp {
          overflow: hidden;
          width: 246px;
          z-index: 5;
          position: relative;
          margin-left: 16px;
          /deep/ .el-input__inner{
            border: none;
            background: #f5f5f5;
          }
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
      .model-tabs-c {
        display: flex;
        align-items: center;

        .model-tabs {
          flex: 1;
          overflow: hidden;
          margin-right: 5px;
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
        color: rgba(0,102,255,1);
        line-height: 40px;
        font-size: 18px;
        font-weight: 700;
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
.el-tree /deep/ .is-checked .el-tree-node__content .model-icon:not([stroke]){
  fill: #101010 !important;
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
  .item-wrap { 
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    .item-c{
      color: rgba(16, 16, 16, 0.8);
      display: flex;
      align-items: center;
      width: 100%;
      .item{
        display: flex;
        align-items: center;
        svg{
          margin-right: 10px;
        }
      }
    }
  }
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
