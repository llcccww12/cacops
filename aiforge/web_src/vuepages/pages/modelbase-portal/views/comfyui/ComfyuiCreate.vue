<template>
  <div class="content-box">
    <Header></Header>
    <div class="main-body">
      <div class="main-content">
        <div class="err-msg-box-c" v-if="errorMsgBoxShow">
          <div class="err-msg-box">
            <p>{{ errorMsg }}</p>
          </div>
        </div>
        <div v-if="!alreadyMsgBoxShow" class="create-tips">
          <i class="ri-information-line"></i>
          <span style="margin-left:4px;">{{ $t('cloudbrainObj.runningLimit', {
            count: limitCount, taskType:
              $t('modelSquare.cvComfyui')
          }) }} </span>
        </div>
        <TaskLimitDialog :visible="limitDialogVisible" :title="limitDialogTitle" :desc="limitDialogDesc"
          :taskList="limitTaskList" @close="limitDialogVisible = false" @stopSuccess="onLimitStopSuccess"
          @deleteSuccess="onLimitDeleteSuccess" @taskStatusUpdate="onLimitTaskStatusUpdate"
          @allTasksTerminal="onLimitAllTasksTerminal">
        </TaskLimitDialog>
        <div class="area">
          <div class="area-content">
            <FormTopV2 ref="formTopRef" :resourceObj="pageCfg" :queueNum="queueNum" @change="changeComputeResouce">
            </FormTopV2>
            <SpecSelect ref="specRef" v-model="state.spec" :required="true" :configs="specConfigs" networkType="all"
              :loading="loading"></SpecSelect>
            <ImageSelectV1 ref="imagev1Ref" v-model="state.image" :configs="pageCfg" :spec="state.spec" :type="2"
              @changeImage="changeImage">
            </ImageSelectV1>
            <ModelSelect ref="modelRef" v-if="modelShowFlag" v-model="state.model" :multiple="true"
              :maxCount="modelCount"></ModelSelect>
            <TaskName ref="taskNameRef" v-model="state.taskName" :required="true" :userName="loginName"></TaskName>
            <ForbidPenetrationTipCheck v-model="agreeForbidPenetration"></ForbidPenetrationTipCheck>
            <div class="form-row">
              <div class="left-area">
                <div class="title"></div>
                <div class="content">
                  <el-button
                    :disabled="maskLoading || !agreeForbidPenetration || !specConfigs.specs['all'] || alreadyMsgBoxShow"
                    size="default" class="submit-btn" @click="submit">{{ $t('cloudbrainObj.createTask')
                    }}</el-button>
                  <el-button class="cancel-btn" size="default" @click="cancel">{{ $t('cancel') }}</el-button>
                </div>
              </div>
              <div class="right-area">
                <div class="btn-select">
                  <i class="question circle icon link" style="margin-top:-7px"></i>
                  <a :href="sampleUrl" target="_blank">{{ $t('modelManage.viewSamples') }}</a>
                </div>
              </div>
            </div>
            <div class="form-row" style="display:flex;justify-content:center;margin-top:45px;margin-bottom:-10px;">
              <AcknowledgementsTips></AcknowledgementsTips>
            </div>
          </div>
          <LoadingMask :loading="maskLoading" :content="maskLoadingContent"></LoadingMask>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from '~/pages/modelbase-portal/components/Header.vue';
