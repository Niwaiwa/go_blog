BEGIN;
CREATE TABLE `users` (
    `id` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL,
    `account` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT current_timestamp(),
    `updated_at` DATETIME NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`),
    UNIQUE INDEX `account` (`account`),
    UNIQUE INDEX `username` (`username`)
);
COMMIT;
