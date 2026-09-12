<template>
  <div class="chart-container">
    <div class="wait-wrap" v-if="evalRunning">
      <svg xmlns="http://www.w3.org/2000/svg" style="margin-right:5px" viewBox="0 0 24 24" width="14" height="14"
        class="rotating" fill="#101010">
        <path
          d="M6 4H4V2H20V4H18V6C18 7.61543 17.1838 8.91468 16.1561 9.97667C15.4532 10.703 14.598 11.372 13.7309 12C14.598 12.628 15.4532 13.297 16.1561 14.0233C17.1838 15.0853 18 16.3846 18 18V20H20V22H4V20H6V18C6 16.3846 6.81616 15.0853 7.8439 14.0233C8.54682 13.297 9.40202 12.628 10.2691 12C9.40202 11.372 8.54682 10.703 7.8439 9.97667C6.81616 8.91468 6 7.61543 6 6V4ZM8 4V6C8 6.88457 8.43384 7.71032 9.2811 8.58583C10.008 9.33699 10.9548 10.0398 12 10.7781C13.0452 10.0398 13.992 9.33699 14.7189 8.58583C15.5662 7.71032 16 6.88457 16 6V4H8ZM12 13.2219C10.9548 13.9602 10.008 14.663 9.2811 15.4142C8.43384 16.2897 8 17.1154 8 18V20H16V18C16 17.1154 15.5662 16.2897 14.7189 15.4142C13.992 14.663 13.0452 13.9602 12 13.2219Z">
        </path>
      </svg>
      <span>{{ this.$t('cloudbrainObj.eval_task_ing') }}</span>
    </div>
    <template v-else>
      <div class="tips">
        * {{$t('modelFinetune.evalTips')}}
      </div>
      <div>
        <el-select v-model="dataset" @change="changeDataset" style="margin: 10px 0;color: #101010;">
          <el-option v-for="item in selectDataset" :value="item.v" :key="item.k" :label="item.k"></el-option>
        </el-select>
        <el-select v-model="resultType" @change="changeDataset" style="margin: 10px 0;color: #101010;">
          <el-option v-for="item in resultTypeList" :value="item.v" :key="item.name" :label="item.k"></el-option>
        </el-select>
      </div>
      <div class="table-wrap">
        <el-table :data="tableData" style="width: 100%" :fit="true" border>
          <el-table-column prop="type" :label="$t('repos.dataset')" width="180"></el-table-column>
          <el-table-column prop="raw_input" :label="$t('modelFinetune.evalModelInput')" min-width="400">
            <template slot-scope="scope">
              <div 
                class="file-view markdown chat raw-content" 
                :class="{'ellipsis-text': !scope.row.rawExpanded && scope.row.rawNeedsExpand}"
                :data-id="scope.row.id"
                style="font-size: 14px;" 
                v-html="scope.row.rawContent" 
              />
              <!-- <span 
                v-if="scope.row.rawNeedsExpand" 
                class="show-more-btn"
                @click.stop="toggleExpand(scope.row, 'rawExpanded')"
              >
                {{ scope.row.rawExpanded ? '收起' : '展开更多' }}
              </span> -->
            </template>
          </el-table-column>
          <el-table-column prop="predict" :label="$t('modelFinetune.evalModelOutput')" min-width="400">
            <template slot-scope="scope">
              <div 
                class="file-view markdown chat predict-content" 
                :class="{'ellipsis-text': !scope.row.predictExpanded && scope.row.predictNeedsExpand}"
                :data-id="scope.row.id"
                style="font-size: 14px;" 
                v-html="scope.row.predictContent" 
              />
              <span 
                v-if="scope.row.predictNeedsExpand" 
                class="show-more-btn"
                @click.stop="toggleExpand(scope.row, 'predictExpanded')"
              >
                {{ scope.row.predictExpanded ? $t('taskTmplObj.collapsed') : $t('expandMore') }}
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="gold" :label="$t('modelFinetune.standerAnswer')" width="160"></el-table-column>
          <el-table-column prop="pred" :label="$t('modelFinetune.evalModelOutAnswer')" width="160">
            <template slot-scope="scope">
              <div 
                class="file-view markdown chat pred-content" 
                :class="{'ellipsis-text': !scope.row.predExpanded && scope.row.predNeedsExpand}"
                :data-id="scope.row.id"
                style="font-size: 12px;" 
                v-html="scope.row.predContent" 
              />
              <span 
                v-if="scope.row.predNeedsExpand" 
                class="show-more-btn"
                @click.stop="toggleExpand(scope.row, 'predExpanded')"
              >
                {{ scope.row.predExpanded ? $t('taskTmplObj.collapsed') : $t('expandMore') }}
              </span>
            </template>
          </el-table-column>
        </el-table>
        <div style="text-align: center;margin-top: 40px;">
          <el-pagination
            background
            @current-change="handleCurrentChange"
            :current-page="currentPage"
            :page-size="pageSize"
            layout="total, prev, pager, next"
            :total="totalNum">
          </el-pagination>
        </div>
      </div>
    </template>
    
    
  </div>
</template>

