import service from '~/apis/service';

// 获取promote配置数据
export const getPromoteData = (filePathName) => {
  return service({
    url: '/dashboard/invitation',
    method: 'get',
    params: {
      filename: filePathName
    },
  });
}
