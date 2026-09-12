<template>
<div class="item-container">
  <div class="btn-wrap" @click="refresh">
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="20" height="20"><defs></defs><g><path fill="#202565" d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"></path></g></svg>
    <span>{{$t('cloudbrainObj.refresh')}}</span>
  </div>
  <div class="chart-container" v-loading="loading">
    <div ref="chartRef" class="chart"></div>
  </div>
</div>
  
</template>

<script>
import { getAiTaskLoss } from '~/apis/modules/cloudbrain';
import * as echarts from "echarts";

const chartOptions = {
  xAxis: {
    type:'value',
    splitLine:false,
    name:'epoch',
    nameLocation:'center',
    nameGap:30,
    nameTextStyle:{
      fontSize:16,
      fontWeight:'bold'
    },
    axisLabel:{
      fontSize:14,
      fontWeight:'bold'
    }
  },
  yAxis: {
    type:'value',
    splitLine:false,
    name:'loss',
    nameLocation:'center',
    nameGap:30,
    nameTextStyle:{
      fontSize:16,
      fontWeight:'bold'
    },
    axisLabel:{
      fontSize:14,
      fontWeight:'bold'
    },
    scale:true,
    
  },
  
  series: [
    {
      data: [],
      type: 'line',
      showSymbol: false,
      smooth: 0.2
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
      lossValue: [],
      loading: false,
    };
  },
  methods: {
    refresh() {
      const task = this.data.task;
      if (!task) return;
      this.loading = true;
      getAiTaskLoss({
        repoOwnerName: task.repoOwnerName,
        repoName: task.repoName,
        id: task.id,
      }).then(res => {
        this.loading = false;
        res = res.data;
        const tmpArray = []
        if (res.code == 0 && res.data) {
          res.data.forEach(item => {
            tmpArray.push([item.epoch,item.loss])
          });
          this.lossValue = tmpArray
          chartOptions.series[0].data = tmpArray
          chartHandler && chartHandler.dispose();
          chartHandler = echarts.init(this.$refs.chartRef);
          chartHandler.setOption(chartOptions);
          // const events = res.data.events || [];
          // events.forEach(item => {
          //   item.timestampStr = item.timestamp ? formatDate(new Date(item.timestamp), 'yyyy/MM/dd HH:mm:ss') : '';
          // });
          // this.events = events;
        } else {
          this.$message.error(res.msg)
        }
      }).catch(err => {
        this.loading = false;
        this.$message.error(err)
        console.log(err);
      });
    },
    resize() {
      chartHandler && chartHandler.resize();
    }
  },
  mounted() {
    window.addEventListener('resize', this.resize);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.resize);
    chartHandler && chartHandler.dispose();
  },
};
</script>

<style scoped lang="less">
.item-container{
  padding-top: 40px;
  .btn-wrap{
    position: absolute;
    right: 10px;
    top: 10px;
    height: 30px;
    display: flex;
    align-items: center;
    padding: 0 10px;
    border-radius: 6px;
    background-color: rgba(16,16,16,0.1);
    color: rgba(32,37,101,0.9);
    cursor: pointer;
    z-index: 999;
  }
}
.chart-container {
  height: 500px;
  width: 100%;
  .chart {
    height: 100%;
  }
}
</style>
