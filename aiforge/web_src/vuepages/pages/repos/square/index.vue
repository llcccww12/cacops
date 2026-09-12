<template>
  <div class="content">
    <div class="content-l" v-if="!isWorkspacePage">
      <Filters @changeCondition="conditionChange"></Filters>
    </div>
    <div class="content-r">
      <div v-if="isWorkspacePage" class="workspace-t">
        <span>{{ $t('repos.repos') }}</span>
      </div>
      <div class="filter-c">
        <div class="tab-c" :class="IsOrganization ? 'org-tab': ''">
          <div class="tab-item nowrap" :class="tab == item.key ? 'active' : ''" v-for="(item) in tabList" :key="item.key"
            @click="changeTab(item)">{{
              item.label }}</div>
        </div>
        <div class="right">
          <template v-if="isWorkspacePage">
            <a class="btn-loc" :href="issueLink"> 
              <svg width="16" height="16" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M24 44C29.5228 44 34.5228 41.7614 38.1421 38.1421C41.7614 34.5228 44 29.5228 44 24C44 18.4772 41.7614 13.4772 38.1421 9.85786C34.5228 6.23858 29.5228 4 24 4C18.4772 4 13.4772 6.23858 9.85786 9.85786C6.23858 13.4772 4 18.4772 4 24C4 29.5228 6.23858 34.5228 9.85786 38.1421C13.4772 41.7614 18.4772 44 24 44Z" fill="none" stroke="#101010" stroke-width="4" stroke-linejoin="round"/><path fill-rule="evenodd" clip-rule="evenodd" d="M24 37C25.3807 37 26.5 35.8807 26.5 34.5C26.5 33.1193 25.3807 32 24 32C22.6193 32 21.5 33.1193 21.5 34.5C21.5 35.8807 22.6193 37 24 37Z" fill="#101010"/><path d="M24 12V28" stroke="#101010" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
              <span style="margin-left:4px;">{{ $t('cloudbrainObj.issues') }}</span>
            </a>
            <a class="btn-loc" :href="pullLink"> 
              <svg width="16" height="16" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M37 44C39.2091 44 41 42.2091 41 40C41 37.7909 39.2091 36 37 36C34.7909 36 33 37.7909 33 40C33 42.2091 34.7909 44 37 44Z" fill="none" stroke="#101010" stroke-width="4" stroke-linejoin="round"/><path d="M11 12C13.2091 12 15 10.2091 15 8C15 5.79086 13.2091 4 11 4C8.79086 4 7 5.79086 7 8C7 10.2091 8.79086 12 11 12Z" fill="none" stroke="#101010" stroke-width="4" stroke-linejoin="round"/><path d="M11 44C13.2091 44 15 42.2091 15 40C15 37.7909 13.2091 36 11 36C8.79086 36 7 37.7909 7 40C7 42.2091 8.79086 44 11 44Z" fill="none" stroke="#101010" stroke-width="4" stroke-linejoin="round"/><path d="M11 12V36" stroke="#101010" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M24 10H33C35.2091 10 37 11.7909 37 14V36" stroke="#101010" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M30 16L24 10L30 4" stroke="#101010" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>
              <span style="margin-left:4px;">{{ $t('cloudbrainObj.pullRequests') }}</span>
            </a>
          </template>
          <div class="ui small icon input" style="height: 32px;">
            <input type="text" :placeholder="$t('repos.searchRepositories')" v-model="q" @keyup.enter="search">
            <i class="search icon" style="cursor: pointer;pointer-events: auto;" @click="search"></i>
          </div>
          <el-select v-model="sort" placeholder="请选择" @change="changeSort" style="width: 120px;">
            <el-option
              v-for="item in sortList"
              :key="item.key"
              :label="item.label"
              :value="item.key">
            </el-option>
          </el-select>
          <a class="btn-add" :href="repoCreateLink"> 
            <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 32 32" width="20" height="20"><defs></defs><g><path d="M5.333 4h21.333c0.736 0 1.333 0.597 1.333 1.333v0 21.333c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-21.333c0-0.736 0.597-1.333 1.333-1.333v0zM6.667 6.667v18.667h18.667v-18.667h-18.667zM14.667 14.667v-5.333h2.667v5.333h5.333v2.667h-5.333v5.333h-2.667v-5.333h-5.333v-2.667h5.333z"></path></g></svg>
            <span style="margin-left:6px;">{{ $t('cloudbrainObj.newRepo') }}</span>
          </a>
          <a class="btn-add" :href="repoMirateLink"> 
            <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 16 16" width="16" height="16"><defs></defs><g>
              <path fill-rule="evenodd" d="M15 0H9v7c0 .55.45 1 1 1h1v1h1V8h3c.55 0 1-.45 1-1V1c0-.55-.45-1-1-1zm-4 7h-1V6h1v1zm4 0h-3V6h3v1zm0-2h-4V1h4v4zM4 5H3V4h1v1zm0-2H3V2h1v1zM2 1h6V0H1C.45 0 0 .45 0 1v12c0 .55.45 1 1 1h2v2l1.5-1.5L6 16v-2h5c.55 0 1-.45 1-1v-3H2V1zm9 10v2H6v-1H3v1H1v-2h10zM3 8h1v1H3V8zm1-1H3V6h1v1z"></path>
            </g></svg>
            <span style="margin-left:6px;">{{ $t('cloudbrainObj.migrateRepo') }}</span>
          </a>
        </div>
      </div>
      <List :params="params" :isWorkspacePage="isWorkspacePage"></List>
    </div>
  </div>
</template>

