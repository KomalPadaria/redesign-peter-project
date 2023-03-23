-- +migrate Up
ALTER TYPE public.hosting_provider ADD VALUE 'CoreWeave';
-- +migrate Down
