-- +migrate Up
ALTER TABLE public.users ADD current_company_uuid uuid;

ALTER TABLE public.users ADD CONSTRAINT fk_companies FOREIGN KEY (current_company_uuid) REFERENCES public.companies(company_uuid);

alter table public.users alter column username set not null;
alter table public.users add constraint username_unique unique (username);

-- +migrate Down
ALTER TABLE public.users DROP CONSTRAINT fk_companies;

ALTER TABLE public.users DROP current_company_uuid;

alter table public.users drop constraint username_unique;
