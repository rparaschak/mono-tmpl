CREATE EXTENSION IF NOT EXISTS postgis;

-- Modify "samples" table
ALTER TABLE "samples" ADD COLUMN "geolocation" geography(Point,4326) NULL;

UPDATE "samples"
SET "geolocation" = ST_SetSRID(ST_MakePoint(21.0122, 52.2297), 4326)::geography
WHERE "name" = 'sample-01';

UPDATE "samples"
SET "geolocation" = ST_SetSRID(ST_MakePoint(19.9445, 50.0497), 4326)::geography
WHERE "name" = 'sample-02';

UPDATE "samples"
SET "geolocation" = ST_SetSRID(ST_MakePoint(17.0385, 51.1079), 4326)::geography
WHERE "name" = 'sample-03';

UPDATE "samples"
SET "geolocation" = ST_SetSRID(ST_MakePoint(18.6466, 54.3520), 4326)::geography
WHERE "name" = 'sample-04';

UPDATE "samples"
SET "geolocation" = ST_SetSRID(ST_MakePoint(16.9252, 52.4064), 4326)::geography
WHERE "name" = 'sample-05';

UPDATE "samples"
SET "geolocation" = ST_SetSRID(ST_MakePoint(19.4550, 51.7592), 4326)::geography
WHERE "name" = 'sample-06';

UPDATE "samples"
SET "geolocation" = ST_SetSRID(ST_MakePoint(22.5684, 51.2465), 4326)::geography
WHERE "name" = 'sample-07';

UPDATE "samples"
SET "geolocation" = ST_SetSRID(ST_MakePoint(18.0084, 53.1235), 4326)::geography
WHERE "name" = 'sample-08';

UPDATE "samples"
SET "geolocation" = ST_SetSRID(ST_MakePoint(23.1688, 53.1325), 4326)::geography
WHERE "name" = 'sample-09';

UPDATE "samples"
SET "geolocation" = ST_SetSRID(ST_MakePoint(14.5528, 53.4285), 4326)::geography
WHERE "name" = 'sample-10';

ALTER TABLE "samples" ALTER COLUMN "geolocation" SET NOT NULL;

CREATE INDEX "idx_sample_geolocation_gist" ON "samples" USING gist ("geolocation");
