<template>
  <div class="ui container" style="margin-top: 2rem;margin-bottom:-40px" id="dialog">
    <headerModel :modelName="modelName" :minutes="minutes" :seconds="seconds"/>
    <div class="model-dialog-wrapper">
      <dialogLeft :pattern="pattern" :kbName="kbName" :modelName="modelName" :maxlength="maxlength" :counts="counts" :maxTrie="maxTrie" :expireMinutes="expireMinutes"></dialogLeft>
      <dialogRight @radioChange="radioChange" @changeKbName="changeKbName" :commonKB="commonKB" :modelName="modelName"></dialogRight>
    </div>
    <div class="model-dialog-footer" :style="{marginTop:`${height}rem`}">
      <span class="text" v-html="$t('modelSquare.modelProvide')">
        
      </span>
    </div>
    <el-dialog
      style="border-radius:2rem;margin-top: 20vh;"
      :visible.sync="dialogVisible"
      :title="$t('modelSquare.useNotice')"
      width="30%"
      center
      :show-close="false"
      :before-close="handleClose">
      <span v-html="$t('modelSquare.agreeNotice')"></span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel" size="mini">{{$t('modelSquare.cancel')}}</el-button>
        <el-button type="primary" @click="confirm" size="mini">{{$t('modelSquare.ok')}}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { llmAgree} from '~/apis/modules/llmchat';
