<template>
  <div>
    <div style="min-height: 164px;padding-top: 24px;">
      <div style="display: flex;justify-content: space-between;">
        <div></div>
        <div class="countdouwn">
          <div v-if="minutes">
            <span>{{$t('modelSquare.experienceDuration')}}：</span>
            <span style="color: #f2711c;">{{ minutes }} </span>{{$t('timeObj.raw_minutes')}}
          </div>
          <el-button size="small" type="primary" @click="opStop" v-if="canStop"><span style="font-size:14px">{{$t('modelSquare.stopExperience')}}</span></el-button>
        </div>
      </div>

      <div class="model-text-wrapper">
          <p class="model-title">{{modelName}} {{$t('modelSquare.llmHeader')}}</p>
          <p class="model-desc">{{modelDesc}}</p>
      </div>
    </div>
    <LoadingMask :loading="maskLoading" :content="maskLoadingContent"></LoadingMask>
  </div>
</template>
<script>
import { stopAiTask,getAiTask } from '~/apis/modules/cloudbrain';
import LoadingMask from '~/components/cloudbrain/LoadingMask.vue';
export default {
components: { LoadingMask },
props: {
  modelName:{type:String,default:''},
  config:{type:Object,default:()=>{}},
},
data() {
  return {
    modelDesc:'',
    taskparams: {
        id: '',
        repoName: '',
        repoOwnerName:'',
    },
    minutes: 0,
    canStop: false,
    timer: null,
    maskLoading: false,
    maskLoadingContent: this.$t('modelSquare.experienceStopTips'),
  };
},

methods:{
  opStop() {
    clearInterval(this.timer);
    this.maskLoading = true
    stopAiTask(this.taskparams,{model_experience:true}).then((res) => {
      if(res.data.code===0){
        this.$message.success(this.$t('modelSquare.experienceStopSuccess'))
        this.canStop = false
        this.minutes = 0
      }else{
        this.$message.error(res.data.msg)
      }
      this.maskLoading = false
    }).catch((err) => {
      this.maskLoading = false
      this.$message.error(err)
    })
  },
  initParams(){
    this.modelDesc = this.config?.descr
    this.taskparams.repoName = this.config?.configs?.repo_name
    this.taskparams.repoOwnerName = this.config?.configs?.repo_owner_name
    this.getTaskInfo()
  },
   timeToMinutes(timeString) {
  // 使用正则表达式分割字符串
      const [hours, minutes, seconds] = timeString.split(':').map(Number);

      // 将小时转换为分钟并加上分钟和秒转换为分钟的部分
      const totalMinutes = hours * 60 + minutes + (seconds / 60);

      return Math.floor(totalMinutes);
  },
  getTaskInfo() {
    getAiTask(this.taskparams).then((res => {
        res = res.data;
        if (res.code == 0) {
          const data = res.data;

          if(data.task.status!=="RUNNING"){
            return
          }
          this.canStop = data.can_experience
          let totalMinutes = this.timeToMinutes(data.task.formatted_duration)
          this.minutes = totalMinutes
          let vm = this
          this.timer = setInterval(() => {
              vm.minutes = vm.minutes + 1
          }, 60000);
        }
    })).catch((err) => {
        this.$message.error(err)
    })
 },

},
watch:{
  config:function(value){
    console.log("ssssss",value)
      if(JSON.stringify(value) !== '{}'){
        this.initParams()
      }

  }
},
computed: {

},
beforeMount() {
    const urlParams = new URLSearchParams(location.search)
  if (urlParams.has('id')) {
      this.taskparams.id = +atob(urlParams.get('id'))
    }
},
beforeDestroy() {
    // 实例销毁之前对点击事件进行解绑
    clearInterval(this.timer);
},
}
</script>
<style lang="less" scoped>
.model-type{
  display: inline-flex;
  padding: 0.6rem 1.2rem;
  border: 1px solid rgba(229, 231, 235, 1);
  border-radius: 5px;
  color: rgba(16, 16, 16, 1);
  margin-right: 2rem;
  font-size: 16px;
  position: relative;
  p{
    line-height: 100%;
  }
}
.model-activate{
  border: 1px solid rgb(1, 145, 255);
  color: rgb(1, 145, 255);

}
.model-activate::before{
  position: absolute;
  content: "";
  width: 0;
  height: 0;
  border-top:7px solid rgb(1, 145, 255) ;
  border-left: 7px solid transparent;
  border-right: 7px solid transparent;
  top: 100%;
  left: 42%;

}
.model-text-wrapper{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .model-title{
    color: rgba(16, 16, 16, 1);
    font-size: 28px;
    margin-bottom: 0.5rem;
  }
  .model-desc{
    color: rgba(136, 136, 136, 0.87);
    font-size: 14px;
  }
}
.countdouwn {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 30px;
  color: rgba(16,16,16,1);
  font-size: 16px;
}


</style>
