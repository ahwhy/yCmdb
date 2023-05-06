SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for resource_cost
-- ----------------------------
DROP TABLE IF EXISTS `resource_cost`;
CREATE TABLE `resource_cost` (
  `resource_id` varchar(255) NOT NULL COMMENT '资源Id',
  `pay_mode` tinyint(2) NOT NULL COMMENT '支付方式',
  `pay_mode_detail` text NOT NULL COMMENT '字符说明',
  `sale_price` decimal(12,4) NOT NULL COMMENT '官网价,原价（分）',
  `real_cost` decimal(12,4) NOT NULL COMMENT '实际支付金额（分）',
  `policy` float(4,2) NOT NULL COMMENT '折扣率',
  `unit_price` decimal(12,4) NOT NULL COMMENT '单价（分）',
  PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for resource_host
-- ----------------------------
DROP TABLE IF EXISTS `resource_host`;
CREATE TABLE `resource_host` (
  `resource_id` varchar(64) NOT NULL COMMENT '关联的资源Id',
  `cpu` tinyint(4) NOT NULL COMMENT 'cpu核数',
  `memory` int(13) NOT NULL COMMENT '内存大小',
  `gpu_amount` tinyint(4) DEFAULT NULL COMMENT 'gpu核数',
  `gpu_spec` varchar(255) DEFAULT NULL COMMENT 'gpu规格',
  `os_type` varchar(255) DEFAULT NULL COMMENT '操作系统类型',
  `os_name` varchar(255) DEFAULT NULL COMMENT '操作系统名称',
  `serial_number` varchar(120) DEFAULT NULL COMMENT '系统序列号',
  `image_id` char(64) DEFAULT NULL COMMENT '镜像Id',
  `internet_max_bandwidth_out` int(10) DEFAULT NULL COMMENT '外网最大出口带宽',
  `internet_max_bandwidth_in` int(10) DEFAULT NULL COMMENT '外网最大入口带宽',
  `key_pair_name` varchar(255) DEFAULT NULL COMMENT 'ssh key关联Id',
  `security_groups` varchar(255) DEFAULT NULL COMMENT '安全组Id列表',
  PRIMARY KEY (`resource_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='服务器主机信息';

-- ----------------------------
-- Table structure for resource_meta
-- ----------------------------
DROP TABLE IF EXISTS `resource_meta`;
CREATE TABLE `resource_meta` (
  `id` char(64) NOT NULL COMMENT '全局唯一Id, 直接使用个云商自己的Id',
  `domain` varchar(64) NOT NULL COMMENT '资源所属域',
  `namespace` varchar(255) NOT NULL COMMENT '资源所属空间',
  `env` varchar(255) NOT NULL COMMENT '资源所属环境',
  `sync_at` bigint(13) NOT NULL COMMENT '同步时间',
  `credential_id` varchar(64) NOT NULL COMMENT '关联的同于同步的credential id',
  `create_at` int(10) NOT NULL COMMENT '创建时间',
  `serial_number` varchar(255) NOT NULL COMMENT '序列号',
  `usage_mode` tinyint(2) NOT NULL COMMENT '使用方式',
  `shared_policy` text NOT NULL COMMENT '共享策略, 当一个资源被多个应用共享时, 可以指定允许的应用',
  `spec_hash` varchar(255) NOT NULL COMMENT '规格数据Hash',
  `cost_hash` varchar(255) NOT NULL COMMENT '费用数据Hash',
  `status_hash` varchar(255) NOT NULL COMMENT '状态数据Hash',
  `tag_hash` varchar(255) NOT NULL COMMENT '标签数据Hash',
  `relation_hash` varchar(255) NOT NULL COMMENT '关系数据Hash',
  `custom_hash` varchar(255) NOT NULL COMMENT '资源自定义属性Hash',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_domain` (`domain`) USING HASH,
  KEY `idx_namespace` (`namespace`) USING HASH,
  KEY `idx_env` (`env`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源基础信息,所有资源的公共信息, 用于全局解索';

-- ----------------------------
-- Table structure for resource_metric
-- ----------------------------
DROP TABLE IF EXISTS `resource_metric`;
CREATE TABLE `resource_metric` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '资源Id',
  `resource_type` tinyint(4) NOT NULL COMMENT '适于那种资源',
  `group` varchar(255) NOT NULL COMMENT '分组描述',
  `name` varchar(255) NOT NULL COMMENT '指标名称',
  `description` text NOT NULL COMMENT '指标描述',
  `expr_temp` varchar(255) NOT NULL COMMENT '表达式模版',
  `expr_type` tinyint(4) NOT NULL COMMENT '表达式类型',
  `unit` varchar(64) NOT NULL COMMENT '指标单位',
  `create_at` bigint(20) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_rt` (`resource_type`) USING BTREE,
  KEY `idx_group` (`group`) USING BTREE,
  KEY `idx_expr` (`expr_temp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for resource_rds
-- ----------------------------
DROP TABLE IF EXISTS `resource_rds`;
CREATE TABLE `resource_rds` (
  `resource_id` varchar(64) NOT NULL COMMENT '资源Id',
  `engine_type` varchar(255) NOT NULL COMMENT '引擎 比如 MYSQL, SQLServer, PGSQL',
  `engine_version` varchar(255) NOT NULL COMMENT '引擎版本',
  `instance_class` varchar(255) NOT NULL COMMENT '实例规格: 对应ALI(DBInstanceClass)',
  `class_type` varchar(255) NOT NULL COMMENT '实例规格族，取值：s：共享型；x：通用型；d：独享套餐；h：独占物理机',
  `export_type` varchar(255) NOT NULL COMMENT '实例是内网或外网 (Internet：外网/Intranet：内网)',
  `network_type` varchar(255) NOT NULL COMMENT '实例的网络类型 (Classic：经典网络/VPC：专有网络。)',
  `type` varchar(255) NOT NULL COMMENT '实例类型 Primary：主实例, Readonly：只读实例, Guard：灾备实例, Temp：临时实例',
  `cpu` int(11) NOT NULL COMMENT 'CPU 核数',
  `memory` int(11) NOT NULL COMMENT '实例内存，单位：M。',
  `db_max_quantity` int(11) NOT NULL COMMENT '一个实例下可创建最大数据库数量',
  `account_max_quantity` int(11) NOT NULL COMMENT '可创建账号的最大数量',
  `max_connections` int(11) NOT NULL COMMENT '最大并发连接数',
  `max_iops` int(11) NOT NULL COMMENT '最大每秒IO请求次数',
  `collation` varchar(255) NOT NULL COMMENT '系统字符集排序规则',
  `time_zone` varchar(64) NOT NULL COMMENT '时区',
  `storage_capacity` int(11) NOT NULL COMMENT '实例存储空间，单位：GB。',
  `storage_type` varchar(255) NOT NULL COMMENT '实例储存类型 local_ssd/ephemeral_ssd：本地SSD盘, cloud_ssd：SSD云盘；cloud_essd：ESSD云盘',
  `security_ip_mode` varchar(255) NOT NULL COMMENT '安全名单模式, 默认白名单',
  `security_ip_list` text NOT NULL COMMENT 'IP白名单',
  `connection_mode` varchar(255) NOT NULL COMMENT '实例的访问模式，取值：Standard：标准访问模式；Safe：数据库代理模式。',
  `ip_type` varchar(255) NOT NULL COMMENT 'IP类型',
  `lock_mode` varchar(255) NOT NULL COMMENT '实例锁定模式; Unlock：正常；ManualLock：手动触发锁定；LockByExpiration：实例过期自动锁定；LockByRestoration：实例回滚前的自动锁定；LockByDiskQuota：实例空间满自动锁定',
  `lock_reason` varchar(255) NOT NULL COMMENT '锁定原因',
  `deploy_mode` varchar(255) NOT NULL COMMENT '部署模式(腾讯云独有)',
  `port` int(11) NOT NULL COMMENT '端口',
  `extra` text NOT NULL COMMENT '额外的无法通用的一些属性, 比如只有腾讯云独有的一些属性',
  PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='rds实例信息';

-- ----------------------------
-- Table structure for resource_spec
-- ----------------------------
DROP TABLE IF EXISTS `resource_spec`;
CREATE TABLE `resource_spec` (
  `resource_id` varchar(64) NOT NULL COMMENT '资源Id',
  `vendor` tinyint(2) NOT NULL COMMENT '资源厂商',
  `resource_type` tinyint(2) NOT NULL COMMENT '资源类型',
  `region` varchar(255) NOT NULL COMMENT '地域',
  `zone` varchar(255) NOT NULL COMMENT '区域',
  `owner` varchar(255) NOT NULL COMMENT '资源归属账号',
  `name` varchar(255) NOT NULL COMMENT '资源名称',
  `category` varchar(255) NOT NULL COMMENT '资源种类',
  `type` varchar(255) NOT NULL COMMENT '规格',
  `description` varchar(255) NOT NULL COMMENT '描述',
  `expire_at` int(10) NOT NULL COMMENT '过期时间',
  `update_at` int(10) NOT NULL COMMENT '更新时间',
  `release_protection` tinyint(1) NOT NULL COMMENT '是否开启实例释放保护',
  `cpu` int(5) NOT NULL COMMENT '资源占用Cpu数量',
  `gpu` int(5) NOT NULL COMMENT 'GPU数量',
  `memory` int(10) NOT NULL COMMENT '资源使用的内存, 单位M',
  `storage` int(10) NOT NULL COMMENT '资源使用的存储, 单位G',
  `band_width` int(10) NOT NULL COMMENT '公网IP带宽, 单位M',
  `extra` text NOT NULL COMMENT '额外的通用属性',
  PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for resource_status
-- ----------------------------
DROP TABLE IF EXISTS `resource_status`;
CREATE TABLE `resource_status` (
  `resource_id` varchar(64) NOT NULL COMMENT '资源Id',
  `phase` varchar(255) NOT NULL COMMENT '资源当前状态',
  `lock_mode` varchar(255) NOT NULL COMMENT '实例锁定模式; Unlock：正常；ManualLock：手动触发锁定；LockByExpiration：实例过期自动锁定；LockByRestoration：实例回滚前的自动锁定；LockByDiskQuota：实例空间满自动锁定',
  `lock_reason` varchar(255) NOT NULL COMMENT '锁定原因',
  `public_ip` varchar(255) NOT NULL COMMENT '公网IP, 或者域名',
  `private_ip` varchar(255) NOT NULL COMMENT '内网IP, 或者域名',
  PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for resource_tag
-- ----------------------------
DROP TABLE IF EXISTS `resource_tag`;
CREATE TABLE `resource_tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '标签Id',
  `t_key` varchar(255) NOT NULL COMMENT '标签的名称',
  `t_value` varchar(255) NOT NULL COMMENT '标签的值',
  `description` varchar(255) NOT NULL COMMENT '值的描述信息',
  `resource_id` varchar(64) NOT NULL COMMENT '标签关联的资源Id',
  `weight` int(11) NOT NULL COMMENT '标签权重',
  `type` tinyint(4) NOT NULL COMMENT '标签类型',
  `create_at` bigint(13) NOT NULL COMMENT '标签创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_id` (`t_key`,`t_value`,`resource_id`) COMMENT '一个资源同一个key value只允许有一对',
  KEY `idx_key` (`t_key`) USING HASH,
  KEY `idx_value` (`t_value`) USING BTREE,
  KEY `idx_resource_id` (`resource_id`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='资源标签';

-- ----------------------------
-- Table structure for secret
-- ----------------------------
DROP TABLE IF EXISTS `secret`;
CREATE TABLE `secret` (
  `id` varchar(64) NOT NULL COMMENT '凭证Id',
  `create_at` bigint(13) NOT NULL COMMENT '创建时间',
  `description` varchar(255) NOT NULL COMMENT '凭证描述',
  `vendor` tinyint(1) NOT NULL COMMENT '资源提供商',
  `address` varchar(255) NOT NULL COMMENT '体验提供方访问地址',
  `allow_regions` text NOT NULL COMMENT '允许同步的Region列表',
  `crendential_type` tinyint(1) NOT NULL COMMENT '凭证类型',
  `api_key` varchar(255) NOT NULL COMMENT '凭证key',
  `api_secret` text NOT NULL COMMENT '凭证secret',
  `request_rate` int(11) NOT NULL COMMENT '请求速率',
  `domain` varchar(255) NOT NULL COMMENT '所属域',
  `namespace` varchar(255) NOT NULL COMMENT '所属空间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_key` (`api_key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源提供商同步凭证管理';

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `id` varchar(64) NOT NULL COMMENT '任务Id',
  `region` varchar(64) NOT NULL COMMENT '资源所属Region',
  `resource_type` tinyint(1) NOT NULL COMMENT '资源类型',
  `credential_id` varchar(64) NOT NULL COMMENT '用于操作资源的凭证Id',
  `credential_desc` text NOT NULL COMMENT '凭证描述',
  `timeout` int(11) NOT NULL COMMENT '任务超时时间',
  `status` tinyint(1) NOT NULL COMMENT '任务当前状态',
  `message` text NOT NULL COMMENT '任务失败相关信息',
  `start_at` bigint(20) NOT NULL COMMENT '任务开始时间',
  `end_at` bigint(20) NOT NULL COMMENT '任务结束时间',
  `total_succeed` int(11) NOT NULL COMMENT '总共操作成功的资源数量',
  `total_failed` int(11) NOT NULL COMMENT '总共操作失败的资源数量',
  `domain` varchar(255) NOT NULL COMMENT '所属域',
  `namespace` varchar(255) NOT NULL COMMENT '所属空间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源操作任务管理';

-- ----------------------------
-- Table structure for task_record
-- ----------------------------
DROP TABLE IF EXISTS `task_record`;
CREATE TABLE `task_record` (
  `instance_id` varchar(64) NOT NULL COMMENT '实例Id',
  `instance_name` varchar(255) NOT NULL COMMENT '实例名称',
  `is_success` tinyint(2) NOT NULL COMMENT '是否同步成功',
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '通过异常时的信息',
  `task_id` varchar(64) NOT NULL COMMENT 'task id',
  `create_at` bigint(20) NOT NULL COMMENT '记录创建时间',
  KEY `idx_task` (`task_id`) USING HASH,
  KEY `idx_instance` (`instance_id`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='任务操作详情';

SET FOREIGN_KEY_CHECKS = 1;