<script>
import { getAiEvalDetailResult } from '~/apis/modules/cloudbrain';
import createRenderer from '~/pages/model/llms/componenes/render'
let md = null
export default {
  name: 'EvalDetail',
  props: {
    data: { type: Object, default: () => { return {} } },
  },
  data() {
    return {
      tableData: [],
      resultType: -1,
      resultTypeList: [{ k: this.$t('modelFinetune.allResult'), v: -1 }, { k: this.$t('modelFinetune.correctResult'), v: 1 }, { k: this.$t('modelFinetune.wrongResult'), v: 0 }],
      selectDataset: [],
      dataset: 'all',
      currentPage: 1,
      pageSize: 10,
      totalNum: 0,
      evalRunning: true 
    };
  },
  watch: {
    data: {
      deep: true,  // 深度观察，检测嵌套属性的变化
      handler(newVal) {
        console.log("xxxxxxxxxxx")
        if (newVal.task && ["STOPPED", "FAILED", "START_FAILED", "COMPLETED", "SUCCEEDED", "CREATED_FAILED"].includes(newVal.task.status)) {
          this.evalRunning = false;
        }
      },
    },
    tableData: {
      handler() {
        this.checkContentHeights();
      },
      deep: true
    }
  },
  methods: {
    changeDataset() {
      this.currentPage = 1;
      this.refresh();
    },
    handleCurrentChange(val) {
      this.currentPage = val;
      this.refresh();
    },
    toggleExpand(row, field) {
      this.$set(row, field, !row[field]);
    },
    checkContentHeights() {
      this.$nextTick(() => {
        this.tableData.forEach(row => {
          // 检查rawContent
          if (row.rawContent) {
            const rawEl = document.querySelector(`.raw-content[data-id="${row.id}"]`);
            if (rawEl) {
              const lineHeight = parseInt(window.getComputedStyle(rawEl).lineHeight);
              this.$set(row, 'rawNeedsExpand', rawEl.scrollHeight > lineHeight * 2);
            }
          }
          
          // 检查predictContent
          if (row.predictContent) {
            const predictEl = document.querySelector(`.predict-content[data-id="${row.id}"]`);
            if (predictEl) {
              const lineHeight = parseInt(window.getComputedStyle(predictEl).lineHeight);
              this.$set(row, 'predictNeedsExpand', predictEl.scrollHeight > lineHeight * 3);
            }
          }
          
          // 检查predContent
          if (row.predContent) {
            const predEl = document.querySelector(`.pred-content[data-id="${row.id}"]`);
            if (predEl) {
              const lineHeight = parseInt(window.getComputedStyle(predEl).lineHeight);
              this.$set(row, 'predNeedsExpand', predEl.scrollHeight > lineHeight * 2);
            }
          }
        });
      });
    },
    refresh() {
      const task = this.data.task;
      if (!task || this.evalRunning) return;
      const queryParams = {
        task_id: task.id,
        filter: this.dataset,
        page: this.currentPage,
        pagesize: this.pageSize,
        result: this.resultType
      }
      getAiEvalDetailResult(queryParams).then(res => {
        res = res.data;
        console.log(res)
        if (res.code == 0 && res.data) {
          this.totalNum = res.count
          this.tableData = res.data.map((item) => {
            console.log(JSON.parse(item.raw_input))
            let inputBlock = `\`\`\`json\n${item.raw_input}\n\`\`\``;
            return {
              ...item,
              rawContent: md.render(inputBlock),
              predictContent: md.render(item.predict),
              predContent: md.render(item.pred),
              rawExpanded: false,
              predictExpanded: false,
              predExpanded: false,
              rawNeedsExpand: false,
              predictNeedsExpand: false,
              predNeedsExpand: false
            }
          });
          this.checkContentHeights();
        } else {
          
        }
      }).catch(err => {
        this.$message.error(err)
        console.log(err);
      });
    },
  },
  created() {
    md = createRenderer()
  },
  mounted() {
    let selectDatasetTemp = [];
    const datasetParam = this.data.task.parameters.parameter.find(item => item.label === 'datasets');
    if (datasetParam && datasetParam.value) {
      selectDatasetTemp = datasetParam.value.split(' ')
        .sort((a, b) => a.localeCompare(b))
        .map(item => ({ k: item, v: item }));
    }
    this.selectDataset = [{ k: this.$t('modelFinetune.allDataset'), v: 'all' }, ...selectDatasetTemp]
    // if ([ "STOPPED","FAILED","START_FAILED","COMPLETED","SUCCEEDED","CREATED_FAILED"].includes(this.data?.task?.status)) {
    //   this.evalRunning = false
    // }
  }
};
</script>

<style scoped lang="less">
.el-table /deep/ .el-table__header th {
  background: rgb(245, 245, 246);
  color: #101010;
}
.wait-wrap{
  min-height: 400px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.chart-container {
  width: 100%;
  height: 100%;
  margin-top: 12px;
  .tips{
    line-height: 20px;
    color: rgba(16,16,16,0.5);
    font-size: 14px;
  }
  
  .ellipsis-text {
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .show-more-btn {
    color: #409EFF;
    font-size: 12px;
    cursor: pointer;
    margin-top: 4px;
    display: inline-block;
    &:hover {
      text-decoration: underline;
    }
  }
  
  .table-wrap {
    height: calc(100% - 40px);
    overflow: auto;
  }
}
@keyframes rotation {
  from {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(360deg);
  }
}
.rotating {
  animation: rotation 4s linear infinite;
}
</style>