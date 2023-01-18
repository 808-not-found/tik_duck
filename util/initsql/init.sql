CREATE TABLE `video` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `author_id` int NOT NULL,
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `file_path` varchar(32) NOT NULL,
  `cover_path` varchar(32) NOT NULL,
  `favorite_count` int DEFAULT 0,
  `comment_count` int DEFAULT 0,
  `title` varchar(20) NOT NULL
);

CREATE TABLE `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `name` varchar(20) UNIQUE NOT NULL,
  `password` char(32) NOT NULL,
  `salt` char(32) NOT NULL,
  `follow_count` int NOT NULL DEFAULT 0,
  `follower_count` int NOT NULL DEFAULT 0
);

CREATE TABLE `follow` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `follow_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `from_user_id` int NOT NULL,
  `to_user_id` int NOT NULL
);

CREATE TABLE `comment` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `comment_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_id` int NOT NULL,
  `video_id` int NOT NULL,
  `content` text NOT NULL
);

CREATE TABLE `message` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `message_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `from_user_id` int NOT NULL,
  `to_user_id` int NOT NULL,
  `content` text NOT NULL
);

CREATE TABLE `like` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `like_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_id` int NOT NULL,
  `video_id` int NOT NULL
);

CREATE INDEX `video_index_0` ON `video` (`id`);

CREATE INDEX `video_index_1` ON `video` (`author_id`);

CREATE INDEX `user_index_2` ON `user` (`id`);

CREATE INDEX `follow_index_3` ON `follow` (`id`);

CREATE INDEX `follow_index_4` ON `follow` (`from_user_id`);

CREATE INDEX `follow_index_5` ON `follow` (`to_user_id`);

CREATE INDEX `comment_index_6` ON `comment` (`id`);

CREATE INDEX `comment_index_7` ON `comment` (`user_id`);

CREATE INDEX `comment_index_8` ON `comment` (`video_id`);

CREATE INDEX `message_index_9` ON `message` (`id`);

CREATE INDEX `message_index_10` ON `message` (`from_user_id`);

CREATE INDEX `message_index_11` ON `message` (`to_user_id`);

CREATE INDEX `like_index_12` ON `like` (`id`);

CREATE INDEX `like_index_13` ON `like` (`user_id`);

CREATE INDEX `like_index_14` ON `like` (`video_id`);

ALTER TABLE `video` ADD FOREIGN KEY (`author_id`) REFERENCES `user` (`id`);

ALTER TABLE `follow` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);

ALTER TABLE `follow` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);

ALTER TABLE `comment` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `comment` ADD FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);

ALTER TABLE `message` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);

ALTER TABLE `message` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);

ALTER TABLE `like` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `like` ADD FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);
