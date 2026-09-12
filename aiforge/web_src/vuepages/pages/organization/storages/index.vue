<template>
  <div class="content">
    <StorageList type="org" :orgName="orgName"></StorageList>
  </div>
</template>
<script>
import StorageList from '~/pages/storage/index.vue'

export default {
  data() {
    return {
      orgName: '',
    };
  },
  components: { StorageList },
  methods: {
  },
  beforeMount() {
    const { pathname } = window.location;
    const orgMatch = pathname.match(/^\/org\/([^\/]+)(\/.*)?$/);
    this.orgName = orgMatch[1]
  },
  mounted() {
   
  },
  beforeDestroy() { },
};
</script>
<style scoped lang="less">

</style>