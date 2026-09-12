/******/ (function() { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./web_src/js/standalone/cloudbrainNew.js":
/*!************************************************!*\
  !*** ./web_src/js/standalone/cloudbrainNew.js ***!
  \************************************************/
/***/ (function(__unused_webpack_module, __webpack_exports__, __webpack_require__) {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! core-js/modules/es.array.concat */ "./node_modules/core-js/modules/es.array.concat.js");
/* harmony import */ var core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_concat__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! core-js/modules/es.array.find */ "./node_modules/core-js/modules/es.array.find.js");
/* harmony import */ var core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_find__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! core-js/modules/es.array.for-each */ "./node_modules/core-js/modules/es.array.for-each.js");
/* harmony import */ var core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_for_each__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! core-js/modules/es.array.includes */ "./node_modules/core-js/modules/es.array.includes.js");
/* harmony import */ var core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_array_includes__WEBPACK_IMPORTED_MODULE_3__);
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! core-js/modules/es.object.to-string */ "./node_modules/core-js/modules/es.object.to-string.js");
/* harmony import */ var core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_object_to_string__WEBPACK_IMPORTED_MODULE_4__);
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! core-js/modules/es.promise */ "./node_modules/core-js/modules/es.promise.js");
/* harmony import */ var core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_5___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_promise__WEBPACK_IMPORTED_MODULE_5__);
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! core-js/modules/es.regexp.exec */ "./node_modules/core-js/modules/es.regexp.exec.js");
/* harmony import */ var core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_6___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_regexp_exec__WEBPACK_IMPORTED_MODULE_6__);
/* harmony import */ var core_js_modules_es_string_includes__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! core-js/modules/es.string.includes */ "./node_modules/core-js/modules/es.string.includes.js");
/* harmony import */ var core_js_modules_es_string_includes__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_includes__WEBPACK_IMPORTED_MODULE_7__);
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! core-js/modules/es.string.split */ "./node_modules/core-js/modules/es.string.split.js");
/* harmony import */ var core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_8___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_es_string_split__WEBPACK_IMPORTED_MODULE_8__);
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! core-js/modules/web.dom-collections.for-each */ "./node_modules/core-js/modules/web.dom-collections.for-each.js");
/* harmony import */ var core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9___default = /*#__PURE__*/__webpack_require__.n(core_js_modules_web_dom_collections_for_each__WEBPACK_IMPORTED_MODULE_9__);











