"use strict";
(self["webpackChunk"] = self["webpackChunk"] || []).push([["web_src_vuepages_components_cloudbrain_details_TaskDetailCollapseTabs_vue"],{

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=script&lang=js":
/*!********************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=script&lang=js ***!
  \********************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
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
/* harmony default export */ __webpack_exports__["default"] = ({
  name: "BaseDialog",
  props: {
    visible: {
      type: Boolean,
      "default": false
    },
    title: {
      type: String,
      "default": ""
    },
    width: {
      type: String,
      "default": ""
    },
    fullscreen: {
      type: Boolean,
      "default": false
    },
    top: {
      type: String
    },
    modal: {
      type: Boolean,
      "default": true
    },
    modalAppendToBody: {
      type: Boolean,
      "default": true
    },
    appendToBody: {
      type: Boolean,
      "default": false
    },
    lockScroll: {
      type: Boolean,
      "default": false
    },
    customClass: {
      type: String,
      "default": ""
    },
    closeOnClickModal: {
      type: Boolean,
      "default": false
    },
    closeOnPressEscape: {
      type: Boolean,
      "default": true
    },
    showClose: {
      type: Boolean,
      "default": true
    },
    beforeClose: {
      type: Function
    },
    center: {
      type: Boolean,
      "default": false
    },
    destroyOnClose: {
      type: Boolean,
      "default": false
    }
  },
  data: function data() {
    return {
      dialogShow: false
    };
  },
  watch: {
    visible: function visible(val) {
      this.dialogShow = val;
    }
  },
  methods: {
    open: function open() {
      this.$emit("open");
    },
    opened: function opened() {
      this.$emit("opened");
    },
    close: function close() {
      this.$emit("close");
    },
    closed: function closed() {
      this.$emit("closed");
      this.$emit("update:visible", false);
    }
  },
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/NotFound.vue?vue&type=script&lang=js":
/*!******************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/NotFound.vue?vue&type=script&lang=js ***!
  \******************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
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
/* harmony default export */ __webpack_exports__["default"] = ({
  name: "NotFound",
  props: {
    visible: {
      type: Boolean,
      "default": false
    }
  },
  data: function data() {
    return {};
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=script&lang=js":
/*!****************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=script&lang=js ***!
  \****************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
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
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'GeneralTaskCodeTips',
  props: {
    isTaskDetail: {
      "default": false
    }
  },
  data: function data() {
    return {};
  },
  methods: {},
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=script&lang=js":
/*!********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=script&lang=js ***!
  \********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
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
/* harmony default export */ __webpack_exports__["default"] = ({
  name: "LoadingMask",
  props: {
    loading: {
      type: Boolean,
      "default": false
    },
    content: {
      type: String,
      "default": ''
    }
  },
  data: function data() {
    return {};
  },
  computed: {
    tipsContent: function tipsContent() {
      return this.content || '加载中，请稍后';
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=script&lang=js":
/*!***************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=script&lang=js ***!
  \***************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.symbol.description */ "./node_modules/core-js/modules/es.symbol.description.js");
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.find */ "./node_modules/core-js/modules/es.array.find.js");
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.join */ "./node_modules/core-js/modules/es.array.join.js");
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_18___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_18__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_19___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_19__);
/* harmony import */ var _apis_modules_common__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! ~/apis/modules/common */ "./web_src/vuepages/apis/modules/common.js");
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var _GeneralTaskCodeTips_vue__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! ../GeneralTaskCodeTips.vue */ "./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_24__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");
/* harmony import */ var element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_25__ = __webpack_require__(/*! element-ui/lib/utils/date-util */ "./node_modules/element-ui/lib/utils/date-util.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_26__ = __webpack_require__(/*! spark-md5 */ "./node_modules/spark-md5/spark-md5.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_26___default = /*#__PURE__*/__webpack_require__.n(spark_md5__WEBPACK_IMPORTED_MODULE_26__);
/* harmony import */ var highlight_js__WEBPACK_IMPORTED_MODULE_27__ = __webpack_require__(/*! highlight.js */ "./node_modules/highlight.js/lib/index.js");
/* harmony import */ var highlight_js__WEBPACK_IMPORTED_MODULE_27___default = /*#__PURE__*/__webpack_require__.n(highlight_js__WEBPACK_IMPORTED_MODULE_27__);




















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








/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'ConfigInfo',
  props: {
    configs: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    nosdkcode: {
      type: Boolean,
      "default": false
    },
    isSubscriber: {
      type: Boolean,
      "default": false
    }
  },
  components: {
    GeneralTaskCodeTips: _GeneralTaskCodeTips_vue__WEBPACK_IMPORTED_MODULE_22__["default"]
  },
  data: function data() {
    return {
      clipboardHandler: null,
      sourceFtList: [],
      flag: false
    };
  },
  watch: {
    data: {
      immediate: true,
      deep: true,
      handler: function handler(newVal) {
        var _this = this;

        return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_19___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_17___default().mark(function _callee() {
          var res, data, repoName, repoOwnerName;
          return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_17___default().wrap(function _callee$(_context) {
            while (1) {
              switch (_context.prev = _context.next) {
                case 0:
                  if (!(_this.data.task && _this.data.task.source_id && !_this.flag)) {
                    _context.next = 10;
                    break;
                  }

                  _context.next = 3;
                  return (0,_apis_modules_common__WEBPACK_IMPORTED_MODULE_20__.getPromoteData)('model/modelfinetune.json');

                case 3:
                  res = _context.sent;
                  data = JSON.parse(res.data);
                  repoName = data.llm.repo_name;
                  repoOwnerName = data.llm.repo_owner_name;

                  _this.sourceFtList.push({
                    name: _this.data.task.app_name,
                    url: "/modelbase/nlp/sft/detail/".concat(repoOwnerName, "/").concat(repoName, "/").concat(_this.data.task.source_id)
                  });

                  console.log('this.sourceFtList', _this.sourceFtList);
                  _this.flag = true;

                case 10:
                  _this.$nextTick(function () {
                    _this.clipboardHandler && _this.clipboardHandler.destroy();
                    _this.clipboardHandler = (0,_utils__WEBPACK_IMPORTED_MODULE_23__.initClipboard)();
                  });

                case 11:
                case "end":
                  return _context.stop();
              }
            }
          }, _callee);
        }))();
      }
    }
  },
  computed: {// 计算字段分组配置
  },
  methods: {
    getCurrentFields: function getCurrentFields(groupName) {
      return this.configs.fields[groupName] || [];
    },
    sparkMD5Hash: function sparkMD5Hash() {
      var str = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : '';
      return spark_md5__WEBPACK_IMPORTED_MODULE_26___default().hash(str) + Math.random().toString().replace('0.', '');
    },
    renderHljs: function renderHljs() {
      var str = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : '';
      return highlight_js__WEBPACK_IMPORTED_MODULE_27___default().highlight('python', str).value;
    },
    renderTitle: function renderTitle(key) {
      if (!this.data.is_subscriber && (key === 'endPort' || key === 'port')) {
        return;
      }

      var task = this.data.task;
      var result = key;

      switch (key) {
        case 'taskName':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.taskName');
          break;

        case 'appName':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.appName');
          break;

        case 'spec':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.resourceSpec');
          break;

        case 'computerRes':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.computeResource');
          break;

        case 'status':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('status');
          break;

        case 'imagev1':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.image');
          break;

        case 'imagev2':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.image');
          break;

        case 'codeObsPath':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.codeObsPath');
          break;

        case 'aiCenter':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('resourcesManagement.aiCenter');
          break;

        case 'hasInternet':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.networkType');
          break;

        case 'creator':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('modelManage.creator');
          break;

        case 'repo':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('repos.repos');
          break;

        case 'branch':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.codeBranch');
          break;

        case 'runVersion':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.runVersion');
          break;

        case 'createTime':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.createTime');
          break;

        case 'startTime':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.startRunTime');
          break;

        case 'endTime':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.endRunTime');
          break;

        case 'duration':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.runDuration');
          break;

        case 'modelName':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('modelManage.modelName');
          break;

        case 'modelVersion':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('modelManage.modelVersion');
          break;

        case 'modelFiles':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('modelManage.modelFiles');
          break;

        case 'bootFile':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('modelManage.bootFile');
          break;

        case 'runParameters':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('modelManage.runParameters');
          break;

        case 'workServerNum':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('modelManage.workServerNumber');
          break;

        case 'descr':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('modelManage.descr');
          break;

        case 'codePath':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.codePath');
          break;

        case 'datasetPath':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.datasetPath');
          break;

        case 'modelPath':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.modelPath');
          break;

        case 'outputPath':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.outputPath');
          break;

        case 'bootFile':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('modelManage.bootFile');
          break;

        case 'failedReason':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.failedReason');
          break;

        case 'timeLimit':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.taskIsAutomaticStop');
          break;

        case 'endPort':
          result = task.end_point == '' ? '' : _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.customPath');
          break;

        case 'port':
          result = task.port == 0 ? '' : _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.port');
          break;

        case 'visualization':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.visualization');
          break;

        case 'datasetList':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('dataset');
          break;

        case 'modelList':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('repos.model');
          break;

        case 'sourceFtName':
          result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.sourceFtName');
          break;

        default:
          break;
      }

      return result;
    },
    renderContent: function renderContent(key) {
      if (!this.data.is_subscriber && (key === 'endPort' || key === 'port')) {
        return;
      }

      var task = this.data.task;
      var result = '--';

      switch (key) {
        case 'taskName':
          if (task.display_job_name) {
            result = task.display_job_name;
          }

          break;

        case 'appName':
          if (task.app_name) {
            result = task.app_name;
          }

          break;

        case 'imagev1':
          if (task.image_name || task.image_url) {
            result = "<span class=\"ui poping up clipboard\"\n            id=\"clipboard-".concat(spark_md5__WEBPACK_IMPORTED_MODULE_26___default().hash(task.image_url || task.image_name), "\"\n            data-position=\"top center\"\n            data-variation=\"inverted tiny\"\n            data-success=\"").concat(this.$t('copyedLink'), "\"\n            data-content=\"").concat(this.$t('copy'), "\" \n            data-original=\"").concat(this.$t('copy'), "\"\n            data-clipboard-text=\"").concat(task.image_url || task.image_name, "\"><span title=\"").concat(task.image_name || task.image_url, "\">").concat(task.image_name || task.image_url, "</span></span>");
          }

          break;

        case 'imagev2':
          if (task.image_name) {
            result = "<span class=\"ui poping up clipboard\"\n            id=\"clipboard-".concat(spark_md5__WEBPACK_IMPORTED_MODULE_26___default().hash(task.image_name), "\"\n            data-position=\"top center\"\n            data-variation=\"inverted tiny\"\n            data-success=\"").concat(this.$t('copySuccess'), "\"\n            data-content=\"").concat(this.$t('copy'), "\" \n            data-original=\"").concat(this.$t('copy'), "\"\n            data-clipboard-text=\"").concat(task.image_name, "\"><span title=\"").concat(task.image_name, "\">").concat(task.image_name, "</span></span>");
          }

          break;

        case 'codeObsPath':
          if (task.code_url) {
            result = "<span class=\"ui poping up clipboard\"\n            id=\"clipboard-".concat(spark_md5__WEBPACK_IMPORTED_MODULE_26___default().hash(task.code_url), "\"\n            data-position=\"top center\"\n            data-variation=\"inverted tiny\"\n            data-success=\"").concat(this.$t('copyedLink'), "\"\n            data-content=\"").concat(this.$t('copy'), "\" \n            data-original=\"").concat(this.$t('copy'), "\"\n            data-clipboard-text=\"").concat(task.code_url, "\"><span title=\"").concat(task.code_url, "\">").concat(task.code_url, "</span></span>");
          }

          break;

        case 'spec':
          if (task.spec) {
            var specObj = (0,_utils__WEBPACK_IMPORTED_MODULE_23__.renderSpecObject)(task.spec, false);
            result = specObj.specStr;
          }

          break;

        case 'computerRes':
          if (task.computeSourceShow || task.compute_source) {
            result = task.computeSourceShow || task.compute_source;
          }

          break;

        case 'status':
          if (task.status) {
            result = "<span style=\"display:flex;align-items: center;\"><i style=\"margin-right:4px;\" class=\"".concat(task.status, "\"></i><span>").concat(task.status, "</span>");

            if (task.status === 'WAITING') {
              if (task.detailed_status === 'dataMigrating') {
                result += "<i style=\"margin-left:8px;\" class=\"dataMigrating\" title=\"".concat(_langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.migratingData'), "\"></i>");
              }

              if (task.detailed_status === 'centerPending') {
                result += "<i style=\"margin-left:8px;\" class=\"centerPending\" title=\"".concat(_langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.centerPending'), "\"></i>");
              }

              if (task.detailed_status === 'ImagePulling') {
                result += "<i style=\"margin-left:8px;\" class=\"ImagePulling\" title=\"".concat(_langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.ImagePulling'), "\"></i>");
              }
            }

            result += "</span>";
          }

          break;

        case 'aiCenter':
          if (task.ai_center) {
            result = task.ai_center;
          }

          break;

        case 'hasInternet':
          if (task.has_internet == 1) {
            result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.noInternet');
          } else if (task.has_internet == 2) {
            result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.hasInternet');
          }

          break;

        case 'repo':
          if (task.repo_name && task.repo_owner_name) {
            result = "<a style=\"color:#0066ff;\" target=\"_blank\" href=\"/".concat(task.repo_owner_name, "/").concat(task.repo_name, "\">").concat(task.repo_owner_name, " / ").concat(task.repo_alias, "</a>");
          } else if (task.repo_id) {
            result = "<span style=\"color:#888888\">-- \uFF08".concat(_langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('repositoryWasDel'), "\uFF09</span>");
          }

          break;

        case 'branch':
          if (task.branch_name) {
            result = "".concat(task.branch_name) + (task.commit_id ? "<span class=\"commit-id\">".concat(task.commit_id.slice(0, 10), "</span>") : '');
          }

          break;

        case 'creator':
          if (task.creator_name) {
            result = task.creator_name;
          }

          break;

        case 'runVersion':
          if (task.current_version_name) {
            task.current_version_name;
          }

          break;

        case 'createTime':
          if (task.createTimeStr) {
            result = task.createTimeStr;
          }

          break;

        case 'startTime':
          if (task.start_time) {
            result = (0,element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_25__.formatDate)(new Date(task.start_time * 1000), 'yyyy-MM-dd HH:mm:ss');
          }

          break;

        case 'endTime':
          if (task.end_time) {
            result = (0,element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_25__.formatDate)(new Date(task.end_time * 1000), 'yyyy-MM-dd HH:mm:ss');
          }

          break;

        case 'duration':
          if (task.formatted_duration) {
            result = task.formatted_duration;
          }

          break;

        case 'modelName':
          if (task.pretrain_model_name) {
            result = task.pretrain_model_name;
          }

          break;

        case 'modelVersion':
          if (task.pretrain_model_version) {
            result = task.pretrain_model_version;
          }

          break;

        case 'modelFiles':
          if (task.pretrain_model_ckpt_name) {
            result = "<span title=\"".concat(task.pretrain_model_ckpt_name, "\">").concat(task.pretrain_model_ckpt_name, "</span>");
          }

          break;

        case 'bootFile':
          if (task.boot_file) {
            result = task.boot_file;
          }

          break;

        case 'runParameters':
          var parameters = task.parameters || {
            parameter: []
          };
          var paramList = parameters.parameter || [];
          var paramsStr = paramList.map(function (item) {
            return "".concat(item.label, " = ").concat(item.value);
          }).join('; ');

          if (paramsStr) {
            result = "<span class=\"ui poping up clipboard\"\n            id=\"clipboard-".concat(spark_md5__WEBPACK_IMPORTED_MODULE_26___default().hash(paramsStr), "\"\n            data-position=\"top center\"\n            data-variation=\"inverted tiny\"\n            data-success=\"").concat(this.$t('copySuccess'), "\"\n            data-content=\"").concat(this.$t('copy'), "\" \n            data-original=\"").concat(this.$t('copy'), "\"\n            data-clipboard-text=\"").concat(paramsStr, "\"><span title=\"").concat(paramsStr, "\">").concat(paramsStr, "</span></span>");
          }

          break;

        case 'workServerNum':
          if (task.work_server_number) {
            result = task.work_server_number;
          }

          break;

        case 'descr':
          if (task.description) {
            result = "<span title=\"".concat((0,_utils__WEBPACK_IMPORTED_MODULE_23__.escapeHTML)(task.description), "\">").concat((0,_utils__WEBPACK_IMPORTED_MODULE_23__.escapeHTML)(task.description), "</span>");
          }

          break;

        case 'codePath':
          result = task.code_path;
          break;

        case 'datasetPath':
          result = task.dataset_path;
          break;

        case 'modelPath':
          result = task.pretrain_model_path;
          break;

        case 'outputPath':
          result = task.output_path;
          break;

        case 'bootFile':
          if (task.boot_file) {
            result = task.boot_file;
          }

          break;

        case 'failedReason':
          if (task.failed_reason) {
            result = task.failed_reason;
          }

          break;

        case 'timeLimit':
          if (task.time_limit == 0) {
            result = '';
          } else if (task.time_limit == -1) {
            result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.manualStop');
          } else {
            result = _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.automaticStop') + "(".concat(_langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.numOfHours', {
              num: task.time_limit
            }), ")");
          }

          break;

        case 'endPort':
          if (task.end_point) {
            result = task.end_point;
          }

          break;

        case 'port':
          if (task.port) {
            result = task.port;
          }

          break;

        case 'visualization':
          result = !!task.visualize_required ? _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.tensorBoardVisualization') : _langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('none');
          break;

        case 'datasetList':
          if (task.dataset_list && task.dataset_list.length) {
            result = '<div class="ui list">';
            task.dataset_list.forEach(function (item) {
              if (item.is_delete) {
                result += "<div style=\"line-height:20px\" class=\"item nowrap\" title=".concat(item.dataset_alias, ">").concat(item.dataset_alias, " (").concat(_langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.fileWasDeleted'), ")</div>");
              } else {
                result += "<a style=\"line-height:20px;color:#0066ff;\" target=\"_blank\" href=\"/datasets/detail/".concat(item.owner_name, "/").concat(item.dataset_name, "\" class=\"item nowrap\" title=").concat(item.owner_name, "/").concat(item.dataset_alias, ">").concat(item.owner_name, "/").concat(item.dataset_alias, "</a>");
              }
            });
            result += '</div>';
          }

          if (task.job_type === 'EVAL') {
            var datasetParam = task.parameters.parameter.find(function (item) {
              return item.label === 'datasets';
            });
            var selectDatasetTemp = [];

            if (datasetParam && datasetParam.value) {
              selectDatasetTemp = datasetParam.value.split(' ').sort(function (a, b) {
                return a.localeCompare(b);
              }).map(function (item) {
                return {
                  k: item,
                  v: item
                };
              });
              if (selectDatasetTemp.length > 0) result = '<div class="ui list">';
              selectDatasetTemp.forEach(function (item) {
                result += "<a style=\"line-height:20px;color:#0066ff;\" target=\"_blank\" href=\"/datasets/detail/Open_Dataset/".concat(item.k, "\" class=\"item nowrap\" title=").concat(item.k, ">").concat(item.k, "</a>");
              });
            }
          }

          break;

        case 'modelList':
          if (task.pretrain_model_list && task.pretrain_model_list.length) {
            result = '<div class="ui list">';
            task.pretrain_model_list.forEach(function (item) {
              if (item.is_delete) {
                result += "<div style=\"line-height:20px\" class=\"item nowrap\" title=".concat(item.alias, ">").concat(item.alias, " (").concat(_langs__WEBPACK_IMPORTED_MODULE_24__.i18n.t('cloudbrainObj.fileWasDeleted'), ")</div>");
              } else {
                result += "<a style=\"line-height:20px;color:#0066ff;\" target=\"_blank\" href=\"/models/detail/".concat(item.owner_name, "/").concat(item.name, "\" class=\"item nowrap\" title=").concat(item.owner_name, "/").concat(item.alias, ">").concat(item.owner_name, "/").concat(item.alias, "</a>");
              }
            });
            result += '</div>';
          }

          break;

        case 'sourceFtName':
          if (task.source_id && task.app_name) {
            result = '<div class="ui list">';
            result += "<a style=\"line-height:20px;color:#0066ff;\" target=\"_blank\" href=\"/modelbase/nlp/sft/detail/".concat(task.source_id, "\" class=\"item nowrap\" title=").concat(task.app_name, ">").concat(task.app_name, "</a>");
            result += '</div>';
          }

          break;

        default:
          break;
      }

      return result;
    },
    refresh: function refresh() {}
  },
  beforeMount: function beforeMount() {},
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=script&lang=js":
/*!***************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=script&lang=js ***!
  \***************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.find */ "./node_modules/core-js/modules/es.array.find.js");
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.includes */ "./node_modules/core-js/modules/es.array.includes.js");
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.object.define-properties */ "./node_modules/core-js/modules/es.object.define-properties.js");
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.object.define-property */ "./node_modules/core-js/modules/es.object.define-property.js");
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptor */ "./node_modules/core-js/modules/es.object.get-own-property-descriptor.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptors */ "./node_modules/core-js/modules/es.object.get-own-property-descriptors.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_parse_int__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.parse-int */ "./node_modules/core-js/modules/es.parse-int.js");
/* harmony import */ var core_js_modules_es_parse_int__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_parse_int__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! @babel/runtime/helpers/toConsumableArray */ "./node_modules/@babel/runtime/helpers/toConsumableArray.js");
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_18___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_18__);
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var _pages_model_llms_componenes_render__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! ~/pages/model/llms/componenes/render */ "./web_src/vuepages/pages/model/llms/componenes/render.js");




















function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_18___default()(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }

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


var md = null;
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'EvalDetail',
  props: {
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  data: function data() {
    return {
      tableData: [],
      resultType: -1,
      resultTypeList: [{
        k: this.$t('modelFinetune.allResult'),
        v: -1
      }, {
        k: this.$t('modelFinetune.correctResult'),
        v: 1
      }, {
        k: this.$t('modelFinetune.wrongResult'),
        v: 0
      }],
      selectDataset: [],
      dataset: 'all',
      currentPage: 1,
      pageSize: 10,
      totalNum: 0,
      evalRunning: true
    };
  },
  watch: {
    data: {
      deep: true,
      // 深度观察，检测嵌套属性的变化
      handler: function handler(newVal) {
        console.log("xxxxxxxxxxx");

        if (newVal.task && ["STOPPED", "FAILED", "START_FAILED", "COMPLETED", "SUCCEEDED", "CREATED_FAILED"].includes(newVal.task.status)) {
          this.evalRunning = false;
        }
      }
    },
    tableData: {
      handler: function handler() {
        this.checkContentHeights();
      },
      deep: true
    }
  },
  methods: {
    changeDataset: function changeDataset() {
      this.currentPage = 1;
      this.refresh();
    },
    handleCurrentChange: function handleCurrentChange(val) {
      this.currentPage = val;
      this.refresh();
    },
    toggleExpand: function toggleExpand(row, field) {
      this.$set(row, field, !row[field]);
    },
    checkContentHeights: function checkContentHeights() {
      var _this = this;

      this.$nextTick(function () {
        _this.tableData.forEach(function (row) {
          // 检查rawContent
          if (row.rawContent) {
            var rawEl = document.querySelector(".raw-content[data-id=\"".concat(row.id, "\"]"));

            if (rawEl) {
              var lineHeight = parseInt(window.getComputedStyle(rawEl).lineHeight);

              _this.$set(row, 'rawNeedsExpand', rawEl.scrollHeight > lineHeight * 2);
            }
          } // 检查predictContent


          if (row.predictContent) {
            var predictEl = document.querySelector(".predict-content[data-id=\"".concat(row.id, "\"]"));

            if (predictEl) {
              var _lineHeight = parseInt(window.getComputedStyle(predictEl).lineHeight);

              _this.$set(row, 'predictNeedsExpand', predictEl.scrollHeight > _lineHeight * 3);
            }
          } // 检查predContent


          if (row.predContent) {
            var predEl = document.querySelector(".pred-content[data-id=\"".concat(row.id, "\"]"));

            if (predEl) {
              var _lineHeight2 = parseInt(window.getComputedStyle(predEl).lineHeight);

              _this.$set(row, 'predNeedsExpand', predEl.scrollHeight > _lineHeight2 * 2);
            }
          }
        });
      });
    },
    refresh: function refresh() {
      var _this2 = this;

      var task = this.data.task;
      if (!task || this.evalRunning) return;
      var queryParams = {
        task_id: task.id,
        filter: this.dataset,
        page: this.currentPage,
        pagesize: this.pageSize,
        result: this.resultType
      };
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_19__.getAiEvalDetailResult)(queryParams).then(function (res) {
        res = res.data;
        console.log(res);

        if (res.code == 0 && res.data) {
          _this2.totalNum = res.count;
          _this2.tableData = res.data.map(function (item) {
            console.log(JSON.parse(item.raw_input));
            var inputBlock = "```json\n".concat(item.raw_input, "\n```");
            return _objectSpread(_objectSpread({}, item), {}, {
              rawContent: md.render(inputBlock),
              predictContent: md.render(item.predict),
              predContent: md.render(item.pred),
              rawExpanded: false,
              predictExpanded: false,
              predExpanded: false,
              rawNeedsExpand: false,
              predictNeedsExpand: false,
              predNeedsExpand: false
            });
          });

          _this2.checkContentHeights();
        } else {}
      })["catch"](function (err) {
        _this2.$message.error(err);

        console.log(err);
      });
    }
  },
  created: function created() {
    md = (0,_pages_model_llms_componenes_render__WEBPACK_IMPORTED_MODULE_20__["default"])();
  },
  mounted: function mounted() {
    var selectDatasetTemp = [];
    var datasetParam = this.data.task.parameters.parameter.find(function (item) {
      return item.label === 'datasets';
    });

    if (datasetParam && datasetParam.value) {
      selectDatasetTemp = datasetParam.value.split(' ').sort(function (a, b) {
        return a.localeCompare(b);
      }).map(function (item) {
        return {
          k: item,
          v: item
        };
      });
    }

    this.selectDataset = [{
      k: this.$t('modelFinetune.allDataset'),
      v: 'all'
    }].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_17___default()(selectDatasetTemp)); // if ([ "STOPPED","FAILED","START_FAILED","COMPLETED","SUCCEEDED","CREATED_FAILED"].includes(this.data?.task?.status)) {
    //   this.evalRunning = false
    // }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=script&lang=js":
/*!*****************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=script&lang=js ***!
  \*****************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.includes */ "./node_modules/core-js/modules/es.array.includes.js");
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.is-array */ "./node_modules/core-js/modules/es.array.is-array.js");
/* harmony import */ var core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.join */ "./node_modules/core-js/modules/es.array.join.js");
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.object.define-properties */ "./node_modules/core-js/modules/es.object.define-properties.js");
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.object.define-property */ "./node_modules/core-js/modules/es.object.define-property.js");
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptor */ "./node_modules/core-js/modules/es.object.get-own-property-descriptor.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptors */ "./node_modules/core-js/modules/es.object.get-own-property-descriptors.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_string_includes__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.string.includes */ "./node_modules/core-js/modules/es.string.includes.js");
/* harmony import */ var core_js_modules_es_string_includes__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_includes__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var echarts__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! echarts */ "./node_modules/echarts/index.js");


















function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_16___default()(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }

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


var chartOptions = {
  xAxis: {
    type: 'category',
    axisLabel: {
      interval: 0,
      // 强制显示所有标签，不省略
      show: true,
      fontSize: 10,
      formatter: function formatter(value) {
        // 对长标签进行换行处理
        if (value.length > 7) {
          if (value.includes('-')) {
            return value.split('-').join('-\n');
          } // 没有连接符，从中间位置换行


          var mid = Math.ceil(value.length / 2);
          return value.substring(0, mid) + '\n' + value.substring(mid);
          return value.split('-').join('-\n');
        }

        return value;
      } // rotate: 45, // 旋转45度
      // margin: 10 // 增加边距

    },
    data: []
  },
  yAxis: {
    type: 'value',
    max: 1
  },
  series: [{
    data: [],
    type: 'bar',
    barMaxWidth: 10,
    itemStyle: {
      color: 'rgb(80, 135, 236)' // 设置柱子的颜色

    },
    label: {
      show: true,
      position: 'top'
    }
  }]
};
var chartHandler;
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'Loss',
  props: {
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  data: function data() {
    return {
      tableData: [],
      modelName: '',
      evalRunning: true,
      lang: 'zh-CN',
      loading: false
    };
  },
  watch: {
    data: {
      deep: true,
      // 深度观察，检测嵌套属性的变化
      handler: function handler(newVal) {
        if (newVal.task && ["STOPPED", "FAILED", "START_FAILED", "COMPLETED", "SUCCEEDED", "CREATED_FAILED"].includes(newVal.task.status)) {
          this.evalRunning = false;
        }
      }
    }
  },
  methods: {
    refresh: function refresh() {
      var _this = this;

      var task = this.data.task;
      if (!task) return;
      this.loading = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_17__.getAiEvalResult)({
        id: task.id
      }).then(function (res) {
        _this.loading = false;
        res = res.data;
        console.log(res);
        var tmpArray = [];
        var keyArray = [];
        var valueArray = [];

        if (Array.isArray(res.result) && res.result.length > 0) {
          res.result.forEach(function (item) {
            tmpArray.push(_objectSpread({}, item));
            keyArray.push(item.name);
            valueArray.push(item.value);
          });
          _this.tableData = tmpArray;
          chartOptions.series[0].data = valueArray;
          chartOptions.xAxis.data = keyArray;
          console.log(chartOptions);
          chartHandler && chartHandler.dispose();
          chartHandler = echarts__WEBPACK_IMPORTED_MODULE_18__.init(_this.$refs.chartRef);
          chartHandler.setOption(chartOptions);
        } // if (res.code == 0) {
        //   if (res.status == "RUNNING") {
        //   } else {
        //   }
        // }

      })["catch"](function (err) {
        _this.loading = false;

        _this.$message.error(err);

        console.log(err);
      });
    },
    resize: function resize() {
      chartHandler && chartHandler.resize();
    }
  },
  mounted: function mounted() {
    var _this$data, _this$data$task;

    this.lang = document.querySelector('html').getAttribute('lang');
    this.modelName = (_this$data = this.data) === null || _this$data === void 0 ? void 0 : (_this$data$task = _this$data.task) === null || _this$data$task === void 0 ? void 0 : _this$data$task.app_name;
    window.addEventListener('resize', this.resize);
  },
  beforeDestroy: function beforeDestroy() {
    window.removeEventListener('resize', this.resize);
    chartHandler && chartHandler.dispose();
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=script&lang=js":
/*!******************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=script&lang=js ***!
  \******************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.find-index */ "./node_modules/core-js/modules/es.array.find-index.js");
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.reduce */ "./node_modules/core-js/modules/es.array.reduce.js");
/* harmony import */ var core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.array.splice */ "./node_modules/core-js/modules/es.array.splice.js");
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.number.to-fixed */ "./node_modules/core-js/modules/es.number.to-fixed.js");
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_number_to_precision__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.number.to-precision */ "./node_modules/core-js/modules/es.number.to-precision.js");
/* harmony import */ var core_js_modules_es_number_to_precision__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_to_precision__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.object.assign */ "./node_modules/core-js/modules/es.object.assign.js");
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.object.define-properties */ "./node_modules/core-js/modules/es.object.define-properties.js");
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.object.define-property */ "./node_modules/core-js/modules/es.object.define-property.js");
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptor */ "./node_modules/core-js/modules/es.object.get-own-property-descriptor.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptors */ "./node_modules/core-js/modules/es.object.get-own-property-descriptors.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_18___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_18__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_19___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_19__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_20___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_20__);
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! core-js/modules/es.string.trim */ "./node_modules/core-js/modules/es.string.trim.js");
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_21___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_21__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_22___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_22__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_23___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_23__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_24__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_24___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_24__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_25__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_25___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_25__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_26__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_26___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_26__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_27__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_27___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_27__);
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_28__ = __webpack_require__(/*! @babel/runtime/helpers/toConsumableArray */ "./node_modules/@babel/runtime/helpers/toConsumableArray.js");
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_28___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_28__);
/* harmony import */ var _components_BaseDialog_vue__WEBPACK_IMPORTED_MODULE_29__ = __webpack_require__(/*! ~/components/BaseDialog.vue */ "./web_src/vuepages/components/BaseDialog.vue");
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var _apis_modules_storage__WEBPACK_IMPORTED_MODULE_31__ = __webpack_require__(/*! ~/apis/modules/storage */ "./web_src/vuepages/apis/modules/storage.js");
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_32__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");






























function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_27___default()(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }

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




var UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB'];
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'ExportDataset',
  props: {
    disabled: {
      type: Boolean,
      "default": true
    },
    configs: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  components: {
    BaseDialog: _components_BaseDialog_vue__WEBPACK_IMPORTED_MODULE_29__["default"]
  },
  data: function data() {
    return {
      dlgShow: false,
      loading: false,
      state: {
        dataset_name: '',
        tab: 1,
        filesStr: ''
      },
      dataset_id: '',
      progressId: '',
      firstOpen: false,
      treeData: [],
      selectedData: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      uploading: false,
      progressTimer: null,
      remaining_storage: 0,
      totalSize: 0,
      isStorageExceeded: false,
      dlgActiveName: 'first',
      dlgSearchValue: '',
      dlgLoading: false,
      dlgModelTreeData: [],
      dlgTotal: 0,
      dlgPage: 1,
      dlgPageSize: 6,
      popoverVisible: false,
      loginName: '',
      ownerName: '',
      showOwenerTips: false
    };
  },
  computed: {
    // 格式化的剩余空间显示（自动单位转换）
    formattedRemaining: function formattedRemaining() {
      return this.formatBytes(Math.max(this.remaining_storage, 0));
    },
    formattedTotal: function formattedTotal() {
      return this.formatBytes(this.totalSize);
    }
  },
  methods: {
    changeDataset: function changeDataset(id) {
      var _this = this;

      this.dlgModelTreeData.forEach(function (item) {
        if (item.id == id) {
          _this.state.dataset_name = item.alias;
          console.log(item.alias);
          _this.ownerName = item.owner_name;

          if (_this.loginName != _this.ownerName) {
            _this.showOwenerTips = true;
          }
        }
      });
      this.getStorageSummary(id);
      console.log(this.state.dataset_name);
    },
    dlgTabClick: function dlgTabClick(tab, event) {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.searchModelData();
    },
    inputSearch: function inputSearch() {
      this.dlgTotal = 0;
      this.dlgPage = 1;
      this.searchModelData();
    },
    dlgPageChange: function dlgPageChange(page) {
      this.dlgPage = page;
      this.searchModelData();
    },
    searchModelData: function searchModelData() {
      var _this2 = this;

      var tabName = this.dlgActiveName;
      var tabPrams = {
        'first': '/owned',
        'second': '/collaborated'
      };
      var params = {
        url: tabPrams[tabName],
        q: this.dlgSearchValue.trim(),
        page: this.dlgPage,
        page_size: this.dlgPageSize
      };
      this.dlgLoading = true;
      (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_32__.getDatasets)(params).then(function (res) {
        console.log(res);
        _this2.dlgLoading = false;

        if (res.data.code === 0) {
          var _res$data, _res$data$data, _res$data2, _res$data2$data;

          var data = ((_res$data = res.data) === null || _res$data === void 0 ? void 0 : (_res$data$data = _res$data.data) === null || _res$data$data === void 0 ? void 0 : _res$data$data.datasets) || [];
          _this2.dlgModelTreeData = data;
          _this2.dlgTotal = ((_res$data2 = res.data) === null || _res$data2 === void 0 ? void 0 : (_res$data2$data = _res$data2.data) === null || _res$data2$data === void 0 ? void 0 : _res$data2$data.total) || 0;
        } else {
          _this2.$message.error(rs.data.msg || '获取数据集失败');

          console.log(err);
        }
      })["catch"](function (err) {
        _this2.dlgLoading = false;
        console.log(err);
      });
    },
    onFileCheckChange: function onFileCheckChange() {
      var selectedData = this.$refs.fileTreeRef.getCheckedNodes();
      var fliterfile = selectedData.filter(function (item) {
        return !item.isDir;
      });
      this.state.filesStr = fliterfile.reduce(function (pre, cur) {
        return cur.isDir ? pre : (pre ? pre + ',' : pre) + cur.FileName;
      }, '');
      this.selectedData = _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_28___default()(fliterfile.map(function (item) {
        return _objectSpread(_objectSpread({}, item), {}, {
          statusCode: -99
        });
      }));
      this.totalSize = fliterfile.reduce(function (pre, cur) {
        return cur.isDir ? pre : pre + cur.Size;
      }, 0);
      this.isStorageExceeded = this.totalSize > this.remaining_storage;
    },
    removeFile: function removeFile(data) {
      this.$refs.fileTreeRef.setChecked(data.id, false, false);
      var index = this.selectedData.findIndex(function (item) {
        return item.id === data.id;
      });
      this.selectedData.splice(index, 1);
      this.state.filesStr = this.selectedData.reduce(function (pre, cur) {
        return cur.IsDir ? pre : (pre ? pre + ',' : pre) + cur.FileName;
      }, '');
    },
    getFiles: function getFiles() {
      var _this3 = this;

      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.getAiTaskOutputResultAll)({
        id: this.data.id
      }).then(function (res) {
        res = res.data;

        if (res.code == 0) {
          var _res$data3, _res$data3$output;

          var data = ((_res$data3 = res.data) === null || _res$data3 === void 0 ? void 0 : (_res$data3$output = _res$data3.output) === null || _res$data3$output === void 0 ? void 0 : _res$data3$output.file_list) || [];
          var nodeMap = {};

          for (var i = 0, iLen = data.length; i < iLen; i++) {
            var dataI = data[i];
            var path = dataI.FileName.split('/');
            var curNode = nodeMap;

            for (var j = 0, jLen = path.length; j < jLen; j++) {
              var cur = path[j];

              if (!curNode[cur]) {
                curNode[cur] = {};
              }

              if (j == jLen - 1) {
                dataI._isLeaf = true;
                curNode[cur] = dataI;
              }

              curNode = curNode[cur];
            }
          }

          var nodeData = [];

          var walkNode = function walkNode(curNode, nodeList) {
            if (curNode._isLeaf) return;

            for (var key in curNode) {
              var node = {
                label: key,
                isDir: !curNode[key]._isLeaf,
                children: [],
                id: curNode[key].FileName
              };
              nodeList.push(node);

              if (curNode[key]._isLeaf) {
                delete node.children;
                Object.assign(node, curNode[key]);
              }

              walkNode(curNode[key], node.children);
            }
          };

          walkNode(nodeMap, nodeData);
          _this3.treeData = nodeData;
        } else {
          _this3.treeData = [];
        }
      })["catch"](function (err) {
        console.log(err);
      });
    },
    startGetProgressTimer: function startGetProgressTimer() {
      var _this4 = this;

      this.progressTimer && clearInterval(this.progressTimer);
      this.progressTimer = setInterval(function () {
        _this4.getProgress();
      }, 5 * 1000);
    },
    getProgress: function getProgress(isFirst) {
      var _this5 = this;

      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.getAiTaskExportDatasetProgress)({
        id: this.progressId // ${this.state.tab}_

      }).then(function (res) {
        console.log(res);
        res = res.data;

        if (res.code == 0) {
          var result = res.data;

          if (isFirst) {
            if (result && Object.keys(result).length > 0) {
              if (result['##type##'] !== undefined) {
                _this5.state.tab = Number(result['##type##']);
              }

              var files = Object.keys(result).filter(function (item) {
                return item !== '##type##';
              });
              files.forEach(function (item) {
                var statusCode = result[item];

                _this5.selectedData.push({
                  label: item,
                  id: item,
                  FileName: item,
                  children: [],
                  statusCode: statusCode
                });

                if (statusCode == 0) {
                  _this5.uploading = true;

                  _this5.startGetProgressTimer();
                }
              });
            }
          } else {
            if (result && Object.keys(result).length > 0) {
              var _files = Object.keys(result).filter(function (item) {
                return item !== '##type##';
              });

              var endStatusCount = 0;

              for (var i = 0, iLen = _files.length; i < iLen; i++) {
                var item = _files[i];

                for (var j = 0, jLen = _this5.selectedData.length; j < jLen; j++) {
                  var selectedFile = _this5.selectedData[j];

                  if (selectedFile.id == item) {
                    var statusCode = result[item];

                    if (statusCode == -1 || statusCode == -2 || statusCode == 100) {
                      endStatusCount++;
                    }

                    selectedFile.statusCode = statusCode;
                    break;
                  }
                }
              }

              if (endStatusCount == _this5.selectedData.length) {
                _this5.uploading = false;
                _this5.progressTimer && clearInterval(_this5.progressTimer);
              }
            }
          }
        }
      })["catch"](function (err) {
        _this5.progressTimer && clearInterval(_this5.progressTimer);
        console.log(err);
      });
    },
    submit: function submit() {
      var _this6 = this;

      if (!this.dataset_id) {
        this.$message({
          type: 'info',
          message: this.$t('cloudbrainObj.exportDataset.please_select_dataset')
        });
        return;
      }

      if (!this.selectedData.length || !this.state.filesStr.length) {
        this.$message({
          type: 'info',
          message: this.$t('cloudbrainObj.exportDataset.please_select_file')
        });
        return;
      }

      var subData = {
        task_id: this.data.id,
        file_list: this.state.filesStr
      };
      this.uploading = true;
      this.selectedData.forEach(function (item) {
        return item.statusCode = 0;
      });
      console.log("this.selectedData", this.selectedData);
      console.log("subData", subData);
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.setAiTaskExportDataset)({
        dataset_id: this.dataset_id,
        id: this.data.id
      }, subData).then(function (res) {
        console.log("res", res);
        res = res.data;

        if (res.code == 0) {
          _this6.progressId = res.data.id;

          _this6.startGetProgressTimer();
        } else {
          _this6.$message({
            type: 'error',
            message: res.msg
          });
        }
      })["catch"](function (err) {
        _this6.uploading = false;
        console.log(err);
      });
    },
    cancel: function cancel() {
      this.dlgShow = false;
    },
    open: function open() {
      if (this.firstOpen) return;
      this.uploading = false;
      this.getFiles();
      this.firstOpen = true;
    },
    closed: function closed() {
      this.progressTimer && clearInterval(this.progressTimer);
    },
    getStorageSummary: function getStorageSummary(id) {
      var _this7 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_26___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_24___default().mark(function _callee() {
        var res;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_24___default().wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                _context.next = 2;
                return (0,_apis_modules_storage__WEBPACK_IMPORTED_MODULE_31__.getStorageSummary)({
                  subject_id: id,
                  subject_type: '1'
                });

              case 2:
                res = _context.sent;
                _this7.remaining_storage = res.data.remaining_storage;

                _this7.onFileCheckChange();

              case 5:
              case "end":
                return _context.stop();
            }
          }
        }, _callee);
      }))();
    },
    // 智能单位格式化
    formatBytes: function formatBytes(bytes) {
      var unitIndex = 0;
      var value = bytes;

      while (value >= 1024 && unitIndex < UNITS.length - 1) {
        value /= 1024;
        unitIndex++;
      }

      return [this.toPrecision(value, 2), UNITS[unitIndex]];
    },
    // 精确小数处理
    toPrecision: function toPrecision(value) {
      var decimals = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 2;
      if (value < 0.001 && value > 0) return '<0.001';
      return Number(value.toFixed(decimals)).toString();
    }
  },
  beforeMount: function beforeMount() {},
  mounted: function mounted() {
    console.log('this.data', this.data);
    this.loginName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext');
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=script&lang=js":
/*!****************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=script&lang=js ***!
  \****************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.symbol.description */ "./node_modules/core-js/modules/es.symbol.description.js");
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.find-index */ "./node_modules/core-js/modules/es.array.find-index.js");
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.join */ "./node_modules/core-js/modules/es.array.join.js");
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.array.reduce */ "./node_modules/core-js/modules/es.array.reduce.js");
/* harmony import */ var core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.array.splice */ "./node_modules/core-js/modules/es.array.splice.js");
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.number.to-fixed */ "./node_modules/core-js/modules/es.number.to-fixed.js");
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_number_to_precision__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.number.to-precision */ "./node_modules/core-js/modules/es.number.to-precision.js");
/* harmony import */ var core_js_modules_es_number_to_precision__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_to_precision__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/es.object.assign */ "./node_modules/core-js/modules/es.object.assign.js");
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! core-js/modules/es.object.define-properties */ "./node_modules/core-js/modules/es.object.define-properties.js");
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! core-js/modules/es.object.define-property */ "./node_modules/core-js/modules/es.object.define-property.js");
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptor */ "./node_modules/core-js/modules/es.object.get-own-property-descriptor.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_18___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_18__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptors */ "./node_modules/core-js/modules/es.object.get-own-property-descriptors.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_19___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_19__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_20___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_20__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_21___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_21__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_22___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_22__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_23___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_23__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_24__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_24___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_24__);
/* harmony import */ var core_js_modules_es_string_starts_with__WEBPACK_IMPORTED_MODULE_25__ = __webpack_require__(/*! core-js/modules/es.string.starts-with */ "./node_modules/core-js/modules/es.string.starts-with.js");
/* harmony import */ var core_js_modules_es_string_starts_with__WEBPACK_IMPORTED_MODULE_25___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_starts_with__WEBPACK_IMPORTED_MODULE_25__);
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_26__ = __webpack_require__(/*! core-js/modules/es.string.trim */ "./node_modules/core-js/modules/es.string.trim.js");
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_26___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_26__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_27__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_27___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_27__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_28__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_28___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_28__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_29__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_29___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_29__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_30__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_30___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_30__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_31__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_31___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_31__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_32__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_32___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_32__);
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_33__ = __webpack_require__(/*! @babel/runtime/helpers/toConsumableArray */ "./node_modules/@babel/runtime/helpers/toConsumableArray.js");
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_33___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_33__);
/* harmony import */ var _components_BaseDialog_vue__WEBPACK_IMPORTED_MODULE_34__ = __webpack_require__(/*! ~/components/BaseDialog.vue */ "./web_src/vuepages/components/BaseDialog.vue");
/* harmony import */ var _const__WEBPACK_IMPORTED_MODULE_35__ = __webpack_require__(/*! ~/const */ "./web_src/vuepages/const/index.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_36__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_37__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var _apis_modules_modelmanage__WEBPACK_IMPORTED_MODULE_38__ = __webpack_require__(/*! ~/apis/modules/modelmanage */ "./web_src/vuepages/apis/modules/modelmanage.js");
/* harmony import */ var _apis_modules_storage__WEBPACK_IMPORTED_MODULE_39__ = __webpack_require__(/*! ~/apis/modules/storage */ "./web_src/vuepages/apis/modules/storage.js");
/* harmony import */ var _apis_modules_organization__WEBPACK_IMPORTED_MODULE_40__ = __webpack_require__(/*! ~/apis/modules/organization */ "./web_src/vuepages/apis/modules/organization.js");
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_41__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");



































function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_32___default()(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }

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








var UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB'];
var MAX_LABEL_COUNT = 5;
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'ExportModel',
  props: {
    disabled: {
      type: Boolean,
      "default": true
    },
    configs: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  components: {
    BaseDialog: _components_BaseDialog_vue__WEBPACK_IMPORTED_MODULE_34__["default"]
  },
  data: function data() {
    return {
      dlgShow: false,
      loading: false,
      state: {
        taskName: '',
        name: '',
        alias: '',
        owner_id: '',
        engine: 0,
        filesStr: '',
        label: '',
        license: '',
        isPrivate: '0'
      },
      licenseList: [],
      nameErr: false,
      aliasErr: false,
      modelFileErr: false,
      engineList: _const__WEBPACK_IMPORTED_MODULE_35__.MODEL_ENGINES,
      treeData: [],
      selectedData: [],
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      remaining_storage: 0,
      totalSize: 0,
      isStorageExceeded: false,
      owners: [],
      progressId: '',
      progressTimer: null,
      orgName: ''
    };
  },
  watch: {
    'state.owner_id': {
      handler: function handler(newVal, oldVal) {
        // 在这里处理owner_id变化后的逻辑
        if (newVal !== oldVal) {
          console.log('owner_id从', oldVal, '变更为', newVal, this.owners); // 调用相关处理函数

          var findItem = this.owners.filter(function (item) {
            return item.ID === newVal;
          });
          console.log("findItem", findItem);

          if (findItem[0].IsOrganization) {
            this.orgName = findItem[0].Name;
          } else {
            this.orgName = '';
          }

          this.getStorageSummary(); // this.handleOwnerIdChange(newVal, oldVal)
        }
      },
      deep: false // 由于是基本类型监听，不需要深度监听

    }
  },
  computed: {
    // 格式化的剩余空间显示（自动单位转换）
    formattedRemaining: function formattedRemaining() {
      return this.formatBytes(Math.max(this.remaining_storage, 0));
    },
    formattedTotal: function formattedTotal() {
      return this.formatBytes(this.totalSize);
    }
  },
  methods: {
    handleKeyDown: function handleKeyDown(e) {
      if (e.key === ' ' || e.keyCode === 32) {
        e.preventDefault(); // 阻止空格输入
      }
    },
    checkName: function checkName() {
      console.log("checkName", /^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,98}[a-zA-Z0-9]$/.test(this.state.name));

      if (/^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,98}[a-zA-Z0-9]$/.test(this.state.name)) {
        this.nameErr = false;
      } else {
        this.nameErr = true;
      }

      return !this.nameErr;
    },
    checkAlias: function checkAlias() {
      console.log("checkAlias", /^[\u4e00-\u9fa5a-zA-Z0-9-_.]{0,100}$/.test(this.state.alias));

      if (/^[\u4e00-\u9fa5a-zA-Z0-9-_.]{0,100}$/.test(this.state.alias)) {
        this.aliasErr = false;
      } else {
        this.aliasErr = true;
      }

      return !this.aliasErr;
    },
    checkModelFile: function checkModelFile() {
      this.modelFileErr = !this.state.filesStr;
      return !this.modelFileErr;
    },
    labelInput: function labelInput() {
      var hasEndSpace = this.state.label[this.state.label.length - 1] == ' ';
      var list = this.state.label.trim().split(' ').filter(function (label) {
        return label != '';
      });
      this.state.label = list.slice(0, MAX_LABEL_COUNT).join(' ') + (hasEndSpace && list.length < MAX_LABEL_COUNT ? ' ' : '');
    },
    handleClearLicenseClick: function handleClearLicenseClick() {
      this.state.license = '';
    },
    startGetProgressTimer: function startGetProgressTimer() {
      var _this = this;

      this.progressTimer && clearInterval(this.progressTimer);
      this.progressTimer = setInterval(function () {
        _this.getProgress();
      }, 5 * 1000);
    },
    getProgress: function getProgress(isFirst) {
      var _this2 = this;

      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_37__.getAiTaskExportModelProgress)({
        id: this.progressId // ${this.state.tab}_

      }).then(function (res) {
        console.log(res);
        res = res.data;

        if (res.code == 0) {
          var result = res.data;

          if (isFirst) {
            console.log("xxxxxxxxxx");

            if (result && Object.keys(result).length > 0) {
              if (result['##type##'] !== undefined) {
                _this2.state.tab = Number(result['##type##']);
              }

              var files = Object.keys(result).filter(function (item) {
                return item !== '##type##';
              });
              files.forEach(function (item) {
                var statusCode = result[item];

                _this2.selectedData.push({
                  label: item,
                  FileName: item,
                  children: [],
                  statusCode: statusCode
                });

                if (statusCode == 0) {
                  _this2.uploading = true;

                  _this2.startGetProgressTimer();
                }
              });
            }
          } else {
            if (result && Object.keys(result).length > 0) {
              var _files = Object.keys(result).filter(function (item) {
                return item !== '##type##';
              });

              var endStatusCount = 0;
              console.log("files.length", _files.length);
              console.log("selectedData", _this2.selectedData.length);

              for (var i = 0, iLen = _files.length; i < iLen; i++) {
                var item = _files[i];

                for (var j = 0, jLen = _this2.selectedData.length; j < jLen; j++) {
                  var selectedFile = _this2.selectedData[j];

                  if (selectedFile.id == item) {
                    var statusCode = result[item];

                    if (statusCode == -1 || statusCode == -2 || statusCode == -3 || statusCode == 100) {
                      endStatusCount++;
                    }

                    selectedFile.statusCode = statusCode;
                    break;
                  }
                }
              }

              if (endStatusCount == _this2.selectedData.length) {
                _this2.loading = false;
                _this2.progressTimer && clearInterval(_this2.progressTimer);
              }
            }
          }
        }
      })["catch"](function (err) {
        _this2.loading = false;
        _this2.progressTimer && clearInterval(_this2.progressTimer);
        console.log(err);
      });
    },
    submit: function submit() {
      var _this3 = this;

      console.log("sbmit", this.checkName(), this.checkAlias());
      this.state.name = this.state.name.trim();

      if (!this.checkName()) {
        this.$message({
          type: 'error',
          message: this.$t('modelManage.pleaseInputModelName')
        });
        return;
      }

      if (!this.checkAlias()) {
        return;
      }

      if (!this.checkModelFile()) {
        this.$message({
          type: 'error',
          message: this.$t('modelObj.model_export_placeholder')
        });
        return;
      }

      var subData = {
        name: this.state.name,
        alias: this.state.alias,
        owner_id: this.state.owner_id,
        aimodel_type: 0,
        is_private: this.state.isPrivate == 1 ? true : false,
        engine: this.state.engine,
        label: this.state.label.split(/\s+/).join(' ').trim(),
        license: this.state.license,
        task_id: this.data.id,
        file_list: this.state.filesStr
      };
      this.loading = true;
      this.selectedData.forEach(function (item) {
        return item.statusCode = 0;
      });
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_37__.setAiTaskExportModel)({
        id: this.data.id
      }, subData).then(function (res) {
        res = res.data;
        console.log("res", res);

        if (res.code == 0) {
          _this3.progressId = res.data.process_id;

          _this3.startGetProgressTimer();
        } else {
          _this3.selectedData.forEach(function (item) {
            return item.statusCode = -99;
          });

          _this3.loading = false;

          _this3.$message({
            type: 'error',
            message: res.msg
          });
        }
      })["catch"](function (err) {
        _this3.loading = false;

        _this3.selectedData.forEach(function (item) {
          return item.statusCode = -99;
        });

        console.log(err);
      });
    },
    cancel: function cancel() {
      this.dlgShow = false;
    },
    isMindSporeEngine: function isMindSporeEngine(obj) {
      if (obj.engine_name != null && obj.engine_name.toLowerCase().startsWith("mindspore")) {
        return true;
      }

      if (obj.engine_id == 122 || obj.engine_id == 35 || obj.engine_id == -1 || obj.engine_id == 37) {
        return true;
      }

      return false;
    },
    onFileCheckChange: function onFileCheckChange(data) {
      var selectedData = this.$refs.fileTreeRef.getCheckedNodes();
      var fliterfile = selectedData.filter(function (item) {
        return !item.isDir;
      });
      this.state.filesStr = fliterfile.reduce(function (pre, cur) {
        return cur.isDir ? pre : (pre ? pre + ',' : pre) + cur.FileName;
      }, '');
      this.selectedData = _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_33___default()(fliterfile.map(function (item) {
        return _objectSpread(_objectSpread({}, item), {}, {
          statusCode: -99
        });
      }));
      this.totalSize = fliterfile.reduce(function (pre, cur) {
        return cur.isDir ? pre : pre + cur.Size;
      }, 0);
      this.isStorageExceeded = this.totalSize > this.remaining_storage;
      this.checkModelFile();
    },
    removeFile: function removeFile(data) {
      console.log("data", data);
      this.$refs.fileTreeRef.setChecked(data.id, false, false);
      var index = this.selectedData.findIndex(function (item) {
        return item.id === data.id;
      });
      this.selectedData.splice(index, 1);
      this.totalSize = this.selectedData.reduce(function (pre, cur) {
        return cur.isDir ? pre : pre + cur.Size;
      }, 0);
      this.state.filesStr = this.selectedData.reduce(function (pre, cur) {
        return cur.IsDir ? pre : (pre ? pre + ',' : pre) + cur.FileName;
      }, '');
      this.isStorageExceeded = this.totalSize > this.remaining_storage;
    },
    open: function open() {
      var _this4 = this;

      this.state.filesStr = '';
      this.state.taskName = this.data.display_job_name;
      this.state.name = this.data.display_job_name + '_model_' + Math.random().toString(36).substr(2, 4);

      if (this.isMindSporeEngine(this.data)) {
        this.state.engine = 2;
      } else {
        if (this.data.engine_id == 121 || this.data.engine_id == 38) {
          this.state.engine = 1;
        } else {
          this.state.engine = 0;
        }
      }

      this.selectedData = [];
      this.state.label = '';
      this.state.description = '';
      this.loading = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_37__.getAiTaskOutputResultAll)({
        id: this.data.id
      }).then(function (res) {
        res = res.data;
        _this4.loading = false;

        if (res.code == 0) {
          var _res$data, _res$data$output;

          var data = ((_res$data = res.data) === null || _res$data === void 0 ? void 0 : (_res$data$output = _res$data.output) === null || _res$data$output === void 0 ? void 0 : _res$data$output.file_list) || [];
          var nodeMap = {};

          for (var i = 0, iLen = data.length; i < iLen; i++) {
            var dataI = data[i];
            var path = dataI.FileName.split('/');
            var curNode = nodeMap;

            for (var j = 0, jLen = path.length; j < jLen; j++) {
              var cur = path[j];

              if (!curNode[cur]) {
                curNode[cur] = {};
              }

              if (j == jLen - 1) {
                dataI._isLeaf = true;
                curNode[cur] = dataI;
              }

              curNode = curNode[cur];
            }
          }

          var nodeData = [];

          var walkNode = function walkNode(curNode, nodeList) {
            if (curNode._isLeaf) return;

            for (var key in curNode) {
              var node = {
                label: key,
                isDir: !curNode[key]._isLeaf,
                children: [],
                id: curNode[key].FileName
              };
              nodeList.push(node);

              if (curNode[key]._isLeaf) {
                delete node.children;
                Object.assign(node, curNode[key]);
              }

              walkNode(curNode[key], node.children);
            }
          };

          walkNode(nodeMap, nodeData);
          _this4.treeData = nodeData;
        }
      })["catch"](function (err) {
        console.log(err);
        _this4.loading = false;
      });
      (0,_apis_modules_modelmanage__WEBPACK_IMPORTED_MODULE_38__.getModelLicenseList)().then(function (res) {
        res = res.data;

        try {
          var license = JSON.parse(res) || [];
          _this4.licenseList = license;

          if (license.length) {
            _this4.state.license = license[0].id;
          }
        } catch (err) {
          console.log(err);
        }
      })["catch"](function (err) {
        console.log(err);
      });
    },
    closed: function closed() {
      this.progressTimer && clearInterval(this.progressTimer);
    },
    getStorageSummary: function getStorageSummary() {
      var _this5 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_31___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_29___default().mark(function _callee() {
        var queryParams, response, res;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_29___default().wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                queryParams = !_this5.orgName ? (0,_apis_modules_storage__WEBPACK_IMPORTED_MODULE_39__.getStorageSummary)({}) : (0,_apis_modules_organization__WEBPACK_IMPORTED_MODULE_40__.getOrgStorageSummary)(_this5.orgName, {});
                _context.next = 3;
                return queryParams;

              case 3:
                response = _context.sent;
                res = !_this5.orgName ? response.data : response.data.data;
                console.log("res", res);
                _this5.remaining_storage = res.remaining_storage || 0;
                _this5.totalSize = _this5.selectedData.reduce(function (pre, cur) {
                  return cur.isDir ? pre : pre + cur.Size;
                }, 0);
                _this5.isStorageExceeded = _this5.totalSize > _this5.remaining_storage;

              case 9:
              case "end":
                return _context.stop();
            }
          }
        }, _callee);
      }))();
    },
    // 智能单位格式化
    formatBytes: function formatBytes(bytes) {
      var unitIndex = 0;
      var value = bytes;

      while (value >= 1024 && unitIndex < UNITS.length - 1) {
        value /= 1024;
        unitIndex++;
      }

      return [this.toPrecision(value, 2), UNITS[unitIndex]];
    },
    // 精确小数处理
    toPrecision: function toPrecision(value) {
      var decimals = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 2;
      if (value < 0.001 && value > 0) return '<0.001';
      return Number(value.toFixed(decimals)).toString();
    },
    getUsersList: function getUsersList() {
      var _this6 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_31___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_29___default().mark(function _callee2() {
        var response, res;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_29___default().wrap(function _callee2$(_context2) {
          while (1) {
            switch (_context2.prev = _context2.next) {
              case 0:
                _context2.prev = 0;
                _context2.next = 3;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_41__.getAvailableUsers)({}, 'aimodel');

              case 3:
                response = _context2.sent;
                res = response.data;
                _this6.loading = false;

                if (!(res.code === 0)) {
                  _context2.next = 13;
                  break;
                }

                _this6.owners = res.data.users || [];

                if (!(_this6.owners.length === 0)) {
                  _context2.next = 10;
                  break;
                }

                return _context2.abrupt("return");

              case 10:
                _this6.state.owner_id = _this6.owners[0].ID;
                _context2.next = 14;
                break;

              case 13:
                _this6.$message.error(response.data.msg);

              case 14:
                _context2.next = 20;
                break;

              case 16:
                _context2.prev = 16;
                _context2.t0 = _context2["catch"](0);
                _this6.loading = false;

                _this6.$message.error(_context2.t0);

              case 20:
              case "end":
                return _context2.stop();
            }
          }
        }, _callee2, null, [[0, 16]]);
      }))();
    }
  },
  beforeMount: function beforeMount() {
    this.getUsersList();
  },
  mounted: function mounted() {// this.getStorageSummary()
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=script&lang=js":
/*!**********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
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
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'Icons',
  props: {
    type: {
      type: String,
      "default": '1'
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=script&lang=js":
/*!*********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=script&lang=js ***!
  \*********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");





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

/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'Logs',
  props: {
    configs: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  data: function data() {
    return {
      common: {
        nodeSel: 0,
        content: '',
        lines: 60,
        order: 'up',
        startLine: '',
        endLine: '',
        loading: false,
        scrollTop: 0,
        downloadUrl: ''
      },
      fullscreen: {
        nodeSel: 0,
        content: '',
        lines: 100,
        order: 'up',
        startLine: '',
        endLine: '',
        loading: false,
        scrollTop: 0,
        downloadUrl: ''
      },
      multiNodesData: [],
      dialogShow: false,
      canLogDownload: 0,
      downloading: false
    };
  },
  methods: {
    escapeHTML: function escapeHTML(str) {
      return str.toString().replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;").replace(/'/g, "&apos;");
    },
    getLogs: function getLogs(callback) {
      var _this$multiNodesData$,
          _this$multiNodesData$2,
          _this = this;

      var key = this.dialogShow ? 'fullscreen' : 'common';
      this[key].loading = true;
      var task = this.data.task;
      var params = {
        id: task.id,
        base_line: this[key].order == 'up' ? this[key].startLine : this[key].endLine,
        lines: this[key].lines,
        order: this[key].order,
        node_id: this.configs.multiNodes ? (_this$multiNodesData$ = this.multiNodesData[this[key].nodeSel]) === null || _this$multiNodesData$ === void 0 ? void 0 : _this$multiNodesData$.id : undefined,
        log_file_name: this.configs.multiNodes ? (_this$multiNodesData$2 = this.multiNodesData[this[key].nodeSel]) === null || _this$multiNodesData$2 === void 0 ? void 0 : _this$multiNodesData$2.log_file_name : undefined
      };
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_5__.getAiTaskLogs)(params).then(function (res) {
        _this[key].loading = false;
        res = res.data;

        if (res.code == 0) {
          res = res.data || {};

          if (_this.canLogDownload === 0) {
            _this.canLogDownload = res.can_log_download;
          }

          if (res.lines > 0) {
            if (_this[key].order == 'up') {
              _this[key].content = "<pre>".concat(_this.escapeHTML(res.content), "</pre>") + _this[key].content;
              _this[key].startLine = res.start_line;

              if (_this[key].endLine) {
                _this.stayScrollPos();
              }

              if (!_this[key].endLine) {
                _this[key].endLine = res.end_line;
              }
            } else if (_this[key].order == 'down') {
              _this[key].content += "<pre>".concat(_this.escapeHTML(res.content), "</pre>");
              _this[key].endLine = res.end_line;

              if (!_this[key].startLine) {
                _this[key].startLine = res.start_line;
              }
            }
          } else {
            if (_this.configs.noScroll) {
              _this[key].content = "<pre>".concat(_this.escapeHTML(res.content), "</pre>");
            } else {
              var msg = '';

              if (!_this[key].content) {
                // log content is empty
                msg = _this[key].order == 'down' ? _this.$t('cloudbrainObj.scrolledToTopTip') : _this.$t('cloudbrainObj.scrolledToBottomTip');
              } else {
                msg = _this[key].order == 'up' ? _this.$t('cloudbrainObj.scrolledToTopTip') : _this.$t('cloudbrainObj.scrolledToBottomTip');
              }

              _this.$message({
                type: 'info',
                message: msg
              });
            }
          }

          callback && callback();
        } else {
          _this.$message({
            type: 'error',
            message: res.msg
          });
        }
      })["catch"](function (err) {
        _this[key].loading = false;
        console.log(err);
      });
    },
    scrollHandler: function scrollHandler(evt) {
      if (this.configs.noScroll) return;
      var key = this.dialogShow ? 'fullscreen' : 'common';
      if (this[key].loading) return;
      var logEle = this.$refs[key + 'LogContentRef'];

      if (logEle) {
        var scrollHeight = logEle.scrollHeight;
        var clientHeight = logEle.clientHeight;
        var scrollTop = logEle.scrollTop;

        if (scrollTop != this[key].scrollTop) {
          if (scrollTop == 0) {
            this[key].order = 'up';
            this.getLogs();
          } else if (scrollTop + clientHeight >= scrollHeight - 1) {
            this[key].order = 'down';
            this.getLogs();
          }
        }

        this[key].scrollTop = scrollTop;
      }
    },
    goTop: function goTop() {
      var _this2 = this;

      var key = this.dialogShow ? 'fullscreen' : 'common';
      this[key].content = '';
      this[key].startLine = '';
      this[key].endLine = '';
      this[key].order = 'down';
      this.getLogs(function () {
        _this2.stayScrollPos();

        _this2.$nextTick(function () {
          var logEle = _this2.$refs[key + 'LogContentRef'];

          if (logEle) {
            logEle.scrollTo({
              top: 0,
              behavior: 'smooth'
            });
          }
        });
      });
    },
    goBottom: function goBottom() {
      this.refresh();
    },
    stayScrollPos: function stayScrollPos() {
      var key = this.dialogShow ? 'fullscreen' : 'common';
      var logEle = this.$refs[key + 'LogContentRef'];

      if (logEle) {
        var scrollHeight = logEle.scrollHeight;
        var scrollTop = logEle.scrollTop;
        this.$nextTick(function () {
          var scrollHeightNew = logEle.scrollHeight;
          var scrollTopNew = logEle.scrollTop;
          logEle.scrollTo({
            top: scrollTop + (scrollHeightNew - scrollHeight),
            behavior: 'instant'
          });
        });
      }
    },
    scrollBottomAnimation: function scrollBottomAnimation() {
      var key = this.dialogShow ? 'fullscreen' : 'common';
      var logEle = this.$refs[key + 'LogContentRef'];

      if (logEle) {
        var scrollHeight = logEle.scrollHeight;
        var clientHeight = logEle.clientHeight;
        var scrollTop = logEle.scrollTop;
        logEle.scrollTo({
          top: scrollHeight - clientHeight,
          behavior: 'smooth'
        });
      }
    },
    getFirstLogs: function getFirstLogs() {
      var _this3 = this,
          _this$multiNodesData$3,
          _this$multiNodesData$4;

      this.getLogs(function () {
        _this3.$nextTick(function () {
          _this3.scrollBottomAnimation();
        });
      });
      var key = this.dialogShow ? 'fullscreen' : 'common';
      var task = this.data.task;
      this[key].downloadUrl = (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_5__.getAiTaskLogsDownloadUrl)({
        id: task.id,
        node_id: this.configs.multiNodes ? (_this$multiNodesData$3 = this.multiNodesData[this[key].nodeSel]) === null || _this$multiNodesData$3 === void 0 ? void 0 : _this$multiNodesData$3.id : undefined,
        log_file_name: this.configs.multiNodes ? (_this$multiNodesData$4 = this.multiNodesData[this[key].nodeSel]) === null || _this$multiNodesData$4 === void 0 ? void 0 : _this$multiNodesData$4.log_file_name : undefined
      });
      console.log("this[key].downloadUrl", this[key].downloadUrl);
    },
    refresh: function refresh() {
      var _this4 = this;

      var key = this.dialogShow ? 'fullscreen' : 'common';
      this[key].nodeSel = 0;
      this[key].content = '';
      this[key].startLine = '';
      this[key].endLine = '';
      this[key].order = 'up';
      this[key].downloadUrl = '';

      if (this.configs.multiNodes) {
        var task = this.data.task;
        (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_5__.getAiTaskNodeInfo)({
          repoOwnerName: task.repoOwnerName,
          repoName: task.repoName,
          id: task.id
        }).then(function (res) {
          res = res.data;

          if (res && res.code == 0) {
            var _res$data;

            _this4.multiNodesData = ((_res$data = res.data) === null || _res$data === void 0 ? void 0 : _res$data.nodes) || [];
          }

          _this4.getFirstLogs();
        })["catch"](function (err) {
          console.log(err);
        });
      } else {
        this.getFirstLogs();
      }
    },
    changeNode: function changeNode() {
      var key = this.dialogShow ? 'fullscreen' : 'common';
      this[key].content = '';
      this[key].startLine = '';
      this[key].endLine = '';
      this[key].order = 'up';
      this[key].downloadUrl = '';
      this.getFirstLogs();
    },
    open: function open() {
      var key = 'fullscreen';
      this[key].nodeSel = 0;
      this[key].content = '';
      this[key].startLine = '';
      this[key].endLine = '';
      this[key].order = 'up';
      this[key].downloadUrl = '';
    },
    opened: function opened() {
      this.refresh();
    },
    closed: function closed() {
      var key = 'fullscreen';
      this[key].nodeSel = 0;
      this[key].content = '';
      this[key].startLine = '';
      this[key].endLine = '';
      this[key].order = 'up';
      this[key].downloadUrl = '';
    }
  },
  beforeMount: function beforeMount() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=script&lang=js":
/*!*********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=script&lang=js ***!
  \*********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var echarts__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! echarts */ "./node_modules/echarts/index.js");


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


var chartOptions = {
  xAxis: {
    type: 'value',
    splitLine: false,
    name: 'epoch',
    nameLocation: 'center',
    nameGap: 30,
    nameTextStyle: {
      fontSize: 16,
      fontWeight: 'bold'
    },
    axisLabel: {
      fontSize: 14,
      fontWeight: 'bold'
    }
  },
  yAxis: {
    type: 'value',
    splitLine: false,
    name: 'loss',
    nameLocation: 'center',
    nameGap: 30,
    nameTextStyle: {
      fontSize: 16,
      fontWeight: 'bold'
    },
    axisLabel: {
      fontSize: 14,
      fontWeight: 'bold'
    },
    scale: true
  },
  series: [{
    data: [],
    type: 'line',
    showSymbol: false,
    smooth: 0.2
  }]
};
var chartHandler;
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'Loss',
  props: {
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  data: function data() {
    return {
      lossValue: [],
      loading: false
    };
  },
  methods: {
    refresh: function refresh() {
      var _this = this;

      var task = this.data.task;
      if (!task) return;
      this.loading = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_2__.getAiTaskLoss)({
        repoOwnerName: task.repoOwnerName,
        repoName: task.repoName,
        id: task.id
      }).then(function (res) {
        _this.loading = false;
        res = res.data;
        var tmpArray = [];

        if (res.code == 0 && res.data) {
          res.data.forEach(function (item) {
            tmpArray.push([item.epoch, item.loss]);
          });
          _this.lossValue = tmpArray;
          chartOptions.series[0].data = tmpArray;
          chartHandler && chartHandler.dispose();
          chartHandler = echarts__WEBPACK_IMPORTED_MODULE_3__.init(_this.$refs.chartRef);
          chartHandler.setOption(chartOptions); // const events = res.data.events || [];
          // events.forEach(item => {
          //   item.timestampStr = item.timestamp ? formatDate(new Date(item.timestamp), 'yyyy/MM/dd HH:mm:ss') : '';
          // });
          // this.events = events;
        } else {
          _this.$message.error(res.msg);
        }
      })["catch"](function (err) {
        _this.loading = false;

        _this.$message.error(err);

        console.log(err);
      });
    },
    resize: function resize() {
      chartHandler && chartHandler.resize();
    }
  },
  mounted: function mounted() {
    window.addEventListener('resize', this.resize);
  },
  beforeDestroy: function beforeDestroy() {
    window.removeEventListener('resize', this.resize);
    chartHandler && chartHandler.dispose();
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=script&lang=js":
/*!*********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=script&lang=js ***!
  \*********************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! element-ui/lib/utils/date-util */ "./node_modules/element-ui/lib/utils/date-util.js");



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


/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'OperationProfile',
  props: {
    configs: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  data: function data() {
    return {
      events: [],
      loading: false
    };
  },
  methods: {
    refresh: function refresh() {
      var _this = this;

      // console.log('OperationProfile refresh');
      var task = this.data.task;
      if (!task) return;
      this.loading = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_3__.getAiTaskOperationProfile)({
        repoOwnerName: task.repoOwnerName,
        repoName: task.repoName,
        id: task.id
      }).then(function (res) {
        _this.loading = false;
        res = res.data;

        if (res.code == 0 && res.data) {
          var events = res.data.events || [];
          events.forEach(function (item) {
            item.timestampStr = item.timestamp ? (0,element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_4__.formatDate)(new Date(item.timestamp), 'yyyy/MM/dd HH:mm:ss') : '';
          });
          _this.events = events;
        } else {
          _this.events = [];
        }
      })["catch"](function (err) {
        _this.loading = false;
        console.log(err);
      });
    }
  },
  beforeMount: function beforeMount() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=script&lang=js":
/*!*******************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=script&lang=js ***!
  \*******************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_from__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.from */ "./node_modules/core-js/modules/es.array.from.js");
/* harmony import */ var core_js_modules_es_array_from__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_from__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.index-of */ "./node_modules/core-js/modules/es.array.index-of.js");
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.join */ "./node_modules/core-js/modules/es.array.join.js");
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.number.to-fixed */ "./node_modules/core-js/modules/es.number.to-fixed.js");
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.string.iterator */ "./node_modules/core-js/modules/es.string.iterator.js");
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var echarts__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! echarts */ "./node_modules/echarts/index.js");














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



var sortBy = function sortBy(arr, k) {
  return arr.concat().sort(function (a, b) {
    return a[k] > b[k] ? 1 : a[k] < b[k] ? -1 : 0;
  });
};

