<template>
  <div class="list-container">
    <div style="min-height:540px;">
      <div class="list-item-container" v-loading="loading">
        <div class="item-container" v-for="(item, index) in list" :key="item.ID">
          <PrjResultsItem :data="item"></PrjResultsItem>
        </div>
      </div>
      <div v-show="(!list.length && !loading)" class="no-data">
        <div class="item-empty">
          <div class="item-empty-icon bgtask-header-pic"></div>
          <div class="item-empty-tips">没有找到相关的项目</div>
        </div>
      </div>
    </div>
    <div class="center" v-show="list.length">
      <el-pagination ref="paginationRef" background @current-change="currentChange" @size-change="sizeChange"
        :current-page.sync="iPage" :page-sizes="iPageSizes" :page-size.sync="iPageSize"
        layout="total, sizes, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </div>
  </div>
</template>

<script>
import PrjResultsItem from './PrjResultsItem.vue';
import LetterAvatar from '~/utils/letteravatar';
import { getTechOpenISearch } from '~/apis/modules/tech';

export default {
  name: "PrjResultsList",
  props: {
    condition: { type: Object, default: () => ({}) },
  },
  components: { PrjResultsItem },
  data() {
    return {
      loading: false,
      list: [],
      iPageSizes: [15, 30, 50],
      iPageSize: 15,
      iPage: 1,
      total: 0,
    };
  },
  methods: {
    getListData() {
      this.loading = true;
      getTechOpenISearch({
        name: this.condition.q,
        tech_name: this.condition.project_name,
        institution_name: this.condition.institution_name,
        topic: this.condition.topic,
        page: this.condition.page,
        pageSize: this.condition.pageSize,
        sort: this.condition.sort,
      }).then(res => {
        res = res.data;
        this.loading = false;
        this.total = res.total || 0;
        this.list = res.data || [];
        this.$nextTick(() => {
          LetterAvatar.transform();
        });
      }).catch(err => {
        console.log(err);
        this.loading = false;
        this.list = [];
        this.total = 0;
      });
    },
    search() {
      this.getListData();
    },
    currentChange(page) {
      this.iPage = page;
      this.$emit('changeCondition', {
        page: this.iPage,
        pageSize: this.iPageSize,
        changePage: true,
      });
    },
    sizeChange(pageSize) {
      this.iPageSize = pageSize;
      this.$emit('changeCondition', {
        page: this.iPage,
        pageSize: this.iPageSize,
      });
    },
  },
  watch: {
    condition: {
      handler(newVal) {
        this.iPage = newVal.page;
        this.iPageSize = newVal.pageSize;
      },
      immediate: true,
      deep: true,
    },
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.list-container {
  .list-item-container {
    display: flex;
    flex-wrap: wrap;

    .item-container {
      width: 50%;
      padding: 12px;
    }
  }
}

.center {
  text-align: center;
  margin-top: 10px;
}

.no-data {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 12px;

  .item-empty {
    height: 180px;
    width: 100%;
    padding: 12px;
    border-color: rgb(232, 224, 236);
    border-width: 1px;
    border-style: solid;
    box-shadow: rgba(168, 157, 226, 0.2) 0px 5px 10px 0px;
    background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(-0.01900000000000005%2C%200.997%2C%20-0.06169646324801269%2C%20-0.01900000000000005%2C%200.995%2C%200.014)%22%3E%3Cstop%20stop-color%3D%22%23f2edf5%22%20stop-opacity%3D%221%22%20offset%3D%220.01%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ffffff%22%20stop-opacity%3D%221%22%20offset%3D%220.31%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
    display: flex;
    flex-direction: column;
    justify-content: center;

    .item-empty-icon {
      height: 80px;
      width: 100%;
    }

    .item-empty-tips {
      font-size: 16px;
      color: rgb(16, 16, 16);
      text-align: center;
      margin-top: 2px;
    }
  }
}
</style>
