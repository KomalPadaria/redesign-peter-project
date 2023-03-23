-- +migrate Up
ALTER TABLE public.application_envs ADD name text;

-- +migrate Down

ALTER TABLE public.application_envs DROP name;

