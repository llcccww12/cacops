<template>
  <div class="list-container">
    <div class="list-item-container" v-loading="loading">
      <div class="item-container" v-for="(item) in list" :key="item.id">
        <ImageItem :data="item" :condition="params" @changeImage="changeImage" @refreshImage="refreshImage"
          :key="item.id">
        </ImageItem>
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
import ImageItem from './Item.vue';
import { getImages } from '~/apis/modules/images';
import { formatDate } from 'element-ui/lib/utils/date-util';
import { COMPUTER_RESOURCES_COLORS } from '~/const';

export default {
  name: "List",
  props: {
    params: { type: Object, default: () => ({}) },
  },
  components: { ImageItem },
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
      const params = {
        q: this.params.q,
        type: this.params.tab,
        sort: this.params.sort,
        trainType: this.params.trainType,
        computeResource: this.params.compute_resource,
        framework: this.params.framework,
        frameworkVersion: this.params.framework_version,
        python: this.params.python,
        cuda: this.params.cuda,
        cann: this.params.cann,
        dtk: this.params.dtk,
        page: this.iPage,
        pageSize: this.iPageSize,
      };
      console.log(params)
      getImages(params).then(res => {
        res = res.data;
        this.loading = false;
        this.total = res.count || 0;
        this.list = (res.images || []).map(item => {
          const thirdPackagesList = [];
          const thirdPackages = item.thirdPackages.split('\n');
          thirdPackages.forEach(pkgLine => {
            if (pkgLine) {
              thirdPackagesList.push(pkgLine.trim().replace('==', ' '));
            }
          });
          const trainTypeList = [];
          const trainTypes = item.trainType.split('&');
          trainTypes.forEach(type => {
            if (type) {
              if (type === 'Notebook') {
                if (['GCU', 'GPU'].includes(item.compute_resource)) {
                  trainTypeList.push(this.$t('TaskTypeTitle.Notebook'));
                } else {
                  trainTypeList.push(this.$t('TaskTypeTitle.Notebook1'));
                }
              } else {
                trainTypeList.push(this.$t('TaskTypeTitle.' + type));
              }
            }
          });
          const compute_resource = item.compute_resource || 'GPU';
          return {
            ...item,
            computeResourceColor: COMPUTER_RESOURCES_COLORS[compute_resource],
            computeResourceShow: this.$t('computeResourceTitle.' + compute_resource),
            thirdPackagesShow: thirdPackagesList.join('; '),
            trainTypeShow: trainTypeList.join('、'),
            createTimeStr: formatDate(new Date(item.createdUnix * 1000), 'yyyy-MM-dd HH:mm:ss'),
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
    changeImage() {
      this.getListData();
    },
    refreshImage(data) {
      const find = this.list.find(item => item.id == data.id);
      find && Object.assign(find, data);
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

@media only screen and (max-width: 800px) {
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
