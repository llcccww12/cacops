import service from "../service";

// 获取中心详细情况
// 
export const getAiCenterOverview = () => {
  return service({
    url: `/api/v1/cloudbrainboard/ai_center_overview`,
    method: "get",
  });
};

export const getAiLocation = () => {
    return service({
      url: `/api/v1/cloudbrainboard/location`,
      method: "get",
    });
  };


// 获取中心详细情况
// 
export const getAiCenterCardInfo = () => {
  return service({
    url: '/api/v1/resources/ai_center/card_info/active',
    method: "get",
  });
};

export const getAiActivate = (params) => {
    return service({
      url: '/api/v1/resources/ai_center/active',
      method: "get",
      params
    });
  };