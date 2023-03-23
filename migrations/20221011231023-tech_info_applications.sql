-- +migrate Up

CREATE TABLE public.tech_info_applications (
    tech_info_application_uuid uuid,
    company_uuid uuid NOT NULL,
    company_website_uuid uuid NOT NULL,
    name VARCHAR(250) NOT NULL,
    env VARCHAR(250),
    account VARCHAR(250),
    mfa VARCHAR(250),
    has_permissions boolean,
    has_ids_ips boolean,
    is_whitelisted boolean,
    is_3rd_party_hosted boolean,
    created_at timestamptz NULL default NOW(),
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    PRIMARY KEY(tech_info_application_uuid),
    UNIQUE (company_uuid, name),
    CONSTRAINT fk_companies FOREIGN KEY(company_uuid) REFERENCES companies(company_uuid),
    CONSTRAINT fk_company_websites FOREIGN KEY(company_website_uuid) REFERENCES company_websites(company_website_uuid)
);
-- +migrate Down
DROP TABLE IF EXISTS public.tech_info_applications;
