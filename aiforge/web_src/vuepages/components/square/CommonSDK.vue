
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
import { getMarkdownHtml } from '~/apis/modules/common';
import { getFileSdkCode } from "~/apis/modules/dataset";
import SparkMD5 from "spark-md5";
import hljs from 'highlight.js';
import { initClipboard, transFileSize } from '~/utils';
export default {
  name: "CommonTipsDialog",
  components: { BaseDialog },
  props: {
    title: { type: String, default: 'Title' },
    appendToBody: { type: Boolean, default: false },
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
        copyBtn.innerHTML = `<span style="color: #0366d6;cursor: pointer;" class="ui poping inline up clipboard" id="clipboard-${this.sparkMD5Hash(txt)}"
            data-position="top center" data-variation="inverted tiny" data-success="${this.$t('copySuccess')}"
            data-content="${this.$t('copy')}" data-original="${this.$t('copy')}" 
            data-clipboard-text=""><i style="font-size:14px;" class="copy outline icon"></i></span>`;
            copyBtn.querySelector('span').setAttribute('data-clipboard-text', txt);
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
    getSdkContent(size) {
      let contentStr
      getFileSdkCode(this.data).then((res) => {
        res = res.data
        if (res.code == 0) {
            contentStr = this.$t('modelObj.model_sdk_tips1',{fileSize:transFileSize(size)})
            contentStr += `\n\`\`\`python\n${res.data.code}\`\`\`\n\n${this.$t('modelObj.model_sdk_tips2')}\n\`\`\`bash\n${res.data.cli}\n\`\`\`\n`;
            contentStr += this.$t('modelObj.model_sdk_tips3')       
          this.getMarkdown(contentStr)
        } else {
          this.$message.error(res.msg)
        }
      }).catch((err) => {
        this.$message.error(err)
      });
    },
    handlerOpen(event, size=0) {
      this.getSdkContent(size)
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
  },
  beforeMount() {
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