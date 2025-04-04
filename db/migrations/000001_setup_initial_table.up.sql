CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar,
  "email" varchar,
  "password" varchar,
  "timezone" varchar,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "activities" (
  "id" serial PRIMARY KEY,
  "user_id" integer,
  "name" varchar,
  "target" integer,
  "start_time" time,
  "recurrence_period" integer,
  "last_alert_sent_at" timestamp,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "activity_histories" (
  "id" serial PRIMARY KEY,
  "activity_id" integer,
  "target" integer,
  "achieved" integer,
  "achieved_at" timestamp,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "activities" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "activity_histories" ADD FOREIGN KEY ("activity_id") REFERENCES "activities" ("id");
