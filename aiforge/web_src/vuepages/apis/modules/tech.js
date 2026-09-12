/* 科技项目汇聚功能相关 */
import service from '../service';

// 新建申请页面科技项目查询接口 no-项目编号,name-科技项目名称,institution-参与单位
export const getTechs = (params) => {
  return service({
    url: '/api/v1/tech',
    method: 'get',
    params: params,
  });
}

// 新建启智项目申请页面提交 url-启智项目地址,no-项目立项编号,institution-贡献单位，多个单位用逗号分隔
export const setOpenIApply = (data) => {
  return service({
    url: '/api/v1/tech/openi',
    method: 'post',
    data: data,
    params: {},
  });
}

// 新建启智项目申请页面提交 
// url-启智项目地址,uid-启智项目uid,repo_name-启智项目名称,topics-关键词,description-简介,
// no-项目立项编号,institution-贡献单位，多个单位用逗号分隔
export const setNoOpenIApply = (data) => {
  return service({
    url: '/api/v1/tech/no_openi',
    method: 'post',
    data: data,
    params: {},
  });
}

// 新建非启智项目申请页面项目路径可选的用户和组织
// 返回 [{id:用户ID,name:用户名字,rel_avatar_link:用户图像地址,short_name:用户短名称},...]
export const getCreateRepoUser = () => {
  return service({
    url: '/api/v1/user/owners',
    method: 'get',
    params: {},
  });
}

// 项目过滤信息列表-筛选信息 type=0查询科技项目页， type=1查询启智项目页
// 返回 {type_name:项目类型的名称, 字符串数组,institution_name:参与单位的名称，字符串数组,execute_year:执行年份，字符串数组,
// apply_year:申请年份，字符串数组,topic-关键词，字符串数组,project_name:科技项目名称，字符串数组}
export const getTechFilterInfo = (params) => {
  return service({
    url: '/api/v1/tech/filter',
    method: 'get',
    params: {
      type: params.type,
    },
  });
}

// 按科技项目页查询功能 
// 输入 name-科技项目名称，模糊匹配,type_name-项目类型的名称，精确匹配,institution_name-参与单位的名称，精确匹配,execute_year-执行年份，精确匹配,apply_year-申请年份，精确匹配,page,pageSize,sort
// 返回 {"total":10,"data":[{}]}
// data: {id:科技项目id,project_name:科技项目名称,institution:承担单位,all_institution:参与单位,apply_year:申报年份,execute_period:执行周期,repo_numer:项目成果数,repo:[]项目信息数组}
export const getTechSearch = (params) => {
  return service({
    url: '/api/v1/tech/search',
    method: 'get',
    params: params,
  });
}

// 按启智项目页查询功能
// 输入 name-项目名称，模糊匹配,tech_name-科技项目名称，精确匹配,institution_name-参与单位的名称，精确匹配,topic-关键词，精确匹配,page,pageSize,sort
// 返回 {"total":10,"data":[{}]}
// data: {id:项目id,owner_id:项目拥有者id,name:项目名称,alias:项目别名,topics:关键词 字符串数组,description:项目简介,updated_unix:最后更新时间,institution:贡献单位}
export const getTechOpenISearch = (params) => {
  return service({
    url: '/api/v1/tech/repo_search',
    method: 'get',
    params: params,
  });
}

// 我申请的项目列表
// 输入 page,pageSize
// 返回 {"total":10,"data":[{}]}
// data: {id:记录的id,name:启智项目名称,owner_name:启智项目拥有者名称,tech_name:科技项目名称,tech_number:科技项目编号,institution:项目承担单位,
// execute_period:执行周期,contact:联系人,contact_phone:联系电话,contact_email:联系邮件,apply_user:申请账号,status:状态}
export const getTechMyList = (params) => {
  return service({
    url: '/api/v1/tech/my',
    method: 'get',
    params: params,
  });
}

// 后台管理项目列表
// 输入 page,pageSize
// 返回 {"total":10,"data":[{}]}
// data: {id:记录的id,name:启智项目名称,owner_name:启智项目拥有者名称,tech_name:科技项目名称,tech_number:科技项目编号,institution:项目承担单位,
// execute_period:执行周期,contact:联系人,contact_phone:联系电话,contact_email:联系邮件,apply_user:申请账号,status:状态}
export const getTechAdminList = (params) => {
  return service({
    url: '/api/v1/tech/admin',
    method: 'get',
    params: params,
  });
}

// 后台管理项目列表交互(显示或隐藏)
// 输入 type-show/hide,id:[1,2,3]
// 返回 { code:0, message:"" }
export const setTechAdminOperation = (data) => {
  return service({
    url: `/api/v1/tech/admin/${data.type}`,
    method: 'post',
    data: {
      id: data.id
    },
    params: {},
  });
}

// 后台管理科技项目编号基本信息导入功能（科技项目管理员有权限)
// 输入 form file
export const setTechImportExcel = () => {
  return service({
    url: `/api/v1/tech/basic`,
    method: 'post',
    data: {},
    params: {},
  });
}

// 增加科技项目管理员(管理员才有权限)
// 输入 name:['111', '222']
// 返回 { code:0, message:"" }
export const setTechAdminAdd = (data) => {
  return service({
    url: '/api/v1/tech/admin_add',
    method: 'post',
    data: data,
    params: {},
  });
}

// 判断是否是科技项目管理员
// 输入 name:['111', '222']
export const getIsTechAdmin = () => {
  return service({
    url: '/api/v1/tech/is_admin',
    method: 'get',
    params: {},
  });
}

// =========================================================
// 搜索Topics q-topic
// 返回 {topics: [{id,repo_count,topic_name,created,updated}]}
export const getTopics = (params) => {
  return service({
    url: '/api/v1/topics/search',
    method: 'get',
    params: params,
  });
}

// 检测 repo name q:名称, owner:所属者 
export const getCheckRepoName = (params) => {
  return service({
    url: '/repo/check_name',
    method: 'get',
    params: params,
  });
}
