<template>
  <div class="container ui" style="padding:0 75px;">
    <div class="menu">
      <div class="menu-l">
        <div class="menu-item" v-for="(item, index) in listL" :key="index"
          :class="focusMenu == item.key ? 'focused' : ''" @click="changeMenu(item)">
          {{ item.title }}
        </div>
      </div>
      <div class="menu-r">
        <div class="menu-item" v-for="(item, index) in listR" :key="index"
          :class="focusMenu == item.key ? 'focused' : ''" @click="changeMenu(item)">
          {{ item.title }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getIsTechAdmin } from '~/apis/modules/tech';

export default {
  name: "TopMenu",
  props: {
    menu: { type: String, default: '' },
  },
  components: {},
  data() {
    return {
      focusMenu: '',
      listL: [{
        key: 'tech_view',
        title: '按科技项目查看',
        url: '/tech/tech_view',

      }, {
        key: 'repo_view',
        title: '按项目成果查看',
        url: '/tech/repo_view',
      }],
      listR: [{
        key: 'my_view',
        title: '我申请的项目',
        url: '/tech/my_view',
      },],
      isTechAdmin: false,
    };
  },
  methods: {
    changeMenu(item, index) {
      this.focusMenu = item.key;
      window.location.href = item.url;
    }
  },
  beforeMount() {
    this.focusMenu = this.menu;
    getIsTechAdmin().then(res => {
      res = res.data;
      if (res.data && res.data.is_admin) {
        this.isTechAdmin = true;
        this.listR.push({
          key: 'admin_view',
          title: '管理展示项目',
          url: '/tech/admin_view',
        });
      }
    }).catch(err => {
      console.log(err);
    });
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.menu {
  display: flex;
  align-items: center;
  justify-content: space-between;

  .menu-l {
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  .menu-r {
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }

  .menu-item {
    display: flex;
    border-color: rgb(187, 187, 187);
    border-width: 1px 1px 0px 0px;
    border-style: solid;
    font-size: 14px;
    line-height: 20px;
    color: rgba(255, 255, 255, 0.7);
    align-items: center;
    text-align: center;
    justify-content: center;
    background: rgb(24, 52, 155);
    height: 40px;
    width: 178px;
    cursor: pointer;

    &:first-child {
      border-left-width: 1px;
    }

    &.focused {      
      border-color: rgb(22, 206, 255);
      color: rgb(255, 255, 255);
      background: rgb(22, 47, 255);
    }
  }
}
</style>
