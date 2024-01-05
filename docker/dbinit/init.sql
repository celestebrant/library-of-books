CREATE TABLE books
(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `creation_time` timestamp DEFAULT NULL,
    `update_time` timestamp DEFAULT NULL,
    `title` VARCHAR(255) DEFAULT NULL,
    `author` VARCHAR(255) DEFAULT NULL,
    PRIMARY KEY (id)
);