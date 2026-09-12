import { param } from 'jquery';
import service from '../service';
let csrf = window.config ? window.config.csrf : ''


//下载预设speaker音频
export const getSpeakerVoice = (params) => {
  return service({
    url: "/api/v1/tts/get_speaker_voice",
    headers:{
      'Content-Type': 'audio/wav'
    },
    responseType: 'blob',
    method: "get",
    params: params
  });
}
//发送语音合成
export const getSynthesize = (data) => {
  return service({
    url: "/api/v1/tts/synthesize",
    headers:{
      'Content-Type': 'audio/wav'
    },
    responseType: 'blob',
    method: "post",
    params: {},
    data: data
  });
}
// 在线体验
export const onlineExperience = (params) => {
  return service({
    url: `/api/v1/llm/chat/visit`,
    method: 'post',
    params: params,
    data: {}
  });
}
// 创建知识库
export const llmAgree = (params) => {
  return service({
    url: `/api/v1/llm/chat/agree`,
    method: 'post',
    params: params,
  });
}
//llm count
// 查询知识库列表
export const llmCount = (params) => {
  return service({
    url: '/api/v1/llm/chat/counts',
    method: 'get',
    params: params,
  });
}
// llm 模型对话
export const llmChat = (data) => {
  return fetch(`/api/v1/llm/chat/completion?_csrf=${csrf}`,{
    headers:{
      'Content-Type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify(data)
  })
}

// llm 知识库对话
export const llmKbChat = (data,model_name) => {
  return fetch(`/api/v1/llm/chat/knowledge_base_chat?_csrf=${csrf}&model_name=${model_name}`,{
    headers:{
      'Content-Type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify(data)
  })
}

// 创建知识库
export const llmKbcreate = (data,params) => {
  return service({
    url: `/api/v1/llm/knowledge_base/create`,
    method: 'post',
    params: params,
    data: data
  });
}

// 删除知识库
export const llmKbDelete = (params) => {
  return service({
    url: '/api/v1/llm/knowledge_base/delete',
    method: 'post',
    params: params,

  });
}

// 删除知识库文件
export const llmKbDeleteDoc = (data,params) => {
  return service({
    url: '/api/v1/llm/knowledge_base/delete_doc',
    method: 'post',
    params: params,
    data: data
  });
}


// 查询知识库列表
export const llmKbList = (params) => {
  return service({
    url: '/api/v1/llm/knowledge_base/list',
    method: 'get',
    params: params,
  });
}

// 知识库文件上传
export const llmKbUploadDocUrl = (params,data) => {
  return service({
    url: '/api/v1/llm/knowledge_base/upload_doc',
    method: 'post',
    params: params,
    data: data
  });
}
export const llmKbUploadDoc = (url,data) => {
    return service({
      url: url,
      method: 'post',
      params: {},
      data: data
    });
}
// 查询当前知识库文件列表


// llmchat
export const llmChatCompletion = (data) => {
  return service({
    url: '/api/v1/llm/chat/completion',
    method: 'post',
    params: {},
    data: data
  });
}

export const llmKbFileList = (params) => {
  return service({
    url: '/api/v1/llm/knowledge_base/list_files',
    method: 'get',
    params: params,
  });
}

export const llmRecreateVectorStore = (params) => {
  return fetch(`/api/v1/llm/knowledge_base/recreate_vector_store?_csrf=${csrf}&knowledge_base_name=${params.knowledge_base_name}&model_name=${params.model_name}`, {
      headers: {
        'Content-Type': 'application/json'
      },
      method: 'POST',
  })
}


export const getLoraStage = (params) => {
  return service({
    url: '/api/v1/sd/finetune/stage',
    method: 'get',
    params: params,
  });
}


export const sdLorauploadImages = (url,data) => {
  return service({
    url: `${url}/upload_images`,
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    method: 'post',
    data: data,
  });
}
export const sdLorauploadZip = (url,data) => {
  return service({
    url: `${url}/upload_zip`,
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    method: 'post',
    data: data,
  });
}

// export const getUploadImagesId = (url,uuid) => {
//   return service({
//     url: `${url}/get_image/${uuid}`,
//     responseType: 'blob',
//     method: 'get',
//     params: {},
//   });
// }


export const postClearImagesId = (url) => {
  return service({
    url: `${url}/clear_all_uuids`,
    method: 'post',
    params: {},
  });
}

export const postDeleteImagesId = (url,uuid) => {
  return service({
    url: `${url}/delete/${uuid}`,
    method: 'post',
    params: {},
  });
}

export const postMakeCaption = (url,params) => {
  return service({
    url: `${url}/make_captions`,
    method: 'post',
    params: params,
  });
}

export const getCaptionStatus= (url) => {
  return service({
    url: `${url}/captions_status`,
    method: 'get',
    params: {},
  });
}
export const getCaptionText= (url,uuid) => {
  return service({
    url: `${url}/get_caption/${uuid}`,
    method: 'get',
    params: {},
  });
}

export const getAllCaptionText= (url) => {
  return service({
    url: `${url}/get_all_captions`,
    method: 'get',
    params: {},
  });
}

export const getAllImages= (url) => {
  return service({
    url: `${url}/list_all_uuids`,
    method: 'get',
    params: {},
  });
}

export const postAllAcaption = (url,params) => {
  return service({
    url: `${url}/add_all_captions`,
    method: 'post',
    params: params,
  });
}


export const postSingleCaption = (url,uuid,params) => {
  return service({
    url: `${url}/modify_caption/${uuid}`,
    method: 'post',
    params: params,
  });
}

export const postTrainLora = (params,data) => {
  return service({
    url: `/api/v1/sd/finetune/train_lora`,
    method: 'post',
    params: params,
    data: data,
    
  });
}

export const getTrainLoraProcess = (url) => {
  return service({
    url: `${url}/lora_process`,
    method: 'get',
    params: {},
  });
}



export const postLoraGenerateImg = (params,data) => {
  return service({
    url: `/api/v1/sd/finetune/text2img`,
    method: 'post',
    params: params,
    data: data,
  });
}

export const getLoraGenerateImgProcess = (url) => {
  return service({
    url: `${url}/text2img_process`,
    method: 'get',
    params: {},
  });
}


export const getExampleImages= (url) => {
  return service({
    url: `${url}/load_sample_dataset`,
    method: 'post',
    params: {},
  });
}