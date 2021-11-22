-- MySQL dump 10.13  Distrib 5.7.34, for Linux (x86_64)
--
-- Host: localhost    Database: bieshu-oa
-- ------------------------------------------------------
-- Server version	5.7.34-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `user_name` varchar(80) DEFAULT NULL COMMENT '用户名',
  `real_name` varchar(80) DEFAULT NULL COMMENT '真实姓名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `mobile` varchar(20) DEFAULT NULL COMMENT '手机号码',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(3) DEFAULT NULL COMMENT '状态（1-开启，0-关闭）',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `group_id` int(11) DEFAULT NULL COMMENT '组织架构',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `created_id` int(11) DEFAULT NULL COMMENT '添加人',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `updated_id` int(11) DEFAULT NULL COMMENT '最后更新人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `deleted_id` int(11) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_user_name` (`user_name`) USING BTREE,
  KEY `admin_real_name` (`real_name`),
  KEY `admin_status` (`status`),
  KEY `admin_deleted_at` (`deleted_at`),
  KEY `admin_moblie` (`mobile`),
  KEY `admin_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,'admin','超级管理员','$2a$10$xu3oikTgUGScwac6txad0uUwlDffgTUXlfRS9WTAUQ.LRyYEdr.Dq','123123','111',1,'http://127.0.0.1:8199/admin-avatar/cfsr6ef25drgmqlroo.jpg',13,'2021-11-12 11:10:16',NULL,'2021-11-18 17:25:56',NULL,NULL,NULL),(2,'admin11','测试账号','$2a$10$yIYmoHUnTePICFV2NJYJDu37Bu93GMnBizAggnLZaj/IwURMk477i','','',1,'',15,'2021-11-12 11:22:59',NULL,'2021-11-20 11:16:24',1,NULL,NULL),(3,'admin1133','超级管理员','$2a$10$yAEEks4ePrT1jbtVmYQxUerhcBw7yqsvzSEwj6caOCw05vH2aoA3q','','',1,'',14,'2021-11-12 13:25:56',NULL,'2021-11-20 11:08:12',1,'2021-11-20 11:08:19',1),(4,'ceshi123123','测试ddddd','$2a$10$mUOwdvLMNBWGiy64rRvGuObvNc3o36FsYmqwVml3gbg6DE5q9kHgO','','',1,'',14,'2021-11-18 17:26:34',NULL,'2021-11-20 10:32:01',NULL,NULL,NULL);
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin_role`
--

DROP TABLE IF EXISTS `admin_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin_role` (
  `admin_id` int(11) NOT NULL COMMENT '管理员ID',
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`admin_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员角色关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_role`
--

LOCK TABLES `admin_role` WRITE;
/*!40000 ALTER TABLE `admin_role` DISABLE KEYS */;
INSERT INTO `admin_role` VALUES (1,1),(1,14),(2,14);
/*!40000 ALTER TABLE `admin_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '组织架构ID',
  `name` varchar(150) DEFAULT NULL COMMENT '组织架构名称',
  `pid` int(11) DEFAULT '0' COMMENT '组织架构上级ID',
  `link` varchar(255) DEFAULT NULL COMMENT '组织架构路径',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `created_id` int(11) DEFAULT NULL COMMENT '添加人',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `updated_id` int(11) DEFAULT NULL COMMENT '最后更新人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `deleted_id` int(11) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`),
  KEY `group_pid` (`pid`),
  KEY `group_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COMMENT='组织架构表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group`
--

