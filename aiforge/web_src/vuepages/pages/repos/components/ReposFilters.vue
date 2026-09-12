<template>
  <div>
    <div class="item" :class="(focusIndex == index) ? 'item-focus' : ''" v-for="(item, index) in list" :key="item.key">
      <a href="javascript:;" @click="changeFilters(item, index)">{{ item.label }}</a>
    </div>
  </div>
</template>

<script>
export default {
  name: "ReposFilters",
  props: {
    defaultsort: { type: String, default: 'mostpopular' },
  },
  components: {},
  data() {
    return {
      focusIndex: 0,
      list: [{
        key: 'mostpopular',
        label: this.$t('repos.mostPopular'),
      }, {
        key: 'mostactive',
        label: this.$t('repos.mostActive'),
      }, {
        key: 'recentupdate',
        label: this.$t('repos.recentlyUpdated'),
      }, {
        key: 'newest',
        label: this.$t('repos.newest'),
      }, {
        key: 'moststars',
        label: this.$t('repos.mostStars'),
      }, {
        key: 'mostforks',
        label: this.$t('repos.mostForks'),
      }, {
        key: 'mostdatasets',
        label: this.$t('repos.mostDatasets'),
      }, {
        key: 'mostaitasks',
        label: this.$t('repos.mostAiTasks'),
      }, {
        key: 'mostmodels',
        label: this.$t('repos.mostModels'),
      }]
    };
  },
  methods: {
    changeFilters(item, index) {
      this.focusIndex = index;
      this.$emit('change', this.list[this.focusIndex]);
    },
    setDefaultFilter(sort) {
      const index = this.list.findIndex((item) => item.key == sort);
      this.focusIndex = index >= 0 ? index : 0;
    }
  },
  mounted() {
  },
};
</script>

<style scoped lang="less">
.item {
  height: 40px;
  border-color: rgba(16, 16, 16, 0.05);
  border-width: 0px 0px 1px;
  border-style: solid;
  color: rgba(16, 16, 16, 0.8);
  font-size: 14px;
  padding: 0px;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  position: relative;
}

.item a {
  color: inherit;
  height: 100%;
  display: flex;
  align-items: center;

  &:hover {
    opacity: 0.8;
  }
}

.item-focus {
  font-weight: bold;
  border-color: rgba(0, 108, 205, 0.3);
  color: rgb(50, 145, 248);

  a {
    cursor: default;

    &:hover {
      opacity: 1;
    }
  }

}

.item-focus:before {
  content: "";
  position: absolute;
  width: 7px;
  height: 7px;
  bottom: -4px;
  background: rgb(178, 211, 240);
  left: 0;
  border-radius: 100%;
}
</style>
