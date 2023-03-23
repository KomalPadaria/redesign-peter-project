-- +migrate Up
CREATE TYPE public.signature_status AS ENUM(
    'Signed', 'Not Signed'
);

CREATE TABLE public.signatures (
    signature_uuid uuid NOT NULL,
    "name" text NOT NULL,
    data text NULL,
    company_types industry_type[] NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT signatures_pkey PRIMARY KEY (signature_uuid)
);

CREATE TABLE public.company_signatures (
    company_signature_uuid uuid NOT NULL,
    signature_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    "name" text,
    status signature_status DEFAULT 'Not Signed',
    signature_data jsonb NULL,
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT company_signatures_pkey PRIMARY KEY (company_signature_uuid)
);

ALTER TABLE public.company_signatures ADD CONSTRAINT fk_company FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);
ALTER TABLE public.company_signatures ADD CONSTRAINT fk_signature FOREIGN KEY (signature_uuid) REFERENCES public.signatures(signature_uuid);

-- +migrate Down
ALTER TABLE public.company_signatures DROP CONSTRAINT IF EXISTS fk_company;
ALTER TABLE public.company_signatures DROP CONSTRAINT IF EXISTS fk_signature;

DROP TABLE IF EXISTS public.company_signatures;
DROP TABLE IF EXISTS public.signatures;