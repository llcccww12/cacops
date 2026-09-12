<template>
  <div class="list-container">
    <div style="min-height:540px;" v-loading="loading">
      <div class="repos-item-container" v-for="(item, index) in list" :key="item.ID">
        <ReposItem :data="item" :topic="topic"></ReposItem>
      </div>
      <div v-show="(!list.length && !loading)" class="repos-no-data">{{ $t('repos.noReposfound') }}</div>
    </div>
    <div class="center">
      <el-pagination ref="paginationRef" background @current-change="currentChange" @size-change="sizeChange"
        :current-page.sync="iPage" :page-sizes="iPageSizes" :page-size.sync="iPageSize"
        layout="total, sizes, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </div>
  </div>
</template>

<script>
import ReposItem from '../components/ReposItem.vue';
import { getReposListData } from '~/apis/modules/repos';
import LetterAvatar from '~/utils/letteravatar';

export default {
  name: "ReposList",
  props: {
    q: { type: String, default: '' },
    sort: { type: String, default: 'mostpopular' },
    topic: { type: String, default: '' },
    page: { type: Number, default: 1 },
    pageSize: { type: Number, default: 15 },
    pageSizes: { type: Array, default: () => [15, 30, 50] }
  },
  components: { ReposItem },
  data() {
    return {
      loading: false,
      list: [],
      iPageSizes: [15, 30, 50],
      iPageSize: 15,
      iPage: 1,
      total: 0,
    };
  },
  methods: {
    getListData() {
      this.loading = true;
      getReposListData({
        q: this.q || '',
        topic: this.topic || '',
        sort: this.sort || 'mostpopular',
        pageSize: this.iPageSize || 15,
        page: this.iPage || 1,
      }).then(res => {
        res = res.data;
        this.loading = false;
        if (res.Code == 0) {
          const list = res.Data.Repos || [];
          this.list = list.map((item) => {
            item.Contributors = (item.Contributors || []).map((_item) => {
              return {
                ..._item,
                bgColor: this.randomColor((_item.Email[0] || '').toLocaleUpperCase()),
              }
            });
            const contributors = item.Contributors || [];
            return {
              ...item,
              NameShow: this.handlerSearchStr(item.Alias, this.q),
              DescriptionShow: this.handlerSearchStr(item.Description, this.q),
              TopicsShow: (item.Topics || []).map((_item) => {
                return {
                  topic: _item,
                  topicShow: this.handlerSearchStr(_item, this.q)
                }
              }),
            }
          });
          this.total = res.Data.Total;
          this.iPage = this.iPage;
          this.iPageSize = this.iPageSize;
          this.$nextTick(() => {
            LetterAvatar.transform();
          });
        } else {
          this.list = [];
          this.total = 0;
          this.iPage = this.iPage;
          this.iPageSize = this.iPageSize;
        }
      }).catch(err => {
        console.log(err);
        this.loading = false;
        this.list = [];
        this.total = 0;
        this.iPage = this.iPage;
        this.iPageSize = this.iPageSize;
      });
    },
    search() {
      this.getListData();
    },
    currentChange(page) {
      this.iPage = page;
      this.$emit('current-change', {
        page: this.iPage,
        pageSize: this.iPageSize,
      });
    },
    sizeChange(pageSize) {
      this.iPageSize = pageSize;
      this.$emit('size-change', {
        page: this.iPage,
        pageSize: this.iPageSize,
      });
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
    page: {
      handler(val) {
        this.iPage = val;
      },
      immediate: true,
    },
    pageSize: {
      handler(val) {
        this.iPageSize = val;
      },
      immediate: true,
    }
  },
  mounted() { },
};
</script>

<style scoped lang="less">
@media only screen and (max-width: 767px) {
  /deep/ .btn-prev, /deep/ .btn-next {
    display: none;
  }
  /deep/ .el-pager {
    display: block;
    margin: 10px 0;
  }
  /deep/ .el-pagination__jump {
    margin-left: 0;
  }
}
.list-container {
  margin-left: 12px;
  margin-right: 12px;
}

.center {
  text-align: center;
}

.repos-no-data {
  height: 60px;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
