import service from "../service";
import Qs from 'qs';



//创建数据集
export const createDatasets = (data) => {
  return service({
    url: '/api/v1/dataset/create',
    method: "post",
    params: {},
    data: data,
  });
};

//获取管理后台数据集列表
export const getAdminDatasets = (params,type='dataset') => {
  return service({
    url: `/api/v1/admin/${type}/list`,
    method: "get",
    params: params
  });
};
// 推荐数据集
export const putRecommendDatasets = (params,type='dataset') => {
  return service({
    url: `/api/v1/admin/${type}/recommend`,
    method: "put",
    params: params
  });
};
// 取消推荐数据集
export const delRecommendDatasets = (params,type='dataset') => {
  return service({
    url: `/api/v1/admin/${type}/recommend`,
    method: "delete",
    params: params
  });
};

//获取数据集列表
export const getDatasets = ({ url, ...params }) => {
  return service({
    url: `/api/v1/dataset/list${url}`,
    method: "get",
    params: params
  });
};
// 获取数据集详情
export const getDatasetsDetail = (params) => {
  return service({
    url: '/api/v1/dataset',
    method: "get",
    params: params
  });
};
// 删除数据集
export const delDataset = (params,type='dataset') => {
  return service({
    url: `/api/v1/${type}`,
    method: "delete",
    params: params
  });
};

// 删除单个数据集
export const delSingleDataset = (params,type) => {
  return service({
    url: `/api/v1/${type}/file`,
    method: "delete",
    params: params
  });
};

// 编辑数据集基本信息
export const editDatasetsDetail = (params,data,type) => {
  return service({
    url: `/api/v1/${type}/edit`,
    method: "post",
    params: params,
    data: data
  });
};

//获取readme
export const getDatasetReadMe = (params, type) => {
  return service({
    url: `/api/v1/${type}/readme`,
    method: "get",
    params: params
  });
};

//创建readme
export const postDatasetReadMe = (params, data, type) => {
  return service({
    url: `/api/v1/${type}/readme`,
    method: "post",
    params: params,
    data: data
  });
};
//获取可创建数据集的用户/组织列表
export const getAvailableUsers = (params,type='dataset') => {
  return service({
    url: `/api/v1/${type}/create/available_users`,
    method: "get",
    params: params,
  });
};
//查询用户表
export const getUsers = (params) => {
  return service({
    url: '/api/v1/users/search',
    method: "get",
    params: params,
  });
};

//查询团队表
export const getTeams = (params) => {
  return service({
    url: `/api/v1/orgs/${params.org}/teams/search`,
    method: "get",
    params: params,
  });
};
//数据集添加协作者
export const postCollaborate = (data) => {
  return service({
    url: '/api/v1/access/collaboration',
    method: "post",
    params: data,
    data: {}
  });
};
//数据集删除协作者
export const deleteCollaborate = (data) => {
  return service({
    url: '/api/v1/access/collaboration',
    method: "delete",
    params: data,
    
  });
};
//数据集添加团队
export const postTeam = (data,type) => {
  return service({
    url: `/api/v1/access/collaboration/${type}/team`,
    method: "post",
    params: data,
    data: {}
  });
};
//数据集添加团队
export const deleteTeam = (data,type) => {
  return service({
    url: `/api/v1/access/collaboration/${type}/team/delete`,
    method: "post",
    params: data,
    data: {}
  });
};

//修改协作者权限
export const modifyCollaborateAcess = (data) => {
  return service({
    url: '/api/v1/access/collaboration/access_mode',
    method: "post",
    params: data,
    data: {}
  });
};
//数据集已有协作者列表包括组织
export const getCollaborate = (params) => {
  return service({
    url: '/api/v1/access/collaboration',
    method: "get",
    params: params,
  });
};

//数据集预览 dataset_id,parent_dir,file_name
export const getFilePreview = (params,type) => {
  return service({
    url: `/api/v1/${type}/preview`,
    method: "get",
    params: params,
  });
};

export const putDatasetStar = (url) => {
  return service({
    url: url,
    method: "put"
  });
};

export const getModelFile = (params) => {
  return service({
    url: '',
    params: params,
    method:"get"
  })
}

export const exportExistDataset = (data) => {
  return service({
    url: '',
    data: data,
    method:"post"
  })
}

