<template>
  <div class="event-list" v-loading="loading">
    <div class="btn-wrap" @click="refresh">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" width="20" height="20">
        <defs></defs>
        <g>
          <path fill="#202565"
            d="M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z">
          </path>
        </g>
      </svg>
      <span>{{ $t('cloudbrainObj.refresh') }}</span>
    </div>
    <div class="mesg-wrap">
      <template v-for="(item, index) in events">
        <p :key="index + '-1'"><b>[{{ item.reason }}]</b> <span>{{ item.timestampStr }}</span></p>
        <p :key="index + '-2'">{{ item.message }}</p>
      </template>
      <p v-if="events.length == 0 && !loading">{{ this.$t('noMessage') }}</p>
    </div>
  </div>
</template>

<script>
import { getAiTaskOperationProfile } from '~/apis/modules/cloudbrain';
import { formatDate } from 'element-ui/lib/utils/date-util';

export default {
  name: 'OperationProfile',
  props: {
    configs: { type: Object, default: () => { return {} } },
    data: { type: Object, default: () => { return {} } },
  },
  data() {
    return {
      events: [],
      loading: false,
    };
  },
  methods: {
    refresh() {
      // console.log('OperationProfile refresh');
      const task = this.data.task;
      if (!task) return;
      this.loading = true;
      getAiTaskOperationProfile({
        repoOwnerName: task.repoOwnerName,
        repoName: task.repoName,
        id: task.id,
      }).then(res => {
        this.loading = false;
        res = res.data;
        if (res.code == 0 && res.data) {
          const events = res.data.events || [];
          events.forEach(item => {
            item.timestampStr = item.timestamp ? formatDate(new Date(item.timestamp), 'yyyy/MM/dd HH:mm:ss') : '';
          });
          this.events = events;
        } else {
          this.events = [];
        }
      }).catch(err => {
        this.loading = false;
        console.log(err);
      });
    }
  },
  beforeMount() { }
};
</script>

<style scoped lang="less">
.event-list {
  padding: 30px 50px 24px 20px;
  margin: 0 0 20px;
  height: 100%;
  // overflow: auto;
  width: 100%;
  position: relative;
  max-height: 600px;
  overflow: auto;

  .btn-wrap {
    position: absolute;
    right: 10px;
    top: 10px;
    height: 30px;
    display: flex;
    align-items: center;
    padding: 0 10px;
    border-radius: 6px;
    background-color: rgba(16, 16, 16, 0.1);
    color: rgba(32, 37, 101, 0.9);
    cursor: pointer;
    z-index: 999;
  }

  .mesg-wrap {

    height: 100%;
    width: 100%
  }
}
</style>
