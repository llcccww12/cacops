<template>
    <div class="content-box1">
        <div class="main-body1">
            <div class="wrap">
                <div class="left-wrap">
                    <div class="params-box">
                        <div class="label-wrap">
                            <div class="item-label">{{$t('modelFinetune.loraModelName')}}</div>
                            <div>
                                <el-select v-model="model" style="width:100%">
                                    <el-option v-for="(item ,index) in modelOptionList" :key="item" :label="item" :value="item" @click.native ="handleSelectChange(index)"/>
                                </el-select>
                                <div class="img-wrap">
                                    <div class="img-item" v-for="(item ,index) in displayedImages" :key="item" :class="{ 'active': index === selectedIndex }" @click="handleImageClick(index)">
                                        <img :src="item" alt="">
                                    </div>
                                </div>
                                <div v-if="imgList.length > 4" class="toggle-button" >
                                    <span @click="toggleExpand">{{ isExpanded ? $t('org.fold') : $t('org.unfold') }}</span>
                                </div>
                            </div>
                        </div>
                        <div class="label-wrap">
                            <div class="item-label">{{$t('modelFinetune.picWidth')}}</div>
                            <div>
                                <el-slider v-model="params.width" show-input 
                                    :show-input-controls="false" :min="128" 
                                    :max="1024" :step="1"></el-slider>
                            </div>
                        </div>
                        <div class="label-wrap">
                            <div class="item-label">{{$t('modelFinetune.picHeight')}}</div>
                            <div>
                                <el-slider v-model="params.height" show-input 
                                    :show-input-controls="false" :min="128" 
                                    :max="1024" :step="1"></el-slider>
                            </div>
                        </div>
                        <div class="label-wrap">
                            <div class="item-label">{{$t('modelFinetune.loraRandomSeed')}}</div>
                            <div>
                                <el-input placeholder="" v-model="params.seed" @input="handleInput">
                                    <el-button slot="append" icon="el-icon-refresh" @click="setSeed">Random seed</el-button>
                                </el-input>
                                <!-- <el-slider v-model="params.seed" show-input 
                                    :show-input-controls="false" :min="-1" 
                                    :max="10000" :step="1"></el-slider> -->
                            </div>
                        </div>
                        <div class="label-wrap">
                            <div class="item-label">{{$t('modelFinetune.loraSamplingSteps')}}</div>
                            <div>
                                <el-slider v-model="params.steps" show-input 
                                    :show-input-controls="false" :min="1" 
                                    :max="60" :step="1"></el-slider>
                            </div>
                        </div>
                        <div class="label-wrap">
                            <div class="item-label">{{$t('modelFinetune.loraCFGScale')}}</div>
                            <div>
                                <el-slider v-model="params.cfg_scale" show-input 
                                    :show-input-controls="false" :min="1.0" 
                                    :max="30.0" :step="0.5"></el-slider>
                            </div>
                        </div>
                        <div class="label-wrap">
                            <div class="item-label">{{$t('modelFinetune.loraNegativePrompt')}}</div>
                            <div>
                                <el-input
                                    type="textarea"
                                    :placeholder="$t('modelSquare.negativePromptPlaceholder')"
                                    v-model="params.negative_prompt"
                                    maxlength="1024"
                                    show-word-limit
                                    rows="4"
                                    resize="none"
                                    style="margin-left: -2px;"
                                >
                                </el-input>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="right-wrap">
                    <div class="main-code-body">
                        <div class="output-text">
                            <div v-if="params.prompt" class="out-prompt">
                                <span>{{ params.prompt }}</span>
                                <div class="prompt-time">{{ unixTime }}</div>
                            </div>
                        </div>
                        <div class="output-img" v-if="showLoading">
                            <div>
                                <img src="/img/loading.svg" style="width: 100px;height:100px;">
                            </div>
                        </div>
                        <div class="output-img" v-else>
                            <div class="img-box">
                                <img :src="imgUrl" alt="">
                            </div>
                        </div>
                    </div>
                    <div class="main-input-body">
                        <div class="input-box">
                            <div class="ask-input">
                                <textarea id="code-input" @keydown.enter="carriageReturn($event)" :placeholder="$t('modelSquare.sdPlaceholder')" style="height: auto;" v-model="showPrompt"></textarea>
                                <button class="opera-submit" @click="submit()"></button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
   
