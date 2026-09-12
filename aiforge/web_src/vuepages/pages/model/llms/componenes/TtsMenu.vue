<template>
   <div class="menu-area">
        <div class="menu-select-container">
            <div>
                <div class="text-select-container">
                    <div class="text-content">
                        <div v-for="key in Object.keys(config)" class="label-wrapper">
                            <template v-if="key==='lang'">
                                <div class="item-label">
                                    <span>{{$t('modelSquare.language')}}</span>
                                    <el-tooltip class="item" effect="dark" :content="config[key].desc" placement="right">
                                        <svg viewBox="64 64 896 896" focusable="false" data-icon="question-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M623.6 316.7C593.6 290.4 554 276 512 276s-81.6 14.5-111.6 40.7C369.2 344 352 380.7 352 420v7.6c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V420c0-44.1 43.1-80 96-80s96 35.9 96 80c0 31.1-22 59.6-56.1 72.7-21.2 8.1-39.2 22.3-52.1 40.9-13.1 19-19.9 41.8-19.9 64.9V620c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-22.7a48.3 48.3 0 0130.9-44.8c59-22.7 97.1-74.7 97.1-132.5.1-39.3-17.1-76-48.3-103.3zM472 732a40 40 0 1080 0 40 40 0 10-80 0z"></path></svg>
                                    </el-tooltip>
                                </div>
                                <div>
                                    <el-select v-model="params.lang" @change="changeLang" style="width:100%">
                                        <el-option v-for="(item, index) in config.lang.opt" :key="item.id" :value="item.id" :label="item.name"></el-option>
                                    </el-select>
                                </div>
                            </template>
                            <template v-if="key==='speakers'">
                                <div class="item-label">
                                    <span>{{$t('modelSquare.soundStyle')}}</span>
                                    <el-tooltip class="item" effect="dark" :content="config[key].desc" placement="right">
                                        <svg viewBox="64 64 896 896" focusable="false" data-icon="question-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M623.6 316.7C593.6 290.4 554 276 512 276s-81.6 14.5-111.6 40.7C369.2 344 352 380.7 352 420v7.6c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V420c0-44.1 43.1-80 96-80s96 35.9 96 80c0 31.1-22 59.6-56.1 72.7-21.2 8.1-39.2 22.3-52.1 40.9-13.1 19-19.9 41.8-19.9 64.9V620c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-22.7a48.3 48.3 0 0130.9-44.8c59-22.7 97.1-74.7 97.1-132.5.1-39.3-17.1-76-48.3-103.3zM472 732a40 40 0 1080 0 40 40 0 10-80 0z"></path></svg>
                                    </el-tooltip>
                                </div>
                                <div>
                                    <div class="tab-item-c">
                                        <div class="tab-item" :class="tabIndex == 0 ? 'tab-item-focus' : ''" @click="changeTab(0)">{{$t('modelSquare.systemDefault')}}</div>
                                        <div class="tab-item" :class="tabIndex == 1 ? 'tab-item-focus' : ''" @click="changeTab(1)">{{$t('modelSquare.uploadSound1')}}</div>
                                    </div>
                                    <div class="tab-content-c">
                                        <div v-show="tabIndex == 0" class="tab-content">
                                            <div v-for="(item, index) in config.speakers.opt" @click="changeSpeaker(item,index)"  :key="item.id" class="tab-content-item">
                                                <div class="speak-wrap">
                                                    <div class="speak-name" :class="{'active': index===speakIndex}">
                                                        {{item.name}}
                                                        <i v-if="index===speakIndex" class="ri-checkbox-circle-line" style="margin-left: 6px;"></i>    
                                                    </div>
                                                    <div class="speak-desc">{{item.desc}}</div>
                                                </div>
                                                <div @click="getVoice(item)">
                                                    <img src="/img/model/Play.svg" style="margin-right: 6px;width:22px;height:22px;" alt="">
                                                </div>                                                
                                            </div>
                                        </div>
                                        <div v-show="tabIndex == 1" class="tab-content">
                                            <div class="dropzone" id="myDropzone">
                                                <div class="am-text-success dz-message">
                                                    {{$t('modelSquare.dragDrop')}}<br>- {{$t('or')}} - <br>{{$t('clickUpload')}}
                                                </div>
                                            </div>
                                            
                                        </div>
                                    </div>
                                    <div v-show="tabIndex == 1"  class="upload-tips"><span>*</span>{{ $t('modelSquare.soundTemplateTips3') }}</div>
                                </div>
                            </template>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import { getSpeakerVoice} from '~/apis/modules/llmchat';
