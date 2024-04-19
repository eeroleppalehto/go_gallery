CREATE TABLE `user` (
  `user_id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `password` varchar(100) NOT NULL,
  `creted_at` datetime NOT NULL DEFAULT NOW(),
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_id_UNIQUE` (`user_id`),
  UNIQUE KEY `username_UNIQUE` (`username`),
  UNIQUE KEY `email_UNIQUE` (`email`)
);

CREATE TABLE `photo` (
  `photo_id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `title` varchar(45) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `filename` varchar(255) NOT NULL,
  `date` datetime NOT NULL,
  PRIMARY KEY (`photo_id`),
  UNIQUE KEY `photo_id_UNIQUE` (`photo_id`),
  KEY `fk_photo_user_idx` (`user_id`),
  CONSTRAINT `fk_photo_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
);