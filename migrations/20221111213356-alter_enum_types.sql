-- +migrate Up
ALTER TYPE public.tech_info_wireless_protocol_type ADD VALUE 'Others' AFTER '802.11b';
ALTER TYPE public.tech_info_wireless_security_type ADD VALUE 'Others' AFTER 'WEP';
-- +migrate Down
