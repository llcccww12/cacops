(function () {
  const ZH_PREFIX = "/pages/zh/";
  const EN_PREFIX = "/pages/";

  function isZhPath(pathname) {
    return pathname.includes(ZH_PREFIX);
  }

  function normalizePath(pathname) {
    if (!pathname || pathname === "/") return "/pages/index.html";
    if (pathname === "/pages/" || pathname === "/pages") return "/pages/index.html";
    return pathname;
  }

  function toEnPath(pathname) {
    pathname = normalizePath(pathname);
    if (!isZhPath(pathname)) return pathname;
    return pathname.replace(ZH_PREFIX, EN_PREFIX);
  }

  function toZhPath(pathname) {
    pathname = normalizePath(pathname);
    if (isZhPath(pathname)) return pathname;
    if (pathname.startsWith(EN_PREFIX)) {
      return pathname.replace(EN_PREFIX, ZH_PREFIX);
    }
    return "/pages/zh/index.html";
  }

  function rememberLang(lang) {
    try {
      localStorage.setItem("preferred-lang", lang);
    } catch (_) {}
  }

  function mount() {
    if (document.getElementById("lang-switch-root")) return;

    const pathname = normalizePath(window.location.pathname);
    const zh = isZhPath(pathname);
    const enHref = toEnPath(pathname);
    const zhHref = toZhPath(pathname);

    const root = document.createElement("div");
    root.id = "lang-switch-root";
    root.className = "lang-switch lang-switch--fixed";
    root.setAttribute("role", "navigation");
    root.setAttribute("aria-label", zh ? "语言切换" : "Language switcher");

    // Use real links with target=_top so Framer SPA cannot swallow navigation.
    root.innerHTML =
      '<a class="lang-switch__btn' +
      (zh ? "" : " is-active") +
      '" href="' +
      enHref +
      '" target="_top" rel="noopener" lang="en" hreflang="en">EN</a>' +
      '<span class="lang-switch__sep" aria-hidden="true">|</span>' +
      '<a class="lang-switch__btn' +
      (zh ? " is-active" : "") +
      '" href="' +
      zhHref +
      '" target="_top" rel="noopener" lang="zh-CN" hreflang="zh-CN">中文</a>';

    document.body.appendChild(root);

    root.querySelectorAll("a.lang-switch__btn").forEach(function (link) {
      link.addEventListener(
        "click",
        function () {
          rememberLang(link.getAttribute("lang") || "en");
          // Force hard navigation even if another handler tries to stop it.
          window.setTimeout(function () {
            if (window.location.pathname === pathname) {
              window.top.location.href = link.href;
            }
          }, 0);
        },
        true
      );
    });
  }

  function ensureMounted() {
    mount();
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", ensureMounted);
  } else {
    ensureMounted();
  }

  window.addEventListener("load", ensureMounted);
  [300, 800, 1500, 3000].forEach(function (delay) {
    window.setTimeout(ensureMounted, delay);
  });

  if (document.body) {
    new MutationObserver(function () {
      if (!document.getElementById("lang-switch-root")) {
        ensureMounted();
      }
    }).observe(document.body, { childList: true, subtree: false });
  }
})();
