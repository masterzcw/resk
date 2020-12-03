# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 47.110.141.191 (MySQL 5.6.49-log)
# Database: resk
# Generation Time: 2020-12-02 02:06:50 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table account
# ------------------------------------------------------------

DROP TABLE IF EXISTS `account`;

CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账户ID',
  `account_no` varchar(32) NOT NULL COMMENT '账户编号,账户唯一标识 ',
  `account_name` varchar(64) NOT NULL COMMENT '账户名称,用来说明账户的简短描述,账户对应的名称或者命名，比如xxx积分、xxx零钱',
  `account_type` tinyint(2) NOT NULL COMMENT '账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户',
  `currency_code` char(3) NOT NULL DEFAULT 'CNY' COMMENT '货币类型编码：CNY人民币，EUR欧元，USD美元 。。。',
  `user_id` varchar(40) NOT NULL COMMENT '用户编号, 账户所属用户 ',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名称',
  `balance` decimal(30,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '账户可用余额',
  `status` tinyint(2) NOT NULL COMMENT '账户状态，账户状态：0账户初始化，1启用，2停用 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account_no_idx` (`account_no`) USING BTREE,
  UNIQUE KEY `idx_userid_type` (`user_id`,`account_type`) USING BTREE,
  KEY `id_user_idx` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=32938 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='账户';

LOCK TABLES `account` WRITE;
/*!40000 ALTER TABLE `account` DISABLE KEYS */;

INSERT INTO `account` (`id`, `account_no`, `account_name`, `account_type`, `currency_code`, `user_id`, `username`, `balance`, `status`, `created_at`, `updated_at`)
VALUES
	(32937,'10000020190101010000000000000001','系统红包账户',2,'CNY','100001','系统红包账户',0.000000,1,'2019-05-01 08:41:10.346','2019-05-12 09:37:55.462');

/*!40000 ALTER TABLE `account` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table account_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `account_log`;

CREATE TABLE `account_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `trade_no` varchar(32) NOT NULL COMMENT '交易单号 全局不重复字符或数字，唯一性标识 ',
  `log_no` varchar(32) NOT NULL COMMENT '流水编号 全局不重复字符或数字，唯一性标识 ',
  `account_no` varchar(32) NOT NULL COMMENT '账户编号 账户ID',
  `target_account_no` varchar(32) NOT NULL COMMENT '账户编号 账户ID',
  `user_id` varchar(40) NOT NULL COMMENT '用户编号',
  `username` varchar(64) NOT NULL COMMENT '用户名称',
  `target_user_id` varchar(40) NOT NULL COMMENT '目标用户编号',
  `target_username` varchar(64) NOT NULL COMMENT '目标用户名称',
  `amount` decimal(30,6) NOT NULL DEFAULT '0.000000' COMMENT '交易金额,该交易涉及的金额 ',
  `balance` decimal(30,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '交易后余额,该交易后的余额 ',
  `change_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义',
  `change_flag` tinyint(2) NOT NULL DEFAULT '0' COMMENT '交易变化标识：-1 出账 1为进账，枚举',
  `status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '交易状态：',
  `decs` varchar(128) NOT NULL COMMENT '交易描述 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `id_log_no_idx` (`log_no`) USING BTREE,
  KEY `id_user_idx` (`user_id`) USING BTREE,
  KEY `id_account_idx` (`account_no`) USING BTREE,
  KEY `id_trade_idx` (`trade_no`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=43209 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='账户明细';

LOCK TABLES `account_log` WRITE;
/*!40000 ALTER TABLE `account_log` DISABLE KEYS */;

INSERT INTO `account_log` (`id`, `trade_no`, `log_no`, `account_no`, `target_account_no`, `user_id`, `username`, `target_user_id`, `target_username`, `amount`, `balance`, `change_type`, `change_flag`, `status`, `decs`, `created_at`)
VALUES
	(43208,'20190501084054283000000002110000','20190501084054283000000002110000','10000020190101010000000000000001','10000020190101010000000000000001','100001','系统红包账户','100001','系统红包账户',0.000000,0.000000,0,0,0,'开户','2019-05-01 08:41:10.371');

/*!40000 ALTER TABLE `account_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goods
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goods`;

CREATE TABLE `goods` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `envelope_no` varchar(32) NOT NULL COMMENT '红包编号,红包唯一标识 ',
  `remain_amount` decimal(30,6) NOT NULL DEFAULT '0.000000' COMMENT '红包剩余金额额',
  `remain_quantity` int(10) NOT NULL COMMENT '红包剩余数量 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `envelope_no_idx` (`envelope_no`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=156 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='红包商品表的简化版';



# Dump of table goods_unsigned
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goods_unsigned`;

CREATE TABLE `goods_unsigned` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `envelope_no` varchar(32) NOT NULL COMMENT '红包编号,红包唯一标识 ',
  `remain_amount` decimal(30,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '红包剩余金额额',
  `remain_quantity` int(10) unsigned NOT NULL COMMENT '红包剩余数量 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `envelope_no_idx` (`envelope_no`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=162 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='红包商品表的无符号类型字段版, 另一个简化版';



# Dump of table red_envelope_goods
# ------------------------------------------------------------

DROP TABLE IF EXISTS `red_envelope_goods`;

CREATE TABLE `red_envelope_goods` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `envelope_no` varchar(32) NOT NULL COMMENT '红包编号,红包唯一标识 ',
  `envelope_type` tinyint(2) NOT NULL COMMENT '红包类型：普通红包，碰运气红包,过期红包',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名称',
  `user_id` varchar(40) NOT NULL COMMENT '用户编号, 红包所属用户 ',
  `blessing` varchar(64) DEFAULT NULL COMMENT '祝福语',
  `amount` decimal(30,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '红包总金额',
  `amount_one` decimal(30,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '单个红包金额，碰运气红包无效',
  `quantity` int(10) unsigned NOT NULL COMMENT '红包总数量 ',
  `remain_amount` decimal(30,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '红包剩余金额额',
  `remain_quantity` int(10) unsigned NOT NULL COMMENT '红包剩余数量 ',
  `expired_at` datetime(3) NOT NULL COMMENT '过期时间',
  `status` tinyint(2) NOT NULL COMMENT '红包/订单状态：0 创建、1 发布启用、2过期、3失效',
  `order_type` tinyint(2) NOT NULL COMMENT '订单类型：发布单、退款单 ',
  `pay_status` tinyint(2) NOT NULL COMMENT '支付状态：未支付，支付中，已支付，支付失败 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  `origin_envelope_no` varchar(32) DEFAULT '' COMMENT '原订单号',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `envelope_no_idx` (`envelope_no`) USING BTREE,
  KEY `id_user_idx` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1273 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='红包';



# Dump of table red_envelope_item
# ------------------------------------------------------------

DROP TABLE IF EXISTS `red_envelope_item`;

CREATE TABLE `red_envelope_item` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `item_no` varchar(32) NOT NULL COMMENT '红包订单详情编号 ',
  `envelope_no` varchar(32) NOT NULL COMMENT '红包编号,红包唯一标识 ',
  `recv_username` varchar(64) DEFAULT NULL COMMENT '红包接收者用户名称',
  `recv_user_id` varchar(40) NOT NULL COMMENT '红包接收者用户编号 ',
  `amount` decimal(30,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '收到金额',
  `quantity` int(10) unsigned NOT NULL COMMENT '收到数量：对于收红包来说是1 ',
  `remain_amount` decimal(30,6) unsigned NOT NULL DEFAULT '0.000000' COMMENT '收到后红包剩余金额',
  `account_no` varchar(32) NOT NULL COMMENT '红包接收者账户ID',
  `pay_status` tinyint(2) NOT NULL COMMENT '支付状态：未支付，支付中，已支付，支付失败 ',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  `desc` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `item_no_idx` (`item_no`) USING BTREE,
  KEY `envelope_no_idx` (`envelope_no`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=141 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='红包详情';




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
