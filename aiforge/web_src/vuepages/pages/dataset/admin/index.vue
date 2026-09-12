<template>
  <div>
    <h4 class="ui top attached header">
        数据集管理 (总计：{{total}})
    </h4>
    <div>
        <div style="width: 60%;margin-top: 15px;">
          <!-- <input v-model="q" autofocus @keydown.enter="search">
          <button class="ui blue button" @click="search">搜索</button> -->
          <el-input :placeholder="$t('userRole.pleaseEnterContent')" size="medium" v-model="q" @keydown.enter.stop.native.prevent="search">
            <el-select v-model="selectVal" slot="prepend" style="background-color: #fff;width: 120px;">
              <el-option :label="$t('datasetObj.name')" value="1"></el-option>
              <el-option :label="$t('datasetObj.dataset_owner')" value="2"></el-option>
              <el-option :label="$t('modelManage.creator')" value="3"></el-option>
            </el-select>
            <el-button type="primary" slot="append" @click="search">{{$t('repos.search')}}</el-button>
          </el-input>
        </div>
        <div style="display: flex;justify-content: space-between;align-items: center;margin:15px 0">
          <el-select size="medium" v-model="visibility" @change="selectVisibilityChange">
            <el-option v-for="item in visibilityList" :key="item.k" :label="item.v" :value="item.k" />
          </el-select>
          <div style="margin-right: 30px">
              <el-dropdown trigger="click" size="default">
                <span class="el-dropdown-link">
                    {{ $t('datasets.sort') }}<i class="el-icon-caret-bottom el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item :style="{color: sort == item.key ? '#409eff' : ''}" v-for="item in baseSorts" :key="item.key"
                        @click.native="changeSort(item)">
                        {{ item.label }}
                    </el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
          </div>
        </div>
    </div>
    <div class="ui attached segment">
        <div class="ui ten wide column">
            <el-checkbox @change="changeRecommend">{{$t('datasets.platform_recommendations')}}</el-checkbox>
        </div>
    </div>
    <div class="ui attached table segment">
        <el-table :data="list" style="width: 100%;font-size: 14px;" stripe v-loading="loading">
            <el-table-column  prop="name" :label="$t('cloudbrainObj.datasetFiles')" width="500">
                <template slot-scope="scope">
                    <div style="display: flex;align-items: center;">
                        <a style="margin-right:8px" class="nowrap" :href="`/datasets/detail/${scope.row.owner_name}/${scope.row.name}`" target="_blank">{{scope.row.alias}}</a>
                        <svg v-if="scope.row.recommend" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="16" height="16"><defs></defs><g><path fill="#ff6200" d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z"></path></g></svg>
                    </div>
              </template>
            </el-table-column>
            <el-table-column  prop="creator_name" :label="$t('modelManage.creator')" width="200" align="center" header-align="center">
            </el-table-column>
            <el-table-column  prop="owwner_name" :label="$t('datasetObj.dataset_owner')" width="200" align="center" header-align="center">
            </el-table-column>
            <el-table-column  prop="size_str" :label="$t('modelManage.fileSize')" width="200" align="center" header-align="center">
            </el-table-column>
            <el-table-column prop="is_private" :label="$t('status')"  width="200" align="center" header-align="center">
              <template slot-scope="scope">
                <div style="display:flex;justify-content: center;width: 100%;" v-if="scope.row.is_private">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="16" height="16"><defs></defs><g><path fill="#ff9929" d="M24.512 14.304v-3.84c0-4.672-3.808-8.48-8.48-8.48-0.002 0-0.004 0-0.006 0-4.68 0-8.474 3.794-8.474 8.474 0 0.002 0 0.004 0 0.006v-0 3.84h-1.824v15.712h20.576v-15.712h-1.792zM10.72 10.464c0-2.912 2.4-5.28 5.312-5.28s5.28 2.368 5.28 5.28v3.84h-10.592v-3.84zM17.472 25.696h-2.912l0.448-4.32c-0.541-0.33-0.896-0.917-0.896-1.587 0-0.004 0-0.009 0-0.013v0.001c0-1.056 0.864-1.92 1.92-1.92s1.888 0.864 1.888 1.92c0 0.672-0.352 1.28-0.896 1.6l0.448 4.32z"></path></g></svg>
                    <span style="color: #ff9959;margin-left:8px">{{ $t('modelManage.modelAccessPrivate') }}</span>
                </div>
                <div style="display:flex;justify-content: center;width: 100%;" v-else>
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="16" height="16"><defs></defs><g><path fill="#27b148" d="M29.696 14.464v15.040c0 0.288-0.224 0.48-0.48 0.48h-17.28c-0.288 0-0.512-0.192-0.512-0.48v-15.040c0-0.256 0.224-0.48 0.512-0.48h1.344c0.288 0 0.512-0.096 0.512-0.256v-0.224c-0.032-1.6 0-3.2-0.032-4.8-0.096-2.112-2.016-3.936-4.128-4.064-2.208-0.096-4.16 1.408-4.608 3.552-0.061 0.331-0.096 0.711-0.096 1.1 0 0.007 0 0.014 0 0.021v-0.001 4.416s-0.224 0.16-0.48 0.16h-1.664c-0.256 0-0.48-0.224-0.48-0.512 0-1.632-0.032-3.264 0.032-4.864 0.16-2.976 2.336-5.504 5.216-6.272 4.032-1.088 8.16 1.6 8.768 5.76 0.096 0.672 0.096 1.408 0.096 2.080v3.68s0.224 0.224 0.48 0.224h12.32c0.256 0 0.48 0.224 0.48 0.48z"></path></g></svg>
                    <span style="color: #27b128;margin-left:8px">{{ $t('modelManage.modelAccessPublic') }}</span>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="created_unix_str" :label="$t('modelManage.createTime')" align="center"
              header-align="center" width="200">
            </el-table-column>
            <el-table-column prop="operation" :label="$t('operation')" align="center" min-width="100"
              fixed="right" header-align="center">
              <template slot-scope="scope">
                <div class="op-wrap">
                    <a href="javascript:;" :style="{color: scope.row.recommend ? 'red': '' }"  @click="recommendData(scope.row)">{{scope.row.recommend ? '取消推荐' : '设为推荐'}}</a>
                </div>
              </template>
            </el-table-column>
        </el-table>
        <div style="margin:2rem 0" class="center" v-show="list.length">
            <el-pagination ref="paginationRef" background @current-change="currentChange"
                :current-page.sync="params.page" :page-size.sync="params.page_size"
                layout="total, prev, pager, next, jumper" :total="total">
            </el-pagination>
        </div>
    </div>
  </div>
