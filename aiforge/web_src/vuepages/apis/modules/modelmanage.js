import service from "../service";
import Qs from 'qs';

// 保存本地模型
export const saveLocalModel = (data) => {
  return service({
    url: `${data.repo}/modelmanage/create_local_model`,
    method: 'post',
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    params: {},
    data: Qs.stringify(data),
  });
};
//创建模型
export const createModels = (data) => {
  return service({
    url: '/api/v1/aimodel/create',
    method: "post",
    params: {},
    data: data,
  });
};
// 修改模型
// data: {id,type,name,version,engine,label,description:}
export const modifyModel = (data) => {
  return service({
    url: `${data.repo}/modelmanage/modify_model`,
    method: 'put',
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    params: {},
    data: Qs.stringify(data),
  });
};

export const modifyModelStatus = (data) => {
  return service({
    url: `${data.repo}/modelmanage/modify_model_status`,
    method: 'put',
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    params: {},
    data: Qs.stringify(data),
  });
};

// 求模型信息
export const getModelInfoByName = (params) => {
  return service({
    url: '/api/v1/aimodel',
    method: 'get',
    params
  });
};

//或许训练任务列表
export const getTrainJobList = (params) => {
  return service({
    url: `${params.repo}/modelmanage/query_train_job?repoId=${params.repoId}`,
    method: 'get',
    params: {},
    data: {},
  });
};

// 求模型中文件列表
// params {repo, ID, parentDir}
export const getModelFiles = (params) => {
  return service({
    url: `${params.repo}/modelmanage/query_onelevel_modelfile`,
    method: 'get',
    params,
    data: {},
  });
};

// 删除模型文件
// params {repo, id, fileName}
export const deleteModelFile = (params) => {
  return service({
    url: `${params.repo}/modelmanage/delete_model_file`,
    method: 'delete',
    params,
    data: {},
  });
};

/* 文件上传相关 */
// 上传文件1: 获取文件chunks信息
// params: { md5, type: 0-CPU/GPU,1-NPU, file_name, scene: 'model', modeluuid }
// return: uploadID, uuid, uploaded, chunks, attachID, modeluuid, modelName, fileName
export const getChunks = (params) => {
  return service({
    url: `/attachments/model/get_chunks`,
    method: 'get',
    params,
    data: {},
  });
};

// 上传文件2: 上传新文件
// params: { totalChunkCounts, md5, size, fileType, type, file_name, scene=model, modeluuid=xxxx }
// return: uploadID, uuid
export const getNewMultipart = (params) => {
  return service({
    url: `/attachments/model/new_multipart`,
    method: 'get',
    params,
    data: {},
  });
};

// 上传文件3: 获取分片上传地址
// params: { uuid, uploadID, size, chunkNumber, type, file_name, scene=model }
// return: url
export const getMultipartUrl = (params) => {
  return service({
    url: `/attachments/model/get_multipart_url`,
    method: 'get',
    params,
    data: {},
  });
};

// 上传文件4: 完成上传后
// data: { uuid, uploadID, size, type, file_name, dataset_id, description, scene=model, modeluuid=xxxx }
export const setCompleteMultipart = (data) => {
  return service({
    url: `/attachments/model/complete_multipart`,
    method: 'post',
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    params: {},
    data: Qs.stringify(data),
  });
};

// markdown预览
// data: { userName, repoName, context, text }
export const getMarkdownPreview = (data) => {
  return service({
    url: `/api/v1/repos/${data.repoOwnerName}/${data.repoName}/markdown`,
    method: 'post',
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    params: {},
    data: Qs.stringify({
      mode: 'gfm',
      context: '',
      text: data.text,
    }),
  });
}


// 获取模型演化图谱数据
export const getModelEvolutionMap = (params) => {
  return service({
    url: `/api/v1/aimodel/evolution`,
    method: 'get',
    params: { aimodel_id: params.aimodel_id },
  });
}

// 获取平台模型许可证列表
export const getModelLicenseList = () => {
  return service({
    url: `/dashboard/invitation`,
    method: 'get',
    params: { filename: 'model/license.json' },
  });
}

// 模型迁移新建
// data: { hf_repo_id, version, engine, label, license, description, isPrivate }
export const setModelMigrate = (data) => {
  return service({
    url: `/api/v1/hf_model/new`,
    method: 'post',
    params: {},
    data: { ...data },
  });
}

// 模型迁移查询文件迁移状态
// params: { hf_repo_id }
export const getModelMigrateStatus = (params) => {
  return service({
    url: `/api/v1/hf_model/status`,
    method: 'get',
    params: { ...params },
  });
}

// 模型迁移重试
// data: { hf_repo_id }
export const setModelMigrateRetry = (data) => {
  return service({
    url: `/api/v1/hf_model/retry`,
    method: 'post',
    params: {},
    data: { ...data },
  });
}

// 模型迁移查询迁移者信息
// params: { model_id }
export const getModelMigrateUserInfo = (params) => {
  return service({
    url: `/api/v1/hf_model/get_transfer_user`,
    method: 'get',
    params: { ...params },
  });
}

// 迁移模型同步文件对比列表
// hf_repo_id
export const getModelMigrateUpdateInfo = (params) => {
  return service({
    url: `/api/v1/hf_model/fetch_update`,
    method: 'get',
    params,
    data: {},
  });
};

// 迁移模型发起同步
// hf_repo_id
export const setModelMigrateUpdate = (data) => {
  return service({
    url: `/api/v1/hf_model/update`,
    method: 'post',
    params: {},
    data: { ...data },
  });
};

//获取个人信息页面中的模型列表
export const getProfileModel = (params) => {
  return service({
    url: '/api/v1/aimodel/list/accessible',
    method: "get",
    params: params
  });
}

//获取个人信息页面中的模型列表（未登录态使用，仅返回公开模型）
export const getProfileModelPublic = (params) => {
  return service({
    url: '/api/v1/aimodel/list/public',
    method: "get",
    params: params
  });
}
