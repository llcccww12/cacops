<template>
  <div class="ui container center" style="min-height:80%;">
    <div class="ui basic very padded segment">
      <img class="ui centered medium image" src="/img/icon-403@2x.png">
      <h2>{{ $t('emptyPage') }} </h2>
      <p v-html="$t('emptyPageDescr')"></p>
    </div>
  </div>
</template>

<script>
export default {
  name: "NotFound",
  props: {
    visible: { type: Boolean, default: false },
  },
  data() {
    return {};
  }
};
</script>
<style scoped lang="less"></style>
