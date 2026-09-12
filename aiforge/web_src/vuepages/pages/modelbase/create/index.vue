<template>
  <div>
    <div class="ui container area">
      <div class="area-title">新建鹏城·脑海(原鹏城·盘古)大模型微调任务</div>
      <div v-if="alreadyMsgBoxShow">
        <div class="err-msg-box-already">
          <div class="msg-content">
            <i class="ri-information-line"></i>
            <div class="msg-content-tip">
              <div class="line-1" v-html="$t('cloudbrainObj.sameTaskTips1', { count: notStopTaskCount })"></div>
              <div class="line-2" v-html="$t('cloudbrainObj.sameTaskTips2')"></div>
            </div>
          </div>
        </div>
      </div>
      <div class="area-content">
        <div class="content">
          <div class="ui container">
            <div class="tips-c">
              <p><span>*</span> 本次新建的训练任务会放在您名下项目<span>openi-notebook</span>中，如果没有该项目系统会自动新建一个。</p>
            </div>
            <div class="main-title">基本信息：</div>
            <div class="row">
              <div class="row-title"><span>模型名称</span></div>
              <div class="row-content" style="padding-top:6px;"><span>鹏城 · 脑海(原鹏城 · 盘古)</span></div>
            </div>
            <div class="row">
              <div class="row-title"><span class="required">模型规模</span></div>
              <div class="row-content">
                <el-select v-model="modelSize">
                  <el-option v-for="(item, index) in modelSizeList" :key="index" :value="item.key"
                    :label="item.value"></el-option>
                </el-select>
              </div>
            </div>
            <TaskName ref="taskNameRef" :userName="userName"></TaskName>
            <div class="row">
              <div class="row-title"><span>任务描述</span></div>
              <div class="row-content">
                <el-input type="textarea" :rows="4" placeholder="描述字符不超过255个字符" v-model="descr"
                  :maxlength="255"></el-input>
              </div>
            </div>
            <div class="main-title">任务配置：</div>
            <ModelBaseDatasetSelect ref="modelBaseDatasetSelectRef" :tabindex="datasetTab"
              :exampleType="exampleDatasetType" :datasetId="datasetId" :userName="userName" :repoName="repoName">
            </ModelBaseDatasetSelect>
            <ResourceSpecification ref="resourceSpecificationRef" :specData="specData" :blance="pointBlance"
              :showPoint="pointShow" :specOri="spec"></ResourceSpecification>
            <RunParameters ref="runParametersRef" :required="true" :oriData="paramsOriData"></RunParameters>
            <div class="row">
              <div class="row-title"></div>
              <div class="row-content" style="padding-top:6px;">
                <el-button type="primary" class="btn confirm-btn" size="default" :disabled="alreadyMsgBoxShow"
                  @click="repoCheck" :loading="loading">
                  提 交</el-button>
                <el-button @click="cancel" class="btn" size="default">取 消</el-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <LoadingMask :loading="loading" :tips="loadingTips"></LoadingMask>
  </div>
</template>

<script>
import TaskName from '../components/cloudbrain/TaskName.vue';
import ModelBaseDatasetSelect from '../components/ModelBaseDatasetSelect.vue';
import ResourceSpecification from '../components/cloudbrain/ResourceSpecification.vue';
import RunParameters from '../components/cloudbrain/RunParameters.vue';
import LoadingMask from '../components/cloudbrain/LoadingMask.vue';
import { getUrlSearchParams, getListValueWithKey, transFileSize, renderSpecStr } from '~/utils';
import { getCheckRepo, getSpecInfo, setFinetuneCreate } from '~/apis/modules/modelbase';
import { getPointAccountInfo } from '~/apis/modules/common';

const notStopTaskCount = window.notStopTaskCount || 0;

