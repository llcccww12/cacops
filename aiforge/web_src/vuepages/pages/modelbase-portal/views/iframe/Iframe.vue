<template>
  <div class="iframe-c">
    <iframe class="iframe" :src="iframeSrc" frameborder="0"></iframe>
  </div>
</template>

<script>
export default {
  name: 'Iframe',
  data() {
    return {};
  },
  components: {},
  methods: {},
  computed: {
    iframeSrc() {
      return this.$route.meta.iframeLink
    }
  },
  beforeMount() { },
  mounted() { },
};
</script>

<style scoped lang="less">
.iframe-c {
  position: relative;
  width: 100%;
  height: 100%;
  background-color: white;

  .iframe {
    border: none;
    width: 100%;
    height: 100%;
    margin-bottom: 0;
  }
}
</style>
