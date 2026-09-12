<template>
  <div>
    <el-dialog :title="$t('notebook.createNewNotebook')" :visible.sync="dialogVisible" width="50%"
      :close-on-click-modal="false" @closed="handleClose">
      <div class="wrapper" v-loading="loading" element-loading-spinner="el-icon-loading">
        <!--<div style="text-align: center;padding-bottom: 12px;">
          <span class="text-tip">{{ $t('notebook.tips') }}</span>
        </div>-->
        <div v-show="alertCb" class="ui message alert-info">
          <div style="display: flex;align-items: center;">
            <i class="ri-information-line" style="font-size: 35px;color:  rgba(242, 113, 28, 1);;"></i>
            <div style="text-align: left;margin-left: 1rem;">
              <div style="font-weight: 600;line-height: 2;">{{ $t('notebook.sameTaskTips1') }} <span
                  style="color:rgba(242, 113, 28, 1);">{{ $t('notebook.sameTaskTips2') }}</span>
                {{ $t('notebook.sameTaskTips3') }}</div>
              <div style="color:#939393">{{ $t('notebook.sameTaskTips4') }} “<a href="/cloudbrains" target="_blank">{{
                $t('notebook.sameTaskTips5') }} &gt; {{ $t('notebook.sameTaskTips6') }}</a>”
                {{ $t('notebook.sameTaskTips7') }}</div>
            </div>
          </div>
        </div>
        <div class="three-resource-type" :class="{ active: selectIndex == 0 }" @click="selectResource(0)">
          <div class="resource-child-node">
            <div class="resource-type-icon background-C">
              <span class="text">C</span>
            </div>
            <div class="resource-type-detail">
              <div class="detail-title"><span>{{ $t('notebook.cpuEnv') }}</span></div>
              <div class="detail-spec">
                <span>{{ cpuSpec }}</span>
              </div>
              <div class="detail-spec">
                <span>{{ $t('image') }}：{{ notebookInfo.imageCpuDescription }}</span>
              </div>
            </div>
            <div class="resource-select" :style="notebookInfo.cloudBrainPaySwitch ? { marginBottom: '10px' } : ''">
              <i v-if="selectIndex === 0" :class="{ 'slide-in-bottom': !slideActive && !initSelect }"
                class="ri-checkbox-circle-line green"></i>
              <i class="ri-checkbox-blank-circle-line gray" :class="{ 'fade-out': selectIndex === 0 }"></i>
            </div>
            <div class="resource-unit-price" v-if="notebookInfo.cloudBrainPaySwitch && !loading">{{
              notebookInfo.specCpu.unit_price.toFixed(2) }} {{ $t('resourcesManagement.point_hr') }}</div>
          </div>
        </div>

        <div class="three-resource-type" :class="{ active: selectIndex == 2 }" @click="selectResource(2)">
          <div class="resource-child-node">
            <div class="resource-type-icon background-N">
              <span class="text">N</span>
            </div>
            <div class="resource-type-detail">
              <div class="detail-title"><span>{{ $t('notebook.npuEnv') }}</span></div>
              <div class="detail-spec">
                <span>{{ npuSpec }}</span>
              </div>
              <div class="detail-spec">
                <!-- <span>{{$t('image')}}：{{notebookInfo.imageNpuDescription}}</span> -->
                <div class="select-image">
                  <label>{{ $t('image') }}：</label>
                  <select @click.stop="" class="select-image-label" v-model="selectedImage">
                    <option v-for="item in npuImageList" :key="item.image_id" :value="item.image_id">
                      {{ item.image_name }}
                    </option>
                  </select>
                </div>
              </div>
            </div>
            <div class="resource-select" :style="notebookInfo.cloudBrainPaySwitch ? { marginBottom: '10px' } : ''">
              <i v-if="selectIndex === 2" :class="[slideActive && !initSelect ? 'slide-in-top' : 'slide-in-bottom']"
                class="ri-checkbox-circle-line green"></i>
              <i class="ri-checkbox-blank-circle-line gray" :class="{ 'fade-out': selectIndex === 2 }"></i>
            </div>
            <div class="resource-unit-price" v-if="notebookInfo.cloudBrainPaySwitch && !loading">{{
              notebookInfo.specNpu.unit_price.toFixed(2) }} {{ $t('resourcesManagement.point_hr') }}</div>
          </div>
        </div>
        <div class="three-resource-type" :class="{ active: selectIndex == 1 }" @click="selectResource(1)">
          <div class="resource-child-node">
            <div class="resource-type-icon background-G">
              <span class="text">G</span>
            </div>
            <div class="resource-type-detail">
              <div class="detail-title"><span>{{ $t('notebook.gpuEnv') }}</span></div>
              <div class="detail-spec">
                <span>{{ gpuSpec }}</span>
              </div>
              <div class="detail-spec">
                <span>{{ $t('image') }}：{{ notebookInfo.imageGpuDescription }}</span>
              </div>
            </div>
            <div class="resource-select" :style="notebookInfo.cloudBrainPaySwitch ? { marginBottom: '10px' } : ''">
              <i v-if="selectIndex === 1 && !initSelect" :class="{ 'slide-in-top': slideActive && !initSelect }"
                class="ri-checkbox-circle-line green"></i>
              <i v-if="selectIndex === 1 && initSelect" class="ri-checkbox-circle-line green"></i>
              <i class="ri-checkbox-blank-circle-line gray" :class="{ 'fade-out': selectIndex === 1 }"></i>
            </div>
            <div class="resource-unit-price" v-if="notebookInfo.cloudBrainPaySwitch && !loading">{{
              notebookInfo.specGpu.unit_price.toFixed(2) }} {{ $t('resourcesManagement.point_hr') }}</div>
          </div>
        </div>
        <div class="resource-footer">
          <div class="resource-operate" v-if="selectIndex == 0">
            <div v-if="btnStatus[0] === 0">
              <button class="ui green small button" @click="createTask(0)">{{ $t('notebook.newTask') }}</button>
              <span v-if="notebookInfo.waitCountGpu == 0" class="text">{{ $t('notebook.noQuene') }}</span>
              <span v-else class="text">{{ $t('notebook.queneTips1') }} <span style="color: red;">{{
                notebookInfo.waitCountGpu + 1 }}</span> {{ $t('notebook.queneTips2') }}</span>
            </div>
            <div v-else-if="btnStatus[0] === 1">
              <button class="ui disabled small button" style="background-color: #888;"><i class="loading spinner icon"
                  style="color: #fff;"></i>{{ $t('notebook.newTask') }}</button>
              <span class="text">{{ $t('notebook.watiResource') }}</span>
            </div>

            <div v-else-if="btnStatus[0] === 2">
              <button class="ui small button" style="background-color: #1684fc;" @click="goDebug(0)">
                <a style="color:#fff" href="javascript:;">{{ $t('notebook.debug') }}</a>
              </button>
              <button class="ui small button" @click="stopDebug(0)">{{ $t('notebook.stop') }}</button>
              <span class="text">{{ $t('notebook.notebookRunning') }}</span>
            </div>
            <div v-else>
              <button class="ui disabled small button" style="background-color: #888;"><i class="loading spinner icon"
                  style="color: #fff;"></i>{{ $t('notebook.stopping') }}</button>
            </div>
          </div>

          <div class="resource-operate" v-if="selectIndex == 2">
            <div v-if="btnStatus[2] === 0">
              <button class="ui green small button" @click="createTask(2)">{{ $t('notebook.newTask') }}</button>
              <span v-if="notebookInfo.waitCountNpu == 0" class="text">{{ $t('notebook.noQuene') }}</span>
              <span v-else class="text">{{ $t('notebook.queneTips1') }} <span style="color: red;">{{
                notebookInfo.waitCountNpu + 1 }}</span> {{ $t('notebook.queneTips2') }}</span>
            </div>
            <div v-else-if="btnStatus[2] === 1">
              <button class="ui disabled small button" style="background-color: #888;"><i class="loading spinner icon"
                  style="color: #fff;"></i>{{ $t('notebook.newTask') }}</button>
              <span class="text">{{ $t('notebook.watiResource') }}</span>
            </div>

            <div v-else-if="btnStatus[2] === 2">
              <button class="ui small button" style="background-color: #1684fc;" @click="goDebug(2)">
                <a style="color:#fff" href="javascript:;">{{ $t('notebook.debug') }}</a>
              </button>
              <button class="ui small button" @click="stopDebug(2)">{{ $t('notebook.stop') }}</button>
              <span class="text">{{ $t('notebook.notebookRunning') }}</span>
            </div>
            <div v-else>
              <button class="ui disabled small button" style="background-color: #888;"><i class="loading spinner icon"
                  style="color: #fff;"></i>{{ $t('notebook.stopping') }}</button>
            </div>
          </div>
          <div class="resource-operate" v-if="selectIndex == 1">
            <div v-if="btnStatus[1] === 0">
              <button class="ui green small button" @click="createTask(1)">{{ $t('notebook.newTask') }}</button>
              <span v-if="notebookInfo.waitCountGpu == 0" class="text">{{ $t('notebook.noQuene') }}</span>
              <span v-else class="text">{{ $t('notebook.queneTips1') }} <span style="color: red;">{{
                notebookInfo.waitCountGpu + 1 }}</span> {{ $t('notebook.queneTips2') }}</span>
            </div>
            <div v-else-if="btnStatus[1] === 1">
              <button class="ui disabled small button" style="background-color: #888;"><i class="loading spinner icon"
                  style="color: #fff;"></i>{{ $t('notebook.newTask') }}</button>
              <span class="text">{{ $t('notebook.watiResource') }}</span>
            </div>

            <div v-else-if="btnStatus[1] === 2">
              <button class="ui small button" style="background-color: #1684fc;" @click="goDebug(1)">
                <a style="color:#fff" href="javascript:;">{{ $t('notebook.debug') }}</a>
              </button>
              <button class="ui small button" @click="stopDebug(1)">{{ $t('notebook.stop') }}</button>
              <span class="text">{{ $t('notebook.notebookRunning') }}</span>
            </div>
            <div v-else>
              <button class="ui disabled small button" style="background-color: #888;"><i class="loading spinner icon"
                  style="color: #fff;"></i>{{ $t('notebook.stopping') }}</button>
            </div>
          </div>
          <div style="margin-left:15px;">
            <div v-if="notebookInfo.cloudBrainPaySwitch && notebookInfo.PointAccount && !loading" class="text">
              <span>{{ $t('notebook.balanceOfPoints') }}
                <span style="color:rgb(250, 140, 22)">{{ notebookInfo.PointAccount.Balance.toFixed(2) }}</span>
                {{ $t('notebook.points') }}<span v-show="useUnitPrice > 0">{{ $t('notebook.expected_time') }}
                  <span style="color:rgb(250, 140, 22)"> {{ useTime }} </span>
                  {{ $t('notebook.hours') }}</span>
              </span>
              <a href="/reward/point/rule" target="_blank" style="margin: 0 1rem;text-decoration: underline;">{{
                $t('calcPointAcquisitionInstructions') }}</a>
            </div>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import { getFileNotebook, createNotebook, getCb1Notebook, getCb2Notebook, getFileInfoNotebook, stopNotebook } from "~/apis/modules/notobook";
