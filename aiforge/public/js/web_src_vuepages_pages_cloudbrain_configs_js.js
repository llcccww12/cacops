"use strict";
(self["webpackChunk"] = self["webpackChunk"] || []).push([["web_src_vuepages_pages_cloudbrain_configs_js"],{

/***/ "./web_src/vuepages/pages/cloudbrain/configs.js":
/*!******************************************************!*\
  !*** ./web_src/vuepages/pages/cloudbrain/configs.js ***!
  \******************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   COMPUTER_RESOURCES_TITLE: function() { return /* binding */ COMPUTER_RESOURCES_TITLE; },
/* harmony export */   CreatePageConfigs: function() { return /* binding */ CreatePageConfigs; },
/* harmony export */   DetailPageConfigs: function() { return /* binding */ DetailPageConfigs; },
/* harmony export */   FieldTemplates: function() { return /* binding */ FieldTemplates; },
/* harmony export */   configCreateManager: function() { return /* binding */ configCreateManager; },
/* harmony export */   configDetailManager: function() { return /* binding */ configDetailManager; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.symbol.description */ "./node_modules/core-js/modules/es.symbol.description.js");
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_symbol_iterator__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.symbol.iterator */ "./node_modules/core-js/modules/es.symbol.iterator.js");
/* harmony import */ var core_js_modules_es_symbol_iterator__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol_iterator__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.find */ "./node_modules/core-js/modules/es.array.find.js");
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_from__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.from */ "./node_modules/core-js/modules/es.array.from.js");
/* harmony import */ var core_js_modules_es_array_from__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_from__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.array.is-array */ "./node_modules/core-js/modules/es.array.is-array.js");
/* harmony import */ var core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.array.iterator */ "./node_modules/core-js/modules/es.array.iterator.js");
/* harmony import */ var core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_map__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.map */ "./node_modules/core-js/modules/es.map.js");
/* harmony import */ var core_js_modules_es_map__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_map__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_object_entries__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.object.entries */ "./node_modules/core-js/modules/es.object.entries.js");
/* harmony import */ var core_js_modules_es_object_entries__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_entries__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! core-js/modules/es.string.iterator */ "./node_modules/core-js/modules/es.string.iterator.js");
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! core-js/modules/web.dom-collections.iterator */ "./node_modules/core-js/modules/web.dom-collections.iterator.js");
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var _babel_runtime_helpers_slicedToArray__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! @babel/runtime/helpers/slicedToArray */ "./node_modules/@babel/runtime/helpers/slicedToArray.js");
/* harmony import */ var _babel_runtime_helpers_slicedToArray__WEBPACK_IMPORTED_MODULE_18___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_slicedToArray__WEBPACK_IMPORTED_MODULE_18__);
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! @babel/runtime/helpers/toConsumableArray */ "./node_modules/@babel/runtime/helpers/toConsumableArray.js");
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19__);
/* harmony import */ var _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! @babel/runtime/helpers/classCallCheck */ "./node_modules/@babel/runtime/helpers/classCallCheck.js");
/* harmony import */ var _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_20___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_20__);
/* harmony import */ var _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! @babel/runtime/helpers/createClass */ "./node_modules/@babel/runtime/helpers/createClass.js");
/* harmony import */ var _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_21___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_21__);
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");























