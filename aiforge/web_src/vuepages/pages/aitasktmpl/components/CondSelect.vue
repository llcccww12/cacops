<template>
  <div class="cond-select-c">
    <el-popover v-model="show" placement="bottom-start" width="422" trigger="click" popper-class="cond-select-pop">
      <div class="pop-panel">
        <div class="search-bar">
          <el-input :placeholder="searchPlaceholder" v-model="keyword" @keyup.enter.native="search">
            <i slot="suffix" class="el-input__icon el-icon-search" @click="search"></i>
          </el-input>
        </div>
        <div class="list">
          <div class="list-item" v-for="(item) in list" :key="item.id" :title="item.name"
            :class="(selected && selected.id == item.id ? 'active' : '')" @click="changeSelect(item)">{{ item.fullname
            }}</div>
        </div>
      </div>
      <div class="cond-select" :class="[show ? 'show' : '', selected ? 'selected' : '']" slot="reference">
        <span>{{ selected ? selected.name : placeholder }}</span>
        <i class="el-icon-arrow-up"></i>
        <i class="el-icon-circle-close" @click.stop.prevent="clear"></i>
      </div>
    </el-popover>
  </div>
</template>

<script>
import { doSearchAiTaskTmplConds, getDefaultAiTaskTmplConds } from '~/apis/modules/aitasktmpl';

export default {
  name: "CondSelect",
  props: {
    type: { type: String, default: 'model' } // model|dataset|repo
  },
  components: {},
  data() {
    return {
      show: false,
      placeholder: '',
      searchPlaceholder: '',
      selected: null,
      keyword: '',
      defaultList: [],
      list: [],
      page: 1,
      pageSize: 50,
      qureying: false,
      timer: null,
    };
  },
  methods: {
    getDefaultData() {
      let configPath = '';
      if (this.type == 'model') {
        configPath = 'home_v2/model.json'
      }
      if (this.type == 'dataset') {
        configPath = 'home_v2/dataset.json'
      }
      if (this.type == 'repo') {
        return;
      }
      getDefaultAiTaskTmplConds({ key: configPath }).then(res => {
        res = res.data;
        const list = res.data || [];
        this.defaultList = list.map((item) => {
          return {
            id: item.ID,
            fullname: `${item.OwnerName}/${item.Name}`,
            name: item.Alias || item.Name
          }
        })
        this.list = [...this.defaultList];
      }).catch(err => {
        console.log(err);
      })
    },
    getData() {
      this.qureying = true;
      this.timer && clearTimeout(this.timer);
      this.timer = setTimeout(() => {
        doSearchAiTaskTmplConds({
          type: this.type,
          q: this.keyword,
          page: this.page,
          pageSize: this.pageSize
        }).then(res => {
          res = res.data;
          const list = res.Result || [];
          if (this.type == 'model') {
            this.list = list.map((item) => {
              const showName = item.title.replaceAll("<font color='red'>", '').replaceAll('</font>', '') || item.real_name;
              return {
                id: item.id,
                fullname: `${item.owerName}/${showName}`,
                name: showName,
              }
            })
          } else if (this.type == 'dataset') {
            this.list = list.map((item) => {
              const showName = item.title.replaceAll("<font color='red'>", '').replaceAll('</font>', '') || item.name;
              return {
                id: item.id,
                fullname: `${item.owerName}/${showName}`,
                name: showName,
              }
            })
          } else if (this.type = 'repo') {
            this.list = list.map((item) => {
              const showName = item.lower_alias || item.real_name;
              return {
                id: item.id,
                fullname: `${item.owner_name}/${showName}`,
                name: showName,
              }
            })
          }
          this.total = res.Total;
        }).catch(err => {
          console.log(err)
        }).finally(() => {
          this.qureying = false;
        })
      }, 100)
    },
    search() {
      if (this.qureying) return;
      this.keyword = this.keyword.trim();
      if (!this.keyword) {
        this.list = [...this.defaultList];
        return;
      }
      this.getData();
    },
    changeSelect(item) {
      this.selected = item;
      this.show = false;
      this.$emit('change', item);
    },
    clear() {
      this.selected = null;
      this.$emit('change', null);
    }
  },
  beforeMount() {
    switch (this.type) {
      case 'repo':
        this.placeholder = this.$t('repos.repos');
        this.searchPlaceholder = this.$t('org.searchRepos');
        break;
      case 'dataset':
        this.placeholder = this.$t('repos.dataset');
        this.searchPlaceholder = this.$t('org.searchDatasets');
        break;
      case 'model':
      default:
        this.placeholder = this.$t('repos.model');
        this.searchPlaceholder = this.$t('org.searchModels');
        break;
    }
  },
  mounted() {
    this.getDefaultData();
  },
};
</script>

<style scoped lang="less">
.cond-select {
  position: relative;
  height: 32px;
  width: 140px;
  height: 32px;
  line-height: 32px;
  cursor: pointer;
  border-radius: 4px;
  background-color: rgba(16, 16, 16, 0.05);
  border: 1px solid #DCDFE6;
  box-sizing: border-box;
  color: #606266;
  padding: 0 15px;
  padding-right: 30px;
  display: flex;
  align-items: center;

  span {
    display: inline-block;
    overflow: hidden;
    height: 100%;
    white-space: nowrap;
  }

  .el-icon-circle-close,
  .el-icon-arrow-up {
    right: 8px;
    transition: all 0.3s;
    position: absolute;
    top: 8px;
    color: #C0C4CC;
    text-align: center;
    transform: rotate(180deg);
  }

  .el-icon-circle-close {
    display: none;
  }

  &.show {
    background-color: #FFF;
    border: 1px solid #409EFF;

    .el-icon-arrow-up {
      transform: rotate(0deg);
    }
  }

  &.selected {
    background-color: #FFF;

    &:hover {
      .el-icon-circle-close {
        display: block;
      }

      .el-icon-arrow-up {
        display: none;
      }
    }
  }
}

.pop-panel {
  .search-bar {
    padding: 0 12px;

    .el-icon-search {
      cursor: pointer;
      color: rgba(16, 16, 16, 1);
    }
  }

  .list {
    margin-top: 5px;
    max-height: 256px;
    overflow-y: auto;

    .list-item {
      height: 32px;
      line-height: 32px;
      cursor: pointer;
      padding: 0 16px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;

      &:hover {
        background-color: #F5F7FA;
      }

      &.active {
        color: #409EFF;
      }
    }

    &::-webkit-scrollbar {
      -webkit-appearance: none;
      width: 7px;
      height: 7px
    }

    &::-webkit-scrollbar-track {
      background: transparent;
      border-radius: 7px !important;
    }

    &::-webkit-scrollbar-thumb {
      background: rgb(210, 210, 216) !important;
      border-radius: 7px !important;
    }

    &::-webkit-scrollbar-thumb:hover {
      background: rgba(210, 210, 216, 0.8) !important;
      border-radius: 7px !important;
    }

    &::-webkit-scrollbar-thumb:active {
      background: rgba(210, 210, 216, 0.8) !important;
      border-radius: 7px !important;
    }
  }

  .summary {
    height: 32px;
    display: flex;
    align-items: center;
    padding: 0 16px;
    color: rgba(16, 16, 16, 0.5);
  }
}
</style>
<style lang="less">
.cond-select-pop {
  margin-top: 8px !important;
  padding-left: 0 !important;
  padding-right: 0 !important;
  padding-bottom: 6px !important;
}
</style>
