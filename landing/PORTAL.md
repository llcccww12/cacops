# CacOps 官网落地页 ↔ 控制台

落地页**顶栏**「立即开始 / 开始使用 / Get Started」跳转到：

```text
http://127.0.0.1:8787/dashboard
```

其它同类 CTA（启动项目、预约咨询等）仍进 Vben 登录页：

```text
http://127.0.0.1:4180/auth/login?from=landing
```

修改跳转地址：编辑各页 `<meta name="cacops-portal-login" content="...">`，或运行时设置：

```js
window.__CACOPS_PORTAL_LOGIN__ = "https://console.example.com/auth/login";
```

重新注入脚本：

```bash
node scripts/inject-portal-login.mjs
```

本地预览：

```bash
npx serve . -l 4176
# 或在 ui 仓库：pnpm dev:landing
```
