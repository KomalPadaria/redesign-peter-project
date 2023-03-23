-- +migrate Up
CREATE TYPE public.tech_info_wireless_protocol_type AS ENUM(
    '802.11ax (Wi-Fi 6)', '802.11ac (Wi-Fi 5)', '802.11n (Wi-Fi 4)', '802.11g', '802.11a', '802.11b'
);

CREATE TYPE public.tech_info_wireless_security_type AS ENUM(
    'WPA3', 'WPA2', 'WPA1', 'WEP'
);

CREATE TYPE public.tech_info_wireless_status AS ENUM(
    'Active', 'Inactive'
);

ALTER TABLE public.tech_info_wireless ADD protocol_type tech_info_wireless_protocol_type;
ALTER TABLE public.tech_info_wireless ADD protocol text;
ALTER TABLE public.tech_info_wireless ADD security_type tech_info_wireless_security_type;
ALTER TABLE public.tech_info_wireless ADD security text;
ALTER TABLE public.tech_info_wireless ADD status tech_info_wireless_status  DEFAULT 'Active';
ALTER TABLE public.tech_info_wireless ADD company_facility_uuid uuid;
ALTER TABLE public.tech_info_wireless ADD description text;


ALTER TABLE public.tech_info_wireless DROP env;
ALTER TABLE public.tech_info_wireless DROP location;
ALTER TABLE public.tech_info_wireless DROP has_ids_ips;
ALTER TABLE public.tech_info_wireless DROP is_3rd_party_hosted;
ALTER TABLE public.tech_info_wireless DROP has_permissions;
ALTER TABLE public.tech_info_wireless DROP company_website_uuid;

ALTER TABLE public.tech_info_wireless DROP CONSTRAINT IF EXISTS fk_company_websites;
ALTER TABLE public.tech_info_wireless ADD CONSTRAINT fk_company_facilities FOREIGN KEY (company_facility_uuid) REFERENCES public.company_facilities(company_facility_uuid);

-- +migrate Down