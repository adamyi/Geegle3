CREATE TABLE `users` (
    `username` VARCHAR(255),
    `password` TEXT,
    `avatar` LONGBLOB,
    PRIMARY KEY (`username`)
);

CREATE TABLE `apis` (
    `name` VARCHAR(255),
    `owner` VARCHAR(255),
    `origin` TEXT,
    `description` BLOB,
    `internal` BOOLEAN,
    `public_access` BOOLEAN,
    `dependsOn` TEXT,
    `key` CHAR(36),
    PRIMARY KEY (`name`)
);

CREATE TABLE `user_api` (
    `user` VARCHAR(255),
    `api` VARCHAR(255)
);

INSERT INTO `users` VALUES ('meme@geegle.org', 'b0439fae31f8cbba6294af86234d5a28', '/images/1.png');
