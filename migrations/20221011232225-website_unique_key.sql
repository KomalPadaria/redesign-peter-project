-- +migrate Up
ALTER TABLE public.company_websites ADD CONSTRAINT company_id_url UNIQUE (company_uuid, url);
-- +migrate Down
ALTER TABLE public.company_websites DROP CONSTRAINT IF EXISTS company_id_url;
