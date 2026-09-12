<template>
<div style="height: 100%;">  
    <div class="__mobile-tip" >
        <img style="width: 132px;height: 97px;" src="/img/model/pc-view.png" alt="">
        <div style="margin-top: 2rem;">{{$t('useInPcWeb')}}</div>
    </div>
    <div class="ui container  __content-box">
        <headerModel :config="config" :modelName="modelName" />
        <div style="height: calc(100vh - 228px);width:100%;display: grid;gap: 10px;grid-template-columns:1.6fr 4fr;padding: 1rem 0;overflow: hidden;">
            <sd-menu @changeParams="changeParams" :config="parametersConfig" ref="childMenu"></sd-menu>
            <div style="position: relative;">
                <div class="main-code-body">
                    <div class="output-text">
                        <div class="output-content">
                            <div class="output-container">
                                <div v-if="showPrompt" class="out-prompt">
                                    <span>{{ showPrompt }}</span>
                                    <div class="prompt-time">{{ unixTime }}</div>
                                </div>
                                <template v-if="!showPrompt">
                                            <div style="padding: 20px 0 0 20px;">
                                                <div class="wel-wrap">{{$t('modelSquare.chatHeaderTips',{modelName:modelName})}}</div>
                                                <div class="example-wrap" @click="setInitPrompt">
                                                    <div style="margin-bottom: 12px;">样例（点击加入到对话框）：</div>
                                                    <div>提示词：{{promptExample}}</div>
                                                    <div>负向提示词：{{negPromptExample}}</div>
                                                </div>
                                            </div>
                                        </template>
                                <div class="output-img-wrap">
                                    <main class="text2img-main">
                                        <div class="img-list" v-if="showLoading">
                                            <div>
                                                <img src="/img/loading.svg" style="width: 100px;height:100px;">
                                            </div>
                                        </div>
                                        <div class="img-list" v-if="showResultFlag">
                                            <div class="img-show-item">
                                                    <div v-for="(item, index) in numImgArray" class="item" :style="{width:spliWidth + 'px'}">
                                                        <div class="gallery" style="width:100%">
                                                            <img :src="'data:image/Jpeg;base64,'+resultImg[index]" alt="" style="border-radius: 1rem;z-index: 9;width: 100%;height:100%;">
                                                        </div>
                                                </div>
                                            </div>
                                        </div>
                                    </main>
                                </div>
                            </div>
                        </div>
                    </div>
                    
                </div>
                <div class="main-code-input">
                    <div class="input-container">
                        <div class="input-box">
                            <div class="ask-input">
                                <textarea id="code-input" @keydown.enter="carriageReturn($event)" :placeholder="$t('modelSquare.sdPlaceholder')" style="height: auto;" v-model="inputParams"></textarea>
                                <button class="opera-submit" @click="submit()"></button>
                            </div>
                        </div>
                        <!-- <div class="input-tips" style="text-align: end;" v-if="inputParams">{{$t('modelSquare.submitTips')}}</div> -->
                    </div>
                </div>
            </div>
        </div>
        <div style="text-align: center;">
            <span> <a href="/home/model_privacy"  target="_blank">{{ $t('modelSquare.modelProvide') }}</a></span>
        </div>
    </div>
</div>
</template>
  
<script>

