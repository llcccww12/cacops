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
          type="aimodel"
        />
        <Intro 
          :dataObj="datasetObj"
          v-if="tab=='intro'"
          type="aimodel"
        />
        <FileList
          :dataObj="datasetObj"
          @editSuccess="editSuccess"
          v-if="tab=='files'"
          type="aimodel"
        />
        <Graph
          :dataObj="datasetObj"
          v-if="tab=='graph'"
        />
        <UsageIntro
          :dataObj="datasetObj"
          v-if="tab=='usage'"
          type="model"
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
import Graph from '~/components/square/detail/graph/Graph.vue';
import UsageIntro from '~/components/square/detail/usage/UsageIntro.vue';
import Setting from '~/components/square/detail/setting/SettingModel.vue';
import { getModelInfoByName, getModelLicenseList } from '~/apis/modules/modelmanage';
import { MODEL_ENGINES } from '~/const';
import { getUrlSearchParams, setWebpackPublicPath, getListValueWithKey } from '~/utils';
import { formatDate } from 'element-ui/lib/utils/date-util';

export default {
  data() {
    return {
      emptyPage: false,
      loading: false,
      tab: 'intro',
      dataset_name: '',
      datasetObj: {},
      windowWidth: window.innerWidth
    };
  },
  components: { Header,  NotFound, Intro, FileList, Graph, UsageIntro, Setting },
  computed: {
    isMobile() {
      return this.windowWidth <= 768;
    },
    tabList() {
      return this.isMobile ? [
        {name: this.$t('modelManage.modelIntroduction'), icon: 'ri-database-2-line', key: 'intro'},
        {name: this.$t('modelManage.fileShort'), icon: 'ri-list-check', key: 'files'},
        {name: this.$t('modelManage.evolutionMap'), icon: 'ri-git-merge-line',  key: 'graph'},
        {name: this.$t('cloudbrainObj.useImage'), icon: 'ri-ticket-line', key: 'usage'},
      ] : [
        {name: this.$t('modelManage.modelIntroduction'), icon: 'ri-database-2-line',key: 'intro'},
        {name: this.$t('modelManage.modelFiles'), icon: 'ri-list-check', key: 'files'},
        {name: this.$t('modelManage.modelEvolutionMap'), icon: 'ri-git-merge-line',  key: 'graph'},
        {name: this.$t('modelManage.useModel'), icon: 'ri-ticket-line',  key: 'usage'}
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
         const response = await getModelInfoByName({aimodel_name: this.dataset_name})
         const res = response.data
         let ObjTemp = {}
         console.log("xxxx1",res)
         if(res.code === 0){
            ObjTemp = res.data
            console.log("xxxx2",ObjTemp)
            if(ObjTemp.aimodel_type === 2){
              console.log(ObjTemp.migration,ObjTemp.migration.status!==null)
              if(ObjTemp.migration && ObjTemp.migration.status!==null){
                if(ObjTemp.migration.status!==1){
                  location.href = `/models/detail/model_migrating/?model_id=${ObjTemp.id}`
                }
              }
            }
            ObjTemp.labels = ObjTemp.tags ? ObjTemp.tags.trim().split(/\s+/) : [],
            ObjTemp.tags = ''

            const response1 = await getModelLicenseList()
            console.log("xxxx2",response1)
            const res1 = response1.data
            const license = JSON.parse(res1) || [];
            const matchLicense = license.filter(item => item.id == ObjTemp.licenses);
            if (matchLicense.length) {
                ObjTemp.licenseInfo = matchLicense[0];
            }
            ObjTemp.licenses = ''
            ObjTemp.engineName = getListValueWithKey(MODEL_ENGINES, ObjTemp.engine)
            this.datasetObj = {
                ...ObjTemp,
                ...ObjTemp.permission
            }
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