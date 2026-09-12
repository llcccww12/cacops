<template>
<div style="height: 100%;">    
    <div class="__mobile-tip" >
        <img style="width: 132px;height: 97px;" src="/img/model/pc-view.png" alt="">
        <div style="margin-top: 2rem;">{{$t('useInPcWeb')}}</div>
    </div>
    <div class="content-box __content-box" v-loading="loading">
        <div class="main-header">
            <div class="area-l-title">
                <div class="label link" @click="$router.push({name:'cv'})">{{$t('modelSquare.sdModelFinetuen')}}</div>
                <div class="separator"> / </div>
                <div class="label">{{tabList[activeTab]}}</div>
            </div>
            <div class="tab-c">
                <span  v-for="(item, index) in tabList" class="tab-item" :class="activeTab == index ? 'active' : ''">{{item}}</span>
            </div>
        </div>
        <div class="main-body">
            <LoraImage v-if="activeTab===0 && sections.length" :taskId="taskId" :taskUrl="taskUrl" :sections="sections" :modelName="modelName" @changePage="changePage"></LoraImage>
            <LoraTrain v-if="activeTab===1" :taskUrl="taskUrl" :taskId="taskId" :taskName="taskName" @getModel="getModel"></LoraTrain>
            <LoraOutput v-if="activeTab===2"  :taskUrl="taskUrl" :taskId="taskId" :modelObj="modelObj"></LoraOutput>
            <div v-if="activeTab===3" class="task-stop-wrap">
                <div> 
                    <img src="/img/model/del_failed.png" alt="" style="min-height: 260px;">
                    <div class="tips">{{$t('modelFinetune.cvTaskfaildTips')}}</div>
                </div>
            </div>
        </div>
    </div>
</div>
</template>

<script>
import LoraImage from './LoraImage.vue'
import LoraTrain from './LoraTrain.vue'
import LoraOutput from './LoraOutput.vue'
import { getLoraStage } from '~/apis/modules/llmchat'
import { getPromoteData } from '~/apis/modules/common';
export default {
  data() {
    return {
        activeTab: -1,
        tabList: [this.$t('cloudbrainObj.paramsSetting'),this.$t('modelFinetune.modelTraining'),this.$t('modelFinetune.modelTesting')],
        taskId: '',
        taskUrl: '',
        loading: false,
        sections: [],
        modelObj: {},
        modelName: '',
        taskName: ''
    };
  },
  components: { LoraImage, LoraTrain, LoraOutput },
  computed:{
    
  },
  watch: {
    
  },
  methods: {
    changePage(){
        this.getFinetuneStage()
    },
    getFinetuneStage() {
        this.loading = true
        getLoraStage({task_id:atob(this.taskId)}).then(async (res) => {
            this.taskUrl = res.data.url
            if (!this.taskUrl) {
                this.$message.warning(this.$t('modelSquare.ComfyUiWarn'))
                return
            }
            if (res.status === 200) {
                if (res.data.code === 404) {
                    this.activeTab = 3
                } else {
                    this.activeTab = res.data.stage
                    this.taskName = res.data.task.display_job_name
                    if (res.data.lora_models.length > 0) {
                        this.modelObj.modelList = res.data.lora_models.map((item) => item.split('/').pop().split('.')[0])
                        this.modelObj.model = this.modelObj.modelList[0]
                        this.modelObj.imgList = res.data.lora_samples.map((item) => { return this.taskUrl + '/sample_image/' + item.split('/').pop().split('.')[0] })
                        this.modelObj.index = 0
                        console.log(res.data)
                        console.log(this.modelObj)
                    }
                    
                    if (res.data.stage === 0) {
                        const res = await getPromoteData(`model/${this.modelName}.json`)
                        const data = JSON.parse(res.data); 
                        this.sections = data.params_config
                    }
                }
            }
            this.loading = false
        }).catch((error) => {
            this.loading = false
            this.$message.error(error)
        })
    },
    
    getModel(data) {
        this.modelObj = data
        this.activeTab = 2
        // this.getFinetuneStage()
        
    },
  },
  beforeCreate() {
  },
  mounted() {
  },
  beforeMount() {
    if (Object.keys(this.$route.query).length) {
        this.taskId = this.$route.query.id
        this.modelName = this.$route.query.model
        // this.taskName = this.$route.query.name
    }
    this.getFinetuneStage()  
    // this.getTrainProcess()
  },
  beforeDestroy() {

  },
};
</script>

<style scoped lang="less">
.content-box{
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    .main-header{
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 16px 20px 0 20px;
        .area-l-title{
            display: flex;
            align-items: center;
            flex-wrap: wrap;
            font-size: 28px;
            .separator {
                color: rgba(111, 118, 165, 1);
                font-family: SourceHanSansSC;
                margin-right: 10px;
            }
            .label {
                line-height: 39px;
                margin-right: 10px;
                font-family: SourceHanSansSC;
                font-weight: 500;

                &.link {
                color: rgba(3, 102, 214, 1);
                font-weight: 400;
                cursor: pointer;

                &:hover {
                    color: rgba(3, 102, 214, .8);
                }
                }
            }
        }
        
        .tab-c{
            display: flex;
            height: 32px;
            min-width: 300px;
            .tab-item{
                border: 1px solid rgba(41,45,69,0.33);
                background: #EDEBFE;
                // width: 50%;
                padding: 0px 12px;
                display: flex;
                align-items: center;
                justify-content: center;
                color:#101010;
                &.active{
                    color: #1874FF;
                    background: #fff;
                }
                &:first-child{
                    border-radius: 4px 0 0 4px;
                    border-right: none;
                }
                &:last-child{
                    border-left: none;
                    border-radius: 0 4px 4px 0;
                }
            }
        }
    }
    .main-body{
        flex: 1;
        position: relative;
        height: calc(100% - 55px);
        .task-stop-wrap{
            display: flex;
            height: 80%;
            justify-content: center;
            align-items: center;
            .tips{
                text-align: center;
                color: red;
            }
        }
    }
    
    
}

</style>