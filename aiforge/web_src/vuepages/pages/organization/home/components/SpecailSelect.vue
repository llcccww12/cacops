<template>
  <div class="special-select-c" v-show="pageData.IsOrganizationOwner || dataList.length">
    <div class="head">
      <div class="head-l">
        ⭐️ <span>{{ lang.headTitle }}</span>
      </div>
      <div class="head-r">
        <div class="btn-c" v-if="pageData.IsOrganizationOwner" @click="dlgShow = true">
          <i class="ri-settings-2-line"></i>
          <span>{{ $t('org.customize') }}</span>
        </div>
      </div>
    </div>
    <div v-if="type == 'repo'" class="body list-item-container repo-c">
      <div class="item-container" v-for="(item, index) in dataList" :key="item.ID">
        <RepoItem :data="item"></RepoItem>
      </div>
    </div>
    <div v-if="type == 'model'" class="body list-item-container model-c">
      <div class="item-container" v-for="(item, index) in dataList" :key="item.ID">
        <DatasetItem :data="item" :canChangeFav="false" type="aimodel"></DatasetItem>
      </div>
    </div>
    <div v-if="type == 'dataset'" class="body list-item-container dataset-c">
      <div class="item-container" v-for="(item, index) in dataList" :key="item.id">
        <DatasetItem :data="item" :canChangeFav="false"></DatasetItem>
      </div>
    </div>
    <BaseDialog class="model-dlg" :visible.sync="dlgShow" :title="lang.dlgTitle" width="1000px" :modal="true"
      :close-on-click-modal="false" :show-close="true" :destroy-on-close="false" :before-close="beforeClose"
      @open="open" @closed="closed">
      <div class="dlg-content">
        <div class="tips">{{ lang.dlgMaxCountTip }}</div>
        <div class="search-bar"><input type="text" v-model="searchKey" @input="searchChange"
            :placeholder="$t('org.searching')"></div>
        <div class="list-c">
          <el-checkbox-group class="list" v-model="checkList" @change="handleChange">
            <el-checkbox class="item" v-for="(item) in list" :key="item.id" :label="item.id" v-show="item.show"
              :title="item.name">
              <span v-html="renderMatchText(item.name)"></span>
            </el-checkbox>
          </el-checkbox-group>
        </div>
        <div class="tips-remaining">{{ $t('org.remainNum', { num: remaining }) }}</div>
        <div class="btn-c">
          <el-button type="primary" class="btn confirm-btn" @click="save">{{ $t('confirm') }}</el-button>
          <el-button class="btn" @click="cancel">{{ $t('cancel') }}</el-button>
        </div>
      </div>
    </BaseDialog>
  </div>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';

import RepoItem from './RepoItem.vue';
import LetterAvatar from '~/utils/letteravatar';
import { getOrgSelectedRepoList, getOrgSelectedRepoSetList, setOrgSelectedRepo } from '~/apis/modules/organization';

import { MODEL_ENGINES } from '~/const';
import { getListValueWithKey } from '~/utils';
import { formatDate } from 'element-ui/lib/utils/date-util';
import { getOrgSelectedModelList, getOrgSelectedModelSetList, setOrgSelectedModel } from '~/apis/modules/organization';

import DatasetItem from '~/components/square/RightItem.vue';
import { getOrgSelectedDatasetList, getOrgSelectedDatasetSetList, setOrgSelectedDataset } from '~/apis/modules/organization';

