-- +migrate Up
CREATE TYPE public.gap_status_company AS ENUM(
    'draft', 'submitted'
);

ALTER TABLE public.companies ADD gap_status gap_status_company;

-- +migrate Down
ALTER TABLE public.companies DROP gap_status;
DROP TYPE IF EXISTS gap_status_company;