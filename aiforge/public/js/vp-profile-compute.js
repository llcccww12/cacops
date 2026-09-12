/******/ (function() { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/Icons.vue?vue&type=script&lang=js":
/*!**********************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/Icons.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_0__);

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
      "default": 'dataset'
    },
    isPrivate: {
      type: Boolean,
      "default": false
    },
    modelType: {
      type: Number,
      "default": 0
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=script&lang=js":
/*!*****************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=script&lang=js ***!
  \*****************************************************************************************************************************************************************************************************************/
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
/* harmony default export */ __webpack_exports__["default"] = ({
  name: "Empty",
  data: function data() {
    return {};
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=script&lang=js":
/*!*******************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=script&lang=js ***!
  \*******************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.find */ "./node_modules/core-js/modules/es.array.find.js");
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.string.trim */ "./node_modules/core-js/modules/es.string.trim.js");
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_2__);



//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: 'Filters',
  props: {
    // 搜索相关
    placeholder: {
      type: String,
      "default": ''
    },
    searchBtnText: {
      type: String,
      "default": 'Search'
    },
    // 排序相关
    sortOptions: {
      type: Array,
      required: true,
      "default": function _default() {
        return [];
      }
    },
    defaultSortKey: {
      type: String,
      "default": 'default'
    }
  },
  data: function data() {
    return {
      localKeyword: '',
      currentSortKey: this.defaultSortKey,
      currentSortLabel: '',
      isDropdownActive: false,
      isDropdownClicked: false
    };
  },
  mounted: function mounted() {
    // 添加全局点击事件监听器，用于检测点击外部区域
    document.addEventListener('click', this.handleClickOutside);
  },
  beforeDestroy: function beforeDestroy() {
    // 移除事件监听器，防止内存泄漏
    document.removeEventListener('click', this.handleClickOutside);
  },
  watch: {
    sortOptions: {
      immediate: true,
      handler: function handler(newVal) {
        var _this = this;

        var defaultItem = newVal.find(function (item) {
          return item.key === _this.defaultSortKey;
        });

        if (defaultItem) {
          this.currentSortLabel = defaultItem.label;
        }
      }
    },
    // 监听下拉菜单显示状态变化
    isDropdownActive: function isDropdownActive(newVal) {
      var _this2 = this;

      if (!newVal && this.isDropdownClicked) {
        // 当下拉菜单关闭且曾经被点击过时，激活字体颜色样式
        setTimeout(function () {
          _this2.isDropdownClicked = true; // 保持点击状态
        }, 10);
      }
    }
  },
  methods: {
    handleSearch: function handleSearch() {
      this.$emit('search', {
        keyword: this.localKeyword.trim(),
        order_by: this.currentSortKey
      });
    },
    selectSort: function selectSort(item) {
      this.currentSortKey = item.key;
      this.currentSortLabel = item.label;
      this.handleSearch(); // 自动触发搜索
    },
    onDropdownVisibleChange: function onDropdownVisibleChange(visible) {
      this.isDropdownActive = visible;

      if (visible) {
        this.isDropdownClicked = true;
      }
    },
    handleDropdownClick: function handleDropdownClick(event) {
      // 阻止事件冒泡，避免被handleClickOutside处理
      event.stopPropagation();
      this.isDropdownClicked = true;
    },
    handleClickOutside: function handleClickOutside(event) {
      var _this$$refs$filterOrd;

      // 如果点击的不是filter-order元素或其子元素
      var filterOrderElement = (_this$$refs$filterOrd = this.$refs.filterOrder) === null || _this$$refs$filterOrd === void 0 ? void 0 : _this$$refs$filterOrd.$el;

      if (filterOrderElement && !filterOrderElement.contains(event.target)) {
        this.isDropdownClicked = false;
      }
    },
    // 提供外部重置方法（可选）
    reset: function reset() {
      var _this3 = this;

      this.localKeyword = '';
      this.currentSortKey = this.defaultSortKey;
      var defaultItem = this.sortOptions.find(function (item) {
        return item.key === _this3.defaultSortKey;
      });
      this.currentSortLabel = defaultItem ? defaultItem.label : '';
      this.isDropdownClicked = false;
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Item.vue?vue&type=script&lang=js":
/*!****************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Item.vue?vue&type=script&lang=js ***!
  \****************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var _components_square_Icons_vue__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ~/components/square/Icons.vue */ "./web_src/vuepages/components/square/Icons.vue");


//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "Item",
  props: {
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  components: {
    Icons: _components_square_Icons_vue__WEBPACK_IMPORTED_MODULE_2__["default"]
  },
  data: function data() {
    return {
      isMobile: false
    };
  },
  mounted: function mounted() {
    this.checkScreenSize();
    window.addEventListener('resize', this.checkScreenSize);
  },
  beforeDestroy: function beforeDestroy() {
    window.removeEventListener('resize', this.checkScreenSize);
  },
  methods: {
    checkScreenSize: function checkScreenSize() {
      this.isMobile = window.innerWidth <= 767;
    },
    getItemLink: function getItemLink() {
      var typeToPathMap = {
        aitasktmpl: function aitasktmpl(data) {
          return "/ai_task_tmpl/detail/".concat(data.ID || data.id);
        },
        dataset: function dataset(data) {
          return "/datasets/detail/".concat(data.owner_name, "/").concat(data.name);
        },
        model: function model(data) {
          return "/models/detail/".concat(data.owner_name, "/").concat(data.name);
        }
      };
      var getPath = typeToPathMap[this.data.type];

      if (!getPath) {
        console.warn("Unknown data type: ".concat(this.data.type));
        return;
      }

      var href = getPath(this.data);
      window.open(href, '_blank');
    },
    goRun: function goRun(item) {
      window.open("/cloudbrains/create?tmpl=".concat(item.ID), '_blank');
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/List.vue?vue&type=script&lang=js":
/*!****************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/List.vue?vue&type=script&lang=js ***!
  \****************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Item_vue__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Item.vue */ "./web_src/vuepages/pages/profile/components/Item.vue");
//
//
//
//
//
//
//
//

/* harmony default export */ __webpack_exports__["default"] = ({
  name: "List",
  props: {
    params: {
      type: Array,
      "default": function _default() {
        return [];
      }
    }
  },
  components: {
    Item: _Item_vue__WEBPACK_IMPORTED_MODULE_0__["default"]
  },
  data: function data() {
    return {};
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/compute/index.vue?vue&type=script&lang=js":
/*!**************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/compute/index.vue?vue&type=script&lang=js ***!
  \**************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.join */ "./node_modules/core-js/modules/es.array.join.js");
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_6__);
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
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var _components_List_vue__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! ../components/List.vue */ "./web_src/vuepages/pages/profile/components/List.vue");
/* harmony import */ var _components_Empty_vue__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! ../components/Empty.vue */ "./web_src/vuepages/pages/profile/components/Empty.vue");
/* harmony import */ var _components_Filters_vue__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! ../components/Filters.vue */ "./web_src/vuepages/pages/profile/components/Filters.vue");
/* harmony import */ var _apis_modules_aitasktmpl__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! ~/apis/modules/aitasktmpl */ "./web_src/vuepages/apis/modules/aitasktmpl.js");
/* harmony import */ var element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! element-ui/lib/utils/date-util */ "./node_modules/element-ui/lib/utils/date-util.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _pages_aitasktmpl_tools__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! ~/pages/aitasktmpl/tools */ "./web_src/vuepages/pages/aitasktmpl/tools.js");















function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_13___default()(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  components: {
    List: _components_List_vue__WEBPACK_IMPORTED_MODULE_14__["default"],
    Empty: _components_Empty_vue__WEBPACK_IMPORTED_MODULE_15__["default"],
    Filters: _components_Filters_vue__WEBPACK_IMPORTED_MODULE_16__["default"]
  },
  data: function data() {
    return {
      keyword: '',
      order_by: 'default',
      page: 1,
      page_sizes: [10],
      page_size: 10,
      total: 0,
      currentList: [],
      isLogin: false,
      sortList: [{
        key: 'default',
        label: this.$t('datasets.default')
      }, {
        key: 'recentupdate',
        label: this.$t('datasets.recentupdate')
      }, {
        key: 'newest',
        label: this.$t('datasets.newest')
      }, {
        key: 'collections',
        label: this.$t('datasets.moststars')
      }, {
        key: 'usecount',
        label: this.$t('taskTmplObj.runTimes')
      }, {
        key: 'name_asc',
        label: this.$t('datasets.alphabetasc')
      }, {
        key: 'name',
        label: this.$t('datasets.alphabetdesc')
      }],
      loading: false
    };
  },
  methods: {
    getList: function getList() {
      var _window$__INITIAL_STA,
          _this = this;

      this.loading = true;
      var params = {
        q: this.keyword || undefined,
        order_by: this.order_by,
        page: this.page,
        page_size: this.page_size,
        owner_name: ((_window$__INITIAL_STA = window.__INITIAL_STATE__) === null || _window$__INITIAL_STA === void 0 ? void 0 : _window$__INITIAL_STA.ownerName) || 'default_user'
      };
      var request = this.isLogin ? (0,_apis_modules_aitasktmpl__WEBPACK_IMPORTED_MODULE_17__.getProfileAITaskTemplate)(params) : (0,_apis_modules_aitasktmpl__WEBPACK_IMPORTED_MODULE_17__.getProfileAITaskTemplatePublic)(params);
      request.then(function (res) {
        _this.loading = false;
        res = res.data;

        if (res.code == 0) {
          var list = res.data.Templates || [];
          _this.currentList = list.map(function (item) {
            var datasets = item.DatasetLists || [];
            var datasetsStr = datasets.map(function (item) {
              return "".concat(item.OwnerName ? item.OwnerName + '/' : '').concat(item.DatasetAlias || item.DatasetName);
            }).join(', ');
            var models = item.ModelLists || [];
            var modelsStr = models.map(function (item) {
              return "".concat(item.OwnerName ? item.OwnerName + '/' : '').concat(item.ModelAlias || item.ModelName);
            }).join(', ');
            return _objectSpread(_objectSpread({}, item), {}, {
              alias: item.Name,
              updated_time: _this.formatTime(item.UpdatedUnix),
              recommend: item.Recommend,
              is_collected: item.IsCollected,
              num_stars: item.NumCollections,
              use_count: item.UseCount,
              tags: item.Tags,
              DatasetsStr: datasetsStr,
              ModelsStr: modelsStr,
              JobTypeStr: (0,_utils__WEBPACK_IMPORTED_MODULE_19__.getListValueWithKey)(_pages_aitasktmpl_tools__WEBPACK_IMPORTED_MODULE_20__.TmplTaskTypes, item.JobType),
              ComputeSourceStr: (0,_utils__WEBPACK_IMPORTED_MODULE_19__.getListValueWithKey)(_pages_aitasktmpl_tools__WEBPACK_IMPORTED_MODULE_20__.TmplComputerResouces, item.ComputeSource),
              type: 'aitasktmpl',
              is_private: item.IsPrivate
            });
          });
          _this.total = res.data.Total;
        } else {
          _this.resetPagination();
        }
      })["catch"](function (err) {
        console.error(err);
        _this.loading = false;
        _this.currentList = [];

        _this.resetPagination();
      });
    },
    resetPagination: function resetPagination() {
      this.keyword = '';
      this.order_by = 'default';
      this.page = 1;
      this.page_size = 10;
      this.total = 0;
    },
    formatTime: function formatTime(time) {
      return (0,element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_18__.formatDate)(new Date(time * 1000), 'yyyy-MM-dd');
    },
    onFilterSearch: function onFilterSearch(_ref) {
      var keyword = _ref.keyword,
          order_by = _ref.order_by;
      this.keyword = keyword;
      this.order_by = order_by;
      this.page = 1;
      this.getList();
    },
    currentChange: function currentChange(page) {
      this.page = page;
      this.getList();
    },
    sizeChange: function sizeChange(pageSize) {
      this.page_size = pageSize;
      this.page = 1;
      this.getList();
    }
  },
  mounted: function mounted() {
    this.isLogin = !!document.querySelector('meta[name="_uid"]');
    this.getList();
  }
});

/***/ }),

/***/ "./web_src/vuepages/apis/modules/aitasktmpl.js":
/*!*****************************************************!*\
  !*** ./web_src/vuepages/apis/modules/aitasktmpl.js ***!
  \*****************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   createAiTaskTmpl: function() { return /* binding */ createAiTaskTmpl; },
/* harmony export */   deleteAiTaskTmpl: function() { return /* binding */ deleteAiTaskTmpl; },
/* harmony export */   deleteCollectAiTaskTmpl: function() { return /* binding */ deleteCollectAiTaskTmpl; },
/* harmony export */   deleteRecommendAiTaskTmpl: function() { return /* binding */ deleteRecommendAiTaskTmpl; },
/* harmony export */   doSearchAiTaskTmplConds: function() { return /* binding */ doSearchAiTaskTmplConds; },
/* harmony export */   editAiTaskTmpl: function() { return /* binding */ editAiTaskTmpl; },
/* harmony export */   getAiTaskTmpl: function() { return /* binding */ getAiTaskTmpl; },
/* harmony export */   getAiTaskTmplList: function() { return /* binding */ getAiTaskTmplList; },
/* harmony export */   getDefaultAiTaskTmplConds: function() { return /* binding */ getDefaultAiTaskTmplConds; },
/* harmony export */   getProfileAITaskTemplate: function() { return /* binding */ getProfileAITaskTemplate; },
/* harmony export */   getProfileAITaskTemplatePublic: function() { return /* binding */ getProfileAITaskTemplatePublic; },
/* harmony export */   putCollectAiTaskTmpl: function() { return /* binding */ putCollectAiTaskTmpl; },
/* harmony export */   putRecommendAiTaskTmpl: function() { return /* binding */ putRecommendAiTaskTmpl; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.object.define-properties */ "./node_modules/core-js/modules/es.object.define-properties.js");
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.object.define-property */ "./node_modules/core-js/modules/es.object.define-property.js");
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptor */ "./node_modules/core-js/modules/es.object.get-own-property-descriptor.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptors */ "./node_modules/core-js/modules/es.object.get-own-property-descriptors.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var _service__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! ../service */ "./web_src/vuepages/apis/service.js");
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");











function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_9___default()(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }



/* 查询计算任务模板列表 */
// type 
// - public 公开资源-公开模板
// - recommend 公开资源-平台推荐模板
// - created 我的工作台-我创建的
// - collected 我的工作台-我收藏的
// - admin 管理后台-计算任务模板列表
// - repo 项目主页-引用过改项目的所有模板
// q - 搜索关键词
// tags - 标签，如果有多个值用|分开
// recommend  - all 所有，only只展示推荐
// job_type - 任务类型
// compute_source - 计算资源
// dataset_id	- 数据集id
// model_id	- 模型id
// repo_id	- 项目id
// order_by - default 默认,newest 最新创建, recentupdate 最近更新,  collections 收藏数量, usecount 运行次数, name 名称降序, name_asc 名称升序 
// page - 页码 从1 开始
// page_size - 每页数量

