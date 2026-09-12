<template>
  <CommonTipsDialog type="dataset" :promotePath="promotePath" :appendToBody="true"
    :title="$t('datasetObj.codeUseDlgTitle')" :closeText="$t('cloudbrainObj.dialogTips.tips8')">
    <a class="use-dlg-btn">{{ $t('datasetObj.codeUseDlgTriggerTxt') }}</a>
  </CommonTipsDialog>
</template>
<script>

import CommonTipsDialog from '~/components/CommonTipsDialog.vue';
import { lang } from '~/langs';

export default {
  data() {
    return {
      promotePath: `tips/dataset/sdkcode${lang == 'zh-CN' ? '' : '_en'}.md`
    }
  },
  components: { CommonTipsDialog },
  methods: {},
  mounted() { },
}
</script>
<style scoped>
.use-dlg-btn {
  text-decoration: underline;
}
</style>