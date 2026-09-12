<template>
  <div>
    <div v-if="showDebugTips" class="debug-task-tips" v-html="$t('cloudbrainObj.debugTaskTimeLimitTips')"></div>
    <GeneralTaskCodeTips v-if="showGeneralTaskCodeTip" :hasCode="codeContent"></GeneralTaskCodeTips>
    <div class="title" style="font-weight:bold">
      <p><span v-html="$t('cloudbrainObj.sdkCodeTip1')"></span></p>
    </div>
    <div class="title">
      <p v-show="codeContent"><span>{{ $t('cloudbrainObj.sdkCodeTip2') }}</span></p>
    </div>
    <div class="content" v-show="codeContent">
      <div class="code-c">
        <div class="code-content markdown">
          <pre><code class="python hljs" v-html="codeHtml"></code></pre>
        </div>
        <div class="copy-btn sdk-code-copy">
          <a href="javascript:;" class="ui poping up clipboard" :id="`clipboard-${clipboardIDStr}`"
            data-position="top center" data-variation="inverted tiny" :data-success="$t('copySuccess')"
            :data-content="$t('copy')" :data-original="$t('copy')" :data-clipboard-text="codeContent">
            <i style="font-size:14px;" class="copy outline icon"></i>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import hljs from 'highlight.js';
import { getSDKCode } from '~/apis/modules/common';
import { initClipboard } from '~/utils';
import GeneralTaskCodeTips from './GeneralTaskCodeTips.vue'

export default {
  name: 'SDKCode',
  props: {
    pageConfigs: { type: Object, required: true, },
    formConfigs: { type: Object, required: true, },
    data: { type: Object, required: true },
    specConfigs: { type: Object, required: true },
  },
  components: { GeneralTaskCodeTips },
  data() {
    return {
      codeContent: '',
      clipboardIDStr: Math.random().toString().replace('0.', ''),
      delayTimer: null,
    };
  },
  watch: {
    data: {
      immediate: true,
      deep: true,
      handler(newVal) {
        this.delayTimer && clearTimeout(this.delayTimer);
        this.delayTimer = setTimeout(() => {
          this.getCode();
        }, 1500);
      }
    },
  },
  computed: {
    codeHtml() {
      return hljs.highlight('python', this.codeContent).value;
    },
    showGeneralTaskCodeTip() {
      let specLen = 0;
      for (let key in this.specConfigs.specs) {
        specLen += this.specConfigs.specs[key].length;
      }
      return this.pageConfigs.taskType == 'GENERAL' && specLen;
    },
    showDebugTips() {
      return this.pageConfigs.taskType == 'DEBUG';
    }
  },
  methods: {
    getCode() {
      const job_type = this.pageConfigs.taskType;
      const compute_source = this.pageConfigs.computerResouce;
      const cluster_type = this.pageConfigs.cluster;
      const datasetList = this.data.dataset;
      const modelList = this.data.model;
      const runParameterList = this.data.runParameters;
      const pretrain_model_name = modelList.map(itm => itm.name);
      const dataset_name = datasetList.map(itm => itm.name);
      const param_key = runParameterList.filter(itm => itm.label.trim()).map(itm => itm.label.trim());
      const visualize_required = this.data.visualizeRequired;
      const params = {
        pretrain_model_name,
        dataset_name,
        param_key,
        job_type,
        compute_source,
        cluster_type,
        visualize_required,
      };
      const curParamsStr = JSON.stringify(params);
      if (this.paramsStr == curParamsStr) return;
      this.paramsStr = curParamsStr;
      getSDKCode(params).then(res => {
        res = res.data;
        if (res.code == 0 && res.data) {
          this.codeContent = res.data.code || '';
          this.$nextTick(() => {
            initClipboard('.sdk-code-copy.copy-btn .clipboard');
          });
        }
      }).catch(err => {
        console.log(err);
      })
    },
    check() {
      return true;
    },
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.debug-task-tips {
  border: 1px solid rgb(242, 113, 28);
  background: rgba(242, 113, 28, 0.05);
  font-weight: 700;
  padding: 6px;
  margin-bottom: 10px;
  line-height: 22px;

  /deep/ span {
    color: rgba(255, 37, 37, 1);
  }
}

.title {
  margin-bottom: 10px;
}

.code-c {
  position: relative;
  font-size: 14px;

  .code-content {
    position: relative;
    overflow: auto;
    font-size: 14px;
  }

  .copy-btn {
    position: absolute;
    top: 0.5em;
    right: 0.8em;
  }
}
</style>
