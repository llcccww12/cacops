/******/ (function() { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

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

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseTitle.vue?vue&type=script&lang=js":
/*!*******************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseTitle.vue?vue&type=script&lang=js ***!
  \*******************************************************************************************************************************************************************************************************/
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
  name: 'BaseTitle',
  props: {
    title: {
      type: String,
      "default": ''
    }
  },
  data: function data() {
    return {};
  },
  mounted: function mounted() {},
  methods: {}
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

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=script&lang=js":
/*!**************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=script&lang=js ***!
  \**************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var _components_BaseDialog_vue__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ~/components/BaseDialog.vue */ "./web_src/vuepages/components/BaseDialog.vue");
/* harmony import */ var _apis_modules_common__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ~/apis/modules/common */ "./web_src/vuepages/apis/modules/common.js");
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! spark-md5 */ "./node_modules/spark-md5/spark-md5.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(spark_md5__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var highlight_js__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! highlight.js */ "./node_modules/highlight.js/lib/index.js");
/* harmony import */ var highlight_js__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(highlight_js__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");








//
//
//
//
//
//
//
//
//
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
  name: "CommonTipsDialog",
  components: {
    BaseDialog: _components_BaseDialog_vue__WEBPACK_IMPORTED_MODULE_8__["default"]
  },
  props: {
    title: {
      type: String,
      "default": 'Title'
    },
    appendToBody: {
      type: Boolean,
      "default": false
    },
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    closeText: {
      type: String,
      "default": ''
    }
  },
  data: function data() {
    return {
      showClose: true,
      visible: false,
      loading: false,
      content: ''
    };
  },
  methods: {
    sparkMD5Hash: function sparkMD5Hash() {
      var str = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : '';
      return spark_md5__WEBPACK_IMPORTED_MODULE_11___default().hash(str) + Math.random().toString().replace('0.', '');
    },
    hljsAndInsertCopyButton: function hljsAndInsertCopyButton(htmlStr) {
      var html = document.createElement('div');
      html.innerHTML = htmlStr;
      var codeBlocks = html.querySelectorAll('.code-block');

      for (var i = 0, iLen = codeBlocks.length; i < iLen; i++) {
        var codeBlockI = codeBlocks[i];
        var txt = codeBlockI.textContent.slice(0, -1);
        var codeEl = codeBlockI.querySelector('code');
        var code = codeEl.textContent;
        codeEl.innerHTML = highlight_js__WEBPACK_IMPORTED_MODULE_12___default().highlight('python', code).value;
        var copyBtn = document.createElement('div');
        copyBtn.classList = ['copy-btn'];
        copyBtn.innerHTML = "<span style=\"color: #0366d6;cursor: pointer;\" class=\"ui poping inline up clipboard\" id=\"clipboard-".concat(this.sparkMD5Hash(txt), "\"\n            data-position=\"top center\" data-variation=\"inverted tiny\" data-success=\"").concat(this.$t('copySuccess'), "\"\n            data-content=\"").concat(this.$t('copy'), "\" data-original=\"").concat(this.$t('copy'), "\" \n            data-clipboard-text=\"\"><i style=\"font-size:14px;\" class=\"copy outline icon\"></i></span>");
        copyBtn.querySelector('span').setAttribute('data-clipboard-text', txt);
        codeBlockI.outerHTML = "<div class=\"code-content\">".concat(codeBlockI.outerHTML).concat(copyBtn.outerHTML, "</div>");
      }

      return html.innerHTML;
    },
    getMarkdown: function getMarkdown(str) {
      var _this = this;

      (0,_apis_modules_common__WEBPACK_IMPORTED_MODULE_9__.getMarkdownHtml)(str).then(function (res) {
        _this.loading = false;
        var html = res.data;
        _this.content = _this.hljsAndInsertCopyButton(html);

        _this.$nextTick(function () {
          (0,_utils__WEBPACK_IMPORTED_MODULE_13__.initClipboard)('.base-dlg .clipboard');
        });
      })["catch"](function (err) {
        _this.loading = false;
        console.log(err);
      });
    },
    getSdkContent: function getSdkContent(size) {
      var _this2 = this;

      var contentStr;
      (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_10__.getFileSdkCode)(this.data).then(function (res) {
        res = res.data;

        if (res.code == 0) {
          contentStr = _this2.$t('modelObj.model_sdk_tips1', {
            fileSize: (0,_utils__WEBPACK_IMPORTED_MODULE_13__.transFileSize)(size)
          });
          contentStr += "\n```python\n".concat(res.data.code, "```\n\n").concat(_this2.$t('modelObj.model_sdk_tips2'), "\n```bash\n").concat(res.data.cli, "\n```\n");
          contentStr += _this2.$t('modelObj.model_sdk_tips3');

          _this2.getMarkdown(contentStr);
        } else {
          _this2.$message.error(res.msg);
        }
      })["catch"](function (err) {
        _this2.$message.error(err);
      });
    },
    handlerOpen: function handlerOpen(event) {
      var size = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : 0;
      this.getSdkContent(size);
      this.visible = true;
    },
    closeDialog: function closeDialog() {
      var _this3 = this;

      this.visible = false;
      setTimeout(function () {
        _this3.content = '';
      }, 300);
    }
  },
  mounted: function mounted() {},
  beforeMount: function beforeMount() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Header.vue?vue&type=script&lang=js":
/*!******************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Header.vue?vue&type=script&lang=js ***!
  \******************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var _Icons_vue__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./Icons.vue */ "./web_src/vuepages/components/square/detail/Icons.vue");
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");




//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: "Header",
  props: {
    dataObj: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    tabList: {
      type: Array,
      "default": function _default() {
        return [];
      }
    },
    tab: {
      type: String,
      "default": 'intro'
    },
    type: {
      type: String,
      "default": 'dataset'
    }
  },
  components: {
    Icons: _Icons_vue__WEBPACK_IMPORTED_MODULE_4__["default"]
  },
  data: function data() {
    return {
      isCollected: false,
      collected_count: 0,
      isSetting: false,
      isLogin: false,
      tabIndex: this.tab,
      lang: _langs__WEBPACK_IMPORTED_MODULE_7__.lang
    };
  },
  methods: {
    changeFav: function changeFav(item) {
      var _this = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_3___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0___default().mark(function _callee() {
        var res, params, redirect;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_0___default().wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                if (!_this.isSetting) {
                  _context.next = 2;
                  break;
                }

                return _context.abrupt("return");

              case 2:
                _this.isSetting = true;
                _context.prev = 3;
                params = _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_1___default()({}, _this.type === 'dataset' ? 'dataset_id' : 'aimodel_id', item.id);

                if (!_this.isCollected) {
                  _context.next = 11;
                  break;
                }

                _context.next = 8;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_5__.unsetFavorite)(params, _this.type);

              case 8:
                res = _context.sent;
                _context.next = 14;
                break;

              case 11:
                _context.next = 13;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_5__.setFavorite)(params, _this.type);

              case 13:
                res = _context.sent;

              case 14:
                if (!(res.data.code === 401)) {
                  _context.next = 18;
                  break;
                }

                redirect = encodeURIComponent(window.location.href);
                window.location.href = "/user/login?redirect_to=".concat(redirect);
                return _context.abrupt("return");

              case 18:
                if (!(res.data.code === 0)) {
                  _context.next = 25;
                  break;
                }

                if (_this.isCollected) {
                  _this.collected_count -= 1;
                } else {
                  _this.collected_count += 1;
                }

                _this.isCollected = !_this.isCollected;

                _this.$message.success(_this.isCollected ? _this.$t('datasets.starSuccess') : _this.$t('datasets.unstarSuccess')); // 通知父组件状态变化


                _this.$emit('changeFav');

                _context.next = 26;
                break;

              case 25:
                throw new Error(res.data.msg || '操作失败');

              case 26:
                _context.next = 32;
                break;

              case 28:
                _context.prev = 28;
                _context.t0 = _context["catch"](3);
                console.error('操作收藏失败:', _context.t0);

                _this.$message.error(_context.t0);

              case 32:
                _context.prev = 32;
                _this.isSetting = false;
                return _context.finish(32);

              case 35:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, null, [[3, 28, 32, 35]]);
      }))();
    },
    createOnlineDebug: function createOnlineDebug() {},
    changeTab: function changeTab(item) {
      this.tabIndex = item.key;
      this.$emit('changeTab', item);
    }
  },
  watch: {},
  mounted: function mounted() {
    this.$nextTick(function () {
      (0,_utils__WEBPACK_IMPORTED_MODULE_6__.initClipboard)('.copy-btn .clipboard');
    });
    this.isCollected = this.dataObj.is_collected;
    this.collected_count = this.dataObj.num_stars;
    this.isLogin = !!document.querySelector('meta[name="_uid"]');
    console.log(this.dataObj);
  },
  beforeMount: function beforeMount() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Icons.vue?vue&type=script&lang=js":
/*!*****************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Icons.vue?vue&type=script&lang=js ***!
  \*****************************************************************************************************************************************************************************************************************/
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
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=script&lang=js":
/*!**************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=script&lang=js ***!
  \**************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.iterator */ "./node_modules/core-js/modules/es.array.iterator.js");
/* harmony import */ var core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.join */ "./node_modules/core-js/modules/es.array.join.js");
/* harmony import */ var core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_join__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.sort */ "./node_modules/core-js/modules/es.array.sort.js");
/* harmony import */ var core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_sort__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_map__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.map */ "./node_modules/core-js/modules/es.map.js");
/* harmony import */ var core_js_modules_es_map__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_map__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.string.iterator */ "./node_modules/core-js/modules/es.string.iterator.js");
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/web.dom-collections.iterator */ "./node_modules/core-js/modules/web.dom-collections.iterator.js");
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_web_url__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/web.url */ "./node_modules/core-js/modules/web.url.js");
/* harmony import */ var core_js_modules_web_url__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_url__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! @babel/runtime/helpers/toConsumableArray */ "./node_modules/@babel/runtime/helpers/toConsumableArray.js");
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var _pages_guide_components_FileUpload_vue__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! ~/pages/guide/components/FileUpload.vue */ "./web_src/vuepages/pages/guide/components/FileUpload.vue");
/* harmony import */ var _components_square_CommonSDK_vue__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! ~/components/square/CommonSDK.vue */ "./web_src/vuepages/components/square/CommonSDK.vue");
/* harmony import */ var _components_BaseDialog_vue__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! ~/components/BaseDialog.vue */ "./web_src/vuepages/components/BaseDialog.vue");
/* harmony import */ var _pages_guide_components_CreateForm_vue__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! ~/pages/guide/components/CreateForm.vue */ "./web_src/vuepages/pages/guide/components/CreateForm.vue");
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_24__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");


















//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//







var maxModelFilesSize = window.MAX_MODEL_SIZE || 1024 * 1024 * 1024 * 0.5;
/* harmony default export */ __webpack_exports__["default"] = ({
  name: 'AiforgeFileList1',
  props: {
    dataObj: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    type: {
      type: String,
      "default": 'dataset'
    }
  },
  data: function data() {
    return {
      emptyPage: false,
      filesList: [],
      loading: false,
      params: {
        parent_dir: '',
        page_size: 10,
        marker: ''
      },
      has_next: false,
      filePath: [this.dataObj.name],
      showUploadPage: false,
      visible: false,
      title: '',
      previewCache: new Map(),
      // 使用Map来缓存预览数据
      fileImg: '',
      textContent: '',
      supportImgReg: /(\.jpg|\.jpeg|\.png|\.gif|\.bmp|\.webp)$/i,
      supportTxtReg: /(\.txt|\.xml|\.html|\.json|\.py|\.sh|\.md|\.csv|\.log|\.js|\.css|\.ipynb)$/i,
      sdkData: {},
      codeUsePromotePath: "tips/modelDown/sdkcode".concat(_langs__WEBPACK_IMPORTED_MODULE_24__.lang == 'zh-CN' ? '' : '_en', ".md"),
      showFlag: false,
      showMessage: ''
    };
  },
  components: {
    FileUpload: _pages_guide_components_FileUpload_vue__WEBPACK_IMPORTED_MODULE_18__["default"],
    CreateForm: _pages_guide_components_CreateForm_vue__WEBPACK_IMPORTED_MODULE_21__["default"],
    CommonSDK: _components_square_CommonSDK_vue__WEBPACK_IMPORTED_MODULE_19__["default"],
    BaseDialog: _components_BaseDialog_vue__WEBPACK_IMPORTED_MODULE_20__["default"]
  },
  mounted: function mounted() {
    if (this.type === 'dataset') {
      this.params.dataset_id = this.dataObj.id;
    } else {
      this.params.aimodel_id = this.dataObj.id;
    }

    this.getDatasetFileList();
  },
  methods: {
    handleCopy: function handleCopy(e) {
      e.preventDefault(); // 构造要复制的文本

      var textToCopy = this.filePath.join('/'); // 将文本放入剪贴板

      e.clipboardData.setData('text/plain', textToCopy);
    },
    cancel: function cancel() {
      this.showUploadPage = false;
    },
    uploadFinish: function uploadFinish() {
      this.showUploadPage = false;
      this.getDatasetFileList();
      this.$emit('editSuccess');
    },
    sortChange: function sortChange(options) {
      var prop = options.prop,
          order = options.order;
      console.log(prop);

      if (!order) {
        prop = 'FileName';
        order = 'ascending';
      }

      this.filesList.sort(function (a, b) {
        if (a.IsDir != b.IsDir && order == 'ascending') {
          return a.IsDir ? -1 : 1;
        }

        if (a.IsDir != b.IsDir && order == 'descending') {
          return a.IsDir ? 1 : -1;
        }

        if (prop == 'FileName' && order == 'ascending') {
          return a.FileName.toLocaleLowerCase().localeCompare(b.FileName.toLocaleLowerCase());
        }

        if (prop == 'FileName' && order == 'descending') {
          return b.FileName.toLocaleLowerCase().localeCompare(a.FileName.toLocaleLowerCase());
        }

        if (prop == 'Size' && order == 'ascending') {
          return a.Size - b.Size;
        }

        if (prop == 'Size' && order == 'descending') {
          return b.Size - a.Size;
        }

        if (prop == 'ModTime' && order == 'ascending') {
          return a.ModTime.localeCompare(b.ModTime);
        }

        if (prop == 'ModTime' && order == 'descending') {
          return b.ModTime.localeCompare(a.ModTime);
        }

        return 0;
      });
    },
    goNextDir: function goNextDir(row) {
      this.params.parent_dir = "".concat(row.ParenDir, "/").concat(row.FileName);
      this.filePath.push(row.FileName);
      this.params.marker = ''; // 重置分页标记

      this.getDatasetFileList();
    },
    goBackDir: function goBackDir(index) {
      if (index === this.filePath.length - 1) return; // 计算要返回的路径

      var pathParts = this.filePath.slice(0, index + 1);
      this.filePath = pathParts; // 如果是根目录

      if (index === 0) {
        this.params.parent_dir = '';
      } else {
        // 找到对应的目录对象
        var result = '/' + pathParts.slice(1).join('/');
        this.params.parent_dir = result;
      }

      this.params.marker = ''; // 重置分页标记

      this.getDatasetFileList();
    },
    previewFile: function previewFile(file) {
      var _this = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().mark(function _callee() {
        var cacheKey, cachedData, _params, params, queryParams, response;

        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                _this.visible = true;
                _this.title = file.FileName;
                cacheKey = "".concat(file.ParenDir, "/").concat(file.FileName);
                console.log("cacheKey", cacheKey);
                console.log("previewCache", _this.previewCache); // 检查缓存中是否已有数据

                if (!_this.previewCache.has(cacheKey)) {
                  _context.next = 9;
                  break;
                }

                cachedData = _this.previewCache.get(cacheKey);

                if (cachedData.type === 'image') {
                  _this.fileImg = cachedData.content;
                  _this.textContent = '';
                } else {
                  _this.fileImg = '';
                  _this.textContent = cachedData.content;
                }

                return _context.abrupt("return");

              case 9:
                _context.prev = 9;
                params = (_params = {}, _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15___default()(_params, _this.type === 'dataset' ? 'dataset_id' : 'aimodel_id', _this.dataObj.id), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15___default()(_params, "parent_dir", file.ParenDir), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15___default()(_params, "file_name", file.FileName), _params);
                console.log("this.supportImgReg.test(file.FileName)", _this.supportImgReg.test(file.FileName));

                if (!_this.supportImgReg.test(file.FileName)) {
                  _context.next = 19;
                  break;
                }

                _this.textContent = '';
                queryParams = new URLSearchParams(params);
                _this.fileImg = "/api/v1/dataset/preview?".concat(queryParams.toString()); // 缓存图片URL

                _this.previewCache.set(cacheKey, {
                  type: 'image',
                  content: _this.fileImg
                });

                _context.next = 28;
                break;

              case 19:
                if (!_this.supportTxtReg.test(file.FileName)) {
                  _context.next = 28;
                  break;
                }

                _this.fileImg = '';
                _context.next = 23;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_22__.getFilePreview)(params, _this.type);

              case 23:
                response = _context.sent;
                console.log(response);

                if ((response === null || response === void 0 ? void 0 : response.data) && (response === null || response === void 0 ? void 0 : response.data.code) === 4004) {
                  _this.showMessage = response.data.msg;
                  _this.showFlag = true;
                  _this.textContent = '';
                }

                _this.textContent = response.data; // 缓存文本内容

                _this.previewCache.set(cacheKey, {
                  type: 'text',
                  content: _this.textContent
                });

              case 28:
                _context.next = 33;
                break;

              case 30:
                _context.prev = 30;
                _context.t0 = _context["catch"](9);

                _this.$message.error(_context.t0);

              case 33:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, null, [[9, 30]]);
      }))();
    },
    closeDialog: function closeDialog() {
      this.visible = false;
      this.showFlag = false;
    },
    downloadAllModel: function downloadAllModel() {
      var _this2 = this;

      if (!this.filesList.length) return;

      if (this.dataObj.size <= maxModelFilesSize) {
        var downloadElement = document.createElement('a');
        var href = "/api/v1/".concat(this.type, "/download?").concat(this.type, "_id=").concat(this.dataObj.id);
        downloadElement.href = href;
        document.body.appendChild(downloadElement);
        downloadElement.click(); //点击下载

        document.body.removeChild(downloadElement); //下载完成移除元素
      } else {
        if (this.type === 'dataset') {
          this.sdkData.dataset_id = this.dataObj.id;
        } else {
          this.sdkData.aimodel_id = this.dataObj.id;
        }

        this.sdkData.type = this.type;
        this.sdkData.file_name = '';
        this.$nextTick(function () {
          _this2.$refs.childModelTips.handlerOpen('', maxModelFilesSize);
        });
      }
    },
    downLoadFile: function downLoadFile(file) {
      var _this3 = this;

      if (file.Size <= maxModelFilesSize) {
        var params = "".concat(this.type, "_id=").concat(this.dataObj.id, "&parent_dir=").concat(file.ParenDir, "&file_name=").concat(file.FileName);
        var downloadElement = document.createElement('a');
        var href = "/api/v1/".concat(this.type, "/file?").concat(params);
        downloadElement.href = href;
        document.body.appendChild(downloadElement);
        downloadElement.click(); //点击下载

        document.body.removeChild(downloadElement); //下载完成移除元素
      } else {
        if (this.type === 'dataset') {
          this.sdkData.dataset_id = this.dataObj.id;
        } else {
          this.sdkData.aimodel_id = this.dataObj.id;
        }

        this.sdkData.type = this.type;
        this.sdkData.file_name = '';
        this.sdkData.file_name = file.FileName;
        this.$nextTick(function () {
          _this3.$refs.childModelTips.handlerOpen('', maxModelFilesSize);
        });
      }
    },
    deleteFile: function deleteFile(file) {
      var _this4 = this;

      this.$confirm(this.$t('datasetObj.deleteDataFileConfirmTips', {
        name: file.FileName
      }), this.$t('tips'), {
        confirmButtonText: this.$t('confirm1'),
        cancelButtonText: this.$t('cancel'),
        type: 'warning',
        lockScroll: false
      }).then(function () {
        var _params2;

        _this4.loading = true;
        var params = (_params2 = {}, _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15___default()(_params2, _this4.type === 'dataset' ? 'dataset_id' : 'aimodel_id', _this4.dataObj.id), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15___default()(_params2, "parent_dir", file.ParenDir), _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_15___default()(_params2, "file_name", file.FileName), _params2);
        (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_22__.delSingleDataset)(params, _this4.type).then(function (response) {
          var res = response.data;

          if (res.code === 0) {
            _this4.params.marker = '';

            _this4.getDatasetFileList();

            _this4.$emit('editSuccess');
          } else {
            _this4.loading = false;

            _this4.$message({
              type: 'error',
              message: res.msg || _this4.$t('datasetObj.dataFileDeleteFailed')
            });
          }
        })["catch"](function (err) {
          _this4.loading = false;
          console.log(err);

          _this4.$message({
            type: 'error',
            message: err || _this4.$t('datasetObj.dataFileDeleteFailed')
          });
        });
      })["catch"](function () {});
    },
    goUploadPage: function goUploadPage() {
      this.showUploadPage = true;
    },
    getDatasetFileList: function getDatasetFileList() {
      var _this5 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().mark(function _callee2() {
        var response, res;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().wrap(function _callee2$(_context2) {
          while (1) {
            switch (_context2.prev = _context2.next) {
              case 0:
                _this5.loading = true;
                console.log(_this5.params);
                _context2.prev = 2;
                _context2.next = 5;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_22__.getFileList)(_this5.params, _this5.type);

              case 5:
                response = _context2.sent;
                console.log(response);
                res = response.data;

                if (res.code === 0) {
                  if (_this5.params.marker) {
                    console.log("marker", _this5.params.marker);
                    console.log(res.data.file_list);
                    _this5.filesList = [].concat(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_13___default()(_this5.filesList), _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_13___default()(res.data.file_list));
                    console.log(_this5.filesList);
                  } else {
                    _this5.filesList = res.data.file_list;
                  }

                  _this5.has_next = res.data.has_next;
                  _this5.params.marker = res.data.marker;

                  _this5.$nextTick(function () {
                    (0,_utils__WEBPACK_IMPORTED_MODULE_23__.initClipboard)('.tbl-file-name .clipboard-model-name');
                  });
                } else {
                  _this5.$message.error(res.msg || '加载文件列表出错');
                }

                _context2.next = 15;
                break;

              case 11:
                _context2.prev = 11;
                _context2.t0 = _context2["catch"](2);
                console.error('加载文件列表出错:', _context2.t0);

                _this5.$message.error(_context2.t0);

              case 15:
                _context2.prev = 15;
                _this5.loading = false;
                return _context2.finish(15);

              case 18:
              case "end":
                return _context2.stop();
            }
          }
        }, _callee2, null, [[2, 11, 15, 18]]);
      }))();
    },
    formatFileSize: function formatFileSize(size) {
      return (0,_utils__WEBPACK_IMPORTED_MODULE_23__.transFileSize)(size);
    },
    headerStyle: function headerStyle() {
      return {
        backgroundColor: 'rgba(247,247,247,1)',
        color: 'rgba(16, 16, 16)',
        fontSize: '14px',
        fontWeight: '400'
      };
    },
    cellStyle: function cellStyle() {
      return {
        fontSize: '14px',
        color: 'rgba(16, 16, 16)',
        fontWeight: '400'
      };
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=script&lang=js":
/*!***********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=script&lang=js ***!
  \***********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.map */ "./node_modules/core-js/modules/es.array.map.js");
/* harmony import */ var core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_map__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.promise */ "./node_modules/core-js/modules/es.promise.js");
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var _MigrateModelSync_vue__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! ./MigrateModelSync.vue */ "./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue");
/* harmony import */ var _apis_modules_common__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ~/apis/modules/common */ "./web_src/vuepages/apis/modules/common.js");
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");
/* harmony import */ var _apis_modules_modelsquare__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! ~/apis/modules/modelsquare */ "./web_src/vuepages/apis/modules/modelsquare.js");
/* harmony import */ var element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! element-ui/lib/utils/date-util */ "./node_modules/element-ui/lib/utils/date-util.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");











//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
    dataObj: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    type: {
      type: String,
      "default": 'dataset'
    }
  },
  data: function data() {
    return {
      loading: false,
      introFileName: '',
      introEmpty: false,
      canEdit: false,
      editor: null,
      editLoading: false,
      editing: false,
      editTab: 'edit',
      // edit|preview
      emptyDefaultContent: '',
      content: '',
      htmlContent: '',
      editContent: '',
      previewLoading: false,
      previewContent: '',
      submitLoading: false,
      trainUsedDataList: [],
      modelUseTaskList: [],
      repoUseTaskList: []
    };
  },
  components: {
    MigrateModelSync: _MigrateModelSync_vue__WEBPACK_IMPORTED_MODULE_11__["default"]
  },
  methods: {
    formatAccess: function formatAccess(item) {
      return item ? this.$t('modelManage.modelAccessPrivate') : this.$t('modelManage.modelAccessPublic');
    },
    formatSize: function formatSize(size) {
      return (0,_utils__WEBPACK_IMPORTED_MODULE_16__.transFileSize)(size);
    },
    formatTime: function formatTime(time) {
      return (0,element_ui_lib_utils_date_util__WEBPACK_IMPORTED_MODULE_15__.formatDate)(new Date(time * 1000), 'yyyy-MM-dd HH:mm:ss');
    },
    createIntro: function createIntro() {
      this.content = this.emptyDefaultContent;
      this.toggleEdit(true);
    },
    changeEditTab: function changeEditTab(tab) {
      var _this = this;

      console.log(tab == this.editTab);
      if (tab == this.editTab) return;
      this.editTab = tab;

      if (tab == 'preview') {
        (0,_apis_modules_common__WEBPACK_IMPORTED_MODULE_12__.getMarkdownHtml)(this.editContent).then(function (res) {
          _this.previewContent = res.data;
        })["catch"](function (err) {
          console.log(err);
        });
      } else {
        this.$nextTick(function () {
          _this.editor && _this.editor.layout();
        });
      }
    },
    toggleEdit: function toggleEdit(state) {
      var _this2 = this;

      if (state) {
        this.editing = true;
        this.editTab = 'edit';
        this.editContent = this.content;
        this.$nextTick(function () {
          _this2.initEdit();
        });
      } else {
        if (this.editContent != this.content) {
          this.$confirm(this.$t('modelManage.discardFileChanges'), this.$t('tips'), {
            confirmButtonText: this.$t('confirm1'),
            cancelButtonText: this.$t('cancel'),
            type: 'warning',
            lockScroll: false
          }).then(function () {
            _this2.editing = false;

            _this2.$nextTick(function () {
              window.initMarkdownCatalog && window.initMarkdownCatalog();
            });
          })["catch"](function () {});
        } else {
          this.editing = false;
          this.$nextTick(function () {
            window.initMarkdownCatalog && window.initMarkdownCatalog();
          });
        }
      }
    },
    resize: function resize() {
      this.editor && this.editor.layout();
    },
    initEdit: function initEdit() {
      var _this3 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_10___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8___default().mark(function _callee() {
        var monaco, model;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8___default().wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                _this3.editLoading = true;
                _context.next = 3;
                return Promise.all(/*! import() */[__webpack_require__.e("monaco-vendors-include-loader_node_modules_monaco-editor_esm_vs_editor_editor_main_js"), __webpack_require__.e("monaco")]).then(__webpack_require__.bind(__webpack_require__, /*! monaco-editor */ "include-loader!./node_modules/monaco-editor/esm/vs/editor/editor.main.js"));

              case 3:
                monaco = _context.sent;
                _this3.editor && _this3.editor.dispose();
                _this3.$refs.editContainerRefinnerHTML = '';
                _this3.editor = monaco.editor.create(_this3.$refs.editContainerRef, {
                  value: _this3.editContent,
                  language: 'markdown',
                  theme: document.documentElement.classList.contains('theme-arc-green') ? 'vs-dark' : 'vs',
                  wordWrap: 'on'
                });
                model = _this3.editor.getModel();
                model.onDidChangeContent(function () {
                  _this3.editContent = _this3.editor.getValue();
                });
                window.removeEventListener('resize', _this3.resize);
                window.addEventListener('resize', _this3.resize);
                _this3.editLoading = false;

              case 12:
              case "end":
                return _context.stop();
            }
          }
        }, _callee);
      }))();
    },
    submit: function submit() {
      var _this4 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_10___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8___default().mark(function _callee2() {
        var params, response, res, urlText;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8___default().wrap(function _callee2$(_context2) {
          while (1) {
            switch (_context2.prev = _context2.next) {
              case 0:
                if (!(!_this4.dataObj.id || _this4.submitLoading)) {
                  _context2.next = 2;
                  break;
                }

                return _context2.abrupt("return");

              case 2:
                if (!(_this4.editContent == '')) {
                  _context2.next = 5;
                  break;
                }

                _this4.$message({
                  type: 'info',
                  message: _this4.$t('modelManage.editFileContentFirst')
                });

                return _context2.abrupt("return");

              case 5:
                _this4.submitLoading = true;
                _context2.prev = 6;
                params = _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_7___default()({}, _this4.type === 'dataset' ? 'dataset_id' : 'aimodel_id', _this4.dataObj.id);
                _context2.next = 10;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_13__.postDatasetReadMe)(params, {
                  content: _this4.editContent
                }, _this4.type);

              case 10:
                response = _context2.sent;
                res = response.data;

                if (res.code == 0) {
                  urlText = _this4.type == 'dataset' ? 'datasets' : 'models';
                  window.location.href = "/".concat(urlText, "/detail/").concat(_this4.dataObj.owner_name, "/").concat(_this4.dataObj.name);
                } else {
                  _this4.$message({
                    type: 'error',
                    message: res.msg
                  });
                }

                _context2.next = 19;
                break;

              case 15:
                _context2.prev = 15;
                _context2.t0 = _context2["catch"](6);
                console.log(_context2.t0);

                _this4.$message({
                  type: 'error',
                  message: _this4.$t('submittedFailed')
                });

              case 19:
                _context2.prev = 19;
                _this4.submitLoading = false;
                return _context2.finish(19);

              case 22:
              case "end":
                return _context2.stop();
            }
          }
        }, _callee2, null, [[6, 15, 19, 22]]);
      }))();
    },
    getIntroInfo: function getIntroInfo() {
      var _this5 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_10___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8___default().mark(function _callee3() {
        var params, reponse, res, data;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8___default().wrap(function _callee3$(_context3) {
          while (1) {
            switch (_context3.prev = _context3.next) {
              case 0:
                if (_this5.dataObj.id) {
                  _context3.next = 2;
                  break;
                }

                return _context3.abrupt("return");

              case 2:
                _this5.loading = true;
                _context3.prev = 3;
                params = _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_7___default()({}, _this5.type === 'dataset' ? 'dataset_id' : 'aimodel_id', _this5.dataObj.id);
                _context3.next = 7;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_13__.getDatasetReadMe)(params, _this5.type);

              case 7:
                reponse = _context3.sent;
                res = reponse.data;
                data = res.data;

                if (res && res.code == 0) {
                  if (!data.isExistMDFile) {
                    _this5.introEmpty = true;
                    _this5.emptyDefaultContent = data.content;
                  } else {
                    _this5.introEmpty = false;
                    _this5.content = data.content;
                    _this5.htmlContent = data.htmlcontent;
                  }

                  _this5.introFileName = data.fileName;
                  window.setTimeout(function () {
                    _this5.$nextTick(function () {
                      window.initMarkdownCatalog && window.initMarkdownCatalog();
                    });
                  }, 200);
                }

                _context3.next = 16;
                break;

              case 13:
                _context3.prev = 13;
                _context3.t0 = _context3["catch"](3);
                console.log(_context3.t0);

              case 16:
                _context3.prev = 16;
                _this5.loading = false;
                return _context3.finish(16);

              case 19:
              case "end":
                return _context3.stop();
            }
          }
        }, _callee3, null, [[3, 13, 16, 19]]);
      }))();
    },
    getModelRelated: function getModelRelated() {
      var _this6 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_10___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8___default().mark(function _callee4() {
        var params, reponse, res, data;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_8___default().wrap(function _callee4$(_context4) {
          while (1) {
            switch (_context4.prev = _context4.next) {
              case 0:
                if (!(!_this6.dataObj.id || _this6.type === 'dataset')) {
                  _context4.next = 2;
                  break;
                }

                return _context4.abrupt("return");

              case 2:
                _context4.prev = 2;
                params = {
                  aimodel_id: _this6.dataObj.id
                };
                _context4.next = 6;
                return (0,_apis_modules_modelsquare__WEBPACK_IMPORTED_MODULE_14__.getModelRelatedInfo)(params);

              case 6:
                reponse = _context4.sent;
                res = reponse.data;
                console.log(res);

                if (res.data && res.code == 0) {
                  data = res.data;
                  _this6.trainUsedDataList = data.datasets.map(function (item) {
                    item.url = "/datasets/detail/{item.owner_name}/".concat(item.name);
                    item.showName = "".concat(item.owner_name, "/").concat(item.name);
                    return item;
                  });
                  _this6.repoUseTaskList = data.repos.map(function (item) {
                    item.url = "/".concat(item.owner_name, "/").concat(item.name);
                    item.showName = "".concat(item.owner_name, "/").concat(item.name);
                    return item;
                  });
                }

                _context4.next = 15;
                break;

              case 12:
                _context4.prev = 12;
                _context4.t0 = _context4["catch"](2);
                console.log(_context4.t0);

              case 15:
                _context4.prev = 15;
                return _context4.finish(15);

              case 17:
              case "end":
                return _context4.stop();
            }
          }
        }, _callee4, null, [[2, 12, 15, 17]]);
      }))();
    }
  },
  beforeMount: function beforeMount() {},
  mounted: function mounted() {
    if (this.type === 'dataset') {
      var _this$dataObj;

      this.canEdit = (_this$dataObj = this.dataObj) === null || _this$dataObj === void 0 ? void 0 : _this$dataObj.can_edit_file;
    } else {
      var _this$dataObj2, _this$dataObj2$permis;

      this.canEdit = (_this$dataObj2 = this.dataObj) === null || _this$dataObj2 === void 0 ? void 0 : (_this$dataObj2$permis = _this$dataObj2.permission) === null || _this$dataObj2$permis === void 0 ? void 0 : _this$dataObj2$permis.can_edit_file;
    }

    this.getIntroInfo();
    this.getModelRelated();
    console.log(this.dataObj);
  },
  beforeDestroy: function beforeDestroy() {
    window.removeEventListener('resize', this.resize);
    this.editor && this.editor.dispose();
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=script&lang=js":
/*!**********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************************************************************************************************************************************************/
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
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.string.trim */ "./node_modules/core-js/modules/es.string.trim.js");
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var _components_BaseDialog_vue__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ~/components/BaseDialog.vue */ "./web_src/vuepages/components/BaseDialog.vue");
/* harmony import */ var _apis_modules_modelmanage__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ~/apis/modules/modelmanage */ "./web_src/vuepages/apis/modules/modelmanage.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");













