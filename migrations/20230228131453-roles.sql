-- +migrate Up
CREATE TYPE public.company_type AS ENUM(
    'customer',
    'engineering'
);

ALTER TABLE public.companies ADD type company_type DEFAULT 'customer';

ALTER TYPE public.company_user_role ADD VALUE 'csc';
ALTER TYPE public.company_user_role ADD VALUE 'engineer';

ALTER TABLE public.company_users ALTER COLUMN role SET DEFAULT 'user';
ALTER TABLE public.users ALTER COLUMN user_group SET DEFAULT 'customer';

ALTER TABLE public.companies DROP CONSTRAINT companies_external_id_key;
ALTER TABLE public.companies ADD CONSTRAINT companies_external_id_key UNIQUE (name, external_id, type);
-- +migrate Down
