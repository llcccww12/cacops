<template>
  <div style="margin: 0 20px;height: 100%;" v-loading="loading">
    <div class="page-title"><span>{{ $t('userRole.editUserPermissions') }}：{{ name }}</span></div>
    <div class="operate-wrap">
      <div style="height:100%">
        <div class="sub-content-wrap"><span class="content">{{ $t('userRole.OperationalPermissions') }}</span></div>
        <div class="operate-container">
          <div class="operate-list">
            <div class="title">{{ $t('userRole.opCategoryRole') }}：</div>
            <div class="checkbox-wrap">
              <el-checkbox-group v-model="operateCheckList" class="check-wrap"
                style="display:flex;flex-direction:column">
                <el-checkbox v-for="item in operateList" :key="item.ID" :label="item.ID">
                  <span class="nowrap" style="display:inline-block;max-width:100%"
                    :title="item.Name + '(' + item.Description + ')'">{{ item.Name }} ({{ item.Description }})</span>
                  <span class="detail-a" @click.stop.prevent="showOpDetail(item)">{{ $t('userRole.detail') }}</span>
                </el-checkbox>
              </el-checkbox-group>
            </div>
          </div>
          <div class="operate-collection">
            <div class="title">{{ $t('userRole.permissionList') }}：</div>
            <ul>
              <li v-for="item in showOperList" :key="item">{{ item }}</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
    <div class="resource-wrap">
      <div class="sub-content-wrap"><span class="content">{{ $t('userRole.CResourcePermissions') }}</span></div>
      <div class="resource-container">
        <div class="title">{{ $t('userRole.reCategoryRole') }}：</div>
        <div>
          <el-checkbox-group v-model="resourceCheckList" class="check-wrap">
            <el-checkbox v-for="item in resourceList" :key="item.ID" :label="item.ID">
              <span class="nowrap" style="display:inline-block;max-width:90%"
                :title="item.Name + '(' + item.Description + ')'">{{ item.Name }} ({{ item.Description }}) </span>
              <span class="detail-a" @click.stop.prevent="showReDetail(item)">{{ $t('userRole.detail') }}</span>
            </el-checkbox>
          </el-checkbox-group>
        </div>
      </div>
    </div>
    <div class="resource-wrap">
      <div class="sub-content-wrap"><span class="content">{{ $t('userRole.StoragePermissions') }}</span></div>
      <div class="resource-container">
        <div class="title">{{ $t('userRole.stCategoryRole') }}：</div>
        <div>
          <el-checkbox-group v-model="storageCheckList" class="check-wrap">
            <el-checkbox v-for="item in storageList" :key="item.ID" :label="item.ID">
              <span class="nowrap" style="display:inline-block;max-width:90%"
                :title="item.Name + '(' + item.Description + ')'">{{ item.Name }} ({{ item.Description }}) </span>
              <span class="detail-a" @click.stop.prevent="showStDetail(item)">{{ $t('userRole.detail') }}</span>
            </el-checkbox>
          </el-checkbox-group>
        </div>
      </div>
    </div>
    <div style="display:flex">
      <div class="btn-wrap" @click="updateRole"><span>{{ $t('userRole.updatePermissions') }}</span></div>
      <div class="btn-wrap" @click="cancel" style="margin-left:1rem;background:#c2c7cc;"><span
          style="color: #020004;">{{ $t('cancel') }}</span></div>
    </div>

    <div class="resource-title">
      <span>{{ name }} {{ $t('userRole.resourcePermissions') }}：</span>
    </div>
    <resourceTable v-if="Object.keys(showSpecTableData).length" ref="resourceTable"
      :showSpecTableData="showSpecTableData" :specificationList="specificationList"></resourceTable>
    <BaseDialog :visible.sync="reRoleDialogShow" :title="$t('userRole.viewReRole')" width="80%" @closed="closed">
      <resourceTableCopy ref="reDialogTable" :radioInit="radioInit" :showSpecTableData="reRoleDialogData"
        :specificationList="specificationList" :reRoleCommonData="reRoleCommonData"></resourceTableCopy>
    </BaseDialog>
    <opRoleDialog :visible.sync="opRoleDialogShow" type="view" :data="opRoleDialogData"></opRoleDialog>
    <storageRoleDialog :visible.sync="stRoleDialogShow" type="view" :data="stRoleDialogData"></storageRoleDialog>
  </div>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';