function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); if (enumerableOnly) symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; }); keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i] != null ? arguments[i] : {}; if (i % 2) { ownKeys(Object(source), true).forEach(function (key) { _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_11___default()(target, key, source[key]); }); } else if (Object.getOwnPropertyDescriptors) { Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)); } else { ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } } return target; }

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: 'MigrateModelSync',
  props: {
    data: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  components: {
    BaseDialog: _components_BaseDialog_vue__WEBPACK_IMPORTED_MODULE_12__["default"]
  },
  data: function data() {
    return {
      dlgShow: false,
      loading: true,
      submitLoading: false,
      hasUpdate: false,
      canUpdate: false,
      filesList: [],
      mainData: {},
      gated: false,
      hf_token: ''
    };
  },
  methods: {
    open: function open() {
      var _this = this;

      this.loading = true;
      this.submitLoading = false;
      this.hf_token = '';
      this.gated = false;
      this.hasUpdate = true;
      this.canUpdate = false;
      this.filesList = [];
      (0,_apis_modules_modelmanage__WEBPACK_IMPORTED_MODULE_13__.getModelMigrateUpdateInfo)({
        hf_repo_id: this.data.external_name
      }).then(function (res) {
        _this.loading = false;
        res = res.data;

        if (res.code == 200) {
          var data = res.data;
          _this.mainData = data;
          _this.hasUpdate = data.is_update;
          _this.canUpdate = _this.hasUpdate;
          _this.gated = data.gated;
          _this.filesList = (data.files || []).map(function (item) {
            return _objectSpread(_objectSpread({}, item), {}, {
              FileName: item.filename,
              SizeShow: item.update_status == 1 ? '--' : (0,_utils__WEBPACK_IMPORTED_MODULE_14__.transFileSize)(item.size),
              NewSizeShow: item.update_status == 2 ? '--' : (0,_utils__WEBPACK_IMPORTED_MODULE_14__.transFileSize)(item.size_new),
              Status: item.update_status
            });
          });
        } else {
          _this.$message.error(res.msg || _this.$t('operationFailed'));
        }
      })["catch"](function (err) {
        _this.loading = false;

        if (err.response.status == 401) {
          window.location.href = "/user/login?redirect_to=".concat(encodeURIComponent(window.location.href));
        } else {
          _this.$message.error(err.response.data.message || _this.$t('operationFailed'));
        }
      });
    },
    closed: function closed() {},
    update: function update() {
      var _this2 = this;

      if (this.submitLoading) return;

      if (this.canUpdate && this.gated && this.hf_token.trim() == '') {
        this.$message.info(this.$t('modelManage.modelLabelGatedTips'));
        return;
      }

      this.submitLoading = true;
      (0,_apis_modules_modelmanage__WEBPACK_IMPORTED_MODULE_13__.setModelMigrateUpdate)(_objectSpread(_objectSpread({}, this.mainData), {}, {
        hf_token: this.hf_token.trim()
      })).then(function (res) {
        var data = res.data;

        if (data.code == 1) {
          window.location.reload();
        } else {
          _this2.submitLoading = false;

          _this2.$message.error(data.msg || _this2.$t('operationFailed'));
        }
      })["catch"](function (err) {
        _this2.submitLoading = false;

        _this2.$message.error(err.response.data.message || _this2.$t('operationFailed'));
      });
    },
    cancel: function cancel() {
      this.dlgShow = false;
    }
  },
  beforeMount: function beforeMount() {},
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=script&lang=js":
/*!**************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=script&lang=js ***!
  \**************************************************************************************************************************************************************************************************************************/
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
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.object.define-properties */ "./node_modules/core-js/modules/es.object.define-properties.js");
/* harmony import */ var core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_properties__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.object.define-property */ "./node_modules/core-js/modules/es.object.define-property.js");
/* harmony import */ var core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_define_property__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptor */ "./node_modules/core-js/modules/es.object.get-own-property-descriptor.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptor__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.object.get-own-property-descriptors */ "./node_modules/core-js/modules/es.object.get-own-property-descriptors.js");
/* harmony import */ var core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_get_own_property_descriptors__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.object.keys */ "./node_modules/core-js/modules/es.object.keys.js");
/* harmony import */ var core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_keys__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.string.trim */ "./node_modules/core-js/modules/es.string.trim.js");
/* harmony import */ var core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_trim__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! @babel/runtime/helpers/defineProperty */ "./node_modules/@babel/runtime/helpers/defineProperty.js");
/* harmony import */ var _babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_defineProperty__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var _components_BaseTitle_vue__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! ~/components/BaseTitle.vue */ "./web_src/vuepages/components/BaseTitle.vue");
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");
/* harmony import */ var lodash_function__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! lodash/function */ "./node_modules/lodash/function.js");
/* harmony import */ var lodash_function__WEBPACK_IMPORTED_MODULE_19___default = /*#__PURE__*/__webpack_require__.n(lodash_function__WEBPACK_IMPORTED_MODULE_19__);


















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
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: 'AiforgeAccess',
  props: {
    id: {
      type: String,
      "default": ''
    },
    title: {
      type: String,
      "default": ''
    },
    btnText: {
      type: String,
      "default": ''
    },
    placeholder: {
      type: String,
      "default": ''
    },
    type: {
      type: String,
      "default": ''
    },
    org: {
      type: String,
      "default": ''
    },
    filesType: {
      type: String,
      "default": 'dataset'
    }
  },
  data: function data() {
    return {
      loading: false,
      collaborator: '',
      searching: false,
      showResults: false,
      searchResults: [],
      showEmpty: false,
      usersList: [],
      canChangeDatasetTeams: false,
      showMessage: false,
      showMessageText: '',
      successFlag: false,
      collaboratorUser: ''
    };
  },
  components: {
    BaseTitle: _components_BaseTitle_vue__WEBPACK_IMPORTED_MODULE_17__["default"]
  },
  mounted: function mounted() {
    var _this2 = this;

    this.$nextTick(function () {
      if (_this2.$refs.myInput && document.activeElement === _this2.$refs.myInput) {
        _this2.$refs.myInput.blur();
      }
    });
    this.getCollaborateList();
  },
  computed: {},
  // filters: {
  //   permissionFilter(value) {
  //     const permissionMap = {
  //       1: "可读权限",
  //       2: "可写权限",
  //       3: "管理员",
  //       4: "所有者",
  //     };
  //     return permissionMap[value] || "未知权限";
  //   },
  // },
  methods: {
    handlSearch: (0,lodash_function__WEBPACK_IMPORTED_MODULE_19__.debounce)(function () {
      this.searchUsers();
    }, 500),
    searchUsers: function searchUsers() {
      var _this3 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_16___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().mark(function _callee() {
        var fetchMethod, queryArgs, response, res;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                if (_this3.collaborator.trim()) {
                  _context.next = 6;
                  break;
                }

                _this3.searchResults = [];
                _this3.showResults = false;
                _this3.showEmpty = false;
                _this3.searching = false;
                return _context.abrupt("return");

              case 6:
                _this3.searching = true;
                _context.prev = 7;
                // 调用远程搜索API
                fetchMethod = _this3.type === 'Teams' ? _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_18__.getTeams : _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_18__.getUsers;
                queryArgs = _this3.type === 'Teams' ? {
                  q: _this3.collaborator,
                  org: _this3.org
                } : {
                  q: _this3.collaborator
                };
                _context.next = 12;
                return fetchMethod(queryArgs);

              case 12:
                response = _context.sent;
                console.log(response);
                res = response.data;

                if (res.ok) {
                  if (res.data.length > 2) {}

                  _this3.searchResults = res.data;
                  _this3.showResults = !!res.data.length;
                  _this3.showEmpty = !res.data.length;
                } else {
                  _this3.$message.error('搜索失败，请稍后重试');
                } // this.searchResults = response.data;


                _context.next = 23;
                break;

              case 18:
                _context.prev = 18;
                _context.t0 = _context["catch"](7);
                console.error('搜索用户失败:', _context.t0);

                _this3.$message.error('搜索失败，请稍后重试');

                _this3.searchResults = [];

              case 23:
                _context.prev = 23;
                _this3.searching = false;
                return _context.finish(23);

              case 26:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, null, [[7, 18, 23, 26]]);
      }))();
    },
    // 选择用户
    selectUser: function selectUser(user) {
      var _this4 = this;

      if (this.type === 'Teams') {
        this.collaborator = user.name;
      } else {
        this.collaborator = user.full_name ? "".concat(user.username, " (").concat(user.full_name, ")") : user.username;
        this.collaboratorUser = user.username;
      }

      this.showResults = false; // 使用 $nextTick 确保 DOM 更新后聚焦输入框

      this.$nextTick(function () {
        _this4.$refs.myInput.focus();
      });
    },
    getCollaborateList: function getCollaborateList() {
      var _this5 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_16___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().mark(function _callee2() {
        var data, response, res, canOper;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().wrap(function _callee2$(_context2) {
          while (1) {
            switch (_context2.prev = _context2.next) {
              case 0:
                _context2.prev = 0;
                // 调用远程搜索API
                data = {
                  subject_id: _this5.id,
                  subject_type: _this5.filesType === 'dataset' ? '1' : '2'
                };
                _this5.loading = true;
                _context2.next = 5;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_18__.getCollaborate)(data);

              case 5:
                response = _context2.sent;
                console.log(response);
                res = response.data;

                if (res.code === 0) {
                  _this5.usersList = res.data;

                  _this5.usersList.Teams.forEach(function (item) {
                    item.canDelete = _this5.filesType === 'dataset' ? item.AllowedToChangeDatasetTeams : item.AllowedToChangeAimodelTeams;
                  });

                  console.log(_this5.usersList);
                  canOper = _this5.filesType === 'dataset' ? _this5.usersList.CanChangeDatasetTeams : _this5.usersList.CanChangeAimodelTeams;
                  _this5.canChangeDatasetTeams = canOper || _this5.type === 'Collaborators';
                } else {
                  _this5.$message.error(res.msg || '获取失败，请稍后重试');
                }

                _context2.next = 15;
                break;

              case 11:
                _context2.prev = 11;
                _context2.t0 = _context2["catch"](0);
                console.error('搜索用户失败:', _context2.t0);

                _this5.$message.error(_context2.t0);

              case 15:
                _context2.prev = 15;
                _this5.loading = false;
                return _context2.finish(15);

              case 18:
              case "end":
                return _context2.stop();
            }
          }
        }, _callee2, null, [[0, 11, 15, 18]]);
      }))();
    },
    sumbitCollaborate: function sumbitCollaborate() {
      var _arguments = arguments,
          _this6 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_16___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().mark(function _callee3() {
        var flag, fetchMethod, queryArgs, response, res;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().wrap(function _callee3$(_context3) {
          while (1) {
            switch (_context3.prev = _context3.next) {
              case 0:
                flag = _arguments.length > 0 && _arguments[0] !== undefined ? _arguments[0] : false;

                if (flag) {
                  _this6.collaboratorUser = _this6.collaborator;
                }

                _this6.$refs.myForm.reportValidity();

                if (_this6.collaborator) {
                  _context3.next = 5;
                  break;
                }

                return _context3.abrupt("return");

              case 5:
                _context3.prev = 5;
                fetchMethod = _this6.type === 'Teams' ? _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_18__.postTeam : _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_18__.postCollaborate;
                queryArgs = _objectSpread({
                  subject_id: _this6.id,
                  subject_type: _this6.filesType === 'dataset' ? '1' : '2'
                }, _this6.type === 'Teams' ? {
                  team: _this6.collaborator
                } : {
                  collaborator: _this6.collaboratorUser
                });

                if (!(_this6.type === 'Teams')) {
                  _context3.next = 14;
                  break;
                }

                _context3.next = 11;
                return fetchMethod(queryArgs, _this6.filesType);

              case 11:
                _context3.t0 = _context3.sent;
                _context3.next = 17;
                break;

              case 14:
                _context3.next = 16;
                return fetchMethod(queryArgs);

              case 16:
                _context3.t0 = _context3.sent;

              case 17:
                response = _context3.t0;
                console.log(response);
                res = response.data;

                if (!(res.code === 0)) {
                  _context3.next = 28;
                  break;
                }

                _this6.collaborator = '';
                _context3.next = 24;
                return _this6.getCollaborateList();

              case 24:
                _this6.showMessageText = _this6.type === 'Teams' ? _this6.filesType === 'dataset' ? _this6.$t('datasetObj.add_team_success') : _this6.$t('modelManage.add_team_success') : _this6.$t('datasetObj.add_collaborator_success');
                _this6.successFlag = true;
                _context3.next = 30;
                break;

              case 28:
                _this6.showMessageText = res.msg;
                _this6.successFlag = false;

              case 30:
                _context3.next = 36;
                break;

              case 32:
                _context3.prev = 32;
                _context3.t1 = _context3["catch"](5);
                console.error('添加用户失败:', _context3.t1);

                _this6.$message.error(_context3.t1);

              case 36:
                _context3.prev = 36;
                _this6.collaborator = '';
                _this6.showResults = false;
                _this6.showEmpty = false;
                _this6.showMessage = true;
                setTimeout(function () {
                  _this6.showMessage = false;
                }, 2000);
                return _context3.finish(36);

              case 43:
              case "end":
                return _context3.stop();
            }
          }
        }, _callee3, null, [[5, 32, 36, 43]]);
      }))();
    },
    permissionFilter: function permissionFilter(value) {
      var permissionMap = {
        1: this.$t('datasetObj.readPermission'),
        2: this.$t('datasetObj.writePermission'),
        3: this.$t('datasetObj.administrators'),
        4: this.$t('datasetObj.dataset_owner')
      };
      return permissionMap[value] || "未知权限";
    },
    deleteUser: function deleteUser(id) {
      var _this7 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_16___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().mark(function _callee4() {
        var fetchMethod, queryArgs, response, res;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().wrap(function _callee4$(_context4) {
          while (1) {
            switch (_context4.prev = _context4.next) {
              case 0:
                _context4.prev = 0;
                _this7.loading = true;
                fetchMethod = _this7.type === 'Teams' ? _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_18__.deleteTeam : _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_18__.deleteCollaborate;
                queryArgs = _objectSpread({
                  subject_id: _this7.id,
                  subject_type: _this7.filesType === 'dataset' ? '1' : '2'
                }, _this7.type === 'Teams' ? {
                  id: id
                } : {
                  uid: id
                }); // 根据方法类型调整参数

                if (!(_this7.type === 'Teams')) {
                  _context4.next = 10;
                  break;
                }

                _context4.next = 7;
                return fetchMethod(queryArgs, _this7.filesType);

              case 7:
                _context4.t0 = _context4.sent;
                _context4.next = 13;
                break;

              case 10:
                _context4.next = 12;
                return fetchMethod(queryArgs);

              case 12:
                _context4.t0 = _context4.sent;

              case 13:
                response = _context4.t0;
                console.log(response);
                res = response.data;

                if (!(res.code === 0)) {
                  _context4.next = 24;
                  break;
                }

                _this7.collaborator = '';
                _context4.next = 20;
                return _this7.getCollaborateList();

              case 20:
                _this7.showMessageText = _this7.type === 'Teams' ? _this7.filesType === 'dataset' ? _this7.$t('datasetObj.remove_team_success') : _this7.$t('modelManage.remove_team_success') : _this7.$t('datasetObj.remove_collaborator_success');
                _this7.successFlag = true;
                _context4.next = 26;
                break;

              case 24:
                _this7.showMessageText = res.msg;
                _this7.successFlag = false;

              case 26:
                _context4.next = 32;
                break;

              case 28:
                _context4.prev = 28;
                _context4.t1 = _context4["catch"](0);
                console.error('添加用户失败:', _context4.t1);

                _this7.$message.error(_context4.t1);

              case 32:
                _context4.prev = 32;
                _this7.loading = false;
                _this7.showMessage = true;
                setTimeout(function () {
                  _this7.showMessage = false;
                }, 2000);
                return _context4.finish(32);

              case 37:
              case "end":
                return _context4.stop();
            }
          }
        }, _callee4, null, [[0, 28, 32, 37]]);
      }))();
    },
    deleteImage: function deleteImage(id) {
      var _this = this;

      var ele = "#".concat(_this.type, "-").concat(_this.id);
      $(ele).modal({
        onDeny: function onDeny() {},
        onApprove: function onApprove() {
          _this.deleteUser(id);
        }
      }).modal("show");
    },
    handleCommand: function handleCommand(command) {
      var _this8 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_16___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().mark(function _callee5() {
        var data, response, res;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_14___default().wrap(function _callee5$(_context5) {
          while (1) {
            switch (_context5.prev = _context5.next) {
              case 0:
                _context5.prev = 0;
                data = {
                  uid: command.id,
                  mode: command.mode,
                  subject_id: _this8.id,
                  subject_type: _this8.filesType === 'dataset' ? '1' : '2'
                };
                _context5.next = 4;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_18__.modifyCollaborateAcess)(data);

              case 4:
                response = _context5.sent;
                console.log(response);
                res = response.data;

                if (res.code === 0) {
                  _this8.$message.success('修改成功！');

                  _this8.getCollaborateList();
                } else {
                  _this8.$message.error(res.msg || '修改失败，请稍后重试');
                }

                _context5.next = 14;
                break;

              case 10:
                _context5.prev = 10;
                _context5.t0 = _context5["catch"](0);
                console.error('修改失败，请稍后重试:', _context5.t0);

                _this8.$message.error(_context5.t0);

              case 14:
                _context5.prev = 14;
                return _context5.finish(14);

              case 16:
              case "end":
                return _context5.stop();
            }
          }
        }, _callee5, null, [[0, 10, 14, 16]]);
      }))();
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=script&lang=js":
/*!***************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=script&lang=js ***!
  \***************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @babel/runtime/helpers/toConsumableArray */ "./node_modules/@babel/runtime/helpers/toConsumableArray.js");
/* harmony import */ var _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var _components_BaseTitle_vue__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ~/components/BaseTitle.vue */ "./web_src/vuepages/components/BaseTitle.vue");
/* harmony import */ var _Access_vue__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ./Access.vue */ "./web_src/vuepages/components/square/detail/setting/Access.vue");
/* harmony import */ var _pages_dataset_square_constant_js__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ~/pages/dataset/square/constant.js */ "./web_src/vuepages/pages/dataset/square/constant.js");
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");






//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
  name: 'AiforgeSetting',
  props: {
    dataObj: {
      type: Object,
      "default": function _default() {
        return {};
      }
    }
  },
  data: function data() {
    var _this2 = this;

    return {
      form: {
        name: this.dataObj.name || "",
        alias: this.dataObj.alias || "",
        is_private: this.dataObj.is_private || false,
        tags: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_5___default()(this.dataObj.tags || []),
        tasks: _babel_runtime_helpers_toConsumableArray__WEBPACK_IMPORTED_MODULE_5___default()(this.dataObj.tasks || []),
        licenses: this.dataObj.licenses || ""
      },
      baseLoading: false,
      delLoading: false,
      visible: false,
      name: '',
      owners: [{
        name: this.dataObj.owner_name
      }],
      Category: _pages_dataset_square_constant_js__WEBPACK_IMPORTED_MODULE_8__.Category,
      Task: _pages_dataset_square_constant_js__WEBPACK_IMPORTED_MODULE_8__.Task,
      License: _pages_dataset_square_constant_js__WEBPACK_IMPORTED_MODULE_8__.License,
      rules: {
        name: [{
          required: true,
          message: this.$t('datasetObj.dataset_name_required3'),
          trigger: "blur"
        }, {
          validator: function validator(rule, value, callback) {
            if (/^[a-zA-Z0-9][a-zA-Z0-9-_.]{0,98}[a-zA-Z0-9]$/.test(value) == false) {
              callback(new Error(_this2.$t('datasetObj.dataset_name_required2')));
            } else {
              callback();
            }
          },
          trigger: "blur"
        }],
        alias: [// { required: true, message: this.$t('datasetObj.dataset_name_required3'), trigger: "blur" },
        {
          validator: function validator(rule, value, callback) {
            if (/^[\u4e00-\u9fa5a-zA-Z0-9-_.]{0,100}$/.test(value) == false) {
              callback(new Error(_this2.$t('datasetObj.dataset_name_required2')));
            } else {
              callback();
            }
          },
          trigger: "blur"
        }],
        tags: [{
          required: true,
          message: this.$t('datasetObj.category_required'),
          trigger: "change"
        }],
        tasks: [{
          required: true,
          message: this.$t('datasetObj.select_task_required'),
          trigger: "change"
        }]
      }
    };
  },
  components: {
    BaseTitle: _components_BaseTitle_vue__WEBPACK_IMPORTED_MODULE_6__["default"],
    Access: _Access_vue__WEBPACK_IMPORTED_MODULE_7__["default"]
  },
  mounted: function mounted() {
    console.log(this.dataObj.id);
  },
  computed: {},
  methods: {
    editBaseInfo: function editBaseInfo() {
      var _this3 = this;

      this.$refs['form'].validate(function (valid) {
        if (valid) {
          _this3.baseLoading = true;

          _this3.editDataset();
        } else {
          return false;
        }
      });
    },
    editDataset: function editDataset() {
      var _this4 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default().mark(function _callee() {
        var response;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default().wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                _context.prev = 0;
                _context.next = 3;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_9__.editDatasetsDetail)({
                  dataset_id: _this4.dataObj.id
                }, _this4.form, 'dataset');

              case 3:
                response = _context.sent;
                console.log(response);
                _this4.baseLoading = false;

                if (response.data.code === 0) {
                  _this4.$message.success("更新成功" || 0);

                  if (_this4.dataObj.name !== _this4.form.name) {
                    location.href = "/datasets/detail/".concat(_this4.dataObj.owner_name, "/").concat(_this4.form.name);
                  } else {
                    _this4.$emit('editSuccess');
                  }
                } else {
                  _this4.$message.error(response.data.msg);
                }

                _context.next = 13;
                break;

              case 9:
                _context.prev = 9;
                _context.t0 = _context["catch"](0);
                _this4.baseLoading = false;

                _this4.$message.error(_context.t0);

              case 13:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, null, [[0, 9]]);
      }))();
    },
    showModal: function showModal() {
      var _this = this;

      var ele = "#dataset-del";
      $(ele).modal({
        onDeny: function onDeny() {
          console.log('Denied');
        },
        onApprove: function onApprove() {
          console.log('Approved'); // _this.deleteUser(id)
        }
      }).modal("show");
    },
    deleteDataset: function deleteDataset() {
      var _this5 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_4___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default().mark(function _callee2() {
        var response;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_2___default().wrap(function _callee2$(_context2) {
          while (1) {
            switch (_context2.prev = _context2.next) {
              case 0:
                if (!(_this5.name !== _this5.dataObj.alias)) {
                  _context2.next = 3;
                  break;
                }

                _this5.$message.error(_this5.$t('datasetObj.deleteDatasetNameError'));

                return _context2.abrupt("return");

              case 3:
                _context2.prev = 3;
                _this5.delLoading = true;
                _context2.next = 7;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_9__.delDataset)({
                  dataset_id: _this5.dataObj.id
                });

              case 7:
                response = _context2.sent;
                console.log(response);
                _this5.delLoading = false;

                if (response.data.code === 0) {
                  _this5.$message.success(_this5.$t('imagesObj.deleteSuccessTips'));

                  window.location.href = '/explore/datasets';
                } else {
                  _this5.$message.error(response.data.msg);
                }

                _context2.next = 17;
                break;

              case 13:
                _context2.prev = 13;
                _context2.t0 = _context2["catch"](3);
                _this5.delLoading = false;

                _this5.$message.error(_context2.t0);

              case 17:
              case "end":
                return _context2.stop();
            }
          }
        }, _callee2, null, [[3, 13]]);
      }))();
    }
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=script&lang=js":
/*!****************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=script&lang=js ***!
  \****************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.string.replace */ "./node_modules/core-js/modules/es.string.replace.js");
/* harmony import */ var core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_replace__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var _components_BaseTitle_vue__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ~/components/BaseTitle.vue */ "./web_src/vuepages/components/BaseTitle.vue");
/* harmony import */ var _apis_modules_common__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ~/apis/modules/common */ "./web_src/vuepages/apis/modules/common.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! spark-md5 */ "./node_modules/spark-md5/spark-md5.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(spark_md5__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var highlight_js__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! highlight.js */ "./node_modules/highlight.js/lib/index.js");
/* harmony import */ var highlight_js__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(highlight_js__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");








//
//
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
  name: 'UsageIntro',
  components: {
    BaseTitle: _components_BaseTitle_vue__WEBPACK_IMPORTED_MODULE_8__["default"]
  },
  props: {
    dataObj: {
      type: Object,
      "default": function _default() {
        return {};
      }
    },
    type: {
      type: String,
      "default": 'dataset'
    }
  },
  data: function data() {
    return {
      loading: true,
      content: '',
      promotePath: "tips/".concat(this.type, "/sdkcode").concat(_langs__WEBPACK_IMPORTED_MODULE_13__.lang == 'zh-CN' ? '' : '_en', ".md")
    };
  },
  methods: {
    sparkMD5Hash: function sparkMD5Hash() {
      var str = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : '';
      return spark_md5__WEBPACK_IMPORTED_MODULE_10___default().hash(str) + Math.random().toString().replace('0.', '');
    },
    hljsAndInsertCopyButton: function hljsAndInsertCopyButton(htmlStr) {
      var html = document.createElement('div');
      html.innerHTML = htmlStr;
      var codeBlocks = html.querySelectorAll('.code-block');

      for (var i = 0, iLen = codeBlocks.length; i < iLen; i++) {
        var codeBlockI = codeBlocks[i];
        var txt = codeBlockI.textContent.slice(0, -1);
        var codeEl = codeBlockI.querySelector('code');
        var code = codeEl.textContent;
        codeEl.innerHTML = highlight_js__WEBPACK_IMPORTED_MODULE_11___default().highlight('python', code).value;
        var copyBtn = document.createElement('div');
        copyBtn.classList = ['copy-btn'];

        if (this.type !== 'modelFile') {
          copyBtn.innerHTML = "<span style=\"color: #0366d6;cursor: pointer;\" class=\"ui poping inline up clipboard\" id=\"clipboard-".concat(this.sparkMD5Hash(txt), "\"\n            data-position=\"top center\" data-variation=\"inverted tiny\" data-success=\"").concat(this.$t('copySuccess'), "\"\n            data-content=\"").concat(this.$t('copy'), "\" data-original=\"").concat(this.$t('copy'), "\" \n            data-clipboard-text=\"\"><i style=\"font-size:14px;\" class=\"copy outline icon\"></i></span>");
          copyBtn.querySelector('span').setAttribute('data-clipboard-text', txt);
        } else {
          copyBtn.innerHTML = "<span style=\"color: #0366d6;cursor: pointer;\" class=\"ui poping inline up clipboard\" id=\"clipboard-".concat(this.sparkMD5Hash(txt), "\"\n            data-position=\"top center\" data-variation=\"inverted tiny\" data-success=\"").concat(this.i18n['cloudeBrainMirror']['copy_succeeded'], "\"\n            data-content=\"").concat(this.i18n['cloudeBrainMirror']['copy'], "\" data-original=\"").concat(this.i18n['cloudeBrainMirror']['copy'], "\" \n            data-clipboard-text=\"\"><i style=\"font-size:14px;\" class=\"copy outline icon\"></i></span>");
          copyBtn.querySelector('span').setAttribute('data-clipboard-text', txt);
        }

        codeBlockI.outerHTML = "<div class=\"code-content\">".concat(codeBlockI.outerHTML).concat(copyBtn.outerHTML, "</div>");
      }

      return html.innerHTML;
    },
    getMarkdown: function getMarkdown(str) {
      var _this = this;

      (0,_apis_modules_common__WEBPACK_IMPORTED_MODULE_9__.getMarkdownHtml)(str).then(function (res) {
        _this.loading = false;
        var html = res.data;
        _this.content = _this.hljsAndInsertCopyButton(html);

        _this.$nextTick(function () {
          (0,_utils__WEBPACK_IMPORTED_MODULE_12__.initClipboard)('.base-dlg .clipboard');
        });
      })["catch"](function (err) {
        _this.loading = false;
        console.log(err);
      });
    },
    getContent: function getContent() {
      var _this2 = this;

      if (!this.promotePath) return;
      this.loading = true;
      (0,_apis_modules_common__WEBPACK_IMPORTED_MODULE_9__.getPromoteData)(this.promotePath).then(function (res) {
        var contentStr = res.data;

        if (contentStr) {
          var params = _this2.type == 'dataset' ? {
            dataset_name: [_this2.dataObj.name]
          } : {
            pretrain_model_name: [_this2.dataObj.name]
          };
          (0,_apis_modules_common__WEBPACK_IMPORTED_MODULE_9__.getSDKCode)(params).then(function (res) {
            res = res.data;

            if (res.code == 0 && res.data && res.data.code) {
              contentStr += "\n```python\n".concat(res.data.code, "```\n");
            }

            _this2.getMarkdown(contentStr);
          })["catch"](function (err) {
            console.log(err);
          });
        }
      })["catch"](function (err) {
        _this2.loading = false;
        console.log(err);
      });
    }
  },
  mounted: function mounted() {
    this.getContent();
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=script&lang=js":
/*!*************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=script&lang=js ***!
  \*************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var _components_NotFound_vue__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ~/components/NotFound.vue */ "./web_src/vuepages/components/NotFound.vue");
/* harmony import */ var _components_square_detail_Header_vue__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! ~/components/square/detail/Header.vue */ "./web_src/vuepages/components/square/detail/Header.vue");
/* harmony import */ var _components_square_detail_intro_Intro_vue__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! ~/components/square/detail/intro/Intro.vue */ "./web_src/vuepages/components/square/detail/intro/Intro.vue");
/* harmony import */ var _components_square_detail_files_FileList_vue__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! ~/components/square/detail/files/FileList.vue */ "./web_src/vuepages/components/square/detail/files/FileList.vue");
/* harmony import */ var _components_square_detail_usage_UsageIntro_vue__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! ~/components/square/detail/usage/UsageIntro.vue */ "./web_src/vuepages/components/square/detail/usage/UsageIntro.vue");
/* harmony import */ var _components_square_detail_setting_Setting_vue__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! ~/components/square/detail/setting/Setting.vue */ "./web_src/vuepages/components/square/detail/setting/Setting.vue");
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");






//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
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
      emptyPage: false,
      loading: false,
      // tabList:[
      //   {name: this.$t('datasetObj.dataset_intro'), icon: 'ri-database-2-line',key: 'intro'},
      //   {name: this.$t('modelManage.datasetfile'), icon: 'ri-list-check', key: 'files'},
      //   {name: this.$t('datasetObj.use_dataset'), icon: 'ri-ticket-line',  key: 'usage'},
      // ],
      tab: 'intro',
      dataset_name: '',
      datasetObj: {},
      windowWidth: window.innerWidth
    };
  },
  components: {
    Header: _components_square_detail_Header_vue__WEBPACK_IMPORTED_MODULE_7__["default"],
    NotFound: _components_NotFound_vue__WEBPACK_IMPORTED_MODULE_6__["default"],
    Intro: _components_square_detail_intro_Intro_vue__WEBPACK_IMPORTED_MODULE_8__["default"],
    FileList: _components_square_detail_files_FileList_vue__WEBPACK_IMPORTED_MODULE_9__["default"],
    UsageIntro: _components_square_detail_usage_UsageIntro_vue__WEBPACK_IMPORTED_MODULE_10__["default"],
    Setting: _components_square_detail_setting_Setting_vue__WEBPACK_IMPORTED_MODULE_11__["default"]
  },
  computed: {
    isMobile: function isMobile() {
      return this.windowWidth <= 768;
    },
    tabList: function tabList() {
      return this.isMobile ? [{
        name: this.$t('datasetObj.dataset_intro'),
        icon: 'ri-database-2-line',
        key: 'intro'
      }, {
        name: this.$t('modelManage.fileShort'),
        icon: 'ri-list-check',
        key: 'files'
      }, {
        name: this.$t('cloudbrainObj.useImage'),
        icon: 'ri-ticket-line',
        key: 'usage'
      }] : [{
        name: this.$t('datasetObj.dataset_intro'),
        icon: 'ri-database-2-line',
        key: 'intro'
      }, {
        name: this.$t('modelManage.datasetfile'),
        icon: 'ri-list-check',
        key: 'files'
      }, {
        name: this.$t('datasetObj.use_dataset'),
        icon: 'ri-ticket-line',
        key: 'usage'
      }];
    }
  },
  methods: {
    changeTab: function changeTab(item) {
      this.tab = item.key;
    },
    getDatasetDetail: function getDatasetDetail() {
      var _this = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_5___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3___default().mark(function _callee() {
        var response, res;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_3___default().wrap(function _callee$(_context) {
          while (1) {
            switch (_context.prev = _context.next) {
              case 0:
                _context.prev = 0;
                _this.loading = true;
                _context.next = 4;
                return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_12__.getDatasetsDetail)({
                  dataset_name: _this.dataset_name
                });

              case 4:
                response = _context.sent;
                res = response.data;

                if (res.code === 0) {
                  _this.datasetObj = res.data;
                } else {
                  if (res.code === 9004) {
                    _this.emptyPage = true;
                  } else {
                    _this.$message.error(res.msg);
                  }
                }

                _context.next = 12;
                break;

              case 9:
                _context.prev = 9;
                _context.t0 = _context["catch"](0);

                _this.$message.error(_context.t0);

              case 12:
                _context.prev = 12;
                _this.loading = false;
                return _context.finish(12);

              case 15:
              case "end":
                return _context.stop();
            }
          }
        }, _callee, null, [[0, 9, 12, 15]]);
      }))();
    },
    editSuccess: function editSuccess() {
      this.getDatasetDetail();
    },
    handleResize: function handleResize() {
      this.windowWidth = window.innerWidth;
      console.log("isMobile", this.isMobile);
    }
  },
  beforeMount: function beforeMount() {
    var currentUrl = window.location.pathname; // 提取路径部分

    var pathParts = currentUrl.split('/'); // 获取最后一个路径段

    var lastPathSegment = pathParts[pathParts.length - 1]; // "create"
    // 获取倒数第二个路径段（即父级目录）

    var parentSegment = pathParts[pathParts.length - 2]; // "notebook"

    this.dataset_name = "".concat(parentSegment, "/").concat(lastPathSegment);
    (0,_utils__WEBPACK_IMPORTED_MODULE_13__.setWebpackPublicPath)();
    window.addEventListener('resize', this.handleResize);
  },
  mounted: function mounted() {
    var urlParams = (0,_utils__WEBPACK_IMPORTED_MODULE_13__.getUrlSearchParams)();

    if (urlParams.tab) {
      this.tab = urlParams.tab;
    }

    this.getDatasetDetail();
  },
  beforeDestroy: function beforeDestroy() {
    window.removeEventListener('resize', this.handleResize);
  }
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=script&lang=js":
/*!********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=script&lang=js ***!
  \********************************************************************************************************************************************************************************************************************/
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
/* harmony default export */ __webpack_exports__["default"] = ({
  name: "CreateForm",
  props: {
    title: {
      type: Array,
      "default": function _default() {
        return [];
      }
    },
    active: {
      type: Number,
      "default": 0
    },
    showSteps: {
      type: Boolean,
      "default": true
    }
  },
  components: {},
  data: function data() {
    return {
      formactive: 0,
      formTtitle: this.title[0]
    };
  },
  watch: {
    active: function active(val) {
      this.formactive = val;
      this.formTtitle = this.title[val];
    }
  },
  methods: {},
  mounted: function mounted() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=script&lang=js":
/*!********************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=script&lang=js ***!
  \********************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.index-of */ "./node_modules/core-js/modules/es.array.index-of.js");
/* harmony import */ var core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_index_of__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.reduce */ "./node_modules/core-js/modules/es.array.reduce.js");
/* harmony import */ var core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.number.to-fixed */ "./node_modules/core-js/modules/es.number.to-fixed.js");
/* harmony import */ var core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_to_fixed__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_number_to_precision__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.number.to-precision */ "./node_modules/core-js/modules/es.number.to-precision.js");
/* harmony import */ var core_js_modules_es_number_to_precision__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_to_precision__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.promise */ "./node_modules/core-js/modules/es.promise.js");
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! @babel/runtime/regenerator */ "./node_modules/@babel/runtime/regenerator/index.js");
/* harmony import */ var _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! regenerator-runtime/runtime */ "./node_modules/regenerator-runtime/runtime.js");
/* harmony import */ var regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(regenerator_runtime_runtime__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! @babel/runtime/helpers/asyncToGenerator */ "./node_modules/@babel/runtime/helpers/asyncToGenerator.js");
/* harmony import */ var _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(_babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var _components_NotFound_vue__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! ~/components/NotFound.vue */ "./web_src/vuepages/components/NotFound.vue");
/* harmony import */ var dropzone_dist_dropzone_css__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! dropzone/dist/dropzone.css */ "./node_modules/dropzone/dist/dropzone.css");
/* harmony import */ var dropzone__WEBPACK_IMPORTED_MODULE_20__ = __webpack_require__(/*! dropzone */ "./node_modules/dropzone/dist/dropzone.js");
/* harmony import */ var dropzone__WEBPACK_IMPORTED_MODULE_20___default = /*#__PURE__*/__webpack_require__.n(dropzone__WEBPACK_IMPORTED_MODULE_20__);
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_21__ = __webpack_require__(/*! spark-md5 */ "./node_modules/spark-md5/spark-md5.js");
/* harmony import */ var spark_md5__WEBPACK_IMPORTED_MODULE_21___default = /*#__PURE__*/__webpack_require__.n(spark_md5__WEBPACK_IMPORTED_MODULE_21__);
/* harmony import */ var _apis_modules_dataset__WEBPACK_IMPORTED_MODULE_22__ = __webpack_require__(/*! ~/apis/modules/dataset */ "./web_src/vuepages/apis/modules/dataset.js");
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_23__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");
/* harmony import */ var _apis_modules_storage__WEBPACK_IMPORTED_MODULE_24__ = __webpack_require__(/*! ~/apis/modules/storage */ "./web_src/vuepages/apis/modules/storage.js");
/* harmony import */ var _pages_modelmanage_components_FolderUploadSelect_vue__WEBPACK_IMPORTED_MODULE_25__ = __webpack_require__(/*! ~/pages/modelmanage/components/FolderUploadSelect.vue */ "./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue");


