import FormTopV2 from '~/components/cloudbrain/FormTopV2.vue';
import TaskName from '~/components/cloudbrain/TaskName.vue';
import SpecSelect from '~/components/cloudbrain/SpecSelect.vue';
import ModelSelect from '~/components/cloudbrain/ModelSelectV2.vue';
import LoadingMask from '~/pages/modelbase/components/cloudbrain/LoadingMask.vue';
import ImageSelectV1 from '~/components/cloudbrain/ImageSelectV1.vue';
import AcknowledgementsTips from '~/components/AcknowledgementsTips.vue';
import TaskLimitDialog from '~/components/cloudbrain/TaskLimitDialog.vue';
import ForbidPenetrationTipCheck from '~/components/cloudbrain/ForbidPenetrationTipCheck.vue';
import { getAiTaskPrepareInfo, createAiTask, getAiTask, getMyAiTasks } from '~/apis/modules/cloudbrain';
import { getPromoteData } from '~/apis/modules/common';
import { CLUSTERS, JOB_TYPE, ACC_CARD_TYPE } from '~/const';
import { i18n } from '~/langs';
import { getUrlSearchParams, getListValueWithKey } from '~/utils';
export default {
  name: 'ComfyuiCreate',
  data() {
    return {
      pageCfg: {
        cluster: 'C2Net',
        taskType: 'ComfyuiExperience',
        clusters: CLUSTERS.filter(item => item.k == 'C2Net').map(item => ({ ...item, key: item.k, label: item.v })),
        computerResouce: 'GPU',
        computerResouces: [{ key: 'GPU', label: i18n.t('computeResourceTitle.GPU') }],
        configs: {
          hideCluster: true,
          hideTips2: true,
          showMindTorchHelper: false
        }
      },
      state: {
        spec: '',
        taskName: '',
        model: [],
        image: { url: '' },
      },
      specConfigs: {
        specs: {},
        blance: 0,
        showPoint: false,
      },
      maskLoading: false,
      maskLoadingContent: '',
      errorMsgBoxShow: false,
      errorMsg: '',
      queueNum: 1,
      modelCount: 0,
      computeSource: '',
      loading: false,
      alreadyMsgBoxShow: false,
      notStopTaskCount: 0,
      limitCount: 0,
      limitDialogVisible: false,
      limitTaskList: [],
      taskCreateInfo: {},
      modelList: [],
      modelIndex: 0,
      modelShowFlag: false,
      sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/ComfyUI_Example',
      agreeForbidPenetration: true,
    };
  },
  components: { Header, FormTopV2, SpecSelect, LoadingMask, TaskName, AcknowledgementsTips, ForbidPenetrationTipCheck, ModelSelect, ImageSelectV1, TaskLimitDialog },
  computed: {
    limitDialogTitle() {
      return this.$t('cloudbrainObj.runningLimit', { count: this.limitCount, taskType: this.$t('modelSquare.cvComfyui') });
    },
    limitDialogDesc() {
      return this.$t('cloudbrainObj.sameTaskTips1', { count: this.notStopTaskCount });
    },
  },
  methods: {

    async prepare() {
      this.loading = true;
      try {
        const res = await getAiTaskPrepareInfo({
          jobType: 'ComfyuiExperience',
          clusterType: this.pageCfg.cluster,
          computeSource: this.pageCfg.computerResouce,
        });
        const data = res.data;
        this.loading = false;
        if (data.code === 0) {
          const resultData = data.data;
          this.queueNum = resultData.wait_count || 1;
          this.alreadyMsgBoxShow = !resultData.can_create_more;

          this.notStopTaskCount = resultData.not_stop_task_count || 0;
          this.limitCount = resultData.limit_count || 0;
          if (!resultData.can_create_more) {
            this.fetchLimitTaskList();
          } else {
            this.limitDialogVisible = false;
          }
          this.specConfigs.showPoint = resultData.pay_switch;
          this.specConfigs.blance = resultData.point_account ? resultData.point_account.balance : 0;
          this.specConfigs.specs = { 'all': resultData.specs.all };
          this.state.spec = this.specConfigs.specs['all'][0]?.id?.toString() || '';
          this.state.taskName = resultData.display_job_name;
          if (resultData.config && resultData.config.model_max_num) {
            this.modelCount = resultData.config.model_max_num;
          }
        }
      } catch (err) {
        this.loading = false;
        this.$message.error(err?.response?.data?.message || err);
      }
    },
    changeComputeResouce(data) {
      if (this.pageCfg.cluster != data.cluster || this.pageCfg.computerResouce != data.computerResouce) {
        this.pageCfg.cluster = data.cluster;
        this.pageCfg.computerResouce = data.computerResouce;
        this.prepare();
      }
    },
    fetchLimitTaskList() {
      const params = {
        job_type: 'ComfyuiExperience',
        job_status: 'nostop',
        ai_center: '', cluster: '', compute_source: '', q: '',
        page: 1, pageSize: 5
      };
      getMyAiTasks(params).then(res => {
        if (res.data.code === 0) {
          this.limitTaskList = res.data.data.tasks || [];
          this.limitTaskList.forEach((item) => {
            const task = item.task;
            task.computeSourceShow = task.compute_source == 'GPU' ? 'CPU/GPU' : task.compute_source;
            task.accCardTypeShow = getListValueWithKey(ACC_CARD_TYPE, task.acc_card_type);
            task.jobTypeShow = ['Comfy', 'UI']
          });
          this.limitDialogVisible = true;
        }
      });
    },
    onLimitStopSuccess() { this.prepare(); },
    onLimitDeleteSuccess() { this.fetchLimitTaskList(); this.prepare(); },
    onLimitTaskStatusUpdate(newTask) {
      const target = this.limitTaskList.find(t => t.task.id === newTask.id);
      if (target) target.task = newTask;
    },
    onLimitAllTasksTerminal() { this.prepare(); },
    changeImage() {

    },
    async getTaskInfo(taskId) {
      try {
        const res = await getAiTask({ id: taskId, });
        const data = res.data;
        if (data.code == 0) {
          const task = {};
          Object.assign(task, data.data.task);
          delete data.data.task;
          delete data.data.early_version_list;
          Object.assign(task, data.data);
          const specs = this.specConfigs.specs['all'] || [];
          let spec = '';
          if (task.spec.source_spec_id) {
            const find = specs.filter(item => {
              if (this.state.visualizeRequired) {
                return item.source_spec_id == task.spec.source_spec_id && item.enable_visualization;
              } else {
                return item.source_spec_id == task.spec.source_spec_id;
              }
            });
            if (find.length) {
              spec = find[0].id.toString();
              this.$nextTick(() => { this.state.spec = find[0].id.toString(); });
            }
          } else if (task.spec.id) {
            if (specs.filter(item => item.id == task.spec.id).length) {
              spec = task.spec.id.toString();
              this.$nextTick(() => { this.state.spec = task.spec.id.toString(); });
            }
          }
          if (task.pretrain_model_list) {
            let modelList = [];
            const models = task.pretrain_model_list;
            for (let i = 0, iLen = models.length; i < iLen; i++) {
              const _model = models[i];
              if (_model.is_delete) continue;
              const model = {
                ..._model,
              };
              modelList.push(model);
            }
            modelList = modelList.filter(item =>
              !this.modelList.some(removeItem => removeItem.id === item.id)
            );
            this.state.model = modelList;
          }
          if (task.image_url) {
            this.state.image.image_url = task.image_url;
            this.state.image.image_name = task.image_name
            this.state.image.image_id = task.image_id
          }
        }
      } catch (err) {
        console.log(err);
      }
    },
    submit() {
      if (this.maskLoading) return;
      let canSubmit = true;
      for (let key in this.state) {
        if (this.$refs[key + 'Ref']) {
          if (!this.$refs[key + 'Ref'].check()) {
            canSubmit = false
          }
        }
      }
      if (!canSubmit) return;
      const subObj = {
        repoOwnerName: this.taskCreateInfo.repo_owner_name,
        repoName: this.taskCreateInfo.repo_name,
        job_type: 'ComfyuiExperience',
        cluster: this.pageCfg.cluster,
        compute_source: this.pageCfg.computerResouce,
      }
      subObj['boot_file'] = this.taskCreateInfo.boot_file
      subObj['image_url'] = this.state.image.image_url //this.taskCreateInfo.image_url
      subObj['app_name'] = ''
      subObj['branch_name'] = this.taskCreateInfo.branch_name
      const combined = [...this.modelList, ...this.state.model];
      const uniqueArray = combined.filter((item, index) =>
        combined.findIndex(i => i.id === item.id) === index
      );
      subObj['pretrain_model_id_str'] = uniqueArray.map(item => item.id).join(';');
      subObj['display_job_name'] = this.state.taskName
      subObj['has_internet'] = this.specConfigs.specs.all[0].has_internet;
      subObj['spec_id'] = +this.state.spec;
      subObj['work_server_number'] = 1;
      subObj['endPoint'] = `${this.generateRandomString(5)}${Date.now()}`
      subObj['port'] = 8110
      this.maskLoadingContent = this.$t('cloudbrainObj.taskPrepareTips');
      this.maskLoading = true;
      this.errorMsg = '';
      this.errorMsgBoxShow = false;
      createAiTask(subObj).then(res => {
        const data = res.data;
        if (data.code == 0) {
          this.$router.push('/cv/comfyui');
        } else {
          this.maskLoading = false;
          this.errorMsg = data.msg;
          this.errorMsgBoxShow = true;
          document.querySelector('.main-content').scrollTo({ top: 0, behavior: 'smooth' });
        }
      }).catch(err => {
        this.maskLoading = false;
        this.$message.error(err)
      });
    },
    cancel() {
      if (this.$route.query.backurl) {
        window.location.href = this.$route.query.backurl;
        return;
      }
      if (this.$route.query.backpath) {
        this.$router.replace(this.$route.query.backpath);
        return;
      }
      this.$router.replace('/cv/comfyui');
    },
    generateRandomString(length) {
      // 生成随机字符串（包含小写字母和数字）
      let randomString = '';
      const characters = 'abcdefghijklmnopqrstuvwxyz';

      for (let i = 0; i < length; i++) {
        const randomIndex = Math.floor(Math.random() * characters.length);
        randomString += characters[randomIndex];
      }

      return randomString;
    }
  },
  beforeMount() {
    const isLogin = !!document.querySelector('meta[name="_uid"]');
    if (isLogin) {
      this.loginName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext')
    }

  },
  mounted() {
    const urlParams = getUrlSearchParams();
    const modify = urlParams.modify;
    const taskId = urlParams.id;
    getPromoteData('model/comfyui_experience.json').then(async (res) => {
      // this.loading = false;
      try {
        const data = JSON.parse(res.data);
        this.taskCreateInfo = { ...data }
        this.state.image.image_url = this.taskCreateInfo.image_url
        this.state.image.image_name = this.taskCreateInfo.image_url
        this.modelList = this.taskCreateInfo.model_name
        this.modelShowFlag = true
        await this.prepare();
        if (modify && taskId) {
          // this.isModifyTask = true;
          // this.modeifyTaskId = taskId;
          await this.getTaskInfo(taskId);
        }
      } catch (err) {
        this.loading = false;
        this.$message.error(err);
      }
    })
  },
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';
@import '../../components/createcommon.less';


.base-model-select {
  display: flex;
  margin-top: 14px;
  align-items: center;
  flex-direction: row;
  padding: 0 24px;
  column-gap: 20px;
  flex-wrap: wrap;

  .model-wrap {
    width: 100px;
    height: 100px;
    border-radius: 8px;
    position: relative;
    cursor: pointer;

    .model-item {
      width: 100px;
      height: 100px;
      object-fit: cover;
    }

    .model-text {
      background-image: linear-gradient(to right, #3a71fa, #9b4bff);
      position: absolute;
      bottom: 0;
      left: 0;
      width: 100%;
      height: 28px;
      font-size: 14px;
      line-height: 24px;
      display: flex;
      justify-content: center;
      color: #fff;
    }

    .border-out {
      position: absolute;
      top: 0;
      left: 0;
      background: #0000;
      width: 100%;
      height: 100%;
      border: 2px solid #624aff;
      border-radius: 8px;
    }

    .border-innder {
      position: absolute;
      top: 2px;
      left: 2px;
      background: #0000;
      width: calc(100% - 4px);
      height: calc(100% - 4px);
      border: 2px solid #fff;
      border-radius: 6px;
    }
  }
}

@media (max-width: 768px) {
  .base-model-select {

    .model-wrap {
      width: 80px;
      height: 80px;

      .model-item {
        width: 80px;
        height: 80px;
      }

      .model-text {
        font-size: 12px;
        height: 20px;
        line-height: 20px;
      }
    }
  }


}
</style>
