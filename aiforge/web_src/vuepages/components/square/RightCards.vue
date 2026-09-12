<template>    
<div class="content-r">
  <div class="filter-c-m">
    <span class="title">{{ type === 'dataset' ? $t('dataset') : $t('repos.model')}}</span>
    <div class="right-m">
      <div class="icon-box" @click="createDataset"><i class="ri-add-box-line"></i></div>
      <div class="icon-box" v-if="!showSearch" @click="showSearch = true"><i class="ri-search-line"></i></div>
      <div v-else  class="ui small icon input" style="height: 32px;margin-left: 6px;">
        <input type="text" :placeholder="type === 'dataset' ? $t('org.searchDatasets') : $t('modelObj.model_search')" v-model="q" @keyup.enter="search">
        <i class="search icon" style="cursor: pointer;pointer-events: auto;" @click="search"></i>
      </div> 
      <div v-if="!isWorkspacePage" class="icon-box filter-btn" :class="{ 'active': showFilter }" @click="toggleFilter">
        <i class="ri-filter-line"></i>
      </div>
    </div>
  </div>
  <div v-if="isWorkspacePage" class="workspace-t">
    <span>{{ $t(title) }}</span>
  </div>
  <Recommend v-if="!isWorkspacePage" :type="type"></Recommend>
  <div class="filter-c">
    <div class="tab-c">
        <div class="tab-item nowrap" :class="tabIndex == item.key ? 'active' : ''"
        v-for="(item) in tabList" :key="item.key" @click="changeTab(item)">{{
            item.label }}</div>
    </div>
    <div class="right">
      <div class="ui small icon input" style="height: 32px;">
        <input type="text" :placeholder="type === 'dataset' ? $t('org.searchDatasets') : $t('modelObj.model_search')" v-model="q" @keyup.enter="search">
        <i class="search icon" style="cursor: pointer;pointer-events: auto;" @click="search"></i>
      </div> 
      <el-select v-model="sort" placeholder="请选择" @change="changeSort" style="width: 120px;">
        <el-option
          v-for="item in mergedSortList"
          :key="item.key"
          :label="item.label"
          :value="item.key">
        </el-option>
      </el-select>
      <div class="btn-add" @click="createDataset"> 
        <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 32 32" width="20" height="20"><defs></defs><g><path d="M5.333 4h21.333c0.736 0 1.333 0.597 1.333 1.333v0 21.333c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-21.333c0-0.736 0.597-1.333 1.333-1.333v0zM6.667 6.667v18.667h18.667v-18.667h-18.667zM14.667 14.667v-5.333h2.667v5.333h5.333v2.667h-5.333v5.333h-2.667v-5.333h-5.333v-2.667h5.333z"></path></g></svg>
        <span style="margin-left:6px;">{{ createBtnText  }}</span>
      </div>
    </div>
    <div class="right-m">
      <div class="icon-box" @click="createDataset"><i class="ri-add-box-line"></i></div>
      <div class="icon-box" v-if="!showSearch" @click="showSearch = true"><i class="ri-search-line"></i></div>
      <div v-else  class="ui small icon input" style="height: 32px;margin-left: 6px;">
        <input type="text" :placeholder="type === 'dataset' ? $t('org.searchDatasets') : $t('modelObj.model_search')" v-model="q" @keyup.enter="search">
        <i class="search icon" style="cursor: pointer;pointer-events: auto;" @click="search"></i>
      </div> 
    </div>
  </div>
  <right-list :list="list" :type="type" :pageParams="pageParams" :canChangeFav="canChangeFav" :reloadFlag="reloadFlag" 
    @changePage="changePage" v-loading="loading" :isWorkspacePage="isWorkspacePage" :operaFlag="operaFlag" @deleteEvent="deleteEvent"></right-list>
</div>
</template>

