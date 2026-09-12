"use strict";
(self["webpackChunk"] = self["webpackChunk"] || []).push([["web_src_vuepages_components_AcknowledgementsTips_vue-web_src_vuepages_components_cloudbrain_F-65ba32"],{

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=script&lang=js":
/*!******************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=script&lang=js ***!
  \******************************************************************************************************************************************************************************************************************/
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
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "AcknowledgementsTips",
  props: {
    usePop: {
      type: Boolean,
      "default": false
    }
  },
  data: function data() {
    return {};
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=script&lang=js":
/*!**********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************************************************************************************************************************************************/
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
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'ForbidPenetrationTipCheck',
  props: {
    value: {
      type: Boolean,
      required: true
    },
    required: {
      type: Boolean,
      "default": false
    }
  },
  data: function data() {
    return {
      currentValue: false,
      errStatus: false
    };
  },
  watch: {
    value: {
      immediate: true,
      handler: function handler(newVal) {
        this.currentValue = !!newVal;
      }
    }
  },
  methods: {
    check: function check() {
      var isValid = this.currentValue || !this.required;
      this.errStatus = !isValid;
      return isValid;
    },
    handleChange: function handleChange(value) {
      this.currentValue = value;
      this.$emit('input', value);
      this.check();
    }
  },
  beforeMount: function beforeMount() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=script&lang=js":
/*!******************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=script&lang=js ***!
  \******************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var _pages_cloudbrain_configs__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ~/pages/cloudbrain/configs */ "./web_src/vuepages/pages/cloudbrain/configs.js");
/* harmony import */ var _const__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ~/const */ "./web_src/vuepages/const/index.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");


//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "FormTopV2",
  props: {
    queueNum: {
      type: Number,
      "default": 1
    },
    taskTypeObj: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    //静态资源-基础任务
    resourceObj: {
      type: Object,
      "default": function _default() {
        return {};
      }
    } //动态资源配置项-大模型基地任务

  },
  data: function data() {
    return {
      clusters: [],
      cluster: '',
      computerResouces: [],
      computerResouce: '',
      configs: {}
    };
  },
  watch: {
    taskTypeObj: {
      deep: true,
      handler: function handler(newVal) {
        if (JSON.stringify(newVal) !== '{}') {
          var _clusters = _pages_cloudbrain_configs__WEBPACK_IMPORTED_MODULE_2__.configCreateManager.getTaskTypeAllClusters(newVal.taskType);

          this.clusters = _clusters.map(function (item) {
            return {
              key: item,
              label: (0,_utils__WEBPACK_IMPORTED_MODULE_4__.getListValueWithKey)(_const__WEBPACK_IMPORTED_MODULE_3__.CLUSTERS, item)
            };
          });
          this.cluster = newVal.cluster;

          var _computerResouces = _pages_cloudbrain_configs__WEBPACK_IMPORTED_MODULE_2__.configCreateManager.getComputerResources(newVal.taskType);

          this.computerResouces = _computerResouces.map(function (item) {
            return {
              key: item,
              label: (0,_utils__WEBPACK_IMPORTED_MODULE_4__.getListValueWithKey)(_pages_cloudbrain_configs__WEBPACK_IMPORTED_MODULE_2__.COMPUTER_RESOURCES_TITLE, item)
            };
          });
          this.computerResouce = newVal.computerResouce;

          var _configParams = _pages_cloudbrain_configs__WEBPACK_IMPORTED_MODULE_2__.configCreateManager.getResourceConfig(newVal.taskType, newVal.computerResouce);

          this.configs = _configParams || {};
        }
      }
    },
    resourceObj: {
      deep: true,
      immediate: true,
      handler: function handler(newVal) {
        if (JSON.stringify(newVal) !== '{}') {
          this.computerResouces = (newVal === null || newVal === void 0 ? void 0 : newVal.computerResouces) || [];
          this.computerResouce = newVal.computerResouce || '';
          this.clusters = (newVal === null || newVal === void 0 ? void 0 : newVal.clusters) || [];
          this.cluster = (newVal === null || newVal === void 0 ? void 0 : newVal.cluster) || '';
          this.configs = (newVal === null || newVal === void 0 ? void 0 : newVal.configs) || {};
        }
      }
    }
  },
  methods: {
    check: function check() {
      return true;
    },
    changeCluster: function changeCluster(cluster) {
      this.cluster = cluster;
      return; // this.$emit('change', {
      //   cluster: cluster,
      //   computerResouce: '',
      // })
    },
    changeComputerResouce: function changeComputerResouce(computerResouce) {
      this.computerResouce = computerResouce;
      this.$emit('change', {
        cluster: this.cluster,
        computerResouce: this.computerResouce
      });
    }
  },
  mounted: function mounted() {// console.log('taskType', this.taskType)
    // const clusters = configCreateManager.getTaskTypeAllClusters(this.taskType)
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=script&lang=js":
/*!*******************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=script&lang=js ***!
  \*******************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.number.to-fixed */ "./node_modules/core-js/modules/es.number.to-fixed.js");
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");







//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "SpecSelect",
  props: {
    value: {
      type: String,
      required: true
    },
    configs: {
      type: Object,
      required: true
    },
    required: {
      type: Boolean,
      "default": true
    },
    networkType: {
      type: String,
      "default": 'no_internet'
    },
    visualize: {
      type: Boolean,
      "default": false
    },
    workServerNum: {
      type: Number,
      "default": 1
    },
    loading: {
      type: Boolean,
      "default": false
    },
    tips: {
      type: String,
      "default": ''
    }
  },
  data: function data() {
    return {
      list: [],
      spec: '',
      selIconType: '',
      selPointStr: '',
      showUseTime: false,
      canUseTime: '',
      errStatus: false
    };
  },
  watch: {
    value: {
      immediate: true,
      handler: function handler(newVal) {
        this.spec = newVal.toString();
        this.renderSpec();
      }
    },
    configs: {
      // {specs: [],  blance: 0, showPoint: true },
      immediate: true,
      deep: true,
      handler: function handler(newVal) {
        this.renderSpec();
      }
    },
    networkType: {
      immediate: true,
      handler: function handler() {
        var resetSpec = true;
        this.renderSpec(resetSpec);
      }
    },
    visualize: {
      immediate: true,
      handler: function handler() {
        var resetSpec = true;
        this.renderSpec(resetSpec);
      }
    },
    workServerNum: {
      immediate: true,
      handler: function handler() {
        this.renderSpec();
      }
    }
  },
  methods: {
    renderSpec: function renderSpec(resetSpec) {
      var showPoint = this.configs.showPoint || false;
      var specs = this.configs.specs[this.networkType] || [];

      if (this.visualize) {
        specs = specs.filter(function (item) {
          return item.VisualizeCapableQueuesExist;
        });
      }

      this.list = specs.map(function (item) {
        return (0,_utils__WEBPACK_IMPORTED_MODULE_7__.renderSpecObject)(item, showPoint);
      });

      if (resetSpec) {
        this.spec = specs.length ? specs[0].id.toString() : '';
      }

      this.changeSpec(this.spec);
    },
    changeSpec: function changeSpec() {
      var _this = this;

      var seldSpecItem = this.list.filter(function (item) {
        return item.id == _this.spec;
      })[0];

      if (seldSpecItem) {
        this.selIconType = seldSpecItem.type;
        this.selPointStr = seldSpecItem.pointStr;
        this.errStatus = false;
      } else {
        this.selIconType = '';
        this.selPointStr = '';
      }

      this.refreshPointInfo();
      this.$emit('input', this.spec);
      this.$emit('change', this.spec);
    },
    refreshPointInfo: function refreshPointInfo() {
      var _this2 = this;

      var showPoint = this.configs.showPoint || false;
      if (!showPoint) return;
      var seldSpecItem = this.list.filter(function (item) {
        return item.id == _this2.spec;
      })[0];

      if (seldSpecItem) {
        var unitPrice = seldSpecItem.unit_price;
        var blance = this.configs.blance || 0;
        var workServerNum = this.workServerNum;

        if (unitPrice == 0) {
          this.showUseTime = false;
        } else {
          var canUseTime = Number(blance) / (Number(unitPrice) * Number(workServerNum));

          if (Number(blance) < Number(unitPrice) * Number(workServerNum)) {
            // 余额不足一个单位单价时可用时间提示为 0
            canUseTime = 0;
          }

          this.canUseTime = canUseTime.toFixed(2);
          this.showUseTime = true;
        }
      }
    },
    check: function check() {
      var _this3 = this;

      var seldSpecItem = this.list.filter(function (item) {
        return item.id == _this3.spec;
      })[0];

      if (!seldSpecItem) {
        this.errStatus = true;
        return false;
      }

      this.errStatus = false;
      return true;
    },
    getSpecData: function getSpecData() {
      var _this4 = this;

      return this.list.filter(function (item) {
        return item.id == _this4.spec;
      })[0];
    }
  },
  beforeMount: function beforeMount() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=script&lang=js":
/*!************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=script&lang=js ***!
  \************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.includes */ "./node_modules/core-js/modules/es.array.includes.js");
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.is-array */ "./node_modules/core-js/modules/es.array.is-array.js");
/* harmony import */ var core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_some__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.some */ "./node_modules/core-js/modules/es.array.some.js");
/* harmony import */ var core_js_modules_es_array_some__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_some__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.promise */ "./node_modules/core-js/modules/es.promise.js");
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_promise_finally__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.promise.finally */ "./node_modules/core-js/modules/es.promise.finally.js");
/* harmony import */ var core_js_modules_es_promise_finally__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_promise_finally__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var dayjs__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! dayjs */ "./node_modules/dayjs/dayjs.min.js");
/* harmony import */ var dayjs__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(dayjs__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var _apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ~/apis/modules/cloudbrain */ "./web_src/vuepages/apis/modules/cloudbrain.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _const__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! ~/const */ "./web_src/vuepages/const/index.js");











//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: 'TaskLimitDialog',
  props: {
    visible: {
      type: Boolean,
      "default": false
    },
    title: {
      type: String,
      "default": ''
    },
    desc: {
      type: String,
      "default": ''
    },
    taskList: {
      type: Array,
      "default": function _default() {
        return [];
      }
    }
  },
  data: function data() {
    return {
      statusPollingTimer: null
    };
  },
  watch: {
    taskList: {
      handler: function handler() {
        this.checkAndStartPolling();
      },
      deep: true
    },
    visible: function visible(newVal) {
      if (!newVal) {
        this.stopStatusPolling();
      }
    }
  },
  methods: {
    // 将 jobTypeShow 转为行数组：数组直接返回，字符串按2字分行
    getJobTypeLines: function getJobTypeLines(jobTypeShow) {
      if (Array.isArray(jobTypeShow)) return jobTypeShow;
      if (!jobTypeShow) return []; // 中文按2字分行，英文按空格/单词分行

      var lines = []; // 检测是否包含中文

      if (/[\u4e00-\u9fff]/.test(jobTypeShow)) {
        for (var i = 0; i < jobTypeShow.length; i += 2) {
          lines.push(jobTypeShow.slice(i, i + 2));
        }
      } else {
        lines.push(jobTypeShow);
      }

      return lines;
    },
    dateFormat: function dateFormat(unix) {
      return dayjs__WEBPACK_IMPORTED_MODULE_11___default()(unix * 1000).format('YYYY-MM-DD HH:mm:ss');
    },
    // 判断任务是否在运行或等待中（可操作停止的状态）
    isTaskRunning: function isTaskRunning(task) {
      return task.status === 'WAITING' || task.status === 'RUNNING';
    },
    // 判断任务是否正在停止中
    isTaskStopping: function isTaskStopping(task) {
      return task.status === 'STOPPING';
    },
    // 判断任务是否可删除（只有终态才可以删除）
    canDeleteTask: function canDeleteTask(task) {
      var terminalStatuses = ['STOPPED', 'SUCCEEDED', 'FAILED', 'CREATE_FAILED'];
      return terminalStatuses.includes(task.status);
    },
    // 处理停止任务
    handleStopTask: function handleStopTask(task) {
      var _this = this;

      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_12__.stopAiTask)({
        repoOwnerName: task.owner_name,
        repoName: task.repo_name,
        id: task.id
      }).then(function (res) {
        res = res.data;

        if (res.code === 0) {
          _this.$message.success(_this.$t('notebook.stopSuccess'));

          _this.$emit('stopSuccess', task);
        } else {
          _this.$message.error(res.msg);
        }
      })["catch"](function () {
        _this.$message.error(_this.$t('notebook.stopFailed'));
      });
    },
    // 处理删除任务
    handleDeleteTask: function handleDeleteTask(task) {
      var _this2 = this;

      (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_12__.deleteAiTask)({
        repoOwnerName: task.owner_name,
        repoName: task.repo_name,
        id: task.id
      }).then(function (res) {
        res = res.data;

        if (res.code === 0) {
          _this2.$message.success(_this2.$t('notebook.deleteSuccess'));

          _this2.$emit('deleteSuccess', task);
        } else {
          _this2.$message.error(res.msg);
        }
      })["catch"](function () {
        _this2.$message.error(_this2.$t('notebook.deleteFailed'));
      });
    },
    // 检查是否有非终态任务并开启轮询（WAITING、RUNNING、STOPPING 需要轮询）
    checkAndStartPolling: function checkAndStartPolling() {
      var pollingStatuses = ['WAITING', 'RUNNING', 'STOPPING'];
      var hasPollingTask = this.taskList.some(function (item) {
        return pollingStatuses.includes(item.task.status);
      });

      if (hasPollingTask) {
        this.startStatusPolling();
      } else {
        this.stopStatusPolling();
      }
    },
    // 开启状态轮询
    startStatusPolling: function startStatusPolling() {
      var _this3 = this;

      if (this.statusPollingTimer) return;
      this.statusPollingTimer = setInterval(function () {
        _this3.pollTaskStatus();
      }, 10000);
    },
    // 停止状态轮询
    stopStatusPolling: function stopStatusPolling() {
      if (this.statusPollingTimer) {
        clearInterval(this.statusPollingTimer);
        this.statusPollingTimer = null;
      }
    },
    // 轮询任务状态（逐个查询非终态任务）
    pollTaskStatus: function pollTaskStatus() {
      var _this4 = this;

      var pollingStatuses = ['WAITING', 'RUNNING', 'STOPPING'];
      var pollingTasks = this.taskList.filter(function (item) {
        return pollingStatuses.includes(item.task.status);
      });

      if (!pollingTasks.length) {
        this.stopStatusPolling();
        this.$emit('allTasksTerminal');
        return;
      }

      var pendingCount = pollingTasks.length;
      pollingTasks.forEach(function (item) {
        (0,_apis_modules_cloudbrain__WEBPACK_IMPORTED_MODULE_12__.getAiTask)({
          id: item.task.id
        }).then(function (res) {
          var _res$data;

          res = res.data;

          if (res.code === 0 && ((_res$data = res.data) === null || _res$data === void 0 ? void 0 : _res$data.task)) {
            var newTask = res.data.task;
            newTask.computeSourceShow = newTask.compute_source == 'GPU' ? 'CPU/GPU' : newTask.compute_source;
            newTask.accCardTypeShow = (0,_utils__WEBPACK_IMPORTED_MODULE_13__.getListValueWithKey)(_const__WEBPACK_IMPORTED_MODULE_14__.ACC_CARD_TYPE, newTask.acc_card_type);

            if (newTask.job_type === 'ComfyuiExperience') {
              newTask.jobTypeShow = ['Comfy', 'UI'];
            } else if (newTask.job_type === 'EVAL') {
              newTask.jobTypeShow = _this4.$t('modelSquare.modelEvaluate');
            } else if (newTask.job_type === 'FINETUNE') {
              var evalText = _this4.$t('modelSquare.sftFinetune');

              newTask.jobTypeShow = [evalText.slice(0, 3), evalText.slice(3)];
            } else {
              newTask.jobTypeShow = (0,_utils__WEBPACK_IMPORTED_MODULE_13__.getListValueWithKey)(_const__WEBPACK_IMPORTED_MODULE_14__.JOB_TYPE, newTask.job_type);
            }

            _this4.$emit('taskStatusUpdate', newTask);
          }
        })["catch"](function (err) {
          console.log('pollTaskStatus error', err);
        })["finally"](function () {
          pendingCount--;

          if (pendingCount === 0) {
            _this4.$nextTick(function () {
              var stillPolling = _this4.taskList.some(function (item) {
                return pollingStatuses.includes(item.task.status);
              });

              if (!stillPolling) {
                _this4.stopStatusPolling();

                _this4.$emit('allTasksTerminal');
              }
            });
          }
        });
      });
    },
    handleClose: function handleClose() {
      this.stopStatusPolling();
      this.$emit('close');
    }
  },
  beforeDestroy: function beforeDestroy() {
    this.stopStatusPolling();
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=script&lang=js":
/*!*****************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=script&lang=js ***!
  \*****************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.date.now */ "./node_modules/core-js/modules/es.date.now.js");
