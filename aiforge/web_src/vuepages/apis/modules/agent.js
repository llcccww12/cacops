import service from '../service';

const API_BASE = '/api/v1'

// 1. 获取智能体广场列表
export const getAgentList = (params) => {
  return service({
    url: `${API_BASE}/agent/list`,
    method: 'get',
    params: { q: params.q, page: params.page, page_size: params.pageSize },
  });
}

// 2. 获取用户智能体列表
export const getUserAgentList = (params) => {
  return service({
    url: `${API_BASE}/ai_task/agent/user_list`,
    method: 'get',
    params: { page: params.page, page_size: params.pageSize },
  });
}

// 3. 启动智能体
export const runAgent = (id) => {
  return service({
    url: `${API_BASE}/ai_task/agent/${id}/run`,
    method: 'post',
    params: {}
  });
}

// 4. 停止智能体
export const stopAgent = (id) => {
  return service({
    url: `${API_BASE}/ai_task/agent/${id}/stop`,
    method: 'post',
    params: {}
  });
}

// 5. 创建智能体（管理员）
export const createAgent = (data) => {
  return service({
    url: `${API_BASE}/ai_task/agentAdmin/create`,
    method: 'post',
    params: {},
    data,
  });
}

// 5. 删除智能体（管理员）
export const deleteAgent = (id) => {
  return service({
    url: `${API_BASE}/ai_task/agentAdmin/${id}`,
    method: 'delete',
    params: {},
  });
}