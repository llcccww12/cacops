<template>
  <div v-else class="no-data">
    <div class="item-empty">
      <div class="item-empty-icon"></div>
      <div class="item-empty-tips">{{ $t('org.no_result') }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Empty",
  data() {
    return {

    }
  },
}
</script>

<style scoped lang="less">
.no-data {
  display: flex;
  justify-content: center;
  padding: 12px 0;
  width: 100%;

  .item-empty {
    height: 391px;
    width: 100%;
    overflow: hidden;
    padding: 15px;
    background: transparent;
    display: flex;
    flex-direction: column;
    justify-content: center;
    // background-color: rgba(245, 245, 246, 0.5);

    .item-empty-icon {
      height: 80px;
      width: 100%;
      background: url(/img/empty-box.svg) center center no-repeat;
    }

    .item-empty-tips {
      text-align: center;
      margin-top: 20px;
      font-size: 18px;
      color: rgb(63, 63, 64);
    }
  }
}

</style>