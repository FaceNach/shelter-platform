-- +goose Up
ALTER TABLE shelters
ADD COLUMN contact_email TEXT NOT NULL;

-- +goose Down
ALTER TABLE shelters
DROP COLUMN contact_email;