import Dropzone from 'dropzone';
import 'dropzone/dist/dropzone.css';
Dropzone.autoDiscover = false;
export default {
   name: 'SDMenu',
   props: {
    config:{type:Object,default:()=>({})},
   },
   
   data() {
     return {
        params: {
            lang: 'zh',
            speaker_id: '',
            text: '',
            task_id: 0,
            wav_data: '',
        },
        tabIndex: 0,
        dropzoneHandler: null,
        speakIndex: 0,
        maxFlag: false,
        typeFlag: false
     }
   },
   computed: {
     
   },
   watch: {
       config: {
        deep: true,
        handler(value) {
            if(JSON.stringify(value) !== '{}'){
                this.initParams()
                this.$nextTick(()=>{
                    this.initDropZone()
                })
            }
        }
       }
   },
   methods: {
    changeLang(value){
      this.params.lang = value
      this.$emit('changeLang',value)
    },
    changeTab(tab) {
      this.tabIndex = tab;
    },
    changeSpeaker(item,index){
        this.speakIndex = index
        this.params.speaker_id = item.id
    },
    sendPamras(){
        if(this.tabIndex == 0){
            this.params.wav_data = ''
            if(this.params.speaker_id = -1){
                this.params.speaker_id = this.config.speakers.opt[this.speakIndex].id
            }
        }else{
            this.params.speaker_id = -1
            let fileList = this.dropzoneHandler.getAcceptedFiles()
            this.params.wav_data = fileList.length > 0 ? fileList[0] : ''
        }
        this.$emit('changeParams',this.params);
    },
    getVoice(item){
        if(document.getElementsByTagName("audio").length){
            document.body.removeChild(document.getElementsByTagName("audio")[0])
        }
        const params = { speaker_id:item.id, task_id:this.params.task_id }
        getSpeakerVoice(params).then((res)=>{
            let blob = new Blob([res.data], { type: 'audio/wav' });
            let url = URL.createObjectURL(blob);
            const audio = document.createElement("audio"); //创建标签
            document.body.appendChild(audio); //将标签插入到body中
            audio.autoplay='autoplay'
            audio.style.display = "none";
            audio.src = url; // 指定链接
            // // 语音播放完毕后，需要手动释放内存
            audio.onended = function () {
                document.body.removeChild(audio);
                URL.revokeObjectURL(url);
            };
        }).catch((err)=>{
            let vm = this
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
    initParams(){
        let params = this.config
        for(let key in params){
           if(params[key] !== null  ){
            if(key==='speakers'){
                this.params['speaker_id'] = params[key]['opt'][0].id
            }else{
                this.params[key] = params[key]['opt'][0].id
            }
                
            }
        }
    },
    initDropZone(){
        let vm = this
        let previewTemplate = `
                <div class="dz-preview dz-file-preview" style="margin: 12px;"> 
                <div class="dz-image"> 
                    <img data-dz-thumbnail /> 
                </div> 
                <div class="dz-details"> 
                    <div class="dz-size"><span data-dz-size></span></div> 
                    <div class="dz-filename"><span data-dz-name></span></div> 
                </div> 

                <div style="opacity:0" class="dz-progress"><span class="dz-upload" data-dz-uploadprogress></span></div> 
                <div class="dz-error-message" style="line-height: 1.5;top:150px"><span data-dz-errormessage></span></div> 
                <div class="dz-success-mark"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="54" height="54"><path fill="none" d="M0 0h24v24H0z"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10-4.477 10-10 10zm0-2a8 8 0 1 0 0-16 8 8 0 0 0 0 16zm-.997-4L6.76 11.757l1.414-1.414 2.829 2.829 5.656-5.657 1.415 1.414L11.003 16z" fill="rgba(47,204,113,1)"/></svg></div> 
                <div class="dz-error-mark"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="54" height="54"><path fill="none" d="M0 0h24v24H0z"/><path d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10-4.477 10-10 10zm0-2a8 8 0 1 0 0-16 8 8 0 0 0 0 16zm0-9.414l2.828-2.829 1.415 1.415L13.414 12l2.829 2.828-1.415 1.415L12 13.414l-2.828 2.829-1.415-1.415L10.586 12 7.757 9.172l1.415-1.415L12 10.586z" fill="rgba(231,76,60,1)"/></svg></div> 
                </div> `;
        this.dropzoneHandler = new Dropzone('#myDropzone', {
            url: "/",
            autoProcessQueue:false,
            previewTemplate:previewTemplate,
            maxFiles:1,
            addRemoveLinks: true,
            accept: function(file, done) {
                if (file.type == 'audio/mpeg' || file.type == 'audio/wav' ) {
                    done()
                }
                else {
                    this.removeFile(file)
                    if(!vm.typeFlag){
                        vm.typeFlag = true
                        vm.$message({
                            message: vm.$t('modelSquare.soundTemplateTips'),
                            type: 'error',
                            onClose:()=>{
                                vm.typeFlag = false
                            }
                        })
                    }
                }
            },
            dictRemoveFile: this.$t('modelManage.removeFile'),

        });
        this.dropzoneHandler.on("maxfilesexceeded", file => {
            this.dropzoneHandler.removeFile(file)
            if(!vm.maxFlag){
                vm.maxFlag = true
                vm.$message({
                    message: vm.$t('modelSquare.soundTemplateTips1'),
                    type: 'error',
                    onClose:()=>{
                        vm.maxFlag = false
                    }
                })
            }
            
        });
    }
   },
   mounted() {
    
   },
   beforeMount() {
        const urlParams = new URLSearchParams(location.search)
        if(urlParams.has('id')){
            this.params.task_id = +atob(urlParams.get('id'))
        }
    },
  
};
</script>
<style lang='less' scoped>
.menu-area{
    border-radius: 10px;
    background-color: rgba(249,250,251,1);
    border: 1px solid rgba(16,16,16,0.15);
    color: rgba(16,16,16,1);
    padding: 26px 14px 20px 22px;
    overflow-y: auto;
    min-height: 400px;
    min-width: 368px;
    height: calc(100vh - 250px);
    .menu-select-container{
        display: flex;
        flex-direction: column;
    }
    .label-wrapper{
        line-height: 32px;
        padding-top: 12px;
        .item-label{
            display: flex;
            align-items: center;
            svg{
                color: rgb(186, 191, 199);
                margin-left: 6px;
            }
        }
        .tab-item-c {
            display: flex;
            .tab-item {
                border: 1px solid #DCDFE6;
                display: flex;
                justify-content: center;
                align-items: center;
                height: 36px;
                padding: 0 20px;
                cursor: pointer;
                border-right: none;
                color: rgb(96, 98, 102);
                background: rgb(245, 245, 246);
                &:first-child {
                    border-radius: 5px 0px 0px 0px;
                }
                &:last-child {
                    border-right: 1px solid #DCDFE6;
                    border-radius: 0px 5px 0px 0px;
                }
                &.tab-item-focus {
                    color: rgb(50, 145, 248);
                    background: rgb(255, 255, 255);
                    border-bottom: none;
                }
            }
        }
        .upload-tips{
            color: rgb(242, 113, 28);
            line-height: 2;
            font-size: 12px;
            margin-top: 6px;
        }
        .tab-content-c {
            margin-top: -1px;
            border: 1px solid #DCDFE6;
            border-radius: 0px 5px 5px 5px;
            padding: 10px 16px;
            background: #fff;
            .tab-content {
                min-height: 34px;
                
                .tab-content-item{
                    border-bottom: 1px solid rgba(225,227,230,1);
                    display: flex;
                    align-items: center;
                    justify-content: space-between;
                    cursor: pointer;
                    .speak-wrap{
                        height: 70px;
                        .active{
                            color: rgba(0,102,255,1);
                        }
                        .speak-name{
                            height: 40px;
                            line-height: 50px;
                            display: flex;
                        }
                        .speak-desc{
                            font-size: 12px;
                            color:#888;
                            line-height: 24px;
                        }
                    }
                }
                .dropzone{
                    width: 100%;
                    min-height: 135px;
                    text-align: center;
                    line-height: 22px;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    flex-wrap: wrap;
                    color: #888;
                    padding: 0;
                    border: none;
                }
            }
        }
    }
}

</style>