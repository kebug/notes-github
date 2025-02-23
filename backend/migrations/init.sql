CREATE DATABASE IF NOT EXISTS note_sync CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE note_sync;

-- 如果需要手动创建表（通常由GORM自动处理）
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    username VARCHAR(255) NOT NULL,
    github_id BIGINT NOT NULL,
    avatar_url VARCHAR(255) NOT NULL,
    github_token VARCHAR(255) NOT NULL,
    INDEX idx_users_deleted_at (deleted_at),
    UNIQUE INDEX idx_users_github_id (github_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS articles (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    git_path VARCHAR(255) NULL,
    INDEX idx_articles_deleted_at (deleted_at),
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS questions (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    deleted_at DATETIME(3) NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    answer TEXT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    issue_url VARCHAR(255) NULL,
    INDEX idx_questions_deleted_at (deleted_at),
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci; 