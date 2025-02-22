CREATE DATABASE gollery;

use gollery;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO user (username, email, password)
VALUES ('admin', 'admin@gollery.com', 'admin');

INSERT INTO photo (user_id, title, description, filename, date)
VALUES 
    (1, 'Test Photo 1', 'This is a test photo', 'image-1.jpg', NOW()),
    (1, 'Test Photo 2', 'This is a test photo', 'image-2.jpg', NOW()),
    (1, 'Test Photo 3', 'This is a test photo', 'image-3.jpg', NOW()),
    (1, 'Test Photo 4', 'This is a test photo', 'image-4.jpg', NOW()),
    (1, 'Test Photo 5', 'This is a test photo', 'image-5.jpg', NOW()),
    (1, 'Test Photo 6', 'This is a test photo', 'image-6.jpg', NOW()),
    (1, 'Test Photo 7', 'This is a test photo', 'image-7.jpg', NOW()),
    (1, 'Test Photo 8', 'This is a test photo', 'image-8.jpg', NOW()),
    (1, 'Test Photo 9', 'This is a test photo', 'image-9.jpg', NOW()),
    (1, 'Test Photo 10', 'This is a test photo', 'image-10.jpg', NOW()),
    (1, 'Test Photo 11', 'This is a test photo', 'image-11.jpg', NOW()),
    (1, 'Test Photo 12', 'This is a test photo', 'image-12.jpg', NOW()),
    (1, 'Test Photo 13', 'This is a test photo', 'image-13.jpg', NOW()),
    (1, 'Test Photo 14', 'This is a test photo', 'image-14.jpg', NOW()),
    (1, 'Test Photo 15', 'This is a test photo', 'image-15.jpg', NOW()),
    (1, 'Test Photo 16', 'This is a test photo', 'image-16.jpg', NOW());
