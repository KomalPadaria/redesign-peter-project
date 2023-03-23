-- +migrate Up
ALTER TABLE public.tech_info_applications DROP CONSTRAINT IF EXISTS uk_tech_info_applications_name_type;
ALTER TABLE public.tech_info_applications ADD CONSTRAINT uk_tech_info_applications_name_type UNIQUE (name, type, company_uuid);

-- +migrate Down
ALTER TABLE public.tech_info_applications DROP CONSTRAINT IF EXISTS uk_tech_info_applications_name_type;
