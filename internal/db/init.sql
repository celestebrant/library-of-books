CREATE TABLE books
(
    `id` VARCHAR(26),
    `creation_time` DATETIME(6) DEFAULT NULL,
    `update_time` DATETIME(6) DEFAULT NULL,
    `title` VARCHAR(255) DEFAULT NULL,
    `author` VARCHAR(255) DEFAULT NULL,
    PRIMARY KEY (id)
);