var chartOptions = {
  legend: {
    width: '95%',
    data: []
  },
  grid: {
    top: "60",
    bottom: "5%",
    left: '32',
    right: '80',
    x: "2%",
    containLabel: true
  },
  tooltip: {
    trigger: "axis",
    backgroundColor: "rgb(51, 56, 84)",
    borderColor: "rgb(51, 51, 51)",
    borderWidth: 0,
    textStyle: {
      color: "#fff"
    },
    axisPointer: {
      type: "cross"
    },
    appendToBody: true
  },
  xAxis: {
    type: "category",
    data: [],
    boundaryGap: false,
    axisLabel: {
      interval: "auto"
    },
    name: ""
  },
  yAxis: [{
    show: true,
    name: "(%)",
    position: 'left',
    axisLine: {
      show: true
    },
    axisTick: {
      show: true
    }
  }, {
    show: false,
    name: "Value",
    position: 'right',
    axisLine: {
      show: true
    },
    axisTick: {
      show: true
    }
  }],
  series: []
};
var chartHandler;
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'ResourceUseage',
  props: {
    configs: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  data: function data() {
    return {
      chartData: {},
      loading: false,
      nodeSel: 0,
      multiNodesData: [],
      showErrorMsg: false,
      errorMsg: ''
    };
  },
  methods: {
    checkIsValueType: function checkIsValueType(name) {
      return name.indexOf('Bytes') >= 0 || name.indexOf('Rate') >= 0 || name.indexOf('numProcesses') >= 0;
    },
    getChartData: function getChartData(useRefreshBtn) {
      var _this$multiNodesData$,
          _this$multiNodesData$2,
          _this = this;

      var task = this.data.task;
      this.loading = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_14__.getAiTaskResourceUseage)({
        repoOwnerName: task.repoOwnerName,
        repoName: task.repoName,
        id: task.id,
        node_id: this.configs.multiNodes ? (_this$multiNodesData$ = this.multiNodesData[this.nodeSel]) === null || _this$multiNodesData$ === void 0 ? void 0 : _this$multiNodesData$.id : undefined,
        log_file_name: this.configs.multiNodes ? (_this$multiNodesData$2 = this.multiNodesData[this.nodeSel]) === null || _this$multiNodesData$2 === void 0 ? void 0 : _this$multiNodesData$2.log_file_name : undefined
      }).then(function (res) {
        _this.loading = false;
        res = res.data;

        if (res.code == 0) {
          var data = res.data || {};
          var metricsInfo = data.metrics_info || [];
          var filterData = metricsInfo.filter(function (item) {
            // return !["recvBytesRate", "diskWriteRate", "sendBytesRate", "diskReadRate",].includes(item.name);
            return item.value && item.value.length;
          });
          filterData = sortBy(filterData, "name");
          var legenData = filterData.map(function (item) {
            return item.name;
          });
          var valueTypeCount = 0;
          var seriesData = filterData.map(function (item) {
            var value = (item.value || []).map(function (item) {
              return item > 0 ? Number(Number(item).toFixed(3)) : "0";
            });

            var valueType = _this.checkIsValueType(item.name);

            valueTypeCount += valueType ? 1 : 0;
            var seriesOption = {
              name: item.name,
              type: "line",
              symbol: "circle",
              symbolSize: 10,
              smooth: true,
              showSymbol: false,
              yAxisIndex: valueType ? 1 : 0,
              lineStyle: {
                width: 2,
                shadowColor: "rgba(0,0,0,0.3)",
                shadowBlur: 10,
                shadowOffsetY: 8
              },
              data: value
            };
            return seriesOption;
          });
          var xLength = metricsInfo.length ? metricsInfo[0].value.length : 0;
          var xInterval = data.interval || 1;
          chartOptions.xAxis.data = Array.from({
            length: xLength
          }, function (_, index) {
            return index * xInterval;
          });
          chartOptions.legend.data = legenData;
          var legendSelected = {};
          legenData.forEach(function (element) {
            var valueType = _this.checkIsValueType(element);

            legendSelected[element] = !valueType;
          });

          if (valueTypeCount > 0) {
            chartOptions.yAxis[1].show = true;
          } else if (chartOptions.yAxis[1]) {
            chartOptions.yAxis.pop();
          }

          if (useRefreshBtn && chartHandler && Object.keys(legendSelected).join('') == Object.keys(chartOptions.legend.selected || {}).join('')) {
            try {
              var selected = chartHandler.getOption().legend[0].selected;
              chartOptions.legend.selected = selected;
            } catch (_unused) {}
          } else {
            chartOptions.legend.selected = legendSelected;
          }

          chartOptions.series = seriesData;
          chartOptions.grid.top = '60';

          if (legenData.length == 0) {
            chartOptions.grid.top = '50';
          }

          if (legenData.length >= 8) {
            chartOptions.grid.top = '80';
          }

          if (legenData.length >= 16) {
            chartOptions.grid.top = '100';
          }

          chartHandler && chartHandler.dispose();
          chartHandler = echarts__WEBPACK_IMPORTED_MODULE_15__.init(_this.$refs.chartRef);
          chartHandler.setOption(chartOptions);
        } else {
          _this.showErrorMsg = true;
          _this.errorMsg = res.msg || _this.$t('common.unknownError');
          return;
        }
      })["catch"](function (err) {
        _this.loading = false;
        console.log(err);
      });
    },
    refresh: function refresh(useRefreshBtn) {
      var _this2 = this;

      var task = this.data.task;

      if (this.configs.multiNodes) {
        this.loading = true;
        (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_14__.getAiTaskNodeInfo)({
          repoOwnerName: task.repoOwnerName,
          repoName: task.repoName,
          id: task.id
        }).then(function (res) {
          res = res.data;

          if (res && res.code == 0) {
            var _res$data;

            _this2.multiNodesData = ((_res$data = res.data) === null || _res$data === void 0 ? void 0 : _res$data.nodes) || [];
          }

          _this2.getChartData(useRefreshBtn);
        })["catch"](function (err) {
          _this2.loading = false;
          console.log(err);
        });
      } else {
        this.getChartData(useRefreshBtn);
      }
    },
    changeNode: function changeNode() {
      this.getChartData();
    },
    resize: function resize() {
      chartHandler && chartHandler.resize();
    }
  },
  beforeMount: function beforeMount() {
    chartOptions.xAxis.name = this.$t('cloudbrainObj.chartTime');
    chartOptions.yAxis[0].name = this.$t('cloudbrainObj.chartResourceUsage');
  },
  mounted: function mounted() {
    window.addEventListener('resize', this.resize);
  },
  beforeDestroy: function beforeDestroy() {
    window.removeEventListener('resize', this.resize);
    chartHandler && chartHandler.dispose();
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=script&lang=js":
/*!*******************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=script&lang=js ***!
  \*******************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.find-index */ "./node_modules/core-js/modules/es.array.find-index.js");
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.join */ "./node_modules/core-js/modules/es.array.join.js");
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");








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


/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'ResultDownload',
  props: {
    configs: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  data: function data() {
    return {
      filesList: [],
      filePath: [],
      loading: false,
      isTaskTerminal: null,
      canDownload: false,
      canReschedule: false,
      resultStatus: 0,
      resultPath: '',
      downloadAllUrl: '',
      size: 0
    };
  },
  methods: {
    getDirFiles: function getDirFiles(dir) {
      var _this = this;

      dir = dir.length ? dir.slice(1) : '';
      var task = this.data.task;
      this.loading = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_8__.getAiTaskOutputResult)({
        repoOwnerName: task.repoOwnerName,
        repoName: task.repoName,
        id: task.id,
        parent_dir: dir
      }).then(function (res) {
        _this.loading = false;
        res = res.data;

        if (res.code == 0 && res.data && res.data.output) {
          var data = res.data.output;
          _this.isTaskTerminal = data.is_task_terminal;
          _this.canDownload = data.can_download;
          _this.canReschedule = data.can_reschedule;
          _this.resultStatus = data.status;
          _this.resultPath = data.path;
          _this.size = data.output_size_limit;
          _this.downloadAllUrl = (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_8__.getDownLoadAiTaskResultFileAllUrl)({
            repoOwnerName: task.repoOwnerName,
            repoName: task.repoName,
            id: task.id
          });

          if (_this.resultStatus == 0) {
            var _this$$refs$tableRef;

            // 成功 0
            var list = data.file_list || [];
            list.forEach(function (item) {
              item.SizeShow = item.IsDir ? '' : (0,_utils__WEBPACK_IMPORTED_MODULE_9__.transFileSize)(item.Size);
              item.ModTimeNum = new Date(item.ModTime).getTime();
              item.downloadUrl = (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_8__.getDownLoadAiTaskResultFileUrl)({
                repoOwnerName: task.repoOwnerName,
                repoName: task.repoName,
                id: task.id,
                parent_dir: dir,
                file_name: item.FileName
              });
            });
            list.sort(function (a, b) {
              return a.FileName.localeCompare(b.FileName);
            });
            list.sort(function (a, b) {
              return b.ModTimeNum - a.ModTimeNum;
            });
            _this.filesList = list;
            (_this$$refs$tableRef = _this.$refs['tableRef']) === null || _this$$refs$tableRef === void 0 ? void 0 : _this$$refs$tableRef.clearSort();
          }
        }
      })["catch"](function (err) {
        _this.loading = false;
        console.log(err);
      });
    },
    goNextDir: function goNextDir(item) {
      this.filePath.push({
        label: item.FileName,
        path: item.FileName
      });
      var dir = this.filePath.map(function (item) {
        return item.path;
      }).join('/');
      this.getDirFiles(dir);
    },
    goBackDir: function goBackDir(item) {
      var index = this.filePath.findIndex(function (pth) {
        return item === pth;
      });
      this.filePath = this.filePath.slice(0, index + 1);
      var dir = this.filePath.map(function (item) {
        return item.path;
      }).join('/');
      this.getDirFiles(dir);
    },
    retry: function retry() {
      var _this2 = this;

      var task = this.data.task;
      var experienceParams = {};
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_8__.setAiTaskOutputReschedule)({
        id: task.id
      }).then(function (res) {
        res = res.data;

        if (res.code == 0) {
          _this2.refresh();
        } else {
          _this2.$message({
            type: 'error',
            message: res.msg
          });
        }
      })["catch"](function (err) {
        console.log(err);
      });
    },
    refresh: function refresh() {
      // console.log(this.configs);
      // console.log(this.data);
      var task = this.data.task;
      this.filesList = [];
      var version = task.current_version_name;

      if (version) {
        this.filePath = [{
          label: version,
          path: ""
        }];
        this.getDirFiles("");
      } else {
        this.filePath = [{
          label: 'result',
          path: ''
        }];
        this.getDirFiles('');
      }
    }
  },
  beforeMount: function beforeMount() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=script&lang=js":
/*!***************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=script&lang=js ***!
  \***************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.includes */ "./node_modules/core-js/modules/es.array.includes.js");
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.index-of */ "./node_modules/core-js/modules/es.array.index-of.js");
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.date.now */ "./node_modules/core-js/modules/es.date.now.js");
/* harmony import */ var core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.object.assign */ "./node_modules/core-js/modules/es.object.assign.js");
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! @babel/runtime/helpers/toConsumableArray */ "./node_modules/@babel/runtime/helpers/toConsumableArray.js");
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var _components_NotFound_vue__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ~/components/NotFound.vue */ "./web_src/vuepages/components/NotFound.vue");
/* harmony import */ var _components_cloudbrain_LoadingMask_vue__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ~/components/cloudbrain/LoadingMask.vue */ "./web_src/vuepages/components/cloudbrain/LoadingMask.vue");
/* harmony import */ var _components_cloudbrain_details_ConfigInfo_vue__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! ~/components/cloudbrain/details/ConfigInfo.vue */ "./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue");
/* harmony import */ var _components_cloudbrain_details_OperationProfile_vue__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! ~/components/cloudbrain/details/OperationProfile.vue */ "./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue");
/* harmony import */ var _components_cloudbrain_details_Logs_vue__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! ~/components/cloudbrain/details/Logs.vue */ "./web_src/vuepages/components/cloudbrain/details/Logs.vue");
/* harmony import */ var _components_cloudbrain_details_Loss_vue__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! ~/components/cloudbrain/details/Loss.vue */ "./web_src/vuepages/components/cloudbrain/details/Loss.vue");
/* harmony import */ var _components_cloudbrain_details_ResourceUseage_vue__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! ~/components/cloudbrain/details/ResourceUseage.vue */ "./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue");
/* harmony import */ var _components_cloudbrain_details_ResultDownload_vue__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! ~/components/cloudbrain/details/ResultDownload.vue */ "./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue");
/* harmony import */ var _components_cloudbrain_details_ExportModel_vue__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! ~/components/cloudbrain/details/ExportModel.vue */ "./web_src/vuepages/components/cloudbrain/details/ExportModel.vue");
/* harmony import */ var _components_cloudbrain_details_ExportDataset_vue__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! ~/components/cloudbrain/details/ExportDataset.vue */ "./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue");
/* harmony import */ var _components_cloudbrain_details_EvalOverview_vue__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! ~/components/cloudbrain/details/EvalOverview.vue */ "./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue");
/* harmony import */ var _components_cloudbrain_details_EvalDetail_vue__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! ~/components/cloudbrain/details/EvalDetail.vue */ "./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue");
/* harmony import */ var _Icons_vue__WEBPACK_IMPORTED_MODULE_24__ = __webpack_require__(/*! ./Icons.vue */ "./web_src/vuepages/components/cloudbrain/details/Icons.vue");
/* harmony import */ var _pages_cloudbrain_tools__WEBPACK_IMPORTED_MODULE_25__ = __webpack_require__(/*! ~/pages/cloudbrain/tools */ "./web_src/vuepages/pages/cloudbrain/tools.js");
/* harmony import */ var _pages_tasktmpl_tools__WEBPACK_IMPORTED_MODULE_26__ = __webpack_require__(/*! ~/pages/tasktmpl/tools */ "./web_src/vuepages/pages/tasktmpl/tools.js");
/* harmony import */ var _const__WEBPACK_IMPORTED_MODULE_27__ = __webpack_require__(/*! ~/const */ "./web_src/vuepages/const/index.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_28__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _pages_cloudbrain_configs__WEBPACK_IMPORTED_MODULE_29__ = __webpack_require__(/*! ~/pages/cloudbrain/configs */ "./web_src/vuepages/pages/cloudbrain/configs.js");
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var _apis_modules_common__WEBPACK_IMPORTED_MODULE_31__ = __webpack_require__(/*! ~/apis/modules/common */ "./web_src/vuepages/apis/modules/common.js");
/* harmony import */ var element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_32__ = __webpack_require__(/*! element-ui/lib/utils/date-util */ "./node_modules/element-ui/lib/utils/date-util.js");












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





















var cloudBrainTools = new _pages_cloudbrain_tools__WEBPACK_IMPORTED_MODULE_25__.CloudBrainTools();
var taskTmplTools = new _pages_tasktmpl_tools__WEBPACK_IMPORTED_MODULE_26__.TaskTmplTools();
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'TaskDetailCollapseTabs',
  props: {
    taskId: {
      type: String,
      "default": '',
      required: true
    },
    autoInit: {
      type: Boolean,
      "default": true
    }
  },
  data: function data() {
    return {
      notFound: false,
      pageCfg: {},
      tabNameList: [],
      tabConfigs: {},
      operationList: [],
      state: {},
      collapseValue: '0',
      mainData: [],
      loading: false,
      errorMsgBoxShow: false,
      errorMsg: '',
      maskLoading: false,
      maskLoadingContent: '',
      operating: false,
      taskAlreadyDialogShow: false,
      taskAlreadyDialogShowMsg: ''
    };
  },
  components: {
    NotFound: _components_NotFound_vue__WEBPACK_IMPORTED_MODULE_12__["default"],
    LoadingMask: _components_cloudbrain_LoadingMask_vue__WEBPACK_IMPORTED_MODULE_13__["default"],
    ConfigInfo: _components_cloudbrain_details_ConfigInfo_vue__WEBPACK_IMPORTED_MODULE_14__["default"],
    OperationProfile: _components_cloudbrain_details_OperationProfile_vue__WEBPACK_IMPORTED_MODULE_15__["default"],
    Logs: _components_cloudbrain_details_Logs_vue__WEBPACK_IMPORTED_MODULE_16__["default"],
    ResourceUseage: _components_cloudbrain_details_ResourceUseage_vue__WEBPACK_IMPORTED_MODULE_18__["default"],
    ResultDownload: _components_cloudbrain_details_ResultDownload_vue__WEBPACK_IMPORTED_MODULE_19__["default"],
    ExportModel: _components_cloudbrain_details_ExportModel_vue__WEBPACK_IMPORTED_MODULE_20__["default"],
    ExportDataset: _components_cloudbrain_details_ExportDataset_vue__WEBPACK_IMPORTED_MODULE_21__["default"],
    Loss: _components_cloudbrain_details_Loss_vue__WEBPACK_IMPORTED_MODULE_17__["default"],
    EvalOverview: _components_cloudbrain_details_EvalOverview_vue__WEBPACK_IMPORTED_MODULE_22__["default"],
    EvalDetail: _components_cloudbrain_details_EvalDetail_vue__WEBPACK_IMPORTED_MODULE_23__["default"],
    Icons: _Icons_vue__WEBPACK_IMPORTED_MODULE_24__["default"]
  },
  methods: {
    tabChange: function tabChange(item) {
      var _this = this;

      console.log('tabChange', item);
      this.$nextTick(function () {
        _this.refresh(item);
      });
    },
    refresh: function refresh(item, useRefreshBtn) {
      var activeTab = item.activeName;
      var tabContent = this.$refs[activeTab + '-Ref'];
      console.log("activeTab.indexOf('configInfo-') > -1", activeTab.indexOf('configInfo-') > -1);

      if (activeTab.indexOf('configInfo-') > -1) {
        var task = item.task;
        var taskId = task.id;
        (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.getAiTask)({
          id: taskId
        }).then(function (res) {
          res = res.data;

          if (res.code == 0) {
            Object.assign(task, res.data.task);
            delete res.data.task;
            delete res.data.early_version_list;
            Object.assign(item, res.data);
            Object.assign(task, res.data);
            task.createdFromNow = (0,_utils__WEBPACK_IMPORTED_MODULE_28__.timeSinceUnix)(task.created_unix, Date.now() / 1000);
            task.createTimeStr = (0,element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_32__.formatDate)(new Date(task.created_unix * 1000), 'yyyy-MM-dd HH:mm:ss');

            if (['FINETUNE', 'MODELEXPERIENCE', 'SDFINETUNE', 'ComfyuiExperience', 'EVAL'].includes(task.job_type)) {
              task.sdk_code = '';
            }

            cloudBrainTools.checkRunningLeftTime(task);
            cloudBrainTools.checkOperation(task);
          }
        })["catch"](function (err) {
          console.log(err);
        });
      } else {
        tabContent && tabContent[0] && tabContent[0].refresh && tabContent[0].refresh(useRefreshBtn);
      }
    },
    calcFromNow: function calcFromNow(unix) {
      return (0,_utils__WEBPACK_IMPORTED_MODULE_28__.timeSinceUnix)(unix, Date.now() / 1000);
    },
    dateFormat: function dateFormat(unix) {
      return dayjs(unix * 1000).format('YYYY-MM-DD HH:mm:ss');
    },
    // ops
    opVisual: function opVisual(row) {
      var _this2 = this;

      if (this.operating) return;
      this.operating = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.getAiEndpointUrl)({
        id: row.task.id
      }).then(function (res) {
        _this2.operating = false;
        res = res.data;

        if (res.code == 0) {
          if (res.data && res.data.url) {
            if (!window.open(res.data.url)) {
              window.location.href = res.data.url;
            }
          }
        } else {
          _this2.$message({
            type: 'error',
            message: res.msg
          });
        }
      })["catch"](function (err) {
        _this2.operating = false;
        console.log(err);

        _this2.$message({
          type: 'error',
          message: _this2.$t('operationFailed')
        });
      });
    },
    opReDebug: function opReDebug(row) {
      var _this3 = this;

      if (this.operating) return;
      this.operating = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.getAiTaskRestart)({
        id: row.task.id
      }).then(function (res) {
        _this3.operating = false;
        res = res.data;

        if (res.code == 0) {
          window.location.href = "/cloudbrains/detail/".concat(res.data.id);
        } else if (res.code == 2004) {
          _this3.taskAlreadyDialogShow = true;
          _this3.taskAlreadyDialogShowMsg = res.msg;
        } else {
          _this3.$message({
            type: 'error',
            message: res.msg
          });
        }
      })["catch"](function (err) {
        console.log(err);
        _this3.operating = false;

        _this3.$message({
          type: 'error',
          message: _this3.$t('operationFailed')
        });
      });
    },
    opStop: function opStop(row) {
      var _this4 = this;

      if (this.operating) return;
      this.operating = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.stopAiTask)({
        id: row.task.id
      }).then(function (res) {
        _this4.operating = false;
        res = res.data;

        if (res.code == 0) {
          var data = res.data;
          Object.assign(row.task, data);
          row.task.createdFromNow = _this4.calcFromNow(row.task.created_unix);
          cloudBrainTools.checkOperation(row.task);
        } else {
          _this4.$message({
            type: 'error',
            message: res.msg
          });
        }
      })["catch"](function (err) {
        _this4.operating = false;

        _this4.$message({
          type: 'error',
          message: _this4.$t('operationFailed')
        });
      });
    },
    opOnline: function opOnline(row) {
      var url = '';
      var task = row.task;

      if (task.label_name === "chat") {
        url = "/extension/modelexperience/chat?id=".concat(btoa(task.id), "&modelName=").concat(task.app_name);
      }

      if (row.task.label_name === "sd") {
        url = "/extension/modelexperience/sd?id=".concat(btoa(task.id), "&modelName=").concat(task.app_name);
      }

      if (row.task.label_name === "tts") {
        url = "/extension/modelexperience/tts?id=".concat(btoa(task.id), "&modelName=").concat(task.app_name);
      }

      window.open(url, '_blank');
    },
    opLoraTrain: function opLoraTrain(row) {
      var url = "/modelbase/cv/sft/lora?id=".concat(btoa(row.task.id), "&model=").concat(row.task.app_name);
      window.open(url, '_blank');
    },
    opComfyUI: function opComfyUI(row) {
      (0,_apis_modules_common__WEBPACK_IMPORTED_MODULE_31__.getComfyuiUrl)({
        task_id: row.task.id
      }).then(function (res) {
        window.open(res.data.url, '_blank');
      });
    },
    opAim: function opAim(row) {
      var _this5 = this;

      if (this.operating) return;
      this.operating = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.getAimUrl)({
        job_name: [row.task.display_job_name]
      }).then(function (res) {
        _this5.operating = false;
        res = res.data;

        if (res.code == 0) {
          if (res.data && res.data.url) {
            if (!window.open(res.data.url)) {
              window.location.href = res.data.url;
            }
          }
        } else {
          _this5.$message({
            type: 'error',
            message: res.msg || _this5.$t('operationFailed')
          });
        }
      })["catch"](function (err) {
        _this5.operating = false;
        console.log(err);

        _this5.$message({
          type: 'error',
          message: _this5.$t('operationFailed')
        });
      });
    },
    saveTmpl: function saveTmpl(row) {
      if (this.operating) return;
      window.location.href = "/ai_task_tmpl/create?task=".concat(row.task.id);
    },
    scowFunc: function scowFunc(url, ptoken) {
      var form = document.createElement('form');
      form.style.display = 'none';
      form.action = url;
      form.method = 'post';
      form.target = '_blank';
      var input = document.createElement('input');
      input.type = 'hidden';
      input.name = 'password';
      input.value = ptoken;
      form.appendChild(input);
      document.body.appendChild(form);
      form.submit();
      document.body.removeChild(form);
    },
    deployModel: function deployModel(row) {
      var item = row.task;
      var url = "/modelbase/experience/create?id=".concat(item.id, "&job_name=").concat(item.display_job_name, "&compute_resource=").concat(item.compute_source, "&modelName=").concat(item.app_name);
      window.location.href = url;
    },
    opDebug: function opDebug(row) {
      var _this6 = this;

      if (this.operating) return;
      this.operating = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.getAiTaskDebugUrl)({
        id: row.task.id
      }).then(function (res) {
        _this6.operating = false;
        res = res.data;

        if (res.code == 0) {
          if (res.data && res.data.url) {
            var _res$data;

            if (!!((_res$data = res.data) === null || _res$data === void 0 ? void 0 : _res$data.ptoken)) {
              _this6.scowFunc(res.data.url, res.data.ptoken);

              return;
            }

            if (!window.open(res.data.url)) {
              window.location.href = res.data.url;
            }
          }
        } else {
          _this6.$message({
            type: 'error',
            message: res.msg
          });
        }
      })["catch"](function (err) {
        _this6.operating = false;
        console.log(err);

        _this6.$message({
          type: 'error',
          message: _this6.$t('operationFailed')
        });
      });
    },
    opTensorBoard: function opTensorBoard(row) {
      var _this7 = this;

      if (this.operating) return;
      this.operating = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.getAiTaskVisualizeUrl)({
        id: row.task.id
      }).then(function (res) {
        _this7.operating = false;
        res = res.data;

        if (res.code == 0) {
          if (res.data && res.data.url) {
            if (!window.open(res.data.url)) {
              window.location.href = res.data.url;
            }
          }
        } else {
          _this7.$message({
            type: 'error',
            message: res.msg
          });
        }
      })["catch"](function (err) {
        _this7.operating = false;
        console.log(err);

        _this7.$message({
          type: 'error',
          message: _this7.$t('operationFailed')
        });
      });
    },
    emitUpdate: function emitUpdate() {
      var _this8 = this;

      console.log("emitUpdate");
      this.$nextTick(function () {
        _this8.$emit('update', {
          notFound: _this8.notFound,
          mainData: _this8.mainData,
          pageCfg: _this8.pageCfg
        });
      });
    },
    init: function init() {
      var _this9 = this;

      this.loading = true;
      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_30__.getAiTask)({
        id: this.taskId
      }).then(function (res) {
        _this9.loading = false;
        res = res.data;

        if (res.code == 0 && res.data) {
          (function () {
            var earlyVersionList = (res.data.early_version_list || []).map(function (item) {
              return Object.assign({}, res.data, {
                task: item
              });
            });
            var extraData = Object.assign({}, res.data);
            delete extraData.early_version_list;
            delete extraData.task;
            var data = [res.data].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_11___default()(earlyVersionList));
            var computeSource = '',
                taskType = '',
                cluster = '';
            data.forEach(function (item) {
              var task = item.task;
              Object.assign(task, extraData);
              task.computeSourceShow = task.compute_source == 'GPU' ? 'CPU/GPU' : task.compute_source;
              computeSource = task.compute_source;
              taskType = task.job_type;
              cluster = task.cluster;
              item.activeName = 'configInfo-' + task.id;
              task.clusterName = (0,_utils__WEBPACK_IMPORTED_MODULE_28__.getListValueWithKey)(_const__WEBPACK_IMPORTED_MODULE_27__.CLUSTERS, task.cluster);
              task.jobTypeShow = _const__WEBPACK_IMPORTED_MODULE_27__.BenchmarkTypeList.indexOf(taskType) >= 0 ? _this9.$t('benchmarkTask') : (0,_utils__WEBPACK_IMPORTED_MODULE_28__.getListValueWithKey)(_const__WEBPACK_IMPORTED_MODULE_27__.JOB_TYPE, taskType);
              task.createdFromNow = (0,_utils__WEBPACK_IMPORTED_MODULE_28__.timeSinceUnix)(task.created_unix, Date.now() / 1000);
              task.createTimeStr = (0,element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_32__.formatDate)(new Date(task.created_unix * 1000), 'yyyy-MM-dd HH:mm:ss');
              cloudBrainTools.checkRunningLeftTime(task);
              cloudBrainTools.checkOperation(task);
            });
            console.log(taskType, cluster, computeSource);
            var configs = _pages_cloudbrain_configs__WEBPACK_IMPORTED_MODULE_29__.configDetailManager.getConfig(taskType, cluster, computeSource);
            _this9.pageCfg = configs;
            console.log(_this9.pageCfg);
            var tabs = configs.tabs || [];
            var tabsName = tabs.map(function (item) {
              return item.name;
            });
            var tabConfigs = {};
            tabs.forEach(function (item) {
              tabConfigs[item.name] = item;
            });
            _this9.tabNameList = tabsName;
            _this9.tabConfigs = tabConfigs;
            _this9.operationList = configs.operations || [];

            var _loop = function _loop(i, iLen) {
              var _task = data[i].task;

              if (['PREPARING', 'CONNECTING', 'CREATING', 'WAITING', 'INIT', 'STARTING'].includes(_task.status) && tabsName.indexOf('operationProfile') >= 0) {
                data[i].activeName = 'operationProfile-' + _task.id;
              }

              if (['RUNNING'].includes(_task.status) && tabsName.indexOf('logs') >= 0) {
                data[i].activeName = 'logs-' + _task.id;
              }

              if (data[i].activeName.indexOf('configInfo-') < 0) {
                setTimeout(function () {
                  _this9.tabChange(data[i]);
                }, 80);
              }
            };

            for (var i = 0, iLen = data.length; i < iLen; i++) {
              _loop(i, iLen);
            }

            _this9.mainData = data;
            cloudBrainTools.initRefreshData(_this9.mainData);
          })();
        } else {
          _this9.notFound = true;
        }

        _this9.emitUpdate();
      })["catch"](function (err) {
        _this9.loading = false;
        _this9.notFound = true;
        console.log("xxxxxxxxxxx");
        console.log(err);

        _this9.emitUpdate();
      });
    }
  },
  beforeMount: function beforeMount() {},
  mounted: function mounted() {
    if (this.autoInit) {
      this.init();
    }
  }
});

/***/ }),

/***/ "./web_src/vuepages/apis/modules/modelmanage.js":
/*!******************************************************!*\
  !*** ./web_src/vuepages/apis/modules/modelmanage.js ***!
  \******************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   createModels: function() { return /* binding */ createModels; },
/* harmony export */   deleteModelFile: function() { return /* binding */ deleteModelFile; },
/* harmony export */   getChunks: function() { return /* binding */ getChunks; },
/* harmony export */   getMarkdownPreview: function() { return /* binding */ getMarkdownPreview; },
/* harmony export */   getModelEvolutionMap: function() { return /* binding */ getModelEvolutionMap; },
/* harmony export */   getModelFiles: function() { return /* binding */ getModelFiles; },
/* harmony export */   getModelInfoByName: function() { return /* binding */ getModelInfoByName; },
/* harmony export */   getModelLicenseList: function() { return /* binding */ getModelLicenseList; },
/* harmony export */   getModelMigrateStatus: function() { return /* binding */ getModelMigrateStatus; },
/* harmony export */   getModelMigrateUpdateInfo: function() { return /* binding */ getModelMigrateUpdateInfo; },
/* harmony export */   getModelMigrateUserInfo: function() { return /* binding */ getModelMigrateUserInfo; },
/* harmony export */   getMultipartUrl: function() { return /* binding */ getMultipartUrl; },
/* harmony export */   getNewMultipart: function() { return /* binding */ getNewMultipart; },
/* harmony export */   getProfileModel: function() { return /* binding */ getProfileModel; },
/* harmony export */   getProfileModelPublic: function() { return /* binding */ getProfileModelPublic; },
/* harmony export */   getTrainJobList: function() { return /* binding */ getTrainJobList; },
/* harmony export */   modifyModel: function() { return /* binding */ modifyModel; },
/* harmony export */   modifyModelStatus: function() { return /* binding */ modifyModelStatus; },
/* harmony export */   saveLocalModel: function() { return /* binding */ saveLocalModel; },
/* harmony export */   setCompleteMultipart: function() { return /* binding */ setCompleteMultipart; },
/* harmony export */   setModelMigrate: function() { return /* binding */ setModelMigrate; },
/* harmony export */   setModelMigrateRetry: function() { return /* binding */ setModelMigrateRetry; },
/* harmony export */   setModelMigrateUpdate: function() { return /* binding */ setModelMigrateUpdate; }
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


 // 保存本地模型

var saveLocalModel = function saveLocalModel(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "".concat(data.repo, "/modelmanage/create_local_model"),
    method: 'post',
    headers: {
      'Content-type': 'application/x-www-form-urlencoded'
    },
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_12___default().stringify(data)
  });
}; //创建模型

var createModels = function createModels(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: '/api/v1/aimodel/create',
    method: "post",
    params: {},
    data: data
  });
}; // 修改模型
// data: {id,type,name,version,engine,label,description:}

var modifyModel = function modifyModel(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "".concat(data.repo, "/modelmanage/modify_model"),
    method: 'put',
    headers: {
      'Content-type': 'application/x-www-form-urlencoded'
    },
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_12___default().stringify(data)
  });
};
var modifyModelStatus = function modifyModelStatus(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "".concat(data.repo, "/modelmanage/modify_model_status"),
    method: 'put',
    headers: {
      'Content-type': 'application/x-www-form-urlencoded'
    },
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_12___default().stringify(data)
  });
}; // 求模型信息

var getModelInfoByName = function getModelInfoByName(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: '/api/v1/aimodel',
    method: 'get',
    params: params
  });
}; //或许训练任务列表

var getTrainJobList = function getTrainJobList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "".concat(params.repo, "/modelmanage/query_train_job?repoId=").concat(params.repoId),
    method: 'get',
    params: {},
    data: {}
  });
}; // 求模型中文件列表
// params {repo, ID, parentDir}

var getModelFiles = function getModelFiles(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "".concat(params.repo, "/modelmanage/query_onelevel_modelfile"),
    method: 'get',
    params: params,
    data: {}
  });
}; // 删除模型文件
// params {repo, id, fileName}

var deleteModelFile = function deleteModelFile(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "".concat(params.repo, "/modelmanage/delete_model_file"),
    method: 'delete',
    params: params,
    data: {}
  });
};
/* 文件上传相关 */
// 上传文件1: 获取文件chunks信息
// params: { md5, type: 0-CPU/GPU,1-NPU, file_name, scene: 'model', modeluuid }
// return: uploadID, uuid, uploaded, chunks, attachID, modeluuid, modelName, fileName

var getChunks = function getChunks(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/attachments/model/get_chunks",
    method: 'get',
    params: params,
    data: {}
  });
}; // 上传文件2: 上传新文件
// params: { totalChunkCounts, md5, size, fileType, type, file_name, scene=model, modeluuid=xxxx }
// return: uploadID, uuid

var getNewMultipart = function getNewMultipart(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/attachments/model/new_multipart",
    method: 'get',
    params: params,
    data: {}
  });
}; // 上传文件3: 获取分片上传地址
// params: { uuid, uploadID, size, chunkNumber, type, file_name, scene=model }
// return: url

var getMultipartUrl = function getMultipartUrl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/attachments/model/get_multipart_url",
    method: 'get',
    params: params,
    data: {}
  });
}; // 上传文件4: 完成上传后
// data: { uuid, uploadID, size, type, file_name, dataset_id, description, scene=model, modeluuid=xxxx }

var setCompleteMultipart = function setCompleteMultipart(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/attachments/model/complete_multipart",
    method: 'post',
    headers: {
      'Content-type': 'application/x-www-form-urlencoded'
    },
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_12___default().stringify(data)
  });
}; // markdown预览
// data: { userName, repoName, context, text }

var getMarkdownPreview = function getMarkdownPreview(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/repos/".concat(data.repoOwnerName, "/").concat(data.repoName, "/markdown"),
    method: 'post',
    headers: {
      'Content-type': 'application/x-www-form-urlencoded'
    },
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_12___default().stringify({
      mode: 'gfm',
      context: '',
      text: data.text
    })
  });
}; // 获取模型演化图谱数据

var getModelEvolutionMap = function getModelEvolutionMap(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/aimodel/evolution",
    method: 'get',
    params: {
      aimodel_id: params.aimodel_id
    }
  });
}; // 获取平台模型许可证列表

var getModelLicenseList = function getModelLicenseList() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/dashboard/invitation",
    method: 'get',
    params: {
      filename: 'model/license.json'
    }
  });
}; // 模型迁移新建
// data: { hf_repo_id, version, engine, label, license, description, isPrivate }

var setModelMigrate = function setModelMigrate(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/hf_model/new",
    method: 'post',
    params: {},
    data: _objectSpread({}, data)
  });
}; // 模型迁移查询文件迁移状态
// params: { hf_repo_id }

var getModelMigrateStatus = function getModelMigrateStatus(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/hf_model/status",
    method: 'get',
    params: _objectSpread({}, params)
  });
}; // 模型迁移重试
// data: { hf_repo_id }

var setModelMigrateRetry = function setModelMigrateRetry(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/hf_model/retry",
    method: 'post',
    params: {},
    data: _objectSpread({}, data)
  });
}; // 模型迁移查询迁移者信息
// params: { model_id }

var getModelMigrateUserInfo = function getModelMigrateUserInfo(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/hf_model/get_transfer_user",
    method: 'get',
    params: _objectSpread({}, params)
  });
}; // 迁移模型同步文件对比列表
// hf_repo_id

var getModelMigrateUpdateInfo = function getModelMigrateUpdateInfo(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/hf_model/fetch_update",
    method: 'get',
    params: params,
    data: {}
  });
}; // 迁移模型发起同步
// hf_repo_id

var setModelMigrateUpdate = function setModelMigrateUpdate(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: "/api/v1/hf_model/update",
    method: 'post',
    params: {},
    data: _objectSpread({}, data)
  });
}; //获取个人信息页面中的模型列表

var getProfileModel = function getProfileModel(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: '/api/v1/aimodel/list/accessible',
    method: "get",
    params: params
  });
}; //获取个人信息页面中的模型列表（未登录态使用，仅返回公开模型）

var getProfileModelPublic = function getProfileModelPublic(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_11__["default"])({
    url: '/api/v1/aimodel/list/public',
    method: "get",
    params: params
  });
};

/***/ }),

/***/ "./web_src/vuepages/apis/modules/organization.js":
/*!*******************************************************!*\
  !*** ./web_src/vuepages/apis/modules/organization.js ***!
  \*******************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   getOrgDatasetList: function() { return /* binding */ getOrgDatasetList; },
/* harmony export */   getOrgLabel: function() { return /* binding */ getOrgLabel; },
/* harmony export */   getOrgModelList: function() { return /* binding */ getOrgModelList; },
/* harmony export */   getOrgRepoList: function() { return /* binding */ getOrgRepoList; },
/* harmony export */   getOrgSelectedDatasetList: function() { return /* binding */ getOrgSelectedDatasetList; },
/* harmony export */   getOrgSelectedDatasetSetList: function() { return /* binding */ getOrgSelectedDatasetSetList; },
/* harmony export */   getOrgSelectedModelList: function() { return /* binding */ getOrgSelectedModelList; },
/* harmony export */   getOrgSelectedModelSetList: function() { return /* binding */ getOrgSelectedModelSetList; },
/* harmony export */   getOrgSelectedRepoList: function() { return /* binding */ getOrgSelectedRepoList; },
/* harmony export */   getOrgSelectedRepoSetList: function() { return /* binding */ getOrgSelectedRepoSetList; },
/* harmony export */   getOrgStorageDatasetList: function() { return /* binding */ getOrgStorageDatasetList; },
/* harmony export */   getOrgStorageSummary: function() { return /* binding */ getOrgStorageSummary; },
/* harmony export */   listAiforgeOrgRole: function() { return /* binding */ listAiforgeOrgRole; },
/* harmony export */   listMyOrgUser: function() { return /* binding */ listMyOrgUser; },
/* harmony export */   listRightOrgUser: function() { return /* binding */ listRightOrgUser; },
/* harmony export */   setAiforgeRoleToOrgUser: function() { return /* binding */ setAiforgeRoleToOrgUser; },
/* harmony export */   setOrgSelectedDataset: function() { return /* binding */ setOrgSelectedDataset; },
/* harmony export */   setOrgSelectedModel: function() { return /* binding */ setOrgSelectedModel; },
/* harmony export */   setOrgSelectedRepo: function() { return /* binding */ setOrgSelectedRepo; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.join */ "./node_modules/core-js/modules/es.array.join.js");
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @babel/runtime/helpers/objectWithoutProperties */ "./node_modules/@babel/runtime/helpers/objectWithoutProperties.js");
/* harmony import */ var _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var _service__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../service */ "./web_src/vuepages/apis/service.js");
/* harmony import */ var qs__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! qs */ "./node_modules/qs/lib/index.js");
/* harmony import */ var qs__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(qs__WEBPACK_IMPORTED_MODULE_5__);





 // repos
