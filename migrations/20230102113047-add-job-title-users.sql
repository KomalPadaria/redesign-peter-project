-- +migrate Up
ALTER TABLE public.users ADD job text;

-- +migrate Down
ALTER TABLE public.users DROP job;

