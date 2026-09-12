<template>
  <div class="task-limit-overlay" v-if="visible">
    <div class="task-limit-dialog">
      <div class="dialog-header">
        <div class="dialog-tip-wrap">
          <div class="icon-wrap">
            <i class="ri-error-warning-fill"></i>
          </div>
          <div class="tip-wrap">
            <p class="tip-title">{{ title }}</p>
            <p class="tip-desc" v-if="desc">{{ desc }}</p>
          </div>
        </div>
        <div class="close-wrap" @click="handleClose">
          <i class="ri-close-line"></i>
        </div>
      </div>
      <div class="task-list">
        <div v-for="item in taskList" :key="item.task.id" class="task-item">
          <div class="task-info-wrap">
            <div class="cb-job" :class="item.task.job_type">
              <span v-for="(line, i) in getJobTypeLines(item.task.jobTypeShow)" :key="i">{{ line }}</span>
            </div>
            <div class="task-info">
              <div class="task-name">{{ item.task.display_job_name }}</div>
              <div class="task-meta">
                <div class="task-status" :class="'status-' + item.task.status">
                  <div class="dot"></div>
                  <span>{{ item.task.status }}</span>
                </div>
                <div class="task-unix">
                  <span> {{ item.task.computeSourceShow }}，</span>
                  <span> {{ item.task.accCardTypeShow }} </span>
                </div>
                <div class="task-unix">
                  <span> {{ $t('cloudbrainObj.runDuration') }}：</span>
                  <span> {{ item.task.formatted_duration }}</span>
                </div>
                <div class="task-unix">
                  <span> {{ $t('cloudbrainObj.createTime') }}： </span>
                  <span> {{ dateFormat(item.task.created_unix) }} </span>
                </div>
              </div>
            </div>
          </div>
          <div class="task-actions">
            <span v-if="isTaskRunning(item.task)" class="action-btn stop" @click.stop="handleStopTask(item.task)">{{
              $t('cloudbrainObj.stop') }}</span>
            <span v-else-if="isTaskStopping(item.task)" class="action-btn stop disabled">{{ $t('cloudbrainObj.stop')
            }}</span>
            <span v-else class="action-btn delete" :class="{ disabled: !canDeleteTask(item.task) }"
              @click.stop="canDeleteTask(item.task) && handleDeleteTask(item.task)">{{ $t('cloudbrainObj.delete')
              }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from 'dayjs';
import { getAiTask, stopAiTask, deleteAiTask } from '~/apis/modules/cloudbrain';
import { getListValueWithKey } from '~/utils';
import { ACC_CARD_TYPE, JOB_TYPE } from '~/const';

export default {
  name: 'TaskLimitDialog',
  props: {
    visible: { type: Boolean, default: false },
    title: { type: String, default: '' },
    desc: { type: String, default: '' },
    taskList: { type: Array, default: () => [] },
  },
  data() {
    return {
      statusPollingTimer: null,
    };
  },
  watch: {
    taskList: {
      handler() {
        this.checkAndStartPolling();
      },
      deep: true,
    },
    visible(newVal) {
      if (!newVal) {
        this.stopStatusPolling();
      }
    },
  },
  methods: {
    // 将 jobTypeShow 转为行数组：数组直接返回，字符串按2字分行
    getJobTypeLines(jobTypeShow) {
      if (Array.isArray(jobTypeShow)) return jobTypeShow;
      if (!jobTypeShow) return [];
      // 中文按2字分行，英文按空格/单词分行
      const lines = [];
      // 检测是否包含中文
      if (/[\u4e00-\u9fff]/.test(jobTypeShow)) {
        for (let i = 0; i < jobTypeShow.length; i += 2) {
          lines.push(jobTypeShow.slice(i, i + 2));
        }
      } else {
        lines.push(jobTypeShow);
      }
      return lines;
    },
    dateFormat(unix) {
      return dayjs(unix * 1000).format('YYYY-MM-DD HH:mm:ss');
    },

    // 判断任务是否在运行或等待中（可操作停止的状态）
    isTaskRunning(task) {
      return task.status === 'WAITING' || task.status === 'RUNNING';
    },

    // 判断任务是否正在停止中
    isTaskStopping(task) {
      return task.status === 'STOPPING';
    },

    // 判断任务是否可删除（只有终态才可以删除）
    canDeleteTask(task) {
      const terminalStatuses = ['STOPPED', 'SUCCEEDED', 'FAILED', 'CREATE_FAILED'];
      return terminalStatuses.includes(task.status);
    },

    // 处理停止任务
    handleStopTask(task) {
      stopAiTask({
        repoOwnerName: task.owner_name,
        repoName: task.repo_name,
        id: task.id,
      }).then(res => {
        res = res.data;
        if (res.code === 0) {
          this.$message.success(this.$t('notebook.stopSuccess'));
          this.$emit('stopSuccess', task);
        } else {
          this.$message.error(res.msg);
        }
      }).catch(() => {
        this.$message.error(this.$t('notebook.stopFailed'));
      });
    },

    // 处理删除任务
    handleDeleteTask(task) {
      deleteAiTask({
        repoOwnerName: task.owner_name,
        repoName: task.repo_name,
        id: task.id,
      }).then(res => {
        res = res.data;
        if (res.code === 0) {
          this.$message.success(this.$t('notebook.deleteSuccess'));
          this.$emit('deleteSuccess', task);
        } else {
          this.$message.error(res.msg);
        }
      }).catch(() => {
        this.$message.error(this.$t('notebook.deleteFailed'));
      });
    },

    // 检查是否有非终态任务并开启轮询（WAITING、RUNNING、STOPPING 需要轮询）
    checkAndStartPolling() {
      const pollingStatuses = ['WAITING', 'RUNNING', 'STOPPING'];
      const hasPollingTask = this.taskList.some(item => pollingStatuses.includes(item.task.status));
      if (hasPollingTask) {
        this.startStatusPolling();
      } else {
        this.stopStatusPolling();
      }
    },

    // 开启状态轮询
    startStatusPolling() {
      if (this.statusPollingTimer) return;
      this.statusPollingTimer = setInterval(() => {
        this.pollTaskStatus();
      }, 10000);
    },

    // 停止状态轮询
    stopStatusPolling() {
      if (this.statusPollingTimer) {
        clearInterval(this.statusPollingTimer);
        this.statusPollingTimer = null;
      }
    },

    // 轮询任务状态（逐个查询非终态任务）
    pollTaskStatus() {
      const pollingStatuses = ['WAITING', 'RUNNING', 'STOPPING'];
      const pollingTasks = this.taskList.filter(item => pollingStatuses.includes(item.task.status));
      if (!pollingTasks.length) {
        this.stopStatusPolling();
        this.$emit('allTasksTerminal');
        return;
      }
      let pendingCount = pollingTasks.length;
      pollingTasks.forEach(item => {
        getAiTask({ id: item.task.id }).then(res => {
          res = res.data;
          if (res.code === 0 && res.data?.task) {
            const newTask = res.data.task;
            newTask.computeSourceShow = newTask.compute_source == 'GPU' ? 'CPU/GPU' : newTask.compute_source;
            newTask.accCardTypeShow = getListValueWithKey(ACC_CARD_TYPE, newTask.acc_card_type);
            if (newTask.job_type === 'ComfyuiExperience') {
              newTask.jobTypeShow = ['Comfy', 'UI']
            } else if (newTask.job_type === 'EVAL') {
              newTask.jobTypeShow = this.$t('modelSquare.modelEvaluate');
            } else if (newTask.job_type === 'FINETUNE') {
              const evalText = this.$t('modelSquare.sftFinetune');
              newTask.jobTypeShow = [evalText.slice(0, 3), evalText.slice(3)]
            } else {
              newTask.jobTypeShow = getListValueWithKey(JOB_TYPE, newTask.job_type);
            }
            this.$emit('taskStatusUpdate', newTask);
          }
        }).catch(err => {
          console.log('pollTaskStatus error', err);
        }).finally(() => {
          pendingCount--;
          if (pendingCount === 0) {
            this.$nextTick(() => {
              const stillPolling = this.taskList.some(item => pollingStatuses.includes(item.task.status));
              if (!stillPolling) {
                this.stopStatusPolling();
                this.$emit('allTasksTerminal');
              }
            });
          }
        });
      });
    },

    handleClose() {
      this.stopStatusPolling();
      this.$emit('close');
    },
  },
  beforeDestroy() {
    this.stopStatusPolling();
  },
};
</script>

<style scoped lang="less">
.task-limit-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 2000;
  display: flex;
  align-items: center;
  justify-content: center;
}

