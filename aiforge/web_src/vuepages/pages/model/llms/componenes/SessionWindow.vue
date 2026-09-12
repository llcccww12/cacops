<template>
    <div style="width: 100%;position: relative">
        <div v-for="(item,index) in sessionRecordData" :key="index">
            <WindowAssistant v-if="item.role === 'assistant'" :content="item.content" :data="item" :img="img" :useLoading="useLoading" 
              @stopGeneration="$emit('stopGeneration')"></WindowAssistant>
            <WindowUser v-if="item.role==='user'" :content="item.content"></WindowUser>
            <WindowTts v-if="item.role==='tts'" :content="item.content"></WindowTts>
            <div class="chat-op-wrap" v-if="item.role === 'assistant' && !item.loading && opFlag">
                <el-tooltip class="item" effect="dark" :content="$t('copy')" placement="top">
                    <i class="ri-file-copy-line" @click="copy(item.content)"></i>
                </el-tooltip>
                <el-tooltip class="item" effect="dark" :content="$t('modelSquare.refreshChat')" placement="top">
                    <i class="ri-refresh-line" @click="refreshAnswer(index)"></i>
                </el-tooltip>
                <el-tooltip class="item" effect="dark" :content="$t('modelSquare.like')" placement="top">
                    <i :class="[isGood ? 'ri-thumb-up-fill' : 'ri-thumb-up-line']" @click="feedGood(index,true)"></i>
                </el-tooltip>
                <el-tooltip class="item" effect="dark" :content="$t('modelSquare.unlike')" placement="top">
                    <i :class="[isBad ? 'ri-thumb-down-fill' : 'ri-thumb-down-line']" @click="feedBad(index,false)"></i>
                </el-tooltip>
            </div>
        </div>
    </div>
</template>
<script>
import WindowAssistant from "./WindowAssistant.vue";
import WindowUser from "./WindowUser.vue";
import WindowTts from "./WindowTts.vue";
import { llmChatFeedBack } from '~/apis/modules/deepseek';
export default {
    name: "SessionWindow",
    props: {
        img: {
            type: String,
            default: "/img/chatbot.png"
        },
        useLoading: {
            type: Boolean,
            default: false
        },
        opFlag: {
            type: Boolean,
            default: false
        },
    },
    
    components: {
        WindowUser,
        WindowAssistant,
        WindowTts
    },
    data() {
        return {
            sessionRecordData:[],
            copyValue: '',
            isGood: false,
            isBad: false,
        }
    },
    watch: {
    },
    created() {
    },
    mounted() {
    },
    methods: {
        setSessionRecord(val) {
            this.sessionRecordData = val;
        },
        async copyToClipboard(text) {
            console.log("text",text)
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
        async copy(context){
            let val = context.split('</think>\n\n')[1]
            console.log(val)
            const success = await this.copyToClipboard(val);
            if (success) {
                this.$message.success("文本复制成功")
            }
        },
        feedGood(index){
            if (this.isGood){
                this.cancelFeedback(index, null);
            }else{
                this.submitFeedback(index, true);
            }
        },
        feedBad(index){
            if (this.isBad) {
                this.cancelFeedback(index, null);
            } else {
                this.submitFeedback(index, false);
            }
        },


        submitFeedback(index, flag) {
            this.simulateRequest(index, flag, 'submit');
        },
        cancelFeedback( index, flag) {
            this.simulateRequest(index, flag, 'cancel');
        },
        simulateRequest(index, flag, action){
            const params = {
                messages: this.sessionRecordData[index-1].content,
                feedBack: flag,
                info: this.sessionRecordData[index].c2netName
            }
            llmChatFeedBack(params).then((res)=>{
                console.log(res)
                if(res.data.msg==='success'){
                    if(action==='submit'){
                        if(flag){
                            this.isGood = true
                            this.isBad = false
                        }else{
                            this.isBad = true
                            this.isGood = false
                        }
                    }else{
                        this.isGood = false
                        this.isBad = false
                    }
                }
            }).catch((err)=>{
                this.$message.error(err)
            })
        },
        refreshAnswer(index){
            this.isGood = false
            this.isBad = false
            this.$emit('refreshAnswer',index)
        }
        
    },
}
</script>



<style lang="less" scoped>
.chat-op-wrap{
    margin-left: 44px;
    display: flex;
    gap: 12px;
    i{
        cursor: pointer;
        opacity: 0.7;
        font-size: 16px;
    }
}
</style>