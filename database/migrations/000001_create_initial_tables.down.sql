-- 000001_create_initial_tables.down.sql

-- 作成したテーブルを削除 (依存関係に注意して逆順にDROP)
DROP TABLE IF EXISTS play_sessions;
DROP TABLE IF EXISTS fixed_events;
DROP TABLE IF EXISTS games;
DROP TABLE IF EXISTS users;

-- トリガー関数も削除
DROP FUNCTION IF EXISTS update_updated_at_column;