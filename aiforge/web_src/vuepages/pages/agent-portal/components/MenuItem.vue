<template>
  <div>
    <div v-if="data.children && data.children.length" class="menu-group" :class="deepLevel ? 'deep-level' : ''">
      <div class="menu-line" v-if="!deepLevel"></div>
      <div class="menu-group-title-c" :class="data.expanded ? 'expanded' : ''" v-if="data.label"
        @click="toggleSubMenu(data)">
        <div class="menu-group-title">
          <div class="icon-c">
            <MenuIcon v-if="data.icon" :icon="data.icon" :type="data.iconType || 'iconfont'"
              :active="activePath == data.path">
            </MenuIcon>
          </div>
          <div class="label-c">{{ data.label }}</div>
        </div>
        <div class="menu-group-title-toggle-icon">
          <svg xmlns="http://www.w3.org/2000/svg" fill="rgba(16, 16, 16, 0.5)" viewBox="0 0 32 32" width="16"
            height="16">
            <defs />
            <g>
              <path d="M16 17.563l6.6-6.6 1.885 1.885-8.485 8.485-8.485-8.485 1.885-1.885z" />
            </g>
          </svg>
        </div>
      </div>
      <div class="menu-group-list">
        <template v-for="_item in data.children">
          <MenuItem :data="_item" :activePath="activePath" :key="_item.path" @change="menuClick" :deepLevel="true" />
        </template>
      </div>
    </div>
    <template v-else>
      <div v-if="data.topSplitLine" class="menu-line"></div>
      <div class="menu-item" :class="activePath == data.path ? 'active' : ''" @click.stop.prevent="menuClick(data)">
        <div class="icon-c" :class="data.iconClass">
          <MenuIcon v-if="data.icon" :icon="data.icon" :type="data.iconType || 'iconfont'"
            :active="activePath == data.path">
          </MenuIcon>
        </div>
        <div class="label-c">{{ data.label }}</div>
      </div>
    </template>
  </div>
</template>

<script>
import MenuIcon from './MenuIcon.vue';

export default {
  name: 'MenuItem',
  props: {
    activePath: { type: String, default: '' },
    data: { type: Object, default: () => ({}) },
    deepLevel: { type: Boolean, default: false },
  },
  data() {
    return {};
  },
  components: { MenuIcon },
  methods: {
    menuClick(item) {
      this.$emit('change', item);
    },
    toggleSubMenu(item) {
      item.expanded = !item.expanded;
    },
  },
  created() { },
  mounted() { console.log(this.data.path)},
};
</script>

<style scoped lang="less">
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
  .m-left{
    margin-left: 10px;
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
      height: 100%;
      width: 30px;
      display: flex;
      align-items: center;
      justify-content: center;
      color: #6f76a5;
      font-size: 16px;
    }
  }

  .menu-group-list {
    transition: all .2s ease-in-out;
    overflow: hidden;
  }

  &.deep-level {
    .menu-group-list {
      padding-left: 15px;
    }
  }
}
</style>
