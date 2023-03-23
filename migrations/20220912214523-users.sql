-- +migrate Up
CREATE TABLE public.users (
    user_uuid uuid NOT NULL,
    username text NULL,
    first_name text NULL,
    last_name text NULL,
    email text NULL,
    phone text NULL,
    is_first_login boolean DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    created_by uuid NULL,
    updated_by uuid NULL,
    CONSTRAINT users_pkey PRIMARY KEY (user_uuid)
);

CREATE TRIGGER update_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_update_timestamp();
-- +migrate Down

DROP TABLE IF EXISTS public.users;