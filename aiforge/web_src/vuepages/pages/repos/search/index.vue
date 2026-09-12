<template>
  <div>
    <div class="ui container">
      <SearchBar ref="searchBarRef" type="search" :static="true" :staticTopicsData="staticSquareTopics"
        :sort="reposListSortType" :searchValue="reposListQurey" :topic="reposListTopic" @change="searchBarChange">
      </SearchBar>
    </div>
    <div class="ui container">
      <div class="ui grid">
        <div class="computer only ui two wide computer column">
          <ReposFilters ref="reposFiltersRef" @change="filtersChange"></ReposFilters>
        </div>
        <div class="ui sixteen wide mobile twelve wide tablet ten wide computer column">
          <ReposList ref="reposListRef" :sort="reposListSortType" :q="reposListQurey" :topic="reposListTopic"
            :page="page" :pageSize="pageSize" :pageSizes="pageSizes" @current-change="currentChange"
            @size-change="sizeChange">
          </ReposList>
        </div>
        <div class="computer only ui four wide computer column">
          <div>
            <ActiveUsers></ActiveUsers>
          </div>
          <div class="active-org-c">
            <ActiveOrgs></ActiveOrgs>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SearchBar from '../components/SearchBar.vue';
import ReposFilters from '../components/ReposFilters.vue';
import ReposList from '../components/ReposList.vue';
import ActiveUsers from '../components/ActiveUsers.vue';
import ActiveOrgs from '../components/ActiveOrgs.vue';
import { getUrlSearchParams } from '~/utils';

const staticSquareTopics = JSON.stringify(window.staticSquareTopics || []);

export default {
  data() {
    return {
      reposListSortType: 'mostpopular',
      reposListQurey: '',
      reposListTopic: '',

      page: 1,
      pageSize: 15,
      pageSizes: [15, 30, 50],

      staticSquareTopics: staticSquareTopics,
    };
  },
  components: {
    SearchBar,
    ReposFilters,
    ReposList,
    ActiveUsers,
    ActiveOrgs,
  },
  methods: {
    filtersChange(condition) {
      this.page = 1;
      this.reposListSortType = condition.key;
      this.search();
    },
    searchBarChange(params) {
      this.page = 1;
      this.reposListQurey = params.q || '';
      this.reposListTopic = params.topic || '';
      this.search();
    },
    currentChange({ page, pageSize }) {
      this.page = page;
      this.search();
    },
    sizeChange({ page, pageSize }) {
      this.page = 1;
      this.pageSize = pageSize;
      this.search();
    },
    search() {
      window.location.href = `/explore/repos?q=${this.reposListQurey.trim()}&sort=${this.reposListSortType}&topic=${this.reposListTopic.trim()}&page=${this.page}&pageSize=${this.pageSize}`;
    }
  },
  beforeMount() {
    const urlParams = getUrlSearchParams();
    this.reposListQurey = urlParams.q || '';
    this.reposListTopic = urlParams.topic || '';
    this.reposListSortType = urlParams.sort || 'mostpopular';
    this.page = Number(urlParams.page) || 1;
    this.pageSize = this.pageSizes.indexOf(Number(urlParams.pageSize)) >= 0 ? Number(urlParams.pageSize) : 15;
  },
  mounted() {
    this.$nextTick(() => {
      this.$refs.reposFiltersRef.setDefaultFilter(this.reposListSortType);
      this.$refs.searchBarRef.setDefaultSearch({
        q: this.reposListQurey,
        topic: this.reposListTopic,
      });
      this.$refs.reposListRef.search();
    });
    window.addEventListener('pageshow', function (e) {
      if (e.persisted) {
        window.location.reload();
      }
    }, false);
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.recommend-repos-c {
  margin: 0 0 54px;
}

.active-org-c {
  margin-top: 32px;
}
</style>