export default {
  name: "SpecialSelect",
  props: {
    type: { type: String, default: "repo" }, // 'repo|model|dataset'
    pageData: { type: Object, default: () => { } },
  },
  data() {
    return {
      loading: false,
      dataList: [],
      dlgShow: false,
      maxCount: 9,
      searchKey: '',
      oCheckList: [],
      checkList: [],
      oList: [],
      list: [],
      lang: {
        headTitle: '',
        dlgTitle: '',
        dlgMaxCountTip: '',
      },
    };
  },
  components: { BaseDialog, RepoItem, DatasetItem },
  computed: {
    remaining: function () {
      return this.maxCount - this.checkList.length;
    },
  },
  methods: {
    ishaveDataset(id){
      let findEl = this.dataList.some(item => item.id === id)
      if(findEl){
        this.getData()
      }
    },
    // get
    getData() {
      this.dataList = [];
      if (this.type == 'repo') {
        this.loading = true;
        getOrgSelectedRepoList({ orgName: this.pageData.Org.Name }).then(res => {
          res = res.data;
          this.loading = false;
          if (res.code == 0) {
            const list = res.data[0]?.RepoList || [];
            this.dataList = list.map((item) => {
              item.Contributors = (item.Contributors || []).map((_item) => {
                return {
                  ..._item,
                  bgColor: this.randomColor((_item.Email[0] || '').toLocaleUpperCase()),
                }
              });
              return {
                ...item,
                NameShow: item.Alias,
                DescriptionShow: item.Description,
                TopicsShow: (item.Topics || []).map((_item) => {
                  return {
                    topic: _item,
                    topicShow: _item
                  }
                }),
              }
            });
            this.$nextTick(() => {
              LetterAvatar.transform();
            });
          }
        }).catch(err => {
          console.log(err);
          this.loading = false;
        });
      }
      if (this.type == 'model') {
        this.loading = true;
        getOrgSelectedModelList({ orgName: this.pageData.Org.Name }).then(res => {
          res = res.data;
          this.loading = false;
          this.dataList = (res.data.aimodels || []).map(item => {
            return {
              ...item,
              labels: item.tags ? item.tags.trim().split(/\s+/) : [],
              tags: '',
              licenses: '',
              engineName: getListValueWithKey(MODEL_ENGINES, item.engine),
              updateTimeStr: formatDate(new Date(item.updated_unix * 1000), 'yyyy-MM-dd'),
            }
          });
        }).catch(err => {
          console.log(err);
          this.loading = false;
        });
      }
      if (this.type == 'dataset') {
        this.loading = true;
        getOrgSelectedDatasetList({ orgName: this.pageData.Org.Name }).then(res => {
          res = res.data;
          this.loading = false;
          if(res.code === 0){
            this.dataList = (res.data.datasets || []).map(item => {
              return {
                ...item,
                updateTimeStr: formatDate(new Date(item.updated_unix * 1000), 'yyyy-MM-dd'),
              }
            });
          }else{
            this.$message.error(res.msg);
          }
          
        }).catch(err => {
          console.log(err);
          this.loading = false;
          this.$message.error(err);
        });
      }
    },
    randomColor(t) {
      const tIndex = t.charCodeAt(0);
      const colorList = ["#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#34495e", "#16a085", "#27ae60", "#2980b9", "#8e44ad",
        "#2c3e50", "#f1c40f", "#e67e22", "#e74c3c", "#00bcd4", "#95a5a6", "#f39c12", "#d35400", "#c0392b", "#bdc3c7", "#7f8c8d"];
      return colorList[tIndex % colorList.length];
    },
    // set
    open() {
      this.olist = [];
      this.list = [];
      this.searchKey = '';
      if (this.type == 'repo') {
        getOrgSelectedRepoSetList({ orgName: this.pageData.Org.Name }).then(res => {
          res = res.data;
          const data = res.data || [];
          this.olist = data.map(item => {
            return {
              id: item.RepoID,
              name: item.RepoName,
              selected: item.Selected,
            }
          });
          this.stateInit();
        }).catch(err => {
          console.log(err);
        });
      }
      if (this.type == 'model') {
        getOrgSelectedModelSetList({ orgName: this.pageData.Org.Name }).then(res => {
          res = res.data;
          if (res.code == 0) {
            const data = res.data.aimodels || [];
            this.olist = data.map(item => {
              return {
                id: item.ModelID,
                name: item.ModelName,
                selected: item.Selected,
              }
            });
            this.stateInit();
          }
        }).catch(err => {
          console.log(err);
        });
      }
      if (this.type == 'dataset') {
        getOrgSelectedDatasetSetList({ orgName: this.pageData.Org.Name }).then(res => {
          res = res.data;
          console.log(res);
          if (res.code == 0) {
            const data = res.data.datasets || [];
            this.olist = data.map(item => {
              return {
                id: item.DatasetId,
                name: item.DatasetName,
                selected: item.Selected,
              }
            });
            this.stateInit();
          }
        }).catch(err => {
          console.log(err);
        });
      }
    },
    stateInit() {
      this.list = this.olist;
      this.searchChange();
      this.checkList = this.list.filter(item => item.selected).map(item => item.id);
      this.oCheckList = [...this.checkList];
    },
    searchChange() {
      this.list = this.list.map(item => {
        return {
          ...item,
          show: item.name.toLowerCase().includes(this.searchKey.toLowerCase()),
        }
      })
    },
    renderMatchText(string) {
      const searchKey = this.searchKey.trim();
      if (!searchKey) return string;
      const reg = new RegExp(searchKey, 'gi');
      return string.replace(reg, (txt) => {
        return `<span style="color:red">${txt}</span>`;
      });
    },
    beforeClose(done) {
      done();
    },
    closed() { },
    handleChange(checkedList) {
      if (this.remaining < 0) {
        checkedList.pop();
        this.checkList = checkedList;
      }
    },
    save() {
      if (this.setLoading) return;
      if (this.type == 'repo') {
        this.setLoading = true;
        setOrgSelectedRepo({ orgName: this.pageData.Org.Name, list: this.checkList }).then(res => {
          res = res.data;
          this.setLoading = false;
          if (res.code == 0) {
            this.dlgShow = false;
            this.getData();
          } else {
            this.$message.error(res.msg);
          }
        }).catch(err => {
          console.log(err);
          this.setLoading = false;
          this.$message.error(this.$t('submittedFailed'));
        });
      }
      if (this.type == 'model') {
        this.setLoading = true;
        setOrgSelectedModel({ orgName: this.pageData.Org.Name, list: this.checkList }).then(res => {
          res = res.data;
          this.setLoading = false;
          if (res.code == 0) {
            this.dlgShow = false;
            this.getData();
          } else {
            this.$message.error(res.msg);
          }
        }).catch(err => {
          console.log(err);
          this.setLoading = false;
          this.$message.error(this.$t('submittedFailed'));
        });
      }
      if (this.type == 'dataset') {
        this.setLoading = true;
        setOrgSelectedDataset({ orgName: this.pageData.Org.Name, list: this.checkList }).then(res => {
          res = res.data;
          this.setLoading = false;
          if (res.code == 0) {
            this.dlgShow = false;
            this.getData();
          } else {
            this.$message.error(res.msg);
          }
        }).catch(err => {
          console.log(err);
          this.setLoading = false;
          this.$message.error(this.$t('submittedFailed'));
        });
      }
    },
    cancel() {
      this.dlgShow = false;
    },
  },
  beforeMount() {
    if (this.type == 'repo') {
      this.lang = {
        headTitle: this.$t('org.selectedProjects'),
        dlgTitle: this.$t('org.customizeSelectedProjects'),
        dlgMaxCountTip: this.$t('org.maxProjects', { num: this.maxCount }),
      }
    }
    if (this.type == 'model') {
      this.lang = {
        headTitle: this.$t('org.selectedModels'),
        dlgTitle: this.$t('org.customizeSelectedModels'),
        dlgMaxCountTip: this.$t('org.maxModels', { num: this.maxCount }),
      }
    }
    if (this.type == 'dataset') {
      this.lang = {
        headTitle: this.$t('org.selectedDatasets'),
        dlgTitle: this.$t('org.customizeSelectedDatasets'),
        dlgMaxCountTip: this.$t('org.maxDatasets', { num: this.maxCount }),
      }
    }
  },
  mounted() {
    this.getData();
  },
};
</script>

