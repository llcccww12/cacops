<template>
  <div class="list-container">
    <div class="list-item-container" v-loading="loading">
      <div class="item-container" v-for="(item, index) in list" :key="item.id">
        <ModelItem :data="item" :condition="params" @changeFav="changeFav" :key="item.id"></ModelItem>
      </div>
      <div v-show="(!list.length && !loading)" class="no-data">
        <div class="item-empty">
          <div class="item-empty-icon"></div>
          <div class="item-empty-tips">{{ $t('modelObj.model_square_empty') }}</div>
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
import ModelItem from './ModelItem.vue';
import { getModelList } from '~/apis/modules/modelsquare';
import { MODEL_ENGINES } from '~/const';
import { getListValueWithKey } from '~/utils';
import { formatDate } from 'element-ui/lib/utils/date-util';

export default {
  name: "ModelList",
  props: {
    params: { type: Object, default: () => ({}) },
  },
  components: { ModelItem },
  data() {
    return {
      loading: false,
      list: [],
      iPageSizes: [30],
      iPageSize: 30,
      iPage: 1,
      total: 0,
    };
  },
  watch: {
    params: {
      handler(val, oval) {
        this.search()
      },
      deep: true,
    }
  },
  methods: {
    getListData() {
      this.loading = true;
      getModelList({
        q: this.params.q,
        queryType: this.params.tab,
        orderBy: this.params.sort,
        recommend: this.params.onlyRecommend,
        hasOnlineUrl: this.params.hasOnlineUrl,
        frame: this.params.engine,
        label: this.params.label,
        page: this.iPage,
        pageSize: this.iPageSize,
        notNeedEmpty: false,
      }).then(res => {
        res = res.data;
        this.loading = false;
        this.total = res.count || 0;
        this.list = (res.data || []).map(item => {
          return {
            ...item,
            labels: item.label ? item.label.trim().split(/\s+/) : [],
            engineName: getListValueWithKey(MODEL_ENGINES, item.engine.toString()),
            createTimeStr: formatDate(new Date(item.createdUnix * 1000), 'yyyy-MM-dd'),
            updateTimeStr: formatDate(new Date(item.updatedUnix * 1000), 'yyyy-MM-dd'),
          }
        });
      }).catch(err => {
        console.log(err);
        this.loading = false;
        this.list = [];
        this.total = 0;
      });
    },
    search() {
      this.iPage = 1;
      this.getListData();
    },
    changeFav() {
      if (this.params.tab == '3') {
        this.getListData();
      }
    },
    currentChange(page) {
      this.iPage = page;
      this.getListData();
    },
    sizeChange(pageSize) {
      this.iPageSize = pageSize;
      this.getListData();
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
      width: 33.3%;
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
  padding: 12px 0;
  width: 100%;

  .item-empty {
    height: 391px;
    width: 100%;
    overflow: hidden;
    padding: 15px;
    background: transparent;
    display: flex;
    flex-direction: column;
    justify-content: center;
    background-color: rgba(245, 245, 246, 0.5);

    .item-empty-icon {
      height: 80px;
      width: 100%;
      background: url(/img/empty-box.svg) center center no-repeat;
    }

    .item-empty-tips {
      text-align: center;
      margin-top: 20px;
      font-size: 18px;
      color: rgb(63, 63, 64);
    }
  }
}
</style>