function _createForOfIteratorHelper(o) { if (typeof Symbol === "undefined" || o[Symbol.iterator] == null) { if (Array.isArray(o) || (o = _unsupportedIterableToArray(o))) { var i = 0; var F = function F() {}; return { s: F, n: function n() { if (i >= o.length) return { done: true }; return { done: false, value: o[i++] }; }, e: function e(_e) { throw _e; }, f: F }; } throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method."); } var it, normalCompletion = true, didErr = false, err; return { s: function s() { it = o[Symbol.iterator](); }, n: function n() { var step = it.next(); normalCompletion = step.done; return step; }, e: function e(_e2) { didErr = true; err = _e2; }, f: function f() { try { if (!normalCompletion && it["return"] != null) it["return"](); } finally { if (didErr) throw err; } } }; }

function _unsupportedIterableToArray(o, minLen) { if (!o) return; if (typeof o === "string") return _arrayLikeToArray(o, minLen); var n = Object.prototype.toString.call(o).slice(8, -1); if (n === "Object" && o.constructor) n = o.constructor.name; if (n === "Map" || n === "Set") return Array.from(o); if (n === "Arguments" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return _arrayLikeToArray(o, minLen); }

function _arrayLikeToArray(arr, len) { if (len == null || len > arr.length) len = arr.length; for (var i = 0, arr2 = new Array(len); i < len; i++) { arr2[i] = arr[i]; } return arr2; }


var COMPUTER_RESOURCES_TITLE = [{
  k: 'GPU',
  v: _langs__WEBPACK_IMPORTED_MODULE_22__.i18n.t('computeResourceTitle.GPU')
}, {
  k: 'NPU',
  v: _langs__WEBPACK_IMPORTED_MODULE_22__.i18n.t('computeResourceTitle.NPU')
}, {
  k: 'GCU',
  v: _langs__WEBPACK_IMPORTED_MODULE_22__.i18n.t('computeResourceTitle.GCU')
}, {
  k: 'MLU',
  v: _langs__WEBPACK_IMPORTED_MODULE_22__.i18n.t('computeResourceTitle.MLU')
}, {
  k: 'DCU',
  v: _langs__WEBPACK_IMPORTED_MODULE_22__.i18n.t('computeResourceTitle.DCU')
}, {
  k: 'ILUVATAR-GPGPU',
  v: _langs__WEBPACK_IMPORTED_MODULE_22__.i18n.t('computeResourceTitle.ILUVATAR-GPGPU')
}, {
  k: 'METAX-GPGPU',
  v: _langs__WEBPACK_IMPORTED_MODULE_22__.i18n.t('computeResourceTitle.METAX-GPGPU')
}, {
  k: 'BIREN-GPU',
  v: _langs__WEBPACK_IMPORTED_MODULE_22__.i18n.t('computeResourceTitle.BIREN-GPU')
}];
var CreatePageConfigs = {
  // 调试任务
  'DEBUG': [{
    'C2Net': [{
      'GPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: false
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {},
          runTimeLimit: {
            required: true
          },
          repo: {
            required: false
          }
        },
        showMindTorchHelper: true,
        hideCluster: true
      }],
      'NPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: false
          },
          // imagev2: { required: true, relatedSpec: true },
          imagev1: {
            required: true,
            type: 2,
            useId: true
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {
            required: true
          },
          runTimeLimit: {
            required: true
          },
          repo: {
            required: false
          }
        },
        hideCluster: true
      }],
      'GCU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: false
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {
            required: true
          },
          runTimeLimit: {
            required: true
          },
          repo: {
            required: false
          }
        },
        hideCluster: true
      }],
      'MLU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: false
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {
            required: true
          },
          runTimeLimit: {
            required: true
          },
          repo: {
            required: false
          }
        },
        hideCluster: true
      }],
      'DCU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: false
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {
            required: true
          },
          runTimeLimit: {
            required: true
          },
          repo: {
            required: false
          }
        },
        hideCluster: true
      }],
      'ILUVATAR-GPGPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: false
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {
            required: true
          },
          runTimeLimit: {
            required: true
          },
          repo: {
            required: false
          }
        },
        hideCluster: true
      }],
      'METAX-GPGPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: false
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {
            required: true
          },
          runTimeLimit: {
            required: true
          },
          repo: {
            required: false
          }
        },
        hideCluster: true
      }],
      'BIREN-GPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: false
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {
            required: true
          },
          runTimeLimit: {
            required: true
          },
          repo: {
            required: false
          }
        },
        hideCluster: true
      }]
    }]
  }],
  // 训练任务
  'TRAIN': [{
    'C2Net': [{
      'GPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: true
          },
          model: {
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/gpu_mnist_example/train.py'
          },
          dataset: {
            required: true
          },
          runParameters: {
            required: false
          },
          networkType: {
            required: true
          },
          visualization: {
            required: false
          },
          spec: {
            required: true
          },
          workServerNum: {
            required: true
          },
          repo: {
            required: true
          }
        },
        showMindTorchHelper: true,
        hideCluster: true
      }],
      'NPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: true
          },
          model: {
            required: false,
            multiple: true
          },
          // imagev2: { required: true, relatedSpec: true },
          imagev1: {
            required: true,
            type: 2,
            useId: true
          },
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/npu_mnist_example/train.py'
          },
          dataset: {
            required: true
          },
          runParameters: {
            required: false
          },
          networkType: {
            required: true
          },
          visualization: {
            required: false
          },
          spec: {
            required: true
          },
          workServerNum: {
            required: true
          },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'GCU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: true
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/gcu_mnist_example/train.py'
          },
          dataset: {
            required: true
          },
          runParameters: {
            required: false
          },
          networkType: {
            required: true
          },
          visualization: {
            required: false
          },
          spec: {
            required: true
          },
          workServerNum: {
            required: true
          },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'DCU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: true
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          //imagev2: { required: true, relatedSpec: true },
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/dcu_mnist_example/train.py'
          },
          dataset: {
            required: true
          },
          runParameters: {
            required: false
          },
          networkType: {
            required: true
          },
          visualization: {
            required: false
          },
          spec: {
            required: true
          },
          workServerNum: {
            required: true
          },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'ILUVATAR-GPGPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: true
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/gpgpu_mnist_example/train.py'
          },
          dataset: {
            required: true
          },
          runParameters: {
            required: false
          },
          networkType: {
            required: true
          },
          visualization: {
            required: false
          },
          spec: {
            required: true
          },
          workServerNum: {
            required: true
          },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'METAX-GPGPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: true
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/OpenI_Cloudbrain_Example/src/branch/master/gpgpu_mnist_example/train.py'
          },
          dataset: {
            required: true
          },
          runParameters: {
            required: false
          },
          networkType: {
            required: true
          },
          visualization: {
            required: false
          },
          spec: {
            required: true
          },
          workServerNum: {
            required: true
          },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'BIREN-GPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: true
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          bootFile: {
            required: true,
            sampleUrl: ''
          },
          dataset: {
            required: true
          },
          runParameters: {
            required: false
          },
          networkType: {
            required: true
          },
          visualization: {
            required: false
          },
          spec: {
            required: true
          },
          workServerNum: {
            required: true
          },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }]
    }]
  }],
  //在线推理
  'ONLINEINFERENCE': [{
    'C2Net': [{
      'GPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {},
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example'
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {},
          selfSshAddress: {
            required: false
          },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'NPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: true
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example'
          },
          // imagev2: { required: true, relatedSpec: true },
          imagev1: {
            required: true,
            type: 2,
            useId: true
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {
            required: true
          },
          // selfSshAddress: { required: false },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'GCU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: true
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example'
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {
            required: true
          },
          selfSshAddress: {
            required: false
          },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'MLU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {},
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example'
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {},
          // selfSshAddress: { required: false },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'DCU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {},
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example'
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {},
          // selfSshAddress: { required: false },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'ILUVATAR-GPGPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {},
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example'
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {},
          // selfSshAddress: { required: false },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'METAX-GPGPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {},
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example'
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {},
          // selfSshAddress: { required: false },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }],
      'BIREN-GPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {},
          bootFile: {
            required: true,
            sampleUrl: 'https://openi.pcl.ac.cn/OpenIOSSG/Online-Inference_Example'
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {},
          // selfSshAddress: { required: false },
          repo: {
            required: true
          }
        },
        hideCluster: true
      }]
    }]
  }],
  // 通用任务
  'GENERAL': [{
    'C2Net': [{
      'GPU': [{
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: false
          },
          model: {
            required: false,
            multiple: true,
            useExceedSize: true
          },
          imagev1: {
            required: true,
            type: -1
          },
          dataset: {
            required: false,
            useExceedSize: true
          },
          networkType: {
            required: true
          },
          spec: {},
          repo: {
            required: false
          }
        },
        hideCluster: true
      }]
    }]
  }],
  'HPC': [{
    'C2Net': [{
      'CPU': [{
        appName: 'MMLSpark',
        form: {
          taskName: {
            required: true
          },
          taskDescr: {
            required: false
          },
          branchName: {
            required: false
          },
          model: {
            required: false,
            multiple: true
          },
          imagev1: {
            required: true,
            type: -1,
            useId: true
          },
          // imagev2: { required: true },
          dataset: {
            required: false
          },
          spec: {
            required: true
          },
          repo: {
            required: false
          }
        },
        hideCluster: true
      }]
    }]
  }]
}; // 基础配置类

