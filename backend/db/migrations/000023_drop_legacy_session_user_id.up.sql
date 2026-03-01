-- session は uuid ベースで user を参照する前提に切り替わったため、
-- 旧 user_id 外部キーを削除する。

ALTER TABLE session
DROP FOREIGN KEY session_ibfk_1;

ALTER TABLE session
DROP COLUMN user_id;
