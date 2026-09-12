<template>
  <div class="container">
    <div class="title">{{ $t('repos.repos') }}</div>
    <div class="block-c">
      <div class="title-c">
        <span class="title">{{ $t('repos.repoTopics') }}</span>
        <span class="clear-btn" v-if="topicFlag" @click="clearSelectLeft('topic')">
          <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 32 32" width="12" height="12"><defs></defs><g><path fill="rgb(0, 102, 255)" d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"></path></g></svg>
        </span>
      </div>
      <div class="list topic-c">
        <div class="item" :class="item.active ? 'active' : ''" v-for="(item, index) in Topics" :key="item.k"
          @click="selectTopic(item)">{{ item.v }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import { getPromoteData } from "~/apis/modules/common";
import { lang } from '~/langs';

export default {
  name: "Filters",
  props: {},
  components: {},
  data() {
    return {
      topicFlag: false,
      Topics: [],
      topicValue: '',
      codeUsePromotePath: `/repos/recommend_topics${lang == 'zh-CN' ? '' : '_en'}`,
    };
  },
  methods: {
    selectTopic(item) {
      this.Topics.forEach(element => {
        if (element.k == item.k) {
          element.active = true
        } else {
          element.active = false
        }
      });
      this.topicFlag = true
      this.topicValue = item.k
      this.search()
    },
    clearSelectLeft(type) {
      if (type == 'topic') {
        this.Topics.forEach(element => {
          element.active = false
        });
        this.topicValue = ''
        this.topicFlag = false
      }
      this.search()
    },
    search() {
      this.$emit('changeCondition', {
        topic: this.topicValue,
      });
    },
    handlerTopicsData(data) {
      try {
        const topicsData = JSON.parse(data);
        const topics = topicsData.map((item) => {
          return {
            k: item.trim().toLocaleLowerCase(),
            v: item.trim(),
            active: false,
          }
        });
        this.Topics = topics;
      } catch (err) {
        console.log(err);
      }
    },
  },
  beforeMount() {
    getPromoteData(this.codeUsePromotePath).then(res => {
      const data = res.data;
      this.handlerTopicsData(data);
    }).catch(err => {
      console.log(err);
      this.handlerTopicsData('[]');
    });
  },
  mounted() { },
};
</script>

<style scoped lang="less">
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
      }

      &.topic-c {
        .item {
          border: 1px solid rgba(145, 213, 255, 0.5);
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
</style>
