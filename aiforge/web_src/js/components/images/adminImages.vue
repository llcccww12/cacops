<template>
  <div>
    <div class="ui container" style="width: 100% !important;padding-right: 0;">
      <div class="ui grid" style="margin: 0 !important">
        <div class="row" style="border: 1px solid #d4d4d5;margin-top:0px;padding-top: 0;">
          <div class="ui attached segment">
            <div class="ui form ignore-dirty">
              <div class="ui fluid action input">
                <input type="text" :placeholder="i18n['cloudeBrainMirror']['placeholder']" v-model="search"
                  @keyup.enter="searchName()">
                <button class="ui blue button" @click="searchName()">{{ i18n['cloudeBrainMirror']['search'] }}</button>
              </div>
            </div>
          </div>
          <div class="ui ten wide column" style="margin: 1rem 0;">
            <el-checkbox v-model="checked" style="padding: 0.5rem 1rem;">{{
              i18n['cloudeBrainMirror']['platform_recommendations'] }}</el-checkbox>
            <el-dropdown @command="handleCommandTaskType" trigger="click"
              style="border: 1px solid rgba(34,36,38,.15);border-radius: 4px;padding: 0.5rem 1rem;">
              <span class="el-dropdown-link">
                {{ dropdownTaskType }}<i class="el-icon-caret-bottom el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item :command="{ label: i18n['cloudeBrainMirror']['all_task_type'], type: '' }">{{
                  i18n['cloudeBrainMirror']['all_task_type'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['TaskTypeTitle']['Notebook'], type: 'Notebook' }">{{
                  i18n['TaskTypeTitle']['Notebook'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['TaskTypeTitle']['TrainJob'], type: 'TrainJob' }">{{
                  i18n['TaskTypeTitle']['TrainJob'] }}</el-dropdown-item>
                <!-- <el-dropdown-item :command="{ label: i18n['TaskTypeTitle']['Ecs'], type: 'Ecs' }">{{
                  i18n['TaskTypeTitle']['Ecs'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['TaskTypeTitle']['Inference'], type: 'Inference' }">{{
                  i18n['TaskTypeTitle']['Inference'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['TaskTypeTitle']['Service'], type: 'Service' }">{{
                  i18n['TaskTypeTitle']['Service'] }}</el-dropdown-item> -->
              </el-dropdown-menu>
            </el-dropdown>
            <el-dropdown @command="handleCommandComputeResource" trigger="click"
              style="border: 1px solid rgba(34,36,38,.15);border-radius: 4px;padding: 0.5rem 1rem;">
              <span class="el-dropdown-link">
                {{ dropdownComputeResource }}<i class="el-icon-caret-bottom el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item :command="{ label: i18n['cloudeBrainMirror']['all_compute_resource'], type: '' }">{{
                  i18n['cloudeBrainMirror']['all_compute_resource'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['computeResourceTitle']['CPU'], type: 'CPU' }">{{
                  i18n['computeResourceTitle']['CPU'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['computeResourceTitle']['GPU'], type: 'GPU' }">{{
                  i18n['computeResourceTitle']['GPU'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['computeResourceTitle']['NPU'], type: 'NPU' }">{{
                  i18n['computeResourceTitle']['NPU'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['computeResourceTitle']['GCU'], type: 'GCU' }">{{
                  i18n['computeResourceTitle']['GCU'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['computeResourceTitle']['MLU'], type: 'MLU' }">{{
                  i18n['computeResourceTitle']['MLU'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['computeResourceTitle']['DCU'], type: 'DCU' }">{{
                  i18n['computeResourceTitle']['DCU'] }}</el-dropdown-item>
                <el-dropdown-item
                  :command="{ label: i18n['computeResourceTitle']['ILUVATAR-GPGPU'], type: 'ILUVATAR-GPGPU' }">{{
                    i18n['computeResourceTitle']['ILUVATAR-GPGPU'] }}</el-dropdown-item>
                <el-dropdown-item
                  :command="{ label: i18n['computeResourceTitle']['METAX-GPGPU'], type: 'METAX-GPGPU' }">{{
                    i18n['computeResourceTitle']['METAX-GPGPU'] }}</el-dropdown-item>
                <el-dropdown-item
                  :command="{ label: 'BIREN-GPU', type: 'BIREN-GPU' }">{{
                    'BIREN-GPU' }}</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
            <el-dropdown @command="handleApplyState" trigger="click"
              style="border: 1px solid rgba(34,36,38,.15);border-radius: 4px;padding: 0.5rem 1rem;">
              <span class="el-dropdown-link">
                {{ dropdownApplyState }}<i class="el-icon-caret-bottom el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item :command="{ label: i18n['cloudeBrainMirror']['all_approval_status'], type: '' }">{{
                  i18n['cloudeBrainMirror']['all_approval_status'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['cloudeBrainMirror']['pending_approval'], type: 2 }">{{
                  i18n['cloudeBrainMirror']['pending_approval'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['cloudeBrainMirror']['approved'], type: 3 }">{{
                  i18n['cloudeBrainMirror']['approved'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['cloudeBrainMirror']['not_approved'], type: 4 }">{{
                  i18n['cloudeBrainMirror']['not_approved'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: i18n['cloudeBrainMirror']['none'], type: 1 }">{{
                  i18n['cloudeBrainMirror']['none'] }}</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
            <el-dropdown @command="handleCommandSort" trigger="click"
              style="border: 1px solid rgba(34,36,38,.15);border-radius: 4px;padding: 0.5rem 1rem;">
              <span class="el-dropdown-link">
                {{ sortCustom | transformSort(vm) }}<i class="el-icon-caret-bottom el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item :command="{ label: 'defaultsort', sort: '' }">{{
                  i18n['cloudeBrainMirror']['defaultsort'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: 'moststars', sort: 'moststars' }">{{
                  i18n['cloudeBrainMirror']['moststars'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: 'mostused', sort: 'mostused' }">{{
                  i18n['cloudeBrainMirror']['mostused'] }}</el-dropdown-item>
                <el-dropdown-item :command="{ label: 'newest', sort: 'newest' }">{{ i18n['cloudeBrainMirror']['newest']
                }}</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </div>
          <div class="ui six wide column right aligned" style="margin:1rem 0;display: flex;align-items:center;
            justify-content:flex-end;">
            <el-button size="medium" icon="el-icon-refresh" @click="syncComputerNetwork" v-loading="syncLoading">
              {{ i18n['cloudeBrainMirror']['syncAiNetwork'] }}</el-button>
            <a class="ui blue small button" style="margin-left:10px"  href="/admin/images/commit_image">{{
              i18n['cloudeBrainMirror']['create_cloud_brain_mirror'] }}</a>
          </div>
          <div class="ui sixteen wide column" style="padding: 0;overflow-x:auto;">
            <el-table :data="tableDataCustom" style="min-width:100%;" :header-cell-style="tableHeaderStyle">
              <el-table-column :label="i18n['cloudeBrainMirror']['mirror_tag']" width="250px" align="left" prop="tag">
                <template slot-scope="scope">
                  <div style="display: flex;align-items: center;">
                    <a class="text-over image_title" :title="scope.row.tag">{{ scope.row.tag }}</a>
                    <img v-if="scope.row.type == 5" src="/img/jian.svg" style="margin-left: 0.5rem;">
                  </div>
                </template>
              </el-table-column>
              <el-table-column :label="i18n['cloudeBrainMirror']['mirror_description']" width="300px" align="left"
                prop="description">
                <template slot-scope="scope">
                  <div class="image_desc" :title="scope.row.description">{{ scope.row.description }}
                  </div>
                  <div v-if="!!scope.row.topics">
                    <span v-for="(topic, index) in scope.row.topics" :key="index" class="ui repo-topic label topic"
                      style="cursor: default;">{{ topic }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column :label="i18n['cloudeBrainMirror']['task_type']" width="150px" align="left" prop="trainType">
                <template slot-scope="scope">
                  <span>{{ scope.row | transTrainType(vm) }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="i18n['model_compute_resource']" width="150px" align="left"
                prop="compute_resource">
                <template slot-scope="scope">
                  <span :title="scope.row.compute_resource">{{ i18n['computeResourceTitle'][scope.row.compute_resource]
                    || scope.row.compute_resource }}</span>
                </template>
              </el-table-column>
              <el-table-column :label="i18n['cloudeBrainMirror']['framework']" width="150px" align="left"
                prop="framework">
                <template slot-scope="scope">
                  {{ scope.row.framework }}<br />{{ scope.row.frameworkVersion }}
                </template>
              </el-table-column>
              <el-table-column :label="'Python'" width="100px" align="left" prop="python">
                <template slot-scope="scope">
                  {{ scope.row.pythonVersion }}
                </template>
              </el-table-column>
              <el-table-column :label="'Cuda'" width="100px" align="left" prop="cudaVersion">
                <template slot-scope="scope">
                  {{ scope.row.cudaVersion || '--' }}
                </template>
              </el-table-column>
              <el-table-column :label="'Cann'" width="100px" align="left" prop="cannVersion">
                <template slot-scope="scope">
                  {{ scope.row.cannVersion || '--' }}
                </template>
              </el-table-column>
              <el-table-column :label="'Dtk'" width="100px" align="left" prop="dtkVersion">
                <template slot-scope="scope">
                  {{ scope.row.dtkVersion || '--' }}
                </template>
              </el-table-column>
              <el-table-column :label="i18n['cloudeBrainMirror']['operationSystem']" width="180px" align="left"
                prop="python">
                <template slot-scope="scope">
                  {{ scope.row.operationSystem }}<br />{{ scope.row.operationSystemVersion }}
                </template>
              </el-table-column>
              <!-- <el-table-column :label="i18n['cloudeBrainMirror']['thirdPackages']" width="220px" align="left"
                prop="thirdpackages">
                <template slot-scope="scope">
                  <div class="image_desc" :title="scope.row.thirdPackages">{{ scope.row.thirdPackages }}</div>
                </template>
              </el-table-column>
               <el-table-column prop="cloudbrainType" :label="i18n['cloudeBrainMirror']['available_clusters']"
                width="120px" align="center">
                <template slot-scope="scope">
                  {{ scope.row.cloudbrainType | transformType(vm) }}
                </template>
              </el-table-column> -->
              <el-table-column prop="creator" :label="i18n['cloudeBrainMirror']['creator']" width="80px" align="center">
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
              <el-table-column prop="createdUnix" :label="i18n['cloudeBrainMirror']['creation_time']" align="center"
                width="160px">
                <template slot-scope="scope">
                  {{ scope.row.createdUnix | transformTimestamp }}
                </template>
              </el-table-column>
              <el-table-column prop="apply_status" :label="i18n['cloudeBrainMirror']['approval_status']" width="120px"
                fixed="right" align="center">
                <template slot-scope="scope">
                  <span v-if="scope.row.apply_status == 0" style="">{{ '--' }}</span>
                  <div v-if="scope.row.apply_status === 1"
                    style="display: flex;align-items:center;justify-content:center;">
                    <span> {{ i18n['cloudeBrainMirror']['none'] }}</span>
                    <el-tooltip v-if="scope.row.message" class="item" effect="dark" :content="scope.row.message"
                      placement="top">
                      <i class="INFO" style="margin-left: 0.3rem"></i>
                    </el-tooltip>
                  </div>
                  <div v-if="scope.row.apply_status === 2"
                    style="display: flex;align-items:center;justify-content:center;">
                    <span style="color: rgb(250, 140, 22);">{{ i18n['cloudeBrainMirror']['pending_approval'] }}</span>
                    <el-tooltip class="item" effect="dark" :content="i18n['cloudeBrainMirror']['pending_approval']"
                      placement="top">
                      <i class="CLOCK" style="margin-left: 0.3rem"></i>
                    </el-tooltip>
                  </div>
                  <div v-if="scope.row.apply_status === 3"
                    style="display: flex;align-items:center;justify-content:center;">
                    <span style="color: rgb(19, 194, 141);">{{ i18n['cloudeBrainMirror']['approved'] }}</span>
                    <el-tooltip class="item" effect="dark" :content="i18n['cloudeBrainMirror']['approved']"
                      placement="top">
                      <i class="SUCCEEDED" style="margin-left: 0.3rem"></i>
                    </el-tooltip>
                  </div>
                  <div v-if="scope.row.apply_status === 4"
                    style="display: flex;align-items:center;justify-content:center;">
                    <span style="color: red">{{ i18n['cloudeBrainMirror']['not_approved'] }}</span>
                    <el-tooltip class="item" effect="dark" :content="scope.row.message" placement="top">
                      <i class="FAILED" style="margin-left: 0.3rem"></i>
                    </el-tooltip>
                  </div>
                </template>
              </el-table-column>
              <el-table-column align="center" width="400px" :label="i18n['cloudeBrainMirror']['operation']"
                fixed="right">
                <template slot-scope="scope">
                  <div style="display: flex;justify-content: center;align-items: center;">
                    <div style="display: flex;align-items: center;padding: 0 1rem;" :title="i18n['citations']">
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
                        i18n['cloudeBrainMirror']['set_as_recommended'] }}</span>
                    <span style="padding: 0 1rem;color: rgb(250, 140, 22);cursor:pointer;"
                      v-if="scope.row.type != 5 && scope.row.apply_status == 2"
                      @click="unSetRecommend(scope.$index, scope.row.id, 'reject')">{{
                        i18n['cloudeBrainMirror']['not_recommend'] }}</span>
                    <span style="padding: 0 1rem;color: rgb(250, 140, 22);cursor:pointer;" v-if="scope.row.type == 5"
                      @click="unSetRecommend(scope.$index, scope.row.id, 'cancel')">{{
                        i18n['cloudeBrainMirror']['cancel_recommendation'] }}</span>
                    <div style="padding-left:1rem;cursor:pointer;">
                      <el-dropdown size="medium">
                        <span class="el-dropdown-link">
                          {{ i18n['cloudeBrainMirror']['more'] }}<i class="el-icon-arrow-down el-icon--right"></i>
                        </span>
                        <el-dropdown-menu slot="dropdown">
                          <el-dropdown-item @click.native="eidtImage(scope.row)">{{ i18n['cloudeBrainMirror']['edit']
                          }}
                          </el-dropdown-item>
                          <el-dropdown-item style="color: red;" @click.native="deleteImage(scope.row.id)">{{
                            i18n['cloudeBrainMirror']['delete'] }}</el-dropdown-item>
                        </el-dropdown-menu>
                      </el-dropdown>
                    </div>
                  </div>
                </template>
              </el-table-column>
            </el-table>
          </div>
          <div class="ui container" style="padding:2rem 0;text-align:center">
            <el-pagination background @size-change="handleSizeChangeCustom" @current-change="handleCurrentChangeCustom"
              :current-page="paramsCustom.page" :page-size="paramsCustom.pageSize" :page-sizes="[5, 15, 20]"
              layout="total, sizes, prev, pager, next, jumper" :total="totalNumCustom">
            </el-pagination>
          </div>
        </div>
      </div>
    </div>
    <el-dialog class="reason-dlg" width="600px" :title="reasonDialogTitle" :visible.sync="reasonDialogShow"
      :before-close="reasonDialogHandleClose" @open="reasonDialogHandleOpen">
      <div class="reason-content">
        <el-form label-width="80px" style="margin-top:36px;padding-right:60px;">
          <el-form-item :label="i18n['cloudeBrainMirror']['reason']" required>
            <el-input v-model="reasonDialogContent" maxlength="64"></el-input>
          </el-form-item>
        </el-form>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button class="btn" @click="reasonDialogHandleClose">{{ i18n['cancel'] }}</el-button>
        <el-button class="btn confirm-btn" type="primary" @click="reasonDialogHandleConfirm">{{ i18n['confirm']
        }}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>

