<template>
  <div class="condition-wrap">
    <div class="filter-c-m">
      <span class="title">{{ $t('taskTmplObj.taskTmpl') }}</span>
      <div class="right-m">
        <div class="icon-box" @click="goCreate"><i class="ri-add-box-line"></i></div>
        <div class="icon-box" v-if="!showSearch" @click="showSearch = true"><i class="ri-search-line"></i></div>
        <div v-else  class="ui small icon input" style="height: 32px;margin-left: 6px;">
          <input type="text" :placeholder="$t('taskTmplObj.searchTaskTmpl')" v-model="conds.q" @keyup.enter="search">
          <i class="search icon" style="cursor: pointer;pointer-events: auto;" @click="search"></i>
        </div> 
        <div class="icon-box filter-btn" :class="{ 'active': showFilter }" @click="toggleFilter">
          <i class="ri-filter-line"></i>
        </div>
      </div>
    </div>
    <div class="condition-a">
      <div class="tab-c">
        <div class="tab-item" v-for="(item, index) in typeList" :class="conds.type == item.key ? 'focus' : ''"
          :key="item.key" @click="changeType(item)">
          {{ item.label }}
        </div>
      </div>
      <div class="condition-b">
        <div class="multi-conds">
          <CondSelect type="model" @change="changeModel" />
          <CondSelect type="dataset" @change="changeDataset" />
          <CondSelect type="repo" @change="changeRepo" />
          <el-input class="search-keyword" :placeholder="$t('taskTmplObj.searchTaskTmpl')" v-model="conds.q"
            @keyup.enter.native="search">
            <i slot="suffix" class="el-input__icon el-icon-search" @click="search"></i>
          </el-input>
        </div>
        <div class="sort-c">
          <el-select v-model="conds.order_by" @change="changeSort" style="width: 120px;">
            <el-option v-for="item in sortList" :key="item.key" :label="item.label" :value="item.key">
            </el-option>
          </el-select>
        </div>
        <el-button class="create-btn" type="primary" @click="goCreate">
          <div class="btn-content">
            <i class="ri-add-box-line"></i>
            <span>{{ $t('taskTmplObj.createTaskTmpl') }}</span>
          </div>
        </el-button>
      </div>
    </div>
  </div>
</template>

<script>
import CondSelect from '../../components/CondSelect.vue';

