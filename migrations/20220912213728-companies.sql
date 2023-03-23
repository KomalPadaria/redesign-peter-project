-- +migrate Up
CREATE TYPE public.industry_type AS ENUM(
    'Entertainment'
);

CREATE TABLE public.companies (
    company_uuid uuid NOT NULL,
    "name" text NOT NULL,
    industry_type industry_type[] NULL,
    onboarding jsonb NULL,
    address jsonb NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT companies_pkey PRIMARY KEY (company_uuid)
);

CREATE TRIGGER update_updated_at
    BEFORE UPDATE ON companies
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_update_timestamp();

-- +migrate Down
DROP TABLE IF EXISTS public.companies;