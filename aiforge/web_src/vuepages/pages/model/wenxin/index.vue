<template>
    <div class="ui container wenxin_model_wrap" style="margin-top: 2rem;">        
        <div>
            <div class="wenxin_title_wrap">
                <h1>基于国产算力的ERNIE-ViLG AI 作画大模型</h1>
                <p>全球规模最大的中文跨模态生成模型，可通过自然语言实现图像生成与编辑，由鹏城云脑Ⅱ提供算力支持 。</p>
            </div>
        </div>
        <div class="wenxin_search_wrap">
            <div style="width: 94%;">
                <input class="wenxin_search_input" ref="inputSearch" @keyup.enter="erniePaint" v-model="searchValue" autocomplete="off" maxlength="100" type="text" placeholder="请用中文输入Prompt，参考形式：画面主体，细节描述，修饰词"></input>
            </div>
        </div>
        <div class="wenxin_tips_wrap">
            <div class="wenxin_tip_box">
                <span style="flex:1"> <span style="color: #888;">Prompt示例：</span><span style="cursor: pointer;" @click="copyText(tipCase)">{{ tipCase }}</span></span>
                <span class="wenxin_replace_tip" @click="changeCase">换示例</span>
            </div>
            <div class="wenxin_button_box">
                <button class="wenxin_button_content"  @click="erniePaint" :disabled="loading" :style="{cursor:(loading?'not-allowed':'pointer')}">
                    <span v-if="!loading">立即生成</span>
                    <div v-else class="ui active primary inline loader"></div>
                </button>
            </div>
        </div>
        <div v-if="showResult">
            <div class="wenxin_wait" v-if="resultLoading">
                正在生成中，预计需要 {{waitTime}}s
            </div>
            <div class="wenxin_result_wrap scale-in-ver-top">
                
                <div class="wenxin_result_left">
                    <div>
                        <div class="content">
                            <div class="desc_header">创意灵感：</div>
                            <div class="desc_text">{{searchValue}}</div>
                        </div>
                        <div class="content">
                            <div class="desc_header">规格：</div>
                            <div class="desc_text">1024 x 1024</div>
                        </div>
                    </div>
                    <div class="wenxin_use_content">
                        <span> <span>{{userCount}}/200</span> <span style="color:#ff5e00">限量体验200次</span> </span>
                    </div>
                </div>
                <div class="wenxin_result_right">
                    <img src="/img/model/loading.gif" style="width: 100%;height:100%;border-radius: 10px;" v-if="resultLoading">
                    <div v-else>
                        <div class="gallery">
                            <img :src="'data:image/Jpeg;base64,'+resultImg" alt="" style="border-radius: 1rem;z-index: 9;">
                        </div>
                        <span class="wenxin_mask">
                            <span class="wenxin_down_wrap" @click="downloadImg">
                                <span class="download_img">
                                    下载
                                </span>
                            </span>
                        </span>
                    </div>
                </div>

            </div>
        </div>
        
        <div class="wenxin_example_wrap gallery">
            <p>ERNIE-ViLG应用场景示例</p>
            <div class="ui stackable four column grid">
                <div class="column">
                    <div class="ui segment wenxin_img_box">
                        <img src="/img/model/wenxin1.jpeg">
                    </div>
                    <div class="wenxin_img_desc">
                        <span class="wenxin_img_tips" :title="caseText[0]">{{caseText[0]}}</span>
                        <span class="wenxin_img_copy" :data-text="caseText[0]">复制</span>
                    </div>
                </div>
                <div class="column">
                    <div class="ui segment wenxin_img_box">
                        <img src="/img/model/wenxin2.jpeg">
                    </div>
                    <div class="wenxin_img_desc">
                        <span class="wenxin_img_tips" :title="caseText[1]">{{caseText[1]}}</span>
                        <span class="wenxin_img_copy" :data-text="caseText[1]">复制</span>
                    </div>
                </div>
                <div class="column">
                    <div class="ui segment wenxin_img_box">
                        <img src="/img/model/wenxin3.jpeg">
                    </div>
                    <div class="wenxin_img_desc">
                        <span class="wenxin_img_tips" :title="caseText[2]">{{caseText[2]}}</span>
                        <span class="wenxin_img_copy" :data-text="caseText[2]">复制</span>
                    </div>
                </div>
                <div class="column">
                    <div class="ui segment wenxin_img_box">
                        <img src="/img/model/wenxin4.jpeg">
                    </div>
                    <div class="wenxin_img_desc">
                        <span class="ui wenxin_img_tips" :title="caseText[3]">{{caseText[3]}}</span>
                        <span class="wenxin_img_copy" :data-text="caseText[3]">复制</span>
                    </div>
                </div>
            </div>
        </div>
        <div class="wenxin_alert_wrap">
            <span> <a href="/home/model_privacy"  target="_blank">《免责声明和服务使用规范》</a></span>
            <span>访问更多关于 <a href="https://wenxin.baidu.com/ernie-vilg" target="_blank"> ERNIE-ViLG AI</a> 作画大模型的内容</span>
        </div>
    </div>