export default {
  name: "Condition",
  props: {
    condition: { type: Object, default: () => ({}) },
    typeList: { type: Array, default: () => [] }
  },
  components: { CondSelect },
  data() {
    return {
      isLogin: false,
      sortList: [{
        key: '',
        label: this.$t('datasets.default'),
      }, {
        key: 'newest',
        label: this.$t('datasets.newest'),
      }, {
        key: 'recentupdate',
        label: this.$t('datasets.recentupdate'),
      }, {
        key: 'collections',
        label: this.$t('datasets.moststars'),
      }, {
        key: 'usecount',
        label: this.$t('taskTmplObj.runTimes'),
      }, {
        key: 'name_asc',
        label: this.$t('datasets.alphabetasc'),
      }, {
        key: 'name',
        label: this.$t('datasets.alphabetdesc'),
      }],
      conds: {
        type: 'public',
        order_by: '',
        model: '',
        dataset: '',
        repo: '',
        q: '',
      },
      showSearch: false,
      showFilter: false,
    };
  },
  methods: {
    changeType(item) {
      this.conds.type = item.key;
      this.search();
    },
    changeModel(item) {
      this.conds.model = item?.id || '';
      this.search();
    },
    changeDataset(item) {
      this.conds.dataset = item?.id || '';
      this.search();
    },
    changeRepo(item) {
      this.conds.repo = item?.id || '';
      this.search();
    },
    changeSort(sort) {
      this.conds.order_by = sort;
      this.search();
    },
    search() {
      this.conds.q = this.conds.q.trim();
      this.$emit('changeCondition', {
        type: this.conds.type,
        order_by: this.conds.order_by,
        model_id: this.conds.model,
        dataset_id: this.conds.dataset,
        repo_id: this.conds.repo,
        q: this.conds.q,
      });
    },
    goCreate() {
      window.location.href = '/ai_task_tmpl/create';
    },
    toggleFilter(){
      this.showFilter = !this.showFilter;
      this.$emit('changeModalOverlay',this.showFilter)
    },
    resetFilter(){
      this.showFilter = false;
    }
  },
  watch: {
    condition: {
      handler(newVal) {
        this.conds.type = newVal.type || 'public';
        this.conds.order_by = newVal.order_by || '';
        this.conds.model = newVal.model_id || '';
        this.conds.dataset = newVal.dataset_id || '';
        this.conds.repo = newVal.repo_id || '';
        this.conds.q = newVal.q || '';
      },
      immediate: true,
      deep: true,
    },
  },
  beforeMount() {
    this.isLogin = !!document.querySelector('meta[name="_uid"]');
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.condition-wrap {
  margin: 10px 0 2px;

  .condition-a {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-wrap: wrap;

    .tab-c {
      display: flex;
      align-items: center;
      margin-bottom: 8px;

      .tab-item {
        height: 32px;
        display: flex;
        align-items: center;
        margin-right: 6px;
        padding: 0 6px;
        border-bottom: 2px solid rgba(51, 38, 98, 0.3);
        font-size: 16px;
        box-sizing: border-box;
        cursor: pointer;
        color: rgba(16, 16, 16, 0.5);

        &:hover {
          color: rgba(16, 16, 16, 1);
          border-color: rgba(51, 38, 98, 1);
        }

        &.focus {
          color: rgba(0, 102, 255, 1);
          border-bottom: 2px solid rgba(0, 102, 255, 1);
        }
      }
    }

    .create-btn {
      margin-left: 14px;
      display: flex;
      align-items: center;
      height: 32px;
      font-size: 14px;
      background: rgba(22, 132, 252, 0.9);
      border-radius: 4px;
      margin-bottom: 8px;

      &:active {
        background: rgb(22, 132, 252, 1);
      }


      &:focus,
      &:hover {
        background: rgba(22, 132, 252, 0.8);
      }

      .btn-content {
        display: flex;
        align-items: center;

        i {
          font-size: 14px;
          margin-right: 10px;
        }
      }
    }
  }


  .condition-b {
    display: flex;
    align-items: center;
    // justify-content: flex-end;
    flex-wrap: wrap;

    .multi-conds {
      display: flex;
      align-items: center;

      >div {
        margin-bottom: 8px;

        /deep/.cond-select,
        /deep/.el-input__inner {
          border-radius: 0;
        }

        &:first-child {

          /deep/.cond-select,
          /deep/.el-input__inner {
            border-top-left-radius: 4px;
            border-bottom-left-radius: 4px;
          }
        }

        &:last-child {

          /deep/.cond-select,
          /deep/.el-input__inner {
            border-top-right-radius: 4px;
            border-bottom-right-radius: 4px;
          }
        }
      }

      .search-keyword {
        width: 220px;
        margin-right: 15px;
        margin-bottom: 8px;

        .el-icon-search {
          cursor: pointer;
          color: rgba(16, 16, 16, 1);
        }
      }
    }

    .sort-c {
      color: rgba(0, 0, 0, 0.87);
      cursor: pointer;
      margin-bottom: 8px;
    }
  }
  .filter-c-m{
    display: none;
    margin-bottom: 14px;
    align-items: center;
    justify-content: space-between;
    .title{
      color: rgb(16,16,16);
      font-size: 18px;
      height: 32px;
      line-height: 35px;
      font-family: Arial-bold;
      font-weight: bold;
    }
    .right-m{
      display: flex;
      .icon-box{
        border: 1px solid rgba(16,16,16,0.2);
        border-radius: 5px;
        width: 32px;
        height: 32px;
        margin-left: 6px;
        display: flex;
        justify-content: center;
        align-items: center;
        i{
          color: #101010;
          font-size: 16px;
        }
      }
      .filter-btn{
        &.active{
          background-color: rgba(255,255,255,1);
          z-index: 999;
          i{
            color: #0066ff
          }
        }
      }
    }
  }  
}

/deep/.el-dropdown-menu__item.active {
  color: #409EFF;
  background-color: rgba(179, 216, 255, 0.3);
}

@media only screen and (max-width: 849.99px) {
  .condition-wrap {
    .condition-a{
      justify-content: center !important;
    }
    .filter-c-m{
      display: flex;
    }
    .condition-b{
      display: none;
    }
  }
}
</style>