//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//








(dropzone__WEBPACK_IMPORTED_MODULE_20___default().autoDiscover) = false;
var uploadChunkSize = 1024 * 1024 * 64;
var md5ChunkSize = 1024 * 1024 * 64;
var maxFileCount = 100;
var maxModelFilesSize = window.MAX_MODEL_SIZE || 536870912; // 200 GB

var UNITS = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB'];
/* harmony default export */ __webpack_exports__["default"] = ({
  props: {
    ownerName: {
      type: String,
      "default": ''
    },
    modelId: {
      type: String,
      "default": ''
    },
    uploadDir: {
      type: String,
      "default": ''
    },
    subjectType: {
      type: String,
      "default": '1'
    }
  },
  data: function data() {
    return {
      emptyPage: false,
      modelData: {},
      uploadType: 'file',
      // file|folder
      dropzoneHandler: null,
      type: '1',
      // 1-修改，其它-新增
      state: {
        type: '',
        id: '',
        name: '',
        version: '',
        engine: '',
        label: '',
        description: '',
        size: 0
      },
      originData: null,
      showUploadErr: false,
      uploadErrTxt: '',
      uploadFiles: [],
      uploadLength: 0,
      uploadSuccessLength: 0,
      uploadStatusList: [],
      maxModelFilesSize: maxModelFilesSize,
      uploading: false,
      btnFlag: false,
      remaining_storage: 0,
      totalSize: 0,
      isStorageExceeded: false,
      storageShow: false,
      showOwenerTips: false,
      canShowError: false
    };
  },
  components: {
    NotFound: _components_NotFound_vue__WEBPACK_IMPORTED_MODULE_18__["default"],
    FolderUploadSelect: _pages_modelmanage_components_FolderUploadSelect_vue__WEBPACK_IMPORTED_MODULE_25__["default"]
  },
  computed: {
    // 格式化的剩余空间显示（自动单位转换）
    formattedRemaining: function formattedRemaining() {
      return this.formatBytes(Math.max(this.remaining_storage, 0));
    },
    formattedTotal: function formattedTotal() {
      return this.formatBytes(this.totalSize);
    },
    formatUploadir: function formatUploadir() {
      if (this.uploadDir) {
        return this.uploadDir.slice(1) + '/';
      } else {
        return '';
      }
    }
  },
  methods: {
    initDropZone: function initDropZone() {
      var _this2 = this;

      var previewTemplate = "\n                <div class=\"dz-preview dz-file-preview\"> \n                <div class=\"dz-image\"> \n                    <img data-dz-thumbnail /> \n                </div> \n                <div class=\"dz-details\"> \n                    <div class=\"dz-size\"><span data-dz-size></span></div> \n                    <div class=\"dz-filename\"><span data-dz-name></span></div> \n                </div> \n\n                <div style=\"opacity:0\" class=\"dz-progress\"><span class=\"dz-upload\" data-dz-uploadprogress></span></div> \n                <div class=\"dz-error-message\" style=\"line-height: 1.5;\"><span data-dz-errormessage></span></div> \n                <div class=\"dz-success-mark\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" width=\"54\" height=\"54\"><path fill=\"none\" d=\"M0 0h24v24H0z\"/><path d=\"M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10-4.477 10-10 10zm0-2a8 8 0 1 0 0-16 8 8 0 0 0 0 16zm-.997-4L6.76 11.757l1.414-1.414 2.829 2.829 5.656-5.657 1.415 1.414L11.003 16z\" fill=\"rgba(47,204,113,1)\"/></svg></div> \n                <div class=\"dz-error-mark\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" width=\"54\" height=\"54\"><path fill=\"none\" d=\"M0 0h24v24H0z\"/><path d=\"M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10-4.477 10-10 10zm0-2a8 8 0 1 0 0-16 8 8 0 0 0 0 16zm0-9.414l2.828-2.829 1.415 1.415L13.414 12l2.829 2.828-1.415 1.415L12 13.414l-2.828 2.829-1.415-1.415L10.586 12 7.757 9.172l1.415-1.415L12 10.586z\" fill=\"rgba(231,76,60,1)\"/></svg></div> \n                </div> ";
      var vm = this;
      this.dropzoneHandler = new (dropzone__WEBPACK_IMPORTED_MODULE_20___default())(this.$refs.dropzoneRef, {
        url: '/',
        maxFiles: maxFileCount,
        parallelUploads: 20,
        uploadMultiple: true,
        filesizeBase: 1024,
        maxFilesize: this.maxModelFilesSize / (1024 * 1024),
        timeout: 0,
        accept: function accept(file, done) {
          if (file.name.length > 128) {
            vm.btnFlag = true;
            vm.showUploadErrInfo(true, vm.getDefaultErrTxt());
            done(vm.$t('modelManage.modelFileNameTips'));
          } else {
            done();
          }
        },
        addRemoveLinks: true,
        autoProcessQueue: false,
        dictDefaultMessage: this.$t('modelManage.modelFileUploadDefaultTips'),
        dictFileTooBig: this.$t('modelManage.fileIstoBig'),
        dictRemoveFile: this.$t('modelManage.removeFile'),
        dictMaxFilesExceeded: this.getDefaultErrTxt(),
        previewTemplate: previewTemplate
      });
      this.dropzoneHandler.on("addedfile", function (file) {
        file.fullname = file.name;

        _this2.checkFiles(file);

        setTimeout(function () {
          return _this2.calculateTotalSize([]);
        }, 0);
      });
      this.dropzoneHandler.on("removedfile", function (file) {
        if (_this2.dropzoneHandler.getRejectedFiles().length === 0) {
          _this2.showUploadErrInfo(false);

          _this2.btnFlag = false;
        }

        _this2.checkFiles();

        setTimeout(function () {
          return _this2.calculateTotalSize([]);
        }, 0);
      });
    },
    getDefaultErrTxt: function getDefaultErrTxt() {
      return this.$t('modelManage.modelFileUploadErrTips', {
        maxCount: maxFileCount,
        size: (0,_utils__WEBPACK_IMPORTED_MODULE_23__.transFileSize)(maxModelFilesSize),
        url: 'https://openi.pcl.ac.cn/docs/index.html#/dataset/sdk'
      });
    },
    showUploadErrInfo: function showUploadErrInfo(state, info) {
      if (state) {
        this.uploadErrTxt = info;
        this.showUploadErr = true;
      } else {
        this.uploadErrTxt = '';
        this.showUploadErr = false;
      }
    },
    changeUploadType: function changeUploadType() {
      this.dropzoneHandler.removeAllFiles();
      this.resetFileStatus();
    },
    checkFiles: function checkFiles(file, countStay) {
      var fileList = this.dropzoneHandler.getAcceptedFiles();
      fileList.forEach(function (item) {
        return item.fullname = item.name;
      });
      var filesCount = fileList.length + (file && !countStay ? 1 : 0);

      if (file && filesCount > maxFileCount) {
        this.dropzoneHandler.removeFile(file);
        this.showUploadErrInfo(true, this.getDefaultErrTxt());
        return false;
      }

      if (file && file.size / (1024 * 1024) > this.dropzoneHandler.options.maxFilesize) {
        this.btnFlag = true;
        this.showUploadErrInfo(true, this.getDefaultErrTxt());
        return false;
      }

      return true;
    },
    resetFileStatus: function resetFileStatus() {
      this.uploadFiles = [];
      this.uploadLength = 0;
      this.uploadSuccessLength = 0;
      this.uploadStatusList = [];
      this.canShowError = false;
    },
    updateFileStatus: function updateFileStatus(file, status, progress, infoCode) {
      var _this3 = this;

      var failedInfo = arguments.length > 4 && arguments[4] !== undefined ? arguments[4] : "";
      this.uploadStatusList.forEach(function (item, index) {
        if (item.uploadUuid === file.upload.uuid) {
          _this3.uploadStatusList[index].status = status;
          _this3.uploadStatusList[index].progress = progress;
          _this3.uploadStatusList[index].infoCode = infoCode;
          _this3.uploadStatusList[index].failedInfo = failedInfo;
        }
      });
    },
    calcFileMd5: function calcFileMd5(file) {
      var blobSlice = File.prototype.slice || File.prototype.mozSlice || File.prototype.webkitSlice;
      var chunkSize = md5ChunkSize;
      var chunks = Math.ceil(file.size / chunkSize);
      var spark = new (spark_md5__WEBPACK_IMPORTED_MODULE_21___default().ArrayBuffer)();
      var fileReader = new FileReader();
      file.totalChunkCounts = chunks;

      if (file.size == 0) {
        file.totalChunkCounts = 1;
      }

      var currentChunk = 0;
      this.uploadStatusList.push({
        uploadUuid: file.upload.uuid,
        name: file.fullname,
        status: this.$t('modelManage.calcFileMd5'),
        progress: 0,
        infoCode: 3
      });
      return new Promise(function (resolve, reject) {
        fileReader.onload = function (e) {
          spark.append(e.target.result);
          currentChunk++;

          if (currentChunk < chunks) {
            loadNext();
          } else {
            var md5 = spark.end();
            spark.destroy();
            file.uniqueIdentifier = md5;
            resolve(md5);
          }
        };

        fileReader.onerror = function (e) {
          console.warn(file.fullname + ': calcFileMd5 went wrong.');
          reject(e);
        };

        function loadNext() {
          var start = currentChunk * chunkSize;
          var end = start + chunkSize >= file.size ? file.size : start + chunkSize;
          fileReader.readAsArrayBuffer(blobSlice.call(file, start, end));
        }

        loadNext();
      });
    },
    getChunksInfo: function getChunksInfo(file) {
      return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_22__.getChunks)({
        md5: file.uniqueIdentifier,
        file_name: "".concat(this.formatUploadir).concat(file.fullname),
        subject_id: this.modelId,
        subject_type: this.subjectType
      }).then(function (resonpse) {
        console.log('getChunksInfo', resonpse);
        var res = resonpse.data;

        if (res.code === 0) {
          console.log("xxxxxxxxxxxxxx", res);
          file.uuid = res.data.uuid;
          file.uploaded = res.data.uploaded;
          file.chunks = res.data.chunks;
          file.realName = res.data.fileName;
          return file;
        } else {
          throw new Error(res.msg); // 抛出一个错误
        }
      })["catch"](function (err) {
        console.info('getChunksInfo', err);
        return err;
      });
    },
    newUpload: function newUpload(file) {
      var _this4 = this;

      return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_22__.getNewMultipart)({
        md5: file.uniqueIdentifier,
        file_name: "".concat(this.formatUploadir).concat(file.fullname),
        subject_id: this.modelId,
        subject_type: this.subjectType,
        file_type: file.type,
        size: file.size,
        total_chunk_counts: file.totalChunkCounts
      }).then(function (resonpse) {
        var res = resonpse.data;

        if (res.code === 0) {
          file.uuid = res.data.uuid;

          if (file.uuid) {
            file.chunks = [];

            _this4.breakpointUpload(file);
          } else {
            _this4.uploadError(file, info);

            _this4.updateFileStatus(file, _this4.$t('modelManage.uploadFailed'), 0, 2);
          }

          return file;
        } else {
          if (_this4.canShowError) {
            _this4.canShowError = false;
            throw new Error(res.msg); // 抛出一个错误
          }
        }
      })["catch"](function (err) {
        console.log('getNewMultipart', err);

        _this4.$message({
          type: 'error',
          message: err.message
        });

        _this4.uploading = false;

        _this4.uploadError(file, _this4.$t('modelManage.uploadFailed'));

        _this4.updateFileStatus(file, _this4.$t('modelManage.uploadFailed'), 0, 2);

        return err;
      });
    },
    breakpointUpload: function breakpointUpload(file) {
      var _this5 = this;

      var blobSlice = File.prototype.slice || File.prototype.mozSlice || File.prototype.webkitSlice;
      var fileReader = new FileReader();
      var time = new Date().getTime();
      var chunkSize = uploadChunkSize;
      var chunks = Math.ceil(file.size / chunkSize);

      var _this = this;

      var currentChunk = 0;
      var successChunks = file.chunks; // const successParts = file.chunks.split(",");
      // for (let i = 0; i < successParts.length; i++) {
      //     successChunks[i] = successParts[i].split("-")[0];
      // }

      var urls = [];
      var etags = [];

      var checkSuccessChunks = function checkSuccessChunks() {
        var index = successChunks.indexOf((currentChunk + 1).toString());

        if (index == -1) {
          return false;
        }

        return true;
      };

      var getUploadChunkUrl = /*#__PURE__*/function () {
        var _ref = _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().mark(function _callee(currentChunk, partSize) {
          var resonpse, res;
          return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().wrap(function _callee$(_context) {
            while (1) {
              switch (_context.prev = _context.next) {
                case 0:
                  _context.prev = 0;
                  _context.next = 3;
                  return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_22__.getMultipartUrl)({
                    uuid: file.uuid,
                    size: partSize,
                    chunk_number: currentChunk + 1
                  });

                case 3:
                  resonpse = _context.sent;
                  res = resonpse.data;

                  if (!(res.code == 0)) {
                    _context.next = 9;
                    break;
                  }

                  urls[currentChunk] = res.data.url;
                  _context.next = 10;
                  break;

                case 9:
                  throw new Error(res.message);

                case 10:
                  console.log("getUploadChunkUrl", urls, etags);
                  _context.next = 16;
                  break;

                case 13:
                  _context.prev = 13;
                  _context.t0 = _context["catch"](0);

                  _this5.$message.error(_context.t0);

                case 16:
                case "end":
                  return _context.stop();
              }
            }
          }, _callee, null, [[0, 13]]);
        }));

        return function getUploadChunkUrl(_x, _x2) {
          return _ref.apply(this, arguments);
        };
      }();

      var uploadMinioNewMethod = /*#__PURE__*/function () {
        var _ref2 = _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().mark(function _callee2(url, e) {
          var xhr;
          return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().wrap(function _callee2$(_context2) {
            while (1) {
              switch (_context2.prev = _context2.next) {
                case 0:
                  xhr = new XMLHttpRequest();
                  xhr.open("PUT", url, false);
                  xhr.send(e.target.result); // if (_this.state.type == 0) {
                  //     xhr.setRequestHeader("Content-Type", "text/plain");
                  //     const etagValue = xhr.getResponseHeader("etag");
                  //     etags[currentChunk] = etagValue;
                  // } else if (_this.state.type == 1) {
                  //     xhr.setRequestHeader("Content-Type", "");
                  //     xhr.send(e.target.result);
                  //     const etagValue = xhr.getResponseHeader("ETag");
                  //     etags[currentChunk] = etagValue;
                  // }

                case 3:
                case "end":
                  return _context2.stop();
              }
            }
          }, _callee2);
        }));

        return function uploadMinioNewMethod(_x3, _x4) {
          return _ref2.apply(this, arguments);
        };
      }();

      var uploadChunk = /*#__PURE__*/function () {
        var _ref3 = _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().mark(function _callee3(e) {
          var start, partSize;
          return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().wrap(function _callee3$(_context3) {
            while (1) {
              switch (_context3.prev = _context3.next) {
                case 0:
                  _context3.prev = 0;

                  if (checkSuccessChunks()) {
                    _context3.next = 14;
                    break;
                  }

                  start = currentChunk * chunkSize;
                  partSize = start + chunkSize >= file.size ? file.size - start : chunkSize; // 获取分片上传url

                  _context3.next = 6;
                  return getUploadChunkUrl(currentChunk, partSize);

                case 6:
                  if (!(urls[currentChunk] != '')) {
                    _context3.next = 13;
                    break;
                  }

                  console.log("ssssssssssssssssssss");
                  _context3.next = 10;
                  return uploadMinioNewMethod(urls[currentChunk], e);

                case 10:
                  if (etags[currentChunk] != '') {} else {
                    console.log("上传到minio uploadChunk etags[currentChunk] == ''"); // TODO
                  }

                  _context3.next = 14;
                  break;

                case 13:
                  console.log("uploadChunk urls[currentChunk] != ''"); // TODO

                case 14:
                  _context3.next = 20;
                  break;

                case 16:
                  _context3.prev = 16;
                  _context3.t0 = _context3["catch"](0);
                  console.log(_context3.t0);

                  _this5.$message({
                    type: 'error',
                    message: err
                  });

                case 20:
                case "end":
                  return _context3.stop();
              }
            }
          }, _callee3, null, [[0, 16]]);
        }));

        return function uploadChunk(_x5) {
          return _ref3.apply(this, arguments);
        };
      }();

      var completeUpload = /*#__PURE__*/function () {
        var _ref4 = _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().mark(function _callee4() {
          return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().wrap(function _callee4$(_context4) {
            while (1) {
              switch (_context4.prev = _context4.next) {
                case 0:
                  _context4.next = 2;
                  return (0,_apis_modules_dataset__WEBPACK_IMPORTED_MODULE_22__.setCompleteMultipart)({
                    uuid: file.uuid
                  });

                case 2:
                  return _context4.abrupt("return", _context4.sent);

                case 3:
                case "end":
                  return _context4.stop();
              }
            }
          }, _callee4);
        }));

        return function completeUpload() {
          return _ref4.apply(this, arguments);
        };
      }();

      function loadNext() {
        var start = currentChunk * chunkSize;
        var end = start + chunkSize >= file.size ? file.size : start + chunkSize;
        fileReader.readAsArrayBuffer(blobSlice.call(file, start, end));
      }

      fileReader.onload = /*#__PURE__*/function () {
        var _ref5 = _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().mark(function _callee5(e) {
          var response, _info, _info2;

          return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().wrap(function _callee5$(_context5) {
            while (1) {
              switch (_context5.prev = _context5.next) {
                case 0:
                  _context5.prev = 0;
                  _context5.next = 3;
                  return uploadChunk(e);

                case 3:
                  fileReader.abort();
                  currentChunk++;

                  if (!(currentChunk < chunks)) {
                    _context5.next = 11;
                    break;
                  }

                  console.log("\u7B2C".concat(currentChunk, "\u4E2A\u5206\u7247\u4E0A\u4F20\u5B8C\u6210, \u5F00\u59CB\u7B2C").concat(currentChunk + 1, "/").concat(chunks, "\u4E2A\u5206\u7247\u4E0A\u4F20"));

                  _this5.updateFileStatus(file, _this5.$t('modelManage.uploading'), Number((currentChunk / chunks * 100).toFixed(2)), 3);

                  loadNext();
                  _context5.next = 31;
                  break;

                case 11:
                  _context5.prev = 11;
                  _context5.next = 14;
                  return completeUpload();

                case 14:
                  response = _context5.sent;
                  console.log("xxxxxxxxxxa", response);
                  console.log("\u6587\u4EF6\u4E0A\u4F20\u5B8C\u6210\uFF1A".concat(file.fullname, " \n\u5206\u7247\uFF1A").concat(chunks, " \u5927\u5C0F:").concat(file.size, " \u7528\u65F6\uFF1A").concat((new Date().getTime() - time) / 1000, " s"));
                  _this5.uploadLength++;
                  _this5.uploadSuccessLength++;

                  _this5.updateFileStatus(file, _this5.$t('modelManage.uploadSuccess'), 100, 0);

                  _this5.uploadSuccess(file);

                  _context5.next = 31;
                  break;

                case 23:
                  _context5.prev = 23;
                  _context5.t0 = _context5["catch"](11);
                  _info = _this5.$t('modelManage.uploadFailed');
                  console.log(_info, file);
                  _this5.uploadLength++;

                  _this5.uploadError(file, _info);

                  _this5.updateFileStatus(file, _info, Number((currentChunk / chunks * 100).toFixed(2)) - 1, 2);

                  _this5.$message({
                    type: 'error',
                    message: _context5.t0
                  });

                case 31:
                  _context5.next = 42;
                  break;

                case 33:
                  _context5.prev = 33;
                  _context5.t1 = _context5["catch"](0);
                  console.log(_context5.t1);
                  _info2 = _this5.$t('modelManage.uploadFailed');
                  console.log(_info2, file);
                  _this5.uploadLength++;

                  _this5.uploadError(file, _info2);

                  _this5.updateFileStatus(file, _info2, Number((currentChunk / chunks * 100).toFixed(2)) - 1, 2);

                  _this5.$message({
                    type: 'error',
                    message: _context5.t1
                  });

                case 42:
                case "end":
                  return _context5.stop();
              }
            }
          }, _callee5, null, [[0, 33], [11, 23]]);
        }));

        return function (_x6) {
          return _ref5.apply(this, arguments);
        };
      }();

      console.log("上传分片...");
      loadNext();
    },
    uploadError: function uploadError(file, info) {
      file.status = 'error';
      file.errTips = info;

      if (file.previewTemplate) {
        file.previewTemplate.querySelector('.dz-success-mark').style.opacity = 0;
        file.previewTemplate.querySelector('.dz-error-mark').style.opacity = 1;
        file.previewTemplate.querySelector('.dz-error-message span').innerHTML = info;
        file.previewTemplate.querySelector(".dz-error-message").style.display = 'block';

        file.previewTemplate.querySelector(".dz-details").onmouseover = function () {
          file.previewTemplate.querySelector('.dz-error-message').style.opacity = 1;
        };

        file.previewTemplate.querySelector(".dz-details").onmouseout = function () {
          file.previewTemplate.querySelector('.dz-error-message').style.opacity = 0;
        };
      }

      this.uploadFinishCheck(file);
    },
    uploadSuccess: function uploadSuccess(file) {
      file.status = 'success';
      file.errTips = '';

      if (file.previewTemplate) {
        file.previewTemplate.querySelector('.dz-error-mark').style.opacity = 0;
        file.previewTemplate.querySelector('.dz-error-message span').innerHTML = '';
        file.previewTemplate.querySelector('.dz-success-mark').style.opacity = 1;
        file.previewTemplate.querySelector(".dz-error-message").style.display = 'none';
        file.previewTemplate.querySelector(".dz-details").onmouseover = null;
        file.previewTemplate.querySelector(".dz-details").onmouseout = null;
      }

      this.uploadFinishCheck(file);
    },
    uploadFinishCheck: function uploadFinishCheck(file) {
      console.log('uploadFinishCheck', file, this.uploadLength, '/', this.uploadFiles.length);

      if (this.uploadLength === this.uploadFiles.length && this.uploadFiles.length != 0) {
        console.log('All file has finish, success ' + this.uploadSuccessLength);
        this.uploading = false;

        if (this.uploadSuccessLength == this.uploadLength) {
          this.$emit('uploadFinish'); // window.setTimeout(() => {
          //     location.href = `/explore/datasets/${this.modelId}?tab=files`
          // }, 1000);
        } else {
          if (this.uploadSuccessLength > 0) {}
        }
      }
    },
    submit: function submit() {
      var _this6 = this;

      var fileList = [];

      if (this.uploadType == 'file') {
        fileList = this.dropzoneHandler.getAcceptedFiles();
        if (!fileList.length) return;

        for (var i = 0, iLen = fileList.length; i < iLen; i++) {
          if (!this.checkFiles(fileList[i], true)) return;
        }
      } else if (this.uploadType == 'folder') {
        fileList = this.$refs['folderUploadSelect'].getAcceptedFiles();
        if (!fileList.length) return;
        if (!this.$refs['folderUploadSelect'].checkFiles()) return;
      }

      this.resetFileStatus();
      this.canShowError = true;
      this.uploadFiles = fileList;
      this.uploading = true;
      var fileNameList = [];

      var _loop = function _loop(_i, _iLen) {
        var file = fileList[_i]; // file.modelUuid = this.state.id;
        // file.modelName = this.state.name;

        _this6.calcFileMd5(file).then(function (res) {
          // 计算MD5
          if (fileNameList.indexOf(file.fullname) > -1) {
            var info = "".concat(_this6.$t('modelManage.fileExistInTheModel'), " ").concat(file.fullname);
            console.log("info=" + info);
            _this6.uploadLength++;

            _this6.uploadError(file, _this6.$t('modelManage.uploadFailed'));

            _this6.updateFileStatus(file, _this6.$t('modelManage.uploadFailed'), 0, 1, info);
          } else {
            fileNameList.push(file.fullname);

            _this6.getChunksInfo(file).then(function (res) {
              // 获取Chunk信息
              console.log(file);

              if (file.uuid == '') {
                // 未上传过
                console.log("file.uuid == ''", file.uuid == '');

                _this6.newUpload(file);
              } else if (file.uploaded == '1') {
                // 已上传成功 
                // if (file.attachID == '0') { // 删除数据集记录，未删除文件
                //     // await addAttachment(file);
                //     console.log(`file.attachID == '0'`);
                // }
                _this6.uploadLength++; // 同一模型上传同一个文件

                if (file.subjectId) {
                  var _info3 = "".concat(_this6.$t('modelManage.fileExistInTheModel'), " ").concat(file.realName);

                  _this6.uploadError(file, _info3);

                  _this6.updateFileStatus(file, _this6.$t('modelManage.uploadFailed'), 0, 1, _info3);
                } else {
                  // 秒传
                  _this6.uploadSuccessLength++;

                  _this6.updateFileStatus(file, _this6.$t('modelManage.uploadSuccess'), 100, 0);

                  _this6.uploadSuccess(file);
                }

                console.log(file.fullname, '文件处理完成');
              } else {
                // 断点续传
                _this6.breakpointUpload(file);
              }
            })["catch"](function (err) {
              console.log("xxxxxxxxxxxxxssssss");
              console.info('getChunksInfo', err);
              _this6.uploadLength++;

              _this6.uploadError(file, _this6.$t('modelManage.uploadFailed'));

              _this6.updateFileStatus(file, _this6.$t('modelManage.uploadFailed'), 0, 2);

              _this6.$message({
                type: 'error',
                message: err
              });
            });
          }
        })["catch"](function (err) {
          console.info('calcFileMd5', err);
          _this6.uploadLength++;

          _this6.uploadError(file, _this6.$t('modelManage.uploadFailed'));

          _this6.updateFileStatus(file, _this6.$t('modelManage.uploadFailed'), 0, 2);
        });
      };

      for (var _i = 0, _iLen = fileList.length; _i < _iLen; _i++) {
        _loop(_i, _iLen);
      }
    },
    cancel: function cancel() {
      this.$emit('cancel');
    },
    getStorageSummary: function getStorageSummary() {
      var _this7 = this;

      return _babel_runtime_helpers_asyncToGenerator__WEBPACK_IMPORTED_MODULE_17___default()( /*#__PURE__*/_babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().mark(function _callee6() {
        var res;
        return _babel_runtime_regenerator__WEBPACK_IMPORTED_MODULE_15___default().wrap(function _callee6$(_context6) {
          while (1) {
            switch (_context6.prev = _context6.next) {
              case 0:
                _context6.next = 2;
                return (0,_apis_modules_storage__WEBPACK_IMPORTED_MODULE_24__.getStorageSummary)({
                  subject_id: _this7.modelId,
                  subject_type: _this7.subjectType
                });

              case 2:
                res = _context6.sent;
                _this7.remaining_storage = res.data.remaining_storage;
                _this7.storageShow = res.data.storage_limit != -1;

              case 5:
              case "end":
                return _context6.stop();
            }
          }
        }, _callee6);
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
    calculateTotalSize: function calculateTotalSize(files) {
      if (!this.storageShow) return;

      if (files.length === 0) {
        files = this.dropzoneHandler ? this.dropzoneHandler.getQueuedFiles() : [];
      }

      this.totalSize = files.reduce(function (sum, file) {
        return sum + file.size;
      }, 0);
      this.isStorageExceeded = this.totalSize > this.remaining_storage;
    },
    folderSelectChange: function folderSelectChange(fileList) {
      this.calculateTotalSize(fileList);
    }
  },
  beforeMount: function beforeMount() {},
  mounted: function mounted() {
    this.initDropZone();
    this.getStorageSummary();
    var loginName = document.querySelector('meta[name="_uid"]').getAttribute('content-ext');
    console.log('loginName', loginName, this.ownerName);

    if (this.ownerName != loginName) {
      this.showOwenerTips = true;
    }
  },
  beforeDestroy: function beforeDestroy() {}
});

/***/ }),

/***/ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=script&lang=js":
/*!**********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.symbol */ "./node_modules/core-js/modules/es.symbol.js");
/* harmony import */ var core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.symbol.description */ "./node_modules/core-js/modules/es.symbol.description.js");
/* harmony import */ var core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol_description__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_symbol_iterator__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.symbol.iterator */ "./node_modules/core-js/modules/es.symbol.iterator.js");
/* harmony import */ var core_js_modules_es_symbol_iterator__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_symbol_iterator__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.find-index */ "./node_modules/core-js/modules/es.array.find-index.js");
/* harmony import */ var core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find_index__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_array_from__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.array.from */ "./node_modules/core-js/modules/es.array.from.js");
/* harmony import */ var core_js_modules_es_array_from__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_from__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.array.is-array */ "./node_modules/core-js/modules/es.array.is-array.js");
/* harmony import */ var core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_is_array__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.array.iterator */ "./node_modules/core-js/modules/es.array.iterator.js");
/* harmony import */ var core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_iterator__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.array.reduce */ "./node_modules/core-js/modules/es.array.reduce.js");
/* harmony import */ var core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_reduce__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/es.array.slice */ "./node_modules/core-js/modules/es.array.slice.js");
/* harmony import */ var core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_slice__WEBPACK_IMPORTED_MODULE_9__);
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! core-js/modules/es.array.splice */ "./node_modules/core-js/modules/es.array.splice.js");
/* harmony import */ var core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_10___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_splice__WEBPACK_IMPORTED_MODULE_10__);
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_11__ = __webpack_require__(/*! core-js/modules/es.date.to-string */ "./node_modules/core-js/modules/es.date.to-string.js");
/* harmony import */ var core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_11___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_date_to_string__WEBPACK_IMPORTED_MODULE_11__);
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_12__ = __webpack_require__(/*! core-js/modules/es.function.name */ "./node_modules/core-js/modules/es.function.name.js");
/* harmony import */ var core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_12___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_function_name__WEBPACK_IMPORTED_MODULE_12__);
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_13__ = __webpack_require__(/*! core-js/modules/es.number.constructor */ "./node_modules/core-js/modules/es.number.constructor.js");
/* harmony import */ var core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_13___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_number_constructor__WEBPACK_IMPORTED_MODULE_13__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_14__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_14___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_14__);
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_15__ = __webpack_require__(/*! core-js/modules/es.regexp.to-string */ "./node_modules/core-js/modules/es.regexp.to-string.js");
/* harmony import */ var core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_15___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_to_string__WEBPACK_IMPORTED_MODULE_15__);
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_16__ = __webpack_require__(/*! core-js/modules/es.string.iterator */ "./node_modules/core-js/modules/es.string.iterator.js");
/* harmony import */ var core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_16___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_iterator__WEBPACK_IMPORTED_MODULE_16__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_17__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_17___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_17__);
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_18__ = __webpack_require__(/*! core-js/modules/web.dom-collections.iterator */ "./node_modules/core-js/modules/web.dom-collections.iterator.js");
/* harmony import */ var core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_18___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_iterator__WEBPACK_IMPORTED_MODULE_18__);
/* harmony import */ var _utils__WEBPACK_IMPORTED_MODULE_19__ = __webpack_require__(/*! ~/utils */ "./web_src/vuepages/utils/index.js");




















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

/* harmony default export */ __webpack_exports__["default"] = ({
  name: "FolderUploadSelect",
  props: {
    uploading: {
      type: Boolean,
      "default": false
    },
    maxFilesCount: {
      type: Number,
      "default": 100
    },
    maxFilesSize: {
      type: Number,
      "default": 200 * 1024 * 1024 * 1024
    }
  },
  data: function data() {
    return {
      fileList: [],
      errorState: false,
      errorInfo: '',
      errorList: []
    };
  },
  computed: {
    allFilesSize: function allFilesSize() {
      var allFilesSize = this.fileList.reduce(function (acc, item, index) {
        return acc + item.size;
      }, 0);
      return (0,_utils__WEBPACK_IMPORTED_MODULE_19__.transFileSize)(allFilesSize);
    }
  },
  methods: {
    addFolder: function addFolder() {
      this.$refs['filepicker'].value = '';
      this.$refs['filepicker'].click();
    },
    folderSelectChange: function folderSelectChange(evt) {
      var files = evt.target.files || [];

      for (var i = 0, iLen = files.length; i < iLen; i++) {
        var file = files[i];
        file._webkitRelativePath = file.webkitRelativePath;
      }

      this.handleFiles(files);
    },
    handleFiles: function handleFiles(files) {
      var _this = this;

      var _loop = function _loop(i, iLen) {
        var file = files[i];
        if (_this.fileList.findIndex(function (item) {
          return item._webkitRelativePath == file._webkitRelativePath;
        }) >= 0) return "continue";
        file.fullname = file._webkitRelativePath ? file._webkitRelativePath : file.name;
        file.file_size = (0,_utils__WEBPACK_IMPORTED_MODULE_19__.transFileSize)(file.size);
        file.upload = {
          uuid: (0,_utils__WEBPACK_IMPORTED_MODULE_19__.uuidv4)()
        };

        _this.fileList.push(file);
      };

      for (var i = 0, iLen = files.length; i < iLen; i++) {
        var _ret = _loop(i, iLen);

        if (_ret === "continue") continue;
      }

      this.checkFiles();
      this.$emit('folderSelectChange', this.fileList);
    },
    handleDragover: function handleDragover(evt) {
      evt.preventDefault();
    },
    handleDrop: function handleDrop(evt) {
      evt.preventDefault();
      var filesList = [];
      var dataTransferItemList = evt.dataTransfer.items;

      var _iterator = _createForOfIteratorHelper(dataTransferItemList),
          _step;

      try {
        for (_iterator.s(); !(_step = _iterator.n()).done;) {
          var dataTransferItem = _step.value;
          var fileEntry = dataTransferItem.webkitGetAsEntry();
          this.handleDropedFileEntry(fileEntry, filesList);
        }
      } catch (err) {
        _iterator.e(err);
      } finally {
        _iterator.f();
      }
    },
    handleDropedFileEntry: function handleDropedFileEntry(fileEntry) {
      var self = this;

      if (fileEntry.isFile) {
        fileEntry.file(function (file) {
          file._webkitRelativePath = fileEntry.fullPath.slice(1);
          self.handleFiles([file]);
          return;
        });
      } else {
        var dirReader = fileEntry.createReader();
        dirReader.readEntries(function (entries) {
          for (var i = 0; i < entries.length; i++) {
            self.handleDropedFileEntry(entries[i]);
          }
        });
      }
    },
    remove: function remove(file) {
      if (file) {
        var index = this.fileList.findIndex(function (item) {
          return item._webkitRelativePath === file._webkitRelativePath;
        });
        this.fileList.splice(index, 1);
      } else {
        this.fileList.splice(0);
      }

      this.checkFiles();
      this.$emit('folderSelectChange', this.fileList);
    },
    showErrInfo: function showErrInfo(state, info) {
      this.errorState = state;
      this.errorInfo = info;
    },
    getDefaultErrTxt: function getDefaultErrTxt() {
      return this.$t('modelManage.modelFileUploadErrTips', {
        maxCount: this.maxFilesCount,
        size: (0,_utils__WEBPACK_IMPORTED_MODULE_19__.transFileSize)(this.maxFilesSize),
        url: 'https://openi.pcl.ac.cn/docs/index.html#/model/sdk'
      });
    },
    checkFiles: function checkFiles() {
      var _this2 = this;

      var flag = false;
      this.fileList.forEach(function (item) {
        if (item.size > _this2.maxFilesSize || item.name.length > 128) {
          flag = true;
          item.sizeStatus = 'error';
        }
      });

      if (this.fileList.length > this.maxFilesCount || flag) {
        this.showErrInfo(true, this.getDefaultErrTxt());
        return false;
      }

      this.showErrInfo(false, '');
      return true;
    },
    getAcceptedFiles: function getAcceptedFiles() {
      return this.fileList;
    }
  },
  beforeMount: function beforeMount() {}
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

/***/ "./web_src/vuepages/pages/dataset/detail/vp-dataset-detial.js":
/*!********************************************************************!*\
  !*** ./web_src/vuepages/pages/dataset/detail/vp-dataset-detial.js ***!
  \********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var vue__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! vue */ "./node_modules/vue/dist/vue.esm.js");
/* harmony import */ var element_ui__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! element-ui */ "./node_modules/element-ui/lib/element-ui.common.js");
/* harmony import */ var element_ui__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(element_ui__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var element_ui_lib_theme_chalk_index_css__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! element-ui/lib/theme-chalk/index.css */ "./node_modules/element-ui/lib/theme-chalk/index.css");
/* harmony import */ var element_ui_lib_locale_lang_en__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! element-ui/lib/locale/lang/en */ "./node_modules/element-ui/lib/locale/lang/en.js");
/* harmony import */ var element_ui_lib_locale_lang_zh_CN__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! element-ui/lib/locale/lang/zh-CN */ "./node_modules/element-ui/lib/locale/lang/zh-CN.js");
/* harmony import */ var _langs__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ~/langs */ "./web_src/vuepages/langs/index.js");
/* harmony import */ var _index_vue__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ./index.vue */ "./web_src/vuepages/pages/dataset/detail/index.vue");







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

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less":
/*!**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less ***!
  \**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseTitle.vue?vue&type=style&index=0&id=94dd892a&lang=less&scoped=true":
/*!*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseTitle.vue?vue&type=style&index=0&id=94dd892a&lang=less&scoped=true ***!
  \*************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=0&id=10bd582e&scoped=true&lang=less":
/*!********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=0&id=10bd582e&scoped=true&lang=less ***!
  \********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=1&id=10bd582e&lang=css":
/*!*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=1&id=10bd582e&lang=css ***!
  \*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Header.vue?vue&type=style&index=0&id=7b0734e5&scoped=true&lang=less":
/*!************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Header.vue?vue&type=style&index=0&id=7b0734e5&scoped=true&lang=less ***!
  \************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=style&index=0&id=36f478aa&lang=less&scoped=true":
/*!********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=style&index=0&id=36f478aa&lang=less&scoped=true ***!
  \********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=style&index=0&id=6704b8e1&scoped=true&lang=less":
/*!*****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=style&index=0&id=6704b8e1&scoped=true&lang=less ***!
  \*****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=style&index=0&id=a96f77d8&scoped=true&lang=less":
/*!****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=style&index=0&id=a96f77d8&scoped=true&lang=less ***!
  \****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=style&index=0&id=9b01554a&lang=less&scoped=true":
/*!********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=style&index=0&id=9b01554a&lang=less&scoped=true ***!
  \********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=style&index=0&id=5368752e&lang=less&scoped=true":
/*!*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=style&index=0&id=5368752e&lang=less&scoped=true ***!
  \*********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=style&index=0&id=3e16625e&lang=less&scoped=true":
/*!**********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=style&index=0&id=3e16625e&lang=less&scoped=true ***!
  \**********************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=style&index=0&id=2cb6c8b3&lang=less&scoped=true":
/*!*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=style&index=0&id=2cb6c8b3&lang=less&scoped=true ***!
  \*******************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=style&index=0&id=e8b9b4e4&lang=less&scoped=true":
/*!**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=style&index=0&id=e8b9b4e4&lang=less&scoped=true ***!
  \**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=style&index=0&id=70071dcb&scoped=true&lang=less":
/*!**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=style&index=0&id=70071dcb&scoped=true&lang=less ***!
  \**************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
// extracted by mini-css-extract-plugin


/***/ }),

