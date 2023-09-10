
ALTER TABLE IF EXISTS "role_permission" DROP CONSTRAINT IF EXISTS "role_permission_permission_id_fkey"
ALTER TABLE IF EXISTS "role_permission" DROP CONSTRAINT IF EXISTS "role_permission_role_id_fkey"
DROP TABLE IF EXISTS role_permission CASCADE;