/* harmony import */ var core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.number.to-fixed */ "./node_modules/core-js/modules/es.number.to-fixed.js");
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var dayjs__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! dayjs */ "./node_modules/dayjs/dayjs.min.js");
/* harmony import */ var dayjs__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(dayjs__WEBPACK_IMPORTED_MODULE_8__);








//
//
//
//
//
//
//
//
//
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
  name: 'TaskName',
  props: {
    value: {
      type: String,
      required: true
    },
    userName: {
      type: String,
      required: true
    },
    autofocus: {
      type: Boolean,
      "default": true
    },
    generate: {
      type: Boolean,
      "default": false
    },
    required: {
      type: Boolean,
      "default": true
    },
    type: {
      type: String,
      "default": ''
    }
  },
  data: function data() {
    return {
      currentValue: '',
      errStatus: false
    };
  },
  watch: {
    value: {
      immediate: true,
      handler: function handler(newVal) {
        newVal = newVal === undefined ? '' : newVal;
        this.currentValue = newVal.toString();
      }
    }
  },
  methods: {
    check: function check() {
      var reg = /^[a-z0-9][a-z0-9\-_]{1,34}[a-z0-9\-]$/;
      var reg1 = /^[a-z][a-z0-9\-]{4,25}$/;

      if (this.type == 'supercompute') {
        this.errStatus = !reg1.test(this.currentValue);
      } else {
        this.errStatus = !reg.test(this.currentValue);
      }

      if (!this.required && this.currentValue == '') {
        this.errStatus = false;
      }

      return !this.errStatus;
    },
    generateName: function generateName() {
      var str = this.userName.toLocaleLowerCase();
      var reg1 = /[^a-z0-9_\-]+/g;
      var reg2 = /^[_\-]+/g;
      var reg3 = /[_]+$/g;
      str = str.replace(reg1, '').replace(reg2, '').replace(reg3, '');
      str = str.slice(0, 5);
      var now = Date.now();
      return str + dayjs__WEBPACK_IMPORTED_MODULE_8___default()(now).format('YYYYMMDDHH') + (now / 1000).toFixed(0).slice(-5);
    },
    handleInput: function handleInput(value) {
      this.currentValue = value;
      this.$emit('input', value);
      this.check();
    },
    handleInputChange: function handleInputChange(value) {
      this.$emit('change', value);
    }
  },
  beforeMount: function beforeMount() {
    if (!this.value && this.generate) {
      this.currentValue = this.generateName();
      this.$emit('input', this.currentValue);
    }
  }
});

/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=style&index=0&id=aeeb1646&scoped=true&lang=less":
/*!************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=style&index=0&id=aeeb1646&scoped=true&lang=less ***!
  \************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=0&id=0f089c42&scoped=true&lang=less":
/*!****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=0&id=0f089c42&scoped=true&lang=less ***!
  \****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=1&id=0f089c42&lang=less":
/*!****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=1&id=0f089c42&lang=less ***!
  \****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=style&index=0&id=67f642aa&scoped=true&lang=less":
/*!************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=style&index=0&id=67f642aa&scoped=true&lang=less ***!
  \************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=style&index=0&id=0e79f3ae&scoped=true&lang=less":
/*!*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=style&index=0&id=0e79f3ae&scoped=true&lang=less ***!
  \*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=style&index=0&id=6151937c&scoped=true&lang=less":
/*!******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=style&index=0&id=6151937c&scoped=true&lang=less ***!
  \******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=style&index=0&id=09a23d7c&scoped=true&lang=less":
/*!***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=style&index=0&id=09a23d7c&scoped=true&lang=less ***!
  \***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./web_src/vuepages/components/AcknowledgementsTips.vue":
/*!**************************************************************!*\
  !*** ./web_src/vuepages/components/AcknowledgementsTips.vue ***!
  \**************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _AcknowledgementsTips_vue_vue_type_template_id_aeeb1646_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./AcknowledgementsTips.vue?vue&type=template&id=aeeb1646&scoped=true */ "./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=template&id=aeeb1646&scoped=true");
/* harmony import */ var _AcknowledgementsTips_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./AcknowledgementsTips.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=script&lang=js");
/* harmony import */ var _AcknowledgementsTips_vue_vue_type_style_index_0_id_aeeb1646_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./AcknowledgementsTips.vue?vue&type=style&index=0&id=aeeb1646&scoped=true&lang=less */ "./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=style&index=0&id=aeeb1646&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _AcknowledgementsTips_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _AcknowledgementsTips_vue_vue_type_template_id_aeeb1646_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _AcknowledgementsTips_vue_vue_type_template_id_aeeb1646_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "aeeb1646",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/AcknowledgementsTips.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue":
/*!******************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue ***!
  \******************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ForbidPenetrationTipCheck_vue_vue_type_template_id_0f089c42_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ForbidPenetrationTipCheck.vue?vue&type=template&id=0f089c42&scoped=true */ "./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=template&id=0f089c42&scoped=true");
