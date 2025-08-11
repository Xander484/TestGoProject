-- 20230701_add_columns.sql
ALTER TABLE workspaces
ADD COLUMN allow_add_client BOOLEAN DEFAULT TRUE,
ADD COLUMN allow_edit_client BOOLEAN DEFAULT TRUE;
