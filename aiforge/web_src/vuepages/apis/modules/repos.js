import service from '../service';

// 获取首页数据
export const getHomePageData = () => {
  return service({
    url: '/recommend/home',
    method: 'get',
    params: {},
  });
}

// 获取项目广场上方tab数据 tab=preferred 项目优选|incubation 启智孵化管道|hot-paper 热门论文项目
export const getReposSquareTabData = (tab) => {
  return service({
    url: '/explore/repos/square/tab',
    method: 'get',
    params: {
      type: tab
    },
  });
}

// 搜索项目
// q	string	否	关键词
// topic	string	否	标签名
// sort	string	是	mostpopular 近期热门 | mostactive 近期活跃 | recentupdate 最近更新 | newest 最近创建 
//                  moststars 点赞最多 | mostforks 派生最多 | mostdatasets 数据集最多 | mostaitasks AI任务最多 | mostmodels 模型最多                
// pageSize	int	是	每页大小，可选值为15 | 30 | 50
// page	int	是	页码
export const getReposListData = (params) => {
  return service({
    url: '/explore/repos/search',
    method: 'get',
    params: {
      q: params.q || '',
      topic: params.topic || '',
      sort: params.sort || 'mostpopular',
      pageSize: params.pageSize || 15,
      page: params.page || 1,
    },
  });
}
export const getReposCollectListData = (params) => {
  return service({
    url: 'api/v1/user/subscriptions',
    method: 'get',
    params: params,
  });
}
// 获取活跃用户列表
export const getActiveUsers = () => {
  return service({
    url: '/explore/repos/square/active-user',
    method: 'get',
    params: {},
  });
}

// 关注用户
export const followingUsers = (userName, isFollowing) => {
  return service({
    url: `/api/v1/user/following/${userName}`,
    method: isFollowing ? 'put' : 'delete',
    params: {},
  });
}

// 获取活跃组织列表
export const getActiveOrgs = () => {
  return service({
    url: '/explore/repos/square/active-org',
    method: 'get',
    params: {},
  });
}

// 获取项目列表
export const getRepoList = (params) => {
  return service({
    url: '/api/v1/repos/search',
    method: 'get',
    params: params,
  });
}
// 获取项目列表新接口
export const getNewRepoList = (params) => {
  return service({
    url: '/api/v1/repos/platform/search',
    method: 'get',
    params: params,
  });
}
// 获取项目分支
export const getRepoBranch = (params) => {
  return service({
    url: `/api/v1/${params.repoOwnerName}/${params.repoName}/ai_task/repo_branch`,
    method: 'get',
    params: {},
  });
}


// 关注用户
export const deleteRepos = (params) => {
  return service({
    url: `/api/v1/repos/${params.ownerName}/${params.name}/project`,
    method: 'delete',
    params: {},
  });
}

// 关注用户
export const pinnedRepos = (userName, reposName, isPinning) => {
  return service({
    url: `/api/v1/repos/${userName}/${reposName}/project/pinned`,
    method: !isPinning ? 'post' : 'delete',
    params: {},
  });
}