CREATE TABLE "Account" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "Entries" ADD FOREIGN KEY ("account_id") REFERENCES "Account" ("id");

ALTER TABLE "Transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "Account" ("id");

ALTER TABLE "Transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "Account" ("id");

CREATE INDEX ON "Account" ("owner");

CREATE INDEX ON "Entries" ("account_id");

CREATE INDEX ON "Transfers" ("from_account_id");

CREATE INDEX ON "Transfers" ("to_account_id");

CREATE INDEX ON "Transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "Entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "Transfers"."amount" IS 'must be positive';