/* 选择数据集组件相关 */
// 获取当前仓库的数据集
// params - username, reponame, q, page, type
export const getCurrentRepoDataset = (params) => {
  return service({
    url: `/${params.userName}/${params.repoName}/datasets/current_repo_m`,
    method: 'get',
    params: {
      q: params.q || '',
      page: params.page || 1,
      type: params.type == undefined ? '-1' : params.type,
    },
  });
}

// 获取我上传的数据集
// params - username, reponame, q, page, type
export const getMyUploadedDataset = (params) => {
  return service({
    url: `/${params.userName}/${params.repoName}/datasets/my_datasets_m`,
    method: 'get',
    params: {
      q: params.q || '',
      page: params.page || 1,
      type: params.type == undefined ? '-1' : params.type,
    },
  });
}

// 获取公开数据集
// params - username, reponame, q, page, type
export const getPulicDataset = (params) => {
  return service({
    url: `/${params.userName}/${params.repoName}/datasets/public_datasets_m`,
    method: 'get',
    params: {
      q: params.q || '',
      page: params.page || 1,
      type: params.type == undefined ? '-1' : params.type,
    },
  });
}

// 获取我收藏的数据集
// params - username, reponame, q, page, type
export const getMyFavoriteDataset = (params) => {
  return service({
    url: `/${params.userName}/${params.repoName}/datasets/my_favorite_m`,
    method: 'get',
    params: {
      q: params.q || '',
      page: params.page || 1,
      type: params.type == undefined ? '-1' : params.type,
    },
  });
}

/* 数据集文件上传 */
// 上传文件1: 获取文件chunks信息
// params: { md5, file_name, subject_id,  subject_type: 1, }
// return: uuid, uploaded, chunks, subjectId, fileName
export const getChunks = (params) => {
  return service({
    url: `/api/v1/upload/get_chunks`,
    method: 'get',
    params: params,
    data: {},
  });
};

// 上传文件2: 上传新文件
// params: { total_chunk_counts, md5, size, file_type, subject_type, file_name, subject_id }
// return: uuid
export const getNewMultipart = (params) => {
  return service({
    url: `/api/v1/upload/new_multipart`,
    method: 'get',
    params,
    data: {},
  });
};

// 上传文件3: 获取分片上传地址
// params: { uuid, size, chunk_number }
// return: url
export const getMultipartUrl = (params) => {
  return service({
    url: `/api/v1/upload/get_multipart_url`,
    method: 'get',
    params,
    data: {},
  });
};

// 上传文件4: 完成上传后
// data: { uuid }
// return "code": 0
export const setCompleteMultipart = (data) => {
  return service({
    url: `/api/v1/upload/complete_multipart`,
    method: 'post',
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    params: {},
    data: Qs.stringify(data),
  });
};


// 数据集广场上方推荐数据集信息
// 数据集广场左侧分类、研究方向、证书等标签的源信息
export const getPromoteDataset = (params) => {
  return service({
    url: `/api/v1/dync_config`,
    method: 'get',
    params,
    data: {},
  });
};
//收藏数据集 
// params: {dataset_id:数据集id}
export const setFavorite = (data,type) => {
  return service({
    url: `/api/v1/${type}/collect`,
    method: "put",
    params: data
  });
}
//取消收藏数据集 
// params: {dataset_id:数据集id}
export const unsetFavorite = (data,type) => {
  return service({
    url: `/api/v1/${type}/collect`,
    method: "delete",
    params: data
  });
}


//获取数据集文件列表
// params: {dataset_id:数据集id, parent_dir:文件路径, page_size: 10, marker: ''}
export const getFileList = (params,type) => {
  return service({
    url: `/api/v1/${type}/files`,
    method: "get",
    params: params
  });
}


//弹窗sdk下载代码
export const getFileSdkCode = (params) => {
  return service({
    url: `/api/v1/${params.type}/sdk_download_code`,
    method: "get",
    params: params
  });
}

//弹窗sdk下载代码
export const getCheckoldDataset = (params) => {
  return service({
    url: '/api/v1/dataset/check_old',
    method: "get",
    params: params
  });
}

//获取个人信息页面中的数据集列表
export const getProfileDataset = (params) => {
  return service({
    url: '/api/v1/dataset/list/accessible',
    method: "get",
    params: params
  });
}

//获取个人信息页面中的数据集列表（未登录态使用，仅返回公开数据集）
export const getProfileDatasetPublic = (params) => {
  return service({
    url: '/api/v1/dataset/list/public',
    method: "get",
    params: params
  });
}