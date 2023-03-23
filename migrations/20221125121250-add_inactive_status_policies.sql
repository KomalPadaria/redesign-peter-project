-- +migrate Up
ALTER TYPE public.policies_status ADD VALUE 'Inactive';
ALTER TABLE public.policies
    ADD COLUMN status_updated_by uuid NULL,
    ADD COLUMN status_updated_at timestamptz NULL;

ALTER TABLE public.policies ADD CONSTRAINT fk_status_updated_by_users FOREIGN KEY (status_updated_by) REFERENCES public.users(user_uuid);

-- +migrate Down
