var token;
if (isEmpty(token)) {
  var meta = $("meta[name=_uid]");
  if (!isEmpty(meta)) {
    token = meta.attr("content");
  }
}

var html = document.documentElement;
var lang = html.attributes["lang"];
var isZh = true;
if (lang != null && lang.nodeValue == "en-US") {
  isZh = false;
} else {
}
function isEmpty(str) {
  if (typeof str == "undefined" || str == null || str == "") {
    return true;
  }
  return false;
}

function escapeHTML(str) {
  const stayBefore = `<font color='red'>`;
  const stayAfter = `</font>`;
  const strTemp = str.replaceAll(stayBefore, '_FONT_STAY_BEFORE_').replaceAll(stayAfter, '_FONT_STAY_AFTER_')
  return strTemp.replace(/[<>&"']/g, function (match) {
    switch (match) {
      case '<': return '&lt;';
      case '>': return '&gt;';
      case '&': return '&amp;';
      case '"': return '&quot;';
      case "'": return '&#39;';
    }
  }).replaceAll('_FONT_STAY_BEFORE_', stayBefore).replaceAll('_FONT_STAY_AFTER_', stayAfter)
}

var itemType = {
  1: "repository",
  2: "issue",
  3: "user",
  4: "org",
  5: "dataset",
  6: "pr",
  7: "model",
};

var sortBy = {
  10: "default",
  11: "updated_unix.keyword",
  12: "num_watches",
  13: "num_stars",
  14: "num_forks",
  20: "default",
  21: "updated_unix.keyword",
  30: "default",
  31: "name.keyword",
  32: "name.keyword",
  33: "created_unix.keyword",
  34: "created_unix.keyword",
  40: "default",
  41: "name.keyword",
  42: "name.keyword",
  43: "created_unix.keyword",
  44: "created_unix.keyword",
  50: "default",
  51: "download_times",
  60: "default",
  61: "updated_unix.keyword",
  70: "default",
  71: "reference_count",
  72: "download_count",
  73: "created_unix.keyword",
};

var sortAscending = {
  10: "false",
  11: "false",
  12: "false",
  13: "false",
  14: "false",
  20: "false",
  21: "false",
  30: "false",
  31: "true",
  32: "false",
  33: "false",
  34: "true",
  40: "false",
  41: "true",
  42: "false",
  43: "false",
  44: "true",
  50: "false",
  51: "false",
  60: "false",
  61: "false",
  70: "false",
  71: "false",
  72: "false",
  73: "false",
};

var currentPage = 1;
var pageSize = 15;
var currentSearchTableName = "repository";
var currentSearchKeyword = "";
var currentSearchSortBy = "";
var currentSearchAscending = "false";
var OnlySearchLabel = false;
var startIndex = 1;
var endIndex = 5;
var totalPage = 1;
var totalNum = 0;
var privateTotal = 0;

function initPageInfo() {
  currentPage = 1;
  startIndex = 1;
  endIndex = 5;
}

function searchItem(type, sortType) {
  if (OnlySearchLabel) {
    doSearchLabel(
      currentSearchTableName,
      currentSearchKeyword,
      sortBy[sortType],
      sortAscending[sortType]
    );
  } else {
    currentSearchKeyword = document.getElementById("keyword_input").value;
    if (!isEmpty(currentSearchKeyword)) {
      initPageInfo();
      currentSearchTableName = itemType[type];
      currentSearchSortBy = sortBy[sortType];
      currentSearchAscending = sortAscending[sortType];
      OnlySearchLabel = false;
      page(currentPage);
    } else {
      emptySearch();
    }
  }
}

function search() {
  currentSearchKeyword = document.getElementById("keyword_input").value;
  if (!isEmpty(currentSearchKeyword)) {
    currentSearchKeyword = currentSearchKeyword.trim();
  }
  if (!isEmpty(currentSearchKeyword)) {
    doSpcifySearch(
      currentSearchTableName,
      currentSearchKeyword,
      sortBy[10],
      "false"
    );
  } else {
    emptySearch();
  }
}

function emptySearch() {
  initDiv(false);
  initPageInfo();
  $("#searchForm").addClass("hiddenSearch");
  document.getElementById("find_id").innerHTML = getLabel(isZh, "search_empty");
  $("#find_title").html("");
  document.getElementById("sort_type").innerHTML = "";
  document.getElementById("child_search_item").innerHTML = "";
  document.getElementById("page_menu").innerHTML = "";
  $("#repo_total").text("");
  $("#pr_total").text("");
  $("#issue_total").text("");
  $("#dataset_total").text("");
  $("#model_total").text("");
  $("#user_total").text("");
  $("#org_total").text("");
  setActivate(null);
}

function initDiv(isSearchLabel = false) {
  if (isSearchLabel) {
    document.getElementById("search_div").style.display = "none";
    document.getElementById("search_label_div").style.display = "block";
    document.getElementById("dataset_item").style.display = "none";
    document.getElementById("model_item").style.display = "none";
    document.getElementById("issue_item").style.display = "none";
    document.getElementById("pr_item").style.display = "none";
    document.getElementById("user_item").style.display = "none";
    document.getElementById("org_item").style.display = "none";
    document.getElementById("find_id").innerHTML = "";
  } else {
    document.getElementById("search_div").style.display = "block";
    document.getElementById("search_label_div").style.display = "none";
    document.getElementById("dataset_item").style.display = "block";
    document.getElementById("model_item").style.display = "block";
    document.getElementById("issue_item").style.display = "block";
    document.getElementById("pr_item").style.display = "block";
    document.getElementById("user_item").style.display = "block";
    document.getElementById("org_item").style.display = "block";
    document.getElementById("find_id").innerHTML = getLabel(
      isZh,
      "search_finded"
    );
  }
}

function doSpcifySearch(tableName, keyword, sortBy = "", ascending = "false") {
  initDiv(false);
  $("#searchForm").addClass("hiddenSearch");
  document.getElementById("find_id").innerHTML = getLabel(
    isZh,
    "search_finded"
  );
  currentSearchKeyword = keyword;
  initPageInfo();
  currentSearchTableName = tableName;
  currentSearchSortBy = sortBy;
  currentSearchAscending = ascending;
  OnlySearchLabel = false;

  page(currentPage);

  if (currentSearchTableName != "repository") {
    doSearch("repository", currentSearchKeyword, 1, pageSize, true, "", false);
  }
  if (currentSearchTableName != "issue") {
    doSearch("issue", currentSearchKeyword, 1, pageSize, true, "", false);
  }
  if (currentSearchTableName != "user") {
    doSearch("user", currentSearchKeyword, 1, pageSize, true, "", false);
  }
  if (currentSearchTableName != "org") {
    doSearch("org", currentSearchKeyword, 1, pageSize, true, "", false);
  }
  if (currentSearchTableName != "dataset") {
    doSearch("dataset", currentSearchKeyword, 1, pageSize, true, "", false);
  }
  if (currentSearchTableName != "model") {
    doSearch("model", currentSearchKeyword, 1, pageSize, true, "", false);
  }
  if (currentSearchTableName != "pr") {
    doSearch("pr", currentSearchKeyword, 1, pageSize, true, "", false);
  }
}