import opRoleDialog from '../role/opRoleDialog.vue';
import resourceRoleDialog from '../role/resourceRoleDialog.vue';
import storageRoleDialog from '../role/storageRoleDialog.vue';
import resourceTable from './resourceTable.vue'
import resourceTableCopy from './resourceTable.vue'
import { listAiforgeRole, getResSpecificationListAll, listOperation, setAiforgeRoleToUser } from '~/apis/modules/resources';
import { getListValueWithKey, getUrlSearchParams } from '~/utils';
import { ACC_CARD_TYPE, NETWORK_TYPE_VALUE, NEW_JOB_TYPE_OBJ, NEW_JOB_TYPE } from '~/const';
export default {
  data() {
    return {
      operateList: [],
      resourceList: [],
      storageList: [],
      operateCheckList: [],
      resourceCheckList: [],
      storageCheckList: [],
      operMap: new Map(),
      showOperList: [],
      accCardTypeList: [...ACC_CARD_TYPE],
      networkTypeList: [...NETWORK_TYPE_VALUE],
      specificationList: [],
      radioInit: '',
      taskTypeList: [...NEW_JOB_TYPE],
      showSpecTableData: {},
      filterData: [],
      id: 0,
      name: '',
      opRoleDialogShow: false, // 操作权限弹窗
      opRoleDialogData: {},
      reRoleDialogShow: false, // 资源权限弹窗
      reRoleDialogData: {},
      stRoleDialogShow: false, // 操作权限弹窗
      stRoleDialogData: {},
      reRoleCommonData: {},
      cancelUrl: '',
      loading: false
    };
  },
  components: { opRoleDialog, resourceRoleDialog, storageRoleDialog, resourceTable, resourceTableCopy, BaseDialog },
  computed: {
    operJsonList() {
      return this.operateList.map((item) => {
        const data = JSON.parse(item.RightInfo)
        const operArray = data.map((item) => item.operName)
        return {
          ID: item.ID,
          operArray: operArray
        }
      })
    },
    resourceJsonList() {
      return this.resourceList.map((item) => {
        const data = JSON.parse(item.RightInfo)
        const taskTypes = data.map((item) => item.taskType)
        const specIds = data.map((item) => item.specId)
        return {
          ID: item.ID,
          taskTypes: taskTypes,
          specIds: specIds
        }
      })
    }
  },
  watch: {
    operateCheckList(val) {
      if (val.length > 0) {
        let tempArray = []
        this.operJsonList.forEach(element => {
          if (val.includes(element.ID)) {
            tempArray.push(...element.operArray)
          }
        });
        tempArray = [...new Set(tempArray)]
        this.showOperList = tempArray.map((item) => {
          return this.operMap.get(item)
        })
      } else {
        this.showOperList = []
      }
    },
    resourceCheckList(val) {
      if (val.length > 0) {
        Object.keys(NEW_JOB_TYPE_OBJ).forEach((key) => {
          NEW_JOB_TYPE_OBJ[key] = []
        })
        let specList = Object.assign({}, NEW_JOB_TYPE_OBJ)
        console.log("this.resourceJsonList", this.resourceJsonList)
        this.resourceJsonList.forEach(element => {
          if (val.includes(element.ID)) {
            element.taskTypes.forEach((item, index) => {
              if (item) {
                console.log(item, specList)
                specList[item].push(element.specIds[index])
              }
            })
          }
        });
        Object.keys(specList).forEach((key) => {
          specList[key] = [...new Set(specList[key])]
        })
        this.showSpecTableData = specList
        this.$nextTick(() => {
          this.$refs.resourceTable.tabPosition = NEW_JOB_TYPE[0].k
          this.$refs.resourceTable.changeTaskType(NEW_JOB_TYPE[0].k)
        })
      } else {
        this.showSpecTableData = {}
      }
    }
  },
  methods: {
    changeTaskType(type) {
      this.filterData = this.specificationList.filter((item) => {
        return this.showSpecTableData[type].includes(String(item.ID))
      })
    },
    showOpDetail(item) {
      const jsonData = JSON.parse(item.RightInfo)
      let operNameArray = []
      let operNumArray = []
      let opercomputeResourceArray = []
      jsonData.forEach(element => {
        operNameArray.push(element.operName)
        operNumArray.push(element.num)
        opercomputeResourceArray.push(element.computeResource)
      });
      this.opRoleDialogData = {
        ...item,
        operName: operNameArray,
        operNum: operNumArray,
        opercomputeResource: opercomputeResourceArray,
      }
      this.opRoleDialogShow = true

    },
    showStDetail(item) {
      const jsonData = JSON.parse(item.RightInfo)
      let codeSize = 0
      let outputSize = 0
      let stNameArray = []
      jsonData.forEach(element => {
        stNameArray.push(element.num)
        if (element.codeSize) codeSize = parseInt(element.codeSize) || 0
        if (element.outputSize) outputSize = parseInt(element.outputSize) || 0
      });
      this.stRoleDialogData = {
        ...item,
        operNum: stNameArray,
        codeSize: codeSize,
        outputSize: outputSize,
      }
      this.stRoleDialogShow = true
    },
    showReDetail(item) {
      const jsonData = JSON.parse(item.RightInfo)
      let uniqueArr
      if (jsonData.length) {
        uniqueArr = Array.from(new Set(jsonData.map(item => JSON.stringify(item)))).map(item => JSON.parse(item));
      }
      Object.keys(NEW_JOB_TYPE_OBJ).forEach((key) => {
        NEW_JOB_TYPE_OBJ[key] = []
      })
      let specList = Object.assign({}, NEW_JOB_TYPE_OBJ)
      uniqueArr.forEach(item => {
        specList[item.taskType].push(item.specId)
      });
      this.reRoleDialogData = specList
      this.reRoleDialogShow = true
      this.reRoleCommonData = {
        Name: item.Name,
        Description: item.Description,
        IsCommon: item.IsCommon,

      }
      this.$nextTick(() => {
        this.radioInit = jsonData.length && jsonData[0].taskType
        this.$refs.reDialogTable.changeTaskType(this.radioInit)
      })
    },
    closed() {
      this.radioInit = ''
    },
    cancel() {
      if (this.cancelUrl) {
        window.location.href = this.cancelUrl;
      } else {
        window.history.back();
      }
    },
    initRoleList() {
      return listAiforgeRole({}).then((res) => {
        const data = res.data
        this.operateList = data.filter((item) => { return item.Type === 0 })
        this.resourceList = data.filter((item) => { return item.Type === 1 })
        this.storageList = data.filter((item) => { return item.Type === 2 })
      }).catch((err) => {
        this.$message.error(err)
      })
    },
    initOperationList() {
      // const map = {}
      return listOperation().then((res) => {
        const data = res.data
        data.forEach((item) => {
          this.operMap.set(item.Name, item.Description)
        })
      }).catch((err) => {
        this.$message.error(err)
      })
    },
    getInitCheckedList() {
      const params = new URLSearchParams(location.search)
      if (params.has('opRole')) {
        let opRole = params.get('opRole')
        if (opRole) {
          this.operateCheckList = opRole.split(',').map((item) => Number(item))
          if (this.operateCheckList.length) {
            this.operateList.sort((a, b) => {
              if (this.operateCheckList.includes(a.ID)) {
                return -1
              } else {
                return 0
              }
            })
          }
        }
      }
      if (params.has('reRole')) {
        let reRole = params.get('reRole')
        this.resourceCheckList = reRole.split(',').map((item) => Number(item))
        if (this.resourceCheckList.length) {
          this.resourceList.sort((a, b) => {
            if (this.resourceCheckList.includes(a.ID)) {
              return -1
            } else {
              return 0
            }
          })
        }
      }
      if (params.has('stRole')) {
        let stRole = params.get('stRole')
        this.storageCheckList = stRole.split(',').map((item) => Number(item))
        if (this.storageCheckList.length) {
          this.storageList.sort((a, b) => {
            if (this.storageCheckList.includes(a.ID)) {
              return -1
            } else {
              return 0
            }
          })
        }
      }
    },
    initSpecificationList() {
      return getResSpecificationListAll({ available: 1, cluster: 'C2Net' }).then((res) => {
        res = res.data
        if (res.Code === 0) {
          const list = res.Data.Specs;
          const data = list.map((item) => {
            const NGPU = `${item.ComputeResource}:${item.AccCardsNum + '*' + getListValueWithKey(this.accCardTypeList, item.AccCardType)}`;
            const queueName = item.QueueName ? `【${item.QueueName}】` : '';
            const queueType = item.QueueType ? `【${item.QueueType}】` : '';
            return {
              ...item,
              SpecStr: `${NGPU}(${this.$t('resourcesManagement.gpuMem')}:${item.GPUMemGiB}GB), CPU:${item.CpuCores}, ${this.$t('resourcesManagement.mem')}:${item.MemGiB}GB`,
              QueueStr: `${item.QueueCode}${queueName}${queueType}`,
              NetworkTypeStr: `, ${this.$t('cloudbrainObj.networkType')}:${getListValueWithKey(this.networkTypeList, item.HasInternet)}`,
              visualizationStr: `, ${this.$t('cloudbrainObj.visualization')}:${item.EnableVisualization ? this.$t('resourcesManagement.enable') : this.$t('resourcesManagement.notEnable')}`,
            }
          });
          this.specificationList = data
          this.getInitCheckedList()
        } else {
          this.$message.error(res.Msg)
        }
      }).catch((err) => {
        this.$message.error(err)
      })
    },
    updateRole() {
      const roleIds = [
        ...this.operateCheckList,
        ...this.resourceCheckList,
        ...this.storageCheckList
      ].join(',')
      const data = {
        userIds: this.id,
        roleIds: roleIds
      }
      setAiforgeRoleToUser(data).then((res) => {
        if (res.data.code === '0') {
          this.$message.success('更新成功')
          setTimeout(() => {
            location.href = '/admin/access'
          }, 0)
        } else {
          this.$message.error(res.data.msg)
        }
      }).catch((err) => {
        this.$message.error(err || '更新失败')
      })
    },
    async initPageData() {
      try {
        this.loading = true
        await this.initRoleList()
        await this.initOperationList()
        await this.initSpecificationList()
      } catch (error) {
        this.$$message.error(error)
      } finally {
        this.loading = false
      }

    }
  },
  mounted() {
    let temp = location.pathname.split('/')
    this.id = temp[temp.length - 1]
    const params = new URLSearchParams(location.search)
    this.name = params.get('name')
    this.initPageData()

  },
  beforeMount() {
    const urlParams = getUrlSearchParams();
    if (urlParams.backurl) {
      this.cancelUrl = urlParams.backurl;
    }

  },
  beforeDestroy() {
  },
};
</script>