</template>

<script>
import { getAdminDatasets, putRecommendDatasets, delRecommendDatasets } from "~/apis/modules/dataset";
import { formatDate } from 'element-ui/lib/utils/date-util';
import { transFileSize } from '~/utils';
export default {
  name: "adminDatasets",
  props: {
  },
  data() {
    return {
      loading: false,
      params: {
        q: '',
        tasks: '',
        tags: '',
        license: '',
        recommend: 'all',
        order_by: '',
        page: 1,
        page_size: 15,
        visibility: 'all',
        owner_name: undefined,
        creator_name: undefined
      },
      selectVal: '1',
      list: [],
      total: 0,
      baseSorts: [{
            key: '',
            label: this.$t('datasets.default'),
        }, {
            key: 'newest',
            label: this.$t('datasets.newest'),
        }, {
            key: 'recentupdate',
            label: this.$t('datasets.recentupdate'),
        }, {
            key: 'downloadcount',
            label: this.$t('datasets.downloadtimes'),
        }, {
            key: 'collections',
            label: this.$t('datasets.moststars'),
        }, {
            key: 'usecount',
            label: this.$t('datasets.mostusecount'),
        }, {
            key: 'alias_asc',
            label: this.$t('datasets.alphabetasc'),
        }, {
            key: 'alias',
            label: this.$t('datasets.alphabetdesc'),
        }],
      sort: '',
      visibility: 'all',
      visibilityList: [{
          k: 'all',
          v: this.$t('resourcesManagement.allStatus'),
        }, {
          k: 'public',
          v: this.$t('modelManage.modelAccessPublic'),
        }, {
          k: 'private',
          v: this.$t('modelManage.modelAccessPrivate'),
        }],
      q: ''
    }
  },
  watch: {
    q: {
      handler(val, oval) {
        if(val===''){
            this.params.q = ''
            this.getDatasetList()
        }
      },
      deep: true,
    }
  },
  methods: {
    search(){
        console.log("xxxxxxxxxxxxxxxxxx")
        this.params.q = ''
        if(this.selectVal==='1'){
          this.params.owner_name = undefined
          this.params.creator_name = undefined
          this.params.q = this.q
        } else if(this.selectVal==='2'){
          this.params.owner_name = this.q
          this.params.creator_name = undefined
        } else if(this.selectVal==='3'){
          this.params.owner_name = undefined
          this.params.creator_name = this.q
        }
      // this.params.q = this.q
      this.getDatasetList()
    },
    selectVisibilityChange(){
      this.params.page = 1
      this.params.visibility = this.visibility
      this.getDatasetList();
    },
    changeSort(item){
        console.log("xxxxxxxxxxxxx",item)
      this.sort = item.key
      this.params.order_by = item.key
      this.getDatasetList()
    },
    
    changeRecommend(val){
      this.params.page = 1
      if(val){
        this.params.recommend = 'only'
      }else{
        this.params.recommend = 'all'
      }
      this.getDatasetList();
    },
    currentChange(){
      this.getDatasetList();
    },
    async getDatasetList() {
      try {
            this.loading = true;
            console.log(this.params)
            const response = await getAdminDatasets(this.params)
            this.loading = false;
            
            if(response.data.code === 0){
              const res = response.data.data
              const list = (res.datasets || []).map(item => {
                return {
                  ...item,
                  created_unix_str: formatDate(new Date(item.created_unix * 1000), 'yyyy-MM-dd HH:mm:ss'),
                  size_str: transFileSize(item.size),
                  owwner_name: item?.Owner?.Name || '',
                  creator_name: item.creator_name || ''
                }
              })
              this.total = +res.total || 0
              this.list = list
            }else{
              this.$message.error(res.data.msg)
            }
            
        } catch (error) {
            this.loading = false;
            console.log(error)
            this.$message.error(error)
        }
    },
    async recommendData(item){
        
        try { 
            const params = {
                dataset_id: item.id,
                dataset_name: item.name
            }
            const response = item.recommend ? await delRecommendDatasets(params) : await putRecommendDatasets(params)
            console.log("response",response)
            const res = response.data
            if(res.code === 0){
                // this.$message.success(res.msg)
                this.getDatasetList()
            }else{
                this.$message.error(res.msg)

            }
        } catch (error) {
            this.$message.error(error)
        }
    }
  },
  mounted() {
    this.getDatasetList();
  },
}

</script>

<style lang="less" scoped>

</style>