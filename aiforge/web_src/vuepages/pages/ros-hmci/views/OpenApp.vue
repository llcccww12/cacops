<template>
  <div>
    <div class="group">
      <NavigationBar></NavigationBar>
      <div class="info">
        <div class="line1 flex-row">
          <span class="line1_1">不可错过的社区活动</span>
          <div class="line1_sqhd">
            <span class="line1_2">社区活动</span>
          </div>
        </div>
        <!-- <img src="/img/ros-hmci/平台动态.png" style="width: 50%"> -->
        <div class="info-div1">开源社区活动来袭</div>
        <div class="info-div">
          <span class="info2">欢迎您使用鹏城·盘古SDK 发布版——pcl_pangu v1.2！</span>
        </div>
      </div>
    </div>

    <div class="ui container">
      <SearchBar :static="true" :staticTopicsData="staticSquareTopics" ref="searchBarRef" type="square" :sort="``"
        :searchValue="reposListQurey" :topic="``" @change="searchBarChange"></SearchBar>
    </div>

    <div class="ui container">
      <div class="ui grid">
        <div class="computer only ui two wide computer column">
          <!-- <ReposFilters ref="reposFiltersRef" @change="filtersChange"></ReposFilters> -->
        </div>
        <div class="ui sixteen wide mobile twelve wide tablet ten wide computer column">
          <ReposList ref="reposListRef" :sort="reposListSortType" :q="reposListQurey" :topic="reposListTopic" :page="page"
            :pageSize="pageSize" :pageSizes="pageSizes" @current-change="currentChange" @size-change="sizeChange">
          </ReposList>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import NavigationBar from "../components/NavigationBar.vue";
//以下三个组件做过更改
import SearchBar from "../components/SearchBar.vue";
import ReposFilters from "../components/ReposFilters.vue";
import ReposList from "../components/ReposList.vue";
import { getUrlSearchParams } from "~/utils";

const staticSquareBanners = JSON.stringify(window.staticSquareBanners || []);
const staticSquarePreferredRepos = window.staticSquarePreferredRepos || [];
const staticSquareTopics = JSON.stringify(window.staticSquareTopics || []);
const staticSquareRecommendRepos = window.staticSquareRecommendRepos || [];

export default {
  name: "OpenApp",
  data() {
    return {
      reposListSortType: "mostpopular",
      reposListQurey: "",
      reposListTopic: "",

      page: 1,
      pageSize: 15,
      pageSizes: [15, 30, 50],

      staticSquareBanners: staticSquareBanners,
      staticSquarePreferredRepos: staticSquarePreferredRepos,
      staticSquareTopics: staticSquareTopics,
      staticSquareRecommendRepos: staticSquareRecommendRepos,
    };
  },
  components: {
    NavigationBar,
    SearchBar,
    ReposFilters,
    ReposList,
  },
  methods: {
    filtersChange(condition) {
      this.page = 1;
      this.reposListSortType = condition.key;
      this.search();
    },
    searchBarChange(params) {
      this.page = 1;
      this.reposListQurey = params.q || "";
      this.reposListTopic = params.topic || "";
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
      window.location.href = `/explore/repos/square?q=${this.reposListQurey.trim()}&sort=${this.reposListSortType
        }&topic=${this.reposListTopic.trim()}&page=${this.page}&pageSize=${this.pageSize}`;
    },
  },
  beforeMount() {
    const urlParams = getUrlSearchParams();
    this.reposListQurey = urlParams.q || "";
    this.reposListTopic = urlParams.topic || "";
    this.reposListSortType = urlParams.sort || "mostpopular";
    this.page = Number(urlParams.page) || 1;
    this.pageSize =
      this.pageSizes.indexOf(Number(urlParams.pageSize)) >= 0
        ? Number(urlParams.pageSize)
        : 15;
  },
  mounted() {
    this.$nextTick(() => {
      // this.$refs.reposFiltersRef.setDefaultFilter(this.reposListSortType);
      this.$refs.searchBarRef.setDefaultSearch({
        q: this.reposListQurey,
        topic: this.reposListTopic,
      });
      const urlParams = getUrlSearchParams();
      const page = Number(urlParams.page) || 1;
      const reposListSortType = urlParams.sort;
      if (page != 1 || reposListSortType) {
        window.location.href = "#search";
      }
      this.$refs.reposListRef.search();
    });
    window.addEventListener(
      "pageshow",
      function (e) {
        if (e.persisted) {
          window.location.reload();
        }
      },
      false
    );
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.group {
  position: relative;
  top: 0px;
  background: url(/img/ros-hmci/banner.png);
  background-position: center;
  background-size: cover;
  background-repeat: no-repeat;
  height: 600px;
  width: 100%;
}

.line1_1 {
  font-family: Alibaba PuHuiTi;
  color:#f3bf12;
  font-size: 20px;
}

.line1_sqhd {
  margin-left: 10px;
  display: flex;
  width: 69px;
  background-image: linear-gradient(133.55deg, #edbd4a 0%, #19bf77 100%);
  border-radius: 4px;
  justify-content: center;
}

.line1_2 {
  font-family: Alibaba PuHuiTi;
  color: #ffffff;
  font-size: 14px;
  padding: 4px 0;
}

.info {
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: relative;
  top: 165px;
  width: 1200px;
  margin: 0 auto;
}

// .info>*:not(:first-child) {
//   margin-top: 36px;
// }
.line1{
  align-items: center;
}
.info-div1{
  font-size: 40px;
  line-height: normal;
  margin: 20px 0 30px;
  background-image: linear-gradient(178.32deg,#06c08c 0%,#24a19b 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  font-weight: 700;
  font-family: sucaijishikufangti;
}
.info-div {
  display: inline-flex;
  flex-direction: row;
  align-items: center;
  height: 43px;
  background-image: linear-gradient(150.44deg, #19bf77 0%, rgba(26, 209, 195, 0) 100%);
  border-radius: 6px;
  padding: 0 60px 0 20px;
  width: fit-content;
}

.info1 {
  font-family: sucaijishikufangti;
  color: #ffffff;
  font-size: 45px;
  letter-spacing: 7px;
}

.info2 {
  font-family: Alibaba PuHuiTi;
  color: #ffffff;
  font-size: 16px;
  letter-spacing: 1px;
}

.recommend-repos-c {
  margin: 0 0 54px;
}

.active-org-c {
  margin-top: 32px;
}
</style>
