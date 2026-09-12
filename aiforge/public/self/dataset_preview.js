(function () {
  function DatasetPreview() {
    this.ajaxTimeOut = 1000 * 60;
    this.datasetData = DATASET_DATA;
    this.supportImgReg = /(\.jpg|\.jpeg|\.png|\.gif|\.bmp)$/i;
    this.supportTxtReg = /(\.txt|\.xml|\.html|\.json|\.py|\.sh|\.md|\.csv|\.log|\.js|\.css|\.ipynb)$/i;
    this.data = {
      filePath: [{ path: '', name: this.datasetData.path[0], ParenDir: '', marker: '' }],
      dirMarkerChildrenMap: {},
      fileContentMap: {},
      currentPage: '',
    };
    this.currentFileList = [];
    this.currentFile = '';
    this.pageSize = 100;
    this.getPathChildren();
    this.eventInit();
  }

  DatasetPreview.prototype.renderFilePath = function () {
    var container = $('#file_path_container').empty();
    var filePath = this.data.filePath;
    for (let i = 0, iLen = filePath.length; i < iLen; i++) {
      var path = filePath[i];
      var pathEle;
      if (i == iLen - 1) {
        pathEle = $(`<div>
        <span class="" href="javascript:;">${path.name}</span><div class="divider"> / </div>
        </div>`);
      } else {
        pathEle = $(`<div>
        <a class="section" href="javascript:;">${path.name}</a><div class="divider"> / </div>
        </div>`);
      }
      pathEle.data('data', path);
      container.append(pathEle);
    }
  };

  DatasetPreview.prototype.prevDirs = function () {
    this.data.filePath.pop();
    this.getPathChildren();
  };

  DatasetPreview.prototype.nextDirs = function (pathObj) {
    this.data.filePath.push({
      ...pathObj,
      path: pathObj.FileName,
      name: pathObj.FileName,
    });
    this.getPathChildren();
  };

  DatasetPreview.prototype.showGetPathChildrenTimeout = function () {
    $('#myCanvas_div .tabpannel.ui.form .query-timeout').show();
    $('#filelist').hide();
  }

  DatasetPreview.prototype.getPathChildrenMore = function () {
    var self = this;
    var lastDir = this.data.filePath[this.data.filePath.length - 1];
    var path = this.data.filePath.map((item) => item.path).join('/');
    if (!lastDir) return;
    $('#myCanvas_div .tabpannel.ui.form').addClass('loading');
    $.ajax({
      type: "get",
      url: `/api/v1/attachments/get_dir`,
      dataType: "json",
      timeout: this.ajaxTimeOut,
      data: {
        _csrf: this.datasetData.csrf,
        uuid: this.datasetData.uuid,
        marker: lastDir.marker,
        pageSize: this.pageSize,
        prefix: lastDir.ParenDir ? '/' + lastDir.ParenDir.replace(/^\//, '') : '',
      },
      success: function (res) {
        $('#myCanvas_div .tabpannel.ui.form').removeClass('loading');
        if (res.result_code == 0) {
          var result = res.data.sort((a, b) => {
            var a1 = a.IsDir ? 1 : 0;
            var b1 = b.IsDir ? 1 : 0;
            return b1 - a1;
          });
          var fileList = [];
          for (let i = 0, iLen = result.length; i < iLen; i++) {
            const file = result[i];
            if (self.currentFileList.findIndex(itm => `${itm.ParenDir}${itm.FileName}` == `${file.ParenDir}${file.FileName}`) >= 0) {
              continue;
            }
            fileList.push(file);
          }
          self.data.dirMarkerChildrenMap[path + '-' + lastDir.marker] = result;
          self.currentFileList = self.currentFileList.concat(fileList);
          lastDir.marker = res.marker;
          self.renderFileListAdd(fileList, result.length);
        } else {
          console.log(res);
        }
      },
      error: function (err) {
        $('#myCanvas_div .tabpannel.ui.form').removeClass('loading');
        console.log(err);
        if (err.statusText == 'timeout') {
          self.showGetPathChildrenTimeout();
        }
      }
    });
  }

  DatasetPreview.prototype.getPathChildren = function () {
    this.renderFilePath();
    var lastDir = this.data.filePath[this.data.filePath.length - 1];
    if (!lastDir) return;
    var path = this.data.filePath.map((item) => item.path).join('/');
    var cache = this.data.dirMarkerChildrenMap[path + '-' + lastDir.marker];
    var self = this;
    if (false && cache) {
      self.currentFileList = cache;
      this.renderFileList(cache);
    } else {
      $('#myCanvas_div .tabpannel.ui.form').addClass('loading');
      $.ajax({
        type: "get",
        url: `/api/v1/attachments/get_dir`,
        dataType: "json",
        timeout: this.ajaxTimeOut,
        data: {
          _csrf: this.datasetData.csrf,
          uuid: this.datasetData.uuid,
          marker: '',
          pageSize: this.pageSize,
          prefix: lastDir.ParenDir ? '/' + lastDir.ParenDir.replace(/^\//, '') : '',
        },
        success: function (res) {
          $('#myCanvas_div .tabpannel.ui.form').removeClass('loading');
          if (res.result_code == 0) {
            var result = res.data.sort((a, b) => {
              var a1 = a.IsDir ? 1 : 0;
              var b1 = b.IsDir ? 1 : 0;
              return b1 - a1;
            });
            self.data.dirMarkerChildrenMap[path + '-' + lastDir.marker] = result;
            lastDir.marker = res.marker;
            self.currentFileList = result;
            self.currentFile = '';
            self.renderFileList(result);
          } else {
            console.log(res);
          }
        },
        error: function (err) {
          $('#myCanvas_div .tabpannel.ui.form').removeClass('loading');
          console.log(err);
          if (err.statusText == 'timeout') {
            self.showGetPathChildrenTimeout();
          }
        }
      });
    }
  };

  DatasetPreview.prototype.renderFileListAdd = function (files, length) {
    var domC = $('#filelist');
    domC.find('.file-more').remove();
    for (let i = 0, iLen = files.length; i < iLen; i++) {
      const file = files[i];
      const fileEl = $(`<div class="file-item" title="${file.FileName}">
        <i width="16" height="16" aria-hidden="true" style="color: ${file.IsDir ? 'rgb(91, 185, 115);' : ''}"  class="icon ${file.IsDir ? 'folder outline' : 'file alternate outline'}"></i>
        <span>${file.FileName}</span>
      </div>`);
      fileEl.data('data', file);
      domC.append(fileEl);
    }
    var lastDir = this.data.filePath[this.data.filePath.length - 1];
    if (!lastDir) return;
    if (length >= this.pageSize && lastDir.marker) {
      domC.append(`<div class="file-more">${$('#lang-seemore').text()}</div>`);
    }
  };

  DatasetPreview.prototype.renderFileList = function (files) {
    var domC = $('#filelist').empty();
    var findFile = false;
    for (let i = 0, iLen = files.length; i < iLen; i++) {
      const file = files[i];
      const fileEl = $(`<div class="file-item" title="${file.FileName}">
        <i width="16" height="16" aria-hidden="true" style="color: ${file.IsDir ? 'rgb(91, 185, 115);' : ''}"  class="icon ${file.IsDir ? 'folder outline' : 'file alternate outline'}"></i>
        <span>${file.FileName}</span>
      </div>`);
      fileEl.data('data', file);
      domC.append(fileEl);
    }
    if (!findFile) {
      this.currentFile = null;
      this.renderPleseSelectFile();
    }
    var lastDir = this.data.filePath[this.data.filePath.length - 1];
    if (!lastDir) return;
    if (files.length >= this.pageSize && lastDir.marker) {
      domC.append(`<div class="file-more">${$('#lang-seemore').text()}</div>`);
    }
  };

  DatasetPreview.prototype.renderPreview = function (file) {
    var fileName = file.FileName;
    if (this.supportImgReg.test(fileName)) {
      console.log('renderImage');
      $('#textcontent').hide();
      $('#win_canvas .select-file').hide();
      $('#win_canvas .not-support').hide();
      $('#imgcontent').attr('src', `/api/v1/attachments/get_image_content?uuid=${this.datasetData.uuid}&filePath=${file.ParenDir + file.FileName}&type=${this.datasetData.type}&_csrf=${this.datasetData.csrf}`).show();
    } else if (this.supportTxtReg.test(fileName)) {
      console.log('getFileText');
      this.getFileText(file);
    } else {
      console.log('not support');
      this.renderNotSupport();
    }
  };

  DatasetPreview.prototype.renderNotSupport = function () {
    $('#imgcontent').hide();
    $('#textcontent').hide();
    $('#win_canvas .select-file').hide();
    $('#win_canvas .not-support').show();
  }

  DatasetPreview.prototype.renderPleseSelectFile = function () {
    $('#imgcontent').hide();
    $('#textcontent').hide();
    $('#win_canvas .not-support').hide();
    $('#win_canvas .select-file').show();
  }

  DatasetPreview.prototype.renderText = function (content) {
    $('#imgcontent').hide();
    $('#win_canvas .select-file').hide();
    $('#win_canvas .not-support').hide();
    $('#textcontent').empty().text(content).show();
  };

  DatasetPreview.prototype.getFileText = function (file) {
    var cache = this.data.fileContentMap[file.ParenDir + file.FileName];
    var self = this;
    if (false && cache) {
      self.renderText(cache);
    } else {
      $.ajax({
        type: "get",
        url: "/api/v1/attachments/get_txt_content",
        headers: { authorization: this.datasetData.csrf, },
        dataType: "json",
        data: {
          _csrf: this.datasetData.csrf,
          uuid: this.datasetData.uuid,
          type: this.datasetData.type,
          filePath: file.ParenDir + file.FileName,
        },
        success: function (res) {
          self.data.fileContentMap[file.ParenDir + file.FileName] = res.data;
          const data = res.data || [];
          if (res.result_code == 0) {
            self.renderText(data.join(''));
          } else if (res.result_code == -1) {
            self.renderText(res.msg);
          } else {
            self.renderNotSupport();
          }
        },
        error: function (err) {
          console.log(err);
          self.renderNotSupport();
        }
      });
    }
  };

  DatasetPreview.prototype.offsetPrview = function (offset) {
    var fileList = this.currentFileList.filter(function (item) { return !item.IsDir });
    if (!fileList.length) {
      this.renderPleseSelectFile();
      return;
    }
    var file = this.currentFile;
    var index = -1;
    var nextIndex = 0;
    if (file) {
      index = fileList.findIndex(function (item) { return (item.ParenDir + item.FileName) == (file.ParenDir + file.FileName); });
      index = Math.max(0, index);
      nextIndex = (offset + index + fileList.length) % fileList.length;
    }
    this.currentFile = fileList[nextIndex];
    var self = this;
    $('#filelist .file-item').removeClass('active');
    $('#filelist .file-item').map((_, item) => {
      var eleObj = $(item);
      var data = eleObj.data('data');
      if ((data.ParenDir + data.FileName) == (self.currentFile.ParenDir + self.currentFile.FileName)) {
        eleObj.addClass('active');
        self.renderPreview(data);
      }
    })
  };

  DatasetPreview.prototype.eventInit = function () {
    var self = this;
    $('#filelist').on('click', '.file-item', function () {
      var eleObj = $(this);
      var file = eleObj.data('data');
      if (file.IsDir) {
        self.nextDirs(file);
      } else {
        $('#filelist .file-item').removeClass('active');
        eleObj.addClass('active');
        self.currentFile = file;
        self.renderPreview(file);
      }
    }).on('click', '.file-more', function () {
      self.getPathChildrenMore();
    });
    $('#file_path_container').on('click', '.section', function () {
      var eleObj = $(this).parent();
      var pathData = eleObj.data('data');
      var index = self.data.filePath.findIndex(function (item) { return item.path == pathData.path });
      var newPath = self.data.filePath.slice(0, index + 1);
      self.data.filePath = newPath;
      self.getPathChildren();
    });
    $('#myCanvas_div').on('click', '.prev_view_btn', function () {
      self.offsetPrview(-1);
    }).on('click', '.next_view_btn', function () {
      self.offsetPrview(1);
    });
  };

  $(document).ready(function () {
    window.DatasetPreviewController = new DatasetPreview();
  })
})();
