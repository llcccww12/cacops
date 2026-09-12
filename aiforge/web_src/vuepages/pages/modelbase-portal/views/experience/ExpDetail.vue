<template>
  <div class="content-box">
    <div v-if="notFound">
      <NotFound></NotFound>
    </div>
    <template v-else> 
      <Header :lastPath="lastPath" :mainData="mainData"></Header>
      <div class="main-body">
        <div class="main-content">
          <TaskDetailCollapseTabs :taskId="$route.params.taskid" @update="updateData">
          </TaskDetailCollapseTabs>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import NotFound from '~/components/NotFound.vue';
import Header from '../../components/Header.vue';
import TaskDetailCollapseTabs from '~/components/cloudbrain/details/TaskDetailCollapseTabs.vue';

export default {
  name: 'ExpDetail',
  data() {
    return {
      lastPath: null,
      task: {},
      mainData: [],
      notFound: false,
    };
  },
  components: { NotFound, Header, TaskDetailCollapseTabs },
  methods: {
    updateData(data) {
      const jobName = data.mainData[0]?.task?.display_job_name;
      if (jobName) {
        this.lastPath = {
          label: jobName,
          path: jobName,
        }
      }
      this.mainData = data.mainData;
      this.notFound = data.notFound;
    },
  },
  beforeMount() { },
  mounted() { },
};
</script>

<style scoped lang="less">
.content-box {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;

  .main-body {
    flex: 1;
    height: 0;
    position: relative;
    margin-right: 20px;

    .main-content {
      height: calc(100% - 40px);
      margin-left: 20px;
      margin-top: 20px;
      margin-bottom: 20px;
      padding: 20px;
      border-color: rgba(157, 197, 226, 0.4);
      border-width: 1px;
      border-style: solid;
      border-radius: 10px;
      background: rgb(255, 255, 255);
      display: flex;
      flex-direction: column;
      overflow: hidden;
    }
  }
}
</style>
