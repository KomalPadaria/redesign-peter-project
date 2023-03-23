-- +migrate Up
CREATE TYPE public.tech_info_ip_ranges_status AS ENUM(
    'Active', 'Inactive'
);

CREATE TABLE public.tech_info_ip_ranges (
    tech_info_ip_range_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    company_facility_uuid uuid NOT NULL,
    ip_address VARCHAR(50) NOT NULL,
    ip_size INT NOT NULL,
    is_external boolean NOT NULL,
    status tech_info_ip_ranges_status DEFAULT 'Active',
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT tech_info_ip_ranges_pkey PRIMARY KEY (tech_info_ip_range_uuid)
);

ALTER TABLE public.tech_info_ip_ranges ADD CONSTRAINT fk_company FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);
ALTER TABLE public.tech_info_ip_ranges ADD CONSTRAINT fk_company_facilities FOREIGN KEY (company_facility_uuid) REFERENCES public.company_facilities(company_facility_uuid);

ALTER TABLE public.tech_info_ip_ranges ADD CONSTRAINT uk_company_company_facility UNIQUE (company_uuid, company_facility_uuid);


-- +migrate Down
ALTER TABLE public.tech_info_ip_ranges DROP CONSTRAINT IF EXISTS fk_company;
ALTER TABLE public.tech_info_ip_ranges DROP CONSTRAINT IF EXISTS fk_company_facilities;

DROP TABLE IF EXISTS public.tech_info_ip_ranges;
