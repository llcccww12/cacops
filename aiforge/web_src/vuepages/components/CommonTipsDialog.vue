
<template>
  <div class="common-tips-dlg">
    <div class="trigger-container" @click="handlerOpen">
      <slot></slot>
    </div>
    <BaseDialog :visible="visible" :title="title" :show-close="showClose" @closed="closeDialog"
      :appendToBody="appendToBody" >
      <el-skeleton style=" padding:0 20px;" :rows="10" :loading="loading" animated/>
      <div class="markdown" v-html="content"></div>
      <div slot="footer" class="dialog-footer">
        <el-button class="close-btn" size="default" type="success" @click="closeDialog">
        {{ closeText }}
        </el-button>
      </div>
    </BaseDialog>
  </div>
</template>
<script>
import BaseDialog from '~/components/BaseDialog.vue';
import { getPromoteData, getMarkdownHtml, getSDKCode,getModelFileSDKCode } from '~/apis/modules/common';
import SparkMD5 from "spark-md5";
import hljs from 'highlight.js';
import { initClipboard, transFileSize } from '~/utils';
export default {
  name: "CommonTipsDialog",
  components: { BaseDialog },
  props: {
    title: { type: String, default: 'Title' },
    promotePath: { type: String, default: '' },
    appendToBody: { type: Boolean, default: false },
    type: { type: String, default: '' },
    data: { type: Object, default: () => ({}) },
    closeText: { type: String, default: '' },
  },
  data() {
    return {
      showClose: true,
      visible: false,
      loading: false,
      content: '',
    }
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
          if (this.type == 'model' && this.data.name) {
            getSDKCode({
              pretrain_model_name: [this.data.name],
            }).then(res => {
              res = res.data;
              if (res.code == 0 && res.data && res.data.code) {
                contentStr += `\n\`\`\`python\n${res.data.code}\`\`\`\n`;
              }
              this.getMarkdown(contentStr);
            }).catch(err => {
              console.log(err);
            })
          } else {
            this.getMarkdown(contentStr);
          }
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    },
    getSdkContent(size) {
      let contentStr
      getModelFileSDKCode(this.data).then((res) => {
        res = res.data
        if (res.code == 0) {
          if (this.type == 'model') {
            // let helpAddress = this.data.filename ? 'download_model_file' : 'download_model'
            contentStr = this.$t('modelObj.model_sdk_tips1',{fileSize:transFileSize(size)})
            contentStr += `\n\`\`\`python\n${res.data.code}\`\`\`\n\n${this.$t('modelObj.model_sdk_tips2')}\n\`\`\`bash\n${res.data.cli}\n\`\`\`\n`;
            contentStr += this.$t('modelObj.model_sdk_tips3')
          } else {
            contentStr = this.i18n['modelObj']['model_sdk_tips1'].format(transFileSize(size))
            contentStr += `\n\`\`\`python\n${res.data.code}\`\`\`\n\n${this.i18n['modelObj']['model_sdk_tips2']}\n\`\`\`bash\n${res.data.cli}\n\`\`\`\n`;
            contentStr += this.i18n['modelObj']['model_sdk_tips3']
          }          
          
          this.getMarkdown(contentStr)
        } else {
          this.$message.error(res.msg)
        }
      }).catch((err) => {
        this.$message.error(err)
      });
    },
    handlerOpen(event, flag = false,size=0) {
      if (!flag) {
        this.getContent();
      } else {
        this.getSdkContent(size)
      }
      
      this.visible = true;
    },
    closeDialog() {
      this.visible = false;
      setTimeout(() => {
        this.content = '';
      }, 300);
    }
  },
  mounted() {
    this.i18n = window.i18n;
  },
  beforeMount() {
    // this.i18n = window.i18n;
  }
}
</script>

<style scoped lang="less">
.trigger-container {
  display: inline-block;
}

.markdown {
  font-size: 14px;
  max-height: 65vh;
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

.dialog-footer {
  text-align: center;

  .close-btn {
    color: #fff;
    background-color: #21ba45;
    border-color: #21ba45;
    font-size: 1rem;
    font-weight: 700;
  }
}
</style>
<style>
.ui.popup {
  z-index: 2003;
}
</style>