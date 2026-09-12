<template>
<div>
    <div class="dlg-content" v-if="JSON.stringify(reRoleCommonData) != '{}'">
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
                  <el-input style="width:80%;" disabled :value="reRoleCommonData.Name"></el-input>
              </div>
          </div>
      </div>
      <div class="form" style="margin-left:2rem;flex:1">
          <div class="form-row">
              <label class="required">{{$t('userRole.isItDefault')}}：</label>
              <div class="content">
                  <el-input style="width:80%;" disabled :value="reRoleCommonData.IsCommon ? $t('userRole.notDefault') : $t('userRole.default')"></el-input>
              </div>
          </div>
          <div class="form-row">
              <label for="">{{$t('userRole.roleDescription')}}：</label>
              <div class="content">
                <el-input type="textarea" disabled :rows="1" style="width:80%;" :value="reRoleCommonData.Description"></el-input>
              </div>
          </div>
      </div>
    </div>
    <el-radio-group v-model="tabPosition" @change="changeTaskType" style="margin-top: 12px;" size="medium">
        <el-radio-button v-for="item in taskTypeList" :label="item.k" :key="item.k">{{item.v}} ({{showSpecTableData[item.k].length}})</el-radio-button>
    </el-radio-group>
    <el-table :data="filterData" :span-method="objectSpanMethod" style="width:100%;" border max-height="550">
        <el-table-column prop="SpecStr" :label="$t('resourcesManagement.resourceSpecification')" min-width="42%">
        <template slot-scope="scope">
            <span>{{scope.row.SpecStr}}</span>
        </template>
        </el-table-column>
        <el-table-column prop="AiCenterName" :label="$t('resourcesManagement.aiCenter')" min-width="15%">
        <template slot-scope="scope">
            <span>{{scope.row.AiCenterName}}</span>
        </template>
        </el-table-column>
        <el-table-column prop="QueueStr" :label="$t('resourcesManagement.resQueue')" align="left" header-align="center" min-width="34%">
        <template slot-scope="scope">
            <span v-html="scope.row.QueueStr + scope.row.NetworkTypeStr +  scope.row.visualizationStr"></span>
        </template>
        </el-table-column>
        <el-table-column prop="UnitPrice" :label="`${$t('resourcesManagement.unitPrice')}(${$t('resourcesManagement.point_hr')})`" min-width="10%"></el-table-column>
        <el-table-column prop="Status" :label="$t('resourcesManagement.status')" min-width="9%">
          <template slot-scope="scope">
            <el-tag type="warning" effect="dark" v-if="scope.row.Status===1">{{$t('resourcesManagement.willOnShelf')}}</el-tag>
            <el-tag type="success" effect="dark" v-else-if="scope.row.Status===2">{{$t('resourcesManagement.onShelf')}}</el-tag>
            <el-tag type="danger" effect="dark" v-else>{{$t('resourcesManagement.offShelf')}}</el-tag>
          </template>
        </el-table-column>
    </el-table>
</div>
</template>
<script>
import { NEW_JOB_TYPE } from '~/const';
export default {
   name: '',
   components: {
     
   },
   mixins: [],
   props: {
     specificationList: {type: Array, default: () => ([])},
     showSpecTableData: {type: Object, default: () => ({})},
     radioInit: {type: String, default: ''},
     reRoleCommonData: {type: Object, default: () => ({})},
   },
   data() {
     return {
       tabPosition: NEW_JOB_TYPE[0].k,
       taskTypeList: [...NEW_JOB_TYPE],
       filterData: [],
     }
   },
   computed: {
     mergingRows() { 
      const mergingRows = []
      let mergingPos = 0
      for (let i = 0; i < this.filterData.length; i++) { // tabledata 表格数据源
        if (i === 0) {
          mergingRows.push(1)
          mergingPos = 0
        } else {
          if (this.filterData[i]['SourceSpecId'] === this.filterData[i - 1]['SourceSpecId']) { 
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
   },
   watch: {
     radioInit(val) {
      if(val){
        this.tabPosition = val
      }
     }
   },
   mounted() {
   },
   methods: {
    changeTaskType(type){
      this.filterData = this.specificationList.filter((item)=>{
          return this.showSpecTableData[type].includes(String(item.ID))
      })
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
   }
};
</script>
<style lang='less' scoped>
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
/deep/ .el-table--border td{
  border-right:none;
  border-bottom: 1px solid #11001130;
}
/deep/ .el-table--border th{
  border-right:none
}
/deep/ .el-table thead tr th {
  background-color: #F5F5F6;
}
/deep/ .el-input.is-disabled .el-input__inner{
  color:#101010
}
::-webkit-scrollbar-track {
    background: transparent;
    border-radius: 0;
}
::-webkit-scrollbar {
  -webkit-appearance: none;
  width: 4px;
  height: 4px;
}
::-webkit-scrollbar-thumb {
    cursor: pointer;
    border-radius: 5px;
    background: rgba(0, 0, 0, 0.15);
    transition: color 0.2s ease;
}
</style>