<template lang="">
    <div class="model-dialog-left">
        <div class="model-dialog-text" @scroll="onScroll" ref="chatContainer" :class="dialogHeight == true ? 'dialog-height-llm':'dialog-height-kb'">
          <div class="model-dialog-title" style="margin-bottom: 1.5rem;">
            <div class="text">
              <p>{{$t('modelSquare.dialogtips1')}}</p>
              <p>{{$t('modelSquare.dialogtips21')}} {{modelName}} {{$t('modelSquare.dialogtips22')}}</p>
              <p>{{$t('modelSquare.dialogtips3')}}</p>
              <p>{{$t('modelSquare.dialogtips4')}}</p>
              <p>{{$t('modelSquare.dialogtips5')}}</p>
            </div>
          </div>
          <div class="model-dialog-title">
            <div class="text">
              <p>{{$t('modelSquare.dialogtips6')}}</p>
            </div>
          </div>
          <div style="width:100%; position: relative;" >
            <SessionWindow
              ref="sessionWindow"
              :session-data="sessionData"
            ></SessionWindow>
          </div>
          
        </div>
        <div class="model-dialog-input">
          <el-input 
          ref="inputRef"
          v-model:value="prompt"
          type="textarea"
          :placeholder="placeholder"
          :maxlength="maxlength"
          :autosize="{ minRows: 1, maxRows: 4}"
          :show-word-limit="showLimit"
          @input="inputPrompt"
          @keydown.enter.native="carriageReturn($event)"/>
          <div class="chat-count">
            <span>{{countsRatio}}</span>
            <span  style="color: #ff5e00;">{{$t('modelSquare.maxTries',{maxTries:maxTrie})}}</span>
          </div>
        </div>
      </div>