function doSearchLabel(tableName, keyword, sortBy = "", ascending = "false") {
  initDiv(true);
  //document.getElementById("search_div").style.display="none";
  //document.getElementById("search_label_div").style.display="block";
  document.getElementById("search_label_div").innerHTML =
    '<p class="searchlabel">#' + escapeHTML(keyword) + "</p>";

  currentSearchKeyword = keyword;
  initPageInfo();
  currentSearchTableName = tableName;
  currentSearchSortBy = sortBy;
  currentSearchAscending = ascending;
  OnlySearchLabel = true;

  page(currentPage);
}

function searchLabel(tableName, keyword, sortBy = "", ascending = "false") {
  sessionStorage.setItem("keyword", keyword);
  sessionStorage.setItem("tableName", tableName);
  sessionStorage.setItem("searchLabel", true);
  sessionStorage.setItem("sortBy", sortBy);
  sessionStorage.setItem("ascending", ascending);
  window.open("/all/search/");
}

function doSearch(
  tableName,
  keyword,
  page,
  pageSize = 15,
  onlyReturnNum = true,
  sortBy = "",
  OnlySearchLabel = false
) {
  var language = "zh-CN";
  if (!isZh) {
    language = "en-US";
  }
  $.ajax({
    type: "GET",
    url: "/all/dosearch/",
    headers: {
      authorization: token,
    },
    dataType: "json",
    data: {
      TableName: tableName,
      Key: keyword,
      Page: page,
      PageSize: pageSize,
      OnlyReturnNum: onlyReturnNum,
      SortBy: sortBy,
      OnlySearchLabel: OnlySearchLabel,
      Ascending: currentSearchAscending,
      WebTotal: totalNum,
      PrivateTotal: privateTotal,
      language: language,
    },
    async: true,
    success: function (json) {
      displayResult(tableName, page, json, onlyReturnNum, keyword);
    },
    error: function (response) { },
  });
}

function displayResult(tableName, page, jsonResult, onlyReturnNum, keyword) {
  keyword = escapeHTML(keyword)
  if (tableName == "repository") {
    displayRepoResult(page, jsonResult, onlyReturnNum, keyword);
  } else if (tableName == "issue") {
    displayIssueResult(page, jsonResult, onlyReturnNum, keyword);
  } else if (tableName == "user") {
    displayUserResult(page, jsonResult, onlyReturnNum, keyword);
  } else if (tableName == "org") {
    displayOrgResult(page, jsonResult, onlyReturnNum, keyword);
  } else if (tableName == "dataset") {
    displayDataSetResult(page, jsonResult, onlyReturnNum, keyword);
  } else if (tableName == "model") {
    displayModelResult(page, jsonResult, onlyReturnNum, keyword);
  } else if (tableName == "pr") {
    displayPrResult(page, jsonResult, onlyReturnNum, keyword);
  }
  if (!onlyReturnNum) {
    totalPage = Math.ceil(jsonResult.Total / pageSize);
    totalNum = jsonResult.Total;
    privateTotal = jsonResult.PrivateTotal;
    setPage(page);
  }
}

function displayPrResult(page, jsonResult, onlyReturnNum, keyword) {
  var data = jsonResult.Result;
  var total = jsonResult.Total;
  $("#pr_total").text(total);
  if (!onlyReturnNum) {
    setActivate("pr_item");
    //$('#keyword_desc').text(keyword);
    //$('#obj_desc').text(getLabel(isZh,"search_pr"));
    //$('#child_total').text(total);
    $("#find_title").html(
      getLabel(isZh, "find_title")
        .replace("{keyword}", keyword)
        .replace("{tablename}", getLabel(isZh, "search_pr"))
        .replace("{total}", total)
    );

    setIssueOrPrInnerHtml(data, "pulls");
  }
}

var categoryDesc = {
  computer_vision: "计算机视觉",
  natural_language_processing: "自然语言处理",
  speech_processing: "语音处理",
  computer_vision_natural_language_processing: "计算机视觉、自然语言处理",
  medical_imaging: "医学影像",
};

var categoryENDesc = {
  computer_vision: "computer vision",
  natural_language_processing: "natural language processing",
  speech_processing: "speech processing",
  computer_vision_natural_language_processing:
    "computer vision and natural language processing",
  medical_imaging: "medical imaging",
};

var taskDesc = {
  machine_translation: "机器翻译",
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
  speech_recognition: "语音识别",
  speech_synthesis: "语音合成",
  ros_hmci_datasets: "开源开放社区",
};

var taskENDesc = {
  machine_translation: "machine translation",
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
  speech_recognition: "speech recognition",
  speech_synthesis: "speech synthesis",
  ros_hmci_datasets: "ROS-hmci datasets",
};

function getCategoryDescNew(isZh, key) {
  var result = "";
  try {
    let jsonRe = JSON.parse(key);
    for (var i = 0; i < jsonRe.length; i++) {
      result += getCategoryDesc(isZh, jsonRe[i]);
      if (i < (jsonRe.length - 1)) {
        result += ",";
      }
    }
    return result;
  } catch (error) {
    console.error('Parsing error:', error);
  }
  return getCategoryDesc(isZh, key);
}

function getCategoryDesc(isZh, key) {
  var re = key;
  if (isZh) {
    re = categoryDesc[key];
  } else {
    re = categoryENDesc[key];
  }
  if (isEmpty(re)) {
    return key;
  }
  return re;
}

function getTaskDescNew(isZh, key) {
  var result = "";
  try {
    let jsonRe = JSON.parse(key);
    for (var i = 0; i < jsonRe.length; i++) {
      result += getTaskDesc(isZh, jsonRe[i]);
      if (i < (jsonRe.length - 1)) {
        result += ",";
      }
    }
    return result;
  } catch (error) {
    console.error('Parsing error:', error);
  }
  return getTaskDesc(isZh, key);
}

function getTaskDesc(isZh, key) {
  var re = key;
  if (isZh) {
    re = taskDesc[key];
  } else {
    re = taskENDesc[key];
  }
  if (isEmpty(re)) {
    return key;
  }
  return re;
}

function getActiveItem(sort_type) {
  if (
    currentSearchSortBy == sortBy[sort_type] &&
    currentSearchAscending == sortAscending[sort_type]
  ) {
    return "active ";
  } else {
    return "";
  }
}

