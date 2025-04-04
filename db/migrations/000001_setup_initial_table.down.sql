-- Hapus foreign key constraints terlebih dahulu
ALTER TABLE "activities" DROP CONSTRAINT "activities_user_id_fkey";
ALTER TABLE "activity_histories" DROP CONSTRAINT "activity_histories_activity_id_fkey";

-- Hapus tabel
DROP TABLE IF EXISTS "activity_histories";
DROP TABLE IF EXISTS "activities";
DROP TABLE IF EXISTS "users";
