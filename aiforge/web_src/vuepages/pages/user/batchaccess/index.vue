<template>
<div style="margin: 0 20px;height: 100%;" v-loading="loading">
  <div class="page-title"><span>{{$t('userRole.batchSetOrCancel')}}</span></div>
  <div class="operate-wrap">
    <div style="height:100%">
      <div class="sub-content-wrap"><span class="content">{{$t('userRole.OperationalPermissions')}}</span></div>
      <div class="operate-container">
        <div class="operate-list">
          <div class="title">{{$t('userRole.opCategoryRole')}}：</div>
          <div class="checkbox-wrap">
            <el-checkbox-group v-model="operateCheckList" class="check-wrap" style="display:flex;flex-direction:column">
              <el-checkbox v-for="item in operateList" :key="item.ID" :label="item.ID">
                <span class="nowrap" style="display:inline-block;max-width:100%" :title="item.Name + '(' + item.Description + ')'">{{item.Name}}  ({{item.Description}})</span>
                <span class="detail-a" @click.stop.prevent="showOpDetail(item)">{{$t('userRole.detail')}}</span>
              </el-checkbox>
            </el-checkbox-group>
          </div>
        </div>
        <div class="operate-collection">
          <div class="title">{{$t('userRole.permissionList')}}：</div>
          <ul>
            <li v-for="item in showOperList" :key="item">{{item}}</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
  <div class="resource-wrap">
    <div class="sub-content-wrap"><span class="content">{{$t('userRole.CResourcePermissions')}}</span></div>
    <div class="resource-container">
      <div class="title">{{$t('userRole.reCategoryRole')}}：</div>
      <div>
        <el-checkbox-group v-model="resourceCheckList" class="check-wrap">
          <el-checkbox v-for="item in resourceList" :key="item.ID" :label="item.ID">
            <span class="nowrap" style="display:inline-block;max-width:90%" :title="item.Name + '(' + item.Description + ')'">{{item.Name}}   ({{item.Description}}) </span>
            <span class="detail-a" @click.stop.prevent="showReDetail(item)">{{$t('userRole.detail')}}</span>
          </el-checkbox>
        </el-checkbox-group>
      </div>
    </div>
  </div>
  <div class="resource-wrap">
    <div class="sub-content-wrap"><span class="content">{{$t('userRole.StoragePermissions')}}</span></div>
    <div class="resource-container">
      <div class="title">{{$t('userRole.stCategoryRole')}}：</div>
      <div>
        <el-checkbox-group v-model="storageCheckList" class="check-wrap">
          <el-checkbox v-for="item in storageList" :key="item.ID" :label="item.ID">
            <span class="nowrap" style="display:inline-block;max-width:90%" :title="item.Name + '(' + item.Description + ')'">{{item.Name}}   ({{item.Description}}) </span>
            <span class="detail-a" @click.stop.prevent="showStDetail(item)">{{$t('userRole.detail')}}</span>
          </el-checkbox>
        </el-checkbox-group>
      </div>
    </div>
  </div>
  <div class="resource-wrap">
    <div class="sub-content-wrap"><span class="content">{{$t('userRole.enterUserId')}}</span></div>
    <div class="resource-container">
      <div class="title">{{$t('userRole.userId')}}：</div>
      <div class="user-info">
        <div style="flex:1;max-width:50%">
          <el-input v-model="userInfo" :rows="10" type="textarea" @input="handleInput" :placeholder="`请输入用户ID，一个用户一行，如：
