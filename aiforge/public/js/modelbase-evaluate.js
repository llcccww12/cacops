"use strict";
(self["webpackChunk"] = self["webpackChunk"] || []).push([["modelbase-evaluate"],{

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=script&lang=js":
/*!*********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=script&lang=js ***!
  \*********************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.date.now */ "./node_modules/core-js/modules/es.date.now.js");
/* harmony import */ var core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.object.assign */ "./node_modules/core-js/modules/es.object.assign.js");
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var dayjs__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! dayjs */ "./node_modules/dayjs/dayjs.min.js");
/* harmony import */ var dayjs__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(dayjs__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var _components_Header_vue__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ../../components/Header.vue */ "./web_src/vuepages/pages/modelbase-portal/components/Header.vue");
/* harmony import */ var _pages_cloudbrain_tools__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ~/pages/cloudbrain/tools */ "./web_src/vuepages/pages/cloudbrain/tools.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _const__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! ~/const */ "./web_src/vuepages/const/index.js");
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");






//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//






var cloudBrainTools = new _pages_cloudbrain_tools__WEBPACK_IMPORTED_MODULE_8__.CloudBrainTools();
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'Experience',
  data: function data() {
    return {
      loading: false,
      tableData: [],
      page: 1,
      pageSize: 10,
      total: 0,
      pageSizes: [10, 20, 30, 50]
    };
  },
  components: {
    Header: _components_Header_vue__WEBPACK_IMPORTED_MODULE_7__["default"]
  },
  methods: {
    getTableData: function getTableData() {
      var _this = this;

      var params = {
        job_type: 'EVAL',
        job_status: '',
        ai_center: '',
        cluster: '',
        compute_source: '',
        q: '',
        page: this.page,
        pageSize: this.pageSize,
        exclude_status: undefined
      };
      var routerBase = this.$router.options.base;
      this.loading = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_11__.getMyAiTasks)(params).then(function (res) {
        _this.loading = false;
        res = res.data;

        if (res.code == 0) {
          var data = res.data;
          _this.canCreateTask = data.can_create_task;
          _this.isRepoEmpty = data.is_repo_empty;
          _this.total = data.total;
          data.tasks.forEach(function (item) {
            var obj = Object.assign({}, item);
            delete obj.task;
            var task = item.task;
            Object.assign(task, obj);
            task.computeSourceShow = task.compute_source == 'GPU' ? 'CPU/GPU' : task.compute_source;
            task.clusterName = (0,_utils__WEBPACK_IMPORTED_MODULE_9__.getListValueWithKey)(_const__WEBPACK_IMPORTED_MODULE_10__.CLUSTERS, task.cluster);
            task.accCardTypeShow = (0,_utils__WEBPACK_IMPORTED_MODULE_9__.getListValueWithKey)(_const__WEBPACK_IMPORTED_MODULE_10__.ACC_CARD_TYPE, task.acc_card_type);
            task.repoOwnerName = item.owner_name;
            task.repoName = item.repo_name;
            task.createdFromNow = _this.calcFromNow(task.created_unix);
            task.jobNameShow = task.job_name;
            task.jobTypeShow = (0,_utils__WEBPACK_IMPORTED_MODULE_9__.getListValueWithKey)(_const__WEBPACK_IMPORTED_MODULE_10__.JOB_TYPE, task.job_type);
            cloudBrainTools.checkOperation(task);
          });
          _this.tableData = data.tasks || [];
          cloudBrainTools.initRefreshData(_this.tableData);
        }
      })["catch"](function (err) {
        _this.loading = false;
        console.log(err);
      });
    },
    currentChange: function currentChange(page) {
      this.page = page;
      this.getTableData();
    },
    sizeChange: function sizeChange(pageSize) {
      this.page = 1;
      this.pageSize = pageSize;
      this.getTableData();
    },
    calcFromNow: function calcFromNow(unix) {
      return (0,_utils__WEBPACK_IMPORTED_MODULE_9__.timeSinceUnix)(unix, Date.now() / 1000);
    },
    dateFormat: function dateFormat(unix) {
      return dayjs__WEBPACK_IMPORTED_MODULE_6___default()(unix * 1000).format('YYYY-MM-DD HH:mm:ss');
    },
    goCreate: function goCreate() {
      this.$router.push(this.$route.path + '/create');
    },
    goDetail: function goDetail(row) {
      this.$router.push("/eval/evaluate/detail/".concat(row.task.id));
    },
    opStop: function opStop(row) {
      var _this2 = this;

      if (this.operating) return;
      this.operating = true;
      var experienceParams = {};
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_11__.stopAiTask)({
        id: row.task.id
      }).then(function (res) {
        _this2.operating = false;
        res = res.data;

        if (res.code == 0) {
          var data = res.data;
          Object.assign(row.task, data);
          row.task.createdFromNow = _this2.calcFromNow(row.task.created_unix);
          cloudBrainTools.checkOperation(row.task);
        } else {
          _this2.$message({
            type: 'error',
            message: res.msg
          });
        }
      })["catch"](function (err) {
        _this2.operating = false;

        _this2.$message({
          type: 'error',
          message: _this2.$t('operationFailed')
        });
      });
    },
    opDelete: function opDelete(row) {
      var _this3 = this;

      if (this.operating) return;
      this.$confirm(this.$t('cloudbrainObj.deleteConfirmTips'), this.$t('tips'), {
        confirmButtonText: this.$t('confirm'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false
      }).then(function () {
        _this3.operating = true;
        _this3.maskLoading = true;
        _this3.maskLoadingContent = _this3.$t('cloudbrainObj.deletingTips');
        (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_11__.deleteAiTask)({
          repoOwnerName: row.task.repoOwnerName,
          repoName: row.task.repoName,
          id: row.task.id
        }).then(function (res) {
          _this3.operating = false;
          _this3.maskLoading = false;

          if (_this3.total % _this3.pageSize == 1 && _this3.page > 1) {
            _this3.page -= 1;
          }

          _this3.getTableData();
        })["catch"](function (err) {
          _this3.maskLoading = false;
          _this3.operating = false;

          _this3.$message({
            type: 'error',
            message: _this3.$t('operationFailed')
          });
        });
      })["catch"](function () {
        _this3.$message({
          type: 'info',
          message: _this3.$t('cancelOperate')
        });
      });
    },
    opOnline: function opOnline(row) {
      var url = '';
      var task = row.task;

      if (task.label_name == 'chat') {
        url = "/extension/modelexperience/chat?id=".concat(btoa(task.id), "&modelName=").concat(task.app_name);
      }

      if (row.task.label_name == 'sd') {
        url = "/extension/modelexperience/sd?id=".concat(btoa(task.id), "&modelName=").concat(task.app_name);
      }

      if (row.task.label_name == 'tts') {
        url = "/extension/modelexperience/tts?id=".concat(btoa(task.id), "&modelName=").concat(task.app_name);
      } // 创建隐藏的<a>标签并模拟点击


      var link = document.createElement('a');
      link.href = url;
      link.target = '_blank'; // 确保在新标签页打开

      link.rel = 'noopener noreferrer'; // 安全设置

      link.style.display = 'none'; // 隐藏元素

      document.body.appendChild(link); // 添加到DOM

      link.click(); // 触发点击

      document.body.removeChild(link); // 移除元素
    }
  },
  beforeMount: function beforeMount() {
    this.getTableData();
  },
  mounted: function mounted() {}
});

