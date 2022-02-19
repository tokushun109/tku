ALTER TABLE
    sns
ADD
    icon VARCHAR(100);

ALTER TABLE
    sns
MODIFY
    icon TEXT
AFTER
    url;

ALTER TABLE
    skill_market
ADD
    icon VARCHAR(100);

ALTER TABLE
    skill_market
MODIFY
    icon TEXT
AFTER
    url;

ALTER TABLE
    sales_site
ADD
    icon VARCHAR(100);

ALTER TABLE
    sales_site
MODIFY
    icon TEXT
AFTER
    url;