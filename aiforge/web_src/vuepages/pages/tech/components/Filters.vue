<template>
  <div>
    <div class="filter-main" v-show="!showSecond">
      <div class="filter-c" v-for="(item, index) in mainData">
        <div class="filter-title">{{ item.title }}</div>
        <div class="filter-item-c">
          <div class="filter-item" v-for="(_item, _index) in item.showData" @click="changeFilter(item, _item)"
            :style="conds[item.key] == _item ? { backgroundColor: item.focusBgColor, color: item.focusColor } : { backgroundColor: item.bgColor, color: item.color }">
            {{ _item }}
          </div>
        </div>
        <div class="filter-view-more-c" v-if="item.data.length > item.showMaxLen">
          <span class="filter-view-more" @click="goMore(item)">
            <i class="el-icon-arrow-down"></i><span>展开更多</span>
          </span>
        </div>
      </div>
    </div>
    <div class="filter-second" v-show="showSecond">
      <div class="filter-second-hd">
        <span class="filter-second-go-back" @click="goBack()">
          <i class="el-icon-back"></i><span>返回上一级</span>
        </span>
      </div>
      <div class="filter-title">{{ secondData.title }}</div>
      <div class="filter-search">
        <el-input placeholder="搜索" prefix-icon="el-icon-search" v-model="searchKeyword" clearable
          @input="searchFilterItem">
        </el-input>
      </div>
      <div class="filter-item-c">
        <div class="filter-item" v-for="(_item, _index) in secondData.secondShowData" @click="changeFilter(secondData, _item)"
          :style="conds[secondData.key] == _item ? { backgroundColor: secondData.focusBgColor, color: secondData.focusColor } : { backgroundColor: secondData.bgColor, color: secondData.color }">
          {{ _item }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getTechFilterInfo } from '~/apis/modules/tech';

export default {
  name: "Filters",
  props: {
    type: { type: Number, default: -1 }, // 0-tech_view, 1-repo_view
    condition: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {
      list_tech: [{
        title: '项目类型',
        key: 'type_name',
        bgColor: 'rgb(237, 234, 251)',
        color: 'rgb(100, 59, 159)',
        focusBgColor: 'rgb(100, 59, 159)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 6,
      }, {
        title: '项目参与单位',
        key: 'institution_name',
        bgColor: 'rgb(234, 241, 251)',
        color: 'rgb(18, 76, 157)',
        focusBgColor: 'rgb(18, 76, 157)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 10,
        sortOr: true,
      }, {
        title: '执行周期包含年份',
        key: 'execute_year',
        bgColor: 'rgb(225, 242, 234)',
        color: 'rgb(8, 96, 96)',
        focusBgColor: 'rgb(8, 96, 96)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 20,
      }, {
        title: '申报年份',
        key: 'apply_year',
        bgColor: 'rgb(231, 249, 222)',
        color: 'rgb(55, 94, 2)',
        focusBgColor: 'rgb(55, 94, 2)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 20,
      }],
      list_repo: [{
        title: '关键词',
        key: 'topic',
        bgColor: 'rgb(234, 250, 251)',
        color: 'rgb(0, 167, 132)',
        focusBgColor: 'rgb(0, 167, 132)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 10,
      }, {
        title: '所属科技项目',
        key: 'project_name',
        bgColor: 'rgb(234, 245, 251)',
        color: 'rgb(8, 0, 148)',
        focusBgColor: 'rgb(8, 0, 148)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 8,
        sortOr: true,
      }, {
        title: '成果贡献单位',
        key: 'institution_name',
        bgColor: 'rgb(234, 241, 251)',
        color: 'rgb(18, 76, 157)',
        focusBgColor: 'rgb(18, 76, 157)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 10,
        sortOr: true,
      },
      ],
      conds: {
        type_name: '',
        institution_name: '',
        execute_year: '',
        apply_year: '',

        topic: '',
        project_name: '',
      },
      mainData: [],
      showSecond: false,
      secondData: {},
      searchKeyword: '',
    };
  },
  methods: {
    goMore(item) {
      this.secondData = item;
      this.searchKeyword = '';
      this.secondData.secondShowData = item.data;
      this.showSecond = true;
    },
    goBack() {
      this.showSecond = false;
    },
    changeFilter(item, _item) {
      const value = this.conds[item.key] == _item ? '' : _item;
      this.$emit('changeCondition', {
        [item.key]: value
      });
    },
    searchFilterItem() {
      const keyword = this.searchKeyword.trim().toLocaleLowerCase();
      this.secondData.secondShowData = this.secondData.data.filter(item => {
        return item.toString().toLocaleLowerCase().indexOf(keyword) >= 0;
      });
    },
  },
  watch: {
    condition: {
      handler(newVal) {
        this.conds.type_name = newVal.type_name || '';
        this.conds.institution_name = newVal.institution_name || '';
        this.conds.execute_year = newVal.execute_year || '';
        this.conds.apply_year = newVal.apply_year || '';

        this.conds.topic = newVal.topic || '';
        this.conds.project_name = newVal.project_name || '';
      },
      immediate: true,
      deep: true,
    },
  },
  beforeMount() {
    if (this.type == 0) {
      this.mainData = this.list_tech;
    } else if (this.type == 1) {
      this.mainData = this.list_repo;
    }
    getTechFilterInfo({
      type: this.type,
    }).then(res => {
      const data = res.data;
      if (data) {
        for (let i = 0, iLen = this.mainData.length; i < iLen; i++) {
          const filterItem = this.mainData[i];
          const key = filterItem.key;
          const max = filterItem.showMaxLen;
          if (data[key]) {
            filterItem.data = data[key].map(item => item.toString());
            if (filterItem.sortOr) {
              filterItem.data = filterItem.data.sort((a, b) => a.localeCompare(b));
            }
            filterItem.showData = filterItem.data.slice(0, max);
            if (this.conds[key] && filterItem.showData.indexOf(this.conds[key]) < 0) {
              filterItem.showData.push(this.conds[key]);
            }
          }
        }
      }
    }).catch(err => {
      console.log(err);
    });
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.filter-c {
  margin-bottom: 32px;

  .filter-title {
    font-size: 18px;
    color: rgb(16, 16, 16);
    margin: 14px 0;
  }

  .filter-item-c {
    display: flex;
    flex-wrap: wrap;

    .filter-item {
      border-radius: 3px;
      padding: 3px 10px;
      margin-right: 10px;
      margin-bottom: 10px;
      cursor: pointer;
    }
  }

  .filter-view-more-c {
    margin-top: 6px;

    .filter-view-more {
      color: rgb(50, 145, 248);
      font-size: 14px;
      cursor: pointer;

      i {
        margin-right: 4px;
      }
    }
  }
}

.filter-second {
  .filter-second-hd {
    margin: 14px 0 24px 0;

    .filter-second-go-back {
      color: rgb(50, 145, 248);
      font-size: 14px;
      cursor: pointer;

      i {
        margin-right: 4px;
      }
    }
  }

  .filter-title {
    font-size: 18px;
    color: rgb(16, 16, 16);
    margin: 14px 0;
  }

  .filter-search {
    margin: 5px 10px 10px 0;
  }

  .filter-item-c {
    display: flex;
    flex-wrap: wrap;
    max-height: 605px;
    overflow-y: auto;

    .filter-item {
      border-radius: 3px;
      padding: 3px 10px;
      margin-right: 10px;
      margin-bottom: 10px;
      cursor: pointer;
    }
  }
}
</style>
