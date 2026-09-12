import { getAiTaskBrief } from '~/apis/modules/cloudbrain';
import { timeSinceUnix } from '~/utils';

export class CloudBrainTools {
  constructor() {
    this.intervalTimer = null;
    this.refreshIntervalTime = 1000 * 8;
    this.list = [];
  }
  initRefreshData(data) {
    this.stop();
    this.list = data;
    this.refreshStatus();
    this.intervalTimer = setInterval(() => {
      this.refreshStatus();
    }, this.refreshIntervalTime);
  }
  refreshStatus() {
    for (let i = 0, iLen = this.list.length; i < iLen; i++) {
      const taskInfo = this.list[i];
      const task = taskInfo.task;
      task.createdFromNow = timeSinceUnix(task.created_unix, Date.now() / 1000);
      const finalState = [
        "STOPPED",
        "CREATE_FAILED",
        "CREATED_FAILED",
        "UNAVAILABLE",
        "DELETED",
        "RESIZE_FAILED",
        "SUCCEEDED",
        "IMAGE_FAILED",
        "SUBMIT_FAILED",
        "DELETE_FAILED",
        "KILLED",
        "COMPLETED",
        "FAILED",
        "CANCELED",
        "LOST",
        "START_FAILED",
        "SUBMIT_MODEL_FAILED",
        "DEPLOY_SERVICE_FAILED",
        "CHECK_FAILED",
      ];
      if (finalState.includes(task.status)) {
        continue;
      }
      const self = this;
      (function (taskInfo) {
        const task = taskInfo.task;
        getAiTaskBrief({
          repoOwnerName: task.repoOwnerName,
          repoName: task.repoName,
          id: task.id,
        }).then(res => {
          res = res.data;
          if (res.code == 0) {
            const data = res.data;
            Object.assign(task, data);
            task.createdFromNow = timeSinceUnix(task.created_unix, Date.now() / 1000);
            self.checkRunningLeftTime(task);
            self.checkOperation(task);
          }
        }).catch(err => {
          console.log(err);
        });
      })(taskInfo);
    }
  }
  stop() {
    clearInterval(this.intervalTimer)
  }
  checkOperation(task) {
    if (task.status == 'RUNNING') {
      task.canDebug = true;
      task.canSaveImage = true;
    } else {
      task.canDebug = false;
      task.canSaveImage = false;
    }
    if (task.job_type == 'TRAIN' && task.status == 'RUNNING' && task.visualize_required) {
      task.canVisualize = true;
    } else {
      task.canVisualize = false;
    }
    if (task.aim_required && ["RUNNING", "STOPPED", "FAILED", "SUCCEEDED", "COMPLETED"].includes(task.status)) {
      task.canAim = true;
    } else {
      task.canAim = false;
    }
    if (["PREPARING", "CONNECTING", "CREATING", "STOPPING", "WAITING", "STARTING"].includes(task.status)) {
      task.canDebug = false;
    }
    if (["STOPPED", "FAILED", "START_FAILED", "CREATE_FAILED", "SUCCEEDED"].includes(task.status)) {
      if (task.is_file_notebook) {
        task.canReDebug = false;
      } else {
        task.canReDebug = true;
      }
    } else {
      task.canReDebug = false;
    }
    if (["INIT", "RUNNING", "WAITING"].includes(task.status)) {
      task.canStop = true;
    } else {
      task.canStop = false;
    }
    if (["STOPPED", "FAILED", "START_FAILED", "KILLED", "COMPLETED", "SUCCEEDED", "CREATE_FAILED", "CREATED_FAILED"].includes(task.status)) {
      task.canDelete = true;
    } else {
      task.canDelete = false;
    }
    task.hasMore = true;
    task.canSaveTmpl = ['DEBUG', 'TRAIN', 'ONLINEINFERENCE', 'GENERAL', 'HPC'].includes(task.job_type);
    if ((['GPU', 'NPU', 'GCU', 'MLU', 'ILUVATAR-GPGPU', 'METAX-GPGPU', 'BIREN-GPU'].includes(task.compute_source) || task.ai_center_code === 'sugon-ai') && ['DEBUG', 'ONLINEINFERENCE', 'ComfyuiExperience'].includes(task.job_type)) {
      task.hasMore = true;
      if (task.canDebug) {
        task.canSaveImage = true;
        if (task.job_type === 'ComfyuiExperience') {
          task.saveImageUrl = `/notebook/${task.id}/commit_image?type=comfyui`;
        } else {
          task.saveImageUrl = `/notebook/${task.id}/commit_image`;
        }
      } else {
        task.canSaveImage = false;
      }
      if (task.can_modify && task.cluster == 'OpenI' && !['PREPARING', 'CONNECTING'].includes(task.status)) {
        task.canDownloadModel = true;
        task.downloadModelUrl = `/${task.repoOwnerName}/${task.repoName}/cloudbrain/${task.id}/models`;
      } else {
        task.canDownloadModel = false;
      }
      if (task.is_file_notebook) {
        task.hasDebugMore = false;
        task.canSaveImage = false;
        task.canDownloadModel = false;
      }
    } else {
      task.hasDebugMore = false;
      task.canSaveImage = false;
      task.canDownloadModel = false;
    }
    task.canModify = true;
    if (task.is_fine_tune_task || task.is_file_notebook) {
      task.canModify = false;
    }
    if ((task.job_type == 'TRAIN' || task.job_type == 'FINETUNE') && task.can_download) {
      task.canExportOutput = true;
    }
  }
  checkRunningLeftTime(task) {
    const timeLimit = task.time_limit > -1 ? (task.time_limit > 0 ? task.time_limit : task.default_time_limit) : 0;
    if (task.job_type == 'DEBUG' && task.status == 'RUNNING' && timeLimit > 0) {
      const hmsList = task.formatted_duration.split(':').map((item) => Number(item));
      const durationSecond = hmsList[0] * 60 * 60 + hmsList[1] * 60 + hmsList[2];
      const allSecond = timeLimit * 60 * 60;
      task.runningLeftTime = Math.max(1, Math.floor((allSecond - durationSecond) / 60));
    } else {
      task.runningLeftTime = undefined;
    }
  }
  // 优化后的getAiJobLink方法
  getAiJobLink(taskInfo) {
    
    return '';
  }
}
