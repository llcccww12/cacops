<template>
<div style="margin: 0 20px">
  <div class="title"><span>{{$t('userRole.OrgPermissionConfig')}}</span></div>
  <div style="margin-top: 15px;">
    <el-input :placeholder="$t('userRole.pleaseEnterContent')" size="medium" v-model="searchValue" class="input-with-select" @keydown.enter.stop.native.prevent="searchUser">
      <el-select v-model="select" slot="prepend">
        <el-option :label="$t('userRole.OrganizationName')" value="1"></el-option>
        <el-option :label="$t('userRole.organizationId')" value="2"></el-option>
      </el-select>
      <el-button type="primary" slot="append" @click="searchUser">{{$t('repos.search')}}</el-button>
    </el-input>
  </div>
  <div class="tools-bar">
    <div class="left">
      <!-- <el-select class="select" size="medium" filterable v-model="opRoleSelect" @change="changeOpSelect">
          <el-option v-for="item in opRoleSelectList" :key="item.k" :label="item.v" :value="item.k" />
      </el-select>
      <el-select class="select" size="medium" filterable v-model="reRoleSelect" @change="changeReSelect">
          <el-option v-for="item in reRoleSelectList" :key="item.k" :label="item.v" :value="item.k" />
      </el-select> -->
      <el-select class="select" size="medium" filterable v-model="stRoleSelect" @change="changeStSelect">
          <el-option v-for="item in stRoleSelectList" :key="item.k" :label="item.v" :value="item.k" />
      </el-select>
    </div>
    <div style="margin-left: auto">
      <!-- <el-button type="primary"  @click="batchAddRoles" size="medium" style="margin-right:12px">批量设置用户权限</el-button> -->
      <el-select class="select" size="medium" v-model="sortSelect" @change="changeSortSelect">
          <el-option v-for="item in sortSelectList" :key="item.k" :label="item.v" :value="item.k" />
      </el-select>
    </div>
  </div>
  <div class="table-container">
    <div style="width: 100%">
      <el-table border :data="filtedData" v-tableSticky style="width: 100%" v-loading="loading" stripe>
        <el-table-column prop="ID" fixed label="ID" width="80"></el-table-column>
        <el-table-column prop="Name" fixed :label="$t('userRole.OrganizationName')" min-width="200">
          <template slot-scope="scope">
            <a :href="'/'+ scope.row.Name">{{ scope.row.Name}}</a>
          </template>
        </el-table-column>
        <!-- <el-table-column prop="operRoles" :label="$t('userRole.opCategoryRole')" width="280">
          <template slot-scope="scope">
            <span>{{ scope.row.operRoles}}</span>
          </template>
        </el-table-column>
      
        <el-table-column prop="resourceRoles" :label="$t('userRole.reCategoryRole')"  min-width="320">
          <template slot-scope="scope">
            <span>{{ scope.row.resourceRoles}}</span>
          </template>
        </el-table-column> -->
        <el-table-column prop="storageRoles" :label="$t('userRole.stCategoryRole')"  min-width="140">
          <template slot-scope="scope">
            <span>{{ scope.row.storageRoles}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="CreateTime" :label="$t('modelManage.createTime')" align="center" width="260" />
        <el-table-column prop="UpdatedTime" :label="$t('modelManage.updateTime')" align="center" width="260" />
        <el-table-column :label="$t('edit')" align="center" width="60" fixed="right">
          <template slot-scope="scope">
            <span class="edit-btn" @click="openEdit(scope.row)"><i class="ri-edit-box-line"></i></span>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
  <div class="__r_p_pagination" v-if="filtedData.length">
    <div style="margin-top: 2rem">
      <div class="center">
        <el-pagination background @size-change="handleSizeChange" @current-change="currentChange" :current-page="params.page"
          :page-sizes="[10,15,20]" :page-size="params.pageSize"
          layout="total, sizes, prev, pager, next, jumper" :total="pageTotal">
        </el-pagination>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import { listRightUser,listAiforgeRole } from '~/apis/modules/resources';
import { listAiforgeOrgRole, listRightOrgUser } from '~/apis/modules/organization';
import { formatDate } from 'element-ui/lib/utils/date-util';
export default {
  data() {
    return {
      opRoleSelectList:[{k:'', v:this.$t('userRole.allOpCategory')}],
      opRoleSelect:'',
      reRoleSelectList:[{k:'', v:this.$t('userRole.allReCategory')}],
      reRoleSelect:'',
      stRoleSelectList:[{k:'', v:this.$t('userRole.allStCategory')}],
      stRoleSelect:'',
      sortSelectList:[
        {k:'', v:this.$t('userRole.permissionUpdateUnixDesc')},
        // {k:'updated_unix asc', v:this.$t('userRole.permissionUpdateUnixAsc')},
        // {k:'updated_unix desc', v:this.$t('userRole.permissionUpdateUnixDesc')},
        {k:'created_unix desc', v:this.$t('userRole.orgPermissionCreatedUnixDesc')},
        {k:'created_unix asc', v:this.$t('userRole.orgPermissionCreatedUnixAsc')},
      ],
      sortSelect:'',
      tableData:[],
      loading:false,
      params:{
        page:1,
        pageSize:15
      },
      pageTotal: 0,
      select: '1',
      searchValue: '',
    };
  },
  components: {},
  computed: {
    filtedData(){
      return this.tableData
    }
  },
  methods: {
    batchAddRoles(){
       window.location.href = '/admin/access/batch'
    },
    handleSizeChange(val){
      this.params.pageSize = val
      this.initRightUserList()
    },
    currentChange: function (val) {
      this.params.page = val
      this.initRightUserList()
    },
    changeOpSelect(val){
      this.params.operRoleId = val
      this.initRightUserList()
    },
    changeStSelect(val){
      this.params.storageRoleId = val
      this.initRightUserList()
    },
    changeReSelect(val){
      this.params.resourceRoleId = val
      this.initRightUserList()
    },
    changeSortSelect(val){
      this.params.orderBy = val
      this.initRightUserList()
    },
    searchUser(){
      if(this.select == '1'){
        this.params.userId = ''
        this.params.userName = this.searchValue
      }else{
        this.params.userName = ''
        this.params.userId = Number(this.searchValue)
      }
      this.initRightUserList()
    },
    openEdit(item){
      window.location.href =`/admin/org_access/${item.ID}?name=${item.Name}&opRole=${item.operRolesIds}&reRole=${item.resourceRolesIds}&stRole=${item.storageRolesIds}`
    },
    initRoleList(){
      listAiforgeOrgRole({}).then(res=>{
        const data = res.data
        data.forEach((item)=>{
          if(item.Type===0){
            this.opRoleSelectList.push({k:item.ID,v:item.Name})
          }else if(item.Type===1){
            this.reRoleSelectList.push({k:item.ID,v:item.Name})
          }else{
            this.stRoleSelectList.push({k:item.ID,v:item.Name})
          }
        })
      }).catch((error)=>{
        this.$message.error(error)
      })
    },
    initRightUserList(){
      this.loading = true
      listRightOrgUser(this.params).then((res)=>{
        if (!res?.data) {
          this.$message.warning('获取数据为空')
          return
        }
        const { count, data = [] } = res.data
        this.pageTotal = count
        this.tableData = data.map(item => this.processTableItem(item))
      }).catch((error)=>{
        const message = error?.message || '请求失败，请稍后重试'
        this.$message.error(message)
      }).finally(() => {
        this.loading = false
      })
    },
    processTableItem(item) {
    const processRoles = (roles, type) => {
      const validRoles = Array.isArray(roles) ? roles : []
        return {
          names: validRoles.map(role => role.Name).join('、'),
          ids: validRoles.map(role => role.ID).join(',')
        }
      }

      const oper = processRoles(item.operRole)
      const resource = processRoles(item.resourceRole)
      const storage = processRoles(item.storageRole)
      return {
        ID: item.id,
        Name: item.name,
        operRole: item.operRole,
        resourceRole: item.resourceRole,
        operRoles: oper.names,
        resourceRoles: resource.names,
        storageRoles: storage.names,
        operRolesIds: oper.ids,
        resourceRolesIds: resource.ids,
        storageRolesIds: storage.ids,
        CreateTime: this.formatUnixTime(item.createdUnix),
        UpdatedTime: this.formatUnixTime(item.updateUnix)
      }
    },
    formatUnixTime(timestamp) {
      return timestamp ? 
        formatDate(new Date(timestamp).getTime(), 'yyyy-MM-dd HH:mm:ss') : 
        '--'
    }
  },
  mounted() {
    this.initRoleList()
  },
  created() {
    this.initRightUserList()
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
.input-with-select{
  width: 50%;
  margin-bottom: 18px;
  .el-input-group__prepend {
    background-color: #fff;
  }
  
}
.tools-bar{
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}
.table-container {
  margin-bottom: 16px;

  /deep/ .el-table__header {
    th {
      background: #e8e8e8;;
      font-size: 12px;
      color: #101010;
    }
  }

  /deep/ .el-table__body {
    td {
      font-size: 12px;
    }
    tr:hover{
      td{
        background-color: #eff9ff !important;
      }
    }
  }
  /deep/ .el-table--border{
    td,th {
      border-right: none;
    }
  }
  
  .edit-btn {
    font-size: 16px;
    color: #0066ff;
    font-weight: bold;
    cursor: pointer;
  }
}
/deep/ .el-input-group__append {
  background-color: #0191ff;
  color: #fff;
  padding: 0 32px;
  border: 1px solid #0191ff;
} 
/deep/ .el-input-group__prepend{
  background-color: #fff;
  width: 120px;
}
</style>
