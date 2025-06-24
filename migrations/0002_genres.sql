-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."genres";

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS genres_id_seq;

-- Table Definition
CREATE TABLE "public"."genres" (
    "id" int4 NOT NULL DEFAULT nextval('genres_id_seq'::regclass),
    "name" text NOT NULL,
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX genres_name_key ON public.genres USING btree (name);

INSERT INTO "public"."genres" ("id", "name") VALUES
(1, 'Adventure'),
(2, 'Action'),
(3, 'Thriller'),
(4, 'Crime'),
(6, 'Fantasy'),
(11, 'Family'),
(18, 'Science Fiction'),
(25, 'Animation'),
(26, 'Western'),
(35, 'Comedy'),
(58, 'Drama'),
(76, 'Romance'),
(118, 'Horror'),
(264, 'Mystery'),
(346, 'History'),
(350, 'War'),
(868, 'Music'),
(1545, 'Documentary'),
(6323, 'Foreign'),
(8892, 'TV Movie');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."genres";
-- +goose StatementEnd
