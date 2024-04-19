CREATE DATABASE gollery;

use gollery;

CREATE TABLE `user` (
  `iduser` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `password` varchar(100) NOT NULL,
  PRIMARY KEY (`iduser`),
  UNIQUE KEY `iduser_UNIQUE` (`iduser`),
  UNIQUE KEY `username_UNIQUE` (`username`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `photo` (
  `idphoto` int unsigned NOT NULL AUTO_INCREMENT,
  `iduser` int unsigned NOT NULL,
  `title` varchar(45) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `filename` varchar(255) NOT NULL,
  `date` datetime NOT NULL,
  PRIMARY KEY (`idphoto`),
  UNIQUE KEY `idphoto_UNIQUE` (`idphoto`),
  KEY `fk_photo_user_idx` (`iduser`),
  CONSTRAINT `fk_photo_user` FOREIGN KEY (`iduser`) REFERENCES `user` (`iduser`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;