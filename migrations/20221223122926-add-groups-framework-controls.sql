-- +migrate Up
ALTER TABLE public.framework_controls ADD groups TEXT [];
-- +migrate Down
ALTER TABLE public.framework_controls DROP groups;