/***/ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=style&index=0&id=1a87d32a&scoped=true&lang=less":
/*!****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=style&index=0&id=1a87d32a&scoped=true&lang=less ***!
  \****************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************************/
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

/***/ "./web_src/vuepages/components/BaseTitle.vue":
/*!***************************************************!*\
  !*** ./web_src/vuepages/components/BaseTitle.vue ***!
  \***************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _BaseTitle_vue_vue_type_template_id_94dd892a_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./BaseTitle.vue?vue&type=template&id=94dd892a&scoped=true */ "./web_src/vuepages/components/BaseTitle.vue?vue&type=template&id=94dd892a&scoped=true");
/* harmony import */ var _BaseTitle_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./BaseTitle.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/BaseTitle.vue?vue&type=script&lang=js");
/* harmony import */ var _BaseTitle_vue_vue_type_style_index_0_id_94dd892a_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./BaseTitle.vue?vue&type=style&index=0&id=94dd892a&lang=less&scoped=true */ "./web_src/vuepages/components/BaseTitle.vue?vue&type=style&index=0&id=94dd892a&lang=less&scoped=true");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _BaseTitle_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _BaseTitle_vue_vue_type_template_id_94dd892a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _BaseTitle_vue_vue_type_template_id_94dd892a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "94dd892a",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/BaseTitle.vue"
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

/***/ "./web_src/vuepages/components/square/CommonSDK.vue":
/*!**********************************************************!*\
  !*** ./web_src/vuepages/components/square/CommonSDK.vue ***!
  \**********************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _CommonSDK_vue_vue_type_template_id_10bd582e_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./CommonSDK.vue?vue&type=template&id=10bd582e&scoped=true */ "./web_src/vuepages/components/square/CommonSDK.vue?vue&type=template&id=10bd582e&scoped=true");
/* harmony import */ var _CommonSDK_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./CommonSDK.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/square/CommonSDK.vue?vue&type=script&lang=js");
/* harmony import */ var _CommonSDK_vue_vue_type_style_index_0_id_10bd582e_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./CommonSDK.vue?vue&type=style&index=0&id=10bd582e&scoped=true&lang=less */ "./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=0&id=10bd582e&scoped=true&lang=less");
/* harmony import */ var _CommonSDK_vue_vue_type_style_index_1_id_10bd582e_lang_css__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./CommonSDK.vue?vue&type=style&index=1&id=10bd582e&lang=css */ "./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=1&id=10bd582e&lang=css");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! !../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;



/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_4__["default"])(
  _CommonSDK_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _CommonSDK_vue_vue_type_template_id_10bd582e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _CommonSDK_vue_vue_type_template_id_10bd582e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "10bd582e",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/square/CommonSDK.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/Header.vue":
/*!**************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/Header.vue ***!
  \**************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Header_vue_vue_type_template_id_7b0734e5_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Header.vue?vue&type=template&id=7b0734e5&scoped=true */ "./web_src/vuepages/components/square/detail/Header.vue?vue&type=template&id=7b0734e5&scoped=true");
/* harmony import */ var _Header_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Header.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/square/detail/Header.vue?vue&type=script&lang=js");
/* harmony import */ var _Header_vue_vue_type_style_index_0_id_7b0734e5_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Header.vue?vue&type=style&index=0&id=7b0734e5&scoped=true&lang=less */ "./web_src/vuepages/components/square/detail/Header.vue?vue&type=style&index=0&id=7b0734e5&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Header_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Header_vue_vue_type_template_id_7b0734e5_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Header_vue_vue_type_template_id_7b0734e5_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "7b0734e5",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/square/detail/Header.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/Icons.vue":
/*!*************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/Icons.vue ***!
  \*************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Icons_vue_vue_type_template_id_405d1b1c__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Icons.vue?vue&type=template&id=405d1b1c */ "./web_src/vuepages/components/square/detail/Icons.vue?vue&type=template&id=405d1b1c");
/* harmony import */ var _Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Icons.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/square/detail/Icons.vue?vue&type=script&lang=js");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");





/* normalize component */
;
var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_2__["default"])(
  _Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Icons_vue_vue_type_template_id_405d1b1c__WEBPACK_IMPORTED_MODULE_0__.render,
  _Icons_vue_vue_type_template_id_405d1b1c__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  null,
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/square/detail/Icons.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/files/FileList.vue":
/*!**********************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/files/FileList.vue ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _FileList_vue_vue_type_template_id_36f478aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./FileList.vue?vue&type=template&id=36f478aa&scoped=true */ "./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=template&id=36f478aa&scoped=true");
/* harmony import */ var _FileList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./FileList.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=script&lang=js");
/* harmony import */ var _FileList_vue_vue_type_style_index_0_id_36f478aa_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./FileList.vue?vue&type=style&index=0&id=36f478aa&lang=less&scoped=true */ "./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=style&index=0&id=36f478aa&lang=less&scoped=true");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _FileList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _FileList_vue_vue_type_template_id_36f478aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _FileList_vue_vue_type_template_id_36f478aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "36f478aa",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/square/detail/files/FileList.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/intro/Intro.vue":
/*!*******************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/intro/Intro.vue ***!
  \*******************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Intro_vue_vue_type_template_id_6704b8e1_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Intro.vue?vue&type=template&id=6704b8e1&scoped=true */ "./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=template&id=6704b8e1&scoped=true");
/* harmony import */ var _Intro_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Intro.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=script&lang=js");
/* harmony import */ var _Intro_vue_vue_type_style_index_0_id_6704b8e1_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Intro.vue?vue&type=style&index=0&id=6704b8e1&scoped=true&lang=less */ "./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=style&index=0&id=6704b8e1&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Intro_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Intro_vue_vue_type_template_id_6704b8e1_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Intro_vue_vue_type_template_id_6704b8e1_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "6704b8e1",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/square/detail/intro/Intro.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue":
/*!******************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue ***!
  \******************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _MigrateModelSync_vue_vue_type_template_id_a96f77d8_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./MigrateModelSync.vue?vue&type=template&id=a96f77d8&scoped=true */ "./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=template&id=a96f77d8&scoped=true");
/* harmony import */ var _MigrateModelSync_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./MigrateModelSync.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=script&lang=js");
/* harmony import */ var _MigrateModelSync_vue_vue_type_style_index_0_id_a96f77d8_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./MigrateModelSync.vue?vue&type=style&index=0&id=a96f77d8&scoped=true&lang=less */ "./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=style&index=0&id=a96f77d8&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _MigrateModelSync_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _MigrateModelSync_vue_vue_type_template_id_a96f77d8_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _MigrateModelSync_vue_vue_type_template_id_a96f77d8_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "a96f77d8",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/setting/Access.vue":
/*!**********************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/setting/Access.vue ***!
  \**********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Access_vue_vue_type_template_id_9b01554a_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Access.vue?vue&type=template&id=9b01554a&scoped=true */ "./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=template&id=9b01554a&scoped=true");
/* harmony import */ var _Access_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Access.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=script&lang=js");
/* harmony import */ var _Access_vue_vue_type_style_index_0_id_9b01554a_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Access.vue?vue&type=style&index=0&id=9b01554a&lang=less&scoped=true */ "./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=style&index=0&id=9b01554a&lang=less&scoped=true");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Access_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Access_vue_vue_type_template_id_9b01554a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Access_vue_vue_type_template_id_9b01554a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "9b01554a",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/square/detail/setting/Access.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/setting/Setting.vue":
/*!***********************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/setting/Setting.vue ***!
  \***********************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _Setting_vue_vue_type_template_id_5368752e_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./Setting.vue?vue&type=template&id=5368752e&scoped=true */ "./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=template&id=5368752e&scoped=true");
/* harmony import */ var _Setting_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./Setting.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=script&lang=js");
/* harmony import */ var _Setting_vue_vue_type_style_index_0_id_5368752e_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Setting.vue?vue&type=style&index=0&id=5368752e&lang=less&scoped=true */ "./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=style&index=0&id=5368752e&lang=less&scoped=true");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _Setting_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _Setting_vue_vue_type_template_id_5368752e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _Setting_vue_vue_type_template_id_5368752e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "5368752e",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/square/detail/setting/Setting.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/usage/UsageIntro.vue":
/*!************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/usage/UsageIntro.vue ***!
  \************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _UsageIntro_vue_vue_type_template_id_3e16625e_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./UsageIntro.vue?vue&type=template&id=3e16625e&scoped=true */ "./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=template&id=3e16625e&scoped=true");
/* harmony import */ var _UsageIntro_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./UsageIntro.vue?vue&type=script&lang=js */ "./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=script&lang=js");
/* harmony import */ var _UsageIntro_vue_vue_type_style_index_0_id_3e16625e_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./UsageIntro.vue?vue&type=style&index=0&id=3e16625e&lang=less&scoped=true */ "./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=style&index=0&id=3e16625e&lang=less&scoped=true");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _UsageIntro_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _UsageIntro_vue_vue_type_template_id_3e16625e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _UsageIntro_vue_vue_type_template_id_3e16625e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "3e16625e",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/components/square/detail/usage/UsageIntro.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/dataset/detail/index.vue":
/*!*********************************************************!*\
  !*** ./web_src/vuepages/pages/dataset/detail/index.vue ***!
  \*********************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _index_vue_vue_type_template_id_2cb6c8b3_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./index.vue?vue&type=template&id=2cb6c8b3&scoped=true */ "./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=template&id=2cb6c8b3&scoped=true");
/* harmony import */ var _index_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./index.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=script&lang=js");
/* harmony import */ var _index_vue_vue_type_style_index_0_id_2cb6c8b3_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./index.vue?vue&type=style&index=0&id=2cb6c8b3&lang=less&scoped=true */ "./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=style&index=0&id=2cb6c8b3&lang=less&scoped=true");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _index_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _index_vue_vue_type_template_id_2cb6c8b3_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _index_vue_vue_type_template_id_2cb6c8b3_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "2cb6c8b3",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/dataset/detail/index.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/guide/components/CreateForm.vue":
/*!****************************************************************!*\
  !*** ./web_src/vuepages/pages/guide/components/CreateForm.vue ***!
  \****************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _CreateForm_vue_vue_type_template_id_e8b9b4e4_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./CreateForm.vue?vue&type=template&id=e8b9b4e4&scoped=true */ "./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=template&id=e8b9b4e4&scoped=true");
/* harmony import */ var _CreateForm_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./CreateForm.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=script&lang=js");
/* harmony import */ var _CreateForm_vue_vue_type_style_index_0_id_e8b9b4e4_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./CreateForm.vue?vue&type=style&index=0&id=e8b9b4e4&lang=less&scoped=true */ "./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=style&index=0&id=e8b9b4e4&lang=less&scoped=true");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _CreateForm_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _CreateForm_vue_vue_type_template_id_e8b9b4e4_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _CreateForm_vue_vue_type_template_id_e8b9b4e4_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "e8b9b4e4",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/guide/components/CreateForm.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/guide/components/FileUpload.vue":
/*!****************************************************************!*\
  !*** ./web_src/vuepages/pages/guide/components/FileUpload.vue ***!
  \****************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _FileUpload_vue_vue_type_template_id_70071dcb_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./FileUpload.vue?vue&type=template&id=70071dcb&scoped=true */ "./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=template&id=70071dcb&scoped=true");
/* harmony import */ var _FileUpload_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./FileUpload.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=script&lang=js");
/* harmony import */ var _FileUpload_vue_vue_type_style_index_0_id_70071dcb_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./FileUpload.vue?vue&type=style&index=0&id=70071dcb&scoped=true&lang=less */ "./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=style&index=0&id=70071dcb&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _FileUpload_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _FileUpload_vue_vue_type_template_id_70071dcb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _FileUpload_vue_vue_type_template_id_70071dcb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "70071dcb",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/guide/components/FileUpload.vue"
/* harmony default export */ __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ "./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue":
/*!******************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue ***!
  \******************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _FolderUploadSelect_vue_vue_type_template_id_1a87d32a_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./FolderUploadSelect.vue?vue&type=template&id=1a87d32a&scoped=true */ "./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=template&id=1a87d32a&scoped=true");
/* harmony import */ var _FolderUploadSelect_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./FolderUploadSelect.vue?vue&type=script&lang=js */ "./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=script&lang=js");
/* harmony import */ var _FolderUploadSelect_vue_vue_type_style_index_0_id_1a87d32a_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./FolderUploadSelect.vue?vue&type=style&index=0&id=1a87d32a&scoped=true&lang=less */ "./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=style&index=0&id=1a87d32a&scoped=true&lang=less");
/* harmony import */ var _node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! !../../../../../node_modules/vue-loader/lib/runtime/componentNormalizer.js */ "./node_modules/vue-loader/lib/runtime/componentNormalizer.js");



;


/* normalize component */

var component = (0,_node_modules_vue_loader_lib_runtime_componentNormalizer_js__WEBPACK_IMPORTED_MODULE_3__["default"])(
  _FolderUploadSelect_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_1__["default"],
  _FolderUploadSelect_vue_vue_type_template_id_1a87d32a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render,
  _FolderUploadSelect_vue_vue_type_template_id_1a87d32a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns,
  false,
  null,
  "1a87d32a",
  null
  
)

/* hot reload */
if (false) { var api; }
component.options.__file = "web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue"
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

/***/ "./web_src/vuepages/components/BaseTitle.vue?vue&type=script&lang=js":
/*!***************************************************************************!*\
  !*** ./web_src/vuepages/components/BaseTitle.vue?vue&type=script&lang=js ***!
  \***************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseTitle_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./BaseTitle.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseTitle.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseTitle_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

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

/***/ "./web_src/vuepages/components/square/CommonSDK.vue?vue&type=script&lang=js":
/*!**********************************************************************************!*\
  !*** ./web_src/vuepages/components/square/CommonSDK.vue?vue&type=script&lang=js ***!
  \**********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_CommonSDK_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./CommonSDK.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_CommonSDK_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/Header.vue?vue&type=script&lang=js":
/*!**************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/Header.vue?vue&type=script&lang=js ***!
  \**************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Header_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Header.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Header.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Header_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/Icons.vue?vue&type=script&lang=js":
/*!*************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/Icons.vue?vue&type=script&lang=js ***!
  \*************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Icons.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Icons.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=script&lang=js":
/*!**********************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_FileList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FileList.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_FileList_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=script&lang=js":
/*!*******************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=script&lang=js ***!
  \*******************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Intro_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Intro.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Intro_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=script&lang=js":
/*!******************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=script&lang=js ***!
  \******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_MigrateModelSync_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./MigrateModelSync.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_MigrateModelSync_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=script&lang=js":
/*!**********************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=script&lang=js ***!
  \**********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Access_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Access.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Access_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=script&lang=js":
/*!***********************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=script&lang=js ***!
  \***********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Setting_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Setting.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_Setting_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=script&lang=js":
/*!************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=script&lang=js ***!
  \************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_UsageIntro_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./UsageIntro.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_UsageIntro_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=script&lang=js":
/*!*********************************************************************************!*\
  !*** ./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=script&lang=js ***!
  \*********************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./index.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=script&lang=js":
/*!****************************************************************************************!*\
  !*** ./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=script&lang=js ***!
  \****************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_CreateForm_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./CreateForm.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_CreateForm_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=script&lang=js":
/*!****************************************************************************************!*\
  !*** ./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=script&lang=js ***!
  \****************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_FileUpload_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FileUpload.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_FileUpload_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=script&lang=js":
/*!******************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=script&lang=js ***!
  \******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_FolderUploadSelect_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FolderUploadSelect.vue?vue&type=script&lang=js */ "./node_modules/babel-loader/lib/index.js??clonedRuleSet-3.use[0]!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=script&lang=js");
 /* harmony default export */ __webpack_exports__["default"] = (_node_modules_babel_loader_lib_index_js_clonedRuleSet_3_use_0_node_modules_vue_loader_lib_index_js_vue_loader_options_FolderUploadSelect_vue_vue_type_script_lang_js__WEBPACK_IMPORTED_MODULE_0__["default"]); 

/***/ }),

/***/ "./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less":
/*!*************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less ***!
  \*************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseDialog_vue_vue_type_style_index_0_id_165bc8c5_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../node_modules/less-loader/dist/cjs.js!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseDialog.vue?vue&type=style&index=0&id=165bc8c5&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/BaseTitle.vue?vue&type=style&index=0&id=94dd892a&lang=less&scoped=true":
/*!************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/BaseTitle.vue?vue&type=style&index=0&id=94dd892a&lang=less&scoped=true ***!
  \************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseTitle_vue_vue_type_style_index_0_id_94dd892a_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../node_modules/less-loader/dist/cjs.js!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./BaseTitle.vue?vue&type=style&index=0&id=94dd892a&lang=less&scoped=true */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseTitle.vue?vue&type=style&index=0&id=94dd892a&lang=less&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=0&id=10bd582e&scoped=true&lang=less":
/*!*******************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=0&id=10bd582e&scoped=true&lang=less ***!
  \*******************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_CommonSDK_vue_vue_type_style_index_0_id_10bd582e_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../node_modules/less-loader/dist/cjs.js!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./CommonSDK.vue?vue&type=style&index=0&id=10bd582e&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=0&id=10bd582e&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=1&id=10bd582e&lang=css":
/*!******************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=1&id=10bd582e&lang=css ***!
  \******************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_CommonSDK_vue_vue_type_style_index_1_id_10bd582e_lang_css__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../node_modules/less-loader/dist/cjs.js!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./CommonSDK.vue?vue&type=style&index=1&id=10bd582e&lang=css */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=style&index=1&id=10bd582e&lang=css");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/Header.vue?vue&type=style&index=0&id=7b0734e5&scoped=true&lang=less":
/*!***********************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/Header.vue?vue&type=style&index=0&id=7b0734e5&scoped=true&lang=less ***!
  \***********************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Header_vue_vue_type_style_index_0_id_7b0734e5_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Header.vue?vue&type=style&index=0&id=7b0734e5&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Header.vue?vue&type=style&index=0&id=7b0734e5&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=style&index=0&id=36f478aa&lang=less&scoped=true":
/*!*******************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=style&index=0&id=36f478aa&lang=less&scoped=true ***!
  \*******************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_FileList_vue_vue_type_style_index_0_id_36f478aa_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FileList.vue?vue&type=style&index=0&id=36f478aa&lang=less&scoped=true */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=style&index=0&id=36f478aa&lang=less&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=style&index=0&id=6704b8e1&scoped=true&lang=less":
/*!****************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=style&index=0&id=6704b8e1&scoped=true&lang=less ***!
  \****************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Intro_vue_vue_type_style_index_0_id_6704b8e1_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Intro.vue?vue&type=style&index=0&id=6704b8e1&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=style&index=0&id=6704b8e1&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=style&index=0&id=a96f77d8&scoped=true&lang=less":
/*!***************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=style&index=0&id=a96f77d8&scoped=true&lang=less ***!
  \***************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_MigrateModelSync_vue_vue_type_style_index_0_id_a96f77d8_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./MigrateModelSync.vue?vue&type=style&index=0&id=a96f77d8&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=style&index=0&id=a96f77d8&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=style&index=0&id=9b01554a&lang=less&scoped=true":
/*!*******************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=style&index=0&id=9b01554a&lang=less&scoped=true ***!
  \*******************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Access_vue_vue_type_style_index_0_id_9b01554a_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Access.vue?vue&type=style&index=0&id=9b01554a&lang=less&scoped=true */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=style&index=0&id=9b01554a&lang=less&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=style&index=0&id=5368752e&lang=less&scoped=true":
/*!********************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=style&index=0&id=5368752e&lang=less&scoped=true ***!
  \********************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_Setting_vue_vue_type_style_index_0_id_5368752e_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Setting.vue?vue&type=style&index=0&id=5368752e&lang=less&scoped=true */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=style&index=0&id=5368752e&lang=less&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=style&index=0&id=3e16625e&lang=less&scoped=true":
/*!*********************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=style&index=0&id=3e16625e&lang=less&scoped=true ***!
  \*********************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_UsageIntro_vue_vue_type_style_index_0_id_3e16625e_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../../node_modules/less-loader/dist/cjs.js!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./UsageIntro.vue?vue&type=style&index=0&id=3e16625e&lang=less&scoped=true */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=style&index=0&id=3e16625e&lang=less&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=style&index=0&id=2cb6c8b3&lang=less&scoped=true":
/*!******************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=style&index=0&id=2cb6c8b3&lang=less&scoped=true ***!
  \******************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_style_index_0_id_2cb6c8b3_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./index.vue?vue&type=style&index=0&id=2cb6c8b3&lang=less&scoped=true */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=style&index=0&id=2cb6c8b3&lang=less&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=style&index=0&id=e8b9b4e4&lang=less&scoped=true":
/*!*************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=style&index=0&id=e8b9b4e4&lang=less&scoped=true ***!
  \*************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_CreateForm_vue_vue_type_style_index_0_id_e8b9b4e4_lang_less_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./CreateForm.vue?vue&type=style&index=0&id=e8b9b4e4&lang=less&scoped=true */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=style&index=0&id=e8b9b4e4&lang=less&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=style&index=0&id=70071dcb&scoped=true&lang=less":
/*!*************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=style&index=0&id=70071dcb&scoped=true&lang=less ***!
  \*************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_FileUpload_vue_vue_type_style_index_0_id_70071dcb_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FileUpload.vue?vue&type=style&index=0&id=70071dcb&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=style&index=0&id=70071dcb&scoped=true&lang=less");


/***/ }),

/***/ "./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=style&index=0&id=1a87d32a&scoped=true&lang=less":
/*!***************************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=style&index=0&id=1a87d32a&scoped=true&lang=less ***!
  \***************************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_mini_css_extract_plugin_dist_loader_js_node_modules_css_loader_dist_cjs_js_clonedRuleSet_4_use_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_clonedRuleSet_4_use_2_node_modules_less_loader_dist_cjs_js_node_modules_vue_loader_lib_index_js_vue_loader_options_FolderUploadSelect_vue_vue_type_style_index_0_id_1a87d32a_scoped_true_lang_less__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/mini-css-extract-plugin/dist/loader.js!../../../../../node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!../../../../../node_modules/vue-loader/lib/loaders/stylePostLoader.js!../../../../../node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FolderUploadSelect.vue?vue&type=style&index=0&id=1a87d32a&scoped=true&lang=less */ "./node_modules/mini-css-extract-plugin/dist/loader.js!./node_modules/css-loader/dist/cjs.js??clonedRuleSet-4.use[1]!./node_modules/vue-loader/lib/loaders/stylePostLoader.js!./node_modules/postcss-loader/dist/cjs.js??clonedRuleSet-4.use[2]!./node_modules/less-loader/dist/cjs.js!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=style&index=0&id=1a87d32a&scoped=true&lang=less");


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

/***/ "./web_src/vuepages/components/BaseTitle.vue?vue&type=template&id=94dd892a&scoped=true":
/*!*********************************************************************************************!*\
  !*** ./web_src/vuepages/components/BaseTitle.vue?vue&type=template&id=94dd892a&scoped=true ***!
  \*********************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseTitle_vue_vue_type_template_id_94dd892a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseTitle_vue_vue_type_template_id_94dd892a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_BaseTitle_vue_vue_type_template_id_94dd892a_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./BaseTitle.vue?vue&type=template&id=94dd892a&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseTitle.vue?vue&type=template&id=94dd892a&scoped=true");


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

/***/ "./web_src/vuepages/components/square/CommonSDK.vue?vue&type=template&id=10bd582e&scoped=true":
/*!****************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/CommonSDK.vue?vue&type=template&id=10bd582e&scoped=true ***!
  \****************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_CommonSDK_vue_vue_type_template_id_10bd582e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_CommonSDK_vue_vue_type_template_id_10bd582e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_CommonSDK_vue_vue_type_template_id_10bd582e_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./CommonSDK.vue?vue&type=template&id=10bd582e&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=template&id=10bd582e&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/Header.vue?vue&type=template&id=7b0734e5&scoped=true":
/*!********************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/Header.vue?vue&type=template&id=7b0734e5&scoped=true ***!
  \********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Header_vue_vue_type_template_id_7b0734e5_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Header_vue_vue_type_template_id_7b0734e5_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Header_vue_vue_type_template_id_7b0734e5_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Header.vue?vue&type=template&id=7b0734e5&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Header.vue?vue&type=template&id=7b0734e5&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/Icons.vue?vue&type=template&id=405d1b1c":
/*!*******************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/Icons.vue?vue&type=template&id=405d1b1c ***!
  \*******************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_template_id_405d1b1c__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_template_id_405d1b1c__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Icons_vue_vue_type_template_id_405d1b1c__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Icons.vue?vue&type=template&id=405d1b1c */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Icons.vue?vue&type=template&id=405d1b1c");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=template&id=36f478aa&scoped=true":
/*!****************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=template&id=36f478aa&scoped=true ***!
  \****************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FileList_vue_vue_type_template_id_36f478aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FileList_vue_vue_type_template_id_36f478aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FileList_vue_vue_type_template_id_36f478aa_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FileList.vue?vue&type=template&id=36f478aa&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=template&id=36f478aa&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=template&id=6704b8e1&scoped=true":
/*!*************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=template&id=6704b8e1&scoped=true ***!
  \*************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Intro_vue_vue_type_template_id_6704b8e1_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Intro_vue_vue_type_template_id_6704b8e1_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Intro_vue_vue_type_template_id_6704b8e1_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Intro.vue?vue&type=template&id=6704b8e1&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=template&id=6704b8e1&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=template&id=a96f77d8&scoped=true":
/*!************************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=template&id=a96f77d8&scoped=true ***!
  \************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_MigrateModelSync_vue_vue_type_template_id_a96f77d8_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_MigrateModelSync_vue_vue_type_template_id_a96f77d8_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_MigrateModelSync_vue_vue_type_template_id_a96f77d8_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./MigrateModelSync.vue?vue&type=template&id=a96f77d8&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=template&id=a96f77d8&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=template&id=9b01554a&scoped=true":
/*!****************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=template&id=9b01554a&scoped=true ***!
  \****************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Access_vue_vue_type_template_id_9b01554a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Access_vue_vue_type_template_id_9b01554a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Access_vue_vue_type_template_id_9b01554a_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Access.vue?vue&type=template&id=9b01554a&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=template&id=9b01554a&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=template&id=5368752e&scoped=true":
/*!*****************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=template&id=5368752e&scoped=true ***!
  \*****************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Setting_vue_vue_type_template_id_5368752e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Setting_vue_vue_type_template_id_5368752e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_Setting_vue_vue_type_template_id_5368752e_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./Setting.vue?vue&type=template&id=5368752e&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=template&id=5368752e&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=template&id=3e16625e&scoped=true":
/*!******************************************************************************************************************!*\
  !*** ./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=template&id=3e16625e&scoped=true ***!
  \******************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_UsageIntro_vue_vue_type_template_id_3e16625e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_UsageIntro_vue_vue_type_template_id_3e16625e_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_UsageIntro_vue_vue_type_template_id_3e16625e_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./UsageIntro.vue?vue&type=template&id=3e16625e&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=template&id=3e16625e&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=template&id=2cb6c8b3&scoped=true":
/*!***************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=template&id=2cb6c8b3&scoped=true ***!
  \***************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_template_id_2cb6c8b3_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_template_id_2cb6c8b3_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_index_vue_vue_type_template_id_2cb6c8b3_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./index.vue?vue&type=template&id=2cb6c8b3&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=template&id=2cb6c8b3&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=template&id=e8b9b4e4&scoped=true":
/*!**********************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=template&id=e8b9b4e4&scoped=true ***!
  \**********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_CreateForm_vue_vue_type_template_id_e8b9b4e4_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_CreateForm_vue_vue_type_template_id_e8b9b4e4_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_CreateForm_vue_vue_type_template_id_e8b9b4e4_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./CreateForm.vue?vue&type=template&id=e8b9b4e4&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=template&id=e8b9b4e4&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=template&id=70071dcb&scoped=true":
/*!**********************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=template&id=70071dcb&scoped=true ***!
  \**********************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FileUpload_vue_vue_type_template_id_70071dcb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FileUpload_vue_vue_type_template_id_70071dcb_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FileUpload_vue_vue_type_template_id_70071dcb_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FileUpload.vue?vue&type=template&id=70071dcb&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=template&id=70071dcb&scoped=true");


/***/ }),

