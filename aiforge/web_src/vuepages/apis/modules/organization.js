import service from '../service';
import Qs from 'qs';
// repos
// 获取组织精选项目列表 orgName
export const getOrgSelectedRepoList = (params) => {
  return service({
    url: `/org/${params.orgName}/org_card_repo`,
    method: 'get',
    params: {},
  });
}

// 获取组织精选项目可设置的列表 orgName
export const getOrgSelectedRepoSetList = (params) => {
  return service({
    url: `/org/${params.orgName}/org_tag/repo_list`,
    method: 'get',
    params: { tagId: 1 },
  });
}

// 设置组织精选项目 orgName,list
export const setOrgSelectedRepo = (data) => {
  return service({
    url: `/org/${data.orgName}/org_tag/repo_submit`,
    method: 'post',
    params: { tagId: 1 },
    data: { repoList: data.list },
  });
}

// 获取组织项目列表 orgName
export const getOrgRepoList = (params) => {
  return service({
    url: `/org/${params.orgName}/org_list_repo`,
    method: 'get',
    params: {
      q: params.q,
      filter: params.label,
      sort: params.sort,
      page: params.page,
      pageSize: params.pageSize,
    }
  });
}

// model
// 获取组织精选模型列表 orgName
export const getOrgSelectedModelList = (params) => {
  return service({
    url: `/api/v1/orgs/${params.orgName}/org_tag/aimodel/list`,
    method: 'get',
    params: {},
  });
}

// 获取组织精选模型可设置的列表 orgName
export const getOrgSelectedModelSetList = (params) => {
  return service({
    url: `/api/v1/orgs/${params.orgName}/org_tag/aimodel/available`,
    method: 'get',
    params: {},
  });
}

// 设置组织精选模型
export const setOrgSelectedModel = (data) => {
  return service({
    url: `/api/v1/orgs/${data.orgName}/org_tag/aimodel/submit`,
    method: 'post',
    params: {},
    data: { modelids: data.list.join(',') },
  });
}

// 获取组织模型列表
export const getOrgModelList = ({owner_name, ...getParams }) => {
  return service({
    url: `/api/v1/orgs/${owner_name}/aimodel/list`,
    method: 'get',
    params: getParams,
  });
}

// dataset
// 获取组织精选数据集列表
export const getOrgSelectedDatasetList = (params) => {
  return service({
    url: `/api/v1/orgs/${params.orgName}/org_tag/dataset/list`,
    method: 'get',
    params: {},
  });
}

// 获取组织精选数据集可设置的列表 orgName
export const getOrgSelectedDatasetSetList = (params) => {
  return service({
    url: `/api/v1/orgs/${params.orgName}/org_tag/dataset/available`,
    method: 'get',
    params: {},
  });
}

// 设置组织精选数据集 orgName, list
export const setOrgSelectedDataset = (data) => {
  return service({
    url: `/api/v1/orgs/${data.orgName}/org_tag/dataset/submit`,
    method: 'post',
    params: {},
    data: { datasetids: data.list.join(',') },
  });
}

// 获取组织数据集列表 orgName
export const getOrgDatasetList = ({owner_name, ...getParams }) => {
  return service({
    url: `/api/v1/orgs/${owner_name}/dataset/list`,
    method: 'get',
    params: getParams,
  });
}
// 获取组织数据集列表 orgName
export const getOrgStorageDatasetList = ({owner_name, ...getParams },type) => {
  return service({
    url: `/api/v1/orgs/${owner_name}/storage/${type}`,
    method: 'get',
    params: getParams,
  });
}
//查询组织数据集标签列表
export const getOrgLabel = (params,type='dataset') => {
  return service({
    url: `/api/v1/orgs/${params.orgName}/${type}/tags`,
    method: "get",
    params: {}
  });
}

export const getOrgStorageSummary = (orgName,params) => {
  return service({
      url: `/api/v1/orgs/${orgName}/storage/summary`,
      method: 'get',
      params: params,
  });
}

export const listAiforgeOrgRole = (params) => {
  return service({
    url: `/api/v1/admin/role/org/list`,
    method: 'get',
    params: params,

  });
}

export const listRightOrgUser = (params) => {
  return service({
    url: `/api/v1/admin/role/list_right_org`,
    method: 'get',
    params: params,
  });
}

export const setAiforgeRoleToOrgUser = (data) => {
  return service({
    url: `/api/v1/admin/role/set_to_org`,
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    method: 'post',
    params: { },
    data: Qs.stringify(data),
  });
}

//提供工作台用户组织
export const listMyOrgUser = (params) => {
  return service({
    url: `/api/v1/platform/org`,
    method: 'get',
    params: params,
  });
}
