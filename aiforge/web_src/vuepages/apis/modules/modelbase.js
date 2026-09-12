import service from '../service';

// 检验openi-notebook是否存在，如果不存在先创建openi-notebook仓库，再返回仓库存在
export const getCheckRepo = () => {
  return service({
    url: `/api/v1/finetune/checkRepo`,
    method: 'get',
    params: {},
  });
}

// 获取大模型微调可用规格
export const getSpecInfo = (params) => {
  return service({
    url: `/api/v1/finetune/spec`,
    method: 'get',
    params: {},
  });
}

// 创建大模型微调训练任务
export const setFinetuneCreate = (data) => {
  return service({
    url: `/api/v1/repos/${data.userName}/openi-notebook/finetune/create`,
    method: 'post',
    params: {},
    data: {
      type: data.type, // 0为脑海(盘古2.6模型)
      display_job_name: data.display_job_name,
      description: data.description,
      attachment: data.attachment,
      dataset_name: data.dataset_name,
      sample_dataset_type: data.sample_dataset_type,
      spec_id: data.spec_id,
      run_para_list: data.run_para_list,
    },
  });
}

// 查询训练任务状态
export const getJobStatus = (params) => {
  return service({
    url: `/api/v1/repos/${params.userName}/openi-notebook/modelarts/train-job/${params.jobId}`,
    method: 'get',
    params: {},
  });
}

// 停止训练任务
export const stopTrainJob = (params) => {
  return service({
    url: `/api/v1/repos/${params.userName}/openi-notebook/modelarts/train-job/${params.jobId}/stop_version`,
    method: 'post',
    params: {},
    data: {},
  });
}

// 删除训练任务
export const deleteTrainJob = (params) => {
  return service({
    // url: `/api/v1/repos/${params.userName}/openi-notebook/modelarts/train-job/${params.jobId}/del_version`,    
    url: `/api/v1/${params.userName}/openi-notebook/ai_task/del?id=${params.id}`,
    method: 'post',
    params: {},
  });
}

// 查询微调任务列表
export const getFinetuneList = (params) => {
  return service({
    url: `/api/v1/repos/${params.userName}/openi-notebook/finetune`,
    method: 'get',
    params: {},
  });
}

// 新建ModelArt部署
export const setFinetuneService = (data) => {
  return service({
    url: `/api/v1/repos/${data.userName}/openi-notebook/finetune/deploy/create`,
    method: 'post',
    params: {},
    data: {
      type: data.type, // 0为脑海(盘古2.6模型)
      job_id: data.jobId,
      sample_dataset_type: data.sample_dataset_type, // 1文本分类, 2中英翻译, 3开放问答, 0自定义
    },
  });
}

// 查询部署状态
// params job_id
// return job_id, DEPLOYING部署中, SUCCEEDED部署成功, STOP停止, FAIL部署失败
export const getFinetuneServiceStatus = (params) => {
  return service({
    url: `/api/v1/repos/${params.userName}/openi-notebook/finetune/deploy/${params.jobId}`,
    method: 'get',
    params: {},
  });
}

// 更新部署状态
// return { code-0成功,1失败 }
export const updateFinetuneService = (data) => {
  return service({
    url: `/api/v1/repos/${data.userName}/openi-notebook/finetune/deploy/update`,
    method: 'post',
    params: {},
    data: {
      job_id: data.jobId,
      status: data.status, // 停止或重启部署服务，分别对应stopped或running
    }
  });
}

// 删除部署
// return { code-0成功,1失败 }
export const deleteFinetuneService = (data) => {
  return service({
    url: `/api/v1/repos/${data.userName}/openi-notebook/finetune/deploy/delete`,
    method: 'post',
    params: {},
    data: { job_id: data.jobId, }
  });
}

// 推理体验
// params jobId, text,
export const getFinetuneServiceInference = (data) => {
  return service({
    url: `/api/v1/repos/${data.userName}/openi-notebook/finetune/deploy/inference`,
    method: 'post',
    params: {},
    data: {
      job_id: data.jobId,
      text: data.text,
    }
  });
}
