########################
# CHANGELOG MANAGEMENT #
########################

CREATE TABLE DELETED_IDS(
    `TableName` VARCHAR(64) NOT NULL DEFAULT "",
    `Id` INT NOT NULL DEFAULT 0,
    `Counter` DATETIME NOT NULL DEFAULT UTC_TIMESTAMP(),
    KEY(`TableName`,`Counter`));
    
CREATE TABLE MODIFICATIONS(
    `Id` INTEGER UNSIGNED AUTO_INCREMENT,
    `UserId` INTEGER NOT NULL,
    `Date` DATETIME NOT NULL DEFAULT UTC_TIMESTAMP(),
    PRIMARY KEY(`Id`));

DELIMITER $$

CREATE TRIGGER MODIFICATIONS_AFTER_INSERT AFTER INSERT ON MODIFICATIONS
FOR EACH ROW
BEGIN
    SET @ModificationIndex = NEW.Id;
END$$

CREATE PROCEDURE StartWithModification (IN UserId INTEGER)
BEGIN
    START TRANSACTION;
    INSERT INTO MODIFICATIONS VALUES(NULL, UserId, UTC_TIMESTAMP());
END$$     

CREATE PROCEDURE DeleteIds(IN TableName CHAR(64), IN Ids TEXT, IN IdField TEXT)
BEGIN
    SET @SkipModificationIndexing = FIND_IN_SET(TableName,@DoNotSetModificationIndex);

    IF NOT @SkipModificationIndexing THEN
        IF ISNULL(@ModificationIndex) THEN
            SIGNAL SQLSTATE '45000' 
            SET MESSAGE_TEXT = 'DELETE failed: ModificationIndex is NULL.';
        END IF;

        SET @sql = CONCAT ('UPDATE ',TableName,' SET ModificationIndex=',@ModificationIndex,' WHERE ',IdField,' IN (',Ids,');');
        PREPARE stmt FROM @sql;
        EXECUTE stmt;
    END IF;

    SET @DELETES_UPDATED = 1;

    IF NOT @SkipModificationIndexing THEN
        SAVEPOINT Updated;
    END IF;

    SET @sql = CONCAT ('DELETE FROM ',TableName,' WHERE ',Id,' IN (',Ids,');');
    PREPARE stmt FROM @sql;
    EXECUTE stmt;

    SET @DELETES_UPDATED = NULL;

    IF NOT @SkipModificationIndexing THEN
        RELEASE SAVEPOINT Updated;
    END IF;
    
    DEALLOCATE PREPARE stmt;
END$$

CREATE PROCEDURE Commit()
BEGIN
    SET @DoNotSetModificationIndex = NULL;
    SET @ModificationIndex = NULL;
    SET @DELETES_UPDATED = NULL;
    COMMIT;
END$$

CREATE PROCEDURE Rollback()
BEGIN
    SET @DoNotSetModificationIndex = NULL;
    SET @ModificationIndex = NULL;
    SET @DELETES_UPDATED = NULL;
    ROLLBACK;
END$$

CREATE PROCEDURE ClearVariables()
BEGIN
    SET @DoNotSetModificationIndex = NULL;
    SET @ModificationIndex = NULL;
    SET @DELETES_UPDATED = NULL;
END$$

DELIMITER ;


#######
# FAQ #
#######

CREATE TABLE FAQ(
    `FaqId` INTEGER AUTO_INCREMENT,
    `Question` TEXT NOT NULL DEFAULT "",
    `Answer` TEXT NOT NULL DEFAULT "",
    `SortIndex` INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY(`FaqId`));


#########
# FILES #
#########

CREATE TABLE FILES(
    `FileId` INTEGER AUTO_INCREMENT,
    `FileName` VARCHAR(255) NOT NULL DEFAULT "",
    `Name` VARCHAR(255) NOT NULL DEFAULT "", 
    `AllowedMimeTypes` TEXT NOT NULL DEFAULT "" COMMENT "Allowed mime type list, separated by comma.",
    `OnlyLogined` BOOLEAN NOT NULL DEFAULT FALSE COMMENT "Only logined users can download it.",
    `Permission` VARCHAR(64) NOT NULL DEFAULT "" COMMENT "The user need to has this permission to download it.",
    PRIMARY KEY(`FileId`));


#############
# LANGUAGES #
#############

CREATE TABLE LANGUAGES(
    `Id` CHAR(2) NOT NULL DEFAULT "",
    `Name` VARCHAR(64) NOT NULL DEFAULT "",
    `Active` BOOLEAN NOT NULL DEFAULT FALSE,
    `UiLang` BOOLEAN NOT NULL DEFAULT FALSE,
    `SortIndex` TINYINT NOT NULL DEFAULT 0
    );

########
# LOGS #
########

CREATE TABLE ADMIN_LOGS(
    `LogId` INT AUTO_INCREMENT,
    `Date` DATETIME NOT NULL DEFAULT UTC_TIMESTAMP(),
    `UserId` INT NULL DEFAULT NULL,
    `Method` ENUM("CREATE","UPDATE","DELETE") NOT NULL,
    `ItemType` VARCHAR(64) NOT NULL,
    `ItemId` INT NULL DEFAULT NULL,
    `Note` VARCHAR(255) NOT NULL DEFAULT "",
    PRIMARY KEY(`LogId`),
    KEY(`Date`));


