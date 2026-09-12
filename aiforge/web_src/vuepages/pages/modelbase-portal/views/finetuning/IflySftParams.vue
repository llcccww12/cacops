<template>
<div class="form-row">
    <div class="left-area">
        <div class="title">
            <span class="required">{{$t('cloudbrainObj.paramsSetting')}}</span>
        </div>
        <div class="content">
            <div class="basic-params-wrap" v-for="(item, index) in paramterList" :key="index">
                <div>
                    <div class="title-params">{{item.title}}</div>
                    <div>
                        <el-table
                            :data="item.data"
                            :header-cell-style="{background:'#eef1f6',color:'#606266'}"
                            :cell-style="columnStyle"
                            style="width: 100%">
                            <el-table-column
                            :label="$t('cloudbrainObj.parameters')"
                            width="180"
                            >
                                <template slot-scope="scope">
                                    {{$t('modelFinetune.' + scope.row.parameter)}}
                                </template>
                            </el-table-column>
                            <el-table-column
                            :label="$t('cloudbrainObj.valueAndContent')"
                            width="180">
                            <template slot-scope="scope">
                                <div v-if="scope.row.type=='input'">
                                    <el-input v-model="scope.row.value"></el-input>
                                </div>
                                <div v-else-if="scope.row.type=='slider'">
                                    <el-slider v-model="scope.row.value" show-input :show-tooltip="false" 
                                    :show-input-controls="false" input-size="mini"
                                    :min="scope.row.min" :max="scope.row.max" :step="scope.row.step"></el-slider>
                                </div>
                                <div v-else-if="scope.row.type=='select'">
                                    <el-select v-model="scope.row.value" placeholder="请选择">
                                        <el-option
                                        v-for="item in scope.row.options"
                                        :key="item.value"
                                        :label="item.label"
                                        :value="item.value">
                                        </el-option>
                                    </el-select>
                                </div>
                            </template>
                            </el-table-column>
                            <el-table-column :label="$t('explanation')">
                                <template slot-scope="scope">
                                {{$t(scope.row.desc)}}
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                </div>
            </div>
            
        </div>
    </div>
    <div class="right-area"></div>
</div>
</template>

<script>
import {sparkPamrams} from './constParams'
export default {
    props: {
    
    },
    data() {
        return {
            paramterList:[
                {title:this.$t('cloudbrainObj.basicParameters'),data:sparkPamrams},
            ],
        };
    },
    components: { },

    methods: {
        columnStyle({ row, column, rowIndex, columnIndex }){
            if(columnIndex == 0){
                return "color:#101010";
            }
        },
        getParams(){
            const sumArray = this.paramterList.reduce((accumulator, currentValue)=>{return accumulator.concat(currentValue.data)},[])
            let paramsList = sumArray.reduce((accumulator, currentValue)=>{
                let temp
                temp = {
                    label:`${currentValue.parameter}.${currentValue.parameter}`,
                    value:String(currentValue.value)
                }
                return accumulator.concat(temp)
            },[])
            return paramsList
        }
    },
    mounted() { },
};
</script>

<style scoped lang="less">
@import '~/components/cloudbrain/cloudbrain.less';
/deep/ .el-slider__input{
    width:80px
}
/deep/ .el-table tbody tr:hover>td { 
    background-color:rgb(230,240,255)!important
}
/deep/ .el-slider__runway.show-input{
    margin-right: 90px;
}
/deep/ .el-slider__button { // 拖动的滑块的样式
    width: 10px;
    height: 10px;
}
.basic-params-wrap{
    border-radius: 5px;
    // height: 151px;
    width: 850px;
    color: rgba(0,122,255,1);
    border: 1px solid rgba(225,227,230,1);
    position: relative;
    padding: 24px;
    margin-bottom: 14px;
    .title-params{
    position: absolute;
    top: -10px;
    left: 10px;
    z-index: 10;
    padding: 0 5px;
    background: #fff;
    color: rgb(111, 118, 165);
}
}

</style>
    