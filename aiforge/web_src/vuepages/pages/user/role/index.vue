<template>
  <div style="margin: 0 20px">
    <div class="title"><span>{{ $t('userRole.userRoleManagement') }}</span></div>
    <div style="margin-top: 15px;">
      <el-input :placeholder="$t('userRole.pleaseEnterRoleContent')" size="medium" v-model="searchValue"
        class="input-with-select" @keydown.enter.stop.native.prevent="searchUser">
        <el-button type="primary" slot="append" @click="searchUser">{{ $t('repos.search') }}</el-button>
      </el-input>
    </div>
    <div class="tools-bar">
      <div class="left">
        <el-select class="select" size="medium" v-model="rolePlay" @change="change">
          <el-option v-for="item in roleList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
        <el-select class="select" size="medium" v-model="isCommon" @change="change">
          <el-option v-for="item in commonList" :key="item.k" :label="item.v" :value="item.k" />
        </el-select>
      </div>
      <div class="tools-bar-btn-c">
        <el-button size="medium" type="primary" icon="el-icon-plus" @click="showOperRoleDialog('add')">
          {{ $t('userRole.newOpRole') }}</el-button>
        <el-button type="primary" icon="el-icon-plus" size="medium" @click="showResourceRoleDialog('add')">
          {{ $t('userRole.newReRole') }}</el-button>
        <el-button type="primary" icon="el-icon-plus" size="medium" @click="showStorageRoleDialog('add')">
          {{ $t('userRole.newStorageRole') }}</el-button>
      </div>
    </div>
    <div class="table-container">
      <div>
        <el-table border :data="filtedData.slice((currentPage - 1) * pageSize, currentPage * pageSize)"
          style="width: 100%" v-loading="loading" stripe v-tableSticky >
          <el-table-column prop="ID" fixed label="ID" align="center" header-align="center" width="80"></el-table-column>
          <el-table-column prop="Name" fixed :label="$t('userRole.roleName')" width="280" />
          <el-table-column prop="Type" :label="$t('userRole.roleType')" align="center" width="180">
            <template slot-scope="scope">
              <span>{{ getTypeName(scope.row.Type) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="Description" :label="$t('userRole.roleDescription')" align="center" width="320" />
          <el-table-column prop="IsCommon" :label="$t('userRole.isItDefault')" align="center" width="120">
            <template slot-scope="scope">
              <span>{{ scope.row.IsCommon ? $t('userRole.notDefault') : $t('userRole.default') }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="CreateTime" :label="$t('modelManage.createTime')" align="center" width="200" />
          <el-table-column prop="UpdatedTime" :label="$t('modelManage.updateTime')" align="center" width="200" />
          <el-table-column prop="CreatedUserId" :label="$t('modelManage.creator')" align="center" width="140" />
          <el-table-column :label="$t('operation')" align="right" min-width="200" fixed="right">
            <template slot-scope="scope">
              <span class="op-btn" @click="editDialog('view', scope.row)">{{ $t('userRole.viewDetail') }}</span>
              <span class="op-btn" @click="editDialog('edit', scope.row)">{{ $t('modelManage.edit') }}</span>
              <span class="op-btn" style="color:red" @click="deleteRow(scope.row)">{{ $t('delete') }}</span>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
    <div class="__r_p_pagination" v-if="filtedData.length">
      <div style="margin-top: 2rem">
        <div class="center">
          <el-pagination background @size-change="handleSizeChange" @current-change="currentChange"
            :current-page="currentPage" :page-sizes="pageSizes" :page-size="pageSize"
            layout="total, sizes, prev, pager, next, jumper" :total="pageTotal">
          </el-pagination>
        </div>
      </div>
    </div>
    <opRoleDialog @refresh="initRoleList({})" :visible.sync="opRoleDialogShow" :type="opRoleDialogType"
      :data="opRoleDialogData"></opRoleDialog>
    <resourceRoleDialog @refresh="initRoleList({})" :visible.sync="resourceRoleDialogShow"
      :type="resourceRoleDialogType" :data="resourceRoleDialogData"></resourceRoleDialog>
    <storageRoleDialog @refresh="initRoleList({})" :visible.sync="storageRoleDialogShow" :type="storageRoleDialogType"
      :data="storageRoleDialogData"></storageRoleDialog>
  </div>
</template>

<script>
import opRoleDialog from './opRoleDialog.vue';
import resourceRoleDialog from './resourceRoleDialog.vue';
import storageRoleDialog from './storageRoleDialog.vue';
import { listAiforgeRole, delAiforgeRole, listRightUser } from '~/apis/modules/resources'; import { formatDate } from 'element-ui/lib/utils/date-util';
export default {
  data() {
    return {
      roleList: [
        { k: '', v: this.$t('userRole.allRoleType') },
        { k: 0, v: this.$t('userRole.opCategory') },
        { k: 1, v: this.$t('userRole.reCategory') },
        { k: 2, v: this.$t('userRole.stCategoryRole') },
      ],
      rolePlay: '',
      commonList: [{ k: '', v: this.$t('userRole.isItDefault') }, { k: 1, v: this.$t('userRole.notDefault') }, { k: 0, v: this.$t('userRole.default') }],
      isCommon: '',
      opRoleDialogShow: false,
      opRoleDialogType: 'add',
      opRoleDialogData: {},
      resourceRoleDialogShow: false,
      resourceRoleDialogType: 'add',
      resourceRoleDialogData: {},
      storageRoleDialogShow: false,
      storageRoleDialogType: 'add',
      storageRoleDialogData: {},
      loading: false,
      tableData: [],
      currentPage: 1, //当前页
      pageTotal: 0, //总条数
      pageSize: 15, //当前页容量
      pageSizes: [10, 15, 20],
      searchValue: '',
    };
  },
  components: { opRoleDialog, resourceRoleDialog, storageRoleDialog },
  computed: {
    filtedData() {
      return this.tableData.filter((item) => {
        return this.rolePlay === item.Type || this.rolePlay === ''
      }).filter((item) => {
        return this.isCommon === item.IsCommon || this.isCommon === ''
      })
    }
  },
  watch: {
    filtedData(val) {
      if (val.length) {
        this.pageTotal = val.length
        // this.currentPage = 1
      }
    },
    searchValue(val) {
      if (!val) {
        this.initRoleList({})
      }
    }
  },
  methods: {
    change() {
      this.currentPage = 1
    },
    handleSizeChange(val) {
      this.currentPage = 1
      this.pageSize = val
      // this.initRoleList()
    },
    currentChange(val) {
      this.currentPage = val
      // this.initRoleList()
    },
    showOperRoleDialog(type) {
      this.opRoleDialogShow = true
      this.opRoleDialogType = type
    },
    showResourceRoleDialog(type) {
      this.resourceRoleDialogShow = true
      this.resourceRoleDialogType = type
    },
    showStorageRoleDialog(type) {
      this.storageRoleDialogShow = true
      this.storageRoleDialogType = type
    },
    searchUser() {
      this.initRoleList({ name: this.searchValue })
    },
    initRoleList(params) {
      this.loading = true
      listAiforgeRole(params).then(res => {
        this.loading = false
        const data = res.data
        this.pageTotal = data.length || 0
        this.tableData = data.map((item, index) => {
          const jsonData = JSON.parse(item.RightInfo)
          console.log("jsonData", jsonData)
          let operNameArray = []
          let operNumArray = []
          let opercomputeResourceArray = []
          let taskTypeArray = []
          let specIdArray = []
          let codeSize = 0
          let outputSize = 0
          jsonData.forEach((element) => {
            operNameArray.push(element.operName)
            operNumArray.push(element.num)
            opercomputeResourceArray.push(element.computeResource)
            let hasRepeatTaskType = taskTypeArray.indexOf(element.taskType)
            let hasRepeatSpecId = specIdArray.indexOf(element.specId)
            if (hasRepeatTaskType > -1 && hasRepeatTaskType > -1 && hasRepeatTaskType === hasRepeatSpecId) {
              return
            } else {
              taskTypeArray.push(element.taskType)
              specIdArray.push(element.specId)
            }
            if (element.codeSize) codeSize = parseInt(element.codeSize) || 0
            if (element.outputSize) outputSize = parseInt(element.outputSize) || 0
          });
          console.log("item", item)
          return {
            CreateTime: formatDate(new Date(item.CreatedUnix * 1000), 'yyyy-MM-dd HH:mm:ss'),
            UpdatedTime: item.UpdatedUnix ? formatDate(new Date(item.UpdatedUnix * 1000), 'yyyy-MM-dd HH:mm:ss') : '--',
            Type: item.Type,
            Name: item.Name,
            ID: item.ID,
            IsCommon: item.IsCommon,
            CreatedUserId: item.userName,
            TaskType: '--',
            Description: item.Description,
            operName: operNameArray,
            operNum: operNumArray,
            opercomputeResource: opercomputeResourceArray,
            taskType: taskTypeArray,
            specId: specIdArray,
            codeSize: codeSize,
            outputSize: outputSize,
          }
        })
      }).catch((error) => {
        this.$message.error(error)
        this.loading = false
      })
    },
    deleteRow(item) {
      let tips
      if (item.IsCommon) {
        tips = this.$t('userRole.deleteRoleConfirm1', { name: item.Name })
      } else {
        tips = `<span style="color:red">${this.$t('userRole.deleteRoleConfirm')}</span>` + this.$t('userRole.deleteRoleConfirm1', { name: item.Name })
      }
      this.$confirm(tips, this.$t('tips'), {
        confirmButtonText: this.$t('confirm1'),
        cancelButtonText: this.$t('cancel'),
        dangerouslyUseHTMLString: true,
        type: 'warning',
        lockScroll: false,
        closeOnClickModal: false,
        showClose: false,
        closeOnPressEscape: false,
      }).then(() => {
        delAiforgeRole({ id: item.ID }).then((res) => {
          if (res.data.code === '0') {
            this.$message.success(this.$t('imagesObj.deleteSuccessTips'))
            this.initRoleList({})
          } else {
            this.$message.error(res.data.msg)
          }
        }).catch((error) => {
          this.$message.error(error)
        })
      }).catch((error) => {

      })
    },
    editDialog(type, item) {
      const dialogConfig = {
        0: {
          show: 'opRoleDialogShow',
          type: 'opRoleDialogType',
          data: 'opRoleDialogData',
        },
        1: {
          show: 'resourceRoleDialogShow',
          type: 'resourceRoleDialogType',
          data: 'resourceRoleDialogData',
        },
        default: {
          show: 'storageRoleDialogShow',
          type: 'storageRoleDialogType',
          data: 'storageRoleDialogData',
        },
      };
      const config = dialogConfig[item.Type] || dialogConfig.default;

      this[config.show] = true;
      this[config.type] = type;
      this[config.data] = item;
    },

    getTypeName(val) {
      const labels = {
        0: this.$t('userRole.opCategory'),
        1: this.$t('userRole.reCategory'),
        2: this.$t('userRole.storageCategory'),
      };
      return labels[val] ?? '--';
    }
  },
  mounted() {
    this.initRoleList({})
  },
  beforeDestroy() {
  },
};
</script>

<style scoped lang="less">
.title {
  height: 30px;
  display: flex;
  align-items: center;
  margin-bottom: 20px;

  span {
    font-weight: 700;
    font-size: 16px;
    color: rgb(16, 16, 16);
  }
}

.tools-bar {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.input-with-select {
  width: 50%;
  margin-bottom: 18px;

  .el-input-group__prepend {
    background-color: #fff;
  }

}

/deep/ .el-input-group__append {
  background-color: #0191ff;
  color: #fff;
  padding: 0 32px;
  border: 1px solid #0191ff;
}

.table-container {
  margin-bottom: 16px;

  /deep/ .el-table__header {
    th {
      background: #fafafa;
      ;
      font-size: 12px;
      color: rgb(36, 36, 36);
    }
  }

  /deep/ .el-table__body {
    td {
      font-size: 12px;
    }

    tr:hover {
      td {
        background-color: #eff9ff !important;
      }
    }
  }

  /deep/ .el-table--border {
    td,
    th {
      border-right: 2px solid #fafafa;
      ;
    }

  }

  .op-btn {
    cursor: pointer;
    font-size: 12px;
    color: rgb(25, 103, 252);
    margin: 0 5px;
  }
}

.center {
  display: flex;
  justify-content: center;
}
</style>
