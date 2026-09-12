<template>
    <div :style="{ display: isVisible ? 'block' : 'none' }" style="height: 80%;">
        <div class="pic-upload" ref="pickerArea">
            <div class="pic-upload-drag">
                <div class="upload-icon"></div>
                <div class="upload-text-wrap">
                    <p class="upload-tips">{{$t('modelFinetune.dragUploadImg')}}</p>
                    <p class="accpet-tips">{{$t('modelFinetune.uploadMaxImg')}}</p>
                </div>
                <Upload ref="uploadBtn" @fileList="handleFileList"></Upload>
            </div>
        </div>
        <div class="segmente">
            <div class="line"></div>{{$t('or')}}<div class="line"></div>
        </div>
        <div class="zip-upload">
            <div class="upload-btn" @click="uploadZipClick">{{$t('modelFinetune.uploadHavedImg')}}</div>
            <span class="zip-tips">{{$t('modelFinetune.uploadImgTips')}}</span>
            <input ref="fileInput" type="file" accept=".zip" style="display: none;" @change="handleZipChange">
        </div>
    </div>
</template>
<script>
import Upload from './Upload.vue'
import { sdLorauploadImages, getUploadImagesId, sdLorauploadZip} from '~/apis/modules/llmchat';
export default {
   name: 'UploadDrag',
   components: { Upload },
   props: {
    taskUrl: { type: String, default: "", },
    isVisible: {type: Boolean, default: true}
   },
   data() {
     return {
        bindDrop: false,
        uuidList: [],
        uploadUrl: '',
     }
   },
   computed: {
     
   },
   watch: {
    taskUrl: {
        immediate: true,
        handler(newValue, oValue) {
            if (newValue) {
                this.uploadUrl = newValue
            }
        }
    }
   },
   mounted() {
    this.bindEvents()
   },
   methods: {
    handleFileList(fileList) {
        if (!fileList || fileList.length === 0) {
            this.$message.error(this.$t('modelFinetune.uploadImgTips1'));
            return;
        }
        if (fileList.length > 200) {
            this.$message.warning(this.$t('modelFinetune.uploadImgTips2'))
            fileList = fileList.slice(0, 20);
        }
        // 过滤不符合格式的文件（双重保险）
        const validFiles = Array.from(fileList).filter(file => {
            const fileType = file.name.split('.').pop().toLowerCase();
            return ['.png', '.jpg', '.jpeg'].includes(`.${fileType}`);
        });

        if (validFiles.length === 0) {
            this.$message.error(this.$t('modelFinetune.uploadImgTips3'));
            return;
        }
        this.uploadImage(fileList,false)
    },
    handleZipChange(event) {
        const fileList = event.target.files;
        this.uploadImage(fileList,true)
    },
    async uploadImage(fileList, zipflag) {
        if (fileList.length && this.uploadUrl) {
            const formData = new FormData();
            if (zipflag) {
                formData.append('file', fileList[0]);
            } else {
                fileList.forEach((file, index) => {
                    formData.append('files', file);
                });
            }
            try {
                this.$emit('start-loading');
                const response = zipflag ? await sdLorauploadZip(this.uploadUrl,formData) : await sdLorauploadImages(this.uploadUrl,formData)
                const res = response.data
                if (res.code == 200) {
                    this.processItems(res,this.uploadUrl)
                } else {
                    this.$message.error(res.msg)
                    this.$emit('stop-loading');
                }
                
            } catch (error) {
                this.$message.error(error)
                this.$emit('stop-loading');
            }
        }
    },
    async processItems(res,url) {
        try {
            for (let item of res.data) {
                let imgUrl = `${url}/get_image/${item.data.uuid}`
                this.uuidList.push({image:imgUrl,...item.data});
            }
            // 所有异步操作完成后触发事件
            this.$emit('uploadSuccess', this.uuidList);
            this.$emit('stop-loading');
            this.uuidList = []
        } catch (error) {
            this.$emit('stop-loading');
            this.$message.error(error)
            // 根据需要处理错误，例如触发一个失败事件
            // this.$emit('uploadFailed', error);
        }
    },
    handleDrop(e) {
        e.stopPropagation();
        e.preventDefault();
        const files = e.dataTransfer.files;
        if (!files || files.length === 0) {
            this.$message.error(this.$t('modelFinetune.uploadImgTips1'));
            return;
        }

        // 过滤不符合格式的文件
        const validFiles = Array.from(files).filter(file => {
            const fileType = file.name.split('.').pop().toLowerCase();
            return ['.png', '.jpg', '.jpeg'].includes(`.${fileType}`);
        });

        if (validFiles.length === 0) {
            this.$message.error(this.$t('modelFinetune.uploadImgTips3'));
            return;
        }
        

        // 处理符合格式的文件
        this.handleFileList(validFiles)
    },
    handleDragLeave(e) {
        e.stopPropagation();
        e.preventDefault();
        console.log("2")
    },
    handleDragOver(e) {
        e.stopPropagation();
        e.preventDefault();
        console.log("3")
        
       },
    bindEvents() {
        const dropbox = this.$refs.pickerArea;
        console.log(dropbox)
        // 防止重复绑定事件，需要在 data 中初始化 bindDrop 为 false
        if (!dropbox || this.bindDrop) { return; }
        // 绑定拖拽事件，在组件销毁时解绑
        dropbox.addEventListener('drop', this.handleDrop, false);
        dropbox.addEventListener('dragleave', this.handleDragLeave);
        dropbox.addEventListener('dragover', this.handleDragOver);
        this.bindDrop = true;
    },
    uploadClick() {
        this.$refs.uploadBtn.handleUpClick()
    },
    uploadZipClick() {
        this.$refs.fileInput.click();
    },
   },
   beforeDestroy() {
        // 组件销毁前解绑拖拽事件
        try {
            const dropbox = this.$refs.pickerArea;
            dropbox.removeEventListener('drop', this.handleDrop);
            dropbox.removeEventListener('dragleave', this.handleDragLeave);
            dropbox.removeEventListener('dragover', this.handleDragOver);
            this.bindDrop = false;
        } catch (e) {}
   },

};
</script>
<style lang='less' scoped>
.pic-upload{
    margin: 0 24px;
    text-align: center;
    height: 60%;
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    padding-bottom: 2rem;
    .pic-upload-drag{
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100%;
        background: #fff;
        border-radius: 12px;
        margin-bottom: 16px;
        .upload-icon{
            width: 50px;
            height: 50px;
            margin-top: 30px;
            margin-bottom: 10px;
            background: url('/img/model/icon_upload.png') no-repeat 50%;
            background-size: contain;
        }
        .upload-text-wrap{
            .upload-tips{
                line-height: 24px;
                font-weight: 500;
                color: rgba(0,0,0,.85);
                font-size: 16px;
                margin: 0;
            }
            .accpet-tips{
                line-height: 16px;
                font-size: 12px;
                color: rgba(0,0,0,.45);
                margin: 0;
                margin-top: 4px;
            }
        }
    }
}
.segmente {
    color: #999;
    font-family: PingFang SC;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
    position: relative;
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: center;
    .line{
        width: 148px;
        height: 1px;
        background: rgba(157,197,226,0.4);
    }
}
.zip-upload{
    display: flex;
    /* justify-content: center; */
    flex-direction: column;
    align-items: center;
    margin-top: 20px;
    .upload-btn{
        color:#101010;
        padding-bottom: 1px;
        border-bottom: 1px solid #101010;
        cursor: pointer;
    }
    .zip-tips{
        font-size: 12px;
        color: rgba(0,0,0,.45);
        line-height: 32px;
    }
}  
</style>