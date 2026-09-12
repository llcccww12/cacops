<template>
  <div class="org-area-c">
    <div class="org-area-l">
      <SpecailSelect ref="curatedData" v-if="pageType !== 'storage'" :type="pageType" :pageData="pageData"></SpecailSelect>
      <LabelSelect :labels="labels" :type="pageType" :pageData="pageData" @change="changeLabel"></LabelSelect>
      <RepoList v-if="pageType == 'repo'" :conds="conds" :pageData="pageData" @update-labels="updateLabels"></RepoList>
      <ModelList v-if="pageType == 'model'" :conds="conds" :pageData="pageData" @update-labels="updateLabels" @refreshFav="onRefreshFav">
      </ModelList>
      <DatasetList v-if="pageType == 'dataset'" :conds="conds" :pageData="pageData" @update-labels="updateLabels" @refreshFav="onRefreshFav">
      </DatasetList>
      <div v-if="pageType == 'storage'" style="margin: 20px;float:right">{{$t('storage.storageOrgTips1')}}
        <a download href="/OpenIOSSG/promote/raw/branch/master/%e5%90%af%e6%99%baAI%e5%8d%8f%e4%bd%9c%e5%b9%b3%e5%8f%b0%e7%bb%84%e7%bb%87%e8%b4%a6%e6%88%b7%e5%ad%98%e5%82%a8%e9%85%8d%e9%a2%9d%e7%94%b3%e8%af%b7%e8%a1%a8.docx">{{$t('storage.storageOrgTips2')}}</a>，
        {{$t('storage.storageOrgTips3')}}
      </div>
      <StorageList v-if="pageType == 'storage'" type="org" :orgName="pageData.Org.Name"></StorageList>
    </div>
    <div class="org-area-r" style="display: none;">
      <Members :pageData="pageData"></Members>
      <div class="line" v-if="pageData.IsOrganizationMember && pageData.Teams.length"></div>
      <Teams :pageData="pageData"></Teams>
    </div>
  </div>
</template>

<script>
import SpecailSelect from './components/SpecailSelect.vue';
import LabelSelect from './components/LabelSelect.vue';
import Members from './components/Members.vue';
import Teams from './components/Teams.vue';
import RepoList from './components/RepoList.vue';
import ModelList from './components/ModelList.vue';
import DatasetList from './components/DatasetList.vue';
import StorageList from '~/pages/storage/index.vue'
const pageData = window.PageData || {};
console.log(pageData)
export default {
  data() {
    return {
      pageType: '', // 'repo|model|dataset',
      pageData: pageData,
      conds: {
        label: '',
      },
      labels: [],
      list: [],
      page: 1,
      pageSize: 15,
      total: 0,
    };
  },
  components: { SpecailSelect, LabelSelect, RepoList, ModelList, DatasetList, Members, Teams, StorageList },
  methods: {
    changeLabel(label) {
      this.conds.label = label;
    },
    updateLabels(labels) {
      this.labels = [...labels];
    },
    onRefreshFav(id){
      this.$refs.curatedData.ishaveDataset(id)
    }
  },
  beforeMount() {
    const pageData = window.PageData || {};
    if (pageData.PageIsOrgHomeModel) {
      this.pageType = 'model';
    } else if (pageData.PageIsOrgHomeDataset) {
      this.pageType = 'dataset';
    } else if(pageData.PageIsOrgHomeStorage) {
      this.pageType = 'storage';
    } else{
      this.pageType = 'repo';
    }
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.org-area-c {
  display: flex;

  .org-area-l {
    width: 0;
    flex: 1;
  }

  .org-area-r {
    margin-left: 28px;
    width: 320px;

    .line {
      border-bottom: 1px solid rgb(225, 227, 230);
      margin: 16px 0;
    }
  }
}

@media only screen and (max-width: 767px) {
  .org-area-r {
    display: none;
  }
}
</style>
