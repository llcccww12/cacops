<template>
  <div style="height: 100%;">
    <div v-if="notFound">
      <NotFound></NotFound>
    </div>
    <div v-else style="height: 100%;">
      <div class="cb-detail-container" v-loading="loading">
        <div class="detail-header-wrap">
          <div class="area-l-title">
            <a class="link label" href="/cloudbrains">{{ $t('notebook.sameTaskTips6') }}</a>
            <span class="separator">/</span>
            <span class="label">{{ mainData[0] && mainData[0].task.display_job_name }}</span>
            <a v-if="mainData[0] && mainData[0].task.display_job_name" href="javascript:;"
              class="ui poping up clipboard" id="clipboard-btn" data-position="top center"
              data-variation="inverted tiny" :data-success="$t('copySuccess')" :data-content="$t('copy')"
              :data-original="$t('copy')" :data-clipboard-text="mainData[0].task.display_job_name">
              <i class="ri-file-copy-line"></i>
            </a>
            <div class="cb-job" :class="mainData[0] && mainData[0].task.job_type">
              <span>{{ mainData[0] && mainData[0].task.jobTypeShow }}</span>
            </div>
            <span class="task-status">
              <i :class="mainData[0] && mainData[0].task.status"></i>
              <span>{{ mainData[0] && mainData[0].task.status }}</span>
              <i v-if="mainData[0] && mainData[0].task.detailed_status === 'dataMigrating' && mainData[0].task.status === 'WAITING'"
                :class="mainData[0].task.detailed_status" :title="$t('cloudbrainObj.migratingData')"></i>
              <i v-if="mainData[0] && mainData[0].task.detailed_status === 'centerPending' && mainData[0].task.status === 'WAITING'"
                :class="mainData[0].task.detailed_status" :title="$t('cloudbrainObj.centerPending')"></i>
              <i v-if="mainData[0] && mainData[0].task.detailed_status === 'ImagePulling' && mainData[0].task.status === 'WAITING'"
                :class="mainData[0].task.detailed_status" :title="$t('cloudbrainObj.imagePulling')"></i>
            </span>
          </div>
          <div class="area-r-title" @click.stop.prevent="">
            <div class="auto-stop-tips" v-if="mainData[0] && mainData[0].task.runningLeftTime">
              <i class="ri-alarm-line"></i>
              <span>{{ $t('cloudbrainObj.autoStopTimeTips', { min: mainData[0].task.runningLeftTime }) }}</span>
            </div>
          </div>
        </div>
        <TaskDetailCollapseTabs ref="taskDetailCollapseTabsRef" :taskId="taskId" @update="updateData">
        </TaskDetailCollapseTabs>
      </div>
    </div>
  </div>
</template>

<script>
import NotFound from '~/components/NotFound.vue';
import TaskDetailCollapseTabs from '~/components/cloudbrain/details/TaskDetailCollapseTabs.vue';

export default {
  data() {
    return {
      notFound: false,
      pageCfg: {},
      taskId: '',
      mainData: [],
      loading: true,
    };
  },
  components: { NotFound, TaskDetailCollapseTabs },
  methods: {
    updateData(data) {
      console.log('updateData', data)
      this.loading = false;
      this.pageCfg = data.pageCfg;
      this.mainData = data.mainData;
      this.notFound = data.notFound;
    }
  },
  beforeMount() {
    this.taskId = window.location.pathname.split('/').pop();
  },
  mounted() { },
};
</script>

<style scoped lang="less">
.cb-detail-container {
  padding: 32px 20px 24px 40px;
  display: flex;
  flex-direction: column;
  height: 100%;
  background-color: #f9f9f9;

  .detail-header-wrap {
    display: flex;
    justify-content: space-between;
    align-items: center;

    .area-l-title {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      font-size: 18px;
      margin-bottom: 20px;

      .link {
        color: rgba(0, 102, 255, 1);
        font-weight: 400;
        cursor: pointer;
      }

      .separator {
        color: #6f76a5;
        font-family: SourceHanSansSC;
        margin-right: 10px;
      }

      .label {
        line-height: 39px;
        margin-right: 10px;
        font-family: SourceHanSansSC;
        font-weight: 500;
      }

      i {
        color: rgba(145, 145, 145, 1);
        font-size: 16px;
        margin-right: 20px;
        margin-left: 8px;
      }

      .task-status {
        display: flex;
        align-items: center;
        margin-right: 12px;
        margin-left: 10px;

        i,
        span {
          margin-right: 4px;
          font-size: 14px;
        }
      }
    }

    .area-r-title {
      display: flex;
      align-items: center;
      margin-right: 16px;
      margin-bottom: 20px;

      .auto-stop-tips {
        display: flex;
        align-items: center;
        font-size: 14px;
        color: rgba(255, 37, 37, 1);
        line-height: 39px;

        i {
          margin-right: 3px;
          font-size: 16px;
        }
      }
    }
  }

}

@media screen and (max-width: 767px) {
  .cb-detail-container {
    padding: 20px 10px 16px 10px;
  }
}
</style>
