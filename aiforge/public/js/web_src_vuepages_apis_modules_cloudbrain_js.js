"use strict";
(self["webpackChunk"] = self["webpackChunk"] || []).push([["web_src_vuepages_apis_modules_cloudbrain_js"],{

/***/ "./web_src/vuepages/apis/modules/cloudbrain.js":
/*!*****************************************************!*\
  !*** ./web_src/vuepages/apis/modules/cloudbrain.js ***!
  \*****************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   createAiTask: function() { return /* binding */ createAiTask; },
/* harmony export */   deleteAiTask: function() { return /* binding */ deleteAiTask; },
/* harmony export */   deleteMulAiTask: function() { return /* binding */ deleteMulAiTask; },
/* harmony export */   getAdminAiTasks: function() { return /* binding */ getAdminAiTasks; },
/* harmony export */   getAiEndpointUrl: function() { return /* binding */ getAiEndpointUrl; },
/* harmony export */   getAiEvalDetailResult: function() { return /* binding */ getAiEvalDetailResult; },
/* harmony export */   getAiEvalResult: function() { return /* binding */ getAiEvalResult; },
/* harmony export */   getAiTask: function() { return /* binding */ getAiTask; },
/* harmony export */   getAiTaskBrief: function() { return /* binding */ getAiTaskBrief; },
/* harmony export */   getAiTaskDebugUrl: function() { return /* binding */ getAiTaskDebugUrl; },
/* harmony export */   getAiTaskExportDatasetProgress: function() { return /* binding */ getAiTaskExportDatasetProgress; },
/* harmony export */   getAiTaskExportModelProgress: function() { return /* binding */ getAiTaskExportModelProgress; },
/* harmony export */   getAiTaskImgesBySpec: function() { return /* binding */ getAiTaskImgesBySpec; },
/* harmony export */   getAiTaskList: function() { return /* binding */ getAiTaskList; },
/* harmony export */   getAiTaskLogs: function() { return /* binding */ getAiTaskLogs; },
/* harmony export */   getAiTaskLogsDownloadUrl: function() { return /* binding */ getAiTaskLogsDownloadUrl; },
/* harmony export */   getAiTaskLoss: function() { return /* binding */ getAiTaskLoss; },
/* harmony export */   getAiTaskNodeInfo: function() { return /* binding */ getAiTaskNodeInfo; },
/* harmony export */   getAiTaskOperationProfile: function() { return /* binding */ getAiTaskOperationProfile; },
/* harmony export */   getAiTaskOutputResult: function() { return /* binding */ getAiTaskOutputResult; },
/* harmony export */   getAiTaskOutputResultAll: function() { return /* binding */ getAiTaskOutputResultAll; },
/* harmony export */   getAiTaskPrepareInfo: function() { return /* binding */ getAiTaskPrepareInfo; },
/* harmony export */   getAiTaskResourceUseage: function() { return /* binding */ getAiTaskResourceUseage; },
/* harmony export */   getAiTaskRestart: function() { return /* binding */ getAiTaskRestart; },
/* harmony export */   getAiTaskVisualizeUrl: function() { return /* binding */ getAiTaskVisualizeUrl; },
/* harmony export */   getAimRight: function() { return /* binding */ getAimRight; },
/* harmony export */   getAimUrl: function() { return /* binding */ getAimUrl; },
/* harmony export */   getCanCreateTask: function() { return /* binding */ getCanCreateTask; },
/* harmony export */   getDownLoadAiTaskResultFileAllUrl: function() { return /* binding */ getDownLoadAiTaskResultFileAllUrl; },
/* harmony export */   getDownLoadAiTaskResultFileUrl: function() { return /* binding */ getDownLoadAiTaskResultFileUrl; },
/* harmony export */   getMyAiTasks: function() { return /* binding */ getMyAiTasks; },
/* harmony export */   getRepoDatasetInfo: function() { return /* binding */ getRepoDatasetInfo; },
/* harmony export */   getTmplEditAdress: function() { return /* binding */ getTmplEditAdress; },
/* harmony export */   setAiTaskExportDataset: function() { return /* binding */ setAiTaskExportDataset; },
/* harmony export */   setAiTaskExportDataset1: function() { return /* binding */ setAiTaskExportDataset1; },
/* harmony export */   setAiTaskExportModel: function() { return /* binding */ setAiTaskExportModel; },
/* harmony export */   setAiTaskOutputReschedule: function() { return /* binding */ setAiTaskOutputReschedule; },
/* harmony export */   setAiTaskResultToModel: function() { return /* binding */ setAiTaskResultToModel; },
/* harmony export */   setAiTaskResultToModelApi: function() { return /* binding */ setAiTaskResultToModelApi; },
/* harmony export */   stopAiTask: function() { return /* binding */ stopAiTask; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.object.define-properties */ "./node_modules/core-js/modules/es.object.define-properties.js");
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.object.define-property */ "./node_modules/core-js/modules/es.object.define-property.js");
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptor */ "./node_modules/core-js/modules/es.object.get-own-property-descriptor.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptors */ "./node_modules/core-js/modules/es.object.get-own-property-descriptors.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var _service__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! ../service */ "./web_src/vuepages/apis/service.js");
/* harmony import */ var qs__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! qs */ "./node_modules/qs/lib/index.js");
/* harmony import */ var qs__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(qs__WEBPACK_IMPORTED_MODULE_12__);












