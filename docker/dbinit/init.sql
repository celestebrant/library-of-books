CREATE TABLE books
(
    `id` VARCHAR(26),
    `creation_time` timestamp DEFAULT NULL,
    `update_time` timestamp DEFAULT NULL,
    `title` VARCHAR(255) DEFAULT NULL,
    `author` VARCHAR(255) DEFAULT NULL,
    PRIMARY KEY (id)
);