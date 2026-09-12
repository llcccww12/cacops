<template>

  <div class="chart-container" v-loading="loading">
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
      <div class="main-wrap">
        <div class="chart-table">
          <el-table :data="tableData" style="width: 100%;margin-top: 58px;">
            <el-table-column :prop="lang === 'zh-CN' ? 'class' : 'classen'"  :label="$t('modelFinetune.evalTaskCategory')" width="260"></el-table-column>
            <el-table-column prop="name" :label="$t('repos.dataset')" width="160"></el-table-column>
            <el-table-column prop="value" :label="$t('modelFinetune.evalScore')" min-width="72"></el-table-column>
          </el-table>
        </div>
        <div class="chart">
          <div ref="chartRef" class="chart"></div>
          <div class="chart-custom-element">
            <div class="rect"></div>
            <span>{{ modelName }}</span>
          </div>
        </div>
      </div>
      
    </template>
    
    
  </div>
</template>

<script>
import { getAiEvalResult } from '~/apis/modules/cloudbrain';
import * as echarts from "echarts";

const chartOptions = {
  xAxis: {
    type: 'category',
    axisLabel: {
      interval: 0,  // 强制显示所有标签，不省略
      show: true,
      fontSize: 10,
      formatter: function(value) {
        // 对长标签进行换行处理
        if (value.length > 7) {
          if (value.includes('-')) {
            return value.split('-').join('-\n');
          }
          // 没有连接符，从中间位置换行
          var mid = Math.ceil(value.length / 2);
          return value.substring(0, mid) + '\n' + value.substring(mid);
            return value.split('-').join('-\n');
        }
        return value;
      }
      // rotate: 45, // 旋转45度
      // margin: 10 // 增加边距
    },
    data: []
  },
  yAxis: {
    type: 'value',
    max: 1
  },
  series: [
    {
      data: [],
      type: 'bar',
      barMaxWidth: 10,
      itemStyle: {
        color: 'rgb(80, 135, 236)'  // 设置柱子的颜色
      },
      label: {
        show: true,
        position: 'top'
      }
    }
  ]
};

let chartHandler;
export default {
  name: 'Loss',
  props: {
    data: { type: Object, default: () => { return {} } },
  },
  data() {
    return {
      tableData: [],
      modelName: '',
      evalRunning: true ,
      lang: 'zh-CN',
      loading: false
    };
  },
  watch: {
    data: {
      deep: true,  // 深度观察，检测嵌套属性的变化
      handler(newVal) {
        if (newVal.task && ["STOPPED", "FAILED", "START_FAILED", "COMPLETED", "SUCCEEDED", "CREATED_FAILED"].includes(newVal.task.status)) {
          this.evalRunning = false;
        }
      },
    },
  },
  methods: {
    refresh() {
      const task = this.data.task;
      if (!task) return;
      this.loading = true
      getAiEvalResult({id: task.id}).then(res => {
        this.loading = false
        res = res.data;
        console.log((res))
        const tmpArray = []
        const keyArray = []
        const valueArray = []
        if(Array.isArray(res.result) && res.result.length > 0 ){
          res.result.forEach(item => {
              tmpArray.push({ ...item})
              keyArray.push(item.name)
              valueArray.push(item.value)
            });
            this.tableData = tmpArray
            chartOptions.series[0].data = valueArray
            chartOptions.xAxis.data = keyArray
            console.log(chartOptions)
            chartHandler && chartHandler.dispose();
            chartHandler = echarts.init(this.$refs.chartRef);
            chartHandler.setOption(chartOptions);
        }
        // if (res.code == 0) {
        //   if (res.status == "RUNNING") {
        //   } else {
            
            
        //   }
        // }
      }).catch(err => {
        this.loading = false
        this.$message.error(err)
        console.log(err);
      });
    },
    resize() {
      chartHandler && chartHandler.resize();
    }
  },
  
  mounted() {
    this.lang = document.querySelector('html').getAttribute('lang');

    this.modelName = this.data?.task?.app_name
    window.addEventListener('resize', this.resize);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.resize);
    chartHandler && chartHandler.dispose();
  },
};
</script>

<style scoped lang="less">
.el-table /deep/ .el-table__header th {
  background: rgb(245, 245, 246);
  color: #101010;
}
.chart-container {
  width: 100%;
  height: 100%;
  overflow: auto; /* 允许水平滚动 */
  margin-top: 12px;
  .tips{
    line-height: 20px;
    color: rgba(16,16,16,0.5);
    font-size: 14px;
  }
  .wait-wrap{
    min-height: 400px;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  .main-wrap{
    min-width: 1100px; /* 设置最小宽度，确保两个元素都能显示 */
    min-height: 400px;
    width: 100%;
    margin: 0 auto;
    display: flex;
    justify-content: center;
    gap: 60px;
    .chart-table {
      width: 500px;
    }

    .chart {
      min-height: 400px;
      width: 540px;
      position: relative; /* 父容器设为相对定位 */
      .chart-custom-element{
        position: absolute; /* 子元素绝对定位 */
        top: 10px;          /* 距离顶部 */
        left: 50%;          /* 水平居中 */
        transform: translateX(-50%); /* 修正居中偏移 */
        z-index: 10;        /* 确保在图表上方 */
        display: flex;
        align-items: center;
        .rect{
          width: 24px;
          height: 11px;
          border-radius: 3px;
          margin-right: 4px;
          background: rgb(80, 135, 236);
          display: inline-block;
        }
      }
    }
  }
  

  /* 中等屏幕调整 */
  @media (max-width: 1200px) {
    .chart-table {
      width: 45%;
    }
    .chart {
      width: 50%;
    }
  }

  /* 小屏幕调整（如平板） */
  @media (max-width: 768px) {
    flex-direction: column; /* 改为垂直排列 */
    .chart-table, .chart {
      width: 100%;
    }
    .chart {
      margin-top: 20px;
    }
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