</template>
<script>
import dayjs from 'dayjs';
import { postLoraGenerateImg, getLoraGenerateImgProcess } from '~/apis/modules/llmchat';
import { Thread } from '~/pages/model/wenxin/constant.js'
export default {
   name: '',
   components: {
     
   },
   mixins: [],
   props: {
    taskUrl: { type: String, default: "", },
    taskId: { type: String, default: "", },
    modelObj: { type: Object, default: () => {} },
   },
   data() {
     return {
         params: {
            negative_prompt: '',
            prompt:'',
            width: 1024,
            height: 1024,
            seed: -1,
            steps: 50,
            cfg_scale: 1.0,
            guidance: 3.5,
            lora_weights: ''
         },
         showPrompt: '',
         unixTime: '',
         showLoading: false,
         imgUrl: '',
         modelOptionList: [],
         model: '',
         imgList: [],
         selectedIndex: 0, // 当前选中的索引
         isExpanded: false,
         
     }
   },
   computed: {
    displayedImages() {
      return this.isExpanded ? this.imgList : this.imgList.slice(0, 4);
    },
   },
   watch: {
   },
   mounted() {
       if (JSON.stringify(this.modelObj) !== '{}') {
        console.log(this.modelObj)
        this.modelOptionList = this.modelObj.modelList
        this.model = this.modelObj.model
        this.selectedIndex = this.modelObj.index
        console.log("this.modelObj.imgList",this.modelObj.imgList)
        this.imgList = this.modelObj.imgList//this.modelObj.imgList.map((item)=>{return this.taskUrl + '/sample_image/' + item})
    }
    // this.getImgProcess()
   },
   methods: {
    // 下拉框选项变化时触发
    handleSelectChange(index) {
      this.selectedIndex = index;
    },
    // 图片点击时触发
    handleImageClick(index) {
        this.selectedIndex = index;
        this.model = this.modelOptionList[index]
    },
    // 切换展开/收缩状态
    toggleExpand() {
      this.isExpanded = !this.isExpanded;
    },
    carriageReturn(event){
        event.preventDefault()
        if(event.ctrlKey && event.keyCode ==13){
            this.params.prompt = this.params.prompt + '\n'
        }else{
            this.submit()
        }
    },
    async submit() {
        try {
            if (this.showPrompt) {
                this.showLoading = true
                this.params.prompt = this.showPrompt
                const date = new Date()
                this.unixTime = dayjs(date).format('YYYY-MM-DD HH:mm:ss');
                this.showPrompt = ''
                this.params.lora_weights = this.model
                const response = await postLoraGenerateImg({task_id:atob(this.taskId)},this.params)
                const res = response.data
                if (res.code === 404) {
                    this.$message.error(this.$t('modelFinetune.cvTaskfaildTips'))
                } else {
                    if (res.code === 200) {
                        this.getImgProcess()
                    } else {
                        this.$message.error(res.msg)
                        // this.getImgProcess()
                    }
                }
                
            } else {
                this.$message.error(this.$t('modelFinetune.loraInputPrompt'))
            }
        } catch (error) {
            this.showLoading = false
            this.$message.error(error)
        }
    },
    async getImgProcess() {
        try {
            const vm = this
            const thread = new Thread({
                start: function () {
                    getLoraGenerateImgProcess(vm.taskUrl).then((response => { 
                        const res = response.data
                        if (res.code === 200) {
                            if (res.data.status === -1) {
                                thread.stop()
                                const length = res.data.results.length
                                const imgName = length ? res.data.results[length-1].image : ''
                                if (imgName) {
                                    let img = imgName.split('.')[0]
                                    vm.imgUrl = `${vm.taskUrl}/infer_image/${img}`
                                    vm.params.prompt = res.data.results[length-1].user_prompt.prompt
                                }
                            }
                        } else {
                            vm.$message.error(res.msg)
                            // vm.showLoading = false
                            thread.stop()
                        }
                        
                    })).catch((err) => { 
                        thread.stop()
                        vm.$message.error(err)
                    })
                },
                stop: function () {
                    vm.showLoading = false
                },
                number: 0, //这里是轮询次数配置，不配置默认无线轮询
                time: 4000 //这里是轮询的时间 不配置默认 300ms
            })
            // 开始轮询
            thread.run();
        } catch (error) {
            this.$message.error(error)
            vm.showLoading = false
        }
    },
    setSeed() {
        const firstDigit = Math.floor(Math.random() * 9) + 1;
        const remainingDigits = Array.from({ length: 9 }, () => Math.floor(Math.random() * 10));
        const randomNumber = parseInt([firstDigit, ...remainingDigits].join(''), 10);
        this.params.seed = randomNumber 
    },
    handleInput(value) {
        // 将输入值转换为数字
        const numValue = Number(value);

        // 如果输入值小于 -1，则赋值为 -1
        if (numValue < -1) {
            this.params.seed = -1;
        }
    },
   }
};
</script>
<style lang='less' scoped>
.content-box1{
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    .main-body1{
        flex: 1;
        position: relative;
        height: 100%;
        .wrap{
            width:100%;
            display: grid;
            height: calc(100% - 40px);
            gap: 3rem;
            grid-template-columns:minmax(0, 1fr) 3fr;
            padding: 3rem;
            padding-right: 10rem;
            overflow-y: auto;
            &::-webkit-scrollbar {
                width: 0; /* 尝试隐藏滚动条，但效果可能因浏览器而异 */
                background: transparent; /* 设置滚动条背景为透明 */
            }
            .left-wrap{
                border-radius: 10px;
                background-color: rgba(249,250,251,1);
                border: 1px solid rgba(16,16,16,0.15);
                color: rgba(16,16,16,1);
                padding: 26px 14px 20px 22px;
                height: 100%;
                overflow-y: auto;
                .params-box{
                    display: flex;
                    flex-direction: column;
                    .label-wrap{
                        line-height: 32px;
                        margin: 6px 0;
                        .item-label{
                            display: flex;
                            align-items: center;
                        }
                        .img-wrap{
                            display: flex;
                            gap: 10px;
                            margin-top: 10px;
                            flex-wrap: wrap;
                            
                            .img-item{
                                width: 80px;
                                height: 80px;
                                
                                &.active{
                                    border:2px solid #0066FF;
                                    border-radius: 4px;
                                }
                                img{
                                    width: 100%;
                                    height:100%;
                                    object-fit: cover !important;
                                }
                            }

                        }
                        .toggle-button{
                            display: flex;
                            justify-content: flex-end;
                            font-size: 12px;
                            color: #0066FF;
                            margin-top: 4px;
                            margin-bottom: -24px;
                            span{
                                cursor: pointer;
                            }
                        }
                    }
                }
            }
            .right-wrap{
                display: flex;
                flex-direction: column;
                gap: 14px;
                .main-code-body{
                    flex: 1;
                    border-radius: 10px;
                    border: 1px solid rgba(16,16,16,0.15);
                    color: rgba(16,16,16,1);
                    display: flex;
                    flex-direction: column;
                    .output-text{
                        .out-prompt{
                            border-radius: 10px;
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
                    }
                    .output-img{
                        flex: 1;
                        display: flex;
                        align-items: center;
                        justify-content: center;
                        .img-box{
                            width: 40%;
                            overflow: hidden;
                            img{
                                max-width: 100%;
                            }
                        }
                    }
                }
                .main-input-body{
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
                }
            }
        }
    }
}

</style>