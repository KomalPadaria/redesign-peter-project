-- +migrate Up
CREATE TYPE public.mfa_methods AS ENUM(
    'sms', 'app'
);

ALTER TABLE public.users ADD mfa mfa_methods [];

-- +migrate Down
ALTER TABLE public.users DROP mfa;
DROP TYPE IF EXISTS mfa_methods;