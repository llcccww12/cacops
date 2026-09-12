<template>
  <div class="condition-wrap">
    <div>
      <div class="tab-c">
        <div class="tab-item" v-for="(item) in tabList" :class="conds.tab == item.key ? 'focus' : ''" :key="item.key"
          @click="changeTab(item)">
          {{ item.label }}
        </div>
      </div>
    </div>
    <div class="condition-b">
      <div class="only-recommend-c"></div>
      <div>
        <el-dropdown class="sort-c" trigger="click" size="default">
          <span class="el-dropdown-link">
            {{ $t('datasets.sort') }}<i class="el-icon-caret-bottom el-icon--right"></i>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item :class="conds.sort == item.key ? 'active' : ''" v-for="item in sortList" :key="item.key"
              @click.native="changeSort(item)">
              {{ item.label }}
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>

<script>
import { i18n } from '~/langs';

const SortList = [{
  key: '',
  label: i18n.t('datasets.default'),
}, {
  key: 'moststars',
  label: i18n.t('datasets.mostCollections'),
}, {
  key: 'mostused',
  label: i18n.t('datasets.mostusecount'),
}, {
  key: 'newest',
  label: i18n.t('datasets.newest'),
}];

export default {
  name: "Condition",
  props: {
    condition: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {
      isLogin: false,
      tabList: [{
        key: '0',
        label: this.$t('imagesObj.image_public'),
      }, {
        key: '1',
        label: this.$t('imagesObj.image_my'),
      }, {
        key: '2',
        label: this.$t('imagesObj.image_collected'),
      }],
      sortList: [],
      conds: {
        tab: '0',
        sort: '',
      },
    };
  },
  methods: {
    changeTab(item) {
      this.conds.tab = item.key;
      if (item.key == '1') {
        this.conds.sort = 'newest';
        this.sortList = SortList.slice(1);
      } else {
        this.conds.sort = '';
        this.sortList = SortList.slice(0);
      }
      this.$emit('changeCondition', {
        tab: item.key,
        sort: this.conds.sort,
      });
    },
    changeSort(item) {
      this.conds.sort = item.key;
      this.$emit('changeCondition', {
        sort: this.conds.sort,
      });
    }
  },
  watch: {
    condition: {
      handler(newVal) {
        this.conds.tab = newVal.tab || '0';
        this.conds.sort = newVal.sort || '';
        this.sortList = SortList.slice(newVal.tab == '1' ? 1 : 0);
      },
      immediate: true,
      deep: true,
    },
  },
  beforeMount() {
    this.isLogin = !!document.querySelector('meta[name="_uid"]');
    if (!this.isLogin) {
      this.tabList.splice(1, Infinity);
    }
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.condition-wrap {
  margin: 10px 0 2px;

  .tab-c {
    display: flex;
    align-items: center;

    .tab-item {
      margin-right: 32px;
      padding: 4px 0;
      text-align: center;
      font-size: 14px;
      color: rgb(65, 80, 88);
      cursor: pointer;

      &.focus {
        font-weight: bold;
        color: #3291f8;
        font-weight: bold;
        border-bottom: 2px solid;
      }
    }
  }

  .condition-b {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 8px;

    .only-recommend-c {
      margin-right: 22px;
    }

    .sort-c {
      color: rgba(0, 0, 0, 0.87);
      cursor: pointer;
    }
  }
}

/deep/.el-dropdown-menu__item.active {
  color: #409EFF;
  background-color: rgba(179, 216, 255, 0.3);
}
</style>
