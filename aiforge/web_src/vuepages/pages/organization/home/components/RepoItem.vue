<template>
  <a class="repo-item" :href="`/${data.OwnerName}/${data.Name}`">
    <div class="title-c">
      <div class="repo-icon">
        <img v-if="data.Avatar" class="avatar" :src="`/repo-avatars/${data.Avatar}`" />
        <img v-else class="avatar" :avatar="data.Name" />
      </div>
      <div class="repo-name"><span v-html="data.NameShow" :title="`${data.Name}`"></span></div>
    </div>
    <div class="descr" v-show="data.DescriptionShow" v-html="data.DescriptionShow" :title="data.Description"></div>
  </a>
</template>

<script>
export default {
  name: "RepoItem",
  props: {
    data: { type: Object, default: () => ({}) },
  },
  data() {
    return {};
  },
  methods: {},
  mounted() { },
};
</script>

<style scoped lang="less">
.repo-item {
  border-color: rgba(157, 197, 226, 0.4);
  border-width: 1px;
  border-style: solid;
  border-radius: 6px;
  padding: 15px;
  background: transparent;
  height: 100%;
  position: relative;
  display: flex;
  flex-direction: column;
  min-height: 82px;

  .title-c {
    display: flex;
    align-items: center;

    .repo-icon {
      width: 28px;
      height: 28px;
      margin-right: 10px;

      img {
        width: 100%;
        height: 100%;
        border-radius: 100%;
      }
    }

    .repo-name {
      flex: 1;
      color: rgb(50, 145, 248);
      font-size: 14px;
      font-family: SourceHanSansSC;
      font-weight: 500;
      overflow: hidden;
      width: 0;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  }

  .descr {
    margin-top: 4px;
    color: rgb(16, 16, 16);
    font-size: 12px;
    font-family: SourceHanSansSC;
    font-weight: 300;
    overflow: hidden;
    text-overflow: ellipsis;
    word-break: break-all;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    max-height: 40px;
    white-space: break-spaces;
  }

  &:hover {
    border-color: rgba(204, 204, 255, 0.6);
    box-shadow: rgba(204, 204, 255, 0.6) 0px 4px 8px 0px;
    background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.34000000000000014%2C%20-0.643%2C%200.07087403200000002%2C%20-0.34000000000000014%2C%200.935%2C%200.997)%22%3E%3Cstop%20stop-color%3D%22%23ccfff4%22%20stop-opacity%3D%220.6%22%20offset%3D%220.01%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ece2ff%22%20stop-opacity%3D%220%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  }
}
</style>
