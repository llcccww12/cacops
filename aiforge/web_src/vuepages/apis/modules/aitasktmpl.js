import service from '../service';
import { lang } from '~/langs'

/* 查询计算任务模板列表 */
// type 
// - public 公开资源-公开模板
// - recommend 公开资源-平台推荐模板
// - created 我的工作台-我创建的
// - collected 我的工作台-我收藏的
// - admin 管理后台-计算任务模板列表
// - repo 项目主页-引用过改项目的所有模板
// q - 搜索关键词
// tags - 标签，如果有多个值用|分开
// recommend  - all 所有，only只展示推荐
// job_type - 任务类型
// compute_source - 计算资源
// dataset_id	- 数据集id
// model_id	- 模型id
// repo_id	- 项目id
// order_by - default 默认,newest 最新创建, recentupdate 最近更新,  collections 收藏数量, usecount 运行次数, name 名称降序, name_asc 名称升序 
// page - 页码 从1 开始
// page_size - 每页数量
export const getAiTaskTmplList = (params) => {
  let url = ``
  switch (params.type) {
    case 'created':
      url = `/api/v1/ai_task_template/list/created`
      break;
    case 'collected':
      url = `/api/v1/ai_task_template/list/collected`
      break;
    case 'admin':
      url = `/api/v1/admin/ai_task_template/list`
      break;
    case 'recommend':
      url = `/api/v1/ai_task_template/list/public?recommend=only`
      break;
    case 'repo':
      url = `/api/v1/ai_task_template/list/all?repo_id=${params.repo}`
      break;
    case 'public':
    default:
      url = `/api/v1/ai_task_template/list/public`
      break;
  }
  return service({
    url: url,
    method: 'get',
    params: {
      ...params
    },
  });
}

/* 创建计算任务模板 */
export const createAiTaskTmpl = (data, params = {}) => {
  return service({
    url: `/api/v1/ai_task_template/create`,
    method: 'post',
    params: {},
    data: { ...data, }
  });
}

/* 编辑计算任务模板 */
export const editAiTaskTmpl = (data) => {
  return service({
    url: `/api/v1/ai_task_template/edit`,
    method: 'post',
    params: { id: data.id },
    data: { ...data, }
  });
}

/* 查询计算任务模板详情 */
export const getAiTaskTmpl = (params) => {
  return service({
    url: `/api/v1/ai_task_template`,
    method: 'get',
    params: { id: params.id }
  });
}

/* 删除计算任务模板 */
export const deleteAiTaskTmpl = (params) => {
  return service({
    url: `/api/v1/ai_task_template`,
    method: 'delete',
    params: { id: params.id }
  });
}

/* 收藏计算任务模板 */
export const putCollectAiTaskTmpl = (params) => {
  return service({
    url: `/api/v1/ai_task_template/collect`,
    method: 'put',
    params: { id: params.id }
  });
}

/* 取消收藏计算任务模板 */
export const deleteCollectAiTaskTmpl = (params) => {
  return service({
    url: `/api/v1/ai_task_template/collect`,
    method: 'delete',
    params: { id: params.id }
  });
}

/* 推荐计算任务模板 */
export const putRecommendAiTaskTmpl = (params) => {
  return service({
    url: `/api/v1/admin/ai_task_template/recommend`,
    method: 'put',
    params: { id: params.id }
  });
}

/* 取消推荐计算任务模板 */
export const deleteRecommendAiTaskTmpl = (params) => {
  return service({
    url: `/api/v1/admin/ai_task_template/recommend`,
    method: 'delete',
    params: { id: params.id }
  });
}

/* 搜索模型、数据集、项目 */
// type -model|dataset|repo
// q - keyword
// page
// pageSize
export const doSearchAiTaskTmplConds = (params) => {
  let tableName = ''
  if (params.type == 'model') {
    tableName = 'model';
  }
  if (params.type == 'dataset') {
    tableName = 'dataset';
  }
  if (params.type == 'repo') {
    tableName = 'repository';
  }
  return service({
    url: `/all/dosearch/`,
    method: 'get',
    params: {
      TableName: tableName,
      Key: params.q,
      Page: params.page,
      PageSize: params.pageSize,
      OnlyReturnNum: false,
      OnlySearchLabel: false,
      WebTotal: 0,
      PrivateTotal: 0,
      language: lang
    },
  });
}

/* 搜索默认的模型、数据集、项目 */
// key -配置文件路径
export const getDefaultAiTaskTmplConds = (params) => {
  return service({
    url: `/api/v1/dync_config`,
    method: 'get',
    params,
    data: {},
  });
}

//获取个人信息页面中的计算任务模板列表
export const getProfileAITaskTemplate = (params) => {
  return service({
    url: '/api/v1/ai_task_template/list/accessible',
    method: "get",
    params: params
  });
}

//获取个人信息页面中的计算任务模板列表（未登录态使用，仅返回公开模板）
export const getProfileAITaskTemplatePublic = (params) => {
  return service({
    url: '/api/v1/ai_task_template/list/public',
    method: "get",
    params: params
  });
}
