<template>
   <div class="menu-area">
        <div class="menu-select-container">
            <div>
                <div class="text-select-container">
                    <div class="text-content">
                        <div v-for="key in Object.keys(config)" class="label-wrapper">
                            <template v-if="key==='negative_prompt'">
                                <div class="item-label">
                                    <span>{{$t('modelSquare.'+key)}}</span>
                                    <el-tooltip class="item" effect="dark" :content="config[key].desc" placement="right">
                                        <svg viewBox="64 64 896 896" focusable="false" data-icon="question-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M623.6 316.7C593.6 290.4 554 276 512 276s-81.6 14.5-111.6 40.7C369.2 344 352 380.7 352 420v7.6c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V420c0-44.1 43.1-80 96-80s96 35.9 96 80c0 31.1-22 59.6-56.1 72.7-21.2 8.1-39.2 22.3-52.1 40.9-13.1 19-19.9 41.8-19.9 64.9V620c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-22.7a48.3 48.3 0 0130.9-44.8c59-22.7 97.1-74.7 97.1-132.5.1-39.3-17.1-76-48.3-103.3zM472 732a40 40 0 1080 0 40 40 0 10-80 0z"></path></svg>
                                    </el-tooltip>
                                </div>
                                <div>
                                    <el-input
                                        type="textarea"
                                        :placeholder="$t('modelSquare.negativePromptPlaceholder')"
                                        v-model="params.negative_prompt"
                                        maxlength="1024"
                                        show-word-limit
                                        rows="3"
                                        resize="none"
                                        style="margin-left: -2px;"
                                    >
                                    </el-input>
                                </div>
                            </template>
                            <template v-else-if="key==='scheduler_name'">
                                <div class="item-label">
                                    <span>{{$t('modelSquare.'+key)}}</span>
                                        <el-tooltip class="item" effect="dark" :content="config[key].desc" placement="right">
                                        <svg viewBox="64 64 896 896" focusable="false" data-icon="question-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M623.6 316.7C593.6 290.4 554 276 512 276s-81.6 14.5-111.6 40.7C369.2 344 352 380.7 352 420v7.6c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V420c0-44.1 43.1-80 96-80s96 35.9 96 80c0 31.1-22 59.6-56.1 72.7-21.2 8.1-39.2 22.3-52.1 40.9-13.1 19-19.9 41.8-19.9 64.9V620c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-22.7a48.3 48.3 0 0130.9-44.8c59-22.7 97.1-74.7 97.1-132.5.1-39.3-17.1-76-48.3-103.3zM472 732a40 40 0 1080 0 40 40 0 10-80 0z"></path></svg>
                                    </el-tooltip>
                                </div>
                                <div class="picture-num-wrap">
                                    <div class="picture-num-item" :class="index===schedulerIndex ? 'active' : ''"
                                        v-for="(item, index) in schedulerList" :key="index" @click="selectScheduler(index,item)">
                                        {{ item }}
                                </div>
                            </div>
                            </template>
                            <template v-else-if="key!=='prompt'">
                                <div class="item-label">
                                    <span>{{$t('modelSquare.'+key)}}</span>
                                    <el-tooltip class="item" effect="dark" :content="config[key].desc" placement="right">
                                        <svg viewBox="64 64 896 896" focusable="false" data-icon="question-circle" width="1em" height="1em" fill="currentColor" aria-hidden="true"><path d="M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372s166.6-372 372-372 372 166.6 372 372-166.6 372-372 372z"></path><path d="M623.6 316.7C593.6 290.4 554 276 512 276s-81.6 14.5-111.6 40.7C369.2 344 352 380.7 352 420v7.6c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V420c0-44.1 43.1-80 96-80s96 35.9 96 80c0 31.1-22 59.6-56.1 72.7-21.2 8.1-39.2 22.3-52.1 40.9-13.1 19-19.9 41.8-19.9 64.9V620c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8v-22.7a48.3 48.3 0 0130.9-44.8c59-22.7 97.1-74.7 97.1-132.5.1-39.3-17.1-76-48.3-103.3zM472 732a40 40 0 1080 0 40 40 0 10-80 0z"></path></svg>
                                    </el-tooltip>
                                </div>
                                <div>
                                    <el-slider v-model="params[key]" show-input 
                                    :show-input-controls="false" :min="config[key].min" 
                                    :max="config[key].max" :step="config[key].step"></el-slider>
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
export default {
   name: 'SDMenu',
   props: {
    config:{type:Object,default:()=>({})},
   },
   
   data() {
     return {
        params: {
            width: 768,
            height: 768,
            num_images_per_prompt: 1,
            negative_prompt: '',
            steps: 20,
            seed: 0,
            task_id: 0,
            prompt: '',
            scheduler_name:'',
            guidance_scale: 0,
        },
        schedulerList: [],
        schedulerIndex: 0,
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
            }
        }
       }
   },
   methods: {
    selectScheduler(index,item) {
        this.schedulerIndex = index
        this.params.scheduler_name = item
    },
    sendPamras(){
        this.$emit('changeParams',this.params);
    },
    initParams(){
        let params = this.config
        for(let key in params){
           if(params[key] !== null  ){
                if(key==='scheduler_name'){
                    this.schedulerList = params[key].scheduler_list
                    this.schedulerIndex = this.schedulerList.findIndex((item)=>{
                        return item === params[key].default
                    })
                }
                this.params[key] = params[key].default
            }
        }
    },
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
    .text-model{
        font-size: 18px;
        font-weight: 700;
        margin-bottom: 12px;
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
        .picture-num-wrap{
            display: flex;
            flex-wrap: wrap;
            gap: 10px;
            .picture-num-item{
                border: 1px solid rgba(16,16,16,0.15);
                padding: 0 6px;
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

</style>