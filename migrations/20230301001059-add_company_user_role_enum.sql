-- +migrate Up
ALTER TYPE public.company_user_role ADD VALUE 'superadmin';


-- +migrate Down
