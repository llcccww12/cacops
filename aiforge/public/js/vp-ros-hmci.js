/******/ (function() { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=script&lang=js":
/*!*************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=script&lang=js ***!
  \*************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.splice */ "./node_modules/core-js/modules/es.array.splice.js");
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_1__);


//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "ModelCondition",
  props: {
    condition: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  components: {},
  data: function data() {
    return {
      isLogin: false,
      tabList: [{
        key: '1',
        label: this.$t('modelObj.model_public')
      }, {
        key: '2',
        label: this.$t('modelObj.model_my')
      }, {
        key: '3',
        label: this.$t('modelObj.model_collected')
      }, {
        key: '5',
        label: this.$t('modelObj.model_my_migrate')
      }],
      sortList: [{
        key: '',
        label: this.$t('datasets.default')
      }, {
        key: 'created_unix',
        label: this.$t('datasets.newest')
      }, {
        key: 'updated_unix',
        label: this.$t('datasets.recentupdate')
      }, {
        key: 'download_count',
        label: this.$t('datasets.downloadtimes')
      }, {
        key: 'collected_count',
        label: this.$t('datasets.moststars')
      }, {
        key: 'reference_count',
        label: this.$t('datasets.mostusecount')
      }, {
        key: 'derivative_count',
        label: this.$t('modelManage.mostDerivative')
      }],
      conds: {
        tab: 'public',
        sort: '',
        onlyRecommend: false
      }
    };
  },
  methods: {
    changeTab: function changeTab(item) {
      this.conds.tab = item.key;
      this.$emit('changeCondition', {
        tab: item.key
      });
    },
    changeRecommend: function changeRecommend(item, _item) {
      this.$emit('changeCondition', {
        onlyRecommend: this.conds.onlyRecommend
      });
    },
    changeOnline: function changeOnline(item) {
      this.$emit('changeCondition', {
        hasOnlineUrl: this.conds.hasOnlineUrl
      });
    },
    changeSort: function changeSort(item) {
      this.conds.sort = item.key;
      this.$emit('changeCondition', {
        sort: this.conds.sort
      });
    }
  },
  watch: {
    condition: {
      handler: function handler(newVal) {
        this.conds.tab = newVal.tab || '1';
        this.conds.sort = newVal.sort || '';
        this.conds.onlyRecommend = newVal.onlyRecommend || false;
        this.conds.hasOnlineUrl = newVal.hasOnlineUrl || false;
      },
      immediate: true,
      deep: true
    }
  },
  beforeMount: function beforeMount() {
    this.isLogin = !!document.querySelector('meta[name="_uid"]');

    if (!this.isLogin) {
      this.tabList.splice(1, Infinity);
    }
  },
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=script&lang=js":
/*!***********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=script&lang=js ***!
  \***********************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.string.search */ "./node_modules/core-js/modules/es.string.search.js");
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var _apis_modules_modelsquare__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ~/apis/modules/modelsquare */ "./web_src/vuepages/apis/modules/modelsquare.js");






//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "ModelFilters",
  props: {
    condition: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  components: {},
  data: function data() {
    return {
      engineFlag: false,
      labelFlag: false,
      Engine: [],
      Label: [],
      engineValue: '',
      labelValue: ''
    };
  },
  methods: {
    selectEngine: function selectEngine(item) {
      this.Engine.forEach(function (element) {
        if (element.k == item.k) {
          element.active = true;
        } else {
          element.active = false;
        }
      });
      this.engineFlag = true;
      this.engineValue = item.k;
      this.search();
    },
    selectLabel: function selectLabel(item) {
      this.Label.forEach(function (element) {
        if (element.k == item.k) {
          element.active = true;
        } else {
          element.active = false;
        }
      });
      this.labelValue = item.k;
      this.labelFlag = true;
      this.search();
    },
    clearSelectLeft: function clearSelectLeft(type) {
      if (type === 'engine') {
        this.Engine.forEach(function (element) {
          element.active = false;
        });
        this.engineValue = '';
        this.engineFlag = false;
      } else if (type === 'label') {
        this.Label.forEach(function (element) {
          element.active = false;
        });
        this.labelValue = '';
        this.labelFlag = false;
      }

      this.search();
    },
    search: function search() {
      this.$emit('changeCondition', {
        engine: this.engineValue,
        label: this.labelValue
      });
    }
  },
  beforeMount: function beforeMount() {
    var _this = this;

    (0,_apis_modules_modelsquare__WEBPACK_IMPORTED_MODULE_6__.getModelSqaureFilters)().then(function (res) {
      var oriData = res.data;

      if (oriData && oriData[0]) {
        var data = {
          engine: oriData[0].frame.split(',').map(function (item) {
            return {
              k: item.split(':')[0],
              v: item.split(':')[1],
              active: false
            };
          }),
          label: oriData[0].label.split(',').map(function (item) {
            return {
              k: item,
              v: item,
              active: false
            };
          })
        };
        _this.Engine = data.engine;
        _this.Label = data.label;
      }
    })["catch"](function (err) {
      console.log(err);
    });
  },
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=script&lang=js":
/*!********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=script&lang=js ***!
  \********************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _apis_modules_modelsquare__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ~/apis/modules/modelsquare */ "./web_src/vuepages/apis/modules/modelsquare.js");
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "ModelItem",
  props: {
    condition: {
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
  components: {},
  data: function data() {
    return {
      isCollected: false,
      collectedCount: 0,
      isSetting: false,
      hasOnlineUrl: 0,
      canChangeFav: true
    };
  },
  watch: {
    data: function data(val, oVal) {
      this.updateData();
    }
  },
  methods: {
    changeFav: function changeFav(item) {
      var _this = this;

      if (this.condition.tab == 2 || this.condition.tab == 5) return;
      if (this.isSetting) return;
      this.isSetting = true;
      (0,_apis_modules_modelsquare__WEBPACK_IMPORTED_MODULE_0__.setModelFav)({
        id: item.id,
        collected: this.isCollected ? false : true
      }).then(function (res) {
        _this.isSetting = false;

        if (res.data.code == '0') {
          _this.isCollected = !_this.isCollected;
          _this.collectedCount = _this.collectedCount + (_this.isCollected ? 1 : -1);

          _this.$message.success(_this.isCollected ? _this.$t('datasets.starSuccess') : _this.$t('datasets.unstarSuccess'));

          _this.$emit('changeFav');
        } else if (res.data.code == '401') {
          window.location.href = "/user/login?redirect_to=".concat(encodeURIComponent(window.location.href));
        } else {
          _this.$message.error(res.data.msg);
        }
      })["catch"](function (err) {
        console.log(err);

        _this.$message.error(err);

        _this.isSetting = false;
      });
    },
    updateData: function updateData() {
      this.isCollected = this.data.isCollected;
      this.collectedCount = this.data.collectedCount;
      this.hasOnlineUrl = this.data.hasOnlineUrl;
      this.canChangeFav = !(this.condition.tab == 2 || this.condition.tab == 5);
    }
  },
  beforeMount: function beforeMount() {
    this.updateData();
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=script&lang=js":
/*!********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=script&lang=js ***!
  \********************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.object.define-properties */ "./node_modules/core-js/modules/es.object.define-properties.js");
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.object.define-property */ "./node_modules/core-js/modules/es.object.define-property.js");
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptor */ "./node_modules/core-js/modules/es.object.get-own-property-descriptor.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptors */ "./node_modules/core-js/modules/es.object.get-own-property-descriptors.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.string.search */ "./node_modules/core-js/modules/es.string.search.js");
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! core-js/modules/es.string.trim */ "./node_modules/core-js/modules/es.string.trim.js");
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_18___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_18__);
/* harmony import */ var _ModelItem_vue__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! ./ModelItem.vue */ "./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue");
/* harmony import */ var _apis_modules_modelsquare__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! ~/apis/modules/modelsquare */ "./web_src/vuepages/apis/modules/modelsquare.js");
/* harmony import */ var _const__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! ~/const */ "./web_src/vuepages/const/index.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! element-ui/lib/utils/date-util */ "./node_modules/element-ui/lib/utils/date-util.js");




















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





/* harmony default export */ __webpack_exports__["default"] = ({
  name: "ModelList",
  props: {
    params: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  components: {
    ModelItem: _ModelItem_vue__WEBPACK_IMPORTED_MODULE_19__["default"]
  },
  data: function data() {
    return {
      loading: false,
      list: [],
      iPageSizes: [30],
      iPageSize: 30,
      iPage: 1,
      total: 0
    };
  },
  watch: {
    params: {
      handler: function handler(val, oval) {
        this.search();
      },
      deep: true
    }
  },
  methods: {
    getListData: function getListData() {
      var _this = this;

      this.loading = true;
      (0,_apis_modules_modelsquare__WEBPACK_IMPORTED_MODULE_20__.getModelList)({
        q: this.params.q,
        queryType: this.params.tab,
        orderBy: this.params.sort,
        recommend: this.params.onlyRecommend,
        hasOnlineUrl: this.params.hasOnlineUrl,
        frame: this.params.engine,
        label: this.params.label,
        page: this.iPage,
        pageSize: this.iPageSize,
        notNeedEmpty: false
      }).then(function (res) {
        res = res.data;
        _this.loading = false;
        _this.total = res.count || 0;
        _this.list = (res.data || []).map(function (item) {
          return _objectSpread(_objectSpread({}, item), {}, {
            labels: item.label ? item.label.trim().split(/\s+/) : [],
            engineName: (0,_utils__WEBPACK_IMPORTED_MODULE_22__.getListValueWithKey)(_const__WEBPACK_IMPORTED_MODULE_21__.MODEL_ENGINES, item.engine.toString()),
            createTimeStr: (0,element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_23__.formatDate)(new Date(item.createdUnix * 1000), 'yyyy-MM-dd'),
            updateTimeStr: (0,element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_23__.formatDate)(new Date(item.updatedUnix * 1000), 'yyyy-MM-dd')
          });
        });
      })["catch"](function (err) {
        console.log(err);
        _this.loading = false;
        _this.list = [];
        _this.total = 0;
      });
    },
    search: function search() {
      this.iPage = 1;
      this.getListData();
    },
    changeFav: function changeFav() {
      if (this.params.tab == '3') {
        this.getListData();
      }
    },
    currentChange: function currentChange(page) {
      this.iPage = page;
      this.getListData();
    },
    sizeChange: function sizeChange(pageSize) {
      this.iPageSize = pageSize;
      this.getListData();
    }
  },
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=script&lang=js":
/*!*******************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=script&lang=js ***!
  \*******************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.date.now */ "./node_modules/core-js/modules/es.date.now.js");
/* harmony import */ var core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_now__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var dayjs_plugin_relativeTime__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! dayjs/plugin/relativeTime */ "./node_modules/dayjs/plugin/relativeTime.js");
/* harmony import */ var dayjs_plugin_relativeTime__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(dayjs_plugin_relativeTime__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var dayjs_plugin_localizedFormat__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! dayjs/plugin/localizedFormat */ "./node_modules/dayjs/plugin/localizedFormat.js");
/* harmony import */ var dayjs_plugin_localizedFormat__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(dayjs_plugin_localizedFormat__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var dayjs_locale_zh_cn__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! dayjs/locale/zh-cn */ "./node_modules/dayjs/locale/zh-cn.js");
/* harmony import */ var dayjs_locale_zh_cn__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(dayjs_locale_zh_cn__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var dayjs_locale_en__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! dayjs/locale/en */ "./node_modules/dayjs/locale/en.js");
/* harmony import */ var dayjs_locale_en__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(dayjs_locale_en__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var dayjs__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! dayjs */ "./node_modules/dayjs/dayjs.min.js");
/* harmony import */ var dayjs__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(dayjs__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");


//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//







dayjs__WEBPACK_IMPORTED_MODULE_6___default().locale(_langs__WEBPACK_IMPORTED_MODULE_7__.lang == 'zh-CN' ? 'zh-cn' : 'en');
dayjs__WEBPACK_IMPORTED_MODULE_6___default().extend((dayjs_plugin_relativeTime__WEBPACK_IMPORTED_MODULE_2___default()));
dayjs__WEBPACK_IMPORTED_MODULE_6___default().extend((dayjs_plugin_localizedFormat__WEBPACK_IMPORTED_MODULE_3___default()));
/* harmony default export */ __webpack_exports__["default"] = ({
  name: "ReposItem",
  props: {
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    topic: {
      type: String,
      "default": ''
    },
    topicLink: {
      type: Boolean,
      "default": true
    }
  },
  components: {},
  data: function data() {
    return {
      contributors: []
    };
  },
  methods: {
    calcFromNow: function calcFromNow(unix) {
      // return dayjs(unix * 1000).fromNow();
      return (0,_utils__WEBPACK_IMPORTED_MODULE_8__.timeSinceUnix)(unix, Date.now() / 1000);
    },
    dateFormat: function dateFormat(unix) {
      return _langs__WEBPACK_IMPORTED_MODULE_7__.lang == 'zh-CN' ? dayjs__WEBPACK_IMPORTED_MODULE_6___default()(unix * 1000).format('YYYY年MM月DD日 HH时mm分ss秒') : dayjs__WEBPACK_IMPORTED_MODULE_6___default()(unix * 1000).format('ddd, D MMM YYYY HH:mm:ss [CST]');
    }
  },
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=script&lang=js":
/*!*****************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=script&lang=js ***!
  \*****************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
//
//
//
//
/* harmony default export */ __webpack_exports__["default"] = ({
  name: "App"
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=script&lang=js":
/*!**********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _NavigationBar_vue__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./NavigationBar.vue */ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue");
//
//
//
//
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
  name: "AppBanner",
  props: {},
  components: {
    NavigationBar: _NavigationBar_vue__WEBPACK_IMPORTED_MODULE_0__["default"]
  },
  data: function data() {
    return {
      list: []
    };
  },
  methods: {},
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=script&lang=js":
/*!**************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=script&lang=js ***!
  \**************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_3__);




//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  data: function data() {
    return {};
  },
  mounted: function mounted() {},
  computed: {
    imgSrc: function imgSrc() {
      var _this = this;

      return function (id) {
        var isActiveRoute = _this.$route.name === _this.getRouteNameById(id);

        return "/img/ros-hmci/icon".concat(isActiveRoute ? '-' : '').concat(id, ".png");
      };
    }
  },
  methods: {
    getRouteNameById: function getRouteNameById(id) {
      var routeNames = {
        1: 'HomePage',
        2: 'OSSystem',
        3: 'OpenApp',
        4: 'OpenDataset',
        5: 'OpenModel',
        6: 'CommunitySource',
        7: 'HelpCenter'
      };
      return routeNames[id] || '';
    },
    addLineToImgSrc: function addLineToImgSrc(id) {
      if (this.$route.name === this.getRouteNameById(id)) return;
      var imgElement = document.getElementById("nav".concat(id)).getElementsByTagName('img')[0];
      imgElement.src = imgElement.src.replace("icon".concat(id), "icon-".concat(id));
    },
    removeLineFromImgSrc: function removeLineFromImgSrc(id) {
      if (this.$route.name === this.getRouteNameById(id)) return;
      var imgElement = document.getElementById("nav".concat(id)).getElementsByTagName('img')[0];
      imgElement.src = imgElement.src.replace("icon-".concat(id), "icon".concat(id));
    },
    gotoHomePage: function gotoHomePage() {
      // console.log(this.$router);
      this.$router.push("/");
    },
    gotoOSSystem: function gotoOSSystem() {
      // console.log(this.$router);
      this.$router.push("/os-system");
    },
    gotoOpenApp: function gotoOpenApp() {
      // console.log(this.$router);
      this.$router.push("/open-app");
    },
    gotoOpenData: function gotoOpenData() {
      // console.log(this.$router);
      this.$router.push("/open-data");
    },
    gotoOpenModel: function gotoOpenModel() {
      // console.log(this.$router);
      this.$router.push("/open-model");
    },
    gotoComunitySource: function gotoComunitySource() {
      // console.log(this.$router);
      this.$router.push("/community-source");
    },
    gotoHelpCenter: function gotoHelpCenter() {
      // console.log(this.$router);
      this.$router.push("/help-center");
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=script&lang=js":
/*!**************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=script&lang=js ***!
  \**************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_date_to_iso_string__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.date.to-iso-string */ "./node_modules/core-js/modules/es.date.to-iso-string.js");
/* harmony import */ var core_js_modules_es_date_to_iso_string__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_iso_string__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");
/* harmony import */ var element_ui__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! element-ui */ "./node_modules/element-ui/lib/element-ui.common.js");
/* harmony import */ var element_ui__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(element_ui__WEBPACK_IMPORTED_MODULE_10__);









//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  props: {
    isSigned: {
      type: String,
      "default": "false"
    },
    dataGet: {
      type: String,
      "default": ""
    },
    searchValue: {
      type: String,
      "default": ""
    },
    searchFlag: {
      type: Boolean,
      "default": false
    },
    categoryValue: {
      type: String,
      "default": ""
    },
    taskValue: {
      type: String,
      "default": ""
    },
    licenseValue: {
      type: String,
      "default": ""
    }
  },
  data: function data() {
    return {
      publicDataList: [],
      total: 0,
      params: {
        sort: "",
        q: "",
        page: 1,
        pageSize: 30,
        recommend: false,
        category: "",
        task: "ros_hmci_datasets",
        license: ""
      },
      checked: false,
      sortList: [{
        name: "default",
        active: true
      }, {
        name: "latest",
        active: false
      }, {
        name: "oldest",
        active: false
      }, {
        name: "recentupdate",
        active: false
      }, {
        name: "leastupdate",
        active: false
      }, {
        name: "downloadtimes",
        active: false
      }, {
        name: "moststars",
        active: false
      }, {
        name: "mostusecount",
        active: false
      }],
      showEmpty: false,
      loading: false
    };
  },
  watch: {
    searchValue: function searchValue(newVal) {
      if (!newVal) {
        this.params.page = 1;
        this.params.q = newVal;
        this.getDataList(this.dataGet);
      }
    },
    searchFlag: function searchFlag(newVal) {
      this.params.page = 1;
      this.params.q = this.searchValue;
      this.getDataList(this.dataGet);
    },
    categoryValue: function categoryValue(val) {
      this.params.category = val;
      this.params.page = 1;
      this.getDataList(this.dataGet);
    },
    taskValue: function taskValue(val) {
      this.params.task = val;
      this.params.page = 1;
      this.getDataList(this.dataGet);
    },
    licenseValue: function licenseValue(val) {
      this.params.license = val;
      this.params.page = 1;
      this.getDataList(this.dataGet);
    }
  },
  methods: {
    getDataList: function getDataList(dataType) {
      var _this = this;

      var url = "/explore/".concat(dataType);
      this.loading = true;
      (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_9__.getDatasets)(url, this.params).then(function (res) {
        if (res.data.result_code === "0") {
          if (res.data.data === "null" || res.data.data == "[]") {
            _this.showEmpty = true;
            _this.total = 0;
          } else {
            _this.publicDataList = JSON.parse(res.data.data);

            if (dataType === "my_favorite_datasets") {
              _this.publicDataList.forEach(function (ele, index) {
                _this.publicDataList[index].IsStaring = true;
              });
            }

            _this.total = Number(res.data.count);
            _this.showEmpty = false;
          }

          _this.loading = false;
        } else {
          element_ui__WEBPACK_IMPORTED_MODULE_10__.Message.error(res.data.error_msg);
          _this.loading = false;
        }
      })["catch"](function (err) {
        element_ui__WEBPACK_IMPORTED_MODULE_10__.Message.error(err);
        _this.loading = false;
      });
    },
    handleCurrentChange: function handleCurrentChange(val) {
      this.params.page = val;
      this.getDataList(this.dataGet);
    },
    handleSizeChange: function handleSizeChange(val) {
      this.params.pageSize = val;
      this.getDataList(this.dataGet);
    },
    gotoDataset: function gotoDataset(item) {
      location.href = "/".concat(item.Repo.OwnerName, "/").concat(item.Repo.Name, "/datasets");
    },
    postSquareStar: function postSquareStar(item, index) {
      var _this2 = this;

      if (this.isSigned === "false" || !this.isSigned || this.dataGet == "my_datasets") return;
      var baseUrl = "/".concat(item.Repo.OwnerName, "/").concat(item.Repo.Name, "/datasets/").concat(item.ID, "/");
      var url = item.IsStaring ? baseUrl + "unstar" : baseUrl + "star";
      var changeItem = item;
      (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_9__.putDatasetStar)(url).then(function (res) {
        if (res.data.Code === 0) {
          if (_this2.dataGet == "my_favorite_datasets") {
            _this2.getDataList("my_favorite_datasets");

            element_ui__WEBPACK_IMPORTED_MODULE_10__.Message.success(_this2.$t("datasets.unstarSuccess"));
            return;
          }

          if (item.IsStaring) {
            changeItem.IsStaring = false;
            changeItem.NumStars = changeItem.NumStars - 1;

            _this2.$set(_this2.publicDataList, index, changeItem);

            _this2.$set(_this2.publicDataList, index, changeItem);

            element_ui__WEBPACK_IMPORTED_MODULE_10__.Message.success(_this2.$t("datasets.unstarSuccess"));
          } else {
            changeItem.IsStaring = true;
            changeItem.NumStars = changeItem.NumStars + 1;

            _this2.$set(_this2.publicDataList, index, changeItem);

            _this2.$set(_this2.publicDataList, index, changeItem);

            element_ui__WEBPACK_IMPORTED_MODULE_10__.Message.success(_this2.$t("datasets.starSuccess"));
          }
        } else {
          element_ui__WEBPACK_IMPORTED_MODULE_10__.Message.error(res.data.Message);
        }
      })["catch"](function (err) {
        element_ui__WEBPACK_IMPORTED_MODULE_10__.Message.error(err);
      });
    },
    handleCheckedChange: function handleCheckedChange(val) {
      this.params.recommend = val;
      this.getDataList(this.dataGet);
    },
    handleSort: function handleSort(item) {
      this.sortList.forEach(function (element) {
        element.active = false;
      });
      item.active = true;
      this.params.sort = item.name;
      this.getDataList(this.dataGet);
    },
    chooseLabel: function chooseLabel(item, type) {
      var data = {
        name: item,
        active: false,
        type: type
      };
      this.$emit("getLabel", data);
    }
  },
  filters: {
    DateTransfer: function DateTransfer(unix) {
      var date = new Date(unix * 1000);
      return date.toISOString().slice(0, 10);
    }
  },
  mounted: function mounted() {
    this.getDataList(this.dataGet);
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=script&lang=js":
/*!*************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=script&lang=js ***!
  \*************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.find-index */ "./node_modules/core-js/modules/es.array.find-index.js");
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_0__);

//
//
//
//
//
//
//
//
/* harmony default export */ __webpack_exports__["default"] = ({
  name: "ReposFilters",
  props: {
    defaultsort: {
      type: String,
      "default": 'mostpopular'
    }
  },
  components: {},
  data: function data() {
    return {
      focusIndex: 0,
      list: [{
        key: 'mostpopular',
        label: this.$t('repos.mostPopular')
      }, {
        key: 'mostactive',
        label: this.$t('repos.mostActive')
      }, {
        key: 'recentupdate',
        label: this.$t('repos.recentlyUpdated')
      }, {
        key: 'newest',
        label: this.$t('repos.newest')
      }, {
        key: 'moststars',
        label: this.$t('repos.mostStars')
      }, {
        key: 'mostforks',
        label: this.$t('repos.mostForks')
      } // , {
      //   key: 'mostdatasets',
      //   label: this.$t('repos.mostDatasets'),
      // }, {
      //   key: 'mostaitasks',
      //   label: this.$t('repos.mostAiTasks'),
      // }, {
      //   key: 'mostmodels',
      //   label: this.$t('repos.mostModels'),
      // }
      ]
    };
  },
  methods: {
    changeFilters: function changeFilters(item, index) {
      this.focusIndex = index;
      this.$emit('change', this.list[this.focusIndex]);
    },
    setDefaultFilter: function setDefaultFilter(sort) {
      var index = this.list.findIndex(function (item) {
        return item.key == sort;
      });
      this.focusIndex = index >= 0 ? index : 0;
    }
  },
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=script&lang=js":
/*!**********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.object.define-properties */ "./node_modules/core-js/modules/es.object.define-properties.js");
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.object.define-property */ "./node_modules/core-js/modules/es.object.define-property.js");
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptor */ "./node_modules/core-js/modules/es.object.get-own-property-descriptor.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptors */ "./node_modules/core-js/modules/es.object.get-own-property-descriptors.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_regexp_constructor__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.regexp.constructor */ "./node_modules/core-js/modules/es.regexp.constructor.js");
/* harmony import */ var core_js_modules_es_regexp_constructor__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_constructor__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var _pages_repos_components_ReposItem_vue__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! ~/pages/repos/components/ReposItem.vue */ "./web_src/vuepages/pages/repos/components/ReposItem.vue");
/* harmony import */ var _apis_modules_repos__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! ~/apis/modules/repos */ "./web_src/vuepages/apis/modules/repos.js");
/* harmony import */ var _utils_letteravatar__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! ~/utils/letteravatar */ "./web_src/vuepages/utils/letteravatar.js");


















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



/* harmony default export */ __webpack_exports__["default"] = ({
  name: "ReposList",
  props: {
    q: {
      type: String,
      "default": ""
    },
    sort: {
      type: String,
      "default": "mostpopular"
    },
    topic: {
      type: String,
      "default": ""
    },
    page: {
      type: Number,
      "default": 1
    },
    pageSize: {
      type: Number,
      "default": 15
    },
    pageSizes: {
      type: Array,
      "default": function _default() {
        return [15, 30, 50];
      }
    }
  },
  components: {
    ReposItem: _pages_repos_components_ReposItem_vue__WEBPACK_IMPORTED_MODULE_17__["default"]
  },
  data: function data() {
    return {
      loading: false,
      list: [],
      iPageSizes: [15, 30, 50],
      iPageSize: 15,
      iPage: 1,
      total: 0
    };
  },
  methods: {
    getListData: function getListData() {
      var _this = this;

      this.loading = true;
      (0,_apis_modules_repos__WEBPACK_IMPORTED_MODULE_18__.getReposListData)({
        q: this.q || "",
        topic: this.topic || "ros-hmci-app",
        sort: this.sort || "mostpopular",
        pageSize: this.iPageSize || 15,
        page: this.iPage || 1
      }).then(function (res) {
        res = res.data;
        _this.loading = false;

        if (res.Code == 0) {
          var list = res.Data.Repos || [];
          _this.list = list.map(function (item) {
            item.Contributors = (item.Contributors || []).map(function (_item) {
              return _objectSpread(_objectSpread({}, _item), {}, {
                bgColor: _this.randomColor((_item.Email[0] || '').toLocaleUpperCase())
              });
            });
            var contributors = item.Contributors || [];
            return _objectSpread(_objectSpread({}, item), {}, {
              NameShow: _this.handlerSearchStr(item.Alias, _this.q),
              DescriptionShow: _this.handlerSearchStr(item.Description, _this.q),
              TopicsShow: (item.Topics || []).map(function (_item) {
                return {
                  topic: _item,
                  topicShow: _this.handlerSearchStr(_item, _this.q)
                };
              })
            });
          });
          _this.total = res.Data.Total;
          _this.iPage = _this.iPage;
          _this.iPageSize = _this.iPageSize;

          _this.$nextTick(function () {
            _utils_letteravatar__WEBPACK_IMPORTED_MODULE_19__["default"].transform();
          });
        } else {
          _this.list = [];
          _this.total = 0;
          _this.iPage = _this.iPage;
          _this.iPageSize = _this.iPageSize;
        }
      })["catch"](function (err) {
        console.log(err);
        _this.loading = false;
        _this.list = [];
        _this.total = 0;
        _this.iPage = _this.iPage;
        _this.iPageSize = _this.iPageSize;
      });
    },
    search: function search() {
      this.getListData();
    },
    currentChange: function currentChange(page) {
      this.iPage = page;
      this.$emit("current-change", {
        page: this.iPage,
        pageSize: this.iPageSize
      });
    },
    sizeChange: function sizeChange(pageSize) {
      this.iPageSize = pageSize;
      this.$emit("size-change", {
        page: this.iPage,
        pageSize: this.iPageSize
      });
    },
    handlerSearchStr: function handlerSearchStr(oStr, searchKey) {
      if (!searchKey) return oStr;
      return oStr.replace(new RegExp("(".concat(searchKey, ")"), "ig"), "<font color=\"red\">$1</font>");
    },
    randomColor: function randomColor(t) {
      var tIndex = t.charCodeAt(0);
      var colorList = ["#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#34495e", "#16a085", "#27ae60", "#2980b9", "#8e44ad", "#2c3e50", "#f1c40f", "#e67e22", "#e74c3c", "#00bcd4", "#95a5a6", "#f39c12", "#d35400", "#c0392b", "#bdc3c7", "#7f8c8d"];
      return colorList[tIndex % colorList.length];
    }
  },
  watch: {
    page: {
      handler: function handler(val) {
        this.iPage = val;
      },
      immediate: true
    },
    pageSize: {
      handler: function handler(val) {
        this.iPageSize = val;
      },
      immediate: true
    }
  },
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=script&lang=js":
/*!**********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.find-index */ "./node_modules/core-js/modules/es.array.find-index.js");
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.index-of */ "./node_modules/core-js/modules/es.array.index-of.js");
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.splice */ "./node_modules/core-js/modules/es.array.splice.js");
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.string.search */ "./node_modules/core-js/modules/es.string.search.js");
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.string.trim */ "./node_modules/core-js/modules/es.string.trim.js");
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var _apis_modules_common__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ~/apis/modules/common */ "./web_src/vuepages/apis/modules/common.js");









//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

var COLOR_LIST = ["rgb(255, 104, 104)", "rgb(22, 132, 252)", "rgb(2, 202, 253)", "rgb(164, 145, 215)", "rgb(232, 64, 247)", "rgb(245, 182, 110)", "rgb(54, 187, 166)", "rgb(123, 50, 178)"];
/* harmony default export */ __webpack_exports__["default"] = ({
  name: "SearchBar",
  props: {
    type: {
      type: String,
      "default": "square"
    },
    // square|search
    searchValue: {
      type: String,
      "default": ""
    },
    topic: {
      type: String,
      "default": ""
    },
    sort: {
      type: String,
      "default": ""
    },
    "static": {
      type: Boolean,
      "default": false
    },
    staticTopicsData: {
      type: String,
      "default": "[]"
    }
  },
  components: {},
  data: function data() {
    return {
      searchInputValue: "",
      topicColors: COLOR_LIST,
      defaultColor: "#F6F6F6",
      selectedColor: "",
      selectTopic: "",
      topicOri: [],
      topics: []
    };
  },
  methods: {
    setDefaultSearch: function setDefaultSearch(params) {
      this.searchInputValue = params.q || "";
      this.selectTopic = params.topic || "";
      this.selectTopic && this.changeTopic({
        k: this.selectTopic.toLocaleLowerCase(),
        v: this.selectTopic
      }, true);
    },
    changeTopic: function changeTopic(topicItem, noSearch) {
      var _this = this;

      var index_ori = this.topicOri.findIndex(function (item) {
        return item.k == _this.selectTopic.toLocaleLowerCase();
      });

      if (index_ori < 0 && this.selectTopic) {
        var index = this.topics.findIndex(function (item) {
          return item.k == _this.selectTopic.toLocaleLowerCase();
        });

        if (index > -1) {
          this.topics.splice(index, 1);
        }
      }

      this.selectTopic = topicItem.v;

      if (this.selectTopic && this.topics.indexOf(this.selectTopic) < 0) {
        var _index = this.topics.findIndex(function (item) {
          return item.k == _this.selectTopic.toLocaleLowerCase();
        });

        if (_index < 0) {
          this.topics.push({
            k: this.selectTopic.toLocaleLowerCase(),
            v: this.selectTopic
          });
        }
      }

      !noSearch && this.search();
    },
    handlerTopicsData: function handlerTopicsData(data) {
      try {
        var topicsData = JSON.parse(data);
        var topics = topicsData.map(function (item) {
          return {
            k: item.trim().toLocaleLowerCase(),
            v: item.trim()
          };
        });
        this.topicOri = JSON.parse(JSON.stringify(topics));
        this.topics = topics;
        var selectTopic_key = this.selectTopic.toLocaleLowerCase();

        if (selectTopic_key) {
          var index = this.topics.findIndex(function (item) {
            return item.k == selectTopic_key;
          });

          if (index < 0) {
            this.topics.push({
              k: this.selectTopic.toLocaleLowerCase(),
              v: this.selectTopic
            });
          }
        }
      } catch (err) {
        console.log(err);
      }
    },
    search: function search() {
      this.searchInputValue = this.searchInputValue.trim();

      if (this.type == "square") {
        window.location.href = "/explore/repos?q=".concat(this.searchInputValue, "&sort=").concat(this.sort, "&topic=ros-hmci-app");
      } else {
        this.$emit("change", {
          q: this.searchInputValue,
          topic: this.selectTopic
        });
      }
    }
  },
  mounted: function mounted() {
    var _this2 = this;

    if (this["static"]) {
      try {
        this.handlerTopicsData(this.staticTopicsData);
      } catch (err) {
        console.log(err);
      }
    } else {
      (0,_apis_modules_common__WEBPACK_IMPORTED_MODULE_9__.getPromoteData)("/repos/recommend_topics").then(function (res) {
        var data = res.data;

        _this2.handlerTopicsData(data);
      })["catch"](function (err) {
        console.log(err);

        _this2.handlerTopicsData("[]");
      });
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=script&lang=js":
/*!************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=script&lang=js ***!
  \************************************************************************************************************************************************************************************************************************/
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
/* harmony default export */ __webpack_exports__["default"] = ({
  data: function data() {
    return {
      activeNames: ['1']
    };
  },
  methods: {
    handleChange: function handleChange(val) {// console.log(val);
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=script&lang=js":
/*!***********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=script&lang=js ***!
  \***********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.includes */ "./node_modules/core-js/modules/es.array.includes.js");
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.promise */ "./node_modules/core-js/modules/es.promise.js");
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_string_includes__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.string.includes */ "./node_modules/core-js/modules/es.string.includes.js");
/* harmony import */ var core_js_modules_es_string_includes__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_includes__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var _components_AppBanner_vue__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! ../components/AppBanner.vue */ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue");
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! axios */ "./node_modules/axios/index.js");
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(axios__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var _links_js__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ./links.js */ "./web_src/vuepages/pages/ros-hmci/views/links.js");










//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: 'CommunitySource',
  data: function data() {
    return {
      searchQuery: '',
      links: _links_js__WEBPACK_IMPORTED_MODULE_12__["default"],
      resources: [],
      // 存储所有资源
      currentPage: 1,
      // 当前页码
      pageSize: 15,
      // 每页显示的资源数量
      totalItems: 0 // 资源总数

    };
  },
  created: function created() {
    var _this = this;

    return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_9___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_7___default().mark(function _callee() {
      var response, data;
      return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_7___default().wrap(function _callee$(_context) {
        while (1) {
          switch (_context.prev = _context.next) {
            case 0:
              _context.next = 2;
              return fetch(_this.links.source_Api);

            case 2:
              response = _context.sent;
              _context.next = 5;
              return response.json();

            case 5:
              data = _context.sent;
              // 将数据设置为resources
              _this.resources = data;
              _this.totalItems = _this.resources.length;

            case 8:
            case "end":
              return _context.stop();
          }
        }
      }, _callee);
    }))();
  },
  components: {
    AppBanner: _components_AppBanner_vue__WEBPACK_IMPORTED_MODULE_10__["default"]
  },
  methods: {
    handleSizeChange: function handleSizeChange(newSize) {
      this.pageSize = newSize;
    },
    handleCurrentChange: function handleCurrentChange(newPage) {
      this.currentPage = newPage;
    }
  },
  mounted: function mounted() {},
  computed: {
    filteredResources: function filteredResources() {
      if (!this.searchQuery) {
        return this.resources;
      }

      var query = this.searchQuery.toLowerCase();
      return this.resources.filter(function (resource) {
        return resource.name.toLowerCase().includes(query) || resource.synopsis.toLowerCase().includes(query);
      });
    },
    paginatedResources: function paginatedResources() {
      var start = (this.currentPage - 1) * this.pageSize;
      var end = this.currentPage * this.pageSize;
      return this.filteredResources.slice(start, end);
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=script&lang=js":
/*!******************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=script&lang=js ***!
  \******************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _components_NavigationBar_vue__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../components/NavigationBar.vue */ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue");
/* harmony import */ var _components_helpCollaps_vue__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../components/helpCollaps.vue */ "./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue");
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "HelpCenter",
  data: function data() {
    return {};
  },
  components: {
    NavigationBar: _components_NavigationBar_vue__WEBPACK_IMPORTED_MODULE_0__["default"],
    helpCollaps: _components_helpCollaps_vue__WEBPACK_IMPORTED_MODULE_1__["default"]
  },
  methods: {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=script&lang=js":
/*!****************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=script&lang=js ***!
  \****************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.promise */ "./node_modules/core-js/modules/es.promise.js");
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var _components_NavigationBar_vue__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ../components/NavigationBar.vue */ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue");
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! axios */ "./node_modules/axios/index.js");
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(axios__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var _links_js__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ./links.js */ "./web_src/vuepages/pages/ros-hmci/views/links.js");





//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "HomePage",
  data: function data() {
    return {
      hoveredImage: -1,
      activeIndex: -1,
      statisticsData: {
        developer: "",
        organization: ""
      },
      links: _links_js__WEBPACK_IMPORTED_MODULE_7__["default"],
      showText: [false, false],
      feedbackImgSrc: "/img/ros-hmci/group7305.png",
      qrcodeImgSrc: "/img/ros-hmci/group7307.png",
      systemComponentTotal: 0,
      applicationSoftwareTotal: 0,
      typical: []
    };
  },
  components: {
    NavigationBar: _components_NavigationBar_vue__WEBPACK_IMPORTED_MODULE_5__["default"]
  },
  mounted: function mounted() {
    window.addEventListener('scroll', this.handleScroll);
    this.fetchStatisticsData();
    this.fetchTypical();
    this.fetchTotal("ros-hmci-os", "systemComponentTotal");
    this.fetchTotal("ros-hmci-app", "applicationSoftwareTotal");
  },
  beforeDestroy: function beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll);
  },
  computed: {
    textVisible: function textVisible() {
      return {
        0: this.showText[0],
        1: this.showText[1]
      };
    }
  },
  methods: {
    fetchStatisticsData: function fetchStatisticsData() {
      var _this = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default().mark(function _callee() {
        var response;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default().wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                _context.prev = 0;
                _context.next = 3;
                return axios__WEBPACK_IMPORTED_MODULE_6___default().get(_this.links.home_statisticsData);

              case 3:
                response = _context.sent;

                /*
                json格式如下
                {
                  "developer": "500+",
                  "organization": "100+"
                }
                */
                _this.statisticsData.developer = response.data.developer;
                _this.statisticsData.organization = response.data.organization;
                _context.next = 11;
                break;

              case 8:
                _context.prev = 8;
                _context.t0 = _context["catch"](0);
                console.error("Error fetching data:", _context.t0);

              case 11:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, null, [[0, 8]]);
      }))();
    },
    fetchTypical: function fetchTypical() {
      var _this2 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default().mark(function _callee2() {
        var response, data;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default().wrap(function _callee2$(_context2) {
          while (1) {
            switch (_context2.prev = _context2.next) {
              case 0:
                _context2.prev = 0;
                _context2.next = 3;
                return fetch("/fanshuai/ROS_hmci_Typical_app/raw/branch/master/typical.json");

              case 3:
                response = _context2.sent;
                _context2.next = 6;
                return response.json();

              case 6:
                data = _context2.sent;
                _this2.typical = data;
                _context2.next = 13;
                break;

              case 10:
                _context2.prev = 10;
                _context2.t0 = _context2["catch"](0);
                console.error("Error fetching data:", _context2.t0);

              case 13:
              case "end":
                return _context2.stop();
            }
          }
        }, _callee2, null, [[0, 10]]);
      }))();
    },
    handleMouseEnter: function handleMouseEnter(id) {
      if (id === 1) {
        this.showText[0] = true;
        this.feedbackImgSrc = "/img/ros-hmci/mbz668.png";
      } else if (id === 2) {
        this.showText[1] = true;
        this.qrcodeImgSrc = "/img/ros-hmci/mbz669.png";
      }
    },
    handleMouseLeave: function handleMouseLeave(id) {
      if (id === 1) {
        this.showText[0] = false;
        this.feedbackImgSrc = "/img/ros-hmci/group7305.png";
      } else if (id === 2) {
        this.showText[1] = false;
        this.qrcodeImgSrc = "/img/ros-hmci/group7307.png";
      }
    },
    fetchTotal: function fetchTotal(topic, targetDataField) {
      var _this3 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default().mark(function _callee3() {
        var response;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default().wrap(function _callee3$(_context3) {
          while (1) {
            switch (_context3.prev = _context3.next) {
              case 0:
                _context3.prev = 0;
                _context3.next = 3;
                return axios__WEBPACK_IMPORTED_MODULE_6___default().get("/explore/repos/search?topic=".concat(encodeURIComponent(topic)));

              case 3:
                response = _context3.sent;
                _this3[targetDataField] = response.data.Data.Total;
                _context3.next = 10;
                break;

              case 7:
                _context3.prev = 7;
                _context3.t0 = _context3["catch"](0);
                console.error("接口错误", _context3.t0);

              case 10:
              case "end":
                return _context3.stop();
            }
          }
        }, _callee3, null, [[0, 7]]);
      }))();
    },
    getdownload: function getdownload(path, id) {
      var path1 = path;
      var Id = id;
      this.$router.push({
        path: path1,
        query: {
          id: Id
        }
      });
    },
    mouseEnter: function mouseEnter(index) {
      this.hoveredImage = index;
      this.activeIndex = index;
      var tabbar = document.querySelector(".tabBar");
      var tabIds = ["#tab1", "#tab2", "#tab3", "#tab4"];

      for (var i = 0; i < tabIds.length; i++) {
        if (i === index) {
          document.querySelector(tabIds[i]).style.display = "";
        } else {
          document.querySelector(tabIds[i]).style.display = "none";
        }
      }

      var positions = [50, 350, 640, 974];
      tabbar.style.left = positions[index] + "px";
    },
    xtjg: function xtjg(index) {
      var xgtIds = ["#jgt1", "#jgt2", "#jgt3", "#jgt4", "#jgt5", "#jgt6"];

      for (var i = 0; i < xgtIds.length; i++) {
        if (i === index - 1) {
          document.querySelector(xgtIds[i]).style.display = "";
        } else {
          document.querySelector(xgtIds[i]).style.display = "none";
        }
      }
    },
    unxtjg: function unxtjg(index) {
      var xgtIds = ["#jgt1", "#jgt2", "#jgt3", "#jgt4", "#jgt5", "#jgt6"];

      for (var i = 0; i < xgtIds.length; i++) {
        if (i === index - 1) {
          document.querySelector(xgtIds[i]).style.display = "none";
        } else {
          document.querySelector(xgtIds[i]).style.display = "none";
        }
      }
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=script&lang=js":
/*!****************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=script&lang=js ***!
  \****************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.symbol.description */ "./node_modules/core-js/modules/es.symbol.description.js");
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_symbol_iterator__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.symbol.iterator */ "./node_modules/core-js/modules/es.symbol.iterator.js");
/* harmony import */ var core_js_modules_es_symbol_iterator__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol_iterator__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4__);
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
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_parse_int__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.parse-int */ "./node_modules/core-js/modules/es.parse-int.js");
/* harmony import */ var core_js_modules_es_parse_int__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_parse_int__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/es.string.iterator */ "./node_modules/core-js/modules/es.string.iterator.js");
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! core-js/modules/web.dom-collections.iterator */ "./node_modules/core-js/modules/web.dom-collections.iterator.js");
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_18___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_18__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_19___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_19__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_20___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_20__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_21___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_21__);
/* harmony import */ var _components_AppBanner_vue__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! ../components/AppBanner.vue */ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue");
/* harmony import */ var _links_js__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! ./links.js */ "./web_src/vuepages/pages/ros-hmci/views/links.js");
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_24__ = __webpack_require__(/*! axios */ "./node_modules/axios/index.js");
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_24___default = /*#__PURE__*/__webpack_require__.n(axios__WEBPACK_IMPORTED_MODULE_24__);























function _createForOfIteratorHelper(o) { if (typeof Symbol === "undefined" || o[Symbol.iterator] == null) { if (Array.isArray(o) || (o = _unsupportedIterableToArray(o))) { var i = 0; var F = function F() {}; return { s: F, n: function n() { if (i >= o.length) return { done: true }; return { done: false, value: o[i++] }; }, e: function e(_e) { throw _e; }, f: F }; } throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method."); } var it, normalCompletion = true, didErr = false, err; return { s: function s() { it = o[Symbol.iterator](); }, n: function n() { var step = it.next(); normalCompletion = step.done; return step; }, e: function e(_e2) { didErr = true; err = _e2; }, f: function f() { try { if (!normalCompletion && it["return"] != null) it["return"](); } finally { if (didErr) throw err; } } }; }

function _unsupportedIterableToArray(o, minLen) { if (!o) return; if (typeof o === "string") return _arrayLikeToArray(o, minLen); var n = Object.prototype.toString.call(o).slice(8, -1); if (n === "Object" && o.constructor) n = o.constructor.name; if (n === "Map" || n === "Set") return Array.from(o); if (n === "Arguments" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return _arrayLikeToArray(o, minLen); }

function _arrayLikeToArray(arr, len) { if (len == null || len > arr.length) len = arr.length; for (var i = 0, arr2 = new Array(len); i < len; i++) { arr2[i] = arr[i]; } return arr2; }

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "OSSystem",
  data: function data() {
    return {
      activeTab: 'tab1',
      links: _links_js__WEBPACK_IMPORTED_MODULE_23__["default"],
      repos: [],
      description: '',
      attachments: [],
      selectedRepoAlias: '',
      selectedRepoName: '',
      selectedRepoOwner: '',
      showLeft: true,
      showRight: true
    };
  },
  filters: {
    formatDate: function formatDate(timestamp) {
      var date = new Date(timestamp * 1000);
      var year = date.getFullYear();
      var month = ('0' + (date.getMonth() + 1)).slice(-2);
      var day = ('0' + date.getDate()).slice(-2);
      return "".concat(year, "-").concat(month, "-").concat(day);
    }
  },
  components: {
    AppBanner: _components_AppBanner_vue__WEBPACK_IMPORTED_MODULE_22__["default"]
  },
  created: function created() {
    this.fetchRepos();
  },
  mounted: function mounted() {
    // window.addEventListener('scroll', this.handleScroll);
    // this.xtxz("sec1");
    this.toLocal();
  },
  beforeDestroy: function beforeDestroy() {// window.removeEventListener('scroll', this.handleScroll);
  },
  methods: {
    handleScroll: function handleScroll() {
      var leftElem = this.$refs.leftElem;
      var rightElem = this.$refs.rightElem;
      var windowHeight = window.innerHeight;
      var scrollY = window.scrollY;
      var leftElemTop = leftElem.getBoundingClientRect().top + scrollY;
      var rightElemTop = rightElem.getBoundingClientRect().top + scrollY;
      var threshold = windowHeight * 0.5;

      if (scrollY + windowHeight > leftElemTop + threshold) {
        this.showLeft = true;
      } else {
        this.showLeft = false;
      }

      if (scrollY + windowHeight > rightElemTop + threshold) {
        this.showRight = true;
      } else {
        this.showRight = false;
      }
    },
    fetchRepos: function fetchRepos() {
      var _this = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_21___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_19___default().mark(function _callee() {
        var apiUrl, response, _iterator, _step, repo;

        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_19___default().wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                _context.prev = 0;
                apiUrl = '/explore/repos/search?q=&topic=ros-hmci-os-xjykd&sort=mostpopular&pageSize=15&page=1';
                _context.next = 4;
                return axios__WEBPACK_IMPORTED_MODULE_24___default().get(apiUrl);

              case 4:
                response = _context.sent;

                if (!(response.data.Code === 0)) {
                  _context.next = 26;
                  break;
                }

                _this.repos = response.data.Data.Repos; // 遍历仓库列表并调用fetchCurrentRepoData方法

                _iterator = _createForOfIteratorHelper(_this.repos);
                _context.prev = 8;

                _iterator.s();

              case 10:
                if ((_step = _iterator.n()).done) {
                  _context.next = 16;
                  break;
                }

                repo = _step.value;
                _context.next = 14;
                return _this.fetchCurrentRepoData(repo.OwnerName, repo.Name);

              case 14:
                _context.next = 10;
                break;

              case 16:
                _context.next = 21;
                break;

              case 18:
                _context.prev = 18;
                _context.t0 = _context["catch"](8);

                _iterator.e(_context.t0);

              case 21:
                _context.prev = 21;

                _iterator.f();

                return _context.finish(21);

              case 24:
                _context.next = 27;
                break;

              case 26:
                // 处理请求失败的情况
                console.error('请求失败：', response.data.Msg);

              case 27:
                _context.next = 32;
                break;

              case 29:
                _context.prev = 29;
                _context.t1 = _context["catch"](0);
                // 处理请求异常的情况
                console.error('请求异常：', _context.t1);

              case 32:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, null, [[0, 29], [8, 18, 21, 24]]);
      }))();
    },
    fetchCurrentRepoData: function fetchCurrentRepoData(repoOwnerName, repoName) {
      var _this2 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_21___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_19___default().mark(function _callee2() {
        var apiUrl, response, responseData, data;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_19___default().wrap(function _callee2$(_context2) {
          while (1) {
            switch (_context2.prev = _context2.next) {
              case 0:
                _context2.prev = 0;
                apiUrl = "/".concat(repoOwnerName, "/").concat(repoName, "/datasets/current_repo_m?q=&page=1&q=&type=0");
                _context2.next = 4;
                return axios__WEBPACK_IMPORTED_MODULE_24___default().get(apiUrl);

              case 4:
                response = _context2.sent;

                // 处理请求成功的情况
                if (response.data.result_code === "0") {
                  // const test1 = response.data.data;
                  // console.log(test1);
                  responseData = JSON.parse(response.data.data); // 解析 "data" 字段的 JSON 字符串
                  // console.log(responseData);

                  data = responseData[0]; // 假设 data 数组中的第一个对象包含所需数据

                  _this2.description = data.Description;
                  _this2.attachments = data.Attachments;
                } else {
                  // 处理请求失败的情况
                  console.error("请求失败：", response.data.Msg);
                }

                _context2.next = 11;
                break;

              case 8:
                _context2.prev = 8;
                _context2.t0 = _context2["catch"](0);
                // 处理请求异常的情况
                console.error("请求异常：", _context2.t0);

              case 11:
              case "end":
                return _context2.stop();
            }
          }
        }, _callee2, null, [[0, 8]]);
      }))();
    },
    toLocal: function toLocal() {
      // console.log(this.$route);
      // 查找存储的锚点id
      var Id = this.$route.query.id;
      var toElement = document.getElementById(Id); // console.log(toElement, "toElement");
      //锚点存在跳转

      if (Id) {
        toElement.scrollIntoView();
      }
    },
    xtjg: function xtjg(tabID) {
      var tabs = document.querySelectorAll(".tab");
      var tabbar = document.querySelector(".section_10");
      tabs.forEach(function (tab) {
        if (tab.id == tabID) {
          tab.classList.add("visible");
        } else {
          tab.classList.remove("visible");
        }
      });
      var tabLefts = {
        "tab1": "0px",
        "tab2": "205px",
        "tab3": "385px",
        "tab4": "575px",
        "tab5": "815px",
        "tab6": "1029px"
      };
      tabbar.style.left = tabLefts[tabID] || "0px"; // 更新激活的标签

      this.activeTab = tabID;
    },
    xtxz: function xtxz(secID) {
      var secs = document.querySelectorAll(".sec");
      secs.forEach(function (sec) {
        var active = sec.querySelector(".xz_active");
        var normal = sec.querySelector(".xz_normal");

        if (sec.id == secID) {
          active.style.display = "";
          normal.style.display = "none";
        } else {
          active.style.display = "none";
          normal.style.display = "";
        }
      });
      var index = parseInt(secID.replace("sec", "")) - 1;
      this.currentRepo = this.repos[index];
    },
    destroyed: function destroyed() {
      localStorage.setItem("toId", "");
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=script&lang=js":
/*!***************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=script&lang=js ***!
  \***************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.index-of */ "./node_modules/core-js/modules/es.array.index-of.js");
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.string.search */ "./node_modules/core-js/modules/es.string.search.js");
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.string.trim */ "./node_modules/core-js/modules/es.string.trim.js");
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var _components_NavigationBar_vue__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ../components/NavigationBar.vue */ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue");
/* harmony import */ var _components_SearchBar_vue__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ../components/SearchBar.vue */ "./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue");
/* harmony import */ var _components_ReposFilters_vue__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ../components/ReposFilters.vue */ "./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue");
/* harmony import */ var _components_ReposList_vue__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! ../components/ReposList.vue */ "./web_src/vuepages/pages/ros-hmci/components/ReposList.vue");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");







//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
 //以下三个组件做过更改





var staticSquareBanners = JSON.stringify(window.staticSquareBanners || []);
var staticSquarePreferredRepos = window.staticSquarePreferredRepos || [];
var staticSquareTopics = JSON.stringify(window.staticSquareTopics || []);
var staticSquareRecommendRepos = window.staticSquareRecommendRepos || [];
/* harmony default export */ __webpack_exports__["default"] = ({
  name: "OpenApp",
  data: function data() {
    return {
      reposListSortType: "mostpopular",
      reposListQurey: "",
      reposListTopic: "",
      page: 1,
      pageSize: 15,
      pageSizes: [15, 30, 50],
      staticSquareBanners: staticSquareBanners,
      staticSquarePreferredRepos: staticSquarePreferredRepos,
      staticSquareTopics: staticSquareTopics,
      staticSquareRecommendRepos: staticSquareRecommendRepos
    };
  },
  components: {
    NavigationBar: _components_NavigationBar_vue__WEBPACK_IMPORTED_MODULE_7__["default"],
    SearchBar: _components_SearchBar_vue__WEBPACK_IMPORTED_MODULE_8__["default"],
    ReposFilters: _components_ReposFilters_vue__WEBPACK_IMPORTED_MODULE_9__["default"],
    ReposList: _components_ReposList_vue__WEBPACK_IMPORTED_MODULE_10__["default"]
  },
  methods: {
    filtersChange: function filtersChange(condition) {
      this.page = 1;
      this.reposListSortType = condition.key;
      this.search();
    },
    searchBarChange: function searchBarChange(params) {
      this.page = 1;
      this.reposListQurey = params.q || "";
      this.reposListTopic = params.topic || "";
      this.search();
    },
    currentChange: function currentChange(_ref) {
      var page = _ref.page,
          pageSize = _ref.pageSize;
      this.page = page;
      this.search();
    },
    sizeChange: function sizeChange(_ref2) {
      var page = _ref2.page,
          pageSize = _ref2.pageSize;
      this.page = 1;
      this.pageSize = pageSize;
      this.search();
    },
    search: function search() {
      window.location.href = "/explore/repos/square?q=".concat(this.reposListQurey.trim(), "&sort=").concat(this.reposListSortType, "&topic=").concat(this.reposListTopic.trim(), "&page=").concat(this.page, "&pageSize=").concat(this.pageSize);
    }
  },
  beforeMount: function beforeMount() {
    var urlParams = (0,_utils__WEBPACK_IMPORTED_MODULE_11__.getUrlSearchParams)();
    this.reposListQurey = urlParams.q || "";
    this.reposListTopic = urlParams.topic || "";
    this.reposListSortType = urlParams.sort || "mostpopular";
    this.page = Number(urlParams.page) || 1;
    this.pageSize = this.pageSizes.indexOf(Number(urlParams.pageSize)) >= 0 ? Number(urlParams.pageSize) : 15;
  },
  mounted: function mounted() {
    var _this = this;

    this.$nextTick(function () {
      // this.$refs.reposFiltersRef.setDefaultFilter(this.reposListSortType);
      _this.$refs.searchBarRef.setDefaultSearch({
        q: _this.reposListQurey,
        topic: _this.reposListTopic
      });

      var urlParams = (0,_utils__WEBPACK_IMPORTED_MODULE_11__.getUrlSearchParams)();
      var page = Number(urlParams.page) || 1;
      var reposListSortType = urlParams.sort;

      if (page != 1 || reposListSortType) {
        window.location.href = "#search";
      }

      _this.$refs.reposListRef.search();
    });
    window.addEventListener("pageshow", function (e) {
      if (e.persisted) {
        window.location.reload();
      }
    }, false);
  },
  beforeDestroy: function beforeDestroy() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=script&lang=js":
/*!*******************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=script&lang=js ***!
  \*******************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _components_AppBanner_vue__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../components/AppBanner.vue */ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue");
/* harmony import */ var _pages_dataset_square_constant_js__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ~/pages/dataset/square/constant.js */ "./web_src/vuepages/pages/dataset/square/constant.js");
/* harmony import */ var _components_PublicDataset_vue__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ../components/PublicDataset.vue */ "./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue");



//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

 //以下组件做过更改


/* harmony default export */ __webpack_exports__["default"] = ({
  name: "OpenDataset",
  components: {
    AppBanner: _components_AppBanner_vue__WEBPACK_IMPORTED_MODULE_3__["default"],
    PublicDataset: _components_PublicDataset_vue__WEBPACK_IMPORTED_MODULE_5__["default"]
  },
  data: function data() {
    return {
      Category: _pages_dataset_square_constant_js__WEBPACK_IMPORTED_MODULE_4__.Category,
      Task: _pages_dataset_square_constant_js__WEBPACK_IMPORTED_MODULE_4__.Task,
      License: _pages_dataset_square_constant_js__WEBPACK_IMPORTED_MODULE_4__.License,
      categoryFlag: false,
      taskFlag: false,
      licenseFlag: false,
      categoryValue: "",
      taskValue: "",
      licenseValue: "",
      activeName: "public_datasets",
      isSigned: "false",
      searchValue: "",
      searchFlag: false
    };
  },
  computed: {},
  methods: {
    handleClick: function handleClick(tab, event) {
      this.searchValue = "";
      this.clearAllSelctLeft();
    },
    selectCategory: function selectCategory(item) {
      this.Category.forEach(function (element) {
        if (element.name === item.name) {
          element.active = true;
        } else {
          element.active = false;
        }
      });
      this.categoryFlag = true;
      this.categoryValue = item.name;
    },
    selectTask: function selectTask(item) {
      this.Task.forEach(function (element) {
        if (element.name === item.name) {
          element.active = true;
        } else {
          element.active = false;
        }
      });
      this.taskValue = item.name;
      this.taskFlag = true;
    },
    selectLicense: function selectLicense(item) {
      this.License.forEach(function (element) {
        if (element.name === item.name) {
          element.active = true;
        } else {
          element.active = false;
        }
      });
      this.licenseValue = item.name;
      this.licenseFlag = true;
    },
    clearSelectLeft: function clearSelectLeft(type) {
      if (type === "category") {
        this.Category.forEach(function (element) {
          element.active = false;
        });
        this.categoryValue = "";
        this.categoryFlag = false;
      } else if (type === "task") {
        this.Task.forEach(function (element) {
          element.active = false;
        });
        this.taskValue = "";
        this.taskFlag = false;
      } else {
        this.License.forEach(function (element) {
          element.active = false;
        });
        this.licenseValue = "";
        this.licenseFlag = false;
      }
    },
    clearAllSelctLeft: function clearAllSelctLeft() {
      if (this.categoryFlag) {
        this.clearSelectLeft("category");
      }

      if (this.taskFlag) {
        this.clearSelectLeft("task");
      }

      if (this.licenseFlag) {
        this.clearSelectLeft("license");
      }
    },
    getChildLabel: function getChildLabel(data) {
      if (data.type === "category") {
        this.selectCategory(data);
      } else if (data.type === "task") {
        this.selectTask(data);
      } else {
        this.selectLicense(data);
      }
    }
  },
  mounted: function mounted() {
    var datasets_tmpl = document.getElementById("datasets-square"); // this.isSigned = datasets_tmpl.getAttribute("data-issigned");

    this.isSigned = true;
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=script&lang=js":
/*!*****************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=script&lang=js ***!
  \*****************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.index-of */ "./node_modules/core-js/modules/es.array.index-of.js");
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.object.define-properties */ "./node_modules/core-js/modules/es.object.define-properties.js");
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.object.define-property */ "./node_modules/core-js/modules/es.object.define-property.js");
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptor */ "./node_modules/core-js/modules/es.object.get-own-property-descriptor.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptors */ "./node_modules/core-js/modules/es.object.get-own-property-descriptors.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.string.search */ "./node_modules/core-js/modules/es.string.search.js");
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.string.trim */ "./node_modules/core-js/modules/es.string.trim.js");
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var _components_AppBanner_vue__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! ../components/AppBanner.vue */ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue");
/* harmony import */ var _pages_modelsquare_square_components_ModelCondition_vue__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! ~/pages/modelsquare/square/components/ModelCondition.vue */ "./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue");
/* harmony import */ var _pages_modelsquare_square_components_ModelFilters_vue__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! ~/pages/modelsquare/square/components/ModelFilters.vue */ "./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue");
/* harmony import */ var _pages_modelsquare_square_components_ModelList_vue__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! ~/pages/modelsquare/square/components/ModelList.vue */ "./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");

















function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15___default()(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  data: function data() {
    return {
      condition: {
        q: '',
        tab: '1',
        sort: '',
        onlyRecommend: false,
        engine: '',
        label: '',
        page: 1,
        pageSize: 30
      },
      pageSizes: [30, 50]
    };
  },
  components: {
    AppBanner: _components_AppBanner_vue__WEBPACK_IMPORTED_MODULE_16__["default"],
    ModelFilters: _pages_modelsquare_square_components_ModelFilters_vue__WEBPACK_IMPORTED_MODULE_18__["default"],
    ModelCondition: _pages_modelsquare_square_components_ModelCondition_vue__WEBPACK_IMPORTED_MODULE_17__["default"],
    ModelList: _pages_modelsquare_square_components_ModelList_vue__WEBPACK_IMPORTED_MODULE_19__["default"]
  },
  methods: {
    conditionChange: function conditionChange() {
      var params = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};
      this.condition = _objectSpread(_objectSpread({}, this.condition), params);

      if (!params.changePage) {
        this.condition.page = 1;
      }

      window.location.href = "/explore/models?" + "q=".concat(encodeURIComponent(this.condition.q.trim())) + "&tab=".concat(encodeURIComponent(this.condition.tab)) + "&sort=".concat(encodeURIComponent(this.condition.sort)) + "&onlyRecommend=".concat(encodeURIComponent(this.condition.onlyRecommend)) + "&engine=".concat(encodeURIComponent(this.condition.engine)) + "&label=".concat(encodeURIComponent(this.condition.label)) + "&page=".concat(encodeURIComponent(this.condition.page)) + "&pageSize=".concat(encodeURIComponent(this.condition.pageSize));
    }
  },
  beforeMount: function beforeMount() {
    var _this = this;

    var urlParams = (0,_utils__WEBPACK_IMPORTED_MODULE_20__.getUrlSearchParams)();
    this.condition.q = urlParams.q || '';
    this.condition.tab = urlParams.tab || '1';
    this.condition.sort = urlParams.sort || '';
    this.condition.onlyRecommend = urlParams.onlyRecommend == 'true' ? true : false;
    this.condition.engine = urlParams.engine || '';
    this.condition.label = urlParams.label || 'ros-hmci-models'; //默认搜索ros-hmci-models标签

    this.condition.page = Number(urlParams.page) || 1;
    this.condition.pageSize = this.pageSizes.indexOf(Number(urlParams.pageSize)) >= 0 ? Number(urlParams.pageSize) : 30;
    this.$nextTick(function () {
      _this.$refs.modelListRef.search();
    });
  },
  mounted: function mounted() {},
  beforeDestroy: function beforeDestroy() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=script&lang=js":
/*!**********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.promise */ "./node_modules/core-js/modules/es.promise.js");
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var _components_AppBanner_vue__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ../components/AppBanner.vue */ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue");
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! axios */ "./node_modules/axios/index.js");
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(axios__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var markdown_it__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! markdown-it */ "./node_modules/markdown-it/index.js");
/* harmony import */ var markdown_it__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(markdown_it__WEBPACK_IMPORTED_MODULE_8__);






//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "ResourceDetail",
  props: {
    resourceName: {
      type: String,
      required: true
    }
  },
  data: function data() {
    return {
      resourceDetails: {},
      htmlContent: ""
    };
  },
  created: function created() {
    var _this = this;

    return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_5___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3___default().mark(function _callee() {
      var resourceName, response, data;
      return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3___default().wrap(function _callee$(_context) {
        while (1) {
          switch (_context.prev = _context.next) {
            case 0:
              resourceName = _this.$route.params.name;
              _context.next = 3;
              return fetch("/fanshuai/ROS-hmci-resource/raw/branch/master/resource/".concat(resourceName, ".json"));

            case 3:
              response = _context.sent;
              _context.next = 6;
              return response.json();

            case 6:
              data = _context.sent;
              _this.resourceDetails = data[0]; // 假设响应中的第一个对象包含所需的详细信息

            case 8:
            case "end":
              return _context.stop();
          }
        }
      }, _callee);
    }))();
  },
  components: {
    AppBanner: _components_AppBanner_vue__WEBPACK_IMPORTED_MODULE_6__["default"]
  },
  methods: {},
  mounted: function mounted() {
    var _this2 = this;

    return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_5___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3___default().mark(function _callee2() {
      var resourceName, response, md;
      return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3___default().wrap(function _callee2$(_context2) {
        while (1) {
          switch (_context2.prev = _context2.next) {
            case 0:
              _context2.prev = 0;
              resourceName = _this2.$route.params.name;
              _context2.next = 4;
              return axios__WEBPACK_IMPORTED_MODULE_7___default().get("/fanshuai/ROS-hmci-resource/raw/branch/master/markdown/".concat(resourceName, ".md"));

            case 4:
              response = _context2.sent;
              _this2.markdownContent = response.data;
              md = new (markdown_it__WEBPACK_IMPORTED_MODULE_8___default())();
              _this2.htmlContent = md.render(_this2.markdownContent);
              _context2.next = 13;
              break;

            case 10:
              _context2.prev = 10;
              _context2.t0 = _context2["catch"](0);
              console.error(_context2.t0);

            case 13:
            case "end":
              return _context2.stop();
          }
        }
      }, _callee2, null, [[0, 10]]);
    }))();
  }
});

/***/ }),

/***/ "./web_src/vuepages/apis/modules/common.js":
/*!*************************************************!*\
  !*** ./web_src/vuepages/apis/modules/common.js ***!
  \*************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   commonFormPost: function() { return /* binding */ commonFormPost; },
/* harmony export */   createRepo: function() { return /* binding */ createRepo; },
/* harmony export */   forkRepo: function() { return /* binding */ forkRepo; },
/* harmony export */   getAction: function() { return /* binding */ getAction; },
/* harmony export */   getAiCenterList: function() { return /* binding */ getAiCenterList; },
/* harmony export */   getComfyuiUrl: function() { return /* binding */ getComfyuiUrl; },
/* harmony export */   getHeatMap: function() { return /* binding */ getHeatMap; },
/* harmony export */   getMarkdownHtml: function() { return /* binding */ getMarkdownHtml; },
/* harmony export */   getMlopsRight: function() { return /* binding */ getMlopsRight; },
/* harmony export */   getModelExperience: function() { return /* binding */ getModelExperience; },
/* harmony export */   getModelFileSDKCode: function() { return /* binding */ getModelFileSDKCode; },
/* harmony export */   getPointAccountInfo: function() { return /* binding */ getPointAccountInfo; },
/* harmony export */   getPromoteData: function() { return /* binding */ getPromoteData; },
/* harmony export */   getResQueueCode: function() { return /* binding */ getResQueueCode; },
/* harmony export */   getSDKCode: function() { return /* binding */ getSDKCode; },
/* harmony export */   getStaticFile: function() { return /* binding */ getStaticFile; },
/* harmony export */   getUserLastestAiTaskRepoInfo: function() { return /* binding */ getUserLastestAiTaskRepoInfo; },
/* harmony export */   getUserRepoList: function() { return /* binding */ getUserRepoList; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.iterator */ "./node_modules/core-js/modules/es.array.iterator.js");
/* harmony import */ var core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.string.iterator */ "./node_modules/core-js/modules/es.string.iterator.js");
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/web.dom-collections.iterator */ "./node_modules/core-js/modules/web.dom-collections.iterator.js");
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_web_url__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/web.url */ "./node_modules/core-js/modules/web.url.js");
/* harmony import */ var core_js_modules_web_url__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_url__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var _service__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ../service */ "./web_src/vuepages/apis/service.js");
/* harmony import */ var qs__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! qs */ "./node_modules/qs/lib/index.js");
/* harmony import */ var qs__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(qs__WEBPACK_IMPORTED_MODULE_13__);













 // 获取静态文件内容

var getStaticFile = function getStaticFile(filePathName) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: "".concat(filePathName),
    method: 'get',
    params: {}
  });
}; // 获取promote配置数据

var getPromoteData = function getPromoteData(filePathName) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: '/dashboard/invitation',
    method: 'get',
    params: {
      filename: filePathName
    }
  });
}; // 获取模型体验运行的云脑任务列表

var getModelExperience = function getModelExperience() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: '/extension/modelexperience/queryall',
    method: 'get',
    params: {}
  });
}; // 获取markdown渲染结果

var getMarkdownHtml = function getMarkdownHtml(str, mode) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: '/api/v1/markdown',
    method: 'post',
    parmas: {},
    data: {
      mode: mode || 'gfm',
      text: str
    }
  });
}; // 获取个人积分信息
// return {pointAccount:{id,account_code,balance,total_earned,total_consumed,status,version,created_unix,updated_unix },cloudBrainPaySwitch}

var getPointAccountInfo = function getPointAccountInfo() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: '/api/v1/user/point_account',
    method: 'get',
    params: {}
  });
}; // 由模型名称列表或数据集名称列表或运行参数列表获取SDK的代码使用方式
// dataset_name-string[], datasetNameList
// pretrain_model_name -string[] modelNameList
// param_key -string[] param_key
// job_type,compute_source,cluster_type

var getSDKCode = function getSDKCode(params) {
  var datasetNames = params.dataset_name || [];
  var modelNames = params.pretrain_model_name || [];
  var parameterKeys = params.param_key || [];
  var searchParams = new URLSearchParams();
  datasetNames.forEach(function (name) {
    searchParams.append('dataset_name', name);
  });
  modelNames.forEach(function (name) {
    searchParams.append('pretrain_model_name', name);
  });
  parameterKeys.forEach(function (name) {
    searchParams.append('param_key', name);
  });
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: "/api/v1/ai_task/generate_sdk_code?".concat(searchParams.toString()),
    method: 'get',
    params: {
      job_type: params.job_type,
      compute_source: params.compute_source,
      cluster_type: params.cluster_type,
      visualize_required: params.visualize_required
    },
    data: {}
  });
}; // 获取单个模型文件的SDK code

var getModelFileSDKCode = function getModelFileSDKCode(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: "/api/v1/".concat(params.owner, "/").concat(params.repo, "/sdk/generate_model_download_code"),
    method: 'get',
    params: {
      model_name: params.name,
      model_file_name: params.filename
    },
    data: {}
  });
}; // 查询智算列表

var getAiCenterList = function getAiCenterList() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: "/resources/queue/centers",
    method: 'get',
    params: {}
  });
}; // 查询所有资源队列名称列表

var getResQueueCode = function getResQueueCode(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: "/explore/card_request/resources/queue/codes",
    method: 'get',
    params: params
  });
}; // common form post

var commonFormPost = function commonFormPost(url, data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: url,
    method: 'post',
    params: {},
    data: qs__WEBPACK_IMPORTED_MODULE_13___default().stringify(data)
  });
}; // 获取用户最近AI任务的项目仓信息

var getUserLastestAiTaskRepoInfo = function getUserLastestAiTaskRepoInfo() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: "/api/v1/user/get_latest_cloudbrain_repo",
    method: 'get',
    params: {}
  });
}; // 用户有权限的仓库列表
// uid,sort-updated,order-desc,asc,type-cloudbrain,model,dataset

var getUserRepoList = function getUserRepoList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: "/api/v1/repos/search_for_ai_task",
    method: 'get',
    params: {
      uid: params.uid,
      sort: params.sort || 'updated',
      order: params.order || 'desc',
      type: params.type || 'cloudbrain'
    }
  });
}; // 新建代码仓
// auto_init-true,default_branch-master,name,private-true

var createRepo = function createRepo(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: "/api/v1/user/repos",
    method: 'post',
    data: {
      auto_init: data.auto_init == undefined ? true : data.auto_init,
      default_branch: data.default_branch || 'master',
      name: data.name,
      "private": data["private"] == undefined ? true : data["private"]
    },
    params: {}
  });
}; // fork代码仓
// owner,repo,repo_name

var forkRepo = function forkRepo(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: "/api/v1/repos/".concat(data.owner, "/").concat(data.repo, "/forks"),
    method: 'post',
    data: {
      owner: data.owner,
      repo: data.repo,
      repo_name: data.repo_name
    },
    params: {}
  });
}; // 获取comyui路由

var getComfyuiUrl = function getComfyuiUrl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: '/api/v1/comfyui/experience/online_url',
    method: 'get',
    params: params
  });
}; // 获取 Mlops 权限

var getMlopsRight = function getMlopsRight() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: '/api/v1/monitor/mlops_right',
    method: 'get',
    params: {}
  });
}; // 热力图

var getHeatMap = function getHeatMap(userName) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: "/api/v1/users/".concat(userName, "/heatmap"),
    method: 'get',
    params: {}
  });
}; // 用户action 

var getAction = function getAction() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_12__["default"])({
    url: '/api/v1/platform/action',
    method: 'get',
    params: {}
  });
};

/***/ }),

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

/***/ }),

/***/ "./web_src/vuepages/apis/modules/modelsquare.js":
/*!******************************************************!*\
  !*** ./web_src/vuepages/apis/modules/modelsquare.js ***!
  \******************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   getModelList: function() { return /* binding */ getModelList; },
/* harmony export */   getModelRelatedInfo: function() { return /* binding */ getModelRelatedInfo; },
/* harmony export */   getModelSqaureFilters: function() { return /* binding */ getModelSqaureFilters; },
/* harmony export */   setModelFav: function() { return /* binding */ setModelFav; }
/* harmony export */ });
/* harmony import */ var _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @babel/runtime/helpers/objectWithoutProperties */ "./node_modules/@babel/runtime/helpers/objectWithoutProperties.js");
/* harmony import */ var _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var _service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../service */ "./web_src/vuepages/apis/service.js");
/* harmony import */ var qs__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! qs */ "./node_modules/qs/lib/index.js");
/* harmony import */ var qs__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(qs__WEBPACK_IMPORTED_MODULE_2__);


 // 获取模型广场列表

var getModelList = function getModelList(_ref) {
  var url = _ref.url,
      params = _babel_runtime_helpers_objectWithoutProperties__WEBPACK_IMPORTED_MODULE_0___default()(_ref, ["url"]);

  return (0,_service__WEBPACK_IMPORTED_MODULE_1__["default"])({
    url: "/api/v1/aimodel/list".concat(url),
    method: "get",
    params: params
  });
}; // 模型详情侧边栏相关信息

var getModelRelatedInfo = function getModelRelatedInfo(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_1__["default"])({
    url: "/api/v1/aimodel/related",
    method: "get",
    params: params
  });
}; // 获取模型广场筛选项

var getModelSqaureFilters = function getModelSqaureFilters() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_1__["default"])({
    url: '/modelsquare/main_query_label',
    method: 'get',
    params: {}
  });
}; // 模型收藏/取消收藏
// data: id, collected-为true表示收藏此模型，为false表示取消收藏此模型

var setModelFav = function setModelFav(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_1__["default"])({
    url: '/modelsquare/modify_model_collect',
    method: 'put',
    data: qs__WEBPACK_IMPORTED_MODULE_2___default().stringify(data)
  });
};

/***/ }),

/***/ "./web_src/vuepages/apis/modules/repos.js":
/*!************************************************!*\
  !*** ./web_src/vuepages/apis/modules/repos.js ***!
  \************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   deleteRepos: function() { return /* binding */ deleteRepos; },
/* harmony export */   followingUsers: function() { return /* binding */ followingUsers; },
/* harmony export */   getActiveOrgs: function() { return /* binding */ getActiveOrgs; },
/* harmony export */   getActiveUsers: function() { return /* binding */ getActiveUsers; },
/* harmony export */   getHomePageData: function() { return /* binding */ getHomePageData; },
/* harmony export */   getNewRepoList: function() { return /* binding */ getNewRepoList; },
/* harmony export */   getRepoBranch: function() { return /* binding */ getRepoBranch; },
/* harmony export */   getRepoList: function() { return /* binding */ getRepoList; },
/* harmony export */   getReposCollectListData: function() { return /* binding */ getReposCollectListData; },
/* harmony export */   getReposListData: function() { return /* binding */ getReposListData; },
/* harmony export */   getReposSquareTabData: function() { return /* binding */ getReposSquareTabData; },
/* harmony export */   pinnedRepos: function() { return /* binding */ pinnedRepos; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _service__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../service */ "./web_src/vuepages/apis/service.js");



 // 获取首页数据

var getHomePageData = function getHomePageData() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: '/recommend/home',
    method: 'get',
    params: {}
  });
}; // 获取项目广场上方tab数据 tab=preferred 项目优选|incubation 启智孵化管道|hot-paper 热门论文项目

var getReposSquareTabData = function getReposSquareTabData(tab) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: '/explore/repos/square/tab',
    method: 'get',
    params: {
      type: tab
    }
  });
}; // 搜索项目
// q	string	否	关键词
// topic	string	否	标签名
// sort	string	是	mostpopular 近期热门 | mostactive 近期活跃 | recentupdate 最近更新 | newest 最近创建 
//                  moststars 点赞最多 | mostforks 派生最多 | mostdatasets 数据集最多 | mostaitasks AI任务最多 | mostmodels 模型最多                
// pageSize	int	是	每页大小，可选值为15 | 30 | 50
// page	int	是	页码

var getReposListData = function getReposListData(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: '/explore/repos/search',
    method: 'get',
    params: {
      q: params.q || '',
      topic: params.topic || '',
      sort: params.sort || 'mostpopular',
      pageSize: params.pageSize || 15,
      page: params.page || 1
    }
  });
};
var getReposCollectListData = function getReposCollectListData(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: 'api/v1/user/subscriptions',
    method: 'get',
    params: params
  });
}; // 获取活跃用户列表

var getActiveUsers = function getActiveUsers() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: '/explore/repos/square/active-user',
    method: 'get',
    params: {}
  });
}; // 关注用户

var followingUsers = function followingUsers(userName, isFollowing) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: "/api/v1/user/following/".concat(userName),
    method: isFollowing ? 'put' : 'delete',
    params: {}
  });
}; // 获取活跃组织列表

var getActiveOrgs = function getActiveOrgs() {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: '/explore/repos/square/active-org',
    method: 'get',
    params: {}
  });
}; // 获取项目列表

var getRepoList = function getRepoList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: '/api/v1/repos/search',
    method: 'get',
    params: params
  });
}; // 获取项目列表新接口

var getNewRepoList = function getNewRepoList(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: '/api/v1/repos/platform/search',
    method: 'get',
    params: params
  });
}; // 获取项目分支

var getRepoBranch = function getRepoBranch(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: "/api/v1/".concat(params.repoOwnerName, "/").concat(params.repoName, "/ai_task/repo_branch"),
    method: 'get',
    params: {}
  });
}; // 关注用户

var deleteRepos = function deleteRepos(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: "/api/v1/repos/".concat(params.ownerName, "/").concat(params.name, "/project"),
    method: 'delete',
    params: {}
  });
}; // 关注用户

var pinnedRepos = function pinnedRepos(userName, reposName, isPinning) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_3__["default"])({
    url: "/api/v1/repos/".concat(userName, "/").concat(reposName, "/project/pinned"),
    method: !isPinning ? 'post' : 'delete',
    params: {}
  });
};

/***/ }),

/***/ "./web_src/vuepages/apis/service.js":
/*!******************************************!*\
  !*** ./web_src/vuepages/apis/service.js ***!
  \******************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.object.assign */ "./node_modules/core-js/modules/es.object.assign.js");
/* harmony import */ var core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_assign__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.promise */ "./node_modules/core-js/modules/es.promise.js");
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! axios */ "./node_modules/axios/index.js");
/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(axios__WEBPACK_IMPORTED_MODULE_3__);




var service = axios__WEBPACK_IMPORTED_MODULE_3___default().create({
  baseURL: '/'
});
service.interceptors.request.use(function (config) {
  config.data && Object.assign(config.data, {
    _csrf: window.config ? window.config.csrf : ''
  });
  config.params && Object.assign(config.params, {
    _csrf: window.config ? window.config.csrf : ''
  });
  return config;
}, function (error) {
  return Promise.reject(error);
});
service.interceptors.response.use(function (response) {
  if (response.status == 200 && response.data && response.data.code == 9002) {
    // 绑定微信
    window.location.href = "/authentication/wechat/bind?redirect_to=".concat(encodeURIComponent(window.location.href));
    return Promise.reject(response);
  }

  return response;
}, function (error) {
  return Promise.reject(error);
});
/* harmony default export */ __webpack_exports__["default"] = (service);

/***/ }),

/***/ "./web_src/vuepages/const/index.js":
/*!*****************************************!*\
  !*** ./web_src/vuepages/const/index.js ***!
  \*****************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   ACC_CARD_TYPE: function() { return /* binding */ ACC_CARD_TYPE; },
/* harmony export */   AI_CENTER: function() { return /* binding */ AI_CENTER; },
/* harmony export */   BenchmarkTypeList: function() { return /* binding */ BenchmarkTypeList; },
/* harmony export */   CLUSTERS: function() { return /* binding */ CLUSTERS; },
/* harmony export */   COMPUTER_RESOURCES: function() { return /* binding */ COMPUTER_RESOURCES; },
/* harmony export */   COMPUTER_RESOURCES_COLORS: function() { return /* binding */ COMPUTER_RESOURCES_COLORS; },
/* harmony export */   CONSUME_STATUS: function() { return /* binding */ CONSUME_STATUS; },
/* harmony export */   JOB_TYPE: function() { return /* binding */ JOB_TYPE; },
/* harmony export */   MODEL_ENGINES: function() { return /* binding */ MODEL_ENGINES; },
/* harmony export */   NETWORK_TYPE: function() { return /* binding */ NETWORK_TYPE; },
/* harmony export */   NETWORK_TYPE_VALUE: function() { return /* binding */ NETWORK_TYPE_VALUE; },
/* harmony export */   NEW_JOB_TYPE: function() { return /* binding */ NEW_JOB_TYPE; },
/* harmony export */   NEW_JOB_TYPE_OBJ: function() { return /* binding */ NEW_JOB_TYPE_OBJ; },
/* harmony export */   OPERATION_TYPE: function() { return /* binding */ OPERATION_TYPE; },
/* harmony export */   POINT_ACTIONS: function() { return /* binding */ POINT_ACTIONS; },
/* harmony export */   SOURCE_TYPE: function() { return /* binding */ SOURCE_TYPE; },
/* harmony export */   SPECIFICATION_STATUS: function() { return /* binding */ SPECIFICATION_STATUS; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");


var SOURCE_TYPE = [{
  k: 'ACCOMPLISH_TASK',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('accomplishTask')
}, {
  k: 'ADMIN_OPERATE',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('adminOperate')
}, {
  k: 'RUN_CLOUDBRAIN_TASK',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('runCloudBrainTask')
}];
var CONSUME_STATUS = [{
  k: 'OPERATING',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('operating')
}, {
  k: 'SUCCEEDED',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('succeeded')
}];
var POINT_ACTIONS = [{
  k: 'CreatePublicRepo',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('createPublicProject')
}, {
  k: 'CreateIssue',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('dailyPutforwardTasks')
}, {
  k: 'CreatePullRequest',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('dailyPR')
}, {
  k: 'CommentIssue',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('comment')
}, {
  k: 'UploadAttachment',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('uploadDatasetFile')
}, {
  k: 'CreateNewModelTask',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('importNewModel')
}, {
  k: 'BindWechat',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('completeWechatCodeScanningVerification')
}, {
  k: 'CreateCloudbrainTask',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('dailyRunCloudbrainTasks')
}, {
  k: 'DatasetRecommended',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('datasetRecommendedByThePlatform')
}, {
  k: 'CreateImage',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('submitNewPublicImage')
}, {
  k: 'ImageRecommend',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('imageRecommendedByThePlatform')
}, {
  k: 'ChangeUserAvatar',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('firstChangeofAvatar')
}, {
  k: 'PushCommits',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('dailyCommit')
}, {
  k: 'TaskInviteFriendRegister',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('user.inviteFriends')
}, {
  k: 'TaskCreateDataset',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('dailyCreateDataset')
}, {
  k: 'TaskCreateAimodel',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('dailyCreateAimodel1')
}, {
  k: 'TaskAimodelRecommended',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('aimodelRecommendedByThePlatform')
}];
var JOB_TYPE = [{
  k: 'DEBUG',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('debugTask'),
  train_type: 'Notebook',
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('debugTask')
}, {
  k: 'TRAIN',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('trainTask'),
  train_type: 'TrainJob',
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('trainTask')
}, {
  k: 'INFERENCE',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('inferenceTask'),
  train_type: 'TrainJob',
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('inferenceTask')
}, {
  k: 'BENCHMARK',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('benchmarkTask'),
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('benchmarkTask')
}, {
  k: 'ONLINEINFERENCE',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('onlineinfer'),
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('onlineinferTask'),
  train_type: 'Notebook'
}, {
  k: 'HPC',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('superComputeTask'),
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('superComputeTask')
}, {
  k: 'GENERAL',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('generalTask'),
  train_type: 'Notebook',
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('generalTask')
}, {
  k: 'MODELEXPERIENCE',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelManage.onlineInference'),
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelManage.onlineInferenceTask')
}, {
  k: 'FINETUNE',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.sftFinetune'),
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.sftFinetuneTask')
}, {
  k: 'SDFINETUNE',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.sdModelFinetuen'),
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.sdModelFinetuenTask')
}, {
  k: 'ComfyuiExperience',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.cvComfyui'),
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.cvComfyuiTask')
}, {
  k: 'EVAL',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.modelPerEvaluate'),
  alias: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.modelEvaluateTask')
}];
var BenchmarkTypeList = ['BENCHMARK', 'SIM2BRAIN_SNN', 'SNN4ECOSET', 'SNN4IMAGENET', 'BRAINSCORE', 'MODELSAFETY']; // 资源管理

var CLUSTERS = [{
  k: 'OpenI',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.OpenI')
}, {
  k: 'C2Net',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.C2Net')
}, {
  k: 'IFLYTEKTraining',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.IFLYTEKTraining')
}];
var AI_CENTER = [{
  k: 'OpenIOne',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.OpenIOne')
}, {
  k: 'OpenITwo',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.OpenITwo')
}, {
  k: 'OpenIChengdu',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.OpenIChengdu')
}, {
  k: 'pclcci',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.pclcci')
}, {
  k: 'hefei',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.hefeiCenter')
}, {
  k: 'xuchang',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.xuchangCenter')
}, {
  k: 'OpenI-huoshi',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.huoshi')
}];
var COMPUTER_RESOURCES = [{
  k: 'CPU',
  v: 'CPU'
}, {
  k: 'GPU',
  v: 'GPU'
}, {
  k: 'NPU',
  v: 'NPU'
}, {
  k: 'GCU',
  v: 'GCU'
}, {
  k: 'MLU',
  v: 'MLU'
}, {
  k: 'DCU',
  v: 'DCU'
}, {
  k: 'ILUVATAR-GPGPU',
  v: 'ILUVATAR-GPGPU'
}, {
  k: 'METAX-GPGPU',
  v: 'METAX-GPGPU'
}, {
  k: 'BIREN-GPU',
  v: 'BIREN-GPU'
}];
var COMPUTER_RESOURCES_COLORS = {
  'GPU': '#4fb62f',
  'NPU': '#c31d20',
  'GCU': '#e73828',
  'DCU': '#b01f24',
  'MLU': '#0077ed',
  'ILUVATAR-GPGPU': '#0038bd',
  'METAX-GPGPU': '#5c246a',
  'BIREN-GPU': '#50c878'
};
var ACC_CARD_TYPE = [{
  k: 'T4',
  v: 'T4'
}, {
  k: 'A100',
  v: 'A100'
}, {
  k: 'V100',
  v: 'V100'
}, {
  k: 'ASCEND910',
  v: 'Ascend 910'
}, {
  k: 'ASCEND-D910B',
  v: 'Ascend 910B'
}, {
  k: 'ASCEND-910C',
  v: 'Ascend 910C'
}, {
  k: 'MLU270',
  v: 'MLU270'
}, {
  k: 'MLU290',
  v: 'MLU290'
}, {
  k: 'RTX3080',
  v: 'RTX3080'
}, {
  k: '3090',
  v: '3090'
}, {
  k: '4090',
  v: '4090'
}, {
  k: 'ENFLAME-T20',
  v: 'ENFLAME-T20'
}, {
  k: 'ENFLAME-I20',
  v: 'ENFLAME-I20'
}, {
  k: 'DCU',
  v: 'DCU'
}, {
  k: 'Z100L',
  v: 'Z100L'
}, {
  k: 'K100_AI',
  v: 'K100_AI'
}, {
  k: 'BI-V100',
  v: 'BI-V100'
}, {
  k: 'MR-V100',
  v: 'MR-V100'
}, {
  k: 'N100',
  v: 'N100'
}, {
  k: 'N260',
  v: 'N260'
}, {
  k: 'C500',
  v: 'C500'
}, {
  k: 'L20',
  v: 'L20'
}, {
  k: 'S60',
  v: 'S60'
}, {
  k: 'BIREN106M',
  v: 'BIREN106M'
}, {
  k: 'BW1000',
  v: 'BW1000'
}];
var SPECIFICATION_STATUS = [{
  k: '1',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.willOnShelf')
}, {
  k: '2',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.onShelf')
}, {
  k: '3',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('resourcesManagement.offShelf')
}];
var NETWORK_TYPE = [{
  k: 1,
  v: "".concat(_langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('cloudbrainObj.networkType'), "(").concat(_langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('cloudbrainObj.noInternet'), ")")
}, {
  k: 2,
  v: "".concat(_langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('cloudbrainObj.networkType'), "(").concat(_langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('cloudbrainObj.hasInternet'), ")")
}];
var NETWORK_TYPE_VALUE = [{
  k: 1,
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('cloudbrainObj.noInternet')
}, {
  k: 2,
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('cloudbrainObj.hasInternet')
}];
var OPERATION_TYPE = [{
  k: 'edit',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('logManagement.resourceSpecificationEdit')
}, {
  k: 'on-shelf',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('logManagement.resourceSpecificationOnshelf')
}, {
  k: 'off-shelf',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('logManagement.resourceSpecificationOffshelf')
}, {
  k: 'create',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('logManagement.resourceSpecificationAddition')
}, {
  k: 'auto-update',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('logManagement.resourceSpecificationAutoUpdate')
}]; // 模型

var MODEL_ENGINES = [{
  k: 0,
  v: 'PyTorch'
}, {
  k: 1,
  v: 'TensorFlow'
}, {
  k: 2,
  v: 'MindSpore'
}, {
  k: 4,
  v: 'PaddlePaddle'
}, {
  k: 5,
  v: 'OneFlow'
}, {
  k: 6,
  v: 'MXNet'
}, {
  k: 3,
  v: 'Other'
}];
var NEW_JOB_TYPE = [{
  k: 'DEBUG',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('debugTask')
}, {
  k: 'TRAIN',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('trainTask')
}, {
  k: 'ONLINEINFERENCE',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('onlineinfer')
}, {
  k: 'HPC',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('superComputeTask')
}, {
  k: 'GENERAL',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('generalTask')
}, {
  k: 'MODELEXPERIENCE',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelManage.onlineInference')
}, {
  k: 'FINETUNE',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.sftFinetune')
}, {
  k: 'SDFINETUNE',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.sdModelFinetuen')
}, {
  k: 'ComfyuiExperience',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.cvComfyui')
}, {
  k: 'EVAL',
  v: _langs__WEBPACK_IMPORTED_MODULE_1__.i18n.t('modelSquare.modelPerEvaluate')
}];
var NEW_JOB_TYPE_OBJ = {
  'DEBUG': [],
  'TRAIN': [],
  'ONLINEINFERENCE': [],
  'HPC': [],
  'GENERAL': [],
  'MODELEXPERIENCE': [],
  'FINETUNE': [],
  'SDFINETUNE': [],
  "ComfyuiExperience": [],
  "EVAL": []
};

/***/ }),

/***/ "./web_src/vuepages/langs/config/en-US.js":
/*!************************************************!*\
  !*** ./web_src/vuepages/langs/config/en-US.js ***!
  \************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0__);


var _notebook, _modelObj;

var en = {
  loading: 'Loading...',
  noData: 'No Data',
  noMessage: "No Message",
  noDataset: 'Empty, with nothing at all.',
  date: 'Date',
  confirm: 'Confirm',
  cancel: 'Cancel',
  confirm1: 'Confirm',
  confirmOp: "Confirm Operate",
  cancelOp: "Cancel Operate",
  enterReason: "Please enter the reason",
  clear: "Clear",
  none: "None",
  pleaseCompleteTheInformationFirst: 'Please Complete the Information first!',
  submittedSuccessfully: 'Submitted Successfully!',
  submittedFailed: 'Submitted Failed!',
  operation: 'Operation',
  edit: 'Edit',
  "delete": 'Delete',
  tips: 'Tips',
  expandMore: 'Expand More',
  goBack: 'Go back',
  star: 'Star',
  unStar: 'UnStar',
  submit: 'Submit',
  selectAll: 'Select All',
  selectNone: 'Select None',
  copy: 'Copy',
  copySuccess: 'Copied successfully',
  copyLink: 'Copy Link',
  copyedLink: 'Link has been copied',
  downloadAddress: 'Download address',
  operationFailed: 'Operation failed',
  cancelOperate: 'You have canceled the operation',
  accomplishTask: 'Accomplish Task',
  adminOperate: 'Administrator Operation',
  runCloudBrainTask: 'Run Computing Task',
  operating: 'Operating',
  succeeded: 'Succeeded',
  debugTask: 'Debug Task',
  trainTask: 'Train Task',
  trainDuration: "Train task duration",
  inferenceTask: 'Inference Task',
  benchmarkTask: 'Benchmark Task',
  onlineinfer: "Online Inference",
  onlineinferTask: "Online Inference Task",
  consoleHome: "Console Home",
  superComputeTask: "HPC Task",
  generalTask: "General Task",
  createPublicProject: 'Create Public Projects',
  dailyPutforwardTasks: 'Daily Put Forward Tasks',
  dailyPR: 'Daily PR',
  comment: 'Comment',
  uploadDatasetFile: 'Upload Dataset Files',
  importNewModel: 'Import New Models',
  completeWechatCodeScanningVerification: 'Complete Wechat Code Scanning Verification',
  dailyRunCloudbrainTasks: 'Daily Run Computing Tasks',
  datasetRecommendedByThePlatform: 'Dataset Recommended by the Platform',
  aimodelRecommendedByThePlatform: "Model Recommended by the Platform",
  dailyCreateDataset: "Create Datasets",
  dailyCreateDataset1: "created dataset",
  dailyCreateAimodel: "Create Models",
  dailyCreateAimodel1: "created model",
  submitNewPublicImage: 'Submit New Public Images',
  imageRecommendedByThePlatform: 'Image Recommended by the Platform',
  firstChangeofAvatar: 'First Change of Avatar',
  dailyCommit: 'Daily Commit',
  calcPointDetails: 'Calculation Points Details',
  calcPointAcquisitionInstructions: 'Calculation Points Acquisition Instructions',
  CurrAvailableCalcPoints: 'Currently Available Calculation Points',
  totalGainCalcPoints: 'Total Gain of Calculation Points',
  totalConsumeCalcPoints: 'Total Consume of Calculation Points',
  gainDetail: 'Gain Detail',
  consumeDetail: 'Consume Detail',
  serialNumber: 'Serial Number',
  time: 'Time',
  scene: 'Scene',
  behaviorOfPoint: 'Behavior Of Point',
  explanation: 'Explanation',
  points: 'Points',
  status: 'Status',
  runTime: 'Run Time',
  taskName: 'Task Name',
  createdRepository: 'created repository ',
  repositoryWasDel: 'repository was deleted',
  userWasDel: 'user was deleted',
  userAccountWasDel: "user account was deleted",
  openedIssue: 'opened issue ',
  createdPullRequest: 'created pull request ',
  commentedOnIssue: 'commented on issue ',
  uploadDataset: 'upload dataset ',
  createdNewModel: 'created new model ',
  invitedFriend: 'invited friend ',
  firstBindingWechatRewards: 'first binding wechat rewards',
  created: 'created ',
  type: ' type ',
  dataset: 'Datasets ',
  setAsRecommendedDataset: ' was set as recommended dataset',
  setAsRecommendedAimodel: ' was set as recommended model',
  committedImage: 'committed image ',
  image: 'image ',
  setAsRecommendedImage: ' was set as recommended image',
  updatedAvatar: 'updated avatar',
  pushedBranch: 'pushed to {branch} at ',
  deleteBranch: 'deleted branch {branch} from {repo}',
  pushedTag: ' pushed tag {tag} to ',
  deleteTag: ' deleted tag {tag} from {repo}',
  dailyMaxTips: "can't get full points when reach the daily upper limit",
  memory: 'Memory',
  sharedMemory: 'Shared Memory',
  ';': ', ',
  noPointGainRecord: 'No Point Earn Record Yet',
  noPointConsumeRecord: 'No Point Consume Record Yet',
  consumeDescr: 'Explanation: Starting from late October 2024, the billing cycle for Computing tasks will be adjusted. For specific billing rules, please refer to the "Billing Rules" section on the "<a target="_blank" href="https://openi.pcl.ac.cn/docs/index.html#/quickstart/resources?id=%e8%ae%a1%e8%b4%b9%e8%a7%84%e5%88%99">Platform Computing Resources and Points Explanation</a>" page in the online help.',
  emptyPage: 'Request forbidden by administrative rules',
  emptyPageDescr: 'The page you are trying to reach either <strong>does not exist</strong> or <strong>you are not authorized</strong> to view it.',
  handleTask: 'Processing tasks',
  freeCompute: 'Inclusive computing power',
  repoPathTips: 'The path only allows letters, numbers, and - _, Up to 100 characters.',
  useInPcWeb: 'For a better user experience, please open this page on your PC.',
  acknowledgementsTips: {
    main1: 'If the OpenI Community has provided assistance to your research work, please acknowledge it in your paper (',
    main2: ').',
    moreContent: 'more content',
    p1: "English version: Thanks for the support provided by OpenI Community (<a href=\"https://openi.pcl.ac.cn\" target=\"_blank\">https://openi.pcl.ac.cn</a>).",
    p2: "Chinese version: \u611F\u8C22\u542F\u667A\u793E\u533A\u63D0\u4F9B\u7684\u6280\u672F\u652F\u6301(<a href=\"https://openi.pcl.ac.cn\" target=\"_blank\">https://openi.pcl.ac.cn</a>)\u3002",
    p3: "If your work references this platform, you are also welcome to submit your work information in the following open source projects:",
    p4: "<a href=\"https://openi.pcl.ac.cn/OpenIOSSG/references\" target=\"_blank\">https://openi.pcl.ac.cn/OpenIOSSG/references</a>"
  },
  computeResourceTitle: {
    CPU: 'CPU',
    VCPU: 'VCPU',
    'CPU/GPU': 'NVIDIA GPU',
    GPU: 'NVIDIA GPU',
    NPU: 'Ascend NPU',
    GCU: 'Enflame GCU',
    MLU: 'Cambricon MLU',
    DCU: 'HYGON DCU',
    'ILUVATAR-GPGPU': 'Iluvatar CoreX GPGPU',
    'METAX-GPGPU': 'MetaX GPGPU',
    'BIREN-GPU': 'BIREN GPU'
  },
  TaskTypeTitle: {
    Notebook: 'Notebook/Online Inference',
    Notebook1: 'Notebook',
    TrainJob: 'Train Job',
    Ecs: 'Ecs',
    Inference: 'Inference',
    Service: 'Service'
  },
  resourcesManagement: {
    OpenI: 'OpenI',
    C2Net: 'C2Net',
    IFLYTEKTraining: "IFLYTEKTraining",
    OpenIOne: 'OpenI One',
    OpenITwo: 'OpenI Two',
    OpenIChengdu: 'OpenI ChengDu AI Chenter',
    chengduCenter: 'ChengDu AI Center',
    pclcci: 'PCL Cloud Computer Institute',
    hefeiCenter: 'HeFei AI Center',
    xuchangCenter: 'XuChang AI Center',
    huoshi: "HuoShi AI Center",
    willOnShelf: 'To Be On Shelf',
    onShelf: 'On Shelf',
    offShelf: 'Off Shelf',
    toOnShelf: 'To On Shelf',
    toOffShelf: 'To Off Shelf',
    toSetPriceAndOnShelf: 'To Set Price and On Shelf',
    status: 'Status',
    reason: "Reason",
    allStatus: 'All Status',
    syncAiNetwork: 'Sync AI Network',
    resQueue: 'Resources Queue',
    allResQueue: 'All Resources Queues',
    addResQueue: 'Add Resources Queue',
    addResQueueBtn: 'Add Resources Queue',
    editResQueue: 'Edit Resources Queue',
    resQueueCode: 'Resources Queue Code',
    resQueueName: 'Resources Queue Name',
    resQueueType: "Resources Queue Type",
    resQueueIsAvailable: "Resources Queue Is Available",
    resQueueIsAvailableAll: "Resources Queue Is Available(All)",
    allResQueueType: "All Resources Queue Type",
    whichCluster: 'Cluster',
    allCluster: 'All Clusters',
    aiCenter: 'AI Center',
    aiCenterID: 'AI Center ID',
    allAiCenter: 'All AI Centers',
    computeResource: 'Compute Resource',
    allComputeResource: 'All Compute Resources',
    accCardType: 'Acc Card Type',
    allAccCardType: 'All Acc Card Type',
    allNetworkType: 'All Network Type',
    cardsTotalNum: 'Cards Total Number',
    accCardsNum: 'Acc Cards Number',
    allCardsNum: 'All Cards Number',
    remark: 'Remark',
    pleaseEnterRemark: 'Please Enter Remark(The maximum length shall not exceed 255)',
    pleaseEnterPositiveIntegerCardsTotalNum: 'Please Enter Positive Integer Cards Total Number!',
    addResSpecificationAndPriceInfo: 'Add Resources Specification and Price Info',
    addResSpecificationBtn: 'Add Resources Specification',
    editResSpecificationAndPriceInfo: 'Edit Resources Specification and Price Info',
    resSpecificationAndPriceManagement: 'Resources Specification Management',
    sourceSpecCode: 'Source Specification Code',
    sourceSpecCodeTips: 'OpenI Two Should Enter the Source Specification Code',
    sourceSpecId: 'Source Specification ID',
    cpuNum: 'CPU Number',
    gpuMem: 'GPU Memory',
    mem: 'Memory',
    shareMem: 'Share Memory',
    unitPrice: 'Unit Price',
    point_hr: 'Point/hr',
    ShelfReason: 'Reason for Shelf',
    node: 'node',
    free: 'Free',
    onShelfConfirm: 'Are you sure to on shelf the resources specification?',
    offShelfConfirm: 'Are you sure to off shelf the resources specification?',
    onShelfCode1001: 'On shelf failed, the resources queues not available.',
    onShelfCode1003: 'On shelf failed, the resources specification not available.',
    offShelfDlgTip1: 'The current resource specification has been used in the following roles:',
    offShelfDlgTip2: 'Please confirm to off shelf?',
    characterLengthPrompt: "Please enter no more than 255 characters",
    resSceneManagement: 'Resources Scene Management',
    addResScene: 'Add Resources Scene',
    addResSceneBtn: 'Add Resources Scene',
    editResScene: 'Edit Resources Scene',
    resSceneName: 'Resources Scene Name',
    jobType: 'Job Type',
    allJobType: 'All Job Types',
    allJobStatus: 'All Job Status',
    sceneType: "Community Scene Type",
    allSceneType: "All Community Scene Type",
    isExclusiveSpec: 'Is Exclusive Spec?',
    allExclusiveAndCommonUseSpec: 'All Exclusive and Common Use Spec',
    "public": "Public",
    exclusive: "Exclusive",
    exclusiveTxt: "Exclusive",
    exclusiveSpec: 'Exclusive Spec',
    commonUseSpec: 'Common Use Spec',
    exclusiveOrg: 'Exclusive Organization',
    exclusiveOrgTips: 'Multiple organization names are separated by semicolons',
    computeCluster: 'Compute Cluster',
    resourceSpecification: 'Resource Specification',
    lastUpdateTime: 'Last Update Time',
    resSceneDeleteConfirm: 'Are you sure to delete the current Resource Scene?',
    resourceSpecificationIsAvailable: 'Specification Is Available',
    resourceSpecificationIsAvailableAll: 'Specification Is Available(All)',
    available: 'Available',
    notAvailable: 'Not Available',
    selectSpec: 'Select resource specification',
    accordingSpec: 'According to resource specification',
    accordingQueue: 'According to resources queue',
    accordingSelect: 'According to the selected option',
    allIsEnableVisualization: 'Is Enable Visualization(All)',
    notEnable: 'Not Enable',
    enable: 'Enable',
    roleNameTips: 'Maximum input of 80 characters',
    roleDescTips: 'Maximum input of 800 characters'
  },
  logManagement: {
    operationLog: "Operation Log",
    allOperationType: "All Operation Types",
    resourceSpecificationOnshelf: 'Resource Specification Onshelf',
    resourceSpecificationEdit: 'Resource Specification Edit',
    resourceSpecificationOffshelf: 'Resource Specification Offshelf',
    resourceSpecificationAddition: 'Resource Specification Addition',
    resourceSpecificationAutoUpdate: 'Resource Specification Auto Update',
    operationType: 'Operation Type',
    operationReason: 'Operation Reason',
    operator: 'Operator',
    operationTime: 'Operation Time'
  },
  user: {
    inviteFriends: 'Invite Friends',
    inviteFriendsTips: 'Copy QR code or invite registration link to share with friends',
    clickToViewTheEventDetails: 'Click to view the event details',
    copyRegistrationInvitationLink: 'Copy registration invitation link',
    registrationAdress: 'Registration Adress: ',
    recommender: 'Recommender: ',
    invitedFriends: 'Invited friends',
    registrationTime: 'Registration time',
    theSharedContentHasBeenCopiedToTheClipboard: 'The shared content has been copied to the clipboard',
    copyError: 'Copy error',
    Activated: 'Activated',
    notActive: 'Not active',
    normal: 'Normal'
  },
  tranformImageFailed: 'Picture desensitization failed',
  originPicture: 'Origin picture',
  desensitizationPicture: 'Desensitization picture',
  desensitizationObject: 'Desensitization object',
  example: 'Example',
  startDesensitization: 'Start desensitization',
  all: 'All',
  others: 'Others',
  onlyFace: 'Only face',
  onlyLicensePlate: 'Only license plate',
  dragThePictureHere: 'Drag the picture here',
  or: ' or ',
  clickUpload: 'Click upload',
  dataDesensitizationModelExperience: 'Data desensitization model experience',
  dataDesensitizationModelDesc: 'Use AI technology to desensitize the face and license plate number in the picture. For more information about this model, please visit the project',
  limitFilesUpload: 'Only jpg/jpeg/png files can be uploaded',
  limitSizeUpload: 'The size of the uploaded file cannot exceed 20M!',
  bootPlaceholder: 'Please enter the startup file',
  notebook: (_notebook = {
    createNewNotebook: "Create new notebook debug task",
    sameTaskTips1: "You have created an",
    sameTaskTips2: "equivalent task",
    sameTaskTips3: "that is waiting or running, please wait for the task to finish before creating it.",
    sameTaskTips4: "You can view all your Cloudbrain tasks in",
    sameTaskTips5: "My Workspace",
    sameTaskTips6: "Computing Task",
    sameTaskTips7: "",
    cpuEnv: "CPU Environment",
    gpuEnv: "GPU Environment",
    npuEnv: "NPU environment",
    newTask: "Create new task",
    noQuene: "The current resources are sufficient and no tasks are queued",
    queneTips1: "Your current queue position is",
    queneTips2: "",
    watiResource: "Waiting for resources to be allocated, please be patient",
    debug: "Debug",
    stop: "Stop",
    stopping: "Stopping",
    notebookRunning: "The debugging environment has been started, and it will automatically close after 4 hours ",
    stopSuccess: "Stop task succeeded",
    specification: "Specification",
    graphicMemory: "Graphic Memory",
    memory: "Memory",
    sharedMemory: "Shared Memory",
    tips: 'The newly created debugging task will be placed in the project openi-notebook under your name. If there is no such project, the system will automatically create a new one.',
    balanceOfPoints: 'Balance of Points:',
    points: 'points',
    hours: 'Hours',
    expected_time: ', expected to be available for',
    limitTitle: 'The number of {type} has reached the limit.',
    limitReason: 'The current account can create up to {count} and simultaneously run {runCount} debugging tasks. Please stop running tasks and delete some historical tasks before submitting a new one.',
    stopReason: 'The current account already has {count} tasks waiting or running. Please stop them before submitting a new task.',
    deleteReason: 'The current account can create up to {count} {type}. Please delete historical tasks before submitting a new one.'
  }, _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_notebook, "stopSuccess", 'Stop task succeeded'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_notebook, "stopFailed", 'Stop task failed'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_notebook, "deleteSuccess", 'Delete task succeeded'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_notebook, "deleteFailed", 'Delete task failed'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_notebook, "handleSuccess", 'Operation succeeded. Please continue creating the task.'), _notebook),
  modelManage: {
    modelManage: 'Model management',
    modelName: 'Model name',
    useCluster: 'Available clusters',
    license: 'License',
    selectLicense: 'Select license',
    local: 'Local',
    online: 'Online',
    external: 'External',
    createModel: 'Create model',
    importLocalModel: 'Import Local Model',
    importOnlineModel: 'Import Online Model',
    importFrom: 'Where to import from',
    createNewModel: 'Create a new model',
    localModelTips: 'Import locally stored model files to the OpenI collaboration platform',
    onlineModelTips1: 'Import the model output from the training task into model management. To import an online model, you need to first',
    onlineModelTips2: 'create a training task',
    nextSteps: 'Next step',
    back: 'Back',
    step1: 'Select import type and project',
    step2: 'Set model information',
    step3: 'Upload model file',
    project: 'Belonging to repository',
    newRepoPath: 'New repository path',
    modelInfo: 'Model information',
    modelBriefIntro: 'Model brief introduction',
    trainingInfo: 'Training information',
    modifyModelInfo: 'Modify model information',
    modelList: 'Model list',
    modelFiles: 'Model files',
    addModelFiles: 'Add model files',
    uploadModelFiles: 'Upload model files',
    pleaseInputModelName: 'Please input model name',
    version: 'Version',
    modelVersion: 'Model version',
    modelEngine: 'Model engine',
    modelSource: 'Model source',
    modelLabel: 'Model label',
    externalToken: 'External Token',
    modelLabelInputTips: 'Input labels, multiple labels are separated by spaces',
    modelLabelGatedTips: "To migrate Gated's model files, it is necessary to input HuggingFace's Token",
    modelDescr: 'Model description',
    modelDescrInputTips: 'The description should not exceed 255 characters',
    confirm: 'Confirm',
    cancel: 'Cancel',
    modelCreateFailed: 'Model create failed',
    modelModifyFailed: 'Model modify failed',
    fileUpload: 'File upload',
    folderUpload: 'Folder upload',
    upload: 'Upload',
    uploadPath: 'Upload Path',
    uploadStatus: 'Upload status',
    modelFileUploadDefaultTips: 'Click to add files or drag files here directly',
    modelFileUploadErrTips: 'File name should not exceed 128 characters, Upload a maximum of {maxCount} files at a time, single file exceeding {size}, please use <a target="_blank" href="{url}" style="cursor: pointer;">SDK to upload</a>,  When there are a large number of files, it is recommended to compress them before uploading',
    modelFileNameTips: 'The file name should not exceed 128 characters',
    modelFileNamSspaceTips: 'The file name contains spaces',
    fileIstoBig: 'File size ({{filesize}} MiB) exceeds the maximum size of ({{maxFilesize}} MiB)',
    removeFile: 'Rmove file',
    uploadSuccess: 'upload success',
    uploadFailed: 'upload failed',
    calcFileMd5: 'Calculating file MD5...',
    uploading: 'Uploading...',
    fileHasAlreadyInTheModel: 'This file has already in the model: ',
    fileEqualInTheModel: 'This file content is the same as other file: ',
    fileExistInTheModel: 'This file name is the same as other file: ',
    uploadFile: 'Upload files',
    uploadFolder: 'Upload folders',
    addUploadFolder: 'Add upload folder',
    clearAll: 'Clear all',
    fileCountAndSize: 'Selected files: {count}, total size: {size}',
    folderUploadPlaceholder: 'Click to add a folder or drag folders/files here directly',
    folderUploadTips: 'The number of files uploaded this time does not exceed {maxCount}, and the total size does not exceed {maxSize}, please use <a target="blank" href="https://openi.pcl.ac.cn/docs/index.html#/api/list">SDK</a> to upload too many large files.',
    folderUploadErrorTips: 'The number of files uploaded this time does not exceed {maxCount}, and single file size does not exceed {maxSize}',
    basicInfo: 'Basic information',
    modelSize: 'Model size',
    descr: 'Description',
    createTime: 'Create Time',
    label: 'Label',
    trainTaskInfo: 'Train task information',
    trainTask: 'Train task',
    codeBranch: 'Code branch',
    bootFile: 'Boot file',
    viewSamples: 'View samples',
    trainDataset: 'Train dataset',
    datasetfile: 'Dataset files',
    fileShort: 'files',
    specInfo: 'Specifications',
    workServerNumber: 'Amount of compute node',
    runParameters: 'Run parameters',
    seeMore: 'View more',
    collapseDetails: 'Collapse details',
    modelFilesList: 'Mode files list',
    fileName: 'File name',
    fileSize: 'File size',
    updateTime: 'Update Time',
    operate: 'Operation',
    download: 'Download',
    "delete": 'Delete',
    infoModificationFailed: 'Information modify failed',
    deleteModelFileConfirmTips: 'Are you sure you want to delete the current model file?',
    modelFileDeleteFailed: 'Model file delete failed',
    modelAccess: 'Model Access',
    modelAccessPublic: 'Public',
    modelAccessPrivate: 'Private',
    modelAccessTips: 'Only public projects can be set as public models',
    modelSettings: 'Model Settings',
    edit: 'Edit',
    editFiles: 'Edit files',
    preview: 'Preview',
    modelIntroduction: 'Model Introduction',
    hasNoIntroForModel: 'No detailed model introduction yet',
    createModelIntro: 'Create model Introduction',
    briefIntroduction: 'Introduction',
    addLabels: 'Add labels',
    discardFileChanges: 'Discard current content modifications?',
    editFileContentFirst: 'Please edit the file content first!',
    ownerRepository: 'Repository',
    creator: 'Creator',
    migrator: 'Migrator',
    useModel: 'Use the model',
    modelEvolutionMap: 'Model Evolution Map',
    evolutionMap: 'Evolution Map',
    settings: 'Settings',
    parentModel: 'Parent Model',
    currentModel: 'Current Model',
    derivedModel: 'Derived Model',
    publicDerivedModel: 'Public Derived Model',
    privateDerivedModel: 'Private Derived Model',
    refRepository: 'Reference repository',
    publicRefRepository: 'Public Reference repository',
    privateRefRepository: 'Private Reference repository',
    modelDownloadAll: 'Download All',
    otherOnline: 'Other online url',
    trainUsedDataList: 'The dataset used with this model',
    trainUsedRepo: 'Code used with this modely',
    modelUseTaskList: 'The task used for the model',
    forkModelSuccess: 'The model content has been copied and repository forked successfully!',
    debugModel: 'Debug Model',
    onlineInference: 'Online Inference',
    onlineInferenceTask: 'Online Inference Task',
    onlineLoraTrain: 'Online Train',
    onlineWorkflow: 'online Workflow',
    deleted: 'Deleted',
    derivativeTimes: 'Derivative times',
    derivativeTimes1: 'Derivative',
    mostDerivative: 'Most Derivative',
    modelBaseInfo: 'Basic information of the model',
    modelCollaborator: 'Model collaborators',
    managementTeam: 'Model Management Teams',
    add_team_success: 'The team now have access to the model.',
    remove_team_success: 'The team access to the model has been removed.',
    deletCollaboratedTips: 'After deleting the collaborator, he will no longer be able to access this model. Do you want to continue with the operation?',
    total_size: 'Total size',
    total_size_asc: 'Total size in descending order'
  },
  repos: {
    activeOrganization: 'Active Organization',
    activeUsers: 'Active Users',
    follow: 'Follows',
    unFollow: 'Unfollow',
    selectedFields: 'Recommend Repositories',
    mostPopular: 'Most Popular',
    mostActive: 'Most Active',
    newest: 'Newest',
    recentlyUpdated: 'Recently Updated',
    mostStars: 'Most Stars',
    mostForks: 'Most Forks',
    mostDatasets: 'Most Datasets',
    mostAiTasks: 'Most AI Tasks',
    mostModels: 'Most Models',
    repos: 'Repositories',
    publicRepos: 'Public Repositories',
    repoTopics: 'Repositories Topics',
    dataset: 'Datasets',
    model: 'Models',
    aiTask: 'AI Tasks',
    updated: 'Updated',
    contributors: 'Contributors',
    searchRepositories: 'Search Repositories...',
    search: 'Search',
    allFields: 'All Fields',
    preferred: 'Preferred',
    openIIncubation: 'OpenI Incubation',
    hotPapers: 'Hot Papers',
    watch: 'Watch',
    star: 'Star',
    fork: 'Fork',
    noReposfound: 'No matching repositories found.',
    source: 'Sources',
    mirrors: 'Mirrors',
    collaborative: 'Collaborative',
    forks: 'Forks',
    selectRepo: 'Select repositories',
    selectRepoPlaceholder: 'Search repositories name...',
    cancelToping: 'Unpin',
    topping: 'Pin',
    toppingSuccess: 'Pinned successfully！',
    cancancelTopingSuccess: 'Unpinned successfully！',
    deleteThisRepos: 'Delete This Repository',
    delete_notices_2: '- This operation will permanently delete the <strong>{repos}</strong> repository including the code, issues, merge requests, and other contents',
    sureReposName: 'Enter the repository name as confirmation：',
    repos_name1: 'Repository Name',
    deleteReposNameError: 'Repository name input error!'
  },
  timeObj: {
    ago: '{msg} ago',
    from_now: '{msg} from now',
    now: 'now',
    future: 'future',
    '1s': '1 second',
    '1m': '1 minute',
    '1h': '1 hour',
    '1d': '1 day',
    '1w': '1 week',
    '1mon': '1 month',
    '1y': '1 year',
    seconds: '{msg} seconds',
    minutes: '{msg} minutes',
    hours: '{msg} hours',
    days: '{msg} days',
    weeks: '{msg} weeks',
    months: '{msg} months',
    years: '{msg} years',
    raw_seconds: 'seconds',
    raw_minutes: 'minutes'
  },
  modelObj: (_modelObj = {
    model_label: 'Model',
    model_select_placeholder: 'Select Models',
    model_export_placeholder: 'Please select model file',
    model_select: 'Select Models',
    model_current_repo: 'Current repository',
    model_my: 'My Models',
    model_collaborate: 'My Collaborated Models',
    model_collected: 'My Favorite Models',
    model_my_migrate: 'My Migrate Models',
    model_public: 'Public Models',
    model_recommend: 'Recommended Models',
    model_migrate: 'Migrate Models',
    model_search_placeholder: 'Search model name...',
    model_selected: 'Selected model',
    model_ok: 'OK',
    model_most: 'Up to {msg} models.',
    model_not_equal_file: 'Cannot select models with the same name',
    model_exceeds_failed: 'Model size exceeds',
    model_should_same_model: 'Select the files should in the same model.',
    model_search: 'Search model',
    model_square_empty: "It's empty, with nothing",
    model_suport_file_tips: 'The supported format of the model file is [ckpt, pb, h5, json, pkl, pth, t7, pdparams, onnx, pbtxt, keras, mlmodel, cfg, pt]',
    boot_file_helper: 'The startup file is the entry file that your program executes, and it must be a file ending in .py',
    can_online_infer: 'Online',
    can_fine_tune: 'Finetune'
  }, _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_exceeds_failed", "Model size exceeds "), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_sdk_use_way", 'Model SDK use way'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "codeUseDlgTriggerTxt", 'How to use models in OpenI'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "codeUseDlgTitle", 'How to use models on the OpenI'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "codeDownDlgTitle", 'How to download files on the OpenI'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_not_equal_file", "Cannot select models with the same name."), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_external_model", 'New Model Migration'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "external_model_name", 'External model name'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_external_model_name_tips", 'Currently only support migrating models from HuggingFace. Please enter the model name, such as THUDM/chatglm3-6b'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "openi_model_repo_url", 'OpenI model repository URL'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_model_should_be_public", 'Migrated model should be public'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_model", 'Migrate Model'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "please_enter_right_external_model_name", 'Please enter the correct external model name'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_external_model_failed", 'Migrate external model failed!'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "create_model_migrate_loading_content", 'Creating model migration in progress, please wait~'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "create_model_migrate_exists", 'The migrated model already exists, please check the <a href="{url}">details</a>.'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_from", 'Migrating from {url}'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_watting", 'Waiting'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrating", 'Migrating'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_success", 'Successful'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_failed", 'Failed'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_failed_tips", 'Model migration failed! Please '), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_retry", 'Re migrate'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "migrate_try_later", ', or try again later'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "re_migrate_failed", 'Re migrate failed'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "sync_now", 'Synchronize Now'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_sync", 'Model Synchronize Update'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "current_size", 'Current Size'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "new_size", 'Latest size'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "status_add", 'Added'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "status_del", 'Deleted'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "status_update", 'Updated'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "start_sync", 'Start Synchronize'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_no_update_tips", 'The model file remains unchanged and does not require updating'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_sdk_tips1", '## Download OpenI file to local folder:\n *File larger than {fileSize} When using SDK to download, please modify the parameters in the following code according to the specific **local save path*** \n ### Download using Python code'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_sdk_tips2", '#### Download using the command line'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_sdk_tips3", '\nFor installation and usage instructions of the openi library, please refer to the <a target="_blank" href="https://openi.pcl.ac.cn/docs/index.html#/model/sdk">User Manual</a>。'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_download", 'Download'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_source", 'Source model：'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "model_name1", 'Model name'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "deleteThisModel", 'Delete this model'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "deleteThisModelTips", 'Deleting a model is permanent and cannot be undone. Please proceed with caution.'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "delete_notices_2", '- This operation will permanently delete the model <strong>{model}</strong>.'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "sureModelName", 'Enter model name for confirmation：'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_modelObj, "deleteModelNameError", 'Model name input error!'), _modelObj),
  datasetObj: {
    dataset_label: "Datasets",
    dataset_select_placeholder: "Select Datasets",
    eval_select_placeholder: "Select up to 5 evaluation datasets",
    dataset_select: "Select Datasets",
    dataset_search_placeholder: "Search dataset name ...",
    dataset_unziping: "Decompressing",
    dataset_success: "Decompression successed",
    dataset_unzip_failed: "Decompression failed",
    dataset_status: 'Decompression state',
    dataset_exceeds_failed: "Dataset size exceeds ",
    dataset_my_upload: "Upload by me",
    dataset_current_repo: "Current Repository",
    dataset_public: "Public dataset",
    dataset_collected: "My collection",
    dataset_relate: "Related dataset",
    dataset_selected: "Selected datasets",
    dataset_ok: "OK",
    dataset_not_equal_file: "Cannot select a dataset with the same name.",
    dataset_most: "Up to {msg} datasets.",
    dataset_file_was_deleted: "The file has been deleted",
    dataset_sdk_use_way: 'Dataset SDK use way',
    codeUseDlgTriggerTxt: 'How to use datasets in OpenI',
    codeUseDlgTitle: 'How to use datasets on the OpenI',
    create_new_dataset: 'Create New Dataset',
    dataset_name1: 'Dataset name',
    dataset_name: 'English name of dataset',
    dataset_zh_name: 'Chinese name of dataset',
    dataset_description: 'Dataset Description',
    select_category: 'Select Category',
    task: 'Task',
    application: 'Application',
    select_task: 'Select research direction/application field',
    dataset_name_tooltips: 'Please enter letters, numbers, _, and -, which can only start and end with numbers or letters, with a maximum length of 100 characters.',
    dataset_name_cn_tooltips: 'Please enter Chinese characters, letters, numbers, _, and - up to 100 characters.',
    category: 'Category',
    license: 'License',
    license_helper: 'Select a license file',
    dataset_name_required1: 'Please enter the Chinese name of the dataset',
    dataset_name_required2: 'Input does not comply with dataset name rules',
    dataset_name_required3: 'Please enter the English name of the dataset',
    dataset_description_required: 'Please enter dataset description details',
    category_required: 'Please select a category',
    select_task_required: 'Please select research direction/application field',
    dataset_repo_tips: 'Since each repository can only create one dataset, only repository without a dataset can be specified here',
    exampleDataset: 'Example dataset',
    selectOpeniDataset: 'Select platform dataset',
    buildDatasetTips: 'How to build a dataset',
    create_public_tips: '*The default dataset created is a public dataset, which can be set as a private dataset on the dataset page',
    dataset_owner: 'Owner',
    name: 'Name',
    dataset_acess: 'Dataset access',
    dataset_size: 'Dataset size',
    hasNoIntroForModel: 'No detailed dataset introduction yet',
    createModelIntro: 'Create dataset Introduction',
    dataset_intro: 'Dataset Introduction',
    use_dataset: 'Use the dataset',
    deleteDataFileConfirmTips: 'Are you sure you want to delete the current file {name}?',
    dataFileDeleteFailed: 'Dataset file delete failed',
    dataDownloadAll: 'Download All',
    createDataset: 'Create dataset',
    deletCollaboratedTitle: 'Deleting the collaborator',
    deletCollaboratedTips: 'After deleting the collaborator, he will no longer be able to access this dataset. Do you want to continue with the operation?',
    dataBaseInfo: 'Basic information of the dataset',
    dataSearch: 'Search...',
    notFoundUser: 'No match found',
    administrators: 'Administrator',
    writePermission: 'Write',
    readPermission: 'Read',
    datasetCollaborator: 'Dataset collaborators',
    searchUsers: 'Search Users',
    addCollaborators: 'Add collaborator',
    managementTeam: 'Dataset Management Teams',
    searchTeams: 'search Teams',
    addTeam: 'Add Team',
    dangerousOperationZone: 'Dangerous operation zone',
    deleteThisDataset: 'Delete this dataset',
    deleteThisDatasetTips: 'Deleting a dataset is permanent and cannot be undone. Please proceed with caution.',
    sureDatasetName: 'Enter dataset name for confirmation：',
    deleteDatasetNameError: 'Dataset name input error!',
    add_collaborator_success: 'The collaborator has been added.',
    add_team_success: 'The team now have access to the dataset.',
    remove_collaborator_success: 'The collaborator has been removed.',
    remove_team_success: 'The team access to the dataset has been removed.',
    delete_notices_1: '- This operation <strong> cannot be </strong> rolled back.',
    delete_notices_2: '- This operation will permanently delete the dataset <strong>{dataset}</strong>.'
  },
  imagesObj: {
    cloudbrain_images: 'Image',
    openIGPU: 'OpenI GPU',
    c2netGPU: 'C2Net GPU',
    images_search: 'Search Image tag/description/operating system/python library/label...',
    image_public: 'Recommended Image',
    image_my: 'My Image',
    image_collected: 'My Favorite Image',
    framework: 'Framework',
    version: 'Version',
    frameworkName: 'Framework name',
    frameworkVersion: 'Framework version',
    pyVersion: 'Python version',
    cudaVersion: 'Cuda version',
    cannVersion: 'Cann version',
    dtkVersion: 'Dtk version',
    deleteTips: 'Are you sure you want to delete this image? Once this image is deleted, it cannot be recovered.',
    deleteSuccessTips: 'Successfully deleted',
    editImage: 'Modify image',
    submitImage: 'Submit image',
    appyImage: 'Apply to become a platform recommended image',
    imageTag: 'Image tag',
    imageDesc: 'Image description',
    imageTagPlaceholder: 'Please enter the image tag',
    imageTagInputTips: 'Please enter letters, numbers, _, . and -, with a maximum length of 50 characters and starting with a letter.',
    imageAdress: 'Image address',
    imageAdressPlaceholder: 'Please enter the image adress',
    imageAiCenterPlaceholder: 'Please enter the ai center. For example:\n[{"aiCenterId":"","trainJobUrl":"","notebookUrl":"","accDeviceModel":"",poolIds:""}]',
    operationSystem: 'Operating system',
    operationSystemName: 'Operating system name',
    operationSystemNamePlaceholder: 'Please enter the operating system name',
    operationSystemVersionPlaceholder: 'Please enter the operating system version',
    pyPackge: 'Python libraries',
    thirdPackages: 'Python libraries and version',
    thirdPackagesPlaceholder: "Please enter the installed Python library and version, one per line, up to 10 lines. For example:\ntorch==1.13.1\ntransformers==4.25.1",
    topic: 'Label',
    topicPlaceholder: 'After entering, press enter to confirm the label',
    topicTips: "Please provide additional information in the label field for the image, such as the compatible card type: <span class=\"light\">T4</span>, <span class=\"light\">Ascend 910</span>",
    descr: 'Image description',
    descrPlaceholder: 'Please enter the image description, which should not exceed 1000 characters',
    recommend: 'Recommend',
    notRecommend: 'Unrecommend',
    submitTips: 'The code directory /tmp/code will not be submitted with the image, and new files (folders) in other directories will be packaged into the image.',
    submitNpuTip1: 'The directory /home/ma-user/work will not be submitted with the image, and other directories will be packaged into the image.',
    submitNpuTip2: 'During the image saving process, the Notebook will be unavailable and the status of this task will change from Running to Waiting.',
    submitDcuTip1: '目录/home/ma-user/work不会随镜像提交，其他目录都会打包到镜像中。',
    submitDcuTip2: 'The image must be smaller than 30G to be saved successfully. If the image is not used for 30 days, the system will be cleaned up and cannot be used',
    submitApply: 'Submit Apply',
    imageCommitting: 'Image committing...',
    imageCommitSuccess: 'Image committed successfully',
    imageCommitErrorTips1: 'Check whether the size of the submitted image exceeds 20G!',
    imageCommitErrorTips2: 'Check whether the size of the submitted image exceeds 40G!',
    committing: 'Commiting',
    commitSuccess: 'Committed successfully',
    commitFailed: 'Committed failed',
    recommendNeedReview: 'Pending approval',
    recommendReviewApproved: 'Approved',
    recommendReviewFailed: 'Not approved',
    applyForRecommend: 'Apply recommendation',
    copyAdress: 'Copy adress',
    filterImages: 'Filter images',
    filterImagesPlaceholderCuda: 'Framework name/Framework version/Python version/Cuda version',
    filterImagesPlaceholderCann: 'Framework name/Framework version/Python version/Cann version',
    filterImagesPlaceholderDtk: 'Framework name/Framework version/Python version/Dtk version',
    imageTaskType: 'Task type',
    imageTaskTips: 'Apply to',
    approval_status: 'Approval status',
    all_approval_status: 'All Approval Status',
    pending_approval: 'Pending approval',
    approved: 'Approved',
    not_approved: 'Not approved',
    none: 'None',
    create_cloud_brain_mirror: 'Create cloud brain image',
    not_recommend: 'Not recommend'
  },
  specObj: {
    resSelectTips: 'The "resource specification" is the hardware you use to run the task. In order for more people to use the resources of this platform, please select according to your actual needs',
    no_use_resource: 'No resources available'
  },
  datasets: {
    computer_vision: "computer vision",
    natural_language_processing: "natural language processing",
    speech_processing: "speech processing",
    computer_vision_natural_language_processing: "computer vision and natural language processing",
    machine_translation: "machine translation",
    medical_imaging: "medical imaging",
    question_answering_system: "question answering system",
    information_retrieval: "information retrieval",
    knowledge_graph: "knowledge graph",
    text_annotation: "text annotation",
    text_categorization: "text categorization",
    emotion_analysis: "emotion analysis",
    language_modeling: "language modeling",
    speech_recognition: "speech recognition",
    automatic_digest: "automatic digest",
    information_extraction: "information extraction",
    description_generation: "description generation",
    image_classification: "image classification",
    face_recognition: "face recognition",
    image_search: "image search",
    target_detection: "target detection",
    image_description_generation: "image description generation",
    vehicle_license_plate_recognition: "vehicle license plate recognition",
    medical_image_analysis: "medical image analysis",
    unmanned: "unmanned",
    unmanned_security: "unmanned security",
    drone: "drone",
    vr_ar: "VR/AR",
    "2_d_vision": "2.D vision",
    "2.5_d_vision": "2.5D vision",
    "3_d_reconstruction": "3Dreconstruction",
    image_processing: "image processing",
    video_processing: "video processing",
    visual_input_system: "visual input system",
    speech_coding: "speech coding",
    speech_enhancement: "speech enhancement",
    speech_synthesis: "speech synthesis",
    ros_hmci_datasets: "ROS-hmci datasets",
    downloadtimes: 'Download times',
    downloadtimes1: 'Download',
    citations: 'Citations times',
    citations1: 'Citations',
    "default": "Default",
    newest: "Newest",
    oldest: "Oldest",
    recentupdate: "Recently updated",
    leastupdate: "Least recently updated",
    moststars: "Collection Nums",
    mostCollections: "Most Stars",
    mostusecount: "Most Quote",
    alphabetasc: "Alphabetical order ascending",
    alphabetdesc: "Alphabetical order descending",
    chinesename: "Chinese name",
    chinesenameasc: "Reverse Order of Chinese Names",
    size: "Size",
    sort: "Sort",
    unstarSuccess: "Cancel favorite successfully!",
    starSuccess: "favorite successfully！",
    platform_recommendations: 'Show platform recommendations only',
    category: 'Category',
    task: 'Task',
    license: 'License',
    publick_dataset: 'Public Dataset',
    my_dataset: 'My Datasets',
    favorite_dataset: 'My Favorite Datasets',
    recommend_dataset: 'Recommended Datasets',
    collaborated_dataset: 'My Collaborated Datasets',
    views: 'Views',
    loadMore: 'Load more'
  },
  cloudbrainObj: {
    cloudbrain: 'Cloudbrain',
    openi: 'OpenI Resource Cluster',
    c2net: 'China Computing NET(Beta)',
    create: 'New ',
    createTask: 'Create Task',
    cluster: 'Resource cluster',
    computeResource: 'Computing resource',
    sameTaskTips1: 'You have already {count} running or waiting task(s), which has reached the maximum limit. Please wait until the tasks are finished before creating new ones.',
    sameTaskTips2: 'You can view all your Computing tasks in <a href="/cloudbrains" target="_blank"> My Workspace &gt; Computing Task </a>.',
    pathTips1: 'The code is storaged in <strong style="color:#010101">{code}</strong>, the dataset is storaged in <strong style="color:#010101">{dataset}</strong>, the pre-trained model is storaged in the run parameter <strong style="color:#010101">{model}</strong>, and please put your model into <strong style="color:#010101">{output}</strong> then you can download it online.',
    pathTips11: 'The code is storaged in <strong style="color:#010101">{code}</strong>, the dataset is storaged in <strong style="color:#010101">{dataset}</strong>, the pre-trained model is storaged in the <strong style="color:#010101">{model}</strong>, and please put your model into <strong style="color:#010101">{output}</strong> then you can download it online.',
    pathTips2: 'The code is storaged in <strong style="color:#010101">{code}</strong>, the dataset is storaged in <strong style="color:#010101">{dataset}</strong>, the pre-trained model is storaged in the <strong style="color:#010101">{model}</strong>, and please put your model into <strong style="color:#010101">{output}</strong> then you can download it online.',
    pathTips3: 'The code is storaged in <strong style="color:#010101">{code}</strong>, the dataset is storaged in <strong style="color:#010101">{dataset}</strong>, the pre-trained model is storaged in the <strong style="color:#010101">{model}</strong>.',
    pathTips4: 'The code is storaged in <strong style="color:#010101">{code}</strong>, the pre-trained model is storaged in the run parameter <strong style="color:#010101">{model}</strong>, and please put your model into <strong style="color:#010101">{output}</strong> then you can download it online.',
    pathTips5: 'The dataset is stored in <strong style="color:#010101">{dataset}</strong>, the model file is stored in <strong style="color:#010101">{model}</strong>, please store the inference output in <strong style="color:#010101">{output}</strong> for subsequent downloads.',
    pathTips6: 'The dataset location is stored in the run parameter <strong style="color:#010101">{dataset}</strong>, the pre-trained model is storaged in the run parameter <strong style="color:#010101">{model}</strong>, and the output path is stored in the run parameter <strong style="color:#010101">{output}</strong>.',
    pathTips66: 'The dataset location is stored in the run parameter <strong style="color:#010101">{dataset}</strong>, the pre-trained model is storaged in the run parameter <strong style="color:#010101">{model}</strong>, and please put your model into <strong style="color:#010101">{output}</strong> then you can download it online.',
    pathTips7: 'The dataset location is stored in the run parameter <strong style="color:#010101">{dataset}</strong>, the model file is stored in the run parameter <strong style="color:#010101">{model}</strong>, and the inference output path is stored in the run parameter <strong style="color:#010101">{output}</strong>.',
    pathTips71: 'The code is storaged in <strong style="color:#010101">{code}</strong>, the dataset is storaged in <strong style="color:#010101">{dataset}</strong>, the model file is storaged in the <strong style="color:#010101">{model}</strong>, and please put your inference output into <strong style="color:#010101">{output}</strong> then you can download it online.',
    basicInfo: 'Basic Info',
    paramsSetting: 'Parameter setting',
    resourceSetting: 'Resource setting',
    runningStatus: 'Running status',
    basicParameters: 'Basic parameters',
    otherParameters: 'Other parameters',
    parameters: 'Parameters',
    valueAndContent: 'Value&Content',
    taskPrepareTips: 'The task is currently being prepared. Drink a glass of water and come back to take a look~',
    waitCountStart: 'Your current queue position is ',
    waitCountEnd: '',
    mindTorchHelper: 'Don\'t want to queue up? View how to migrate PyTorch to MindSpore with just one step',
    task: ' Task',
    taskList: 'Task List',
    taskName: 'Task name',
    appName: 'App name',
    appList: 'App List',
    taskNameTips: 'Name must start with a lowercase letter or number,can include lowercase letter,number,_ and -,can not end with _, and can be up to 36 characters long.',
    taskNameTips1: 'Name must start with a lowercase letter,can include lowercase letter,number,and -, and between 5 and 26 characters.',
    imageInnerUrlErrTips: 'The image url you specified is not applicable to this task. Please specify a different image.',
    imageUrlErrTips: 'Please enter a valid image url.',
    taskDescr: 'Description',
    taskDescrPlaceholder: 'The description should not exceed 255 characters',
    codeBranch: 'Code branch',
    image: 'Image',
    selectImage: 'Select Image',
    selectImagePlaceholder: 'Select image or input image url',
    dataset: 'Datasets',
    networkType: 'Access Internet',
    allNetworkType: 'All network',
    codeGenerate: 'Code Generate',
    noInternet: 'No',
    hasInternet: 'Yes',
    networkTypeDesc: '「Access Internet」Describe whether the computing resource center supports internet access. When \'No\' is selected, more computing resource centers can be allocated, and the dataset and models on the OpenI platform can still be used.',
    resourceSpec: 'Specification',
    specPlaceholder: 'Select resource specification',
    specDescr: 'Resource Note',
    PointGainDescr: 'Points acquisition instructions',
    balanceOfPoints: 'Balance of Points',
    points: 'Points',
    canUseTime: ', expected to be available for ',
    hours: 'Hours',
    versionCount: 'Version Nums',
    runVersion: 'Run Version',
    createTime: 'Create time',
    startRunTime: 'Start run time',
    endRunTime: 'End run time',
    runDuration: 'Running Time',
    refresh: 'Refresh',
    codePath: 'Code storage path',
    datasetPath: 'Dataset storage path',
    modelPath: 'Model storage path',
    outputPath: 'Output storage path',
    codeObsPath: 'Code OBS address',
    clusterAndComputeResource: 'Cluster/Compute Resources',
    runParameter: 'Run Parameters',
    addRunParameter: 'Add Run Parameters',
    parameterName: 'Parameter Name',
    parameterValue: 'Parameter Value',
    customPath: 'Custom Path',
    port: 'Port',
    customPathTips: 'Custom paths are only allowed to contain lowercase letters and numbers, with the first character being a letter, and the length must not exceed 32 characters; The port number can be selected from the range of 8000-8800',
    debug: 'Debug',
    reDebug: 'Restart',
    reInfer: 'ReInference',
    stop: 'Stop',
    modify: 'Modify',
    "delete": 'Delete',
    startUse: 'Start use',
    more: 'More',
    commitImage: 'Submit Image',
    downloadModel: 'Download',
    configurationInfo: 'Configuration Information',
    taskRuntimeInfo: 'Task Runtime Information',
    log: 'Logs',
    logFile: 'Log file',
    downloadLog: 'Download log file',
    viewFullScreen: 'Fullscreen',
    exitFullScreen: 'Exit fullscreen',
    scrolledToTopTip: 'You have scrolled to the top of the log',
    scrolledToBottomTip: 'You have scrolled to the bottom of the log',
    resourceOccupancy: 'Resources Occupancy',
    modelDownload: 'File Download',
    lossPlot: 'loss',
    evalOverview: 'Overview of Evaluation Results',
    evalDetail: 'Evaluation Results Details',
    failedReason: 'Failed reason',
    publicImage: 'Public Image',
    recommendImage: 'Recommend Images',
    myImage: 'My Images',
    myFavImage: 'My collected images',
    searchImagePlaceholder: 'Search image tag/description/operating system/python library/label...',
    useImage: 'Use',
    submitting: 'Commiting',
    submitFailed: 'Commit failed',
    checkImageSizeTips: 'Check whether the size of the submitted image exceeds 20G.',
    maxTaskTips: '<p><span>*</span>The platform only retains the results of debug, train, inference and evaluation tasks for nearly<span> 30 </span> days. <span>Tasks over 30 days will not be able to download results and view logs, and cannot be debugged、evaluation and trained again</span></p>',
    datasetFiles: 'Datasets',
    fileWasDeleted: 'The file has been deleted',
    debugTaskEmptyTitle: 'Debug task has not been created',
    debugTaskEmptyTip0: 'Code version: You have not initialized the code repository, please <a href="{url}">initialized</a> first;',
    debugTaskEmptyTip1: 'Running time: no more than 4 hours, it will automatically stop if it exceeds 4 hours;',
    debugTaskEmptyTip2: 'Dataset: Cloud Brain 1 provides CPU/GPU,Cloud Brain 2 provides Ascend NPU. And dataset also needs to be uploaded to the corresponding environment;',
    debugTaskEmptyTip3: 'Instructions for use: You can refer to the OpenI AI collaboration platform<a href="{url}"> Help Center </a>.',
    onlineInferTaskEmptyTitle: 'Online Inference task has not been created',
    onlineInferEmptyTip2: 'Dataset: Cloud Brain 1 provides CPU/GPU,Cloud Brain 2 provides Ascend NPU. And dataset also needs to be uploaded to the corresponding environment;',
    superTaskEmptyTitle: 'HPC task has not been created',
    generalTaskEmptyTitle: 'General task has not been created',
    generalTaskEmptyTip1: 'This task type has the following characteristics and can only be used after submitting a usage application and approval on the <a href="/computingpower/demand">Computing resources</a> page:<br/> (1) Provide a Jupyter debugging environment.<br/> (2) The running time is unlimited, but points are required.<br/> (3) Supports open ports.<br/> (4) Saving images is not supported, and debugging again is not supported.',
    deleteConfirmTips: 'Are you sure you want to delete this task? Once this task is deleted, it cannot be recovered.',
    deleteBatchConfirmTips: 'Are you sure to batch delete the selected tasks? Once a task is deleted, it cannot be restored.',
    deletingTips: 'Task deletion in progress, please wait',
    tabTitDebug: 'Debug Task',
    tabTitTrain: 'Train Task',
    tabTitInference: 'Inference Task',
    tabTitBenchmark: 'Model Evaluation',
    trainTaskEmptyTitle: 'Train task has not been created',
    trainTaskEmptyTip2: 'Dataset: Cloud Brain 1 provides CPU/GPU,Cloud Brain 2 provides Ascend NPU. And dataset also needs to be uploaded to the corresponding environment;',
    inferenceTaskEmptyTitle: 'Inference task has not been created',
    inferenceTaskEmptyTip2: 'Dataset: Cloud Brain 1 provides CPU/GPU,Cloud Brain 2 provides Ascend NPU. And dataset also needs to be uploaded to the corresponding environment;',
    bootFileTips: 'The startup file is the entry file that your program executes, and it must be a file ending in .py',
    viewSample: 'View sample',
    tabTitOnlineInference: 'Online Inference',
    tabTitGeneral: 'General Task',
    allResultDownload: 'All result download',
    downloadDisplayMaxCountTips: 'A maximum of {count} files or folders are displayed in a single directory. The storage quota for compute task results is {size} GiB. Exceeding the quota will result in data loss.',
    file_sync_ing: "File synchronization in progress, please wait",
    file_sync_wait: "File synchronization in waiting, please wait",
    file_sync_fail: "File synchronization failed",
    eval_task_ing: "The evaluation task is in progress, come back later to take a look",
    no_file_to_download: "No files can be downloaded",
    task_not_finished: "Task not finished yet, please wait",
    retrieve_results: "Retrieve results",
    reuseLastResult: 'Reuse last result',
    continue_helper: 'Check Reuse to copy the output result file of the last training task',
    computeNode: 'Compute Node',
    computeNodeCount: 'Compute node count',
    exportDataset: {
      exportDatasetTitle: 'Export the results to a dataset',
      exportDatasetTips: '<span style="color:red">*</span> The exported file can ultimately be viewed under the dataset of your choice.',
      export_failed: 'Export failed',
      export_has_same_file: 'The same file already exists in the current dataset',
      export_has_same_file1: 'The same file already exists in the current model',
      export_exceed_storage: 'The current file exceeds the size limit',
      export_success: 'Export success',
      exporting: ' Exporting',
      please_select_file: 'Please select a file first',
      please_select_output_file: 'Select result file',
      select_file: 'Select File',
      file_descr: 'File description',
      no_dataset: 'No dataset has been created yet, ',
      create_dataset: 'to create the dataset.',
      please_select_dataset: 'Please select the dataset to export to first'
    },
    chartResourceUsage: 'Usage(%)',
    chartTime: 'Time(min)',
    scrollToTop: 'Scroll to top',
    scrollToBottom: 'Scroll to bottom',
    migratingData: 'Data migration in progress',
    centerPending: 'Queuing in sub centers',
    imagePulling: 'Pulling image in progress',
    sdkUseWay: 'c2net library usage',
    dialogTips: {
      title1: 'When using online inference to provide webui or API services, it is necessary to understand the following points',
      tips1: "The platform does not directly expose custom ports to the external network for service provision. Instead, FastAPI can be used to forward requests to a designated externally exposed URL for service delivery.",
      tips1_1: 'Webui service reference:',
      tips1_2: 'API service reference:',
      tips2: 'Service startup instruction reference',
      tips3: 'Online reasoning tasks only have logs returned when the task is terminated',
      tips61: 'Detailed usage methods can be found in ',
      tips62: 'the code repository',
      tips7: "Don't prompt again",
      tips8: 'Close',
      tips9: 'OpenI Online Inference Deployment Requirements'
    },
    codeUseDlgTriggerTxt: 'How to access data resources in code',
    codeUseDlgTitle: 'How to obtain models, datasets, and output paths in code via the c2net library',
    sdkCodeTip1: 'Please use c2net library to access relevant resources in the container, which can be referred to as <a target="_blank" href="https://openi.pcl.ac.cn/docs/index.html#/cloudbrain/codepath">Help</a>.',
    sdkCodeTip2: 'Based on the type of computing task you choose, as well as the specified datasets, models, etc. The example code for accessing datasets and models and uploading training output are as follows:',
    generalTaskSdkCodeTip0: 'How to open ports for general task',
    generalTaskSdkCodeTip1: 'Run the following command to configure HTTP proxy in the environment:',
    generalTaskSdkCodeTip2: '/root/bin/scc tunnel http http://127.0.0.1:7860 (Port number 7860 can be customized)',
    generalTaskSdkCodeTip3: 'After successful configuration, output the following information:',
    generalTaskSdkCodeTip4: 'You can use the address <code>https://be22fe9f-d140-474c-81f4-30fb068157b5.tunnel.paracloud.com</code> to access the corresponding service.',
    searchTaskName: 'Search Task Name...',
    searchTaskNameOrCreator: 'Search Task Name/Creator...',
    downloadReport: 'Download Report',
    cloudbrainTaskType: 'Task Type',
    ComputingResourceInfo: 'Computing resource information',
    repo: 'Repository',
    cloudbrainTaskName: 'Cloudbrain Name',
    taskIsAutomaticStop: 'Task is automatic stop',
    automaticStop: 'Automatic stop',
    manualStop: 'Manual stop',
    automaticStopTips: 'This task will automatically stop when the running time exceeds the time you have selected, or when the point balance is insufficient.',
    manualStopTips: 'This task will be manually stopped by you or automatically stopped when your point balance is insufficient.',
    customize: 'Customize',
    numOfHours: '{num} hours',
    customizeTimeLimitPlaceholder: 'Please enter an integer between 1 and 24',
    debugTaskTimeLimitTips: 'The debug type task has a usage time limit, and will automatically end if it expires. <span>Please save the debugging results before the deadline<span>.',
    autoStopTimeTips: 'The task will automatically stop after {min} minutes. Please save the results',
    forkRepo: 'Fork existing repository',
    newRepo: 'New repository',
    migrateRepo: 'New migrate',
    pullRequests: 'Pull requests',
    issues: 'Issues',
    selectRepo: 'Select repository',
    nextStep: 'Next Step',
    previousStep: 'Previous Step',
    instructionsForUse: 'Instructions for use',
    setUp: 'Set ',
    aiTaskType: 'AI task type',
    newCloudbrainAiTask: 'New Cloudbrain(AI) Task',
    selectAiTaskTypeAndRepo: 'Select AI Task Type and Repository',
    specifyComputingResourcesAndRarameters: 'Specify Computing Resources and Parameters',
    dataPreparing: 'Data preparation in progress, please wait a moment~',
    selectTheTaskType: 'Select the type of task you need to create.',
    debugTaskDesc: 'Interactive editing environment',
    debugTaskDescLong: 'Provide a Jupyter or Modelarts interactive debugging environment, suitable for debugging and building early models.',
    debugTaskDescLong1: 'Debug task will automatically stop after running for more than 4 hours.',
    trainTaskDesc: 'Submit to task queue',
    trainTaskDescLong: 'Suitable for scenarios where long-term execution does not require frequent modification of algorithm code.',
    inferenceTaskDesc: 'Supercomputing power computing task',
    inferenceTaskDescLong: 'Support running multiple applications on supercomputing power.',
    onlineinferTaskDesc: 'Build a real-time inference interface',
    onlineinferTaskDescLong: 'Provide online deployment capabilities for the model, using the Gradio framework to generate real-time inference interfaces.',
    generalTaskDesc: 'Provide Jupyter debugging environment',
    generalTaskDescLong: 'Provide a Jupyter debugging environment, which requires submitting a usage application on the computing resources page and approval before it can be used.',
    noAiTasksInThisRepo: 'There are no Computing tasks under this repository yet',
    createNewAitask: 'Create computing task',
    visualization: 'Visualization',
    sourceFtName: 'Source fine-tuning task name',
    tensorBoardVisualization: 'TensorBoard Visualization',
    viewVisualization: 'View TensorBoard visualization results',
    startDebug: 'Start Debug',
    stopTask: 'Stop Task',
    taskTmpl: 'Computing task Template',
    saveTaskTmpl: 'Save As Computing Task Template',
    saveNewModel: 'Export To New Model',
    exportData: 'Export Data',
    exportToDataset: 'Export To Dataset',
    deploymentExperience: 'Deployment Experience',
    aimTrainCompare: 'Training Comparison',
    aimTrainCompareDlgTitle: 'Select Training Tasks for Comparison',
    aimTrainCompareSelectTips: 'Please select {minCount} to {maxCount} tasks for training comparison!',
    aimVisualization: 'AIM Visualization',
    batchDelete: 'Batch delete',
    deleteFailed: 'Delete Failed',
    all: 'All',
    ihave: 'My Datasets',
    iCollaborate: 'My Collaborated Datasets',
    iCollect: 'My Favorite Datasets',
    selectApp: 'Select App',
    cbOffline: 'The Cloudbrain tab under the project is about to go offline',
    cbOffflineTips: 'Please go to the <a href="/cloudbrains/create"> "My Workspace>Computing Tasks" </a> page to view the existing Computing Tasks on this page。',
    exportOutputTis: 'Export the results to the dataset or model, please move to the PC web page',
    repoStorageTips: 'Control project size within the storage quota of {size} GiB for compute tasks.',
    runningLimit: 'This account can run {count} {taskType} simultaneously.',
    forbidPenetrationStatement: 'Declaration of Prohibition of Any Form of Penetration Behavior',
    penaltyTipContent: 'This platform strictly prohibits any form of <span class="highlight">penetration behavior</span>. Violators will have their computing power points deducted, login restricted, and serious cases will bear <span class="highlight">legal responsibility</span>!'
  },
  taskTmplObj: {
    tmplName: 'Template Name',
    tmplNamePlacehoulder: 'Please enter a template name of up to 20 characters',
    tmplDescr: 'Template Description',
    quickRunTmplTips: 'One click run computing task template',
    editTmplFile: 'Edit template file',
    addTmpl: 'New template',
    dragSortTips: 'Drag tab to adjust order',
    deleteTmplTips: 'Do you want to delete the current Computing Task template?',
    completeTmplTips: 'Please complete the required content in the position {n} template!',
    more: 'More',
    collapsed: 'Collapsed',
    edit: 'Edit',
    run: 'Run',
    runTmplForkTips: 'You do not have permission for this repository to create a new computing task. We suggest that you fork this repository and then run the computing task with just one click.',
    forkRepo: 'Fork this repository',
    runTmplFailedTips: 'Run template failed!',
    runTmplWithConfigsErrorTips: 'Template configuration exception, unable to run properly!',
    useTaskTmpl: 'Using task template',
    saveTaskTmpl: 'Save computing task template',
    unAvailableTmplSpec: 'Template Specification "{msg}" due to reasons such as non-existent permissions or resources, <span style="color:red">currently unavailable.</span>',
    unAvailableTmplImage: 'Template Image "{msg}" due to reasons such as non-existent permissions or resources, <span style="color:red">currently unavailable.</span>',
    unAvailableTmplBranch: 'Template Code Branch "{msg}" due to reasons such as non-existent permissions or resources, <span style="color:red">currently unavailable.</span>',
    taskTmpl: 'Computing Task Template',
    publicTmpl: 'Public Templates',
    recommendTmpl: 'Recommended Templates',
    runTimes: 'Number of runs',
    searchTaskTmpl: 'Search for template name',
    createTaskTmpl: 'Create computing task template',
    myCreated: 'My Templates',
    myCollected: 'My Favorite Templates',
    deleteTaskTmplConfirmTips: 'Confirm delete the current computing task template?',
    tmplAccessRight: 'Template Access',
    accessRight: 'Access',
    collectedNum: 'Collection Nums',
    tagsAndDescr: 'Tags and Description',
    modelDatasetAndRepo: 'Model, Dataset and Repository',
    cancelRecommend: 'Cancel Recommendation',
    setRecommend: 'Set Recommended',
    cancelRecommendSuccess: 'Cancel recommendation success',
    setRecommendSuccess: 'Set recommended success',
    tmplCoreElements: 'Core Elements',
    editTmpl: 'Edit Template',
    tmplInfo: 'Template Information',
    tmplTags: 'Template tags',
    tagsPlaceholder: 'After entering, press enter to confirm the label, enter up to 10',
    editTaskTmpl: 'Edit Computing Task Template',
    createNewTaskTmpl: 'New Computing Task Template',
    createTaskTmplErrTips: 'Please complete the computing task template information',
    taskTmplSaveTips: 'The computing task template is being saved, please wait a moment',
    taskTmplSaveFailedTips: 'Failed to save computing task template',
    taskTmplReferencedRepo: 'Computing task templates that have referenced this repository'
  },
  superComputeObj: {
    mmlSparkDescr: "The full name of MMLSpark is Microsoft Machine Learning for Apache Spark, which enables users to run customized container images and grants them root accesses within the container. Users can directly use Microsoft's MMLSpark provided by the platform.\nNote: MMLSpark is a Spark version provided by Microsoft for machine learning environments\uFF08 <a target=\"_blank\" href=\"https://github.com/Azure/mmlspark\">https://github.com/Azure/mmlspark</a> \uFF09Regarding mmlspark, please refer to the following paper: <a target=\"_blank\" href=\"https://arxiv.org/pdf/1810.08744.pdf\">https://arxiv.org/pdf/1810.08744.pdf</a>"
  },
  modelSquare: {
    llmHeader: 'Model experience',
    chatGlm_intro: 'is an open source conversational language model that supports Chinese and English bilingualism, provided by Zhipu AI.',
    llama2: 'is a collection of pretrained and fine-tuned generative text models ranging in scale from 7 billion to 70 billion parameters. This is the repository for the 7B fine-tuned model, optimized for dialogue use cases and converted for the Hugging Face Transformers format.',
    dialogtips1: 'Hello 👋! Welcome to experience the large model knowledge base Q&A',
    dialogtips21: 'This experience is based on',
    dialogtips22: 'language model and m3-base vector model',
    dialogtips3: 'Please choose to talk to the model directly or ask questions based on the local knowledge base on the right.',
    dialogtips4: 'In the knowledge base Q&A mode, after selecting the name of the knowledge base, you can start Q&A.',
    dialogtips5: 'If necessary, you can upload files/folders to the knowledge base or delete files from the knowledge base after selecting the knowledge base name.',
    dialogtips6: 'The model has been successfully loaded, you can start the conversation, or select the mode from the right to start the conversation',
    promptPlaceholder: 'Please input the question content (Ctrl+Enter=line feed, press Enter to submit)',
    dialogModeSelect: 'Please select the dialogue mode',
    dialogLLM: 'LLM dialogue',
    dialogKb: 'Knowledge Base Q&A',
    configKb: 'Configure Knowledge Base',
    updatekb: 'Update existing knowledge base options',
    recreateKb: 'Vector library reconstruction in progress, please be patient and do not refresh or close the webpage',
    selectKb: 'Please select the knowledge base to load:',
    createKb: 'Create knowledge base',
    deleteKb: 'Delete knowledge base',
    deleteKbTips: 'Are you sure to delete the {knowledgeValue} knowledge base?',
    deleteVbTips: 'Are you sure to delete the knowledge base files？',
    uploadFile: 'Upload files',
    uploadFileTips1: 'Drag the file here, or click <em>to upload</em>',
    uploadFileTips2: 'Single file upload size limit is 1MB • HTML, MD, JSON, CSV, TXT, XML, DOCX',
    addFileToKb: 'Add files to the knowledge base',
    manageFile: 'Manage files',
    deleteKbFileSelect: 'Please select the file to delete from the existing files in the knowledge base',
    deleteKbFile: 'Delete files from the knowledge base',
    recreateKbSuccess: '{knowledgeValue} vector library reconstruction successful',
    noPermission: 'You do not have permission to operate',
    fileExit: 'File already exists',
    fileExceed: 'Files exceeding 1MB',
    fileError: 'File error, please re-upload!',
    fileTypeError: 'File type error, please re-upload!',
    fileUploadSuccess: '{fileName} File uploaded successfully!',
    kbName: 'Knowledge Base Name',
    createKbPlaceholder: 'The new knowledge base name can only be numbers and letters',
    vectorType: 'Vector Library Type',
    embedModel: 'Embedding Model',
    cancel: 'Cancel',
    create: 'Create',
    ok: 'Ok',
    kbNameDetect1: 'Name cannot be empty',
    kbNameDetect2: 'The name can only be numbers and letters',
    chatExceedCount: 'If the usage limit is exceeded, you will not be able to experience the model service!',
    useNotice: 'User experience instructions',
    agreeNotice: 'Agree to <a href="/home/model_privacy" target="_blank"> 《OpenI Qizhi Community AI Collaboration Platform Disclaimer and Service Usage Specification》</a> <p style="text-align: center;margin-top: 1rem;color: red;">Kind reminder: Unreasonable use may result in account closure</p>',
    modelProvide: '《Disclaimer and Service Usage Specifications》',
    modelNotExist: 'Model not exist',
    maxTries: 'Limited experience {maxTries} times',
    modelChatTask: 'Create a model online experience task',
    createChatTips1: 'Click the button below to create an online experience task. After successful creation, you can experience {expireMinutes} minutes online',
    createChatTips2: 'The online experience tasks created are exclusive to your account only.',
    createChatTips3: 'You have already created an online experience task for the current model. Click the button below to directly enter the experience interface',
    createChatBtns1: 'Create Online Experience task',
    createEvalBtns1: 'Create Evaluation task',
    createChatBtns2: 'online experience',
    experienceTime: 'Experience Countdown：',
    uploadFIleLimit: 'Upload up to 10 files',
    inputNotEmpty: 'The input content cannot be empty!',
    sessionChating: 'Session loading, don not be impatient!',
    sessionSding: 'Image generation in progress, don not be impatient！',
    chatBanned: 'Your account has been banned, please contact the website administrator',
    chatIllegal: "I'm very sorry, as an artificial intelligence assistant I can only provide objective information. Do you have anything else to ask?",
    chatExpireMins: 'If the experience time exceeds {expireMinutes} minutes, the [online experience task]({locaRefresh}) needs to be re-created.',
    chatExpired: 'Chat session expired, please create a new chat.',
    stopExperience: 'Stop the experience',
    experienceDuration: 'Experience duration',
    requireTimeout: 'Request timed out, please try again later',
    sdPlaceholder: 'Please describe the desired image content in Chinese or English, and using English prompts will result in better results.',
    submitTips: 'Ctrl + Enter line break',
    negativePromptPlaceholder: 'Please describe the screen content that you do not want to generate.',
    width: 'Width',
    height: 'Height',
    num_images_per_prompt: 'Number images',
    negative_prompt: 'Negative prompt',
    steps: 'Steps',
    seed: 'Seed',
    scheduler_name: 'Scheduler',
    guidance_scale: 'Guidance scale',
    newChat: 'New Chat',
    stopChat: 'Stop generating',
    stopedChat: 'Stoped generating',
    refreshChat: 'Regenerate',
    like: 'Like',
    unlike: 'Dislike',
    chatHeaderTips: 'Hello 👋！ Welcome to experience the {modelName} model',
    deepseekHeaderTips: "hello 👋！ Welcome to experience DeepSeek's comprehensive service for multi-scale models",
    //deepseekHeaderTips: "Hello! 👋！ Welcome to experience DeepSeek's intelligent Q&A service! We use intelligent scheduling technology to automatically match the model with the optimal parameter size based on the complexity of the problem, providing you with accurate and efficient answers",
    systemPlaceholder: 'Please enter the system persona, for example, "You are an AI assistant"',
    temperature: 'Temperature',
    top_p: 'Top_p',
    repetition_penalty: 'Repetition_penalty',
    max_tokens: 'Max_tokens',
    system_message: 'System_message',
    history: 'History',
    experienceStopTips: 'Task stopping, please wait!',
    experienceStopSuccess: 'Task stopped successfully, the experience service has been paused!',
    modelFinetuen: 'Model Finetune',
    pcMind: 'PengCheng · Mind',
    sdModelFinetuen: 'CV Finetune',
    sdModelFinetuenTask: 'CV Finetune Task',
    modelExperience: 'Model experience',
    modelExperienceTips: 'Provide online experience function for large models, where you can choose suitable models and computing resources according to your task scenarios, create online experience tasks for models, and test the response effect of the models online.',
    modelEvaluateTips: 'Evaluate the performance of the user fine-tuning generated model using the dataset provided by OpenCompass as the testing standard.',
    largeModelsList: 'Large models List',
    newModelExperience: 'New Model Experience Task',
    ModelExperienceDetail: 'Model Experience Task Details',
    ModelEvaluateDetail: 'Model Evaluation Task Details',
    modelEvaluate: 'Model evaluation',
    modelPerEvaluate: 'Model performance evaluation',
    modelSafeEvaluate: 'Model security evaluation',
    modelEvaluateTask: 'Model Performance Evaluation Task',
    evaluationModel: 'Evaluation Model',
    sftFinetune: 'NLP Finetune',
    sftFinetuneTask: 'NLP Finetune Task',
    newSftFinetune: 'Create NLP Finetune task',
    newLoraFinetune: 'Create CV Finetune task',
    newComfyUi: 'Create Comfy UI task',
    sftFinetuneDetail: 'Details of NLP Finetune task',
    sftFinetuneLimit: 'Each account can create 3 NLP Finetune tasks simultaneously',
    ModelExperienceLimit: 'Each account can create 3 Online Experience tasks simultaneously',
    ModelEvaluateLimit: 'Each account can create 1 Evaluation tasks simultaneously',
    sdLoraLimit: 'Each account can create 3 CV Finetune tasks simultaneously',
    comfyuiLimit: 'Each account can create 1 Comfy UI task simultaneously',
    sftFinetuneTips: 'At present, LoRA (Low Rank Adaptation) training mode is provided, which allows you to choose suitable training data and adjust training parameters according to your own task scenario, thereby achieving ideal model performance.',
    cvloraTips: 'LoRA is a lightweight model tuning method that trains specific network layer weights and inserts them into the base model to achieve fast, efficient, and minimal image training, ultimately optimizing the model parameter count and inference performance',
    ComfyUiTips: 'Comfy UI is a node-based graphical user interface (GUI) for Stable Diffusion, primarily designed to build highly flexible image generation workflows. It caters to diverse user groups, including digital artists, AI researchers, content creators, and more.',
    ComfyUiWarn: 'The service is starting up, please try again later!',
    uploadSound: 'Please upload a sound template',
    language: 'Language: ',
    soundStyle: 'Sound style: ',
    systemDefault: 'system Default',
    uploadSound1: 'upload sound template',
    dragDrop: 'Drag and drop audio to this location',
    outputAduio: 'output Aduio',
    appDev: 'Application Development',
    relatedTools: 'Related Tool Systems',
    inputLength: 'Input characters exceed the limit',
    ttsPlaceholder: 'Please input the text to be converted into speech, The Chinese and English speech synthesis should not exceed 82 and 250 characters respectively (Ctrl+Enter=line break, press Enter to submit)',
    soundTemplateTips: 'The sound template only supports MP3 and WAV format files',
    soundTemplateTips1: 'The sound template only supports a single file',
    soundTemplateTips3: 'Currently, only MP3 and WAV audio files are supported for uploading sound templates, and only a single file is supported',
    cvLoraTrain: 'CV Finetune',
    cvComfyui: 'Comfy UI',
    cvComfyuiTask: 'Comfy UI Task',
    modelArena: 'Model Arena',
    send: 'Send',
    newRound: 'New Round',
    pleaseSelectModelAndConversation: 'Select Model and Start Conversation',
    model_name: 'English name of model',
    model_zh_name: 'Chinese name of model'
  },
  modelFinetune: {
    foldParameters: 'Fold Parameters',
    expandParameters: 'Expand Parameters',
    epochs: 'Total number of training epochs to perform',
    learningRate: 'Initial learning rate for AdamW',
    batchSize: 'Number of samples processed on each GPU',
    valSize: 'Proportion of data in the dev set',
    evalSteps: 'Number of steps should be taken for verification',
    saveSteps: 'Number of steps between two checkpoints',
    maxsamples: 'Maximum samples per dataset',
    computeType: 'Whether to use mixed precision training',
    gradientAccumulation: 'Number of steps for gradient accumulation',
    lrScheduler: 'Name of the learning rate scheduler',
    maximumGradientNorm: 'Norm for gradient clipping',
    cutoffLen: 'Max tokens in input sequence',
    preprocessingNumWorkers: 'Number of processes used for processing',
    loggingSteps: 'Number of steps between two logs',
    warmupSteps: 'Number of steps used for warmup',
    packing: 'Pack sequences into samples of fixed length',
    optim: 'The optimizer to use: adamw_torch, adamw_8bit or adafactor',
    loraRank: 'The rank of LoRA matrices',
    loraAlpha: 'Lora scaling coefficient',
    loraDropout: 'Dropout ratio of LoRA weights',
    loraTarget: 'Name(s) of modules to apply LoRA. Use commas to separate multiple modules',
    additionalTarget: 'Name(s) of modules apart from LoRA layers to be set as trainable. Use commas to separate multiple modules',
    maxTaskTips: '<p><span>*</span>The platform only retains the results of finetune tasks for nearly<span> 30 </span> days. <span>Tasks over 30 days will not be able to download results and view logs, and cannot be finetune again</span></p>',
    maxEvalTaskTips: '<p><span>*</span>The platform only retains the results of evaluate tasks for nearly<span> 30 </span> days. <span>Tasks over 30 days will not be able to view results and logs</span></p>',
    max_epoch: 'Epoch',
    max_epoch_tips: 'Iteration rounds, control the number of iterations during the training process.',
    io_strategy: 'Model strategy',
    io_strategy_tips: '--',
    keep_interval_update: 'Checkpoint save interval',
    keep_interval_update_tips: '--',
    save_interval_update: 'Number of saved iteration steps',
    save_interval_update_tips: 'Every how many steps, save',
    lr: 'Learning rate',
    lr_tips: '--',
    vitual_pipeline_model_parallel_size: 'Virtual pipeline parallelism',
    vitual_pipeline_model_parallel_size_tips: '--',
    pipeline_model_parallel_size: 'Pipeline ',
    pipeline_model_parallel_size_tips: 'Pipeline parallelism, splitting the model according to the hierarchy of neurons, and placing different layers on different GPUs for computation',
    tensor_model_parallel_size: 'Tensor parallelism',
    tensor_model_parallel_size_tips: 'Tensor parallelism, which means combining multiple GPUs to perform tensor calculations simultaneously, such as matrix multiplication',
    num_micro_batch: 'Num micro batch',
    num_micro_batch_tips: 'The range of values for the number of micro batches is [1,32], which represents the accumulated number of micro batches in each training iteration. Setting a large number of micro batches can accelerate training, but it may cause memory issues',
    micro_batch_size: 'Micro batch size',
    micro_batch_size_tips: 'If it can only be 1 or 2, it will exceed the video memory and use gradient accumulation instead',
    recomputer_granularity: 'Recomputer granularity',
    recomputer_granularity_tips: 'Choosing to enable recalculation of compressible video memory usage may result in a decrease in throughput. It is recommended to close it if there is enough video memory available',
    modelTraining: 'model Training',
    modelTesting: 'model Testing',
    trainStart: 'started training',
    imageLabeling: 'image Captioning',
    totalPic: 'Total',
    numPic: 'pictures',
    clearAllPic: 'Clear all images',
    addPic: 'Add picture',
    captionMethod: 'captioning algorithm',
    captionthres: 'captioning threshold',
    modelTrigger: 'Model trigger word',
    modelTriggerph: 'Please enter the trigger word',
    captioning: 'Captioning',
    imgUploading: 'Picture uploading in progress',
    sureClearAllPic: 'Do you want to delete all images?',
    sureClearOnePic: 'Do you want to delete this image',
    clearAllPicSucc: 'All images have been successfully deleted',
    clearAllPicFail: 'Failed to delete all images',
    clearOnePicSucc: 'Image deleted successfully',
    clearOnePicFail: 'Image deletion failed',
    autoCaption: 'Automatically captioning images...',
    trainProgress: 'Training progress',
    currentSteps: 'Current Steps/Total Steps',
    trainingRounds: 'Training epochs/total epochs',
    trainParameters: 'Training parameters',
    viewLoss: 'View Loss',
    selected: 'Selected',
    startGenerateImg: 'Start generating images',
    sampleImage: 'Real time sample image',
    sampleImageGenerate: 'Sample image generation in progress',
    trainProcessing: 'During data processing and model loading...',
    trainingProgress: 'Training in progress',
    trainingCompleted: 'Training completed',
    loraModelName: 'LoRA model',
    picWidth: 'Width',
    picHeight: 'Height',
    loraRandomSeed: 'Random Seed',
    loraSamplingSteps: 'Sampling Steps',
    loraCFGScale: 'CFG Scale',
    loraNegativePrompt: 'Negative Prompt',
    loraInputPrompt: 'Please enter prompt!',
    capationAddPic: 'Add tags to all images',
    capationAddBegin: 'Add begin of the line',
    capationAddEnd: 'Add end of the line',
    capationAddMsg: 'Label input cannot be empty!',
    ProfessionalParameters: 'Professional Parameters',
    useBasicModel: 'Basic Model',
    Repeat: 'Repeat',
    Epoch: 'Epoch',
    totalSteps: 'total Steps',
    uploadImgCalc: 'Calculate after uploading the image',
    Prompt: 'Prompt',
    ProfessionalSetting: 'Professional Setting',
    loraPromptPlace: 'During the training process, real-time sample images will be generated based on prompt words',
    uploadImg: 'upload Images',
    uploadImgAccept: 'The file format does not meet the requirements. Please upload. jpg JPEG,. png format files',
    dragUploadImg: 'Drag and drop the image here to upload',
    uploadMaxImg: 'Add up to 200 images, supporting PNG/JPG/JPEG',
    uploadHavedImg: 'Upload existing image collection',
    uploadImgTips: 'Please ensure that the images correspond one-to-one with the marked files. The zip package cannot contain multi-level directories, and the total number of characters in the zip package name and file name should be less than 255',
    uploadImgTips1: 'File not detected',
    uploadImgTips2: 'The current maximum allowed upload is 200 image files',
    uploadImgTips3: 'The file format does not meet the requirements. Please upload an image in PNG/JPG/JPEG format',
    cvTaskfaildTips: 'The task has stopped or failed. Please recreate the task!',
    evalDatasetLimit: 'Evaluate the number of datasets',
    sftFinetuneName: 'fine-tuning task name',
    selectBaseModel: 'Select the basic model',
    selectFtModel: 'Select NLP fine-tuning model',
    evalTips: 'The model capability evaluation is based on the dataset provided by OpenCompass as the testing standard, for reference only.',
    evalTaskCategory: 'Task category',
    evalScore: 'Score',
    evalModelInput: 'Model input',
    evalModelOutput: 'Model output results',
    standerAnswer: 'Standard answer',
    evalModelOutAnswer: 'Model output answer',
    allResult: 'All result',
    correctResult: 'Correct result',
    wrongResult: 'Wrong result',
    allDataset: 'All Datasets'
  },
  userRole: {
    userRoleManagement: 'User role management',
    newOpRole: 'Create a new operational role',
    editOpRole: 'Edit operational role',
    viewOpRole: 'View operational role',
    newReRole: 'Create a new role for computing resources',
    editReRole: 'Edit computing power resources role',
    viewReRole: 'View computing resource roles',
    newStorageRole: 'Create a new storage resource role',
    editStorageRole: 'Edit storage resource role',
    viewStorageRole: 'View storage resource role',
    newContainerStorageRole: 'Create a new container storage quota role',
    editContainerStorageRole: 'Edit container storage quota role',
    viewContainerStorageRole: 'View container storage quota role',
    roleName: 'Role name',
    roleType: 'Role type',
    roleDescription: 'Role description',
    isItDefault: 'Is it default',
    viewDetail: 'View detail',
    "default": 'Default',
    notDefault: 'Not default',
    reCategory: 'Computing resources type',
    opCategory: 'Operational type',
    storageCategory: 'Storage resource type',
    allRoleType: 'All role type',
    deleteRoleConfirm: 'Deleting the default role will affect all users.',
    deleteRoleConfirm1: 'Are you sure you want to delete the current {name} role?',
    roleSelected: 'Please specify the permissions that the role has (multiple choices are allowed)',
    roleSelectedTips: 'Select All/Select None',
    allUserDefault: 'All user default',
    UserPermissionConfig: 'User permission configuration',
    OrgPermissionConfig: 'Organization permission configuration',
    allOpCategory: 'All operational roles',
    allReCategory: 'All computing power resource roles',
    allStCategory: 'All storage resource roles',
    userName: 'User name',
    OrganizationName: 'Organization name',
    opCategoryRole: 'Operational roles',
    reCategoryRole: 'Computing resource roles',
    stCategoryRole: 'Storage resource roles',
    codeSizeLimit: 'Code directory size limit',
    outputSizeLimit: 'Output directory size limit',
    pleaseEnterContent: 'Please enter the content',
    pleaseEnterRoleContent: 'Please enter the role name to search...',
    userId: 'User ID',
    organizationId: 'Organization ID',
    editUserPermissions: 'Edit user permissions',
    editOrgPermissions: 'Edit organization permissions',
    OperationalPermissions: 'Operational permissions',
    detail: 'Detail',
    permissionList: 'Permission List',
    updatePermissions: 'Update permissions',
    resourcePermissions: 'computing resource permission list',
    CResourcePermissions: 'Computing resource permissions',
    StoragePermissions: 'Storage resource permissions',
    permissionUpdateUnixDesc: 'Descending by user permission update time',
    permissionUpdateUnixAsc: 'Ascending by user update time',
    permissionCreatedUnixDesc: 'Descending by user creation time ',
    permissionCreatedUnixAsc: 'Ascending by user creation time ',
    orgPermissionCreatedUnixDesc: 'Descending by organization creation time ',
    orgPermissionCreatedUnixAsc: 'Ascending by organization creation time ',
    batchSetOrCancel: 'Batch setting or canceling user roles',
    enterUserId: 'Please enter the list of users who need to set permissions',
    failed: 'Failed',
    times: 'times',
    failedUserList: 'Failed user list',
    addAcessSuccess: 'Add access successfully',
    cancelAcessSuccess: 'Access cancelled successfully',
    operaNodeTips1: 'When running training tasks can use',
    operaNodeTips2: '-node',
    operaNodeTips3: 'computing resources',
    operaTaskTips1: 'Can run simultaneously',
    operaTaskTips2: 'task'
  },
  org: {
    organization: 'Organization',
    createOrg: 'Create organization',
    searchOrg: 'Search organization',
    orgName: 'Organization name',
    orgFullName: 'Full organization name',
    Members: 'Members',
    Teams: 'Teams',
    orgTeam: 'Teams',
    orgMember: 'People',
    numMember: 'member',
    numMembers: 'members',
    numRepo: 'repository',
    numRepos: 'repositories',
    unfold: 'Unfold',
    fold: 'Fold',
    searchRepos: 'Search Repositories',
    searchModels: 'Search Models',
    searchDatasets: 'Search Datasets',
    searchImages: 'Search Images',
    customize: 'Customize',
    selectedProjects: 'Selected Projects',
    selectedModels: 'Selected Models',
    selectedDatasets: 'Selected Datasets',
    customizeSelectedProjects: 'Customize selected projects',
    customizeSelectedModels: 'Customize selected models',
    customizeSelectedDatasets: 'Customize selected datasets',
    searching: 'Search...',
    remainNum: 'Remain: {num}',
    maxProjects: 'Select up to {num} public projects',
    maxModels: 'Select up to {num} public models',
    maxDatasets: 'Select up to {num} public datasets',
    order_newest: 'Newest',
    order_oldest: 'Oldest',
    order_recentupdate: 'Recently updated',
    order_leastupdate: 'Least recently updated',
    order_reversealphabetically: 'Reverse alphabetically',
    order_alphabetically: 'Alphabetically',
    order_moststars: 'Most stars',
    order_feweststars: 'Fewest stars',
    order_mostforks: 'Most forks',
    order_fewestforks: 'Fewest forks',
    no_result: 'No Data',
    myOrg: 'My Organization',
    orgDataset: 'Organization Datasets',
    orgModel: 'Organization Models'
  },
  storage: {
    capacity_details: 'Storage Capacity Details',
    quota: 'Storage Quota',
    data_usaged: 'Dataset Used',
    model_usaged: 'Model Used',
    remaining_available: 'Free Space',
    upload_time: 'Upload time',
    deleteModelConfirm: 'Are you sure to delete the current model <span style="color:red;">{name}</span> ？',
    deleteModelConSuccess: 'Successfully deleted the current model {name}!',
    deleteDataSetConfirm: 'Are you sure to delete the current dataset  <span style="color:red;">{name}</span> ？',
    deleteDataSetSuccess: 'Successfully deleted the current dataset {name}!',
    feedback_issue: 'Feedback issue',
    remain_storage: 'Remaining storage quota',
    selected_file_size: 'Selected file size',
    exceedStorage: 'Exceeded storage limit',
    deleteBatchData: 'Are you sure to delete the datasets selected in bulk? Once deleted, it cannot be restored',
    deleteBatchModel: 'Are you sure to delete the models selected in bulk? Once deleted, it cannot be restored',
    deletingTips: 'Files deletion in progress, please wait',
    OwenerTips: 'Please check the storage capacity details of {ownerName}',
    storageOrgTips1: 'To apply for more quotas, please click',
    storageOrgTips2: 'to download the template',
    storageOrgTips3: 'fill it out and send it to the email secretariat@openi.org.cn'
  },
  evalCaterary: {
    Examination: "Examination",
    Reasoning: "Reasoning",
    Knowledge: "Knowledge",
    Code: "Code",
    Math: "Math",
    Understanding: "Understanding",
    Language: "Language",
    Instruct: "Instruct"
  },
  thirdPartyApp: {
    thirdPartyApp: 'Third-party Application',
    oauth2App: 'OAuth2 Application',
    oauth2AppDescr: '(Verified OAuth2 application, after user authorization, the application can obtain the user\'s basic information (user ID, username, user avatar link), email, and phone number)',
    authStatus: 'Verify Status',
    canGetUserInfo: 'Authorized User Information',
    searchThirdPartyAppPlaceholder: 'Search for third-party application name',
    clientID: 'Client ID',
    appName: 'Application Name',
    appCanGetUserInfo: 'Authorized User Information',
    appCreator: 'Creator',
    authenticated: 'Verified',
    notAuthenticated: 'Not Verified',
    setAuthenticated: 'Set Verification',
    cancelAuthentication: 'Cancel Verification',
    basicInfo: 'Basic information',
    basicInfoAll: 'Basic information (user ID, username, user avatar link)',
    email: 'Email',
    phoneNumber: 'Phone number',
    editOauth2AppAuthInfo: 'Edit OAuth2 application authentication information',
    authenticatedCanGetInfoTips: 'After user authorization, the application can obtain the following user information',
    verify: 'Verify',
    cancelVerify: 'Cancel Verify',
    verifyAgain: 'Verify Again',
    cancelVerifyTips: 'After canceling verify, the application will only be able to obtain basic user information and will not be able to obtain the user\'s email or phone number. Do you want to continue?'
  },
  computingPowerObj: {
    computeResource: 'Compute Resources',
    computeResourceDescr: 'In order to meet the personalized computing power needs of users, OpenI Community can customize computing power pools according to their needs. Welcome to submit your computing power requirements.',
    freeComputingPower: 'Free Compute',
    paidComputingPower: 'Paid Compute',
    paidComputingPowerTip: 'Paid computing power and after-sales service provided by partners',
    createNewComputingTask: 'Create New Computing Task',
    createComputingTask: 'Create Computing Task',
    displayFilteringConditions: 'Display Filtering Conditions',
    priceRange: 'Price Range',
    point_hr: 'Point/hr',
    use: 'Ues Now',
    buyNow: 'Buy Now',
    provider: 'Provider',
    systemDisk: 'Dystem Disk',
    computingPowerPartner: 'Compute Partner',
    computingPowerOperationPlatform: 'Compute Platform',
    chipManufacturers: 'Chip Manufacturers',
    accCardCount: 'Acc Card Count'
  },
  dashboard: {
    welcomeTips: 'Welcome to OpenI AI Collaboration Platform',
    newUserGuide: 'New User Guide',
    quickAccess: 'Quick Access',
    createTask: 'Create computing task',
    taskDesc: 'Use computing cards to create tasks such as debugging, training, and inference.',
    myTaskCount: 'My computing tasks',
    runningTasks: 'Running',
    totalGPUHourd: 'Cumulative Runtime',
    createProject: 'Create repository',
    createRepoDesc: 'Create a code repository and use the code in computing tasks.',
    uploadDatasetDesc: 'Upload datasets and use them in computing tasks.',
    uploadModelDesc: 'Upload AI models and use them in computing tasks.',
    oneClickCreate: 'One-Click Creation',
    remainingQuota: 'Remaining Available Storage Quota',
    occupiedStorage: 'Occupied Storage',
    datasetStorage: 'Dataset Occupancy',
    modelStorage: 'Dataset Occupancy',
    credits: 'Calculation Points',
    currentAvailable: 'Current Available Calculation Points',
    currentAvailable1: 'Current Available',
    totalGained: 'Total Gained Calculation Points',
    totalGained1: 'Total Gained',
    totalConsumed: 'Total Consumed Calculation Points',
    totalConsumed1: 'Total Consumed',
    announcements: 'Announcements',
    moreAnnouncements: 'More Announcements',
    activities: 'Events',
    weChatLinked: 'WeChat Account Linked',
    weChatNotLinked: 'WeChat Account Not Linked',
    datasetNum: 'Datasets',
    modelNum: 'Models',
    unitPoints: '',
    unitTask: '',
    unitdata: '',
    runtimeCards: '',
    templateNum: 'Total templates',
    accumulatedRuns: 'Total runs',
    unitTemplate: '',
    heatmapLoading: 'Loading Heatmap…',
    contributionCount: 'contributions in the last 12 months：{count} times'
  },
  agentPortal: {
    taskRunningTip: 'You have {count} agents running, consuming {numPoint} points/hour',
    viewRunningAgents: 'View running agents',
    bannerTitle: 'Start Your AI Agent Journey',
    bannerDesc: 'Based on the powerful computing power of China Computing Network, deploy cutting-edge large model capabilities with one click to experience AI charm anytime, anywhere.',
    searchPlaceholder: 'Search agents',
    agentSquare: 'Agent Square',
    runningAgents: 'Running Agents',
    viewDetails: 'View Details',
    oneClickDeploy: 'One-click Deploy',
    deployFailed: 'Running Failed',
    deploying: 'Deploying',
    stop: 'Stop',
    startUsing: 'Start Using',
    startSuccess: 'Started Successfully',
    startFailed: 'Start Failed',
    stopSuccess: 'Stopped Successfully',
    stopFailed: 'Stop Failed',
    agentsRunning: 'Agent {names} is running',
    updated: 'Updated',
    noData: 'No Data'
  }
};
/* harmony default export */ __webpack_exports__["default"] = (en);

/***/ }),

/***/ "./web_src/vuepages/langs/config/zh-CN.js":
/*!************************************************!*\
  !*** ./web_src/vuepages/langs/config/zh-CN.js ***!
  \************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0__);


var _notebook;

var zh = {
  loading: "加载中...",
  noData: "暂无数据",
  noMessage: "暂无信息",
  date: "日期",
  noDataset: "空荡荡的，什么都没有",
  confirm: "确定",
  cancel: "取消",
  confirm1: "确认",
  confirmOp: "确认操作",
  cancelOp: "取消操作",
  enterReason: "请输入原因",
  clear: "清除",
  none: "无",
  pleaseCompleteTheInformationFirst: "请先完善信息！",
  submittedSuccessfully: "提交成功！",
  submittedFailed: "提交失败！",
  operation: "操作",
  edit: "修改",
  "delete": "删除",
  tips: "提示",
  expandMore: '展开更多',
  goBack: '返回上一级',
  star: '收藏',
  unStar: '取消收藏',
  submit: '提交',
  selectAll: '全选',
  selectNone: '全不选',
  copy: '复制',
  copySuccess: '复制成功',
  copyLink: '复制链接',
  copyedLink: '已复制链接',
  downloadAddress: '下载地址',
  operationFailed: '操作失败',
  cancelOperate: '您已取消操作',
  accomplishTask: "积分任务",
  adminOperate: "管理员操作",
  runCloudBrainTask: "运行计算任务",
  operating: "消耗中",
  succeeded: "已完成",
  debugTask: "调试任务",
  trainTask: "训练任务",
  trainDuration: "训练时长",
  inferenceTask: "推理任务",
  benchmarkTask: "评测任务",
  onlineinfer: "在线推理",
  onlineinferTask: "在线推理任务",
  consoleHome: "控制台",
  superComputeTask: "超算任务",
  generalTask: "通用任务",
  createPublicProject: "创建公开项目",
  dailyPutforwardTasks: "每日提出任务",
  dailyPR: "每日提出PR",
  comment: "发表评论",
  uploadDatasetFile: "上传数据集文件",
  importNewModel: "导入新模型",
  completeWechatCodeScanningVerification: "完成微信扫码验证",
  dailyRunCloudbrainTasks: "每日运行计算任务",
  datasetRecommendedByThePlatform: "数据集被平台推荐",
  aimodelRecommendedByThePlatform: "模型被平台推荐",
  dailyCreateDataset: "创建了数据集",
  dailyCreateDataset1: "创建了数据集",
  dailyCreateAimodel: "创建了模型",
  dailyCreateAimodel1: "创建了模型",
  submitNewPublicImage: "提交新公开镜像",
  imageRecommendedByThePlatform: "镜像被平台推荐",
  firstChangeofAvatar: "首次更换头像",
  dailyCommit: "每日commit",
  calcPointDetails: "算力积分明细",
  calcPointAcquisitionInstructions: "积分获取说明",
  CurrAvailableCalcPoints: "当前可用算力积分(分)",
  totalGainCalcPoints: "总获取算力积分(分)",
  totalConsumeCalcPoints: "总消耗算力积分(分)",
  gainDetail: "获取明细",
  consumeDetail: "消耗明细",
  serialNumber: "流水号",
  time: "时间",
  scene: "场景",
  behaviorOfPoint: "积分行为",
  explanation: "说明",
  points: "积分",
  status: "状态",
  runTime: "运行时长",
  taskName: "任务名称",
  createdRepository: "创建了项目",
  repositoryWasDel: "项目已删除",
  userWasDel: "用户已删除",
  userAccountWasDel: "用户账号已注销",
  openedIssue: "创建了任务",
  createdPullRequest: "创建了合并请求",
  commentedOnIssue: "评论了任务",
  uploadDataset: "上传了数据集文件",
  createdNewModel: "导入了新模型",
  invitedFriend: "邀请了好友",
  firstBindingWechatRewards: "首次绑定微信奖励",
  created: "创建了",
  type: "类型",
  dataset: "数据集",
  setAsRecommendedDataset: "被设置为推荐数据集",
  setAsRecommendedAimodel: "被设置为推荐模型",
  committedImage: "提交了镜像",
  image: "镜像",
  setAsRecommendedImage: "被设置为推荐镜像",
  updatedAvatar: "更新了头像",
  pushedBranch: "推送了{branch}分支代码到",
  deleteBranch: "从{repo}删除分支{branch}",
  pushedTag: "推送了标签{tag}到",
  deleteTag: "从{repo}删除了标签{tag}",
  dailyMaxTips: "达到每日上限积分，不能拿满分",
  memory: "内存",
  sharedMemory: "共享内存",
  ";": "；",
  noPointGainRecord: "还没有积分获取记录",
  noPointConsumeRecord: "还没有积分消耗记录",
  consumeDescr: '说明：从2024年10月下旬开始，计算任务扣费周期有调整，具体的计费规则请参考联机帮助中“<a target="_blank" href="https://openi.pcl.ac.cn/docs/index.html#/quickstart/resources?id=%e8%ae%a1%e8%b4%b9%e8%a7%84%e5%88%99">平台算力资源与积分说明</a>”页面的“计费规则”章节。',
  emptyPage: '您的访问受限！',
  emptyPageDescr: '您正尝试访问的页面 <strong>不存在</strong> 或 <strong>您尚未被授权</strong> 查看该页面。',
  handleTask: '处理任务',
  freeCompute: '普惠算力',
  repoPathTips: '路径只允许字母、数字和-_ .，最多100个字符。',
  useInPcWeb: '为获得更好的使用体验，请在PC端打开此页面。',
  acknowledgementsTips: {
    main1: '如果启智社区对您的科研工作提供了帮助，请在论文中给我们添加致谢（',
    main2: '）。',
    moreContent: '更多内容',
    p1: "\u82F1\u6587\u7248\uFF1AThanks for the support provided by OpenI Community (<a href=\"https://openi.pcl.ac.cn\" target=\"_blank\">https://openi.pcl.ac.cn</a>).",
    p2: "\u4E2D\u6587\u7248\uFF1A\u611F\u8C22\u542F\u667A\u793E\u533A\u63D0\u4F9B\u7684\u6280\u672F\u652F\u6301(<a href=\"https://openi.pcl.ac.cn\" target=\"_blank\">https://openi.pcl.ac.cn</a>)\u3002",
    p3: "\u5982\u679C\u60A8\u7684\u6210\u679C\u4E2D\u5F15\u7528\u4E86\u672C\u5E73\u53F0\uFF0C\u4E5F\u6B22\u8FCE\u5728\u4E0B\u8FF0\u5F00\u6E90\u9879\u76EE\u4E2D\u63D0\u4EA4\u60A8\u7684\u6210\u679C\u4FE1\u606F\uFF1A",
    p4: "<a href=\"https://openi.pcl.ac.cn/OpenIOSSG/references\" target=\"_blank\">https://openi.pcl.ac.cn/OpenIOSSG/references</a>"
  },
  computeResourceTitle: {
    CPU: 'CPU',
    VCPU: 'VCPU',
    'CPU/GPU': '英伟达GPU',
    GPU: '英伟达GPU',
    NPU: '昇腾NPU',
    GCU: '燧原GCU',
    MLU: '寒武纪MLU',
    DCU: '海光DCU',
    'ILUVATAR-GPGPU': '天数智芯GPGPU',
    'METAX-GPGPU': '沐曦GPGPU',
    'BIREN-GPU': '壁仞GPU'
  },
  TaskTypeTitle: {
    Notebook: '调试任务/在线推理',
    Notebook1: '调试任务',
    TrainJob: '训练任务',
    Ecs: '虚拟机',
    Inference: '推理任务',
    Service: '在线服务'
  },
  resourcesManagement: {
    OpenI: "启智集群",
    C2Net: "智算集群",
    IFLYTEKTraining: "讯飞集群",
    OpenIOne: "云脑一",
    OpenITwo: "云脑二",
    OpenIChengdu: "启智成都智算",
    chengduCenter: "成都智算",
    pclcci: "鹏城云计算所",
    hefeiCenter: "合肥类脑类脑智能开放平台",
    xuchangCenter: "中原人工智能计算中心",
    huoshi: "火石",
    willOnShelf: "待上架",
    onShelf: "已上架",
    offShelf: "已下架",
    toOnShelf: "上架",
    toOffShelf: "下架",
    toSetPriceAndOnShelf: "定价上架",
    status: "状态",
    reason: "原因",
    allStatus: "全部状态",
    syncAiNetwork: "同步智算网络",
    resQueue: "资源池（队列）",
    allResQueue: "全部资源池（队列）",
    addResQueue: "新建资源池（队列）",
    addResQueueBtn: "新增资源池",
    editResQueue: "修改资源池（队列）",
    resQueueCode: "资源池（队列）编码",
    resQueueName: "资源池（队列）名称",
    resQueueType: "资源池（队列）类型",
    resQueueIsAvailable: "资源池（队列）是否可用",
    resQueueIsAvailableAll: "资源池（队列）是否可用（全部）",
    allResQueueType: "全部资源池（队列）类型",
    whichCluster: "所属集群",
    allCluster: "全部集群",
    aiCenter: "智算中心",
    aiCenterID: "智算中心ID",
    allAiCenter: "全部智算中心",
    computeResource: "计算资源",
    allComputeResource: "全部计算资源",
    accCardType: "卡类型",
    allAccCardType: "全部卡类型",
    allNetworkType: "全部网络类型",
    cardsTotalNum: "卡数",
    accCardsNum: "卡数",
    allCardsNum: "全部卡数",
    remark: "备注",
    pleaseEnterRemark: "请输入备注(最大长度不超过255)",
    pleaseEnterPositiveIntegerCardsTotalNum: "请输入正整数的卡数！",
    addResSpecificationAndPriceInfo: "新增资源规格和单价信息",
    addResSpecificationBtn: "新增资源规格",
    editResSpecificationAndPriceInfo: "修改资源规格和单价信息",
    resSpecificationAndPriceManagement: "资源规格管理",
    sourceSpecCode: "对应资源编码",
    sourceSpecCodeTips: "云脑II需要填写对应的资源编码",
    sourceSpecId: "智算网络资源规格ID",
    cpuNum: "CPU数",
    gpuMem: "显存",
    mem: "内存",
    shareMem: "共享内存",
    unitPrice: "单价",
    point_hr: "积分/时",
    ShelfReason: "上架原因",
    node: "节点",
    free: "免费",
    onShelfConfirm: "请确认上架该规格？",
    offShelfConfirm: "请确认下架该规格？",
    onShelfCode1001: "上架失败，资源池（队列）不可用。",
    onShelfCode1003: "上架失败，资源规格不可用。",
    offShelfDlgTip1: "当前资源规格已在以下角色中被使用：",
    offShelfDlgTip2: "请确认进行下架操作？",
    characterLengthPrompt: "请输入不超过255个字符",
    resSceneManagement: "算力资源应用场景管理",
    addResScene: "新建算力资源应用场景",
    addResSceneBtn: "新增应用场景",
    editResScene: "修改算力资源应用场景",
    resSceneName: "应用场景名称",
    jobType: "任务类型",
    allJobType: "全部任务类型",
    allJobStatus: "全部任务状态",
    sceneType: "社区资源性质",
    allSceneType: "全部社区资源性质",
    isExclusiveSpec: "是否专属资源规格",
    allExclusiveAndCommonUseSpec: "全部专属和通用资源规格",
    "public": "共享(public)",
    exclusive: "独占(exclusive)",
    exclusiveTxt: "独占",
    exclusiveSpec: "专属资源规格",
    commonUseSpec: "通用资源规格",
    exclusiveOrg: "专属组织",
    exclusiveOrgTips: "多个组织名之间用英文分号隔开",
    computeCluster: "算力集群",
    resourceSpecification: "资源规格",
    lastUpdateTime: "最后更新时间",
    resSceneDeleteConfirm: "是否确认删除当前应用场景？",
    resourceSpecificationIsAvailable: "资源规格是否可用",
    resourceSpecificationIsAvailableAll: "资源规格是否可用（全部）",
    available: "可用",
    notAvailable: "不可用",
    selectSpec: '选择资源规格',
    accordingSpec: '按资源规格',
    accordingQueue: '按资源池队列',
    accordingSelect: '按已勾选',
    allIsEnableVisualization: '是否支持可视化（全部）',
    notEnable: '不支持',
    enable: '支持',
    roleNameTips: '最长输入80个字符',
    roleDescTips: '最长输入800个字符'
  },
  logManagement: {
    operationLog: "操作日志",
    allOperationType: "全部操作类型",
    resourceSpecificationOnshelf: '资源规格上架',
    resourceSpecificationEdit: '资源规格修改',
    resourceSpecificationOffshelf: '资源规格下架',
    resourceSpecificationAddition: '新增资源规格',
    resourceSpecificationAutoUpdate: '资源规格是否可用状态变化',
    operationType: '操作类型',
    operationReason: '操作原因',
    operator: '操作者',
    operationTime: '操作时间'
  },
  user: {
    inviteFriends: "邀请好友",
    inviteFriendsTips: "复制二维码或者注册邀请链接分享给好友",
    clickToViewTheEventDetails: "点击查看活动详情",
    copyRegistrationInvitationLink: "复制注册邀请链接",
    registrationAdress: "注册地址：",
    recommender: "推荐人：",
    invitedFriends: "已邀请好友",
    registrationTime: "注册时间",
    theSharedContentHasBeenCopiedToTheClipboard: "分享内容已复制到剪切板",
    copyError: "复制错误",
    Activated: "已激活",
    notActive: "未激活",
    normal: '正常'
  },
  notebook: (_notebook = {
    createNewNotebook: "新建Notebook调试任务",
    sameTaskTips1: "您已经有",
    sameTaskTips2: "同类任务",
    sameTaskTips3: "正在等待或运行中，请等待任务结束再创建",
    sameTaskTips4: "可以在",
    sameTaskTips5: "我的工作台",
    sameTaskTips6: "计算任务",
    sameTaskTips7: "查看您所有的计算任务。",
    cpuEnv: "CPU 环境",
    gpuEnv: "GPU 环境",
    npuEnv: "NPU 环境",
    newTask: "新建任务",
    noQuene: "当前资源充足，没有任务排队",
    queneTips1: "您当前排队位置是第",
    queneTips2: "位",
    watiResource: "正在等待分配资源，请耐心等待",
    debug: "调试",
    stop: "停止",
    stopping: "停止中",
    notebookRunning: "调试环境已启动，单次连接 4 小时后自动关闭",
    stopSuccess: "停止任务成功",
    specification: "规格",
    graphicMemory: "显存",
    memory: "内存",
    sharedMemory: "共享内存",
    tips: '本次新建的调试任务会放在您名下项目openi-notebook中，如果没有该项目系统会自动新建一个。',
    balanceOfPoints: '积分余额：',
    points: '积分',
    hours: '小时',
    expected_time: '，预计可用',
    limitTitle: '{type}数量已达上限',
    limitReason: '当前账号最多创建{count}个和同时运行{runCount}个调试任务，请停止运行中的任务并删除部分历史任务后提交新任务',
    stopReason: '当前账号已有{count}个正在等待或运行的任务，请停止后再提交新任务',
    deleteReason: '当前账号最多创建{count}个{type}，请删除历史任务后提交新任务'
  }, _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_notebook, "stopSuccess", '停止任务成功'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_notebook, "stopFailed", '停止任务失败'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_notebook, "deleteSuccess", '删除任务成功'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_notebook, "deleteFailed", '删除任务失败'), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_0___default()(_notebook, "handleSuccess", '操作成功，请继续创建任务'), _notebook),
  tranformImageFailed: '图片脱敏失败',
  originPicture: '原始图片',
  desensitizationPicture: '脱敏图片',
  desensitizationObject: '脱敏对象',
  example: '示例',
  startDesensitization: '开始处理',
  all: '全部',
  others: '其它',
  onlyFace: '仅人脸',
  onlyLicensePlate: '仅车牌',
  dragThePictureHere: '拖动图片到这里',
  or: '或',
  clickUpload: '点击上传',
  dataDesensitizationModelExperience: '数据脱敏模型体验',
  dataDesensitizationModelDesc: '利用人工智能AI技术，把图片中的人脸、车牌号码进行脱敏处理。该模型更多信息请访问项目',
  limitFilesUpload: '只能上传 jpg/jpeg/png 格式的文件',
  limitSizeUpload: '上传文件大小不能超过 20M !',
  bootPlaceholder: '请输入启动文件',
  modelManage: {
    modelManage: '模型管理',
    modelName: '模型名称',
    useCluster: '可用集群',
    license: '许可证',
    selectLicense: '选择许可证',
    local: '本地',
    online: '线上',
    external: '外部',
    createModel: '创建模型',
    importLocalModel: '导入本地模型',
    importOnlineModel: '导入线上模型',
    importFrom: '从哪里导入',
    createNewModel: '新建模型',
    localModelTips: '将本地存放的模型文件导入至启智AI协作平台',
    onlineModelTips1: '将训练任务输出的模型导入至模型管理中。导入线上模型需先',
    onlineModelTips2: '创建训练任务',
    nextSteps: '下一步',
    back: '上一步',
    step1: '选择导入类型和项目',
    step2: '设置模型信息',
    step3: '上传模型文件',
    project: '所属项目',
    newRepoPath: '新项目路径',
    modelInfo: '模型信息',
    modelBriefIntro: '模型简介',
    trainingInfo: '训练相关信息',
    modifyModelInfo: '修改模型信息',
    modelList: '模型列表',
    modelFiles: '模型文件',
    addModelFiles: '增加模型文件',
    uploadModelFiles: '上传模型文件',
    pleaseInputModelName: '请输入模型名称',
    version: '版本',
    modelVersion: '模型版本',
    modelEngine: '模型框架',
    modelLabel: '模型标签',
    externalToken: '外部Token ',
    modelSource: '模型来源',
    modelLabelInputTips: '输入标签，多个标签用空格区分',
    modelLabelGatedTips: '迁移Gated的模型文件需要输入HuggingFace的Token',
    modelDescr: '模型描述',
    modelDescrInputTips: '描述字数不超过255个字符',
    confirm: '确定',
    cancel: '取消',
    modelCreateFailed: '模型创建失败',
    modelModifyFailed: '模型修改失败',
    fileUpload: '文件上传',
    folderUpload: '文件夹上传',
    upload: '上传',
    uploadPath: '上传路径',
    uploadStatus: '上传状态',
    modelFileUploadDefaultTips: '点击添加文件或直接拖拽文件到此处',
    modelFileUploadErrTips: '文件名不超过128个字符，单次最多上传 {maxCount} 个文件，单个文件超过{size}请使用<a target="_blank" href="{url}" style="cursor: pointer;">SDK上传</a></a>, 文件数量较多时，建议压缩后再上传',
    modelFileNameTips: '文件名长度不超过128个字符',
    modelFileNamSspaceTips: '文件名包含空格',
    fileIstoBig: '文件大小（{{filesize}} MiB）超过了最大允许大小（{{maxFilesize}} MiB）',
    removeFile: '移除文件',
    uploadSuccess: '上传成功',
    uploadFailed: '上传失败',
    calcFileMd5: '计算文件MD5...',
    uploading: '上传中...',
    fileHasAlreadyInTheModel: '该文件已上传在模型：',
    fileEqualInTheModel: '该文件与其它文件内容相同：',
    fileExistInTheModel: '已存在相同名称文件：',
    uploadFile: '上传文件',
    uploadFolder: '上传文件夹',
    addUploadFolder: '添加上传文件夹',
    clearAll: '清空所有',
    fileCountAndSize: '已选择文件数: {count}，总大小: {size}',
    folderUploadPlaceholder: '点击添加文件夹或直接拖拽文件夹/文件到此处',
    folderUploadTips: '本次上传的文件数量不超过 {maxCount} 个，总大小不超过 {maxSize}，超多超大的文件请通过 <a target="blank" href="https://openi.pcl.ac.cn/docs/index.html#/api/list">SDK</a> 方式上传。',
    folderUploadErrorTips: '本次上传的文件数量不超过 {maxCount} 个，单个文件超过 {maxSize}请使用<a target="_blank" href="{url}" style="cursor: pointer;">SDK上传</a>',
    basicInfo: '基本信息',
    modelSize: '模型大小',
    descr: '描述',
    createTime: '创建时间',
    label: '标签',
    trainTaskInfo: '训练相关信息',
    trainTask: '训练任务',
    codeBranch: '代码分支',
    bootFile: '启动文件',
    viewSamples: '查看样例',
    trainDataset: '训练数据集',
    datasetfile: '数据集文件',
    fileShort: '文件',
    specInfo: '规格',
    workServerNumber: '计算节点',
    runParameters: '运行参数',
    seeMore: '查看更多信息',
    collapseDetails: '折叠详细信息',
    modelFilesList: '模型文件列表',
    fileName: '文件名称',
    fileSize: '文件大小',
    updateTime: '更新时间',
    operate: '操作',
    download: '下载',
    "delete": '删除',
    infoModificationFailed: '信息修改失败',
    deleteModelFileConfirmTips: '请确认是否删除当前模型文件？',
    modelFileDeleteFailed: '模型文件删除失败',
    modelAccess: '模型权限',
    modelAccessPublic: '公开',
    modelAccessPrivate: '私有',
    modelAccessTips: '公开项目才可以设置模型为公开',
    modelSettings: '模型信息设置',
    edit: '编辑',
    editFiles: '编辑文件',
    preview: '预览',
    modelIntroduction: '模型介绍',
    hasNoIntroForModel: '还没有详细的模型介绍',
    createModelIntro: '创建模型介绍',
    briefIntroduction: '简介',
    addLabels: '新增标签',
    discardFileChanges: '是否放弃当前内容修改？',
    editFileContentFirst: '请先编辑文件内容！',
    ownerRepository: '所属项目',
    creator: '创建者',
    migrator: '迁移者',
    useModel: '使用模型',
    modelEvolutionMap: '模型演化图谱',
    evolutionMap: '图谱',
    settings: '设置',
    parentModel: '父模型',
    currentModel: '当前模型',
    derivedModel: '衍生模型',
    publicDerivedModel: '公开衍生模型',
    privateDerivedModel: '私有衍生模型',
    refRepository: '引用项目',
    publicRefRepository: '公开引用项目',
    privateRefRepository: '私有引用项目',
    modelDownloadAll: '下载全部模型文件',
    otherOnline: '其他在线体验入口',
    trainUsedDataList: '和该模型一起使用过的数据集',
    trainUsedRepo: '和该模型一起使用过的代码',
    modelUseTaskList: '使用该模型的任务',
    forkModelSuccess: '模型内容复制完成，项目fork成功！',
    debugModel: '调试模型',
    onlineInference: '在线体验',
    onlineInferenceTask: '在线体验任务',
    onlineLoraTrain: '在线训练',
    onlineWorkflow: '在线工作流',
    deleted: '已删除',
    derivativeTimes: '衍生次数',
    derivativeTimes1: '衍生',
    mostDerivative: '最多衍生',
    modelBaseInfo: '模型基本信息',
    modelCollaborator: '模型协作者',
    managementTeam: '模型管理团队',
    add_team_success: '团队现在可以访问模型',
    remove_team_success: '团队访问模型的权限已被删除。',
    deletCollaboratedTips: '删除该协作者后，他将无法访问此模型，继续操作吗？',
    total_size: '总大小',
    total_size_asc: '总大小倒序'
  },
  repos: {
    activeOrganization: '活跃组织',
    activeUsers: '活跃用户',
    follow: '关注',
    unFollow: '取消关注',
    selectedFields: '领域精选',
    mostPopular: '近期热门',
    mostActive: '近期活跃',
    newest: '最近创建',
    recentlyUpdated: '最近更新',
    mostStars: '点赞最多',
    mostForks: '派生最多',
    mostDatasets: '数据集最多',
    mostAiTasks: 'AI任务最多',
    mostModels: '模型最多',
    repos: '项目',
    publicRepos: '公开项目',
    repoTopics: '项目领域',
    dataset: '数据集',
    model: '模型',
    aiTask: 'AI任务',
    updated: '最后更新于',
    contributors: '贡献者',
    searchRepositories: '搜项目...',
    search: '搜索',
    allFields: '全部领域',
    preferred: '项目优选',
    openIIncubation: '启智孵化管道',
    hotPapers: '热门论文项目',
    watch: '关注',
    star: '点赞',
    fork: '派生',
    noReposfound: '未找到匹配的项目。',
    source: '自建',
    mirrors: '镜像',
    collaborative: '协作',
    forks: '派生',
    selectRepo: '选择项目',
    selectRepoPlaceholder: '搜索项目名称...',
    cancelToping: '取消置顶',
    topping: '置顶',
    toppingSuccess: '置顶成功！',
    cancancelTopingSuccess: '取消置顶成功！',
    deleteThisRepos: '删除本项目',
    delete_notices_2: '- 此操作将永久删除项目 <strong>{repos}</strong>。包括该项目中的代码、任务、合并请求等内容',
    sureReposName: '输入项目名称以做确认：',
    repos_name1: '项目名称',
    deleteReposNameError: '项目名称输入错误！'
  },
  timeObj: {
    ago: '{msg}前',
    from_now: '{msg} 之后',
    now: '现在',
    future: '将来',
    '1s': '1 秒',
    '1m': '1 分钟',
    '1h': '1 小时',
    '1d': '1 天',
    '1w': '1 周',
    '1mon': '1 个月',
    '1y': '1 年',
    seconds: '{msg} 秒',
    minutes: '{msg} 分钟',
    hours: '{msg} 小时',
    days: '{msg} 天',
    weeks: '{msg} 周',
    months: '{msg} 个月',
    years: '{msg} 年',
    raw_seconds: '秒',
    raw_minutes: '分钟'
  },
  modelObj: {
    model_label: '选择模型',
    model_select_placeholder: '选择模型',
    model_export_placeholder: '请选择模型文件',
    model_select: '选择模型',
    model_current_repo: '本项目',
    model_my: '我拥有的',
    model_collaborate: '我协作的',
    model_collected: '我收藏的',
    model_my_migrate: '我迁移的',
    model_public: '公开模型',
    model_recommend: '平台推荐模型',
    model_migrate: '外部迁移模型',
    model_search_placeholder: '搜索模型名称...',
    model_selected: '已选模型',
    model_ok: '确定',
    model_most: '最多不超过 {msg} 个模型',
    model_not_equal_file: '不能选择相同名称的模型',
    model_exceeds_failed: '模型大小超过',
    model_should_same_model: '请选择同一模型下的文件',
    model_search: '搜索模型',
    model_square_empty: '空荡荡的，什么都没有',
    model_suport_file_tips: '模型文件支持的格式为 [ckpt, pb, h5, json, pkl, pth, t7, pdparams, onnx, pbtxt, keras, mlmodel, cfg, pt]',
    boot_file_helper: '启动文件是您程序执行的入口文件，必须是以.py结尾的文件。比如train.py、main.py、example/train.py、case/main.py。',
    can_online_infer: '可体验',
    can_fine_tune: '可微调',
    model_sdk_use_way: '模型SDK使用方式',
    codeUseDlgTriggerTxt: '在OpenI如何使用模型',
    codeUseDlgTitle: '如何在OpenI协作平台使用模型',
    codeDownDlgTitle: '如何在OpenI协作平台下载文件',
    migrate_external_model: '迁移外部模型',
    external_model_name: '外部模型名称',
    migrate_external_model_name_tips: '目前仅支持迁移HuggingFace的模型，请输入其模型名称，如：THUDM/chatglm3-6b',
    openi_model_repo_url: '启智模型仓库URL地址',
    migrate_model_should_be_public: '迁移的模型须公开',
    migrate_model: '迁移模型',
    please_enter_right_external_model_name: '请输入正确的外部模型名称',
    migrate_external_model_failed: '迁移外部模型失败！',
    create_model_migrate_loading_content: '创建模型迁移中，请稍后~',
    create_model_migrate_exists: '迁移的模型已存在，查看 <a href="{url}">模型详情</a>。',
    migrate_from: '正在从 {url} 迁移',
    migrate_watting: '等待迁移',
    migrating: '迁移中',
    migrate_success: '迁移成功',
    migrate_failed: '迁移失败',
    migrate_failed_tips: '模型迁移失败！请',
    migrate_retry: '重新迁移',
    migrate_try_later: '，或者稍后再试',
    re_migrate_failed: '重新迁移失败！',
    sync_now: '立即同步',
    model_sync: '模型同步更新',
    current_size: '当前大小',
    new_size: '最新大小',
    status_add: '新增',
    status_del: '删除',
    status_update: '有更新',
    start_sync: '开始同步',
    model_no_update_tips: '模型文件无变化，无需更新',
    model_sdk_tips1: '## 下载启智平台文件到本地文件夹：\n *文件大于{fileSize}时使用openi SDK下载，请根据具体的 **本地保存路径** 修改下列代码中的参数* \n#### 使用Python代码下载',
    model_sdk_tips2: '#### 使用命令行下载',
    model_sdk_tips3: '\n关于 `openi` SDK库的安装及使用方式请参考 <a target="_blank" href="https://openi.pcl.ac.cn/docs/index.html#/model/sdk">使用帮助</a>。',
    model_download: '下载',
    model_source: '源模型：',
    model_name1: '模型名称',
    deleteThisModel: '删除本模型',
    deleteThisModelTips: '删除模型是永久性的，无法撤销的，请您谨慎操作。',
    delete_notices_2: '- 此操作将永久删除模型 <strong>{model}</strong>。',
    sureModelName: '输入模型名称以做确认：',
    deleteModelNameError: '模型名称输入错误！'
  },
  datasetObj: {
    dataset_label: "数据集",
    dataset_select_placeholder: "选择数据集",
    eval_select_placeholder: "最多选择5个评测数据集",
    dataset_select: "选择数据集",
    dataset_search_placeholder: "搜数据集名称",
    dataset_unziping: "正在解压缩",
    dataset_success: "解压成功",
    dataset_unzip_failed: "解压失败",
    dataset_status: '解压状态',
    dataset_exceeds_failed: "数据集大小超过",
    dataset_my_upload: "我上传的",
    dataset_current_repo: "本项目",
    dataset_public: "公开数据集",
    dataset_relate: "关联数据集",
    dataset_collected: "我收藏的",
    dataset_selected: "已选数据集",
    dataset_ok: "确定",
    dataset_not_equal_file: "不能选择相同名称的数据集",
    dataset_most: "最多不超过 {msg} 个数据集",
    dataset_file_was_deleted: "文件已经被删除",
    dataset_sdk_use_way: '数据集SDK使用方式',
    codeUseDlgTriggerTxt: '在OpenI如何使用数据集',
    codeUseDlgTitle: '如何在OpenI协作平台使用数据集',
    create_new_dataset: '新建数据集',
    dataset_name1: '数据集名称',
    dataset_name: '数据集英文名称',
    dataset_zh_name: '数据集中文名称',
    dataset_description: '数据集描述',
    select_category: '选择分类',
    task: '研究方向/应用领域',
    application: '应用领域',
    select_task: '选择研究方向/应用领域',
    dataset_name_tooltips: '请输入字母、数字和-_ .，以数字或字母开头和结尾，最长100个字符。',
    dataset_name_cn_tooltips: '请输入中文、字母、数字和-_ .，最长100个字符。',
    category: '分类',
    license: '授权许可',
    license_helper: '选择授权许可文件',
    dataset_name_required1: '请输入数据集中文名称',
    dataset_name_required2: '输入不符合数据集名称规则',
    dataset_name_required3: '请输入数据集英文名称',
    dataset_description_required: '请输入数据集描述详情',
    category_required: '请选择分类',
    select_task_required: '请选择研究方向/应用领域',
    dataset_repo_tips: '因每个项目只能创建一个数据集，此处仅能指定未创建数据集的项目',
    exampleDataset: '示例数据集',
    selectOpeniDataset: '选择平台数据集',
    buildDatasetTips: '如何构建数据集',
    create_public_tips: '*默认创建的数据集为公开数据集，可在数据集页面将其设置为私有数据集',
    dataset_owner: '拥有者',
    name: '名称',
    dataset_acess: '数据集权限',
    dataset_size: '数据集大小',
    hasNoIntroForModel: '还没有详细的数据集介绍',
    createModelIntro: '创建数据集介绍',
    dataset_intro: '数据集介绍',
    use_dataset: ' 使用数据集',
    deleteDataFileConfirmTips: '请确认是否删除当前文件 {name}？',
    dataFileDeleteFailed: '文件删除失败',
    dataDownloadAll: '下载全部数据集文件',
    createDataset: '创建数据集',
    deletCollaboratedTitle: '删除协作者',
    deletCollaboratedTips: '删除该协作者后，他将无法访问此数据集，继续操作吗？',
    dataBaseInfo: '数据集基本信息',
    dataSearch: '搜索中...',
    notFoundUser: '未找到匹配的',
    administrators: '管理员',
    writePermission: '可写权限',
    readPermission: '可读权限',
    datasetCollaborator: '数据集协作者',
    searchUsers: '搜索用户',
    addCollaborators: '增加协作者',
    managementTeam: '数据集管理团队',
    searchTeams: '搜索团队',
    addTeam: '添加团队',
    dangerousOperationZone: '危险操作区',
    deleteThisDataset: '删除本数据集',
    deleteThisDatasetTips: '删除数据集是永久性的，无法撤销的，请您谨慎操作。',
    sureDatasetName: '输入数据集名称以做确认：',
    deleteDatasetNameError: '数据集名称输入错误！',
    add_collaborator_success: '协作者添加成功！',
    add_team_success: '团队现在可以访问数据集',
    remove_collaborator_success: '协作者删除成功！',
    remove_team_success: '团队访问数据集的权限已被删除。',
    delete_notices_1: '- 此操作 <strong>不可以</strong> 回滚。',
    delete_notices_2: '- 此操作将永久删除数据集 <strong>{dataset}</strong>。'
  },
  imagesObj: {
    cloudbrain_images: '镜像',
    openIGPU: '启智GPU',
    c2netGPU: '智算GPU',
    images_search: '搜镜像Tag/描述/操作系统/安装的软件包/标签...',
    image_public: '平台推荐镜像',
    image_my: '我的镜像',
    image_collected: '我收藏的镜像',
    framework: '框架',
    version: '版本',
    frameworkName: '框架名称',
    frameworkVersion: '框架版本',
    pyVersion: 'Python版本',
    cudaVersion: 'Cuda版本',
    cannVersion: 'Cann版本',
    dtkVersion: 'Dtk版本',
    deleteTips: '你确认删除该镜像么？此镜像一旦删除不可恢复。',
    deleteSuccessTips: '删除成功',
    editImage: '修改镜像',
    submitImage: '提交镜像',
    appyImage: '申请成为平台推荐镜像',
    imageTag: '镜像Tag',
    imageDesc: '镜像描述',
    imageTagPlaceholder: '请输入镜像Tag',
    imageTagInputTips: '请输入字母、数字、_-.，最长50个字符，且以字母开头。',
    imageAdress: '镜像地址',
    imageAdressPlaceholder: '请输入镜像地址',
    imageAiCenterPlaceholder: '请输入智算中心JSON格式，如：\n[{"aiCenterId":"","trainJobUrl":"","notebookUrl":"","accDeviceModel":"",poolIds:""}]',
    operationSystem: '操作系统',
    operationSystemName: '操作系统名称',
    operationSystemNamePlaceholder: '请输入操作系统名称',
    operationSystemVersionPlaceholder: '请输入操作系统版本',
    pyPackge: 'Python依赖库',
    thirdPackages: 'Python依赖库及版本',
    thirdPackagesPlaceholder: "\u8BF7\u8F93\u5165\u5B89\u88C5\u7684\u8F6F\u4EF6\u5305\u53CA\u7248\u672C\uFF0C\u4E00\u884C\u4E00\u4E2A\uFF0C\u6700\u591A10\u884C\u3002\u5982\uFF1A\ntorch==1.13.1\ntransformers==4.25.1",
    topic: '标签',
    topicPlaceholder: '输入完成后回车键完成标签确定',
    topicTips: "\u8BF7\u5728\u6807\u7B7E\u5B57\u6BB5\u8865\u5145\u955C\u50CF\u4E2D\u5176\u4ED6\u4FE1\u606F\uFF0C\u5982\uFF1A\u9002\u914D\u7684\u5361\u7C7B\u578B\uFF1A<span class=\"light\">T4</span>\u3001<span class=\"light\">Ascend 910</span>",
    descr: '镜像简介',
    descrPlaceholder: '请输入镜像简介，不超过1000个字符',
    recommend: '推荐',
    notRecommend: '不推荐',
    submitTips: '代码目录/tmp/code不会随镜像提交，其他目录下新增的文件(夹)都会打包到镜像中。',
    submitNpuTip1: '目录/home/ma-user/work不会随镜像提交，其他目录都会打包到镜像中。',
    submitNpuTip2: '镜像保存过程中，Notebook将不可用，本任务状态也将由Running变为Waiting。',
    submitDcuTip1: '目录/home/ma-user/work不会随镜像提交，其他目录都会打包到镜像中。',
    submitDcuTip2: '镜像小于40G才能保存成功。镜像30天未使用系统将会清理，无法使用',
    submitApply: '提交申请',
    imageCommitting: '镜像提交中...',
    imageCommitSuccess: '镜像提交成功',
    imageCommitErrorTips1: '检测提交镜像是否大小超过20G!',
    imageCommitErrorTips2: '检测提交镜像是否大小超过40G!',
    committing: '提交中',
    commitSuccess: '提交成功',
    commitFailed: '提交失败',
    recommendNeedReview: '推荐待审核',
    recommendReviewApproved: '通过推荐审核',
    recommendReviewFailed: '未通过推荐审核',
    applyForRecommend: '申请推荐',
    copyAdress: '复制地址',
    filterImages: '筛选镜像',
    filterImagesPlaceholderCuda: '框架名称/框架版本/Python版本/Cuda版本',
    filterImagesPlaceholderCann: '框架名称/框架版本/Python版本/Cann版本',
    filterImagesPlaceholderDtk: '框架名称/框架版本/Python版本/Dtk版本',
    imageTaskType: '适用任务类型',
    imageTaskTips: '适用于',
    approval_status: '评审状态',
    all_approval_status: '全部评审状态',
    pending_approval: '待审核',
    approved: '通过审核',
    not_approved: '未通过审核',
    none: '无',
    create_cloud_brain_mirror: '创建云脑镜像',
    not_recommend: '不同意推荐'
  },
  specObj: {
    resSelectTips: '「资源规格」是您运行该任务使用的硬件，为了更多人能够使用本平台的资源，请按照您的实际需求进行选择。',
    no_use_resource: '暂无可用资源'
  },
  datasets: {
    computer_vision: "计算机视觉",
    natural_language_processing: "自然语言处理",
    speech_processing: "语音处理",
    computer_vision_natural_language_processing: "计算机视觉、自然语言处理",
    machine_translation: "机器翻译",
    medical_imaging: "医学影像",
    question_answering_system: "问答系统",
    information_retrieval: "信息检索",
    knowledge_graph: "知识图谱",
    text_annotation: "文本标注",
    text_categorization: "文本分类",
    emotion_analysis: "情感分析",
    language_modeling: "语言建模",
    speech_recognition: "语音识别",
    automatic_digest: "自动文摘",
    information_extraction: "信息抽取",
    description_generation: "说明生成",
    image_classification: "图像分类",
    face_recognition: "人脸识别",
    image_search: "图像搜索",
    target_detection: "目标检测",
    image_description_generation: "图像描述生成",
    vehicle_license_plate_recognition: "车辆车牌识别",
    medical_image_analysis: "医学图像分析",
    unmanned: "无人驾驶",
    unmanned_security: "无人安防",
    drone: "无人机",
    vr_ar: "VR/AR",
    "2_d_vision": "2-D视觉",
    "2_5_d_vision": "2.5-D视觉",
    "3_d_reconstruction": "3D重构",
    image_processing: "图像处理",
    video_processing: "视频处理",
    visual_input_system: "视觉输入系统",
    speech_coding: "语音编码",
    speech_enhancement: "语音增强",
    speech_synthesis: "语音合成",
    ros_hmci_datasets: "开源开放社区",
    downloadtimes: '下载次数',
    downloadtimes1: '下载',
    citations: '引用次数',
    citations1: '引用',
    "default": "默认排序",
    newest: "最新创建",
    oldest: "最早创建",
    recentupdate: "最近更新",
    leastupdate: "最少更新",
    moststars: "收藏数量",
    mostCollections: "最多收藏",
    mostusecount: "最多引用",
    alphabetasc: "字母升序",
    alphabetdesc: "字母降序",
    chinesename: "中文名",
    chinesenameasc: "中文名倒序",
    size: "大小",
    sort: "排序",
    unstarSuccess: "取消收藏成功！",
    starSuccess: "收藏成功！",
    platform_recommendations: '仅显示平台推荐',
    category: '分类',
    task: '研究方向/应用领域',
    license: '授权许可',
    publick_dataset: '公开数据集',
    my_dataset: '我拥有的',
    favorite_dataset: '我收藏的',
    recommend_dataset: '平台推荐数据集',
    collaborated_dataset: '我协作的',
    views: '查看',
    loadMore: '加载更多'
  },
  cloudbrainObj: {
    cloudbrain: '云脑',
    openi: '启智集群',
    c2net: '智算网络集群(Beta)',
    create: '新建',
    createTask: '新建任务',
    cluster: '算力集群',
    computeResource: '计算资源',
    sameTaskTips1: '您已经有 {count}个同类任务正在等待或运行中，已达到最大值，请停止任务再创建。',
    sameTaskTips2: '可以在 “<a href="/cloudbrains" target="_blank">我的工作台 &gt; 计算任务</a>” 查看您所有的计算任务。',
    pathTips1: '训练脚本存储在 <strong style="color:#010101">{code}</strong> 中，数据集存储在 <strong style="color:#010101">{dataset}</strong> 中，预训练模型存放在运行参数 <strong style="color:#010101">{model}</strong> 中，训练输出请存储在 <strong style="color:#010101">{output}</strong> 中以供后续下载。',
    pathTips11: '训练脚本存储在 <strong style="color:#010101">{code}</strong> 中，数据集存储在 <strong style="color:#010101">{dataset}</strong> 中，预训练模型存储在 <strong style="color:#010101">{model}</strong> 中，训练输出请存储在 <strong style="color:#010101">{output}</strong> 中以供后续下载。',
    pathTips2: '项目代码存储在 <strong style="color:#010101">{code}</strong> 中，数据集存储在 <strong style="color:#010101">{dataset}</strong> 中，选择的模型存储在 <strong style="color:#010101">{model}</strong> 中，调试输出请存储在 <strong style="color:#010101">{output}</strong> 中以供后续下载。',
    pathTips3: '项目代码存储在 <strong style="color:#010101">{code}</strong> 中，数据集存储在 <strong style="color:#010101">{dataset}</strong> 中，选择的模型存储在 <strong style="color:#010101">{model}</strong> 中。',
    pathTips4: '训练脚本存储在 <strong style="color:#010101">{code}</strong> 中，预训练模型存放在运行参数 <strong style="color:#010101">{model}</strong> 中，训练输出请存储在 <strong style="color:#010101">{output}</strong> 中以供后续下载。',
    pathTips5: '数据集存储在 <strong style="color:#010101">{dataset}</strong> 中，模型文件存储在 <strong style="color:#010101">{model}</strong> 中，推理输出请存储在 <strong style="color:#010101">{output}</strong> 中以供后续下载。',
    pathTips6: '数据集位置存储在运行参数 <strong style="color:#010101">{dataset}</strong> 中，预训练模型存放在运行参数 <strong style="color:#010101">{model}</strong> 中，训练输出路径存储在运行参数 <strong style="color:#010101">{output}</strong> 中。',
    pathTips66: '数据集位置存储在运行参数 <strong style="color:#010101">{dataset}</strong> 中，预训练模型存放在运行参数 <strong style="color:#010101">{model}</strong> 中，训练输出请存储在 <strong style="color:#010101">{output}</strong> 中以供后续下载。',
    pathTips7: '数据集位置存储在运行参数 <strong style="color:#010101">{dataset}</strong> 中，模型存放在运行参数 <strong style="color:#010101">{model}</strong> 中，推理输出路径存储在运行参数 <strong style="color:#010101">{output}</strong> 中。',
    pathTips71: '推理脚本存储在 <strong style="color:#010101">{code}</strong> 中，数据集存储在 <strong style="color:#010101">{dataset}</strong> 中，模型存储在 <strong style="color:#010101">{model}</strong> 中，推理输出请存储在 <strong style="color:#010101">{output}</strong> 中以供后续下载。',
    basicInfo: '基本信息',
    paramsSetting: '参数设置',
    resourceSetting: '资源设置',
    runningStatus: '运行状态',
    basicParameters: '基本参数',
    otherParameters: '其他参数',
    parameters: '参数',
    valueAndContent: '数值&内容',
    taskPrepareTips: '任务正在准备中，喝杯水回来再看看~',
    waitCountStart: '您当前排队位置是第',
    waitCountEnd: '位',
    mindTorchHelper: '不想排队？查看如何一键将PyTorch迁移至MindSpore',
    task: '任务',
    taskList: '任务列表',
    taskName: '任务名称',
    appName: '应用名称',
    appList: '应用列表',
    taskNameTips: '只能以小写字母或数字开头且只包含小写字母、数字、_和-，不能以_结尾，最长36个字符。',
    taskNameTips1: '只能以小写字母开头，包含数字、小写字母和短横线(-)，长度为5～26个字符。',
    imageInnerUrlErrTips: '您指定的镜像地址不适用于该任务，请重新指定其它镜像。',
    imageUrlErrTips: '请输入合法的镜像地址。',
    taskDescr: '任务描述',
    taskDescrPlaceholder: '描述字数不超过255个字符',
    codeBranch: '代码分支',
    image: '镜像',
    selectImage: '选择镜像',
    selectImagePlaceholder: '选择镜像或输入镜像地址',
    dataset: '数据集',
    networkType: '访问Internet',
    allNetworkType: '全部网络',
    codeGenerate: '代码生成',
    noInternet: '否',
    hasInternet: '是',
    networkTypeDesc: '「访问Internet」是描述您的任务分配到的计算资源中心是否支持互联网访问的情况。选择“否”时可分配的计算资源中心更多，且依然可以使用启智平台上的数据集和模型。',
    resourceSpec: '资源规格',
    specPlaceholder: '请选择资源规格',
    specDescr: '资源说明',
    PointGainDescr: '积分获取说明',
    balanceOfPoints: '积分余额',
    points: '积分',
    canUseTime: '，预计可用',
    hours: '小时',
    versionCount: '版本数',
    runVersion: '运行版本',
    createTime: '创建时间',
    startRunTime: '开始运行时间',
    endRunTime: '结束运行时间',
    runDuration: '运行时长',
    refresh: '刷新',
    codePath: '代码存放路径',
    datasetPath: '数据集存放路径',
    modelPath: '模型存放路径',
    outputPath: '输出存放路径',
    codeObsPath: '代码obs地址',
    clusterAndComputeResource: '集群/计算资源',
    runParameter: '运行参数',
    addRunParameter: '新增运行参数',
    parameterName: '参数名',
    parameterValue: '参数值',
    customPath: '自定义路径',
    port: '端口',
    customPathTips: '自定义路径只允许输入小写字母和数字，首字符必须是字母，长度不超过32个字符; 端口号可选择范围为8000-8800',
    debug: '调试',
    reDebug: '再次调试',
    reInfer: '再次推理',
    stop: '停止',
    modify: '修改',
    "delete": '删除',
    startUse: '开始使用',
    more: '更多',
    commitImage: '提交镜像',
    downloadModel: '模型下载',
    configurationInfo: '配置信息',
    taskRuntimeInfo: '任务运行简况',
    log: '日志',
    logFile: '日志文件',
    downloadLog: '下载日志',
    viewFullScreen: '全屏',
    exitFullScreen: '退出全屏',
    scrolledToTopTip: '您已翻阅到日志顶部',
    scrolledToBottomTip: '您已翻阅到日志底部',
    resourceOccupancy: '资源占用情况',
    modelDownload: '结果下载',
    lossPlot: 'loss',
    evalOverview: '评测结果概览',
    evalDetail: '评测结果详情',
    failedReason: '运行失败原因',
    publicImage: '公开镜像',
    recommendImage: '平台推荐镜像',
    myImage: '我的镜像',
    myFavImage: '我收藏的镜像',
    searchImagePlaceholder: '搜镜像Tag/描述/操作系统/安装的软件包/标签...',
    useImage: '使用',
    submitting: '提交中',
    submitFailed: '提交失败',
    checkImageSizeTips: '检测提交镜像是否大小超过20G！',
    maxTaskTips: '<p><span>*</span> 平台仅留存近 <span>30</span> 天的调试、训练、推理、评测任务结果；<span>超过 30 天的任务将不能下载结果和查看日志，且不能再次调试、训练和微调。</span></p>',
    datasetFiles: '数据集',
    fileWasDeleted: '文件已经被删除',
    debugTaskEmptyTitle: '未创建过调试任务',
    debugTaskEmptyTip0: '代码版本：您还没有初始化代码仓库，请先 <a href="{url}">创建代码版本</a>；',
    debugTaskEmptyTip1: '运行时长：最长不超过4个小时，超过4个小时将自动停止；',
    debugTaskEmptyTip2: '数据集：云脑1提供 CPU / GPU 资源，云脑2提供 Ascend NPU 资源，调试使用的数据集也需要上传到对应的环境；',
    debugTaskEmptyTip3: '使用说明：可以参考启智AI协作平台 <a href="{url}" target="_blank">帮助中心</a>。',
    onlineInferTaskEmptyTitle: '未创建过在线推理任务',
    onlineInferEmptyTip2: '数据集：云脑1提供 CPU / GPU 资源，云脑2提供 Ascend NPU 资源，在线推理使用的数据集也需要上传到对应的环境；',
    superTaskEmptyTitle: '未创建过超算任务',
    generalTaskEmptyTitle: '未创建过通用任务',
    generalTaskEmptyTip1: '该任务类型具有如下特点，需在 <a href="/computingpower/demand">算力资源</a> 页面提交使用申请并审核通过后才能使用：<br/>（1）提供Jupyter调试环境。<br/>（2）运行时间不限时，但需消耗积分。<br/>（3）支持开放端口。<br/>（4）不支持保存镜像，不支持再次调试。',
    deleteConfirmTips: '您确认删除该任务么？此任务一旦删除不可恢复。',
    deleteBatchConfirmTips: '您确认批量删除选中的任务么？任务一旦删除不可恢复。',
    deletingTips: '任务删除中，请稍后',
    tabTitDebug: '调试任务',
    tabTitTrain: '训练任务',
    tabTitInference: '推理任务',
    tabTitBenchmark: '评测任务',
    trainTaskEmptyTitle: '未创建过训练任务',
    trainTaskEmptyTip2: '数据集：云脑1提供 CPU / GPU 资源，云脑2提供 Ascend NPU 资源，训练使用的数据集也需要上传到对应的环境；',
    inferenceTaskEmptyTitle: '未创建过推理任务',
    inferenceTaskEmptyTip2: '数据集：云脑1提供 CPU / GPU 资源，云脑2提供 Ascend NPU 资源，推理使用的数据集也需要上传到对应的环境；',
    bootFileTips: '启动文件是您程序执行的入口文件，必须是以.py结尾的文件。比如train.py、main.py、example/train.py、case/main.py。',
    viewSample: '查看样例',
    tabTitOnlineInference: '在线推理',
    tabTitGeneral: '通用任务',
    allResultDownload: '下载全部结果',
    downloadDisplayMaxCountTips: '单目录下最多显示 {count} 个文件或文件夹。计算任务容器内结果存储配额为 {size} GiB，超出配额将会导致结果丢失。',
    file_sync_ing: "文件同步中，请稍候",
    file_sync_wait: "文件等待同步中，请稍候",
    file_sync_fail: "文件同步失败",
    eval_task_ing: "评测任务进行中，稍后再来看看",
    no_file_to_download: "没有文件可以下载，稍后再来看看",
    task_not_finished: "任务还未结束，稍后再来看看",
    retrieve_results: "重新获取结果",
    reuseLastResult: '复用上次结果',
    continue_helper: '勾选复用将拷贝上次训练任务输出结果文件',
    computeNode: '计算节点',
    computeNodeCount: '计算节点数',
    exportDataset: {
      exportDatasetTitle: '导出结果至数据集',
      exportDatasetTips: '<span style="color:red">*</span> 导出的文件最终可以在您选择的数据集下查看。',
      export_failed: '导出失败',
      export_has_same_file: '当前数据集已存在相同的文件',
      export_has_same_file1: '当前模型已存在相同的文件',
      export_exceed_storage: '当前文件大小已经超出存储限额',
      export_success: '导出成功',
      exporting: '正在导出',
      please_select_file: '请先选择文件',
      please_select_output_file: '选择结果文件',
      select_file: '选择文件',
      file_descr: '文件描述',
      no_dataset: '还未创建过数据集',
      create_dataset: '去创建数据集',
      please_select_dataset: '请先选择要导出到的数据集'
    },
    chartResourceUsage: '占有率(%)',
    chartTime: '时间(min)',
    scrollToTop: '滚动到顶部',
    scrollToBottom: '滚动到底部',
    migratingData: '数据迁移中',
    centerPending: '分中心排队中',
    imagePulling: '拉取镜像中',
    sdkUseWay: '通过c2net库访问方式',
    dialogTips: {
      title1: '使用在线推理提供webui或api服务时需要了解以下几点',
      tips1: "平台没有对外直接提供自定义端口对外提供服务，可以使用fastapi转发到对外暴露的指定的URL提供服务。",
      tips1_1: 'webui服务参考：',
      tips1_2: 'api服务参考：',
      tips2: '服务启动指令参考',
      tips3: '在线推理任务只有在任务终止状态下才有日志返回',
      tips61: '详细使用方法可查看',
      tips62: '使用样例代码仓',
      tips7: '不再提示',
      tips8: '关闭',
      tips9: 'OpenI在线推理部署须知'
    },
    codeUseDlgTriggerTxt: '代码中如何访问数据资源',
    codeUseDlgTitle: '如何在代码中通过c2net库方式获取模型、数据集和输出路径',
    sdkCodeTip1: '请使用c2net库方式在容器中访问相关资源，可参考<a target="_blank" href="https://openi.pcl.ac.cn/docs/index.html#/cloudbrain/codepath">使用帮助</a>。',
    sdkCodeTip2: '根据您选择的计算任务类型，指定的数据集、模型等，访问数据集和模型，回传结果的示例代码如下：',
    generalTaskSdkCodeTip0: '通用任务如何开放端口',
    generalTaskSdkCodeTip1: '运行以下命令可在环境中配置HTTP代理：',
    generalTaskSdkCodeTip2: '/root/bin/scc tunnel http http://127.0.0.1:7860 （端口号7860可自定义）',
    generalTaskSdkCodeTip3: '配置成功后，输出如下信息：',
    generalTaskSdkCodeTip4: '即可使用地址<code>https://be22fe9f-d140-474c-81f4-30fb068157b5.tunnel.paracloud.com</code>访问相应服务。',
    searchTaskName: '搜索任务名称...',
    searchTaskNameOrCreator: '搜索任务名称/创建者...',
    downloadReport: '下载此报告',
    cloudbrainTaskType: '任务类型',
    ComputingResourceInfo: '算力资源信息',
    repo: '项目',
    cloudbrainTaskName: '云脑侧任务',
    taskIsAutomaticStop: '任务是否自动停止',
    automaticStop: '自动停止',
    manualStop: '手动停止',
    automaticStopTips: '该任务将在运行时长超过您所选择的时长后，或积分余额不足时，自动停止。',
    manualStopTips: '该任务将由您手动停止或积分余额不足时自动停止。',
    customize: '自定义',
    numOfHours: '{num} 小时',
    customizeTimeLimitPlaceholder: '请输入1到24之间的整数',
    debugTaskTimeLimitTips: '调试任务有使用时长限制，超时将自动结束任务，<span>请在到时前自行保存调试结果。</span>',
    autoStopTimeTips: '{min} 分钟后，该调试任务将自动结束，请自行保存调试结果',
    forkRepo: 'Fork已有项目',
    newRepo: '新建项目',
    migrateRepo: '迁移外部项目',
    pullRequests: '合并请求',
    issues: '任务管理',
    selectRepo: '指定已有项目',
    nextStep: '下一步',
    previousStep: '上一步',
    instructionsForUse: '使用须知',
    setUp: '设置',
    aiTaskType: 'AI任务类型',
    newCloudbrainAiTask: '新建云脑（AI）任务',
    selectAiTaskTypeAndRepo: '选择AI任务类型和项目',
    specifyComputingResourcesAndRarameters: '指定计算资源和参数',
    dataPreparing: '数据准备中，请稍后~',
    selectTheTaskType: '选择你需要创建的任务类型。',
    debugTaskDesc: '交互编辑环境',
    debugTaskDescLong: '提供Jupyter或Modelarts交互式调试环境，适合早期模型的调试和构建。',
    debugTaskDescLong1: '调试任务运行超过4个小时将自动停止。',
    trainTaskDesc: '提交至任务队列',
    trainTaskDescLong: '适用于长时间执行，不需要频繁修改算法代码的场景。',
    inferenceTaskDesc: '超算算力计算任务',
    inferenceTaskDescLong: '支持在超算算力上运行多种应用。',
    onlineinferTaskDesc: '搭建实时推理界面',
    onlineinferTaskDescLong: '为模型提供在线部署功能，使用Gradio框架，生成实时推理界面。',
    generalTaskDesc: '提供Jupyter调试环境',
    generalTaskDescLong: '提供Jupyter调试环境，需在算力资源页面提交使用申请并审核通过后即可使用。',
    noAiTasksInThisRepo: '该项目下还没有计算任务',
    createNewAitask: '新建计算任务',
    visualization: '可视化',
    sourceFtName: '源微调任务名称',
    tensorBoardVisualization: 'TensorBoard可视化',
    viewVisualization: '查看TensorBoard可视化结果',
    startDebug: '开始调试',
    stopTask: '停止任务',
    taskTmpl: '计算任务模板',
    saveTaskTmpl: '保存为计算任务模板',
    saveNewModel: '导出为新模型',
    exportData: '导出数据',
    exportToDataset: '导出结果至数据集',
    deploymentExperience: '部署体验',
    aimTrainCompare: '训练对比',
    aimTrainCompareDlgTitle: '选择对比的训练任务',
    aimTrainCompareSelectTips: '请选择 {minCount} 至 {maxCount} 个任务进行训练对比！',
    aimVisualization: 'AIM可视化',
    batchDelete: '批量删除',
    deleteFailed: '删除失败',
    all: '所有',
    ihave: '我拥有的',
    iCollaborate: '我协作的',
    iCollect: '我收藏的',
    selectApp: '选择应用',
    cbOffline: '项目下的“云脑”页签即将下线',
    cbOffflineTips: '请到<a href="/cloudbrains/create">“我的工作台 > 计算任务”</a>页面查看该页面原有的计算任务。',
    exportOutputTis: '导出结果至数据集或者模型，请移步pc网页端',
    repoStorageTips: '计算任务容器内项目存储配额为 {size} GiB，请控制项目大小。',
    runningLimit: '您的账户可以同时创建{count}个{taskType}。',
    forbidPenetrationStatement: '禁止任何形式的穿透行为声明',
    penaltyTipContent: '本平台严禁任何形式的<span class="highlight">穿透行为</span>。违规者将被扣除算力积分、限制登录，严重者将承担<span class="highlight">法律责任</span>！'
  },
  taskTmplObj: {
    tmplName: '模板名称',
    tmplNamePlacehoulder: '请输入模板名称，最多20个字符',
    tmplDescr: '模板描述',
    quickRunTmplTips: '一键运行计算任务模板',
    editTmplFile: '编辑模板文件',
    addTmpl: '新增模板',
    dragSortTips: '可拖动页签调整顺序',
    deleteTmplTips: '是否要删除当前计算任务模板？',
    completeTmplTips: '请完善第 {n} 个模板中的必填内容！',
    more: '更多',
    collapsed: '收起',
    edit: '编辑',
    run: '运行',
    runTmplForkTips: '您对本项目没有写入权限，无法新建计算任务，建议您fork本项目，然后再一键运行计算任务。',
    forkRepo: 'Fork本项目',
    runTmplFailedTips: '运行模板失败！',
    runTmplWithConfigsErrorTips: '模板配置异常，无法正常运行！',
    useTaskTmpl: '使用任务模板',
    saveTaskTmpl: '保存计算任务模板',
    unAvailableTmplSpec: '模板资源规格"{msg}"因权限或资源不存在等原因，<span style="color:red">目前不可用</span>',
    unAvailableTmplImage: '模板镜像"{msg}"因权限或资源不存在等原因，<span style="color:red">目前不可用</span>',
    unAvailableTmplBranch: '模板代码分支"{msg}"因权限或资源不存在等原因，<span style="color:red">目前不可用</span>',
    taskTmpl: '计算任务模板',
    publicTmpl: '公开模板',
    recommendTmpl: '平台推荐模板',
    runTimes: '运行次数',
    searchTaskTmpl: '搜索计算任务模板名称',
    createTaskTmpl: '新建计算任务模板',
    myCreated: '我创建的',
    myCollected: '我收藏的',
    deleteTaskTmplConfirmTips: '是否确认删除当前计算任务模板？',
    tmplAccessRight: '模板权限',
    accessRight: '权限',
    collectedNum: '收藏次数',
    tagsAndDescr: '标签及描述',
    modelDatasetAndRepo: '模型、数据集及代码',
    cancelRecommend: '取消推荐',
    setRecommend: '设为推荐',
    cancelRecommendSuccess: '取消推荐成功',
    setRecommendSuccess: '设为推荐成功',
    tmplCoreElements: '核心要素',
    editTmpl: '编辑模板',
    tmplInfo: '模板信息',
    tmplTags: '模板标签',
    tagsPlaceholder: '输入完成后回车键完成标签确定，最多输入10个',
    editTaskTmpl: '编辑计算任务模板',
    createNewTaskTmpl: '新建计算任务模板',
    createTaskTmplErrTips: '请完善计算任务模板信息',
    taskTmplSaveTips: '计算任务模板保存中，请稍后',
    taskTmplSaveFailedTips: '保存计算任务模板失败',
    taskTmplReferencedRepo: '引用过该代码仓的计算任务模板'
  },
  superComputeObj: {
    mmlSparkDescr: "MMLSpark\u5168\u79F0\u4E3AMicrosoft Machine Learning for Apache Spark\uFF0C\u652F\u6301\u7528\u6237\u8FD0\u884C\u81EA\u5236\u5BB9\u5668\u955C\u50CF\uFF0C\u4E14\u8D4B\u4E88\u4E86\u7528\u6237\u5BB9\u5668\u5185root\u6743\u9650\u3002\u7528\u6237\u53EF\u76F4\u63A5\u4F7F\u7528\u5E73\u53F0\u63D0\u4F9B\u7684\u5FAE\u8F6F\u7684MMLSpark\u3002\n\u6CE8\uFF1AMMLSpark\u662F\u5FAE\u8F6F\u63D0\u4F9B\u9488\u5BF9\u673A\u5668\u5B66\u4E60\u73AF\u5883\u7684Spark\u7248\u672C(<a target=\"_blank\" href=\"https://github.com/Azure/mmlspark\">https://github.com/Azure/mmlspark</a>)\uFF0C\u5173\u4E8Emmlspark\u53EF\u53C2\u8003\u8BBA\u6587\uFF1A<a target=\"_blank\" href=\"https://arxiv.org/pdf/1810.08744.pdf\">https://arxiv.org/pdf/1810.08744.pdf</a>"
  },
  modelSquare: {
    llmHeader: '模型体验',
    chatGlm_intro: '是一个开源的、支持中英双语的对话语言模型，由智谱AI提供。',
    llama2: 'is a collection of pretrained and fine-tuned generative text models ranging in scale from 7 billion to 70 billion parameters. This is the repository for the 7B fine-tuned model, optimized for dialogue use cases and converted for the Hugging Face Transformers format.',
    dialogtips1: '你好👋！欢迎体验大模型知识库问答',
    dialogtips21: '此体验基于',
    dialogtips22: '语言模型与m3-base向量模型',
    dialogtips3: '请在右侧选择直接与模型对话或基于本地知识库问答',
    dialogtips4: '知识库问答模式，选择知识库名称后，即可开始问答。',
    dialogtips5: '如有需要可以在选择知识库名称后上传文件/文件夹至知识库，或从知识库中删除文件。',
    dialogtips6: '模型已成功加载，可以开始对话，或从右侧选择模式后开始对话',
    promptPlaceholder: '请输入提问内容...(Ctrl + Enter = 换行, 按回车键进行提交)',
    dialogModeSelect: '请选择使用模式',
    dialogLLM: 'LLM 对话',
    dialogKb: '知识库问答',
    configKb: '配置知识库',
    updatekb: '更新已有知识库选项',
    recreateKb: '向量库重构中，请耐心等待，请勿刷新或关闭网页',
    selectKb: '请选择要加载的知识库:',
    createKb: '新建知识库',
    deleteKb: '删除本知识库',
    deleteKbTips: '确定删除{knowledgeValue}知识库吗？',
    deleteVbTips: '确定删除已选择的知识库文件吗？',
    uploadFile: '上传文件',
    uploadFileTips1: '将文件拖到此处，或<em>点击上传</em>',
    uploadFileTips2: '单个文件上传大小限制为1MB • HTML, MD, JSON, CSV, TXT, XML, DOCX',
    addFileToKb: '添加文件到知识库',
    manageFile: '管理文件',
    deleteKbFileSelect: '请从知识库已有文件中选择要删除的文件',
    deleteKbFile: '从知识库中删除文件',
    recreateKbSuccess: '{knowledgeValue}向量库重构成功',
    noPermission: '您没有权限操作',
    fileExit: '文件已经存在',
    fileExceed: '文件超过1MB',
    fileError: '文件错误,请重新上传!',
    fileTypeError: '文类型件错误,请重新上传!',
    fileUploadSuccess: '{fileName}文件上传成功!',
    kbName: '知识库名称',
    createKbPlaceholder: '新知识库名称只能是数字和字母',
    vectorType: '向量库类型',
    embedModel: 'Embedding 模型',
    cancel: '取 消',
    create: '新 建',
    ok: '确 定',
    kbNameDetect1: '名称不能为空',
    kbNameDetect2: '名称只能是数字和字母',
    chatExceedCount: '超出使用次数，无法体验模型服务！',
    useNotice: '使用体验须知',
    agreeNotice: '同意 <a href="/home/model_privacy" target="_blank"> 《OpenI启智社区AI协作平台免责声明和服务使用规范》 </a>中所述内容 <p style="text-align: center;margin-top: 1rem;color: red;">温馨提示：不合理使用可能会被封号!</p>',
    modelProvide: '《免责声明和服务使用规范》',
    modelNotExist: '模型不存在',
    maxTries: '限量体验{maxTries}次',
    modelChatTask: '创建模型在线体验任务',
    createChatTips1: '单击下方按钮创建在线体验任务，创建成功后可在线体验{expireMinutes}分钟',
    createChatTips2: '本人创建的在线体验任务只限本人体验',
    createChatTips3: '您已经创建过当前模型的在线体验任务，点击下方按钮直接进入体验界面',
    createChatBtns1: '创建在线体验任务',
    createEvalBtns1: '创建评测任务',
    createChatBtns2: '在线体验',
    experienceTime: '体验倒计时：',
    uploadFIleLimit: '最多上传10个文件',
    inputNotEmpty: '输入内容不能为空！',
    sessionChating: '会话加载中，不要心急哦！',
    sessionSding: '图像生成中，不要心急哦！',
    chatBanned: '您的帐户被禁止登录，请与网站管理员联系',
    chatIllegal: '非常抱歉，作为一个人工智能助手我只能提供客观的信息。您还有什么要问的吗？',
    chatExpireMins: '体验时间超过{expireMinutes}分钟，需重新创建[在线体验任务]({locaRefresh})',
    chatExpired: '对话已过期，请重新创建对话',
    stopExperience: '结束体验',
    experienceDuration: '已体验时长',
    requireTimeout: '请求超时，请稍后再试',
    sdPlaceholder: '请使用英文prompt描述想要生成的画面内容。',
    submitTips: 'Ctrl + Enter换行',
    negativePromptPlaceholder: '请描述不想要生成的画面内容。',
    width: '图宽',
    height: '图高',
    num_images_per_prompt: '图像数量',
    negative_prompt: '负向提示词',
    steps: '采样步数',
    seed: '随机种子',
    scheduler_name: '采样方式',
    guidance_scale: '提示词相关性',
    newChat: '全新对话',
    stopChat: '停止生成',
    stopedChat: '已停止生成',
    refreshChat: '重新生成',
    like: '喜欢',
    unlike: '不喜欢',
    chatHeaderTips: '你好👋！欢迎体验 {modelName} 模型',
    deepseekHeaderTips: '你好👋！欢迎体验 DeepSeek 多尺寸模型综合服务',
    //deepseekHeaderTips: '您好👋！欢迎体验DeepSeek智能问答服务！我们采用智能调度技术，可根据问题复杂程度自动匹配最佳参数规模的模型，为您提供精准高效的解答',
    systemPlaceholder: '请输入系统人设，例如“你是一个AI助手”',
    temperature: '温度',
    top_p: '多样性',
    repetition_penalty: '重复惩罚',
    max_tokens: '最大生成长度',
    system_message: '系统人设',
    history: '多轮对话',
    experienceStopTips: '任务停止中,请稍后!',
    experienceStopSuccess: '任务停止成功,该体验服务已经暂停!',
    modelFinetuen: '模型微调',
    pcMind: '鹏城 · 脑海',
    sdModelFinetuen: 'CV微调',
    sdModelFinetuenTask: 'CV微调任务',
    modelExperience: '模型体验',
    modelExperienceTips: '提供大模型在线体验功能，您可以根据自己的任务场景选择合适的模型和计算资源，创建模型在线体验任务，从而在线检验模型的反应效果。',
    modelEvaluateTips: '以opencompass提供的数据集为测试标准，对用户微调产生模型的性能（performance）进行评估。',
    largeModelsList: '大模型列表',
    newModelExperience: '新建模型体验任务',
    ModelExperienceDetail: '模型体验任务详情',
    ModelEvaluateDetail: '模型评测任务详情',
    modelEvaluate: '模型评测',
    modelPerEvaluate: '模型性能评测',
    modelSafeEvaluate: '模型安全评测',
    modelEvaluateTask: '模型性能评测任务',
    evaluationModel: '评测模型',
    sftFinetune: 'NLP微调',
    sftFinetuneTask: 'NLP微调任务',
    newSftFinetune: '创建NLP微调任务',
    newLoraFinetune: '创建CV微调任务',
    newComfyUi: '创建Comfy UI任务',
    sftFinetuneDetail: 'NLP微调任务详情',
    sftFinetuneLimit: '每个账户可以同时创建3个NLP微调任务',
    ModelExperienceLimit: '每个账户可以同时创建3个模型体验任务',
    ModelEvaluateLimit: '每个账户可以同时创建1个模型评测任务',
    sdLoraLimit: '每个账户可以同时创建3个CV微调任务',
    comfyuiLimit: '每个账户可以同时创建1个Comfy UI任务',
    sftFinetuneTips: '目前提供LoRA（Low-Rank Adaptation）训练模式，您可以根据自己的任务场景选择合适的训练数据、调整训练参数，从而实现理想的模型效果。',
    cvloraTips: 'LoRA是一种轻量化的模型调校方法，通过训练特定网络层的权重并插入到基础模型中，实现快速、高效且只需少量图片的训练，最终优化了模型的参数量和推理性能。',
    ComfyUiTips: 'Comfy UI 是一个基于节点的稳定扩散图形用户界面，主要用于构建高度灵活的图像生成工作流。它适用于数字艺术家、AI研究人员、内容创作者等多种用户群体',
    ComfyUiWarn: '服务正在启动中，请稍后再试！',
    uploadSound: '请上传声音模板',
    language: '语言: ',
    soundStyle: '声音风格: ',
    systemDefault: '系统默认',
    uploadSound1: '上传声音模板',
    dragDrop: '拖放音频至此处',
    outputAduio: '输出音频',
    appDev: '应用开发',
    relatedTools: '相关工具系统',
    inputLength: '输入字符超过限制',
    ttsPlaceholder: '请输入要转换成语音的文案，中、英文语音合成分别不超过82和250个字符...(Ctrl + Enter = 换行, 按回车键进行提交)',
    soundTemplateTips: '声音模板只支持mp3和wav格式文件',
    soundTemplateTips1: '声音模板仅支持单个文件',
    soundTemplateTips3: '上传声音模板目前只支持mp3,wav两种格式的音频文件，并且只支持单个文件',
    cvLoraTrain: 'CV微调',
    cvComfyui: 'Comfy UI',
    cvComfyuiTask: 'Comfy UI任务',
    modelArena: '大模型竞技场',
    send: '发送',
    newRound: '新开一轮',
    pleaseSelectModelAndConversation: '请选择模型后进行对话',
    model_name1: '模型名称',
    model_name: '模型英文名称',
    model_zh_name: '模型中文名称'
  },
  modelFinetune: {
    foldParameters: '折叠参数',
    expandParameters: '展开参数',
    epochs: '需要执行的训练总轮数',
    learningRate: 'AdamW 优化器的初始学习率',
    batchSize: '每张卡处理的样本数量',
    valSize: '验证集占全部样本的百分比',
    evalSteps: '每隔多少步进行验证',
    saveSteps: '每两次断点保存间的更新步数',
    maxsamples: '每个数据集使用的最大样本数',
    computeType: '是否使用混合精度训练。',
    gradientAccumulation: '梯度累计的步数',
    lRScheduler: '采用的学习率调节器名称',
    maximumGradientNorm: '用于梯度裁剪的范数',
    cutoffLen: '输入序列分词后的最大长度',
    preprocessingNumWorkers: '用于处理的进程数',
    loggingSteps: '每两次日志输出间的更新参数',
    warmupSteps: '学习率预热采用的步数',
    packing: '在有监督微调阶段将序列打包为相同长度的样本',
    optim: '使用的优化器：adamw_torch、adamw_8bit 或 adafactor',
    loraRank: 'LoRA 矩阵的秩大小',
    loraAlpha: 'LoRA 缩放系数大小',
    loraDropout: 'LoRA 权重随机丢弃的概率',
    loraTarget: '应用 LoRA 的模块名称。使用英文逗号分隔多个名称',
    additionalTarget: '除 LoRA 层以外的可训练模块名称。使用英文逗号分隔多个名称',
    maxTaskTips: '<p><span>*</span> 平台仅留存近 <span>30</span> 天的微调任务结果；<span>超过 30 天的任务将不能下载结果和查看日志，且不能再次微调。</span></p>',
    maxEvalTaskTips: '<p><span>*</span> 平台仅留存近 <span>30</span> 天的评测任务结果；<span>超过 30 天的任务将不能查看评测结果和日志。</span></p>',
    max_epoch: '迭代轮次',
    max_epoch_tips: '迭代轮次，控制训练过程中的迭代轮数。',
    io_strategy: '模型策略',
    io_strategy_tips: '--',
    keep_interval_update: '模型保存间隔',
    keep_interval_update_tips: '--',
    save_interval_update: '保存的迭代步数',
    save_interval_update_tips: '每隔多少个step，保存',
    lr: '学习率',
    lr_tips: '--',
    vitual_pipeline_model_parallel_size: '虚拟流水线并行度',
    vitual_pipeline_model_parallel_size_tips: '--',
    pipeline_model_parallel_size: '流水线并行度',
    pipeline_model_parallel_size_tips: '流水线并行，把模型按照神经元的层次进行拆分，不同层放到不同的 GPU 上去计算',
    tensor_model_parallel_size: '张量并行度',
    tensor_model_parallel_size_tips: '张量并行，也就是联合多个 GPU 同时做一个张量计算，比如说矩阵乘法',
    num_micro_batch: '微批处理数量',
    num_micro_batch_tips: '"微批处理数量取值范围是[1,32]，表示在每次训练迭代中累加的微批处理数。微批处理数量设置较大可以加速训练，但可能会导致内存问题。',
    micro_batch_size: '微批处理大小',
    micro_batch_size_tips: '一般只能是1, 2的话就会超显存，使用梯度累加代替',
    recomputer_granularity: '重计算策略',
    recomputer_granularity_tips: '选择开启重计算可压缩显存使用量，但会造成吞吐量下降。在显存够用的情况下建议关闭',
    modelTraining: '模型训练',
    modelTesting: '模型测试',
    trainStart: '立即训练',
    imageLabeling: '图片打标',
    totalPic: '共',
    numPic: '张',
    clearAllPic: '清空所有图片',
    addPic: '添加图片',
    captionMethod: '打标算法',
    captionthres: '打标阈值',
    modelTrigger: '模型触发词',
    modelTriggerph: '请输入触发词',
    captioning: '打标',
    imgUploading: '图片上传中...',
    sureClearAllPic: '是否要删除所有图片?',
    sureClearOnePic: '是否要删除本张图片?',
    clearAllPicSucc: '所有图片删除成功',
    clearAllPicFail: '所有图片删除失败',
    clearOnePicSucc: '图片删除成功',
    clearOnePicFail: '图片删除失败',
    autoCaption: '图片自动打标中...',
    trainProgress: '训练进度',
    currentSteps: '当前步数/总步数',
    trainingRounds: '训练轮数/总轮数',
    trainParameters: '训练参数',
    viewLoss: '查看loss',
    selected: '已选择',
    startGenerateImg: '开始生图',
    sampleImage: '实时样图',
    sampleImageGenerate: '样图生成中',
    trainProcessing: '数据处理和模型加载中...',
    trainingProgress: '训练中',
    trainingCompleted: '训练完成',
    loraModelName: 'LoRA模型',
    picWidth: '图片宽度（Width）',
    picHeight: '图片高度（Height）',
    loraRandomSeed: '随机种子 （Random Seed）',
    loraSamplingSteps: '采样步数 （Sampling Steps）',
    loraCFGScale: '提示词引导系数 （CFG Scale）',
    loraNegativePrompt: '负向提示词 （Negative Prompt）',
    loraInputPrompt: '请输入prompt!',
    capationAddPic: '为所有图片添加标签',
    capationAddBegin: '添加到行首',
    capationAddEnd: '添加到行尾',
    capationAddMsg: '标签输入不能为空！',
    ProfessionalParameters: '专业参数',
    useBasicModel: '使用底模',
    Repeat: '单张次数',
    Epoch: '循环轮次',
    totalSteps: '总步数',
    uploadImgCalc: '上传图片后计算',
    Prompt: '提示词',
    ProfessionalSetting: '专业设置',
    loraPromptPlace: '训练过程中会根据提示词实时生成样图',
    uploadImg: '上传图片',
    uploadImgAccept: '文件格式不符合要求，请上传.jpg, .jpeg, .png格式的文件',
    dragUploadImg: '拖拽图片到这里上传',
    uploadMaxImg: '最多添加 200 张图片，支持 PNG / JPG / JPEG ',
    uploadHavedImg: '上传已有图片集',
    uploadImgTips: '请保证图片和打标文件一一对应，zip包内不能包含多级目录，zip包名称和文件名称总字符数低于255',
    uploadImgTips1: '未检测到文件',
    uploadImgTips2: '当前最大允许上传200个图片文件',
    uploadImgTips3: '文件格式不符合要求，请上传 PNG / JPG / JPEG 格式的图片',
    cvTaskfaildTips: '任务已停止或失败请重新创建任务！',
    evalDatasetLimit: '评估数据集条数',
    sftFinetuneName: '微调任务名称',
    selectBaseModel: '选择基础模型',
    selectFtModel: '选择NLP微调模型',
    evalTips: '模型能力评测以 opencompass 提供的数据集为测试标准，仅供参考。',
    evalTaskCategory: '任务类别',
    evalScore: '得分',
    evalModelInput: '模型输入',
    evalModelOutput: '模型输出结果',
    standerAnswer: '标准答案',
    evalModelOutAnswer: '模型输出答案',
    allResult: '全部结果',
    correctResult: '正确结果',
    wrongResult: '错误结果',
    allDataset: '全部数据集'
  },
  userRole: {
    userRoleManagement: '用户角色管理',
    newOpRole: '新建操作角色',
    editOpRole: '编辑操作角色',
    viewOpRole: '查看操作角色',
    newReRole: '新建算力资源角色',
    editReRole: '编辑算力资源角色',
    viewReRole: '查看算力资源角色',
    newStorageRole: '新建存储资源角色',
    editStorageRole: '编辑存储资源角色',
    viewStorageRole: '查看存储资源角色',
    newContainerStorageRole: '新建容器内储存配额角色',
    editContainerStorageRole: '编辑容器内储存配额角色',
    viewContainerStorageRole: '查看容器内储存配额角色',
    roleName: '角色名称',
    roleType: '角色类型',
    roleDescription: '角色描述',
    isItDefault: '是否默认',
    viewDetail: '查看详情',
    "default": '默认',
    notDefault: '非默认',
    reCategory: '算力资源类',
    opCategory: '操作类',
    storageCategory: '存储资源类',
    allRoleType: '全部角色',
    deleteRoleConfirm: '删除【默认】角色会影响所有用户，',
    deleteRoleConfirm1: '是否确认删除当前 {name} 角色？',
    roleSelected: '请指定角色拥有的权限（可多选）',
    roleSelectedTips: '全选/全不选',
    allUserDefault: '所有用户默认',
    UserPermissionConfig: '用户权限配置',
    OrgPermissionConfig: '组织权限配置',
    allOpCategory: '全部操作类角色',
    allReCategory: '全部算力资源类角色',
    allStCategory: '全部存储资源类角色',
    userName: '用户名',
    OrganizationName: '组织名',
    opCategoryRole: '操作类角色',
    reCategoryRole: '算力资源类角色',
    stCategoryRole: '存储资源类角色',
    codeSizeLimit: '运行的计算任务容器内code目录存储配额',
    outputSizeLimit: '运行的计算任务容器内output目录存储配额',
    pleaseEnterContent: '请输入内容',
    pleaseEnterRoleContent: '请输入角色名称搜索...',
    userId: '用户ID',
    organizationId: '组织ID',
    editUserPermissions: '编辑用户权限',
    editOrgPermissions: '编辑组织权限',
    OperationalPermissions: '操作类权限',
    detail: ' 详情',
    permissionList: '权限清单',
    updatePermissions: '更新权限',
    resourcePermissions: '的算力资源权限清单',
    CResourcePermissions: '算力资源类权限',
    StoragePermissions: '存储资源类权限',
    permissionUpdateUnixDesc: '按权限更新时间降序',
    permissionUpdateUnixAsc: '按用户更新时间升序',
    permissionCreatedUnixDesc: '按用户创建时间降序',
    permissionCreatedUnixAsc: '按用户创建时间升序',
    orgPermissionCreatedUnixDesc: '按组织创建时间降序',
    orgPermissionCreatedUnixAsc: '按组织创建时间升序',
    batchSetOrCancel: '批量设置或取消用户角色',
    enterUserId: '请输入需设置权限的用户名单',
    failed: '失败了',
    times: '个',
    failedUserList: '失败用户列表',
    addAcessSuccess: '增加权限成功',
    cancelAcessSuccess: '取消权限成功',
    operaNodeTips1: '运行训练任务时可以使用',
    operaNodeTips2: '节点的',
    operaNodeTips3: '计算资源',
    operaTaskTips1: '可以同时运行',
    operaTaskTips2: '个'
  },
  org: {
    organization: '组织',
    createOrg: '新建组织',
    searchOrg: '搜索组织',
    orgName: '组织名称',
    orgFullName: '组织全称',
    Members: '成员',
    Teams: '团队',
    orgTeam: '组织团队',
    orgMember: '组织成员',
    numMember: '名成员',
    numMembers: '名成员',
    numRepo: '个项目',
    numRepos: '个项目',
    unfold: '展开',
    fold: '收起',
    searchRepos: '搜索项目',
    searchModels: '搜索模型',
    searchDatasets: '搜索数据集',
    searchImages: '搜索镜像',
    customize: '自定义',
    selectedProjects: '精选项目',
    selectedModels: '精选模型',
    selectedDatasets: '精选数据集',
    customizeSelectedProjects: '自定义精选项目',
    customizeSelectedModels: '自定义精选模型',
    customizeSelectedDatasets: '自定义精选数据集',
    searching: '搜索...',
    remainNum: '还能推荐 {num} 个',
    maxProjects: '最多可选择 {num} 个公开项目',
    maxModels: '最多可选择 {num} 个公开模型',
    maxDatasets: '最多可选择 {num} 个公开数据集',
    order_newest: '最新创建',
    order_oldest: '最早创建',
    order_recentupdate: '最近更新',
    order_leastupdate: '最少更新',
    order_reversealphabetically: '按字母逆序排序',
    order_alphabetically: '按字母顺序排序',
    order_moststars: '点赞由多到少',
    order_feweststars: '点赞由少到多',
    order_mostforks: '派生由多到少',
    order_fewestforks: '派生由少到多',
    no_result: '空荡荡的，什么都没有',
    myOrg: '我的组织',
    orgDataset: '组织数据集',
    orgModel: '组织模型'
  },
  storage: {
    capacity_details: '存储容量明细',
    quota: '存储配额',
    data_usaged: '数据集已用',
    model_usaged: '模型已用',
    remaining_available: '剩余可用',
    upload_time: '上传时间',
    deleteModelConfirm: '是否确认删除当前模型 <span style="color:red;word-break: break-all;">{name}</span> ？',
    deleteModelConSuccess: '删除当前模型 {name} 成功！',
    deleteDataSetConfirm: '是否确认删除当前数据集 <span style="color:red;word-break: break-all;">{name}</span> ？',
    deleteDataSetSuccess: '删除当前数据集 {name} 成功 ！',
    feedback_issue: '反馈问题',
    remain_storage: '剩余存储配额',
    selected_file_size: '已选文件大小',
    exceedStorage: '已经超出存储限额',
    deleteBatchData: '您确认删除批量选中的数据集吗？一旦删除不可恢复',
    deleteBatchModel: '您确认删除批量选中的模型吗？一旦删除不可恢复',
    deletingTips: '文件删除中，请稍后',
    owenerTips: '请查看 {ownerName}（拥有者）的存储容量明细',
    storageOrgTips1: '申请更多配额，请单击',
    storageOrgTips2: '下载模板',
    storageOrgTips3: '填写好后发送至邮箱 secretariat@openi.org.cn'
  },
  evalCaterary: {
    Examination: "学科",
    Reasoning: "推理",
    Knowledge: "知识",
    Code: "代码",
    Math: "数学",
    Understanding: "理解",
    Language: "语言",
    Instruct: "指令跟随"
  },
  thirdPartyApp: {
    thirdPartyApp: '第三方应用',
    oauth2App: 'OAuth2应用',
    oauth2AppDescr: '（已认证过的OAuth2应用，用户授权后，该应用可获取用户的基本信息（用户id、用户名、用户头像链接）、电子邮箱、手机号码）',
    authStatus: '认证状态',
    canGetUserInfo: '可获得的用户信息',
    searchThirdPartyAppPlaceholder: '搜第三方应用名称',
    clientID: '客户端ID',
    appName: '应用名称',
    appCanGetUserInfo: '应用可获得的用户信息',
    appCreator: '应用创建者',
    authenticated: '已认证',
    notAuthenticated: '未认证',
    setAuthenticated: '通过认证',
    cancelAuthentication: '取消认证',
    basicInfo: '基本信息',
    basicInfoAll: '基本信息（用户id、用户名、用户头像链接）',
    email: '电子邮箱',
    phoneNumber: '手机号码',
    editOauth2AppAuthInfo: '修改OAuth2应用认证信息',
    authenticatedCanGetInfoTips: '用户授权后，该应用可获得如下用户信息',
    verify: '认证',
    cancelVerify: '取消认证',
    verifyAgain: '再次认证',
    cancelVerifyTips: '取消认证后，该应用将只能获得用户基本信息，不能获得用户电子邮箱、手机号码，是否继续操作？'
  },
  computingPowerObj: {
    computeResource: '算力资源',
    computeResourceDescr: '为了满足用户的个性化算力需求，启智社区可以根据用户的需求定制算力池，欢迎提交您的算力需求。',
    freeComputingPower: '普惠算力',
    paidComputingPower: '付费算力',
    paidComputingPowerTip: '付费算力及售后服务由合作伙伴提供',
    createNewComputingTask: '新建计算任务',
    createComputingTask: '创建算力任务',
    displayFilteringConditions: '显示筛选条件',
    priceRange: '价格区间',
    point_hr: '积分/卡时',
    use: '去使用',
    buyNow: '立即购买',
    provider: '提供方',
    systemDisk: '系统盘',
    computingPowerPartner: '算力合作伙伴',
    computingPowerOperationPlatform: '算力运营平台',
    chipManufacturers: '芯片厂商',
    accCardCount: '卡数'
  },
  dashboard: {
    welcomeTips: '欢迎来到启智AI协作平台',
    newUserGuide: '新用户指引',
    quickAccess: '快捷入口',
    createTask: '创建计算任务',
    taskDesc: '使用算力卡，创建调试、训练、推理等计算任务',
    myTaskCount: '我的计算任务',
    runningTasks: '运行中',
    totalGPUHourd: '累计运行',
    createProject: '创建项目',
    createRepoDesc: '创建代码仓，在计算任务中使用这些代码',
    uploadDatasetDesc: '上传数据集，在计算任务中使用这些数据集',
    uploadModelDesc: '上传AI模型，在计算任务中使用这些模型',
    oneClickCreate: '一键创建',
    remainingQuota: '剩余可用存储配额',
    occupiedStorage: '占用存储',
    datasetStorage: '数据集占用',
    modelStorage: '模型占用',
    credits: '算力积分',
    currentAvailable: '当前可用算力积分',
    currentAvailable1: '当前可用',
    totalGained: '总获取算力积分',
    totalGained1: '总获取',
    totalConsumed: '总消耗算力积分',
    totalConsumed1: '总消耗',
    announcements: '公告',
    moreAnnouncements: '更多公告',
    activities: '活动',
    weChatLinked: '微信已绑定',
    weChatNotLinked: '微信未绑定',
    datasetNum: '数据集',
    modelNum: '模型',
    unitPoints: '分',
    unitTask: '个',
    unitdata: '个',
    runtimeCards: '卡时',
    templateNum: '模板总数量',
    accumulatedRuns: '累计运行',
    unitTemplate: '次',
    heatmapLoading: '正在加载热图...',
    contributionCount: ' 最近一年贡献：{count} 次'
  },
  agentPortal: {
    taskRunningTip: '你有 {count} 个智能体正在运行，消耗 {numPoint} 积分/每小时',
    viewRunningAgents: '查看正在运行的智能体',
    bannerTitle: '开启您的 AI 智能体之旅',
    bannerDesc: '基于中国算力网强大算力底座，将前沿大模型能力一键部署，即可随时随地体验Al魅力。',
    searchPlaceholder: '搜索智能体',
    agentSquare: '智能体广场',
    runningAgents: '运行中的智能体',
    viewDetails: '了解详情',
    oneClickDeploy: '一键部署',
    deployFailed: '运行失败',
    deploying: '部署中',
    stop: '停止',
    startUsing: '开始使用',
    startSuccess: '启动成功',
    startFailed: '启动失败',
    stopSuccess: '停止成功',
    stopFailed: '停止失败',
    agentsRunning: '智能体 {names} 正在运行中',
    updated: '更新于',
    noData: '暂无数据'
  }
};
/* harmony default export */ __webpack_exports__["default"] = (zh);

/***/ }),

/***/ "./web_src/vuepages/langs/index.js":
/*!*****************************************!*\
  !*** ./web_src/vuepages/langs/index.js ***!
  \*****************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   i18n: function() { return /* binding */ i18n; },
/* harmony export */   lang: function() { return /* binding */ lang; }
/* harmony export */ });
/* harmony import */ var vue__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! vue */ "./node_modules/vue/dist/vue.esm.js");
/* harmony import */ var vue_i18n__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! vue-i18n */ "./node_modules/vue-i18n/dist/vue-i18n.esm.js");
/* harmony import */ var _config_zh_CN__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./config/zh-CN */ "./web_src/vuepages/langs/config/zh-CN.js");
/* harmony import */ var _config_en_US__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./config/en-US */ "./web_src/vuepages/langs/config/en-US.js");




vue__WEBPACK_IMPORTED_MODULE_3__["default"].use(vue_i18n__WEBPACK_IMPORTED_MODULE_0__["default"]);
var lang = window.config.lang;
var i18n = new vue_i18n__WEBPACK_IMPORTED_MODULE_0__["default"]({
  locale: lang,
  messages: {
    'zh-CN': _config_zh_CN__WEBPACK_IMPORTED_MODULE_1__["default"],
    'en-US': _config_en_US__WEBPACK_IMPORTED_MODULE_2__["default"]
  },
  silentTranslationWarn: true
});

/***/ }),

/***/ "./web_src/vuepages/pages/dataset/square/constant.js":
/*!***********************************************************!*\
  !*** ./web_src/vuepages/pages/dataset/square/constant.js ***!
  \***********************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   Category: function() { return /* binding */ Category; },
/* harmony export */   License: function() { return /* binding */ License; },
/* harmony export */   Task: function() { return /* binding */ Task; }
/* harmony export */ });
var Category = [{
  name: "computer_vision",
  active: false
}, {
  name: "natural_language_processing",
  active: false
}, {
  name: "speech_processing",
  active: false
}, {
  name: "computer_vision_natural_language_processing",
  active: false
}, {
  name: "machine_translation",
  active: false
}, {
  name: "medical_imaging",
  active: false
}];
var Task = [{
  name: "machine_translation",
  active: false
}, {
  name: "question_answering_system",
  active: false
}, {
  name: "information_retrieval",
  active: false
}, {
  name: "knowledge_graph",
  active: false
}, {
  name: "text_annotation",
  active: false
}, {
  name: "text_categorization",
  active: false
}, {
  name: "language_modeling",
  active: false
}, {
  name: "speech_recognition",
  active: false
}, {
  name: "information_extraction",
  active: false
}, {
  name: "description_generation",
  active: false
}, {
  name: "image_classification",
  active: false
}, {
  name: "face_recognition",
  active: false
}, {
  name: "image_search",
  active: false
}, {
  name: "target_detection",
  active: false
}, {
  name: "image_description_generation",
  active: false
}, {
  name: "vehicle_license_plate_recognition",
  active: false
}, {
  name: "medical_image_analysis",
  active: false
}, {
  name: "unmanned",
  active: false
}, {
  name: "unmanned_security",
  active: false
}, {
  name: "drone",
  active: false
}, {
  name: "vr_ar",
  active: false
}, {
  name: "2_d_vision",
  active: false
}, {
  name: "2_5_d_vision",
  active: false
}, {
  name: "3_d_reconstruction",
  active: false
}, {
  name: "image_processing",
  active: false
}, {
  name: "video_processing",
  active: false
}, {
  name: "visual_input_system",
  active: false
}, {
  name: "speech_coding",
  active: false
}, {
  name: "speech_enhancement",
  active: false
}, {
  name: "speech_synthesis",
  active: false
}, {
  name: "ros_hmci_datasets",
  active: false
}];
var License = [{
  name: "MIT",
  active: false
}, {
  name: "GPL",
  active: false
}, {
  name: "GFDL",
  active: false
}, {
  name: "Apache 2.0",
  active: false
}, {
  name: "CC-BY",
  active: false
}, {
  name: "CC-BY-SA",
  active: false
}, {
  name: "CC-BY-NC",
  active: false
}, {
  name: "CC-BY-ND",
  active: false
}, {
  name: "CC0: Public Domain",
  active: false
}, {
  name: "CDLA Permissive 1.0",
  active: false
}, {
  name: "Open Database License",
  active: false
}, {
  name: "Research or commercial",
  active: false
}, {
  name: "Research and commercial",
  active: false
}, {
  name: "CC-BY-NC-ND",
  active: false
}, {
  name: "CC-BY-NC-SA",
  active: false
}, {
  name: "CC-BY 2.0",
  active: false
}, {
  name: "CC-BY-SA 2.0",
  active: false
}, {
  name: "GNU Free Documentation License",
  active: false
}, {
  name: "CC-BY-NC 2.0",
  active: false
}, {
  name: "CC-BY-ND 2.0",
  active: false
}, {
  name: "CC-BY-NC-ND 2.0",
  active: false
}, {
  name: "CC-BY-NC-SA 2.0",
  active: false
}, {
  name: "Non-commercial",
  active: false
}, {
  name: "CC-BY 3.0",
  active: false
}, {
  name: "CC-BY-SA 3.0",
  active: false
}, {
  name: "CC-BY-NC 3.0",
  active: false
}, {
  name: "CC-BY-ND 3.0",
  active: false
}, {
  name: "CC-BY-NC-ND 3.0",
  active: false
}, {
  name: "CC-BY-NC-SA 3.0",
  active: false
}, {
  name: "CC-BY 4.0",
  active: false
}, {
  name: "CC-BY-SA 4.0",
  active: false
}, {
  name: "CC-BY-NC 4.0",
  active: false
}, {
  name: "CC-BY-ND 4.0",
  active: false
}, {
  name: "CC-BY-NC-ND 4.0",
  active: false
}, {
  name: "CC-BY-NC-SA 4.0",
  active: false
}];

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/router/index.js":
/*!*********************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/router/index.js ***!
  \*********************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var vue__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! vue */ "./node_modules/vue/dist/vue.esm.js");
/* harmony import */ var vue_router__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! vue-router */ "./node_modules/vue-router/dist/vue-router.esm.js");
/* harmony import */ var _views_HomePage_vue__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../views/HomePage.vue */ "./web_src/vuepages/pages/ros-hmci/views/HomePage.vue");
/* harmony import */ var _views_OSSystem_vue__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../views/OSSystem.vue */ "./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue");
/* harmony import */ var _views_OpenApp_vue__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../views/OpenApp.vue */ "./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue");
/* harmony import */ var _views_OpenDataset_vue__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../views/OpenDataset.vue */ "./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue");
/* harmony import */ var _views_OpenModel_vue__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ../views/OpenModel.vue */ "./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue");
/* harmony import */ var _views_CommunitySource_vue__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ../views/CommunitySource.vue */ "./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue");
/* harmony import */ var _views_HelpCenter_vue__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ../views/HelpCenter.vue */ "./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue");
/* harmony import */ var _views_ResourceDetail_vue__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ../views/ResourceDetail.vue */ "./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue");










vue__WEBPACK_IMPORTED_MODULE_8__["default"].use(vue_router__WEBPACK_IMPORTED_MODULE_9__["default"]);
/* harmony default export */ __webpack_exports__["default"] = (new vue_router__WEBPACK_IMPORTED_MODULE_9__["default"]({
  base: '/ros-hmci',
  routes: [{
    path: '/',
    name: 'HomePage',
    component: _views_HomePage_vue__WEBPACK_IMPORTED_MODULE_0__["default"]
  }, {
    path: '/os-system',
    name: 'OSSystem',
    component: _views_OSSystem_vue__WEBPACK_IMPORTED_MODULE_1__["default"]
  }, {
    path: '/open-app',
    name: 'OpenApp',
    component: _views_OpenApp_vue__WEBPACK_IMPORTED_MODULE_2__["default"]
  }, {
    path: '/open-data',
    name: 'OpenDataset',
    component: _views_OpenDataset_vue__WEBPACK_IMPORTED_MODULE_3__["default"]
  }, {
    path: '/open-model',
    name: 'OpenModel',
    component: _views_OpenModel_vue__WEBPACK_IMPORTED_MODULE_4__["default"]
  }, {
    path: '/community-source',
    name: 'CommunitySource',
    component: _views_CommunitySource_vue__WEBPACK_IMPORTED_MODULE_5__["default"]
  }, {
    path: '/help-center',
    name: 'HelpCenter',
    component: _views_HelpCenter_vue__WEBPACK_IMPORTED_MODULE_6__["default"]
  }, {
    path: '/source-detail/:name',
    name: 'ResourceDetail',
    component: _views_ResourceDetail_vue__WEBPACK_IMPORTED_MODULE_7__["default"]
  }]
}));

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/links.js":
/*!********************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/links.js ***!
  \********************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony default export */ __webpack_exports__["default"] = ({
  // 主页的六个tab栏
  "home_tab1": "/root/343432",
  "home_tab2": "/root/343432",
  "home_tab3": "/root/343432",
  "home_tab4": "/root/343432",
  "home_tab5": "/root/343432",
  "home_tab6": "/root/343432",
  //主页的项目协同后两项获取地址
  "home_statisticsData": "/fanshuai/ROS-hmci-resource/raw/branch/master/homePage_statistics.json",

  /*
  json格式如下
  {
      "developer": "500+",
      "organization": "100+"
  }
   */
  //主页底部的三个项目介绍
  "home_app1": "/root/343432",
  "home_app2": "/root/343432",
  "home_app3": "/root/343432",
  // 主页意见反馈跳转地址
  "home_advice": "/fanshuai/ROS-hmci_ISSUE/issues/new",
  // 操作系统界面系统概述
  "os_info": "/root/343432",
  //操作系统界面六个tab栏
  "os_tab1": "/root/343432",
  "os_tab2": "/root/343432",
  "os_tab3": "/root/343432",
  "os_tab4": "/root/343432",
  "os_tab5": "/root/343432",
  "os_tab6": "/root/343432",
  // 资源发布跳转
  "source": "/fanshuai/ros-hmci-resource/issues/new",

  /*
  ------------------以下为API接口地址-----------------------
  */
  // 资源发布获取仓库地址
  "source_Api": "/fanshuai/ROS-hmci-resource/raw/branch/master/resource.json" // getRepoApiUrl(repoName) {
  //     return `/root/${repoName}/datasets/current_repo_m?q=&page=1&q=&type=0`;
  // },

});

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/vp-ros-hmci.js":
/*!********************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/vp-ros-hmci.js ***!
  \********************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var vue__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! vue */ "./node_modules/vue/dist/vue.esm.js");
/* harmony import */ var _App_vue__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./App.vue */ "./web_src/vuepages/pages/ros-hmci/App.vue");
/* harmony import */ var _router__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./router */ "./web_src/vuepages/pages/ros-hmci/router/index.js");
/* harmony import */ var element_ui_lib_locale_lang_en__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! element-ui/lib/locale/lang/en */ "./node_modules/element-ui/lib/locale/lang/en.js");
/* harmony import */ var element_ui_lib_locale_lang_zh_CN__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! element-ui/lib/locale/lang/zh-CN */ "./node_modules/element-ui/lib/locale/lang/zh-CN.js");
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");
/* harmony import */ var element_ui__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! element-ui */ "./node_modules/element-ui/lib/element-ui.common.js");
/* harmony import */ var element_ui__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(element_ui__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var element_ui_lib_theme_chalk_index_css__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! element-ui/lib/theme-chalk/index.css */ "./node_modules/element-ui/lib/theme-chalk/index.css");
/* harmony import */ var _style_index_css__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ./style/index.css */ "./web_src/vuepages/pages/ros-hmci/style/index.css");









vue__WEBPACK_IMPORTED_MODULE_8__["default"].use((element_ui__WEBPACK_IMPORTED_MODULE_5___default()), {
  locale: _langs__WEBPACK_IMPORTED_MODULE_4__.lang === 'zh-CN' ? element_ui_lib_locale_lang_zh_CN__WEBPACK_IMPORTED_MODULE_3__["default"] : element_ui_lib_locale_lang_en__WEBPACK_IMPORTED_MODULE_2__["default"]
}), vue__WEBPACK_IMPORTED_MODULE_8__["default"].config.productionTip = false;
new vue__WEBPACK_IMPORTED_MODULE_8__["default"]({
  i18n: _langs__WEBPACK_IMPORTED_MODULE_4__.i18n,
  router: _router__WEBPACK_IMPORTED_MODULE_1__["default"],
  render: function render(h) {
    return h(_App_vue__WEBPACK_IMPORTED_MODULE_0__["default"]);
  }
}).$mount('#__vue-root');

/***/ }),

/***/ "./web_src/vuepages/utils/index.js":
/*!*****************************************!*\
  !*** ./web_src/vuepages/utils/index.js ***!
  \*****************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   RandomGenerator: function() { return /* binding */ RandomGenerator; },
/* harmony export */   escapeHTML: function() { return /* binding */ escapeHTML; },
/* harmony export */   getListValueWithKey: function() { return /* binding */ getListValueWithKey; },
/* harmony export */   getUrlSearchParams: function() { return /* binding */ getUrlSearchParams; },
/* harmony export */   initClipboard: function() { return /* binding */ initClipboard; },
/* harmony export */   renderSpecObject: function() { return /* binding */ renderSpecObject; },
/* harmony export */   renderSpecStr: function() { return /* binding */ renderSpecStr; },
/* harmony export */   setWebpackPublicPath: function() { return /* binding */ setWebpackPublicPath; },
/* harmony export */   timeSinceUnix: function() { return /* binding */ timeSinceUnix; },
/* harmony export */   toBoolean: function() { return /* binding */ toBoolean; },
/* harmony export */   transFileSize: function() { return /* binding */ transFileSize; },
/* harmony export */   uuidv4: function() { return /* binding */ uuidv4; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.index-of */ "./node_modules/core-js/modules/es.array.index-of.js");
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.iterator */ "./node_modules/core-js/modules/es.array.iterator.js");
/* harmony import */ var core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.number.to-fixed */ "./node_modules/core-js/modules/es.number.to-fixed.js");
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_7__);
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
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_parse_float__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.parse-float */ "./node_modules/core-js/modules/es.parse-float.js");
/* harmony import */ var core_js_modules_es_parse_float__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_parse_float__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! core-js/modules/es.string.iterator */ "./node_modules/core-js/modules/es.string.iterator.js");
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_18___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_18__);
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! core-js/modules/es.string.search */ "./node_modules/core-js/modules/es.string.search.js");
/* harmony import */ var core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_19___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_search__WEBPACK_IMPORTED_MODULE_19__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_20___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_20__);
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! core-js/modules/web.dom-collections.iterator */ "./node_modules/core-js/modules/web.dom-collections.iterator.js");
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_21___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_21__);
/* harmony import */ var core_js_modules_web_url__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! core-js/modules/web.url */ "./node_modules/core-js/modules/web.url.js");
/* harmony import */ var core_js_modules_web_url__WEBPACK_IMPORTED_MODULE_22___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_url__WEBPACK_IMPORTED_MODULE_22__);
/* harmony import */ var _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! @babel/runtime/helpers/classCallCheck */ "./node_modules/@babel/runtime/helpers/classCallCheck.js");
/* harmony import */ var _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_23___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_23__);
/* harmony import */ var _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_24__ = __webpack_require__(/*! @babel/runtime/helpers/createClass */ "./node_modules/@babel/runtime/helpers/createClass.js");
/* harmony import */ var _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_24___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_24__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_25__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_25___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_25__);
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_26__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");
/* harmony import */ var _const__WEBPACK_IMPORTED_MODULE_27__ = __webpack_require__(/*! ~/const */ "./web_src/vuepages/const/index.js");
/* harmony import */ var clipboard__WEBPACK_IMPORTED_MODULE_28__ = __webpack_require__(/*! clipboard */ "./node_modules/clipboard/dist/clipboard.js");
/* harmony import */ var clipboard__WEBPACK_IMPORTED_MODULE_28___default = /*#__PURE__*/__webpack_require__.n(clipboard__WEBPACK_IMPORTED_MODULE_28__);



























function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_25___default()(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }




var getListValueWithKey = function getListValueWithKey(list, key) {
  var k = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : 'k';
  var v = arguments.length > 3 && arguments[3] !== undefined ? arguments[3] : 'v';

  for (var i = 0, iLen = list.length; i < iLen; i++) {
    var listI = list[i];
    if (listI[k] === key) return listI[v];
  }

  return key;
};
var getUrlSearchParams = function getUrlSearchParams() {
  var params = new URLSearchParams(location.search);
  var obj = {};
  params.forEach(function (value, key) {
    obj[key] = value;
  });
  return obj;
};
var uuidv4 = function uuidv4() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
    var r = Math.random() * 16 | 0,
        v = c === 'x' ? r : r & 0x3 | 0x8;
    return v.toString(16);
  });
};
var transFileSize = function transFileSize(srcSize) {
  if (null == srcSize || srcSize == '') {
    return '0 Bytes';
  }

  var unitArr = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PB', 'EB', 'ZB', 'YB'];
  srcSize = parseFloat(srcSize);
  var index = Math.floor(Math.log(srcSize) / Math.log(1024));
  var size;

  if (srcSize % Math.pow(1024, index) !== 0) {
    size = (srcSize / Math.pow(1024, index)).toFixed(2);
  } else {
    size = srcSize / Math.pow(1024, index);
  }

  return size + ' ' + unitArr[index];
};
var renderSpecStr = function renderSpecStr(spec, showPoint) {
  if (!spec) return '';
  var ngpu = "".concat(spec.ComputeResource, ": ").concat(spec.AccCardsNum + '*' + getListValueWithKey(_const__WEBPACK_IMPORTED_MODULE_27__.ACC_CARD_TYPE, spec.AccCardType));
  var gpuMemStr = spec.GPUMemGiB != 0 ? "(".concat(_langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('resourcesManagement.gpuMem'), ": ").concat(spec.GPUMemGiB, "GB)") : '';
  var memStr = spec.MemGiB != 0 ? ", ".concat(_langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('resourcesManagement.mem'), ": ").concat(spec.MemGiB, "GB") : '';
  var sharedMemStr = spec.ShareMemGiB != 0 ? ", ".concat(_langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('resourcesManagement.shareMem'), ": ").concat(spec.ShareMemGiB, "GB") : '';
  var pointStr = showPoint ? ", ".concat(spec.UnitPrice == 0 ? _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('resourcesManagement.free') : spec.UnitPrice.toFixed(2) + _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('resourcesManagement.point_hr')) : '';
  var specStr = "".concat(ngpu).concat(gpuMemStr, ", CPU: ").concat(spec.CpuCores).concat(memStr).concat(sharedMemStr).concat(pointStr);
  return specStr;
};
var renderSpecObject = function renderSpecObject(spec, showPoint) {
  if (!spec) {
    return {
      type: '',
      specStr: '',
      pointStr: ''
    };
  }

  var ngpu = "".concat(spec.compute_resource, ": ").concat(spec.acc_cards_num + '*' + getListValueWithKey(_const__WEBPACK_IMPORTED_MODULE_27__.ACC_CARD_TYPE, spec.acc_card_type));
  var gpuMemStr = spec.gpu_mem_gi_b != 0 ? "(".concat(_langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('resourcesManagement.gpuMem'), ": ").concat(spec.gpu_mem_gi_b, "GB)") : '';
  var memStr = spec.mem_gi_b != 0 ? ", ".concat(_langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('resourcesManagement.mem'), ": ").concat(spec.mem_gi_b, "GB") : '';
  var sharedMemStr = spec.share_mem_gi_b != 0 ? ", ".concat(_langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('resourcesManagement.shareMem'), ": ").concat(spec.share_mem_gi_b, "GB") : '';
  var pointStr = showPoint ? "".concat(spec.unit_price == 0 ? _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('resourcesManagement.free') : spec.unit_price.toFixed(2) + _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('resourcesManagement.point_hr')) : '';
  var specStr = "".concat(ngpu).concat(gpuMemStr, ", CPU: ").concat(spec.cpu_cores).concat(memStr).concat(sharedMemStr);
  return _objectSpread(_objectSpread({}, spec), {}, {
    id: spec.id.toString(),
    type: spec.compute_resource,
    specStr: specStr,
    pointStr: pointStr
  });
};
var Minute = 60;
var Hour = 60 * Minute;
var Day = 24 * Hour;
var Week = 7 * Day;
var Month = 30 * Day;
var Year = 12 * Month;

var computeTimeDiff = function computeTimeDiff(diff) {
  var diffStr = '';

  switch (true) {
    case diff <= 0:
      diff = 0;
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.now');
      break;

    case diff < 2:
      diff = 0;
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.1s');
      break;

    case diff < 1 * Minute:
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.seconds', {
        msg: Math.floor(diff)
      });
      diff = 0;
      break;

    case diff < 2 * Minute:
      diff -= 1 * Minute;
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.1m');
      break;

    case diff < 1 * Hour:
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.minutes', {
        msg: Math.floor(diff / Minute)
      });
      diff -= diff / Minute * Minute;
      break;

    case diff < 2 * Hour:
      diff -= 1 * Hour;
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.1h');
      break;

    case diff < 1 * Day:
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.hours', {
        msg: Math.floor(diff / Hour)
      });
      diff -= diff / Hour * Hour;
      break;

    case diff < 2 * Day:
      diff -= 1 * Day;
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.1d');
      break;

    case diff < 1 * Week:
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.days', {
        msg: Math.floor(diff / Day)
      });
      diff -= diff / Day * Day;
      break;

    case diff < 2 * Week:
      diff -= 1 * Week;
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.1w');
      break;

    case diff < 1 * Month:
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.weeks', {
        msg: Math.floor(diff / Week)
      });
      diff -= diff / Week * Week;
      break;

    case diff < 2 * Month:
      diff -= 1 * Month;
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.1mon');
      break;

    case diff < 1 * Year:
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.months', {
        msg: Math.floor(diff / Month)
      });
      diff -= diff / Month * Month;
      break;

    case diff < 2 * Year:
      diff -= 1 * Year;
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.1y');
      break;

    default:
      diffStr = _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.years', {
        msg: Math.floor(diff / Year)
      });
      diff -= diff / Year * Year;
      break;
  }

  return {
    diff: diff,
    diffStr: diffStr
  };
};

var timeSinceUnix = function timeSinceUnix(then, now) {
  var lbl = 'timeObj.ago';
  var diff = now - then;

  if (then > now) {
    lbl = 'timeObj.from_now';
    diff = then - now;
  }

  if (diff <= 10) {
    return _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t('timeObj.now');
  }

  var out = computeTimeDiff(diff);
  return _langs__WEBPACK_IMPORTED_MODULE_26__.i18n.t(lbl, {
    msg: out.diffStr
  });
};
var setWebpackPublicPath = function setWebpackPublicPath() {
  var _window$config;

  // This sets up webpack's chunk loading to load resources from the 'public'
  // directory. This file must be imported before any lazy-loading is being attempted.
  if (document.currentScript && document.currentScript.src) {
    var url = new URL(document.currentScript.src);
    __webpack_require__.p = url.pathname.replace(/\/[^/]*?\/[^/]*?$/, '/');
  } else {
    // compat: IE11
    var script = document.querySelector('script[src*="/index.js"]');
    __webpack_require__.p = script.getAttribute('src').replace(/\/[^/]*?\/[^/]*?$/, '/');
  } // rewrite to solve dynamic import cache error


  var AppVer = ((_window$config = window.config) === null || _window$config === void 0 ? void 0 : _window$config.AppVer) || Math.random().toString().replace('0.', '');
  var ElementAppendChild = Element.prototype.appendChild;

  Element.prototype.appendChild = function (node) {
    if (node.tagName == 'SCRIPT' && node.src.indexOf('?') < 0) {
      node.src = node.src + '?v=' + AppVer;
    }

    if (node.tagName == 'LINK' && node.href.indexOf('?') < 0) {
      node.href = node.href + '?v=' + AppVer;
    }

    return ElementAppendChild.call(this, node);
  };
};
var initClipboard = function initClipboard(_els) {
  var els = _els || document.querySelectorAll(".clipboard");

  if (!els || !els.length) return;
  window.$ && $().popup && $(els).popup();
  var clipboard = new (clipboard__WEBPACK_IMPORTED_MODULE_28___default())(els);
  clipboard.on("success", function (e) {
    e.clearSelection();
    var popUpEl = $(e.trigger);
    popUpEl.popup("destroy");
    e.trigger.setAttribute("data-content", e.trigger.getAttribute("data-success"));
    popUpEl.popup("show");
    e.trigger.setAttribute("data-content", e.trigger.getAttribute("data-original"));
  });
  clipboard.on("error", function (e) {
    var popUpEl = $(e.trigger);
    popUpEl.popup("destroy");
    e.trigger.setAttribute("data-content", e.trigger.getAttribute("data-error"));
    popUpEl.popup("show");
    e.trigger.setAttribute("data-content", e.trigger.getAttribute("data-original"));
  });
  return clipboard;
};
var escapeHTML = function escapeHTML(str) {
  return str.replace(/[<>&"']/g, function (match) {
    switch (match) {
      case '<':
        return '&lt;';

      case '>':
        return '&gt;';

      case '&':
        return '&amp;';

      case '"':
        return '&quot;';

      case "'":
        return '&#39;';
    }
  });
};
var RandomGenerator = /*#__PURE__*/function () {
  function RandomGenerator() {
    _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_23___default()(this, RandomGenerator);
  }

  _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_24___default()(RandomGenerator, null, [{
    key: "numbers",
    // 生成随机数字
    value: function numbers() {
      var length = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 4;
      var result = '';

      for (var i = 0; i < length; i++) {
        result += Math.floor(Math.random() * 10);
      }

      return result;
    } // 生成随机字母

  }, {
    key: "letters",
    value: function letters() {
      var length = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 4;
      var uppercase = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : false;
      var result = '';
      var chars = uppercase ? 'ABCDEFGHIJKLMNOPQRSTUVWXYZ' : 'abcdefghijklmnopqrstuvwxyz';

      for (var i = 0; i < length; i++) {
        result += chars.charAt(Math.floor(Math.random() * chars.length));
      }

      return result;
    } // 生成字母数字混合

  }, {
    key: "alphanumeric",
    value: function alphanumeric() {
      var length = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 4;
      var result = '';
      var chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

      for (var i = 0; i < length; i++) {
        result += chars.charAt(Math.floor(Math.random() * chars.length));
      }

      return result;
    } // 生成自定义字符集的随机码

  }, {
    key: "custom",
    value: function custom(charset) {
      var length = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 4;
      var result = '';

      for (var i = 0; i < length; i++) {
        result += charset.charAt(Math.floor(Math.random() * charset.length));
      }

      return result;
    }
  }]);

  return RandomGenerator;
}();
var toBoolean = function toBoolean(value) {
  return value === 'true' ? true : value === 'false' ? false : !!value;
};

/***/ }),

/***/ "./web_src/vuepages/utils/letteravatar.js":
/*!************************************************!*\
  !*** ./web_src/vuepages/utils/letteravatar.js ***!
  \************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_3__);





function LetterAvatar(name, size, color) {
  name = name || "";
  size = size || 60;
  var colours = ["#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#34495e", "#16a085", "#27ae60", "#2980b9", "#8e44ad", "#2c3e50", "#f1c40f", "#e67e22", "#e74c3c", "#00bcd4", "#95a5a6", "#f39c12", "#d35400", "#c0392b", "#bdc3c7", "#7f8c8d"],
      nameSplit = String(name).split(" "),
      initials,
      charIndex,
      colourIndex,
      canvas,
      context,
      dataURI;

  if (nameSplit.length == 1) {
    initials = nameSplit[0] ? nameSplit[0].charAt(0) : "?";
  } else {
    initials = nameSplit[0].charAt(0) + nameSplit[1].charAt(0);
  }

  var initials1 = initials.toUpperCase();

  if (window.devicePixelRatio) {
    size = size * window.devicePixelRatio;
  }

  charIndex = (initials == "?" ? 72 : initials.charCodeAt(0)) - 64;
  colourIndex = charIndex % 20;
  canvas = document.createElement("canvas");
  canvas.width = size;
  canvas.height = size;
  context = canvas.getContext("2d");
  context.fillStyle = color ? color : colours[colourIndex - 1];
  context.fillRect(0, 0, canvas.width, canvas.height);
  context.font = Math.round(canvas.width / 2) + "px 'Microsoft Yahei'";
  context.textAlign = "center";
  context.fillStyle = "#FFF";
  context.fillText(initials1, size / 2, size / 1.5);
  dataURI = canvas.toDataURL();
  canvas = null;
  return dataURI;
}

LetterAvatar.transform = function () {
  Array.prototype.forEach.call(document.querySelectorAll("img[avatar]"), function (img, name, color) {
    name = img.getAttribute("avatar");
    color = img.getAttribute("color");
    img.src = LetterAvatar(name, img.getAttribute("width"), color);
    img.removeAttribute("avatar");
    img.setAttribute("alt", name);
  });
};

/* harmony default export */ __webpack_exports__["default"] = (LetterAvatar);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/style/index.css":
/*!*********************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/style/index.css ***!
  \*********************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=style&index=0&id=2e7d188a&scoped=true&lang=less":
/*!*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=style&index=0&id=2e7d188a&scoped=true&lang=less ***!
  \*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=style&index=0&id=cf90edac&scoped=true&lang=less":
/*!*****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=style&index=0&id=cf90edac&scoped=true&lang=less ***!
  \*****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=style&index=0&id=6ad19374&scoped=true&lang=less":
/*!**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=style&index=0&id=6ad19374&scoped=true&lang=less ***!
  \**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=style&index=0&id=74d70b7f&scoped=true&lang=less":
/*!**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=style&index=0&id=74d70b7f&scoped=true&lang=less ***!
  \**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=style&index=0&id=6b484811&scoped=true&lang=less":
/*!*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=style&index=0&id=6b484811&scoped=true&lang=less ***!
  \*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=style&index=0&id=cf1d6d6e&scoped=true&lang=css":
/*!***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=style&index=0&id=cf1d6d6e&scoped=true&lang=css ***!
  \***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=style&index=0&id=42a186fb&scoped=true&lang=css":
/*!*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=style&index=0&id=42a186fb&scoped=true&lang=css ***!
  \*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=style&index=0&id=48f3bf6b&scoped=true&lang=less":
/*!********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=style&index=0&id=48f3bf6b&scoped=true&lang=less ***!
  \********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=style&index=0&id=b0252fe4&scoped=true&lang=less":
/*!*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=style&index=0&id=b0252fe4&scoped=true&lang=less ***!
  \*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=style&index=0&id=885791ca&scoped=true&lang=less":
/*!****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=style&index=0&id=885791ca&scoped=true&lang=less ***!
  \****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=style&index=0&id=8f7ec232&scoped=true&lang=less":
/*!****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=style&index=0&id=8f7ec232&scoped=true&lang=less ***!
  \****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=style&index=0&id=0230c3f3&scoped=true&lang=css":
/*!*****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=style&index=0&id=0230c3f3&scoped=true&lang=css ***!
  \*****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=style&index=0&id=6586b47e&scoped=true&lang=css":
/*!****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=style&index=0&id=6586b47e&scoped=true&lang=css ***!
  \****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=style&index=0&id=13112268&scoped=true&lang=css":
/*!***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=style&index=0&id=13112268&scoped=true&lang=css ***!
  \***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=0&id=c76be938&scoped=true&lang=css":
/*!*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=0&id=c76be938&scoped=true&lang=css ***!
  \*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=1&id=c76be938&scoped=true&lang=css":
/*!*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=1&id=c76be938&scoped=true&lang=css ***!
  \*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=style&index=0&id=6cebc5e9&scoped=true&lang=css":
/*!*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=style&index=0&id=6cebc5e9&scoped=true&lang=css ***!
  \*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=style&index=0&id=4aaca311&scoped=true&lang=less":
/*!*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=style&index=0&id=4aaca311&scoped=true&lang=less ***!
  \*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=style&index=0&id=726aa430&scoped=true&lang=css":
/*!************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=style&index=0&id=726aa430&scoped=true&lang=css ***!
  \************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=style&index=0&id=8071ebce&scoped=true&lang=less":
/*!***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=style&index=0&id=8071ebce&scoped=true&lang=less ***!
  \***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=style&index=0&id=c265b356&scoped=true&lang=css":
/*!***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=style&index=0&id=c265b356&scoped=true&lang=css ***!
  \***************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue":
/*!*********************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue ***!
  \*********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ModelCondition_vue_vue_type_template_id_2e7d188a_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ModelCondition.vue?vue&type=template&id=2e7d188a&scoped=true */ "./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=template&id=2e7d188a&scoped=true");
/* harmony import */ var _ModelCondition_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ModelCondition.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=script&lang=js");
/* harmony import */ var _ModelCondition_vue_vue_type_style_index_0_id_2e7d188a_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ModelCondition.vue?vue&type=style&index=0&id=2e7d188a&scoped=true&lang=less */ "./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=style&index=0&id=2e7d188a&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ModelCondition_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ModelCondition_vue_vue_type_template_id_2e7d188a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ModelCondition_vue_vue_type_template_id_2e7d188a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "2e7d188a",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue":
/*!*******************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue ***!
  \*******************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ModelFilters_vue_vue_type_template_id_cf90edac_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ModelFilters.vue?vue&type=template&id=cf90edac&scoped=true */ "./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=template&id=cf90edac&scoped=true");
/* harmony import */ var _ModelFilters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ModelFilters.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=script&lang=js");
/* harmony import */ var _ModelFilters_vue_vue_type_style_index_0_id_cf90edac_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ModelFilters.vue?vue&type=style&index=0&id=cf90edac&scoped=true&lang=less */ "./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=style&index=0&id=cf90edac&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ModelFilters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ModelFilters_vue_vue_type_template_id_cf90edac_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ModelFilters_vue_vue_type_template_id_cf90edac_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "cf90edac",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue":
/*!****************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ModelItem_vue_vue_type_template_id_6ad19374_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ModelItem.vue?vue&type=template&id=6ad19374&scoped=true */ "./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=template&id=6ad19374&scoped=true");
/* harmony import */ var _ModelItem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ModelItem.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=script&lang=js");
/* harmony import */ var _ModelItem_vue_vue_type_style_index_0_id_6ad19374_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ModelItem.vue?vue&type=style&index=0&id=6ad19374&scoped=true&lang=less */ "./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=style&index=0&id=6ad19374&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ModelItem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ModelItem_vue_vue_type_template_id_6ad19374_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ModelItem_vue_vue_type_template_id_6ad19374_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "6ad19374",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue":
/*!****************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue ***!
  \****************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ModelList_vue_vue_type_template_id_74d70b7f_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ModelList.vue?vue&type=template&id=74d70b7f&scoped=true */ "./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=template&id=74d70b7f&scoped=true");
/* harmony import */ var _ModelList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ModelList.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=script&lang=js");
/* harmony import */ var _ModelList_vue_vue_type_style_index_0_id_74d70b7f_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ModelList.vue?vue&type=style&index=0&id=74d70b7f&scoped=true&lang=less */ "./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=style&index=0&id=74d70b7f&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ModelList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ModelList_vue_vue_type_template_id_74d70b7f_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ModelList_vue_vue_type_template_id_74d70b7f_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "74d70b7f",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/modelsquare/square/components/ModelList.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/repos/components/ReposItem.vue":
/*!***************************************************************!*\
  !*** ./web_src/vuepages/pages/repos/components/ReposItem.vue ***!
  \***************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ReposItem_vue_vue_type_template_id_6b484811_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ReposItem.vue?vue&type=template&id=6b484811&scoped=true */ "./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=template&id=6b484811&scoped=true");
/* harmony import */ var _ReposItem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ReposItem.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=script&lang=js");
/* harmony import */ var _ReposItem_vue_vue_type_style_index_0_id_6b484811_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ReposItem.vue?vue&type=style&index=0&id=6b484811&scoped=true&lang=less */ "./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=style&index=0&id=6b484811&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ReposItem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ReposItem_vue_vue_type_template_id_6b484811_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ReposItem_vue_vue_type_template_id_6b484811_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "6b484811",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/repos/components/ReposItem.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/App.vue":
/*!*************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/App.vue ***!
  \*************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _App_vue_vue_type_template_id_b942a2c8__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./App.vue?vue&type=template&id=b942a2c8 */ "./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=template&id=b942a2c8");
/* harmony import */ var _App_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./App.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=script&lang=js");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! !../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");





/* normalize component */
;
var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
  _App_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _App_vue_vue_type_template_id_b942a2c8__WEBPACK_IMPORTED_MODULE_0__.render,
  _App_vue_vue_type_template_id_b942a2c8__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  null,
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/App.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue":
/*!******************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue ***!
  \******************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _AppBanner_vue_vue_type_template_id_cf1d6d6e_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./AppBanner.vue?vue&type=template&id=cf1d6d6e&scoped=true */ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=template&id=cf1d6d6e&scoped=true");
/* harmony import */ var _AppBanner_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./AppBanner.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=script&lang=js");
/* harmony import */ var _AppBanner_vue_vue_type_style_index_0_id_cf1d6d6e_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./AppBanner.vue?vue&type=style&index=0&id=cf1d6d6e&scoped=true&lang=css */ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=style&index=0&id=cf1d6d6e&scoped=true&lang=css");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _AppBanner_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _AppBanner_vue_vue_type_template_id_cf1d6d6e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _AppBanner_vue_vue_type_template_id_cf1d6d6e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "cf1d6d6e",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/components/AppBanner.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue":
/*!**********************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _NavigationBar_vue_vue_type_template_id_42a186fb_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./NavigationBar.vue?vue&type=template&id=42a186fb&scoped=true */ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=template&id=42a186fb&scoped=true");
/* harmony import */ var _NavigationBar_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./NavigationBar.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=script&lang=js");
/* harmony import */ var _NavigationBar_vue_vue_type_style_index_0_id_42a186fb_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./NavigationBar.vue?vue&type=style&index=0&id=42a186fb&scoped=true&lang=css */ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=style&index=0&id=42a186fb&scoped=true&lang=css");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _NavigationBar_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _NavigationBar_vue_vue_type_template_id_42a186fb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _NavigationBar_vue_vue_type_template_id_42a186fb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "42a186fb",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue":
/*!**********************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _PublicDataset_vue_vue_type_template_id_48f3bf6b_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./PublicDataset.vue?vue&type=template&id=48f3bf6b&scoped=true */ "./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=template&id=48f3bf6b&scoped=true");
/* harmony import */ var _PublicDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./PublicDataset.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=script&lang=js");
/* harmony import */ var _PublicDataset_vue_vue_type_style_index_0_id_48f3bf6b_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./PublicDataset.vue?vue&type=style&index=0&id=48f3bf6b&scoped=true&lang=less */ "./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=style&index=0&id=48f3bf6b&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _PublicDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _PublicDataset_vue_vue_type_template_id_48f3bf6b_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _PublicDataset_vue_vue_type_template_id_48f3bf6b_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "48f3bf6b",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue":
/*!*********************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue ***!
  \*********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ReposFilters_vue_vue_type_template_id_b0252fe4_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ReposFilters.vue?vue&type=template&id=b0252fe4&scoped=true */ "./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=template&id=b0252fe4&scoped=true");
/* harmony import */ var _ReposFilters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ReposFilters.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=script&lang=js");
/* harmony import */ var _ReposFilters_vue_vue_type_style_index_0_id_b0252fe4_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ReposFilters.vue?vue&type=style&index=0&id=b0252fe4&scoped=true&lang=less */ "./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=style&index=0&id=b0252fe4&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ReposFilters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ReposFilters_vue_vue_type_template_id_b0252fe4_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ReposFilters_vue_vue_type_template_id_b0252fe4_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "b0252fe4",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/ReposList.vue":
/*!******************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/ReposList.vue ***!
  \******************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ReposList_vue_vue_type_template_id_885791ca_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ReposList.vue?vue&type=template&id=885791ca&scoped=true */ "./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=template&id=885791ca&scoped=true");
/* harmony import */ var _ReposList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ReposList.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=script&lang=js");
/* harmony import */ var _ReposList_vue_vue_type_style_index_0_id_885791ca_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ReposList.vue?vue&type=style&index=0&id=885791ca&scoped=true&lang=less */ "./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=style&index=0&id=885791ca&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ReposList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ReposList_vue_vue_type_template_id_885791ca_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ReposList_vue_vue_type_template_id_885791ca_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "885791ca",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/components/ReposList.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue":
/*!******************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue ***!
  \******************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _SearchBar_vue_vue_type_template_id_8f7ec232_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./SearchBar.vue?vue&type=template&id=8f7ec232&scoped=true */ "./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=template&id=8f7ec232&scoped=true");
/* harmony import */ var _SearchBar_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./SearchBar.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=script&lang=js");
/* harmony import */ var _SearchBar_vue_vue_type_style_index_0_id_8f7ec232_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./SearchBar.vue?vue&type=style&index=0&id=8f7ec232&scoped=true&lang=less */ "./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=style&index=0&id=8f7ec232&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _SearchBar_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _SearchBar_vue_vue_type_template_id_8f7ec232_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _SearchBar_vue_vue_type_template_id_8f7ec232_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "8f7ec232",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/components/SearchBar.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue":
/*!********************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue ***!
  \********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _helpCollaps_vue_vue_type_template_id_0230c3f3_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./helpCollaps.vue?vue&type=template&id=0230c3f3&scoped=true */ "./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=template&id=0230c3f3&scoped=true");
/* harmony import */ var _helpCollaps_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./helpCollaps.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=script&lang=js");
/* harmony import */ var _helpCollaps_vue_vue_type_style_index_0_id_0230c3f3_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./helpCollaps.vue?vue&type=style&index=0&id=0230c3f3&scoped=true&lang=css */ "./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=style&index=0&id=0230c3f3&scoped=true&lang=css");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _helpCollaps_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _helpCollaps_vue_vue_type_template_id_0230c3f3_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _helpCollaps_vue_vue_type_template_id_0230c3f3_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "0230c3f3",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue":
/*!*******************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue ***!
  \*******************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _CommunitySource_vue_vue_type_template_id_6586b47e_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./CommunitySource.vue?vue&type=template&id=6586b47e&scoped=true */ "./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=template&id=6586b47e&scoped=true");
/* harmony import */ var _CommunitySource_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./CommunitySource.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=script&lang=js");
/* harmony import */ var _CommunitySource_vue_vue_type_style_index_0_id_6586b47e_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./CommunitySource.vue?vue&type=style&index=0&id=6586b47e&scoped=true&lang=css */ "./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=style&index=0&id=6586b47e&scoped=true&lang=css");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _CommunitySource_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _CommunitySource_vue_vue_type_template_id_6586b47e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _CommunitySource_vue_vue_type_template_id_6586b47e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "6586b47e",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue":
/*!**************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue ***!
  \**************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _HelpCenter_vue_vue_type_template_id_13112268_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./HelpCenter.vue?vue&type=template&id=13112268&scoped=true */ "./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=template&id=13112268&scoped=true");
/* harmony import */ var _HelpCenter_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./HelpCenter.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=script&lang=js");
/* harmony import */ var _HelpCenter_vue_vue_type_style_index_0_id_13112268_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./HelpCenter.vue?vue&type=style&index=0&id=13112268&scoped=true&lang=css */ "./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=style&index=0&id=13112268&scoped=true&lang=css");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _HelpCenter_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _HelpCenter_vue_vue_type_template_id_13112268_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _HelpCenter_vue_vue_type_template_id_13112268_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "13112268",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/HomePage.vue":
/*!************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/HomePage.vue ***!
  \************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _HomePage_vue_vue_type_template_id_c76be938_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./HomePage.vue?vue&type=template&id=c76be938&scoped=true */ "./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=template&id=c76be938&scoped=true");
/* harmony import */ var _HomePage_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./HomePage.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=script&lang=js");
/* harmony import */ var _HomePage_vue_vue_type_style_index_0_id_c76be938_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./HomePage.vue?vue&type=style&index=0&id=c76be938&scoped=true&lang=css */ "./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=0&id=c76be938&scoped=true&lang=css");
/* harmony import */ var _HomePage_vue_vue_type_style_index_1_id_c76be938_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./HomePage.vue?vue&type=style&index=1&id=c76be938&scoped=true&lang=css */ "./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=1&id=c76be938&scoped=true&lang=css");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;



/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_4__["default"])(
  _HomePage_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _HomePage_vue_vue_type_template_id_c76be938_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _HomePage_vue_vue_type_template_id_c76be938_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "c76be938",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/views/HomePage.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue":
/*!************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue ***!
  \************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _OSSystem_vue_vue_type_template_id_6cebc5e9_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./OSSystem.vue?vue&type=template&id=6cebc5e9&scoped=true */ "./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=template&id=6cebc5e9&scoped=true");
/* harmony import */ var _OSSystem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./OSSystem.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=script&lang=js");
/* harmony import */ var _OSSystem_vue_vue_type_style_index_0_id_6cebc5e9_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./OSSystem.vue?vue&type=style&index=0&id=6cebc5e9&scoped=true&lang=css */ "./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=style&index=0&id=6cebc5e9&scoped=true&lang=css");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _OSSystem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _OSSystem_vue_vue_type_template_id_6cebc5e9_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _OSSystem_vue_vue_type_template_id_6cebc5e9_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "6cebc5e9",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/views/OSSystem.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue":
/*!***********************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue ***!
  \***********************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _OpenApp_vue_vue_type_template_id_4aaca311_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./OpenApp.vue?vue&type=template&id=4aaca311&scoped=true */ "./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=template&id=4aaca311&scoped=true");
/* harmony import */ var _OpenApp_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./OpenApp.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=script&lang=js");
/* harmony import */ var _OpenApp_vue_vue_type_style_index_0_id_4aaca311_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./OpenApp.vue?vue&type=style&index=0&id=4aaca311&scoped=true&lang=less */ "./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=style&index=0&id=4aaca311&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _OpenApp_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _OpenApp_vue_vue_type_template_id_4aaca311_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _OpenApp_vue_vue_type_template_id_4aaca311_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "4aaca311",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/views/OpenApp.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue":
/*!***************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue ***!
  \***************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _OpenDataset_vue_vue_type_template_id_726aa430_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./OpenDataset.vue?vue&type=template&id=726aa430&scoped=true */ "./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=template&id=726aa430&scoped=true");
/* harmony import */ var _OpenDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./OpenDataset.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=script&lang=js");
/* harmony import */ var _OpenDataset_vue_vue_type_style_index_0_id_726aa430_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./OpenDataset.vue?vue&type=style&index=0&id=726aa430&scoped=true&lang=css */ "./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=style&index=0&id=726aa430&scoped=true&lang=css");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _OpenDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _OpenDataset_vue_vue_type_template_id_726aa430_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _OpenDataset_vue_vue_type_template_id_726aa430_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "726aa430",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue":
/*!*************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue ***!
  \*************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _OpenModel_vue_vue_type_template_id_8071ebce_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./OpenModel.vue?vue&type=template&id=8071ebce&scoped=true */ "./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=template&id=8071ebce&scoped=true");
/* harmony import */ var _OpenModel_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./OpenModel.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=script&lang=js");
/* harmony import */ var _OpenModel_vue_vue_type_style_index_0_id_8071ebce_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./OpenModel.vue?vue&type=style&index=0&id=8071ebce&scoped=true&lang=less */ "./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=style&index=0&id=8071ebce&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _OpenModel_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _OpenModel_vue_vue_type_template_id_8071ebce_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _OpenModel_vue_vue_type_template_id_8071ebce_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "8071ebce",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/views/OpenModel.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue":
/*!******************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue ***!
  \******************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ResourceDetail_vue_vue_type_template_id_c265b356_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ResourceDetail.vue?vue&type=template&id=c265b356&scoped=true */ "./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=template&id=c265b356&scoped=true");
/* harmony import */ var _ResourceDetail_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ResourceDetail.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=script&lang=js");
/* harmony import */ var _ResourceDetail_vue_vue_type_style_index_0_id_c265b356_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ResourceDetail.vue?vue&type=style&index=0&id=c265b356&scoped=true&lang=css */ "./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=style&index=0&id=c265b356&scoped=true&lang=css");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _ResourceDetail_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _ResourceDetail_vue_vue_type_template_id_c265b356_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _ResourceDetail_vue_vue_type_template_id_c265b356_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "c265b356",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=script&lang=js":
/*!*********************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=script&lang=js ***!
  \*********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelCondition_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelCondition.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelCondition_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=script&lang=js":
/*!*******************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=script&lang=js ***!
  \*******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelFilters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelFilters.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelFilters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=script&lang=js":
/*!****************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=script&lang=js ***!
  \****************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelItem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelItem.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelItem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=script&lang=js":
/*!****************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=script&lang=js ***!
  \****************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelList.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=script&lang=js":
/*!***************************************************************************************!*\
  !*** ./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=script&lang=js ***!
  \***************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposItem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ReposItem.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposItem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=script&lang=js":
/*!*************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=script&lang=js ***!
  \*************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_App_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./App.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_App_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=script&lang=js":
/*!******************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=script&lang=js ***!
  \******************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_AppBanner_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./AppBanner.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_AppBanner_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=script&lang=js":
/*!**********************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_NavigationBar_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./NavigationBar.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_NavigationBar_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=script&lang=js":
/*!**********************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_PublicDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./PublicDataset.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_PublicDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=script&lang=js":
/*!*********************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=script&lang=js ***!
  \*********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposFilters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ReposFilters.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposFilters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=script&lang=js":
/*!******************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=script&lang=js ***!
  \******************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ReposList.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=script&lang=js":
/*!******************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=script&lang=js ***!
  \******************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_SearchBar_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./SearchBar.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_SearchBar_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=script&lang=js":
/*!********************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=script&lang=js ***!
  \********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_helpCollaps_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./helpCollaps.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_helpCollaps_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=script&lang=js":
/*!*******************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=script&lang=js ***!
  \*******************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_CommunitySource_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./CommunitySource.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_CommunitySource_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=script&lang=js":
/*!**************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=script&lang=js ***!
  \**************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_HelpCenter_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./HelpCenter.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_HelpCenter_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=script&lang=js":
/*!************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=script&lang=js ***!
  \************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_HomePage_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./HomePage.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_HomePage_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=script&lang=js":
/*!************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=script&lang=js ***!
  \************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_OSSystem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OSSystem.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_OSSystem_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=script&lang=js":
/*!***********************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=script&lang=js ***!
  \***********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenApp_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OpenApp.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenApp_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=script&lang=js":
/*!***************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=script&lang=js ***!
  \***************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OpenDataset.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenDataset_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=script&lang=js":
/*!*************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=script&lang=js ***!
  \*************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenModel_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OpenModel.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenModel_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=script&lang=js":
/*!******************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=script&lang=js ***!
  \******************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceDetail_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ResourceDetail.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceDetail_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=style&index=0&id=2e7d188a&scoped=true&lang=less":
/*!******************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=style&index=0&id=2e7d188a&scoped=true&lang=less ***!
  \******************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelCondition_vue_vue_type_style_index_0_id_2e7d188a_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelCondition.vue?vue&type=style&index=0&id=2e7d188a&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=style&index=0&id=2e7d188a&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=style&index=0&id=cf90edac&scoped=true&lang=less":
/*!****************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=style&index=0&id=cf90edac&scoped=true&lang=less ***!
  \****************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelFilters_vue_vue_type_style_index_0_id_cf90edac_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelFilters.vue?vue&type=style&index=0&id=cf90edac&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=style&index=0&id=cf90edac&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=style&index=0&id=6ad19374&scoped=true&lang=less":
/*!*************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=style&index=0&id=6ad19374&scoped=true&lang=less ***!
  \*************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelItem_vue_vue_type_style_index_0_id_6ad19374_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelItem.vue?vue&type=style&index=0&id=6ad19374&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=style&index=0&id=6ad19374&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=style&index=0&id=74d70b7f&scoped=true&lang=less":
/*!*************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=style&index=0&id=74d70b7f&scoped=true&lang=less ***!
  \*************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelList_vue_vue_type_style_index_0_id_74d70b7f_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelList.vue?vue&type=style&index=0&id=74d70b7f&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=style&index=0&id=74d70b7f&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=style&index=0&id=6b484811&scoped=true&lang=less":
/*!************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=style&index=0&id=6b484811&scoped=true&lang=less ***!
  \************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposItem_vue_vue_type_style_index_0_id_6b484811_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ReposItem.vue?vue&type=style&index=0&id=6b484811&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=style&index=0&id=6b484811&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=style&index=0&id=cf1d6d6e&scoped=true&lang=css":
/*!**************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=style&index=0&id=cf1d6d6e&scoped=true&lang=css ***!
  \**************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_AppBanner_vue_vue_type_style_index_0_id_cf1d6d6e_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./AppBanner.vue?vue&type=style&index=0&id=cf1d6d6e&scoped=true&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=style&index=0&id=cf1d6d6e&scoped=true&lang=css");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=style&index=0&id=42a186fb&scoped=true&lang=css":
/*!******************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=style&index=0&id=42a186fb&scoped=true&lang=css ***!
  \******************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_NavigationBar_vue_vue_type_style_index_0_id_42a186fb_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./NavigationBar.vue?vue&type=style&index=0&id=42a186fb&scoped=true&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=style&index=0&id=42a186fb&scoped=true&lang=css");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=style&index=0&id=48f3bf6b&scoped=true&lang=less":
/*!*******************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=style&index=0&id=48f3bf6b&scoped=true&lang=less ***!
  \*******************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_PublicDataset_vue_vue_type_style_index_0_id_48f3bf6b_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./PublicDataset.vue?vue&type=style&index=0&id=48f3bf6b&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=style&index=0&id=48f3bf6b&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=style&index=0&id=b0252fe4&scoped=true&lang=less":
/*!******************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=style&index=0&id=b0252fe4&scoped=true&lang=less ***!
  \******************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposFilters_vue_vue_type_style_index_0_id_b0252fe4_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ReposFilters.vue?vue&type=style&index=0&id=b0252fe4&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=style&index=0&id=b0252fe4&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=style&index=0&id=885791ca&scoped=true&lang=less":
/*!***************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=style&index=0&id=885791ca&scoped=true&lang=less ***!
  \***************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposList_vue_vue_type_style_index_0_id_885791ca_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ReposList.vue?vue&type=style&index=0&id=885791ca&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=style&index=0&id=885791ca&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=style&index=0&id=8f7ec232&scoped=true&lang=less":
/*!***************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=style&index=0&id=8f7ec232&scoped=true&lang=less ***!
  \***************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_SearchBar_vue_vue_type_style_index_0_id_8f7ec232_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./SearchBar.vue?vue&type=style&index=0&id=8f7ec232&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=style&index=0&id=8f7ec232&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=style&index=0&id=0230c3f3&scoped=true&lang=css":
/*!****************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=style&index=0&id=0230c3f3&scoped=true&lang=css ***!
  \****************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_helpCollaps_vue_vue_type_style_index_0_id_0230c3f3_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./helpCollaps.vue?vue&type=style&index=0&id=0230c3f3&scoped=true&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=style&index=0&id=0230c3f3&scoped=true&lang=css");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=style&index=0&id=6586b47e&scoped=true&lang=css":
/*!***************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=style&index=0&id=6586b47e&scoped=true&lang=css ***!
  \***************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_CommunitySource_vue_vue_type_style_index_0_id_6586b47e_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./CommunitySource.vue?vue&type=style&index=0&id=6586b47e&scoped=true&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=style&index=0&id=6586b47e&scoped=true&lang=css");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=style&index=0&id=13112268&scoped=true&lang=css":
/*!**********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=style&index=0&id=13112268&scoped=true&lang=css ***!
  \**********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_HelpCenter_vue_vue_type_style_index_0_id_13112268_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./HelpCenter.vue?vue&type=style&index=0&id=13112268&scoped=true&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=style&index=0&id=13112268&scoped=true&lang=css");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=0&id=c76be938&scoped=true&lang=css":
/*!********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=0&id=c76be938&scoped=true&lang=css ***!
  \********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_HomePage_vue_vue_type_style_index_0_id_c76be938_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./HomePage.vue?vue&type=style&index=0&id=c76be938&scoped=true&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=0&id=c76be938&scoped=true&lang=css");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=1&id=c76be938&scoped=true&lang=css":
/*!********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=1&id=c76be938&scoped=true&lang=css ***!
  \********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_HomePage_vue_vue_type_style_index_1_id_c76be938_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./HomePage.vue?vue&type=style&index=1&id=c76be938&scoped=true&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=style&index=1&id=c76be938&scoped=true&lang=css");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=style&index=0&id=6cebc5e9&scoped=true&lang=css":
/*!********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=style&index=0&id=6cebc5e9&scoped=true&lang=css ***!
  \********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_OSSystem_vue_vue_type_style_index_0_id_6cebc5e9_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OSSystem.vue?vue&type=style&index=0&id=6cebc5e9&scoped=true&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=style&index=0&id=6cebc5e9&scoped=true&lang=css");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=style&index=0&id=4aaca311&scoped=true&lang=less":
/*!********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=style&index=0&id=4aaca311&scoped=true&lang=less ***!
  \********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenApp_vue_vue_type_style_index_0_id_4aaca311_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OpenApp.vue?vue&type=style&index=0&id=4aaca311&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=style&index=0&id=4aaca311&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=style&index=0&id=726aa430&scoped=true&lang=css":
/*!***********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=style&index=0&id=726aa430&scoped=true&lang=css ***!
  \***********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenDataset_vue_vue_type_style_index_0_id_726aa430_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OpenDataset.vue?vue&type=style&index=0&id=726aa430&scoped=true&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=style&index=0&id=726aa430&scoped=true&lang=css");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=style&index=0&id=8071ebce&scoped=true&lang=less":
/*!**********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=style&index=0&id=8071ebce&scoped=true&lang=less ***!
  \**********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenModel_vue_vue_type_style_index_0_id_8071ebce_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OpenModel.vue?vue&type=style&index=0&id=8071ebce&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=style&index=0&id=8071ebce&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=style&index=0&id=c265b356&scoped=true&lang=css":
/*!**************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=style&index=0&id=c265b356&scoped=true&lang=css ***!
  \**************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceDetail_vue_vue_type_style_index_0_id_c265b356_scoped_true_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ResourceDetail.vue?vue&type=style&index=0&id=c265b356&scoped=true&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=style&index=0&id=c265b356&scoped=true&lang=css");


/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=template&id=2e7d188a&scoped=true":
/*!***************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=template&id=2e7d188a&scoped=true ***!
  \***************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelCondition_vue_vue_type_template_id_2e7d188a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelCondition_vue_vue_type_template_id_2e7d188a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelCondition_vue_vue_type_template_id_2e7d188a_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelCondition.vue?vue&type=template&id=2e7d188a&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=template&id=2e7d188a&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=template&id=cf90edac&scoped=true":
/*!*************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=template&id=cf90edac&scoped=true ***!
  \*************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelFilters_vue_vue_type_template_id_cf90edac_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelFilters_vue_vue_type_template_id_cf90edac_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelFilters_vue_vue_type_template_id_cf90edac_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelFilters.vue?vue&type=template&id=cf90edac&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=template&id=cf90edac&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=template&id=6ad19374&scoped=true":
/*!**********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=template&id=6ad19374&scoped=true ***!
  \**********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelItem_vue_vue_type_template_id_6ad19374_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelItem_vue_vue_type_template_id_6ad19374_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelItem_vue_vue_type_template_id_6ad19374_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelItem.vue?vue&type=template&id=6ad19374&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=template&id=6ad19374&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=template&id=74d70b7f&scoped=true":
/*!**********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=template&id=74d70b7f&scoped=true ***!
  \**********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelList_vue_vue_type_template_id_74d70b7f_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelList_vue_vue_type_template_id_74d70b7f_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ModelList_vue_vue_type_template_id_74d70b7f_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ModelList.vue?vue&type=template&id=74d70b7f&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=template&id=74d70b7f&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=template&id=6b484811&scoped=true":
/*!*********************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=template&id=6b484811&scoped=true ***!
  \*********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposItem_vue_vue_type_template_id_6b484811_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposItem_vue_vue_type_template_id_6b484811_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposItem_vue_vue_type_template_id_6b484811_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ReposItem.vue?vue&type=template&id=6b484811&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=template&id=6b484811&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=template&id=b942a2c8":
/*!*******************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=template&id=b942a2c8 ***!
  \*******************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_App_vue_vue_type_template_id_b942a2c8__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_App_vue_vue_type_template_id_b942a2c8__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_App_vue_vue_type_template_id_b942a2c8__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./App.vue?vue&type=template&id=b942a2c8 */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=template&id=b942a2c8");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=template&id=cf1d6d6e&scoped=true":
/*!************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=template&id=cf1d6d6e&scoped=true ***!
  \************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_AppBanner_vue_vue_type_template_id_cf1d6d6e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_AppBanner_vue_vue_type_template_id_cf1d6d6e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_AppBanner_vue_vue_type_template_id_cf1d6d6e_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./AppBanner.vue?vue&type=template&id=cf1d6d6e&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=template&id=cf1d6d6e&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=template&id=42a186fb&scoped=true":
/*!****************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=template&id=42a186fb&scoped=true ***!
  \****************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_NavigationBar_vue_vue_type_template_id_42a186fb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_NavigationBar_vue_vue_type_template_id_42a186fb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_NavigationBar_vue_vue_type_template_id_42a186fb_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./NavigationBar.vue?vue&type=template&id=42a186fb&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=template&id=42a186fb&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=template&id=48f3bf6b&scoped=true":
/*!****************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=template&id=48f3bf6b&scoped=true ***!
  \****************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_PublicDataset_vue_vue_type_template_id_48f3bf6b_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_PublicDataset_vue_vue_type_template_id_48f3bf6b_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_PublicDataset_vue_vue_type_template_id_48f3bf6b_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./PublicDataset.vue?vue&type=template&id=48f3bf6b&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=template&id=48f3bf6b&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=template&id=b0252fe4&scoped=true":
/*!***************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=template&id=b0252fe4&scoped=true ***!
  \***************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposFilters_vue_vue_type_template_id_b0252fe4_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposFilters_vue_vue_type_template_id_b0252fe4_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposFilters_vue_vue_type_template_id_b0252fe4_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ReposFilters.vue?vue&type=template&id=b0252fe4&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=template&id=b0252fe4&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=template&id=885791ca&scoped=true":
/*!************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=template&id=885791ca&scoped=true ***!
  \************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposList_vue_vue_type_template_id_885791ca_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposList_vue_vue_type_template_id_885791ca_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ReposList_vue_vue_type_template_id_885791ca_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ReposList.vue?vue&type=template&id=885791ca&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=template&id=885791ca&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=template&id=8f7ec232&scoped=true":
/*!************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=template&id=8f7ec232&scoped=true ***!
  \************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_SearchBar_vue_vue_type_template_id_8f7ec232_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_SearchBar_vue_vue_type_template_id_8f7ec232_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_SearchBar_vue_vue_type_template_id_8f7ec232_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./SearchBar.vue?vue&type=template&id=8f7ec232&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=template&id=8f7ec232&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=template&id=0230c3f3&scoped=true":
/*!**************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=template&id=0230c3f3&scoped=true ***!
  \**************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_helpCollaps_vue_vue_type_template_id_0230c3f3_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_helpCollaps_vue_vue_type_template_id_0230c3f3_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_helpCollaps_vue_vue_type_template_id_0230c3f3_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./helpCollaps.vue?vue&type=template&id=0230c3f3&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=template&id=0230c3f3&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=template&id=6586b47e&scoped=true":
/*!*************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=template&id=6586b47e&scoped=true ***!
  \*************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_CommunitySource_vue_vue_type_template_id_6586b47e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_CommunitySource_vue_vue_type_template_id_6586b47e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_CommunitySource_vue_vue_type_template_id_6586b47e_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./CommunitySource.vue?vue&type=template&id=6586b47e&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=template&id=6586b47e&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=template&id=13112268&scoped=true":
/*!********************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=template&id=13112268&scoped=true ***!
  \********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_HelpCenter_vue_vue_type_template_id_13112268_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_HelpCenter_vue_vue_type_template_id_13112268_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_HelpCenter_vue_vue_type_template_id_13112268_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./HelpCenter.vue?vue&type=template&id=13112268&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=template&id=13112268&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=template&id=c76be938&scoped=true":
/*!******************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=template&id=c76be938&scoped=true ***!
  \******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_HomePage_vue_vue_type_template_id_c76be938_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_HomePage_vue_vue_type_template_id_c76be938_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_HomePage_vue_vue_type_template_id_c76be938_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./HomePage.vue?vue&type=template&id=c76be938&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=template&id=c76be938&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=template&id=6cebc5e9&scoped=true":
/*!******************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=template&id=6cebc5e9&scoped=true ***!
  \******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OSSystem_vue_vue_type_template_id_6cebc5e9_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OSSystem_vue_vue_type_template_id_6cebc5e9_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OSSystem_vue_vue_type_template_id_6cebc5e9_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OSSystem.vue?vue&type=template&id=6cebc5e9&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=template&id=6cebc5e9&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=template&id=4aaca311&scoped=true":
/*!*****************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=template&id=4aaca311&scoped=true ***!
  \*****************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenApp_vue_vue_type_template_id_4aaca311_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenApp_vue_vue_type_template_id_4aaca311_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenApp_vue_vue_type_template_id_4aaca311_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OpenApp.vue?vue&type=template&id=4aaca311&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=template&id=4aaca311&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=template&id=726aa430&scoped=true":
/*!*********************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=template&id=726aa430&scoped=true ***!
  \*********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenDataset_vue_vue_type_template_id_726aa430_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenDataset_vue_vue_type_template_id_726aa430_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenDataset_vue_vue_type_template_id_726aa430_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OpenDataset.vue?vue&type=template&id=726aa430&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=template&id=726aa430&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=template&id=8071ebce&scoped=true":
/*!*******************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=template&id=8071ebce&scoped=true ***!
  \*******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenModel_vue_vue_type_template_id_8071ebce_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenModel_vue_vue_type_template_id_8071ebce_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_OpenModel_vue_vue_type_template_id_8071ebce_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./OpenModel.vue?vue&type=template&id=8071ebce&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=template&id=8071ebce&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=template&id=c265b356&scoped=true":
/*!************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=template&id=c265b356&scoped=true ***!
  \************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceDetail_vue_vue_type_template_id_c265b356_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceDetail_vue_vue_type_template_id_c265b356_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_ResourceDetail_vue_vue_type_template_id_c265b356_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./ResourceDetail.vue?vue&type=template&id=c265b356&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=template&id=c265b356&scoped=true");


/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=template&id=2e7d188a&scoped=true":
/*!******************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelCondition.vue?vue&type=template&id=2e7d188a&scoped=true ***!
  \******************************************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "condition-wrap" }, [
    _c("div", [
      _c(
        "div",
        { staticClass: "tab-c" },
        _vm._l(_vm.tabList, function(item, index) {
          return _c(
            "div",
            {
              key: item.key,
              staticClass: "tab-item",
              class: _vm.conds.tab == item.key ? "focus" : "",
              on: {
                click: function($event) {
                  return _vm.changeTab(item)
                }
              }
            },
            [_vm._v("\n        " + _vm._s(item.label) + "\n      ")]
          )
        }),
        0
      )
    ]),
    _vm._v(" "),
    _c("div", { staticClass: "condition-b" }, [
      _c(
        "div",
        { staticClass: "only-recommend-c" },
        [
          _c(
            "el-checkbox",
            {
              on: { change: _vm.changeRecommend },
              model: {
                value: _vm.conds.onlyRecommend,
                callback: function($$v) {
                  _vm.$set(_vm.conds, "onlyRecommend", $$v)
                },
                expression: "conds.onlyRecommend"
              }
            },
            [
              _vm._v(
                "\n        " +
                  _vm._s(_vm.$t("datasets.platform_recommendations")) +
                  "\n      "
              )
            ]
          ),
          _vm._v(" "),
          _c(
            "el-checkbox",
            {
              on: { change: _vm.changeOnline },
              model: {
                value: _vm.conds.hasOnlineUrl,
                callback: function($$v) {
                  _vm.$set(_vm.conds, "hasOnlineUrl", $$v)
                },
                expression: "conds.hasOnlineUrl"
              }
            },
            [
              _vm._v(
                "\n        " +
                  _vm._s(_vm.$t("modelObj.can_online_infer")) +
                  "\n      "
              )
            ]
          )
        ],
        1
      ),
      _vm._v(" "),
      _c(
        "div",
        [
          _c(
            "el-dropdown",
            {
              staticClass: "sort-c",
              attrs: { trigger: "click", size: "default" }
            },
            [
              _c("span", { staticClass: "el-dropdown-link" }, [
                _vm._v("\n          " + _vm._s(_vm.$t("datasets.sort"))),
                _c("i", { staticClass: "el-icon-caret-bottom el-icon--right" })
              ]),
              _vm._v(" "),
              _c(
                "el-dropdown-menu",
                { attrs: { slot: "dropdown" }, slot: "dropdown" },
                _vm._l(_vm.sortList, function(item) {
                  return _c(
                    "el-dropdown-item",
                    {
                      key: item.key,
                      class: _vm.conds.sort == item.key ? "active" : "",
                      nativeOn: {
                        click: function($event) {
                          return _vm.changeSort(item)
                        }
                      }
                    },
                    [
                      _vm._v(
                        "\n            " + _vm._s(item.label) + "\n          "
                      )
                    ]
                  )
                }),
                1
              )
            ],
            1
          )
        ],
        1
      )
    ])
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=template&id=cf90edac&scoped=true":
/*!****************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelFilters.vue?vue&type=template&id=cf90edac&scoped=true ***!
  \****************************************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "container" }, [
    _c("div", { staticClass: "title" }, [
      _vm._v(_vm._s(_vm.$t("repos.model")))
    ]),
    _vm._v(" "),
    _c("div", { staticClass: "block-c" }, [
      _c("div", { staticClass: "title-c" }, [
        _c("span", { staticClass: "title" }, [
          _vm._v(_vm._s(_vm.$t("modelManage.modelEngine")))
        ]),
        _vm._v(" "),
        _vm.engineFlag
          ? _c(
              "span",
              {
                staticClass: "clear-btn",
                on: {
                  click: function($event) {
                    return _vm.clearSelectLeft("engine")
                  }
                }
              },
              [_vm._v("Clear")]
            )
          : _vm._e()
      ]),
      _vm._v(" "),
      _c(
        "div",
        { staticClass: "list engine-c" },
        _vm._l(_vm.Engine, function(item, index) {
          return _c(
            "div",
            {
              key: item.k,
              staticClass: "item",
              class: item.active ? "active" : "",
              on: {
                click: function($event) {
                  return _vm.selectEngine(item)
                }
              }
            },
            [_vm._v(_vm._s(item.v))]
          )
        }),
        0
      )
    ]),
    _vm._v(" "),
    _c("div", { staticClass: "block-c" }, [
      _c("div", { staticClass: "title-c" }, [
        _c("span", { staticClass: "title" }, [
          _vm._v(_vm._s(_vm.$t("modelManage.label")))
        ]),
        _vm._v(" "),
        _vm.labelFlag
          ? _c(
              "span",
              {
                staticClass: "clear-btn",
                on: {
                  click: function($event) {
                    return _vm.clearSelectLeft("label")
                  }
                }
              },
              [_vm._v("Clear")]
            )
          : _vm._e()
      ]),
      _vm._v(" "),
      _c(
        "div",
        { staticClass: "list label-c" },
        _vm._l(_vm.Label, function(item, index) {
          return _c(
            "div",
            {
              key: item.k,
              staticClass: "item",
              class: item.active ? "active" : "",
              on: {
                click: function($event) {
                  return _vm.selectLabel(item)
                }
              }
            },
            [_vm._v(_vm._s(item.v))]
          )
        }),
        0
      )
    ])
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=template&id=6ad19374&scoped=true":
/*!*************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelItem.vue?vue&type=template&id=6ad19374&scoped=true ***!
  \*************************************************************************************************************************************************************************************************************************************************************/
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
    "a",
    {
      staticClass: "item",
      attrs: {
        target: "_blank",
        href:
          "/" +
          _vm.data.repoOwnerName +
          "/" +
          _vm.data.repoName +
          "/modelmanage/model_readme_tmpl?name=" +
          _vm.data.name
      }
    },
    [
      _c("div", { staticClass: "top" }, [
        _c("div", { staticClass: "top-head" }, [
          _vm._m(0),
          _vm._v(" "),
          _c("div", { staticClass: "name" }, [
            _c("span", { attrs: { title: _vm.data.name } }, [
              _vm._v(_vm._s(_vm.data.name))
            ])
          ]),
          _vm._v(" "),
          _vm.data.recommend == 1
            ? _c("div", { staticClass: "reconmend-icon" }, [
                _c("img", { attrs: { src: "/img/recommend.png", alt: "" } })
              ])
            : _vm._e()
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "top-bottom" }, [
          _c(
            "div",
            { staticClass: "labels" },
            [
              _vm.data.engineName || _vm.data.engine
                ? _c("span", { staticClass: "label engine" }, [
                    _vm._v(
                      " " +
                        _vm._s(_vm.data.engineName || _vm.data.engine) +
                        "\n        "
                    )
                  ])
                : _vm._e(),
              _vm._v(" "),
              _vm._l(_vm.data.labels, function(item, index) {
                return _c("a", { key: index, staticClass: "label normal" }, [
                  _vm._v(_vm._s(item))
                ])
              })
            ],
            2
          )
        ])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "footer" }, [
        _c("div", { staticClass: "footer-l" }, [
          _c(
            "a",
            {
              staticClass: "avatar-c",
              attrs: { href: "/" + _vm.data.userName }
            },
            [
              _c("img", {
                staticClass: "avatar",
                attrs: { src: _vm.data.userRelAvatarLink }
              })
            ]
          ),
          _vm._v(" "),
          _c(
            "span",
            { staticStyle: { "margin-left": "3px", "margin-right": "8px" } },
            [
              _vm._v(
                " " +
                  _vm._s(_vm.$t("repos.updated")) +
                  " " +
                  _vm._s(_vm.data.updateTimeStr) +
                  " "
              )
            ]
          ),
          _vm._v(" "),
          _vm.hasOnlineUrl
            ? _c("span", { staticStyle: { "white-space": "nowrap" } }, [
                _c("span", { staticClass: "greenPoint" }),
                _vm._v(
                  "\n        " +
                    _vm._s(_vm.$t("modelObj.can_online_infer")) +
                    "\n      "
                )
              ])
            : _vm._e()
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "footer-r" }, [
          _c(
            "div",
            {
              staticClass: "fav-c r-item",
              class: _vm.canChangeFav ? "" : "fav-disabled",
              on: {
                click: function($event) {
                  $event.preventDefault()
                  $event.stopPropagation()
                  return _vm.changeFav(_vm.data)
                }
              }
            },
            [
              !_vm.isCollected
                ? _c("i", {
                    staticClass: "heart outline icon",
                    attrs: {
                      title: _vm.canChangeFav
                        ? _vm.$t("star")
                        : _vm.$t("datasets.moststars")
                    }
                  })
                : _vm._e(),
              _vm._v(" "),
              _vm.isCollected
                ? _c("i", {
                    staticClass: "heart icon",
                    attrs: {
                      title: _vm.canChangeFav
                        ? _vm.$t("unStar")
                        : _vm.$t("datasets.moststars")
                    }
                  })
                : _vm._e(),
              _vm._v(" "),
              _c("span", [_vm._v(_vm._s(_vm.collectedCount))])
            ]
          ),
          _vm._v(" "),
          _c("div", { staticClass: "line" }),
          _vm._v(" "),
          _c(
            "span",
            {
              staticClass: "r-item",
              attrs: { title: _vm.$t("datasets.citations") }
            },
            [
              _c("i", { staticClass: "el-icon-link" }),
              _vm._v(" "),
              _c("span", [_vm._v(_vm._s(_vm.data.referenceCount))])
            ]
          ),
          _vm._v(" "),
          _c("div", { staticClass: "line" }),
          _vm._v(" "),
          _c(
            "span",
            {
              staticClass: "r-item",
              attrs: { title: _vm.$t("datasets.downloadtimes") }
            },
            [
              _c("i", { staticClass: "el-icon-download" }),
              _vm._v(" "),
              _c("span", [_vm._v(_vm._s(_vm.data.downloadCount))])
            ]
          ),
          _vm._v(" "),
          _c("div", { staticClass: "line" }),
          _vm._v(" "),
          _c(
            "span",
            {
              staticClass: "r-item",
              attrs: { title: _vm.$t("modelManage.derivativeTimes") }
            },
            [
              _c("i", { staticClass: "ri-git-merge-line" }),
              _vm._v(" "),
              _c("span", [_vm._v(_vm._s(_vm.data.derivativeCount))])
            ]
          )
        ])
      ])
    ]
  )
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "icon-c" }, [
      _c("img", { attrs: { src: "/img/icons/model-icon.png", alt: "" } })
    ])
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=template&id=74d70b7f&scoped=true":
/*!*************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelsquare/square/components/ModelList.vue?vue&type=template&id=74d70b7f&scoped=true ***!
  \*************************************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "list-container" }, [
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
        staticClass: "list-item-container"
      },
      [
        _vm._l(_vm.list, function(item, index) {
          return _c(
            "div",
            { key: item.id, staticClass: "item-container" },
            [
              _c("ModelItem", {
                key: item.id,
                attrs: { data: item, condition: _vm.params },
                on: { changeFav: _vm.changeFav }
              })
            ],
            1
          )
        }),
        _vm._v(" "),
        _c(
          "div",
          {
            directives: [
              {
                name: "show",
                rawName: "v-show",
                value: !_vm.list.length && !_vm.loading,
                expression: "(!list.length && !loading)"
              }
            ],
            staticClass: "no-data"
          },
          [
            _c("div", { staticClass: "item-empty" }, [
              _c("div", { staticClass: "item-empty-icon" }),
              _vm._v(" "),
              _c("div", { staticClass: "item-empty-tips" }, [
                _vm._v(_vm._s(_vm.$t("modelObj.model_square_empty")))
              ])
            ])
          ]
        )
      ],
      2
    ),
    _vm._v(" "),
    _c(
      "div",
      {
        directives: [
          {
            name: "show",
            rawName: "v-show",
            value: _vm.list.length,
            expression: "list.length"
          }
        ],
        staticClass: "center"
      },
      [
        _c("el-pagination", {
          ref: "paginationRef",
          attrs: {
            background: "",
            "current-page": _vm.iPage,
            "page-sizes": _vm.iPageSizes,
            "page-size": _vm.iPageSize,
            layout: "total, sizes, prev, pager, next, jumper",
            total: _vm.total
          },
          on: {
            "current-change": _vm.currentChange,
            "size-change": _vm.sizeChange,
            "update:currentPage": function($event) {
              _vm.iPage = $event
            },
            "update:current-page": function($event) {
              _vm.iPage = $event
            },
            "update:pageSize": function($event) {
              _vm.iPageSize = $event
            },
            "update:page-size": function($event) {
              _vm.iPageSize = $event
            }
          }
        })
      ],
      1
    )
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=template&id=6b484811&scoped=true":
/*!************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/repos/components/ReposItem.vue?vue&type=template&id=6b484811&scoped=true ***!
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
  return _c("div", [
    _c("div", { staticClass: "item" }, [
      _c("div", { staticClass: "item-top" }, [
        _vm.data.RelAvatarLink
          ? _c("img", {
              staticClass: "avatar",
              attrs: { src: _vm.data.RelAvatarLink }
            })
          : _c("img", {
              staticClass: "avatar",
              attrs: { avatar: _vm.data.Name }
            }),
        _vm._v(" "),
        _c("div", { staticClass: "content" }, [
          _c("div", { staticClass: "title" }, [
            _c("div", { staticClass: "title-l" }, [
              _c(
                "a",
                {
                  attrs: {
                    href: "/" + _vm.data.OwnerName + "/" + _vm.data.Name,
                    title: _vm.data.OwnerName + "/" + _vm.data.Name
                  }
                },
                [
                  _c("span", { staticClass: "title-1" }, [
                    _vm._v(_vm._s(_vm.data.OwnerName))
                  ]),
                  _vm._v(" "),
                  _c("span", { staticClass: "title-1" }, [_vm._v(" / ")]),
                  _vm._v(" "),
                  _c("span", {
                    staticClass: "title-2",
                    domProps: { innerHTML: _vm._s(_vm.data.NameShow) }
                  })
                ]
              ),
              _vm._v(" "),
              _vm.data.IsArchived
                ? _c("i", { staticClass: "archive icon archived-icon" })
                : _vm._e(),
              _vm._v(" "),
              _vm.data.IsFork
                ? _c(
                    "svg",
                    {
                      staticClass: "svg octicon-repo-forked",
                      attrs: {
                        width: "15",
                        height: "15",
                        "aria-hidden": "true"
                      }
                    },
                    [
                      _c("use", {
                        attrs: { "xlink:href": "#octicon-repo-forked" }
                      })
                    ]
                  )
                : _vm._e(),
              _vm._v(" "),
              _vm.data.IsMirror
                ? _c(
                    "svg",
                    {
                      staticClass: "svg octicon-repo-clone",
                      attrs: {
                        width: "15",
                        height: "15",
                        "aria-hidden": "true"
                      }
                    },
                    [
                      _c("use", {
                        attrs: { "xlink:href": "#octicon-repo-clone" }
                      })
                    ]
                  )
                : _vm._e(),
              _vm._v(" "),
              _vm.data.IsPrivate || _vm.data.IsOwnerPrivate
                ? _c(
                    "svg",
                    {
                      staticClass: "svg octicon-lock",
                      staticStyle: { color: "#a1882b!important" },
                      attrs: {
                        width: "15",
                        height: "15",
                        "aria-hidden": "true"
                      }
                    },
                    [_c("use", { attrs: { "xlink:href": "#octicon-lock" } })]
                  )
                : _vm._e()
            ]),
            _vm._v(" "),
            _c("span", { staticClass: "title-r only-mobile-hidden" }, [
              _c(
                "span",
                {
                  staticClass: "t-item",
                  attrs: { title: _vm.$t("repos.watch") }
                },
                [
                  _c("i", { staticClass: "ri-eye-line" }),
                  _vm._v(" "),
                  _c("span", [_vm._v(_vm._s(_vm.data.NumWatches))])
                ]
              ),
              _vm._v(" "),
              _c(
                "span",
                {
                  staticClass: "t-item",
                  attrs: { title: _vm.$t("repos.star") }
                },
                [
                  _c("i", { staticClass: "ri-star-line" }),
                  _vm._v(" "),
                  _c("span", [_vm._v(_vm._s(_vm.data.NumStars))])
                ]
              ),
              _vm._v(" "),
              _c(
                "span",
                {
                  staticClass: "t-item",
                  attrs: { title: _vm.$t("repos.fork") }
                },
                [
                  _c(
                    "svg",
                    {
                      staticClass: "svg octicon-repo-forked",
                      attrs: {
                        width: "13",
                        height: "13",
                        "aria-hidden": "true"
                      }
                    },
                    [
                      _c("use", {
                        attrs: { "xlink:href": "#octicon-repo-forked" }
                      })
                    ]
                  ),
                  _vm._v(" "),
                  _c("span", [_vm._v(_vm._s(_vm.data.NumForks))])
                ]
              )
            ])
          ]),
          _vm._v(" "),
          _c("div", {
            directives: [
              {
                name: "show",
                rawName: "v-show",
                value: _vm.data.DescriptionShow,
                expression: "data.DescriptionShow"
              }
            ],
            staticClass: "descr",
            domProps: { innerHTML: _vm._s(_vm.data.DescriptionShow) }
          }),
          _vm._v(" "),
          _c(
            "div",
            {
              directives: [
                {
                  name: "show",
                  rawName: "v-show",
                  value: _vm.data.Topics && _vm.data.Topics.length,
                  expression: "data.Topics && data.Topics.length"
                }
              ],
              staticClass: "tags",
              class: _vm.topicLink ? "" : "hide-link"
            },
            _vm._l(_vm.data.TopicsShow, function(item, index) {
              return _c("a", {
                key: index,
                staticClass: "tag",
                class:
                  item.topic.toLocaleLowerCase() ==
                  _vm.topic.toLocaleLowerCase()
                    ? "tag-focus"
                    : "",
                attrs: {
                  href: _vm.topicLink
                    ? "/explore/repos?q=&topic=" + item.topic + "&sort=hot"
                    : "javascript:;"
                },
                domProps: { innerHTML: _vm._s(item.topicShow) }
              })
            }),
            0
          )
        ])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "item-bottom" }, [
        _c(
          "div",
          [
            _c("span", [_vm._v(_vm._s(_vm.$t("repos.updated")))]),
            _vm._v(" "),
            _c(
              "el-tooltip",
              {
                attrs: {
                  effect: "dark",
                  content: _vm.dateFormat(_vm.data.UpdatedUnix),
                  placement: "top-start"
                }
              },
              [
                _c("span", [
                  _vm._v(_vm._s(_vm.calcFromNow(_vm.data.UpdatedUnix)))
                ])
              ]
            ),
            _vm._v(" "),
            _vm.data.PrimaryLanguage
              ? _c("span", { staticStyle: { "margin-left": "8px" } }, [
                  _c("i", {
                    staticClass: "color-icon",
                    style: { backgroundColor: _vm.data.PrimaryLanguage.Color }
                  }),
                  _vm._v(_vm._s(_vm.data.PrimaryLanguage.Language))
                ])
              : _vm._e()
          ],
          1
        ),
        _vm._v(" "),
        _c("div", { staticClass: "contributors" }, [
          _c(
            "span",
            {
              directives: [
                {
                  name: "show",
                  rawName: "v-show",
                  value: _vm.data.Contributors && _vm.data.Contributors.length,
                  expression: "data.Contributors && data.Contributors.length"
                }
              ],
              staticClass: "contributors-count"
            },
            [
              _vm._v(
                "\n          " +
                  _vm._s(_vm.$t("repos.contributors")) +
                  " \n        "
              )
            ]
          ),
          _vm._v(" "),
          _c(
            "span",
            { staticClass: "contributors-avatar" },
            _vm._l(_vm.data.Contributors, function(item, index) {
              return _c(
                "a",
                {
                  key: index,
                  staticClass: "avatar-c",
                  attrs: {
                    href: item.UserName
                      ? "/" + item.UserName
                      : item.Email
                      ? "mailto:" + item.Email
                      : "javascript:;"
                  }
                },
                [
                  _c("img", {
                    directives: [
                      {
                        name: "show",
                        rawName: "v-show",
                        value: item.UserName,
                        expression: "item.UserName"
                      }
                    ],
                    staticClass: "avatar",
                    attrs: { src: item.RelAvatarLink }
                  }),
                  _vm._v(" "),
                  _c(
                    "span",
                    {
                      directives: [
                        {
                          name: "show",
                          rawName: "v-show",
                          value: !item.UserName,
                          expression: "!item.UserName"
                        }
                      ],
                      staticClass: "avatar",
                      style: { backgroundColor: item.bgColor }
                    },
                    [
                      _vm._v(
                        "\n              " +
                          _vm._s((item.Email[0] || "").toLocaleUpperCase())
                      )
                    ]
                  )
                ]
              )
            }),
            0
          )
        ])
      ])
    ])
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=template&id=b942a2c8":
/*!**********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/App.vue?vue&type=template&id=b942a2c8 ***!
  \**********************************************************************************************************************************************************************************************************************/
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
  return _c("div", [_c("router-view")], 1)
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=template&id=cf1d6d6e&scoped=true":
/*!***************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/AppBanner.vue?vue&type=template&id=cf1d6d6e&scoped=true ***!
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
  return _c(
    "div",
    { staticClass: "group" },
    [_c("NavigationBar"), _vm._v(" "), _vm._m(0)],
    1
  )
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "info" }, [
      _c("span", { staticClass: "info1" }, [_vm._v("人机协同智能操作系统")]),
      _vm._v(" "),
      _c("div", { staticClass: "info-div" }, [
        _c("span", { staticClass: "info2" }, [
          _vm._v("智能操作软件库和工具  构建人工智能应用程序")
        ])
      ]),
      _vm._v(" "),
      _c("span", { staticClass: "info3" }, [_vm._v("ROS - hmci")])
    ])
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=template&id=42a186fb&scoped=true":
/*!*******************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/NavigationBar.vue?vue&type=template&id=42a186fb&scoped=true ***!
  \*******************************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "navbar" }, [
    _c(
      "div",
      { staticClass: "nanbar_cont justify-center flex-row items-center" },
      [
        _c(
          "div",
          {
            staticClass: "flex-row items-center nb_section",
            class: { "active-route": _vm.$route.name === "HomePage" },
            attrs: { id: "nav1" },
            on: {
              mouseover: function($event) {
                return _vm.addLineToImgSrc(1)
              },
              mouseout: function($event) {
                return _vm.removeLineFromImgSrc(1)
              }
            }
          },
          [
            _c("img", { staticClass: "nb_img", attrs: { src: _vm.imgSrc(1) } }),
            _vm._v(" "),
            _c(
              "a",
              { staticClass: "nb_txt", on: { click: _vm.gotoHomePage } },
              [_vm._v("首页")]
            )
          ]
        ),
        _vm._v(" "),
        _c("div", { staticClass: "line" }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-row items-center nb_section",
            class: { "active-route": _vm.$route.name === "OSSystem" },
            attrs: { id: "nav2" },
            on: {
              mouseover: function($event) {
                return _vm.addLineToImgSrc(2)
              },
              mouseout: function($event) {
                return _vm.removeLineFromImgSrc(2)
              }
            }
          },
          [
            _c("img", { staticClass: "nb_img", attrs: { src: _vm.imgSrc(2) } }),
            _vm._v(" "),
            _c(
              "a",
              { staticClass: "nb_txt", on: { click: _vm.gotoOSSystem } },
              [_vm._v("操作系统")]
            )
          ]
        ),
        _vm._v(" "),
        _c("div", { staticClass: "line" }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-row items-center nb_section",
            class: { "active-route": _vm.$route.name === "OpenApp" },
            attrs: { id: "nav3" },
            on: {
              mouseover: function($event) {
                return _vm.addLineToImgSrc(3)
              },
              mouseout: function($event) {
                return _vm.removeLineFromImgSrc(3)
              }
            }
          },
          [
            _c("img", { staticClass: "nb_img", attrs: { src: _vm.imgSrc(3) } }),
            _vm._v(" "),
            _c("a", { staticClass: "nb_txt", on: { click: _vm.gotoOpenApp } }, [
              _vm._v("开源应用")
            ])
          ]
        ),
        _vm._v(" "),
        _c("div", { staticClass: "line" }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-row items-center nb_section",
            class: { "active-route": _vm.$route.name === "OpenDataset" },
            attrs: { id: "nav4" },
            on: {
              mouseover: function($event) {
                return _vm.addLineToImgSrc(4)
              },
              mouseout: function($event) {
                return _vm.removeLineFromImgSrc(4)
              }
            }
          },
          [
            _c("img", { staticClass: "nb_img", attrs: { src: _vm.imgSrc(4) } }),
            _vm._v(" "),
            _c(
              "a",
              { staticClass: "nb_txt", on: { click: _vm.gotoOpenData } },
              [_vm._v("社区数据集")]
            )
          ]
        ),
        _vm._v(" "),
        _c("div", { staticClass: "line" }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-row items-center nb_section",
            class: { "active-route": _vm.$route.name === "OpenModel" },
            attrs: { id: "nav5" },
            on: {
              mouseover: function($event) {
                return _vm.addLineToImgSrc(5)
              },
              mouseout: function($event) {
                return _vm.removeLineFromImgSrc(5)
              }
            }
          },
          [
            _c("img", { staticClass: "nb_img", attrs: { src: _vm.imgSrc(5) } }),
            _vm._v(" "),
            _c(
              "a",
              { staticClass: "nb_txt", on: { click: _vm.gotoOpenModel } },
              [_vm._v("社区智能模型")]
            )
          ]
        ),
        _vm._v(" "),
        _c("div", { staticClass: "line" }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-row items-center nb_section",
            class: { "active-route": _vm.$route.name === "CommunitySource" },
            attrs: { id: "nav6" },
            on: {
              mouseover: function($event) {
                return _vm.addLineToImgSrc(6)
              },
              mouseout: function($event) {
                return _vm.removeLineFromImgSrc(6)
              }
            }
          },
          [
            _c("img", { staticClass: "nb_img", attrs: { src: _vm.imgSrc(6) } }),
            _vm._v(" "),
            _c(
              "a",
              { staticClass: "nb_txt", on: { click: _vm.gotoComunitySource } },
              [_vm._v("社区资源")]
            )
          ]
        ),
        _vm._v(" "),
        _c("div", { staticClass: "line" }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-row items-center nb_section",
            class: { "active-route": _vm.$route.name === "HelpCenter" },
            attrs: { id: "nav7" },
            on: {
              mouseover: function($event) {
                return _vm.addLineToImgSrc(7)
              },
              mouseout: function($event) {
                return _vm.removeLineFromImgSrc(7)
              }
            }
          },
          [
            _c("img", { staticClass: "nb_img", attrs: { src: _vm.imgSrc(7) } }),
            _vm._v(" "),
            _c(
              "a",
              { staticClass: "nb_txt", on: { click: _vm.gotoHelpCenter } },
              [_vm._v("帮助中心")]
            )
          ]
        )
      ]
    )
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=template&id=48f3bf6b&scoped=true":
/*!*******************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/PublicDataset.vue?vue&type=template&id=48f3bf6b&scoped=true ***!
  \*******************************************************************************************************************************************************************************************************************************************************/
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
      ]
    },
    [
      _c(
        "div",
        { staticClass: "dataset_head_wrap" },
        [
          _c(
            "el-checkbox",
            {
              staticStyle: { padding: "0.5rem 1rem" },
              on: { change: _vm.handleCheckedChange },
              model: {
                value: _vm.checked,
                callback: function($$v) {
                  _vm.checked = $$v
                },
                expression: "checked"
              }
            },
            [_vm._v(_vm._s(_vm.$t("datasets.platform_recommendations")))]
          ),
          _vm._v(" "),
          _c(
            "el-dropdown",
            { staticStyle: { cursor: "pointer" }, attrs: { trigger: "click" } },
            [
              _c("span", { staticClass: "el-dropdown-link" }, [
                _vm._v("\n        " + _vm._s(_vm.$t("datasets.sort"))),
                _c("i", { staticClass: "el-icon-caret-bottom el-icon--right" })
              ]),
              _vm._v(" "),
              _c(
                "el-dropdown-menu",
                { attrs: { slot: "dropdown" }, slot: "dropdown" },
                _vm._l(_vm.sortList, function(item) {
                  return _c(
                    "el-dropdown-item",
                    {
                      key: item.name,
                      class: { active: item.active },
                      nativeOn: {
                        click: function($event) {
                          return _vm.handleSort(item)
                        }
                      }
                    },
                    [_vm._v(_vm._s(_vm.$t("datasets." + item.name)))]
                  )
                }),
                1
              )
            ],
            1
          )
        ],
        1
      ),
      _vm._v(" "),
      !_vm.showEmpty
        ? _c(
            "div",
            { staticClass: "ui two cards" },
            _vm._l(_vm.publicDataList, function(item, index) {
              return _c(
                "div",
                {
                  staticClass: "ui card dataset_card_wrap",
                  on: {
                    click: function($event) {
                      return _vm.gotoDataset(item)
                    }
                  }
                },
                [
                  _c(
                    "div",
                    {
                      staticClass: "content",
                      staticStyle: { "border-bottom": "none" }
                    },
                    [
                      _c("div", { staticClass: "dataset_content_wrap" }, [
                        _c(
                          "span",
                          {
                            staticClass: "nowrap",
                            staticStyle: { display: "inline-block" },
                            attrs: { title: item.Title }
                          },
                          [_vm._v(_vm._s(item.Title))]
                        ),
                        _vm._v(" "),
                        item.Recommend
                          ? _c("img", {
                              staticStyle: { "margin-left": "0.5rem" },
                              attrs: { src: "/img/jian.svg" }
                            })
                          : _vm._e(),
                        _vm._v(" "),
                        _c(
                          "span",
                          {
                            staticClass: "dataset_icon_wrap",
                            on: {
                              click: function($event) {
                                $event.stopPropagation()
                                return _vm.postSquareStar(item, index)
                              }
                            }
                          },
                          [
                            _c("div", { staticClass: "dataset_icon_content" }, [
                              _c(
                                "svg",
                                {
                                  staticClass: "heart-stroke",
                                  class: { stars_active: item.IsStaring },
                                  attrs: {
                                    width: "1.4em",
                                    height: "1.4em",
                                    viewBox: "0 0 32 32"
                                  }
                                },
                                [
                                  _c("path", {
                                    attrs: {
                                      d:
                                        "M4.4 6.54c-1.761 1.643-2.6 3.793-2.36 6.056.24 2.263 1.507 4.521 3.663 6.534a29110.9 29110.9 0 0010.296 9.633l10.297-9.633c2.157-2.013 3.424-4.273 3.664-6.536.24-2.264-.599-4.412-2.36-6.056-1.73-1.613-3.84-2.29-6.097-1.955-1.689.25-3.454 1.078-5.105 2.394l-.4.319-.398-.319c-1.649-1.316-3.414-2.143-5.105-2.394a7.612 7.612 0 00-1.113-.081c-1.838 0-3.541.694-4.983 2.038z"
                                    }
                                  })
                                ]
                              )
                            ]),
                            _vm._v(" "),
                            _c(
                              "span",
                              {
                                staticStyle: {
                                  "line-height": "1",
                                  color: "#101010"
                                }
                              },
                              [_vm._v(_vm._s(item.NumStars))]
                            )
                          ]
                        )
                      ]),
                      _vm._v(" "),
                      _c("div", { staticClass: "dataset_label_content" }, [
                        item.Category
                          ? _c(
                              "span",
                              {
                                staticClass: "ui repo-topic label topic",
                                on: {
                                  click: function($event) {
                                    $event.stopPropagation()
                                    return _vm.chooseLabel(
                                      item.Category,
                                      "category"
                                    )
                                  }
                                }
                              },
                              [
                                _vm._v(
                                  _vm._s(_vm.$t("datasets." + item.Category))
                                )
                              ]
                            )
                          : _vm._e(),
                        _vm._v(" "),
                        item.Task
                          ? _c(
                              "span",
                              {
                                staticClass: "ui repo-topic label topic",
                                on: {
                                  click: function($event) {
                                    $event.stopPropagation()
                                    return _vm.chooseLabel(item.Task, "task")
                                  }
                                }
                              },
                              [_vm._v(_vm._s(_vm.$t("datasets." + item.Task)))]
                            )
                          : _vm._e(),
                        _vm._v(" "),
                        item.License
                          ? _c(
                              "span",
                              {
                                staticClass: "ui repo-topic label topic",
                                on: {
                                  click: function($event) {
                                    $event.stopPropagation()
                                    return _vm.chooseLabel(
                                      item.License,
                                      "license"
                                    )
                                  }
                                }
                              },
                              [_vm._v(_vm._s(item.License))]
                            )
                          : _vm._e()
                      ]),
                      _vm._v(" "),
                      _c(
                        "div",
                        { staticClass: "description dataset_desc_wrap" },
                        [_c("p", [_vm._v(_vm._s(item.Description))])]
                      )
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass: "extra content",
                      staticStyle: { "border-top": "none !important" }
                    },
                    [
                      _c(
                        "div",
                        {
                          staticStyle: {
                            display: "flex",
                            "align-items": "center"
                          }
                        },
                        [
                          item.UserID === 0
                            ? _c(
                                "a",
                                {
                                  attrs: {
                                    href: "/" + item.Repo.OwnerName,
                                    title: item.Repo.OwnerName
                                  }
                                },
                                [
                                  _c("img", {
                                    staticClass: "ui avatar image",
                                    staticStyle: {
                                      width: "22px",
                                      height: "22px"
                                    },
                                    attrs: {
                                      src:
                                        "/user/avatar/" +
                                        item.Repo.OwnerName +
                                        "/-1"
                                    }
                                  })
                                ]
                              )
                            : _c(
                                "a",
                                {
                                  attrs: {
                                    href: "/" + item.User.Name,
                                    title: item.User.Name
                                  }
                                },
                                [
                                  _c("img", {
                                    staticClass: "ui avatar image",
                                    staticStyle: {
                                      width: "22px",
                                      height: "22px"
                                    },
                                    attrs: {
                                      src:
                                        "/user/avatar/" + item.User.Name + "/-1"
                                    }
                                  })
                                ]
                              ),
                          _vm._v(" "),
                          _c("span", { staticClass: "dataset_extra_time" }, [
                            _vm._v(
                              _vm._s(_vm._f("DateTransfer")(item.CreatedUnix))
                            )
                          ]),
                          _vm._v(" "),
                          _c(
                            "span",
                            {
                              staticClass: "dataset_extra_link",
                              attrs: { title: _vm.$t("datasets.downloadtimes") }
                            },
                            [
                              _c("i", { staticClass: "ri-link" }),
                              _vm._v(" "),
                              _c(
                                "span",
                                { staticClass: "dataset_extra_content" },
                                [_vm._v(_vm._s(item.UseCount))]
                              )
                            ]
                          ),
                          _vm._v(" "),
                          _c(
                            "span",
                            {
                              staticClass: "dataset_extra_download",
                              attrs: { title: _vm.$t("datasets.citations") }
                            },
                            [
                              _c("i", { staticClass: "ri-download-line" }),
                              _vm._v(" "),
                              _c(
                                "span",
                                { staticClass: "dataset_extra_content" },
                                [_vm._v(_vm._s(item.DownloadTimes))]
                              )
                            ]
                          )
                        ]
                      )
                    ]
                  )
                ]
              )
            }),
            0
          )
        : _c("el-empty", {
            staticStyle: {
              "background-color": "rgba(245, 245, 246, 0.5)",
              "min-height": "400px",
              padding: "40px 0 80px 0"
            },
            attrs: { description: _vm.$t("noDataset"), "image-size": 100 }
          }),
      _vm._v(" "),
      !_vm.showEmpty
        ? _c(
            "div",
            { staticClass: "center", staticStyle: { "margin-top": "2rem" } },
            [
              _c("el-pagination", {
                attrs: {
                  background: "",
                  "current-page": _vm.params.page,
                  "page-sizes": [30],
                  "page-size": _vm.params.pageSize,
                  layout: "total, sizes, prev, pager, next, jumper",
                  total: _vm.total
                },
                on: {
                  "size-change": _vm.handleSizeChange,
                  "current-change": _vm.handleCurrentChange
                }
              })
            ],
            1
          )
        : _vm._e()
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=template&id=b0252fe4&scoped=true":
/*!******************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposFilters.vue?vue&type=template&id=b0252fe4&scoped=true ***!
  \******************************************************************************************************************************************************************************************************************************************************/
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
    _vm._l(_vm.list, function(item, index) {
      return _c(
        "div",
        {
          key: item.key,
          staticClass: "item",
          class: _vm.focusIndex == index ? "item-focus" : ""
        },
        [
          _c(
            "a",
            {
              attrs: { href: "javascript:;" },
              on: {
                click: function($event) {
                  return _vm.changeFilters(item, index)
                }
              }
            },
            [_vm._v(_vm._s(item.label))]
          )
        ]
      )
    }),
    0
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=template&id=885791ca&scoped=true":
/*!***************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/ReposList.vue?vue&type=template&id=885791ca&scoped=true ***!
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
  return _c("div", { staticClass: "list-container" }, [
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
        staticStyle: { "min-height": "540px" }
      },
      [
        _vm._l(_vm.list, function(item, index) {
          return _c(
            "div",
            { key: item.ID, staticClass: "repos-item-container" },
            [_c("ReposItem", { attrs: { data: item, topic: _vm.topic } })],
            1
          )
        }),
        _vm._v(" "),
        _c(
          "div",
          {
            directives: [
              {
                name: "show",
                rawName: "v-show",
                value: !_vm.list.length && !_vm.loading,
                expression: "!list.length && !loading"
              }
            ],
            staticClass: "repos-no-data"
          },
          [_vm._v("\n      " + _vm._s(_vm.$t("repos.noReposfound")) + "\n    ")]
        )
      ],
      2
    ),
    _vm._v(" "),
    _c(
      "div",
      { staticClass: "center" },
      [
        _c("el-pagination", {
          ref: "paginationRef",
          attrs: {
            background: "",
            "current-page": _vm.iPage,
            "page-sizes": _vm.iPageSizes,
            "page-size": _vm.iPageSize,
            layout: "total, sizes, prev, pager, next, jumper",
            total: _vm.total
          },
          on: {
            "current-change": _vm.currentChange,
            "size-change": _vm.sizeChange,
            "update:currentPage": function($event) {
              _vm.iPage = $event
            },
            "update:current-page": function($event) {
              _vm.iPage = $event
            },
            "update:pageSize": function($event) {
              _vm.iPageSize = $event
            },
            "update:page-size": function($event) {
              _vm.iPageSize = $event
            }
          }
        })
      ],
      1
    )
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=template&id=8f7ec232&scoped=true":
/*!***************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/SearchBar.vue?vue&type=template&id=8f7ec232&scoped=true ***!
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
  return _c("div", [
    _c("div", { staticClass: "_repo_search" }, [
      _c("div", { staticClass: "_repo_search_input_c" }, [
        _c("div", { staticClass: "_repo_search_input" }, [
          _c("input", {
            directives: [
              {
                name: "model",
                rawName: "v-model",
                value: _vm.searchInputValue,
                expression: "searchInputValue"
              }
            ],
            attrs: {
              type: "text",
              placeholder: _vm.$t("repos.searchRepositories"),
              autocomplete: "off"
            },
            domProps: { value: _vm.searchInputValue },
            on: {
              keyup: function($event) {
                if (
                  !$event.type.indexOf("key") &&
                  _vm._k($event.keyCode, "enter", 13, $event.key, "Enter")
                ) {
                  return null
                }
                return _vm.search($event)
              },
              input: function($event) {
                if ($event.target.composing) {
                  return
                }
                _vm.searchInputValue = $event.target.value
              }
            }
          })
        ]),
        _vm._v(" "),
        _c(
          "div",
          { staticClass: "_repo_search_btn", on: { click: _vm.search } },
          [
            _c(
              "svg",
              {
                staticClass:
                  "styles__StyledSVGIconPathComponent-sc-16fsqc8-0 kdvdTY svg-icon-path-icon fill",
                attrs: {
                  xmlns: "http://www.w3.org/2000/svg",
                  viewBox: "0 0 32 32",
                  width: "24",
                  height: "24"
                }
              },
              [
                _c(
                  "defs",
                  { attrs: { "data-reactroot": "" } },
                  [
                    _c(
                      "linearGradient",
                      {
                        attrs: {
                          id:
                            "ilac9fwnydq3lcx1,1,rs,1,f0dwf0xj,ezu9f0dw,f000,00lwrsktrs,rs1bhv8urs",
                          x1: "0",
                          x2: "100%",
                          y1: "0",
                          y2: "0",
                          gradientTransform:
                            "matrix(-0.7069999999999999, -0.707, 0.707, -0.7069999999999999, 16, 38.624)",
                          gradientUnits: "userSpaceOnUse"
                        }
                      },
                      [
                        _c("stop", {
                          attrs: {
                            "stop-color": "#c9ffbf",
                            "stop-opacity": "1",
                            offset: "0"
                          }
                        }),
                        _vm._v(" "),
                        _c("stop", {
                          attrs: {
                            "stop-color": "#0ca451",
                            "stop-opacity": "1",
                            offset: "1"
                          }
                        })
                      ],
                      1
                    )
                  ],
                  1
                ),
                _vm._v(" "),
                _c("g", [
                  _c("path", {
                    attrs: {
                      fill:
                        "url(#ilac9fwnydq3lcx1,1,rs,1,f0dwf0xj,ezu9f0dw,f000,00lwrsktrs,rs1bhv8urs)",
                      d:
                        "M14.667 2.667c6.624 0 12 5.376 12 12s-5.376 12-12 12-12-5.376-12-12 5.376-12 12-12zM14.667 24c5.156 0 9.333-4.177 9.333-9.333 0-5.157-4.177-9.333-9.333-9.333-5.157 0-9.333 4.176-9.333 9.333 0 5.156 4.176 9.333 9.333 9.333zM25.98 24.095l3.772 3.771-1.887 1.887-3.771-3.772 1.885-1.885z"
                    }
                  })
                ])
              ]
            ),
            _vm._v(" "),
            _c("span", { staticStyle: { "margin-left": "10px" } }, [
              _vm._v(_vm._s(_vm.$t("repos.search")))
            ])
          ]
        )
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "_repo_search_label_c" }, [
        _c(
          "div",
          { staticClass: "_repo_search_label" },
          [
            _vm._l(_vm.topics, function(item, index) {
              return _vm.type == "square"
                ? _c(
                    "a",
                    {
                      key: index,
                      style: {
                        backgroundColor:
                          _vm.topicColors[index % _vm.topicColors.length]
                      },
                      attrs: {
                        href:
                          "/explore/repos?q=" +
                          _vm.searchInputValue.trim() +
                          "&topic=" +
                          item.v +
                          "&sort=" +
                          _vm.sort
                      }
                    },
                    [_vm._v(_vm._s(item.v))]
                  )
                : _vm._e()
            }),
            _vm._v(" "),
            _vm.type == "search"
              ? _c(
                  "a",
                  {
                    staticStyle: { "font-weight": "bold" },
                    style: {
                      backgroundColor:
                        _vm.selectTopic == ""
                          ? _vm.selectedColor
                          : _vm.defaultColor,
                      color: _vm.selectTopic == "" ? "white" : "#40485b"
                    },
                    attrs: { href: "javascript:;" },
                    on: {
                      click: function($event) {
                        return _vm.changeTopic({ k: "", v: "" })
                      }
                    }
                  },
                  [_vm._v(_vm._s(_vm.$t("repos.allFields")))]
                )
              : _vm._e(),
            _vm._v(" "),
            _vm._l(_vm.topics, function(item, index) {
              return _vm.type == "search"
                ? _c(
                    "a",
                    {
                      key: index,
                      style: {
                        backgroundColor:
                          _vm.selectTopic.toLocaleLowerCase() == item.k
                            ? _vm.selectedColor
                            : _vm.defaultColor,
                        color:
                          _vm.selectTopic.toLocaleLowerCase() == item.k
                            ? "white"
                            : "#40485b"
                      },
                      attrs: { href: "javascript:;" },
                      on: {
                        click: function($event) {
                          return _vm.changeTopic(item)
                        }
                      }
                    },
                    [_vm._v(_vm._s(item.v))]
                  )
                : _vm._e()
            })
          ],
          2
        )
      ])
    ])
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=template&id=0230c3f3&scoped=true":
/*!*****************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/components/helpCollaps.vue?vue&type=template&id=0230c3f3&scoped=true ***!
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
  return _c(
    "div",
    { staticClass: "section_8" },
    [
      _c(
        "el-collapse",
        {
          on: { change: _vm.handleChange },
          model: {
            value: _vm.activeNames,
            callback: function($$v) {
              _vm.activeNames = $$v
            },
            expression: "activeNames"
          }
        },
        [
          _c("div", { staticClass: "coll_tag" }, [
            _c("span", { staticClass: "self-start text_9 font_3" }, [
              _vm._v("常见问题")
            ])
          ]),
          _vm._v(" "),
          _c(
            "el-collapse-item",
            {
              attrs: {
                title: "如何将项目加入到人机协同智能操作系统社区？",
                name: "1"
              }
            },
            [
              _c("div", { staticClass: "section_10" }, [
                _c("span", { staticClass: "font_5" }, [
                  _vm._v(
                    "给您的项目赋予标签ros-hmci-app(开源应用)，ros-hmci-datasets(社区数据集)，ros-hmci-models(社区智能模型)。"
                  )
                ])
              ])
            ]
          ),
          _vm._v(" "),
          _c(
            "el-collapse-item",
            { attrs: { title: "如何使用人机协同智能操作系统？", name: "5" } },
            [
              _c("div", { staticClass: "section_10" }, [
                _c("span", { staticClass: "font_5" }, [
                  _vm._v(
                    "在操作系统下载列表中选择您需要的操作系统下载至本地，解压zip的包，得到iso镜像文件后进行安装使用。"
                  )
                ])
              ])
            ]
          ),
          _vm._v(" "),
          _c(
            "el-collapse-item",
            { attrs: { title: "如何发布新的资源？", name: "6" } },
            [
              _c("div", { staticClass: "section_10" }, [
                _c("span", { staticClass: "font_5" }, [
                  _vm._v(
                    "在社区资源栏中点击资源发布，按格式填写需要发布的资源后提交，等待管理员审核完成后即可在社区资源找到您发布的资源。"
                  )
                ])
              ])
            ]
          ),
          _vm._v(" "),
          _c(
            "el-collapse-item",
            { attrs: { title: "是否能发布其他平台的资源？", name: "7" } },
            [
              _c("div", { staticClass: "section_10" }, [
                _c("span", { staticClass: "font_5" }, [_vm._v("可以。")])
              ])
            ]
          ),
          _vm._v(" "),
          _c("div", { staticClass: "coll_tag" }, [
            _c("span", { staticClass: "self-start font_3 text_15" }, [
              _vm._v("入门指南")
            ])
          ]),
          _vm._v(" "),
          _c(
            "el-collapse-item",
            { attrs: { title: "此处为常见问题2标题", name: "2" } },
            [
              _c("div", { staticClass: "section_10" }, [
                _c("span", { staticClass: "font_5" }, [
                  _vm._v("此处为常见问题2内容")
                ])
              ])
            ]
          ),
          _vm._v(" "),
          _c("div", { staticClass: "coll_tag" }, [
            _c("span", { staticClass: "self-start font_3 text_15" }, [
              _vm._v("安装手册")
            ])
          ]),
          _vm._v(" "),
          _c(
            "el-collapse-item",
            { attrs: { title: "此处为常见问题3标题", name: "3" } },
            [
              _c("div", { staticClass: "section_10" }, [
                _c("span", { staticClass: "font_5" }, [
                  _vm._v("此处为常见问题3内容")
                ])
              ])
            ]
          ),
          _vm._v(" "),
          _c("div", { staticClass: "coll_tag" }, [
            _c("span", { staticClass: "self-start font_3 text_15" }, [
              _vm._v("API接口")
            ])
          ]),
          _vm._v(" "),
          _c(
            "el-collapse-item",
            { attrs: { title: "此处为常见问题4标题", name: "4" } },
            [
              _c("div", { staticClass: "section_10" }, [
                _c("span", { staticClass: "font_5" }, [
                  _vm._v("此处为常见问题4内容")
                ])
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

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=template&id=6586b47e&scoped=true":
/*!****************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/CommunitySource.vue?vue&type=template&id=6586b47e&scoped=true ***!
  \****************************************************************************************************************************************************************************************************************************************************/
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
    { staticClass: "flex-col page" },
    [
      _c("AppBanner"),
      _vm._v(" "),
      _c("div", { staticClass: "flex-col justify-start relative group_3" }, [
        _c("div", { staticClass: "flex-col relative section_8" }, [
          _c("span", { staticClass: "self-start font_3" }, [
            _vm._v("ROS-hmci资源发布")
          ]),
          _vm._v(" "),
          _c(
            "div",
            { staticClass: "flex-col group_5 space-y-18" },
            _vm._l(_vm.paginatedResources, function(resource, index) {
              return _c(
                "div",
                { key: index, staticClass: "flex-col resource" },
                [
                  _c(
                    "div",
                    {
                      staticClass: "flex-row items-center self-start space-x-20"
                    },
                    [
                      _c("img", {
                        staticClass: "shrink-0 image",
                        attrs: {
                          src:
                            "/img/ros-hmci/9b12ac07bff17055334642687e7ea577.png"
                        }
                      }),
                      _vm._v(" "),
                      _c(
                        "router-link",
                        {
                          attrs: {
                            to: {
                              name: "ResourceDetail",
                              params: { name: resource.name }
                            }
                          }
                        },
                        [
                          _c("span", { staticClass: "name" }, [
                            _vm._v(_vm._s(resource.name))
                          ])
                        ]
                      )
                    ],
                    1
                  ),
                  _vm._v(" "),
                  _c("span", { staticClass: "self-start font_4 text_19" }, [
                    _vm._v(
                      "\n                        " +
                        _vm._s(resource.synopsis) +
                        "\n                    "
                    )
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "flex-col self-start group_6" }, [
                    _c("img", {
                      staticClass: "shrink-0 self-start image_14",
                      attrs: { src: "/img/ros-hmci/mbz627.png" }
                    }),
                    _vm._v(" "),
                    _c("span", { staticClass: "self-center font_2 text_20" }, [
                      _vm._v("提交日期：" + _vm._s(resource.create_time))
                    ])
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "divider" })
                ]
              )
            }),
            0
          ),
          _vm._v(" "),
          _c(
            "div",
            { staticClass: "center pagination" },
            [
              _c("el-pagination", {
                attrs: {
                  "current-page": _vm.currentPage,
                  "page-sizes": [15],
                  "page-size": 15,
                  layout: "total, sizes, prev, pager, next, jumper",
                  total: _vm.totalItems
                },
                on: {
                  "size-change": _vm.handleSizeChange,
                  "current-change": _vm.handleCurrentChange
                }
              })
            ],
            1
          )
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "section_7 pos_3" }, [
          _c(
            "div",
            {
              staticClass:
                "flex-row justify-between items-center section_7_content"
            },
            [
              _c("div"),
              _vm._v(" "),
              _c(
                "div",
                {
                  staticClass:
                    "flex-col justify-start items-start relative group_4"
                },
                [
                  _c("input", {
                    directives: [
                      {
                        name: "model",
                        rawName: "v-model",
                        value: _vm.searchQuery,
                        expression: "searchQuery"
                      }
                    ],
                    staticClass: "font_1 text_15",
                    attrs: {
                      "outline:none": "",
                      type: "text",
                      placeholder: "请输入资源名称以搜索"
                    },
                    domProps: { value: _vm.searchQuery },
                    on: {
                      input: function($event) {
                        if ($event.target.composing) {
                          return
                        }
                        _vm.searchQuery = $event.target.value
                      }
                    }
                  }),
                  _vm._v(" "),
                  _c(
                    "button",
                    { staticClass: "ui green button image_11 pos_4" },
                    [_vm._v(_vm._s(_vm.$t("repos.search")))]
                  )
                ]
              ),
              _vm._v(" "),
              _c("div", { staticClass: "flex-row space-x-28" }, [
                _c(
                  "div",
                  {
                    staticClass: "flex-row space-x-8",
                    attrs: { id: "openForm" }
                  },
                  [
                    _c("img", {
                      staticClass: "shrink-0 image_12",
                      attrs: {
                        src:
                          "/img/ros-hmci/4b4c4662629e074227c63d23f46271ac.png"
                      }
                    }),
                    _vm._v(" "),
                    _c(
                      "a",
                      {
                        staticClass: "font_1 text_16",
                        attrs: { href: _vm.links.source }
                      },
                      [_vm._v("资源发布")]
                    )
                  ]
                )
              ])
            ]
          )
        ])
      ])
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=template&id=13112268&scoped=true":
/*!***********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HelpCenter.vue?vue&type=template&id=13112268&scoped=true ***!
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
  return _c("div", { staticClass: "flex-col relative page" }, [
    _c(
      "div",
      { staticClass: "banner flex-col relative" },
      [_c("NavigationBar"), _vm._v(" "), _vm._m(0)],
      1
    ),
    _vm._v(" "),
    _c("img", {
      staticClass: "image_8 pos_5",
      attrs: { src: "/img/ros-hmci/group5727@2x.png" }
    }),
    _vm._v(" "),
    _c("div", { staticClass: "flex-col relative space-y-64" }, [
      _c("div", { staticClass: "self-center" }, [
        _c(
          "div",
          { staticClass: "flex-row group_3 space-x-14" },
          [_vm._m(1), _vm._v(" "), _c("helpCollaps")],
          1
        )
      ])
    ])
  ])
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "flex-col pos_6" }, [
      _c("span", { staticClass: "self-center text_6" }, [
        _vm._v("在这里,"),
        _c("br"),
        _vm._v("我们为你提供帮助与支持！")
      ])
    ])
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col shrink-0 self-start relative section_7" },
      [
        _c(
          "div",
          { staticClass: "flex-col justify-start items-start text-wrapper_2" },
          [_c("span", { staticClass: "font_2 text_8" }, [_vm._v("帮助中心")])]
        ),
        _vm._v(" "),
        _c("div", { staticClass: "self-center section_9" }),
        _vm._v(" "),
        _c("div", { staticClass: "flex-row group_6 space-x-8" }, [
          _c("img", {
            staticClass: "image_10",
            attrs: { src: "/img/ros-hmci/mbz640.png" }
          }),
          _vm._v(" "),
          _c("span", { staticClass: "font_4" }, [_vm._v("常见问题")])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-row group_6 space-x-8" }, [
          _c("img", {
            staticClass: "image_10",
            attrs: { src: "/img/ros-hmci/mbz637.png" }
          }),
          _vm._v(" "),
          _c("span", { staticClass: "font_4" }, [_vm._v("入门指南")])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-row group_6 space-x-8" }, [
          _c("img", {
            staticClass: "image_10",
            attrs: { src: "/img/ros-hmci/mbz636.png" }
          }),
          _vm._v(" "),
          _c("span", { staticClass: "font_4" }, [_vm._v("安装手册")])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-row group_6 space-x-8" }, [
          _c("img", {
            staticClass: "image_10",
            attrs: { src: "/img/ros-hmci/mbz635.png" }
          }),
          _vm._v(" "),
          _c("span", { staticClass: "font_4" }, [_vm._v("API接口")])
        ])
      ]
    )
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=template&id=c76be938&scoped=true":
/*!*********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/HomePage.vue?vue&type=template&id=c76be938&scoped=true ***!
  \*********************************************************************************************************************************************************************************************************************************************/
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
      staticClass: "flex-col relative page",
      staticStyle: { "padding-bottom": "0px" }
    },
    [
      _c("div", { staticClass: "flex-col section_6" }, [
        _c(
          "div",
          { staticClass: "section_6_cont" },
          [
            _c("div", { staticClass: "homeHeadCont" }, [
              _vm._m(0),
              _vm._v(" "),
              _vm._m(1),
              _vm._v(" "),
              _c("div", { staticClass: "homeBox3" }, [
                _vm._v("AI学习者的一体化智能学习交流平台")
              ]),
              _vm._v(" "),
              _c(
                "a",
                {
                  staticClass: "font_5 gotoDownload",
                  on: {
                    click: function($event) {
                      return _vm.getdownload("/os-system", "osDownload")
                    }
                  }
                },
                [_vm._v("立即下载")]
              )
            ]),
            _vm._v(" "),
            _vm._m(2),
            _vm._v(" "),
            _c("NavigationBar")
          ],
          1
        )
      ]),
      _vm._v(" "),
      _vm._m(3),
      _vm._v(" "),
      _c("div", [
        _c("div", { staticClass: "flex-row items-center group_14 pos_7" }, [
          _c(
            "div",
            {
              staticClass: "flex-col group_14 space-y-10",
              on: {
                mouseenter: function() {
                  return _vm.mouseEnter(0)
                }
              }
            },
            [
              _c("img", {
                staticClass: "image_18",
                class: { "bubble-animation": _vm.hoveredImage === 0 },
                attrs: { src: "/img/ros-hmci/作业环境.png" }
              }),
              _vm._v(" "),
              _c(
                "a",
                {
                  staticClass: "text_31",
                  style: { color: _vm.activeIndex === 0 ? "#24a19b" : "" }
                },
                [_vm._v("架构设计与系统集成")]
              )
            ]
          ),
          _vm._v(" "),
          _c(
            "div",
            {
              staticClass: "flex-col group_15 space-y-10",
              on: {
                mouseenter: function() {
                  return _vm.mouseEnter(1)
                }
              }
            },
            [
              _c("img", {
                staticClass: "self-center image_18",
                class: { "bubble-animation": _vm.hoveredImage === 1 },
                attrs: { src: "/img/ros-hmci/人机协同体系架构.png" }
              }),
              _vm._v(" "),
              _c(
                "a",
                {
                  staticClass: "text_31",
                  style: { color: _vm.activeIndex === 1 ? "#24a19b" : "" }
                },
                [_vm._v("作业规划器与示范应用")]
              )
            ]
          ),
          _vm._v(" "),
          _c(
            "div",
            {
              staticClass: "flex-col group_15 space-y-10",
              on: {
                mouseenter: function() {
                  return _vm.mouseEnter(2)
                }
              }
            },
            [
              _c("img", {
                staticClass: "image_18",
                class: { "bubble-animation": _vm.hoveredImage === 2 },
                attrs: { src: "/img/ros-hmci/持续自主学习.png" }
              }),
              _vm._v(" "),
              _c(
                "a",
                {
                  staticClass: "text_31",
                  style: { color: _vm.activeIndex === 2 ? "#24a19b" : "" }
                },
                [_vm._v("作业规划器与示范应用")]
              )
            ]
          ),
          _vm._v(" "),
          _c(
            "div",
            {
              staticClass: "flex-col group_17 space-y-10",
              on: {
                mouseenter: function() {
                  return _vm.mouseEnter(3)
                }
              }
            },
            [
              _c("img", {
                staticClass: "image_18",
                class: { "bubble-animation": _vm.hoveredImage === 3 },
                attrs: { src: "/img/ros-hmci/复杂环境建模.png" }
              }),
              _vm._v(" "),
              _c(
                "a",
                {
                  staticClass: "text_31",
                  style: { color: _vm.activeIndex === 3 ? "#24a19b" : "" }
                },
                [_vm._v("作业环境仿真器与数字孪生平台")]
              )
            ]
          ),
          _vm._v(" "),
          _c("div", { staticClass: "tabBar" })
        ]),
        _vm._v(" "),
        _vm._m(4),
        _vm._v(" "),
        _vm._m(5),
        _vm._v(" "),
        _vm._m(6),
        _vm._v(" "),
        _vm._m(7)
      ]),
      _vm._v(" "),
      _c(
        "div",
        { staticClass: "flex-col justify-start items-center pos_89 " },
        [
          _c("span", { staticClass: "font_8 pos_80" }, [_vm._v("项目协同")]),
          _vm._v(" "),
          _c(
            "div",
            { staticClass: "xmxt justify-center items-center flex-row" },
            [
              _c(
                "div",
                { staticClass: "justify-center items-center flex-col sjx" },
                [
                  _c("img", {
                    staticClass: "image_42",
                    attrs: { src: "/img/ros-hmci/mbz545.png" }
                  }),
                  _vm._v(" "),
                  _c("a", { staticClass: "text_73" }, [
                    _vm._v(_vm._s(_vm.systemComponentTotal))
                  ]),
                  _vm._v(" "),
                  _c("a", { staticClass: "text_74" }, [_vm._v("系统组件")])
                ]
              ),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "justify-center items-center flex-col sjx" },
                [
                  _c("img", {
                    staticClass: "image_42",
                    attrs: { src: "/img/ros-hmci/mbz546.png" }
                  }),
                  _vm._v(" "),
                  _c("a", { staticClass: "text_73" }, [
                    _vm._v(_vm._s(_vm.applicationSoftwareTotal))
                  ]),
                  _vm._v(" "),
                  _c("a", { staticClass: "text_74" }, [_vm._v("应用软件")])
                ]
              ),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "justify-center items-center flex-col sjx" },
                [
                  _c("img", {
                    staticClass: "image_42",
                    attrs: { src: "/img/ros-hmci/mbz547.png" }
                  }),
                  _vm._v(" "),
                  _c("a", { staticClass: "text_73" }, [
                    _vm._v(_vm._s(_vm.statisticsData.developer))
                  ]),
                  _vm._v(" "),
                  _c("a", { staticClass: "text_74" }, [_vm._v("开发者")])
                ]
              ),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "justify-center items-center flex-col sjx" },
                [
                  _c("img", {
                    staticClass: "image_42",
                    attrs: { src: "/img/ros-hmci/mbz548.png" }
                  }),
                  _vm._v(" "),
                  _c("a", { staticClass: "text_73" }, [
                    _vm._v(_vm._s(_vm.statisticsData.organization))
                  ]),
                  _vm._v(" "),
                  _c("a", { staticClass: "text_74" }, [_vm._v("组织")])
                ]
              )
            ]
          )
        ]
      ),
      _vm._v(" "),
      _c(
        "div",
        {
          staticClass: "flex-col justify-start items-center section_19 pos_13"
        },
        [
          _c("span", { staticClass: "pos_15" }, [
            _vm._v("人机协同智能操作系统整体架构")
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "Ovearc" }, [
            _c("div", { staticClass: "section_22 pos_19" }),
            _vm._v(" "),
            _c("div", { staticClass: "section_24 pos_25" }),
            _vm._v(" "),
            _c("div", { staticClass: "section_25 pos_28" }),
            _vm._v(" "),
            _c("div", { staticClass: "section_27 pos_35" }),
            _vm._v(" "),
            _c("img", {
              staticClass: "image_30 pos_24",
              attrs: {
                src: "/img/ros-hmci/39ea174f887d877ea76ba31228a6315d.png"
              }
            }),
            _vm._v(" "),
            _c("img", {
              staticClass: "image_30 pos_21",
              attrs: {
                src: "/img/ros-hmci/f21dd041539c46ca20943e97f003c425.png"
              }
            }),
            _vm._v(" "),
            _c("img", {
              staticClass: "image_32 pos_29",
              attrs: {
                src: "/img/ros-hmci/58c00cc05cb3a7614393125e1a0bd7ba.png"
              }
            }),
            _vm._v(" "),
            _c("img", {
              staticClass: "image_33 pos_32",
              attrs: {
                src: "/img/ros-hmci/f21dd041539c46ca20943e97f003c425.png"
              }
            }),
            _vm._v(" "),
            _c("img", {
              staticClass: "image_30 pos_31",
              attrs: {
                src: "/img/ros-hmci/c283a7269004ff2edbff7619d4346ae1.png"
              }
            }),
            _vm._v(" "),
            _c("img", {
              staticClass: "image_30 pos_34",
              attrs: {
                src: "/img/ros-hmci/c283a7269004ff2edbff7619d4346ae1.png"
              }
            }),
            _vm._v(" "),
            _c("img", {
              staticClass: "image_30 pos_39",
              attrs: {
                src: "/img/ros-hmci/c283a7269004ff2edbff7619d4346ae1.png"
              }
            }),
            _vm._v(" "),
            _c("img", {
              staticClass: "image_34 pos_37",
              attrs: {
                src: "/img/ros-hmci/3d0cdf6c59b3734bdcaa492f430697ae.png"
              }
            }),
            _vm._v(" "),
            _vm._m(8),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col items-center space-y-2 pos_91" },
              [
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start items-center section_23 pos_jgt",
                    on: {
                      mouseenter: function() {
                        return _vm.xtjg(1)
                      },
                      mouseleave: function($event) {
                        return _vm.unxtjg(1)
                      }
                    }
                  },
                  [
                    _c("a", { staticClass: "font_15 text_43" }, [
                      _vm._v("作业")
                    ]),
                    _vm._v(" "),
                    _c("a", { staticClass: "font_15 text_43" }, [
                      _vm._v("环境模型")
                    ])
                  ]
                ),
                _vm._v(" "),
                _vm._m(9)
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col items-center space-y-2 pos_92" },
              [
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start items-center text-wrapper_6 pos_jgt",
                    on: {
                      mouseenter: function() {
                        return _vm.xtjg(2)
                      },
                      mouseleave: function($event) {
                        return _vm.unxtjg(2)
                      }
                    }
                  },
                  [
                    _c("a", { staticClass: "font_15 text_43" }, [
                      _vm._v("感知层")
                    ])
                  ]
                ),
                _vm._v(" "),
                _vm._m(10)
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col items-center space-y-2 pos_93" },
              [
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col items-center section_23 space-y-2 pos_jgt",
                    on: {
                      mouseenter: function() {
                        return _vm.xtjg(3)
                      },
                      mouseleave: function($event) {
                        return _vm.unxtjg(3)
                      }
                    }
                  },
                  [
                    _c("a", { staticClass: "font_14" }, [_vm._v("无人平台")]),
                    _vm._v(" "),
                    _c("a", { staticClass: "font_14" }, [_vm._v("硬件虚拟层")])
                  ]
                ),
                _vm._v(" "),
                _vm._m(11)
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col items-center space-y-2 pos_94" },
              [
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col items-center section_23 space-y-2 pos_jgt",
                    on: {
                      mouseenter: function() {
                        return _vm.xtjg(4)
                      },
                      mouseleave: function($event) {
                        return _vm.unxtjg(4)
                      }
                    }
                  },
                  [
                    _c("a", { staticClass: "font_13" }, [_vm._v("虚拟空间")]),
                    _vm._v(" "),
                    _c("a", { staticClass: "font_13" }, [
                      _vm._v("自主学习平台")
                    ])
                  ]
                ),
                _vm._v(" "),
                _vm._m(12)
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col items-center space-y-2 pos_95" },
              [
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start items-center text-wrapper_6 pos_jgt",
                    on: {
                      mouseenter: function() {
                        return _vm.xtjg(5)
                      },
                      mouseleave: function($event) {
                        return _vm.unxtjg(5)
                      }
                    }
                  },
                  [
                    _c("a", { staticClass: "font_5 text_46" }, [
                      _vm._v("人机协同层")
                    ])
                  ]
                ),
                _vm._v(" "),
                _vm._m(13)
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col items-center space-y-2 pos_96" },
              [
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col items-center section_29 space-y-2 pos_jgt",
                    on: {
                      mouseenter: function() {
                        return _vm.xtjg(6)
                      },
                      mouseleave: function($event) {
                        return _vm.unxtjg(6)
                      }
                    }
                  },
                  [
                    _c("a", { staticClass: "font_13" }, [_vm._v("规划")]),
                    _vm._v(" "),
                    _c("a", { staticClass: "font_13" }, [_vm._v("推理层")])
                  ]
                ),
                _vm._v(" "),
                _vm._m(14)
              ]
            ),
            _vm._v(" "),
            _vm._m(15),
            _vm._v(" "),
            _vm._m(16),
            _vm._v(" "),
            _vm._m(17)
          ])
        ]
      ),
      _vm._v(" "),
      _c("div", { staticClass: "pos_88" }, [
        _c("span", { staticClass: "font_8 pos_44" }, [
          _vm._v("系统典型应用场景")
        ]),
        _vm._v(" "),
        _c("span", { staticClass: "font_5 text_50 pos_45" }),
        _vm._v(" "),
        _vm.typical.length >= 3
          ? _c(
              "div",
              { staticClass: "flex-row justify-center section_30 space-x-42" },
              [
                _c(
                  "div",
                  { staticClass: "flex-col justify-start relative group_26" },
                  [
                    _c("div", { staticClass: "flex-col relative section_35" }, [
                      _c(
                        "span",
                        { staticClass: "self-start font_16 text_51" },
                        [_vm._v(_vm._s(_vm.typical[0].name))]
                      ),
                      _vm._v(" "),
                      _c(
                        "span",
                        { staticClass: "self-start font_17 text_54" },
                        [_vm._v(_vm._s(_vm.typical[0].developers))]
                      ),
                      _vm._v(" "),
                      _c(
                        "div",
                        { staticClass: "flex-col items-start group_27" },
                        [
                          _c("span", { staticClass: "font_12 text_55" }, [
                            _vm._v("“")
                          ]),
                          _vm._v(" "),
                          _c(
                            "span",
                            {
                              staticClass: "font_11 text_56 double-line",
                              attrs: { title: _vm.typical[0].desc }
                            },
                            [
                              _vm._v(
                                "\n              " +
                                  _vm._s(_vm.typical[0].desc) +
                                  "\n            "
                              )
                            ]
                          )
                        ]
                      ),
                      _vm._v(" "),
                      _c(
                        "div",
                        { staticClass: "flex-row group_28 space-x-4" },
                        [
                          _c(
                            "a",
                            {
                              staticClass: "font_18",
                              attrs: { href: _vm.typical[0].url }
                            },
                            [_vm._v("查看详情")]
                          ),
                          _vm._v(" "),
                          _c("img", {
                            staticClass: "image_38",
                            attrs: {
                              src:
                                "/img/ros-hmci/5520ba673bc80210912d830a304adddd.png"
                            }
                          })
                        ]
                      )
                    ]),
                    _vm._v(" "),
                    _c("img", {
                      staticClass: "image_35 pos_49",
                      attrs: {
                        src:
                          "/img/ros-hmci/cd2dbfa8c0212c06144c86f45ad70df8.png"
                      }
                    }),
                    _vm._v(" "),
                    _vm._m(18)
                  ]
                ),
                _vm._v(" "),
                _c(
                  "div",
                  { staticClass: "flex-col justify-start relative group_26" },
                  [
                    _c("div", { staticClass: "flex-col relative section_35" }, [
                      _c(
                        "span",
                        { staticClass: "self-start font_16 text_52" },
                        [_vm._v(_vm._s(_vm.typical[1].name))]
                      ),
                      _vm._v(" "),
                      _c(
                        "span",
                        { staticClass: "self-start font_17 text_54" },
                        [_vm._v(_vm._s(_vm.typical[1].developers))]
                      ),
                      _vm._v(" "),
                      _c(
                        "div",
                        { staticClass: "flex-col items-start group_27" },
                        [
                          _c("span", { staticClass: "font_12 text_55" }, [
                            _vm._v("“")
                          ]),
                          _vm._v(" "),
                          _c(
                            "span",
                            {
                              staticClass: "font_11 text_56 double-line",
                              attrs: { title: _vm.typical[1].desc }
                            },
                            [
                              _vm._v(
                                "\n              " +
                                  _vm._s(_vm.typical[1].desc) +
                                  "\n            "
                              )
                            ]
                          )
                        ]
                      ),
                      _vm._v(" "),
                      _c(
                        "div",
                        { staticClass: "flex-row group_28 space-x-4" },
                        [
                          _c(
                            "a",
                            {
                              staticClass: "font_18",
                              attrs: { href: _vm.typical[1].url }
                            },
                            [_vm._v("查看详情")]
                          ),
                          _vm._v(" "),
                          _c("img", {
                            staticClass: "image_38",
                            attrs: {
                              src:
                                "/img/ros-hmci/5520ba673bc80210912d830a304adddd.png"
                            }
                          })
                        ]
                      )
                    ]),
                    _vm._v(" "),
                    _c("img", {
                      staticClass: "image_36 pos_50",
                      attrs: {
                        src:
                          "/img/ros-hmci/7bd4e4622f6f137f0c4f348339b275c5.png"
                      }
                    }),
                    _vm._v(" "),
                    _c("img", {
                      staticClass: "image_37 pos_52",
                      attrs: {
                        src:
                          "/img/ros-hmci/6f6bf64d608d346bd40464912be1df76.png"
                      }
                    }),
                    _vm._v(" "),
                    _vm._m(19)
                  ]
                ),
                _vm._v(" "),
                _c(
                  "div",
                  { staticClass: "flex-col justify-start relative group_26" },
                  [
                    _c("div", { staticClass: "flex-col relative section_35" }, [
                      _c(
                        "span",
                        { staticClass: "self-start font_16 text_53" },
                        [_vm._v(_vm._s(_vm.typical[2].name))]
                      ),
                      _vm._v(" "),
                      _c(
                        "span",
                        { staticClass: "self-start font_17 text_54" },
                        [_vm._v(_vm._s(_vm.typical[2].developers))]
                      ),
                      _vm._v(" "),
                      _c(
                        "div",
                        { staticClass: "flex-col items-start group_27" },
                        [
                          _c("span", { staticClass: "font_12 text_55" }, [
                            _vm._v("“")
                          ]),
                          _vm._v(" "),
                          _c(
                            "span",
                            {
                              staticClass: "font_11 text_56 double-line",
                              attrs: { title: _vm.typical[2].desc }
                            },
                            [
                              _vm._v(
                                "\n              " +
                                  _vm._s(_vm.typical[2].desc) +
                                  "\n            "
                              )
                            ]
                          )
                        ]
                      ),
                      _vm._v(" "),
                      _c(
                        "div",
                        { staticClass: "flex-row group_28 space-x-4" },
                        [
                          _c(
                            "a",
                            {
                              staticClass: "font_18",
                              attrs: { href: _vm.typical[2].url }
                            },
                            [_vm._v("查看详情")]
                          ),
                          _vm._v(" "),
                          _c("img", {
                            staticClass: "image_38",
                            attrs: {
                              src:
                                "/img/ros-hmci/5520ba673bc80210912d830a304adddd.png"
                            }
                          })
                        ]
                      )
                    ]),
                    _vm._v(" "),
                    _c("img", {
                      staticClass: "image_36 pos_51",
                      attrs: {
                        src:
                          "/img/ros-hmci/7bd4e4622f6f137f0c4f348339b275c5.png"
                      }
                    }),
                    _vm._v(" "),
                    _c("img", {
                      staticClass: "image_37 pos_53",
                      attrs: {
                        src:
                          "/img/ros-hmci/6f6bf64d608d346bd40464912be1df76.png"
                      }
                    }),
                    _vm._v(" "),
                    _vm._m(20)
                  ]
                )
              ]
            )
          : _vm._e(),
        _vm._v(" "),
        _c("div", [
          _c("div", { staticClass: "pos_86", attrs: { id: "cbl1" } }, [
            _c(
              "div",
              {
                directives: [
                  {
                    name: "show",
                    rawName: "v-show",
                    value: _vm.showText[0],
                    expression: "showText[0]"
                  }
                ],
                staticClass: "cbltext"
              },
              [_c("span", { staticClass: "text_96" }, [_vm._v("意见反馈")])]
            ),
            _vm._v(" "),
            _c("a", { attrs: { href: _vm.links.home_advice } }, [
              _c("img", {
                staticClass: "cblimg",
                attrs: { src: _vm.feedbackImgSrc },
                on: {
                  mouseenter: function($event) {
                    return _vm.handleMouseEnter(1)
                  },
                  mouseleave: function($event) {
                    return _vm.handleMouseLeave(1)
                  }
                }
              })
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "pos_85", attrs: { id: "cbl2" } }, [
            _c(
              "div",
              {
                directives: [
                  {
                    name: "show",
                    rawName: "v-show",
                    value: _vm.showText[1],
                    expression: "showText[1]"
                  }
                ],
                staticClass: "cbltext",
                staticStyle: { position: "relative", top: "0px", left: "-30px" }
              },
              [_vm._m(21)]
            ),
            _vm._v(" "),
            _c("img", {
              staticClass: "cblimg",
              staticStyle: {
                position: "absolute",
                right: "-15px",
                top: "-12px"
              },
              attrs: { src: _vm.qrcodeImgSrc },
              on: {
                mouseenter: function($event) {
                  return _vm.handleMouseEnter(2)
                },
                mouseleave: function($event) {
                  return _vm.handleMouseLeave(2)
                }
              }
            })
          ])
        ])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "flex-col justify-start section_36 pos_54" }, [
        _c(
          "div",
          { staticClass: "flex-col items-center section_37 space-y-50" },
          [
            _c("span", { staticClass: "font_8 text_58" }, [
              _vm._v("现在开始！进入人机协同智能操作系统社区。")
            ]),
            _vm._v(" "),
            _c(
              "div",
              {
                staticClass:
                  "flex-col justify-start items-center text-wrapper_10",
                on: {
                  click: function($event) {
                    return _vm.getdownload("/os-system", "osDownload")
                  }
                }
              },
              [
                _c("a", { staticClass: "font_10 text_59" }, [
                  _vm._v("立即下载")
                ])
              ]
            )
          ]
        )
      ])
    ]
  )
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "flexCont" }, [
      _c("span", [_vm._v("欢迎来到")]),
      _vm._v(" "),
      _c("span", { staticClass: "homeBox1" }, [_vm._v("智能学习")])
    ])
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", [
      _c("span", { staticClass: "homeBox2" }, [_vm._v("人机协同")]),
      _vm._v(" "),
      _c("span", [_vm._v("智能操作系统")])
    ])
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "banner_right" }, [
      _c("img", {
        staticClass: "homepage1",
        attrs: { src: "/img/ros-hmci/homepage1.png" }
      }),
      _vm._v(" "),
      _c("img", {
        staticClass: "homepage2",
        attrs: { src: "/img/ros-hmci/homepage2.png" }
      }),
      _vm._v(" "),
      _c("img", {
        staticClass: "homepage3",
        attrs: { src: "/img/ros-hmci/homepage3.png" }
      }),
      _vm._v(" "),
      _c("img", {
        staticClass: "homepage4",
        attrs: { src: "/img/ros-hmci/homepage4.png" }
      })
    ])
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col justify-start items-center section_10 pos_3" },
      [
        _c("div", [
          _c("span", { staticClass: "font_8 pos_80" }, [
            _vm._v("人机协同智能操作系统简介")
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-row section_11 space-x-40 pos_5" }, [
          _c("div", { staticClass: "flex-auto self-start group_13 " }, [
            _c("span", { staticClass: "font_9 text_28" }, [_vm._v("”")]),
            _vm._v(" "),
            _c("span", { staticClass: "text_27" }, [
              _vm._v(
                "\n          瞄准人机协同自主作业需求，针对“可持续自主学习型智能操作系统”和“人机物融合、多主体协同计算平台”两大科学问题，着眼构建新一代人工智能基础软件，创新和突破相关基础理论和关键技术，研制人机协同智能操作系统原型版本，构建支持操作系统开源生态构建和智能资源共享的开源开放基础平台，面向典型场景和任务开展应用验证。\n        "
              )
            ])
          ]),
          _vm._v(" "),
          _c("img", {
            staticClass: "shrink-0 self-center image_17",
            attrs: { src: "/img/ros-hmci/mbz549.png" }
          })
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "flex-col justify-start items-end group_20 pos_9",
        staticStyle: { display: "''" },
        attrs: { id: "tab1" }
      },
      [
        _c("div", { staticClass: "flex-col relative section_17" }, [
          _c("div", { staticClass: "flex-row group_21" }, [
            _c("span", { staticClass: "self-center font_11 text_36" }, [
              _vm._v(
                "\n            针对复杂应用环境，开展人机自主协同架构设计，突破无人平台硬件虚拟层设计、适应环境的感知/规划/推理、人机自主协同操控、面向作业的实时无线通信等关键技术，构建人机自主协同操控子系统，研制支持多种异构智能硬件/软件/数据资源和人机互理解的智能操作系统原型。\n          "
              )
            ]),
            _vm._v(" "),
            _c("span", { staticClass: "self-start font_9 text_35" }, [
              _vm._v("”")
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col items-start relative group_22" }, [
            _c(
              "div",
              {
                staticClass:
                  "flex-col justify-start items-center text-wrapper_3"
              },
              [
                _c("span", { staticClass: "font_5 text_37" }, [
                  _vm._v("面向开源")
                ])
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              {
                staticClass:
                  "flex-col justify-start items-center group_23 pos_11"
              },
              [
                _c("img", {
                  staticClass: "image_26",
                  attrs: { src: "/img/ros-hmci/白色背景.png" }
                }),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start items-center text-wrapper_4 pos_12"
                  },
                  [
                    _c("span", { staticClass: "text_38" }, [
                      _vm._v("质量保证机制")
                    ])
                  ]
                )
              ]
            )
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col justify-start section_15 pos_10" }, [
          _c("div", { staticClass: "flex-col tab_sec1 space-y-134" })
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "flex-col justify-start items-end group_20 pos_9",
        staticStyle: { display: "none" },
        attrs: { id: "tab2" }
      },
      [
        _c("div", { staticClass: "flex-col relative section_17" }, [
          _c("div", { staticClass: "flex-row group_21" }, [
            _c("span", { staticClass: "self-center font_11 text_36" }, [
              _vm._v(
                "\n            针对复杂环境下无人系统作业的机器可理解可执行需求，聚焦无人系统作业规划与作业调度问题，突破作业环境模型构建、协同作业规划、作业状态切换引擎、面向作业的人机互理解等关键技术，设计实现适应复杂环境和支持复杂作业的作业规划器子系统，开展面向反恐侦察和农田作业的人机协同智能操作系统典型应用示范。\n          "
              )
            ]),
            _vm._v(" "),
            _c("span", { staticClass: "self-start font_9 text_35" }, [
              _vm._v("”")
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col items-start relative group_22" }, [
            _c(
              "div",
              {
                staticClass:
                  "flex-col justify-start items-center text-wrapper_3"
              },
              [
                _c("span", { staticClass: "font_5 text_37" }, [
                  _vm._v("面向开源")
                ])
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              {
                staticClass:
                  "flex-col justify-start items-center group_23 pos_11"
              },
              [
                _c("img", {
                  staticClass: "image_26",
                  attrs: { src: "/img/ros-hmci/白色背景.png" }
                }),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start items-center text-wrapper_4 pos_12"
                  },
                  [
                    _c("span", { staticClass: "text_38" }, [
                      _vm._v("质量保证机制")
                    ])
                  ]
                )
              ]
            )
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col justify-start section_15 pos_10" }, [
          _c("div", { staticClass: "flex-col tab_sec2 space-y-134" })
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "flex-col justify-start items-end group_20 pos_9",
        staticStyle: { display: "none" },
        attrs: { id: "tab3" }
      },
      [
        _c("div", { staticClass: "flex-col relative section_17" }, [
          _c("div", { staticClass: "flex-row group_21" }, [
            _c("span", { staticClass: "self-center font_11 text_36" }, [
              _vm._v(
                "\n            针对人机协同智能系统快速赋能和持续增能的需求，突破智能体依托虚拟环境自主学习行为策略的核心机理，构造支撑虚拟空间内自主持续学习和利用人类智能数据进行学习的虚拟学习引擎；突破操作系统开源生态构建和演化的核心使能机理，以及运行时知识、模型、算法等智能资源的规模化共享机制，搭建面向人机协同智能操作系统软件生态链构造和运行时后端支撑的开源开放基础平台。\n          "
              )
            ]),
            _vm._v(" "),
            _c("span", { staticClass: "self-start font_9 text_35" }, [
              _vm._v("”")
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col items-start relative group_22" }, [
            _c(
              "div",
              {
                staticClass:
                  "flex-col justify-start items-center text-wrapper_3"
              },
              [
                _c("span", { staticClass: "font_5 text_37" }, [
                  _vm._v("面向开源")
                ])
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              {
                staticClass:
                  "flex-col justify-start items-center group_23 pos_11"
              },
              [
                _c("img", {
                  staticClass: "image_26",
                  attrs: { src: "/img/ros-hmci/白色背景.png" }
                }),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start items-center text-wrapper_4 pos_12"
                  },
                  [
                    _c("span", { staticClass: "text_38" }, [
                      _vm._v("质量保证机制")
                    ])
                  ]
                )
              ]
            )
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col justify-start section_15 pos_10" }, [
          _c("div", { staticClass: "flex-col tab_sec3 space-y-134" })
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "flex-col justify-start items-end group_20 pos_9",
        staticStyle: { display: "none" },
        attrs: { id: "tab4" }
      },
      [
        _c("div", { staticClass: "flex-col relative section_17" }, [
          _c("div", { staticClass: "flex-row group_21" }, [
            _c("span", { staticClass: "self-center font_11 text_36" }, [
              _vm._v(
                "\n            针对虚拟空间自主学习对作业环境仿真的需求与挑战，开展跨传统计算机操作系统的并行计算架构设计，研究人机协同仿真技术，攻克复杂作业环境智能建模、视景渲染、物理过程仿真技术，构建具有“高可扩展、高逼真、高保真、高实时”的作业环境仿真器与数字孪生平台，为可持续自主学习提供大数据支撑，以课题1和课题2的应用需求为牵引开展数字孪生平台的验证和优化。\n          "
              )
            ]),
            _vm._v(" "),
            _c("span", { staticClass: "self-start font_9 text_35" }, [
              _vm._v("”")
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col items-start relative group_22" }, [
            _c(
              "div",
              {
                staticClass:
                  "flex-col justify-start items-center text-wrapper_3"
              },
              [
                _c("span", { staticClass: "font_5 text_37" }, [
                  _vm._v("面向开源")
                ])
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              {
                staticClass:
                  "flex-col justify-start items-center group_23 pos_11"
              },
              [
                _c("img", {
                  staticClass: "image_26",
                  attrs: { src: "/img/ros-hmci/白色背景.png" }
                }),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start items-center text-wrapper_4 pos_12"
                  },
                  [
                    _c("span", { staticClass: "text_38" }, [
                      _vm._v("质量保证机制")
                    ])
                  ]
                )
              ]
            )
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col justify-start section_15 pos_10" }, [
          _c("div", { staticClass: "flex-col tab_sec4 space-y-134" })
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "flex-col section_26 space-y-14 pos_33" }, [
      _c("div", { staticClass: "flex-row items-center group_25 space-x-12" }, [
        _c("img", {
          staticClass: "image",
          attrs: { src: "/img/ros-hmci/81de95962288a1aa874c40ee5f622d90.png" }
        }),
        _vm._v(" "),
        _c("span", { staticClass: "font_5 text_47" }, [
          _vm._v("人机协同智能操作系统")
        ])
      ]),
      _vm._v(" "),
      _c("span", { staticClass: "text_48" }, [
        _vm._v(
          "\n          人机协同智能操作系统是元操作系统（Meta\n          OS），它运行于传统的计算机操作系统之上，实现智能无人平台的传感器、执行器、通信系统等硬件资源的虚化，管理智能无人系统的感知、规划、推理等自主行为，并对人机自主协同和持续自主学习提供支撑，研究人机协同智能操作系统架构层次性功能设计、协同机制以及接口标准，并持续迭代进行设计优化。\n        "
        )
      ])
    ])
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticStyle: { display: "none" }, attrs: { id: "jgt1" } },
      [
        _c("img", {
          staticClass: "image_31 pos_22",
          staticStyle: {
            position: "relative",
            left: "-178px",
            top: "129px",
            transform: "rotate(180deg)"
          },
          attrs: { src: "/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" }
        }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-col justify-start items-start section_20",
            staticStyle: { position: "relative", top: "5px", right: "470px" }
          },
          [
            _c("span", { staticClass: "font_12 text_42" }, [_vm._v("“")]),
            _vm._v(" "),
            _c("span", { staticClass: "font_3 text_41 pos_17" }, [
              _vm._v(
                "\n              作业环境模型中记录了机器和算法可理解的地图和障碍物、自身的位姿和速度等信息，以及用于识别是否关注目标的目标特征库。其关键技术包括高效的地图构建与更新技术、基于多传感器的自身位姿与速度估计、目标特征的检索/匹配/评估技术等。\n            "
              )
            ])
          ]
        )
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticStyle: { display: "none" }, attrs: { id: "jgt2" } },
      [
        _c("img", {
          staticClass: "image_31 pos_22",
          staticStyle: {
            position: "relative",
            left: "-161px",
            top: "129px",
            transform: "rotate(180deg)"
          },
          attrs: { src: "/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" }
        }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-col justify-start items-start section_20",
            staticStyle: {
              position: "relative",
              top: "-33px",
              right: "453px",
              height: "154px"
            }
          },
          [
            _c("span", { staticClass: "font_12 text_42" }, [_vm._v("“")]),
            _vm._v(" "),
            _c("span", { staticClass: "font_3 text_41 pos_17" }, [
              _vm._v(
                "\n              感知层不断综合各种传感器数据，基于各类感知算法，构建无人系统的作业环境模型，并确保作业环境模型总是被实时更新。其关键技术包手异构传感器融合框架以支撑多传感器数据融合技术，目标检测\\识别\\跟踪与障碍物检测技术、随机复杂高动态强鲁棒感知认知技术与算法等。\n            "
              )
            ])
          ]
        )
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticStyle: { display: "none" }, attrs: { id: "jgt3" } },
      [
        _c("img", {
          staticClass: "image_31 pos_22",
          staticStyle: {
            position: "relative",
            left: "-5px",
            top: "129px",
            "z-index": "1"
          },
          attrs: { src: "/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" }
        }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-col justify-start items-start section_20",
            staticStyle: {
              position: "relative",
              top: "-33px",
              right: "-10px",
              "z-index": "1",
              height: "142px"
            }
          },
          [
            _c("span", { staticClass: "font_12 text_42" }, [_vm._v("“")]),
            _vm._v(" "),
            _c("span", { staticClass: "font_3 text_41 pos_17" }, [
              _vm._v(
                "\n              无人平台硬件虚拟层对无人平台各类硬件进行虚拟化，从而屏蔽硬件的物理差异，向上提供统一的虚拟模型和访问接口，为自主协同行为提供支撑。其关键技术包括各类无人平台的传感器及执行器的虚拟模型和访问接口设计、无线通信虚拟技术及其接口等。\n            "
              )
            ])
          ]
        )
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticStyle: { display: "none" }, attrs: { id: "jgt4" } },
      [
        _c("img", {
          staticClass: "image_31 pos_22",
          staticStyle: { position: "relative", left: "-5px", top: "129px" },
          attrs: { src: "/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" }
        }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-col justify-start items-start section_20",
            staticStyle: {
              position: "relative",
              top: "-33px",
              right: "-10px",
              height: "190px"
            }
          },
          [
            _c("span", { staticClass: "font_12 text_42" }, [_vm._v("“")]),
            _vm._v(" "),
            _c("span", { staticClass: "font_3 text_41 pos_17" }, [
              _vm._v(
                "\n              虚拟空间自主学习平台运行在后台服务器集群中，用于构建虚拟空间，并实现并行可持续自主学习。其核心模块包括支持高性能、高精度、高保真、高动态的模拟仿真引擎，面向复杂环境和多型无人系统的数据生成引擎，虚拟空间自主学习引擎等。/匹配/评估技术等。针对实体空间的人机物全要素场景，研究构建全视景虚拟作业环境的智能方法和视景渲染技术。\n            "
              )
            ])
          ]
        )
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticStyle: { display: "none" }, attrs: { id: "jgt5" } },
      [
        _c("img", {
          staticClass: "image_31 pos_22",
          staticStyle: { position: "relative", top: "126px", right: "5px" },
          attrs: { src: "/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" }
        }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-col justify-start items-start section_20",
            staticStyle: {
              position: "relative",
              top: "-33px",
              right: "-10px",
              height: "230px"
            }
          },
          [
            _c("span", { staticClass: "font_12 text_42" }, [_vm._v("“")]),
            _vm._v(" "),
            _c("span", { staticClass: "font_3 text_41 pos_17" }, [
              _vm._v(
                "\n              人机协同层支持人机自主协同，包括人机互理解模块和人机自主协同操作模块。人机互理解模块为人机自主协同提供界面、指令和基础算法，其关键技术包括作业指令设计、基于作业环境模型的人机交互以及协同目标识别等算法，协同作业规划模块实现人向无人系统的作业分配方法，作业预规划和动态作业规划等技。人机自主协同操控模块实现作业过程的人对无人系统的粗粒度控制，其关键技术包括自主协同操控通用框架、协同任务规划交互控制、协同路径规划与控制等技术。\n            "
              )
            ])
          ]
        )
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticStyle: { display: "none" }, attrs: { id: "jgt6" } },
      [
        _c("img", {
          staticClass: "image_31 pos_22",
          staticStyle: { position: "relative", top: "127px", right: "5px" },
          attrs: { src: "/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" }
        }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass: "flex-col justify-start items-start section_20",
            staticStyle: {
              position: "relative",
              top: "-33px",
              right: "-10px",
              height: "180px"
            }
          },
          [
            _c("span", { staticClass: "font_12 text_42" }, [_vm._v("“")]),
            _vm._v(" "),
            _c("span", { staticClass: "font_3 text_41 pos_17" }, [
              _vm._v(
                "\n              规划与推理层基于感知信息和作业环境模型，按作业任务要求，开展高层次的规划和推理，将作业分解为可执行的状态机，并规划为无人平台可以执行的动作序列，由作业状态切换引擎和运动规划器组成。其关键技术包括作业状态机管理技术、基于状态机的推理技术、失效和异常处理技术等。运动规划器研究多目标融合决策与规划、长时多模协同混合控制等技术等。\n            "
              )
            ])
          ]
        )
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass:
          "flex-col justify-start items-center text-wrapper_9 pos_97 small"
      },
      [_c("a", { staticClass: "text_97" }, [_vm._v("执行器")])]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass:
          "flex-col justify-start items-center text-wrapper_7 pos_98 small"
      },
      [_c("a", { staticClass: "text_98" }, [_vm._v("运动规划")])]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass:
          "flex-col justify-start items-center text-wrapper_7 pos_99 small"
      },
      [_c("a", { staticClass: "text_99" }, [_vm._v("传感器")])]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col justify-start section_31 pos_46" },
      [_c("div", { staticClass: "section_32" })]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col justify-start section_33 pos_47" },
      [_c("div", { staticClass: "section_32" })]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col justify-start section_34 pos_48" },
      [_c("div", { staticClass: "shrink-0 section_32" })]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticStyle: {
          display: "flex",
          "flex-direction": "column",
          "align-content": "center"
        }
      },
      [
        _c("img", { attrs: { src: "/img/ros-hmci/qrcode.png" } }),
        _vm._v(" "),
        _c("span", { staticClass: "text_95" }, [_vm._v("用户交流群")])
      ]
    )
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=template&id=6cebc5e9&scoped=true":
/*!*********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OSSystem.vue?vue&type=template&id=6cebc5e9&scoped=true ***!
  \*********************************************************************************************************************************************************************************************************************************************/
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
    { staticClass: "flex-col page" },
    [
      _c("AppBanner"),
      _vm._v(" "),
      _c(
        "div",
        {
          staticClass: "flex-col justify-start items-center relative section_6"
        },
        [
          _c("div", { staticClass: "flex-col space-y-72" }, [
            _vm._m(0),
            _vm._v(" "),
            _c("div", { staticClass: "flex-row space-x-92" }, [
              _c(
                "div",
                {
                  ref: "leftElem",
                  staticClass:
                    "flex-col self-start group_5 space-y-54 slide-in-left",
                  class: { show: _vm.showLeft }
                },
                [
                  _vm._m(1),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass: "flex-row items-center self-start section_8"
                    },
                    [
                      _c(
                        "a",
                        {
                          staticClass: "text_19",
                          attrs: { href: _vm.links.os_info }
                        },
                        [_vm._v("更多详情")]
                      ),
                      _vm._v(" "),
                      _c("img", {
                        staticClass: "shrink-0 image_14",
                        attrs: { src: "/img/ros-hmci/mbz607.png" }
                      })
                    ]
                  )
                ]
              ),
              _vm._v(" "),
              _c("img", {
                ref: "rightElem",
                staticClass: "image_27 slide-in-right",
                class: { show: _vm.showRight },
                attrs: { src: "/img/ros-hmci/group4654.png" }
              })
            ])
          ])
        ]
      ),
      _vm._v(" "),
      _c("div", { staticClass: "flex-col justify-start relative" }, [
        _vm._m(2),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col relative" }, [
          _c(
            "div",
            { staticClass: "flex-col justify-start items-center group_1" },
            [
              _c("div", { staticClass: "flex-col group_36 space-y-64" }, [
                _vm._m(3),
                _vm._v(" "),
                _c("div", { staticClass: "flex-col pos_97" }, [
                  _c("div", { staticClass: "flex-col" }, [
                    _c(
                      "div",
                      {
                        staticClass: "flex-row group_10 items-center space-x-82"
                      },
                      [
                        _c(
                          "a",
                          {
                            staticClass: "font_4",
                            class: { sel: _vm.activeTab === "tab1" },
                            on: {
                              mouseenter: function($event) {
                                return _vm.xtjg("tab1")
                              }
                            }
                          },
                          [_vm._v("无人平台硬件虚拟层")]
                        ),
                        _vm._v(" "),
                        _c(
                          "a",
                          {
                            staticClass: "font_4",
                            class: { sel: _vm.activeTab === "tab2" },
                            on: {
                              mouseenter: function($event) {
                                return _vm.xtjg("tab2")
                              }
                            }
                          },
                          [_vm._v("感知层")]
                        ),
                        _vm._v(" "),
                        _c(
                          "a",
                          {
                            staticClass: "font_4",
                            class: { sel: _vm.activeTab === "tab3" },
                            on: {
                              mouseenter: function($event) {
                                return _vm.xtjg("tab3")
                              }
                            }
                          },
                          [_vm._v("规划与推理层")]
                        ),
                        _vm._v(" "),
                        _c(
                          "a",
                          {
                            staticClass: "font_4",
                            class: { sel: _vm.activeTab === "tab4" },
                            on: {
                              mouseenter: function($event) {
                                return _vm.xtjg("tab4")
                              }
                            }
                          },
                          [_vm._v("人机协同层")]
                        ),
                        _vm._v(" "),
                        _c(
                          "a",
                          {
                            staticClass: "font_4",
                            class: { sel: _vm.activeTab === "tab5" },
                            on: {
                              mouseenter: function($event) {
                                return _vm.xtjg("tab5")
                              }
                            }
                          },
                          [_vm._v("虚拟空间自主学习平台")]
                        ),
                        _vm._v(" "),
                        _c(
                          "a",
                          {
                            staticClass: "font_4",
                            class: { sel: _vm.activeTab === "tab6" },
                            on: {
                              mouseenter: function($event) {
                                return _vm.xtjg("tab6")
                              }
                            }
                          },
                          [_vm._v("作业环境模型")]
                        )
                      ]
                    ),
                    _vm._v(" "),
                    _vm._m(4)
                  ]),
                  _vm._v(" "),
                  _vm._m(5),
                  _vm._v(" "),
                  _vm._m(6),
                  _vm._v(" "),
                  _vm._m(7),
                  _vm._v(" "),
                  _vm._m(8),
                  _vm._v(" "),
                  _vm._m(9),
                  _vm._v(" "),
                  _vm._m(10)
                ])
              ])
            ]
          ),
          _vm._v(" "),
          _c(
            "div",
            { staticClass: "flex-col justify-start items-center section_13" },
            [
              _c(
                "div",
                {
                  staticClass: "flex-col space-y-48",
                  attrs: { id: "osDownload" }
                },
                [
                  _vm._m(11),
                  _vm._v(" "),
                  _c("div", { staticClass: "flex-row space-x-14" }, [
                    _c("div", { staticClass: "flex-row relative" }, [
                      _c(
                        "div",
                        { staticClass: "flex-col section_14" },
                        _vm._l(_vm.repos, function(repo, index) {
                          return _c(
                            "div",
                            {
                              ref: "sec" + (index + 1),
                              refInFor: true,
                              staticClass: "sec repo",
                              attrs: { id: "sec" + (index + 1) },
                              on: {
                                mouseenter: function($event) {
                                  _vm.xtxz("sec" + (index + 1))
                                  _vm.selectedRepoName = repo.Name
                                  _vm.selectedRepoAlias = repo.Alias
                                  _vm.selectedRepoOwner = repo.OwnerName
                                }
                              }
                            },
                            [
                              _c(
                                "div",
                                {
                                  staticClass:
                                    "flex-col items-start section_16 xz_active",
                                  style: { display: index === 0 ? "" : "none" }
                                },
                                [
                                  _c("span", { staticClass: "font_10" }, [
                                    _vm._v("人机协同智能操作系统")
                                  ]),
                                  _vm._v(" "),
                                  _c(
                                    "div",
                                    {
                                      staticClass:
                                        "flex-col justify-start text-wrapper_7"
                                    },
                                    [
                                      _c(
                                        "span",
                                        { staticClass: "font_2 text_36" },
                                        [_vm._v(_vm._s(repo.Alias))]
                                      )
                                    ]
                                  )
                                ]
                              ),
                              _vm._v(" "),
                              _c(
                                "div",
                                {
                                  staticClass:
                                    "flex-col justify-start items-center section_31 xz_normal",
                                  style: { display: index === 0 ? "none" : "" }
                                },
                                [
                                  _c("span", { staticClass: "text_37" }, [
                                    _vm._v("人机协同智能操作系统")
                                  ]),
                                  _vm._v(" "),
                                  _c("span", { staticClass: "text_39" }, [
                                    _vm._v("-" + _vm._s(repo.Alias))
                                  ])
                                ]
                              )
                            ]
                          )
                        }),
                        0
                      )
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "flex-col group_34" }, [
                      _c(
                        "div",
                        {
                          ref: "dataset_info",
                          staticClass: "flex-col section_17 visible"
                        },
                        [
                          _c(
                            "span",
                            { staticClass: "self-start font_10 text_35" },
                            [
                              _vm._v(
                                "人机协同智能操作系统-" +
                                  _vm._s(_vm.selectedRepoAlias) +
                                  "简介"
                              )
                            ]
                          ),
                          _vm._v(" "),
                          _c(
                            "div",
                            {
                              staticClass:
                                "flex-col items-start group_13 space-y-2"
                            },
                            [
                              _c("span", { staticClass: "font_11" }, [
                                _vm._v(" " + _vm._s(_vm.description) + " ")
                              ])
                            ]
                          ),
                          _vm._v(" "),
                          _c(
                            "div",
                            {
                              staticClass:
                                "flex-row self-start section_18 space-x-2"
                            },
                            [
                              _c(
                                "a",
                                {
                                  staticClass: "font_8 text_38",
                                  attrs: {
                                    href:
                                      "/" +
                                      _vm.selectedRepoOwner +
                                      "/" +
                                      _vm.selectedRepoName +
                                      "/datasets"
                                  }
                                },
                                [_vm._v("更多详情")]
                              ),
                              _vm._v(" "),
                              _c("img", {
                                staticClass: "shrink-0 image_21",
                                attrs: {
                                  src:
                                    "/img/ros-hmci/8aacfa799145b95562e808dcf6004f2d.png"
                                }
                              })
                            ]
                          ),
                          _vm._v(" "),
                          _c(
                            "div",
                            { staticClass: "flex-col relative section_19" },
                            [
                              _vm._m(12),
                              _vm._v(" "),
                              _vm._l(_vm.attachments, function(attachment) {
                                return _c("div", [
                                  _c("div", { staticClass: "flex-col" }, [
                                    _c(
                                      "div",
                                      { staticClass: "group_14 view_5" },
                                      [
                                        _c("span", { staticClass: "font_14" }, [
                                          _vm._v(_vm._s(attachment.Name))
                                        ]),
                                        _vm._v(" "),
                                        _c(
                                          "a",
                                          {
                                            attrs: {
                                              href:
                                                "/attachments/" +
                                                attachment.UUID +
                                                "?type=0"
                                            }
                                          },
                                          [
                                            _c("img", {
                                              staticClass: "shrink-0 image_22",
                                              attrs: {
                                                src: "/img/ros-hmci/mbz614.png"
                                              }
                                            })
                                          ]
                                        )
                                      ]
                                    ),
                                    _vm._v(" "),
                                    _c(
                                      "div",
                                      {
                                        staticClass:
                                          "flex-row items-baseline group_14"
                                      },
                                      [
                                        _c(
                                          "a",
                                          {
                                            staticClass: "font_13 text_42",
                                            attrs: {
                                              href:
                                                "/" +
                                                _vm.selectedRepoOwner +
                                                "/" +
                                                _vm.selectedRepoName +
                                                "/datasets"
                                            }
                                          },
                                          [_vm._v("本次改动内容...")]
                                        ),
                                        _vm._v(" "),
                                        _c("span", { staticClass: "font_15" }, [
                                          _vm._v(
                                            _vm._s(
                                              _vm._f("formatDate")(
                                                attachment.CreatedUnix
                                              )
                                            )
                                          )
                                        ])
                                      ]
                                    )
                                  ]),
                                  _vm._v(" "),
                                  _c("div", { staticClass: "divider_2" })
                                ])
                              })
                            ],
                            2
                          )
                        ]
                      )
                    ])
                  ])
                ]
              )
            ]
          )
        ])
      ])
    ],
    1
  )
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col justify-start self-center relative group_3" },
      [
        _c("span", { staticClass: "text_16 pos_99" }, [
          _vm._v("关于人机协同智能操作系统")
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "flex-col" }, [
      _c("div", { staticClass: "flex-row space-x-16" }, [
        _c("img", {
          staticClass: "image_12",
          attrs: { src: "/img/ros-hmci/mbz606.png" }
        }),
        _vm._v(" "),
        _c("span", { staticClass: "font_41 text_17" }, [_vm._v("系统概述")])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "flex-col items-center group_6 space-y-2" }, [
        _c("span", { staticClass: "font_5" }, [
          _vm._v(
            "人机协同智能操作系统是一组软件库和工具，可帮助您构建人工智能应用程序。\n                从驱动程序到最先进的算法，海量的数据集，模型，以及强大的开发人员工具...\n              "
          )
        ])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "flex-row group_7 space-x-8" }, [
        _c("div", { staticClass: "flex-col justify-start text-wrapper_2" }, [
          _c("span", { staticClass: "font_6 text_18" }, [
            _vm._v("人机协同智能操作系统...")
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col justify-start text-wrapper_2" }, [
          _c("span", { staticClass: "font_6 text_18" }, [
            _vm._v("复杂环境与作业规划器")
          ])
        ])
      ])
    ])
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col justify-start section_9 pos_9" },
      [
        _c(
          "div",
          {
            staticClass:
              "flex-col justify-start items-center relative section_30"
          },
          [
            _c("img", {
              staticClass: "image_1",
              attrs: {
                src: "/img/ros-hmci/6da4828190161924d3e2c084e55fa79e.png"
              }
            }),
            _vm._v(" "),
            _c("img", {
              staticClass: "image_13 pos_1",
              attrs: {
                src: "/img/ros-hmci/95ecb911f420e03fb3ade00d4bc53614.png"
              }
            })
          ]
        )
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col self-center group_8 space-y-26" },
      [
        _c(
          "div",
          {
            staticClass: "flex-col justify-start items-center relative group_9"
          },
          [
            _c("span", { staticClass: "text_16 pos_98" }, [
              _vm._v("人机协同智能操作系统架构")
            ])
          ]
        ),
        _vm._v(" "),
        _c("span", { staticClass: "text_21" }, [
          _vm._v(
            "\n              运行于传统的计算机操作系统之上，实现智能无人平台的传感器、执行器、通信系统等硬件资源的虚拟化，管理智能无人系统的感知、规划、推理等自主行为，并对人机自主协同和持续自主学习提供支撑\n            "
          )
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col justify-start items-start divider" },
      [_c("div", { staticClass: "section_10" })]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "tab visible flex-row items-center group_31 space-x-40",
        attrs: { id: "tab1" }
      },
      [
        _c("div", { staticClass: "flex-col group_26 space-y-18" }, [
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_1",
              attrs: { src: "/img/ros-hmci/mbz609.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7 text_25" }, [
                _vm._v("传感器")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "flex-col space-y-9" }, [
                _c("div", { staticClass: "flex-row self-start space-x-9" }, [
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_1" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("激光扫描仪")
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_3" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("可见光/红外像机")
                      ])
                    ]
                  )
                ]),
                _vm._v(" "),
                _c("div", { staticClass: "flex-row space-x-9" }, [
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start items-center text-wrapper_1"
                    },
                    [_c("span", { staticClass: "text_26" }, [_vm._v("GNSS")])]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_3" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("加速度计/陀螺仪")
                      ])
                    ]
                  )
                ])
              ])
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_1",
              attrs: { src: "/img/ros-hmci/mbz610.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7" }, [
                _vm._v("实时无线通信")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "flex-col" }, [
                _c("div", { staticClass: "flex-row group_30 space-x-9" }, [
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start items-center text-wrapper_1"
                    },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("自组织链路")
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_3" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("高可靠无人链路")
                      ])
                    ]
                  )
                ]),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start self-start text-wrapper_16"
                  },
                  [
                    _c("span", { staticClass: "text_26" }, [
                      _vm._v("实时信息通信")
                    ])
                  ]
                )
              ])
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_1",
              attrs: { src: "/img/ros-hmci/mbz611.png" }
            }),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col items-start group_29 space-y-12" },
              [
                _c("span", { staticClass: "font_7" }, [_vm._v("执行器")]),
                _vm._v(" "),
                _c("div", { staticClass: "flex-col space-y-9" }, [
                  _c("div", { staticClass: "flex-row space-x-9" }, [
                    _c(
                      "div",
                      {
                        staticClass:
                          "flex-col justify-start items-center text-wrapper_1"
                      },
                      [_c("span", { staticClass: "text_26" }, [_vm._v("姿态")])]
                    ),
                    _vm._v(" "),
                    _c(
                      "div",
                      {
                        staticClass:
                          "flex-col justify-start items-center text-wrapper_5"
                      },
                      [_c("span", { staticClass: "text_26" }, [_vm._v("位置")])]
                    )
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start items-center self-start text-wrapper_6"
                    },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("速度/加速度")
                      ])
                    ]
                  )
                ])
              ]
            )
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col group_33 space-y-56" }, [
          _c("div", { staticClass: "flex-col space-y-32" }, [
            _c("img", {
              staticClass: "image_17",
              attrs: { src: "/img/ros-hmci/架构图1.png" }
            }),
            _vm._v(" "),
            _c("span", { staticClass: "font_5 text_31" }, [
              _vm._v(
                "\n                    无人平台硬件虚拟层对无人平台各类硬件进行虚拟化，从而屏蔽硬件的物理差异，向上提供统一的虚拟模型和访问接口，为自主协同行为提供支撑。其关键技术包括各类无人平台的传感器及执行器的虚拟模型和访问接口设计、无线通信虚拟技术及其接口等。\n                  "
              )
            ])
          ])
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "tab flex-row items-center group_31 space-x-40",
        attrs: { id: "tab2" }
      },
      [
        _c("div", { staticClass: "flex-col group_26 space-y-18" }, [
          _c("div", { staticClass: "flex-col justify-start section_123" }, [
            _c("img", {
              staticClass: "imgpos_5",
              attrs: { src: "/img/ros-hmci/group668.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7 text_25" }, [
                _vm._v("感知模块")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "flex-col space-y-9" }, [
                _c("div", { staticClass: "flex-row self-start space-x-9" }, [
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_1" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("目标检测")
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_3" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("障碍物检测")
                      ])
                    ]
                  )
                ]),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start self-start text-wrapper_16"
                  },
                  [
                    _c("span", { staticClass: "text_26" }, [
                      _vm._v("目标识别与跟踪")
                    ])
                  ]
                ),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start self-start text-wrapper_16"
                  },
                  [
                    _c("span", { staticClass: "text_26" }, [
                      _vm._v("异构传感器数据融合框架")
                    ])
                  ]
                ),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start self-start text-wrapper_16"
                  },
                  [
                    _c("span", { staticClass: "text_26" }, [
                      _vm._v("随机复杂高动态强鲁棒感知认知")
                    ])
                  ]
                )
              ])
            ])
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col group_33 space-y-56" }, [
          _c("div", { staticClass: "flex-col space-y-32" }, [
            _c("img", {
              staticClass: "image_17",
              attrs: { src: "/img/ros-hmci/架构图2.png" }
            }),
            _vm._v(" "),
            _c("span", { staticClass: "font_5 text_31" }, [
              _vm._v(
                "\n                    感知层不断综合各种传感器数据，基于各类感知算法，构建无人系统的作业环境模型，并确保作业环境模型总是被实时更新。其关键技术包手异构传感器融合框架以支撑多传感器数据融合技术，目标检测\\识别\\跟踪与障碍物检测技术、随机复杂高动态强鲁棒感知认知技术与算法等。\n                  "
              )
            ])
          ])
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "tab flex-row items-center group_31 space-x-40",
        attrs: { id: "tab3" }
      },
      [
        _c("div", { staticClass: "flex-col group_26 space-y-18" }, [
          _c("div", { staticClass: "flex-col justify-start section_122" }, [
            _c("img", {
              staticClass: "imgpos_4",
              attrs: { src: "/img/ros-hmci/group7079.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7 text_25" }, [
                _vm._v("作业状态切换引擎")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "flex-col space-y-9" }, [
                _c("div", { staticClass: "flex-row self-start space-x-9" }, [
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_16" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("作业状态机管理")
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_16" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("基于状态机的推理")
                      ])
                    ]
                  )
                ]),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start items-center self-start text-wrapper_16"
                  },
                  [_c("span", { staticClass: "text_26" }, [_vm._v("失效处理")])]
                )
              ])
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col justify-start section_122" }, [
            _c("img", {
              staticClass: "imgpos_4",
              attrs: { src: "/img/ros-hmci/group7103.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7" }, [
                _vm._v("失效处理")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "flex-col" }, [
                _c("div", { staticClass: "flex-row group_30 space-x-9" }, [
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start items-center text-wrapper_16"
                    },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("多目标融合规划...")
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_16" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("长时多模协同混...")
                      ])
                    ]
                  )
                ])
              ])
            ])
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col group_33 space-y-56" }, [
          _c("div", { staticClass: "flex-col space-y-32" }, [
            _c("img", {
              staticClass: "image_17",
              attrs: { src: "/img/ros-hmci/架构图3.png" }
            }),
            _vm._v(" "),
            _c("span", { staticClass: "font_5 text_31" }, [
              _vm._v(
                "\n                    规划与推理层基于感知信息和作业环境模型，按作业任务要求，开展高层次的规划和推理，将作业分解为可执行的状态机，并规划为无人平台可以执行的动作序列，由作业状态切换引擎和运动规划器组成。其关键技术包括作业状态机管理技术、基于状态机的推理技术、失效和异常处理技术等。运动规划器研究多目标融合决策与规划、长时多模协同混合控制等技术等。\n                  "
              )
            ])
          ])
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "tab flex-row items-center group_31 space-x-40",
        attrs: { id: "tab4" }
      },
      [
        _c("div", { staticClass: "flex-col group_26 space-y-18" }, [
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_1",
              attrs: { src: "/img/ros-hmci/mbz609.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7 text_25" }, [
                _vm._v("人机互理解")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "flex-col space-y-9" }, [
                _c("div", { staticClass: "flex-row self-start space-x-9" }, [
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_16" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("作业指令设计")
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_16" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("协同目标识别")
                      ])
                    ]
                  )
                ]),
                _vm._v(" "),
                _c(
                  "div",
                  {
                    staticClass:
                      "flex-col justify-start items-center self-start text-wrapper_16"
                  },
                  [
                    _c("span", { staticClass: "text_26" }, [
                      _vm._v("基于作业环境模型的人机交互")
                    ])
                  ]
                )
              ])
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_1",
              attrs: { src: "/img/ros-hmci/mbz610.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7" }, [
                _vm._v("协同作业规划")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "flex-col" }, [
                _c("div", { staticClass: "flex-row group_30 space-x-9" }, [
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start items-center text-wrapper_16"
                    },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("作业预规划")
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_16" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("动态作业规划")
                      ])
                    ]
                  )
                ])
              ])
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_1",
              attrs: { src: "/img/ros-hmci/mbz611.png" }
            }),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col items-start group_29 space-y-12" },
              [
                _c("span", { staticClass: "font_7" }, [
                  _vm._v("人机自主协同操控")
                ]),
                _vm._v(" "),
                _c("div", { staticClass: "flex-col space-y-9" }, [
                  _c("div", { staticClass: "flex-row space-x-9" }, [
                    _c(
                      "div",
                      {
                        staticClass:
                          "flex-col justify-start items-center text-wrapper_16"
                      },
                      [
                        _c("span", { staticClass: "text_26" }, [
                          _vm._v("自主协同操控...")
                        ])
                      ]
                    ),
                    _vm._v(" "),
                    _c(
                      "div",
                      {
                        staticClass:
                          "flex-col justify-start items-center text-wrapper_16"
                      },
                      [
                        _c("span", { staticClass: "text_26" }, [
                          _vm._v("协同任务规划...")
                        ])
                      ]
                    )
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start items-center self-start text-wrapper_16"
                    },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("协同路径规划与控制")
                      ])
                    ]
                  )
                ])
              ]
            )
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col group_33 space-y-56" }, [
          _c("div", { staticClass: "flex-col space-y-32" }, [
            _c("img", {
              staticClass: "image_17",
              attrs: { src: "/img/ros-hmci/架构图4.png" }
            }),
            _vm._v(" "),
            _c("span", { staticClass: "font_5 text_31" }, [
              _vm._v(
                "\n                    人机协同层支持人机自主协同，包括人机互理解模块和人机自主协同操作模块。人机互理解模块为人机自主协同提供界面、指令和基础算法，其关键技术包括作业指令设计、基于作业环境模型的人机交互以及协同目标识别等算法，协同作业规划模块实现人向无人系统的作业分配方法，作业预规划和动态作业规划等技。人机自主协同操控模块实现作业过程的人对无人系统的粗粒度控制，其关键技术包括自主协同操控通用框架、协同任务规划交互控制、协同路径规划与控制等技术。\n                  "
              )
            ])
          ])
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "tab flex-row items-center group_31 space-x-40",
        attrs: { id: "tab5" }
      },
      [
        _c("div", { staticClass: "flex-col group_26 space-y-181" }, [
          _c("div", { staticClass: "flex-col justify-start section_121" }, [
            _c("img", {
              staticClass: "imgpos_3",
              attrs: { src: "/img/ros-hmci/mbz658.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7 text_25" }, [
                _vm._v("管理平台")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "flex-col space-y-9" }, [
                _c("div", { staticClass: "flex-row self-start space-x-9" }, [
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_16" },
                    [_c("span", { staticClass: "text_26" }, [_vm._v("资源")])]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_16" },
                    [_c("span", { staticClass: "text_26" }, [_vm._v("数据")])]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_16" },
                    [_c("span", { staticClass: "text_26" }, [_vm._v("任务")])]
                  )
                ])
              ])
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_1",
              attrs: { src: "/img/ros-hmci/mbz610.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7" }, [
                _vm._v("虚拟学习引擎")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "flex-col" }, [
                _c("div", { staticClass: "flex-row group_30 space-x-9" }, [
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start items-center text-wrapper_16"
                    },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("自主强化学习")
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex-col justify-start text-wrapper_16" },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("人机融合学习")
                      ])
                    ]
                  )
                ]),
                _vm._v(" "),
                _c("div", { staticClass: "flex-row group_30 space-x-9" }, [
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start self-start text-wrapper_16"
                    },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("大时空尺度持...")
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start self-start text-wrapper_16"
                    },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("规模化场景学习...")
                      ])
                    ]
                  )
                ])
              ])
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_1",
              attrs: { src: "/img/ros-hmci/mbz611.png" }
            }),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col items-start group_29 space-y-12" },
              [
                _c("span", { staticClass: "font_7" }, [_vm._v("数据生成引擎")]),
                _vm._v(" "),
                _c("div", { staticClass: "flex-col space-y-9" }, [
                  _c("div", { staticClass: "flex-row space-x-9" }, [
                    _c(
                      "div",
                      {
                        staticClass:
                          "flex-col justify-start items-center text-wrapper_6"
                      },
                      [
                        _c("span", { staticClass: "text_26" }, [
                          _vm._v("作业环境生成")
                        ])
                      ]
                    ),
                    _vm._v(" "),
                    _c(
                      "div",
                      {
                        staticClass:
                          "flex-col justify-start items-center text-wrapper_6"
                      },
                      [
                        _c("span", { staticClass: "text_26" }, [
                          _vm._v("无人系统生成")
                        ])
                      ]
                    )
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start items-center self-start text-wrapper_6"
                    },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("虚拟目标生成")
                      ])
                    ]
                  )
                ])
              ]
            )
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_1",
              attrs: { src: "/img/ros-hmci/mbz611.png" }
            }),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col items-start group_29 space-y-12" },
              [
                _c("span", { staticClass: "font_7" }, [_vm._v("模拟仿真引擎")]),
                _vm._v(" "),
                _c("div", { staticClass: "flex-col space-y-9" }, [
                  _c("div", { staticClass: "flex-row space-x-9" }, [
                    _c(
                      "div",
                      {
                        staticClass:
                          "flex-col justify-start items-center text-wrapper_16"
                      },
                      [
                        _c("span", { staticClass: "text_26" }, [
                          _vm._v("精准碰撞检测技术")
                        ])
                      ]
                    ),
                    _vm._v(" "),
                    _c(
                      "div",
                      {
                        staticClass:
                          "flex-col justify-start items-center text-wrapper_16"
                      },
                      [
                        _c("span", { staticClass: "text_26" }, [
                          _vm._v("全要素模拟技术")
                        ])
                      ]
                    )
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    {
                      staticClass:
                        "flex-col justify-start items-center self-start text-wrapper_16"
                    },
                    [
                      _c("span", { staticClass: "text_26" }, [
                        _vm._v("并行计算仿真技术")
                      ])
                    ]
                  )
                ])
              ]
            )
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col group_331 space-y-56" }, [
          _c("div", { staticClass: "flex-col space-y-32" }, [
            _c("img", {
              staticClass: "image_17",
              attrs: { src: "/img/ros-hmci/架构图1.png" }
            }),
            _vm._v(" "),
            _c("span", { staticClass: "font_5 text_31" }, [
              _vm._v(
                "\n                    虚拟空间自主学习平台运行在后台服务器集群中，用于构建虚拟空间，并实现并行可持续自主学习。其核心模块包括支持高性能、高精度、高保真、高动态的模拟仿真引擎，面向复杂环境和多型无人系统的数据生成引擎，虚拟空间自主学习引擎等。/匹配/评估技术等。针对实体空间的人机物全要素场景，研究构建全视景虚拟作业环境的智能方法和视景渲染技术。\n                  "
              )
            ])
          ])
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "tab flex-row items-center group_31 space-x-40",
        attrs: { id: "tab6" }
      },
      [
        _c("div", { staticClass: "flex-col group_26 space-y-18" }, [
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_2",
              attrs: { src: "/img/ros-hmci/group7310.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7 text_25" }, [
                _vm._v("地图和障碍物")
              ])
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col justify-start section_12" }, [
            _c("img", {
              staticClass: "imgpos_2",
              attrs: { src: "/img/ros-hmci/group7311.png" }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_29 space-y-12" }, [
              _c("span", { staticClass: "self-start font_7" }, [
                _vm._v("自身位姿和速度")
              ])
            ])
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col group_33 space-y-56" }, [
          _c("div", { staticClass: "flex-col space-y-32" }, [
            _c("img", {
              staticClass: "image_17",
              attrs: { src: "/img/ros-hmci/架构图2.png" }
            }),
            _vm._v(" "),
            _c("span", { staticClass: "font_5 text_31" }, [
              _vm._v(
                "\n                    作业环境模型中记录了机器和算法可理解的地图和障碍物、自身的位姿和速度等信息，以及用于识别是否关注目标的目标特征库。其关键技术包括高效的地图构建与更新技术、基于多传感器的自身位姿与速度估计、目标特征的检索/匹配/评估技术等。\n                  "
              )
            ])
          ])
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass:
          "flex-col justify-start items-center self-center relative group_12"
      },
      [
        _c("span", { staticClass: "font_3 text_14 text_34 pos_17" }, [
          _vm._v("系统下载")
        ])
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      {
        staticClass: "flex-col justify-start items-center text-wrapper_8 pos_18"
      },
      [_c("span", { staticClass: "font_2" }, [_vm._v("版本列表")])]
    )
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=template&id=4aaca311&scoped=true":
/*!********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenApp.vue?vue&type=template&id=4aaca311&scoped=true ***!
  \********************************************************************************************************************************************************************************************************************************************/
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
    _c(
      "div",
      { staticClass: "group" },
      [_c("NavigationBar"), _vm._v(" "), _vm._m(0)],
      1
    ),
    _vm._v(" "),
    _c(
      "div",
      { staticClass: "ui container" },
      [
        _c("SearchBar", {
          ref: "searchBarRef",
          attrs: {
            static: true,
            staticTopicsData: _vm.staticSquareTopics,
            type: "square",
            sort: "",
            searchValue: _vm.reposListQurey,
            topic: ""
          },
          on: { change: _vm.searchBarChange }
        })
      ],
      1
    ),
    _vm._v(" "),
    _c("div", { staticClass: "ui container" }, [
      _c("div", { staticClass: "ui grid" }, [
        _c("div", { staticClass: "computer only ui two wide computer column" }),
        _vm._v(" "),
        _c(
          "div",
          {
            staticClass:
              "ui sixteen wide mobile twelve wide tablet ten wide computer column"
          },
          [
            _c("ReposList", {
              ref: "reposListRef",
              attrs: {
                sort: _vm.reposListSortType,
                q: _vm.reposListQurey,
                topic: _vm.reposListTopic,
                page: _vm.page,
                pageSize: _vm.pageSize,
                pageSizes: _vm.pageSizes
              },
              on: {
                "current-change": _vm.currentChange,
                "size-change": _vm.sizeChange
              }
            })
          ],
          1
        )
      ])
    ])
  ])
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "info" }, [
      _c("div", { staticClass: "line1 flex-row" }, [
        _c("span", { staticClass: "line1_1" }, [_vm._v("不可错过的社区活动")]),
        _vm._v(" "),
        _c("div", { staticClass: "line1_sqhd" }, [
          _c("span", { staticClass: "line1_2" }, [_vm._v("社区活动")])
        ])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "info-div1" }, [_vm._v("开源社区活动来袭")]),
      _vm._v(" "),
      _c("div", { staticClass: "info-div" }, [
        _c("span", { staticClass: "info2" }, [
          _vm._v("欢迎您使用鹏城·盘古SDK 发布版——pcl_pangu v1.2！")
        ])
      ])
    ])
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=template&id=726aa430&scoped=true":
/*!************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenDataset.vue?vue&type=template&id=726aa430&scoped=true ***!
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
  return _c(
    "div",
    [
      _c("AppBanner"),
      _vm._v(" "),
      _c("div", { staticClass: "repos--seach datasetSearch" }, [
        _c("div", { staticClass: "ui container" }, [
          _c("div", { staticClass: "ui two column centered grid" }, [
            _c(
              "div",
              {
                staticClass:
                  "fourteen wide mobile ten wide tablet ten wide computer column ui form ignore-dirty"
              },
              [
                _c(
                  "div",
                  { staticClass: "ui fluid action input datasetSearchInput" },
                  [
                    _c("input", {
                      directives: [
                        {
                          name: "model",
                          rawName: "v-model",
                          value: _vm.searchValue,
                          expression: "searchValue"
                        }
                      ],
                      attrs: {
                        name: "q",
                        value: "",
                        placeholder: "搜索数据集",
                        autofocus: ""
                      },
                      domProps: { value: _vm.searchValue },
                      on: {
                        keyup: function($event) {
                          if (
                            !$event.type.indexOf("key") &&
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
                          _vm.searchFlag = !_vm.searchFlag
                        },
                        input: function($event) {
                          if ($event.target.composing) {
                            return
                          }
                          _vm.searchValue = $event.target.value
                        }
                      }
                    }),
                    _vm._v(" "),
                    _c(
                      "button",
                      {
                        staticClass: "ui green button",
                        on: {
                          click: function($event) {
                            _vm.searchFlag = !_vm.searchFlag
                          }
                        }
                      },
                      [
                        _vm._v(
                          "\n              " +
                            _vm._s(_vm.$t("repos.search")) +
                            "\n            "
                        )
                      ]
                    )
                  ]
                )
              ]
            )
          ])
        ])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "ui container" }, [
        _c("div", { staticClass: "ui grid" }, [
          _c(
            "div",
            { staticClass: "computer only four wide computer column" },
            [
              _c("div", { staticClass: "ui sixteen wide column" }, [
                _c(
                  "div",
                  {
                    staticStyle: {
                      "font-size": "24px",
                      color: "rgba(16, 16, 16, 1)",
                      height: "40px",
                      "line-height": "40px",
                      "margin-bottom": "2rem"
                    }
                  },
                  [
                    _vm._v(
                      "\n            " +
                        _vm._s(_vm.$t("dataset")) +
                        "\n          "
                    )
                  ]
                ),
                _vm._v(" "),
                _c("div", { staticClass: "mg-b-2" }, [
                  _c("div", { staticClass: "flex mg-b-1" }, [
                    _c("h3", { staticClass: "font-medium" }, [
                      _vm._v(
                        "\n                " +
                          _vm._s(_vm.$t("datasets.license")) +
                          "\n                "
                      ),
                      _vm.licenseFlag
                        ? _c(
                            "span",
                            {
                              staticClass:
                                "mg-l-1 underline text-gray-400 text-sm",
                              staticStyle: { cursor: "pointer" },
                              on: {
                                click: function($event) {
                                  return _vm.clearSelectLeft("license")
                                }
                              }
                            },
                            [_vm._v("Clear")]
                          )
                        : _vm._e()
                    ])
                  ]),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "flex flex-wrap history-content" },
                    _vm._l(_vm.License, function(item) {
                      return _c(
                        "a",
                        {
                          key: item.name,
                          staticClass: "tag",
                          class: item.active ? "tag-active" : "tag-gray",
                          on: {
                            click: function($event) {
                              return _vm.selectLicense(item)
                            }
                          }
                        },
                        [_c("span", [_vm._v(_vm._s(item.name))])]
                      )
                    }),
                    0
                  )
                ])
              ])
            ]
          ),
          _vm._v(" "),
          _c(
            "div",
            {
              staticClass:
                "ui sixteen wide mobile sixteen wide tablet twelve wide computer column"
            },
            [
              _c(
                "div",
                { staticClass: "ui sixteen wide column" },
                [
                  _c(
                    "el-tabs",
                    {
                      on: { "tab-click": _vm.handleClick },
                      model: {
                        value: _vm.activeName,
                        callback: function($$v) {
                          _vm.activeName = $$v
                        },
                        expression: "activeName"
                      }
                    },
                    [
                      _c(
                        "el-tab-pane",
                        {
                          attrs: {
                            label: _vm.$t("datasets.publick_dataset"),
                            name: "public_datasets"
                          }
                        },
                        [
                          _vm.activeName === "public_datasets"
                            ? _c(
                                "div",
                                [
                                  _c("public-dataset", {
                                    attrs: {
                                      isSigned: _vm.isSigned,
                                      dataGet: _vm.activeName,
                                      searchValue: _vm.searchValue,
                                      searchFlag: _vm.searchFlag,
                                      categoryValue: _vm.categoryValue,
                                      taskValue: _vm.taskValue,
                                      licenseValue: _vm.licenseValue
                                    },
                                    on: { getLabel: _vm.getChildLabel }
                                  })
                                ],
                                1
                              )
                            : _vm._e()
                        ]
                      ),
                      _vm._v(" "),
                      _vm.isSigned === "true"
                        ? _c(
                            "el-tab-pane",
                            {
                              attrs: {
                                label: _vm.$t("datasets.my_dataset"),
                                name: "my_datasets"
                              }
                            },
                            [
                              _vm.activeName === "my_datasets"
                                ? _c(
                                    "div",
                                    [
                                      _c("public-dataset", {
                                        attrs: {
                                          isSigned: _vm.isSigned,
                                          dataGet: _vm.activeName,
                                          searchValue: _vm.searchValue,
                                          searchFlag: _vm.searchFlag,
                                          categoryValue: _vm.categoryValue,
                                          taskValue: _vm.taskValue,
                                          licenseValue: _vm.licenseValue
                                        }
                                      })
                                    ],
                                    1
                                  )
                                : _vm._e()
                            ]
                          )
                        : _vm._e(),
                      _vm._v(" "),
                      _vm.isSigned === "true"
                        ? _c(
                            "el-tab-pane",
                            {
                              attrs: {
                                label: _vm.$t("datasets.favorite_dataset"),
                                name: "my_favorite_datasets"
                              }
                            },
                            [
                              _vm.activeName === "my_favorite_datasets"
                                ? _c(
                                    "div",
                                    [
                                      _c("public-dataset", {
                                        attrs: {
                                          isSigned: _vm.isSigned,
                                          dataGet: _vm.activeName,
                                          searchValue: _vm.searchValue,
                                          searchFlag: _vm.searchFlag,
                                          categoryValue: _vm.categoryValue,
                                          taskValue: _vm.taskValue,
                                          licenseValue: _vm.licenseValue
                                        }
                                      })
                                    ],
                                    1
                                  )
                                : _vm._e()
                            ]
                          )
                        : _vm._e()
                    ],
                    1
                  )
                ],
                1
              )
            ]
          )
        ])
      ])
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=template&id=8071ebce&scoped=true":
/*!**********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/OpenModel.vue?vue&type=template&id=8071ebce&scoped=true ***!
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
  return _c(
    "div",
    [
      _c("AppBanner"),
      _vm._v(" "),
      _c("div", { staticClass: "search-bar-wrap" }, [
        _c("div", { staticClass: "search-bar" }, [
          _c("input", {
            directives: [
              {
                name: "model",
                rawName: "v-model",
                value: _vm.condition.q,
                expression: "condition.q"
              }
            ],
            attrs: {
              type: "text",
              placeholder: _vm.$t("modelObj.model_search")
            },
            domProps: { value: _vm.condition.q },
            on: {
              keyup: function($event) {
                if (
                  !$event.type.indexOf("key") &&
                  _vm._k($event.keyCode, "enter", 13, $event.key, "Enter")
                ) {
                  return null
                }
                return _vm.conditionChange($event)
              },
              input: function($event) {
                if ($event.target.composing) {
                  return
                }
                _vm.$set(_vm.condition, "q", $event.target.value)
              }
            }
          }),
          _vm._v(" "),
          _c("button", { on: { click: _vm.conditionChange } }, [
            _vm._v(_vm._s(_vm.$t("repos.search")))
          ])
        ])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "ui container" }, [
        _c("div", { staticClass: "content" }, [
          _c(
            "div",
            { staticClass: "content-l" },
            [
              _c("ModelFilters", {
                attrs: { condition: _vm.condition },
                on: { changeCondition: _vm.conditionChange }
              })
            ],
            1
          ),
          _vm._v(" "),
          _c(
            "div",
            { staticClass: "content-r" },
            [
              _c("ModelCondition", {
                attrs: { condition: _vm.condition },
                on: { changeCondition: _vm.conditionChange }
              }),
              _vm._v(" "),
              _c("ModelList", {
                ref: "modelListRef",
                attrs: { condition: _vm.condition },
                on: { changeCondition: _vm.conditionChange }
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



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=template&id=c265b356&scoped=true":
/*!***************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/ros-hmci/views/ResourceDetail.vue?vue&type=template&id=c265b356&scoped=true ***!
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
  return _c(
    "div",
    [
      _c("AppBanner"),
      _vm._v(" "),
      _c("div", { staticClass: "flex-col" }, [
        _vm._m(0),
        _vm._v(" "),
        _c("div", { staticClass: "flex-col section_7 space-y-17" }, [
          _c("div", { staticClass: "flex-col section_9" }, [
            _vm._m(1),
            _vm._v(" "),
            _c("div", { staticClass: "flex-col group_5" }, [
              _c(
                "div",
                { staticClass: "flex-row items-center self-start space-x-16" },
                [
                  _c("span", { staticClass: "font_4 text_bq" }, [
                    _vm._v("标签")
                  ]),
                  _vm._v(" "),
                  _c("span", { staticClass: "font_2" }, [
                    _vm._v(
                      "\n                            " +
                        _vm._s(_vm.resourceDetails.name) +
                        "\n                        "
                    )
                  ])
                ]
              ),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "flex-row items-center self-start group_7" },
                [
                  _c("span", { staticClass: "font_4" }, [_vm._v("版本号")]),
                  _vm._v(" "),
                  _c("span", { staticClass: "font_2 text_331" }, [
                    _vm._v(_vm._s(_vm.resourceDetails.version))
                  ])
                ]
              ),
              _vm._v(" "),
              _c("div", { staticClass: "flex-row items-center group_7" }, [
                _c("span", { staticClass: "font_4" }, [_vm._v("资源地址")]),
                _vm._v(" "),
                _c("img", {
                  staticClass: "image_13",
                  attrs: {
                    src: "/img/ros-hmci/e6222f48f76ae9ba015d4748d05ca0f6.png"
                  }
                }),
                _vm._v(" "),
                _c(
                  "a",
                  {
                    staticClass: "font_2",
                    attrs: {
                      href: _vm.resourceDetails.address,
                      target: "_blank"
                    }
                  },
                  [_vm._v(_vm._s(_vm.resourceDetails.address))]
                )
              ]),
              _vm._v(" "),
              _c(
                "div",
                {
                  staticClass:
                    "flex-row items-center self-start group_7 space-x-16"
                },
                [
                  _c("span", { staticClass: "font_4" }, [_vm._v("开源协议")]),
                  _vm._v(" "),
                  _c("span", { staticClass: "font_2" }, [
                    _vm._v(_vm._s(_vm.resourceDetails.License))
                  ])
                ]
              ),
              _vm._v(" "),
              _c(
                "div",
                {
                  staticClass:
                    "flex-row items-center self-start group_7 space-x-16"
                },
                [
                  _c("span", { staticClass: "font_4" }, [_vm._v("最后更新")]),
                  _vm._v(" "),
                  _c("span", { staticClass: "font_2" }, [
                    _vm._v(_vm._s(_vm.resourceDetails.update_time))
                  ])
                ]
              ),
              _vm._v(" "),
              _c(
                "div",
                {
                  staticClass:
                    "flex-row items-center self-start group_7 space-x-16"
                },
                [
                  _c("span", { staticClass: "font_4" }, [_vm._v("资源简介")]),
                  _vm._v(" "),
                  _c("span", { staticClass: "font_2" }, [
                    _vm._v(_vm._s(_vm.resourceDetails.synopsis))
                  ])
                ]
              )
            ])
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "flex-col markdown" }, [
            _vm._m(2),
            _vm._v(" "),
            _c("div", {
              staticClass: "md_content",
              domProps: { innerHTML: _vm._s(_vm.htmlContent) }
            })
          ])
        ])
      ])
    ],
    1
  )
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col items-start section_6 space-y-5" },
      [
        _c("div", { staticClass: "flex-row items-center space-x-14" }, [
          _c("img", {
            staticClass: "shrink-0 image_10",
            attrs: { src: "/img/ros-hmci/aabc8fe78ef76d15c2805169fa04c4c9.png" }
          }),
          _vm._v(" "),
          _c("span", { staticClass: "text_15" }, [
            _vm._v("资源列表/人机协同侦察系统")
          ])
        ]),
        _vm._v(" "),
        _c(
          "div",
          { staticClass: "flex-row justify-center group_3 space-x-10" },
          [
            _c(
              "div",
              { staticClass: "flex-col justify-start text-wrapper_2" },
              [
                _c("span", { staticClass: "font_3 text_16" }, [
                  _vm._v("人机协同智能操作系统...")
                ])
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col justify-start text-wrapper_2" },
              [
                _c("span", { staticClass: "font_3 text_17" }, [
                  _vm._v("虚拟学习引擎")
                ])
              ]
            ),
            _vm._v(" "),
            _c(
              "div",
              { staticClass: "flex-col justify-start text-wrapper_2" },
              [
                _c("span", { staticClass: "font_3 text_18" }, [
                  _vm._v("作业环境仿真器与数字...")
                ])
              ]
            )
          ]
        )
      ]
    )
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "flex-col items-start section_8" }, [
      _c("div", { staticClass: "flex-row justify-center group_4 space-x-36" }, [
        _c("div", { staticClass: "flex-row items-start justify-center" }, [
          _c("img", {
            staticClass: "shrink-0 image_11",
            attrs: { src: "/img/ros-hmci/mbz628.png" }
          }),
          _vm._v(" "),
          _c("span", { staticClass: "text_19" }, [_vm._v("概述")])
        ])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "section_10" })
    ])
  },
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c(
      "div",
      { staticClass: "flex-col items-start justify-center section_88" },
      [_c("span", { staticClass: "text_detail" }, [_vm._v("详细介绍")])]
    )
  }
]
render._withStripped = true



/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			id: moduleId,
/******/ 			loaded: false,
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		__webpack_modules__[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/ 	
/******/ 		// Flag the module as loaded
/******/ 		module.loaded = true;
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = __webpack_modules__;
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/chunk loaded */
/******/ 	!function() {
/******/ 		var deferred = [];
/******/ 		__webpack_require__.O = function(result, chunkIds, fn, priority) {
/******/ 			if(chunkIds) {
/******/ 				priority = priority || 0;
/******/ 				for(var i = deferred.length; i > 0 && deferred[i - 1][2] > priority; i--) deferred[i] = deferred[i - 1];
/******/ 				deferred[i] = [chunkIds, fn, priority];
/******/ 				return;
/******/ 			}
/******/ 			var notFulfilled = Infinity;
/******/ 			for (var i = 0; i < deferred.length; i++) {
/******/ 				var chunkIds = deferred[i][0];
/******/ 				var fn = deferred[i][1];
/******/ 				var priority = deferred[i][2];
/******/ 				var fulfilled = true;
/******/ 				for (var j = 0; j < chunkIds.length; j++) {
/******/ 					if ((priority & 1 === 0 || notFulfilled >= priority) && Object.keys(__webpack_require__.O).every(function(key) { return __webpack_require__.O[key](chunkIds[j]); })) {
/******/ 						chunkIds.splice(j--, 1);
/******/ 					} else {
/******/ 						fulfilled = false;
/******/ 						if(priority < notFulfilled) notFulfilled = priority;
/******/ 					}
/******/ 				}
/******/ 				if(fulfilled) {
/******/ 					deferred.splice(i--, 1)
/******/ 					var r = fn();
/******/ 					if (r !== undefined) result = r;
/******/ 				}
/******/ 			}
/******/ 			return result;
/******/ 		};
/******/ 	}();
/******/ 	
/******/ 	/* webpack/runtime/compat get default export */
/******/ 	!function() {
/******/ 		// getDefaultExport function for compatibility with non-harmony modules
/******/ 		__webpack_require__.n = function(module) {
/******/ 			var getter = module && module.__esModule ?
/******/ 				function() { return module['default']; } :
/******/ 				function() { return module; };
/******/ 			__webpack_require__.d(getter, { a: getter });
/******/ 			return getter;
/******/ 		};
/******/ 	}();
/******/ 	
/******/ 	/* webpack/runtime/define property getters */
/******/ 	!function() {
/******/ 		// define getter functions for harmony exports
/******/ 		__webpack_require__.d = function(exports, definition) {
/******/ 			for(var key in definition) {
/******/ 				if(__webpack_require__.o(definition, key) && !__webpack_require__.o(exports, key)) {
/******/ 					Object.defineProperty(exports, key, { enumerable: true, get: definition[key] });
/******/ 				}
/******/ 			}
/******/ 		};
/******/ 	}();
/******/ 	
/******/ 	/* webpack/runtime/global */
/******/ 	!function() {
/******/ 		__webpack_require__.g = (function() {
/******/ 			if (typeof globalThis === 'object') return globalThis;
/******/ 			try {
/******/ 				return this || new Function('return this')();
/******/ 			} catch (e) {
/******/ 				if (typeof window === 'object') return window;
/******/ 			}
/******/ 		})();
/******/ 	}();
/******/ 	
/******/ 	/* webpack/runtime/hasOwnProperty shorthand */
/******/ 	!function() {
/******/ 		__webpack_require__.o = function(obj, prop) { return Object.prototype.hasOwnProperty.call(obj, prop); }
/******/ 	}();
/******/ 	
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	!function() {
/******/ 		// define __esModule on exports
/******/ 		__webpack_require__.r = function(exports) {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	}();
/******/ 	
/******/ 	/* webpack/runtime/node module decorator */
/******/ 	!function() {
/******/ 		__webpack_require__.nmd = function(module) {
/******/ 			module.paths = [];
/******/ 			if (!module.children) module.children = [];
/******/ 			return module;
/******/ 		};
/******/ 	}();
/******/ 	
/******/ 	/* webpack/runtime/publicPath */
/******/ 	!function() {
/******/ 		var scriptUrl;
/******/ 		if (__webpack_require__.g.importScripts) scriptUrl = __webpack_require__.g.location + "";
/******/ 		var document = __webpack_require__.g.document;
/******/ 		if (!scriptUrl && document) {
/******/ 			if (document.currentScript)
/******/ 				scriptUrl = document.currentScript.src;
/******/ 			if (!scriptUrl) {
/******/ 				var scripts = document.getElementsByTagName("script");
/******/ 				if(scripts.length) {
/******/ 					var i = scripts.length - 1;
/******/ 					while (i > -1 && (!scriptUrl || !/^http(s?):/.test(scriptUrl))) scriptUrl = scripts[i--].src;
/******/ 				}
/******/ 			}
/******/ 		}
/******/ 		// When supporting browsers where an automatic publicPath is not supported you must specify an output.publicPath manually via configuration
/******/ 		// or pass an empty string ("") and set the __webpack_public_path__ variable from your code to use your own logic.
/******/ 		if (!scriptUrl) throw new Error("Automatic publicPath is not supported in this browser");
/******/ 		scriptUrl = scriptUrl.replace(/#.*$/, "").replace(/\?.*$/, "").replace(/\/[^\/]+$/, "/");
/******/ 		__webpack_require__.p = scriptUrl + "../";
/******/ 	}();
/******/ 	
/******/ 	/* webpack/runtime/jsonp chunk loading */
/******/ 	!function() {
/******/ 		// no baseURI
/******/ 		
/******/ 		// object to store loaded and loading chunks
/******/ 		// undefined = chunk not loaded, null = chunk preloaded/prefetched
/******/ 		// [resolve, reject, Promise] = chunk loading, 0 = chunk loaded
/******/ 		var installedChunks = {
/******/ 			"vp-ros-hmci": 0
/******/ 		};
/******/ 		
/******/ 		// no chunk on demand loading
/******/ 		
/******/ 		// no prefetching
/******/ 		
/******/ 		// no preloaded
/******/ 		
/******/ 		// no HMR
/******/ 		
/******/ 		// no HMR manifest
/******/ 		
/******/ 		__webpack_require__.O.j = function(chunkId) { return installedChunks[chunkId] === 0; };
/******/ 		
/******/ 		// install a JSONP callback for chunk loading
/******/ 		var webpackJsonpCallback = function(parentChunkLoadingFunction, data) {
/******/ 			var chunkIds = data[0];
/******/ 			var moreModules = data[1];
/******/ 			var runtime = data[2];
/******/ 			// add "moreModules" to the modules object,
/******/ 			// then flag all "chunkIds" as loaded and fire callback
/******/ 			var moduleId, chunkId, i = 0;
/******/ 			if(chunkIds.some(function(id) { return installedChunks[id] !== 0; })) {
/******/ 				for(moduleId in moreModules) {
/******/ 					if(__webpack_require__.o(moreModules, moduleId)) {
/******/ 						__webpack_require__.m[moduleId] = moreModules[moduleId];
/******/ 					}
/******/ 				}
/******/ 				if(runtime) var result = runtime(__webpack_require__);
/******/ 			}
/******/ 			if(parentChunkLoadingFunction) parentChunkLoadingFunction(data);
/******/ 			for(;i < chunkIds.length; i++) {
/******/ 				chunkId = chunkIds[i];
/******/ 				if(__webpack_require__.o(installedChunks, chunkId) && installedChunks[chunkId]) {
/******/ 					installedChunks[chunkId][0]();
/******/ 				}
/******/ 				installedChunks[chunkId] = 0;
/******/ 			}
/******/ 			return __webpack_require__.O(result);
/******/ 		}
/******/ 		
/******/ 		var chunkLoadingGlobal = self["webpackChunk"] = self["webpackChunk"] || [];
/******/ 		chunkLoadingGlobal.forEach(webpackJsonpCallback.bind(null, 0));
/******/ 		chunkLoadingGlobal.push = webpackJsonpCallback.bind(null, chunkLoadingGlobal.push.bind(chunkLoadingGlobal));
/******/ 	}();
/******/ 	
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module depends on other loaded chunks and execution need to be delayed
/******/ 	var __webpack_exports__ = __webpack_require__.O(undefined, ["vendor-libs"], function() { return __webpack_require__("./web_src/vuepages/pages/ros-hmci/vp-ros-hmci.js"); })
/******/ 	__webpack_exports__ = __webpack_require__.O(__webpack_exports__);
/******/ 	
/******/ })()
;