<template>
  <div class="filters-container">
    <el-input
      :placeholder="placeholder"
      v-model="localKeyword"
      @keyup.native.enter="handleSearch"
      style="width:50%"
    >
      <el-button slot="append" @click="handleSearch">{{ searchBtnText }}</el-button>
    </el-input>

    <el-dropdown 
      ref="filterOrder"
      class="filter-order"
      :class="{ 'filter-order-active': isDropdownActive }"
      trigger="click" 
      @visible-change="onDropdownVisibleChange"
      @click.native="handleDropdownClick"
    >
      <span class="el-dropdown-link">
        {{ currentSortLabel }}<i class="el-icon-arrow-down el-icon--right"></i>
      </span>
      <el-dropdown-menu slot="dropdown" class="order-item">
        <el-dropdown-item
          v-for="item in sortOptions"
          :key="item.key"
          :class="{ active: item.key === currentSortKey }"
          @click.native="selectSort(item)"
        >
          {{ item.label }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
  </div>
</template>

<script>
export default {
  name: 'Filters',
  props: {
    // 搜索相关
    placeholder: { type: String, default: '' },
    searchBtnText: { type: String, default: 'Search' },

    // 排序相关
    sortOptions: {
      type: Array,
      required: true,
      default: () => []
    },
    defaultSortKey: { type: String, default: 'default' }
  },
  data() {
    return {
      localKeyword: '',
      currentSortKey: this.defaultSortKey,
      currentSortLabel: '',
      isDropdownActive: false,
      isDropdownClicked: false
    }
  },
  mounted() {
    // 添加全局点击事件监听器，用于检测点击外部区域
    document.addEventListener('click', this.handleClickOutside);
  },
  beforeDestroy() {
    // 移除事件监听器，防止内存泄漏
    document.removeEventListener('click', this.handleClickOutside);
  },
  watch: {
    sortOptions: {
      immediate: true,
      handler(newVal) {
        const defaultItem = newVal.find(item => item.key === this.defaultSortKey);
        if (defaultItem) {
          this.currentSortLabel = defaultItem.label;
        }
      }
    },
    // 监听下拉菜单显示状态变化
    isDropdownActive(newVal) {
      if (!newVal && this.isDropdownClicked) {
        // 当下拉菜单关闭且曾经被点击过时，激活字体颜色样式
        setTimeout(() => {
          this.isDropdownClicked = true; // 保持点击状态
        }, 10);
      }
    }
  },
  methods: {
    handleSearch() {
      this.$emit('search', {
        keyword: this.localKeyword.trim(),
        order_by: this.currentSortKey
      });
    },
    selectSort(item) {
      this.currentSortKey = item.key;
      this.currentSortLabel = item.label;
      this.handleSearch(); // 自动触发搜索
    },
    onDropdownVisibleChange(visible) {
      this.isDropdownActive = visible;
      if (visible) {
        this.isDropdownClicked = true;
      }
    },
    handleDropdownClick(event) {
      // 阻止事件冒泡，避免被handleClickOutside处理
      event.stopPropagation();
      this.isDropdownClicked = true;
    },
    handleClickOutside(event) {
      // 如果点击的不是filter-order元素或其子元素
      const filterOrderElement = this.$refs.filterOrder?.$el;
      if (filterOrderElement && !filterOrderElement.contains(event.target)) {
        this.isDropdownClicked = false;
      }
    },
    // 提供外部重置方法（可选）
    reset() {
      this.localKeyword = '';
      this.currentSortKey = this.defaultSortKey;
      const defaultItem = this.sortOptions.find(item => item.key === this.defaultSortKey);
      this.currentSortLabel = defaultItem ? defaultItem.label : '';
      this.isDropdownClicked = false;
    }
  }
}
</script>

<style scoped lang="less">
.filters-container {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 1.2rem;
  padding-bottom: 2rem;
}

.filter-order {
  position: absolute;
  right: 0;
  color: rgba(16, 16, 16, 1);
  cursor: pointer;

  // width: 113px;
  height: 36px;
  display: flex;
  align-content: center;
  justify-content: center;
  flex-direction: row;
  flex-wrap: wrap;
  border-bottom: 1px solid transparent;
  transition: all 0.3s ease;
  padding: 0 10px;

  &:hover {
    background-color: rgba(0,0,0,.05);
  }
}

.filter-order-active {
  background-color: rgba(0,0,0,.05);
  border-bottom: 1px solid #40a9ff !important;
}

.filter-order-clicked {
  color: #40a9ff !important;
}

::v-deep .el-dropdown-link {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  transition: color 0.3s ease;
}

// 当元素被点击过且不在激活状态时，应用字体颜色样式
.filter-order:not(.filter-order-active) {
  &.filter-order-clicked-style {
    ::v-deep .el-dropdown-link {
      color: #40a9ff !important;
    }
  }
}

.order-item {
  // width: 130px;
  left: 1403px;
  border: 1px solid rgba(34, 36, 38, .15);
  border-radius: 5px;
  top: 176px;
  white-space: nowrap;
  min-width: 170px !important;
  
  ::v-deep .el-dropdown-menu__item {
    line-height: 1em;
    color: rgba(0, 0, 0, .87) !important;
    min-height: 2.57142857rem;
    text-align: left;
    font-size: 1em !important;
    padding: .78571429em 1.14285714em !important;
  }
  ::v-deep .popper__arrow {
    display: none
  }
}

::v-deep .el-dropdown-menu__item.active {
  color: rgba(0,0,0,.95) !important;
  font-weight: 700 !important;
  background-color: rgba(0,0,0,.03) !important;
  border-bottom: 1px solid #409eff;
}
::v-deep .el-dropdown-menu__item:hover {
  background-color: rgba(0,0,0,.03) !important;
}

::v-deep .el-input__inner {
  height: 38px;
}

::v-deep .el-input-group__append {
  background-color: rgba(16,16,16,0.05);
  color: rgba(16,16,16,1);
  font-size: 14px;
  width: 70px;
  height: 38px;
  text-align: center;
}

::v-deep .el-input-group {
  width: 62.5%;
}

@media only screen and (min-width: 992px) {
  .filters-container {
    width: 100%;
  }
}
@media only screen and (max-width: 767px) {
  .filters-container {
    justify-content: space-between;
  }
}
</style>