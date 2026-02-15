ALTER TABLE
    product
ADD
    (
        price INT NOT NULL DEFAULT 0,
        is_active BOOLEAN NOT NULL DEFAULT 0
    );