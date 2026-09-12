<template>
  <div>
    <div class="container">
      <div class="title">
        <i style="margin-left:10px;margin-right:8px;font-size:20px;" class="ri-blaze-line"></i>
        <span>{{ $t('repos.activeOrganization') }}</span>
      </div>
      <div class="content">
        <div class="item" v-for="(item, index) in list" :key="index">
          <div class="item-l">
            <a class="name" :href="`/${item.Name}`" :title="item.Name">
              <img class="avatar" :src="item.RelAvatarLink">
              <div class="name-c"><span>{{ item.Name }}</span></div>
            </a>
          </div>
          <div class="item-r">
            <i class="ri-user-2-line" style="color:rgb(250, 140, 22);margin-right:4px;"></i>
            <span>{{ item.NumMembers }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getActiveOrgs } from '~/apis/modules/repos';

export default {
  name: "ActiveOrgs",
  props: {},
  components: {},
  data() {
    return {
      list: [],
    };
  },
  methods: {},
  mounted() {
    getActiveOrgs().then(res => {
      res = res.data;
      if (res.Code == 0) {
        this.list = res.Data.Orgs || [];
      } else {
        this.list = [];
      }
    }).catch(err => {
      console.log(err);
      this.list = [];
    });
  },
};
</script>

<style scoped lang="less">
.title {
  height: 43px;
  border-color: rgba(16, 16, 16, 0.05);
  border-width: 1px 0px;
  border-style: solid;
  display: flex;
  align-items: center;
  font-size: 18px;
  color: rgba(47, 9, 69, 0.74);
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-1.007%2C%200.0010000000000001494%2C%20-0.000018400023883213844%2C%20-1.007%2C%201.003%2C%200.008)%22%3E%3Cstop%20stop-color%3D%22%23eee9da%22%20stop-opacity%3D%220%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23f3e7f7%22%20stop-opacity%3D%220.26%22%20offset%3D%220.29%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23d0e7ff%22%20stop-opacity%3D%220.3%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
}

.content {
  padding: 6px 0;
}

.item {
  display: flex;
  align-items: center;
  height: 48px;
  padding: 0 8px;
  font-size: 14px;
  color: rgba(16, 16, 16, 1);
}

.item>div {
  display: flex;
  align-items: center;
}

.item-l {
  flex: 1;
  overflow: hidden;
  a {    
    display: flex;
    align-items: center;
    overflow: hidden;
  }
}

.item-r {
  width: 80px;
  justify-content: flex-end;
}

.item .avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  margin-right: 10px;
}

.item .name-c {
  flex: 1;
  overflow: hidden;
  width: 100%;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.item .name {
  color: rgba(16, 16, 16, 1);

  &:hover {
    opacity: 0.8;
  }
}
</style>
