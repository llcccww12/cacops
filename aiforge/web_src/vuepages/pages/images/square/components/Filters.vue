<template>
  <div>
    <div class="top-title">{{ $t('imagesObj.cloudbrain_images') }}</div>
    <div class="filter-main" v-show="!showSecond">
      <div class="filter-c" v-for="(item, index) in mainData" :key="index"
        v-show="item.key != 'framework_version' || item.data.length > 0">
        <div class="filter-title-c">
          <div class="filter-title">{{ item.title }}</div>
          <span class="clear-btn" v-if="conds[item.key]" @click="clearSelect(item.key)">
            <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 32 32" width="12" height="12"><defs></defs><g><path fill="rgb(0, 102, 255)" d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"></path></g></svg>
          </span>
        </div>
        <div class="filter-item-c">
          <div class="filter-item" :class="[item.key, conds[item.key] == _item.k ? 'active' : '']"
            v-for="(_item) in item.showData" @click="changeFilter(item, _item)" :key="_item.k" :style="conds[item.key] == _item.k ?
              { backgroundColor: _item.focusBgColor || item.focusBgColor, color: _item.focusColor || item.focusColor }
              : { borderColor: _item.bgColor || item.bgColor, color: _item.color || item.color }">
            {{ _item.v }}
          </div>
        </div>
        <div class="filter-view-more-c" v-if="item.data.length > item.showMaxLen">
          <span class="filter-view-more" @click="goMore(item)">
            <i class="el-icon-arrow-down"></i><span>{{ $t('expandMore') }}</span>
          </span>
        </div>
      </div>
    </div>
    <div class="filter-second" v-show="showSecond">
      <div class="filter-second-hd">
        <span class="filter-second-go-back" @click="goBack()">
          <i class="el-icon-back"></i><span>{{ $t('goBack') }}</span>
        </span>
      </div>
      <div class="filter-title">{{ secondData.title }}</div>
      <div class="filter-search">
        <el-input :placeholder="$t('repos.search')" prefix-icon="el-icon-search" v-model="searchKeyword" clearable
          @input="searchFilterItem">
        </el-input>
      </div>
      <div class="filter-item-c">
        <div class="filter-item" :class="[secondData.key, conds[secondData.key] == _item.k ? 'active' : '']"
          v-for="(_item) in secondData.secondShowData" :key="_item.k" @click="changeFilter(secondData, _item)" :style="conds[secondData.key] == _item.k ?
            { backgroundColor: _item.focusBgColor || secondData.focusBgColor, color: _item.focusColor || secondData.focusColor }
            : { borderColor: _item.bgColor || secondData.bgColor, color: _item.color || secondData.color }">
          {{ _item.v }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getStaticFile } from '~/apis/modules/common';
import { COMPUTER_RESOURCES_COLORS } from '~/const';

