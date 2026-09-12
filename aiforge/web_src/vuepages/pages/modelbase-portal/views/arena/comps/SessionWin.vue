<template>
  <div class="session-win">
    <div class="header" :style="headBg ? `background:${headBg}` : ''">{{ model.name }}</div>
    <div class="body" ref="chatContainer">
      <div v-for="(item, index) in sessionRecordData" :key="index">
        <SessionAssistant v-if="item.role === 'assistant'" :model="model" :content="item.content" :data="item"
          :img="model.icon || img" :useLoading="useLoading" @stopGeneration="stopGeneration"></SessionAssistant>
        <SessionUser v-if="item.role === 'user'" :content="item.content" :username="loginName" :avatar="userAvatar">
        </SessionUser>
        <div class="chat-op-wrap" v-if="item.role === 'assistant' && !item.loading && opFlag && item.content">
          <el-tooltip class="item" effect="dark" :content="$t('copy')" placement="top">
            <i class="ri-file-copy-line" @click="copy(item.content)"></i>
          </el-tooltip>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import SessionUser from './SessionUser.vue';
import SessionAssistant from './SessionAssistant.vue';
import { llmChat } from '~/apis/modules/deepseek';
import dayjs from 'dayjs';

export default {
  props: {
    headBg: { type: String, default: '' },
    model: { type: Object, default: () => { } },
    config: { type: Object, default: () => { } },
    img: {
      type: String,
      default: "/img/chatbot.png"
    },
    useLoading: {
      type: Boolean,
      default: true
    },
    opFlag: {
      type: Boolean,
      default: true
    },
  },
  data() {
    return {
      sessionRecordData: [],
      isAutoScroll: true,
      chatFlag: false,
      modelName: '',
      loginName: '',
      userAvatar: '',
      params: {
        messages: [],
        model: '',
        chat_id: '',
        flag: false
      },
      abortController: null,
      regexConfig: {},
      c2netErrMsg: '',
    };
  },
  components: {
    SessionUser,
    SessionAssistant,
  },
  watch: {
    sessionRecordData: {
      handler(val) {
        if (val != null) {
          this.$nextTick(() => {
            this.scrollBottom();
          })
        }
      },
      deep: true,
    },
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
    clearChat() {
      this.sessionRecordData = []
      this.params.messages = []
    },
    async submit(prompt) {
      this.prompt = prompt
      this.modelName = this.model.name
      this.isAutoScroll = true
      this.scrollBottom()
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
      this.params.messages.push(newMessage)
      this.prompt = ''
      this.sessionRecordData.push({ "role": "user", "content": this.params.messages[this.params.messages.length - 1].content })
      this.sessionRecordData.push({ "role": "assistant", "content": '', "loading": true, "c2netNameTips": '', "c2netName": '' })
      this.params.model = this.modelName === 'DeepSeek' ? 'auto' : this.modelName
      this.chat(this.params)
    },
    // 停止生成
    stopGeneration() {
      if (this.abortController) {
        this.abortController.abort(); // 中止请求和流
        this.abortController = null;
        this.chatFlag = false;
        this.$emit('finish')
        const popData = this.sessionRecordData[this.sessionRecordData.length - 1];
        popData.loading = false;
      }
    },
    refreshAnswer(index) {
      this.sessionRecordData.splice(index, 1, { "role": "assistant", "content": '', "loading": true, "c2netNameTips": '', "c2netName": '' })
      const params = {
        flag: true,
        chat_id: this.generateName(this.loginName),
        model: this.params.model
      }
      const messagesCopy = Array.from(this.messagesSubVal)
      params.messages = messagesCopy.splice(0, index)
      this.chat(params, index)
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
        if (targetIndex !== null) {
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
                vm.$emit('finish')
                if (targetIndex === null) {
                  if (!popData.content) {
                    this.params.messages.pop()
                  } else {
                    this.params.messages.push({ role: "assistant", content: popData.content })
                  }
                } else {
                  this.params.messages.splice(targetIndex, 1, { role: "assistant", content: popData.content })
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
                this.prompt = this.params.messages[this.params.messages.length - 1].content
                vm.chatFlag = false
                vm.$emit('finish')
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
                    // console.log("result", result)
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
            }
            this.scrollBottom()
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
        vm.$emit('finish')
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
    async copyToClipboard(text) {
      try {
        if (navigator.clipboard) {
          await navigator.clipboard.writeText(text);
          return true;
        }
        // 降级方案
        const textarea = document.createElement('textarea');
        textarea.value = text;
        textarea.style.position = 'fixed';
        document.body.appendChild(textarea);

        if (navigator.userAgent.match(/iphone|ipad|ipod/i)) {
          textarea.contentEditable = true;
          textarea.readOnly = true;
          const range = document.createRange();
          range.selectNodeContents(textarea);
          const selection = window.getSelection();
          selection.removeAllRanges();
          selection.addRange(range);
          textarea.setSelectionRange(0, 999999);
        } else {
          textarea.select();
        }

        document.execCommand('copy');
        document.body.removeChild(textarea);
        return true;
      } catch (err) {
        console.error('复制操作失败:', err);
        return false;
      }
    },
    async copy(context) {
      let val = context.split('</think>\n\n')[1] || context
      const success = await this.copyToClipboard(val);
      if (success) {
        this.$message.success("文本复制成功")
      }
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
  },
  beforeMount() {
    const isLogin = !!document.querySelector('meta[name="_uid"]');
    if (isLogin) {
      this.loginName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext')
      this.userAvatar = document.querySelector('div.userAvatar').getAttribute('src')
    }
  },
  mounted() {
    if (this.modelName === 'DeepSeek') {
      this.regexConfig = this.config.regexConfig['auto']
    } else {
      this.regexConfig = this.config.regexConfig['default']
    }
    this.c2netErrMsg = this.config.errorMessage
  },
  beforeDestroy() {
    this.stopGeneration();
  },
};
</script>

<style scoped lang="less">
.session-win {
  height: 100%;
  overflow-y: auto;
  position: relative;

  .header {
    height: 40px;
    background: rgba(214, 219, 246, 0.3);
    color: rgb(0, 102, 255);
    font-size: 14px;
    padding: 0px 0px 0px 20px;
    display: flex;
    align-items: center;
    border-bottom: 1px solid rgba(16, 16, 16, 0.1);
    font-weight: 700;
    font-family: PingFangSC;
  }

  .body {
    height: calc(100% - 40px);
    overflow-y: auto;
    padding: 25px;

    .chat-op-wrap {
      margin-left: 44px;
      display: flex;
      gap: 12px;
      margin-bottom: 10px;

      i {
        cursor: pointer;
        opacity: 0.7;
        font-size: 16px;
      }
    }
  }
}

@media (max-width: 768px) {
  .session-win {
    .body {
      padding: 12px;
      padding-bottom: 0px;
    }
  }
}
</style>