var getAiTaskTmplList = function getAiTaskTmplList(params) {
  var url = "";

  switch (params.type) {
    case 'created':
      url = "/api/v1/ai_task_template/list/created";
      break;

    case 'collected':
      url = "/api/v1/ai_task_template/list/collected";
      break;

    case 'admin':
      url = "/api/v1/admin/ai_task_template/list";
      break;

    case 'recommend':
      url = "/api/v1/ai_task_template/list/public?recommend=only";
      break;

    case 'repo':
      url = "/api/v1/ai_task_template/list/all?repo_id=".concat(params.repo);
      break;

    case 'public':
    default:
      url = "/api/v1/ai_task_template/list/public";
      break;
  }

  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: url,
    method: 'get',
    params: _objectSpread({}, params)
  });
};
/* 创建计算任务模板 */

var createAiTaskTmpl = function createAiTaskTmpl(data) {
  var params = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : {};
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: "/api/v1/ai_task_template/create",
    method: 'post',
    params: {},
    data: _objectSpread({}, data)
  });
};
/* 编辑计算任务模板 */

var editAiTaskTmpl = function editAiTaskTmpl(data) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: "/api/v1/ai_task_template/edit",
    method: 'post',
    params: {
      id: data.id
    },
    data: _objectSpread({}, data)
  });
};
/* 查询计算任务模板详情 */

var getAiTaskTmpl = function getAiTaskTmpl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: "/api/v1/ai_task_template",
    method: 'get',
    params: {
      id: params.id
    }
  });
};
/* 删除计算任务模板 */

var deleteAiTaskTmpl = function deleteAiTaskTmpl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: "/api/v1/ai_task_template",
    method: 'delete',
    params: {
      id: params.id
    }
  });
};
/* 收藏计算任务模板 */

var putCollectAiTaskTmpl = function putCollectAiTaskTmpl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: "/api/v1/ai_task_template/collect",
    method: 'put',
    params: {
      id: params.id
    }
  });
};
/* 取消收藏计算任务模板 */

var deleteCollectAiTaskTmpl = function deleteCollectAiTaskTmpl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: "/api/v1/ai_task_template/collect",
    method: 'delete',
    params: {
      id: params.id
    }
  });
};
/* 推荐计算任务模板 */

var putRecommendAiTaskTmpl = function putRecommendAiTaskTmpl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: "/api/v1/admin/ai_task_template/recommend",
    method: 'put',
    params: {
      id: params.id
    }
  });
};
/* 取消推荐计算任务模板 */

var deleteRecommendAiTaskTmpl = function deleteRecommendAiTaskTmpl(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: "/api/v1/admin/ai_task_template/recommend",
    method: 'delete',
    params: {
      id: params.id
    }
  });
};
/* 搜索模型、数据集、项目 */
// type -model|dataset|repo
// q - keyword
// page
// pageSize

var doSearchAiTaskTmplConds = function doSearchAiTaskTmplConds(params) {
  var tableName = '';

  if (params.type == 'model') {
    tableName = 'model';
  }

  if (params.type == 'dataset') {
    tableName = 'dataset';
  }

  if (params.type == 'repo') {
    tableName = 'repository';
  }

  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: "/all/dosearch/",
    method: 'get',
    params: {
      TableName: tableName,
      Key: params.q,
      Page: params.page,
      PageSize: params.pageSize,
      OnlyReturnNum: false,
      OnlySearchLabel: false,
      WebTotal: 0,
      PrivateTotal: 0,
      language: _langs__WEBPACK_IMPORTED_MODULE_11__.lang
    }
  });
};
/* 搜索默认的模型、数据集、项目 */
// key -配置文件路径

var getDefaultAiTaskTmplConds = function getDefaultAiTaskTmplConds(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: "/api/v1/dync_config",
    method: 'get',
    params: params,
    data: {}
  });
}; //获取个人信息页面中的计算任务模板列表

var getProfileAITaskTemplate = function getProfileAITaskTemplate(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: '/api/v1/ai_task_template/list/accessible',
    method: "get",
    params: params
  });
}; //获取个人信息页面中的计算任务模板列表（未登录态使用，仅返回公开模板）

var getProfileAITaskTemplatePublic = function getProfileAITaskTemplatePublic(params) {
  return (0,_service__WEBPACK_IMPORTED_MODULE_10__["default"])({
    url: '/api/v1/ai_task_template/list/public',
    method: "get",
    params: params
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

/***/ "./web_src/vuepages/pages/aitasktmpl/tools.js":
/*!****************************************************!*\
  !*** ./web_src/vuepages/pages/aitasktmpl/tools.js ***!
  \****************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   TaskTmplTools: function() { return /* binding */ TaskTmplTools; },
/* harmony export */   TmplComputerResouces: function() { return /* binding */ TmplComputerResouces; },
/* harmony export */   TmplTaskTypes: function() { return /* binding */ TmplTaskTypes; },
/* harmony export */   getGradientColor: function() { return /* binding */ getGradientColor; }
/* harmony export */ });
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.symbol.description */ "./node_modules/core-js/modules/es.symbol.description.js");
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.filter */ "./node_modules/core-js/modules/es.array.filter.js");
/* harmony import */ var core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_filter__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @babel/runtime/helpers/toConsumableArray */ "./node_modules/@babel/runtime/helpers/toConsumableArray.js");
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @babel/runtime/helpers/classCallCheck */ "./node_modules/@babel/runtime/helpers/classCallCheck.js");
/* harmony import */ var _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @babel/runtime/helpers/createClass */ "./node_modules/@babel/runtime/helpers/createClass.js");
/* harmony import */ var _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");
/* harmony import */ var _pages_cloudbrain_configs__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ~/pages/cloudbrain/configs */ "./web_src/vuepages/pages/cloudbrain/configs.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! spark-md5 */ "./node_modules/spark-md5/spark-md5.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(spark_md5__WEBPACK_IMPORTED_MODULE_10__);











var TaskTmplTools = /*#__PURE__*/function () {
  function TaskTmplTools() {
    _babel_runtime_helpers_classCallCheck__WEBPACK_IMPORTED_MODULE_6___default()(this, TaskTmplTools);
  }

  _babel_runtime_helpers_createClass__WEBPACK_IMPORTED_MODULE_7___default()(TaskTmplTools, [{
    key: "transformDataTmplToForm",
    value: function transformDataTmplToForm(obj) {
      var runParameters = JSON.parse(obj.Parameters || '[]').map(function (item) {
        return {
          label: item.Label || item.label,
          value: item.Value || item.value
        };
      });
      var models = (obj.ModelLists || []).filter(function (item) {
        return !item.IsDeleted;
      }).map(function (item) {
        return {
          id: item.ModelID,
          name: item.ModelName,
          owner_name: item.OwnerName,
          alias: item.ModelAlias
        };
      });
      var datasets = (obj.DatasetLists || []).filter(function (item) {
        return !item.IsDeleted;
      }).map(function (item) {
        return {
          id: item.DatasetID,
          name: item.DatasetName,
          owner_name: item.OwnerName,
          alias: item.DatasetAlias
        };
      });
      var data = {
        id: obj.ID,
        ownerId: obj.OwnerId,
        name: obj.Name || '',
        descr: obj.Description || '',
        tags: obj.Tags,
        isPrivate: obj.IsPrivate,
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
          image_id: obj.ImageID || '',
          image_name: obj.ImageName || '',
          image_url: obj.ImageUrl || ''
        },
        model: models,
        dataset: datasets,
        repoID: obj.RepoID,
        repoOwnerName: obj.RepoOwnerName,
        repoName: obj.RepoName,
        branchName: obj.BranchName || '',
        bootFile: obj.BootFile || '',
        runParameters: runParameters
      };
      return data;
    }
  }, {
    key: "transformDataFormToTmpl",
    value: function transformDataFormToTmpl(obj) {
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
        Tags: obj.tags,
        IsPrivate: obj.isPrivate,
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
        ImageID: obj.image.image_id || '',
        ImageName: obj.image.image_name || '',
        ImageUrl: obj.image.image_url,
        RepoOwnerName: obj.repoOwnerName,
        RepoName: obj.repoName,
        BranchName: obj.branchName,
        BootFile: obj.bootFile,
        ModelIDs: modelList.map(function (item) {
          return item.ID;
        }),
        DatasetIDs: datasetList.map(function (item) {
          return item.ID;
        }),
        Parameters: JSON.stringify(runParameterList)
      };
      return data;
    }
  }, {
    key: "transformDataTaskToForm",
    value: function transformDataTaskToForm(obj) {
      var _obj$parameters;

      var runParameters = (((_obj$parameters = obj.parameters) === null || _obj$parameters === void 0 ? void 0 : _obj$parameters.parameter) || []).map(function (item) {
        return {
          label: item.Label || item.label,
          value: item.Value || item.value
        };
      });
      var models = (obj.pretrain_model_list || []).filter(function (item) {
        return !item.is_delete;
      }).map(function (item) {
        return {
          id: item.id,
          name: item.name,
          owner_name: item.owner_name,
          alias: item.alias
        };
      });
      var datasets = (obj.dataset_list || []).filter(function (item) {
        return !item.IsDeleted;
      }).map(function (item) {
        return {
          id: item.uuid,
          name: item.dataset_name,
          owner_name: item.owner_name,
          alias: item.dataset_alias
        };
      });
      var spec = obj.spec || {};
      var data = {
        taskId: obj.id,
        name: obj.display_job_name || '',
        descr: obj.description || '',
        tags: [],
        isPrivate: false,
        taskType: obj.job_type || '',
        cluster: obj.cluster || '',
        computeResource: obj.compute_source || '',
        networkType: obj.has_internet == 1 ? 'no_internet' : obj.has_internet == 2 ? 'has_internet' : 'has_internet',
        visualizeRequired: !!obj.visualize_required,
        spec: '',
        acc_cards_num: spec.acc_cards_num || 0,
        acc_card_type: spec.acc_card_type || '',
        cpu_cores: spec.cpu_cores || 0,
        mem_gi_b: spec.mem_gi_b || 0,
        gpu_mem_gi_b: spec.gpu_mem_gi_b || 0,
        share_mem_gi_b: spec.share_mem_gi_b || 0,
        image: {
          image_id: obj.image_id || '',
          image_name: obj.image_name || '',
          image_url: obj.image_url || ''
        },
        model: models,
        dataset: datasets,
        repoID: obj.repo_id,
        repoOwnerName: obj.repo_owner_name,
        repoName: obj.repo_name,
        branchName: obj.branch_name || '',
        bootFile: obj.boot_file || '',
        runParameters: runParameters
      };
      return data;
    }
  }]);

  return TaskTmplTools;
}();
var TmplTaskTypes = [{
  k: 'DEBUG',
  v: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.tabTitDebug'),
  cluster: 'C2Net',
  computerResouce: 'NPU',
  desc: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.debugTaskDesc'),
  descLong: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.debugTaskDescLong')
}, {
  k: 'TRAIN',
  v: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.tabTitTrain'),
  cluster: 'C2Net',
  computerResouce: 'NPU',
  desc: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.trainTaskDesc'),
  descLong: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.trainTaskDescLong')
}, {
  k: 'ONLINEINFERENCE',
  v: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.tabTitOnlineInference'),
  cluster: 'C2Net',
  computerResouce: 'GPU',
  desc: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.onlineinferTaskDesc'),
  descLong: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.onlineinferTaskDescLong')
}, {
  k: 'GENERAL',
  v: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.tabTitGeneral'),
  cluster: 'C2Net',
  computerResouce: 'GPU',
  desc: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.generalTaskDesc'),
  descLong: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.generalTaskDescLong')
}, {
  k: 'HPC',
  v: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('superComputeTask'),
  cluster: 'C2Net',
  computerResouce: 'CPU',
  desc: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.inferenceTaskDesc'),
  descLong: _langs__WEBPACK_IMPORTED_MODULE_8__.i18n.t('cloudbrainObj.inferenceTaskDescLong')
}];
var TmplComputerResouces = _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_5___default()(_pages_cloudbrain_configs__WEBPACK_IMPORTED_MODULE_9__.COMPUTER_RESOURCES_TITLE);
var GradientColors = [['rgba(178,195,255,0.85)', 'rgba(160,246,246,1)'], ['rgba(189,231,255,0.85)', 'rgba(200,202,255,1)'], ['rgba(255,237,189,0.85)', 'rgba(214,255,200,1)'], ['rgba(217,186,255,0.85)', 'rgba(255,233,201,1)'], ['rgba(200,255,232,1)', 'rgba(189,235,255,0.85)'], ['rgba(242,255,189,0.85)', 'rgba(200,205,255,1)'], ['rgba(203,178,255,0.85)', 'rgba(203,246,160,1)'], ['rgba(255,213,189,0.85)', 'rgba(200,255,232,1)'], ['rgba(255,189,237,0.85)', 'rgba(200,229,255,1)'], ['rgba(255,186,186,0.85)', 'rgba(255,233,201,1)'], ['rgba(255,250,200,1)', 'rgba(189,215,255,0.85)'], ['rgba(189,255,196,0.85)', 'rgba(200,255,239,1)']];
var getGradientColor = function getGradientColor(name) {
  var nameHash = spark_md5__WEBPACK_IMPORTED_MODULE_10___default().hash(name);
  var colourIndex = (nameHash.charCodeAt(0) + nameHash.charCodeAt(1) + nameHash.charCodeAt(2) + nameHash.charCodeAt(3)) % 12;
  return GradientColors[colourIndex];
};

/***/ }),

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

/***/ }),

/***/ "./web_src/vuepages/pages/profile/compute/vp-profile-compute.js":
/*!**********************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/compute/vp-profile-compute.js ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var vue__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! vue */ "./node_modules/vue/dist/vue.esm.js");
/* harmony import */ var element_ui__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! element-ui */ "./node_modules/element-ui/lib/element-ui.common.js");
/* harmony import */ var element_ui__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(element_ui__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var element_ui_lib_theme_chalk_index_css__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! element-ui/lib/theme-chalk/index.css */ "./node_modules/element-ui/lib/theme-chalk/index.css");
/* harmony import */ var element_ui_lib_locale_lang_en__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! element-ui/lib/locale/lang/en */ "./node_modules/element-ui/lib/locale/lang/en.js");
/* harmony import */ var element_ui_lib_locale_lang_zh_CN__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! element-ui/lib/locale/lang/zh-CN */ "./node_modules/element-ui/lib/locale/lang/zh-CN.js");
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");
/* harmony import */ var _index_vue__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ./index.vue */ "./web_src/vuepages/pages/profile/compute/index.vue");







