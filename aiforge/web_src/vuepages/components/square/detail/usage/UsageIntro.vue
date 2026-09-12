<template>
  <div class="ui container usage-intro">
    <BaseTitle :title="type==='dataset' ? $t('datasetObj.codeUseDlgTitle') : $t('modelObj.codeUseDlgTitle')" class="content">
        <div class="main-box">
            <el-skeleton style=" padding:0 20px;" :rows="10" :loading="loading" animated/>
            <div class="markdown" v-html="content"></div>
        </div>
    </BaseTitle>
  </div>
</template>

<script>
import BaseTitle from '~/components/BaseTitle.vue'
import { getPromoteData, getMarkdownHtml, getSDKCode } from '~/apis/modules/common';
import SparkMD5 from "spark-md5";
import hljs from 'highlight.js';
import { initClipboard, transFileSize } from '~/utils';
import { lang } from '~/langs';
export default {
  name: 'UsageIntro',
  components: {
    BaseTitle,
  },
  props: {
    dataObj: { type: Object, default: () => ({}) },
    type: { type: String, default: 'dataset' },
  },
  data() {
    return {
      loading: true,
      content: '',
      promotePath: `tips/${this.type}/sdkcode${lang == 'zh-CN' ? '' : '_en'}.md`
    };
  },

  methods: {
    sparkMD5Hash(str = '') {
      return SparkMD5.hash(str) + Math.random().toString().replace('0.', '');
    },
    hljsAndInsertCopyButton(htmlStr) {
      const html = document.createElement('div');
      html.innerHTML = htmlStr;
      const codeBlocks = html.querySelectorAll('.code-block')
      for (let i = 0, iLen = codeBlocks.length; i < iLen; i++) {
        const codeBlockI = codeBlocks[i];
        const txt = codeBlockI.textContent.slice(0,-1);
        const codeEl = codeBlockI.querySelector('code');
        const code = codeEl.textContent;
        codeEl.innerHTML = hljs.highlight('python', code).value;
        const copyBtn = document.createElement('div');
        copyBtn.classList = ['copy-btn'];
        if (this.type !== 'modelFile') {
          copyBtn.innerHTML = `<span style="color: #0366d6;cursor: pointer;" class="ui poping inline up clipboard" id="clipboard-${this.sparkMD5Hash(txt)}"
            data-position="top center" data-variation="inverted tiny" data-success="${this.$t('copySuccess')}"
            data-content="${this.$t('copy')}" data-original="${this.$t('copy')}" 
            data-clipboard-text=""><i style="font-size:14px;" class="copy outline icon"></i></span>`;
            copyBtn.querySelector('span').setAttribute('data-clipboard-text', txt);
        } else {
          copyBtn.innerHTML = `<span style="color: #0366d6;cursor: pointer;" class="ui poping inline up clipboard" id="clipboard-${this.sparkMD5Hash(txt)}"
            data-position="top center" data-variation="inverted tiny" data-success="${this.i18n['cloudeBrainMirror']['copy_succeeded']}"
            data-content="${this.i18n['cloudeBrainMirror']['copy']}" data-original="${this.i18n['cloudeBrainMirror']['copy']}" 
            data-clipboard-text=""><i style="font-size:14px;" class="copy outline icon"></i></span>`;
            copyBtn.querySelector('span').setAttribute('data-clipboard-text', txt);
        }
        codeBlockI.outerHTML = `<div class="code-content">${codeBlockI.outerHTML}${copyBtn.outerHTML}</div>`;
      }
      return html.innerHTML;
    },
    getMarkdown(str) {
      getMarkdownHtml(str).then(res => {
        this.loading = false;
        const html = res.data;
        this.content = this.hljsAndInsertCopyButton(html);
        this.$nextTick(() => {
          initClipboard('.base-dlg .clipboard');
        });
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    getContent() {
      if (!this.promotePath) return;
      this.loading = true;
      getPromoteData(this.promotePath).then(res => {
        let contentStr = res.data;
        if (contentStr) {
          let params = this.type == 'dataset' ? { dataset_name: [this.dataObj.name] } : { pretrain_model_name: [this.dataObj.name] };
          getSDKCode(params).then(res => {
            res = res.data;
            if (res.code == 0 && res.data && res.data.code) {
              contentStr += `\n\`\`\`python\n${res.data.code}\`\`\`\n`;
            }
            this.getMarkdown(contentStr);
          }).catch(err => {
            console.log(err);
          })
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
  },
  mounted() {
    this.getContent()
  },
};
</script>
<style lang="less" scoped>
.content{
  margin-top: 3rem;
  box-shadow: none;
  border: 1px solid rgba(224,227,234,1);
  .main-box{
    padding: 3rem 12rem 3rem 2rem;
    .markdown {
      font-size: 14px;
      overflow: auto;
      padding-left: 20px;
      padding-right: 20px;

      /deep/ .code-content {
        position: relative;

        .copy-btn {
          position: absolute;
          right: 4px;
          top: 4px;
        }
      }
    }
  }
}
@media screen and (max-width: 767px) {
/* 当视口宽度 ≤ 767px 时生效 */
  .main-box{
    padding: 24px 0 0 0 !important;
    .markdown{
      padding: 0 0 0 10px !important;
    }
  }
  .usage-intro{
    margin: 0 !important;
  }
}
</style>