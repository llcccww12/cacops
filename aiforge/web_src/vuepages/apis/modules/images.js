import service from '../service';
import Qs from 'qs';

// 获取镜像
// params: { type-0|1|2, q, page, pageSize, cloudbrainType-0,1, sort, framework, frameworkVersion, cuda, python }
export const getImages = (params) => {
  const typeMap = {
    '0': 'recommend',
    '1': 'custom',
    '2': 'star',
  };
  let type = 'recommend';
  if (params.type && typeMap[params.type]) {
    type = typeMap[params.type];
  }
  return service({
    url: `/api/v1/images/${type}`,
    method: "get",
    params: {
      q: params.q || '',
      cloudbrainType: params.cloudbrainType || '-1',
      jobType: params.jobType,
      sort: params.sort,
      computeResource: params.computeResource,
      framework: params.framework,
      frameworkVersion: params.frameworkVersion,
      cuda: params.cuda,
      dtk: params.dtk,
      cann: params.cann,
      python: params.python,
      spec: params.spec,
      has_internet: params.hasInternet,
      visualize_required: params.visualizeRequired,
      trainType: params.trainType,
      onlyOpenIImage: params.onlyOpenIImage,
      page: params.page || 1,
      pageSize: params.pageSize || 5,
    }
  });
};

export const putImageAction = (params) => {
  return service({
    url: `/image/${params.id}/action/${params.action}`,
    method: 'put',
    params: {}
  });
};

export const deleteImage = (params) => {
  return service({
    url: `/image/${params.id}`,
    method: 'delete',
    params: {}
  });
};

export const submitImage = (data) => {
  return service({
    url: data.link,
    method: 'post',
    params: {},
    data: Qs.stringify(data),
  });
};

export const getImageById = (params) => {
  return service({
    url: `/image/${params.id}`,
    method: 'get',
    params: {},
  });
};

export const searchImageTopics = (params) => {
  return service({
    url: `/api/v1/image/topics/search`,
    method: 'get',
    params: { q: params.q }
  });
};

// 查询选择镜像过滤条件
// params: index-查询类型:0(可用框架)|1(可用框架版本)|2(可用python版本)|3(可用cuda版本)
//         framework-框架,version-框架版本,python-python版本
export const getImageAvailabelFilter = (params) => {
  return service({
    url: `/api/v1/images/availableFilter`,
    method: 'get',
    params: { ...params }
  });
};



export const syncImage = () => {
  return service({
    url: '/admin/resources/image/sync',
    method: 'post',
    params: {},
  });
};

export const getImageListCustom = (params) => {
  return service({
    url: `/admin/images/data`,
    method: 'get',
    params: params
  });
};


export const putImageRecommend = (id) => {
  return service({
    url: `/admin/image/${id}/action/recommend`,
    method: 'put',
    params: {}
  });
};

export const putImageUnRecommend = (id, data) => {
  return service({
    url: `/admin/image/${id}/action/unrecommend`,
    method: 'put',
    params: {},
    data: data
  });
};