function displayDataSetResult(page, jsonResult, onlyReturnNum, keyword) {
  var data = jsonResult.Result;
  var total = jsonResult.Total;
  $("#dataset_total").text(total);
  if (!onlyReturnNum) {
    setActivate("dataset_item");
    //$('#keyword_desc').text(keyword);
    //$('#obj_desc').text(getLabel(isZh,"search_dataset"));
    //$('#child_total').text(total);
    $("#find_title").html(
      getLabel(isZh, "find_title")
        .replace("{keyword}", keyword)
        .replace("{tablename}", getLabel(isZh, "search_dataset"))
        .replace("{total}", total)
    );

    var sortHtml = "";
    sortHtml +=
      '<a class="' +
      getActiveItem(50) +
      'item" href="javascript:searchItem(5,50);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_matched") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(51) +
      'item" href="javascript:searchItem(5,51);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_matched_download") +
      "</a>";
    document.getElementById("sort_type").innerHTML = sortHtml;

    var html = "";
    var currentTime = new Date().getTime();
    for (var i = 0; i < data.length; i++) {
      var recordMap = data[i];
      html += '<div class="item">';
      html += '  <div class="content">';
      html += '       <div class="ui right metas">';
      if (!isEmpty(recordMap["category"])) {
        html +=
          '       <span class="text grey"><svg class="svg octicon-tasklist" width="16" height="16" aria-hidden="true"><use xlink:href="#octicon-tasklist" /></svg> ' +
          getCategoryDescNew(isZh, recordMap["category"]) +
          "</span>";
      }
      if (!isEmpty(recordMap["task"])) {
        html +=
          '        <span class="text grey"><svg class="svg octicon-tag" width="16" height="16" aria-hidden="true"><use xlink:href="#octicon-tag" /></svg>' +
          getTaskDescNew(isZh, recordMap["task"]) +
          "</span>";
      }
      html +=
        '           <span class="text grey"><i class="ri-fire-line"></i> ' +
        recordMap["download_times"] +
        "</span>     ";
      html += "        </div>";
      html += '       <div class="ui header">';
      html +=
        '           <a class="name" href="/datasets/detail/' +
        recordMap["owerName"] + '/' + recordMap["name"] +
        '" target="_blank">' +
        recordMap["title"] +
        "</a>";
      html +=
        '            <span class="middle"><svg class="svg octicon-repo-clone" width="16" height="16" aria-hidden="true"><use xlink:href="#octicon-repo-clone"></use></svg></span>';
      html += "        </div>";
      html += '       <div class="description">';
      html +=
        '           <p class="has-emoji"> ' + escapeHTML(recordMap["description"]) + "</p>";
      if (!isEmpty(recordMap["file_name"])) {
        html +=
          '           <p class="has-emoji"> ' + escapeHTML(recordMap["file_name"]) + "</p>";
      }
      html += '            <p class="time">';
      html +=
        ' <span class="am-ml-10"></span>  ' +
        getLabel(isZh, "search_lasted_update") +
        " " +
        recordMap["updated_html"];
      html += "            </p>";
      html += "        </div>";
      html += "   </div>";
      html += "</div>";
    }
    document.getElementById("child_search_item").innerHTML = html;
  }
}

function displayModelResult(page, jsonResult, onlyReturnNum, keyword) {
  var data = jsonResult.Result;
  var total = jsonResult.Total;
  $("#model_total").text(total);
  if (!onlyReturnNum) {
    setActivate("model_item");
    //$('#keyword_desc').text(keyword);
    //$('#obj_desc').text(getLabel(isZh,"search_model"));
    //$('#child_total').text(total);
    $("#find_title").html(
      getLabel(isZh, "find_title")
        .replace("{keyword}", keyword)
        .replace("{tablename}", getLabel(isZh, "search_model"))
        .replace("{total}", total)
    );

    var sortHtml = "";
    sortHtml +=
      '<a class="' +
      getActiveItem(70) +
      'item" href="javascript:searchItem(7,70);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_matched") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(71) +
      'item" href="javascript:searchItem(7,71);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_matched_reference") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(72) +
      'item" href="javascript:searchItem(7,72);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_matched_download") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(73) +
      'item" href="javascript:searchItem(7,73);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_lasted_create") +
      "</a>";
    document.getElementById("sort_type").innerHTML = sortHtml;

    var html = "";
    var currentTime = new Date().getTime();
    const engineMap = {
      '0': 'PyTorch',
      '1': 'TensorFlow',
      '2': 'MindSpore',
      '4': 'PaddlePaddle',
      '5': 'OneFlow',
      '6': 'MXNet',
      '3': 'Other',
    };
    for (var i = 0; i < data.length; i++) {
      var recordMap = data[i];
      var createDate = new Date(recordMap['created_unix'] * 1000);
      var createYear = createDate.getFullYear().toString();
      var createMonth = (createDate.getMonth() + 1).toString();
      var createDay = createDate.getDate().toString();
      html += `<div class="item">
        <div class="content">
          <div class="ui right metas" style="color:#767676">${engineMap[recordMap['engine']] || 'Other'}</div>
          <div class="ui header">
            <a class="name" href="/models/detail/${recordMap["owerName"]}/${recordMap["real_name"]}" target="_blank">
              ${recordMap["title"]}
            </a>
          </div>
          <div class="description">
            <p class="labels has-emoji">${recordMap['label'] ? recordMap['label'].replace(/font\scolor=/g, 'font_color=').trim().split(/\s+/).map(item => {
        return '<span style="color:rgba(16, 16, 16, 0.8);border-radius:4px;font-size:12px;background:rgba(232, 232, 232, 0.6);padding:2px 6px;margin-right:8px">'
          + item.replace(/font_color=/g, 'font color=') + '</span>';
      }).join('') : ''}</p>
            <p class="descr has-emoji">${escapeHTML(recordMap['description'])}</p>
            <p class="filename has-emoji">${recordMap['file_name'] || ''}</p>
            <p class="time" style="display:flex;align-items:center">
              <a style="margin-right:8px;display:inline-block;height:22px;" href="/${recordMap['owerName']}"><img src="/user/avatar/${recordMap['owerName']}/-1" 
                style="display:inline-block;width:22px;height:22px;border-radius:100%;"></a>
              <span style="margin-right:8px" title="${getLabel(isZh, "create_time")}">
                ${createYear}-${createMonth.length < 2 ? '0' + createMonth : createMonth}-${createDay.length < 2 ? '0' + createDay : createDay}
              </span>
              <span title="${getLabel(isZh, "search_matched_reference")}" style="display:flex;align-items:center">
                <i class="ri-link" style="margin-right:4px"></i><span style="margin-right:4px">${recordMap['reference_count']}</span>
              </span>
              <span title="${getLabel(isZh, "search_matched_download")}" style="display:flex;align-items:center">
                <i class="ri-download-line" style="margin-right:4px"></i><span style="margin-right:8px">${recordMap['download_count']}</span>
              </span>
              ${getLabel(isZh, "search_lasted_update")}&nbsp;${recordMap["updated_html"]}
            </p>
          </div>
        </div>
      </div>
      `;
    }
    document.getElementById("child_search_item").innerHTML = html;
  }
}