const { _AppSubUrl, _StaticUrlPrefix, csrf } = window.config;
import qs from 'qs'
export default {
  components: {},
  data() {
    return {
      search: '',
      dropdownPrivate: '',
      dropdownType: '',
      dropdownTaskType: '',
      dropdownComputeResource: '',
      dropdownApplyState: '',
      checked: false,
      currentPageCustom: 1,
      pageSizeCustom: 15,
      totalNumCustom: 0,
      paramsCustom: { page: 1, pageSize: 15, q: '', recommend: false, cloudbrainType: -1, sort: '', apply: '', },
      tableDataCustom: [],
      starCustom: [],
      loadingCustom: false,
      vm: this,
      firstSearch: false,
      sortCustom: 'defaultsort',

      reasonDialogShow: false,
      reasonDialogType: '',
      reasonDialogTitle: '',
      reasonDialogContent: '',
      reasonDialogData: null,

      syncLoading: false,
    };
  },
  methods: {
    tableHeaderStyle({ row, column, rowIndex, columnIndex }) {
      if (rowIndex === 0) {
        return 'background:#f5f5f6;color:#606266'
      }
    },
    handleSizeChangeCustom(val) {
      this.paramsCustom.pageSize = val
      this.getImageListCustom()
    },
    handleCurrentChangeCustom(val) {
      this.paramsCustom.page = val
      this.getImageListCustom()
    },
    getImageListCustom() {
      this.loadingCustom = true
      if (this.firstSearch) {
        history.replaceState('', '', location.href.split('?')[0])
      }
      this.$axios.get('/admin/images/data', {
        params: this.paramsCustom
      }).then((res) => {
        this.totalNumCustom = res.data.count
        this.tableDataCustom = res.data.images
        this.tableDataCustom.forEach(element => {
          this.starCustom.push({ id: element.id, })
        });
        this.loadingCustom = false
      })
    },
    deleteImage(id) {
      let flag = 1
      let _this = this
      $('.ui.basic.modal.images')
        .modal({
          onDeny: function () {
            flag = false
          },
          onApprove: function () {
            _this.$axios.delete('/image/' + id).then((res) => {
              _this.getImageListCustom()
            })
            flag = true
          },
          onHidden: function () {
            if (flag == false) {
              $('.alert').html(_this.i18n['canceled_operation']).removeClass('alert-success').addClass('alert-danger').show().delay(1500).fadeOut();
            } else {
              $('.alert').html(_this.i18n['successfully_deleted']).removeClass('alert-danger').addClass('alert-success').show().delay(1500).fadeOut();
            }
          }
        })
        .modal('show')
    },
    eidtImage(row) {
      console.log("row", row)
      location.href = `/image/${row.id}/imageAdmin?trainType=${row.trainType.replace(/&/g, '_')}`
    },
    syncComputerNetwork() {
      this.syncLoading = true;
      this.$axios.post(`/admin/resources/image/sync?_csrf=${csrf}`).then(res => {
        this.syncLoading = false;
        res = res.data;
        if (res.Code === 0) {
          this.$message({
            type: 'success',
            message: this.i18n['submittedSuccessfully']
          });
          this.getImageListCustom()
        } else {
          this.$message({
            type: 'error',
            message: this.i18n['submittedFailed']
          });
        }
      }).catch(err => {
        console.log(err);
        this.syncLoading = false;
        this.$message({
          type: 'error',
          message: this.i18n['submittedFailed']
        });
      });
    },
    imageStar(index, id, isStar) {
      if (isStar) {
        this.$axios.put(`/image/${id}/action/unstar`).then((res) => {
          this.tableDataPublic[index].numStars = this.tableDataPublic[index].numStars - 1
          this.tableDataPublic[index].isStar = false
        })
      } else {
        this.$axios.put(`/image/${id}/action/star`).then((res) => {
          this.tableDataPublic[index].numStars = this.tableDataPublic[index].numStars + 1
          this.tableDataPublic[index].isStar = true
        })
      }
    },
    copyUrl(url) {
      const cInput = document.createElement('input')
      cInput.value = url
      document.body.appendChild(cInput)
      cInput.select()
      document.execCommand('Copy')
      cInput.remove()
    },
    searchName() {
      this.paramsCustom.q = this.search
      this.paramsCustom.page = 1
      this.getImageListCustom()
    },
    setRecommend(index, id) {
      this.$axios.put(`/admin/image/${id}/action/recommend`).then((res) => {
        // this.tableDataCustom[index].type = 5
        this.getImageListCustom()
      })
    },
    unSetRecommend(index, id, type) {
      // this.$axios.put(`/admin/image/${id}/action/unrecommend`).then((res) => {
      //     this.tableDataCustom[index].type = 0
      // })
      this.reasonDialogContent = '';
      this.reasonDialogType = type;
      this.reasonDialogData = this.tableDataCustom[index];
      if (type == 'reject') {
        this.reasonDialogTitle = this.i18n['cloudeBrainMirror']['not_recommend'];
        this.reasonDialogShow = true;
      } else if (type == 'cancel') {
        this.reasonDialogTitle = this.i18n['cloudeBrainMirror']['cancel_recommendation'];
        this.reasonDialogShow = true;
      }
    },
    handleCommand(command) {
      this.dropdownPrivate = command.label
      this.paramsCustom.private = command.private
      this.paramsCustom.page = 1
      this.getImageListCustom()
    },
    handleCommandType(command) {
      this.dropdownType = command.label
      this.paramsCustom.cloudbrainType = command.type
      this.paramsCustom.page = 1
      this.getImageListCustom()
    },
    handleCommandTaskType(command) {
      this.dropdownTaskType = command.label
      this.paramsCustom.trainType = command.type
      this.paramsCustom.page = 1
      this.getImageListCustom()
    },
    handleCommandComputeResource(command) {
      this.dropdownComputeResource = command.label
      this.paramsCustom.computeResource = command.type
      this.paramsCustom.page = 1
      this.getImageListCustom()
    },
    handleApplyState(command) {
      this.dropdownApplyState = command.label
      this.paramsCustom.apply = command.type
      this.paramsCustom.page = 1
      this.getImageListCustom()
    },
    handleCommandSort(command) {
      this.sortCustom = command.label
      this.paramsCustom.sort = command.sort
      this.paramsCustom.page = 1
      this.getImageListCustom()
    },
    reasonDialogHandleOpen() { },
    reasonDialogHandleClose() {
      this.reasonDialogShow = false;
    },
    reasonDialogHandleConfirm() {
      // console.log(this.reasonDialogType, this.reasonDialogContent, this.reasonDialogData);
      const reasonContent = this.reasonDialogContent.trim();
      if (reasonContent == '') {
        this.$message(this.i18n['cloudeBrainMirror']['pleaseEnterReason']);
        return;
      }
      this.$axios.put(`/admin/image/${this.reasonDialogData.id}/action/unrecommend`, {
        message: reasonContent,
      }).then((res) => {
        this.reasonDialogShow = false;
        this.getImageListCustom();
      });
    }
  },
  filters: {
    transformType(val, vm) {
      if (val == 0) {
        return `${vm.i18n['cloudeBrainMirror']['openi']}/GPU`;
      } else {
        return `${vm.i18n['cloudeBrainMirror']['c2net']}/GPU`;
      }
    },
    transformSort(val, vm) {
      if (val === 'moststars') {
        return vm.i18n['cloudeBrainMirror']['moststars'];
      } else if (val === 'mostused') {
        return vm.i18n['cloudeBrainMirror']['mostused'];
      } else if (val === 'defaultsort') {
        return vm.i18n['cloudeBrainMirror']['defaultsort'];
      } else {
        return vm.i18n['cloudeBrainMirror']['newest'];
      }
    },
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
      if(row){
        const typeList = row.trainType.split('&').map((type)=>{
          if(type){
            if(type === 'Notebook'){
              if(['GCU', 'GPU'].includes(row.compute_resource)){
                return vm.i18n['TaskTypeTitle']['Notebook']
              }else{
                return vm.i18n['TaskTypeTitle']['Notebook1']
              }
            }else{
              return vm.i18n['TaskTypeTitle'][type]
            }
          }
        })
        return typeList.join('、')
      }
      return ''
    },
  },
  watch: {
    checked(val) {
      if (this.firstSearch) {
        this.firstSearch = false
        return
      };
      this.paramsCustom.page = 1
      this.paramsCustom.recommend = val
      this.getImageListCustom()
    }
  },
  mounted() {
    this.getImageListCustom()
  },
  created() {
    this.i18n = window.i18n;
    this.dropdownPrivate = this.i18n['all'];
    this.dropdownType = this.i18n['cloudeBrainMirror']['all_cluster'];
    this.dropdownApplyState = this.i18n['cloudeBrainMirror']['all_approval_status'];
    this.dropdownTaskType = this.i18n['cloudeBrainMirror']['all_task_type'];
    this.dropdownComputeResource = this.i18n['cloudeBrainMirror']['all_compute_resource'];
    let params = new URLSearchParams(location.search)
    if (location.search) {
      this.firstSearch = true
      this.paramsCustom = qs.parse(location.search.split('?')[1]);
      if (params.has('private')) {
        this.dropdownPrivate = !params.get('private') ? this.i18n['cloudeBrainMirror']['private'] : this.i18n['cloudeBrainMirror']['public']
      }
      if (params.has('recommend')) {
        this.checked = params.get('recommend') === 'true' ? true : false
      }
      if (params.has('cloudbrainType')) {
        this.dropdownType = params.get('cloudbrainType') === 0 ? this.i18n['cloudeBrainMirror']['openi'] : this.i18n['cloudeBrainMirror']['c2net']
      }
      if (params.has('trainType')) {
        this.dropdownTaskType = this.i18n['TaskTypeTitle'][params.get('trainType')];
      }
      if (params.has('computeResource')) {
        this.dropdownComputeResource = this.i18n['computeResourceTitle'][params.get('computeResource')];
      }
      if (params.has('apply')) {
        const apply = params.get('apply');
        switch (apply) {
          case '1':
            this.dropdownApplyState = this.i18n['cloudeBrainMirror']['none']
            break;
          case '2':
            this.dropdownApplyState = this.i18n['cloudeBrainMirror']['pending_approval']
            break;
          case '3':
            this.dropdownApplyState = this.i18n['cloudeBrainMirror']['approved']
            break;
          case '4':
            this.dropdownApplyState = this.i18n['cloudeBrainMirror']['not_approved']
            break;
          case '':
            this.dropdownApplyState = this.i18n['cloudeBrainMirror']['all_approval_status']
          default:
            break;
        }
      }
      if (params.has('page')) {
        this.paramsCustom.page = Number(params.get('page')) || 1
      }
      if (params.has('pageSize')) {
        this.paramsCustom.pageSize = Number(params.get('pageSize')) || 15
      }
    }
  }
};
</script>

