<template>
  <div class="container">
    <div class="title">{{ $t('taskTmplObj.taskTmpl') }}</div>
    <div class="block-c">
      <div class="title-c">
        <span class="title">{{ $t('resourcesManagement.jobType') }}</span>
        <span class="clear-btn" v-if="taskTypeFlag" @click="clearSelectLeft('taskType')">
          <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 32 32" width="12" height="12"><defs></defs><g><path fill="rgb(0, 102, 255)" d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"></path></g></svg>
        </span>
      </div>
      <div class="list task-type-c">
        <div class="item" :class="taskTypeValue == item.k ? 'active' : ''" v-for="(item, index) in TaskType"
          :key="item.k" @click="selectTaskType(item)">{{ item.v }}</div>
      </div>
    </div>
    <div class="block-c">
      <div class="title-c">
        <span class="title">{{ $t('resourcesManagement.computeResource') }}</span>
        <span class="clear-btn" v-if="resourceFlag" @click="clearSelectLeft('resource')">
          <svg xmlns="http://www.w3.org/2000/svg" class="fill" viewBox="0 0 32 32" width="12" height="12"><defs></defs><g><path d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"></path></g></svg>
        </span>
      </div>
      <div class="list resource-c">
        <div class="item" :class="resourceValue == item.k ? 'active' : ''" v-for="(item, index) in Resource"
          :key="item.k" @click="selectResource(item)">{{ item.v }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import { TmplTaskTypes, TmplComputerResouces } from '~/pages/aitasktmpl/tools';

export default {
  name: "Filters",
  props: {
    condition: { type: Object, default: () => ({}) },
  },
  components: {},
  data() {
    return {
      resourceFlag: false,
      taskTypeFlag: false,
      TaskType: [...TmplTaskTypes],
      Resource: [...TmplComputerResouces],
      resourceValue: '',
      taskTypeValue: ''
    };
  },
  methods: {
    selectTaskType(item) {
      this.taskTypeValue = item.k
      this.taskTypeFlag = true
      this.search()
    },
    selectResource(item) {
      this.resourceFlag = true
      this.resourceValue = item.k
      this.search()
    },
    clearSelectLeft(type) {
      if (type === 'resource') {
        this.resourceValue = ''
        this.resourceFlag = false
      } else if (type === 'taskType') {
        this.taskTypeValue = ''
        this.taskTypeFlag = false
      }
      this.search()
    },
    search() {
      this.$emit('changeCondition', {
        job_type: this.taskTypeValue,
        compute_source: this.resourceValue
      });
    }
  },
  beforeMount() { },
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

      &.resource-c {
        .item {
          border: 1px solid rgba(145, 213, 255, 0.5);
          background: linear-gradient(92.08deg, rgba(226, 245, 255, 1) -0.83%, rgba(247, 252, 255, 0.5) 19.83%, rgba(247, 252, 255, 0.5) 97.53%);

          &.active {
            background: rgba(255, 255, 255, 1);
            border: 1px solid rgba(0, 158, 255, 1);
          }
        }
      }

      &.task-type-c {
        .item {
          border: 1px solid rgba(255, 198, 145, 0.5);
          background: linear-gradient(92.08deg, rgba(255, 245, 226, 1) -0.83%, rgba(255, 253, 247, 0.5) 19.83%, rgba(255, 253, 247, 0.5) 97.53%);

          &.active {
            background: rgba(255, 255, 255, 1);
            border: 1px solid rgba(255, 122, 0, 1);
          }
        }
      }
    }
  }
}
</style>