/* harmony import */ var _ForbidPenetrationTipCheck_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ForbidPenetrationTipCheck.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=script&lang=js");
/* harmony import */ var _ForbidPenetrationTipCheck_vue_vue_type_style_index_0_id_0f089c42_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ForbidPenetrationTipCheck.vue?vue&type=style&index=0&id=0f089c42&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=0&id=0f089c42&scoped=true&lang=less");
/* harmony import */ var _ForbidPenetrationTipCheck_vue_vue_type_style_index_1_id_0f089c42_lang_less__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./ForbidPenetrationTipCheck.vue?vue&type=style&index=1&id=0f089c42&lang=less */ "./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=1&id=0f089c42&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! !../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;



/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_4__["default"])(
  _ForbidPenetrationTipCheck_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ForbidPenetrationTipCheck_vue_vue_type_template_id_0f089c42_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ForbidPenetrationTipCheck_vue_vue_type_template_id_0f089c42_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "0f089c42",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/FormTopV2.vue":
/*!**************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/FormTopV2.vue ***!
  \**************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _FormTopV2_vue_vue_type_template_id_67f642aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./FormTopV2.vue?vue&type=template&id=67f642aa&scoped=true */ "./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=template&id=67f642aa&scoped=true");
/* harmony import */ var _FormTopV2_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./FormTopV2.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=script&lang=js");
/* harmony import */ var _FormTopV2_vue_vue_type_style_index_0_id_67f642aa_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./FormTopV2.vue?vue&type=style&index=0&id=67f642aa&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=style&index=0&id=67f642aa&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _FormTopV2_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _FormTopV2_vue_vue_type_template_id_67f642aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _FormTopV2_vue_vue_type_template_id_67f642aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "67f642aa",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/FormTopV2.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/SpecSelect.vue":
/*!***************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/SpecSelect.vue ***!
  \***************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _SpecSelect_vue_vue_type_template_id_0e79f3ae_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./SpecSelect.vue?vue&type=template&id=0e79f3ae&scoped=true */ "./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=template&id=0e79f3ae&scoped=true");
/* harmony import */ var _SpecSelect_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./SpecSelect.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=script&lang=js");
/* harmony import */ var _SpecSelect_vue_vue_type_style_index_0_id_0e79f3ae_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./SpecSelect.vue?vue&type=style&index=0&id=0e79f3ae&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=style&index=0&id=0e79f3ae&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _SpecSelect_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _SpecSelect_vue_vue_type_template_id_0e79f3ae_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _SpecSelect_vue_vue_type_template_id_0e79f3ae_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "0e79f3ae",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/SpecSelect.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue":
/*!********************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue ***!
  \********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _TaskLimitDialog_vue_vue_type_template_id_6151937c_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./TaskLimitDialog.vue?vue&type=template&id=6151937c&scoped=true */ "./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=template&id=6151937c&scoped=true");
/* harmony import */ var _TaskLimitDialog_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./TaskLimitDialog.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=script&lang=js");
/* harmony import */ var _TaskLimitDialog_vue_vue_type_style_index_0_id_6151937c_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./TaskLimitDialog.vue?vue&type=style&index=0&id=6151937c&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=style&index=0&id=6151937c&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _TaskLimitDialog_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _TaskLimitDialog_vue_vue_type_template_id_6151937c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _TaskLimitDialog_vue_vue_type_template_id_6151937c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "6151937c",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/TaskName.vue":
/*!*************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/TaskName.vue ***!
  \*************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _TaskName_vue_vue_type_template_id_09a23d7c_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./TaskName.vue?vue&type=template&id=09a23d7c&scoped=true */ "./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=template&id=09a23d7c&scoped=true");