export default {
  data() {
    return {
      userName: '',
      repoName: 'openi-notebook',

      modelSize: '2.6',
      modelSizeList: [{
        key: '2.6',
        value: '2.6 B',
      }],
      descr: '',

      datasetTab: 0,
      exampleDatasetType: 1,
      datasetId: '',

      spec: '',
      specData: [],
      pointBlance: 0,
      pointShow: false,

      paramsOriData: [{ k: 'train_iters', v: '40' }],

      alreadyMsgBoxShow: notStopTaskCount >= 1,
      notStopTaskCount: notStopTaskCount,
      loading: false,
      loadingTips: '任务正在准备中，喝杯水回来再看看~',
    };
  },
  components: { TaskName, ModelBaseDatasetSelect, ResourceSpecification, RunParameters, LoadingMask },
  methods: {
    repoCheck() {
      this.loading = true;
      getCheckRepo().then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 0) {
          this.datasetId = res.message;
          this.submitCheck();
        } else {
          this.$message({
            type: 'error',
            message: res.message,
          });
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    submitCheck() {
      const r1 = this.$refs.taskNameRef.check();
      const r2 = this.$refs.modelBaseDatasetSelectRef.check();
      const r3 = this.$refs.resourceSpecificationRef.check();
      const r4 = this.$refs.runParametersRef.check();
      if (r1 && r2 && r3 && r4) {
        const subData = {
          userName: this.userName,
          type: 0, // 0为脑海(盘古2.6模型)
          model_size: this.modelSize,
          display_job_name: this.$refs.taskNameRef.taskName,
          description: this.descr.trim(),
        };

        const specData = this.$refs.resourceSpecificationRef.getSpecData();
        subData.spec_id = Number(specData.id);

        const parameters = this.$refs['runParametersRef'].getParameter();
        const parameterObj = {
          parameter: parameters.map((item) => { return { ...item } }),
        };
        subData.run_para_list = JSON.stringify(parameterObj);

        const datasetTab = this.$refs.modelBaseDatasetSelectRef.tabIndex;
        if (datasetTab == 1) { // 本地上传
          this.loading = true;
          this.$refs.modelBaseDatasetSelectRef.localUpload().then(res => {
            if (res) {
              if (res.status == '1') {
                this.loading = false;
                this.$message({
                  type: 'error',
                  message: res.msg,
                });
              } else if (res.status == '0') {
                if (res.total == 1) {
                  const file = res.files[0];
                  if (res.total == res.success) { // 上传完成
                    subData.attachment = file.uuid;
                    subData.dataset_name = file.name;
                    this.submit(subData);
                    return;
                  }
                  if (file.uploaded == '1') { // 文件已经上传过
                    subData.attachment = file.uuid;
                    subData.dataset_name = file._fileName;
                    this.submit(subData);
                    return;
                  }
                  this.loading = false;
                  this.$message({
                    type: 'error',
                    message: '本地数据集文件上传失败',
                  });
                } else {
                  this.loading = false;
                  this.$message({
                    type: 'error',
                    message: '本地数据集文件上传失败',
                  });
                }
              }
            } else {
              this.loading = false;
              this.$message({
                type: 'error',
                message: '本地数据集文件上传失败',
              });
            }
          }).catch(err => {
            console.log(err);
            this.loading = false;
            this.$message({
              type: 'error',
              message: '本地数据集文件上传失败',
            });
          });
        } else {
          if (datasetTab == 0) {
            subData.sample_dataset_type = this.$refs.modelBaseDatasetSelectRef.sampleDataset;
          } else if (datasetTab == 2) {
            const datasetList = this.$refs.modelBaseDatasetSelectRef.getPlatformDataset();
            subData.attachment = datasetList[0].id;
            subData.dataset_name = datasetList[0].name;
          }
          this.submit(subData);
        }
      }
    },
    submit(data) {
      this.loading = true;
      setFinetuneCreate(data).then(res => {
        res = res.data;
        if (res && res.code == 0) {
          window.location.href = '/extension/modelbase/pangufinetune';
        } else {
          this.loading = false;
          this.$message({
            type: 'error',
            duration: 4000,
            message: res.message,
          });
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
        this.$message({
          type: 'error',
          message: '创建任务失败',
        });
      });
    },
    cancel() {
      window.location.href = `/extension/modelbase/pangufinetune`;
    }
  },
  beforeMount() {
    const urlParams = getUrlSearchParams();
    const type = urlParams.type;
    if (type == '0' || type == '1' || type == '2' || type == '3') {
      this.exampleDatasetType = Number(type);
      if (type == '0') {
        this.datasetTab = 1;
      }
    }
    const metaEl = document.querySelectorAll('meta[name="_uid"]');
    if (!metaEl.length) {
      window.location.href = `/user/login?redirect_to=${encodeURIComponent(window.location.href)}`;
      return;
    } else {
      const uid = metaEl[0].getAttribute('content');
      const uname = metaEl[0].getAttribute('content-ext');
      this.userName = uname;
    }
    getPointAccountInfo().then(res => {
      const data = res.data;
      this.pointShow = data.cloudBrainPaySwitch ? true : false;
      this.pointBlance = data.pointAccount ? data.pointAccount.balance : 0;
    }).catch(err => {
      console.log(err);
    });
    getSpecInfo({}).then(res => {
      const data = res.data || [];
      this.specData = data;
      if (data.length) {
        this.spec = data[0].id.toString();
      }
    }).catch(err => {
      console.log(err);
    });
  },
  mounted() { },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.area {
  width: 1050px !important;
  margin-top: 40px;

  .area-title {
    height: 45px;
    border-color: rgb(212, 212, 213);
    border-width: 1px;
    border-style: solid;
    border-radius: 5px 5px 0px 0px;
    font-size: 14px;
    background: rgb(240, 240, 240);
    line-height: 45px;
    padding-left: 15px;
    font-weight: 550;
    font-size: 16px;
    color: rgb(16, 16, 16);
  }

  .err-msg-box-already {
    margin: 1em 0;
    padding: 1em 1.5em;
    background-color: rgba(242, 113, 28, 0.05);
    border: 1px solid rgba(242, 113, 28, 1);
    border-radius: 5px;

    .msg-content {
      display: flex;
      align-items: center;

      i {
        font-size: 35px;
        color: rgba(242, 113, 28, 1);
      }

      .msg-content-tip {
        text-align: left;
        margin-left: 1rem;

        .line-1 {
          font-weight: 600;
          line-height: 2;

          span {
            color: rgba(242, 113, 28, 1);
          }
        }

        .line-2 {
          color: #939393
        }
      }
    }
  }

  .area-content {
    border-color: rgb(212, 212, 213);
    border-width: 1px;
    border-style: solid;
    margin-top: -1px;

    .content {
      padding: 20px 180px 24px 0;

      .main-title {
        font-weight: 550;
        font-size: 16px;
        color: rgb(16, 16, 16);
        height: 22px;
        margin: 30px 0 30px 110px;
      }

      .row {
        display: flex;
        margin-bottom: 28px;

        .row-title {
          width: 200px;
          text-align: right;
          margin-right: 24px;
          color: #101010;
          font-size: 14px;
          display: flex;
          justify-content: flex-end;
          padding-top: 6px;

          .required {
            position: relative;

            &::after {
              position: absolute;
              content: "*";
              top: -3px;
              right: -10px;
              color: red;
            }
          }
        }

        .row-content {
          flex: 1;
        }
      }
    }
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

    &.is-disabled {
      opacity: .45 !important;
    }
  }
}

.tips-c {
  color: #888888;
  font-size: 14px;
  margin-top: 10px;
  margin-bottom: -6px;
  margin-left: 173px;
  text-align: center;

  p {
    span {
      color: #f2711c;
    }
  }
}
</style>
