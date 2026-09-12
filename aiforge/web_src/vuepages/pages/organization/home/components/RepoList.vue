<template>
  <div class="list-container">
    <div class="list-head">
      <div class="title">{{ $t('repos.repos') }}</div>
      <div class="title-r">
        <div class="search-c">
          <el-input v-model="q" :placeholder="$t('org.searchRepos')" @keyup.native.enter="search">
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
      <div class="item-container" v-for="(item, index) in list" :key="item.ID">
        <ReposItem :data="item" :key="item.ID"></ReposItem>
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
        :current-page.sync="iPage" :page-sizes="iPageSizes" :page-size.sync="iPageSize"
        layout="total, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </div>
  </div>
</template>

<script>
import ReposItem from '~/pages/repos/square/components/Item.vue';
// import ReposItem from '~/pages/repos/components/ReposItem.vue';
import { getOrgRepoList } from '~/apis/modules/organization';
import LetterAvatar from '~/utils/letteravatar';

export default {
  name: "ReposList",
  props: {
    conds: { type: Object, default: () => ({}) },
    pageData: { type: Object, default: () => { } },
  },
  components: { ReposItem },
  data() {
    return {
      loading: false,
      q: '',
      sort: 'recentupdate',
      sortList: [{
        key: 'newest',
        label: this.$t('org.order_newest'),
      }, {
        key: 'oldest',
        label: this.$t('org.order_oldest'),
      }, {
        key: 'alphabetically',
        label: this.$t('org.order_alphabetically'),
      }, {
        key: 'reversealphabetically',
        label: this.$t('org.order_reversealphabetically'),
      }, {
        key: 'recentupdate',
        label: this.$t('org.order_recentupdate'),
      }, {
        key: 'leastupdate',
        label: this.$t('org.order_leastupdate'),
      }, {
        key: 'moststars',
        label: this.$t('org.order_moststars'),
      }, {
        key: 'feweststars',
        label: this.$t('org.order_feweststars'),
      }, {
        key: 'mostforks',
        label: this.$t('org.order_mostforks'),
      }, {
        key: 'fewestforks',
        label: this.$t('org.order_fewestforks'),
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
    getListData() {
      this.loading = true;
      getOrgRepoList({
        orgName: this.pageData.Org.Name,
        q: this.q.trim(),
        label: this.conds.label || '',
        sort: this.sort || 'mostpopular',
        page: this.iPage || 1,
        pageSize: this.iPageSize || 15
      }).then(res => {
        res = res.data;
        this.loading = false;
        if (res.code == 0) {
          const list = res.data || [];
          this.list = list.map((item) => {
            item.Contributors = (item.Contributors || []).map((_item) => {
              return {
                ..._item,
                bgColor: this.randomColor((_item.Email[0] || '').toLocaleUpperCase()),
              }
            });
            return {
              ...item,
              NameShow: item.Alias, // this.handlerSearchStr(item.Alias, this.conds.q),
              DescriptionShow: item.Description, // this.handlerSearchStr(item.Description, this.conds.q),
              TopicsShow: (item.Topics || []).map((_item) => {
                return {
                  topic: _item,
                  topicShow: _item, // this.handlerSearchStr(_item, this.conds.q)
                }
              }),
            }
          });
          this.total = res.total;
          this.iPage = this.iPage;
          this.iPageSize = this.iPageSize;
          this.$nextTick(() => {
            LetterAvatar.transform();
          });
          if (!this.labels.length) {
            this.labels = (res.orgTopics || []).map(item => {
              return {
                k: item.Name,
                v: item.Name,
                ...item,
              }
            });
            this.$emit('update-labels', [...this.labels]);
          }
        } else {
          this.list = [];
          this.total = 0;
        }
      }).catch(err => {
        console.log(err);
        this.loading = false;
        this.list = [];
        this.total = 0;
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
    handlerSearchStr(oStr, searchKey) {
      if (!searchKey) return oStr;
      return oStr.replace(new RegExp(`(${searchKey})`, 'ig'), `<font color="red">$1</font>`);
    },
    randomColor(t) {
      const tIndex = t.charCodeAt(0);
      const colorList = ["#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#34495e", "#16a085", "#27ae60", "#2980b9", "#8e44ad",
        "#2c3e50", "#f1c40f", "#e67e22", "#e74c3c", "#00bcd4", "#95a5a6", "#f39c12", "#d35400", "#c0392b", "#bdc3c7", "#7f8c8d"];
      return colorList[tIndex % colorList.length];
    }
  },
  watch: {
    conds: {
      handler(newVal) {
        this.search();
      },
      deep: true,
    }
  },
  mounted() {
    this.search();
  },
};
</script>

<style scoped lang="less">
.list-container {
  margin-top: 16px;

  .list-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 16px;

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
    .item-container {
      width: 100%;
      // margin-bottom: -20px;
      
    }
  }
}

.center {
  text-align: center;
}


.no-data {
  display: flex;
  justify-content: center;
  padding: 12px 0;
  width: 100%;
  margin-top: -10px;

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

@media only screen and (max-width: 850px) {

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

  }
  
}
@media only screen and (max-width: 1200px) {
  .list-item-container {
    .item-container{
      margin-bottom: 20px;
    }
    
  }
}
@media only screen and (min-width: 1200px) {
  .list-item-container {
    .item-container {
      width: 50% !important;
      padding: 14px;
      // margin-bottom: -40px;
      &:nth-child(odd) {
        padding-left: 0;
      }

      &:nth-child(even) {
        padding-right: 0;
      }
    }
  }
}
</style>
