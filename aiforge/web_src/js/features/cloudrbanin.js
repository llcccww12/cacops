export default async function initCloudrain() {
  function paddingZeros(str, len) {
    str = str.toString();
    if (str.length < len) {
      str = new Array(len - str.length).fill('0').join('') + str;
    }
    return str;
  }

  function timeFormat(date) {
    return `${date.getFullYear()}-${paddingZeros(date.getMonth() + 1, 2)}-${paddingZeros(date.getDate(), 2)} ${paddingZeros(date.getHours(), 2)}:${paddingZeros(date.getMinutes(), 2)}:${paddingZeros(date.getSeconds(), 2)}`;
  }

  let debug_button = $(".cloudbrain_debug").data("debug");
  let debug_again_button = $(".cloudbrain_debug").data("debug-again");
  let timeid
  let timeidShow
  if (document.getElementsByClassName('job-status').length!==0) {
    timeid = window.setInterval(loadJobStatus, 15000);
    timeidShow = window.setInterval(loadShowJobStatus, 15000);
  }
  $(document).ready(loadJobStatus);
  $(document).ready(loadShowJobStatus);
  function loadJobStatus() {
    let flagStopInterval = 0
    $(".job-status").each((index, job) => {
      const ID = job.dataset.jobid;
      const cloudbrainID = job.dataset.cloudbrainid;
      if (!ID) return;
      const repoPath = job.dataset.repopath;
      const dataMigrate = job.dataset.datamigrate
      const centerPend = job.dataset.centerpend
      // const computeResource = job.dataset.resource
      const versionname = job.dataset.version;
      const bootfile = job.dataset.bootfile;
      const status_text = $(`#${ID}-text`).text();
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
      if (finalState.includes(status_text)) {
        flagStopInterval++
        if (flagStopInterval === $(".job-status").length) { 
          clearInterval(timeid)
        }
        return;
      }
      // const diffResource = computeResource == "NPU" ? 'modelarts/notebook' : 'cloudbrain'
      $.get(
        `/api/v1/repos/${repoPath}/${ID}?version_name=${versionname}&id=${cloudbrainID}`,
        (data) => {
          const ID = data.ID || data.JobID;
          const status = data.JobStatus;
          const duration = data.JobDuration;
          const aiCenter = data.AiCenter || '--'
          const detailStatus = data.DetailedStatus
          $("#duration-" + ID).text(duration);
          data.AiCenter != undefined && $("#cluster-" + ID).text(aiCenter);
          data.AiCenter != undefined && $("#" + versionname + "-ai_center").text(data.AiCenter);
          if ( (detailStatus === 'dataMigrating' || detailStatus === 'centerPending') && status==='WAITING') {
            if (document.getElementById(`${ID}-icon-detail`)) {
              document.getElementById(`${ID}-icon-detail`).remove()
            }
            $("#" + ID + "-text").parent().append(`<i id="${ID}-icon-detail" class="${detailStatus}" style="vertical-align: middle;" title="${detailStatus === 'dataMigrating'? dataMigrate:centerPend}"></i>`)
          } else { 
            if (document.getElementById(`${ID}-icon-detail`)) {
              document.getElementById(`${ID}-icon-detail`).remove()
            }
          }
          if (status != status_text) {
            $("#" + ID + "-icon")
              .removeClass()
              .addClass(status);
            $("#" + ID + "-text").text(status);
            finalState.includes(status) &&
              $("#" + ID + "-stop")
                .removeClass("blue")
                .addClass("disabled");
          }
          if (status === "RUNNING") {
            $("#ai-debug-" + ID)
              .removeClass("disabled")
              .addClass("blue")
              .text(debug_button)
              .css("margin", "0 1rem");
            $("#model-image-" + ID)
              .removeClass("disabled")
              .addClass("blue");
            $("#ai-debug-infer-" + ID)
            .removeClass("disabled")
            .addClass("blue")
          }
          if (status !== "RUNNING") {
            // $('#model-debug-'+ID).removeClass('blue')
            // $('#model-debug-'+ID).addClass('disabled')
            $("#model-image-" + ID)
              .removeClass("blue")
              .addClass("disabled");
          }
          if (
            ["CREATING", "STOPPING", "WAITING", "STARTING"].includes(status)
          ) {
            $("#ai-debug-" + ID)
              .removeClass("blue")
              .addClass("disabled");
            $("#ai-debug-infer-" + ID)
            .removeClass("blue")
            .addClass("disabled")
          }
          if (
            [
              "STOPPED",
              "FAILED",
              "START_FAILED",
              "CREATE_FAILED",
              "SUCCEEDED",
            ].includes(status)
          ) {
            $("#ai-debug-infer-" + ID)
            .removeClass("blue")
            .addClass("disabled")
            if (!bootfile) {
              $("#ai-debug-" + ID)
                .removeClass("disabled")
                .addClass("blue")
                .text(debug_again_button)
                .css("margin", "0");
            } else {
              $("#ai-debug-" + ID).remove()
            }

          }
          if (["RUNNING", "WAITING"].includes(status)) {
            $("#ai-stop-" + ID)
              .removeClass("disabled")
              .addClass("blue");
          }
          if (
            [
              "PREPARING",
              "CONNECTING",
              "CREATING",
              "STOPPING",
              "STARTING",
              "STOPPED",
              "FAILED",
              "START_FAILED",
              "SUCCEEDED",
              "COMPLETED",
              "CREATE_FAILED",
              "CREATED_FAILED",
            ].includes(status)
          ) {
            $("#ai-stop-" + ID)
              .removeClass("blue")
              .addClass("disabled");
          }

          if (
            [
              "STOPPED",
              "FAILED",
              "START_FAILED",
              "KILLED",
              "COMPLETED",
              "SUCCEEDED",
              "CREATE_FAILED",
              "CREATED_FAILED",
            ].includes(status)
          ) {
            $("#ai-delete-" + ID)
              .removeClass("disabled")
              .addClass("blue");
          } else {
            $("#ai-delete-" + ID)
              .removeClass("blue")
              .addClass("disabled");
          }
        }
      ).fail(function (err) {
        console.log(err);
      });
    });
  }

  function loadShowJobStatus() {
    $(".ui.accordion.border-according").each((index, job) => {
      const jobID = job.dataset.jobid;
      if (!jobID) return;
      const repoPath = job.dataset.repopath;
      const versionname = job.dataset.version;
      // ['IMAGE_FAILED','SUBMIT_FAILED','DELETE_FAILED','KILLED','COMPLETED','FAILED','CANCELED','LOST','START_FAILED']
      // if (job.textContent.trim() == 'IMAGE_FAILED' || job.textContent.trim() == 'SUBMIT_FAILED' || job.textContent.trim() == 'DELETE_FAILED'
      //     || job.textContent.trim() == 'KILLED' || job.textContent.trim() == 'COMPLETED' || job.textContent.trim() == 'FAILED'
      //     || job.textContent.trim() == 'CANCELED' || job.textContent.trim() == 'LOST') {
      //     return
      // }
      let status = $(`#${versionname}-status-span`).text().trim();

      if (
        [
          "IMAGE_FAILED",
          "SUBMIT_FAILED",
          "DELETE_FAILED",
          "KILLED",
          "COMPLETED",
          "FAILED",
          "CANCELED",
          "LOST",
          "START_FAILED",
          "SUCCEEDED",
          "STOPPED",
          "CREATE_FAILED",
          "CREATED_FAILED",
        ].includes(status)
      ) {
        clearInterval(timeidShow)
        return;
      }
      let stopArray = [
        "PREPARING",
        "CONNECTING",
        "KILLED",
        "FAILED",
        "START_FAILED",
        "KILLING",
        "COMPLETED",
        "SUCCEEDED",
        "CREATE_FAILED",
        "CREATED_FAILED",
        "STOPPED",
      ];
      let deleteArray = [
        "KILLED",
        "FAILED",
        "START_FAILED",
        "COMPLETED",
        "SUCCEEDED",
        "CREATE_FAILED",
        "CREATED_FAILED",
        "STOPPED",
      ];
      $.get(
        `/api/v1/repos/${repoPath}/${jobID}?version_name=${versionname}`,
        (data) => {
          $(`#${versionname}-duration-span`).text(data.JobDuration);
          $(`#${versionname}-status-span span`).text(data.JobStatus);
          $(`#${versionname}-status-span i`).attr("class", data.JobStatus == 'TESTING' ? 'RUNNING' : data.JobStatus);
          // detail status and duration
          data.StartTime !== undefined && data.StartTime > 0 && $("#" + versionname + "-startTime").text(timeFormat(new Date(data.StartTime * 1000)));
          $("#" + versionname + "-duration").text(data.JobDuration);
          $("#" + versionname + "-status").text(data.JobStatus);
          data.AiCenter != undefined && $("#" + versionname + "-ai_center").text(data.AiCenter);

          if (stopArray.includes(data.JobStatus)) {
            $("#" + versionname + "-stop").addClass("disabled");
          }
          if (deleteArray.includes(data.JobStatus)) {
            $(`#${versionname}-delete`).removeClass("disabled");
            $(`#${versionname}-delete`).addClass("blue");
          }
          if (data.JobStatus === "COMPLETED") {
            $("#" + versionname + "-create-model")
              .removeClass("disabled")
              .addClass("blue");
          }
        }
      ).fail(function (err) {
        console.log(err);
      });
    });
  }

  function assertDelete(obj, versionName, repoPath, cloudbrainID) {
    if (obj.style.color == "rgb(204, 204, 204)") {
      return;
    } else {
      const delId = obj.parentNode.getAttribute('id');
      let flag = 1;
      $(".ui.basic.modal")
        .modal({
          onDeny: function () {
            flag = false;
          },
          onApprove: function () {
            if (!versionName) {
              document.getElementById(delId).submit();
            } else {
              deleteVersion(versionName, repoPath, cloudbrainID);
            }
            flag = true;
          },
          onHidden: function () {
            if (flag == false) {
              $(".alert")
                .html(i18n.canceled_operation)
                .removeClass("alert-success")
                .addClass("alert-danger")
                .show()
                .delay(1500)
                .fadeOut();
            }
          },
        })
        .modal("show");
    }
  }
  function deleteVersion(versionName, repoPath, cloudbrainID) {
    const url = `/api/v1/repos/${repoPath}`;
    $.post(url, { version_name: versionName, id: cloudbrainID }, (data) => {
      if (data.StatusOK === 0 || data.Code === 0) {
        location.reload();
      } else {
        $(".alert")
          .html(data.Message)
          .removeClass("alert-success")
          .addClass("alert-danger")
          .show()
          .delay(1500)
          .fadeOut();
      }
    }).fail(function (err) {
      console.log(err);
    });
  }
  $(".ui.basic.ai_delete").click(function () {
    const repoPath = this.dataset.repopath;
    const versionName = this.dataset.version;
    const cloudbrainID = this.dataset.cloudbrainid;
    if (repoPath && versionName) {
      assertDelete(this, versionName, repoPath, cloudbrainID);
    } else {
      assertDelete(this);
    }
  });
  function stopDebug(ID, stopUrl, bootFile) {
    $.ajax({
      type: "POST",
      url: stopUrl,
      data: $("#stopForm-" + ID).serialize(),
      success: function (res) {
        if (res.result_code === "0") {
          $("#" + ID + "-icon")
            .removeClass()
            .addClass(res.status);
          $("#" + ID + "-text").text(res.status);
          if (res.status === "STOPPED") {
            if (!bootFile) {
              $("#ai-debug-" + ID)
                .removeClass("disabled")
                .addClass("blue")
                .text(debug_again_button)
                .css("margin", "0");
            } else {
              $("#ai-debug-" + ID).remove()
            }
            $("#ai-image-" + ID)
              .removeClass("blue")
              .addClass("disabled");
            $("#ai-model-debug-" + ID)
              .removeClass("blue")
              .addClass("disabled");
            $("#ai-delete-" + ID)
              .removeClass("disabled")
              .addClass("blue");
            $("#ai-stop-" + ID)
              .removeClass("blue")
              .addClass("disabled");
          } else {
            $("#ai-debug-" + ID)
              .removeClass("blue")
              .addClass("disabled");
            $("#ai-debug-infer-" + ID)
              .removeClass("blue")
              .addClass("disabled");
            $("#ai-stop-" + ID)
              .removeClass("blue")
              .addClass("disabled");
          }
        } else {
          $(".alert")
            .html(res.error_msg)
            .removeClass("alert-success")
            .addClass("alert-danger")
            .show()
            .delay(2000)
            .fadeOut();
        }
      },
      error: function (res) {
        console.log(res);
      },
    });
  }
  $(".ui.basic.ai_stop").click(function () {
    const ID = this.dataset.jobid;
    const repoPath = this.dataset.repopath;
    const bootFile = this.dataset.bootfile
    stopDebug(ID, repoPath, bootFile);
  });

  function stopVersion(version_name, ID, repoPath, cloudbrainID) {
    const url = `/api/v1/repos/${repoPath}/${ID}/stop_version`;
    $.post(url, { version_name: version_name, id: cloudbrainID }, (data) => {
      if (data.StatusOK === 0) {
        $("#ai-stop-" + ID).removeClass("blue");
        $("#ai-stop-" + ID).addClass("disabled");
        refreshStatus(version_name, ID, repoPath, cloudbrainID);
      } else { 
        $(".alert")
          .html(data.error_msg)            
          .removeClass("alert-success")
          .addClass("alert-danger")
          .show()
          .delay(2000)
          .fadeOut();
      }
    }).fail(function (err) {
      console.log(err);
    });
  }
  function refreshStatus(version_name, ID, repoPath, cloudbrainID) {
    const url = `/api/v1/repos/${repoPath}/${ID}?version_name${version_name}&id=${cloudbrainID}`;
    $.get(url, (data) => {
      $(`#${ID}-icon`).attr("class", data.JobStatus);
      // detail status and duration
      $(`#${ID}-text`).text(data.JobStatus);
      if (
        [
          "STOPPED",
          "FAILED",
          "START_FAILED",
          "KILLED",
          "COMPLETED",
          "SUCCEEDED",
          "CREATE_FAILED",
          "CREATED_FAILED",
        ].includes(data.JobStatus)
      ) {
        $("#ai-delete-" + ID)
          .removeClass("disabled")
          .addClass("blue");
      }
      const detailStatus = data.DetailedStatus
      const dataMigrate = $(`${ID}`).data('datamigrate')
      const centerPend = $(`${ID}`).data('centerpend')
      if ((detailStatus === 'dataMigrating' || detailStatus === 'centerPending') && data.JobStatus==='WAITING') {
        if (document.getElementById(`${ID}-icon-detail`)) {
          document.getElementById(`${ID}-icon-detail`).remove()
        }
        $("#" + ID + "-text").parent().append(`<i id="${ID}-icon-detail" class="${detailStatus}" style="vertical-align: middle;" title="${detailStatus === 'dataMigrating'? dataMigrate:centerPend}"></i>`)
      } else { 
        if (document.getElementById(`${ID}-icon-detail`)) {
          document.getElementById(`${ID}-icon-detail`).remove()
        }
      }
    }).fail(function (err) {
      console.log(err);
    });
  }
  $(".ui.basic.ai_stop_version").click(function () {
    const ID = this.dataset.jobid;
    const cloudbrainID = this.dataset.cloudbrainid;
    const repoPath = this.dataset.repopath;
    const versionName = this.dataset.version;
    stopVersion(versionName, ID, repoPath, cloudbrainID);
  });
  function getModelInfo(repoPath, modelId, modelName, versionName, jobName) {
    $.get(
      `/api/v1/repos${repoPath}/modelmanage/query_model_byId?id=${modelId}`,
      (data) => {
        if (data && data.id == modelId) {
          location.href = `${repoPath}/modelmanage/model_readme_tmpl?name=${data.name}`;
        } else {
          $(`#${jobName}`).popup("toggle");
        }
      }
    ).catch(err => {
      console.log(err);
      $(`#${jobName}`).popup("toggle");
    });
  }
  $(".goto_modelmanage").click(function () {
    let repoPath = this.dataset.repopath;
    const modelId = this.dataset.modelid;
    const modelName = this.dataset.modelname;
    const versionName = this.dataset.version;
    const jobName = this.dataset.jobname;
    const modelRepoOwnerName = this.dataset.modelrepoownername;
    const modelRepoName = this.dataset.modelreponame;
    if (modelRepoOwnerName && modelRepoName) {
      repoPath = `/${modelRepoOwnerName}/${modelRepoName}`
    }
    getModelInfo(repoPath, modelId, modelName, versionName, jobName);
  });
  function debugAgain(ID, debugUrl, redirect_to) {
    if ($("#" + ID + "-text").text() === "RUNNING") {
      window.open(debugUrl + "debug");
    } else {
      $.ajax({
        type: "POST",
        url: debugUrl + "restart?redirect_to=" + redirect_to,
        data: $("#debugAgainForm-" + ID).serialize(),
        success: function (res) {
          if (res["WechatRedirectUrl"]) {
            window.location.href = res["WechatRedirectUrl"];
          } else if (res.result_code === "0") {
            if (res.id !== ID) {
              location.reload();
            } else {
              $("#" + ID + "-icon")
                .removeClass()
                .addClass(res.status);
              $("#" + ID + "-text").text(res.status);
              $("#ai-debug-" + ID)
                .removeClass("blue")
                .addClass("disabled");
              $("#ai-delete-" + ID)
                .removeClass("blue")
                .addClass("disabled");
              $("#ai-debug-" + ID)
                .text(debug_button)
                .css("margin", "0 1rem");
            }
          } else if (res.result_code == "2") {
            $(".ui.modal.debug-again-alert").modal("show");
          } else {
            $(".alert")
              .html(res.error_msg)
              .removeClass("alert-success")
              .addClass("alert-danger")
              .show()
              .delay(3000)
              .fadeOut();
          }
        },
        error: function (res) {
          console.log(res);
        },
      });
    }
  }
  $(".ui.basic.ai_debug").click(function () {
    const ID = this.dataset.jobid;
    const repoPath = this.dataset.repopath;
    const redirect_to = this.dataset.linkpath;
    debugAgain(ID, repoPath, redirect_to);
  });
}

