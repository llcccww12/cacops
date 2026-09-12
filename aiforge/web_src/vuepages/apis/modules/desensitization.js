import service from "../service";

// 数据脱敏接口
export const postImgDesensitization = (params, data) => {
  return service({
    url: "/extension/tuomin/upload",
    headers: { Accept: "image/png" },
    responseType: "blob",
    method: "POST",
    params: params,
    data: data,
  });
};
// 文心接口
export const postERNIEPaintNew = (params) => {
  return service({
    url: "/extension/wenxin/paint_new",
    method: "get",
    params: params
  });
};
export const postERNIEPaintResult = (params) => {
  return service({
    url: "/extension/wenxin/query_paint_image",
    method: "get",
    params: params
  });
};
export const getERNIEPaintCount = () => {
  return service({
    url: "/extension/wenxin/query_paint_result",
    method: "get",
    params: {}
  });
};
// sd接口
export const postSdPaintNew = (params) => {
  return service({
    url: "/extension/sd/paint_new",
    method: "get",
    params: params
  });
};