###########
# METRICS #
###########

CREATE TABLE METRICS(
    `MetricId` INT AUTO_INCREMENT,
    `Name` VARCHAR(255) NOT NULL DEFAULT "",
    PRIMARY KEY(`MetricId`));

CREATE TABLE METRICS_DATA(
    `MetricId` INT NOT NULL,
    `TimeStamp` DATETIME NOT NULL,
    `Value` DOUBLE NOT NULL,
    KEY(`MetricId`));

CREATE TABLE METRICS_AGGREGATED_COUNTERS(
    `MetricId` INT NOT NULL,
    `TimeStamp` DATETIME NOT NULL,
    `Value` DOUBLE NULL DEFAULT NULL,
    KEY(`MetricId`),
    KEY(`TimeStamp`));
    
CREATE TABLE METRICS_AGGREGATED_GAUGES(
    `MetricId` INT NOT NULL,
    `TimeStamp` DATETIME NOT NULL,
    `MinValue` DOUBLE NULL DEFAULT NULL,
    `MaxValue` DOUBLE NULL DEFAULT NULL,
    `AvgValue` DOUBLE NULL DEFAULT NULL,
    KEY(`MetricId`),
    KEY(`TimeStamp`));


#############
# MICRODATA #
#############

CREATE TABLE MICRODATA(
    `Key` VARCHAR(255) NOT NULL DEFAULT "",
    `Value` MEDIUMTEXT NOT NULL DEFAULT "",
    PRIMARY KEY(`Key`)
    );


#########
# PAGES #
#########

CREATE TABLE PAGES(
    `PageId` INTEGER AUTO_INCREMENT,
    `Title` VARCHAR(255) NOT NULL DEFAULT "",
    `Slug` VARCHAR(255) NOT NULL DEFAULT "",
    `Content` MEDIUMTEXT NOT NULL DEFAULT "",
    PRIMARY KEY(`PageId`));


##################
# SERVER MODULES #
##################

CREATE TABLE SERVER_MODULES(
    `Id` INTEGER AUTO_INCREMENT,
    `ModuleName` VARCHAR(128) NOT NULL DEFAULT "",
    `Title` VARCHAR(255) NOT NULL DEFAULT "",
    `Enabled` BOOLEAN NOT NULL DEFAULT FALSE,
    `Event` VARCHAR(128) NOT NULL DEFAULT "",
    `Config` JSON NOT NULL DEFAULT (JSON_OBJECT()),
    `LastRunTime` DATETIME NULL DEFAULT NULL,
    `LastRunError` VARCHAR(255) NOT NULL DEFAULT "",
    `Group` VARCHAR(64) NOT NULL DEFAULT "",
    `SortIndex` INTEGER NOT NULL DEFAULT 0,
    `Hidden` BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY(`Id`));

CREATE TABLE SERVER_MODULE_QUEUE(
    `Id` INTEGER AUTO_INCREMENT,
    `Event` VARCHAR(128) NOT NULL DEFAULT "",
    PRIMARY KEY(`Id`));

CREATE TABLE SERVER_MODULE_LOGS(
    `Id` INTEGER AUTO_INCREMENT,
    `ModuleId` INTEGER NOT NULL DEFAULT 0,
    `Date` DATETIME NOT NULL DEFAULT UTC_TIMESTAMP(),
    `Log` TEXT NOT NULL DEFAULT "",
    `HasError` BOOLEAN NOT NULL DEFAULT FALSE,
    `RunDuration` INT NOT NULL DEFAULT 0 COMMENT 'msec',
    PRIMARY KEY(`Id`),
    KEY(`ModuleId`),
    KEY(`Date`));


###################
# SERVER SETTINGS #
###################

CREATE TABLE SETTINGS(
    `Settings` JSON NOT NULL DEFAULT (JSON_OBJECT()));


####################
# SERVER VARIABLES #
####################

CREATE TABLE SERVER_VARIABLES(
    `Name` VARCHAR(255) NOT NULL DEFAULT "",
    `Value` TEXT NOT NULL DEFAULT "",
    KEY(`Name`));

INSERT INTO SERVER_VARIABLES VALUES("IMAGE_CACHE_MAX_TIME","2005-07-13 21:09:00");


###################
# USER MANAGEMENT #
###################

