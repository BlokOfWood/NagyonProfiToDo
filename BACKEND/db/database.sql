SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for todos
-- ----------------------------
DROP TABLE IF EXISTS `Todos`;
CREATE TABLE `Todos` (
  `TodoID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `UserID` int(10) unsigned NOT NULL,
  `Name` varchar(64) NOT NULL,
  `Description` text NOT NULL,
  `Priority` enum('Critical','Urgent','Important','Normal','Eventually') NOT NULL,
  `Done` tinyint(1) NOT NULL DEFAULT 0 COMMENT "Boolean",
  `Deadline` datetime NOT NULL,
  `CreatedAt` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`TodoID`),
  KEY `UserID` (`UserID`),
  FOREIGN KEY (`UserID`) REFERENCES `Users` (`UserID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `Users`;
CREATE TABLE `Users` (
  `UserID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `SessionID` char(16) DEFAULT NULL,
  `Name` varchar(128) NOT NULL,
  `Email` varchar(128) NOT NULL,
  `PasswordHash` char(64) DEFAULT NULL,
  `Salt` char(8) DEFAULT NULL,
  PRIMARY KEY (`UserID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
