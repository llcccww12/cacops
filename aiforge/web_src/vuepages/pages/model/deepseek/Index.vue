<template>
  <div>
    <div class="menu-mask" v-if="isMobileMenuOpen" @click="toggleMobileMenu"></div>
    <div class="ui container chat-wrap">
      <div class="model-text-wrapper">
        <p class="model-title">{{ modelName }} {{ $t('modelSquare.llmHeader') }}</p>
      </div>
      <div class="main-container">
        <div class="main-left-menu" :class="{ 'mobile-menu-active': isMobileMenuOpen }">
          <div class="menu-wrap">
            <div class="sidbar-header"> 
              <div class="sidbar-title">
                <div> ⭐ {{c2netContent.mainTitle}} </div>
                  </div>
                  <div class="sidbar-sub-title">{{c2netContent.subTitle}}</div>
                </div>
            <div class="sidbar-body">
              <div class="sidbar-main">
                <a v-for="item in c2netList" :key="item.name" class="sidebar-item" :href="item.address" target="_blank">● {{ item.name }}</a>
              </div>
            </div>
            <div class="sidbar-footer">
              <img src="/img/model/c2net.png" alt="">
              <img src="/img/model/openi.png" alt="">
              <div style="text-align: center;font-size: 12px;">
                <span> <a href="/home/model_privacy" target="_blank">{{ $t('modelSquare.modelProvide') }}</a></span>
              </div>
            </div>
          </div>
        </div>
        <div class="main-code-wrap">
          <div style="position: relative;">
            <div class="main-code-body">
              <div class="output-text">
                <div class="output-content" ref="chatContainer" @scroll="onScroll">
                  <div class="out-header-wrap">
                    <div class="c2net-support-list" @click="toggleMobileMenu"> 
                      <span>算力</span>
                      <span>支持</span>
                    </div>
                    <div class="output-welcome-banner">
                      <div v-if="modelName==='DeepSeek'">{{ $t('modelSquare.deepseekHeaderTips') }}</div>
                      <div v-else>{{ $t('modelSquare.chatHeaderTips', { modelName: modelName }) }}</div>
                    </div>
                  </div>
                  <div style="width:100%; position: relative;">
                    <SessionWindow ref="sessionWindow" :img="imgValue" :useLoading="true" :opFlag="modelName === 'DeepSeek'" @stopGeneration="stopGeneration" @refreshAnswer="refreshAnswer"></SessionWindow>
                  </div>
                </div>
              </div>
            </div>
            <div class="main-code-input">
              <el-button type="text" @click="clearChat">{{ $t('modelSquare.newChat') }}</el-button>
              <div class="input-container">
                <div class="input-box">
                  <div class="ask-input">
                    <textarea id="code-input" :placeholder="$t('modelSquare.promptPlaceholder')" style="height: auto;"
                      v-model="prompt" @keydown.enter="carriageReturn($event)"></textarea>
                    <button class="opera-submit" @click="submit"></button>
                  </div>
                </div>
                <div class="tips-privacy">
                  <span> <a href="/home/model_privacy" target="_blank">{{ $t('modelSquare.modelProvide') }}</a></span>
                </div>
              </div>
              
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { llmChat } from '~/apis/modules/deepseek';
import SessionWindow from '~/pages/model/llms/componenes/SessionWindow.vue'
import ChatMenu from './ChatMenu.vue';
import { getPromoteData } from '~/apis/modules/common';
import dayjs from 'dayjs';