function displayOrgResult(page, jsonResult, onlyReturnNum, keyword) {
  var data = jsonResult.Result;
  var total = jsonResult.Total;
  $("#org_total").text(total);
  if (!onlyReturnNum) {
    setActivate("org_item");
    //$('#keyword_desc').text(keyword);
    //$('#obj_desc').text(getLabel(isZh,"search_org"));
    //$('#child_total').text(total);
    $("#find_title").html(
      getLabel(isZh, "find_title")
        .replace("{keyword}", keyword)
        .replace("{tablename}", getLabel(isZh, "search_org"))
        .replace("{total}", total)
    );

    var sortHtml = "";
    sortHtml +=
      '<a class="' +
      getActiveItem(40) +
      'item" href="javascript:searchItem(4,40);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_matched") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(41) +
      'item" href="javascript:searchItem(4,41);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_letter_asc") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(42) +
      'item" href="javascript:searchItem(4,42);" tabindex="-1" role="menuitem" id="menuitem_2">' +
      getLabel(isZh, "search_letter_desc") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(43) +
      'item" href="javascript:searchItem(4,43);" tabindex="-1" role="menuitem" id="menuitem_2">' +
      getLabel(isZh, "search_lasted_create") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(44) +
      'item" href="javascript:searchItem(4,44);" tabindex="-1" role="menuitem" id="menuitem_2">' +
      getLabel(isZh, "search_early_create") +
      "</a>";
    document.getElementById("sort_type").innerHTML = sortHtml;

    var html = "";
    var currentTime = new Date().getTime();
    for (var i = 0; i < data.length; i++) {
      var recordMap = data[i];

      html += '<div class="item members">';
      html +=
        '<img class="ui avatar image" src="' + recordMap["avatar"] + '"></img>';
      html += '  <div class="content">';
      html += '       <div class="ui header">';
      html +=
        '           <a class="name" href="/' +
        recordMap["real_name"] +
        '" target="_blank">' +
        recordMap["name"] +
        "&nbsp;&nbsp;" +
        escapeHTML(recordMap["full_name"]) +
        "</a>";
      html += "        </div>";
      html += '       <div class="description">';
      html +=
        '           <p class="has-emoji"> ' + escapeHTML(recordMap["description"]) + "</p>";
      html += '            <p class="has-emoji">';
      if (!isEmpty(recordMap["location"]) && recordMap["location"] != "null") {
        html +=
          '               <i class="ri-map-pin-2-line"></i> ' +
          escapeHTML(recordMap["location"]);
      }
      html += '               <span class="am-ml-10"></span>';
      if (!isEmpty(recordMap["website"]) && recordMap["website"] != "null") {
        html +=
          '           <i class="ri-links-line"></i>' +
          '<a href="' +
          escapeHTML(recordMap["website"]) +
          '" target="_blank">' +
          escapeHTML(recordMap["website"]) +
          "</a>";
      }
      html +=
        '               <i class="ri-time-line am-ml-10"></i>  ' +
        getLabel(isZh, "search_add_by") +
        " ";
      html += recordMap["add_time"];
      html += "            </p>";
      html += "        </div>";
      html += "   </div>";
      html += "</div>";
    }
    document.getElementById("child_search_item").innerHTML = html;
  }
}
var monthDisplay = new Array(
  "Jan",
  "Feb",
  "Mar",
  "Apr",
  "May",
  "Jun",
  "Jul",
  "Aug",
  "Spt",
  "Oct",
  "Nov",
  "Dec"
);
function displayUserResult(page, jsonResult, onlyReturnNum, keyword) {
  var data = jsonResult.Result;
  var total = jsonResult.Total;
  $("#user_total").text(total);
  if (!onlyReturnNum) {
    setActivate("user_item");
    //$('#keyword_desc').text(keyword);
    //$('#obj_desc').text(getLabel(isZh,"search_user"));
    //$('#child_total').text(total);

    $("#find_title").html(
      getLabel(isZh, "find_title")
        .replace("{keyword}", keyword)
        .replace("{tablename}", getLabel(isZh, "search_user"))
        .replace("{total}", total)
    );

    var sortHtml = ""; //equal user sort by
    sortHtml +=
      '<a class="' +
      getActiveItem(30) +
      'item" href="javascript:searchItem(3,30);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_matched") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(31) +
      'item" href="javascript:searchItem(3,31);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_letter_asc") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(32) +
      'item" href="javascript:searchItem(3,32);" tabindex="-1" role="menuitem" id="menuitem_2">' +
      getLabel(isZh, "search_letter_desc") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(33) +
      'item" href="javascript:searchItem(3,33);" tabindex="-1" role="menuitem" id="menuitem_2">' +
      getLabel(isZh, "search_lasted_create") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(34) +
      'item" href="javascript:searchItem(3,34);" tabindex="-1" role="menuitem" id="menuitem_2">' +
      getLabel(isZh, "search_early_create") +
      "</a>";

    document.getElementById("sort_type").innerHTML = sortHtml;

    var html = "";
    var currentTime = new Date().getTime();
    for (var i = 0; i < data.length; i++) {
      var recordMap = data[i];
      html += '<div class="item members">';
      html +=
        '<img class="ui avatar image" src="' + recordMap["avatar"] + '"></img>';
      html += '  <div class="content">';
      html += '       <div class="ui header">';
      html +=
        '           <a class="name" href="/' +
        recordMap["real_name"] +
        '" target="_blank">' +
        recordMap["name"] +
        "&nbsp;&nbsp;" +
        escapeHTML(recordMap["full_name"]) +
        "</a>";
      html += "        </div>";
      html += '       <div class="description">';
      html +=
        '           <p class="has-emoji"> ' + escapeHTML(recordMap["description"]) + "</p>";
      html += '            <p class="has-emoji">';
      if (!isEmpty(recordMap["email"]) && recordMap["email"] != "null") {
        html +=
          '               <i class="ri-mail-line"></i>&nbsp;<a href="mailto:' +
          recordMap["email"] +
          '" rel="nofollow">' +
          recordMap["email"] +
          "</a>";
      }
      html +=
        '               <i class="ri-time-line am-ml-10"></i>  ' +
        getLabel(isZh, "search_add_by") +
        " ";
      html += recordMap["add_time"];
      html += "            </p>";
      html += "        </div>";
      html += "   </div>";
      html += "</div>";
    }
    document.getElementById("child_search_item").innerHTML = html;
  }
}

function setIssueOrPrInnerHtml(data, path) {
  var sortHtml = "";
  if (path == "issues") {
    sortHtml +=
      '<a class="' +
      getActiveItem(20) +
      'item" href="javascript:searchItem(2,20);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_matched") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(21) +
      'item" href="javascript:searchItem(2,21);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_lasted") +
      "</a>";
  } else {
    sortHtml +=
      '<a class="' +
      getActiveItem(60) +
      'item" href="javascript:searchItem(6,60);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_matched") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(61) +
      'item" href="javascript:searchItem(6,61);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_lasted") +
      "</a>";
  }

  document.getElementById("sort_type").innerHTML = sortHtml;

  var html = "";
  var currentTime = new Date().getTime();
  for (var i = 0; i < data.length; i++) {
    var recordMap = data[i];
    html += '<div class="item">';
    html += '  <div class="content">';
    html += '       <div class="ui header">';
    html +=
      '           <a class="name" href="/' +
      recordMap["repoUrl"] +
      "/" +
      path +
      "/" +
      recordMap["index"] +
      '" target="_blank">' +
      recordMap["name"] +
      "</a>";
    html += "        </div>";
    html += '       <div class="description">';
    html += '           <p class="has-emoji"> ' + escapeHTML(recordMap["content"]) + "</p>";
    html += '            <p class="time">';
    html += '               <i class="ri-code-box-line"></i>';
    html +=
      '               <a class="am-text grey" href="/' +
      recordMap["repoUrl"] +
      "/" +
      path +
      "/" +
      recordMap["index"] +
      '" target="_blank"> ' +
      addBlank(recordMap["repoUrl"]) +
      " #" +
      recordMap["index"] +
      "</a>&nbsp;&nbsp;&nbsp;&nbsp;";
    html += '               <i class="ri-information-line am-ml-10"></i>&nbsp;';
    if (
      recordMap["is_closed"] != null &&
      (!recordMap["is_closed"] || recordMap["is_closed"] == "f")
    ) {
      html += getLabel(isZh, "search_open");
    } else {
      html += getLabel(isZh, "search_closed");
    }
    html +=
      ' &nbsp;&nbsp;&nbsp;&nbsp;<i class="ri-message-2-line am-ml-10"></i>&nbsp;' +
      recordMap["num_comments"];

    html +=
      ' <span class="am-ml-10">&nbsp;&nbsp;</span>&nbsp;&nbsp;' +
      getLabel(isZh, "search_lasted_update") +
      " " +
      recordMap["updated_html"];

    html += "            </p>";
    html += "        </div>";
    html += "   </div>";
    html += "</div>";
  }
  document.getElementById("child_search_item").innerHTML = html;
}