function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_10___default()(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }




var _csrf = window.config ? window.config.csrf : ''; // 查询创建任务时所需信息


var getAiTaskPrepareInfo = function getAiTaskPrepareInfo(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/creation/required",
    method: 'get',
    params: {
      job_type: params.jobType,
      compute_source: params.computeSource,
      cluster_type: params.clusterType,
      app_name: params.appName
    }
  });
}; // 由资源规格查询可用镜像
// params: repoOwnerName, repoName, job_type,compute_source,cluster_type,spec_id

var getAiTaskImgesBySpec = function getAiTaskImgesBySpec(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/creation/image_by_spec",
    method: 'get',
    params: {
      job_type: params.jobType,
      compute_source: params.computeSource,
      cluster_type: params.clusterType,
      spec_id: params.spec,
      has_internet: params.hasInternet,
      visualize_required: params.visualizeRequired
    }
  });
}; // 是否能够创建AI任务

var getCanCreateTask = function getCanCreateTask(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/".concat(params.repoOwnerName, "/").concat(params.repoName, "/ai_task/can_create"),
    method: 'get',
    params: {},
    data: {}
  });
}; // 获取AI任务模板编辑页地址

var getTmplEditAdress = function getTmplEditAdress(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/".concat(params.repoOwnerName, "/").concat(params.repoName, "/ai_task/template/get_edit_address").concat(params.branchName ? '/branch/' + params.branchName : ''),
    method: 'get',
    params: {},
    data: {}
  });
}; // 创建AI任务

var createAiTask = function createAiTask(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/".concat(data.repoOwnerName, "/").concat(data.repoName, "/ai_task/create"),
    method: 'post',
    params: {},
    data: _objectSpread({}, data)
  });
}; // 查询AI任务列表
// page,
// job_type-DEBUG|TRAIN|INFERENCE|BENCHMARK,
// compute_source-GPU|NPU|GCU|MLU

var getAiTaskList = function getAiTaskList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/".concat(params.repoOwnerName, "/").concat(params.repoName, "/ai_task/list"),
    method: 'get',
    params: _objectSpread({}, params)
  });
}; // 查看任务详情

var getAiTask = function getAiTask(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task",
    method: 'get',
    params: {
      id: params.id
    }
  });
}; // 查询任务最新状态

var getAiTaskBrief = function getAiTaskBrief(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/brief",
    method: 'get',
    params: {
      id: params.id
    }
  });
}; // 获取调试页面地址

var getAiTaskDebugUrl = function getAiTaskDebugUrl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/debug_url",
    method: 'get',
    params: {
      id: params.id,
      file: params.file
    }
  });
}; // 获取可视化页面地址

