-- MySQL dump 10.13  Distrib 5.7.38, for osx10.17 (x86_64)
--
-- Host: rm-hp32w80372fy6clynoo.mysql.huhehaote.rds.aliyuncs.com    Database: beego-admin
-- ------------------------------------------------------
-- Server version	5.7.37-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN= 0;

--
-- Table structure for table `demo_crud`
--

DROP TABLE IF EXISTS `demo_crud`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `demo_crud` (
  `title` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` varchar(255) DEFAULT NULL,
  `type` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `demo_crud`
--

/*!40000 ALTER TABLE `demo_crud` DISABLE KEYS */;
INSERT INTO `demo_crud` VALUES ('5',3,'5','5',NULL,NULL,NULL),('23',4,'555','545',NULL,NULL,NULL),('',5,'','','2022-05-24 17:28:22',NULL,NULL),('',6,'团行织府阶','装','2022-05-24 17:28:27',NULL,NULL),('123',7,'555','55','2022-05-24 17:35:55',NULL,NULL),('123312',8,'132312','123312','2022-05-24 17:36:13',NULL,NULL),('123',9,'555','123','2022-05-25 10:36:26',NULL,NULL),('123',10,'1','1','2022-05-25 10:36:31',NULL,NULL),('1',11,'1','1','2022-05-25 10:36:35',NULL,NULL),('21',12,'1','1','2022-05-25 10:46:59',NULL,NULL),('1',13,'1','1','2022-05-25 10:47:02',NULL,NULL),('',14,'型元家些','复','2022-05-25 17:22:27',NULL,NULL),('',15,'酸国月写','制','2022-05-25 17:23:33',NULL,NULL),('该说断会志约',16,'更元装信','要','2022-05-25 17:23:35',NULL,NULL),('该说断会志约',17,'更元装信','要','2022-05-25 17:27:07',NULL,NULL),('该说断会志约',18,'更元装信','要','2022-05-25 17:27:33',NULL,NULL),('',19,'','','2022-05-25 17:29:21',NULL,NULL),('发头因照',20,'名查断条','月','2022-05-25 17:29:24',NULL,NULL),('发头因照',21,'名查断条','月','2022-05-25 17:29:47',NULL,NULL),('发头因照',22,'名查断条','月','2022-05-25 17:31:51',NULL,NULL);
/*!40000 ALTER TABLE `demo_crud` ENABLE KEYS */;

--
-- Table structure for table `log_login`
--

DROP TABLE IF EXISTS `log_login`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_login` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `equipment` varchar(255) DEFAULT NULL COMMENT '设备型号',
  `ip` varchar(255) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL COMMENT '用户id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `log_login_sys_user_id_fk` (`user_id`),
  CONSTRAINT `log_login_sys_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COMMENT='登录日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_login`
--

/*!40000 ALTER TABLE `log_login` DISABLE KEYS */;
INSERT INTO `log_login` VALUES (1,'Mac','1.24.24.216',11,'2022-05-25 15:54:17',NULL,NULL),(2,'Mac','1.24.24.216',11,'2022-05-25 15:58:34',NULL,NULL),(3,'Mac','1.24.24.216',11,'2022-05-25 15:59:15',NULL,NULL),(4,'Mac','1.24.24.216',11,'2022-05-25 16:00:15',NULL,NULL),(5,'Mac','1.24.24.216',11,'2022-05-25 16:01:39',NULL,NULL),(6,'Mac','1.24.24.216',11,'2022-05-25 16:05:01',NULL,NULL),(7,'Mac','1.24.24.216',11,'2022-05-25 16:11:01',NULL,NULL),(8,'Mac','1.24.24.216',11,'2022-05-25 16:19:14',NULL,NULL),(9,'Mac','1.24.24.216',11,'2022-05-25 16:20:33',NULL,NULL),(10,'Mac','1.24.24.216',11,'2022-05-25 16:31:34',NULL,NULL),(11,'Mac','1.24.24.216',11,'2022-05-25 16:32:56',NULL,NULL),(12,'Mac','1.24.24.216',11,'2022-05-25 16:33:17',NULL,NULL),(13,'Mac','1.24.24.216',11,'2022-05-25 16:34:07',NULL,NULL),(14,'Mac','1.24.24.216',11,'2022-05-25 16:34:59',NULL,NULL),(15,'Mac','1.24.24.216',11,'2022-05-25 16:35:35',NULL,NULL),(16,'Mac','1.24.24.216',11,'2022-05-25 16:37:53',NULL,NULL),(17,'Mac','1.24.24.216',11,'2022-05-25 16:38:13',NULL,NULL),(18,'Mac','1.24.24.216',11,'2022-05-25 16:38:28',NULL,NULL),(19,'Mac','1.24.24.216',11,'2022-05-25 16:39:00',NULL,NULL),(20,'Mac','1.24.24.216',11,'2022-05-25 16:45:38',NULL,NULL),(21,'Mac','1.24.24.216',11,'2022-05-25 16:47:51',NULL,NULL),(22,'Mac','1.24.24.216',11,'2022-05-25 16:48:18',NULL,NULL),(23,'Mac','1.24.24.216',11,'2022-05-25 16:48:40',NULL,NULL),(24,'Mac','1.24.24.216',11,'2022-05-25 16:49:19',NULL,NULL),(25,'Mac','1.24.24.216',11,'2022-05-25 16:51:45',NULL,NULL),(26,'Mac','1.24.24.216',11,'2022-05-25 17:37:43',NULL,NULL);
/*!40000 ALTER TABLE `log_login` ENABLE KEYS */;

--
-- Table structure for table `log_request`
--

DROP TABLE IF EXISTS `log_request`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_request` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `params` varchar(255) DEFAULT NULL,
  `request_address` varchar(255) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int(32) DEFAULT NULL,
  `method` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `log_request_sys_user_id_fk` (`user_id`),
  CONSTRAINT `log_request_sys_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_request`
--

/*!40000 ALTER TABLE `log_request` DISABLE KEYS */;
INSERT INTO `log_request` VALUES (1,'{\"content\":\"名查断条\",\"title\":\"发头因照\",\"type\":\"月\"}','/api/v1/demoCrud','','2022-05-25 17:31:51',NULL,NULL,11,'POST'),(2,'','/api/v1/logLog?pageNum=1&pageSize=10','','2022-05-25 17:35:08',NULL,NULL,11,'GET'),(3,'','/api/v1/logLog?pageNum=1&pageSize=10','','2022-05-25 17:36:03',NULL,NULL,11,'GET'),(4,'{\"pageNum\":[\"1\"],\"pageSize\":[\"10\"]}','/api/v1/logLog?pageNum=1&pageSize=10','','2022-05-25 17:37:17',NULL,NULL,11,'GET'),(5,'{}','/api/v1/userInfo','1.24.24.216','2022-05-25 17:37:43',NULL,NULL,11,'GET'),(6,'{\"pageNum\":[\"1\"],\"pageSize\":[\"10\"]}','/api/v1/logLog?pageNum=1&pageSize=10','1.24.24.216','2022-05-25 17:37:43',NULL,NULL,11,'GET'),(7,'{\"pageNum\":[\"1\"],\"pageSize\":[\"10\"]}','/api/v1/logRequest?pageNum=1&pageSize=10','1.24.24.216','2022-05-25 17:38:00',NULL,NULL,11,'GET'),(8,'{\"pageNum\":[\"1\"],\"pageSize\":[\"10\"]}','/api/v1/logRequest?pageNum=1&pageSize=10','1.24.24.216','2022-05-25 17:38:12',NULL,NULL,11,'GET');
/*!40000 ALTER TABLE `log_request` ENABLE KEYS */;

--
-- Table structure for table `sys_dept`
--

DROP TABLE IF EXISTS `sys_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dept` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(11) NOT NULL COMMENT '父级id',
  `name` varchar(50) NOT NULL COMMENT '部门名称',
  `order_num` int(11) NOT NULL COMMENT '排序',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='部门管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dept`
--

/*!40000 ALTER TABLE `sys_dept` DISABLE KEYS */;
INSERT INTO `sys_dept` VALUES (1,0,'测试',0,NULL,'2021-11-12 10:19:19',NULL);
/*!40000 ALTER TABLE `sys_dept` ENABLE KEYS */;

--
-- Table structure for table `sys_dict`
--

DROP TABLE IF EXISTS `sys_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dict` (
  `name` varchar(255) NOT NULL COMMENT '字典名称',
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `serial_number` varchar(255) DEFAULT NULL COMMENT '字典编号',
  `describe` varchar(255) DEFAULT NULL COMMENT '描述',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict`
--

/*!40000 ALTER TABLE `sys_dict` DISABLE KEYS */;
INSERT INTO `sys_dict` VALUES ('测试哈123',1,'','',NULL,NULL,NULL),('测试',2,'','',NULL,NULL,NULL),('512312313',4,'123','1231233123','2022-05-24 17:08:22',NULL,NULL);
/*!40000 ALTER TABLE `sys_dict` ENABLE KEYS */;

--
-- Table structure for table `sys_dict_item`
--

DROP TABLE IF EXISTS `sys_dict_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dict_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `label` varchar(100) NOT NULL COMMENT '名称',
  `value` varchar(100) NOT NULL COMMENT '数据值',
  `dict_id` int(11) NOT NULL COMMENT '字典id',
  `describe` varchar(255) DEFAULT NULL COMMENT '描述',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_item`
--

/*!40000 ALTER TABLE `sys_dict_item` DISABLE KEYS */;
INSERT INTO `sys_dict_item` VALUES (1,'','',0,'',NULL,NULL,NULL),(2,'','',0,'',NULL,NULL,NULL),(3,'55','55',1,'555',NULL,NULL,NULL),(4,'123','55555',1,'555',NULL,NULL,NULL);
/*!40000 ALTER TABLE `sys_dict_item` ENABLE KEYS */;

--
-- Table structure for table `sys_menu`
--

DROP TABLE IF EXISTS `sys_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_menu` (
  `order_num` int(11) NOT NULL COMMENT '排序',
  `redirect` varchar(100) DEFAULT NULL COMMENT '重定向地址',
  `icon` varchar(100) DEFAULT NULL COMMENT '图标',
  `hidden` tinyint(4) DEFAULT NULL COMMENT '是否隐藏',
  `is_home` tinyint(4) DEFAULT NULL COMMENT '是否首页',
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(11) NOT NULL COMMENT '父级id',
  `title` varchar(32) NOT NULL COMMENT '菜单名称',
  `type` tinyint(4) NOT NULL COMMENT '菜单类型： 1. 目录 2. 菜单  3. 按钮',
  `perms` varchar(100) DEFAULT NULL COMMENT '权限标识',
  `name` varchar(100) DEFAULT NULL COMMENT '菜单标识',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu`
--

/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;
INSERT INTO `sys_menu` VALUES (1,NULL,NULL,0,0,29,0,'系统管理',1,NULL,'Sys',NULL,'2021-11-12 16:51:33',NULL),(0,NULL,NULL,0,0,30,29,'用户管理',2,NULL,'SysUser',NULL,NULL,NULL),(1,NULL,NULL,0,0,31,29,'角色管理',2,NULL,'SysRole',NULL,NULL,NULL),(2,NULL,NULL,0,0,32,29,'部门管理',2,NULL,'SysDepart',NULL,NULL,NULL),(3,NULL,NULL,0,0,33,29,'菜单管理',2,NULL,'SysMenu',NULL,NULL,NULL),(4,NULL,NULL,0,0,34,29,'字典管理',2,NULL,'SysDict',NULL,NULL,NULL),(0,NULL,NULL,NULL,NULL,35,33,'添加',3,'add','add',NULL,NULL,NULL),(1,NULL,NULL,NULL,NULL,36,33,'修改',3,'edit','edit',NULL,NULL,NULL),(2,NULL,NULL,NULL,NULL,37,33,'删除',3,'del','del',NULL,NULL,NULL),(0,'','',0,0,39,30,'添加',3,'add','add',NULL,NULL,NULL),(1,'','',0,0,40,30,'修改',3,'edit','edit',NULL,NULL,NULL),(2,'','',0,0,41,30,'删除',3,'del','del',NULL,NULL,NULL),(0,'','',0,0,42,32,'添加',3,'add','add',NULL,NULL,NULL),(0,'','',0,0,43,32,'修改',3,'edit','edit',NULL,NULL,NULL),(1,'','',0,0,44,32,'删除',3,'del','del',NULL,NULL,NULL),(0,'','',0,0,45,31,'添加',3,'add','add',NULL,NULL,NULL),(1,'','',0,0,46,31,'编辑',3,'edit','edit',NULL,NULL,NULL),(2,'','',0,0,48,31,'删除',3,'del','del',NULL,NULL,NULL),(0,'','',0,0,49,30,'分配角色',3,'power','power',NULL,'2021-11-12 11:13:59',NULL),(2,'','',0,0,50,31,'分配菜单',3,'power','power',NULL,NULL,NULL),(0,'','',0,0,51,34,'添加',3,'add','add',NULL,NULL,NULL),(1,'','',0,0,52,34,'修改',3,'edit','edit',NULL,NULL,NULL),(2,'','',0,0,53,34,'删除',3,'del','del',NULL,NULL,NULL),(4,'','',0,0,54,34,'字典项设置',3,'setting','setting',NULL,NULL,NULL),(0,'','',0,1,55,0,'首页',2,'','ViewsHome',NULL,'2021-11-12 16:56:22',NULL),(5,'','',0,0,56,30,'修改密码',3,'updatePassword','updateName',NULL,'2021-12-20 22:33:17',NULL),(1,'','',0,0,57,0,'crud',1,'','Crud',NULL,'2021-12-21 21:17:42',NULL),(0,'','',0,0,58,57,'avue展示基础',2,'','DemoAvue',NULL,'2021-12-21 21:20:19',NULL),(0,'','',0,0,60,0,'日志',1,'','Journal','2022-05-25 11:23:52',NULL,NULL),(0,'','',0,0,61,60,'登录日志',2,'','LoginLog','2022-05-25 11:24:30',NULL,NULL),(1,'','',0,0,62,60,'接口请求',2,'','Request','2022-05-25 11:28:43',NULL,NULL);
/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;

--
-- Table structure for table `sys_oss`
--

DROP TABLE IF EXISTS `sys_oss`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_oss` (
  `id` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '图片等url链接',
  `location` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件存放位置，暂存本地服务器，不采用云 oss',
  `type` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件类型',
  `size` int(11) NOT NULL COMMENT '文件大小 size',
  `old_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '原文件名称',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_oss`
--

/*!40000 ALTER TABLE `sys_oss` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_oss` ENABLE KEYS */;

--
-- Table structure for table `sys_role`
--

DROP TABLE IF EXISTS `sys_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role` (
  `remark` varchar(100) NOT NULL COMMENT '角色备注',
  `role_name` varchar(100) NOT NULL COMMENT '角色名称',
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role`
--

/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;
INSERT INTO `sys_role` VALUES ('测试','test',1,NULL,'2021-11-12 15:58:06',NULL),('超级管理员','admin',4,NULL,'2021-11-12 15:58:12',NULL),('123123312','55555',6,'2022-05-24 16:44:39',NULL,NULL);
/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;

--
-- Table structure for table `sys_role_dept`
--

DROP TABLE IF EXISTS `sys_role_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_dept` (
  `id` varchar(32) NOT NULL,
  `role_id` varchar(32) DEFAULT NULL COMMENT '角色ID',
  `dept_id` varchar(32) DEFAULT NULL COMMENT '部门ID',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色与部门对应关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_dept`
--

/*!40000 ALTER TABLE `sys_role_dept` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_role_dept` ENABLE KEYS */;

--
-- Table structure for table `sys_role_menu`
--

DROP TABLE IF EXISTS `sys_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `menu_id` varchar(255) NOT NULL COMMENT '菜单id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=125 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色与菜单对应关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_menu`
--

/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;
INSERT INTO `sys_role_menu` VALUES (45,1,'30',NULL,NULL,NULL),(52,1,'39',NULL,NULL,NULL),(53,1,'49',NULL,NULL,NULL),(54,1,'40',NULL,NULL,NULL),(55,1,'41',NULL,NULL,NULL),(56,1,'29',NULL,NULL,NULL),(84,0,'29','2022-05-24 16:57:09',NULL,NULL),(85,0,'30','2022-05-24 16:57:09',NULL,NULL),(86,0,'39','2022-05-24 16:57:09',NULL,NULL),(87,0,'49','2022-05-24 16:57:09',NULL,NULL),(88,0,'40','2022-05-24 16:57:09',NULL,NULL),(89,0,'41','2022-05-24 16:57:09',NULL,NULL),(90,6,'29','2022-05-24 16:57:30',NULL,NULL),(91,6,'30','2022-05-24 16:57:30',NULL,NULL),(92,6,'39','2022-05-24 16:57:30',NULL,NULL),(93,6,'49','2022-05-24 16:57:30',NULL,NULL),(94,4,'29','2022-05-25 11:37:39',NULL,NULL),(95,4,'30','2022-05-25 11:37:39',NULL,NULL),(96,4,'31','2022-05-25 11:37:39',NULL,NULL),(97,4,'32','2022-05-25 11:37:39',NULL,NULL),(98,4,'33','2022-05-25 11:37:39',NULL,NULL),(99,4,'34','2022-05-25 11:37:39',NULL,NULL),(100,4,'35','2022-05-25 11:37:39',NULL,NULL),(101,4,'36','2022-05-25 11:37:39',NULL,NULL),(102,4,'37','2022-05-25 11:37:39',NULL,NULL),(103,4,'39','2022-05-25 11:37:39',NULL,NULL),(104,4,'40','2022-05-25 11:37:39',NULL,NULL),(105,4,'41','2022-05-25 11:37:39',NULL,NULL),(106,4,'42','2022-05-25 11:37:39',NULL,NULL),(107,4,'43','2022-05-25 11:37:39',NULL,NULL),(108,4,'44','2022-05-25 11:37:39',NULL,NULL),(109,4,'45','2022-05-25 11:37:39',NULL,NULL),(110,4,'46','2022-05-25 11:37:39',NULL,NULL),(111,4,'48','2022-05-25 11:37:39',NULL,NULL),(112,4,'49','2022-05-25 11:37:39',NULL,NULL),(113,4,'50','2022-05-25 11:37:39',NULL,NULL),(114,4,'51','2022-05-25 11:37:39',NULL,NULL),(115,4,'52','2022-05-25 11:37:39',NULL,NULL),(116,4,'53','2022-05-25 11:37:39',NULL,NULL),(117,4,'54','2022-05-25 11:37:39',NULL,NULL),(118,4,'55','2022-05-25 11:37:39',NULL,NULL),(119,4,'56','2022-05-25 11:37:39',NULL,NULL),(120,4,'57','2022-05-25 11:37:39',NULL,NULL),(121,4,'58','2022-05-25 11:37:39',NULL,NULL),(122,4,'60','2022-05-25 11:37:39',NULL,NULL),(123,4,'61','2022-05-25 11:37:39',NULL,NULL),(124,4,'62','2022-05-25 11:37:39',NULL,NULL);
/*!40000 ALTER TABLE `sys_role_menu` ENABLE KEYS */;

--
-- Table structure for table `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user` (
  `account` varchar(255) NOT NULL COMMENT '账户',
  `nick_name` varchar(255) NOT NULL COMMENT '昵称',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `dept_id` int(11) NOT NULL COMMENT '部门id',
  `phone_num` varchar(255) DEFAULT NULL COMMENT '手机号码',
  `status` tinyint(4) NOT NULL COMMENT '所属状态是否有效  1是有效 0是失效',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user`
--

/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;
INSERT INTO `sys_user` VALUES ('test','5555',NULL,1,NULL,1,NULL,10,'e10adc3949ba59abbe56e057f20f883e',NULL,'2021-12-20 22:55:26',NULL),('admin','超级管理员',NULL,1,NULL,1,NULL,11,'e10adc3949ba59abbe56e057f20f883e',NULL,NULL,NULL),('5555555555','123123123','851708184@qq.com',1,'',1,'',15,'e10adc3949ba59abbe56e057f20f883e','2022-05-24 16:24:00',NULL,NULL);
/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;

--
-- Table structure for table `sys_user_role`
--

DROP TABLE IF EXISTS `sys_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL COMMENT '角色id',
  `role_id` varchar(255) NOT NULL COMMENT '角色id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户与角色对应关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_role`
--

/*!40000 ALTER TABLE `sys_user_role` DISABLE KEYS */;
INSERT INTO `sys_user_role` VALUES (13,10,'4',NULL,NULL,NULL),(14,10,'1',NULL,NULL,NULL),(16,11,'4','2022-05-24 16:37:50',NULL,NULL),(17,11,'1','2022-05-24 16:37:50',NULL,NULL),(18,15,'1','2022-05-24 16:37:55',NULL,NULL);
/*!40000 ALTER TABLE `sys_user_role` ENABLE KEYS */;

--
-- Table structure for table `test`
--

DROP TABLE IF EXISTS `test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `test` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `value` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test`
--

/*!40000 ALTER TABLE `test` DISABLE KEYS */;
INSERT INTO `test` VALUES (1,'123'),(2,'222');
/*!40000 ALTER TABLE `test` ENABLE KEYS */;

--
-- GTID state at the end of the backup 
--

SET @@GLOBAL.GTID_PURGED='4e25c795-d9b2-11ec-98bc-00163e0055d4:1-17793';
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-25 17:39:14
