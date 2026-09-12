import service from '../service';

// 算力积分概要
export const getPointAccount = () => {
  return service({
    url: '/reward/point/account',
    method: 'get',
    params: {},
  });
}

// 算力积分获取、消耗明细
// operate-INCREASE 表示获取明细  DECREASE表示消耗明细, page-当前页, pageSize-每页条数
export const getPointList = (params) => {
  return service({
    url: '/reward/point/record/list',
    method: 'get',
    params,
  });
}

// 管理员充值、扣减用户积分
// TargetUserId, OperateType-INCREASE,DECREASE, Amount, Remark, RewardType-POINT
export const setPointOperate = (data) => {
  return service({
    url: '/operation/reward/point/account/operate',
    method: 'post',
    data,
    params: {}
  });
}

// 算力积分页面
export const getPoint = () => {
  return service({
    url: '/reward/point',
    method: 'get',
    params: {},
    data: {},
  });
}