var getAiTaskVisualizeUrl = function getAiTaskVisualizeUrl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/visualize_url",
    method: 'get',
    params: {
      id: params.id
    }
  });
}; // 获取自定义路径页面地址

var getAiEndpointUrl = function getAiEndpointUrl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/self_endpoint_url",
    method: 'get',
    params: {
      id: params.id,
      file: params.file
    }
  });
}; // 再次调试

var getAiTaskRestart = function getAiTaskRestart(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/restart?id=".concat(data.id),
    method: 'post',
    params: {},
    data: {}
  });
}; // 停止任务

var stopAiTask = function stopAiTask(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/stop?id=".concat(data.id),
    method: 'post',
    params: {},
    data: {}
  });
}; // 删除任务

var deleteAiTask = function deleteAiTask(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/del?id=".concat(data.id),
    method: 'post',
    params: {},
    data: {}
  });
}; // 批量删除任务

var deleteMulAiTask = function deleteMulAiTask(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: '/api/v1/ai_task/batch_del',
    method: 'post',
    params: {},
    data: data
  });
}; // 查询任务运行简况

var getAiTaskOperationProfile = function getAiTaskOperationProfile(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/operation_profile",
    method: 'get',
    params: {
      id: params.id
    }
  });
}; // 查询任务loss曲线

var getAiTaskLoss = function getAiTaskLoss(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/loss",
    method: 'get',
    params: {
      id: params.id
    }
  });
}; // 查询任务日志
// params: id, base_line, lines, order-asc|desc,node_id,file_name

var getAiTaskLogs = function getAiTaskLogs(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/log",
    method: 'get',
    params: {
      id: params.id,
      base_line: params.base_line,
      lines: params.lines,
      order: params.order,
      node_id: params.node_id,
      file_name: params.log_file_name
    }
  });
}; // 下载任务日志

var getAiTaskLogsDownloadUrl = function getAiTaskLogsDownloadUrl(params) {
  return "/api/v1/ai_task/log/download?id=".concat(params.id, "&node_id=").concat(params.node_id, "&file_name=").concat(params.log_file_name, "&_csrf=").concat(_csrf);
}; // 查询任务资源占用情况
// params: id,node_id,log_file_name

var getAiTaskResourceUseage = function getAiTaskResourceUseage(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/resource_usage",
    method: 'get',
    params: {
      id: params.id,
      node_id: params.node_id,
      file_name: params.log_file_name
    }
  });
}; // 查询任务结果下载
// params: id, parent_dir

var getAiTaskOutputResult = function getAiTaskOutputResult(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/output",
    method: 'get',
    params: {
      id: params.id,
      parent_dir: params.parent_dir
    }
  });
}; // 查询任务结果下载(所有)
// params: id,suffix-.zip|.tar.gz

var getAiTaskOutputResultAll = function getAiTaskOutputResultAll(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/output/all",
    method: 'get',
    params: {
      id: params.id,
      suffix: params.suffix
    }
  });
}; // 下载训练任务的某一个结果文件
// params-id,file_name,parent_dir

var getDownLoadAiTaskResultFileUrl = function getDownLoadAiTaskResultFileUrl(params) {
  return "/api/v1/ai_task/output/download?id=".concat(params.id, "&file_name=").concat(params.file_name, "&parent_dir=").concat(params.parent_dir, "&_csrf=").concat(_csrf);
}; // 下载训练任务的所有结果文件链接
// params-id

var getDownLoadAiTaskResultFileAllUrl = function getDownLoadAiTaskResultFileAllUrl(params) {
  return "/api/v1/ai_task/output/download/all?id=".concat(params.id, "&_csrf=").concat(_csrf);
}; // 结果下载重新获取

var setAiTaskOutputReschedule = function setAiTaskOutputReschedule(data) {
  var paramsObj = {
    id: data.id
  };
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/output/reschedule",
    method: 'post',
    data: {
      id: data.id
    },
    params: paramsObj
  });
}; // 获取训练任务的节点信息

