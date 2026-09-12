# CacOps 部署（官网 + 智算平台）

本仓库是 **ui4 官网** 和 **aiforge 智算平台** 的组合。不包含 Vben mock 控制台。

用户可见品牌名：**CacOps**。

---

## 你将得到什么

| 服务 | 地址 | 作用 |
|------|------|------|
| 官网落地页 | http://127.0.0.1:4176 | 介绍页，CTA 跳进平台 |
| 智算平台 | http://127.0.0.1:8787 | 登录、数据集、模型、任务、公开资源 |
| Postgres | `127.0.0.1:5433` | 平台主库 `gitea` + 统计库 `statistic` |
| Redis | `127.0.0.1:6379` | 锁 / 队列 |
| MinIO | http://127.0.0.1:9000 | 数据集和模型文件 |

推荐路径：打开官网 → 点「立即开始」→ 进入 `8787/dashboard`。

---

## 环境要求

- Docker Desktop（Postgres / Redis / MinIO）
- Go **1.18+**（编译 `opendata`）
- Node.js **14+**（如需重编前端；仓库已带编译好的 `aiforge/public`）
- 可选：`pnpm` / `npx` 用来起官网

```bash
git clone https://github.com/llcccww12/cacops.git
cd cacops
```

---

## 1. 启动依赖

```bash
cd deploy
docker compose up -d
docker compose ps
```

等 Postgres、Redis、MinIO 都 healthy。MinIO 控制台：http://127.0.0.1:9001 （`minioadmin` / `minioadmin`）。

首次起来后，MinIO 会在平台第一次启动时自动建 bucket `gitea`。

---

## 2. 配置并编译智算平台

```bash
cd ../aiforge
cp custom/conf/app.ini.example custom/conf/app.ini
```

`app.ini.example` 已按本地 Docker 端口写好。如果要对外提供，至少改：

- `[server] DOMAIN` / `ROOT_URL`
- `[security] SECRET_KEY`
- MinIO / 数据库密码

编译后端（在 `aiforge/` 目录）：

```bash
go build -mod=vendor -o opendata .
```

需要重编前端时（通常不必）：

```bash
npm install
npx webpack --config webpack_pro.config.js
```

启动：

```bash
./opendata web
```

浏览器打开 http://127.0.0.1:8787 。首次可注册账号，或用管理员创建用户：

```bash
./opendata admin create-user \
  --username admin \
  --password 'ChangeMe123' \
  --email admin@local \
  --admin \
  --must-change-password=false
```

---

## 3. 启动官网

另开终端，在仓库根目录：

```bash
npx --yes serve ./landing -l 4176
```

打开：

- http://127.0.0.1:4176
- 中文：http://127.0.0.1:4176/zh

顶栏「立即开始」→ `http://127.0.0.1:8787/dashboard`  
其它 CTA → `http://127.0.0.1:8787/user/login`

改跳转：编辑 `landing/assets/js/portal-login.js` 的 `NAV_DASHBOARD` / `DEFAULT_LOGIN`。

---

## 4. 本地验收

- [ ] `deploy` 三个容器在跑
- [ ] `./opendata web` 监听 `8787`
- [ ] 能注册 / 登录智算平台
- [ ] 官网 `4176` 打开，点「立即开始」进 dashboard
- [ ] 公开资源里能看到本地数据集、模型（若已导入）
- [ ] 创建任务时能选到这些本地资源，而不是魔塔外链

存储：数据集 / 模型文件走 MinIO。角色里若没有存储配额，上传会被拒。本地可在库里加一条通用存储角色（`num=-1` 表示不限额）。

---

## 生产环境要点

1. 把 `ROOT_URL`、官网域名、`portal-login.js` 换成公网 HTTPS。
2. 改掉 example 里的数据库 / MinIO / `SECRET_KEY`。
3. Postgres、Redis、MinIO 用独立磁盘和备份，不要用开发默认密码。
4. 反向代理：`4176` 或静态托管官网；`8787` 反代智算平台。
5. 不要把 `custom/conf/app.ini`、`data/`、MinIO 磁盘提交进 Git。

---

## 常见问题

**8787 起不来 / 连不上库**  
看 `docker compose ps`，确认 `5433`、`6379`、`9000` 没被占用。

**上传数据集提示存储限额**  
给用户配存储角色，或按本地开发把 `flow_control.IGNORE_FLAG` 打开。

**官网点了没反应**  
必须先起 `8787`。跳转地址写死在 `landing/assets/js/portal-login.js`。

**这和 zhisuan-platform 仓库有什么区别**  
`zhisuan-platform` 是另一套 Vben 门户 + mock。本仓库才是官网 + 智算平台这一套。
