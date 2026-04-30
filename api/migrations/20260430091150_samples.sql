-- Create "samples" table
CREATE TABLE "samples" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "created_at" timestamptz NULL DEFAULT now(),
  "updated_at" timestamptz NULL DEFAULT now(),
  PRIMARY KEY ("id")
);
-- Create index "idx_sample_created_at" to table: "samples"
CREATE INDEX "idx_sample_created_at" ON "samples" ("created_at");
-- Create index "idx_sample_name" to table: "samples"
CREATE INDEX "idx_sample_name" ON "samples" ("name");

INSERT INTO "samples" ("name") VALUES
  ('sample-01'),
  ('sample-02'),
  ('sample-03'),
  ('sample-04'),
  ('sample-05'),
  ('sample-06'),
  ('sample-07'),
  ('sample-08'),
  ('sample-09'),
  ('sample-10');
