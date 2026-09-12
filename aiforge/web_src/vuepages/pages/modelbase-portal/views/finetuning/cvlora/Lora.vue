<template>
  <div class="list-wrap">
    <Header>
      <template slot="right">
        <el-button class="op-btn" type="primary" @click="goCreate">
          <div class="btn-content">
            <i class="ri-add-box-line"></i>
            <span>{{$t('modelSquare.newLoraFinetune')}}</span>
          </div>
        </el-button>
      </template>
    </Header>
    <div class="tips" v-html="$t('modelFinetune.maxTaskTips')"></div>
    <div class="header-img-card">
      <div v-for="model in modelList" :key="model.alias" class="card-box"> 
        <img :src="`/img/model/${model.alias}-w.png`" alt="">
      </div>
    </div>
    <div class="main-body">
      <div class="content">
        <div class="body">
          <div class="table-container">
            <el-table :data="tableData" style="min-width:100%" height="100%" v-loading="loading" stripe>
              <el-table-column :label="$t('cloudbrainObj.taskName')" align="left" header-align="left" min-width="250">
                <template slot-scope="scope">
                  <a class="dispaly-job-name" @click.prevent="goDetail(scope.row)">
                    {{ scope.row.task.display_job_name }}
                  </a>
                </template>
              </el-table-column>
              <el-table-column prop="Status" :label="$t('status')" align="left" header-align="left" min-width="142">
                <template slot-scope="scope">
                  <div class="status-wrap">
                    <i :class="scope.row.task.status"></i>
                    <span>{{ scope.row.task.status }}</span>
                    <i v-if="scope.row.task.detailed_status === 'dataMigrating' && scope.row.task.status === 'WAITING'"
                      :class="scope.row.task.detailed_status" :title="$t('cloudbrainObj.migratingData')"></i>
                    <i v-if="scope.row.task.detailed_status === 'centerPending' && scope.row.task.status === 'WAITING'"
                      :class="scope.row.task.detailed_status" :title="$t('cloudbrainObj.centerPending')"></i>
                    <i v-if="scope.row.task.detailed_status === 'ImagePulling' && scope.row.task.status === 'WAITING'"
                      :class="scope.row.task.detailed_status" :title="$t('cloudbrainObj.imagePulling')"></i>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="created_unix" :label="$t('cloudbrainObj.createTime')" align="center"
                header-align="center" width="190">
                <template slot-scope="scope">
                  <span>{{ dateFormat(scope.row.task.created_unix) }}</span>
                </template>
              </el-table-column>
              <el-table-column prop="formatted_duration" :label="$t('cloudbrainObj.runDuration')" align="center"
                header-align="center" width="120">
                <template slot-scope="scope">
                  <span>{{ scope.row.task.formatted_duration }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('cloudbrainObj.computeResource')" align="center" header-align="center" min-width="150">
                <template slot-scope="scope">
                  <span>{{ scope.row.task.computeSourceShow }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('resourcesManagement.aiCenter')" align="center" header-align="center"
                width="200">
                <template slot-scope="scope">
                  <span>{{ scope.row.task.ai_center }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('resourcesManagement.accCardType')" align="center" header-align="center"
                width="150">
                <template slot-scope="scope">
                  <span>{{ scope.row.task.accCardTypeShow }}</span>
                </template>
              </el-table-column>
              <el-table-column prop="remark" :label="$t('operation')" align="left" :width="($isMobile && isSubscriber) ? 300 : ''" min-width="200" header-align="center"
                fixed="right">
                <template slot-scope="scope">
                  <div class="op-wrap">
                    <a href="javascript:;"
                    v-if="isSubscriber"
                    @click="opDebug(scope.row)"
                    :class="scope.row.can_modify && scope.row.task.canDebug ? '' : 'disabled'">{{
                      $t('consoleHome')
                    }}</a>
                    <a href="javascript:;" :class="scope.row.can_experience && scope.row.task.canDebug ? '' : 'disabled'"
                    @click="opOnline(scope.row)">{{ $t('modelManage.onlineLoraTrain') }}</a>
                    <a href="javascript:;" :class="scope.row.can_delete && scope.row.task.canStop ? '' : 'disabled'"
                      @click="opStop(scope.row)">{{ $t('cloudbrainObj.stop') }}</a>
                    <a href="javascript:;" class="delete"
                      :class="scope.row.can_delete && scope.row.task.canDelete ? '' : 'disabled'"
                      @click="opDelete(scope.row)">{{ $t('cloudbrainObj.delete') }}</a>
                  </div>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </div>
        <div class="foot">
          <el-pagination ref="paginationRef" background @current-change="currentChange" @size-change="sizeChange"
            :current-page.sync="page" :page-sizes="pageSizes" :page-size.sync="pageSize"
            layout="total, sizes, prev, pager, next, jumper" :total="total">
          </el-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs';
import Header from '../../../components/Header.vue';
import { CloudBrainTools } from '~/pages/cloudbrain/tools';
import { timeSinceUnix, getListValueWithKey } from '~/utils';
import { CLUSTERS, JOB_TYPE, ACC_CARD_TYPE } from '~/const';
import { getMyAiTasks, stopAiTask, deleteAiTask, getAiEndpointUrl, getAiTaskDebugUrl } from '~/apis/modules/cloudbrain';
import { getPromoteData } from '~/apis/modules/common';
import { getLoraStage } from '~/apis/modules/llmchat';
const cloudBrainTools = new CloudBrainTools();
export default {
  name: 'Experience',
  data() {
    return {
      loading: false,
      tableData: [],

      page: 1,
      pageSize: 10,
      total: 0,
      pageSizes: [10, 20, 30, 50],
      isSubscriber: false,
      modelList: [],
    };
  },
  components: { Header },
  methods: {
    getTableData() {
      const params = {
        job_type: 'SDFINETUNE',
        job_status: '',
        ai_center: '',
        cluster: '',
        compute_source: '',
        q: '',
        page: this.page,
        pageSize: this.pageSize,
        exclude_status: undefined,
      };
      this.loading = true;
      const routerBase = this.$router.options.base;
      getMyAiTasks(params).then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 0) {
          const data = res.data;
          this.isSubscriber = data.is_admin;
          this.total = data.total;
          data.tasks.forEach(item => {
            const obj = Object.assign({}, item);
            delete obj.task;
            const task = item.task;
            Object.assign(task, obj);
            task.computeSourceShow = task.compute_source == 'GPU' ? 'CPU/GPU' : task.compute_source;
            task.clusterName = getListValueWithKey(CLUSTERS, task.cluster);
            task.accCardTypeShow = getListValueWithKey(ACC_CARD_TYPE, task.acc_card_type);
            task.repoOwnerName = item.owner_name;
            task.repoName = item.repo_name;
            task.createdFromNow = this.calcFromNow(task.created_unix);
            task.jobNameShow = task.job_name;
            task.jobTypeShow = getListValueWithKey(JOB_TYPE, task.job_type);
            cloudBrainTools.checkOperation(task);
          });
          this.tableData = data.tasks || [];
          cloudBrainTools.initRefreshData(this.tableData);
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    currentChange(page) {
      this.page = page;
      this.getTableData();
    },
    sizeChange(pageSize) {
      this.page = 1;
      this.pageSize = pageSize;
      this.getTableData();
    },
    calcFromNow(unix) {
      return timeSinceUnix(unix, Date.now() / 1000);
    },
    dateFormat(unix) {
      return dayjs(unix * 1000).format('YYYY-MM-DD HH:mm:ss');
    },
    goCreate() {
      this.$router.push(this.$route.path + '/create');
    },
    goDetail(row) {     
      this.$router.push(`/cv/sft/detail/${row.task.id}`);
    },
    opOnline(row) {
      if (this.operating) return;
      this.operating = true;
      getLoraStage({task_id: row.task.id}).then(res => {
        this.operating = false;
        if (res.status === 200 && res.data.url) {
          this.$router.push({name:'FinetuneSFTLora',query:{id:btoa(row.task.id),model:row.task.app_name}});
        } else {
          this.$message.warning(this.$t('modelSquare.ComfyUiWarn'));
        }
      }).catch(err => {
        this.operating = false;
        console.log(err);
        this.$message.warning(this.$t('modelSquare.ComfyUiWarn'));
      });
    },
    opStop(row) {
      if (this.operating) return;
      this.operating = true;
      stopAiTask({id: row.task.id}).then(res => {
        this.operating = false;
        res = res.data;
        if (res.code == 0) {
          const data = res.data;
          Object.assign(row.task, data);
          row.task.createdFromNow = this.calcFromNow(row.task.created_unix);
          cloudBrainTools.checkOperation(row.task);
        } else {
          this.$message({
            type: 'error',
            message: res.msg
          });
        }
      }).catch(err => {
        this.operating = false;
        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      });
    },
    opDelete(row) {
      if (this.operating) return;
      this.$confirm(this.$t('cloudbrainObj.deleteConfirmTips'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        this.operating = true;
        this.maskLoading = true;
        this.maskLoadingContent = this.$t('cloudbrainObj.deletingTips');
        deleteAiTask({
          repoOwnerName: row.task.repoOwnerName,
          repoName: row.task.repoName,
          id: row.task.id,
        }).then(res => {
          this.operating = false;
          this.maskLoading = false;
          if (this.total % this.pageSize == 1 && this.page > 1) {
            this.page -= 1;
          }
          this.getTableData();
        }).catch(err => {
          this.maskLoading = false;
          this.operating = false;
          this.$message({
            type: 'error',
            message: this.$t('operationFailed'),
          });
        });
      }).catch(() => {
        this.$message({
          type: 'info',
          message: this.$t('cancelOperate'),
        });
      });
    },
    opDebug(row) {
      if (this.operating) return;
      this.operating = true;
      getAiTaskDebugUrl({
        repoOwnerName: row.task.repoOwnerName,
        repoName: row.task.repoName,
        id: row.task.id,
      }).then(res => {
        this.operating = false;
        res = res.data;
        if (res.code == 0) {
          if (res.data && res.data.url) {
            if (!window.open(res.data.url)) {
              window.location.href = res.data.url;
            }
          }
        } else {
          this.$message({
            type: 'error',
            message: res.msg,
          });
        }
      }).catch(err => {
        this.operating = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('operationFailed'),
        });
      });
    },
  
    
  },
  beforeMount() {
    this.getTableData();
    getPromoteData('model/sdfinetune.json').then(res => {
      // this.loading = false;
      try {
        const data = JSON.parse(res.data);
        this.taskCreateInfo = { ...data }
        this.modelList = this.taskCreateInfo.model_name
      } catch (err) {
        this.$message.error(err);
      }
    })
  },
  mounted() { },
};
</script>

<style scoped lang="less">
@import '../../../components/createcommon.less';
.header-img-card{
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
  padding-left: 22px;
  margin-top: 12px;
  .card-box{
    img{width:100%;height: 100%;}
  }
  
}
@media (max-width: 768px){
  .header-img-card{
    .card-box{
      max-width: 160px;
      img{width:100%;height: 100%;}
    }
}
}
</style>
  