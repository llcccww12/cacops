<template>
   <div class="content">
    <div class="head-wrap">
        <div class="title">LoRA {{taskName}}</div>
        <div class="title-sup">
            <span>{{$t('modelFinetune.trainProgress')}}</span>
            <div class="train-text">{{ trainStatusText }}</div>
        </div>
        <div class="progress-wrap">
            <div class="progress-title">
                <span>{{$t('modelFinetune.currentSteps')}}: {{ currentStepsNum }}/{{ stepsNum }}</span>&nbsp;&nbsp;&nbsp;
                <span>{{$t('modelFinetune.trainingRounds')}}: {{ currentEpochNum }}/{{ epochNum }}</span>
            </div>
            <div>
                <el-progress :stroke-width="10" :percentage="0" class="load-progress" v-if="activateFlag===0"></el-progress>
                <el-progress v-else :stroke-width="10" :percentage="percentSum"></el-progress>
            </div>
        </div>
    </div>
    <div class="main">
        <div class="left-wrap">
            <div class="btn-group">
                <span class="title">LoRA</span>
                <div class="btn-right">
                    <div class="params-wrap" @click="dialogShow=true">
                        <span class="title1">{{$t('modelFinetune.trainParameters')}}</span>
                        <i class="ri-file-copy-line"></i>
                    </div>
                    <div class="log-wrap">
                        <el-switch v-model="showLogFlag" active-color="#3162ff"></el-switch>
                        <span style="margin-left: 6px;">{{$t('modelFinetune.viewLoss')}}</span>
                    </div>
                </div>

            </div>
            <ul class="img-wrap" v-show="!showLogFlag">
                <section v-for="(item,index) in fileList" :key="index" class="imgs-box-wrap">
                    <div class="thum-img" v-if="item.img" :class="{ 'active': index===imgIndex }" @click="selectModel(item,index)">
                        <img v-if="item.img" :src="item.img" alt="">
                        <div class="img-text">
                        <div class="model-name">
                                <span :title="item.model">{{ item.model }}</span>
                            </div>
                            
                        </div>
                        
                    </div>
                    <div class="no-img" v-else>
                        <el-progress type="circle" :percentage="item.percent"></el-progress>
                    </div>
                </section>
            </ul>
            <div class="btn-wrap" v-show="!showLogFlag">
                <span>{{$t('modelFinetune.selected')}}：{{ modelName }}</span>
                <el-button @click="goImgPath" size="medium" class="generate-btn" :disabled="!finishFlag">{{$t('modelFinetune.startGenerateImg')}}</el-button>
            </div>
            <div class="chart-wrap" v-show="showLogFlag">
                <LossChart :lossData="lossData"></LossChart>
            </div>
        </div>
        <div class="right-wrap">
            <div class="title">
                {{$t('modelFinetune.sampleImage')}}
            </div>
            <ul class="show-img-wrap">
                <div class="img-miss" v-if="!previewImage">
                    <div class="img-miss-box">
                        <img src="/img/model/imgLoad.png" alt="">
                        <span>{{$t('modelFinetune.sampleImageGenerate')}}</span>
                    </div>
                </div>
                <div v-else class="img-preview">
                    <img :src="previewImage" alt="">
                    <div class="img-text">
                        <p :title="prompts">{{ prompts }}</p>
                    </div>
                </div>

            </ul>
        </div>
    </div>
    <el-dialog 
        :visible.sync="dialogShow"
        custom-class="container"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        class="fixed-close-dialog">
        <div class="params-wrap">
            <div class="section-wrap" v-for="(item,index) in paramsList" :key="index">
                <span class="title">{{ item.k }}:</span>
                <span class="label">{{ item.v }}</span>
            </div>
        </div>
    </el-dialog>
   </div>
</template>
<script>

