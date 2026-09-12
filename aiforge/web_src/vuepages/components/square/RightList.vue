<template>
  <div class="list-container">
    <div class="list-item-container">
      <div class="item-container" v-for="(item, index) in list" :key="item.id">
        <RightOwnerItem 
          v-if="isWorkspacePage" 
          :data="item" 
          :type="type" 
          :canChangeFav="canChangeFav"
          :reloadFlag="reloadFlag"
          :operaFlag="operaFlag"
          @reloadPage="reloadPage"
          @deleteEvent="deleteEvent"
          :key="item.id"
        />
        <RightItem 
          v-else
          :data="item" 
          :type="type" 
          :canChangeFav="canChangeFav" 
          :reloadFlag="reloadFlag" 
          @reloadPage="reloadPage"
          :key="item.id"/>
      </div>
      <div v-show="!list.length " class="no-data">
        <div class="item-empty">
          <div class="item-empty-icon"></div>
          <div class="item-empty-tips">{{ $t('modelObj.model_square_empty') }}</div>
        </div>
      </div>
    </div>
    <div class="center" v-show="list.length">
      <el-pagination ref="paginationRef" background @current-change="currentChange" @size-change="sizeChange"
        :current-page.sync="iPage" :pager-count="5" :page-sizes="iPageSizes" :page-size.sync="iPageSize"
        layout="total, sizes, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </div>
  </div>
</template>

<script>
import RightItem from './RightItem.vue';
import RightOwnerItem from './RightOwnerItem.vue';
export default {
  name: "ModelList",
  props: {
    list: { type: Array, default: () => []},
    pageParams: { 
      type: Object,
      default: () => ({}) 
    },
    canChangeFav: {  type: Boolean, default: true },
    reloadFlag: {  type: Boolean, default: false },
    operaFlag: {  type: Boolean, default: false },
    type: {  type: String, default: 'dataset' },
    isWorkspacePage: { type: Boolean, default: false }
  },
  components: { RightItem, RightOwnerItem },
  data() {
    return {
      iPageSizes: [30],
      iPageSize: 30,
      iPage: 1,
      total: 0,
    };
  },
  watch: {
    // canChangeFav(val){

    // },
    pageParams(val){
      this.iPageSizes = [val.page_size];
      this.iPageSize = val.page_size;
      this.total = val.total;
      this.iPage = val.page;
    }

  },
  methods: {
    reloadPage() {
      const isLastPage = this.iPage === Math.ceil(this.total / this.iPageSize);
      const isSingleItemOnPage = this.list.length === 1;
      let newPage = this.iPage;
      if (isLastPage && isSingleItemOnPage && this.iPage > 1) {
        newPage = this.iPage - 1;
      }
      this.$emit("changePage", newPage);
      // Update the current page if it changed
      if (newPage !== this.iPage) {
        this.iPage = newPage;
      }
    },
    deleteEvent(data){
      console.log("deleteEvent",data)
      this.$emit("deleteEvent", data)
    },
    currentChange(page) {
      this.iPage = page;
      this.$emit("changePage",page)
    },
    sizeChange(pageSize) {
      this.iPageSize = pageSize;
    },
  },
  mounted() { 
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
@media only screen and (min-width: 850px) and (max-width: 1440px) {
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

@media only screen and (min-width: 1440px) {
  /deep/ .list-item-container {
    .item-container {
      width: 33.3% !important;

      &:nth-child(3n+1) {
        padding-left: 0;
      }

      &:nth-child(3n) {
        padding-right: 0;
      }
    }
  }
}
</style>
