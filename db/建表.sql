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

 Date: 16/04/2023 18:50:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for 53986_2023
-- ----------------------------
DROP TABLE IF EXISTS `53986_2023`;
CREATE TABLE `53986_2023`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `date` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日期',
  `day_info` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '白天天气',
  `day_temperature` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '白天温度',
  `day_direct` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '白天风向',
  `day_power` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '白天风力',
  `night_info` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '夜间天气',
  `night_temperature` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '夜间温度',
  `night_direct` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '夜间风向',
  `night_power` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '夜间风力',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 53986_2023
-- ----------------------------
INSERT INTO `53986_2023` VALUES (1, '2023-04-17', '阴', '27', '西南风', '3~4级', '多云', '17', '东北风', '3~4级');
INSERT INTO `53986_2023` VALUES (2, '2023-04-18', '晴', '26', '东北风', '4~5级', '多云', '15', '东北风', '3~4级');
INSERT INTO `53986_2023` VALUES (3, '2023-04-19', '晴', '27', '南风', '微风', '多云', '15', '南风', '微风');
INSERT INTO `53986_2023` VALUES (4, '2023-04-20', '小雨', '31', '东北风', '5~6级', '小雨', '9', '东北风', '5~6级');
INSERT INTO `53986_2023` VALUES (5, '2023-04-21', '小雨', '15', '东北风', '4~5级', '中雨', '7', '东北风', '3~4级');
INSERT INTO `53986_2023` VALUES (6, '2023-04-22', '小雨', '10', '东北风', '3~4级', '小雨', '5', '东北风', '3~4级');
INSERT INTO `53986_2023` VALUES (7, '2023-04-18', '多云', '26', '东北风', '4~5级', '多云', '15', '东北风', '3~4级');

-- ----------------------------
-- Table structure for 53986r_2023
-- ----------------------------
DROP TABLE IF EXISTS `53986r_2023`;
CREATE TABLE `53986r_2023`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `date` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日期',
  `time` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '时间',
  `temperature` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '温度',
  `humidity` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '相对湿度',
  `rain` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '降水量mm',
  `icomfort` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '舒适度\r\n(4很热，极不适应\r\n3热，很不舒适\r\n2暖，不舒适\r\n1温暖，较舒适\r\n0舒适，最可接受\r\n-1凉爽，较舒适\r\n-2凉，不舒适\r\n-3冷，很不舒适\r\n-4很冷，极不适应)',
  `info` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '天气',
  `feelst` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '体感温度',
  `wind_direct` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '风向',
  `wind_power` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '风力',
  `wind_speed` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '风速',
  `warn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '预警',
  `aqi` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '空气质量',
  `aq` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '1 优 空气质量令人满意,基本无空气污染,各类人群可正常活动\r\n2 良 空气质量可接受,但某些污染物可能,对极少数异常,敏感人群健康有较弱影响,极少数异常敏感人群应减少户外活动\r\n3 轻度污染 易感人群症状有轻度加剧,健康人群出现刺激症状,儿童、老年人及心脏病、呼吸系统疾病患者应减少长时间、高强度的户外锻炼\r\n4 中度污染 进一步加剧易感人群症状,可能对健康人群心脏、呼吸系统有影响,儿童、老年人及心脏病、呼吸系统疾病患者避免长时间、高强度的户外锻炼,一般人群适量减少户外运动\r\n5 重度污染 心脏病和肺病患者症状显著加剧,运动耐受力减低,健康人群普遍出现症状,老年人和心脏病、肺病患者应停留在室内，停止户外活动，一般人群减少户外活动\r\n6 严重污染 健康人运动耐力减低,有显著强烈症状,提前出现某些疾病,老年人和病人应当留在室内，避免体力消耗，一般人群应避免户外活动',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of 53986r_2023
-- ----------------------------
INSERT INTO `53986r_2023` VALUES (1, '2023-04-16', '17:45', '22.2', '41.0', '0.0', '0', '多云', '20.5', '东南风', '微风', '1.6', '', '72', '2');
INSERT INTO `53986r_2023` VALUES (2, '2023-04-16', '17:50', '22.2', '39.0', '0.0', '0', '多云', '20.3', '东南风', '微风', '2.0', '', '72', '2');
INSERT INTO `53986r_2023` VALUES (3, '2023-04-16', '17:55', '22.2', '39.0', '0.0', '0', '多云', '20.3', '东南风', '微风', '1.9', '', '72', '2');
INSERT INTO `53986r_2023` VALUES (4, '2023-04-16', '18:25', '21.6', '43.0', '0.0', '0', '多云', '20.1', '东南风', '微风', '1.3', '', '72', '2');

-- ----------------------------
-- Table structure for location
-- ----------------------------
DROP TABLE IF EXISTS `location`;
CREATE TABLE `location`  (
  `id` int(11) NOT NULL,
  `stationid` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `country` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `city` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `valid` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of location
-- ----------------------------
INSERT INTO `location` VALUES (1, '53986', '中国', '河南省', '新乡市', 1);
INSERT INTO `location` VALUES (2, '54455', '中国', '辽宁省', '兴城市', 1);
INSERT INTO `location` VALUES (3, '54526', '中国', '天津市', '东丽区', 1);
INSERT INTO `location` VALUES (4, '54453', '中国', '辽宁省', '葫芦岛市', 0);

SET FOREIGN_KEY_CHECKS = 1;