/***/ }),

/***/ "./web_src/vuepages/pages/cloudbrain/tools.js":
/*!****************************************************!*\
  !*** ./web_src/vuepages/pages/cloudbrain/tools.js ***!
  \****************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   CloudBrainTools: function() { return /* binding */ CloudBrainTools; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.includes */ "./node_modules/core-js/modules/es.array.includes.js");
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.date.now */ "./node_modules/core-js/modules/es.date.now.js");
/* harmony import */ var core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.object.assign */ "./node_modules/core-js/modules/es.object.assign.js");
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! @babel/runtime/helpers/classCallCheck */ "./node_modules/@babel/runtime/helpers/classCallCheck.js");
/* harmony import */ var _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! @babel/runtime/helpers/createClass */ "./node_modules/@babel/runtime/helpers/createClass.js");
/* harmony import */ var _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");














var CloudBrainTools = /*#__PURE__*/function () {
  function CloudBrainTools() {
    _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_10___default()(this, CloudBrainTools);

    this.intervalTimer = null;
    this.refreshIntervalTime = 1000 * 8;
    this.list = [];
  }

  _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_11___default()(CloudBrainTools, [{
    key: "initRefreshData",
    value: function initRefreshData(data) {
      var _this = this;

      this.stop();
      this.list = data;
      this.refreshStatus();
      this.intervalTimer = setInterval(function () {
        _this.refreshStatus();
      }, this.refreshIntervalTime);
    }
  }, {
    key: "refreshStatus",
    value: function refreshStatus() {
      var _this2 = this;

      var _loop = function _loop(i, iLen) {
        var taskInfo = _this2.list[i];
        var task = taskInfo.task;
        task.createdFromNow = (0,_utils__WEBPACK_IMPORTED_MODULE_13__.timeSinceUnix)(task.created_unix, Date.now() / 1000);
        var finalState = ["STOPPED", "CREATE_FAILED", "CREATED_FAILED", "UNAVAILABLE", "DELETED", "RESIZE_FAILED", "SUCCEEDED", "IMAGE_FAILED", "SUBMIT_FAILED", "DELETE_FAILED", "KILLED", "COMPLETED", "FAILED", "CANCELED", "LOST", "START_FAILED", "SUBMIT_MODEL_FAILED", "DEPLOY_SERVICE_FAILED", "CHECK_FAILED"];

        if (finalState.includes(task.status)) {
          return "continue";
        }

        var self = _this2;

        (function (taskInfo) {
          var task = taskInfo.task;
          (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_12__.getAiTaskBrief)({
            repoOwnerName: task.repoOwnerName,
            repoName: task.repoName,
            id: task.id
          }).then(function (res) {
            res = res.data;

            if (res.code == 0) {
              var data = res.data;
              Object.assign(task, data);
              task.createdFromNow = (0,_utils__WEBPACK_IMPORTED_MODULE_13__.timeSinceUnix)(task.created_unix, Date.now() / 1000);
              self.checkRunningLeftTime(task);
              self.checkOperation(task);
            }
          })["catch"](function (err) {
            console.log(err);
          });
        })(taskInfo);
      };

      for (var i = 0, iLen = this.list.length; i < iLen; i++) {
        var _ret = _loop(i, iLen);

        if (_ret === "continue") continue;
      }
    }
  }, {
    key: "stop",
    value: function stop() {
      clearInterval(this.intervalTimer);
    }
  }, {
    key: "checkOperation",
    value: function checkOperation(task) {
      if (task.status == 'RUNNING') {
        task.canDebug = true;
        task.canSaveImage = true;
      } else {
        task.canDebug = false;
        task.canSaveImage = false;
      }

      if (task.job_type == 'TRAIN' && task.status == 'RUNNING' && task.visualize_required) {
        task.canVisualize = true;
      } else {
        task.canVisualize = false;
      }

      if (task.aim_required && ["RUNNING", "STOPPED", "FAILED", "SUCCEEDED", "COMPLETED"].includes(task.status)) {
        task.canAim = true;
      } else {
        task.canAim = false;
      }

      if (["PREPARING", "CONNECTING", "CREATING", "STOPPING", "WAITING", "STARTING"].includes(task.status)) {
        task.canDebug = false;
      }

      if (["STOPPED", "FAILED", "START_FAILED", "CREATE_FAILED", "SUCCEEDED"].includes(task.status)) {
        if (task.is_file_notebook) {
          task.canReDebug = false;
        } else {
          task.canReDebug = true;
        }
      } else {
        task.canReDebug = false;
      }

      if (["INIT", "RUNNING", "WAITING"].includes(task.status)) {
        task.canStop = true;
      } else {
        task.canStop = false;
      }

      if (["STOPPED", "FAILED", "START_FAILED", "KILLED", "COMPLETED", "SUCCEEDED", "CREATE_FAILED", "CREATED_FAILED"].includes(task.status)) {
        task.canDelete = true;
      } else {
        task.canDelete = false;
      }

      task.hasMore = true;
      task.canSaveTmpl = ['DEBUG', 'TRAIN', 'ONLINEINFERENCE', 'GENERAL', 'HPC'].includes(task.job_type);

      if ((['GPU', 'NPU', 'GCU', 'MLU', 'ILUVATAR-GPGPU', 'METAX-GPGPU', 'BIREN-GPU'].includes(task.compute_source) || task.ai_center_code === 'sugon-ai') && ['DEBUG', 'ONLINEINFERENCE', 'ComfyuiExperience'].includes(task.job_type)) {
        task.hasMore = true;

        if (task.canDebug) {
          task.canSaveImage = true;

          if (task.job_type === 'ComfyuiExperience') {
            task.saveImageUrl = "/notebook/".concat(task.id, "/commit_image?type=comfyui");
          } else {
            task.saveImageUrl = "/notebook/".concat(task.id, "/commit_image");
          }
        } else {
          task.canSaveImage = false;
        }

        if (task.can_modify && task.cluster == 'OpenI' && !['PREPARING', 'CONNECTING'].includes(task.status)) {
          task.canDownloadModel = true;
          task.downloadModelUrl = "/".concat(task.repoOwnerName, "/").concat(task.repoName, "/cloudbrain/").concat(task.id, "/models");
        } else {
          task.canDownloadModel = false;
        }

        if (task.is_file_notebook) {
          task.hasDebugMore = false;
          task.canSaveImage = false;
          task.canDownloadModel = false;
        }
      } else {
        task.hasDebugMore = false;
        task.canSaveImage = false;
        task.canDownloadModel = false;
      }

      task.canModify = true;

      if (task.is_fine_tune_task || task.is_file_notebook) {
        task.canModify = false;
      }

      if ((task.job_type == 'TRAIN' || task.job_type == 'FINETUNE') && task.can_download) {
        task.canExportOutput = true;
      }
    }
  }, {
    key: "checkRunningLeftTime",
    value: function checkRunningLeftTime(task) {
      var timeLimit = task.time_limit > -1 ? task.time_limit > 0 ? task.time_limit : task.default_time_limit : 0;

      if (task.job_type == 'DEBUG' && task.status == 'RUNNING' && timeLimit > 0) {
        var hmsList = task.formatted_duration.split(':').map(function (item) {
          return Number(item);
        });
        var durationSecond = hmsList[0] * 60 * 60 + hmsList[1] * 60 + hmsList[2];
        var allSecond = timeLimit * 60 * 60;
        task.runningLeftTime = Math.max(1, Math.floor((allSecond - durationSecond) / 60));
      } else {
        task.runningLeftTime = undefined;
      }
    } // 优化后的getAiJobLink方法

  }, {
    key: "getAiJobLink",
    value: function getAiJobLink(taskInfo) {
      return '';
    }
  }]);

  return CloudBrainTools;
}();

