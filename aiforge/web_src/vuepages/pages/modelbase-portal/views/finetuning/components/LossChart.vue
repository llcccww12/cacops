<template>
<div ref="chartContainer" class="chart" style="width: 100%; height: 100%;"></div>
</template>
  
<script>

import * as echarts from "echarts";
import {debounce} from 'lodash/function';
const chartOptions = {
  grid: {
    left: '8%',   // 左侧留出足够的空间
    // right: '10%',  // 右侧留出足够的空间
    // top: '10%',    // 顶部留出足够的空间
    bottom: '10%', // 底部留出足够的空间
    containLabel: true, // 确保坐标轴标签和名称在 grid 区域内
  },
  xAxis: {
    type:'value',
    splitLine:false,
    name:'step',
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
    nameGap:50,
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
    lossData: {type: Array,default: () => []},
  },
  data() {
    return {
      resizeObserver: null,
    };
  },
  watch: {
    lossData(val) {
      if (val.length) {
        // this.initializeChart()
        // this.updataChart()
        this.loadChart()
      }
    }
  },
  methods: {
    resize() {
      chartHandler && chartHandler.resize();
    },
    initializeChart() {
      
      chartHandler && chartHandler?.dispose();
      this.disposeResizeObserver();
      chartHandler = echarts.init(this.$refs.chartContainer);
      // chartHandler.setOption(chartOptions);
      this.initResizeObserver()
      this.loadChart()
    },
    // 开启监视ResizeObserver
    loadResizeObserver() {
      this.resizeObserver?.observe(this.$refs.chartContainer);
    },
    // 销毁ResizeObserver
    disposeResizeObserver() {
      this.resizeObserver?.unobserve(this.$refs.chartContainer);
      this.resizeObserver?.disconnect()
    },
    // 加载表格
    loadChart() {
      chartOptions.series[0].data = this.lossData
      chartHandler?.setOption(chartOptions, { notMerge: true }); //设置为true时不会合并数据，而是重新刷新数据
      this.resizeObserver && chartHandler?.on('finished', () => this.loadResizeObserver());
    },
    initResizeObserver() {
      if (!chartHandler) return;

      const __resizeHandler = debounce(() => {
        chartHandler?.resize()
      }, 0)

      this.resizeObserver = new ResizeObserver((entries, observer) => {
        __resizeHandler();
      });
    },
  },
  mounted() {
    this.initializeChart()
    window.addEventListener('resize', this.resize);
  },
  beforeDestroy() {
    if (!chartHandler) return;
    window.removeEventListener('resize', this.resize);
    chartHandler && chartHandler.dispose();
    this.disposeResizeObserver()
    this.resizeObserver = null
    chartHandler=null
  },
};
</script>
