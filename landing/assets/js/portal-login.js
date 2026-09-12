/**
 * CacOps 官网 CTA 跳转。
 *
 * 顶栏「立即开始 / 开始使用 / Get Started」→ 智算平台 dashboard。
 * 其它同类 CTA（启动项目、预约咨询等）→ 智算平台登录页。
 * Framer 会拦 <a> 的默认跳转，所以不能只改 href，必须在捕获阶段自己 location.assign。
 */
(function () {
  if (window.__CACOPS_PORTAL_LOGIN_BOOTED__) return;
  window.__CACOPS_PORTAL_LOGIN_BOOTED__ = true;
  const DEFAULT_LOGIN = "http://127.0.0.1:8787/user/login";
  const NAV_DASHBOARD = "http://127.0.0.1:8787/dashboard";
  const NAV_START_TEXT = /立即开始|开始使用|Get Started/i;
  const CTA_TEXT =
    /立即开始|开始使用|Get Started|Book a call|预约通话|预约咨询|预约需求沟通|预约免费评估|启动项目|Sign in|免费评估/i;
  const CTA_HREF =
    /(\/pages\/(?:zh\/)?contact(?:\.html)?)|(^\.\/contact)|(^https?:\/\/(?:www\.)?cal\.com)|(\/auth\/login)/i;

  function portalLoginUrl() {
    if (typeof window.__CACOPS_PORTAL_LOGIN__ === "string" && window.__CACOPS_PORTAL_LOGIN__) {
      return withFrom(window.__CACOPS_PORTAL_LOGIN__);
    }
    const meta = document.querySelector('meta[name="cacops-portal-login"]');
    const fromMeta = meta && meta.getAttribute("content");
    if (fromMeta) return withFrom(fromMeta);
    return DEFAULT_LOGIN;
  }

  function withFrom(url) {
    try {
      const u = new URL(url, window.location.href);
      if (!u.searchParams.has("from")) u.searchParams.set("from", "landing");
      return u.toString();
    } catch {
      return url;
    }
  }

  function linkText(el) {
    return (el.textContent || "").replace(/\s+/g, " ").trim();
  }

  function isInTopBar(el) {
    if (!el.closest) return false;
    if (el.closest("nav, header")) return true;
    if (el.closest('[data-framer-name="Logo and menu"]')) return true;
    let node = el;
    for (let i = 0; i < 8 && node; i += 1) {
      const style = window.getComputedStyle(node);
      if (
        (style.position === "fixed" || style.position === "sticky") &&
        node.getBoundingClientRect().top < 80
      ) {
        return true;
      }
      node = node.parentElement;
    }
    return false;
  }

  function isNavStartCta(el) {
    return Boolean(el && NAV_START_TEXT.test(linkText(el)) && isInTopBar(el));
  }

  function destUrl(el) {
    if (el && el.getAttribute("data-cacops-portal") === "dashboard") return NAV_DASHBOARD;
    return isNavStartCta(el) ? NAV_DASHBOARD : portalLoginUrl();
  }

  function isPortalCta(el) {
    if (!el || el.nodeType !== 1) return false;
    if (el.closest && el.closest("#lang-switch-root")) return false;
    const tag = el.tagName;
    if (tag !== "A" && tag !== "BUTTON") return false;
    if (el.getAttribute("data-cacops-portal")) return true;
    const href = el.getAttribute("href") || "";
    const text = linkText(el);
    return CTA_TEXT.test(text) || CTA_HREF.test(href);
  }

  function goDest(el, event) {
    if (event) {
      event.preventDefault();
      event.stopPropagation();
      if (event.stopImmediatePropagation) event.stopImmediatePropagation();
    }
    window.location.assign(destUrl(el));
  }

  function rewriteAnchors() {
    document.querySelectorAll("a[href], button").forEach(function (el) {
      if (!isPortalCta(el) && !isNavStartCta(el)) return;
      const url = destUrl(el);
      el.setAttribute("data-cacops-portal", isNavStartCta(el) ? "dashboard" : "1");
      if (el.tagName === "A") {
        el.setAttribute("href", url);
        el.setAttribute("target", "_self");
        el.setAttribute("rel", "noopener");
      }
    });
  }

  function onPointer(event) {
    const hit = event.target && event.target.closest ? event.target.closest("a, button") : null;
    if (!isPortalCta(hit) && !isNavStartCta(hit)) return;
    goDest(hit, event);
  }

  function boot() {
    const extra = document.getElementById("cacops-portal-login");
    if (extra) extra.remove();
    rewriteAnchors();
  }

  document.addEventListener("click", onPointer, true);
  document.addEventListener("pointerup", onPointer, true);

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", boot);
  } else {
    boot();
  }

  const obs = new MutationObserver(function () {
    boot();
  });
  obs.observe(document.documentElement, { childList: true, subtree: true });
})();
