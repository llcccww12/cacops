<template>
  <div class="list-container">
    <div class="list-head">
      <div class="title">{{ $t('repos.dataset') }}</div>
      <div class="title-r">
        <div class="search-c">
          <el-input v-model="q" :placeholder="$t('org.searchDatasets')" @keyup.native.enter="search">
            <el-button slot="append" class="search-btn" @click="search">{{ $t('repos.search') }}</el-button>
          </el-input>
        </div>
        <el-dropdown class="sort-c" trigger="click" size="default">
          <span class="el-dropdown-link">
            {{ $t('datasets.sort') }}<i class="el-icon-caret-bottom el-icon--right"></i>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item :class="sort == item.key ? 'active' : ''" v-for="item in sortList" :key="item.key"
              @click.native="changeSort(item)">
              {{ item.label }}
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
    <div class="list-item-container" v-loading="loading">
      <div class="item-container" v-for="(item, index) in list" :key="item.id">
        <DatasetItem :data="item" :canChangeFav="true" @changeFav="onChangeFav"></DatasetItem>
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
        layout="total, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </div>
  </div>
</template>

<script>
import DatasetItem from '~/components/square/RightItem.vue';
import { getOrgDatasetList, getOrgLabel } from '~/apis/modules/organization';
import { formatDate } from 'element-ui/lib/utils/date-util';

export default {
  name: "DatasetList",
  props: {
    conds: { type: Object, default: () => ({}) },
    pageData: { type: Object, default: () => { } },
  },
  components: { DatasetItem },
  data() {
    return {
      loading: false,
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
      list: [],
      labels: [],
      iPageSizes: [12],
      iPageSize: 12,
      iPage: 1,
      total: 0,
    };
  },
  methods: {
    onChangeFav(id){
      this.$emit('refreshFav',id);
    },
    getListData() {
      this.loading = true;
      console.log("this.conds",this.conds)
      const filter = this.conds.label || [];

      const requestParams = {
        owner_name: this.pageData.Org.Name,
        q: this.q.trim(),
        tasks: '',
        tags: '',
        license: '',
        recommend: '',
        order_by: this.sort,
        page: this.iPage,
        page_size: this.iPageSize,
      };
      filter.forEach(item => {
        const [type, value] = item.split('|'); 
        
        if (type === 'tags' && value) {
          requestParams.tags = value;
        } else if (type === 'tasks' && value) {
          requestParams.tasks = value;
        } else if (type === 'license' && value) {
          requestParams.license = value;
        }
      });
      getOrgDatasetList(requestParams).then(res => {
        res = res.data;
        this.loading = false;
        if (res.code === 0) {
          const list = res.data.datasets || [];
          console.log("xxxxxxxxx",list)
          this.list = list.map(item => {
            return {
              ...item,
              updateTimeStr: formatDate(new Date(item.updated_unix * 1000), 'yyyy-MM-dd'),
            }
          })
          this.total = res.data.total || 0;
        } else {
          this.list = [];
        }
      }).catch(err => {
        console.log(err);
        this.loading = false;
        this.list = [];
        this.total = 0;
      });
    },
    getLabelList(){
      getOrgLabel({ orgName: this.pageData.Org.Name}).then((res)=>{
        res = res.data
        if(res.code === 0){
          console.log(res.data)
          const orgTopics = res.data;
          const category = (orgTopics.Tag || []).filter(item => item.trim() != '').map(item => {
            return {
              k: [`tags|${item}`],
              v: this.$t('datasets.' + item),
            }
          })
          const task = (orgTopics.Task || []).filter(item => item.trim() != '').map(item => {
            return {
              k: [`tasks|${item}`],
              v: this.$t('datasets.' + item),
            }
          })
          const license = (orgTopics.License || []).filter(item => item.trim() != '').map(item=>{
            return{
              k: [`license|${item}`],
              v: item,
            }
          })
          this.labels = [...category, ...task, ...license];
          this.$emit('update-labels', [...this.labels]);
        }
      }).catch(err => {
        console.log(err);
        
      });
    },
    search() {
      this.iPage = 1;
      this.getListData();
    },
    changeSort(item) {
      this.sort = item.key;
      this.search();
    },
    currentChange(page) {
      this.iPage = page;
      this.getListData();
    },
    sizeChange(pageSize) {
      this.iPageSize = pageSize;
      this.search();
    },
  },
  watch: {
    conds: {
      handler(newVal) {
        this.search();
      },
      deep: true,
    },
  },
  mounted() {
    this.search();
    this.getLabelList()
  },
};
</script>

<style scoped lang="less">
.list-container {

  .list-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 6px;

    .title {
      color: rgb(16, 16, 16);
      font-size: 24px;
      font-family: SourceHanSansSC;
      font-weight: 500;
      line-height: 35px;
    }

    .title-r {
      display: flex;
      align-items: center;

      .search-c {
        margin-right: 16px;

        .el-input {
          .el-button {
            color: rgb(16, 16, 16);
          }
        }
      }
    }
  }

  .list-item-container {
    display: flex;
    flex-wrap: wrap;
    min-height: 120px;

    .item-container {
      width: 33.3%;
      padding: 12px;

      &:nth-child(3n+1) {
        padding-left: 0;
      }

      &:nth-child(3n) {
        padding-right: 0;
      }
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

.sort-c {
  color: rgb(16, 16, 16);
  cursor: pointer;
}

/deep/.el-dropdown-menu__item.active {
  color: #409EFF;
  background-color: rgba(179, 216, 255, 0.3);
}

@media only screen and (max-width: 767px) {

  /deep/ .el-pagination__total,
  /deep/ .el-pagination__jump {
    display: none !important;
  }

  .list-container {
    .list-head {
      .title {
        font-size: 16px;
        margin-right: 8px;
      }

      .title-r {
        width: 0;
        flex: 1;

        .search-c {
          width: 0;
          flex: 1;
          margin-right: 8px;
        }
      }
    }

    .list-item-container {
      .item-container {
        width: 100% !important;
        padding: 8px !important;
      }
    }
  }
}
@media only screen and (min-width: 767px) and (max-width: 1440px) {
  .list-item-container {
    .item-container {
      width: 50% !important;

      &:nth-child(2n+1) {
        padding-left: 0 !important;
        padding-right: 12px !important;
      }

      &:nth-child(2n) {
        padding-left: 12px !important;
        padding-right: 0 !important;
      }
    }
  }
}
</style>