function addBlank(url) {
  if (url == null) {
    return url;
  }
  var tmps = url.split("/");
  if (tmps.length == 2) {
    return tmps[0] + " / " + tmps[1];
  }
  return url;
}

function displayIssueResult(page, jsonResult, onlyReturnNum, keyword) {
  var data = jsonResult.Result;
  var total = jsonResult.Total;
  $("#issue_total").text(total);
  if (!onlyReturnNum) {
    setActivate("issue_item");
    //$('#keyword_desc').text(keyword);
    //$('#obj_desc').text(getLabel(isZh,"search_issue"));
    //$('#child_total').text(total);
    $("#find_title").html(
      getLabel(isZh, "find_title")
        .replace("{keyword}", keyword)
        .replace("{tablename}", getLabel(isZh, "search_issue"))
        .replace("{total}", total)
    );

    setIssueOrPrInnerHtml(data, "issues");
  }
}

function setActivate(name) {
  $("#repo_item").removeClass("active");
  $("#user_item").removeClass("active");
  $("#issue_item").removeClass("active");
  $("#dataset_item").removeClass("active");
  $("#model_item").removeClass("active");
  $("#org_item").removeClass("active");
  $("#pr_item").removeClass("active");
  if (name == null) {
    return;
  }
  var tmp = "#" + name;
  $(tmp).addClass("active");
}

function LetterAvatar(name, size, color) {
  name = name || "";
  size = size || 60;
  var colours = [
    "#1abc9c",
    "#2ecc71",
    "#3498db",
    "#9b59b6",
    "#34495e",
    "#16a085",
    "#27ae60",
    "#2980b9",
    "#8e44ad",
    "#2c3e50",
    "#f1c40f",
    "#e67e22",
    "#e74c3c",
    "#00bcd4",
    "#95a5a6",
    "#f39c12",
    "#d35400",
    "#c0392b",
    "#bdc3c7",
    "#7f8c8d",
  ],
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
  let initials1 = initials.toUpperCase();
  initials.toUpperCase();
  if (w.devicePixelRatio) {
    size = size * w.devicePixelRatio;
  }

  charIndex = (initials == "?" ? 72 : initials.charCodeAt(0)) - 64;
  colourIndex = charIndex % 20;
  canvas = d.createElement("canvas");
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
  Array.prototype.forEach.call(
    d.querySelectorAll("img[avatar]"),
    function (img, name, color) {
      name = img.getAttribute("avatar");
      color = img.getAttribute("color");
      img.src = LetterAvatar(name, img.getAttribute("width"), color);
      img.removeAttribute("avatar");
      img.setAttribute("alt", name);
    }
  );
};

function displayRepoResult(page, jsonResult, onlyReturnNum, keyword) {
  var data = jsonResult.Result;
  var total = jsonResult.Total;
  $("#repo_total").text(total);

  if (!onlyReturnNum) {
    setActivate("repo_item");
    // $('#keyword_desc').text(keyword);
    //$('#obj_desc').text(getLabel(isZh,"search_repo"));
    //$('#child_total').text(total);
    $("#find_title").html(
      getLabel(isZh, "find_title")
        .replace("{keyword}", keyword)
        .replace("{tablename}", getLabel(isZh, "search_repo"))
        .replace("{total}", total)
    );

    var sortHtml = "";
    sortHtml +=
      '<a class="' +
      getActiveItem(10) +
      'item" href="javascript:searchItem(1,10);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_matched") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(11) +
      'item" href="javascript:searchItem(1,11);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_lasted") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(12) +
      'item" href="javascript:searchItem(1,12);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_watched") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(13) +
      'item" href="javascript:searchItem(1,13);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_star") +
      "</a>";
    sortHtml +=
      '<a class="' +
      getActiveItem(14) +
      'item" href="javascript:searchItem(1,14);" tabindex="-1" role="menuitem" id="menuitem_1">' +
      getLabel(isZh, "search_fork") +
      "</a>";

    document.getElementById("sort_type").innerHTML = sortHtml;

    var html = "";
    var currentTime = new Date().getTime();
    for (var i = 0; i < data.length; i++) {
      var recordMap = data[i];
      html += '<div class="item">';
      if (recordMap["avatar"]) {
        html += `<img class="ui avatar image" src="${recordMap["avatar"]}">`;
      } else {
        html += `<img class="ui avatar image" avatar="${recordMap["owner_name"]}">`;
      }

      html += '  <div class="content">';
      html += '       <div class="ui header">';
      html +=
        '           <a class="name" href="/' +
        recordMap["owner_name"] +
        "/" +
        recordMap["real_name"] +
        '" target="_blank"> <span>' +
        recordMap["owner_name"] +
        "</span> <span>/</span> <strong>" +
        recordMap["alias"] +
        "</strong></a>";
      if (recordMap["is_private"]) {
        html +=
          '        <span class="middle text gold"><svg class="svg octicon-lock" width="16" height="16" aria-hidden="true"><use xlink:href="#octicon-lock" /></svg></span>';
      }
      html += "        </div>";
      html += '       <div class="description">';
      html +=
        '           <p class="has-emoji"> ' + escapeHTML(recordMap["description"]) + "</p>";
      html += '           <div class="ui tags">';
      if (!isEmpty(recordMap["topics"]) && recordMap["topics"] != "null") {
        for (var j = 0; j < recordMap["topics"].length; j++) {
          //function searchLabel(tableName,keyword,sortBy="",ascending=false)
          html +=
            "       <a href=\"javascript:searchLabel('repository','" +
            recordMap["topics"][j] +
            "','updated_unix.keyword',false);\" ><div class=\"ui small label topic\">" +
            recordMap["hightTopics"][j] +
            "</div></a>";
        }
      }
      html += "            </div>";
      html += '            <p class="time">';
      html +=
        '                <i class="icon fa-eye outline"></i>&nbsp;' +
        recordMap["num_watches"] +
        '&nbsp;&nbsp;<i class="icon star outline"></i>&nbsp;' +
        recordMap["num_stars"] +
        '&nbsp;&nbsp;<i class="icon code branch"></i>&nbsp;' +
        recordMap["num_forks"] +
        "&nbsp;&nbsp;";
      html +=
        "&nbsp;&nbsp;&nbsp;&nbsp;" +
        getLabel(isZh, "search_lasted_update") +
        " " +
        recordMap["updated_html"];
      if (!isEmpty(recordMap["lang"])) {
        var lang = recordMap["lang"];
        var tmpLang = recordMap["lang"].split(",");
        if (tmpLang.length > 0) {
          lang = tmpLang[0];
        }
        var backColor = "#3572A5";
        if (LanguagesColor[lang] != null) {
          backColor = LanguagesColor[lang];
        }
        html +=
          '            <span class="text grey am-ml-10"><i class="color-icon" style="background-color: ' +
          backColor +
          '"></i>&nbsp;' +
          lang +
          "</span>";
      }
      html += "            </p>";
      html += "        </div>";
      html += "   </div>";
      html += "</div>";
    }

    document.getElementById("child_search_item").innerHTML = html;
    LetterAvatar.transform();
  }
}