42
31
……
`"></el-input>
        </div>
        <div style="flex:1;" v-if="ErrorList.length!=0">
          <div>
            <span class="title_e">{{$t('userRole.failed')}} {{ErrorList.length}} {{$t('userRole.times')}}</span>
            <span>{{$t('userRole.failedUserList')}}：</span>
          </div>
          <ul>
            <li v-for="(item,index) in ErrorList" :key="item.id">
              {{ item.id }}  <span style="color:#888;margin-left: 8px;">{{item.reason}}</span>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
  <div style="display:flex;margin-top:1rem;">
    <el-button type="primary"  @click="batchAddAccess" size="medium">增加权限</el-button>
    <el-button type="warning"  @click="batchDelAccess" size="medium" style="margin-left:1rem;">取消权限</el-button>
    <!-- <div class="btn-wrap" type="primary" @click="batchAddAccess"><span>增加权限</span></div>
    <div class="btn-wrap" type="warning" @click="batchDelAccess" style="margin-left:1rem;"><span>取消权限</span></div> -->
  </div>
  <opRoleDialog  :visible.sync="opRoleDialogShow" type="view" :data="opRoleDialogData"></opRoleDialog>
  <BaseDialog :visible.sync="reRoleDialogShow" :title="$t('userRole.viewReRole')" width="80%" @closed="closed">
    <resourceTableCopy ref="reDialogTable" :radioInit="radioInit" :showSpecTableData="reRoleDialogData" :specificationList="specificationList" :reRoleCommonData="reRoleCommonData"></resourceTableCopy>
  </BaseDialog>
  <storageRoleDialog  :visible.sync="stRoleDialogShow" type="view" :data="stRoleDialogData"></storageRoleDialog>
</div>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';
import opRoleDialog from '../role/opRoleDialog.vue';
import resourceTableCopy from '../accessedit/resourceTable.vue'
import storageRoleDialog from '../role/storageRoleDialog.vue';
import { listAiforgeRole, listOperation, getResSpecificationListAll, bacthAddRoleToUser, bacthDelRoleToUser } from '~/apis/modules/resources';
import { getListValueWithKey } from '~/utils';
import { ACC_CARD_TYPE, NETWORK_TYPE_VALUE, NEW_JOB_TYPE_OBJ } from '~/const';
export default {
  data() {
    return {
      loading:false,
      operateCheckList:[],
      operateList:[],
      showOperList:[],
      operMap: new Map(),
      resourceCheckList:[],
      resourceList:[],
      storageCheckList:[],
      storageList:[],
      userInfo:'',
      ErrorList:[],
      opRoleDialogShow:false, // 操作权限弹窗
      opRoleDialogData:{},
      reRoleDialogShow:false, // 资源权限弹窗
      reRoleDialogData:{},
      radioInit:'',
      specificationList:[],
      reRoleCommonData:{},
      accCardTypeList:[...ACC_CARD_TYPE],
      networkTypeList:[...NETWORK_TYPE_VALUE],
      stRoleDialogShow:false, // 操作权限弹窗
      stRoleDialogData:{},
    };
  },
  components: { BaseDialog, opRoleDialog, resourceTableCopy, storageRoleDialog },
  computed: {
    operJsonList(){
     return this.operateList.map((item)=>{
      const data = JSON.parse(item.RightInfo)
      const operArray = data.map((item)=>item.operName)
       return {
         ID: item.ID,
         operArray:operArray
       }
     })
   }
  },
  watch: {
    operateCheckList(val){
      if(val.length > 0){
        let tempArray = []
        this.operJsonList.forEach(element => {
          if(val.includes(element.ID)){
            tempArray.push(...element.operArray)
          }
        });
        tempArray = [...new Set(tempArray)]
        this.showOperList = tempArray.map((item)=>{
          return this.operMap.get(item)
        })
      }else{
        this.showOperList = []
      }
    },
  },
  methods: {
    handleInput(value){
      this.userInfo = value.replace(/[^\d\n]/g, '');
    },
    batchAddAccess(){
      this.batchRequest(true)
    },
    batchDelAccess(){
      this.batchRequest(false)
    },
    batchRequest(flag){
      const roleIds = [
        ...this.operateCheckList, 
        ...this.resourceCheckList,
        ...this.storageCheckList
      ].join(',')
       const userIds = this.userInfo.replace(/\s+/g, ',');

      console.log(userIds)
      const data = {
        userIds: userIds,
        roleIds: roleIds
      }
      let batchRequest = flag ? bacthAddRoleToUser : bacthDelRoleToUser
      let successTips = flag ? this.$t('userRole.addAcessSuccess') : this.$t('userRole.cancelAcessSuccess')
      console.log(data)
      batchRequest(data).then((res)=>{
        if(res.data.code==='0'){
          this.$message.success(successTips)
          this.ErrorList = []
          setTimeout(()=>{
            location.href = '/admin/access'
          },0)
        }else{
          const failedObject = res.data.failedUsers
          this.ErrorList = Object.entries(failedObject).map(([id, reason]) => ({ id: Number(id), reason }));
        }
      }).catch((err)=>{
        this.$message.error(err || '更新失败')
      })
    },
    closed(){
      this.radioInit = ''
    },
    showOpDetail(item){
      const jsonData = JSON.parse(item.RightInfo)
      let operNameArray = []
      jsonData.forEach(element => {
        operNameArray.push(element.operName)
      });
      this.opRoleDialogData = {
        ...item,
        operName: operNameArray,
      }
      this.opRoleDialogShow = true
    },
    initOperationList(){
      // const map = {}
      return listOperation().then((res)=>{
        const data = res.data
        data.forEach((item)=>{
          this.operMap.set(item.Name,item.Description)
        })
      }).catch((err)=>{
        this.$message.error(err)
      })
    },
    initRoleList(){
      return listAiforgeRole({}).then((res)=>{
        const data = res.data
        this.operateList = data.filter((item)=>{return item.Type === 0})
        this.resourceList = data.filter((item)=>{return item.Type === 1})
        this.storageList = data.filter((item)=>{return item.Type === 2})
      }).catch((err)=>{
        this.$message.error(err)
      })
    },

    showReDetail(item){
      const jsonData = JSON.parse(item.RightInfo)
      let uniqueArr
      if (jsonData.length){
         uniqueArr = Array.from(new Set(jsonData.map(item => JSON.stringify(item)))).map(item => JSON.parse(item));
      }
      Object.keys(NEW_JOB_TYPE_OBJ).forEach((key)=>{
          NEW_JOB_TYPE_OBJ[key] = []
        })
      let specList = Object.assign({},NEW_JOB_TYPE_OBJ)
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
      this.$nextTick(()=>{
        this.radioInit = jsonData.length && jsonData[0].taskType
        this.$refs.reDialogTable.changeTaskType(this.radioInit)
      })
    },
    initSpecificationList(){
      return getResSpecificationListAll({available: 1,cluster:'C2Net'}).then((res)=>{
        res = res.data
        if (res.Code === 0) {
            const list = res.Data.Specs;
            const data = list.map((item) => {
              const NGPU = `${item.ComputeResource}:${item.AccCardsNum + '*' + getListValueWithKey(this.accCardTypeList, item.AccCardType)}`;
              const queueName = item.QueueName ? `【${item.QueueName}】` : '';
              const queueType = item.QueueType ? `【${item.QueueType}】` : '';
              return {
                ...item,
                SpecStr:`${NGPU}(${this.$t('resourcesManagement.gpuMem')}:${item.GPUMemGiB}GB), CPU:${item.CpuCores}, ${this.$t('resourcesManagement.mem')}:${item.MemGiB}GB`,
                QueueStr: `${item.QueueCode}${queueName}${queueType}`,
                NetworkTypeStr: `, ${this.$t('cloudbrainObj.networkType')}:${getListValueWithKey(this.networkTypeList, item.HasInternet)}`,
                visualizationStr: `, ${this.$t('cloudbrainObj.visualization')}:${item.EnableVisualization ? this.$t('resourcesManagement.enable') : this.$t('resourcesManagement.notEnable')}`,
              }
            });
            this.specificationList = data
            // this.getInitCheckedList()
        }else{
          this.$message.error(res.Msg)
        }
      }).catch((err)=>{
        this.$message.error(err)
      })
    },
    showStDetail(item){
      const jsonData = JSON.parse(item.RightInfo)
      let stNameArray = []
      jsonData.forEach(element => {
        stNameArray.push(element.num)
      });
      this.stRoleDialogData = {
        ...item,
        operNum: stNameArray,
      }
      this.stRoleDialogShow = true
    },
    async initPageData(){
      try {
        this.loading = true
        await this.initRoleList()
        await this.initOperationList()
        await this.initSpecificationList()    
      } catch (error) {
        this.$message.error(error)
      }finally{
        this.loading = false
      }
      
    }
  },
  mounted() {
    this.initPageData()
  },
  created() {

  },

  beforeDestroy() {
  },
};
</script>

<style scoped lang="less">
.page-title {
  display: flex;
  height:50px;
  span {
    font-weight: 700;
    font-size: 16px;
    color: rgb(16, 16, 16);
    line-height: 30px;
  }
}
.resource-title{
  margin-top:30px;
  span {
    font-weight: 700;
    font-size: 14px;
    color: rgb(16, 16, 16);
    line-height: 30px;
  }
}
.btn-wrap{
  margin-top: 22px;
  height: 32px;
  width: 91px;
  border-radius: 5px;
  background-color: rgba(91,185,115,1);
  text-align: center;
  cursor: pointer;
  span{
    line-height: 32px;
    color:#fff;
  }
}
.sub-content-wrap{
  height:40px;
  line-height: 40px;
  background-color: #eaeae7;
  .content{
    font-weight: 600;
    color: rgb(16, 16, 16);
    margin-left: 8px;
  }
}
.operate-wrap{
  display: flex;
  flex-direction: column;
  .operate-container{
    max-height: calc(100% - 40px);
    flex: 1;
    border: 1px solid rgba(232,232,232,1);
    display: flex;
    .operate-list{
      width: 45%;
      margin-left: 100px;
      display: flex;
      flex-direction: column;
      .title{
        font-weight: 600;
        color: rgb(16, 16, 16);
        margin-top: 20px;
        margin-bottom: 12px;
      }
      
      .checkbox-wrap{
        flex: 1;
        overflow: hidden;
        overflow-y: auto;
        margin-bottom: 12px;
        // max-height: 200px;
        .check-wrap{
          display: flex;
          flex-direction: column;
          .el-checkbox{
            height:34px;
            font-size:14px;
          }
          .detail-a{
            color: #409EFF;
            margin-left: 6px;
            border-bottom: 1px solid #409EFF;
            vertical-align: super;
          }
          /deep/ .el-checkbox__input{
            vertical-align: baseline;
            overflow: hidden;
          }

          /deep/ .el-checkbox__label{
            width:95%;
          }
        }
        
      }
    }
    .operate-collection{
      display: flex;
      flex-direction: column;
      flex: 1;
      margin-left:12px;
      .title{
        font-weight: 600;
        color: rgb(16, 16, 16);
        margin-top: 20px;
        padding-bottom: 12px;
        padding-left: 20px;
        border-left: 1px solid #e8e8e8;
      }
      ul{
        margin-top: 0px;
        overflow: hidden;
        overflow-y: auto;
        margin-bottom: 12px;
        border-left: 1px solid #e8e8e8;
        li{
          color: #101010;
          line-height: 30px;
        }
        li::marker{
          color: #0191ff;
        }
      }
    }
  }
}
.resource-wrap{
  display: flex;
  overflow: hidden;
  border: 1px solid rgba(232,232,232,1);
  flex-direction: column;
  .resource-container{
    max-height: calc(100% - 40px);
    flex: 1;
    overflow-y: auto;
    .title{
      font-weight: 600;
      color: rgb(16, 16, 16);
      margin-top: 20px;
      margin-bottom: 12px;
      margin-left: 100px;
    }
    .check-wrap{
      display: flex;
      flex-wrap: wrap;
      margin-bottom: 8px;
      .el-checkbox{
        width: 50%;
        padding-left: 100px;
        margin-right: 0;
        height:34px;
        font-size:14px;
        .detail-a{
          color: #409EFF;
          margin-left: 6px;
          border-bottom: 1px solid #409EFF;
          vertical-align: super;
        }
      }
      /deep/ .el-checkbox__input{
        vertical-align: baseline;
        overflow: hidden;
      }

      /deep/ .el-checkbox__label{
        width:95%;
      }
    }
    .user-info{
      display: flex;
      gap: 24px;
      margin: 12px 100px;
      max-height: 208px;
      overflow-y: auto;
      .title_e {
        color: #f56c6c;
        margin: 0 20px;
      }
      ul {
        max-height: 166px;
        overflow: auto;
        margin-top: 10px;
        color: red
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
