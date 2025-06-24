-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."users";

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "name" text NOT NULL,
    "email" text NOT NULL,
    "password_hashed" text NOT NULL,
    "last_login" timestamp,
    "time_created" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "time_confirmed" timestamp,
    "time_deleted" timestamp,
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX users_email_key ON public.users USING btree (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."users";
-- +goose StatementEnd
