-- +migrate Up
ALTER TABLE public.companies ADD knowbe4_token TEXT;
ALTER TABLE public.companies ADD rapid7_site_ids TEXT [];

-- +migrate Down
ALTER TABLE public.companies DROP knowbe4_token;
ALTER TABLE public.companies DROP rapid7_site_ids;