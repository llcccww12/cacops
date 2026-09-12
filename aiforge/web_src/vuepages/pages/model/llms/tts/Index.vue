<template>
<div style="height: 100%;">  
    <div class="__mobile-tip" >
        <img style="width: 132px;height: 97px;" src="/img/model/pc-view.png" alt="">
        <div style="margin-top: 2rem;">{{$t('useInPcWeb')}}</div>
    </div>
    <div class="ui container __content-box">
        <headerModel :config="config" :modelName="modelName" />
        <div style="height: calc(100vh - 228px);width:100%;display: grid;gap: 10px;grid-template-columns:1fr 3fr;padding: 1rem 0;">
            <TtsMenu @changeParams="changeParams" @changeLang="changeLang" :config="parametersConfig" ref="childMenu"></TtsMenu>
            <div style="position: relative;">
                <div class="main-code-body">
                    <div class="output-text" >
                        <div class="output-content" ref="chatContainer" @scroll="onScroll">
                            <div class="output-welcome-banner">
                                <div>{{$t('modelSquare.chatHeaderTips',{modelName:modelName})}}</div>
                            </div>
                        
                            <div style="width:100%; position: relative;" >
                                <SessionWindow ref="sessionWindow"></SessionWindow>
                            </div>
                        </div>
                    </div> 
                </div>
                <div class="main-code-input">
                    <el-button type="text" @click="clearChat">{{$t('modelSquare.newChat')}}</el-button>
                    <div class="input-container">
                        <div class="input-box">
                            <div class="ask-input">
                                <textarea id="code-input" :maxlength="maxlength" :placeholder="$t('modelSquare.ttsPlaceholder')" style="height: auto;" v-model="prompt" @keydown.enter="carriageReturn($event)"></textarea>
                                <button class="opera-submit" @click="submit"></button>
                            </div>
                        </div>
                        <!-- <div class="input-tips" style="text-align: end;" v-if="prompt">{{$t('modelSquare.submitTips')}}</div> -->
                    </div>
                </div>
            </div>
        </div>
        <div style="text-align: center;">
            <span> <a href="/home/model_privacy"  target="_blank">{{$t('modelSquare.modelProvide')}}</a></span>
        </div>
    </div>
</div>
</template>
    
<script>

import headerModel from '../componenes/headerModel.vue'
import { getSynthesize} from '~/apis/modules/llmchat';
import SessionWindow from '../componenes/SessionWindow.vue'
import TtsMenu from '../componenes/TtsMenu.vue';
import { getPromoteData } from '~/apis/modules/common';

export default {
    components: { headerModel,SessionWindow, TtsMenu },
    data() {
        return {
            prompt:'',
            sessionRecordData: [],
            isAutoScroll: true,
            chatFlag: false,
            modelName: '',
            config: {},
            parametersConfig: {},
            maxlength: 82
        }
    },
    watch:{
      sessionRecordData(val){
        if (val != null){
          this.$refs.sessionWindow.setSessionRecord(val)
          this.$nextTick(() => {
            this.scrollBottom();
          })
        }
      },
      
    },
    methods: {
        clearChat(){
            if(this.chatFlag){
              this.$message.error(this.$t('modelSquare.sessionChating'))
              return
            }else{
                this.sessionRecordData = []
            }
            
        },
        carriageReturn(event){
            event.preventDefault()
            if(event.ctrlKey && event.keyCode ==13){
                this.prompt = this.prompt + '\n'
            }else{
                this.submit()
            }
        },
        changeLang(value){
            if(value== 'zh'){
                this.maxlength = 82
            }else{
                this.maxlength = 250
            }
        
        },
        async submit() {
            this.$refs.childMenu.sendPamras()
        },
        changeParams(params){
            const re = new RegExp("^[ ]+$")
            if(!this.prompt || re.test(this.prompt)){
              this.$message.error(this.$t('modelSquare.inputNotEmpty'))
              return
            }
            if(this.chatFlag){
              this.$message.error(this.$t('modelSquare.sessionChating'))
              return
            }
            if(params.speaker_id === -1 && !params.wav_data){
              this.$message.error(this.$t('modelSquare.uploadSound'))
              return
            }
            if(params.lang == 'zh' && params.text.length > 82){
              this.$message.error(this.$t('modelSquare.inputLength'))
              return
            }
            params.text = this.prompt.trim()
            this.prompt = ''
            this.sessionRecordData.push({"role":"user","content":params.text})
            this.sessionRecordData.push({"role": "tts","content": ''})
            this.chat(params)
        },
        chat(params){
            const fd = new FormData();
            fd.append('task_id', params.task_id)
            fd.append('text', params.text)
            fd.append('lang', params.lang)
            fd.append('speaker_id', params.speaker_id)
            fd.append('wav_data', params.wav_data)
            this.chatFlag = true
            let popData = this.sessionRecordData[this.sessionRecordData.length - 1];
            getSynthesize(fd).then((res)=>{         
                this.chatFlag = false
                let blob = new Blob([res.data], { type: 'audio/wav' });
                let url = URL.createObjectURL(blob);
                popData.content = url
            }).catch((err)=>{
                 let vm = this
                 vm.chatFlag = false
                 this.sessionRecordData.pop()
                if(err.response.data.type=='application/json'){
                    const reader  = new FileReader();  //创建一个FileReader实例
                    reader.readAsText(err.response.data, 'utf-8'); //读取文件,结果用字符串形式表示
                    reader.onload=function(){//读取完成后,**获取reader.result**
                        const  {error}  = JSON.parse(reader.result);
                        vm.$message.error(error); //弹出错误提示
                    }
                }else{
                    vm.$message.error(err.error)
                }
            })

        },
        scrollBottom() {
          this.$nextTick(() => {
            if (!this.isAutoScroll) return;
            const scrollDom = this.$refs.chatContainer;
            scrollDom.scrollTop =  scrollDom.scrollHeight;

            // animation(scrollDom, scrollDom.scrollHeight);
          })
        },
        onScroll() {
            const scrollDom = this.$refs.chatContainer;
            const scrollTop = scrollDom.scrollTop;
            const offsetHeight = scrollDom.offsetHeight;
            const scrollHeight = scrollDom.scrollHeight;
            // 当滚动到底部，设置 isAutoScroll 为 true
            if (scrollTop + offsetHeight + 1 >= scrollHeight) {
              this.isAutoScroll = true;
            } else {
              // 否则，用户正在手动滑动，设置为 false，停止自动滚动
              this.isAutoScroll = false;
            }
        },
        getConfig(){
            getPromoteData('model/modelexperiencenew.json').then((res)=>{
                const data = JSON.parse(res.data);
                let DepModelInfo = data.filter((item) => {
                    return item.name === this.modelName
                })
                this.config = DepModelInfo[0]
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
    beforeMount() {
        
    },
}
</script>
    
    
<style lang="less" scoped>
* {
    box-sizing: border-box;
}
.main-code-body{
    display: flex;
    width: 100%;
    overflow-x: auto;
    overflow-y: hidden;
    height: calc(100vh - 250px);
    position: relative;
    overflow: hidden;
    .output-text{
        overflow: auto;
        flex: 1 1;
        margin-bottom: 148px;
        .output-content{
            height: 100%;
            font-size: 14px;
            padding: 32px;
            padding-top: 0;
            overflow: auto;
            word-break: break-all;
            .output-welcome-banner{
                padding: 12px;
                margin-bottom: 20px;
                background-color: aliceblue;
                
            }
            .user-input-wrapper{
                margin-left: 60px;
                margin-bottom: 24px;
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
    margin-bottom: 12px;
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
            text-align: end;
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