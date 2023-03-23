-- +migrate Up
ALTER TABLE public.tech_info_wireless ADD has_permissions boolean DEFAULT false;
-- +migrate Down
ALTER TABLE public.tech_info_wireless DROP has_permissions;
