-- +migrate Up
ALTER TABLE public.tech_info_ip_ranges DROP CONSTRAINT IF EXISTS uk_company_company_facility;

-- +migrate Down
ALTER TABLE public.tech_info_ip_ranges ADD CONSTRAINT uk_company_company_facility UNIQUE (company_uuid, company_facility_uuid);
