<template>
  <div>
    <div class="_repo_search">
      <div class="_repo_search_input_c">
        <div class="_repo_search_input">
          <input type="text" v-model="searchInputValue" :placeholder="$t('repos.searchRepositories')" autocomplete="off"
            @keyup.enter="search" />
        </div>
        <div class="_repo_search_btn" @click="search">
          <svg xmlns="http://www.w3.org/2000/svg"
            class="styles__StyledSVGIconPathComponent-sc-16fsqc8-0 kdvdTY svg-icon-path-icon fill" viewBox="0 0 32 32"
            width="24" height="24">
            <defs data-reactroot="">
              <linearGradient id="ilac9fwnydq3lcx1,1,rs,1,f0dwf0xj,ezu9f0dw,f000,00lwrsktrs,rs1bhv8urs" x1="0" x2="100%"
                y1="0" y2="0"
                gradientTransform="matrix(-0.7069999999999999, -0.707, 0.707, -0.7069999999999999, 16, 38.624)"
                gradientUnits="userSpaceOnUse">
                <stop stop-color="#c9ffbf" stop-opacity="1" offset="0"></stop>
                <stop stop-color="#0ca451" stop-opacity="1" offset="1"></stop>
              </linearGradient>
            </defs>
            <g>
              <path fill="url(#ilac9fwnydq3lcx1,1,rs,1,f0dwf0xj,ezu9f0dw,f000,00lwrsktrs,rs1bhv8urs)"
                d="M14.667 2.667c6.624 0 12 5.376 12 12s-5.376 12-12 12-12-5.376-12-12 5.376-12 12-12zM14.667 24c5.156 0 9.333-4.177 9.333-9.333 0-5.157-4.177-9.333-9.333-9.333-5.157 0-9.333 4.176-9.333 9.333 0 5.156 4.176 9.333 9.333 9.333zM25.98 24.095l3.772 3.771-1.887 1.887-3.771-3.772 1.885-1.885z">
              </path>
            </g>
          </svg>
          <span style="margin-left:10px;">{{ $t('repos.search') }}</span>
        </div>
      </div>
      <div class="_repo_search_label_c">
        <div class="_repo_search_label">
          <a v-if="type == 'square'" :href="`/explore/repos?q=${searchInputValue.trim()}&topic=${item.v}&sort=${sort}`"
            :style="{ backgroundColor: topicColors[index % topicColors.length] }" v-for="(item, index) in topics"
            :key="index">{{ item.v }}</a>
          <a v-if="type == 'search'" href="javascript:;" @click="changeTopic({ k: '', v: '' })"
            style="font-weight:bold;"
            :style="{ backgroundColor: selectTopic == '' ? selectedColor : defaultColor, color: selectTopic == '' ? 'white' : '#40485b' }">{{
                $t('repos.allFields')
            }}</a>
          <a v-if="type == 'search'" href="javascript:;" @click="changeTopic(item)"
            :style="{ backgroundColor: selectTopic.toLocaleLowerCase() == item.k ? selectedColor : defaultColor, color: selectTopic.toLocaleLowerCase() == item.k ? 'white' : '#40485b' }"
            v-for="(item, index) in topics" :key="index">{{ item.v }}</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getPromoteData } from '~/apis/modules/common';
