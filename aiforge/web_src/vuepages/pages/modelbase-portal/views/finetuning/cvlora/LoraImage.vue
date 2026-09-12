<template>
<div class="content" v-loading="loading">
    <div class="base-model-settle">
        <div class="model-wrap">
            <img class="model-item" :src="`/img/model/${modelName}.png`"/>
        </div>
        <button class="start-button" :class="{ 'is-disabled': hasCaption }" @click="submitTask">
            <i class="ri-google-play-line" style="font-size: 24px;margin-bottom: 6px;"></i>
            <span>{{$t('modelFinetune.trainStart')}}</span>
        </button>
    </div>
    <div class="train-wrap">
        <div class="train-inner">
            <LoraParams ref="paramsRef" :fileLength="fileLength" :sections="sections" :modelName="modelName"></LoraParams>
            <div class="train-pic" v-loading="uploadLoadFlag" :element-loading-text="loadingText">
                <div class="top-tool">
                    <div class="number">2</div>
                    <span>{{$t('modelFinetune.imageLabeling')}}</span>
                    <div class="img-nums">
                        {{$t('modelFinetune.totalPic')}}<span style="color: #247CFF;"> {{fileLength}} </span>{{$t('modelFinetune.numPic')}}
                    </div>
                    <a v-if="fileLength>0" class="clear-imgs" @click="handleDeleteAll">
                        {{$t('modelFinetune.clearAllPic')}}
                    </a>
                    <span class="use-ex-data" @click="useExampleData">使用示例数据集</span>
                </div>
                <UploadDrag ref="uploadDarg" @uploadSuccess="uploadSuccess" :taskUrl="taskUrl" 
                :isVisible="isVisible" @start-loading="startLoading" @stop-loading="stopLoading"></UploadDrag>
                <ul class="grid-images" :style="{ display: isVisible ? 'none' : 'grid' }">
                    <div class="imgs-box-wrap upload-wrap" @click="uploadImgClick">
                        <div class="upload-item">
                            <i class="ri-add-line"></i>
                            <p class="label">{{$t('modelFinetune.addPic')}}</p>
                        </div>
                    </div>
                    <section v-for="(item,index) in fileList" :key="item.uuid" class="imgs-box-wrap imgs-wrap">
                        <p class="img-resolution">{{item.width}}x{{ item.height }}</p>
                        <span class="img-del" @click="handleDelete(item.uuid)"><i class="ri-delete-bin-2-line"></i></span>
                        <div class="thum-img" @click="showCaptionDialog(index)">
                            <img :src="item.image" alt="">
                        </div>
                        <div class="caption-text">
                            <p>{{ item.caption }}</p>
                        </div>
                    </section>
                </ul>
                <div style="flex-grow: 1;"></div>
                <div class="caption-wrap">
                    <div class="caption-left">
                        <div class="item">
                            <div class="label">{{$t('modelFinetune.captionMethod')}}</div>
                            <el-select v-model="captionValue" size="medium">
                                <el-option
                                v-for="item in captionOptions"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value">
                                </el-option>
                            </el-select>
                        </div>
                        <div class="item">
                            <div class="label">{{$t('modelFinetune.captionthres')}}</div>
                            <el-input-number size="medium" v-model="captionThreshold" controls-position="right" :min="0" :max="1" :step="0.05" :precision="2"></el-input-number>
                        </div>
                        <div class="item" style="flex:1;">
                            <div class="label">{{$t('modelFinetune.modelTrigger')}}</div>
                            <el-input size="medium" v-model="modelTrickValue" class="trigger-input" :placeholder="$t('modelFinetune.modelTriggerph')"></el-input>
                        </div>
                        <el-button size="medium" type="primary" class="caption-btn" @click="captionClick">{{$t('modelFinetune.captioning')}}</el-button>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <CaptionDialog v-if="fileLength" :fileList="fileList" :currentIndex="currentIndex" :visible.sync="dialogVisible" :taskUrl="taskUrl" @captionChange="captionChange"></CaptionDialog>