function userSearchControll() {
  if ($("#userCloud").length === 0) {
    return;
  }
  const params = new URLSearchParams(window.location.search);
  let cluster;
  if ($(".cloudbrain_debug").length === 1) {
    if (!params.get("cluster")) {
      cluster = $(".cloudbrain_debug").data("all-cluster");
    } else {
      if (params.get("cluster") === "resource_cluster_c2net") {
        cluster = $(".cloudbrain_debug").data("cluster-c2net");
      } else {
        cluster = $(".cloudbrain_debug").data("cluster-openi");
      }
    }
  }
  let jobType;
  if ($(".cloudbrain_debug").length === 1) {
    if (!params.get("jobType")) {
      jobType = $(".cloudbrain_debug").data("allTask");
    } else {
      if (params.get("jobType") === "DEBUG") {
        jobType = $(".cloudbrain_debug").data("debug-task");
      } else if (params.get("jobType") === "TRAIN") {
        jobType = $(".cloudbrain_debug").data("train-task");
      } else if (params.get("jobType") === "INFERENCE") {
        jobType = $(".cloudbrain_debug").data("inference-task");
      } else if (params.get("jobType") === "ONLINEINFERENCE") {
        jobType = $(".cloudbrain_debug").data("inferonline-task");
      } else if (params.get("jobType") === "HPC") {
        jobType = $(".cloudbrain_debug").data("supercompute-task");
      } else {
        jobType = $(".cloudbrain_debug").data("benchmark-task");
      }
    }
  }
  let aiCenter = !params.get("aiCenter")
    ? $(".cloudbrain_debug").data("all-aiCenter")
    : params.get("aiCenter");
  let listType = !params.get("listType")
    ? $(".cloudbrain_debug").data("all-compute")
    : params.get("listType");
  let jobStatus = !params.get("jobStatus")
    ? $(".cloudbrain_debug").data("all-status")
    : params.get("jobStatus").toUpperCase();
  const dropdownValueArray = [cluster, aiCenter, jobType, listType, jobStatus];
  $("#userCloud .default.text ").each(function (index, e) {
    index != 1 && $(e).text(dropdownValueArray[index]);
  });
}

