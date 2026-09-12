<template>
<div class="content-l" v-if="!isWorkspacePage">
  <div class="container">
    <div class="title">{{ $t(title) }}</div>
    <div 
      class="block-c"
      v-for="(filter, index) in localFilters"
      :key="index"
    >
      <div class="title-c">
        <span class="title">{{ $t(filter.titleKey) }}</span>
        <span 
          class="clear-btn" 
          v-if="filter.clearFlag" 
          @click="clearFilter(filter.type)"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 32 32" width="12" height="12"><defs></defs><g><path d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"></path></g></svg>
        </span>
      </div>
      <div class="list">
        <div
          class="item"
          :class="{ 
            active: item.active,
            [filter.styleClass]: true
          }"
          v-for="(item, idx) in filter.items"
          :key="idx"
          @click="selectItem(filter.type, item)"
        >
          {{ item.v }}
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: "CommonFilter",
  props: {
    isWorkspacePage: { type: Boolean, default: false },
    title: { type: String, default: '' },         // 主标题
    filters: {                                    // 过滤项配置
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      // 筛选条件值集合
      selectedFilters: {},
      localFilters: []
    }
  },
  watch: {
    filters: {
      handler(newVal) {
        this.localFilters = JSON.parse(JSON.stringify(newVal))
      },
      deep: true
    }
  },
  methods: {
    selectItem(type, item) {
      this.localFilters = this.localFilters.map(filter => {
        if (filter.type === type) {
          // 更新选中状态（单选模式）
          return {
            ...filter,
            items: filter.items.map(i => ({
              ...i,
              active: i.k === item.k
            })),
            clearFlag: true
          }
        }
        return filter
      })
      this.updateSelectedFilters()
      this.triggerSearch()
    },
    clearFilter(type) {
        this.localFilters = this.localFilters.map(filter => {
        if (filter.type === type) {
          return {
            ...filter,
            items: filter.items.map(i => ({ ...i, active: false })),
            clearFlag: false
          }
        }
        return filter
       
      })
      this.updateSelectedFilters()
      this.triggerSearch()
    
    },
    // 更新选中值集合
    updateSelectedFilters() {
      const selected = {}
      this.localFilters.forEach(filter => {
        const activeItem = filter.items.find(item => item.active)
        selected[filter.type] = activeItem ? activeItem.k : ''
      })
      this.selectedFilters = selected
      this.selectedFilters.page = 1
    },
    // 触发搜索
    triggerSearch() {
      this.$emit('trigger-search', this.selectedFilters)
    },
  },
  mounted() {
    console.log('taskType', this.isWorkspacePage)
  },
};
</script>

<style scoped lang="less">
.content-l {
    width: 25%;
    max-width: 350px;
    min-width: 250px;
    padding: 30px 20px;
    .container {
        >.title {
            color: rgb(16, 16, 16);
            font-size: 18px;
            font-weight: bold;
            margin-bottom: 20px;
        }

        .block-c {
            margin-top: 8px;

            .title-c {
                display: flex;
                align-items: flex-end;
                margin-bottom: 8px;

                .title {
                  color: rgb(16, 16, 16);
                  font-size: 14px;
                }

                .clear-btn {
                  text-decoration: underline;
                  font-size: .875rem;
                  cursor: pointer;
                  font-size: 12px;
                  margin-left: 8px;
                  color: rgb(156, 163, 175);
                  .fill:not([stroke]) {
                    fill: rgb(0, 102, 255);
                  }
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
                border: 1px solid rgba(145,213,255,0.5);
                background-color: rgba(255,255,255,1);

                &.active {
                  background-color: rgba(0,102,255,1);
                  color: #fff;
                }

                &.second-style {
                  border: 1px solid rgba(255, 198, 145, 0.5);
                  background-color: rgba(255,255,255,1);
              
                  &.active {
                    background-color: rgba(0,102,255,1);
                    color: #fff;
                  }
                }
                &.third-style {
                  border: 1px solid rgba(169, 223, 184, 0.5);
                  background-color: rgba(255,255,255,1);
                  
                  &.active {
                    background-color: rgba(0,102,255,1);
                    color: #fff;
                  }
                }
              }
            }
        }
    }
}

/* 保持原有样式结构 */
@media only screen and (max-width: 849.99px) {	
  .content .content-l {	
    display: none;	
  }	
}
</style>