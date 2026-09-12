import service from '../service';

// 邀请好友页面数据
export const getUserInvitationCode = (params) => { // page pageSize
  return service({
    url: '/user/invitation_code',
    method: 'get',
    params: params,
    data: {},
  });
}