import headerModel from './componenes/headerModel.vue'
import dialogLeft from './componenes/dialogLeft.vue'
import dialogRight from './componenes/dialogRight.vue'
export default {
  name: "App",
  components: { headerModel,dialogLeft,dialogRight },
  data(){
    return{
      pattern:'1',
      kbName:'',
      modelName:'',
      commonKB:'',
      dialogVisible:false,
      maxlength:0,
      height:5,
      minutes:"NA",
      seconds:"NA",
      counts:1000,
      maxTrie:0,
      expireMinutes:30
    }
  },
  methods: {
    handleClose(){

    },
    cancel(){
      this.dialogVisible = false
      history.back()
    },
    confirm(){
      llmAgree({model_name:this.modelName}).then((res)=>{
        this.dialogVisible = false
      }).catch((err)=>{
        this.$message.error(err.msg)
      })
    },
    radioChange(val){
      if(val==='1'){
        this.height = 5
      }else{
        this.height = 3
      }
      this.pattern = val
    },
    changeKbName(val){
      this.kbName = val
    },
    durationFormatter(gap) {
      const second = 1000
      const minute = second * 60
      const hour = minute * 60
      let m = Math.floor((gap % hour) / minute)
      let s = Math.floor((gap % minute) / second)
      this.minutes = this.addZero(m)
      this.seconds = this.addZero(s)

    },
    addZero(num) {
      return num < 10 ? '0' + num : '' + num
    },
    countDown(duration){
      const totalDuration = duration;
      let requestRef = null;
      let startTime;
      let prevEndTime;
      let prevTime;
      let currentCount = totalDuration;
      let endTime;
      let timeDifferance = 0; // 每1s倒计时偏差值，单位ms
      let interval = 1000;
      let nextTime = interval;

      const animate = (timestamp) => {
        if (prevTime !== undefined) {
          const deltaTime = timestamp - prevTime;
          if (deltaTime >= nextTime) {
            prevTime = timestamp;
            prevEndTime = endTime;
            endTime = new Date().getTime();
            currentCount = currentCount - 1000;
            this.durationFormatter(currentCount)
            timeDifferance = endTime - startTime - (totalDuration - currentCount);
            nextTime = interval - timeDifferance;
            // 慢太多了，就立刻执行下一个循环
            if (nextTime < 0) {
              nextTime = 0;
            }
            if (currentCount <= 0) {
              currentCount = 0;
              cancelAnimationFrame(requestRef);
              return;
            }
          }
        } else {
          startTime = new Date().getTime();
          prevTime = timestamp;
          endTime = new Date().getTime();
        }
        requestRef = requestAnimationFrame(animate);
      };

      requestRef = requestAnimationFrame(animate);
    },
    addWaterMarker(name,str){
      let that = this
      let can  = document.createElement('canvas')
      let container = document.querySelector('#dialog')
      container.appendChild(can)
      can.width = 180
      can.height = 100
      can.style.display = 'none'
      let cans = can.getContext('2d')
      cans.rotate(-20 * Math.PI / 180)
      cans.font = 'normal 12px Microsoft Jhenghei'
      cans.fillStyle = 'rgba(223,223,223,1)'
      cans.textAlign = 'center'
      cans.textBaseline = 'Middle'
      cans.fillText(name, can.width / 3 , can.height / 2)
      cans.fillText(str, can.width / 2.7, can.height / 1.6)
      const base64Url = can.toDataURL();
      const watermarkNode = document.querySelector(".watermarkNode");
      const watermarkDiv = watermarkNode || document.createElement("div");
      const styleStr = `position:absolute;
        opacity:0.5;
        top:0;
        left:0;
        width:100%;
        height:100%;
        z-index:1000;
        pointer-events:none;
        background-repeat:repeat;
        background-image:url('${base64Url}')`
      watermarkDiv.setAttribute("style", styleStr);
      watermarkDiv.classList.add("watermarkNode");
      if(!watermarkNode){
        container.style.position = 'relative';
      
        container.insertBefore(watermarkDiv, container.firstChild);
      }
      if (MutationObserver) {
        let MOInstance = new MutationObserver(function () {
          const watermarkNode = document.querySelector(".watermarkNode");
          // 只在watermarkNode元素变动才重新调用 createWatermark
          if (
            !watermarkNode ||
            (watermarkNode && watermarkNode.getAttribute("style") !== styleStr)
          ) {
            // 避免一直触发
            MOInstance.disconnect();
            // 重新创建水印
            that.addWaterMarker(document.querySelector('meta[name="_uid"]').getAttribute('content-ext'),'AI生成内容仅供参考');
          }
        });
    
        MOInstance.observe(container, {
          attributes: true,
          subtree: true,
          childList: true,
        });
      }
    
    },
  },
  async mounted() {
    const urlParams = new URLSearchParams(location.search)
    if(urlParams.has('model_name')){
      this.modelName = urlParams.get('model_name')
      this.modelName==='chatglm2-6b' ? this.maxlength=2000 :  this.maxlength=1000
    }
    const userName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext')
    if(window.config.csrf !== '' && userName !==''){
      this.addWaterMarker(userName,'AI生成内容仅供参考')
    } 
    this.commonKB = document.getElementById('dialog-setting').getAttribute('data-common-kb')
    const firstVisit = document.getElementById('dialog-setting').getAttribute('data-first-visit')
    const createdUnix = +document.getElementById('dialog-setting').getAttribute('data-create-unix')
    const endUnix = +document.getElementById('dialog-setting').getAttribute('data-end-unix')
    this.counts = +document.getElementById('dialog-setting').getAttribute('data-counts')
    this.maxTrie = +document.getElementById('dialog-setting').getAttribute('data-max-tries')
    this.expireMinutes = +document.getElementById('dialog-setting').getAttribute('data-expire-minutes')
    this.countDown((endUnix - createdUnix)*1000)
    if(firstVisit=='true'){
      this.dialogVisible = true
    }
  }
};
</script>
  
<style lang="less" scoped>
  
  
  .model-dialog-wrapper{
    display: flex;
    justify-content: space-between;
  }
  .model-dialog-footer{
    align-items: center;
    display: flex;
    flex-direction: column;
    padding: 10px 0;
    margin-top: 5rem;
    .text{
      font-family: PingFangSC-Regular;
      font-size: 12px;
      font-weight: 400;
      letter-spacing: 0;
      line-height: 20px;
      text-align: center;
      color: rgba(103,104,144,.6);
      }
  }
  
  
</style>