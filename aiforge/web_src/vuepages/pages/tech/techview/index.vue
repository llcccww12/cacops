<template>
  <div>
    <TopHeader :menu="'tech_view'"></TopHeader>
    <div class="ui container">
      <SearchBar :type="0" :condition="condition" @changeCondition="changeCondition"></SearchBar>
      <div class="conent-c">
        <div class="filter-c">
          <Filters :type="0" :condition="condition" @changeCondition="changeCondition"></Filters>
        </div>
        <div class="result-c">
          <SciAndTechPrjList ref="resultListRef" :condition="condition" @changeCondition="changeCondition">
          </SciAndTechPrjList>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TopHeader from '../components/TopHeader.vue';
import SearchBar from '../components/SearchBar.vue';
import Filters from '../components/Filters.vue';
import SciAndTechPrjList from '../components/SciAndTechPrjList.vue';
import { getUrlSearchParams } from '~/utils';

export default {
  data() {
    return {
      condition: {
        q: '', // 搜索框
        type_name: '', // 项目类型
        institution_name: '', // 项目参与单位
        execute_year: '', // 执行周期包含年份
        apply_year: '',
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
    SciAndTechPrjList
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
      window.location.href = `/tech/tech_view?` + `q=${encodeURIComponent(this.condition.q.trim())}` +
        `&type_name=${encodeURIComponent(this.condition.type_name)}` +
        `&institution_name=${encodeURIComponent(this.condition.institution_name)}` +
        `&execute_year=${encodeURIComponent(this.condition.execute_year)}` +
        `&apply_year=${encodeURIComponent(this.condition.apply_year)}` +
        `&page=${encodeURIComponent(this.condition.page)}` +
        `&pageSize=${encodeURIComponent(this.condition.pageSize)}` +
        `&sort=${encodeURIComponent(this.condition.sort)}`;
    },
  },
  beforeMount() {
    const urlParams = getUrlSearchParams();
    this.condition.q = urlParams.q || '';
    this.condition.type_name = urlParams.type_name || '';
    this.condition.institution_name = urlParams.institution_name || '';
    this.condition.execute_year = urlParams.execute_year || '';
    this.condition.apply_year = urlParams.apply_year || '';
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
}
</style>
