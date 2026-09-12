<template>
  <div v-if="!showErrorMsg" class="item-container">
    <div class="btn-wrap" @click="refresh">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="20" height="20"><defs></defs><g><path fill="#202565" d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"></path></g></svg>
      <span>{{$t('cloudbrainObj.refresh')}}</span>
    </div>
    <div>
      <el-select v-if="configs.multiNodes && multiNodesData.length > 1" v-model="nodeSel" @change="changeNode">
        <el-option v-for="(item, index) in multiNodesData" :key="item.id" :value="index"
          :label="`${$t('cloudbrainObj.computeNode')} ${index + 1}`"></el-option>
      </el-select>
    </div>
    <div class="chart-container" v-loading="loading">
      <div ref="chartRef" class="chart"></div>
    </div>
  </div>

  <div v-else class="error-message">
    {{ this.$t('noMessage') }}
  </div>
</template>

<script>
import { getAiTaskNodeInfo, getAiTaskResourceUseage } from '~/apis/modules/cloudbrain';
import * as echarts from "echarts";

const sortBy = (arr, k) => arr.concat().sort((a, b) => (a[k] > b[k] ? 1 : a[k] < b[k] ? -1 : 0));

const chartOptions = {
  legend: {
    width: '95%',
    data: [],
  },
  grid: {
    top: "60",
    bottom: "5%",
    left: '32',
    right: '80',
    x: "2%",
    containLabel: true,
  },
  tooltip: {
    trigger: "axis",
    backgroundColor: "rgb(51, 56, 84)",
    borderColor: "rgb(51, 51, 51)",
    borderWidth: 0,
    textStyle: {
      color: "#fff",
    },
    axisPointer: {
      type: "cross",
    },
    appendToBody: true,
  },
  xAxis: {
    type: "category",
    data: [],
    boundaryGap: false,
    axisLabel: {
      interval: "auto",
    },
    name: "",
  },
  yAxis: [{
    show: true,
    name: "(%)",
    position: 'left',
    axisLine: {
      show: true,
    },
    axisTick: { show: true },
  }, {
    show: false,
    name: "Value",
    position: 'right',
    axisLine: {
      show: true,
    },
    axisTick: { show: true },
  }],
  series: [],
};

let chartHandler;

export default {
  name: 'ResourceUseage',
  props: {
    configs: { type: Object, default: () => { return {} } },
    data: { type: Object, default: () => { return {} } },
  },
  data() {
    return {
      chartData: {},
      loading: false,
      nodeSel: 0,
      multiNodesData: [],
      showErrorMsg: false,
      errorMsg: '',  
    };
  },
  methods: {
    checkIsValueType(name) {
      return name.indexOf('Bytes') >= 0 || name.indexOf('Rate') >= 0 || name.indexOf('numProcesses') >= 0;
    },
    getChartData(useRefreshBtn) {
      const task = this.data.task;
      this.loading = true;
      getAiTaskResourceUseage({
        repoOwnerName: task.repoOwnerName,
        repoName: task.repoName,
        id: task.id,
        node_id: this.configs.multiNodes ? this.multiNodesData[this.nodeSel]?.id : undefined,
        log_file_name: this.configs.multiNodes ? this.multiNodesData[this.nodeSel]?.log_file_name : undefined,
      }).then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 0) {
          const data = res.data || {};
          const metricsInfo = data.metrics_info || [];
          let filterData = metricsInfo.filter((item) => {
            // return !["recvBytesRate", "diskWriteRate", "sendBytesRate", "diskReadRate",].includes(item.name);
            return (item.value && item.value.length);
          });
          filterData = sortBy(filterData, "name");
          const legenData = filterData.map((item) => {
            return item.name;
          });
          let valueTypeCount = 0;
          const seriesData = filterData.map((item) => {
            const value = (item.value || []).map((item) => {
              return item > 0 ? Number(Number(item).toFixed(3)) : "0";
            });
            const valueType = this.checkIsValueType(item.name);
            valueTypeCount += (valueType ? 1 : 0);
            const seriesOption = {
              name: item.name,
              type: "line",
              symbol: "circle",
              symbolSize: 10,
              smooth: true,
              showSymbol: false,
              yAxisIndex: valueType ? 1 : 0,
              lineStyle: {
                width: 2,
                shadowColor: "rgba(0,0,0,0.3)",
                shadowBlur: 10,
                shadowOffsetY: 8,
              },
              data: value,
            };
            return seriesOption;
          });
          const xLength = metricsInfo.length ? metricsInfo[0].value.length : 0;
          const xInterval = data.interval || 1;
          chartOptions.xAxis.data = Array.from(
            { length: xLength },
            (_, index) => index * xInterval
          );
          chartOptions.legend.data = legenData;
          const legendSelected = {};
          legenData.forEach(element => {
            const valueType = this.checkIsValueType(element);
            legendSelected[element] = !valueType;
          });
          if (valueTypeCount > 0) {
            chartOptions.yAxis[1].show = true;
          } else if (chartOptions.yAxis[1]) {
            chartOptions.yAxis.pop();
          }
          if (useRefreshBtn && chartHandler && Object.keys(legendSelected).join('') == Object.keys(chartOptions.legend.selected || {}).join('')) {
            try {
              const selected = chartHandler.getOption().legend[0].selected;
              chartOptions.legend.selected = selected;
            } catch { }
          } else {
            chartOptions.legend.selected = legendSelected;
          }
          chartOptions.series = seriesData;
          chartOptions.grid.top = '60';
          if (legenData.length == 0) {
            chartOptions.grid.top = '50';
          }
          if (legenData.length >= 8) {
            chartOptions.grid.top = '80';
          }
          if (legenData.length >= 16) {
            chartOptions.grid.top = '100';
          }
          chartHandler && chartHandler.dispose();
          chartHandler = echarts.init(this.$refs.chartRef);
          chartHandler.setOption(chartOptions);
        } else {
          this.showErrorMsg = true;
          this.errorMsg = res.msg || this.$t('common.unknownError');
          return;
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    refresh(useRefreshBtn) {
      const task = this.data.task;
      if (this.configs.multiNodes) {
        this.loading = true;
        getAiTaskNodeInfo({
          repoOwnerName: task.repoOwnerName,
          repoName: task.repoName,
          id: task.id,
        }).then(res => {
          res = res.data;
          if (res && res.code == 0) {
            this.multiNodesData = res.data?.nodes || [];
          }
          this.getChartData(useRefreshBtn);
        }).catch(err => {
          this.loading = false;
          console.log(err);
        });
      } else {
        this.getChartData(useRefreshBtn);
      }
    },
    changeNode() {
      this.getChartData();
    },
    resize() {
      chartHandler && chartHandler.resize();
    }
  },
  beforeMount() {
    chartOptions.xAxis.name = this.$t('cloudbrainObj.chartTime');
    chartOptions.yAxis[0].name = this.$t('cloudbrainObj.chartResourceUsage');
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
  position: relative;
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
  height: 400px;
  .chart {
    height: 100%;
  }
}
.error-message {
  padding: 30px 50px 24px 20px;
}
</style>
