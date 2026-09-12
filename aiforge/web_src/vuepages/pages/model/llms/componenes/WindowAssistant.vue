<template>
    <div>
        <div class="chat-main-content">
            <img :src="img" alt="" />
            <div class="chat-output-wrapper">
                <div class="chat-container">
                    <div style="overflow: hidden;background-color: transparent;height: auto;">
                        <div class="loading" v-if="useLoading && !htmlContent && data.loading"><i
                                class="el-icon-loading"></i></div>
                        <div class="file-view markdown chat" style="font-size: 14px;" v-html="htmlContent" />
                    </div>  
                </div>        
            </div>
        </div>
        <div v-if="!!htmlContent && data.loading" class="chat-c2net-stop" @click="$emit('stopGeneration')"><el-button style="font-size: 12px;" type="text">{{ $t('modelSquare.stopChat') }}</el-button></div>
        <div class="chat-c2net-tips" v-html="data.c2netNameTips"></div>
    </div>
</template>

<script>
// import MarkdownIt from "markdown-it";
// import markdownItKatexGpt from 'markdown-it-katex-gpt'
// import 'katex/dist/katex.css'
import createRenderer from './render'
const thinkStart = ['<', '<t', '<th', '<thi', '<thin', '<think', '<think']
const thinkStop = ['<', '</', '</t', '</th', '</thi', '</thin', '</think', '</think']
let md = null
export default {
    name: "WindowAssistant",
    props: {
        content: { type: String, required: true },
        img: { type: String, required: true, default: "/img/chatbot.png" },
        data: { type: Object, default: () => { return {} } },
        useLoading: { type: Boolean, default: false }
    },
    data() {
        return {
            htmlContent: "",
        }
    },
    watch: {
        content: {
            deep: true,
            handler(val) {
                if (thinkStart.indexOf(val) > -1) val = ''
                const regex = /^<think>([\S\s]*?)/im;
                const match = val.match(regex);
                if (match) {
                    val = val.replace('<think>', '')
                    const valueList = val.split('</think>')
                    var flag = false
                    thinkStop.map(item => {
                        const lastIndex = valueList[0].lastIndexOf(item)
                        if (!flag && lastIndex > -1 && lastIndex + item.length == valueList[0].length) {
                            valueList[0] = valueList[0].substring(0, lastIndex)
                            flag = true
                        }
                    })
                    const quote = valueList[0].trim() ? valueList[0].trim().split('\n').map(item => '> ' + item).join('\n') : ''
                    val = quote + (valueList[1] || '')
                }
                // this.htmlContent = 'md.render(val)'
                this.htmlContent  = md.render(val)
            }
        },
    },
    methods: {

    },
    mounted() {

    },
    created() {
        md = createRenderer()
    }           
}
</script>

<style scoped lang="less">
.chat-main-content {
    margin-right: 60px;
    padding-top: 10px;
    display: flex;

    img {
        width: 28px;
        height: 28px;
        margin-right: 12px;
    }

    .chat-output-wrapper {
        display: inline-block;
        background: rgba(25, 117, 255, .1);
        color: #43436b;
        border-radius: 6px;
        padding: 8px;
        line-height: 20px;

        .chat-container {
            display: inline-block;
            line-height: 20px;
            padding: 6px 12px;
            word-break: break-all;
        }
    }
}
.chat-c2net-stop {
    margin-left: 44px;
    font-size: 12px;
    display: flex;
    align-items: center;
}
.chat-c2net-tips {
    margin-left: 44px;
    font-size: 12px;
    color: rgba(187, 187, 187, 1);
    margin-bottom: 8px;
    ::v-deep span{
        color: rgba(0, 71, 169, 0.5);
        font-weight: 400;
    }
}
.file-view {
    border-bottom: 1px solid rgba(1, 145, 255, 0.1);
    word-break: break-word;
    &:first-child {
        padding-bottom: 1rem;
    }

    &:last-child {
        padding-bottom: 0;
        border: none;
    }
}
@media (max-width: 768px){
    .chat-main-content{
        margin-right: 32px;
        padding-top: 0;
        .chat-output-wrapper{
            padding: 8px 0;
            .chat-container{
                .chat{
                    padding: 0 !important;
                }
            }
        }
    }
    
}



</style>