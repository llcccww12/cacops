<template>
  <div class="list-container">
    <div class="list-item-container" v-loading="loading">
      <div class="item-container" v-for="(item, index) in list" :key="item.id">
        <Item :data="item" :condition="params" @changeFav="changeFav" :key="item.id"></Item>
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
import Item from './Item.vue';
import { getAiTaskTmplList } from '~/apis/modules/aitasktmpl';
import { getListValueWithKey } from '~/utils';
import { formatDate } from 'element-ui/lib/utils/date-util';
import { TmplTaskTypes, TmplComputerResouces, getGradientColor } from '~/pages/aitasktmpl/tools';

export default {
  name: "List",
  props: {
    params: { type: Object, default: () => ({}) },
  },
  components: { Item },
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
      getAiTaskTmplList({
        ...this.params,
        page: this.iPage,
        page_size: this.iPageSize,
      }).then(res => {
        res = res.data;
        this.loading = false;
        if (res.code == 0) {
          res = res.data;
          this.total = res.Total;
          this.list = (res.Templates || []).map(item => {
            const datasets = item.DatasetLists || [];
            const datasetsStr = datasets.map(item => {
              return `${item.OwnerName ? (item.OwnerName + '/') : ''}${item.DatasetAlias || item.DatasetName}`
            }).join(', ');
            const models = item.ModelLists || [];
            const modelsStr = models.map(item => {
              return `${item.OwnerName ? (item.OwnerName + '/') : ''}${item.ModelAlias || item.ModelName}`
            }).join(', ');
            return {
              ...item,
              NumCollections: Math.max(0, item.NumCollections),
              JobTypeStr: getListValueWithKey(TmplTaskTypes, item.JobType),
              ComputeSourceStr: getListValueWithKey(TmplComputerResouces, item.ComputeSource),
              UpdatedUnixStr: formatDate(new Date(item.UpdatedUnix * 1000), 'yyyy-MM-dd'),
              BgColor: getGradientColor(item.Name),
              DatasetsStr: datasetsStr,
              ModelsStr: modelsStr,
            }
          })
        } else {
          this.list = [];
          this.total = 0;
        }
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
      // this.getListData();
    },
    currentChange(page) {
      this.iPage = page;
      this.getListData();
    },
    sizeChange(pageSize) {
      this.iPageSize = pageSize;
      this.search();
    },
  },
  mounted() {
    this.search();
  },
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

@media only screen and (min-width: 1000px) and (max-width: 1700px) {
  /deep/ .list-item-container {
    .item-container {
      width: 50% !important;

      &:nth-child(2n+1) {
        padding-left: 0;
      }

      &:nth-child(2n) {
        padding-right: 0;
      }
    }
  }
}

@media only screen and (max-width: 1000px) {
  /deep/ .list-item-container {
    .item-container {
      width: 100% !important;
      padding-left: 0;
      padding-right: 0;
    }
  }
}
@media only screen and (max-width: 800px){
  .center {
    .el-pagination {
      /deep/ .el-pagination__total,
      /deep/ .el-pagination__sizes,
      /deep/ .el-pagination__jump {
        display: none;
      }
    }
  }
}
</style>
