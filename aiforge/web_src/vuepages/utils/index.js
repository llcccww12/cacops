
import { i18n } from '~/langs';
import { ACC_CARD_TYPE } from '~/const';
import ClipboardJS from 'clipboard';

export const getListValueWithKey = (list, key, k = 'k', v = 'v') => {
  for (let i = 0, iLen = list.length; i < iLen; i++) {
    const listI = list[i];
    if (listI[k] === key) return listI[v];
  }
  return key;
};

export const getUrlSearchParams = () => {
  const params = new URLSearchParams(location.search);
  const obj = {};
  params.forEach((value, key) => {
    obj[key] = value;
  });
  return obj;
};

export const uuidv4 = () => {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
    var r = Math.random() * 16 | 0,
      v = c === 'x' ? r : r & 0x3 | 0x8;
    return v.toString(16);
  });
};

export const transFileSize = (srcSize) => {
  if (null == srcSize || srcSize == '') {
    return '0 Bytes';
  }
  const unitArr = ['Bytes', 'KiB', 'MiB', 'GiB', 'TiB', 'PB', 'EB', 'ZB', 'YB'];
  srcSize = parseFloat(srcSize);
  const index = Math.floor(Math.log(srcSize) / Math.log(1024));
  let size
  if ((srcSize % Math.pow(1024, index)) !== 0) {
    size = (srcSize / Math.pow(1024, index)).toFixed(2);
  } else {
    size = (srcSize / Math.pow(1024, index))
  }
  return size + ' ' + unitArr[index];
};

export const renderSpecStr = (spec, showPoint) => {
  if (!spec) return '';
  var ngpu = `${spec.ComputeResource}: ${spec.AccCardsNum + '*' + getListValueWithKey(ACC_CARD_TYPE, spec.AccCardType)}`;
  var gpuMemStr = spec.GPUMemGiB != 0 ? `(${i18n.t('resourcesManagement.gpuMem')}: ${spec.GPUMemGiB}GB)` : '';
  var memStr = spec.MemGiB != 0 ? `, ${i18n.t('resourcesManagement.mem')}: ${spec.MemGiB}GB` : '';
  var sharedMemStr = spec.ShareMemGiB != 0 ? `, ${i18n.t('resourcesManagement.shareMem')}: ${spec.ShareMemGiB}GB` : '';
  var pointStr = showPoint ? `, ${spec.UnitPrice == 0 ? i18n.t('resourcesManagement.free') : spec.UnitPrice.toFixed(2) + i18n.t('resourcesManagement.point_hr')}` : '';
  var specStr = `${ngpu}${gpuMemStr}, CPU: ${spec.CpuCores}${memStr}${sharedMemStr}${pointStr}`;
  return specStr;
};

export const renderSpecObject = (spec, showPoint) => {
  if (!spec) {
    return {
      type: '',
      specStr: '',
      pointStr: '',
    };
  }
  var ngpu = `${spec.compute_resource}: ${spec.acc_cards_num + '*' + getListValueWithKey(ACC_CARD_TYPE, spec.acc_card_type)}`;
  var gpuMemStr = spec.gpu_mem_gi_b != 0 ? `(${i18n.t('resourcesManagement.gpuMem')}: ${spec.gpu_mem_gi_b}GB)` : '';
  var memStr = spec.mem_gi_b != 0 ? `, ${i18n.t('resourcesManagement.mem')}: ${spec.mem_gi_b}GB` : '';
  var sharedMemStr = spec.share_mem_gi_b != 0 ? `, ${i18n.t('resourcesManagement.shareMem')}: ${spec.share_mem_gi_b}GB` : '';
  var pointStr = showPoint ? `${spec.unit_price == 0 ? i18n.t('resourcesManagement.free') : spec.unit_price.toFixed(2) + i18n.t('resourcesManagement.point_hr')}` : '';
  var specStr = `${ngpu}${gpuMemStr}, CPU: ${spec.cpu_cores}${memStr}${sharedMemStr}`;
  return {
    ...spec,
    id: spec.id.toString(),
    type: spec.compute_resource,
    specStr: specStr,
    pointStr: pointStr,
  };
};

const Minute = 60;
const Hour = 60 * Minute;
const Day = 24 * Hour;
const Week = 7 * Day;
const Month = 30 * Day;
const Year = 12 * Month;

