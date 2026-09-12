<template>
  <div class="preview-container markdown" :class="isTaskDetail ? 'task-detail' : ''">
    <div class="markdown-content">
      <div class="content-c">
        <p v-html="$t('cloudbrainObj.generalTaskSdkCodeTip1')"></p>
        <pre
          class="code-block"><code class="chroma language-text" v-html="$t('cloudbrainObj.generalTaskSdkCodeTip2')"></code></pre>
        <p v-html="$t('cloudbrainObj.generalTaskSdkCodeTip3')"></p>
        <pre class="code-block"><code class="chroma language-text">🎉 You're ready to go live at https://be22fe9f-d140-474c-81f4-30fb068157b5.tunnel.paracloud.com =&gt; http://127.0.0.1:7860
</code></pre>
        <p v-html="$t('cloudbrainObj.generalTaskSdkCodeTip4')"></p>
      </div>
    </div>
    <div class="segment-line"></div>
  </div>
</template>

<script>
export default {
  name: 'GeneralTaskCodeTips',
  props: {
    isTaskDetail: { default: false },
  },
  data() {
    return {};
  },
  methods: {
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.preview-container {
  .markdown-content {
    font-size: 14px;

    p {
      font-size: 14px;
      margin-bottom: 14px;
    }

    .code-block {
      font-size: 85%;
    }
  }

  .segment-line {
    border-top: 1px solid rgba(34, 36, 38, .15);
    border-bottom: 1px solid rgba(255, 255, 255, .1);
    font-size: 1rem;
    line-height: 1;
    height: 0;
    margin-bottom: 12px !important;
  }
}

.task-detail.preview-container {
  .markdown-content {
    .title-h {
      color: #8a8e99;
      font-size: 12px;
      font-weight: 700;
      margin-bottom: 8px;
    }

    .content-c {
      position: relative;
      border-radius: 6px;
      background-color: rgba(240,240,240,1);
      padding: 0.6em 0.8em;
      overflow: auto;
      p {
        font-size: 13px;
        margin-bottom: 10px;
      }
    }

    .segment-line {
      display: none;
    }
  }
}
</style>