<style scoped lang="less">
.page-title {
  display: flex;
  height: 50px;

  span {
    font-weight: 700;
    font-size: 16px;
    color: rgb(16, 16, 16);
    line-height: 30px;
  }
}

.resource-title {
  margin-top: 30px;

  span {
    font-weight: 700;
    font-size: 14px;
    color: rgb(16, 16, 16);
    line-height: 30px;
  }
}

.btn-wrap {
  margin-top: 22px;
  height: 32px;
  width: 91px;
  border-radius: 5px;
  background-color: rgba(91, 185, 115, 1);
  text-align: center;
  cursor: pointer;

  span {
    line-height: 32px;
    color: #fff;
  }
}

.operate-wrap {
  display: flex;
  flex-direction: column;

  .operate-container {
    max-height: calc(100% - 40px);
    flex: 1;
    border: 1px solid rgba(232, 232, 232, 1);
    display: flex;

    .operate-list {
      width: 45%;
      margin-left: 100px;
      display: flex;
      flex-direction: column;

      .title {
        font-weight: 600;
        color: rgb(16, 16, 16);
        margin-top: 20px;
        margin-bottom: 12px;
      }

      .checkbox-wrap {
        flex: 1;
        overflow: hidden;
        overflow-y: auto;
        margin-bottom: 12px;

        // max-height: 200px;
        .check-wrap {
          display: flex;
          flex-direction: column;

          .el-checkbox {
            height: 34px;
            font-size: 14px;
          }

          .detail-a {
            color: #409EFF;
            margin-left: 6px;
            border-bottom: 1px solid #409EFF;
            vertical-align: super;
          }

          /deep/ .el-checkbox__input {
            vertical-align: baseline;
            overflow: hidden;
          }

          /deep/ .el-checkbox__label {
            width: 95%;
          }
        }

      }
    }

    .operate-collection {
      display: flex;
      flex-direction: column;
      flex: 1;
      margin-left: 12px;

      .title {
        font-weight: 600;
        color: rgb(16, 16, 16);
        margin-top: 20px;
        padding-bottom: 12px;
        padding-left: 20px;
        border-left: 1px solid #e8e8e8;
      }

      ul {
        margin-top: 0px;
        overflow: hidden;
        overflow-y: auto;
        margin-bottom: 12px;
        border-left: 1px solid #e8e8e8;

        li {
          color: #101010;
          line-height: 30px;
        }

        li::marker {
          color: #0191ff;
        }
      }
    }
  }
}

