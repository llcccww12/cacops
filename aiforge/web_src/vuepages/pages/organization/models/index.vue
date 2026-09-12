<template>
  <div class="content">
    <div class="content-r">
      <div class="filter-c-m">
        <span class="title">{{ $t('org.orgModel')}}</span>
        <div class="right-m">
          <div class="icon-box" v-if="canCreate" @click="createDataset"><i class="ri-add-box-line"></i></div>
          <div class="icon-box" v-if="!showSearch" @click="showSearch = true"><i class="ri-search-line"></i></div>
          <div v-else  class="ui small icon input" style="height: 32px;margin-left: 6px;">
            <input type="text" :placeholder="$t('modelObj.model_search')" v-model="q" @keyup.enter="search">
            <i class="search icon" style="cursor: pointer;pointer-events: auto;" @click="search"></i>
          </div> 
        </div>
      </div>
      <div class="workspace-t">
        <span>{{ $t('org.orgModel')}}</span>
      </div>
      <div class="filter-c">
        <div class="tab-c">
            <div class="tab-item nowrap" :class="tab == item.key ? 'active' : ''"
            v-for="(item) in tabList" :key="item.key" @click="changeTab(item)">{{
                item.label }}</div>
        </div>
        <div class="right">
          <div class="ui small icon input" style="height: 32px;">
            <input type="text" :placeholder="$t('org.searchModels')" v-model="q" @keyup.enter="search">
            <i class="search icon" style="cursor: pointer;pointer-events: auto;" @click="search"></i>
          </div> 
          <el-select v-model="sort" placeholder="请选择" @change="changeSort" style="width: 120px;">
            <el-option
              v-for="item in sortList"
              :key="item.key"
              :label="item.label"
              :value="item.key">
            </el-option>
          </el-select>
          <div class="btn-add" v-if="canCreate" @click="createDataset"> 
            <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 32 32" width="20" height="20"><defs></defs><g><path d="M5.333 4h21.333c0.736 0 1.333 0.597 1.333 1.333v0 21.333c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-21.333c0-0.736 0.597-1.333 1.333-1.333v0zM6.667 6.667v18.667h18.667v-18.667h-18.667zM14.667 14.667v-5.333h2.667v5.333h5.333v2.667h-5.333v5.333h-2.667v-5.333h-5.333v-2.667h5.333z"></path></g></svg>
            <span style="margin-left:6px;">{{ $t('modelManage.createNewModel')  }}</span>
          </div>
        </div>
      </div>
      <div class="list-item-container" v-loading="loading">
        <div class="item-container" v-for="(item, index) in list" :key="item.id">
          <DatasetItem type="aimodel" :data="item" :canChangeFav="true" :operaFlag="true" @deleteEvent="deleteEvent"></DatasetItem>
        </div>
        <div v-show="(!list.length && !loading)" class="no-data">
          <div class="item-empty">
            <div class="item-empty-icon"></div>
            <div class="item-empty-tips">{{ $t('org.no_result') }}</div>
          </div>
        </div>
      </div>
      <div class="center" v-show="list.length">
        <el-pagination ref="paginationRef" background @current-change="currentChange" @size-change="sizeChange"
          :current-page.sync="iPage" :pager-count="5" :page-sizes="iPageSizes" :page-size.sync="iPageSize"
          layout="total, sizes, prev, pager, next, jumper" :total="total">
        </el-pagination>
      </div>
    </div>
    <!-- 移动端left-lables遮罩层（独立层级，不覆盖按钮） -->
    <!-- 确认删除弹框 -->
    <delete-confirm-modal
      ref="deleteModal"
      type="modelObj"
      :data-obj="dataObj"
      @confirm-delete="handleDeleteModel"
      @delete-success="handleDeleteSuccess"
    />
  </div>
</template>
<script>
import DatasetItem from '~/components/square/RightOwnerItem.vue'
import DeleteConfirmModal from '~/components/square/DeleteConfirmModal.vue'
import { getOrgModelList } from "~/apis/modules/organization";
import { delDataset } from "~/apis/modules/dataset";
import { MODEL_ENGINES } from '~/const';
import { getListValueWithKey } from '~/utils';
import { formatDate } from 'element-ui/lib/utils/date-util';

