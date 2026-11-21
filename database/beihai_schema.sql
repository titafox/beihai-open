-- =====================================================================
-- 北海·NORTH 业务表结构（tbl_*）
--
-- 原项目没有保留建表脚本，本文件根据以下代码反推生成：
--   * backend/models/*.go 中的 Tbl* 结构体（GORM v1，SingularTable，字段名 snake_case）
--   * backend 中的原生 SQL
--   * admin 后台的模型、表单（view/tbl/*）、列表（assets/js/backend/tbl/*）和语言包
-- 字段类型、长度、默认值和索引是按用途推定的，不一定与原线上库完全一致。
--
-- 用法（先导入 FastAdmin 基础表 admin/application/admin/command/Install/fastadmin.sql）：
--   mysql -u<user> -p beihai < database/beihai_schema.sql
--
-- 约定：时间字段统一为 Unix 秒级时间戳；金额为 decimal(10,2)。
-- =====================================================================

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ---------------------------------------------------------------------
-- 用户 / 账户
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `tbl_account` (
  `id`                         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uid`                        varchar(32)  NOT NULL DEFAULT '' COMMENT '用户唯一标识（邀请码）',
  `pid`                        bigint(20)   NOT NULL DEFAULT 0  COMMENT '邀请人 account.id',
  `role`                       tinyint(4)   NOT NULL DEFAULT 0  COMMENT '角色：0 普通用户',
  `status`                     tinyint(4)   NOT NULL DEFAULT 1  COMMENT '状态：0 禁用 1 正常',
  `phone`                      varchar(20)  NOT NULL DEFAULT '' COMMENT '手机号',
  `nickname`                   varchar(64)  NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar`                     varchar(255) NOT NULL DEFAULT '' COMMENT '头像 URL',
  `bsn_address`                varchar(128) NOT NULL DEFAULT '' COMMENT 'BSN 链账户地址',
  `bsn_username`               varchar(64)  NOT NULL DEFAULT '' COMMENT 'BSN 链账户名称',
  `name`                       varchar(32)  NOT NULL DEFAULT '' COMMENT '真实姓名',
  `amount`                     int(11)      NOT NULL DEFAULT 0  COMMENT '余额',
  `id_card_num`                varchar(32)  NOT NULL DEFAULT '' COMMENT '身份证号（建议加密存储）',
  `id_card_positive_image_url` varchar(255) NOT NULL DEFAULT '' COMMENT '身份证正面照',
  `id_card_negative_image_url` varchar(255) NOT NULL DEFAULT '' COMMENT '身份证反面照',
  `authentication`             tinyint(4)   NOT NULL DEFAULT 0  COMMENT '实名认证：0 未认证 1 已认证',
  `description`                varchar(255) NOT NULL DEFAULT '' COMMENT '个人简介',
  `is_handsel`                 tinyint(4)   NOT NULL DEFAULT 1  COMMENT '转赠功能：0 禁用 1 正常',
  `is_wallet`                  tinyint(4)   NOT NULL DEFAULT 0  COMMENT '是否已开通汇付钱包（运行时计算，可不落库）',
  `nums`                       int(11)      NOT NULL DEFAULT 0  COMMENT '真实邀请人数',
  `create_time`                bigint(20)   NOT NULL DEFAULT 0  COMMENT '创建时间',
  `update_time`                bigint(20)   NOT NULL DEFAULT 0  COMMENT '修改时间',
  `del_time`                   bigint(20)   NOT NULL DEFAULT 0  COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_phone` (`phone`),
  KEY `idx_uid` (`uid`),
  KEY `idx_pid` (`pid`),
  KEY `idx_id_card_num` (`id_card_num`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户账户';

CREATE TABLE IF NOT EXISTS `tbl_sms` (
  `id`        int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mobile`    varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `code`      varchar(10) NOT NULL DEFAULT '' COMMENT '验证码',
  `event`     tinyint(4)  NOT NULL DEFAULT 1  COMMENT '事件：1 登录',
  `send_time` bigint(20)  NOT NULL DEFAULT 0  COMMENT '发送时间',
  PRIMARY KEY (`id`),
  KEY `idx_mobile_event` (`mobile`, `event`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='短信验证码';

CREATE TABLE IF NOT EXISTS `tbl_users` (
  `id`    bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `users` bigint(20) NOT NULL DEFAULT 0 COMMENT '在线人数（定时任务统计）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='在线人数统计';

-- ---------------------------------------------------------------------
-- 作品 / 藏品
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `tbl_ips` (
  `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `nickname`     varchar(64)  NOT NULL DEFAULT '' COMMENT '作者 / IP 名称',
  `avatar`       varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `introduction` text                  COMMENT '简介',
  `status`       tinyint(4)   NOT NULL DEFAULT 1 COMMENT '状态：0 禁用 1 正常',
  `create_time`  bigint(20)   NOT NULL DEFAULT 0,
  `update_time`  bigint(20)   NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='IP / 艺术家';

CREATE TABLE IF NOT EXISTS `tbl_works` (
  `id`                    bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `collection`            varchar(64)   NOT NULL DEFAULT '' COMMENT '合集标识',
  `no`                    int(11)       NOT NULL DEFAULT 0  COMMENT '编号',
  `author_id`             bigint(20)    NOT NULL DEFAULT 0  COMMENT '创作者 tbl_ips.id',
  `owner_id`              bigint(20)    NOT NULL DEFAULT 0  COMMENT '拥有者',
  `name`                  varchar(128)  NOT NULL DEFAULT '' COMMENT '藏品名称',
  `description`           text                              COMMENT '藏品介绍',
  `image`                 varchar(255)  NOT NULL DEFAULT '' COMMENT '藏品图片（相对 OSS 路径）',
  `token_id`              varchar(128)  NOT NULL DEFAULT '' COMMENT '链上标识',
  `hash`                  varchar(128)  NOT NULL DEFAULT '' COMMENT '哈希',
  `price`                 decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '价格',
  `stock`                 int(11)       NOT NULL DEFAULT 0  COMMENT '发行总量',
  `reserved`              int(11)       NOT NULL DEFAULT 0  COMMENT '预留份数',
  `stock_sell`            int(11)       NOT NULL DEFAULT 0  COMMENT '已售份数',
  `pub_time`              bigint(20)    NOT NULL DEFAULT 0  COMMENT '公售时间',
  `advance_hour`          bigint(20)    NOT NULL DEFAULT 0  COMMENT '优先购提前时间',
  `sell_time`             bigint(20)    NOT NULL DEFAULT 0  COMMENT '可寄售时间',
  `is_airdrop`            tinyint(4)    NOT NULL DEFAULT 0  COMMENT '是否空投作品',
  `is_synthetic_material` tinyint(4)    NOT NULL DEFAULT 0  COMMENT '是否可作为合成材料：0 否 1 是',
  `typ`                   tinyint(4)    NOT NULL DEFAULT 1  COMMENT '藏品类型：1 普通 2 合成成功 3 合成失败',
  `sold_out`              tinyint(1)    NOT NULL DEFAULT 0  COMMENT '是否售罄：0 未售罄 1 已售罄',
  `status`                tinyint(4)    NOT NULL DEFAULT 0  COMMENT '状态：0 下架 1 上架',
  `create_time`           bigint(20)    NOT NULL DEFAULT 0,
  `update_time`           bigint(20)    NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_author` (`author_id`),
  KEY `idx_status_pub` (`status`, `pub_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='作品（藏品系列）';

CREATE TABLE IF NOT EXISTS `tbl_nfts` (
  `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `wid`          bigint(20)    NOT NULL DEFAULT 0  COMMENT '作品 tbl_works.id',
  `work_no`      bigint(20)    NOT NULL DEFAULT 0  COMMENT '作品内编号',
  `author_id`    bigint(20)    NOT NULL DEFAULT 0  COMMENT '创作者',
  `owner_id`     bigint(20)    NOT NULL DEFAULT 0  COMMENT '拥有者 tbl_account.id',
  `transfer_num` int(11)       NOT NULL DEFAULT 1  COMMENT '流转次数：1 未售出（待分配）',
  `price`        decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '价格 / 寄售价',
  `nft_id`       varchar(128)  NOT NULL DEFAULT '' COMMENT '链上 NFT ID',
  `operation_id` varchar(64)   NOT NULL DEFAULT '' COMMENT 'AVATA 操作 ID',
  `chain_status` tinyint(4)    NOT NULL DEFAULT 0  COMMENT '上链状态：0 上链中 1 已上链 2 上链失败（代码中铸造中也写 2）',
  `status`       tinyint(4)    NOT NULL DEFAULT 0  COMMENT '藏品状态：0 未寄售 1 寄售中 2 锁单中 3 已销毁',
  `mark`         varchar(255)  NOT NULL DEFAULT '' COMMENT '备注',
  `update_time`  bigint(20)    NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_wid_transfer` (`wid`, `transfer_num`),
  KEY `idx_owner_status` (`owner_id`, `status`),
  KEY `idx_status` (`status`),
  KEY `idx_operation` (`operation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='藏品（单个编号）';

CREATE TABLE IF NOT EXISTS `tbl_tx` (
  `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `from_id`      bigint(20)   NOT NULL DEFAULT 0,
  `to_id`        bigint(20)   NOT NULL DEFAULT 0,
  `work_id`      bigint(20)   NOT NULL DEFAULT 0  COMMENT '作品 id',
  `work_no`      bigint(20)   NOT NULL DEFAULT 0  COMMENT '藏品编号',
  `block_height` int(11)      NOT NULL DEFAULT 0  COMMENT '区块高度',
  `tx_hash`      varchar(128) NOT NULL DEFAULT '' COMMENT '交易哈希',
  `typ`          varchar(32)  NOT NULL DEFAULT '' COMMENT '交易类型：mint_nft / transfer_nft …',
  `timestamp`    bigint(20)   NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_tx_hash` (`tx_hash`),
  KEY `idx_work` (`work_id`, `work_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='链上交易记录（BSN 回调）';

-- ---------------------------------------------------------------------
-- 订单 / 市场
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `tbl_orders` (
  `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `no`          varchar(64)   NOT NULL DEFAULT '' COMMENT '订单号（BH 首发 / BY 二级市场）',
  `account_id`  bigint(20)    NOT NULL DEFAULT 0  COMMENT '买家',
  `wid`         bigint(20)    NOT NULL DEFAULT 0  COMMENT '作品 id',
  `nft_id`      bigint(20)    NOT NULL DEFAULT 0  COMMENT '藏品 tbl_nfts.id（二级市场）',
  `price`       decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '单价',
  `num`         int(11)       NOT NULL DEFAULT 1  COMMENT '数量',
  `order_money` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '订单金额',
  `pay_url`     varchar(1024) NOT NULL DEFAULT '' COMMENT '汇付支付链接',
  `typ`         tinyint(4)    NOT NULL DEFAULT 1  COMMENT '1 一级市场 2 二级市场',
  `status`      tinyint(4)    NOT NULL DEFAULT 1  COMMENT '1 待支付 2 已支付 3 已取消',
  `create_time` bigint(20)    NOT NULL DEFAULT 0,
  `pay_time`    bigint(20)    NOT NULL DEFAULT 0,
  `update_time` bigint(20)    NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_no` (`no`),
  KEY `idx_account_status` (`account_id`, `status`),
  KEY `idx_wid_status` (`wid`, `status`),
  KEY `idx_status_create` (`status`, `create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单';

CREATE TABLE IF NOT EXISTS `tbl_market_log` (
  `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `nft_id`      bigint(20)    NOT NULL DEFAULT 0 COMMENT '藏品 tbl_nfts.id',
  `account_id`  bigint(20)    NOT NULL DEFAULT 0,
  `typ`         tinyint(4)    NOT NULL DEFAULT 0 COMMENT '1 寄售 2 买入 …',
  `price`       decimal(10,2) NOT NULL DEFAULT 0.00,
  `status`      tinyint(4)    NOT NULL DEFAULT 1,
  `create_time` bigint(20)    NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_nft` (`nft_id`),
  KEY `idx_account` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='二级市场流通记录';

CREATE TABLE IF NOT EXISTS `tbl_handsel` (
  `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_id`   bigint(20)  NOT NULL DEFAULT 0  COMMENT '接收者 / 操作者',
  `w_id`         bigint(20)  NOT NULL DEFAULT 0  COMMENT '作品 id',
  `work_no`      bigint(20)  NOT NULL DEFAULT 0  COMMENT '藏品编号',
  `typ`          tinyint(4)  NOT NULL DEFAULT 0  COMMENT '方向：0 转入 1 转出',
  `status`       tinyint(4)  NOT NULL DEFAULT 0  COMMENT '0 处理中 1 成功 2 失败',
  `operation_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'AVATA 操作 ID',
  `crete_time`   bigint(20)  NOT NULL DEFAULT 0  COMMENT '操作时间（字段名原样保留拼写）',
  PRIMARY KEY (`id`),
  KEY `idx_account` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='转赠记录';

-- ---------------------------------------------------------------------
-- 合成
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `tbl_composite` (
  `id`              bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title`           varchar(128) NOT NULL DEFAULT '' COMMENT '标题',
  `description`     text                             COMMENT '合成描述',
  `composite_id`    bigint(20)   NOT NULL DEFAULT 0  COMMENT '合成产出的作品 tbl_works.id',
  `stock`           int(11)      NOT NULL DEFAULT 0  COMMENT '可合成库存',
  `remaining_stock` int(11)      NOT NULL DEFAULT 0  COMMENT '剩余可合成数量',
  `material`        text                             COMMENT '合成材料（后台编辑用）',
  `material_info`   text                             COMMENT '合成材料 JSON：[{"wid":1,"name":"","nums":2}]',
  `failure_rate`    int(10) unsigned NOT NULL DEFAULT 0 COMMENT '失败率（%）',
  `status`          tinyint(4)   NOT NULL DEFAULT 0  COMMENT '0 未开始 1 进行中 2 已结束',
  `create_time`     bigint(20)   NOT NULL DEFAULT 0,
  `update_time`     bigint(20)   NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合成活动';

CREATE TABLE IF NOT EXISTS `tbl_composite_log` (
  `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `op`           varchar(64)  NOT NULL DEFAULT '' COMMENT '合成操作 ID（MQ 消息 HC:<op>）',
  `composite_id` bigint(20)   NOT NULL DEFAULT 0  COMMENT '合成活动 id',
  `account_id`   bigint(20)   NOT NULL DEFAULT 0,
  `status`       tinyint(4)   NOT NULL DEFAULT 0  COMMENT '0 待处理 1 成功 2 失败',
  `fail_text`    varchar(255) NOT NULL DEFAULT '' COMMENT '失败原因',
  `create_time`  bigint(20)   NOT NULL DEFAULT 0,
  `update_time`  bigint(20)   NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_op` (`op`),
  KEY `idx_account` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合成记录';

-- ---------------------------------------------------------------------
-- 空投 / 快照（优先购）
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `tbl_airdrops` (
  `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `work_id`     bigint(20) NOT NULL DEFAULT 0 COMMENT '空投作品',
  `whitelist`   text                          COMMENT '空投名单：手机号-数量,手机号-数量',
  `start_time`  bigint(20) NOT NULL DEFAULT 0 COMMENT '空投开始时间',
  `status`      tinyint(4) NOT NULL DEFAULT 0 COMMENT '0 未开始 1 进行中 2 已完成',
  `update_time` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_status_start` (`status`, `start_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='空投活动';

CREATE TABLE IF NOT EXISTS `tbl_airdrops_log` (
  `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_id`   bigint(20)   NOT NULL DEFAULT 0  COMMENT '空投用户',
  `airdrops_id`  bigint(20)   NOT NULL DEFAULT 0  COMMENT '空投活动 id',
  `p_id`         bigint(20)   NOT NULL DEFAULT 0  COMMENT '空投作品',
  `n_id`         int(11)      NOT NULL DEFAULT 0  COMMENT '空投藏品编号',
  `message`      varchar(255) NOT NULL DEFAULT '' COMMENT '处理信息',
  `operation_id` varchar(64)  NOT NULL DEFAULT '' COMMENT 'AVATA 操作 ID',
  `status`       tinyint(4)   NOT NULL DEFAULT 0  COMMENT '0 处理中 1 成功 2 失败',
  `create_time`  bigint(20)   NOT NULL DEFAULT 0,
  `update_time`  bigint(20)   NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_airdrops` (`airdrops_id`),
  KEY `idx_account` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='空投发放记录';

CREATE TABLE IF NOT EXISTS `tbl_snapshot` (
  `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `a_id`         bigint(20) NOT NULL DEFAULT 0 COMMENT '快照作品（持有该作品的用户获得资格）',
  `b_id`         bigint(20) NOT NULL DEFAULT 0 COMMENT '关联作品（可优先购的作品）',
  `start_time`   bigint(20) NOT NULL DEFAULT 0 COMMENT '快照开始时间',
  `status`       tinyint(4) NOT NULL DEFAULT 0 COMMENT '0 未开始 1 进行中 2 已结束',
  `snapshot_mod` tinyint(4) NOT NULL DEFAULT 0 COMMENT '0 手动 1 自动',
  `whitelist`    text                          COMMENT '手动名单：手机号-数量,手机号-数量',
  `update_time`  bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_status_start` (`status`, `start_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='持仓快照（优先购资格）';

CREATE TABLE IF NOT EXISTS `tbl_snapshot2` (
  `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `a_id`        bigint(20) NOT NULL DEFAULT 0 COMMENT '快照作品',
  `role_data`   text                          COMMENT '规则 / 关联作品（JSON）',
  `role_data2`  text                          COMMENT '规则补充（JSON）',
  `start_time`  bigint(20) NOT NULL DEFAULT 0 COMMENT '开始时间',
  `status`      tinyint(4) NOT NULL DEFAULT 0 COMMENT '0 未开始 1 已完成',
  `update_time` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_status_start` (`status`, `start_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='持仓快照（按规则空投）';

CREATE TABLE IF NOT EXISTS `tbl_snapshot_log` (
  `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_id`   bigint(20) NOT NULL DEFAULT 0 COMMENT '用户',
  `work_id`      bigint(20) NOT NULL DEFAULT 0 COMMENT '可提前购作品',
  `snapshot_id`  bigint(20) NOT NULL DEFAULT 0 COMMENT '快照 id',
  `buy_total`    int(11)    NOT NULL DEFAULT 0 COMMENT '最大可购买份数',
  `use_total`    int(11)    NOT NULL DEFAULT 0 COMMENT '已购买份数',
  `snapshot_mod` tinyint(4) NOT NULL DEFAULT 0 COMMENT '0 手动 1 自动',
  `create_time`  bigint(20) NOT NULL DEFAULT 0,
  `update_time`  bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_account_work` (`account_id`, `work_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='优先购资格明细';

-- ---------------------------------------------------------------------
-- 内容运营
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `tbl_banner` (
  `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name`        varchar(128) NOT NULL DEFAULT '' COMMENT '轮播图名称',
  `image`       varchar(255) NOT NULL DEFAULT '' COMMENT '图片链接',
  `link`        varchar(255) NOT NULL DEFAULT '' COMMENT '外链',
  `index`       int(11)      NOT NULL DEFAULT 0  COMMENT '排序',
  `produc_id`   int(11)      NOT NULL DEFAULT 0  COMMENT '关联藏品 id（后端字段名，原拼写保留）',
  `product_id`  int(11)      NOT NULL DEFAULT 0  COMMENT '关联产品 id（后台列表使用）',
  `banner_type` tinyint(4)   NOT NULL DEFAULT 0  COMMENT '类型：0 藏品 1 公告 2 外链 3 无',
  `val`         varchar(255) NOT NULL DEFAULT '' COMMENT '类型值',
  `status`      tinyint(4)   NOT NULL DEFAULT 1  COMMENT '0 禁用 1 正常',
  `start_time`  bigint(20)   NOT NULL DEFAULT 0,
  `end_time`    bigint(20)   NOT NULL DEFAULT 0,
  `create_time` bigint(20)   NOT NULL DEFAULT 0,
  `update_time` bigint(20)   NOT NULL DEFAULT 0,
  `del_time`    bigint(20)   NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='首页轮播图';

CREATE TABLE IF NOT EXISTS `tbl_notice_type` (
  `id`      bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name`    varchar(64) NOT NULL DEFAULT '' COMMENT '分类名称',
  `display` tinyint(1)  NOT NULL DEFAULT 1  COMMENT '是否展示',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='公告分类';

CREATE TABLE IF NOT EXISTS `tbl_notice` (
  `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `nid`         bigint(20)   NOT NULL DEFAULT 0  COMMENT '分类 tbl_notice_type.id',
  `title`       varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `summary`     varchar(512) NOT NULL DEFAULT '' COMMENT '摘要',
  `content`     longtext                         COMMENT '内容（HTML）',
  `display`     tinyint(1)   NOT NULL DEFAULT 1  COMMENT '是否展示',
  `create_time` bigint(20)   NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_nid` (`nid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='公告';

CREATE TABLE IF NOT EXISTS `tbl_version` (
  `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `download_url` varchar(255) NOT NULL DEFAULT '' COMMENT '安装包下载地址',
  `oldversion`   int(11)      NOT NULL DEFAULT 0  COMMENT '最低兼容版本号',
  `newversion`   int(11)      NOT NULL DEFAULT 0  COMMENT '最新版本号',
  `content`      text                             COMMENT '更新说明',
  `enforce`      tinyint(4)   NOT NULL DEFAULT 0  COMMENT '是否强制更新',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='App 版本';

-- ---------------------------------------------------------------------
-- 一期遗留：后台「产品」模块使用，后端 API 不再读取
-- ---------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `tbl_product` (
  `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `author_id`   bigint(20)    NOT NULL DEFAULT 0  COMMENT '创作者',
  `owner_id`    bigint(20)    NOT NULL DEFAULT 0  COMMENT '拥有者',
  `classify`    int(11)       NOT NULL DEFAULT 0  COMMENT '分类 id',
  `c_name`      varchar(64)   NOT NULL DEFAULT '' COMMENT '产品编号',
  `name`        varchar(128)  NOT NULL DEFAULT '' COMMENT '作品名称',
  `description` text                              COMMENT '作品描述',
  `image`       varchar(255)  NOT NULL DEFAULT '' COMMENT '图片链接',
  `price`       decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '价格',
  `stock`       int(11)       NOT NULL DEFAULT 0  COMMENT '库存',
  `stock_sell`  int(11)       NOT NULL DEFAULT 0  COMMENT '已售份数',
  `circulation` int(11)       NOT NULL DEFAULT 0  COMMENT '流通量',
  `index`       int(11)       NOT NULL DEFAULT 0  COMMENT '排序',
  `pay_type`    tinyint(4)    NOT NULL DEFAULT 0  COMMENT '支付方式',
  `tinyint`     tinyint(4)    NOT NULL DEFAULT 0  COMMENT '是否为空投（字段名原样保留）',
  `order_no`    varchar(64)   NOT NULL DEFAULT '' COMMENT '订单号',
  `token_id`    varchar(128)  NOT NULL DEFAULT '' COMMENT '链上标识',
  `tx_id`       varchar(128)  NOT NULL DEFAULT '' COMMENT '交易 id',
  `hash`        varchar(128)  NOT NULL DEFAULT '' COMMENT '哈希',
  `chain_status` tinyint(4)   NOT NULL DEFAULT 0  COMMENT '上链状态',
  `status`      tinyint(4)    NOT NULL DEFAULT 0  COMMENT '状态',
  `sale_time`   bigint(20)    NOT NULL DEFAULT 0  COMMENT '销售时间',
  `create_time` bigint(20)    NOT NULL DEFAULT 0,
  `update_time` bigint(20)    NOT NULL DEFAULT 0,
  `del_time`    bigint(20)    NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='产品（一期遗留）';

-- ---------------------------------------------------------------------
-- 初始数据
-- ---------------------------------------------------------------------
INSERT INTO `tbl_users` (`id`, `users`) VALUES (1, 0);
INSERT INTO `tbl_notice_type` (`name`, `display`) VALUES ('平台公告', 1), ('发售公告', 1);
INSERT INTO `tbl_version` (`download_url`, `oldversion`, `newversion`, `content`, `enforce`) VALUES ('', 100, 100, '初始版本', 0);

SET FOREIGN_KEY_CHECKS = 1;
