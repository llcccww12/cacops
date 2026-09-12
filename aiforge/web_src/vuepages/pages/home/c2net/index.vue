<template>
    <div>
        <div class="__home_screemap">
          <div id="main"></div>
          <div class="left_box_wrap">
              <collapseMenuVue style="margin-right:10px">
                <div>
                    <p v-for="(item,index) in cpResource" @click="cpResourceClick(item,index)" :class="{'active':cpIndex==index}" :key="item.ComputeSource">{{item.ComputeSource}}</p>
                </div>
              </collapseMenuVue>
              <collapseMenuVue title="卡类型">
                <div>
                    <p v-for="(item,index) in cardType" @click="cardTypeClick(item,index)" :class="{'active':cardIndex==index}" :key="item">{{renderCardType(item)}}</p>
                </div>
              </collapseMenuVue>
          </div>
          <div class="middel_box_wrap">
            <div class="context-wrap">
              <div class="ui right icon label basic content" v-if="queryPamras.ComputeSourceList!=='所有'">
                  <span class="text">计算资源：<span>{{queryPamras.ComputeSourceList}}</span></span>
                  <i class="close icon" style="color:#007aff" @click="deleteComputeSource()"></i>
              </div>
              <div class="ui right icon label basic content" v-if="queryPamras.AccCardTypeList!=='所有'">
                  <span class="text">卡类型：<span>{{renderCardType(queryPamras.AccCardTypeList)}}</span></span>
                  <i class="close icon" style="color:#007aff"@click="deleteAccCardType()"> </i>
              </div>
            </div>
          </div>
          <div class="right_container">
            <div class="swiper-container" >
                <div class="swiper-wrapper" style="flex-direction: column;">
                    <div class="swiper-slide item" v-for="(item,index) in swiperList" :key="item.AICenterCode">
                        <div class="card-box" @click="selectedCard(item,index)" :class="{'active':selectedCardIndex==index}">
                            <div class="card-icon" v-if="index===0"><img src="/img/model/first.png"></div>
                            <div class="card-icon" v-else-if="index===1"><img src="/img/model/second.png"></div>
                            <div class="card-icon" v-else-if="index===2"><img src="/img/model/third.png"></div>
                            <div class="card-icon" v-else><img src="/img/model/common_awrad.png"></div>
                            <div style="overflow: hidden;">
                                <div>
                                    <div class="title">{{item.AICenterName}}</div>
                                    <div style="font-size:12px;">
                                        <span>城市：<span style="color: rgba(16,16,16,1);">{{item.City}}</span></span>
                                        <span>接入时间：<span style="color: rgba(16,16,16,1);">{{timeString(item.AccessTime)}}</span></span>
                                    </div>
                                </div>
                                <div class="label-wrap">
                                    <div class="label-item" v-for="item in item.ComputeSourceList">{{item || '--'}}</div>
                                </div>
                                <div class="label-wrap">
                                    <div class="label-item" v-for="item in item.AccCardTypeList" style="background:#e9e9e9;color:#101010">{{renderCardType(item)}}</div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        </div>
    </div>
