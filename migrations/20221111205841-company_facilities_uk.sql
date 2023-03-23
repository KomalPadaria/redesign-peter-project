-- +migrate Up
ALTER TABLE public.company_facilities ADD CONSTRAINT unique_company_address_name UNIQUE ("name",company_address_uuid);
ALTER TABLE public.company_facilities DROP CONSTRAINT IF EXISTS unique_company_address_type;
-- +migrate Down
