(function () {
  if (!window.location.pathname.includes("/pages/zh/")) return;

  var dict = null;
  var pending = false;
  var lastRun = 0;

  function loadDict() {
    return fetch("/locales/zh-runtime.json")
      .then(function (res) {
        return res.json();
      })
      .then(function (data) {
        dict = data || {};
        return dict;
      })
      .catch(function () {
        dict = {};
        return dict;
      });
  }

  function translateNode(el) {
    if (!el || el.closest("#lang-switch-root")) return;

    var spans = el.querySelectorAll(":scope > span");
    if (spans.length > 1) {
      var joined = "";
      for (var i = 0; i < spans.length; i++) {
        joined += spans[i].textContent || "";
      }
      joined = joined.replace(/\s+/g, " ").trim();
      var zhJoined = dict[joined];
      if (zhJoined) {
        el.textContent = zhJoined;
        return;
      }
    }

    if (el.childElementCount === 0) {
      var text = (el.textContent || "").trim();
      if (!text) return;
      var zh = dict[text];
      if (zh) el.textContent = zh;
    }
  }

  function apply() {
    if (!dict) return;
    var now = Date.now();
    if (now - lastRun < 120) return;
    lastRun = now;

    var nodes = document.querySelectorAll(
      "p.framer-text, h1.framer-text, h2.framer-text, h3.framer-text, h4.framer-text, li.framer-text, button.framer-text, a.framer-text, span.framer-text"
    );
    for (var i = 0; i < nodes.length; i++) {
      translateNode(nodes[i]);
    }

    // Common nav / CTA labels that may not use framer-text class alone
    var loose = document.querySelectorAll("a, button, p, h1, h2, h3, li");
    for (var j = 0; j < loose.length; j++) {
      var el = loose[j];
      if (el.closest("#lang-switch-root")) continue;
      if (el.children.length > 0 && !el.querySelector(":scope > span")) continue;
      var t = (el.textContent || "").trim();
      if (!t || t.length > 180) continue;
      if (dict[t]) el.textContent = dict[t];
    }
  }

  function schedule() {
    if (pending) return;
    pending = true;
    window.requestAnimationFrame(function () {
      pending = false;
      apply();
    });
  }

  loadDict().then(function () {
    schedule();
    [200, 600, 1200, 2500, 5000].forEach(function (ms) {
      window.setTimeout(schedule, ms);
    });

    if (document.body) {
      new MutationObserver(function () {
        schedule();
      }).observe(document.body, {
        childList: true,
        subtree: true,
        characterData: true,
      });
    }
  });
})();