</template>
<script>
import { getAiCenterCardInfo, getAiActivate } from "~/apis/modules/homemap";
import collapseMenuVue from './collapseMenu.vue'
import { handelData,sortData } from './utils'
import { ACC_CARD_TYPE } from '~/const'
import { getListValueWithKey } from '~/utils'
export default {
    data() {
        return {
           options: {},
            myChart: null,
            open: true,
            cpIndex: 0,
            cardIndex: 0,
            selectedCardIndex: -1,
            cpResource: [{ComputeSource:'所有'}],
            cardType: [],
            swiperHandler: null,
            data:[],
            aiCenterCount: 0,
            aiCenterList: [],
            swiperList: [],
            queryPamras:{
                ComputeSourceList:'所有',
                AccCardTypeList: '所有',
            }
        }
    },
    components:{collapseMenuVue},
    methods: {
        timeString(timestamp){
            return new Date(timestamp).toISOString().slice(0,10)
        },
        renderCardType(cardType) {
            return getListValueWithKey(ACC_CARD_TYPE, cardType)
        },
        selectedCard(item,index){
            this.selectedCardIndex = index
            this.data = [{
                name:item.Province,
                value:item.ComputeScale ? item.ComputeScale : 1
            }]
            this.$nextTick(()=>{
                this.getOption()
            })

        },
        initCardInfo(){
          getAiCenterCardInfo().then((res)=>{
            const data = res.data
            if(data.code===0){
              const List = data.data.list
              const cardTempList  = []
              List.forEach(item => {
                this.cpResource.push(item)
                item.CardList.forEach(card => {
                  cardTempList.push(card)
                });
                item.CardList.unshift('所有')
                
              });
              this.cardType = cardTempList.sort(sortData)
              this.cardType.unshift('所有')
              this.cpResource[0].CardList = this.cardType
            }
          })
        },
        initAiCenter(){
            getAiActivate().then((res)=>{
                const data = res.data
                if(data.code===0){
                    this.aiCenterList = data.data.list
                    this.flitersListFunc(this.queryPamras)
                }
            })
        },
        flitersListFunc(params){
            const fliteredList = this.aiCenterList.filter((item) => {
                if(params.ComputeSourceList!=='所有'){
                    if(item.ComputeSourceList.indexOf(params.ComputeSourceList)===-1) return false;
                }
                if(params.AccCardTypeList!=='所有'){
                    if(item.AccCardTypeList.indexOf(params.AccCardTypeList)===-1) return false;
                }
                return true;
            });
            this.aiCenterCount = fliteredList.length
            const siwperFliteredList = fliteredList
            this.updateSwiper(siwperFliteredList)
            this.data = handelData(fliteredList)
            this.getOption()
        },
        updateSwiper(newData) {
            // 销毁旧的Swiper实例
            this.swiperHandler.destroy(true);

            // 更新数据源
            this.swiperList = newData;

            // 下一帧重新挂载Swiper实例
            this.$nextTick(() => {
                let loop = this.aiCenterCount <= 5 ? false : true
                this.initSwiper(loop)
            });
        },
        initSwiper(loop=true){
            this.swiperHandler = new Swiper(".swiper-container", {
                direction: "vertical",
                slidesPerView: 3,
                loop: loop,
                spaceBetween: 0,
                autoplay: {
                    delay: 2500,
                    disableOnInteraction: false,
                },
            });
        },
        deleteComputeSource(){
            this.cpIndex = 0
            this.cardIndex = 0
            this.selectedCardIndex = -1
            this.queryPamras.AccCardTypeList = '所有'
            this.queryPamras.ComputeSourceList = '所有'
            this.cardType = this.cpResource[this.cpIndex].CardList
            this.flitersListFunc(this.queryPamras)
        },
        deleteAccCardType(){
            this.cardIndex = 0
            this.selectedCardIndex = -1
            this.queryPamras.AccCardTypeList = '所有'
            this.flitersListFunc(this.queryPamras)
        },
        cpResourceClick(item,index){
            let selectItem = item.ComputeSource 
            this.cpIndex = index
            this.cardType = this.cpResource[this.cpIndex].CardList

            this.queryPamras.AccCardTypeList = '所有'
            this.cardIndex = 0
            this.selectedCardIndex = -1
            this.queryPamras.ComputeSourceList = selectItem
            this.flitersListFunc(this.queryPamras)
        },
        cardTypeClick(item,index){
            let selectItem = item
            this.cardIndex = index
            this.selectedCardIndex = -1
            this.queryPamras.AccCardTypeList = selectItem
            this.flitersListFunc(this.queryPamras)
        },
        getOption(){
            this.options = {
                 visualMap:[{
                    type:'piecewise',
                    showLabel:true,
                    textStyle:{
                      color:'#101010',
                      fontSize:10,
                    },
                    itemWidth: 12,
                    itemGap: 0,
                    itemSymbol:'rect',
                    pieces: [
                        {gt:0, lt: 100, label:'<100', color:'#C5ECFF'},
                        {gte:100, lt:200, label:">=100", color:'#98ECF6'},
                        {gte:200, lt:300, label:">=200", color:'#66E0EF'},
                        {gte:300, lt:400, label:">=300", color:'#65CEF5'},
                        {gte:400, lt:500, label:">=400", color:'#66B8FB'},
                        {gte:500, lt:1000, label:">=500", color:'#6E9FFD'},
                        {gte: 1000, label:">=1000", color:'#7980FF'},
                    ],
                    inverse: true,
                    top:'70%',
                    left:14,
                }],
                series: [
                    {
                        type: 'map',
                        zoom: 1.1,
                        map: 'china',
                        selectedMode:false,
                        label: {
                            show: true,
                            fontSize:10,
                        },
                        itemStyle:{
                            areaColor:'#d4e4ff',
                            borderColor:'#fff'
                        },
                        emphasis:{
                            label:{
                                color:'#fff',
                            },
                            itemStyle:{
                                areaColor:'#459dfd',
                            }
                        },
                        data: this.data,
                    }
                ],
            }   
            this.options && this.myChart.setOption(this.options)
            
        },
       
        resizeEcharts(){
            this.myChart.resize();
        },

    },
    computed: {
    },
    beforeDestroy() {
        window.removeEventListener("resize", this.resizeEcharts);
    },
    mounted() {
        const chartDom = document.getElementById('main');
        this.myChart = this.$echarts.init(chartDom);
        this.initCardInfo()
        this.initAiCenter()
        this.getOption()
        window.addEventListener('resize', this.resizeEcharts)
        this.initSwiper();
    },
    created(){
      
    }
};
</script>
<style scoped lang="less">
*{
    moz-user-select: -moz-none;
    -moz-user-select: none;
    -o-user-select:none;
    -khtml-user-select:none;
    -webkit-user-select:none;
    -ms-user-select:none;
    user-select:none;
}
.__home_screemap{
  height: 865px;
  width:100%;
  background: radial-gradient(farthest-side at 10.7% 7.3%, rgba(232, 228, 255, 1) 0, rgba(234, 232, 249, 0) 100%);
  position: relative;
  #main{
      height: 400px;
      width: 100%;
  }
  .left_box_wrap{
    position: absolute;
    // display: inline-flex;
    left: 14px;
    top: 24px
  }
  .middel_box_wrap{
    position:relative;
    padding: 0 12px;
    .context-wrap{
        display:flex;
        flex-wrap: wrap;
        .content{
            height:30px;
            display:flex;
            align-items: center;
            padding: 0 12px;
            margin-right:10px;
            font-size:12px;
            .text{
                color: #606266;
                font-size:12px;
                font-weight: 400;
                span{
                    color: rgba(16,16,16,1);
                }
            }
        }
    }
  }
  .right_container{
    .swiper-container {
        width: 390px;
        height: 492px;
        overflow: hidden;
        .item{
            height: auto;
            .active{
                border: 1px solid rgba(179,157,226,1) !important;
            }
            .card-box{
                height: 150px;
                // width: 383px;
                border-radius: 6px;
                background-color: rgba(253,253,253,1);
                color: rgba(16,16,16,1);
                font-size: 14px;
                box-shadow: 0px 5px 10px 0px rgba(157,197,226,0.2);
                border: 1px solid rgba(157,197,226,0.2);
                margin: 15px 15px 0 15px;
                padding: 18px;
                display: flex;
                box-sizing: border-box;

                .card-icon{
                    width: 40px;
                    height: 40px;
                    margin-right: 12px;
                }
                .title{
                        color: rgba(16,16,16,1);
                        font-size: 16px;
                        font-family: SourceHanSansSC;
                        font-weight: 600;
                        margin-bottom: 5px;
                }
                .label-wrap{
                    display: flex;
                    margin-top: 8px;
                    height: 24px;
                    overflow: hidden;
                    -webkit-mask: linear-gradient(to right, #fff calc(100% - 20px), transparent);
                    .label-item{
                        display: flex;
                        align-items: center;
                        justify-content: center;
                        border-radius: 4px;
                        margin-right: 8px;
                        background: #BAEEFF;
                        color: #101010;
                        font-size: 12px;
                        padding: 0 8px;
                        white-space: nowrap;
                        flex-wrap: wrap;
                    }
                }
            }
        }
    }
}
}