const COLOR_LIST = [
  'rgb(255, 104, 104)',
  'rgb(22, 132, 252)',
  'rgb(2, 202, 253)',
  'rgb(164, 145, 215)',
  'rgb(232, 64, 247)',
  'rgb(245, 182, 110)',
  'rgb(54, 187, 166)',
  'rgb(123, 50, 178)'
];
export default {
  name: "SearchBar",
  props: {
    type: { type: String, default: 'square' }, // square|search
    searchValue: { type: String, default: '' },
    topic: { type: String, default: '' },
    sort: { type: String, default: '' },
    static: { type: Boolean, default: false },
    staticTopicsData: { type: String, default: '[]' },
  },
  components: {},
  data() {
    return {
      searchInputValue: '',
      topicColors: COLOR_LIST,
      defaultColor: '#F6F6F6',
      selectedColor: '',
      selectTopic: '',
      topicOri: [],
      topics: [],
    };
  },
  methods: {
    setDefaultSearch(params) {
      this.searchInputValue = params.q || '';
      this.selectTopic = params.topic || '';
      this.selectTopic && this.changeTopic({
        k: this.selectTopic.toLocaleLowerCase(),
        v: this.selectTopic,
      }, true);
    },
    changeTopic(topicItem, noSearch) {      
      const index_ori = this.topicOri.findIndex((item) => {
        return item.k == this.selectTopic.toLocaleLowerCase();
      });
      if (index_ori < 0 && this.selectTopic) {
        const index = this.topics.findIndex((item) => {
          return item.k == this.selectTopic.toLocaleLowerCase();
        });
        if (index > -1) {
          this.topics.splice(index, 1);
        }
      }
      this.selectTopic = topicItem.v;
      if (this.selectTopic && this.topics.indexOf(this.selectTopic) < 0) {
        const index = this.topics.findIndex(item => {
          return item.k == this.selectTopic.toLocaleLowerCase();
        })
        if (index < 0) {
          this.topics.push({
            k: this.selectTopic.toLocaleLowerCase(),
            v: this.selectTopic,
          });
        }
      }
      !noSearch && this.search();
    },
    handlerTopicsData(data) {
      try {
        const topicsData = JSON.parse(data);
        const topics = topicsData.map((item) => {
          return {
            k: item.trim().toLocaleLowerCase(),
            v: item.trim(),
          }
        });
        this.topicOri = JSON.parse(JSON.stringify(topics));
        this.topics = topics;
        const selectTopic_key = this.selectTopic.toLocaleLowerCase();
        if (selectTopic_key) {
          const index = this.topics.findIndex((item) => {
            return item.k == selectTopic_key;
          });
          if (index < 0) {
            this.topics.push({
              k: this.selectTopic.toLocaleLowerCase(),
              v: this.selectTopic,
            });
          }
        }
      } catch (err) {
        console.log(err);
      }
    },
    search() {
      this.searchInputValue = this.searchInputValue.trim();
      if (this.type == 'square') {
        window.location.href = `/explore/repos?q=${this.searchInputValue}&sort=${this.sort}&topic=${this.selectTopic}`;
      } else {
        this.$emit('change', {
          q: this.searchInputValue,
          topic: this.selectTopic,
        });
      }
    }
  },
  mounted() {
    if (this.static) {
      try {
        this.handlerTopicsData(this.staticTopicsData);
      } catch (err) {
        console.log(err);
      }
    } else {
      getPromoteData('/repos/recommend_topics').then(res => {
        const data = res.data;
        this.handlerTopicsData(data);
      }).catch(err => {
        console.log(err);
        this.handlerTopicsData('[]');
      });
    }
  },
};
</script>

<style scoped lang="less">
._repo_search {
  margin: 54px 0;
}

._repo_search_input_c {
  margin: 0 0 35px;
  display: flex;
  justify-content: center;
  align-items: center;
}

._repo_search_input {
  display: flex;
  align-items: center;
  width: 437px;
  height: 41px;
  border-color: rgba(47, 9, 69, 0.64);
  border-width: 1px;
  border-style: solid;
  color: rgba(16, 16, 16, 0.5);
  border-radius: 20px;
  font-size: 14px;
  padding: 20px;
  text-align: left;
  line-height: 20px;
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736765e-17%2C%20-0.9999999999999999%2C%200.008802475794500678%2C%206.123233995736765e-17%2C%201%2C%201.003)%22%3E%3Cstop%20stop-color%3D%22%23eeeade%22%20stop-opacity%3D%220.2%22%20offset%3D%220.76%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23ece0e9%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
}

._repo_search_input input {
  width: 100%;
  height: 30px;
  border: none;
  background: transparent;
  outline: none;
}

._repo_search_btn {
  margin-left: 10px;
  width: 110px;
  height: 40px;
  border-style: none;
  border-color: unset;
  color: rgb(255, 255, 255);
  border-radius: 21px;
  font-size: 14px;
  text-align: center;
  font-weight: normal;
  font-style: normal;
  background: url("data:image/svg+xml;charset=utf8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20xmlns%3Axlink%3D%22http%3A%2F%2Fwww.w3.org%2F1999%2Fxlink%22%20version%3D%221.1%22%3E%3Cdefs%3E%3ClinearGradient%20id%3D%221%22%20x1%3D%220%22%20x2%3D%221%22%20y1%3D%220%22%20y2%3D%220%22%20gradientTransform%3D%22matrix(6.123233995736766e-17%2C%201%2C%20-0.17728531855955676%2C%206.123233995736766e-17%2C%201%2C%200)%22%3E%3Cstop%20stop-color%3D%22%232f0945%22%20stop-opacity%3D%221%22%20offset%3D%220%22%3E%3C%2Fstop%3E%3Cstop%20stop-color%3D%22%23341a7b%22%20stop-opacity%3D%221%22%20offset%3D%221%22%3E%3C%2Fstop%3E%3C%2FlinearGradient%3E%3C%2Fdefs%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22url(%231)%22%3E%3C%2Frect%3E%3C%2Fsvg%3E");
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

._repo_search_label_c {
  display: flex;
  justify-content: center;
}

._repo_search_label {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  max-width: 1200px;
}

._repo_search_label a {
  background-color: rgb(2, 202, 253);
  color: rgb(255, 255, 255);
  border-radius: 5px;
  margin: 0 5px 10px 5px;
  padding: 5px 10px;
  cursor: pointer;
  font-size: 12px;
}
</style>
