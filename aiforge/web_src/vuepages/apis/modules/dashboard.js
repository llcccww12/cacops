import service from "../service";

//新增概览页面接口
export const getOverviewData = () => {
  return service({
    url: `/api/v1/platform/overview`,
    method: "get",
    params: {}
  });
};


export const getOverviewConfig = () => {
  return service({
    url: `/api/v1/platform/official_config`,
    method: "get",
    params: {}
  });
};