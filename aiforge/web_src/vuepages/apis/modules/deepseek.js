import service from '../service';
const csrf = window.config ? window.config.csrf : "";
// llm 模型对话
export const llmChat = (data, { signal } = {}) => {
  return fetch(`/api/v1/llm/chat/deepseek?_csrf=${csrf}`,{
    headers:{
      'Content-Type': 'application/json'
    },
    method: 'POST',
    body: JSON.stringify(data),
    signal
  })
}

export const llmChatFeedBack = (data) => {
  return service({
    url: `/api/v1/llm/feedback/deepseek`,
    method: 'post',
    params: {},
    data: data,
  });
}