const { AppVer } = window.config;

// compat: IE11
if (!Element.prototype.matches) {
  Element.prototype.matches = Element.prototype.msMatchesSelector || Element.prototype.webkitMatchesSelector;
}

// compat: IE11
if (!Element.prototype.closest) {
  Element.prototype.closest = function (s) {
    let el = this;

    do {
      if (el.matches(s)) return el;
      el = el.parentElement || el.parentNode;
    } while (el !== null && el.nodeType === 1);
    return null;
  };
}

// rewrite to solve dynamic import cache error
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
