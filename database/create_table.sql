CREATE TABLE `tldr_index` (
  `name` varchar(255) NOT NULL,
  `platform` json DEFAULT NULL,
  `language` json DEFAULT NULL,
  `targets` json DEFAULT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tldr_content` (
  `name` varchar(255) NOT NULL,
  `platform` varchar(255) NOT NULL,
  `language` varchar(255) NOT NULL,
  `content` text,
  PRIMARY KEY (`name`,`platform`,`language`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `tldr_hot` (
  `name` varchar(255) NOT NULL,
  `platform` varchar(255) NOT NULL,
  `language` varchar(255) NOT NULL,
  `num` int(11) DEFAULT NULL,
  PRIMARY KEY (`name`,`platform`,`language`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
