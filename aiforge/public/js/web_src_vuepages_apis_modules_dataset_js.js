"use strict";
(self["webpackChunk"] = self["webpackChunk"] || []).push([["web_src_vuepages_apis_modules_dataset_js"],{

/***/ "./web_src/vuepages/apis/modules/dataset.js":
/*!**************************************************!*\
  !*** ./web_src/vuepages/apis/modules/dataset.js ***!
  \**************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   createDatasets: function() { return /* binding */ createDatasets; },
/* harmony export */   delDataset: function() { return /* binding */ delDataset; },
/* harmony export */   delRecommendDatasets: function() { return /* binding */ delRecommendDatasets; },
/* harmony export */   delSingleDataset: function() { return /* binding */ delSingleDataset; },
/* harmony export */   deleteCollaborate: function() { return /* binding */ deleteCollaborate; },
/* harmony export */   deleteTeam: function() { return /* binding */ deleteTeam; },
/* harmony export */   editDatasetsDetail: function() { return /* binding */ editDatasetsDetail; },
/* harmony export */   exportExistDataset: function() { return /* binding */ exportExistDataset; },
/* harmony export */   getAdminDatasets: function() { return /* binding */ getAdminDatasets; },
/* harmony export */   getAvailableUsers: function() { return /* binding */ getAvailableUsers; },
/* harmony export */   getCheckoldDataset: function() { return /* binding */ getCheckoldDataset; },
/* harmony export */   getChunks: function() { return /* binding */ getChunks; },
/* harmony export */   getCollaborate: function() { return /* binding */ getCollaborate; },
/* harmony export */   getCurrentRepoDataset: function() { return /* binding */ getCurrentRepoDataset; },
/* harmony export */   getDatasetReadMe: function() { return /* binding */ getDatasetReadMe; },
/* harmony export */   getDatasets: function() { return /* binding */ getDatasets; },
/* harmony export */   getDatasetsDetail: function() { return /* binding */ getDatasetsDetail; },
/* harmony export */   getFileList: function() { return /* binding */ getFileList; },
/* harmony export */   getFilePreview: function() { return /* binding */ getFilePreview; },
/* harmony export */   getFileSdkCode: function() { return /* binding */ getFileSdkCode; },
/* harmony export */   getModelFile: function() { return /* binding */ getModelFile; },
/* harmony export */   getMultipartUrl: function() { return /* binding */ getMultipartUrl; },
/* harmony export */   getMyFavoriteDataset: function() { return /* binding */ getMyFavoriteDataset; },
/* harmony export */   getMyUploadedDataset: function() { return /* binding */ getMyUploadedDataset; },
/* harmony export */   getNewMultipart: function() { return /* binding */ getNewMultipart; },
/* harmony export */   getProfileDataset: function() { return /* binding */ getProfileDataset; },
/* harmony export */   getProfileDatasetPublic: function() { return /* binding */ getProfileDatasetPublic; },
/* harmony export */   getPromoteDataset: function() { return /* binding */ getPromoteDataset; },
/* harmony export */   getPulicDataset: function() { return /* binding */ getPulicDataset; },
/* harmony export */   getTeams: function() { return /* binding */ getTeams; },
/* harmony export */   getUsers: function() { return /* binding */ getUsers; },
/* harmony export */   modifyCollaborateAcess: function() { return /* binding */ modifyCollaborateAcess; },
/* harmony export */   postCollaborate: function() { return /* binding */ postCollaborate; },
/* harmony export */   postDatasetReadMe: function() { return /* binding */ postDatasetReadMe; },
/* harmony export */   postTeam: function() { return /* binding */ postTeam; },
/* harmony export */   putDatasetStar: function() { return /* binding */ putDatasetStar; },
/* harmony export */   putRecommendDatasets: function() { return /* binding */ putRecommendDatasets; },
/* harmony export */   setCompleteMultipart: function() { return /* binding */ setCompleteMultipart; },
/* harmony export */   setFavorite: function() { return /* binding */ setFavorite; },
/* harmony export */   unsetFavorite: function() { return /* binding */ unsetFavorite; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @babel/runtime/helpers/objectWithoutProperties */ "./node_modules/@babel/runtime/helpers/objectWithoutProperties.js");
/* harmony import */ var _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var _service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../service */ "./web_src/vuepages/apis/service.js");
/* harmony import */ var qs__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! qs */ "./node_modules/qs/lib/index.js");
/* harmony import */ var qs__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(qs__WEBPACK_IMPORTED_MODULE_3__);



 //创建数据集

