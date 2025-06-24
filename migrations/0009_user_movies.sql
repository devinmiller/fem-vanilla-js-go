-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."user_movies";

-- Table Definition
CREATE TABLE "public"."user_movies" (
    "user_id" int4 NOT NULL,
    "movie_id" int4 NOT NULL,
    "relation_type" text NOT NULL CHECK (relation_type = ANY (ARRAY['favorite'::text, 'watchlist'::text])),
    "time_added" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user_id","movie_id","relation_type")
);

ALTER TABLE "public"."user_movies" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE;
ALTER TABLE "public"."user_movies" ADD FOREIGN KEY ("movie_id") REFERENCES "public"."movies"("id") ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."user_movies";
-- +goose StatementEnd