var CreatePageConfigManager = /*#__PURE__*/function () {
  function CreatePageConfigManager(config) {
    _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_20___default()(this, CreatePageConfigManager);

    this.config = config;
    this.cache = new Map(); // 缓存查询结果
  }
  /**
   * 获取任务类型配置
   * @param {string} taskType - 任务类型: DEBUG, TRAIN, ONLINEINFERENCE, GENERAL, HPC
   * @returns {Array|null} 任务类型配置
   */


  _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_21___default()(CreatePageConfigManager, [{
    key: "getTaskTypeConfig",
    value: function getTaskTypeConfig(taskType) {
      var cacheKey = "taskType_".concat(taskType);

      if (this.cache.has(cacheKey)) {
        return this.cache.get(cacheKey);
      }

      var config = this.config[taskType];
      this.cache.set(cacheKey, config);
      return config;
    }
    /**
     * 获取集群类型配置
     * @param {string} taskType - 任务类型
     * @param {string} clusterType - 集群类型 (目前主要是 C2Net)
     * @returns {Array|null} 集群类型配置
     */

  }, {
    key: "getClusterConfig",
    value: function getClusterConfig(taskType) {
      var clusterType = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 'C2Net';
      var cacheKey = "network_".concat(taskType, "_").concat(clusterType);

      if (this.cache.has(cacheKey)) {
        console.log("cache cacheKey");
        return this.cache.get(cacheKey);
      }

      var taskConfig = this.getTaskTypeConfig(taskType);
      if (!taskConfig) return null;
      var clusterConfig = taskConfig.find(function (item) {
        return item[clusterType];
      });
      var result = clusterConfig ? clusterConfig[clusterType] : null;
      this.cache.set(cacheKey, result);
      return result;
    }
    /**
     * 获取具体资源配置
     * @param {string} taskType - 任务类型
     * @param {string} resourceType - 资源类型: GPU, NPU, GCU, etc.
     * @param {string} clusterType - 网络类型
     * @returns {Object|null} 资源配置
     */

  }, {
    key: "getResourceConfig",
    value: function getResourceConfig(taskType, resourceType) {
      var clusterType = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : 'C2Net';
      var cacheKey = "resource_".concat(taskType, "_").concat(clusterType, "_").concat(resourceType);

      if (this.cache.has(cacheKey)) {
        return this.cache.get(cacheKey);
      }

      var clusterConfig = this.getClusterConfig(taskType, clusterType);
      if (!clusterConfig || !clusterConfig[0]) return null;
      var resourceConfig = clusterConfig[0][resourceType];
      var result = resourceConfig && resourceConfig[0] ? resourceConfig[0] : null;
      this.cache.set(cacheKey, result);
      return result;
    }
    /**
     * 获取计算资源列表
     * @param {string} taskType - 任务类型
     * @param {string} networkType - 网络类型
     * @returns {Array|null} 可用的计算资源列表
     */

  }, {
    key: "getComputerResources",
    value: function getComputerResources(taskType) {
      var clusterType = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 'C2Net';
      var cacheKey = "resources_".concat(taskType, "_").concat(clusterType);

      if (this.cache.has(cacheKey)) {
        return this.cache.get(cacheKey);
      }

      var clusterConfig = this.getClusterConfig(taskType, clusterType);
      if (!clusterConfig || !clusterConfig[0]) return null;
      var result = Object.keys(clusterConfig[0]) || null;
      this.cache.set(cacheKey, result);
      return result;
    }
    /**
     * 获取所有支持的任务类型
     * @returns {string[]} 任务类型列表
     */

  }, {
    key: "getAllTaskTypes",
    value: function getAllTaskTypes() {
      return Object.keys(this.config);
    }
    /**
     * 获取任务支持的集群类型
     * @param {string} taskType - 任务类型
     * @returns {string[]} 任务类型列表
     */

  }, {
    key: "getTaskTypeAllClusters",
    value: function getTaskTypeAllClusters(taskType) {
      var taskConfig = this.getTaskTypeConfig(taskType);
      console.log(taskConfig);
      if (!taskConfig) return null;
      return Object.keys(taskConfig[0]);
    }
    /**
     * 清空缓存
     */

  }, {
    key: "clearCache",
    value: function clearCache() {
      this.cache.clear();
    }
  }]);

  return CreatePageConfigManager;
}(); // 创建配置管理器实例


