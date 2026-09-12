<template>
  <div class="list-container" v-loading="loading">
    <div class="list-item-container">
      <div class="item-container" v-for="(item, index) in list" :key="item.ID">
        <Item v-if="!isWorkspacePage" :data="item" :condition="params" :key="item.ID"></Item>
        <ItemOwner v-else :data="item" :condition="params" :key="item.ID" @changePinned="changePinned" @deleteEvent="deleteEvent"></ItemOwner>
      </div>
      <div v-show="(!list.length && !loading)" class="no-data">
        <div class="item-empty">
          <div class="item-empty-icon"></div>
          <div class="item-empty-tips">{{ $t('modelObj.model_square_empty') }}</div>
        </div>
      </div>
    </div>
    <div class="center" v-show="list.length">
      <el-pagination ref="paginationRef" background @current-change="currentChange" @size-change="sizeChange"
        :current-page.sync="iPage" :page-sizes="iPageSizes" :page-size.sync="iPageSize"
        layout="total, sizes, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </div>
    <delete-confirm-modal
      ref="deleteModal"
      type="repos"
      :data-obj="dataObj"
      @confirm-delete="handleDeleteModel"
      @delete-success="handleDeleteSuccess"
    />
  </div>
</template>

<script>
import Item from './Item.vue';
import ItemOwner from './ItemOwner.vue';
import { getReposListData, getNewRepoList, deleteRepos } from '~/apis/modules/repos';
import DeleteConfirmModal from '~/components/square/DeleteConfirmModal.vue'
import LetterAvatar from '~/utils/letteravatar';

export default {
  name: "List",
  props: {
    params: { type: Object, default: () => ({}) },
    isWorkspacePage: { type: Boolean, default: false },
  },
  components: { Item, ItemOwner, DeleteConfirmModal },
  data() {
    return {
      loading: false,
      list: [],
      iPageSizes: [30],
      iPageSize: 30,
      iPage: 1,
      total: 0,
      uid: '',
      dataObj: {},
    };
  },
  watch: {
    params: {
      handler(val, oval) {
        this.search()
      },
      deep: true,
    }
  },
  methods: {
    changePinned(){
      this.search()
    },
    async getListData() {
      this.loading = true;
      try {
        let params
        if(this.isWorkspacePage){
          params = {
            q: this.params.q || '',
            sort: this.params.sort || 'updated',
            order: 'desc',
            uid: this.uid,
            page: this.iPage,
            limit: this.iPageSize,
            mode: this.params.mode,
            exclusive: this.params.exclusive,
            // exclusive: 1,
            archived: false
          }         
        }else{
          params = {
            q: this.params.q || '',
            topic: this.params.topic || '',
            sort: this.params.sort || 'mostpopular',
            page: this.iPage,
            pageSize: this.iPageSize,
          }  
        }
        let getApi = this.isWorkspacePage ? getNewRepoList : getReposListData
        const res = await getApi(params);
        const resData = res.data;
        this.loading = false;
        if (resData.Data && resData.Data.Repos?.length) {
          const list = resData.Data.Repos || [];
          this.list = list.map((item) => ({
            ...item,
            NameShow: item.Alias,
            DescriptionShow: item.Description,
            TopicsShow: (item.Topics || []).map((_item) => ({
              topic: _item,
              topicShow: _item,
            })),
          }));
          this.total = resData.Data.Total;
          
          this.$nextTick(() => {
            LetterAvatar.transform();
          });
        } else {
          this.list = [];
          this.total = 0;
          this.iPage = this.iPage;
          this.iPageSize = this.iPageSize;
        }
      } catch (err) {
        console.log(err);
        this.loading = false;
        this.list = [];
        this.total = 0;
        this.iPage = this.iPage;
        this.iPageSize = this.iPageSize;
      }
    },
    search() {
      this.iPage = 1;
      this.getListData();
    },
    currentChange(page) {
      this.iPage = page;
      this.getListData();
    },
    sizeChange(pageSize) {
      this.iPageSize = pageSize;
      this.getListData();
    },
    deleteEvent(data){
      console.log("xxxxxdele",data)
      this.dataObj = data
      this.$refs.deleteModal.showModal()
    },
    async handleDeleteModel({ type, data, callback }) {
      try {
        const response = await deleteRepos(data)
        if (response.data.Code === 0) {
          callback({
            success: true,
            message: this.$t('imagesObj.deleteSuccessTips')
          })
        } else {
          callback({
            success: false,
            error: response.data.message
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
      this.search() // 刷新列表
    },
  },
  beforeMount() {
    const metaEl = document.querySelector('meta[name="_context_uid"]');
    if (metaEl) {
      this.uid = metaEl.getAttribute('content');
    }
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.list-container {
  min-height: 500px;
  .list-item-container {
    display: flex;
    flex-wrap: wrap;
    .item-container {
      width: 33.3%;
      padding: 12px;
    }
  }
}

.center {
  text-align: center;
  margin-top: 10px;
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
</style>
