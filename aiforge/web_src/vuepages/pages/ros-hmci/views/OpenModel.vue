<template>
  <div>
    <AppBanner></AppBanner>
    <div class="search-bar-wrap">
      <div class="search-bar">
        <input type="text" v-model="condition.q" :placeholder="$t('modelObj.model_search')"
          @keyup.enter="conditionChange" />
        <button @click="conditionChange">{{ $t('repos.search') }}</button>
      </div>
    </div>
    <div class="ui container">
      <div class="content">
        <div class="content-l">
          <ModelFilters :condition="condition" @changeCondition="conditionChange"></ModelFilters>
        </div>
        <div class="content-r">
          <ModelCondition :condition="condition" @changeCondition="conditionChange"></ModelCondition>
          <ModelList ref="modelListRef" :condition="condition" @changeCondition="conditionChange"></ModelList>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AppBanner from "../components/AppBanner.vue";
import ModelCondition from '~/pages/modelsquare/square/components/ModelCondition.vue';
import ModelFilters from '~/pages/modelsquare/square/components/ModelFilters.vue';
import ModelList from '~/pages/modelsquare/square/components/ModelList.vue';
import { getUrlSearchParams } from '~/utils';

export default {
  data() {
    return {
      condition: {
        q: '',
        tab: '1',
        sort: '',
        onlyRecommend: false,
        engine: '',
        label: '',
        page: 1,
        pageSize: 30,
      },
      pageSizes: [30, 50],
    };
  },
  components: { AppBanner, ModelFilters, ModelCondition, ModelList },
  methods: {
    conditionChange(params = {}) {
      this.condition = {
        ...this.condition,
        ...params,
      };
      if (!params.changePage) {
        this.condition.page = 1;
      }
      window.location.href = `/explore/models?` + `q=${encodeURIComponent(this.condition.q.trim())}` +
        `&tab=${encodeURIComponent(this.condition.tab)}` +
        `&sort=${encodeURIComponent(this.condition.sort)}` +
        `&onlyRecommend=${encodeURIComponent(this.condition.onlyRecommend)}` +
        `&engine=${encodeURIComponent(this.condition.engine)}` +
        `&label=${encodeURIComponent(this.condition.label)}` +
        `&page=${encodeURIComponent(this.condition.page)}` +
        `&pageSize=${encodeURIComponent(this.condition.pageSize)}`;
    }
  },
  beforeMount() {
    const urlParams = getUrlSearchParams();
    this.condition.q = urlParams.q || '';
    this.condition.tab = urlParams.tab || '1';
    this.condition.sort = urlParams.sort || '';
    this.condition.onlyRecommend = urlParams.onlyRecommend == 'true' ? true : false;
    this.condition.engine = urlParams.engine || '';
    this.condition.label = urlParams.label || 'ros-hmci-models';//默认搜索ros-hmci-models标签

    this.condition.page = Number(urlParams.page) || 1;
    this.condition.pageSize = this.pageSizes.indexOf(Number(urlParams.pageSize)) >= 0 ? Number(urlParams.pageSize) : 30;
    this.$nextTick(() => {
      this.$refs.modelListRef.search();
    });
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.search-bar-wrap {
  height: 88px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f5f5f6;

  .search-bar {
    height: 40px;
    font-size: 14px;
    text-align: left;
    display: flex;

    input {
      width: 602px;
      height: 100%;
      padding: 0 8px;
      outline: none;
      border-color: #bbbbbb;
      border-width: 1px;
      border-style: solid;
      border-radius: 5px 0px 0px 5px;
      background-color: rgba(0,0,0,0);

      &:focus {
        border-color: #85b7d9;
        -webkit-box-shadow: 0 0 0 0 rgba(34, 36, 38, .35) inset;
        box-shadow: 0 0 0 0 rgba(34, 36, 38, .35) inset
      }
    }

    button {
      padding: 0 20px;
      height: 40px;
      color: #ffffff;
      border-radius: 0px 4px 4px 0px;
      background: #5bb973;
      border: none;
      cursor: pointer;
    }
  }
}

.content {
  display: flex;

  .content-l {
    flex: 1;
    padding-left: 12px;
  }

  .content-r {
    margin-left: 10px;
    flex: 3;
  }
}
</style>
