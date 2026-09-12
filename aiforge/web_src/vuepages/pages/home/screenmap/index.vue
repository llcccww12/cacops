<template>
    <div class="__home_screemap1">
        <button v-if="showFlag" @click="roamMap(0)" class="screemap_button_add">+</button>
        <button v-if="showFlag" @click="roamMap(1)" class="screemap_button_min">-</button>
        <div class="screemap_content_wrap" v-if="showData">
          <div class="screemap_content_part1">
            <div class="text_wrap_right">
              <span>{{ $t('handleTask') }}</span>
            </div>
            <div style="width: 100px;"></div>
            <div class="text_wrap_left">
              <span>{{ $t('freeCompute') }}</span>
            </div>
          </div>
          <div class="screemap_content_part2">
            <div class="inner-container">
              <div class="container_wrap">
                <div class="content_container" v-for="item in cardAndJobCountList">
                  <div style="width: 145px;">
                    <div :style="{width:item.percentTask + '%',padding:(item.num==0?'0px':'')}" class="left_task_num" :tooltip="item.num +'个'" position="left"></div>
                  </div>
                  <div style="width: 100px;text-align: center;white-space: nowrap;overflow: hidden;" :title="item.ai_center">{{item.ai_center.slice(0,6)}}</div>
                  <div style="width: 145px;">
                    <div :style="{width:item.percentDuration + '%',padding:(item.card_duration==0.00?'0px':'')}" class="right_card_duration" :tooltip="item.card_duration + '卡时'" position="right"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div id="main" style="width:100%;height:759px"></div>
    </div>
</template>
<script>
import { getAiCenterOverview,getAiLocation } from "~/apis/modules/homemap";
import * as echarts from 'echarts'
import 'echarts/extension/bmap/bmap';

