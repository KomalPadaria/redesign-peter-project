-- +migrate Up
CREATE TYPE public.user_group_type AS ENUM(
    'customer', 'superadmin', 'engineer', 'csc'
);

ALTER TABLE public.users ADD user_group user_group_type;

-- +migrate Down
ALTER TABLE public.users DROP user_group;
DROP TYPE IF EXISTS user_group_type;
