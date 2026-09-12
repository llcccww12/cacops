<template>
    <el-dialog
        :visible.sync="dialogShow"
        custom-class="container"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        @closed="closed"
        >
        <div class="img-wrap">
            <div class="left-box">
                <div class="large-swiper">
                    <svg @click="prevSlide" width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg" class="arrow"><g  filter="url(#filter0_b_649_7589)"><circle  cx="24" cy="24" r="24" transform="matrix(-1 0 0 1 48 0)" fill="#F1F2F6" fill-opacity="0.7"></circle></g><path  opacity="0.4" d="M27 32L19 24L27 16" stroke="#191919" stroke-width="2" stroke-linejoin="round"></path><defs ><filter  id="filter0_b_649_7589" x="-10" y="-10" width="68" height="68" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood  flood-opacity="0" result="BackgroundImageFix"></feFlood><feGaussianBlur  in="BackgroundImageFix" stdDeviation="5"></feGaussianBlur><feComposite  in2="SourceAlpha" operator="in" result="effect1_backgroundBlur_649_7589"></feComposite><feBlend  mode="normal" in="SourceGraphic" in2="effect1_backgroundBlur_649_7589" result="shape"></feBlend></filter></defs></svg>
                    <img v-if="fileLength > 0 && fileList[imgIndex]" :src="fileList[imgIndex].image" alt="">
                    <svg @click="nextSlide" width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg" class="arrow"><g  filter="url(#filter0_b_649_7593)"><circle  cx="24" cy="24" r="24" fill="#F1F2F6" fill-opacity="0.9"></circle></g><path opacity="0.4" d="M21 32L29 24L21 16" stroke="#191919" stroke-width="2" stroke-linejoin="round"></path><defs ><filter  id="filter0_b_649_7593" x="-10" y="-10" width="68" height="68" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood  flood-opacity="0" result="BackgroundImageFix"></feFlood><feGaussianBlur  in="BackgroundImageFix" stdDeviation="5"></feGaussianBlur><feComposite  in2="SourceAlpha" operator="in" result="effect1_backgroundBlur_649_7593"></feComposite><feBlend  mode="normal" in="SourceGraphic" in2="effect1_backgroundBlur_649_7593" result="shape"></feBlend></filter></defs></svg>
                </div>
                <div class="thum-swiper">
                    <div class="swiper-box">
                        <div class="wrap-imgs" ref="wrapImgs">
                            <ul class="swiper-imgs" :style="trackStyle">
                                <li @click="selectImg(index)" v-for="(item,index) in fileList" :key="index" class="img-item" :class="{'active':imgIndex==index}">
                                    <img :src="item.image" alt="">
                                </li>
                            </ul>
                        </div>
                    </div>
                    <i class="ri-arrow-left-s-line img-icon" style="left: -12px;" @click="prevSlide"></i>
                    <i class="ri-arrow-right-s-line img-icon" style="right: -12px;" @click="nextSlide"></i>
                </div>
            </div>
            <div class="right-box">
                <div class="label-container">
                    <div class="tags-wrap">
                        <el-tag :key="tag + '' + index" v-for="(tag,index) in captionList[imgIndex]" closable :disable-transitions="false" @close="handleClose(tag)">
                            {{ tag }}
                        </el-tag>
                        <el-input class="input-new-tag" v-if="inputVisible" v-model="inputValue" ref="saveTagInput" size="small" @keyup.enter.native="handleInputConfirm">
                        </el-input>
                        <el-button v-else class="button-new-tag" size="small" @click="showInput">+ New Tag</el-button>
                    </div>
                </div>
                <div class="edit-label-wrap">
                    <div class="edit-label-title">
                        {{$t('modelFinetune.capationAddPic')}}
                    </div>
                    <el-input v-model="addTagValue" :placeholder="$t('modelFinetune.loraInputPrompt')"></el-input>
                    <div class="edit-btn-box">
                        <el-button size="medium" type="primary" class="caption-btn" @click="addAllTagClick(false)">{{$t('modelFinetune.capationAddBegin')}}</el-button>
                        <el-button size="medium" type="primary" class="caption-btn" @click="addAllTagClick(true)">{{$t('modelFinetune.capationAddEnd')}}</el-button>
                    </div>
                </div>
            </div>
        </div>
    </el-dialog>
</template>
<script>
import { postAllAcaption, postSingleCaption } from '~/apis/modules/llmchat';
export default {
   name: 'CaptionDialog',
   components: {
     
   },
   mixins: [],
   props: {
    fileList: { type: Array, default: () => [] },
    visible: { type: Boolean, default: false },
    currentIndex: { type: Number, default: 0 },
    taskUrl: { type: String, default: "", },
   },
   data() {
     return {
        itemWidth: 75,
        imgIndex: 0,
        dialogShow: false,
        inputVisible: false,
        inputValue: '',
        addTagValue: '',
        fileCopyList: this.fileList
     }
   },
   computed: {
    trackStyle() {
        const translateX = -this.imgIndex * this.itemWidth;
        return {
            transform: `translateX(${translateX}px)`,
            transition: 'transform 0.3s ease', // 添加平滑过渡效果
        };
    },
    fileLength() {
        return this.fileCopyList.length;
    },
    captionList() {
        let list = []
        this.fileCopyList.forEach(item => {
            if (item.caption) {
                list.push(item.caption.split(","))
            } else {
                list.push([])
            }
        });
        return list
    }
   },
   watch: {
    currentIndex(value) {
        console.log(value)
        this.imgIndex = value
    },
    visible: function (val) {
      this.dialogShow = val;
    },
   },
   mounted() {
     
   },
   methods: {
    closed() {
        this.$emit("update:visible", false);
    },
    async handleClose(tag) { 
        const uuid = this.fileCopyList[this.imgIndex].uuid
        let parts = this.fileCopyList[this.imgIndex].caption.split(',').map(part => part.trim());
        let newParts = parts.filter(part => part !== tag);
        let newString = newParts.join(',');
        newString = newString.trim();
        let text = newString
        await this.eidtSingeCaption(uuid,text)
    },
    showInput() {
        this.inputVisible = true;
        this.$nextTick(_ => {
            this.$refs.saveTagInput.$refs.input.focus();
        });
    },
    async eidtSingeCaption(uuid,text) {
        const response = await postSingleCaption(this.taskUrl,uuid,{full_text:text});
        const res = response.data
        console.log(res)
        if (res.code === 200) {
            this.inputValue = ''
            const mergeFileList = this.fileCopyList.map((item) => {
                console.log(item)
                if (item.uuid === res.data.uuid) {
                    item.caption=res.data.text
                }
                return item
            })
            this.fileCopyList = mergeFileList
            console.log(this.fileCopyList)
        } else {
            this.$message.error(res.data.msg || '编辑标签失败！')
        }
    },
    async handleInputConfirm() {
        try {
            if (this.inputValue) {
                const uuid = this.fileCopyList[this.imgIndex].uuid
                console.log(this.fileCopyList[this.imgIndex])
                const text = this.fileCopyList[this.imgIndex].caption + ',' + this.inputValue
                await this.eidtSingeCaption(uuid,text)
                
            }
            this.inputVisible = false;
        } catch (error) {
            this.inputVisible = false;
            this.$message.error(error)
        }
    },
    
    async addAllTagClick(append) {
        try {
            if (this.addTagValue) {
                const response = await postAllAcaption(this.taskUrl,{text:this.addTagValue,append:append});
                const res = response.data
                console.log(res)
                if (res.code === 200 && res.data.length > 0) {
                    this.addTagValue = ''
                    const mergeFileList =  this.fileCopyList.map((item1) => {
                        const item2 = res.data.find(item => item.uuid === item1.uuid)
                        if (item2) {
                            item1.caption = item2.text
                        }
                        return item1
                    })
                    this.fileCopyList = mergeFileList
                    this.$emit("captionChange",this.fileCopyList)
                } else {
                    this.$message.error(res.data.msg || '添加标签失败！')
                }
            } else {
                this.$message.error(this.$t('modelFinetune.loraInputPrompt'))
            }
            
        } catch (error) {
            this.$message.error(error)
        }
    },
    selectImg(index) {
        this.imgIndex = index 
    },
    prevSlide() {
        if (this.imgIndex > 0) {
            this.imgIndex -= 1;
        } else {
            this.imgIndex = this.fileLength - 1;
        }
    },
    nextSlide() {
        if (this.imgIndex < this.fileLength - 1) {
            this.imgIndex += 1;
        } else {
            this.imgIndex = 0;
        }
    },
   }
};
</script>
<style lang='less' scoped>
/deep/ .container{
    width: 75vw;
    min-width: 800px;
    min-height: 500px;
    background: #fff;
    border-radius: 16px;
    margin-top: 5vh;
    .el-dialog__body{
        padding: 0 28px 28px 28px;
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
    .img-wrap{
        display: flex;
        flex-direction: row;
        .left-box{
            width: 50%;
            position: relative;
            display: flex;
            flex-direction: column;
            align-items: center;
            .large-swiper{
                display: flex;
                align-items: center;
                justify-content: space-between;
                width: 100%;
                height: 50vh !important;
                -webkit-user-select: none;
                -moz-user-select: none;
                user-select: none;
                img{
                    border-radius: 12px;
                    max-width: calc(100% - 110px);
                    text-align: center;
                    max-height: 100%;
                }
                .arrow {
                    cursor: pointer;
                    &:hover path{opacity: 1;}
                }
                .wrap{
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                    width: 100%;
                    overflow: scroll;
                    
                }
            }
            ul, li {
                margin: 0;
                padding: 0;
                list-style: none; /* 移除项目符号 */
            }
            .thum-swiper{
                height: 75px;
                padding-left: 20px;
                padding-right: 20px;
                width: 100%;
                margin-top: 2rem;
                position: relative;
                .swiper-box{
                    height: 100%;
                    position: relative;
                    .wrap-imgs{
                        overflow: hidden;
                        height: 100%;
                        position: relative;
                        .swiper-imgs{
                            display: flex;
                            height: 100%;
                            transition-property: color, background-color, border-color, outline-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter, backdrop-filter, -webkit-backdrop-filter;
                            transition-timing-function: cubic-bezier(.4,0,.2,1);
                            transition-duration: .15s;
                            gap: 10px;
                            .img-item{
                                height: auto;
                                cursor: pointer;
                                text-align: center;
                                border: 0;
                                box-sizing: border-box;
                                &.active{
                                    img{
                                        padding: 2px;
                                        background: #fff;
                                        border: 2px solid #3162ff;
                                        border-radius: 8px;
                                        transition: border .1s, padding .2s;
                                    }
                                }
                                img{
                                    border-radius: 8px;
                                    object-fit: cover;
                                    height: 100%;
                                    text-align: center;
                                    width: 72px !important;
                                }
                            }
                        }
                    }
                }
                .img-icon{
                    position: absolute;
                    top: 50%;
                    cursor: pointer;
                    font-size: 32px;
                    transform: translateY(-50%);
                }
            }
        }
        .right-box{
            width: 50%;
            &::before{
                position: absolute;
                top: 0;
                height: 100%;
                margin-left: 25px;
                border-left: 1px solid #e8eaef;
                content: "";
            }
            .label-container{
                margin-top: 15px;
                margin-bottom: 32px;
                margin-left: 57px;
                .tags-wrap{
                    height: 34vh;
                    padding: 16px;
                    overflow-y: auto;
                    background: #f1f2f6;
                    border-radius: 8px;
                    width: 100%;
                    border: 1px solid #dcdfe6;
                    box-sizing: border-box;
                    color: #606266;
                    display: inline-block;
                    outline: none;
                    position: relative;
                    font-size: 14px;
                    .button-new-tag {
                        height: 32px;
                        line-height: 30px;
                        padding-top: 0;
                        padding-bottom: 0;
                    }
                    .input-new-tag {
                        width: 90px;
                        // vertical-align: bottom;
                    }
                }
            }
            .edit-label-wrap{
                margin-left: 57px;
                .edit-label-title{
                    margin-bottom: 16px;
                    margin-left: 8px;
                    color: #191919;
                    font-weight: 500;
                    font-size: 18px;
                    font-style: normal;
                    line-height: 24px;
                }
                .edit-btn-box{
                    display: flex;
                    justify-content: flex-end;
                    margin-top: 24px;
                    .caption-btn{
                        background-color: #0066ff;
                        color: #fff;
                        border-radius: 8px;
                    }
                }
            }
            
        }
    }
}
/deep/ .el-tag{
    background: #fff;
    color: #101010;
    border: none;
    line-height: 32px;
    height: 32px;
    font-size: 14px;
    margin-right: 8px;
    margin-bottom: 8px;
}
</style>