</template>
<script>
import { llmChat,llmKbChat,llmCount} from '~/apis/modules/llmchat';
import SessionWindow from './SessionWindow.vue'
export default {
    name: "dialogLeft",
    props: {
      pattern:{type:String,required:true,default:'1'},
      kbName:{type:String,default:''},
      modelName:{type:String,default:''},
      maxlength:{type:Number,default:0},
      counts:{type:Number,default:0},
      maxTrie:{type:Number,default:0},
      expireMinutes:{type:Number,default:30}
    },
    components: { SessionWindow,},
    data() {
        return {
            prompt:'',
            placeholder:this.$t('modelSquare.promptPlaceholder'),
            sessionData:[],
            history:[],
            sessionRecordData:[],
            isAutoScroll:true,
            countsRatio:0,
            elementHeight:0,
            chatFlag:false,
            kbChatFlag:false,
            showLimit:false,
            dialogHeight:true,
            maxTries:0,
            count:0
        };
    },
    watch:{
      sessionRecordData(val){
        if (val != null){
          this.$refs.sessionWindow.setSessionRecord(val)
          this.$nextTick(() => {
            this.scrollBottom();
          })
        }
      },
      counts(){
        this.countsRatio = `${this.counts}/${this.maxTrie}`
      },
      pattern(val){
        if(val==='2'){
          this.dialogHeight = false
        }else{
          this.dialogHeight = true
        }
      }
    },
    methods:{
        inputPrompt(value){
          if(value.length>=this.maxlength){
            this.showLimit = true
          }else{
            this.showLimit = false
          }
        },
        onScroll() {
            const scrollDom = this.$refs.chatContainer;
            const scrollTop = scrollDom.scrollTop;
            const offsetHeight = scrollDom.offsetHeight;
            const scrollHeight = scrollDom.scrollHeight;
            // 当滚动到底部，设置 isAutoScroll 为 true
            if (scrollTop + offsetHeight >= scrollHeight) {
              this.isAutoScroll = true;
            } else {
              // 否则，用户正在手动滑动，设置为 false，停止自动滚动
              this.isAutoScroll = false;
            }
        },
         /**
        * 获取窗口高度并滚动至最底层
        */
        scrollBottom() {
          this.$nextTick(() => {
          if (!this.isAutoScroll) return;
              const scrollDom = this.$refs.chatContainer;
              scrollDom.scrollTop =  scrollDom.scrollHeight;
            // animation(scrollDom, scrollDom.scrollHeight);
          })
        },
        carriageReturn(event){
            event.preventDefault()
            if(event.ctrlKey && event.keyCode ==13){
                this.prompt = this.prompt + '\n'
            }else{
                this.sendInputMessage()
            }
        },
        sendInputMessage(){
            const re = new RegExp("^[ ]+$")
            if(!this.prompt || re.test(this.prompt)){
              this.$message.error(this.$t('modelSquare.inputNotEmpty'))
              return
            }
            if(this.chatFlag || this.kbChatFlag){
              this.$message.error(this.$t('modelSquare.sessionChating'))
              return
            }
            
            let query = this.prompt.trim()
            this.prompt = ''
            this.sessionRecordData.push({"role":"user","content":query})
            if(this.pattern === '1'){
              this.sessionRecordData.push({"role": "assistant","content": ''})
              let data = {
                  "query":query,
                  "stream": true,
                  "model_name": this.modelName,
                  history:this.history
              }
              llmCount({'model_name':this.modelName}).then((res)=>{
                this.countsRatio =  `${res.data.counts}/${res.data.max_tries}`
                this.maxTries = res.data.max_tries
                this.count = res.data.counts
                if(res.status===200 && res.data.can_chat===true){
                  this.chat(data)
                }else{
                  this.sessionRecordData.pop()
                  this.sessionRecordData.push({"role": "assistant","content": this.$t('modelSquare.chatExceedCount')})
                }
                
              })
              
            }else{
              if(!this.kbName){
                this.$message.error(this.$t('modelSquare.chatExpired'))
                return
              }
              this.sessionRecordData.push({"role": "assistant","content": '',"docs":[]})
              let data = {
                "query":query,
                "knowledge_base_name": this.kbName,
                "top_k": 5,
                "score_threshold": 1,
                "stream": true,
                "model_name": this.modelName,
                history:this.history,
                "local_doc_url": false
              }
              llmCount({'model_name':this.modelName}).then((res)=>{
                this.countsRatio =  `${res.data.counts}/${res.data.max_tries}`
                this.maxTries = res.data.max_tries
                this.count = res.data.counts
                if(res.status===200 && res.data.can_chat===true){
                  this.kbChat(data)
                }else{
                  this.sessionRecordData.pop()
                  this.sessionRecordData.push({"role": "assistant","content": this.$t('modelSquare.chatExceedCount'),"docs":[]})
                }
              })
            }
            this.isAutoScroll = true
            this.$nextTick(() => {
              this.scrollBottom();
            })
        },
        chat(data){
          let that = this
          this.chatFlag = true
          this.countsRatio = `${this.count + 1}/${this.maxTries}`
          llmChat(data,this.modelName).then((response)=>{
            const reader = response.body.getReader();
            const processBinaryData = async () => {
              while (true) {
                const { done, value } = await reader.read();
                if (done) {
                  that.chatFlag = false
                  if(new TextDecoder().decode(value)==='<illegal>'){
                      popData.content = this.$t('modelSquare.chatIllegal')
                  }
                  // The entire response has been processed
                  //this.history.push({"role":"user","content":data.query},{"role": "assistant","content": res.data})
                  break;
                }
                // Handle the binary data in the 'value' variable
                
                let chars = new TextDecoder().decode(value)
                let popData = that.sessionRecordData[that.sessionRecordData.length - 1];
                if(chars==='<expired>'){
                  popData.content = this.$t('modelSquare.chatExpireMins',{expireMinutes:this.expireMinutes,locaRefresh:location.href})
                  this.countsRatio = `${this.count}/${this.maxTries}`
                  that.chatFlag = false
                  return
                }
                if(chars==='<illegal>'){
                  popData.content = this.$t('modelSquare.chatIllegal')
                  that.chatFlag = false
                  return
                }
                if(chars==='<banned>'){
                  popData.content = this.$t('modelSquare.chatBanned')
                  that.chatFlag = false
                  setTimeout(()=>{
                    location.reload()
                  },1000)
                  return
                }
                popData.content += chars;
                this.scrollBottom();
                // You can process the binary data here and update your UI as needed
              }
            };
        
            // Start processing the binary data
            processBinaryData();
          }).catch((err)=>{
            this.$message({
              type: 'error',
              message: err.message,
            });
          })
          
        },
        kbChat(data){
          let that = this
          let docs = []
          this.kbChatFlag = true 
          this.countsRatio = `${this.count + 1}/${this.maxTries}`
          llmKbChat(data,this.modelName).then((response)=>{
            const reader = response.body.getReader();
            const processBinaryData = async () => {
              while (true) {
                const { done, value } = await reader.read();
                if (done) {
                  that.kbChatFlag = false
                  //if(popData.docs.length===0){
                    //popData.docs.push(JSON.parse(chars.split('<end>')[1]).docs)
                  //}
                  //const docsArray = JSON.parse(chars.split('<end>')[1])
                  // The entire response has been processed
                  break;
                }
                
                // Handle the binary data in the 'value' variable
                let chars = new TextDecoder().decode(value)
                let popData = that.sessionRecordData[that.sessionRecordData.length - 1];
                if(chars==='<expired>'){
                  popData.content = this.$t('modelSquare.chatExpireMins',{expireMinutes:this.expireMinutes})
                  this.countsRatio = `${this.count}/${this.maxTries}`
                  that.kbChatFlag = false
                  return
                }
                if(chars==='<illegal>'){
                  popData.content = this.$t('modelSquare.chatIllegal')
                  that.kbChatFlag = false
                  return
                }
                if(chars==='<banned>'){
                  popData.content = this.$t('modelSquare.chatBanned')
                  that.kbChatFlag = false
                  return
                  setTimeout(()=>{
                    location.reload()
                  },1000)
                }
                if(chars.indexOf('<docs>')!==0){
                  popData.content += chars;
                }else{
                  if(popData.docs.length===0){
                    that.kbChatFlag = false
                    popData.docs.push(JSON.parse(chars.split('<docs>')[1]).docs)
                  }
                }
                this.scrollBottom();
                // You can process the binary data here and update your UI as needed
              }
            };
        
            // Start processing the binary data
            processBinaryData();
          })
          /* llmKbChat(data).then((res)=>{
            if(res.data){
              this.history.push({"role":"user","content":data.query},{"role": "assistant","content": res.data.answer})
              this.sessionData.push({"role": "assistant","content": res.data.answer,"docs":res.data.docs})
            }
          })*/
        },
    },
    mounted() {
    }
}
</script>
<style lang="less" scoped>
  .model-dialog-left{
    width: 70%;
    margin-right:3rem;
    position: relative;
    .dialog-height-llm{
      min-height: 60vh;
      max-height: 635px;
    }
    .dialog-height-kb{
      max-height: 635px;
      height: 100%;
    }
    .model-dialog-text{
      
      overflow-y: auto;
      .model-dialog-title{
        border-color: rgb(229, 231, 235);
        border-width: 1px;
        border-style: solid;
        border-radius: 10px 10px 10px 0px;
        font-size: 14px;
        padding: 20px;
        text-align: left;
        line-height: 23px;
        font-weight: normal;
        font-style: normal;
        background: rgb(249, 250, 251);
        .text{
          width: 100%;
          word-break: break-word;
          word-wrap: break-word;
          p{
            padding: 0;
            margin-bottom: 5px;
            white-space: pre-wrap;
            &:last-child{
              margin-bottom: 0px;
            }
          }
        }
      }
    }
    .model-dialog-input{
      position: absolute;
      width:100%;
      margin-top: 1rem;
      .chat-count{
        float: right;
        margin-top:0.5rem;
      }
    }
  }
  /deep/ .el-textarea__inner{
    resize: none;
  }
</style>