<style scoped lang="less">
.header-wrapper {
  background-color: #f5f5f6;
  padding-top: 15px;
}

.image_text {
  padding: 25px 0 55px 0;
}

#header {
  position: relative;
  top: -40px;
}

.el-dropdown-menu__item--divided {
  border-top: 1px solid blue;
}

.el-table thead {
  background-color: #f5f5f6;
}

/deep/ .el-tabs__item:hover {
  color: #000;
  font-weight: 500;

}

/deep/ .el-tabs__item.is-active {
  color: #000;
  font-weight: 500;
}

/deep/ .el-tabs__active-bar {
  background-color: #000
}

#success {
  background-color: #5bb973;
  color: white;
}

.text-over {
  overflow: hidden;
  text-overflow: ellipsis;
  vertical-align: middle;
  white-space: nowrap;
}

.image_title {
  display: inline-block;
  cursor: default;
  color: rgb(66, 98, 144);
}

.image_desc {
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  display: -webkit-box;
  text-overflow: ellipsis;
  overflow: hidden;
}

.heart-stroke {
  stroke: #FA8C16;
  stroke-width: 2;
  fill: #fff
}

.stars_active {
  fill: #FA8C16 !important;
  stroke: #FA8C16 !important
}

.header-new-drop {
  width: 100%;
}

.copy-adress {
  padding: 0 1rem;
  color: #0366d6;
  cursor: pointer;
}

.copy-adress-no {
  cursor: pointer;
  pointer-events: none;
  opacity: .45 !important;
  color: rgba(0, 0, 0, .6);
  padding: 0 1rem;
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