import service from "../service";

// Notebook新建页面需要的信息
export const getFileNotebook = () => {
  return service({
    url: "/api/v1/file_notebook",
    method: "get",
    params: {},
  });
};

// Notebook新建调试任务
// type, file, branch_name, owner_name, project_name
export const createNotebook = (data) => {
  return service({
    url: "/api/v1/file_notebook/create",
    method: "post",
    data,
    params: {},
  });
};

// Notebook获取云脑I调试任务状态
export const getCb1Notebook = (path,jobid) => {
  return service({
    url: `/api/v1/${path}/cloudbrain/${jobid}`,
    method: "get",
    params: {},
  });
};

// Notebook获取云脑II调试任务状态
export const getCb2Notebook = (path,jobid) => {
  return service({
    url: `/api/v1/${path}/modelarts/notebook/${jobid}`,
    method: "get",
    params: {},
  });
};
// Notebook查询文件在环境中是否已准备好
// type, file, branch_name, owner_name, project_name,job_id
export const getFileInfoNotebook = (data) => {
  return service({
    url: "/api/v1/file_notebook/status",
    method: "post",
    data,
    params: {},
  });
};
export const stopNotebook = (url) => {
    return service({
      url: url,
      method: "post",
      params: {},
    });
  };