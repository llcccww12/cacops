import service from '../service';
import Qs from 'qs';
let _csrf  = window.config ? window.config.csrf : ''

// 查询创建任务时所需信息
export const getAiTaskPrepareInfo = (params) => {
  return service({
    url: `/api/v1/ai_task/creation/required`,
    method: 'get',
    params: {
      job_type: params.jobType,
      compute_source: params.computeSource,
      cluster_type: params.clusterType,
      app_name: params.appName,
    },
  });
}

// 由资源规格查询可用镜像
// params: repoOwnerName, repoName, job_type,compute_source,cluster_type,spec_id
export const getAiTaskImgesBySpec = (params) => {
  return service({
    url: `/api/v1/ai_task/creation/image_by_spec`,
    method: 'get',
    params: {
      job_type: params.jobType,
      compute_source: params.computeSource,
      cluster_type: params.clusterType,
      spec_id: params.spec,
      has_internet: params.hasInternet,
      visualize_required: params.visualizeRequired,
    },
  });
}

// 是否能够创建AI任务
export const getCanCreateTask = (params) => {
  return service({
    url: `/${params.repoOwnerName}/${params.repoName}/ai_task/can_create`,
    method: 'get',
    params: {},
    data: {}
  });
}

// 获取AI任务模板编辑页地址
export const getTmplEditAdress = (params) => {
  return service({
    url: `/api/v1/${params.repoOwnerName}/${params.repoName}/ai_task/template/get_edit_address${params.branchName ? '/branch/' + params.branchName : ''}`,
    method: 'get',
    params: {},
    data: {}
  });
}

// 创建AI任务
export const createAiTask = (data) => {
  return service({
    url: `/api/v1/${data.repoOwnerName}/${data.repoName}/ai_task/create`,
    method: 'post',
    params: {},
    data: { ...data, }
  });
}

// 查询AI任务列表
// page,
// job_type-DEBUG|TRAIN|INFERENCE|BENCHMARK,
// compute_source-GPU|NPU|GCU|MLU
export const getAiTaskList = (params) => {
  return service({
    url: `/api/v1/${params.repoOwnerName}/${params.repoName}/ai_task/list`,
    method: 'get',
    params: { ...params },
  });
}

// 查看任务详情
export const getAiTask = (params) => {
  return service({
    url: `/api/v1/ai_task`,
    method: 'get',
    params: { id: params.id },
  });
}

// 查询任务最新状态
export const getAiTaskBrief = (params) => {
  return service({
    url: `/api/v1/ai_task/brief`,
    method: 'get',
    params: { id: params.id },
  });
}

// 获取调试页面地址
export const getAiTaskDebugUrl = (params) => {
  return service({
    url: `/api/v1/ai_task/debug_url`,
    method: 'get',
    params: { id: params.id, file: params.file },
  });
}

// 获取可视化页面地址
export const getAiTaskVisualizeUrl = (params) => {
  return service({
    url: `/api/v1/ai_task/visualize_url`,
    method: 'get',
    params: { id: params.id },
  });
}

// 获取自定义路径页面地址
export const getAiEndpointUrl = (params) => {
  return service({
    url: `/api/v1/ai_task/self_endpoint_url`,
    method: 'get',
    params: { id: params.id, file: params.file },
  });
}

// 再次调试
export const getAiTaskRestart = (data) => {
  return service({
    url: `/api/v1/ai_task/restart?id=${data.id}`,
    method: 'post',
    params: {},
    data: {},
  });
}

// 停止任务
export const stopAiTask = (data) => {
  return service({
    url: `/api/v1/ai_task/stop?id=${data.id}`,
    method: 'post',
    params: {},
    data: {},
  });
}

// 删除任务
export const deleteAiTask = (data) => {
  return service({
    url: `/api/v1/ai_task/del?id=${data.id}`,
    method: 'post',
    params: {},
    data: {},
  });
}
// 批量删除任务
export const deleteMulAiTask = (data) => {
  return service({
    url: '/api/v1/ai_task/batch_del',
    method: 'post',
    params: {},
    data: data,
  });
}
// 查询任务运行简况
export const getAiTaskOperationProfile = (params) => {
  return service({
    url: `/api/v1/ai_task/operation_profile`,
    method: 'get',
    params: { id: params.id },
  });
}
// 查询任务loss曲线
export const getAiTaskLoss = (params) => {
  return service({
    url: `/api/v1/ai_task/loss`,
    method: 'get',
    params: { id: params.id },
  });
}
// 查询任务日志
// params: id, base_line, lines, order-asc|desc,node_id,file_name
export const getAiTaskLogs = (params) => {
  return service({
    url: `/api/v1/ai_task/log`,
    method: 'get',
    params: {
      id: params.id,
      base_line: params.base_line,
      lines: params.lines,
      order: params.order,
      node_id: params.node_id,
      file_name: params.log_file_name,
    },
  });
}

// 下载任务日志
export const getAiTaskLogsDownloadUrl = (params) => {
  return `/api/v1/ai_task/log/download?id=${params.id}&node_id=${params.node_id}&file_name=${params.log_file_name}&_csrf=${_csrf}`
}

// 查询任务资源占用情况
// params: id,node_id,log_file_name
export const getAiTaskResourceUseage = (params) => {
  return service({
    url: `/api/v1/ai_task/resource_usage`,
    method: 'get',
    params: {
      id: params.id,
      node_id: params.node_id,
      file_name: params.log_file_name,
    },
  });
}

// 查询任务结果下载
// params: id, parent_dir
export const getAiTaskOutputResult = (params) => {
  return service({
    url: `/api/v1/ai_task/output`,
    method: 'get',
    params: {
      id: params.id,
      parent_dir: params.parent_dir,
    },
  });
}