import { getAiTaskBrief, getAiTaskDebugUrl, stopAiTask, getAiTaskImgesBySpec } from '~/apis/modules/cloudbrain';
import { Message } from "element-ui";
let timerCb1, timerCb2
let { AppSubUrl } = window.config
const finalState = [
  "STOPPED",
  "CREATE_FAILED",
  "UNAVAILABLE",
  "DELETED",
  "RESIZE_FAILED",
  "SUCCEEDED",
  "IMAGE_FAILED",
  "SUBMIT_FAILED",
  "DELETE_FAILED",
  "KILLED",
  "COMPLETED",
  "FAILED",
  "CANCELED",
  "LOST",
  "START_FAILED",
  "SUBMIT_MODEL_FAILED",
  "DEPLOY_SERVICE_FAILED",
  "CHECK_FAILED",
  "STOPPING",
  "CREATED_FAILED"
];
export default {
  data() {
    return {
      dialogVisible: false,
      selectIndex: 0,
      slideActive: true,
      initSelect: true,
      notebookInfo: { specCpu: { acc_cards_num: 0 }, specGpu: { acc_cards_num: 0 }, specNpu: { acc_cards_num: 0 } },
      fileInfo: {
        file: '',
        branch_name: '',
        owner_name: '',
        project_name: '',
        sign_name: ''
      },
      btnStatus: { 0: 0, 1: 0, 2: 0 },
      alertCb: false,
      deubgUrlNpu: '',
      deubgUrlGpu: '',
      deubgUrlNpuStop: '',
      deubgUrlGpuStop: '',
      loading: false,
      activeLoadFirst: true,

      useUnitPrice: '',
      useTime: '',
      selectedImage: '',
      npuImageList: [],
    };
  },
  methods: {
    handleClose(done) {
      this.initSelect = true
      this.alertCb = false
    },
    selectResource(index) {
      this.getNotebookInfo()
      if (index == this.selectIndex) {
        return
      }
      if (index > this.selectIndex && this.selectIndex !== 1) {
        this.slideActive = true
      } else if (index < this.selectIndex && index == 1) {
        this.slideActive = true
      } else {
        this.slideActive = false
      }
      this.selectIndex = index;
      this.initSelect = false
      this.alertCb = false
      this.getPointInfo()
    },
    getPointInfo() {
      if (this.notebookInfo.cloudBrainPaySwitch && this.notebookInfo.PointAccount) {
        const blance = this.notebookInfo.PointAccount.Balance;
        if (this.selectIndex == 0) {
          this.useUnitPrice = this.notebookInfo.specCpu.unit_price;
        } else if (this.selectIndex == 1) {
          this.useUnitPrice = this.notebookInfo.specGpu.unit_price;
        } else if (this.selectIndex == 2) {
          this.useUnitPrice = this.notebookInfo.specNpu.unit_price;
        }
        this.useTime = (blance / this.useUnitPrice).toFixed(2);
        if (blance < this.useUnitPrice) this.useTime = '0.00';
      }
    },
    getNpuImages(specId) {
      getAiTaskImgesBySpec({
        repoOwnerName: this.fileInfo.owner_name,
        repoName: this.fileInfo.project_name,
        jobType: 'DEBUG',
        computeSource: 'NPU',
        clusterType: 'C2Net',
        spec: specId,
        hasInternet: '2',
      }).then(res => {
        res = res.data
        if (res.code == 0) {
          this.npuImageList = res.data?.images || []
          const imageId = new URLSearchParams(window.location.search).has("image")
          if (imageId && this.npuImageList.map(item => item.image_id).includes(imageId)) {
            this.selectedImage = imageId
          } else {
            if (this.npuImageList[0]) {
              this.selectedImage = this.npuImageList[0].image_id
            }
          }
        }
      }).catch(err => {
        console.log(err);
      })
    },
    getNotebookInfo() {
      if (this.activeLoadFirst) {
        this.loading = true
      }
      getFileNotebook().then((res) => {
        if (res.data.code == 0) {
          this.notebookInfo = res.data
          if (this.activeLoadFirst) {
            this.getNpuImages(this.notebookInfo.specNpu.id)
          }
          this.getPointInfo()
        } else {
          Message.error(res.data.message)
        }
        this.loading = false
        this.activeLoadFirst = false
      }).catch((err) => {
        Message.error(err)
        this.loading = false
        this.activeLoadFirst = false
      })
    },
    getCb1NotebookInfo(id, index, data) {
      getAiTaskBrief({
        id: id
      }).then(res => {
        res = res.data;
        if (res.code == 0) {
          const _data = res.data;
          if (_data.status === "RUNNING") {
            let fileData = { id: Number(id), ...data }
            setTimeout(this.getFileInfoReadyNotebook(fileData, index), 0)
            clearInterval(timerCb1)
          }
          if (finalState.includes(_data.status)) {
            this.btnStatus[index] = 0
            clearInterval(timerCb1)
          }
        }
      }).catch(err => {
        console.log(err);
        this.btnStatus[index] = 0
        clearInterval(timerCb1)
        Message.error(err)
      });
    },
    getCb2NotebookInfo(id, index, data) {
      getAiTaskBrief({
        id: id
      }).then(res => {
        res = res.data;
        if (res.code == 0) {
          const _data = res.data;
          if (_data.status === "RUNNING") {
            let fileData = { id: Number(id), ...data }
            setTimeout(this.getFileInfoReadyNotebook(fileData, 2), 0)
            clearInterval(timerCb2)
          }
          if (finalState.includes(_data.status)) {
            Message.error(_data.status)
            this.btnStatus[2] = 0
            clearInterval(timerCb2)
          }
        }
      }).catch(err => {
        this.btnStatus[2] = 0
        clearInterval(timerCb2)
        Message.error(err)
      });
    },
    getFileInfoReadyNotebook(data, index) {
      let retryCount = 5;
      const self = this;
      function success(data, index) {
        if (index === 2) {
          self.btnStatus[2] = 2
          self.deubgUrlNpu = data.id
          self.deubgUrlNpuStop = data.id
        } else {
          self.btnStatus[index] = 2
          self.deubgUrlGpu = data.id
          self.deubgUrlGpuStop = data.id
        }
      }
      function getStatus(data, index) {
        getFileInfoNotebook(data).then((res) => {
          retryCount--;
          if (res.data.code === 0) {
            success(data, index)
          } else {
            if (retryCount > 0) {
              setTimeout(() => { getStatus(data, index) }, 200)
            } else {
              success(data, index)
              Message.error(res.data.message)
            }
          }
        }).catch((err) => {
          self.btnStatus[index] = 0
          Message.error(err)
        })
      }
      getStatus(data, index)
    },
    goDebug(index) {
      let id = index === 2 ? this.deubgUrlNpu : this.deubgUrlGpu
      getAiTaskDebugUrl({
        id: id,
        repoOwnerName: this.fileInfo.sign_name,
        repoName: this.notebookInfo.projectName,
        file: `${this.fileInfo.owner_name}/${this.fileInfo.project_name}/${this.fileInfo.branch_name}/${this.fileInfo.file}`,
      }).then((res) => {
        res = res.data;
        if (res.code == 0) {
          if (res.data && res.data.url) {
            window.open(res.data.url)
          }
        } else {
          Message.error(res.msg)
        }
      }).catch((err) => {
        this.btnStatus[index] = 0
        Message.error(err)
      })
    },
    stopDebug(index) {
      this.btnStatus[index] = 3
      let id = index === 2 ? this.deubgUrlNpuStop : this.deubgUrlGpuStop
      stopAiTask({
        id: id,
        repoOwnerName: this.fileInfo.sign_name,
        repoName: this.notebookInfo.projectName,
      }).then((res) => {
        res = res.data;
        if (res.code == 0) {
          this.btnStatus[index] = 0
          Message.success(this.$t("notebook.stopSuccess"))
        } else {
          this.btnStatus[index] = 0
          Message.error(res.msg)
        }
      }).catch((err) => {
        this.btnStatus[index] = 0
        Message.error(err)
      })
    },
    createTask(index) {
      if (index === 2) {
        const selectedImage = this.npuImageList.filter(item => item.image_id == this.selectedImage)[0]
        if (selectedImage) {
          this.fileInfo.image_id = selectedImage.image_id
          this.fileInfo.image = selectedImage.image_name
        } else {
          Message.error(this.$t("cloudbrainObj.selectImage"))
          return
        }
      }
      const data = { type: index, ...this.fileInfo }
      let repoPath = `repos/${this.fileInfo.owner_name}/${this.fileInfo.project_name}`
      this.btnStatus[index] = 1
      createNotebook(data).then((res) => {
        if (res.data.code === 0 && res.status === 200) {
          if (index === 2) {
            timerCb2 = setInterval(() => {
              setTimeout(this.getCb2NotebookInfo(res.data.message, index, data), 0)
            }, 10000)
          } else {
            timerCb1 = setInterval(() => {
              setTimeout(this.getCb1NotebookInfo(res.data.message, index, data), 0)
            }, 10000)
          }
          this.alertCb = false

        } else if (res.data.code == 2) {
          this.btnStatus[index] = 0
          this.alertCb = true
        } else {
          this.btnStatus[index] = 0
          Message.error(res.data.message)
        }
      }).catch((err) => {
        if (err.response.status === 403 && err.response.data.code === 1) {
          location.href = `${AppSubUrl}/authentication/wechat/bind`
        }
        if (err.response.status === 401) {
          location.href = `${AppSubUrl}/user/login?redirect_to=${encodeURIComponent(location.origin + location.pathname + '?card=' + this.selectIndex + '&image=' + this.selectedImage)}`
          return
        }
        this.btnStatus[index] = 0
        this.alertCb = false
        Message.error(err)

      })
    },
    handler() {
      this.getNotebookInfo()
      this.dialogVisible = true;
      this.getPointInfo()
    }
  },
  computed: {
    cpuSpec() {
      let gpuMemStr = this.notebookInfo.specCpu.gpu_mem_gi_b !== 0 ? `(${this.$t("notebook.graphicMemory")}: ${this.notebookInfo.specCpu.gpu_mem_gi_b}GB)` : '';
      let cpu_spec = `${this.$t("notebook.specification")}：GPU: ${this.notebookInfo.specCpu.acc_cards_num}*${this.notebookInfo.specCpu.acc_card_type}${gpuMemStr}, CPU: ${this.notebookInfo.specCpu.cpu_cores}`      
      if (this.notebookInfo.specCpu.mem_gi_b !== 0) {
        cpu_spec += `, ${this.$t("notebook.memory")}: ${this.notebookInfo.specCpu.mem_gi_b}GB`
      }
      if (this.notebookInfo.specCpu.share_mem_gi_b !== 0) {
        cpu_spec += `, ${this.$t("notebook.sharedMemory")}: ${this.notebookInfo.specCpu.share_mem_gi_b}GB`
      }
      return cpu_spec
    },
    npuSpec() {
      let acc_card_type = ''
      if (this.notebookInfo.specNpu.acc_card_type === "ASCEND910") {
        acc_card_type = "Ascend 910"
      }
      let gpuMemStr = this.notebookInfo.specNpu.gpu_mem_gi_b !== 0 ? `(${this.$t("notebook.graphicMemory")}: ${this.notebookInfo.specNpu.gpu_mem_gi_b}GB)` : '';
      let npu_spec = `${this.$t("notebook.specification")}：NPU: ${this.notebookInfo.specNpu.acc_cards_num}*${acc_card_type}${gpuMemStr}, CPU: ${this.notebookInfo.specNpu.cpu_cores}`      
      if (this.notebookInfo.specNpu.mem_gi_b !== 0) {
        npu_spec += `, ${this.$t("notebook.memory")}: ${this.notebookInfo.specNpu.mem_gi_b}GB`
      }
      if (this.notebookInfo.specNpu.share_mem_gi_b !== 0) {
        npu_spec += `, ${this.$t("notebook.sharedMemory")}: ${this.notebookInfo.specNpu.share_mem_gi_b}GB`
      }
      return npu_spec
    },
    npuImageType() {
      return this.notebookInfo.imageNpuDescription
    },
    gpuSpec() {
      let gpuMemStr = this.notebookInfo.specGpu.gpu_mem_gi_b !== 0 ? `(${this.$t("notebook.graphicMemory")}: ${this.notebookInfo.specGpu.gpu_mem_gi_b}GB)` : '';
      let gpu_spec = `${this.$t("notebook.specification")}：GPU: ${this.notebookInfo.specGpu.acc_cards_num}*${this.notebookInfo.specGpu.acc_card_type}${gpuMemStr}, CPU: ${this.notebookInfo.specGpu.cpu_cores}`     
      if (this.notebookInfo.specGpu.mem_gi_b !== 0) {
        gpu_spec += `, ${this.$t("notebook.memory")}: ${this.notebookInfo.specGpu.mem_gi_b}GB`
      }
      if (this.notebookInfo.specGpu.share_mem_gi_b !== 0) {
        gpu_spec += `, ${this.$t("notebook.sharedMemory")}: ${this.notebookInfo.specGpu.share_mem_gi_b}GB`
      }
      return gpu_spec
    }
  },
  beforeDestroy() {
    clearInterval(timerCb1)
    clearInterval(timerCb2)
    document
      .querySelector("#notebook-debug")
      .removeEventListener("click", this.handler);
  },
  mounted() {
    const selfData = document.querySelector('#__vue-self-data')
    this.fileInfo.file = selfData.getAttribute('data-file')
    this.fileInfo.branch_name = selfData.getAttribute('data-branch')
    this.fileInfo.owner_name = selfData.getAttribute('data-owner')
    this.fileInfo.project_name = selfData.getAttribute('data-project')
    this.fileInfo.sign_name = selfData.getAttribute('data-name')
    let that = this;
    const params = new URLSearchParams(window.location.search)
    if (params.has("card")) {
      that.getNotebookInfo()
      that.dialogVisible = true;
      that.selectIndex = Number(new URLSearchParams(window.location.search).get("card"))
      that.getPointInfo()
    }
    document
      .querySelector("#notebook-debug")
      .addEventListener("click", this.handler);
  },
};
</script>
<style scoped lang="less">
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
  font-family: Roboto;

  .el-dialog__title {
    font-weight: 600;
    font-size: 15px;
    color: rgb(16, 16, 16);
  }

  .el-dialog__headerbtn {
    top: 15px;
    right: 15px;
  }
}

