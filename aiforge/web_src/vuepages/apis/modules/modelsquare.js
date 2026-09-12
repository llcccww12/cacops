import service from "../service";
import Qs from 'qs';

// 获取模型广场列表
export const getModelList = ({ url, ...params }) => {
  return service({
    url: `/api/v1/aimodel/list${url}`,
    method: "get",
    params: params
  });
};

// 模型详情侧边栏相关信息
export const getModelRelatedInfo = (params) => {
  return service({
    url: `/api/v1/aimodel/related`,
    method: "get",
    params: params
  });
};

// 获取模型广场筛选项
export const getModelSqaureFilters = () => {
  return service({
    url: '/modelsquare/main_query_label',
    method: 'get',
    params: {},
  });
}

// 模型收藏/取消收藏
// data: id, collected-为true表示收藏此模型，为false表示取消收藏此模型
export const setModelFav = (data) => {
  return service({
    url: '/modelsquare/modify_model_collect',
    method: 'put',
    data: Qs.stringify(data),
  });
}

