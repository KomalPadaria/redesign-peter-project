-- +migrate Up
CREATE TABLE public.meetings (
    meetings_uuid uuid NOT NULL,
    "name" text NOT NULL,
    description text NOT NULL,
    duration interval NOT NULL,
    data jsonb NULL,
    company_types industry_type[] NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NULL,
    created_by uuid NOT NULL,
    updated_by uuid NULL,
    CONSTRAINT meetings_pkey PRIMARY KEY (meetings_uuid)
);

CREATE TABLE public.company_meetings (
    company_meetings_uuid uuid NOT NULL,
    meetings_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    start_at timestamptz NOT NULL,
    utc_start_at timestamptz NOT NULL,
    host text,
    "name" text,
    "data" jsonb NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NULL,
    created_by uuid NOT NULL,
    updated_by uuid NULL,
    CONSTRAINT company_meetings_pkey PRIMARY KEY (company_meetings_uuid)
);

ALTER TABLE public.company_meetings ADD CONSTRAINT fk_companies FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);
ALTER TABLE public.company_meetings ADD CONSTRAINT fk_meetings FOREIGN KEY (meetings_uuid) REFERENCES public.meetings(meetings_uuid);

-- +migrate Down
ALTER TABLE public.company_meetings DROP CONSTRAINT IF EXISTS fk_companies;
ALTER TABLE public.company_meetings DROP CONSTRAINT IF EXISTS fk_meetings;

DROP TABLE IF EXISTS public.company_meetings;
DROP TABLE IF EXISTS public.meetings;
