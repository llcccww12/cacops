import service from '../service';
import Qs from 'qs';
// 查询智算列表
export const getAiCenterList = () => {
  return service({
    url: '/admin/resources/queue/centers',
    method: 'get',
    params: {},
    data: {},
  });
}

// 查询资源队列列表
// page	当前页数，从1开始
// cluster 所属集群 ：OpenI 启智集群，C2Net 智算集群
// center 智算中心：OpenIOne 云脑一，OpenITwo 云脑二， chendu 成都人工智能计算中心， pclcci 鹏城云计算所 ，hefei 合肥类脑类脑智能开放平台， xuchang 中原人工智能计算中心
// resource 计算资源： GPU NPU
// card XPU类型： T4、A100、V100、Ascend 910
export const getResQueueList = (params) => {
  return service({
    url: '/admin/resources/queue/list',
    method: 'get',
    params,
  });
}

// 新增资源队列
export const addResQueue = (data) => { // Cluster,QueueCode,AiCenterCode,ComputeResource,AccCardType,CardsTotalNum,Remark
  return service({
    url: '/admin/resources/queue/add',
    method: 'post',
    params: {},
    data,
  });
}

// 更新资源队列
export const updateResQueue = (data) => { // CardsTotalNum,Remark
  return service({
    url: `/admin/resources/queue/update/${data.ID}`,
    method: 'post',
    params: {},
    data,
  });
}

// 查询所有资源队列名称列表
export const getResQueueCode = (params) => { // cluster
  return service({
    url: '/admin/resources/queue/codes',
    method: 'get',
    params,
    data: {},
  });
}
// 查询所有算力中心名称列表
export const getAiCenterCode = (params) => { // cluster
  return service({
    url: '/admin/resources/queue/centers',
    method: 'get',
    params,
    data: {},
  });
}
// 同步智算网络资源池（队列）
export const syncResQueue = () => {
  return service({
    url: '/admin/resources/queue/grampus/sync',
    method: 'post',
    params: {},
    data: {},
  });
}

// 新增资源规格
export const addResSpecification = (data) => {
  return service({
    url: '/admin/resources/specification/add',
    method: 'post',
    params: {},
    data,
  });
}

// 查询资源规格所属场景 - 下架时提醒
export const getResSpecificationScenes = (data) => { // data => { ID: 1 }
  return service({
    url: `/admin/resources/specification/scenes/${data.ID}`,
    method: 'get',
    params: {},
    data: {}
  });
}

// 更新资源规格
// params: action edit-编辑 on-shelf 上架 off-shelf 下架
// data: UnitPrice
export const updateResSpecification = (data) => { // data => { ID: 1, action: 'edit|on-shelf|off-shelf', UnitPrice: 1 | undefined }
  return service({
    url: `/admin/resources/specification/update/${data.ID}`,
    method: 'post',
    params: { action: data.action, comment: data.comment },
    data: { 
      UnitPrice: data.action === 'edit' || data.action === 'on-shelf' ? data.UnitPrice : undefined,
    }
  });
}

// 查询资源规格列表
// page
// cluster 所属集群 ：OpenI 启智集群，C2Net 智算集群
// queue 所属队列id
// status 状态 ： 1 待审核 2已上架 3已下架
// cardType 卡类型
// cardsNum 卡数
// resource 计算资源
export const getResSpecificationList = (params) => {
  return service({
    url: '/admin/resources/specification/list',
    method: 'get',
    params,
    data: {},
  });
}

// 查询资源规格上下架操作日志
export const getResSpecificationShelfLogList = (params) => {
  return service({
    url: '/admin/resources/specification/log',
    method: 'get',
    params,
    data: {},
  });
}

// 查询资源规格列表(所有)
// cluster 所属集群 ：OpenI 启智集群，C2Net 智算集群
// queue 所属队列id
// status 状态 ： 1 待审核 2已上架 3已下架
// available 是否可用 -1全部 1可用 2不可用
export const getResSpecificationListAll = (params) => {
  return service({
    url: '/admin/resources/specification/list/all',
    method: 'get',
    params,
    data: {},
  });
}

// 同步智算网络资源池（队列）
export const syncResSpecification = () => {
  return service({
    url: '/admin/resources/specification/grampus/sync',
    method: 'post',
    params: {},
    data: {},
  });
}

// 新增资源应用场景
/*
{
  "SceneName":"启智集群调试任务",    //应用场景名
  "JobType":"TRAIN",  //任务类型  DEBUG调试任务  BENCHMARK 评测任务 TRAIN 训练 INFERENCE 推理
  "sceneType":"public",
  "IsSpecExclusive":true,  //是否专属
  "ExclusiveOrg":"123,456",  //专属组织
  "SpecIds":[2,3]   // 资源规格id
}
*/
export const addResScene = (data) => {
  return service({
    url: '/admin/resources/scene/add',
    method: 'post',
    params: {},
    data,
  });
}

// 更新资源应用场景
// params: action:edit-编辑 delete-删除, 
// data: {
//  "SceneName":"启智集群调试任务",    //应用场景名
//	"IsSpecExclusive":true,  //是否专属
//	"ExclusiveOrg":"123,456",  //专属组织
//   "SpecIds":[2,3]   // 资源规格id
//}
export const updateResScene = (data) => {
  return service({
    url: `/admin/resources/scene/update/${data.ID}`,
    method: 'post',
    params: { action: data.action },
    data: {
      ...data
    },
  });
}

// 查询资源应用场景
// page
// jobType 
// center
// queue 所属队列
// sceneType
// IsSpecExclusive 是否专属 1 专属 2 非专属
// cardType 卡类型
// resource 计算资源
export const getResSceneList = (params) => {
  return service({
    url: '/admin/resources/scene/list',
    method: 'get',
    params,
    data: {},
  });
}

export const getResourceType = () => {
  return service({
    url: '/api/v1/admin/role/list_quene',
    method: 'get',
    params:{},
    data: {},
  });
}


export const addAiforgeRole = (data) => {
  return service({
    url: `/api/v1/admin/role/add`,
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    method: 'post',
    params: { },
    data: Qs.stringify(data),
  });
}


export const updateAiforgeRole = (data) => {
  return service({
    url: `/api/v1/admin/role/update`,
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    method: 'post',
    params: { },
    data: Qs.stringify(data),
  });
}

export const listAiforgeRole = (params) => {
  return service({
    url: `/api/v1/admin/role/list`,
    method: 'get',
    params: params,

  });
}

export const listOperation = () => {
  return service({
    url: `/api/v1/admin/role/listOperation`,
    method: 'get',
    params: { },
  });
}

export const delAiforgeRole = (params) => {
  return service({
    url: `/api/v1/admin/role/del`,
    method: 'delete',
    params: params,

  });
}

export const listRightUser = (params) => {
  return service({
    url: `/api/v1/admin/role/list_right_user`,
    method: 'get',
    params: params,
  });
}


export const setAiforgeRoleToUser = (data) => {
  return service({
    url: `/api/v1/admin/role/settouser`,
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    method: 'post',
    params: { },
    data: Qs.stringify(data),
  });
}


export const bacthAddRoleToUser = (data) => {
  return service({
    url: `/api/v1/admin/role/batchAddRoleToUser`,
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    method: 'post',
    params: { },
    data: Qs.stringify(data),
  });
}

export const bacthDelRoleToUser = (data) => {
  return service({
    url: `/api/v1/admin/role/batchDelRoleToUser`,
    headers: { 'Content-type': 'application/x-www-form-urlencoded' },
    method: 'post',
    params: { },
    data: Qs.stringify(data),
  });
}