function getTime(UpdatedUnix, currentTime) {
  UpdatedUnix = UpdatedUnix;
  currentTime = currentTime / 1000;
  var timeEscSecond = currentTime - UpdatedUnix;
  if (timeEscSecond < 0) {
    timeEscSecond = 1;
  }

  var hours = Math.floor(timeEscSecond / 3600);
  //计算相差分钟数
  var leave2 = Math.floor(timeEscSecond % 3600); //计算小时数后剩余的秒数
  var minutes = Math.floor(leave2 / 60); //计算相差分钟数

  var leave3 = Math.floor(leave2 % 60); //计算分钟数后剩余的秒数
  var seconds = leave3;

  if (hours == 0 && minutes == 0) {
    return seconds + getRepoOrOrg(6, isZh);
  } else {
    if (hours > 0) {
      if (hours >= 24) {
        var days = Math.ceil(hours / 24);
        if (days >= 30 && days < 365) {
          return Math.ceil(days / 30) + getRepoOrOrg(8, isZh);
        } else if (days >= 365) {
          return Math.ceil(days / 365) + getRepoOrOrg(9, isZh);
        }
        return Math.ceil(hours / 24) + getRepoOrOrg(7, isZh);
      } else {
        return hours + getRepoOrOrg(4, isZh);
      }
    } else {
      return minutes + getRepoOrOrg(5, isZh);
    }
  }
}

function getRepoOrOrg(key, isZhLang) {
  if (isZhLang) {
    return repoAndOrgZH[key];
  } else {
    return repoAndOrgEN[key];
  }
}

var repoAndOrgZH = {
  1: "项目",
  2: "成员",
  3: "团队",
  4: "小时前",
  5: "分钟前",
  6: "秒前",
  7: "天前",
  8: "个月前",
  9: "年前",
};

var repoAndOrgEN = {
  1: "repository",
  2: "Members ",
  3: "Teams",
  4: " hours ago",
  5: " minutes ago",
  6: " seconds ago",
  7: " day ago",
  8: " month ago",
  9: " year ago",
};

function page(current) {
  currentPage = current;

  doSearch(
    currentSearchTableName,
    currentSearchKeyword,
    current,
    pageSize,
    false,
    currentSearchSortBy,
    OnlySearchLabel
  );
}

function nextPage() {
  currentPage = currentPage + 1;
  page(currentPage);
}

function prePage() {
  if (currentPage > 1) {
    currentPage = currentPage - 1;
    page(currentPage);
  }
}

function getXPosition(e) {
  var x = e.offsetLeft;
  while ((e = e.offsetParent)) {
    x += e.offsetLeft;
  }
  return x + 20; //-260防止屏幕超出
}
//获取y坐标
function getYPosition(e) {
  var y = e.offsetTop;
  while ((e = e.offsetParent)) {
    y += e.offsetTop;
  }
  return y + 20; //80为input高度
}

function goPage(event) {
  var inputpage = document.getElementById("inputpage_div");
  var left = getXPosition(event.target);
  var top = getYPosition(event.target);
  var goNum = $("#inputpage").val();
  if (goNum <= 0) {
    showTip(getLabel(isZh, "search_input_large_0"), "warning", left + 5, top);
  } else if (goNum <= totalPage) {
    page(parseInt(goNum, 10));
  } else {
    showTip(getLabel(isZh, "search_input_maxed"), "warning", left + 5, top);
  }
}

function showTip(tip, type, left, top) {
  var $tip = $("#tipmsg");
  var tipmsg = document.getElementById("tipmsg");
  var style =
    "z-index:10024;top:" +
    top +
    "px;left:" +
    left +
    "px;position:absolute;width:200px;height:60px;vertical-align:middle;";
  tipmsg.style = style;
  var html = "<p>" + tip + "</p>";
  $tip
    .stop(true)
    .prop("class", "alert alert-" + type)
    .html(html)
    .fadeIn(500)
    .delay(2000)
    .fadeOut(500);
}

function setPage(currentPage) {
  var html = "";
  startIndex = currentPage - 1;
  if (startIndex < 1) {
    startIndex = 1;
  }
  endIndex = currentPage + 2;
  if (endIndex >= totalPage) {
    endIndex = totalPage;
  }
  html +=
    '<span class="item">' +
    getLabel(isZh, "search_input_total") +
    " " +
    totalNum +
    " " +
    getLabel(isZh, "search_srtip") +
    "</span>";
  if (currentPage > 1) {
    html +=
      '<a class="item navigation" href="javascript:page(1)"><span class="navigation_label">' +
      getLabel(isZh, "search_home_page") +
      "</span></a>";
    html +=
      '<a class="item navigation" href="javascript:prePage()"><i class="left arrow icon"></i></a>';
  } else {
    html +=
      '<a class="disabled item navigation" href="javascript:page(1)"><span class="navigation_label">' +
      getLabel(isZh, "search_home_page") +
      "</span></a>";
    html +=
      '<a class="disabled item navigation" href="javascript:prePage()"><i class="left arrow icon"></i></a>';
  }

  for (var i = startIndex; i <= endIndex; i++) {
    var page_i = i;
    if (page_i > totalPage) {
      break;
    }
    if (page_i == currentPage) {
      html +=
        '<a id="page_' +
        page_i +
        '" class="active item" href="javascript:page(' +
        page_i +
        ')">' +
        page_i +
        "</a>";
    } else {
      html +=
        '<a id="page_' +
        page_i +
        '" class="item" href="javascript:page(' +
        page_i +
        ')">' +
        page_i +
        "</a>";
    }
  }

  if (endIndex < totalPage - 1) {
    html += "...";
    html +=
      '<a id="page_' +
      totalPage +
      '" class="item" href="javascript:page(' +
      totalPage +
      ')">' +
      totalPage +
      "</a>";
  }

  if (currentPage >= totalPage) {
    html +=
      '<a class="disabled item navigation" href="javascript:nextPage()"><i class="icon right arrow"></i></a>';
    html +=
      '<a class="disabled item navigation" href="javascript:page(' +
      totalPage +
      ')"><span class="navigation_label">' +
      getLabel(isZh, "search_last_page") +
      "</span></a>";
  } else {
    html +=
      '<a class="item navigation" href="javascript:nextPage()"><i class="icon right arrow"></i></a>';
    html +=
      '<a class="item navigation" href="javascript:page(' +
      totalPage +
      ')"><span class="navigation_label">' +
      getLabel(isZh, "search_last_page") +
      "</span></a>";
  }

  html +=
    '<div class="item">  ' +
    getLabel(isZh, "search_go_to") +
    '<div id="inputpage_div" class="ui input"><input id="inputpage" type="text"></div>' +
    getLabel(isZh, "search_go_page") +
    "</div>";
  document.getElementById("page_menu").innerHTML = html;
  $("#inputpage").on("keypress", function (event) {
    if (event.keyCode == 13) {
      goPage(event);
    }
  });
}

