<template>
  <div>
    <div v-if="emptyPage" style="padding-top:50px">
      <NotFound></NotFound>
    </div>
    <div v-else>
      <template v-if="JSON.stringify(datasetObj) !== '{}'">
        <Header 
          :tab="tab" 
          :tabList="tabList"
          :dataObj="datasetObj"
          @changeTab="changeTab"
          class="datset-header"
        />
        <Intro 
          :dataObj="datasetObj"
          v-if="tab=='intro'"
        />
        <FileList
          :dataObj="datasetObj"
          @editSuccess="editSuccess"
          v-if="tab=='files'"
        />
        <UsageIntro
          :dataObj="datasetObj"
          v-if="tab=='usage'"
        />
        <Setting
          :dataObj="datasetObj"
          @editSuccess="editSuccess"
          v-if="tab=='settings'"
        />
      </template>
    </div>
  </div>
</template>

<script>
import NotFound from '~/components/NotFound.vue';
import Header from '~/components/square/detail/Header.vue';
import Intro from '~/components/square/detail/intro/Intro.vue';
import FileList from '~/components/square/detail/files/FileList.vue';
import UsageIntro from '~/components/square/detail/usage/UsageIntro.vue';
import Setting from '~/components/square/detail/setting/Setting.vue';
import { getDatasetsDetail } from "~/apis/modules/dataset";
import { getUrlSearchParams, setWebpackPublicPath } from '~/utils';
import { lang } from '~/langs';

export default {
  data() {
    return {
      emptyPage: false,
      loading: false,
      // tabList:[
      //   {name: this.$t('datasetObj.dataset_intro'), icon: 'ri-database-2-line',key: 'intro'},
      //   {name: this.$t('modelManage.datasetfile'), icon: 'ri-list-check', key: 'files'},
      //   {name: this.$t('datasetObj.use_dataset'), icon: 'ri-ticket-line',  key: 'usage'},
      // ],
      tab: 'intro',
      dataset_name: '',
      datasetObj: {},
      windowWidth: window.innerWidth
    };
  },
  components: { Header,  NotFound, Intro, FileList, UsageIntro, Setting },
  computed: {
    isMobile() {
      return this.windowWidth <= 768;
    },
    tabList() {
      return this.isMobile ? [
        {name: this.$t('datasetObj.dataset_intro'), icon: 'ri-database-2-line', key: 'intro'},
        {name: this.$t('modelManage.fileShort'), icon: 'ri-list-check', key: 'files'},
        {name: this.$t('cloudbrainObj.useImage'), icon: 'ri-ticket-line', key: 'usage'},
      ] : [
        {name: this.$t('datasetObj.dataset_intro'), icon: 'ri-database-2-line', key: 'intro'},
        {name: this.$t('modelManage.datasetfile'), icon: 'ri-list-check', key: 'files'},
        {name: this.$t('datasetObj.use_dataset'), icon: 'ri-ticket-line', key: 'usage'},
      ];
    }
  },
  methods: {
    changeTab(item){
        this.tab = item.key
    },
    async getDatasetDetail(){
        try {
         this.loading = true
         const response = await getDatasetsDetail({dataset_name: this.dataset_name})
         const res = response.data
         if(res.code === 0){
            this.datasetObj = res.data
         } else {
          if (res.code === 9004) {
            this.emptyPage = true
          } else {
            this.$message.error(res.msg)
          }
         }
        } catch (error) {
            this.$message.error(error)
        }finally{
          this.loading = false
        }
    },
    editSuccess(){
      this.getDatasetDetail()
    },
    handleResize() {
      this.windowWidth = window.innerWidth;
      console.log("isMobile", this.isMobile);
    }
  },
  beforeMount() {
    const currentUrl = window.location.pathname;
    // 提取路径部分
    const pathParts = currentUrl.split('/');
    // 获取最后一个路径段
    const lastPathSegment = pathParts[pathParts.length - 1]; // "create"
    // 获取倒数第二个路径段（即父级目录）
    const parentSegment = pathParts[pathParts.length - 2]; // "notebook"
    this.dataset_name = `${parentSegment}/${lastPathSegment}`
    setWebpackPublicPath()

    window.addEventListener('resize', this.handleResize);
  },
  mounted() {
    const urlParams = getUrlSearchParams();
    if(urlParams.tab){
      this.tab = urlParams.tab
    } 
    this.getDatasetDetail()
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.handleResize);
  },
};
</script>

<style lang="less" scoped>
.datset-header{
  background: url('/img/dataset-square.png') center center no-repeat;
  background-size: cover;
}
</style>