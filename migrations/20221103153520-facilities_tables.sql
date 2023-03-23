-- +migrate Up
CREATE TYPE public.company_address_status AS ENUM(
    'Active', 'Inactive'
);

CREATE TABLE public.company_address (
    company_address_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    zip VARCHAR(10),
    country VARCHAR(2),
    state VARCHAR(2),
    city VARCHAR(250),
    address1 VARCHAR(250),
    address2 VARCHAR(250),
    status company_address_status DEFAULT 'Active',
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT company_address_pkey PRIMARY KEY (company_address_uuid)
);

CREATE TYPE public.company_facility_types AS ENUM(
    'Corporate Office', 'Studio', 'Data Center', 'Manufacturing Facility', 'Warehouse', 'Others'
);

CREATE TYPE public.company_facility_status AS ENUM(
    'Active', 'Inactive'
);

CREATE TABLE public.company_facilities (
    company_facility_uuid uuid NOT NULL,
    company_address_uuid uuid NOT NULL,
    company_uuid uuid NOT NULL,
    "name" text NOT NULL,
    "type" company_facility_types NULL,
    status company_facility_status DEFAULT 'Active',
    created_at timestamptz NULL,
    updated_at timestamptz NULL,
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT company_facilities_pkey PRIMARY KEY (company_facility_uuid)
);

ALTER TABLE public.company_facilities ADD CONSTRAINT fk_company FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);
ALTER TABLE public.company_facilities ADD CONSTRAINT fk_company_address FOREIGN KEY (company_address_uuid) REFERENCES public.company_address(company_address_uuid);
ALTER TABLE public.company_address ADD CONSTRAINT fk_company FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);

ALTER TABLE public.company_facilities ADD CONSTRAINT unique_company_address_type UNIQUE ("type",company_address_uuid);

-- +migrate Down
ALTER TABLE public.company_facilities DROP CONSTRAINT IF EXISTS fk_company;
ALTER TABLE public.company_facilities DROP CONSTRAINT IF EXISTS fk_company_address;
ALTER TABLE public.company_address DROP CONSTRAINT IF EXISTS fk_company;

DROP TABLE IF EXISTS public.company_facilities;
DROP TABLE IF EXISTS public.company_address;