/***/ "./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=template&id=1a87d32a&scoped=true":
/*!************************************************************************************************************************!*\
  !*** ./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=template&id=1a87d32a&scoped=true ***!
  \************************************************************************************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   render: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FolderUploadSelect_vue_vue_type_template_id_1a87d32a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.render; },
/* harmony export */   staticRenderFns: function() { return /* reexport safe */ _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FolderUploadSelect_vue_vue_type_template_id_1a87d32a_scoped_true__WEBPACK_IMPORTED_MODULE_0__.staticRenderFns; }
/* harmony export */ });
/* harmony import */ var _node_modules_vue_loader_lib_loaders_templateLoader_js_vue_loader_options_node_modules_vue_loader_lib_index_js_vue_loader_options_FolderUploadSelect_vue_vue_type_template_id_1a87d32a_scoped_true__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! -!../../../../../node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!../../../../../node_modules/vue-loader/lib/index.js??vue-loader-options!./FolderUploadSelect.vue?vue&type=template&id=1a87d32a&scoped=true */ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=template&id=1a87d32a&scoped=true");


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

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseTitle.vue?vue&type=template&id=94dd892a&scoped=true":
/*!************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/BaseTitle.vue?vue&type=template&id=94dd892a&scoped=true ***!
  \************************************************************************************************************************************************************************************************************************************/
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
  return _c("div", { staticClass: "form-container" }, [
    _c("div", { staticClass: "form-head" }, [
      _c("h4", [_vm._v(_vm._s(_vm.title))])
    ]),
    _vm._v(" "),
    _c("div", { staticClass: "form-content" }, [_vm._t("default")], 2)
  ])
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

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=template&id=10bd582e&scoped=true":
/*!*******************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/CommonSDK.vue?vue&type=template&id=10bd582e&scoped=true ***!
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
    { staticClass: "common-tips-dlg" },
    [
      _c(
        "div",
        { staticClass: "trigger-container", on: { click: _vm.handlerOpen } },
        [_vm._t("default")],
        2
      ),
      _vm._v(" "),
      _c(
        "BaseDialog",
        {
          attrs: {
            visible: _vm.visible,
            title: _vm.title,
            "show-close": _vm.showClose,
            appendToBody: _vm.appendToBody
          },
          on: { closed: _vm.closeDialog }
        },
        [
          _c("el-skeleton", {
            staticStyle: { padding: "0 20px" },
            attrs: { rows: 10, loading: _vm.loading, animated: "" }
          }),
          _vm._v(" "),
          _c("div", {
            staticClass: "markdown",
            domProps: { innerHTML: _vm._s(_vm.content) }
          }),
          _vm._v(" "),
          _c(
            "div",
            {
              staticClass: "dialog-footer",
              attrs: { slot: "footer" },
              slot: "footer"
            },
            [
              _c(
                "el-button",
                {
                  staticClass: "close-btn",
                  attrs: { size: "default", type: "success" },
                  on: { click: _vm.closeDialog }
                },
                [_vm._v("\n      " + _vm._s(_vm.closeText) + "\n      ")]
              )
            ],
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

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Header.vue?vue&type=template&id=7b0734e5&scoped=true":
/*!***********************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Header.vue?vue&type=template&id=7b0734e5&scoped=true ***!
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
  return _c("div", { staticClass: "bg" }, [
    _c("div", { staticClass: "ui container" }, [
      _c("div", { staticClass: "title" }, [
        _c(
          "div",
          {
            staticClass: "title-l",
            style: { "max-width": _vm.type === "dataset" ? "74%" : "69%" },
            on: {
              mouseover: function($event) {
                _vm.hover = true
              },
              mouseleave: function($event) {
                _vm.hover = false
              }
            }
          },
          [
            _c("Icons", {
              staticStyle: { "flex-shrink": "0" },
              attrs: {
                type: _vm.type,
                isPrivate: _vm.dataObj.is_private,
                modelType: _vm.dataObj.aimodel_type
              }
            }),
            _vm._v(" "),
            _c("div", { staticClass: "t-name nowrap" }, [
              _c(
                "span",
                {
                  staticClass: "t-name-data",
                  attrs: {
                    title:
                      _vm.lang == "en-US" ? _vm.dataObj.name : _vm.dataObj.alias
                  }
                },
                [
                  _vm._v(
                    _vm._s(
                      _vm.lang == "en-US" ? _vm.dataObj.name : _vm.dataObj.alias
                    )
                  )
                ]
              )
            ]),
            _vm._v(" "),
            _vm.dataObj.recommend
              ? _c("div", { staticClass: "reconmend-icon" }, [
                  _c(
                    "svg",
                    {
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
                            fill: "#FF6200",
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
        _c(
          "div",
          {
            staticClass: "title-r",
            style: { "min-width": _vm.type === "dataset" ? "340px" : "440px" }
          },
          [
            _c("div", { staticClass: "btn-group" }, [
              _c(
                "div",
                {
                  staticClass: "base-btn like-btn",
                  staticStyle: { cursor: "pointer" },
                  on: {
                    click: function($event) {
                      return _vm.changeFav(_vm.dataObj)
                    }
                  }
                },
                [
                  !_vm.isCollected
                    ? [
                        _c("i", {
                          staticClass: "heart outline icon",
                          attrs: { title: _vm.$t("star") }
                        }),
                        _vm._v(" "),
                        _c("span", [_vm._v(_vm._s(_vm.$t("star")))])
                      ]
                    : [
                        _c("i", {
                          staticClass: "heart icon",
                          attrs: { title: _vm.$t("unStar") }
                        }),
                        _vm._v(" "),
                        _c("span", [_vm._v(_vm._s(_vm.$t("unStar")))])
                      ]
                ],
                2
              ),
              _vm._v(" "),
              _c("div", { staticClass: "base-btn" }, [
                _vm._v(_vm._s(_vm.collected_count))
              ])
            ]),
            _vm._v(" "),
            _c("div", { staticClass: "btn-group" }, [
              _c("div", { staticClass: "base-btn" }, [
                _c("i", { staticClass: "el-icon-link" }),
                _vm._v(" "),
                _c("span", [_vm._v(_vm._s(_vm.$t("datasets.citations1")))])
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "base-btn" }, [
                _vm._v(_vm._s(_vm.dataObj.use_count))
              ])
            ]),
            _vm._v(" "),
            _c("div", { staticClass: "btn-group" }, [
              _c("div", { staticClass: "base-btn" }, [
                _c("i", { staticClass: "el-icon-download" }),
                _vm._v(" "),
                _c("span", [_vm._v(_vm._s(_vm.$t("datasets.downloadtimes1")))])
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "base-btn" }, [
                _vm._v(_vm._s(_vm.dataObj.download_count))
              ])
            ]),
            _vm._v(" "),
            _vm.type !== "dataset"
              ? _c("div", { staticClass: "btn-group" }, [
                  _c("div", { staticClass: "base-btn" }, [
                    _c("i", { staticClass: "ri-git-merge-line" }),
                    _vm._v(" "),
                    _c("span", [
                      _vm._v(_vm._s(_vm.$t("modelManage.derivativeTimes")))
                    ])
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "base-btn" }, [
                    _vm._v(_vm._s(_vm.dataObj.derivative_count))
                  ])
                ])
              : _vm._e()
          ]
        )
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "sub-data-title" }, [
        _c("div", { staticClass: "nowrap" }, [
          _c(
            "a",
            {
              staticStyle: { color: "rgba(3, 102, 214, 1)" },
              attrs: { href: "/" + _vm.dataObj.owner_name }
            },
            [_vm._v(_vm._s(_vm.dataObj.owner_name))]
          ),
          _vm._v(" "),
          _c("span", [_vm._v(" / ")]),
          _vm._v(" "),
          _c("span", { staticStyle: { color: "rgba(16, 16, 16, 0.8)" } }, [
            _vm._v(_vm._s(_vm.dataObj.name))
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
                id: "clipboard-btn",
                "data-position": "top center",
                "data-variation": "inverted tiny",
                "data-success": _vm.$t("copySuccess"),
                "data-content": _vm.$t("copy"),
                "data-original": _vm.$t("copy"),
                "data-clipboard-text":
                  _vm.dataObj.owner_name + "/" + _vm.dataObj.name
              }
            },
            [_c("i", { staticClass: "ri-file-copy-line" })]
          )
        ])
      ]),
      _vm._v(" "),
      _vm.dataObj.aimodel_type !== 2
        ? _c("div", { staticClass: "sub-title" }, [
            _c(
              "div",
              { staticClass: "labels" },
              [
                _vm._l(_vm.dataObj.tags, function(item, index) {
                  return _c("span", { key: item, staticClass: "label" }, [
                    _vm._v(_vm._s(_vm.$t("datasets." + item)))
                  ])
                }),
                _vm._v(" "),
                _vm._l(_vm.dataObj.tasks, function(item, index) {
                  return _c("span", { key: item, staticClass: "label" }, [
                    _vm._v(_vm._s(_vm.$t("datasets." + item)))
                  ])
                }),
                _vm._v(" "),
                _vm.dataObj.licenses
                  ? _c("span", { staticClass: "label" }, [
                      _vm._v(" " + _vm._s(_vm.dataObj.licenses) + " ")
                    ])
                  : _vm._e(),
                _vm._v(" "),
                _vm.dataObj.engineName || _vm.dataObj.engine
                  ? _c("span", { staticClass: "label" }, [
                      _vm._v(
                        " " +
                          _vm._s(_vm.dataObj.engineName || _vm.dataObj.engine)
                      )
                    ])
                  : _vm._e(),
                _vm._v(" "),
                _vm._l(_vm.dataObj.labels, function(item, index) {
                  return _c(
                    "span",
                    {
                      key: index,
                      staticClass: "label nowrap",
                      attrs: { title: item }
                    },
                    [_vm._v(_vm._s(item))]
                  )
                })
              ],
              2
            )
          ])
        : _c("div", { staticClass: "sub-title" }, [
            _c("div", { staticClass: "labels" }, [
              _c("span", { staticClass: "label nowrap" }, [
                _vm._v(
                  _vm._s(_vm.$t("modelObj.model_source")) +
                    _vm._s(_vm.dataObj.external_name)
                )
              ])
            ])
          ]),
      _vm._v(" "),
      _c("div", { staticClass: "tabs-wrap" }, [
        _c("div", { staticClass: "tabs" }, [
          _c(
            "div",
            { staticClass: "tabs-l" },
            _vm._l(_vm.tabList, function(item) {
              return _c(
                "div",
                {
                  key: item.key,
                  staticClass: "tab",
                  class: _vm.tabIndex == item.key ? "focus" : "",
                  on: {
                    click: function($event) {
                      return _vm.changeTab(item)
                    }
                  }
                },
                [
                  _c("i", { class: item.icon }),
                  _vm._v(" "),
                  _c("span", [_vm._v(_vm._s(_vm.$t(item.name)))])
                ]
              )
            }),
            0
          ),
          _vm._v(" "),
          _c("div", { staticClass: "tabs-r" }, [
            _vm.dataObj && _vm.dataObj.can_manage
              ? _c(
                  "div",
                  {
                    staticClass: "tab",
                    class: _vm.tabIndex == "settings" ? "focus" : "",
                    on: {
                      click: function($event) {
                        return _vm.changeTab({ key: "settings" })
                      }
                    }
                  },
                  [
                    _c("i", { staticClass: "ri-settings-2-line" }),
                    _c("span", [_vm._v(_vm._s(_vm.$t("modelManage.settings")))])
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

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Icons.vue?vue&type=template&id=405d1b1c":
/*!**********************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/Icons.vue?vue&type=template&id=405d1b1c ***!
  \**********************************************************************************************************************************************************************************************************************************/
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
                fill: "#101010",
                "p-id": "3601"
              }
            }),
            _c("path", {
              attrs: {
                d:
                  "M768 409.6c-145.92 0-256 55.04-256 128v358.4c0 72.96 110.08 128 256 128s256-55.04 256-128v-358.4c0-72.96-110.08-128-256-128z m0 324.2496c-125.0304 0-204.8-45.4656-204.8-76.8v-40.96c46.08 30.4128 119.1936 49.3568 204.8 49.3568s158.72-18.944 204.8-49.3568v40.96c0 31.3344-79.7696 76.8-204.8 76.8z m-204.8 1.8944c46.08 30.3616 119.1936 49.3056 204.8 49.3056s158.72-18.944 204.8-49.3056v40.96c0 31.2832-79.7696 76.8-204.8 76.8s-204.8-45.5168-204.8-76.8zM768 460.8c125.0304 0 204.8 45.4656 204.8 76.8s-79.7696 76.8-204.8 76.8-204.8-45.4656-204.8-76.8 79.7696-76.8 204.8-76.8z m0 512c-125.0304 0-204.8-45.4656-204.8-76.8v-40.96c46.08 30.4128 119.1936 49.3568 204.8 49.3568s158.72-18.944 204.8-49.3568v40.96c0 31.3344-79.7696 76.8-204.8 76.8z",
                fill: "#101010",
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
                  fill: "#101010",
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
                  fill: "#101010",
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
                  fill: "#101010",
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
            _c("defs", [
              _c(
                "clipPath",
                { attrs: { id: "clip248_15-imb94pib95nx105-UmZeied51Isooy" } },
                [
                  _c("rect", {
                    attrs: {
                      width: "200.000000",
                      height: "200.000000",
                      fill: "white",
                      "fill-opacity": "0"
                    }
                  })
                ]
              )
            ]),
            _vm._v(" "),
            _c(
              "g",
              {
                attrs: {
                  "clip-path": "url(#clip248_15-imb94pib95nx105-UmZeied51Isooy)"
                }
              },
              [
                _c("path", {
                  attrs: {
                    id: "path-imb94pib95nx105-UmZeied51Isooy",
                    d:
                      "M102 139C101.05 139 100.1 138.75 99.25 138.26L66.74 119.42C66.32 119.18 65.94 118.89 65.6 118.55C65.26 118.21 64.97 117.83 64.73 117.41C64.49 116.99 64.31 116.55 64.18 116.09C64.06 115.62 64 115.15 64 114.66L64 77C64 75.03 65.04 73.22 66.74 72.24L74.1 67.97C76.72 66.45 80.08 67.35 81.59 69.98C83.1 72.61 82.2 75.97 79.58 77.49L74.96 80.17L74.96 111.49L102 127.15L129.03 111.49L129.03 80.17L124 77.25C121.38 75.74 120.48 72.37 121.99 69.74C123.51 67.11 126.86 66.21 129.48 67.73L137.25 72.23C137.67 72.48 138.05 72.77 138.39 73.11C138.73 73.45 139.02 73.83 139.26 74.25C139.5 74.66 139.68 75.11 139.81 75.57C139.93 76.04 140 76.51 140 77L140 114.66C140 116.63 138.95 118.44 137.25 119.42L104.74 138.26C103.89 138.75 102.97 139 102 139Z",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    id: "path-imb94pib95nx105-UmZeied51Isooy",
                    d:
                      "M102 53C102.94 53 103.89 53.24 104.74 53.73L137.25 72.57C137.67 72.81 138.05 73.1 138.39 73.44C138.73 73.78 139.02 74.16 139.26 74.58C139.5 75 139.68 75.44 139.81 75.9C139.93 76.37 140 76.84 140 77.33L140 114.99C140 116.96 138.95 118.77 137.25 119.75L129.89 124.02C127.27 125.54 123.91 124.64 122.4 122.01C120.89 119.38 121.79 116.02 124.41 114.5L129.03 111.82L129.03 80.5L102 64.84L74.96 80.5L74.96 111.82L79.99 114.74C82.61 116.25 83.51 119.62 82 122.25C80.48 124.88 77.13 125.78 74.51 124.26L66.74 119.76C66.32 119.52 65.94 119.22 65.6 118.88C65.26 118.54 64.97 118.16 64.73 117.74C64.49 117.33 64.31 116.88 64.18 116.42C64.06 115.95 64 115.48 64 115L64 77.33C64 75.36 65.04 73.55 66.74 72.57L99.25 53.73C100.1 53.24 101.02 53 102 53Z",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    id: "path-imb94pib95nx105-UmZeied51Isooy",
                    d:
                      "M101.5 200C100.55 200 99.61 199.75 98.77 199.27L62.72 178.52C62.31 178.28 61.93 177.99 61.59 177.65C61.26 177.31 60.97 176.94 60.73 176.52C60.49 176.11 60.31 175.67 60.18 175.21C60.06 174.75 60 174.28 60 173.8L60 132.3C60 130.36 61.04 128.56 62.72 127.59L79.86 117.73C82.47 116.22 85.81 117.11 87.31 119.72C88.82 122.32 87.93 125.65 85.32 127.15L70.91 135.45L70.91 170.66L101.5 188.27L132.08 170.66L132.08 135.45L117.93 127.3C115.32 125.8 114.42 122.47 115.93 119.86C117.44 117.26 120.78 116.37 123.39 117.87L140.27 127.59C140.68 127.83 141.06 128.11 141.4 128.45C141.74 128.79 142.02 129.17 142.26 129.58C142.5 129.99 142.69 130.43 142.81 130.89C142.93 131.35 143 131.82 143 132.3L143 173.8C143 175.75 141.96 177.54 140.27 178.52L104.23 199.27C103.38 199.75 102.47 200 101.5 200Z",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    id: "path-imb94pib95nx105-UmZeied51Isooy",
                    d:
                      "M50.83 173C49.88 173 48.94 172.75 48.09 172.26L11.73 151.25C11.31 151 10.94 150.72 10.6 150.38C10.26 150.03 9.97 149.66 9.73 149.24C9.49 148.82 9.31 148.38 9.18 147.92C9.06 147.46 9 146.98 9 146.5L9 104.47C9 101.45 11.44 99 14.46 99C17.49 99 19.93 101.45 19.93 104.47L19.93 143.34L50.83 161.2L81.72 143.34L81.72 125.76C81.72 122.74 84.17 120.29 87.19 120.29C90.21 120.29 92.66 122.74 92.66 125.76L92.66 146.5C92.66 148.46 91.62 150.27 89.92 151.25L53.56 172.26C52.72 172.75 51.8 173 50.83 173ZM151.16 173C150.22 173 149.28 172.75 148.43 172.26L112.07 151.25C111.65 151 111.28 150.72 110.94 150.38C110.6 150.03 110.31 149.66 110.07 149.24C109.83 148.83 109.64 148.39 109.52 147.92C109.4 147.46 109.33 146.98 109.33 146.5L109.33 126.07C109.33 123.04 111.78 120.59 114.8 120.59C117.83 120.59 120.27 123.04 120.27 126.07L120.27 143.34L151.16 161.2L182.06 143.34L182.06 104.47C182.06 101.45 184.5 99 187.53 99C190.55 99 193 101.45 193 104.47L193 146.5C193 148.46 191.95 150.27 190.26 151.25L153.9 172.26C153.05 172.75 152.14 173 151.16 173Z",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    id: "path-imb94pib95nx105-UmZeied51Isooy",
                    d:
                      "M155.14 112C154.19 112 153.25 111.75 152.41 111.26L131.52 99.16C128.91 97.64 128.02 94.3 129.52 91.67C131.03 89.05 134.37 88.16 136.98 89.67L155.14 100.19L182.08 84.58L182.08 53.35L155.14 37.73L127.77 53.59C126.95 55.57 125 56.96 122.73 56.96C119.72 56.96 117.28 54.51 117.28 51.49L117.28 50.19C117.28 48.23 118.32 46.42 120.01 45.44L152.41 26.67C152.82 26.42 153.26 26.24 153.72 26.12C154.19 25.99 154.66 25.93 155.14 25.93C155.61 25.93 156.09 25.99 156.55 26.12C157.01 26.24 157.45 26.42 157.86 26.67L190.27 45.44C190.68 45.68 191.06 45.97 191.4 46.31C191.74 46.65 192.03 47.03 192.26 47.45C192.5 47.86 192.69 48.3 192.81 48.77C192.93 49.23 193 49.71 193 50.19L193 87.74C193 89.7 191.96 91.5 190.27 92.48L157.86 111.26C157.02 111.75 156.11 112 155.14 112ZM46.85 112C45.91 112 44.97 111.75 44.13 111.26L11.72 92.48C11.31 92.24 10.93 91.95 10.59 91.61C10.26 91.27 9.97 90.89 9.73 90.48C9.49 90.06 9.31 89.62 9.18 89.16C9.06 88.69 9 88.22 9 87.74L9 50.19C9 48.23 10.04 46.42 11.72 45.44L44.13 26.67C44.54 26.42 44.98 26.24 45.44 26.12C45.9 25.99 46.38 25.93 46.85 25.93C47.33 25.93 47.8 25.99 48.27 26.12C48.73 26.24 49.17 26.42 49.58 26.67L81.99 45.44C82.4 45.68 82.78 45.97 83.12 46.31C83.45 46.65 83.74 47.03 83.98 47.45C84.22 47.86 84.4 48.3 84.53 48.77C84.65 49.23 84.71 49.71 84.71 50.19L84.71 52.01C84.71 55.04 82.27 57.49 79.26 57.49C76.74 57.49 74.63 55.78 73.99 53.46L46.86 37.74L19.91 53.35L19.91 84.58L46.86 100.2L66.35 88.9C68.96 87.39 72.3 88.28 73.8 90.9C75.31 93.52 74.42 96.87 71.81 98.39L49.59 111.26C48.74 111.75 47.83 112 46.85 112ZM69.07 33.84C66.06 33.84 63.62 31.39 63.62 28.36L63.62 24.25C63.62 22.29 64.66 20.49 66.35 19.51L98.75 0.73C99.16 0.49 99.6 0.31 100.07 0.18C100.53 0.06 101 0 101.48 0C101.96 0 102.43 0.06 102.89 0.18C103.35 0.31 103.79 0.49 104.21 0.73L136.61 19.51C137.02 19.75 137.4 20.04 137.74 20.38C138.08 20.72 138.37 21.1 138.61 21.51C138.85 21.93 139.03 22.37 139.15 22.83C139.28 23.3 139.34 23.77 139.34 24.25L139.34 28.19C139.34 31.22 136.89 33.67 133.88 33.67C130.87 33.67 128.42 31.22 128.42 28.19L128.42 27.41L101.48 11.8L74.53 27.41L74.53 28.37C74.53 31.39 72.09 33.84 69.07 33.84Z",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                })
              ]
            )
          ])
        ]
      )
    : _vm.type != "dataset" && _vm.isPrivate
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
          _c("g", [
            _c("desc", [_vm._v("\n        Created with Pixso.\n    ")]),
            _vm._v(" "),
            _c("defs", [
              _c(
                "clipPath",
                {
                  attrs: {
                    id: "clip255_38-imb94pibaiq0q7q-UmZeiedeUmZepo268ClcGM"
                  }
                },
                [
                  _c("rect", {
                    attrs: {
                      width: "200.000000",
                      height: "200.000000",
                      fill: "white",
                      "fill-opacity": "0"
                    }
                  })
                ]
              )
            ]),
            _vm._v(" "),
            _c(
              "g",
              {
                attrs: {
                  "clip-path":
                    "url(#clip255_38-imb94pibaiq0q7q-UmZeiedeUmZepo268ClcGM)"
                }
              },
              [
                _c("path", {
                  attrs: {
                    d:
                      "M102 139C101.05 139 100.1 138.75 99.25 138.26L66.74 119.42C66.32 119.18 65.94 118.89 65.6 118.55C65.26 118.21 64.97 117.83 64.73 117.41C64.49 116.99 64.31 116.55 64.18 116.09C64.06 115.62 64 115.15 64 114.66L64 77C64 75.03 65.04 73.22 66.74 72.24L74.1 67.97C76.72 66.45 80.08 67.35 81.59 69.98C83.1 72.61 82.2 75.97 79.58 77.49L74.96 80.17L74.96 111.49L102 127.15L129.03 111.49L129.03 80.17L124 77.25C121.38 75.74 120.48 72.37 121.99 69.74C123.51 67.11 126.86 66.21 129.48 67.73L137.25 72.23C137.67 72.48 138.05 72.77 138.39 73.11C138.73 73.45 139.02 73.83 139.26 74.25C139.5 74.66 139.68 75.11 139.81 75.57C139.93 76.04 140 76.51 140 77L140 114.66C140 116.63 138.95 118.44 137.25 119.42L104.74 138.26C103.89 138.75 102.97 139 102 139Z",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    d:
                      "M102 53C102.94 53 103.89 53.24 104.74 53.73L137.25 72.57C137.67 72.81 138.05 73.1 138.39 73.44C138.73 73.78 139.02 74.16 139.26 74.58C139.5 75 139.68 75.44 139.81 75.9C139.93 76.37 140 76.84 140 77.33L140 114.99C140 116.96 138.95 118.77 137.25 119.75L129.89 124.02C127.27 125.54 123.91 124.64 122.4 122.01C120.89 119.38 121.79 116.02 124.41 114.5L129.03 111.82L129.03 80.5L102 64.84L74.96 80.5L74.96 111.82L79.99 114.74C82.61 116.25 83.51 119.62 82 122.25C80.48 124.88 77.13 125.78 74.51 124.26L66.74 119.76C66.32 119.52 65.94 119.22 65.6 118.88C65.26 118.54 64.97 118.16 64.73 117.74C64.49 117.33 64.31 116.88 64.18 116.42C64.06 115.95 64 115.48 64 115L64 77.33C64 75.36 65.04 73.55 66.74 72.57L99.25 53.73C100.1 53.24 101.02 53 102 53Z",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    d:
                      "M101.5 200C100.55 200 99.61 199.75 98.77 199.27L62.72 178.52C62.31 178.28 61.93 177.99 61.59 177.65C61.26 177.31 60.97 176.94 60.73 176.52C60.49 176.11 60.31 175.67 60.18 175.21C60.06 174.75 60 174.28 60 173.8L60 132.3C60 130.36 61.04 128.56 62.72 127.59L79.86 117.73C82.47 116.22 85.81 117.11 87.31 119.72C88.82 122.32 87.93 125.65 85.32 127.15L70.91 135.45L70.91 170.66L101.5 188.27L132.08 170.66L132.08 135.45L117.93 127.3C115.32 125.8 114.42 122.47 115.93 119.86C117.44 117.26 120.78 116.37 123.39 117.87L140.27 127.59C140.68 127.83 141.06 128.11 141.4 128.45C141.74 128.79 142.02 129.17 142.26 129.58C142.5 129.99 142.69 130.43 142.81 130.89C142.93 131.35 143 131.82 143 132.3L143 173.8C143 175.75 141.96 177.54 140.27 178.52L104.23 199.27C103.38 199.75 102.47 200 101.5 200Z",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    d:
                      "M50.83 173C49.88 173 48.94 172.75 48.09 172.26L11.73 151.25C11.31 151 10.94 150.72 10.6 150.38C10.26 150.03 9.97 149.66 9.73 149.24C9.49 148.82 9.31 148.38 9.18 147.92C9.06 147.46 9 146.98 9 146.5L9 104.47C9 101.45 11.44 99 14.46 99C17.49 99 19.93 101.45 19.93 104.47L19.93 143.34L50.83 161.2L81.72 143.34L81.72 125.76C81.72 122.74 84.17 120.29 87.19 120.29C90.21 120.29 92.66 122.74 92.66 125.76L92.66 146.5C92.66 148.46 91.62 150.27 89.92 151.25L53.56 172.26C52.72 172.75 51.8 173 50.83 173ZM151.16 173C150.22 173 149.28 172.75 148.43 172.26L112.07 151.25C111.65 151 111.28 150.72 110.94 150.38C110.6 150.03 110.31 149.66 110.07 149.24C109.83 148.83 109.64 148.39 109.52 147.92C109.4 147.46 109.33 146.98 109.33 146.5L109.33 126.07C109.33 123.04 111.78 120.59 114.8 120.59C117.83 120.59 120.27 123.04 120.27 126.07L120.27 143.34L151.16 161.2L182.06 143.34L182.06 104.47C182.06 101.45 184.5 99 187.53 99C190.55 99 193 101.45 193 104.47L193 146.5C193 148.46 191.95 150.27 190.26 151.25L153.9 172.26C153.05 172.75 152.14 173 151.16 173Z",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    d:
                      "M155.14 112C154.19 112 153.25 111.75 152.41 111.26L131.52 99.16C128.91 97.64 128.02 94.3 129.52 91.67C131.03 89.05 134.37 88.16 136.98 89.67L155.14 100.19L182.08 84.58L182.08 53.35L155.14 37.73L127.77 53.59C126.95 55.57 125 56.96 122.73 56.96C119.72 56.96 117.28 54.51 117.28 51.49L117.28 50.19C117.28 48.23 118.32 46.42 120.01 45.44L152.41 26.67C152.82 26.42 153.26 26.24 153.72 26.12C154.19 25.99 154.66 25.93 155.14 25.93C155.61 25.93 156.09 25.99 156.55 26.12C157.01 26.24 157.45 26.42 157.86 26.67L190.27 45.44C190.68 45.68 191.06 45.97 191.4 46.31C191.74 46.65 192.03 47.03 192.26 47.45C192.5 47.86 192.69 48.3 192.81 48.77C192.93 49.23 193 49.71 193 50.19L193 87.74C193 89.7 191.96 91.5 190.27 92.48L157.86 111.26C157.02 111.75 156.11 112 155.14 112ZM46.85 112C45.91 112 44.97 111.75 44.13 111.26L11.72 92.48C11.31 92.24 10.93 91.95 10.59 91.61C10.26 91.27 9.97 90.89 9.73 90.48C9.49 90.06 9.31 89.62 9.18 89.16C9.06 88.69 9 88.22 9 87.74L9 50.19C9 48.23 10.04 46.42 11.72 45.44L44.13 26.67C44.54 26.42 44.98 26.24 45.44 26.12C45.9 25.99 46.38 25.93 46.85 25.93C47.33 25.93 47.8 25.99 48.27 26.12C48.73 26.24 49.17 26.42 49.58 26.67L81.99 45.44C82.4 45.68 82.78 45.97 83.12 46.31C83.45 46.65 83.74 47.03 83.98 47.45C84.22 47.86 84.4 48.3 84.53 48.77C84.65 49.23 84.71 49.71 84.71 50.19L84.71 52.01C84.71 55.04 82.27 57.49 79.26 57.49C76.74 57.49 74.63 55.78 73.99 53.46L46.86 37.74L19.91 53.35L19.91 84.58L46.86 100.2L66.35 88.9C68.96 87.39 72.3 88.28 73.8 90.9C75.31 93.52 74.42 96.87 71.81 98.39L49.59 111.26C48.74 111.75 47.83 112 46.85 112ZM69.07 33.84C66.06 33.84 63.62 31.39 63.62 28.36L63.62 24.25C63.62 22.29 64.66 20.49 66.35 19.51L98.75 0.73C99.16 0.49 99.6 0.31 100.07 0.18C100.53 0.06 101 0 101.48 0C101.96 0 102.43 0.06 102.89 0.18C103.35 0.31 103.79 0.49 104.21 0.73L136.61 19.51C137.02 19.75 137.4 20.04 137.74 20.38C138.08 20.72 138.37 21.1 138.61 21.51C138.85 21.93 139.03 22.37 139.15 22.83C139.28 23.3 139.34 23.77 139.34 24.25L139.34 28.19C139.34 31.22 136.89 33.67 133.88 33.67C130.87 33.67 128.42 31.22 128.42 28.19L128.42 27.41L101.48 11.8L74.53 27.41L74.53 28.37C74.53 31.39 72.09 33.84 69.07 33.84Z",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    d:
                      "M97 31.88L46.37 61.11C43.28 62.89 41.37 66.2 41.37 69.77L41.37 128.22C41.37 131.79 43.28 135.1 46.37 136.88L97 166.11C100.09 167.9 103.9 167.9 107 166.11L157.62 136.88C160.71 135.1 162.62 131.79 162.62 128.22L162.62 69.77C162.62 66.2 160.71 62.89 157.62 61.11L107 31.88C103.9 30.1 100.09 30.1 97 31.88Z",
                    fill: "#FFFFFF",
                    "fill-opacity": "1.000000",
                    "fill-rule": "evenodd"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    d:
                      "M46.37 61.11C43.28 62.89 41.37 66.2 41.37 69.77L41.37 128.22C41.37 131.79 43.28 135.1 46.37 136.88L97 166.11C100.09 167.9 103.9 167.9 107 166.11L157.62 136.88C160.71 135.1 162.62 131.79 162.62 128.22L162.62 69.77C162.62 66.2 160.71 62.89 157.62 61.11L107 31.88C103.9 30.1 100.09 30.1 97 31.88L46.37 61.11ZM51.37 69.77L51.37 128.22L102 157.45L152.62 128.22L152.62 69.77L102 40.54L51.37 69.77Z",
                    fill: "#000000",
                    "fill-opacity": "1.000000",
                    "fill-rule": "evenodd"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    d:
                      "M84 73L84 89C84 91.48 84.43 93.78 85.31 95.9C86.19 98.03 87.51 99.97 89.27 101.72C91.02 103.48 92.96 104.8 95.09 105.68C97.21 106.56 99.51 107 102 107C104.48 107 106.78 106.56 108.91 105.68C111.03 104.8 112.97 103.48 114.72 101.72C116.48 99.97 117.8 98.03 118.68 95.9C119.56 93.78 120 91.48 120 89L120 73C120 70.51 119.56 68.21 118.68 66.09C117.8 63.96 116.48 62.02 114.72 60.27C112.97 58.51 111.03 57.19 108.91 56.31C106.78 55.43 104.48 55 102 55C99.51 55 97.21 55.43 95.09 56.31C92.96 57.19 91.02 58.51 89.27 60.27C87.51 62.02 86.19 63.96 85.31 66.09C84.43 68.21 84 70.51 84 73ZM92 73C92 67.47 96.47 63 102 63C107.52 63 112 67.47 112 73L112 89C112 94.52 107.52 99 102 99C96.47 99 92 94.52 92 89L92 73Z",
                    fill: "#000000",
                    "fill-opacity": "1.000000",
                    "fill-rule": "evenodd"
                  }
                }),
                _vm._v(" "),
                _c("rect", {
                  attrs: {
                    x: "79.000000",
                    y: "89.000000",
                    rx: "5.000000",
                    width: "46.000000",
                    height: "36.000000",
                    fill: "#101010",
                    "fill-opacity": "1.000000"
                  }
                }),
                _vm._v(" "),
                _c("rect", {
                  attrs: {
                    x: "75.000000",
                    y: "85.000000",
                    rx: "5.000000",
                    width: "54.000000",
                    height: "44.000000",
                    stroke: "#000000",
                    "stroke-opacity": "1.000000",
                    "stroke-width": "8.000000"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    d:
                      "M107.6 101.3C107.6 101.04 107.58 100.79 107.55 100.53C107.52 100.28 107.47 100.03 107.41 99.78C107.35 99.53 107.27 99.28 107.18 99.04C107.09 98.8 106.98 98.57 106.86 98.34C106.74 98.11 106.61 97.9 106.46 97.69C106.31 97.47 106.15 97.27 105.98 97.08C105.81 96.89 105.63 96.71 105.43 96.54C105.24 96.37 105.03 96.22 104.82 96.07C104.61 95.93 104.39 95.8 104.16 95.68C103.93 95.57 103.7 95.46 103.45 95.38C103.21 95.29 102.97 95.22 102.72 95.16C102.47 95.1 102.21 95.06 101.96 95.03C101.7 95 101.45 94.99 101.19 95C100.93 95 100.68 95.02 100.42 95.06C100.17 95.09 99.92 95.14 99.67 95.21C99.42 95.28 99.18 95.36 98.94 95.45C98.7 95.55 98.47 95.66 98.24 95.78C98.02 95.91 97.8 96.05 97.59 96.2C97.39 96.35 97.19 96.51 97 96.69C96.81 96.86 96.64 97.05 96.47 97.24C96.31 97.44 96.15 97.65 96.01 97.86C95.87 98.08 95.75 98.3 95.63 98.53C95.52 98.76 95.42 99 95.34 99.24C95.26 99.48 95.19 99.73 95.13 99.98C95.08 100.24 95.04 100.49 95.02 100.74C95 101 94.99 101.26 95 101.51C95.01 101.77 95.03 102.03 95.07 102.28C95.11 102.53 95.17 102.78 95.24 103.03C95.31 103.28 95.39 103.52 95.5 103.76C95.6 103.99 95.71 104.22 95.84 104.44C95.97 104.67 96.11 104.88 96.26 105.09C96.42 105.29 96.58 105.49 96.76 105.67C97.05 105.97 97.31 106.29 97.54 106.64C97.77 106.99 97.96 107.36 98.12 107.75C98.28 108.13 98.39 108.53 98.47 108.94C98.55 109.35 98.59 109.76 98.6 110.18L98.6 115.6C98.59 115.77 98.6 115.93 98.63 116.09C98.66 116.26 98.7 116.42 98.76 116.57C98.81 116.73 98.88 116.88 98.96 117.02C99.05 117.16 99.14 117.3 99.25 117.42C99.35 117.55 99.47 117.67 99.6 117.77C99.73 117.88 99.86 117.97 100.01 118.05C100.15 118.13 100.3 118.2 100.46 118.26C100.61 118.31 100.77 118.35 100.94 118.38C101.13 118.4 101.32 118.41 101.51 118.39C101.7 118.38 101.89 118.34 102.07 118.29C102.26 118.23 102.43 118.16 102.6 118.06C102.77 117.97 102.93 117.86 103.07 117.73C103.22 117.61 103.35 117.47 103.46 117.31C103.58 117.16 103.68 116.99 103.76 116.82C103.84 116.64 103.9 116.46 103.94 116.27C103.98 116.08 104 115.89 104 115.7L104 110.18C104 109.77 104.04 109.36 104.11 108.95C104.19 108.54 104.31 108.14 104.46 107.76C104.62 107.37 104.8 107.01 105.03 106.66C105.26 106.31 105.51 105.98 105.8 105.69C106.08 105.4 106.34 105.08 106.56 104.74C106.78 104.4 106.97 104.05 107.13 103.67C107.28 103.29 107.4 102.91 107.48 102.51C107.56 102.11 107.59 101.71 107.6 101.3Z",
                    fill: "#FFFFFF",
                    "fill-opacity": "1.000000",
                    "fill-rule": "nonzero"
                  }
                })
              ]
            )
          ])
        ]
      )
    : _vm.type != "dataset" && _vm.modelType == 2
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
          _c("g", [
            _c("desc", [_vm._v("\n            Created with Pixso.\n    ")]),
            _vm._v(" "),
            _c("defs", [
              _c(
                "clipPath",
                {
                  attrs: {
                    id: "clip255_52-imb95sh8twfa2dm-UmZmLeRUUmZmNXPsOGaVLb"
                  }
                },
                [
                  _c("rect", {
                    attrs: {
                      id: "模型镜像-imb95sh8twfa2dm-UmZmLeRUUmZmNXPsOGaVLb",
                      width: "200.000000",
                      height: "200.000000",
                      fill: "white",
                      "fill-opacity": "0"
                    }
                  })
                ]
              )
            ]),
            _vm._v(" "),
            _c(
              "g",
              {
                attrs: {
                  "clip-path":
                    "url(#clip255_52-imb95sh8twfa2dm-UmZmLeRUUmZmNXPsOGaVLb)"
                }
              },
              [
                _c("path", {
                  attrs: {
                    id: "多边形 4-imb95sh8twfa2dm-UmZmLeRUUmZmNXPsOGaVLb",
                    d:
                      "M116.5 11.88L60.24 44.36C57.15 46.14 55.24 49.45 55.24 53.02L55.24 117.97C55.24 121.54 57.15 124.85 60.24 126.63L116.5 159.11C119.59 160.89 123.4 160.89 126.5 159.11L182.75 126.63C185.84 124.85 187.75 121.54 187.75 117.97L187.75 53.02C187.75 49.45 185.84 46.14 182.75 44.36L126.5 11.88C123.4 10.1 119.59 10.1 116.5 11.88Z",
                    fill: "#FFFFFF",
                    "fill-opacity": "1.000000",
                    "fill-rule": "evenodd"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    id: "多边形 4-imb95sh8twfa2dm-UmZmLeRUUmZmNXPsOGaVLb",
                    d:
                      "M60.24 44.36C57.15 46.14 55.24 49.45 55.24 53.02L55.24 117.97C55.24 121.54 57.15 124.85 60.24 126.63L116.5 159.11C119.59 160.89 123.4 160.89 126.5 159.11L182.75 126.63C185.84 124.85 187.75 121.54 187.75 117.97L187.75 53.02C187.75 49.45 185.84 46.14 182.75 44.36L126.5 11.88C123.4 10.1 119.59 10.1 116.5 11.88L60.24 44.36ZM65.24 53.02L65.24 117.97L121.5 150.45L177.75 117.97L177.75 53.02L121.5 20.54L65.24 53.02Z",
                    fill: "#000000",
                    "fill-opacity": "1.000000",
                    "fill-rule": "evenodd"
                  }
                }),
                _vm._v(" "),
                _c("path", {
                  attrs: {
                    id: "减去顶层-imb95sh8twfa2dm-UmZmLeRUUmZmNXPsOGaVLb",
                    d:
                      "M112.204 5.36689L65.3843 32.3984C64.3013 31.3976 63.6228 29.9625 63.6228 28.3679L63.6228 24.2557C63.6228 22.2992 64.6641 20.4917 66.3513 19.5124L98.7539 0.734482C99.1687 0.493904 99.6074 0.311348 100.07 0.18681C100.533 0.0622673 101.003 0 101.482 0C101.961 0 102.432 0.0622673 102.895 0.18681C103.357 0.311348 103.796 0.493904 104.211 0.734482L112.204 5.36689ZM182.06 138.035L191.411 132.637C191.983 132.307 192.514 131.925 193 131.499L193 146.508C193 148.464 191.958 150.271 190.265 151.25L153.904 172.266C153.058 172.756 152.146 173.001 151.169 173C150.224 173 149.281 172.755 148.434 172.266L143 169.125L143 173.806C143 175.751 141.96 177.547 140.27 178.521L104.23 199.271C103.385 199.757 102.475 200 101.5 200C100.557 200 99.6155 199.756 98.7703 199.271L62.7297 178.521C62.3149 178.282 61.938 177.993 61.5994 177.655C61.2607 177.317 60.9714 176.942 60.7319 176.528C60.4924 176.114 60.3105 175.677 60.1865 175.215C60.0623 174.754 60.0002 174.284 60 173.806L60 168.547L53.5659 172.266C52.72 172.756 51.8083 173.001 50.8308 173C49.8857 173 48.9426 172.755 48.0959 172.266L11.7349 151.25C11.3193 151.01 10.9419 150.72 10.6025 150.38C10.2629 150.04 9.97339 149.662 9.7334 149.246C9.49341 148.829 9.31104 148.389 9.18677 147.925Q9.09375 147.578 9.04712 147.224L9.04712 147.224L9.04688 147.222Q9.00024 146.868 9 146.508L9 104.476C9 101.451 11.4482 99 14.4697 99C17.4912 99 19.9397 101.451 19.9397 104.476L19.9397 143.347L50.8308 161.203L60 155.903L60 137.493L70.9187 143.797L70.9187 149.592L75.9346 146.693L86.8792 153.012L70.9187 162.237L70.9187 170.664L101.5 188.27L132.081 170.664L132.081 166.891L146.556 158.534L151.169 161.201L182.06 143.345L182.06 138.035ZM46.5889 100.042L46.5889 111.993C45.7373 111.951 44.895 111.708 44.1311 111.266L11.7285 92.4881C11.3137 92.2476 10.9373 91.9574 10.5986 91.6174C10.26 91.2775 9.97095 90.8994 9.73169 90.4831C9.49219 90.0668 9.3103 89.6266 9.18628 89.1623C9.06226 88.698 9.00024 88.2256 9 87.7449L9 50.1912C9 48.2347 10.0413 46.4273 11.7285 45.4479L44.1311 26.67C44.5459 26.4294 44.9844 26.2469 45.447 26.1223C45.9097 25.9978 46.3806 25.9355 46.8594 25.9355C47.3384 25.9355 47.8093 25.9978 48.272 26.1223C48.7346 26.2469 49.1731 26.4294 49.5879 26.67L62.4229 34.1082L51.5889 40.3633C51.5552 40.3826 51.522 40.4021 51.4888 40.4218L46.8616 37.7405L19.9158 53.3548L19.9158 84.5856L46.5889 100.042Z",
                    "clip-rule": "evenodd",
                    fill: "#101010",
                    "fill-opacity": "1.000000",
                    "fill-rule": "evenodd"
                  }
                }),
                _vm._v(" "),
                _c("line", {
                  attrs: {
                    id: "直线 1-imb95sh8twfa2dm-UmZmLeRUUmZmNXPsOGaVLb",
                    x1: "85.500000",
                    y1: "67.060120",
                    x2: "121.000000",
                    y2: "85.702629",
                    stroke: "#000000",
                    "stroke-opacity": "1.000000",
                    "stroke-width": "10.000000"
                  }
                }),
                _vm._v(" "),
                _c("line", {
                  attrs: {
                    id: "直线 1-imb95sh8twfa2dm-UmZmLeRUUmZmNXPsOGaVLb",
                    x1: "156.500000",
                    y1: "67.060127",
                    x2: "121.000000",
                    y2: "85.702637",
                    stroke: "#000000",
                    "stroke-opacity": "1.000000",
                    "stroke-width": "10.000000"
                  }
                }),
                _vm._v(" "),
                _c("line", {
                  attrs: {
                    id: "直线 1-imb95sh8twfa2dm-UmZmLeRUUmZmNXPsOGaVLb",
                    x1: "121.000000",
                    y1: "126.060120",
                    x2: "121.000000",
                    y2: "85.702637",
                    stroke: "#000000",
                    "stroke-opacity": "1.000000",
                    "stroke-width": "10.000000"
                  }
                })
              ]
            )
          ])
        ]
      )
    : _vm._e()
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=template&id=36f478aa&scoped=true":
/*!*******************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/files/FileList.vue?vue&type=template&id=36f478aa&scoped=true ***!
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
    [
      _vm.emptyPage
        ? _c(
            "div",
            { staticStyle: { "padding-top": "50px" } },
            [_c("NotFound")],
            1
          )
        : _c("div", [
            _c(
              "div",
              { staticClass: "ui container" },
              [
                !_vm.showUploadPage
                  ? _c("div", { staticClass: "content" }, [
                      _c("div", { staticClass: "top-wrap" }, [
                        _c(
                          "div",
                          { staticClass: "left-path-wrap" },
                          _vm._l(_vm.filePath, function(item, index) {
                            return _c(
                              "div",
                              {
                                key: index,
                                staticClass: "file-path nowrap",
                                on: { copy: _vm.handleCopy }
                              },
                              [
                                index == _vm.filePath.length - 1
                                  ? _c("span", [_vm._v(_vm._s(item))])
                                  : _vm._e(),
                                _vm._v(" "),
                                index != _vm.filePath.length - 1
                                  ? _c(
                                      "a",
                                      {
                                        staticClass: "canback",
                                        on: {
                                          click: function($event) {
                                            return _vm.goBackDir(index)
                                          }
                                        }
                                      },
                                      [_vm._v(_vm._s(item))]
                                    )
                                  : _vm._e(),
                                _vm._v(" "),
                                _c("span", { staticClass: "divider" }, [
                                  _vm._v(" / ")
                                ])
                              ]
                            )
                          }),
                          0
                        ),
                        _vm._v(" "),
                        _c(
                          "div",
                          { staticClass: "right-btn-c" },
                          [
                            _vm.dataObj.can_download
                              ? _c(
                                  "a",
                                  {
                                    staticClass: "download-btn",
                                    on: {
                                      click: function($event) {
                                        return _vm.downloadAllModel()
                                      }
                                    }
                                  },
                                  [
                                    _vm._v(
                                      "\n                " +
                                        _vm._s(
                                          _vm.type === "dataset"
                                            ? _vm.$t(
                                                "datasetObj.dataDownloadAll"
                                              )
                                            : _vm.$t(
                                                "modelManage.modelDownloadAll"
                                              )
                                        ) +
                                        "\n              "
                                    )
                                  ]
                                )
                              : _vm._e(),
                            _vm._v(" "),
                            _vm.dataObj.can_edit_file
                              ? _c(
                                  "el-button",
                                  {
                                    staticStyle: {
                                      "background-color": "#0066ff"
                                    },
                                    attrs: { type: "primary" },
                                    on: { click: _vm.goUploadPage }
                                  },
                                  [
                                    _vm._v(
                                      "\n                " +
                                        _vm._s(
                                          _vm.type === "dataset"
                                            ? _vm.$t("uploadDatasetFile")
                                            : _vm.$t(
                                                "modelManage.uploadModelFiles"
                                              )
                                        ) +
                                        "\n              "
                                    )
                                  ]
                                )
                              : _vm._e()
                          ],
                          1
                        )
                      ]),
                      _vm._v(" "),
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
                              ref: "tableRef",
                              staticStyle: { width: "100%" },
                              attrs: {
                                data: _vm.filesList,
                                "header-cell-style": _vm.headerStyle,
                                "cell-style": _vm.cellStyle,
                                "max-height": "800px"
                              },
                              on: { "sort-change": _vm.sortChange }
                            },
                            [
                              _c("el-table-column", {
                                attrs: {
                                  "column-key": "FileName",
                                  prop: "FileName",
                                  sortable: "custom",
                                  label: _vm.$t("modelManage.fileName"),
                                  align: "left",
                                  "header-align": "center",
                                  "min-width": "260"
                                },
                                scopedSlots: _vm._u(
                                  [
                                    {
                                      key: "default",
                                      fn: function(scope) {
                                        return [
                                          _c(
                                            "div",
                                            { staticClass: "tbl-file-name" },
                                            [
                                              scope.row.IsDir
                                                ? _c(
                                                    "div",
                                                    {
                                                      staticClass:
                                                        "fitted folder",
                                                      on: {
                                                        click: function(
                                                          $event
                                                        ) {
                                                          return _vm.goNextDir(
                                                            scope.row
                                                          )
                                                        }
                                                      }
                                                    },
                                                    [
                                                      _c("i", {
                                                        staticClass:
                                                          "ri-folder-5-fill",
                                                        staticStyle: {
                                                          color: "#ffb02c",
                                                          "font-size": "16px"
                                                        }
                                                      }),
                                                      _vm._v(" "),
                                                      _c(
                                                        "span",
                                                        {
                                                          staticClass: "nowrap",
                                                          attrs: {
                                                            title:
                                                              scope.row.FileName
                                                          }
                                                        },
                                                        [
                                                          _vm._v(
                                                            _vm._s(
                                                              scope.row.FileName
                                                            )
                                                          )
                                                        ]
                                                      )
                                                    ]
                                                  )
                                                : _c(
                                                    "div",
                                                    { staticClass: "fitted" },
                                                    [
                                                      _c("i", {
                                                        staticClass:
                                                          "ri-file-text-fill",
                                                        staticStyle: {
                                                          color: "#8ca2aa",
                                                          "font-size": "16px"
                                                        }
                                                      }),
                                                      _vm._v(" "),
                                                      _c(
                                                        "span",
                                                        {
                                                          staticClass: "nowrap",
                                                          attrs: {
                                                            title:
                                                              scope.row.FileName
                                                          }
                                                        },
                                                        [
                                                          _vm._v(
                                                            _vm._s(
                                                              scope.row.FileName
                                                            )
                                                          )
                                                        ]
                                                      ),
                                                      _vm._v(" "),
                                                      _c("i", {
                                                        staticClass:
                                                          "ri-file-copy-2-line ui poping up clipboard-model-name",
                                                        attrs: {
                                                          "data-position":
                                                            "top center",
                                                          "data-variation":
                                                            "inverted tiny",
                                                          "data-success": _vm.$t(
                                                            "copySuccess"
                                                          ),
                                                          "data-content": _vm.$t(
                                                            "copy"
                                                          ),
                                                          "data-original": _vm.$t(
                                                            "copy"
                                                          ),
                                                          "data-clipboard-text":
                                                            scope.row.FileName
                                                        }
                                                      })
                                                    ]
                                                  )
                                            ]
                                          )
                                        ]
                                      }
                                    }
                                  ],
                                  null,
                                  false,
                                  1369410280
                                )
                              }),
                              _vm._v(" "),
                              _c("el-table-column", {
                                attrs: {
                                  "column-key": "SizeShow",
                                  prop: "Size",
                                  sortable: "custom",
                                  label: _vm.$t("modelManage.fileSize"),
                                  align: "center",
                                  "header-align": "center",
                                  width: "200"
                                },
                                scopedSlots: _vm._u(
                                  [
                                    {
                                      key: "default",
                                      fn: function(scope) {
                                        return [
                                          _vm._v(
                                            "\n                      " +
                                              _vm._s(
                                                scope.row.IsDir
                                                  ? "--"
                                                  : _vm.formatFileSize(
                                                      scope.row.Size
                                                    )
                                              ) +
                                              "\n                  "
                                          )
                                        ]
                                      }
                                    }
                                  ],
                                  null,
                                  false,
                                  3209513433
                                )
                              }),
                              _vm._v(" "),
                              _c("el-table-column", {
                                attrs: {
                                  "column-key": "ModTime",
                                  prop: "ModTime",
                                  sortable: "custom",
                                  label: _vm.$t("modelManage.updateTime"),
                                  align: "center",
                                  "header-align": "center",
                                  width: "200"
                                }
                              }),
                              _vm._v(" "),
                              _c("el-table-column", {
                                attrs: {
                                  "column-key": "operate",
                                  prop: "operate",
                                  fixed: "right",
                                  label: _vm.$t("modelManage.operate"),
                                  align: "center",
                                  "header-align": "center",
                                  width: "200"
                                },
                                scopedSlots: _vm._u(
                                  [
                                    {
                                      key: "default",
                                      fn: function(scope) {
                                        return [
                                          scope.row.IsSupportPreview &&
                                          _vm.type == "dataset"
                                            ? _c(
                                                "span",
                                                {
                                                  staticClass: "btn-del",
                                                  on: {
                                                    click: function($event) {
                                                      return _vm.previewFile(
                                                        scope.row
                                                      )
                                                    }
                                                  }
                                                },
                                                [
                                                  _vm._v(
                                                    "\n                      " +
                                                      _vm._s(
                                                        _vm.$t(
                                                          "modelManage.preview"
                                                        )
                                                      ) +
                                                      "\n                    "
                                                  )
                                                ]
                                              )
                                            : _vm._e(),
                                          _vm._v(" "),
                                          !scope.row.IsDir &&
                                          _vm.dataObj.can_download
                                            ? _c(
                                                "span",
                                                {
                                                  staticClass: "btn-del",
                                                  on: {
                                                    click: function($event) {
                                                      return _vm.downLoadFile(
                                                        scope.row
                                                      )
                                                    }
                                                  }
                                                },
                                                [
                                                  _vm._v(
                                                    "\n                      " +
                                                      _vm._s(
                                                        _vm.$t(
                                                          "modelManage.download"
                                                        )
                                                      ) +
                                                      "\n                    "
                                                  )
                                                ]
                                              )
                                            : _vm._e(),
                                          _vm._v(" "),
                                          !scope.row.IsDir &&
                                          _vm.dataObj.can_edit_file
                                            ? _c(
                                                "span",
                                                {
                                                  staticClass: "btn-del",
                                                  on: {
                                                    click: function($event) {
                                                      return _vm.deleteFile(
                                                        scope.row
                                                      )
                                                    }
                                                  }
                                                },
                                                [
                                                  _vm._v(
                                                    "\n                      " +
                                                      _vm._s(
                                                        _vm.$t(
                                                          "modelManage.delete"
                                                        )
                                                      ) +
                                                      "\n                    "
                                                  )
                                                ]
                                              )
                                            : _vm._e()
                                        ]
                                      }
                                    }
                                  ],
                                  null,
                                  false,
                                  3650748601
                                )
                              })
                            ],
                            1
                          ),
                          _vm._v(" "),
                          _vm.has_next
                            ? _c(
                                "el-button",
                                {
                                  staticStyle: { "margin-top": "24px" },
                                  on: { click: _vm.getDatasetFileList }
                                },
                                [_vm._v(_vm._s(_vm.$t("datasets.loadMore")))]
                              )
                            : _vm._e()
                        ],
                        1
                      )
                    ])
                  : _c(
                      "CreateForm",
                      {
                        staticClass: "upload-wrap",
                        attrs: {
                          showSteps: false,
                          title: [
                            _vm.type === "dataset"
                              ? _vm.$t("uploadDatasetFile")
                              : _vm.$t("modelManage.uploadModelFiles")
                          ]
                        }
                      },
                      [
                        _c("FileUpload", {
                          attrs: {
                            modelId: _vm.dataObj.id,
                            subjectType: _vm.type === "dataset" ? "1" : "2",
                            ownerName: _vm.dataObj.owner_name,
                            uploadDir: _vm.params.parent_dir
                          },
                          on: {
                            cancel: _vm.cancel,
                            uploadFinish: _vm.uploadFinish
                          }
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
        "BaseDialog",
        {
          attrs: { visible: _vm.visible, title: _vm.title },
          on: { closed: _vm.closeDialog }
        },
        [
          _c(
            "div",
            { staticClass: "body-c" },
            [
              _vm.showFlag
                ? _c("div", { staticClass: "not-support" }, [
                    _c(
                      "svg",
                      {
                        staticStyle: { "enable-background": "new 0 0 200 200" },
                        attrs: {
                          width: "168",
                          height: "168",
                          xmlns: "http://www.w3.org/2000/svg",
                          "xml:space": "preserve",
                          viewBox: "0 0 200 200"
                        }
                      },
                      [
                        _c("path", {
                          staticStyle: { fill: "#66c8ff" },
                          attrs: {
                            d:
                              "M64.2 137.3H40s1.3-1.6 1.3-10.7c0-10.2-5.8-11.1-5.8-24.5 0-10.5 4-15.1 4-24.3 0-8-3.2-15.9-3.2-15.9h27.9s-4.3 30.6-4.3 39.7c0 9 4.3 35.7 4.3 35.7zm110.7 0-19.4.6s-4.2.4-4.2-8.7c0-10.2 5.6-13.7 5.6-27 0-10.5-5.9-16.3-5.9-25.5 0-8 2-14.7 2-14.7h21.8s4.4 30.6 4.4 39.7-4.3 35.6-4.3 35.6z"
                          }
                        }),
                        _c("path", {
                          staticStyle: { fill: "#78c6ff" },
                          attrs: {
                            d:
                              "M185.4 6.8H14.6c-6.8 0-12.3 5.5-12.3 12.3v117.5c0 6.8 5.5 12.3 12.3 12.3h9.3c.6 1.7 1.6 3.5 3.1 5 .4.4.8.7 1.4.8 15.5 4 19.3 9 20.3 11.6 1 2.9-.7 5.2-.8 5.4-1 1.3-.8 3.1.5 4.1 1.3 1 3.1.8 4.2-.5 1.4-1.8 3.4-6 1.8-10.8-2.2-6.5-10.1-11.6-23.5-15.2a5.03 5.03 0 0 1-1.4-3.3c0-1.2.7-2.4 2.1-3.4 0 0 .1 0 .1-.1.4-.3.8-.6 1.3-.8.1-.1.2-.1.4-.2.6-.3 1.2-.5 1.9-.7 1-.2 2.9-.6 5.9-.6 5.2 0 11.2 1 18 3 1.1.5 2.3 1 3.4 1.6.4.2.7.4 1 .6 1.4.8 2.9 1.7 4.3 2.6.1.1.2.1.4.2 22.6 15.9 20.1 24 20.1 24-.8 1.4-.2 3.2 1.2 4 .4.2.9.4 1.4.4 1.1 0 2.1-.6 2.6-1.6 1-1.8 4-10.6-14.7-26H131c-18.8 15.5-15.7 24.3-14.8 26.1.6 1 1.6 1.6 2.7 1.6.5 0 .9-.1 1.3-.4 1.4-.8 2-2.5 1.3-3.9 0-.1-2.7-8.1 20.1-24.1 1.6-1.1 3.2-2 4.8-3 .3-.2.6-.3.8-.5 1.2-.6 2.4-1.2 3.5-1.7 6.7-2 12.7-3 17.9-3 3 0 4.8.4 5.8.6.6.2 1.3.5 1.9.7.1 0 .2.1.3.1.5.2 1 .6 1.4.8 1.4 1 2.1 2.2 2.1 3.4 0 .9-.4 2-1.4 3.3-13.4 3.6-21.3 8.7-23.5 15.3-1.6 4.8.4 9.1 1.8 10.8a3 3 0 0 0 4.2.5 3 3 0 0 0 .5-4.2s-1.8-2.4-.8-5.3c.9-2.6 4.7-7.7 20.3-11.7.5-.1 1-.4 1.4-.8 1.6-1.6 2.6-3.4 3.1-5.1 6.7-.1 12.1-5.6 12.1-12.3V19.1a12.5 12.5 0 0 0-12.4-12.3zM33.7 63.7h26.4c1.1 3.9 3 13.9-1.4 26.4l-.4.9c-2.2 6-4.8 13.5 1.1 29 3.6 9.4 1.4 14.8-.1 17.1a46.37 46.37 0 0 0-11-2.9l-.8-.1c-1.5-.2-3-.3-4.4-.3h-1c-1.2 0-2.3.1-3.4.2-.3 0-.6 0-.8.1-.2 0-.5.1-.7.1 1-4.9 1.7-12.2-1.5-17-3.5-5.2-4.3-21.3-.1-33.4 1.9-5.3-.3-14.7-1.9-20.1zm116.4 56.2c6-15.5 3.3-23 1.1-29l-.3-.8c-4.4-12.6-2.5-22.5-1.4-26.4h26.4c-1.6 5.4-3.7 14.9-1.9 20.1 4.2 12.2 3.4 28.3 0 33.4-3.2 4.9-2.5 12.1-1.5 17-.2 0-.5-.1-.7-.1-.2 0-.4 0-.6-.1-1.2-.2-2.5-.2-3.8-.3h-.5c-4.9 0-10.5 1-16.4 3.3-1.8-2.2-4-7.6-.4-17.1zm41.7 16.8c0 3.5-2.8 6.3-6.2 6.4-.4-1.1-1-2.2-1.9-3.2l-.4-.4-.1-.1c-.4-.5-1-.9-1.4-1.3-.2-.1-.4-.2-.5-.4-.4-.3-.9-.6-1.4-.8-.2-.1-.4-.2-.6-.4-.1 0-.1 0-.2-.1-1.1-3.8-2.8-12-.2-15.8 4.9-7.4 5.3-25.5.7-38.7-1.2-3.6.7-12.3 2.5-18.2h1.2a3 3 0 0 0 3-3 3 3 0 0 0-3-3h-38.4a3 3 0 0 0-3 3c0 1.2.7 2.1 1.6 2.6-1.3 5.3-2.8 15.7 1.9 28.8l.3.9c1.9 5.4 4.1 11.4-1.1 24.8-4.2 10.9-2 18 .2 21.7-2 1.1-4 2.2-6 3.6H70.7c-.2-.1-.4-.2-.6-.4-1.6-1.1-3.2-2-4.8-3-.2-.1-.4-.2-.6-.4 2.2-3.8 4.4-10.9.2-21.7-5.2-13.4-3-19.4-1.1-24.8l.4-.9c4.6-13.1 3.2-23.5 1.9-28.8 1-.5 1.6-1.4 1.6-2.6a3 3 0 0 0-3-3H26.4a3 3 0 0 0-3 3 3 3 0 0 0 3 3h1.2c1.8 5.9 3.7 14.6 2.5 18.2-4.4 12.9-4.1 31.4.7 38.7 2.6 3.8.9 12-.2 15.8-.1 0-.2.1-.2.1-.2.1-.4.2-.6.4-.5.2-.8.5-1.3.8l-.6.4c-.5.4-1 .8-1.4 1.2l-.2.2-.4.4c-.4.4-.6.8-1 1.3-.4.6-.7 1.3-1 1.9h-9.3c-3.6 0-6.4-2.9-6.4-6.4V19.2c0-3.6 2.9-6.4 6.4-6.4h170.8c3.6 0 6.4 2.9 6.4 6.4v117.5zM93.7 49.9a3 3 0 0 0-3-3c-10.1 0-18.6-6.2-22.6-9.2-1-.7-1.7-1.2-2.2-1.5-1.4-.8-3.2-.4-4.1 1-.8 1.4-.4 3.2 1 4.1.4.2 1 .7 1.7 1.2 4.6 3.4 14.1 10.3 26.1 10.3 1.8 0 3.1-1.3 3.1-2.9zm42.5-13.7c-.5.3-1.2.8-2.2 1.5-4.1 3-12.6 9.2-22.6 9.2a3 3 0 0 0-3 3 3 3 0 0 0 3 3c12 0 21.6-7 26.1-10.3.7-.6 1.3-1 1.7-1.2 1.4-.8 1.9-2.6 1-4.1-.8-1.5-2.6-1.9-4-1.1zm-15 80.5-19.9-9c-.8-.4-1.8-.3-2.5 0l-17.6 9c-1.4.7-2 2.5-1.3 4 .7 1.4 2.5 2 4 1.3l16.4-8.3 18.6 8.4c.4.2.8.2 1.2.2 1.1 0 2.2-.6 2.7-1.8.5-1.4-.2-3.1-1.6-3.8z"
                          }
                        }),
                        _c("path", {
                          staticStyle: { fill: "#444" },
                          attrs: {
                            d:
                              "M185.4 6.8H14.6c-6.8 0-12.3 5.5-12.3 12.3v117.5c0 6.8 5.5 12.3 12.3 12.3h9.3c.6 1.7 1.6 3.5 3.1 5 .4.4.8.7 1.4.8 15.5 4 19.3 9 20.3 11.6 1 2.9-.7 5.2-.8 5.4-1 1.3-.8 3.1.5 4.1 1.3 1 3.1.8 4.2-.5 1.4-1.8 3.4-6 1.8-10.8-2.2-6.5-10.1-11.6-23.5-15.2a5.03 5.03 0 0 1-1.4-3.3c0-1.2.7-2.4 2.1-3.4 0 0 .1 0 .1-.1.4-.3.8-.6 1.3-.8.1-.1.2-.1.4-.2.6-.3 1.2-.5 1.9-.7 1-.2 2.9-.6 5.9-.6 5.2 0 11.2 1 18 3 1.1.5 2.3 1 3.4 1.6.4.2.7.4 1 .6 1.4.8 2.9 1.7 4.3 2.6.1.1.2.1.4.2 22.6 15.9 20.1 24 20.1 24-.8 1.4-.2 3.2 1.2 4 .4.2.9.4 1.4.4 1.1 0 2.1-.6 2.6-1.6 1-1.8 4-10.6-14.7-26H131c-18.8 15.5-15.7 24.3-14.8 26.1.6 1 1.6 1.6 2.7 1.6.5 0 .9-.1 1.3-.4 1.4-.8 2-2.5 1.3-3.9 0-.1-2.7-8.1 20.1-24.1 1.6-1.1 3.2-2 4.8-3 .3-.2.6-.3.8-.5 1.2-.6 2.4-1.2 3.5-1.7 6.7-2 12.7-3 17.9-3 3 0 4.8.4 5.8.6.6.2 1.3.5 1.9.7.1 0 .2.1.3.1.5.2 1 .6 1.4.8 1.4 1 2.1 2.2 2.1 3.4 0 .9-.4 2-1.4 3.3-13.4 3.6-21.3 8.7-23.5 15.3-1.6 4.8.4 9.1 1.8 10.8a3 3 0 0 0 4.2.5 3 3 0 0 0 .5-4.2s-1.8-2.4-.8-5.3c.9-2.6 4.7-7.7 20.3-11.7.5-.1 1-.4 1.4-.8 1.6-1.6 2.6-3.4 3.1-5.1 6.7-.1 12.1-5.6 12.1-12.3V19.1a12.5 12.5 0 0 0-12.4-12.3zM33.7 63.7h26.4c1.1 3.9 3 13.9-1.4 26.4l-.4.9c-2.2 6-4.8 13.5 1.1 29 3.6 9.4 1.4 14.8-.1 17.1a46.37 46.37 0 0 0-11-2.9l-.8-.1c-1.5-.2-3-.3-4.4-.3h-1c-1.2 0-2.3.1-3.4.2-.3 0-.6 0-.8.1-.2 0-.5.1-.7.1 1-4.9 1.7-12.2-1.5-17-3.5-5.2-4.3-21.3-.1-33.4 1.9-5.3-.3-14.7-1.9-20.1zm116.4 56.2c6-15.5 3.3-23 1.1-29l-.3-.8c-4.4-12.6-2.5-22.5-1.4-26.4h26.4c-1.6 5.4-3.7 14.9-1.9 20.1 4.2 12.2 3.4 28.3 0 33.4-3.2 4.9-2.5 12.1-1.5 17-.2 0-.5-.1-.7-.1-.2 0-.4 0-.6-.1-1.2-.2-2.5-.2-3.8-.3h-.5c-4.9 0-10.5 1-16.4 3.3-1.8-2.2-4-7.6-.4-17.1zm41.7 16.8c0 3.5-2.8 6.3-6.2 6.4-.4-1.1-1-2.2-1.9-3.2l-.4-.4-.1-.1c-.4-.5-1-.9-1.4-1.3-.2-.1-.4-.2-.5-.4-.4-.3-.9-.6-1.4-.8-.2-.1-.4-.2-.6-.4-.1 0-.1 0-.2-.1-1.1-3.8-2.8-12-.2-15.8 4.9-7.4 5.3-25.5.7-38.7-1.2-3.6.7-12.3 2.5-18.2h1.2a3 3 0 0 0 3-3 3 3 0 0 0-3-3h-38.4a3 3 0 0 0-3 3c0 1.2.7 2.1 1.6 2.6-1.3 5.3-2.8 15.7 1.9 28.8l.3.9c1.9 5.4 4.1 11.4-1.1 24.8-4.2 10.9-2 18 .2 21.7-2 1.1-4 2.2-6 3.6H70.7c-.2-.1-.4-.2-.6-.4-1.6-1.1-3.2-2-4.8-3-.2-.1-.4-.2-.6-.4 2.2-3.8 4.4-10.9.2-21.7-5.2-13.4-3-19.4-1.1-24.8l.4-.9c4.6-13.1 3.2-23.5 1.9-28.8 1-.5 1.6-1.4 1.6-2.6a3 3 0 0 0-3-3H26.4a3 3 0 0 0-3 3 3 3 0 0 0 3 3h1.2c1.8 5.9 3.7 14.6 2.5 18.2-4.4 12.9-4.1 31.4.7 38.7 2.6 3.8.9 12-.2 15.8-.1 0-.2.1-.2.1-.2.1-.4.2-.6.4-.5.2-.8.5-1.3.8l-.6.4c-.5.4-1 .8-1.4 1.2l-.2.2-.4.4c-.4.4-.6.8-1 1.3-.4.6-.7 1.3-1 1.9h-9.3c-3.6 0-6.4-2.9-6.4-6.4V19.2c0-3.6 2.9-6.4 6.4-6.4h170.8c3.6 0 6.4 2.9 6.4 6.4v117.5zM93.7 49.9a3 3 0 0 0-3-3c-10.1 0-18.6-6.2-22.6-9.2-1-.7-1.7-1.2-2.2-1.5-1.4-.8-3.2-.4-4.1 1-.8 1.4-.4 3.2 1 4.1.4.2 1 .7 1.7 1.2 4.6 3.4 14.1 10.3 26.1 10.3 1.8 0 3.1-1.3 3.1-2.9zm42.5-13.7c-.5.3-1.2.8-2.2 1.5-4.1 3-12.6 9.2-22.6 9.2a3 3 0 0 0-3 3 3 3 0 0 0 3 3c12 0 21.6-7 26.1-10.3.7-.6 1.3-1 1.7-1.2 1.4-.8 1.9-2.6 1-4.1-.8-1.5-2.6-1.9-4-1.1zm-15 80.5-19.9-9c-.8-.4-1.8-.3-2.5 0l-17.6 9c-1.4.7-2 2.5-1.3 4 .7 1.4 2.5 2 4 1.3l16.4-8.3 18.6 8.4c.4.2.8.2 1.2.2 1.1 0 2.2-.6 2.7-1.8.5-1.4-.2-3.1-1.6-3.8z"
                          }
                        })
                      ]
                    ),
                    _vm._v(" "),
                    _c("div", [_vm._v(_vm._s(_vm.showMessage))])
                  ])
                : [
                    _vm.fileImg
                      ? _c("div", { staticClass: "img-c" }, [
                          _c("img", { attrs: { src: _vm.fileImg, alt: "" } })
                        ])
                      : _vm._e(),
                    _vm._v(" "),
                    _vm.textContent
                      ? _c("div", { staticClass: "file-content" }, [
                          _c("pre", [_vm._v(_vm._s(_vm.textContent))])
                        ])
                      : _vm._e()
                  ]
            ],
            2
          )
        ]
      ),
      _vm._v(" "),
      _c("CommonSDK", {
        ref: "childModelTips",
        attrs: {
          data: _vm.sdkData,
          title: _vm.$t("modelObj.codeDownDlgTitle"),
          closeText: _vm.$t("cloudbrainObj.dialogTips.tips8")
        }
      })
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=template&id=6704b8e1&scoped=true":
/*!****************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/Intro.vue?vue&type=template&id=6704b8e1&scoped=true ***!
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
  return _c("div", { staticClass: "ui container content" }, [
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
        staticClass: "content-l"
      },
      [
        !(_vm.introEmpty && !_vm.editing) && !_vm.loading
          ? _c("div", { staticClass: "intro-content" }, [
              !_vm.editing
                ? _c("div", { staticClass: "read-mode" }, [
                    _c("div", { staticClass: "head" }, [
                      _c("div", [
                        _c("i", { staticClass: "el-icon-warning-outline" }),
                        _vm._v(" "),
                        _c("span", [_vm._v(_vm._s(_vm.introFileName))])
                      ]),
                      _vm._v(" "),
                      _c("div", [
                        _vm.canEdit
                          ? _c("i", {
                              staticClass: "icon pencil alternate",
                              attrs: { title: _vm.$t("modelManage.edit") },
                              on: {
                                click: function($event) {
                                  return _vm.toggleEdit(true)
                                }
                              }
                            })
                          : _vm._e()
                      ])
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "content-box" }, [
                      _c("div", {
                        staticClass: "file-view markdown markdown",
                        domProps: { innerHTML: _vm._s(_vm.htmlContent) }
                      })
                    ])
                  ])
                : _vm._e(),
              _vm._v(" "),
              _vm.canEdit && _vm.editing
                ? _c("div", { staticClass: "edit-mode" }, [
                    _c("div", { staticClass: "head" }, [
                      _c("div", [
                        _c("i", { staticClass: "el-icon-warning-outline" }),
                        _vm._v(" "),
                        _c("span", [_vm._v(_vm._s(_vm.introFileName))])
                      ]),
                      _vm._v(" "),
                      _c("div", [
                        _vm.canEdit
                          ? _c("i", {
                              staticClass: "icon reply",
                              attrs: { title: _vm.$t("cancel") },
                              on: {
                                click: function($event) {
                                  return _vm.toggleEdit(false)
                                }
                              }
                            })
                          : _vm._e()
                      ])
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "item-tab-c" }, [
                      _c(
                        "div",
                        {
                          staticClass: "item-tab",
                          class: _vm.editTab == "edit" ? "focus" : "",
                          attrs: { tab: "edit" },
                          on: {
                            click: function($event) {
                              return _vm.changeEditTab("edit")
                            }
                          }
                        },
                        [
                          _c(
                            "svg",
                            {
                              staticClass: "svg octicon-code",
                              attrs: {
                                width: "16",
                                height: "16",
                                "aria-hidden": "true"
                              }
                            },
                            [
                              _c("use", {
                                attrs: { "xlink:href": "#octicon-code" }
                              })
                            ]
                          ),
                          _vm._v(" "),
                          _c("span", [
                            _vm._v(_vm._s(_vm.$t("modelManage.editFiles")))
                          ])
                        ]
                      ),
                      _vm._v(" "),
                      _c(
                        "div",
                        {
                          staticClass: "item-tab",
                          class: _vm.editTab == "preview" ? "focus" : "",
                          attrs: { tab: "preview" },
                          on: {
                            click: function($event) {
                              return _vm.changeEditTab("preview")
                            }
                          }
                        },
                        [
                          _c(
                            "svg",
                            {
                              staticClass: "svg octicon-eye",
                              attrs: {
                                width: "16",
                                height: "16",
                                "aria-hidden": "true"
                              }
                            },
                            [
                              _c("use", {
                                attrs: { "xlink:href": "#octicon-eye" }
                              })
                            ]
                          ),
                          _vm._v(" "),
                          _c("span", [
                            _vm._v(_vm._s(_vm.$t("modelManage.preview")))
                          ])
                        ]
                      )
                    ]),
                    _vm._v(" "),
                    _c(
                      "div",
                      {
                        directives: [
                          {
                            name: "show",
                            rawName: "v-show",
                            value: _vm.editTab == "edit",
                            expression: "editTab == 'edit'"
                          }
                        ],
                        staticClass: "tab-content edit-content"
                      },
                      [
                        _c("div", {
                          directives: [
                            {
                              name: "loading",
                              rawName: "v-loading",
                              value: _vm.editLoading,
                              expression: "editLoading"
                            }
                          ],
                          ref: "editContainerRef",
                          staticClass: "monaco-editor-container"
                        })
                      ]
                    ),
                    _vm._v(" "),
                    _c(
                      "div",
                      {
                        directives: [
                          {
                            name: "show",
                            rawName: "v-show",
                            value: _vm.editTab == "preview",
                            expression: "editTab == 'preview'"
                          },
                          {
                            name: "loading",
                            rawName: "v-loading",
                            value: _vm.previewLoading,
                            expression: "previewLoading"
                          }
                        ],
                        staticClass: "tab-content preview-content"
                      },
                      [
                        _c("div", {
                          ref: "previewContainerRef",
                          staticClass: "preview-container markdown",
                          domProps: { innerHTML: _vm._s(_vm.previewContent) }
                        })
                      ]
                    ),
                    _vm._v(" "),
                    _c(
                      "div",
                      { staticClass: "conmit-btn-c" },
                      [
                        _c(
                          "el-button",
                          {
                            directives: [
                              {
                                name: "loading",
                                rawName: "v-loading",
                                value: _vm.submitLoading,
                                expression: "submitLoading"
                              }
                            ],
                            staticClass: "btn confirm-btn",
                            class:
                              this.content == this.editContent
                                ? "_disabled_"
                                : "",
                            attrs: {
                              size: "default",
                              disabled: this.content == this.editContent
                            },
                            on: { click: _vm.submit }
                          },
                          [
                            _vm._v(
                              "\n            " + _vm._s(_vm.$t("submit")) + " "
                            )
                          ]
                        ),
                        _vm._v(" "),
                        _c(
                          "el-button",
                          {
                            staticClass: "btn",
                            attrs: { size: "default" },
                            on: {
                              click: function($event) {
                                return _vm.toggleEdit(false)
                              }
                            }
                          },
                          [_vm._v(_vm._s(_vm.$t("cancel")))]
                        )
                      ],
                      1
                    )
                  ])
                : _vm._e()
            ])
          : _vm._e(),
        _vm._v(" "),
        _vm.introEmpty && !_vm.editing && !_vm.loading
          ? _c("div", { staticClass: "empty" }, [
              _vm._m(0),
              _vm._v(" "),
              _c("div", { staticClass: "tips" }, [
                _vm._v(
                  _vm._s(
                    _vm.type == "dataset"
                      ? _vm.$t("datasetObj.hasNoIntroForModel")
                      : _vm.$t("modelManage.hasNoIntroForModel")
                  )
                )
              ]),
              _vm._v(" "),
              _vm.canEdit
                ? _c(
                    "div",
                    { staticClass: "ops" },
                    [
                      _c(
                        "el-button",
                        {
                          staticClass: "create-btn",
                          attrs: { size: "default" },
                          on: { click: _vm.createIntro }
                        },
                        [
                          _vm._v(
                            "\n          " +
                              _vm._s(
                                _vm.type == "dataset"
                                  ? _vm.$t("datasetObj.createModelIntro")
                                  : _vm.$t("modelManage.createModelIntro")
                              ) +
                              "\n        "
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
    ),
    _vm._v(" "),
    _c(
      "div",
      {
        directives: [
          {
            name: "show",
            rawName: "v-show",
            value: !_vm.editing,
            expression: "!editing"
          }
        ],
        staticClass: "content-r",
        style: { marginTop: !_vm.introEmpty ? "28px" : "" }
      },
      [
        _c(
          "div",
          { staticClass: "summary" },
          [
            _c("div", { staticClass: "row" }, [
              _c("div", { staticClass: "label" }, [
                _vm._v(_vm._s(_vm.$t("datasetObj.dataset_owner")) + "：")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "value" }, [
                _c("a", { attrs: { href: "/" + _vm.dataObj.owner_name } }, [
                  _vm._v(_vm._s(_vm.dataObj.owner_name))
                ])
              ])
            ]),
            _vm._v(" "),
            _vm.type === "dataset"
              ? [
                  _c("div", { staticClass: "row" }, [
                    _c("div", { staticClass: "label" }, [
                      _vm._v(_vm._s(_vm.$t("datasetObj.category")) + "：")
                    ]),
                    _vm._v(" "),
                    _c(
                      "div",
                      { staticClass: "value" },
                      _vm._l(_vm.dataObj.tags, function(item) {
                        return _vm.dataObj.tags.length
                          ? _c("span", { key: item }, [
                              _vm._v(
                                "\n              " +
                                  _vm._s(_vm.$t("datasets." + item)) +
                                  "\n            "
                              )
                            ])
                          : _c("span", [_vm._v("--")])
                      }),
                      0
                    )
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "row" }, [
                    _c("div", { staticClass: "label" }, [
                      _vm._v(_vm._s(_vm.$t("datasetObj.application")) + "：")
                    ]),
                    _vm._v(" "),
                    _c(
                      "div",
                      { staticClass: "value" },
                      _vm._l(_vm.dataObj.tasks, function(item) {
                        return _vm.dataObj.tasks.length
                          ? _c("span", { key: item }, [
                              _vm._v(
                                "\n              " +
                                  _vm._s(_vm.$t("datasets." + item)) +
                                  "\n            "
                              )
                            ])
                          : _c("span", [_vm._v("--")])
                      }),
                      0
                    )
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "row" }, [
                    _c("div", { staticClass: "label" }, [
                      _vm._v(_vm._s(_vm.$t("datasetObj.license")) + "：")
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "value" }, [
                      _c("span", [_vm._v(_vm._s(_vm.dataObj.licenses || "--"))])
                    ])
                  ])
                ]
              : [
                  _c("div", { staticClass: "row" }, [
                    _c("div", { staticClass: "label" }, [
                      _vm._v(_vm._s(_vm.$t("modelManage.license")) + "：")
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "value" }, [
                      _vm.dataObj &&
                      _vm.dataObj.licenseInfo &&
                      _vm.dataObj.licenseInfo.name
                        ? _c(
                            "a",
                            {
                              attrs: {
                                target: "_blank",
                                href:
                                  _vm.dataObj &&
                                  _vm.dataObj.licenseInfo &&
                                  _vm.dataObj.licenseInfo.linkUrl
                              }
                            },
                            [
                              _vm._v(
                                "\n              " +
                                  _vm._s(
                                    _vm.dataObj &&
                                      _vm.dataObj.licenseInfo &&
                                      _vm.dataObj.licenseInfo.name
                                  ) +
                                  "\n            "
                              )
                            ]
                          )
                        : _c("span", [_vm._v("--")])
                    ])
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "row" }, [
                    _c("div", { staticClass: "label" }, [
                      _vm._v(_vm._s(_vm.$t("modelManage.modelEngine")) + "：")
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "value" }, [
                      _c("span", [
                        _vm._v(_vm._s(_vm.dataObj.engineName || "--"))
                      ])
                    ])
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "row" }, [
                    _c("div", { staticClass: "label" }, [
                      _vm._v(_vm._s(_vm.$t("modelManage.modelSource")) + "：")
                    ]),
                    _vm._v(" "),
                    _c("div", { staticClass: "value" }, [
                      _vm.dataObj.aimodel_type == 0
                        ? _c("span", [
                            _c("span", { staticClass: "model-type online" }, [
                              _vm._v(_vm._s(_vm.$t("modelManage.online")))
                            ]),
                            _c("span", {
                              domProps: {
                                innerHTML: _vm._s(
                                  _vm.dataObj.displayJobNameHtml
                                )
                              }
                            })
                          ])
                        : _vm._e(),
                      _vm._v(" "),
                      _vm.dataObj.aimodel_type == 1
                        ? _c("span", { staticClass: "model-type local" }, [
                            _vm._v(_vm._s(_vm.$t("modelManage.local")))
                          ])
                        : _vm._e(),
                      _vm._v(" "),
                      _c(
                        "div",
                        { staticStyle: { display: "flex" } },
                        [
                          _vm.dataObj.aimodel_type == 2
                            ? _c(
                                "span",
                                { staticClass: "model-type external" },
                                [_vm._v(_vm._s(_vm.$t("modelManage.external")))]
                              )
                            : _vm._e(),
                          _vm._v(" "),
                          _vm.dataObj.aimodel_type == 2
                            ? _c("MigrateModelSync", {
                                attrs: { data: _vm.dataObj }
                              })
                            : _vm._e()
                        ],
                        1
                      )
                    ])
                  ])
                ],
            _vm._v(" "),
            _c("div", { staticClass: "row" }, [
              _c("div", { staticClass: "label" }, [
                _vm._v(
                  _vm._s(
                    _vm.type == "dataset"
                      ? _vm.$t("datasetObj.dataset_acess")
                      : _vm.$t("modelManage.modelAccess")
                  ) + "：\n        "
                )
              ]),
              _vm._v(" "),
              _vm.dataObj.is_private
                ? _c(
                    "div",
                    { staticClass: "value", staticStyle: { height: "21px" } },
                    [
                      _c("span", { staticClass: "model-private" }, [
                        _c(
                          "svg",
                          {
                            attrs: {
                              xmlns: "http://www.w3.org/2000/svg",
                              viewBox: "0 0 32 32",
                              width: "16",
                              height: "16"
                            }
                          },
                          [
                            _c("defs"),
                            _vm._v(" "),
                            _c("g", [
                              _c("path", {
                                attrs: {
                                  d:
                                    "M25.333 13.333h1.333c0.736 0 1.333 0.597 1.333 1.333v0 13.333c0 0.736-0.597 1.333-1.333 1.333v0h-21.333c-0.736 0-1.333-0.597-1.333-1.333v0-13.333c0-0.736 0.597-1.333 1.333-1.333v0h1.333v-1.333c0-5.155 4.179-9.333 9.333-9.333s9.333 4.179 9.333 9.333v0 1.333zM6.667 16v10.667h18.667v-10.667h-18.667zM14.667 18.667h2.667v5.333h-2.667v-5.333zM22.667 13.333v-1.333c0-3.682-2.985-6.667-6.667-6.667s-6.667 2.985-6.667 6.667v0 1.333h13.333z"
                                }
                              })
                            ])
                          ]
                        ),
                        _vm._v(
                          "\n            " +
                            _vm._s(_vm.formatAccess(_vm.dataObj.is_private)) +
                            "\n          "
                        )
                      ])
                    ]
                  )
                : _c("div", { staticClass: "value" }, [
                    _vm._v(
                      _vm._s(_vm.formatAccess(_vm.dataObj.is_private)) +
                        "\n        "
                    )
                  ])
            ]),
            _vm._v(" "),
            _c("div", { staticClass: "row" }, [
              _c("div", { staticClass: "label" }, [
                _vm._v(
                  _vm._s(
                    _vm.type == "dataset"
                      ? _vm.$t("datasetObj.dataset_size")
                      : _vm.$t("modelManage.modelSize")
                  ) + "："
                )
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "value" }, [
                _vm._v(_vm._s(_vm.formatSize(_vm.dataObj.size)))
              ])
            ]),
            _vm._v(" "),
            _c("div", { staticClass: "row" }, [
              _c("div", { staticClass: "label" }, [
                _vm._v(_vm._s(_vm.$t("modelManage.createTime")) + "：")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "value" }, [
                _vm._v(_vm._s(_vm.formatTime(_vm.dataObj.created_unix)))
              ])
            ]),
            _vm._v(" "),
            _c("div", { staticClass: "row" }, [
              _c("div", { staticClass: "label" }, [
                _vm._v(_vm._s(_vm.$t("modelManage.updateTime")) + "：")
              ]),
              _vm._v(" "),
              _c("div", { staticClass: "value" }, [
                _vm._v(_vm._s(_vm.formatTime(_vm.dataObj.updated_unix)))
              ])
            ])
          ],
          2
        ),
        _vm._v(" "),
        _vm.repoUseTaskList.length
          ? [
              _c("div", { staticClass: "divider-column-vertical" }),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "summary" },
                [
                  _c("div", { staticClass: "title" }, [
                    _vm._v(_vm._s(_vm.$t("modelManage.trainUsedRepo")) + "：")
                  ]),
                  _vm._v(" "),
                  _vm._l(_vm.repoUseTaskList, function(item, index) {
                    return _c(
                      "div",
                      {
                        key: index,
                        staticClass: "detail-address",
                        staticStyle: { "margin-left": "1rem" }
                      },
                      [
                        _c("li", { staticClass: "nowrap" }, [
                          _c(
                            "a",
                            { attrs: { href: item.url, title: item.showName } },
                            [_vm._v(_vm._s(item.showName))]
                          )
                        ])
                      ]
                    )
                  })
                ],
                2
              )
            ]
          : _vm._e(),
        _vm._v(" "),
        _vm.trainUsedDataList.length
          ? [
              _c("div", { staticClass: "divider-column-vertical" }),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "summary" },
                [
                  _c("div", { staticClass: "title" }, [
                    _vm._v(
                      _vm._s(_vm.$t("modelManage.trainUsedDataList")) + "："
                    )
                  ]),
                  _vm._v(" "),
                  _vm._l(_vm.trainUsedDataList, function(item) {
                    return _c(
                      "div",
                      {
                        key: item.name,
                        staticClass: "detail-address",
                        staticStyle: { "margin-left": "1rem" }
                      },
                      [
                        !item.is_delete
                          ? _c("span", { staticStyle: { display: "flex" } }, [
                              _c("i", {
                                staticClass: "ri-stack-line",
                                staticStyle: { "margin-right": "0.5rem" }
                              }),
                              _vm._v(" "),
                              _c(
                                "a",
                                {
                                  staticClass: "nowrap",
                                  attrs: {
                                    href: item.url,
                                    title: item.showName
                                  }
                                },
                                [_vm._v(_vm._s(item.showName))]
                              )
                            ])
                          : _c(
                              "span",
                              {
                                staticClass: "nowrap",
                                staticStyle: {
                                  display: "flex",
                                  color: "#888888"
                                }
                              },
                              [
                                _c("i", {
                                  staticClass: "ri-stack-line",
                                  staticStyle: { "margin-right": "0.5rem" }
                                }),
                                _vm._v(" "),
                                _c(
                                  "span",
                                  {
                                    staticClass: "nowrap",
                                    staticStyle: { width: "70%" }
                                  },
                                  [_vm._v(_vm._s(item.showName))]
                                ),
                                _vm._v(" "),
                                _c("span", [
                                  _vm._v(_vm._s(_vm.$t("modelManage.deleted")))
                                ])
                              ]
                            )
                      ]
                    )
                  })
                ],
                2
              )
            ]
          : _vm._e()
      ],
      2
    )
  ])
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "icon-c" }, [
      _c("div", { staticClass: "empty-icon" })
    ])
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=template&id=a96f77d8&scoped=true":
/*!***************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/intro/MigrateModelSync.vue?vue&type=template&id=a96f77d8&scoped=true ***!
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
    [
      _c(
        "a",
        {
          staticClass: "operate-btn",
          on: {
            click: function($event) {
              $event.stopPropagation()
              $event.preventDefault()
              _vm.dlgShow = true
            }
          }
        },
        [
          _c("i", { staticClass: "ri-restart-line" }),
          _vm._v(_vm._s(_vm.$t("modelObj.sync_now")) + "\n  ")
        ]
      ),
      _vm._v(" "),
      _c(
        "BaseDialog",
        {
          staticClass: "export-model-dlg base-dlg",
          attrs: {
            visible: _vm.dlgShow,
            title: _vm.data.name + " " + _vm.$t("modelObj.model_sync"),
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
          _c(
            "div",
            {
              directives: [
                {
                  name: "loading",
                  rawName: "v-loading",
                  value: _vm.loading || _vm.submitLoading,
                  expression: "loading || submitLoading"
                }
              ],
              staticClass: "dlg-content"
            },
            [
              !_vm.loading && _vm.hasUpdate
                ? _c("div", { staticClass: "compare-files" }, [
                    _c(
                      "div",
                      { staticClass: "table-container" },
                      [
                        _c(
                          "el-table",
                          {
                            ref: "tableRef",
                            staticStyle: { width: "100%", height: "100%" },
                            attrs: {
                              data: _vm.filesList,
                              height: 400,
                              stripe: ""
                            }
                          },
                          [
                            _c("el-table-column", {
                              attrs: {
                                prop: "FileName",
                                label: _vm.$t("modelManage.fileName"),
                                align: "left",
                                "header-align": "left"
                              },
                              scopedSlots: _vm._u(
                                [
                                  {
                                    key: "default",
                                    fn: function(scope) {
                                      return [
                                        _c(
                                          "div",
                                          { staticClass: "fitted file-name" },
                                          [
                                            _c("i", {
                                              staticClass:
                                                "icon file alternate outline",
                                              attrs: {
                                                width: "16",
                                                height: "16",
                                                "aria-hidden": "true"
                                              }
                                            }),
                                            _vm._v(" "),
                                            _c(
                                              "span",
                                              {
                                                class:
                                                  scope.row.Status == 1
                                                    ? "color-add"
                                                    : scope.row.Status == 2
                                                    ? "color-del"
                                                    : scope.row.Status == 3
                                                    ? "color-update"
                                                    : "",
                                                attrs: {
                                                  title: scope.row.FileName
                                                }
                                              },
                                              [
                                                _vm._v(
                                                  _vm._s(scope.row.FileName)
                                                )
                                              ]
                                            )
                                          ]
                                        )
                                      ]
                                    }
                                  }
                                ],
                                null,
                                false,
                                3730353096
                              )
                            }),
                            _vm._v(" "),
                            _c("el-table-column", {
                              attrs: {
                                prop: "SizeShow",
                                label: _vm.$t("modelObj.current_size"),
                                align: "center",
                                "header-align": "center",
                                width: "200"
                              }
                            }),
                            _vm._v(" "),
                            _c("el-table-column", {
                              attrs: {
                                prop: "NewSizeShow",
                                label: _vm.$t("modelObj.new_size"),
                                align: "center",
                                "header-align": "center",
                                width: "200"
                              },
                              scopedSlots: _vm._u(
                                [
                                  {
                                    key: "default",
                                    fn: function(scope) {
                                      return [
                                        _c(
                                          "span",
                                          {
                                            class:
                                              scope.row.Status == 1
                                                ? "color-add"
                                                : scope.row.Status == 2
                                                ? "color-del"
                                                : scope.row.Status == 3
                                                ? "color-update"
                                                : ""
                                          },
                                          [
                                            _vm._v(
                                              _vm._s(scope.row.NewSizeShow)
                                            )
                                          ]
                                        )
                                      ]
                                    }
                                  }
                                ],
                                null,
                                false,
                                2191672782
                              )
                            }),
                            _vm._v(" "),
                            _c("el-table-column", {
                              attrs: {
                                prop: "Status",
                                label: _vm.$t("status"),
                                align: "center",
                                "header-align": "center",
                                width: "200"
                              },
                              scopedSlots: _vm._u(
                                [
                                  {
                                    key: "default",
                                    fn: function(scope) {
                                      return [
                                        scope.row.Status == 0
                                          ? _c("span")
                                          : _vm._e(),
                                        _vm._v(" "),
                                        scope.row.Status == 1
                                          ? _c(
                                              "span",
                                              { staticClass: "color-add" },
                                              [
                                                _vm._v(
                                                  _vm._s(
                                                    _vm.$t(
                                                      "modelObj.status_add"
                                                    )
                                                  )
                                                )
                                              ]
                                            )
                                          : _vm._e(),
                                        _vm._v(" "),
                                        scope.row.Status == 2
                                          ? _c(
                                              "span",
                                              { staticClass: "color-del" },
                                              [
                                                _vm._v(
                                                  _vm._s(
                                                    _vm.$t(
                                                      "modelObj.status_del"
                                                    )
                                                  )
                                                )
                                              ]
                                            )
                                          : _vm._e(),
                                        _vm._v(" "),
                                        scope.row.Status == 3
                                          ? _c(
                                              "span",
                                              { staticClass: "color-update" },
                                              [
                                                _vm._v(
                                                  _vm._s(
                                                    _vm.$t(
                                                      "modelObj.status_update"
                                                    )
                                                  )
                                                )
                                              ]
                                            )
                                          : _vm._e()
                                      ]
                                    }
                                  }
                                ],
                                null,
                                false,
                                273032725
                              )
                            })
                          ],
                          1
                        )
                      ],
                      1
                    ),
                    _vm._v(" "),
                    _c("div", { staticClass: "op-btns-c" }, [
                      _c("div", [
                        _vm.canUpdate && _vm.gated
                          ? _c(
                              "div",
                              { staticClass: "hf-token-c" },
                              [
                                _c("span", { staticClass: "title-requred" }, [
                                  _vm._v(
                                    _vm._s(_vm.$t("modelManage.externalToken"))
                                  )
                                ]),
                                _vm._v(" "),
                                _c("el-input", {
                                  attrs: {
                                    placeholder: _vm.$t(
                                      "modelManage.modelLabelGatedTips"
                                    )
                                  },
                                  model: {
                                    value: _vm.hf_token,
                                    callback: function($$v) {
                                      _vm.hf_token = $$v
                                    },
                                    expression: "hf_token"
                                  }
                                })
                              ],
                              1
                            )
                          : _vm._e()
                      ]),
                      _vm._v(" "),
                      _c(
                        "div",
                        { staticClass: "btns-c" },
                        [
                          _vm.canUpdate
                            ? _c(
                                "el-button",
                                {
                                  staticClass: "submit-btn",
                                  attrs: { type: "primary", size: "default" },
                                  on: { click: _vm.update }
                                },
                                [_vm._v(_vm._s(_vm.$t("modelObj.start_sync")))]
                              )
                            : _vm._e(),
                          _vm._v(" "),
                          _c(
                            "el-button",
                            {
                              staticClass: "cancel-btn",
                              attrs: { size: "default" },
                              on: { click: _vm.cancel }
                            },
                            [_vm._v(_vm._s(_vm.$t("cancel")))]
                          )
                        ],
                        1
                      )
                    ])
                  ])
                : _vm._e(),
              _vm._v(" "),
              !_vm.loading && !_vm.hasUpdate
                ? _c("div", { staticClass: "no-update" }, [
                    _c("div", { staticClass: "content" }, [
                      _c("div", { staticClass: "icon-c" }, [
                        _c("i", { staticClass: "ri-information-line" })
                      ]),
                      _vm._v(" "),
                      _c("div", { staticClass: "txt" }, [
                        _vm._v(_vm._s(_vm.$t("modelObj.model_no_update_tips")))
                      ])
                    ]),
                    _vm._v(" "),
                    _c(
                      "div",
                      { staticClass: "op-btns-c" },
                      [
                        _c("div"),
                        _vm._v(" "),
                        _c(
                          "el-button",
                          {
                            staticClass: "cancel-btn",
                            attrs: { size: "default" },
                            on: { click: _vm.cancel }
                          },
                          [
                            _vm._v(
                              _vm._s(_vm.$t("cloudbrainObj.dialogTips.tips8"))
                            )
                          ]
                        )
                      ],
                      1
                    )
                  ])
                : _vm._e()
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

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=template&id=9b01554a&scoped=true":
/*!*******************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Access.vue?vue&type=template&id=9b01554a&scoped=true ***!
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
  return _c("div", { staticClass: "form-container" }, [
    _c(
      "div",
      {
        directives: [
          {
            name: "show",
            rawName: "v-show",
            value: _vm.showMessage,
            expression: "showMessage"
          }
        ],
        staticClass: "ui message",
        class: _vm.successFlag ? "success" : "negative",
        staticStyle: { "margin-top": "10px" }
      },
      [_c("p", [_vm._v(_vm._s(_vm.showMessageText))])]
    ),
    _vm._v(" "),
    _c(
      "div",
      [
        _c(
          "BaseTitle",
          {
            directives: [
              {
                name: "loading",
                rawName: "v-loading",
                value: _vm.loading,
                expression: "loading"
              }
            ],
            staticClass: "content-collar",
            attrs: { title: _vm.title }
          },
          [
            _c("div", { staticClass: "main-box" }, [
              _vm.canChangeDatasetTeams
                ? _c("div", { staticClass: "top-op-area" }, [
                    _c(
                      "div",
                      {
                        staticClass: "ui search focus",
                        attrs: { id: "search-user-box" }
                      },
                      [
                        _c(
                          "form",
                          {
                            ref: "myForm",
                            on: {
                              submit: function($event) {
                                $event.preventDefault()
                              }
                            }
                          },
                          [
                            _c("div", { staticClass: "ui input" }, [
                              _c("input", {
                                directives: [
                                  {
                                    name: "model",
                                    rawName: "v-model",
                                    value: _vm.collaborator,
                                    expression: "collaborator"
                                  }
                                ],
                                ref: "myInput",
                                staticClass: "prompt",
                                attrs: {
                                  placeholder: _vm.placeholder,
                                  required: ""
                                },
                                domProps: { value: _vm.collaborator },
                                on: {
                                  keydown: function($event) {
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
                                    $event.preventDefault()
                                    return _vm.sumbitCollaborate(true)
                                  },
                                  input: [
                                    function($event) {
                                      if ($event.target.composing) {
                                        return
                                      }
                                      _vm.collaborator = $event.target.value
                                    },
                                    _vm.handlSearch
                                  ]
                                }
                              })
                            ])
                          ]
                        ),
                        _vm._v(" "),
                        _vm.searching
                          ? _c(
                              "div",
                              { staticClass: "results transition visible" },
                              [
                                _c("a", { staticClass: "result" }, [
                                  _c("div", { staticClass: "content" }, [
                                    _c("div", { staticClass: "title" }, [
                                      _vm._v(
                                        _vm._s(_vm.$t("datasetObj.dataSearch"))
                                      )
                                    ])
                                  ])
                                ])
                              ]
                            )
                          : _vm.showResults
                          ? _c(
                              "div",
                              { staticClass: "results transition visible" },
                              _vm._l(
                                _vm.searchResults.length > 7
                                  ? _vm.searchResults.slice(0, 7)
                                  : _vm.searchResults,
                                function(user) {
                                  return _c(
                                    "a",
                                    {
                                      key: user.id,
                                      staticClass: "result",
                                      on: {
                                        click: function($event) {
                                          return _vm.selectUser(user)
                                        }
                                      }
                                    },
                                    [
                                      _vm.type === "Collaborators"
                                        ? _c("div", { staticClass: "image" }, [
                                            _c("img", {
                                              attrs: {
                                                src: user.avatar_url,
                                                alt: ""
                                              }
                                            })
                                          ])
                                        : _vm._e(),
                                      _vm._v(" "),
                                      _c("div", { staticClass: "content" }, [
                                        _c(
                                          "div",
                                          { staticClass: "title" },
                                          [
                                            _vm.type === "Teams"
                                              ? [
                                                  _vm._v(
                                                    "\n                        " +
                                                      _vm._s(user.name) +
                                                      " (" +
                                                      _vm._s(
                                                        _vm.filesType ===
                                                          "dataset"
                                                          ? user.dataset_permission
                                                          : user.aimodel_permission
                                                      ) +
                                                      " access)\n                      "
                                                  )
                                                ]
                                              : [
                                                  _vm._v(
                                                    "\n                        " +
                                                      _vm._s(
                                                        user.full_name
                                                          ? user.username +
                                                              " (" +
                                                              user.full_name +
                                                              ")"
                                                          : user.username
                                                      ) +
                                                      "\n                      "
                                                  )
                                                ]
                                          ],
                                          2
                                        )
                                      ])
                                    ]
                                  )
                                }
                              ),
                              0
                            )
                          : _vm.showEmpty
                          ? _c(
                              "div",
                              { staticClass: "results transition visible" },
                              [
                                _c("a", { staticClass: "result" }, [
                                  _c("div", { staticClass: "content" }, [
                                    _c("div", { staticClass: "title" }, [
                                      _vm._v(
                                        _vm._s(
                                          _vm.$t("datasetObj.notFoundUser")
                                        )
                                      )
                                    ])
                                  ])
                                ])
                              ]
                            )
                          : _vm._e()
                      ]
                    ),
                    _vm._v(" "),
                    _c(
                      "button",
                      {
                        staticClass: "ui button add-user",
                        on: { click: _vm.sumbitCollaborate }
                      },
                      [_vm._v(_vm._s(_vm.btnText))]
                    )
                  ])
                : _vm._e(),
              _vm._v(" "),
              _c(
                "div",
                { staticClass: "middle-list-area" },
                _vm._l(_vm.usersList[_vm.type], function(item) {
                  return _c(
                    "div",
                    { staticClass: "list-wrap" },
                    [
                      _vm.type === "Teams"
                        ? [
                            _c("div", { staticClass: "name-img" }, [
                              _c(
                                "a",
                                {
                                  attrs: {
                                    href:
                                      "/org/" +
                                      item.OwnerName +
                                      "/teams/" +
                                      item.Name
                                  }
                                },
                                [_vm._v(_vm._s(item.Name))]
                              )
                            ]),
                            _vm._v(" "),
                            _c("div", { staticClass: "access-i" }, [
                              _c("i", {
                                staticClass: "ri-shield-keyhole-fill"
                              }),
                              _vm._v(" "),
                              _c("span", [
                                _vm._v(
                                  _vm._s(
                                    _vm.filesType === "dataset"
                                      ? _vm.permissionFilter(
                                          item.DatasetAuthorize
                                        )
                                      : _vm.permissionFilter(
                                          item.AimodelAuthorize
                                        )
                                  )
                                )
                              ])
                            ]),
                            _vm._v(" "),
                            _vm.canChangeDatasetTeams
                              ? _c(
                                  "div",
                                  { staticClass: "btn-group" },
                                  [
                                    _c(
                                      "el-button",
                                      {
                                        style: {
                                          backgroundColor: !item.canDelete
                                            ? ""
                                            : "#ff2525"
                                        },
                                        attrs: {
                                          type: "danger",
                                          disabled: !item.canDelete,
                                          round: ""
                                        },
                                        on: {
                                          click: function($event) {
                                            return _vm.deleteImage(item.ID)
                                          }
                                        }
                                      },
                                      [
                                        _vm._v(
                                          "\n                    " +
                                            _vm._s(_vm.$t("delete")) +
                                            "\n                  "
                                        )
                                      ]
                                    )
                                  ],
                                  1
                                )
                              : _vm._e()
                          ]
                        : _vm._e(),
                      _vm._v(" "),
                      _vm.type === "Collaborators"
                        ? [
                            _c("div", { staticClass: "name-img" }, [
                              _c("img", {
                                attrs: { src: item.User.RelAvatarLink }
                              }),
                              _vm._v(" "),
                              _c(
                                "a",
                                { attrs: { href: "/" + item.User.Name } },
                                [
                                  _vm._v(
                                    _vm._s(item.User.FullName || item.User.Name)
                                  )
                                ]
                              )
                            ]),
                            _vm._v(" "),
                            _c(
                              "div",
                              { staticClass: "access-i" },
                              [
                                _c("i", {
                                  staticClass: "ri-shield-keyhole-fill"
                                }),
                                _vm._v(" "),
                                _c(
                                  "el-dropdown",
                                  {
                                    attrs: { size: "medium", trigger: "click" },
                                    on: { command: _vm.handleCommand }
                                  },
                                  [
                                    _c(
                                      "span",
                                      { staticClass: "el-dropdown-link" },
                                      [
                                        _vm._v(
                                          "\n                          " +
                                            _vm._s(
                                              _vm.permissionFilter(
                                                item.Collaboration.Mode
                                              )
                                            )
                                        ),
                                        _c("i", {
                                          staticClass:
                                            "el-icon-arrow-down el-icon--right"
                                        })
                                      ]
                                    ),
                                    _vm._v(" "),
                                    _c(
                                      "el-dropdown-menu",
                                      {
                                        attrs: { slot: "dropdown" },
                                        slot: "dropdown"
                                      },
                                      [
                                        _c(
                                          "el-dropdown-item",
                                          {
                                            attrs: {
                                              disabled:
                                                item.Collaboration.Mode === 3,
                                              command: {
                                                mode: 3,
                                                id: item.User.ID
                                              }
                                            }
                                          },
                                          [
                                            _vm._v(
                                              _vm._s(
                                                _vm.$t(
                                                  "datasetObj.administrators"
                                                )
                                              )
                                            )
                                          ]
                                        ),
                                        _vm._v(" "),
                                        _c(
                                          "el-dropdown-item",
                                          {
                                            attrs: {
                                              disabled:
                                                item.Collaboration.Mode === 2,
                                              command: {
                                                mode: 2,
                                                id: item.User.ID
                                              }
                                            }
                                          },
                                          [
                                            _vm._v(
                                              _vm._s(
                                                _vm.$t(
                                                  "datasetObj.writePermission"
                                                )
                                              )
                                            )
                                          ]
                                        ),
                                        _vm._v(" "),
                                        _c(
                                          "el-dropdown-item",
                                          {
                                            attrs: {
                                              disabled:
                                                item.Collaboration.Mode === 1,
                                              command: {
                                                mode: 1,
                                                id: item.User.ID
                                              }
                                            }
                                          },
                                          [
                                            _vm._v(
                                              _vm._s(
                                                _vm.$t(
                                                  "datasetObj.readPermission"
                                                )
                                              )
                                            )
                                          ]
                                        )
                                      ],
                                      1
                                    )
                                  ],
                                  1
                                )
                              ],
                              1
                            ),
                            _vm._v(" "),
                            _c(
                              "div",
                              { staticClass: "btn-group" },
                              [
                                _c(
                                  "el-button",
                                  {
                                    staticStyle: {
                                      "background-color": "#ff2525"
                                    },
                                    attrs: { type: "danger", round: "" },
                                    on: {
                                      click: function($event) {
                                        return _vm.deleteImage(item.User.ID)
                                      }
                                    }
                                  },
                                  [_vm._v(_vm._s(_vm.$t("delete")))]
                                )
                              ],
                              1
                            )
                          ]
                        : _vm._e()
                    ],
                    2
                  )
                }),
                0
              )
            ])
          ]
        )
      ],
      1
    ),
    _vm._v(" "),
    _c(
      "div",
      {
        staticClass: "ui small basic delete modal",
        attrs: { id: _vm.type + "-" + _vm.id }
      },
      [
        _c("div", { staticClass: "ui icon header" }, [
          _c("i", { staticClass: "trash icon" }),
          _vm._v(
            " " +
              _vm._s(_vm.$t("datasetObj.deletCollaboratedTitle")) +
              "\n      "
          )
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "content" }, [
          _c("p", [
            _vm._v(
              _vm._s(
                _vm.filesType === "dataset"
                  ? _vm.$t("datasetObj.deletCollaboratedTips")
                  : _vm.$t("modelManage.deletCollaboratedTips")
              )
            )
          ])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "actions" }, [
          _c("div", { staticClass: "ui red basic inverted cancel button" }, [
            _c("i", { staticClass: "remove icon" }),
            _vm._v("\n          " + _vm._s(_vm.$t("cancelOp")) + "\n        ")
          ]),
          _vm._v(" "),
          _c("div", { staticClass: "ui green basic inverted ok button" }, [
            _c("i", { staticClass: "checkmark icon" }),
            _vm._v("\n          " + _vm._s(_vm.$t("confirmOp")) + "\n        ")
          ])
        ])
      ]
    )
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=template&id=5368752e&scoped=true":
/*!********************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/setting/Setting.vue?vue&type=template&id=5368752e&scoped=true ***!
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
    { staticClass: "ui container setting-wrap" },
    [
      _c(
        "BaseTitle",
        {
          directives: [
            {
              name: "loading",
              rawName: "v-loading",
              value: _vm.baseLoading,
              expression: "baseLoading"
            }
          ],
          attrs: { title: _vm.$t("datasetObj.dataBaseInfo") }
        },
        [
          _c(
            "el-form",
            {
              ref: "form",
              staticClass: "form-box",
              attrs: { size: "medium", model: _vm.form, rules: _vm.rules }
            },
            [
              _c(
                "el-form-item",
                {
                  staticStyle: { "margin-bottom": "0" },
                  attrs: {
                    label: _vm.$t("datasetObj.dataset_name"),
                    prop: "name"
                  }
                },
                [
                  _c("el-input", {
                    attrs: { maxlength: "100" },
                    model: {
                      value: _vm.form.name,
                      callback: function($$v) {
                        _vm.$set(_vm.form, "name", $$v)
                      },
                      expression: "form.name"
                    }
                  }),
                  _vm._v(" "),
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
                        "\n                    " +
                          _vm._s(_vm.$t("datasetObj.dataset_name_tooltips")) +
                          "\n                "
                      )
                    ]
                  )
                ],
                1
              ),
              _vm._v(" "),
              _c(
                "el-form-item",
                {
                  staticStyle: { "margin-bottom": "0" },
                  attrs: {
                    label: _vm.$t("datasetObj.dataset_zh_name"),
                    prop: "alias"
                  }
                },
                [
                  _c("el-input", {
                    attrs: { maxlength: "100" },
                    model: {
                      value: _vm.form.alias,
                      callback: function($$v) {
                        _vm.$set(_vm.form, "alias", $$v)
                      },
                      expression: "form.alias"
                    }
                  }),
                  _vm._v(" "),
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
                        "\n                    " +
                          _vm._s(
                            _vm.$t("datasetObj.dataset_name_cn_tooltips")
                          ) +
                          "\n                "
                      )
                    ]
                  )
                ],
                1
              ),
              _vm._v(" "),
              _c(
                "el-form-item",
                {
                  attrs: {
                    label: _vm.$t("datasetObj.dataset_owner"),
                    prop: "owner_id"
                  }
                },
                [
                  _c(
                    "el-select",
                    {
                      attrs: { disabled: "" },
                      model: {
                        value: _vm.owners[0].name,
                        callback: function($$v) {
                          _vm.$set(_vm.owners[0], "name", $$v)
                        },
                        expression: "owners[0].name"
                      }
                    },
                    _vm._l(_vm.owners, function(item) {
                      return _c("el-option", {
                        key: item.name,
                        attrs: { value: item.name, label: item.name }
                      })
                    }),
                    1
                  )
                ],
                1
              ),
              _vm._v(" "),
              _c(
                "el-form-item",
                { attrs: { label: "", prop: "is_private" } },
                [
                  _c(
                    "el-radio-group",
                    {
                      model: {
                        value: _vm.form.is_private,
                        callback: function($$v) {
                          _vm.$set(_vm.form, "is_private", $$v)
                        },
                        expression: "form.is_private"
                      }
                    },
                    [
                      _c("el-radio", { attrs: { label: false } }, [
                        _vm._v(_vm._s(_vm.$t("modelManage.modelAccessPublic")))
                      ]),
                      _vm._v(" "),
                      _c("el-radio", { attrs: { label: true } }, [
                        _vm._v(_vm._s(_vm.$t("modelManage.modelAccessPrivate")))
                      ])
                    ],
                    1
                  )
                ],
                1
              ),
              _vm._v(" "),
              _c(
                "el-form-item",
                {
                  attrs: { label: _vm.$t("datasetObj.category"), prop: "tags" }
                },
                [
                  _c(
                    "el-select",
                    {
                      attrs: {
                        placeholder: _vm.$t("datasetObj.select_category"),
                        filterable: ""
                      },
                      model: {
                        value: _vm.form.tags[0],
                        callback: function($$v) {
                          _vm.$set(_vm.form.tags, 0, $$v)
                        },
                        expression: "form.tags[0]"
                      }
                    },
                    _vm._l(_vm.Category, function(item) {
                      return _c("el-option", {
                        key: item.name,
                        attrs: {
                          value: item.name,
                          label: _vm.$t("datasets." + item.name)
                        }
                      })
                    }),
                    1
                  )
                ],
                1
              ),
              _vm._v(" "),
              _c(
                "el-form-item",
                { attrs: { label: _vm.$t("datasetObj.task"), prop: "tasks" } },
                [
                  _c(
                    "el-select",
                    {
                      attrs: {
                        placeholder: _vm.$t("datasetObj.select_task"),
                        filterable: ""
                      },
                      model: {
                        value: _vm.form.tasks[0],
                        callback: function($$v) {
                          _vm.$set(_vm.form.tasks, 0, $$v)
                        },
                        expression: "form.tasks[0]"
                      }
                    },
                    _vm._l(_vm.Task, function(item) {
                      return _c("el-option", {
                        key: item.name,
                        attrs: {
                          value: item.name,
                          label: _vm.$t("datasets." + item.name)
                        }
                      })
                    }),
                    1
                  )
                ],
                1
              ),
              _vm._v(" "),
              _c(
                "el-form-item",
                {
                  attrs: {
                    label: _vm.$t("datasetObj.license"),
                    prop: "licenses"
                  }
                },
                [
                  _c(
                    "el-select",
                    {
                      attrs: {
                        placeholder: _vm.$t("datasetObj.license_helper"),
                        filterable: ""
                      },
                      model: {
                        value: _vm.form.licenses,
                        callback: function($$v) {
                          _vm.$set(_vm.form, "licenses", $$v)
                        },
                        expression: "form.licenses"
                      }
                    },
                    _vm._l(_vm.License, function(item) {
                      return _c("el-option", {
                        key: item.name,
                        attrs: { value: item.name, label: item.name }
                      })
                    }),
                    1
                  )
                ],
                1
              ),
              _vm._v(" "),
              _c(
                "el-form-item",
                { staticClass: "btn-wrap" },
                [
                  _c(
                    "el-button",
                    {
                      staticClass: "submit",
                      staticStyle: { background: "rgb(0, 102, 255)" },
                      attrs: { round: "", type: "primary" },
                      on: {
                        click: function($event) {
                          return _vm.editBaseInfo()
                        }
                      }
                    },
                    [_vm._v(_vm._s(_vm.$t("confirm")))]
                  )
                ],
                1
              )
            ],
            1
          )
        ],
        1
      ),
      _vm._v(" "),
      _c("Access", {
        attrs: {
          id: _vm.dataObj.id,
          title: _vm.$t("datasetObj.datasetCollaborator"),
          placeholder: _vm.$t("datasetObj.searchUsers"),
          btnText: _vm.$t("datasetObj.addCollaborators"),
          type: "Collaborators"
        }
      }),
      _vm._v(" "),
      _vm.dataObj.Owner.IsOrganization
        ? _c("Access", {
            attrs: {
              id: _vm.dataObj.id,
              title: _vm.$t("datasetObj.managementTeam"),
              placeholder: _vm.$t("datasetObj.searchTeams"),
              btnText: _vm.$t("datasetObj.addTeam"),
              type: "Teams",
              org: _vm.dataObj.Owner.Name
            }
          })
        : _vm._e(),
      _vm._v(" "),
      _vm.dataObj.can_delete
        ? _c(
            "BaseTitle",
            {
              staticClass: "content-del",
              attrs: { title: _vm.$t("datasetObj.dangerousOperationZone") }
            },
            [
              _c(
                "div",
                { staticClass: "main-box" },
                [
                  _c("div", { staticClass: "text-wrap" }, [
                    _c("p", { staticStyle: { "font-weight": "700" } }, [
                      _vm._v(_vm._s(_vm.$t("datasetObj.deleteThisDataset")))
                    ]),
                    _vm._v(" "),
                    _c("p", [
                      _vm._v(_vm._s(_vm.$t("datasetObj.deleteThisDatasetTips")))
                    ])
                  ]),
                  _vm._v(" "),
                  _c(
                    "el-button",
                    {
                      staticClass: "del-btn",
                      attrs: { size: "medium", round: "", plain: "" },
                      on: { click: _vm.showModal }
                    },
                    [_vm._v(_vm._s(_vm.$t("datasetObj.deleteThisDataset")))]
                  )
                ],
                1
              )
            ]
          )
        : _vm._e(),
      _vm._v(" "),
      _c(
        "div",
        { staticClass: "ui small modal", attrs: { id: "dataset-del" } },
        [
          _c("div", { staticClass: "header" }, [
            _vm._v(
              "\n\t\t\t" +
                _vm._s(_vm.$t("datasetObj.deleteThisDataset")) +
                "\n\t\t"
            )
          ]),
          _vm._v(" "),
          _c(
            "div",
            {
              directives: [
                {
                  name: "loading",
                  rawName: "v-loading",
                  value: _vm.delLoading,
                  expression: "delLoading"
                }
              ],
              staticClass: "content"
            },
            [
              _c(
                "div",
                { staticClass: "ui warning message text left" },
                [
                  _c("p", {
                    domProps: {
                      innerHTML: _vm._s(_vm.$t("datasetObj.delete_notices_1"))
                    }
                  }),
                  _vm._v(" "),
                  _c("P", {
                    domProps: {
                      innerHTML: _vm._s(
                        _vm.$t("datasetObj.delete_notices_2", {
                          dataset: _vm.dataObj.alias
                        })
                      )
                    }
                  })
                ],
                1
              ),
              _vm._v(" "),
              _c(
                "form",
                {
                  staticClass: "ui form",
                  on: {
                    submit: function($event) {
                      $event.preventDefault()
                    }
                  }
                },
                [
                  _c("div", { staticClass: "field" }, [
                    _c("label", [
                      _vm._v(
                        "\n\t\t\t\t\t\t" +
                          _vm._s(_vm.$t("datasetObj.sureDatasetName")) +
                          "\n\t\t\t\t\t\t"
                      ),
                      _c("span", { staticClass: "text red" }, [
                        _vm._v(_vm._s(_vm.dataObj.alias))
                      ])
                    ])
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "required field" }, [
                    _c("label", { attrs: { for: "dataset_name" } }, [
                      _vm._v(_vm._s(_vm.$t("datasetObj.dataset_name1")))
                    ]),
                    _vm._v(" "),
                    _c("input", {
                      directives: [
                        {
                          name: "model",
                          rawName: "v-model",
                          value: _vm.name,
                          expression: "name"
                        }
                      ],
                      attrs: {
                        id: "dataset_name",
                        name: "dataset_name",
                        required: ""
                      },
                      domProps: { value: _vm.name },
                      on: {
                        input: function($event) {
                          if ($event.target.composing) {
                            return
                          }
                          _vm.name = $event.target.value
                        }
                      }
                    })
                  ]),
                  _vm._v(" "),
                  _c("div", { staticClass: "text right actions" }, [
                    _c("div", { staticClass: "ui cancel button" }, [
                      _vm._v(_vm._s(_vm.$t("cancel")))
                    ]),
                    _vm._v(" "),
                    _c(
                      "button",
                      {
                        staticClass: "ui red button",
                        on: { click: _vm.deleteDataset }
                      },
                      [_vm._v(_vm._s(_vm.$t("confirm1")))]
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
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=template&id=3e16625e&scoped=true":
/*!*********************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/components/square/detail/usage/UsageIntro.vue?vue&type=template&id=3e16625e&scoped=true ***!
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
    { staticClass: "ui container usage-intro" },
    [
      _c(
        "BaseTitle",
        {
          staticClass: "content",
          attrs: {
            title:
              _vm.type === "dataset"
                ? _vm.$t("datasetObj.codeUseDlgTitle")
                : _vm.$t("modelObj.codeUseDlgTitle")
          }
        },
        [
          _c(
            "div",
            { staticClass: "main-box" },
            [
              _c("el-skeleton", {
                staticStyle: { padding: "0 20px" },
                attrs: { rows: 10, loading: _vm.loading, animated: "" }
              }),
              _vm._v(" "),
              _c("div", {
                staticClass: "markdown",
                domProps: { innerHTML: _vm._s(_vm.content) }
              })
            ],
            1
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

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=template&id=2cb6c8b3&scoped=true":
/*!******************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/dataset/detail/index.vue?vue&type=template&id=2cb6c8b3&scoped=true ***!
  \******************************************************************************************************************************************************************************************************************************************/
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
    _vm.emptyPage
      ? _c(
          "div",
          { staticStyle: { "padding-top": "50px" } },
          [_c("NotFound")],
          1
        )
      : _c(
          "div",
          [
            JSON.stringify(_vm.datasetObj) !== "{}"
              ? [
                  _c("Header", {
                    staticClass: "datset-header",
                    attrs: {
                      tab: _vm.tab,
                      tabList: _vm.tabList,
                      dataObj: _vm.datasetObj
                    },
                    on: { changeTab: _vm.changeTab }
                  }),
                  _vm._v(" "),
                  _vm.tab == "intro"
                    ? _c("Intro", { attrs: { dataObj: _vm.datasetObj } })
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tab == "files"
                    ? _c("FileList", {
                        attrs: { dataObj: _vm.datasetObj },
                        on: { editSuccess: _vm.editSuccess }
                      })
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tab == "usage"
                    ? _c("UsageIntro", { attrs: { dataObj: _vm.datasetObj } })
                    : _vm._e(),
                  _vm._v(" "),
                  _vm.tab == "settings"
                    ? _c("Setting", {
                        attrs: { dataObj: _vm.datasetObj },
                        on: { editSuccess: _vm.editSuccess }
                      })
                    : _vm._e()
                ]
              : _vm._e()
          ],
          2
        )
  ])
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=template&id=e8b9b4e4&scoped=true":
/*!*************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/CreateForm.vue?vue&type=template&id=e8b9b4e4&scoped=true ***!
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
  return _c(
    "div",
    { staticClass: "ui container" },
    [
      _vm.showSteps
        ? _c(
            "el-steps",
            {
              staticClass: "steps",
              attrs: { active: _vm.formactive, "finish-status": "finish" }
            },
            [
              _c("el-step", { attrs: { title: _vm.title[0] } }),
              _vm._v(" "),
              _c("el-step", { attrs: { title: _vm.title[1] } })
            ],
            1
          )
        : _vm._e(),
      _vm._v(" "),
      _c("div", { staticClass: "form-container" }, [
        _c("div", { staticClass: "form-head" }, [
          _c("h4", [_vm._v(_vm._s(_vm.formTtitle))])
        ]),
        _vm._v(" "),
        _c("div", { staticClass: "form-content" }, [_vm._t("default")], 2)
      ])
    ],
    1
  )
}
var staticRenderFns = []
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=template&id=70071dcb&scoped=true":
/*!*************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/guide/components/FileUpload.vue?vue&type=template&id=70071dcb&scoped=true ***!
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
  return _c("div", { staticClass: "ui form" }, [
    _c("div", { staticClass: "row-c" }, [
      _c("div", { staticClass: "row" }, [
        _c("div", { staticClass: "r-title" }, [
          _c("label", { staticClass: "required" }, [
            _vm._v(_vm._s(_vm.$t("modelManage.uploadPath")))
          ])
        ]),
        _vm._v(" "),
        _c(
          "div",
          { staticClass: "r-content" },
          [
            _c("el-input", {
              staticClass: "input-disabled",
              attrs: {
                size: "medium",
                value: _vm.uploadDir || "/",
                readonly: ""
              }
            })
          ],
          1
        )
      ]),
      _vm._v(" "),
      _c("div", { staticClass: "row" }, [
        _c("div", { staticClass: "r-title" }),
        _vm._v(" "),
        _c(
          "div",
          { staticClass: "r-content" },
          [
            _c(
              "el-radio-group",
              {
                staticClass: "upload-type-sel",
                attrs: { disabled: _vm.uploading },
                on: { change: _vm.changeUploadType },
                model: {
                  value: _vm.uploadType,
                  callback: function($$v) {
                    _vm.uploadType = $$v
                  },
                  expression: "uploadType"
                }
              },
              [
                _c("el-radio", { attrs: { label: "file" } }, [
                  _vm._v(_vm._s(_vm.$t("modelManage.uploadFile")))
                ]),
                _vm._v(" "),
                _c("el-radio", { attrs: { label: "folder" } }, [
                  _vm._v(_vm._s(_vm.$t("modelManage.uploadFolder")))
                ])
              ],
              1
            )
          ],
          1
        )
      ]),
      _vm._v(" "),
      _vm.storageShow
        ? _c("div", { staticClass: "row" }, [
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
                    " " +
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
                  _vm._v(_vm._s(_vm.formattedTotal[1]) + "\n              "),
                  _vm.isStorageExceeded
                    ? _c("div", { staticClass: "limit-tip-wrap" }, [
                        _c("div", { staticClass: "limit-tip" }, [
                          _c("i", { staticClass: "ri-information-line" }),
                          _vm._v(_vm._s(_vm.$t("storage.exceedStorage")))
                        ]),
                        _vm._v(" "),
                        !_vm.showOwenerTips
                          ? _c("a", { attrs: { href: "/storages" } }, [
                              _vm._v(_vm._s(_vm.$t("storage.capacity_details")))
                            ])
                          : _c("span", [
                              _vm._v(
                                _vm._s(
                                  _vm.$t("storage.owenerTips", {
                                    ownerName: _vm.ownerName
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
          ])
        : _vm._e(),
      _vm._v(" "),
      _c(
        "div",
        { staticClass: "row", staticStyle: { "align-items": "flex-start" } },
        [
          _c("div", { staticClass: "r-title" }, [
            _vm.uploadType == "file"
              ? _c("label", { staticClass: "required" }, [
                  _vm._v(_vm._s(_vm.$t("modelManage.fileUpload")))
                ])
              : _vm._e(),
            _vm._v(" "),
            _vm.uploadType == "folder"
              ? _c("label", { staticClass: "required" }, [
                  _vm._v(_vm._s(_vm.$t("modelManage.folderUpload")))
                ])
              : _vm._e()
          ]),
          _vm._v(" "),
          _c(
            "div",
            { staticClass: "r-content" },
            [
              _c(
                "div",
                {
                  directives: [
                    {
                      name: "show",
                      rawName: "v-show",
                      value: _vm.uploadType == "file",
                      expression: "uploadType == 'file'"
                    }
                  ],
                  staticStyle: { position: "relative" }
                },
                [
                  _c("form", { ref: "dropzoneRef", staticClass: "dropzone" }, [
                    _c("div", {
                      directives: [
                        {
                          name: "show",
                          rawName: "v-show",
                          value: _vm.showUploadErr,
                          expression: "showUploadErr"
                        }
                      ],
                      staticClass: "dropzon-err-tips ui red message",
                      staticStyle: { display: "none", margin: "2.5rem" },
                      domProps: { innerHTML: _vm._s(_vm.uploadErrTxt) }
                    })
                  ]),
                  _vm._v(" "),
                  _c("div", {
                    staticClass: "tips",
                    domProps: { innerHTML: _vm._s(_vm.getDefaultErrTxt()) }
                  }),
                  _vm._v(" "),
                  _c("div", {
                    directives: [
                      {
                        name: "show",
                        rawName: "v-show",
                        value: _vm.uploading,
                        expression: "uploading"
                      }
                    ],
                    staticClass: "not-allowed-placeholder"
                  })
                ]
              ),
              _vm._v(" "),
              _vm.uploadType == "folder"
                ? _c("FolderUploadSelect", {
                    ref: "folderUploadSelect",
                    attrs: {
                      maxFilesSize: _vm.maxModelFilesSize,
                      uploading: _vm.uploading
                    },
                    on: { folderSelectChange: _vm.folderSelectChange }
                  })
                : _vm._e()
            ],
            1
          )
        ]
      ),
      _vm._v(" "),
      _c("div", { staticClass: "row", staticStyle: { "margin-top": "10px" } }, [
        _vm._m(0),
        _vm._v(" "),
        _c(
          "div",
          { staticClass: "r-content" },
          [
            _c(
              "el-button",
              {
                staticClass: "green",
                attrs: {
                  size: "medium",
                  round: "",
                  disabled:
                    _vm.uploading || _vm.btnFlag || _vm.isStorageExceeded
                },
                on: { click: _vm.submit }
              },
              [_vm._v(_vm._s(_vm.$t("modelManage.upload")))]
            ),
            _vm._v(" "),
            _c(
              "el-button",
              {
                attrs: {
                  size: "medium",
                  round: "",
                  type: "info",
                  disabled: _vm.uploading
                },
                on: { click: _vm.cancel }
              },
              [_vm._v(_vm._s(_vm.$t("modelManage.cancel")))]
            )
          ],
          1
        )
      ]),
      _vm._v(" "),
      _c(
        "div",
        { staticClass: "row", staticStyle: { "align-items": "flex-start" } },
        [
          _c("div", { staticClass: "r-title" }, [
            _c("label", [
              _vm._v(_vm._s(_vm.$t("modelManage.uploadStatus")) + "：")
            ])
          ]),
          _vm._v(" "),
          _c(
            "div",
            { staticClass: "r-content" },
            _vm._l(_vm.uploadFiles, function(item, index) {
              return _c(
                "div",
                {
                  key: item.upload.uuid,
                  staticClass: "datast-upload-progress"
                },
                [
                  _c(
                    "span",
                    {
                      staticClass: "dataset-name nowrap",
                      attrs: { title: item.fullname }
                    },
                    [_vm._v(_vm._s(item.fullname))]
                  ),
                  _vm._v(" "),
                  _c(
                    "div",
                    { staticClass: "dataset-progress" },
                    [
                      _c("el-progress", {
                        attrs: {
                          "text-inside": true,
                          "stroke-width": 14,
                          percentage: _vm.uploadStatusList[index].progress
                        }
                      })
                    ],
                    1
                  ),
                  _vm._v(" "),
                  _c("div", { staticClass: "dataset-status nowrap" }, [
                    _c(
                      "div",
                      { staticClass: "status-flex" },
                      [
                        _vm.uploadStatusList[index].infoCode === 1 ||
                        _vm.uploadStatusList[index].infoCode === 2
                          ? _c("i", {
                              staticClass: "ri-close-circle-line failed"
                            })
                          : _vm._e(),
                        _vm._v(" "),
                        _vm.uploadStatusList[index].infoCode === 0
                          ? _c("i", {
                              staticClass: "ri-checkbox-circle-line success"
                            })
                          : _vm._e(),
                        _vm._v(" "),
                        _c("span", [
                          _vm._v(_vm._s(_vm.uploadStatusList[index].status))
                        ]),
                        _vm._v(" "),
                        _vm.uploadStatusList[index].infoCode === 1
                          ? _c(
                              "el-tooltip",
                              {
                                staticClass: "item",
                                attrs: { effect: "dark", placement: "top" }
                              },
                              [
                                _c(
                                  "div",
                                  {
                                    attrs: { slot: "content" },
                                    slot: "content"
                                  },
                                  [
                                    _vm._v(
                                      " " +
                                        _vm._s(
                                          _vm.uploadStatusList[index].failedInfo
                                        ) +
                                        " "
                                    )
                                  ]
                                ),
                                _vm._v(" "),
                                _c("i", {
                                  staticClass: "ri-question-fill",
                                  staticStyle: {
                                    "font-size": "16px",
                                    "margin-left": "0.5rem",
                                    cursor: "pointer"
                                  }
                                })
                              ]
                            )
                          : _vm._e()
                      ],
                      1
                    )
                  ])
                ]
              )
            }),
            0
          )
        ]
      )
    ])
  ])
}
var staticRenderFns = [
  function() {
    var _vm = this
    var _h = _vm.$createElement
    var _c = _vm._self._c || _h
    return _c("div", { staticClass: "r-title" }, [_c("label")])
  }
]
render._withStripped = true



/***/ }),