let {AppSubUrl} = window.config
export default {
data() {
    return {
      
      options:{},
      myChart:'',
      CCRFTETW:[],
      C2NET:[],
      SuperCompute:[],  
      TaskLine:[],
      cardAndJobCountList:[],
      showFlag:false,
      showData:false

    }
},
methods: {
    getOption(){
      this.options = {
        tooltip: {
          trigger: 'item',
          // formatter: '{b}',
        },
      color: [
        '#5470c6',
        '#0191FF',
        '#fac858',
        '#ee6666',
        '#73c0de',
        '#3ba272',
        '#fc8452',
        '#9a60b4',
        '#ea7ccc'
      ],
      bmap: {
        center: [98.114129, 31.550339],
        zoom: 6,
        roam: 'move',
        mapStyle: {
          styleJson: [
            {
              featureType: 'background',
              elementType: 'all',
              stylers: {
                color: '#FFFFFF'
              }
            },
            {
              featureType: 'water',
              elementType: 'all',
              stylers: {
                color: '#DBEAFE'
              }
            },
            {
              featureType: 'road',
              elementType: 'all',
              stylers: {
                visibility: 'off'
              }
            },
            {
              featureType: 'poi',
              elementType: 'all',
              stylers: {
                visibility: 'off'
              }
            },
            {
              featureType: 'local',
              elementType: 'all',
              stylers: {
                color: 'off'
              }
            },
            {
              featureType: 'arterial',
              elementType: 'labels',
              stylers: {
                visibility: 'off'
              }
            },
            {
              featureType: 'boundary',
              elementType: 'geometry',
              stylers: {
                color: '#9bacf0'
              }
            },
            {
              featureType: 'boundary',
              elementType: 'geometry.fill',
              stylers: {
                color: '#C7D1FB'
              }
            },
            {
              featureType: 'building',
              elementType: 'all',
              stylers: {
                visibility: 'off'
              }
            },
            {
              featureType: 'label',
              elementType: 'all',
              stylers: {
                visibility: 'off'
              }
            }
          ]
        }
      },
      series: [
        {
          name: '超算中心',
          type: 'scatter',
          coordinateSystem: 'bmap',
          data: this.SuperCompute,
          symbolSize: function (val) {
            return val[2] / 7;
          },
          encode: {
            value: 2
          },
          label: {
            formatter: '{b}',
            position: 'right',
            show: false
          },
          tooltip:{
            formatter: (params)=>{
              const marker = `
              <div>${params.seriesName}</div>
              <span style=\"display:inline-block;margin-right:4px;border-radius:10px;width:10px;height:10px;background-color:rgb(251, 118, 123);\"></span>
              <span>${params.data.name}</span>
              `
              return marker
            }
          },
          emphasis: {
            label: {
              show: true
            }
          },
          itemStyle: {
            shadowBlur: 10,
            shadowColor: 'rgba(25, 100, 150, 0.5)',
            shadowOffsetY: 5,
            color: new echarts.graphic.RadialGradient(0.4, 0.3, 1, [
              {
                offset: 0,
                color: 'rgb(67, 235, 255)'
              },
              {
                offset: 1,
                color: 'rgb(0, 123, 255)'
              }
            ])
          },
          zlevel: 1
        },
        {
          name: '东数西算',
          type: 'scatter',
          coordinateSystem: 'bmap',
          data: this.CCRFTETW,
          symbolSize: function (val) {
            return val[2] / 7;
          },
          encode: {
            value: 2
          },
          label: {
            formatter: '{b}',
            position: 'right',
            show: false
          },
          tooltip:{
            formatter: (params)=>{
              const marker = `
              <div>${params.seriesName}</div>
              <span style=\"display:inline-block;margin-right:4px;border-radius:10px;width:10px;height:10px;background-color:rgb(251, 118, 123);\"></span>
              <span>${params.data.name}</span>
              `
              return marker
            }
          },
          emphasis: {
            label: {
              show: true
            }
          },
          itemStyle: {
            shadowBlur: 10,
            shadowColor: 'rgba(25, 100, 150, 0.5)',
            shadowOffsetY: 5,
            color: new echarts.graphic.RadialGradient(0.4, 0.3, 1, [
              {
                offset: 0,
                color: 'rgb(217, 1, 217)'
              },
              {
                offset: 1,
                color: 'rgb(99, 1, 237)'
              }
            ])
          },
          zlevel: 3
        },
        {
          name: '智算中心',
          type: 'scatter',
          coordinateSystem: 'bmap',
          data: this.C2NET,
          symbolSize: function (val) {
            return val[2] / 7;
          },
          encode: {
            value: 2
          },
          label: {
            formatter: function(datas){
              if(datas.dataIndex<10){
                return datas.name
              }else{
                return ''
              }
            },
            position: 'right',
            show: true
          },
          tooltip:{
            formatter: (params)=>{
              const marker = `
              <div>${params.seriesName}</div>
              <span style=\"display:inline-block;margin-right:4px;border-radius:10px;width:10px;height:10px;background-color:rgb(251, 118, 123);\"></span>
              <span>${params.data.name}</span>
              `
              return marker
            }
          },
          showEffectOn: 'render',
          rippleEffect: {
            brushType: 'stroke'
          },

          itemStyle: {
            shadowBlur: 10,
            shadowColor: 'rgba(120, 36, 50, 0.5)',
            shadowOffsetY: 5,
            color: new echarts.graphic.RadialGradient(0.4, 0.3, 1, [
              {
                offset: 0,
                color: 'rgb(251, 118, 123)'
              },
              {
                offset: 1,
                color: 'rgb(204, 46, 72)'
              }
            ])
          },
          emphasis: {
            scale: true
          },
          zlevel: 4
        },
        {
          name: '实时调度任务',
          type: 'lines',
          coordinateSystem: 'bmap',
          zlevel: 2,
          effect: {
            show: true,
            symbol:'rect',
            period: 2.5, // 速度
            trailLength: 0.5, // 特效拖尾
            color: '#007BFF',
            symbolSize: 3,
          },
          lineStyle: {
            color: 'rgba(255,255,255,0.1)',
            width: 0,
            curveness: 0.2
          },
          data: this.TaskLine
        }
      ]
    }   
    this.options && this.myChart.setOption(this.options)
  },
  roamMap(flag){
    let currentZoom = this.myChart.getOption().bmap[0].zoom; // 当前的缩放比例   
    if (flag == 1) {
      currentZoom = currentZoom - 1
    }else{
      currentZoom = currentZoom + 1
    }
    if(currentZoom===4 && flag===0){
      this.myChart.setOption({
        bmap: {
            zoom: 5
        }
      })
    }else{
      this.myChart.setOption({
        bmap: {
            zoom: currentZoom
        }
      })
    }
    
  },
  async getAiCenterInfo(){
    const allPromise = Promise.allSettled([getAiCenterOverview(),getAiLocation()])
    try{
      const reslut = await allPromise;
      this.showFlag=true
      const [AiCenterInfo,localtionInfo] = reslut
      if(AiCenterInfo.status ==='fulfilled'){
        const GeomapData = AiCenterInfo.value.data.locationInfo
        this.CCRFTETW = GeomapData['东数西算'].map((item)=>{
          return {name:item.name,value:[Number(item.longitude),Number(item.latitude),item.value]}
        })
        this.C2NET = GeomapData['智算中心'].map((item)=>{
          return {name:item.name,value:[Number(item.longitude),Number(item.latitude),item.value]}
        })
        this.SuperCompute = GeomapData['超算中心'].map((item)=>{
          return {name:item.name,value:[Number(item.longitude),Number(item.latitude),item.value]}
        })
        const cardAndJobCount = AiCenterInfo.value.data.cardAndJobCount
        if(cardAndJobCount.length!==0){
          this.showData = true
          const cardAndJobCountSort = cardAndJobCount.map(({num})=>{return Number(num)})
          let sumDuration = Number(cardAndJobCount[0].card_duration)
          let sumTask = Math.max(...cardAndJobCountSort)
          this.cardAndJobCountList = cardAndJobCount.map((item)=>{
            const obj = {percentDuration:(Number(item.card_duration)/sumDuration).toFixed(2)*100,percentTask:(Number(item.num)/sumTask).toFixed(2)*100}
            item.card_duration = (Number(item.card_duration) / 3600).toFixed(2)
            return {...item,...obj}
          })
        }
        
      }else{

      }
      if(localtionInfo.status ==='fulfilled'){
        const schedulingTasks = localtionInfo.value.data
        if(schedulingTasks.length===0){
          this.TaskLine = []
        }else{
          this.TaskLine = schedulingTasks.map((item)=>{
            return {coords:[[Number(item.from_longitude),Number(item.from_latitude)],[Number(item.to_longitude),Number(item.to_latitude)]]}
          })
        }
      }else{

      }
      this.getOption()
    }catch(error){
      console.log(error)
    }
  },
  resizeEcharts(){
    this.myChart.resize();
  }
},
computed: {


},
beforeDestroy() {
  window.removeEventListener("resize", this.resizeEcharts);
},
mounted() {
  const chartDom = document.getElementById('main');
  this.myChart = echarts.init(chartDom);
  this.getAiCenterInfo()
  window.addEventListener('resize', this.resizeEcharts)
},
};
</script>
<style scoped lang="less">

</style>