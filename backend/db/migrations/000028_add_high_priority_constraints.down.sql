ALTER TABLE session
DROP FOREIGN KEY fk_session_user_uuid;

DROP INDEX idx_session_user_uuid ON session;
CREATE INDEX idx_session_user_uuid ON session (user_uuid);

ALTER TABLE session
ADD CONSTRAINT fk_session_user_uuid
FOREIGN KEY (user_uuid) REFERENCES user (uuid);

DROP INDEX idx_sales_site_active_name ON sales_site;
ALTER TABLE sales_site DROP COLUMN active_name;

DROP INDEX idx_tag_active_name ON tag;
ALTER TABLE tag DROP COLUMN active_name;

DROP INDEX idx_target_active_name ON target;
ALTER TABLE target DROP COLUMN active_name;

DROP INDEX idx_category_active_name ON category;
ALTER TABLE category DROP COLUMN active_name;
