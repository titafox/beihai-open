# 北海 · NORTH — 数字藏品（NFT）平台

> 北海二期项目源码：移动端 App / H5 + Go API 服务 + FastAdmin 管理后台。
>
> ⚠️ **项目已停止运营，代码以归档/学习为目的开源**。公开前已移除全部密钥、证书和用户数据，
> 并修复了已知的高危漏洞（见文末[安全说明](#安全说明)），但未经生产环境重新验证，**请勿直接用于生产**。

北海是一个基于 **BSN 文昌链（边界智能 AVATA 开放平台）** 的数字藏品发行与交易平台，主要功能：

- 手机号短信登录、实名认证、邀请/排行
- 藏品首发（发售日历、作品详情、下单支付）
- 链上铸造：支付成功 → RabbitMQ 队列 → AVATA 接口创建链账户 / 铸造 / 转移 NFT
- 二级市场：藏品挂售 / 下架 / 购买，流通记录
- 藏品转赠、合成（以材料藏品合成新藏品）、空投、持仓快照
- 汇付天下钱包（开户、余额支付、提现）
- 公告、Banner、IP/艺术家、App 版本管理

## 目录结构

```
beihai/
├── frontend/   # uni-app (Vue 2) 客户端：App + H5，"北海·NORTH"
├── backend/    # Go (Gin) API 服务，模块名 coltstail-server
└── admin/      # FastAdmin (ThinkPHP 5 + Bootstrap) 运营后台
```

### 技术栈

| 模块 | 技术 |
|------|------|
| frontend | uni-app（Vue 2）、uView UI、z-paging、Vuex |
| backend | Go 1.13+、Gin、GORM v1、MySQL、Redis、RabbitMQ、beego task（定时任务）、Swagger、go-ethereum、腾讯云 COS、gRPC |
| admin | PHP ≥ 7.1、ThinkPHP 5、FastAdmin、阿里云 OSS 插件 |
| 链/支付 | BSN AVATA（文昌链）、汇付天下企账通（CFCA 加签，依赖一个 Java/Tomcat 签名服务） |

## 安装部署

### 0. 准备

| 软件 | 版本 | 用途 |
|------|------|------|
| MySQL | 5.7 / 8.0 | 后端与后台共用一个库（默认 `beihai`） |
| Redis | 5+ | 首发库存、防重复提交、验证码限频 |
| RabbitMQ | 3.8+ | 铸造队列 `mint` |
| Go | 1.18+（`go.mod` 声明 1.13） | 编译后端 |
| PHP | 7.1 ~ 7.4，扩展 `pdo_mysql` `curl` `json` `bcmath` `mbstring` `gd` `fileinfo` | 运营后台 |
| Composer | 2.x | 安装后台依赖 |
| Nginx | 任意 | 反向代理 API、托管后台和 H5 |
| HBuilderX | 最新版 | 运行 / 打包 uni-app 客户端 |

还需要以下第三方账号（都要自行申请）：

- **BSN 文昌链 / AVATA**：API Key、API Secret、NFT 类别 ID（class_id）；
- **汇付天下企账通**：商户号、分账账户、CFCA 加签证书（`.pfx`）和验签证书（`.cer`），以及汇付提供的 Java 加签/验签服务 `hfpcfca`（不在本仓库中）；
- **阿里云云市场**：短信和实名认证接口的 APPCODE；
- **腾讯云 COS**（用户头像）和 **阿里云 OSS**（后台上传的作品图片）。

```bash
git clone <仓库地址> beihai && cd beihai
```

### 1. 数据库

```bash
mysql -uroot -p -e "CREATE DATABASE beihai DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
  CREATE USER 'beihai'@'%' IDENTIFIED BY '<密码>';
  GRANT ALL ON beihai.* TO 'beihai'@'%';"

# 后台基础表（fa_ 前缀：管理员、权限、配置、附件等）
mysql -uroot -p beihai < admin/application/admin/command/Install/fastadmin.sql

# 业务表（tbl_ 前缀，22 张）
mysql -uroot -p beihai < database/beihai_schema.sql
```

设置后台管理员密码（FastAdmin 的密码算法是 `md5(md5(密码) + salt)`），并删掉 SQL 里自带的前台演示用户：

```sql
UPDATE fa_admin SET salt = 'Xy7Kq2', password = MD5(CONCAT(MD5('<你的后台密码>'), 'Xy7Kq2')) WHERE username = 'admin';
DELETE FROM fa_user WHERE id = 1;
```

> `database/beihai_schema.sql` 是**根据代码反推出来的**（原项目没有保留建表脚本）：字段名与 Go 模型、后台表单一一对应，
> 但字段类型、长度、默认值和索引是按用途推定的，未与原线上库核对。有原库时，以 `mysqldump --no-data` 导出的为准。

### 2. Redis 和 RabbitMQ

Redis 用默认配置即可，**不要对公网开放**（代码里没有设置 Redis 密码）。

RabbitMQ 需要建虚拟主机、用户，并**手动创建 `mint` 队列**（代码只消费、不声明队列）：

```bash
rabbitmqctl add_vhost bh
rabbitmqctl add_user beihai '<密码>'
rabbitmqctl set_permissions -p bh beihai ".*" ".*" ".*"
rabbitmq-plugins enable rabbitmq_management
rabbitmqadmin -u beihai -p '<密码>' -V bh declare queue name=mint durable=true
```

对应的连接串：`RABBITMQ_URL=amqp://beihai:<密码>@127.0.0.1:5672/bh`。

### 3. 汇付加签服务

后端通过 HTTP 调用汇付提供的 Java 服务做 CFCA 加签和验签（`HF_MAKE_SIGN_URL` / `HF_VERIFY_SIGN_URL`）。
把汇付给的 `hfpcfca` war 包和证书放到 Tomcat 的 `webapps/` 下运行，监听 `127.0.0.1:8080`。
**这个服务只能内网访问**：一旦对公网开放，任何人都能用它伪造支付回调的签名。

### 4. 后端 API

```bash
cd backend
cp .env.example .env
vi .env           # 必填：JWT_SECRET（openssl rand -hex 32）、DB_*、RABBITMQ_URL、AVATA_*、HF_*、COS_*、ALIYUN_MARKET_APPCODE
go mod download
go build -o coltstail-server main.go

set -a && source .env && set +a
./coltstail-server            # 默认监听 :8000
```

启动日志里出现 `Using dev environment` 和 `Listening and serving HTTP on :8000`，就说明启动成功；`curl localhost:8000/api/v1/banner/list` 应返回 JSON。

生产环境：

- 设置 `ENV=prod`：缺少 `JWT_SECRET`（至少 32 位）、`DB_PASSWORD`、`RABBITMQ_URL`、`AVATA_API_KEY/SECRET` 任一项都会拒绝启动；Swagger 会关闭。
- 后端放在 Nginx 之后时，把 Nginx 的地址写进 `TRUSTED_PROXIES`，否则按 IP 限频拿到的都是代理的 IP。
- 在 AVATA 控制台把回调地址配置为 `https://<api 域名>/api/v1/bsn/notify`。
- 汇付的三个回调地址（`HF_BG_RET_URL` 等）必须是公网可以访问的 HTTPS 地址。
- 用 Docker 部署时，可以参考 `backend/build/`（`Dockerfile`、`docker-compose.yml`、`build.sh`）。

Nginx 反向代理示例：

```nginx
server {
    listen 443 ssl;
    server_name api.example.com;
    # ssl_certificate ...;

    location /api/ {
        proxy_pass http://127.0.0.1:8000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

### 5. 运营后台

```bash
cd admin
composer install --no-dev
cp .env.example .env
vi .env           # 填 [database] 段，[token] key 填 32 位随机串（openssl rand -hex 16）
chmod -R 775 runtime public/uploads    # Web 进程需要可写（runtime 目录不存在就先 mkdir）
php think menu -c all-controller       # 新库的 fa_auth_rule 里没有业务模块菜单，用这条命令生成
```

Nginx 示例（ThinkPHP 需要 URL 重写，Web 根指向 `admin/public`）：

```nginx
server {
    listen 443 ssl;
    server_name admin.example.com;
    root /path/to/beihai/admin/public;
    index index.php;

    location / {
        if (!-e $request_filename) { rewrite ^(.*)$ /index.php?s=$1 last; }
    }
    location ~ \.php(/|$) {
        fastcgi_pass 127.0.0.1:9000;
        fastcgi_split_path_info ^(.+\.php)(/.*)$;
        fastcgi_param PATH_INFO $fastcgi_path_info;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include fastcgi_params;
    }
    location ^~ /uploads/ { location ~ \.php$ { deny all; } }   # 上传目录禁止执行脚本
}
```

后台入口是 `https://admin.example.com/sa.php`，用第 1 步设置的 `admin` 账号登录。
登录后在「插件管理 → 阿里云OSS」里填写 OSS 的 AccessKey、Bucket 和域名。
**正式部署前把 `public/sa.php` 改成一个随机文件名**，入口一旦被人知道，后台就会被爆破。

### 6. 客户端（App / H5）

1. 修改 `frontend/utils/constant.js`：`BASEURL` 改成第 4 步的 API 地址（以 `/` 结尾），`INVURL` 改成 H5 的地址；
2. `cd frontend && npm install`；
3. 用 HBuilderX 打开 `frontend/`：
   - 调试：运行 → 运行到浏览器 / 手机或模拟器；
   - H5：发行 → 网站-H5，把 `unpackage/dist/build/h5` 部署到 Nginx；
   - App：发行 → 原生 App 云打包，Android 证书需要自己生成，`manifest.json` 里的 appid 换成自己 DCloud 账号下的应用。

### 7. 上架首发作品

1. 后台「作品管理」新建作品，填写发行总量、预留份数、价格、发售时间；
2. **在 Redis 里写入该作品的可售库存**（后台不会自动写入，不写的话下单会报"脚本执行失败"）：

   ```bash
   # 可售份数 = 发行总量 - 预留份数；库存脚本判断的是"剩余 - 购买 > 0"，所以要多加 1 才能卖完
   redis-cli SET work_<作品ID> <可售份数 + 1>
   ```

3. 用户支付成功后，后端经 `mint` 队列调用 AVATA 铸造，链上结果通过 `/api/v1/bsn/notify` 回调更新。

### 常见问题

| 现象 | 原因 |
|------|------|
| 启动时报 `Failed to connect to OrderConn RabbitMQ` | `RABBITMQ_URL` 配错，或者 vhost `bh` 不存在 |
| 下单时报「脚本执行失败」 | Redis 里没有 `work_<作品ID>` 这个库存 key，见第 7 步 |
| 收不到短信验证码 | 没有配置 `ALIYUN_MARKET_APPCODE`，或者阿里云市场的短信模板 ID 不一样（`utils/sms.go` 中的 `TPL_09828`） |
| 支付页面拉不起来 | 汇付加签服务没起来，或者商户号、证书与环境（测试 / 生产）不匹配 |
| 后台页面一片空白 | `runtime/` 目录不可写，或者 PHP 缺少扩展 |

## frontend — uni-app 客户端

```
frontend/
├── pages/          # 39 个页面（见下）
├── components/     # 公共组件
├── api/ common/ utils/ store/
├── uni_modules/    # uview-ui、z-paging
├── manifest.json   # 应用名 北海·NORTH，appid __UNI__5E42D80
└── pages.json
```

主要页面：首页 `index`、藏品 `collection` / `allNFTs` / `nftDetail`、市场 `market`、流通 `circulation`、
合成 `CollectionDetails/merge`、转赠 `handselList`、订单 `order`、支付 `pay` / `payResult`、提现 `withdrawal`、
实名 `realNameAuth`、银行卡 `cardManage`、链地址 `BSNAddress`、邀请 `invitation`、公告 `notice`、艺术家 `artist`、个人中心 `mine` / `myProfile` 等。

**接口地址** 在 `frontend/utils/constant.js`：

```js
export const BASEURL = 'https://bhapi.mareboreum.com/';  // API
export const INVURL  = 'https://h5.mareboreum.com/';     // 邀请 H5
```

运行和打包见 [安装部署 · 客户端](#6-客户端app--h5)。

## backend — Go API 服务

```
backend/
├── main.go          # 启动：定时任务 → Redis → RabbitMQ 铸造消费者 → Gin
├── conf/config.go   # 配置全部从环境变量读取（示例见 .env.example）
├── routes/          # 路由，统一前缀 /api/v1
├── controllers/     # 账户、登录、作品、订单、支付、市场、转赠、合成、BSN 回调…
├── models/          # GORM 模型 + Redis + RabbitMQ
├── services/        # 定时任务 cron.go、幸运狗合约、ML gRPC 服务
├── middlewares/     # CORS、i18n、JWT
├── utils/           # AVATA 签名请求、汇付支付、短信、雪花 ID、IPFS…
├── entity/req|res   # 请求/响应结构
├── docs/            # Swagger（swag 生成）
└── build/           # Dockerfile、docker-compose.yml、部署脚本
```

安装和运行见 [安装部署 · 后端 API](#4-后端-api)。
Swagger 文档（仅非 prod 环境）：`http://localhost:8000/swagger/index.html`（重新生成：`swag init`）。

### 定时任务（`services/cron.go`）

| 任务 | 周期 | 作用 |
|------|------|------|
| tkSnapshot | 每 3 分钟 | 持仓快照 |
| tkAirdrops | 每 5 分钟 | 空投发放 |
| tkExpiredOrder | 每 1 分钟 | 关闭超时未支付订单 |
| tkUsers | 每 2 分钟 | 用户数据统计 |

### 主要接口（`/api/v1`）

| 分组 | 接口 |
|------|------|
| 账户 | `POST account/smsLogin/code` `POST account/smsLogin` `POST account/authentication` `GET account/info` `POST account/update` `GET account/invite/list` `GET account/invite/top` |
| 作品/首发 | `GET product/list` `GET product/detail/:id` `GET product/buyList` `GET line` |
| 订单/支付 | `POST order/commit` `GET order/list/:typ` `POST pay/notify` `POST balance-pay/notify` |
| 钱包 | `POST wallet/create` `POST wallet/notify` `GET wallet/info` |
| 市场 | `GET market` `GET market/detail` `POST market/nft/on` `POST market/nft/off` `POST market/nft/buy` `GET market/nfts/circulating` |
| 我的藏品 | `GET account/nfts` `GET account/nfts/series` `GET account/nfts/sells` |
| 转赠/合成 | `POST works/handsel` `GET handsel/list` `GET nft/composite/list` `GET nft/composite/material` `POST nft/composite/start` |
| 链回调 | `POST bsn/notify` |
| 其他 | `GET app/version` `GET banner/list` `GET notice*` |

需登录的接口在 Header 带 `Authorization: <JWT>`。

### 部署

```bash
DEPLOY_HOST=root@<测试服务器> ./build.sh        # 交叉编译 linux/amd64 → scp → 远程执行 build/build.sh
DEPLOY_HOST=root@<生产服务器> ./build-prod.sh   # 生产环境
```

`build/Dockerfile` 把二进制打成 `coltstail-server` 镜像（`ENV=prod`，暴露 8000），运行时通过 `--env-file` / `env_file` 注入配置。
`build/docker-compose.yml` 是整套服务器编排：nginx + certbot、PHP 7.4（后台）、MySQL、Redis、coltstail-server、java-tomcat（汇付 CFCA 签名服务）、portainer。
MySQL、Redis、签名服务、portainer 只绑定 127.0.0.1，不要改成对公网开放。

## admin — FastAdmin 运营后台

```
admin/
├── application/admin/controller/tbl/   # 业务模块
├── application/database.php            # 数据库配置（表前缀 fa_，通过 admin/.env 设置）
├── addons/alioss/                      # 阿里云 OSS 上传插件
├── public/                             # Web 根目录
└── thinkphp/                           # ThinkPHP 5 框架
```

业务模块（`tbl/`）：用户 Users / Account、作品 Works、藏品 Nfts、产品 Product、IP 方 Ips、Banner、公告 Notice（含分类）、
转赠 Handsel、合成 Composite、空投 Airdrops（含日志）、持仓快照 Snapshot / Snapshot2。

安装见 [安装部署 · 运营后台](#5-运营后台)。部署时建议：
- 把后台入口 `public/sa.php` 改成随机文件名，并加 IP 白名单或 HTTP Basic Auth；
- 保持 `app.debug = false`；登录验证码已默认开启；
- 在线命令插件（`addons/command`）已默认禁用，生产环境不要启用；
- 升级到 FastAdmin 最新版（当前为 1.3.4.20220530 / ThinkPHP 5.0.24）。

## 已知限制

- **业务表结构是反推的**：`database/beihai_schema.sql` 根据代码生成，未与原线上库核对，见 [安装部署 · 数据库](#1-数据库)。
- 依赖版本较老（Go 1.13、GORM v1、`dgrijalva/jwt-go`、Vue 2、ThinkPHP 5），使用前建议升级。
- 前端「修改手机号」页面（`pages/contactInfo`）调用的 `/account/updatePhone` 已移除：原接口直接映射到登录，并不能修改手机号。
- 一级市场 Redis 库存脚本（`models/redis.go` 的 `ReduceStock`）用的是 `剩余 - 购买 > 0`，最后一份无法售出，因此初始库存要多写 1（见[上架首发作品](#7-上架首发作品)）。

## 安全说明

开源前做过一轮安全清理：

- 删除了所有硬编码的密钥、证书、Android 签名文件、服务器 IP 和用户个人信息，配置改为环境变量；
- 修复短信验证码可暴力破解 / 复用（改为 6 位、5 分钟过期、一次性、输错 5 次锁定、按 IP 限频）；
- 修复市场搜索接口的 SQL 注入；
- 修复合成接口可销毁他人藏品、配方校验不全；
- BSN 回调改为按 operation_id 向 AVATA 回查结果，不再信任回调内容；
- 二级市场改为原子锁单，防止一件藏品被多人买走；支付回调、铸造消费者增加幂等；
- 首发下单失败自动回补库存，超时订单回补库存方向修正，恢复下单实名校验；
- JWT 只接受 HS256 并校验 issuer；生产环境强制要求随机密钥；
- 对外接口返回的手机号、身份证号打码；Swagger 仅在非生产环境开放；
- 后台默认关闭调试模式、开启登录验证码、禁用在线命令插件、禁止上传 zip/rar。

尚未处理（如需重新上线必须先做）：

- 支付回调验签依赖外部 Java 签名服务，建议改为进程内用汇付公钥验签，并回查订单金额；
- 实名认证接口没有调用频率限制（按次计费）；
- 身份证号明文存储，应加密存储。

发现安全问题请通过 Issue 联系（不要在 Issue 中公开漏洞细节）。

## License

[Apache License 2.0](LICENSE)。`admin/` 基于 [FastAdmin](https://www.fastadmin.net)（Apache-2.0），前端使用 uView UI、z-paging（MIT），各自保留原许可。
