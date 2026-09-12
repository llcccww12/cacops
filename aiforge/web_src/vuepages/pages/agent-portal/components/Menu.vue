<template>
  <div>
    <!-- 手机端菜单图标 -->
    <div class="mobile-menu-icon" v-if="!isMobileMenuOpen" @click="toggleMobileMenu">
      <i class="ri-arrow-right-s-line"></i>
    </div>
    <!-- 菜单遮罩 -->
    <div class="menu-mask" v-if="isMobileMenuOpen" @click="toggleMobileMenu"></div>
    <div class="main-menu" :class="{ 'mobile-menu-active': isMobileMenuOpen }">
      <template v-for="item in menuData">
        <MenuItem :data="item" :activePath="active" :key="item.path" @change="menuClick" />
      </template>
    </div>
  </div>
</template>

<script>
import MenuIcon from './MenuIcon.vue';
import MenuItem from './MenuItem.vue';

export default {
  name: 'Menu',
  data() {
    return {
      active: '',
      menuData: [],
      pathMap: {},
      isMobileMenuOpen: false, // 控制手机端菜单的展开状态
    };
  },
  components: { MenuIcon, MenuItem },
  watch: {
    $route: {
      handler() {
        this.refresh();
      }
    }
  },
  methods: {
    menuClick(item) {
      console.log("xxxxxxxxx",item.path,this.$route.path)
      if (item.path == this.$route.path) return;
      if (item.outLink) {
        console.log('outLink', item.outLink);
        
        // 确保是完整URL
        let outLink = item.outLink;
        if (!outLink.startsWith('http://') && !outLink.startsWith('https://')) {
          outLink = 'https://' + outLink; // 或根据实际情况处理
        }
        
        // 打开新窗口
        const newWindow = window.open(outLink, '_blank');
        
        // 如果弹窗被阻止，则直接跳转
        if (!newWindow) {
          window.location.href = outLink;
        }
        return;
      }
      if (item.path) {
        this.$router.push(item.path);
        console.log('active', item.path);
        this.active = item.path;
      }
      if (this.$isMobile) {
        this.isMobileMenuOpen = false;
      }
    },
    refresh() {
      const pathList = this.$route.path.slice(1).split('/');
      let path = '';
      for (let i = 0, iLen = pathList.length; i < iLen; i++) {
        path += ('/' + pathList[i]);
        const routeI = this.pathMap[path];
        if (routeI) {
          this.active = routeI.path;
          console.log('active', routeI.path);
        }
      }
    },
    toggleMobileMenu() {
      this.isMobileMenuOpen = !this.isMobileMenuOpen;
    },
    toggleSubMenu(item) {
      item.expanded = !item.expanded;
    },
    initMenu() {
      const routes = this.$router.options.routes || [];
      const allAuths = this.$router.auth;
      const pathMap = {};
      const walk = (item, parentPath, out) => {
        const fullPath = `${parentPath}/${item.path}`;
        let cur = null;
        if (item.meta?.menu) {
          cur = {
            index: fullPath,
            path: fullPath,
            name: item.name,
            label: this.$t(item.meta.label),
            icon: item.meta.icon,
            iconType: item.meta.iconType,
            topSplitLine: item.meta.topSplitLine,
            expanded: false,
            iconClass: item.meta.iconClass,
            outLink: item.outLink
          };
        }
        let authCheck = true;
        if (item.meta?.auths) {
          for (let auth in item.meta.auths) {
            if (!allAuths[auth]) {
              authCheck = false;
            }
          }
        }
        let hideCheck = false;
        if (item.meta?.hideAuths) {
          for (let auth in item.meta.hideAuths) {
            if (allAuths[auth]) {
              hideCheck = true;
            }
          }
        }
        const nextOut = [];
        if (cur && authCheck && !hideCheck && item.children && item.children.length) {
          item.children.forEach(_item => {
            walk(_item, cur.path, nextOut);
          });
          if (nextOut.length && cur) {
            cur.children = nextOut;
          }
        }
        if (cur && authCheck && !hideCheck) {
          out.push(cur);
          pathMap[cur.path] = cur;
        }
      };
      const menuData = [];
      (routes[0]?.children || []).forEach(item => {
        walk(item, '', menuData);
      });
      this.menuData = menuData;
      this.pathMap = pathMap;
      this.refresh();
    }
  },
  created() {
    this.initMenu();
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.main-menu {
  width: 180px;
  min-height: 400px;
  max-height: calc(100vh - 62px);
  box-sizing: content-box;
  padding: 20px 10px;

  .menu-item {
    height: 40px;
    display: flex;
    align-items: center;
    cursor: pointer;
    color: rgba(32, 37, 101, 0.9);

    .icon-c {
      height: 100%;
      width: 30px;
      display: flex;
      align-items: center;
      justify-content: center;
      color: rgb(111, 118, 165);
      font-size: 16px;
    }

    .label-c {
      display: flex;
      align-items: center;
      height: 100%;
    }

    &.active {
      color: rgb(0, 26, 199);
      background-color: rgba(230, 240, 255, 1);
      border-radius: 4px;

      .icon-c {
        color: rgb(0, 26, 199);
      }
    }

    &:hover:not(.active) {
      color: rgba(0, 26, 199, 0.8);

      .icon-c {
        color: rgba(0, 26, 199, 0.8);
      }
    }
  }

  .menu-line {
    border-bottom: 1px solid rgba(157, 197, 226, 0.4);
    height: 0;
    margin: 10px 2px;
  }

  .menu-group {
    .menu-group-title-c {
      display: flex;
      align-items: center;
      justify-content: space-between;
      cursor: pointer;
      padding-right: 8px;

      .menu-group-title-toggle-icon {
        display: flex;
        align-items: center;
        justify-content: center;
        transition: all 0.1s linear;
      }

      &.expanded {
        .menu-group-title-toggle-icon {
          transform: rotate(180deg);
        }

        &+.menu-group-list {
          height: 0;
        }
      }
    }

    .menu-group-title {
      color: rgb(111, 118, 165);
      height: 30px;
      display: flex;
      align-items: center;

      .icon-c {
        width: 30px;
      }
    }

    .menu-group-list {
      transition: all .2s ease-in-out;
      overflow: hidden;
    }
  }
}

.mobile-menu-icon {
  display: none;
}

/* 手机端样式 */
@media (max-width: 768px) {
  .main-menu {
    position: fixed;
    top: 62px;
    left: -200px;
    height: 100vh;
    z-index: 1000;
    background-color: #fff;
    box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
    transform: translateX(0);
    transition: transform 0.3s ease;

    &.mobile-menu-active {
      transform: translateX(200px);
    }
  }

  .mobile-menu-icon {
    position: absolute;
    top: 50%;
    /* 垂直居中 */
    left: -1px;
    /* 紧贴左边 */
    transform: translateY(-50%);
    /* 确保完全垂直居中 */
    z-index: 1001;
    cursor: pointer;
    display: flex;
    align-items: center;
    /* 图标垂直居中 */
    justify-content: center;
    /* 图标水平居中 */
    width: 20px;
    height: 40px;
    background-color: rgba(255, 255, 255, 1);
    box-shadow: 0pt 2pt 6pt 0pt rgba(73, 178, 255, 0.5);
    border: 1pt solid rgba(28, 126, 232, 1);
    color: #005cdd;
  }

  .menu-mask {
    position: fixed;
    top: 62px;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 999;
  }
}
</style>
