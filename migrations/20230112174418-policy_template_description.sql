-- +migrate Up
ALTER TABLE public.policy_templates ADD description text;

-- +migrate Down
ALTER TABLE public.policy_templates DROP description;