export default {
  name: "Filters",
  props: {
    condition: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {
      mainData: [{
        title: this.$t('imagesObj.imageTaskType'),
        key: 'trainType',
        bgColor: 'rgba(145, 213, 255, 0.5)',
        color: 'rgb(65, 80, 88)',
        focusBgColor: 'rgb(3, 102, 214)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 10,
      }, {
        title: this.$t('cloudbrainObj.computeResource'),
        key: 'compute_resource',
        bgColor: 'rgba(255, 198, 145, 0.5)',
        color: 'rgb(65, 80, 88)',
        focusBgColor: 'rgb(3, 102, 214)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 10,
      },{
        title: this.$t('imagesObj.framework'),
        key: 'framework',
        bgColor: 'rgba(169, 223, 184, 0.5)',
        color: 'rgb(65, 80, 88)',
        focusBgColor: 'rgb(3, 102, 214)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 10,
      }, {
        title: this.$t('imagesObj.frameworkVersion'),
        key: 'framework_version',
        bgColor: 'rgba(213, 176, 242, 0.5)',
        color: 'rgb(65, 80, 88)',
        focusBgColor: 'rgb(3, 102, 214)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 20,
      }, {
        title: this.$t('imagesObj.pyVersion'),
        key: 'python',
        bgColor: 'rgba(244, 179, 212, 0.5)',
        color: 'rgb(65, 80, 88)',
        focusBgColor: 'rgb(3, 102, 214)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 20,
      }, {
        title: this.$t('imagesObj.cudaVersion'),
        key: 'cuda',
        bgColor: 'rgba(215, 176, 240, 0.5)',
        color: 'rgb(65, 80, 88)',
        focusBgColor: 'rgb(3, 102, 214)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 20,
      }, {
        title: this.$t('imagesObj.cannVersion'),
        key: 'cann',
        bgColor: 'rgba(92, 28, 138, 0.5)',
        color: 'rgb(65, 80, 88)',
        focusBgColor: 'rgb(3, 102, 214)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 20,
      }, {
        title: this.$t('imagesObj.dtkVersion'),
        key: 'dtk',
        bgColor: 'rgba(119, 199, 43, 0.5)',
        color: 'rgb(65, 80, 88)',
        focusBgColor: 'rgb(3, 102, 214)',
        focusColor: 'rgb(255, 255, 255)',
        data: [],
        showData: [],
        showMaxLen: 20,
      }],
      filterData: {},
      conds: {
        trainType: '',
        compute_resource: '',
        framework: '',
        framework_version: '',
        python: '',
        cuda: '',
        cann: '',
        dtk: ''
      },
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
      console.log(this.secondData)
    },
    goBack() {
      this.showSecond = false;
    },
    clearSelect(key) {
      this.conds[key] = ''
      if (key == 'framework') {
        const frameworkVersionObj = this.mainData.find(itm => itm.key == 'framework_version');
        frameworkVersionObj.data = [];
        frameworkVersionObj.showData = [];
      }
      this.$emit('changeCondition', this.conds);
    },
    changeFilter(item, _item) {
      const value = this.conds[item.key] == _item.k ? '' : _item.k;
      const changes = {
        [item.key]: value,
      };
      if (item.key == 'framework') {
        const frameworkVersionObj = this.mainData.find(itm => itm.key == 'framework_version');
        changes['framework_version'] = '';
        frameworkVersionObj.data = value ? ((this.filterData['framework_version'] || {})[value] || []).map(item => ({ k: item, v: item })) : [];
        frameworkVersionObj.showData = frameworkVersionObj.data.slice(0, frameworkVersionObj.showMaxLen);
      }
      this.conds = {
        ...this.conds,
        ...changes,
      }
      this.$emit('changeCondition', this.conds);
    },
    searchFilterItem() {
      const keyword = this.searchKeyword.trim().toLocaleLowerCase();
      this.secondData.secondShowData = this.secondData.data.filter(item => {
        return item.v.toString().toLocaleLowerCase().indexOf(keyword) >= 0;
      });
    },
  },
  beforeMount() {
    getStaticFile('/images_version.json').then(res => {
      const data = res.data;
      if (data) {
        this.filterData = data;
        for (let i = 0, iLen = this.mainData.length; i < iLen; i++) {
          const filterItem = this.mainData[i];
          const key = filterItem.key;
          const max = filterItem.showMaxLen;
          if (data[key] && key != 'framework_version') {
            if (key == 'compute_resource') {
              filterItem.data = data[key].map(item => ({
                k: item,
                v: this.$t('computeResourceTitle.' + item) || item,
                focusBgColor: COMPUTER_RESOURCES_COLORS[item],
              }));
            } else if(key == 'trainType'){
              filterItem.data = data[key].map(item => ({
                k: item,
                v: this.$t('TaskTypeTitle.' + item) || item,
              }));
            } else {
              filterItem.data = data[key].map(item => ({ k: item, v: item }));
            }
            if (filterItem.sortOr) {
              filterItem.data = filterItem.data.sort((a, b) => a.v.localeCompare(b.v));
            }
            filterItem.showData = filterItem.data.slice(0, max);
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
.top-title {
  color: rgb(16, 16, 16);
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 20px;
}

.filter-c {
  margin-bottom: 16px;

  .filter-title-c {
    display: flex;
    align-items: flex-end;
    margin-bottom: 12px;

    .filter-title {
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
    }
  }

  .filter-item-c {
    display: flex;
    flex-wrap: wrap;

    .filter-item {
      border-radius: 4px;
      padding: 3px 10px;
      margin-right: 10px;
      margin-bottom: 10px;
      cursor: pointer;
      border: 1px solid rgba(145, 213, 255, 0.5);
      box-shadow: 0px 1px 1px 0px rgba(16, 16, 16, 0.2);
      background: #fff;
      font-size: 12px;
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
    color: rgb(16, 16, 16);
    font-size: 14px;
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