/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=style&index=0&id=3e9a1786&scoped=true&lang=less":
/*!***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=style&index=0&id=3e9a1786&scoped=true&lang=less ***!
  \***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue":
/*!*****************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue ***!
  \*****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Evaluate_vue_vue_type_template_id_3e9a1786_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Evaluate.vue?vue&type=template&id=3e9a1786&scoped=true */ "./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=template&id=3e9a1786&scoped=true");
/* harmony import */ var _Evaluate_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Evaluate.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=script&lang=js");
/* harmony import */ var _Evaluate_vue_vue_type_style_index_0_id_3e9a1786_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Evaluate.vue?vue&type=style&index=0&id=3e9a1786&scoped=true&lang=less */ "./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=style&index=0&id=3e9a1786&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Evaluate_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Evaluate_vue_vue_type_template_id_3e9a1786_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Evaluate_vue_vue_type_template_id_3e9a1786_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "3e9a1786",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=script&lang=js":
/*!*****************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=script&lang=js ***!
  \*****************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Evaluate_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Evaluate.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Evaluate_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=style&index=0&id=3e9a1786&scoped=true&lang=less":
/*!**************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=style&index=0&id=3e9a1786&scoped=true&lang=less ***!
  \**************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Evaluate_vue_vue_type_style_index_0_id_3e9a1786_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Evaluate.vue?vue&type=style&index=0&id=3e9a1786&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=style&index=0&id=3e9a1786&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=template&id=3e9a1786&scoped=true":
/*!***********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=template&id=3e9a1786&scoped=true ***!
  \***********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Evaluate_vue_vue_type_template_id_3e9a1786_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Evaluate_vue_vue_type_template_id_3e9a1786_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Evaluate_vue_vue_type_template_id_3e9a1786_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Evaluate.vue?vue&type=template&id=3e9a1786&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=template&id=3e9a1786&scoped=true");


/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=template&id=3e9a1786&scoped=true":
/*!**************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelbase-portal/views/evaluate/Evaluate.vue?vue&type=template&id=3e9a1786&scoped=true ***!
  \**************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* binding */ render; },
