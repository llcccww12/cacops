<template>
  <div class="profile-compute" v-loading="loading">
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
import { getProfileAITaskTemplate, getProfileAITaskTemplatePublic } from "~/apis/modules/aitasktmpl";
import { formatDate } from 'element-ui/lib/utils/date-util';
import { getListValueWithKey } from '~/utils';
import { TmplTaskTypes, TmplComputerResouces, getGradientColor } from '~/pages/aitasktmpl/tools';

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
        { key: 'collections', label: this.$t('datasets.moststars') },
        { key: 'usecount', label: this.$t('taskTmplObj.runTimes') },
        { key: 'name_asc', label: this.$t('datasets.alphabetasc') },
        { key: 'name', label: this.$t('datasets.alphabetdesc') },
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

      const request = this.isLogin ? getProfileAITaskTemplate(params) : getProfileAITaskTemplatePublic(params)

      request.then((res) => {
        this.loading = false;
        res = res.data
        if (res.code == 0) {
          const list = res.data.Templates || []
          this.currentList = list.map((item) => {
            const datasets = item.DatasetLists || [];
            const datasetsStr = datasets.map(item => {
              return `${item.OwnerName ? (item.OwnerName + '/') : ''}${item.DatasetAlias || item.DatasetName}`
            }).join(', ');
            const models = item.ModelLists || [];
            const modelsStr = models.map(item => {
              return `${item.OwnerName ? (item.OwnerName + '/') : ''}${item.ModelAlias || item.ModelName}`
            }).join(', ');
            return {
              ...item,
              alias: item.Name,
              updated_time: this.formatTime(item.UpdatedUnix),
              recommend: item.Recommend,
              is_collected: item.IsCollected,
              num_stars: item.NumCollections,
              use_count: item.UseCount,
              tags: item.Tags,
              DatasetsStr: datasetsStr,
              ModelsStr: modelsStr,
              JobTypeStr: getListValueWithKey(TmplTaskTypes, item.JobType),
              ComputeSourceStr: getListValueWithKey(TmplComputerResouces, item.ComputeSource),
              type: 'aitasktmpl',
              is_private: item.IsPrivate
            }
          })
          this.total = res.data.Total
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
.profile-compute {
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