export default {
  data() {
    return {
      tab: 'all',
      tabList: [
       { key: 'all', label: this.$t('cloudbrainObj.all'), },
      ],
      q: '',
      sort: 'recentupdate',
      sortList: [{
        key: 'recentupdate',
        label: this.$t('datasets.recentupdate'),
      }, {
        key: 'newest',
        label: this.$t('datasets.newest'),
      }, {
        key: 'downloadcount',
        label: this.$t('datasets.downloadtimes'),
      }, {
        key: 'collections',
        label: this.$t('datasets.moststars'),
      }, {
        key: 'usecount',
        label: this.$t('datasets.mostusecount'),
      }, {
        key: 'alias_asc',
        label: this.$t('datasets.alphabetasc'),
      }, {
        key: 'alias',
        label: this.$t('datasets.alphabetdesc'),
      }],
      
      loading: false,
      iPageSizes: [30],
      iPageSize: 30,
      iPage: 1,
      total: 0,
      list: [],
      dataObj: {},
      delLoading: false,
      orgName: '',
      showSearch: false,
      uid: '',
      canCreate: false,
    };
  },
  components: { DatasetItem, DeleteConfirmModal },
  methods: {
    changeTab(){

    },
    createDataset(){
      location.href = `/guide/create_model?org=${this.uid}`
    },
    search(){
      this.iPage = 1;
      this.getCardList()
    },
    changeSort(item) {
      this.sort = item
      this.search();
    },
    currentChange(page) {
      this.iPage = page;
      this.getCardList()
    },
    sizeChange(pageSize) {
      this.iPageSize = pageSize;
      this.search();
    },
    async getCardList() {
      try {
        this.loading = true;
        const requestParams = {
          owner_name: this.orgName,
          q: this.q.trim(),
          engine: '',
          label: '',
          recommend: '',
          order_by: this.sort,
          page: this.iPage,
          page_size: this.iPageSize,
        };
        const response = await getOrgModelList(requestParams)
        this.loading = false;
        
        if(response.data.code === 0){
          const res = response.data.data
          const list = (res.aimodels || []).map(item=>{
            return {
              ...item,
              labels: item.tags ? item.tags.trim().split(/\s+/) : [],
              tags: '',
              licenses: '',
              engineName: getListValueWithKey(MODEL_ENGINES, item.engine),
              updateTimeStr: formatDate(new Date(item.updated_unix * 1000), 'yyyy-MM-dd')
            }
          })
          this.total = +res.total || 0;
          this.list = list
        }else{
          this.$message.error(res.msg)
        }
      } catch (error) {
        this.loading = false;
        console.log(error)
        this.$message.error(error)
      }
    },
    deleteEvent(data){
      this.dataObj = data
      this.$refs.deleteModal.showModal()
    },
    async handleDeleteModel({ type, data, callback }) {
      try {
        const response = await delDataset({ aimodel_id: data.id },'aimodel')
        
        if (response.data.code === 0) {
          callback({
            success: true,
            message: this.$t('imagesObj.deleteSuccessTips')
          })
        } else {
          callback({
            success: false,
            error: response.data.msg
          })
        }
      } catch (error) {
        callback({
          success: false,
          error: error.message
        })
      }
    },
    // 删除成功后的回调
    handleDeleteSuccess() {
      this.search();
    },
  },
  beforeMount() {
    const { pathname } = window.location;
    const orgMatch = pathname.match(/^\/org\/([^\/]+)(\/.*)?$/);
    this.orgName = orgMatch[1]
  },
  mounted() {
    const metaEl = document.querySelector('#org-info');
    console.log(metaEl)
    if (metaEl) {
      this.canCreate = metaEl.getAttribute('data-cancreate');
      this.uid = metaEl.getAttribute('data-orgid');
      console.log("this.uid",this.uid)
    }
    this.getCardList();
  },
  beforeDestroy() { },
};
</script>
<style scoped lang="less">
.content {
  display: flex;
  min-height: 100%;
  padding-left: 20px;
}
.content-r {
  flex: 1;
  width: 0;
  padding: 30px 20px;
  padding-right: 36px;
  .filter-c-m{
    display: none;
    margin-bottom: 14px;
    align-items: center;
    justify-content: space-between;
    .title{
      color: rgb(16,16,16);
      font-size: 18px;
      height: 32px;
      line-height: 35px;
      font-family: Arial-bold;
      font-weight: bold;
    }
    .right-m{
      display: flex;
      .icon-box{
        border: 1px solid rgba(16,16,16,0.2);
        border-radius: 5px;
        width: 32px;
        height: 32px;
        margin-left: 6px;
        display: flex;
        justify-content: center;
        align-items: center;
        i{
          color: #101010;
          font-size: 16px;
        }
      }
      .filter-btn{
        &.active{
          background-color: rgba(255,255,255,1);
          z-index: 999;
          i{
            color: #0066ff
          }
        }
      }
    }
  }
  .workspace-t{
    color: rgb(16,16,16);
    font-size: 18px;
    font-weight: 700;
    padding: 6px 0 16px 0;
  }  
  .filter-c {
    display: flex;
    align-items: center;
    justify-content: space-between;

    .tab-c {
      display: flex;
      align-items: center;
      
      .tab-item {
        height: 32px;
        display: flex;
        align-items: center;
        margin-right: 6px;
        padding: 0 6px;
        border-bottom: 2px solid rgba(51, 38, 98, 0.3);
        font-size: 16px;
        box-sizing: border-box;
        cursor: pointer;
        color: rgba(16,16,16,0.5);
        &:hover {
          color: rgba(16,16,16,1);
          border-color: rgba(51, 38, 98, 1);
        }
        &.active {
          color: rgba(0,102,255,1);
          border-bottom: 2px solid rgba(0,102,255,1);
        }
      }
    }

    .right {
      display: flex;
      align-items: center;
      gap: 14px;
      .check-c {
        margin-right: 16px;
      }
      .btn-add{
        height: 32px;
        background-color: rgba(50,145,248,1);
        font-size: 14px;
        display: flex;
        align-items: center;
        padding: 0 10px;
        color: #fff;
        border-radius: 4px;
        cursor: pointer;
        .fill:not([store]){
          fill: rgb(255, 255, 255);
        }
      }
      .search-c {
        margin-right: 16px;

        .el-input {
          .el-button {
            color: rgb(16, 16, 16);
          }
        }
      }

      .sort-c {
        cursor: pointer;
      }
    }
    .right-m{
      display: none;
      align-items: center;
      .icon-box{
        border: 1px solid rgba(16,16,16,0.2);
        border-radius: 5px;
        width: 32px;
        height: 32px;
        margin-left: 6px;
        display: flex;
        justify-content: center;
        align-items: center;
        i{
          color: #101010;
          font-size: 16px;
        }
      }
    }
  }
  .list-item-container {
    display: flex;
    flex-wrap: wrap;

    .item-container {
      width: 33.3%;
      padding: 12px;
    }
    .no-data {
      display: flex;
      justify-content: center;
      padding: 12px 0;
      width: 100%;

      .item-empty {
        height: 391px;
        width: 100%;
        overflow: hidden;
        padding: 15px;
        background: transparent;
        display: flex;
        flex-direction: column;
        justify-content: center;
        background-color: rgba(245, 245, 246, 0.5);

        .item-empty-icon {
          height: 80px;
          width: 100%;
          background: url(/img/empty-box.svg) center center no-repeat;
        }

        .item-empty-tips {
          text-align: center;
          margin-top: 20px;
          font-size: 18px;
          color: rgb(63, 63, 64);
        }
      }
    }
  }
  .center {
    text-align: center;
    margin-top: 10px;
  }
}

