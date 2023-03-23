-- +migrate Up
ALTER TABLE public.users ADD CONSTRAINT email_unique_key UNIQUE (email);

alter table public.users drop constraint if exists username_unique;
alter table public.users drop constraint if exists username_unique_constraint;
alter table public.users drop constraint if exists users_external_id_key;

-- +migrate Down
alter table public.users drop constraint if exists email_unique_key;