/* harmony import */ var _TaskName_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./TaskName.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=script&lang=js");
/* harmony import */ var _TaskName_vue_vue_type_style_index_0_id_09a23d7c_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./TaskName.vue?vue&type=style&index=0&id=09a23d7c&scoped=true&lang=less */ "./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=style&index=0&id=09a23d7c&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _TaskName_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _TaskName_vue_vue_type_template_id_09a23d7c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _TaskName_vue_vue_type_template_id_09a23d7c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "09a23d7c",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/cloudbrain/TaskName.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=script&lang=js":
/*!**************************************************************************************!*\
  !*** ./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=script&lang=js ***!
  \**************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_AcknowledgementsTips_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./AcknowledgementsTips.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_AcknowledgementsTips_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=script&lang=js":
/*!******************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=script&lang=js ***!
  \******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ForbidPenetrationTipCheck_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ForbidPenetrationTipCheck.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ForbidPenetrationTipCheck_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=script&lang=js":
/*!**************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=script&lang=js ***!
  \**************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_FormTopV2_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FormTopV2.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_FormTopV2_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=script&lang=js":
/*!***************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=script&lang=js ***!
  \***************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_SpecSelect_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./SpecSelect.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_SpecSelect_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=script&lang=js":
/*!********************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=script&lang=js ***!
  \********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskLimitDialog_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./TaskLimitDialog.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskLimitDialog_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=script&lang=js":
/*!*************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=script&lang=js ***!
  \*************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskName_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./TaskName.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskName_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=style&index=0&id=aeeb1646&scoped=true&lang=less":
/*!***********************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=style&index=0&id=aeeb1646&scoped=true&lang=less ***!
  \***********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_AcknowledgementsTips_vue_vue_type_style_index_0_id_aeeb1646_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../node_modules/less-loader/dist/cjs.js!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./AcknowledgementsTips.vue?vue&type=style&index=0&id=aeeb1646&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=style&index=0&id=aeeb1646&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=0&id=0f089c42&scoped=true&lang=less":
/*!***************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=0&id=0f089c42&scoped=true&lang=less ***!
  \***************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ForbidPenetrationTipCheck_vue_vue_type_style_index_0_id_0f089c42_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../node_modules/less-loader/dist/cjs.js!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ForbidPenetrationTipCheck.vue?vue&type=style&index=0&id=0f089c42&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=0&id=0f089c42&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=1&id=0f089c42&lang=less":
/*!***************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=1&id=0f089c42&lang=less ***!
  \***************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ForbidPenetrationTipCheck_vue_vue_type_style_index_1_id_0f089c42_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../node_modules/less-loader/dist/cjs.js!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ForbidPenetrationTipCheck.vue?vue&type=style&index=1&id=0f089c42&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=style&index=1&id=0f089c42&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=style&index=0&id=67f642aa&scoped=true&lang=less":
/*!***********************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=style&index=0&id=67f642aa&scoped=true&lang=less ***!
  \***********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_FormTopV2_vue_vue_type_style_index_0_id_67f642aa_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../node_modules/less-loader/dist/cjs.js!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FormTopV2.vue?vue&type=style&index=0&id=67f642aa&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=style&index=0&id=67f642aa&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=style&index=0&id=0e79f3ae&scoped=true&lang=less":
/*!************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=style&index=0&id=0e79f3ae&scoped=true&lang=less ***!
  \************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_SpecSelect_vue_vue_type_style_index_0_id_0e79f3ae_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../node_modules/less-loader/dist/cjs.js!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./SpecSelect.vue?vue&type=style&index=0&id=0e79f3ae&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=style&index=0&id=0e79f3ae&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=style&index=0&id=6151937c&scoped=true&lang=less":
/*!*****************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=style&index=0&id=6151937c&scoped=true&lang=less ***!
  \*****************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskLimitDialog_vue_vue_type_style_index_0_id_6151937c_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../node_modules/less-loader/dist/cjs.js!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./TaskLimitDialog.vue?vue&type=style&index=0&id=6151937c&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=style&index=0&id=6151937c&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=style&index=0&id=09a23d7c&scoped=true&lang=less":
/*!**********************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=style&index=0&id=09a23d7c&scoped=true&lang=less ***!
  \**********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskName_vue_vue_type_style_index_0_id_09a23d7c_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../node_modules/less-loader/dist/cjs.js!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./TaskName.vue?vue&type=style&index=0&id=09a23d7c&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=style&index=0&id=09a23d7c&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=template&id=aeeb1646&scoped=true":
/*!********************************************************************************************************!*\
  !*** ./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=template&id=aeeb1646&scoped=true ***!
  \********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_AcknowledgementsTips_vue_vue_type_template_id_aeeb1646_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_AcknowledgementsTips_vue_vue_type_template_id_aeeb1646_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_AcknowledgementsTips_vue_vue_type_template_id_aeeb1646_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./AcknowledgementsTips.vue?vue&type=template&id=aeeb1646&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=template&id=aeeb1646&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=template&id=0f089c42&scoped=true":
/*!************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=template&id=0f089c42&scoped=true ***!
  \************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ForbidPenetrationTipCheck_vue_vue_type_template_id_0f089c42_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ForbidPenetrationTipCheck_vue_vue_type_template_id_0f089c42_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ForbidPenetrationTipCheck_vue_vue_type_template_id_0f089c42_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ForbidPenetrationTipCheck.vue?vue&type=template&id=0f089c42&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=template&id=0f089c42&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=template&id=67f642aa&scoped=true":
/*!********************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=template&id=67f642aa&scoped=true ***!
  \********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FormTopV2_vue_vue_type_template_id_67f642aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FormTopV2_vue_vue_type_template_id_67f642aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FormTopV2_vue_vue_type_template_id_67f642aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FormTopV2.vue?vue&type=template&id=67f642aa&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=template&id=67f642aa&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=template&id=0e79f3ae&scoped=true":
/*!*********************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=template&id=0e79f3ae&scoped=true ***!
  \*********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_SpecSelect_vue_vue_type_template_id_0e79f3ae_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_SpecSelect_vue_vue_type_template_id_0e79f3ae_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_SpecSelect_vue_vue_type_template_id_0e79f3ae_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./SpecSelect.vue?vue&type=template&id=0e79f3ae&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=template&id=0e79f3ae&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=template&id=6151937c&scoped=true":
/*!**************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=template&id=6151937c&scoped=true ***!
  \**************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskLimitDialog_vue_vue_type_template_id_6151937c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskLimitDialog_vue_vue_type_template_id_6151937c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskLimitDialog_vue_vue_type_template_id_6151937c_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./TaskLimitDialog.vue?vue&type=template&id=6151937c&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=template&id=6151937c&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=template&id=09a23d7c&scoped=true":
/*!*******************************************************************************************************!*\
  !*** ./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=template&id=09a23d7c&scoped=true ***!
  \*******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskName_vue_vue_type_template_id_09a23d7c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskName_vue_vue_type_template_id_09a23d7c_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_TaskName_vue_vue_type_template_id_09a23d7c_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./TaskName.vue?vue&type=template&id=09a23d7c&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=template&id=09a23d7c&scoped=true");


/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=template&id=aeeb1646&scoped=true":
/*!***********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/AcknowledgementsTips.vue?vue&type=template&id=aeeb1646&scoped=true ***!
  \***********************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "ack-tips-c" }, [
    _c("div", { staticClass: "icon-c" }, [
      _c(
        "svg",
        {
          staticClass:
            "styles__StyledSVGIconPathComponent-sc-4n1c4t-0 kmDyzV svg-icon-path-icon",
          attrs: {
            xmlns: "http://www.w3.org/2000/svg",
            viewBox: "0 0 1072 1024",
            width: "60",
            height: "60"
          }
        },
        [
          _c("defs"),
          _vm._v(" "),
          _c("g", [
            _c("path", {
              attrs: {
                d:
                  "M667.989333 247.466667h-31.719619c-7.460571 0-13.06819-3.705905-13.06819-11.166477 0-7.43619 5.607619-11.166476 13.06819-11.166476h31.719619V193.487238c0-7.43619 5.607619-13.019429 13.068191-13.019428 7.460571 0 13.06819 5.583238 13.06819 13.019428v31.646476h27.989334c7.460571 0 13.06819 3.730286 13.06819 11.166476 0 7.460571-5.607619 11.166476-13.06819 11.166477H694.125714v27.940571c0 7.43619-5.607619 13.019429-13.06819 13.019429a12.726857 12.726857 0 0 1-13.068191-13.019429V247.466667z",
                fill: "#6DD400",
                "p-id": "6404"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M821.248 50.273524V5.583238c0-1.852952 3.730286-5.583238 9.337905-5.583238s9.337905 3.730286 9.337905 5.607619v42.788571c0 3.730286-3.730286 7.460571-9.337905 7.460572s-9.337905-3.730286-9.337905-5.607619z",
                fill: "#FA6400",
                "p-id": "6405"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M852.992 37.229714h-42.934857c-3.730286 0-7.460571-3.730286-7.460572-9.313524s3.730286-9.289143 5.607619-9.289142h42.910477c3.730286 0 7.460571 3.705905 7.460571 9.289142 0 5.607619-3.730286 9.313524-5.607619 9.313524z",
                fill: "#FA6400",
                "p-id": "6406"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M953.782857 567.905524l-16.822857-16.749714c-5.607619-5.583238-5.607619-13.019429 0-18.602667a14.677333 14.677333 0 0 1 20.528762 0l16.822857 16.749714 16.774095-16.749714a14.677333 14.677333 0 0 1 20.528762 0c5.607619 5.583238 5.607619 13.019429 0 18.602667l-16.798476 16.749714 16.822857 16.774095c5.583238 5.583238 5.583238 13.019429 0 18.602667-3.754667 1.852952-5.607619 3.730286-9.362286 3.730285a14.262857 14.262857 0 0 1-9.313523-3.730285l-16.822858-16.749715-16.774095 16.749715a14.287238 14.287238 0 0 1-9.337905 3.730285 14.262857 14.262857 0 0 1-9.337904-3.730285c-5.607619-5.607619-5.607619-13.019429 0-18.627048l13.06819-16.749714zM174.445714 365.007238l-31.695238-29.769143c-1.877333-3.730286-1.877333-7.460571 0-9.313524 3.705905-3.730286 7.43619-3.730286 9.313524 0l31.719619 31.646477c1.877333 1.852952 1.877333 7.43619 0 9.313523-1.852952 1.852952-5.607619 1.852952-9.337905-1.877333z",
                fill: "#F7B500",
                "p-id": "6407"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M183.832381 335.238095L152.064 364.982857c-1.877333 1.877333-7.460571 1.877333-9.337905 0-1.877333-1.852952-1.877333-5.583238 0-9.313524l31.719619-31.622095c3.730286-1.877333 7.484952-1.877333 9.337905 1.852952 3.730286 1.852952 3.730286 5.583238 0 9.313524zM348.208762 221.549714c-26.136381 0-46.665143-20.48-46.665143-44.665904 0-26.087619 20.528762-46.567619 46.665143-46.56762s46.665143 20.48 46.665143 44.690286a46.153143 46.153143 0 0 1-46.665143 46.543238z m0-65.145904c-11.215238 0-18.67581 7.43619-18.67581 18.602666s7.484952 18.627048 18.67581 18.627048c11.215238 0 18.651429-7.460571 18.651428-18.627048s-7.460571-18.602667-18.651428-18.602666z",
                fill: "#F7B500",
                "p-id": "6408"
              }
            }),
            _vm._v(" "),
            _c("path", {
              staticClass: "selected-ilysg2xf064v02p-UIxb7NtTX3AVjZ",
              attrs: {
                d:
                  "M763.928381 982.528c0 22.113524-96.743619 40.106667-216.064 40.106667-119.344762 0-216.161524-17.92-216.161524-40.106667 0-22.113524 96.743619-40.106667 216.161524-40.106667 119.320381 0 216.039619 17.993143 216.039619 40.106667z",
                fill: "#e6e6e6",
                "p-id": "6409",
                "data-spm-anchor-id": "a313x.search_index.0.i2.399d3a81HlCAyy"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M60.318476 591.530667a15.11619 15.11619 0 0 0 15.067429-15.043048 15.018667 15.018667 0 0 0-15.067429-15.043048H15.067429A15.11619 15.11619 0 0 1 0 546.401524a15.018667 15.018667 0 0 1 15.067429-15.043048h150.79619a15.11619 15.11619 0 0 1 15.067429 15.043048 15.018667 15.018667 0 0 1-15.067429 15.043047H135.704381a15.11619 15.11619 0 0 0-15.09181 15.043048 15.018667 15.018667 0 0 0 15.09181 15.043048h60.294095a15.11619 15.11619 0 0 1 15.09181 15.043047 15.018667 15.018667 0 0 1-15.09181 15.018667H45.226667a15.11619 15.11619 0 0 1-15.067429-15.018667 15.018667 15.018667 0 0 1 15.067429-15.043047h15.091809zM888.07619 364.446476a18.480762 18.480762 0 0 0 18.456381-18.432 18.407619 18.407619 0 0 0-18.480761-18.407619h-55.393524a18.505143 18.505143 0 0 1-18.480762-18.432 18.407619 18.407619 0 0 1 18.480762-18.432h184.685714a18.505143 18.505143 0 0 1 18.480762 18.432 18.407619 18.407619 0 0 1-18.480762 18.432h-36.937143a18.480762 18.480762 0 0 0-18.456381 18.407619 18.407619 18.407619 0 0 0 18.456381 18.432h73.874286a18.480762 18.480762 0 0 1 18.480762 18.407619 18.407619 18.407619 0 0 1-18.480762 18.432h-184.685714a18.505143 18.505143 0 0 1-18.480762-18.432 18.407619 18.407619 0 0 1 18.480762-18.407619h18.480761zM774.656 469.528381c1.219048 1.950476 2.413714 3.82781 3.584 5.753905-1.170286-1.877333-2.438095-3.803429-3.608381-5.753905zM462.531048 859.062857c-0.950857-0.341333-1.877333-0.707048-2.779429-0.975238l2.779429 0.975238z m7.94819 2.779429a72.94781 72.94781 0 0 1-3.291428-1.121524c1.121524 0.341333 2.194286 0.75581 3.291428 1.121524z m-4.291048-1.462857c-0.902095-0.316952-1.77981-0.609524-2.608761-0.975239 0.902095 0.365714 1.77981 0.658286 2.608761 0.975239z m-15.018666-5.729524z m306.712381-408.210286c1.414095 1.77981 2.82819 3.535238 4.242285 5.36381l-4.242285-5.36381zM459.093333 857.770667l-3.535238-1.340953c1.194667 0.487619 2.364952 0.926476 3.535238 1.340953z m-3.754666-1.462857c-1.072762-0.414476-2.121143-0.877714-3.242667-1.292191 1.048381 0.487619 2.121143 0.877714 3.242667 1.292191z m34.742857 11.044571c-1.219048-0.292571-2.486857-0.585143-3.657143-0.950857 1.243429 0.365714 2.438095 0.658286 3.657143 0.950857z m-42.422857-14.214095l-2.80381-1.219048c0.926476 0.463238 1.877333 0.853333 2.828191 1.219048z m353.816381-328.899048c0.341333 1.048381 0.75581 2.096762 1.121523 3.169524a62.244571 62.244571 0 0 1-1.121523-3.169524z m-1.243429-3.364571c0.414476 1.048381 0.75581 2.194286 1.170286 3.242666-0.341333-1.072762-0.75581-2.121143-1.170286-3.242666zM474.38019 863.061333c-1.219048-0.341333-2.413714-0.75581-3.584-1.121523l3.584 1.121523z m7.314286 2.121143z m308.614095-367.640381c0.487619 0.999619 0.950857 2.072381 1.462858 3.072-0.463238-1.072762-0.975238-2.072381-1.462858-3.072zM493.763048 868.230095c-0.877714-0.24381-1.828571-0.414476-2.681905-0.633905 0.853333 0.24381 1.755429 0.414476 2.681905 0.633905z m-135.996953-78.994285z m58.197334 47.616a60.952381 60.952381 0 0 1-3.291429-2.048c1.048381 0.731429 2.194286 1.365333 3.291429 2.048z m-75.922286-69.12c-0.707048-0.950857-1.340952-1.901714-2.048-2.828191 0.633905 0.926476 1.340952 1.877333 2.048 2.80381z m72.508952 67.023238l-3.072-1.950477 3.072 1.950477z m9.898667 5.924571l-2.608762-1.462857c0.902095 0.487619 1.77981 0.926476 2.608762 1.462857zM349.817905 780.190476l-2.218667-2.633143c0.731429 0.877714 1.511619 1.755429 2.218667 2.633143z m46.006857 43.25181c-0.828952-0.658286-1.706667-1.219048-2.535619-1.877334 0.877714 0.633905 1.706667 1.219048 2.535619 1.877334z m-3.072-2.29181z m12.921905 9.216a87.064381 87.064381 0 0 1-2.950096-1.999238c0.975238 0.658286 1.999238 1.29219 2.925715 1.999238z m-25.234286-19.139047z m56.149333 36.888381c-0.902095-0.414476-1.706667-0.877714-2.608762-1.292191 0.828952 0.414476 1.706667 0.877714 2.608762 1.292191z m7.655619 3.584l-3.486476-1.584762 3.486476 1.584762z m-14.384762-6.997334c-1.121524-0.585143-2.29181-1.170286-3.413333-1.828571l3.413333 1.828571z m-3.535238-1.877333l-3.120762-1.706667 3.120762 1.706667z m6.826667 3.584c-0.926476-0.487619-1.877333-0.999619-2.876952-1.462857 0.950857 0.463238 1.877333 0.975238 2.901333 1.462857z m68.022857-28.038095c145.724952 0 263.875048-117.833143 263.875048-263.168a261.607619 261.607619 0 0 0-44.714667-146.627048l2.608762 2.194286a263.143619 263.143619 0 0 0-170.520381-62.342096c-145.700571 0-263.850667 117.833143-263.850667 263.168 0 57.10019 18.261333 109.909333 49.176381 153.112381-1.536-2.145524-2.998857-4.315429-4.486095-6.41219a263.43619 263.43619 0 0 0 167.936 60.099048z m162.084572-334.604191c26.209524 0 47.469714 21.211429 47.469714 47.34781a47.420952 47.420952 0 0 1-47.469714 47.347809 47.420952 47.420952 0 0 1-47.469715-47.347809c0-26.136381 21.26019-47.34781 47.469715-47.34781z m7.899428 136.874667c11.605333 0 20.041143 9.99619 18.505143 21.040762a137.508571 137.508571 0 0 1-59.855238 92.525714l-0.048762 0.048762a28.281905 28.281905 0 0 1-1.755429 1.121524l-1.950476 1.219047c-18.773333 11.702857-40.521143 18.919619-63.853714 20.504381h-0.121905l-2.876952 0.195048h-0.341334c-0.902095 0.048762-1.828571 0.048762-2.730666 0.121905h-0.463238l-3.169524 0.048762c-1.072762 0-2.121143 0-3.169524-0.073143h-0.560762a41.935238 41.935238 0 0 1-2.681905-0.097524h-0.438857l-2.82819-0.195048h-0.097524a137.752381 137.752381 0 0 1-63.804952-20.48c-0.048762-0.073143-0.097524-0.073143-0.170667-0.121904-0.585143-0.365714-1.170286-0.78019-1.828571-1.121524a137.679238 137.679238 0 0 1-59.416381-82.651429h-0.121905a78.409143 78.409143 0 0 1-2.121143-10.971428v-0.146286a18.383238 18.383238 0 0 1 18.505143-21.016381h237.397333v0.048762z m-229.522286-136.874667c26.209524 0 47.469714 21.211429 47.469715 47.34781a47.420952 47.420952 0 0 1-47.469715 47.347809 47.420952 47.420952 0 0 1-47.469714-47.347809c0-26.136381 21.211429-47.34781 47.469714-47.34781z m84.821334 389.729524l-2.706286-0.316952c0.877714 0.121905 1.755429 0.24381 2.681905 0.316952z m287.695238-295.619048c0.195048 1.170286 0.316952 2.340571 0.414476 3.535238-0.097524-1.194667-0.292571-2.364952-0.414476-3.535238z m0.414476 3.584l0.365714 3.462096c-0.073143-1.170286-0.24381-2.340571-0.365714-3.462096zM530.456381 873.862095c-0.877714-0.048762-1.755429-0.170667-2.657524-0.243809l2.681905 0.243809z m-7.899429-0.731428a71.119238 71.119238 0 0 1-3.754666-0.487619c1.219048 0.170667 2.535619 0.292571 3.754666 0.487619z m-3.998476-0.560762l-3.535238-0.463238c1.194667 0.170667 2.364952 0.365714 3.535238 0.487619z m28.745143 2.121143l-2.706286-0.048762 2.706286 0.048762z m-4.071619-0.121905c-0.877714-0.048762-1.828571-0.048762-2.681905-0.121905 0.853333 0.073143 1.80419 0.121905 2.681905 0.121905z m-140.580571-46.32381c-1.048381-0.78019-2.194286-1.487238-3.242667-2.243047 1.072762 0.75581 2.121143 1.511619 3.242667 2.218666z m-41.935239-35.84z m174.640762 81.822477c-1.243429-0.048762-2.486857-0.170667-3.705904-0.24381 1.219048 0.073143 2.438095 0.121905 3.705904 0.24381z m3.998477 0.243809c-1.29219-0.073143-2.584381-0.121905-3.82781-0.243809 1.219048 0.048762 2.535619 0.121905 3.82781 0.243809z m270.214095-322.291809c0.24381 1.121524 0.536381 2.243048 0.755809 3.364571-0.292571-1.121524-0.512-2.243048-0.755809-3.364571zM506.148571 870.692571l-3.730285-0.707047c1.243429 0.316952 2.486857 0.487619 3.705904 0.731428z m3.754667 0.658286a43.398095 43.398095 0 0 1-2.706286-0.487619c0.950857 0.195048 1.828571 0.365714 2.706286 0.487619z",
                fill: "#FFD73A",
                "p-id": "6410"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M769.462857 761.319619c-1.828571 2.584381-3.657143 5.168762-5.534476 7.704381 1.877333-2.535619 3.705905-5.071238 5.534476-7.704381z m-7.484952 10.24c-1.29219 1.633524-2.584381 3.31581-3.949715 4.973714 1.365333-1.657905 2.657524-3.291429 3.949715-4.998095z m-56.783238 54.613333c-1.706667 1.219048-3.462095 2.413714-5.241905 3.584 1.77981-1.121524 3.486476-2.340571 5.241905-3.584z m7.655619-5.632z m19.334095-16.286476z m-9.508571 8.411429c-1.560381 1.340952-3.218286 2.681905-4.876191 3.998476 1.657905-1.365333 3.31581-2.657524 4.876191-3.998476z m29.305904-28.842667z m21.016381-27.721143l-3.535238 5.217524 3.535238-5.217524z m11.459048-19.114666z m30.47619-98.499048z m-2.779428 19.675429z m4.144762-43.178667c0 2.29181-0.073143 4.534857-0.195048 6.753524 0.121905-2.218667 0.121905-4.534857 0.170667-6.753524z m-2.56 33.401905c-0.292571 2.194286-0.633905 4.388571-0.975238 6.534095 0.341333-2.169905 0.682667-4.33981 0.975238-6.509714z m-33.962667 96.938666c-1.072762 1.828571-2.194286 3.657143-3.31581 5.412572 1.072762-1.828571 2.194286-3.584 3.31581-5.412572z m-79.823238 84.48z m94.354286-112.835047l-2.633143 5.802666 2.633143-5.802666z m-116.077715 126.000762l-5.656381 2.998857c1.877333-0.999619 3.779048-1.950476 5.656381-2.998857z m-17.188571 8.533333z",
                fill: "#FEC43C",
                "p-id": "6411"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M816.274286 607.939048v-1.170286l-0.048762-2.438095c0-0.438857 0-0.926476-0.048762-1.389715 0-0.75581-0.073143-1.462857-0.073143-2.243047l-0.048762-1.462857c-0.073143-0.707048-0.073143-1.414095-0.121905-2.121143l-0.048762-1.462857c-0.073143-0.707048-0.121905-1.414095-0.121904-2.121143-0.073143-0.463238-0.073143-0.975238-0.121905-1.462857l-0.170667-2.169905c-0.073143-0.487619-0.073143-0.926476-0.121904-1.414095a33.401905 33.401905 0 0 0-0.24381-2.413715c-0.048762-0.390095-0.048762-0.731429-0.121905-1.170285l-0.341333-3.462096v-0.048762l-0.414476-3.535238-0.170667-1.121523a32.402286 32.402286 0 0 1-0.316952-2.340572l-0.170667-1.340952a137.142857 137.142857 0 0 1-0.341333-2.194286c-0.073143-0.463238-0.195048-0.926476-0.24381-1.389714a146.090667 146.090667 0 0 0-0.341333-2.072381c-0.073143-0.463238-0.195048-0.926476-0.24381-1.414095l-0.341333-2.048c-0.073143-0.487619-0.195048-0.926476-0.24381-1.414096l-0.414476-2.121143c-0.048762-0.463238-0.170667-0.877714-0.243809-1.340952-0.170667-0.75581-0.292571-1.536-0.487619-2.291809l-0.219429-1.121524a66.243048 66.243048 0 0 0-0.755809-3.340191v-0.048762a69.095619 69.095619 0 0 0-1.121524-4.461714c-0.170667-0.78019-0.414476-1.536-0.585143-2.31619l-0.365714-1.292191c-0.170667-0.682667-0.341333-1.389714-0.585143-2.048l-0.341334-1.340952c-0.195048-0.658286-0.365714-1.365333-0.609523-1.999238a12.434286 12.434286 0 0 0-0.390096-1.365334l-0.609523-1.974857c-0.097524-0.487619-0.292571-0.902095-0.414477-1.365333l-0.633904-2.048a9.630476 9.630476 0 0 0-0.414477-1.219048l-0.755809-2.243047-0.365714-1.072762c-0.341333-1.048381-0.707048-2.121143-1.121524-3.169524v-0.048762c-0.414476-1.072762-0.75581-2.194286-1.170286-3.242667l-0.365714-0.975238a54.442667 54.442667 0 0 1-0.804572-2.243047l-0.487619-1.194667-0.755809-1.974857a30.47619 30.47619 0 0 1-0.536381-1.292191c-0.24381-0.658286-0.536381-1.29219-0.75581-1.950476l-0.536381-1.29219-0.828952-1.950477-0.536381-1.219047-0.877714-1.950476-0.536381-1.170286c-0.292571-0.707048-0.633905-1.414095-0.975238-2.121143l-0.487619-0.975238c-0.487619-0.999619-0.950857-2.072381-1.462857-3.072v-0.048762c-0.487619-1.072762-0.999619-2.048-1.536-3.047619a6.460952 6.460952 0 0 0-0.48762-0.902095c-0.341333-0.707048-0.707048-1.462857-1.097142-2.169905a21.577143 21.577143 0 0 0-0.585143-1.121524c-0.365714-0.633905-0.658286-1.26781-1.024-1.950476-0.24381-0.390095-0.414476-0.731429-0.633905-1.145905-0.365714-0.585143-0.658286-1.219048-0.999619-1.828571l-0.658286-1.170286a28.111238 28.111238 0 0 1-0.975238-1.828571l-0.658286-1.170286c-0.365714-0.633905-0.731429-1.219048-1.072761-1.877333l-0.658286-1.072762-1.219048-2.048-0.536381-0.877715c-1.170286-1.950476-2.413714-3.876571-3.584-5.753904l-0.536381-0.75581c-0.487619-0.731429-0.877714-1.414095-1.365333-2.072381-0.24381-0.341333-0.463238-0.633905-0.633905-0.975238l-1.219047-1.828571-0.731429-1.072762c-0.414476-0.585143-0.828952-1.170286-1.170286-1.755429l-0.78019-1.048381a43.008 43.008 0 0 0-1.219048-1.706666l-0.78019-1.072762-1.292191-1.755429-0.707047-0.975238a42.398476 42.398476 0 0 0-1.462858-1.950476l-0.585142-0.780191a121.295238 121.295238 0 0 0-4.778667-5.973333l-1.584762-1.950476a10.727619 10.727619 0 0 0-0.78019-0.877714l-1.414096-1.706667c-0.292571-0.292571-0.512-0.633905-0.828952-0.950857L751.299048 438.857143l-0.902096-0.975238-1.414095-1.609143a11.361524 11.361524 0 0 1-0.828952-0.926476l-1.462857-1.657905-0.828953-0.877714-1.633524-1.755429c-0.24381-0.24381-0.414476-0.487619-0.658285-0.707048-1.584762-1.633524-3.169524-3.291429-4.778667-4.87619l-0.585143-0.585143c-0.585143-0.585143-1.219048-1.170286-1.828571-1.828571l-0.828953-0.828953a20.382476 20.382476 0 0 0-1.633523-1.511619l-0.877715-0.828952-1.609143-1.462857-0.926476-0.877715a34.864762 34.864762 0 0 1-1.609143-1.462857 10.971429 10.971429 0 0 0-0.926476-0.828952l-1.657905-1.462857-0.877714-0.780191-1.828571-1.584762-0.707048-0.633904-2.584381-2.194286a261.144381 261.144381 0 0 1 44.690286 146.627048c0 145.359238-118.150095 263.192381-263.850667 263.192381-63.780571 0-122.270476-22.552381-167.911619-60.14781 1.462857 2.169905 2.925714 4.266667 4.486095 6.38781l0.170667 0.292571 2.048 2.82819 0.292571 0.341334c2.31619 3.120762 4.656762 6.119619 7.070477 9.118476 0.073143 0.121905 0.195048 0.170667 0.243809 0.292571 0.707048 0.877714 1.462857 1.755429 2.243048 2.633143l0.292571 0.365715c2.462476 2.925714 4.998095 5.802667 7.606857 8.630857 0.097524 0.121905 0.170667 0.24381 0.292572 0.292571l2.291809 2.462476 0.341334 0.365715c2.657524 2.755048 5.315048 5.510095 8.070095 8.167619l0.365714 0.341333 2.413714 2.291809 0.414477 0.341334c2.755048 2.584381 5.607619 5.168762 8.533333 7.655619l0.414476 0.341333 2.462476 2.121143c0.195048 0.121905 0.292571 0.24381 0.487619 0.414476a525.360762 525.360762 0 0 0 9.411048 7.509334c0.828952 0.658286 1.706667 1.29219 2.535619 1.877333l0.536381 0.414476 2.998857 2.194286h0.048762c1.072762 0.731429 2.121143 1.511619 3.242667 2.194286 0.073143 0 0.073143 0.073143 0.121905 0.073142 0.999619 0.658286 1.950476 1.365333 2.925714 1.999238 0.24381 0.121905 0.438857 0.292571 0.658286 0.414477 0.828952 0.585143 1.706667 1.121524 2.608761 1.706666 0.24381 0.121905 0.414476 0.292571 0.633905 0.414477l3.072 1.950476 0.048762 0.048762c1.121524 0.707048 2.243048 1.340952 3.31581 2.048l0.292571 0.170666c0.926476 0.585143 1.950476 1.170286 2.876952 1.706667l0.731429 0.414476 2.56 1.462857 0.731429 0.414476c1.048381 0.585143 2.048 1.170286 3.096381 1.706667 0.073143 0 0.073143 0.048762 0.121904 0.048762 1.121524 0.609524 2.243048 1.243429 3.413334 1.828571l0.414476 0.24381 2.901333 1.462857c0.292571 0.121905 0.512 0.292571 0.804572 0.414476 0.902095 0.414476 1.706667 0.877714 2.608761 1.292191 0.292571 0.121905 0.512 0.292571 0.828953 0.414476 1.048381 0.512 2.121143 0.975238 3.169524 1.511619l0.170666 0.073143c1.194667 0.512 2.29181 1.048381 3.486477 1.584762 0.170667 0.121905 0.414476 0.170667 0.585142 0.243809l2.828191 1.219048 0.926476 0.414476c0.902095 0.341333 1.706667 0.707048 2.608762 1.121524l0.877714 0.341333c1.072762 0.414476 2.121143 0.877714 3.242667 1.292191 0.048762 0.073143 0.170667 0.073143 0.243809 0.121905 1.170286 0.487619 2.340571 0.950857 3.535239 1.340952l0.707047 0.316952c0.926476 0.341333 1.80419 0.682667 2.755048 0.975238l1.072762 0.365715c0.877714 0.292571 1.706667 0.585143 2.584381 0.926476l0.975238 0.365714c1.072762 0.341333 2.194286 0.707048 3.315809 1.121524 0.121905 0 0.170667 0.048762 0.292572 0.121905 1.194667 0.390095 2.438095 0.731429 3.608381 1.097143l0.804571 0.243809 2.779429 0.828952 1.121523 0.341334c0.877714 0.24381 1.755429 0.487619 2.584381 0.755809l1.072762 0.316953 3.340191 0.877714 0.292571 0.048762c1.243429 0.292571 2.486857 0.633905 3.657143 0.950857 0.292571 0.048762 0.658286 0.170667 0.950857 0.24381 0.877714 0.219429 1.828571 0.390095 2.681905 0.633904 0.438857 0.121905 0.853333 0.170667 1.219048 0.292572l2.633142 0.585143c0.341333 0.048762 0.75581 0.170667 1.121524 0.243809l3.413334 0.682667c0.121905 0 0.170667 0.073143 0.292571 0.073143l3.705905 0.707047 1.072762 0.170667c0.877714 0.170667 1.828571 0.292571 2.681904 0.487619l1.316572 0.170667c0.877714 0.097524 1.755429 0.26819 2.657524 0.390095l1.170285 0.195048a72.996571 72.996571 0 0 0 3.779048 0.512c1.219048 0.195048 2.535619 0.292571 3.754667 0.487619l1.121523 0.097523c0.877714 0.121905 1.828571 0.195048 2.706286 0.292572 0.414476 0.073143 0.877714 0.073143 1.292191 0.121905l2.681904 0.243809c0.390095 0.048762 0.804571 0.048762 1.145905 0.121905 1.243429 0.121905 2.486857 0.170667 3.730286 0.243809h0.170667c1.29219 0.048762 2.584381 0.170667 3.827809 0.219429l1.170286 0.048762c0.877714 0.073143 1.828571 0.073143 2.706286 0.121905 0.487619 0 0.902095 0.073143 1.365333 0.073142l2.681905 0.048762h1.267809l3.900953 0.048762a263.265524 263.265524 0 0 0 142.214095-41.447619l5.290666-3.486476c1.77981-1.170286 3.486476-2.413714 5.241905-3.584 0.902095-0.585143 1.706667-1.219048 2.608762-1.828571 1.706667-1.219048 3.413333-2.511238 5.046857-3.754667 1.657905-1.219048 3.364571-2.584381 5.022476-3.925333 1.633524-1.29219 3.291429-2.657524 4.876191-3.998477l2.438095-2.048a247.734857 247.734857 0 0 0 18.310095-17.456761 231.789714 231.789714 0 0 0 14.628572-16.62781c1.26781-1.633524 2.608762-3.291429 3.900952-4.998095 0.658286-0.804571 1.316571-1.706667 1.950476-2.511238 1.901714-2.535619 3.779048-5.12 5.534477-7.704381 1.194667-1.755429 2.438095-3.462095 3.535238-5.217524a166.034286 166.034286 0 0 0 3.413333-5.36381c1.121524-1.80419 2.243048-3.584 3.31581-5.38819a297.593905 297.593905 0 0 0 9.118476-16.749714 174.32381 174.32381 0 0 0 2.779428-5.753905l2.633143-5.827048a256.585143 256.585143 0 0 0 17.895619-58.758095c0.195048-1.024 0.365714-2.145524 0.609524-3.218286a240.688762 240.688762 0 0 0 1.414095-9.801143 262.022095 262.022095 0 0 0 1.950476-23.381333l0.170667-6.753524v-3.413333c0.24381-1.511619 0.24381-2.681905 0.24381-3.925333z",
                fill: "#FEC43C",
                "p-id": "6412"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M394.166857 531.163429a47.250286 47.250286 0 0 0 29.305905 43.739428 47.591619 47.591619 0 0 0 62.025143-25.624381 47.225905 47.225905 0 0 0-10.288762-51.614476 47.542857 47.542857 0 0 0-81.042286 33.499429zM615.814095 531.163429a47.201524 47.201524 0 0 0 29.281524 43.739428 47.567238 47.567238 0 0 0 62.025143-25.624381 47.201524 47.201524 0 0 0-25.673143-61.878857 47.567238 47.567238 0 0 0-62.025143 25.624381 47.225905 47.225905 0 0 0-3.608381 18.139429zM433.688381 620.617143z m255.439238 21.040762a81.237333 81.237333 0 0 1-2.121143 11.044571H417.499429a137.679238 137.679238 0 0 0 59.416381 82.651429 133.753905 133.753905 0 0 1-10.99581-7.92381c2.121143-24.210286 22.674286-43.178667 47.469714-43.178666 16.359619 0 30.622476 7.875048 39.033905 19.968 8.94781-12.092952 23.210667-19.992381 39.058286-19.992381a47.542857 47.542857 0 0 1 47.469714 43.203047 137.362286 137.362286 0 0 0 50.663619-85.820952 18.383238 18.383238 0 0 0-18.529524-21.040762c11.142095 0.048762 20.163048 10.044952 18.041905 21.089524z",
                fill: "#873A18",
                "p-id": "6413"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M417.401905 652.678095h0.097524a135.192381 135.192381 0 0 1-2.243048-10.971428c0.536381 3.705905 1.072762 7.338667 2.121143 10.971428z",
                fill: "#FFFFFF",
                "p-id": "6414"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M415.256381 641.657905v0.121905c0.585143 3.681524 1.365333 7.314286 2.243048 10.971428h269.507047c1.048381-3.705905 1.584762-7.338667 2.121143-11.044571 2.121143-11.044571-6.826667-21.016381-17.968762-21.016381h-237.470476a18.285714 18.285714 0 0 0-18.432 20.967619z",
                fill: "#FFFFFF",
                "p-id": "6415"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M549.254095 757.443048l3.169524 0.048762c1.072762 0 2.121143 0 3.169524-0.048762l-3.169524 0.048762c-1.048381-0.048762-2.121143-0.048762-3.169524-0.048762z m9.606095-0.121905c-0.902095 0.048762-1.828571 0.048762-2.730666 0.121905 0.902095-0.073143 1.77981-0.073143 2.730666-0.121905z m-10.093714 0.121905a41.935238 41.935238 0 0 1-2.681905-0.121905c0.853333 0.048762 1.755429 0.048762 2.681905 0.121905z m13.312-0.292572a2377.630476 2377.630476 0 0 1-2.876952 0.170667l2.876952-0.170667z m-19.358476 0z m83.334095-20.504381z",
                fill: "#F44444",
                "p-id": "6416"
              }
            }),
            _vm._v(" "),
            _c("path", {
              attrs: {
                d:
                  "M591.481905 684.30019c-15.847619 0-30.110476 7.875048-39.058286 19.992381a47.225905 47.225905 0 0 0-39.033905-19.992381 47.542857 47.542857 0 0 0-47.469714 43.178667 145.16419 145.16419 0 0 0 12.824381 9.045333l0.170667 0.121905a139.337143 139.337143 0 0 0 63.804952 20.48h0.097524l2.82819 0.195048h0.414476c0.902095 0.048762 1.77981 0.048762 2.706286 0.121905h0.536381l3.169524 0.048762c1.072762 0 2.145524 0 3.193905-0.048762h0.487619c0.877714 0 1.80419-0.073143 2.681905-0.121905h0.365714l2.876952-0.170667h0.121905a138.654476 138.654476 0 0 0 63.853714-20.504381l1.950476-1.219047c0.585143-0.365714 1.170286-0.78019 1.755429-1.121524l0.048762-0.073143c3.120762-2.121143 6.192762-4.33981 9.191619-6.680381a47.616 47.616 0 0 0-47.542857-43.25181z",
                fill: "#F44444",
                "p-id": "6417"
              }
            })
          ])
        ]
      )
    ]),
    _vm._v(" "),
    _c("div", { staticClass: "tips" }, [
      _c(
        "span",
        [
          _vm._v(_vm._s(_vm.$t("acknowledgementsTips.main1"))),
          _vm.usePop
            ? _c(
                "el-popover",
                { attrs: { placement: "top", width: "620", trigger: "hover" } },
                [
                  _c("a", { attrs: { slot: "reference" }, slot: "reference" }, [
                    _vm._v(_vm._s(_vm.$t("acknowledgementsTips.moreContent")))
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "pop-tips" }, [
                    _c("p", {
                      domProps: {
                        innerHTML: _vm._s(_vm.$t("acknowledgementsTips.p1"))
                      }
                    }),
                    _vm._v(" "),
                    _c("p", {
                      domProps: {
                        innerHTML: _vm._s(_vm.$t("acknowledgementsTips.p2"))
                      }
                    }),
                    _vm._v(" "),
                    _c("br"),
                    _vm._v(" "),
                    _c("p", {
                      domProps: {
                        innerHTML: _vm._s(_vm.$t("acknowledgementsTips.p3"))
                      }
                    }),
                    _vm._v(" "),
                    _c("p", {
                      domProps: {
                        innerHTML: _vm._s(_vm.$t("acknowledgementsTips.p4"))
                      }
                    })
                  ])
                ]
              )
            : _c(
                "a",
                {
                  attrs: {
                    target: "_blank",
                    href: "https://bbs.openi.org.cn/forums/7877"
                  }
                },
                [_vm._v(_vm._s(_vm.$t("acknowledgementsTips.moreContent")))]
              ),
          _vm._v(_vm._s(_vm.$t("acknowledgementsTips.main2")))
        ],
        1
      )
    ])
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=template&id=0f089c42&scoped=true":
/*!***************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/ForbidPenetrationTipCheck.vue?vue&type=template&id=0f089c42&scoped=true ***!
  \***************************************************************************************************************************************************************************************************************************************************************/
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
    { staticClass: "forbid-penetration-check-container" },
    [
      _c(
        "el-popover",
        {
          ref: "penaltyPopover",
          attrs: {
            placement: "right",
            trigger: "hover",
            "popper-class": "penalty-tip-popover"
          }
        },
        [
          _c("div", {
            staticClass: "penalty-tip-content",
            domProps: {
              innerHTML: _vm._s(_vm.$t("cloudbrainObj.penaltyTipContent"))
            }
          }),
          _vm._v(" "),
          _c(
            "el-checkbox",
            {
              attrs: { slot: "reference" },
              on: { change: _vm.handleChange },
              slot: "reference",
              model: {
                value: _vm.currentValue,
                callback: function($$v) {
                  _vm.currentValue = $$v
                },
                expression: "currentValue"
              }
            },
            [
              _c("span", { staticClass: "checkbox-label" }, [
                _vm._v(
                  _vm._s(_vm.$t("cloudbrainObj.forbidPenetrationStatement"))
                )
              ])
            ]
          )
        ],
        1
      )
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=template&id=67f642aa&scoped=true":
/*!***********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/FormTopV2.vue?vue&type=template&id=67f642aa&scoped=true ***!
  \***********************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", [
    _vm.configs.hideCluster !== true
      ? _c("div", { staticClass: "form-row form-row-cluster" }, [
          _c("div", { staticClass: "left-area" }, [
            _c("div", { staticClass: "title align-items-center" }, [
              _c("span", { staticClass: "required" }, [
                _vm._v(_vm._s(_vm.$t("cloudbrainObj.cluster")))
              ])
            ]),
            _vm._v(" "),
            _c("div", { staticClass: "content" }, [
              _c(
                "div",
                { staticClass: "list" },
                _vm._l(_vm.clusters, function(item) {
                  return _c(
                    "a",
                    {
                      key: item.key,
                      staticClass: "item",
                      class: item.key == _vm.cluster ? "focus" : "",
                      on: {
                        click: function($event) {
                          return _vm.changeCluster(item.key)
                        }
                      }
                    },
                    [
                      _c("i", { staticClass: "icon ri-global-line" }),
                      _vm._v(" "),
                      _c("span", [_vm._v(_vm._s(item.label))])
                    ]
                  )
                }),
                0
              )
            ])
          ])
        ])
      : _vm._e(),
    _vm._v(" "),
    _vm.configs.hideComputerResource !== true
      ? _c("div", { staticClass: "form-row form-row-computer-resource" }, [
          _c("div", { staticClass: "left-area" }, [
            _c("div", { staticClass: "title" }, [
              _c("span", { staticClass: "required" }, [
                _vm._v(_vm._s(_vm.$t("cloudbrainObj.computeResource")))
              ])
            ]),
            _vm._v(" "),
            _c("div", { staticClass: "content" }, [
              _c(
                "div",
                { staticClass: "list" },
                _vm._l(_vm.computerResouces, function(item) {
                  return _c(
                    "a",
                    {
                      key: item.key,
                      staticClass: "item",
                      class: item.key == _vm.computerResouce ? "focus" : "",
                      on: {
                        click: function($event) {
                          return _vm.changeComputerResouce(item.key)
                        }
                      }
                    },
                    [_c("span", [_vm._v(_vm._s(item.label))])]
                  )
                }),
                0
              )
            ])
          ])
        ])
      : _vm._e(),
    _vm._v(" "),
    _c("div", { staticClass: "form-row tips-c" }, [
      _c("div", { staticClass: "left-area" }, [
        _c("div", { staticClass: "title" }),
        _vm._v(" "),
        _c("div", { staticClass: "content" }, [
          _c("div", { staticClass: "tips tips-1" }, [
            _c("span", { staticClass: "wait-count-c" }, [
              _c("i", { staticClass: "ri-error-warning-line" }),
              _vm._v(" "),
              _c("span", [
                _vm._v(
                  "\n              " +
                    _vm._s(_vm.$t("cloudbrainObj.waitCountStart")) +
                    "\n              "
                ),
                _c("span", [_vm._v(_vm._s(_vm.queueNum))]),
                _vm._v(
                  "\n              " +
                    _vm._s(_vm.$t("cloudbrainObj.waitCountEnd")) +
                    "\n            "
                )
              ])
            ]),
            _vm._v(" "),
            _vm.configs.showMindTorchHelper
              ? _c(
                  "a",
                  {
                    staticClass: "mind-torch-helper",
                    attrs: {
                      href: "https://openi.pcl.ac.cn/OpenI/mindtorch_tutorial",
                      target: "_blank"
                    }
                  },
                  [
                    _c("i", { staticClass: "ri-arrow-right-line" }),
                    _vm._v(" "),
                    _c("span", [
                      _vm._v(_vm._s(_vm.$t("cloudbrainObj.mindTorchHelper")))
                    ])
                  ]
                )
              : _vm._e()
          ])
        ])
      ])
    ])
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=template&id=0e79f3ae&scoped=true":
/*!************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/SpecSelect.vue?vue&type=template&id=0e79f3ae&scoped=true ***!
  \************************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "form-row" }, [
    _c("div", { staticClass: "left-area" }, [
      _c("div", { staticClass: "title" }, [
        _c("span", { class: _vm.required ? "required" : "" }, [
          _vm._v(_vm._s(_vm.$t("cloudbrainObj.resourceSpec")))
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
              value: _vm.loading,
              expression: "loading"
            }
          ],
          staticClass: "content",
          class: _vm.errStatus ? "error" : ""
        },
        [
          !_vm.list.length
            ? _c("div", { staticClass: "spec-list-c" }, [
                _c(
                  "div",
                  {
                    staticClass: "spec-item-placeholder",
                    staticStyle: { color: "red" }
                  },
                  [
                    _vm._v(
                      "\n          " +
                        _vm._s(_vm.$t("specObj.no_use_resource")) +
                        "\n        "
                    )
                  ]
                )
              ])
            : _c(
                "div",
                { staticClass: "spec-info" },
                [
                  _c(
                    "el-select",
                    {
                      staticClass: "spec-sel field-input",
                      class: _vm.configs.showPoint ? "spec-show-point" : "",
                      attrs: {
                        placeholder: _vm.$t("cloudbrainObj.specPlaceholder")
                      },
                      on: { change: _vm.changeSpec },
                      model: {
                        value: _vm.spec,
                        callback: function($$v) {
                          _vm.spec = $$v
                        },
                        expression: "spec"
                      }
                    },
                    [
                      _c(
                        "div",
                        {
                          staticClass: "spec-sel-icon spec-op-icon",
                          attrs: { slot: "prefix" },
                          slot: "prefix"
                        },
                        [
                          _c(
                            "div",
                            { class: _vm.selIconType + "_icon _icon" },
                            [_vm._v(_vm._s(_vm.selIconType[0]))]
                          )
                        ]
                      ),
                      _vm._v(" "),
                      _vm._l(_vm.list, function(item) {
                        return _c(
                          "el-option",
                          {
                            key: item.id,
                            attrs: { label: item.specStr, value: item.id }
                          },
                          [
                            _c(
                              "span",
                              {
                                staticClass: "spec-op-icon",
                                staticStyle: { float: "left" }
                              },
                              [
                                _c(
                                  "div",
                                  { class: item.type + "_icon _icon" },
                                  [_vm._v(_vm._s(item.type[0]))]
                                )
                              ]
                            ),
                            _vm._v(" "),
                            _c(
                              "span",
                              {
                                staticClass: "spec-op-spec",
                                staticStyle: { float: "left" }
                              },
                              [_vm._v(_vm._s(item.specStr))]
                            ),
                            _vm._v(" "),
                            _vm.configs.showPoint
                              ? _c(
                                  "span",
                                  {
                                    staticClass: "spec-op-point",
                                    staticStyle: { float: "right" }
                                  },
                                  [_vm._v(_vm._s(item.pointStr))]
                                )
                              : _vm._e()
                          ]
                        )
                      }),
                      _vm._v(" "),
                      _vm.configs.showPoint
                        ? _c(
                            "div",
                            {
                              staticClass: "spec-sel-point",
                              attrs: { slot: "prefix" },
                              slot: "prefix"
                            },
                            [_vm._v(" " + _vm._s(_vm.selPointStr) + " ")]
                          )
                        : _vm._e()
                    ],
                    2
                  ),
                  _vm._v(" "),
                  _vm.configs.showPoint
                    ? _c("div", { staticClass: "self-point-info" }, [
                        _c("span", [
                          _vm._v(
                            _vm._s(_vm.$t("cloudbrainObj.balanceOfPoints")) +
                              "："
                          ),
                          _c("span", { staticStyle: { color: "red" } }, [
                            _vm._v(" " + _vm._s(_vm.configs.blance.toFixed(2)))
                          ]),
                          _vm._v(
                            "\n            " +
                              _vm._s(_vm.$t("cloudbrainObj.points"))
                          ),
                          _vm.showUseTime
                            ? _c("span", [
                                _vm._v(
                                  _vm._s(_vm.$t("cloudbrainObj.canUseTime")) +
                                    " "
                                ),
                                _c("span", { staticStyle: { color: "red" } }, [
                                  _vm._v(_vm._s(_vm.canUseTime))
                                ]),
                                _vm._v(
                                  " " + _vm._s(_vm.$t("cloudbrainObj.hours"))
                                )
                              ])
                            : _vm._e()
                        ]),
                        _vm._v(" "),
                        _c("span", [
                          _c("i", { staticClass: "el-icon-question" }),
                          _vm._v(" "),
                          _c(
                            "a",
                            {
                              attrs: {
                                target: "_blank",
                                href: "/reward/point/rule"
                              }
                            },
                            [
                              _vm._v(
                                _vm._s(_vm.$t("cloudbrainObj.PointGainDescr"))
                              )
                            ]
                          )
                        ])
                      ])
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tips
                    ? _c("div", {
                        staticClass: "dynamic-tips",
                        domProps: { innerHTML: _vm._s(_vm.tips) }
                      })
                    : _vm._e()
                ],
                1
              )
        ]
      )
    ]),
    _vm._v(" "),
    _c("div", { staticClass: "right-area resource-descr-c" }, [
      !_vm.configs.hideHelpLink
        ? _c(
            "div",
            { staticClass: "resource-descr" },
            [
              _c(
                "el-tooltip",
                { attrs: { placement: "top", effect: "light" } },
                [
                  _c("i", {
                    staticClass: "question circle icon link",
                    staticStyle: { "margin-top": "-7px" }
                  }),
                  _vm._v(" "),
                  _c("div", { attrs: { slot: "content" }, slot: "content" }, [
                    _c(
                      "div",
                      {
                        staticStyle: { width: "200px", "text-align": "center" }
                      },
                      [_vm._v(_vm._s(_vm.$t("specObj.resSelectTips")))]
                    )
                  ])
                ]
              ),
              _vm._v(" "),
              _c(
                "a",
                {
                  attrs: {
                    target: "_blank",
                    href:
                      "https://openi.pcl.ac.cn/docs/index.html#/quickstart/resources"
                  }
                },
                [_vm._v(_vm._s(_vm.$t("cloudbrainObj.specDescr")))]
              )
            ],
            1
          )
        : _vm._e()
    ])
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=template&id=6151937c&scoped=true":
/*!*****************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskLimitDialog.vue?vue&type=template&id=6151937c&scoped=true ***!
  \*****************************************************************************************************************************************************************************************************************************************************/
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
  return _vm.visible
    ? _c("div", { staticClass: "task-limit-overlay" }, [
        _c("div", { staticClass: "task-limit-dialog" }, [
          _c("div", { staticClass: "dialog-header" }, [
            _c("div", { staticClass: "dialog-tip-wrap" }, [
              _vm._m(0),
              _vm._v(" "),
              _c("div", { staticClass: "tip-wrap" }, [
                _c("p", { staticClass: "tip-title" }, [
                  _vm._v(_vm._s(_vm.title))
                ]),
                _vm._v(" "),
                _vm.desc
                  ? _c("p", { staticClass: "tip-desc" }, [
                      _vm._v(_vm._s(_vm.desc))
                    ])
                  : _vm._e()
              ])
            ]),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "close-wrap", on: { click: _vm.handleClose } },
              [_c("i", { staticClass: "ri-close-line" })]
            )
          ]),
          _vm._v(" "),
          _c(
            "div",
            { staticClass: "task-list" },
            _vm._l(_vm.taskList, function(item) {
              return _c(
                "div",
                { key: item.task.id, staticClass: "task-item" },
                [
                  _c("div", { staticClass: "task-info-wrap" }, [
                    _c(
                      "div",
                      { staticClass: "cb-job", class: item.task.job_type },
                      _vm._l(
                        _vm.getJobTypeLines(item.task.jobTypeShow),
                        function(line, i) {
                          return _c("span", { key: i }, [_vm._v(_vm._s(line))])
                        }
                      ),
                      0
                    ),
                    _vm._v(" "),
                    _c("div", { staticClass: "task-info" }, [
                      _c("div", { staticClass: "task-name" }, [
                        _vm._v(_vm._s(item.task.display_job_name))
                      ]),
                      _vm._v(" "),
                      _c("div", { staticClass: "task-meta" }, [
                        _c(
                          "div",
                          {
                            staticClass: "task-status",
                            class: "status-" + item.task.status
                          },
                          [
                            _c("div", { staticClass: "dot" }),
                            _vm._v(" "),
                            _c("span", [_vm._v(_vm._s(item.task.status))])
                          ]
                        ),
                        _vm._v(" "),
                        _c("div", { staticClass: "task-unix" }, [
                          _c("span", [
                            _vm._v(
                              " " + _vm._s(item.task.computeSourceShow) + "，"
                            )
                          ]),
                          _vm._v(" "),
                          _c("span", [
                            _vm._v(
                              " " + _vm._s(item.task.accCardTypeShow) + " "
                            )
                          ])
                        ]),
                        _vm._v(" "),
                        _c("div", { staticClass: "task-unix" }, [
                          _c("span", [
                            _vm._v(
                              " " +
                                _vm._s(_vm.$t("cloudbrainObj.runDuration")) +
                                "："
                            )
                          ]),
                          _vm._v(" "),
                          _c("span", [
                            _vm._v(" " + _vm._s(item.task.formatted_duration))
                          ])
                        ]),
                        _vm._v(" "),
                        _c("div", { staticClass: "task-unix" }, [
                          _c("span", [
                            _vm._v(
                              " " +
                                _vm._s(_vm.$t("cloudbrainObj.createTime")) +
                                "： "
                            )
                          ]),
                          _vm._v(" "),
                          _c("span", [
                            _vm._v(
                              " " +
                                _vm._s(_vm.dateFormat(item.task.created_unix)) +
                                " "
                            )
                          ])
                        ])
                      ])
                    ])
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "task-actions" }, [
                    _vm.isTaskRunning(item.task)
                      ? _c(
                          "span",
                          {
                            staticClass: "action-btn stop",
                            on: {
                              click: function($event) {
                                $event.stopPropagation()
                                return _vm.handleStopTask(item.task)
                              }
                            }
                          },
                          [_vm._v(_vm._s(_vm.$t("cloudbrainObj.stop")))]
                        )
                      : _vm.isTaskStopping(item.task)
                      ? _c(
                          "span",
                          { staticClass: "action-btn stop disabled" },
                          [_vm._v(_vm._s(_vm.$t("cloudbrainObj.stop")))]
                        )
                      : _c(
                          "span",
                          {
                            staticClass: "action-btn delete",
                            class: { disabled: !_vm.canDeleteTask(item.task) },
                            on: {
                              click: function($event) {
                                $event.stopPropagation()
                                _vm.canDeleteTask(item.task) &&
                                  _vm.handleDeleteTask(item.task)
                              }
                            }
                          },
                          [_vm._v(_vm._s(_vm.$t("cloudbrainObj.delete")))]
                        )
                  ])
                ]
              )
            }),
            0
          )
        ])
      ])
    : _vm._e()
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "icon-wrap" }, [
      _c("i", { staticClass: "ri-error-warning-fill" })
    ])
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=template&id=09a23d7c&scoped=true":
/*!**********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/cloudbrain/TaskName.vue?vue&type=template&id=09a23d7c&scoped=true ***!
  \**********************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "form-row" }, [
    _c("div", { staticClass: "left-area" }, [
      _c("div", { staticClass: "title" }, [
        _c("span", { class: _vm.required ? "required" : "" }, [
          _vm._v(_vm._s(_vm.$t("cloudbrainObj.taskName")))
        ])
      ]),
      _vm._v(" "),
      _c(
        "div",
        { staticClass: "content", class: _vm.errStatus ? "error" : "" },
        [
          _c("el-input", {
            staticClass: "field-input",
            attrs: {
              maxlength: _vm.type == "supercompute" ? 26 : 36,
              placeholder: _vm.$t("cloudbrainObj.taskName"),
              autofocus: _vm.autofocus
            },
            on: { input: _vm.handleInput, change: _vm.handleInputChange },
            model: {
              value: _vm.currentValue,
              callback: function($$v) {
                _vm.currentValue = $$v
              },
              expression: "currentValue"
            }
          }),
          _vm._v(" "),
          _vm.type === "supercompute"
            ? _c("div", { staticClass: "tips" }, [
                _vm._v(_vm._s(_vm.$t("cloudbrainObj.taskNameTips1")))
              ])
            : _c("div", { staticClass: "tips" }, [
                _vm._v(_vm._s(_vm.$t("cloudbrainObj.taskNameTips")))
              ])
        ],
        1
      )
    ]),
    _vm._v(" "),
    _c("div", { staticClass: "right-area" })
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ })

}]);