</div>
</template>
<script>
import UploadDrag from '../components/UploadDrag.vue'
import CaptionDialog from '../components/CaptionDialog.vue'
import LoraParams from '../components/LoraParams.vue'
import {postAllAcaption, postClearImagesId, postDeleteImagesId, postMakeCaption,
    getCaptionStatus, getAllCaptionText, getAllImages, postTrainLora, getExampleImages
} from '~/apis/modules/llmchat';
import { Thread } from '~/pages/model/wenxin/constant.js'
export default {
   name: 'LoraImage',
   components: { UploadDrag, CaptionDialog, LoraParams },
   mixins: [],
   props: {
      taskUrl: { type: String, default: "", },
      taskId: { type: String, default: "", },
      sections: { type: Array, default: () => [] },
      modelName: { type: String, default: "", },
   },
   data() {
     return {
        fileList: [],
        isVisible: true,
        dialogVisible:false,
        
        currentIndex: 0 ,
        captionValue: 'wd-swinv2-tagger-v3',
        captionOptions: [
            { label: 'wd-v1-4-vit-tagger-v2', value: 'wd-v1-4-vit-tagger-v2' },
            { label: 'wd-v1-4-swinv2-tagger-v2', value: 'wd-v1-4-swinv2-tagger-v2' },
            { label: 'wd-swinv2-tagger-v3', value: 'wd-swinv2-tagger-v3' },
            { label: 'wd-v1-4-convnext-tagger-v2', value: 'wd-v1-4-convnext-tagger-v2'}
         ],
        captionThreshold:0.3,
        uploadLoadFlag: false,
        loadingText: this.$t('modelFinetune.imgUploading'),
        modelTrickValue: '',
        loading: false
     }
   },
   computed:{
      fileLength() {
        return this.fileList.length;
       },
       hasCaption() {
        // 使用 some 方法检查是否至少有一个 caption 有值
            return !this.fileList.some(file => file.caption.trim() !== '');
        },
   },
   watch: {
    fileLength(newVal) {
      if (newVal === 0) {
        this.isVisible = true
      } else {
        this.isVisible = false
      }
    }, 
   },

   mounted() {
    this.getAllImagesId()
   },
   methods: {
    async useExampleData(){
        try {
            this.startLoading()
            await postClearImagesId(this.taskUrl)
            this.fileList = []
            const res = await getExampleImages(this.taskUrl)
            if(res.data.code == 200){
                if(res.data.data.length > 0){
                    res.data.data.forEach(item => {
                        let imgUrl = `${this.taskUrl}/get_image/${item.data.uuid}`
                        this.fileList.push({image:imgUrl,...item.data})
                    })
                }
            }
        } catch (error) {
            this.$message.error(error)
        } finally{
            this.stopLoading()
        }
    },
    startLoading() {
      this.uploadLoadFlag = true;
      this.loadingText = this.$t('modelFinetune.imgUploading')
    },
    stopLoading() {
      this.uploadLoadFlag = false;
    },
    showCaptionDialog(index) {
      this.currentIndex = index
      this.dialogVisible = true
    },
    async getAllImagesId() {
        try {
            const res = await getAllImages(this.taskUrl)
            if (res.data.code == 200) {
                res.data = res.data.data.reduce((acc, user) => {
                    acc.push({data:user})
                    return acc;
                }, []);
                this.$refs.uploadDarg.processItems(res,this.taskUrl)
                
            } else {
                // this.$messagee.error(res.data.msg)
            }
            this.uploadLoadFlag = false
        } catch (error) {
            this.uploadLoadFlag = false
            this.$message.error(error)
        }
    },
    uploadSuccess(files) {
        console.log(files)
        if (files.length) {
            this.fileList.push(...files)
        }
    },
    captionChange(files) {
        if (files.length) {
            this.fileList = files
        }
    },
    uploadImgClick() {
        this.$refs.uploadDarg.uploadClick()
    },
    handleDeleteAll() {
      this.$confirm(this.$t('modelFinetune.sureClearAllPic'),  this.$t('tips'), {
        confirmButtonText: this.$t('confirm1'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        // 用户点击了确定按钮
        this.clearAllImages();
      }).catch(() => {
        // 用户点击了取消按钮或关闭了对话框
       
      });
    },
    async clearAllImages() {
        try {
            const res = await postClearImagesId(this.taskUrl)
            if (res.data.code==200) {
                this.$message.success(this.$t('modelFinetune.clearAllPicSucc'))
                this.fileList = []
            } else {
                this.$messagee.error(res.data.msg || this.$t('modelFinetune.clearAllPicFail'))
            }
        } catch (error) {
            this.$message.error(error)
        }
    },
    
    handleDelete(uuid) {
      this.$confirm(this.$t('modelFinetune.sureClearOnePic'),  this.$t('tips'), {
        confirmButtonText: this.$t('confirm1'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false,
      }).then(() => {
        // 用户点击了确定按钮
        this.clearSingleIamge(uuid);
      }).catch(() => {
        // 用户点击了取消按钮或关闭了对话框

      });
    },
    async clearSingleIamge(uuid) {
        try {
            const res = await postDeleteImagesId(this.taskUrl, uuid);
            if (res.data.code == 200) {
                this.fileList = this.fileList.filter((item, i) => item.uuid !== uuid);
                this.$message.success(this.$t('modelFinetune.clearOnePicSucc'))
            } else {
                this.$messagee.error(res.data.msg || this.$t('modelFinetune.clearOnePicFail'))
            }
        } catch (error) {
            this.$message.error(error)
        }
    },
    async captionClick() {
        try {
            this.uploadLoadFlag = true
            this.loadingText = this.$t('modelFinetune.autoCaption')
            const response = await postMakeCaption(this.taskUrl, { model: this.captionValue })
            const res = response.data
            let vm = this
            if (res.code === 200) {
                
                const thread = new Thread({
                    start: function () {
                        getCaptionStatus(vm.taskUrl).then((async (response) => { 
                            const res = response.data
                            if (res.code === 200) {
                                if (res.msg === 'captions finished') {
                                    if (vm.modelTrickValue) {
                                        postAllAcaption(vm.taskUrl, { text: vm.modelTrickValue, append: false }).then(async (res) => {
                                            if (res.data.code === 200) {
                                                await vm.getAllCaptions()
                                                thread.stop()
                                                vm.modelTrickValue = ''
                                            }
                                        })
                                    } else {
                                        await vm.getAllCaptions()
                                        thread.stop()
                                    }
                                }
                            } else {
                                vm.$message.error(res.msg)
                                thread.stop()
                            }
                            
                        })).catch((err) => { 
                            vm.uploadLoadFlag = false
                            vm.$message.error(err)
                        })
                    },
                    stop: function () {
                        vm.uploadLoadFlag = false 
                        vm.sdFlag = false
                    },
                    number: 0, //这里是轮询次数配置，不配置默认无线轮询
                    time: 2000 //这里是轮询的时间 不配置默认 300ms
                })
                // 开始轮询
                thread.run();
            }
        } catch (error) {
            this.$message.error(error)
        }
    },
    async getAllCaptions() {
        try {
            const response = await getAllCaptionText(this.taskUrl);
            const res = response.data
            if (res.code === 200 && res.data.length > 0) {
                const mergeFileList =  this.fileList.map((item1) => {
                    const item2 = res.data.find(item => item.uuid === item1.uuid)
                    if (item2) {
                        item1.caption = item2.text
                    }
                    return item1
                })
                this.fileList = mergeFileList
            }
            
        } catch (error) {
            this.$message.error(error)
        }
    },
    async submitTask() {
        if(this.hasCaption) return
        try {
            this.loading = true
            let data = this.$refs.paramsRef.submitParams()
            
            const response = await postTrainLora({task_id:atob(this.taskId)},data);
            const res = response.data
            if (res.code === 404) {
                this.$message.error(this.$t('modelFinetune.cvTaskfaildTips'))
            } else {
                if (res.code === 200) {
                    this.$emit("changePage")
                }
                else {
                    this.$message.error(res.msg)
                }
            }
             
            this.loading = false
        } catch (error) {
            this.loading = false
            this.$message.error(error)
        }
    },
    
   }
};
</script>
<style lang='less' scoped>
.content{
    height: 100%;
    padding: 20px;
    // padding: 0 0 10px 0;
    display: flex;
    flex-direction: column;
    min-height: 750px;
    .base-model-settle{
        display: flex;
        min-height: 120px;
        align-items: center;
        justify-content: space-between;
        flex-direction: row;
        background-color: rgba(255,255,255,1);
        border: 1px solid rgba(157,197,226,0.4);
        border-radius: 10px;
        padding: 0 24px;
        .model-wrap{
            display: flex;
            height: 76px;
            flex-direction: row;
            align-items: center;
            justify-content: flex-start;
            .model-item{
                margin-right: 22px;
                position: relative;
                flex-shrink: 0;
                width: 76px;
                height: 76px;
                border-radius: 8px;
                cursor: pointer;
                overflow: hidden;
                background-repeat: no-repeat;
                background-position: 50%;
                background-size: cover;
            }
        }
        .start-button{
            width: 90px;
            height: 80px;
            background-color: #0066ff;
            color: rgba(255, 255, 255, 1);
            box-shadow: 0px 2px 10px 0px rgba(0,199,255,0.8);
            border-radius: 8px;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            cursor: pointer;
            border: none;
            &.is-disabled{
                cursor: not-allowed;
                opacity: 0.5;
                pointer-events: none;
            }
        }
    }
    .train-wrap{
        margin-top: 10px;
        // flex-grow: 1;
        background: #fff;
        height: calc(100% - 120px);
        .train-inner{
            display: flex;
            height: 100%;
            border-radius: 8px;
            .number{
                display: flex;
                align-items: center;
                justify-content: center;
                width: 20px;
                height: 20px;
                margin-right: 8px;
                background: #3162ff;
                border-radius: 10px;
                color: #fff;
                font-size: 14px
            }
            
            .train-pic{
                display: flex;
                flex-direction: column;
                flex-grow: 1;
                overflow-y: auto;
                padding: 15px;
                height: 100%;
                .top-tool{
                    display: flex;
                    flex-direction: row;
                    align-items: center;
                    margin-bottom: 16px;
                    .img-nums{
                        color: rgba(41,45,69,0.4);
                        font-size: 12px;
                        margin-left: 12px;
                    }
                    .clear-imgs{
                        text-decoration: underline;
                        font-size: 12px;
                        margin-left: 12px;
                    }
                    .use-ex-data{
                        margin-left: auto;
                        cursor: pointer;
                        color: #0066ff;
                    }
                }
                
                
                .grid-images{
                    padding: 0;
                    overflow-y: auto;
                    gap: 12px;
                    display: grid;
                    grid-auto-rows: max-content;
                    // flex: 1;
                    // max-height: 60%;
                    .imgs-box-wrap{
                        aspect-ratio: 1;
                        border-width: 1px;
                        border-radius: 8px;
                        background-color:rgba(0, 0, 0, 0.06);
                        border-color:rgba(0, 0, 0, 0.06);
                        cursor: pointer;
                    }
                    .upload-wrap{
                        width: 100%;
                        // height: 100%;
                        .upload-item{
                            color:rgba(0, 0, 0, .65);
                            gap: 10px;
                            justify-content: center;
                            align-items: center;
                            flex-direction: column;
                            display: flex;
                            width: 100%;
                            height: 100%;
                            flex: 1;
                            .label{
                                margin: 0;
                                font-size:14px;
                                line-height: 20px;
                            }
                            i{
                                font-size: 24px;
                            }
                        }
                    }
                    .imgs-wrap{
                        overflow: hidden;
                        justify-content: center;
                        align-items: center;
                        flex-direction: column;
                        display: flex;
                        position: relative;
                        .img-resolution{
                            line-height: 16px;
                            font-size: 12px;
                            padding: 4px;
                            color: #fff;
                            top: 0;
                            left: 0;
                            position: absolute;
                            z-index: 1;
                        }
                        .img-del{
                            cursor: pointer;
                            position: absolute;
                            right: 4px;
                            top: 4px;
                            font-size: 16px;
                            color: #fff;
                            display: none;
                            z-index: 1;
                        }
                        &:hover .img-del{
                            display: block;
                        }
                        .thum-img{
                            overflow: hidden;
                            width: 100%;
                            height: 100%;
                            position: relative;
                            img{
                                width: 100%;
                                height: 100%;
                                object-fit: cover !important;
                            }
                        }
                        .caption-text{
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
                @media (min-width: 1280px) {
                    .grid-images {
                        grid-template-columns: repeat(5, 1fr);
                    }
                }
                .caption-wrap{
                    display: flex;
                    align-items: center;
                    width: 100%;
                    background-color: #fff;
                    -webkit-backdrop-filter: blur(2px);
                    backdrop-filter: blur(2px);
                    padding-top: 15px;
                    .caption-left{
                        display: flex;
                        flex-grow: 1;
                        .item{
                            margin-right: 12px;
                            display: flex;
                            flex-direction: column;
                            align-items: flex-start;
                            .label{
                                color: #999;
                                font-family: PingFang SC;
                                font-size: 12px;
                                font-style: normal;
                                font-weight: 400;
                                line-height: normal;
                                margin-bottom: 4px;
                            }
                            .trigger-input{
                                flex-grow: 1;
                                min-width: 148px;
                            }
                        }
                        .caption-btn{
                            width: 100px;
                            align-self: flex-end;
                            background-color: #0066ff;
                        }
                    }
                    
                }
            }
        }
    }
}
</style>