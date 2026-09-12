<template>
  <div class="base-dlg">
    <BaseDialog :visible.sync="dialogShow" width="60%"
      :title="dialogTitle"
      @open="open" @opened="opened" @close="close" @closed="closed">
      <div class="dlg-content">
        <div class="form">
          <div class="form-row">
            <label class="required" for="">{{$t('userRole.roleType')}}：</label>
            <div class="content">
              <span style="color:#101010">{{$t('userRole.opCategory')}}</span>
            </div>
          </div>
          <div class="form-row">
            <label class="required" for="">{{$t('userRole.roleName')}}：</label>
            <div class="content">
              <el-input v-model="params.name" :disabled="type==='view'" style="width:60%;"
                maxlength="80" :placeholder="$t('resourcesManagement.roleNameTips')">
              </el-input>
            </div>
          </div>
          <div class="form-row">
            <label for="">{{$t('userRole.roleDescription')}}：</label>
            <div class="content">
              <el-input type="textarea" :rows="2" v-model="params.description"  :disabled="type==='view'" 
              style="width:60%;"  maxlength="800" :placeholder="$t('resourcesManagement.roleDescTips')">
              </el-input>
            </div>
          </div>
          <div class="form-row border">
            <div class="title" style="position:absolute">{{$t('userRole.roleSelected')}} </div>
            <div class="content box">
              <el-checkbox :disabled="type==='view'" :indeterminate="isIndeterminate" v-model="checkAll" @change="handleCheckAllChange">{{$t('userRole.roleSelectedTips')}}</el-checkbox>
                <el-checkbox-group :disabled="type==='view'" v-model="checkedRoleOps" @change="handleCheckedChange">
                  <el-checkbox v-for="item in roleOptios" :label="item.Name" :key="item.Name">
                    <el-row v-if="item.Name === 'multi_node'">
                      {{$t('userRole.operaNodeTips1')}}&nbsp;&nbsp;<el-input v-model="numNode" size="mini" style="width:46px"></el-input>&nbsp;&nbsp;{{$t('userRole.operaNodeTips2')}}&nbsp;&nbsp;
                      <el-select v-model="computeResourceNode" size="mini" placeholder="请选择" style="width:140px">
                        <el-option v-for="item in computingTypeList" :key="item.k" :label="item.v" :value="item.k" />
                      </el-select>&nbsp;&nbsp;{{$t('userRole.operaNodeTips3')}}
                    </el-row>
                    <el-row v-else-if="item.Name === 'multi_task'">
                      {{$t('userRole.operaTaskTips1')}}&nbsp;&nbsp;<el-input v-model="numTask" size="mini" style="width:66px"></el-input>&nbsp;&nbsp;{{$t('userRole.operaTaskTips2')}}&nbsp;&nbsp;
                      <el-select v-model="jobTypeItem" size="mini" placeholder="请选择" style="width:140px">
                        <el-option v-for="item in jobTypeList" :key="item.k" :label="item.v" :value="item.k" />
                      </el-select>
                    </el-row>
                    <span v-else>{{item.Description}}</span>
                  </el-checkbox>
              </el-checkbox-group>
            </div>
          </div>
          <div style="margin-left:80px" v-if="type!=='view'">
            <el-button type="primary" :disabled="type==='view'" class="btn confirm-btn" @click="confirm">{{ $t('confirm') }}</el-button>
            <el-button class="btn" @click="cancel">{{ $t('cancel') }}</el-button>
          </div>
        </div>
      </div>
    </BaseDialog>
  </div>
</template>