// 获取组织精选项目列表 orgName

var getOrgSelectedRepoList = function getOrgSelectedRepoList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/org/".concat(params.orgName, "/org_card_repo"),
    method: 'get',
    params: {}
  });
}; // 获取组织精选项目可设置的列表 orgName

var getOrgSelectedRepoSetList = function getOrgSelectedRepoSetList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/org/".concat(params.orgName, "/org_tag/repo_list"),
    method: 'get',
    params: {
      tagId: 1
    }
  });
}; // 设置组织精选项目 orgName,list

var setOrgSelectedRepo = function setOrgSelectedRepo(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/org/".concat(data.orgName, "/org_tag/repo_submit"),
    method: 'post',
    params: {
      tagId: 1
    },
    data: {
      repoList: data.list
    }
  });
}; // 获取组织项目列表 orgName

var getOrgRepoList = function getOrgRepoList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/org/".concat(params.orgName, "/org_list_repo"),
    method: 'get',
    params: {
      q: params.q,
      filter: params.label,
      sort: params.sort,
      page: params.page,
      pageSize: params.pageSize
    }
  });
}; // model
// 获取组织精选模型列表 orgName

var getOrgSelectedModelList = function getOrgSelectedModelList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(params.orgName, "/org_tag/aimodel/list"),
    method: 'get',
    params: {}
  });
}; // 获取组织精选模型可设置的列表 orgName

var getOrgSelectedModelSetList = function getOrgSelectedModelSetList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(params.orgName, "/org_tag/aimodel/available"),
    method: 'get',
    params: {}
  });
}; // 设置组织精选模型

var setOrgSelectedModel = function setOrgSelectedModel(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(data.orgName, "/org_tag/aimodel/submit"),
    method: 'post',
    params: {},
    data: {
      modelids: data.list.join(',')
    }
  });
}; // 获取组织模型列表

var getOrgModelList = function getOrgModelList(_ref) {
  var owner_name = _ref.owner_name,
      getParams = _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_3___default()(_ref, ["owner_name"]);

  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(owner_name, "/aimodel/list"),
    method: 'get',
    params: getParams
  });
}; // dataset
// 获取组织精选数据集列表

var getOrgSelectedDatasetList = function getOrgSelectedDatasetList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(params.orgName, "/org_tag/dataset/list"),
    method: 'get',
    params: {}
  });
}; // 获取组织精选数据集可设置的列表 orgName

var getOrgSelectedDatasetSetList = function getOrgSelectedDatasetSetList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(params.orgName, "/org_tag/dataset/available"),
    method: 'get',
    params: {}
  });
}; // 设置组织精选数据集 orgName, list

var setOrgSelectedDataset = function setOrgSelectedDataset(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(data.orgName, "/org_tag/dataset/submit"),
    method: 'post',
    params: {},
    data: {
      datasetids: data.list.join(',')
    }
  });
}; // 获取组织数据集列表 orgName

var getOrgDatasetList = function getOrgDatasetList(_ref2) {
  var owner_name = _ref2.owner_name,
      getParams = _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_3___default()(_ref2, ["owner_name"]);

  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(owner_name, "/dataset/list"),
    method: 'get',
    params: getParams
  });
}; // 获取组织数据集列表 orgName

var getOrgStorageDatasetList = function getOrgStorageDatasetList(_ref3, type) {
  var owner_name = _ref3.owner_name,
      getParams = _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_3___default()(_ref3, ["owner_name"]);

  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(owner_name, "/storage/").concat(type),
    method: 'get',
    params: getParams
  });
}; //查询组织数据集标签列表

var getOrgLabel = function getOrgLabel(params) {
  var type = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 'dataset';
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(params.orgName, "/").concat(type, "/tags"),
    method: "get",
    params: {}
  });
};
var getOrgStorageSummary = function getOrgStorageSummary(orgName, params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/orgs/".concat(orgName, "/storage/summary"),
    method: 'get',
    params: params
  });
};
var listAiforgeOrgRole = function listAiforgeOrgRole(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/admin/role/org/list",
    method: 'get',
    params: params
  });
};
var listRightOrgUser = function listRightOrgUser(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/admin/role/list_right_org",
    method: 'get',
    params: params
  });
};
var setAiforgeRoleToOrgUser = function setAiforgeRoleToOrgUser(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/admin/role/set_to_org",
    headers: {
      'Content-type': 'application/x-www-form-urlencoded'
    },
    method: 'post',
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_5___default().stringify(data)
  });
}; //提供工作台用户组织

var listMyOrgUser = function listMyOrgUser(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_4__["default"])({
    url: "/api/v1/platform/org",
    method: 'get',
    params: params
  });
};

/***/ }),

/***/ "./web_src/vuepages/apis/modules/storage.js":
/*!**************************************************!*\
  !*** ./web_src/vuepages/apis/modules/storage.js ***!
  \**************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   batchDelStorageDataset: function() { return /* binding */ batchDelStorageDataset; },
/* harmony export */   getStorageDataset: function() { return /* binding */ getStorageDataset; },
/* harmony export */   getStorageSummary: function() { return /* binding */ getStorageSummary; }
/* harmony export */ });
/* harmony import */ var _service__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../service */ "./web_src/vuepages/apis/service.js");

var getStorageSummary = function getStorageSummary(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_0__["default"])({
    url: '/api/v1/storage/summary',
    method: 'get',
    params: params
  });
};
var getStorageDataset = function getStorageDataset(params, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_0__["default"])({
    url: "/api/v1/storage/".concat(type),
    method: 'get',
    params: params
  });
};
var batchDelStorageDataset = function batchDelStorageDataset(params, type) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_0__["default"])({
    url: "/api/v1/".concat(type, "/batch_delete"),
    method: 'post',
    params: params
  });
};

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

/***/ "./web_src/vuepages/pages/model/llms/componenes/markdown-copy.js":
/*!***********************************************************************!*\
  !*** ./web_src/vuepages/pages/model/llms/componenes/markdown-copy.js ***!
  \***********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! spark-md5 */ "./node_modules/spark-md5/spark-md5.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(spark_md5__WEBPACK_IMPORTED_MODULE_7__);









try {
  undefined === window;
  (0,_utils__WEBPACK_IMPORTED_MODULE_6__.initClipboard)('.markdown-it-code-copy');
} catch (_err) {}

function renderCode(origRule, options) {
  return function () {
    for (var _len = arguments.length, args = new Array(_len), _key = 0; _key < _len; _key++) {
      args[_key] = arguments[_key];
    }

    var tokens = args[0],
        idx = args[1];
    var content = tokens[idx].content.replaceAll('"', '&quot;').replaceAll("'", "&apos;");
    var origRendered = origRule.apply(void 0, args);
    if (content.length === 0) return origRendered;
    return "\n<div style=\"position: relative\">\n\t".concat(origRendered, "\n\t<button class=\"ui poping inline up markdown-it-code-copy \" id=\"clipboard-").concat(sparkMD5Hash(content), "\" \n        data-clipboard-text=\"").concat(content, "\" style=\"").concat(options.buttonStyle, "\" title=\"Copy\"\n        data-position=\"top center\" data-variation=\"inverted tiny\" data-success=\"\u590D\u5236\u6210\u529F\"\n        data-content=\"\u590D\u5236\" data-original=\"\u590D\u5236\">\n\t\t<span style=\"").concat(options.iconStyle, "\" class=\"").concat(options.iconClass, "\">").concat(options.element, "</span>\n\t</button>\n</div>\n");
  };
}

function sparkMD5Hash() {
  var str = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : '';
  return spark_md5__WEBPACK_IMPORTED_MODULE_7___default().hash(str) + Math.random().toString().replace('0.', '');
}

var markdownCopy = function markdownCopy(md, options) {
  md.renderer.rules.code_block = renderCode(md.renderer.rules.code_block, options);
  md.renderer.rules.fence = renderCode(md.renderer.rules.fence, options);
};

/* harmony default export */ __webpack_exports__["default"] = (markdownCopy);

/***/ }),

/***/ "./web_src/vuepages/pages/model/llms/componenes/render.js":
/*!****************************************************************!*\
  !*** ./web_src/vuepages/pages/model/llms/componenes/render.js ***!
  \****************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var markdown_it__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! markdown-it */ "./node_modules/markdown-it/index.js");
/* harmony import */ var markdown_it__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(markdown_it__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var markdown_it_katex_gpt__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! markdown-it-katex-gpt */ "./node_modules/markdown-it-katex-gpt/index.js");
/* harmony import */ var _markdown_copy__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./markdown-copy */ "./web_src/vuepages/pages/model/llms/componenes/markdown-copy.js");
// import 'highlight.js/styles/tomorrow-night.css';
 // import 'katex/dist/katex.min.css';




function createRenderer() {
  var renderer = new (markdown_it__WEBPACK_IMPORTED_MODULE_0___default())({}) // .use(require('markdown-it-texmath'), {
  // 	engine: katex,
  // 	...config.texmath
  // })
  // .use(require('markdown-it-mermaid-plugin'))
  .use(__webpack_require__(/*! markdown-it-highlightjs */ "./node_modules/markdown-it-highlightjs/index.js"), {
    auto: true
  }).use(_markdown_copy__WEBPACK_IMPORTED_MODULE_2__["default"], {
    iconClass: "",
    buttonStyle: 'position: absolute; top: 1px; right: 6px; cursor: pointer; outline: none;',
    iconStyle: 'font-size: 16px; opacity: 0.4;',
    element: '<i class="ri-file-copy-line"></i>'
  }).use(markdown_it_katex_gpt__WEBPACK_IMPORTED_MODULE_1__["default"], {
    delimiters: [{
      left: '\\[',
      right: '\\]',
      display: true
    }, {
      left: '\\(',
      right: '\\)',
      display: false
    }, {
      left: '$$',
      right: '$$',
      display: true
    }, {
      left: '$',
      right: '$',
      display: false
    }]
  });
  return renderer;
}

/* harmony default export */ __webpack_exports__["default"] = (createRenderer);

/***/ }),

/***/ "./web_src/vuepages/pages/tasktmpl/tools.js":
/*!**************************************************!*\
  !*** ./web_src/vuepages/pages/tasktmpl/tools.js ***!
  \**************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   TaskTmplTools: function() { return /* binding */ TaskTmplTools; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @babel/runtime/helpers/classCallCheck */ "./node_modules/@babel/runtime/helpers/classCallCheck.js");
/* harmony import */ var _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @babel/runtime/helpers/createClass */ "./node_modules/@babel/runtime/helpers/createClass.js");
/* harmony import */ var _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_4__);





var TaskTmplTools = /*#__PURE__*/function () {
  function TaskTmplTools() {
    _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_3___default()(this, TaskTmplTools);
  }

  _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_4___default()(TaskTmplTools, [{
    key: "transformData",
    value: function transformData(obj) {
      var _obj$Image, _obj$Image2, _obj$Image3;

      var runParameters = (obj.Parameters || []).map(function (item) {
        return {
          label: item.Label || item.label,
          value: item.Value || item.value
        };
      });
      var models = (obj.PretrainModelList || []).filter(function (item) {
        return item.ID && item.ModelName && item.OwnerName && item.ModelAlias;
      }).map(function (item) {
        return {
          id: item.ID,
          name: item.ModelName,
          owner_name: item.OwnerName,
          alias: item.ModelAlias
        };
      });
      var datasets = (obj.DatasetList || []).filter(function (item) {
        return item.ID && item.DatasetName && item.OwnerName && item.DatasetAlias;
      }).map(function (item) {
        return {
          id: item.ID,
          name: item.DatasetName,
          owner_name: item.OwnerName,
          alias: item.DatasetAlias
        };
      }); // const datasets = (obj.DatasetList || [])
      //   .map((item) => {
      //   if (item.ID && item.DatasetName && item.OwnerName && item.DatasetAlias) {
      //     return {
      //       id: item.ID,
      //       name: item.DatasetName,
      //       owner_name:item.OwnerName,
      //       alias: item.DatasetAlias,
      //     };
      //   }
      // });

      var data = {
        name: obj.Name || '',
        descr: obj.Description || '',
        taskType: obj.JobType || '',
        cluster: obj.Cluster || '',
        computeResource: obj.ComputeSource || '',
        networkType: obj.HasInternet == 1 ? 'no_internet' : obj.HasInternet == 2 ? 'has_internet' : 'has_internet',
        visualizeRequired: !!obj.VisualizeRequired,
        spec: '',
        acc_cards_num: obj.AccCardsNum || 0,
        acc_card_type: obj.AccCardType || '',
        cpu_cores: obj.CpuCores || 0,
        mem_gi_b: obj.MemGiB || 0,
        gpu_mem_gi_b: obj.GPUMemGiB || 0,
        share_mem_gi_b: obj.ShareMemGiB || 0,
        image: {
          image_id: ((_obj$Image = obj.Image) === null || _obj$Image === void 0 ? void 0 : _obj$Image.ImageID) || '',
          image_name: ((_obj$Image2 = obj.Image) === null || _obj$Image2 === void 0 ? void 0 : _obj$Image2.ImageName) || '',
          image_url: ((_obj$Image3 = obj.Image) === null || _obj$Image3 === void 0 ? void 0 : _obj$Image3.ImageUrl) || obj.ImageUrl || ''
        },
        model: models,
        dataset: datasets,
        branchName: obj.BranchName || '',
        bootFile: obj.BootFile || '',
        runParameters: runParameters
      };
      return data;
    }
  }, {
    key: "transformDataReverse",
    value: function transformDataReverse(obj) {
      var modelList = (obj.model || []).map(function (item) {
        return {
          ID: item.id,
          ModelName: item.name,
          ModelAlias: item.alias,
          OwnerName: item.owner_name
        };
      });
      var datasetList = (obj.dataset || []).map(function (item) {
        return {
          ID: item.id,
          DatasetName: item.name,
          DatasetAlias: item.alias,
          OwnerName: item.owner_name
        };
      });
      var runParameterList = (obj.runParameters || []).map(function (item) {
        return {
          Label: item.label,
          Value: item.value
        };
      });
      var data = {
        Name: obj.name,
        Description: obj.descr,
        JobType: obj.taskType,
        Cluster: obj.cluster,
        ComputeSource: obj.computeResource,
        HasInternet: obj.networkType == 'no_internet' ? 1 : obj.networkType == 'has_internet' ? 2 : 0,
        VisualizeRequired: obj.visualizeRequired,
        AccCardsNum: obj.acc_cards_num || 0,
        AccCardType: obj.acc_card_type || '',
        CpuCores: obj.cpu_cores || 0,
        MemGiB: obj.mem_gi_b || 0,
        GPUMemGiB: obj.gpu_mem_gi_b || 0,
        ShareMemGiB: obj.share_mem_gi_b || 0,
        Image: {
          ImageID: obj.image.image_id || '',
          ImageName: obj.image.image_name || '',
          ImageUrl: obj.image.image_url
        },
        PretrainModelList: modelList,
        DatasetList: datasetList,
        BranchName: obj.branchName,
        BootFile: obj.bootFile,
        Parameters: runParameterList
      };
      return data;
    }
  }, {
    key: "transformDataTaskToTmpl",
    value: function transformDataTaskToTmpl(obj) {
      var _obj$parameters;

      var modelList = (obj.pretrain_model_list || []).map(function (item) {
        return {
          ID: item.id,
          ModelName: item.name,
          ModelAlias: item.alias,
          OwnerName: item.owner_name
        };
      });
      var datasetList = (obj.dataset_list || []).map(function (item) {
        return {
          ID: item.uuid,
          DatasetName: item.dataset_name,
          DatasetAlias: item.dataset_alias,
          OwnerName: item.owner_name
        };
      });
      var runParameterList = ((obj === null || obj === void 0 ? void 0 : (_obj$parameters = obj.parameters) === null || _obj$parameters === void 0 ? void 0 : _obj$parameters.parameter) || []).map(function (item) {
        return {
          Label: item.label,
          Value: item.value
        };
      });
      var specObj = obj.spec || {};
      var data = {
        Name: obj.name || '',
        Description: obj.descr || '',
        JobType: obj.job_type,
        Cluster: obj.cluster == 'OpenICloudbrainOne' || obj.cluster == 'OpenICloudbrainTwo' ? 'OpenI' : obj.cluster,
        ComputeSource: obj.compute_source,
        HasInternet: obj.has_internet || 2,
        VisualizeRequired: !!obj.visualize_required,
        AccCardsNum: specObj.acc_cards_num || 0,
        AccCardType: specObj.acc_card_type || '',
        CpuCores: specObj.cpu_cores || 0,
        MemGiB: specObj.mem_gi_b || 0,
        GPUMemGiB: specObj.gpu_mem_gi_b || 0,
        ShareMemGiB: specObj.share_mem_gi_b || 0,
        Image: {
          ImageID: obj.image_id || '',
          ImageName: obj.image_id ? obj.image_name : '',
          ImageUrl: obj.image_url
        },
        PretrainModelList: modelList,
        DatasetList: datasetList,
        BranchName: obj.branch_name,
        BootFile: obj.boot_file,
        Parameters: runParameterList
      };
      return data;
    }
  }]);

  return TaskTmplTools;
}();

/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less":
/*!**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less ***!
  \**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=style&index=0&id=36b8b8d0&scoped=true&lang=less":
/*!**********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=style&index=0&id=36b8b8d0&scoped=true&lang=less ***!
  \**********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=style&index=0&id=6de925a6&scoped=true&lang=less":
/*!**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=style&index=0&id=6de925a6&scoped=true&lang=less ***!
  \**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=style&index=0&id=07fd758f&scoped=true&lang=less":
/*!*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=style&index=0&id=07fd758f&scoped=true&lang=less ***!
  \*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=style&index=0&id=88d6d7e8&scoped=true&lang=less":
/*!*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=style&index=0&id=88d6d7e8&scoped=true&lang=less ***!
  \*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=style&index=0&id=23568f18&scoped=true&lang=less":
/*!***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=style&index=0&id=23568f18&scoped=true&lang=less ***!
  \***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=style&index=0&id=0d0d2975&scoped=true&lang=less":
/*!************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=style&index=0&id=0d0d2975&scoped=true&lang=less ***!
  \************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=style&index=0&id=b8c5f634&scoped=true&lang=less":
/*!**********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=style&index=0&id=b8c5f634&scoped=true&lang=less ***!
  \**********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=style&index=0&id=f63d172a&scoped=true&lang=less":
/*!****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=style&index=0&id=f63d172a&scoped=true&lang=less ***!
  \****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=style&index=0&id=77aca1ee&scoped=true&lang=less":
/*!***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=style&index=0&id=77aca1ee&scoped=true&lang=less ***!
  \***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=style&index=0&id=e7b26d3c&scoped=true&lang=less":
/*!***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=style&index=0&id=e7b26d3c&scoped=true&lang=less ***!
  \***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=style&index=0&id=4ea8dafe&scoped=true&lang=less":
/*!***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=style&index=0&id=4ea8dafe&scoped=true&lang=less ***!
  \***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=style&index=0&id=f5065876&scoped=true&lang=less":
/*!*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=style&index=0&id=f5065876&scoped=true&lang=less ***!
  \*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=style&index=0&id=a5b33778&scoped=true&lang=less":
/*!*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=style&index=0&id=a5b33778&scoped=true&lang=less ***!
  \*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=style&index=0&id=570c19a0&scoped=true&lang=less":
/*!*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=style&index=0&id=570c19a0&scoped=true&lang=less ***!
  \*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./web_src/vuepages/components/BaseDialog.vue":
/*!****************************************************!*\
  !*** ./web_src/vuepages/components/BaseDialog.vue ***!
  \****************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _BaseDialog_vue_vue_type_template_id_165bc8c5_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./BaseDialog.vue?vue&type=template&id=165bc8c5&scoped=true */ "./web_src/vuepages/components/BaseDialog.vue?vue&type=template&id=165bc8c5&scoped=true");
/* harmony import */ var _BaseDialog_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./BaseDialog.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/BaseDialog.vue?vue&type=script&lang=js");
/* harmony import */ var _BaseDialog_vue_vue_type_style_index_0_id_165bc8c5_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less */ "./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _BaseDialog_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _BaseDialog_vue_vue_type_template_id_165bc8c5_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _BaseDialog_vue_vue_type_template_id_165bc8c5_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "165bc8c5",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/BaseDialog.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/NotFound.vue":
/*!**************************************************!*\
  !*** ./web_src/vuepages/components/NotFound.vue ***!
  \**************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _NotFound_vue_vue_type_template_id_5f1293bb_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./NotFound.vue?vue&type=template&id=5f1293bb&scoped=true */ "./web_src/vuepages/components/NotFound.vue?vue&type=template&id=5f1293bb&scoped=true");
/* harmony import */ var _NotFound_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./NotFound.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/NotFound.vue?vue&type=script&lang=js");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! !../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");





/* normalize component */
;
var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
  _NotFound_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _NotFound_vue_vue_type_template_id_5f1293bb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _NotFound_vue_vue_type_template_id_5f1293bb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "5f1293bb",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/NotFound.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue":
/*!************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _GeneralTaskCodeTips_vue_vue_type_template_id_36b8b8d0_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./GeneralTaskCodeTips.vue?vue&type=template&id=36b8b8d0&scoped=true */ "./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=template&id=36b8b8d0&scoped=true");
/* harmony import */ var _GeneralTaskCodeTips_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./GeneralTaskCodeTips.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=script&lang=js");
/* harmony import */ var _GeneralTaskCodeTips_vue_vue_type_style_index_0_id_36b8b8d0_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./GeneralTaskCodeTips.vue?vue&type=style&index=0&id=36b8b8d0&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=style&index=0&id=36b8b8d0&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _GeneralTaskCodeTips_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _GeneralTaskCodeTips_vue_vue_type_template_id_36b8b8d0_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _GeneralTaskCodeTips_vue_vue_type_template_id_36b8b8d0_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "36b8b8d0",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/LoadingMask.vue":
/*!****************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/LoadingMask.vue ***!
  \****************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _LoadingMask_vue_vue_type_template_id_6de925a6_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./LoadingMask.vue?vue&type=template&id=6de925a6&scoped=true */ "./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=template&id=6de925a6&scoped=true");
/* harmony import */ var _LoadingMask_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./LoadingMask.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=script&lang=js");
/* harmony import */ var _LoadingMask_vue_vue_type_style_index_0_id_6de925a6_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./LoadingMask.vue?vue&type=style&index=0&id=6de925a6&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=style&index=0&id=6de925a6&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _LoadingMask_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _LoadingMask_vue_vue_type_template_id_6de925a6_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _LoadingMask_vue_vue_type_template_id_6de925a6_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "6de925a6",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/LoadingMask.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue":
/*!***********************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue ***!
  \***********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ConfigInfo_vue_vue_type_template_id_07fd758f_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ConfigInfo.vue?vue&type=template&id=07fd758f&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=template&id=07fd758f&scoped=true");
/* harmony import */ var _ConfigInfo_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ConfigInfo.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=script&lang=js");
/* harmony import */ var _ConfigInfo_vue_vue_type_style_index_0_id_07fd758f_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ConfigInfo.vue?vue&type=style&index=0&id=07fd758f&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=style&index=0&id=07fd758f&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ConfigInfo_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ConfigInfo_vue_vue_type_template_id_07fd758f_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ConfigInfo_vue_vue_type_template_id_07fd758f_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "07fd758f",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue":
/*!***********************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue ***!
  \***********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _EvalDetail_vue_vue_type_template_id_88d6d7e8_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./EvalDetail.vue?vue&type=template&id=88d6d7e8&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=template&id=88d6d7e8&scoped=true");
/* harmony import */ var _EvalDetail_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./EvalDetail.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=script&lang=js");
/* harmony import */ var _EvalDetail_vue_vue_type_style_index_0_id_88d6d7e8_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./EvalDetail.vue?vue&type=style&index=0&id=88d6d7e8&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=style&index=0&id=88d6d7e8&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _EvalDetail_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _EvalDetail_vue_vue_type_template_id_88d6d7e8_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _EvalDetail_vue_vue_type_template_id_88d6d7e8_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "88d6d7e8",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/EvalDetail.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue":
/*!*************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue ***!
  \*************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _EvalOverview_vue_vue_type_template_id_23568f18_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./EvalOverview.vue?vue&type=template&id=23568f18&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=template&id=23568f18&scoped=true");
/* harmony import */ var _EvalOverview_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./EvalOverview.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=script&lang=js");
/* harmony import */ var _EvalOverview_vue_vue_type_style_index_0_id_23568f18_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./EvalOverview.vue?vue&type=style&index=0&id=23568f18&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=style&index=0&id=23568f18&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _EvalOverview_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _EvalOverview_vue_vue_type_template_id_23568f18_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _EvalOverview_vue_vue_type_template_id_23568f18_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "23568f18",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/EvalOverview.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue":
/*!**************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ExportDataset_vue_vue_type_template_id_0d0d2975_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ExportDataset.vue?vue&type=template&id=0d0d2975&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=template&id=0d0d2975&scoped=true");
/* harmony import */ var _ExportDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ExportDataset.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=script&lang=js");
/* harmony import */ var _ExportDataset_vue_vue_type_style_index_0_id_0d0d2975_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ExportDataset.vue?vue&type=style&index=0&id=0d0d2975&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=style&index=0&id=0d0d2975&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ExportDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ExportDataset_vue_vue_type_template_id_0d0d2975_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ExportDataset_vue_vue_type_template_id_0d0d2975_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "0d0d2975",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/ExportDataset.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ExportModel.vue":
/*!************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ExportModel.vue ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ExportModel_vue_vue_type_template_id_b8c5f634_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ExportModel.vue?vue&type=template&id=b8c5f634&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=template&id=b8c5f634&scoped=true");
/* harmony import */ var _ExportModel_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ExportModel.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=script&lang=js");
/* harmony import */ var _ExportModel_vue_vue_type_style_index_0_id_b8c5f634_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ExportModel.vue?vue&type=style&index=0&id=b8c5f634&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=style&index=0&id=b8c5f634&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ExportModel_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ExportModel_vue_vue_type_template_id_b8c5f634_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ExportModel_vue_vue_type_template_id_b8c5f634_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "b8c5f634",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/ExportModel.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Icons.vue":
/*!******************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Icons.vue ***!
  \******************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Icons_vue_vue_type_template_id_f63d172a_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Icons.vue?vue&type=template&id=f63d172a&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=template&id=f63d172a&scoped=true");
/* harmony import */ var _Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Icons.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=script&lang=js");
/* harmony import */ var _Icons_vue_vue_type_style_index_0_id_f63d172a_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Icons.vue?vue&type=style&index=0&id=f63d172a&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=style&index=0&id=f63d172a&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Icons_vue_vue_type_template_id_f63d172a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Icons_vue_vue_type_template_id_f63d172a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "f63d172a",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/Icons.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Logs.vue":
/*!*****************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Logs.vue ***!
  \*****************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Logs_vue_vue_type_template_id_77aca1ee_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Logs.vue?vue&type=template&id=77aca1ee&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=template&id=77aca1ee&scoped=true");
/* harmony import */ var _Logs_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Logs.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=script&lang=js");
/* harmony import */ var _Logs_vue_vue_type_style_index_0_id_77aca1ee_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Logs.vue?vue&type=style&index=0&id=77aca1ee&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=style&index=0&id=77aca1ee&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Logs_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Logs_vue_vue_type_template_id_77aca1ee_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Logs_vue_vue_type_template_id_77aca1ee_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "77aca1ee",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/Logs.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Loss.vue":
/*!*****************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Loss.vue ***!
  \*****************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Loss_vue_vue_type_template_id_e7b26d3c_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Loss.vue?vue&type=template&id=e7b26d3c&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=template&id=e7b26d3c&scoped=true");
/* harmony import */ var _Loss_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Loss.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=script&lang=js");
/* harmony import */ var _Loss_vue_vue_type_style_index_0_id_e7b26d3c_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Loss.vue?vue&type=style&index=0&id=e7b26d3c&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=style&index=0&id=e7b26d3c&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Loss_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Loss_vue_vue_type_template_id_e7b26d3c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Loss_vue_vue_type_template_id_e7b26d3c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "e7b26d3c",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/Loss.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue":
/*!*****************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue ***!
  \*****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _OperationProfile_vue_vue_type_template_id_4ea8dafe_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./OperationProfile.vue?vue&type=template&id=4ea8dafe&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=template&id=4ea8dafe&scoped=true");
/* harmony import */ var _OperationProfile_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./OperationProfile.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=script&lang=js");
/* harmony import */ var _OperationProfile_vue_vue_type_style_index_0_id_4ea8dafe_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./OperationProfile.vue?vue&type=style&index=0&id=4ea8dafe&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=style&index=0&id=4ea8dafe&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _OperationProfile_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _OperationProfile_vue_vue_type_template_id_4ea8dafe_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _OperationProfile_vue_vue_type_template_id_4ea8dafe_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "4ea8dafe",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/OperationProfile.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue":
/*!***************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue ***!
  \***************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ResourceUseage_vue_vue_type_template_id_f5065876_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ResourceUseage.vue?vue&type=template&id=f5065876&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=template&id=f5065876&scoped=true");
/* harmony import */ var _ResourceUseage_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ResourceUseage.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=script&lang=js");
/* harmony import */ var _ResourceUseage_vue_vue_type_style_index_0_id_f5065876_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ResourceUseage.vue?vue&type=style&index=0&id=f5065876&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=style&index=0&id=f5065876&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ResourceUseage_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ResourceUseage_vue_vue_type_template_id_f5065876_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ResourceUseage_vue_vue_type_template_id_f5065876_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "f5065876",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue":
/*!***************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue ***!
  \***************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ResultDownload_vue_vue_type_template_id_a5b33778_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ResultDownload.vue?vue&type=template&id=a5b33778&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=template&id=a5b33778&scoped=true");
/* harmony import */ var _ResultDownload_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ResultDownload.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=script&lang=js");
/* harmony import */ var _ResultDownload_vue_vue_type_style_index_0_id_a5b33778_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ResultDownload.vue?vue&type=style&index=0&id=a5b33778&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=style&index=0&id=a5b33778&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ResultDownload_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ResultDownload_vue_vue_type_template_id_a5b33778_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ResultDownload_vue_vue_type_template_id_a5b33778_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "a5b33778",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/ResultDownload.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue":
/*!***********************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue ***!
  \***********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _TaskDetailCollapseTabs_vue_vue_type_template_id_570c19a0_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./TaskDetailCollapseTabs.vue?vue&type=template&id=570c19a0&scoped=true */ "./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=template&id=570c19a0&scoped=true");
