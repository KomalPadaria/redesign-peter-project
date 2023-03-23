-- +migrate Up
ALTER TABLE public.users ALTER COLUMN username SET NOT NULL;
alter table users drop constraint if exists username_unique_constraint;
ALTER TABLE users ADD CONSTRAINT username_unique_constraint UNIQUE (username);

-- +migrate Down
alter table users drop constraint if exists username_unique_constraint;
