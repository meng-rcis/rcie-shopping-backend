\set userdb `echo "$APP_DB_USER"`
\set passdb `echo "$APP_DB_PASS"`
\set dbname `echo "$APP_DB_NAME"`
CREATE DATABASE IF NOT EXISTS :'dbname';
CREATE USER :'userdb' WITH ENCRYPTED PASSWORD :'passdb';
\c `shopping_db`

CREATE TABLE IF NOT EXISTS `user` (
    `id` UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    `first_name` VARCHAR(50) NOT NULL,
    `last_name` VARCHAR(50) NOT NULL,
    `email` VARCHAR(62) NOT NULL,
    `username` VARCHAR(50) NOT NULL,
    `password_hash` VARCHAR(255) NOT NULL,
    `password_salt` VARCHAR(32) NOT NULL,
    `mobile` VARCHAR(15),
    `address` VARCHAR(255),
    `role_id` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
)
