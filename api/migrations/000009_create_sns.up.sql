CREATE TABLE IF NOT EXISTS sns(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30),
    url VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);