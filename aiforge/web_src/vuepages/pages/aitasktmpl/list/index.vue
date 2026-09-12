<template>
  <div class="content">
    <left-labels
      :title="'taskTmplObj.taskTmpl'"
      :filters="filtersConfig"
      @trigger-search="changeCondition"
    />
    <div class="content-r">
      <Recommend :type="'ai_task_template'"></Recommend>
      <Condition ref="rightCards" @changeCondition="changeCondition" @changeModalOverlay="changeModalOverlay" :typeList="typeList" />
      <List :params="params" />
      <!-- 移动端left-lables遮罩层（独立层级，不覆盖按钮） -->
      <div class="modal-overlay" v-if="showFilter">
        <div class="filter-panel">
          <!-- 弹窗内容（与图片描述一致） -->
          <div class="filter-section">
            <div class="filter-content">
              <div  class="block-c" v-for="(filter, index) in filtersConfig" :key="index">
                <div class="title-c">
                  <span class="title">{{ $t(filter.titleKey) }}</span>
                </div>
                <div class="list">
                  <div class="item" :class="{ active: item.active, [filter.styleClass]: true }"
                    v-for="(item, idx) in filter.items" :key="idx" @click="selectItem(filter.type, item)">
                    {{ item.v }}
                  </div>
                </div>
              </div>
              <!-- 其他分类标签 -->
            </div>
            <div class="filter-buttons">
              <div class="reset-button btn" @click="resetFilters">重置</div>
              
            </div>
          </div>
          
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import LeftLabels from '~/components/square/LeftLabels.vue'
import Condition from './components/Condition.vue'
import List from './components/List.vue'
import Recommend from '~/components/square/Recommend.vue';
import { TmplTaskTypes, TmplComputerResouces } from '~/pages/aitasktmpl/tools';

export default {
  data() {
    return {
      type: 'public',
      typeList: [{
        key: 'public',
        label: `⭕${this.$t('taskTmplObj.publicTmpl')}`,
      }, {
        key: 'recommend',
        label: `🏆${this.$t('taskTmplObj.recommendTmpl')}`,
      }],
      params: {
        type: '',
        q: '',
        tags: '',
        job_type: '',
        compute_source: '',
        dataset_id: '',
        model_id: '',
        repo_id: '',
        order_by: '',
        page: 1,
        page_size: 30,
      },
      pageParams: {
        page: 1,
        page_size: 30,
        total: 0
      },
      showFilter: false,
      filtersConfig: []
    };
  },
  components: { Condition, List, Recommend, LeftLabels },
  methods: {
    changeCondition(conds) {
      this.params = {
        ...this.params,
        ...conds,
      }
    },
    changeModalOverlay(val){
      this.showFilter = val;
    },
    // 移动端filter
    selectItem(type, clickedItem){
       // 1. 找到对应的筛选组
      const filterGroup = this.filtersConfig.find(f => f.type === type);
      if (!filterGroup) return;

      this.params[type] = clickedItem.k;

      filterGroup.items.forEach(item => {
        item.active = (item === clickedItem); // 只有当前点击项为 true
      });
      this.showFilter = false
      this.$refs.rightCards.resetFilter()
    },
    resetFilters() {
      this.filtersConfig.forEach(group => {
        this.params[group.type] = ''
        group.items.forEach(item => {
          item.active = false;
        });
      });
      this.showFilter = false
      this.$refs.rightCards.resetFilter()
    },
   
  },
  beforeMount() {},
  mounted() {
    TmplTaskTypes.forEach(item => {
      item.active = false;
    });
    TmplComputerResouces.forEach(item => {
      item.active = false;
    });
    this.filtersConfig = [
      {
          type: 'job_type',
          titleKey: 'resourcesManagement.jobType',
          items: TmplTaskTypes,
          clearFlag: false,
          styleClass: 'first-style'
      },
      {
          type: 'compute_source',
          titleKey: 'resourcesManagement.computeResource',
          items: TmplComputerResouces,
          clearFlag: false,
          styleClass: 'second-style'
      }
    ]
  },
  beforeDestroy() { },
};
</script>

<style scoped lang="less">
.content {
  display: flex;

  .content-r {
    flex: 1;
    width: 0;
    padding: 30px 20px;
    padding-right: 36px;
  }
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.2);
    z-index: 100; /* 低于按钮的z-index */
    display: flex;
    justify-content: center;
    align-items: flex-start;
    /* 弹窗样式 */
    .filter-panel {
      width: 100%;
      margin: 0 14px;
      margin-top: 120px;
      box-shadow: 0px 4px 4px 0px rgba(16,16,16,0.1);
      border-radius: 5px;
      background-color: rgba(250,250,255,1);
      border: 1px solid rgba(16,16,16,0.2);
      height: 480px;
      z-index: 101; /* 高于遮罩层 */
      .filter-section{
        display: flex;
        flex-direction: column;
        height: 100%;
        .filter-content{
          flex: 1;
          overflow: auto;
          margin: 24px 8px 0 16px;
          .block-c{
            margin-top: 8px;
            .title-c {
              display: flex;
              align-items: flex-end;
              margin-bottom: 8px;

              .title {
                color: rgb(16, 16, 16);
                font-size: 14px;
              }
            }
            .list {
              display: flex;
              flex-wrap: wrap;
              .item {
                border-radius: 4px;
                font-size: 12px;
                color: rgba(14, 37, 69, 1);
                box-shadow: 0px 1px 1px 0px rgba(16, 16, 16, 0.2);
                margin: 0 10px 10px 0;
                padding: 2px 4px;
                display: flex;
                align-items: center;
                justify-content: center;
                cursor: pointer;

                /* 默认样式 */
                border: 1px solid rgba(145, 213, 255, 0.5);
                &.active {
                  background-color: #0066ff;
                  color: #fff;
                }

                /* 不同样式类型 */
                // &.first-style {
                //     border-color: rgba(0, 158, 255, 1);
                // }

                &.second-style {
                  border: 1px solid rgba(255, 198, 145, 0.5);
                  &.active {
                    background-color: #0066ff;
                    color: #fff;
                  }
                }
                &.third-style {
                  border: 1px solid rgba(169, 223, 184, 0.5);
                  &.active {
                    background-color: #0066ff;
                    color: #fff;
                  }
                }
              }
            }
          }
        }
        .filter-buttons{
          height: 60px;
          display: flex;
          align-items: center;
          justify-content: center;
          .btn{
            cursor: pointer;
            display: flex;
            align-items: center;
            justify-content: center;
            height: 40px;
            border-radius: 4px;
            font-size: 14px;
          }
          .reset-button{
            width: 90px;
            background-color: rgba(255,255,255,1);
            color: rgba(16,16,16,0.75);
            border: 1px solid rgba(16,16,16,0.5);
            margin-right: 20px;
            
          }
          .search-button{
            width: 171px;
            background-color: rgba(22,132,252,1);
            color: rgba(255,255,255,1);
          }
        }
      }
    }
  }
}
/* 保持原有样式结构 */
@media only screen and (max-width: 849.99px) {	

  .content-r {
    padding: 20px 16px !important;
  }
}
</style>
