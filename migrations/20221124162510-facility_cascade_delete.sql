-- +migrate Up
ALTER TABLE public.tech_info_ip_ranges
DROP CONSTRAINT fk_company_facilities,
ADD CONSTRAINT fk_company_facilities
	FOREIGN KEY (company_facility_uuid)
	REFERENCES company_facilities(company_facility_uuid)
	ON DELETE CASCADE;

ALTER TABLE public.tech_info_wireless
DROP CONSTRAINT fk_company_facilities,
ADD CONSTRAINT fk_company_facilities
	FOREIGN KEY (company_facility_uuid)
	REFERENCES company_facilities(company_facility_uuid)
	ON DELETE CASCADE;
-- +migrate Down
ALTER TABLE public.tech_info_ip_ranges
DROP CONSTRAINT fk_company_facilities,
ADD CONSTRAINT fk_company_facilities
	FOREIGN KEY (company_facility_uuid)
	REFERENCES company_facilities(company_facility_uuid);

ALTER TABLE public.tech_info_wireless
DROP CONSTRAINT fk_company_facilities,
ADD CONSTRAINT fk_company_facilities
	FOREIGN KEY (company_facility_uuid)
	REFERENCES company_facilities(company_facility_uuid);
