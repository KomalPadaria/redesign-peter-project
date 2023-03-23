-- +migrate Up
CREATE TYPE public.company_users_status AS ENUM(
    'ACTIVE',
    'INACTIVE',
    'PENDING'
);
ALTER TABLE public.company_users ADD status company_users_status DEFAULT 'ACTIVE';

-- +migrate Down
ALTER TABLE public.company_users DROP status;
