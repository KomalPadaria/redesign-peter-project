-- +migrate Up
ALTER TABLE public.companies ADD campaign_users jsonb NULL;

-- +migrate Down
ALTER TABLE public.companies DROP campaign_users;