const computeTimeDiff = (diff) => {
  let diffStr = '';
  switch (true) {
    case diff <= 0:
      diff = 0;
      diffStr = i18n.t('timeObj.now');
      break;
    case diff < 2:
      diff = 0;
      diffStr = i18n.t('timeObj.1s');
      break;
    case diff < 1 * Minute:
      diffStr = i18n.t('timeObj.seconds', { msg: Math.floor(diff) });
      diff = 0;
      break;
    case diff < 2 * Minute:
      diff -= 1 * Minute;
      diffStr = i18n.t('timeObj.1m');
      break;
    case diff < 1 * Hour:
      diffStr = i18n.t('timeObj.minutes', { msg: Math.floor(diff / Minute) });
      diff -= diff / Minute * Minute;
      break;
    case diff < 2 * Hour:
      diff -= 1 * Hour;
      diffStr = i18n.t('timeObj.1h');
      break;
    case diff < 1 * Day:
      diffStr = i18n.t('timeObj.hours', { msg: Math.floor(diff / Hour) });
      diff -= diff / Hour * Hour;
      break;
    case diff < 2 * Day:
      diff -= 1 * Day;
      diffStr = i18n.t('timeObj.1d');
      break;
    case diff < 1 * Week:
      diffStr = i18n.t('timeObj.days', { msg: Math.floor(diff / Day) });
      diff -= diff / Day * Day;
      break;
    case diff < 2 * Week:
      diff -= 1 * Week;
      diffStr = i18n.t('timeObj.1w');
      break;
    case diff < 1 * Month:
      diffStr = i18n.t('timeObj.weeks', { msg: Math.floor(diff / Week) });
      diff -= diff / Week * Week;
      break;
    case diff < 2 * Month:
      diff -= 1 * Month;
      diffStr = i18n.t('timeObj.1mon');
      break;
    case diff < 1 * Year:
      diffStr = i18n.t('timeObj.months', { msg: Math.floor(diff / Month) });
      diff -= diff / Month * Month;
      break;
    case diff < 2 * Year:
      diff -= 1 * Year;
      diffStr = i18n.t('timeObj.1y');
      break;
    default:
      diffStr = i18n.t('timeObj.years', { msg: Math.floor(diff / Year) });
      diff -= (diff / Year) * Year;
      break;
  }
  return { diff, diffStr };
};

export const timeSinceUnix = (then, now) => {
  let lbl = 'timeObj.ago';
  let diff = now - then;
  if (then > now) {
    lbl = 'timeObj.from_now';
    diff = then - now;
  }
  if (diff <= 10) {
    return i18n.t('timeObj.now');
  }
  const out = computeTimeDiff(diff);
  return i18n.t(lbl, { msg: out.diffStr });
};

export const setWebpackPublicPath = () => {
  // This sets up webpack's chunk loading to load resources from the 'public'
  // directory. This file must be imported before any lazy-loading is being attempted.
  if (document.currentScript && document.currentScript.src) {
    const url = new URL(document.currentScript.src);
    __webpack_public_path__ = url.pathname.replace(/\/[^/]*?\/[^/]*?$/, '/');
  } else {
    // compat: IE11
    const script = document.querySelector('script[src*="/index.js"]');
    __webpack_public_path__ = script.getAttribute('src').replace(/\/[^/]*?\/[^/]*?$/, '/');
  }

  // rewrite to solve dynamic import cache error
  const AppVer = window.config?.AppVer || Math.random().toString().replace('0.', '');
  const ElementAppendChild = Element.prototype.appendChild;
  Element.prototype.appendChild = function (node) {
    if (node.tagName == 'SCRIPT' && node.src.indexOf('?') < 0) {
      node.src = node.src + '?v=' + AppVer;
    }
    if (node.tagName == 'LINK' && node.href.indexOf('?') < 0) {
      node.href = node.href + '?v=' + AppVer;
    }
    return ElementAppendChild.call(this, node);
  }
};

export const initClipboard = (_els) => {
  const els = _els || document.querySelectorAll(".clipboard");
  if (!els || !els.length) return;
  window.$ && $().popup && $(els).popup();
  const clipboard = new ClipboardJS(els);
  clipboard.on("success", (e) => {
    e.clearSelection();
    const popUpEl = $(e.trigger);
    popUpEl.popup("destroy");
    e.trigger.setAttribute(
      "data-content",
      e.trigger.getAttribute("data-success")
    );
    popUpEl.popup("show");
    e.trigger.setAttribute(
      "data-content",
      e.trigger.getAttribute("data-original")
    );
  });

  clipboard.on("error", (e) => {
    const popUpEl = $(e.trigger);
    popUpEl.popup("destroy");
    e.trigger.setAttribute(
      "data-content",
      e.trigger.getAttribute("data-error")
    );
    popUpEl.popup("show");
    e.trigger.setAttribute(
      "data-content",
      e.trigger.getAttribute("data-original")
    );
  });
  return clipboard;
}

export const escapeHTML = (str) => {
  return str.replace(/[<>&"']/g, function (match) {
    switch (match) {
      case '<': return '&lt;';
      case '>': return '&gt;';
      case '&': return '&amp;';
      case '"': return '&quot;';
      case "'": return '&#39;';
    }
  })
}

export class RandomGenerator {
  // 生成随机数字
  static numbers(length = 4) {
      let result = '';
      for (let i = 0; i < length; i++) {
          result += Math.floor(Math.random() * 10);
      }
      return result;
  }
  
  // 生成随机字母
  static letters(length = 4, uppercase = false) {
      let result = '';
      const chars = uppercase ? 'ABCDEFGHIJKLMNOPQRSTUVWXYZ' : 'abcdefghijklmnopqrstuvwxyz';
      
      for (let i = 0; i < length; i++) {
          result += chars.charAt(Math.floor(Math.random() * chars.length));
      }
      return result;
  }
  
  // 生成字母数字混合
  static alphanumeric(length = 4) {
      let result = '';
      const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
      
      for (let i = 0; i < length; i++) {
          result += chars.charAt(Math.floor(Math.random() * chars.length));
      }
      return result;
  }
  
  // 生成自定义字符集的随机码
  static custom(charset, length = 4) {
      let result = '';
      for (let i = 0; i < length; i++) {
          result += charset.charAt(Math.floor(Math.random() * charset.length));
      }
      return result;
  }
}

export const toBoolean = (value)=>{
  return value === 'true' ? true : value === 'false' ? false : !!value;
}