var configCreateManager = new CreatePageConfigManager(CreatePageConfigs);
var FieldTemplates = {
  // 基础字段组
  basicFields: ['taskName', 'creator', 'descr'],
  resourceFields: ['aiCenter', 'computerRes', 'spec'],
  paramsFields1: ['imagev1', 'repo', 'branch'],
  paramsFields2: ['datasetList', 'modelList'],
  statusFields: ['status', 'createTime', 'startTime', 'endTime', 'duration'] // 特殊字段组

};
var DetailPageConfigs = {
  // 调试任务
  'DEBUG': [{
    listUrl: 'debugjob?debugListType=all',
    clusters: ['OpenI', 'C2Net'],
    'OpenI': [{
      'GPU': [{
        detailUrl: 'cloudbrain/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'timeLimit']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }]
      }],
      'NPU': [{
        detailUrl: 'modelarts/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'timeLimit']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          },
          showDatasetDownload: false,
          showModelFileDownload: false
        }, {
          name: 'resultDownload'
        }]
      }]
    }],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'timeLimit']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }]
      }],
      'NPU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'timeLimit']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }]
      }],
      'GCU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'timeLimit']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }]
      }],
      'MLU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'timeLimit']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }]
      }],
      'DCU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'timeLimit']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }]
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'timeLimit']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }]
      }],
      'METAX-GPGPU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'timeLimit']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }]
      }],
      'BIREN-GPU': [{
        detailUrl: 'grampus/notebook/',
        summary: [],
        operations: ['debug', 'redebug', 'stop', 'saveTaskTmpl'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'timeLimit']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'operationProfile'
        }, {
          name: 'resultDownload'
        }]
      }]
    }]
  }],
  // 训练任务
  'TRAIN': [{
    listUrl: 'modelarts/train-job?listType=all',
    clusters: ['OpenI', 'C2Net'],
    'OpenI': [{
      'GPU': [{
        detailUrl: 'cloudbrain/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'workServerNum', 'visualization']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile', 'runVersion', 'runParameters'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs'
        }, {
          name: 'resultDownload'
        }]
      }],
      'NPU': [{
        detailUrl: 'modelarts/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'workServerNum', 'visualization']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile', 'runVersion', 'runParameters'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'logs',
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }]
    }],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'workServerNum', 'visualization']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile', 'runVersion', 'runParameters'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'NPU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'workServerNum', 'visualization']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile', 'runVersion', 'runParameters'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'GCU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'workServerNum', 'visualization']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile', 'runVersion', 'runParameters'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'DCU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'workServerNum', 'visualization']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile', 'runVersion', 'runParameters'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'MLU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'workServerNum', 'visualization']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile', 'runVersion', 'runParameters'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'workServerNum', 'visualization']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile', 'runVersion', 'runParameters'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'METAX-GPGPU': [{
        detailUrl: 'grampus/train-job/',
        summary: [],
        operations: ['stop', 'tensorBoard', 'saveTaskTmpl', 'saveModel', 'exportDataset'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet', 'workServerNum', 'visualization']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile', 'runVersion', 'runParameters'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            showSdkCode: [true]
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }]
    }]
  }],
  // 推理任务
  'INFERENCE': [{
    listUrl: 'modelarts/inference-job',
    clusters: ['OpenI', 'C2Net'],
    'OpenI': [{
      'NPU': [{
        detailUrl: 'modelarts/inference-job/',
        summary: [],
        operations: [],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'logs',
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }]
    }],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/inference-job/',
        summary: [],
        operations: [],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'NPU': [{
        detailUrl: 'grampus/inference-job/',
        summary: [],
        operations: [],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/inference-job/',
        summary: [],
        operations: [],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }]
      }]
    }]
  }],
  'ONLINEINFERENCE': [{
    listUrl: 'grampus/onlineinfer',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'GCU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'NPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'MLU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'DCU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'METAX-GPGPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'BIREN-GPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['consoleHome', 'onlineInfer', 'reinfer', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'resultDownload'
        }]
      }]
    }]
  }],
  'MODELEXPERIENCE': [{
    listUrl: '',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2), ['sourceFtName']),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }]
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2), ['sourceFtName']),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }]
      }],
      'GCU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2), ['sourceFtName']),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }]
      }],
      'NPU': [{
        detailUrl: 'grampus/onlineinfer',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2), ['sourceFtName']),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }]
      }]
    }]
  }],
  'FINETUNE': [{
    listUrl: '',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['stop', 'saveModel', "deployModel"],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }, {
          name: 'loss'
        }]
      }],
      'NPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['stop', 'saveModel', 'deployModel'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }, {
          name: 'loss'
        }]
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['stop', 'saveModel'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'resourceUseage',
          multiNodes: true
        }, {
          name: 'resultDownload'
        }, {
          name: 'loss'
        }]
      }]
    }]
  }],
  'SDFINETUNE': [{
    listUrl: '',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlineLoraTrain', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'resultDownload'
        }]
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }]
      }],
      'GCU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }]
      }]
    }]
  }],
  'ComfyuiExperience': [{
    listUrl: '',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlineWorkflow', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }]
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }]
      }],
      'GCU': [{
        detailUrl: 'grampus/sftfinetune',
        summary: [],
        operations: ['onlinexperience', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }]
      }]
    }]
  }],
  'EVAL': [{
    listUrl: '',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/evaluate',
        summary: [],
        operations: ['stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2), ['sourceFtName']),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'evalOverview'
        }, {
          name: 'evalDetail'
        }]
      }],
      'ILUVATAR-GPGPU': [{
        detailUrl: 'grampus/evaluate',
        summary: [],
        operations: ['stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2), ['sourceFtName']),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'evalOverview'
        }, {
          name: 'evalDetail'
        }]
      }],
      'GCU': [{
        detailUrl: 'grampus/evaluate',
        summary: [],
        operations: ['stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2), ['sourceFtName']),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'logs',
          noScroll: true
        }, {
          name: 'evalOverview'
        }, {
          name: 'evalDetail'
        }]
      }],
      'NPU': [{
        detailUrl: 'grampus/evaluate',
        summary: [],
        operations: ['stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), ['bootFile'], _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2), ['sourceFtName']),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }, {
          name: 'logs',
          noScroll: true,
          multiNodes: true
        }, {
          name: 'evalOverview'
        }, {
          name: 'evalDetail'
        }]
      }]
    }]
  }],
  // 通用任务
  'GENERAL': [{
    listUrl: 'grampus/general',
    clusters: ['C2Net'],
    'C2Net': [{
      'GPU': [{
        detailUrl: 'grampus/general/',
        summary: [],
        operations: ['debug', 'stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields),
            resourceFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields), ['hasInternet']),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields),
            generalTaskCodeTips: [true]
          }
        }, {
          name: 'operationProfile'
        }]
      }]
    }]
  }],
  'HPC': [{
    listUrl: 'supercompute/job',
    clusters: ['C2Net'],
    'C2Net': [{
      'CPU': [{
        detailUrl: 'supercompute/job',
        summary: [],
        operations: ['stop'],
        tabs: [{
          name: 'configInfo',
          fields: {
            basicFiled: ['appName'].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.basicFields)),
            resourceFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.resourceFields),
            paramsFiled: [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields1), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.paramsFields2)),
            runstatusFiled: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_19___default()(FieldTemplates.statusFields)
          }
        }, {
          name: 'operationProfile'
        }]
      }]
    }]
  }]
};

