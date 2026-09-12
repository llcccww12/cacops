import service from "../service";
import Qs from 'qs';

// 算力需求-算力资源查询智算中心列表
export const getAvailableAiCenterList = () => {
  return service({
    url: `/api/v1/resources/ai_center/available`,
    method: 'get',
    params: {},
    data: {},
  });
}

// 查询智算列表
export const getAiCenterList = () => {
  return service({
    url: `/explore/card_request/resources/queue/centers`,
    method: 'get',
    params: {},
    data: {},
  });
}

// 查询所有资源队列名称列表
export const getResQueueCode = (params) => { // cluster
  return service({
    url: `/explore/card_request/resources/queue/codes`,
    method: 'get',
    params,
  });
}

// 获取资源规格清单
// params -cluster,resource,available,
export const getSpecificationList = (params) => {
  return service({
    url: `/explore/card_request/specification/list`,
    method: 'get',
    params,
  });
};

/* 算力资源 */
// 查询卡类型数据
export const getAccCardList = () => {
  return service({
    url: `/api/v1/resources/acc_card/list`,
    method: 'get',
    params: {}
  });
};

// 查询算力资源列表
// params-page,pageSize,resource-GPU|NPU...,accCardType-ASCEND910|...,accCardNum:-1|1|2|4|8...,excludeAccCardNums-"1|2|4|8",centerCode,minPrice,maxPrice-积分值，未填请传-1
export const getResourceList = (params) => {
  return service({
    url: `/explore/card_request/resource/list`,
    method: 'get',
    params: { ...params },
    paramsSerializer: _params => Qs.stringify(_params, { arrayFormat: 'repeat' }),
  });
};

/* 付费算力 */
export const getPayResourceList = (params) => {
  return service({
    url: `/api/v1/pay_computility`,
    method: 'get',
    params: { ...params },
  });
};

/* 合作伙伴 */
export const getPartners = (params) => {
  return service({
    url: `/api/v1/computility_partner`,
    method: 'get',
    params: { ...params },
  });
};

/* 算力需求 */
// 获取创建算力计算资源和卡类型信息
export const getDemandCreationRequired = (params) => {
  return service({
    url: `/explore/card_request/creation/required`,
    method: 'get',
    params: {}
  });
};

// 提交算力需求
// data-compute_resource,card_type,acc_cards_num,disk_capacity,resource_type,begin_date,end_date,contact,phone_number,email_address,org,description
export const postDemand = (data) => {
  return service({
    url: `/explore/card_request/create`,
    method: 'post',
    params: {},
    data: { ...data },
  });
};

// 用户修改算力需求
// data-compute_resource,card_type,acc_cards_num,disk_capacity,resource_type,begin_date,end_date,contact,phone_number,email_address,org,description
export const updateDemand = (data) => {
  return service({
    url: `/explore/card_request/update/${data.id}`,
    method: 'put',
    params: {},
    data: { ...data },
  });
};

// 算力需求广场列表
// page,pageSize
export const getDemandList = (params) => {
  return service({
    url: `/explore/card_request/list`,
    method: 'get',
    params: { page: params.page, pageSize: params.pageSize }
  });
};

// 我提交的算力需求列表
// page,pageSize
export const getDemandMyList = (params) => {
  return service({
    url: `/explore/card_request/my_list`,
    method: 'get',
    params: { page: params.page, pageSize: params.pageSize }
  });
};

/* 国产算力 */
// 获取国产算力卡使用情况相关数据
// params - type-all|7|30, category-card|user|task
export const getDomesticCardData = (params) => {
  return service({
    url: `/api/v1/cloudbrainboard/card_data`,
    method: 'get',
    params: {
      type: params.type,
      category: params.category,
    }
  });
};
