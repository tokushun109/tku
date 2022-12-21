CREATE TABLE IF NOT EXISTS contact(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL,
    company VARCHAR(30),
    phone_number VARCHAR(30),
    email VARCHAR(50) NOT NULL,
    content TEXT(1000) NOT NULL,
    is_read BOOLEAN NOT NULL DEFAULT 0,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);