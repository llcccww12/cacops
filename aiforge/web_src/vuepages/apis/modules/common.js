import service from '../service';
import Qs from 'qs';

// 获取静态文件内容
export const getStaticFile = (filePathName) => {
  return service({
    url: `${filePathName}`,
    method: 'get',
    params: {},
  });
}

// 获取promote配置数据
export const getPromoteData = (filePathName) => {
  return service({
    url: '/dashboard/invitation',
    method: 'get',
    params: {
      filename: filePathName
    },
  });
}

// 获取模型体验运行的云脑任务列表
export const getModelExperience = () => {
  return service({
    url: '/extension/modelexperience/queryall',
    method: 'get',
    params: {},
  });
}

// 获取markdown渲染结果
export const getMarkdownHtml = (str, mode) => {
  return service({
    url: '/api/v1/markdown',
    method: 'post',
    parmas: {},
    data: {
      mode: mode || 'gfm',
      text: str,
    },
  });
}

// 获取个人积分信息
// return {pointAccount:{id,account_code,balance,total_earned,total_consumed,status,version,created_unix,updated_unix },cloudBrainPaySwitch}
export const getPointAccountInfo = () => {
  return service({
    url: '/api/v1/user/point_account',
    method: 'get',
    params: {},
  });
}

// 由模型名称列表或数据集名称列表或运行参数列表获取SDK的代码使用方式
// dataset_name-string[], datasetNameList
// pretrain_model_name -string[] modelNameList
// param_key -string[] param_key
// job_type,compute_source,cluster_type
export const getSDKCode = (params) => {
  const datasetNames = params.dataset_name || [];
  const modelNames = params.pretrain_model_name || [];
  const parameterKeys = params.param_key || [];
  const searchParams = new URLSearchParams();
  datasetNames.forEach(name => {
    searchParams.append('dataset_name', name);
  });
  modelNames.forEach(name => {
    searchParams.append('pretrain_model_name', name);
  });
  parameterKeys.forEach(name => {
    searchParams.append('param_key', name);
  });
  return service({
    url: `/api/v1/ai_task/generate_sdk_code?${searchParams.toString()}`,
    method: 'get',
    params: {
      job_type: params.job_type,
      compute_source: params.compute_source,
      cluster_type: params.cluster_type,
      visualize_required: params.visualize_required
    },
    data: {},
  });
}

// 获取单个模型文件的SDK code
export const getModelFileSDKCode = (params) => {
  return service({
    url: `/api/v1/${params.owner}/${params.repo}/sdk/generate_model_download_code`,
    method: 'get',
    params: { model_name: params.name, model_file_name: params.filename },
    data: {},
  });
}

// 查询智算列表
export const getAiCenterList = () => {
  return service({
    url: `/resources/queue/centers`,
    method: 'get',
    params: {}
  });
};

// 查询所有资源队列名称列表
export const getResQueueCode = (params) => {
  return service({
    url: `/explore/card_request/resources/queue/codes`,
    method: 'get',
    params,
  });
};

// common form post
export const commonFormPost = (url, data) => {
  return service({
    url: url,
    method: 'post',
    params: {},
    data: Qs.stringify(data),
  });
}

// 获取用户最近AI任务的项目仓信息
export const getUserLastestAiTaskRepoInfo = () => {
  return service({
    url: `/api/v1/user/get_latest_cloudbrain_repo`,
    method: 'get',
    params: {},
  });
};

// 用户有权限的仓库列表
// uid,sort-updated,order-desc,asc,type-cloudbrain,model,dataset
export const getUserRepoList = (params) => {
  return service({
    url: `/api/v1/repos/search_for_ai_task`,
    method: 'get',
    params: {
      uid: params.uid,
      sort: params.sort || 'updated',
      order: params.order || 'desc',
      type: params.type || 'cloudbrain'
    },
  });
};

// 新建代码仓
// auto_init-true,default_branch-master,name,private-true
export const createRepo = (data) => {
  return service({
    url: `/api/v1/user/repos`,
    method: 'post',
    data: {
      auto_init: data.auto_init == undefined ? true : data.auto_init,
      default_branch: data.default_branch || 'master',
      name: data.name,
      private: data.private == undefined ? true : data.private,
    },
    params: {},
  });
};

// fork代码仓
// owner,repo,repo_name
export const forkRepo = (data) => {
  return service({
    url: `/api/v1/repos/${data.owner}/${data.repo}/forks`,
    method: 'post',
    data: {
      owner: data.owner,
      repo: data.repo,
      repo_name: data.repo_name,
    },
    params: {},
  });
};


// 获取comyui路由
export const getComfyuiUrl = (params) => {
  return service({
    url: '/api/v1/comfyui/experience/online_url',
    method: 'get',
    params: params
  });
}

// 获取 Mlops 权限
export const getMlopsRight = () => {
  return service({
    url: '/api/v1/monitor/mlops_right',
    method: 'get',
    params: {}
  });
}

// 热力图
export const getHeatMap = (userName) => {
  return service({
    url: `/api/v1/users/${userName}/heatmap`,
    method: 'get',
    params: {}
  });
}

// 用户action 

export const getAction = () => {
  return service({
    url: '/api/v1/platform/action',
    method: 'get',
    params: {}
  });
}