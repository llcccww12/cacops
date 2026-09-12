<template>
    <!--upload.vue 点击上传组件 -->
    <div class="file-selector" @click="handleUpClick">
        <i class="ri-upload-cloud-2-line" style="margin-right: 4px;"></i>
        <span>{{$t('modelFinetune.uploadImg')}}</span>
        <input
            ref="input"
            class="file-selector-input"
            type="file"
            :multiple="multiple"
            :accept="accept"
            @change="handleFiles"
        />
    </div>
</template>
<script>
import {debounce} from 'lodash/function';
export default {
    data() {
        return {
            accept: '.jpg,.jpeg,.png,',
            multiple: true,
            list: [], // 已选择的文件对象
            uploadFinished: true, // 上传状态
            startIndex: 0, // 开始上传的下标，用于追加文件
            maxSize: 10 * 1024 * 1024, //10M(size单位为byte)
            // source: this.$axios.CancelToken.source(), // axios 取消请求
        };
    },
    methods: {
        // 重置
        reset() {
            this.list = [];
            this.source.cancel();
            this.startIndex = 0;
            this.uploadFinished = true;
            this.$refs.input && (this.$refs.input.value = null);
        },
        // 调用上传功能
        handleUpClick: debounce(function () {
            // 可在此维护一个上传状态，上传过程中禁用上传按钮
            // if (!this.uploadFinished) this.$message.info('即将覆盖之前的文件~');
            
            this.$refs.input.dispatchEvent(new MouseEvent('click'))
        }, 300),
        handleFiles(e) {
            const files = Array.from(e?.target?.files)
            console.log(files)
             // 检查文件格式
            const isValid = files.every(file => {
                const fileType = file.name.split('.').pop().toLowerCase();
                return this.accept.split(',').includes(`.${fileType}`);
            });

            if (!isValid) {
                this.$message.error(this.$t('modelFinetune.uploadImgAccept'));
                this.$refs.input.value = ''; // 清空文件输入框
                return;
            }
            this.$emit('fileList', files);
            // this.readFiles(files);
        },
        // 上传之前将文件处理为对象
        readFiles(files) {
            if (!files || files.length <= 0) {
                return;
            }
            for (const file of files) {
                const url = window.URL.createObjectURL(file);
                const obj = {
                    title: file.name.replace(/(.*\/)*([^.]+).*/ig, '$2'), // 去掉文件后缀
                    url,
                    file,
                    fileType: file.type,
                    status: 0, // 状态 -> 0 等待中，1 完成， 2 正在上传，3 上传失败
                    percent: 0, // 上传进度
                };
                // 提前在 data 中定义 list，用来保存需要上传的文件
                this.list.unshift(obj);
            }
            this.$emit('fileList', this.list);
            // 在 data 中定义 startIndex 初始值为 0，上传完成后更新，用于追加上传文件
            // this.startUpload(this.startIndex);
        },

    }
};
</script>
<style lang="less" scoped>
.file-selector {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0 12px;
    height: 40px;
    background-color: rgba(0,199,255,0.8);
    color: rgba(255,255,255,1);
    box-shadow: 0px 2px 6px 0px rgba(0,199,255,0.8);
    font-family: PingFangSC-regular;
    border: 1px solid rgba(0,199,255,0.8);
    border-radius: 12px;
    margin-top: 16px;
    cursor: pointer;
    &:hover {
        background-color: rgba(0,199,255,1);
        transition: background 180ms;
    }

    .file-selector-input {
        display: none;
    }
}
</style>