.task-limit-dialog {
  max-width: 900px;
  width: 100%;
  max-height: 80vh;
  background-color: #fff;
  box-shadow: 0 4px 24px rgba(0, 0, 0, 0.15);

  .dialog-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    min-height: 80px;
    background: linear-gradient(86.11deg, rgba(9, 60, 241, 1) -0.14%, rgba(72, 134, 255, 1) 93.93%);
    padding: 12px 30px;

    .dialog-tip-wrap {
      display: flex;
      align-items: center;

      .icon-wrap {
        width: 50px;
        height: 50px;
        border-radius: 10px;
        background-color: rgba(255, 255, 255, 0.2);
        display: flex;
        align-items: center;
        justify-content: center;
        flex-shrink: 0;

        i {
          font-size: 24px;
          color: #fff;
        }
      }

      .tip-wrap {
        margin-left: 14px;
        min-height: 50px;

        .tip-title {
          color: rgba(255, 255, 255, 1);
          font-size: 16px;
          font-weight: 700;
          margin: 0;
          margin-bottom: 6px;
        }

        .tip-desc {
          color: rgba(255, 245, 0, 1);
          line-height: 14px;
          margin: 0;
        }
      }
    }

    .close-wrap {
      width: 30px;
      height: 30px;
      border-radius: 5px;
      background-color: rgba(255, 255, 255, 0.2);
      display: flex;
      align-items: center;
      justify-content: center;
      flex-shrink: 0;

      i {
        cursor: pointer;
        font-size: 20px;
        color: #fff;

        &:hover {
          color: #409eff;
        }
      }
    }
  }

  .task-list {
    max-height: 700px;
    overflow-y: auto;
    padding: 20px 24px;

    .task-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 12px;
      background-color: rgba(248, 251, 255, 1);
      border-radius: 6px;
      margin-top: 10px;

      .task-info-wrap {
        display: flex;
        flex: 1;
        min-width: 0;

        .cb-job {
          // width: 40px;
          height: 40px;
          border-radius: 6px;
          color: rgba(249, 249, 249, 1);
          font-size: 14px;
          line-height: 20px;
          flex-shrink: 0;
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
        }

        .task-info {
          margin-left: 12px;

          .task-name {
            height: 20px;
            line-height: 20px;
            color: rgba(16, 16, 16, 1);
          }

          .task-meta {
            display: flex;
            align-items: center;
            gap: 8px;
            font-size: 12px;
            color: rgba(16, 16, 16, 0.5);

            .task-status {
              color: rgba(80, 85, 89, 1);
              display: flex;
              align-items: center;
              padding: 0 10px;
              background-color: rgba(233, 233, 235, 1);
              border-radius: 10px;

              .dot {
                width: 6px;
                height: 6px;
                border-radius: 100%;
                margin-right: 6px;
                background-color: rgba(80, 85, 89, 1);
              }

              &.status-RUNNING,
              &.status-WAITING {
                color: rgba(39, 177, 72, 1);
                background-color: rgba(217, 255, 226, 1);

                .dot {
                  background-color: rgba(39, 177, 72, 1);
                }
              }

              &.status-STOPPING {
                color: rgba(242, 113, 28, 1);
                background-color: rgba(255, 237, 218, 1);

                .dot {
                  background-color: rgba(242, 113, 28, 1);
                }
              }
            }
          }
        }
      }

      .task-actions {
        flex-shrink: 0;
        margin-left: 16px;

        .action-btn {
          cursor: pointer;
          font-size: 12px;
          padding: 0 10px;
          border-radius: 4px;
          display: inline-block;

          &.stop {
            color: rgba(0, 86, 255, 1);
            background-color: rgba(236, 245, 255, 1);
            border: 1px solid rgba(0, 86, 255, 1);

            &.disabled {
              color: rgba(200, 200, 200, 1);
              background-color: rgba(245, 245, 245, 1);
              border-color: rgba(200, 200, 200, 1);
              cursor: not-allowed;
            }
          }

          &.delete {
            color: rgba(244, 56, 56, 1);
            background-color: rgba(255, 255, 255, 1);
            border: 1px solid rgba(244, 56, 56, 1);

            &.disabled {
              color: rgba(200, 200, 200, 1);
              background-color: rgba(245, 245, 245, 1);
              border-color: rgba(200, 200, 200, 1);
              cursor: not-allowed;
            }
          }
        }
      }
    }
  }
}

@media screen and (max-width: 767px) {
  .task-meta {
    .task-unix {
      display: none;
    }
  }
}
</style>
