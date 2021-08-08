/*
 Navicat Premium Data Transfer

 Source Server         : test1
 Source Server Type    : MySQL
 Source Server Version : 80026
 Source Host           : localhost:3306
 Source Schema         : test01

 Target Server Type    : MySQL
 Target Server Version : 80026
 File Encoding         : 65001

 Date: 08/08/2021 08:12:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for clothes
-- ----------------------------
DROP TABLE IF EXISTS `clothes`;
CREATE TABLE `clothes`  (
  `clothing code` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci NOT NULL,
  `clothing size` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci NOT NULL,
  `clothing price` int(0) NOT NULL,
  `clothing type` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci NOT NULL,
  PRIMARY KEY (`clothing code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_czech_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of clothes
-- ----------------------------
INSERT INTO `clothes` VALUES ('A01', 's', 100, 'female');
INSERT INTO `clothes` VALUES ('A02', 'm', 88, 'famale');
INSERT INTO `clothes` VALUES ('A03', 's', 98, 'male');
INSERT INTO `clothes` VALUES ('B01', 'l', 115, 'male');

-- ----------------------------
-- Table structure for storage
-- ----------------------------
DROP TABLE IF EXISTS `storage`;
CREATE TABLE `storage`  (
  `storage code` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci NOT NULL,
  `storage volume` int(0) NOT NULL,
  PRIMARY KEY (`storage code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_czech_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of storage
-- ----------------------------
INSERT INTO `storage` VALUES ('A', 10);
INSERT INTO `storage` VALUES ('B', 14);
INSERT INTO `storage` VALUES ('C', 17);

-- ----------------------------
-- Table structure for supplier
-- ----------------------------
DROP TABLE IF EXISTS `supplier`;
CREATE TABLE `supplier`  (
  `supplier code` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci NOT NULL,
  `supplier name` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci NOT NULL,
  PRIMARY KEY (`supplier code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_czech_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of supplier
-- ----------------------------
INSERT INTO `supplier` VALUES ('1', 'zhongshan');
INSERT INTO `supplier` VALUES ('2', 'doudou');
INSERT INTO `supplier` VALUES ('3', 'sense');
INSERT INTO `supplier` VALUES ('4', 'Rafina');

-- ----------------------------
-- Table structure for supply
-- ----------------------------
DROP TABLE IF EXISTS `supply`;
CREATE TABLE `supply`  (
  `clothes code` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci NOT NULL,
  `supplier infomation` int(0) NOT NULL,
  `clothes quality degree` varchar(255) CHARACTER SET utf8 COLLATE utf8_czech_ci NOT NULL,
  PRIMARY KEY (`clothes code`, `supplier infomation`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_czech_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of supply
-- ----------------------------
INSERT INTO `supply` VALUES ('A01', 1, 'qualified');
INSERT INTO `supply` VALUES ('A02', 2, 'unqualified');
INSERT INTO `supply` VALUES ('A03', 3, 'qualified');
INSERT INTO `supply` VALUES ('B01', 4, 'unqualified');

SET FOREIGN_KEY_CHECKS = 1;