export default {
  components: { SessionWindow, ChatMenu },
  data() {
    return {
      prompt: '',
      sessionRecordData: [],
      isAutoScroll: true,
      chatFlag: false,
      modelName: 'DeepSeek',
      config: {},
      loginName: 'xxx',
      params: {
        messages: [],
        model: 4,
        chat_id: '',
        flag: false,
        id: 0,
        service: 'auto'
      },
      abortController: null,
      isMobileMenuOpen: false,
      c2netList:[],
      c2netErrMsg: '',
      c2netContent: {},
      promotePath: `model/modelexperiencedeepseeknpu${document.documentElement.attributes["lang"].nodeValue == "zh-CN" ? '' : '_en'}.json`,
      imgContent: {},
      regexConfig: {}
    }
  },
  watch: {
    sessionRecordData(val) {
      if (val != null) {
        this.$refs.sessionWindow.setSessionRecord(val)
        this.$nextTick(() => {
          this.scrollBottom();
        })
      }
    },
  },
  computed: {
    imgValue(){
      // 方法一：使用 for...of 提前返回
      for (const key of Object.keys(this.imgContent)) {
        if (this.modelName.includes(key)) {
          return this.imgContent[key];
        }
      }
      return null; // 无匹配时返回默认值
    },
    messagesSubVal(){
      return this.params.messages
    }
  },
  methods: {
    generateName(name) {
      let str = name.toLocaleLowerCase();
      const reg1 = /[^a-z0-9_\-]+/g;
      const reg2 = /^[_\-]+/g;
      const reg3 = /[_]+$/g;
      str = str.replace(reg1, '').replace(reg2, '').replace(reg3, '');
      str = str.slice(0, 5);
      const now = Date.now();
      return str + dayjs(now).format('YYYYMMDDHH') + (now / 1000).toFixed(0).slice(-5);
    },
    toggleMobileMenu() {
      
      this.isMobileMenuOpen = !this.isMobileMenuOpen;
    },
    clearChat() {
      this.sessionRecordData = []
      this.params.messages = []
    },
    carriageReturn(event) {
      event.preventDefault()
      if (event.ctrlKey && event.keyCode == 13) {
        this.prompt = this.prompt + '\n'
      } else {
        this.submit()
      }
    },
    async submit() {
      this.isAutoScroll = true
      this.scrollBottom()
      // this.$refs.childMenu.sendPamras()
      const re = new RegExp("^[ ]+$")
      if (!this.prompt || re.test(this.prompt)) {
        this.$message.error(this.$t('modelSquare.inputNotEmpty'))
        return
      }
      if (this.chatFlag) {
        this.$message.error(this.$t('modelSquare.sessionChating'))
        return
      }
      this.params.chat_id = this.generateName(this.loginName)
      const newMessage = { role: "user", content: this.prompt.trim() }
      this.params.messages = []
      this.params.messages.push(newMessage)
      this.prompt = ''
      this.sessionRecordData.push({ "role": "user", "content":  this.params.messages[this.params.messages.length-1].content })
      this.sessionRecordData.push({ "role": "assistant", "content": '', "loading": true, "c2netNameTips": '', "c2netName": '' })
      // this.params.model = this.modelName === 'DeepSeek' ?  'auto' : this.modelName
      this.chat(this.params)
    },
    // 停止生成
    stopGeneration() {
      
      if (this.abortController) {
        this.abortController.abort(); // 中止请求和流
        this.abortController = null;
        this.chatFlag = false;
        const popData = this.sessionRecordData[this.sessionRecordData.length - 1];
        popData.loading = false;
        this.$message({ type: 'success', message: this.$t('modelSquare.stopedChat') });
      }
    },
    refreshAnswer(index){
      console.log(index)
      console.log("xxxxxxxxxxxxxxxx")
      console.log(this.messagesSubVal)
      
      this.sessionRecordData.splice(index, 1, { "role": "assistant", "content": '', "loading": true, "c2netNameTips": '', "c2netName": '' })
      const params = {
        flag:true,
        chat_id: this.generateName(this.loginName),
        model: this.params.model
      }
      const messagesCopy = Array.from(this.messagesSubVal)
      params.messages = messagesCopy.splice(0,index)
      
      // this.params.flag = true
      // this.params.chat_id = this.generateName(this.loginName)
      // // this.params.messages.pop()
      console.log("params.messages",params.messages)
      this.chat(params,index)
      
    },
    chat(customParams = null, targetIndex = null) {
      this.chatFlag = true
      
      // 创建 AbortController
      this.abortController = new AbortController();
      let vm = this
      llmChat(customParams, { signal: this.abortController.signal }).then((response) => {

        const reader = response.body.getReader();
        const decoder = new TextDecoder()
        let popData = this.sessionRecordData[this.sessionRecordData.length - 1];
        if(targetIndex!==null){
            popData = this.sessionRecordData[targetIndex]
        }
        let isDoneTriggered = false
        let c2netName = ''
        let buffer = ''; // 新增缓冲区
        const ERROR_MARKER = '[ERROR]';
        const ERROR_MARKER_LENGTH = ERROR_MARKER.length;
        const processBinaryData = async () => {
          try {
            while (true) {
              const { done, value } = await reader.read();
              if (done) {
                if (isDoneTriggered) {
                    // 处理流结束时的剩余内容
                    const errorIndex = buffer.indexOf(ERROR_MARKER);
                    if (errorIndex !== -1) {
                        popData.content += buffer.substring(0, errorIndex);
                        buffer = buffer.substring(errorIndex + ERROR_MARKER_LENGTH);
                        this.$message({ type: 'error', message: this.c2netErrMsg });
                    }
                    // 添加剩余内容
                    popData.content += buffer;
                    buffer = '';
                }
                buffer = '';
                vm.chatFlag = false
                if(targetIndex===null){
                  if(!popData.content){
                    this.params.messages.pop()
                  }else{
                    this.params.messages.push({  role: "assistant", content: popData.content })
                  }
                }else{
                  
                  this.params.messages.splice(targetIndex,1,{  role: "assistant", content: popData.content })
                  console.log(this.params.messages)
                }
                
                popData.loading = false
                break;
              }
              let chars = decoder.decode(value)
              buffer += chars; // 将新数据追加到缓冲区
              if (response.status === 503) {
                const err = JSON.parse(chars)
                this.$message({
                  type: 'error',
                  message: err.error,
                });
                this.prompt = this.params.messages[this.params.messages.length-1].content
                vm.chatFlag = false
                popData.loading = false;
                popData.content = err.error
                break
              }
              // 1. 处理[DONE]标记（仅在首次触发）
              if (!isDoneTriggered) {
                const doneIndex = buffer.indexOf('[DONE]');
                if (doneIndex !== -1) {
                    isDoneTriggered = true;
                    c2netName = buffer.substring(0, doneIndex);
                    const regex = new RegExp(vm.regexConfig.pattern);
                    let match = c2netName.match(regex);
                    if (match) {
                        let result = vm.regexConfig.template;
                        for (let i = 1; i < match.length; i++) {
                          result = result.replace(new RegExp(`\\$${i}`, 'g'), match[i]);
                        }
                        console.log("result",result)
                        popData.c2netNameTips = result
                        popData.c2netName = match[2]
                    } else {
                        console.log("字符串格式不匹配");
                    }
                    
                    // 分割[DONE]后的内容，并清空缓冲区
                    const postDone = buffer.substring(doneIndex + '[DONE]'.length);
                    buffer = postDone; // 保留后续数据供后续处理
                }
              }

              // [DONE] 触发后的处理
              if (isDoneTriggered) {
                  let hasError = false;
                  do {
                      const errorIndex = buffer.indexOf(ERROR_MARKER);
                      if (errorIndex === -1) break;
                      // 发现错误标记
                      popData.content += buffer.substring(0, errorIndex);
                      buffer = ''
                      this.$message({ type: 'error', message: this.c2netErrMsg });
                      hasError = true;
                  } while (false);

                  if (!hasError) {
                      // 保留可能形成跨块标记的部分（保留最后 N-1 个字符）
                      const maxKeep = ERROR_MARKER_LENGTH - 1;
                      const keepLength = Math.min(maxKeep, buffer.length);
                      const output = buffer.substring(0, buffer.length - keepLength);
                      popData.content += output;
                      buffer = buffer.substring(buffer.length - keepLength);
                  }
              }
              this.scrollBottom()
            }
          } catch (error) {
            if (error.name === 'AbortError') {
              // 主动中止时忽略错误
              return;
            }
            throw error;
          }
        };
        processBinaryData();
      }).catch((err) => {
        vm.chatFlag = false;

        const popData = vm.sessionRecordData[vm.sessionRecordData.length - 1];
        popData.content += buffer;
        popData.loading = false;
        if (err.name === 'AbortError') {
          // 用户主动停止
          vm.params.messages.push({ role: 'assistant', content: popData.content });
          vm.scrollBottom();
        } else {
          this.$message({ type: 'error', message: err.message });
        }
      })
    },
    scrollBottom() {
      this.$nextTick(() => {
        if (!this.isAutoScroll) return;
        const scrollDom = this.$refs.chatContainer;
        scrollDom.scrollTop = scrollDom.scrollHeight;
      })
    },
    onScroll() {
      const scrollDom = this.$refs.chatContainer;
      const scrollTop = scrollDom.scrollTop;
      const offsetHeight = scrollDom.offsetHeight;
      const scrollHeight = scrollDom.scrollHeight;
      // 当滚动到底部，设置 isAutoScroll 为 true
      if (scrollTop + offsetHeight + 1 >= scrollHeight) {
        this.isAutoScroll = true;
      } else {
        // 否则，用户正在手动滑动，设置为 false，停止自动滚动
        this.isAutoScroll = false;
      }
      this.delayTimer && clearTimeout(this.delayTimer)
      this.delayTimer = setTimeout(() => { this.isAutoScroll = true }, 3500)
    },
    getConfig(){
      getPromoteData(this.promotePath).then((res)=>{
          const data = JSON.parse(res.data);
          this.c2netList = data.c2netList
          this.c2netErrMsg = data.errorMessage
          this.c2netContent = data.headerContent
          this.imgContent = data.imgList
          console.log(data.regexConfig['auto'])
          this.regexConfig = data.regexConfig['auto']
          // if(this.modelName === 'DeepSeek'){
          //   this.regexConfig = data.regexConfig['auto']
          // }else{
          //    this.regexConfig = data.regexConfig['default']
          // }
      }).catch((err)=>{
          this.$message.error(err)
      })
  },
  },
  mounted() {
    const urlParams = new URLSearchParams(location.search)
    if (urlParams.has('name')) {
      this.modelName = urlParams.get('name')
    }
    this.getConfig()
  },
  beforeMount() {
    const isLogin = !!document.querySelector('meta[name="_uid"]');
    if (isLogin) {
      this.loginName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext')
    }
  }
}
</script>

