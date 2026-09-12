<template>
  <div class="dataset-select">
    <div class="title" v-if="showTitle"><span :class="required ? 'required' : ''">{{ selfTitle }}</span></div>
    <div class="content">
      <div class="dataset-list-c" :class="errStatus ? 'error' : ''">
        <div class="dataset-item" v-for="(item) in selectList" :key="item.id">
          {{ item.name }};
        </div>
        <div v-if="selectList.length == 0" class="dataset-item-placeholder">
          {{ $t('datasetObj.dataset_select_placeholder') }}
        </div>
      </div>
      <div class="btn-select" @click="dlgShow = true">
        <i class="el-icon-plus"></i>
        <span>{{ $t('datasetObj.dataset_select') }}</span>
      </div>
    </div>
    <el-dialog class="dataset-dlg" :visible.sync="dlgShow" :title="$t('datasetObj.dataset_select')" width="1000px"
      :modal="true" :close-on-click-modal="false" :show-close="true" :destroy-on-close="false" :before-close="beforeClose"
      @open="open" @closed="closed">
      <div class="dlg-content">
        <div class="left-area" v-loading="dlgLoading">
          <div class="dataset-tabs-c">
            <el-tabs class="dataset-tabs" v-model="dlgActiveName" @tab-click="dlgTabClick">
              <!-- <el-tab-pane :label="$t('datasetObj.dataset_current_repo')" name="first"></el-tab-pane> -->
              <el-tab-pane :label="$t('datasetObj.dataset_my_upload')" name="second"></el-tab-pane>
              <el-tab-pane :label="$t('datasetObj.dataset_public')" name="third"></el-tab-pane>
              <el-tab-pane :label="$t('datasetObj.dataset_collected')" name="fourth"></el-tab-pane>
            </el-tabs>
            <el-input class="search-inp" :placeholder="$t('datasetObj.dataset_search_placeholder')"
              v-model="dlgSearchValue" @keyup.enter.native="inputSearch">
              <div slot="suffix" class="search-inp-icon" @click="inputSearch">
                <i class="el-icon-search"></i>
              </div>
            </el-input>
          </div>
          <el-tree :data="dlgDatasetTreeData" ref="dlgTreeRef" highlight-current show-checkbox node-key="id"
            :default-expanded-keys="dlgInitTreeNode" :props="dlgTreeProps" :index="10" accordion
            @check="onTreeCheckChange">
            <span slot-scope="{ node, data }" class="slot-wrap">
              <span v-if="data.parent" class="custom-tree-node">
                <el-tooltip v-if="data.Description" placement="top-start">
                  <div slot="content" class="multiple-wrap"> {{ data.Description }}</div>
                  <span class="dataset-title dataset-nowrap">
                    <div class="dataset_flex">
                      <span style="flex: inherit" class="dataset-nowrap">{{ node.label }}</span>
                      <img v-if="data.Recommend" style="margin-left: 0.4rem" src="/img/jian.svg" />
                    </div>
                  </span>
                </el-tooltip>
                <span v-else class="dataset-title dataset-nowrap">
                  <div class="dataset_flex">
                    <span style="flex: inherit" class="dataset-nowrap">{{ node.label }}</span>
                    <img v-if="data.Recommend" style="margin-left: 0.4rem" src="/img/jian.svg" />
                  </div>
                </span>
                <span class="dataset-repolink dataset-nowrap" @click.stop="return false;">
                  <i class="ri-links-line" style="color: #21ba45; margin-right: 0.3rem"
                    :title="$t('datasetObj.dataset_relate')"
                    v-if="dlgActiveName == 'first' && '/' + data.Repo.OwnerName + '/' + data.Repo.Name !== '/' + userName + '/' + repoName"></i>
                  <a :href="'/' + data.Repo.OwnerName + '/' + data.Repo.Name + '/datasets'" target="_blank"
                    :title="data.Repo.OwnerName + '/' + data.Repo.Alias">
                    {{ data.Repo.OwnerName }}/{{ data.Repo.Alias }}
                  </a>
                </span>
              </span>
              <span v-else style="display: flex">
                <span class="dataset-nowrap" :title="node.label">
                  {{ node.label }}
                </span>
                <span class="zip-loading" v-if="data.DecompressState == 2">
                  {{ $t('datasetObj.dataset_unziping') }}
                </span>
                <span class="unzip-failed" v-if="data.DecompressState == 3">
                  {{ $t('datasetObj.dataset_unzip_failed') }}
                </span>
                <span class="unzip-failed" v-if="exceedSize && data.Size > exceedSize">
                  {{ $t('datasetObj.dataset_exceeds_failed') }}{{ exceedSize / (1024 * 1024 * 1024) }}G
                </span>
              </span>
            </span>
          </el-tree>
          <div class="pagination-c">
            <el-pagination background @current-change="dlgPageChange" :current-page="dlgPage" :page-size="dlgPageSize"
              layout="total, prev, pager, next" :total="dlgTotal">
            </el-pagination>
          </div>
        </div>
        <div class="right-area">
          <div class="right-title"><span>{{ $t('datasetObj.dataset_selected') }}</span></div>
          <div class="right-selected-list">
            <el-checkbox-group v-model="dlgSelectedDataset">
              <el-checkbox v-for="(item) in dlgSelectedDatasetList" :key="item.id" :label="item.id"
                :true-label="item.id" :title="item.label" @change="(checked) => dlgChangeSelect(checked, item)">
                {{ item.label }}
              </el-checkbox>
            </el-checkbox-group>
          </div>
          <div class="right-btn-c">
            <el-button type="primary" @click="confirm">{{ $t('datasetObj.dataset_ok') }}</el-button>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getCurrentRepoDataset, getMyUploadedDataset, getPulicDataset, getMyFavoriteDataset } from '~/apis/modules/dataset';