import LossChart from '../components/LossChart.vue';
import { getTrainLoraProcess, } from '~/apis/modules/llmchat';
import { Thread } from '~/pages/model/wenxin/constant.js'
export default {
   name: 'LoraTrain',
   components: {
    LossChart
   },
   mixins: [],
   props: {
    taskUrl: { type: String, default: "", },
    taskId: { type: String, default: "", },
    taskName: { type: String, default: "", },
   },
   data() {
     return {
        showLogFlag: false,
        epochNum: 0,
        stepsNum: 0,
        currentEpochNum: 0,
        currentStepsNum: 0, 
        lossData: [],
        fileList: [],
        previewImage: '',
        prompts: '',
        percentSum: 0,
        trainStatusText: '',
        dialogShow: false,
        paramsList: [],
        activateFlag: 0,
        imgIndex: -1,
        finishFlag: false,
        modelName: '',
        imgList: []
     }
   },
   computed: {
     
   },
   watch: {
    finishFlag(val) {
        if (val) {
            this.imgIndex = 0
            this.modelName = this.fileList[0].model
        }
    }
   },
   mounted() {
    this.getTrainProcess()
    this.trainStatusText = this.$t('modelFinetune.trainProcessing')
   },
   methods: {
    selectModel(item, index) {
        if (this.finishFlag) {
            this.imgIndex = index
            this.modelName = item.model
            this.previewImage = item.img
        }
    },
    goImgPath() {
        const data = {
            model: this.modelName,
            modelList: this.modelList,
            imgList: this.imgList,
            index: this.imgIndex
        }
        this.$emit('getModel',data)
    },
    async getTrainProcess() {
        try {
            const vm = this
            const thread = new Thread({
                start: () => {
                    
                    getTrainLoraProcess(vm.taskUrl).then((response => { 
                        const res = response.data
                        if (res.code === 200) {
                            if (res.data.training_status === 0) {
                                vm.trainStatusText = '训练中'
                                vm.activateFlag  = 1
                            }
                            vm.epochNum = res.data.total_epoch
                            vm.stepsNum = res.data.total_steps
                            let lossValue = res.data.step_loss.map((item, index) => [index, item])
                            vm.currentStepsNum = lossValue.length
                            vm.currentEpochNum = res.data.epoch_loss.length
                            vm.lossData = lossValue
                            vm.prompts = res.data.user_config.PromptsConfig.prompts
                            
                            vm.percentSum = Number(((vm.currentStepsNum*100)/vm.stepsNum).toFixed(2))
                            let save_every_n_epochs = res.data.user_config.TrainConfig.save_every_n_epochs
                            let imageNum = Math.ceil(res.data.total_epoch/res.data.user_config.TrainConfig.save_every_n_epochs)
                            let initImgList = Array.from({ length:imageNum }, (_, index) => {
                                let startEpoch = save_every_n_epochs * index
                                let endEpoch = Math.min((index + 1) * save_every_n_epochs - 1, vm.epochNum - 1);
                                let startStep = startEpoch * (vm.stepsNum / vm.epochNum); // 假设steps是均匀分布的
                                let endStep = (endEpoch + 1) * (vm.stepsNum / vm.epochNum)
                                let totalpercent = endStep - startStep;
                                let percent = Number(((vm.currentStepsNum - startStep)*100 / totalpercent).toFixed(2)) 
                                let percentNum =  percent < 0 ? 0 : percent
                                return {
                                    img: '', // 初始时img为空字符串
                                    percent: percentNum
                                }
                            });
                            
                            let sampleImgLength = res.data.sample_imgs.length

                            sampleImgLength && res.data.sample_imgs.forEach((item,index) => {
                                let imgUrl = vm.taskUrl + '/sample_image/' + item.split('.png')[0]
                                initImgList[index].img = imgUrl
                                initImgList[index].model = res.data.saved_models[index].split('/').pop().split('.')[0]
                                if (index === (sampleImgLength - 1)) {
                                    vm.previewImage = imgUrl
                                }
                            });
                            vm.fileList = initImgList

                            //参数列表整理
                            const TrainConfigList = vm.convertConfigToArray(res.data.user_config.TrainConfig)
                            const DatasetConfigList = vm.convertConfigToArray(res.data.user_config.DatasetConfig)
                            vm.paramsList = [...DatasetConfigList,...TrainConfigList]
                            if (res.data.training_status === 1 && sampleImgLength=== res.data.total_epoch) {
                                vm.modelList = res.data.saved_models.map((item) => item.split('/').pop().split('.')[0])
                                vm.imgList = res.data.sample_imgs.map((item) => {return vm.taskUrl + '/sample_image/' + item.split('/').pop().split('.')[0]})
                                vm.trainStatusText = '训练完成'
                                vm.activateFlag  = 1
                                vm.finishFlag = true
                                thread.stop()
                            }
                        } else {
                            vm.$message.error(res.msg)
                            thread.stop()
                        }
                        
                    })).catch((err) => { 
                        thread.stop()
                        console.log("11111111111111")
                        vm.$message.error(err)
                    })
                },
                stop: function () {
                },
                number: 0, //这里是轮询次数配置，不配置默认无线轮询
                time: 4000 //这里是轮询的时间 不配置默认 300ms
            })
            // 开始轮询
            thread.run();
        } catch (error) {
            this.$message.error(error)
        }
    },
    // 定义一个函数来转换对象为所需的数组格式
    convertConfigToArray(config) {
        return Object.keys(config).filter(key => key !== '') // 排除 resolution 键
            .map(key => ({ k: key, v: config[key] }));
    }
   }
};
</script>
<style lang='less' scoped>
.content{
    height: calc(100% - 40px);
    margin: 20px;
    border: 1px solid rgba(157,197,226,0.4);
    border-radius: 10px;
    background-color: rgba(255,255,255,1);
    display: flex;
    flex-direction: column;
    padding: 24px;
    .head-wrap{
        margin-bottom: 1.5rem;
        .title{
            height: 40px;
            line-height: 20px;
            font-weight: 600;
            font-size: 20px;
            text-align: left;
        }
        .title-sup{
            height: 40px;
            line-height: 20px;
            font-size: 18px;
            display: flex;
            align-items: center;
            .train-text{
                color: rgba(0,102,255,1);
                background-color: rgba(16,16,16,0.05);
                font-size: 12px;
                border: 1px solid rgba(157,197,226,0.4);
                border-radius: 20px;
                padding: 0px 14px;
                margin-left: 8px;
            }
        }
        .progress-wrap{
            .progress-title{
                color:#888;
                margin-bottom:6px;
            }
            /deep/ .el-progress-bar {
                margin-right: -75px;
                padding-right: 75px;
            }
            /deep/ .load-progress .el-progress-bar__inner{
                animation: progress 1.5s infinite linear;
                background: #409EFF;
                
            }
            @keyframes progress {
                to {
                    width: 100%
                }
            }
            .activate{
                transition: width 1s ease-in-out; /* 你可以根据需要调整动画效果 */
                box-shadow: 0 0 10px rgba(64, 158, 255, 0.5); /* 可选的流光阴影效果 */
                animation: stream 3s linear infinite; /* 根据需要调整动画时长和速度 */
                @keyframes stream {
                    0% {
                        transform: translateX(0);
                    }
                    100% {
                        transform: translateX(100%);
                    }
                }
            }
        }
    }
    .main{
        flex-grow: 1;
        display: grid;
        grid-template-columns: 2fr 1fr; /* 左侧3份，右侧1份，共4份 */
        column-gap: 48px; /* 设置列间的间隔为24px */
        width: 100%; /* 容器宽度设为100% */
        box-sizing: border-box; 
        overflow-y: auto;
        .left-wrap{
            display: flex;
            flex-direction: column;
            padding: 10px; /* 内边距设置 */
            box-sizing: border-box; /* 确保内边距不增加总宽度 */
            flex: 1;
            overflow-y: auto;
            .title{
                font-size: 16px;
                font-weight: 400;
                color: #101010;
                height: 24px;
                line-height: 24px;
            }
            .btn-group{
                display: flex;
                justify-content: space-between;
                align-items: center;
                .btn-right{
                    display: flex;
                    .params-wrap{
                        display: flex;
                        align-items: center;
                        color: #888888;
                        margin-right: 14px;
                        cursor: pointer;
                        // overflow-y: auto;
                        // max-height: 600px;
                        i{
                            font-size: 18px;
                        }
                    }
                    .log-wrap{
                        display: flex;
                        align-items: center;
                        color: #888888;
                    }
                }
            }
            .img-wrap{
                padding: 0;
                overflow-y: auto;
                gap: 24px;
                display: grid;
                grid-template-columns: repeat(4, 1fr); /* 三等分 */
                &::-webkit-scrollbar {
                    width: 0; /* 尝试隐藏滚动条，但效果可能因浏览器而异 */
                    background: transparent; /* 设置滚动条背景为透明 */
                }
                .imgs-box-wrap{
                    aspect-ratio: 1;
                    border-width: 1px;
                    border-radius: 8px;
                    
                    cursor: pointer;
                    .thum-img{
                        overflow: hidden;
                        width: 100%;
                        height: 100%;
                        position: relative;
                        background-color:rgba(0, 0, 0, 0.06);
                        border-radius: 8px;
                        &.active{
                            border:3px solid #0066FF;
                        }
                        .img-text{
                            position: absolute;
                            background: linear-gradient(180deg, transparent, rgba(0, 0, 0, .65));
                            bottom: 0;
                            left: 0;
                            box-sizing: border-box;
                            width: 100%;
                            padding: 12px;
                            transition: opacity .3s;
                            .model-name{
                                width: 100%;
                                display: flex;
                                justify-content: center;
                                span{
                                    color: #fff;
                                }
                            }
                        }
                        img{
                            width: 100%;
                            height: 100%;
                            object-fit: cover !important;
                        }
                    }
                    .no-img{
                        display: flex;
                        align-items: center;
                        justify-content: center;
                        width: 100%;
                        height: 100%;
                        border: 2px dashed #dfe1e7;
                        background: #f5f6fa;
                        border-radius: 12px;
                    }
                }
            }
            .btn-wrap{
              display: flex;
              justify-content: space-between;  
              margin-top: auto;
              .generate-btn{
                background: linear-gradient(74.52deg, rgba(0,102,255,1) 11.24%,rgba(255,66,244,0.99) 97.08%);
                color: #fff;
              }
            }
            .chart-wrap{
                flex: 1;
                width: 80%;
                height: 80%;
                max-height: 600px;
            }
            
        }
        .right-wrap{
            display: flex;
            flex-direction: column;
            padding: 10px;
            .title{
                font-size: 16px;
                font-weight: 400;
                color: #101010;
                height: 24px;
                line-height: 24px;
            }
            .show-img-wrap{
                border-radius: 10px;
                background-color: rgba(255,255,255,0.2);
                border: 1px solid rgba(157,197,226,0.4);
                padding: 20px;
                display: flex;
                align-items: center;
                justify-content: center;
                flex: 1;
                .img-miss{
                    display: flex;
                    align-items: center;
                    flex-direction: column;
                    justify-content: center;
                    height: 100%;
                    .img-miss-box{
                        display: flex;
                        flex-direction: column;
                        align-items: center;
                        span{
                            color: #888;
                            margin-top: 14px;
                            height: 100%;
                        }
                    }
                    
                }
                .img-preview{
                    position: relative;
                    width: 100%;
                    height: 100%;
                    img{
                        width: 100%;
                        height: 100%;
                    }
                    .img-text{
                        background: linear-gradient(180deg, transparent, rgba(0, 0, 0, .65));
                        padding-top: 12px;
                        padding-bottom: 4px;
                        padding-left: 8px;
                        padding-right: 8px;
                        z-index: 1;
                        right: 0;
                        left: 0;
                        bottom: 0;
                        position: absolute;
                        align-items: center;
                        display: flex;
                        p{
                            line-height: 16px;
                            color:#fff;
                            font-size: 12px;
                            flex: 1 1 0%;
                            display: -webkit-box;
                            overflow: hidden;
                            -webkit-box-orient: vertical;
                            text-overflow: ellipsis;
                            word-break: break-all;
                            -webkit-line-clamp: 1;
                            margin: 0;
                        }
                    }
                }
            }
            
        }
    }
}
/deep/ .container{
    width: 460px;
    // min-width: 800px;
    // max-height: 600px;
    background: #fff;
    border-radius: 16px;
    margin-top: 5vh;
    // overflow-y: auto;
    .el-dialog__body{
        padding: 0 28px 28px 28px;
        max-height: 600px;
        overflow: auto;
        &::-webkit-scrollbar {
            width: 0; /* 尝试隐藏滚动条，但效果可能因浏览器而异 */
            background: transparent; /* 设置滚动条背景为透明 */
        }
    }
    .el-dialog__header {
        display: flex;
        justify-content: space-between;
        padding: 18px 28px;
        border-bottom: 0;
    }
    .el-dialog__headerbtn {
        top: 14px;
        right: 14px;
        width: 24px;
        height: 24px;
    }
    .section-wrap{
        font-size: 16px;
        line-height: 32px;
        .title{
            width: 240px;
            display: inline-block;
            color: #8a8e99;
        }
        .label{
            color: rgba(0, 0, 0, 0.87);
        }
    }
    
}

</style>