import SparkMD5 from "spark-md5";
import initClipboard from "./clipboard.js";
import highlight from "./highlight.js";

; (function () {
  const WAIT_COUNT = 10;
  const exceptPages = ['/home/term'];
  const csrf = window.config.csrf;
  const mdConfigReg = /^\---\n((.|\s)*?)\n---\n/;
  let userID = '';
  let curStorageKey = '';
  const lang = document.querySelector('html').getAttribute('lang');

  function init() {
    userID = $('meta[name="_uid"]').attr('content');
    if (!userID) { // 未登录，不显示
      return;
    }
    curStorageKey = `g-models-${userID}`;
    var pathName = window.location.pathname;
    if (exceptPages.indexOf(pathName) > -1 || window.IsActivatePage || window.IsBindPhonePage) return; // 排除页，不显示
    $.ajax({
      type: "GET",
      url: "/dashboard/invitation",
      dataType: "json",
      data: { filename: `tips/global/change${lang == 'zh-CN' ? '' : '_en'}.md` },
      success: function (res) {
        try {
          res && renderMdStr(res);
        } catch (err) {
          console.log(err);
        }
      },
      error: function (err) {
        console.log(err);
      }
    });
    delete window.globalModalInit;
  }

  function renderMdStr(str) {
    if (!str) return;
    const hash = SparkMD5.hash(str);
    const configs = parseMdConfigs(str);
    str = str.replace(mdConfigReg, '');
    const dataInfoStr = window.localStorage.getItem(curStorageKey) || '{}';
    let dataInfoObj = JSON.parse(dataInfoStr);
    var showOr = dataInfoObj[hash] === false ? false : true;
    dataInfoObj[hash] = showOr;
    configs.hash = hash;
    if (!showOr) return;
    $.ajax({
      type: "POST",
      url: `/api/v1/markdown?_csrf=${csrf}`,
      data: {
        mode: 'gfm',
        text: str,
      },
      success: function (res) {
        try {
          if (res) {
            createDialog(res, configs);
            window.config.HighlightJS = true;
            $('.__g-modal pre code').each((function () {
              highlight(this);
            }));
            setTimeout(() => {
              initClipboard('.__g-modal .clipboard');
              $('.__g-modal .clipboard').popup('hide');
            }, 300);
          }
        } catch (err) {
          console.log(err);
        }
      },
      error: function (err) {
        console.log(err);
      }
    });
  }

  function sparkMD5Hash(str = '') {
    return SparkMD5.hash(str) + Math.random().toString().replace('0.', '');
  }

  function parseMdConfigs(mdStr) {
    const obj = {};
    const regex = mdConfigReg;
    const res = mdStr.match(regex);
    if (res && res[1]) {
      const str = res[1];
      const rows = str.split('\n');
      for (let i = 0, iLen = rows.length; i < iLen; i++) {
        const row = rows[i];
        const col = row.split(':');
        if (col.length > 1) {
          obj[col[0].trim()] = col[1].trim();
        }
      }
    }
    return obj
  }

  function insertCodeCopyBtn(htmlStr) {
    const html = $(htmlStr);
    const codeBlocks = html.find('.code-block');
    for (let i = 0, iLen = codeBlocks.length; i < iLen; i++) {
      const codeBlockI = codeBlocks.eq(i);
      const txt = codeBlockI.text();
      const copyBtn = $(`<div class="copy-btn"><a 
      href="javascript:;" class="ui poping up clipboard" id="clipboard-${sparkMD5Hash(txt)}"
        data-position="top center" data-variation="inverted tiny" data-success="${window.i18n.cloudeBrainMirror.copy_succeeded}"
        data-content="${window.i18n.cloudeBrainMirror.copy}" 
        data-original="${window.i18n.cloudeBrainMirror.copy}" 
        data-clipboard-text=""><i style="font-size:14px;" class="copy outline icon"></i></a></div>`);
      copyBtn.find('a').attr('data-clipboard-text', txt);
      codeBlockI[0].outerHTML = `<div class="code-content">${codeBlockI[0].outerHTML}${copyBtn[0].outerHTML}</div>`;
    }
    return html.html();
  }

  function createDialog(html, configs) {
    const dataInfoStr = window.localStorage.getItem(curStorageKey) || '{}';
    let dataInfoObj = JSON.parse(dataInfoStr);
    const hash = configs.hash;
    function startCount(modelEl, count) {
      var timer = setInterval(function () {
        count--;
        modelEl.data('count', count);
        if (count <= 0) {
          modelEl.find('.button.positive').removeClass('disabled');
          modelEl.find('.count-down-c').hide();
          clearInterval(timer);
        } else {
          modelEl.find('.count-down-c .count-down').text(count);
        }
      }, 1000);
      modelEl.data('timer', timer);
    }
    const renderHtml = insertCodeCopyBtn(html);
    var showOr = dataInfoObj[hash] === false ? false : true;
    dataInfoObj[hash] = showOr;
    if (!showOr || !renderHtml) return;
    var el = $(`
      <div class="ui longer large modal __g-modal" _id="g-modal-${hash}">        
        <div class="header" style="line-height:0.7em;font-size:1.2em;">${configs.title || window.i18n.warmPrompt}</div>
        <style>
          .__g-modal .code-content {position:relative;}
          .__g-modal .code-block {position:relative;}
          .__g-modal .copy-btn {position:absolute;right:4px;top:4px;}
        </style>
        <div class="content scrolling markdown">
          ${renderHtml}
        </div>
        <div class="actions" style="text-align:center;">
          <div style="padding:0.78571429em 0.78571429em 0.78571429em;height:36px;display: inline-block;">
            <div class="ui checkbox" >
              <input type="checkbox" name="notRemindAgain">
              <label>${window.i18n.notRemind}</label>
            </div>
          </div>         
          <div class="ui positive button ${WAIT_COUNT ? 'disabled' : ''}">
            ${window.i18n.close}${WAIT_COUNT ? `<span class="count-down-c"> (<span class="count-down">${WAIT_COUNT}</span>S)</span>` : ''}
          </div>
        </div>
      </div>`);
    $('body').append(el);
    var modelEl = $(`.ui.modal[_id="g-modal-${hash}"]`);
    modelEl.modal({
      closable: false,
      onDeny: function () { return false; },
      onApprove: function (trigger) {
        var modelEl = $(trigger).closest('.__g-modal');
        var count = modelEl.data('count');
        var indexNum = modelEl.data('index-num');
        var notRemindAgain = modelEl.find('input[name="notRemindAgain"]').prop('checked');
        if (Number(count) <= 0) {
          if (notRemindAgain) {
            const dataInfoStr = window.localStorage.getItem(curStorageKey) || '{}';
            let dataInfoObj = JSON.parse(dataInfoStr);
            dataInfoObj[indexNum] = false;
            window.localStorage.setItem(curStorageKey, JSON.stringify(dataInfoObj));
          }
        }
      }
    }).modal('show');
    modelEl.data('index-num', hash);
    startCount(modelEl, WAIT_COUNT);
    window.localStorage.setItem(curStorageKey, JSON.stringify(dataInfoObj));
  }

  window.globalModalInit = init;
  setTimeout(function () {
    if ($('.modal.network-security').hasClass('scale') || $('.modal.network-security').hasClass('active'))
      return;
    init();
  }, 0);
})();
