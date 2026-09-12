# 官网 ↔ 智算平台

落地页不再走 mock 控制台。

- 顶栏「立即开始 / 开始使用 / Get Started」→ `http://127.0.0.1:8787/dashboard`
- 其它同类 CTA → `http://127.0.0.1:8787/user/login`

改线上地址：编辑 `assets/js/portal-login.js` 里的 `NAV_DASHBOARD` / `DEFAULT_LOGIN`，或运行时：

```js
window.__CACOPS_PORTAL_LOGIN__ = "https://your-host/user/login";
```
