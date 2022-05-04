BEGIN;
CREATE TABLE `articles` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `content` LONGTEXT,
    `user_id` INT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT current_timestamp(),
    `updated_at` DATETIME NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`),
    INDEX `user_id` (`user_id`),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
COMMIT;
