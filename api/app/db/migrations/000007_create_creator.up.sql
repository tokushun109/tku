CREATE TABLE IF NOT EXISTS creator(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL,
    introduction TEXT(1000),
    mime_type VARCHAR(30),
    logo VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);