/deep/ .nav-bar-content{
    ::-webkit-scrollbar-track {
        background: transparent;
        border-radius: 0;
    }

    ::-webkit-scrollbar {
    -webkit-appearance: none;
        width: 6px;
        height: 6px;
    }

    ::-webkit-scrollbar-thumb {
        cursor: pointer;
        border-radius: 5px;
        background: rgba(0, 0, 0, 0.15);
        transition: color 0.2s ease;
    }
    display: inline-block;
    vertical-align: top;
    .header-wrap{
        line-height: 30px;
        background: linear-gradient(91.34deg, rgba(0,122,255,1) 3.31%,rgba(65,99,227,1) 99.31%);
        color: #fff;
        display: flex;
        justify-content: space-between;
        padding: 0 10px;
        cursor: pointer;
        border-radius: 5px;
        width: 120px;
        margin-bottom: 8px;
        .arrow {
            // font-size: 16px;
            transition: transform 0.2s ease;
        }
        .open {
            transform: rotate(90deg);
            transition: transform 0.2s ease;
        }
    }
    .content{
      background-color: rgba(249,250,251,1);
      max-height: 300px;
      overflow-y: auto;
      p{
          padding: 0 10px;
          line-height: 28px;
          margin-bottom: 0;
      }
      .active{
          background-color: rgba(238,241,254,1);
          color: rgba(0,122,255,1);
      }
    }
}
// .swiper-container {
//     width: 413px;
//     height: 860px;
//     overflow: hidden;
//     .item{height: auto;}
// } 
// .__home_screemap{
//   display: flex;
//   min-height: 880px;
//   min-width:1200px;
//   overflow: auto;
//   background: radial-gradient(farthest-side at 10.7% 7.3%, rgba(232, 228, 255, 1) 0, rgba(234, 232, 249, 0) 100%);
//   border-radius: 15px;
//   position: relative;
//   #main{
//         height: 616px;
//         width: 800px;
//         margin:auto;
//         flex: 1;
//   }
//   .left_box_wrap{
//     border-radius: 10px;
//     background-color: rgba(249,250,251,1);
//     border: 1px solid rgba(213,217,224,1);
//     padding: 15px;
//     width: 183px;
//     margin: 15px;
//   }
//   .card-box{
//     height: 153px;
//     width: 383px;
//     border-radius: 6px;
//     background-color: rgba(253,253,253,1);
//     color: rgba(16,16,16,1);
//     font-size: 14px;
//     box-shadow: 0px 5px 10px 0px rgba(157,197,226,0.2);
//     border: 1px solid rgba(157,197,226,0.2);
//     margin: 15px 15px 0 15px;

//   }
// }
// @media only screen and (min-width: 767.98px) and (max-width: 1199.98px){
//     .__home_screemap_wrap{
//         overflow-x: scroll;
//     }
// }
</style>
  