<style lang="less" scoped>
* {
  box-sizing: border-box;
}
.chat-wrap{
  display: flex;
  flex-direction: column;
  height: calc(100vh - 60px);
  margin-bottom: -40px;
  .model-text-wrapper {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    min-height: 80px;

    .model-title {
      color: rgba(16, 16, 16, 1);
      font-size: 28px;
      margin-bottom: 0.5rem;
    }

    .model-desc {
      color: rgba(136, 136, 136, 0.87);
      font-size: 14px;
      height: 20px;
    }
  }
  .main-container{
    position: relative;
    display: flex;
    justify-content: center;
    height: calc(100% - 80px);
    .main-left-menu{
      width: 280px;
      height: calc(100% - 12px);
      overflow-y: auto; 
      background-color: rgba(255,255,255,1);
      border: 1px solid rgba(229,231,235,1);
      border-radius: 6px;
      flex-shrink: 0;
      .menu-wrap{
        display: flex;
        flex-direction: column;
        height: 100%;
        .sidbar-header{
          flex-shrink: 0;
          .sidbar-title{
            height: 60px;
            font-size: 18px;
            align-items: center;
            display: flex;
            background: url('/img/model/c2net-header.png');
            color: rgba(16,16,16,1);
            div{
              margin-left: 12px;
            }
          }
          .sidbar-sub-title{
            line-height: 20px;
            color: rgba(16,16,16,1);
            margin: 0 12px 16px 24px;
          }
        }
        .sidbar-body{
          flex: 1;
          overflow: auto;
          /* 整个滚动条 */
          &::-webkit-scrollbar {
            width: 4px !important; /* 控制垂直滚动条的宽度 */
          }

          /* 水平滚动条的高度 */
          &::-webkit-scrollbar:horizontal {
            height: 4px !important; /* 控制水平滚动条的高度 */
          }
          &::-webkit-scrollbar-thumb {
            background: #ccc;     /* 更浅的滑块颜色 */
          }
          .sidbar-main{
            margin-left:24px;
            .sidebar-item{
              line-height: 32px;
              color: rgba(0,71,169,1);
              display: block;
            }
          }
          
        }
        .sidbar-footer{
          margin: 12px 0;
          padding: 0 24px;
          flex-shrink: 0;
        }
      }
      
    }
    .main-code-wrap{
      width: 100%;
      max-width: 1000px;
      .main-code-body {
        display: flex;
        width: 100%;
        overflow-x: auto;
        overflow-y: hidden;
        height: calc(100vh - 140px);
        position: relative;
        overflow: hidden;

        .output-text {
          overflow: auto;
          flex: 1 1;
          margin-bottom: 152px;
          
          .output-content {
            height: 100%;
            font-size: 14px;
            padding: 0 32px;
            padding-top: 0;
            overflow: auto;
            word-break: break-all;
            &::-webkit-scrollbar {
              width: 0 !important; /* Chrome/Safari/Edge */
            }
            
            scrollbar-width: none; /* Firefox */
            -ms-overflow-style: none; /* IE/Edge */
            .out-header-wrap{
              
              margin-bottom: 20px;
              .c2net-support-list{
                display: none;
                align-items: center;
                justify-content: center;
                background-color: rgba(250,140,22,1);
                border-radius: 6px;
                line-height: 14px;
                text-align: center;
                border: 1px solid rgba(104,83,155,1);
                color: rgba(251,251,251,1);
                margin-right: 4px;
                flex-direction: column;
              }
              .output-welcome-banner {
                padding: 12px;
                background-color: aliceblue;
              }
            }
            .user-input-wrapper {
              margin-left: 60px;
              margin-bottom: 24px;
              display: flex;
              align-items: top;
              flex-direction: row-reverse;
              white-space: pre-line;

              .user-input {
                display: inline-block;
                background: rgba(25, 117, 255, .1);
                color: #43436b;
                border-radius: 6px;
                padding: 8px;
                line-height: 20px;
              }
            }

            .user-output-wrapper {
              margin-right: 60px;
              margin-bottom: 24px;
              padding: 10px;
              display: flex;
              align-items: top;

              .user-output {
                display: inline-block;
                background: antiquewhite;
                color: #43436b;
                border-radius: 6px;
                padding: 8px;
                line-height: 20px;
                position: relative;
                word-break: break-all;

              }
            }
          }
        }

        .ask-text-input {
          position: relative;
          height: 140px;
          background-color: blue;
        }
      }
    }
  }
  
  

  .main-code-input {
    bottom: 0;
    // display: flex;
    // justify-content: center;
    margin-bottom: 12px;
    position: absolute;
    width: 100%;
    z-index: 99;
    padding: 0 32px;

    .input-container {
      flex: 1 1;
      position: relative;
      width: 100%;

      .input-box {
        background: #fff;
        border: 1px solid #2468f2;
        border-radius: 6px;
        width: 100%;

        .ask-input {
          padding: 12px 12px 28px;
          position: relative;

          #code-input {
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

          .opera-submit {
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

      .input-tips {
        color: #999;
        text-align: end;
        font-size: 12px;
        padding-top: 6px;
      }
    }
    .tips-privacy{
      text-align: center;
      font-size: 12px;
      display: none;
    }
  }
}


::v-deep .el-slider__input {
  width: 55px;
}

::v-deep .el-slider__runway.show-input {
  margin-right: 70px;
}

::v-deep .el-input__suffix {
  right: 0px;
}

::v-deep .el-input-number.is-without-controls .el-input__inner {
  padding-left: 10px;
  padding-right: 10px;
}
@media (max-width: 768px){
  .menu-mask{
    position: fixed;
    top: 62px;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 999;
    
  }
  .chat-wrap{
    margin-bottom: -80px;
    
    .model-text-wrapper{
      min-height: 60px;
      .model-title{
        font-size: 16px;
      }
    }
    .main-container{
      .main-left-menu{
        position: fixed;
        height: calc(100% - 62px);
        left: -280px;
        top: 62px;
        z-index: 1000;
        transform: translateX(0);
        transition: transform 0.3s ease;
        &.mobile-menu-active {
          transform: translateX(280px);
        }
      }
      .main-code-wrap
        .main-code-body{
        height: calc(100vh - 130px);
        .output-text{
          .output-content{
            padding: 0;
            .out-header-wrap{
              display: flex;
              font-size: 12px;
              .c2net-support-list{
                width: 40px;
                display: flex;
              }
              .output-welcome-banner{
                flex: 1;
              }
            }
          }
        }
      }
    }
    
    .main-code-input{
      padding: 0;
      margin-bottom: 0;
      .tips-privacy{
        display: block;
      }
    }
  }
}
</style>