var DetailPageConfigManager = /*#__PURE__*/function () {
  function DetailPageConfigManager(configs) {
    _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_20___default()(this, DetailPageConfigManager);

    this.configMap = new Map();
    this.initialize(configs);
  }

  _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_21___default()(DetailPageConfigManager, [{
    key: "initialize",
    value: function initialize(configs) {
      // 遍历所有配置，建立多层索引
      for (var _i = 0, _Object$entries = Object.entries(configs); _i < _Object$entries.length; _i++) {
        var _Object$entries$_i = _babel_runtime_helpers_slicedToArray__WEBPACK_IMPORTED_MODULE_18___default()(_Object$entries[_i], 2),
            taskType = _Object$entries$_i[0],
            clustersConfig = _Object$entries$_i[1];

        var taskMap = new Map();

        var _iterator = _createForOfIteratorHelper(clustersConfig),
            _step;

        try {
          for (_iterator.s(); !(_step = _iterator.n()).done;) {
            var clusterConfig = _step.value;

            for (var _i2 = 0, _Object$entries2 = Object.entries(clusterConfig); _i2 < _Object$entries2.length; _i2++) {
              var _Object$entries2$_i = _babel_runtime_helpers_slicedToArray__WEBPACK_IMPORTED_MODULE_18___default()(_Object$entries2[_i2], 2),
                  clusterName = _Object$entries2$_i[0],
                  hardwareConfigs = _Object$entries2$_i[1];

              if (clusterName === 'listUrl' || clusterName === 'clusters') continue;
              var clusterMap = new Map();

              var _iterator2 = _createForOfIteratorHelper(hardwareConfigs),
                  _step2;

              try {
                for (_iterator2.s(); !(_step2 = _iterator2.n()).done;) {
                  var hardwareConfig = _step2.value;

                  for (var _i3 = 0, _Object$entries3 = Object.entries(hardwareConfig); _i3 < _Object$entries3.length; _i3++) {
                    var _Object$entries3$_i = _babel_runtime_helpers_slicedToArray__WEBPACK_IMPORTED_MODULE_18___default()(_Object$entries3[_i3], 2),
                        hardwareType = _Object$entries3$_i[0],
                        configArray = _Object$entries3$_i[1];

                    // 每个硬件类型对应一个配置对象
                    clusterMap.set(hardwareType, configArray[0]);
                  }
                }
              } catch (err) {
                _iterator2.e(err);
              } finally {
                _iterator2.f();
              }

              taskMap.set(clusterName, clusterMap);
            }
          }
        } catch (err) {
          _iterator.e(err);
        } finally {
          _iterator.f();
        }

        this.configMap.set(taskType, {
          taskMap: taskMap,
          listUrl: clustersConfig[0].listUrl,
          clusters: clustersConfig[0].clusters
        });
      }
    } // 快速获取配置

  }, {
    key: "getConfig",
    value: function getConfig(taskType, cluster, hardware) {
      var taskConfig = this.configMap.get(taskType);
      if (!taskConfig) return null;
      var clusterConfig = taskConfig.taskMap.get(cluster);
      if (!clusterConfig) return null;
      return clusterConfig.get(hardware);
    } // 获取列表URL

  }, {
    key: "getListUrl",
    value: function getListUrl(taskType) {
      var _this$configMap$get;

      return (_this$configMap$get = this.configMap.get(taskType)) === null || _this$configMap$get === void 0 ? void 0 : _this$configMap$get.listUrl;
    } // 获取支持的集群

  }, {
    key: "getClusters",
    value: function getClusters(taskType) {
      var _this$configMap$get2;

      return (_this$configMap$get2 = this.configMap.get(taskType)) === null || _this$configMap$get2 === void 0 ? void 0 : _this$configMap$get2.clusters;
    } // 获取特定集群支持的硬件类型

  }, {
    key: "getHardwareTypes",
    value: function getHardwareTypes(taskType, cluster) {
      var taskConfig = this.configMap.get(taskType);
      if (!taskConfig) return [];
      var clusterConfig = taskConfig.taskMap.get(cluster);
      if (!clusterConfig) return [];
      return Array.from(clusterConfig.keys());
    }
  }]);

  return DetailPageConfigManager;
}();

var configDetailManager = new DetailPageConfigManager(DetailPageConfigs); // 快速访问
// const debugGPUConfig = configManager.getConfig('DEBUG', 'OpenI', 'GPU');

/***/ })

}]);