function AdaminSearchControll() {
  if ($("#adminCloud").length === 0) {
    return;
  }
  const params = new URLSearchParams(window.location.search);
  let cluster;
  if ($(".cloudbrain_debug").length === 1) {
    if (!params.get("cluster")) {
      cluster = $(".cloudbrain_debug").data("all-cluster");
    } else {
      if (params.get("cluster") === "resource_cluster_c2net") {
        cluster = $(".cloudbrain_debug").data("cluster-c2net");
      } else {
        cluster = $(".cloudbrain_debug").data("cluster-openi");
      }
    }
  }
  let aiCenter = !params.get("aiCenter")
    ? $(".cloudbrain_debug").data("all-aiCenter")
    : params.get("aiCenter");
  let jobType = !params.get("jobType")
    ? $(".cloudbrain_debug").data("all-task")
    : params.get("jobType");
  let listType = !params.get("listType")
    ? $(".cloudbrain_debug").data("all-compute")
    : params.get("listType");
  let jobStatus = !params.get("jobStatus")
    ? $(".cloudbrain_debug").data("all-status")
    : params.get("jobStatus").toUpperCase();
  const dropdownValueArray = [cluster, aiCenter, jobType, listType, jobStatus];
  $("#adminCloud .default.text ").each(function (index, e) {
    index != 1 && $(e).text(dropdownValueArray[index]);
  });
}
userSearchControll();
AdaminSearchControll();
$(".message .close").on("click", function () {
  $(this).closest(".message").transition("fade");
});