/deep/.el-dropdown-menu__item.active {
  color: #409EFF;
  background-color: rgba(179, 216, 255, 0.3);
}

@media only screen and (max-width: 849.99px) {
  .content {
    padding-left: 0 !important;
  }
  .content-r {
    padding: 20px 16px;
    .filter-c-m{
      display: flex;
    }
    .workspace-t{
      display: none;
    }
    .filter-c{
      justify-content: center;
      .right{
        display: none;
      }
      .tab-c{
        display: none;
      }
    }
    .list-item-container {
      .item-container {
        width: 100% !important;
        padding: 12px 0 !important;
      }
    }
    .center {
      .el-pagination {
        /deep/ .el-pagination__total,
        /deep/ .el-pagination__sizes,
        /deep/ .el-pagination__jump {
          display: none;
        }
      }
    }
  }
}
@media only screen and (min-width: 850px) and (max-width: 1440px) {
  .content-r{
    .filter-c{
      .right{
        display: none;
      }
      .right-m{
        display: flex;
      }
    }
    .list-item-container {
      .item-container {
        width: 50% !important;

        &:nth-child(2n+1) {
          padding-left: 0;
        }

        &:nth-child(2n) {
          padding-right: 0;
        }
      }
    }
  }
}
@media only screen and (min-width: 1440px) {
  /deep/ .list-item-container {
    .item-container {
      width: 33.3% !important;

      &:nth-child(3n+1) {
        padding-left: 0;
      }

      &:nth-child(3n) {
        padding-right: 0;
      }
    }
  }
}
</style>