<style scoped lang="less">
.special-select-c {
  border-color: rgb(225, 227, 230);
  border-width: 1px;
  border-style: solid;
  border-radius: 15px;
  padding: 18px 28px;
  margin-bottom: 32px;

  .head {
    display: flex;
    align-items: center;
    justify-content: space-between;

    .head-l {
      display: flex;
      align-items: center;
      height: 40px;
      font-size: 18px;

      span {
        margin-left: 5px;
        color: rgb(16, 16, 16);
        font-family: SourceHanSansSC;
        font-weight: 550;
      }
    }

    .head-r {
      .btn-c {
        display: flex;
        align-items: center;
        color: rgba(3, 102, 214, 1);
        font-size: 14px;
        line-height: 20px;

        i {
          margin-right: 5px;
        }

        cursor: pointer;

        &:hover {
          color: rgba(3, 102, 214, 0.8);
        }
      }
    }
  }

  .body.list-item-container {
    &.repo-c {
      display: flex;
      flex-wrap: wrap;

      .item-container {
        width: 33.3%;
        padding: 8px;

        &:nth-child(3n+1) {
          padding-left: 0;
        }

        &:nth-child(3n) {
          padding-right: 0;
        }
      }
    }

    &.model-c,
    &.dataset-c {
      display: flex;
      flex-wrap: wrap;

      .item-container {
        width: 50%;
        padding: 12px;

        &:nth-child(odd) {
          padding-left: 0;
        }

        &:nth-child(even) {
          padding-right: 0;
        }
      }
    }
  }

  .dlg-content {
    padding: 30px 60px 20px 60px;

    .tips {
      color: rgba(68, 68, 68, 1);
      font-size: 14px;
    }

    .search-bar {
      height: 40px;
      position: relative;
      margin-top: 12px;

      input {
        width: 100%;
        height: 100%;
        padding: 0px 8px;
        letter-spacing: 0px;
        color: rgb(136, 136, 136);
        position: relative;
        font-family: PingFangSC;
        font-weight: 400;
        outline: none;
        box-shadow: none;
        border: none;
        border-radius: 5px;
        border-color: rgb(187, 187, 187);
        border-width: 1px;
        border-style: solid;
        border-radius: 5px;

        &:focus {
          border-color: #66b1ff;
        }
      }
    }

    .list-c {
      min-height: 200px;

      .list {
        display: flex;
        flex-wrap: wrap;
        max-height: 240px;
        overflow-y: auto;
        margin-top: 12px;

        .el-checkbox {
          width: 33%;
          margin-bottom: 6px;
          overflow: hidden;
          position: relative;
          display: flex;
          align-items: center;
          margin-right: 0px;
          height: 22px;

          /deep/.el-checkbox__input {
            margin-top: 2px;
          }

          /deep/.el-checkbox__label {
            width: calc(100% - 24px);
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
          }
        }
      }
    }

    .tips-remaining {
      color: rgb(16, 16, 16);
      font-size: 14px;
      font-family: SourceHanSansSC;
      font-weight: 400;
      line-height: 20px;
      margin-top: 16px;
    }

    .btn-c {
      display: flex;
      align-items: center;
      justify-content: center;
      margin-top: 20px;

      .btn {
        color: rgb(2, 0, 4);
        background-color: rgb(194, 199, 204);
        border-color: rgb(194, 199, 204);

        &.confirm-btn {
          color: #fff;
          background-color: rgb(56, 158, 13);
          border-color: rgb(56, 158, 13);
          margin-right: 10px;
        }
      }
    }
  }
}

@media only screen and (max-width: 767px) {
  .special-select-c {
    border: none;
    background: rgb(245, 245, 246);
    border-radius: 0;
    padding: 12px 12px;
    margin-bottom: 20px;
    width: calc(100% + 2em);
    margin-left: -1em;

    .head {
      justify-content: flex-start;

      .head-r {
        margin-left: 10px;
      }
    }

    .repo-c {
      .item-container {
        width: 50% !important;
        padding: 8px !important;

        &:nth-child(2n+1) {
          padding-left: 2px !important;
        }

        &:nth-child(2n) {
          padding-right: 2px !important;
        }

        .repo-item {
          background-color: white;
        }
      }
    }

    .model-c,
    .dataset-c {
      .item-container {
        width: 100% !important;
        padding: 8px !important;
        padding-left: 8px !important;
        padding-right: 8px !important;

        .item {
          background-color: white;
        }
      }
    }
  }

  .model-dlg {

    /deep/ .el-dialog {
      width: 95% !important;

      .dlg-content {
        padding: 20px 20px 20px 20px !important;

        .el-checkbox-group {
          .el-checkbox {
            width: 100%;
          }
        }
      }
    }
  }
}
</style>
