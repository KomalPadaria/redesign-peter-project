-- +migrate Up
CREATE TYPE public.company_subscription_type AS ENUM(
    'active',
    'inactive'
);

ALTER TABLE public.company_subscriptions ALTER COLUMN status TYPE company_subscription_type USING status::company_subscription_type;

ALTER TABLE public.company_subscriptions ALTER COLUMN status SET DEFAULT 'active';

-- +migrate Down
