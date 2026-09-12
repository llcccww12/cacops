<template>
  <div class="search-bar">
    <div class="search-bar-l">
      <div class="search-c">
        <div class="search-input-c">
          <input type="text" :placeholder="serchPlaceHolder" v-model="keyword" @keyup.enter="search">
        </div>
        <div class="search-btn" @click="search">{{ $t('repos.search') }}</div>
      </div>
      <el-button class="apply-btn" type="primary" icon="el-icon-plus" size="medium" @click="apply">申请展示成果</el-button>
      <div class="openi-link-c">
        <a class="openi-link" target="_blank"
          href="https://openi.pcl.ac.cn/OpenIOSSG/promote/raw/branch/master/other/OpenI%E5%90%AF%E6%99%BA%E7%A4%BE%E5%8C%BA%E9%A1%B9%E7%9B%AE%E5%BC%80%E6%BA%90%E6%8C%87%E5%8D%97.pdf">OpenI启智社区开源指南</a>
      </div>
    </div>
    <div class="sort-c">
      <el-select class="select" size="medium" v-model="sortType" @change="changeSort" placeholder="排序" clearable>
        <el-option v-for="item in sortList" :key="item.k" :label="item.v" :value="item.k" />
      </el-select>
    </div>
  </div>
</template>

<script>

export default {
  name: "SearchBar",
  props: {
    type: { type: Number, default: -1 }, // 0-tech_view, 1-repo_view
    condition: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {
      keyword: '',
      serchPlaceHolder: '',
      sortType: '',
      techSortList: [{
        k: 'time',
        v: this.$t('repos.recentlyUpdated'),
      }, {
        k: 'count',
        v: '项目成果数',
      }],
      repoSortList: [
        {
          k: 'mostpopular',
          v: this.$t('repos.mostPopular'),
        }, {
          k: 'recentupdate',
          v: this.$t('repos.recentlyUpdated'),
        }, {
          k: 'newest',
          v: this.$t('repos.newest'),
        }
      ],
      sortList: [],
      isTechAdmin: false,
    };
  },
  methods: {
    search() {
      this.$emit('changeCondition', {
        q: this.keyword.trim()
      });
    },
    apply() {
      window.location.href = '/tech/new';
    },
    manage() {
      window.location.href = '/tech/admin_view';
    },
    changeSort() {
      this.$emit('changeCondition', {
        sort: this.sortType
      });
    }
  },
  watch: {
    condition: {
      handler(newVal) {
        this.keyword = newVal.q;
        this.sortType = newVal.sort;
      },
      immediate: true,
      deep: true,
    },
  },
  beforeMount() {
    if (this.type == 0) {
      this.sortList = this.techSortList;
      this.serchPlaceHolder = '搜索科技项目名称';
    } else if (this.type == 1) {
      this.serchPlaceHolder = '搜索项目';
      this.sortList = this.repoSortList;
    }
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.search-bar {
  margin: 30px 0;
  display: flex;
  justify-content: space-between;

  .search-bar-l {
    display: flex;

    .search-c {
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 20px;

      .search-input-c {
        width: 268px;
        height: 40px;
        border-color: rgba(0, 61, 192, 0.73);
        border-width: 2px;
        border-style: solid;
        font-size: 14px;
        line-height: 20px;
        padding: 8px;
        display: flex;
        color: rgb(136, 136, 136);
        align-items: center;
        background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.014005111865831027%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%23e2d1ea%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%220.3%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");

        input {
          width: 100%;
          border: none;
          outline: none;
        }
      }

      .search-btn {
        height: 40px;
        font-size: 14px;
        line-height: 20px;
        padding: 0px;
        display: flex;
        color: rgb(255, 255, 255);
        align-items: center;
        text-align: center;
        justify-content: center;
        padding: 0 20px;
        cursor: pointer;
        background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-1.1800000000000002%2C%200.9499999999999998%2C%20-0.23749999999999996%2C%20-1.1800000000000002%2C%201.024%2C%200.047)%22%3E%3Cstop%20stop-color%3D%22%23bbd2f2%22%20stop-opacity%3D%221%22%20offset%3D%220.02%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23003dc0%22%20stop-opacity%3D%220.73%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
      }
    }

    .apply-btn {
      margin-left: 20px;
      border-color: rgb(31, 1, 115);
      border-width: 1px;
      border-style: solid;
      border-radius: 6px;
      font-size: 14px;
      background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.08281144868278038%2C%206.123233995736766e-17%2C%200.5%2C%200)%22%3E%3Cstop%20stop-color%3D%22%233291f8%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23060075%22%20stop-opacity%3D%220.73%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");

      &:hover {
        opacity: 0.9;
      }

      &:active {
        opacity: 0.8;
      }
    }

    .openi-link-c {
      margin-left: 20px;
      display: flex;
      align-items: center;
      justify-content: center;

      .openi-link {
        text-decoration: underline;
        color: rgb(50, 145, 248);
      }
    }
  }

  .sort-c {
    display: flex;
    align-items: center;
    justify-content: center;
    justify-self: flex-end;
  }
}
</style>