// 查询任务结果下载(所有)
// params: id,suffix-.zip|.tar.gz
export const getAiTaskOutputResultAll = (params) => {
  return service({
    url: `/api/v1/ai_task/output/all`,
    method: 'get',
    params: { id: params.id, suffix: params.suffix },
  });
}

// 下载训练任务的某一个结果文件
// params-id,file_name,parent_dir
export const getDownLoadAiTaskResultFileUrl = (params) => {
  return `/api/v1/ai_task/output/download?id=${params.id}&file_name=${params.file_name}&parent_dir=${params.parent_dir}&_csrf=${_csrf}`
  
}

// 下载训练任务的所有结果文件链接
// params-id
export const getDownLoadAiTaskResultFileAllUrl = (params) => {
  return `/api/v1/ai_task/output/download/all?id=${params.id}&_csrf=${_csrf}`
  
}

// 结果下载重新获取
export const setAiTaskOutputReschedule = (data) => {
  const paramsObj = { id: data.id };
  return service({
    url: `/api/v1/ai_task/output/reschedule`,
    method: 'post',
    data: { id: data.id, },
    params: paramsObj,
  });
}

// 获取训练任务的节点信息
export const getAiTaskNodeInfo = (params) => {
  return service({
    url: `/api/v1/ai_task/node_info`,
    method: 'get',
    params: { id: params.id, },
  });
}

// 导出训练任务的输出至模型
// trainTaskCreate-true,cloudbrain_id(jobId,versionName),name,version,engine,engine_name,modelSelectedFile,label,isPrivate,description
export const setAiTaskResultToModel = (data) => {
  return service({
    url: `/${data.repoOwnerName}/${data.repoName}/modelmanage/create_new_model`,
    method: 'post',
    params: {},
    data: Qs.stringify(data),
  });
}
// 导出训练任务的输出至模型
// trainTaskCreate-true,cloudbrain_id(jobId,versionName),name,version,engine,engine_name,modelSelectedFile,label,isPrivate,description
export const setAiTaskResultToModelApi = (data) => {
  return service({
    url: `/api/v1/repos/${data.repoOwnerName}/${data.repoName}/modelmanage/create_new_model`,
    method: 'post',
    params: {},
    data: Qs.stringify(data),
  });
}
/* 导出训练任务的输出至数据集相关 */
// 获取项目仓库的数据集信息
export const getRepoDatasetInfo = (params) => {
  return service({
    url: `/${params.repoOwnerName}/${params.repoName}/datasets/model/getcurrentdataset`,
    method: 'get',
    params: {},
  });
}

// 获取AI任务结果导出到数据集的导出进度信息
export const getAiTaskExportDatasetProgress = (params) => {
  return service({
    url: `/api/v1/dataset/get_export_process`,
    method: 'get',
    params: params,
  });
}
// AI任务结果中导入数据集提交
// cloudbrain_id(jobId,versionName),datasetId,modelSelectedFile,type,description,
export const setAiTaskExportDataset = (params, data) => {
  return service({
    url: `/api/v1/ai_task/export2dataset`,
    method: 'post',
    params: params,
    data: data,
  });
}


// AI任务结果中导入模型提交
// cloudbrain_id(jobId,versionName),datasetId,modelSelectedFile,type,description,
export const setAiTaskExportModel = (params, data) => {
  return service({
    url: `/api/v1/ai_task/export2aimodel`,
    method: 'post',
    params: params,
    data: data,
  });
}

// 获取AI任务结果导出到模型的导出进度信息
export const getAiTaskExportModelProgress = (params) => {
  return service({
    url: '/api/v1/aimodel/get_export_process',
    method: 'get',
    params: { id: params.id, },
  });
}
// AI任务结果中导入数据集提交
// cloudbrain_id(jobId,versionName),datasetId,modelSelectedFile,type,description,
export const setAiTaskExportDataset1 = (data) => {
  return service({
    url: `/${data.repoOwnerName}/${data.repoName}/datasets/model/export_exist_dataset`,
    method: 'post',
    params: {},
    data: Qs.stringify(data),
  });
}

// 我的工作台-云脑任务列表
// job_type,job_status,ai_center,cluster,compute_source,q,page,pageSize
export const getMyAiTasks = (params) => {
  return service({
    url: `/api/v1/ai_task/my_list`,
    method: 'get',
    params: { ...params },
  });
}

// 管理后台-云脑任务列表
// job_type,job_status,ai_center,cluster,compute_source,q,page,pageSize
export const getAdminAiTasks = (params) => {
  return service({
    url: `/api/v1/admin/ai_task/list`,
    method: 'get',
    params: { ...params },
  });
}

/* Aim可视化相关 */ 
// 获取Aim功能权限状态
export const getAimRight = () => {
  return service({
    url: `/api/v1/monitor/aim_right`,
    method: 'get',
    params: {},
  });
}

// 获取Aim跳转地址
// job_name-["jobname1","jobname2"]
export const getAimUrl = (data) => {
  return service({
    url: `/api/v1/monitor/aim_url`,
    method: 'post',
    data: {
      job_name: data.job_name,
    },
  });
}


// 查询任务评测总览
export const getAiEvalResult = (params) => {
  return service({
    url: `/api/v1/ai_task/eval_result`,
    method: 'get',
    params: { id: params.id },
  });
}

// 查询任务评测详情
export const getAiEvalDetailResult = (data) => {
  return service({
    url: `/api/v1/ai_task/eval_detail_result`,
    method: 'post',
    params: {id: data.task_id},
    data: data
  });
}