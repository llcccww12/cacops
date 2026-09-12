<template>
  <div class="base-dlg">
    <BaseDialog :visible.sync="dialogShow" width="80%"
      :title="dialogTitle" @open="open" @opened="opened" @close="close" @closed="closed">
      <div class="dlg-wrap">
        <div class="dlg-content">
          <div class="form" style="width:50%">
              <div class="form-row" style="height:32px">
                  <label>{{$t('userRole.roleType')}}：</label>
                  <div class="content">
                      <span style="color:#101010">{{$t('userRole.reCategory')}}</span>
                  </div>
              </div>
              <div class="form-row">
                  <label class="required">{{$t('userRole.roleName')}}：</label>
                  <div class="content">
                      <el-input v-model="params.name" :disabled="type=='view'" style="width:80%;"
                      maxlength="80" :placeholder="$t('resourcesManagement.roleNameTips')"></el-input>
                  </div>
              </div>
          </div>
          <div class="form" style="margin-left:2rem;flex:1">
              <div class="form-row">
                  <label class="required">{{$t('userRole.isItDefault')}}：</label>
                  <div class="content">
                      <el-select v-model="params.isCommon" :disabled="type !== 'add'" style="width:80%;">
                        <el-option v-for="item in commonTypeList" :key="item.k" :label="item.v" :value="item.k" />
                      </el-select>
                  </div>
              </div>
              <div class="form-row">
                  <label for="">{{$t('userRole.roleDescription')}}：</label>
                  <div class="content">
                    <el-input type="textarea" :rows="1" v-model="params.description" :disabled="type === 'view'"
                    style="width:80%;" maxlength="800" :placeholder="$t('resourcesManagement.roleDescTips')"></el-input>
                  </div>
              </div>
          </div>
        </div>
        <div class="resource-wrap">
          <div class="resource-container">
            <div style="display:flex;justify-content: space-between;">
              <div>
                <el-select size="medium" v-model="aiCenter" filterable>
                  <el-option v-for="item in aiCenterList" :key="item.k" :label="item.v" :value="item.k" />
                </el-select>
                <el-cascader size="medium" v-model="resource" :options="resourceList" :props="{ checkStrictly: true }">
                  <template slot-scope="{ node, data }">
                    <span>{{ data.label }}</span>
                    <span v-if="!node.isLeaf"> ({{ data.children.length }}) </span>
                  </template>
                </el-cascader>
                <el-select size="medium" v-model="specStatus">
                  <el-option v-for="item in specStatusList" :key="item.k" :label="item.v" :value="item.k" />
                </el-select>
              </div>
              <div>
                <el-select size="medium" v-model="sortStatus" @change="sortChange">
                  <el-option v-for="item in sortStatusList" :key="item.k" :label="item.v" :value="item.k" />
                </el-select>
              </div>
            </div>
            <el-radio-group v-model="tabPosition" @change="changeTaskType" style="margin-top: 12px;" size="medium">
              <el-radio-button v-for="item in taskTypeList" :label="item.k" :key="item.k">{{item.v}} ({{selectListMap[item.k].length}})</el-radio-button>
            </el-radio-group>
            <div class="table-wrap" ref="tableWrap" v-loading="loading">
              <el-table :data="filterData" :span-method="objectSpanMethod" style="width:100%;" v-if="sortStatus==='SourceSpecId'" max-height="490">
                <el-table-column prop="SpecStr" :label="$t('resourcesManagement.resourceSpecification')" min-width="38%">
                  <template slot-scope="scope">
                    <el-checkbox :disabled="type=='view'" v-model="scope.row.firstChecked" @change="firstChange(scope.row,scope.row.firstChecked)"></el-checkbox>
                    <span>{{scope.row.SpecStr}}</span>
                  </template>
                </el-table-column>
                <el-table-column prop="AiCenterName" :label="$t('resourcesManagement.aiCenter')" header-align="center" min-width="16%">
                  <template slot-scope="scope">
                    <el-checkbox :disabled="type=='view'" v-model="scope.row.secondChecked" @change="secondChange(scope.row,scope.row.secondChecked)"></el-checkbox>
                    <span>{{scope.row.AiCenterName}}</span>
                  </template>
                </el-table-column>
                <el-table-column prop="QueueStr" :label="$t('resourcesManagement.resQueue')" align="left" header-align="center" min-width="26%">
                  <template slot-scope="scope">
                    <span v-html="scope.row.QueueStr + scope.row.NetworkTypeStr + scope.row.visualizationStr"></span>
                  </template>
                </el-table-column>
                <el-table-column prop="UnitPrice" :label="`${$t('resourcesManagement.unitPrice')}(${$t('resourcesManagement.point_hr')})`" header-align="center" align="center" min-width="10%"></el-table-column>
                <el-table-column prop="Status" :label="$t('resourcesManagement.status')" header-align="center" align="center" min-width="10%">
                  <template slot-scope="scope">
                    <el-tag type="warning" effect="dark" v-if="scope.row.Status===1">{{$t('resourcesManagement.willOnShelf')}}</el-tag>
                    <el-tag type="success" effect="dark" v-else-if="scope.row.Status===2">{{$t('resourcesManagement.onShelf')}}</el-tag>
                    <el-tag type="danger" effect="dark" v-else>{{$t('resourcesManagement.offShelf')}}</el-tag>
                  </template>
                </el-table-column>
              </el-table>
              <el-table :data="filterData" :span-method="objectSpanMethod" style="width:100%;" v-if="sortStatus==='QueueId'" max-height="490">
                <el-table-column prop="AiCenterName" :label="$t('resourcesManagement.aiCenter')" min-width="16%">
                  <template slot-scope="scope">
                    <el-checkbox :disabled="type=='view'" v-model="scope.row.firstChecked" @change="firstChange(scope.row,scope.row.firstChecked,'QueueId')"></el-checkbox>
                    <span>{{scope.row.AiCenterName}}</span>
                  </template>
                </el-table-column>
                <el-table-column prop="QueueStr" :label="$t('resourcesManagement.resQueue')" align="left" header-align="center" min-width="26%">
                  <template slot-scope="scope">
                    <span v-html="scope.row.QueueStr + scope.row.NetworkTypeStr + scope.row.visualizationStr"></span>
                  </template>
                </el-table-column>
                <el-table-column prop="SpecStr" :label="$t('resourcesManagement.resourceSpecification')" min-width="38%">
                  <template slot-scope="scope">
                    <el-checkbox :disabled="type=='view'" v-model="scope.row.secondChecked" @change="secondChange(scope.row,scope.row.secondChecked,false,'QueueId')"></el-checkbox>
                    <span>{{scope.row.SpecStr}}</span>
                  </template>
                </el-table-column>
                <el-table-column prop="UnitPrice" :label="`${$t('resourcesManagement.unitPrice')}(${$t('resourcesManagement.point_hr')})`" header-align="center" align="center" min-width="10%"></el-table-column>
                <el-table-column prop="Status" :label="$t('resourcesManagement.status')" header-align="center" align="center" min-width="10%">
                  <template slot-scope="scope">
                    <el-tag type="warning" effect="dark" v-if="scope.row.Status===1">{{$t('resourcesManagement.willOnShelf')}}</el-tag>
                    <el-tag type="success" effect="dark" v-else-if="scope.row.Status===2">{{$t('resourcesManagement.onShelf')}}</el-tag>
                    <el-tag type="danger" effect="dark" v-else>{{$t('resourcesManagement.offShelf')}}</el-tag>
                  </template>
                </el-table-column>
              </el-table>
              <el-table :data="filterData" :span-method="objectSpanMethod" style="width:100%;" v-if="sortStatus==='Selected'" max-height="490">
                <el-table-column prop="SpecStr" :label="$t('resourcesManagement.resourceSpecification')" min-width="38%">
                  <template slot-scope="scope">
                    <el-checkbox :disabled="true" v-model="scope.row.firstChecked" @change.native="firstChange(scope.row,scope.row.firstChecked)"></el-checkbox>
                    <span>{{scope.row.SpecStr}}</span>
                  </template>
                </el-table-column>
                <el-table-column prop="AiCenterName" :label="$t('resourcesManagement.aiCenter')" min-width="16%">
                  <template slot-scope="scope">
                    <el-checkbox :disabled="true" v-model="scope.row.secondChecked" @change.native="secondChange(scope.row,scope.row.secondChecked,false)"></el-checkbox>
                    <span>{{scope.row.AiCenterName}}</span>
                  </template>
                </el-table-column>
                <el-table-column prop="QueueStr" :label="$t('resourcesManagement.resQueue')" align="left" header-align="center" min-width="26%">
                  <template slot-scope="scope">
                    <span v-html="scope.row.QueueStr + scope.row.NetworkTypeStr + scope.row.visualizationStr"></span>
                  </template>
                </el-table-column>
                <el-table-column prop="UnitPrice" :label="`${$t('resourcesManagement.unitPrice')}(${$t('resourcesManagement.point_hr')})`" header-align="center" align="center" min-width="10%"></el-table-column>
                <el-table-column prop="Status" :label="$t('resourcesManagement.status')" header-align="center" align="center" min-width="10%">
                  <template slot-scope="scope">
                    <el-tag type="warning" effect="dark" v-if="scope.row.Status===1">{{$t('resourcesManagement.willOnShelf')}}</el-tag>
                    <el-tag type="success" effect="dark" v-else-if="scope.row.Status===2">{{$t('resourcesManagement.onShelf')}}</el-tag>
                    <el-tag type="danger" effect="dark" v-else>{{$t('resourcesManagement.offShelf')}}</el-tag>
                  </template>
                </el-table-column>
              </el-table>
            </div>
            <div style="padding: 12px 0;">
              <el-button type="primary" :disabled="type==='view'" class="btn confirm-btn" @click="confirm">{{ $t('confirm') }}</el-button>
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
import { getResSpecificationListAll, getResourceType, getAiCenterList, getResQueueCode,addAiforgeRole, updateAiforgeRole  } from '~/apis/modules/resources';
import { getListValueWithKey } from '~/utils';
import { NEW_JOB_TYPE, CLUSTERS, ACC_CARD_TYPE, SPECIFICATION_STATUS, NETWORK_TYPE_VALUE, NEW_JOB_TYPE_OBJ } from '~/const';


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
      clusterList: [...CLUSTERS],
      taskType:'',
      taskTypeList: [...NEW_JOB_TYPE],
      accCardTypeList: [...ACC_CARD_TYPE],
      networkTypeList: [...NETWORK_TYPE_VALUE],
      // commonType: 0,
      commonTypeList: [{k: 0,v: this.$t('userRole.default'),},{k: 1,v: this.$t('userRole.notDefault')}],
      aiCenter: '',
      aiCenterList: [{ k: '', v: this.$t('resourcesManagement.allAiCenter') }],
      resource: [''],
      resourceList: [],
      specStatus:'',
      specStatusList: [{ k: '', v: this.$t('resourcesManagement.allStatus') }, ...SPECIFICATION_STATUS],
      sortStatus:'SourceSpecId',
      sortStatusList:[
        {k: 'SourceSpecId', v: this.$t('resourcesManagement.accordingSpec')},
        {k: 'QueueId', v: this.$t('resourcesManagement.accordingQueue')},
        {k: 'Selected', v: this.$t('resourcesManagement.accordingSelect')},
      ],
      tableData: [],
      tabPosition: NEW_JOB_TYPE[0].k,
      selectListMap: NEW_JOB_TYPE_OBJ,
      params: {
        name:"",
        type:1,
        isCommon:0,
        description:'',
        jobType:'',
        specId:'',
      },
      getResourceFlag: false,
      loading: false,
      getAiCenterFlag: false,
      getQueueFlag: false,
    };
  },
  watch: {
    visible: function (val) {
      this.dialogShow = val;
    },
    filterData:function(val){
      // if(this.sortStatus === 'Selected'){
      //   this.filterData.sort((a,b)=>{
      //     return b.secondChecked-a.secondChecked
      //   })
      // }
    }

  },
  computed: {
    mergingRows() {
      let props = ''
      if(this.sortStatus==='QueueId'){
        this.filterData.sort(this.queuedSort)
        props = 'QueueId'
      }else if(this.sortStatus==='SourceSpecId'){
        this.filterData.sort(this.specIdSort)
        props = 'SourceSpecId'
      }else{
        this.filterData.sort(this.selectedSort)
        props = 'SourceSpecId'
      }
      const mergingRows = []
      let mergingPos = 0
      for (let i = 0; i < this.filterData.length; i++) { // tabledata 表格数据源
        if (i === 0) {
          mergingRows.push(1)
          mergingPos = 0
        } else {
          if (this.filterData[i][props] === this.filterData[i - 1][props]) {
          // 哪些数据是要合并的 合并的条件是什么 此处合并条件为categoryName 相同则进行合并
            mergingRows[mergingPos] += 1
            mergingRows.push(0)
          } else {
            mergingRows.push(1)
            mergingPos = i
          }
        }
      }
      return mergingRows
    },
    filterData(){
      return this.tableData.filter((item)=>{
        return item.AiCenterCode == this.aiCenter ||  this.aiCenter == ''
      }).filter((item)=>{
        if(this.resource.length===1){
          if(this.resource[0] === 'GPGPU'){
            return item.ComputeResource.includes(this.resource[0])
          }else{
            return item.ComputeResource == this.resource[0] ||  this.resource[0] == ''
          }
        }
        return true
      }).filter((item)=>{
        if(this.resource.length===2){
          return item.AccCardType == this.resource[1]
        }
        return true
      }).filter((item)=>{
        return item.Status == this.specStatus ||  this.specStatus == ''
      })
    },
    dialogTitle(){
      switch(this.type){
        case 'add':
          return this.$t('userRole.newReRole')
        case 'edit':
          return this.$t('userRole.editReRole')
        case 'view':
          return this.$t('userRole.viewReRole')
      }
    },
  },
  methods: {
    queuedSort(a,b){
      if(a.QueueId === b.QueueId){
        if(a.ComputeResource === b.ComputeResource){
          if(a.AccCardsNum === b.AccCardsNum){
            if(a.AccCardType === b.AccCardType){
              if(a.GPUMemGiB === b.GPUMemGiB){
                if(a.CpuCores === b.CpuCores){
                  if(a.MemGiB === b.MemGiB){
                    return 0
                  }else{
                    return a.MemGiB-b.MemGiB
                  }
                }else{
                  return a.CpuCores-b.CpuCores
                }
              }
            }else{
              return a.AccCardType.localeCompare(b.AccCardType)
            }
          }else{
            return a.AccCardsNum-b.AccCardsNum
          }
        }else{
          return a.ComputeResource.localeCompare(b.ComputeResource)
        }
        return a.SpecStr.localeCompare(b.SpecStr)
      }else{
        return b.QueueId-a.QueueId
      }
    },
    specIdSort(a,b){
      if(a.ComputeResource === b.ComputeResource){
        if(a.AccCardType === b.AccCardType){
          if(a.AccCardsNum === b.AccCardsNum){
            if(a.CpuCores === b.CpuCores){
              if(a.MemGiB === b.MemGiB){
                if(a.ShareMemGiB === b.ShareMemGiB){
                  if(a.SourceSpecId === b.SourceSpecId){
                    return 0
                  }else{
                    return a.SourceSpecId.localeCompare(b.SourceSpecId)
                  }
                }else{
                  return a.ShareMemGiB-b.ShareMemGiB
                }
              }else{
                return a.MemGiB-b.MemGiB
              }
            }else{
              return a.CpuCores-b.CpuCores
            }
          }else{
            return a.AccCardsNum-b.AccCardsNum
          }
        }else{
          return a.AccCardType.localeCompare(b.AccCardType)
        }
      }else{
        return a.ComputeResource.localeCompare(b.ComputeResource)
      }
    },
    selectedSort(a,b){
      if(a.secondChecked === b.secondChecked){
        return this.specIdSort(a,b)
      }else{
        return b.secondChecked-a.secondChecked
      }
    },
    firstChange(row,checkValue,props='SourceSpecId'){
      this.filterData.forEach((item)=>{
        if(checkValue){
          if(item[props] == row[props]){
            item.secondChecked = true
            this.selectListMap[this.tabPosition].push(item.ID)
          }
        }else{
          if(item[props] == row[props]){
            console.log(props)
            item.secondChecked = false
            let index = this.selectListMap[this.tabPosition].indexOf(item.ID)
            if(index>-1){
              this.selectListMap[this.tabPosition].splice(index,1)
            }
          }
        }
      })
    },
    secondChange(row,checkValue,flag=false,props='SourceSpecId'){
      let checkNum = 0
      this.filterData.forEach((item) => {
        if (item[props] == row[props]) {
          checkNum += item.secondChecked ? 1 : 0
        }
      })
      this.filterData.forEach((item)=>{
        if(item[props] == row[props] && checkValue){
          item.firstChecked = true
          if(flag && item['ID'] == row['ID']){
            item.secondChecked = true
          }
        }
        if(item.ID == row.ID && checkValue){
          if(!flag){
            this.selectListMap[this.tabPosition].push(item.ID)
          }
        }
        if(item.ID == row.ID && !checkValue){
          let index = this.selectListMap[this.tabPosition].indexOf(item.ID)
          if(index>-1){
            this.selectListMap[this.tabPosition].splice(index,1)
          }
        }
        if(item[props] == row[props] && !checkValue){
          if(checkNum>0){
            item.firstChecked = true
          }else{
            item.firstChecked = false
          }
        }
      })
    },
    sortChange(val){
      this.changeTaskType('')
      this.$nextTick(()=>{
        let scrollElem = this.$refs.tableWrap
        scrollElem.scrollTo({ top: 0 });
      })
    },
    changeTaskType(type){
      this.resetTableData(this.selectListMap[this.tabPosition])
      this.$nextTick(()=>{
        let scrollElem = this.$refs.tableWrap
        scrollElem.scrollTo({ top: 0 });
      })
    },
    resetTableData(selectList){
      this.filterData.forEach((item)=>{
        item.firstChecked = false
        item.secondChecked = false
        selectList && selectList.length && selectList.forEach((selectItem)=>{
          if(item.ID == selectItem){
            if(this.sortStatus==='Selected'){
              this.secondChange(item,true,true,'ID')
            }else{
              this.secondChange(item,true,true,this.sortStatus)
            }
          }
        })
      })
      // if(this.sortStatus === 'Selected'){
      //   this.filterData.sort((a,b)=>{
      //     return b.secondChecked-a.secondChecked
      //   })
      // }
    },
    initResourceType(){
      if(!this.getQueueFlag){
        getResourceType().then((res)=>{
          const data = res.data
          this.resourceList = this.flatArrayToTree(data)
          this.getQueueFlag = true
        }).catch((error)=>{
          this.$message.error(error)
        })
      }

    },
    objectSpanMethod({row, column, rowIndex, columnIndex }) {
      if (columnIndex === 0) { // 第一列
        const _row = this.mergingRows[rowIndex]
        const _col = _row > 0 ? 1 : 0
        return {
          rowspan: _row,
          colspan: _col
        }
      }
    },
    flatArrayToTree(flatData){
      const map = {};
      const resultArray = [{ value: '', label: this.$t('resourcesManagement.allComputeResource')}]
      flatData.forEach(item => {
        if(item.ComputeResource.includes('GPGPU') ){
          item.ComputeResource = 'GPGPU'
        }
        if(map.hasOwnProperty(item.ComputeResource)){
          map[item.ComputeResource].children.push({'label':getListValueWithKey(this.accCardTypeList, item.AccCardType),'value':item.AccCardType})
        }else{
          map[item.ComputeResource] = {"label":item.ComputeResource,"value":item.ComputeResource, children: [{'label':getListValueWithKey(this.accCardTypeList, item.AccCardType),'value':item.AccCardType}] };
        }
      })

      for (let key in map){
        const child = map[key].children
        child.sort((a,b)=>{
          return a.label.localeCompare(b.label)
        })
        const childMap = new Map()
        const childResult = []
        for (let item of child){
          if(!childMap.has(item.label)){
            childMap.set(item.label,item)
          }
        }
        childResult = [...childMap.values()]
        map[key].children = childResult
      }
      for (let key in map){
        resultArray.push(map[key])
        resultArray.sort((a,b)=>{
          return a.label.localeCompare(b.label)
        })
      }
      return resultArray
    },
    open() {
      this.initTableData()
      this.initResourceType()
      this.getAiCenterList()
      // this.getQueueList()

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
    addResourceRole(params){
      addAiforgeRole(params).then((res)=>{
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
    },
    confirm() {
      let sumSelectedList = []
      let sumJobTypeList = []
      Object.keys(this.selectListMap).forEach((key) => {
        const selectedList = this.selectListMap[key]
        if(selectedList.length){
          sumSelectedList = sumSelectedList.concat(selectedList)
          sumJobTypeList = sumJobTypeList.concat(new Array(selectedList.length).fill(key))
        }
      })
      this.params.jobType = sumJobTypeList.join(",")
      this.params.specId = sumSelectedList.join(",")
      if(!this.params.name){
        this.$message.error("请输入角色名称")
        return
      }
      if(!sumJobTypeList.length){
        this.$message.error("请选择资源")
        return
      }
      if(this.type === 'add'){
        if(!this.params.isCommon){
          this.$confirm('当前创建角色应用于所有用户', "提示", {
            confirmButtonText: this.$t('confirm1'),
            cancelButtonText: this.$t('cancel'),
            type: 'warning',
            lockScroll: false,
            closeOnClickModal:false,
            showClose:false,
            closeOnPressEscape:false,
          }).then(()=>{
            this.addResourceRole(this.params)
          }).catch(()=>{

          })
        }else{
          this.addResourceRole(this.params)
        }

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
    },
    getAiCenterList() {
      if(!this.getAiCenterFlag){
        getAiCenterList().then(res => {
          res = res.data;
          this.getAiCenterFlag = true
          if (res.Code === 0) {
            const list = res.Data;
            const data = list.map(item => {
              return {
                k: item.AiCenterCode,
                v: item.AiCenterName
              };
            });
            this.aiCenterList.splice(1, Infinity, ...data);
          }
        }).catch(err => {
          console.log(err);
        });
      }
    },
    getQueueList() {
      getResQueueCode().then(res => {
        res = res.data;
        if (res.Code === 0) {
          const data = res.Data;
          const list = [];
          for (let i = 0, iLen = data.length; i < iLen; i++) {
            const item = data[i];
            const queueName = item.QueueName ? `【${item.QueueName}】` : '';
            const queueType = item.QueueType ? `【${item.QueueType}】` : '';
            list.push({
              k: item.ID,
              v: `${item.QueueCode}${queueName}${queueType}(${getListValueWithKey(this.clusterList, item.Cluster)} - ${item.AiCenterName}) ${item.ComputeResource}(${item.AccCardType})`,
            });
          }
          // this.queueList.push(...list);
        }
      }).catch(err => {
        console.log(err);
      });
    },
    restorePamras(){
      Object.keys(this.selectListMap).forEach((key)=>{
        this.selectListMap[key] = []
      })
      this.resetSelectOption()
      const map = {}
      this.params.name = this.data.Name
      this.params.description = this.data.Description
      this.params.isCommon = this.data.IsCommon
      this.data.taskType.forEach((item,index)=>{
        if(!map.hasOwnProperty(item)){
          map[item] = [Number(this.data.specId[index])]
        }else{
          map[item].push(Number(this.data.specId[index]))
        }
      })
      Object.keys(map).length && Object.keys(map).forEach((key)=>{
        if(key){
          this.selectListMap[key] = map[key]
        }
      })
      this.sortStatus = 'Selected'
      this.tabPosition = this.data.taskType[0]
      if(this.tabPosition){
        this.changeTaskType(this.tabPosition)
      }else{
        this.restParams()
        this.params.name = this.data.Name
        this.params.description = this.data.Description
        this.params.isCommon = this.data.IsCommon
      }
    },
    restParams(){
      this.params = { name:"", type:1, isCommon:0, description:'', jobType:'', specId:'' }
      this.tabPosition = NEW_JOB_TYPE[0].k
      Object.keys(this.selectListMap).forEach((key)=>{
        this.selectListMap[key] = []
      })
      this.resetSelectOption()
      this.changeTaskType(this.tabPosition)
    },
    resetSelectOption(){
      this.aiCenter = '',
      this.resource = ['']
      this.specStatus = ''
      this.sortStatus = 'SourceSpecId'
    },
    initTableData(){
      if(this.getResourceFlag){
          if (this.type === 'add'){
            this.restParams()
          }else{
            this.restorePamras()
          }
      }else{
        this.loading = true
        getResSpecificationListAll({available: 1}).then((res=>{
          res = res.data;
          this.loading = false
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
                firstChecked: false,
                secondChecked: false,
                visualizationStr: `, ${this.$t('cloudbrainObj.visualization')}:${item.EnableVisualization ? this.$t('resourcesManagement.enable') : this.$t('resourcesManagement.notEnable')}`,
              }
            });

            this.tableData = data;
            this.getResourceFlag = true
            if (this.type === 'add') {
              //
            } else {
              this.restorePamras()
            }
          }
        })).catch((err) => {
          this.loading = false
          console.log(err)
        })
      }

    }
  },
  mounted() {
  },
};
</script>

