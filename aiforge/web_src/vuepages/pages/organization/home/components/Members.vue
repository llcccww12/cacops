<template>
  <div class="members-c">
    <div class="members-head">
      <div class="members-head-l"><i class="icon ri-user-line"></i>{{ $t('org.orgMember') }}</div>
      <a class="members-head-r" :href="`${pageData.OrgLink}/members`">
        {{ pageData.MembersTotal }} <i class="icon ri-arrow-right-s-line"></i>
      </a>
    </div>
    <div class="members-list">
      <a class="members-item" v-for="(item, index) in pageData.Members" :href="item.HomeLink"
        :title="item.Name + (item.FullName ? ' (' + item.FullName + ')' : '')">
        <img class="ui avatar" :src="item.RelAvatarLink"></a>
    </div>
  </div>
</template>

<script>

export default {
  name: "Members",
  props: {
    pageData: { type: Object, default: () => { } },
  },
};
</script>

<style scoped lang="less">
.members-c {
  .members-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 40px;

    .members-head-l {
      display: flex;
      align-items: center;
      color: rgba(16, 16, 16, 1);
      font-size: 16px;
      font-weight: 700;

      .icon {
        font-size: 24px;
      }
    }

    .members-head-r {
      display: flex;
      align-items: center;
      color: rgba(136, 136, 136, 1);
      font-size: 14px;

      &:hover {
        color: rgba(136, 136, 136, 0.8);
      }

      .icon {
        font-size: 20px;
      }
    }
  }

  .members-list {
    display: flex;
    flex-wrap: wrap;
    margin-bottom: 6px;

    .members-item {
      height: 45px;
      width: 45px;
      margin: 4px;

      img {
        height: 100%;
        width: 100%;
        border-radius: 100%;
      }
    }
  }
}
</style>
