import axios from 'axios';

const service = axios.create({
  baseURL: '/',
});

service.interceptors.request.use((config) => {
  const csrf = window.config ? window.config.csrf : '';
  if (csrf) {
    config.headers = config.headers || {};
    config.headers['X-Csrf-Token'] = csrf;
  }
  config.data && Object.assign(config.data, {
    _csrf: csrf,
  });
  config.params && Object.assign(config.params, {
    _csrf: csrf,
  });
  return config;
}, (error) => {
  return Promise.reject(error);
});

service.interceptors.response.use((response) => {
  if (response.status == 200 && response.data && response.data.code == 9002) { // 绑定微信
    window.location.href = `/authentication/wechat/bind?redirect_to=${encodeURIComponent(window.location.href)}`;
    return Promise.reject(response);
  }
  return response;
}, (error) => {
  return Promise.reject(error);
});

export default service;