import { postSdPaintNew, postERNIEPaintResult } from "~/apis/modules/desensitization";
import { Thread } from '~/pages/model/wenxin/constant.js'
import dayjs from 'dayjs';
import headerModel from '../componenes/headerModel.vue'
import SdMenu from '../componenes/SdMenu.vue';
import { getPromoteData } from '~/apis/modules/common';
export default {
    components: { headerModel, SdMenu },
    data() {
        return {
            inputParams:'',
            waitTime: 0,
            showPrompt: '',
            unixTime: '',
            numImgArray: [],
            resultImg: '',
            spliWidth: 0,
            showResultFlag: false,
            showLoading: false,
            sdFlag: false,
            modelName: '',
            config: {},
            negPromptExample: '',
            promptExample: '',
            parametersConfig:{},

        }
    },
    methods: {

        setInitPrompt() {
            this.inputParams = this.promptExample
            this.parametersConfig.negative_prompt.default = this.negPromptExample
        },
        carriageReturn(event){
            event.preventDefault()
            if(event.ctrlKey && event.keyCode ==13){
                this.inputParams = this.inputParams + '\n'
            }else{
                this.submit()
            }
        },
        submit() {
            this.$refs.childMenu.sendPamras()
            
        },
        changeParams(params){
            if (!this.inputParams) {
                this.$message.error(this.$t('modelSquare.inputNotEmpty'))
                return
            }
            if(this.sdFlag){
              this.$message.error(this.$t('modelSquare.sessionSding'))
              return
            }
            let clientWidth = document.getElementsByClassName('text2img-main')[0].clientWidth
            this.spliWidth = clientWidth/2 -20
            const date = new Date()
            this.showPrompt = this.inputParams
            params.prompt = this.inputParams
            this.inputParams = ''
            this.unixTime = dayjs(date).format('YYYY-MM-DD HH:mm:ss');
            this.resultImg = []
            switch (params.num_images_per_prompt) {
                case 1:
                    this.numImgArray = new Array(1).fill(0)
                    break
                case 2:
                    this.numImgArray = new Array(2).fill(0)
                    break
                case 3:
                    this.numImgArray = new Array(3).fill(0)
                    break
                case 4:
                    this.numImgArray = new Array(4).fill(0)
                    break
                default:
                    this.numImgArray = new Array(1).fill(0)
                    break
            }
            this.showLoading = true
            this.sdFlag = true
            this.showResultFlag = false
            postSdPaintNew(params).then((res) => {
                if (res.data.result == -1) {
                    this.$message.error(res.data.msg)
                    this.showLoading = false
                    this.sdFlag = false
                    return
                } else {
                    let { result, wait } = res.data
                    this.waitTime = wait
                    this.main(result, wait,params.prompt)
                }
            }).catch((err) => { 
                this.showLoading = false
                this.sdFlag = false
                this.$message.error(err)
            })
        },
        main(id, count,prompt) { 
            let vm = this
            let countRequest=0
            const thread = new Thread({
                start: function () {
                    postERNIEPaintResult({ id: id }).then((res => { 
                        if (res.data.Status === 0) { 
                            if (res.data.Picture) {
                                vm.showResultFlag = true
                                vm.showLoading = false 
                                this.sdFlag = false
                                vm.resultImg = res.data.Picture.split(";")
                                thread.stop()
                            }
                        } else {
                            vm.$message.error(res.data.Picture||'发生错误')
                            vm.inputParams = prompt
                            vm.showPrompt = ''
                            thread.stop()
                        }
                        
                        countRequest++
                    })).catch((err) => { 
                        vm.showLoading = false
                        vm.sdFlag = false
                        vm.$message.error(err)
                    })
                },
                stop: function () {
                    vm.showLoading = false 
                    vm.sdFlag = false
                    if (countRequest === Math.ceil(count / 2) - 1) { 
                        vm.$message.error(vm.$t('modelSquare.requireTimeout'))
                    }
                },
                number: Math.ceil(Number(count)), //这里是轮询次数配置，不配置默认无线轮询
                time: 3000 //这里是轮询的时间 不配置默认 300ms
            })
            // 开始轮询
            thread.run();
            
        },
        getConfig(){
            getPromoteData('model/modelexperiencenew.json').then((res)=>{
                const data = JSON.parse(res.data);
                let DepModelInfo = data.filter((item) => {
                    return item.name === this.modelName
                })
                this.config = DepModelInfo[0]
                // this.inputParams = this.config.parameters.prompt.default
                this.promptExample = this.config.parameters.prompt.default
                this.negPromptExample = this.config.parameters.negative_prompt.default
                this.config.parameters.negative_prompt.default = ''
                this.parametersConfig = this.config.parameters
            }).catch((err)=>{
                this.$message.error(err)
            })
        },
    },
    mounted(){
        const urlParams = new URLSearchParams(location.search)
        if(urlParams.has('modelName')){
            this.modelName = urlParams.get('modelName')
        }
        this.getConfig()

    },
    
}
</script>
<style lang="less" scoped>
* {
    box-sizing: border-box;
}

