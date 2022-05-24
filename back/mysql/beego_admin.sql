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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `demo_crud`
--

/*!40000 ALTER TABLE `demo_crud` DISABLE KEYS */;
/*!40000 ALTER TABLE `demo_crud` ENABLE KEYS */;

--
-- Table structure for table `sys_dept`
--

DROP TABLE IF EXISTS `sys_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dept` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(11) NOT NULL DEFAULT '-1' COMMENT '父级id',
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
/*!40000 ALTER TABLE `sys_dept` ENABLE KEYS */;

--
-- Table structure for table `sys_dict`
--

DROP TABLE IF EXISTS `sys_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_dict` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) NOT NULL COMMENT '字典名称',
  `serial_number` varchar(255) DEFAULT NULL COMMENT '字典编号',
  `describe` varchar(255) DEFAULT NULL COMMENT '描述',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict`
--

/*!40000 ALTER TABLE `sys_dict` DISABLE KEYS */;
INSERT INTO `sys_dict` VALUES (3,'','','','2022-05-23 10:43:25',NULL,NULL),(5,'','123','上','2022-05-23 10:44:53',NULL,NULL),(6,'Knsfhftu Btst Kmield','555','Qghoqs Ovikhhdc Vhyt Rjhpbv Urptdxp','2022-05-23 10:45:54',NULL,NULL),(7,'Wpuh Lroidvzvb Dlnzbbxye Rsigsj Wnjx','123','运','2022-05-23 15:24:22',NULL,NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_item`
--

/*!40000 ALTER TABLE `sys_dict_item` DISABLE KEYS */;
INSERT INTO `sys_dict_item` VALUES (6,'Dcyilrsr Iczl Nxjfkgy Yimkr Foojaiyqa Ncdvdsixh','Cfhhij Pplhryuc Sywwer Vflhkl Rmkrdh Ugscugrco',7,'51231231','2022-05-23 15:24:40',NULL,NULL);
/*!40000 ALTER TABLE `sys_dict_item` ENABLE KEYS */;

--
-- Table structure for table `sys_menu`
--

DROP TABLE IF EXISTS `sys_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_num` int(11) NOT NULL COMMENT '排序',
  `redirect` varchar(100) DEFAULT NULL COMMENT '重定向地址',
  `icon` varchar(100) DEFAULT NULL COMMENT '图标',
  `hidden` tinyint(4) DEFAULT NULL COMMENT '是否隐藏',
  `is_home` tinyint(4) DEFAULT NULL COMMENT '是否首页',
  `parent_id` int(11) NOT NULL DEFAULT '-1' COMMENT '父级id',
  `title` varchar(32) NOT NULL COMMENT '菜单名称',
  `type` tinyint(4) NOT NULL COMMENT '菜单类型： 1. 目录 2. 菜单  3. 按钮',
  `perms` varchar(100) DEFAULT NULL COMMENT '权限标识',
  `name` varchar(100) DEFAULT NULL COMMENT '菜单标识',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menu`
--

/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;
INSERT INTO `sys_menu` VALUES (60,1,'','',0,0,0,'',0,'','队场光花','2022-05-23 10:32:40',NULL,NULL);
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
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `remark` varchar(100) NOT NULL COMMENT '角色备注',
  `role_name` varchar(100) NOT NULL COMMENT '角色名称',
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
INSERT INTO `sys_role` VALUES (6,'测试','test','2022-05-23 11:11:06',NULL,NULL),(7,'测试1','test1','2022-05-23 14:41:24',NULL,NULL);
/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;

--
-- Table structure for table `sys_role_dept`
--

DROP TABLE IF EXISTS `sys_role_dept`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_dept` (
  `id` int(11) NOT NULL,
  `role_id` int(11) DEFAULT NULL COMMENT '角色ID',
  `dept_id` int(11) DEFAULT NULL COMMENT '部门ID',
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
  `menu_id` int(11) NOT NULL COMMENT '菜单id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色与菜单对应关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_menu`
--

/*!40000 ALTER TABLE `sys_role_menu` DISABLE KEYS */;
INSERT INTO `sys_role_menu` VALUES (61,6,60,'2022-05-23 16:09:00',NULL,NULL);
/*!40000 ALTER TABLE `sys_role_menu` ENABLE KEYS */;

--
-- Table structure for table `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `nick_name` varchar(255) NOT NULL COMMENT '昵称',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `dept_id` varchar(255) NOT NULL COMMENT '部门id',
  `phone_num` varchar(255) DEFAULT NULL COMMENT '手机号码',
  `status` tinyint(4) NOT NULL COMMENT '所属状态是否有效  1是有效 0是失效',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `account` varchar(255) NOT NULL COMMENT '账户',
  `password` varchar(500) DEFAULT NULL COMMENT '密码',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user`
--

/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;
INSERT INTO `sys_user` VALUES (31,'周凯','851708184@qq.com','1','123456789101',1,'http://dummyimage.com/300x250','zhoukai','ab261b409c8c97ecfd9bf90f24f9b321','2022-05-23 11:12:10',NULL,NULL),(32,'周凯测试 123456','851708184@qq.com','1','123456789101',1,'http://dummyimage.com/125x125','zhoukai','e10adc3949ba59abbe56e057f20f883e','2022-05-23 17:05:44',NULL,NULL);
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
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户与角色对应关系';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_role`
--

/*!40000 ALTER TABLE `sys_user_role` DISABLE KEYS */;
INSERT INTO `sys_user_role` VALUES (26,32,6,'2022-05-23 15:59:09',NULL,NULL);
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

SET @@GLOBAL.GTID_PURGED='4e25c795-d9b2-11ec-98bc-00163e0055d4:1-10248';
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-24 11:38:09
