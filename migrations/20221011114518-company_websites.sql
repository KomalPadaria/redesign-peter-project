-- +migrate Up
CREATE TABLE public.company_websites (
    company_website_uuid uuid,
    company_uuid uuid NOT NULL,
    url TEXT NOT NULL,
    industry_type industry_type NULL,
    zip VARCHAR(10),
    country VARCHAR(2),
    state VARCHAR(2),
    city VARCHAR(250),
    address1 VARCHAR(250),
    address2 VARCHAR(250),
    created_at timestamptz NULL DEFAULT now(),
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    PRIMARY KEY(company_website_uuid),
    CONSTRAINT fk_companies FOREIGN KEY(company_uuid) REFERENCES companies(company_uuid)
);

-- +migrate Down

DROP TABLE IF EXISTS public.company_websites;