.menu-area{
    border-radius: 10px;
    background-color: rgba(249,250,251,1);
    border: 1px solid rgba(16,16,16,0.15);
    color: rgba(16,16,16,1);
    padding: 6px 14px 20px 22px;
    overflow-y: auto;
    min-height: 400px;
    min-width: 368px;
    .menu-select-container{
        height: 100%;
        display: flex;
        flex-direction: column;
    }
    .text-model{
        font-size: 18px;
        font-weight: 700;
        margin-bottom: 12px;
    }
    .label-wrapper{
        line-height: 32px;
        padding-top: 12px;
        .picture-size-wrap{
            display: flex;
            align-items: center;
            justify-content: space-between;
            .picture-item-container{
                width: 60px;
                height: 75px;
                display: flex;
                flex-direction: column;
                align-content: space-between;
                justify-content: center;
                align-items: center;
                border-radius: 5px;
                border: 1px solid rgba(16,16,16,0.15);
                font-size: 12px;
                cursor: pointer;
                &.active{
                    border-color: #0191ff;
                    color:#0366d6
                }
                .pic-bg{
                    background-color: #e0e0e0;
                    margin-top: 12px;
                    text-align: center;
                }
                .active{
                    background-color: rgba(1,145,255,0.3);
                }
                .pg-text{
                    text-align: center;
                    height: 23px;
                    line-height: 23px;
                }
                .pic-item-1{
                    width: 36px;
                    height: 36px;
                }
                .pic-item-2{
                    width: 27px;
                    height: 36px;
                }
                .pic-item-3{
                    width: 36px;
                    height: 27px;
                }
                .pic-item-4{
                    width: 24px;
                    height: 36px;
                }
                .pic-item-5{
                    width: 36px;
                    height: 24px;
                }
            }
            
        }
        .picture-num-wrap{
            display: flex;
            .picture-num-item{
                border: 1px solid rgba(16,16,16,0.15);
                width: 60px;
                height: 30px;
                line-height: 30px;
                border-radius: 5px;
                background-color: rgba(255,255,255,1);
                color: rgba(0,0,0,1);
                font-size: 12px;
                text-align: center;
                cursor: pointer;
                margin-right: 12px;
                &.active{
                    background-color: rgba(255,255,255,1);
                    color: rgba(1,145,255,1);
                    border: 1px solid rgba(1,145,255,1);
                }
            }

        }
    }
}
.main-code-body{
    max-height: calc(100vh - 264px);
    display: flex;
    width: 100%;
    overflow-x: auto;
    overflow-y: hidden;
    height: 100%;
    position: relative;
    overflow: hidden;
    // height: calc(100vh - 82px);
    .output-text{
        overflow: auto;
        flex: 1 1;
        // margin-bottom: 148px;
        .output-content{
            height: 100%;
            font-size: 14px;
            padding-left: 32px;
            overflow: auto;
            word-break: break-all;
            overflow-y: auto;
            .output-container{
                height: calc(100vh - 384px);
                border: 1px solid rgba(229,231,235,1);
                border-radius: 10px;
                display: flex;
                flex-direction: column;
                overflow-y: auto;
                .wel-wrap{
                    width: 733px;
                    height: 63px;
                    padding: 12px;
                    display: flex;
                    align-items: center;
                    border-radius: 10px 10px 10px 0px;
                    background-color: rgba(249,250,251,1);
                    color: rgba(16,16,16,0.8);
                    font-size: 16px;
                    text-align: left;
                    font-family: Roboto;
                    border: 1px solid rgba(229,231,235,1);
                    margin-bottom: 20px;
                }
                .example-wrap{
                    width: 733px;
                    height: 253px;
                    line-height: 23px;
                    border-radius: 10px 10px 10px 0px;
                    background-color: rgba(249,250,251,1);
                    color: rgba(16,16,16,0.8);
                    font-size: 16px;
                    text-align: left;
                    font-family: Roboto;
                    border: 1px solid rgba(229,231,235,1);
                    padding: 20px;
                }
                .out-prompt{
                    font-size: 16px;
                    line-height: 23px;
                    color: rgba(16,16,16,1);
                    background-color: rgba(249,250,251,1);
                    padding: 12px 24px;
                    word-break: break-all;
                    border-bottom: 1px solid rgba(229,231,235,1);
                    .prompt-time{
                        color: rgba(130,130,130,1);
                        font-size: 12px;
                        line-height: 17px;
                        padding-top: 5px;
                    }
                }
                .output-img-wrap{
                    flex: 1;
                    display: flex;
                    .text2img-main{
                        margin: 24px;
                        flex: 1 1;
                        display: flex;
                        align-items: center;
                        justify-content: center;
                        .img-list{
                            display: flex;
                            align-items: center;
                            justify-content: center;
                            width: 100%;
                            .img-show-item{
                                display:flex;
                                flex-wrap: wrap;
                                justify-content: flex-start;
                                align-items: center;
                                overflow: auto;
                                .item:nth-child(2){
                                    margin-left: 20px;
                                }
                                .item:nth-child(4){
                                    margin-left: 20px;
                                }
                            }
                            
                        }
                    }
                }
            }
            .user-input-wrapper{
                margin-left: 60px;
                margin-bottom: 8px;
                display: flex;
                align-items: top;
                flex-direction: row-reverse;
                white-space: pre-line;
                .user-input{
                    display: inline-block;
                    background: rgba(25, 117, 255, .1);
                    color: #43436b;
                    border-radius: 6px;
                    padding: 8px;
                    line-height: 20px;
                }
            }
            .user-output-wrapper{
                margin-right: 60px;
                margin-bottom: 24px;
                padding: 10px;
                display: flex;
                align-items: top;
                .user-output{
                    display: inline-block;
                    background: antiquewhite;
                    color: #43436b;
                    border-radius: 6px;
                    padding: 8px;
                    line-height: 20px;
                    position: relative;
                    word-break: break-all;

                }
            }
        }
    }
    .ask-text-input{
        position: relative;
        height: 140px;
        background-color: blue;
    }
}
.main-code-input{
    bottom: 0;
    // display: flex;
    // justify-content: center;
    margin-bottom: 6px;
    position: absolute;
    width: 100%;
    z-index: 99;
    padding-left: 32px;
    .input-container{
        flex: 1 1;
        position: relative;
        width: 100%;
        .input-box{
            background: #fff;
            border: 1px solid #2468f2;
            border-radius: 6px;
            width: 100%;
            .ask-input{
                padding: 12px 12px 28px;
                position: relative;
                #code-input{
                    border: none;
                    color: #151b26;
                    font-family: PingFangSC-Regular;
                    font-size: 14px;
                    font-weight: 400;
                    line-height: 20px;
                    max-height: 120px;
                    min-height: 42px;
                    overflow-y: auto;
                    resize: none;
                    width: 100%;
                    outline: none;
                }
                .opera-submit{
                    background-image: url(data:image/svg+xml;base64,PHN2ZyBoZWlnaHQ9IjMyIiB3aWR0aD0iNDgiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHJhZGlhbEdyYWRpZW50IGlkPSJhIiBjeD0iMTUzLjc5JSIgY3k9IjEyMS41OTklIiBncmFkaWVudFRyYW5zZm9ybT0ibWF0cml4KC0uNjI4MjYgLS4zMzQ1MyAuMjIzMDIgLS45NDIzOCAyLjIzMyAyLjg3NikiIHI9IjMzMi4wMjElIj48c3RvcCBvZmZzZXQ9IjAiIHN0b3AtY29sb3I9IiNlODUzZjYiLz48c3RvcCBvZmZzZXQ9Ii42MzQiIHN0b3AtY29sb3I9IiMyMTY0ZWQiLz48c3RvcCBvZmZzZXQ9IjEiIHN0b3AtY29sb3I9IiMyMTY0ZWQiLz48L3JhZGlhbEdyYWRpZW50PjxsaW5lYXJHcmFkaWVudCBpZD0iYiIgeDE9Ii02MS40NTElIiB4Mj0iMTAwJSIgeTE9IjY1LjAzNiUiIHkyPSI1MCUiPjxzdG9wIG9mZnNldD0iMCIgc3RvcC1jb2xvcj0iI2ZmZiIgc3RvcC1vcGFjaXR5PSIwIi8+PHN0b3Agb2Zmc2V0PSIxIiBzdG9wLWNvbG9yPSIjZmZmIi8+PC9saW5lYXJHcmFkaWVudD48ZyBmaWxsPSJub25lIiBmaWxsLXJ1bGU9ImV2ZW5vZGQiPjxyZWN0IGZpbGw9InVybCgjYSkiIGhlaWdodD0iMzIiIHJ4PSI2IiB3aWR0aD0iNDgiLz48cGF0aCBkPSJNMS41MzQgMTQuMDlhLjUxMi41MTIgMCAwIDEtLjM2Ny0uMTU2LjU5Ny41OTcgMCAwIDEtLjEzMy0uNjEzbDEuNDIyLTQuMDkzYS41NTkuNTU5IDAgMCAxIC4yODMtLjMyMkw1LjYgNy41MzUgMi43NDMgNi4yYS41NjIuNTYyIDAgMCAxLS4yODctLjMyNkwxLjAzNCAxLjc3YS41OTcuNTk3IDAgMCAxIC4xMzMtLjYxNC41MDYuNTA2IDAgMCAxIC41OC0uMTA4TDE0LjU3NiA3LjAyYy4xOTUuMDkuMzIuMjk2LjMyLjUyNGEuNTcyLjU3MiAwIDAgMS0uMzIuNTIzTDEuNzQ4IDE0LjA0M2EuNTA0LjUwNCAwIDAgMS0uMjE0LjA0OHoiIGZpbGw9InVybCgjYikiIHRyYW5zZm9ybT0idHJhbnNsYXRlKDE2IDgpIi8+PC9nPjwvc3ZnPg==);
                    border: none;
                    border-radius: 6px;
                    bottom: 8px;
                    cursor: pointer;
                    height: 32px;
                    margin: 0;
                    outline: none;
                    padding: 0;
                    position: absolute;
                    right: 12px;
                    width: 48px;
                }
            }
            
        }
        .input-tips{
            color: #999;
            
            font-size: 12px;
            padding-top: 6px;
        }
    }
}
::v-deep .el-slider__input{
    width: 55px;
}
::v-deep .el-slider__runway.show-input{
    margin-right: 70px;
}
::v-deep .el-input__suffix{
    right: 0px;
}
::v-deep .el-input-number.is-without-controls .el-input__inner{
        padding-left: 10px;
        padding-right: 10px;
}

</style>