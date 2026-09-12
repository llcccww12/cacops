<template>
  <div class="profile-model" v-loading="loading">
    <Filters
      :placeholder="$t('org.searching')"
      :search-btn-text="$t('repos.search')"
      :sort-options="sortList"
      :default-sort-key="order_by"
      @search="onFilterSearch"
    />

    <List v-if="currentList.length" :params="currentList" />
    <Empty v-if="(!currentList.length && !loading)" />

    <div class="center" v-show="currentList.length">
      <el-pagination
        ref="paginationRef"
        background
        @current-change="currentChange"
        @size-change="sizeChange"
        :current-page.sync="page"
        :pager-count="5"
        :page-sizes="page_sizes"
        :page-size.sync="page_size"
        layout="total, prev, pager, next, jumper"
        :total="total"
      />
    </div>
  </div>
</template>

<script>
import List from '../components/List.vue';
import Empty from '../components/Empty.vue';
import Filters from '../components/Filters.vue';
import { getProfileModel, getProfileModelPublic } from "~/apis/modules/modelmanage";
import { formatDate } from 'element-ui/lib/utils/date-util';
import { transFileSize, getListValueWithKey } from '~/utils';
import { MODEL_ENGINES } from '~/const';

export default {
  components: { List, Empty, Filters },
  data() {
    return {
      keyword: '',
      order_by: 'default',
      page: 1,
      page_sizes: [10],
      page_size: 10,
      total: 0,
      currentList: [],
      isLogin: false,
      sortList: [
        { key: 'default', label: this.$t('datasets.default') },
        { key: 'recentupdate', label: this.$t('datasets.recentupdate') },
        { key: 'newest', label: this.$t('datasets.newest') },
        { key: 'downloadcount', label: this.$t('datasets.downloadtimes') },
        { key: 'collections', label: this.$t('datasets.moststars') },
        { key: 'usecount', label: this.$t('datasets.mostusecount') },
        { key: 'derivativecount', label: this.$t('modelManage.mostDerivative') },
        // { key: 'size', label: this.$t('modelManage.total_size') },
        // { key: 'size_asc', label: this.$t('modelManage.total_size_asc') },
        { key: 'alias_asc', label: this.$t('datasets.alphabetasc') },
        { key: 'alias', label: this.$t('datasets.alphabetdesc') },
      ],
      loading: false
    }
  },
  methods: {
    getList() {
      this.loading = true;
      const params = {
        q: this.keyword || undefined,
        order_by: this.order_by,
        page: this.page,
        page_size: this.page_size,
        owner_name: window.__INITIAL_STATE__?.ownerName || 'default_user'
      }

      const request = this.isLogin ? getProfileModel(params) : getProfileModelPublic(params)

      request.then((res) => {
        this.loading = false
        res = res.data
        if (res.code == 0) {
          const list = res.data.aimodels || []
          this.currentList = list
          .filter((item) => item.Owner && item.Owner.IsOrganization === false)
          .map((item) => ({
            ...item,
            updated_time: this.formatTime(item.updated_unix),
            size: transFileSize(item.size),
            tags: this.parseTags(item.tags),
            engineName: getListValueWithKey(MODEL_ENGINES, item.engine),
            licenses: '',
            type: 'model'
          }))
          this.total = res.data.total
        } else {
          this.resetPagination();
        }
      }).catch((err) => {
        console.error(err)
        this.loading = false;
        this.currentList = []
        this.resetPagination();
      })
    },
    parseTags(tags) {
      return (tags || '')
        .split(' ')
        .map(t => t.trim())
        .filter(t => t !== '');
    },
    resetPagination() {
      this.keyword = ''
      this.order_by = 'default'
      this.page = 1
      this.page_size = 10
      this.total = 0
    },
    formatTime(time) {
      return formatDate(new Date(time * 1000), 'yyyy-MM-dd')
    },
    onFilterSearch({ keyword, order_by }) {
      this.keyword = keyword
      this.order_by = order_by
      this.page = 1
      this.getList()
    },
    currentChange(page) {
      this.page = page
      this.getList()
    },
    sizeChange(pageSize) {
      this.page_size = pageSize
      this.page = 1
      this.getList()
    }
  },
  mounted() {
    this.isLogin = !!document.querySelector('meta[name="_uid"]');
    this.getList()
  }
}
</script>

<style scoped lang="less">
.profile-model {
  display: flex;
  justify-content: center;
  padding-top: 0;
  flex-direction: column;
}

.center {
  text-align: center;
  margin-top: 10px;
}
</style>