$("#keyword_input").on("keypress", function (event) {
  if (event.keyCode == 13) {
    search();
  }
});

var LanguagesColor = {
  "1C Enterprise": "#814CCC",
  ABAP: "#E8274B",
  "AGS Script": "#B9D9FF",
  AMPL: "#E6EFBB",
  ANTLR: "#9DC3FF",
  "API Blueprint": "#2ACCA8",
  APL: "#5A8164",
  ASP: "#6a40fd",
  ATS: "#1ac620",
  ActionScript: "#882B0F",
  Ada: "#02f88c",
  Agda: "#315665",
  Alloy: "#64C800",
  AngelScript: "#C7D7DC",
  AppleScript: "#101F1F",
  Arc: "#aa2afe",
  AspectJ: "#a957b0",
  Assembly: "#6E4C13",
  Asymptote: "#4a0c0c",
  AutoHotkey: "#6594b9",
  AutoIt: "#1C3552",
  Ballerina: "#FF5000",
  Batchfile: "#C1F12E",
  BlitzMax: "#cd6400",
  Boo: "#d4bec1",
  Brainfuck: "#2F2530",
  C: "#555555",
  "C#": "#178600",
  "C++": "#f34b7d",
  CSS: "#563d7c",
  Ceylon: "#dfa535",
  Chapel: "#8dc63f",
  Cirru: "#ccccff",
  Clarion: "#db901e",
  Clean: "#3F85AF",
  Click: "#E4E6F3",
  Clojure: "#db5855",
  CoffeeScript: "#244776",
  ColdFusion: "#ed2cd6",
  "Common Lisp": "#3fb68b",
  "Common Workflow Language": "#B5314C",
  "Component Pascal": "#B0CE4E",
  Crystal: "#000100",
  Cuda: "#3A4E3A",
  D: "#ba595e",
  DM: "#447265",
  Dart: "#00B4AB",
  DataWeave: "#003a52",
  Dhall: "#dfafff",
  Dockerfile: "#384d54",
  Dogescript: "#cca760",
  Dylan: "#6c616e",
  E: "#ccce35",
  ECL: "#8a1267",
  EQ: "#a78649",
  Eiffel: "#946d57",
  Elixir: "#6e4a7e",
  Elm: "#60B5CC",
  "Emacs Lisp": "#c065db",
  EmberScript: "#FFF4F3",
  Erlang: "#B83998",
  "F#": "#b845fc",
  "F*": "#572e30",
  FLUX: "#88ccff",
  Factor: "#636746",
  Fancy: "#7b9db4",
  Fantom: "#14253c",
  Faust: "#c37240",
  Forth: "#341708",
  Fortran: "#4d41b1",
  FreeMarker: "#0050b2",
  Frege: "#00cafe",
  "G-code": "#D08CF2",
  GAML: "#FFC766",
  GDScript: "#355570",
  "Game Maker Language": "#71b417",
  Genie: "#fb855d",
  Gherkin: "#5B2063",
  Glyph: "#c1ac7f",
  Gnuplot: "#f0a9f0",
  Go: "#00ADD8",
  Golo: "#88562A",
  Gosu: "#82937f",
  "Grammatical Framework": "#79aa7a",
  Groovy: "#e69f56",
  HTML: "#e34c26",
  Hack: "#878787",
  Harbour: "#0e60e3",
  Haskell: "#5e5086",
  Haxe: "#df7900",
  HiveQL: "#dce200",
  HolyC: "#ffefaf",
  Hy: "#7790B2",
  IDL: "#a3522f",
  "IGOR Pro": "#0000cc",
  Idris: "#b30000",
  Io: "#a9188d",
  Ioke: "#078193",
  Isabelle: "#FEFE00",
  J: "#9EEDFF",
  JSONiq: "#40d47e",
  Java: "#b07219",
  JavaScript: "#f1e05a",
  Jolie: "#843179",
  Jsonnet: "#0064bd",
  Julia: "#a270ba",
  "Jupyter Notebook": "#DA5B0B",
  KRL: "#28430A",
  Kotlin: "#F18E33",
  LFE: "#4C3023",
  LLVM: "#185619",
  LOLCODE: "#cc9900",
  LSL: "#3d9970",
  Lasso: "#999999",
  Lex: "#DBCA00",
  LiveScript: "#499886",
  LookML: "#652B81",
  Lua: "#000080",
  MATLAB: "#e16737",
  MAXScript: "#00a6a6",
  MLIR: "#5EC8DB",
  MQL4: "#62A8D6",
  MQL5: "#4A76B8",
  MTML: "#b7e1f4",
  Makefile: "#427819",
  Mask: "#f97732",
  Max: "#c4a79c",
  Mercury: "#ff2b2b",
  Meson: "#007800",
  Metal: "#8f14e9",
  Mirah: "#c7a938",
  "Modula-3": "#223388",
  NCL: "#28431f",
  Nearley: "#990000",
  Nemerle: "#3d3c6e",
  NetLinx: "#0aa0ff",
  "NetLinx+ERB": "#747faa",
  NetLogo: "#ff6375",
  NewLisp: "#87AED7",
  Nextflow: "#3ac486",
  Nim: "#37775b",
  Nit: "#009917",
  Nix: "#7e7eff",
  Nu: "#c9df40",
  OCaml: "#3be133",
  ObjectScript: "#424893",
  "Objective-C": "#438eff",
  "Objective-C++": "#6866fb",
  "Objective-J": "#ff0c5a",
  Odin: "#60AFFE",
  Omgrofl: "#cabbff",
  Opal: "#f7ede0",
  OpenQASM: "#AA70FF",
  Oxygene: "#cdd0e3",
  Oz: "#fab738",
  P4: "#7055b5",
  PHP: "#4F5D95",
  PLSQL: "#dad8d8",
  Pan: "#cc0000",
  Papyrus: "#6600cc",
  Parrot: "#f3ca0a",
  Pascal: "#E3F171",
  Pawn: "#dbb284",
  Pep8: "#C76F5B",
  Perl: "#0298c3",
  PigLatin: "#fcd7de",
  Pike: "#005390",
  PogoScript: "#d80074",
  PostScript: "#da291c",
  PowerBuilder: "#8f0f8d",
  PowerShell: "#012456",
  Processing: "#0096D8",
  Prolog: "#74283c",
  "Propeller Spin": "#7fa2a7",
  Puppet: "#302B6D",
  PureBasic: "#5a6986",
  PureScript: "#1D222D",
  Python: "#3572A5",
  QML: "#44a51c",
  Quake: "#882233",
  R: "#198CE7",
  RAML: "#77d9fb",
  RUNOFF: "#665a4e",
  Racket: "#3c5caa",
  Ragel: "#9d5200",
  Raku: "#0000fb",
  Rascal: "#fffaa0",
  Reason: "#ff5847",
  Rebol: "#358a5b",
  Red: "#f50000",
  "Ren'Py": "#ff7f7f",
  Ring: "#2D54CB",
  Riot: "#A71E49",
  Roff: "#ecdebe",
  Rouge: "#cc0088",
  Ruby: "#701516",
  Rust: "#dea584",
  SAS: "#B34936",
  SQF: "#3F3F3F",
  "SRecode Template": "#348a34",
  SaltStack: "#646464",
  Scala: "#c22d40",
  Scheme: "#1e4aec",
  Self: "#0579aa",
  Shell: "#89e051",
  Shen: "#120F14",
  Slash: "#007eff",
  Slice: "#003fa2",
  SmPL: "#c94949",
  Smalltalk: "#596706",
  Solidity: "#AA6746",
  SourcePawn: "#5c7611",
  Squirrel: "#800000",
  Stan: "#b2011d",
  "Standard ML": "#dc566d",
  Starlark: "#76d275",
  SuperCollider: "#46390b",
  Swift: "#ffac45",
  SystemVerilog: "#DAE1C2",
  "TI Program": "#A0AA87",
  Tcl: "#e4cc98",
  TeX: "#3D6117",
  Terra: "#00004c",
  Turing: "#cf142b",
  TypeScript: "#2b7489",
  UnrealScript: "#a54c4d",
  V: "#5d87bd",
  VBA: "#867db1",
  VBScript: "#15dcdc",
  VCL: "#148AA8",
  VHDL: "#adb2cb",
  Vala: "#fbe5cd",
  Verilog: "#b2b7f8",
  "Vim script": "#199f4b",
  "Visual Basic .NET": "#945db7",
  Volt: "#1F1F1F",
  Vue: "#2c3e50",
  WebAssembly: "#04133b",
  Wollok: "#a23738",
  X10: "#4B6BEF",
  XC: "#99DA07",
  XQuery: "#5232e7",
  XSLT: "#EB8CEB",
  YARA: "#220000",
  YASnippet: "#32AB90",
  Yacc: "#4B6C4B",
  ZAP: "#0d665e",
  ZIL: "#dc75e5",
  ZenScript: "#00BCD1",
  Zephir: "#118f9e",
  Zig: "#ec915c",
  eC: "#913960",
  "mIRC Script": "#926059",
  mcfunction: "#E22837",
  nesC: "#94B0C7",
  ooc: "#b0b77e",
  q: "#0040cd",
  sed: "#64b970",
  wdl: "#42f1f4",
  wisp: "#7582D1",
  xBase: "#403a40",
};

