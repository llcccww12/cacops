(function () {
  let form = document.getElementById("form_id");
  let createFlag = false;
  let flag;
  form.onsubmit = function (e) {
    if (createFlag) return false;
    createFlag = true;
  };
  // $("select.dropdown").dropdown();
  $(document).keydown(function (event) {
    switch (event.keyCode) {
      case 13:
        return false;
    }
  });
  $(".menu .item").tab();
  $(document).ready(createParamter());
  function createParamter() {
    let params = $(".dynamic.field").data("params");
    params &&
      params.parameter.forEach((item, index) => {
        Add_parameter(index, (flag = true), item);
      });
  }
  // 参数增加、删除、修改、保存
  function Add_parameter(i, flag = false, paramsObject = {}) {
    let value = "";
    value += `<div class="two fields width85" id= "para${i}">`;
    value += '<div class="field">';
    let placeholder_value = $(".dynamic.field").data("params-value");
    let placeholder_name = $(".dynamic.field").data("params-name");
    if (flag) {
      value += `<input type="text" class="shipping_first-name"  value="${paramsObject.label}">`;
    } else {
      value +=
        '<input type="text" class="shipping_first-name" required placeholder="' +
        placeholder_name +
        '">';
    }
    value += "</div>";
    value += '<div class="field">';
    if (flag) {
      value += `<input type="text" class="shipping_last-name"  value="${paramsObject.value}">`;
    } else {
      value +=
        '<input type="text" class="shipping_last-name" required placeholder="' +
        placeholder_value +
        '">';
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
          rules: [
            {
              type: "regExp[/.+.py$/g]",
            },
          ],
        },
        job_name: {
          identifier: "job_name",
          rules: [
            {
              type: "regExp[/^[a-z0-9][a-z0-9-_]{1,34}[a-z0-9-]$/]",
            },
          ],
        },
        display_job_name: {
          identifier: "display_job_name",
          rules: [
            {
              type: "regExp[/^[a-z0-9][a-z0-9-_]{1,34}[a-z0-9-]$/]",
            },
          ],
        },
        attachment: {
          identifier: "attachment",
          rules: [
            {
              type: "empty",
            },
          ],
        },
        spec_id: {
          identifier: "spec_id",
          rules: [{ type: "empty" }],
        },
        branch_name: {
          identifier: "branch_name",
          rules: [{ type: "empty" }],
        },
      },
      onSuccess: function () {
        // $('.ui.page.dimmer').dimmer('show')
        
        document.getElementById("mask").style.display = "block";
        isValidate = true;
      },
      onFailure: function (e) {
        isValidate = false;
        createFlag = false;
        return false;
      },
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
    let paraFlag = true;
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
      run_parameters.push({ label: para_name, value: para_value });
    });
    msg["parameter"] = run_parameters;
    msg = JSON.stringify(msg);
    $("#store_run_para").val(msg);
    return paraFlag;
  }
  function get_name() {
    let name1 = $("#engine_name .text").text();
    let name2 = $("#flaver_name .text").text();
    $("input#ai_engine_name").val(name1);
    $("input#ai_flaver_name").val(name2);
    if ($(".cloudbrain_image .text").text()) { 
      $("input[name='image']").val($(".cloudbrain_image .text").text());
    }
  }

  validate();
  $(".ui.create_train_job.green.button").click(function (e) {
    get_name();
    let paramNotValue = send_run_para();
    if (!paramNotValue) {
      return false;
    }
    if (e.target.getAttribute('data-required-model') && !$('input[name="model_name"]').val()) {
      $('input[name="model_name"]').parent().addClass("error")
      return false
    }
    if($('input[name="model_name"]').val() && !$('input[name="ckpt_name"]').val()){
        $('input[name="ckpt_name"]').parent().addClass("error")
        return false
    }
    validate();
  });


  //管理镜像相关的东西
  let nameMap, nameList;
  let RepoLink = $(".cloudbrain-type").data("repo-link");
  let type = $(".cloudbrain-type").data("cloudbrain-type");
  let flagModel = $(".cloudbrain-type").data("flag-model");
  // 获取模型列表和模型名称对应的模型版本
  $(document).ready(function () {
    return; // 改用模型选择组件了
    if (!flagModel) return;
    else {
      $.get(
        `${RepoLink}/modelmanage/query_model_for_predict?type=${type}`,
        (data) => {
          nameMap = data.nameMap;
          nameList = data.nameList;
          let html = `<div class="item"></div>`;
          nameList.forEach((element) => {
            html += `<div class="item" data-value=${element}>${element}</div>`;
          });
          if (nameList.length !== 0) {
            $("#model_name").append(html);
          }
          let faildModelName = $('input[name="model_name"]').val();
          let faildModelVersion = $('input[name="model_version"]').val();
          let dataID;
          // 新建错误的表单返回初始化
          if (faildModelName && nameList.includes(faildModelName)) {
            $("#select_model").dropdown("set text", faildModelName);
            $("#select_model").dropdown("set value", faildModelName);
            nameMap[faildModelName].forEach((element) => {
              if (element.version === faildModelVersion) {
                dataID = element.id;
              }
            });
            initModelVerison(faildModelName, nameMap, faildModelVersion);
            initModelckpt(dataID);
          }
        }
      );
    }

    $("#select_model").dropdown({
      onChange: function (value, text, $selectedItem) {
        $("#model_name_version").empty();
        if (value) {
          $("#select_model").removeClass("error");
          let html = "";
          nameMap[value].forEach((element) => {
            //let { trainTaskInfo } = element;
            //trainTaskInfo = JSON.parse(trainTaskInfo);
            html += `<div class="item" data-label="${element.label}" data-id="${element.id}" data-value="${element.path}">${element.version}</div>`;
          });
          $("#model_name_version").append(html);
          const initVersionText = $(
            "#model_name_version div.item:first-child"
          ).text();
          const initVersionValue = $(
            "#model_name_version div.item:first-child"
          ).data("value");

          $("#select_model_version").dropdown("set text", initVersionText);
          $("#select_model_version").dropdown(
            "set value",
            initVersionValue,
            initVersionText,
            $("#model_name_version div.item:first-child")
          );
        } else {
          $("#select_model_version").dropdown("set text", "");
          $("#select_model_version").dropdown("set value", "");
          $("#select_model_checkpoint").dropdown("set text", "");
          $("#select_model_checkpoint").dropdown("set value", "");
          $("#model_checkpoint").empty();
        }
      },
    });

    $("#select_model_version").dropdown({
      onChange: function (value, text, $selectedItem) {
        if (!value) return;
        const dataID =
          $selectedItem && $selectedItem[0].getAttribute("data-id");
        $("input#ai_model_version").val(text);
        $("#select_model_checkpoint").dropdown("set text", "");
        $("#select_model_checkpoint").addClass("loading");

        $("#model_checkpoint").empty();
        let html = "";
        loadCheckpointList(dataID).then((res) => {
          res.forEach((element) => {
            const ckptSuffix = element.FileName.split(".");
            const loadCheckpointFile = [
              "ckpt",
              "pb",
              "h5",
              "json",
              "pkl",
              "pth",
              "t7",
              "pdparams",
              "onnx",
              "pbtxt",
              "keras",
              "mlmodel",
              "cfg",
              "pt",
            ];
            if (
              !element.IsDir &&
              loadCheckpointFile.includes(ckptSuffix[ckptSuffix.length - 1])
            ) {
              html += `<div class="item" data-value="${element.FileName}">${element.FileName}</div>`;
            }
          });
          $("#model_checkpoint").append(html);
          $("#select_model_checkpoint").removeClass("loading");
          if (html) { 
            $("#select_model_checkpoint").removeClass("error");
          }
          const initVersionText = $(
            "#model_checkpoint div.item:first-child"
          ).text();
          const initVersionValue = $(
            "#model_checkpoint div.item:first-child"
          ).data("value");

          $("#select_model_checkpoint").dropdown("set text", initVersionText);
          $("#select_model_checkpoint").dropdown(
            "set value",
            initVersionValue,
            initVersionText,
            $("#model_name_version div.item:first-child")
          );
        });
      },
    });
  });
  function initModelVerison(value, nameMap, faildModelVersion) {
    let faildTrainUrl = $('input[name="pre_train_model_url"]').val();
    let html = "";
    nameMap[value].forEach((element) => {
      html += `<div class="item" data-label="${element.label}" data-id="${element.id}" data-value="${element.path}">${element.version}</div>`;
    });
    $("#model_name_version").append(html);
    $("#select_model_version").dropdown("set text", faildModelVersion);
    $("#select_model_version").dropdown("set value", faildTrainUrl);
  }
  function initModelckpt(dataID) {
    let faildCkptName = $('input[name="ckpt_name"]').val();
    $("#select_model_checkpoint").addClass("loading");
    $("#model_checkpoint").empty();
    let html = "";
    loadCheckpointList(dataID).then((res) => {
      res.forEach((element) => {
        const ckptSuffix = element.FileName.split(".");
        const loadCheckpointFile = [
          "ckpt",
          "pb",
          "h5",
          "json",
          "pkl",
          "pth",
          "t7",
          "pdparams",
          "onnx",
          "pbtxt",
          "keras",
          "mlmodel",
          "cfg",
          "pt",
        ];
        if (
          !element.IsDir &&
          loadCheckpointFile.includes(ckptSuffix[ckptSuffix.length - 1])
        ) {
          html += `<div class="item" data-value=${element.FileName}>${element.FileName}</div>`;
        }
      });
      $("#model_checkpoint").append(html);
      $("#select_model_checkpoint").removeClass("loading");
      $("#select_model_checkpoint").dropdown("set text", faildCkptName);
      $("#select_model_checkpoint").dropdown("set value", faildCkptName);
    });
  }
  function loadCheckpointList(value) {
    return new Promise((resolve, reject) => {
      $.get(
        `${RepoLink}/modelmanage/query_modelfile_for_predict`,
        { id: value },
        (data) => {
          resolve(data);
        }
      );
    });
  }
  // 评测任务相关创建func
  let repoLink = $(".cloudbrain-type").data("repo-link");
  function setChildType(type_id=1) {
    if (type_id == 3) {
        $('#train_href_id').attr('href', 'https://openi.pcl.ac.cn/CV_benchmark/CV_MOT_benchmark');
        $('#test_href_id').attr('href', 'https://openi.pcl.ac.cn/CV_benchmark/CV_MOT_benchmark');
    } else {
        $('#train_href_id').attr('href', 'https://openi.pcl.ac.cn/CV_benchmark/CV_reID_benchmark');
        $('#test_href_id').attr('href', 'https://openi.pcl.ac.cn/CV_benchmark/CV_reID_benchmark');
    }
    let child_selected_id = $('#benchmark_child_types_id_hidden').val();
    $.get(`${repoLink}/cloudbrain/benchmark/get_child_types?benchmark_type_id=${type_id}`, (data) => {
        const n_length = data['child_types'].length
        let html = ''
        for (let i = 0; i < n_length; i++) {
            if (child_selected_id == data['child_types'][i].id) {
                html += `<option value="${data['child_types'][i].id}" selected="true">${data['child_types'][i].value}</option>`;
            } else {
                html += `<option value="${data['child_types'][i].id}">${data['child_types'][i].value}</option>`;
            }
        }
        let el = document.getElementById("benchmark_child_types_id");
        el && (el.innerHTML = html);
    })
  }
  $(document).ready(function () {
    if ($('input[name=benchmarkMode]').val() === 'alogrithm' || $('input[name=benchmarkMode]').val() === '') {
        setChildType();
    }
    $(".ui.selection.dropdown.benchmark_types_id").dropdown({
      onChange:function (value, text, $selectedItem){
        setChildType(value)
      }
    })
    $('.ui.search.dropdown.job_type').dropdown({
      onChange: function (value, text, $selectedItem) {
          if (value === "BRAINSCORE") {
              $('#brainscore_child_type').css('display', 'block')
              $('#sim2brain_child_type').css('display', 'none')
              $('#benchmark_model_example').attr('href', 'https://openi.pcl.ac.cn/BDIP/similarity2brain_ann')
          }else if(value === "SIM2BRAIN_SNN"){
              $('#sim2brain_child_type').css('display', 'block')
              $('#brainscore_child_type').css('display', 'none')
              $('#benchmark_model_example').attr('href', 'https://openi.pcl.ac.cn/BDIP/sim2brian_snn')
          }else if (value === "SNN4ECOSET") {
              $('#brainscore_child_type').css('display', 'none')
              $('#sim2brain_child_type').css('display', 'none')
              $('#benchmark_model_example').attr('href', 'https://openi.pcl.ac.cn/BDIP/snn4ecoset')
          }else {
              $('#brainscore_child_type').css('display', 'none')
              $('#sim2brain_child_type').css('display', 'none')
              $('#benchmark_model_example').attr('href', 'https://openi.pcl.ac.cn/BDIP/snn4imagenet')
          }
      }
    })
  })
})();