(function () {
  var form = document.getElementById("form_id");
  var createFlag = false;
  var flag;

  form.onsubmit = function (e) {
    if (createFlag) return false;
    createFlag = true;
  }; // $("select.dropdown").dropdown();


  $(document).keydown(function (event) {
    switch (event.keyCode) {
      case 13:
        return false;
    }
  });
  $(".menu .item").tab();
  $(document).ready(createParamter());

  function createParamter() {
    var params = $(".dynamic.field").data("params");
    params && params.parameter.forEach(function (item, index) {
      Add_parameter(index, flag = true, item);
    });
  } // 参数增加、删除、修改、保存


  function Add_parameter(i) {
    var flag = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : false;
    var paramsObject = arguments.length > 2 && arguments[2] !== undefined ? arguments[2] : {};
    var value = "";
    value += "<div class=\"two fields width85\" id= \"para".concat(i, "\">");
    value += '<div class="field">';
    var placeholder_value = $(".dynamic.field").data("params-value");
    var placeholder_name = $(".dynamic.field").data("params-name");

    if (flag) {
      value += "<input type=\"text\" class=\"shipping_first-name\"  value=\"".concat(paramsObject.label, "\">");
    } else {
      value += '<input type="text" class="shipping_first-name" required placeholder="' + placeholder_name + '">';
    }

    value += "</div>";
    value += '<div class="field">';

    if (flag) {
      value += "<input type=\"text\" class=\"shipping_last-name\"  value=\"".concat(paramsObject.value, "\">");
    } else {
      value += '<input type="text" class="shipping_last-name" required placeholder="' + placeholder_value + '">';
    }

    value += "</div>";
    value += '<span><i class="trash icon"></i></span>';
    value += "</div>";
    $(".dynamic.field").append(value);
  }

  $("#add_run_para").click(function () {
    var len = $(".dynamic.field .two.fields").length;
    Add_parameter(len);
  });
  $(".dynamic.field").on("click", ".trash.icon", function () {
    var index = $(this).parent().parent().index();
    $(this).parent().parent().remove();
    var len = $(".dynamic.field .two.fields").length;
    $(".dynamic.field .two.fields").each(function () {
      var cur_index = $(this).index();
      $(this).attr("id", "para" + cur_index);
    });
  });
  var isValidate = false;

  function validate() {
    $(".ui.form").form({
      on: "blur",
      fields: {
        boot_file: {
          identifier: "boot_file",
          rules: [{
            type: "regExp[/.+.py$/g]"
          }]
        },
        job_name: {
          identifier: "job_name",
          rules: [{
            type: "regExp[/^[a-z0-9][a-z0-9-_]{1,34}[a-z0-9-]$/]"
          }]
        },
        display_job_name: {
          identifier: "display_job_name",
          rules: [{
            type: "regExp[/^[a-z0-9][a-z0-9-_]{1,34}[a-z0-9-]$/]"
          }]
        },
        attachment: {
          identifier: "attachment",
          rules: [{
            type: "empty"
          }]
        },
        spec_id: {
          identifier: "spec_id",
          rules: [{
            type: "empty"
          }]
        },
        branch_name: {
          identifier: "branch_name",
          rules: [{
            type: "empty"
          }]
        }
      },
      onSuccess: function onSuccess() {
        // $('.ui.page.dimmer').dimmer('show')
        document.getElementById("mask").style.display = "block";
        isValidate = true;
      },
      onFailure: function onFailure(e) {
        isValidate = false;
        createFlag = false;
        return false;
      }
    });
  }

  document.onreadystatechange = function () {
    if (document.readyState === "complete") {
      document.getElementById("mask").style.display = "none";
    }
  };

  function send_run_para() {
    var run_parameters = [];
    var msg = {};
    var paraFlag = true;
    $(".dynamic.field .two.fields").each(function () {
      var para_name = $(this).find("input.shipping_first-name").val();
      var para_value = $(this).find("input.shipping_last-name").val();

      if (!para_name) {
        $(this).find("input.shipping_first-name").parent().addClass("error");
        paraFlag = false;
        return;
      } else {
        $(this).find("input.shipping_first-name").parent().removeClass("error");
      }

      if (!para_value) {
        $(this).find("input.shipping_last-name").parent().addClass("error");
        paraFlag = false;
        return;
      } else {
        $(this).find("input.shipping_last-name").parent().removeClass("error");
      }

      run_parameters.push({
        label: para_name,
        value: para_value
      });
    });
    msg["parameter"] = run_parameters;
    msg = JSON.stringify(msg);
    $("#store_run_para").val(msg);
    return paraFlag;
  }

  function get_name() {
    var name1 = $("#engine_name .text").text();
    var name2 = $("#flaver_name .text").text();
    $("input#ai_engine_name").val(name1);
    $("input#ai_flaver_name").val(name2);

    if ($(".cloudbrain_image .text").text()) {
      $("input[name='image']").val($(".cloudbrain_image .text").text());
    }
  }

  validate();
  $(".ui.create_train_job.green.button").click(function (e) {
    get_name();
    var paramNotValue = send_run_para();

    if (!paramNotValue) {
      return false;
    }

    if (e.target.getAttribute('data-required-model') && !$('input[name="model_name"]').val()) {
      $('input[name="model_name"]').parent().addClass("error");
      return false;
    }

    if ($('input[name="model_name"]').val() && !$('input[name="ckpt_name"]').val()) {
      $('input[name="ckpt_name"]').parent().addClass("error");
      return false;
    }

    validate();
  }); //管理镜像相关的东西

  var nameMap, nameList;
  var RepoLink = $(".cloudbrain-type").data("repo-link");
  var type = $(".cloudbrain-type").data("cloudbrain-type");
  var flagModel = $(".cloudbrain-type").data("flag-model"); // 获取模型列表和模型名称对应的模型版本

  $(document).ready(function () {
    return; // 改用模型选择组件了

    if (!flagModel) return;else {
      $.get("".concat(RepoLink, "/modelmanage/query_model_for_predict?type=").concat(type), function (data) {
        nameMap = data.nameMap;
        nameList = data.nameList;
        var html = "<div class=\"item\"></div>";
        nameList.forEach(function (element) {
          html += "<div class=\"item\" data-value=".concat(element, ">").concat(element, "</div>");
        });

        if (nameList.length !== 0) {
          $("#model_name").append(html);
        }

        var faildModelName = $('input[name="model_name"]').val();
        var faildModelVersion = $('input[name="model_version"]').val();
        var dataID; // 新建错误的表单返回初始化

        if (faildModelName && nameList.includes(faildModelName)) {
          $("#select_model").dropdown("set text", faildModelName);
          $("#select_model").dropdown("set value", faildModelName);
          nameMap[faildModelName].forEach(function (element) {
            if (element.version === faildModelVersion) {
              dataID = element.id;
            }
          });
          initModelVerison(faildModelName, nameMap, faildModelVersion);
          initModelckpt(dataID);
        }
      });
    }
    $("#select_model").dropdown({
      onChange: function onChange(value, text, $selectedItem) {
        $("#model_name_version").empty();

        if (value) {
          $("#select_model").removeClass("error");
          var html = "";
          nameMap[value].forEach(function (element) {
            //let { trainTaskInfo } = element;
            //trainTaskInfo = JSON.parse(trainTaskInfo);
            html += "<div class=\"item\" data-label=\"".concat(element.label, "\" data-id=\"").concat(element.id, "\" data-value=\"").concat(element.path, "\">").concat(element.version, "</div>");
          });
          $("#model_name_version").append(html);
          var initVersionText = $("#model_name_version div.item:first-child").text();
          var initVersionValue = $("#model_name_version div.item:first-child").data("value");
          $("#select_model_version").dropdown("set text", initVersionText);
          $("#select_model_version").dropdown("set value", initVersionValue, initVersionText, $("#model_name_version div.item:first-child"));
        } else {
          $("#select_model_version").dropdown("set text", "");
          $("#select_model_version").dropdown("set value", "");
          $("#select_model_checkpoint").dropdown("set text", "");
          $("#select_model_checkpoint").dropdown("set value", "");
          $("#model_checkpoint").empty();
        }
      }
    });
    $("#select_model_version").dropdown({
      onChange: function onChange(value, text, $selectedItem) {
        if (!value) return;
        var dataID = $selectedItem && $selectedItem[0].getAttribute("data-id");
        $("input#ai_model_version").val(text);
        $("#select_model_checkpoint").dropdown("set text", "");
        $("#select_model_checkpoint").addClass("loading");
        $("#model_checkpoint").empty();
        var html = "";
        loadCheckpointList(dataID).then(function (res) {
          res.forEach(function (element) {
            var ckptSuffix = element.FileName.split(".");
            var loadCheckpointFile = ["ckpt", "pb", "h5", "json", "pkl", "pth", "t7", "pdparams", "onnx", "pbtxt", "keras", "mlmodel", "cfg", "pt"];

            if (!element.IsDir && loadCheckpointFile.includes(ckptSuffix[ckptSuffix.length - 1])) {
              html += "<div class=\"item\" data-value=\"".concat(element.FileName, "\">").concat(element.FileName, "</div>");
            }
          });
          $("#model_checkpoint").append(html);
          $("#select_model_checkpoint").removeClass("loading");

          if (html) {
            $("#select_model_checkpoint").removeClass("error");
          }

          var initVersionText = $("#model_checkpoint div.item:first-child").text();
          var initVersionValue = $("#model_checkpoint div.item:first-child").data("value");
          $("#select_model_checkpoint").dropdown("set text", initVersionText);
          $("#select_model_checkpoint").dropdown("set value", initVersionValue, initVersionText, $("#model_name_version div.item:first-child"));
        });
      }
    });
  });

  function initModelVerison(value, nameMap, faildModelVersion) {
    var faildTrainUrl = $('input[name="pre_train_model_url"]').val();
    var html = "";
    nameMap[value].forEach(function (element) {
      html += "<div class=\"item\" data-label=\"".concat(element.label, "\" data-id=\"").concat(element.id, "\" data-value=\"").concat(element.path, "\">").concat(element.version, "</div>");
    });
    $("#model_name_version").append(html);
    $("#select_model_version").dropdown("set text", faildModelVersion);
    $("#select_model_version").dropdown("set value", faildTrainUrl);
  }

  function initModelckpt(dataID) {
    var faildCkptName = $('input[name="ckpt_name"]').val();
    $("#select_model_checkpoint").addClass("loading");
    $("#model_checkpoint").empty();
    var html = "";
    loadCheckpointList(dataID).then(function (res) {
      res.forEach(function (element) {
        var ckptSuffix = element.FileName.split(".");
        var loadCheckpointFile = ["ckpt", "pb", "h5", "json", "pkl", "pth", "t7", "pdparams", "onnx", "pbtxt", "keras", "mlmodel", "cfg", "pt"];

        if (!element.IsDir && loadCheckpointFile.includes(ckptSuffix[ckptSuffix.length - 1])) {
          html += "<div class=\"item\" data-value=".concat(element.FileName, ">").concat(element.FileName, "</div>");
        }
      });
      $("#model_checkpoint").append(html);
      $("#select_model_checkpoint").removeClass("loading");
      $("#select_model_checkpoint").dropdown("set text", faildCkptName);
      $("#select_model_checkpoint").dropdown("set value", faildCkptName);
    });
  }

  function loadCheckpointList(value) {
    return new Promise(function (resolve, reject) {
      $.get("".concat(RepoLink, "/modelmanage/query_modelfile_for_predict"), {
        id: value
      }, function (data) {
        resolve(data);
      });
    });
  } // 评测任务相关创建func


  var repoLink = $(".cloudbrain-type").data("repo-link");

  function setChildType() {
    var type_id = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : 1;

    if (type_id == 3) {
      $('#train_href_id').attr('href', 'https://openi.pcl.ac.cn/CV_benchmark/CV_MOT_benchmark');
      $('#test_href_id').attr('href', 'https://openi.pcl.ac.cn/CV_benchmark/CV_MOT_benchmark');
    } else {
      $('#train_href_id').attr('href', 'https://openi.pcl.ac.cn/CV_benchmark/CV_reID_benchmark');
      $('#test_href_id').attr('href', 'https://openi.pcl.ac.cn/CV_benchmark/CV_reID_benchmark');
    }

    var child_selected_id = $('#benchmark_child_types_id_hidden').val();
    $.get("".concat(repoLink, "/cloudbrain/benchmark/get_child_types?benchmark_type_id=").concat(type_id), function (data) {
      var n_length = data['child_types'].length;
      var html = '';

      for (var i = 0; i < n_length; i++) {
        if (child_selected_id == data['child_types'][i].id) {
          html += "<option value=\"".concat(data['child_types'][i].id, "\" selected=\"true\">").concat(data['child_types'][i].value, "</option>");
        } else {
          html += "<option value=\"".concat(data['child_types'][i].id, "\">").concat(data['child_types'][i].value, "</option>");
        }
      }

      var el = document.getElementById("benchmark_child_types_id");
      el && (el.innerHTML = html);
    });
  }

  $(document).ready(function () {
    if ($('input[name=benchmarkMode]').val() === 'alogrithm' || $('input[name=benchmarkMode]').val() === '') {
      setChildType();
    }

    $(".ui.selection.dropdown.benchmark_types_id").dropdown({
      onChange: function onChange(value, text, $selectedItem) {
        setChildType(value);
      }
    });
    $('.ui.search.dropdown.job_type').dropdown({
      onChange: function onChange(value, text, $selectedItem) {
        if (value === "BRAINSCORE") {
          $('#brainscore_child_type').css('display', 'block');
          $('#sim2brain_child_type').css('display', 'none');
          $('#benchmark_model_example').attr('href', 'https://openi.pcl.ac.cn/BDIP/similarity2brain_ann');
        } else if (value === "SIM2BRAIN_SNN") {
          $('#sim2brain_child_type').css('display', 'block');
          $('#brainscore_child_type').css('display', 'none');
          $('#benchmark_model_example').attr('href', 'https://openi.pcl.ac.cn/BDIP/sim2brian_snn');
        } else if (value === "SNN4ECOSET") {
          $('#brainscore_child_type').css('display', 'none');
          $('#sim2brain_child_type').css('display', 'none');
          $('#benchmark_model_example').attr('href', 'https://openi.pcl.ac.cn/BDIP/snn4ecoset');
        } else {
          $('#brainscore_child_type').css('display', 'none');
          $('#sim2brain_child_type').css('display', 'none');
          $('#benchmark_model_example').attr('href', 'https://openi.pcl.ac.cn/BDIP/snn4imagenet');
        }
      }
    });
  });
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
/******/ 			"cloudbrainNew": 0
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
/******/ 	var __webpack_exports__ = __webpack_require__.O(undefined, ["vendor-libs"], function() { return __webpack_require__("./web_src/js/standalone/cloudbrainNew.js"); })
/******/ 	__webpack_exports__ = __webpack_require__.O(__webpack_exports__);
/******/ 	
/******/ })()
;