var getAiTaskNodeInfo = function getAiTaskNodeInfo(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/node_info",
    method: 'get',
    params: {
      id: params.id
    }
  });
}; // 导出训练任务的输出至模型
// trainTaskCreate-true,cloudbrain_id(jobId,versionName),name,version,engine,engine_name,modelSelectedFile,label,isPrivate,description

var setAiTaskResultToModel = function setAiTaskResultToModel(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/".concat(data.repoOwnerName, "/").concat(data.repoName, "/modelmanage/create_new_model"),
    method: 'post',
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_12___default().stringify(data)
  });
}; // 导出训练任务的输出至模型
// trainTaskCreate-true,cloudbrain_id(jobId,versionName),name,version,engine,engine_name,modelSelectedFile,label,isPrivate,description

var setAiTaskResultToModelApi = function setAiTaskResultToModelApi(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/repos/".concat(data.repoOwnerName, "/").concat(data.repoName, "/modelmanage/create_new_model"),
    method: 'post',
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_12___default().stringify(data)
  });
};
/* 导出训练任务的输出至数据集相关 */
// 获取项目仓库的数据集信息

var getRepoDatasetInfo = function getRepoDatasetInfo(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/".concat(params.repoOwnerName, "/").concat(params.repoName, "/datasets/model/getcurrentdataset"),
    method: 'get',
    params: {}
  });
}; // 获取AI任务结果导出到数据集的导出进度信息

var getAiTaskExportDatasetProgress = function getAiTaskExportDatasetProgress(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/dataset/get_export_process",
    method: 'get',
    params: params
  });
}; // AI任务结果中导入数据集提交
// cloudbrain_id(jobId,versionName),datasetId,modelSelectedFile,type,description,

var setAiTaskExportDataset = function setAiTaskExportDataset(params, data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/export2dataset",
    method: 'post',
    params: params,
    data: data
  });
}; // AI任务结果中导入模型提交
// cloudbrain_id(jobId,versionName),datasetId,modelSelectedFile,type,description,

var setAiTaskExportModel = function setAiTaskExportModel(params, data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/export2aimodel",
    method: 'post',
    params: params,
    data: data
  });
}; // 获取AI任务结果导出到模型的导出进度信息

var getAiTaskExportModelProgress = function getAiTaskExportModelProgress(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: '/api/v1/aimodel/get_export_process',
    method: 'get',
    params: {
      id: params.id
    }
  });
}; // AI任务结果中导入数据集提交
// cloudbrain_id(jobId,versionName),datasetId,modelSelectedFile,type,description,

var setAiTaskExportDataset1 = function setAiTaskExportDataset1(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/".concat(data.repoOwnerName, "/").concat(data.repoName, "/datasets/model/export_exist_dataset"),
    method: 'post',
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_12___default().stringify(data)
  });
}; // 我的工作台-云脑任务列表
// job_type,job_status,ai_center,cluster,compute_source,q,page,pageSize

var getMyAiTasks = function getMyAiTasks(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/my_list",
    method: 'get',
    params: _objectSpread({}, params)
  });
}; // 管理后台-云脑任务列表
// job_type,job_status,ai_center,cluster,compute_source,q,page,pageSize

var getAdminAiTasks = function getAdminAiTasks(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/admin/ai_task/list",
    method: 'get',
    params: _objectSpread({}, params)
  });
};
/* Aim可视化相关 */
// 获取Aim功能权限状态

var getAimRight = function getAimRight() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/monitor/aim_right",
    method: 'get',
    params: {}
  });
}; // 获取Aim跳转地址
// job_name-["jobname1","jobname2"]

var getAimUrl = function getAimUrl(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/monitor/aim_url",
    method: 'post',
    data: {
      job_name: data.job_name
    }
  });
}; // 查询任务评测总览

var getAiEvalResult = function getAiEvalResult(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/eval_result",
    method: 'get',
    params: {
      id: params.id
    }
  });
}; // 查询任务评测详情

var getAiEvalDetailResult = function getAiEvalDetailResult(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/ai_task/eval_detail_result",
    method: 'post',
    params: {
      id: data.task_id
    },
    data: data
  });
};

/***/ })

}]);