/* MLU Debug List Operations */
$(".ui.basic.ai_mlu_debug").click(function () {
  const debug_button = $(".cloudbrain_debug").data("debug");
  const taskID = this.dataset.jobid;
  const linkPath = this.dataset.linkpath;
  const repoOwnerName = linkPath.split('/')[1];
  const repoName = linkPath.split('/')[2];
  if ($("#" + taskID + "-text").text() === "RUNNING") {
    $.ajax({
      type: 'get',
      url: `/api/v1/${repoOwnerName}/${repoName}/ai_task/debug_url?id=${taskID}`,
      data: { _csrf: csrf, },
      success: function (res) {
        if (res.code == 0) {
          window.open(res.data.url);
        } else {
          $(".alert")
            .html(res.msg)
            .removeClass("alert-success")
            .addClass("alert-danger")
            .show()
            .delay(3000)
            .fadeOut();
        }
      },
      error: function (err) {
        console.log(err);
      },
    });
  } else {
    $.ajax({
      type: 'post',
      url: `/api/v1/${repoOwnerName}/${repoName}/ai_task/restart?id=${taskID}`,
      data: { _csrf: csrf, },
      success: function (res) {
        if (res.code == 0) {
          if (res.data.id != taskID) {
            window.location.reload();
          } else {
            $("#" + taskID + "-icon")
              .removeClass()
              .addClass(res.data.status);
            $("#" + taskID + "-text").text(res.data.status);
            $("#ai-debug-" + taskID)
              .removeClass("blue")
              .addClass("disabled");
            $("#ai-delete-" + taskID)
              .removeClass("blue")
              .addClass("disabled");
            $("#ai-debug-" + taskID)
              .text(debug_button)
              .css("margin", "0 1rem");
          }
        } else {
          $(".alert")
            .html(res.msg)
            .removeClass("alert-success")
            .addClass("alert-danger")
            .show()
            .delay(3000)
            .fadeOut();
        }
      },
      error: function (err) {
        console.log(err);
      },
    });
  }
});