<script>
import Recommend from './Recommend.vue';
import RightList from './RightList.vue';
export default {
  props: {
    tabList: {
      type: Array,
      default: () => []
    },
    extraSorts: {
      type: Array,
      default: () => []
    },
    showCheckbox: {
      type: Boolean,
      default: false
    },
    isWorkspacePage: {
      type: Boolean,
      default: false
    },
    loading: {
      type: Boolean,
      default: false
    },
    tab: {
      type: String,
      default: ''
    },
    list: {
      type: Array,
      default: () => []
    },
    pageParams: { 
      type: Object,
      default: () => ({}) 
    },
    type: {
      type: String,
      default: 'dataset'
    },
    title: { 
      type: String, 
      default: '' 
    },
  },
  data() {
    return {
        q: '',
        baseSorts: [{
            key: '',
            label: this.$t('datasets.default'),
        }, {
            key: 'newest',
            label: this.$t('datasets.newest'),
        }, {
            key: 'recentupdate',
            label: this.$t('datasets.recentupdate'),
        }, {
            key: 'downloadcount',
            label: this.$t('datasets.downloadtimes'),
        }, {
            key: 'collections',
            label: this.$t('datasets.moststars'),
        }, {
            key: 'alias_asc',
            label: this.$t('datasets.alphabetasc'),
        }, {
            key: 'alias',
            label: this.$t('datasets.alphabetdesc'),
        }],
        tabIndex: this.tab,
        sort: '',
        hasOnlineUrl: false,
        // 筛选条件值集合
        selectedFilters: {},
        showSearch: false,
        showFilter: false,
        createUrl: '',
    };
  },
  components: { RightList, Recommend },
  computed: {
    mergedSortList() {
      return [...this.baseSorts, ...this.extraSorts]
    },
    canChangeFav(){
      const data = this.tabList.find((item)=>{
        return item.key === this.tabIndex
      })
      return data.favFlag
    },
    reloadFlag(){
      const data = this.tabList.find((item)=>{
        console.log(item.key,this.tabIndex)
        return item.key === this.tabIndex
      })
      return !!data.reloadFlag
    },
    operaFlag(){
      const data = this.tabList.find((item)=>{
        return item.key === this.tabIndex
      })
      return data.operaFlag
    },
    createBtnText(){
      const data = this.tabList.find((item)=>{
        return item.key === this.tabIndex
      })
      if(data.btnText){
        this.createUrl = data.btnUrl
        return data.btnText
      }else{
        if(this.type==='dataset'){
          return this.$t('datasetObj.create_new_dataset')
        }else{
          return this.$t('modelManage.createNewModel')
        }
      }
    },
  },
  watch: {
    q: {
      handler(val, oval) {
        if(val===''){
            this.selectedFilters.q = ''
            this.changeParams()
        }
      },
      deep: true,
    }
  },
  methods: {
    toggleFilter(){
      this.showFilter = !this.showFilter;
      this.$emit('changeModalOverlay',this.showFilter)
    },
    resetFilter(){
      this.showFilter = false
    },
    parseQueryString(key) {
      let path = key.split('?')[0];
      const queryString = key.includes('?') ? key.split('?')[1] : '';
      const queryParams = {};
      if (queryString){
        const params = new URLSearchParams(queryString);
        console.log("xxxxxxxxx1",params)
        params.forEach((value, key) => {
          // 处理布尔值（如 recommend=false → false）
          console.log("xxxxxxxxx2",value,key)
          if (value === 'true' || value === 'false') {
            queryParams[key] = value === 'true';
          } 
          // 处理数字（如 page=2 → 2）
          else if (!isNaN(value) && !isNaN(parseFloat(value))) {
            queryParams[key] = Number(value);
          }
          // 其他情况直接赋值
          else {
            queryParams[key] = value;
          }
        })
      }
      return {
        url: path, // 覆盖默认路径
        ...queryParams, // 动态覆盖查询参数
      };
    },
    changeTab(item){
        this.tabIndex = item.key
        this.selectedFilters = this.parseQueryString(item.key)
        console.log("this.selectedFilters",this.selectedFilters,item)
        this.selectedFilters.page = 1
        let index = this.tabList.findIndex((item)=>{
          return item.key === this.tabIndex
        })
        if(index==0){
          this.selectedFilters.owner_type = 'individual'
        }else{
          this.selectedFilters.owner_type = ''
        }
        this.changeParams()
    },
    changeSort(item){
      console.log(item)
        // this.sort = item
        this.selectedFilters.order_by = item
        this.changeParams()
    },
    changeOnline(){
        this.selectedFilters.hasOnlineUrl = this.hasOnlineUrl
        this.changeParams()
    },
    search(){
      this.selectedFilters.page = 1
      this.selectedFilters.q = this.q
      this.changeParams()
    },
    changePage(item){
        this.selectedFilters.page = item
        this.changeParams()
    },
    changeParams(){
        console.log("this.selectedFilters",this.selectedFilters)
        this.$emit('changeParams', this.selectedFilters)
    },
    deleteEvent(data){
      this.$emit("deleteEvent", data)
    },
    createDataset(){
      if(this.createUrl){
        location.href = this.createUrl
      }else{
        if(this.type==='dataset'){
          location.href = `/guide/create_dataset`
        }else{
          location.href = `/guide/create_model`
        }
      }
      
    }
  },
  beforeMount() {
  },
  mounted() {

  },
  beforeDestroy() { },
};
</script>
<style scoped lang="less">

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
}

/deep/.el-dropdown-menu__item.active {
  color: #409EFF;
  background-color: rgba(179, 216, 255, 0.3);
}

@media only screen and (max-width: 849.99px) {
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
    }
  }


  /deep/ .list-item-container {
    .item-container {
      width: 100% !important;
      padding: 12px 0 !important;
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
  }
  
}
</style>
