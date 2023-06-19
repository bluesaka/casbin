DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(255) NOT NULL DEFAULT '',
  `v0` varchar(255) NOT NULL DEFAULT '',
  `v1` varchar(255) NOT NULL DEFAULT '',
  `v2` varchar(255) NOT NULL DEFAULT '',
  `v3` varchar(255) NOT NULL DEFAULT '',
  `v4` varchar(255) NOT NULL DEFAULT '',
  `v5` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `casbin_rule` WRITE;

INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES
	(1,'g','admin_user1','admin','','','',''),
	(2,'g','guest_user1','guest','','','',''),
	(3,'p','admin','/api/v1/user','GET','','','');

UNLOCK TABLES;
