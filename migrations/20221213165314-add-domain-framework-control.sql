-- +migrate Up
ALTER TABLE public.framework_controls ADD domain text;

-- +migrate Down

ALTER TABLE public.framework_controls DROP domain;

