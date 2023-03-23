-- +migrate Up
ALTER TABLE public.policies ADD CONSTRAINT uk_name_company UNIQUE (name, company_uuid);

-- +migrate Down
ALTER TABLE public.policies DROP CONSTRAINT IF EXISTS uk_name_company;
