<template>
  <div>
    <TopHeader :menu="'repo_view'"></TopHeader>
    <div class="ui container">
      <SearchBar :type="1" :condition="condition" @changeCondition="changeCondition"></SearchBar>
      <div class="conent-c">
        <div class="filter-c">
          <Filters :type="1" :condition="condition" @changeCondition="changeCondition"></Filters>
        </div>
        <div class="result-c">
          <PrjResultsList ref="resultListRef" :condition="condition" @changeCondition="changeCondition">
          </PrjResultsList>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TopHeader from '../components/TopHeader.vue';
import SearchBar from '../components/SearchBar.vue';
import Filters from '../components/Filters.vue';
import PrjResultsList from '../components/PrjResultsList.vue';
import { getUrlSearchParams } from '~/utils';

export default {
  data() {
    return {
      condition: {
        q: '', // 搜索框
        topic: '', // 关键词
        project_name: '', // 项目类型
        institution_name: '', // 成果贡献单位

        page: '',
        pageSize: '',
        sort: '',
      },
      pageSizes: [15, 30, 50],
    };
  },
  components: {
    TopHeader,
    SearchBar,
    Filters,
    PrjResultsList
  },
  methods: {
    changeCondition(params) {
      this.condition = {
        ...this.condition,
        ...params,
      };
      if (!params.changePage) {
        this.condition.page = 1;
      }
      window.location.href = `/tech/repo_view?` + `q=${encodeURIComponent(this.condition.q.trim())}` +
        `&topic=${encodeURIComponent(this.condition.topic)}` +
        `&project_name=${encodeURIComponent(this.condition.project_name)}` +
        `&institution_name=${encodeURIComponent(this.condition.institution_name)}` +
        `&page=${encodeURIComponent(this.condition.page)}` +
        `&pageSize=${encodeURIComponent(this.condition.pageSize)}` +
        `&sort=${encodeURIComponent(this.condition.sort)}`;
    },
  },
  beforeMount() {
    const urlParams = getUrlSearchParams();
    this.condition.q = urlParams.q || '';
    this.condition.topic = urlParams.topic || '';
    this.condition.project_name = urlParams.project_name || '';
    this.condition.institution_name = urlParams.institution_name || '';
    this.condition.sort = urlParams.sort || '';
    this.condition.page = Number(urlParams.page) || 1;
    this.condition.pageSize = this.pageSizes.indexOf(Number(urlParams.pageSize)) >= 0 ? Number(urlParams.pageSize) : 15;
    this.$nextTick(() => {
      this.$refs.resultListRef.search();
    });
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.conent-c {
  display: flex;
}

.filter-c {
  flex: 1;
  padding-left: 12px;
}

.result-c {
  margin-left: 10px;
  flex: 3;
  width: 0;
}
</style>

