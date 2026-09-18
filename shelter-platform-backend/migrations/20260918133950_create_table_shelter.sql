-- +goose Up
 CREATE TABLE shelters (
       id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       name TEXT NOT NULL,
       city TEXT NOT NULL,
       state TEXT NOT NULL,
       country TEXT NOT NULL,
       zip_code TEXT NOT NULL,
       phone_number TEXT NOT NULL,
       founded_at DATE,
       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
   );


-- +goose Down
DROP TABLE shelters;
