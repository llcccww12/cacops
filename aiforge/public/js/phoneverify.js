/******/ (function() { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./web_src/js/standalone/phoneverify.js":
/*!**********************************************!*\
  !*** ./web_src/js/standalone/phoneverify.js ***!
  \**********************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.find */ "./node_modules/core-js/modules/es.array.find.js");
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_parse_int__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.parse-int */ "./node_modules/core-js/modules/es.parse-int.js");
/* harmony import */ var core_js_modules_es_parse_int__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_parse_int__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/web.timers */ "./node_modules/core-js/modules/web.timers.js");
/* harmony import */ var core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_timers__WEBPACK_IMPORTED_MODULE_2__);



;

(function () {
  function PhoneVerifyCode(dom, options) {
    if (!dom) return;
    this.countDownNumber = 120;
    this.options = options;
    this.init(dom);
  }

  PhoneVerifyCode.prototype.init = function (dom) {
    if (!dom) return;
    this.dom = $(dom);
    this.isMoving = false;
    this.verifySucess = false;
    this.canSendCode = false;
    this.countDownEnd = false;
    this.imgID = ''; // 初始化触摸冲突处理

    this.initTouchConflictHandler();
    this.eventInit();
    this.refreshImages();
    var wrap = this.dom.closest('div.use-type');
    var readonly = wrap.attr('readonly');

    if (readonly) {
      this.dom.find('.phone-area-c select').attr('disabled', true);
      this.dom.find('.phone-num-c input').attr('disabled', true);
      this.dom.find('.slide-bar-wrap').hide();
      this.dom.find('.verify-code-c').hide();
      this.dom.find('.modify-phone-number').show();
    }

    var oldPhoneNum = wrap.attr('ophonenumber');

    if (oldPhoneNum) {
      this.dom.find('input.phoneNumber').val(oldPhoneNum);
    }

    var showlabel = wrap.attr('showlabel');

    if (showlabel) {
      this.dom.find('._label-c ').show();
    } else {
      this.dom.find('._label-c ').hide();
    }

    var showNewpwd = wrap.attr('shownewpwd');

    if (showNewpwd) {
      this.dom.find('.new-pass-word-wrap').show();
    } else {
      this.dom.find('.new-pass-word-wrap').remove();
    }

    var autofocus = wrap.attr('autofocus');

    if (autofocus) {
      this.dom.find('input.phoneNumber').focus();
    }

    var verifyCodeNoRequired = wrap.attr('verifycodenorequired');

    if (verifyCodeNoRequired) {
      this.dom.find('input.verifyCode').removeAttr('required');
    }
  };

  PhoneVerifyCode.prototype.eventInit = function () {
    if (!this.dom) return;
    var self = this;
    var clientX = 0,
        oLeft = 0,
        imgHideTimer = null;
    this.dom.find('.slide-bar-bg').on('mouseenter touchstart', function (e) {
      if (self.verifySucess) return;
      imgHideTimer && clearTimeout(imgHideTimer);
      self.dom.find('.slide-image-big').slideDown();
    });
    this.dom.find('.slide-bar-bg').on('mouseleave', function (e) {
      imgHideTimer && clearTimeout(imgHideTimer);
      imgHideTimer = setTimeout(function () {
        self.dom.find('.slide-image-big').slideUp();
      }, 200);
    });
    this.dom.find('.slide-image-big').on('mouseenter', function (e) {
      imgHideTimer && clearTimeout(imgHideTimer);
    });
    this.dom.find('.slide-image-big').on('mouseleave', function (e) {
      imgHideTimer && clearTimeout(imgHideTimer);
      imgHideTimer = setTimeout(function () {
        self.dom.find('.slide-image-big').slideUp();
      }, 200);
    });

    function mouseMove(e) {
      var _clientX = e.clientX !== undefined ? e.clientX : e.targetTouches[0].clientX;

      var offset = _clientX - clientX;
      var triggerEl = self.dom.find('.slide-trigger');
      var triggerWidth = triggerEl.width();
      var parentWidth = triggerEl.parent().width();
      var maxLeft = parentWidth - triggerWidth;
      var left = oLeft + offset;
      if (oLeft + offset < 0) left = 0;
      if (oLeft + offset > maxLeft) left = maxLeft;
      triggerEl.css('left', left + 'px');
      self.dom.find('.slide-bar').css('width', left + triggerWidth);
      var imageBigWidth = self.dom.find('.slide-image-big').width();
      var imageSmallWidth = self.dom.find('.slide-image-small').width();
      self.dom.find('.slide-image-small').css('left', left / maxLeft * (imageBigWidth - imageSmallWidth) + 'px');
      self.isMoving = true;
      self.dom.find('.slide-txt').hide();
      imgHideTimer && clearTimeout(imgHideTimer);
    }

    function mouseUp(e) {
      $(document).off('mousemove', mouseMove);
      $(document).off('mouseup', mouseUp);
      $(document).off('touchmove', mouseMove);
      $(document).off('touchend', mouseUp);
      self.isMoving = false;
      $.ajax({
        url: '/verifySlideImage',
        type: 'post',
        dataType: 'json',
        data: {
          slide_id: self.imgID,
          x: parseInt(self.dom.find('.slide-image-small').position().left / self.dom.find('.slide-image-big').attr('scale'))
        },
        success: function success(res) {
          if (res && res.Code === 0) {
            self.verifySucess = true;
            self.canSendCode = true;
            self.dom.find('.slide-bar').addClass('sucess');
            self.dom.find('.slide-trigger').addClass('sucess');
            self.dom.find('.slide-trigger .icon').hide();
            self.dom.find('.slide-trigger .icon.check').show();
            self.dom.find('.slide-image-big').slideUp();
            self.dom.find('.verify-code-send-btn').removeClass('__disabled');
          } else {
            self.dom.find('.slide-bar').addClass('error');
            self.dom.find('.slide-trigger').addClass('error');
            self.dom.find('.slide-trigger .icon').hide();
            self.dom.find('.slide-trigger .icon.close').show();
            setTimeout(function () {
              self.refreshImages();
            }, 300);
          }
        },
        error: function error(err) {
          self.dom.find('.slide-bar').addClass('error');
          self.dom.find('.slide-trigger').addClass('error');
          setTimeout(function () {
            self.refreshImages();
          }, 300);
        }
      });
    }

    function mouseDown(e) {
      if (self.verifySucess) return;
      clientX = e.clientX !== undefined ? e.clientX : e.targetTouches[0].clientX;
      oLeft = $(this).position().left;
      $(document).on('mousemove', mouseMove);
      $(document).on('mouseup', mouseUp);
      $(document).on('touchmove', mouseMove);
      $(document).on('touchend', mouseUp);
    }

    this.dom.find('.slide-trigger').on('mousedown', mouseDown);
    this.dom.find('.slide-trigger').on('touchstart', mouseDown);
    this.dom.find('.verify-code-send-btn').on('click', function () {
      if (!self.canSendCode) return;

      if (self.countDownEnd) {
        self.refreshImages();
        return;
      }

      var phoneNumber = self.dom.find('.phone-num-c input').val();

      if (!/^1[3-9]\d{9}$/.test(phoneNumber)) {
        self.dom.find('.phone-num-c').addClass('error');
        return;
      } else {
        self.dom.find('.phone-num-c').removeClass('error');
        var useType = self.dom.closest('div.use-type').attr('usetype') || 0;
        $.ajax({
          url: '/sendVerifyCode',
          type: 'post',
          dataType: 'json',
          data: {
            phone_number: phoneNumber,
            mode: useType,
            // 0注册，1登录 ，2修改手机号，3找回密码
            slide_id: self.imgID
          },
          success: function success(res) {
            if (res && res.Code === 0) {
              self.countDown();
            } else {
              if ($('.ui.negative.message').length) {
                $('.ui.negative.message').eq(0).show().find('p').text(res.Message);
              } else {
                $('body').toast({
                  message: res.Message,
                  showProgress: 'bottom',
                  showIcon: 'warning circle',
                  "class": 'warning',
                  position: 'top right'
                });
              }

              self.refreshImages();
            }
          },
          error: function error(err) {
            console.log(err);
          }
        });
      }
    });
    this.dom.find('.modify-phone-number a').on('click', function () {
      self.dom.find('.phone-area-c select').attr('disabled', false);
      self.dom.find('.phone-num-c input').attr('disabled', false);
      self.dom.find('.slide-bar-wrap').css('display', 'flex');
      self.dom.find('.verify-code-c').css('display', 'flex');
      self.dom.find('.modify-phone-number').hide();
      self.refreshImages();
    });
  };

  PhoneVerifyCode.prototype.refreshImages = function () {
    this.isMoving = false;
    this.verifySucess = false;
    this.canSendCode = false;
    this.countDownEnd = false;
    this.imgID = '';
    this.dom.find('.slide-bar').removeClass('sucess error').css('width', '30px');
    this.dom.find('.slide-trigger').removeClass('sucess error').css('left', '0px');
    var scale = this.dom.find('.slide-bar-bg').width() / 391;
    this.dom.find('.slide-image-big').css('transform', "scale(".concat(scale, ")")).attr('scale', scale);
    this.dom.find('.slide-trigger .icon').hide();
    this.dom.find('.slide-trigger .icon.arrow').show();
    this.dom.find('.slide-txt').show();
    this.dom.find('.slide-image-small').css('left', '0');
    this.dom.find('.verify-code-send-btn').addClass('__disabled');
    var self = this;
    $.ajax({
      url: '/slideImage',
      type: 'get',
      success: function success(res) {
        if (res && res.Code === 0) {
          self.imgID = res.Message;
          self.dom.find('.slide-image-big').css('background', "url(\"/slideimage/".concat(res.Message, ".png\")"));
          self.dom.find('.slide-image-small').css('background', "url(\"/slideimage/".concat(res.Message, "screenshot.png\")"));
        }
      },
      error: function error(err) {
        console.log(err);
      }
    });
  };

  PhoneVerifyCode.prototype.countDown = function () {
    var self = this;
    var sendBtnEl = this.dom.find('.verify-code-send-btn');
    var count = this.countDownNumber;
    sendBtnEl.addClass('__disabled').text(count + (self.options && self.options.Lang && self.options.Lang.second_resend ? self.options.Lang.second_resend : 'S后重发'));
    this.canSendCode = false;
    this.countDownEnd = false;
    var timer = setInterval(function () {
      count--;
      sendBtnEl.addClass('__disabled').text(count + (self.options && self.options.Lang && self.options.Lang.second_resend ? self.options.Lang.second_resend : 'S后重发'));

      if (count <= 0) {
        sendBtnEl.removeClass('__disabled').text(self.options && self.options.Lang && self.options.Lang.get_verification_code ? self.options.Lang.get_verification_code : '获取验证码');
        clearInterval(timer);
        self.canSendCode = true;
        self.countDownEnd = true;
        self.refreshImages();
      }
    }, 1000);
  };

  PhoneVerifyCode.prototype.initTouchConflictHandler = function () {
    var self = this;
    var sliderTrigger = this.dom.find('.slide-trigger')[0];
    var sliderBg = this.dom.find('.slide-bar-bg')[0];
    if (!sliderTrigger) return; // 最小化影响的CSS方案

    var style = document.createElement('style');
    style.textContent = "\n      /* \u53EA\u5BF9\u6ED1\u5757\u80CC\u666F\u5E94\u7528\uFF0C\u4E0D\u5F71\u54CD\u5185\u90E8\u5143\u7D20 */\n      .slide-bar-bg {\n        touch-action: pan-y;\n      }\n      \n      /* \u6ED1\u5757\u672C\u8EAB\u4E0D\u5E94\u7528touch-action\uFF0C\u907F\u514D\u5F71\u54CD\u529F\u80FD */\n      .slide-trigger {\n        -webkit-touch-callout: none;\n        -webkit-user-select: none;\n        user-select: none;\n      }\n      \n      /* \u786E\u4FDD\u6240\u6709\u56FE\u7247\u76F8\u5173\u5143\u7D20\u53EF\u4EE5\u6B63\u5E38\u4EA4\u4E92 */\n      .slide-image-big *,\n      .slide-image-small * {\n        touch-action: auto !important;\n        pointer-events: auto !important;\n      }\n    ";
    document.head.appendChild(style); // 只在必要的时候阻止默认行为

    var touchStartX = 0;
    var isHorizontalMove = false;
    sliderBg.addEventListener('touchstart', function (e) {
      touchStartX = e.touches[0].clientX;
      isHorizontalMove = false;
    }, {
      passive: true
    });
    sliderBg.addEventListener('touchmove', function (e) {
      if (!e.touches || e.touches.length === 0) return;
      var deltaX = e.touches[0].clientX - touchStartX; // 如果检测到明显的水平滑动，阻止浏览器后退

      if (Math.abs(deltaX) > 10) {
        isHorizontalMove = true;
        e.preventDefault();
      }
    }, {
      passive: false
    });
    sliderBg.addEventListener('touchend', function () {
      // 短暂延迟后重置状态
      setTimeout(function () {
        isHorizontalMove = false;
      }, 100);
    }, {
      passive: true
    });
  };

  window.PhoneVerifyCode = PhoneVerifyCode;
})();

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
/******/ 	/* webpack/runtime/jsonp chunk loading */
/******/ 	!function() {
/******/ 		// no baseURI
/******/ 		
/******/ 		// object to store loaded and loading chunks
/******/ 		// undefined = chunk not loaded, null = chunk preloaded/prefetched
/******/ 		// [resolve, reject, Promise] = chunk loading, 0 = chunk loaded
/******/ 		var installedChunks = {
/******/ 			"phoneverify": 0
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
/******/ 	var __webpack_exports__ = __webpack_require__.O(undefined, ["vendor-libs"], function() { return __webpack_require__("./web_src/js/standalone/phoneverify.js"); })
/******/ 	__webpack_exports__ = __webpack_require__.O(__webpack_exports__);
/******/ 	
/******/ })()
;