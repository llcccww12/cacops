<template>
  <div class="content">
    <left-labels
      :title="'dataset'"
      :filters="filtersConfig"
      @trigger-search="changeFilter"
      :isWorkspacePage="isWorkspacePage"
    />
    <right-cards
      ref="rightCards"
      :title="'dataset'"
      :loading="loading"
      :tabList="(isWorkspacePage ? tabList2 : tabList1)"
      :isWorkspacePage="isWorkspacePage"
      :tab="tab"
      :type="'dataset'"
      :list="list"
      :pageParams="pageParams"
      @changeParams="changeParams"
      @changeModalOverlay="changeModalOverlay"
      @deleteEvent="deleteEvent"
    >
    </right-cards>
    <!-- 移动端left-lables遮罩层（独立层级，不覆盖按钮） -->
    <div class="modal-overlay" v-if="showFilter">
      <div class="filter-panel">
        <!-- 弹窗内容（与图片描述一致） -->
        <div class="filter-section">
          <div class="filter-content">
            <div  class="block-c" v-for="(filter, index) in filtersConfig" :key="index">
              <div class="title-c">
                <span class="title">{{ $t(filter.titleKey) }}</span>
              </div>
              <div class="list">
                <div class="item" :class="{ active: item.active, [filter.styleClass]: true }"
                  v-for="(item, idx) in filter.items" :key="idx" @click="selectItem(filter.type, item)">
                  {{ item.v }}
                </div>
              </div>
            </div>
            <!-- 其他分类标签 -->
          </div>
          <div class="filter-buttons">
            <div class="reset-button btn" @click="resetFilters">重置</div>
            <div class="search-button btn" @click="searchFilters">搜索</div>
          </div>
        </div>
        
      </div>
    </div>
    <!-- 确认删除弹框 -->
    <delete-confirm-modal
      ref="deleteModal"
      type="datasetObj"
      :data-obj="dataObj"
      @confirm-delete="handleDeleteModel"
      @delete-success="handleDeleteSuccess"
    />
  </div>
</template>