/deep/ .el-dialog__body {
  padding: 55px 15px 0 15px;
}

/deep/ .el-icon-loading {
  font-size: 40px;
}

.wrapper {
  width: 100%;

  .active {
    background: linear-gradient(269.2deg,
        rgba(183, 247, 255, 0.5) 0%,
        rgba(233, 233, 255, 0) 78.67%);
    border-radius: 5px 5px 0px 0px;
  }

  .text-tip {
    color: #888;
    font-size: 12px;
  }

  .text-tip::before {
    content: '*';
    color: #f2711c;
  }

  .alert-info {
    width: 75%;
    background-color: rgba(242, 113, 28, 0.05);
    border: 1px solid rgb(242, 113, 28);
    border-radius: 5px;
    margin: 0 auto;
    padding-bottom: 10px;
  }

  &>.three-resource-type:nth-child(2) {
    border: none;
  }

  .three-resource-type {
    width: 75%;
    margin: 0 auto;
    display: flex;
    border-top: 1px solid rgba(16, 16, 16, 0.1);
    cursor: pointer;

    .resource-child-node {
      display: flex;
      align-items: center;
      width: 100%;
      height: 115px;
      position: relative;

      .resource-type-icon {
        width: 50px;
        height: 50px;
        line-height: 20px;
        border-radius: 25px;
        text-align: center;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-shrink: 0;

        .text {
          font-size: 26px;
          color: rgba(251, 251, 251, 1);
          font-family: ZKXiaoWeiLogo-regular;
        }
      }

      .background-C {
        background: linear-gradient(134.2deg,
            rgba(130, 209, 246, 1) 0%,
            rgba(41, 182, 244, 1) 51.94%,
            rgba(0, 137, 205, 1) 102.83%);
      }

      .background-N {
        background: linear-gradient(151.47deg,
            rgba(123, 50, 178, 1) 20.02%,
            rgba(64, 26, 93, 1) 100%);
      }

      .background-G {
        background: linear-gradient(-25.25deg,
            rgba(254, 86, 77, 1) 9.3%,
            rgba(251, 155, 54, 1) 38.86%,
            rgba(249, 202, 38, 1) 67.95%);
      }
    }

    .resource-type-detail {
      margin-left: 23px;

      .detail-title {
        font-family: SourceHanSansSC;
        font-weight: 500;
        font-size: 16px;
        color: rgb(16, 16, 16);
        font-style: normal;
        letter-spacing: 0px;
        line-height: 32px;
        text-decoration: none;
      }

      .detail-spec {
        font-family: SourceHanSansSC;
        font-weight: 400;
        font-size: 14px;
        color: rgba(136, 136, 136, 1);
        font-style: normal;
        letter-spacing: 0px;
        line-height: 24px;
        text-decoration: none;

        .select-image {
          color: rgb(50, 145, 248);
          display: flex;
          align-items: center;

        }

        .select-image-label {
          border: none !important;
          background: transparent;
          color: rgb(50, 145, 248);
          padding-left: 0;
          margin-left: -5px;
          outline: none;
        }
      }
    }

    .resource-select {
      margin-left: auto;
      margin-right: 20px;
      font-size: 20px;
      height: 100%;
      display: flex;
      align-items: center;

      .green {
        color: green;
      }

      .gray {
        color: rgba(16, 16, 16, 0.1);
      }
    }

    .resource-unit-price {
      position: absolute;
      bottom: 30px;
      right: 20px;
      color: rgb(16, 16, 16);
    }
  }

  .resource-footer {
    margin-top: 40px;
    border-top: 1px solid rgba(16, 16, 16, 0.1);
    height: 71px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .resource-operate {
      display: flex;
      align-items: center;

      .text {
        color: #101010;
        margin-left: 20px;
      }
    }

    .text {
      color: #101010;
    }
  }

  @media screen and (max-width:1600px) {
    .resource-footer {
      display: block;

      .text {
        text-align: right;
      }
    }
  }

  .slide-in-top {
    -webkit-animation: slide-in-top 1s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    animation: slide-in-top 1s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
  }

  .slide-in-bottom {
    -webkit-animation: slide-in-bottom 1s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    animation: slide-in-bottom 1s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
  }

  .fade-out {
    -webkit-animation: fade-in 1.2s cubic-bezier(0.39, 0.575, 0.565, 1) both;
    animation: fade-in 1.2s cubic-bezier(0.39, 0.575, 0.565, 1) both;
    position: absolute
  }

  @-webkit-keyframes slide-in-top {
    0% {
      -webkit-transform: translateY(-50px);
      transform: translateY(-50px);
      opacity: 0;
    }

    100% {
      -webkit-transform: translateY(0);
      transform: translateY(0);
      opacity: 1;
    }
  }

  @keyframes slide-in-top {
    0% {
      -webkit-transform: translateY(-50px);
      transform: translateY(-50px);
      opacity: 0;
    }

    100% {
      -webkit-transform: translateY(0);
      transform: translateY(0);
      opacity: 1;
    }
  }

  @-webkit-keyframes slide-in-bottom {
    0% {
      -webkit-transform: translateY(50px);
      transform: translateY(50px);
      opacity: 0;
    }

    100% {
      -webkit-transform: translateY(0);
      transform: translateY(0);
      opacity: 1;
    }
  }

  @keyframes slide-in-bottom {
    0% {
      -webkit-transform: translateY(50px);
      transform: translateY(50px);
      opacity: 0;
    }

    100% {
      -webkit-transform: translateY(0);
      transform: translateY(0);
      opacity: 1;
    }
  }

  @-webkit-keyframes fade-in {
    0% {
      opacity: 1;
    }

    100% {
      opacity: 0;
    }
  }

  @keyframes fade-in {
    0% {
      opacity: 1;
    }

    100% {
      opacity: 0;
    }
  }

}
</style>