/* harmony import */ var _TaskDetailCollapseTabs_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./TaskDetailCollapseTabs.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=script&lang=js");
/* harmony import */ var _TaskDetailCollapseTabs_vue_vue_type_style_index_0_id_570c19a0_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./TaskDetailCollapseTabs.vue?vue&type=style&index=0&id=570c19a0&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=style&index=0&id=570c19a0&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _TaskDetailCollapseTabs_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _TaskDetailCollapseTabs_vue_vue_type_template_id_570c19a0_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _TaskDetailCollapseTabs_vue_vue_type_template_id_570c19a0_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "570c19a0",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/BaseDialog.vue?vue&type=script&lang=js":
/*!****************************************************************************!*\
  !*** ./web_src/vuepages/components/BaseDialog.vue?vue&type=script&lang=js ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseDialog_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./BaseDialog.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseDialog_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/NotFound.vue?vue&type=script&lang=js":
/*!**************************************************************************!*\
  !*** ./web_src/vuepages/components/NotFound.vue?vue&type=script&lang=js ***!
  \**************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_NotFound_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./NotFound.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/NotFound.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_NotFound_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=script&lang=js":
/*!************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=script&lang=js ***!
  \************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_GeneralTaskCodeTips_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./GeneralTaskCodeTips.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_GeneralTaskCodeTips_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=script&lang=js":
/*!****************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=script&lang=js ***!
  \****************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_LoadingMask_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./LoadingMask.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_LoadingMask_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=script&lang=js":
/*!***********************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=script&lang=js ***!
  \***********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ConfigInfo_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ConfigInfo.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ConfigInfo_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=script&lang=js":
/*!***********************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=script&lang=js ***!
  \***********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalDetail_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./EvalDetail.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalDetail_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=script&lang=js":
/*!*************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=script&lang=js ***!
  \*************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalOverview_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./EvalOverview.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalOverview_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=script&lang=js":
/*!**************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=script&lang=js ***!
  \**************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ExportDataset.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=script&lang=js":
/*!************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=script&lang=js ***!
  \************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportModel_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ExportModel.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportModel_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=script&lang=js":
/*!******************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=script&lang=js ***!
  \******************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Icons.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=script&lang=js":
/*!*****************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=script&lang=js ***!
  \*****************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Logs_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Logs.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Logs_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=script&lang=js":
/*!*****************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=script&lang=js ***!
  \*****************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Loss_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Loss.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Loss_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=script&lang=js":
/*!*****************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=script&lang=js ***!
  \*****************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_OperationProfile_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OperationProfile.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_OperationProfile_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=script&lang=js":
/*!***************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=script&lang=js ***!
  \***************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceUseage_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ResourceUseage.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceUseage_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=script&lang=js":
/*!***************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=script&lang=js ***!
  \***************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ResultDownload_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ResultDownload.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ResultDownload_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=script&lang=js":
/*!***********************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=script&lang=js ***!
  \***********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskDetailCollapseTabs_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./TaskDetailCollapseTabs.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskDetailCollapseTabs_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less":
/*!*************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less ***!
  \*************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseDialog_vue_vue_type_style_index_0_id_165bc8c5_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../node_modules/less-loader/dist/cjs.js!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=style&index=0&id=36b8b8d0&scoped=true&lang=less":
/*!*********************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=style&index=0&id=36b8b8d0&scoped=true&lang=less ***!
  \*********************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_GeneralTaskCodeTips_vue_vue_type_style_index_0_id_36b8b8d0_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../node_modules/less-loader/dist/cjs.js!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./GeneralTaskCodeTips.vue?vue&type=style&index=0&id=36b8b8d0&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=style&index=0&id=36b8b8d0&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=style&index=0&id=6de925a6&scoped=true&lang=less":
/*!*************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=style&index=0&id=6de925a6&scoped=true&lang=less ***!
  \*************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_LoadingMask_vue_vue_type_style_index_0_id_6de925a6_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../node_modules/less-loader/dist/cjs.js!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./LoadingMask.vue?vue&type=style&index=0&id=6de925a6&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=style&index=0&id=6de925a6&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=style&index=0&id=07fd758f&scoped=true&lang=less":
/*!********************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=style&index=0&id=07fd758f&scoped=true&lang=less ***!
  \********************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ConfigInfo_vue_vue_type_style_index_0_id_07fd758f_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ConfigInfo.vue?vue&type=style&index=0&id=07fd758f&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=style&index=0&id=07fd758f&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=style&index=0&id=88d6d7e8&scoped=true&lang=less":
/*!********************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=style&index=0&id=88d6d7e8&scoped=true&lang=less ***!
  \********************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalDetail_vue_vue_type_style_index_0_id_88d6d7e8_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./EvalDetail.vue?vue&type=style&index=0&id=88d6d7e8&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=style&index=0&id=88d6d7e8&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=style&index=0&id=23568f18&scoped=true&lang=less":
/*!**********************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=style&index=0&id=23568f18&scoped=true&lang=less ***!
  \**********************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalOverview_vue_vue_type_style_index_0_id_23568f18_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./EvalOverview.vue?vue&type=style&index=0&id=23568f18&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=style&index=0&id=23568f18&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=style&index=0&id=0d0d2975&scoped=true&lang=less":
/*!***********************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=style&index=0&id=0d0d2975&scoped=true&lang=less ***!
  \***********************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportDataset_vue_vue_type_style_index_0_id_0d0d2975_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ExportDataset.vue?vue&type=style&index=0&id=0d0d2975&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=style&index=0&id=0d0d2975&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=style&index=0&id=b8c5f634&scoped=true&lang=less":
/*!*********************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=style&index=0&id=b8c5f634&scoped=true&lang=less ***!
  \*********************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportModel_vue_vue_type_style_index_0_id_b8c5f634_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ExportModel.vue?vue&type=style&index=0&id=b8c5f634&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=style&index=0&id=b8c5f634&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=style&index=0&id=f63d172a&scoped=true&lang=less":
/*!***************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=style&index=0&id=f63d172a&scoped=true&lang=less ***!
  \***************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_style_index_0_id_f63d172a_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Icons.vue?vue&type=style&index=0&id=f63d172a&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=style&index=0&id=f63d172a&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=style&index=0&id=77aca1ee&scoped=true&lang=less":
/*!**************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=style&index=0&id=77aca1ee&scoped=true&lang=less ***!
  \**************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Logs_vue_vue_type_style_index_0_id_77aca1ee_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Logs.vue?vue&type=style&index=0&id=77aca1ee&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=style&index=0&id=77aca1ee&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=style&index=0&id=e7b26d3c&scoped=true&lang=less":
/*!**************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=style&index=0&id=e7b26d3c&scoped=true&lang=less ***!
  \**************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Loss_vue_vue_type_style_index_0_id_e7b26d3c_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Loss.vue?vue&type=style&index=0&id=e7b26d3c&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=style&index=0&id=e7b26d3c&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=style&index=0&id=4ea8dafe&scoped=true&lang=less":
/*!**************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=style&index=0&id=4ea8dafe&scoped=true&lang=less ***!
  \**************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_OperationProfile_vue_vue_type_style_index_0_id_4ea8dafe_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OperationProfile.vue?vue&type=style&index=0&id=4ea8dafe&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=style&index=0&id=4ea8dafe&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=style&index=0&id=f5065876&scoped=true&lang=less":
/*!************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=style&index=0&id=f5065876&scoped=true&lang=less ***!
  \************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceUseage_vue_vue_type_style_index_0_id_f5065876_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ResourceUseage.vue?vue&type=style&index=0&id=f5065876&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=style&index=0&id=f5065876&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=style&index=0&id=a5b33778&scoped=true&lang=less":
/*!************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=style&index=0&id=a5b33778&scoped=true&lang=less ***!
  \************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ResultDownload_vue_vue_type_style_index_0_id_a5b33778_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ResultDownload.vue?vue&type=style&index=0&id=a5b33778&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=style&index=0&id=a5b33778&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=style&index=0&id=570c19a0&scoped=true&lang=less":
/*!********************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=style&index=0&id=570c19a0&scoped=true&lang=less ***!
  \********************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskDetailCollapseTabs_vue_vue_type_style_index_0_id_570c19a0_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./TaskDetailCollapseTabs.vue?vue&type=style&index=0&id=570c19a0&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=style&index=0&id=570c19a0&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/BaseDialog.vue?vue&type=template&id=165bc8c5&scoped=true":
/*!**********************************************************************************************!*\
  !*** ./web_src/vuepages/components/BaseDialog.vue?vue&type=template&id=165bc8c5&scoped=true ***!
  \**********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseDialog_vue_vue_type_template_id_165bc8c5_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseDialog_vue_vue_type_template_id_165bc8c5_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseDialog_vue_vue_type_template_id_165bc8c5_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./BaseDialog.vue?vue&type=template&id=165bc8c5&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=template&id=165bc8c5&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/NotFound.vue?vue&type=template&id=5f1293bb&scoped=true":
/*!********************************************************************************************!*\
  !*** ./web_src/vuepages/components/NotFound.vue?vue&type=template&id=5f1293bb&scoped=true ***!
  \********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_NotFound_vue_vue_type_template_id_5f1293bb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_NotFound_vue_vue_type_template_id_5f1293bb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_NotFound_vue_vue_type_template_id_5f1293bb_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./NotFound.vue?vue&type=template&id=5f1293bb&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/NotFound.vue?vue&type=template&id=5f1293bb&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=template&id=36b8b8d0&scoped=true":
/*!******************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=template&id=36b8b8d0&scoped=true ***!
  \******************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_GeneralTaskCodeTips_vue_vue_type_template_id_36b8b8d0_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_GeneralTaskCodeTips_vue_vue_type_template_id_36b8b8d0_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_GeneralTaskCodeTips_vue_vue_type_template_id_36b8b8d0_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./GeneralTaskCodeTips.vue?vue&type=template&id=36b8b8d0&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=template&id=36b8b8d0&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=template&id=6de925a6&scoped=true":
/*!**********************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=template&id=6de925a6&scoped=true ***!
  \**********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_LoadingMask_vue_vue_type_template_id_6de925a6_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_LoadingMask_vue_vue_type_template_id_6de925a6_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_LoadingMask_vue_vue_type_template_id_6de925a6_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./LoadingMask.vue?vue&type=template&id=6de925a6&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=template&id=6de925a6&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=template&id=07fd758f&scoped=true":
/*!*****************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=template&id=07fd758f&scoped=true ***!
  \*****************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ConfigInfo_vue_vue_type_template_id_07fd758f_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ConfigInfo_vue_vue_type_template_id_07fd758f_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ConfigInfo_vue_vue_type_template_id_07fd758f_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ConfigInfo.vue?vue&type=template&id=07fd758f&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=template&id=07fd758f&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=template&id=88d6d7e8&scoped=true":
/*!*****************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=template&id=88d6d7e8&scoped=true ***!
  \*****************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalDetail_vue_vue_type_template_id_88d6d7e8_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalDetail_vue_vue_type_template_id_88d6d7e8_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalDetail_vue_vue_type_template_id_88d6d7e8_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./EvalDetail.vue?vue&type=template&id=88d6d7e8&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=template&id=88d6d7e8&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=template&id=23568f18&scoped=true":
/*!*******************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=template&id=23568f18&scoped=true ***!
  \*******************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalOverview_vue_vue_type_template_id_23568f18_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalOverview_vue_vue_type_template_id_23568f18_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_EvalOverview_vue_vue_type_template_id_23568f18_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./EvalOverview.vue?vue&type=template&id=23568f18&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=template&id=23568f18&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=template&id=0d0d2975&scoped=true":
/*!********************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=template&id=0d0d2975&scoped=true ***!
  \********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportDataset_vue_vue_type_template_id_0d0d2975_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportDataset_vue_vue_type_template_id_0d0d2975_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportDataset_vue_vue_type_template_id_0d0d2975_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ExportDataset.vue?vue&type=template&id=0d0d2975&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=template&id=0d0d2975&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=template&id=b8c5f634&scoped=true":
/*!******************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=template&id=b8c5f634&scoped=true ***!
  \******************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportModel_vue_vue_type_template_id_b8c5f634_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportModel_vue_vue_type_template_id_b8c5f634_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ExportModel_vue_vue_type_template_id_b8c5f634_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ExportModel.vue?vue&type=template&id=b8c5f634&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=template&id=b8c5f634&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=template&id=f63d172a&scoped=true":
/*!************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=template&id=f63d172a&scoped=true ***!
  \************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_template_id_f63d172a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_template_id_f63d172a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_template_id_f63d172a_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Icons.vue?vue&type=template&id=f63d172a&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=template&id=f63d172a&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=template&id=77aca1ee&scoped=true":
/*!***********************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=template&id=77aca1ee&scoped=true ***!
  \***********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Logs_vue_vue_type_template_id_77aca1ee_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Logs_vue_vue_type_template_id_77aca1ee_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Logs_vue_vue_type_template_id_77aca1ee_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Logs.vue?vue&type=template&id=77aca1ee&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=template&id=77aca1ee&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=template&id=e7b26d3c&scoped=true":
/*!***********************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=template&id=e7b26d3c&scoped=true ***!
  \***********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Loss_vue_vue_type_template_id_e7b26d3c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Loss_vue_vue_type_template_id_e7b26d3c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Loss_vue_vue_type_template_id_e7b26d3c_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Loss.vue?vue&type=template&id=e7b26d3c&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=template&id=e7b26d3c&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=template&id=4ea8dafe&scoped=true":
/*!***********************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=template&id=4ea8dafe&scoped=true ***!
  \***********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OperationProfile_vue_vue_type_template_id_4ea8dafe_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OperationProfile_vue_vue_type_template_id_4ea8dafe_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OperationProfile_vue_vue_type_template_id_4ea8dafe_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OperationProfile.vue?vue&type=template&id=4ea8dafe&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=template&id=4ea8dafe&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=template&id=f5065876&scoped=true":
/*!*********************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=template&id=f5065876&scoped=true ***!
  \*********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceUseage_vue_vue_type_template_id_f5065876_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceUseage_vue_vue_type_template_id_f5065876_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceUseage_vue_vue_type_template_id_f5065876_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ResourceUseage.vue?vue&type=template&id=f5065876&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=template&id=f5065876&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=template&id=a5b33778&scoped=true":
/*!*********************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=template&id=a5b33778&scoped=true ***!
  \*********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ResultDownload_vue_vue_type_template_id_a5b33778_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ResultDownload_vue_vue_type_template_id_a5b33778_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ResultDownload_vue_vue_type_template_id_a5b33778_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ResultDownload.vue?vue&type=template&id=a5b33778&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=template&id=a5b33778&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=template&id=570c19a0&scoped=true":
/*!*****************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=template&id=570c19a0&scoped=true ***!
  \*****************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskDetailCollapseTabs_vue_vue_type_template_id_570c19a0_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskDetailCollapseTabs_vue_vue_type_template_id_570c19a0_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskDetailCollapseTabs_vue_vue_type_template_id_570c19a0_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./TaskDetailCollapseTabs.vue?vue&type=template&id=570c19a0&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=template&id=570c19a0&scoped=true");


/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=template&id=165bc8c5&scoped=true":
/*!*************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=template&id=165bc8c5&scoped=true ***!
  \*************************************************************************************************************************************************************************************************************************************/
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
    { staticClass: "base-dlg" },
    [
      _c("el-dialog", {
        staticClass: "base-dlg",
        attrs: {
          visible: _vm.dialogShow,
          title: _vm.title,
          width: _vm.width,
          fullscreen: _vm.fullscreen,
          top: _vm.top,
          modal: _vm.modal,
          "modal-append-to-body": _vm.modalAppendToBody,
          "append-to-body": _vm.appendToBody,
          "lock-scroll": _vm.lockScroll,
          "custom-class": _vm.customClass,
          "close-on-click-modal": _vm.closeOnClickModal,
          "close-on-press-escape": _vm.closeOnPressEscape,
          "show-close": _vm.showClose,
          center: _vm.center,
          "destroy-on-close": _vm.destroyOnClose,
          "before-close": _vm.beforeClose
        },
        on: {
          "update:visible": function($event) {
            _vm.dialogShow = $event
          },
          open: _vm.open,
          opened: _vm.opened,
          close: _vm.close,
          closed: _vm.closed
        },
        scopedSlots: _vm._u(
          [
            {
              key: "title",
              fn: function() {
                return [_vm._t("title")]
              },
              proxy: true
            },
            {
              key: "default",
              fn: function() {
                return [_vm._t("default")]
              },
              proxy: true
            },
            {
              key: "footer",
              fn: function() {
                return [_vm._t("footer")]
              },
              proxy: true
            }
          ],
          null,
          true
        )
      })
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/NotFound.vue?vue&type=template&id=5f1293bb&scoped=true":
/*!***********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/NotFound.vue?vue&type=template&id=5f1293bb&scoped=true ***!
  \***********************************************************************************************************************************************************************************************************************************/
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
    {
      staticClass: "ui container center",
      staticStyle: { "min-height": "80%" }
    },
    [
      _c("div", { staticClass: "ui basic very padded segment" }, [
        _c("img", {
          staticClass: "ui centered medium image",
          attrs: { src: "/img/icon-403@2x.png" }
        }),
        _vm._v(" "),
        _c("h2", [_vm._v(_vm._s(_vm.$t("emptyPage")) + " ")]),
        _vm._v(" "),
        _c("p", { domProps: { innerHTML: _vm._s(_vm.$t("emptyPageDescr")) } })
      ])
    ]
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=template&id=36b8b8d0&scoped=true":
/*!*********************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/GeneralTaskCodeTips.vue?vue&type=template&id=36b8b8d0&scoped=true ***!
  \*********************************************************************************************************************************************************************************************************************************************************/
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
    {
      staticClass: "preview-container markdown",
      class: _vm.isTaskDetail ? "task-detail" : ""
    },
    [
      _c("div", { staticClass: "markdown-content" }, [
        _c("div", { staticClass: "content-c" }, [
          _c("p", {
            domProps: {
              innerHTML: _vm._s(_vm.$t("cloudbrainObj.generalTaskSdkCodeTip1"))
            }
          }),
          _vm._v(" "),
          _c("pre", { staticClass: "code-block" }, [
            _c("code", {
              staticClass: "chroma language-text",
              domProps: {
                innerHTML: _vm._s(
                  _vm.$t("cloudbrainObj.generalTaskSdkCodeTip2")
                )
              }
            })
          ]),
          _vm._v(" "),
          _c("p", {
            domProps: {
              innerHTML: _vm._s(_vm.$t("cloudbrainObj.generalTaskSdkCodeTip3"))
            }
          }),
          _vm._v(" "),
          _vm._m(0),
          _vm._v(" "),
          _c("p", {
            domProps: {
              innerHTML: _vm._s(_vm.$t("cloudbrainObj.generalTaskSdkCodeTip4"))
            }
          })
        ])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "segment-line" })
    ]
  )
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("pre", { staticClass: "code-block" }, [
      _c("code", { staticClass: "chroma language-text" }, [
        _vm._v(
          "🎉 You're ready to go live at https://be22fe9f-d140-474c-81f4-30fb068157b5.tunnel.paracloud.com => http://127.0.0.1:7860\n"
        )
      ])
    ])
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=template&id=6de925a6&scoped=true":
/*!*************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/LoadingMask.vue?vue&type=template&id=6de925a6&scoped=true ***!
  \*************************************************************************************************************************************************************************************************************************************************/
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
  return _vm.loading
    ? _c("div", { staticClass: "loading-mask" }, [
        _c("div", { staticClass: "loading-wrap" }, [
          _vm._m(0),
          _vm._v(" "),
          _c("div", { staticClass: "loading-tips" }, [
            _vm._v(" " + _vm._s(_vm.tipsContent))
          ])
        ])
      ])
    : _vm._e()
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "loading-icon" }, [
      _c("div", { staticClass: "rect1" }),
      _vm._v(" "),
      _c("div", { staticClass: "rect2" }),
      _vm._v(" "),
      _c("div", { staticClass: "rect3" }),
      _vm._v(" "),
      _c("div", { staticClass: "rect4" }),
      _vm._v(" "),
      _c("div", { staticClass: "rect5" })
    ])
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=template&id=07fd758f&scoped=true":
/*!********************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ConfigInfo.vue?vue&type=template&id=07fd758f&scoped=true ***!
  \********************************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "item-container" }, [
    _c("div", { staticClass: "layout-wrapper" }, [
      _c("div", { staticClass: "left-column" }, [
        _vm.getCurrentFields("basicFiled").length > 0
          ? _c("div", { staticClass: "field-group" }, [
              _c("div", { staticClass: "field-title" }, [
                _vm._v(_vm._s(_vm.$t("cloudbrainObj.basicInfo")))
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "field-divider" }),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "field-grid" },
                _vm._l(_vm.getCurrentFields("basicFiled"), function(fieldKey) {
                  return _c(
                    "div",
                    { key: fieldKey, staticClass: "field-item" },
                    [
                      _c(
                        "label",
                        { attrs: { title: _vm.renderTitle(fieldKey) } },
                        [_vm._v(_vm._s(_vm.renderTitle(fieldKey)) + ":")]
                      ),
                      _vm._v(" "),
                      _c("span", {
                        staticClass: "nowrap",
                        domProps: {
                          innerHTML: _vm._s(_vm.renderContent(fieldKey))
                        }
                      })
                    ]
                  )
                }),
                0
              )
            ])
          : _vm._e(),
        _vm._v(" "),
        _vm.getCurrentFields("resourceFiled").length > 0
          ? _c("div", { staticClass: "field-group" }, [
              _c("div", { staticClass: "field-title" }, [
                _vm._v(_vm._s(_vm.$t("cloudbrainObj.paramsSetting")))
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "field-divider" }),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "field-grid" },
                _vm._l(_vm.getCurrentFields("resourceFiled"), function(
                  fieldKey
                ) {
                  return _c(
                    "div",
                    { key: fieldKey, staticClass: "field-item" },
                    [
                      _c(
                        "label",
                        { attrs: { title: _vm.renderTitle(fieldKey) } },
                        [_vm._v(_vm._s(_vm.renderTitle(fieldKey)) + ":")]
                      ),
                      _vm._v(" "),
                      _c("span", {
                        staticClass: "nowrap",
                        domProps: {
                          innerHTML: _vm._s(_vm.renderContent(fieldKey))
                        }
                      })
                    ]
                  )
                }),
                0
              )
            ])
          : _vm._e(),
        _vm._v(" "),
        _vm.getCurrentFields("paramsFiled").length > 0
          ? _c("div", { staticClass: "field-group" }, [
              _c("div", { staticClass: "field-title" }, [
                _vm._v(_vm._s(_vm.$t("cloudbrainObj.resourceSetting")))
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "field-divider" }),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "field-grid" },
                _vm._l(_vm.getCurrentFields("paramsFiled"), function(fieldKey) {
                  return _c(
                    "div",
                    { key: fieldKey, staticClass: "field-item" },
                    [
                      _c(
                        "label",
                        { attrs: { title: _vm.renderTitle(fieldKey) } },
                        [_vm._v(_vm._s(_vm.renderTitle(fieldKey)) + ":")]
                      ),
                      _vm._v(" "),
                      _c("span", {
                        staticClass: "nowrap",
                        domProps: {
                          innerHTML: _vm._s(_vm.renderContent(fieldKey))
                        }
                      })
                    ]
                  )
                }),
                0
              )
            ])
          : _vm._e()
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "right-column" }, [
        _vm.getCurrentFields("runstatusFiled").length > 0
          ? _c("div", { staticClass: "field-group" }, [
              _c("div", { staticClass: "field-title" }, [
                _vm._v(_vm._s(_vm.$t("cloudbrainObj.runningStatus")))
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "field-divider" }),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "field-grid" },
                _vm._l(_vm.getCurrentFields("runstatusFiled"), function(
                  fieldKey
                ) {
                  return _c(
                    "div",
                    { key: fieldKey, staticClass: "field-item" },
                    [
                      _c(
                        "label",
                        { attrs: { title: _vm.renderTitle(fieldKey) } },
                        [_vm._v(_vm._s(_vm.renderTitle(fieldKey)) + ":")]
                      ),
                      _vm._v(" "),
                      _c("span", {
                        staticClass: "nowrap",
                        domProps: {
                          innerHTML: _vm._s(_vm.renderContent(fieldKey))
                        }
                      })
                    ]
                  )
                }),
                0
              )
            ])
          : _vm._e(),
        _vm._v(" "),
        _vm.getCurrentFields("showSdkCode").length > 0
          ? _c(
              "div",
              { staticClass: "field-group item-block item-model-code" },
              [
                _c("div", { staticClass: "field-title" }, [
                  _vm._v(_vm._s(_vm.$t("cloudbrainObj.sdkUseWay")))
                ]),
                _vm._v(" "),
                _c("div", { staticClass: "field-grid" }, [
                  _c("div", { staticClass: "code-content-c" }, [
                    _c("div", { staticClass: "code-content" }, [
                      _c("pre", [
                        _c("code", {
                          staticClass: "python hljs",
                          domProps: {
                            innerHTML: _vm._s(
                              _vm.renderHljs(_vm.data.task.sdk_code)
                            )
                          }
                        })
                      ])
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "copy-btn" }, [
                      _c(
                        "a",
                        {
                          staticClass: "ui poping up clipboard",
                          attrs: {
                            href: "javascript:;",
                            id:
                              "clipboard-" +
                              _vm.sparkMD5Hash(_vm.data.task.sdk_code),
                            "data-position": "top center",
                            "data-variation": "inverted tiny",
                            "data-success": _vm.$t("copySuccess"),
                            "data-content": _vm.$t("copy"),
                            "data-original": _vm.$t("copy"),
                            "data-clipboard-text": _vm.data.task.sdk_code
                          }
                        },
                        [
                          _c("i", {
                            staticClass: "copy outline icon",
                            staticStyle: { "font-size": "14px" }
                          })
                        ]
                      )
                    ])
                  ])
                ])
              ]
            )
          : _vm._e(),
        _vm._v(" "),
        _vm.getCurrentFields("generalTaskCodeTips").length > 0 &&
        _vm.data.can_modify
          ? _c(
              "div",
              { staticClass: "field-group item-block item-model-code" },
              [
                _c("div", { staticClass: "field-title" }, [
                  _vm._v(_vm._s(_vm.$t("cloudbrainObj.generalTaskSdkCodeTip0")))
                ]),
                _vm._v(" "),
                _c(
                  "div",
                  { staticClass: "field-grid" },
                  [
                    _c("GeneralTaskCodeTips", { attrs: { isTaskDetail: true } })
                  ],
                  1
                )
              ]
            )
          : _vm._e()
      ])
    ])
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=template&id=88d6d7e8&scoped=true":
/*!********************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalDetail.vue?vue&type=template&id=88d6d7e8&scoped=true ***!
  \********************************************************************************************************************************************************************************************************************************************************/
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
    { staticClass: "chart-container" },
    [
      _vm.evalRunning
        ? _c("div", { staticClass: "wait-wrap" }, [
            _c(
              "svg",
              {
                staticClass: "rotating",
                staticStyle: { "margin-right": "5px" },
                attrs: {
                  xmlns: "http://www.w3.org/2000/svg",
                  viewBox: "0 0 24 24",
                  width: "14",
                  height: "14",
                  fill: "#101010"
                }
              },
              [
                _c("path", {
                  attrs: {
                    d:
                      "M6 4H4V2H20V4H18V6C18 7.61543 17.1838 8.91468 16.1561 9.97667C15.4532 10.703 14.598 11.372 13.7309 12C14.598 12.628 15.4532 13.297 16.1561 14.0233C17.1838 15.0853 18 16.3846 18 18V20H20V22H4V20H6V18C6 16.3846 6.81616 15.0853 7.8439 14.0233C8.54682 13.297 9.40202 12.628 10.2691 12C9.40202 11.372 8.54682 10.703 7.8439 9.97667C6.81616 8.91468 6 7.61543 6 6V4ZM8 4V6C8 6.88457 8.43384 7.71032 9.2811 8.58583C10.008 9.33699 10.9548 10.0398 12 10.7781C13.0452 10.0398 13.992 9.33699 14.7189 8.58583C15.5662 7.71032 16 6.88457 16 6V4H8ZM12 13.2219C10.9548 13.9602 10.008 14.663 9.2811 15.4142C8.43384 16.2897 8 17.1154 8 18V20H16V18C16 17.1154 15.5662 16.2897 14.7189 15.4142C13.992 14.663 13.0452 13.9602 12 13.2219Z"
                  }
                })
              ]
            ),
            _vm._v(" "),
            _c("span", [_vm._v(_vm._s(this.$t("cloudbrainObj.eval_task_ing")))])
          ])
        : [
            _c("div", { staticClass: "tips" }, [
              _vm._v(
                "\n      * " +
                  _vm._s(_vm.$t("modelFinetune.evalTips")) +
                  "\n    "
              )
            ]),
            _vm._v(" "),
            _c(
              "div",
              [
                _c(
                  "el-select",
                  {
                    staticStyle: { margin: "10px 0", color: "#101010" },
                    on: { change: _vm.changeDataset },
                    model: {
                      value: _vm.dataset,
                      callback: function($$v) {
                        _vm.dataset = $$v
                      },
                      expression: "dataset"
                    }
                  },
                  _vm._l(_vm.selectDataset, function(item) {
                    return _c("el-option", {
                      key: item.k,
                      attrs: { value: item.v, label: item.k }
                    })
                  }),
                  1
                ),
                _vm._v(" "),
                _c(
                  "el-select",
                  {
                    staticStyle: { margin: "10px 0", color: "#101010" },
                    on: { change: _vm.changeDataset },
                    model: {
                      value: _vm.resultType,
                      callback: function($$v) {
                        _vm.resultType = $$v
                      },
                      expression: "resultType"
                    }
                  },
                  _vm._l(_vm.resultTypeList, function(item) {
                    return _c("el-option", {
                      key: item.name,
                      attrs: { value: item.v, label: item.k }
                    })
                  }),
                  1
                )
              ],
              1
            ),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "table-wrap" },
              [
                _c(
                  "el-table",
                  {
                    staticStyle: { width: "100%" },
                    attrs: { data: _vm.tableData, fit: true, border: "" }
                  },
                  [
                    _c("el-table-column", {
                      attrs: {
                        prop: "type",
                        label: _vm.$t("repos.dataset"),
                        width: "180"
                      }
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        prop: "raw_input",
                        label: _vm.$t("modelFinetune.evalModelInput"),
                        "min-width": "400"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("div", {
                                staticClass:
                                  "file-view markdown chat raw-content",
                                class: {
                                  "ellipsis-text":
                                    !scope.row.rawExpanded &&
                                    scope.row.rawNeedsExpand
                                },
                                staticStyle: { "font-size": "14px" },
                                attrs: { "data-id": scope.row.id },
                                domProps: {
                                  innerHTML: _vm._s(scope.row.rawContent)
                                }
                              })
                            ]
                          }
                        }
                      ])
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        prop: "predict",
                        label: _vm.$t("modelFinetune.evalModelOutput"),
                        "min-width": "400"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("div", {
                                staticClass:
                                  "file-view markdown chat predict-content",
                                class: {
                                  "ellipsis-text":
                                    !scope.row.predictExpanded &&
                                    scope.row.predictNeedsExpand
                                },
                                staticStyle: { "font-size": "14px" },
                                attrs: { "data-id": scope.row.id },
                                domProps: {
                                  innerHTML: _vm._s(scope.row.predictContent)
                                }
                              }),
                              _vm._v(" "),
                              scope.row.predictNeedsExpand
                                ? _c(
                                    "span",
                                    {
                                      staticClass: "show-more-btn",
                                      on: {
                                        click: function($event) {
                                          $event.stopPropagation()
                                          return _vm.toggleExpand(
                                            scope.row,
                                            "predictExpanded"
                                          )
                                        }
                                      }
                                    },
                                    [
                                      _vm._v(
                                        "\n              " +
                                          _vm._s(
                                            scope.row.predictExpanded
                                              ? _vm.$t("taskTmplObj.collapsed")
                                              : _vm.$t("expandMore")
                                          ) +
                                          "\n            "
                                      )
                                    ]
                                  )
                                : _vm._e()
                            ]
                          }
                        }
                      ])
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        prop: "gold",
                        label: _vm.$t("modelFinetune.standerAnswer"),
                        width: "160"
                      }
                    }),
                    _vm._v(" "),
                    _c("el-table-column", {
                      attrs: {
                        prop: "pred",
                        label: _vm.$t("modelFinetune.evalModelOutAnswer"),
                        width: "160"
                      },
                      scopedSlots: _vm._u([
                        {
                          key: "default",
                          fn: function(scope) {
                            return [
                              _c("div", {
                                staticClass:
                                  "file-view markdown chat pred-content",
                                class: {
                                  "ellipsis-text":
                                    !scope.row.predExpanded &&
                                    scope.row.predNeedsExpand
                                },
                                staticStyle: { "font-size": "12px" },
                                attrs: { "data-id": scope.row.id },
                                domProps: {
                                  innerHTML: _vm._s(scope.row.predContent)
                                }
                              }),
                              _vm._v(" "),
                              scope.row.predNeedsExpand
                                ? _c(
                                    "span",
                                    {
                                      staticClass: "show-more-btn",
                                      on: {
                                        click: function($event) {
                                          $event.stopPropagation()
                                          return _vm.toggleExpand(
                                            scope.row,
                                            "predExpanded"
                                          )
                                        }
                                      }
                                    },
                                    [
                                      _vm._v(
                                        "\n              " +
                                          _vm._s(
                                            scope.row.predExpanded
                                              ? _vm.$t("taskTmplObj.collapsed")
                                              : _vm.$t("expandMore")
                                          ) +
                                          "\n            "
                                      )
                                    ]
                                  )
                                : _vm._e()
                            ]
                          }
                        }
                      ])
                    })
                  ],
                  1
                ),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticStyle: {
                      "text-align": "center",
                      "margin-top": "40px"
                    }
                  },
                  [
                    _c("el-pagination", {
                      attrs: {
                        background: "",
                        "current-page": _vm.currentPage,
                        "page-size": _vm.pageSize,
                        layout: "total, prev, pager, next",
                        total: _vm.totalNum
                      },
                      on: { "current-change": _vm.handleCurrentChange }
                    })
                  ],
                  1
                )
              ],
              1
            )
          ]
    ],
    2
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=template&id=23568f18&scoped=true":
/*!**********************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/EvalOverview.vue?vue&type=template&id=23568f18&scoped=true ***!
  \**********************************************************************************************************************************************************************************************************************************************************/
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
    {
      directives: [
        {
          name: "loading",
          rawName: "v-loading",
          value: _vm.loading,
          expression: "loading"
        }
      ],
      staticClass: "chart-container"
    },
    [
      _vm.evalRunning
        ? _c("div", { staticClass: "wait-wrap" }, [
            _c(
              "svg",
              {
                staticClass: "rotating",
                staticStyle: { "margin-right": "5px" },
                attrs: {
                  xmlns: "http://www.w3.org/2000/svg",
                  viewBox: "0 0 24 24",
                  width: "14",
                  height: "14",
                  fill: "#101010"
                }
              },
              [
                _c("path", {
                  attrs: {
                    d:
                      "M6 4H4V2H20V4H18V6C18 7.61543 17.1838 8.91468 16.1561 9.97667C15.4532 10.703 14.598 11.372 13.7309 12C14.598 12.628 15.4532 13.297 16.1561 14.0233C17.1838 15.0853 18 16.3846 18 18V20H20V22H4V20H6V18C6 16.3846 6.81616 15.0853 7.8439 14.0233C8.54682 13.297 9.40202 12.628 10.2691 12C9.40202 11.372 8.54682 10.703 7.8439 9.97667C6.81616 8.91468 6 7.61543 6 6V4ZM8 4V6C8 6.88457 8.43384 7.71032 9.2811 8.58583C10.008 9.33699 10.9548 10.0398 12 10.7781C13.0452 10.0398 13.992 9.33699 14.7189 8.58583C15.5662 7.71032 16 6.88457 16 6V4H8ZM12 13.2219C10.9548 13.9602 10.008 14.663 9.2811 15.4142C8.43384 16.2897 8 17.1154 8 18V20H16V18C16 17.1154 15.5662 16.2897 14.7189 15.4142C13.992 14.663 13.0452 13.9602 12 13.2219Z"
                  }
                })
              ]
            ),
            _vm._v(" "),
            _c("span", [_vm._v(_vm._s(this.$t("cloudbrainObj.eval_task_ing")))])
          ])
        : [
            _c("div", { staticClass: "tips" }, [
              _vm._v(
                "\n      * " +
                  _vm._s(_vm.$t("modelFinetune.evalTips")) +
                  "\n    "
              )
            ]),
            _vm._v(" "),
            _c("div", { staticClass: "main-wrap" }, [
              _c(
                "div",
                { staticClass: "chart-table" },
                [
                  _c(
                    "el-table",
                    {
                      staticStyle: { width: "100%", "margin-top": "58px" },
                      attrs: { data: _vm.tableData }
                    },
                    [
                      _c("el-table-column", {
                        attrs: {
                          prop: _vm.lang === "zh-CN" ? "class" : "classen",
                          label: _vm.$t("modelFinetune.evalTaskCategory"),
                          width: "260"
                        }
                      }),
                      _vm._v(" "),
                      _c("el-table-column", {
                        attrs: {
                          prop: "name",
                          label: _vm.$t("repos.dataset"),
                          width: "160"
                        }
                      }),
                      _vm._v(" "),
                      _c("el-table-column", {
                        attrs: {
                          prop: "value",
                          label: _vm.$t("modelFinetune.evalScore"),
                          "min-width": "72"
                        }
                      })
                    ],
                    1
                  )
                ],
                1
              ),
              _vm._v(" "),
              _c("div", { staticClass: "chart" }, [
                _c("div", { ref: "chartRef", staticClass: "chart" }),
                _vm._v(" "),
                _c("div", { staticClass: "chart-custom-element" }, [
                  _c("div", { staticClass: "rect" }),
                  _vm._v(" "),
                  _c("span", [_vm._v(_vm._s(_vm.modelName))])
                ])
              ])
            ])
          ]
    ],
    2
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=template&id=0d0d2975&scoped=true":
/*!***********************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportDataset.vue?vue&type=template&id=0d0d2975&scoped=true ***!
  \***********************************************************************************************************************************************************************************************************************************************************/
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
    [
      _c(
        "a",
        {
          staticClass: "operate-btn",
          class: _vm.disabled ? "disabled" : "",
          on: {
            click: function($event) {
              $event.stopPropagation()
              $event.preventDefault()
              _vm.dlgShow = true
            }
          }
        },
        [_vm._v(_vm._s(_vm.$t("cloudbrainObj.exportToDataset")))]
      ),
      _vm._v(" "),
      _c(
        "BaseDialog",
        {
          staticClass: "export-model-dlg base-dlg",
          attrs: {
            visible: _vm.dlgShow,
            title: _vm.$t("cloudbrainObj.exportToDataset"),
            width: "900px",
            modal: true,
            modalAppendToBody: true,
            appendToBody: true,
            "close-on-click-modal": false,
            "show-close": true,
            lockScroll: true,
            "destroy-on-close": false
          },
          on: {
            "update:visible": function($event) {
              _vm.dlgShow = $event
            },
            open: _vm.open,
            closed: _vm.closed
          }
        },
        [
          _c(
            "div",
            {
              staticClass: "dlg-content",
              staticStyle: { "min-width": "900px" }
            },
            [
              _c(
                "div",
                {
                  directives: [
                    {
                      name: "loading",
                      rawName: "v-loading",
                      value: _vm.loading,
                      expression: "loading"
                    }
                  ],
                  staticClass: "row-c"
                },
                [
                  _c("div", {
                    staticClass: "tips",
                    domProps: {
                      innerHTML: _vm._s(
                        _vm.$t("cloudbrainObj.exportDataset.exportDatasetTips")
                      )
                    }
                  }),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass: "row",
                      staticStyle: { "align-items": "flex-start" }
                    },
                    [
                      _c(
                        "div",
                        {
                          staticClass: "r-title",
                          staticStyle: { "margin-top": "5px" }
                        },
                        [
                          _c("label", { staticClass: "required" }, [
                            _vm._v(_vm._s(_vm.$t("datasetObj.dataset_name1")))
                          ])
                        ]
                      ),
                      _vm._v(" "),
                      _c("div", { staticClass: "r-content" }, [
                        _c(
                          "div",
                          { staticClass: "select-dataset" },
                          [
                            _c("el-input", {
                              staticStyle: { flex: "1" },
                              attrs: {
                                size: "medium",
                                maxLength: 255,
                                readonly: ""
                              },
                              model: {
                                value: _vm.state.dataset_name,
                                callback: function($$v) {
                                  _vm.$set(_vm.state, "dataset_name", $$v)
                                },
                                expression: "state.dataset_name"
                              }
                            }),
                            _vm._v(" "),
                            _c(
                              "el-popover",
                              {
                                attrs: {
                                  placement: "right",
                                  width: "500",
                                  "popper-class": "export_dataset",
                                  trigger: "click"
                                },
                                on: { show: _vm.searchModelData },
                                model: {
                                  value: _vm.popoverVisible,
                                  callback: function($$v) {
                                    _vm.popoverVisible = $$v
                                  },
                                  expression: "popoverVisible"
                                }
                              },
                              [
                                _c("div", { staticClass: "dataset-wrap" }, [
                                  _c("div", { staticClass: "dataset-title" }, [
                                    _c("span", [_vm._v("选择数据集")]),
                                    _vm._v(" "),
                                    _c("i", {
                                      staticClass: "el-icon-close",
                                      staticStyle: { cursor: "pointer" },
                                      on: {
                                        click: function($event) {
                                          _vm.popoverVisible = false
                                        }
                                      }
                                    })
                                  ]),
                                  _vm._v(" "),
                                  _c("div", { staticClass: "datatset-main" }, [
                                    _c("div", { staticClass: "model-tabs-c" }, [
                                      _c(
                                        "div",
                                        { staticClass: "model-tabs-w" },
                                        [
                                          _c(
                                            "el-tabs",
                                            {
                                              staticClass: "model-tabs",
                                              on: {
                                                "tab-click": _vm.dlgTabClick
                                              },
                                              model: {
                                                value: _vm.dlgActiveName,
                                                callback: function($$v) {
                                                  _vm.dlgActiveName = $$v
                                                },
                                                expression: "dlgActiveName"
                                              }
                                            },
                                            [
                                              _c("el-tab-pane", {
                                                attrs: {
                                                  label: _vm.$t(
                                                    "cloudbrainObj.ihave"
                                                  ),
                                                  name: "first"
                                                }
                                              }),
                                              _vm._v(" "),
                                              _c("el-tab-pane", {
                                                attrs: {
                                                  label: _vm.$t(
                                                    "cloudbrainObj.iCollaborate"
                                                  ),
                                                  name: "second"
                                                }
                                              })
                                            ],
                                            1
                                          ),
                                          _vm._v(" "),
                                          _c(
                                            "el-input",
                                            {
                                              staticClass: "search-inp",
                                              attrs: {
                                                size: "small",
                                                placeholder: _vm.$t(
                                                  "datasetObj.dataset_search_placeholder"
                                                )
                                              },
                                              nativeOn: {
                                                keydown: function($event) {
                                                  if (
                                                    !$event.type.indexOf(
                                                      "key"
                                                    ) &&
                                                    _vm._k(
                                                      $event.keyCode,
                                                      "enter",
                                                      13,
                                                      $event.key,
                                                      "Enter"
                                                    )
                                                  ) {
                                                    return null
                                                  }
                                                  $event.stopPropagation()
                                                  $event.preventDefault()
                                                  return _vm.inputSearch($event)
                                                }
                                              },
                                              model: {
                                                value: _vm.dlgSearchValue,
                                                callback: function($$v) {
                                                  _vm.dlgSearchValue = $$v
                                                },
                                                expression: "dlgSearchValue"
                                              }
                                            },
                                            [
                                              _c(
                                                "div",
                                                {
                                                  staticClass:
                                                    "search-inp-icon",
                                                  attrs: { slot: "suffix" },
                                                  on: {
                                                    click: _vm.inputSearch
                                                  },
                                                  slot: "suffix"
                                                },
                                                [
                                                  _c("i", {
                                                    staticClass:
                                                      "el-icon-search"
                                                  })
                                                ]
                                              )
                                            ]
                                          )
                                        ],
                                        1
                                      ),
                                      _vm._v(" "),
                                      _c(
                                        "div",
                                        { staticClass: "datalist-wrap" },
                                        [
                                          _c(
                                            "el-radio-group",
                                            {
                                              staticClass: "data-wrap",
                                              on: { input: _vm.changeDataset },
                                              model: {
                                                value: _vm.dataset_id,
                                                callback: function($$v) {
                                                  _vm.dataset_id = $$v
                                                },
                                                expression: "dataset_id"
                                              }
                                            },
                                            _vm._l(
                                              _vm.dlgModelTreeData,
                                              function(item, index) {
                                                return _c(
                                                  "el-radio",
                                                  {
                                                    key: item.id,
                                                    staticClass: "data-item",
                                                    attrs: {
                                                      label: item.id,
                                                      title:
                                                        item.owner_name +
                                                        "/" +
                                                        item.alias
                                                    }
                                                  },
                                                  [
                                                    _c(
                                                      "div",
                                                      {
                                                        staticClass:
                                                          "data-info-c"
                                                      },
                                                      [
                                                        _c(
                                                          "div",
                                                          {
                                                            staticClass:
                                                              "data-info nowrap"
                                                          },
                                                          [
                                                            _c("span", [
                                                              _vm._v(
                                                                " " +
                                                                  _vm._s(
                                                                    item.owner_name
                                                                  ) +
                                                                  " / "
                                                              )
                                                            ]),
                                                            _vm._v(" "),
                                                            _c(
                                                              "span",
                                                              {
                                                                staticStyle: {
                                                                  color:
                                                                    "#101010",
                                                                  "font-weight":
                                                                    "700"
                                                                }
                                                              },
                                                              [
                                                                _vm._v(
                                                                  _vm._s(
                                                                    item.alias
                                                                  )
                                                                )
                                                              ]
                                                            )
                                                          ]
                                                        ),
                                                        _vm._v(" "),
                                                        _c(
                                                          "a",
                                                          {
                                                            attrs: {
                                                              href:
                                                                "/datasets/detail/" +
                                                                item.owner_name +
                                                                "/" +
                                                                item.name,
                                                              target: "_blank"
                                                            }
                                                          },
                                                          [
                                                            _c(
                                                              "svg",
                                                              {
                                                                attrs: {
                                                                  width: "16",
                                                                  height: "16",
                                                                  viewBox:
                                                                    "0 0 48 48",
                                                                  fill: "none",
                                                                  xmlns:
                                                                    "http://www.w3.org/2000/svg"
                                                                }
                                                              },
                                                              [
                                                                _c("path", {
                                                                  attrs: {
                                                                    d:
                                                                      "M28 6H42V20",
                                                                    stroke:
                                                                      "#10101080",
                                                                    "stroke-width":
                                                                      "4",
                                                                    "stroke-linecap":
                                                                      "round",
                                                                    "stroke-linejoin":
                                                                      "round"
                                                                  }
                                                                }),
                                                                _vm._v(" "),
                                                                _c("path", {
                                                                  attrs: {
                                                                    d:
                                                                      "M42 29.4737V39C42 40.6569 40.6569 42 39 42H9C7.34315 42 6 40.6569 6 39V9C6 7.34315 7.34315 6 9 6L18 6",
                                                                    stroke:
                                                                      "#10101080",
                                                                    "stroke-width":
                                                                      "4",
                                                                    "stroke-linecap":
                                                                      "round",
                                                                    "stroke-linejoin":
                                                                      "round"
                                                                  }
                                                                }),
                                                                _vm._v(" "),
                                                                _c("path", {
                                                                  attrs: {
                                                                    d:
                                                                      "M25.7998 22.1999L41.0998 6.8999",
                                                                    stroke:
                                                                      "#10101080",
                                                                    "stroke-width":
                                                                      "4",
                                                                    "stroke-linecap":
                                                                      "round",
                                                                    "stroke-linejoin":
                                                                      "round"
                                                                  }
                                                                })
                                                              ]
                                                            )
                                                          ]
                                                        )
                                                      ]
                                                    )
                                                  ]
                                                )
                                              }
                                            ),
                                            1
                                          )
                                        ],
                                        1
                                      ),
                                      _vm._v(" "),
                                      _c(
                                        "div",
                                        { staticClass: "pagination-c" },
                                        [
                                          _c("el-pagination", {
                                            attrs: {
                                              background: "",
                                              "current-page": _vm.dlgPage,
                                              "page-size": _vm.dlgPageSize,
                                              "pager-count": 5,
                                              layout:
                                                "total, prev, pager, next",
                                              total: _vm.dlgTotal
                                            },
                                            on: {
                                              "current-change":
                                                _vm.dlgPageChange
                                            }
                                          })
                                        ],
                                        1
                                      )
                                    ])
                                  ])
                                ]),
                                _vm._v(" "),
                                _c(
                                  "div",
                                  {
                                    staticClass: "btn-select",
                                    attrs: { slot: "reference" },
                                    slot: "reference"
                                  },
                                  [
                                    _c("i", { staticClass: "el-icon-plus" }),
                                    _vm._v(" "),
                                    _c("span", [
                                      _vm._v(
                                        _vm._s(
                                          _vm.$t("datasetObj.dataset_select")
                                        )
                                      )
                                    ])
                                  ]
                                )
                              ]
                            )
                          ],
                          1
                        )
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  !!_vm.dataset_id
                    ? _c(
                        "div",
                        {
                          staticClass: "row",
                          staticStyle: { margin: "0", "margin-left": "-190px" }
                        },
                        [
                          _c("div", { staticClass: "r-title" }),
                          _vm._v(" "),
                          _c(
                            "div",
                            {
                              staticClass: "r-content",
                              staticStyle: {
                                display: "flex",
                                "line-height": "28px",
                                "flex-wrap": "wrap"
                              }
                            },
                            [
                              _c("span", { staticClass: "storage-t" }, [
                                _vm._v(
                                  "\n              " +
                                    _vm._s(_vm.$t("storage.remain_storage")) +
                                    "："
                                ),
                                _c("span", { staticClass: "storage-v" }, [
                                  _vm._v(
                                    _vm._s(_vm.formattedRemaining[0]) +
                                      "\n              "
                                  )
                                ]),
                                _vm._v(
                                  _vm._s(_vm.formattedRemaining[1]) +
                                    "\n              (" +
                                    _vm._s(
                                      _vm.$t("storage.selected_file_size")
                                    ) +
                                    "：\n            "
                                )
                              ]),
                              _vm._v(" "),
                              _c(
                                "span",
                                {
                                  staticClass: "storage-limit-wrap",
                                  staticStyle: { display: "inline" }
                                },
                                [
                                  _c("span", { staticClass: "select-v" }, [
                                    _vm._v(_vm._s(_vm.formattedTotal[0]))
                                  ]),
                                  _vm._v(
                                    _vm._s(_vm.formattedTotal[1]) +
                                      "\n              "
                                  ),
                                  _vm.isStorageExceeded
                                    ? _c(
                                        "span",
                                        {
                                          staticClass: "limit-tip-wrap",
                                          staticStyle: { display: "inline" }
                                        },
                                        [
                                          _c(
                                            "span",
                                            {
                                              staticClass: "limit-tip",
                                              staticStyle: {
                                                display: "inline-flex"
                                              }
                                            },
                                            [
                                              _c("i", {
                                                staticClass:
                                                  "ri-information-line"
                                              }),
                                              _vm._v(
                                                _vm._s(
                                                  _vm.$t(
                                                    "storage.exceedStorage"
                                                  )
                                                )
                                              )
                                            ]
                                          ),
                                          _vm._v(" "),
                                          !_vm.showOwenerTips
                                            ? _c(
                                                "a",
                                                {
                                                  attrs: { href: "/storages" }
                                                },
                                                [
                                                  _vm._v(
                                                    _vm._s(
                                                      _vm.$t(
                                                        "storage.capacity_details"
                                                      )
                                                    )
                                                  )
                                                ]
                                              )
                                            : _c("span", [
                                                _vm._v(
                                                  _vm._s(
                                                    _vm.$t(
                                                      "storage.owenerTips",
                                                      {
                                                        ownerName: _vm.ownerName
                                                      }
                                                    )
                                                  )
                                                )
                                              ])
                                        ]
                                      )
                                    : _vm._e(),
                                  _vm._v("\n              )\n            ")
                                ]
                              )
                            ]
                          )
                        ]
                      )
                    : _vm._e(),
                  _vm._v(" "),
                  _c("div", { staticClass: "row" }, [
                    _c("div", { staticClass: "r-title" }, [
                      _c("label", { staticClass: "required" }, [
                        _vm._v(
                          _vm._s(
                            _vm.$t(
                              "cloudbrainObj.exportDataset.please_select_output_file"
                            )
                          )
                        )
                      ])
                    ]),
                    _vm._v(" "),
                    _c(
                      "div",
                      {
                        staticClass: "r-content",
                        staticStyle: { display: "flex" }
                      },
                      [
                        _c(
                          "el-popover",
                          {
                            attrs: {
                              placement: "bottom-start",
                              width: "735",
                              trigger: "click",
                              disabled: _vm.uploading
                            }
                          },
                          [
                            _c(
                              "div",
                              { staticClass: "treeContainer" },
                              [
                                _c("el-tree", {
                                  ref: "fileTreeRef",
                                  attrs: {
                                    data: _vm.treeData,
                                    "show-checkbox": "",
                                    "default-expand-all": "",
                                    "node-key": "id",
                                    props: _vm.defaultProps
                                  },
                                  on: { check: _vm.onFileCheckChange },
                                  scopedSlots: _vm._u([
                                    {
                                      key: "default",
                                      fn: function(ref) {
                                        var data = ref.data
                                        return _c(
                                          "span",
                                          {
                                            staticClass: "slot-wrap",
                                            staticStyle: {
                                              display: "flex",
                                              flex: "1"
                                            }
                                          },
                                          [
                                            _c("i", {
                                              staticClass: "icon",
                                              class: data.IsDir
                                                ? "folder"
                                                : "file",
                                              attrs: {
                                                width: "16",
                                                height: "16",
                                                "aria-hidden": "true"
                                              }
                                            }),
                                            _vm._v(" "),
                                            _c("span", [
                                              _vm._v(_vm._s(data.label))
                                            ]),
                                            _vm._v(" "),
                                            !data.isDir
                                              ? _c(
                                                  "span",
                                                  {
                                                    staticStyle: {
                                                      "margin-left": "auto"
                                                    }
                                                  },
                                                  [
                                                    _vm._v(
                                                      _vm._s(
                                                        _vm.formatBytes(
                                                          data.Size
                                                        )[0]
                                                      ) +
                                                        _vm._s(
                                                          _vm.formatBytes(
                                                            data.Size
                                                          )[1]
                                                        )
                                                    )
                                                  ]
                                                )
                                              : _vm._e()
                                          ]
                                        )
                                      }
                                    }
                                  ])
                                })
                              ],
                              1
                            ),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "add-param-btn",
                                attrs: { slot: "reference" },
                                slot: "reference"
                              },
                              [
                                _c(
                                  "a",
                                  {
                                    class: _vm.uploading ? "disabled" : "",
                                    attrs: { href: "javascript:;" }
                                  },
                                  [
                                    _c("i", {
                                      staticClass: "plus square outline icon"
                                    }),
                                    _vm._v(" "),
                                    _c("span", [
                                      _vm._v(
                                        _vm._s(
                                          _vm.$t(
                                            "cloudbrainObj.exportDataset.select_file"
                                          )
                                        )
                                      )
                                    ])
                                  ]
                                )
                              ]
                            )
                          ]
                        )
                      ],
                      1
                    )
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass: "row",
                      staticStyle: { "margin-top": "-2px" }
                    },
                    [
                      _c("div", { staticClass: "r-title" }, [_c("label")]),
                      _vm._v(" "),
                      _c("div", { staticClass: "r-content" }, [
                        _c(
                          "div",
                          { staticClass: "file-item-list" },
                          _vm._l(_vm.selectedData, function(item) {
                            return _c(
                              "div",
                              { key: item.FileName, staticClass: "file-item" },
                              [
                                _c(
                                  "span",
                                  {
                                    staticClass: "file-name",
                                    attrs: { title: item.FileName }
                                  },
                                  [_vm._v(_vm._s(item.FileName))]
                                ),
                                _vm._v(" "),
                                item.statusCode == -99
                                  ? _c("i", {
                                      staticClass: "icon delete icon-delete",
                                      on: {
                                        click: function($event) {
                                          return _vm.removeFile(item)
                                        }
                                      }
                                    })
                                  : _vm._e(),
                                _vm._v(" "),
                                item.statusCode == 0
                                  ? _c("div", { staticClass: "file-status" }, [
                                      _c("i", {
                                        staticClass: "icon el-icon-loading",
                                        staticStyle: {
                                          color: "#21ba45",
                                          "margin-top": "0"
                                        }
                                      }),
                                      _vm._v(" "),
                                      _c("span", [
                                        _vm._v(
                                          _vm._s(
                                            _vm.$t(
                                              "cloudbrainObj.exportDataset.exporting"
                                            )
                                          )
                                        )
                                      ])
                                    ])
                                  : _vm._e(),
                                _vm._v(" "),
                                item.statusCode == -1 || item.statusCode == -2
                                  ? _c(
                                      "div",
                                      { staticClass: "file-status" },
                                      [
                                        _c("i", {
                                          staticClass:
                                            "icon ri-close-circle-line",
                                          staticStyle: { color: "red" }
                                        }),
                                        _vm._v(" "),
                                        _c("span", [
                                          _vm._v(
                                            _vm._s(
                                              _vm.$t(
                                                "cloudbrainObj.exportDataset.export_failed"
                                              )
                                            )
                                          )
                                        ]),
                                        _vm._v(" "),
                                        item.statusCode == -2
                                          ? _c(
                                              "el-tooltip",
                                              {
                                                attrs: {
                                                  placement: "top",
                                                  effect: "dark"
                                                }
                                              },
                                              [
                                                _c("i", {
                                                  staticClass:
                                                    "question circle icon"
                                                }),
                                                _vm._v(" "),
                                                _c(
                                                  "div",
                                                  {
                                                    attrs: { slot: "content" },
                                                    slot: "content"
                                                  },
                                                  [
                                                    _c("div", [
                                                      _vm._v(
                                                        _vm._s(
                                                          _vm.$t(
                                                            "cloudbrainObj.exportDataset.export_has_same_file"
                                                          )
                                                        )
                                                      )
                                                    ])
                                                  ]
                                                )
                                              ]
                                            )
                                          : _vm._e()
                                      ],
                                      1
                                    )
                                  : _vm._e(),
                                _vm._v(" "),
                                item.statusCode == 100
                                  ? _c("div", { staticClass: "file-status" }, [
                                      _c("i", {
                                        staticClass:
                                          "icon ri-checkbox-circle-line",
                                        staticStyle: { color: "#21ba45" }
                                      }),
                                      _vm._v(" "),
                                      _c("span", [
                                        _vm._v(
                                          _vm._s(
                                            _vm.$t(
                                              "cloudbrainObj.exportDataset.export_success"
                                            )
                                          )
                                        )
                                      ])
                                    ])
                                  : _vm._e()
                              ]
                            )
                          }),
                          0
                        )
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass: "row",
                      staticStyle: { "margin-top": "20px" }
                    },
                    [
                      _c("div", { staticClass: "r-title" }, [_c("label")]),
                      _vm._v(" "),
                      _c(
                        "div",
                        { staticClass: "r-content btn-c" },
                        [
                          _c(
                            "el-button",
                            {
                              staticClass: "green",
                              attrs: {
                                size: "medium",
                                disabled: _vm.uploading || _vm.isStorageExceeded
                              },
                              on: { click: _vm.submit }
                            },
                            [_vm._v(_vm._s(_vm.$t("modelManage.confirm")))]
                          ),
                          _vm._v(" "),
                          _c(
                            "el-button",
                            {
                              attrs: { size: "medium" },
                              on: { click: _vm.cancel }
                            },
                            [_vm._v(_vm._s(_vm.$t("modelManage.cancel")))]
                          )
                        ],
                        1
                      )
                    ]
                  )
                ]
              )
            ]
          )
        ]
      )
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=template&id=b8c5f634&scoped=true":
/*!*********************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ExportModel.vue?vue&type=template&id=b8c5f634&scoped=true ***!
  \*********************************************************************************************************************************************************************************************************************************************************/
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
    [
      _c(
        "a",
        {
          staticClass: "operate-btn",
          class: _vm.disabled ? "disabled" : "",
          on: {
            click: function($event) {
              $event.stopPropagation()
              $event.preventDefault()
              _vm.dlgShow = true
            }
          }
        },
        [_vm._v(_vm._s(_vm.$t("cloudbrainObj.saveNewModel")))]
      ),
      _vm._v(" "),
      _c(
        "BaseDialog",
        {
          staticClass: "export-model-dlg base-dlg",
          attrs: {
            visible: _vm.dlgShow,
            title: _vm.$t("cloudbrainObj.saveNewModel"),
            width: "950px",
            modal: true,
            modalAppendToBody: true,
            appendToBody: true,
            "close-on-click-modal": false,
            "show-close": true,
            lockScroll: true,
            "destroy-on-close": false
          },
          on: {
            "update:visible": function($event) {
              _vm.dlgShow = $event
            },
            open: _vm.open,
            closed: _vm.closed
          }
        },
        [
          _c("div", { staticClass: "dlg-content" }, [
            _c(
              "div",
              { staticClass: "row-c", class: _vm.loading ? "disabled" : "" },
              [
                _c("div", { staticClass: "row" }, [
                  _c("div", { staticClass: "r-title" }, [
                    _c("label", { staticClass: "required" }, [
                      _vm._v(_vm._s(_vm.$t("modelManage.trainTask")))
                    ])
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass: "r-content",
                      staticStyle: { display: "flex" }
                    },
                    [
                      _c("el-input", {
                        staticClass: "input-disabled",
                        attrs: { size: "medium", readonly: "" },
                        model: {
                          value: _vm.state.taskName,
                          callback: function($$v) {
                            _vm.$set(_vm.state, "taskName", $$v)
                          },
                          expression: "state.taskName"
                        }
                      })
                    ],
                    1
                  )
                ]),
                _vm._v(" "),
                _c(
                  "div",
                  { staticClass: "row", class: _vm.nameErr ? "error" : "" },
                  [
                    _c("div", { staticClass: "r-title" }, [
                      _c("label", { staticClass: "required" }, [
                        _vm._v(_vm._s(_vm.$t("modelSquare.model_name")))
                      ])
                    ]),
                    _vm._v(" "),
                    _c(
                      "div",
                      { staticClass: "r-content" },
                      [
                        _c("el-input", {
                          attrs: { size: "medium", maxlength: 100 },
                          on: { blur: _vm.checkName },
                          nativeOn: {
                            keydown: function($event) {
                              return _vm.handleKeyDown($event)
                            }
                          },
                          model: {
                            value: _vm.state.name,
                            callback: function($$v) {
                              _vm.$set(_vm.state, "name", $$v)
                            },
                            expression: "state.name"
                          }
                        })
                      ],
                      1
                    )
                  ]
                ),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass: "row",
                    staticStyle: { margin: "-6px 0 0 -190px" }
                  },
                  [
                    _c("div", { staticClass: "r-title" }),
                    _vm._v(" "),
                    _c("div", { staticClass: "r-content" }, [
                      _c(
                        "span",
                        {
                          staticStyle: {
                            "font-size": "12px",
                            color: "#888",
                            "line-height": "1",
                            "margin-top": "0.5em",
                            display: "inline-block"
                          }
                        },
                        [
                          _vm._v(
                            "\n              " +
                              _vm._s(
                                _vm.$t("datasetObj.dataset_name_tooltips")
                              ) +
                              "\n            "
                          )
                        ]
                      )
                    ])
                  ]
                ),
                _vm._v(" "),
                _c(
                  "div",
                  { staticClass: "row", class: _vm.aliasErr ? "error" : "" },
                  [
                    _c("div", { staticClass: "r-title" }, [
                      _c("label", [
                        _vm._v(_vm._s(_vm.$t("modelSquare.model_zh_name")))
                      ])
                    ]),
                    _vm._v(" "),
                    _c(
                      "div",
                      { staticClass: "r-content" },
                      [
                        _c("el-input", {
                          attrs: { size: "medium", maxLength: 100 },
                          on: { blur: _vm.checkAlias },
                          nativeOn: {
                            keydown: function($event) {
                              return _vm.handleKeyDown($event)
                            }
                          },
                          model: {
                            value: _vm.state.alias,
                            callback: function($$v) {
                              _vm.$set(_vm.state, "alias", $$v)
                            },
                            expression: "state.alias"
                          }
                        })
                      ],
                      1
                    )
                  ]
                ),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass: "row",
                    staticStyle: { margin: "-6px 0 0 -190px" }
                  },
                  [
                    _c("div", { staticClass: "r-title" }),
                    _vm._v(" "),
                    _c("div", { staticClass: "r-content" }, [
                      _c(
                        "span",
                        {
                          staticStyle: {
                            "font-size": "12px",
                            color: "#888",
                            "line-height": "1",
                            "margin-top": "0.5em",
                            display: "inline-block"
                          }
                        },
                        [
                          _vm._v(
                            "\n              " +
                              _vm._s(
                                _vm.$t("datasetObj.dataset_name_cn_tooltips")
                              ) +
                              "\n            "
                          )
                        ]
                      )
                    ])
                  ]
                ),
                _vm._v(" "),
                _c("div", { staticClass: "row" }, [
                  _c("div", { staticClass: "r-title" }, [
                    _c("label", { staticClass: "required" }, [
                      _vm._v(_vm._s(_vm.$t("datasetObj.dataset_owner")))
                    ])
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "r-content" },
                    [
                      _c(
                        "el-select",
                        {
                          staticStyle: { width: "312px" },
                          attrs: {
                            placeholder: _vm.$t("datasetObj.select_category")
                          },
                          model: {
                            value: _vm.state.owner_id,
                            callback: function($$v) {
                              _vm.$set(_vm.state, "owner_id", $$v)
                            },
                            expression: "state.owner_id"
                          }
                        },
                        _vm._l(_vm.owners, function(item) {
                          return _c("el-option", {
                            key: item.ID,
                            attrs: { value: item.ID, label: item.Name }
                          })
                        }),
                        1
                      )
                    ],
                    1
                  )
                ]),
                _vm._v(" "),
                _c("div", { staticClass: "row" }, [
                  _c("div", { staticClass: "r-title" }, [
                    _c("label", { staticClass: "required" }, [
                      _vm._v(_vm._s(_vm.$t("modelManage.modelEngine")))
                    ])
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "r-content" },
                    [
                      _c(
                        "el-select",
                        {
                          staticStyle: { width: "312px" },
                          attrs: { size: "medium", placeholder: "" },
                          model: {
                            value: _vm.state.engine,
                            callback: function($$v) {
                              _vm.$set(_vm.state, "engine", $$v)
                            },
                            expression: "state.engine"
                          }
                        },
                        _vm._l(_vm.engineList, function(item) {
                          return _c("el-option", {
                            key: item.k,
                            attrs: { label: item.v, value: item.k }
                          })
                        }),
                        1
                      )
                    ],
                    1
                  )
                ]),
                _vm._v(" "),
                _c("div", { staticClass: "row" }, [
                  _c("div", { staticClass: "r-title" }, [
                    _c("label", [_vm._v(_vm._s(_vm.$t("modelManage.license")))])
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "r-content" },
                    [
                      _c(
                        "el-select",
                        {
                          staticClass: "license-sel",
                          staticStyle: { width: "100%" },
                          attrs: {
                            size: "medium",
                            placeholder: _vm.$t("modelManage.selectLicense")
                          },
                          model: {
                            value: _vm.state.license,
                            callback: function($$v) {
                              _vm.$set(_vm.state, "license", $$v)
                            },
                            expression: "state.license"
                          }
                        },
                        [
                          _c("template", { slot: "prefix" }, [
                            _vm.state.license
                              ? _c("i", {
                                  staticClass:
                                    "el-select__caret el-input__icon el-icon-close",
                                  on: {
                                    click: function($event) {
                                      $event.stopPropagation()
                                      $event.preventDefault()
                                      return _vm.handleClearLicenseClick($event)
                                    }
                                  }
                                })
                              : _vm._e()
                          ]),
                          _vm._v(" "),
                          _vm._l(_vm.licenseList, function(item) {
                            return _c("el-option", {
                              key: item.id,
                              attrs: { label: item.name, value: item.id }
                            })
                          })
                        ],
                        2
                      )
                    ],
                    1
                  )
                ]),
                _vm._v(" "),
                _c("div", { staticClass: "row" }, [
                  _c("div", { staticClass: "r-title" }, [
                    _c("label", [
                      _vm._v(_vm._s(_vm.$t("modelManage.modelLabel")))
                    ])
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "r-content" },
                    [
                      _c("el-input", {
                        attrs: {
                          size: "medium",
                          maxLength: 255,
                          placeholder: _vm.$t("modelManage.modelLabelInputTips")
                        },
                        on: { input: _vm.labelInput },
                        model: {
                          value: _vm.state.label,
                          callback: function($$v) {
                            _vm.$set(_vm.state, "label", $$v)
                          },
                          expression: "state.label"
                        }
                      })
                    ],
                    1
                  )
                ]),
                _vm._v(" "),
                _c("div", { staticClass: "row" }, [
                  _c("div", { staticClass: "r-title" }, [
                    _c("label", [
                      _vm._v(_vm._s(_vm.$t("modelManage.modelAccess")))
                    ])
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "r-content" },
                    [
                      _c(
                        "el-radio",
                        {
                          attrs: { label: "0" },
                          model: {
                            value: _vm.state.isPrivate,
                            callback: function($$v) {
                              _vm.$set(_vm.state, "isPrivate", $$v)
                            },
                            expression: "state.isPrivate"
                          }
                        },
                        [
                          _vm._v(
                            _vm._s(_vm.$t("modelManage.modelAccessPublic"))
                          )
                        ]
                      ),
                      _vm._v(" "),
                      _c(
                        "el-radio",
                        {
                          attrs: { label: "1" },
                          model: {
                            value: _vm.state.isPrivate,
                            callback: function($$v) {
                              _vm.$set(_vm.state, "isPrivate", $$v)
                            },
                            expression: "state.isPrivate"
                          }
                        },
                        [
                          _vm._v(
                            _vm._s(_vm.$t("modelManage.modelAccessPrivate"))
                          )
                        ]
                      )
                    ],
                    1
                  )
                ]),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass: "row",
                    staticStyle: { margin: "0", "margin-left": "-190px" }
                  },
                  [
                    _c("div", { staticClass: "r-title" }),
                    _vm._v(" "),
                    _c(
                      "div",
                      {
                        staticClass: "r-content",
                        staticStyle: {
                          display: "flex",
                          "line-height": "28px",
                          "flex-wrap": "wrap"
                        }
                      },
                      [
                        _c("div", { staticClass: "storage-t" }, [
                          _vm._v(
                            "\n              " +
                              _vm._s(_vm.$t("storage.remain_storage")) +
                              "："
                          ),
                          _c("span", { staticClass: "storage-v" }, [
                            _vm._v(_vm._s(_vm.formattedRemaining[0]))
                          ]),
                          _vm._v(
                            "\n              " +
                              _vm._s(_vm.formattedRemaining[1]) +
                              "\n              (" +
                              _vm._s(_vm.$t("storage.selected_file_size")) +
                              "：\n            "
                          )
                        ]),
                        _vm._v(" "),
                        _c("div", { staticClass: "storage-limit-wrap" }, [
                          _c("span", { staticClass: "select-v" }, [
                            _vm._v(_vm._s(_vm.formattedTotal[0]))
                          ]),
                          _vm._v(
                            _vm._s(_vm.formattedTotal[1]) + "\n              "
                          ),
                          _vm.isStorageExceeded
                            ? _c("div", { staticClass: "limit-tip-wrap" }, [
                                _c("div", { staticClass: "limit-tip" }, [
                                  _c("i", {
                                    staticClass: "ri-information-line"
                                  }),
                                  _vm._v(
                                    _vm._s(_vm.$t("storage.exceedStorage"))
                                  )
                                ]),
                                _vm._v(" "),
                                _c("span", [
                                  _vm._v(
                                    _vm._s(
                                      _vm.$t("storage.owenerTips", {
                                        ownerName: _vm.orgName
                                      })
                                    )
                                  )
                                ])
                              ])
                            : _vm._e(),
                          _vm._v("\n              )\n            ")
                        ])
                      ]
                    )
                  ]
                ),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass: "row",
                    class: _vm.modelFileErr ? "error" : ""
                  },
                  [
                    _c("div", { staticClass: "r-title" }, [
                      _c("label", { staticClass: "required" }, [
                        _vm._v(_vm._s(_vm.$t("modelManage.modelFiles")))
                      ])
                    ]),
                    _vm._v(" "),
                    _c(
                      "div",
                      { staticClass: "r-content" },
                      [
                        _c(
                          "el-popover",
                          {
                            attrs: {
                              placement: "bottom",
                              width: "508",
                              trigger: "click"
                            }
                          },
                          [
                            _c(
                              "div",
                              { staticClass: "treeContainer" },
                              [
                                _c("el-tree", {
                                  ref: "fileTreeRef",
                                  attrs: {
                                    data: _vm.treeData,
                                    "show-checkbox": "",
                                    "default-expand-all": "",
                                    "node-key": "id",
                                    props: _vm.defaultProps
                                  },
                                  on: { check: _vm.onFileCheckChange },
                                  scopedSlots: _vm._u([
                                    {
                                      key: "default",
                                      fn: function(ref) {
                                        var data = ref.data
                                        return _c(
                                          "span",
                                          {
                                            staticClass: "slot-wrap",
                                            staticStyle: {
                                              display: "flex",
                                              flex: "1"
                                            }
                                          },
                                          [
                                            _c("i", {
                                              staticClass: "icon",
                                              class: data.isDir
                                                ? "folder"
                                                : "file",
                                              attrs: {
                                                width: "16",
                                                height: "16",
                                                "aria-hidden": "true"
                                              }
                                            }),
                                            _vm._v(" "),
                                            _c("span", [
                                              _vm._v(_vm._s(data.label))
                                            ]),
                                            _vm._v(" "),
                                            !data.isDir
                                              ? _c(
                                                  "span",
                                                  {
                                                    staticStyle: {
                                                      "margin-left": "auto"
                                                    }
                                                  },
                                                  [
                                                    _vm._v(
                                                      _vm._s(
                                                        _vm.formatBytes(
                                                          data.Size
                                                        )[0]
                                                      ) +
                                                        _vm._s(
                                                          _vm.formatBytes(
                                                            data.Size
                                                          )[1]
                                                        )
                                                    )
                                                  ]
                                                )
                                              : _vm._e()
                                          ]
                                        )
                                      }
                                    }
                                  ])
                                })
                              ],
                              1
                            ),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "add-param-btn",
                                attrs: { slot: "reference" },
                                slot: "reference"
                              },
                              [
                                _c("a", { attrs: { href: "javascript:;" } }, [
                                  _c("i", {
                                    staticClass: "plus square outline icon"
                                  }),
                                  _vm._v(" "),
                                  _c("span", [
                                    _vm._v(
                                      _vm._s(
                                        _vm.$t(
                                          "cloudbrainObj.exportDataset.select_file"
                                        )
                                      )
                                    )
                                  ])
                                ])
                              ]
                            )
                          ]
                        )
                      ],
                      1
                    )
                  ]
                ),
                _vm._v(" "),
                _c(
                  "div",
                  { staticClass: "row", staticStyle: { "margin-top": "-2px" } },
                  [
                    _c("div", { staticClass: "r-title" }, [_c("label")]),
                    _vm._v(" "),
                    _c("div", { staticClass: "r-content" }, [
                      _c(
                        "div",
                        { staticClass: "file-item-list" },
                        _vm._l(_vm.selectedData, function(item) {
                          return _c(
                            "div",
                            { key: item.FileName, staticClass: "file-item" },
                            [
                              _c(
                                "span",
                                {
                                  staticClass: "file-name",
                                  attrs: { title: item.FileName }
                                },
                                [_vm._v(_vm._s(item.FileName))]
                              ),
                              _vm._v(" "),
                              item.statusCode == -99
                                ? _c("i", {
                                    staticClass: "icon delete icon-delete",
                                    on: {
                                      click: function($event) {
                                        return _vm.removeFile(item)
                                      }
                                    }
                                  })
                                : _vm._e(),
                              _vm._v(" "),
                              item.statusCode == 0
                                ? _c("div", { staticClass: "file-status" }, [
                                    _c("i", {
                                      staticClass: "icon el-icon-loading",
                                      staticStyle: {
                                        color: "#21ba45",
                                        "margin-top": "0"
                                      }
                                    }),
                                    _vm._v(" "),
                                    _c("span", [
                                      _vm._v(
                                        _vm._s(
                                          _vm.$t(
                                            "cloudbrainObj.exportDataset.exporting"
                                          )
                                        )
                                      )
                                    ])
                                  ])
                                : _vm._e(),
                              _vm._v(" "),
                              item.statusCode == -1 ||
                              item.statusCode == -2 ||
                              item.statusCode == -3
                                ? _c(
                                    "div",
                                    { staticClass: "file-status" },
                                    [
                                      _c("i", {
                                        staticClass:
                                          "icon ri-close-circle-line",
                                        staticStyle: { color: "red" }
                                      }),
                                      _vm._v(" "),
                                      _c("span", [
                                        _vm._v(
                                          _vm._s(
                                            _vm.$t(
                                              "cloudbrainObj.exportDataset.export_failed"
                                            )
                                          )
                                        )
                                      ]),
                                      _vm._v(" "),
                                      item.statusCode == -2
                                        ? _c(
                                            "el-tooltip",
                                            {
                                              attrs: {
                                                placement: "top",
                                                effect: "dark"
                                              }
                                            },
                                            [
                                              _c("i", {
                                                staticClass:
                                                  "question circle icon"
                                              }),
                                              _vm._v(" "),
                                              _c(
                                                "div",
                                                {
                                                  attrs: { slot: "content" },
                                                  slot: "content"
                                                },
                                                [
                                                  _c("div", [
                                                    _vm._v(
                                                      _vm._s(
                                                        _vm.$t(
                                                          "cloudbrainObj.exportDataset.export_has_same_file1"
                                                        )
                                                      )
                                                    )
                                                  ])
                                                ]
                                              )
                                            ]
                                          )
                                        : _vm._e(),
                                      _vm._v(" "),
                                      item.statusCode == -3
                                        ? _c(
                                            "el-tooltip",
                                            {
                                              attrs: {
                                                placement: "top",
                                                effect: "dark"
                                              }
                                            },
                                            [
                                              _c("i", {
                                                staticClass:
                                                  "question circle icon"
                                              }),
                                              _vm._v(" "),
                                              _c(
                                                "div",
                                                {
                                                  attrs: { slot: "content" },
                                                  slot: "content"
                                                },
                                                [
                                                  _c("div", [
                                                    _vm._v(
                                                      _vm._s(
                                                        _vm.$t(
                                                          "cloudbrainObj.exportDataset.export_exceed_storage"
                                                        )
                                                      )
                                                    )
                                                  ])
                                                ]
                                              )
                                            ]
                                          )
                                        : _vm._e()
                                    ],
                                    1
                                  )
                                : _vm._e(),
                              _vm._v(" "),
                              item.statusCode == 100
                                ? _c("div", { staticClass: "file-status" }, [
                                    _c("i", {
                                      staticClass:
                                        "icon ri-checkbox-circle-line",
                                      staticStyle: { color: "#21ba45" }
                                    }),
                                    _vm._v(" "),
                                    _c("span", [
                                      _vm._v(
                                        _vm._s(
                                          _vm.$t(
                                            "cloudbrainObj.exportDataset.export_success"
                                          )
                                        )
                                      )
                                    ])
                                  ])
                                : _vm._e()
                            ]
                          )
                        }),
                        0
                      )
                    ])
                  ]
                ),
                _vm._v(" "),
                _c(
                  "div",
                  { staticClass: "row", staticStyle: { "margin-top": "20px" } },
                  [
                    _c("div", { staticClass: "r-title" }, [_c("label")]),
                    _vm._v(" "),
                    _c(
                      "div",
                      { staticClass: "r-content btn-c" },
                      [
                        _c(
                          "el-button",
                          {
                            staticClass: "green",
                            attrs: {
                              size: "medium",
                              disabled: _vm.isStorageExceeded
                            },
                            on: { click: _vm.submit }
                          },
                          [_vm._v(_vm._s(_vm.$t("modelManage.confirm")))]
                        ),
                        _vm._v(" "),
                        _c(
                          "el-button",
                          {
                            attrs: { size: "medium" },
                            on: { click: _vm.cancel }
                          },
                          [_vm._v(_vm._s(_vm.$t("modelManage.cancel")))]
                        )
                      ],
                      1
                    )
                  ]
                )
              ]
            )
          ])
        ]
      )
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=template&id=f63d172a&scoped=true":
/*!***************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Icons.vue?vue&type=template&id=f63d172a&scoped=true ***!
  \***************************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "icon-c" }, [
    _vm.type == "1"
      ? _c(
          "svg",
          {
            attrs: {
              version: "1.1",
              id: "Layer_1",
              xmlns: "http://www.w3.org/2000/svg",
              "xmlns:xlink": "http://www.w3.org/1999/xlink",
              x: "0px",
              y: "0px",
              width: "20",
              height: "20",
              viewBox: "0 0 50 50",
              "enable-background": "new 0 0 50 50",
              "xml:space": "preserve"
            }
          },
          [
            _c("image", {
              attrs: {
                id: "image0",
                width: "50",
                height: "50",
                x: "0",
                y: "0",
                href:
                  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACcAAAAmCAYAAABH/4KQAAAAAXNSR0IArs4c6QAABQtJREFUWEflmH1oW1UYxp/TZiZtvjba1dW0rNmSORxbs6rMgbYVRGUIhcFwf2w2c52iVfADhKHohhNBcA60ClvnWpygDoT+4eZw0LUb1Dmt6epmJaXp1g9d1w2bjzbJ2hx9z83tbtKb5iZZZeALgeTec8/7y3M+7nlehjs4WK5snPMKAPUAagEMxvtrZYydzrXvrOE45wRzBADBqQWBtjDG9mYLmTFcXCmCIjj4I9JnKAAEIoDNAqwpTsAhyB3ZKKkZTjF8e2Soi+PAjyNzdbHogfuKJUj6HuMYzmM4BWAvY0we+rSCaoLjnL/zb08CiqJrRB0qOZsMudEm3YlDNmsd6lm4hkNjn3Rfjmy7OHzTSh2tKVs0sWmdsf3dLUtc8rwa9gMnfdIwZhLJkLRw9h+fGPuyK3CvMl/Vcv3R5l0lL8l9C7i6j/483/bL5ANqCV/fZMXbm4sEFMHlEgT5hB048N11fHh8QrWruvsLf257tfRBuslIseaOQCP9aKovxvaHTeKhL84G0dg6Lr6/sbkEy2zS9VzjxtUg9n0zJrpJla+hxtxECrKqt4b+7h6MWqnhi49ZEnJ/esovADc4DXj6yXty5RLPf/39KM55wwIsVb6qirsmuveVL2Z69wCPTHP4D1XAbMhLAAiEY7DsGoRex/D+C/bbArf7Mx+05Iu0rGB3NpzWYX2q5m4YC/JzVi+jYdW6IE5emISzzIBKhzFryFB4Bp3nb+AHT0Dbgki3lWyrsSK/0IDWE1dFh46yArgcRlQ6TDAWJM7RVLKGIzF4+kP41RuEd2gKjAGcq7dO2ErkJi2d/onPOwOWc/3SDuuy61Gz1iK2kNDUDHr6Q/B4g+gfCUuQNlLRhEqnEaYUwx29yQVQjzeIvitT4rmVNgPWO00oNQE//RGExyfl2+DQ49lqs99dbREvAYrZNwTn3EdvgsM9qd8AMiSBeoelZA5bgRhqJeT0DBd/xOMN4dLgpGhnLzXA5TTB5TTCXJg4d2lz3lkpmg0yxma3hYzg5H8UCsfQ00+KJEKuXVmA/Lw8odJvAyHRfPkyvYBav8oESxKU3N9thbsFOSMAlUrK98pL9HCtMon5ucSsSzUdxfUFgZMzTkZiAvL3yyFwzlAh1DKiyLpoXqgFVU6Z+bWPB8TP3dvLsXSxNqj/HG7/yys0qaVstKDDSolk5f6XcO1kWo71ZXeozEU58hqPS7PhNGPsUbVNWMBp9QfJEysXuIdsQNxnkJXcoQYnTAwdxUm9TCMXuC2rgTLpnJsSjsyxLxjF9IkB6DL1C9nCERTBxcOutI4J1lC2gOSutBqav0aC6Oj1i6M3BR3p5QNDOvWTwBJUo2eT4Ug94eZJwd5r0KmZZjlp94XrONqh7qLoqFW1rigln2KezVkIc+acfEHN2RMguXtlkGIffHvLRW3daBS3v+oKzevaSC2yh7TxxmOOYinhFJBuALRIRKEm2VBrOW4rXZvsWeMTn7pMW0OZtxyRrKI81JfGgcYD2lxU0yt2UTeRSxJxKE3VJ621ErkGJ+olBFn8vE+nxeKFjyRYyj1a6yRzFkS61RVXkobaXfPeKDr75jfH1asN6HhTmHEqJFIZTHOFKWM4xXysbT0TaHMfvCa2zlRlhZbnlvrrHzHXZVObyxqOHiQV3QfH2lvPBFUrmztrzaOHG0rixa90Y6J+X9Ocm69r8r2eK9FneoeiZmpHpbPkUlZ2aMA/HhTNRaB9XSwAAAAASUVORK5CYII="
              }
            })
          ]
        )
      : _vm._e(),
    _vm._v(" "),
    _vm.type == "2"
      ? _c(
          "svg",
          {
            attrs: {
              version: "1.1",
              id: "Layer_1",
              xmlns: "http://www.w3.org/2000/svg",
              "xmlns:xlink": "http://www.w3.org/1999/xlink",
              x: "0px",
              y: "0px",
              width: "22",
              height: "22",
              viewBox: "0 0 54 54",
              "enable-background": "new 0 0 54 54",
              "xml:space": "preserve"
            }
          },
          [
            _c("image", {
              attrs: {
                id: "image0",
                width: "54",
                height: "54",
                x: "0",
                y: "0",
                href:
                  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACYAAAAmCAYAAACoPemuAAAAAXNSR0IArs4c6QAABB5JREFUWEftls9vE0cUx9/sjte/iEjDEqq0EEcgUKW6iUTFj9YSccOxlRCnioAgfwBU2fZATyFHTlSIe3zoH1ChHiohZKOkEYQe0laqVGqbBZxGQYE4IWST9e4MzMa7Xa/H3o0DoYf46Jn35rPf930zD8Eb/mnZYwkEeIwC5KLpidFW06NWA71xDEgAcYQCuvDfGlUpoEwrgG8ETMumRl4nutL4Iy3AoWh6IhdUiE2BadlUPwLIBj+MZgiYo9H0XdUvpiUw20cA0O93QP16sPJuCIwBAeDzzcsWFJUBCqPR9HiGFxEYzN9H6+npWqW0Ulj6i8xGpLaBuK+iCPjl9QVb9xEdA0CJZlowIE1dUsmKnqIGmtZ/j/aBACWpG+f9AevL2xCM3/58NE1dyJkLq446Dpi9vQVALljQsnmBbI46MBfgjoHYfLg71Mf7RGNxbXr14YIsD+f31oAFbX+rbH8/nycm4R7QEKxKg3eJE21fRBPCTvFD25fufLJSQBZY0PZ3+6ip32yPNdtULS/er4PbBizEAZu71Dfu/gJevkZla3B2rjLVNkRFk72ZTTszfHilLoUDNjOYpHYHxT+NHLAlttT0GDvgLZWTlUKa7S2dSV5ACEYAgNvV/mCeDtqRkg4085EPoAPG9s1+3ZcggnEekPXA1wBywFRZKfRYHrMU8/ykXm0aYco1d3WrWilKKqwJOPTRasoTXgNmr7kAnQffBaZSRDO7h4vWqNQSmFGSfjLn8CmWQIiRiaBgPEAGRhG9YgPZeyywfweTWa9JeYqZ82LGeBRmZnbKwQOT9sTv4K6Os7F0rtSs5ExBfHQ5sXs4XzcOOfeY16QesJyIzKGV39qYiV2DYK1iKIxn4gc7/gEsMviXAPRyJP3rjYANU7Ot5oJ1m1Tq1coI0zIDem9YteanmcHkWCOwSKL9Fm6PDABCggfktmmS7+MnJ+9vBJD7JDHAyNEy2EB2Qh5Y+BC+H/s8JgkC6m12MHmp34h9OXUxKJzvdOFOxANj61KP9Ef8RIQIGHG7eFUt54wFbb+sFPZtKZh9WPxI9F70E+kQBWhn/+lzy5P67IsuoFazPHlnYNb1EUYQ7xen9PKyCAY57FLo3YIxELHD+BP36ElP2bbBLEGEzsozsog1WEPWbNVYMfqjrBTPbZn58QeVSfH9ynG6KI7rD0PdYArdNaVEcNMkcHXPt4XJoFDOWxk0gHddVME+q+Z4Rp7ioqkJS6GELhIKVzuVwi9B87v3bfoe84BZuRFB93Z9lz/WCpAd81bAANDPspL/ahssgMeqIm0rti7E/1Yxa14TzZphkdeVAHBdVgrfbJn57YOenP24X6CIDY0JNxil6IdKhVzrulx8vBmoDV+w3sPYOI736adDncYDCuI1WXkws1kgO/4Vz5JIRYjrynQAAAAASUVORK5CYII="
              }
            })
          ]
        )
      : _vm._e(),
    _vm._v(" "),
    _vm.type == "3"
      ? _c(
          "svg",
          {
            attrs: {
              version: "1.1",
              id: "Layer_1",
              xmlns: "http://www.w3.org/2000/svg",
              "xmlns:xlink": "http://www.w3.org/1999/xlink",
              x: "0px",
              y: "0px",
              width: "22",
              height: "22",
              viewBox: "0 0 54 54",
              "enable-background": "new 0 0 54 54",
              "xml:space": "preserve"
            }
          },
          [
            _c("image", {
              attrs: {
                id: "image0",
                width: "54",
                height: "54",
                x: "0",
                y: "0",
                href:
                  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADQAAAA0CAYAAADFeBvrAAAAAXNSR0IArs4c6QAACxZJREFUaEPdWntUE1ca/2bygoTwDJBAAgmBKCIuFQs+2oqttT7a7rp2e7SL1V0r7erprrrWbqvuSe1qa3Wte2qpUm0pUrVra/es1WprFaqlgKVaFFAIJpBAXjzkFV5JZs9MmDgEMpOgdT29f82Z+71+93vc794ZBH5hA/mF4YE7Cmjtx9ai9m5sJr5IupZB0FntxHrpdsfT61Fj6O+k5W0Xqjgh5lYhOJ0owYeiToiO6IIM1UAL1yyMPapOGWBywG0Dem6/9T92B/br/PNd3nUVKr3qUW4qe9jWw/nGaA2htVUS2QGCQGyuZlv6aTrCMQNaecB88usrffN0LS4vUIdcxAZ5JBvkIg6UavqKr+2Iy/JmhGTtjxgOJjzCCqmpF0EiMQCL7ZLpsLPBaJTClSv3Q1trJOCgwkJ4vGoaT/kNKHuv6diFa/0LPYHMTgk0Ver71ZZcxT6msCDn5+75rvFUqViGg5kz53M3EE9+HNhXXy0kQM2eYtWfWTM1zpsOvwBlbW3Cimr63LJwT1g77Y/2fKA84ysIKh3pnZlZJ0Eq09GKMOjlUFw0n/CS8e3JXu32CVDiS41KaShLQwUTEcTa1LpXvnUsQEge9Nk6DC8Ai5fkefUOSYt76cjhHKJQOAuSxg5IvkabBoBeIkMsKzkAijbG+rQQTGDvOiDx6roUUwd6lTQsLZ7beXmrjL4ceUFh0qyPEifutPxfQ46aM5PiuPsrt8lWMq36aPMO424N2l+tBK6qDolZryJpJr1WdL6yTvaAP0Xh4ckt1rPrMqP8LgpUMFMUvOs/vC4d7y8Yh2FbAepoWgrY4C1WhANORPwhK27zH/GX/pTtkOB2mJMsvnr0RWmqX4CW7bNUf3S+KxlnSpdzuyv+IRP6A8beuH4xixV0GAab3WxY+w1AwhJuieHEgLlDkLrmlCqgWC+7yLSxhofdhL4BNth6gmDJNEHN4dXiCaPZNCK5x68zqK5Z+q/jxHhZZmxbPKQ6jdstSH99JPkas7UA1lQBWEcDICHxgMSkAyIQubl6neFF/IQ3Z6VsPvtlW3fo3NFaHy6391zD9hkPy9c0YGRx+u0UftmxNZKpnqBGAMr5wNyQd7ab2LiCONgj3R8mnvXHO5guByPo7f3gNFYAZnbXFLcYJGoioDHpAGwe8Q6R5/lUNaPX/yQItId2k6DGy3jjrr0hraXaN0xQ6iva8Vf0zhqcICOB21u+Rcb3BwxOSwJy1p8BPMy8xnpYAqDK2X4BwomTXtJtqjM6Xsefs6cHaQtXRVPiGIZ320veNVUe/r6HSDhBMCrpyVWY7jVARCpQQi8zkZ9SppZUk3YOd3V2PREuY8kdd84MhdzP5SEiejbrj5ZrB54ivDQj6Ebhn6KVIwAJcupTe2xQiU8oItmF2rfjl/rrnbsRcqRNpJfwxZeIUOn3m2RNRD6SBEv3mqsOXuh2lUKa8wsTSEybg+FSffYQBoAofCsKVN1T1U3HSjV9C/F3M5MDC4o3xiwbBojcSG8r3Jq21MKgIYlQjFe55grALKNUueiJgEpuVTngKZoRySuxTIs1Yn4oRXJmCZvzVkQR/LdyaGhyQVrgTyfWx6T5IxzTq/eD07piWEdAxB8AsQ/hwMh9KDYdkECRRznCLeGAkyXKZ0lf+4OvusmwozbMIwDNmsArPveq1OsJk6oM0699BtCQj2k7giGGEZ2Ct/ecGABn7+8R2fZDTMCGnc+G0sQFaLk2AOzOXvxRGsbaYHhHvoNJGGZ6qxX6NOHu6sbQERCx4NpyicHUQQAvvhmRbKQNw+xc80BhSTeHEMhGAyFf0ecCtPhaDLA5RJUAJ6jgkLKOEdAYOgJ3bt2hDmJVvuVY7pkuojAAAvFwUNlIAGIvrZtpx9AiFwhkPBQmEL0c3RhLR0Cs1x3sIJ7aY9z5aantrwQGFsy2f6T8xgUoWzPdDsh3+LNCxNql3S0niO51QMkv6xbUNDm+wO3k8+Bx2wHlCVfIZeskAA6i1388jd/6xXrJrXbYC6p7wUORz+vmW3scJwinIM4s+8Gk4hFVbvmDQsh/Poqx+70XAD22vfn901d6nyPWfFiVozR8vl6C3AuAcg5YavLOdblO0p6AsnPNg4Ul3WzqpC855FdHgAv0o4NgOietzrdq3j3TqaR2N+7QeibX8v6hki6X+9ABERQkt9IC0r9UCo6OTJIG6xmlI+BTUpEdDGDvdItkpEf53yJxu4mLf69jqLsZtVOYtMEgrWzu1+PMj0zgN3/zqoSxt8Ku54hAIDaB3cRyA/O8O0CDmqG3+1fIuLwWTLs8AFiRTeDsuLUhe9KzI/ug15yEJO030IJZqokCDDHjNDNU3K+++7vsMfx5WPJTWwlpJJdveFtGdA9MA7ux4glgBfzX83YHHJ1PIgkHj3vyY/UrFwGb++kIemfPEkSRf4RJHz6/YKex5MRl2zT8WS5Cw3S7FTdHAtrW9EZRdd/f8IkHxwWUnN8cO8MX4W7vNGw4AdjN+YCGHEfidjzJxIvpN54DhzULWKGnENlb85joqfPU8xD1ImdEeaYebycEC4XVuVHd/ii6G7RP/NN4/vgl2wO4rozEgNPl6ti5pN4RgNI3N35SoR18Gid4bFLg5dMbYu67G0b6o8Obd0aEHCmU6iWVmPVi7U75HnJuZWHFh8EC5AlwYoChcHzXwnSfzy/+GO2NluqdqYm8r0vV0jlU2lE7guAV+vBwobOVvP8KCsKiH0pp/f5CdWBCZ5dgmC4BvxfuH9+lKVo3zXVS/RmH9M/1SYY2IO7hvJ2svbY4YSs1i0MEnMM4KLHYBCaTmDA1OLgDwiMsxNmmvT0COoYqcEpCK1RtyWBsmcaKN3S5NvSm3dlO8gcLWAs698lPesqjNUCxVruvBzHkWCwxwOX1Q0ZGMcTLNcNkGPQKuFj+ENhsApiYaO67qp4eOFaj6fioW0pKLOe9qu1xq0ajpwWUvqX89YraiE1EgZj3GYhExD42YuBeOvXlIrAPcmB6qvWTkpenLr5ToCaoq7hRnNB+8uthZiKvrkwtdX+S8ctDU98sx0qvRkBq6g8wKa2c1saa6jT4sWI6TExshavqOxN6Mas0Mi6X00jm8n3x3M5LDB/caD0U/EIl1tktgMefPAIhIW20gPCQ+/yzZcDj9UP/gQm3nUuzthr2navpzyGVqiRsc+2OeFci0wxaxaxltZjDwYLFz+wDFsvBJAuOHHoecPrbuahM3tD4m+gQ1ufUD9QqMedY7c64RYwGePZyngz8lVWYrTcAFi4qAD6fvmEYHOTCv488ByzUAY4Cld8eenR783uDduwFKhDcHjaCzrUfVND+PUK1m1Zx8uZyrEYbARmZxZCkqqJdoAZdIlw4PwfiJe3QsGMKAyAMkb+oy0xPDFgrDGA9PdpvNfII1hHdv+RLfPGKz4Ae2nnh5LeXJfOChB0wf8FR4HC8/DuEIXDy5NPQ3hYBynjzqfqt00dvNIfOL3RGJkVzykxWdEFXgZT2POZNBmNoJLxyEbuhDwexuAkyp52FoKDhPyn19vGhrCQLmprkIIm6CcZd6d5legE0LZFXr7XYd5ly5bn+esSvso0Ti9aVqZyD3Ott7cGAAAbKpGsQHm4BBMWgrSUSbmjHET8ZCQU26LP3TR48cP8lb0Yl/KXhrCiMPdDZay9rbnV+2SmVl4Macd4uCJ9DjiQU5vwgkkQO1tY2RIWNpjxGbGm52cmZZMubYryTxo1FFmPIUYUKV19WJcf2fhbARSfiH4ECOJixUsN51vzO5DH9vDQWg5l4/gdrrlGAC46VpgAAAABJRU5ErkJggg=="
              }
            })
          ]
        )
      : _vm._e(),
    _vm._v(" "),
    _vm.type == "4"
      ? _c(
          "svg",
          {
            attrs: {
              version: "1.1",
              id: "Layer_1",
              xmlns: "http://www.w3.org/2000/svg",
              "xmlns:xlink": "http://www.w3.org/1999/xlink",
              x: "0px",
              y: "0px",
              width: "22",
              height: "22",
              viewBox: "0 0 54 54",
              "enable-background": "new 0 0 54 54",
              "xml:space": "preserve"
            }
          },
          [
            _c("image", {
              attrs: {
                id: "image0",
                width: "54",
                height: "54",
                x: "0",
                y: "0",
                href:
                  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAYAAACM/rhtAAAAAXNSR0IArs4c6QAAB1RJREFUWEfNmHtMU2cUwM+lFBCsPFpBHmZUnIggTpIycODUzYmPIOoWN7WgM0PBmKjb8DGZouDURHA6weEQBMIWpohMBcx8TNxEYWYOIiKL1DmqQPFVngW849z2K7fltpbCH56EwL3fvef7fed9oeA1F+o154NhA5y48ronr5u/A+DlDAsLSKjKCcwajsMPGRDBLHsso/pgdrKBBLa8yqeqro/uZgXLhgI6JEC/FRU72GCxi90YlsKrCpArVGA/kveipxeO3DgWsM1cSLMA/VZUzKBpOpOiKE/ceP40YefeGLENgWhQdEFRWQukFsiZW3Y2Fi3tXb1fmON2kwA1bvSszpVcwQ19l1dmUhS9crQDv2tf7DhriY+A00AImlbwCM6UKQCfvXR4SgRFUSV4wB7LHpkp7n8lINuNNE3LZgY4xvuPt1138rIi6PhX3uAusjbqPYT8NKkWJD6jIGaxK6xP/qex7mGHC+rq80BWda4kwZgCg4D6bmQrWRgqYjZjwxG3NjSrYGGoENhW1Xe5m8iKiVEUBDWW9QMANe7MBIAZRMHGj8fmzw50WH7u9yfuJK5wk4jpIggPFerEG76Da8Ri7Fgk78QscgOE3pJa3/lXXSsTuzRNZfXyuxP03a4DiFYDgMsEzG+c3ZWfdk16AwBmGkoAtmVJFpNDWPEpUHXTzCO4hmBsIZbNv9jcqXjebaNx+yoS6/isPiBTNmysLK5XZATcAYDVbIWd3eqrjF/kcPS0OkNR0J27oz21LseNc0sbIbekiVkX2lvC0vecdQBRV+9LABs+wOMn/cnk7MRPuXTorU1ENydgeKiwIilaLCEPobK2LoBbtUrYl10PjU/U8ePiyIfPl42FuUFOnHH+Z50SblYrteUGXbxztRh8xAJQ9fa/YmcNkHVOfWjce8+acYEmAap6ANq7QKss+7wcTpxTW44rUTgpAWB7uowpNShzgkQQJ2XKp47kFMsh6+wgAONXiSUvOnSVoCsulLdoITHGQvzt4dsN4znZtqfXw5myFu3alDcFEBfpCWOcBpYmcniTLbhpmViCMcIl+Rcfw/cF/2mX2NmJN9NOy7XtTv0QDV4edpC+dZIhI8OwAu7PkUFpuQJmBzqCrTVP6z4ERSE1Dq83LPWAuCP3wcXJCvJ2+w8PoMNIS+XaJR6C2YEiToUIWFKugKRoTyYWMWvj02VQUaNknkcwLCu4huInrYQxRgBv1/Unn1EX6xdprphByyEgum31AlfGQkQQVN6s0ukiuDZZWgkWFgDSuW4QOa+/FmJM78+WAQKSg5Wm+KdRFBVLdA7oJBOk5SEfhjgnV9YoJegqPPmcYBF8ECTUUYYKEjUWNOg3zQICEkFXb44UM1CkIrC7krvI2jDgZOnNlTRNZaKZYxe7SdhtimyAyhD61r3WQQE6CngwwpqnjU+ij3QYLENYjmIWuV5et8RjFqcFCWDgpFF3M7ZOmIgPkZGpouYF03uxXZG6NhgL4sFw+iGH1u8+mPnYIo3GIBcg231YuJWdAPuyZVB8XTEgGfRdzZ4HEbA0RZ3FWL6wM3X3qlsddpIhAaJCZUd/R9ENbhpGO1pBztcTdcavqN134dY9DH6KOUhitBgCvAVA2ib7MDwLgNxiOfxQZIYFsZuQIYEoRcDE4/ehpr5Nu4/AlgfSMBfmmkwz7LidN00En8zRnWbYkCYVan0Xoxvwhy2YfZsO1jK3MCPf9rWHorJmzkSOeFcIDU0qbX1M3uANWLq4xCzAZ239biVKSR3EjXBDcr0izBmUbS8BkwkTAHswydDNaTI4/4cCoubr1kGzLejlPkJeuNfXzRggmUoIIIHBeD18Us7MjOEhIkha4wkH89XXQwZkT9QY2FsixeDrpesSAjR9qiM8b+3RdgEsOe8HipiQYIeBq9AKQqY6wM+/NnEC6neTPWvFZ8JDRBEGOwnGob0d/9Cz1h6GLCxIBNL5rtoRqb/VqVXgQRa8I4JlYW5M+WCHAnYKHG5HjuBBa0evDqD+6MaaiIy3OtxgT54suL2V/q7wqiIAr8cIrSAsWAQ0Ddr2hPfRrdhb9ROJne3s+RFBl8xSZzppc0SP5nslo68uJVAU9dCgBdmBW3itaemd+x3f5F1oEutnHo5ayeu9oKVVXXiNSeqpf+HUJfX3CVtY3QQ/1BDsN/1nXvnhji9cq3r25dGCR7uanqps8Psj4+xjJlOPb/M2CZCUEMz02gcd0NDcxRRuiY/ggQYMP3M5xSRAfJOmaV7f750Niq7tYRurtPEXJxXD5PHctU0/AfR6dyLqoyiK9fk0kNFkQPIqTdMTjhXJT+C/PsjkjCUnkpVIRhIA1eRprHbPeGCoVwcNSJQW31BE/V3XnphT0shMrGRuxL8NJEC5BqzEFDDyjNmAREH5nefxN6qVm48VPbIzkACYHZgAqYMBGzZATXwKqupbD6T82PAZKwFw6YAGTj3TmyFDtiB7T5qmp/TlE/77BAWtdtsMJp1XhhVwqDBc7/8Ph50kZZ+3tvYAAAAASUVORK5CYII="
              }
            })
          ]
        )
      : _vm._e(),
    _vm._v(" "),
    _vm.type == "5"
      ? _c(
          "svg",
          {
            attrs: {
              version: "1.1",
              id: "Layer_1",
              xmlns: "http://www.w3.org/2000/svg",
              "xmlns:xlink": "http://www.w3.org/1999/xlink",
              x: "0px",
              y: "0px",
              width: "22",
              height: "22",
              viewBox: "0 0 54 54",
              "enable-background": "new 0 0 54 54",
              "xml:space": "preserve"
            }
          },
          [
            _c("image", {
              attrs: {
                id: "image0",
                width: "54",
                height: "54",
                x: "0",
                y: "0",
                href:
                  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACgAAAAoCAYAAACM/rhtAAAAAXNSR0IArs4c6QAABkVJREFUWEftmGtMU2cYx/9vT4FyaUVLANEgTikoIDhhYtSNhcxt2cVplrioTIwJTrdsxGXGDclY1Cwu2fTDNm8zqLBlH2Y2s4u6ZMqiG2gbMwYBwTjZpqYiF+dBbr284zntW0p7KKVkmR98k5NDy3v5nf/z/J/zvmW4zxu7z/nwAHCiEQpJQVNR3ZIwp5QraZyZBODkmll01zDnVc41HU4OKzSwOiR73eUji9omAhkUYEaR+SUNsIpz5ANIDH5BDg70M7A6znHCEWb/ZrzAAQEJDJxXMjCdN1R4GIMhSotInYRonQb6KAmGKAl3ex2wdtnQ1+/AgM0Judcx4lkYuAOMfdtQlbci2IdUBaQQRkBTxTlLoYkYY8605EjN4iwDSldND3Zu3OgYwBnLHZy9dAeX/+z1AHPwfhtzPtFalX9+rMlUATPXmrsBxAIcuemG9sqytPixJgrm/yfrulB+8BoGbBzg3D6ocT4+FqQf4LyiC585uWYDLdhQlXsAwMZgFqc+Dqerp7iLcZIGoIsaqbp8a6MLEqhprM57PND8foBZReavOccLhmip/pf987MDDSaQfhtgswODI9NNdZgCyoBDJ67j2EkrJInJ9UdzDeMEtFRyzotp0Kk9WZgWF+E3/t6AC8xXqWCVfnFbPbplGybrtfK5ffPHB5i51vzu0EIVYrHNK5Pw/FKjAkpAcl9waqnBnq7rwAdVw2Uxe3Z07+cVc6PHpaAAzDHF8PauQXazY1AZnzdHj6xZehQuNCJxir+qoy1CUPVXevBb613c6nLPNXcSzE3/IHeOnh8pS3dnp/oMfjkoAHeWpNAEKD/Yhhu3ByBAxTQJU8KRaHSBijv9be0cUC5qAkiMyU7VY90zScrHLXtblIeuLEsPWItHBXxz9XQUPz380rjQJOPc7zJO1Xb4LRwoRE/mxyE7NQbZJr1H+for8oQAzw7Zv6BwQSz2ls72rE3519kzjGLtGsCtTlfIhHKkZIIx3KVqgDSYEGBWkbqLfQGDdaxavwkBers4KS4c5GLKRVLEW8H/HfBhUwwutbpiSsm8IF0P0wy9EsLxuNiTAu6UIKNQemzZ04L8DAN2bZwZmknIxUlxEdh+8Jqfg2kR4d5Ed84N56IrL8kYZBAK5+5j1zzGIvd/sWOe0ifeoGxEQgdcvjROmcjcLONis4zaRlkpIb7lQy3cAoTKCUHSZ1F6CHxrUUrIgIqLNzybOGJr5W0Sbwd7AwsIAiEIaoWvWpT7T5/kKqFdU96gwJKKISkoNgszk3TY91aq510ciou9gZbkxKKn1wl6s1CKfFSaFhqgt4spxJtWTvW8h0Nx8bEfbuLo9zdBe0u4D5EER5AhKSgAJ+sldMuuPRSBPrfEiMio0Bz86fG/cPxMO8K0DLtfMylwEzbJjpIUp6W5R2NuvjvCxTR5jkmv5JF4B4/19vDNPWGqeAPsjLGwQDV11Hfx9uJk66rC+ERysOWyjItNMv5uH93B3uXDd8EAgK2MsbSQADevnFqzacW0AjFYmIRKBl3UrJ2Dnp0LqUmlQ601/CGj9MMWROkYDpVlegp9vAEHGGOvhASYNiPq8Fc75ypnE2rBuJiU+rGuU+lPqUCXr0lIadpyURkq39+y69cG2QbgZ7vW3qZ2ZlYLMal2dmjb39ZYnXdaHJqCARRFWTxUdKSEe30uo81L1SseFupH6yRHT7+Djiii1di19vW+kH6A6cW1KVq7tpKKNUGuXpZgf+fl5NljAdJWXtQ4CvfputtKWRGbVOFcAtxx+Crvlu1i7Qo6AzHGUhjj6xuqHjniHXLV96Abct1QR+VsEhujlR/NidUX5Bo9JcI3b4R6osaVvN+Eq9d7lYIs4GiMLgyWxSWW79xzVzRW573nVXuVz2MCig6Za80UbjpEecySMDmcg4GRSq4rHMvyjUruUUGmHHtqUZy7OANf7sriScaIPkmDJkMkXmeM1WYVXSzmnNHpsW1IOVJMESJoBX3VIUXDHFIB54xU9cCKfg9Nizr/xupky/aPrxTLfY5Y8f3CjEm7D79t2qbm0ow1lkrGXMdbd1M9xAf165YasNaupZxR6opNctRQcns9yAy71n50rF+y3BF6jFzcWJ1Xo/YgIQGqTfRfffcAcKLK3vcK/guCAu9HVzfHmQAAAABJRU5ErkJggg=="
              }
            })
          ]
        )
      : _vm._e()
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=template&id=77aca1ee&scoped=true":
/*!**************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Logs.vue?vue&type=template&id=77aca1ee&scoped=true ***!
  \**************************************************************************************************************************************************************************************************************************************************/
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
    { staticClass: "item-container" },
    [
      _c("div", { staticClass: "log-container" }, [
        _c("div", { staticClass: "op-tool-c" }, [
          _c(
            "div",
            [
              _vm.configs.multiNodes && _vm.multiNodesData.length > 1
                ? _c(
                    "el-select",
                    {
                      on: { change: _vm.changeNode },
                      model: {
                        value: _vm.common.nodeSel,
                        callback: function($$v) {
                          _vm.$set(_vm.common, "nodeSel", $$v)
                        },
                        expression: "common.nodeSel"
                      }
                    },
                    _vm._l(_vm.multiNodesData, function(item, index) {
                      return _c("el-option", {
                        key: item.id,
                        attrs: {
                          value: index,
                          label:
                            _vm.$t("cloudbrainObj.computeNode") +
                            " " +
                            (index + 1)
                        }
                      })
                    }),
                    1
                  )
                : _vm._e()
            ],
            1
          ),
          _vm._v(" "),
          _c("div", { staticClass: "op-btn-c" }, [
            _c(
              "a",
              {
                staticClass: "op-btn",
                class: _vm.downloading || !_vm.canLogDownload ? "disabled" : "",
                attrs: { href: _vm.common.downloadUrl }
              },
              [
                _c(
                  "svg",
                  {
                    staticClass: "fill",
                    attrs: {
                      xmlns: "http://www.w3.org/2000/svg",
                      viewBox: "0 0 32 32",
                      width: "20",
                      height: "20",
                      fill: "none"
                    }
                  },
                  [
                    _c("defs"),
                    _c("g", [
                      _c("path", {
                        attrs: {
                          d:
                            "M24 8h-4v2h2.8l4 8h-8.8v4h-4v-4h-8.8l4-8h2.8v-2h-4l-6 12v10h28v-10l-6-12zM28 28h-24v-8h8v4h8v-4h8v8z"
                        }
                      }),
                      _c("path", {
                        attrs: {
                          d:
                            "M11.2 12.8l4.8 4.6 4.8-4.6-1.6-1.6-2.2 2.4v-9.6h-2v9.6l-2.2-2.4z"
                        }
                      })
                    ])
                  ]
                ),
                _vm._v(" "),
                _c("span", [_vm._v(_vm._s(_vm.$t("modelManage.download")))])
              ]
            ),
            _vm._v(" "),
            _c(
              "a",
              {
                staticClass: "op-btn",
                attrs: { href: "javascript:;" },
                on: {
                  click: function($event) {
                    _vm.dialogShow = true
                  }
                }
              },
              [
                _c(
                  "svg",
                  {
                    attrs: {
                      xmlns: "http://www.w3.org/2000/svg",
                      viewBox: "0 0 20 20",
                      width: "20",
                      height: "20",
                      fill: "none"
                    }
                  },
                  [
                    _c("defs"),
                    _c("g", [
                      _c("rect", {
                        attrs: {
                          id: "全屏-imiijmsard5yz2g-V3pOBaFoDeaT9o",
                          width: "20.000000",
                          height: "20.000000",
                          x: "0.000000",
                          y: "0.000000"
                        }
                      }),
                      _vm._v(" "),
                      _c(
                        "g",
                        {
                          attrs: { id: "组合 7-imiijmsard5yz2g-V3pOBaFoDeaT9o" }
                        },
                        [
                          _c("path", {
                            attrs: {
                              id: "合并-imiijmsard5yz2g-V3pOBaFoDeaT9o",
                              d:
                                "M11.9333 7.21802L12.7819 8.06655L16.7998 4.04861L16.7998 8.00002L17.9998 8.00002L17.9998 2.84861L17.9999 2.84853L17.9998 2.84845L17.9998 2.00002L17.1514 2.00002L17.1514 2L17.1514 2.00002L11.9998 2.00002L11.9998 3.20002L15.9514 3.20002L11.9333 7.21802Z",
                              fill: "rgba(32,37,101,0.9)",
                              "fill-rule": "evenodd"
                            }
                          }),
                          _vm._v(" "),
                          _c("path", {
                            attrs: {
                              id: "合并-imiijmsard5yz2g-V3pOBaFoDeaT9o",
                              d:
                                "M0 5.21802L0.848528 6.06655L4.86647 2.04861L4.86647 6.00002L6.06647 6.00002L6.06647 0.84861L6.06655 0.848528L6.06647 0.848446L6.06647 1.74046e-05L5.21804 1.74046e-05L5.21802 0L5.218 1.74046e-05L0.0664654 1.74046e-05L0.0664654 1.20002L4.018 1.20002L0 5.21802Z",
                              fill: "rgba(32,37,101,0.9)",
                              "fill-rule": "evenodd",
                              transform: "matrix(-1,0,0,-1,8.06665,18)"
                            }
                          }),
                          _vm._v(" "),
                          _c("path", {
                            attrs: {
                              id: "合并-imiijmsard5yz2g-V3pOBaFoDeaT9o",
                              d:
                                "M6 1.2L6 0L0 0L0 1.2L4.8 1.2L4.8 6L6 6L6 1.2Z",
                              fill: "rgba(32,37,101,0.9)",
                              "fill-rule": "evenodd",
                              transform: "matrix(0,1,-1,0,18,12)"
                            }
                          }),
                          _vm._v(" "),
                          _c("path", {
                            attrs: {
                              id: "合并-imiijmsard5yz2g-V3pOBaFoDeaT9o",
                              d:
                                "M6 1.2L6 0L0 0L0 1.2L4.8 1.2L4.8 6L6 6L6 1.2Z",
                              fill: "rgba(32,37,101,0.9)",
                              "fill-rule": "evenodd",
                              transform: "matrix(-0,-1,1,0,2,8)"
                            }
                          })
                        ]
                      )
                    ])
                  ]
                ),
                _vm._v(" "),
                _c("span", [
                  _vm._v(_vm._s(_vm.$t("cloudbrainObj.viewFullScreen")))
                ])
              ]
            ),
            _vm._v(" "),
            _c(
              "a",
              {
                staticClass: "op-btn",
                attrs: { href: "javascript:;" },
                on: { click: _vm.refresh }
              },
              [
                _c(
                  "svg",
                  {
                    attrs: {
                      xmlns: "http://www.w3.org/2000/svg",
                      viewBox: "0 0 32 32",
                      width: "20",
                      height: "20"
                    }
                  },
                  [
                    _c("defs"),
                    _c("g", [
                      _c("path", {
                        attrs: {
                          fill: "#202565",
                          d:
                            "M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"
                        }
                      })
                    ])
                  ]
                ),
                _vm._v(" "),
                _c("span", [_vm._v(_vm._s(_vm.$t("cloudbrainObj.refresh")))])
              ]
            )
          ])
        ]),
        _vm._v(" "),
        _c(
          "div",
          {
            directives: [
              {
                name: "loading",
                rawName: "v-loading",
                value: _vm.common.loading,
                expression: "common.loading"
              }
            ],
            staticClass: "log-content-c"
          },
          [
            _c("div", {
              ref: "commonLogContentRef",
              staticClass: "log-content",
              domProps: { innerHTML: _vm._s(_vm.common.content) },
              on: { scroll: _vm.scrollHandler }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "right-btn-c" }, [
              _c(
                "div",
                {
                  staticClass: "icon-wrap",
                  staticStyle: { "margin-top": "8px" },
                  attrs: { title: _vm.$t("cloudbrainObj.scrollToTop") },
                  on: { click: _vm.goTop }
                },
                [
                  _c(
                    "svg",
                    {
                      attrs: {
                        xmlns: "http://www.w3.org/2000/svg",
                        viewBox: "0 0 48 48",
                        width: "16",
                        height: "16",
                        fill: "none"
                      }
                    },
                    [
                      _c("defs"),
                      _c("g", [
                        _c("path", {
                          attrs: {
                            d: "M24.0083 14.1005V42",
                            stroke: "#333",
                            "stroke-width": "4",
                            "stroke-linecap": "round",
                            "stroke-linejoin": "round"
                          }
                        }),
                        _c("path", {
                          attrs: {
                            d: "M12 26L24 14L36 26",
                            stroke: "#333",
                            "stroke-width": "4",
                            "stroke-linecap": "round",
                            "stroke-linejoin": "round"
                          }
                        }),
                        _c("path", {
                          attrs: {
                            d: "M12 6H36",
                            stroke: "#333",
                            "stroke-width": "4",
                            "stroke-linecap": "round",
                            "stroke-linejoin": "round"
                          }
                        })
                      ])
                    ]
                  )
                ]
              ),
              _vm._v(" "),
              _c(
                "div",
                {
                  staticClass: "icon-wrap",
                  staticStyle: { "margin-bottom": "8px" },
                  attrs: { title: _vm.$t("cloudbrainObj.scrollToBottom") },
                  on: { click: _vm.goBottom }
                },
                [
                  _c(
                    "svg",
                    {
                      attrs: {
                        xmlns: "http://www.w3.org/2000/svg",
                        viewBox: "0 0 48 48",
                        width: "16",
                        height: "16"
                      }
                    },
                    [
                      _c("path", {
                        attrs: {
                          d: "M24.0083 33.8995V6",
                          stroke: "#333",
                          "stroke-width": "4",
                          "stroke-linecap": "round",
                          "stroke-linejoin": "round"
                        }
                      }),
                      _c("path", {
                        attrs: {
                          d: "M36 22L24 34L12 22",
                          stroke: "#333",
                          "stroke-width": "4",
                          "stroke-linecap": "round",
                          "stroke-linejoin": "round"
                        }
                      }),
                      _c("path", {
                        attrs: {
                          d: "M36 42H12",
                          stroke: "#333",
                          "stroke-width": "4",
                          "stroke-linecap": "round",
                          "stroke-linejoin": "round"
                        }
                      })
                    ]
                  )
                ]
              )
            ])
          ]
        )
      ]),
      _vm._v(" "),
      _c(
        "el-dialog",
        {
          staticClass: "log-fullscreen-dlg",
          attrs: {
            visible: _vm.dialogShow,
            fullscreen: true,
            modal: true,
            "modal-append-to-body": true,
            "lock-scroll": true
          },
          on: {
            "update:visible": function($event) {
              _vm.dialogShow = $event
            },
            open: _vm.open,
            opened: _vm.opened,
            closed: _vm.closed
          }
        },
        [
          _c("div", { staticClass: "log-fullscreen-container" }, [
            _c("div", { staticClass: "log-fullscreen-header-c" }, [
              _c("div", { staticClass: "title" }, [
                _vm._v(_vm._s(_vm.$t("cloudbrainObj.logFile")))
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "op-tool-c" }, [
                _c(
                  "div",
                  { staticStyle: { "margin-right": "20px" } },
                  [
                    _vm.configs.multiNodes && _vm.multiNodesData.length > 1
                      ? _c(
                          "el-select",
                          {
                            on: { change: _vm.changeNode },
                            model: {
                              value: _vm.fullscreen.nodeSel,
                              callback: function($$v) {
                                _vm.$set(_vm.fullscreen, "nodeSel", $$v)
                              },
                              expression: "fullscreen.nodeSel"
                            }
                          },
                          _vm._l(_vm.multiNodesData, function(item, index) {
                            return _c("el-option", {
                              key: item.id,
                              attrs: {
                                value: index,
                                label:
                                  _vm.$t("cloudbrainObj.computeNode") +
                                  " " +
                                  (index + 1)
                              }
                            })
                          }),
                          1
                        )
                      : _vm._e()
                  ],
                  1
                ),
                _vm._v(" "),
                _c("div", { staticClass: "op-btn-c" }, [
                  _c(
                    "a",
                    {
                      staticClass: "op-btn",
                      class:
                        _vm.downloading || !_vm.canLogDownload
                          ? "disabled"
                          : "",
                      attrs: { href: _vm.fullscreen.downloadUrl }
                    },
                    [
                      _c("i", { staticClass: "ri-download-cloud-2-line" }),
                      _vm._v(" "),
                      _c("span", [
                        _vm._v(_vm._s(_vm.$t("modelManage.download")))
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "a",
                    {
                      staticClass: "op-btn",
                      attrs: { href: "javascript:;" },
                      on: {
                        click: function($event) {
                          _vm.dialogShow = false
                        }
                      }
                    },
                    [
                      _c("i", { staticClass: "ri-fullscreen-exit-fill" }),
                      _vm._v(" "),
                      _c("span", [
                        _vm._v(_vm._s(_vm.$t("cloudbrainObj.exitFullScreen")))
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "a",
                    {
                      staticClass: "op-btn",
                      attrs: { href: "javascript:;" },
                      on: { click: _vm.refresh }
                    },
                    [
                      _c(
                        "svg",
                        {
                          attrs: {
                            xmlns: "http://www.w3.org/2000/svg",
                            viewBox: "0 0 32 32",
                            width: "20",
                            height: "20"
                          }
                        },
                        [
                          _c("defs"),
                          _c("g", [
                            _c("path", {
                              attrs: {
                                fill: "#202565",
                                d:
                                  "M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"
                              }
                            })
                          ])
                        ]
                      ),
                      _vm._v(" "),
                      _c("span", [
                        _vm._v(_vm._s(_vm.$t("cloudbrainObj.refresh")))
                      ])
                    ]
                  )
                ])
              ])
            ]),
            _vm._v(" "),
            _c(
              "div",
              {
                directives: [
                  {
                    name: "loading",
                    rawName: "v-loading",
                    value: _vm.fullscreen.loading,
                    expression: "fullscreen.loading"
                  }
                ],
                staticClass: "log-fullscreen-content-c"
              },
              [
                _c("div", {
                  ref: "fullscreenLogContentRef",
                  staticClass: "log-content",
                  domProps: { innerHTML: _vm._s(_vm.fullscreen.content) },
                  on: { scroll: _vm.scrollHandler }
                }),
                _vm._v(" "),
                _c("div", { staticClass: "right-btn-c" }, [
                  _c(
                    "div",
                    {
                      staticClass: "icon-wrap",
                      staticStyle: { "margin-top": "8px" },
                      attrs: { title: _vm.$t("cloudbrainObj.scrollToTop") },
                      on: { click: _vm.goTop }
                    },
                    [
                      _c(
                        "svg",
                        {
                          attrs: {
                            xmlns: "http://www.w3.org/2000/svg",
                            viewBox: "0 0 48 48",
                            width: "16",
                            height: "16",
                            fill: "none"
                          }
                        },
                        [
                          _c("defs"),
                          _c("g", [
                            _c("path", {
                              attrs: {
                                d: "M24.0083 14.1005V42",
                                stroke: "#333",
                                "stroke-width": "4",
                                "stroke-linecap": "round",
                                "stroke-linejoin": "round"
                              }
                            }),
                            _c("path", {
                              attrs: {
                                d: "M12 26L24 14L36 26",
                                stroke: "#333",
                                "stroke-width": "4",
                                "stroke-linecap": "round",
                                "stroke-linejoin": "round"
                              }
                            }),
                            _c("path", {
                              attrs: {
                                d: "M12 6H36",
                                stroke: "#333",
                                "stroke-width": "4",
                                "stroke-linecap": "round",
                                "stroke-linejoin": "round"
                              }
                            })
                          ])
                        ]
                      )
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass: "icon-wrap",
                      staticStyle: { "margin-bottom": "8px" },
                      attrs: { title: _vm.$t("cloudbrainObj.scrollToBottom") },
                      on: { click: _vm.goBottom }
                    },
                    [
                      _c(
                        "svg",
                        {
                          attrs: {
                            xmlns: "http://www.w3.org/2000/svg",
                            viewBox: "0 0 48 48",
                            width: "16",
                            height: "16"
                          }
                        },
                        [
                          _c("path", {
                            attrs: {
                              d: "M24.0083 33.8995V6",
                              stroke: "#333",
                              "stroke-width": "4",
                              "stroke-linecap": "round",
                              "stroke-linejoin": "round"
                            }
                          }),
                          _c("path", {
                            attrs: {
                              d: "M36 22L24 34L12 22",
                              stroke: "#333",
                              "stroke-width": "4",
                              "stroke-linecap": "round",
                              "stroke-linejoin": "round"
                            }
                          }),
                          _c("path", {
                            attrs: {
                              d: "M36 42H12",
                              stroke: "#333",
                              "stroke-width": "4",
                              "stroke-linecap": "round",
                              "stroke-linejoin": "round"
                            }
                          })
                        ]
                      )
                    ]
                  )
                ])
              ]
            )
          ])
        ]
      )
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=template&id=e7b26d3c&scoped=true":
/*!**************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/Loss.vue?vue&type=template&id=e7b26d3c&scoped=true ***!
  \**************************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "item-container" }, [
    _c("div", { staticClass: "btn-wrap", on: { click: _vm.refresh } }, [
      _c(
        "svg",
        {
          attrs: {
            xmlns: "http://www.w3.org/2000/svg",
            viewBox: "0 0 32 32",
            width: "20",
            height: "20"
          }
        },
        [
          _c("defs"),
          _c("g", [
            _c("path", {
              attrs: {
                fill: "#202565",
                d:
                  "M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"
              }
            })
          ])
        ]
      ),
      _vm._v(" "),
      _c("span", [_vm._v(_vm._s(_vm.$t("cloudbrainObj.refresh")))])
    ]),
    _vm._v(" "),
    _c(
      "div",
      {
        directives: [
          {
            name: "loading",
            rawName: "v-loading",
            value: _vm.loading,
            expression: "loading"
          }
        ],
        staticClass: "chart-container"
      },
      [_c("div", { ref: "chartRef", staticClass: "chart" })]
    )
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=template&id=4ea8dafe&scoped=true":
/*!**************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/OperationProfile.vue?vue&type=template&id=4ea8dafe&scoped=true ***!
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
    {
      directives: [
        {
          name: "loading",
          rawName: "v-loading",
          value: _vm.loading,
          expression: "loading"
        }
      ],
      staticClass: "event-list"
    },
    [
      _c("div", { staticClass: "btn-wrap", on: { click: _vm.refresh } }, [
        _c(
          "svg",
          {
            attrs: {
              xmlns: "http://www.w3.org/2000/svg",
              viewBox: "0 0 32 32",
              width: "20",
              height: "20"
            }
          },
          [
            _c("defs"),
            _vm._v(" "),
            _c("g", [
              _c("path", {
                attrs: {
                  fill: "#202565",
                  d:
                    "M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"
                }
              })
            ])
          ]
        ),
        _vm._v(" "),
        _c("span", [_vm._v(_vm._s(_vm.$t("cloudbrainObj.refresh")))])
      ]),
      _vm._v(" "),
      _c(
        "div",
        { staticClass: "mesg-wrap" },
        [
          _vm._l(_vm.events, function(item, index) {
            return [
              _c("p", { key: index + "-1" }, [
                _c("b", [_vm._v("[" + _vm._s(item.reason) + "]")]),
                _vm._v(" "),
                _c("span", [_vm._v(_vm._s(item.timestampStr))])
              ]),
              _vm._v(" "),
              _c("p", { key: index + "-2" }, [_vm._v(_vm._s(item.message))])
            ]
          }),
          _vm._v(" "),
          _vm.events.length == 0 && !_vm.loading
            ? _c("p", [_vm._v(_vm._s(this.$t("noMessage")))])
            : _vm._e()
        ],
        2
      )
    ]
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=template&id=f5065876&scoped=true":
/*!************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResourceUseage.vue?vue&type=template&id=f5065876&scoped=true ***!
  \************************************************************************************************************************************************************************************************************************************************************/
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
  return !_vm.showErrorMsg
    ? _c("div", { staticClass: "item-container" }, [
        _c("div", { staticClass: "btn-wrap", on: { click: _vm.refresh } }, [
          _c(
            "svg",
            {
              attrs: {
                xmlns: "http://www.w3.org/2000/svg",
                viewBox: "0 0 32 32",
                width: "20",
                height: "20"
              }
            },
            [
              _c("defs"),
              _c("g", [
                _c("path", {
                  attrs: {
                    fill: "#202565",
                    d:
                      "M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"
                  }
                })
              ])
            ]
          ),
          _vm._v(" "),
          _c("span", [_vm._v(_vm._s(_vm.$t("cloudbrainObj.refresh")))])
        ]),
        _vm._v(" "),
        _c(
          "div",
          [
            _vm.configs.multiNodes && _vm.multiNodesData.length > 1
              ? _c(
                  "el-select",
                  {
                    on: { change: _vm.changeNode },
                    model: {
                      value: _vm.nodeSel,
                      callback: function($$v) {
                        _vm.nodeSel = $$v
                      },
                      expression: "nodeSel"
                    }
                  },
                  _vm._l(_vm.multiNodesData, function(item, index) {
                    return _c("el-option", {
                      key: item.id,
                      attrs: {
                        value: index,
                        label:
                          _vm.$t("cloudbrainObj.computeNode") +
                          " " +
                          (index + 1)
                      }
                    })
                  }),
                  1
                )
              : _vm._e()
          ],
          1
        ),
        _vm._v(" "),
        _c(
          "div",
          {
            directives: [
              {
                name: "loading",
                rawName: "v-loading",
                value: _vm.loading,
                expression: "loading"
              }
            ],
            staticClass: "chart-container"
          },
          [_c("div", { ref: "chartRef", staticClass: "chart" })]
        )
      ])
    : _c("div", { staticClass: "error-message" }, [
        _vm._v("\n  " + _vm._s(this.$t("noMessage")) + "\n")
      ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=template&id=a5b33778&scoped=true":
/*!************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/ResultDownload.vue?vue&type=template&id=a5b33778&scoped=true ***!
  \************************************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "item-container" }, [
    _c("div", { staticClass: "content" }, [
      _c(
        "div",
        {
          directives: [
            {
              name: "loading",
              rawName: "v-loading",
              value: _vm.loading,
              expression: "loading"
            }
          ],
          staticClass: "files-info"
        },
        [
          _vm.resultStatus == 0
            ? _c("div", { staticClass: "top" }, [
                _c("div", [
                  _c(
                    "div",
                    { staticClass: "title files-path-c" },
                    _vm._l(_vm.filePath, function(item, index) {
                      return _c(
                        "div",
                        { key: index, staticClass: "file-path" },
                        [
                          index == _vm.filePath.length - 1
                            ? _c("span", { staticClass: "path-name" }, [
                                _vm._v(_vm._s(item.label))
                              ])
                            : _vm._e(),
                          _vm._v(" "),
                          index != _vm.filePath.length - 1
                            ? _c(
                                "a",
                                {
                                  staticClass: "path-name canback",
                                  on: {
                                    click: function($event) {
                                      return _vm.goBackDir(item)
                                    }
                                  }
                                },
                                [_vm._v(_vm._s(item.label))]
                              )
                            : _vm._e(),
                          _vm._v(" "),
                          _c(
                            "span",
                            {
                              staticClass: "divider",
                              staticStyle: { color: "rgba(0,0,0,.4)" }
                            },
                            [_vm._v(" / ")]
                          )
                        ]
                      )
                    }),
                    0
                  )
                ]),
                _vm._v(" "),
                _c("div", { staticClass: "right-btn-c" }, [
                  _vm.filesList.length
                    ? _c(
                        "a",
                        {
                          staticClass: "op-btn",
                          class: !_vm.canDownload ? "disabled-download" : "",
                          attrs: { href: _vm.downloadAllUrl }
                        },
                        [
                          _c(
                            "svg",
                            {
                              staticClass: "fill",
                              attrs: {
                                xmlns: "http://www.w3.org/2000/svg",
                                viewBox: "0 0 32 32",
                                width: "20",
                                height: "20",
                                fill: "none"
                              }
                            },
                            [
                              _c("defs"),
                              _vm._v(" "),
                              _c("g", [
                                _c("path", {
                                  attrs: {
                                    d:
                                      "M24 8h-4v2h2.8l4 8h-8.8v4h-4v-4h-8.8l4-8h2.8v-2h-4l-6 12v10h28v-10l-6-12zM28 28h-24v-8h8v4h8v-4h8v8z"
                                  }
                                }),
                                _vm._v(" "),
                                _c("path", {
                                  attrs: {
                                    d:
                                      "M11.2 12.8l4.8 4.6 4.8-4.6-1.6-1.6-2.2 2.4v-9.6h-2v9.6l-2.2-2.4z"
                                  }
                                })
                              ])
                            ]
                          ),
                          _vm._v(" "),
                          _c("span", [
                            _vm._v(
                              _vm._s(_vm.$t("cloudbrainObj.allResultDownload"))
                            )
                          ])
                        ]
                      )
                    : _vm._e(),
                  _vm._v(" "),
                  _c(
                    "a",
                    {
                      staticClass: "op-btn",
                      attrs: { href: "javascript:;" },
                      on: { click: _vm.refresh }
                    },
                    [
                      _c(
                        "svg",
                        {
                          attrs: {
                            xmlns: "http://www.w3.org/2000/svg",
                            viewBox: "0 0 32 32",
                            width: "20",
                            height: "20"
                          }
                        },
                        [
                          _c("defs"),
                          _vm._v(" "),
                          _c("g", [
                            _c("path", {
                              attrs: {
                                fill: "#202565",
                                d:
                                  "M25.6 15l-1.8 1.8c0.4 3-0.4 6.2-2.8 8.4-3.8 3.8-10.2 3.8-14.2 0-3.8-3.8-3.8-10.2 0-14.2 3-3 7.2-3.6 10.8-2.2l-5 5 1.4 1.4 7.2-7.2-7.2-7-1.4 1.4 3.8 3.8c-3.8-0.8-8 0.2-11 3.2-4.6 4.6-4.6 12.2 0 17 4.6 4.6 12.2 4.6 17 0 3.2-3 4.2-7.4 3.2-11.4z"
                              }
                            })
                          ])
                        ]
                      ),
                      _vm._v(" "),
                      _c("span", [
                        _vm._v(_vm._s(_vm.$t("cloudbrainObj.refresh")))
                      ])
                    ]
                  )
                ])
              ])
            : _vm._e(),
          _vm._v(" "),
          _vm.resultStatus == 0
            ? _c(
                "div",
                { staticClass: "table-container" },
                [
                  _c(
                    "el-table",
                    {
                      ref: "tableRef",
                      attrs: {
                        data: _vm.filesList,
                        "row-key": "sn",
                        height: "100%",
                        tyle: "min-width:100%"
                      }
                    },
                    [
                      _c("el-table-column", {
                        attrs: {
                          "column-key": "FileName",
                          prop: "FileName",
                          sortable: "",
                          "sort-method": function(a, b) {
                            return a.FileName.toLocaleLowerCase().localeCompare(
                              b.FileName.toLocaleLowerCase()
                            )
                          },
                          label: _vm.$t("modelManage.fileName"),
                          align: "left",
                          "header-align": "left",
                          "min-width": "160"
                        },
                        scopedSlots: _vm._u(
                          [
                            {
                              key: "default",
                              fn: function(scope) {
                                return [
                                  _c("div", { staticClass: "tbl-file-name" }, [
                                    scope.row.IsDir
                                      ? _c(
                                          "a",
                                          {
                                            attrs: { href: "javascript:;" },
                                            on: {
                                              click: function($event) {
                                                return _vm.goNextDir(scope.row)
                                              }
                                            }
                                          },
                                          [
                                            _c(
                                              "div",
                                              {
                                                staticClass: "fitted",
                                                attrs: {
                                                  title: scope.row.FileName
                                                }
                                              },
                                              [
                                                _c("i", {
                                                  staticClass: "icon folder",
                                                  attrs: {
                                                    width: "16",
                                                    height: "16",
                                                    "aria-hidden": "true"
                                                  }
                                                }),
                                                _vm._v(" "),
                                                _c("span", [
                                                  _vm._v(
                                                    _vm._s(scope.row.FileName)
                                                  )
                                                ])
                                              ]
                                            )
                                          ]
                                        )
                                      : _c(
                                          "a",
                                          {
                                            class: !_vm.canDownload
                                              ? "disabled-download"
                                              : "",
                                            attrs: {
                                              href: scope.row.downloadUrl
                                            }
                                          },
                                          [
                                            _c(
                                              "div",
                                              {
                                                staticClass: "fitted",
                                                attrs: {
                                                  title: scope.row.FileName
                                                }
                                              },
                                              [
                                                _c("i", {
                                                  staticClass: "icon file",
                                                  attrs: {
                                                    width: "16",
                                                    height: "16",
                                                    "aria-hidden": "true"
                                                  }
                                                }),
                                                _vm._v(" "),
                                                _c("span", [
                                                  _vm._v(
                                                    _vm._s(scope.row.FileName)
                                                  )
                                                ])
                                              ]
                                            )
                                          ]
                                        )
                                  ])
                                ]
                              }
                            }
                          ],
                          null,
                          false,
                          3279437061
                        )
                      }),
                      _vm._v(" "),
                      _c("el-table-column", {
                        attrs: {
                          "column-key": "SizeShow",
                          prop: "SizeShow",
                          sortable: "",
                          "sort-method": function(a, b) {
                            return a.Size - b.Size
                          },
                          label: _vm.$t("modelManage.fileSize"),
                          align: "left",
                          "header-align": "left",
                          width: "200"
                        }
                      }),
                      _vm._v(" "),
                      _c("el-table-column", {
                        attrs: {
                          "column-key": "ModTime",
                          prop: "ModTime",
                          sortable: "",
                          "sort-method": function(a, b) {
                            return a.ModTimeNum - b.ModTimeNum
                          },
                          label: _vm.$t("modelManage.updateTime"),
                          align: "center",
                          "header-align": "center",
                          width: "200"
                        }
                      })
                    ],
                    1
                  ),
                  _vm._v(" "),
                  _c("div", { staticClass: "max-count-tips" }, [
                    _c("i", { staticClass: "el-icon-warning-outline" }),
                    _vm._v(
                      _vm._s(
                        this.$t("cloudbrainObj.downloadDisplayMaxCountTips", {
                          count: 100,
                          size: _vm.size
                        })
                      )
                    )
                  ])
                ],
                1
              )
            : _vm._e(),
          _vm._v(" "),
          _vm.resultStatus != 0
            ? _c("div", [
                _vm.isTaskTerminal === false
                  ? _c("div", { staticClass: "status-tips" }, [
                      _c("div", [
                        _c("i", { staticClass: "ri-time-line" }),
                        _vm._v(" "),
                        _c("span", [
                          _vm._v(
                            _vm._s(this.$t("cloudbrainObj.task_not_finished"))
                          )
                        ])
                      ])
                    ])
                  : _c("div", { staticClass: "status-tips" }, [
                      _vm.resultStatus == 1
                        ? _c("div", [
                            _c(
                              "svg",
                              {
                                staticClass: "rotating",
                                staticStyle: { "margin-right": "5px" },
                                attrs: {
                                  xmlns: "http://www.w3.org/2000/svg",
                                  viewBox: "0 0 24 24",
                                  width: "14",
                                  height: "14",
                                  fill: "#101010"
                                }
                              },
                              [
                                _c("path", {
                                  attrs: {
                                    d:
                                      "M6 4H4V2H20V4H18V6C18 7.61543 17.1838 8.91468 16.1561 9.97667C15.4532 10.703 14.598 11.372 13.7309 12C14.598 12.628 15.4532 13.297 16.1561 14.0233C17.1838 15.0853 18 16.3846 18 18V20H20V22H4V20H6V18C6 16.3846 6.81616 15.0853 7.8439 14.0233C8.54682 13.297 9.40202 12.628 10.2691 12C9.40202 11.372 8.54682 10.703 7.8439 9.97667C6.81616 8.91468 6 7.61543 6 6V4ZM8 4V6C8 6.88457 8.43384 7.71032 9.2811 8.58583C10.008 9.33699 10.9548 10.0398 12 10.7781C13.0452 10.0398 13.992 9.33699 14.7189 8.58583C15.5662 7.71032 16 6.88457 16 6V4H8ZM12 13.2219C10.9548 13.9602 10.008 14.663 9.2811 15.4142C8.43384 16.2897 8 17.1154 8 18V20H16V18C16 17.1154 15.5662 16.2897 14.7189 15.4142C13.992 14.663 13.0452 13.9602 12 13.2219Z"
                                  }
                                })
                              ]
                            ),
                            _vm._v(" "),
                            _c("span", [
                              _vm._v(
                                _vm._s(this.$t("cloudbrainObj.file_sync_ing"))
                              )
                            ])
                          ])
                        : _vm._e(),
                      _vm._v(" "),
                      _vm.resultStatus == 2
                        ? _c("div", [
                            _c("i", { staticClass: "ri-alert-line" }),
                            _vm._v(" "),
                            _c("span", [
                              _vm._v(
                                _vm._s(this.$t("cloudbrainObj.file_sync_fail"))
                              )
                            ]),
                            _vm._v(" "),
                            _vm.canReschedule
                              ? _c(
                                  "a",
                                  {
                                    staticClass: "retry",
                                    attrs: { href: "javascript:void(0)" },
                                    on: { click: _vm.retry }
                                  },
                                  [
                                    _vm._v(
                                      _vm._s(
                                        this.$t(
                                          "cloudbrainObj.retrieve_results"
                                        )
                                      )
                                    )
                                  ]
                                )
                              : _vm._e()
                          ])
                        : _vm._e(),
                      _vm._v(" "),
                      _vm.resultStatus == 3
                        ? _c("div", [
                            _c("i", { staticClass: "ri-time-line" }),
                            _vm._v(" "),
                            _c("span", [
                              _vm._v(
                                _vm._s(this.$t("cloudbrainObj.file_sync_wait"))
                              )
                            ])
                          ])
                        : _vm._e(),
                      _vm._v(" "),
                      _vm.resultStatus == 4
                        ? _c("div", [
                            _c("i", { staticClass: "ri-emotion-unhappy-line" }),
                            _vm._v(" "),
                            _c("span", [
                              _vm._v(
                                _vm._s(
                                  this.$t("cloudbrainObj.no_file_to_download")
                                )
                              )
                            ])
                          ])
                        : _vm._e()
                    ])
              ])
            : _vm._e()
        ]
      )
    ])
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=template&id=570c19a0&scoped=true":
/*!********************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/details/TaskDetailCollapseTabs.vue?vue&type=template&id=570c19a0&scoped=true ***!
  \********************************************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "task-detail-container" }, [
    _c(
      "div",
      { staticClass: "body-content" },
      [
        _vm._l(_vm.mainData, function(item, index) {
          return _c(
            "div",
            { staticClass: "content" },
            [
              _c(
                "div",
                { staticClass: "operation-area" },
                [
                  item.task.isSubscriber &&
                  _vm.operationList.indexOf("consoleHome") >= 0
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_modify && item.task.canDebug
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "1" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opDebug(item)
                                  }
                                }
                              },
                              [
                                _vm._v(
                                  _vm._s(_vm.$t("cloudbrainObj.consoleHome"))
                                )
                              ]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("onlineInfer") >= 0 &&
                  item.task.canDebug
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_modify && item.task.canDebug
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "1" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opVisual(item)
                                  }
                                }
                              },
                              [_vm._v(_vm._s(_vm.$t("onlineinfer")))]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("onlinexperience") >= 0
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_experience && item.task.canDebug
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "1" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opOnline(item)
                                  }
                                }
                              },
                              [
                                _vm._v(
                                  _vm._s(_vm.$t("modelManage.onlineInference"))
                                )
                              ]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("onlineLoraTrain") >= 0
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_experience && item.task.canDebug
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "1" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opLoraTrain(item)
                                  }
                                }
                              },
                              [
                                _vm._v(
                                  _vm._s(_vm.$t("modelManage.onlineLoraTrain"))
                                )
                              ]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("onlineWorkflow") >= 0
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_experience && item.task.canDebug
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "1" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opComfyUI(item)
                                  }
                                }
                              },
                              [
                                _vm._v(
                                  _vm._s(_vm.$t("modelManage.onlineWorkflow"))
                                )
                              ]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("debug") >= 0 && item.task.canDebug
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_modify && item.task.canDebug
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "1" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opDebug(item)
                                  }
                                }
                              },
                              [
                                _vm._v(
                                  _vm._s(_vm.$t("cloudbrainObj.startDebug"))
                                )
                              ]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("redebug") >= 0 &&
                  !item.task.canDebug &&
                  !item.task.is_file_notebook
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_modify && item.task.canReDebug
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "1" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opReDebug(item)
                                  }
                                }
                              },
                              [_vm._v(_vm._s(_vm.$t("cloudbrainObj.reDebug")))]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("reinfer") >= 0 &&
                  !item.task.canDebug &&
                  !item.task.is_file_notebook
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_modify && item.task.canReDebug
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "1" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opReDebug(item)
                                  }
                                }
                              },
                              [_vm._v(_vm._s(_vm.$t("cloudbrainObj.reInfer")))]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("tensorBoard") >= 0 &&
                  item.task.visualize_required
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_modify && item.task.canVisualize
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "2" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opTensorBoard(item)
                                  }
                                }
                              },
                              [
                                _vm._v(
                                  _vm._s(
                                    _vm.$t(
                                      "cloudbrainObj.tensorBoardVisualization"
                                    )
                                  )
                                )
                              ]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  item.can_modify && item.task.canAim
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          { staticClass: "operate-btn-c" },
                          [
                            _c("Icons", { attrs: { type: "1" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opAim(item)
                                  }
                                }
                              },
                              [
                                _vm._v(
                                  _vm._s(
                                    _vm.$t("cloudbrainObj.aimVisualization")
                                  )
                                )
                              ]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("stop") >= 0
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_delete && item.task.canStop
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "1" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.opStop(item)
                                  }
                                }
                              },
                              [_vm._v(_vm._s(_vm.$t("cloudbrainObj.stopTask")))]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  ["ComfyuiExperience"].includes(item.task.job_type)
                    ? [
                        item.can_experience && item.task.canSaveImage
                          ? _c("div", { staticClass: "operate-item" }, [
                              item.can_experience && item.task.canSaveImage
                                ? _c(
                                    "a",
                                    {
                                      staticClass: "operate-btn-c",
                                      attrs: {
                                        target: "_blank",
                                        href: item.task.saveImageUrl
                                      }
                                    },
                                    [
                                      _c("Icons", { attrs: { type: "1" } }),
                                      _vm._v(" "),
                                      _c(
                                        "div",
                                        { staticClass: "operate-btn" },
                                        [
                                          _vm._v(
                                            " " +
                                              _vm._s(
                                                _vm.$t(
                                                  "cloudbrainObj.commitImage"
                                                )
                                              )
                                          )
                                        ]
                                      )
                                    ],
                                    1
                                  )
                                : _vm._e()
                            ])
                          : _vm._e()
                      ]
                    : [
                        item.can_modify && item.task.canSaveImage
                          ? _c("div", { staticClass: "operate-item" }, [
                              item.can_modify && item.task.canSaveImage
                                ? _c(
                                    "a",
                                    {
                                      staticClass: "operate-btn-c",
                                      attrs: {
                                        target: "_blank",
                                        href: item.task.saveImageUrl
                                      }
                                    },
                                    [
                                      _c("Icons", { attrs: { type: "1" } }),
                                      _vm._v(" "),
                                      _c(
                                        "div",
                                        { staticClass: "operate-btn" },
                                        [
                                          _vm._v(
                                            " " +
                                              _vm._s(
                                                _vm.$t(
                                                  "cloudbrainObj.commitImage"
                                                )
                                              )
                                          )
                                        ]
                                      )
                                    ],
                                    1
                                  )
                                : _vm._e()
                            ])
                          : _vm._e()
                      ],
                  _vm._v(" "),
                  item.can_create_template && item.task.canSaveTmpl
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class:
                              item.can_create_template && item.task.canSaveTmpl
                                ? ""
                                : "btn-disabled"
                          },
                          [
                            _c("Icons", { attrs: { type: "3" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.saveTmpl(item)
                                  }
                                }
                              },
                              [
                                _vm._v(
                                  _vm._s(_vm.$t("cloudbrainObj.saveTaskTmpl"))
                                )
                              ]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("saveModel") >= 0
                    ? _c("div", { staticClass: "operate-item export-output" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class: !item.task.canExportOutput
                              ? "btn-disabled"
                              : ""
                          },
                          [
                            _c("Icons", { attrs: { type: "4" } }),
                            _vm._v(" "),
                            _c("ExportModel", {
                              attrs: {
                                data: item.task,
                                configs: _vm.pageCfg,
                                disabled: !item.task.canExportOutput
                              }
                            })
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("exportDataset") >= 0
                    ? _c("div", { staticClass: "operate-item export-output" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class: !item.task.canExportOutput
                              ? "btn-disabled"
                              : ""
                          },
                          [
                            _c("Icons", { attrs: { type: "5" } }),
                            _vm._v(" "),
                            _c("ExportDataset", {
                              attrs: {
                                data: item.task,
                                configs: _vm.pageCfg,
                                disabled: !item.task.canExportOutput
                              }
                            })
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("deployModel") >= 0
                    ? _c("div", { staticClass: "operate-item" }, [
                        _c(
                          "div",
                          {
                            staticClass: "operate-btn-c",
                            class: !item.can_fintune_experience
                              ? "btn-disabled"
                              : ""
                          },
                          [
                            _c("Icons", { attrs: { type: "5" } }),
                            _vm._v(" "),
                            _c(
                              "div",
                              {
                                staticClass: "operate-btn",
                                on: {
                                  click: function($event) {
                                    return _vm.deployModel(item)
                                  }
                                }
                              },
                              [
                                _vm._v(
                                  _vm._s(
                                    _vm.$t("cloudbrainObj.deploymentExperience")
                                  )
                                )
                              ]
                            )
                          ],
                          1
                        )
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.operationList.indexOf("exportDataset") >= 0 ||
                  _vm.operationList.indexOf("saveModel") >= 0
                    ? _c("div", { staticClass: "mobile-export-output" }, [
                        _vm._v(
                          "\n          " +
                            _vm._s(_vm.$t("cloudbrainObj.exportOutputTis")) +
                            "\n        "
                        )
                      ])
                    : _vm._e()
                ],
                2
              ),
              _vm._v(" "),
              _c(
                "el-tabs",
                {
                  on: {
                    "tab-click": function($event) {
                      return _vm.tabChange(item)
                    }
                  },
                  model: {
                    value: item.activeName,
                    callback: function($$v) {
                      _vm.$set(item, "activeName", $$v)
                    },
                    expression: "item.activeName"
                  }
                },
                [
                  _vm.tabNameList.indexOf("configInfo") >= 0
                    ? _c(
                        "el-tab-pane",
                        {
                          attrs: {
                            label: _vm.$t("cloudbrainObj.configurationInfo"),
                            name: "configInfo-" + item.task.id
                          }
                        },
                        [
                          _c("ConfigInfo", {
                            ref: "configInfo-" + item.task.id + "-Ref",
                            refInFor: true,
                            attrs: {
                              data: item,
                              configs: _vm.tabConfigs["configInfo"]
                            }
                          })
                        ],
                        1
                      )
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tabNameList.indexOf("operationProfile") >= 0
                    ? _c(
                        "el-tab-pane",
                        {
                          attrs: {
                            label: _vm.$t("cloudbrainObj.taskRuntimeInfo"),
                            name: "operationProfile-" + item.task.id
                          }
                        },
                        [
                          _c("OperationProfile", {
                            ref: "operationProfile-" + item.task.id + "-Ref",
                            refInFor: true,
                            attrs: {
                              data: item,
                              configs: _vm.tabConfigs["operationProfile"]
                            }
                          })
                        ],
                        1
                      )
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tabNameList.indexOf("logs") >= 0
                    ? _c(
                        "el-tab-pane",
                        {
                          attrs: {
                            label: _vm.$t("cloudbrainObj.log"),
                            name: "logs-" + item.task.id
                          }
                        },
                        [
                          _c("Logs", {
                            ref: "logs-" + item.task.id + "-Ref",
                            refInFor: true,
                            attrs: {
                              data: item,
                              configs: _vm.tabConfigs["logs"]
                            }
                          })
                        ],
                        1
                      )
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tabNameList.indexOf("resourceUseage") >= 0
                    ? _c(
                        "el-tab-pane",
                        {
                          attrs: {
                            label: _vm.$t("cloudbrainObj.resourceOccupancy"),
                            name: "resourceUseage-" + item.task.id
                          }
                        },
                        [
                          _c("ResourceUseage", {
                            ref: "resourceUseage-" + item.task.id + "-Ref",
                            refInFor: true,
                            attrs: {
                              data: item,
                              configs: _vm.tabConfigs["resourceUseage"]
                            }
                          })
                        ],
                        1
                      )
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tabNameList.indexOf("resultDownload") >= 0
                    ? _c(
                        "el-tab-pane",
                        {
                          attrs: {
                            label: _vm.$t("cloudbrainObj.modelDownload"),
                            name: "resultDownload-" + item.task.id
                          }
                        },
                        [
                          _c("ResultDownload", {
                            ref: "resultDownload-" + item.task.id + "-Ref",
                            refInFor: true,
                            attrs: {
                              data: item,
                              configs: _vm.tabConfigs["resultDownload"]
                            }
                          })
                        ],
                        1
                      )
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tabNameList.indexOf("loss") >= 0
                    ? _c(
                        "el-tab-pane",
                        {
                          attrs: {
                            label: _vm.$t("cloudbrainObj.lossPlot"),
                            name: "loss-" + item.task.id
                          }
                        },
                        [
                          _c("Loss", {
                            ref: "loss-" + item.task.id + "-Ref",
                            refInFor: true,
                            attrs: { data: item }
                          })
                        ],
                        1
                      )
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tabNameList.indexOf("evalOverview") >= 0
                    ? _c(
                        "el-tab-pane",
                        {
                          attrs: {
                            label: _vm.$t("cloudbrainObj.evalOverview"),
                            name: "evalOverview-" + item.task.id
                          }
                        },
                        [
                          _c("EvalOverview", {
                            ref: "evalOverview-" + item.task.id + "-Ref",
                            refInFor: true,
                            attrs: { data: item }
                          })
                        ],
                        1
                      )
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tabNameList.indexOf("evalDetail") >= 0
                    ? _c(
                        "el-tab-pane",
                        {
                          attrs: {
                            label: _vm.$t("cloudbrainObj.evalDetail"),
                            name: "evalDetail-" + item.task.id
                          }
                        },
                        [
                          _c("EvalDetail", {
                            ref: "evalDetail-" + item.task.id + "-Ref",
                            refInFor: true,
                            attrs: { data: item }
                          })
                        ],
                        1
                      )
                    : _vm._e()
                ],
                1
              )
            ],
            1
          )
        }),
        _vm._v(" "),
        _c(
          "el-dialog",
          {
            staticClass: "task-already-dlg",
            attrs: {
              visible: _vm.taskAlreadyDialogShow,
              "lock-scroll": false,
              "show-close": false
            },
            on: {
              "update:visible": function($event) {
                _vm.taskAlreadyDialogShow = $event
              }
            }
          },
          [
            _c("div", { staticClass: "err-msg-box-already" }, [
              _c("div", { staticClass: "msg-content" }, [
                _c("i", { staticClass: "ri-information-line" }),
                _vm._v(" "),
                _c("div", { staticClass: "msg-content-tip" }, [
                  _c("div", {
                    staticClass: "line-1",
                    domProps: {
                      innerHTML: _vm._s(_vm.taskAlreadyDialogShowMsg)
                    }
                  }),
                  _vm._v(" "),
                  _c("div", {
                    staticClass: "line-2",
                    domProps: {
                      innerHTML: _vm._s(_vm.$t("cloudbrainObj.sameTaskTips2"))
                    }
                  })
                ])
              ])
            ])
          ]
        )
      ],
      2
    )
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ })

}]);