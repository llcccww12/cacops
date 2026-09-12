import service from '../service';

// 查询管理后台oauth应用列表
// q, scopes -user.base,user.email  verify_flag -0 所有，1 未认证 2 已认证 3 认证取消
export const getApplicaitonList = (params) => {
  return service({
    url: '/api/v1/admin/oauth2/application/list',
    method: 'get',
    params: params,
  });
}

// 编辑管理后台oauth应用
// ID,VerifyFlag-1 未认证 2 已认证 3 认证取消 ScopeList-user.base,user.email
export const setEditApplicaiton = (data) => {
  return service({
    url: '/api/v1/admin/oauth2/application/edit',
    method: 'post',
    params: {},
    data
  });
}
