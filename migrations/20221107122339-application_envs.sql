-- +migrate Up
CREATE TYPE public.application_env_type AS ENUM(
    'Production', 'Staging','Testing','Development'
);

CREATE TYPE public.mfa AS ENUM(
    'Microsoft Authenticator', 'Google Authenticator','Okta Verify','Duo Mobile', 'Last Pass Authenticator', 'Others'
);

CREATE TYPE public.hosting_provider AS ENUM(
    'Amazon Web Services', 'Google Cloud Platform','Microsoft Azure','Others'
);

CREATE TYPE public.ids_ips_solution AS ENUM(
    'Microsoft Authenticator', 'Zscaler','CrowdStrike','Cisco', 'Fortigate', 'Splunk', 'Zeek', 'SolarWinds', 'Fidelis', 'Snort', 'BlueVector', 'Others'
);

CREATE TYPE public.application_envs_status AS ENUM(
    'Active', 'Inactive'
);

CREATE TABLE public.application_envs (
    application_env_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    tech_info_application_uuid uuid NOT NULL,
    type application_env_type NOT NULL,
    description text NULL,
    url text NULL,
    hosting_provider_type hosting_provider NULL,
    hosting_provider text NULL,
    mfa_type mfa NULL,
    mfa text NULL,
    ids_ips_type ids_ips_solution NULL,
    ids_ips text,
    status application_envs_status DEFAULT 'Inactive',
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT application_envs_pkey PRIMARY KEY (application_env_uuid)
);

ALTER TABLE public.application_envs ADD CONSTRAINT fk_company FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);
ALTER TABLE public.application_envs ADD CONSTRAINT fk_tech_info_application FOREIGN KEY (tech_info_application_uuid) REFERENCES public.tech_info_applications(tech_info_application_uuid);

CREATE TYPE public.tech_info_applications_status AS ENUM(
    'Active', 'Inactive'
);

CREATE TYPE public.application_type AS ENUM(
    'Website', 'Desktop','SaaS'
);

ALTER TABLE public.tech_info_applications DROP CONSTRAINT IF EXISTS name_env_website_unique_constraint;

ALTER TABLE public.tech_info_applications
DROP COLUMN company_website_uuid,
DROP COLUMN env,
DROP COLUMN account,
DROP COLUMN mfa,
DROP COLUMN has_permissions,
DROP COLUMN has_ids_ips,
DROP COLUMN is_whitelisted,
DROP COLUMN is_3rd_party_hosted;

ALTER TABLE public.tech_info_applications
ADD COLUMN status tech_info_applications_status DEFAULT 'Inactive' NOT NULL,
ADD COLUMN "type" application_type NOT NULL;

-- +migrate Down
ALTER TABLE public.tech_info_applications
ADD COLUMN company_website_uuid,
ADD COLUMN env VARCHAR(250),
ADD COLUMN account VARCHAR(250),
ADD COLUMN mfa VARCHAR(250),
ADD COLUMN has_permissions boolean,
ADD COLUMN has_ids_ips boolean,
ADD COLUMN is_whitelisted boolean,
ADD COLUMN is_3rd_party_hosted boolean;

ALTER TABLE public.tech_info_applications
DROP COLUMN status,
DROP COLUMN "type";

ALTER TABLE public.tech_info_applications ADD CONSTRAINT name_env_website_unique_constraint UNIQUE (name, env, company_website_uuid);
ALTER TABLE public.application_envs DROP CONSTRAINT IF EXISTS fk_company;
ALTER TABLE public.application_envs ADD CONSTRAINT IF EXISTS fk_tech_info_application;

DROP TABLE IF EXISTS public.application_envs;