<script>
import Filters from './components/Filters.vue';
import List from './components/List.vue';
const isWorkspacePage = window.MENU_CONFIG.activeTopMenu == 'workspace'
export default {
  data() {
    return {
      tab: 'all',
      mode: '',
      tabPublicList: [{ key: 'all',label: `⭕${this.$t('repos.publicRepos')}`, mode: '' }],
      tabOwnedList: [
        { key: 'all',label: `${this.$t('cloudbrainObj.all')}`, mode: ''},
        { key: 'source',label: `${this.$t('repos.source')}`, mode: 'source'},
        { key: 'fork',label: `${this.$t('repos.forks')}`, mode: 'fork'},
        { key: 'mirror',label: `${this.$t('repos.mirrors')}`, mode: 'mirror'},
        { key: 'collaborative',label: `${this.$t('repos.collaborative')}`, mode: 'collaborative'},
      ],
      tabOrgOwnedList: [
        { key: 'all',label: `${this.$t('cloudbrainObj.all')}`, mode: ''},
      ],
      sorPublictList: [
        { key: '', label: this.$t('datasets.default') },
        { key: 'mostpopular', label: this.$t('repos.mostPopular') }, 
        { key: 'mostactive', label: this.$t('repos.mostActive') }, 
        { key: 'recentupdate', label: this.$t('repos.recentlyUpdated') },
        { key: 'newest', label: this.$t('repos.newest') },
        { key: 'moststars', label: this.$t('repos.mostStars') },
        { key: 'mostforks', label: this.$t('repos.mostForks') }
      ],
      sorOwnedList:[
        { key: '', label: this.$t('datasets.default') }
      ],
      q: '',
      sort: '',
      topic: '',
      params: {
        tab: '',
        q: '',
        sort: '',
        topic: '',
        mode: '',
        exclusive: 0,
      },
      isWorkspacePage: isWorkspacePage,
      IsOrganization: false,
      orgName: '',
      orgId: 1,
    };
  },
  components: { Filters, List },
  computed: {
    tabList() {
      return this.isWorkspacePage ? this.IsOrganization ? this.tabOrgOwnedList :  this.tabOwnedList : this.tabPublicList
    },
    sortList() {
      return this.isWorkspacePage ? this.sorOwnedList : this.sorPublictList
    },
    issueLink(){
      return this.IsOrganization ? `/org/${this.orgName}/repositories/issues` : '/repositories/issues'
    },
    pullLink(){
      return this.IsOrganization ? `/org/${this.orgName}/repositories/pulls` : '/repositories/pulls'
    },
    repoCreateLink(){
      return this.IsOrganization ? `/repo/create?org=${this.orgId}` : '/repo/create' 
    },
    repoMirateLink(){
      return this.IsOrganization ? `/repo/migrate?org=${this.orgId}` : '/repo/migrate'
    }
  },
  methods: {
    changeTab(item) {
      this.tab = item.key;
      this.mode = item.mode
      this.search()
    },
    changeSort(item) {
      this.sort = item
      this.search()
    },
    search() {
      this.params = {
        q: this.q,
        tab: this.tab,
        sort: this.sort,
        topic: this.topic,
        mode: this.mode,
        exclusive: this.mode ? 1 : 0
      }
    },
    conditionChange(params = {}) {
      this.topic = params.topic;
      this.search()
    },
  },
  beforeMount() { 
    const orgInfo = document.querySelector('#org-info');
    if (orgInfo && orgInfo.getAttribute('data-orgname')) {
      this.IsOrganization = true
      this.orgName = orgInfo.getAttribute('data-orgname')
      this.orgId = orgInfo.getAttribute('data-orgid')
    }
  },
  mounted() {
    const params = new URLSearchParams(window.location.search);
    if (params.get('q')) {
      this.q = params.get('q').trim()
    }
    if(params.get('tab')){
      this.tab =  params.get('tab')
      this.mode = params.get('tab') === 'all' ? '' : params.get('tab')
      history.pushState(null, '', location.pathname);
    }
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
    padding: 30px 20px;
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
          color: rgba(16,16,16,0.5);
          &:hover {
            color: rgba(16,16,16,1);
            border-color: rgba(51, 38, 98, 1);
          }
          &.active {
            color: rgba(0,102,255,1);
            border-bottom: 2px solid rgba(0,102,255,1);
          }
        }
      }

      .right {
        display: flex;
        align-items: center;
        gap: 14px;
        .btn-loc{
          height: 32px;
          display: flex;
          align-items: center;
          padding: 0 10px;
          color: rgb(0, 0, 0);
          cursor: pointer;
        }
        .btn-add{
          height: 32px;
          background-color: rgba(50,145,248,1);
          font-size: 14px;
          display: flex;
          align-items: center;
          padding: 0 10px;
          color: #fff;
          border-radius: 4px;
          cursor: pointer;
          .fill:not([store]){
            fill: rgb(255, 255, 255);
          }
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
  .content{
    padding-left: 0 !important;
    .content-l {
      display: none;
    }
    .content-r{
      padding: 20px 16px !important;
      .filter-c{
        flex-direction: column;
        justify-content: center !important;
        gap: 12px;
        .input, .btn-add, .el-select{
          display: none !important;
        }
      }
      .org-tab{
        display: none !important;
      }
    }
  }

  /deep/ .list-item-container {
    .item-container {
      width: 100% !important;
      padding: 12px 0 !important;
    }
  }
}

@media only screen and (min-width: 768px) and (max-width: 1250px) {

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

@media only screen and (min-width: 1700px) {
  /deep/ .list-item-container {
    .item-container {
      width: 33.3% !important;

      &:nth-child(3n+1) {
        padding-left: 0;
      }

      &:nth-child(3n) {
        padding-right: 0;
      }
    }
  }
}
</style>
