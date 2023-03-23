-- +migrate Up
ALTER TABLE public.tech_info_applications ADD CONSTRAINT name_env_website_unique_constraint UNIQUE (name, env, company_website_uuid);
ALTER TABLE public.tech_info_external_infra ADD CONSTRAINT env_website_unique_constraint UNIQUE (env, company_website_uuid);
ALTER TABLE public.tech_info_wireless ADD CONSTRAINT location_env_website_unique_constraint UNIQUE (location, env, company_website_uuid);

-- +migrate Down
ALTER TABLE public.tech_info_applications DROP CONSTRAINT IF EXISTS name_env_website_unique_constraint;
ALTER TABLE public.tech_info_external_infra DROP CONSTRAINT IF EXISTS env_website_unique_constraint;
ALTER TABLE public.tech_info_wireless DROP CONSTRAINT IF EXISTS location_env_website_unique_constraint;