var createDatasets = function createDatasets(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '/api/v1/dataset/create',
    method: "post",
    params: {},
    data: data
  });
}; //获取管理后台数据集列表

var getAdminDatasets = function getAdminDatasets(params) {
  var type = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 'dataset';
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/admin/".concat(type, "/list"),
    method: "get",
    params: params
  });
}; // 推荐数据集

var putRecommendDatasets = function putRecommendDatasets(params) {
  var type = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 'dataset';
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/admin/".concat(type, "/recommend"),
    method: "put",
    params: params
  });
}; // 取消推荐数据集

var delRecommendDatasets = function delRecommendDatasets(params) {
  var type = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 'dataset';
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/admin/".concat(type, "/recommend"),
    method: "delete",
    params: params
  });
}; //获取数据集列表

var getDatasets = function getDatasets(_ref) {
  var url = _ref.url,
      params = _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_1___default()(_ref, ["url"]);

  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/dataset/list".concat(url),
    method: "get",
    params: params
  });
}; // 获取数据集详情

var getDatasetsDetail = function getDatasetsDetail(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '/api/v1/dataset',
    method: "get",
    params: params
  });
}; // 删除数据集

var delDataset = function delDataset(params) {
  var type = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 'dataset';
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(type),
    method: "delete",
    params: params
  });
}; // 删除单个数据集

var delSingleDataset = function delSingleDataset(params, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(type, "/file"),
    method: "delete",
    params: params
  });
}; // 编辑数据集基本信息

var editDatasetsDetail = function editDatasetsDetail(params, data, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(type, "/edit"),
    method: "post",
    params: params,
    data: data
  });
}; //获取readme

var getDatasetReadMe = function getDatasetReadMe(params, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(type, "/readme"),
    method: "get",
    params: params
  });
}; //创建readme

var postDatasetReadMe = function postDatasetReadMe(params, data, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(type, "/readme"),
    method: "post",
    params: params,
    data: data
  });
}; //获取可创建数据集的用户/组织列表

var getAvailableUsers = function getAvailableUsers(params) {
  var type = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 'dataset';
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(type, "/create/available_users"),
    method: "get",
    params: params
  });
}; //查询用户表

var getUsers = function getUsers(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '/api/v1/users/search',
    method: "get",
    params: params
  });
}; //查询团队表

var getTeams = function getTeams(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/orgs/".concat(params.org, "/teams/search"),
    method: "get",
    params: params
  });
}; //数据集添加协作者

var postCollaborate = function postCollaborate(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '/api/v1/access/collaboration',
    method: "post",
    params: data,
    data: {}
  });
}; //数据集删除协作者

var deleteCollaborate = function deleteCollaborate(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '/api/v1/access/collaboration',
    method: "delete",
    params: data
  });
}; //数据集添加团队

var postTeam = function postTeam(data, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/access/collaboration/".concat(type, "/team"),
    method: "post",
    params: data,
    data: {}
  });
}; //数据集添加团队

var deleteTeam = function deleteTeam(data, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/access/collaboration/".concat(type, "/team/delete"),
    method: "post",
    params: data,
    data: {}
  });
}; //修改协作者权限

var modifyCollaborateAcess = function modifyCollaborateAcess(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '/api/v1/access/collaboration/access_mode',
    method: "post",
    params: data,
    data: {}
  });
}; //数据集已有协作者列表包括组织

var getCollaborate = function getCollaborate(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '/api/v1/access/collaboration',
    method: "get",
    params: params
  });
}; //数据集预览 dataset_id,parent_dir,file_name

var getFilePreview = function getFilePreview(params, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(type, "/preview"),
    method: "get",
    params: params
  });
};
var putDatasetStar = function putDatasetStar(url) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: url,
    method: "put"
  });
};
var getModelFile = function getModelFile(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '',
    params: params,
    method: "get"
  });
};
var exportExistDataset = function exportExistDataset(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '',
    data: data,
    method: "post"
  });
};
/* 选择数据集组件相关 */
// 获取当前仓库的数据集
// params - username, reponame, q, page, type