vue__WEBPACK_IMPORTED_MODULE_6__["default"].use((element_ui__WEBPACK_IMPORTED_MODULE_0___default()), {
  locale: _langs__WEBPACK_IMPORTED_MODULE_4__.lang === 'zh-CN' ? element_ui_lib_locale_lang_zh_CN__WEBPACK_IMPORTED_MODULE_3__["default"] : element_ui_lib_locale_lang_en__WEBPACK_IMPORTED_MODULE_2__["default"],
  size: 'small'
});
new vue__WEBPACK_IMPORTED_MODULE_6__["default"]({
  i18n: _langs__WEBPACK_IMPORTED_MODULE_4__.i18n,
  render: function render(h) {
    return h(_index_vue__WEBPACK_IMPORTED_MODULE_5__["default"]);
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

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=style&index=0&id=64a91ffc&scoped=true&lang=less":
/*!***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=style&index=0&id=64a91ffc&scoped=true&lang=less ***!
  \***********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=style&index=0&id=56015720&scoped=true&lang=less":
/*!*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=style&index=0&id=56015720&scoped=true&lang=less ***!
  \*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Item.vue?vue&type=style&index=0&id=bbde0424&scoped=true&lang=less":
/*!**********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Item.vue?vue&type=style&index=0&id=bbde0424&scoped=true&lang=less ***!
  \**********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/compute/index.vue?vue&type=style&index=0&id=2bf59048&scoped=true&lang=less":
/*!********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/compute/index.vue?vue&type=style&index=0&id=2bf59048&scoped=true&lang=less ***!
  \********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./web_src/vuepages/components/square/Icons.vue":
/*!******************************************************!*\
  !*** ./web_src/vuepages/components/square/Icons.vue ***!
  \******************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Icons_vue_vue_type_template_id_0b139034__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Icons.vue?vue&type=template&id=0b139034 */ "./web_src/vuepages/components/square/Icons.vue?vue&type=template&id=0b139034");
/* harmony import */ var _Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Icons.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/square/Icons.vue?vue&type=script&lang=js");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! !../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");





/* normalize component */
;
var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
  _Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Icons_vue_vue_type_template_id_0b139034__WEBPACK_IMPORTED_MODULE_0__.render,
  _Icons_vue_vue_type_template_id_0b139034__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  null,
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/square/Icons.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Empty.vue":
/*!*************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Empty.vue ***!
  \*************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Empty_vue_vue_type_template_id_64a91ffc_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Empty.vue?vue&type=template&id=64a91ffc&scoped=true */ "./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=template&id=64a91ffc&scoped=true");
/* harmony import */ var _Empty_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Empty.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=script&lang=js");
/* harmony import */ var _Empty_vue_vue_type_style_index_0_id_64a91ffc_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Empty.vue?vue&type=style&index=0&id=64a91ffc&scoped=true&lang=less */ "./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=style&index=0&id=64a91ffc&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Empty_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Empty_vue_vue_type_template_id_64a91ffc_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Empty_vue_vue_type_template_id_64a91ffc_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "64a91ffc",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/profile/components/Empty.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Filters.vue":
/*!***************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Filters.vue ***!
  \***************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Filters_vue_vue_type_template_id_56015720_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Filters.vue?vue&type=template&id=56015720&scoped=true */ "./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=template&id=56015720&scoped=true");
/* harmony import */ var _Filters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Filters.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=script&lang=js");
/* harmony import */ var _Filters_vue_vue_type_style_index_0_id_56015720_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Filters.vue?vue&type=style&index=0&id=56015720&scoped=true&lang=less */ "./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=style&index=0&id=56015720&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Filters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Filters_vue_vue_type_template_id_56015720_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Filters_vue_vue_type_template_id_56015720_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "56015720",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/profile/components/Filters.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Item.vue":
/*!************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Item.vue ***!
  \************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Item_vue_vue_type_template_id_bbde0424_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Item.vue?vue&type=template&id=bbde0424&scoped=true */ "./web_src/vuepages/pages/profile/components/Item.vue?vue&type=template&id=bbde0424&scoped=true");
/* harmony import */ var _Item_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Item.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/profile/components/Item.vue?vue&type=script&lang=js");
/* harmony import */ var _Item_vue_vue_type_style_index_0_id_bbde0424_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Item.vue?vue&type=style&index=0&id=bbde0424&scoped=true&lang=less */ "./web_src/vuepages/pages/profile/components/Item.vue?vue&type=style&index=0&id=bbde0424&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Item_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Item_vue_vue_type_template_id_bbde0424_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Item_vue_vue_type_template_id_bbde0424_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "bbde0424",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/profile/components/Item.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/List.vue":
/*!************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/List.vue ***!
  \************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _List_vue_vue_type_template_id_a7d3140e__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./List.vue?vue&type=template&id=a7d3140e */ "./web_src/vuepages/pages/profile/components/List.vue?vue&type=template&id=a7d3140e");
/* harmony import */ var _List_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./List.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/profile/components/List.vue?vue&type=script&lang=js");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");





/* normalize component */
;
var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
  _List_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _List_vue_vue_type_template_id_a7d3140e__WEBPACK_IMPORTED_MODULE_0__.render,
  _List_vue_vue_type_template_id_a7d3140e__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  null,
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/profile/components/List.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/profile/compute/index.vue":
/*!**********************************************************!*\
  !*** ./web_src/vuepages/pages/profile/compute/index.vue ***!
  \**********************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _index_vue_vue_type_template_id_2bf59048_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./index.vue?vue&type=template&id=2bf59048&scoped=true */ "./web_src/vuepages/pages/profile/compute/index.vue?vue&type=template&id=2bf59048&scoped=true");
/* harmony import */ var _index_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./index.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/profile/compute/index.vue?vue&type=script&lang=js");
/* harmony import */ var _index_vue_vue_type_style_index_0_id_2bf59048_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./index.vue?vue&type=style&index=0&id=2bf59048&scoped=true&lang=less */ "./web_src/vuepages/pages/profile/compute/index.vue?vue&type=style&index=0&id=2bf59048&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _index_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _index_vue_vue_type_template_id_2bf59048_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _index_vue_vue_type_template_id_2bf59048_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "2bf59048",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/profile/compute/index.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/square/Icons.vue?vue&type=script&lang=js":
/*!******************************************************************************!*\
  !*** ./web_src/vuepages/components/square/Icons.vue?vue&type=script&lang=js ***!
  \******************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Icons.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/Icons.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=script&lang=js":
/*!*************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=script&lang=js ***!
  \*************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Empty_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Empty.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Empty_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=script&lang=js":
/*!***************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=script&lang=js ***!
  \***************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Filters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Filters.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Filters_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Item.vue?vue&type=script&lang=js":
/*!************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Item.vue?vue&type=script&lang=js ***!
  \************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Item_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Item.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Item.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Item_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/List.vue?vue&type=script&lang=js":
/*!************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/List.vue?vue&type=script&lang=js ***!
  \************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_List_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./List.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/List.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_List_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/profile/compute/index.vue?vue&type=script&lang=js":
/*!**********************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/compute/index.vue?vue&type=script&lang=js ***!
  \**********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./index.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/compute/index.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=style&index=0&id=64a91ffc&scoped=true&lang=less":
/*!**********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=style&index=0&id=64a91ffc&scoped=true&lang=less ***!
  \**********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Empty_vue_vue_type_style_index_0_id_64a91ffc_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Empty.vue?vue&type=style&index=0&id=64a91ffc&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=style&index=0&id=64a91ffc&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=style&index=0&id=56015720&scoped=true&lang=less":
/*!************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=style&index=0&id=56015720&scoped=true&lang=less ***!
  \************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Filters_vue_vue_type_style_index_0_id_56015720_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Filters.vue?vue&type=style&index=0&id=56015720&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=style&index=0&id=56015720&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Item.vue?vue&type=style&index=0&id=bbde0424&scoped=true&lang=less":
/*!*********************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Item.vue?vue&type=style&index=0&id=bbde0424&scoped=true&lang=less ***!
  \*********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Item_vue_vue_type_style_index_0_id_bbde0424_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Item.vue?vue&type=style&index=0&id=bbde0424&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Item.vue?vue&type=style&index=0&id=bbde0424&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/profile/compute/index.vue?vue&type=style&index=0&id=2bf59048&scoped=true&lang=less":
/*!*******************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/compute/index.vue?vue&type=style&index=0&id=2bf59048&scoped=true&lang=less ***!
  \*******************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_style_index_0_id_2bf59048_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./index.vue?vue&type=style&index=0&id=2bf59048&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/compute/index.vue?vue&type=style&index=0&id=2bf59048&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/square/Icons.vue?vue&type=template&id=0b139034":
/*!************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/Icons.vue?vue&type=template&id=0b139034 ***!
  \************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_template_id_0b139034__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_template_id_0b139034__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_template_id_0b139034__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Icons.vue?vue&type=template&id=0b139034 */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/Icons.vue?vue&type=template&id=0b139034");


/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=template&id=64a91ffc&scoped=true":
/*!*******************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=template&id=64a91ffc&scoped=true ***!
  \*******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Empty_vue_vue_type_template_id_64a91ffc_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Empty_vue_vue_type_template_id_64a91ffc_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Empty_vue_vue_type_template_id_64a91ffc_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Empty.vue?vue&type=template&id=64a91ffc&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=template&id=64a91ffc&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=template&id=56015720&scoped=true":
/*!*********************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=template&id=56015720&scoped=true ***!
  \*********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Filters_vue_vue_type_template_id_56015720_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Filters_vue_vue_type_template_id_56015720_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Filters_vue_vue_type_template_id_56015720_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Filters.vue?vue&type=template&id=56015720&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=template&id=56015720&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/Item.vue?vue&type=template&id=bbde0424&scoped=true":
/*!******************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/Item.vue?vue&type=template&id=bbde0424&scoped=true ***!
  \******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Item_vue_vue_type_template_id_bbde0424_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Item_vue_vue_type_template_id_bbde0424_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Item_vue_vue_type_template_id_bbde0424_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Item.vue?vue&type=template&id=bbde0424&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Item.vue?vue&type=template&id=bbde0424&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/profile/components/List.vue?vue&type=template&id=a7d3140e":
/*!******************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/components/List.vue?vue&type=template&id=a7d3140e ***!
  \******************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_List_vue_vue_type_template_id_a7d3140e__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_List_vue_vue_type_template_id_a7d3140e__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_List_vue_vue_type_template_id_a7d3140e__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./List.vue?vue&type=template&id=a7d3140e */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/List.vue?vue&type=template&id=a7d3140e");


/***/ }),