const apiMap = {
  // 'first': getCurrentRepoDataset,
  'second': getMyUploadedDataset,
  'third': getPulicDataset,
  'fourth': getMyFavoriteDataset,
};

export default {
  name: "DatasetSelect",
  props: {
    title: { type: String, default: '' },
    showTitle: { type: Boolean, default: true },
    required: { type: Boolean, default: true },
    type: { type: Number, default: 0 },
    maxCount: { type: Number, default: 5 },
    userName: { type: String, default: '' },
    repoName: { type: String, default: '' },
    oriData: { type: Array, default: () => [] },
  },
  data() {
    return {
      selectList: [],
      dlgShow: false,
      dlgLoading: false,
      dlgActiveName: 'second',
      dlgSearchValue: '',

      dlgTreeProps: {
        children: "Attachments",
        label: "label",
      },
      dlgDatasetTreeData: [],
      dlgInitTreeNode: [],

      dlgPage: 1,
      dlgPageSize: 5,
      dlgTotal: 0,

      dlgSelectedDataset: [],
      dlgSelectedDatasetList: [],

      exceedSize: 0,

      errStatus: false,
    };
  },
  computed: {
    selfTitle() {
      return this.title || this.$t('dataset_label');
    }
  },
  watch: {
    oriData: function (val) {

    },
  },
  methods: {
    open() {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.dlgSelectedDataset = [];
      this.dlgSelectedDatasetList = [];
      for (let i = 0, iLen = this.selectList.length; i < iLen; i++) {
        const item = this.selectList[i];
        this.dlgSelectedDataset.push(item.id);
        this.dlgSelectedDatasetList.push({
          id: item.id,
          label: item.name,
        });
      }
      this.searchDatasetData();
    },
    beforeClose(done) {
      done();
    },
    closed() { },
    dlgTabClick(tab, event) {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.searchDatasetData();
    },
    transformTreeData(data) {
      return data.reduce((preParent, curParent) => {
        curParent.id = curParent.ID;
        curParent.disabled = true;
        curParent.parent = true;
        curParent.label = curParent.Title;
        const childrenData = curParent.Attachments.reduce(
          (preChild, curchild) => {
            curchild.id = curchild.UUID;
            if (curchild.DecompressState !== 1) {
              curchild.disabled = true;
            }
            if (curchild.Size > this.exceedSize && this.exceedSize) {
              curchild.disabled = true;
            }
            curchild.ref = 'dlgTreeRef';
            curchild.label = curchild.Name;
            preChild.push(curchild);
            return preChild;
          },
          []
        );
        preParent.push(curParent);
        return preParent;
      }, []);
    },
    inputSearch() {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.searchDatasetData();
    },
    searchDatasetData() {
      const tabName = this.dlgActiveName;
      const getApi = apiMap[tabName];
      this.dlgLoading = true;
      getApi({
        userName: this.userName,
        repoName: this.repoName,
        q: this.dlgSearchValue.trim(),
        type: this.type,
        page: this.dlgPage,
      }).then(res => {
        this.dlgLoading = false;
        const data = JSON.parse(res.data.data);
        this.dlgDatasetTreeData = this.transformTreeData(data);
        this.dlgInitTreeNode = this.dlgDatasetTreeData[0]?.id
          ? [this.dlgDatasetTreeData[0].id]
          : [];
        this.dlgTotal = parseInt(res.data.count);
        const setCheckedKeysList = this.dlgDatasetTreeData.reduce((pre, cur) => {
          cur.Attachments.forEach((item) => {
            if (this.dlgSelectedDataset.includes(item.id)) {
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
      const index = this.dlgSelectedDatasetList.findIndex((item) => {
        return item.id === data.id;
      });
      this.dlgSelectedDatasetList.splice(index, 1);
      this.dlgSelectedDataset = this.dlgSelectedDatasetList.map(item => item.id);
    },
    dlgPageChange(page) {
      this.dlgPage = page;
      this.searchDatasetData();
    },
    onTreeCheckChange(data) {
      if (
        this.dlgSelectedDatasetList.length === 0 ||
        this.dlgSelectedDatasetList.every((item) => item.id !== data.id)
      ) {
        if (
          this.dlgSelectedDatasetList.some((item) => {
            const itemIndex = item.label.lastIndexOf('.')
            const dataIndex = data.label.lastIndexOf('.')
            return item.label.substr(0, itemIndex) === data.label.substr(0, dataIndex);
          })
        ) {
          this.$refs.dlgTreeRef.setChecked(data.id, false, false);
          this.$message.warning(this.$t('datasetObj.dataset_not_equal_file'));
        } else if (this.dlgSelectedDatasetList.length === this.maxCount) {
          this.$refs.dlgTreeRef.setChecked(data.id, false, false);
          this.$message.error(this.$t('datasetObj.dataset_most', { msg: this.maxCount }));
        } else {
          this.dlgSelectedDatasetList.push(data);
        }
      } else {
        const index = this.dlgSelectedDatasetList.findIndex((item) => {
          return item.id === data.id;
        });
        this.dlgSelectedDatasetList.splice(index, 1);
      }
      this.dlgSelectedDataset = this.dlgSelectedDatasetList.map((item) => {
        return item.id;
      });
    },
    confirm() {
      const ids = this.dlgSelectedDatasetList.map(item => item.id);
      const names = this.dlgSelectedDatasetList.map(item => item.label);
      this.selectList.splice(0, Infinity);
      for (let i = 0, iLen = ids.length; i < iLen; i++) {
        this.selectList.push({
          id: ids[i],
          name: names[i]
        });
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
  beforeMount() { }
};
</script>

<style scoped lang="less">
.dataset-select {
  display: flex;
  margin-bottom: 28px;

  .title {
    width: 200px;
    text-align: right;
    margin-right: 24px;
    color: #101010;
    font-size: 14px;
    display: flex;
    justify-content: flex-end;
    padding-top: 6px;

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
    flex: 1;
    display: flex;

    .dataset-list-c {
      margin-right: 5px;
      flex: 1;
      min-height: 32px;
      border-radius: 4px;
      border: 1px solid #DCDFE6;
      box-sizing: border-box;
      color: #606266;
      padding: 0 15px;

      .dataset-item {
        line-height: 32px;
        font-size: 13px;
      }

      .dataset-item-placeholder {
        line-height: 32px;
        color: rgba(0, 0, 0, 0.6);
        opacity: 0.45 !important;
        font-size: 13px;
      }

      &.error {
        color: #9f3a38;
        background: #fff6f6;
        border-color: #e0b4b4;
      }
    }

    .btn-select {
      display: flex;
      justify-content: center;
      align-items: center;
      cursor: pointer;
      color: rgb(3, 102, 214);

      i {
        margin-right: 2px;
      }
    }
  }
}

.dataset-dlg {
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

      .dataset-tabs-c {
        display: flex;
        align-items: center;

        .dataset-tabs {
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

.custom-tree-node .dataset-title {
  font-size: 14px;
  color: #101010;
  font-weight: 600;
  flex: 1;
}

.custom-tree-node .dataset-repolink {
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

.dataset-nowrap {
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

.dataset-search-vue {
  z-index: 9999;
  position: absolute;
  right: 31%;
  height: 30px;
  top: 6px;
}

.select-dataset-label {
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-right: 1rem;
  white-space: nowrap;
}

.dataset_flex {
  display: flex;
  align-items: center;
}
</style>
