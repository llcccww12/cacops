<template>
  <div class="condition-wrap">
    <div>
      <div class="tab-c">
        <div class="tab-item" v-for="(item, index) in tabList" :class="conds.tab == item.key ? 'focus' : ''"
          :key="item.key" @click="changeTab(item)">
          {{ item.label }}
        </div>
      </div>
    </div>
    <div class="condition-b">
      <div class="only-recommend-c">
        <el-checkbox v-model="conds.onlyRecommend" @change="changeRecommend">
          {{ $t('datasets.platform_recommendations') }}
        </el-checkbox>
        <el-checkbox v-model="conds.hasOnlineUrl" @change="changeOnline">
          {{ $t('modelObj.can_online_infer') }}
        </el-checkbox>
      </div>
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

export default {
  name: "ModelCondition",
  props: {
    condition: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {
      isLogin: false,
      tabList: [{
        key: '1',
        label: this.$t('modelObj.model_public'),
      }, {
        key: '2',
        label: this.$t('modelObj.model_my'),
      }, {
        key: '3',
        label: this.$t('modelObj.model_collected'),
      }, {
        key: '5',
        label: this.$t('modelObj.model_my_migrate'),
      }],
      sortList: [{
        key: '',
        label: this.$t('datasets.default'),
      }, {
        key: 'created_unix',
        label: this.$t('datasets.newest'),
      }, {
        key: 'updated_unix',
        label: this.$t('datasets.recentupdate'),
      }, {
        key: 'download_count',
        label: this.$t('datasets.downloadtimes'),
      }, {
        key: 'collected_count',
        label: this.$t('datasets.moststars'),
      }, {
        key: 'reference_count',
        label: this.$t('datasets.mostusecount'),
      }, {
        key: 'derivative_count',
        label: this.$t('modelManage.mostDerivative'),
      }],
      conds: {
        tab: 'public',
        sort: '',
        onlyRecommend: false,
      },
    };
  },
  methods: {
    changeTab(item) {
      this.conds.tab = item.key;
      this.$emit('changeCondition', {
        tab: item.key,
      });
    },
    changeRecommend(item, _item) {
      this.$emit('changeCondition', {
        onlyRecommend: this.conds.onlyRecommend,
      });
    },
    changeOnline(item) {
      this.$emit('changeCondition', {
        hasOnlineUrl: this.conds.hasOnlineUrl,
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
        this.conds.tab = newVal.tab || '1';
        this.conds.sort = newVal.sort || '';
        this.conds.onlyRecommend = newVal.onlyRecommend || false;
        this.conds.hasOnlineUrl = newVal.hasOnlineUrl || false;
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
    margin-top: 20px;
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
