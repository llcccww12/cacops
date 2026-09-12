<template>
  <div class="menu-icon-c" :class="active ? 'active' : ''">
    <div class="svg-c" v-if="type == 'svg'">
      <svg v-if="icon == 'arena'" width="16" height="16" viewBox="0 0 48 48" fill="none"
        xmlns="http://www.w3.org/2000/svg">
        <path d="M44 26L24 36L4 26L24 16L44 26Z" stroke-width="4" stroke-miterlimit="2" stroke-linejoin="round" />
        <path d="M24 7L44 17L24 27L4 17L24 7Z" stroke-width="4" stroke-miterlimit="2" stroke-linejoin="round" />
        <path d="M44 26V34L24 44L4 34V26L24 36L44 26Z" fill="none" stroke-width="4" stroke-miterlimit="2"
          stroke-linecap="round" stroke-linejoin="round" />
        <path d="M44 14V17V26" stroke-width="4" stroke-miterlimit="2" stroke-linecap="round" stroke-linejoin="round" />
        <path d="M4 26V17V14" stroke-width="4" stroke-miterlimit="2" stroke-linecap="round" stroke-linejoin="round" />
        <path d="M24 36V24" stroke-width="4" stroke-miterlimit="2" stroke-linecap="round" stroke-linejoin="round" />
        <path d="M24 16V4" stroke-width="4" stroke-miterlimit="2" stroke-linecap="round" stroke-linejoin="round" />
      </svg>
      <svg v-else-if="icon == 'evaluate'" width="16" height="16" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
        <rect x="6" y="6" width="36" height="36" rx="3" fill="none" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
        <path d="M12 25H15L19 14L22 36L27 23L31 29L34 25H37" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
      </svg>
      <svg v-else-if="icon == 'evallist'"  width="16" height="16" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M20 12L32 24L20 36V12Z" fill="none" stroke="#001ac7" stroke-width="4" stroke-linejoin="round"/></svg>
    </div>
    <i v-else :class="icon"></i>
  </div>
</template>

<script>
export default {
  name: 'MenuIcon',
  props: {
    type: { type: String, default: 'iconfont' }, // iconfont, svg
    active: { type: Boolean, default: false, },
    icon: { type: String, default: '' },
  },
  data() {
    return {};
  },
  components: {},
  methods: {},
  mounted() { },
};
</script>

<style scoped lang="less">
.menu-icon-c {
  display: flex;
  align-items: center;

  .svg-c {
    display: flex;
    align-items: center;
  }

  svg {
    stroke: #6f76a5;
    fill: none;
  }

  i {
    color: #6f76a5;
  }

  &.active {
    svg {
      path {
        stroke: #001ac7;
      }
      rect {
        stroke: #001ac7;
      }
    }

    i {
      color: #001ac7;
    }
  }
}
</style>
