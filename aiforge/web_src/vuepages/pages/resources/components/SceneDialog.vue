<template>
  <div class="base-dlg">
    <BaseDialog :visible.sync="dialogShow" :width="`900px`"
      :title="type === 'add' ? $t('resourcesManagement.addResScene') : $t('resourcesManagement.editResScene')"
      @open="open" @opened="opened" @close="close" @closed="closed">
      <div class="dlg-content">
        <div class="form">
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.resSceneName') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.SceneName" placeholder="" maxlength="255"></el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.jobType') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.JobType" :disabled="type === 'edit'">
                <el-option v-for="item in taskTypeList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.sceneType') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.SceneType" @change="changeSceneType">
                <el-option v-for="item in sceneTypeList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row" v-if="dataInfo.SceneType == 'public'">
            <div class="title required">
              <span>{{ $t('resourcesManagement.isExclusiveSpec') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.IsSpecExclusive" @change="changeIsSpecExclusive">
                <el-option v-for="item in isSpecExclusiveList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row" v-if="dataInfo.IsSpecExclusive === 'exclusive' || dataInfo.SceneType == 'exclusive'">
            <div class="title required">
              <span>{{ $t('resourcesManagement.exclusiveOrg') }}</span>
            </div>
            <div class="content">
              <el-input v-model="dataInfo.ExclusiveOrg" :placeholder="$t('resourcesManagement.exclusiveOrgTips')"
                maxlength="255">
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.computeCluster') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.Cluster" @change="changeCluster" :disabled="type === 'edit'">
                <el-option v-for="item in clusterList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <div class="form-row">
            <div class="title required">
              <span>{{ $t('resourcesManagement.computeResource') }}</span>
            </div>
            <div class="content">
              <el-select v-model="dataInfo.Resource" @change="changeResource" :disabled="type === 'edit'">
                <el-option v-for="item in resourceList" :key="item.k" :label="item.v" :value="item.k" />
              </el-select>
            </div>
          </div>
          <SpecSelect v-model="dataInfo.SpecIds" :specs="specsList"></SpecSelect>
          <div class="form-row" style="margin-top:20px">
            <div class="title" style="width:285px;"></div>
            <div class="content">
              <el-button type="primary" class="btn confirm-btn" @click="confirm">{{ $t('confirm') }}</el-button>
              <el-button class="btn" @click="cancel">{{ $t('cancel') }}</el-button>
            </div>
          </div>
        </div>
      </div>
    </BaseDialog>
  </div>
</template>
<script>
import BaseDialog from '~/components/BaseDialog.vue';
import SpecSelect from './SpecSelect.vue';
import { getResSpecificationListAll, addResScene, updateResScene } from '~/apis/modules/resources';
import { JOB_TYPE, CLUSTERS, ACC_CARD_TYPE, SPECIFICATION_STATUS, COMPUTER_RESOURCES, NETWORK_TYPE_VALUE } from '~/const';
import { getListValueWithKey } from '~/utils';

export default {
  name: "SceneDialog",
  props: {
    visible: { type: Boolean, default: false },
    title: { type: String, default: '' },
    type: { type: String, defalut: 'add' },
    data: { type: Object, default: () => ({}) },
  },
  components: { BaseDialog, SpecSelect },
  data() {
    return {
      dialogShow: false,
      dataInfo: {},
      taskTypeList: [...JOB_TYPE],
      clusterList: [...CLUSTERS],
      accCardTypeList: [...ACC_CARD_TYPE],
      statusList: [...SPECIFICATION_STATUS],
      sceneTypeList: [{ k: 'public', v: this.$t('resourcesManagement.public') }, { k: 'exclusive', v: this.$t('resourcesManagement.exclusive') }],
      isSpecExclusiveList: [{ k: 'public', v: this.$t('resourcesManagement.commonUseSpec') }, { k: 'exclusive', v: this.$t('resourcesManagement.exclusiveSpec') }],
      resourceList: [...COMPUTER_RESOURCES],
      networkTypeList: [...NETWORK_TYPE_VALUE],
      specsList: [],
    };
  },
  watch: {
    visible: function (val) {
      this.dialogShow = val;
    },
  },
  methods: {
    resetDataInfo() {
      this.dataInfo = {
        SceneName: '',
        JobType: '',
        SceneType: '',
        IsSpecExclusive: '',
        ExclusiveOrg: '',
        Cluster: 'OpenI',
        Resource: 'GPU',
        SpecIds: [],
      }
      this.specsList.splice(0, Infinity);
    },
    getResSpecificationList() {
      const params = {
        cluster: this.dataInfo.Cluster,
        resource: this.dataInfo.Resource,
        queue: this.dataInfo.QueueId === '-1' ? '' : this.dataInfo.QueueId,
        available: 1,
      };
      return getResSpecificationListAll(params).then(res => {
        res = res.data;
        if (res.Code === 0) {
          const list = res.Data.Specs;
          const data = list.map((item) => {
            const NGPU = `${item.ComputeResource}:${item.AccCardsNum + '*' + getListValueWithKey(this.accCardTypeList, item.AccCardType)}`;
            const statusStr = item.Status != '2' ? `<span style="color:rgb(245, 34, 45)">(${getListValueWithKey(this.statusList, item.Status.toString())})</span>` : '';
            const queueName = item.QueueName ? `【${item.QueueName}】` : '';
            const queueType = item.QueueType ? `【${item.QueueType}】` : '';
            return {
              ...item,
              StatusStr: statusStr,
              QueueStr: `${item.QueueCode}${queueName}${queueType}(${getListValueWithKey(this.clusterList, item.Cluster)} - ${item.AiCenterName})`,
              QueueIsExclusiveStr: item.IsQueueExclusive ? `<span style="color:rgb(245, 34, 45);font-weight:bold">(${this.$t('resourcesManagement.exclusiveTxt')})</span>` : '',
              SpecStr: `${NGPU}(${this.$t('resourcesManagement.gpuMem')}:${item.GPUMemGiB}GB), CPU:${item.CpuCores}, ${this.$t('resourcesManagement.mem')}:${item.MemGiB}GB, ${this.$t('resourcesManagement.shareMem')}:${item.ShareMemGiB}GB`,
              PriceStr: `, ${this.$t('resourcesManagement.unitPrice')}:${item.UnitPrice.toFixed(2)}${this.$t('resourcesManagement.point_hr')}`,
              NetworkTypeStr: `, ${this.$t('cloudbrainObj.networkType')}:${getListValueWithKey(this.networkTypeList, item.HasInternet)}`,
              VisualizationStr: `, ${this.$t('cloudbrainObj.visualization')}:${item.EnableVisualization ? this.$t('resourcesManagement.enable') : this.$t('resourcesManagement.notEnable')}`,
            }
          });
          this.specsList.splice(0, Infinity, ...data);
        }
      }).catch(err => {
        console.log(err);
      });
    },
    changeSceneType() {
      this.dataInfo.IsSpecExclusive = '';
      this.changeIsSpecExclusive();
    },
    changeIsSpecExclusive() {
      this.dataInfo.ExclusiveOrg = '';
    },
    changeCluster() {
      this.dataInfo.SpecIds = [];
      this.specsList.splice(0, Infinity);
      this.getResSpecificationList();
    },
    changeResource() {
      this.dataInfo.SpecIds = [];
      this.specsList.splice(0, Infinity);
      this.getResSpecificationList();
    },
    open() {
      this.resetDataInfo();
      if (this.type === 'add') {
        //
      } else if (this.type === 'edit') {
        Object.assign(this.dataInfo, {
          ID: this.data.ID,
          SceneName: this.data.SceneName,
          JobType: this.data.JobType,
          SceneType: this.data.SceneType,
          IsSpecExclusive: this.data.IsSpecExclusive,
          ExclusiveOrg: this.data.ExclusiveOrg,
          Cluster: this.data.Cluster,
          Resource: this.data.ComputeResource,
          SpecIds: [...this.data.SpecIds],
        });
      }
      this.$emit("open");
      this.getResSpecificationList();
    },
    opened() {
      this.$emit("opened");
    },
    close() {
      this.$emit("close");
    },
    closed() {
      this.$emit("closed");
      this.$emit("update:visible", false);
    },
    confirm() {
      if (!this.dataInfo.SceneName || !this.dataInfo.JobType || !this.dataInfo.SceneType || !this.dataInfo.SpecIds.length
        || (this.dataInfo.IsSpecExclusive === 'exclusive' && !this.dataInfo.ExclusiveOrg)
        || (this.dataInfo.SceneType === 'exclusive' && !this.dataInfo.ExclusiveOrg)
      ) {
        this.$message({
          type: 'info',
          message: this.$t('pleaseCompleteTheInformationFirst')
        });
        return;
      }
      // console.log('submit', this.dataInfo);
      // return;
      const setApi = this.type === 'add' ? addResScene : updateResScene;
      setApi({
        ...this.dataInfo,
        action: this.type === 'edit' ? 'edit' : undefined,
        IsSpecExclusive: this.dataInfo.IsSpecExclusive,
      }).then(res => {
        res = res.data;
        if (res.Code === 0) {
          this.$message({
            type: 'success',
            message: this.$t('submittedSuccessfully')
          });
          this.$emit("confirm");
        } else {
          this.$message({
            type: 'error',
            message: this.$t('submittedFailed')
          });
        }
      }).catch(err => {
        console.log(err);
        this.$message({
          type: 'error',
          message: this.$t('submittedFailed')
        });
      })
    },
    cancel() {
      this.dialogShow = false;
      this.$emit("update:visible", false);
    }
  },
  mounted() {
    this.resetDataInfo();
  },
};
</script>
<style scoped lang="less">
.dlg-content {
  margin: 20px 0 25px 0;
  display: flex;
  justify-content: center;

  .form {
    width: 800px;

    .form-row {
      display: flex;
      min-height: 42px;
      margin-bottom: 4px;

      .title {
        width: 255px;
        display: flex;
        justify-content: flex-end;
        align-items: center;
        margin-right: 20px;
        color: rgb(136, 136, 136);
        font-size: 14px;

        &.required {
          span {
            position: relative;
          }

          span::after {
            position: absolute;
            right: -10px;
            top: -2px;
            vertical-align: top;
            content: '*';
            color: #db2828;
          }
        }
      }

      .content {
        width: 300px;
        display: flex;
        align-items: center;

        /deep/ .el-select {
          width: 100%;
        }
      }

      .specSel {
        /deep/ .el-tag.el-tag--info {
          max-width: 81%;
          display: flex;
          align-items: center;

          .el-select__tags-text {
            overflow: hidden;
            text-overflow: ellipsis;
          }

          .el-tag__close {
            flex-shrink: 0;
            right: -5px;
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
    }
  }
}

.el-select-dropdown__item {
  padding-left: 26px !important;
}

.el-select-dropdown__item.selected::after {
  right: 0;
  left: 6px;
}
</style>