LOCK TABLES `group` WRITE;
/*!40000 ALTER TABLE `group` DISABLE KEYS */;
INSERT INTO `group` VALUES (13,'销售中心',0,'0:','2021-11-15 16:35:04',NULL,'2021-11-15 16:47:43',NULL,NULL,NULL),(14,'销售一组',13,'0:13:','2021-11-15 16:35:12',NULL,'2021-11-15 16:47:43',NULL,NULL,NULL),(15,'销售二组',13,'0:13:','2021-11-15 16:35:25',NULL,'2021-11-15 16:47:43',NULL,NULL,NULL),(16,'财务部',0,'0:','2021-11-15 16:35:34',NULL,'2021-11-15 17:17:19',NULL,'2021-11-20 10:17:43',NULL),(17,'asdasd',16,'0:16:','2021-11-15 16:38:21',NULL,'2021-11-15 16:47:43',NULL,'2021-11-15 17:08:34',NULL),(18,'ddd',16,'0:16:','2021-11-15 16:38:54',NULL,'2021-11-15 16:47:43',NULL,'2021-11-15 17:08:36',NULL),(19,'333',16,'0:16:','2021-11-15 17:11:45',NULL,'2021-11-15 17:11:45',NULL,'2021-11-15 17:56:54',NULL);
/*!40000 ALTER TABLE `group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `name` varchar(150) DEFAULT NULL COMMENT '菜单名称',
  `router` varchar(255) DEFAULT NULL COMMENT '请求地址',
  `method` tinyint(3) DEFAULT NULL COMMENT '请求类型(1 = "Get",2 = "Post",3 = "Put",4 = "DELETE")',
  `key` varchar(80) DEFAULT NULL COMMENT '标识',
  `pid` int(11) DEFAULT '0' COMMENT '菜单上级ID',
  `type` tinyint(3) DEFAULT NULL COMMENT '菜单类型（1= "目录"2 = “菜单”，3 = “按钮”，4 = “隐藏”）',
  `link` varchar(255) DEFAULT NULL COMMENT '菜单路径',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `created_id` int(11) DEFAULT NULL COMMENT '添加人',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `updated_id` int(11) DEFAULT NULL COMMENT '最后更新人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `deleted_id` int(11) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`),
  KEY `menu_pid` (`pid`),
  KEY `menu_deleted_at` (`deleted_at`),
  KEY `menu_router` (`router`),
  KEY `menu_method` (`method`),
  KEY `menu_key` (`key`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES (2,'组织架构列表','/group',1,'Group-Index',7,2,'0:7:','2021-11-10 16:53:36',NULL,'2021-11-15 17:38:12',NULL,NULL,NULL),(3,'新增组织架构','/group',2,'Group-Store',2,3,'0:7:2:','2021-11-10 16:54:40',NULL,'2021-11-15 17:37:32',NULL,NULL,NULL),(5,'更新组织架构','/group/:id',3,'Group-Update',2,3,'0:7:2:','2021-11-15 15:14:22',NULL,'2021-11-16 15:37:43',NULL,NULL,NULL),(6,'删除组织架构','/group/:id',4,'Group-Delete',2,3,'0:7:2:','2021-11-15 15:14:48',NULL,'2021-11-16 15:37:46',NULL,NULL,NULL),(7,'系统设置','',0,'System',0,1,'0:','2021-11-15 15:22:28',NULL,'2021-11-15 16:14:50',NULL,NULL,NULL),(8,'管理员列表','/admin',1,'Admin-Index',7,2,'0:7:','2021-11-15 17:44:59',NULL,'2021-11-15 17:45:21',NULL,NULL,NULL),(9,'菜单列表','/menu',1,'Menu-Index',7,2,'0:7:','2021-11-15 17:46:48',NULL,'2021-11-15 17:47:02',NULL,NULL,NULL),(10,'新增菜单','/menu',2,'Menu-Store',9,3,'0:7:9:','2021-11-15 17:49:04',NULL,'2021-11-15 17:49:04',NULL,NULL,NULL),(11,'更新菜单','/menu/:id',3,'Menu-Update',9,3,'0:7:9:','2021-11-15 17:49:37',NULL,'2021-11-16 15:37:14',NULL,NULL,NULL),(12,'删除菜单','/menu/:id',4,'Menu-Delete',9,3,'0:7:9:','2021-11-15 17:50:05',NULL,'2021-11-16 15:37:24',NULL,NULL,NULL),(13,'新增管理员','/admin',2,'Admin-Store',8,3,'0:7:8:','2021-11-15 17:51:15',NULL,'2021-11-15 17:51:15',NULL,NULL,NULL),(14,'更新管理员','/admin/:id',3,'Admin-Update',8,3,'0:7:8:','2021-11-15 17:51:38',NULL,'2021-11-16 15:37:36',NULL,NULL,NULL),(15,'删除管理员','/delete/:id',4,'Admin-Delete',8,3,'0:7:8:','2021-11-15 17:52:11',NULL,'2021-11-16 15:37:30',NULL,NULL,NULL),(16,'角色列表','/role',1,'Role-Index',7,2,'0:7:','2021-11-16 15:14:35',NULL,'2021-11-16 15:14:59',NULL,NULL,NULL),(17,'新增角色','/role',2,'Role-Store',16,3,'0:7:16:','2021-11-16 15:15:21',NULL,'2021-11-16 15:15:21',NULL,NULL,NULL),(18,'更新角色','/role/:id',3,'Role-Update',16,3,'0:7:16:','2021-11-16 15:15:53',NULL,'2021-11-16 15:37:05',NULL,NULL,NULL),(19,'删除角色','/role/:id',4,'Role-Delete',16,3,'0:7:16:','2021-11-16 15:16:18',NULL,'2021-11-16 15:37:17',NULL,NULL,NULL),(20,'角色详情','/role/:id',1,'Role-Show',16,4,'0:7:16:','2021-11-16 15:36:52',NULL,'2021-11-16 15:36:52',NULL,NULL,NULL),(21,'管理员详情','/admin/:id',1,'Admin-Show',8,4,'0:7:8:','2021-11-16 15:38:23',NULL,'2021-11-16 15:38:23',NULL,NULL,NULL);
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(150) DEFAULT NULL COMMENT '角色名称',
  `dp` tinyint(3) DEFAULT '0' COMMENT '数据权限（0-仅自己，1-所在组织架构和下级组织架构，2-所有人）',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `created_id` int(11) DEFAULT NULL COMMENT '添加人',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `updated_id` int(11) DEFAULT NULL COMMENT '最后更新人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `deleted_id` int(11) DEFAULT NULL COMMENT '删除人',
  PRIMARY KEY (`id`),
  KEY `role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role`
--

LOCK TABLES `role` WRITE;
/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'超级管理员角色',2,'2021-11-16 10:39:50',NULL,'2021-11-16 13:45:55',NULL,NULL,NULL),(14,'测试角色',0,'2021-11-16 10:39:50',NULL,'2021-11-18 11:30:52',NULL,NULL,NULL),(15,'测试111',0,'2021-11-16 10:44:25',NULL,'2021-11-16 11:49:14',NULL,'2021-11-16 11:49:59',NULL);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_menu`
--

DROP TABLE IF EXISTS `role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_menu` (
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `menu_id` int(11) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_menu`
--

LOCK TABLES `role_menu` WRITE;
/*!40000 ALTER TABLE `role_menu` DISABLE KEYS */;
INSERT INTO `role_menu` VALUES (14,2),(14,3),(14,5),(14,7);
/*!40000 ALTER TABLE `role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'bieshu-oa'
--

--
-- Dumping routines for database 'bieshu-oa'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-11-20  5:38:47
