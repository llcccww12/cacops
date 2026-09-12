export default async function initCloudrainSow() {

  const { csrf } = window.config;

  function paddingZeros(str, len) {
    str = str.toString();
    if (str.length < len) {
      str = new Array(len - str.length).fill('0').join('') + str;
    }
    return str;
  }

  function timeFormat(date) {
    return `${date.getFullYear()}-${paddingZeros(date.getMonth() + 1, 2)}-${paddingZeros(date.getDate(), 2)} ${paddingZeros(date.getHours(), 2)}:${paddingZeros(date.getMinutes(), 2)}:${paddingZeros(date.getSeconds(), 2)}`;
  }

  $(".run_info").on('click', function () {
    let version_name = $(this).data("version");
    let ID = $(`#accordion${version_name}`).data("jobid");
    let repoPath = $(`#accordion${version_name}`).data("repopath");
    $(`#info${version_name} .ui.inverted.active.dimmer`).css({
      "background-color": "#fff",
      display: "block",
    });
    $.get(
      `/api/v1/repos/${repoPath}/${ID}/job_event`,
      (data) => {
        $(`#info${version_name} .ui.inverted.active.dimmer`).css(
          "display", "none",
        );
        parseInfo(data, version_name)
      })
  });

  function parseInfo(jsonObj, version_name) {
    let html = "";
    if (jsonObj != null) {
      let podEventArray = jsonObj['JobEvents'];
      if (podEventArray != null) {
        for (let i = 0; i < podEventArray.length; i++) {
          if (podEventArray[i]["reason"] != "") {
            let time = podEventArray[i]["timestamp"] && new Date(podEventArray[i]["timestamp"])
            html += `<p><b>[${podEventArray[i]["reason"]}]</b> <span>${time.toLocaleString()}</span></p>`
            html += `<p>${podEventArray[i]["message"]}</p>`;
          }
        }
      }
    }
    $(`#info${version_name} .info_text`)[0].innerHTML = html
  }

  /* New Log Start */
  function LogTool() {
    this.$container = null;
    this.version = '';
    this.taskID = '';
    this.logApiUrlUrl = '';
    this.logDownloadUrl = '';
    this.workservernumber = 0;
    this.nodeIndex = '0';
    this.multipleNode = false;
    this.noScroll = false;

    this.content = '';
    this.lines = 60;
    this.order = 'up';
    this.startLine = '';
    this.endLine = '';
    this.scrollTop = 0;
    this.loading = false;
  }

  LogTool.prototype.init = function (container, options) {
    if (!container) return;
    this.$container = container;
    const data = container.data();
    this.version = data.version;
    this.taskID = data.taskid;
    this.logApiUrl = data.logapiurl;
    this.logDownloadUrl = data.logdownloadurl;
    this.multipleNode = data.multiplenode;
    this.noScroll = data.noscroll;
    this.workservernumber = data.workservernumber;

    if (options) {
      if (options.lines) this.lines = options.lines;
    }
    this.eventInit();
    this.refresh();
  };

  LogTool.prototype.escapeHTML = function (str) {
    return str.toString().replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;").replace(/'/g, "&apos;");
  };

  LogTool.prototype.getLogs = function (callback) {
    const _this = this;
    _this.loading = true;
    _this.$container.find('.log-scroll .dimmer').show();
    let node_id = '';
    if (this.multipleNode) {
      node_id = this.nodeIndex;
    }
    $.ajax({
      url: `/api/v1${_this.logApiUrl}/ai_task/log`,
      type: 'get',
      data: {
        id: this.taskID,
        base_line: this.order == 'up' ? this.startLine : this.endLine,
        lines: this.lines,
        order: this.order,
        _csrf: csrf,
      },
      dataType: 'json',
      success: function (res) {
        _this.loading = false;
        _this.$container.find('.log-scroll .dimmer').hide();
        if (res.code == 0) {
          res = res.data || {};
          if (res && res.can_log_download) {
            _this.$container.find('.log-download').removeClass('disabled');
          }
          if (res.lines > 0) {
            if (_this.order == 'up') {
              const logEle = _this.$container.find('.log-scroll')[0];
              const scrollHeight = logEle.scrollHeight;
              const scrollTop = logEle.scrollTop;
              _this.content = `${_this.escapeHTML(res.content)}` + _this.content;
              _this.$container.find('.log-content').html(`<pre>${_this.content}</pre>`);
              _this.startLine = res.start_line;
              if (_this.endLine) {
                _this.stayScrollPos(scrollHeight, scrollTop);
              }
              if (!_this.endLine) {
                _this.endLine = res.end_line;
              }
            } else if (_this.order == 'down') {
              _this.content += `${_this.escapeHTML(res.content)}`;
              _this.$container.find('.log-content').html(`<pre>${_this.content}</pre>`);
              _this.endLine = res.end_line;
              if (!_this.startLine) {
                _this.startLine = res.start_line;
              }
            }
          } else {
            if (_this.noScroll) {
              _this.content = `${_this.escapeHTML(res.content)}`;
              _this.$container.find('.log-content').html(`<pre>${_this.content}</pre>`);
            } else {
              let msg = '';
              if (!_this.content) { // log content is empty
                msg = _this.order == 'down' ? i18n['scrolled_logs_top']
                  : i18n['scrolled_logs_bottom'];
              } else {
                msg = _this.order == 'up' ? i18n['scrolled_logs_top']
                  : i18n['scrolled_logs_bottom'];
              }
              _this.showMessage(msg);
            }
          }
          callback && callback();
        } else {
          _this.showMessage(res.msg);
        }
      },
      error: function (err) {
        console.log(err);
        _this.loading = false;
        _this.$container.find('.log-scroll .dimmer').hide();
        _this.showMessage('请求错误');
      }
    })
  };

  LogTool.prototype.scrollHandler = function (evt) {
    if (this.loading) return;
    const logEle = this.$container.find('.log-scroll')[0];
    if (logEle) {
      const scrollHeight = logEle.scrollHeight;
      const clientHeight = logEle.clientHeight;
      const scrollTop = logEle.scrollTop;
      if (scrollTop != this.scrollTop) {
        if (scrollTop == 0) {
          this.order = 'up';
          this.getLogs();
        } else if (scrollTop + clientHeight >= scrollHeight - 1) {
          this.order = 'down';
          this.getLogs();
        }
      }
      this.scrollTop = scrollTop;
    }
  };

  LogTool.prototype.stayScrollPos = function (scrollHeight, scrollTop) {
    const logEle = this.$container.find('.log-scroll')[0];
    if (logEle) {
      const scrollHeightNew = logEle.scrollHeight;
      logEle.scrollTo({
        top: scrollTop + (scrollHeightNew - scrollHeight),
        behavior: 'instant',
      });
    }
  };

  LogTool.prototype.goBottom = function () {
    this.refresh();
  };

  LogTool.prototype.goTop = function () {
    this.content = '';
    this.startLine = '';
    this.endLine = '';
    this.order = 'down';
    this.getLogs(() => {
      const logEle = this.$container.find('.log-scroll')[0];
      if (logEle) {
        logEle.scrollTo({
          top: 0,
          behavior: 'smooth',
        });
      }
    });
  };

  LogTool.prototype.refresh = function () {
    this.content = '';
    this.$container.find('.log-content').empty();
    this.startLine = '';
    this.endLine = '';
    this.order = 'up';
    this.getLogs(() => {
      this.scrollBottomAnimation();
    });
  };

  LogTool.prototype.scrollBottomAnimation = function () {
    const logEle = this.$container.find('.log-scroll')[0];
    if (logEle) {
      const scrollHeight = logEle.scrollHeight;
      const clientHeight = logEle.clientHeight;
      logEle.scrollTo({
        top: scrollHeight - clientHeight,
        behavior: 'smooth',
      });
    }
  };

  LogTool.prototype.showMessage = function (msg) {
    this.$container.find('.message .msg').text(msg);
    this.msgTimer && clearTimeout(this.msgTimer);
    this.$container.find('.message').fadeIn();
    this.msgTimer = setTimeout(() => {
      this.$container.find('.message').fadeOut();
    }, 2000)
  };

  LogTool.prototype.showFullscreen = function () {
    const version = this.version
    $(`.ui.modal.full-log-dlg-${version}`).modal({
      closable: false,
      onShow: function () {
        const detailLogContent = $(`.detail-log-fullscreen-content-${version}`);
        if (detailLogContent.length) {
          const logTool = new LogTool();
          logTool.init(detailLogContent, {
            lines: 100,
          });
        }
      },
      onVisible: function () { },
      onHide: function () { },
    }).modal("show");
    $(`.ui.modal.full-log-dlg-${version}`).find('.full-log-dialog-exit').off('click').on('click', function () {
      $(`.ui.modal.full-log-dlg-${version}`).modal("hide");
    });
  };

  LogTool.prototype.eventInit = function () {
    const _this = this;
    this.$container.find('.log-download').attr('href', `${_this.logDownloadUrl}`);
    this.$container.find('.node-select').hide();
    if (this.multipleNode) {
      this.$container.find('.node-select').show();
      this.$container.find('.node-select .ui.dropdown').dropdown({
        values: new Array(this.workservernumber).fill(0).map((item, index) => ({
          name: `${i18n['computeNode']} ${index + 1}`,
          value: index.toString(),
          selected: index === 0,
        })),
        allowTab: false,
      }).dropdown('set value', '0');
      this.$container.find('.log-download').attr('href', `${_this.logDownloadUrl}/${_this.nodeIndex}`);
      this.$container.find('.node-select .ui.dropdown').dropdown({
        allowTab: false,
        onChange: function (value) {
          if (_this.nodeIndex == value) return;
          _this.nodeIndex = value;
          _this.$container.find('.log-download').attr('href', `${_this.logDownloadUrl}/${_this.nodeIndex}`);
          _this.refresh();
        },
      });
    }
    this.$container.find('.log_top').off('click').on('click', () => {
      this.goTop();
    });
    this.$container.find('.log_bottom').off('click').on('click', () => {
      this.goBottom();
    });
    if (!this.noScroll) {
      this.$container.find('.log-scroll').off('scroll').on('scroll', this.scrollHandler.bind(this));
    }
    this.$container.find('.full-log-dialog').off('click').on('click', () => {
      _this.showFullscreen();
    });
  };

  $('.detail-log-tab').on('click', function () {
    const self = $(this);
    const version = self.data('version');
    const detailLogContent = $(`.detail-log-content-${version}`);
    if (detailLogContent.length) {
      const logTool = new LogTool();
      logTool.init(detailLogContent);
    }
  });
  /* New Log End */

  $(".refresh-status").click(function (e) {
    let version_name = $(this).data("version");
    let ID = $(`#accordion${version_name}`).data("jobid");
    let repoPath = $(`#accordion${version_name}`).data("repopath");
    refreshStatusShow(version_name, ID, repoPath);
    e.stopPropagation();
  });
  $(".stop-show-version").click(function (e) {
    const ID = this.dataset.jobid;
    const repoPath = this.dataset.repopath;
    const version_name = this.dataset.version;
    const url = `/api/v1/repos/${repoPath}/${ID}/stop_version`;
    $.post(url, { version_name: version_name }, (data) => {
      if (data.StatusOK === 0) {
        $(`#${version_name}-stop`).removeClass("blue");
        $(`#${version_name}-stop`).addClass("disabled");
        refreshStatusShow(version_name, ID, repoPath);
      }
    }).fail(function (err) {
      console.log(err);
    });
    e.stopPropagation();
  });
  $(".delete-show-version").click(function (e) {
    const ID = this.dataset.jobid;
    const repoPath = this.dataset.repopath;
    const version_name = this.dataset.version;
    const url = `/api/v1/repos/${repoPath}/${ID}/del_version`;
    $(".ui.basic.modal")
      .modal({
        onApprove: function () {
          $.post(url, { version_name: version_name }, (data) => {
            if (data.StatusOK === 0) {
              if (data.VersionListCount === 0) {
                location.href = `/${repoPath}`;
              } else {
                $("#accordion" + version_name).remove();
              }
              refreshStatusShow(version_name, ID, repoPath);
            } else {
              setTimeout(() => {
                $(".alert")
                  .html(data.Message)
                  .removeClass("alert-success")
                  .addClass("alert-danger")
                  .show()
                  .delay(1500)
                  .fadeOut();
              }, 520);
              return;
            }
          }).fail(function (err) {
            console.log(err);
          });
        },
      })
      .modal("show");

    e.stopPropagation();
  });
  let initShowExportDataset = true
  let initShowNoDataset = true
  let canExportDataset = true
  let fileList = []
  let timer = null
  let last_version = ''
  let datasetID = ''
  $('.ui.accordion .export-dataset').on('click', function (e) {
    const version_name = this.dataset.version;
    const jobId = this.dataset.jobid;
    const repoPath = this.dataset.repopath;
    const Fileurl = `${repoPath}/getmodelfile`
    const dataUrl = `${repoPath}/getcurrentdataset`
    const exportUrl = `${repoPath}/export_exist_dataset`
    const getProgressUrl = `${repoPath}/getprogress`

    if (!initShowExportDataset && last_version === version_name) {
      if (initShowNoDataset) {
        $(".ui.export_dataset.modal").modal({
          onShow: function () {
            $('.ui.dimmer').css({ "background-color": "rgb(136, 136, 136,0.7)" })
            $(".ui.export_dataset.modal .cancel").on('click', function () {
              $(".ui.export_dataset.modal").modal('hide')
            })
          },
          onHide: function (params) {
            $('.ui.modal.export_dataset #container').off("click")
            $('#export-dataset-select').off("click")
            $('#export-dataset-type').off("click")
            $('.ui.export_dataset.modal .error.message').text('').hide()
          },
          onApprove: function () {
            const modelFileSelectEle = $(".ui.export_dataset.modal #export-dataset-file .items").find('.file_item')
            if (modelFileSelectEle.length !== 0 && canExportDataset) {
              modelFileSelectEle.each(function (index) {
                fileList.push($(this).attr('data-index'))
              })
              const type = Number($('.ui.modal.export_dataset input[name="type"]').val())
              const csrf = $('.ui.modal.export_dataset input[name="_csrf"]').val()
              const desc = $('.ui.modal.export_dataset textarea[name="description"]').val() //,description:desc
              let params = { _csrf: csrf, jobId: jobId, versionName: version_name, datasetId: datasetID, modelSelectedFile: fileList.join(';'), type: type, description: desc }
              postExportDataset(exportUrl, params, getProgressUrl)
            } else {
              $('.ui.export_dataset.modal .error.message').text(`${i18n['exportDataset']['please_select_file']}`).show()
            }
            return false;
          },
        }).modal('show').modal('setting', 'closable', false)
      } else {
        $(".ui.no_export_dataset.modal").modal({
          onShow: function () {
            $('.ui.dimmer').css({ "background-color": "rgb(136, 136, 136,0.7)" })
          },
        }).modal("show")
      }

    } else {
      $(`.ui.accordion #${version_name}-export-dataset .export-popup`).show()
      $('.ui.export_dataset.modal .error.message').text('').hide()
      $.get(dataUrl, (data) => {
        initShowExportDataset = false
        last_version = version_name
        if (data.code === 0) {
          datasetID = data.dataset.ID
          getModelFileList(Fileurl, version_name, jobId)
          getInitEXportDataset(getProgressUrl, data.dataset.ID, jobId, version_name)
          $(".ui.export_dataset.modal").modal({
            onApprove: function () {
              const modelFileSelectEle = $(".ui.export_dataset.modal #export-dataset-file .items").find('.file_item')
              if (modelFileSelectEle.length !== 0 && canExportDataset) {
                modelFileSelectEle.each(function (index) {
                  fileList.push($(this).attr('data-index'))
                })
                const type = Number($('.ui.modal.export_dataset input[name="type"]').val())
                const csrf = $('.ui.modal.export_dataset input[name="_csrf"]').val()
                const desc = $('.ui.modal.export_dataset textarea[name="description"]').val() //,description:desc
                let params = { _csrf: csrf, jobId: jobId, versionName: version_name, datasetId: data.dataset.ID, modelSelectedFile: fileList.join(';'), type: type, description: desc }
                postExportDataset(exportUrl, params, getProgressUrl)
              } else {
                $('.ui.export_dataset.modal .error.message').text(`${i18n['exportDataset']['please_select_file']}`).show()
              }
              return false;
            },
            onShow: function () {
              $(`.ui.accordion #${version_name}-export-dataset .export-popup`).hide()
              $('.ui.dimmer').css({ "background-color": "rgb(136, 136, 136,0.7)" })
            },
            onHide: function () {
              $('.ui.modal.export_dataset #container').off("click")
              $('#export-dataset-select').off("click")
              $('#export-dataset-type').off("click")
              $('.ui.export_dataset.modal .error.message').text('').hide()
            }
          })
            .modal("show")
            .modal('setting', 'closable', false)
        } else {
          $(".ui.no_export_dataset.modal").modal({
            onShow: function () {
              $(`.ui.accordion #${version_name}-export-dataset .export-popup`).hide()
              $('.ui.dimmer').css({ "background-color": "rgb(136, 136, 136,0.7)" })
            },
          }).modal("show")
          initShowNoDataset = false
        }
      })
    }

    $('#export-dataset-select').on('click', function () {
      $(this).find('#model-file-wrap').show()
    })
    $('#model-file-export').on('click', '.delete.icon', function () {
      let fileEle = $(this).siblings('span').text()
      $(this).parent().remove()
      const $parentCheckbox = $('#model-file-result').find(`input[name="${fileEle}"]`).parent()
      $parentCheckbox.checkbox('set unchecked')
    })
    $('.ui.modal.export_dataset #container').on('click', function (e) {
      if ($(e.target).closest('#model-file-wrap').length === 0 && $(e.target).closest('#export-dataset-select').length !== 1) {
        $(this).find('#model-file-wrap').hide()
      }
    })
    $('#export-dataset-type').on('click', function (e, arg1) {
      document.querySelectorAll('#export-dataset-type a').forEach((item) => {
        item.classList.remove('active')
      })
      if (arg1) {
        $('#export-dataset-type a')[arg1].classList.add('active')
        document.querySelector('input[name="type"]').value = arg1
      } else {
        e.target.classList.add('active')
        document.querySelector('input[name="type"]').value = e.target.dataset.type
      }

    })
    e.stopPropagation();
  })
  function getInitEXportDataset(getProgressUrl, datasetId, jobId, version_name) {
    let setIntervalFlag = false
    $('.ui.modal.export_dataset #model-file-export').empty()
    $.get(getProgressUrl, { progressId: `${datasetId}_${jobId}_${version_name}` }, (data) => {
      const result = data && JSON.parse(data)
      if (Object.keys(result).length > 0) {
        canExportDataset = false
        let fileInitList = Object.keys(result).filter((item) => item !== '##type##')
        $($('#export-dataset-type')).trigger('click', `${result['##type##']}`)

        fileInitList.forEach((item) => {
          if (result[item] === -1) {
            let itemHtml = `<div data-index="${item}" class="file_item">
                            <span class="nowrap" style="width:80%" title="${item}">${item}</span>
                            <div class="file_error"><i class="ri-close-circle-line failed"></i><span>${i18n['exportDataset']['export_failed']}</span></div> 
                          </div>`
            $('.ui.modal.export_dataset #model-file-export').append(itemHtml)
          }
          if (result[item] === -2) {
            let itemHtml = `<div data-index="${item}" class="file_item">
                            <span class="nowrap" style="width:80%" title="${item}">${item}</span>
                            <div class="file_error"">
                              <i class="ri-close-circle-line failed"></i>
                              <span>${i18n['exportDataset']['export_failed']}</span>  
                              <span data-tooltip="${i18n['exportDataset']['export_has_same_file']}" data-inverted="" data-variation="tiny">
                                <i class="ri-question-fill question"></i>
                              </span>
                            </div>
                          </div>`
            $('.ui.modal.export_dataset #model-file-export').append(itemHtml)
          }
          if (result[item] === 100) {
            let itemHtml = `<div data-index="${item}" class="file_item">
                            <span class="nowrap" style="width:80%" title="${item}">${item}</span>
                            <div class="file_success"">
                              <i class="ri-checkbox-circle-line success" style="vertical-align: middle;"></i>
                              <span>${i18n['exportDataset']['export_success']}</span>
                            </div>
                          </div>`
            $('.ui.modal.export_dataset #model-file-export').append(itemHtml)
          }
          if (result[item] === 0) {
            let itemHtml = `<div data-index="${item}" class="file_item">
                              <span class="nowrap" style="width:80%" title="${item}">${item}</span>
                              <div class="file_wait""><i class="ri-loader-2-line waiting spin"></i><span>${i18n['exportDataset']['exporting']}</span></div>
                            </div>`
            $('.ui.modal.export_dataset #model-file-export').append(itemHtml)
            setIntervalFlag = true
          }
        })
        if (setIntervalFlag) {
          timer && clearInterval(timer)
          timer = setIntervalImmediately(getProgress, 5000, getProgressUrl, `${datasetId}_${jobId}`)
        }
      }

    })
  }
  function getModelFileList(Fileurl, version_name, jobId) {
    $.get(Fileurl, { versionName: version_name, jobId: jobId }, (data) => {
      $('.ui.modal.export_dataset #model-file-result').empty()
      let html = ''
      if (data.code === 0 && data.files.length !== 0) {
        let dataFileList = data.files.sort((a, b) => {
          return a.FileName.localeCompare(b.FileName)
        })
        dataFileList.forEach(element => {
          html += `<div class="item">
                      <div class="ui child checkbox">
                        <input type="checkbox" name="${element.FileName}">
                        <label>${element.FileName}</label>
                      </div>
                    </div>`
        });
        $('.ui.modal.export_dataset #model-file-result').append(html)
        $('#model-file-result.list .child.checkbox').checkbox({
          onChecked: function () {
            $('.ui.export_dataset.modal .error.message').text('').hide()
            if (!canExportDataset) {
              $('.ui.modal.export_dataset #model-file-export').empty()
              canExportDataset = true
            }
            let itemHtml = ''
            let fileName = $(this).attr('name')
            itemHtml += `<div data-index="${fileName}" class="file_item">
                          <span class="nowrap" style="width:80%" title="${fileName}">${fileName}</span>
                          <i class="delete icon" style="cursor:pointer"></i>
                        </div>`
            $('.ui.modal.export_dataset #model-file-export').append(itemHtml)

          },
          onUnchecked: function () {
            let fileName = $(this).attr('name')
            $('.ui.modal.export_dataset #model-file-export').find(`div[data-index="${fileName}"]`).remove()

          },
        })
      }
    }).fail(function (err) {
      console.log(err);
    });

  }
  function getProgress(getProgressUrl, progressId) {
    const $statusEle = $('#model-file-export .file_item').find('div')
    let fileLength = $statusEle.length
    let count = 0
    $.get(getProgressUrl, { progressId: progressId }, (data) => {
      const result = data && JSON.parse(data)
      let sortResult = Object.keys(result).sort((a, b) => {
        return fileList.indexOf(a) - fileList.indexOf(b)
      })
      let filterResult = sortResult.filter((item) => item !== '##type##')
      filterResult.forEach((item, index) => {
        console.log(item)
        if (result[item] === -1) {

          if (!$($statusEle[index]).hasClass('file_error')) {
            $($statusEle[index]).replaceWith(`<div class="file_error"><i class="ri-close-circle-line failed"></i><span>${i18n['exportDataset']['export_failed']}</span></div>`)
          }
          count++
        }
        if (result[item] === -2) {
          if (!$($statusEle[index]).hasClass('file_error')) {
            $($statusEle[index]).replaceWith(`<div class="file_error""><i class="ri-close-circle-line failed"></i><span>${i18n['exportDataset']['export_failed']}</span><span data-tooltip="${i18n['exportDataset']['export_has_same_file']}" data-inverted="" data-variation="tiny"><i class="ri-question-fill question"></i></span></div>`)
          }
          count++
        }
        if (result[item] === 100) {
          if (!$($statusEle[index]).hasClass('file_success')) {
            $($statusEle[index]).replaceWith(`<div class="file_success""><i class="ri-checkbox-circle-line success" style="vertical-align: middle;"></i><span>${i18n['exportDataset']['export_success']}</span></div>`)
          }
          count++
        }
        console.log("count:", count)
        if (count === fileLength) {
          $(".ui.export_dataset.modal").modal('refresh')
          timer && clearInterval(timer)
          $('#export-dataset-select').on('click', function () {
            $(this).find('#model-file-wrap').show()
          })
          $('.ui.modal.export_dataset .ui.approve').removeClass('disabled')
          fileList = []
          canExportDataset = false
          $('#model-file-result.list .child.checkbox.checked').each(function () {
            $(this).checkbox('set unchecked')
          })
          return
        }

      })
    })
  }
  function postExportDataset(url, params, getProgressUrl) {
    $.post(url, params, (data) => {
      if (data.code === '0') {
        $('#model-file-export .delete.icon').each(function () {
          $(this).replaceWith(`<div class="file_wait""><i class="ri-loader-2-line waiting spin"></i><span>${i18n['exportDataset']['exporting']}</span></div>`)
        })
        $('.ui.modal.export_dataset .ui.approve').addClass('disabled')
        $('#export-dataset-select').off("click")
        // $('.ui.modal.export_dataset').off("click")
        getProgress(getProgressUrl, data.progressId)

        timer && clearInterval(timer)
        timer = setIntervalImmediately(getProgress, 5000, getProgressUrl, data.progressId)
      }

    })
  }
  function setIntervalImmediately(func, interval, ...args) {
    func(...args)
    return setInterval(func, interval, ...args)
  };
  // $('.ui.pointing.secondary.menu .item:eq(0)').click(function(e) {
  //   const self = $(this);
  //   setTimeout(function() {
  //     self.closest('.accordion').find('.refresh-status').trigger('click');
  //   }, 20);

  // });
  function refreshStatusShow(version_name, ID, repoPath) {
    $.get(
      `/api/v1/repos/${repoPath}/${ID}?version_name=${version_name}`,
      (data) => {
        //accroding下的状态
        $(`#${version_name}-status-span span`).text(data.JobStatus);
        //accroding下的状态图标
        $(`#${version_name}-status-span i`).attr("class", data.JobStatus);
        //accroding下的运行时长
        $(`#${version_name}-duration-span`).text(data.JobDuration);
        //配置信息详情页的状态
        data.StartTime !== undefined && data.StartTime > 0 && $(`#${version_name}-startTime`).text(timeFormat(new Date(data.StartTime * 1000)));
        //配置信息详情页的状态
        $(`#${version_name}-status`).text(data.JobStatus);
        //配置信息详情页的状态
        $(`#${version_name}-duration`).text(data.JobDuration);
        //配置信息详情页的状态
        $(`#${version_name}-ai_center`).text(data.AiCenter);
      }
    ).fail(function (err) {
      console.log(err);
    });
    const accordionEl = $(`#accordion${version_name}`);
    const activeTab = accordionEl.find('.ui.pointing.secondary.menu .item:not(:eq(0)).active');
    activeTab.trigger('click');
  }
  //

  $(".content-pad").on("click", ".load-model-file", function () {
    let downloadFlag = $(this).data("download-flag") || "";
    let gpuFlag = $(this).data("gpu-flag") || "";
    let version_name = $(this).data("version");
    let parents = $(this).data("parents");
    let filename = $(this).data("filename");
    let init = $(this).data("init") || "";
    let path = $(this).data("path");
    let retryPath = `/api/v1/repos${$(this).data("retry-path")}`;
    const rescheduleFlag = $(this).data("can-reschedule") || "";
    $(`#dir_list${version_name}`).empty();
    let url = `/api/v1/repos${path}?version_name=${version_name}&parentDir=${parents}`;
    $.get(url, (data) => {

      if (data.StatusOK == 0) { // 成功 0
        if (data.Dirs) {
          data.Dirs.length !== 0 && $(`#${version_name}-result-down`).show()
          renderDir(path, data, version_name, downloadFlag, gpuFlag);
        }
        if (init === "init") {
          $(`input[name=model${version_name}]`).val("");
          $(`input[name=modelback${version_name}]`).val(version_name);
          $(`#file_breadcrumb${version_name}`).empty();
          let htmlBread = "";
          if (version_name) {
            htmlBread += `<div class='active section'>${version_name}</div>`;
          } else {
            htmlBread += `<div class='active section'>result</div>`;
          }
          htmlBread += "<div class='divider'> / </div>";
          $(`#file_breadcrumb${version_name}`).append(htmlBread);
        } else {
          renderBrend(
            this,
            path,
            version_name,
            parents,
            filename,
            init,
            downloadFlag,
            gpuFlag
          );
        }
      } else if (data.StatusOK == -1) { // 任务未结束 -1
        $(`#file_breadcrumb${version_name}`).empty();
        $(`#dir_list${version_name}`).html(`<div style="height:200px;display:flex;justify-content:center;align-items:center;font-size:14px;color:rgb(16, 16, 16);">          
          <div style="display:flex;justify-content:center;align-items:center;height:24px;width:24px;margin-right:5px;">
            <svg xmlns="http://www.w3.org/2000/svg" class="styles__StyledSVGIconPathComponent-sc-16fsqc8-0 iKfgJk svg-icon-path-icon fill" viewBox="0 0 32 32" width="16" height="16"><defs data-reactroot=""></defs><g><path d="M16 29.333c-7.364 0-13.333-5.969-13.333-13.333s5.969-13.333 13.333-13.333 13.333 5.969 13.333 13.333-5.969 13.333-13.333 13.333zM16 26.667c5.891 0 10.667-4.776 10.667-10.667s-4.776-10.667-10.667-10.667v0c-5.891 0-10.667 4.776-10.667 10.667s4.776 10.667 10.667 10.667v0zM17.333 16h5.333v2.667h-8v-9.333h2.667v6.667z"></path></g></svg>
            </div>
          <span>${i18n['task_not_finished']}</span>
        </div>`);
      } else if (data.StatusOK == 1) { // 处理中 1
        $(`#file_breadcrumb${version_name}`).empty();
        $(`#dir_list${version_name}`).html(`<div style="height:200px;display:flex;justify-content:center;align-items:center;font-size:14px;color:rgb(16, 16, 16);">
          <style>
          @-webkit-keyframes spinning {
            0% { -webkit-transform: rotate(0deg); }
            100% { -webkit-transform: rotate(360deg); }
          }          
          </style>
          <div style="display:flex;justify-content:center;align-items:center;height:24px;width:24px;margin-right:5px;animation-duration:3s;animation-iteration-count:infinite;animation-name:spinning;animation-timing-function:linear;animation-fill-mode:backwards;">
            <svg xmlns="http://www.w3.org/2000/svg" class="styles__StyledSVGIconPathComponent-sc-16fsqc8-0 iKfgJk svg-icon-path-icon fill" viewBox="0 0 48 48" width="16" height="16"><defs data-reactroot=""></defs><g><g><rect width="48" height="48" fill="white" fill-opacity="0.01" stroke-linejoin="round" stroke-width="4" stroke="none" fill-rule="evenodd"></rect><g transform="translate(7.000000, 3.500000)"><path d="M0,0.5 L34,0.5" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" stroke="#333" fill="none" fill-rule="evenodd"></path><path d="M0,40.5 L34,40.5" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" stroke="#333" fill="none" fill-rule="evenodd"></path><path d="M30,20.4999965 C27.3333333,33.8388874 23,40.5055541 17,40.4999965 C11,40.494439 6.66666667,33.8277723 4,20.4999965 L30,20.4999965 Z" fill="none" fill-rule="nonzero" transform="translate(17.000000, 30.499998) rotate(180.000000) translate(-17.000000, -30.499998) " stroke-linejoin="round" stroke-width="4" stroke="#333"></path><path d="M30,0.5 C27.3333333,13.8388909 23,20.5055575 17,20.5 C11,20.4944425 6.66666667,13.8277758 4,0.5 L30,0.5 Z" fill="none" fill-rule="nonzero" stroke-linejoin="round" stroke-width="4" stroke="#333"></path></g></g></g></svg>
          </div>          
          <span>${i18n['file_sync_ing']}</span>
        </div>`);
      } else if (data.StatusOK == 2) { // 失败 2
        $(`#file_breadcrumb${version_name}`).empty();
        if (rescheduleFlag) {
          $(`#dir_list${version_name}`).html(`<div style="height:200px;display:flex;justify-content:center;align-items:center;font-size:14px;color:rgb(16, 16, 16);">          
            <div style="display:flex;justify-content:center;align-items:center;height:24px;width:24px;margin-right:5px;">
            <svg xmlns="http://www.w3.org/2000/svg" class="styles__StyledSVGIconPathComponent-sc-16fsqc8-0 iKfgJk svg-icon-path-icon fill" viewBox="64 64 896 896" width="16" height="16"><defs data-reactroot=""></defs><g><path d="M464 720a48 48 0 1 0 96 0 48 48 0 1 0-96 0zm16-304v184c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V416c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8zm475.7 440l-416-720c-6.2-10.7-16.9-16-27.7-16s-21.6 5.3-27.7 16l-416 720C56 877.4 71.4 904 96 904h832c24.6 0 40-26.6 27.7-48zm-783.5-27.9L512 239.9l339.8 588.2H172.2z"></path></g></svg>
            </div>
            <span>${i18n['file_sync_fail']}</span>
            <a href="javascript:void(0)" id="retry_result" style='text-decoration: underline;margin-left:0.5rem'>${i18n['retrieve_results']}</a>
          </div>`);
        }
        else {
          $(`#dir_list${version_name}`).html(`<div style="height:200px;display:flex;justify-content:center;align-items:center;font-size:14px;color:rgb(16, 16, 16);">          
              <div style="display:flex;justify-content:center;align-items:center;height:24px;width:24px;margin-right:5px;">
              <svg xmlns="http://www.w3.org/2000/svg" class="styles__StyledSVGIconPathComponent-sc-16fsqc8-0 iKfgJk svg-icon-path-icon fill" viewBox="64 64 896 896" width="16" height="16"><defs data-reactroot=""></defs><g><path d="M464 720a48 48 0 1 0 96 0 48 48 0 1 0-96 0zm16-304v184c0 4.4 3.6 8 8 8h48c4.4 0 8-3.6 8-8V416c0-4.4-3.6-8-8-8h-48c-4.4 0-8 3.6-8 8zm475.7 440l-416-720c-6.2-10.7-16.9-16-27.7-16s-21.6 5.3-27.7 16l-416 720C56 877.4 71.4 904 96 904h832c24.6 0 40-26.6 27.7-48zm-783.5-27.9L512 239.9l339.8 588.2H172.2z"></path></g></svg>
              </div>
              <span>${i18n['file_sync_fail']}</span>
            </div>`);
        }
      } else if (data.StatusOK == 3) { // 等待同步 3
        $(`#file_breadcrumb${version_name}`).empty();
        $(`#dir_list${version_name}`).html(`<div style="height:200px;display:flex;justify-content:center;align-items:center;font-size:14px;color:rgb(16, 16, 16);">          
          <div style="display:flex;justify-content:center;align-items:center;height:24px;width:24px;margin-right:5px;">
            <svg xmlns="http://www.w3.org/2000/svg" class="styles__StyledSVGIconPathComponent-sc-16fsqc8-0 iKfgJk svg-icon-path-icon fill" viewBox="0 0 32 32" width="16" height="16"><defs data-reactroot=""></defs><g><path d="M16 29.333c-7.364 0-13.333-5.969-13.333-13.333s5.969-13.333 13.333-13.333 13.333 5.969 13.333 13.333-5.969 13.333-13.333 13.333zM16 26.667c5.891 0 10.667-4.776 10.667-10.667s-4.776-10.667-10.667-10.667v0c-5.891 0-10.667 4.776-10.667 10.667s4.776 10.667 10.667 10.667v0zM17.333 16h5.333v2.667h-8v-9.333h2.667v6.667z"></path></g></svg>
          </div>
          <span>${i18n['file_sync_wait']}</span>
        </div>`);
      } else if (data.StatusOK == 4) { // 无文件 4
        $(`#file_breadcrumb${version_name}`).empty();
        $(`#dir_list${version_name}`).html(`<div style="height:200px;display:flex;justify-content:center;align-items:center;font-size:14px;color:rgb(16, 16, 16);">          
          <div style="display:flex;justify-content:center;align-items:center;height:24px;width:24px;margin-right:5px;">
            <svg xmlns="http://www.w3.org/2000/svg" class="styles__StyledSVGIconPathComponent-sc-16fsqc8-0 iKfgJk svg-icon-path-icon fill" viewBox="0 0 24 24" width="16" height="16"><defs data-reactroot=""></defs><g><circle cx="15.5" cy="9.5" r="1.5"></circle><circle cx="8.5" cy="9.5" r="1.5"></circle><path d="M11.99 2C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zM12 20c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8zm0-6c-2.33 0-4.32 1.45-5.12 3.5h1.67c.69-1.19 1.97-2 3.45-2s2.75.81 3.45 2h1.67c-.8-2.05-2.79-3.5-5.12-3.5z"></path></g></svg>
          </div>
          <span>${i18n['no_file_to_download']}</span>
        </div>`);
      }
      $('#retry_result').on('click', function () {
        $.post(retryPath, (data) => {
          if (data.code === 0) {
            $('.load-model-file').trigger('click');
          }
        }).fail(function (err) {
          console.log(err);
        });
      })
    }).fail(function (err) {
      console.log(err, version_name);
    });
  });


  function renderSize(value) {
    if (null == value || value == "") {
      return "0 Bytes";
    }
    var unitArr = new Array(
      "Bytes",
      "KB",
      "MB",
      "GB",
      "TB",
      "PB",
      "EB",
      "ZB",
      "YB"
    );
    var index = 0;
    var srcsize = parseFloat(value);
    index = Math.floor(Math.log(srcsize) / Math.log(1024));
    var size = srcsize / Math.pow(1024, index);
    size = size.toFixed(0); //保留的小数位数
    return size + unitArr[index];
  }
  function renderBrend(
    that,
    path,
    version_name,
    parents,
    filename,
    init,
    downloadFlag,
    gpuFlag
  ) {
    if (init == "folder") {
      let htmlBrend = "";
      let sectionName = $(
        `#file_breadcrumb${version_name} .active.section`
      ).text();
      let parents1 = $(`input[name=model${version_name}]`).val();
      let filename1 = $(`input[name=modelback${version_name}]`).val();
      if (parents1 === "") {
        $(`#file_breadcrumb${version_name} .active.section`).replaceWith(
          `<a class='section load-model-file' data-download-flag='${downloadFlag}' data-gpu-flag='${gpuFlag}' data-path='${path}' data-version='${version_name}' data-parents='${parents1}' data-filename='' data-init='init'>${sectionName}</a>`
        );
      } else {
        $(`#file_breadcrumb${version_name} .active.section`).replaceWith(
          `<a class='section load-model-file' data-download-flag='${downloadFlag}' data-gpu-flag='${gpuFlag}' data-path='${path}' data-version='${version_name}' data-parents='${parents1}' data-filename='${filename1}'>${sectionName}</a>`
        );
      }

      htmlBrend += `<div class='active section'>${filename}</div>`;
      htmlBrend += "<div class='divider'> / </div>";
      $(`#file_breadcrumb${version_name}`).append(htmlBrend);
      $(`input[name=model${version_name}]`).val(parents);
      $(`input[name=modelback${version_name}]`).val(filename);
    } else {
      $(`input[name=model${version_name}]`).val(parents);
      $(`input[name=modelback${version_name}]`).val(filename);
      $(that).nextAll().remove();
      $(that).after("<div class='divider'> / </div>");
      $(that).replaceWith(`<div class='active section'>${filename}</div>`);
    }
  }

  function renderDir(path, data, version_name, downloadFlag, gpuFlag) {
    let html = "";
    html += "<div class='ui grid' style='margin:0;'>";
    html += "<div class='row' style='padding: 0;'>";
    html += "<div class='ui sixteen wide column' style='padding:1rem;'>";
    html += "<div class='dir list'>";
    html += "<table id='repo-files-table' class='ui single line table pad20'>";
    html += "<tbody>";
    // html += "</tbody>"
    for (let i = 0; i < data.Dirs.length; i++) {
      let dirs_size = renderSize(data.Dirs[i].Size);
      html += "<tr>";
      html += "<td class='name six wid'>";
      html += "<span class='truncate'>";
      html += "<span class='octicon octicon-file-directory'>";
      html += "</span>";
      if (data.Dirs[i].IsDir) {
        html += `<a  class='load-model-file' data-download-flag='${downloadFlag}' data-gpu-flag='${gpuFlag}' data-path='${path}' data-version='${version_name}' data-parents='${data.Dirs[i].ParenDir}' data-filename='${data.Dirs[i].FileName}' data-init='folder'>`;
        html +=
          "<span class='fitted'><i class='folder icon' width='16' height='16' aria-hidden='true'></i>" +
          data.Dirs[i].FileName +
          "</span>";
      } else {
        if (downloadFlag) {
          if (gpuFlag) {
            if (path.includes("model_list")) {
              html += `<a href="${location.href}/download_model?version_name=${version_name}&fileName=${data.Dirs[i].FileName}&parentDir=${data.Dirs[i].ParenDir}&jobName=${data.task.JobName}">`;
            } else {
              html += `<a href="${location.href}/result_download?version_name=${version_name}&fileName=${data.Dirs[i].FileName}&parentDir=${data.Dirs[i].ParenDir}&jobName=${data.task.JobName}">`;
            }
          } else {
            if (path.includes("model_list")) {
              html += `<a href="${location.href}/model_download?version_name=${version_name}&file_name=${data.Dirs[i].FileName}&parent_dir=${data.Dirs[i].ParenDir}">`;
            } else {
              html += `<a href="${location.href}/result_download?version_name=${version_name}&file_name=${data.Dirs[i].FileName}&parent_dir=${data.Dirs[i].ParenDir}">`;
            }
          }
        } else {
          html += `<a class="disabled">`;
        }
        html +=
          "<span class='fitted'><i class='file icon' width='16' height='16' aria-hidden='true'></i>" +
          data.Dirs[i].FileName +
          "</span>";
      }
      html += "</a>";
      html += "</span>";
      html += "</td>";
      html += "<td class='message1 seven wide'>";
      if (data.Dirs[i].IsDir) {
        html += "<span class='truncate has-emoji'></span>";
      } else {
        html +=
          "<span class='truncate has-emoji'>" + `${dirs_size}` + "</span>";
      }

      html += "</td>";

      html += "<td class='text right age three wide'>";
      html +=
        "<span class='truncate has-emoji'>" + data.Dirs[i].ModTime + "</span>";
      html += "</td>";
      html += "</tr>";
    }
    html += "</tbody>";
    html += "</table>";
    html += "</div>";
    html += "</div>";
    html += "</div>";
    html += "</div>";
    $(`#dir_list${version_name}`).append(html);
  }
}
