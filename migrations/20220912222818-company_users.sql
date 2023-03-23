-- +migrate Up
CREATE TYPE public.company_user_role AS ENUM(
    'admin',
    'manager'
);

CREATE TABLE public.company_users (
    company_uuid uuid,
    user_uuid uuid,
    role company_user_role,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT company_users_pkey PRIMARY KEY (company_uuid, user_uuid)
);

create trigger update_updated_at before
    update
    on
        public.company_users for each row execute function trigger_update_timestamp();

ALTER TABLE public.company_users ADD CONSTRAINT fk_companies FOREIGN KEY (company_uuid) REFERENCES public.companies(company_uuid);
ALTER TABLE public.company_users ADD CONSTRAINT fk_users FOREIGN KEY (user_uuid) REFERENCES public.users(user_uuid);

-- +migrate Down
DROP TABLE IF EXISTS public.company_users;