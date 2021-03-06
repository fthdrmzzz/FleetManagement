CREATE TYPE "bags_status" AS ENUM (
  'created',
  'unloaded',
  'loaded'
);

CREATE TYPE "packages_status" AS ENUM (
  'created',
  'unloaded',
  'loaded_to_bag',
  'loaded'
);

CREATE TABLE "bags" (
  "barcode" varchar PRIMARY KEY,
  "bag_status" bags_status NOT NULL,
  "delivery_id" int NOT NULL
);

CREATE TABLE "packages" (
  "barcode" varchar PRIMARY KEY,
  "package_status" packages_status NOT NULL,
  "package_weight" int NOT NULL,
  "delivery_id" int NOT NULL
);

CREATE TABLE "vehicles" (
  "plate" varchar PRIMARY KEY NOT NULL
);

CREATE TABLE "delivery_points" (
  "id" int PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "package_bag" (
  "bag_barcode" varchar,
  "package_barcode" varchar,
  PRIMARY KEY ("bag_barcode", "package_barcode")
);

CREATE TABLE "bag_vehicle" (
  "vehicle_plate" varchar,
  "bag_barcode" varchar,

  PRIMARY KEY ("bag_barcode", "vehicle_plate")
);

CREATE TABLE "package_vehicle" (
  "vehicle_plate" varchar,
  "package_barcode" varchar,
  PRIMARY KEY ("package_barcode", "vehicle_plate")
);

ALTER TABLE "bags" ADD FOREIGN KEY ("delivery_id") REFERENCES "delivery_points" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("delivery_id") REFERENCES "delivery_points" ("id");

ALTER TABLE "package_bag" ADD FOREIGN KEY ("bag_barcode") REFERENCES "bags" ("barcode");

ALTER TABLE "package_bag" ADD FOREIGN KEY ("package_barcode") REFERENCES "packages" ("barcode");

ALTER TABLE "bag_vehicle" ADD FOREIGN KEY ("bag_barcode") REFERENCES "bags" ("barcode");

ALTER TABLE "bag_vehicle" ADD FOREIGN KEY ("vehicle_plate") REFERENCES "vehicles" ("plate");

ALTER TABLE "package_vehicle" ADD FOREIGN KEY ("package_barcode") REFERENCES "packages" ("barcode");

ALTER TABLE "package_vehicle" ADD FOREIGN KEY ("vehicle_plate") REFERENCES "vehicles" ("plate");