function getLabel(isZh, key) {
  if (isZh) {
    return zhCN[key];
  } else {
    return esUN[key];
  }
}

var zhCN = {
  search: "搜索",
  search_repo: "项目",
  search_dataset: "数据集",
  search_model: "模型",
  search_issue: "任务",
  search_pr: "合并请求",
  search_user: "用户",
  search_org: "组织",
  search_finded: "找到",
  search_matched: "最佳匹配",
  search_matched_download: "下载次数",
  search_matched_reference: "引用次数",
  search_lasted_update: "最后更新于",
  search_letter_asc: "字母顺序排序",
  search_letter_desc: "字母逆序排序",
  search_lasted_create: "最新创建",
  search_early_create: "最早创建",
  search_add_by: "加入于",
  search_lasted: "最近更新",
  search_open: "开启中",
  search_closed: "已关闭",
  search_watched: "关注数",
  search_star: "点赞数",
  search_fork: "Fork数",
  search_input_large_0: "请输入大于0的数值。",
  search_input_maxed: "不能超出总页数。",
  search_input_total: "共",
  search_srtip: "条",
  search_home_page: "首页",
  search_last_page: "末页",
  search_go_to: "前往",
  search_go_page: "页",
  find_title:
    '“<strong class="highlight" id="keyword_desc">{keyword}</strong>”相关{tablename}约为{total}个',
  search_empty: "<strong>请输入任意关键字开始搜索。</strong>",
  create_time: "创建时间"
};

var esUN = {
  search: "Search",
  search_repo: "Repository",
  search_dataset: "DataSet",
  search_model: "Model",
  search_issue: "Issue",
  search_pr: "Pull Request",
  search_user: "User",
  search_org: "Organization",
  search_finded: "Find",
  search_matched: "Best Match",
  search_matched_download: "Most downloads",
  search_matched_reference: "Most reference",
  search_lasted_update: "Updated ",
  search_letter_asc: "Alphabetically",
  search_letter_desc: "Reverse alphabetically",
  search_lasted_create: "Recently created",
  search_early_create: "First created",
  search_add_by: "Joined on",
  search_lasted: "Recently updated",
  search_open: "Open",
  search_closed: "Closed",
  search_watched: "Watches",
  search_star: "Stars",
  search_fork: "Forks",
  search_input_large_0: "Please enter a value greater than 0.",
  search_input_maxed: "Cannot exceed total pages.",
  search_input_total: "Total",
  search_srtip: "",
  search_home_page: "First",
  search_last_page: "Last",
  search_go_to: "Go",
  search_go_page: "Page",
  find_title:
    ' {total} "<strong class="highlight" id="keyword_desc">{keyword}</strong>" related {tablename}',
  search_empty:
    "<strong>Please enter any keyword to start the search.</strong>",
  create_time: "Create Time"
};
initDiv(false);
document.onreadystatechange = function () {
  if (document.readyState === "complete") {
    var tmpSearchLabel = sessionStorage.getItem("searchLabel");
    if (tmpSearchLabel) {
      sessionStorage.removeItem("searchLabel");
      doSearchLabel(
        sessionStorage.getItem("tableName"),
        sessionStorage.getItem("keyword"),
        sessionStorage.getItem("sortBy"),
        sessionStorage.getItem("ascending")
      );
    } else {
      var specifySearch = sessionStorage.getItem("specifySearch");
      if (specifySearch) {
        sessionStorage.removeItem("specifySearch");
        document.getElementById("keyword_input").value =
          sessionStorage.getItem("keyword");
        doSpcifySearch(
          sessionStorage.getItem("tableName"),
          sessionStorage.getItem("keyword"),
          sessionStorage.getItem("sortBy"),
          sessionStorage.getItem("ascending")
        );
      } else {
        search();
      }
    }
  }
};