<script>
import LeftLabels from '~/components/square/LeftLabels.vue'
import RightCards from '~/components/square/RightCards.vue'
import DeleteConfirmModal from '~/components/square/DeleteConfirmModal.vue'
import { getDatasets, getPromoteDataset, delDataset } from "~/apis/modules/dataset";
import { formatDate } from 'element-ui/lib/utils/date-util';
import { lang } from '~/langs';
const isWorkspacePage = window.MENU_CONFIG.activeTopMenu == 'workspace'
export default {
  data() {
    return {
      isWorkspacePage: isWorkspacePage,
      tab: isWorkspacePage ? '/owned' : '/public?recommend=all',
      tabList1: [
        { key: '/public?recommend=all', label: `⭕${this.$t('datasets.publick_dataset')}`, favFlag: true },
        { key: '/public?recommend=only', label: `🏆${this.$t('datasets.recommend_dataset')}`, favFlag: true },
        { key: '/modelscope', label: '🌐 魔塔社区', favFlag: false, modelscope: true },
      ],
      tabList2: [
        { key: '/owned', label: this.$t('datasets.my_dataset'), favFlag: false, operaFlag: true },
        { key: '/collaborated', label: this.$t('datasets.collaborated_dataset'), favFlag: false },
        { key: '/collected', label: this.$t('datasets.favorite_dataset'), favFlag: true, reloadFlag: true },
      ],
      filtersConfig: [],
      params: {
        url: '',
        q: '',
        tasks: '',
        tags: '',
        license: '',
        recommend: false,
        order_by: '',
        owner_type: isWorkspacePage ? 'individual' : '',
        page: 1,
        page_size: 30,
      },
      loading: false,
      pageParams: {
        page: 1,
        page_size: 30,
        total: 0
      },
      list: [],
      showFilter: false,
      dataObj: {},
      name: '',
      delLoading: false,
      codeUsePromotePath: `home_v2/dataset_square${lang == 'zh-CN' ? '' : '_en'}.json`,
    };
  },
  components: { LeftLabels, RightCards, DeleteConfirmModal },
  methods: {
    changeModalOverlay(val){
      this.showFilter = val;
    },
    selectItem(type, clickedItem){
       // 1. 找到对应的筛选组
      const filterGroup = this.filtersConfig.find(f => f.type === type);
      if (!filterGroup) return;

      this.params[type] = clickedItem.k;

      filterGroup.items.forEach(item => {
        item.active = (item === clickedItem); // 只有当前点击项为 true
      });
    },
    resetFilters() {
      this.filtersConfig.forEach(group => {
        this.params[group.type] = ''
        group.items.forEach(item => {
          item.active = false;
        });
      });
    },
    searchFilters(){
      this.getCardList() 
      this.showFilter = false
      console.log()
      this.$refs.rightCards.resetFilter()
    },
    async getCardList() {
        try {
            this.loading = true;
            if (this.params.url === '/modelscope') {
              const response = await fetch(`/api/v1/modelscope/datasets?page=${this.params.page || 1}&page_size=${this.params.page_size || 30}&q=${encodeURIComponent(this.params.q || '')}&order_by=${encodeURIComponent(this.params.order_by || '')}`)
              const payload = await response.json()
              this.loading = false
              if (payload.code === 0) {
                const res = payload.data || {}
                this.list = (res.datasets || []).map(item => ({
                  ...item,
                  updateTimeStr: formatDate(new Date(item.updated_unix * 1000), 'yyyy-MM-dd'),
                }))
                this.pageParams = {
                  page: +res.page || 1,
                  page_size: +res.page_size || 30,
                  total: +res.total || 0,
                }
              } else {
                this.$message.error(payload.msg || '加载魔塔社区数据集失败')
              }
              return
            }
            // 处理 URL 中的重复参数
             const requestParams = { ...this.params };
    
            // 如果URL中有参数，提取出来
            if (requestParams.url.includes('?')) {
              const [path, query] = requestParams.url.split('?');
              const urlParams = new URLSearchParams(query);
              
              // 优先使用URL中的参数
              if (urlParams.has('recommend')) {
                const urlRecommend = urlParams.get('recommend');
                requestParams.recommend = urlRecommend;
                // 从URL中删除recommend参数，避免重复
                urlParams.delete('recommend');
              }
              
              // 重新构建URL
              const newQuery = urlParams.toString();
              requestParams.url = newQuery ? `${path}?${newQuery}` : path;
            }
            const response = await getDatasets(requestParams)
            this.loading = false;
            
            if(response.data.code === 0){
              const res = response.data.data
              const list = (res.datasets || []).map(item=>{
                return {
                  ...item,
                  updateTimeStr: formatDate(new Date(item.updated_unix * 1000), 'yyyy-MM-dd'),
                }
              })
              // this.total = +res.total || 0;
              this.pageParams = {
                page: +res.page || 0,
                page_size: +res.page_size || 0,
                total: +res.total || 0
              }
              this.list = list
            }else{
              this.$message.error(res.data.msg)
            }
            
        } catch (error) {
            this.loading = false;
            console.log(error)
            this.$message.error(error)
        }
    },
    changeFilter(obj){
        this.params = {...this.params, ...obj}
        this.getCardList()
    },
    changeParams(obj){
        this.params = {...this.params, ...obj}
        console.log(this.params)
        this.getCardList()
    },
    async getDatasetTags() {
        try {
          const response = await getPromoteDataset({key: this.codeUsePromotePath})
          console.log(response)
          if(response.data.code === 0){
            const res = JSON.parse(response.data.data || '{}') || {}
            const {category = {}, license = {}, task = {}} = res.tagInformation || {}
            const createTagArray = (source, keyPrefix) => {
              return Object.keys(source).map(key => ({
                k: key,
                v: source[key],
                active: false
              }));
            };

            const categoryTag = createTagArray(category);
            const licenseTag = createTagArray(license);
            const taskTag = createTagArray(task);
            this.filtersConfig = [
              {
                  type: 'tags',
                  titleKey: 'datasets.category',
                  items: categoryTag,
                  clearFlag: false,
                  styleClass: 'first-style'
              },
              {
                  type: 'tasks',
                  titleKey: 'datasets.task',
                  items: taskTag,
                  clearFlag: false,
                  styleClass: 'second-style'
              },
              {
                  type: 'license',
                  titleKey: 'datasets.license',
                  items: licenseTag,
                  clearFlag: false,
                  styleClass: 'third-style'
              }
            ]
            console.log(this.filtersConfig)
          }else{
            this.$message.error(response.data.msg)
          }
        }catch(error){
          console.log(error)
          this.$message.error(error)
        }
    },
    deleteEvent(data){
      console.log("ssssssssssss",data)
      this.dataObj = data
      this.$refs.deleteModal.showModal()
    },
    async handleDeleteModel({ type, data, callback }) {
      try {
        const response = await delDataset({ dataset_id: data.id })
        
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
      this.getCardList() // 刷新列表
    },
  },
  beforeMount() {
    
  },
  mounted() {
    
    console.log("isWorkspacePage",isWorkspacePage)
    if(isWorkspacePage){
      this.params.url = '/owned'
    }else{
      this.params.url = '/public?recommend=all'
      this.getDatasetTags()
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
  /* 遮罩层样式（不覆盖按钮） */
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.2);
    z-index: 100; /* 低于按钮的z-index */
    display: flex;
    justify-content: center;
    align-items: flex-start;
    /* 弹窗样式 */
    .filter-panel {
      width: 100%;
      margin: 0 14px;
      margin-top: 120px;
      box-shadow: 0px 4px 4px 0px rgba(16,16,16,0.1);
      border-radius: 5px;
      background-color: rgba(250,250,255,1);
      border: 1px solid rgba(16,16,16,0.2);
      height: 480px;
      z-index: 101; /* 高于遮罩层 */
      .filter-section{
        display: flex;
        flex-direction: column;
        height: 100%;
        .filter-content{
          flex: 1;
          overflow: auto;
          margin: 24px 8px 0 16px;
          .block-c{
            margin-top: 8px;
            .title-c {
              display: flex;
              align-items: flex-end;
              margin-bottom: 8px;

              .title {
                color: rgb(16, 16, 16);
                font-size: 14px;
              }
            }
            .list {
              display: flex;
              flex-wrap: wrap;
              .item {
                border-radius: 4px;
                font-size: 12px;
                color: rgba(14, 37, 69, 1);
                box-shadow: 0px 1px 1px 0px rgba(16, 16, 16, 0.2);
                margin: 0 10px 10px 0;
                padding: 2px 4px;
                display: flex;
                align-items: center;
                justify-content: center;
                cursor: pointer;

                /* 默认样式 */
                border: 1px solid rgba(145, 213, 255, 0.5);
                &.active {
                  background-color: #0066ff;
                  color: #fff;
                }

                /* 不同样式类型 */
                // &.first-style {
                //     border-color: rgba(0, 158, 255, 1);
                // }

                &.second-style {
                  border: 1px solid rgba(255, 198, 145, 0.5);
                  &.active {
                    background-color: #0066ff;
                    color: #fff;
                  }
                }
                &.third-style {
                  border: 1px solid rgba(169, 223, 184, 0.5);
                  &.active {
                    background-color: #0066ff;
                    color: #fff;
                  }
                }
              }
            }
          }
        }
        .filter-buttons{
          height: 60px;
          display: flex;
          align-items: center;
          justify-content: center;
          .btn{
            cursor: pointer;
            display: flex;
            align-items: center;
            justify-content: center;
            height: 40px;
            border-radius: 4px;
            font-size: 14px;
          }
          .reset-button{
            width: 90px;
            background-color: rgba(255,255,255,1);
            color: rgba(16,16,16,0.75);
            border: 1px solid rgba(16,16,16,0.5);
            margin-right: 20px;
            
          }
          .search-button{
            width: 171px;
            background-color: rgba(22,132,252,1);
            color: rgba(255,255,255,1);
          }
        }
      }
    }
  }

  
}
@media only screen and (max-width: 849.99px) {
  .content {
    padding-left: 0 !important;
  }
}
/* 通用滚动条样式 */
::-webkit-scrollbar {
  width: 4px;
  height: 4px;
}
::-webkit-scrollbar-track {
  background: #f1f1f1;
}
::-webkit-scrollbar-thumb {
  background: #88888880;
  border-radius: 3px;
}
/deep/.el-dropdown-menu__item.active {
  color: #409EFF;
  background-color: rgba(179, 216, 255, 0.3);
}

</style>
