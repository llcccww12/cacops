<template>
  <div class="content">
    <div class="title">{{ $t('taskTmplObj.taskTmpl') }}</div>
    <div class="condition-a">
      <div class="tab-c">
        <div class="tab-item" v-for="(item, index) in tabList" :class="conds.tab == item.key ? 'focus' : ''"
          :key="item.key" @click="changeTab(item)">
          {{ item.label }}
        </div>
      </div>
      <div class="condition-b">
        <div class="search-c">
          <el-select v-model="conds.taskType" @change="changeCondition">
            <el-option v-for="item in taskTypeList" :key="item.k" :label="item.v" :value="item.k">
            </el-option>
          </el-select>
          <el-select v-model="conds.resource" @change="changeCondition">
            <el-option v-for="item in resourceList" :key="item.k" :label="item.v" :value="item.k">
            </el-option>
          </el-select>
          <el-input class="search-keyword" :placeholder="$t('taskTmplObj.searchTaskTmpl')" v-model="conds.q"
            @keyup.enter.native="search">
            <i slot="suffix" class="el-input__icon el-icon-search" @click="search"></i>
          </el-input>
        </div>
        <div class="sort-c">
          <el-select v-model="conds.order_by" @change="changeSort" style="width: 120px;">
            <el-option v-for="item in sortList" :key="item.key" :label="item.label" :value="item.key">
            </el-option>
          </el-select>
        </div>
        <el-button class="create-btn" type="primary" @click="goCreate">
          <div class="btn-content">
            <i class="ri-add-box-line"></i>
            <span>{{ $t('taskTmplObj.createTaskTmpl') }}</span>
          </div>
        </el-button>
      </div>
    </div>
    <div class="table-c table-container" :class="isZh ? 'zh' : ''">
      <el-table :data="tableData" @selection-change="handleSelectionChange" style="min-width:100%" v-loading="loading"
        stripe>
        <!-- <el-table-column type="selection" width="45" :selectable="checkRowSelectable" fixed></el-table-column> -->
        <el-table-column :label="$t('taskTmplObj.tmplName')" align="left" header-align="center" min-width="240" :fixed="!isMiniScreen">
          <template slot-scope="scope">
            <a class="name" target="_blank" :href="`/ai_task_tmpl/detail/${scope.row.ID}`">
              <span class="nowrap" :title="scope.row.Name">{{ scope.row.Name }}</span>
              <svg v-if="scope.row.Recommend" xmlns="http://www.w3.org/2000/svg" fill="rgb(255, 98, 0)"
                viewBox="0 0 24 24" width="20" height="20">
                <defs></defs>
                <g>
                  <path
                    d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z">
                  </path>
                </g>
              </svg>
            </a>
          </template>
        </el-table-column>
        <el-table-column :label="$t('resourcesManagement.jobType')" align="center" header-align="center" width="140">
          <template slot-scope="scope">
            <div class="cb-job" :class="scope.row.JobType">
              <span>{{ scope.row.JobTypeStr }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column :label="$t('resourcesManagement.computeResource')" align="center" header-align="center"
          width="140">
          <template slot-scope="scope">
            <span>
              {{ scope.row.ComputeSourceStr }}
            </span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('taskTmplObj.modelDatasetAndRepo')" align="left" header-align="center"
          min-width="280">
          <template slot-scope="scope">
            <div class="row">
              <div class="row-l">{{ $t('repos.model') }}：</div>
              <div class="row-r">
                <div class="nowrap" v-for="model in (scope.row.ModelLists || [])" :key="model.ModelID"
                  :title="`${model.OwnerName ? `${model.OwnerName}/` : ''}${model.ModelAlias || model.ModelName}`">
                  {{ model.OwnerName ? `${model.OwnerName}/` : '' }}{{ model.ModelAlias || model.ModelName }}</div>
                <div v-if="!(scope.row.ModelLists || []).length">--</div>
              </div>
            </div>
            <div class="row">
              <div class="row-l">{{ $t('cloudbrainObj.dataset') }}：</div>
              <div class="row-r">
                <div class="nowrap" v-for="dataset in (scope.row.DatasetLists || [])" :key="dataset.DatasetID"
                  :title="`${dataset.OwnerName ? `${dataset.OwnerName}/` : ''}${dataset.DatasetAlias || dataset.DatasetName}`">
                  {{ dataset.OwnerName ? `${dataset.OwnerName}/` : '' }}{{ dataset.DatasetAlias || dataset.DatasetName
                  }}
                </div>
                <div v-if="!(scope.row.DatasetLists || []).length">--</div>
              </div>
            </div>
            <div class="row">
              <div class="row-l">{{ $t('repos.repos') }}：</div>
              <div class="row-r nowrap">
                <span v-if="scope.row.RepoOwnerName && scope.row.RepoName"
                  :title="`${scope.row.RepoOwnerName}/${scope.row.RepoName}`">{{
                    `${scope.row.RepoOwnerName}/${scope.row.RepoName}` }}</span>
                <span v-else>--</span>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column :label="$t('modelManage.creator')" align="center" header-align="center" width="90">
          <template slot-scope="scope">
            <a v-if="scope.row.Owner" :title="scope.row.Owner.FullName || scope.row.Owner.Name"
              :href="`/${scope.row.Owner.Name}`" class="avatar-c">
              <img class="avatar" :src="scope.row.Owner.RelAvatarLink">
            </a>
          </template>
        </el-table-column>
        <el-table-column :label="$t('cloudbrainObj.createTime')" align="center" header-align="center" width="140">
          <template slot-scope="scope">
            <span>{{ scope.row.CreatedUnixStr }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('modelManage.updateTime')" align="center" header-align="center" width="140">
          <template slot-scope="scope">
            <span>{{ scope.row.UpdatedUnixStr }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('taskTmplObj.accessRight')" align="center" header-align="center" width="145">
          <template slot-scope="scope">
            <div v-if="scope.row.IsPrivate" class="private"><i class="ri-lock-line"></i>{{
              $t('modelManage.modelAccessPrivate') }}</div>
            <div v-else>{{ $t('modelManage.modelAccessPublic') }}</div>
          </template>
        </el-table-column>
        <el-table-column :label="$t('taskTmplObj.collectedNum')" align="center" header-align="center" width="145">
          <template slot-scope="scope">
            <span>{{ scope.row.NumCollections }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="$t('taskTmplObj.runTimes')" align="center" header-align="center" width="140">
          <template slot-scope="scope">
            <span>{{ scope.row.UseCount }}</span>
          </template>
        </el-table-column>
        <el-table-column v-if="false" :label="$t('taskTmplObj.tagsAndDescr')" align="left" header-align="center"
          min-width="300">
          <template slot-scope="scope">
            <div class="descr" :title="scope.row.Description">{{ scope.row.Description }}</div>
            <div class="labels">
              <span v-for="tag in (scope.row.Tags || [])" :key="tag">{{ tag }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column :label="$t('operation')" align="center" header-align="center" min-width="220" fixed="right">
          <template slot-scope="scope">
            <div class="op-c">
              <div class="btn run-btn" @click="opRun(scope.row)">
                <svg xmlns="http://www.w3.org/2000/svg" fill="rgb(255, 255, 255)" viewBox="0 0 48 48" width="12"
                  height="12">
                  <defs></defs>
                  <g>
                    <path d="M24 44C12.9543 44 4 35.0457 4 24C4 12.9543 12.9543 4 24 4C35.0457 4 44 12.9543 44 24"
                      fill="none" stroke="rgb(255, 255, 255)" stroke-width="4" stroke-linecap="round"
                      stroke-linejoin="round"></path>
                    <path d="M20 24V17.0718L26 20.5359L32 24L26 27.4641L20 30.9282V24Z" fill="none"
                      stroke="rgb(255, 255, 255)" stroke-width="4" stroke-linejoin="round"></path>
                    <path d="M37.0508 32L37.0508 42" fill="none" stroke="rgb(255, 255, 255)" stroke-width="4"
                      stroke-linecap="round" stroke-linejoin="round"></path>
                    <path d="M42 36.9497L32 36.9497" fill="none" stroke="rgb(255, 255, 255)" stroke-width="4"
                      stroke-linecap="round" stroke-linejoin="round"></path>
                  </g>
                </svg>
                {{ $t('taskTmplObj.run') }}
              </div>
              <a class="btn edit-btn" v-if="scope.row.CanEdit && conds.tab == 'created'"
                :href="`/ai_task_tmpl/edit/${scope.row.ID}`">
                {{ $t('edit') }}
              </a>
              <div class="btn del-btn" v-if="scope.row.CanDelete && conds.tab == 'created'" @click="opDel(scope.row)">
                {{ $t('delete') }}
              </div>
              <div class="btn del-btn" v-if="conds.tab == 'collected'" @click="opCancelCollect(scope.row)">
                {{ $t('unStar') }}
              </div>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="list-foot">
      <!-- <el-button v-if="conds.tab == 'created'" class="left-btn" plain icon="el-icon-delete" @click="batchDelete">{{
        $t('cloudbrainObj.batchDelete')
        }}</el-button> -->
      <div class="pagination-c" v-show="tableData.length">
        <el-pagination background @current-change="currentChange" @size-change="sizeChange"
          :current-page.sync="pageParams.page" :page-sizes="pageParams.pageSizes" :page-size.sync="pageParams.pageSize"
          layout="total, sizes, prev, pager, next, jumper" :total="pageParams.total">
        </el-pagination>
      </div>
    </div>
  </div>
</template>

<script>
import { getAiTaskTmplList, deleteAiTaskTmpl, deleteCollectAiTaskTmpl } from '~/apis/modules/aitasktmpl';
import { formatDate } from 'element-ui/lib/utils/date-util';
import { getListValueWithKey } from '~/utils';
import { TmplTaskTypes, TmplComputerResouces } from '~/pages/aitasktmpl/tools';
import { lang } from '~/langs'

export default {
  data() {
    return {
      isZh: lang == 'zh-CN',
      tabList: [{
        key: 'created',
        label: this.$t('taskTmplObj.myCreated'),
      }, {
        key: 'collected',
        label: this.$t('taskTmplObj.myCollected'),
      }],
      taskTypeList: [
        { k: '', v: this.$t('resourcesManagement.allJobType') }, ...TmplTaskTypes
      ],
      resourceList: [
        { k: '', v: this.$t('resourcesManagement.allComputeResource') }, ...TmplComputerResouces
      ],
      sortList: [
        /*{ key: '', label: this.$t('datasets.default') },*/
        {
          key: 'recentupdate',
          label: this.$t('datasets.recentupdate'),
        }, {
          key: 'newest',
          label: this.$t('datasets.newest'),
        }, {
          key: 'collections',
          label: this.$t('datasets.moststars'),
        }, {
          key: 'usecount',
          label: this.$t('taskTmplObj.runTimes'),
        }, {
          key: 'name_asc',
          label: this.$t('datasets.alphabetasc'),
        }, {
          key: 'name',
          label: this.$t('datasets.alphabetdesc'),
        }],
      conds: {
        tab: 'created',
        taskType: '',
        resource: '',
        tags: '',
        model: '',
        dataset: '',
        repo: '',
        q: '',
        order_by: 'recentupdate',
      },
      multipleSelection: [],
      pageParams: {
        page: 1,
        pageSize: 10,
        pageSizes: [10, 20, 30, 50],
        total: 0
      },
      tableData: [],
      loading: false,
      operating: false,
      isMiniScreen: false,
    };
  },
  components: {},
  methods: {
    calcScreenInfo() {
      this.isMiniScreen = document.documentElement.clientWidth <= 800;
    },
    changeTab(item) {
      this.conds.tab = item.key;
      this.search();
    },
    changeCondition() {
      this.search()
    },
    handleSelectionChange(rows) {
      this.multipleSelection = []
      rows.forEach((item) => {
        this.multipleSelection.push(item.ID);
      })
    },
    checkRowSelectable(row) {
      return row.CanDelete;
    },
    changeSort(sort) {
      this.conds.order_by = sort;
      this.search()
    },
    getListData() {
      this.loading = true;
      getAiTaskTmplList({
        type: this.conds.tab,
        tags: this.conds.tags,
        job_type: this.conds.taskType,
        compute_source: this.conds.resource,
        model: this.conds.model,
        dataset: this.conds.dataset,
        repo: this.conds.repo,
        q: this.conds.q,
        order_by: this.conds.order_by || 'recentupdate',
        page: this.pageParams.page,
        page_size: this.pageParams.pageSize,
      }).then(res => {
        res = res.data;
        this.loading = false;
        if (res.code == 0) {
          res = res.data;
          this.pageParams.total = res.Total;
          this.tableData = (res.Templates || []).map(item => {
            return {
              ...item,
              NumCollections: Math.max(0, item.NumCollections),
              JobTypeStr: getListValueWithKey(TmplTaskTypes, item.JobType),
              ComputeSourceStr: getListValueWithKey(TmplComputerResouces, item.ComputeSource),
              CreatedUnixStr: formatDate(new Date(item.CreatedUnix * 1000), 'yyyy-MM-dd HH:mm:ss'),
              UpdatedUnixStr: formatDate(new Date(item.UpdatedUnix * 1000), 'yyyy-MM-dd HH:mm:ss'),
            }
          })
        } else {
          this.tableData = [];
          this.pageParams.total = 0;
        }
      }).catch(err => {
        console.log(err);
        this.loading = false;
        this.tableData = [];
        this.pageParams.total = 0;
      });
    },
    search() {
      this.pageParams.page = 1;
      this.conds.q = this.conds.q.trim();
      this.getListData();
    },
    currentChange(page) {
      this.pageParams.page = page;
      this.getListData();
    },
    sizeChange(pageSize) {
      this.pageParams.pageSize = pageSize;
      this.search();
    },
    goCreate() {
      window.location.href = '/ai_task_tmpl/create';
    },
    opRun(data) {
      window.location.href = `/cloudbrains/create?tmpl=${data.ID}`;
    },
    opEdit(data) {
      window.location.href = `/ai_task_tmpl/edit/${data.ID}`;
    },
    opDel(data) {
      if (this.operating) return;
      this.$confirm(this.$t('taskTmplObj.deleteTaskTmplConfirmTips'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning'
      }).then(() => {
        this.operating = true;
        deleteAiTaskTmpl({
          id: data.ID,
        }).then(res => {
          res = res.data;
          this.operating = false;
          if (res.code == 0) {
            if (this.pageParams.total % this.pageParams.pageSize == 1 && this.pageParams.page > 1) {
              this.page -= 1;
            }
            this.getListData();
          } else {
            this.$message({
              type: 'error',
              message: res.msg || this.$t('operationFailed'),
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
      }).catch((err) => {
        console.log(err);
        this.$message({
          type: 'info',
          message: this.$t('cancelOperate'),
        });
      });
    },
    opCancelCollect(data) {
      if (this.operating) return;
      this.operating = true;
      deleteCollectAiTaskTmpl({ id: data.ID }).then(res => {
        res = res.data;
        this.operating = false;
        if (res.code == '0') {
          if (this.pageParams.total % this.pageParams.pageSize == 1 && this.pageParams.page > 1) {
            this.page -= 1;
          }
          this.$message.success(this.$t('datasets.unstarSuccess'));
          this.getListData();
        } else {
          this.$message.error(res.msg || this.$t('operationFailed'));
        }
      }).catch(err => {
        console.log(err);
        this.operating = false;
      })
    },
    batchDelete() { },
  },
  beforeMount() {
    this.calcScreenInfo();
  },
  mounted() {
    this.search();
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';

.content {
  padding: 24px 24px 24px 32px;

  .title {
    color: rgb(16, 16, 16);
    font-size: 18px;
    text-align: left;
    font-family: 微软雅黑;
    font-weight: bold;
    margin-bottom: 16px;
  }

  .condition-a {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 4px;
    flex-wrap: wrap;

    .tab-c {
      display: flex;
      align-items: center;
      margin-bottom: 16px;

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
        color: rgba(16, 16, 16, 0.5);

        &:hover {
          color: rgba(16, 16, 16, 1);
          border-color: rgba(51, 38, 98, 1);
        }

        &.focus {
          color: rgba(0, 102, 255, 1);
          border-bottom: 2px solid rgba(0, 102, 255, 1);
        }
      }
    }

    .create-btn {
      margin-left: 14px;
      display: flex;
      align-items: center;
      height: 36px;
      font-size: 14px;
      background: rgba(22, 132, 252, 0.9);
      border-radius: 4px;

      &:active {
        background: rgb(22, 132, 252, 1);
      }


      &:focus,
      &:hover {
        background: rgba(22, 132, 252, 0.8);
      }

      .btn-content {
        display: flex;
        align-items: center;

        i {
          font-size: 14px;
          margin-right: 10px;
        }
      }
    }
  }


  .condition-b {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 16px;

    .search-c {
      display: flex;
      align-items: center;

      .el-select {
        margin-right: 15px;

        /deep/.el-input__inner {
          height: 36px;
        }
      }

      .el-input {
        width: 320px;
        margin-right: 15px;

        /deep/.el-input__inner {
          height: 36px;
        }
      }

      /deep/ .el-icon-search {
        cursor: pointer;
        color: rgb(16, 16, 16);
      }
    }

    .sort-c {
      color: rgba(0, 0, 0, 0.87);
      cursor: pointer;

      /deep/.el-input__inner {
        height: 36px;
      }
    }
  }

  .table-c {

    /deep/ .el-table__header {
      th {
        border: 1px solid rgb(208, 208, 208);
        border-right: 0;

        &:last-child {
          border-right: 1px solid rgb(208, 208, 208);
        }
      }
    }

    /deep/ .el-table__body {
      td {
        font-size: 12px;
        vertical-align: top;
        color: rgb(16, 16, 16);
      }
    }

    .name {
      color: rgb(0, 92, 255);
      font-size: 14px;
      font-weight: 700;
      display: flex;
      align-items: center;

      svg {
        margin-left: 4px;
        flex-shrink: 0;
      }
    }

    .task-type-c {
      display: flex;
      align-items: center;
      justify-content: center;

      .task-type {
        display: flex;
        align-items: center;
        padding: 0 6px;
        color: rgb(255, 255, 255);
        border-radius: 4px;
        background: rgb(27, 194, 134);
        height: 24px;
      }
    }

    .avatar-c {
      display: inline-block;
      height: 24px;

      .avatar {
        display: inline-block;
        width: 24px;
        height: 24px;
        border-radius: 100%;
      }
    }

    .row {
      display: flex;

      .row-l {
        color: rgba(16, 16, 16, 0.6);
        width: 80px;
        text-align: right;
      }

      .row-r {
        flex: 1;
        width: 0;
      }
    }

    .nowrap {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .private {
      display: flex;
      align-items: center;
      color: rgb(255, 98, 0);
      justify-content: center;

      i {
        margin-right: 2px;
      }
    }

    .descr {
      overflow: hidden;
      text-overflow: ellipsis;
      word-break: break-all;
      display: -webkit-box;
      -webkit-box-orient: vertical;
      -webkit-line-clamp: 2;
      max-height: 46px;
      white-space: break-spaces;
      margin-bottom: 2px;
    }

    .labels {
      display: flex;
      align-items: center;
      flex-wrap: wrap;

      span {
        height: 22px;
        border-radius: 4px;
        background-color: rgba(240, 240, 240, 1);
        color: rgba(16, 16, 16, 0.8);
        padding: 0 4px;
        margin-right: 6px;
        margin-bottom: 4px;
      }
    }

    .op-c {
      display: flex;
      align-items: center;
      justify-content: center;

      .run-btn {
        margin-left: 10px;
        display: flex;
        align-items: center;
        background: rgb(0, 102, 255);
        height: 24px;
        font-size: 12px;
        padding: 0 6px;
        border-color: rgba(157, 197, 226, 0.4);
        border-style: solid;
        border-width: 1px;
        border-radius: 4px;
        color: rgb(255, 255, 255);
        cursor: pointer;
        margin: 0 8px;

        svg {
          margin-right: 4px;
        }
      }

      .edit-btn,
      .del-btn {
        color: rgb(0, 102, 255);
        font-size: 14px;
        margin: 0 8px;
        cursor: pointer;
      }
    }

    &.zh {
      .row {
        .row-l {
          width: 50px;
        }

        .row-r {}
      }
    }
  }

  .list-foot {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 20px;

    .left-btn {
      border-color: #ff4d4f;
      color: #ff4d4f;
    }

    .pagination-c {
      margin-right: auto;
      margin-left: auto;
    }
  }
}
@media only screen and (max-width: 800px) {
  .content {
    margin: 10px 10px;
    padding: 0 !important;
    .condition-b {
      margin-left: auto;
      .search-c{
        display: none;
      }
    }
  }
  .el-pagination {
    /deep/ .el-pagination__total,
    /deep/ .el-pagination__sizes,
    /deep/ .el-pagination__jump {
      display: none;
    }
  }
    
}
</style>
