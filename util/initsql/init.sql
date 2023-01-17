CREATE TABLE `video` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `author_id` int NOT NULL,
  `publish_time` datetime NOT NULL DEFAULT "now()",
  `file_path` varchar(32) NOT NULL,
  `cover_path` varchar(32) NOT NULL,
  `favorite_count` int DEFAULT 0,
  `comment_count` int DEFAULT 0,
  `title` varchar(20) NOT NULL
);

CREATE TABLE `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT "now()",
  `name` varchar(20) UNIQUE NOT NULL,
  `password` char(32) NOT NULL,
  `salt` char(32) NOT NULL,
  `follow_count` int NOT NULL DEFAULT 0,
  `follower_count` int NOT NULL DEFAULT 0
);

CREATE TABLE `follow` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `follow_time` datetime NOT NULL DEFAULT "now()",
  `from_user_id` int NOT NULL,
  `to_user_id` int NOT NULL
);

CREATE TABLE `comment` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `comment_time` datetime NOT NULL DEFAULT "now()",
  `user_id` int NOT NULL,
  `video_id` int NOT NULL,
  `content` text NOT NULL
);

CREATE TABLE `message` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `message_time` datetime NOT NULL DEFAULT "now()",
  `from_user_id` int NOT NULL,
  `to_user_id` int NOT NULL,
  `content` text NOT NULL
);

CREATE TABLE `like` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `like_time` datetime NOT NULL DEFAULT "now()",
  `user_id` int NOT NULL,
  `video_id` int NOT NULL
);

ALTER TABLE `video` ADD FOREIGN KEY (`author_id`) REFERENCES `user` (`id`);

ALTER TABLE `follow` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);

ALTER TABLE `follow` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);

ALTER TABLE `comment` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `comment` ADD FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);

ALTER TABLE `message` ADD FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`);

ALTER TABLE `message` ADD FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`);

ALTER TABLE `like` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `like` ADD FOREIGN KEY (`video_id`) REFERENCES `video` (`id`);