var getCurrentRepoDataset = function getCurrentRepoDataset(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/".concat(params.userName, "/").concat(params.repoName, "/datasets/current_repo_m"),
    method: 'get',
    params: {
      q: params.q || '',
      page: params.page || 1,
      type: params.type == undefined ? '-1' : params.type
    }
  });
}; // 获取我上传的数据集
// params - username, reponame, q, page, type

var getMyUploadedDataset = function getMyUploadedDataset(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/".concat(params.userName, "/").concat(params.repoName, "/datasets/my_datasets_m"),
    method: 'get',
    params: {
      q: params.q || '',
      page: params.page || 1,
      type: params.type == undefined ? '-1' : params.type
    }
  });
}; // 获取公开数据集
// params - username, reponame, q, page, type

var getPulicDataset = function getPulicDataset(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/".concat(params.userName, "/").concat(params.repoName, "/datasets/public_datasets_m"),
    method: 'get',
    params: {
      q: params.q || '',
      page: params.page || 1,
      type: params.type == undefined ? '-1' : params.type
    }
  });
}; // 获取我收藏的数据集
// params - username, reponame, q, page, type

var getMyFavoriteDataset = function getMyFavoriteDataset(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/".concat(params.userName, "/").concat(params.repoName, "/datasets/my_favorite_m"),
    method: 'get',
    params: {
      q: params.q || '',
      page: params.page || 1,
      type: params.type == undefined ? '-1' : params.type
    }
  });
};
/* 数据集文件上传 */
// 上传文件1: 获取文件chunks信息
// params: { md5, file_name, subject_id,  subject_type: 1, }
// return: uuid, uploaded, chunks, subjectId, fileName

var getChunks = function getChunks(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/upload/get_chunks",
    method: 'get',
    params: params,
    data: {}
  });
}; // 上传文件2: 上传新文件
// params: { total_chunk_counts, md5, size, file_type, subject_type, file_name, subject_id }
// return: uuid

var getNewMultipart = function getNewMultipart(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/upload/new_multipart",
    method: 'get',
    params: params,
    data: {}
  });
}; // 上传文件3: 获取分片上传地址
// params: { uuid, size, chunk_number }
// return: url

var getMultipartUrl = function getMultipartUrl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/upload/get_multipart_url",
    method: 'get',
    params: params,
    data: {}
  });
}; // 上传文件4: 完成上传后
// data: { uuid }
// return "code": 0

var setCompleteMultipart = function setCompleteMultipart(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/upload/complete_multipart",
    method: 'post',
    headers: {
      'Content-type': 'application/x-www-form-urlencoded'
    },
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_3___default().stringify(data)
  });
}; // 数据集广场上方推荐数据集信息
// 数据集广场左侧分类、研究方向、证书等标签的源信息

var getPromoteDataset = function getPromoteDataset(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/dync_config",
    method: 'get',
    params: params,
    data: {}
  });
}; //收藏数据集 
// params: {dataset_id:数据集id}

var setFavorite = function setFavorite(data, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(type, "/collect"),
    method: "put",
    params: data
  });
}; //取消收藏数据集 
// params: {dataset_id:数据集id}

var unsetFavorite = function unsetFavorite(data, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(type, "/collect"),
    method: "delete",
    params: data
  });
}; //获取数据集文件列表
// params: {dataset_id:数据集id, parent_dir:文件路径, page_size: 10, marker: ''}

var getFileList = function getFileList(params, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(type, "/files"),
    method: "get",
    params: params
  });
}; //弹窗sdk下载代码

var getFileSdkCode = function getFileSdkCode(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: "/api/v1/".concat(params.type, "/sdk_download_code"),
    method: "get",
    params: params
  });
}; //弹窗sdk下载代码

var getCheckoldDataset = function getCheckoldDataset(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '/api/v1/dataset/check_old',
    method: "get",
    params: params
  });
}; //获取个人信息页面中的数据集列表

var getProfileDataset = function getProfileDataset(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '/api/v1/dataset/list/accessible',
    method: "get",
    params: params
  });
}; //获取个人信息页面中的数据集列表（未登录态使用，仅返回公开数据集）

var getProfileDatasetPublic = function getProfileDatasetPublic(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_2__["default"])({
    url: '/api/v1/dataset/list/public',
    method: "get",
    params: params
  });
};

/***/ })

}]);