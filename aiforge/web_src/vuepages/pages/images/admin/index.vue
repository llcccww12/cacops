<template>
  <div>
    <div class="admin-page">
      <div class="ui attached segment">
        <div class="ui form ignore-dirty" style="display: flex;justify-content: space-between;align-items: center;">
          <div class="ui fluid action input" style="width: 80%">
            <input :value="conds.q" @input="onInput" @keyup.enter="search"
              :placeholder="$t('imagesObj.images_search')" />
            <button class="ui blue button" @click="search">{{ $t('repos.search') }}</button>
          </div>
          <div style="margin-right: 30px">
            <el-dropdown trigger="click" size="default">
              <span class="el-dropdown-link">
                {{ $t('datasets.sort') }}<i class="el-icon-caret-bottom el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item :style="{ color: conds.sort == item.k ? '#409eff' : '' }" v-for="item in sortList"
                  :key="item.k" @click.native="changeSort(item)">
                  {{ item.v }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </div>
        </div>
        <div class="content">
          <div class="list-head">
            <div class="filters-c">
              <el-checkbox v-model="conds.recommend" @change="changeConds">{{
                $t('datasets.platform_recommendations') }}</el-checkbox>
              <el-select v-model="conds.center" @change="changeConds" filterable>
                <el-option v-for="(item, index) in aiCenterList" :key="item.k + index" :value="item.k"
                  :label="item.v"></el-option>
              </el-select>
              <el-select v-model="conds.trainType" @change="changeConds" filterable>
                <el-option v-for="(item) in taskTypeList" :key="item.k" :value="item.k" :label="item.v"></el-option>
              </el-select>
              <el-select v-model="conds.computeResource" @change="changeConds" filterable>
                <el-option v-for="(item) in computeResourceList" :key="item.k" :value="item.k"
                  :label="item.v"></el-option>
              </el-select>
              <el-select v-model="conds.accCardType" @change="changeConds" filterable>
                <el-option v-for="item in cardTypeList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
              <el-select v-model="conds.apply" @change="changeConds" filterable>
                <el-option v-for="(item) in statusList" :key="item.k" :value="item.k" :label="item.v"></el-option>
              </el-select>
            </div>
            <div class="right">
              <el-button icon="el-icon-refresh" @click="syncComputerNetwork" v-loading="syncLoading">
                {{ $t('resourcesManagement.syncAiNetwork') }}</el-button>
              <!-- <a class="ui blue small button" style="margin-left:10px" href="/admin/images/commit_image">
              {{ $t('imagesObj.create_cloud_brain_mirror') }}</a> -->
              <el-button type="primary" icon="el-icon-plus" @click="goCreate">
                {{ $t('imagesObj.create_cloud_brain_mirror') }}</el-button>
            </div>
          </div>
          <div class="table-container">
            <el-table :data="tableData" style="min-width:100%" v-loading="loading" stripe>
              <el-table-column :label="$t('imagesObj.imageTag')" width="250" align="left" prop="tag">
                <template slot-scope="scope">
                  <div style="display: flex;align-items: center;">
                    <a class="nowrap image_title" :title="scope.row.tag">{{ scope.row.tag }}</a>
                    <img v-if="scope.row.type == 5" src="/img/jian.svg" style="margin-left: 0.5rem;">
                  </div>
                </template>
              </el-table-column>
              <el-table-column :label="$t('imagesObj.imageTag')" width="300" align="left" prop="description">
                <template slot-scope="scope">
                  <div class="image_desc" :title="scope.row.description">{{ scope.row.description }}
                  </div>
                  <div v-if="!!scope.row.topics">
                    <span v-for="(topic, index) in scope.row.topics" :key="index" class="ui repo-topic label topic"
                      style="cursor: default;">{{ topic }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column :label="$t('imagesObj.imageTaskType')" width="150" align="left" prop="trainType">
                <template slot-scope="scope">
                  <span>{{ scope.row | transTrainType(vm) }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('resourcesManagement.computeResource')" width="120" align="left"
                prop="compute_resource">
                <template slot-scope="scope">
                  <span :title="scope.row.compute_resource">{{ $t('computeResourceTitle.' + scope.row.compute_resource)
                    ||
                    scope.row.compute_resource }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('resourcesManagement.aiCenter')" width="150" align="left" prop="aiCenterName">
                <template slot-scope="scope">
                  <span :title="scope.row.aiCenterName">{{ scope.row.aiCenterName }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('resourcesManagement.accCardType')" width="120" align="left"
                prop="accCardType">
                <template slot-scope="scope">
                  <span :title="scope.row.accCardType">{{ scope.row.accCardTypeShow }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="$t('imagesObj.framework')" width="150" align="left" prop="framework">
                <template slot-scope="scope">
                  {{ scope.row.framework }}<br />{{ scope.row.frameworkVersion }}
                </template>
              </el-table-column>
              <el-table-column :label="'Python'" width="100" align="left" prop="python">
                <template slot-scope="scope">
                  {{ scope.row.pythonVersion }}
                </template>
              </el-table-column>
              <el-table-column :label="'Cuda'" width="100" align="left" prop="cudaVersion">
                <template slot-scope="scope">
                  {{ scope.row.cudaVersion || '--' }}
                </template>
              </el-table-column>
              <el-table-column :label="'Cann'" width="100" align="left" prop="cannVersion">
                <template slot-scope="scope">
                  {{ scope.row.cannVersion || '--' }}
                </template>
              </el-table-column>
              <el-table-column :label="'Dtk'" width="100" align="left" prop="dtkVersion">
                <template slot-scope="scope">
                  {{ scope.row.dtkVersion || '--' }}
                </template>
              </el-table-column>
              <el-table-column :label="$t('imagesObj.operationSystem')" width="180" align="left" prop="python">
                <template slot-scope="scope">
                  {{ scope.row.operationSystem }}<br />{{ scope.row.operationSystemVersion }}
                </template>
              </el-table-column>
              <el-table-column prop="creator" :label="$t('modelManage.creator')" width="80" align="center">
                <template slot-scope="scope">
                  <a v-if="scope.row.userName || scope.row.relAvatarLink" :href="'/' + scope.row.userName"
                    :title="scope.row.userName">
                    <img :src="scope.row.relAvatarLink" class="ui avatar image">
                  </a>
                  <a v-else>
                    <img class="ui avatar image" title="Ghost" src="/user/avatar/ghost/-1">
                  </a>
                </template>
              </el-table-column>
              <el-table-column prop="createdUnix" :label="$t('cloudbrainObj.createTime')" align="center" width="160">
                <template slot-scope="scope">
                  {{ scope.row.createdUnix | transformTimestamp }}
                </template>
              </el-table-column>
              <el-table-column prop="apply_status" :label="$t('imagesObj.approval_status')" width="120" fixed="right"
                align="center">
                <template slot-scope="scope">
                  <span v-if="scope.row.apply_status == 0" style="">{{ '--' }}</span>
                  <div v-if="scope.row.apply_status === 1"
                    style="display: flex;align-items:center;justify-content:center;">
                    <span> {{ $t('imagesObj.none') }}</span>
                    <el-tooltip v-if="scope.row.message" class="item" effect="dark" :content="scope.row.message"
                      placement="top">
                      <i class="INFO" style="margin-left: 0.3rem"></i>
                    </el-tooltip>
                  </div>
                  <div v-if="scope.row.apply_status === 2"
                    style="display: flex;align-items:center;justify-content:center;">
                    <span style="color: rgb(250, 140, 22);">{{ $t('imagesObj.pending_approval') }}</span>
                    <el-tooltip class="item" effect="dark" :content="$t('imagesObj.pending_approval')" placement="top">
                      <i class="CLOCK" style="margin-left: 0.3rem"></i>
                    </el-tooltip>
                  </div>
                  <div v-if="scope.row.apply_status === 3"
                    style="display: flex;align-items:center;justify-content:center;">
                    <span style="color: rgb(19, 194, 141);">{{ $t('imagesObj.approved') }}</span>
                    <el-tooltip class="item" effect="dark" :content="$t('imagesObj.approved')" placement="top">
                      <i class="SUCCEEDED" style="margin-left: 0.3rem"></i>
                    </el-tooltip>
                  </div>
                  <div v-if="scope.row.apply_status === 4"
                    style="display: flex;align-items:center;justify-content:center;">
                    <span style="color: red">{{ $t('imagesObj.not_approved') }}</span>
                    <el-tooltip class="item" effect="dark" :content="scope.row.message" placement="top">
                      <i class="FAILED" style="margin-left: 0.3rem"></i>
                    </el-tooltip>
                  </div>
                </template>
              </el-table-column>
              <el-table-column align="center" width="360" :label="$t('operation')" fixed="right">
                <template slot-scope="scope">
                  <div style="display: flex;justify-content: center;align-items: center;">
                    <div style="display: flex;align-items: center;padding: 0 1rem;" :title="$t('citations')">
                      <i class="ri-links-line" style="font-size: 16px;"></i>
                      <span style="line-height: 2;margin-left: 0.3rem;">{{ scope.row.useCount }}</span>
                    </div>
                    <div style="display: flex;align-items: center;cursor: default;;padding: 0 1rem;">
                      <svg width="1.4em" height="1.4em" viewBox="0 0 32 32" class="heart-stroke">
                        <path
                          d="M4.4 6.54c-1.761 1.643-2.6 3.793-2.36 6.056.24 2.263 1.507 4.521 3.663 6.534a29110.9 29110.9 0 0010.296 9.633l10.297-9.633c2.157-2.013 3.424-4.273 3.664-6.536.24-2.264-.599-4.412-2.36-6.056-1.73-1.613-3.84-2.29-6.097-1.955-1.689.25-3.454 1.078-5.105 2.394l-.4.319-.398-.319c-1.649-1.316-3.414-2.143-5.105-2.394a7.612 7.612 0 00-1.113-.081c-1.838 0-3.541.694-4.983 2.038z">
                        </path>
                      </svg>
                      <span style="line-height: 2;margin-left:0.3rem;">{{ scope.row.numStars }}</span>
                    </div>
                    <span style="padding: 0 1rem;color: rgb(19, 194, 141);cursor:pointer;" v-if="scope.row.type !== 5 && scope.row.status == '1'
                      && (scope.row.apply_status == 0 || scope.row.apply_status == 1 || scope.row.apply_status == 2)"
                      @click="setRecommend(scope.$index, scope.row.id)">{{
                        $t('taskTmplObj.setRecommend') }}</span>
                    <span style="padding: 0 1rem;color: rgb(250, 140, 22);cursor:pointer;"
                      v-if="scope.row.type != 5 && scope.row.apply_status == 2"
                      @click="unSetRecommend(scope.$index, scope.row.id, 'reject')">{{
                        $t('imagesObj.not_recommend') }}</span>
                    <span style="padding: 0 1rem;color: rgb(250, 140, 22);cursor:pointer;" v-if="scope.row.type == 5"
                      @click="unSetRecommend(scope.$index, scope.row.id, 'cancel')">{{
                        $t('taskTmplObj.cancelRecommend') }}</span>
                    <div style="padding-left:1rem;cursor:pointer;">
                      <el-dropdown size="medium">
                        <span class="el-dropdown-link">
                          {{ $t('cloudbrainObj.more') }}<i class="el-icon-arrow-down el-icon--right"></i>
                        </span>
                        <el-dropdown-menu slot="dropdown">
                          <el-dropdown-item @click.native="eidtImage(scope.row)">{{ $t('taskTmplObj.edit')
                            }}
                          </el-dropdown-item>
                          <el-dropdown-item style="color: red;" @click.native="deleteImage(scope.row.id)">{{
                            $t('delete') }}</el-dropdown-item>
                        </el-dropdown-menu>
                      </el-dropdown>
                    </div>
                  </div>
                </template>
              </el-table-column>
            </el-table>
          </div>
          <div class="list-foot">
            <div class="pagination-c" v-show="tableData.length">
              <el-pagination background @current-change="currentChange" @size-change="sizeChange"
                :current-page.sync="conds.page" :page-sizes="pageSizes" :page-size.sync="conds.pageSize"
                layout="total, sizes, prev, pager, next, jumper" :total="total">
              </el-pagination>
            </div>
          </div>
        </div>
      </div>
      <el-dialog class="reason-dlg" width="600px" :title="reasonDialogTitle" :visible.sync="reasonDialogShow"
        :before-close="reasonDialogHandleClose" @open="reasonDialogHandleOpen">
        <div class="reason-content">
          <el-form label-width="80px" style="margin-top:36px;padding-right:60px;">
            <el-form-item :label="$t('resourcesManagement.reason')" required>
              <el-input v-model="reasonDialogContent" maxlength="64"></el-input>
            </el-form-item>
          </el-form>
        </div>
        <span slot="footer" class="dialog-footer">
          <el-button class="btn" @click="reasonDialogHandleClose">{{ $t('cancel') }}</el-button>
          <el-button class="btn confirm-btn" type="primary" @click="reasonDialogHandleConfirm">{{ $t('confirm')
            }}</el-button>
        </span>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { COMPUTER_RESOURCES, ACC_CARD_TYPE } from '~/const';
import { syncImage, getImageListCustom, deleteImage, putImageRecommend, putImageUnRecommend } from '~/apis/modules/images';
import { getAiCenterList } from '~/apis/modules/resources';
import { getListValueWithKey } from '~/utils';
export default {
  data() {
    return {
      conds: {
        q: '',
        recommend: false,
        cloudbrainType: -1,
        trainType: '',
        sort: '',
        apply: '',
        computeResource: '',
        center: '',
        accCardType: '',
        page: 1,
        pageSize: 15,

      },
      aiCenterList: [{ k: '', v: this.$t('resourcesManagement.allAiCenter') }],
      taskTypeList: [
        { k: '', v: this.$t('resourcesManagement.allJobType') },
        { k: 'Notebook', v: this.$t('TaskTypeTitle.Notebook') },
        { k: 'TrainJob', v: this.$t('TaskTypeTitle.TrainJob') }
      ],
      computeResourceList: [{ k: '', v: this.$t('resourcesManagement.allComputeResource') }, ...COMPUTER_RESOURCES],
      cardTypeList: [{ k: '', v: this.$t('resourcesManagement.allAccCardType') }, ...ACC_CARD_TYPE],
      statusList: [
        { k: '', v: this.$t('imagesObj.all_approval_status') },
        { k: 2, v: this.$t('imagesObj.pending_approval') },
        { k: 3, v: this.$t('imagesObj.approved') },
        { k: 4, v: this.$t('imagesObj.not_approved') },
        { k: 1, v: this.$t('imagesObj.none') }
      ],
      sortList: [
        { k: '', v: this.$t('datasets.default') },
        { k: 'moststars', v: this.$t('datasets.mostCollections') },
        { k: 'mostused', v: this.$t('datasets.mostusecount') },
        { k: 'newest', v: this.$t('datasets.newest') },
      ],
      searchValue: '',
      syncLoading: false,
      loading: false,
      tableData: [],
      total: 0,
      pageSizes: [10, 15, 20],
      vm: this,

      reasonDialogShow: false,
      reasonDialogType: '',
      reasonDialogTitle: '',
      reasonDialogContent: '',
      reasonDialogData: null,
    }
  },
  components: {},
  watch: {
    'searchValue': function (val) {
      if (val === '') {
        setTimeout(() => {
          this.search();
        }, 100);
      }
    }
  },
  computed: {
  },
  filters: {
    transformTimestamp(timestamp) {
      const date = new Date(parseInt(timestamp) * 1000);
      const Y = date.getFullYear() + '-';
      const M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
      const D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) + '  ';
      const h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
      const m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
      const s = (date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds()); // 秒 
      const dateString = Y + M + D + h + m + s;
      return dateString;
    },
    transTrainType(row, vm) {
      if (row) {
        const typeList = row.trainType.split('&').map((type) => {
          if (type) {
            if (type === 'Notebook') {
              if (['GCU', 'GPU'].includes(row.compute_resource)) {
                return vm.$t('TaskTypeTitle.Notebook')
              } else {
                return vm.$t('TaskTypeTitle.Notebook1')
              }
            } else {
              return vm.$t(`TaskTypeTitle.${type}`)
            }
          }
        })
        return typeList.join('、')
      }
      return ''
    },
  },
  methods: {
    goCreate() {
      location.href = '/admin/images/commit_image'
    },
    onInput(event) {
      this.searchValue = event.target.value;
    },
    search() {
      this.conds.q = this.searchValue;
      this.conds.page = 1;
      this.getTableData();
    },
    changeConds() {
      this.search();
    },
    changeSort(item) {
      this.conds.sort = item.k
      this.search()
    },
    currentChange(page) {
      this.conds.page = page;
      this.getTableData();
    },
    sizeChange(pageSize) {
      this.conds.page = 1;
      this.conds.pageSize = pageSize;
      this.getTableData();
    },
    async getTableData() {
      try {
        this.loading = true
        const response = await getImageListCustom(this.conds);
        console.log("response", response)
        const res = response.data
        this.total = res.count
        this.tableData = res.images.map(item => {
          return {
            accCardTypeShow: getListValueWithKey(ACC_CARD_TYPE, item.accCardType),
            ...item,
          }
        })
      } catch (error) {
        this.$message.error(error)
      } finally {
        this.loading = false
      }
    },
    async syncComputerNetwork() {
      try {
        this.syncLoading = true;
        const response = await syncImage();
        console.log(response)
        const res = response.data;
        if (res.Code === 0) {
          this.$message.success(this.$t('submittedSuccessfully'))
          this.search()
        } else {
          this.$message.error(this.$t('submittedFailed'))
        }
      } catch (error) {
        this.$message.error(error || this.$t('submittedFailed'))
      } finally {
        this.syncLoading = false;
      }
    },
    deleteImage(id) {
      this.$confirm(this.$t('imagesObj.deleteTips'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm1'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        deleteImage({ id: id }).then(res => {
          this.isSetting = false;
          res = res.data;
          if (res.Code == '0') {
            this.$message({
              type: 'success',
              message: this.$t('imagesObj.deleteSuccessTips'),
            });
            this.search()
          } else {
            this.$message({
              type: 'error',
              message: res.Message,
            });
          }
        }).catch(err => {
          this.isSetting = false;
          console.log(err);
          this.$message({
            type: 'error',
            message: this.$t('operationFailed'),
          });
        });
      }).catch(() => { });
    },
    eidtImage(row) {
      location.href = `/image/${row.id}/imageAdmin?trainType=${row.trainType.replace(/&/g, '_')}`
    },
    setRecommend(index, id) {
      putImageRecommend(id).then((res) => {
        this.search()
      })
    },
    unSetRecommend(index, id, type) {

      this.reasonDialogContent = '';
      this.reasonDialogType = type;
      this.reasonDialogData = this.tableData[index];
      if (type == 'reject') {
        this.reasonDialogTitle = this.$t('imagesObj.not_recommend')
        this.reasonDialogShow = true;
      } else if (type == 'cancel') {
        this.reasonDialogTitle = this.$t('taskTmplObj.cancelRecommend')
        this.reasonDialogShow = true;
      }
    },
    reasonDialogHandleOpen() { },
    reasonDialogHandleClose() {
      this.reasonDialogShow = false;
    },
    reasonDialogHandleConfirm() {
      const reasonContent = this.reasonDialogContent.trim();
      if (reasonContent == '') {
        this.$message.error($t('enterReason'));
        return;
      }
      putImageUnRecommend(this.reasonDialogData.id, { message: reasonContent }).then((res) => {
        this.reasonDialogShow = false;
        this.search();
      });
    },
    getAiCenterList() {
      getAiCenterList().then(res => {
        res = res.data;
        if (res.Code === 0) {
          const list = res.Data;
          const data = list.map(item => {
            return {
              k: item.AiCenterCode,
              v: item.AiCenterName
            };
          });
          this.aiCenterList.splice(1, Infinity, ...data);
        }
      }).catch(err => {
        console.log(err);
      });
    },
  },
  beforeMount() {
    this.getAiCenterList();
  },
  mounted() {
    console.log(this)
    this.getTableData()
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.admin-page {
  border: 1px solid #d4d4d5;
  margin-top: 0px;
  padding-top: 0;
  box-shadow: 0 1px 2px 0 rgba(34, 36, 38, 0.15);

  .content {
    margin: 10px 10px;

    .list-head {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-top: 16px;
      margin-bottom: 12px;

      .filters-c {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        gap: 8px;
      }

      .right {
        display: flex;
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

    .table-container {
      /deep/ .el-table__header {
        th {
          background: #f0f0f0;
          color: rgba(0, 0, 0, .87);
          font-weight: 400;
          font-size: 14px;
          height: 48px;
          border: 1px solid #d0d0d0;
          border-right: 0;

          &:last-child {
            border-right: 1px solid rgb(208, 208, 208);
          }
        }
      }

      /deep/ .el-table__body {
        td {
          font-size: 12px;
        }
      }

      .image_title {
        display: inline-block;
        cursor: default;
        color: rgb(66, 98, 144);
      }

      .heart-stroke {
        stroke: #FA8C16;
        stroke-width: 2;
        fill: #fff
      }
    }
  }
}

.reason-dlg {
  /deep/ .el-dialog__header {
    text-align: left;
    height: 45px;
    background: rgb(240, 240, 240);
    border-radius: 5px 5px 0px 0px;
    border-bottom: 1px solid rgb(212, 212, 213);
    padding: 0 15px;
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: rgb(16, 16, 16);

    .el-dialog__title {
      font-weight: 500;
      font-size: 16px;
      color: rgb(16, 16, 16);
    }

    .el-dialog__headerbtn {
      top: 15px;
      right: 15px;
    }
  }

  /deep/ .el-dialog__body {
    padding: 15px 15px;
  }

  .el-input {
    /deep/ .el-input__inner:focus {
      border-color: #85b7d9;
      outline: 0;
    }
  }

  .btn {
    color: rgb(2, 0, 4);
    background-color: rgb(194, 199, 204);
    border-color: rgb(194, 199, 204);

    &.confirm-btn {
      color: #fff;
      background-color: rgb(56, 158, 13);
      border-color: rgb(56, 158, 13);
    }
  }
}
</style>