/***/ "./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=template&id=1a87d32a&scoped=true":
/*!***************************************************************************************************************************************************************************************************************************************************************!*\
  !*** ./node_modules/vue-loader/lib/loaders/templateLoader.js??vue-loader-options!./node_modules/vue-loader/lib/index.js??vue-loader-options!./web_src/vuepages/pages/modelmanage/components/FolderUploadSelect.vue?vue&type=template&id=1a87d32a&scoped=true ***!
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
  return _c("div", { staticClass: "folder-upload-select-c" }, [
    _c("div", { staticClass: "folder-upload-select" }, [
      _c(
        "div",
        { staticClass: "op-c" },
        [
          _c(
            "el-button",
            {
              attrs: {
                size: "mini",
                icon: "el-icon-circle-plus",
                disabled: _vm.uploading
              },
              on: { click: _vm.addFolder }
            },
            [_vm._v(_vm._s(_vm.$t("modelManage.addUploadFolder")))]
          ),
          _vm._v(" "),
          _c(
            "div",
            { staticClass: "op-right" },
            [
              _c("div", { staticClass: "summary" }, [
                _vm._v(
                  _vm._s(
                    _vm.$t("modelManage.fileCountAndSize", {
                      count: _vm.fileList.length,
                      size: _vm.allFilesSize
                    })
                  ) + "\n        "
                )
              ]),
              _vm._v(" "),
              _c(
                "el-button",
                {
                  staticStyle: { color: "red" },
                  attrs: {
                    size: "mini",
                    type: "text",
                    disabled: _vm.uploading
                  },
                  on: {
                    click: function($event) {
                      return _vm.remove()
                    }
                  }
                },
                [_vm._v(_vm._s(_vm.$t("modelManage.clearAll")))]
              )
            ],
            1
          )
        ],
        1
      ),
      _vm._v(" "),
      _c("div", { staticClass: "err-tips-c" }, [
        _c("div", {
          directives: [
            {
              name: "show",
              rawName: "v-show",
              value: _vm.errorState,
              expression: "errorState"
            }
          ],
          staticClass: "err-tips ui red message",
          domProps: { innerHTML: _vm._s(_vm.errorInfo) }
        })
      ]),
      _vm._v(" "),
      _c(
        "div",
        {
          staticClass: "file-list-c",
          on: { dragover: _vm.handleDragover, drop: _vm.handleDrop }
        },
        [
          _vm._l(_vm.fileList, function(file, index) {
            return _c(
              "div",
              {
                key: index,
                staticClass: "file-item",
                on: {
                  click: function($event) {
                    $event.stopPropagation()
                    $event.preventDefault()
                  }
                }
              },
              [
                _c(
                  "div",
                  {
                    staticClass: "file-name",
                    style: { color: file.sizeStatus == "error" ? "red" : "" }
                  },
                  [_vm._v(_vm._s(file.fullname))]
                ),
                _vm._v(" "),
                _c("div", { staticClass: "file-size" }, [
                  _vm._v(_vm._s(file.file_size))
                ]),
                _vm._v(" "),
                _c(
                  "div",
                  { staticClass: "file-status" },
                  [
                    file.status == "error"
                      ? _c(
                          "el-tooltip",
                          {
                            attrs: {
                              effect: "dark",
                              content: file.errTips,
                              placement: "top"
                            }
                          },
                          [
                            _c("i", {
                              staticClass: "el-icon-warning-outline",
                              staticStyle: { color: "red" }
                            })
                          ]
                        )
                      : _vm._e(),
                    _vm._v(" "),
                    file.status == "success"
                      ? _c("i", {
                          staticClass: "el-icon-circle-check",
                          staticStyle: { color: "#21ba45" }
                        })
                      : _vm._e()
                  ],
                  1
                ),
                _vm._v(" "),
                _c("div", { staticClass: "file-op" }, [
                  _c("i", {
                    staticClass: "el-icon-circle-close",
                    attrs: { disabled: _vm.uploading },
                    on: {
                      click: function($event) {
                        $event.stopPropagation()
                        $event.preventDefault()
                        return _vm.remove(file)
                      }
                    }
                  })
                ])
              ]
            )
          }),
          _vm._v(" "),
          !_vm.fileList.length
            ? _c("div", {
                staticClass: "empty-place-holder",
                domProps: {
                  innerHTML: _vm._s(
                    _vm.$t("modelManage.folderUploadPlaceholder")
                  )
                },
                on: { click: _vm.addFolder }
              })
            : _vm._e()
        ],
        2
      ),
      _vm._v(" "),
      _c("input", {
        ref: "filepicker",
        staticStyle: { display: "none" },
        attrs: {
          type: "file",
          webkitdirectory: "",
          directory: "",
          multiple: ""
        },
        on: { change: _vm.folderSelectChange }
      })
    ]),
    _vm._v(" "),
    _c("div", {
      staticClass: "tips",
      domProps: { innerHTML: _vm._s(_vm.getDefaultErrTxt()) }
    })
  ])
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
/******/ 	/* webpack/runtime/amd options */
/******/ 	!function() {
/******/ 		__webpack_require__.amdO = {};
/******/ 	}();
/******/ 	
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
/******/ 	/* webpack/runtime/ensure chunk */
/******/ 	!function() {
/******/ 		__webpack_require__.f = {};
/******/ 		// This file contains only the entry chunk.
/******/ 		// The chunk loading function for additional chunks
/******/ 		__webpack_require__.e = function(chunkId) {
/******/ 			return Promise.all(Object.keys(__webpack_require__.f).reduce(function(promises, key) {
/******/ 				__webpack_require__.f[key](chunkId, promises);
/******/ 				return promises;
/******/ 			}, []));
/******/ 		};
/******/ 	}();
/******/ 	
/******/ 	/* webpack/runtime/get javascript chunk filename */
/******/ 	!function() {
/******/ 		// This function allow to reference async chunks
/******/ 		__webpack_require__.u = function(chunkId) {
/******/ 			// return url for filenames based on template
/******/ 			return "js/" + chunkId + ".js";
/******/ 		};
/******/ 	}();
/******/ 	
/******/ 	/* webpack/runtime/get mini-css chunk filename */
/******/ 	!function() {
/******/ 		// This function allow to reference async chunks
/******/ 		__webpack_require__.miniCssF = function(chunkId) {
/******/ 			// return url for filenames based on template
/******/ 			return "css/" + chunkId + ".css";
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
/******/ 	/* webpack/runtime/load script */
/******/ 	!function() {
/******/ 		var inProgress = {};
/******/ 		// data-webpack is not used as build has no uniqueName
/******/ 		// loadScript function to load a script via script tag
/******/ 		__webpack_require__.l = function(url, done, key, chunkId) {
/******/ 			if(inProgress[url]) { inProgress[url].push(done); return; }
/******/ 			var script, needAttach;
/******/ 			if(key !== undefined) {
/******/ 				var scripts = document.getElementsByTagName("script");
/******/ 				for(var i = 0; i < scripts.length; i++) {
/******/ 					var s = scripts[i];
/******/ 					if(s.getAttribute("src") == url) { script = s; break; }
/******/ 				}
/******/ 			}
/******/ 			if(!script) {
/******/ 				needAttach = true;
/******/ 				script = document.createElement('script');
/******/ 		
/******/ 				script.charset = 'utf-8';
/******/ 				script.timeout = 120;
/******/ 				if (__webpack_require__.nc) {
/******/ 					script.setAttribute("nonce", __webpack_require__.nc);
/******/ 				}
/******/ 		
/******/ 		
/******/ 				script.src = url;
/******/ 			}
/******/ 			inProgress[url] = [done];
/******/ 			var onScriptComplete = function(prev, event) {
/******/ 				// avoid mem leaks in IE.
/******/ 				script.onerror = script.onload = null;
/******/ 				clearTimeout(timeout);
/******/ 				var doneFns = inProgress[url];
/******/ 				delete inProgress[url];
/******/ 				script.parentNode && script.parentNode.removeChild(script);
/******/ 				doneFns && doneFns.forEach(function(fn) { return fn(event); });
/******/ 				if(prev) return prev(event);
/******/ 			}
/******/ 			var timeout = setTimeout(onScriptComplete.bind(null, undefined, { type: 'timeout', target: script }), 120000);
/******/ 			script.onerror = onScriptComplete.bind(null, script.onerror);
/******/ 			script.onload = onScriptComplete.bind(null, script.onload);
/******/ 			needAttach && document.head.appendChild(script);
/******/ 		};
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
/******/ 	/* webpack/runtime/css loading */
/******/ 	!function() {
/******/ 		if (typeof document === "undefined") return;
/******/ 		var createStylesheet = function(chunkId, fullhref, oldTag, resolve, reject) {
/******/ 			var linkTag = document.createElement("link");
/******/ 		
/******/ 			linkTag.rel = "stylesheet";
/******/ 			linkTag.type = "text/css";
/******/ 			if (__webpack_require__.nc) {
/******/ 				linkTag.nonce = __webpack_require__.nc;
/******/ 			}
/******/ 			var onLinkComplete = function(event) {
/******/ 				// avoid mem leaks.
/******/ 				linkTag.onerror = linkTag.onload = null;
/******/ 				if (event.type === 'load') {
/******/ 					resolve();
/******/ 				} else {
/******/ 					var errorType = event && event.type;
/******/ 					var realHref = event && event.target && event.target.href || fullhref;
/******/ 					var err = new Error("Loading CSS chunk " + chunkId + " failed.\n(" + errorType + ": " + realHref + ")");
/******/ 					err.name = "ChunkLoadError";
/******/ 					err.code = "CSS_CHUNK_LOAD_FAILED";
/******/ 					err.type = errorType;
/******/ 					err.request = realHref;
/******/ 					if (linkTag.parentNode) linkTag.parentNode.removeChild(linkTag)
/******/ 					reject(err);
/******/ 				}
/******/ 			}
/******/ 			linkTag.onerror = linkTag.onload = onLinkComplete;
/******/ 			linkTag.href = fullhref;
/******/ 		
/******/ 		
/******/ 			if (oldTag) {
/******/ 				oldTag.parentNode.insertBefore(linkTag, oldTag.nextSibling);
/******/ 			} else {
/******/ 				document.head.appendChild(linkTag);
/******/ 			}
/******/ 			return linkTag;
/******/ 		};
/******/ 		var findStylesheet = function(href, fullhref) {
/******/ 			var existingLinkTags = document.getElementsByTagName("link");
/******/ 			for(var i = 0; i < existingLinkTags.length; i++) {
/******/ 				var tag = existingLinkTags[i];
/******/ 				var dataHref = tag.getAttribute("data-href") || tag.getAttribute("href");
/******/ 				if(tag.rel === "stylesheet" && (dataHref === href || dataHref === fullhref)) return tag;
/******/ 			}
/******/ 			var existingStyleTags = document.getElementsByTagName("style");
/******/ 			for(var i = 0; i < existingStyleTags.length; i++) {
/******/ 				var tag = existingStyleTags[i];
/******/ 				var dataHref = tag.getAttribute("data-href");
/******/ 				if(dataHref === href || dataHref === fullhref) return tag;
/******/ 			}
/******/ 		};
/******/ 		var loadStylesheet = function(chunkId) {
/******/ 			return new Promise(function(resolve, reject) {
/******/ 				var href = __webpack_require__.miniCssF(chunkId);
/******/ 				var fullhref = __webpack_require__.p + href;
/******/ 				if(findStylesheet(href, fullhref)) return resolve();
/******/ 				createStylesheet(chunkId, fullhref, null, resolve, reject);
/******/ 			});
/******/ 		}
/******/ 		// object to store loaded CSS chunks
/******/ 		var installedCssChunks = {
/******/ 			"vp-dataset-detial": 0
/******/ 		};
/******/ 		
/******/ 		__webpack_require__.f.miniCss = function(chunkId, promises) {
/******/ 			var cssChunks = {"monaco-vendors-include-loader_node_modules_monaco-editor_esm_vs_editor_editor_main_js":1};
/******/ 			if(installedCssChunks[chunkId]) promises.push(installedCssChunks[chunkId]);
/******/ 			else if(installedCssChunks[chunkId] !== 0 && cssChunks[chunkId]) {
/******/ 				promises.push(installedCssChunks[chunkId] = loadStylesheet(chunkId).then(function() {
/******/ 					installedCssChunks[chunkId] = 0;
/******/ 				}, function(e) {
/******/ 					delete installedCssChunks[chunkId];
/******/ 					throw e;
/******/ 				}));
/******/ 			}
/******/ 		};
/******/ 		
/******/ 		// no hmr
/******/ 		
/******/ 		// no prefetching
/******/ 		
/******/ 		// no preloaded
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
/******/ 			"vp-dataset-detial": 0
/******/ 		};
/******/ 		
/******/ 		__webpack_require__.f.j = function(chunkId, promises) {
/******/ 				// JSONP chunk loading for javascript
/******/ 				var installedChunkData = __webpack_require__.o(installedChunks, chunkId) ? installedChunks[chunkId] : undefined;
/******/ 				if(installedChunkData !== 0) { // 0 means "already installed".
/******/ 		
/******/ 					// a Promise means "currently loading".
/******/ 					if(installedChunkData) {
/******/ 						promises.push(installedChunkData[2]);
/******/ 					} else {
/******/ 						if(true) { // all chunks have JS
/******/ 							// setup Promise in chunk cache
/******/ 							var promise = new Promise(function(resolve, reject) { installedChunkData = installedChunks[chunkId] = [resolve, reject]; });
/******/ 							promises.push(installedChunkData[2] = promise);
/******/ 		
/******/ 							// start chunk loading
/******/ 							var url = __webpack_require__.p + __webpack_require__.u(chunkId);
/******/ 							// create error before stack unwound to get useful stacktrace later
/******/ 							var error = new Error();
/******/ 							var loadingEnded = function(event) {
/******/ 								if(__webpack_require__.o(installedChunks, chunkId)) {
/******/ 									installedChunkData = installedChunks[chunkId];
/******/ 									if(installedChunkData !== 0) installedChunks[chunkId] = undefined;
/******/ 									if(installedChunkData) {
/******/ 										var errorType = event && (event.type === 'load' ? 'missing' : event.type);
/******/ 										var realSrc = event && event.target && event.target.src;
/******/ 										error.message = 'Loading chunk ' + chunkId + ' failed.\n(' + errorType + ': ' + realSrc + ')';
/******/ 										error.name = 'ChunkLoadError';
/******/ 										error.type = errorType;
/******/ 										error.request = realSrc;
/******/ 										installedChunkData[1](error);
/******/ 									}
/******/ 								}
/******/ 							};
/******/ 							__webpack_require__.l(url, loadingEnded, "chunk-" + chunkId, chunkId);
/******/ 						}
/******/ 					}
/******/ 				}
/******/ 		};
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
/******/ 	var __webpack_exports__ = __webpack_require__.O(undefined, ["vendor-libs"], function() { return __webpack_require__("./web_src/vuepages/pages/dataset/detail/vp-dataset-detial.js"); })
/******/ 	__webpack_exports__ = __webpack_require__.O(__webpack_exports__);
/******/ 	
/******/ })()
;