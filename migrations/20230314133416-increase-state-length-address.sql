-- +migrate Up
ALTER TABLE public.company_address
ALTER COLUMN state TYPE VARCHAR(250);
-- +migrate Down