$(".ui.basic.ai_mlu_stop").click(function () {
  const debug_again_button = $(".cloudbrain_debug").data("debug-again");
  const ID = this.dataset.jobid;
  const bootFile = this.dataset.bootfile;
  const linkPath = this.dataset.linkpath;
  const repoOwnerName = linkPath.split('/')[1];
  const repoName = linkPath.split('/')[2];
  $.ajax({
    type: "post",
    url: `/api/v1/${repoOwnerName}/${repoName}/ai_task/stop?id=${ID}`,
    data: { _csrf: csrf, },
    success: function (res) {
      if (res.code == "0") {
        $("#" + ID + "-icon")
          .removeClass()
          .addClass(res.data.status);
        $("#" + ID + "-text").text(res.data.status);
        if (res.data.status === "STOPPED") {
          if (!bootFile) {
            $("#ai-debug-" + ID)
              .removeClass("disabled")
              .addClass("blue")
              .text(debug_again_button)
              .css("margin", "0");
          } else {
            $("#ai-debug-" + ID).remove()
          }
          $("#ai-image-" + ID)
            .removeClass("blue")
            .addClass("disabled");
          $("#ai-model-debug-" + ID)
            .removeClass("blue")
            .addClass("disabled");
          $("#ai-delete-" + ID)
            .removeClass("disabled")
            .addClass("blue");
          $("#ai-stop-" + ID)
            .removeClass("blue")
            .addClass("disabled");
        } else {
          $("#ai-debug-" + ID)
            .removeClass("blue")
            .addClass("disabled");
          $("#ai-stop-" + ID)
            .removeClass("blue")
            .addClass("disabled");
        }
      } else {
        $(".alert")
          .html(res.msg)
          .removeClass("alert-success")
          .addClass("alert-danger")
          .show()
          .delay(2000)
          .fadeOut();
      }
    },
    error: function (res) {
      console.log(res);
    },
  });
});

$(".ui.basic.ai_mlu_delete").click(function () {
  const ID = this.dataset.jobid;
  const linkPath = this.dataset.linkpath;
  const repoOwnerName = linkPath.split('/')[1];
  const repoName = linkPath.split('/')[2];
  let submitFlag = false;
  $(".ui.basic.modal")
    .modal({
      onDeny: function () {
        submitFlag = false;
      },
      onApprove: function () {
        $.ajax({
          type: "post",
          url: `/api/v1/${repoOwnerName}/${repoName}/ai_task/del?id=${ID}`,
          data: { _csrf: csrf, },
          success: function (res) {
            if (res.code == 0) {
              window.location.reload();
            } else {
              $(".alert")
                .html(res.msg)
                .removeClass("alert-success")
                .addClass("alert-danger")
                .show()
                .delay(1500)
                .fadeOut();
            }
          },
          error: function (res) {
            console.log(res);
          },
        });
        submitFlag = true;
      },
      onHidden: function () {
        if (submitFlag == false) {
          $(".alert")
            .html(i18n.canceled_operation)
            .removeClass("alert-success")
            .addClass("alert-danger")
            .show()
            .delay(1500)
            .fadeOut();
        }
      },
    })
    .modal("show");
});