<style scoped lang="less">
.dlg-wrap{
  display: flex;
  flex-direction:column;
  height:100%;
  .dlg-content {
    margin: 12px;
    margin-bottom: 0px;
    display: flex;
    .form{
      .form-row{
        display: flex;
        align-items: center;
        margin-bottom: 20px;
          &.border{
              border: 1px solid #d4d4d5;
              border-radius: 5px;
              min-height: 100px;
              margin-top: 24px;
              position: relative;
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
              width: 140px;
              text-align: right;
              color: rgba(136,136,136,1);
              font-size: 14px;
              box-sizing: border-box;
          }
          .content{
              flex: 1;
              &.box{
                  padding: 20px;
                  .el-checkbox{
                      display: flex;
                      margin-bottom: 12px;
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
  }
  .resource-wrap{
    margin: 24px;
    margin-top: 0;
    height: calc(100% - 128px);
    .resource-container{
      height: 100%;
      display: flex;
      flex-direction: column;
    }
    .table-wrap{
      margin: 12px 0px;
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
}
/deep/ .el-dialog{
  margin-top: 5% !important;
  // height: 85%;
  flex-direction: column;
  display: flex;
}
/deep/ .el-dialog__body{
  height: 0;
  flex: 1;
}
/deep/ .el-radio-button:first-child .el-radio-button__inner{
  border-radius: 4px 0 0 0;
}
/deep/ .el-radio-button:last-child .el-radio-button__inner{
  border-radius: 0 4px 0 0;
}
/deep/ .el-radio-button__orig-radio:checked + .el-radio-button__inner {
  color:#3894ff;
  background-color: #f5f5f6;
  border-color: #DCDFE6;
  box-shadow: none;
}
/deep/ .el-table thead tr th {
  background-color: #F5F5F6;
}
/deep/ .el-table  td{
  border-right:none;
  border-bottom: 1px solid #11001130;
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
</style>