</template>
<script>
import Clipboard from 'clipboard';
import base64Img,{ ERNIE_CASE, ERNIE_CASE_IMG,Thread } from './constant.js'
import { postERNIEPaintNew,postERNIEPaintResult,getERNIEPaintCount } from "~/apis/modules/desensitization";
import { Message } from "element-ui";

export default {
    
    data() {
        return {
            caseIndex: 0,
            caseText: ERNIE_CASE_IMG,
            searchValue: "",
            loading: false,
            showResult: false,
            resultLoading: false,
            resultImg: '',
            userCount: 0,
            startFlag:false,
            resultImgName: '',
            waitTime:0
        }
    },
    computed: {
        tipCase() { 
            return ERNIE_CASE[this.caseIndex]
        }
    },
    methods: {
        initCopy() {
            const clipboard = new Clipboard('.wenxin_img_copy', {
                text: (e) => {
                    return e.getAttribute('data-text');
                },
            });
            clipboard.on('success', (e) => {
                Message.success('复制成功')
            });
            clipboard.on('error', (e) => {
                Message.error('复制失败')
            });
        },
        changeCase() { 
            if (this.caseIndex === ERNIE_CASE.length-1) { 
                this.caseIndex=0
            } else {
                this.caseIndex++
            }
        },
        erniePaint() {
            if (!this.startFlag) {
                const params = new URLSearchParams(location.search)
                if (params.has('prompt')) { 
                    window.location.href = `/user/login?redirect_to=${encodeURIComponent(location.href)}`
                } else {
                    window.location.href = `/user/login?redirect_to=${encodeURIComponent(location.href)}?prompt=${this.searchValue}`;
                }
                
                return;
            }
            if (!this.searchValue.split(" ").join("").length) { 
                Message.error('输入不能为空')
                return
            }
            this.loading = true
            this.resultLoading = true
            getERNIEPaintCount().then((res) => { 
                if (res.data >= 200) { 
                    Message.error('已达最大生成图片量')
                    this.resultLoading = false
                    this.loading = false
                    return
                }
                this.userCount = res.data
                this.getPaintResult()
            }).catch((err) => {
                Message.error(err)
            })
            
         },
        copyText(value) {
            this.searchValue = value
            
        },
        getPaintResult() {
            let params = { 'textDesc': this.searchValue }
            this.showResult = true
            postERNIEPaintNew(params).then((res) => { 
                
                if (res.data.result == -1) {
                    Message.error(res.data.msg)
                    this.showResult = false
                    this.loading = false
                    this.resultLoading = false
                    this.resultLoading = false
                    return
                } else { 
                    let { result, wait } = res.data
                    this.waitTime = wait
                    this.main(result,wait,params.textDesc)
                }
                
            }).catch((err) => { 
                Message.error(err)
                this.showResult = false
            })

        },
        main(id, count,textDesc) { 
            // 示例2- 通过轮询次数结束轮询
            let vm = this
            let countRequest=0
            const thread = new Thread({
                start: function () {
                    postERNIEPaintResult({ id: id }).then((res => { 
                        if (res.data.Status === 0) { 
                            if (res.data.Picture) { 
                                vm.resultImg = res.data.Picture
                                vm.resultImgName = textDesc
                                vm.resultLoading = false
                                vm.loading = false
                                thread.stop()
                            }
                        } else {
                            Message.error(res.data.Picture||'发生错误')
                            vm.showResult = false
                            vm.resultLoading = false
                            vm.loading = false
                            thread.stop()
                        } 
                        countRequest++
                    })).catch((err) => { 
                        Message.error(err)
                        vm.resultLoading = false
                        vm.loading = false
                    })
                },
                stop: function () {
                    if (countRequest === Math.ceil(count / 2) - 1) { 
                        Message.error('请求超时，请稍后再试')
                        vm.showResult = false
                        vm.resultLoading = false
                        vm.loading = false
                    }
                },
                number: Math.ceil(count/2), //这里是轮询次数配置，不配置默认无线轮询
                time: 2000 //这里是轮询的时间 不配置默认 300ms
            })
            // 开始轮询
            thread.run();
            
        },
        downloadImg() { 
            
            let time = new Date()
            let nowTime = base64Img.timestampToTime(time.toLocaleString('en-US', { hours12: false }).split(' '))
            let fileName = nowTime + '_' + this.resultImgName
            base64Img.downloadFile(this.resultImg,fileName,'jpeg')
        }
    },
    mounted() {
        this.initCopy()
        const signEl = document.querySelector("#isSignd");
        this.startFlag = !!signEl.getAttribute("data-sign");
        const params = new URLSearchParams(location.search)
        if (params.get('prompt')) { 
            this.searchValue = params.get('prompt')
            window.history.pushState("", document.title, location.pathname);
        }
    },  
}
</script>