/***/ "./web_src/vuepages/pages/profile/compute/index.vue?vue&type=template&id=2bf59048&scoped=true":
/*!****************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/profile/compute/index.vue?vue&type=template&id=2bf59048&scoped=true ***!
  \****************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_template_id_2bf59048_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_template_id_2bf59048_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_template_id_2bf59048_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./index.vue?vue&type=template&id=2bf59048&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/compute/index.vue?vue&type=template&id=2bf59048&scoped=true");


/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/Icons.vue?vue&type=template&id=0b139034":
/*!***************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/Icons.vue?vue&type=template&id=0b139034 ***!
  \***************************************************************************************************************************************************************************************************************************/
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
  return _vm.type == "dataset" && !_vm.isPrivate
    ? _c(
        "svg",
        {
          attrs: {
            xmlns: "http://www.w3.org/2000/svg",
            viewBox: "0 0 1024 1024",
            width: "30",
            height: "30"
          }
        },
        [
          _c("defs"),
          _c("g", [
            _c("path", {
              attrs: {
                d:
                  "M460.8 840.7552a726.7328 726.7328 0 0 1-76.8 4.0448c-187.5456 0-307.2-68.2496-307.2-115.2v-61.184C145.7152 713.984 255.5904 742.4 384 742.4c26.4704 0 52.1216-1.28 76.8-3.584v-77.2608a726.7328 726.7328 0 0 1-76.8 4.0448c-187.5456 0-307.2-68.2496-307.2-115.2v-61.44c68.9152 45.824 178.7904 74.24 307.2 74.24 26.4704 0 52.1216-1.28 76.8-3.584V482.3552A726.7328 726.7328 0 0 1 384 486.4c-187.5456 0-307.2-68.2496-307.2-115.2v-61.44C145.7152 355.584 255.5392 384 384 384s238.2848-28.416 307.2-74.0352V358.4h76.8V192C768 82.5344 602.9312 0 384 0S0 82.5344 0 192v537.6C0 839.0656 165.0688 921.6 384 921.6c26.4704 0 52.1216-1.2288 76.8-3.5328zM384 76.8c187.5456 0 307.2 68.2496 307.2 115.2S571.5456 307.2 384 307.2 76.8 238.9504 76.8 192 196.4544 76.8 384 76.8z",
                fill: "#888",
                "p-id": "3601"
              }
            }),
            _c("path", {
              attrs: {
                d:
                  "M768 409.6c-145.92 0-256 55.04-256 128v358.4c0 72.96 110.08 128 256 128s256-55.04 256-128v-358.4c0-72.96-110.08-128-256-128z m0 324.2496c-125.0304 0-204.8-45.4656-204.8-76.8v-40.96c46.08 30.4128 119.1936 49.3568 204.8 49.3568s158.72-18.944 204.8-49.3568v40.96c0 31.3344-79.7696 76.8-204.8 76.8z m-204.8 1.8944c46.08 30.3616 119.1936 49.3056 204.8 49.3056s158.72-18.944 204.8-49.3056v40.96c0 31.2832-79.7696 76.8-204.8 76.8s-204.8-45.5168-204.8-76.8zM768 460.8c125.0304 0 204.8 45.4656 204.8 76.8s-79.7696 76.8-204.8 76.8-204.8-45.4656-204.8-76.8 79.7696-76.8 204.8-76.8z m0 512c-125.0304 0-204.8-45.4656-204.8-76.8v-40.96c46.08 30.4128 119.1936 49.3568 204.8 49.3568s158.72-18.944 204.8-49.3568v40.96c0 31.3344-79.7696 76.8-204.8 76.8z",
                fill: "#888",
                "p-id": "3602"
              }
            })
          ])
        ]
      )
    : _vm.type == "dataset" && _vm.isPrivate
    ? _c(
        "svg",
        {
          attrs: {
            xmlns: "http://www.w3.org/2000/svg",
            viewBox: "0 0 200 200",
            width: "30",
            height: "30",
            fill: "none"
          }
        },
        [
          _c("defs"),
          _vm._v(" "),
          _c("g", [
            _c("g", [
              _c("path", {
                attrs: {
                  id: "path-imb4gmi8ix6s1yc-UmGY9TvFNiqwKS",
                  d:
                    "M185.76 127.37L135.02 127.37C134.88 127.37 134.75 127.36 134.61 127.33C134.48 127.3 134.35 127.27 134.23 127.21C134.11 127.16 133.99 127.1 133.87 127.03C133.76 126.95 133.66 126.87 133.56 126.77C133.47 126.68 133.38 126.57 133.31 126.46C133.23 126.35 133.17 126.23 133.12 126.11C133.07 125.98 133.03 125.85 133 125.72C132.98 125.59 132.96 125.46 132.96 125.32L132.96 113.16C132.96 101.98 141.76 92.55 152.99 92.27C153.67 92.25 154.35 92.27 155.03 92.32C155.71 92.37 156.39 92.46 157.06 92.58C157.74 92.69 158.4 92.85 159.06 93.03C159.72 93.21 160.36 93.43 161 93.68C161.64 93.93 162.26 94.21 162.87 94.52C163.47 94.83 164.06 95.17 164.64 95.54C165.21 95.91 165.77 96.31 166.3 96.73C166.83 97.16 167.35 97.61 167.84 98.08C168.32 98.56 168.79 99.06 169.23 99.58C169.67 100.1 170.08 100.65 170.46 101.21C170.85 101.77 171.2 102.35 171.53 102.95C171.86 103.55 172.15 104.16 172.42 104.79C172.68 105.42 172.92 106.06 173.12 106.71C173.32 107.36 173.49 108.02 173.62 108.69C173.76 109.36 173.86 110.03 173.93 110.71C174 111.38 174.03 112.06 174.03 112.75L174.03 114.05C174.02 114.32 174.04 114.59 174.09 114.85C174.13 115.12 174.2 115.37 174.29 115.63C174.38 115.88 174.49 116.12 174.63 116.36C174.76 116.59 174.92 116.81 175.09 117.01C175.26 117.22 175.45 117.41 175.66 117.58C175.87 117.75 176.09 117.9 176.33 118.03C176.56 118.16 176.81 118.27 177.06 118.36C177.32 118.45 177.58 118.51 177.84 118.55C178.15 118.6 178.46 118.6 178.78 118.58C179.09 118.55 179.4 118.5 179.7 118.41C180 118.32 180.28 118.19 180.56 118.04C180.83 117.89 181.09 117.71 181.32 117.51C181.56 117.3 181.77 117.07 181.96 116.82C182.14 116.57 182.3 116.3 182.43 116.02C182.56 115.74 182.66 115.44 182.73 115.13C182.79 114.83 182.83 114.52 182.83 114.21L182.83 112.75C182.83 111.79 182.78 110.83 182.69 109.88C182.59 108.93 182.45 107.98 182.27 107.04C182.08 106.1 181.85 105.18 181.57 104.26C181.29 103.34 180.97 102.44 180.6 101.56C180.23 100.68 179.82 99.81 179.37 98.97C178.92 98.12 178.43 97.3 177.89 96.51C177.36 95.71 176.79 94.94 176.18 94.2C175.57 93.46 174.93 92.75 174.25 92.07C173.57 91.4 172.86 90.75 172.12 90.15C171.38 89.54 170.61 88.97 169.81 88.44C169.01 87.9 168.19 87.41 167.34 86.96C166.5 86.51 165.63 86.1 164.75 85.73C163.86 85.36 162.96 85.04 162.04 84.76C161.12 84.48 160.19 84.25 159.25 84.06C158.31 83.88 157.36 83.73 156.4 83.64C155.45 83.54 154.49 83.5 153.53 83.5C137.37 83.5 124.16 97.15 124.16 113.25L124.16 125.32C124.16 125.46 124.15 125.59 124.12 125.72C124.1 125.85 124.06 125.98 124.01 126.11C123.95 126.23 123.89 126.35 123.82 126.46C123.74 126.57 123.66 126.68 123.56 126.77C123.47 126.87 123.36 126.95 123.25 127.03C123.14 127.1 123.02 127.16 122.89 127.21C122.77 127.27 122.64 127.3 122.51 127.33C122.38 127.36 122.24 127.37 122.11 127.37L121.23 127.37C120.84 127.37 120.46 127.39 120.08 127.43C119.7 127.47 119.32 127.53 118.94 127.6C118.57 127.68 118.19 127.77 117.83 127.88C117.46 127.99 117.1 128.12 116.74 128.27C116.39 128.42 116.04 128.58 115.7 128.76C115.37 128.94 115.04 129.14 114.72 129.35C114.4 129.57 114.09 129.79 113.79 130.04C113.5 130.28 113.21 130.54 112.94 130.81C112.67 131.08 112.41 131.36 112.17 131.66C111.93 131.95 111.7 132.26 111.48 132.58C111.27 132.9 111.07 133.22 110.89 133.56C110.71 133.9 110.55 134.25 110.4 134.6C110.25 134.95 110.12 135.31 110.01 135.68C109.9 136.04 109.8 136.42 109.73 136.79C109.65 137.17 109.59 137.54 109.56 137.92C109.52 138.31 109.5 138.69 109.5 139.07L109.5 188.8C109.5 189.18 109.52 189.56 109.56 189.94C109.59 190.32 109.65 190.7 109.73 191.08C109.8 191.45 109.9 191.82 110.01 192.19C110.12 192.55 110.25 192.91 110.4 193.27C110.55 193.62 110.71 193.97 110.89 194.3C111.07 194.64 111.27 194.97 111.48 195.29C111.7 195.61 111.93 195.91 112.17 196.21C112.41 196.5 112.67 196.79 112.94 197.06C113.21 197.33 113.5 197.59 113.79 197.83C114.09 198.07 114.4 198.3 114.72 198.51C115.04 198.73 115.37 198.92 115.7 199.1C116.04 199.28 116.39 199.45 116.74 199.6C117.1 199.74 117.46 199.87 117.83 199.98C118.19 200.1 118.57 200.19 118.94 200.26C119.32 200.34 119.7 200.4 120.08 200.44C120.46 200.47 120.84 200.49 121.23 200.5L185.76 200.5C186.15 200.49 186.53 200.47 186.91 200.44C187.29 200.4 187.67 200.34 188.05 200.26C188.43 200.19 188.8 200.1 189.16 199.98C189.53 199.87 189.89 199.74 190.25 199.6C190.6 199.45 190.95 199.28 191.29 199.1C191.62 198.92 191.95 198.73 192.27 198.51C192.59 198.3 192.9 198.07 193.2 197.83C193.49 197.59 193.78 197.33 194.05 197.06C194.32 196.79 194.58 196.5 194.82 196.21C195.07 195.91 195.29 195.61 195.51 195.29C195.72 194.97 195.92 194.64 196.1 194.3C196.28 193.97 196.45 193.62 196.59 193.27C196.74 192.91 196.87 192.55 196.98 192.19C197.09 191.82 197.19 191.45 197.26 191.08C197.34 190.7 197.4 190.32 197.44 189.94C197.47 189.56 197.49 189.18 197.5 188.8L197.5 139.07C197.49 138.69 197.47 138.31 197.44 137.92C197.4 137.54 197.34 137.17 197.26 136.79C197.19 136.42 197.09 136.04 196.98 135.68C196.87 135.31 196.74 134.95 196.59 134.6C196.45 134.25 196.28 133.9 196.1 133.56C195.92 133.22 195.72 132.9 195.51 132.58C195.29 132.26 195.07 131.95 194.82 131.66C194.58 131.36 194.32 131.08 194.05 130.81C193.78 130.54 193.49 130.28 193.2 130.04C192.9 129.79 192.59 129.57 192.27 129.35C191.95 129.14 191.62 128.94 191.29 128.76C190.95 128.58 190.6 128.42 190.25 128.27C189.89 128.12 189.53 127.99 189.16 127.88C188.8 127.77 188.43 127.68 188.05 127.6C187.67 127.53 187.29 127.47 186.91 127.43C186.53 127.39 186.15 127.37 185.76 127.37ZM188.7 185.87C188.7 186.25 188.66 186.64 188.58 187.01C188.51 187.39 188.4 187.75 188.25 188.11C188.1 188.46 187.92 188.8 187.71 189.12C187.49 189.44 187.25 189.74 186.98 190.01C186.7 190.28 186.41 190.52 186.09 190.73C185.77 190.95 185.43 191.13 185.07 191.28C184.72 191.42 184.35 191.53 183.97 191.61C183.6 191.68 183.21 191.72 182.83 191.72L124.16 191.72C123.78 191.72 123.4 191.68 123.02 191.61C122.64 191.53 122.27 191.42 121.92 191.28C121.56 191.13 121.22 190.95 120.9 190.73C120.58 190.52 120.29 190.28 120.01 190.01C119.74 189.74 119.5 189.44 119.28 189.12C119.07 188.8 118.89 188.46 118.74 188.11C118.59 187.75 118.48 187.39 118.41 187.01C118.33 186.64 118.3 186.25 118.3 185.87L118.3 142C118.3 141.61 118.33 141.23 118.41 140.85C118.48 140.48 118.59 140.11 118.74 139.76C118.89 139.4 119.07 139.06 119.28 138.75C119.5 138.43 119.74 138.13 120.01 137.86C120.29 137.59 120.58 137.34 120.9 137.13C121.22 136.92 121.56 136.74 121.92 136.59C122.27 136.44 122.64 136.33 123.02 136.26C123.4 136.18 123.78 136.14 124.16 136.14L182.83 136.14C183.21 136.14 183.6 136.18 183.97 136.26C184.35 136.33 184.72 136.44 185.07 136.59C185.43 136.74 185.77 136.92 186.09 137.13C186.41 137.34 186.7 137.59 186.98 137.86C187.25 138.13 187.49 138.43 187.71 138.75C187.92 139.06 188.1 139.4 188.25 139.76C188.4 140.11 188.51 140.48 188.58 140.85C188.66 141.23 188.7 141.61 188.7 142L188.7 185.87Z",
                  fill: "#888",
                  "fill-opacity": "1.000000",
                  "fill-rule": "nonzero"
                }
              }),
              _vm._v(" "),
              _c("path", {
                attrs: {
                  id: "path-imb4gmi8ix6s1yc-UmGY9TvFNiqwKS",
                  d:
                    "M163.76 155.16C163.76 154.74 163.74 154.32 163.69 153.91C163.64 153.5 163.56 153.09 163.46 152.68C163.36 152.28 163.23 151.88 163.08 151.49C162.93 151.1 162.76 150.72 162.56 150.35C162.37 149.98 162.15 149.63 161.91 149.28C161.67 148.94 161.41 148.61 161.13 148.3C160.85 147.99 160.55 147.7 160.24 147.43C159.92 147.15 159.59 146.9 159.24 146.67C158.89 146.43 158.53 146.22 158.16 146.03C157.79 145.84 157.41 145.67 157.01 145.53C156.62 145.39 156.22 145.27 155.81 145.17C155.4 145.08 154.99 145.01 154.57 144.97C154.16 144.92 153.74 144.91 153.32 144.91C152.9 144.92 152.49 144.95 152.07 145.01C151.66 145.07 151.25 145.15 150.84 145.26C150.44 145.37 150.04 145.5 149.65 145.66C149.27 145.81 148.89 145.99 148.52 146.19C148.16 146.39 147.8 146.62 147.46 146.86C147.13 147.11 146.8 147.37 146.5 147.66C146.19 147.94 145.9 148.25 145.63 148.57C145.36 148.89 145.12 149.22 144.89 149.57C144.66 149.92 144.45 150.28 144.27 150.66C144.09 151.03 143.92 151.42 143.79 151.81C143.65 152.21 143.54 152.61 143.45 153.02C143.37 153.43 143.3 153.84 143.27 154.25C143.23 154.67 143.22 155.09 143.24 155.5C143.25 155.92 143.29 156.34 143.35 156.75C143.42 157.16 143.51 157.57 143.63 157.97C143.74 158.37 143.88 158.76 144.04 159.15C144.21 159.53 144.39 159.9 144.6 160.27C144.81 160.63 145.04 160.98 145.3 161.31C145.55 161.64 145.82 161.96 146.11 162.26C146.58 162.75 147 163.27 147.38 163.84C147.75 164.41 148.06 165.01 148.32 165.63C148.58 166.26 148.77 166.91 148.9 167.57C149.03 168.24 149.1 168.91 149.1 169.59L149.1 178.39C149.09 178.66 149.11 178.93 149.16 179.2C149.2 179.46 149.27 179.72 149.36 179.97C149.45 180.23 149.56 180.47 149.7 180.7C149.83 180.93 149.98 181.15 150.16 181.36C150.33 181.57 150.52 181.75 150.73 181.93C150.94 182.1 151.16 182.25 151.39 182.38C151.63 182.51 151.87 182.62 152.13 182.71C152.38 182.8 152.64 182.86 152.91 182.9C153.22 182.95 153.53 182.95 153.84 182.93C154.16 182.91 154.46 182.85 154.76 182.76C155.06 182.67 155.35 182.55 155.63 182.4C155.9 182.24 156.16 182.07 156.39 181.86C156.63 181.65 156.84 181.43 157.03 181.17C157.22 180.92 157.38 180.66 157.5 180.37C157.64 180.09 157.73 179.79 157.8 179.49C157.87 179.18 157.9 178.87 157.9 178.56L157.9 169.59C157.9 168.92 157.96 168.25 158.09 167.58C158.21 166.92 158.4 166.28 158.65 165.65C158.91 165.03 159.21 164.43 159.58 163.86C159.95 163.3 160.37 162.77 160.83 162.28C161.3 161.81 161.71 161.3 162.08 160.75C162.44 160.2 162.75 159.62 163 159.01C163.25 158.4 163.44 157.77 163.57 157.12C163.7 156.47 163.76 155.82 163.76 155.16Z",
                  fill: "#888",
                  "fill-opacity": "1.000000",
                  "fill-rule": "nonzero"
                }
              }),
              _vm._v(" "),
              _c("path", {
                attrs: {
                  id: "path-imb4gmi8ix6s1yc-UmGY9TvFNiqwKS",
                  d:
                    "M92.5 163.71C87.51 164.23 82.51 164.5 77.5 164.5C40.87 164.5 17.5 151.17 17.5 142L17.5 130.05C30.96 138.95 52.41 144.5 77.5 144.5C82.66 144.5 87.67 144.25 92.5 143.8L92.5 128.71C87.51 129.23 82.51 129.5 77.5 129.5C40.87 129.5 17.5 116.17 17.5 107L17.5 95C30.96 103.95 52.41 109.5 77.5 109.5C82.66 109.5 87.67 109.25 92.5 108.8L92.5 93.71C87.51 94.23 82.51 94.5 77.5 94.5C40.87 94.5 17.5 81.17 17.5 72L17.5 60C30.96 68.94 52.4 74.5 77.5 74.5C102.59 74.5 124.04 68.94 137.5 60.04L137.5 69.5L152.5 69.5L152.5 37C152.5 15.62 120.26 -0.5 77.5 -0.5C34.74 -0.5 2.5 15.62 2.5 37L2.5 142C2.5 163.38 34.74 179.5 77.5 179.5C82.66 179.5 87.67 179.26 92.5 178.81L92.5 163.71ZM77.5 14.5C114.13 14.5 137.5 27.83 137.5 37C137.5 46.17 114.13 59.5 77.5 59.5C40.87 59.5 17.5 46.17 17.5 37C17.5 27.83 40.87 14.5 77.5 14.5Z",
                  fill: "#888",
                  "fill-opacity": "1.000000",
                  "fill-rule": "nonzero"
                }
              })
            ])
          ])
        ]
      )
    : _vm.type != "dataset" && !_vm.isPrivate && _vm.modelType != 2
    ? _c(
        "svg",
        {
          attrs: {
            xmlns: "http://www.w3.org/2000/svg",
            "xmlns:xlink": "http://www.w3.org/1999/xlink",
            viewBox: "0 0 200 200",
            width: "200.000000",
            height: "200.000000",
            fill: "none"
          }
        },
        [
          _c("rect", {
            attrs: {
              id: "公开模型",
              width: "200.000000",
              height: "200.000000",
              x: "0.000000",
              y: "0.000000"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M102 139C101.053 139 100.108 138.754 99.2596 138.263L66.7404 119.429C66.324 119.188 65.9457 118.897 65.6057 118.555C65.2656 118.214 64.9753 117.835 64.7348 117.417C64.4943 116.999 64.3118 116.557 64.1872 116.091C64.0626 115.625 64.0002 115.151 64 114.668L64 77.0025C64 75.0387 65.0457 73.2243 66.7404 72.2413L74.1067 67.9748C76.7288 66.4574 80.0809 67.3569 81.5936 69.9871C83.1063 72.6173 82.2097 75.9799 79.5876 77.4973L74.9617 80.176L74.9617 111.495L102 127.153L129.038 111.495L129.038 80.176L124.005 77.2598C121.383 75.7424 120.486 72.3798 121.999 69.7496C123.511 67.1193 126.863 66.2199 129.485 67.7373L137.26 72.2391C137.676 72.4802 138.055 72.7714 138.395 73.1126C138.735 73.4538 139.025 73.8333 139.266 74.2512C139.506 74.6691 139.689 75.1111 139.813 75.5772C139.938 76.0434 140 76.5177 140 77.0003L140 114.668C140 116.632 138.956 118.446 137.26 119.429L104.74 138.263C103.892 138.754 102.978 139 102 139Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "nonzero"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M38 72C37.0529 72 36.108 71.7537 35.2596 71.2633L2.74044 52.4294C2.32397 52.188 1.94572 51.8966 1.60566 51.5554C1.26561 51.2141 0.975331 50.8346 0.734827 50.4167C0.494323 49.9989 0.311776 49.5569 0.187185 49.0909C0.0625944 48.6249 0.00019938 48.1506 0 47.6681L0 10.0025C0 8.03865 1.04575 6.22431 2.74044 5.24127L10.1067 0.974839C12.7288 -0.542604 16.0809 0.356866 17.5936 2.9871C19.1063 5.61733 18.2097 8.9799 15.5876 10.4973L10.9617 13.176L10.9617 44.4947L38 60.1529L65.0383 44.4947L65.0383 13.176L60.0046 10.2598C57.3826 8.74239 56.4859 5.37982 57.9986 2.74958C59.5113 0.119349 62.8634 -0.780121 65.4855 0.737322L73.2596 5.23907C73.6761 5.48025 74.0545 5.77142 74.3947 6.11258C74.7348 6.45375 75.0251 6.83328 75.2657 7.25117C75.5062 7.66907 75.6887 8.11108 75.8132 8.57722C75.9377 9.04335 76 9.51772 76 10.0003L76 47.6681C76 49.632 74.9564 51.4463 73.2596 52.4294L40.7404 71.2633C39.8919 71.7544 38.9785 72 38 72Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "nonzero",
              transform: "matrix(-1,0,0,-1,140,125)"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M101.5 200C100.557 200 99.6154 199.756 98.7703 199.271L62.7297 178.521C62.3149 178.282 61.9381 177.993 61.5994 177.655C61.2607 177.317 60.9715 176.942 60.7319 176.528C60.4924 176.114 60.3106 175.677 60.1865 175.215C60.0624 174.754 60.0002 174.284 60 173.806L60 132.307C60 130.362 61.0416 128.566 62.7297 127.592L79.8612 117.73C82.473 116.228 85.8119 117.118 87.3187 119.723C88.8255 122.327 87.9324 125.656 85.3206 127.159L70.9188 135.451L70.9188 170.664L101.5 188.27L132.081 170.664L132.081 135.451L117.933 127.305C115.321 125.802 114.428 122.473 115.935 119.868C117.441 117.264 120.78 116.373 123.392 117.876L140.27 127.592C140.685 127.831 141.062 128.119 141.401 128.457C141.74 128.795 142.029 129.171 142.269 129.584C142.508 129.998 142.69 130.436 142.814 130.897C142.938 131.359 143 131.829 143 132.307L143 173.806C143 175.751 141.961 177.547 140.27 178.521L104.23 199.271C103.385 199.757 102.475 200 101.5 200Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "nonzero"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M50.8309 173C49.8857 173 48.9427 172.755 48.096 172.266L11.7349 151.25C11.3193 151.01 10.9418 150.72 10.6024 150.38C10.2631 150.04 9.97336 149.662 9.73335 149.246C9.49333 148.829 9.31115 148.389 9.18681 147.925C9.06247 147.461 9.0002 146.989 9 146.508L9 104.476C9 101.451 11.4483 99 14.4698 99C17.4913 99 19.9396 101.451 19.9396 104.476L19.9396 143.347L50.8309 161.203L81.7222 143.347L81.7222 125.768C81.7222 122.743 84.1705 120.292 87.192 120.292C90.2135 120.292 92.6618 122.743 92.6618 125.768L92.6618 146.508C92.6618 148.464 91.6204 150.271 89.9269 151.25L53.5658 172.266C52.7199 172.756 51.8083 173.001 50.8309 173ZM151.169 173C150.224 173 149.281 172.755 148.434 172.266L112.073 151.25C111.657 151.01 111.28 150.72 110.94 150.38C110.601 150.04 110.311 149.662 110.071 149.246C109.831 148.83 109.649 148.39 109.525 147.925C109.4 147.461 109.338 146.989 109.338 146.508L109.338 126.072C109.338 123.048 111.786 120.597 114.808 120.597C117.83 120.597 120.278 123.048 120.278 126.072L120.278 143.345L151.169 161.201L182.06 143.345L182.06 104.476C182.06 101.451 184.509 99 187.53 99C190.552 99 193 101.451 193 104.476L193 146.508C193 148.464 191.959 150.271 190.265 151.25L153.904 172.266C153.058 172.756 152.146 173.001 151.169 173Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "nonzero"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M155.141 112C154.198 112 153.257 111.755 152.412 111.266L131.526 99.1615C128.915 97.6498 128.022 94.3 129.528 91.6797C131.034 89.0594 134.372 88.1633 136.982 89.675L155.141 100.198L182.086 84.5835L182.086 53.3526L155.141 37.7383L127.776 53.5958C126.953 55.5741 125.008 56.9675 122.738 56.9675C119.723 56.9675 117.281 54.516 117.281 51.4904L117.281 50.1912C117.281 48.2347 118.32 46.4273 120.009 45.448L152.412 26.67C152.827 26.4294 153.265 26.2469 153.728 26.1223C154.191 25.9978 154.662 25.9355 155.141 25.9355C155.619 25.9355 156.09 25.9978 156.553 26.1223C157.016 26.2469 157.454 26.4294 157.869 26.67L190.272 45.448C190.686 45.6882 191.063 45.9783 191.402 46.3182C191.74 46.658 192.029 47.0361 192.269 47.4524C192.508 47.8688 192.69 48.3091 192.814 48.7735C192.938 49.2378 193 49.7104 193 50.1912L193 87.7449C193 89.7013 191.961 91.5088 190.272 92.4881L157.869 111.266C157.024 111.755 156.115 112 155.141 112ZM46.8595 112C45.9166 112 44.9758 111.755 44.1311 111.266L11.7284 92.4881C11.3138 92.2476 10.9372 91.9574 10.5986 91.6174C10.2601 91.2775 9.97106 90.8994 9.73161 90.4831C9.49216 90.0668 9.31041 89.6266 9.18637 89.1623C9.06232 88.698 9.0002 88.2256 9 87.7449L9 50.1912C9 48.2347 10.0412 46.4273 11.7284 45.4479L44.1311 26.67C44.5458 26.4294 44.9844 26.2469 45.4471 26.1223C45.9097 25.9978 46.3805 25.9355 46.8595 25.9355C47.3385 25.9355 47.8093 25.9978 48.2719 26.1223C48.7345 26.2469 49.1732 26.4294 49.5879 26.67L81.9906 45.4479C82.4053 45.6882 82.782 45.9783 83.1207 46.3182C83.4594 46.658 83.7484 47.0361 83.9879 47.4524C84.2274 47.8687 84.4091 48.3091 84.533 48.7735C84.657 49.2378 84.719 49.7104 84.719 50.1912L84.719 52.0162C84.719 55.0418 82.2765 57.4933 79.2622 57.4933C76.7477 57.4933 74.6326 55.7867 73.9996 53.4665L46.8617 37.7405L19.9158 53.3548L19.9158 84.5856L46.8617 100.2L66.3535 88.9038C68.964 87.3921 72.3014 88.2882 73.8075 90.9085C75.3136 93.5288 74.4209 96.8786 71.8103 98.3903L49.5901 111.266C48.745 111.755 47.8349 112 46.8595 112ZM69.0797 33.8451C66.0654 33.8451 63.6229 31.3935 63.6229 28.3679L63.6229 24.2557C63.6229 22.2992 64.664 20.4917 66.3513 19.5124L98.754 0.734481C99.1687 0.493905 99.6073 0.311347 100.07 0.186808C100.533 0.0622689 101.003 -6.15908e-07 101.482 -6.15908e-07C101.961 -6.15908e-07 102.432 0.0622689 102.895 0.186808C103.357 0.311347 103.796 0.493905 104.211 0.734481L136.613 19.5124C137.028 19.7527 137.405 20.0428 137.744 20.3826C138.082 20.7225 138.371 21.1006 138.611 21.5169C138.85 21.9332 139.032 22.3736 139.156 22.8379C139.28 23.3023 139.342 23.7749 139.342 24.2557L139.342 28.197C139.342 31.2226 136.899 33.6742 133.885 33.6742C130.871 33.6742 128.428 31.2226 128.428 28.197L128.428 27.4193L101.482 11.8049L74.5366 27.4193L74.5366 28.3701C74.5366 31.3935 72.0941 33.8451 69.0797 33.8451Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "nonzero"
            }
          })
        ]
      )
    : _vm.type != "dataset" && _vm.isPrivate
    ? _c(
        "svg",
        {
          attrs: {
            xmlns: "http://www.w3.org/2000/svg",
            "xmlns:xlink": "http://www.w3.org/1999/xlink",
            viewBox: "0 0 200 200",
            width: "200.000000",
            height: "200.000000",
            fill: "none"
          }
        },
        [
          _c("rect", {
            attrs: {
              id: "私有模型",
              width: "200.000000",
              height: "200.000000",
              x: "0.000000",
              y: "0.000000"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M102 139C101.053 139 100.108 138.754 99.2596 138.263L66.7404 119.429C66.324 119.188 65.9457 118.897 65.6057 118.555C65.2656 118.214 64.9753 117.835 64.7348 117.417C64.4943 116.999 64.3118 116.557 64.1872 116.091C64.0626 115.625 64.0002 115.151 64 114.668L64 77.0025C64 75.0387 65.0457 73.2243 66.7404 72.2413L74.1067 67.9748C76.7288 66.4574 80.0809 67.3569 81.5936 69.9871C83.1063 72.6173 82.2097 75.9799 79.5876 77.4973L74.9617 80.176L74.9617 111.495L102 127.153L129.038 111.495L129.038 80.176L124.005 77.2598C121.383 75.7424 120.486 72.3798 121.999 69.7496C123.511 67.1193 126.863 66.2199 129.485 67.7373L137.26 72.2391C137.676 72.4802 138.055 72.7714 138.395 73.1126C138.735 73.4538 139.025 73.8333 139.266 74.2512C139.506 74.6691 139.689 75.1111 139.813 75.5772C139.938 76.0434 140 76.5177 140 77.0003L140 114.668C140 116.632 138.956 118.446 137.26 119.429L104.74 138.263C103.892 138.754 102.978 139 102 139Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "nonzero"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M38 72C37.0529 72 36.108 71.7537 35.2596 71.2633L2.74044 52.4294C2.32397 52.188 1.94572 51.8966 1.60566 51.5554C1.26561 51.2141 0.975331 50.8346 0.734827 50.4167C0.494323 49.9989 0.311776 49.5569 0.187185 49.0909C0.0625944 48.6249 0.00019938 48.1506 0 47.6681L0 10.0025C0 8.03865 1.04575 6.22431 2.74044 5.24127L10.1067 0.974839C12.7288 -0.542604 16.0809 0.356866 17.5936 2.9871C19.1063 5.61733 18.2097 8.9799 15.5876 10.4973L10.9617 13.176L10.9617 44.4947L38 60.1529L65.0383 44.4947L65.0383 13.176L60.0046 10.2598C57.3826 8.74239 56.4859 5.37982 57.9986 2.74958C59.5113 0.119349 62.8634 -0.780121 65.4855 0.737322L73.2596 5.23907C73.6761 5.48025 74.0545 5.77142 74.3947 6.11258C74.7348 6.45375 75.0251 6.83328 75.2657 7.25117C75.5062 7.66907 75.6887 8.11108 75.8132 8.57722C75.9377 9.04335 76 9.51772 76 10.0003L76 47.6681C76 49.632 74.9564 51.4463 73.2596 52.4294L40.7404 71.2633C39.8919 71.7544 38.9785 72 38 72Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "nonzero",
              transform: "matrix(-1,0,0,-1,140,125)"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M101.5 200C100.557 200 99.6154 199.756 98.7703 199.271L62.7297 178.521C62.3149 178.282 61.9381 177.993 61.5994 177.655C61.2607 177.317 60.9715 176.942 60.7319 176.528C60.4924 176.114 60.3106 175.677 60.1865 175.215C60.0624 174.754 60.0002 174.284 60 173.806L60 132.307C60 130.362 61.0416 128.566 62.7297 127.592L79.8612 117.73C82.473 116.228 85.8119 117.118 87.3187 119.723C88.8255 122.327 87.9324 125.656 85.3206 127.159L70.9188 135.451L70.9188 170.664L101.5 188.27L132.081 170.664L132.081 135.451L117.933 127.305C115.321 125.802 114.428 122.473 115.935 119.868C117.441 117.264 120.78 116.373 123.392 117.876L140.27 127.592C140.685 127.831 141.062 128.119 141.401 128.457C141.74 128.795 142.029 129.171 142.269 129.584C142.508 129.998 142.69 130.436 142.814 130.897C142.938 131.359 143 131.829 143 132.307L143 173.806C143 175.751 141.961 177.547 140.27 178.521L104.23 199.271C103.385 199.757 102.475 200 101.5 200Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "nonzero"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M50.8309 173C49.8857 173 48.9427 172.755 48.096 172.266L11.7349 151.25C11.3193 151.01 10.9418 150.72 10.6024 150.38C10.2631 150.04 9.97336 149.662 9.73335 149.246C9.49333 148.829 9.31115 148.389 9.18681 147.925C9.06247 147.461 9.0002 146.989 9 146.508L9 104.476C9 101.451 11.4483 99 14.4698 99C17.4913 99 19.9396 101.451 19.9396 104.476L19.9396 143.347L50.8309 161.203L81.7222 143.347L81.7222 125.768C81.7222 122.743 84.1705 120.292 87.192 120.292C90.2135 120.292 92.6618 122.743 92.6618 125.768L92.6618 146.508C92.6618 148.464 91.6204 150.271 89.9269 151.25L53.5658 172.266C52.7199 172.756 51.8083 173.001 50.8309 173ZM151.169 173C150.224 173 149.281 172.755 148.434 172.266L112.073 151.25C111.657 151.01 111.28 150.72 110.94 150.38C110.601 150.04 110.311 149.662 110.071 149.246C109.831 148.83 109.649 148.39 109.525 147.925C109.4 147.461 109.338 146.989 109.338 146.508L109.338 126.072C109.338 123.048 111.786 120.597 114.808 120.597C117.83 120.597 120.278 123.048 120.278 126.072L120.278 143.345L151.169 161.201L182.06 143.345L182.06 104.476C182.06 101.451 184.509 99 187.53 99C190.552 99 193 101.451 193 104.476L193 146.508C193 148.464 191.959 150.271 190.265 151.25L153.904 172.266C153.058 172.756 152.146 173.001 151.169 173Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "nonzero"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M155.141 112C154.198 112 153.257 111.755 152.412 111.266L131.526 99.1615C128.915 97.6498 128.022 94.3 129.528 91.6797C131.034 89.0594 134.372 88.1633 136.982 89.675L155.141 100.198L182.086 84.5835L182.086 53.3526L155.141 37.7383L127.776 53.5958C126.953 55.5741 125.008 56.9675 122.738 56.9675C119.723 56.9675 117.281 54.516 117.281 51.4904L117.281 50.1912C117.281 48.2347 118.32 46.4273 120.009 45.448L152.412 26.67C152.827 26.4294 153.265 26.2469 153.728 26.1223C154.191 25.9978 154.662 25.9355 155.141 25.9355C155.619 25.9355 156.09 25.9978 156.553 26.1223C157.016 26.2469 157.454 26.4294 157.869 26.67L190.272 45.448C190.686 45.6882 191.063 45.9783 191.402 46.3182C191.74 46.658 192.029 47.0361 192.269 47.4524C192.508 47.8688 192.69 48.3091 192.814 48.7735C192.938 49.2378 193 49.7104 193 50.1912L193 87.7449C193 89.7013 191.961 91.5088 190.272 92.4881L157.869 111.266C157.024 111.755 156.115 112 155.141 112ZM46.8595 112C45.9166 112 44.9758 111.755 44.1311 111.266L11.7284 92.4881C11.3138 92.2476 10.9372 91.9574 10.5986 91.6174C10.2601 91.2775 9.97106 90.8994 9.73161 90.4831C9.49216 90.0668 9.31041 89.6266 9.18637 89.1623C9.06232 88.698 9.0002 88.2256 9 87.7449L9 50.1912C9 48.2347 10.0412 46.4273 11.7284 45.4479L44.1311 26.67C44.5458 26.4294 44.9844 26.2469 45.4471 26.1223C45.9097 25.9978 46.3805 25.9355 46.8595 25.9355C47.3385 25.9355 47.8093 25.9978 48.2719 26.1223C48.7345 26.2469 49.1732 26.4294 49.5879 26.67L81.9906 45.4479C82.4053 45.6882 82.782 45.9783 83.1207 46.3182C83.4594 46.658 83.7484 47.0361 83.9879 47.4524C84.2274 47.8687 84.4091 48.3091 84.533 48.7735C84.657 49.2378 84.719 49.7104 84.719 50.1912L84.719 52.0162C84.719 55.0418 82.2765 57.4933 79.2622 57.4933C76.7477 57.4933 74.6326 55.7867 73.9996 53.4665L46.8617 37.7405L19.9158 53.3548L19.9158 84.5856L46.8617 100.2L66.3535 88.9038C68.964 87.3921 72.3014 88.2882 73.8075 90.9085C75.3136 93.5288 74.4209 96.8786 71.8103 98.3903L49.5901 111.266C48.745 111.755 47.8349 112 46.8595 112ZM69.0797 33.8451C66.0654 33.8451 63.6229 31.3935 63.6229 28.3679L63.6229 24.2557C63.6229 22.2992 64.664 20.4917 66.3513 19.5124L98.754 0.734481C99.1687 0.493905 99.6073 0.311347 100.07 0.186808C100.533 0.0622689 101.003 -6.15908e-07 101.482 -6.15908e-07C101.961 -6.15908e-07 102.432 0.0622689 102.895 0.186808C103.357 0.311347 103.796 0.493905 104.211 0.734481L136.613 19.5124C137.028 19.7527 137.405 20.0428 137.744 20.3826C138.082 20.7225 138.371 21.1006 138.611 21.5169C138.85 21.9332 139.032 22.3736 139.156 22.8379C139.28 23.3023 139.342 23.7749 139.342 24.2557L139.342 28.197C139.342 31.2226 136.899 33.6742 133.885 33.6742C130.871 33.6742 128.428 31.2226 128.428 28.197L128.428 27.4193L101.482 11.8049L74.5366 27.4193L74.5366 28.3701C74.5366 31.3935 72.0941 33.8451 69.0797 33.8451Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "nonzero"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "多边形 4",
              d:
                "M157.622 61.1132C160.716 62.8996 162.622 66.2009 162.622 69.7735L162.622 128.227C162.622 131.799 160.716 135.1 157.622 136.887L107 166.113C103.906 167.9 100.094 167.9 97 166.113L46.3782 136.887C43.2842 135.1 41.3782 131.799 41.3782 128.226L41.3782 69.7735C41.3782 66.2009 43.2842 62.8996 46.3782 61.1133L97 31.8868C100.094 30.1004 103.906 30.1004 107 31.8868L157.622 61.1132Z",
              fill: "rgb(255,255,255)",
              "fill-rule": "evenodd"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "多边形 4",
              d:
                "M162.622 69.7735C162.622 66.2009 160.716 62.8996 157.622 61.1132L107 31.8868C103.906 30.1004 100.094 30.1004 97 31.8868L46.3782 61.1133C43.2842 62.8996 41.3782 66.2009 41.3782 69.7735L41.3782 128.226C41.3782 131.799 43.2842 135.1 46.3782 136.887L97 166.113C100.094 167.9 103.906 167.9 107 166.113L157.622 136.887C160.716 135.1 162.622 131.799 162.622 128.227L162.622 69.7735ZM102 40.547L152.622 69.7735L152.622 128.227L102 157.453L51.3782 128.226L51.3782 69.7735L102 40.547Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "evenodd"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "矩形 13",
              d:
                "M84 73L84 81L84 89Q84 90.9374 84.398 92.7892Q84.7512 94.4327 85.4178 96.0089Q86.0778 97.5693 86.9991 98.9562Q87.9857 100.442 89.2721 101.728Q90.5584 103.014 92.0438 104.001Q93.4307 104.922 94.9911 105.582Q96.5674 106.249 98.2109 106.602Q100.063 107 102 107Q103.937 107 105.789 106.602Q107.433 106.249 109.009 105.582Q110.569 104.922 111.956 104.001Q113.442 103.014 114.728 101.728Q116.014 100.442 117.001 98.9563Q117.922 97.5694 118.582 96.0089Q119.249 94.4327 119.602 92.7891Q120 90.9373 120 89L120 73Q120 71.0626 119.602 69.2107Q119.249 67.5672 118.582 65.9911Q117.922 64.4307 117.001 63.0439Q116.014 61.5584 114.728 60.2721Q113.442 58.9857 111.956 57.999Q110.569 57.0778 109.009 56.4178Q107.433 55.7512 105.789 55.398Q103.937 55 102 55Q100.063 55 98.2108 55.398Q96.5673 55.7512 94.9911 56.4178Q93.4307 57.0778 92.0439 57.999Q90.5584 58.9857 89.2721 60.2721Q87.9857 61.5584 86.9991 63.0438Q86.0778 64.4307 85.4178 65.9911Q84.7512 67.5673 84.398 69.2108Q84 71.0626 84 73ZM102 63C96.4771 63 92 67.4771 92 73L92 89C92 94.5228 96.4771 99 102 99C107.523 99 112 94.5228 112 89L112 73C112 67.4771 107.523 63 102 63Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "evenodd"
            }
          }),
          _vm._v(" "),
          _c("rect", {
            attrs: {
              id: "矩形 13",
              width: "46.000000",
              height: "36.000000",
              x: "79.000000",
              y: "89.000000",
              rx: "5.000000",
              fill: "rgb(136,136,136)"
            }
          }),
          _vm._v(" "),
          _c("rect", {
            attrs: {
              id: "矩形 13",
              width: "54.000000",
              height: "44.000000",
              x: "75.000000",
              y: "85.000000",
              rx: "9.000000",
              stroke: "rgb(136,136,136)",
              "stroke-width": "8.000000"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "path",
              d:
                "M107.6 101.305C107.6 101.048 107.585 100.792 107.554 100.537C107.523 100.282 107.476 100.03 107.414 99.7808C107.352 99.5315 107.275 99.2868 107.183 99.0469C107.091 98.807 106.985 98.5735 106.865 98.3466C106.744 98.1196 106.611 97.9009 106.463 97.6903C106.316 97.4797 106.157 97.2789 105.985 97.0878C105.813 96.8968 105.63 96.7169 105.436 96.5482C105.243 96.3795 105.039 96.2232 104.826 96.0794C104.614 95.9356 104.393 95.8053 104.164 95.6885C103.935 95.5717 103.7 95.4693 103.458 95.3813C103.217 95.2933 102.971 95.2203 102.721 95.1623C102.471 95.1044 102.218 95.0619 101.962 95.0349C101.707 95.0079 101.45 94.9966 101.194 95.0009C100.937 95.0052 100.681 95.0252 100.427 95.0608C100.172 95.0965 99.9208 95.1475 99.6726 95.2138C99.4244 95.2802 99.1811 95.3615 98.9428 95.4576C98.7046 95.5537 98.4731 95.664 98.2483 95.7885C98.0235 95.9129 97.8072 96.0506 97.5992 96.2015C97.3913 96.3525 97.1933 96.5155 97.0054 96.6907C96.8174 96.8658 96.6408 97.0518 96.4755 97.2485C96.3103 97.4453 96.1577 97.6514 96.0177 97.8668C95.8777 98.0822 95.7513 98.3054 95.6386 98.5363C95.5259 98.7671 95.4276 99.004 95.3439 99.2469C95.2602 99.4899 95.1916 99.737 95.1381 99.9883C95.0846 100.24 95.0466 100.493 95.0241 100.749C95.0017 101.005 94.9949 101.261 95.0038 101.518C95.0127 101.775 95.0372 102.03 95.0773 102.284C95.1175 102.538 95.1729 102.788 95.2437 103.035C95.3144 103.282 95.4 103.524 95.5003 103.76C95.6007 103.997 95.7151 104.227 95.8435 104.449C95.9719 104.672 96.1134 104.885 96.268 105.091C96.4226 105.296 96.5891 105.491 96.7676 105.676C97.0583 105.975 97.3178 106.299 97.5459 106.649C97.774 106.998 97.9665 107.366 98.1236 107.753C98.2806 108.139 98.3992 108.537 98.4793 108.947C98.5595 109.356 98.5997 109.769 98.6 110.187L98.6 115.605C98.5976 115.77 98.6099 115.934 98.6369 116.097C98.6639 116.261 98.7053 116.42 98.7609 116.576C98.8165 116.732 98.8855 116.881 98.9679 117.025C99.0504 117.168 99.1449 117.303 99.2515 117.429C99.3581 117.556 99.4751 117.672 99.6026 117.777C99.73 117.883 99.8657 117.976 100.01 118.057C100.154 118.139 100.304 118.206 100.46 118.26C100.617 118.315 100.777 118.355 100.94 118.38C101.131 118.406 101.322 118.411 101.514 118.396C101.706 118.381 101.894 118.345 102.079 118.29C102.263 118.234 102.439 118.16 102.608 118.067C102.776 117.974 102.933 117.864 103.078 117.737C103.223 117.611 103.353 117.47 103.468 117.316C103.583 117.161 103.681 116.997 103.76 116.821C103.84 116.646 103.901 116.464 103.942 116.276C103.983 116.088 104.003 115.898 104.004 115.705L104.004 110.187C104.002 109.772 104.04 109.36 104.117 108.952C104.195 108.545 104.311 108.148 104.465 107.763C104.62 107.378 104.809 107.011 105.034 106.662C105.26 106.313 105.516 105.989 105.804 105.69C106.088 105.4 106.342 105.086 106.566 104.747C106.789 104.407 106.978 104.05 107.132 103.674C107.286 103.298 107.402 102.911 107.481 102.512C107.56 102.114 107.599 101.712 107.6 101.305Z",
              fill: "rgb(255,255,255)",
              "fill-rule": "nonzero"
            }
          })
        ]
      )
    : _vm.type != "dataset" && _vm.modelType == 2
    ? _c(
        "svg",
        {
          attrs: {
            xmlns: "http://www.w3.org/2000/svg",
            "xmlns:xlink": "http://www.w3.org/1999/xlink",
            viewBox: "0 0 200 200",
            width: "200.000000",
            height: "200.000000",
            fill: "none"
          }
        },
        [
          _c("rect", {
            attrs: {
              id: "模型镜像",
              width: "200.000000",
              height: "200.000000",
              x: "0.000000",
              y: "0.000000"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "多边形 4",
              d:
                "M181.335 44.1133C184.429 45.8996 186.335 49.2009 186.335 52.7735L186.335 143.227C186.335 146.799 184.429 150.1 181.335 151.887L103 197.113C99.906 198.9 96.094 198.9 93 197.113L14.6654 151.887C11.5714 150.1 9.6654 146.799 9.6654 143.226L9.6654 52.7735C9.6654 49.2009 11.5714 45.8996 14.6654 44.1133L93 -1.11325C96.094 -2.89958 99.906 -2.89958 103 -1.11325L181.335 44.1133Z",
              opacity: "0",
              fill: "rgb(255,255,255)",
              "fill-rule": "evenodd"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "多边形 4",
              d:
                "M186.335 52.7735C186.335 49.2009 184.429 45.8996 181.335 44.1133L103 -1.11325C99.906 -2.89958 96.094 -2.89958 93 -1.11325L14.6654 44.1133C11.5714 45.8996 9.6654 49.2009 9.6654 52.7735L9.6654 143.226C9.6654 146.799 11.5714 150.1 14.6654 151.887L93 197.113C96.094 198.9 99.906 198.9 103 197.113L181.335 151.887C184.429 150.1 186.335 146.799 186.335 143.227L186.335 52.7735ZM98 7.547L176.335 52.7735L176.335 143.227L98 188.453L19.6654 143.227L19.6654 52.7735L98 7.547Z",
              opacity: "0",
              fill: "rgb(0,0,0)",
              "fill-rule": "evenodd"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "多边形 4",
              d:
                "M187.751 53.0235C187.751 49.4509 185.845 46.1496 182.751 44.3633L126.5 11.8868C123.406 10.1004 119.594 10.1004 116.5 11.8868L60.2491 44.3633C57.155 46.1496 55.2491 49.4509 55.2491 53.0235L55.2491 117.976C55.2491 121.549 57.155 124.85 60.2491 126.637L116.5 159.113C119.594 160.9 123.406 160.9 126.5 159.113L182.751 126.637C185.845 124.85 187.751 121.549 187.751 117.977L187.751 53.0235ZM121.5 20.547L177.751 53.0235L177.751 117.977L121.5 150.453L65.2491 117.976L65.2491 53.0235L121.5 20.547Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "evenodd"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "减去顶层",
              d:
                "M65.3843 32.3984C64.3012 31.3976 63.6228 29.9625 63.6228 28.3679L63.6228 24.2557C63.6228 22.2992 64.6641 20.4917 66.3513 19.5124L98.7539 0.734482C99.1687 0.493904 99.6074 0.311348 100.07 0.18681C100.533 0.0622673 101.003 0 101.482 0C101.961 0 102.432 0.0622673 102.895 0.18681C103.357 0.311348 103.796 0.493904 104.211 0.734482L112.204 5.36689L65.3843 32.3984ZM62.4229 34.1082L49.5879 26.67C49.1731 26.4294 48.7346 26.2469 48.272 26.1223C47.8093 25.9978 47.3384 25.9355 46.8594 25.9355C46.3806 25.9355 45.9097 25.9978 45.447 26.1223C44.9844 26.2469 44.5459 26.4294 44.1311 26.67L11.7285 45.4479C10.0413 46.4273 9 48.2347 9 50.1912L9 87.7449C9.00024 88.2256 9.06226 88.698 9.18628 89.1623C9.3103 89.6266 9.49219 90.0668 9.73169 90.4831C9.97095 90.8994 10.26 91.2775 10.5986 91.6174C10.9373 91.9574 11.3137 92.2476 11.7285 92.4881L44.1311 111.266C44.895 111.708 45.7373 111.951 46.5888 111.993L46.5888 100.042L19.9158 84.5856L19.9158 53.3548L46.8616 37.7405L51.4887 40.4218C51.5219 40.4021 51.5553 40.3826 51.5888 40.3633L62.4229 34.1082ZM60 137.493L60 155.903L50.8308 161.203L19.9397 143.347L19.9397 104.476C19.9397 101.451 17.4912 99 14.4697 99C11.4482 99 9 101.451 9 104.476L9 146.508Q9.00018 146.868 9.04682 147.222L9.04705 147.224L9.04706 147.224Q9.09375 147.578 9.18677 147.925C9.31104 148.389 9.49341 148.829 9.7334 149.246C9.97339 149.662 10.2629 150.04 10.6025 150.38C10.9419 150.72 11.3193 151.01 11.7349 151.25L48.0959 172.266C48.9426 172.755 49.8857 173 50.8308 173C51.8083 173.001 52.72 172.756 53.5659 172.266L60 168.547L60 173.806C60.0002 174.284 60.0623 174.754 60.1865 175.215C60.3105 175.677 60.4924 176.114 60.7319 176.528C60.9714 176.942 61.2607 177.317 61.5994 177.655C61.938 177.993 62.3149 178.282 62.7297 178.521L98.7703 199.271C99.6155 199.756 100.557 200 101.5 200C102.475 200 103.385 199.757 104.23 199.271L140.27 178.521C141.96 177.547 143 175.751 143 173.806L143 169.125L148.434 172.266C149.281 172.755 150.224 173 151.169 173C152.146 173.001 153.058 172.756 153.904 172.266L190.265 151.25C191.958 150.271 193 148.464 193 146.508L193 131.499C192.514 131.925 191.983 132.307 191.411 132.637L182.06 138.035L182.06 143.345L151.169 161.201L146.556 158.534L132.081 166.891L132.081 170.664L101.5 188.27L70.9187 170.664L70.9187 162.237L86.8792 153.012L75.9345 146.693L70.9187 149.592L70.9187 143.797L60 137.493Z",
              fill: "rgb(136,136,136)",
              "fill-rule": "evenodd"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "直线 1",
              d: "M0 0L40.0973 0",
              stroke: "rgb(136,136,136)",
              "stroke-linecap": "round",
              "stroke-width": "10.000000",
              transform:
                "matrix(0.885346,0.464932,-0.464932,0.885346,85.5,67.0601)"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "直线 1",
              d: "M0 0L40.0973 0",
              stroke: "rgb(136,136,136)",
              "stroke-linecap": "round",
              "stroke-width": "10.000000",
              transform:
                "matrix(-0.885347,0.464932,0.464932,0.885347,156.5,67.0601)"
            }
          }),
          _vm._v(" "),
          _c("path", {
            attrs: {
              id: "直线 1",
              d: "M0 0L40.3575 0",
              stroke: "rgb(136,136,136)",
              "stroke-linecap": "round",
              "stroke-width": "10.000000",
              transform: "matrix(0,-1,-1,0,121,126.06)"
            }
          })
        ]
      )
    : _vm._e()
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=template&id=64a91ffc&scoped=true":
/*!**********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Empty.vue?vue&type=template&id=64a91ffc&scoped=true ***!
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
  return _c("div", { staticClass: "no-data" }, [
    _c("div", { staticClass: "item-empty" }, [
      _c("div", { staticClass: "item-empty-icon" }),
      _vm._v(" "),
      _c("div", { staticClass: "item-empty-tips" }, [
        _vm._v(_vm._s(_vm.$t("org.no_result")))
      ])
    ])
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=template&id=56015720&scoped=true":
/*!************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Filters.vue?vue&type=template&id=56015720&scoped=true ***!
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
    { staticClass: "filters-container" },
    [
      _c(
        "el-input",
        {
          staticStyle: { width: "50%" },
          attrs: { placeholder: _vm.placeholder },
          nativeOn: {
            keyup: function($event) {
              if (
                !$event.type.indexOf("key") &&
                _vm._k($event.keyCode, "enter", 13, $event.key, "Enter")
              ) {
                return null
              }
              return _vm.handleSearch($event)
            }
          },
          model: {
            value: _vm.localKeyword,
            callback: function($$v) {
              _vm.localKeyword = $$v
            },
            expression: "localKeyword"
          }
        },
        [
          _c(
            "el-button",
            {
              attrs: { slot: "append" },
              on: { click: _vm.handleSearch },
              slot: "append"
            },
            [_vm._v(_vm._s(_vm.searchBtnText))]
          )
        ],
        1
      ),
      _vm._v(" "),
      _c(
        "el-dropdown",
        {
          ref: "filterOrder",
          staticClass: "filter-order",
          class: { "filter-order-active": _vm.isDropdownActive },
          attrs: { trigger: "click" },
          on: { "visible-change": _vm.onDropdownVisibleChange },
          nativeOn: {
            click: function($event) {
              return _vm.handleDropdownClick($event)
            }
          }
        },
        [
          _c("span", { staticClass: "el-dropdown-link" }, [
            _vm._v("\n      " + _vm._s(_vm.currentSortLabel)),
            _c("i", { staticClass: "el-icon-arrow-down el-icon--right" })
          ]),
          _vm._v(" "),
          _c(
            "el-dropdown-menu",
            {
              staticClass: "order-item",
              attrs: { slot: "dropdown" },
              slot: "dropdown"
            },
            _vm._l(_vm.sortOptions, function(item) {
              return _c(
                "el-dropdown-item",
                {
                  key: item.key,
                  class: { active: item.key === _vm.currentSortKey },
                  nativeOn: {
                    click: function($event) {
                      return _vm.selectSort(item)
                    }
                  }
                },
                [_vm._v("\n        " + _vm._s(item.label) + "\n      ")]
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
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Item.vue?vue&type=template&id=bbde0424&scoped=true":
/*!*********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/Item.vue?vue&type=template&id=bbde0424&scoped=true ***!
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
      staticClass: "item-card",
      on: {
        click: function($event) {
          $event.stopPropagation()
          $event.preventDefault()
          return _vm.getItemLink()
        }
      }
    },
    [
      _c("div", { staticClass: "card-part1" }, [
        _c("div", { staticClass: "part1-content" }, [
          _c(
            "div",
            { staticClass: "part1-content-display" },
            [
              _vm.data.type !== "aitasktmpl"
                ? _c("Icons", {
                    staticStyle: {
                      "flex-shrink": "0",
                      width: "24px",
                      height: "24px"
                    },
                    attrs: {
                      type: _vm.data.type,
                      isPrivate: _vm.data.is_private
                    }
                  })
                : _c("div", { staticClass: "icon-c" }, [
                    _c(
                      "svg",
                      {
                        attrs: {
                          xmlns: "http://www.w3.org/2000/svg",
                          viewBox: "0 0 48 48",
                          width: "24",
                          height: "24"
                        }
                      },
                      [
                        _c("defs"),
                        _vm._v(" "),
                        _c("g", [
                          _c("path", {
                            attrs: {
                              d: "M48 0H0V48H48V0Z",
                              "fill-opacity": "0.01"
                            }
                          }),
                          _vm._v(" "),
                          _c("path", {
                            attrs: {
                              d: "M44 14L24 4L4 14V34L24 44L44 34V14Z",
                              fill: "none",
                              stroke: "#888",
                              "stroke-width": "4",
                              "stroke-linejoin": "round"
                            }
                          }),
                          _vm._v(" "),
                          _c("path", {
                            attrs: {
                              d: "M4 14L24 24",
                              fill: "none",
                              stroke: "#888",
                              "stroke-width": "4",
                              "stroke-linecap": "round",
                              "stroke-linejoin": "round"
                            }
                          }),
                          _vm._v(" "),
                          _c("path", {
                            attrs: {
                              d: "M24 44V24",
                              fill: "none",
                              stroke: "#888",
                              "stroke-width": "4",
                              "stroke-linecap": "round",
                              "stroke-linejoin": "round"
                            }
                          }),
                          _vm._v(" "),
                          _c("path", {
                            attrs: {
                              d: "M44 14L24 24",
                              fill: "none",
                              stroke: "#888",
                              "stroke-width": "4",
                              "stroke-linecap": "round",
                              "stroke-linejoin": "round"
                            }
                          }),
                          _vm._v(" "),
                          _c("path", {
                            attrs: {
                              d: "M34 9L14 19",
                              fill: "none",
                              stroke: "#888",
                              "stroke-width": "4",
                              "stroke-linecap": "round",
                              "stroke-linejoin": "round"
                            }
                          })
                        ])
                      ]
                    ),
                    _vm._v(" "),
                    _vm.data.IsPrivate
                      ? _c(
                          "svg",
                          {
                            staticClass: "lock",
                            attrs: {
                              xmlns: "http://www.w3.org/2000/svg",
                              viewBox: "0 0 32 32",
                              width: "12",
                              height: "12"
                            }
                          },
                          [
                            _c("defs"),
                            _vm._v(" "),
                            _c("g", [
                              _c("path", {
                                attrs: {
                                  d:
                                    "M25.333 13.333h1.333c0.736 0 1.333 0.597 1.333 1.333v0 13.333c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-13.333c0-0.736 0.597-1.333 1.333-1.333v0h1.333v-1.333c0-5.155 4.179-9.333 9.333-9.333s9.333 4.179 9.333 9.333v0 1.333zM6.667 16v10.667h18.667v-10.667h-18.667zM14.667 18.667h2.667v5.333h-2.667v-5.333zM22.667 13.333v-1.333c0-3.682-2.985-6.667-6.667-6.667s-6.667 2.985-6.667 6.667v0 1.333h13.333z",
                                  fill: "#888"
                                }
                              })
                            ])
                          ]
                        )
                      : _vm._e()
                  ]),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "item-title", attrs: { title: _vm.data.alias } },
                [_vm._v(_vm._s(_vm.data.alias))]
              ),
              _vm._v(" "),
              _vm.data.recommend
                ? _c("div", { staticClass: "item-recommend" }, [
                    _c(
                      "svg",
                      {
                        staticClass:
                          "styles__StyledSVGIconPathComponent-sc-i3aj97-0 dZJqQS svg-icon-path-icon fill",
                        attrs: {
                          xmlns: "http://www.w3.org/2000/svg",
                          viewBox: "0 0 24 24",
                          width: "20",
                          height: "20"
                        }
                      },
                      [
                        _c("defs"),
                        _c("g", [
                          _c("path", {
                            attrs: {
                              d:
                                "M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zm4.24 16L12 15.45 7.77 18l1.12-4.81-3.73-3.23 4.92-.42L12 5l1.92 4.53 4.92.42-3.73 3.23L16.23 18z"
                            }
                          })
                        ])
                      ]
                    )
                  ])
                : _vm._e()
            ],
            1
          ),
          _vm._v(" "),
          _vm.data.type == "aitasktmpl" && !_vm.isMobile
            ? _c(
                "div",
                { staticClass: "aitasktmpl-content" },
                [
                  _c(
                    "el-row",
                    { attrs: { gutter: 20 } },
                    [
                      _c("el-col", { attrs: { span: 12 } }, [
                        _c("div", { staticClass: "aitasktmpl-item" }, [
                          _c("span", { staticClass: "aitasktmpl-subtitle" }, [
                            _vm._v(
                              _vm._s(_vm.$t("cloudbrainObj.dataset")) + "："
                            )
                          ]),
                          _vm._v(" "),
                          (_vm.data.DatasetLists || []).length
                            ? _c("span", { staticClass: "val" }, [
                                _c(
                                  "span",
                                  { attrs: { title: _vm.data.DatasetsStr } },
                                  [_vm._v(_vm._s(_vm.data.DatasetsStr))]
                                )
                              ])
                            : _c("span", { staticClass: "val" }, [_vm._v("--")])
                        ]),
                        _vm._v(" "),
                        _c("div", { staticClass: "aitasktmpl-item" }, [
                          _c("span", { staticClass: "aitasktmpl-subtitle" }, [
                            _vm._v(_vm._s(_vm.$t("repos.repos")) + "：")
                          ]),
                          _vm._v(" "),
                          _vm.data.RepoOwnerName && _vm.data.RepoName
                            ? _c(
                                "span",
                                {
                                  staticClass: "val",
                                  attrs: {
                                    title:
                                      _vm.data.RepoOwnerName +
                                      "/" +
                                      _vm.data.RepoName
                                  }
                                },
                                [
                                  _vm._v(
                                    _vm._s(
                                      _vm.data.RepoOwnerName +
                                        "/" +
                                        _vm.data.RepoName
                                    )
                                  )
                                ]
                              )
                            : _c("span", { staticClass: "val" }, [_vm._v("--")])
                        ])
                      ]),
                      _vm._v(" "),
                      _c("el-col", { attrs: { span: 12 } }, [
                        _c("div", { staticClass: "aitasktmpl-item" }, [
                          _c("span", { staticClass: "aitasktmpl-subtitle" }, [
                            _vm._v(_vm._s(_vm.$t("repos.model")) + "：")
                          ]),
                          _vm._v(" "),
                          (_vm.data.ModelLists || []).length
                            ? _c("span", { staticClass: "val" }, [
                                _c(
                                  "span",
                                  { attrs: { title: _vm.data.ModelsStr } },
                                  [_vm._v(" " + _vm._s(_vm.data.ModelsStr))]
                                )
                              ])
                            : _c("span", { staticClass: "val" }, [_vm._v("--")])
                        ]),
                        _vm._v(" "),
                        _c("div", { staticClass: "aitasktmpl-item" }, [
                          _c("span", { staticClass: "aitasktmpl-subtitle" }, [
                            _vm._v(_vm._s(_vm.$t("cloudbrainObj.image")) + "：")
                          ]),
                          _vm._v(" "),
                          _vm.data.ImageName || _vm.data.ImageUrl
                            ? _c(
                                "span",
                                {
                                  staticClass: "val",
                                  attrs: { title: _vm.data.ImageName }
                                },
                                [
                                  _vm._v(
                                    _vm._s(
                                      _vm.data.ImageName || _vm.data.ImageUrl
                                    )
                                  )
                                ]
                              )
                            : _c("span", { staticClass: "val" }, [_vm._v("--")])
                        ])
                      ])
                    ],
                    1
                  )
                ],
                1
              )
            : _vm._e(),
          _vm._v(" "),
          _vm.data.type == "aitasktmpl" && _vm.isMobile
            ? _c(
                "div",
                { staticClass: "aitasktmpl-content" },
                [
                  _c("el-row", { attrs: { gutter: 20 } }, [
                    _c("div", { staticClass: "aitasktmpl-item" }, [
                      _c("span", { staticClass: "aitasktmpl-subtitle" }, [
                        _vm._v(_vm._s(_vm.$t("cloudbrainObj.dataset")) + "：")
                      ]),
                      _vm._v(" "),
                      (_vm.data.DatasetLists || []).length
                        ? _c("span", { staticClass: "val" }, [
                            _c(
                              "span",
                              { attrs: { title: _vm.data.DatasetsStr } },
                              [_vm._v(_vm._s(_vm.data.DatasetsStr))]
                            )
                          ])
                        : _c("span", { staticClass: "val" }, [_vm._v("--")])
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "aitasktmpl-item" }, [
                      _c("span", { staticClass: "aitasktmpl-subtitle" }, [
                        _vm._v(_vm._s(_vm.$t("repos.repos")) + "：")
                      ]),
                      _vm._v(" "),
                      _vm.data.RepoOwnerName && _vm.data.RepoName
                        ? _c(
                            "span",
                            {
                              staticClass: "val",
                              attrs: {
                                title:
                                  _vm.data.RepoOwnerName +
                                  "/" +
                                  _vm.data.RepoName
                              }
                            },
                            [
                              _vm._v(
                                _vm._s(
                                  _vm.data.RepoOwnerName +
                                    "/" +
                                    _vm.data.RepoName
                                )
                              )
                            ]
                          )
                        : _c("span", { staticClass: "val" }, [_vm._v("--")])
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "aitasktmpl-item" }, [
                      _c("span", { staticClass: "aitasktmpl-subtitle" }, [
                        _vm._v(_vm._s(_vm.$t("repos.model")) + "：")
                      ]),
                      _vm._v(" "),
                      (_vm.data.ModelLists || []).length
                        ? _c("span", { staticClass: "val" }, [
                            _c(
                              "span",
                              { attrs: { title: _vm.data.ModelsStr } },
                              [_vm._v(" " + _vm._s(_vm.data.ModelsStr))]
                            )
                          ])
                        : _c("span", { staticClass: "val" }, [_vm._v("--")])
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "aitasktmpl-item" }, [
                      _c("span", { staticClass: "aitasktmpl-subtitle" }, [
                        _vm._v(_vm._s(_vm.$t("cloudbrainObj.image")) + "：")
                      ]),
                      _vm._v(" "),
                      _vm.data.ImageName || _vm.data.ImageUrl
                        ? _c(
                            "span",
                            {
                              staticClass: "val",
                              attrs: { title: _vm.data.ImageName }
                            },
                            [
                              _vm._v(
                                _vm._s(_vm.data.ImageName || _vm.data.ImageUrl)
                              )
                            ]
                          )
                        : _c("span", { staticClass: "val" }, [_vm._v("--")])
                    ])
                  ])
                ],
                1
              )
            : _vm._e(),
          _vm._v(" "),
          _c("div", { staticClass: "part1-content-display tag-part" }, [
            _vm.data.type == "aitasktmpl"
              ? _c("div", [
                  _c("div", { staticClass: "aitasktmpl-tag" }, [
                    _c("div", [_vm._v(_vm._s(_vm.data.JobTypeStr))]),
                    _vm._v(" "),
                    _c("div", [_vm._v(_vm._s(_vm.data.ComputeSourceStr))])
                  ])
                ])
              : _c(
                  "div",
                  { staticStyle: { display: "flex" } },
                  [
                    _vm._l(_vm.data.tags, function(item, index) {
                      return _c("div", { key: item }, [
                        _vm.data.type == "dataset"
                          ? _c("span", { staticClass: "item-tag" }, [
                              _vm._v(_vm._s(_vm.$t("datasets." + item)))
                            ])
                          : _vm.data.type == "model"
                          ? _c("span", { staticClass: "item-tag" }, [
                              _vm._v(_vm._s(item))
                            ])
                          : _vm._e()
                      ])
                    }),
                    _vm._v(" "),
                    _vm.data.type == "dataset"
                      ? _c(
                          "div",
                          _vm._l(_vm.data.tasks, function(item, index) {
                            return _c("div", { key: item }, [
                              _c("span", { staticClass: "item-tag" }, [
                                _vm._v(_vm._s(_vm.$t("datasets." + item)))
                              ])
                            ])
                          }),
                          0
                        )
                      : _vm._e(),
                    _vm._v(" "),
                    _c("div", [
                      _vm.data.licenses
                        ? _c("span", { staticClass: "item-tag" }, [
                            _vm._v(" " + _vm._s(_vm.data.licenses) + " ")
                          ])
                        : _vm._e()
                    ]),
                    _vm._v(" "),
                    _c("div", [
                      _vm.data.engineName
                        ? _c("span", { staticClass: "item-tag" }, [
                            _vm._v(_vm._s(_vm.data.engineName))
                          ])
                        : _vm._e()
                    ])
                  ],
                  2
                )
          ])
        ])
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "card-part2" }, [
        _c("div", { staticClass: "time-size" }, [
          _c(
            "span",
            {
              staticStyle: { "margin-right": "5px" },
              attrs: {
                title: _vm.$t("repos.updated") + " " + _vm.data.updated_time
              }
            },
            [
              _vm._v(
                " \n        " +
                  _vm._s(_vm.$t("repos.updated")) +
                  " " +
                  _vm._s(_vm.data.updated_time) +
                  "\n      "
              )
            ]
          ),
          _vm._v(" "),
          _vm.data.type != "aitasktmpl"
            ? _c(
                "span",
                {
                  staticClass: "size",
                  attrs: {
                    title: _vm.$t("datasets.size") + " " + _vm.data.size
                  }
                },
                [
                  _c("span", [_vm._v(_vm._s(_vm.$t("datasets.size")) + "：")]),
                  _vm._v(" "),
                  _c("span", { staticStyle: { color: "rgba(16,16,16,1)" } }, [
                    _vm._v(_vm._s(_vm.data.size))
                  ])
                ]
              )
            : _vm._e()
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "item-link" }, [
          _c("div", { attrs: { title: _vm.$t("datasets.moststars") } }, [
            _c("i", {
              class: _vm.data.is_collected ? "heart icon" : "heart outline icon"
            }),
            _vm._v(" "),
            _c("span", [_vm._v(_vm._s(_vm.data.num_stars))])
          ]),
          _vm._v(" "),
          _c(
            "div",
            {
              staticClass: "r-item",
              attrs: {
                title:
                  _vm.data.type != "aitasktmpl"
                    ? _vm.$t("datasets.citations")
                    : _vm.$t("taskTmplObj.runTimes")
              }
            },
            [
              _vm.data.type != "aitasktmpl"
                ? _c("i", { staticClass: "el-icon-link" })
                : _c(
                    "svg",
                    {
                      attrs: {
                        xmlns: "http://www.w3.org/2000/svg",
                        fill: "rgb(136, 136, 136)",
                        viewBox: "0 0 48 48",
                        width: "12",
                        height: "12"
                      }
                    },
                    [
                      _c("defs"),
                      _vm._v(" "),
                      _c("g", [
                        _c("rect", {
                          attrs: {
                            width: "48",
                            height: "48",
                            "fill-opacity": "0.01"
                          }
                        }),
                        _vm._v(" "),
                        _c("path", {
                          attrs: {
                            d:
                              "M43.8233 25.2305C43.7019 25.9889 43.5195 26.727 43.2814 27.4395C42.763 28.9914 41.9801 30.4222 40.9863 31.6785C38.4222 34.9201 34.454 37 30 37H16C9.39697 37 4 31.6785 4 25C4 18.3502 9.39624 13 16 13L44 13",
                            stroke: "rgb(136, 136, 136)",
                            fill: "none",
                            "stroke-width": "4",
                            "stroke-linecap": "round",
                            "stroke-linejoin": "round"
                          }
                        }),
                        _vm._v(" "),
                        _c("path", {
                          attrs: {
                            d: "M38 7L44 13L38 19",
                            fill: "none",
                            stroke: "rgb(136, 136, 136)",
                            "stroke-width": "4",
                            "stroke-linecap": "round",
                            "stroke-linejoin": "round"
                          }
                        })
                      ])
                    ]
                  ),
              _vm._v(" "),
              _c("span", [_vm._v(_vm._s(_vm.data.use_count))])
            ]
          ),
          _vm._v(" "),
          _vm.data.type != "aitasktmpl"
            ? _c(
                "div",
                {
                  staticClass: "r-item",
                  attrs: { title: _vm.$t("datasets.downloadtimes") }
                },
                [
                  _c("i", { staticClass: "el-icon-download" }),
                  _vm._v(" "),
                  _c("span", [_vm._v(_vm._s(_vm.data.download_count))])
                ]
              )
            : _vm._e(),
          _vm._v(" "),
          _vm.data.type == "model"
            ? _c(
                "div",
                {
                  staticClass: "r-item",
                  attrs: { title: _vm.$t("modelManage.derivativeTimes") }
                },
                [
                  _c("i", { staticClass: "ri-git-merge-line" }),
                  _vm._v(" "),
                  _c("span", [_vm._v(_vm._s(_vm.data.derivative_count))])
                ]
              )
            : _vm._e(),
          _vm._v(" "),
          _vm.data.type == "aitasktmpl"
            ? _c(
                "span",
                {
                  staticClass: "run-btn",
                  attrs: { title: "" },
                  on: {
                    click: function($event) {
                      $event.stopPropagation()
                      $event.preventDefault()
                      return _vm.goRun(_vm.data)
                    }
                  }
                },
                [
                  _c(
                    "svg",
                    {
                      attrs: {
                        xmlns: "http://www.w3.org/2000/svg",
                        fill: "rgb(255, 255, 255)",
                        viewBox: "0 0 48 48",
                        width: "12",
                        height: "12"
                      }
                    },
                    [
                      _c("defs"),
                      _vm._v(" "),
                      _c("g", [
                        _c("path", {
                          attrs: {
                            d:
                              "M24 44C12.9543 44 4 35.0457 4 24C4 12.9543 12.9543 4 24 4C35.0457 4 44 12.9543 44 24",
                            fill: "none",
                            stroke: "rgb(255, 255, 255)",
                            "stroke-width": "4",
                            "stroke-linecap": "round",
                            "stroke-linejoin": "round"
                          }
                        }),
                        _vm._v(" "),
                        _c("path", {
                          attrs: {
                            d:
                              "M20 24V17.0718L26 20.5359L32 24L26 27.4641L20 30.9282V24Z",
                            fill: "none",
                            stroke: "rgb(255, 255, 255)",
                            "stroke-width": "4",
                            "stroke-linejoin": "round"
                          }
                        }),
                        _vm._v(" "),
                        _c("path", {
                          attrs: {
                            d: "M37.0508 32L37.0508 42",
                            fill: "none",
                            stroke: "rgb(255, 255, 255)",
                            "stroke-width": "4",
                            "stroke-linecap": "round",
                            "stroke-linejoin": "round"
                          }
                        }),
                        _vm._v(" "),
                        _c("path", {
                          attrs: {
                            d: "M42 36.9497L32 36.9497",
                            fill: "none",
                            stroke: "rgb(255, 255, 255)",
                            "stroke-width": "4",
                            "stroke-linecap": "round",
                            "stroke-linejoin": "round"
                          }
                        })
                      ])
                    ]
                  ),
                  _vm._v(" "),
                  _c("span", [_vm._v(_vm._s(_vm.$t("taskTmplObj.run")))])
                ]
              )
            : _vm._e()
        ])
      ])
    ]
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/List.vue?vue&type=template&id=a7d3140e":
/*!*********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/components/List.vue?vue&type=template&id=a7d3140e ***!
  \*********************************************************************************************************************************************************************************************************************************/
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
    _vm._l(_vm.params, function(item, index) {
      return _c(
        "div",
        { key: item.id },
        [_c("Item", { key: item.ID, attrs: { data: item } })],
        1
      )
    }),
    0
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/compute/index.vue?vue&type=template&id=2bf59048&scoped=true":
/*!*******************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/profile/compute/index.vue?vue&type=template&id=2bf59048&scoped=true ***!
  \*******************************************************************************************************************************************************************************************************************************************/
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
      staticClass: "profile-compute"
    },
    [
      _c("Filters", {
        attrs: {
          placeholder: _vm.$t("org.searching"),
          "search-btn-text": _vm.$t("repos.search"),
          "sort-options": _vm.sortList,
          "default-sort-key": _vm.order_by
        },
        on: { search: _vm.onFilterSearch }
      }),
      _vm._v(" "),
      _vm.currentList.length
        ? _c("List", { attrs: { params: _vm.currentList } })
        : _vm._e(),
      _vm._v(" "),
      !_vm.currentList.length && !_vm.loading ? _c("Empty") : _vm._e(),
      _vm._v(" "),
      _c(
        "div",
        {
          directives: [
            {
              name: "show",
              rawName: "v-show",
              value: _vm.currentList.length,
              expression: "currentList.length"
            }
          ],
          staticClass: "center"
        },
        [
          _c("el-pagination", {
            ref: "paginationRef",
            attrs: {
              background: "",
              "current-page": _vm.page,
              "pager-count": 5,
              "page-sizes": _vm.page_sizes,
              "page-size": _vm.page_size,
              layout: "total, prev, pager, next, jumper",
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
                _vm.page_size = $event
              },
              "update:page-size": function($event) {
                _vm.page_size = $event
              }
            }
          })
        ],
        1
      )
    ],
    1
  )
}
var staticRenderFns = []
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
/******/ 			"vp-profile-compute": 0
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
/******/ 	var __webpack_exports__ = __webpack_require__.O(undefined, ["vendor-libs"], function() { return __webpack_require__("./web_src/vuepages/pages/profile/compute/vp-profile-compute.js"); })
/******/ 	__webpack_exports__ = __webpack_require__.O(__webpack_exports__);
/******/ 	
/******/ })()
;