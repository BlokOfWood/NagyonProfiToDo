DROP TABLE IF EXISTS `CUSTOMERS`;
CREATE TABLE `CUSTOMERS` (
    `CustomerId` int(10) unsigned NOT NULL AUTO_INCREMENT,
    /*Name*/
    `CustomerUsername` VARCHAR(100) NOT NULL DEFAULT "",
    /*password*/
    `CustomerPassword` VARCHAR(100) NOT NULL DEFAULT "",
    /*email*/
    `CustomerEmail` VARCHAR(255) NOT NULL DEFAULT "",
    PRIMARY KEY (`CustomerId`)
)
