/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : inventory

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 08/01/2022 21:36:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for inventory
-- ----------------------------
DROP TABLE IF EXISTS `inventory`;
CREATE TABLE `inventory` (
  `inventory_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `location` varchar(30) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `comment` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `inventory_admini` varchar(15) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `capacity` bigint DEFAULT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for inventory_item
-- ----------------------------
DROP TABLE IF EXISTS `inventory_item`;
CREATE TABLE `inventory_item` (
  `item_id` varchar(30) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `thumbnail` longtext CHARACTER SET utf8 COLLATE utf8_bin,
  `comment` varchar(30) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `from_location` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `current_location` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `to_location` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `original_price` bigint NOT NULL,
  `current_price` bigint NOT NULL,
  `weight` bigint NOT NULL,
  `url` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin;

-- ----------------------------
-- Records of inventory_item
-- ----------------------------
BEGIN;
INSERT INTO `inventory_item` VALUES ('1479666178681999360', '../../uploads/thumbs/131-thumbnail.png', 'sds', 'NY', 'NY', 'NY', 4500, 860, 34034, 'dsds', 2, '');
INSERT INTO `inventory_item` VALUES ('1479674837138935808', '', 'sds', 'NY', 'NY', 'NY', 4500, 860, 34034, 'dsds', 4, 'sas');
INSERT INTO `inventory_item` VALUES ('1479988642867843072', '', '5555', 'NY', 'NY', 'NY', 4500, 860, 34034, '6324', 5, 'sword');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