.sub-content-wrap {
  height: 40px;
  line-height: 40px;
  background-color: #eaeae7;

  .content {
    font-weight: 600;
    color: rgb(16, 16, 16);
    margin-left: 8px;
  }
}

.resource-wrap {
  display: flex;
  overflow: hidden;
  border: 1px solid rgba(232, 232, 232, 1);
  flex-direction: column;

  .resource-container {
    max-height: calc(100% - 40px);
    flex: 1;
    overflow-y: auto;

    .title {
      font-weight: 600;
      color: rgb(16, 16, 16);
      margin-top: 20px;
      margin-bottom: 12px;
      margin-left: 100px;
    }

    .check-wrap {
      display: flex;
      flex-wrap: wrap;
      margin-bottom: 8px;

      .el-checkbox {
        width: 50%;
        padding-left: 100px;
        margin-right: 0;
        height: 34px;
        font-size: 14px;

        .detail-a {
          color: #409EFF;
          margin-left: 6px;
          border-bottom: 1px solid #409EFF;
          vertical-align: super;
        }
      }

      /deep/ .el-checkbox__input {
        vertical-align: baseline;
        overflow: hidden;
      }

      /deep/ .el-checkbox__label {
        width: 95%;
      }
    }
  }
}

::-webkit-scrollbar-track {
  background: transparent;
  border-radius: 0;
}

::-webkit-scrollbar {
  -webkit-appearance: none;
  width: 0px;
  height: 4px;
}

::-webkit-scrollbar-thumb {
  cursor: pointer;
  border-radius: 5px;
  background: rgba(0, 0, 0, 0.15);
  transition: color 0.2s ease;
}
</style>
