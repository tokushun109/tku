CREATE TABLE IF NOT EXISTS product(
    id INT PRIMARY KEY AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL UNIQUE,
    name VARCHAR(30) NOT NULL,
    description TEXT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);