CREATE TABLE USERS(
    `UserId` INTEGER AUTO_INCREMENT,
    `Username` VARCHAR(100) NOT NULL DEFAULT "",
    `Fullname` VARCHAR(100) NOT NULL DEFAULT "",
    `Password` VARCHAR(100) NOT NULL DEFAULT "",
    `Salt` VARCHAR(8) NOT NULL DEFAULT "",
    `Email` VARCHAR(255) NOT NULL DEFAULT "",
    `Phone` VARCHAR(30) NOT NULL DEFAULT "",
    `Active` BOOLEAN NOT NULL DEFAULT TRUE,
    `ApiEnabled` BOOLEAN NOT NULL DEFAULT FALSE,
    `ApiKey` CHAR(64) NOT NULL DEFAULT "",
    `MustChangePassword` BOOLEAN NOT NULL DEFAULT FALSE,
    `Avatar` MEDIUMBLOB NULL DEFAULT NULL,
    `AvatarContentType` VARCHAR(255) NOT NULL DEFAULT "",
    `AvatarUploadTime` DATETIME NULL DEFAULT NULL,
    PRIMARY KEY(`UserId`));

CREATE TABLE USER_GROUPS(
    `GroupId` INTEGER AUTO_INCREMENT,
    `Name` VARCHAR(100) NOT NULL DEFAULT "",
    PRIMARY KEY(`GroupId`));

CREATE TABLE USER_GROUP_CONNECTIONS(
    `UserId` INT NOT NULL DEFAULT -1,
    `GroupId` INT NOT NULL DEFAULT -1,
    UNIQUE KEY(`UserId`, `GroupId`));

CREATE TABLE `PERMISSIONS`(
    `PermissionId` VARCHAR(100) NOT NULL DEFAULT "",
    `Name` VARCHAR(100) NOT NULL DEFAULT "",
    PRIMARY KEY(`PermissionId`));

INSERT INTO `PERMISSIONS` VALUES
    ('faq', 'Gyakran ismételt kérdések'),
    ('faq.create', 'Új kérdés felvétele'),
    ('faq.read', 'Kérdések megtekintése'),
    ('faq.update', 'Kérdések módosítása'),
    ('faq.delete', 'Kérdések törlése'),
    ('microdata', 'Microdata'),
    ('microdata.read', 'Microdata megtekintése'),
    ('microdata.update', 'Microdata módosítása'),
    ('pages', 'Egyedi tartalmú oldalak'),
    ('pages.read', 'Oldal megtekintése'),
    ('pages.update', 'Oldal módosítása'),
    ('permissions', 'Engedélyek'),
    ('permissions.read', 'Engedélyek megtekintése'),
    ('usergroups', 'Felhasználói csoportok'),
    ('usergroups.create', 'Felhasználói csoportok felvétele'),
    ('usergroups.read', 'Felhasználói csoportok megtekintése'),
    ('usergroups.update', 'Felhasználói csoportok módosítása'),
    ('usergroups.delete', 'Felhasználói csoportok törlése'),
    ('users', 'Felhasználók'),
    ('users.create', 'Felhasználók felvétele'),
    ('users.read', 'Felhasználók megtekintése'),
    ('users.update', 'Felhasználók módosítása'),
    ('users.delete', 'Felhasználók törlése');

CREATE TABLE USER_GROUP_PERMISSIONS(
    `GroupId` INT NOT NULL DEFAULT -1,
    `PermissionId` VARCHAR(100) NOT NULL DEFAULT -1,
    UNIQUE KEY(`GroupId`, `PermissionId`));

INSERT INTO USERS(`UserId`, `Username`,`Fullname`,`Password`,`Active`) VALUES(NULL,"Fox","CodeFox Admin","$2a$04$rZ3MUCiJ0NqyPaUDydcSw.d2lK89OkS986h3VQpu8xiAIft0IUIem",1);
UPDATE USERS SET `UserId`=0;
INSERT INTO USER_GROUPS (`Name`) VALUES ('Adminisztrátor');
INSERT INTO PERMISSIONS (`PermissionId`, `Name`) VALUES ('godmode','Mindenhez van jogosultsága');
INSERT INTO USER_GROUP_CONNECTIONS VALUES (0,1);
INSERT INTO USER_GROUP_PERMISSIONS VALUES (1,"godmode");

CREATE TABLE USER_SESSIONS(
    `Token` CHAR(128) NOT NULL,
    `UserId` INTEGER NULL DEFAULT NULL,
    `Expiration` DATETIME NULL DEFAULT NULL,
    `Vars` BLOB  NOT NULL DEFAULT "{}",
    `Permissions` BLOB  NOT NULL DEFAULT "{}",
    `Language` CHAR(2) NOT NULL DEFAULT "",
    PRIMARY KEY(`Token`));

CREATE TABLE USER_LOCKS(
    `Token` CHAR(128) NOT NULL DEFAULT "",
    `Module` VARCHAR(50) NOT NULL DEFAULT "",
    `Id` VARCHAR(50) NOT NULL DEFAULT "",
    `Updated` DATETIME NOT NULL DEFAULT UTC_TIMESTAMP(),
    UNIQUE KEY(`Token`, `Module`, `Id`));
	
CREATE TABLE USER_SETTINGS(
    `UserId` INTEGER NOT NULL DEFAULT 0,
    `Module` VARCHAR(64) NOT NULL DEFAULT "",
    `Name` VARCHAR(64) NOT NULL DEFAULT "",
    `Data` MEDIUMTEXT NOT NULL DEFAULT "{}",
    PRIMARY KEY(`UserId`,`Module`, `Name`));