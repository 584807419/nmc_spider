/*
 Navicat MySQL Data Transfer

 Source Server         : M_ITX_192.168.1.16 root Zk19924599
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 192.168.1.16:3306
 Source Schema         : www_nmc_cn

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 18/04/2023 15:55:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for province
-- ----------------------------
DROP TABLE IF EXISTS `province`;
CREATE TABLE `province`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '江苏',
  `abbr` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '江苏:http://www.nmc.cn/rest/province/AJS?_=1681720916608',
  `valid` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of province
-- ----------------------------
INSERT INTO `province` VALUES (1, '江苏', 'AJS', 1);
INSERT INTO `province` VALUES (2, '河南', 'AHA', 1);
INSERT INTO `province` VALUES (3, '北京', 'ABJ', 1);
INSERT INTO `province` VALUES (4, '天津', 'ATJ', 1);
INSERT INTO `province` VALUES (5, '河北', 'AHE', 1);
INSERT INTO `province` VALUES (6, '山西', 'ASX', 1);
INSERT INTO `province` VALUES (7, '内蒙古', 'ANM', 1);
INSERT INTO `province` VALUES (8, '辽宁', 'ALN', 1);
INSERT INTO `province` VALUES (9, '吉林', 'AJL', 1);
INSERT INTO `province` VALUES (10, '黑龙江', 'AHL', 1);
INSERT INTO `province` VALUES (11, '上海', 'ASH', 1);
INSERT INTO `province` VALUES (12, '浙江', 'AZJ', 1);
INSERT INTO `province` VALUES (13, '安徽', 'AAH', 1);
INSERT INTO `province` VALUES (14, '福建', 'AFJ', 1);
INSERT INTO `province` VALUES (15, '江西', 'AJX', 1);
INSERT INTO `province` VALUES (16, '山东', 'ASD', 1);
INSERT INTO `province` VALUES (17, '湖北', 'AHB', 1);
INSERT INTO `province` VALUES (18, '湖南', 'AHN', 1);
INSERT INTO `province` VALUES (19, '广东', 'AGD', 1);
INSERT INTO `province` VALUES (20, '广西', 'AGX', 1);
INSERT INTO `province` VALUES (21, '海南', 'AHI', 1);
INSERT INTO `province` VALUES (22, '重庆', 'ACQ', 1);
INSERT INTO `province` VALUES (23, '四川', 'ASC', 1);
INSERT INTO `province` VALUES (24, '贵州', 'AGZ', 1);
INSERT INTO `province` VALUES (25, '云南', 'AYN', 1);
INSERT INTO `province` VALUES (26, '西藏', 'AXZ', 1);
INSERT INTO `province` VALUES (27, '陕西', 'ASN', 1);
INSERT INTO `province` VALUES (28, '甘肃', 'AGS', 1);
INSERT INTO `province` VALUES (29, '青海', 'AQH', 1);
INSERT INTO `province` VALUES (30, '宁夏', 'ANX', 1);
INSERT INTO `province` VALUES (31, '新疆', 'AXJ', 1);
INSERT INTO `province` VALUES (32, '香港', 'AXG', 1);
INSERT INTO `province` VALUES (33, '澳门', 'AAM', 1);
INSERT INTO `province` VALUES (34, '台湾', 'ATW', 1);

SET FOREIGN_KEY_CHECKS = 1;
