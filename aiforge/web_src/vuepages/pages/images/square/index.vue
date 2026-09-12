<template>
  <div class="content">
    <div class="content-l" v-if="!isWorkspacePage">
      <Filters @changeCondition="conditionChange"></Filters>
    </div>
    <div class="content-r">
      <div v-if="isWorkspacePage" class="workspace-t">
        <span>{{ $t('imagesObj.cloudbrain_images') }}</span>
      </div>
      <div class="filter-c">
        <div class="tab-c">
          <div class="tab-item" :class="tab == item.key ? 'active' : ''"
            v-for="(item) in (isWorkspacePage ? tabList2 : tabList1)" :key="item.key" @click="changeTab(item)">{{
              item.label }}</div>
        </div>
        <div class="right">
          <div class="search-c">
            <el-input v-model="q" :placeholder="$t('org.searchImages')" @keyup.native.enter="search">
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
      <List :params="params"></List>
    </div>
  </div>
</template>

<script>
import Filters from './components/Filters.vue';
import Condition from './components/Condition.vue';
import List from './components/List.vue';
import { i18n } from '~/langs';

const isWorkspacePage = window.MENU_CONFIG.activeTopMenu == 'workspace'

const SortList = [{
  key: '',
  label: i18n.t('datasets.default'),
}, {
  key: 'moststars',
  label: i18n.t('datasets.mostCollections'),
}, {
  key: 'mostused',
  label: i18n.t('datasets.mostusecount'),
}, {
  key: 'newest',
  label: i18n.t('datasets.newest'),
}];

export default {
  data() {
    return {
      isWorkspacePage: isWorkspacePage,
      tab: isWorkspacePage ? '1' : '0',
      tabList1: [{
        key: '0',
        label: `🏆${this.$t('imagesObj.image_public')}`,
      }],
      tabList2: [{
        key: '1',
        label: this.$t('imagesObj.image_my'),
      }, {
        key: '2',
        label: this.$t('imagesObj.image_collected'),
      }],
      sort: '',
      sortList: [],
      q: '',
      trainType: '',
      compute_resource: '',
      framework: '',
      framework_version: '',
      python: '',
      cuda: '',
      cann: '',
      dtk: '',
      params: {
        q: '',
        tab: '0',
        sort: '',
        trainType: '',
        compute_resource: '',
        framework: '',
        framework_version: '',
        python: '',
        cuda: '',
        cann: '',
        dtk: '',
      },
    };
  },
  components: { Filters, Condition, List },
  methods: {
    changeTab(item) {
      this.tab = item.key;
      if (item.key == '1') {
        this.sort = 'newest';
        this.sortList = SortList.slice(1);
      } else {
        this.sort = '';
        this.sortList = SortList.slice(0);
      }
      this.search()
    },
    changeSort(item) {
      this.sort = item.key
      this.search()
    },
    search() {
      this.params = {
        q: this.q,
        tab: this.tab,
        sort: this.sort,
        trainType: this.trainType,
        compute_resource: this.compute_resource,
        framework: this.framework,
        framework_version: this.framework_version,
        python: this.python,
        cuda: this.cuda,
        cann: this.cann,
        dtk: this.dtk
      }
    },
    conditionChange(params = {}) {
      console.log(params)
      this.trainType = params.trainType;
      this.compute_resource = params.compute_resource;
      this.framework = params.framework;
      this.framework_version = params.framework_version;
      this.python = params.python;
      this.cuda = params.cuda;
      this.cann = params.cann;
      this.dtk = params.dtk;
      this.search()
    },
  },
  beforeMount() {
    if (this.tab == '1') {
      this.sort = 'newest';
      this.sortList = SortList.slice(1);
    } else {
      this.sort = '';
      this.sortList = SortList.slice(0);
    }
  },
  mounted() {
    this.search()
  },
  beforeDestroy() { },
};
</script>
<style scoped lang="less">
.content {
  display: flex;
  min-height: 100%;
  padding-left: 20px;
  .content-l {
    width: 25%;
    max-width: 350px;
    min-width: 250px;
    padding: 30px 40px;
  }

  .content-r {
    flex: 1;
    width: 0;
    padding: 30px 20px;
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
          color: rgba(16, 16, 16, 0.5);

          &:hover {
            opacity: 0.8;
          }

          &.active {
            color: #0066ff;
            border-bottom: 2px solid #0066ff;
            // background-color: rgba(68, 58, 104, 1);
            // color: rgba(255, 255, 255, 1);
            // border: 1px solid rgba(51, 38, 98, 1);

            &:hover {
              opacity: 1;
            }
          }
        }
      }

      .right {
        display: flex;
        align-items: center;

        .check-c {
          margin-right: 16px;
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
    }
  }
}

/deep/.el-dropdown-menu__item.active {
  color: #409EFF;
  background-color: rgba(179, 216, 255, 0.3);
}

@media only screen and (max-width: 767.98px) {
  .content .content-l {
    display: none;
  }

  .content-r .right .search-c {
    display: none !important;
  }

  /deep/ .list-item-container {
    .item-container {
      width: 100% !important;
      padding: 12px 0 !important;
    }
  }
}

@media only screen and (min-width: 768px) and (max-width: 1250px) {
  .content-r .right .search-c {
    // display: none !important;
  }

  /deep/ .list-item-container {
    .item-container {
      width: 100% !important;
      padding: 12px 0 !important;
    }
  }
}

@media only screen and (min-width: 1250px) and (max-width: 1700px) {
  /deep/ .list-item-container {
    .item-container {
      width: 100% !important;

      &:nth-child(n+1) {
        padding-left: 0;
      }

      &:nth-child(n) {
        padding-right: 0;
      }
    }
  }
}

@media only screen and (min-width: 1700px) {
  /deep/ .list-item-container {
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
</style>