/* harmony export */   staticRenderFns: function() { return /* binding */ staticRenderFns; }
/* harmony export */ });
var render = function() {
  var _vm = this
  var _h = _vm.$createElement
  var _c = _vm._self._c || _h
  return _c(
    "div",
    { staticClass: "list-wrap" },
    [
      _c(
        "Header",
        [
          _c(
            "template",
            { slot: "right" },
            [
              _c(
                "el-button",
                {
                  staticClass: "op-btn",
                  attrs: { type: "primary" },
                  on: { click: _vm.goCreate }
                },
                [
                  _c("div", { staticClass: "btn-content" }, [
                    _c("i", { staticClass: "ri-add-box-line" }),
                    _vm._v(" "),
                    _c("span", [
                      _vm._v(_vm._s(_vm.$t("modelSquare.createEvalBtns1")))
                    ])
                  ])
                ]
              )
            ],
            1
          )
        ],
        2
      ),
      _vm._v(" "),
      _c("div", {
        staticClass: "tips",
        domProps: { innerHTML: _vm._s(_vm.$t("modelFinetune.maxEvalTaskTips")) }
      }),
      _vm._v(" "),
      _c("div", { staticClass: "main-body" }, [
        _c("div", { staticClass: "content" }, [
          _c("div", { staticClass: "body" }, [
            _c(
              "div",
              { staticClass: "table-container" },
              [
                _c(
                  "el-table",
                  {
                    directives: [
                      {
                        name: "loading",
                        rawName: "v-loading",
                        value: _vm.loading,
                        expression: "loading"
                      }
                    ],
                    staticStyle: { "min-width": "100%" },
                    attrs: { data: _vm.tableData, height: "100%", stripe: "" }
                  },
                  [
                    _c("el-table-column", {
                      attrs: {
                        label: _vm.$t("cloudbrainObj.taskName"),
                        align: "left",
                        "header-align": "left",
                        "min-width": "250"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c(
                                "a",
                                {
                                  staticClass: "dispaly-job-name",
                                  on: {
                                    click: function($event) {
                                      $event.preventDefault()
                                      return _vm.goDetail(scope.row)
                                    }
                                  }
                                },
                                [
                                  _vm._v(
                                    "\n                  " +
                                      _vm._s(scope.row.task.display_job_name) +
                                      "\n                "
                                  )
                                ]
                              )
                            ]
                          }
                        }
                      ])
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        label: _vm.$t("modelSquare.evaluationModel"),
                        align: "left",
                        "header-align": "left",
                        "min-width": "250"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("span", [
                                _vm._v(_vm._s(scope.row.task.app_name))
                              ])
                            ]
                          }
                        }
                      ])
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        prop: "Status",
                        label: _vm.$t("status"),
                        align: "left",
                        "header-align": "center",
                        "min-width": "150"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("div", { staticClass: "status-wrap" }, [
                                _c("i", { class: scope.row.task.status }),
                                _vm._v(" "),
                                _c("span", [
                                  _vm._v(_vm._s(scope.row.task.status))
                                ]),
                                _vm._v(" "),
                                scope.row.task.detailed_status ===
                                  "dataMigrating" &&
                                scope.row.task.status === "WAITING"
                                  ? _c("i", {
                                      class: scope.row.task.detailed_status,
                                      attrs: {
                                        title: _vm.$t(
                                          "cloudbrainObj.migratingData"
                                        )
                                      }
                                    })
                                  : _vm._e(),
                                _vm._v(" "),
                                scope.row.task.detailed_status ===
                                  "centerPending" &&
                                scope.row.task.status === "WAITING"
                                  ? _c("i", {
                                      class: scope.row.task.detailed_status,
                                      attrs: {
                                        title: _vm.$t(
                                          "cloudbrainObj.centerPending"
                                        )
                                      }
                                    })
                                  : _vm._e(),
                                _vm._v(" "),
                                scope.row.task.detailed_status ===
                                  "ImagePulling" &&
                                scope.row.task.status === "WAITING"
                                  ? _c("i", {
                                      class: scope.row.task.detailed_status,
                                      attrs: {
                                        title: _vm.$t(
                                          "cloudbrainObj.imagePulling"
                                        )
                                      }
                                    })
                                  : _vm._e()
                              ])
                            ]
                          }
                        }
                      ])
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        prop: "created_unix",
                        label: _vm.$t("cloudbrainObj.createTime"),
                        align: "center",
                        "header-align": "center",
                        width: "180"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("span", [
                                _vm._v(
                                  _vm._s(
                                    _vm.dateFormat(scope.row.task.created_unix)
                                  )
                                )
                              ])
                            ]
                          }
                        }
                      ])
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        prop: "formatted_duration",
                        label: _vm.$t("cloudbrainObj.runDuration"),
                        align: "center",
                        "header-align": "center",
                        width: "120"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("span", [
                                _vm._v(
                                  _vm._s(scope.row.task.formatted_duration)
                                )
                              ])
                            ]
                          }
                        }
                      ])
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        label: _vm.$t("cloudbrainObj.computeResource"),
                        align: "center",
                        "header-align": "center",
                        "min-width": "150"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("span", [
                                _vm._v(_vm._s(scope.row.task.computeSourceShow))
                              ])
                            ]
                          }
                        }
                      ])
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        label: _vm.$t("resourcesManagement.aiCenter"),
                        align: "center",
                        "header-align": "center",
                        width: "180"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("span", [
                                _vm._v(_vm._s(scope.row.task.ai_center))
                              ])
                            ]
                          }
                        }
                      ])
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        label: _vm.$t("resourcesManagement.accCardType"),
                        align: "center",
                        "header-align": "center",
                        width: "150"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("span", [
                                _vm._v(_vm._s(scope.row.task.accCardTypeShow))
                              ])
                            ]
                          }
                        }
                      ])
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        prop: "remark",
                        label: _vm.$t("operation"),
                        align: "left",
                        "min-width": "200",
                        "header-align": "center",
                        fixed: "right"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("div", { staticClass: "op-wrap" }, [
                                _c(
                                  "a",
                                  {
                                    class:
                                      scope.row.can_delete &&
                                      scope.row.task.canStop
                                        ? ""
                                        : "disabled",
                                    attrs: { href: "javascript:;" },
                                    on: {
                                      click: function($event) {
                                        return _vm.opStop(scope.row)
                                      }
                                    }
                                  },
                                  [_vm._v(_vm._s(_vm.$t("cloudbrainObj.stop")))]
                                ),
                                _vm._v(" "),
                                _c(
                                  "a",
                                  {
                                    staticClass: "delete",
                                    class:
                                      scope.row.can_delete &&
                                      scope.row.task.canDelete
                                        ? ""
                                        : "disabled",
                                    attrs: { href: "javascript:;" },
                                    on: {
                                      click: function($event) {
                                        return _vm.opDelete(scope.row)
                                      }
                                    }
                                  },
                                  [
                                    _vm._v(
                                      _vm._s(_vm.$t("cloudbrainObj.delete"))
                                    )
                                  ]
                                )
                              ])
                            ]
                          }
                        }
                      ])
                    })
                  ],
                  1
                )
              ],
              1
            )
          ]),
          _vm._v(" "),
          _c(
            "div",
            { staticClass: "foot" },
            [
              _c("el-pagination", {
                ref: "paginationRef",
                attrs: {
                  background: "",
                  "current-page": _vm.page,
                  "page-sizes": _vm.pageSizes,
                  "page-size": _vm.pageSize,
                  layout: "total, sizes, prev, pager, next, jumper",
                  total: _vm.total
                },
                on: {
                  "current-change": _vm.currentChange,
                  "size-change": _vm.sizeChange,
                  "update:currentPage": function($event) {
                    _vm.page = $event
                  },
                  "update:current-page": function($event) {
                    _vm.page = $event
                  },
                  "update:pageSize": function($event) {
                    _vm.pageSize = $event
                  },
                  "update:page-size": function($event) {
                    _vm.pageSize = $event
                  }
                }
              })
            ],
            1
          )
        ])
      ])
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ })

}]);