<script>
import BaseDialog from '~/components/BaseDialog.vue';
import { listOperation, addAiforgeRole, updateAiforgeRole } from '~/apis/modules/resources';
import { JOB_TYPE } from "~/const/index.js";
export default {
  name: "opRoleDialog",
  props: {
    visible: { type: Boolean, default: false },
    type: { type: String, defalut: 'add' },
    data: { type: Object, default: () => ({}) },
  },
  components: {
    BaseDialog
  },
  data() {
    return {
      dialogShow: false,
      checkAll:false,
      isIndeterminate:false,
      roleOptios:[],
      checkedRoleOps:[],
      params:{
        name:"",
        type:0,
        isCommon:1,
        description:'',
        operName:'',
        operNum: '',
        computeResource: '',
      },
      numNode: 2,
      computeResourceNode: 'NPU',
      computingTypeList: [{ k: 'GPU', v: 'GPU' }, { k: 'NPU', v: 'NPU' }],
      jobTypeList: JOB_TYPE,
      numTask: 1,
      jobTypeItem: JOB_TYPE[0].k,
      getOperationFlag:false,
    };
  },
  watch: {
    visible: function (val) {
      this.dialogShow = val;
    },
  },
  computed: {
    dialogTitle(){
      switch(this.type){
        case 'add':
          return this.$t('userRole.newOpRole')
        case 'edit':
          return this.$t('userRole.editOpRole')
        case 'view':
          return this.$t('userRole.viewOpRole')
      }
    },
  },
  methods: {
    initListOperation(){
      console.log("this.getOperationFlag",this.getOperationFlag)
      console.log("this.type",this.type)
      if(this.getOperationFlag){
          if (this.type === 'add'){
            this.restParams()
          }else{
            this.restorePamras()
          }
      }else{
        listOperation().then((res)=>{
          this.roleOptios = res.data
          this.getOperationFlag = true
          if (this.type === 'add') {
            //
          } else {
            this.restorePamras()
          }
        })
      }
    },
    open() {
      this.initListOperation()
      this.$emit("open");
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
    handleCheckedChange(val){
        let checkedCount = val.length;
        this.checkAll = checkedCount === this.roleOptios.length;
        this.isIndeterminate = checkedCount > 0 && checkedCount < this.roleOptios.length;
    },
    handleCheckAllChange(val){
        this.checkedRoleOps = val ? this.roleOptios.reduce((pre, cur) =>{
            return pre.concat(cur.Name);
        },[]):[];
        this.isIndeterminate = false;
    },
    restParams(){
      this.params = {name:"",type:0,isCommon:1,description:'',operName:'',operNum: '',computeResource: '',jobType: ''}
      this.numNode = 2
      this.computeResourceNode = 'NPU',
      this.numTask = 1,
      this.jobTypeItem = JOB_TYPE[0].k,
      this.checkedRoleOps = []
      this.checkAll = false
      this.isIndeterminate = false
    },
    restorePamras(){
      console.log(this.data)
      this.params.name = this.data.Name
      this.params.description = this.data.Description
      this.checkedRoleOps = this.data.operName
      let multiNodeIndex = this.data.operName.findIndex(item=>item==='multi_node')
      this.numNode = this.data.operNum[multiNodeIndex]
      this.computeResourceNode = this.data.opercomputeResource[multiNodeIndex]
      let multiTaskIndex = this.data.operName.findIndex(item=>item==='multi_task')
      if (multiTaskIndex >= 0) {
        this.numTask = this.data.operNum[multiTaskIndex] || 1;
        this.jobTypeItem = this.data.taskType[multiTaskIndex] || JOB_TYPE[0].k;
      } else {
        this.numTask = 1;
        this.jobTypeItem = JOB_TYPE[0].k;
      }
      this.handleCheckedChange(this.checkedRoleOps)
    },
    confirm() {
      if(this.params.name === ''){
        this.$message.error("请输入角色名称")
        return
      }
      if(this.checkedRoleOps.length === 0){
        this.$message.error("请勾选角色拥有的权限!")
        return
      }
      this.params.operName = this.checkedRoleOps.join(',')
      let operNumTem = this.checkedRoleOps.map((item)=>{
        if(item === 'multi_node'){
          return this.numNode
        }else if(item === 'multi_task'){
          return this.numTask
        }else{
          return ''
        }
      })
      this.params.operNum = operNumTem.join(',')
      let computeResourceNodeTem = this.checkedRoleOps.map((item)=>{
        if(item==='multi_node'){
          return this.computeResourceNode
        }else{
          return ''
        }
      })
      let taskTypeTem = this.checkedRoleOps.map((item)=>{
        if(item==='multi_task'){
          return this.jobTypeItem
        }else{
          return ''
        }
      })
      this.params.computeResource = computeResourceNodeTem.join(',')
      this.params.jobType = taskTypeTem.join(',')
      if(this.type === 'add'){
        addAiforgeRole(this.params).then((res)=>{
          if(res.data.code == 0){
            this.$message.success(this.$t('submittedSuccessfully'))
            this.dialogShow = false
            this.$emit("update:visible", false);
            this.$emit("refresh");
          }else{
            this.$message.error(res.data.msg)
          }
        }).catch((err)=>{
          this.dialogShow = false
          this.$emit("update:visible", false);
          this.$message.error(err)
        })
      }else{
        updateAiforgeRole({id:this.data.ID,...this.params}).then((res)=>{
          if(res.data.code == 0){
            this.$message.success(this.$t('submittedSuccessfully'))
            this.dialogShow = false
            this.$emit("update:visible", false);
            this.$emit("refresh");
          }else{
            this.$message.error(res.data.msg)
          }
        }).catch((err)=>{
          this.dialogShow = false
          this.$emit("update:visible", false);
          this.$message.error(err)
        })
      }
      
    },
    cancel() {
      this.dialogShow = false;
      this.$emit("update:visible", false);
    }
  },
  mounted() {
  },
};
</script>

<style scoped lang="less">
.dlg-content {
   margin: 24px 0px;
  .form{
    .form-row{
       display: flex;
        margin-bottom: 20px;
        &.border{
            border: 1px solid #d4d4d5;
            border-radius: 5px;
            min-height: 100px;
            position: relative;
            margin: 32px 80px 32px 80px;
            .title{
                position: absolute;
                top: -10px;
                left: 12px;
                padding: 0 4px;
                background: #fff;
                color: #101010;
            }
        }
        label{
             &.required::before{
                content: "*";
                color: red;
                margin-right: 5px;
            }
            width: 170px;
            color: rgba(136,136,136,1);
            font-size: 14px;
            box-sizing: border-box;
            text-align: right;
        }
        .content{
            flex: 1;
            &.box{
                padding: 20px;
                .el-checkbox{
                    display: flex;
                    margin-bottom: 12px;
                    align-items: center;
                    &:last-child{
                      margin-top: -4px;
                    }
                }
                .el-checkbox-group{
                    margin-left: 40px;
                    margin-top: 20px;
                }
            }
            &.error {
              /deep/.el-input__inner,
              /deep/.el-textarea__inner {
                color: #9f3a38;
                background: #fff6f6;
                border-color: #e0b4b4;

                &:visited {
                  border-color: #e0b4b4;
                }

                &:focus {
                  border-color: #e0b4b4;
                }

                &:active {
                  border-color: #e0b4b4;
                }
              }
            }
            .tips {
                font-size: 12px;
                color: rgba(136, 136, 136, 1);
                margin-top: 10px;
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
/deep/ .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner{
  background-color: #409eff;
  border-color: #DCDFE6;
}
/deep/ .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner::after {
  border-color: #fff;
}
/deep/ .el-input.is-disabled .el-input__inner{
  color:#101010
}
/deep/ .el-checkbox__input.is-disabled + span.el-checkbox__label {
  color:#101010
}
</style>
